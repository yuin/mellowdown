package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gobwas/glob"
	"github.com/pkg/errors"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/log"
	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/theme"
	"github.com/yuin/mellowdown/util"

	blackfriday "github.com/yuin/blackfriday/v2"
)

type Builder interface {
	LoadConfig() error
	Build() error
	Analyze() (Context, error)
	AnalyzeMarkdown(c Context, path string) (Context, error)
}

type builder struct {
	l               *lua.LState
	logger          log.Logger
	fileSystem      asset.FileSystem
	sourceDirectory string
	outputDirectory string
	renderers       []renderer.Renderer
	config          config
	theme           theme.Theme
	extras          map[string]bool
	extraPatterns   []glob.Glob
	ignores         map[string]bool
	ignorePatterns  []glob.Glob
	vars            map[string]interface{}
}

type Option func(*builder)

func WithSourceDirectory(path string) Option {
	return func(b *builder) {
		b.sourceDirectory, _ = filepath.Abs(path)
	}
}

func WithOutputDirectory(path string) Option {
	return func(b *builder) {
		b.outputDirectory, _ = filepath.Abs(path)
	}
}

func New(logger log.Logger, fileSystem asset.FileSystem, opts ...Option) Builder {
	b := &builder{
		logger:         logger,
		fileSystem:     fileSystem,
		l:              lua.NewState(),
		extras:         map[string]bool{},
		extraPatterns:  []glob.Glob{},
		ignores:        map[string]bool{},
		ignorePatterns: []glob.Glob{},
		vars:           map[string]interface{}{},
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *builder) LoadConfig() error {
	b.renderers = []renderer.Renderer{
		renderer.NewPlantUMLRenderer(),
		renderer.NewPPTRenderer(),
		renderer.NewTOCRenderer(),
		renderer.NewLabelRenderer(),
		renderer.NewSyntaxHighlightingRenderer(),
	}

	option := renderer.NewLuaOption(b.l)
	for _, r := range b.renderers {
		r.AddOption(option)
	}

	confFile := filepath.Join(b.sourceDirectory, "conf.lua")
	b.l.SetGlobal("tostring", b.l.NewFunction(toString))
	if err := b.l.DoFile(confFile); err != nil {
		return err
	}

	if err := readConfig(b.l, &b.config.Resource, "resource"); err != nil {
		return err
	}
	extras := []string{}
	for _, path := range b.config.Resource.Extras {
		v := util.SlashPath(filepath.Clean(path))
		b.extras[v] = true
		extras = append(extras, v)
	}
	b.config.Resource.Extras = extras

	for _, pat := range b.config.Resource.ExtraPatterns {
		g, err := glob.Compile(pat, '/')
		if err != nil {
			return err
		}
		b.extraPatterns = append(b.extraPatterns, g)
	}

	ignores := []string{}
	for _, path := range b.config.Resource.Ignores {
		v := util.SlashPath(filepath.Clean(path))
		b.ignores[v] = true
		ignores = append(ignores, v)
	}
	b.config.Resource.Ignores = ignores

	for _, pat := range b.config.Resource.IgnorePatterns {
		g, err := glob.Compile(pat, '/')
		if err != nil {
			return err
		}
		b.ignorePatterns = append(b.ignorePatterns, g)
	}

	if err := readConfig(b.l, &b.config.Site, "site"); err != nil {
		return err
	}
	if err := readConfig(b.l, &b.config.Theme, "theme"); err != nil {
		return err
	}

	themes := theme.NewThemes(b.fileSystem)
	if err := themes.Load(); err != nil {
		return err
	}

	var ok bool
	b.theme, ok = themes.Get(b.config.Theme.Name)
	if !ok {
		return fmt.Errorf("theme %s not found", b.config.Theme.Name)
	}

	lv := gluamapper.ToGoValue(b.l.GetGlobal("vars"), gluamapper.Option{
		NameFunc: gluamapper.ToUpperCamelCase,
	})
	for k, v := range lv.(map[interface{}]interface{}) {
		b.vars[k.(string)] = v
	}

	for _, r := range b.renderers {
		r.InitOption(option)
	}

	return nil
}

func (b *builder) AnalyzeMarkdown(con Context, path string) (Context, error) {
	if con == nil {
		con = newContext()
	}
	c := con.(*context)
	optList := []blackfriday.Option{blackfriday.WithExtensions(blackfriday.CommonExtensions | blackfriday.Functions)}
	parser := blackfriday.New(optList...)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	b.logger.Debug("parse markdown file")
	ast := parser.Parse(input)
	b.logger.Debug("done")
	title := ""
	c.addAST(path, ast)
	toc := renderer.TOC{}
	visited := map[*blackfriday.Node]bool{}
	labelIds := map[string]bool{}
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if _, ok := visited[node]; ok {
			return blackfriday.GoToNext
		}
		visited[node] = true
		switch node.Type {
		case blackfriday.Function:
			if node.FunctionData.Name == "label" {
				args := node.FunctionData.Arguments
				if len(args) < 1 {
					return blackfriday.GoToNext
				}
				text, ok := node.FunctionData.Arguments[0].(string)
				if !ok {
					return blackfriday.GoToNext
				}
				parts := strings.Split(text, "#")
				if len(parts) != 2 {
					return blackfriday.GoToNext
				}
				b.logger.Debug("find label: %s(#%s)", parts[0], parts[1])
				c.addLabel(path, parts[1], parts[0])
			}
		case blackfriday.Heading:
			id := node.HeadingData.HeadingID
			name := string(node.FirstChild.Literal)
			if len(id) > 0 {
				b.logger.Debug("find label: %s(#%s)", name, id)
			} else {
				id = util.GenId(labelIds, name)
				node.HeadingData.HeadingID = id
				b.logger.Debug("generate label: %s(#%s)", name, id)
			}
			label := c.addLabel(path, id, name)

			if node.HeadingData.Level == 1 {
				title = string(node.FirstChild.Literal)
				b.logger.Debug("title: %s", title)
			}

			heading := renderer.NewHeading(label, node.HeadingData.Level)
			toc = append(toc, heading)
		default:
		}
		return blackfriday.GoToNext
	})
	if len(title) == 0 {
		b.logger.Warn("%s does not have title", path)
	}
	c.addResource(renderer.MarkdownFile, path, title, toc)
	return c, nil
}

func (b *builder) Analyze() (Context, error) {
	b.logger.Debug("Analyze: start")
	defer b.logger.Debug("Analyze: End")
	c := newContext()
	conf := filepath.Join(b.sourceDirectory, "conf.lua")
	if err := filepath.Walk(b.sourceDirectory, func(path string, info os.FileInfo, err error) error {
		rel := util.CleanRelPath(b.sourceDirectory, path)
		b.logger.Debug("process: %s", rel)
		if util.IsDir(path) || path == conf || b.isIgnore(rel) {
			b.logger.Debug("ignored.")
			return nil
		}
		name := filepath.Base(path)
		if !strings.HasSuffix(name, ".md") {
			if b.isExtra(rel) {
				b.logger.Debug("found extra file")
				c.addResource(renderer.ExtraFile, path, "", nil)
			} else {
				b.logger.Debug("found static file")
				c.addResource(renderer.StaticFile, path, "", nil)
			}
			return nil
		}
		b.logger.Debug("markdown file")
		_, err = b.AnalyzeMarkdown(c, path)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return c, nil
}

var ErrBuildFailed = errors.New("Build failed")

func (b *builder) Build() error {
	b.logger.Debug("Build: start")
	defer b.logger.Debug("Build: End")
	c, err := b.Analyze()
	if err != nil {
		return err
	}

	b.logger.Info("%d markdown files", c.NumDocuments())
	b.logger.Info("%d static files", c.NumStatics())
	b.logger.Info("%d extra files", c.NumExtras())

	if err := b.buildAux(c); err != nil {
		return err
	}
	if err := b.copyStaticFiles(c); err != nil {
		return err
	}
	if err := b.copyExtraFiles(c); err != nil {
		return err
	}
	return nil
}

func (b *builder) buildAux(c Context) error {
	cpus := runtime.NumCPU()
	cr := make(chan *renderer.HTMLRenderer, cpus)
	ce := make(chan error)
	b.logger.Info("%d concurrent workers", cpus)

	for i := 0; i < cpus; i++ {
		go b.buildWorker(cr, ce)
	}

	count := 0
	for path, ast := range c.(*context).asts {
		rel, _ := filepath.Rel(b.sourceDirectory, path)
		name := filepath.Base(path)
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		b.logger.Debug("render %s", path)
		outdir := filepath.Join(b.outputDirectory, filepath.Dir(rel))
		vars := map[string]interface{}{}
		vars["Site"] = b.config.Site
		vars["Vars"] = b.vars

		r := renderer.NewHTMLRenderer(
			b.logger,
			b.fileSystem,
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			},
			renderer.WithTemplateVars(vars),
			renderer.WithBuildContext(c),
			renderer.WithSourceAST(ast),
			renderer.WithSourceFile(path),
			renderer.WithSourceDirectory(filepath.Dir(path)),
			renderer.WithOutputDirectory(outdir),
			renderer.WithStaticDirectory(filepath.Join(b.outputDirectory, "statics")),
			renderer.WithOutputFormat(renderer.HTML),
			renderer.WithRenderers(b.renderers...),
			renderer.WithTheme(b.theme),
			renderer.WithSiteStorage(map[string]interface{}{}),
		)
		count++
		cr <- r
	}
	results := 0
	es := []error{}
	for {
		select {
		case err := <-ce:
			results++
			if err != nil {
				es = append(es, err)
			}
		}
		if results == count {
			close(cr)
			break
		}
	}
	for _, e := range es {
		b.logger.Error(e.Error())
	}
	if len(es) > 0 {
		return ErrBuildFailed
	}
	return nil
}

func (b *builder) buildWorker(cr chan *renderer.HTMLRenderer, ce chan error) {
	for {
		select {
		case r, ok := <-cr:
			if !ok {
				return
			}
			sourceFile := r.SourceFile()
			err := r.Render()
			if err != nil {
				ce <- errors.Errorf("%s: %s", sourceFile, err.Error())
			} else {
				ce <- nil
			}
		}
	}
}

func (b *builder) copyStaticFiles(c Context) error {
	b.logger.Debug("Copy static files")
	for r := range c.Resources() {
		if r.Type() != renderer.StaticFile {
			continue
		}
		rel := util.CleanRelPath(b.sourceDirectory, r.Path())
		dest := filepath.Join(b.outputDirectory, "statics", filepath.Base(r.Path()))
		b.logger.Debug("copy %s -> %s", rel, dest)
		if err := b.fileSystem.Copy(rel, dest); err != nil {
			return err
		}
	}
	return nil
}

func (b *builder) copyExtraFiles(c Context) error {
	b.logger.Debug("Copy extra files")
	for r := range c.Resources() {
		if r.Type() != renderer.ExtraFile {
			continue
		}
		rel := util.CleanRelPath(b.sourceDirectory, r.Path())
		dest := filepath.Join(b.outputDirectory, rel)
		b.logger.Debug("copy %s -> %s", rel, dest)
		if err := b.fileSystem.Copy(rel, dest); err != nil {
			return err
		}
	}
	return nil
}

func (b *builder) isExtra(rel string) bool {
	if _, ok := b.extras[rel]; ok {
		return true
	}
	return b.matchesExtraPatterns(rel)
}

func (b *builder) matchesExtraPatterns(rel string) bool {
	for _, g := range b.extraPatterns {
		if g.Match(rel) {
			return true
		}
	}
	return false
}

func (b *builder) isIgnore(rel string) bool {
	if _, ok := b.ignores[rel]; ok {
		return true
	}
	return b.matchesIgnorePatterns(rel)
}

func (b *builder) matchesIgnorePatterns(rel string) bool {
	for _, g := range b.ignorePatterns {
		if g.Match(rel) {
			return true
		}
	}
	return false
}

func readConfig(l *lua.LState, v interface{}, name string) error {
	if err := gluamapper.Map(l.GetGlobal(name).(*lua.LTable), v); err != nil {
		return err
	}
	return nil
}

func toString(l *lua.LState) int {
	arg := l.CheckAny(1)
	if ud, ok := arg.(*lua.LUserData); ok {
		if s, ok := ud.Value.([]byte); ok {
			l.Push(lua.LString(string(s)))
		}
	} else {
		l.Push(lua.LString(arg.String()))
	}
	return 1
}

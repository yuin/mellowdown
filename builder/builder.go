package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/theme"
	"github.com/yuin/mellowdown/util"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Ref interface {
	String() string
	DefinedIn() string
	ID() string
	Name() string
}

type ref struct {
	definedIn string
	id        string
	name      string
}

func newRef(definedIn, id, name string) Ref {
	return &ref{
		definedIn: definedIn,
		id:        id,
		name:      name,
	}
}

func (r *ref) DefinedIn() string {
	return r.definedIn
}

func (r *ref) Name() string {
	return r.name
}

func (r *ref) ID() string {
	return r.id
}

func (r *ref) String() string {
	return fmt.Sprintf("Ref{definedIn:%s, id:%s, name:%s}", r.definedIn, r.id, r.name)
}

type CrossRef interface {
	String() string
	From() string
	To() string
}

type crossRef struct {
	from string
	to   string
}

func newCrossRef(from, to string) CrossRef {
	return &crossRef{
		from: from,
		to:   to,
	}
}

func (r *crossRef) From() string {
	return r.from
}

func (r *crossRef) To() string {
	return r.to
}

func (r *crossRef) String() string {
	return fmt.Sprintf("CorssRef{from:%s, to:%s}", r.from, r.to)
}

type Context interface {
	CrossRefs() []CrossRef
	Refs() []Ref
	Get(path string) (*blackfriday.Node, bool)
}

type context struct {
	crossRefs []CrossRef
	refs      []Ref
	asts      map[string]*blackfriday.Node
}

func newContext() *context {
	return &context{
		crossRefs: []CrossRef{},
		refs:      []Ref{},
		asts:      map[string]*blackfriday.Node{},
	}
}

func (c *context) AddAst(name string, node *blackfriday.Node) {
	c.asts[name] = node
}

func (c *context) AddCrossRef(from, to string) {
	c.crossRefs = append(c.crossRefs, newCrossRef(from, to))
}

func (c *context) AddRef(definedIn, id, name string) {
	c.refs = append(c.refs, newRef(definedIn, id, name))
}

func (c *context) CrossRefs() []CrossRef {
	return c.crossRefs
}

func (c *context) Refs() []Ref {
	return c.refs
}

func (c *context) Get(path string) (*blackfriday.Node, bool) {
	v, ok := c.asts[path]
	return v, ok
}

type Builder interface {
	LoadConfig() error
	Build() error
}

type builder struct {
	l               *lua.LState
	fileSystem      asset.FileSystem
	sourceDirectory string
	outputDirectory string
	renderers       []renderer.Renderer
	config          config
	theme           theme.Theme
}

type Option func(*builder)

func SourceDirectory(path string) Option {
	return func(b *builder) {
		b.sourceDirectory = path
	}
}

func OutputDirectory(path string) Option {
	return func(b *builder) {
		b.outputDirectory = path
	}
}

func New(fileSystem asset.FileSystem, opts ...Option) Builder {
	b := &builder{
		fileSystem: fileSystem,
		l:          lua.NewState(),
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

	for _, r := range b.renderers {
		r.InitOption(option)
	}

	return nil
}

func (b *builder) Analyze() (Context, error) {
	c := newContext()
	visited := map[*blackfriday.Node]bool{}
	if err := filepath.Walk(b.sourceDirectory, func(path string, info os.FileInfo, err error) error {
		name := filepath.Base(path)
		if !strings.HasSuffix(name, ".md") {
			return nil
		}
		optList := []blackfriday.Option{blackfriday.WithExtensions(blackfriday.CommonExtensions)}
		parser := blackfriday.New(optList...)
		input, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		ast := parser.Parse(input)
		c.AddAst(path, ast)
		ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			if _, ok := visited[node]; ok {
				return blackfriday.GoToNext
			}
			visited[node] = true
			switch node.Type {
			case blackfriday.Link:
				dest := string(node.LinkData.Destination)
				if strings.HasPrefix(dest, ".") {
					c.AddCrossRef(path, dest)
				}
			case blackfriday.Heading:
				id := node.HeadingData.HeadingID
				if len(id) > 0 {
					name := string(node.FirstChild.Literal)
					c.AddRef(path, id, name)
				}
			default:
			}
			return blackfriday.GoToNext
		})

		return nil
	}); err != nil {
		return nil, err
	}
	return c, nil
}

func (b *builder) Build() error {
	c, err := b.Analyze()
	if err != nil {
		return err
	}
	/*
		for _, r := range c.CrossRefs() {
			fmt.Println(r.String())
		}
		for _, r := range c.Refs() {
			fmt.Println(r.String())
		}
	*/
	if err := filepath.Walk(b.sourceDirectory, func(path string, info os.FileInfo, err error) error {
		name := filepath.Base(path)
		dir := filepath.Dir(path)
		if util.IsDir(path) {
			if strings.HasPrefix(name, "_") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(name, ".md") {
			return nil
		}

		fmt.Printf("%s\n", path)
		rel, _ := filepath.Rel(b.sourceDirectory, dir)
		outdir := filepath.Join(b.outputDirectory, rel)

		ast, _ := c.Get(path)
		r := renderer.NewHTMLRenderer(
			b.fileSystem,
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			},
			renderer.SourceAST(ast),
			renderer.SourceFile(path),
			renderer.SourceDirectory(filepath.Dir(path)),
			renderer.OutputDirectory(outdir),
			renderer.StaticDirectory(filepath.Join(b.outputDirectory, "statics")),
			renderer.OutputFormat(renderer.HTML),
			renderer.Renderers(b.renderers...),
			renderer.Theme(b.theme),
			renderer.SiteStorage(map[string]interface{}{}),
		)
		if err := r.Render(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
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

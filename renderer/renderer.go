package renderer

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	blackfriday "github.com/yuin/blackfriday/v2"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/log"
	"github.com/yuin/mellowdown/theme"
	"github.com/yuin/mellowdown/util"

	yaml "gopkg.in/yaml.v2"
)

type templateVars map[string]interface{}

func (t templateVars) SetDev(v bool) {
	t["Dev"] = v
}

func (t templateVars) SetTitle(v string) {
	t["Title"] = v
}

func (t templateVars) SetLiveReloadScript(v template.HTML) {
	t["LiveReloadScript"] = v
}

func (t templateVars) SetStyleSheets(v template.HTML) {
	t["StyleSheets"] = v
}

func (t templateVars) SetScripts(v template.HTML) {
	t["Scripts"] = v
}

func (t templateVars) SetHeader(v template.HTML) {
	t["Header"] = v
}

func (t templateVars) SetContent(v template.HTML) {
	t["Content"] = v
}

func (t templateVars) SetFooter(v template.HTML) {
	t["Footer"] = v
}

type Format int

const (
	HTML Format = iota + 1
	PDF
)

func FindFormat(name string) (Format, bool) {
	switch strings.ToLower(name) {
	case "html":
		return HTML, true
	case "pdf":
		return PDF, true
	default:
		return Format(0), false
	}
}

type HTMLRenderer struct {
	*blackfriday.HTMLRenderer
	fileSystem      asset.FileSystem
	logger          log.Logger
	dev             bool
	sourceFile      string
	sourceAST       *blackfriday.Node
	sourceDirectory string
	outputFile      string
	outputDirectory string
	outputFormat    Format
	theme           theme.Theme
	wkhtmltopdfPath string
	renderers       []Renderer
	rendererIndex   map[NodeType]map[string]Renderer
	staticDirectory string
	buildContext    BuildContext
	templateVars    map[string]interface{}
	siteStorage     map[string]interface{}
	pageStorage     map[string]interface{}

	headerWriter bytes.Buffer
	footerWriter bytes.Buffer
}

type HTMLRendererOption func(r *HTMLRenderer)

func WithSourceFile(sourceFile string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceFile, _ = filepath.Abs(sourceFile)
	}
}

func WithSourceAST(node *blackfriday.Node) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceAST = node
	}
}

func WithDev() HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.dev = true
	}
}

func WithSourceDirectory(dir string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceDirectory, _ = filepath.Abs(dir)
	}
}

func WithOutputDirectory(outputDirectory string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputDirectory, _ = filepath.Abs(outputDirectory)
	}
}

func WithOutputFormat(outputFormat Format) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputFormat = outputFormat
	}
}

func WithTheme(theme theme.Theme) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.theme = theme
	}
}

func WithStaticDirectory(path string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.staticDirectory, _ = filepath.Abs(path)
	}
}

func WithRenderers(renderers ...Renderer) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.renderers = renderers
		for _, renderer := range renderers {
			atype, aname := renderer.Acceptable()
			_, ok := r.rendererIndex[atype]
			if !ok {
				r.rendererIndex[atype] = map[string]Renderer{}
			}
			r.rendererIndex[atype][aname] = renderer
		}
	}
}

func WithWkhtmltopdfPath(path string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.wkhtmltopdfPath = path
	}
}

func WithBuildContext(b BuildContext) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.buildContext = b
	}
}

func WithTemplateVars(vars map[string]interface{}) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.templateVars = vars
	}
}

func WithSiteStorage(s map[string]interface{}) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.siteStorage = s
	}
}

func NewHTMLRenderer(logger log.Logger, fileSystem asset.FileSystem, params blackfriday.HTMLRendererParameters, options ...HTMLRendererOption) *HTMLRenderer {
	abscwd, _ := filepath.Abs(".")
	ret := &HTMLRenderer{
		HTMLRenderer:    blackfriday.NewHTMLRenderer(params),
		logger:          logger,
		fileSystem:      fileSystem,
		renderers:       []Renderer{},
		rendererIndex:   map[NodeType]map[string]Renderer{},
		sourceDirectory: abscwd,
		outputDirectory: abscwd,
		outputFormat:    HTML,
		theme:           nil,
		wkhtmltopdfPath: "",
		staticDirectory: filepath.Join(abscwd, "statics"),
		siteStorage:     map[string]interface{}{},
		pageStorage:     nil,
	}
	for _, option := range options {
		option(ret)
	}

	ret.outputFile, _ = filepath.Rel(ret.sourceDirectory, ret.sourceFile)
	ret.outputFile = filepath.Join(ret.outputDirectory, ret.outputFile)

	if ret.outputFormat == PDF {
		if len(ret.wkhtmltopdfPath) == 0 {
			ret.wkhtmltopdfPath = os.Getenv("WKHTMLTOPDF_PATH")
		}
		if len(ret.wkhtmltopdfPath) == 0 {
			ret.wkhtmltopdfPath = "wkhtmltopdf"
		}
		ret.outputFile = util.ReplaceExtension(ret.outputFile, ".md", ".pdf")
	} else {
		ret.outputFile = util.ReplaceExtension(ret.outputFile, ".md", ".html")
	}

	for _, renderer := range ret.renderers {
		renderer.NewDocument(ret)
	}
	for _, r := range ret.renderers {
		r.RenderHeader(ret.HeaderWriter(), ret)
	}
	for _, r := range ret.renderers {
		r.RenderFooter(ret.FooterWriter(), ret)
	}
	return ret
}

var funcRegex *regexp.Regexp

func init() {
	funcRegex = regexp.MustCompile(`(?:(?:.*[^\\])|\A)!(\w[\w0-9_-]+)!`)
}

func (r *HTMLRenderer) resolveReference(node *blackfriday.Node) {
	dest := node.LinkData.Destination
	if len(dest) != 0 {
		d := string(dest)
		if !util.IsUrl(d) {
			if strings.HasPrefix(d, "#") {
				id := strings.TrimLeft(d, "#")
				label, ok := r.FindLabel(id)
				if !ok {
					r.logger.Error("%s : label %s is not defined.", r.SourceFile(), id)
					return
				}
				relpath, _ := filepath.Rel(filepath.Dir(r.SourceFile()), label.DefinedIn())
				node.LinkData.Destination = []byte(util.SlashPath(util.ReplaceExtension(relpath, ".md", ".html") + d))
				text := string(node.FirstChild.Literal)
				if text == " " {
					node.FirstChild.Literal = []byte(label.Name())
				}
				return
			}

			if !strings.HasPrefix(d, ".") {
				d = "./" + d
			}
			parts := strings.Split(d, "#")
			if strings.HasSuffix(parts[0], ".md") {
				parts[0] = util.ReplaceExtension(parts[0], ".md", ".html")
				d = strings.Join(parts, "#")
				node.LinkData.Destination = []byte(d)
				return
			}
			dpath := filepath.Clean(filepath.Join(filepath.Dir(r.SourceFile()), d))
			relpath, _ := filepath.Rel(filepath.Dir(r.SourceFile()), dpath)
			staticFile := filepath.Join(r.StaticDirectory(), relpath)
			node.LinkData.Destination = []byte(r.StaticPath(staticFile))
		}
	}
}

func (r *HTMLRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	r.resolveReference(node)
	n, ok := newNode(node)
	if ok {
		renderer, found := r.findRenderer(n.Type(), n.Identifier())
		if found {
			if err := renderer.Render(w, n, r); err != nil {
				r.logger.Error("%s: %s", r.sourceFile, err.Error())
			}
			return blackfriday.GoToNext
		}
	}
	return r.HTMLRenderer.RenderNode(w, node, entering)
}

func (r *HTMLRenderer) ConvertString(source string) []byte {
	result := r.ConvertBytes([]byte(source))
	return result
}

func (r *HTMLRenderer) ConvertBytes(bs []byte) []byte {
	optList := []blackfriday.Option{blackfriday.WithRenderer(r), blackfriday.WithExtensions(blackfriday.CommonExtensions | blackfriday.Functions)}
	parser := blackfriday.New(optList...)
	return r.ConvertAST(parser.Parse(bs))
}

func (r *HTMLRenderer) ConvertAST(ast *blackfriday.Node) []byte {
	var buf bytes.Buffer
	r.RenderHeader(&buf, ast)
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		return r.RenderNode(&buf, node, entering)
	})
	r.RenderFooter(&buf, ast)
	return buf.Bytes()
}

func (r *HTMLRenderer) HeaderHTML() string {
	return r.headerWriter.String()
}

func (r *HTMLRenderer) HeaderWriter() io.Writer {
	return &r.headerWriter
}

func (r *HTMLRenderer) FooterHTML() string {
	return r.footerWriter.String()
}

func (r *HTMLRenderer) FooterWriter() io.Writer {
	return &r.footerWriter
}

func (r *HTMLRenderer) SourceFile() string {
	return r.sourceFile
}

func (r *HTMLRenderer) OutputFile() string {
	return r.outputFile
}

func (r *HTMLRenderer) OutputDirectory() string {
	return r.outputDirectory
}

func (r *HTMLRenderer) StaticPath(file string) string {
	return util.SlashPath(util.CleanRelPath(filepath.Dir(r.OutputFile()), file))
}

func (r *HTMLRenderer) StaticDirectory() string {
	path := r.staticDirectory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return path
}

func (r *HTMLRenderer) Converter() MarkdownConverter {
	return r
}

func (r *HTMLRenderer) SiteStorage() map[string]interface{} {
	return r.siteStorage
}

func (r *HTMLRenderer) PageStorage() map[string]interface{} {
	return r.pageStorage
}

func (r *HTMLRenderer) copyThemeFiles(files []os.FileInfo, filter func(os.FileInfo) bool, cb func(os.FileInfo, string)) error {
	themeFilesDir := r.theme.FilesDirectory()
	for _, f := range files {
		if !filter(f) {
			continue
		}
		dest := filepath.Join(r.StaticDirectory(), f.Name())
		if err := r.fileSystem.Copy(filepath.Join(themeFilesDir, f.Name()), dest); err != nil {
			return err
		}
		cb(f, dest)
	}
	return nil
}

func (r *HTMLRenderer) FindLabel(id string) (Label, bool) {
	if r.buildContext == nil {
		return nil, false
	}
	return r.buildContext.FindLabel(id)
}

func (r *HTMLRenderer) FindResource(path string) (Resource, bool) {
	if r.buildContext == nil {
		return nil, false
	}
	return r.buildContext.FindResource(path)
}

func (r *HTMLRenderer) NumDocuments() int {
	if r.buildContext == nil {
		return 1
	}
	return r.buildContext.NumDocuments()
}

func (r *HTMLRenderer) NumStatics() int {
	if r.buildContext == nil {
		return 0
	}
	return r.buildContext.NumStatics()
}

func (r *HTMLRenderer) NumExtras() int {
	if r.buildContext == nil {
		return 0
	}
	return r.buildContext.NumExtras()
}

func (r *HTMLRenderer) Render() error {
	optOutputDirectory := r.outputDirectory
	r.pageStorage = map[string]interface{}{}
	if r.outputFormat == PDF {
		tmpOutputDirectory, err := ioutil.TempDir("", "mellowdown-")
		if err != nil {
			return err
		}
		r.outputDirectory = tmpOutputDirectory
		if r.staticDirectory == optOutputDirectory {
			r.staticDirectory = tmpOutputDirectory
		}
		defer os.RemoveAll(tmpOutputDirectory)
	}

	filename := filepath.Base(r.sourceFile)
	var result []byte
	if r.sourceAST != nil {
		result = r.ConvertAST(r.sourceAST)
	} else {
		bs, err := ioutil.ReadFile(r.sourceFile)
		if err != nil {
			return err
		}
		result = r.ConvertBytes(bs)
	}

	tv := templateVars{}
	for k, v := range r.templateVars {
		tv[k] = v
	}
	tv.SetTitle(util.GetTitle(result))
	tv.SetContent(template.HTML(result))
	themeFiles, err := r.theme.Files()
	if err != nil {
		return err
	}
	var cssBuf bytes.Buffer
	if err := r.copyThemeFiles(themeFiles, func(f os.FileInfo) bool {
		return strings.HasSuffix(f.Name(), ".css")
	}, func(f os.FileInfo, dest string) {
		fmt.Fprintf(&cssBuf, `<link rel="stylesheet" type="text/css" href="%s" media="screen">`,
			r.StaticPath(dest))
	}); err != nil {
		return err
	}
	tv.SetStyleSheets(template.HTML(cssBuf.String()))

	var scriptBuf bytes.Buffer
	if err := r.copyThemeFiles(themeFiles, func(f os.FileInfo) bool {
		return strings.HasSuffix(f.Name(), ".js")
	}, func(f os.FileInfo, dest string) {
		fmt.Fprintf(&scriptBuf, `<script src="%s">`,
			r.StaticPath(dest))
	}); err != nil {
		return err
	}
	tv.SetScripts(template.HTML(scriptBuf.String()))

	if err := r.copyThemeFiles(themeFiles, func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), ".css") && !strings.HasSuffix(f.Name(), ".js")
	}, func(f os.FileInfo, dest string) {
	}); err != nil {
		return err
	}

	tv.SetLiveReloadScript(template.HTML(`<script src="http://localhost:35730/livereload.js"></script>`))
	tv.SetDev(r.dev)
	tv.SetHeader(template.HTML(r.HeaderHTML()))
	tv.SetFooter(template.HTML(r.FooterHTML()))

	htmlfile := filepath.Join(r.outputDirectory, util.ReplaceExtension(filename, ".md", ".html"))
	if err := util.EnsureDirectoryExists(r.outputDirectory); err != nil {
		return err
	}
	mainTemplate, err := r.theme.MainTemplate()
	if err != nil {
		return err
	}
	var out bytes.Buffer
	if err := mainTemplate.Execute(&out, &tv); err != nil {
		return err
	}
	if err := util.WriteFile(htmlfile, out.Bytes()); err != nil {
		return err
	}

	if r.outputFormat == PDF {
		pdffile := filepath.Join(optOutputDirectory, util.ReplaceExtension(filename, ".md", ".pdf"))
		wkhtmltopdf.SetPath(r.wkhtmltopdfPath)
		pdfg, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			return err
		}
		page := wkhtmltopdf.NewPage(htmlfile)
		page.FooterRight.Set("[page]")
		page.FooterFontSize.Set(10)
		page.PrintMediaType.Set(true)
		page.NoStopSlowScripts.Set(true)
		page.DisableJavascript.Set(false)
		page.JavascriptDelay.Set(10000)
		pdfg.AddPage(page)

		if err := pdfg.Create(); err != nil {
			return err
		}

		if err := pdfg.WriteFile(pdffile); err != nil {
			return err
		}
	}
	return nil
}

func (r *HTMLRenderer) findRenderer(n NodeType, name string) (Renderer, bool) {
	v1, ok := r.rendererIndex[n]
	if !ok {
		return nil, false
	}
	v2, ok := v1[name]
	if ok {
		return v2, true
	}
	v3, ok := v1[Any]
	return v3, ok
}

type OptionType int

const (
	Cli OptionType = iota + 1
	Lua
)

type Option interface {
	Type() OptionType
	Flag() *flag.FlagSet
	Lua() *lua.LState
}

type option struct {
	typ  OptionType
	flag *flag.FlagSet
	l    *lua.LState
}

func NewCliOption(fs *flag.FlagSet) Option {
	return &option{
		typ:  Cli,
		flag: fs,
		l:    nil,
	}
}

func NewLuaOption(l *lua.LState) Option {
	return &option{
		typ:  Lua,
		flag: nil,
		l:    l,
	}
}

func (o *option) Type() OptionType {
	return o.typ
}

func (o *option) Flag() *flag.FlagSet {
	return o.flag
}

func (o *option) Lua() *lua.LState {
	return o.l
}

type Renderer interface {
	Name() string
	AddOption(Option)
	InitOption(Option)
	NewDocument(context RenderingContext)
	Acceptable() (NodeType, string)
	Render(w io.Writer, node Node, context RenderingContext) error
	RenderHeader(w io.Writer, context RenderingContext) error
	RenderFooter(w io.Writer, context RenderingContext) error
}

type MarkdownConverter interface {
	ConvertString(string) []byte
	ConvertBytes([]byte) []byte
}

func readMeta(data []byte) (map[interface{}]interface{}, []byte, error) {
	s := string(data)
	buf := []string{}
	rest := []string{}
	flag := false
	for _, line := range strings.Split(s, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			flag = true
			continue
		}
		if flag {
			rest = append(rest, line)
		} else {
			buf = append(buf, line)
		}
	}
	if len(buf) > 0 {
		yamlText := strings.Join(buf, "\n")
		var meta interface{}
		err := yaml.Unmarshal([]byte(yamlText), &meta)
		if err != nil {
			return nil, nil, err
		}
		return meta.(map[interface{}]interface{}), []byte(strings.Join(rest, "\n")), nil
	} else {
		return map[interface{}]interface{}{}, data, nil
	}
}

func readMetaStruct(data []byte, dest interface{}) ([]byte, error) {
	s := string(data)
	buf := []string{}
	rest := []string{}
	flag := false
	for _, line := range strings.Split(s, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			flag = true
			continue
		}
		if flag {
			rest = append(rest, line)
		} else {
			buf = append(buf, line)
		}
	}
	if len(buf) > 0 {
		yamlText := strings.Join(buf, "\n")
		if err := yaml.Unmarshal([]byte(yamlText), dest); err != nil {
			return nil, err
		}
		return []byte(strings.Join(rest, "\n")), nil
	} else {
		return data, nil
	}
}

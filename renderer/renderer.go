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
	"github.com/yuin/gopher-lua"
	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/theme"
	"github.com/yuin/mellowdown/util"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type templateVars struct {
	Debug            bool
	Title            string
	LiveReloadScript template.HTML
	StyleSheets      template.HTML
	Scripts          template.HTML
	Header           template.HTML
	Content          template.HTML
	Footer           template.HTML
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
	debug           bool
	sourceFile      string
	sourceAST       *blackfriday.Node
	sourceDirectory string
	outputFile      string
	outputDirectory string
	outputFormat    Format
	theme           theme.Theme
	wkhtmltopdfPath string
	renderers       []Renderer
	staticDirectory string
	siteStorage     map[string]interface{}
	pageStorage     map[string]interface{}

	headerWriter bytes.Buffer
	footerWriter bytes.Buffer
	function     string
}

type HTMLRendererOption func(r *HTMLRenderer)

func SourceFile(sourceFile string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceFile, _ = filepath.Abs(sourceFile)
	}
}

func SourceAST(node *blackfriday.Node) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceAST = node
	}
}

func Debug() HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.debug = true
	}
}

func SourceDirectory(dir string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.sourceDirectory, _ = filepath.Abs(dir)
	}
}

func OutputDirectory(outputDirectory string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputDirectory, _ = filepath.Abs(outputDirectory)
	}
}

func OutputFormat(outputFormat Format) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputFormat = outputFormat
	}
}

func Theme(theme theme.Theme) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.theme = theme
	}
}

func StaticDirectory(path string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.staticDirectory, _ = filepath.Abs(path)
	}
}

func Renderers(renderers ...Renderer) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.renderers = renderers
	}
}

func WkhtmltopdfPath(path string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.wkhtmltopdfPath = path
	}
}

func SiteStorage(s map[string]interface{}) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.siteStorage = s
	}
}

func NewHTMLRenderer(fileSystem asset.FileSystem, params blackfriday.HTMLRendererParameters, options ...HTMLRendererOption) *HTMLRenderer {
	abscwd, _ := filepath.Abs(".")
	ret := &HTMLRenderer{
		HTMLRenderer:    blackfriday.NewHTMLRenderer(params),
		fileSystem:      fileSystem,
		renderers:       []Renderer{},
		sourceDirectory: abscwd,
		outputDirectory: abscwd,
		outputFormat:    HTML,
		theme:           nil,
		wkhtmltopdfPath: "",
		staticDirectory: filepath.Join(abscwd, "statics"),
		siteStorage:     map[string]interface{}{},
		pageStorage:     nil,
		function:        "",
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

	for _, r := range ret.renderers {
		r.NewDocument()
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

func (r *HTMLRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	var n Node
	f := ""
	ok := false
	if len(r.function) != 0 {
		f = r.function
		r.function = ""
		if node.Type == blackfriday.Code {
			n = newFunctionNode(node, f)
			ok = true
		} else {
			fmt.Fprintf(w, "!%s!", f)
		}
	}
	if node.Type == blackfriday.Text {
		result := funcRegex.FindAllSubmatchIndex(node.Literal, -1)
		if len(result) > 0 && result[0][len(result[0])-1] == len(node.Literal)-1 {
			fmt.Fprint(w, string(node.Literal[:result[0][len(result[0])-2]-1]))
			r.function = string(node.Literal[result[0][len(result[0])-2]:result[0][len(result[0])-1]])
			return blackfriday.GoToNext
		}
	}

	if !ok {
		n, ok = newNode(node)
	}
	if ok {
		for _, fr := range r.renderers {
			if fr.Accept(n) {
				if err := fr.Render(w, n, r); err != nil {
					fmt.Fprintf(w, "ERROR:%s", err.Error())
				}
				return blackfriday.GoToNext
			}
		}
	}

	if ok && n.Type() == NodeFunction {
		fmt.Fprintf(w, "!%s!", f)
	}
	return r.HTMLRenderer.RenderNode(w, node, entering)
}

func (r *HTMLRenderer) ConvertString(source string) []byte {
	result := r.ConvertBytes([]byte(source))
	return result
}

func (r *HTMLRenderer) ConvertBytes(bs []byte) []byte {
	parser := blackfriday.New(blackfriday.WithRenderer(r))
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
	relpath, _ := filepath.Rel(filepath.Dir(r.OutputFile()), file)
	return strings.Replace(relpath, "\\", "/", -1)
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

	for _, r := range r.renderers {
		r.NewDocument()
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

	tv := &templateVars{}
	tv.Title = util.GetTitle(result)
	tv.Content = template.HTML(result)
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
	tv.StyleSheets = template.HTML(cssBuf.String())

	var scriptBuf bytes.Buffer
	if err := r.copyThemeFiles(themeFiles, func(f os.FileInfo) bool {
		return strings.HasSuffix(f.Name(), ".js")
	}, func(f os.FileInfo, dest string) {
		fmt.Fprintf(&scriptBuf, `<script src="%s">`,
			r.StaticPath(dest))
	}); err != nil {
		return err
	}
	tv.Scripts = template.HTML(scriptBuf.String())

	if err := r.copyThemeFiles(themeFiles, func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), ".css") && !strings.HasSuffix(f.Name(), ".js")
	}, func(f os.FileInfo, dest string) {
	}); err != nil {
		return err
	}

	tv.LiveReloadScript = template.HTML(`<script src="http://localhost:35730/livereload.js"></script>`)
	tv.Debug = r.debug
	tv.Header = template.HTML(r.HeaderHTML())
	tv.Footer = template.HTML(r.FooterHTML())

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
	NewDocument()
	Accept(node Node) bool
	Render(w io.Writer, node Node, context RenderingContext) error
	RenderHeader(w io.Writer, context RenderingContext) error
	RenderFooter(w io.Writer, context RenderingContext) error
}

type MarkdownConverter interface {
	ConvertString(string) []byte
	ConvertBytes([]byte) []byte
}

type RenderingContext interface {
	Converter() MarkdownConverter
	SourceFile() string
	OutputFile() string
	OutputDirectory() string
	StaticDirectory() string
	StaticPath(file string) string
	SiteStorage() map[string]interface{}
	PageStorage() map[string]interface{}
}

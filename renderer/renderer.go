package renderer

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/pkg/errors"
	"github.com/yuin/mellowdown/style"
	"github.com/yuin/mellowdown/util"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	template = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>%s</title>
%s
<script src="http://localhost:35730/livereload.js"></script>
</head>
<body>
<article class="markdown-body">
%s
</article>
%s
</body>
</html>
`
)

type Format int

const (
	HTML = iota + 1
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
	inputFile       string
	outputDirectory string
	outputFormat    Format
	style           string
	wkhtmltopdfPath string
	renderers       []Renderer
	staticDirectory string

	headerWriter bytes.Buffer
	footerWriter bytes.Buffer
	function     string
}

type HTMLRendererOption func(r *HTMLRenderer)

func InputFile(inputFile string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.inputFile = inputFile
	}
}

func OutputDirectory(outputDirectory string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputDirectory = outputDirectory
	}
}

func OutputFormat(outputFormat Format) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.outputFormat = outputFormat
	}
}

func Style(style string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.style = style
	}
}

func StaticDirectory(path string) HTMLRendererOption {
	return func(r *HTMLRenderer) {
		r.staticDirectory = path
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

func NewHTMLRenderer(params blackfriday.HTMLRendererParameters, options ...HTMLRendererOption) *HTMLRenderer {
	ret := &HTMLRenderer{
		HTMLRenderer:    blackfriday.NewHTMLRenderer(params),
		renderers:       []Renderer{},
		outputDirectory: ".",
		outputFormat:    HTML,
		style:           "github",
		wkhtmltopdfPath: "",
		staticDirectory: ".",
		function:        "",
	}
	for _, option := range options {
		option(ret)
	}

	if ret.outputFormat == PDF {
		if len(ret.wkhtmltopdfPath) == 0 {
			ret.wkhtmltopdfPath = os.Getenv("WKHTMLTOPDF_PATH")
		}
		if len(ret.wkhtmltopdfPath) == 0 {
			ret.wkhtmltopdfPath = "wkhtmltopdf"
		}
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

func (r *HTMLRenderer) InputFile() string {
	return r.inputFile
}

func (r *HTMLRenderer) OutputDirectory() string {
	return r.outputDirectory
}

func (r *HTMLRenderer) ImageDirectory() (string, error) {
	return r.StaticDirectory("images")
}

func (r *HTMLRenderer) StaticPath(file string) string {
	relpath, _ := filepath.Rel(r.OutputDirectory(), file)
	return relpath
}

func (r *HTMLRenderer) StaticDirectory(name string) (string, error) {
	path := filepath.Join(r.staticDirectory, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", errors.WithStack(err)
		}
	}
	return path, nil
}

func (r *HTMLRenderer) Render() error {
	optOutputDirectory := r.outputDirectory
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

	bs, err := ioutil.ReadFile(r.inputFile)
	if err != nil {
		return err
	}
	filename := filepath.Base(r.inputFile)
	fmt.Fprintf(r.HeaderWriter(), "\n<style>\n%s\n</style>\n", style.Get(r.style))
	result := blackfriday.Run(bs, blackfriday.WithRenderer(r))
	htmlfile := filepath.Join(r.outputDirectory, util.ReplaceExtension(filename, ".md", ".html"))
	if err := util.EnsureDirectoryExists(r.outputDirectory); err != nil {
		return err
	}
	htmldata := fmt.Sprintf(template, util.GetTitle(result), r.HeaderHTML(), result, r.FooterHTML())

	if err := ioutil.WriteFile(htmlfile, []byte(htmldata), os.ModePerm); err != nil {
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

type RendererOption func(r Renderer)

type Renderer interface {
	Name() string
	AddOption(*flag.FlagSet)
	InitOption()
	NewDocument()
	Accept(node Node) bool
	Render(w io.Writer, node Node, context RenderingContext) error
	RenderHeader(w io.Writer, context RenderingContext) error
	RenderFooter(w io.Writer, context RenderingContext) error
}

type RenderingContext interface {
	InputFile() string
	OutputDirectory() string
	ImageDirectory() (string, error)
	StaticPath(file string) string
}

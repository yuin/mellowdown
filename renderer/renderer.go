package renderer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type HTMLRenderer struct {
	*blackfriday.HTMLRenderer
	renderers    []Renderer
	headerWriter bytes.Buffer
	footerWriter bytes.Buffer
	function     string
}

func NewHTMLRenderer(params blackfriday.HTMLRendererParameters, rs []Renderer) *HTMLRenderer {
	ret := &HTMLRenderer{
		HTMLRenderer: blackfriday.NewHTMLRenderer(params),
		renderers:    rs,
		function:     "",
	}
	for _, r := range rs {
		r.NewDocument()
	}
	for _, r := range rs {
		r.RenderHeader(ret.HeaderWriter())
	}
	for _, r := range rs {
		r.RenderFooter(ret.FooterWriter())
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
				if err := fr.Render(w, n); err != nil {
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

type Renderer interface {
	Name() string
	SetOutputDirectory(path string)
	SetFile(path string)
	AddOption()
	InitOption()
	NewDocument()
	Accept(node Node) bool
	Render(w io.Writer, node Node) error
	RenderHeader(w io.Writer) error
	RenderFooter(w io.Writer) error
}

func ImageDirectory(outputDir string) (string, error) {
	path := filepath.Join(outputDir, "images")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", errors.WithStack(err)
		}
	}
	return path, nil
}

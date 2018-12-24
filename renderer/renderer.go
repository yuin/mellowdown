package renderer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type HTMLRenderer struct {
	*blackfriday.HTMLRenderer
	renderers    []Renderer
	headerWriter bytes.Buffer
	footerWriter bytes.Buffer
}

func NewHTMLRenderer(params blackfriday.HTMLRendererParameters, rs []Renderer) *HTMLRenderer {
	ret := &HTMLRenderer{
		HTMLRenderer: blackfriday.NewHTMLRenderer(params),
		renderers:    rs,
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

func (r *HTMLRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	n, ok := newNode(node)
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

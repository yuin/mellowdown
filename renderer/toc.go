package renderer

import (
	"fmt"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
	"io"
	"strings"
)

type TOCOption struct {
	Depth int    `yaml:"depth"`
	Title string `yaml:"title"`
}

type TOCRenderer struct {
}

func NewTOCRenderer() Renderer {
	return &TOCRenderer{}
}

func (r *TOCRenderer) Name() string {
	return "toc"
}

func (r *TOCRenderer) AddOption(o Option) {
}

func (r *TOCRenderer) InitOption(o Option) {
}

func (r *TOCRenderer) NewDocument(c RenderingContext) {
}

func (r *TOCRenderer) Acceptable() (NodeType, string) {
	return NodeFencedCode, "toc"
}

func (r *TOCRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	return nil
}

func (r *TOCRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	var tocOpt TOCOption
	err := yaml.Unmarshal(node.Text(), &tocOpt)
	if err != nil {
		return errors.WithStack(err)
	}

	resource, _ := c.FindResource(c.SourceFile())
	headings := resource.TOC()
	current := 0
	depth := 1
	fmt.Fprint(w, `<div class="toc-container">`+"\n")
	if len(tocOpt.Title) > 0 {
		fmt.Fprintf(w, `<p class="toc-title">%s</p>`+"\n", tocOpt.Title)
	}
	fmt.Fprint(w, `<ol class="toc">`+"\n")
	for _, heading := range headings {
		if current != 0 {
			if heading.Level() < current {
				depth--
				fmt.Fprintf(w, "%s</ol></li>\n", strings.Repeat(" ", current))
			} else if heading.Level() != current {
				depth++
				if tocOpt.Depth != 0 && depth > tocOpt.Depth {
					continue
				}
				fmt.Fprintf(w, "%s<li><ol>\n", strings.Repeat(" ", current))
			}
		}
		current = heading.Level()
		fmt.Fprintf(w, `%s<li><a href="#%s">%s</a></li>`+"\n", strings.Repeat(" ", current), heading.ID(), heading.Name())
	}
	if current != 1 {
		fmt.Fprintf(w, "%s</ol></li>\n", strings.Repeat(" ", current))
	}
	fmt.Fprint(w, "</ol>\n")
	fmt.Fprint(w, "</div>\n")
	return nil
}

func (r *TOCRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

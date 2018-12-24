package renderer

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	chromastyles "github.com/alecthomas/chroma/styles"
)

type SyntaxHighlightingRenderer struct {
	Style        *chroma.Style
	optHighlight string
}

func NewSyntaxHighlightingRenderer() Renderer {
	return &SyntaxHighlightingRenderer{}
}

func (r *SyntaxHighlightingRenderer) Name() string {
	return "syntax-highlight"
}

func (r *SyntaxHighlightingRenderer) SetOutputDirectory(path string) {
}

func (r *SyntaxHighlightingRenderer) SetFile(path string) {
}

func (r *SyntaxHighlightingRenderer) AddOption() {
	flag.StringVar(&r.optHighlight, "syntax-highlight", "monokailight", fmt.Sprintf("Syntax Highlightinging Style (Optional, available styles:%s)", strings.Join(chromastyles.Names(), ",")))
}

func (r *SyntaxHighlightingRenderer) InitOption() {
	r.Style = chromastyles.Get(r.optHighlight)
}

func (r *SyntaxHighlightingRenderer) NewDocument() {
}

func (r *SyntaxHighlightingRenderer) Accept(n Node) bool {
	return n.Type() == FencedCode && lexers.Get(n.FencedCodeBlock().Info()) != nil
}

func (r *SyntaxHighlightingRenderer) RenderHeader(w io.Writer) error {
	formatter := html.New(html.WithClasses())
	fmt.Fprint(w, "\n<style>\n")
	formatter.WriteCSS(w, r.Style)
	fmt.Fprint(w, "\n</style>\n")
	return nil
}

func (r *SyntaxHighlightingRenderer) Render(w io.Writer, node Node) error {
	lexer := lexers.Get(node.FencedCodeBlock().Info())
	iterator, err := lexer.Tokenise(nil, fmt.Sprintf("%s", node.Text()))
	if err != nil {
		return err
	}
	formatter := html.New(html.WithClasses())
	var buf bytes.Buffer
	if err := formatter.Format(&buf, r.Style, iterator); err != nil {
		return err
	}
	s := strings.Replace(buf.String(), "<pre class=\"chroma\">", "<pre class=\"chroma\"><code>", -1)
	s = strings.Replace(s, "</pre>", "</code></pre>", -1)
	fmt.Fprint(w, s)
	return nil
}

func (r *SyntaxHighlightingRenderer) RenderFooter(w io.Writer) error {
	return nil
}

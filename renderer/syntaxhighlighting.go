package renderer

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	chromastyles "github.com/alecthomas/chroma/styles"
	lua "github.com/yuin/gopher-lua"
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

func (r *SyntaxHighlightingRenderer) AddOption(o Option) {
	switch o.Type() {
	case Cli:
		o.Flag().StringVar(&r.optHighlight, "syntax-highlight", "monokailight", fmt.Sprintf("Syntax Highlightinging Style (Optional, available styles:%s)", strings.Join(chromastyles.Names(), ",")))
	case Lua:
		o.Lua().DoString(`
	      if theme == nil then
		    theme = {}
		  end
		  theme["syntax_highlight"] = "monokailight"
		`)
	}
}

func (r *SyntaxHighlightingRenderer) InitOption(o Option) {
	if o.Type() == Lua {
		r.optHighlight = string(o.Lua().GetGlobal("theme").(*lua.LTable).RawGetString("syntax_highlight").(lua.LString))
	}

	r.Style = chromastyles.Get(r.optHighlight)
}

func (r *SyntaxHighlightingRenderer) NewDocument() {
}

func (r *SyntaxHighlightingRenderer) Accept(n Node) bool {
	return n.Type() == NodeFencedCode && lexers.Get(n.FencedCodeBlock().Info()) != nil
}

func (r *SyntaxHighlightingRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	formatter := html.New(html.WithClasses())
	fmt.Fprint(w, "\n<style>\n")
	formatter.WriteCSS(w, r.Style)
	fmt.Fprint(w, "\n</style>\n")
	return nil
}

func (r *SyntaxHighlightingRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
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

func (r *SyntaxHighlightingRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

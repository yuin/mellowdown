package renderer

import (
	"fmt"
	"io"
	"strings"
)

type LabelRenderer struct {
}

func NewLabelRenderer() Renderer {
	return &LabelRenderer{}
}

func (r *LabelRenderer) Name() string {
	return "label"
}

func (r *LabelRenderer) AddOption(o Option) {
}

func (r *LabelRenderer) InitOption(o Option) {
}

func (r *LabelRenderer) NewDocument(c RenderingContext) {
}

func (r *LabelRenderer) Acceptable() (NodeType, string) {
	return NodeFunction, "label"
}

func (r *LabelRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	return nil
}

func (r *LabelRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	text := node.Function().Arguments()[0].(string)
	if !strings.Contains(text, "#") {
		return fmt.Errorf("invalid label: %s", text)
	}

	parts := strings.Split(text, "#")
	fmt.Fprintf(w, `<a name="#%s">%s</a>`, parts[1], parts[0])
	return nil
}

func (r *LabelRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

package renderer

import (
	"io"
)

type NullRenderer struct {
}

func NewNullRenderer() Renderer {
	return &NullRenderer{}
}

func (r *NullRenderer) Name() string {
	return ""
}

func (r *NullRenderer) AddOption(o Option) {
}

func (r *NullRenderer) InitOption(o Option) {
}

func (r *NullRenderer) NewDocument() {
}

func (r *NullRenderer) Accept(n Node) bool {
	return false
}

func (r *NullRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	return nil
}

func (r *NullRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	return nil
}

func (r *NullRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

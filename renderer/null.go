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

func (r *NullRenderer) SetOutputDirectory(path string) {
}

func (r *NullRenderer) SetFile(path string) {
}

func (r *NullRenderer) AddOption() {
}

func (r *NullRenderer) InitOption() {
}

func (r *NullRenderer) NewDocument() {
}

func (r *NullRenderer) Accept(n Node) bool {
	return false
}

func (r *NullRenderer) RenderHeader(w io.Writer) error {
	return nil
}

func (r *NullRenderer) Render(w io.Writer, node Node) error {
	return nil
}

func (r *NullRenderer) RenderFooter(w io.Writer) error {
	return nil
}

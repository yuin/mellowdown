// +build !windows

package renderer

func NewPPTRenderer() Renderer {
	return &NullRenderer{}
}

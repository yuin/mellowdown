// +build !windows

package renderer

func NewPPTRenderer(file, outdir string) FenceCodeRenderer {
	return &NullRenderer{}
}

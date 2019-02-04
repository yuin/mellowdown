// +build windows

package renderer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const ppShapeFormatGIF = 0
const ppShapeFormatJPG = 1
const ppShapeFormatPNG = 2
const ppShapeFormatBMP = 3
const ppScaleXY = 4

type PPTRenderer struct {
}

type PPT struct {
	File      string `yaml:"file"`
	ShapeName string `yaml:"shape"`
	Width     int    `yaml:"width"`
	Height    int    `yaml:"height"`
}

func NewPPTRenderer() Renderer {
	return &PPTRenderer{}
}

func (r *PPTRenderer) Name() string {
	return "ppt"
}

func (r *PPTRenderer) AddOption(o Option) {
}

func (r *PPTRenderer) InitOption(o Option) {
}

func (r *PPTRenderer) NewDocument() {
}

func (r *PPTRenderer) Accept(n Node) bool {
	return n.Type() == NodeFencedCode && n.FencedCodeBlock().Info() == "ppt"
}

func (r *PPTRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	return nil
}

func (r *PPTRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	var ppt PPT
	err := yaml.Unmarshal(node.Text(), &ppt)
	if err != nil {
		return errors.WithStack(err)
	}
	dir := c.StaticDirectory()
	outpath := filepath.Join(dir, ppt.ShapeName+".png")
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		file := ppt.File
		if !filepath.IsAbs(file) {
			file = filepath.Clean(filepath.Join(filepath.Dir(c.SourceFile()), file))
		}

		if err := r.ppt2png(file, ppt.ShapeName, outpath, ppt.Width, ppt.Height); err != nil {
			return err
		}
	}
	relpath, _ := filepath.Rel(c.OutputDirectory(), outpath)
	fmt.Fprintf(w, "<img src=\"%s\" style=\"display:block\" />", relpath)
	return nil
}

func (r *PPTRenderer) ppt2png(filename, shapename, outpath string, w, h int) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return errors.Errorf("%s does not exists", filename)
	}

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("PowerPoint.Application")
	if err != nil {
		return errors.WithStack(err)
	}

	app, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return errors.WithStack(err)
	}

	oleutil.PutProperty(app, "DisplayAlerts", 1)
	ps, _ := oleutil.GetProperty(app, "Presentations")
	presentations, err := oleutil.CallMethod(ps.ToIDispatch(), "Open", filename, -1, 0, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	slides, _ := oleutil.GetProperty(presentations.ToIDispatch(), "Slides")
	if w != 0 || h != 0 {
		sm, _ := oleutil.GetProperty(presentations.ToIDispatch(), "SlideMaster")
		ws, _ := oleutil.GetProperty(sm.ToIDispatch(), "Width")
		hs, _ := oleutil.GetProperty(sm.ToIDispatch(), "Height")
		if w == 0 {
			w = int(ws.Value().(float32))
		}
		if h == 0 {
			h = int(hs.Value().(float32))
		}
	}

	if err := oleutil.ForEach(slides.ToIDispatch(), func(slide *ole.VARIANT) error {
		shapes, _ := oleutil.GetProperty(slide.ToIDispatch(), "Shapes")
		if err := oleutil.ForEach(shapes.ToIDispatch(), func(shape *ole.VARIANT) error {
			titlev, _ := oleutil.GetProperty(shape.ToIDispatch(), "Title")
			title := titlev.ToString()
			if title == shapename {
				_, err := oleutil.CallMethod(shape.ToIDispatch(), "Export", outpath, ppShapeFormatPNG, w, h, ppScaleXY)
				if err != nil {
					return errors.WithStack(err)
				}
			}
			return nil
		}); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}); err != nil {
		return err
	}

	_, err = oleutil.CallMethod(app, "Quit")
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *PPTRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

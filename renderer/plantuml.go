package renderer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
)

type PlantUMLRenderer struct {
	optPlantUMLPath string
}

func NewPlantUMLRenderer() Renderer {
	return &PlantUMLRenderer{
		optPlantUMLPath: "plantuml",
	}
}

func (r *PlantUMLRenderer) Name() string {
	return "platnuml"
}

func (r *PlantUMLRenderer) AddOption(o Option) {
	switch o.Type() {
	case Cli:
		o.Flag().StringVar(&r.optPlantUMLPath, "plantuml", "", "PlantUML executable file path(Optional). If this value is empty, PLANTUML_PATH envvar value will be used as an executable file path")
	case Lua:
		o.Lua().DoString(`plantuml = {
			path = ""
		}`)
	}
}

func (r *PlantUMLRenderer) InitOption(o Option) {
	if o.Type() == Lua {
		r.optPlantUMLPath = string(o.Lua().GetGlobal("plantuml").(*lua.LTable).RawGetString("path").(lua.LString))
	}

	if len(r.optPlantUMLPath) == 0 {
		r.optPlantUMLPath = os.Getenv("PLANTUML_PATH")
	}
	if len(r.optPlantUMLPath) == 0 {
		r.optPlantUMLPath = "plantuml"
	}
}

func (r *PlantUMLRenderer) NewDocument() {
}

func (r *PlantUMLRenderer) Accept(n Node) bool {
	return n.Type() == NodeFencedCode && n.FencedCodeBlock().Info() == "uml"
}

func (r *PlantUMLRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	return nil
}

func (r *PlantUMLRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	h := sha256.New()
	h.Write(node.Text())
	dir := c.StaticDirectory()
	temp := filepath.Join(dir, fmt.Sprintf("%x", h.Sum(nil)))
	if _, err := os.Stat(temp + ".png"); os.IsNotExist(err) {
		if _, err := os.Stat(r.optPlantUMLPath); os.IsNotExist(err) {
			return errors.Errorf("PlantUML launcher '%s' not found on path", r.optPlantUMLPath)
		}

		if err := ioutil.WriteFile(temp, node.Text(), os.ModePerm); err != nil {
			return errors.WithStack(err)
		}
		defer os.Remove(temp)
		if err := exec.Command(r.optPlantUMLPath, temp).Run(); err != nil {
			return errors.WithStack(err)
		}
	}
	relpath := c.StaticPath(temp + ".png")
	fmt.Fprintf(w, "<img src=\"%s\" style=\"display:block\" />", relpath)
	return nil
}

func (r *PlantUMLRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	return nil
}

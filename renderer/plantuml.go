package renderer

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

type PlantUMLRenderer struct {
	OutputDirectory string
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

func (r *PlantUMLRenderer) SetOutputDirectory(path string) {
	r.OutputDirectory = path
}

func (r *PlantUMLRenderer) SetFile(path string) {
}

func (r *PlantUMLRenderer) AddOption() {
	flag.StringVar(&r.optPlantUMLPath, "plantuml", "plantuml", "PlantUML Path(Optional). If PLANTUML_PATH envvar is not empty, this option will be overwritten by its value.")
}

func (r *PlantUMLRenderer) InitOption() {
	if len(os.Getenv("PLANTUML_PATH")) != 0 {
		r.optPlantUMLPath = os.Getenv("PLANTUML_PATH")
	}
}

func (r *PlantUMLRenderer) NewDocument() {
}

func (r *PlantUMLRenderer) Accept(n Node) bool {
	return n.Type() == FencedCode && n.FencedCodeBlock().Info() == "uml"
}

func (r *PlantUMLRenderer) RenderHeader(w io.Writer) error {
	return nil
}

func (r *PlantUMLRenderer) Render(w io.Writer, node Node) error {
	h := sha256.New()
	h.Write(node.Text())
	dir, err := ImageDirectory(r.OutputDirectory)
	if err != nil {
		return err
	}
	temp := filepath.Join(dir, fmt.Sprintf("%x", h.Sum(nil)))
	if _, err := os.Stat(temp + ".png"); os.IsNotExist(err) {
		defer os.Remove(temp)
		if err := ioutil.WriteFile(temp, node.Text(), os.ModePerm); err != nil {
			return errors.WithStack(err)
		}
		if err := exec.Command(r.optPlantUMLPath, temp).Run(); err != nil {
			return errors.WithStack(err)
		}
	}
	relpath, _ := filepath.Rel(r.OutputDirectory, temp+".png")
	fmt.Fprintf(w, "<img src=\"%s\" style=\"display:block\" />", relpath)
	return nil
}

func (r *PlantUMLRenderer) RenderFooter(w io.Writer) error {
	return nil
}
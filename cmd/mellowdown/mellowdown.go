package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/mattn/goemon"

	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/style"
)

const (
	goemonconf = `
tasks:
- match: '**/*.md'
  commands:
    - |
      %s -file ${GOEMON_TARGET_FILE}
    - :livereload
`
)

func abort(err interface{}, status int) {
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	os.Exit(status)
}

type renderOptions struct {
	OutputDirectory string
	File            string
	Format          string
	Style           string
	Addr            string
	WkhtmltopdfPath string
	LuaScripts      string
}

const usage string = `Usage of mellowdown: mellowdown [-h] ACTION [ACTION_ARGS]

"mellowdown ACTION -h" shows a usage fo the action.

ACTIONS:
    render: Render a single markdown file`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println(usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "render":
		opts := renderOptions{}
		fs := flag.NewFlagSet("mellowdown "+os.Args[1], flag.ExitOnError)
		fs.StringVar(&opts.OutputDirectory, "out", "", "Output Directory(Optional)")
		fs.StringVar(&opts.File, "file", "", "Markdown file(Required)")
		fs.StringVar(&opts.Format, "format", "html", "Output format(html or pdf)")
		fs.StringVar(&opts.Style, "style", "github", fmt.Sprintf("Style (Optional, available styles:%s)", strings.Join(style.Names(), ",")))
		fs.StringVar(&opts.Addr, "addr", "", "address like localhost:8000, this enables livereloading")
		fs.StringVar(&opts.WkhtmltopdfPath, "wkhtmltopdf", "", "Wkhtmltopdf executable file path(Optional). If this value is empty, WKHTMLTOPDF_PATH envvar value will be used as an executable file path")
		fs.StringVar(&opts.LuaScripts, "lua", "", "comma separeted lua renderers")
		rs := []renderer.Renderer{
			renderer.NewPlantUMLRenderer(),
			renderer.NewPPTRenderer(),
			renderer.NewSyntaxHighlightingRenderer(),
		}

		for _, r := range rs {
			r.AddOption(fs)
		}
		fs.Parse(os.Args[2:])
		format, ok := renderer.FindFormat(opts.Format)

		if len(opts.File) == 0 && len(opts.Addr) == 0 || !ok {
			fs.PrintDefaults()
			os.Exit(1)
		}
		syntaxr := rs[len(rs)-1]
		rs = rs[:len(rs)-1]

		if len(opts.LuaScripts) > 0 {
			for _, script := range strings.Split(opts.LuaScripts, ",") {
				r, err := renderer.NewLuaRenderer(script)
				if err != nil {
					abort(err, 1)
				}
				rs = append(rs, r)
			}
		}
		rs = append(rs, syntaxr)

		if !filepath.IsAbs(opts.File) {
			opts.File, _ = filepath.Abs(opts.File)
		}
		if len(opts.OutputDirectory) == 0 {
			opts.OutputDirectory = filepath.Dir(opts.File)
		}

		for _, r := range rs {
			r.InitOption()
		}

		if len(opts.Addr) != 0 {
			cmdbuf := []string{}
			skip := false
			for _, arg := range os.Args {
				if arg == "-addr" {
					skip = true
					continue
				}
				if skip {
					skip = false
					continue
				}
				cmdbuf = append(cmdbuf, EscapeArg(arg))
			}
			tempfile, err := ioutil.TempFile("", "mellowdown")
			if err != nil {
				abort(err, 1)
			}
			if _, err := tempfile.Write([]byte(fmt.Sprintf(goemonconf, strings.Join(cmdbuf, " ")))); err != nil {
				abort(err, 1)
			}
			defer os.Remove(tempfile.Name())
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				for range c {
					os.Remove(tempfile.Name())
					close(c)
					os.Exit(0)
				}
			}()

			g := goemon.NewWithArgs([]string{})
			g.File = tempfile.Name()
			g.Logger = log.New(os.Stdout, "[mellowdown]", log.LstdFlags)
			g.Run()
			http.Handle("/", http.FileServer(http.Dir(".")))
			if err := http.ListenAndServe(opts.Addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				g.Logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
				http.DefaultServeMux.ServeHTTP(w, r)
			})); err != nil {
				abort(err, 1)
			}
		} else {
			r := renderer.NewHTMLRenderer(
				blackfriday.HTMLRendererParameters{
					Flags: blackfriday.CommonHTMLFlags,
				},
				renderer.InputFile(opts.File),
				renderer.Renderers(rs...),
				renderer.OutputDirectory(opts.OutputDirectory),
				renderer.Style(opts.Style),
				renderer.OutputFormat(format),
				renderer.WkhtmltopdfPath(opts.WkhtmltopdfPath),
				renderer.StaticDirectory(opts.OutputDirectory),
			)
			if err := r.Render(); err != nil {
				abort(err, 1)
			}
		}
	default:
		fmt.Println(usage)
		os.Exit(1)
	}

}

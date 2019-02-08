package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/mattn/goemon"
	blackfriday "github.com/yuin/blackfriday/v2"
	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/builder"
	"github.com/yuin/mellowdown/log"
	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/theme"
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

type renderOptions struct {
	Debug           bool
	OutputDirectory string
	File            string
	Format          string
	Theme           string
	Addr            string
	WkhtmltopdfPath string
	LuaScripts      string
}

func render() {
	opts := renderOptions{}
	fileSystem := asset.NewFileSystem()
	themes := theme.NewThemes(fileSystem)
	themes.AddLoadPath(".")
	if err := themes.Load(); err != nil {
		abort(err, 1)
	}

	fs := flag.NewFlagSet("mellowdown "+os.Args[1], flag.ExitOnError)
	fs.BoolVar(&opts.Debug, "debug", false, "Debug mode")
	fs.StringVar(&opts.OutputDirectory, "out", "", "Output Directory(Optional)")
	fs.StringVar(&opts.File, "file", "", "Markdown file(Required)")
	fs.StringVar(&opts.Format, "format", "html", "Output format(html or pdf)")
	fs.StringVar(&opts.Theme, "theme", "github", fmt.Sprintf("Theme (Optional, available themes:%s)", strings.Join(themes.Names(), ",")))
	fs.StringVar(&opts.Addr, "addr", "", "address like localhost:8000, this enables livereloading")
	fs.StringVar(&opts.WkhtmltopdfPath, "wkhtmltopdf", "", "Wkhtmltopdf executable file path(Optional). If this value is empty, WKHTMLTOPDF_PATH envvar value will be used as an executable file path")
	fs.StringVar(&opts.LuaScripts, "lua", "", "comma separeted lua renderers")
	rs := []renderer.Renderer{
		renderer.NewPlantUMLRenderer(),
		renderer.NewPPTRenderer(),
		renderer.NewTOCRenderer(),
		renderer.NewLabelRenderer(),
		renderer.NewSyntaxHighlightingRenderer(),
	}
	option := renderer.NewCliOption(fs)

	for _, r := range rs {
		r.AddOption(option)
	}
	fs.Parse(os.Args[2:])
	var logger log.Logger
	if opts.Debug {
		logger = log.NewLogger(log.Debug)
	} else {
		logger = log.NewLogger(log.Info)
	}

	format, ok := renderer.FindFormat(opts.Format)

	if len(opts.File) == 0 && len(opts.Addr) == 0 || !ok {
		fs.PrintDefaults()
		os.Exit(1)
	}
	themev, ok := themes.Get(opts.Theme)
	if !ok {
		fmt.Fprintf(os.Stderr, "theme %s not found", opts.Theme)
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
		r.InitOption(option)
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
		g.Logger = stdlog.New(logger, "[mellowdown]", stdlog.LstdFlags)
		g.Run()
		http.Handle("/", http.FileServer(http.Dir(".")))
		if err := http.ListenAndServe(opts.Addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
			http.DefaultServeMux.ServeHTTP(w, r)
		})); err != nil {
			abort(err, 1)
		}
	} else {
		c, err := builder.New(logger, fileSystem).AnalyzeMarkdown(nil, opts.File)
		if err != nil {
			abort(err, 1)
		}
		ast, _ := c.FindAST(opts.File)
		r := renderer.NewHTMLRenderer(
			logger,
			fileSystem,
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			},
			renderer.WithSourceFile(opts.File),
			renderer.WithSourceAST(ast),
			renderer.WithSourceDirectory(filepath.Dir(opts.File)),
			renderer.WithOutputDirectory(opts.OutputDirectory),
			renderer.WithStaticDirectory(filepath.Join(opts.OutputDirectory, "statics")),
			renderer.WithOutputFormat(format),
			renderer.WithRenderers(rs...),
			renderer.WithTheme(themev),
			renderer.WithWkhtmltopdfPath(opts.WkhtmltopdfPath),
			renderer.WithSiteStorage(map[string]interface{}{}),
			renderer.WithBuildContext(c),
		)
		if err := r.Render(); err != nil {
			abort(err, 1)
		}
	}
}

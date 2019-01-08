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
	"regexp"
	"strings"

	"github.com/mattn/goemon"
	"github.com/pkg/errors"

	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/style"
)

const (
	template = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>%s</title>
%s
<script src="http://localhost:35730/livereload.js"></script>
</head>
<body>
<article class="markdown-body">
%s
</article>
%s
</body>
</html>
`
	goemonconf = `
tasks:
- match: '*.md'
  commands:
    - |
      %s -file ${GOEMON_TARGET_FILE}
    - :livereload
`
)

func ensureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func getTitle(html []byte) string {
	for _, r := range []*regexp.Regexp{
		regexp.MustCompile(`<h1>([^<]+)</h1>`),
		regexp.MustCompile(`<h2>([^<]+)</h2>`)} {
		result := r.FindAllSubmatch(html, -1)
		if len(result) > 0 {
			return string(result[0][1])
		}
	}
	return ""
}

func abort(err interface{}, status int) {
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	os.Exit(status)
}

func main() {
	var (
		optOutputDirectory string
		optFile            string
		optFormat          string
		optStyle           string
		optAddr            string
		optWkhtmltopdfPath string
		optLuaScripts      string
	)
	flag.StringVar(&optOutputDirectory, "out", "", "Output Directory(Optional)")
	flag.StringVar(&optFile, "file", "", "Markdown file(Required)")
	flag.StringVar(&optFormat, "format", "html", "Output format(html or pdf)")
	flag.StringVar(&optStyle, "style", "github", fmt.Sprintf("Style (Optional, available styles:%s)", strings.Join(style.Names(), ",")))
	flag.StringVar(&optAddr, "addr", "", "address like localhost:8000, this enables livereloading")
	flag.StringVar(&optWkhtmltopdfPath, "wkhtmltopdf", "", "Wkhtmltopdf executable file path(Optional). If this value is empty, WKHTMLTOPDF_PATH envvar value will be used as an executable file path")
	flag.StringVar(&optLuaScripts, "lua", "", "comma separeted lua renderers")
	rs := []renderer.Renderer{
		renderer.NewPlantUMLRenderer(),
		renderer.NewPPTRenderer(),
		renderer.NewSyntaxHighlightingRenderer(),
	}

	for _, r := range rs {
		r.AddOption()
	}
	flag.Parse()
	if len(optFile) == 0 && len(optAddr) == 0 || (optFormat != "html" && optFormat != "pdf") {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	syntaxr := rs[len(rs)-1]
	rs = rs[:len(rs)-1]

	if len(optLuaScripts) > 0 {
		for _, script := range strings.Split(optLuaScripts, ",") {
			r, err := renderer.NewLuaRenderer(script)
			if err != nil {
				abort(err, 1)
			}
			rs = append(rs, r)
		}
	}
	rs = append(rs, syntaxr)

	if !filepath.IsAbs(optFile) {
		optFile, _ = filepath.Abs(optFile)
	}
	if len(optOutputDirectory) == 0 {
		optOutputDirectory = filepath.Dir(optFile)
	}

	filename := filepath.Base(optFile)
	pdffile := filepath.Join(optOutputDirectory, filename[0:len(filename)-len(".md")]+".pdf")
	if optFormat == "pdf" {
		var err error
		optOutputDirectory, err = ioutil.TempDir("", "mellowdown-")
		if err != nil {
			abort(err, 1)
		}
		defer os.RemoveAll(optOutputDirectory)
	}

	for _, r := range rs {
		r.SetFile(optFile)
		r.SetOutputDirectory(optOutputDirectory)
		r.InitOption()
	}

	if len(optAddr) != 0 {
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
		http.ListenAndServe(optAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
			http.DefaultServeMux.ServeHTTP(w, r)
		}))
	} else {
		r := renderer.NewHTMLRenderer(
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			},
			rs,
		)
		for _, r := range rs {
			r.NewDocument()
		}

		bs, err := ioutil.ReadFile(optFile)
		if err != nil {
			abort(err, 1)
		}
		defer func() {
			if rcv := recover(); rcv != nil {
				if err, ok := rcv.(error); ok {
					abort(err, 1)
				} else {
					abort(rcv, 1)
				}
			}
		}()
		fmt.Fprintf(r.HeaderWriter(), "\n<style>\n%s\n</style>\n", style.Get(optStyle))
		result := blackfriday.Run(bs, blackfriday.WithRenderer(r))
		htmlfile := filepath.Join(optOutputDirectory, filename[0:len(filename)-len(".md")]+".html")
		if err := ensureDirectoryExists(optOutputDirectory); err != nil {
			abort(err, 1)
		}
		htmldata := fmt.Sprintf(template, getTitle(result), r.HeaderHTML(), result, r.FooterHTML())

		if err := ioutil.WriteFile(htmlfile, []byte(htmldata), os.ModePerm); err != nil {
			abort(err, 1)
		}

		if optFormat == "pdf" {
			if len(optWkhtmltopdfPath) == 0 {
				optWkhtmltopdfPath = os.Getenv("WKHTMLTOPDF_PATH")
			}
			if len(optWkhtmltopdfPath) == 0 {
				optWkhtmltopdfPath = "wkhtmltopdf"
			}
			wkhtmltopdf.SetPath(optWkhtmltopdfPath)
			pdfg, err := wkhtmltopdf.NewPDFGenerator()
			if err != nil {
				abort(err, 1)
			}
			page := wkhtmltopdf.NewPage(htmlfile)
			page.FooterRight.Set("[page]")
			page.FooterFontSize.Set(10)
			pdfg.AddPage(page)

			if err := pdfg.Create(); err != nil {
				abort(err, 1)
			}

			if err := pdfg.WriteFile(pdffile); err != nil {
				abort(err, 1)
			}
		}
	}
}

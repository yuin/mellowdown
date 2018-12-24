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
	"syscall"

	"github.com/mattn/goemon"
	"github.com/pkg/errors"

	blackfriday "gopkg.in/russross/blackfriday.v2"

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

func main() {
	var (
		optOutputDirectory string
		optFile            string
		optStyle           string
		optAddr            string
	)
	flag.StringVar(&optOutputDirectory, "out", "", "Output Directory(Optional)")
	flag.StringVar(&optFile, "file", "", "Markdown file(Required)")
	flag.StringVar(&optStyle, "style", "github", fmt.Sprintf("Style (Optional, available styles:%s)", strings.Join(style.Names(), ",")))
	flag.StringVar(&optAddr, "addr", "", "address like localhost:8000, this enables livereloading")
	rs := []renderer.Renderer{
		renderer.NewPlantUMLRenderer(),
		renderer.NewPPTRenderer(),
		renderer.NewLuaRenderer(),
		renderer.NewSyntaxHighlightingRenderer(),
	}
	for _, r := range rs {
		r.AddOption()
	}
	flag.Parse()
	if len(optFile) == 0 && len(optAddr) == 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !filepath.IsAbs(optFile) {
		optFile, _ = filepath.Abs(optFile)
	}
	if len(optOutputDirectory) == 0 {
		optOutputDirectory = filepath.Dir(optFile)
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
			cmdbuf = append(cmdbuf, syscall.EscapeArg(arg))
		}
		tempfile, err := ioutil.TempFile("", "mellowdown")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", errors.WithStack(err))
			os.Exit(1)
		}
		if _, err := tempfile.Write([]byte(fmt.Sprintf(goemonconf, strings.Join(cmdbuf, " ")))); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", errors.WithStack(err))
			os.Exit(1)
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
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			os.Exit(1)
		}
		defer func() {
			if rcv := recover(); rcv != nil {
				if err, ok := rcv.(error); ok {
					fmt.Fprintf(os.Stderr, "%+v\n", errors.WithStack(err))
					os.Exit(1)
				} else {
					fmt.Fprintf(os.Stderr, "%v\n", rcv)
					os.Exit(1)
				}
			}
		}()
		fmt.Fprintf(r.HeaderWriter(), "\n<style>\n%s\n</style>\n", style.Get(optStyle))
		result := blackfriday.Run(bs, blackfriday.WithRenderer(r))
		filename := filepath.Base(optFile)
		htmlfile := filename[0:len(filename)-len(".md")] + ".html"
		if err := ensureDirectoryExists(optOutputDirectory); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", errors.WithStack(err))
			os.Exit(1)
		}
		htmldata := fmt.Sprintf(template, getTitle(result), r.HeaderHTML(), result, r.FooterHTML())

		if err := ioutil.WriteFile(filepath.Join(optOutputDirectory, htmlfile), []byte(htmldata), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", errors.WithStack(err))
			os.Exit(1)
		}
	}
}

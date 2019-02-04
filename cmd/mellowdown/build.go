package main

import (
	"flag"
	"os"

	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/builder"
)

type buildOptions struct {
	OutputDirectory string
	SourceDirectory string
}

func build() {
	opts := buildOptions{}
	fs := flag.NewFlagSet("mellowdown "+os.Args[1], flag.ExitOnError)
	fs.StringVar(&opts.SourceDirectory, "src", "", "Source directory(Required)")
	fs.StringVar(&opts.OutputDirectory, "out", "", "Output directory(Required)")
	fs.Parse(os.Args[2:])
	if len(opts.SourceDirectory) == 0 || len(opts.OutputDirectory) == 0 {
		fs.PrintDefaults()
		os.Exit(1)
	}
	fileSystem := asset.NewFileSystem()
	fileSystem.SetRoot(opts.SourceDirectory)
	builder := builder.New(fileSystem,
		builder.SourceDirectory(opts.SourceDirectory),
		builder.OutputDirectory(opts.OutputDirectory))
	if err := builder.LoadConfig(); err != nil {
		abort(err, 1)
	}
	if err := builder.Build(); err != nil {
		abort(err, 1)
	}
}

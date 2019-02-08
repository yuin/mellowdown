package main

import (
	"flag"
	"os"

	"github.com/yuin/mellowdown/asset"
	"github.com/yuin/mellowdown/builder"
	"github.com/yuin/mellowdown/log"
)

type buildOptions struct {
	Debug           bool
	OutputDirectory string
	SourceDirectory string
}

func build() {
	opts := buildOptions{}
	fs := flag.NewFlagSet("mellowdown "+os.Args[1], flag.ExitOnError)
	fs.BoolVar(&opts.Debug, "debug", false, "Debug mode")
	fs.StringVar(&opts.SourceDirectory, "src", "", "Source directory(Required)")
	fs.StringVar(&opts.OutputDirectory, "out", "", "Output directory(Required)")
	fs.Parse(os.Args[2:])
	var logger log.Logger
	if opts.Debug {
		logger = log.NewLogger(log.Debug)
	} else {
		logger = log.NewLogger(log.Info)
	}

	if len(opts.SourceDirectory) == 0 || len(opts.OutputDirectory) == 0 {
		fs.PrintDefaults()
		os.Exit(1)
	}
	fileSystem := asset.NewFileSystem()
	fileSystem.SetRoot(opts.SourceDirectory)
	builder := builder.New(logger, fileSystem,
		builder.WithSourceDirectory(opts.SourceDirectory),
		builder.WithOutputDirectory(opts.OutputDirectory))
	if err := builder.LoadConfig(); err != nil {
		abort(err, 1)
	}
	if err := builder.Build(); err != nil {
		abort(err, 1)
	}
}

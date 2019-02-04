package main

import (
	"fmt"
	"os"
)

func abort(err interface{}, status int) {
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	os.Exit(status)
}

const usage string = `Usage of mellowdown: mellowdown [-h] ACTION [ACTION_ARGS]

"mellowdown ACTION -h" shows a usage fo the action.

ACTIONS:
    render: Render a single markdown file
    build:  Build documents`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println(usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "render":
		render()
		os.Exit(0)
	case "build":
		build()
		os.Exit(0)
	default:
		fmt.Println(usage)
		os.Exit(1)
	}

}

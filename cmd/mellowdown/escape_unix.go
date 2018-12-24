// +build !windows

package main

import "github.com/alessio/shellescape"

func EscapeArg(v string) string {
	return shellescape.Quote(v)
}

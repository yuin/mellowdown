// +build windows

package main

import "syscall"

func EscapeArg(v string) string {
	return syscall.EscapeArg(v)
}

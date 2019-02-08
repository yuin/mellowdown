package util

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var nonAlphaNumeric = regexp.MustCompile("[^a-zA-Z0-9]+")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SlashPath(path string) string {
	return strings.Replace(path, "\\", "/", -2)
}

func CleanRelPath(base, target string) string {
	rel, _ := filepath.Rel(base, target)
	return SlashPath(filepath.Clean(rel))
}

func IsUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func GenId(exists map[string]bool, name string) string {
	anName := nonAlphaNumeric.ReplaceAllString(name, "")
	if len(anName) == 0 {
		anName = "h"
	}
	if _, ok := exists[anName]; ok {
		for i := 0; ; i++ {
			newName := fmt.Sprintf("%s%d", anName, i)
			if _, ok := exists[newName]; !ok {
				anName = newName
				break
			}
		}
	}
	exists[anName] = true
	return anName
}

func GetTitle(html []byte) string {
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

func ReplaceExtension(path, from, to string) string {
	return path[0:len(path)-len(from)] + to
}

type FileType int

const (
	FtFile FileType = iota
	FtDir
	FtLink
	FtNotExists
	FtOther
)

func PathType(path string) FileType {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return FtNotExists
		}
		return FtOther
	}

	if (fi.Mode() & os.ModeSymlink) == os.ModeSymlink {
		return FtLink
	}
	if fi.IsDir() {
		return FtDir
	}
	return FtFile
}

func IsDir(path string) bool { return PathType(path) == FtDir }

func IsFile(path string) bool { return PathType(path) == FtFile }

func PathExists(path string) bool { return PathType(path) != FtNotExists }

func EnsureDirectoryExists(path string) error {
	dir := filepath.Dir(path)
	if !PathExists(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(path string, data []byte) error {
	if err := EnsureDirectoryExists(path); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, data, 0755); err != nil {
		return err
	}
	return nil
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

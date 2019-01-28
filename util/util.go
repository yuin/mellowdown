package util

import (
	"os"
	"regexp"

	"github.com/pkg/errors"
)

func EnsureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func ReplaceExtension(path, from, to string) string {
	return path[0:len(path)-len(from)] + to
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

package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

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

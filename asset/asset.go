//go:generate vfsgendev -source="github.com/yuin/mellowdown/asset".Assets
package asset

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	//"github.com/pkg/errors"
	"github.com/yuin/mellowdown/util"
)

type FileSystem interface {
	SetRoot(root ...string)
	ListDir(path string) ([]os.FileInfo, error)
	Open(path string) (http.File, error)
	Walk(path string, cb func(path string, file os.FileInfo, err error) error) error
	Copy(source, dest string) error
	CopyFile(source, dest string) error
	// source = src/dir, dest = dest/hoge
	// copy src/dir/{a,b,c} -> dst/hoge/{a,b,c}
	CopyTree(source string, dest string) error
}

type fileSystem struct {
	root []string
}

func NewFileSystem() FileSystem {
	return &fileSystem{
		root: []string{},
	}
}

func (f *fileSystem) SetRoot(root ...string) {
	f.root = root
}

func (f *fileSystem) ListDir(path string) ([]os.FileInfo, error) {
	return ListDir(path, f.root...)
}

func (f *fileSystem) Open(path string) (http.File, error) {
	return Open(path, f.root...)
}

func (f *fileSystem) walk(path string, info os.FileInfo, cb func(path string, file os.FileInfo, err error) error) error {
	if !info.IsDir() {
		return f.walk(path, info, nil)
	}

	infos, err := f.ListDir(path)
	err1 := cb(path, info, err)
	if err != nil || err1 != nil {
		return err1
	}

	for _, fileinfo := range infos {
		filename := filepath.Join(path, fileinfo.Name())
		err = f.walk(filename, fileinfo, cb)
		if err != nil {
			if !fileinfo.IsDir() || err != filepath.SkipDir {
				return err
			}
		}
	}

	return nil
}

func (f *fileSystem) Walk(root string, cb func(path string, file os.FileInfo, err error) error) error {
	file, err := f.Open(root)
	if file != nil {
		file.Close()
	}
	if err != nil {
		err = cb(root, nil, err)
	}
	if err != nil {
		err = cb(root, nil, err)
	} else {
		stat, err := file.Stat()
		if err != nil {
			err = cb(root, nil, err)
		} else {
			err = f.walk(root, stat, cb)
		}
	}

	if err == filepath.SkipDir {
		return nil
	}

	return err
}

func (f *fileSystem) Copy(source, dest string) error {
	file, err := f.Open(source)
	defer file.Close()
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if !info.IsDir() {
		err = f.CopyFile(source, dest)
	} else {
		err = f.CopyTree(source, dest)
	}
	return err
}

func (f *fileSystem) CopyFile(source, dest string) error {
	if err := util.EnsureDirectoryExists(dest); err != nil {
		return err
	}
	file, err := f.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, bs, 0644)
}

func (f *fileSystem) CopyTree(source, dest string) error {
	if !util.PathExists(dest) {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			return err
		}
	}

	lst, err := f.ListDir(source)
	if err != nil {
		return err
	}

	for _, item := range lst {
		srcfile := filepath.Join(source, item.Name())
		destfile := filepath.Join(dest, item.Name())
		err := f.Copy(srcfile, destfile)
		if err != nil {
			return err
		}
	}
	return nil
}

func ListDir(path string, root ...string) ([]os.FileInfo, error) {
	results := []os.FileInfo{}
	for _, r := range root {
		d := filepath.Join(r, path)
		if util.IsDir(d) {
			fp, err := os.Open(d)
			if err != nil {
				return nil, err
			}
			lst, err := fp.Readdir(-1)
			if err != nil {
				return nil, err
			}
			results = append(results, lst...)
		}
	}

	f, err := rawOpen(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if os.IsNotExist(err) {
		return []os.FileInfo{}, nil
	}
	lst, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	results = append(results, lst...)
	return results, nil
}

func Open(path string, root ...string) (http.File, error) {
	for _, r := range root {
		f := filepath.Join(r, path)
		if util.IsFile(f) {
			fp, err := os.Open(f)
			if err == nil {
				return fp, nil
			}
		}
	}
	return rawOpen(path)
}

func rawOpen(path string) (http.File, error) {
	// govfsgen always uses UNIX style paths
	path = strings.Replace(path, "\\", "/", -1)
	return Assets.Open(path)
}

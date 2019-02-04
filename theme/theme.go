package theme

import (
	"github.com/yuin/mellowdown/asset"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	themesDirRoot = "_themes"
)

type Theme interface {
	MainTemplate() (*template.Template, error)
	Files() ([]os.FileInfo, error)
	FilesDirectory() string
}

type theme struct {
	root         string
	fileSystem   asset.FileSystem
	mainTemplate *template.Template
	once         sync.Once
}

func newTheme(fs asset.FileSystem, root string) Theme {
	return &theme{
		root:       root,
		fileSystem: fs,
	}
}

func (t *theme) MainTemplate() (*template.Template, error) {
	var err error
	t.once.Do(func() {
		path := filepath.Join(t.root, "templates", "main.html")
		var file http.File
		file, err = t.fileSystem.Open(path)
		if err == nil {
			var bs []byte
			bs, err = ioutil.ReadAll(file)
			if err == nil {
				t.mainTemplate, err = template.New("main").Parse(string(bs))
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return t.mainTemplate, err
}

func (t *theme) FilesDirectory() string {
	return filepath.Join(t.root, "files")
}

func (t *theme) Files() ([]os.FileInfo, error) {
	lst, err := t.fileSystem.ListDir(filepath.Join(t.root, "files"))
	if os.IsNotExist(err) {
		return []os.FileInfo{}, nil
	}
	return lst, err
}

type Themes interface {
	Load() error
	SetFileSystem(asset.FileSystem)
	AddLoadPath(path string)
	Get(name string) (Theme, bool)
	Names() []string
}

type themes struct {
	values     map[string]Theme
	loadPaths  []string
	fileSystem asset.FileSystem
}

func NewThemes(fileSystem asset.FileSystem) Themes {
	return &themes{
		values:     map[string]Theme{},
		loadPaths:  []string{},
		fileSystem: fileSystem,
	}
}

func (t *themes) Load() error {
	lst, err := t.fileSystem.ListDir(themesDirRoot)
	if err != nil {
		return err
	}
	for _, file := range lst {
		if !file.IsDir() {
			continue
		}
		t.values[file.Name()] = newTheme(t.fileSystem, filepath.Join(themesDirRoot, file.Name()))
	}
	return nil
}

func (t *themes) SetFileSystem(fs asset.FileSystem) {
	t.fileSystem = fs
}

func (t *themes) AddLoadPath(path string) {
	t.loadPaths = append(t.loadPaths, path)
}

func (t *themes) Get(name string) (Theme, bool) {
	v, ok := t.values[name]
	return v, ok
}

func (t *themes) Names() []string {
	result := []string{}
	for k := range t.values {
		result = append(result, k)
	}
	return result
}

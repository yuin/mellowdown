package renderer

import (
	"fmt"
	"strings"
)

type TOC []Heading

type BuildContext interface {
	FindLabel(id string) (Label, bool)
	FindResource(path string) (Resource, bool)
	NumDocuments() int
	NumStatics() int
	NumExtras() int
}

type RenderingContext interface {
	BuildContext
	Converter() MarkdownConverter
	SourceFile() string
	OutputFile() string
	OutputDirectory() string
	StaticDirectory() string
	StaticPath(file string) string
	SiteStorage() map[string]interface{}
	PageStorage() map[string]interface{}
}

type Label interface {
	String() string
	DefinedIn() string
	ID() string
	Name() string
}

type label struct {
	definedIn string
	id        string
	name      string
}

func NewLabel(definedIn, id, name string) Label {
	return &label{
		definedIn: definedIn,
		id:        id,
		name:      name,
	}
}

func (r *label) DefinedIn() string {
	return r.definedIn
}

func (r *label) Name() string {
	return r.name
}

func (r *label) ID() string {
	return r.id
}

func (r *label) String() string {
	return fmt.Sprintf("Label{definedIn:%s, id:%s, name:%s}", r.definedIn, r.id, r.name)
}

type Heading interface {
	Label
	Level() int
}

type heading struct {
	Label
	level int
}

func NewHeading(l Label, level int) Heading {
	return &heading{
		Label: l,
		level: level,
	}
}

func (h *heading) Level() int {
	return h.level
}

type ResourceType int

const (
	MarkdownFile ResourceType = iota
	StaticFile
	ExtraFile
)

func (r ResourceType) String() string {
	switch r {
	case MarkdownFile:
		return "Markdown"
	case StaticFile:
		return "Static"
	case ExtraFile:
		return "Extra"
	}
	return ""
}

type Resource interface {
	Type() ResourceType
	Path() string
	Title() string
	String() string
	TOC() TOC
}

type resource struct {
	typ   ResourceType
	path  string
	title string
	toc   TOC
}

func NewResource(typ ResourceType, path, title string, toc TOC) Resource {
	return &resource{
		typ:   typ,
		path:  strings.Replace(path, "\\", "/", -1),
		title: title,
		toc:   toc,
	}
}

func (r *resource) Type() ResourceType {
	return r.typ
}

func (r *resource) Path() string {
	return r.path
}

func (r *resource) Title() string {
	return r.title
}

func (r *resource) TOC() TOC {
	return r.toc
}

func (r *resource) String() string {
	return fmt.Sprintf("Resource{type:%s, path:%s, title:%s}", r.typ.String(), r.path, r.title)
}

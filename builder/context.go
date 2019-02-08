package builder

import (
	"github.com/yuin/mellowdown/renderer"
	"github.com/yuin/mellowdown/util"
	"runtime"

	blackfriday "github.com/yuin/blackfriday/v2"
)

type Context interface {
	FindLabel(id string) (renderer.Label, bool)
	FindResource(path string) (renderer.Resource, bool)
	FindAST(path string) (*blackfriday.Node, bool)
	Resources() chan renderer.Resource
	NumDocuments() int
	NumStatics() int
	NumExtras() int
}

type context struct {
	labels    map[string]renderer.Label
	resources map[string]renderer.Resource
	asts      map[string]*blackfriday.Node
	numDoc    int
	numStatic int
	numExtra  int
}

func newContext() *context {
	return &context{
		labels:    map[string]renderer.Label{},
		resources: map[string]renderer.Resource{},
		asts:      map[string]*blackfriday.Node{},
	}
}

func (c *context) addAST(path string, node *blackfriday.Node) {
	path = util.SlashPath(path)
	c.asts[path] = node
}

func (c *context) addLabel(definedIn, id, name string) renderer.Label {
	definedIn = util.SlashPath(definedIn)
	label := renderer.NewLabel(definedIn, id, name)
	c.labels[id] = label
	return label
}

func (c *context) incResourceCount(typ renderer.ResourceType) {
	switch typ {
	case renderer.MarkdownFile:
		c.numDoc++
	case renderer.StaticFile:
		c.numStatic++
	case renderer.ExtraFile:
		c.numExtra++
	}
}

func (c *context) NumDocuments() int {
	return c.numDoc
}

func (c *context) NumStatics() int {
	return c.numStatic
}

func (c *context) NumExtras() int {
	return c.numExtra
}

func (c *context) addResource(typ renderer.ResourceType, path, title string, toc renderer.TOC) {
	c.incResourceCount(typ)
	path = util.SlashPath(path)
	c.resources[path] = renderer.NewResource(typ, path, title, toc)
}

func (c *context) FindLabel(id string) (renderer.Label, bool) {
	v, ok := c.labels[id]
	return v, ok
}

func (c *context) FindResource(path string) (renderer.Resource, bool) {
	v, ok := c.resources[util.SlashPath(path)]
	return v, ok
}

func (c *context) FindAST(path string) (*blackfriday.Node, bool) {
	v, ok := c.asts[util.SlashPath(path)]
	return v, ok
}

func (c *context) Resources() chan renderer.Resource {
	ch := make(chan renderer.Resource, runtime.NumCPU())
	go func() {
		for _, r := range c.resources {
			ch <- r
		}
		close(ch)
	}()
	return ch
}

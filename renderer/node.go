package renderer

import (
	"bytes"
	"fmt"

	blackfriday "github.com/yuin/blackfriday/v2"
)

type NodeType int

const Any = "*"

const (
	Invalid NodeType = iota
	NodeFencedCode
	NodeFunction
)

type Identifierer interface {
	Identifier() string
}

type FencedCodeBlock interface {
	Identifierer
}

type fencedCodeBlock struct {
	d          *blackfriday.CodeBlockData
	identifier string
}

func newFencedCodeBlock(d *blackfriday.CodeBlockData) FencedCodeBlock {
	info := d.Info
	endOfLang := bytes.IndexAny(info, "\t ")
	if endOfLang < 0 {
		endOfLang = len(info)
	}
	return &fencedCodeBlock{
		d:          d,
		identifier: fmt.Sprintf("%s", info[:endOfLang]),
	}
}

func (b *fencedCodeBlock) Identifier() string {
	return b.identifier
}

type Function interface {
	Identifierer
	Arguments() []interface{}
}

type function struct {
	identifier string
	arguments  []interface{}
}

func newFunction(name string, args []interface{}) Function {
	return &function{
		identifier: name,
		arguments:  args,
	}
}

func (f *function) Identifier() string {
	return f.identifier
}

func (f *function) Arguments() []interface{} {
	return f.arguments
}

type Node interface {
	Identifierer
	Type() NodeType
	Text() []byte
	FencedCodeBlock() FencedCodeBlock
	Function() Function
}

type node struct {
	nodeType        NodeType
	fencedCodeBlock FencedCodeBlock
	function        Function
	node            *blackfriday.Node
	identifier      string
}

func newNode(n *blackfriday.Node) (Node, bool) {
	if n.Type == blackfriday.CodeBlock && n.CodeBlockData.IsFenced {
		fc := newFencedCodeBlock(&n.CodeBlockData)
		return &node{
			nodeType:        NodeFencedCode,
			fencedCodeBlock: newFencedCodeBlock(&n.CodeBlockData),
			node:            n,
			identifier:      fc.Identifier(),
		}, true
	} else if n.Type == blackfriday.Function {
		fn := newFunction(n.FunctionData.Name, n.FunctionData.Arguments)
		return &node{
			nodeType:   NodeFunction,
			node:       n,
			function:   fn,
			identifier: fn.Identifier(),
		}, true
	} else {
		return nil, false
	}
}

func (n *node) Type() NodeType {
	return n.nodeType
}

func (n *node) Text() []byte {
	return n.node.Literal
}

func (n *node) FencedCodeBlock() FencedCodeBlock {
	return n.fencedCodeBlock
}

func (n *node) Function() Function {
	return n.function
}

func (n *node) Identifier() string {
	return n.identifier
}

package renderer

import (
	"bytes"
	"fmt"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type NodeType int

const (
	Invalid NodeType = iota
	NodeFencedCode
	NodeFunction
)

type FencedCodeBlock interface {
	Info() string
}

type fencedCodeBlock struct {
	d    *blackfriday.CodeBlockData
	info string
}

func newFencedCodeBlock(d *blackfriday.CodeBlockData) FencedCodeBlock {
	info := d.Info
	endOfLang := bytes.IndexAny(info, "\t ")
	if endOfLang < 0 {
		endOfLang = len(info)
	}
	return &fencedCodeBlock{
		d:    d,
		info: fmt.Sprintf("%s", info[:endOfLang]),
	}
}

func (b *fencedCodeBlock) Info() string {
	return b.info
}

type Function interface {
	Name() string
}

type function struct {
	name string
}

func newFunction(name string) Function {
	return &function{
		name: name,
	}
}

func (f *function) Name() string {
	return f.name
}

type Node interface {
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
}

func newFunctionNode(n *blackfriday.Node, name string) Node {
	return &node{
		nodeType: NodeFunction,
		node:     n,
		function: newFunction(name),
	}
}

func newNode(n *blackfriday.Node) (Node, bool) {
	if n.Type == blackfriday.CodeBlock && n.CodeBlockData.IsFenced {
		return &node{
			nodeType:        NodeFencedCode,
			fencedCodeBlock: newFencedCodeBlock(&n.CodeBlockData),
			node:            n,
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

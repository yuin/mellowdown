package renderer

import (
	"bytes"
	"fmt"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type NodeType int

const (
	Invalid = iota
	FencedCode
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

type Node interface {
	Type() NodeType
	Text() []byte
	FencedCodeBlock() FencedCodeBlock
}

type node struct {
	nodeType        NodeType
	fencedCodeBlock FencedCodeBlock
	node            *blackfriday.Node
}

func newNode(n *blackfriday.Node) (Node, bool) {
	if n.Type == blackfriday.CodeBlock && n.CodeBlockData.IsFenced {
		return &node{
			nodeType:        FencedCode,
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

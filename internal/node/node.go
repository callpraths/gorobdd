package node

import (
	"fmt"
)

type Type int

const (
	InternalType Type = iota
	LeafType
)

type Node struct {
	Type Type
	Internal
	Leaf
}

type Internal struct {
	True  *Node
	False *Node
}

type Leaf struct {
	Value bool
}

func (n Node) String(v []string) string {
	switch n.Type {
	case InternalType:
		return n.Internal.String(v)
	case LeafType:
		return n.Leaf.String()
	default:
		return fmt.Sprintf("Garbage Node type: %v", n.Type)
	}
}

func (n Leaf) String() string {
	if n.Value {
		return "T"
	} else {
		return "F"
	}
}

func (n Internal) String(v []string) string {
	return fmt.Sprintf("(%v/T: %v, %v/F: %v)", v[0], n.True.String(v[1:]), v[0], n.False.String(v[1:]))
}

func Uniform(depth int, v bool) *Node {
	if depth == 0 {
		return &Node{
			Type: LeafType,
			Leaf: Leaf{v},
		}
	}
	return &Node{
		Type:     InternalType,
		Internal: Internal{True: Uniform(depth-1, v), False: Uniform(depth-1, v)},
	}
}

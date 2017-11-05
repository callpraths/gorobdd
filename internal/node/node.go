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
	// Ply is an index into the vocabular assigning this node to
	// the corresponding variable.
	Ply   int
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
	return fmt.Sprintf(
		"(%v/T: %v, %v/F: %v)",
		v[n.Ply], n.True.String(v), v[n.Ply], n.False.String(v),
	)

}

func Uniform(depth int, v bool) *Node {
	return uniformHelper(0, depth, v)
}

func uniformHelper(ply int, totalPlies int, v bool) *Node {
	if ply == totalPlies {
		return &Node{
			Type: LeafType,
			Leaf: Leaf{v},
		}
	}
	return &Node{
		Type: InternalType,
		Internal: Internal{
			Ply:   ply,
			True:  uniformHelper(ply+1, totalPlies, v),
			False: uniformHelper(ply+1, totalPlies, v),
		},
	}
}

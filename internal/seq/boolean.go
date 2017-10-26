package seq

import (
	"github.com/callpraths/gorobdd/internal/node"
)

func And(a *node.Node, b *node.Node) *node.Node {
	return &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{a.Value && b.Value},
	}
}

func Or(a *node.Node, b *node.Node) *node.Node {
	return &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{a.Value || b.Value},
	}
}

func Not(a *node.Node) *node.Node {
	return &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{!a.Value},
	}
}

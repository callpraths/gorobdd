package node

import (
)

func (a *Node) And(b*Node) *Node {
	return &Node{
		Type: LeafType,
		Leaf: Leaf{ a.Value && b.Value },
	}
}

func (a *Node) Or(b*Node) *Node {
	return &Node{
		Type: LeafType,
		Leaf: Leaf{ a.Value || b.Value },
	}
}

func (a *Node) Not() *Node {
	return &Node {
		Type: LeafType,
		Leaf: Leaf{ ! a.Value },
	}
}


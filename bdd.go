package gorobdd

import (
	"fmt"
)

type NodeLabel string

// Bdd, or Binary Decision Diagram is a balanced binary tree, with branches labelled True and False.
// The leafs are also labeled True and Fal
// The tree is |vocabulary| deep, and each "ply" is corresponds to a name in the vocabulary.
// The order in which vocabulary names match the plys can be changed.
type BDD struct {
	// Names for all the plys.
	Vocabulary []NodeLabel
	*node
}


func (b BDD) String() string {
	return b.node.String(b.Vocabulary)
}

func FromTuples(voc []NodeLabel, tuples [][]bool)  (*BDD, error) {
	bdd := False(voc)
	for _, t:= range tuples {
		if len(t) != len(voc) {
			return nil, fmt.Errorf("Length of tuple %v does not match vocabulary (len: %d)", t, len(voc))
		}
		p := bdd.node
		for _, b := range t {
			if b {
				p = p.True
			} else {
				p = p.False
			}
		}
		p.Value = true
	}
	return bdd, nil
}

func False(voc []NodeLabel) (*BDD) {
	return &BDD{voc, uniform(len(voc), false)}
}

func True(voc []NodeLabel) (*BDD) {
	return &BDD{voc, uniform(len(voc), true)}
}

func uniform(depth int, v bool) *node {
	if depth == 0 {
		return &node {
			Type: leafNodeType,
			leafNode: leafNode { v },
		}
	}
	return &node {
		Type: internalNodeType,
		internalNode: internalNode{True: uniform(depth - 1, v), False: uniform(depth - 1, v)},
	}
}

type nodeType int

const (
	internalNodeType nodeType = iota
	leafNodeType
)

type node struct {
	Type nodeType
	internalNode
	leafNode
}

type internalNode struct {
	True *node
	False *node
}

type leafNode struct {
	Value bool
}

func (n leafNode) String() string {
	if n.Value {
		return "T"
	} else {
		return "F"
	}
}

func (n internalNode) String(v []NodeLabel) string {
	return fmt.Sprintf("(%v/T: %v, %v/F: %v)", v[0], n.True.String(v[1:]), v[0], n.False.String(v[1:]))
}

func (n node) String(v []NodeLabel) string {
	switch(n.Type) {
	case internalNodeType:
		return n.internalNode.String(v)
	case leafNodeType:
		return n.leafNode.String()
	default:
		return fmt.Sprintf("Garbage node type: %v", n.Type)
	}
}

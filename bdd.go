package gorobdd

import (
	"fmt"
	"reflect"
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

func Equal(a *BDD, b *BDD) bool {
	return a.Equal(b)
}

func And(a *BDD, b *BDD) (*BDD, error) {
	return a.And(b)
}

func Or(a *BDD, b *BDD) (*BDD, error) {
	return a.Or(b)
}

func Not(a *BDD) *BDD {
	return a.Not()
}

func (a *BDD) Equal(b *BDD) bool {
	return reflect.DeepEqual(a, b)
}

func (a *BDD) And(b *BDD) (*BDD, error) {
	if ! reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in And: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return &BDD{a.Vocabulary, a.node.And(b.node)}, nil
}

func (a *BDD) Or(b *BDD) (*BDD, error) {
	if ! reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in And: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return &BDD{a.Vocabulary, a.node.Or(b.node)}, nil
}

func (a *BDD) Not() *BDD {
	return &BDD{a.Vocabulary, a.node.Not()}
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

func (a *node) And(b*node) *node {
	return &node{
		Type: leafNodeType,
		leafNode: leafNode{ a.Value && b.Value },
	}
}

func (a *node) Or(b*node) *node {
	return &node{
		Type: leafNodeType,
		leafNode: leafNode{ a.Value || b.Value },
	}
}

func (a *node) Not() *node {
	return &node {
		Type: leafNodeType,
		leafNode: leafNode{ ! a.Value },
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

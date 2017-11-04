// Package seq has sequential implementations of all BDD operations.
// Currently these operations result in BDDs (as opposed to ROBDDs).
// These operations will be optimized in following iterations to work
// entirely on ROBDDs.
package seq

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
)

type leafOp func(a node.Leaf, b node.Leaf) node.Leaf

func And(a *node.Node, b *node.Node) (*node.Node, error) {
	return walk(a, b, andLeafOp)
}

func Or(a *node.Node, b *node.Node) (*node.Node, error) {
	return walk(a, b, orLeafOp)
}

func Not(a *node.Node) *node.Node {
	return &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{!a.Value},
	}
}

func walk(a *node.Node, b *node.Node, op leafOp) (*node.Node, error) {
	if a.Type != b.Type {
		return nil, fmt.Errorf("Mismatched bdd path heights: %v, %v", a, b)
	}
	switch a.Type {
	case node.LeafType:
		return &node.Node{
			Type: node.LeafType,
			Leaf: op(a.Leaf, b.Leaf),
		}, nil
	case node.InternalType:
		tb, e1 := walk(a.True, b.True, op)
		if e1 != nil {
			return tb, e1
		}
		fb, e2 := walk(a.False, b.False, op)
		if e2 != nil {
			return fb, e2
		}
		return &node.Node{
			Type: node.InternalType,
			Internal: node.Internal{True: tb, False: fb},
		}, nil
	default:
		return nil, fmt.Errorf("Unexpected node type: %v in %v", a.Type, a)
	}
}

func orLeafOp(a node.Leaf, b node.Leaf) node.Leaf {
	return node.Leaf{ Value: a.Value || b.Value }
}

func andLeafOp(a node.Leaf, b node.Leaf) node.Leaf {
	return node.Leaf{ Value: a.Value && b.Value }
}

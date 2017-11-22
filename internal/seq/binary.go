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

// GraphEqual determines if the BDDs rooted at the given nodes have identical
// graph structures. For this purpose, Leaf nodes with the same value are considered
// equal.
func GraphEqual(a *node.Node, b *node.Node) (bool, error) {
	if a.Type != b.Type {
		return false, nil
	}
	switch a.Type {
	case node.LeafType:
		return a.Value == b.Value, nil
	case node.InternalType:
		if a.Ply != b.Ply {
			return false, nil
		}
		if eq, e := GraphEqual(a.True, b.True); e != nil || !eq {
			return eq, e
		}
		return GraphEqual(a.False, b.False)
	default:
		return false, fmt.Errorf("Unexpected node type: %v in %v", a.Type, a)
	}

}

// Equal determines if the BDDs rooted at the two nodes are logically equal.
// We can not use Reduce here since Equal will be used in testing Reduce.
func Equal(a *node.Node, b *node.Node) (bool, error) {
	switch a.Type {
	case node.LeafType:
		switch b.Type {
		case node.LeafType:
			return a.Value == b.Value, nil
		case node.InternalType:
			return equalSkippingRoot(b, a)
		default:
			return false, fmt.Errorf("Unknown node type: %v", b)
		}
	case node.InternalType:
		switch b.Type {
		case node.LeafType:
			return equalSkippingRoot(a, b)
		case node.InternalType:
			if a.Ply == b.Ply {
				if r, e := Equal(a.True, b.True); e != nil || !r {
					return r, e
				}
				return Equal(a.False, b.False)
			} else if a.Ply > b.Ply {
				return equalSkippingRoot(a, b)
			} else { // a.Ply < b.Ply
				return equalSkippingRoot(b, a)
			}
		default:
			return false, fmt.Errorf("Unknown node type: %v", b)
		}
	default:
		return false, fmt.Errorf("Unknown node type: %v", a)
	}
}

// equalSkippingRoot skips one level on tall and compares the rest with short.
func equalSkippingRoot(tall *node.Node, short *node.Node) (bool, error) {
	// tall and short are both node.InternalType
	if r, e := Equal(tall.True, tall.False); e != nil || !r {
		return r, e
	}
	return Equal(tall.True, short)
}

// And returns a BDD that represents the conjunction of the given BDDs.
func And(a *node.Node, b *node.Node) (*node.Node, error) {
	return walk(a, b, andLeafOp)
}

// Or returns a BDD that represents the disjunction of the given BDDs.
func Or(a *node.Node, b *node.Node) (*node.Node, error) {
	return walk(a, b, orLeafOp)
}

// Not returns a BDD that represents the negation of the given BDD.
func Not(a *node.Node) *node.Node {
	return &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{
			Value: !a.Value,
		},
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
			Internal: node.Internal{
				Ply:   a.Ply,
				True:  tb,
				False: fb,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Unexpected node type: %v in %v", a.Type, a)
	}
}

func orLeafOp(a node.Leaf, b node.Leaf) node.Leaf {
	return node.Leaf{Value: a.Value || b.Value}
}

func andLeafOp(a node.Leaf, b node.Leaf) node.Leaf {
	return node.Leaf{Value: a.Value && b.Value}
}

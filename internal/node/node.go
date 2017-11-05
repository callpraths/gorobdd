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

func (n Node) String(v ...[]string) string {
	switch n.Type {
	case InternalType:
		return n.Internal.String(v...)
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

func (n Internal) String(v ...[]string) string {
	switch len(v) {
	case 0:
		return fmt.Sprintf(
			"(%v/T: %v, %v/F: %v)",
			n.Ply, n.True.String(v...), n.Ply, n.False.String(v...),
		)
	case 1:
		return fmt.Sprintf(
			"(%v/T: %v, %v/F: %v)",
			v[0][n.Ply], n.True.String(v...), v[0][n.Ply], n.False.String(v...),
		)
	default:
		return fmt.Sprintf("Unepxected vocabulary: %v", v)
	}
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

// Reduce converts the ROBDD rooted at n to the canonical reduced form.
// Reduce is not part of the operations package because it is a transitional operation.
// Eventually, all ROBDDs should be constructed reduced and stay in reduced form during the operations.
func Reduce(n *Node) (*Node, error) {
	t, f, e := findTrueFalse(n, true, true)
	if e != nil {
		return nil, e
	}
	return reduceHelper(n, t, f)
}

func reduceHelper(n *Node, t *Node, f *Node) (*Node, error) {
	switch n.Type {
	case LeafType:
		if n.Value {
			return t, nil
		} else {
			return f, nil
		}
	case InternalType:
		var e error
		n.True, e = reduceHelper(n.True, t, f)
		if e != nil {
			return nil, e
		}
		n.False, e = reduceHelper(n.False, t, f)
		if e != nil {
			return nil, e
		}
		if n.True == n.False {
			return n.True, nil
		} else {
			return n, nil
		}
	default:
		return nil, fmt.Errorf("Unexpected node type: %v", n)
	}
}

// findTrueFalse finds a True and False leaf node to be used to replace all True/False leaves.
func findTrueFalse(n *Node, findTrue bool, findFalse bool) (*Node, *Node, error) {
	switch n.Type {
	case LeafType:
		if n.Value {
			return n, nil, nil
		} else {
			return nil, n, nil
		}
	case InternalType:
		tt, ft, et := findTrueFalse(n.True, findTrue, findFalse)
		if et != nil {
			return nil, nil, et
		}
		findTrue = findTrue && (tt == nil)
		findFalse = findFalse && (ft == nil)
		if !findTrue && !findFalse {
			return tt, ft, nil
		}

		tf, ff, ef := findTrueFalse(n.False, findTrue, findFalse)
		if ef != nil {
			return nil, nil, ef
		}
		if tt == nil {
			tt = tf
		}
		if ft == nil {
			ft = ff
		}
		return tt, ft, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected node type: %v", n)
	}
}

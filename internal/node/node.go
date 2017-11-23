package node

import (
	"fmt"
)

// Type is the type of a node.
type Type int

// Exhaustive set of valid Type's.
const (
	InternalType Type = iota
	LeafType
)

// Tag is an opaque tag that can be attached to a node.
// See also internal/tag
type Tag interface{}

// Node is a single node in the DAG comprising a ROBDD.
type Node struct {
	Type Type
	Internal
	Leaf
	// Tag is an opaque interface allowing Walking functions to
	// attach arbitrary metadata to the Node.
	Tag Tag
}

// Internal is an internal BDD node.
type Internal struct {
	// Ply is an index into the vocabular assigning this node to
	// the corresponding variable.
	Ply   int
	True  *Node
	False *Node
}

// Leaf is a leaf BDD node.
type Leaf struct {
	Value bool
}

// String strigifies a BDD n labeling the plies with the given vocabulary v.
func (n Node) String(v ...[]string) string {
	var ns string
	switch n.Type {
	case InternalType:
		ns = n.Internal.String(v...)
	case LeafType:
		ns = n.Leaf.String()
	default:
		return fmt.Sprintf("Garbage Node type: %v", n.Type)
	}
	if n.Tag == nil {
		return ns
	}
	return fmt.Sprintf("%s#%s|", ns, n.Tag)
}

// String strgifies a leaf BDD node.
func (n Leaf) String() string {
	if n.Value {
		return "T"
	}
	return "F"
}

// String stringifies an internal BDD node labaeling the plies with the given vocabulary v.
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

// Uniform returns a non-reduced balanced tree BDD of given depth where each leaf node
// has the given boolean value.
func Uniform(depth int, v bool) *Node {
	return uniformHelper(0, depth, v)
}

// Clone creates and returns another ROBDD graph with identical graph to the one given.
func Clone(n *Node) *Node {
	return &Node{
		Type: n.Type,
		Internal: Internal{
			Ply:   n.Internal.Ply,
			True:  Clone(n.Internal.True),
			False: Clone(n.Internal.False),
		},
		Leaf: Leaf{
			Value: n.Leaf.Value,
		},
	}
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
		}
		return f, nil
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
		}
		return n, nil
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
		}
		return nil, n, nil
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

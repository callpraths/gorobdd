package gorobdd

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
	"testing"
)

func ExamplePrintLeaf() {
	fmt.Println(&ROBDD{
		[]string{},
		&node.Node{
			Type: node.LeafType,
			Leaf: node.Leaf{true},
		},
	})
	// Output: T
}

func ExamplePrintInternal() {
	fmt.Println(&ROBDD{
		[]string{"a"},
		&node.Node{
			Type: node.InternalType,
			Internal: node.Internal{
				Ply: 0,
				True: &node.Node{
					Type: node.LeafType,
					Leaf: node.Leaf{false},
				},
				False: &node.Node{
					Type: node.LeafType,
					Leaf: node.Leaf{true},
				},
			},
		},
	})
	// Output: (a/T: F, a/F: T)
}

func ExamplePrintUnbalanced() {
	fmt.Println(&ROBDD{
		[]string{"a", "b"},
		&node.Node{
			Type: node.InternalType,
			Internal: node.Internal{
				Ply: 0,
				True: &node.Node{
					Type: node.LeafType,
					Leaf: node.Leaf{false},
				},
				False: &node.Node{
					Type: node.InternalType,
					Internal: node.Internal{
						Ply: 1,
						True: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{false},
						},
						False: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{true},
						},
					},
				},
			},
		},
	})
	// Output: (a/T: F, a/F: (b/T: F, b/F: T))
}

func ExamplePrintPlySkipped() {
	fmt.Println(&ROBDD{
		[]string{"a", "b", "c"},
		&node.Node{
			Type: node.InternalType,
			Internal: node.Internal{
				Ply: 0,
				True: &node.Node{
					Type: node.LeafType,
					Leaf: node.Leaf{false},
				},
				False: &node.Node{
					Type: node.InternalType,
					Internal: node.Internal{
						Ply: 2,
						True: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{false},
						},
						False: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{true},
						},
					},
				},
			},
		},
	})
	// Output: (a/T: F, a/F: (c/T: F, c/F: T))
}

func ExampleTrivialBDDFromTuples() {
	v, _ := FromTuples([]string{}, [][]bool{})
	fmt.Println(v)
	// Output: F
}

func ExampleFalseBDDFromTuples() {
	v, _ := FromTuples([]string{"a"}, [][]bool{})
	fmt.Println(v)
	// Output: (a/T: F, a/F: F)
}

func ExampleBDDFromSingleTuple() {
	v, _ := FromTuples([]string{"a", "b"}, [][]bool{{true, false}})
	fmt.Println(v)
	// Output: (a/T: (b/T: F, b/F: T), a/F: (b/T: F, b/F: F))
}

func ExampleBDDFromTuples() {
	v, _ := FromTuples([]string{"a", "b"}, [][]bool{{true, false}, {false, true}})
	fmt.Println(v)
	// Output: (a/T: (b/T: F, b/F: T), a/F: (b/T: T, b/F: F))
}

func TestBDDFromTuplesChecksTupleLengths(t *testing.T) {
	v, e := FromTuples([]string{"a", "b"}, [][]bool{{true}})
	if e == nil {
		t.Errorf("Unexpected ROBDD from tuples: %v", v)
	}
}

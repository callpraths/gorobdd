package gorobdd

import (
	"fmt"
	"testing"
)

func ExamplePrintLeaf() {
	fmt.Println(&BDD{
		[]NodeLabel{},
		&node{
			Type: leafNodeType,
			leafNode: leafNode{ true },
		},
	})
	// Output: T
}

func ExamplePrintInternal() {
	fmt.Println(&BDD{
		[]NodeLabel{"a"},
		&node{
			Type: internalNodeType,
			internalNode: internalNode {
				True: &node{
					Type: leafNodeType,
					leafNode: leafNode{ false },
				},
				False: &node{
					Type: leafNodeType,
					leafNode: leafNode{ true },
				},
			},
		},
	})
	// Output: (a/T: F, a/F: T)
}

func ExampleTrivialBDDFromTuples() {
	v, _ := FromTuples([]NodeLabel{}, [][]bool{})
	fmt.Println(v)
	// Output: F
}

func ExampleFalseBDDFromTuples() {
	v, _ := FromTuples([]NodeLabel{"a"}, [][]bool{})
	fmt.Println(v)
	// Output: (a/T: F, a/F: F)
}

func ExampleBDDFromSingleTuple() {
	v, _ := FromTuples([]NodeLabel{"a", "b"}, [][]bool{{true, false}})
	fmt.Println(v)
	// Output: (a/T: (b/T: F, b/F: T), a/F: (b/T: F, b/F: F))
}

func ExampleBDDFromTuples() {
	v, _ := FromTuples([]NodeLabel{"a", "b"}, [][]bool{{true, false}, {false, true}})
	fmt.Println(v)
	// Output: (a/T: (b/T: F, b/F: T), a/F: (b/T: T, b/F: F))
}

func TestBDDFromTuplesChecksTupleLengths(t *testing.T) {
	v, e := FromTuples([]NodeLabel{"a", "b"}, [][]bool{{true}})
	if e == nil {
		t.Error("Unexpected BDD from tuples: %v", v)
	}
}

func ExampleTrivialBDDBasicOps() {
	bt := True([]NodeLabel{})
	bf := False([]NodeLabel{})
	// TODO
	fmt.Println(bt)
	fmt.Println(bf)
	// Output: 
	// T
	// F
}

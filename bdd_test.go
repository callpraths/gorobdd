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

func TestBinaryOpsCheckVocabulary(t *testing.T) {
	var tests = []struct{
		lhs *BDD
		rhs *BDD
	}{
		{True([]NodeLabel{"a"}), True([]NodeLabel{"a", "b"})},
		{True([]NodeLabel{"a", "b"}), True([]NodeLabel{"a"})},
		{True([]NodeLabel{"a", "b"}), True([]NodeLabel{})},
		{True([]NodeLabel{"a", "b"}), True([]NodeLabel{"b", "a"})},
		{True([]NodeLabel{"a", "b"}), True([]NodeLabel{"a", "a"})},
	}
	for _, tt := range tests {
		if _, e := Equal(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from And(%v, %v)", tt.lhs, tt.rhs)
		}
		if _, e := And(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from And(%v, %v)", tt.lhs, tt.rhs)
		}
		if _, e := Or(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from Or(%v, %v)", tt.lhs, tt.rhs)
		}
	}
}

func TestBDDEqual(t *testing.T) {
	var tests = []struct {
		lhs *BDD
		rhs *BDD
		eq bool
	}{
		{True([]NodeLabel{}), True([]NodeLabel{}), true},
		{False([]NodeLabel{}), False([]NodeLabel{}), true},
		{True([]NodeLabel{}), False([]NodeLabel{}), false},
		{False([]NodeLabel{}), True([]NodeLabel{}), false},

	}
	for _, tt := range tests {
		eq, e := Equal(tt.lhs, tt.rhs)
		if e != nil {
			t.Errorf("Equal(%v, %v) failed: %v", tt.lhs, tt.rhs, e)
		}
		if eq != tt.eq {
			t.Errorf("Equal(%v, %v) = %v, want %v", tt.lhs, tt.rhs, eq, tt.eq)
		}
	}
}

func TestTrivialBDDBinaryOps(t *testing.T) {
	var tests = []struct{
		lhs *BDD
		rhs *BDD
		and *BDD
		or *BDD
	} {
		{True([]NodeLabel{}), True([]NodeLabel{}), True([]NodeLabel{}), True([]NodeLabel{})},
		{True([]NodeLabel{}), False([]NodeLabel{}), False([]NodeLabel{}), True([]NodeLabel{})},
		{False([]NodeLabel{}), True([]NodeLabel{}), False([]NodeLabel{}), True([]NodeLabel{})},
		{False([]NodeLabel{}), False([]NodeLabel{}), False([]NodeLabel{}), False([]NodeLabel{})},
	}
	for _, tt := range tests {
		var and, or *BDD
		var eq bool
		var e error
		and, e = And(tt.lhs, tt.rhs)
		if e != nil {
			t.Errorf("And(%v, %v) returned error %v", tt.lhs, tt.rhs, e)
		}
		eq, e = Equal(and, tt.and)
		if e != nil {
			t.Errorf("And(%v, %v) returned error %v", and, tt.and, e)
		}
		if ! eq {
			t.Errorf("And(%v, %v) = %v, want %v", tt.lhs, tt.rhs, and, tt.and)
		}
		or, e = Or(tt.lhs, tt.rhs)
		if e != nil {
			t.Errorf("Or(%v, %v) returned error %v", tt.lhs, tt.rhs, e)
		}
		eq, e = Equal(or, tt.or)
		if e != nil {
			t.Errorf("And(%v, %v) returned error %v", and, tt.and, e)
		}
		if ! eq {
			t.Errorf("Or(%v, %v) = %v, want %v", tt.lhs, tt.rhs, or, tt.or)
		}
	}
}

func TestTrivialBDDNot(t *testing.T) {
	var tests = []struct{
		in *BDD
		ans *BDD
	} {
		{True([]NodeLabel{}), False([]NodeLabel{})},
		{False([]NodeLabel{}), True([]NodeLabel{})},
	}
	for _, tt := range tests {
		ans, e1 := Not(tt.in)
		if e1 != nil {
			t.Errorf("Not(%v) returned error %v", tt.in, e1)
		}
		eq, e2 := Equal(ans, tt.ans)
		if e2 != nil {
			t.Errorf("Equal(%v, %v) returned error %v", ans, tt.ans, e2)
		}
		if ! eq {
			t.Errorf("Not(%v) = %v, want %v", tt.in, ans, tt.ans)
		}
	}
}

package gorobdd

import (
	"github.com/callpraths/gorobdd/internal/node"
	"github.com/callpraths/gorobdd/internal/tag"
	"testing"
)

func TestBinaryOpsCheckVocabulary(t *testing.T) {
	var tests = []struct {
		lhs *ROBDD
		rhs *ROBDD
	}{
		{True([]string{"a"}), True([]string{"a", "b"})},
		{True([]string{"a", "b"}), True([]string{"a"})},
		{True([]string{"a", "b"}), True([]string{})},
		{True([]string{"a", "b"}), True([]string{"b", "a"})},
		{True([]string{"a", "b"}), True([]string{"a", "a"})},
	}
	for _, tt := range tests {
		if _, e := Equal(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from Equal(%v, %v)", tt.lhs, tt.rhs)
		}
		if _, e := And(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from And(%v, %v)", tt.lhs, tt.rhs)
		}
		if _, e := Or(tt.lhs, tt.rhs); e == nil {
			t.Errorf("No error raised from Or(%v, %v)", tt.lhs, tt.rhs)
		}
	}
}

func fromTuplesNoError(t *testing.T, v []string, tu [][]bool) *ROBDD {
	b, e := FromTuples(v, tu)
	if e != nil {
		t.Fatalf("FromTuples(%v, %v) returned error: %v", v, tu, e)
	}
	return b
}

func TestBDDEqual(t *testing.T) {
	type testCase struct {
		lhs *ROBDD
		rhs *ROBDD
		eq  bool
		gEq bool
	}
	tcs := []testCase{
		{True([]string{}), True([]string{}), true, true},
		{False([]string{}), False([]string{}), true, true},
		{True([]string{}), False([]string{}), false, false},
		{False([]string{}), True([]string{}), false, false},
		{True([]string{"a"}), True([]string{"a"}), true, true},
		{False([]string{"a"}), False([]string{"a"}), true, true},
		{True([]string{"a"}), False([]string{"a"}), false, false},
		{False([]string{"a"}), True([]string{"a"}), false, false},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			true,
			true,
		},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{false, false}}),
			false,
			false,
		},
	}

	tfunc := func(tt testCase) {
		eq, e := Equal(tt.lhs, tt.rhs)
		if e != nil {
			t.Errorf("Equal(%v, %v) failed: %v", tt.lhs, tt.rhs, e)
		}
		if eq != tt.eq {
			t.Errorf("Equal(%v, %v) = %v, want %v", tt.lhs, tt.rhs, eq, tt.eq)
		}
		gEq, e2 := GraphEqual(tt.lhs, tt.rhs)
		if e2 != nil {
			t.Errorf("GraphEqual(%v, %v) failed: %v", tt.lhs, tt.rhs, e2)
		}
		if gEq != tt.gEq {
			t.Errorf("GraphEqual(%v, %v) = %v, want %v", tt.lhs, tt.rhs, gEq, tt.gEq)
		}
	}

	s := tag.NewSeenContext()
	for _, tt := range tcs {
		tfunc(tt)
		// Equal operations should completely ignore tags.
		s.MarkSeen(tt.lhs.Node)
		tfunc(tt)
	}
}

func TestBDDEqualLogicallyNotEqualStructurally(t *testing.T) {
	a1 := &ROBDD{
		[]string{"a", "b"},
		&node.Node{
			Type: node.InternalType,
			Internal: node.Internal{
				Ply: 0,
				True: &node.Node{
					Type: node.LeafType,
					Leaf: node.Leaf{
						Value: false,
					},
				},
				False: &node.Node{
					Type: node.InternalType,
					Internal: node.Internal{
						Ply: 1,
						True: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{
								Value: false,
							},
						},
						False: &node.Node{
							Type: node.LeafType,
							Leaf: node.Leaf{
								Value: true,
							},
						},
					},
				},
			},
		},
	}
	tn := &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{
			Value: true,
		},
	}
	fn := &node.Node{
		Type: node.LeafType,
		Leaf: node.Leaf{
			Value: false,
		},
	}
	bt := &node.Node{
		Type:     node.InternalType,
		Internal: node.Internal{Ply: 1, True: fn, False: fn},
	}
	bf := &node.Node{
		Type:     node.InternalType,
		Internal: node.Internal{Ply: 1, True: fn, False: tn},
	}
	a2 := &ROBDD{
		[]string{"a", "b"},
		&node.Node{
			Type:     node.InternalType,
			Internal: node.Internal{Ply: 0, True: bt, False: bf},
		},
	}

	eq, e := Equal(a1, a2)
	if e != nil {
		t.Errorf("Equal(%v, %v) failed: %v", a1, a2, e)
	}
	if !eq {
		t.Errorf("Equal(%v, %v) = false, want true", a1, a2)
	}
	gEq, e2 := GraphEqual(a1, a2)
	if e2 != nil {
		t.Errorf("Equal(%v, %v) failed: %v", a1, a2, e2)
	}
	if gEq {
		t.Errorf("GraphEqual(%v, %v) = true, want false", a1, a2)
	}
}

func TestBDDBooleanOps(t *testing.T) {
	var tests = []struct {
		lhs *ROBDD
		rhs *ROBDD
		and *ROBDD
		or  *ROBDD
	}{
		{True([]string{}), True([]string{}), True([]string{}), True([]string{})},
		{True([]string{}), False([]string{}), False([]string{}), True([]string{})},
		{False([]string{}), True([]string{}), False([]string{}), True([]string{})},
		{False([]string{}), False([]string{}), False([]string{}), False([]string{})},
		{True([]string{"a"}), True([]string{"a"}), True([]string{"a"}), True([]string{"a"})},
		{True([]string{"a"}), False([]string{"a"}), False([]string{"a"}), True([]string{"a"})},
		{False([]string{"a"}), True([]string{"a"}), False([]string{"a"}), True([]string{"a"})},
		{False([]string{"a"}), False([]string{"a"}), False([]string{"a"}), False([]string{"a"})},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}, {false, true}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {true, false}, {false, true}, {false, false}}),
		},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
		},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {true, false}, {false, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {true, false}, {false, true}, {false, false}}),
		},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, true}, {false, false}}),
		},
	}
	for _, tt := range tests {
		var and, or *ROBDD
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
		if !eq {
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
		if !eq {
			t.Errorf("Or(%v, %v) = %v, want %v", tt.lhs, tt.rhs, or, tt.or)
		}
	}
}

func TestTrivialBDDNot(t *testing.T) {
	var tests = []struct {
		in  *ROBDD
		ans *ROBDD
	}{
		{True([]string{}), False([]string{})},
		{False([]string{}), True([]string{})},
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
		if !eq {
			t.Errorf("Not(%v) = %v, want %v", tt.in, ans, tt.ans)
		}
	}
}

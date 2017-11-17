package gorobdd

import (
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
	var tests = []struct {
		lhs *ROBDD
		rhs *ROBDD
		eq  bool
	}{
		{True([]string{}), True([]string{}), true},
		{False([]string{}), False([]string{}), true},
		{True([]string{}), False([]string{}), false},
		{False([]string{}), True([]string{}), false},
		{True([]string{"a"}), True([]string{"a"}), true},
		{False([]string{"a"}), False([]string{"a"}), true},
		{True([]string{"a"}), False([]string{"a"}), false},
		{False([]string{"a"}), True([]string{"a"}), false},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			true,
		},
		{
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{true, false}}),
			fromTuplesNoError(t, []string{"a", "b"}, [][]bool{{false, false}}),
			false,
		},
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
func TestBDDBinaryOps(t *testing.T) {
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

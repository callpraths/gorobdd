package gorobdd

import (
	"testing"
)

func TestBDDBinaryOps(t *testing.T) {
	var tests = []struct {
		lhs *BDD
		rhs *BDD
		and *BDD
		or  *BDD
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
		in  *BDD
		ans *BDD
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

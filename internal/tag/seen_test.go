package tag

import (
	"github.com/callpraths/gorobdd/internal/node"
	"testing"
)

func TestIsSeen(t *testing.T) {
	c := NewSeenContext()
	var tests = []struct{
		n *node.Node
		s bool
	}{
		{&node.Node{Tag: nil}, false},
		{&node.Node{Tag: seenTag{true, NewSeenContext()}}, false},
		{&node.Node{Tag: seenTag{false, c}}, false},
		{&node.Node{Tag: seenTag{true, c}}, true},
	}
	for _, tt := range tests {
		s := c.IsSeen(tt.n)
		if s != tt.s {
			t.Errorf("%v.IsSeen(%v) = %v, want %v", c, tt.n, s, tt.s)
		}
	}

}

func TestMarkSeen(t *testing.T) {
	c := NewSeenContext()
	n := &node.Node{Tag: nil}
	if c.IsSeen(n) {
		t.Errorf("%v.IsSeen(%v) = true, want false", c, n)
	}

	c.MarkSeen(n)
	if !c.IsSeen(n) {
		t.Errorf("%v.IsSeen(%v) = false, want true", c, n)
	}

	c2 := NewSeenContext()
	if c2.IsSeen(n) {
		t.Errorf("%v.IsSeen(%v) = true, want false", c2, n)
	}
}

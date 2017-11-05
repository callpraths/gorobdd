package node

import (
	"testing"
)

func TestCountTreeStructure(t *testing.T) {
	var tests = []struct {
		b *Node
		c int
	}{
		{&Node{Type: LeafType, Leaf: Leaf{true}}, 1},
		{&Node{Type: LeafType, Leaf: Leaf{false}}, 1},
		{
			&Node{
				Type: InternalType,
				Internal: Internal{
					Ply:   1,
					True:  &Node{Type: LeafType, Leaf: Leaf{true}},
					False: &Node{Type: LeafType, Leaf: Leaf{false}},
				},
			},
			3,
		},
		{
			&Node{
				Type: InternalType,
				Internal: Internal{
					Ply: 1,
					True: &Node{
						Type: InternalType,
						Internal: Internal{
							Ply:   2,
							True:  &Node{Type: LeafType, Leaf: Leaf{false}},
							False: &Node{Type: LeafType, Leaf: Leaf{false}},
						},
					},
					False: &Node{Type: LeafType, Leaf: Leaf{false}},
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		c, e := CountNodes(tt.b)
		if e != nil {
			t.Errorf("Encountered error in CountNodes(%v): %v", tt.b, e)
		}
		if c != tt.c {
			t.Errorf("CountNodes(%v) = %v, want %v", tt.b.String(), c, tt.c)
		}
	}
}

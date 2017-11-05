package gorobdd

import (
	"github.com/callpraths/gorobdd/internal/node"
	"testing"
)

func TestCountTreeStructure(t *testing.T) {
	var tests = []struct {
		b *ROBDD
		c int
	}{
		{&ROBDD{[]string{""}, &node.Node{Type: node.LeafType, Leaf: node.Leaf{true}}}, 1},
		{&ROBDD{[]string{""}, &node.Node{Type: node.LeafType, Leaf: node.Leaf{false}}}, 1},
		{
			&ROBDD{
				[]string{"a"},
				&node.Node{
					Type: node.InternalType,
					Internal: node.Internal{
						Ply:   1,
						True:  &node.Node{Type: node.LeafType, Leaf: node.Leaf{true}},
						False: &node.Node{Type: node.LeafType, Leaf: node.Leaf{false}},
					},
				},
			},
			3,
		},
		{
			&ROBDD{
				[]string{"a", "b"},
				&node.Node{
					Type: node.InternalType,
					Internal: node.Internal{
						Ply: 1,
						True: &node.Node{
							Type: node.InternalType,
							Internal: node.Internal{
								Ply:   2,
								True:  &node.Node{Type: node.LeafType, Leaf: node.Leaf{false}},
								False: &node.Node{Type: node.LeafType, Leaf: node.Leaf{false}},
							},
						},
						False: &node.Node{Type: node.LeafType, Leaf: node.Leaf{false}},
					},
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

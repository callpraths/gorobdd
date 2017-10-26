package gorobdd

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
	"reflect"
)

// Bdd, or Binary Decision Diagram is a balanced binary tree, with branches labelled True and False.
// The leafs are also labeled True and Fal
// The tree is |vocabulary| deep, and each "ply" is corresponds to a name in the vocabulary.
// The order in which vocabulary names match the plys can be changed.
type BDD struct {
	// Names for all the plys.
	Vocabulary []string
	*node.Node
}


func (b BDD) String() string {
	return b.Node.String(b.Vocabulary)
}

func FromTuples(voc []string, tuples [][]bool)  (*BDD, error) {
	bdd := False(voc)
	for _, t:= range tuples {
		if len(t) != len(voc) {
			return nil, fmt.Errorf("Length of tuple %v does not match vocabulary (len: %d)", t, len(voc))
		}
		p := bdd.Node
		for _, b := range t {
			if b {
				p = p.True
			} else {
				p = p.False
			}
		}
		p.Value = true
	}
	return bdd, nil
}

func False(voc []string) (*BDD) {
	return &BDD{voc, node.Uniform(len(voc), false)}
}

func True(voc []string) (*BDD) {
	return &BDD{voc, node.Uniform(len(voc), true)}
}

func Equal(a *BDD, b *BDD) (bool, error) {
	return a.Equal(b)
}

func And(a *BDD, b *BDD) (*BDD, error) {
	return a.And(b)
}

func Or(a *BDD, b *BDD) (*BDD, error) {
	return a.Or(b)
}

func Not(a *BDD) (*BDD, error) {
	return a.Not()
}

func (a *BDD) Equal(b *BDD) (bool, error) {
	if ! reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return false, fmt.Errorf("Mismatched vocabularies in Equal: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return reflect.DeepEqual(a, b), nil
}

func (a *BDD) And(b *BDD) (*BDD, error) {
	if ! reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in And: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return &BDD{a.Vocabulary, a.Node.And(b.Node)}, nil
}

func (a *BDD) Or(b *BDD) (*BDD, error) {
	if ! reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in Or: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return &BDD{a.Vocabulary, a.Node.Or(b.Node)}, nil
}

func (a *BDD) Not() (*BDD, error) {
	return &BDD{a.Vocabulary, a.Node.Not()}, nil
}


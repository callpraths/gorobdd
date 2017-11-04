/*
Package gorobdd is an implementation of Reduced Ordered Binary Decision Diagrams.

This package is intended to implement ROBDD operations in a few different ways
for comparison and fun.
Details about the data structure can be found at:
  - http://people.cs.aau.dk/~srba/courses/SV-08/material/09.pdf
*/
package gorobdd

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
	"reflect"
)

// Bdd, or Binary Decision Diagram is a balanced binary tree, with branches labelled True and False.
// The leafs are also labeled True and False.
// The tree is |vocabulary| deep, and each "ply" corresponds to a name in the vocabulary.
// The order in which vocabulary names match the plys can be changed.
type BDD struct {
	// Names for all the plys.
	// Order of vocabulary is significant.
	Vocabulary []string
	*node.Node
}

// String is part of the implementation of Stringer interface.
func (b BDD) String() string {
	return b.Node.String(b.Vocabulary)
}

// FromTuples allows constructing a BDD from a boolean expression in DNF.
// voc is the vocabulary to be used for the BDD.
// tuples lists the values assumed by each ply for which the expression evaluates to True.
func FromTuples(voc []string, tuples [][]bool) (*BDD, error) {
	bdd := False(voc)
	for _, t := range tuples {
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

// False returns the BDD with value False for the given vocabulary.
func False(voc []string) *BDD {
	return &BDD{voc, node.Uniform(len(voc), false)}
}

// True returns the BDD with value True for the given vocabulary.
func True(voc []string) *BDD {
	return &BDD{voc, node.Uniform(len(voc), true)}
}

// Equal determines if two BDDs correspond to equivalent boolean expressions.
func Equal(a *BDD, b *BDD) (bool, error) {
	if !reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return false, fmt.Errorf("Mismatched vocabularies in Equal: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	return reflect.DeepEqual(a, b), nil
}

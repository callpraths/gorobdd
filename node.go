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
)

// ROBDD is short for Reduced Ordered Binary Decision Diagram, a concise
// normalized representation of boolean formulas.  It is a DAG. Internal nodes
// have two branches labeled True and False, and unique True and/or False
// leafs.  Additionally, each internal node is labeled with a variable from a
// given vocabulary, and variables from the vocabulary occur in a fixed order
// (skipping some) on any path down the DAG.  The order in which vocabulary
// names match the plys can be changed.
type ROBDD struct {
	// Names for all the plys.
	// Order of vocabulary is significant.
	Vocabulary []string
	*node.Node
}

// String is part of the implementation of Stringer interface.
func (b ROBDD) String() string {
	return b.Node.String(b.Vocabulary)
}

// FromTuples allows constructing a ROBDD from a boolean expression in DNF.
// voc is the vocabulary to be used for the ROBDD.
// tuples lists the values assumed by each ply for which the expression evaluates to True.
func FromTuples(voc []string, tuples [][]bool) (*ROBDD, error) {
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

// False returns the ROBDD with value False for the given vocabulary.
func False(voc []string) *ROBDD {
	return &ROBDD{voc, node.Uniform(len(voc), false)}
}

// True returns the ROBDD with value True for the given vocabulary.
func True(voc []string) *ROBDD {
	return &ROBDD{voc, node.Uniform(len(voc), true)}
}

// Reduce converts an ROBDD to the the reduced canonicalized form.
func Reduce(a *ROBDD) (*ROBDD, error) {
	n, e := node.Reduce(a.Node)
	if e != nil {
		return nil, e
	}
	return &ROBDD{a.Vocabulary, n}, nil
}

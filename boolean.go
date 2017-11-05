package gorobdd

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/seq"
	"reflect"
)

// And computes the conjuction of the boolean expressions corresponding to given BDDs.
func And(a *ROBDD, b *ROBDD) (*ROBDD, error) {
	if !reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in And: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	r, e := seq.And(a.Node, b.Node)
	if e != nil {
		return nil, e
	}
	return &ROBDD{a.Vocabulary, r}, nil
}

// Or computes the disjunction of the boolean expressions corresponding to given BDDs.
func Or(a *ROBDD, b *ROBDD) (*ROBDD, error) {
	if !reflect.DeepEqual(a.Vocabulary, b.Vocabulary) {
		return nil, fmt.Errorf("Mismatched vocabularies in Or: %v, %v", a.Vocabulary, b.Vocabulary)
	}
	r, e := seq.Or(a.Node, b.Node)
	if e != nil {
		return nil, e
	}
	return &ROBDD{a.Vocabulary, r}, nil
}

// Not computes the negation of the boolean expression corresponding to given ROBDD.
func Not(a *ROBDD) (*ROBDD, error) {
	return &ROBDD{a.Vocabulary, seq.Not(a.Node)}, nil
}

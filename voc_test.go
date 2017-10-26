package gorobdd

import (
	"fmt"
	"testing"
)

func ExampleNewVocabulary() {
	fmt.Println(NewVocabulary())
	fmt.Println(NewVocabulary("a", "b"))
	// Output:
	// Voc: [] <nil>
	// Voc: [a b] <nil>
}

func TestDetectsDuplicate(t *testing.T) {
	if _, e := NewVocabulary("a", "a"); e == nil {
		t.Errorf("No error raised on duplicate vocabulary")
	}
}

package gorobdd

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"testing"
)

func ExampleReduceTrivial() {
	fmt.Println(Reduce(True([]string{})))
	fmt.Println(Reduce(False([]string{})))
	fmt.Println(Reduce(True([]string{"a"})))
	fmt.Println(Reduce(False([]string{"a"})))
	fmt.Println(Reduce(True([]string{"a", "b"})))
	fmt.Println(Reduce(False([]string{"a", "b"})))
	// Output:
	// T <nil>
	// F <nil>
	// T <nil>
	// F <nil>
	// T <nil>
	// F <nil>
}

func ExampleReduceSkipsPlies() {
	n, _ := FromTuples(
		[]string{"a", "b"},
		[][]bool{{true, true}, {true, false}},
	)
	fmt.Println(Reduce(n))
	// Output:
	// (a/T: T, a/F: F) <nil>
}

func TestReduceSkipsPlies(t *testing.T) {
	var es *multierror.Error
	n, e := FromTuples(
		[]string{"a", "b"},
		[][]bool{{true, true}, {true, false}},
	)
	multierror.Append(es, e)
	r, e := FromTuples(
		[]string{"a", "b"},
		[][]bool{{true, true}, {true, false}},
	)
	multierror.Append(es, e)

	r, e = Reduce(r)
	multierror.Append(es, e)

	b, e := Equal(r, n)
	multierror.Append(es, e)
	if !b {
		t.Errorf("Equal(%v, %v)= %v, want equal", r, n, b)
	}

	c, e := CountNodes(n)
	multierror.Append(es, e)
	if c != 7 {
		t.Errorf("CountNodes(%v) = %v, want %v", n, c, 7)
	}
	c, e = CountNodes(r)
	multierror.Append(es, e)
	if c != 3 {
		t.Errorf("CountNodes(%v) = %v, want %v", r, c, 3)
	}

	e = es.ErrorOrNil()
	if e != nil {
		t.Errorf("Encountered errors %s", e)
	}
}

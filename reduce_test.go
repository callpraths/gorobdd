package gorobdd

import (
	"fmt"
	//"testing"
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

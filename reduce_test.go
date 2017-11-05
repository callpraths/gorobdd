package gorobdd

import (
	"fmt"
	//"testing"
)

func ExampleReduceTrivial() {
	fmt.Println(Reduce(True([]string{})))
	fmt.Println(Reduce(False([]string{})))
	// Output:
	// T <nil>
	// F <nil>
}

func ExampleReduceSkipsLevel() {
	fmt.Println(Reduce(True([]string{"a"})))
	fmt.Println(Reduce(False([]string{"a"})))
	// Output:
	// T <nil>
	// F <nil>
}

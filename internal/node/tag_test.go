package node

import (
	"fmt"
)

func tagWithPath(n *Node, p string) {
	n.Tag = p
	switch n.Type {
	case InternalType:
		tagWithPath(n.True, fmt.Sprintf("%s1", p))
		tagWithPath(n.False, fmt.Sprintf("%s0", p))
	}
}
func ExamplePathTag() {
	n := Uniform(2, true)
	fmt.Println(n.String())
	tagWithPath(n, "")
	fmt.Println(n.String())
	CleanTags(n)
	fmt.Println(n.String())
	// Output:
	// (0/T: (1/T: T, 1/F: T), 0/F: (1/T: T, 1/F: T))
	// (0/T: (1/T: T#11|, 1/F: T#10|)#1|, 0/F: (1/T: T#01|, 1/F: T#00|)#0|)#|
	// (0/T: (1/T: T, 1/F: T), 0/F: (1/T: T, 1/F: T))
}

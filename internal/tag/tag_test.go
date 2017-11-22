package tag

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
)

func tagWithPath(n *node.Node, p string) {
	n.Tag = p
	switch n.Type {
	case node.InternalType:
		tagWithPath(n.True, fmt.Sprintf("%s1", p))
		tagWithPath(n.False, fmt.Sprintf("%s0", p))
	}
}

func ExampleTag() {
	n := node.Uniform(2, true)
	fmt.Println(n.String())
	tagWithPath(n, "")
	fmt.Println(n.String())
	Clean(n)
	fmt.Println(n.String())
	// Output:
	// (0/T: (1/T: T, 1/F: T), 0/F: (1/T: T, 1/F: T))
	// (0/T: (1/T: T#11|, 1/F: T#10|)#1|, 0/F: (1/T: T#01|, 1/F: T#00|)#0|)#|
	// (0/T: (1/T: T, 1/F: T), 0/F: (1/T: T, 1/F: T))
}

package node

import (
	"fmt"
)

func CountNodes(n *Node) (int, error) {
	switch n.Type {
	case LeafType:
		return 1, nil
	case InternalType:
		t, et := CountNodes(n.True)
		if et != nil {
			return t, et
		}
		f, ef := CountNodes(n.False)
		if ef != nil {
			return f, ef
		}
		return t + f + 1, nil
	default:
		return -1, fmt.Errorf("Malformed node: %v", n)
	}
}

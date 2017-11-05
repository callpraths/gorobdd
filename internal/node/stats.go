package node

import (
	"fmt"
)

func CountNodes(n *Node) (int, error) {
	m := make(map[*Node]bool)
	return countNodesHelper(n, m)
}

func countNodesHelper(n *Node, m map[*Node]bool) (int, error) {
	_, seen := m[n]
	if seen {
		return 0, nil
	}
	m[n] = true
	switch n.Type {
	case LeafType:
		return 1, nil
	case InternalType:
		t, et := countNodesHelper(n.True, m)
		if et != nil {
			return t, et
		}
		f, ef := countNodesHelper(n.False, m)
		if ef != nil {
			return f, ef
		}
		return t + f + 1, nil
	default:
		return -1, fmt.Errorf("Malformed node: %v", n)
	}

}

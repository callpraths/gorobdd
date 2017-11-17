package stats

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
)

func CountNodes(n *node.Node) (int, error) {
	m := make(map[*node.Node]bool)
	return countNodesHelper(n, m)
}

func countNodesHelper(n *node.Node, m map[*node.Node]bool) (int, error) {
	_, seen := m[n]
	if seen {
		return 0, nil
	}
	m[n] = true
	switch n.Type {
	case node.LeafType:
		return 1, nil
	case node.InternalType:
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

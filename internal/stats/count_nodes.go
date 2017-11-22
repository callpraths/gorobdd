package stats

import (
	"fmt"
	"github.com/callpraths/gorobdd/internal/node"
	"github.com/callpraths/gorobdd/internal/tag"
)

// CountNodes counts the total number of unique graph nodes reachable from The
// given root node.
func CountNodes(n *node.Node) (int, error) {
	s := tag.NewSeenContext()
	return countNodesHelper(n, s)
}

func countNodesHelper(n *node.Node, s tag.SeenContext) (int, error) {
	if s.IsSeen(n) {
		return 0, nil
	}
	s.MarkSeen(n)
	switch n.Type {
	case node.LeafType:
		return 1, nil
	case node.InternalType:
		t, et := countNodesHelper(n.True, s)
		if et != nil {
			return t, et
		}
		f, ef := countNodesHelper(n.False, s)
		if ef != nil {
			return f, ef
		}
		return t + f + 1, nil
	default:
		return -1, fmt.Errorf("Malformed node: %v", n)
	}

}

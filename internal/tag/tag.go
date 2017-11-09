package tag

import (
	"github.com/callpraths/gorobdd/internal/node"
)

// Clean resets tags on the Node graph rooted at n.
// TODO(callpraths): Use a smart Tagger to clean this up without walking
// nodes multiple times.
func Clean(n *node.Node) {
	n.Tag = nil
	switch n.Type {
	case node.InternalType:
		Clean(n.True)
		Clean(n.False)
	}
}

// Package tag provides helper functions to work with the Tag member
// of the Node struct. To tag nodes during an operation, create a new context (e.g. SeenContext) and
// tag all nodes using that context. This ensures that any stale tags from an earlier context do not
// intefere with the new tags.
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

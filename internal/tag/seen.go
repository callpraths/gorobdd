package tag

import (
	"github.com/callpraths/gorobdd/internal/node"
	uuid "github.com/satori/go.uuid"
)

// SeenContext is a tagging context used to mark nodes as having been visited during the current
// operation.
type SeenContext uuid.UUID

// NewSeenContext returns an initialized SeenContext.
func NewSeenContext() SeenContext {
	return SeenContext(uuid.NewV4())
}

// IsSeen returns whether the given node was marked visited using the current context.
func (c SeenContext) IsSeen(n *node.Node) bool {
	switch t := n.Tag.(type) {
	case seenTag:
		return uuid.Equal(uuid.UUID(c), uuid.UUID(t.SeenContext)) && t.marked
	default:
		return false
	}
}

// MarkSeen marks the given node as having been visited in the current context.
func (c SeenContext) MarkSeen(n *node.Node) {
	n.Tag = seenTag{true, c}
}

type seenTag struct {
	marked bool
	SeenContext
}

func (t seenTag) String() string {
	if t.marked {
 		return "s"
        }
	return ""
}

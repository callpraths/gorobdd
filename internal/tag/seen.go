package tag

import (
	"github.com/callpraths/gorobdd/internal/node"
	uuid "github.com/satori/go.uuid"
)

type SeenContext uuid.UUID

func NewSeenContext() SeenContext {
	return SeenContext(uuid.NewV4())
}

type seenTag struct {
	marked bool
	SeenContext
}

func (c SeenContext) IsSeen(n *node.Node) bool {
	switch t := n.Tag.(type) {
	case seenTag:
		return uuid.Equal(uuid.UUID(c), uuid.UUID(t.SeenContext)) && t.marked
	default:
		return false
	}
}

func (c SeenContext) MarkSeen(n *node.Node) {
	n.Tag = seenTag{true, c}
}

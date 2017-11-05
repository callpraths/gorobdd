package gorobdd

import (
	"github.com/callpraths/gorobdd/internal/node"
)

func CountNodes(b *ROBDD) (int, error) {
	return node.CountNodes(b.Node)
}

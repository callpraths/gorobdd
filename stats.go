package gorobdd

import (
	"github.com/callpraths/gorobdd/internal/stats"
)

func CountNodes(b *ROBDD) (int, error) {
	return stats.CountNodes(b.Node)
}

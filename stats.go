package gorobdd

import (
	"github.com/callpraths/gorobdd/internal/stats"
)

// CountNodes counts the total number of unique BDD nodes in the given ROBDD.
func CountNodes(b *ROBDD) (int, error) {
	return stats.CountNodes(b.Node)
}

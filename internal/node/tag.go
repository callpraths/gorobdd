package node

// CleanTags resets tags on the Node graph rooted at n.
// TODO(callpraths): Use a smart Tagger to clean this up without walking
// nodes multiple times.
func CleanTags(n *Node) {
	n.Tag = nil
	switch n.Type {
	case InternalType:
		CleanTags(n.True)
		CleanTags(n.False)
	}
}

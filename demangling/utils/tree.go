package utils

import (
	"strings"
	"swift/demangling"
)

func Tree(node *demangling.Node) string {
	var tree strings.Builder
	recurseNodeTree(&tree, node, 0)
	return strings.TrimSpace(tree.String())
}

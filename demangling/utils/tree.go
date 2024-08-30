package utils

import (
	"github.com/Laky-64/goswift/demangling"
	"strings"
)

func Tree(node *demangling.Node) string {
	var tree strings.Builder
	recurseNodeTree(&tree, node, 0)
	return strings.TrimSpace(tree.String())
}

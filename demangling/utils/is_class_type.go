package utils

import "github.com/Laky-64/goswift/demangling"

func isClassType(node *demangling.Node) bool {
	return node.Kind == demangling.ClassKind
}

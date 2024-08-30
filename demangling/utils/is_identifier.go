package utils

import "swift/demangling"

func IsIdentifier(node *demangling.Node, desired string) bool {
	return node.Kind == demangling.IdentifierKind && node.Text == desired
}

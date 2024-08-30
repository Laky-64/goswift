package utils

import "github.com/Laky-64/goswift/demangling"

func IsIdentifier(node *demangling.Node, desired string) bool {
	return node.Kind == demangling.IdentifierKind && node.Text == desired
}

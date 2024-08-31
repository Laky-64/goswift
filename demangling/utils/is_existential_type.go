package utils

import "github.com/Laky-64/goswift/demangling"

func IsExistentialType(node *demangling.Node) bool {
	switch node.Kind {
	case demangling.ExistentialMetatypeKind,
		demangling.ProtocolListKind,
		demangling.ProtocolListWithClassKind,
		demangling.ProtocolListWithAnyObjectKind:
		return true
	default:
		return false
	}
}

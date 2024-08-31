package utils

import "github.com/Laky-64/goswift/demangling"

func needSpaceBeforeType(t *demangling.Node) bool {
	switch t.Kind {
	case demangling.TypeKind:
		return needSpaceBeforeType(t.FirstChild())
	case demangling.FunctionTypeKind, demangling.NoEscapeFunctionTypeKind, demangling.UncurriedFunctionTypeKind, demangling.CFunctionPointerKind, demangling.DependentGenericTypeKind:
		return false
	default:
		return true
	}
}

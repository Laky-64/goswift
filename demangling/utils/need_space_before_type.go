package utils

import "swift/demangling"

func (ctx *Context) needSpaceBeforeType(t *demangling.Node) bool {
	switch t.Kind {
	case demangling.TypeKind:
		return ctx.needSpaceBeforeType(t.FirstChild())
	case demangling.FunctionTypeKind, demangling.NoEscapeFunctionTypeKind, demangling.UncurriedFunctionTypeKind, demangling.CFunctionPointerKind, demangling.DependentGenericTypeKind:
		return false
	default:
		return true
	}
}

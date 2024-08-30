package utils

import "swift/demangling"

func (ctx *Context) findSugar(node *demangling.Node) sugarType {
	if len(node.Children) == 1 && node.Kind == demangling.TypeKind {
		return ctx.findSugar(node.FirstChild())
	}
	if len(node.Children) != 2 {
		return sugarNone
	}
	if node.Kind != demangling.BoundGenericEnumKind && node.Kind != demangling.BoundGenericStructureKind {
		return sugarNone
	}
	unboundType := node.Children[0].Children[0]
	typeArgs := node.Children[1]
	if node.Kind == demangling.BoundGenericEnumKind {
		if IsIdentifier(unboundType.Children[1], "Optional") &&
			len(typeArgs.Children) == 1 &&
			IsSwiftModule(unboundType.FirstChild()) {
			return sugarOptional
		}
		if IsIdentifier(unboundType.Children[1], "ImplicitlyUnwrappedOptional") &&
			len(typeArgs.Children) == 1 &&
			IsSwiftModule(unboundType.FirstChild()) {
			return sugarImplicitlyUnwrappedOptional
		}
		return sugarNone
	}
	if IsIdentifier(unboundType.Children[1], "Array") &&
		len(typeArgs.Children) == 1 &&
		IsSwiftModule(unboundType.FirstChild()) {
		return sugarArray
	}
	if IsIdentifier(unboundType.Children[1], "Dictionary") &&
		len(typeArgs.Children) == 2 &&
		IsSwiftModule(unboundType.FirstChild()) {
		return sugarDictionary
	}
	return sugarNone
}

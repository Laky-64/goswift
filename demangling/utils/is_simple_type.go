package utils

import "github.com/Laky-64/goswift/demangling"

func IsSimpleType(node *demangling.Node) bool {
	switch node.Kind {
	case demangling.AssociatedTypeKind,
		demangling.AssociatedTypeRefKind,
		demangling.BoundGenericClassKind,
		demangling.BoundGenericEnumKind,
		demangling.BoundGenericStructureKind,
		demangling.BoundGenericProtocolKind,
		demangling.BoundGenericOtherNominalTypeKind,
		demangling.BoundGenericTypeAliasKind,
		demangling.BoundGenericFunctionKind,
		demangling.BuiltinTypeNameKind,
		demangling.BuiltinTupleTypeKind,
		demangling.ClassKind,
		demangling.DependentGenericTypeKind,
		demangling.DependentMemberTypeKind,
		demangling.DependentGenericParamTypeKind,
		demangling.DynamicSelfKind,
		demangling.EnumKind,
		demangling.ErrorTypeKind,
		demangling.ExistentialMetatypeKind,
		demangling.MetatypeKind,
		demangling.MetatypeRepresentationKind,
		demangling.ModuleKind,
		demangling.TupleKind,
		demangling.PackKind,
		demangling.SILPackDirectKind,
		demangling.SILPackIndirectKind,
		demangling.ConstrainedExistentialRequirementListKind,
		demangling.ConstrainedExistentialSelfKind,
		demangling.ProtocolKind,
		demangling.ProtocolSymbolicReferenceKind,
		demangling.ReturnTypeKind,
		demangling.SILBoxTypeKind,
		demangling.SILBoxTypeWithLayoutKind,
		demangling.StructureKind,
		demangling.OtherNominalTypeKind,
		demangling.TupleElementNameKind,
		demangling.TypeAliasKind,
		demangling.TypeListKind,
		demangling.LabelListKind,
		demangling.TypeSymbolicReferenceKind,
		demangling.SugaredOptionalKind,
		demangling.SugaredArrayKind,
		demangling.SugaredDictionaryKind,
		demangling.SugaredParenKind:
		return true
	case demangling.TypeKind:
		return IsSimpleType(node.FirstChild())
	case demangling.ProtocolListKind:
		return len(node.FirstChild().Children) <= 1
	case demangling.ProtocolListWithAnyObjectKind:
		return len(node.FirstChild().FirstChild().Children) == 0
	default:
		return false
	}
}

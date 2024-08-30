package demangling

func isAnyGeneric(kind NodeKind) bool {
	switch kind {
	case
		StructureKind,
		ClassKind,
		EnumKind,
		ProtocolKind,
		ProtocolSymbolicReferenceKind,
		ObjectiveCProtocolSymbolicReferenceKind,
		OtherNominalTypeKind,
		TypeAliasKind,
		TypeSymbolicReferenceKind,
		BuiltinTupleTypeKind:
		return true
	default:
		return false
	}
}

package demangling

// Reference:
// https://github.com/swiftlang/github.com/Laky-64/goswift/blob/main/lib/Demangling/Demangler.cpp#L37C8-L53C2
func isDeclName(kind NodeKind) bool {
	switch kind {
	case IdentifierKind,
		LocalDeclNameKind,
		PrivateDeclNameKind,
		RelatedEntityDeclNameKind,
		PrefixOperatorKind,
		PostfixOperatorKind,
		InfixOperatorKind,
		TypeSymbolicReferenceKind,
		ProtocolSymbolicReferenceKind,
		ObjectiveCProtocolSymbolicReferenceKind:
		return true
	default:
		return false
	}
}

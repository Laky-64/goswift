package demangling

func isMacroExpansionNodeKind(kind NodeKind) bool {
	switch kind {
	case AccessorAttachedMacroExpansionKind,
		MemberAttributeAttachedMacroExpansionKind,
		FreestandingMacroExpansionKind,
		MemberAttachedMacroExpansionKind,
		PeerAttachedMacroExpansionKind,
		ConformanceAttachedMacroExpansionKind,
		ExtensionAttachedMacroExpansionKind,
		MacroExpansionLocKind:
		return true
	default:
		return false
	}
}

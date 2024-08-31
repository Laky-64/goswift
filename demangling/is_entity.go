package demangling

func isEntity(kind NodeKind) bool {
	if kind == TypeKind {
		return true
	}
	return isContext(kind)
}

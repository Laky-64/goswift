package demangling

func (ctx *Context) extendedExistentialShape(specialChar rune) (*Node, error) {
	node := ctx.popNodeKind(TypeKind)
	var genSig *Node
	if specialChar == 'G' {
		genSig = ctx.popNodeKind(DependentGenericSignatureKind)
	}
	if genSig != nil {
		return createWithChildren(ExtendedExistentialTypeShapeKind, genSig, node), nil
	}
	return createWithChildren(ExtendedExistentialTypeShapeKind, node), nil
}

package demangling

func (ctx *Context) associatedTypeSimple(base *Node) *Node {
	atName := ctx.popAssocTypeName()
	var baseTy *Node
	if base != nil {
		baseTy = CreateType(base)
	} else {
		baseTy = ctx.popNodeKind(TypeKind)
	}
	return CreateType(createWithChildren(DependentMemberTypeKind, baseTy, atName))
}

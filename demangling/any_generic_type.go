package demangling

func (ctx *Context) anyGenericType(kind NodeKind) (*Node, error) {
	name := ctx.popNodePred(isDeclName)
	context := ctx.popContext()
	nType := CreateType(createWithChildren(kind, context, name))
	ctx.addSubstitution(nType)
	return nType, nil
}

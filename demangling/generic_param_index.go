package demangling

func (ctx *Context) genericParamIndex() *Node {
	if ctx.nextIf('d') {
		depth := ctx.index()
		index := ctx.index()
		return ctx.getDependentGenericParamType(depth, index)
	}
	if ctx.nextIf('z') {
		return ctx.getDependentGenericParamType(0, 0)
	}
	if ctx.nextIf('s') {
		return createNode(ConstrainedExistentialSelfKind)
	}
	return ctx.getDependentGenericParamType(0, ctx.index()+1)
}

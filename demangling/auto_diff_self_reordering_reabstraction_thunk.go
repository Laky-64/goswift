package demangling

func (ctx *Context) autoDiffSelfReorderingReabstractionThunk() (*Node, error) {
	result := createNode(AutoDiffSelfReorderingReabstractionThunkKind)
	addChild(result, ctx.popNodeKind(DependentGenericSignatureKind))
	result = addChild(result, ctx.popNodeKind(TypeKind))
	result = addChild(result, ctx.popNodeKind(TypeKind))
	if result != nil {
		result.reverseChildren(0)
	}
	result = addChild(result, ctx.autoDiffFunctionKind())
	return result, nil
}

package demangling

func (ctx *Context) plainFunction() (*Node, error) {
	genSig := ctx.popNodeKind(DependentGenericSignatureKind)
	typ := ctx.popFunctionType(FunctionTypeKind, false)
	labelList := ctx.popFunctionParamLabels(typ)
	if genSig != nil {
		typ = CreateType(createWithChildren(DependentGenericTypeKind, genSig, typ))
	}
	name := ctx.popNodePred(isDeclName)
	context := ctx.popContext()
	var result *Node
	if labelList != nil {
		result = createWithChildren(FunctionKind, context, name, labelList, typ)
	} else {
		result = createWithChildren(FunctionKind, context, name, typ)
	}
	result = ctx.setParentForOpaqueReturnTypeNodes(result, typ)
	return result, nil
}

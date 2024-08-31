package demangling

func (ctx *Context) entity(kind NodeKind) *Node {
	typ := ctx.popNodeKind(TypeKind)
	labelList := ctx.popFunctionParamLabels(typ)
	name := ctx.popNodePred(isDeclName)
	context := ctx.popContext()
	var result *Node
	if labelList != nil {
		result = createWithChildren(kind, context, name, labelList, typ)
	} else {
		result = createWithChildren(kind, context, name, typ)
	}
	return ctx.setParentForOpaqueReturnTypeNodes(result, typ)
}

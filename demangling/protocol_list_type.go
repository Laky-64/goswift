package demangling

func (ctx *Context) protocolListType() *Node {
	return CreateType(ctx.protocolList())
}

package demangling

func (ctx *Context) variable() *Node {
	return ctx.accessor(ctx.entity(VariableKind))
}

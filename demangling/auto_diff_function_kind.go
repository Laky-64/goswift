package demangling

func (ctx *Context) autoDiffFunctionKind() *Node {
	kind := ctx.nextChar()
	if kind != 'f' && kind != 'r' && kind != 'd' && kind != 'p' {
		return nil
	}
	return createNodeWithIndex(AutoDiffFunctionKind, kind)
}

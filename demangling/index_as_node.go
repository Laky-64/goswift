package demangling

func (ctx *Context) indexAsNode() *Node {
	idx := ctx.index()
	if idx >= 0 {
		return createNodeWithIndex(NumberKind, idx)
	}
	return nil
}

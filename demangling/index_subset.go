package demangling

func (ctx *Context) indexSubset() *Node {
	var str string
	for c := ctx.peekChar(); c == 'S' || c == 'U'; c = ctx.peekChar() {
		str += string(c)
		ctx.nextChar()
	}
	if len(str) == 0 {
		return nil
	}
	return createNodeWithText(IndexSubsetKind, str)
}

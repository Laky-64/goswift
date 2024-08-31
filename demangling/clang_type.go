package demangling

func (ctx *Context) clangType() *Node {
	numChars := ctx.natural()
	if numChars <= 0 || ctx.Pos+numChars > len(ctx.Data) {
		return nil
	}
	mangledClangType := string(ctx.Data[ctx.Pos : ctx.Pos+numChars])
	ctx.Pos += numChars
	return createNodeWithText(ClangTypeKind, mangledClangType)
}

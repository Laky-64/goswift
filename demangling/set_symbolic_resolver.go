package demangling

func (ctx *Context) SetSymbolicReferenceResolver(f func(SymbolicReferenceKind, Directness, int32) (*Node, error)) {
	ctx.symbolicReferenceResolver = f
}

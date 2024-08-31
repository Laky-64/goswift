package demangling

func (ctx *Context) extensionContext() (*Node, error) {
	genSig := ctx.popNodeKind(DependentGenericSignatureKind)
	module := ctx.popModule()
	typ := ctx.popTypeAndGetAnyGeneric()
	ext := createWithChildren(ExtensionKind, module, typ)
	if genSig != nil {
		ext = addChild(ext, genSig)
	}
	return ext, nil
}

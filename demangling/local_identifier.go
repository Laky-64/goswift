package demangling

func (ctx *Context) localIdentifier() (*Node, error) {
	if ctx.nextIf('L') {
		discriminator := ctx.popNodeKind(IdentifierKind)
		name := ctx.popNodePred(isDeclName)
		return createWithChildren(PrivateDeclNameKind, discriminator, name), nil
	}
	if ctx.nextIf('l') {
		discriminator := ctx.popNodeKind(IdentifierKind)
		return createWithChildren(PrivateDeclNameKind, discriminator), nil
	}
	if ctx.peekChar() >= 'a' && ctx.peekChar() <= 'j' || ctx.peekChar() >= 'A' && ctx.peekChar() <= 'J' {
		relatedEntityKind := ctx.nextChar()
		kindNd := createNodeWithText(IdentifierKind, string(relatedEntityKind))
		name := ctx.popNode()
		result := createNode(RelatedEntityDeclNameKind)
		addChild(result, kindNd)
		return addChild(result, name), nil
	}
	discriminator := ctx.indexAsNode()
	name := ctx.popNodePred(isDeclName)
	return createWithChildren(LocalDeclNameKind, discriminator, name), nil
}

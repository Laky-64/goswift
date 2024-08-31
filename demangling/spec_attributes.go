package demangling

func (ctx *Context) specAttributes(specKind NodeKind) *Node {
	isSerialized := ctx.nextIf('q')
	asyncRemoved := ctx.nextIf('a')

	passID := ctx.nextChar() - '0'
	if passID < 0 || passID >= maxSpecializationPass {
		return nil
	}
	specNd := createNode(specKind)

	if isSerialized {
		specNd.addChild(createNode(IsSerializedKind))
	}
	if asyncRemoved {
		specNd.addChild(createNode(AsyncRemovedKind))
	}
	specNd.addChild(createNodeWithIndex(SpecializationPassIDKind, passID))
	return specNd
}

package demangling

func (ctx *Context) genericSignature(hasParamCounts bool) *Node {
	sig := createNode(DependentGenericSignatureKind)
	if hasParamCounts {
		for !ctx.nextIf('l') {
			count := rune(0)
			if !ctx.nextIf('z') {
				count = ctx.index() + 1
			}
			if count < 0 {
				return nil
			}
			sig.addChild(createNodeWithIndex(DependentGenericParamCountKind, count))
		}
	} else {
		sig.addChild(createNodeWithIndex(DependentGenericParamCountKind, 1))
	}
	numCounts := len(sig.Children)
	for req := ctx.popNodePred(isRequirement); req != nil; req = ctx.popNodePred(isRequirement) {
		sig.addChild(req)
	}
	sig.reverseChildren(numCounts)
	return sig
}

package demangling

func (ctx *Context) getDependentGenericParamType(depth, index rune) *Node {
	if depth < 0 || index < 0 {
		return nil
	}
	paramTy := createNode(DependentGenericParamTypeKind)
	paramTy.addChild(createNodeWithIndex(IndexKind, depth))
	paramTy.addChild(createNodeWithIndex(IndexKind, index))
	return paramTy
}

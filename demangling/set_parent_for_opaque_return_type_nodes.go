package demangling

func (ctx *Context) setParentForOpaqueReturnTypeNodes(parent *Node, visitedNode *Node) *Node {
	if parent == nil || visitedNode == nil {
		return nil
	}
	if visitedNode.Kind == OpaqueReturnTypeKind {
		if visitedNode.Children != nil && visitedNode.LastChild().Kind == OpaqueReturnTypeParentKind {
			return parent
		}
		visitedNode.addChild(createNodeWithIndex(OpaqueReturnTypeParentKind, parent.Index))
		return parent
	}

	if visitedNode.Kind == FunctionKind || visitedNode.Kind == VariableKind || visitedNode.Kind == SubscriptKind {
		return parent
	}

	for _, child := range visitedNode.Children {
		ctx.setParentForOpaqueReturnTypeNodes(parent, child)
	}
	return parent
}

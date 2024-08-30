package demangling

func (ctx *Context) boundGenerics(typeListList *[]*Node, retroActiveConformance *Node) bool {
	retroActiveConformance = ctx.popRetroactiveConformance()
	for {
		typeList := createNode(TypeListKind)
		*typeListList = append(*typeListList, typeList)
		for {
			ty := ctx.popNodeKind(TypeKind)
			if ty == nil {
				break
			}
			typeList.addChild(ty)
		}
		typeList.reverseChildren(0)
		if ctx.popNodeKind(EmptyListKind) != nil {
			break
		}
		if ctx.popNodeKind(FirstElementMarkerKind) == nil {
			return false
		}
	}
	return true
}

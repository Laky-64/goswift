package demangling

func (ctx *Context) protocolList() *Node {
	typeList := createNode(TypeListKind)
	protoList := createWithChildren(ProtocolListKind, typeList)
	if ctx.popNodeKind(EmptyListKind) == nil {
		var firstElem bool
		for {
			firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
			proto := ctx.popProtocol()
			if proto == nil {
				return nil
			}
			typeList.addChild(proto)
			if firstElem {
				break
			}
		}
		typeList.reverseChildren(0)
	}
	return protoList
}

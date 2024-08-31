package demangling

func isProtocolNode(node *Node) bool {
	if node == nil {
		return false
	}
	switch node.Kind {
	case TypeKind:
		return isProtocolNode(node.Children[0])
	case ProtocolKind, ProtocolSymbolicReferenceKind, ObjectiveCProtocolSymbolicReferenceKind:
		return true
	default:
		return false
	}
}

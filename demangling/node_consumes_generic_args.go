package demangling

func nodeConsumesGenericArgs(node *Node) bool {
	switch node.Kind {
	case
		VariableKind,
		SubscriptKind,
		ImplicitClosureKind,
		ExplicitClosureKind,
		DefaultArgumentInitializerKind,
		InitializerKind,
		PropertyWrapperBackingInitializerKind,
		PropertyWrapperInitFromProjectedValueKind,
		StaticKind:
		return false
	default:
		return true
	}
}

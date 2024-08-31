package demangling

func (ctx *Context) addSubstitution(node *Node) {
	if node != nil {
		ctx.substitutions = append(ctx.substitutions, node)
	}
}

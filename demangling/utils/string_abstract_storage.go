package utils

import "github.com/Laky-64/goswift/demangling"

func (ctx *Context) abstractStorage(node *demangling.Node, depth int, asPrefixContent bool, extraName string) *demangling.Node {
	switch node.Kind {
	case demangling.VariableKind:
		return ctx.stringEntity(node, depth, asPrefixContent, withColonType, true, extraName, -1, "")
	case demangling.SubscriptKind:
		return ctx.stringEntity(node, depth, asPrefixContent, withColonType, false, extraName, -1, "subscript")
	default:
		ctx.WriteString("<<unknown abstract storage kind>>")
		return nil
	}
}

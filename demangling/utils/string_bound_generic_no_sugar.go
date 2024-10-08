package utils

import "github.com/Laky-64/goswift/demangling"

func (ctx *Context) stringBoundGenericNoSugar(node *demangling.Node, depth int) {
	if len(node.Children) < 2 {
		return
	}
	typeList := node.Children[1]
	ctx.stringNode(node.FirstChild(), depth+1, false)
	ctx.WriteByte('<')
	ctx.printChildren(typeList, depth, ", ")
	ctx.WriteByte('>')
}

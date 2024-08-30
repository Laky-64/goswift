package utils

import (
	"github.com/Laky-64/swift/demangling"
)

func (ctx *Context) stringWithParens(node *demangling.Node, depth int) {
	needPerens := !IsSimpleType(node)
	if needPerens {
		ctx.WriteString("(")
	}
	ctx.stringNode(node, depth+1, false)
	if needPerens {
		ctx.WriteString(")")
	}
}

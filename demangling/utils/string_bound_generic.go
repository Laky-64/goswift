package utils

import (
	"github.com/Laky-64/swift/demangling"
)

func (ctx *Context) stringBoundGeneric(node *demangling.Node, depth int) {
	if len(node.Children) < 2 {
		return
	}
	if len(node.Children) != 2 {
		ctx.stringBoundGenericNoSugar(node, depth)
		return
	}
	if !ctx.SynthesizeSugarOnTypes || node.Kind == demangling.BoundGenericClassKind {
		ctx.stringBoundGenericNoSugar(node, depth)
		return
	}
	if node.Kind == demangling.BoundGenericProtocolKind {
		ctx.printChildren(node.Children[1], depth, ", ")
		ctx.WriteString(" as ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return
	}
	sugar := ctx.findSugar(node)
	switch sugar {
	case sugarNone:
		ctx.stringBoundGenericNoSugar(node, depth)
	case sugarOptional, sugarImplicitlyUnwrappedOptional:
		ctx.stringWithParens(node.Children[1].Children[0], depth)
		if sugar == sugarOptional {
			ctx.WriteString("?")
		} else {
			ctx.WriteString("!")
		}
	case sugarArray:
		ctx.WriteString("[")
		ctx.stringNode(node.Children[1].Children[0], depth+1, false)
		ctx.WriteString("]")
	case sugarDictionary:
		ctx.WriteString("[")
		ctx.stringNode(node.Children[1].Children[0], depth+1, false)
		ctx.WriteString(" : ")
		ctx.stringNode(node.Children[1].Children[1], depth+1, false)
		ctx.WriteString("]")
	}
}

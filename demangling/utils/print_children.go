package utils

import "github.com/Laky-64/swift/demangling"

func (ctx *Context) printChildren(node *demangling.Node, depth int, separator string) {
	if node == nil {
		return
	}
	for i, child := range node.Children {
		if i > 0 {
			ctx.WriteString(separator)
		}
		ctx.stringNode(child, depth+1, false)
	}
}

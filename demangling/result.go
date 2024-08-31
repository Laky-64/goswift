package demangling

import (
	"fmt"
)

func (ctx *Context) Result() (*Node, error) {
	for ctx.Pos < ctx.Size {
		node, err := ctx.operator()
		if err != nil {
			return nil, err
		} else if node == nil {
			return nil, fmt.Errorf("missing %c node, processed text: %s", ctx.Data[ctx.Pos-1], ctx.Data[:ctx.Pos-1])
		}
		ctx.pushNode(node)
	}
	parent := createNode(GlobalKind)
	for {
		funcAttr := ctx.popNodePred(isFunctionAttr)
		if funcAttr == nil {
			break
		}
		parent.addChild(funcAttr)
		if funcAttr.Kind == PartialApplyForwarderKind || funcAttr.Kind == PartialApplyObjCForwarderKind {
			parent = funcAttr
		}
	}
	for _, node := range ctx.nodeList {
		switch node.Kind {
		case TypeKind:
			parent.addChild(node.FirstChild())
		default:
			parent.addChild(node)
		}
	}
	return parent, nil
}

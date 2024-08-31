package demangling

import "fmt"

func (ctx *Context) subscript() (*Node, error) {
	privateName := ctx.popNodeKind(PrivateDeclNameKind)
	typ := ctx.popNodeKind(TypeKind)
	labelList := ctx.popFunctionParamLabels(typ)
	context := ctx.popContext()

	if typ == nil {
		return nil, fmt.Errorf("subscript: typ is nil")
	}

	subscript := createNode(SubscriptKind)
	subscript = addChild(subscript, context)
	addChild(subscript, labelList)
	subscript = addChild(subscript, typ)
	addChild(subscript, privateName)

	subscript = ctx.setParentForOpaqueReturnTypeNodes(subscript, typ)

	return ctx.accessor(subscript), nil
}

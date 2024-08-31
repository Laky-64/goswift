package demangling

import "fmt"

func (ctx *Context) autoDiffFunctionOrSimpleThunk(nodeKind NodeKind) (*Node, error) {
	result := createNode(nodeKind)
	for {
		originalNode := ctx.popNode()
		if originalNode == nil {
			break
		}
		result = addChild(result, originalNode)
	}
	result.reverseChildren(0)
	kind := ctx.autoDiffFunctionKind()
	result = addChild(result, kind)
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('p') {
		return nil, fmt.Errorf("autoDiffFunctionOrSimpleThunk: expected 'p'")
	}
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('r') {
		return nil, fmt.Errorf("autoDiffFunctionOrSimpleThunk: expected 'r'")
	}
	return result, nil
}

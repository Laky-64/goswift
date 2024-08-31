package demangling

import "fmt"

func (ctx *Context) autoDiffSubsetParametersThunk() (*Node, error) {
	result := createNode(AutoDiffSubsetParametersThunkKind)
	for {
		node := ctx.popNode()
		if node == nil {
			break
		}
		result = addChild(result, node)
	}
	result.reverseChildren(0)
	kind := ctx.autoDiffFunctionKind()
	result = addChild(result, kind)
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('p') {
		return nil, fmt.Errorf("autoDiffSubsetParametersThunk: expected 'p'")
	}
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('r') {
		return nil, fmt.Errorf("autoDiffSubsetParametersThunk: expected 'r'")
	}
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('P') {
		return nil, fmt.Errorf("autoDiffSubsetParametersThunk: expected 'P'")
	}
	return result, nil
}

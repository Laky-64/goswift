package demangling

import "fmt"

func (ctx *Context) demangleLifetimeDependence() (*Node, error) {
	kind := getMangledLifetimeDependenceKind(ctx.nextChar())
	if kind == UnknownLifetime {
		return nil, fmt.Errorf("unexpected lifetime dependence kind: %c", ctx.peekChar())
	}
	result := createNode(LifetimeDependenceKind)
	result = addChild(result, createNodeWithIndex(IndexKind, rune(kind)))
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('_') {
		return nil, fmt.Errorf("expected '_' after lifetime dependence")
	}
	return result, nil
}

package demangling

import "fmt"

func (ctx *Context) boundGenericType() (*Node, error) {
	var retroactiveConformance *Node
	var typeListList []*Node
	if !ctx.boundGenerics(&typeListList, retroactiveConformance) {
		return nil, fmt.Errorf("failed to parse bound generics")
	}
	nominal := ctx.popTypeAndGetAnyGeneric()
	if nominal == nil {
		return nil, fmt.Errorf("expected nominal type")
	}
	boundNode, err := ctx.boundGenericArgs(nominal, typeListList, 0)
	if err != nil {
		return nil, err
	}
	addChild(retroactiveConformance, boundNode)
	nType := CreateType(boundNode)
	ctx.addSubstitution(nType)
	return nType, nil
}

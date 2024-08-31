package demangling

import "fmt"

func (ctx *Context) archetype() (*Node, error) {
	switch ctx.nextChar() {
	case 'a':
		ident := ctx.popNodeKind(IdentifierKind)
		archeTy := ctx.popTypeAndGetChild()
		assocTy := CreateType(createWithChildren(AssociatedTypeRefKind, archeTy, ident))
		ctx.addSubstitution(assocTy)
		return assocTy, nil
	case 'O':
		definingContext := ctx.popContext()
		return createWithChildren(OpaqueReturnTypeOfKind, definingContext), nil
	case 'o':
		index := ctx.index()
		var boundGenericArgs []*Node
		var retroactiveConformances *Node
		if !ctx.boundGenerics(&boundGenericArgs, retroactiveConformances) {
			return nil, nil
		}
		name := ctx.popNode()
		if name == nil {
			return nil, fmt.Errorf("name is nil")
		}
		opaque := createWithChildren(OpaqueTypeKind, name, createNodeWithIndex(IndexKind, index))
		boundGenerics := createNode(TypeListKind)
		for i := len(boundGenericArgs) - 1; i >= 0; i-- {
			boundGenerics.addChild(boundGenericArgs[i])
		}
		opaque.addChild(boundGenerics)
		if retroactiveConformances != nil {
			opaque.addChild(retroactiveConformances)
		}
		opaqueTy := CreateType(opaque)
		ctx.addSubstitution(opaqueTy)
		return opaqueTy, nil
	case 'r':
		return CreateType(createNode(OpaqueReturnTypeKind)), nil
	case 'R':
		ordinal := ctx.index()
		if ordinal < 0 {
			return nil, fmt.Errorf("ordinal is less than 0")
		}
		return CreateType(createWithChildren(OpaqueReturnTypeKind, createNodeWithIndex(OpaqueReturnTypeIndexKind, ordinal))), nil
	case 'x':
		t := ctx.associatedTypeSimple(nil)
		ctx.addSubstitution(t)
		return t, nil
	case 'X':
		t, err := ctx.associatedTypeCompound(nil)
		if err != nil {
			return nil, err
		}
		ctx.addSubstitution(t)
		return t, nil
	case 'y':
		t := ctx.associatedTypeSimple(ctx.genericParamIndex())
		ctx.addSubstitution(t)
		return t, nil
	case 'Y':
		t, err := ctx.associatedTypeCompound(ctx.genericParamIndex())
		if err != nil {
			return nil, err
		}
		ctx.addSubstitution(t)
		return t, nil
	case 'z':
		t := ctx.associatedTypeSimple(ctx.getDependentGenericParamType(0, 0))
		ctx.addSubstitution(t)
		return t, nil
	case 'Z':
		t, err := ctx.associatedTypeCompound(ctx.getDependentGenericParamType(0, 0))
		if err != nil {
			return nil, err
		}
		ctx.addSubstitution(t)
		return t, nil
	case 'p':
		countTy := ctx.popTypeAndGetChild()
		patternTy := ctx.popTypeAndGetChild()
		packExpansionTy := CreateType(createWithChildren(PackExpansionKind, patternTy, countTy))
		return packExpansionTy, nil
	case 'e':
		packTy := ctx.popTypeAndGetChild()
		level := ctx.index()
		if level < 0 {
			return nil, fmt.Errorf("level is less than 0")
		}
		packElementTy := CreateType(createWithChildren(PackElementKind, packTy, createNodeWithIndex(PackElementLevelKind, level)))
		return packElementTy, nil
	case 'P':
		return ctx.popPack(), nil
	case 'S':
		return ctx.popSILPack()
	default:
		return nil, fmt.Errorf("unknown archetype: %c", ctx.Data[ctx.Pos-1])
	}
}

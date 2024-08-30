package demangling

import (
	"fmt"
)

func (ctx *Context) boundGenericArgs(nominal *Node, typeLists []*Node, typeListIdx int) (*Node, error) {
	if nominal == nil {
		return nil, fmt.Errorf("expected nominal type")
	}
	if typeListIdx >= len(typeLists) {
		return nil, fmt.Errorf("type list index out of bounds")
	}
	if nominal.Kind == TypeSymbolicReferenceKind || nominal.Kind == ProtocolSymbolicReferenceKind {
		remainingTypeList := createNode(TypeListKind)
		for i := len(typeLists) - 1; i >= typeListIdx && i < len(typeLists); i-- {
			list := typeLists[i]
			for _, child := range list.Children {
				remainingTypeList.addChild(child)
			}
		}
		return createWithChildren(BoundGenericOtherNominalTypeKind, CreateType(nominal), remainingTypeList), nil
	}
	if len(nominal.Children) == 0 {
		return nil, fmt.Errorf("expected nominal type to have children")
	}
	context := nominal.FirstChild()
	consumesGenericArgs := nodeConsumesGenericArgs(context)
	args := typeLists[typeListIdx]
	if consumesGenericArgs {
		typeListIdx++
	}
	if typeListIdx < len(typeLists) {
		var boundParent *Node
		if context.Kind == ExtensionKind {
			b, err := ctx.boundGenericArgs(context.Children[1], typeLists, typeListIdx)
			if err != nil {
				return nil, err
			}
			boundParent = b
			boundParent = createWithChildren(
				ExtensionKind,
				CreateType(context.Children[1]),
				boundParent,
			)
			if len(context.Children) == 3 {
				addChild(boundParent, context.Children[2])
			}
		} else {
			b, err := ctx.boundGenericArgs(context, typeLists, typeListIdx)
			if err != nil {
				return nil, err
			}
			boundParent = b
		}
		newNominal := createWithChildren(
			nominal.Kind,
			boundParent,
		)
		for i := 1; i < len(nominal.Children); i++ {
			addChild(newNominal, nominal.Children[i])
		}
		nominal = newNominal
	}
	if !consumesGenericArgs {
		return nominal, nil
	}
	if len(args.Children) == 0 {
		return nominal, nil
	}
	var nKind NodeKind
	switch nominal.Kind {
	case ClassKind:
		nKind = BoundGenericClassKind
	case StructureKind:
		nKind = BoundGenericStructureKind
	case EnumKind:
		nKind = BoundGenericEnumKind
	case ProtocolKind:
		nKind = BoundGenericProtocolKind
	case OtherNominalTypeKind:
		nKind = BoundGenericOtherNominalTypeKind
	case TypeAliasKind:
		nKind = BoundGenericTypeAliasKind
	case FunctionKind, ConstructorKind:
		return createWithChildren(BoundGenericFunctionKind, nominal, args), nil
	default:
		return nil, fmt.Errorf("unexpected nominal type kind")
	}
	return createWithChildren(nKind, CreateType(nominal), args), nil
}

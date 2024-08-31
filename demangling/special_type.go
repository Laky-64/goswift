package demangling

import "fmt"

func (ctx *Context) specialType() (*Node, error) {
	switch specialChar := ctx.nextChar(); specialChar {
	case 'E':
		return ctx.popFunctionType(NoEscapeFunctionTypeKind, false), nil
	case 'A':
		return ctx.popFunctionType(EscapingAutoClosureTypeKind, false), nil
	case 'f':
		return ctx.popFunctionType(ThinFunctionTypeKind, false), nil
	case 'K':
		return ctx.popFunctionType(AutoClosureTypeKind, false), nil
	case 'U':
		return ctx.popFunctionType(UncurriedFunctionTypeKind, false), nil
	case 'L':
		return ctx.popFunctionType(EscapingObjCBlockKind, false), nil
	case 'B':
		return ctx.popFunctionType(ObjCBlockKind, false), nil
	case 'C':
		return ctx.popFunctionType(CFunctionPointerKind, false), nil
	case 'g', 'G':
		return ctx.extendedExistentialShape(specialChar)
	case 'j':
		return ctx.symbolicExtendedExistentialType()
	case 'z':
		switch ctx.nextChar() {
		case 'B':
			return ctx.popFunctionType(ObjCBlockKind, true), nil
		case 'C':
			return ctx.popFunctionType(CFunctionPointerKind, true), nil
		default:
			return nil, fmt.Errorf("specialType: unexpected z %c", ctx.Data[ctx.Pos-1])
		}
	case 'o':
		return CreateType(createWithChildren(UnownedKind, ctx.popNodeKind(TypeKind))), nil
	case 'u':
		return CreateType(createWithChildren(UnmanagedKind, ctx.popNodeKind(TypeKind))), nil
	case 'w':
		return CreateType(createWithChildren(WeakKind, ctx.popNodeKind(TypeKind))), nil
	case 'b':
		return CreateType(createWithChildren(SILBoxTypeKind, ctx.popNodeKind(TypeKind))), nil
	case 'D':
		return CreateType(createWithChildren(DynamicSelfKind, ctx.popNodeKind(TypeKind))), nil
	case 'M':
		mtr, err := ctx.metatypeRepresentation()
		if err != nil {
			return nil, err
		}
		return CreateType(createWithChildren(MetatypeKind, mtr, ctx.popNodeKind(TypeKind))), nil
	case 'm':
		mtr, err := ctx.metatypeRepresentation()
		if err != nil {
			return nil, err
		}
		return CreateType(createWithChildren(ExistentialMetatypeKind, mtr, ctx.popNodeKind(TypeKind))), nil
	case 'P':
		reqs, err := ctx.constrainedExistentialRequirementList()
		if err != nil {
			return nil, err
		}
		base := ctx.popNodeKind(TypeKind)
		return CreateType(createWithChildren(ConstrainedExistentialKind, base, reqs)), nil
	case 'p':
		return CreateType(createWithChildren(ExistentialMetatypeKind, ctx.popNodeKind(TypeKind))), nil
	case 'c':
		superClass := ctx.popNodeKind(TypeKind)
		protocols := ctx.protocolList()
		return CreateType(createWithChildren(ProtocolListWithClassKind, protocols, superClass)), nil
	case 'l':
		protocols := ctx.protocolList()
		return CreateType(createWithChildren(ProtocolListWithAnyObjectKind, protocols)), nil
	case 'X', 'x':
		var signature, genericArgs *Node
		if specialChar == 'X' {
			signature = ctx.popNodeKind(DependentGenericSignatureKind)
			if signature == nil {
				return nil, fmt.Errorf("specialType: signature is nil")
			}
			genericArgs = ctx.popTypeList()
			if genericArgs == nil {
				return nil, fmt.Errorf("specialType: genericArgs is nil")
			}
		}

		fieldTypes := ctx.popTypeList()
		if fieldTypes == nil {
			return nil, fmt.Errorf("specialType: fieldTypes is nil")
		}
		layout := createNode(SILBoxLayoutKind)
		for i := 0; i < len(fieldTypes.Children); i++ {
			fieldType := fieldTypes.Children[i]
			var isMutable bool
			if fieldType.Children[0].Kind == InOutKind {
				isMutable = true
				fieldType = CreateType(fieldType.FirstChild().FirstChild())
			}
			var field *Node
			if isMutable {
				field = createNode(SILBoxMutableFieldKind)
			} else {
				field = createNode(SILBoxImmutableFieldKind)
			}
			field.addChild(fieldType)
			layout.addChild(field)
		}
		boxTy := createNode(SILBoxTypeWithLayoutKind)
		boxTy.addChild(layout)
		if signature != nil {
			boxTy.addChild(signature)
			boxTy.addChild(genericArgs)
		}
		return CreateType(boxTy), nil
	case 'Y':
		return ctx.anyGenericType(OtherNominalTypeKind)
	case 'Z':
		types := ctx.popTypeList()
		name := ctx.popNodeKind(IdentifierKind)
		parent := ctx.popContext()
		anon := createNode(AnonymousContextKind)
		anon = addChild(anon, name)
		anon = addChild(anon, parent)
		anon = addChild(anon, types)
		return anon, nil
	case 'e':
		return CreateType(createNode(ErrorTypeKind)), nil
	case 'S':
		switch ctx.nextChar() {
		case 'q':
			return CreateType(createWithChildren(SugaredOptionalKind, ctx.popNodeKind(TypeKind))), nil
		case 'a':
			return CreateType(createWithChildren(SugaredArrayKind, ctx.popNodeKind(TypeKind))), nil
		case 'D':
			value := ctx.popNodeKind(TypeKind)
			key := ctx.popNodeKind(TypeKind)
			return CreateType(createWithChildren(SugaredDictionaryKind, key, value)), nil
		case 'p':
			return CreateType(createWithChildren(SugaredParenKind, ctx.popNodeKind(TypeKind))), nil
		default:
			return nil, fmt.Errorf("specialType: unexpected S %c", ctx.Data[ctx.Pos-1])
		}
	default:
		return nil, fmt.Errorf("specialType: unexpected o %c", ctx.Data[ctx.Pos-1])
	}
}

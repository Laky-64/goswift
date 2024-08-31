package demangling

// Reference:
// https://github.com/swiftlang/swift/blob/main/lib/Demangling/Demangler.cpp#L1002
func (ctx *Context) operator() (*Node, error) {
recur:
	switch ctx.nextChar() {
	case 0xFF:
		goto recur
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xA, 0xB, 0xC:
		return ctx.symbolicReference()
	case 'A':
		return ctx.multiSubstitutions()
	case 'C':
		return ctx.anyGenericType(ClassKind)
	case 'E':
		return ctx.extensionContext()
	case 'F':
		return ctx.plainFunction()
	case 'G':
		return ctx.boundGenericType()
	case 'K':
		return createNode(ThrowsAnnotationKind), nil
	case 'L':
		return ctx.localIdentifier()
	case 'M':
		return ctx.metaType()
	case 'N':
		return createWithChildren(TypeMetadataKind, ctx.popNode()), nil
	case 'O':
		return ctx.anyGenericType(EnumKind)
	case 'P':
		return ctx.anyGenericType(ProtocolKind)
	case 'Q':
		return ctx.archetype()
	case 'R':
		return ctx.genericRequirement()
	case 'S':
		return ctx.standardSubstitution()
	case 'T':
		return ctx.thunkOrSpecialization()
	case 'V':
		return ctx.anyGenericType(StructureKind)
	case 'W':
		return ctx.witness()
	case 'X':
		return ctx.specialType()
	case 'Y':
		return ctx.typeAnnotation()
	case 'Z':
		return createWithChildren(StaticKind, ctx.popNodePred(isEntity)), nil
	case 'a':
		return ctx.anyGenericType(TypeAliasKind)
	case 'c':
		return ctx.popFunctionType(FunctionTypeKind, false), nil
	case 'd':
		return createNode(VariadicMarkerKind), nil
	case 'f':
		return ctx.functionEntity()
	case 'i':
		return ctx.subscript()
	case 'l':
		return ctx.genericSignature(false), nil
	case 'm':
		return CreateType(createWithChildren(MetatypeKind, ctx.popNodeKind(TypeKind))), nil
	case 'o':
		return ctx.operatorIdentifier()
	case 'p':
		return ctx.protocolListType(), nil
	case 'q':
		return CreateType(ctx.genericParamIndex()), nil
	case 'r':
		return ctx.genericSignature(true), nil
	case 's':
		return createNodeWithText(ModuleKind, StdLibName), nil
	case 't':
		return ctx.popTuple()
	case 'v':
		return ctx.variable(), nil
	case 'y':
		return createNode(EmptyListKind), nil
	case 'x':
		return CreateType(ctx.getDependentGenericParamType(0, 0)), nil
	case 'z':
		return CreateType(createWithChildren(InOutKind, ctx.popTypeAndGetChild())), nil
	case '_':
		return createNode(FirstElementMarkerKind), nil
	default:
		ctx.pushBack()
		return ctx.identifier()
	}
}

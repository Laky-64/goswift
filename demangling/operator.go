package demangling

// Reference:
// https://github.com/swiftlang/github.com/Laky-64/goswift/blob/main/lib/Demangling/Demangler.cpp#L1002
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
	case 'F':
		return ctx.plainFunction()
	case 'G':
		return ctx.boundGenericType()
	case 'M':
		return ctx.metaType()
	case 'S':
		return ctx.standardSubstitution()
	case 'V':
		return ctx.anyGenericType(StructureKind)
	case 'Y':
		return ctx.typeAnnotation()
	case 's':
		return createNodeWithText(ModuleKind, StdLibName), nil
	case 't':
		return ctx.popTuple()
	case 'y':
		return createNode(EmptyListKind), nil
	case '_':
		return createNode(FirstElementMarkerKind), nil
	default:
		ctx.pushBack()
		return ctx.identifier()
	}
}

package demangling

func (ctx *Context) funcSpecParam(kind NodeKind) *Node {
	param := createNode(kind)
	switch ctx.nextChar() {
	case 'n':
		return param
	case 'c':
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, ClosureProp))
	case 'p':
		switch ctx.nextChar() {
		case 'f':
			return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, ConstantPropFunction))
		case 'g':
			return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, ConstantPropGlobal))
		case 'i':
			return ctx.addFuncSpecParamNumber(param, ConstantPropInteger)
		case 'd':
			return ctx.addFuncSpecParamNumber(param, ConstantPropFloat)
		case 's':
			var encoding string
			switch ctx.nextChar() {
			case 'b':
				encoding = "u8"
			case 'w':
				encoding = "u16"
			case 'c':
				encoding = "objc"
			default:
				return nil
			}
			addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, ConstantPropString))
			return addChild(param, createNodeWithText(FunctionSignatureSpecializationParamPayloadKind, encoding))
		case 'k':
			return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, ConstantPropKeyPath))
		default:
			return nil
		}
	case 'e':
		var value rune
		if ctx.nextIf('D') {
			value |= Dead
		}
		if ctx.nextIf('G') {
			value |= OwnedToGuaranteed
		}
		if ctx.nextIf('O') {
			value |= GuaranteedToOwned
		}
		if ctx.nextIf('X') {
			value |= SROA
		}
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, value))
	case 'd':
		var value rune
		if ctx.nextIf('G') {
			value |= OwnedToGuaranteed
		}
		if ctx.nextIf('O') {
			value |= GuaranteedToOwned
		}
		if ctx.nextIf('X') {
			value |= SROA
		}
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, value))
	case 'g':
		var value rune
		if ctx.nextIf('X') {
			value |= SROA
		}
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, value))
	case 'o':
		var value rune
		if ctx.nextIf('X') {
			value |= SROA
		}
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, value))
	case 'x':
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, SROA))
	case 'i':
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, BoxToValue))
	case 's':
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, BoxToStack))
	case 'r':
		return addChild(param, createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, InOutToOut))
	default:
		return nil
	}
}

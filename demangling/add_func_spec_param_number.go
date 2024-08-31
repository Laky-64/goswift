package demangling

import "unicode"

func (ctx *Context) addFuncSpecParamNumber(param *Node, kind rune) *Node {
	param.addChild(createNodeWithIndex(FunctionSignatureSpecializationParamKindKind, kind))
	var str []rune
	for unicode.IsDigit(ctx.peekChar()) {
		str = append(str, ctx.nextChar())
	}
	if len(str) == 0 {
		return nil
	}
	return addChild(param, createNodeWithText(FunctionSignatureSpecializationParamPayloadKind, string(str)))
}

package demangling

import "fmt"

func (ctx *Context) differentiableFunctionType() (*Node, error) {
	var mangledKind MangledDifferentiabilityKind
	switch ctx.nextChar() {
	case 'f':
		mangledKind = MangledForward
	case 'r':
		mangledKind = MangledReverse
	case 'd':
		mangledKind = MangledNormal
	case 'l':
		mangledKind = MangledLinear
	default:
		return nil, fmt.Errorf("unexpected differentiability kind: %c", ctx.peekChar())
	}
	return createNodeWithIndex(DifferentiableFunctionTypeKind, rune(mangledKind)), nil
}

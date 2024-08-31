package demangling

import "fmt"

func (ctx *Context) differentiabilityWitness() (*Node, error) {
	result := createNode(DifferentiabilityWitnessKind)
	optionalGenSig := ctx.popNodeKind(DependentGenericSignatureKind)
	for {
		node := ctx.popNode()
		if node == nil {
			break
		}
		result = addChild(result, node)
	}
	result.reverseChildren(0)
	var kind MangledDifferentiabilityKind
	switch ctx.nextChar() {
	case 'f':
		kind = MangledForward
	case 'r':
		kind = MangledReverse
	case 'd':
		kind = MangledNormal
	case 'l':
		kind = MangledLinear
	default:
		return nil, fmt.Errorf("unknown differentiability kind %c", ctx.Data[ctx.Pos-1])
	}
	result = addChild(result, createNodeWithIndex(IndexKind, rune(kind)))
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('p') {
		return nil, fmt.Errorf("expected 'p'")
	}
	result = addChild(result, ctx.indexSubset())
	if !ctx.nextIf('r') {
		return nil, fmt.Errorf("expected 'r'")
	}
	result = addChild(result, optionalGenSig)
	return result, nil
}

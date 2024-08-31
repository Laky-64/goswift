package demangling

import "fmt"

func (ctx *Context) genericSpecializationWithDroppedArguments() (*Node, error) {
	ctx.pushBack()
	tmp := createNode(GenericSpecializationKind)
	for ctx.nextIf('t') {
		n := rune(ctx.natural())
		addChild(tmp, createNodeWithIndex(DroppedArgumentKind, n+1))
	}
	var specKind NodeKind
	switch ctx.nextChar() {
	case 'g':
		specKind = GenericSpecializationKind
	case 'G':
		specKind = GenericSpecializationNotReAbstractedKind
	case 'B':
		specKind = GenericSpecializationInResilienceDomainKind
	default:
		return nil, fmt.Errorf("unsupported specKind: %c", ctx.Data[ctx.Pos-1])
	}
	return ctx.genericSpecialization(specKind, tmp)
}

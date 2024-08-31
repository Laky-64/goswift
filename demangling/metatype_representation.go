package demangling

import "fmt"

func (ctx *Context) metatypeRepresentation() (*Node, error) {
	switch ctx.nextChar() {
	case 't':
		return createNodeWithText(MetatypeRepresentationKind, "@thin"), nil
	case 'T':
		return createNodeWithText(MetatypeRepresentationKind, "@thick"), nil
	case 'o':
		return createNodeWithText(MetatypeRepresentationKind, "@objc_metatype"), nil
	default:
		return nil, fmt.Errorf("metatypeRepresentation: unexpected %c", ctx.Data[ctx.Pos-1])
	}
}

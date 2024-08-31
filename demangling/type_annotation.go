package demangling

import "fmt"

func (ctx *Context) typeAnnotation() (*Node, error) {
	switch ctx.nextChar() {
	case 'a':
		return createNode(AsyncAnnotationKind), nil
	case 'A':
		return createNode(IsolatedAnyFunctionTypeKind), nil
	case 'b':
		return createNode(ConcurrentFunctionTypeKind), nil
	case 'c':
		return createWithChildren(GlobalActorFunctionTypeKind, ctx.popTypeAndGetChild()), nil
	case 'i':
		return CreateType(createWithChildren(IsolatedKind, ctx.popTypeAndGetChild())), nil
	case 'j':
		return ctx.differentiableFunctionType()
	case 'k':
		return CreateType(createWithChildren(NoDerivativeKind, ctx.popTypeAndGetChild())), nil
	case 'K':
		return createWithChildren(TypedThrowsAnnotationKind, ctx.popTypeAndGetChild()), nil
	case 't':
		return CreateType(createWithChildren(CompileTimeConstKind, ctx.popTypeAndGetChild())), nil
	case 'T':
		return createNode(SendingResultFunctionTypeKind), nil
	case 'u':
		return CreateType(createWithChildren(SendingKind, ctx.popTypeAndGetChild())), nil
	case 'l':
		node, err := ctx.lifetimeDependence()
		if err != nil {
			return nil, err
		}
		addChild(node, ctx.popTypeAndGetChild())
		return CreateType(node), nil
	default:
		return nil, fmt.Errorf("unexpected type annotation: %c", ctx.peekChar())
	}
}

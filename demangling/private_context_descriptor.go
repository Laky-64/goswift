package demangling

import "fmt"

func (ctx *Context) privateContextDescriptor() (*Node, error) {
	switch ctx.nextChar() {
	case 'E':
		extension := ctx.popContext()
		if extension == nil {
			return nil, fmt.Errorf("expected an extension context")
		}
		return createWithChildren(ExtensionDescriptorKind, extension), nil
	case 'M':
		module := ctx.popModule()
		if module == nil {
			return nil, fmt.Errorf("expected a module context")
		}
		return createWithChildren(ModuleDescriptorKind, module), nil
	case 'Y':
		discriminator := ctx.popNode()
		if discriminator == nil {
			return nil, fmt.Errorf("expected a discriminator")
		}
		context := ctx.popContext()
		if context == nil {
			return nil, fmt.Errorf("expected a context")
		}
		node := createNode(AnonymousDescriptorKind)
		node.addChild(context)
		node.addChild(discriminator)
		return node, nil
	case 'X':
		context := ctx.popContext()
		if context == nil {
			return nil, fmt.Errorf("expected a context")
		}
		return createWithChildren(AnonymousDescriptorKind, context), nil
	case 'A':
		path := ctx.popAssocTypePath()
		if path == nil {
			return nil, fmt.Errorf("expected an associated type path")
		}
		base := ctx.popNodeKind(TypeKind)
		if base == nil {
			return nil, fmt.Errorf("expected a base node")
		}
		return createWithChildren(AssociatedTypeGenericParamRefKind, base, path), nil
	default:
		return nil, fmt.Errorf("unexpected private context descriptor kind: %c", ctx.Data[ctx.Pos-1])
	}
}

package demangling

import "fmt"

func (ctx *Context) standardSubstitution() (*Node, error) {
	switch ctx.nextChar() {
	case 'o':
		return createNodeWithText(ModuleKind, ManglingModuleObjC), nil
	case 'C':
		return createNodeWithText(ModuleKind, ManglingModuleSynthesized), nil
	case 'g':
		optionalType := CreateType(
			createWithChildren(
				BoundGenericEnumKind,
				createSwiftType(EnumKind, OptionalName),
				createWithChildren(TypeListKind, ctx.popNodeKind(TypeKind)),
			),
		)
		ctx.addSubstitution(optionalType)
		return optionalType, nil
	default:
		ctx.pushBack()
		repeatCount := ctx.natural()
		if repeatCount > maxRepeats {
			return nil, fmt.Errorf("repeat count too large: %d", repeatCount)
		}
		secondLevelSubstitution := ctx.nextIf('c')
		if node := createStandardSubstitution(ctx.nextChar(), secondLevelSubstitution); node != nil {
			for repeatCount > 1 {
				ctx.pushNode(node)
				repeatCount--
			}
			return node, nil
		}
		return nil, fmt.Errorf("unhandled standard substitution %c", ctx.Data[ctx.Pos-1])
	}
}

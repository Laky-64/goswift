package demangling

import "fmt"

func (ctx *Context) functionEntity() (*Node, error) {
	kind := EmptyListKind
	var args FunctionEntityArgsKind
	switch ctx.nextChar() {
	case 'D':
		args = FunctionArgsNone
		kind = DeallocatorKind
	case 'd':
		args = FunctionArgsNone
		kind = DestructorKind
	case 'E':
		args = FunctionArgsNone
		kind = IVarDestroyerKind
	case 'e':
		args = FunctionArgsNone
		kind = IVarInitializerKind
	case 'i':
		args = FunctionArgsNone
		kind = InitializerKind
	case 'C':
		args = FunctionArgsTypeAndMaybePrivateName
		kind = AllocatorKind
	case 'c':
		args = FunctionArgsTypeAndMaybePrivateName
		kind = ConstructorKind
	case 'U':
		args = FunctionArgsTypeAndIndex
		kind = ExplicitClosureKind
	case 'u':
		args = FunctionArgsTypeAndIndex
		kind = ImplicitClosureKind
	case 'A':
		args = FunctionArgsIndex
		kind = DefaultArgumentInitializerKind
	case 'm':
		return ctx.entity(MacroKind), nil
	case 'M':
		return ctx.macroExpansion()
	case 'p':
		return ctx.entity(GenericTypeParamDeclKind), nil
	case 'P':
		args = FunctionArgsNone
		kind = PropertyWrapperBackingInitializerKind
	case 'W':
		args = FunctionArgsNone
		kind = PropertyWrapperInitFromProjectedValueKind
	default:
		return nil, fmt.Errorf("unexpected function entity kind: %c", ctx.Data[ctx.Pos-1])
	}
	var nameOrIndex, paramType, labelList, context *Node
	switch args {
	case FunctionArgsNone:
	case FunctionArgsTypeAndMaybePrivateName:
		nameOrIndex = ctx.popNodeKind(PrivateDeclNameKind)
		paramType = ctx.popNodeKind(TypeKind)
		labelList = ctx.popFunctionParamLabels(paramType)
	case FunctionArgsTypeAndIndex:
		nameOrIndex = ctx.indexAsNode()
		paramType = ctx.popNodeKind(TypeKind)
	case FunctionArgsIndex:
		nameOrIndex = ctx.indexAsNode()
	case FunctionArgsContextArg:
		context = ctx.popNode()
	}
	entity := createWithChildren(kind, ctx.popContext())
	switch args {
	case FunctionArgsNone:
	case FunctionArgsIndex:
		entity = addChild(entity, nameOrIndex)
	case FunctionArgsTypeAndMaybePrivateName:
		addChild(entity, labelList)
		entity = addChild(entity, paramType)
		addChild(entity, nameOrIndex)
	case FunctionArgsTypeAndIndex:
		entity = addChild(entity, nameOrIndex)
		entity = addChild(entity, paramType)
	case FunctionArgsContextArg:
		entity = addChild(entity, context)
	}
	return entity, nil
}

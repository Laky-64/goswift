package demangling

import "fmt"

func (ctx *Context) symbolicExtendedExistentialType() (*Node, error) {
	retroactiveConformances := ctx.popRetroactiveConformances()
	args := createNode(TypeListKind)
	for {
		ty := ctx.popNodeKind(TypeKind)
		if ty == nil {
			break
		}
		args.addChild(ty)
	}
	args.reverseChildren(0)

	shape := ctx.popNode()
	if shape == nil {
		return nil, fmt.Errorf("symbolicExtendedExistentialType: shape is nil")
	}
	if shape.Kind != UniqueExtendedExistentialTypeShapeSymbolicReferenceKind &&
		shape.Kind != NonUniqueExtendedExistentialTypeShapeSymbolicReferenceKind {
		return nil, fmt.Errorf("symbolicExtendedExistentialType: shape is %s", shape.Kind)
	}

	var existentialType *Node
	if retroactiveConformances == nil {
		existentialType = createWithChildren(SymbolicExtendedExistentialTypeKind, shape, args)
	} else {
		existentialType = createWithChildren(SymbolicExtendedExistentialTypeKind, shape, args, retroactiveConformances)
	}
	return CreateType(existentialType), nil
}

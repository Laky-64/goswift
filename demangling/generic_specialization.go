package demangling

import "fmt"

func (ctx *Context) genericSpecialization(specKind NodeKind, droppedArguments *Node) (*Node, error) {
	spec := ctx.specAttributes(specKind)
	if spec == nil {
		return nil, fmt.Errorf("failed to demangle spec attributes")
	}

	if droppedArguments != nil {
		spec.addChild(droppedArguments)
	}

	tyList := ctx.popTypeList()
	if tyList == nil {
		return nil, fmt.Errorf("failed to pop type list")
	}
	spec.addChild(createWithChildren(GenericSpecializationParamKind, tyList))
	return spec, nil
}

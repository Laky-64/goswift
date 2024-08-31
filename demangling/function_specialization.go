package demangling

import "fmt"

func (ctx *Context) functionSpecialization() (*Node, error) {
	spec := ctx.specAttributes(FunctionSignatureSpecializationKind)
	for spec != nil && ctx.nextIf('_') {
		spec = addChild(spec, ctx.funcSpecParam(FunctionSignatureSpecializationParamKind))
	}
	if !ctx.nextIf('n') {
		spec = addChild(spec, ctx.funcSpecParam(FunctionSignatureSpecializationReturnKind))
	}
	if spec == nil {
		return nil, fmt.Errorf("functionSpecialization: spec is nil")
	}
	for idx := 0; idx < len(spec.Children); idx++ {
		param := spec.Children[len(spec.Children)-idx-1]
		if param.Kind != FunctionSignatureSpecializationParamKind {
			continue
		}
		if len(param.Children) == 0 {
			continue
		}
		kindNd := param.Children[0]
		paramKind := kindNd.Index
		switch paramKind {
		case ConstantPropFunction,
			ConstantPropGlobal,
			ConstantPropString,
			ConstantPropKeyPath,
			ClosureProp:
			fixedChildren := len(param.Children)
			for {
				ty := ctx.popNodeKind(TypeKind)
				if ty == nil {
					break
				}
				if paramKind != ClosureProp && paramKind != ConstantPropKeyPath {
					return nil, fmt.Errorf("functionSpecialization: unexpected paramKind")
				}
				param = addChild(param, ty)
			}
			name := ctx.popNodeKind(IdentifierKind)
			if name == nil {
				return nil, fmt.Errorf("functionSpecialization: name is nil")
			}
			text := name.Text
			if paramKind == ConstantPropString && len(text) > 0 && text[0] == '_' {
				text = text[1:]
			}
			param = addChild(param, createNodeWithText(FunctionSignatureSpecializationParamPayloadKind, text))
			param.reverseChildren(fixedChildren)
		default:
			break
		}
	}
	return spec, nil
}

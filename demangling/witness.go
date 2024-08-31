package demangling

import "fmt"

func (ctx *Context) witness() (*Node, error) {
	switch c := ctx.nextChar(); c {
	case 'C':
		return createWithChildren(EnumCaseKind, ctx.popNode()), nil
	case 'V':
		return createWithChildren(ValueWitnessTableKind, ctx.popNodeKind(TypeKind)), nil
	case 'v':
		var directness Directness
		switch ctx.nextChar() {
		case 'd':
			directness = Direct
		case 'i':
			directness = Indirect
		default:
			return nil, fmt.Errorf("unknown directness %c", ctx.Data[ctx.Pos-1])
		}
		return createWithChildren(FieldOffsetKind, createNodeWithIndex(DirectnessKind, rune(directness)), ctx.popNodePred(isEntity)), nil
	case 'S':
		return createWithChildren(ProtocolSelfConformanceWitnessTableKind, ctx.popProtocol()), nil
	case 'P':
		return createWithChildren(ProtocolWitnessTableKind, ctx.popProtocolConformance()), nil
	case 'p':
		return createWithChildren(ProtocolWitnessTablePatternKind, ctx.popProtocolConformance()), nil
	case 'G':
		return createWithChildren(GenericProtocolWitnessTableKind, ctx.popProtocolConformance()), nil
	case 'I':
		return createWithChildren(GenericProtocolWitnessTableInstantiationFunctionKind, ctx.popProtocolConformance()), nil
	case 'r':
		return createWithChildren(ResilientProtocolWitnessTableKind, ctx.popProtocolConformance()), nil
	case 'l':
		conf := ctx.popProtocolConformance()
		ty := ctx.popNodeKind(TypeKind)
		return createWithChildren(LazyProtocolWitnessTableAccessorKind, ty, conf), nil
	case 'L':
		conf := ctx.popProtocolConformance()
		ty := ctx.popNodeKind(TypeKind)
		return createWithChildren(LazyProtocolWitnessTableCacheVariableKind, ty, conf), nil
	case 'a':
		return createWithChildren(ProtocolWitnessTableAccessorKind, ctx.popProtocolConformance()), nil
	case 't':
		name := ctx.popNodePred(isDeclName)
		conf := ctx.popProtocolConformance()
		return createWithChildren(AssociatedTypeMetadataAccessorKind, conf, name), nil
	case 'T':
		protoTy := ctx.popNodeKind(TypeKind)
		conformingType := ctx.popAssocTypePath()
		conf := ctx.popProtocolConformance()
		return createWithChildren(AssociatedTypeWitnessTableAccessorKind, conf, conformingType, protoTy), nil
	case 'b':
		protoTy := ctx.popNodeKind(TypeKind)
		conf := ctx.popProtocolConformance()
		return createWithChildren(BaseWitnessTableAccessorKind, conf, protoTy), nil
	case 'O':
		switch ctx.nextChar() {
		case 'C':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedInitializeWithCopyNoValueWitnessKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedInitializeWithCopyNoValueWitnessKind, ctx.popNodeKind(TypeKind)), nil
		case 'D':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedAssignWithTakeNoValueWitnessKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedAssignWithTakeNoValueWitnessKind, ctx.popNodeKind(TypeKind)), nil
		case 'F':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedAssignWithCopyNoValueWitnessKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedAssignWithCopyNoValueWitnessKind, ctx.popNodeKind(TypeKind)), nil
		case 'H':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedDestroyNoValueWitnessKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedDestroyNoValueWitnessKind, ctx.popNodeKind(TypeKind)), nil
		case 'y':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedCopyKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedCopyKind, ctx.popNodeKind(TypeKind)), nil
		case 'e':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedConsumeKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedConsumeKind, ctx.popNodeKind(TypeKind)), nil
		case 'r':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedRetainKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedRetainKind, ctx.popNodeKind(TypeKind)), nil
		case 's':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedReleaseKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedReleaseKind, ctx.popNodeKind(TypeKind)), nil
		case 'b':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedInitializeWithTakeKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedInitializeWithTakeKind, ctx.popNodeKind(TypeKind)), nil
		case 'c':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedInitializeWithCopyKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedInitializeWithCopyKind, ctx.popNodeKind(TypeKind)), nil
		case 'd':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedAssignWithTakeKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedAssignWithTakeKind, ctx.popNodeKind(TypeKind)), nil
		case 'f':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedAssignWithCopyKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedAssignWithCopyKind, ctx.popNodeKind(TypeKind)), nil
		case 'h':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedDestroyKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedDestroyKind, ctx.popNodeKind(TypeKind)), nil
		case 'g':
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedEnumGetTagKind, ctx.popNodeKind(TypeKind), sig), nil
			}
			return createWithChildren(OutlinedEnumGetTagKind, ctx.popNodeKind(TypeKind)), nil
		case 'i':
			enumCaseIdx := ctx.indexAsNode()
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedEnumTagStoreKind, ctx.popNodeKind(TypeKind), sig, enumCaseIdx), nil
			}
			return createWithChildren(OutlinedEnumTagStoreKind, ctx.popNodeKind(TypeKind), enumCaseIdx), nil
		case 'j':
			enumCaseIdx := ctx.indexAsNode()
			if sig := ctx.popNodeKind(DependentGenericSignatureKind); sig != nil {
				return createWithChildren(OutlinedEnumProjectDataForLoadKind, ctx.popNodeKind(TypeKind), sig, enumCaseIdx), nil
			}
			return createWithChildren(OutlinedEnumProjectDataForLoadKind, ctx.popNodeKind(TypeKind), enumCaseIdx), nil
		default:
			return nil, fmt.Errorf("unknown witness O%c", ctx.Data[ctx.Pos-1])
		}
	case 'Z', 'z':
		declList := createNode(GlobalVariableOnceDeclListKind)
		var vars []*Node
		for ctx.popNodeKind(FirstElementMarkerKind) != nil {
			identifier := ctx.popNodePred(isDeclName)
			if identifier == nil {
				return nil, fmt.Errorf("expected identifier")
			}
			vars = append(vars, identifier)
		}
		for i := len(vars) - 1; i >= 0; i-- {
			declList.addChild(vars[i])
		}
		context := ctx.popContext()
		if context == nil {
			return nil, fmt.Errorf("expected context")
		}
		var kind NodeKind
		if c == 'Z' {
			kind = GlobalVariableOnceFunctionKind
		} else {
			kind = GlobalVariableOnceTokenKind
		}
		return createWithChildren(kind, context, declList), nil
	case 'J':
		return ctx.differentiabilityWitness()
	default:
		return nil, fmt.Errorf("unknown witness %c", c)
	}
}

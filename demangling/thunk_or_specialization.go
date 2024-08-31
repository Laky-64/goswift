package demangling

import "fmt"

func (ctx *Context) thunkOrSpecialization() (*Node, error) {
	switch c := ctx.nextChar(); c {
	case 'c':
		return createWithChildren(CurryThunkKind, ctx.popNode()), nil
	case 'j':
		return createWithChildren(DispatchThunkKind, ctx.popNode()), nil
	case 'q':
		return createWithChildren(MethodDescriptorKind, ctx.popNode()), nil
	case 'o':
		return createNode(ObjCAttributeKind), nil
	case 'O':
		return createNode(NonObjCAttributeKind), nil
	case 'D':
		return createNode(DynamicAttributeKind), nil
	case 'd':
		return createNode(DirectMethodReferenceAttributeKind), nil
	case 'E':
		return createNode(DistributedThunkKind), nil
	case 'F':
		return createNode(DistributedAccessorKind), nil
	case 'a':
		return createNode(PartialApplyObjCForwarderKind), nil
	case 'A':
		return createNode(PartialApplyForwarderKind), nil
	case 'm':
		return createNode(MergedFunctionKind), nil
	case 'X':
		return createNode(DynamicallyReplaceableFunctionVarKind), nil
	case 'x':
		return createNode(DynamicallyReplaceableFunctionKeyKind), nil
	case 'I':
		return createNode(DynamicallyReplaceableFunctionImplKind), nil
	case 'Y', 'Q':
		discriminator := ctx.indexAsNode()
		var kind NodeKind
		if c == 'Q' {
			kind = AsyncAwaitResumePartialFunctionKind
		} else {
			kind = AsyncSuspendResumePartialFunctionKind
		}
		return createWithChildren(kind, discriminator), nil
	case 'C':
		ty := ctx.popNodeKind(TypeKind)
		return createWithChildren(CoroutineContinuationPrototypeKind, ty), nil
	case 'z', 'Z':
		flagMode := ctx.indexAsNode()
		sig := ctx.popNodeKind(DependentGenericSignatureKind)
		resultType := ctx.popNodeKind(TypeKind)
		implType := ctx.popNodeKind(TypeKind)
		var kind NodeKind
		if c == 'z' {
			kind = ObjCAsyncCompletionHandlerImplKind
		} else {
			kind = PredefinedObjCAsyncCompletionHandlerImplKind
		}
		node := createWithChildren(kind, implType, resultType, flagMode)
		if sig != nil {
			addChild(node, sig)
		}
		return node, nil
	case 'V':
		base := ctx.popNodePred(isEntity)
		derived := ctx.popNodePred(isEntity)
		return createWithChildren(VTableThunkKind, derived, base), nil
	case 'W':
		entity := ctx.popNodePred(isEntity)
		conf := ctx.popProtocolConformance()
		return createWithChildren(ProtocolWitnessKind, conf, entity), nil
	case 'S':
		return createWithChildren(ProtocolSelfConformanceWitnessKind, ctx.popNodePred(isEntity)), nil
	case 'R', 'r', 'y':
		var kind NodeKind
		if c == 'R' {
			kind = ReabstractionThunkHelperKind
		} else if c == 'y' {
			kind = ReabstractionThunkHelperWithSelfKind
		} else {
			kind = ReabstractionThunkKind
		}
		thunk := createNode(kind)
		if genSig := ctx.popNodeKind(DependentGenericSignatureKind); genSig != nil {
			addChild(thunk, genSig)
		}
		if kind == ReabstractionThunkHelperWithSelfKind {
			addChild(thunk, ctx.popNodeKind(TypeKind))
		}
		addChild(thunk, ctx.popNodeKind(TypeKind))
		addChild(thunk, ctx.popNodeKind(TypeKind))
		return thunk, nil
	case 'g':
		return ctx.genericSpecialization(GenericSpecializationKind, nil)
	case 'G':
		return ctx.genericSpecialization(GenericSpecializationNotReAbstractedKind, nil)
	case 'B':
		return ctx.genericSpecialization(GenericSpecializationInResilienceDomainKind, nil)
	case 't':
		return ctx.genericSpecializationWithDroppedArguments()
	case 's':
		return ctx.genericSpecialization(GenericSpecializationPrespecializedKind, nil)
	case 'i':
		return ctx.genericSpecialization(InlinedGenericFunctionKind, nil)
	case 'p':
		spec := ctx.specAttributes(GenericPartialSpecializationKind)
		param := createWithChildren(GenericSpecializationParamKind, ctx.popNodeKind(TypeKind))
		return addChild(spec, param), nil
	case 'P':
		spec := ctx.specAttributes(GenericPartialSpecializationNotReAbstractedKind)
		param := createWithChildren(GenericSpecializationParamKind, ctx.popNodeKind(TypeKind))
		return addChild(spec, param), nil
	case 'f':
		return ctx.functionSpecialization()
	case 'K', 'k':
		var kind NodeKind
		if c == 'K' {
			kind = KeyPathGetterThunkHelperKind
		} else {
			kind = KeyPathSetterThunkHelperKind
		}
		isSerialized := ctx.nextIf('q')
		var types []*Node
		node := ctx.popNode()
		if node == nil || node.Kind != TypeKind {
			return nil, fmt.Errorf("expected type node")
		}
		for {
			types = append(types, node)
			node = ctx.popNode()
			if node == nil || node.Kind != TypeKind {
				break
			}
		}
		var result *Node
		if node != nil {
			if node.Kind == DependentGenericSignatureKind {
				decl := ctx.popNode()
				if decl == nil {
					return nil, fmt.Errorf("expected decl node")
				}
				result = createWithChildren(kind, decl, node)
			} else {
				result = createWithChildren(kind, node)
			}
		} else {
			return nil, fmt.Errorf("expected node")
		}
		for i := len(types) - 1; i >= 0; i-- {
			result.addChild(types[i])
		}
		if isSerialized {
			result.addChild(createNode(IsSerializedKind))
		}
		return result, nil
	case 'l':
		assocTypeName := ctx.popAssocTypeName()
		if assocTypeName == nil {
			return nil, fmt.Errorf("expected assocTypeName")
		}
		return createWithChildren(AssociatedTypeDescriptorKind, assocTypeName), nil
	case 'L':
		return createWithChildren(ProtocolRequirementsBaseDescriptorKind, ctx.popProtocol()), nil
	case 'M':
		return createWithChildren(DefaultAssociatedTypeMetadataAccessorKind, ctx.popAssocTypeName()), nil
	case 'n':
		requirementTy := ctx.popProtocol()
		conformingType := ctx.popAssocTypePath()
		protoTy := ctx.popNodeKind(TypeKind)
		return createWithChildren(AssociatedConformanceDescriptorKind, protoTy, conformingType, requirementTy), nil
	case 'N':
		requirementTy := ctx.popProtocol()
		assocTypePath := ctx.popAssocTypePath()
		protoTy := ctx.popNodeKind(TypeKind)
		return createWithChildren(DefaultAssociatedConformanceAccessorKind, protoTy, assocTypePath, requirementTy), nil
	case 'b':
		requirementTy := ctx.popProtocol()
		protoTy := ctx.popNodeKind(TypeKind)
		return createWithChildren(BaseConformanceDescriptorKind, protoTy, requirementTy), nil
	case 'H', 'h':
		var nodeKind NodeKind
		if c == 'H' {
			nodeKind = KeyPathEqualsThunkHelperKind
		} else {
			nodeKind = KeyPathHashThunkHelperKind
		}
		isSerialized := ctx.nextIf('q')
		var genericSig *Node
		var types []*Node
		node := ctx.popNode()
		if node != nil {
			if node.Kind == DependentGenericSignatureKind {
				genericSig = node
			} else if node.Kind == TypeKind {
				types = append(types, node)
			} else {
				return nil, fmt.Errorf("expected type node")
			}
		} else {
			return nil, fmt.Errorf("expected node")
		}
		for {
			node = ctx.popNode()
			if node == nil {
				break
			}
			if node.Kind != TypeKind {
				return nil, fmt.Errorf("expected type node")
			}
			types = append(types, node)
		}
		result := createNode(nodeKind)
		for i := len(types) - 1; i >= 0; i-- {
			result.addChild(types[i])
		}
		if genericSig != nil {
			result.addChild(genericSig)
		}
		if isSerialized {
			result.addChild(createNode(IsSerializedKind))
		}
		return result, nil
	case 'v':
		idx := ctx.index()
		if idx < 0 {
			return nil, fmt.Errorf("expected index")
		}
		if ctx.nextChar() == 'r' {
			return createNodeWithIndex(OutlinedReadOnlyObjectKind, idx), nil
		}
		return createNodeWithIndex(OutlinedVariableKind, idx), nil
	case 'e':
		params := ctx.bridgedMethodParams()
		if len(params) == 0 {
			return nil, fmt.Errorf("expected params")
		}
		return createNodeWithText(OutlinedBridgedMethodKind, params), nil
	case 'u':
		return createNode(AsyncFunctionPointerKind), nil
	case 'U':
		globalActor := ctx.popNodeKind(TypeKind)
		if globalActor == nil {
			return nil, fmt.Errorf("expected globalActor")
		}
		reabstraction := ctx.popNode()
		if reabstraction == nil {
			return nil, fmt.Errorf("expected reabstraction")
		}
		node := createNode(ReabstractionThunkHelperWithGlobalActorKind)
		node.addChild(reabstraction)
		node.addChild(globalActor)
		return node, nil
	case 'J':
		switch ctx.peekChar() {
		case 'S':
			ctx.nextChar()
			return ctx.autoDiffSubsetParametersThunk()
		case 'O':
			ctx.nextChar()
			return ctx.autoDiffSelfReorderingReabstractionThunk()
		case 'V':
			ctx.nextChar()
			return ctx.autoDiffFunctionOrSimpleThunk(AutoDiffDerivativeVTableThunkKind)
		default:
			return ctx.autoDiffFunctionOrSimpleThunk(AutoDiffFunctionKind)
		}
	case 'w':
		switch ctx.nextChar() {
		case 'b':
			return createNode(BackDeploymentThunkKind), nil
		case 'B':
			return createNode(BackDeploymentFallbackKind), nil
		case 'S':
			return createNode(HasSymbolQueryKind), nil
		default:
			return nil, fmt.Errorf("unexpected w-thunk kind: %c", ctx.peekChar())
		}
	default:
		return nil, fmt.Errorf("unexpected thunk or specialization kind: %c", c)
	}
}

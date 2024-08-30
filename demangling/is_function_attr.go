package demangling

func isFunctionAttr(kind NodeKind) bool {
	switch kind {
	case FunctionSignatureSpecializationKind,
		GenericSpecializationKind,
		GenericSpecializationPrespecializedKind,
		InlinedGenericFunctionKind,
		GenericSpecializationNotReAbstractedKind,
		GenericPartialSpecializationKind,
		GenericPartialSpecializationNotReAbstractedKind,
		GenericSpecializationInResilienceDomainKind,
		ObjCAttributeKind,
		NonObjCAttributeKind,
		DynamicAttributeKind,
		DirectMethodReferenceAttributeKind,
		VTableAttributeKind,
		PartialApplyForwarderKind,
		PartialApplyObjCForwarderKind,
		OutlinedVariableKind,
		OutlinedReadOnlyObjectKind,
		OutlinedBridgedMethodKind,
		MergedFunctionKind,
		DistributedThunkKind,
		DistributedAccessorKind,
		DynamicallyReplaceableFunctionImplKind,
		DynamicallyReplaceableFunctionKeyKind,
		DynamicallyReplaceableFunctionVarKind,
		AsyncFunctionPointerKind,
		AsyncAwaitResumePartialFunctionKind,
		AsyncSuspendResumePartialFunctionKind,
		AccessibleFunctionRecordKind,
		BackDeploymentThunkKind,
		BackDeploymentFallbackKind,
		HasSymbolQueryKind:
		return true
	default:
		return false
	}
}

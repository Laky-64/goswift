package demangling

func isRequirement(kind NodeKind) bool {
	switch kind {
	case DependentGenericParamPackMarkerKind,
		DependentGenericSameTypeRequirementKind,
		DependentGenericSameShapeRequirementKind,
		DependentGenericLayoutRequirementKind,
		DependentGenericConformanceRequirementKind,
		DependentGenericInverseConformanceRequirementKind:
		return true
	default:
		return false
	}
}

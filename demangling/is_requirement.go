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

/*
bool isRequirement(Node::Kind kind) {
  switch (kind) {
    case Node::Kind::DependentGenericParamPackMarker:
    case Node::Kind::DependentGenericSameTypeRequirement:
    case Node::Kind::DependentGenericSameShapeRequirement:
    case Node::Kind::DependentGenericLayoutRequirement:
    case Node::Kind::DependentGenericConformanceRequirement:
    case Node::Kind::DependentGenericInverseConformanceRequirement:
      return true;
    default:
      return false;
  }
}
*/

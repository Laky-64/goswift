package demangling

import "fmt"

//goland:noinspection GoDfaConstantCondition
func (ctx *Context) genericRequirement() (*Node, error) {
	var constraintKind GenericReqConstraintKind
	var constraintType GenericReqTypeKind
	switch ctx.nextChar() {
	case 'v':
		constraintKind = GenericReqConstraintPackMarker
		constraintType = GenericReqTypeGeneric
	case 'c':
		constraintKind = GenericReqConstraintBaseClass
		constraintType = GenericReqTypeAssoc
	case 'C':
		constraintKind = GenericReqConstraintBaseClass
		constraintType = GenericReqTypeCompoundAssoc
	case 'b':
		constraintKind = GenericReqConstraintBaseClass
		constraintType = GenericReqTypeGeneric
	case 'B':
		constraintKind = GenericReqConstraintBaseClass
		constraintType = GenericReqTypeSubstitution
	case 't':
		constraintKind = GenericReqConstraintSameType
		constraintType = GenericReqTypeAssoc
	case 'T':
		constraintKind = GenericReqConstraintSameType
		constraintType = GenericReqTypeCompoundAssoc
	case 's':
		constraintKind = GenericReqConstraintSameType
		constraintType = GenericReqTypeGeneric
	case 'S':
		constraintKind = GenericReqConstraintSameType
		constraintType = GenericReqTypeSubstitution
	case 'm':
		constraintKind = GenericReqConstraintLayout
		constraintType = GenericReqTypeAssoc
	case 'M':
		constraintKind = GenericReqConstraintLayout
		constraintType = GenericReqTypeCompoundAssoc
	case 'l':
		constraintKind = GenericReqConstraintLayout
		constraintType = GenericReqTypeGeneric
	case 'L':
		constraintKind = GenericReqConstraintLayout
		constraintType = GenericReqTypeSubstitution
	case 'p':
		constraintKind = GenericReqConstraintProtocol
		constraintType = GenericReqTypeAssoc
	case 'P':
		constraintKind = GenericReqConstraintProtocol
		constraintType = GenericReqTypeCompoundAssoc
	case 'Q':
		constraintKind = GenericReqConstraintProtocol
		constraintType = GenericReqTypeSubstitution
	case 'h':
		constraintKind = GenericReqConstraintSameShape
		constraintType = GenericReqTypeGeneric
	case 'i':
		constraintKind = GenericReqConstraintInverse
		constraintType = GenericReqTypeGeneric
		inverseKind := ctx.indexAsNode()
		if inverseKind == nil {
			return nil, fmt.Errorf("inverseKind is nil")
		}
	case 'I':
		constraintKind = GenericReqConstraintInverse
		constraintType = GenericReqTypeSubstitution
		inverseKind := ctx.indexAsNode()
		if inverseKind == nil {
			return nil, fmt.Errorf("inverseKind is nil")
		}
	default:
		ctx.pushBack()
		constraintKind = GenericReqConstraintProtocol
		constraintType = GenericReqTypeGeneric
	}
	var constrTy *Node
	switch constraintType {
	case GenericReqTypeGeneric:
		constrTy = CreateType(ctx.genericParamIndex())
	case GenericReqTypeAssoc:
		constrTy = ctx.associatedTypeSimple(nil)
	case GenericReqTypeCompoundAssoc:
		tmp, err := ctx.associatedTypeCompound(ctx.genericParamIndex())
		if err != nil {
			return nil, err
		}
		constrTy = tmp
		ctx.addSubstitution(constrTy)
	case GenericReqTypeSubstitution:
		constrTy = ctx.popNodeKind(TypeKind)
	}

	switch constraintKind {
	case GenericReqConstraintPackMarker:
		return createWithChildren(DependentGenericParamPackMarkerKind, constrTy), nil
	case GenericReqConstraintProtocol:
		return createWithChildren(DependentGenericConformanceRequirementKind, constrTy, ctx.popProtocol()), nil
	case GenericReqConstraintInverse:
		return createWithChildren(DependentGenericInverseConformanceRequirementKind, constrTy, ctx.popNodeKind(TypeKind)), nil
	case GenericReqConstraintBaseClass:
		return createWithChildren(DependentGenericConformanceRequirementKind, constrTy, ctx.popNodeKind(TypeKind)), nil
	case GenericReqConstraintSameType:
		return createWithChildren(DependentGenericSameTypeRequirementKind, constrTy, ctx.popNodeKind(TypeKind)), nil
	case GenericReqConstraintSameShape:
		return createWithChildren(DependentGenericSameShapeRequirementKind, constrTy, ctx.popNodeKind(TypeKind)), nil
	case GenericReqConstraintLayout:
		c := ctx.nextChar()
		var size, alignment *Node
		var name string
		switch c {
		case 'U':
			name = "U"
		case 'R':
			name = "R"
		case 'N':
			name = "N"
		case 'C':
			name = "C"
		case 'D':
			name = "D"
		case 'T':
			name = "T"
		case 'B':
			name = "B"
		case 'E':
			size = ctx.indexAsNode()
			if size == nil {
				return nil, fmt.Errorf("size is nil")
			}
			alignment = ctx.indexAsNode()
			name = "E"
		case 'e':
			size = ctx.indexAsNode()
			if size == nil {
				return nil, fmt.Errorf("size is nil")
			}
			name = "e"
		case 'M':
			size = ctx.indexAsNode()
			if size == nil {
				return nil, fmt.Errorf("size is nil")
			}
			alignment = ctx.indexAsNode()
			name = "M"
		case 'm':
			size = ctx.indexAsNode()
			if size == nil {
				return nil, fmt.Errorf("size is nil")
			}
			name = "m"
		case 'S':
			size = ctx.indexAsNode()
			if size == nil {
				return nil, fmt.Errorf("size is nil")
			}
			name = "S"
		default:
			return nil, fmt.Errorf("unknown layout constraint")
		}
		nameNode := createNodeWithText(IdentifierKind, name)
		nodeKind := DependentGenericLayoutRequirementKind
		layoutRequirement := createWithChildren(nodeKind, constrTy, nameNode)
		if size != nil {
			addChild(layoutRequirement, size)
		}
		if alignment != nil {
			addChild(layoutRequirement, alignment)
		}
		return layoutRequirement, nil
	}
	return nil, fmt.Errorf("unknown generic requirement")
}

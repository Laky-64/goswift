package demangling

import "fmt"

func (ctx *Context) associatedTypeCompound(base *Node) (*Node, error) {
	var assocTyNames []*Node
	var firstElem bool
	for {
		firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
		assocTyName := ctx.popAssocTypeName()
		if assocTyName == nil {
			return nil, fmt.Errorf("assocTyName is nil")
		}
		assocTyNames = append(assocTyNames, assocTyName)
		if firstElem {
			break
		}
	}
	var baseTy *Node
	if base != nil {
		baseTy = CreateType(base)
	} else {
		baseTy = ctx.popNodeKind(TypeKind)
	}

	for {
		if len(assocTyNames) == 0 {
			break
		}
		assocTy := assocTyNames[len(assocTyNames)-1]
		assocTyNames = assocTyNames[:len(assocTyNames)-1]
		depTy := createNode(DependentMemberTypeKind)
		depTy = addChild(depTy, baseTy)
		baseTy = CreateType(addChild(depTy, assocTy))
	}
	return baseTy, nil
}

package utils

import "github.com/Laky-64/goswift/demangling"

func (ctx *Context) stringGenericSignature(node *demangling.Node, depth int) {
	ctx.WriteString("<")
	numChildren := len(node.Children)
	var numGenericParams int
	for numGenericParams = 0; numGenericParams < numChildren; numGenericParams++ {
		if node.Children[numGenericParams].Kind != demangling.DependentGenericParamCountKind {
			break
		}
	}
	firstRequirement := numGenericParams
	for firstRequirement = numGenericParams; firstRequirement < numChildren; firstRequirement++ {
		child := node.Children[firstRequirement]
		if child.Kind == demangling.TypeKind {
			child = child.FirstChild()
		}
		if child.Kind != demangling.DependentGenericParamPackMarkerKind {
			break
		}
	}

	isGenericParamPack := func(depth, index rune) bool {
		for i := numGenericParams; i < firstRequirement; i++ {
			child := node.Children[i]
			if child.Kind != demangling.DependentGenericParamPackMarkerKind {
				continue
			}
			child = child.FirstChild()
			if child.Kind != demangling.TypeKind {
				continue
			}
			child = child.FirstChild()
			if child.Kind != demangling.DependentGenericParamTypeKind {
				continue
			}
			if index == child.FirstChild().Index && depth == child.Children[1].Index {
				return true
			}
		}
		return false
	}

	var gpDepth rune
	for gpDepth = 0; gpDepth < rune(numGenericParams); gpDepth++ {
		if gpDepth != 0 {
			ctx.WriteString("><")
		}
		count := node.Children[gpDepth].Index
		for index := rune(0); index < count; index++ {
			if index != 0 {
				ctx.WriteString(", ")
			}
			if index >= 128 {
				ctx.WriteString("...")
				break
			}
			if isGenericParamPack(gpDepth, index) {
				ctx.WriteString("each ")
			}
			ctx.WriteString(genericParameterName(gpDepth, index))
		}
	}

	if firstRequirement != numChildren {
		if ctx.DisplayWhereClauses {
			ctx.WriteString(" where ")
			for i := firstRequirement; i < numChildren; i++ {
				if i > firstRequirement {
					ctx.WriteString(", ")
				}
				ctx.stringNode(node.Children[i], depth+1, false)
			}
		}
	}
	ctx.WriteByte('>')
}

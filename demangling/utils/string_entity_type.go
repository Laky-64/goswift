package utils

import "swift/demangling"

func (ctx *Context) stringEntityType(entity *demangling.Node, t *demangling.Node, genericFunctionTypeList *demangling.Node, depth int) {
	labelList := entity.GetChildIf(demangling.LabelListKind)
	if labelList != nil || genericFunctionTypeList != nil {
		if genericFunctionTypeList != nil {
			ctx.WriteString("<")
			ctx.printChildren(genericFunctionTypeList, depth, ", ")
			ctx.WriteString(">")
		}
		if t.Kind == demangling.DependentGenericTypeKind {
			if genericFunctionTypeList == nil {
				ctx.stringNode(entity.FirstChild(), depth+1, false)
			}
			dependentType := t.Children[1]
			if ctx.needSpaceBeforeType(dependentType) {
				ctx.WriteByte(' ')
			}
			t = dependentType.FirstChild()
		}
		ctx.stringFunctionType(labelList, t, depth)
	} else {
		ctx.stringNode(t, depth+1, false)
	}
}

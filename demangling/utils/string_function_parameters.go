package utils

import (
	"github.com/Laky-64/goswift/demangling"
)

func (ctx *Context) stringFunctionParameters(labelList, parameterType *demangling.Node, depth int, showTypes bool) {
	if parameterType.Kind != demangling.ArgumentTupleKind {
		ctx.isValid = false
		return
	}
	parameters := parameterType.FirstChild().FirstChild()
	if parameters.Kind != demangling.TupleKind {
		if showTypes {
			ctx.WriteByte('(')
			ctx.stringNode(parameters, depth+1, false)
			ctx.WriteByte(')')
		} else {
			ctx.WriteString("(_:)")
		}
		return
	}
	getLabelFor := func(param *demangling.Node, index int) string {
		label := labelList.Children[index]
		if label.Kind == demangling.IdentifierKind {
			return label.Text
		}
		return "_"
	}
	paramIndex := 0
	hasLabels := labelList != nil && len(labelList.Children) > 0
	ctx.WriteByte('(')
	for i, param := range parameters.Children {
		if hasLabels {
			ctx.WriteString(getLabelFor(param, paramIndex) + ":")
		} else if !showTypes {
			if label := param.GetChildIf(demangling.TupleElementNameKind); label != nil {
				ctx.WriteString(label.Text + ":")
			} else {
				ctx.WriteString("_:")
			}
		}
		if hasLabels && showTypes {
			ctx.WriteByte(' ')
		}
		paramIndex++
		if showTypes {
			ctx.stringNode(param, depth+1, false)
		}
		if i < len(parameters.Children)-1 {
			ctx.WriteString(", ")
		}
	}
	ctx.WriteByte(')')
}

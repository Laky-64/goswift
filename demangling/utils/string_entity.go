package utils

import (
	"github.com/Laky-64/goswift/demangling"
	"strconv"
	"strings"
)

func (ctx *Context) stringEntity(entity *demangling.Node, depth int, asPrefixContent bool, typePr typePrinting, hasName bool, extraName string, extraIndex int, overwriteName string) *demangling.Node {
	var genericFunctionTypeList *demangling.Node
	if entity.Kind == demangling.BoundGenericFunctionKind {
		genericFunctionTypeList = entity.Children[1]
		entity = entity.FirstChild()
	}
	multiWordName := strings.Contains(extraName, " ")
	localName := hasName && entity.Children[1].Kind == demangling.LocalDeclNameKind
	if localName && ctx.DisplayLocalName {
		multiWordName = true
	}
	if asPrefixContent && (typePr != noType || multiWordName) {
		return entity
	}
	var postFixContent *demangling.Node
	context := entity.FirstChild()
	if ctx.stringContext(context) {
		if multiWordName {
			postFixContent = context
		} else {
			currentPos := ctx.Len()
			postFixContent = ctx.stringNode(context, depth+1, true)
			if ctx.Len() != currentPos {
				ctx.WriteString(".")
			}
		}
	}
	if hasName || len(overwriteName) > 0 {
		if len(extraName) > 0 && multiWordName {
			ctx.WriteString(extraName)
			if extraIndex > 0 {
				ctx.WriteString(strconv.FormatInt(int64(extraIndex), 10))
			}
			ctx.WriteString(" of ")
			extraName = ""
			extraIndex = -1
		}
		currentPos := ctx.Len()
		if len(overwriteName) > 0 {
			ctx.WriteString(overwriteName)
		} else {
			name := entity.Children[1]
			if name.Kind != demangling.PrivateDeclNameKind {
				ctx.stringNode(name, depth+1, false)
			}
			if privateName := entity.GetChildIf(demangling.PrivateDeclNameKind); privateName != nil {
				ctx.stringNode(privateName, depth+1, false)
			}
		}
		if ctx.Len() != currentPos && len(extraName) > 0 {
			ctx.WriteString(".")
		}
	}
	if len(extraName) > 0 {
		ctx.WriteString(extraName)
		if extraIndex >= 0 {
			ctx.WriteString(strconv.FormatInt(int64(extraIndex), 10))
		}
	}
	if typePr != noType {
		t := entity.GetChildIf(demangling.TypeKind)
		if t == nil {
			ctx.isValid = false
			return nil
		}
		t = t.FirstChild()
		if typePr == functionStyleType {
			t2 := t
			for t2.Kind == demangling.DependentGenericTypeKind {
				t2 = t2.Children[1].FirstChild()
			}
			if t2.Kind != demangling.FunctionTypeKind &&
				t2.Kind != demangling.NoEscapeFunctionTypeKind &&
				t2.Kind != demangling.UncurriedFunctionTypeKind &&
				t2.Kind != demangling.CFunctionPointerKind &&
				t2.Kind != demangling.ThinFunctionTypeKind {
				typePr = withColonType
			}
		}
		if typePr == withColonType {
			if ctx.DisplayEntityTypes {
				ctx.WriteString(" : ")
				ctx.stringEntityType(entity, t, genericFunctionTypeList, depth)
			}
		} else if ctx.shouldShowEntityType(entity.Kind) {
			if multiWordName || needSpaceBeforeType(t) {
				ctx.WriteByte(' ')
			}
			ctx.stringEntityType(entity, t, genericFunctionTypeList, depth)
		}
	}
	if !asPrefixContent && postFixContent != nil && (!localName || ctx.DisplayLocalName) {
		if entity.Kind == demangling.DefaultArgumentInitializerKind ||
			entity.Kind == demangling.InitializerKind ||
			entity.Kind == demangling.PropertyWrapperBackingInitializerKind ||
			entity.Kind == demangling.PropertyWrapperInitFromProjectedValueKind {
			ctx.WriteString(" of ")
		} else {
			ctx.WriteString(" in ")
		}
		ctx.stringNode(postFixContent, depth+1, false)
		postFixContent = nil
	}
	return postFixContent
}

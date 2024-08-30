package utils

import (
	"fmt"
	"github.com/Laky-64/goswift/demangling"
)

// Reference:
// https://github.com/swiftlang/swift/blob/main/lib/Demangling/NodePrinter.cpp
func (ctx *Context) stringNode(node *demangling.Node, depth int, asPrefixContent bool) *demangling.Node {
	if depth > maxDepth {
		ctx.WriteString("<<too complex>>")
		return nil
	}
	switch node.Kind {
	case demangling.TypeKind:
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ClassKind, demangling.StructureKind, demangling.EnumKind, demangling.ProtocolKind, demangling.TypeAliasKind, demangling.OtherNominalTypeKind:
		return ctx.stringEntity(node, depth, asPrefixContent, noType, true, "", -1, "")
	case demangling.ModuleKind:
		if ctx.DisplayModuleNames {
			ctx.WriteString(node.Text)
		}
		return nil
	case demangling.IdentifierKind:
		ctx.WriteString(node.Text)
		return nil
	case demangling.FirstElementMarkerKind:
		ctx.WriteString(" first-element-marker ")
		return nil
	case demangling.TupleKind:
		ctx.WriteByte('(')
		ctx.printChildren(node, depth, ", ")
		ctx.WriteByte(')')
		return nil
	case demangling.TupleElementKind:
		if label := node.GetChildIf(demangling.TupleElementNameKind); label != nil {
			ctx.WriteString(label.Text + ": ")
		}
		t := node.GetChildIf(demangling.TypeKind)
		ctx.stringNode(t, depth+1, false)
		if node.GetChildIf(demangling.VariadicMarkerKind) != nil {
			ctx.WriteString("...")
		}
		return nil
	case demangling.BoundGenericClassKind,
		demangling.BoundGenericStructureKind,
		demangling.BoundGenericEnumKind,
		demangling.BoundGenericProtocolKind,
		demangling.BoundGenericOtherNominalTypeKind,
		demangling.BoundGenericTypeAliasKind:
		ctx.stringBoundGeneric(node, depth)
		return nil
	case demangling.EmptyListKind:
		ctx.WriteString(" empty-list ")
		return nil
	case demangling.GlobalKind:
		ctx.printChildren(node, depth, "")
		return nil
	case demangling.FunctionKind, demangling.BoundGenericFunctionKind:
		return ctx.stringEntity(node, depth, asPrefixContent, functionStyleType, true, "", -1, "")
	case demangling.FunctionTypeKind,
		demangling.UncurriedFunctionTypeKind,
		demangling.NoEscapeFunctionTypeKind,
		demangling.AutoClosureTypeKind,
		demangling.EscapingAutoClosureTypeKind,
		demangling.ThinFunctionTypeKind,
		demangling.CFunctionPointerKind,
		demangling.ObjCBlockKind,
		demangling.EscapingObjCBlockKind:
		ctx.stringFunctionType(nil, node, depth)
		return nil
	case demangling.ReturnTypeKind:
		if len(node.Children) > 0 {
			ctx.stringNode(node.FirstChild(), depth+1, false)
		}
		return nil
	case demangling.DifferentiableFunctionTypeKind:
		ctx.WriteString("@differentiable ")
		switch node.Index {
		case demangling.MangledForward:
			ctx.WriteString("(_forward)")
		case demangling.MangledReverse:
			ctx.WriteString("(reverse)")
		case demangling.MangledLinear:
			ctx.WriteString("(_linear)")
		case demangling.MangledNormal:
		default:
		}
		ctx.WriteString(" ")
		return nil
	case demangling.AsyncAnnotationKind:
		ctx.WriteString(" async")
		return nil
	case demangling.NominalTypeDescriptorKind:
		ctx.WriteString("nominal type descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	default:
		ctx.WriteString(fmt.Sprintf("<<unknown kind %s>>", node.Kind))
		return nil
	}
}

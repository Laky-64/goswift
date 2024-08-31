package utils

import (
	"fmt"
	"github.com/Laky-64/goswift/demangling"
)

// Reference:
// https://github.com/swiftlang/swift/blob/main/lib/Demangling/NodePrinter.cpp
func (ctx *Context) stringNode(node *demangling.Node, depth int, asPrefixContext bool) *demangling.Node {
	if depth > maxDepth {
		ctx.WriteString("<<too complex>>")
		return nil
	}
	switch node.Kind {
	case demangling.TypeKind:
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ClassKind, demangling.StructureKind, demangling.EnumKind, demangling.ProtocolKind, demangling.TypeAliasKind, demangling.OtherNominalTypeKind:
		return ctx.stringEntity(node, depth, asPrefixContext, noType, true, "", -1, "")
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
		return ctx.stringEntity(node, depth, asPrefixContext, functionStyleType, true, "", -1, "")
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
		switch demangling.MangledDifferentiabilityKind(node.Index) {
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
	case demangling.StaticKind:
		ctx.WriteString("static ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ProtocolListKind:
		typeList := node.FirstChild()
		if typeList == nil {
			return nil
		}
		if len(typeList.Children) == 0 {
			ctx.WriteString("Any")
		} else {
			ctx.printChildren(typeList, depth, " & ")
		}
		return nil
	case demangling.ProtocolConformanceDescriptorKind:
		ctx.WriteString("protocol conformance descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ProtocolConformanceKind:
		child0 := node.Children[0]
		child1 := node.Children[1]
		child2 := node.Children[2]
		if len(node.Children) == 4 {
			ctx.WriteString("property behavior storage of ")
			ctx.stringNode(child2, depth+1, false)
			ctx.WriteString(" in ")
			ctx.stringNode(child0, depth+1, false)
			ctx.WriteString(" : ")
			ctx.stringNode(child1, depth+1, false)
		} else {
			ctx.stringNode(child0, depth+1, false)
			if ctx.DisplayProtocolConformances {
				ctx.WriteString(" : ")
				ctx.stringNode(child1, depth+1, false)
				ctx.WriteString(" in ")
				ctx.stringNode(child2, depth+1, false)
			}
		}
		return nil
	case demangling.ProtocolWitnessTableKind:
		ctx.WriteString("protocol witness table for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.TypeMetadataAccessFunctionKind:
		ctx.WriteString("type metadata accessor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.TypeMetadataKind:
		ctx.WriteString("type metadata for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.InOutKind:
		ctx.WriteString("inout ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.GetterKind:
		return ctx.abstractStorage(node.FirstChild(), depth, false, "getter")
	case demangling.PropertyDescriptorKind:
		ctx.WriteString("property descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.VariableKind:
		return ctx.stringEntity(node, depth, asPrefixContext, withColonType, true, "", -1, "")
	case demangling.MetatypeKind:
		var idx rune
		if len(node.Children) == 2 {
			ctx.stringNode(node.Children[idx], depth+1, false)
			ctx.WriteString(" ")
			idx++
		}
		t := node.Children[idx].Children[0]
		ctx.stringWithParens(t, depth)
		if IsExistentialType(t) {
			ctx.WriteString(".Protocol")
		} else {
			ctx.WriteString(".Type")
		}
		return nil
	case demangling.DependentGenericParamTypeKind:
		i := node.Children[1].Index
		d := node.Children[0].Index
		ctx.WriteString(genericParameterName(d, i))
		return nil
	case demangling.DependentGenericSignatureKind:
		ctx.stringGenericSignature(node, depth)
		return nil
	case demangling.MethodDescriptorKind:
		ctx.WriteString("method descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.FieldOffsetKind:
		ctx.stringNode(node.FirstChild(), depth+1, false)
		ctx.WriteString("field offset for ")
		ctx.stringNode(node.Children[1], depth+1, false)
		return nil
	case demangling.DirectnessKind:
		d := demangling.Directness(node.Index)
		ctx.WriteString(d.String() + " ")
		return nil
	case demangling.InitializerKind:
		return ctx.stringEntity(node, depth, asPrefixContext, noType, false, "variable initialization expression", -1, "")
	case demangling.MetaclassKind:
		ctx.WriteString("metaclass for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ConstructorKind:
		return ctx.stringEntity(node, depth, asPrefixContext, functionStyleType, len(node.Children) > 2, "init", -1, "")
	case demangling.DestructorKind:
		return ctx.stringEntity(node, depth, asPrefixContext, noType, false, "deinit", -1, "")
	case demangling.AllocatorKind:
		var extraName string
		if isClassType(node.FirstChild()) {
			extraName = "__allocating_init"
		} else {
			extraName = "init"
		}
		return ctx.stringEntity(node, depth, asPrefixContext, functionStyleType, len(node.Children) > 2, extraName, -1, "")
	case demangling.DeallocatorKind:
		var extraName string
		if isClassType(node.FirstChild()) {
			extraName = "__deallocating_deinit"
		} else {
			extraName = "deinit"
		}
		return ctx.stringEntity(node, depth, asPrefixContext, noType, false, extraName, -1, "")
	case demangling.SetterKind:
		return ctx.abstractStorage(node.FirstChild(), depth, asPrefixContext, "setter")
	case demangling.LabelListKind:
		return nil
	case demangling.ProtocolDescriptorKind:
		ctx.WriteString("protocol descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ProtocolRequirementsBaseDescriptorKind:
		ctx.WriteString("protocol requirements base descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.ModifyAccessorKind:
		return ctx.abstractStorage(node.FirstChild(), depth, asPrefixContext, "modify")
	case demangling.LocalDeclNameKind:
		ctx.stringNode(node.Children[1], depth+1, false)
		if ctx.DisplayLocalNameContexts {
			ctx.WriteString(" #" + fmt.Sprint(node.FirstChild().Index+1))
		}
		return nil
	case demangling.PrivateDeclNameKind:
		if len(node.Children) > 1 {
			if ctx.ShowPrivateDiscriminators {
				ctx.WriteByte('(')
			}
			ctx.stringNode(node.Children[1], depth+1, false)
			if ctx.ShowPrivateDiscriminators {
				ctx.WriteString(" in " + node.Children[0].Text + ")")
			}
		} else {
			if ctx.ShowPrivateDiscriminators {
				ctx.WriteString("(in " + node.Children[0].Text + ")")
			}
		}
		return nil
	case demangling.ThrowsAnnotationKind:
		ctx.WriteString(" throws")
		return nil
	case demangling.UnsafeMutableAddressorKind:
		return ctx.abstractStorage(node.FirstChild(), depth, asPrefixContext, "unsafeMutableAddressor")
	case demangling.InfixOperatorKind:
		ctx.WriteString(node.Text + " infix")
		return nil
	case demangling.PrefixOperatorKind:
		ctx.WriteString(node.Text + " prefix")
		return nil
	case demangling.PostfixOperatorKind:
		ctx.WriteString(node.Text + " postfix")
		return nil
	case demangling.BaseConformanceDescriptorKind:
		ctx.WriteString("base conformance descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		ctx.WriteString(": ")
		ctx.stringNode(node.Children[1], depth+1, false)
		return nil
	case demangling.AssociatedConformanceDescriptorKind:
		ctx.WriteString("associated conformance descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		ctx.WriteString(".")
		ctx.stringNode(node.Children[1], depth+1, false)
		ctx.WriteString(": ")
		ctx.stringNode(node.Children[2], depth+1, false)
		return nil
	case demangling.AssocTypePathKind:
		ctx.printChildren(node, len(node.Children), ".")
		return nil
	case demangling.DependentAssociatedTypeRefKind:
		if len(node.Children) > 1 {
			ctx.stringNode(node.Children[1], depth+1, false)
			ctx.WriteString(".")
		}
		ctx.stringNode(node.Children[0], depth+1, false)
		return nil
	case demangling.DependentGenericConformanceRequirementKind:
		typ := node.FirstChild()
		reqt := node.Children[1]
		ctx.stringNode(typ, depth+1, false)
		ctx.WriteString(": ")
		ctx.stringNode(reqt, depth+1, false)
		return nil
	case demangling.DependentMemberTypeKind:
		base := node.FirstChild()
		ctx.stringNode(base, depth+1, false)
		ctx.WriteString(".")
		assocTy := node.Children[1]
		ctx.stringNode(assocTy, depth+1, false)
		return nil
	case demangling.ExtensionKind:
		if ctx.QualifyEntities && ctx.DisplayExtensionContexts {
			ctx.WriteString("(extension in ")
			ctx.stringNode(node.FirstChild(), depth+1, true)
			ctx.WriteString("):")
		}
		ctx.stringNode(node.Children[1], depth+1, false)
		if len(node.Children) == 3 {
			if !ctx.PrintForTypeName {
				ctx.stringNode(node.Children[2], depth+1, false)
			}
		}
		return nil
	case demangling.SubscriptKind:
		return ctx.stringEntity(node, depth, asPrefixContext, functionStyleType, false, "", -1, "subscript")
	case demangling.AssociatedTypeDescriptorKind:
		ctx.WriteString("associated type descriptor for ")
		ctx.stringNode(node.FirstChild(), depth+1, false)
		return nil
	case demangling.DependentGenericTypeKind:
		sig := node.FirstChild()
		depTy := node.Children[1]
		ctx.stringNode(sig, depth+1, false)
		if needSpaceBeforeType(depTy) {
			ctx.WriteByte(' ')
		}
		ctx.stringNode(depTy, depth+1, false)
		return nil
	case demangling.DependentGenericSameTypeRequirementKind:
		fst := node.FirstChild()
		snd := node.Children[1]
		ctx.stringNode(fst, depth+1, false)
		ctx.WriteString(" == ")
		ctx.stringNode(snd, depth+1, false)
		return nil
	default:
		ctx.WriteString(fmt.Sprintf("<<unknown kind %s>>", node.Kind))
		return nil
	}
}

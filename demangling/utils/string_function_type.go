package utils

import (
	"github.com/Laky-64/swift/demangling"
)

func (ctx *Context) stringFunctionType(labelList, node *demangling.Node, depth int) {
	if len(node.Children) < 2 {
		ctx.isValid = false
		return
	}
	printConventionWithMangledCType := func(convention string) {
		ctx.WriteString("@convention(" + convention)
		if node.Children[0].Kind == demangling.ClangTypeKind {
			ctx.WriteString(", mangledCType: \"")
			ctx.stringNode(node.Children[0], depth+1, false)
			ctx.WriteByte('"')
		}
		ctx.WriteByte(')')
	}
	switch node.Kind {
	case demangling.FunctionTypeKind, demangling.UncurriedFunctionTypeKind, demangling.NoEscapeFunctionTypeKind:
	case demangling.AutoClosureTypeKind, demangling.EscapingAutoClosureTypeKind:
		ctx.WriteString("@autoclosure ")
	case demangling.ThinFunctionTypeKind:
		ctx.WriteString("@convention(thin) ")
	case demangling.CFunctionPointerKind:
		printConventionWithMangledCType("c")
	case demangling.EscapingObjCBlockKind:
		ctx.WriteString("@escaping ")
		fallthrough
	case demangling.ObjCBlockKind:
		printConventionWithMangledCType("block")
	default:
	}
	argIndex := len(node.Children) - 2
	var startIndex int
	var isSendable, isAsync, hasSendingResult bool
	var diffKind demangling.MangledDifferentiabilityKind
	if node.Children[startIndex].Kind == demangling.ClangTypeKind {
		startIndex++
	}
	if node.Children[startIndex].Kind == demangling.IsolatedAnyFunctionTypeKind {
		ctx.stringNode(node.Children[startIndex], depth+1, false)
		startIndex++
	}
	if node.Children[startIndex].Kind == demangling.GlobalActorFunctionTypeKind {
		ctx.stringNode(node.Children[startIndex], depth+1, false)
		startIndex++
	}
	if node.Children[startIndex].Kind == demangling.DifferentiableFunctionTypeKind {
		diffKind = demangling.MangledDifferentiabilityKind(node.Children[startIndex].Index)
		startIndex++
	}
	var thrownErrorNode *demangling.Node
	if node.Children[startIndex].Kind == demangling.ThrowsAnnotationKind || node.Children[startIndex].Kind == demangling.TypedThrowsAnnotationKind {
		thrownErrorNode = ctx.stringNode(node.Children[startIndex], depth+1, false)
		startIndex++
	}
	if node.Children[startIndex].Kind == demangling.ConcurrentFunctionTypeKind {
		startIndex++
		isSendable = true
	}
	if node.Children[startIndex].Kind == demangling.AsyncAnnotationKind {
		startIndex++
		isAsync = true
	}
	if node.Children[startIndex].Kind == demangling.SendingResultFunctionTypeKind {
		startIndex++
		hasSendingResult = true
	}
	switch diffKind {
	case demangling.MangledForward:
		ctx.WriteString("@differentiable(_forward) ")
	case demangling.MangledReverse:
		ctx.WriteString("@differentiable(reverse) ")
	case demangling.MangledLinear:
		ctx.WriteString("@differentiable(_linear) ")
	case demangling.MangledNormal:
		ctx.WriteString("@differentiable ")
	case demangling.MangledNonDifferentiable:
	}
	if isSendable {
		ctx.WriteString("@Sendable ")
	}
	ctx.stringFunctionParameters(labelList, node.Children[argIndex], depth, ctx.ShowFunctionArgumentTypes)
	if !ctx.ShowFunctionArgumentTypes {
		return
	}
	if isAsync {
		ctx.WriteString(" async")
	}
	if thrownErrorNode != nil {
		ctx.stringNode(thrownErrorNode, depth+1, false)
	}
	ctx.WriteString(" -> ")
	if hasSendingResult {
		ctx.WriteString("sending ")
	}
	ctx.stringNode(node.Children[argIndex+1], depth+1, false)
}

/*
  void printFunctionType(NodePointer LabelList, NodePointer node,
                         unsigned depth) {
    switch (diffKind) {
    case MangledDifferentiabilityKind::Forward:
      Printer << "@differentiable(_forward) ";
      break;
    case MangledDifferentiabilityKind::Reverse:
      Printer << "@differentiable(reverse) ";
      break;
    case MangledDifferentiabilityKind::Linear:
      Printer << "@differentiable(_linear) ";
      break;
    case MangledDifferentiabilityKind::Normal:
      Printer << "@differentiable ";
      break;
    case MangledDifferentiabilityKind::NonDifferentiable:
      break;
    }

    if (isSendable)
      Printer << "@Sendable ";

    printFunctionParameters(LabelList, node->getChild(argIndex), depth,
                            Options.ShowFunctionArgumentTypes);

    if (!Options.ShowFunctionArgumentTypes)
      return;

    if (isAsync)
      Printer << " async";

    if (thrownErrorNode) {
      print(thrownErrorNode, depth + 1);
    }

    Printer << " -> ";

    if (hasSendingResult)
      Printer << "sending ";

    print(node->getChild(argIndex + 1), depth + 1);
  }
*/

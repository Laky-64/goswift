package demangling

import (
	"fmt"
)

type Node struct {
	Kind     NodeKind
	Text     string
	Index    rune
	Children []*Node
}

func (node *Node) FirstChild() *Node {
	if len(node.Children) == 0 {
		return nil
	}
	return node.Children[0]
}

func (node *Node) LastChild() *Node {
	if len(node.Children) == 0 {
		return nil
	}
	return node.Children[len(node.Children)-1]
}

func (node *Node) GetChildIf(kind NodeKind) *Node {
	for _, child := range node.Children {
		if child.Kind == kind {
			return child
		}
	}
	return nil
}

func (node *Node) addChild(child *Node) {
	node.Children = append(node.Children, child)
}

func (node *Node) removeChildAt(idx int) {
	if idx < 0 || idx >= len(node.Children) {
		return
	}
	node.Children = append(node.Children[:idx], node.Children[idx+1:]...)
}

func addChild(parent, child *Node) *Node {
	if parent == nil || child == nil {
		return nil
	}
	parent.addChild(child)
	return parent
}

func (node *Node) reverseChildren(startingAt int) {
	if startingAt >= len(node.Children) {
		return
	}
	for i, j := startingAt, len(node.Children)-1; i < j; i, j = i+1, j-1 {
		node.Children[i], node.Children[j] = node.Children[j], node.Children[i]
	}
}

func createNode(kind NodeKind) *Node {
	return &Node{
		Kind: kind,
	}
}

func createNodeWithText(kind NodeKind, text string) *Node {
	return &Node{
		Kind: kind,
		Text: text,
	}
}

func createNodeWithIndex(kind NodeKind, index rune) *Node {
	return &Node{
		Kind:  kind,
		Index: index,
	}
}

func createWithChildren(kind NodeKind, children ...*Node) *Node {
	for _, child := range children {
		if child == nil {
			return nil
		}
	}
	return &Node{
		Kind:     kind,
		Children: children,
	}
}

func (ctx *Context) createWithPoppedType(kind NodeKind) *Node {
	return createWithChildren(kind, ctx.popNodeKind(TypeKind))
}

func CreateType(child *Node) *Node {
	return createWithChildren(TypeKind, child)
}

func createSwiftType(typeKind NodeKind, name string) *Node {
	return CreateType(
		createWithChildren(
			typeKind,
			createNodeWithText(ModuleKind, StdLibName),
			createNodeWithText(IdentifierKind, name),
		),
	)
}

func (ctx *Context) pushNode(node *Node) {
	ctx.nodeList = append(ctx.nodeList, node)
}

func (ctx *Context) popNode() *Node {
	node := ctx.nodeList[len(ctx.nodeList)-1]
	ctx.nodeList = ctx.nodeList[:len(ctx.nodeList)-1]
	return node
}

func (ctx *Context) popNodeKind(kind NodeKind) *Node {
	if len(ctx.nodeList) == 0 {
		return nil
	}
	if ctx.nodeList[len(ctx.nodeList)-1].Kind != kind {
		return nil
	}
	return ctx.popNode()
}

func (ctx *Context) popNodePred(pred func(NodeKind) bool) *Node {
	if len(ctx.nodeList) == 0 {
		return nil
	}
	ndKind := ctx.nodeList[len(ctx.nodeList)-1].Kind
	if !pred(ndKind) {
		return nil
	}
	return ctx.popNode()
}

func (ctx *Context) popModule() *Node {
	if node := ctx.popNodeKind(IdentifierKind); node != nil {
		newNode := createNode(ModuleKind)
		newNode.Text = node.Text
		newNode.Index = node.Index
		return newNode
	}
	return ctx.popNodeKind(ModuleKind)
}

func (ctx *Context) popContext() *Node {
	if mod := ctx.popModule(); mod != nil {
		return mod
	}
	if node := ctx.popNodeKind(TypeKind); node != nil {
		if len(node.Children) != 1 {
			return nil
		}
		child := node.FirstChild()
		if !isContext(child.Kind) {
			return nil
		}
		return child
	}
	return ctx.popNodePred(isContext)
}

func (ctx *Context) popRetroactiveConformance() *Node {
	var conformanceNode *Node
	for {
		conformance := ctx.popNodeKind(RetroactiveConformanceKind)
		if conformance == nil {
			break
		}
		if conformanceNode == nil {
			conformanceNode = createNode(TypeListKind)
		}
		conformanceNode.addChild(conformance)
	}
	if conformanceNode != nil {
		conformanceNode.reverseChildren(0)
	}
	return conformanceNode
}

func (ctx *Context) popTypeAndGetChild() *Node {
	ty := ctx.popNodeKind(TypeKind)
	if ty != nil && len(ty.Children) == 1 {
		return ty.FirstChild()
	}
	return nil
}

func (ctx *Context) popTypeAndGetAnyGeneric() *Node {
	ty := ctx.popTypeAndGetChild()
	if ty != nil && isAnyGeneric(ty.Kind) {
		return ty
	}
	return nil
}

func (ctx *Context) popTuple() (*Node, error) {
	root := createNode(TupleKind)
	if ctx.popNodeKind(EmptyListKind) == nil {
		var firstElem bool
		for !firstElem {
			firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
			tupleElem := createNode(TupleElementKind)
			addChild(tupleElem, ctx.popNodeKind(VariadicMarkerKind))
			if ident := ctx.popNodeKind(IdentifierKind); ident != nil {
				tupleElem.addChild(createNodeWithText(TupleElementNameKind, ident.Text))
			}
			ty := ctx.popNodeKind(TypeKind)
			if ty == nil {
				return nil, fmt.Errorf("expected type in tuple")
			}
			tupleElem.addChild(ty)
			root.addChild(tupleElem)
		}
		root.reverseChildren(0)
	}
	return CreateType(root), nil
}

func (ctx *Context) popFunctionType(kind NodeKind, hasClangType bool) *Node {
	funcType := createNode(kind)
	var clangType *Node
	if hasClangType {
		clangType = ctx.clangType()
	}
	addChild(funcType, clangType)
	addChild(funcType, ctx.popNodeKind(GlobalActorFunctionTypeKind))
	addChild(funcType, ctx.popNodeKind(IsolatedAnyFunctionTypeKind))
	addChild(funcType, ctx.popNodeKind(SendingResultFunctionTypeKind))
	addChild(funcType, ctx.popNodeKind(DifferentiableFunctionTypeKind))
	addChild(funcType, ctx.popNodePred(func(kind NodeKind) bool {
		return kind == ThrowsAnnotationKind || kind == TypedThrowsAnnotationKind
	}))
	addChild(funcType, ctx.popNodeKind(ConcurrentFunctionTypeKind))
	addChild(funcType, ctx.popNodeKind(AsyncAnnotationKind))
	funcType = addChild(funcType, ctx.popFunctionParams(ArgumentTupleKind))
	funcType = addChild(funcType, ctx.popFunctionParams(ReturnTypeKind))
	return CreateType(funcType)
}

func (ctx *Context) popFunctionParams(kind NodeKind) *Node {
	var paramsType *Node
	if ctx.popNodeKind(EmptyListKind) != nil {
		paramsType = CreateType(createNode(TupleKind))
	} else {
		paramsType = ctx.popNodeKind(TypeKind)
	}
	return createWithChildren(kind, paramsType)
}

func (ctx *Context) popFunctionParamLabels(typ *Node) *Node {
	if !ctx.isOldFunctionMangling && ctx.popNodeKind(EmptyListKind) != nil {
		return createNode(LabelListKind)
	}

	if typ == nil || typ.Kind != TypeKind {
		return nil
	}

	funcType := typ.FirstChild()
	if funcType.Kind == DependentGenericTypeKind {
		funcType = funcType.Children[1].FirstChild()
	}

	if funcType.Kind != FunctionTypeKind && funcType.Kind != NoEscapeFunctionTypeKind {
		return nil
	}

	var firstChildIdx int
	if funcType.Children[firstChildIdx].Kind == GlobalActorFunctionTypeKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == IsolatedAnyFunctionTypeKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == SendingResultFunctionTypeKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == DifferentiableFunctionTypeKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == ThrowsAnnotationKind ||
		funcType.Children[firstChildIdx].Kind == TypedThrowsAnnotationKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == ConcurrentFunctionTypeKind {
		firstChildIdx++
	}
	if funcType.Children[firstChildIdx].Kind == AsyncAnnotationKind {
		firstChildIdx++
	}
	paramType := funcType.Children[firstChildIdx]
	paramsType := paramType.FirstChild()
	var numParams int
	params := paramsType.FirstChild()
	if params.Kind == TupleKind {
		numParams = len(params.Children)
	} else {
		numParams = 1
	}
	if numParams == 0 {
		return nil
	}
	getChildIf := func(node *Node, kind NodeKind) (*Node, int) {
		for i, child := range node.Children {
			if child.Kind == kind {
				return child, i
			}
		}
		return nil, 0
	}
	getLabel := func(params *Node, idx int) *Node {
		if ctx.isOldFunctionMangling {
			param := params.Children[idx]
			label, i := getChildIf(param, TupleElementNameKind)
			if label != nil {
				param.removeChildAt(i)
				return createNodeWithText(IdentifierKind, label.Text)
			}
			return createNode(FirstElementMarkerKind)
		}
		return ctx.popNode()
	}
	labelList := createNode(LabelListKind)
	tuple := paramsType.FirstChild().FirstChild()

	if ctx.isOldFunctionMangling && (tuple == nil || tuple.Kind != TupleKind) {
		return labelList
	}

	hasLabels := false
	for i := 0; i < numParams; i++ {
		label := getLabel(tuple, i)
		if label == nil {
			return nil
		}
		if label.Kind != IdentifierKind && label.Kind != FirstElementMarkerKind {
			return nil
		}
		labelList.addChild(label)
		if !hasLabels {
			hasLabels = label.Kind != FirstElementMarkerKind
		}
	}
	if !hasLabels {
		return createNode(LabelListKind)
	}

	if !ctx.isOldFunctionMangling {
		labelList.reverseChildren(0)
	}
	return labelList
}

func (ctx *Context) popProtocol() *Node {
	if t := ctx.popNodeKind(TypeKind); t != nil {
		if len(t.Children) < 1 {
			return nil
		}
		if !isProtocolNode(t) {
			return nil
		}
		return t
	}

	if s := ctx.popNodeKind(ProtocolSymbolicReferenceKind); s != nil {
		return s
	} else if s = ctx.popNodeKind(ObjectiveCProtocolSymbolicReferenceKind); s != nil {
		return s
	}

	name := ctx.popNodePred(isDeclName)
	context := ctx.popContext()
	proto := createWithChildren(ProtocolKind, context, name)
	return CreateType(proto)
}

func (ctx *Context) popProtocolConformance() *Node {
	genSig := ctx.popNodeKind(DependentGenericSignatureKind)
	module := ctx.popModule()
	proto := ctx.popProtocol()
	typ := ctx.popNodeKind(TypeKind)
	var ident *Node
	if typ == nil {
		ident = ctx.popNodeKind(IdentifierKind)
		typ = ctx.popNodeKind(TypeKind)
	}
	if genSig != nil {
		typ = CreateType(createWithChildren(DependentGenericTypeKind, genSig, typ))
	}
	conformance := createWithChildren(ProtocolConformanceKind, typ, proto, module)
	addChild(conformance, ident)
	return conformance
}

func (ctx *Context) popTypeList() *Node {
	root := createNode(TypeListKind)
	if ctx.popNodeKind(EmptyListKind) == nil {
		var firstElem bool
		for !firstElem {
			firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
			ty := ctx.popNodeKind(TypeKind)
			if ty == nil {
				return nil
			}
			root.addChild(ty)
		}
		root.reverseChildren(0)
	}
	return root
}

func (ctx *Context) popAssocTypeName() *Node {
	proto := ctx.popNodeKind(TypeKind)
	if proto != nil && !isProtocolNode(proto) {
		return nil
	}
	if proto == nil {
		proto = ctx.popNodeKind(ProtocolSymbolicReferenceKind)
	}
	if proto == nil {
		proto = ctx.popNodeKind(ObjectiveCProtocolSymbolicReferenceKind)
	}
	id := ctx.popNodeKind(IdentifierKind)
	assocTy := createWithChildren(DependentAssociatedTypeRefKind, id)
	addChild(assocTy, proto)
	return assocTy
}

func (ctx *Context) popAssocTypePath() *Node {
	assocTypePath := createNode(AssocTypePathKind)
	var firstElem bool
	for !firstElem {
		firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
		assocTy := ctx.popAssocTypeName()
		if assocTy == nil {
			return nil
		}
		assocTypePath.addChild(assocTy)
	}
	assocTypePath.reverseChildren(0)
	return assocTypePath
}

func (ctx *Context) popRetroactiveConformances() *Node {
	var conformancesNode *Node
	for {
		conformance := ctx.popNodeKind(RetroactiveConformanceKind)
		if conformance == nil {
			break
		}
		if conformancesNode == nil {
			conformancesNode = createNode(TypeListKind)
		}
		conformancesNode.addChild(conformance)
	}
	if conformancesNode != nil {
		conformancesNode.reverseChildren(0)
	}
	return conformancesNode
}

func (ctx *Context) popPack() *Node {
	root := createNode(PackKind)
	if ctx.popNodeKind(EmptyListKind) == nil {
		var firstElem bool
		for !firstElem {
			firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
			ty := ctx.popNodeKind(TypeKind)
			if ty == nil {
				return nil
			}
			root.addChild(ty)
		}
		root.reverseChildren(0)
	}
	return CreateType(root)
}

func (ctx *Context) popSILPack() (*Node, error) {
	var root *Node
	switch ctx.nextChar() {
	case 'd':
		root = createNode(SILPackDirectKind)
	case 'i':
		root = createNode(SILPackIndirectKind)
	default:
		return nil, fmt.Errorf("unexpected sil pack kind: %c", ctx.nextChar())
	}
	if ctx.popNodeKind(EmptyListKind) == nil {
		var firstElem bool
		for !firstElem {
			firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
			ty := ctx.popNodeKind(TypeKind)
			if ty == nil {
				return nil, fmt.Errorf("expected type in sil pack")
			}
			root.addChild(ty)
		}
		root.reverseChildren(0)
	}
	return CreateType(root), nil
}

package demangling

import "fmt"

func (ctx *Context) macroExpansion() (*Node, error) {
	var kind NodeKind
	var isAttached, isFreestanding bool
	switch ctx.nextChar() {
	case 'a':
		kind = AccessorAttachedMacroExpansionKind
		isAttached = true
	case 'r':
		kind = MemberAttributeAttachedMacroExpansionKind
		isAttached = true
	case 'm':
		kind = MemberAttachedMacroExpansionKind
		isAttached = true
	case 'p':
		kind = PeerAttachedMacroExpansionKind
		isAttached = true
	case 'c':
		kind = ConformanceAttachedMacroExpansionKind
		isAttached = true
	case 'e':
		kind = ExtensionAttachedMacroExpansionKind
		isAttached = true
	case 'b':
		kind = BodyAttachedMacroExpansionKind
		isAttached = true
	case 'f':
		kind = FreestandingMacroExpansionKind
		isFreestanding = true
	case 'u':
		kind = MacroExpansionUniqueNameKind
	case 'X':
		kind = MacroExpansionLocKind
		line := ctx.index()
		col := ctx.index()
		lineNode := createNodeWithIndex(IndexKind, line)
		colNode := createNodeWithIndex(IndexKind, col)
		buffer := ctx.popNodeKind(IdentifierKind)
		module := ctx.popNodeKind(IdentifierKind)
		return createWithChildren(kind, module, buffer, lineNode, colNode), nil
	default:
		return nil, fmt.Errorf("unexpected macro expansion kind: %c", ctx.Data[ctx.Pos-1])
	}
	macroName := ctx.popNodeKind(IdentifierKind)
	var privateDiscriminator *Node
	if isFreestanding {
		privateDiscriminator = ctx.popNodeKind(PrivateDeclNameKind)
	}
	var attachedName *Node
	if isAttached {
		attachedName = ctx.popNodePred(isDeclName)
	}
	context := ctx.popNodePred(isMacroExpansionNodeKind)
	if context == nil {
		context = ctx.popContext()
	}
	var discriminator, result *Node
	if isAttached {
		result = createWithChildren(kind, context, attachedName, macroName, discriminator)
	} else {
		result = createWithChildren(kind, context, macroName, discriminator)
	}
	if privateDiscriminator != nil {
		result.addChild(privateDiscriminator)
	}
	return result, nil
}

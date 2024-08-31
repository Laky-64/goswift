package demangling

func (ctx *Context) accessor(childNode *Node) *Node {
	var kind NodeKind
	switch ctx.nextChar() {
	case 'm':
		kind = MaterializeForSetKind
	case 's':
		kind = SetterKind
	case 'g':
		kind = GetterKind
	case 'G':
		kind = GlobalGetterKind
	case 'w':
		kind = WillSetKind
	case 'W':
		kind = DidSetKind
	case 'r':
		kind = ReadAccessorKind
	case 'M':
		kind = ModifyAccessorKind
	case 'i':
		kind = InitAccessorKind
	case 'a':
		switch ctx.nextChar() {
		case 'O':
			kind = OwningMutableAddressorKind
		case 'o':
			kind = NativeOwningMutableAddressorKind
		case 'P':
			kind = NativePinningMutableAddressorKind
		case 'u':
			kind = UnsafeMutableAddressorKind
		default:
			return nil
		}
	case 'l':
		switch ctx.nextChar() {
		case 'O':
			kind = OwningAddressorKind
		case 'o':
			kind = NativeOwningAddressorKind
		case 'p':
			kind = NativePinningAddressorKind
		case 'u':
			kind = UnsafeAddressorKind
		default:
			return nil
		}
	case 'p':
		return childNode
	default:
		return nil
	}
	return createWithChildren(kind, childNode)
}

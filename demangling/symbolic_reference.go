package demangling

import (
	"fmt"
	"unsafe"
)

func (ctx *Context) symbolicReference() (*Node, error) {
	ctx.pushBack()
	rawKind := ctx.peekChar()
	if ctx.Pos+4 > ctx.Size {
		return nil, fmt.Errorf("unexpected end of buffer")
	}
	offset := ctx.Pos
	ctx.Pos++
	at := ctx.Data[ctx.Pos : ctx.Pos+4]
	value := *(*int32)(unsafe.Pointer(&at[0]))
	ctx.Pos += 4
	value += int32(offset)
	var kind SymbolicReferenceKind
	var direct Directness
	switch rawKind {
	case 0x01:
		kind = SymbolicContext
		direct = Direct
	case 0x02:
		kind = SymbolicContext
		direct = Indirect
	case 0x09:
		kind = SymbolicAccessorFunctionReference
		direct = Direct
	case 0x0A:
		kind = SymbolicUniqueExtendedExistentialTypeShape
		direct = Direct
	case 0x0B:
		kind = SymbolicNonUniqueExtendedExistentialTypeShape
		direct = Direct
	case 0x0C:
		kind = SymbolicObjectiveCProtocol
		direct = Direct
	case 0x03:
		fallthrough
	case 0x04:
		fallthrough
	case 0x05:
		fallthrough
	case 0x06:
		fallthrough
	case 0x07:
		fallthrough
	case 0x08:
		fallthrough
	default:
		return nil, fmt.Errorf("unhandled symbolic reference kind 0x%x", rawKind)
	}
	var resolved *Node
	if ctx.symbolicReferenceResolver != nil {
		r, err := ctx.symbolicReferenceResolver(kind, direct, value)
		if err != nil {
			return nil, err
		}
		resolved = r
	}
	if resolved == nil {
		return nil, fmt.Errorf("unresolved symbolic reference kind 0x%x", value)
	}
	if (kind == SymbolicContext || kind == SymbolicObjectiveCProtocol) &&
		resolved.Kind != OpaqueTypeDescriptorSymbolicReferenceKind &&
		resolved.Kind != OpaqueReturnTypeOfKind {
		ctx.addSubstitution(resolved)
	}
	return resolved, nil
}

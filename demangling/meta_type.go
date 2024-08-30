package demangling

import "fmt"

// Reference:
// https://github.com/swiftlang/swift/blob/4987c3b970036046ba668b0fe779297e37fa9544/lib/Demangling/Demangler.cpp#L2400
func (ctx *Context) metaType() (*Node, error) {
	switch ctx.nextChar() {
	case 'n':
		return ctx.createWithPoppedType(NominalTypeDescriptorKind), nil
	default:
		return nil, fmt.Errorf("unsupported meta type %c", ctx.peekChar())
	}
}

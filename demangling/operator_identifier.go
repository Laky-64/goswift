package demangling

import (
	"fmt"
	"unicode"
)

func (ctx *Context) operatorIdentifier() (*Node, error) {
	ident := ctx.popNodeKind(IdentifierKind)
	if ident == nil {
		return nil, fmt.Errorf("expected identifier")
	}
	opCharTable := []rune("& @/= >    <*!|+?%-~   ^ .")
	var opStr []rune
	for _, c := range ident.Text {
		if c < 0 {
			opStr = append(opStr, c)
			continue
		}
		if !unicode.IsLower(c) {
			return nil, fmt.Errorf("expected lower letter")
		}
		o := opCharTable[c-'a']
		if o == ' ' {
			return nil, fmt.Errorf("expected operator character")
		}
		opStr = append(opStr, o)
	}
	switch ctx.nextChar() {
	case 'i':
		return createNodeWithText(InfixOperatorKind, string(opStr)), nil
	case 'p':
		return createNodeWithText(PrefixOperatorKind, string(opStr)), nil
	case 'P':
		return createNodeWithText(PostfixOperatorKind, string(opStr)), nil
	default:
		return nil, fmt.Errorf("unexpected operator identifier %c", ctx.Data[ctx.Pos-1])
	}
}

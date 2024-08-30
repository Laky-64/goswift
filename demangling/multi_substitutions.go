package demangling

import (
	"fmt"
	"unicode"
)

func (ctx *Context) multiSubstitutions() (*Node, error) {
	repeatCount := -1
	for {
		c := ctx.nextChar()
		if c == 0 {
			return nil, fmt.Errorf("unexpected end of buffer")
		}
		if unicode.IsLower(c) {
			nd, err := ctx.pushMultiSubstitutions(repeatCount, int(c-'a'))
			if err != nil {
				return nil, err
			}
			ctx.pushNode(nd)
			repeatCount = -1
			continue
		}
		if unicode.IsUpper(c) {
			return ctx.pushMultiSubstitutions(repeatCount, int(c-'A'))
		}
		if c == '_' {
			idx := repeatCount + 27
			if idx >= len(ctx.substitutions) {
				return nil, fmt.Errorf("out of range substitution index %d", idx)
			}
			return ctx.substitutions[idx], nil
		}
		ctx.pushBack()
		repeatCount = ctx.natural()
		if repeatCount < 0 {
			return nil, fmt.Errorf("invalid repeat count")
		}
	}
}

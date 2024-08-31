package demangling

import "fmt"

func (ctx *Context) pushMultiSubstitutions(repeatCount, subStIdx int) (*Node, error) {
	if subStIdx >= len(ctx.substitutions) {
		return nil, fmt.Errorf("out of range substitution index %d", subStIdx)
	}
	if repeatCount > maxRepeats {
		return nil, fmt.Errorf("too many repeats")
	}
	nd := ctx.substitutions[subStIdx]
	for i := 1; i < repeatCount; i++ {
		ctx.pushNode(nd)
	}
	return nd, nil
}

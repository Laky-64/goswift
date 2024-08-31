package demangling

func (ctx *Context) index() rune {
	if ctx.nextIf('_') {
		return 0
	}
	num := rune(ctx.natural())
	if num >= 0 && ctx.nextIf('_') {
		return num + 1
	}
	return -1000
}

package demangling

func (ctx *Context) nextChar() rune {
	if ctx.Pos >= ctx.Size {
		return 0
	}
	ch := rune(ctx.Data[ctx.Pos])
	ctx.Pos++
	return ch
}

func (ctx *Context) peekChar() rune {
	if ctx.Pos >= ctx.Size {
		return 0
	}
	return rune(ctx.Data[ctx.Pos])
}

func (ctx *Context) nextIf(c rune) bool {
	if ctx.peekChar() == c {
		ctx.Pos++
		return true
	}
	return false
}

func (ctx *Context) pushBack() {
	ctx.Pos--
}

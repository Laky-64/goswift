package demangling

import "unicode"

func (ctx *Context) natural() int {
	if !unicode.IsDigit(ctx.peekChar()) {
		return -1000
	}
	var num int
	for {
		c := ctx.peekChar()
		if !unicode.IsDigit(c) {
			return num
		}
		newNum := num*10 + int(c-'0')
		if newNum < num {
			return -1000
		}
		num = newNum
		ctx.nextChar()
	}
}

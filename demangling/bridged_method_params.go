package demangling

func (ctx *Context) bridgedMethodParams() string {
	if ctx.nextIf('_') {
		return ""
	}
	var str []rune
	kind := ctx.nextChar()
	switch kind {
	case 'o', 'p', 'a', 'm':
		str = append(str, kind)
	default:
		return ""
	}
	for !ctx.nextIf('_') {
		c := ctx.nextChar()
		if c != 'n' && c != 'b' && c != 'g' {
			return ""
		}
		str = append(str, c)
	}
	return string(str)
}

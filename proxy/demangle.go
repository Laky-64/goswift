package proxy

import "github.com/Laky-64/goswift/demangling"

func (ctx *Context) Demangle(addr uint64) (*demangling.Node, error) {
	buf, err := ctx.buildBuffer(addr)
	if err != nil {
		return nil, err
	}
	demangler, err := demangling.New(buf)
	if err != nil {
		return nil, err
	}
	demangler.SetSymbolicReferenceResolver(func(kind demangling.SymbolicReferenceKind, directness demangling.Directness, index int32) (*demangling.Node, error) {
		return ctx.lookup(kind, addr+uint64(1+int64(index)))
	})
	node, err := demangler.Result()
	if err != nil {
		return nil, err
	}
	return node, nil
}

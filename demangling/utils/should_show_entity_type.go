package utils

import "github.com/Laky-64/swift/demangling"

func (ctx *Context) shouldShowEntityType(kind demangling.NodeKind) bool {
	switch kind {
	case demangling.ExplicitClosureKind, demangling.ImplicitClosureKind:
		return ctx.ShowClosureSignature
	default:
		return true
	}
}

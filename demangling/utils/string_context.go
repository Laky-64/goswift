package utils

import (
	"strings"
	"swift/demangling"
)

func (ctx *Context) stringContext(node *demangling.Node) bool {
	if !ctx.QualifyEntities {
		return false
	}
	if node.Kind == demangling.ModuleKind {
		switch node.Text {
		case demangling.StdLibName:
			return ctx.DisplayStdlibModule
		case demangling.ManglingModuleObjC:
			return ctx.DisplayObjCModule
		case ctx.HidingCurrentModule:
			return false
		}
		if strings.HasPrefix(node.Text, demangling.LLDBExpressionMangling) {
			return ctx.DisplayDebuggerGeneratedModule
		}
	}
	return true
}

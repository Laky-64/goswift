package utils

import "github.com/Laky-64/goswift/demangling"

func IsSwiftModule(node *demangling.Node) bool {
	return node.Kind == demangling.ModuleKind && node.Text == demangling.StdLibName
}

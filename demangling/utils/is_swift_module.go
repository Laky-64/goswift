package utils

import "swift/demangling"

func IsSwiftModule(node *demangling.Node) bool {
	return node.Kind == demangling.ModuleKind && node.Text == demangling.StdLibName
}

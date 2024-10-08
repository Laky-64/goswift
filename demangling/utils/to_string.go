package utils

import (
	"github.com/Laky-64/goswift/demangling"
)

func ToString(node *demangling.Node, opt OptionMode) string {
	p := &Context{
		Options: &Options{
			DisplayLocalName:               true,
			DisplayStdlibModule:            opt&ModeSimplified == 0,
			QualifyEntities:                true,
			DisplayObjCModule:              true,
			DisplayDebuggerGeneratedModule: true,
			DisplayEntityTypes:             true,
			ShowClosureSignature:           true,
			ShowFunctionArgumentTypes:      true,
			DisplayProtocolConformances:    true,
			DisplayWhereClauses:            true,
			DisplayLocalNameContexts:       true,
			ShowPrivateDiscriminators:      true,
			DisplayExtensionContexts:       true,
			PrintForTypeName:               true,
			DisplayModuleNames:             opt&ModeSimplified == 0,
			SynthesizeSugarOnTypes:         opt&ModeNoSugar == 0,
		},
		isValid: true,
	}
	p.stringNode(node, 0, false)
	return p.String()
}

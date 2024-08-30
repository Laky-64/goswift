package utils

import (
	"strings"
)

type OptionMode int

const (
	ModeDefault OptionMode = 1 << iota
	ModeNoSugar
	ModeSimplified
)

type Context struct {
	*Options
	strings.Builder
	isValid bool
}

type Options struct {
	DisplayLocalName               bool
	DisplayStdlibModule            bool
	QualifyEntities                bool
	DisplayObjCModule              bool
	DisplayDebuggerGeneratedModule bool
	DisplayEntityTypes             bool
	ShowClosureSignature           bool
	DisplayModuleNames             bool
	SynthesizeSugarOnTypes         bool
	ShowFunctionArgumentTypes      bool
	HidingCurrentModule            string
}

type typePrinting int
type sugarType int

const maxDepth = 768
const (
	noType typePrinting = iota
	functionStyleType
	withColonType
)
const (
	sugarNone sugarType = iota
	sugarOptional
	sugarImplicitlyUnwrappedOptional
	sugarArray
	sugarDictionary
)

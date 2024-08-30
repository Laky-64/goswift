package demangling

const (
	maxWords   = 26
	maxRepeats = 2048
)

type (
	SymbolicReferenceKind         int
	Directness                    int
	MangledDifferentiabilityKind  rune
	MangledLifetimeDependenceKind rune
)

const (
	ManglingModuleObjC        = "__C"
	ManglingModuleSynthesized = "__C_Synthesized"
	LLDBExpressionMangling    = "__lldb_expr_"
	StdLibName                = "Swift"
	OptionalName              = "Optional"
)

const (
	SymbolicContext SymbolicReferenceKind = iota
	SymbolicAccessorFunctionReference
	SymbolicUniqueExtendedExistentialTypeShape
	SymbolicNonUniqueExtendedExistentialTypeShape
	SymbolicObjectiveCProtocol
)

const (
	Direct Directness = 1 << iota
	Indirect
)

const (
	MangledNonDifferentiable MangledDifferentiabilityKind = 0
	MangledForward                                        = 'f'
	MangledReverse                                        = 'r'
	MangledNormal                                         = 'd'
	MangledLinear                                         = 'l'
)

const (
	UnknownLifetime        MangledLifetimeDependenceKind = 0
	MangledLifetimeInherit                               = 'i'
	MangledLifetimeScope                                 = 's'
)

//go:generate go run ../cmd/generator/.

type Context struct {
	Data                      []byte
	Pos, Size                 int
	words                     [maxWords]string
	isOldFunctionMangling     bool
	numWords                  int
	nodeList                  []*Node
	substitutions             []*Node
	symbolicReferenceResolver func(SymbolicReferenceKind, Directness, int32) (*Node, error)
}

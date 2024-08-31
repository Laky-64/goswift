package demangling

const (
	maxWords              = 26
	maxRepeats            = 2048
	maxSpecializationPass = 10
)

type (
	SymbolicReferenceKind         int
	Directness                    int
	FunctionEntityArgsKind        int
	GenericReqTypeKind            int
	GenericReqConstraintKind      int
	MangledDifferentiabilityKind  rune
	MangledLifetimeDependenceKind rune
)

func (c *Directness) String() string {
	switch *c {
	case Direct:
		return "direct"
	case Indirect:
		return "indirect"
	default:
		return "unknown"
	}
}

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
	MangledForward           MangledDifferentiabilityKind = 'f'
	MangledReverse           MangledDifferentiabilityKind = 'r'
	MangledNormal            MangledDifferentiabilityKind = 'd'
	MangledLinear            MangledDifferentiabilityKind = 'l'
)

const (
	UnknownLifetime        MangledLifetimeDependenceKind = 0
	MangledLifetimeInherit MangledLifetimeDependenceKind = 'i'
	MangledLifetimeScope   MangledLifetimeDependenceKind = 's'
)

const (
	FunctionArgsNone FunctionEntityArgsKind = iota
	FunctionArgsTypeAndMaybePrivateName
	FunctionArgsTypeAndIndex
	FunctionArgsIndex
	FunctionArgsContextArg
)

const (
	ConstantPropFunction = 0
	ConstantPropGlobal   = 1
	ConstantPropInteger  = 2
	ConstantPropFloat    = 3
	ConstantPropString   = 4
	ClosureProp          = 5
	BoxToValue           = 6
	BoxToStack           = 7
	InOutToOut           = 8
	ConstantPropKeyPath  = 9

	Dead                 = 1 << 6
	OwnedToGuaranteed    = 1 << 7
	SROA                 = 1 << 8
	GuaranteedToOwned    = 1 << 9
	ExistentialToGeneric = 1 << 10
)

const (
	GenericReqTypeGeneric GenericReqTypeKind = iota
	GenericReqTypeAssoc
	GenericReqTypeCompoundAssoc
	GenericReqTypeSubstitution
)

const (
	GenericReqConstraintProtocol GenericReqConstraintKind = iota
	GenericReqConstraintBaseClass
	GenericReqConstraintSameType
	GenericReqConstraintSameShape
	GenericReqConstraintLayout
	GenericReqConstraintPackMarker
	GenericReqConstraintInverse
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

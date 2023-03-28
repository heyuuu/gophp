package ast

type Node interface{}

type (
	// ComplexType : PhpParserNodeAbstract
	ComplexType interface {
		Node
		complexTypeNode()
	}

	// Expr : PhpParserNodeAbstract
	Expr interface {
		Node
		exprNode()
	}

	// ExprAssignOp : Expr
	AssignOpExpr interface {
		Expr
		assignOpExprNode()
	}

	// ExprBinaryOp : Expr
	BinaryOpExpr interface {
		Expr
		binaryOpExprNode()
	}

	// ExprCallLike : Expr
	CallLikeExpr interface {
		Expr
		callLikeExprNode()
	}

	// ExprCast : Expr
	CastExpr interface {
		Expr
		castExprNode()
	}

	// FunctionLike
	FunctionLike interface {
		functionLikeNode()
	}

	// Scalar : Expr
	Scalar interface {
		Expr
		scalarNode()
	}

	// ScalarMagicConst : Scalar
	MagicConstScalar interface {
		Scalar
		magicConstScalarNode()
	}

	// Stmt : PhpParserNodeAbstract
	Stmt interface {
		Node
		stmtNode()
	}

	// StmtClassLike : Stmt
	ClassLikeStmt interface {
		Stmt
		classLikeStmtNode()
	}

	// StmtTraitUseAdaptation : Stmt
	TraitUseAdaptationStmt interface {
		Stmt
		traitUseAdaptationStmtNode()
	}
)
type (
	// Arg : PhpParserNodeAbstract
	Arg struct {
		Name       *Identifier // @var Identifier|null Parameter name (for named parameters)
		Value      Expr        // @var Expr Value to pass
		ByRef      bool        // @var bool Whether to pass by ref
		Unpack     bool        // @var bool Whether to unpack the argument
		Attributes any
	}

	// Attribute : PhpParserNodeAbstract
	Attribute struct {
		Name       *Name  // @var Name Attribute name
		Args       []*Arg // @var Arg[] Attribute arguments
		Attributes any
	}

	// AttributeGroup : PhpParserNodeAbstract
	AttributeGroup struct {
		Attrs      []*Attribute // @var Attribute[] Attributes
		Attributes any
	}

	// Const : PhpParserNodeAbstract
	Const struct {
		Name           *Identifier // @var Identifier Name
		Value          Expr        // @var Expr Value
		NamespacedName *Name       // @var Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// ExprArray : Expr
	ArrayExpr struct {
		Items      []*ArrayItemExpr // @var (ArrayItem|null)[] Items
		Attributes any
	}

	// ExprArrayDimFetch : Expr
	ArrayDimFetchExpr struct {
		Var        Expr // @var Expr Variable
		Dim        Expr // @var null|Expr Array index / dim
		Attributes any
	}

	// ExprArrayItem : Expr
	ArrayItemExpr struct {
		Key        Expr // @var null|Expr Key
		Value      Expr // @var Expr Value
		ByRef      bool // @var bool Whether to assign by reference
		Unpack     bool // @var bool Whether to unpack the argument
		Attributes any
	}

	// ExprArrowFunction : Expr, FunctionLike
	ArrowFunctionExpr struct {
		Static     bool              // @var bool
		ByRef      bool              // @var bool
		Params     []*Param          // @var Node\Param[]
		ReturnType any               // @var null|Node\Identifier|Node\Name|Node\ComplexType
		Expr       Expr              // @var Expr
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
		Attributes any
	}

	// ExprAssign : Expr
	AssignExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpBitwiseAnd : ExprAssignOp
	AssignOpBitwiseAndExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpBitwiseOr : ExprAssignOp
	AssignOpBitwiseOrExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpBitwiseXor : ExprAssignOp
	AssignOpBitwiseXorExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpCoalesce : ExprAssignOp
	AssignOpCoalesceExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpConcat : ExprAssignOp
	AssignOpConcatExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpDiv : ExprAssignOp
	AssignOpDivExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpMinus : ExprAssignOp
	AssignOpMinusExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpMod : ExprAssignOp
	AssignOpModExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpMul : ExprAssignOp
	AssignOpMulExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpPlus : ExprAssignOp
	AssignOpPlusExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpPow : ExprAssignOp
	AssignOpPowExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpShiftLeft : ExprAssignOp
	AssignOpShiftLeftExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignOpShiftRight : ExprAssignOp
	AssignOpShiftRightExpr struct {
		Var        Expr // @var Expr Variable
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprAssignRef : Expr
	AssignRefExpr struct {
		Var        Expr // @var Expr Variable reference is assigned to
		Expr       Expr // @var Expr Variable which is referenced
		Attributes any
	}

	// ExprBinaryOpBitwiseAnd : ExprBinaryOp
	BinaryOpBitwiseAndExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpBitwiseOr : ExprBinaryOp
	BinaryOpBitwiseOrExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpBitwiseXor : ExprBinaryOp
	BinaryOpBitwiseXorExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpBooleanAnd : ExprBinaryOp
	BinaryOpBooleanAndExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpBooleanOr : ExprBinaryOp
	BinaryOpBooleanOrExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpCoalesce : ExprBinaryOp
	BinaryOpCoalesceExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpConcat : ExprBinaryOp
	BinaryOpConcatExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpDiv : ExprBinaryOp
	BinaryOpDivExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpEqual : ExprBinaryOp
	BinaryOpEqualExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpGreater : ExprBinaryOp
	BinaryOpGreaterExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpGreaterOrEqual : ExprBinaryOp
	BinaryOpGreaterOrEqualExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpIdentical : ExprBinaryOp
	BinaryOpIdenticalExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpLogicalAnd : ExprBinaryOp
	BinaryOpLogicalAndExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpLogicalOr : ExprBinaryOp
	BinaryOpLogicalOrExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpLogicalXor : ExprBinaryOp
	BinaryOpLogicalXorExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpMinus : ExprBinaryOp
	BinaryOpMinusExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpMod : ExprBinaryOp
	BinaryOpModExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpMul : ExprBinaryOp
	BinaryOpMulExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpNotEqual : ExprBinaryOp
	BinaryOpNotEqualExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpNotIdentical : ExprBinaryOp
	BinaryOpNotIdenticalExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpPlus : ExprBinaryOp
	BinaryOpPlusExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpPow : ExprBinaryOp
	BinaryOpPowExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpShiftLeft : ExprBinaryOp
	BinaryOpShiftLeftExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpShiftRight : ExprBinaryOp
	BinaryOpShiftRightExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpSmaller : ExprBinaryOp
	BinaryOpSmallerExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpSmallerOrEqual : ExprBinaryOp
	BinaryOpSmallerOrEqualExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBinaryOpSpaceship : ExprBinaryOp
	BinaryOpSpaceshipExpr struct {
		Left       Expr // @var Expr The left hand side expression
		Right      Expr // @var Expr The right hand side expression
		Attributes any
	}

	// ExprBitwiseNot : Expr
	BitwiseNotExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprBooleanNot : Expr
	BooleanNotExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastArray : ExprCast
	CastArrayExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastBool : ExprCast
	CastBoolExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastDouble : ExprCast
	CastDoubleExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastInt : ExprCast
	CastIntExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastObject : ExprCast
	CastObjectExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastString : ExprCast
	CastStringExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprCastUnset : ExprCast
	CastUnsetExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprClassConstFetch : Expr
	ClassConstFetchExpr struct {
		Class      any // @var Name|Expr Class name
		Name       any // @var Identifier|Error Constant name
		Attributes any
	}

	// ExprClone : Expr
	CloneExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprClosure : Expr, FunctionLike
	ClosureExpr struct {
		Static     bool              // @var bool Whether the closure is static
		ByRef      bool              // @var bool Whether to return by reference
		Params     []*Param          // @var Node\Param[] Parameters
		Uses       []*ClosureUseExpr // @var ClosureUse[] use()s
		ReturnType any               // @var null|Node\Identifier|Node\Name|Node\ComplexType Return type
		Stmts      []Stmt            // @var Node\Stmt[] Statements
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		Attributes any
	}

	// ExprClosureUse : Expr
	ClosureUseExpr struct {
		Var        *VariableExpr // @var Expr\Variable Variable to use
		ByRef      bool          // @var bool Whether to use by reference
		Attributes any
	}

	// ExprConstFetch : Expr
	ConstFetchExpr struct {
		Name       *Name // @var Name Constant name
		Attributes any
	}

	// ExprEmpty : Expr
	EmptyExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprError : Expr
	ErrorExpr struct {
		Attributes any
	}

	// ExprErrorSuppress : Expr
	ErrorSuppressExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprEval : Expr
	EvalExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprExit : Expr
	ExitExpr struct {
		Expr       Expr // @var null|Expr Expression
		Attributes any
	}

	// ExprFuncCall : ExprCallLike
	FuncCallExpr struct {
		Name       any   // @var Node\Name|Expr Function name
		Args       []any // @var array<Node\Arg|Node\VariadicPlaceholder> Arguments
		Attributes any
	}

	// ExprInclude : Expr
	IncludeExpr struct {
		Expr       Expr // @var Expr Expression
		Type       int  // @var int Type of include
		Attributes any
	}

	// ExprInstanceof : Expr
	InstanceofExpr struct {
		Expr       Expr // @var Expr Expression
		Class      any  // @var Name|Expr Class name
		Attributes any
	}

	// ExprIsset : Expr
	IssetExpr struct {
		Vars       []Expr // @var Expr[] Variables
		Attributes any
	}

	// ExprList : Expr
	ListExpr struct {
		Items      []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
		Attributes any
	}

	// ExprMatch : Expr
	MatchExpr struct {
		Cond       Expr        // @var Node\Expr
		Arms       []*MatchArm // @var MatchArm[]
		Attributes any
	}

	// ExprMethodCall : ExprCallLike
	MethodCallExpr struct {
		Var        Expr  // @var Expr Variable holding object
		Name       any   // @var Identifier|Expr Method name
		Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
		Attributes any
	}

	// ExprNew : ExprCallLike
	NewExpr struct {
		Class      any   // @var Node\Name|Expr|Node\Stmt\Class_ Class name
		Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
		Attributes any
	}

	// ExprNullsafeMethodCall : ExprCallLike
	NullsafeMethodCallExpr struct {
		Var        Expr  // @var Expr Variable holding object
		Name       any   // @var Identifier|Expr Method name
		Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
		Attributes any
	}

	// ExprNullsafePropertyFetch : Expr
	NullsafePropertyFetchExpr struct {
		Var        Expr // @var Expr Variable holding object
		Name       any  // @var Identifier|Expr Property name
		Attributes any
	}

	// ExprPostDec : Expr
	PostDecExpr struct {
		Var        Expr // @var Expr Variable
		Attributes any
	}

	// ExprPostInc : Expr
	PostIncExpr struct {
		Var        Expr // @var Expr Variable
		Attributes any
	}

	// ExprPreDec : Expr
	PreDecExpr struct {
		Var        Expr // @var Expr Variable
		Attributes any
	}

	// ExprPreInc : Expr
	PreIncExpr struct {
		Var        Expr // @var Expr Variable
		Attributes any
	}

	// ExprPrint : Expr
	PrintExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprPropertyFetch : Expr
	PropertyFetchExpr struct {
		Var        Expr // @var Expr Variable holding object
		Name       any  // @var Identifier|Expr Property name
		Attributes any
	}

	// ExprShellExec : Expr
	ShellExecExpr struct {
		Parts      []any // @var array Encapsed string array
		Attributes any
	}

	// ExprStaticCall : ExprCallLike
	StaticCallExpr struct {
		Class      any   // @var Node\Name|Expr Class name
		Name       any   // @var Identifier|Expr Method name
		Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
		Attributes any
	}

	// ExprStaticPropertyFetch : Expr
	StaticPropertyFetchExpr struct {
		Class      any // @var Name|Expr Class name
		Name       any // @var VarLikeIdentifier|Expr Property name
		Attributes any
	}

	// ExprTernary : Expr
	TernaryExpr struct {
		Cond       Expr // @var Expr Condition
		If         Expr // @var null|Expr Expression for true
		Else       Expr // @var Expr Expression for false
		Attributes any
	}

	// ExprThrow : Expr
	ThrowExpr struct {
		Expr       Expr // @var Node\Expr Expression
		Attributes any
	}

	// ExprUnaryMinus : Expr
	UnaryMinusExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprUnaryPlus : Expr
	UnaryPlusExpr struct {
		Expr       Expr // @var Expr Expression
		Attributes any
	}

	// ExprVariable : Expr
	VariableExpr struct {
		Name       any // @var string|Expr Name
		Attributes any
	}

	// ExprYield : Expr
	YieldExpr struct {
		Key        Expr // @var null|Expr Key expression
		Value      Expr // @var null|Expr Value expression
		Attributes any
	}

	// ExprYieldFrom : Expr
	YieldFromExpr struct {
		Expr       Expr // @var Expr Expression to yield from
		Attributes any
	}

	// Identifier : PhpParserNodeAbstract
	Identifier struct {
		Name              string // @var string Identifier as string
		SpecialClassNames any
		Attributes        any
	}

	// IntersectionType : ComplexType
	IntersectionType struct {
		Types      []any // @var (Identifier|Name)[] Types
		Attributes any
	}

	// MatchArm : PhpParserNodeAbstract
	MatchArm struct {
		Conds      []Expr // @var null|Node\Expr[]
		Body       Expr   // @var Node\Expr
		Attributes any
	}

	// Name : PhpParserNodeAbstract
	Name struct {
		Parts             []string // @var string[] Parts of the name
		SpecialClassNames any
		Attributes        any
	}

	// NameFullyQualified : Name
	NameFullyQualified struct {
		Parts      []string // @var string[] Parts of the name
		Attributes any
	}

	// NameRelative : Name
	NameRelative struct {
		Parts      []string // @var string[] Parts of the name
		Attributes any
	}

	// NullableType : ComplexType
	NullableType struct {
		Type       any // @var Identifier|Name Type
		Attributes any
	}

	// Param : PhpParserNodeAbstract
	Param struct {
		Type       any               // @var null|Identifier|Name|ComplexType Type declaration
		ByRef      bool              // @var bool Whether parameter is passed by reference
		Variadic   bool              // @var bool Whether this is a variadic argument
		Var        any               // @var Expr\Variable|Expr\Error Parameter variable
		Default    Expr              // @var null|Expr Default value
		Flags      int               // @var int
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		Attributes any
	}

	// ScalarDNumber : Scalar
	DNumberScalar struct {
		Value      float64 // @var float Number value
		Attributes any
	}

	// ScalarEncapsed : Scalar
	EncapsedScalar struct {
		Parts      []Expr // @var Expr[] list of string parts
		Attributes any
	}

	// ScalarEncapsedStringPart : Scalar
	EncapsedStringPartScalar struct {
		Value      string // @var string String value
		Attributes any
	}

	// ScalarLNumber : Scalar
	LNumberScalar struct {
		Value      int // @var int Number value
		Attributes any
	}

	// ScalarMagicConstClass : ScalarMagicConst
	MagicConstClassScalar struct {
		Attributes any
	}

	// ScalarMagicConstDir : ScalarMagicConst
	MagicConstDirScalar struct {
		Attributes any
	}

	// ScalarMagicConstFile : ScalarMagicConst
	MagicConstFileScalar struct {
		Attributes any
	}

	// ScalarMagicConstFunction : ScalarMagicConst
	MagicConstFunctionScalar struct {
		Attributes any
	}

	// ScalarMagicConstLine : ScalarMagicConst
	MagicConstLineScalar struct {
		Attributes any
	}

	// ScalarMagicConstMethod : ScalarMagicConst
	MagicConstMethodScalar struct {
		Attributes any
	}

	// ScalarMagicConstNamespace : ScalarMagicConst
	MagicConstNamespaceScalar struct {
		Attributes any
	}

	// ScalarMagicConstTrait : ScalarMagicConst
	MagicConstTraitScalar struct {
		Attributes any
	}

	// ScalarString : Scalar
	StringScalar struct {
		Value        string // @var string String value
		Replacements any
		Attributes   any
	}

	// StmtBreak : Stmt
	BreakStmt struct {
		Num        Expr // @var null|Node\Expr Number of loops to break
		Attributes any
	}

	// StmtCase : Stmt
	CaseStmt struct {
		Cond       Expr   // @var null|Node\Expr Condition (null for default)
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtCatch : Stmt
	CatchStmt struct {
		Types      []*Name       // @var Node\Name[] Types of exceptions to catch
		Var        *VariableExpr // @var Expr\Variable|null Variable for exception
		Stmts      []Stmt        // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtClass : StmtClassLike
	ClassStmt struct {
		Flags          int               // @var int Type
		Extends        *Name             // @var null|Node\Name Name of extended class
		Implements     []*Name           // @var Node\Name[] Names of implemented interfaces
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// StmtClassConst : Stmt
	ClassConstStmt struct {
		Flags      int               // @var int Modifiers
		Consts     []*Const          // @var Node\Const_[] Constant declarations
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
		Attributes any
	}

	// StmtClassMethod : Stmt, FunctionLike
	ClassMethodStmt struct {
		Flags      int               // @var int Flags
		ByRef      bool              // @var bool Whether to return by reference
		Name       *Identifier       // @var Node\Identifier Name
		Params     []*Param          // @var Node\Param[] Parameters
		ReturnType any               // @var null|Node\Identifier|Node\Name|Node\ComplexType Return type
		Stmts      []Stmt            // @var Node\Stmt[]|null Statements
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		MagicNames any
		Attributes any
	}

	// StmtConst : Stmt
	ConstStmt struct {
		Consts     []*Const // @var Node\Const_[] Constant declarations
		Attributes any
	}

	// StmtContinue : Stmt
	ContinueStmt struct {
		Num        Expr // @var null|Node\Expr Number of loops to continue
		Attributes any
	}

	// StmtDeclare : Stmt
	DeclareStmt struct {
		Declares   []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
		Stmts      []Stmt                // @var Node\Stmt[]|null Statements
		Attributes any
	}

	// StmtDeclareDeclare : Stmt
	DeclareDeclareStmt struct {
		Key        *Identifier // @var Node\Identifier Key
		Value      Expr        // @var Node\Expr Value
		Attributes any
	}

	// StmtDo : Stmt
	DoStmt struct {
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Cond       Expr   // @var Node\Expr Condition
		Attributes any
	}

	// StmtEcho : Stmt
	EchoStmt struct {
		Exprs      []Expr // @var Node\Expr[] Expressions
		Attributes any
	}

	// StmtElse : Stmt
	ElseStmt struct {
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtElseIf : Stmt
	ElseIfStmt struct {
		Cond       Expr   // @var Node\Expr Condition
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtEnum : StmtClassLike
	EnumStmt struct {
		ScalarType     *Identifier       // @var null|Node\Identifier Scalar Type
		Implements     []*Name           // @var Node\Name[] Names of implemented interfaces
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// StmtEnumCase : Stmt
	EnumCaseStmt struct {
		Name       *Identifier       // @var Node\Identifier Enum case name
		Expr       Expr              // @var Node\Expr|null Enum case expression
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		Attributes any
	}

	// StmtExpression : Stmt
	ExpressionStmt struct {
		Expr       Expr // @var Node\Expr Expression
		Attributes any
	}

	// StmtFinally : Stmt
	FinallyStmt struct {
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtFor : Stmt
	ForStmt struct {
		Init       []Expr // @var Node\Expr[] Init expressions
		Cond       []Expr // @var Node\Expr[] Loop conditions
		Loop       []Expr // @var Node\Expr[] Loop expressions
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtForeach : Stmt
	ForeachStmt struct {
		Expr       Expr   // @var Node\Expr Expression to iterate
		KeyVar     Expr   // @var null|Node\Expr Variable to assign key to
		ByRef      bool   // @var bool Whether to assign value by reference
		ValueVar   Expr   // @var Node\Expr Variable to assign value to
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtFunction : Stmt, FunctionLike
	FunctionStmt struct {
		ByRef          bool              // @var bool Whether function returns by reference
		Name           *Identifier       // @var Node\Identifier Name
		Params         []*Param          // @var Node\Param[] Parameters
		ReturnType     any               // @var null|Node\Identifier|Node\Name|Node\ComplexType Return type
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// StmtGlobal : Stmt
	GlobalStmt struct {
		Vars       []Expr // @var Node\Expr[] Variables
		Attributes any
	}

	// StmtGoto : Stmt
	GotoStmt struct {
		Name       *Identifier // @var Identifier Name of label to jump to
		Attributes any
	}

	// StmtGroupUse : Stmt
	GroupUseStmt struct {
		Type       int           // @var int Type of group use
		Prefix     *Name         // @var Name Prefix for uses
		Uses       []*UseUseStmt // @var UseUse[] Uses
		Attributes any
	}

	// StmtHaltCompiler : Stmt
	HaltCompilerStmt struct {
		Remaining  string // @var string Remaining text after halt compiler statement.
		Attributes any
	}

	// StmtIf : Stmt
	IfStmt struct {
		Cond       Expr          // @var Node\Expr Condition expression
		Stmts      []Stmt        // @var Node\Stmt[] Statements
		Elseifs    []*ElseIfStmt // @var ElseIf_[] Elseif clauses
		Else       *ElseStmt     // @var null|Else_ Else clause
		Attributes any
	}

	// StmtInlineHTML : Stmt
	InlineHTMLStmt struct {
		Value      string // @var string String
		Attributes any
	}

	// StmtInterface : StmtClassLike
	InterfaceStmt struct {
		Extends        []*Name           // @var Node\Name[] Extended interfaces
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// StmtLabel : Stmt
	LabelStmt struct {
		Name       *Identifier // @var Identifier Name
		Attributes any
	}

	// StmtNamespace : Stmt
	NamespaceStmt struct {
		Name       *Name  // @var null|Node\Name Name
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// StmtNop : Stmt
	NopStmt struct {
		Attributes any
	}

	// StmtProperty : Stmt
	PropertyStmt struct {
		Flags      int                     // @var int Modifiers
		Props      []*PropertyPropertyStmt // @var PropertyProperty[] Properties
		Type       any                     // @var null|Identifier|Name|ComplexType Type declaration
		AttrGroups []*AttributeGroup       // @var Node\AttributeGroup[] PHP attribute groups
		Attributes any
	}

	// StmtPropertyProperty : Stmt
	PropertyPropertyStmt struct {
		Name       *VarLikeIdentifier // @var Node\VarLikeIdentifier Name
		Default    Expr               // @var null|Node\Expr Default
		Attributes any
	}

	// StmtReturn : Stmt
	ReturnStmt struct {
		Expr       Expr // @var null|Node\Expr Expression
		Attributes any
	}

	// StmtStatic : Stmt
	StaticStmt struct {
		Vars       []*StaticVarStmt // @var StaticVar[] Variable definitions
		Attributes any
	}

	// StmtStaticVar : Stmt
	StaticVarStmt struct {
		Var        *VariableExpr // @var Expr\Variable Variable
		Default    Expr          // @var null|Node\Expr Default value
		Attributes any
	}

	// StmtSwitch : Stmt
	SwitchStmt struct {
		Cond       Expr        // @var Node\Expr Condition
		Cases      []*CaseStmt // @var Case_[] Case list
		Attributes any
	}

	// StmtThrow : Stmt
	ThrowStmt struct {
		Expr       Expr // @var Node\Expr Expression
		Attributes any
	}

	// StmtTrait : StmtClassLike
	TraitStmt struct {
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
		Attributes     any
	}

	// StmtTraitUse : Stmt
	TraitUseStmt struct {
		Traits      []*Name                   // @var Node\Name[] Traits
		Adaptations []*TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
		Attributes  any
	}

	// StmtTraitUseAdaptationAlias : StmtTraitUseAdaptation
	TraitUseAdaptationAliasStmt struct {
		NewModifier int         // @var null|int New modifier
		NewName     *Identifier // @var null|Node\Identifier New name
		Trait       *Name       // @var Node\Name|null Trait name
		Method      *Identifier // @var Node\Identifier Method name
		Attributes  any
	}

	// StmtTraitUseAdaptationPrecedence : StmtTraitUseAdaptation
	TraitUseAdaptationPrecedenceStmt struct {
		Insteadof  []*Name     // @var Node\Name[] Overwritten traits
		Trait      *Name       // @var Node\Name|null Trait name
		Method     *Identifier // @var Node\Identifier Method name
		Attributes any
	}

	// StmtTryCatch : Stmt
	TryCatchStmt struct {
		Stmts      []Stmt       // @var Node\Stmt[] Statements
		Catches    []*CatchStmt // @var Catch_[] Catches
		Finally    *FinallyStmt // @var null|Finally_ Optional finally node
		Attributes any
	}

	// StmtUnset : Stmt
	UnsetStmt struct {
		Vars       []Expr // @var Node\Expr[] Variables to unset
		Attributes any
	}

	// StmtUse : Stmt
	UseStmt struct {
		Type       int           // @var int Type of alias
		Uses       []*UseUseStmt // @var UseUse[] Aliases
		Attributes any
	}

	// StmtUseUse : Stmt
	UseUseStmt struct {
		Type       int         // @var int One of the Stmt\Use_::TYPE_* constants. Will only differ from TYPE_UNKNOWN for mixed group uses
		Name       *Name       // @var Node\Name Namespace, class, function or constant to alias
		Alias      *Identifier // @var Identifier|null Alias
		Attributes any
	}

	// StmtWhile : Stmt
	WhileStmt struct {
		Cond       Expr   // @var Node\Expr Condition
		Stmts      []Stmt // @var Node\Stmt[] Statements
		Attributes any
	}

	// UnionType : ComplexType
	UnionType struct {
		Types      []any // @var (Identifier|Name|IntersectionType)[] Types
		Attributes any
	}

	// VarLikeIdentifier : Identifier
	VarLikeIdentifier struct {
		Name       string // @var string Identifier as string
		Attributes any
	}

	// VariadicPlaceholder : PhpParserNodeAbstract
	VariadicPlaceholder struct {
		Attributes any
	}
)

// PhpParserNodeAbstract
func (*Arg) phpParserNodeAbstractNode()                 {}
func (*Attribute) phpParserNodeAbstractNode()           {}
func (*AttributeGroup) phpParserNodeAbstractNode()      {}
func (*Const) phpParserNodeAbstractNode()               {}
func (*Identifier) phpParserNodeAbstractNode()          {}
func (*MatchArm) phpParserNodeAbstractNode()            {}
func (*Name) phpParserNodeAbstractNode()                {}
func (*Param) phpParserNodeAbstractNode()               {}
func (*VariadicPlaceholder) phpParserNodeAbstractNode() {}

// Expr
func (*ArrayExpr) exprNode()                 {}
func (*ArrayDimFetchExpr) exprNode()         {}
func (*ArrayItemExpr) exprNode()             {}
func (*ArrowFunctionExpr) exprNode()         {}
func (*AssignExpr) exprNode()                {}
func (*AssignRefExpr) exprNode()             {}
func (*BitwiseNotExpr) exprNode()            {}
func (*BooleanNotExpr) exprNode()            {}
func (*ClassConstFetchExpr) exprNode()       {}
func (*CloneExpr) exprNode()                 {}
func (*ClosureExpr) exprNode()               {}
func (*ClosureUseExpr) exprNode()            {}
func (*ConstFetchExpr) exprNode()            {}
func (*EmptyExpr) exprNode()                 {}
func (*ErrorExpr) exprNode()                 {}
func (*ErrorSuppressExpr) exprNode()         {}
func (*EvalExpr) exprNode()                  {}
func (*ExitExpr) exprNode()                  {}
func (*IncludeExpr) exprNode()               {}
func (*InstanceofExpr) exprNode()            {}
func (*IssetExpr) exprNode()                 {}
func (*ListExpr) exprNode()                  {}
func (*MatchExpr) exprNode()                 {}
func (*NullsafePropertyFetchExpr) exprNode() {}
func (*PostDecExpr) exprNode()               {}
func (*PostIncExpr) exprNode()               {}
func (*PreDecExpr) exprNode()                {}
func (*PreIncExpr) exprNode()                {}
func (*PrintExpr) exprNode()                 {}
func (*PropertyFetchExpr) exprNode()         {}
func (*ShellExecExpr) exprNode()             {}
func (*StaticPropertyFetchExpr) exprNode()   {}
func (*TernaryExpr) exprNode()               {}
func (*ThrowExpr) exprNode()                 {}
func (*UnaryMinusExpr) exprNode()            {}
func (*UnaryPlusExpr) exprNode()             {}
func (*VariableExpr) exprNode()              {}
func (*YieldExpr) exprNode()                 {}
func (*YieldFromExpr) exprNode()             {}

// FunctionLike
func (*ArrowFunctionExpr) functionLikeNode() {}
func (*ClosureExpr) functionLikeNode()       {}
func (*ClassMethodStmt) functionLikeNode()   {}
func (*FunctionStmt) functionLikeNode()      {}

// ExprAssignOp
func (*AssignOpBitwiseAndExpr) exprAssignOpNode() {}
func (*AssignOpBitwiseOrExpr) exprAssignOpNode()  {}
func (*AssignOpBitwiseXorExpr) exprAssignOpNode() {}
func (*AssignOpCoalesceExpr) exprAssignOpNode()   {}
func (*AssignOpConcatExpr) exprAssignOpNode()     {}
func (*AssignOpDivExpr) exprAssignOpNode()        {}
func (*AssignOpMinusExpr) exprAssignOpNode()      {}
func (*AssignOpModExpr) exprAssignOpNode()        {}
func (*AssignOpMulExpr) exprAssignOpNode()        {}
func (*AssignOpPlusExpr) exprAssignOpNode()       {}
func (*AssignOpPowExpr) exprAssignOpNode()        {}
func (*AssignOpShiftLeftExpr) exprAssignOpNode()  {}
func (*AssignOpShiftRightExpr) exprAssignOpNode() {}

// ExprBinaryOp
func (*BinaryOpBitwiseAndExpr) exprBinaryOpNode()     {}
func (*BinaryOpBitwiseOrExpr) exprBinaryOpNode()      {}
func (*BinaryOpBitwiseXorExpr) exprBinaryOpNode()     {}
func (*BinaryOpBooleanAndExpr) exprBinaryOpNode()     {}
func (*BinaryOpBooleanOrExpr) exprBinaryOpNode()      {}
func (*BinaryOpCoalesceExpr) exprBinaryOpNode()       {}
func (*BinaryOpConcatExpr) exprBinaryOpNode()         {}
func (*BinaryOpDivExpr) exprBinaryOpNode()            {}
func (*BinaryOpEqualExpr) exprBinaryOpNode()          {}
func (*BinaryOpGreaterExpr) exprBinaryOpNode()        {}
func (*BinaryOpGreaterOrEqualExpr) exprBinaryOpNode() {}
func (*BinaryOpIdenticalExpr) exprBinaryOpNode()      {}
func (*BinaryOpLogicalAndExpr) exprBinaryOpNode()     {}
func (*BinaryOpLogicalOrExpr) exprBinaryOpNode()      {}
func (*BinaryOpLogicalXorExpr) exprBinaryOpNode()     {}
func (*BinaryOpMinusExpr) exprBinaryOpNode()          {}
func (*BinaryOpModExpr) exprBinaryOpNode()            {}
func (*BinaryOpMulExpr) exprBinaryOpNode()            {}
func (*BinaryOpNotEqualExpr) exprBinaryOpNode()       {}
func (*BinaryOpNotIdenticalExpr) exprBinaryOpNode()   {}
func (*BinaryOpPlusExpr) exprBinaryOpNode()           {}
func (*BinaryOpPowExpr) exprBinaryOpNode()            {}
func (*BinaryOpShiftLeftExpr) exprBinaryOpNode()      {}
func (*BinaryOpShiftRightExpr) exprBinaryOpNode()     {}
func (*BinaryOpSmallerExpr) exprBinaryOpNode()        {}
func (*BinaryOpSmallerOrEqualExpr) exprBinaryOpNode() {}
func (*BinaryOpSpaceshipExpr) exprBinaryOpNode()      {}

// ExprCast
func (*CastArrayExpr) exprCastNode()  {}
func (*CastBoolExpr) exprCastNode()   {}
func (*CastDoubleExpr) exprCastNode() {}
func (*CastIntExpr) exprCastNode()    {}
func (*CastObjectExpr) exprCastNode() {}
func (*CastStringExpr) exprCastNode() {}
func (*CastUnsetExpr) exprCastNode()  {}

// ExprCallLike
func (*FuncCallExpr) exprCallLikeNode()           {}
func (*MethodCallExpr) exprCallLikeNode()         {}
func (*NewExpr) exprCallLikeNode()                {}
func (*NullsafeMethodCallExpr) exprCallLikeNode() {}
func (*StaticCallExpr) exprCallLikeNode()         {}

// ComplexType
func (*IntersectionType) complexTypeNode() {}
func (*NullableType) complexTypeNode()     {}
func (*UnionType) complexTypeNode()        {}

// Name
func (*NameFullyQualified) nameNode() {}
func (*NameRelative) nameNode()       {}

// Scalar
func (*DNumberScalar) scalarNode()            {}
func (*EncapsedScalar) scalarNode()           {}
func (*EncapsedStringPartScalar) scalarNode() {}
func (*LNumberScalar) scalarNode()            {}
func (*StringScalar) scalarNode()             {}

// ScalarMagicConst
func (*MagicConstClassScalar) scalarMagicConstNode()     {}
func (*MagicConstDirScalar) scalarMagicConstNode()       {}
func (*MagicConstFileScalar) scalarMagicConstNode()      {}
func (*MagicConstFunctionScalar) scalarMagicConstNode()  {}
func (*MagicConstLineScalar) scalarMagicConstNode()      {}
func (*MagicConstMethodScalar) scalarMagicConstNode()    {}
func (*MagicConstNamespaceScalar) scalarMagicConstNode() {}
func (*MagicConstTraitScalar) scalarMagicConstNode()     {}

// Stmt
func (*BreakStmt) stmtNode()            {}
func (*CaseStmt) stmtNode()             {}
func (*CatchStmt) stmtNode()            {}
func (*ClassConstStmt) stmtNode()       {}
func (*ClassMethodStmt) stmtNode()      {}
func (*ConstStmt) stmtNode()            {}
func (*ContinueStmt) stmtNode()         {}
func (*DeclareStmt) stmtNode()          {}
func (*DeclareDeclareStmt) stmtNode()   {}
func (*DoStmt) stmtNode()               {}
func (*EchoStmt) stmtNode()             {}
func (*ElseStmt) stmtNode()             {}
func (*ElseIfStmt) stmtNode()           {}
func (*EnumCaseStmt) stmtNode()         {}
func (*ExpressionStmt) stmtNode()       {}
func (*FinallyStmt) stmtNode()          {}
func (*ForStmt) stmtNode()              {}
func (*ForeachStmt) stmtNode()          {}
func (*FunctionStmt) stmtNode()         {}
func (*GlobalStmt) stmtNode()           {}
func (*GotoStmt) stmtNode()             {}
func (*GroupUseStmt) stmtNode()         {}
func (*HaltCompilerStmt) stmtNode()     {}
func (*IfStmt) stmtNode()               {}
func (*InlineHTMLStmt) stmtNode()       {}
func (*LabelStmt) stmtNode()            {}
func (*NamespaceStmt) stmtNode()        {}
func (*NopStmt) stmtNode()              {}
func (*PropertyStmt) stmtNode()         {}
func (*PropertyPropertyStmt) stmtNode() {}
func (*ReturnStmt) stmtNode()           {}
func (*StaticStmt) stmtNode()           {}
func (*StaticVarStmt) stmtNode()        {}
func (*SwitchStmt) stmtNode()           {}
func (*ThrowStmt) stmtNode()            {}
func (*TraitUseStmt) stmtNode()         {}
func (*TryCatchStmt) stmtNode()         {}
func (*UnsetStmt) stmtNode()            {}
func (*UseStmt) stmtNode()              {}
func (*UseUseStmt) stmtNode()           {}
func (*WhileStmt) stmtNode()            {}

// StmtClassLike
func (*ClassStmt) stmtClassLikeNode()     {}
func (*EnumStmt) stmtClassLikeNode()      {}
func (*InterfaceStmt) stmtClassLikeNode() {}
func (*TraitStmt) stmtClassLikeNode()     {}

// StmtTraitUseAdaptation
func (*TraitUseAdaptationAliasStmt) stmtTraitUseAdaptationNode()      {}
func (*TraitUseAdaptationPrecedenceStmt) stmtTraitUseAdaptationNode() {}

// Identifier
func (*VarLikeIdentifier) identifierNode() {}

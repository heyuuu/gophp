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
		Name   *Identifier // @var Identifier|null Parameter name (for named parameters)
		Value  Expr        // @var Expr Value to pass
		ByRef  bool        // @var bool Whether to pass by ref
		Unpack bool        // @var bool Whether to unpack the argument
	}

	// Attribute : PhpParserNodeAbstract
	Attribute struct {
		Name *Name  // @var Name Attribute name
		Args []*Arg // @var Arg[] Attribute arguments
	}

	// AttributeGroup : PhpParserNodeAbstract
	AttributeGroup struct {
		Attrs []*Attribute // @var Attribute[] Attributes
	}

	// Const : PhpParserNodeAbstract
	Const struct {
		Name           *Identifier // @var Identifier Name
		Value          Expr        // @var Expr Value
		NamespacedName *Name       // @var Name|null Namespaced name (if using NameResolver)
	}

	// ExprArray : Expr
	ArrayExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	}

	// ExprArrayDimFetch : Expr
	ArrayDimFetchExpr struct {
		Var Expr // @var Expr Variable
		Dim Expr // @var null|Expr Array index / dim
	}

	// ExprArrayItem : Expr
	ArrayItemExpr struct {
		Key    Expr // @var null|Expr Key
		Value  Expr // @var Expr Value
		ByRef  bool // @var bool Whether to assign by reference
		Unpack bool // @var bool Whether to unpack the argument
	}

	// ExprArrowFunction : Expr, FunctionLike
	ArrowFunctionExpr struct {
		Static     bool              // @var bool
		ByRef      bool              // @var bool
		Params     []*Param          // @var Node\Param[]
		ReturnType any               // @var null|Node\Identifier|Node\Name|Node\ComplexType
		Expr       Expr              // @var Expr
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
	}

	// ExprAssign : Expr
	AssignExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpBitwiseAnd : ExprAssignOp
	AssignOpBitwiseAndExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpBitwiseOr : ExprAssignOp
	AssignOpBitwiseOrExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpBitwiseXor : ExprAssignOp
	AssignOpBitwiseXorExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpCoalesce : ExprAssignOp
	AssignOpCoalesceExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpConcat : ExprAssignOp
	AssignOpConcatExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpDiv : ExprAssignOp
	AssignOpDivExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpMinus : ExprAssignOp
	AssignOpMinusExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpMod : ExprAssignOp
	AssignOpModExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpMul : ExprAssignOp
	AssignOpMulExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpPlus : ExprAssignOp
	AssignOpPlusExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpPow : ExprAssignOp
	AssignOpPowExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpShiftLeft : ExprAssignOp
	AssignOpShiftLeftExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignOpShiftRight : ExprAssignOp
	AssignOpShiftRightExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// ExprAssignRef : Expr
	AssignRefExpr struct {
		Var  Expr // @var Expr Variable reference is assigned to
		Expr Expr // @var Expr Variable which is referenced
	}

	// ExprBinaryOpBitwiseAnd : ExprBinaryOp
	BinaryOpBitwiseAndExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpBitwiseOr : ExprBinaryOp
	BinaryOpBitwiseOrExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpBitwiseXor : ExprBinaryOp
	BinaryOpBitwiseXorExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpBooleanAnd : ExprBinaryOp
	BinaryOpBooleanAndExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpBooleanOr : ExprBinaryOp
	BinaryOpBooleanOrExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpCoalesce : ExprBinaryOp
	BinaryOpCoalesceExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpConcat : ExprBinaryOp
	BinaryOpConcatExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpDiv : ExprBinaryOp
	BinaryOpDivExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpEqual : ExprBinaryOp
	BinaryOpEqualExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpGreater : ExprBinaryOp
	BinaryOpGreaterExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpGreaterOrEqual : ExprBinaryOp
	BinaryOpGreaterOrEqualExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpIdentical : ExprBinaryOp
	BinaryOpIdenticalExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpLogicalAnd : ExprBinaryOp
	BinaryOpLogicalAndExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpLogicalOr : ExprBinaryOp
	BinaryOpLogicalOrExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpLogicalXor : ExprBinaryOp
	BinaryOpLogicalXorExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpMinus : ExprBinaryOp
	BinaryOpMinusExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpMod : ExprBinaryOp
	BinaryOpModExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpMul : ExprBinaryOp
	BinaryOpMulExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpNotEqual : ExprBinaryOp
	BinaryOpNotEqualExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpNotIdentical : ExprBinaryOp
	BinaryOpNotIdenticalExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpPlus : ExprBinaryOp
	BinaryOpPlusExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpPow : ExprBinaryOp
	BinaryOpPowExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpShiftLeft : ExprBinaryOp
	BinaryOpShiftLeftExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpShiftRight : ExprBinaryOp
	BinaryOpShiftRightExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpSmaller : ExprBinaryOp
	BinaryOpSmallerExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpSmallerOrEqual : ExprBinaryOp
	BinaryOpSmallerOrEqualExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBinaryOpSpaceship : ExprBinaryOp
	BinaryOpSpaceshipExpr struct {
		Left  Expr // @var Expr The left hand side expression
		Right Expr // @var Expr The right hand side expression
	}

	// ExprBitwiseNot : Expr
	BitwiseNotExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprBooleanNot : Expr
	BooleanNotExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastArray : ExprCast
	CastArrayExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastBool : ExprCast
	CastBoolExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastDouble : ExprCast
	CastDoubleExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastInt : ExprCast
	CastIntExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastObject : ExprCast
	CastObjectExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastString : ExprCast
	CastStringExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprCastUnset : ExprCast
	CastUnsetExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprClassConstFetch : Expr
	ClassConstFetchExpr struct {
		Class any // @var Name|Expr Class name
		Name  any // @var Identifier|Error Constant name
	}

	// ExprClone : Expr
	CloneExpr struct {
		Expr Expr // @var Expr Expression
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
	}

	// ExprClosureUse : Expr
	ClosureUseExpr struct {
		Var   *VariableExpr // @var Expr\Variable Variable to use
		ByRef bool          // @var bool Whether to use by reference
	}

	// ExprConstFetch : Expr
	ConstFetchExpr struct {
		Name *Name // @var Name Constant name
	}

	// ExprEmpty : Expr
	EmptyExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprError : Expr
	ErrorExpr struct {
	}

	// ExprErrorSuppress : Expr
	ErrorSuppressExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprEval : Expr
	EvalExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprExit : Expr
	ExitExpr struct {
		Expr Expr // @var null|Expr Expression
	}

	// ExprFuncCall : ExprCallLike
	FuncCallExpr struct {
		Name any   // @var Node\Name|Expr Function name
		Args []any // @var array<Node\Arg|Node\VariadicPlaceholder> Arguments
	}

	// ExprInclude : Expr
	IncludeExpr struct {
		Expr Expr // @var Expr Expression
		Type int  // @var int Type of include
	}

	// ExprInstanceof : Expr
	InstanceofExpr struct {
		Expr  Expr // @var Expr Expression
		Class any  // @var Name|Expr Class name
	}

	// ExprIsset : Expr
	IssetExpr struct {
		Vars []Expr // @var Expr[] Variables
	}

	// ExprList : Expr
	ListExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	}

	// ExprMatch : Expr
	MatchExpr struct {
		Cond Expr        // @var Node\Expr
		Arms []*MatchArm // @var MatchArm[]
	}

	// ExprMethodCall : ExprCallLike
	MethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Identifier|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNew : ExprCallLike
	NewExpr struct {
		Class any   // @var Node\Name|Expr|Node\Stmt\Class_ Class name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNullsafeMethodCall : ExprCallLike
	NullsafeMethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Identifier|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNullsafePropertyFetch : Expr
	NullsafePropertyFetchExpr struct {
		Var  Expr // @var Expr Variable holding object
		Name any  // @var Identifier|Expr Property name
	}

	// ExprPostDec : Expr
	PostDecExpr struct {
		Var Expr // @var Expr Variable
	}

	// ExprPostInc : Expr
	PostIncExpr struct {
		Var Expr // @var Expr Variable
	}

	// ExprPreDec : Expr
	PreDecExpr struct {
		Var Expr // @var Expr Variable
	}

	// ExprPreInc : Expr
	PreIncExpr struct {
		Var Expr // @var Expr Variable
	}

	// ExprPrint : Expr
	PrintExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprPropertyFetch : Expr
	PropertyFetchExpr struct {
		Var  Expr // @var Expr Variable holding object
		Name any  // @var Identifier|Expr Property name
	}

	// ExprShellExec : Expr
	ShellExecExpr struct {
		Parts []any // @var array Encapsed string array
	}

	// ExprStaticCall : ExprCallLike
	StaticCallExpr struct {
		Class any   // @var Node\Name|Expr Class name
		Name  any   // @var Identifier|Expr Method name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprStaticPropertyFetch : Expr
	StaticPropertyFetchExpr struct {
		Class any // @var Name|Expr Class name
		Name  any // @var VarLikeIdentifier|Expr Property name
	}

	// ExprTernary : Expr
	TernaryExpr struct {
		Cond Expr // @var Expr Condition
		If   Expr // @var null|Expr Expression for true
		Else Expr // @var Expr Expression for false
	}

	// ExprThrow : Expr
	ThrowExpr struct {
		Expr Expr // @var Node\Expr Expression
	}

	// ExprUnaryMinus : Expr
	UnaryMinusExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprUnaryPlus : Expr
	UnaryPlusExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprVariable : Expr
	VariableExpr struct {
		Name any // @var string|Expr Name
	}

	// ExprYield : Expr
	YieldExpr struct {
		Key   Expr // @var null|Expr Key expression
		Value Expr // @var null|Expr Value expression
	}

	// ExprYieldFrom : Expr
	YieldFromExpr struct {
		Expr Expr // @var Expr Expression to yield from
	}

	// Identifier : PhpParserNodeAbstract
	Identifier struct {
		Name string // @var string Identifier as string
	}

	// IntersectionType : ComplexType
	IntersectionType struct {
		Types []any // @var (Identifier|Name)[] Types
	}

	// MatchArm : PhpParserNodeAbstract
	MatchArm struct {
		Conds []Expr // @var null|Node\Expr[]
		Body  Expr   // @var Node\Expr
	}

	// Name : PhpParserNodeAbstract
	Name struct {
		Parts []string // @var string[] Parts of the name
	}

	// NameFullyQualified : Name
	NameFullyQualified struct {
		Parts []string // @var string[] Parts of the name
	}

	// NameRelative : Name
	NameRelative struct {
		Parts []string // @var string[] Parts of the name
	}

	// NullableType : ComplexType
	NullableType struct {
		Type any // @var Identifier|Name Type
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
	}

	// ScalarDNumber : Scalar
	DNumberScalar struct {
		Value float64 // @var float Number value
	}

	// ScalarEncapsed : Scalar
	EncapsedScalar struct {
		Parts []Expr // @var Expr[] list of string parts
	}

	// ScalarEncapsedStringPart : Scalar
	EncapsedStringPartScalar struct {
		Value string // @var string String value
	}

	// ScalarLNumber : Scalar
	LNumberScalar struct {
		Value int // @var int Number value
	}

	// ScalarMagicConstClass : ScalarMagicConst
	MagicConstClassScalar struct {
	}

	// ScalarMagicConstDir : ScalarMagicConst
	MagicConstDirScalar struct {
	}

	// ScalarMagicConstFile : ScalarMagicConst
	MagicConstFileScalar struct {
	}

	// ScalarMagicConstFunction : ScalarMagicConst
	MagicConstFunctionScalar struct {
	}

	// ScalarMagicConstLine : ScalarMagicConst
	MagicConstLineScalar struct {
	}

	// ScalarMagicConstMethod : ScalarMagicConst
	MagicConstMethodScalar struct {
	}

	// ScalarMagicConstNamespace : ScalarMagicConst
	MagicConstNamespaceScalar struct {
	}

	// ScalarMagicConstTrait : ScalarMagicConst
	MagicConstTraitScalar struct {
	}

	// ScalarString : Scalar
	StringScalar struct {
		Value string // @var string String value
	}

	// StmtBreak : Stmt
	BreakStmt struct {
		Num Expr // @var null|Node\Expr Number of loops to break
	}

	// StmtCase : Stmt
	CaseStmt struct {
		Cond  Expr   // @var null|Node\Expr Condition (null for default)
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtCatch : Stmt
	CatchStmt struct {
		Types []*Name       // @var Node\Name[] Types of exceptions to catch
		Var   *VariableExpr // @var Expr\Variable|null Variable for exception
		Stmts []Stmt        // @var Node\Stmt[] Statements
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
	}

	// StmtClassConst : Stmt
	ClassConstStmt struct {
		Flags      int               // @var int Modifiers
		Consts     []*Const          // @var Node\Const_[] Constant declarations
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
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
	}

	// StmtConst : Stmt
	ConstStmt struct {
		Consts []*Const // @var Node\Const_[] Constant declarations
	}

	// StmtContinue : Stmt
	ContinueStmt struct {
		Num Expr // @var null|Node\Expr Number of loops to continue
	}

	// StmtDeclare : Stmt
	DeclareStmt struct {
		Declares []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
		Stmts    []Stmt                // @var Node\Stmt[]|null Statements
	}

	// StmtDeclareDeclare : Stmt
	DeclareDeclareStmt struct {
		Key   *Identifier // @var Node\Identifier Key
		Value Expr        // @var Node\Expr Value
	}

	// StmtDo : Stmt
	DoStmt struct {
		Stmts []Stmt // @var Node\Stmt[] Statements
		Cond  Expr   // @var Node\Expr Condition
	}

	// StmtEcho : Stmt
	EchoStmt struct {
		Exprs []Expr // @var Node\Expr[] Expressions
	}

	// StmtElse : Stmt
	ElseStmt struct {
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtElseIf : Stmt
	ElseIfStmt struct {
		Cond  Expr   // @var Node\Expr Condition
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtEnum : StmtClassLike
	EnumStmt struct {
		ScalarType     *Identifier       // @var null|Node\Identifier Scalar Type
		Implements     []*Name           // @var Node\Name[] Names of implemented interfaces
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	}

	// StmtEnumCase : Stmt
	EnumCaseStmt struct {
		Name       *Identifier       // @var Node\Identifier Enum case name
		Expr       Expr              // @var Node\Expr|null Enum case expression
		AttrGroups []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
	}

	// StmtExpression : Stmt
	ExpressionStmt struct {
		Expr Expr // @var Node\Expr Expression
	}

	// StmtFinally : Stmt
	FinallyStmt struct {
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtFor : Stmt
	ForStmt struct {
		Init  []Expr // @var Node\Expr[] Init expressions
		Cond  []Expr // @var Node\Expr[] Loop conditions
		Loop  []Expr // @var Node\Expr[] Loop expressions
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtForeach : Stmt
	ForeachStmt struct {
		Expr     Expr   // @var Node\Expr Expression to iterate
		KeyVar   Expr   // @var null|Node\Expr Variable to assign key to
		ByRef    bool   // @var bool Whether to assign value by reference
		ValueVar Expr   // @var Node\Expr Variable to assign value to
		Stmts    []Stmt // @var Node\Stmt[] Statements
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
	}

	// StmtGlobal : Stmt
	GlobalStmt struct {
		Vars []Expr // @var Node\Expr[] Variables
	}

	// StmtGoto : Stmt
	GotoStmt struct {
		Name *Identifier // @var Identifier Name of label to jump to
	}

	// StmtGroupUse : Stmt
	GroupUseStmt struct {
		Type   int           // @var int Type of group use
		Prefix *Name         // @var Name Prefix for uses
		Uses   []*UseUseStmt // @var UseUse[] Uses
	}

	// StmtHaltCompiler : Stmt
	HaltCompilerStmt struct {
		Remaining string // @var string Remaining text after halt compiler statement.
	}

	// StmtIf : Stmt
	IfStmt struct {
		Cond    Expr          // @var Node\Expr Condition expression
		Stmts   []Stmt        // @var Node\Stmt[] Statements
		Elseifs []*ElseIfStmt // @var ElseIf_[] Elseif clauses
		Else    *ElseStmt     // @var null|Else_ Else clause
	}

	// StmtInlineHTML : Stmt
	InlineHTMLStmt struct {
		Value string // @var string String
	}

	// StmtInterface : StmtClassLike
	InterfaceStmt struct {
		Extends        []*Name           // @var Node\Name[] Extended interfaces
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	}

	// StmtLabel : Stmt
	LabelStmt struct {
		Name *Identifier // @var Identifier Name
	}

	// StmtNamespace : Stmt
	NamespaceStmt struct {
		Name  *Name  // @var null|Node\Name Name
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// StmtNop : Stmt
	NopStmt struct {
	}

	// StmtProperty : Stmt
	PropertyStmt struct {
		Flags      int                     // @var int Modifiers
		Props      []*PropertyPropertyStmt // @var PropertyProperty[] Properties
		Type       any                     // @var null|Identifier|Name|ComplexType Type declaration
		AttrGroups []*AttributeGroup       // @var Node\AttributeGroup[] PHP attribute groups
	}

	// StmtPropertyProperty : Stmt
	PropertyPropertyStmt struct {
		Name    *VarLikeIdentifier // @var Node\VarLikeIdentifier Name
		Default Expr               // @var null|Node\Expr Default
	}

	// StmtReturn : Stmt
	ReturnStmt struct {
		Expr Expr // @var null|Node\Expr Expression
	}

	// StmtStatic : Stmt
	StaticStmt struct {
		Vars []*StaticVarStmt // @var StaticVar[] Variable definitions
	}

	// StmtStaticVar : Stmt
	StaticVarStmt struct {
		Var     *VariableExpr // @var Expr\Variable Variable
		Default Expr          // @var null|Node\Expr Default value
	}

	// StmtSwitch : Stmt
	SwitchStmt struct {
		Cond  Expr        // @var Node\Expr Condition
		Cases []*CaseStmt // @var Case_[] Case list
	}

	// StmtThrow : Stmt
	ThrowStmt struct {
		Expr Expr // @var Node\Expr Expression
	}

	// StmtTrait : StmtClassLike
	TraitStmt struct {
		Name           *Identifier       // @var Node\Identifier|null Name
		Stmts          []Stmt            // @var Node\Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	}

	// StmtTraitUse : Stmt
	TraitUseStmt struct {
		Traits      []*Name                   // @var Node\Name[] Traits
		Adaptations []*TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
	}

	// StmtTraitUseAdaptationAlias : StmtTraitUseAdaptation
	TraitUseAdaptationAliasStmt struct {
		NewModifier int         // @var null|int New modifier
		NewName     *Identifier // @var null|Node\Identifier New name
		Trait       *Name       // @var Node\Name|null Trait name
		Method      *Identifier // @var Node\Identifier Method name
	}

	// StmtTraitUseAdaptationPrecedence : StmtTraitUseAdaptation
	TraitUseAdaptationPrecedenceStmt struct {
		Insteadof []*Name     // @var Node\Name[] Overwritten traits
		Trait     *Name       // @var Node\Name|null Trait name
		Method    *Identifier // @var Node\Identifier Method name
	}

	// StmtTryCatch : Stmt
	TryCatchStmt struct {
		Stmts   []Stmt       // @var Node\Stmt[] Statements
		Catches []*CatchStmt // @var Catch_[] Catches
		Finally *FinallyStmt // @var null|Finally_ Optional finally node
	}

	// StmtUnset : Stmt
	UnsetStmt struct {
		Vars []Expr // @var Node\Expr[] Variables to unset
	}

	// StmtUse : Stmt
	UseStmt struct {
		Type int           // @var int Type of alias
		Uses []*UseUseStmt // @var UseUse[] Aliases
	}

	// StmtUseUse : Stmt
	UseUseStmt struct {
		Type  int         // @var int One of the Stmt\Use_::TYPE_* constants. Will only differ from TYPE_UNKNOWN for mixed group uses
		Name  *Name       // @var Node\Name Namespace, class, function or constant to alias
		Alias *Identifier // @var Identifier|null Alias
	}

	// StmtWhile : Stmt
	WhileStmt struct {
		Cond  Expr   // @var Node\Expr Condition
		Stmts []Stmt // @var Node\Stmt[] Statements
	}

	// UnionType : ComplexType
	UnionType struct {
		Types []any // @var (Identifier|Name|IntersectionType)[] Types
	}

	// VarLikeIdentifier : Identifier
	VarLikeIdentifier struct {
		Name string // @var string Identifier as string
	}

	// VariadicPlaceholder : PhpParserNodeAbstract
	VariadicPlaceholder struct {
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

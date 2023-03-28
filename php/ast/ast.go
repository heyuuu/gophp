package ast

type Node interface{}

// ComplexType : PhpParserNodeAbstract
type ComplexType interface {
}

// Expr : PhpParserNodeAbstract
type Expr interface {
}

// ExprAssignOp : Expr
type AssignOpExpr interface {
}

// ExprBinaryOp : Expr
type BinaryOpExpr interface {
}

// ExprCallLike : Expr
type CallLikeExpr interface {
}

// ExprCast : Expr
type CastExpr interface {
}

// FunctionLike
type FunctionLike interface {
}

// Scalar : Expr
type Scalar interface {
}

// ScalarMagicConst : Scalar
type MagicConstScalar interface {
}

// Stmt : PhpParserNodeAbstract
type Stmt interface {
}

// StmtClassLike : Stmt
type ClassLikeStmt interface {
}

// StmtTraitUseAdaptation : Stmt
type TraitUseAdaptationStmt interface {
}

// Arg : PhpParserNodeAbstract
type Arg struct {
	Name       *Identifier // @var Identifier|null Parameter name (for named parameters)
	Value      Expr        // @var Expr Value to pass
	ByRef      bool        // @var bool Whether to pass by ref
	Unpack     bool        // @var bool Whether to unpack the argument
	Attributes any
}

// Attribute : PhpParserNodeAbstract
type Attribute struct {
	Name       *Name  // @var Name Attribute name
	Args       []*Arg // @var Arg[] Attribute arguments
	Attributes any
}

// AttributeGroup : PhpParserNodeAbstract
type AttributeGroup struct {
	Attrs      []*Attribute // @var Attribute[] Attributes
	Attributes any
}

// Const : PhpParserNodeAbstract
type Const struct {
	Name           *Identifier // @var Identifier Name
	Value          Expr        // @var Expr Value
	NamespacedName *Name       // @var Name|null Namespaced name (if using NameResolver)
	Attributes     any
}

// ExprArray : Expr
type ArrayExpr struct {
	Items      []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	Attributes any
}

// ExprArrayDimFetch : Expr
type ArrayDimFetchExpr struct {
	Var        Expr // @var Expr Variable
	Dim        Expr // @var null|Expr Array index / dim
	Attributes any
}

// ExprArrayItem : Expr
type ArrayItemExpr struct {
	Key        Expr // @var null|Expr Key
	Value      Expr // @var Expr Value
	ByRef      bool // @var bool Whether to assign by reference
	Unpack     bool // @var bool Whether to unpack the argument
	Attributes any
}

// ExprArrowFunction : Expr, FunctionLike
type ArrowFunctionExpr struct {
	Static     bool              // @var bool
	ByRef      bool              // @var bool
	Params     []*Param          // @var Node\Param[]
	ReturnType any               // @var null|Node\Identifier|Node\Name|Node\ComplexType
	Expr       Expr              // @var Expr
	AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
	Attributes any
}

// ExprAssign : Expr
type AssignExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpBitwiseAnd : ExprAssignOp
type AssignOpBitwiseAndExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpBitwiseOr : ExprAssignOp
type AssignOpBitwiseOrExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpBitwiseXor : ExprAssignOp
type AssignOpBitwiseXorExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpCoalesce : ExprAssignOp
type AssignOpCoalesceExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpConcat : ExprAssignOp
type AssignOpConcatExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpDiv : ExprAssignOp
type AssignOpDivExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpMinus : ExprAssignOp
type AssignOpMinusExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpMod : ExprAssignOp
type AssignOpModExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpMul : ExprAssignOp
type AssignOpMulExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpPlus : ExprAssignOp
type AssignOpPlusExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpPow : ExprAssignOp
type AssignOpPowExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpShiftLeft : ExprAssignOp
type AssignOpShiftLeftExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignOpShiftRight : ExprAssignOp
type AssignOpShiftRightExpr struct {
	Var        Expr // @var Expr Variable
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprAssignRef : Expr
type AssignRefExpr struct {
	Var        Expr // @var Expr Variable reference is assigned to
	Expr       Expr // @var Expr Variable which is referenced
	Attributes any
}

// ExprBinaryOpBitwiseAnd : ExprBinaryOp
type BinaryOpBitwiseAndExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpBitwiseOr : ExprBinaryOp
type BinaryOpBitwiseOrExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpBitwiseXor : ExprBinaryOp
type BinaryOpBitwiseXorExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpBooleanAnd : ExprBinaryOp
type BinaryOpBooleanAndExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpBooleanOr : ExprBinaryOp
type BinaryOpBooleanOrExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpCoalesce : ExprBinaryOp
type BinaryOpCoalesceExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpConcat : ExprBinaryOp
type BinaryOpConcatExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpDiv : ExprBinaryOp
type BinaryOpDivExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpEqual : ExprBinaryOp
type BinaryOpEqualExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpGreater : ExprBinaryOp
type BinaryOpGreaterExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpGreaterOrEqual : ExprBinaryOp
type BinaryOpGreaterOrEqualExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpIdentical : ExprBinaryOp
type BinaryOpIdenticalExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpLogicalAnd : ExprBinaryOp
type BinaryOpLogicalAndExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpLogicalOr : ExprBinaryOp
type BinaryOpLogicalOrExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpLogicalXor : ExprBinaryOp
type BinaryOpLogicalXorExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpMinus : ExprBinaryOp
type BinaryOpMinusExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpMod : ExprBinaryOp
type BinaryOpModExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpMul : ExprBinaryOp
type BinaryOpMulExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpNotEqual : ExprBinaryOp
type BinaryOpNotEqualExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpNotIdentical : ExprBinaryOp
type BinaryOpNotIdenticalExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpPlus : ExprBinaryOp
type BinaryOpPlusExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpPow : ExprBinaryOp
type BinaryOpPowExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpShiftLeft : ExprBinaryOp
type BinaryOpShiftLeftExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpShiftRight : ExprBinaryOp
type BinaryOpShiftRightExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpSmaller : ExprBinaryOp
type BinaryOpSmallerExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpSmallerOrEqual : ExprBinaryOp
type BinaryOpSmallerOrEqualExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBinaryOpSpaceship : ExprBinaryOp
type BinaryOpSpaceshipExpr struct {
	Left       Expr // @var Expr The left hand side expression
	Right      Expr // @var Expr The right hand side expression
	Attributes any
}

// ExprBitwiseNot : Expr
type BitwiseNotExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprBooleanNot : Expr
type BooleanNotExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastArray : ExprCast
type CastArrayExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastBool : ExprCast
type CastBoolExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastDouble : ExprCast
type CastDoubleExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastInt : ExprCast
type CastIntExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastObject : ExprCast
type CastObjectExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastString : ExprCast
type CastStringExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprCastUnset : ExprCast
type CastUnsetExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprClassConstFetch : Expr
type ClassConstFetchExpr struct {
	Class      any // @var Name|Expr Class name
	Name       any // @var Identifier|Error Constant name
	Attributes any
}

// ExprClone : Expr
type CloneExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprClosure : Expr, FunctionLike
type ClosureExpr struct {
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
type ClosureUseExpr struct {
	Var        *VariableExpr // @var Expr\Variable Variable to use
	ByRef      bool          // @var bool Whether to use by reference
	Attributes any
}

// ExprConstFetch : Expr
type ConstFetchExpr struct {
	Name       *Name // @var Name Constant name
	Attributes any
}

// ExprEmpty : Expr
type EmptyExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprError : Expr
type ErrorExpr struct {
	Attributes any
}

// ExprErrorSuppress : Expr
type ErrorSuppressExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprEval : Expr
type EvalExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprExit : Expr
type ExitExpr struct {
	Expr       Expr // @var null|Expr Expression
	Attributes any
}

// ExprFuncCall : ExprCallLike
type FuncCallExpr struct {
	Name       any   // @var Node\Name|Expr Function name
	Args       []any // @var array<Node\Arg|Node\VariadicPlaceholder> Arguments
	Attributes any
}

// ExprInclude : Expr
type IncludeExpr struct {
	Expr       Expr // @var Expr Expression
	Type       int  // @var int Type of include
	Attributes any
}

// ExprInstanceof : Expr
type InstanceofExpr struct {
	Expr       Expr // @var Expr Expression
	Class      any  // @var Name|Expr Class name
	Attributes any
}

// ExprIsset : Expr
type IssetExpr struct {
	Vars       []Expr // @var Expr[] Variables
	Attributes any
}

// ExprList : Expr
type ListExpr struct {
	Items      []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	Attributes any
}

// ExprMatch : Expr
type MatchExpr struct {
	Cond       Expr        // @var Node\Expr
	Arms       []*MatchArm // @var MatchArm[]
	Attributes any
}

// ExprMethodCall : ExprCallLike
type MethodCallExpr struct {
	Var        Expr  // @var Expr Variable holding object
	Name       any   // @var Identifier|Expr Method name
	Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
	Attributes any
}

// ExprNew : ExprCallLike
type NewExpr struct {
	Class      any   // @var Node\Name|Expr|Node\Stmt\Class_ Class name
	Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
	Attributes any
}

// ExprNullsafeMethodCall : ExprCallLike
type NullsafeMethodCallExpr struct {
	Var        Expr  // @var Expr Variable holding object
	Name       any   // @var Identifier|Expr Method name
	Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
	Attributes any
}

// ExprNullsafePropertyFetch : Expr
type NullsafePropertyFetchExpr struct {
	Var        Expr // @var Expr Variable holding object
	Name       any  // @var Identifier|Expr Property name
	Attributes any
}

// ExprPostDec : Expr
type PostDecExpr struct {
	Var        Expr // @var Expr Variable
	Attributes any
}

// ExprPostInc : Expr
type PostIncExpr struct {
	Var        Expr // @var Expr Variable
	Attributes any
}

// ExprPreDec : Expr
type PreDecExpr struct {
	Var        Expr // @var Expr Variable
	Attributes any
}

// ExprPreInc : Expr
type PreIncExpr struct {
	Var        Expr // @var Expr Variable
	Attributes any
}

// ExprPrint : Expr
type PrintExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprPropertyFetch : Expr
type PropertyFetchExpr struct {
	Var        Expr // @var Expr Variable holding object
	Name       any  // @var Identifier|Expr Property name
	Attributes any
}

// ExprShellExec : Expr
type ShellExecExpr struct {
	Parts      []any // @var array Encapsed string array
	Attributes any
}

// ExprStaticCall : ExprCallLike
type StaticCallExpr struct {
	Class      any   // @var Node\Name|Expr Class name
	Name       any   // @var Identifier|Expr Method name
	Args       []any // @var array<Arg|VariadicPlaceholder> Arguments
	Attributes any
}

// ExprStaticPropertyFetch : Expr
type StaticPropertyFetchExpr struct {
	Class      any // @var Name|Expr Class name
	Name       any // @var VarLikeIdentifier|Expr Property name
	Attributes any
}

// ExprTernary : Expr
type TernaryExpr struct {
	Cond       Expr // @var Expr Condition
	If         Expr // @var null|Expr Expression for true
	Else       Expr // @var Expr Expression for false
	Attributes any
}

// ExprThrow : Expr
type ThrowExpr struct {
	Expr       Expr // @var Node\Expr Expression
	Attributes any
}

// ExprUnaryMinus : Expr
type UnaryMinusExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprUnaryPlus : Expr
type UnaryPlusExpr struct {
	Expr       Expr // @var Expr Expression
	Attributes any
}

// ExprVariable : Expr
type VariableExpr struct {
	Name       any // @var string|Expr Name
	Attributes any
}

// ExprYield : Expr
type YieldExpr struct {
	Key        Expr // @var null|Expr Key expression
	Value      Expr // @var null|Expr Value expression
	Attributes any
}

// ExprYieldFrom : Expr
type YieldFromExpr struct {
	Expr       Expr // @var Expr Expression to yield from
	Attributes any
}

// Identifier : PhpParserNodeAbstract, Stringable
type Identifier struct {
	Name              string // @var string Identifier as string
	SpecialClassNames any
	Attributes        any
}

// IntersectionType : ComplexType
type IntersectionType struct {
	Types      []any // @var (Identifier|Name)[] Types
	Attributes any
}

// MatchArm : PhpParserNodeAbstract
type MatchArm struct {
	Conds      []Expr // @var null|Node\Expr[]
	Body       Expr   // @var Node\Expr
	Attributes any
}

// Name : PhpParserNodeAbstract, Stringable
type Name struct {
	Parts             []string // @var string[] Parts of the name
	SpecialClassNames any
	Attributes        any
}

// NameFullyQualified : Name, Stringable
type NameFullyQualified struct {
	Parts      []string // @var string[] Parts of the name
	Attributes any
}

// NameRelative : Name, Stringable
type NameRelative struct {
	Parts      []string // @var string[] Parts of the name
	Attributes any
}

// NullableType : ComplexType
type NullableType struct {
	Type       any // @var Identifier|Name Type
	Attributes any
}

// Param : PhpParserNodeAbstract
type Param struct {
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
type DNumberScalar struct {
	Value      float64 // @var float Number value
	Attributes any
}

// ScalarEncapsed : Scalar
type EncapsedScalar struct {
	Parts      []Expr // @var Expr[] list of string parts
	Attributes any
}

// ScalarEncapsedStringPart : Scalar
type EncapsedStringPartScalar struct {
	Value      string // @var string String value
	Attributes any
}

// ScalarLNumber : Scalar
type LNumberScalar struct {
	Value      int // @var int Number value
	Attributes any
}

// ScalarMagicConstClass : ScalarMagicConst
type MagicConstClassScalar struct {
	Attributes any
}

// ScalarMagicConstDir : ScalarMagicConst
type MagicConstDirScalar struct {
	Attributes any
}

// ScalarMagicConstFile : ScalarMagicConst
type MagicConstFileScalar struct {
	Attributes any
}

// ScalarMagicConstFunction : ScalarMagicConst
type MagicConstFunctionScalar struct {
	Attributes any
}

// ScalarMagicConstLine : ScalarMagicConst
type MagicConstLineScalar struct {
	Attributes any
}

// ScalarMagicConstMethod : ScalarMagicConst
type MagicConstMethodScalar struct {
	Attributes any
}

// ScalarMagicConstNamespace : ScalarMagicConst
type MagicConstNamespaceScalar struct {
	Attributes any
}

// ScalarMagicConstTrait : ScalarMagicConst
type MagicConstTraitScalar struct {
	Attributes any
}

// ScalarString : Scalar
type StringScalar struct {
	Value        string // @var string String value
	Replacements any
	Attributes   any
}

// StmtBreak : Stmt
type BreakStmt struct {
	Num        Expr // @var null|Node\Expr Number of loops to break
	Attributes any
}

// StmtCase : Stmt
type CaseStmt struct {
	Cond       Expr   // @var null|Node\Expr Condition (null for default)
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtCatch : Stmt
type CatchStmt struct {
	Types      []*Name       // @var Node\Name[] Types of exceptions to catch
	Var        *VariableExpr // @var Expr\Variable|null Variable for exception
	Stmts      []Stmt        // @var Node\Stmt[] Statements
	Attributes any
}

// StmtClass : StmtClassLike
type ClassStmt struct {
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
type ClassConstStmt struct {
	Flags      int               // @var int Modifiers
	Consts     []*Const          // @var Node\Const_[] Constant declarations
	AttrGroups []*AttributeGroup // @var Node\AttributeGroup[]
	Attributes any
}

// StmtClassMethod : Stmt, FunctionLike
type ClassMethodStmt struct {
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
type ConstStmt struct {
	Consts     []*Const // @var Node\Const_[] Constant declarations
	Attributes any
}

// StmtContinue : Stmt
type ContinueStmt struct {
	Num        Expr // @var null|Node\Expr Number of loops to continue
	Attributes any
}

// StmtDeclare : Stmt
type DeclareStmt struct {
	Declares   []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
	Stmts      []Stmt                // @var Node\Stmt[]|null Statements
	Attributes any
}

// StmtDeclareDeclare : Stmt
type DeclareDeclareStmt struct {
	Key        *Identifier // @var Node\Identifier Key
	Value      Expr        // @var Node\Expr Value
	Attributes any
}

// StmtDo : Stmt
type DoStmt struct {
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Cond       Expr   // @var Node\Expr Condition
	Attributes any
}

// StmtEcho : Stmt
type EchoStmt struct {
	Exprs      []Expr // @var Node\Expr[] Expressions
	Attributes any
}

// StmtElse : Stmt
type ElseStmt struct {
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtElseIf : Stmt
type ElseIfStmt struct {
	Cond       Expr   // @var Node\Expr Condition
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtEnum : StmtClassLike
type EnumStmt struct {
	ScalarType     *Identifier       // @var null|Node\Identifier Scalar Type
	Implements     []*Name           // @var Node\Name[] Names of implemented interfaces
	Name           *Identifier       // @var Node\Identifier|null Name
	Stmts          []Stmt            // @var Node\Stmt[] Statements
	AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
	NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	Attributes     any
}

// StmtEnumCase : Stmt
type EnumCaseStmt struct {
	Name       *Identifier       // @var Node\Identifier Enum case name
	Expr       Expr              // @var Node\Expr|null Enum case expression
	AttrGroups []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
	Attributes any
}

// StmtExpression : Stmt
type ExpressionStmt struct {
	Expr       Expr // @var Node\Expr Expression
	Attributes any
}

// StmtFinally : Stmt
type FinallyStmt struct {
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtFor : Stmt
type ForStmt struct {
	Init       []Expr // @var Node\Expr[] Init expressions
	Cond       []Expr // @var Node\Expr[] Loop conditions
	Loop       []Expr // @var Node\Expr[] Loop expressions
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtForeach : Stmt
type ForeachStmt struct {
	Expr       Expr   // @var Node\Expr Expression to iterate
	KeyVar     Expr   // @var null|Node\Expr Variable to assign key to
	ByRef      bool   // @var bool Whether to assign value by reference
	ValueVar   Expr   // @var Node\Expr Variable to assign value to
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtFunction : Stmt, FunctionLike
type FunctionStmt struct {
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
type GlobalStmt struct {
	Vars       []Expr // @var Node\Expr[] Variables
	Attributes any
}

// StmtGoto : Stmt
type GotoStmt struct {
	Name       *Identifier // @var Identifier Name of label to jump to
	Attributes any
}

// StmtGroupUse : Stmt
type GroupUseStmt struct {
	Type       int           // @var int Type of group use
	Prefix     *Name         // @var Name Prefix for uses
	Uses       []*UseUseStmt // @var UseUse[] Uses
	Attributes any
}

// StmtHaltCompiler : Stmt
type HaltCompilerStmt struct {
	Remaining  string // @var string Remaining text after halt compiler statement.
	Attributes any
}

// StmtIf : Stmt
type IfStmt struct {
	Cond       Expr          // @var Node\Expr Condition expression
	Stmts      []Stmt        // @var Node\Stmt[] Statements
	Elseifs    []*ElseIfStmt // @var ElseIf_[] Elseif clauses
	Else       *ElseStmt     // @var null|Else_ Else clause
	Attributes any
}

// StmtInlineHTML : Stmt
type InlineHTMLStmt struct {
	Value      string // @var string String
	Attributes any
}

// StmtInterface : StmtClassLike
type InterfaceStmt struct {
	Extends        []*Name           // @var Node\Name[] Extended interfaces
	Name           *Identifier       // @var Node\Identifier|null Name
	Stmts          []Stmt            // @var Node\Stmt[] Statements
	AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
	NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	Attributes     any
}

// StmtLabel : Stmt
type LabelStmt struct {
	Name       *Identifier // @var Identifier Name
	Attributes any
}

// StmtNamespace : Stmt
type NamespaceStmt struct {
	Name       *Name  // @var null|Node\Name Name
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// StmtNop : Stmt
type NopStmt struct {
	Attributes any
}

// StmtProperty : Stmt
type PropertyStmt struct {
	Flags      int                     // @var int Modifiers
	Props      []*PropertyPropertyStmt // @var PropertyProperty[] Properties
	Type       any                     // @var null|Identifier|Name|ComplexType Type declaration
	AttrGroups []*AttributeGroup       // @var Node\AttributeGroup[] PHP attribute groups
	Attributes any
}

// StmtPropertyProperty : Stmt
type PropertyPropertyStmt struct {
	Name       *VarLikeIdentifier // @var Node\VarLikeIdentifier Name
	Default    Expr               // @var null|Node\Expr Default
	Attributes any
}

// StmtReturn : Stmt
type ReturnStmt struct {
	Expr       Expr // @var null|Node\Expr Expression
	Attributes any
}

// StmtStatic : Stmt
type StaticStmt struct {
	Vars       []*StaticVarStmt // @var StaticVar[] Variable definitions
	Attributes any
}

// StmtStaticVar : Stmt
type StaticVarStmt struct {
	Var        *VariableExpr // @var Expr\Variable Variable
	Default    Expr          // @var null|Node\Expr Default value
	Attributes any
}

// StmtSwitch : Stmt
type SwitchStmt struct {
	Cond       Expr        // @var Node\Expr Condition
	Cases      []*CaseStmt // @var Case_[] Case list
	Attributes any
}

// StmtThrow : Stmt
type ThrowStmt struct {
	Expr       Expr // @var Node\Expr Expression
	Attributes any
}

// StmtTrait : StmtClassLike
type TraitStmt struct {
	Name           *Identifier       // @var Node\Identifier|null Name
	Stmts          []Stmt            // @var Node\Stmt[] Statements
	AttrGroups     []*AttributeGroup // @var Node\AttributeGroup[] PHP attribute groups
	NamespacedName *Name             // @var Node\Name|null Namespaced name (if using NameResolver)
	Attributes     any
}

// StmtTraitUse : Stmt
type TraitUseStmt struct {
	Traits      []*Name                   // @var Node\Name[] Traits
	Adaptations []*TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
	Attributes  any
}

// StmtTraitUseAdaptationAlias : StmtTraitUseAdaptation
type TraitUseAdaptationAliasStmt struct {
	NewModifier int         // @var null|int New modifier
	NewName     *Identifier // @var null|Node\Identifier New name
	Trait       *Name       // @var Node\Name|null Trait name
	Method      *Identifier // @var Node\Identifier Method name
	Attributes  any
}

// StmtTraitUseAdaptationPrecedence : StmtTraitUseAdaptation
type TraitUseAdaptationPrecedenceStmt struct {
	Insteadof  []*Name     // @var Node\Name[] Overwritten traits
	Trait      *Name       // @var Node\Name|null Trait name
	Method     *Identifier // @var Node\Identifier Method name
	Attributes any
}

// StmtTryCatch : Stmt
type TryCatchStmt struct {
	Stmts      []Stmt       // @var Node\Stmt[] Statements
	Catches    []*CatchStmt // @var Catch_[] Catches
	Finally    *FinallyStmt // @var null|Finally_ Optional finally node
	Attributes any
}

// StmtUnset : Stmt
type UnsetStmt struct {
	Vars       []Expr // @var Node\Expr[] Variables to unset
	Attributes any
}

// StmtUse : Stmt
type UseStmt struct {
	Type       int           // @var int Type of alias
	Uses       []*UseUseStmt // @var UseUse[] Aliases
	Attributes any
}

// StmtUseUse : Stmt
type UseUseStmt struct {
	Type       int         // @var int One of the Stmt\Use_::TYPE_* constants. Will only differ from TYPE_UNKNOWN for mixed group uses
	Name       *Name       // @var Node\Name Namespace, class, function or constant to alias
	Alias      *Identifier // @var Identifier|null Alias
	Attributes any
}

// StmtWhile : Stmt
type WhileStmt struct {
	Cond       Expr   // @var Node\Expr Condition
	Stmts      []Stmt // @var Node\Stmt[] Statements
	Attributes any
}

// UnionType : ComplexType
type UnionType struct {
	Types      []any // @var (Identifier|Name|IntersectionType)[] Types
	Attributes any
}

// VarLikeIdentifier : Identifier, Stringable
type VarLikeIdentifier struct {
	Name       string // @var string Identifier as string
	Attributes any
}

// VariadicPlaceholder : PhpParserNodeAbstract
type VariadicPlaceholder struct {
	Attributes any
}

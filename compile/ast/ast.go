package ast

import (
	"strings"
)

type File struct {
	Declares   []*DeclareStmt
	Namespaces []*NamespaceStmt
}

// Node
type Node interface {
	node()
}

// baseNode
type baseNode struct {
	Meta map[string]any `json:"@"`
}

func (*baseNode) node() {}
func (n *baseNode) SetMeta(meta map[string]any) {
	n.Meta = meta
}

// baseExpr
type baseExpr struct {
	baseNode
}

func (*baseExpr) exprNode() {}

// baseStmt
type baseStmt struct {
	baseNode
}

func (*baseStmt) stmtNode() {}

// node interfaces
type (
	Expr interface {
		Node
		exprNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	CallLikeExpr interface {
		Expr
		callLikeExprNode()
	}

	// FunctionLike
	FunctionLike interface {
		Node
		functionLikeNode()
	}

	ClassLikeStmt interface {
		Stmt
		classLikeStmtNode()
	}

	TraitUseAdaptationStmt interface {
		Stmt
		traitUseAdaptationStmtNode()
	}
)

// misc
type (
	Arg struct {
		baseNode
		Value  Expr // @var Expr Value to pass
		Unpack bool // @var bool Whether to unpack the argument
	}

	Ident struct {
		baseNode
		Name string // @var string Ident as string
		/**
		 * Represents a name that is written in source code with a leading dollar,
		 * but is not a proper variable. The leading dollar is not stored as part of the name.
		 *
		 * Examples: Names in property declarations are formatted as variables. Names in static property
		 * lookups are also formatted as variables.
		 */
		VarLike bool
	}

	Param struct {
		baseNode
		Type     TypeHint      // @var TypeHint|null Type declaration
		ByRef    bool          // @var bool Whether parameter is passed by reference
		Variadic bool          // @var bool Whether this is a variadic argument
		Var      *VariableExpr // @var VariableExpr Parameter variable
		Default  Expr          // @var Expr|null Default value
	}
)

// TypeHint 类型标识
type TypeHint interface {
	Node
	typeHintNode()
}

type (
	SimpleType struct {
		baseNode
		Name *Name
	}

	// IntersectionType : A & B & C
	IntersectionType struct {
		baseNode
		Types []TypeHint // possible type: SimpleType
	}

	// UnionType : A | B | C
	UnionType struct {
		baseNode
		Types []TypeHint // possible type: SimpleType or IntersectionType
	}

	// NullableType : ?A
	NullableType struct {
		baseNode
		Type *SimpleType // possible type: SimpleType
	}
)

// Name
type Name struct {
	baseNode
	Kind  NameKind // kind
	Parts []string // @var string[] Parts of the name
}

func NewName(parts ...string) *Name {
	return &Name{Kind: NameNormal, Parts: parts}
}

func (n *Name) IsUnqualified() bool    { return n.Kind == NameNormal && len(n.Parts) == 1 }
func (n *Name) IsQualified() bool      { return n.Kind == NameNormal && len(n.Parts) > 1 }
func (n *Name) IsFullyQualified() bool { return n.Kind == NameFullyQualified }
func (n *Name) IsRelative() bool       { return n.Kind == NameRelative }
func (n *Name) ToString() string       { return strings.Join(n.Parts, "\\") }
func (n *Name) ToCodeString() string {
	switch n.Kind {
	case NameFullyQualified:
		return "\\" + n.ToString()
	case NameRelative:
		return "namespace\\" + n.ToString()
	default: // NameNormal
		return n.ToString()
	}
}

// Expr
type (
	// literal

	IntLit struct {
		baseExpr
		Value int
	}
	FloatLit struct {
		baseExpr
		Value float64
	}
	StringLit struct {
		baseExpr
		Value string // string value
	}

	ArrayExpr struct {
		baseExpr
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	}

	ArrayItemExpr struct {
		baseExpr
		Key    Expr // @var Expr|null Key
		Value  Expr // @var Expr Value
		ByRef  bool // @var bool Whether to assign by reference
		Unpack bool // @var bool Whether to unpack the argument
	}

	// ClosureExpr : FunctionLike
	ClosureExpr struct {
		baseExpr
		Static     bool              // @var bool Whether the closure is static
		ByRef      bool              // @var bool Whether to return by reference
		Params     []*Param          // @var Param[] Parameters
		Uses       []*ClosureUseExpr // @var ClosureUse[] use()s
		ReturnType TypeHint          // @var TypeHint|null Return type
		Stmts      []Stmt            // @var Stmt[] Statements
	}

	ClosureUseExpr struct {
		baseExpr
		Var   *VariableExpr // @var VariableExpr Variable to use
		ByRef bool          // @var bool Whether to use by reference
	}

	// ArrowFunctionExpr : Expr, FunctionLike
	ArrowFunctionExpr struct {
		baseExpr
		Static     bool     // @var bool
		ByRef      bool     // @var bool
		Params     []*Param // @var Param[]
		ReturnType TypeHint // @var TypeHint|null
		Expr       Expr     // @var Expr
	}

	IndexExpr struct {
		baseExpr
		Var Expr // @var Expr       Variable
		Dim Expr // @var Expr|null  Array index / dim
	}

	// CastExpr
	CastExpr struct {
		baseExpr
		Kind CastKind
		Expr Expr // @var Expr Expression
	}

	// UnaryExpr
	UnaryExpr struct {
		baseExpr
		Op  UnaryOpKind
		Var Expr // variable
	}

	// BinaryOpExpr
	BinaryOpExpr struct {
		baseExpr
		Op    BinaryOpKind
		Left  Expr // @var Expr The left-hand side expression
		Right Expr // @var Expr The right-hand side expression
	}

	// AssignExpr
	AssignExpr struct {
		baseExpr
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// AssignOpExpr
	AssignOpExpr struct {
		baseExpr
		Op   AssignOpKind
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	AssignRefExpr struct {
		baseExpr
		Var  Expr // @var Expr Variable reference is assigned to
		Expr Expr // @var Expr Variable which is referenced
	}

	IssetExpr struct {
		baseExpr
		Vars []Expr // @var Expr[] Variables
	}

	EmptyExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	EvalExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	IncludeExpr struct {
		baseExpr
		Kind IncludeKind
		Expr Expr
	}

	CloneExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	ErrorSuppressExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	ExitExpr struct {
		baseExpr
		Expr Expr // @var Expr|null Expression
	}

	ConstFetchExpr struct {
		baseExpr
		Name *Name // @var Name Constant name
	}

	ClassConstFetchExpr struct {
		baseExpr
		Class Node   // @var Name|Expr Class name
		Name  *Ident // @var Ident Constant name
	}

	MagicConstExpr struct {
		baseExpr
		Kind MagicConstKind
	}

	InstanceofExpr struct {
		baseExpr
		Expr  Expr // @var Expr Expression
		Class Node // @var Name|Expr Class name
	}

	ListExpr struct {
		baseExpr
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	}

	PrintExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	PropertyFetchExpr struct {
		baseExpr
		Var      Expr // @var Expr Variable holding object
		Name     Node // @var Ident|Expr Property name
		Nullsafe bool
	}

	StaticPropertyFetchExpr struct {
		baseExpr
		Class Node // @var Name|Expr Class name
		Name  Node // @var Ident|Expr Property name
	}

	ShellExecExpr struct {
		baseExpr
		Parts []Expr // @var array Encapsed string array
	}

	TernaryExpr struct {
		baseExpr
		Cond Expr // @var Expr Condition
		If   Expr // @var Expr|null Expression for true
		Else Expr // @var Expr Expression for false
	}

	ThrowExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression
	}

	VariableExpr struct {
		baseExpr
		Name Node // @var Ident|Expr Name
	}

	YieldExpr struct {
		baseExpr
		Key   Expr // @var Expr|null Key expression
		Value Expr // @var Expr|null Value expression
	}

	YieldFromExpr struct {
		baseExpr
		Expr Expr // @var Expr Expression to yield from
	}

	// FuncCallExpr : CallLikeExpr
	FuncCallExpr struct {
		baseExpr
		Name Node   // @var Name|Expr Function name
		Args []*Arg // @var Arguments
	}

	// NewExpr : CallLikeExpr
	NewExpr struct {
		baseExpr
		Class Node   // @var Name|Expr|ClassStmt Class name
		Args  []*Arg // @var Arguments
	}

	// MethodCallExpr : CallLikeExpr
	MethodCallExpr struct {
		baseExpr
		Var      Expr   // @var Expr Variable holding object
		Name     Node   // @var Ident|Expr Method name
		Args     []*Arg // @var Arguments
		Nullsafe bool
	}

	// StaticCallExpr : CallLikeExpr
	StaticCallExpr struct {
		baseExpr
		Class Node   // @var Name|Expr Class name
		Name  Node   // @var Ident|Expr Method name
		Args  []*Arg // @var Arguments
	}
)

// Stmt
type (
	EmptyStmt struct {
		baseStmt
	}

	BlockStmt struct {
		baseStmt
		List []Stmt
	}

	ExprStmt struct {
		baseStmt
		Expr Expr // @var Expr Expression
	}

	ReturnStmt struct {
		baseStmt
		Expr Expr // @var Expr|null Expression
	}

	LabelStmt struct {
		baseStmt
		Name *Ident // @var Ident Name
	}

	GotoStmt struct {
		baseStmt
		Name *Ident // @var Ident Name of label to jump to
	}

	// IfStmt
	IfStmt struct {
		baseStmt
		Cond    Expr          // @var Expr 			condition expression
		Stmts   []Stmt        // @var Stmt[] 		body statements
		Elseifs []*ElseIfStmt // @var ElseIfStmt[] 	elseif branches
		Else    *ElseStmt     // @var ElseStmt|null else branch; or nil
	}

	ElseIfStmt struct {
		baseStmt
		Cond  Expr   // @var Expr 	Condition
		Stmts []Stmt // @var Stmt[] Statements
	}

	ElseStmt struct {
		baseStmt
		Stmts []Stmt // @var Stmt[] Statements
	}

	// SwitchStmt
	SwitchStmt struct {
		baseStmt
		Cond  Expr        // @var Expr Condition
		Cases []*CaseStmt // @var Case_[] Case list
	}

	CaseStmt struct {
		baseStmt
		Cond  Expr   // @var Expr|null Condition (null for default)
		Stmts []Stmt // @var Stmt[] Statements
	}

	// ForStmt
	ForStmt struct {
		baseStmt
		Init  []Expr // @var Expr[] Init expressions
		Cond  []Expr // @var Expr[] Loop conditions
		Loop  []Expr // @var Expr[] Loop expressions
		Stmts []Stmt // @var Stmt[] Statements
	}

	// ForeachStmt
	ForeachStmt struct {
		baseStmt
		Expr     Expr   // @var Expr Expression to iterate
		KeyVar   Expr   // @var Expr|null Variable to assign key to
		ByRef    bool   // @var bool Whether to assign value by reference
		ValueVar Expr   // @var Expr Variable to assign value to
		Stmts    []Stmt // @var Stmt[] Statements
	}

	BreakStmt struct {
		baseStmt
		Num Expr // @var Expr|null Number of loops to break
	}

	ContinueStmt struct {
		baseStmt
		Num Expr // @var Expr|null Number of loops to continue
	}

	// WhileStmt
	WhileStmt struct {
		baseStmt
		Cond  Expr   // @var Expr Condition
		Stmts []Stmt // @var Stmt[] Statements
	}

	// DoStmt
	DoStmt struct {
		baseStmt
		Stmts []Stmt // @var Stmt[] Statements
		Cond  Expr   // @var Expr Condition
	}

	// try-catch-finally
	TryCatchStmt struct {
		baseStmt
		Stmts   []Stmt       // @var Stmt[] Statements
		Catches []*CatchStmt // @var Catch_[] Catches
		Finally *FinallyStmt // @var Finally_|null Optional finally node
	}

	CatchStmt struct {
		baseStmt
		Types []*Name       // @var Name[] Types of exceptions to catch
		Var   *VariableExpr // @var VariableExpr|null Variable for exception
		Stmts []Stmt        // @var Stmt[] Statements
	}

	FinallyStmt struct {
		baseStmt
		Stmts []Stmt // @var Stmt[] Statements
	}

	ConstStmt struct {
		baseStmt
		Name           *Ident // @var Ident Name
		Value          Expr   // @var Expr Value
		NamespacedName *Name  // @var Name|null Namespaced name (if using NameResolver)
	}

	EchoStmt struct {
		baseStmt
		Exprs []Expr // @var Expr[] Expressions
	}

	GlobalStmt struct {
		baseStmt
		Vars []Expr // @var Expr[] Variables
	}

	HaltCompilerStmt struct {
		baseStmt
		Remaining string // @var string Remaining text after halt compiler statement.
	}

	InlineHTMLStmt struct {
		baseStmt
		Value string // @var string String
	}

	StaticStmt struct {
		baseStmt
		Vars []*StaticVarStmt // @var StaticVar[] Variable definitions
	}

	StaticVarStmt struct {
		baseStmt
		Var     *VariableExpr // @var VariableExpr Variable
		Default Expr          // @var Expr|null Default value
	}

	UnsetStmt struct {
		baseStmt
		Vars []Expr // @var Expr[] Variables to unset
	}

	UseStmt struct {
		baseStmt
		Type  UseType // @var UseType     UseNormal UseFunction Or UseConstant
		Name  *Name   // @var Name        Namespace, class, function or constant to alias
		Alias *Ident  // @var Ident|null  Alias Name, or nil
	}

	// DeclareStmt
	DeclareStmt struct {
		baseStmt
		Declares []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
		Stmts    []Stmt                // @var Stmt[]|null Statements
	}

	DeclareDeclareStmt struct {
		baseStmt
		Key   *Ident // @var Ident Key
		Value Expr   // @var Expr Value
	}

	// NamespaceStmt
	NamespaceStmt struct {
		baseStmt
		Name  *Name  // @var Name|null Name
		Stmts []Stmt // @var Stmt[] Statements
	}

	// FunctionStmt : Stmt, FunctionLike
	FunctionStmt struct {
		baseStmt
		ByRef          bool     // @var bool Whether function returns by reference
		Name           *Ident   // @var Ident Name
		Params         []*Param // @var Param[] Parameters
		ReturnType     TypeHint // @var TypeHint|null Return type
		Stmts          []Stmt   // @var Stmt[] Statements
		NamespacedName *Name    // @var Name|null Namespaced name (if using NameResolver)
	}

	// InterfaceStmt : ClassLikeStmt
	InterfaceStmt struct {
		baseStmt
		Extends        []*Name // @var Name[] Extended interfaces
		Name           *Ident  // @var Ident|null Name
		Stmts          []Stmt  // @var Stmt[] Statements
		NamespacedName *Name   // @var Name|null Namespaced name (if using NameResolver)
	}

	// ClassStmt : ClassLikeStmt
	ClassStmt struct {
		baseStmt
		Flags          Flags   // @var Flags      Type
		Extends        *Name   // @var Name|null  Name of extended class
		Implements     []*Name // @var Name[]     Names of implemented interfaces
		Name           *Ident  // @var Ident|null Name
		Stmts          []Stmt  // @var Stmt[] Statements
		NamespacedName *Name   // @var Name|null Namespaced name (if using NameResolver)
	}

	ClassConstStmt struct {
		baseStmt
		Flags Flags    // @var Flags Modifiers
		Type  TypeHint // @var TypeHint|null Type declaration
		Name  *Ident   // @var Ident Name
		Value Expr     // @var Expr Value
	}

	PropertyStmt struct {
		baseStmt
		Flags   Flags    // @var Flags Modifiers
		Type    TypeHint // @var TypeHint|null Type declaration
		Name    *Ident   // @var Ident Name
		Default Expr     // @var Expr|null Default
	}

	// ClassMethodStmt : Stmt, FunctionLike
	ClassMethodStmt struct {
		baseStmt
		Flags      Flags    // @var Flags Modifiers
		ByRef      bool     // @var bool Whether to return by reference
		Name       *Ident   // @var Ident Name
		Params     []*Param // @var Param[] Parameters
		ReturnType TypeHint // @var TypeHint|null Return type
		Stmts      []Stmt   // @var Stmt[]|null Statements
	}

	// TraitStmt : ClassLikeStmt
	TraitStmt struct {
		baseStmt
		Name           *Ident // @var Ident|null Name
		Stmts          []Stmt // @var Stmt[] Statements
		NamespacedName *Name  // @var Name|null Namespaced name (if using NameResolver)
	}

	TraitUseStmt struct {
		baseStmt
		Traits      []*Name                  // @var Name[] Traits
		Adaptations []TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
	}

	// TraitUseAdaptationAliasStmt : TraitUseAdaptationStmt
	TraitUseAdaptationAliasStmt struct {
		baseStmt
		NewModifier Flags  // @var Flags 	    New modifier, default 0
		NewName     *Ident // @var Ident|null 	New name, or nil
		Trait       *Name  // @var Name|null 	Trait name, or nil
		Method      *Ident // @var Ident Method name
	}

	// TraitUseAdaptationPrecedenceStmt : TraitUseAdaptationStmt
	TraitUseAdaptationPrecedenceStmt struct {
		baseStmt
		Insteadof []*Name // @var Name[] Overwritten traits
		Trait     *Name   // @var Name|null Trait name
		Method    *Ident  // @var Ident Method name
	}
)

// TypeHint
func (*SimpleType) typeHintNode()       {}
func (*IntersectionType) typeHintNode() {}
func (*UnionType) typeHintNode()        {}
func (*NullableType) typeHintNode()     {}

// CallLikeExpr
func (*FuncCallExpr) callLikeExprNode()   {}
func (*NewExpr) callLikeExprNode()        {}
func (*MethodCallExpr) callLikeExprNode() {}
func (*StaticCallExpr) callLikeExprNode() {}

// FunctionLike
func (*ArrowFunctionExpr) functionLikeNode() {}
func (*ClosureExpr) functionLikeNode()       {}
func (*ClassMethodStmt) functionLikeNode()   {}
func (*FunctionStmt) functionLikeNode()      {}

// ClassLikeStmt
func (*ClassStmt) classLikeStmtNode()     {}
func (*InterfaceStmt) classLikeStmtNode() {}
func (*TraitStmt) classLikeStmtNode()     {}

// TraitUseAdaptationStmt
func (*TraitUseAdaptationAliasStmt) traitUseAdaptationStmtNode()      {}
func (*TraitUseAdaptationPrecedenceStmt) traitUseAdaptationStmtNode() {}

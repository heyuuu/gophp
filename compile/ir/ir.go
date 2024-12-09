package ir

import (
	"github.com/heyuuu/gophp/compile/ast"
	"strings"
)

type File struct {
	FilePath    string
	StrictTypes bool
	Namespaces  []*NamespaceStmt
}

// Node
type Node interface {
	node()
	Comments() []*Comment
	SetComments([]*Comment)
}

// baseNode
type baseNode struct {
	comments []*Comment
}

func (*baseNode) node() {}

func (n *baseNode) Comments() []*Comment {
	return n.comments
}

func (n *baseNode) SetComments(comments []*Comment) {
	n.comments = comments
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
		Value  Expr // value to pass
		Unpack bool // whether to unpack the argument
	}

	Ident struct {
		baseNode
		Name string
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

	Comment struct {
		baseNode
		Type ast.CommentType
		Text string
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
	Kind  ast.NameKind
	Parts []string // parts of the name
}

func NewName(parts ...string) *Name {
	return &Name{Kind: ast.NameNormal, Parts: parts}
}

func (n *Name) IsUnqualified() bool    { return n.Kind == ast.NameNormal && len(n.Parts) == 1 }
func (n *Name) IsQualified() bool      { return n.Kind == ast.NameNormal && len(n.Parts) > 1 }
func (n *Name) IsFullyQualified() bool { return n.Kind == ast.NameFullyQualified }
func (n *Name) IsRelative() bool       { return n.Kind == ast.NameRelative }
func (n *Name) ToString() string       { return strings.Join(n.Parts, "\\") }
func (n *Name) ToCodeString() string {
	switch n.Kind {
	case ast.NameFullyQualified:
		return "\\" + n.ToString()
	case ast.NameRelative:
		return "namespace\\" + n.ToString()
	default: // NameNormal
		return n.ToString()
	}
}

// Expr
type (
	// literal

	NullLit struct {
		baseExpr
	}
	BoolLit struct {
		baseExpr
		Value bool
	}
	IntLit struct {
		baseExpr
		Raw   string
		Value int
	}
	FloatLit struct {
		baseExpr
		Raw   string
		Value float64
	}
	StringLit struct {
		baseExpr
		Raw   string
		Value string
	}

	ArrayExpr struct {
		baseExpr
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	}

	ArrayItemExpr struct {
		baseExpr
		Key    Expr // @var Expr|null Key
		Value  Expr // Value
		ByRef  bool // Whether to assign by reference
		Unpack bool // Whether to unpack the argument
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
		Var   *VariableExpr // variable to use
		ByRef bool          // whether to use by reference
	}

	// ArrowFunctionExpr : FunctionLike
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

	CastExpr struct {
		baseExpr
		Kind ast.CastKind
		Expr Expr
	}

	UnaryExpr struct {
		baseExpr
		Op  ast.UnaryOpKind
		Var Expr // variable
	}

	BinaryOpExpr struct {
		baseExpr
		Op    ast.BinaryOpKind
		Left  Expr // the left hand side expression
		Right Expr // the right hand side expression
	}

	AssignExpr struct {
		baseExpr
		Var  Expr // variable
		Expr Expr // expression
	}

	AssignOpExpr struct {
		baseExpr
		Op   ast.AssignOpKind
		Var  Expr // variable
		Expr Expr // expression
	}

	AssignRefExpr struct {
		baseExpr
		Var  Expr // variable reference is assigned to
		Expr Expr // variable which is referenced
	}

	IssetExpr struct {
		baseExpr
		Vars []Expr // @var Expr[] Variables
	}

	EmptyExpr struct {
		baseExpr
		Expr Expr
	}

	EvalExpr struct {
		baseExpr
		Expr Expr
	}

	IncludeExpr struct {
		baseExpr
		Kind ast.IncludeKind
		Expr Expr
	}

	CloneExpr struct {
		baseExpr
		Expr Expr
	}

	ErrorSuppressExpr struct {
		baseExpr
		Expr Expr
	}

	ExitExpr struct {
		baseExpr
		Expr Expr // @var Expr|null Expression
	}

	ConstFetchExpr struct {
		baseExpr
		Name *Name // constant name
	}

	ClassConstFetchExpr struct {
		baseExpr
		Class Node   // @var Name|Expr Class name
		Name  *Ident // @var Ident Constant name
	}

	MagicConstExpr struct {
		baseExpr
		Kind ast.MagicConstKind
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
		Expr Expr
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
		Expr Expr
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
		Expr Expr // Expression to yield from
	}

	// FuncCallExpr : CallLikeExpr
	FuncCallExpr struct {
		baseExpr
		Name Node   // @var Name|Expr Function name
		Args []*Arg // Arguments
	}

	// NewExpr : CallLikeExpr
	NewExpr struct {
		baseExpr
		Class Node   // @var Name|Expr|ClassStmt class name
		Args  []*Arg // Arguments
	}

	// MethodCallExpr : CallLikeExpr
	MethodCallExpr struct {
		baseExpr
		Var      Expr   // @var Expr Variable holding object
		Name     Node   // @var Ident|Expr Method name
		Args     []*Arg // arguments
		Nullsafe bool
	}

	// StaticCallExpr : CallLikeExpr
	StaticCallExpr struct {
		baseExpr
		Class Node   // @var Name|Expr Class name
		Name  Node   // @var Ident|Expr Method name
		Args  []*Arg // Arguments
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
		Expr Expr
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

	IfStmt struct {
		baseStmt
		Cond    Expr          // @var Expr 			condition expression
		Stmts   []Stmt        // @var Stmt[] 		body statements
		Elseifs []*ElseIfStmt // @var ElseIfStmt[] 	elseif branches
		Else    *ElseStmt     // @var ElseStmt|null else branch; or nil
	}

	ElseIfStmt struct {
		baseStmt
		Cond  Expr   // condition
		Stmts []Stmt // statements
	}

	ElseStmt struct {
		baseStmt
		Stmts []Stmt // statements
	}

	SwitchStmt struct {
		baseStmt
		Cond  Expr        // condition
		Cases []*CaseStmt // case list
	}

	CaseStmt struct {
		baseStmt
		Cond  Expr   // @var Expr|null Condition (null for default)
		Stmts []Stmt // statements
	}

	ForStmt struct {
		baseStmt
		Init  []Expr // @var Expr[] Init expressions
		Cond  []Expr // @var Expr[] Loop conditions
		Loop  []Expr // @var Expr[] Loop expressions
		Stmts []Stmt // @var Stmt[] Statements
	}

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
		Num Expr // @var Expr|null number of loops to break
	}

	ContinueStmt struct {
		baseStmt
		Num Expr // @var Expr|null Number of loops to continue
	}

	WhileStmt struct {
		baseStmt
		Cond  Expr   // condition
		Stmts []Stmt // statements
	}

	DoStmt struct {
		baseStmt
		Stmts []Stmt // @var Stmt[] Statements
		Cond  Expr   // @var Expr Condition
	}

	// try-catch-finally
	TryCatchStmt struct {
		baseStmt
		Stmts   []Stmt       // statements
		Catches []*CatchStmt // catches
		Finally *FinallyStmt // @var FinallyStmt|null Optional finally node
	}

	CatchStmt struct {
		baseStmt
		Types []*Name       // types of exceptions to catch
		Var   *VariableExpr // @var VariableExpr|null Variable for exception
		Stmts []Stmt        // statements
	}

	FinallyStmt struct {
		baseStmt
		Stmts []Stmt // @var Stmt[] Statements
	}

	ConstStmt struct {
		baseStmt
		Name           *Ident
		Value          Expr
		NamespacedName *Name // @var Name|null namespaced name (if using NameResolver)
	}

	EchoStmt struct {
		baseStmt
		Exprs []Expr // @var Expr[] Expressions
	}

	GlobalStmt struct {
		baseStmt
		Var Expr // variable
	}

	StaticStmt struct {
		baseStmt
		Var     *VariableExpr // variable
		Default Expr          // @var Expr|null default value
	}

	UnsetStmt struct {
		baseStmt
		Var Expr // variables to unset
	}

	UseStmt struct {
		baseStmt
		Type  ast.UseType // @var UseType     UseNormal UseFunction Or UseConstant
		Name  *Name       // @var Name        Namespace, class, function or constant to alias
		Alias *Ident      // @var Ident|null  Alias Name, or nil
	}

	NamespaceStmt struct {
		baseStmt
		Name  *Name  // @var Name|null Name
		Stmts []Stmt // @var Stmt[] Statements
	}

	// FunctionStmt : FunctionLike
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
		Flags          ast.Flags
		Extends        *Name   // @var Name|null  Name of extended class
		Implements     []*Name // @var Name[]     Names of implemented interfaces
		Name           *Ident  // @var Ident|null Name
		Stmts          []Stmt  // @var Stmt[] Statements
		NamespacedName *Name   // @var Name|null Namespaced name (if using NameResolver)
	}

	ClassConstStmt struct {
		baseStmt
		Flags ast.Flags // @var Flags Modifiers
		Type  TypeHint  // @var TypeHint|null Type declaration
		Name  *Ident    // @var Ident Name
		Value Expr      // @var Expr Value
	}

	PropertyStmt struct {
		baseStmt
		Flags   ast.Flags // @var Flags Modifiers
		Type    TypeHint  // @var TypeHint|null Type declaration
		Name    *Ident    // @var Ident Name
		Default Expr      // @var Expr|null Default
	}

	// ClassMethodStmt : FunctionLike
	ClassMethodStmt struct {
		baseStmt
		Flags      ast.Flags // @var Flags Modifiers
		ByRef      bool      // @var bool Whether to return by reference
		Name       *Ident    // @var Ident Name
		Params     []*Param  // @var Param[] Parameters
		ReturnType TypeHint  // @var TypeHint|null Return type
		Stmts      []Stmt    // @var Stmt[]|null Statements
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
		NewModifier ast.Flags // new modifier, default 0
		NewName     *Ident    // @var Ident|null 	New name, or nil
		Trait       *Name     // @var Name|null 	Trait name, or nil
		Method      *Ident    // @var Ident Method name
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

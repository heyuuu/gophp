package ast

import (
	"strings"
)

type File struct {
	Declares   []*DeclareStmt
	Namespaces []*NamespaceStmt
}

type (
	Node interface {
		node()
	}

	Expr interface {
		Node
		exprNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	// CallLikeExpr : Expr
	CallLikeExpr interface {
		Expr
		callLikeExprNode()
	}

	// FunctionLike
	FunctionLike interface {
		Node
		functionLikeNode()
	}

	// ClassLikeStmt : Stmt
	ClassLikeStmt interface {
		Stmt
		classLikeStmtNode()
	}

	// TraitUseAdaptationStmt : Stmt
	TraitUseAdaptationStmt interface {
		Stmt
		traitUseAdaptationStmtNode()
	}
)

// misc
type (
	Arg struct {
		Value  Expr // @var Expr Value to pass
		Unpack bool // @var bool Whether to unpack the argument
	}

	Ident struct {
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
		Name *Name
	}

	// IntersectionType : A & B & C
	IntersectionType struct {
		Types []TypeHint // possible type: SimpleType
	}

	// UnionType : A | B | C
	UnionType struct {
		Types []TypeHint // possible type: SimpleType or IntersectionType
	}

	// NullableType : ?A
	NullableType struct {
		Type *SimpleType // possible type: SimpleType
	}
)

// Name : Node
type Name struct {
	Kind  NameKind // kind
	Parts []string // @var string[] Parts of the name
}

func NewName(parts ...string) *Name {
	return &Name{Parts: parts}
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
		Value int // number value
	}

	FloatLit struct {
		Value float64 // number value
	}

	StringLit struct {
		Value string // string value
	}

	ArrayExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	}

	ArrayItemExpr struct {
		Key    Expr // @var Expr|null Key
		Value  Expr // @var Expr Value
		ByRef  bool // @var bool Whether to assign by reference
		Unpack bool // @var bool Whether to unpack the argument
	}

	// ClosureExpr : Expr, FunctionLike
	ClosureExpr struct {
		Static     bool              // @var bool Whether the closure is static
		ByRef      bool              // @var bool Whether to return by reference
		Params     []*Param          // @var Param[] Parameters
		Uses       []*ClosureUseExpr // @var ClosureUse[] use()s
		ReturnType TypeHint          // @var TypeHint|null Return type
		Stmts      []Stmt            // @var Stmt[] Statements
	}

	ClosureUseExpr struct {
		Var   *VariableExpr // @var VariableExpr Variable to use
		ByRef bool          // @var bool Whether to use by reference
	}

	// ArrowFunctionExpr : Expr, FunctionLike
	ArrowFunctionExpr struct {
		Static     bool     // @var bool
		ByRef      bool     // @var bool
		Params     []*Param // @var Param[]
		ReturnType TypeHint // @var TypeHint|null
		Expr       Expr     // @var Expr
	}

	// IndexExpr
	IndexExpr struct {
		Var Expr // @var Expr       Variable
		Dim Expr // @var Expr|null  Array index / dim
	}

	// CastExpr
	CastExpr struct {
		Kind CastKind
		Expr Expr // @var Expr Expression
	}

	// UnaryExpr
	UnaryExpr struct {
		Op  UnaryOpKind
		Var Expr // variable
	}

	// BinaryOpExpr
	BinaryOpExpr struct {
		Op    BinaryOpKind
		Left  Expr // @var Expr The left-hand side expression
		Right Expr // @var Expr The right-hand side expression
	}

	// AssignExpr
	AssignExpr struct {
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	// AssignOpExpr
	AssignOpExpr struct {
		Op   AssignOpKind
		Var  Expr // @var Expr Variable
		Expr Expr // @var Expr Expression
	}

	AssignRefExpr struct {
		Var  Expr // @var Expr Variable reference is assigned to
		Expr Expr // @var Expr Variable which is referenced
	}

	IssetExpr struct {
		Vars []Expr // @var Expr[] Variables
	}

	EmptyExpr struct {
		Expr Expr // @var Expr Expression
	}

	EvalExpr struct {
		Expr Expr // @var Expr Expression
	}

	IncludeExpr struct {
		Kind IncludeKind // @var int Type of include
		Expr Expr        // @var Expr Expression
	}

	CloneExpr struct {
		Expr Expr // @var Expr Expression
	}

	ErrorSuppressExpr struct {
		Expr Expr // @var Expr Expression
	}

	ExitExpr struct {
		Expr Expr // @var Expr|null Expression
	}

	ConstFetchExpr struct {
		Name *Name // @var Name Constant name
	}

	ClassConstFetchExpr struct {
		Class Node   // @var Name|Expr Class name
		Name  *Ident // @var Ident Constant name
	}

	MagicConstExpr struct {
		Kind MagicConstKind
	}

	InstanceofExpr struct {
		Expr  Expr // @var Expr Expression
		Class Node // @var Name|Expr Class name
	}

	ListExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	}

	PrintExpr struct {
		Expr Expr // @var Expr Expression
	}

	// PropertyFetchExpr : Expr
	PropertyFetchExpr struct {
		Var      Expr // @var Expr Variable holding object
		Name     Node // @var Ident|Expr Property name
		Nullable bool
	}

	StaticPropertyFetchExpr struct {
		Class Node // @var Name|Expr Class name
		Name  Node // @var Ident|Expr Property name
	}

	// ShellExecExpr : Expr
	ShellExecExpr struct {
		Parts []Expr // @var array Encapsed string array
	}

	// TernaryExpr : Expr
	TernaryExpr struct {
		Cond Expr // @var Expr Condition
		If   Expr // @var Expr|null Expression for true
		Else Expr // @var Expr Expression for false
	}

	// ThrowExpr : Expr
	ThrowExpr struct {
		Expr Expr // @var Expr Expression
	}

	// VariableExpr : Expr
	VariableExpr struct {
		Name Node // @var Ident|Expr Name
	}

	// YieldExpr : Expr
	YieldExpr struct {
		Key   Expr // @var Expr|null Key expression
		Value Expr // @var Expr|null Value expression
	}

	// YieldFromExpr : Expr
	YieldFromExpr struct {
		Expr Expr // @var Expr Expression to yield from
	}

	// FuncCallExpr : CallLikeExpr
	FuncCallExpr struct {
		Name Node   // @var Name|Expr Function name
		Args []*Arg // @var Arguments
	}

	// NewExpr : CallLikeExpr
	NewExpr struct {
		Class Node   // @var Name|Expr|ClassStmt Class name
		Args  []*Arg // @var Arguments
	}

	// MethodCallExpr : CallLikeExpr
	MethodCallExpr struct {
		Var      Expr   // @var Expr Variable holding object
		Name     Node   // @var Ident|Expr Method name
		Args     []*Arg // @var Arguments
		Nullsafe bool
	}

	// StaticCallExpr : CallLikeExpr
	StaticCallExpr struct {
		Class Node   // @var Name|Expr Class name
		Name  Node   // @var Ident|Expr Method name
		Args  []*Arg // @var Arguments
	}
)

// UseType for UseStmt
type UseType int

const (
	_           = iota
	UseNormal   = 1 // Class or namespace import
	UseFunction = 2 // Function import
	UseConstant = 3 // Constant import
)

// flags
type Flags int

func (f Flags) Is(flags Flags) bool { return f&flags != 0 }

const (
	// 此处不写成 1 << iota 形式，为了表示与 PHP Parser 对齐
	FlagPublic    Flags = 1
	FlagProtected Flags = 2
	FlagPrivate   Flags = 4
	FlagStatic    Flags = 8
	FlagAbstract  Flags = 16
	FlagFinal     Flags = 32
	FlagReadonly  Flags = 64

	VisibilityModifierMask = FlagPublic | FlagProtected | FlagPrivate
)

// Stmt
type (
	EmptyStmt struct{}

	BlockStmt struct {
		List []Stmt
	}

	ExprStmt struct {
		Expr Expr // @var Expr Expression
	}

	ReturnStmt struct {
		Expr Expr // @var Expr|null Expression
	}

	LabelStmt struct {
		Name *Ident // @var Ident Name
	}

	GotoStmt struct {
		Name *Ident // @var Ident Name of label to jump to
	}

	// IfStmt
	IfStmt struct {
		Cond    Expr          // @var Expr 			condition expression
		Stmts   []Stmt        // @var Stmt[] 		body statements
		Elseifs []*ElseIfStmt // @var ElseIfStmt[] 	elseif branches
		Else    *ElseStmt     // @var ElseStmt|null else branch; or nil
	}

	ElseIfStmt struct {
		Cond  Expr   // @var Expr 	Condition
		Stmts []Stmt // @var Stmt[] Statements
	}

	ElseStmt struct {
		Stmts []Stmt // @var Stmt[] Statements
	}

	// SwitchStmt
	SwitchStmt struct {
		Cond  Expr        // @var Expr Condition
		Cases []*CaseStmt // @var Case_[] Case list
	}

	CaseStmt struct {
		Cond  Expr   // @var Expr|null Condition (null for default)
		Stmts []Stmt // @var Stmt[] Statements
	}

	// ForStmt
	ForStmt struct {
		Init  []Expr // @var Expr[] Init expressions
		Cond  []Expr // @var Expr[] Loop conditions
		Loop  []Expr // @var Expr[] Loop expressions
		Stmts []Stmt // @var Stmt[] Statements
	}

	// ForeachStmt
	ForeachStmt struct {
		Expr     Expr   // @var Expr Expression to iterate
		KeyVar   Expr   // @var Expr|null Variable to assign key to
		ByRef    bool   // @var bool Whether to assign value by reference
		ValueVar Expr   // @var Expr Variable to assign value to
		Stmts    []Stmt // @var Stmt[] Statements
	}

	BreakStmt struct {
		Num Expr // @var Expr|null Number of loops to break
	}

	ContinueStmt struct {
		Num Expr // @var Expr|null Number of loops to continue
	}

	// WhileStmt
	WhileStmt struct {
		Cond  Expr   // @var Expr Condition
		Stmts []Stmt // @var Stmt[] Statements
	}

	// DoStmt
	DoStmt struct {
		Stmts []Stmt // @var Stmt[] Statements
		Cond  Expr   // @var Expr Condition
	}

	// try-catch-finally
	TryCatchStmt struct {
		Stmts   []Stmt       // @var Stmt[] Statements
		Catches []*CatchStmt // @var Catch_[] Catches
		Finally *FinallyStmt // @var Finally_|null Optional finally node
	}

	CatchStmt struct {
		Types []*Name       // @var Name[] Types of exceptions to catch
		Var   *VariableExpr // @var VariableExpr|null Variable for exception
		Stmts []Stmt        // @var Stmt[] Statements
	}

	FinallyStmt struct {
		Stmts []Stmt // @var Stmt[] Statements
	}

	ConstStmt struct {
		Name           *Ident // @var Ident Name
		Value          Expr   // @var Expr Value
		NamespacedName *Name  // @var Name|null Namespaced name (if using NameResolver)
	}

	EchoStmt struct {
		Exprs []Expr // @var Expr[] Expressions
	}

	GlobalStmt struct {
		Vars []Expr // @var Expr[] Variables
	}

	HaltCompilerStmt struct {
		Remaining string // @var string Remaining text after halt compiler statement.
	}

	InlineHTMLStmt struct {
		Value string // @var string String
	}

	StaticStmt struct {
		Vars []*StaticVarStmt // @var StaticVar[] Variable definitions
	}

	StaticVarStmt struct {
		Var     *VariableExpr // @var VariableExpr Variable
		Default Expr          // @var Expr|null Default value
	}

	UnsetStmt struct {
		Vars []Expr // @var Expr[] Variables to unset
	}

	UseStmt struct {
		Type  UseType // @var UseType     UseNormal UseFunction Or UseConstant
		Name  *Name   // @var Name        Namespace, class, function or constant to alias
		Alias *Ident  // @var Ident|null  Alias Name, or nil
	}

	// DeclareStmt
	DeclareStmt struct {
		Declares []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
		Stmts    []Stmt                // @var Stmt[]|null Statements
	}

	DeclareDeclareStmt struct {
		Key   *Ident // @var Ident Key
		Value Expr   // @var Expr Value
	}

	// NamespaceStmt
	NamespaceStmt struct {
		Name  *Name  // @var Name|null Name
		Stmts []Stmt // @var Stmt[] Statements
	}

	// FunctionStmt : Stmt, FunctionLike
	FunctionStmt struct {
		ByRef          bool     // @var bool Whether function returns by reference
		Name           *Ident   // @var Ident Name
		Params         []*Param // @var Param[] Parameters
		ReturnType     TypeHint // @var TypeHint|null Return type
		Stmts          []Stmt   // @var Stmt[] Statements
		NamespacedName *Name    // @var Name|null Namespaced name (if using NameResolver)
	}

	// InterfaceStmt
	InterfaceStmt struct {
		Extends        []*Name // @var Name[] Extended interfaces
		Name           *Ident  // @var Ident|null Name
		Stmts          []Stmt  // @var Stmt[] Statements
		NamespacedName *Name   // @var Name|null Namespaced name (if using NameResolver)
	}

	// ClassStmt : ClassLikeStmt
	ClassStmt struct {
		Flags          Flags   // @var Flags      Type
		Extends        *Name   // @var Name|null  Name of extended class
		Implements     []*Name // @var Name[]     Names of implemented interfaces
		Name           *Ident  // @var Ident|null Name
		Stmts          []Stmt  // @var Stmt[] Statements
		NamespacedName *Name   // @var Name|null Namespaced name (if using NameResolver)
	}

	// ClassConstStmt : Stmt
	ClassConstStmt struct {
		Flags Flags  // @var Flags Modifiers
		Name  *Ident // @var Ident Name
		Value Expr   // @var Expr Value
	}

	// PropertyStmt : Stmt
	PropertyStmt struct {
		Flags   Flags    // @var Flags Modifiers
		Type    TypeHint // @var TypeHint|null Type declaration
		Name    *Ident   // @var Ident     Name
		Default Expr     // @var Expr|null Default
	}

	// ClassMethodStmt : Stmt, FunctionLike
	ClassMethodStmt struct {
		Flags      Flags    // @var Flags Modifiers
		ByRef      bool     // @var bool Whether to return by reference
		Name       *Ident   // @var Ident Name
		Params     []*Param // @var Param[] Parameters
		ReturnType TypeHint // @var TypeHint|null Return type
		Stmts      []Stmt   // @var Stmt[]|null Statements
	}

	// TraitStmt : ClassLikeStmt
	TraitStmt struct {
		Name           *Ident // @var Ident|null Name
		Stmts          []Stmt // @var Stmt[] Statements
		NamespacedName *Name  // @var Name|null Namespaced name (if using NameResolver)
	}

	TraitUseStmt struct {
		Traits      []*Name                  // @var Name[] Traits
		Adaptations []TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
	}

	// TraitUseAdaptationAliasStmt : TraitUseAdaptationStmt
	TraitUseAdaptationAliasStmt struct {
		NewModifier Flags  // @var Flags 	    New modifier, default 0
		NewName     *Ident // @var Ident|null 	New name, or nil
		Trait       *Name  // @var Name|null 	Trait name, or nil
		Method      *Ident // @var Ident Method name
	}

	// TraitUseAdaptationPrecedenceStmt : TraitUseAdaptationStmt
	TraitUseAdaptationPrecedenceStmt struct {
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

// Expr
func (*IntLit) exprNode()        {}
func (*FloatLit) exprNode()      {}
func (*StringLit) exprNode()     {}
func (*ArrayExpr) exprNode()     {}
func (*ArrayItemExpr) exprNode() {}

func (*ClosureExpr) exprNode()       {}
func (*ClosureUseExpr) exprNode()    {}
func (*ArrowFunctionExpr) exprNode() {}

func (*IndexExpr) exprNode()     {}
func (*CastExpr) exprNode()      {}
func (*UnaryExpr) exprNode()     {}
func (*BinaryOpExpr) exprNode()  {}
func (*AssignExpr) exprNode()    {}
func (*AssignOpExpr) exprNode()  {}
func (*AssignRefExpr) exprNode() {}

func (*IssetExpr) exprNode()         {}
func (*EmptyExpr) exprNode()         {}
func (*EvalExpr) exprNode()          {}
func (*IncludeExpr) exprNode()       {}
func (*CloneExpr) exprNode()         {}
func (*ErrorSuppressExpr) exprNode() {}
func (*ExitExpr) exprNode()          {}

func (*ConstFetchExpr) exprNode()      {}
func (*ClassConstFetchExpr) exprNode() {}
func (*MagicConstExpr) exprNode()      {}

func (*InstanceofExpr) exprNode()          {}
func (*ListExpr) exprNode()                {}
func (*PrintExpr) exprNode()               {}
func (*PropertyFetchExpr) exprNode()       {}
func (*StaticPropertyFetchExpr) exprNode() {}
func (*ShellExecExpr) exprNode()           {}
func (*TernaryExpr) exprNode()             {}
func (*ThrowExpr) exprNode()               {}
func (*VariableExpr) exprNode()            {}
func (*YieldExpr) exprNode()               {}
func (*YieldFromExpr) exprNode()           {}

func (*FuncCallExpr) exprNode()   {}
func (*NewExpr) exprNode()        {}
func (*MethodCallExpr) exprNode() {}
func (*StaticCallExpr) exprNode() {}

// Stmt
func (*EmptyStmt) stmtNode()  {}
func (*BlockStmt) stmtNode()  {}
func (*ExprStmt) stmtNode()   {}
func (*ReturnStmt) stmtNode() {}

func (*LabelStmt) stmtNode()    {}
func (*GotoStmt) stmtNode()     {}
func (*IfStmt) stmtNode()       {}
func (*ElseIfStmt) stmtNode()   {}
func (*ElseStmt) stmtNode()     {}
func (*SwitchStmt) stmtNode()   {}
func (*CaseStmt) stmtNode()     {}
func (*ForStmt) stmtNode()      {}
func (*ForeachStmt) stmtNode()  {}
func (*BreakStmt) stmtNode()    {}
func (*ContinueStmt) stmtNode() {}
func (*WhileStmt) stmtNode()    {}
func (*DoStmt) stmtNode()       {}
func (*TryCatchStmt) stmtNode() {}
func (*CatchStmt) stmtNode()    {}
func (*FinallyStmt) stmtNode()  {}

func (*ConstStmt) stmtNode()        {}
func (*EchoStmt) stmtNode()         {}
func (*GlobalStmt) stmtNode()       {}
func (*HaltCompilerStmt) stmtNode() {}
func (*InlineHTMLStmt) stmtNode()   {}
func (*StaticStmt) stmtNode()       {}
func (*StaticVarStmt) stmtNode()    {}

func (*UnsetStmt) stmtNode()          {}
func (*UseStmt) stmtNode()            {}
func (*DeclareStmt) stmtNode()        {}
func (*DeclareDeclareStmt) stmtNode() {}

func (*NamespaceStmt) stmtNode()                    {}
func (*FunctionStmt) stmtNode()                     {}
func (*InterfaceStmt) stmtNode()                    {}
func (*ClassStmt) stmtNode()                        {}
func (*ClassConstStmt) stmtNode()                   {}
func (*PropertyStmt) stmtNode()                     {}
func (*ClassMethodStmt) stmtNode()                  {}
func (*TraitStmt) stmtNode()                        {}
func (*TraitUseStmt) stmtNode()                     {}
func (*TraitUseAdaptationAliasStmt) stmtNode()      {}
func (*TraitUseAdaptationPrecedenceStmt) stmtNode() {}

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

// All Node types
func (*Arg) node()                              {}
func (*Ident) node()                            {}
func (*Param) node()                            {}
func (*SimpleType) node()                       {}
func (*IntersectionType) node()                 {}
func (*UnionType) node()                        {}
func (*NullableType) node()                     {}
func (*Name) node()                             {}
func (*IntLit) node()                           {}
func (*FloatLit) node()                         {}
func (*StringLit) node()                        {}
func (*ArrayExpr) node()                        {}
func (*ArrayItemExpr) node()                    {}
func (*ClosureExpr) node()                      {}
func (*ClosureUseExpr) node()                   {}
func (*ArrowFunctionExpr) node()                {}
func (*IndexExpr) node()                        {}
func (*CastExpr) node()                         {}
func (*UnaryExpr) node()                        {}
func (*BinaryOpExpr) node()                     {}
func (*AssignExpr) node()                       {}
func (*AssignOpExpr) node()                     {}
func (*AssignRefExpr) node()                    {}
func (*IssetExpr) node()                        {}
func (*EmptyExpr) node()                        {}
func (*EvalExpr) node()                         {}
func (*IncludeExpr) node()                      {}
func (*CloneExpr) node()                        {}
func (*ErrorSuppressExpr) node()                {}
func (*ExitExpr) node()                         {}
func (*ConstFetchExpr) node()                   {}
func (*ClassConstFetchExpr) node()              {}
func (*MagicConstExpr) node()                   {}
func (*InstanceofExpr) node()                   {}
func (*ListExpr) node()                         {}
func (*PrintExpr) node()                        {}
func (*PropertyFetchExpr) node()                {}
func (*StaticPropertyFetchExpr) node()          {}
func (*ShellExecExpr) node()                    {}
func (*TernaryExpr) node()                      {}
func (*ThrowExpr) node()                        {}
func (*VariableExpr) node()                     {}
func (*YieldExpr) node()                        {}
func (*YieldFromExpr) node()                    {}
func (*FuncCallExpr) node()                     {}
func (*NewExpr) node()                          {}
func (*MethodCallExpr) node()                   {}
func (*StaticCallExpr) node()                   {}
func (*EmptyStmt) node()                        {}
func (*BlockStmt) node()                        {}
func (*ExprStmt) node()                         {}
func (*ReturnStmt) node()                       {}
func (*LabelStmt) node()                        {}
func (*GotoStmt) node()                         {}
func (*IfStmt) node()                           {}
func (*ElseIfStmt) node()                       {}
func (*ElseStmt) node()                         {}
func (*SwitchStmt) node()                       {}
func (*CaseStmt) node()                         {}
func (*ForStmt) node()                          {}
func (*ForeachStmt) node()                      {}
func (*BreakStmt) node()                        {}
func (*ContinueStmt) node()                     {}
func (*WhileStmt) node()                        {}
func (*DoStmt) node()                           {}
func (*TryCatchStmt) node()                     {}
func (*CatchStmt) node()                        {}
func (*FinallyStmt) node()                      {}
func (*ConstStmt) node()                        {}
func (*EchoStmt) node()                         {}
func (*GlobalStmt) node()                       {}
func (*HaltCompilerStmt) node()                 {}
func (*InlineHTMLStmt) node()                   {}
func (*StaticStmt) node()                       {}
func (*StaticVarStmt) node()                    {}
func (*UnsetStmt) node()                        {}
func (*UseStmt) node()                          {}
func (*DeclareStmt) node()                      {}
func (*DeclareDeclareStmt) node()               {}
func (*NamespaceStmt) node()                    {}
func (*FunctionStmt) node()                     {}
func (*InterfaceStmt) node()                    {}
func (*ClassStmt) node()                        {}
func (*ClassConstStmt) node()                   {}
func (*PropertyStmt) node()                     {}
func (*ClassMethodStmt) node()                  {}
func (*TraitStmt) node()                        {}
func (*TraitUseStmt) node()                     {}
func (*TraitUseAdaptationAliasStmt) node()      {}
func (*TraitUseAdaptationPrecedenceStmt) node() {}

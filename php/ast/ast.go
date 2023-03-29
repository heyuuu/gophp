package ast

import (
	"gophp/php/token"
	"strings"
)

type (
	Node interface{}

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

	// StmtClassLike : Stmt
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
		Name   *Ident // @var Ident|null Parameter name (for named parameters)
		Value  Expr   // @var Expr Value to pass
		ByRef  bool   // @var bool Whether to pass by ref
		Unpack bool   // @var bool Whether to unpack the argument
	}

	Attribute struct {
		Name *Name  // @var Name Attribute name
		Args []*Arg // @var Arg[] Attribute arguments
	}

	AttributeGroup struct {
		Attrs []*Attribute // @var Attribute[] Attributes
	}

	Const struct {
		Name           *Ident // @var Ident Name
		Value          Expr   // @var Expr Value
		NamespacedName *Name  // @var Name|null Namespaced name (if using NameResolver)
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

	MatchArm struct {
		Conds []Expr // @var Expr|null[]
		Body  Expr   // @var Expr
	}

	Param struct {
		Type       Type              // @var Type|null Type declaration
		ByRef      bool              // @var bool Whether parameter is passed by reference
		Variadic   bool              // @var bool Whether this is a variadic argument
		Var        *VariableExpr     // @var VariableExpr Parameter variable
		Default    Expr              // @var Expr|null Default value
		Flags      int               // @var int
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	// VariadicPlaceholder : PhpParserNodeAbstract
	VariadicPlaceholder struct{}
)

/**
 * Type
 */
type Type interface {
	Node
	typeNode()
}
type (
	// IntersectionType : A
	SimpleType struct {
		Name *Name
	}

	// IntersectionType : A & B & C
	IntersectionType struct {
		Types []Type // possible type: SimpleType
	}

	// UnionType : A | B | C
	UnionType struct {
		Types []Type // possible type: SimpleType or IntersectionType
	}

	// NullableType : ?A
	NullableType struct {
		Type *SimpleType // possible type: SimpleType
	}
)

/**
 *	Name
 */
type NameType int

const (
	NameNormal NameType = iota
	NameFullyQualified
	NameRelative
)

type Name struct {
	Kind  NameType // 0 normal, 1 full-qualified, 2 relative
	Parts []string // @var string[] Parts of the name
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
	default:
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
		Value float64 // @var float Number value
	}

	StringLit struct {
		Value string // @var string String value
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

	// ExprClosure : Expr, FunctionLike
	ClosureExpr struct {
		Static     bool              // @var bool Whether the closure is static
		ByRef      bool              // @var bool Whether to return by reference
		Params     []*Param          // @var Param[] Parameters
		Uses       []*ClosureUseExpr // @var ClosureUse[] use()s
		ReturnType Type              // @var Type|null Return type
		Stmts      []Stmt            // @var Stmt[] Statements
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	ClosureUseExpr struct {
		Var   *VariableExpr // @var VariableExpr Variable to use
		ByRef bool          // @var bool Whether to use by reference
	}

	// ExprArrowFunction : Expr, FunctionLike
	ArrowFunctionExpr struct {
		Static     bool              // @var bool
		ByRef      bool              // @var bool
		Params     []*Param          // @var Param[]
		ReturnType Type              // @var Type|null
		Expr       Expr              // @var Expr
		AttrGroups []*AttributeGroup // @var AttributeGroup[]
	}

	// IndexExpr
	IndexExpr struct {
		Var Expr // @var Expr       Variable
		Dim Expr // @var Expr|null  Array index / dim
	}

	// CastExpr
	CastExpr struct {
		Op   token.Token // token.
		Expr Expr        // @var Expr Expression
	}

	// UnaryExpr
	UnaryExpr struct {
		Kind token.Token // token.Add, token.Sub, token.Not, token.Tilde, token.PreInc, token.PreDec, token.PostInc or token.PostDec
		Var  Expr        // variable
	}

	// BinaryExpr
	BinaryExpr struct {
		Op    token.Token // token.IsBinaryOp()
		Left  Expr        // @var Expr The left-hand side expression
		Right Expr        // @var Expr The right-hand side expression
	}

	// AssignExpr
	AssignExpr struct {
		Op   token.Token // token.IsAssignOp()
		Var  Expr        // @var Expr Variable
		Expr Expr        // @var Expr Expression
	}

	AssignRefExpr struct {
		Var  Expr // @var Expr Variable reference is assigned to
		Expr Expr // @var Expr Variable which is referenced
	}

	// InternalCallExpr
	InternalCallExpr struct {
		Kind token.Token // token.IsInternalCall()
		Args []Expr      // arguments
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

	// Const
	ConstFetchExpr struct {
		Name *Name // @var Name Constant name
	}

	ClassConstFetchExpr struct {
		Class Expr   // @var Name|Expr Class name
		Name  *Ident // @var Ident Constant name
	}

	MagicConstExpr struct {
		Kind token.Token // token.IsMagicConstKind()
	}

	// ExprMatch : Expr
	MatchExpr struct {
		Cond Expr        // @var Expr
		Arms []*MatchArm // @var MatchArm[]
	}

	InstanceofExpr struct {
		Expr  Expr // @var Expr Expression
		Class Expr // @var Name|Expr Class name
	}

	ListExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	}

	PrintExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprPropertyFetch : Expr
	PropertyFetchExpr struct {
		Var  Expr // @var Expr Variable holding object
		Name any  // @var Ident|Expr Property name
	}

	NullsafePropertyFetchExpr struct {
		Var  Expr // @var Expr Variable holding object
		Name any  // @var Ident|Expr Property name
	}

	// ExprStaticPropertyFetch : Expr
	StaticPropertyFetchExpr struct {
		Class any // @var Name|Expr Class name
		Name  any // @var VarLikeIdent|Expr Property name
	}

	// ExprShellExec : Expr
	ShellExecExpr struct {
		Parts []any // @var array Encapsed string array
	}

	// ExprTernary : Expr
	TernaryExpr struct {
		Cond Expr // @var Expr Condition
		If   Expr // @var Expr|null Expression for true
		Else Expr // @var Expr Expression for false
	}

	// ExprThrow : Expr
	ThrowExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprVariable : Expr
	VariableExpr struct {
		Name any // @var string|Expr Name
	}

	// ExprYield : Expr
	YieldExpr struct {
		Key   Expr // @var Expr|null Key expression
		Value Expr // @var Expr|null Value expression
	}

	// ExprYieldFrom : Expr
	YieldFromExpr struct {
		Expr Expr // @var Expr Expression to yield from
	}

	// FuncCallExpr : Expr, CallLikeExpr
	FuncCallExpr struct {
		Name any   // @var Name|Expr Function name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// NewExpr : CallLikeExpr
	NewExpr struct {
		Class Node  // @var Name|Expr|ClassStmt Class name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// MethodCallExpr : CallLikeExpr
	MethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Ident|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// NullsafeMethodCallExpr : CallLikeExpr
	NullsafeMethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Ident|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprStaticCall : ExprCallLike
	StaticCallExpr struct {
		Class any   // @var Name|Expr Class name
		Name  any   // @var Ident|Expr Method name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
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
		Elseifs []*ElseIfStmt // @var ElseIf_[] 	elseif branches
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
		Consts []*Const // @var Const_[] Constant declarations
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

	ThrowStmt struct {
		Expr Expr // @var Expr Expression
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

	// StmtFunction : Stmt, FunctionLike
	FunctionStmt struct {
		ByRef          bool              // @var bool Whether function returns by reference
		Name           *Ident            // @var Ident Name
		Params         []*Param          // @var Param[] Parameters
		ReturnType     Type              // @var Type|null Return type
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	// InterfaceStmt
	InterfaceStmt struct {
		Extends        []*Name           // @var Name[] Extended interfaces
		Name           *Ident            // @var Ident|null Name
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	// StmtClass : Stmt, StmtClassLike
	ClassStmt struct {
		Flags          int               // @var int        Type
		Extends        *Name             // @var Name|null  Name of extended class
		Implements     []*Name           // @var Name[]     Names of implemented interfaces
		Name           *Ident            // @var Ident|null Name
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	// StmtClassConst : Stmt
	ClassConstStmt struct {
		Flags      int               // @var int Modifiers
		Consts     []*Const          // @var Const_[] Constant declarations
		AttrGroups []*AttributeGroup // @var AttributeGroup[]
	}

	// PropertyStmt : Stmt
	PropertyStmt struct {
		Flags      int                     // @var int Modifiers
		Props      []*PropertyPropertyStmt // @var PropertyProperty[] Properties
		Type       Type                    // @var Type|null Type declaration
		AttrGroups []*AttributeGroup       // @var AttributeGroup[] PHP attribute groups
	}

	PropertyPropertyStmt struct {
		Name    *Ident // @var Ident     Name
		Default Expr   // @var Expr|null Default
	}

	// StmtClassMethod : Stmt, FunctionLike
	ClassMethodStmt struct {
		Flags      int               // @var int Flags
		ByRef      bool              // @var bool Whether to return by reference
		Name       *Ident            // @var Ident Name
		Params     []*Param          // @var Param[] Parameters
		ReturnType Type              // @var Type|null Return type
		Stmts      []Stmt            // @var Stmt[]|null Statements
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	// StmtTrait : StmtClassLike
	TraitStmt struct {
		Name           *Ident            // @var Ident|null Name
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	TraitUseStmt struct {
		Traits      []*Name                   // @var Name[] Traits
		Adaptations []*TraitUseAdaptationStmt // @var TraitUseAdaptation[] Adaptations
	}

	// StmtTraitUseAdaptationAlias : StmtTraitUseAdaptation
	TraitUseAdaptationAliasStmt struct {
		NewModifier int    // @var int|null New modifier
		NewName     *Ident // @var Ident|null New name
		Trait       *Name  // @var Name|null Trait name
		Method      *Ident // @var Ident Method name
	}

	// StmtTraitUseAdaptationPrecedence : StmtTraitUseAdaptation
	TraitUseAdaptationPrecedenceStmt struct {
		Insteadof []*Name // @var Name[] Overwritten traits
		Trait     *Name   // @var Name|null Trait name
		Method    *Ident  // @var Ident Method name
	}

	// StmtEnum : StmtClassLike
	EnumStmt struct {
		ScalarType     *Ident            // @var Ident|null Scalar Type
		Implements     []*Name           // @var Name[] Names of implemented interfaces
		Name           *Ident            // @var Ident|null Name
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	EnumCaseStmt struct {
		Name       *Ident            // @var Ident Enum case name
		Expr       Expr              // @var Expr|null Enum case expression
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}
)

// Type
func (*SimpleType) typeNode()       {}
func (*IntersectionType) typeNode() {}
func (*UnionType) typeNode()        {}
func (*NullableType) typeNode()     {}

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
func (*BinaryExpr) exprNode()    {}
func (*AssignExpr) exprNode()    {}
func (*AssignRefExpr) exprNode() {}

func (*InternalCallExpr) exprNode()  {}
func (*CloneExpr) exprNode()         {}
func (*ErrorSuppressExpr) exprNode() {}
func (*ExitExpr) exprNode()          {}

func (*ConstFetchExpr) exprNode()      {}
func (*ClassConstFetchExpr) exprNode() {}
func (*MagicConstExpr) exprNode()      {}

func (*MatchExpr) exprNode()                 {}
func (*InstanceofExpr) exprNode()            {}
func (*ListExpr) exprNode()                  {}
func (*PrintExpr) exprNode()                 {}
func (*PropertyFetchExpr) exprNode()         {}
func (*NullsafePropertyFetchExpr) exprNode() {}
func (*StaticPropertyFetchExpr) exprNode()   {}
func (*ShellExecExpr) exprNode()             {}
func (*TernaryExpr) exprNode()               {}
func (*ThrowExpr) exprNode()                 {}
func (*VariableExpr) exprNode()              {}
func (*YieldExpr) exprNode()                 {}
func (*YieldFromExpr) exprNode()             {}

func (*FuncCallExpr) exprNode()           {}
func (*NewExpr) exprNode()                {}
func (*MethodCallExpr) exprNode()         {}
func (*NullsafeMethodCallExpr) exprNode() {}
func (*StaticCallExpr) exprNode()         {}

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

func (*ThrowStmt) stmtNode()          {}
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
func (*PropertyPropertyStmt) stmtNode()             {}
func (*ClassMethodStmt) stmtNode()                  {}
func (*TraitStmt) stmtNode()                        {}
func (*TraitUseStmt) stmtNode()                     {}
func (*TraitUseAdaptationAliasStmt) stmtNode()      {}
func (*TraitUseAdaptationPrecedenceStmt) stmtNode() {}
func (*EnumStmt) stmtNode()                         {}
func (*EnumCaseStmt) stmtNode()                     {}

// CallLikeExpr
func (*FuncCallExpr) callLikeExprNode()           {}
func (*NewExpr) callLikeExprNode()                {}
func (*MethodCallExpr) callLikeExprNode()         {}
func (*NullsafeMethodCallExpr) callLikeExprNode() {}
func (*StaticCallExpr) callLikeExprNode()         {}

// FunctionLike
func (*ArrowFunctionExpr) functionLikeNode() {}
func (*ClosureExpr) functionLikeNode()       {}
func (*ClassMethodStmt) functionLikeNode()   {}
func (*FunctionStmt) functionLikeNode()      {}

// ClassLikeStmt
func (*ClassStmt) classLikeStmtNode()     {}
func (*EnumStmt) classLikeStmtNode()      {}
func (*InterfaceStmt) classLikeStmtNode() {}
func (*TraitStmt) classLikeStmtNode()     {}

// TraitUseAdaptationStmt
func (*TraitUseAdaptationAliasStmt) traitUseAdaptationStmtNode()      {}
func (*TraitUseAdaptationPrecedenceStmt) traitUseAdaptationStmtNode() {}

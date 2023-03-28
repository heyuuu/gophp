package ast

import (
	"gophp/php/token"
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

	// ComplexType : PhpParserNodeAbstract
	ComplexType interface {
		Node
		complexTypeNode()
	}

	// CallLikeExpr : Expr
	CallLikeExpr interface {
		Expr
		callLikeExprNode()
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
	}

	// IntersectionType : ComplexType
	IntersectionType struct {
		Types []any // @var (Ident|Name)[] Types
	}

	MatchArm struct {
		Conds []Expr // @var Expr|null[]
		Body  Expr   // @var Expr
	}

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
		Type any // @var Ident|Name Type
	}

	Param struct {
		Type       any               // @var Ident|Name|ComplexType|null Type declaration
		ByRef      bool              // @var bool Whether parameter is passed by reference
		Variadic   bool              // @var bool Whether this is a variadic argument
		Var        VariableExpr      // @var VariableExpr Parameter variable
		Default    Expr              // @var Expr|null Default value
		Flags      int               // @var int
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	// UnionType : ComplexType
	UnionType struct {
		Types []any // @var (Ident|Name|IntersectionType)[] Types
	}

	// VarLikeIdent : Ident
	VarLikeIdentifier struct {
		Name string // @var string Ident as string
	}

	// VariadicPlaceholder : PhpParserNodeAbstract
	VariadicPlaceholder struct {
	}
)

// Expr
type (
	BinaryExpr struct {
		Op    token.Token
		Left  Expr // @var Expr The left-hand side expression
		Right Expr // @var Expr The right-hand side expression
	}

	ArrayExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] Items
	}

	ArrayDimFetchExpr struct {
		Var Expr // @var Expr Variable
		Dim Expr // @var Expr|null Array index / dim
	}

	ArrayItemExpr struct {
		Key    Expr // @var Expr|null Key
		Value  Expr // @var Expr Value
		ByRef  bool // @var bool Whether to assign by reference
		Unpack bool // @var bool Whether to unpack the argument
	}

	// ExprArrowFunction : Expr, FunctionLike
	ArrowFunctionExpr struct {
		Static     bool              // @var bool
		ByRef      bool              // @var bool
		Params     []*Param          // @var Param[]
		ReturnType any               // @var Ident|Name|ComplexType|null
		Expr       Expr              // @var Expr
		AttrGroups []*AttributeGroup // @var AttributeGroup[]
	}

	AssignExpr struct {
		Op   token.Token //
		Var  Expr        // @var Expr Variable
		Expr Expr        // @var Expr Expression
	}

	AssignRefExpr struct {
		Var  Expr // @var Expr Variable reference is assigned to
		Expr Expr // @var Expr Variable which is referenced
	}

	BitwiseNotExpr struct {
		Expr Expr // @var Expr Expression
	}

	BooleanNotExpr struct {
		Expr Expr // @var Expr Expression
	}

	CastExpr struct {
		Op   token.Token
		Expr Expr // @var Expr Expression
	}

	ClassConstFetchExpr struct {
		Class any    // @var Name|Expr Class name
		Name  *Ident // @var Ident Constant name
	}

	CloneExpr struct {
		Expr Expr // @var Expr Expression
	}

	// ExprClosure : Expr, FunctionLike
	ClosureExpr struct {
		Static     bool              // @var bool Whether the closure is static
		ByRef      bool              // @var bool Whether to return by reference
		Params     []*Param          // @var Param[] Parameters
		Uses       []*ClosureUseExpr // @var ClosureUse[] use()s
		ReturnType any               // @var Ident|Name|ComplexType|null Return type
		Stmts      []Stmt            // @var Stmt[] Statements
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	ClosureUseExpr struct {
		Var   *VariableExpr // @var VariableExpr Variable to use
		ByRef bool          // @var bool Whether to use by reference
	}

	ConstFetchExpr struct {
		Name *Name // @var Name Constant name
	}

	EmptyExpr struct {
		Expr Expr // @var Expr Expression
	}

	ErrorSuppressExpr struct {
		Expr Expr // @var Expr Expression
	}

	EvalExpr struct {
		Expr Expr // @var Expr Expression
	}

	ExitExpr struct {
		Expr Expr // @var Expr|null Expression
	}

	// ExprFuncCall : ExprCallLike
	FuncCallExpr struct {
		Name any   // @var Name|Expr Function name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	IncludeExpr struct {
		Expr Expr // @var Expr Expression
		Type int  // @var int Type of include
	}

	InstanceofExpr struct {
		Expr  Expr // @var Expr Expression
		Class any  // @var Name|Expr Class name
	}

	IssetExpr struct {
		Vars []Expr // @var Expr[] Variables
	}

	// ExprList : Expr
	ListExpr struct {
		Items []*ArrayItemExpr // @var (ArrayItem|null)[] List of items to assign to
	}

	// ExprMatch : Expr
	MatchExpr struct {
		Cond Expr        // @var Expr
		Arms []*MatchArm // @var MatchArm[]
	}

	// ExprMethodCall : ExprCallLike
	MethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Ident|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNew : ExprCallLike
	NewExpr struct {
		Class any   // @var Name|Expr|ClassStmt Class name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNullsafeMethodCall : ExprCallLike
	NullsafeMethodCallExpr struct {
		Var  Expr  // @var Expr Variable holding object
		Name any   // @var Ident|Expr Method name
		Args []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprNullsafePropertyFetch : Expr
	NullsafePropertyFetchExpr struct {
		Var  Expr // @var Expr Variable holding object
		Name any  // @var Ident|Expr Property name
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
		Name any  // @var Ident|Expr Property name
	}

	// ExprShellExec : Expr
	ShellExecExpr struct {
		Parts []any // @var array Encapsed string array
	}

	// ExprStaticCall : ExprCallLike
	StaticCallExpr struct {
		Class any   // @var Name|Expr Class name
		Name  any   // @var Ident|Expr Method name
		Args  []any // @var array<Arg|VariadicPlaceholder> Arguments
	}

	// ExprStaticPropertyFetch : Expr
	StaticPropertyFetchExpr struct {
		Class any // @var Name|Expr Class name
		Name  any // @var VarLikeIdent|Expr Property name
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
		Key   Expr // @var Expr|null Key expression
		Value Expr // @var Expr|null Value expression
	}

	// ExprYieldFrom : Expr
	YieldFromExpr struct {
		Expr Expr // @var Expr Expression to yield from
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

	MagicConstScalar struct {
		Op token.Token
	}

	// ScalarString : Scalar
	StringScalar struct {
		Value string // @var string String value
	}
)

// Stmt
type (
	EmptyStmt struct{}

	LabelStmt struct {
		Name *Ident // @var Ident Name
	}

	ExprStmt struct {
		Expr Expr // @var Expr Expression
	}

	ReturnStmt struct {
		Expr Expr // @var Expr|null Expression
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

	CatchStmt struct {
		Types []*Name       // @var Name[] Types of exceptions to catch
		Var   *VariableExpr // @var VariableExpr|null Variable for exception
		Stmts []Stmt        // @var Stmt[] Statements
	}

	// StmtClass : StmtClassLike
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

	// StmtClassMethod : Stmt, FunctionLike
	ClassMethodStmt struct {
		Flags      int               // @var int Flags
		ByRef      bool              // @var bool Whether to return by reference
		Name       *Ident            // @var Ident Name
		Params     []*Param          // @var Param[] Parameters
		ReturnType any               // @var Ident|Name|ComplexType|null Return type
		Stmts      []Stmt            // @var Stmt[]|null Statements
		AttrGroups []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
	}

	ConstStmt struct {
		Consts []*Const // @var Const_[] Constant declarations
	}

	ContinueStmt struct {
		Num Expr // @var Expr|null Number of loops to continue
	}

	DeclareStmt struct {
		Declares []*DeclareDeclareStmt // @var DeclareDeclare[] List of declares
		Stmts    []Stmt                // @var Stmt[]|null Statements
	}

	DeclareDeclareStmt struct {
		Key   *Ident // @var Ident Key
		Value Expr   // @var Expr Value
	}

	DoStmt struct {
		Stmts []Stmt // @var Stmt[] Statements
		Cond  Expr   // @var Expr Condition
	}

	// StmtEcho : Stmt
	EchoStmt struct {
		Exprs []Expr // @var Expr[] Expressions
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

	FinallyStmt struct {
		Stmts []Stmt // @var Stmt[] Statements
	}

	// StmtFunction : Stmt, FunctionLike
	FunctionStmt struct {
		ByRef          bool              // @var bool Whether function returns by reference
		Name           *Ident            // @var Ident Name
		Params         []*Param          // @var Param[] Parameters
		ReturnType     any               // @var Ident|Name|ComplexType|null Return type
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	GlobalStmt struct {
		Vars []Expr // @var Expr[] Variables
	}

	GotoStmt struct {
		Name *Ident // @var Ident Name of label to jump to
	}

	GroupUseStmt struct {
		Type   int           // @var int Type of group use
		Prefix *Name         // @var Name Prefix for uses
		Uses   []*UseUseStmt // @var UseUse[] Uses
	}

	HaltCompilerStmt struct {
		Remaining string // @var string Remaining text after halt compiler statement.
	}

	InlineHTMLStmt struct {
		Value string // @var string String
	}

	InterfaceStmt struct {
		Extends        []*Name           // @var Name[] Extended interfaces
		Name           *Ident            // @var Ident|null Name
		Stmts          []Stmt            // @var Stmt[] Statements
		AttrGroups     []*AttributeGroup // @var AttributeGroup[] PHP attribute groups
		NamespacedName *Name             // @var Name|null Namespaced name (if using NameResolver)
	}

	NamespaceStmt struct {
		Name  *Name  // @var Name|null Name
		Stmts []Stmt // @var Stmt[] Statements
	}

	PropertyStmt struct {
		Flags      int                     // @var int Modifiers
		Props      []*PropertyPropertyStmt // @var PropertyProperty[] Properties
		Type       any                     // @var Ident|Name|ComplexType|null Type declaration
		AttrGroups []*AttributeGroup       // @var AttributeGroup[] PHP attribute groups
	}

	PropertyPropertyStmt struct {
		Name    *VarLikeIdentifier // @var VarLikeIdent Name
		Default Expr               // @var Expr|null Default
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

	TryCatchStmt struct {
		Stmts   []Stmt       // @var Stmt[] Statements
		Catches []*CatchStmt // @var Catch_[] Catches
		Finally *FinallyStmt // @var Finally_|null Optional finally node
	}

	UnsetStmt struct {
		Vars []Expr // @var Expr[] Variables to unset
	}

	UseStmt struct {
		Type int           // @var int Type of alias
		Uses []*UseUseStmt // @var UseUse[] Aliases
	}

	UseUseStmt struct {
		Type  int    // @var int One of the Stmt\Use_::TYPE_* constants. Will only differ from TYPE_UNKNOWN for mixed group uses
		Name  *Name  // @var Name Namespace, class, function or constant to alias
		Alias *Ident // @var Ident|null Alias
	}

	WhileStmt struct {
		Cond  Expr   // @var Expr Condition
		Stmts []Stmt // @var Stmt[] Statements
	}
)

// Expr
func (*ArrayExpr) exprNode()                 {}
func (*ArrayDimFetchExpr) exprNode()         {}
func (*ArrayItemExpr) exprNode()             {}
func (*ArrowFunctionExpr) exprNode()         {}
func (*AssignExpr) exprNode()                {}
func (*AssignRefExpr) exprNode()             {}
func (*BinaryExpr) exprNode()                {}
func (*BitwiseNotExpr) exprNode()            {}
func (*BooleanNotExpr) exprNode()            {}
func (*CastExpr) exprNode()                  {}
func (*ClassConstFetchExpr) exprNode()       {}
func (*CloneExpr) exprNode()                 {}
func (*ClosureExpr) exprNode()               {}
func (*ClosureUseExpr) exprNode()            {}
func (*ConstFetchExpr) exprNode()            {}
func (*EmptyExpr) exprNode()                 {}
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

// ExprCallLike
func (*FuncCallExpr) exprCallLikeNode()           {}
func (*MethodCallExpr) exprCallLikeNode()         {}
func (*NewExpr) exprCallLikeNode()                {}
func (*NullsafeMethodCallExpr) exprCallLikeNode() {}
func (*StaticCallExpr) exprCallLikeNode()         {}

func (*FuncCallExpr) exprNode()           {}
func (*MethodCallExpr) exprNode()         {}
func (*NewExpr) exprNode()                {}
func (*NullsafeMethodCallExpr) exprNode() {}
func (*StaticCallExpr) exprNode()         {}

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
func (*MagicConstScalar) scalarNode()         {}
func (*StringScalar) scalarNode()             {}

func (*DNumberScalar) exprNode()            {}
func (*EncapsedScalar) exprNode()           {}
func (*EncapsedStringPartScalar) exprNode() {}
func (*LNumberScalar) exprNode()            {}
func (*MagicConstScalar) exprNode()         {}
func (*StringScalar) exprNode()             {}

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
func (*ExprStmt) stmtNode()             {}
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
func (*EmptyStmt) stmtNode()            {}
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

// Ident
func (*VarLikeIdentifier) identifierNode() {}

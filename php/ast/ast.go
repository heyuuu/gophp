package ast

type Node interface {}

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
    Name any
    Value any
    ByRef any
    Unpack any
    Attributes any
}

// Attribute : PhpParserNodeAbstract
type Attribute struct {
    Name any
    Args any
    Attributes any
}

// AttributeGroup : PhpParserNodeAbstract
type AttributeGroup struct {
    Attrs any
    Attributes any
}

// Const : PhpParserNodeAbstract
type Const struct {
    Name any
    Value any
    NamespacedName any
    Attributes any
}

// ExprArray : Expr
type ArrayExpr struct {
    Items any
    Attributes any
}

// ExprArrayDimFetch : Expr
type ArrayDimFetchExpr struct {
    Var any
    Dim any
    Attributes any
}

// ExprArrayItem : Expr
type ArrayItemExpr struct {
    Key any
    Value any
    ByRef any
    Unpack any
    Attributes any
}

// ExprArrowFunction : Expr, FunctionLike
type ArrowFunctionExpr struct {
    Static any
    ByRef any
    Params any
    ReturnType any
    Expr any
    AttrGroups any
    Attributes any
}

// ExprAssign : Expr
type AssignExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpBitwiseAnd : ExprAssignOp
type AssignOpBitwiseAndExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpBitwiseOr : ExprAssignOp
type AssignOpBitwiseOrExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpBitwiseXor : ExprAssignOp
type AssignOpBitwiseXorExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpCoalesce : ExprAssignOp
type AssignOpCoalesceExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpConcat : ExprAssignOp
type AssignOpConcatExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpDiv : ExprAssignOp
type AssignOpDivExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpMinus : ExprAssignOp
type AssignOpMinusExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpMod : ExprAssignOp
type AssignOpModExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpMul : ExprAssignOp
type AssignOpMulExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpPlus : ExprAssignOp
type AssignOpPlusExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpPow : ExprAssignOp
type AssignOpPowExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpShiftLeft : ExprAssignOp
type AssignOpShiftLeftExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignOpShiftRight : ExprAssignOp
type AssignOpShiftRightExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprAssignRef : Expr
type AssignRefExpr struct {
    Var any
    Expr any
    Attributes any
}

// ExprBinaryOpBitwiseAnd : ExprBinaryOp
type BinaryOpBitwiseAndExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpBitwiseOr : ExprBinaryOp
type BinaryOpBitwiseOrExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpBitwiseXor : ExprBinaryOp
type BinaryOpBitwiseXorExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpBooleanAnd : ExprBinaryOp
type BinaryOpBooleanAndExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpBooleanOr : ExprBinaryOp
type BinaryOpBooleanOrExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpCoalesce : ExprBinaryOp
type BinaryOpCoalesceExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpConcat : ExprBinaryOp
type BinaryOpConcatExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpDiv : ExprBinaryOp
type BinaryOpDivExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpEqual : ExprBinaryOp
type BinaryOpEqualExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpGreater : ExprBinaryOp
type BinaryOpGreaterExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpGreaterOrEqual : ExprBinaryOp
type BinaryOpGreaterOrEqualExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpIdentical : ExprBinaryOp
type BinaryOpIdenticalExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpLogicalAnd : ExprBinaryOp
type BinaryOpLogicalAndExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpLogicalOr : ExprBinaryOp
type BinaryOpLogicalOrExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpLogicalXor : ExprBinaryOp
type BinaryOpLogicalXorExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpMinus : ExprBinaryOp
type BinaryOpMinusExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpMod : ExprBinaryOp
type BinaryOpModExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpMul : ExprBinaryOp
type BinaryOpMulExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpNotEqual : ExprBinaryOp
type BinaryOpNotEqualExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpNotIdentical : ExprBinaryOp
type BinaryOpNotIdenticalExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpPlus : ExprBinaryOp
type BinaryOpPlusExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpPow : ExprBinaryOp
type BinaryOpPowExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpShiftLeft : ExprBinaryOp
type BinaryOpShiftLeftExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpShiftRight : ExprBinaryOp
type BinaryOpShiftRightExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpSmaller : ExprBinaryOp
type BinaryOpSmallerExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpSmallerOrEqual : ExprBinaryOp
type BinaryOpSmallerOrEqualExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBinaryOpSpaceship : ExprBinaryOp
type BinaryOpSpaceshipExpr struct {
    Left any
    Right any
    Attributes any
}

// ExprBitwiseNot : Expr
type BitwiseNotExpr struct {
    Expr any
    Attributes any
}

// ExprBooleanNot : Expr
type BooleanNotExpr struct {
    Expr any
    Attributes any
}

// ExprCastArray : ExprCast
type CastArrayExpr struct {
    Expr any
    Attributes any
}

// ExprCastBool : ExprCast
type CastBoolExpr struct {
    Expr any
    Attributes any
}

// ExprCastDouble : ExprCast
type CastDoubleExpr struct {
    Expr any
    Attributes any
}

// ExprCastInt : ExprCast
type CastIntExpr struct {
    Expr any
    Attributes any
}

// ExprCastObject : ExprCast
type CastObjectExpr struct {
    Expr any
    Attributes any
}

// ExprCastString : ExprCast
type CastStringExpr struct {
    Expr any
    Attributes any
}

// ExprCastUnset : ExprCast
type CastUnsetExpr struct {
    Expr any
    Attributes any
}

// ExprClassConstFetch : Expr
type ClassConstFetchExpr struct {
    Class any
    Name any
    Attributes any
}

// ExprClone : Expr
type CloneExpr struct {
    Expr any
    Attributes any
}

// ExprClosure : Expr, FunctionLike
type ClosureExpr struct {
    Static any
    ByRef any
    Params any
    Uses any
    ReturnType any
    Stmts any
    AttrGroups any
    Attributes any
}

// ExprClosureUse : Expr
type ClosureUseExpr struct {
    Var any
    ByRef any
    Attributes any
}

// ExprConstFetch : Expr
type ConstFetchExpr struct {
    Name any
    Attributes any
}

// ExprEmpty : Expr
type EmptyExpr struct {
    Expr any
    Attributes any
}

// ExprError : Expr
type ErrorExpr struct {
    Attributes any
}

// ExprErrorSuppress : Expr
type ErrorSuppressExpr struct {
    Expr any
    Attributes any
}

// ExprEval : Expr
type EvalExpr struct {
    Expr any
    Attributes any
}

// ExprExit : Expr
type ExitExpr struct {
    Expr any
    Attributes any
}

// ExprFuncCall : ExprCallLike
type FuncCallExpr struct {
    Name any
    Args any
    Attributes any
}

// ExprInclude : Expr
type IncludeExpr struct {
    Expr any
    Type any
    Attributes any
}

// ExprInstanceof : Expr
type InstanceofExpr struct {
    Expr any
    Class any
    Attributes any
}

// ExprIsset : Expr
type IssetExpr struct {
    Vars any
    Attributes any
}

// ExprList : Expr
type ListExpr struct {
    Items any
    Attributes any
}

// ExprMatch : Expr
type MatchExpr struct {
    Cond any
    Arms any
    Attributes any
}

// ExprMethodCall : ExprCallLike
type MethodCallExpr struct {
    Var any
    Name any
    Args any
    Attributes any
}

// ExprNew : ExprCallLike
type NewExpr struct {
    Class any
    Args any
    Attributes any
}

// ExprNullsafeMethodCall : ExprCallLike
type NullsafeMethodCallExpr struct {
    Var any
    Name any
    Args any
    Attributes any
}

// ExprNullsafePropertyFetch : Expr
type NullsafePropertyFetchExpr struct {
    Var any
    Name any
    Attributes any
}

// ExprPostDec : Expr
type PostDecExpr struct {
    Var any
    Attributes any
}

// ExprPostInc : Expr
type PostIncExpr struct {
    Var any
    Attributes any
}

// ExprPreDec : Expr
type PreDecExpr struct {
    Var any
    Attributes any
}

// ExprPreInc : Expr
type PreIncExpr struct {
    Var any
    Attributes any
}

// ExprPrint : Expr
type PrintExpr struct {
    Expr any
    Attributes any
}

// ExprPropertyFetch : Expr
type PropertyFetchExpr struct {
    Var any
    Name any
    Attributes any
}

// ExprShellExec : Expr
type ShellExecExpr struct {
    Parts any
    Attributes any
}

// ExprStaticCall : ExprCallLike
type StaticCallExpr struct {
    Class any
    Name any
    Args any
    Attributes any
}

// ExprStaticPropertyFetch : Expr
type StaticPropertyFetchExpr struct {
    Class any
    Name any
    Attributes any
}

// ExprTernary : Expr
type TernaryExpr struct {
    Cond any
    If any
    Else any
    Attributes any
}

// ExprThrow : Expr
type ThrowExpr struct {
    Expr any
    Attributes any
}

// ExprUnaryMinus : Expr
type UnaryMinusExpr struct {
    Expr any
    Attributes any
}

// ExprUnaryPlus : Expr
type UnaryPlusExpr struct {
    Expr any
    Attributes any
}

// ExprVariable : Expr
type VariableExpr struct {
    Name any
    Attributes any
}

// ExprYield : Expr
type YieldExpr struct {
    Key any
    Value any
    Attributes any
}

// ExprYieldFrom : Expr
type YieldFromExpr struct {
    Expr any
    Attributes any
}

// Identifier : PhpParserNodeAbstract, Stringable
type Identifier struct {
    Name any
    SpecialClassNames any
    Attributes any
}

// IntersectionType : ComplexType
type IntersectionType struct {
    Types any
    Attributes any
}

// MatchArm : PhpParserNodeAbstract
type MatchArm struct {
    Conds any
    Body any
    Attributes any
}

// Name : PhpParserNodeAbstract, Stringable
type Name struct {
    Parts any
    SpecialClassNames any
    Attributes any
}

// NameFullyQualified : Name, Stringable
type NameFullyQualified struct {
    Parts any
    Attributes any
}

// NameRelative : Name, Stringable
type NameRelative struct {
    Parts any
    Attributes any
}

// NullableType : ComplexType
type NullableType struct {
    Type any
    Attributes any
}

// Param : PhpParserNodeAbstract
type Param struct {
    Type any
    ByRef any
    Variadic any
    Var any
    Default any
    Flags any
    AttrGroups any
    Attributes any
}

// ScalarDNumber : Scalar
type DNumberScalar struct {
    Value any
    Attributes any
}

// ScalarEncapsed : Scalar
type EncapsedScalar struct {
    Parts any
    Attributes any
}

// ScalarEncapsedStringPart : Scalar
type EncapsedStringPartScalar struct {
    Value any
    Attributes any
}

// ScalarLNumber : Scalar
type LNumberScalar struct {
    Value any
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
    Value any
    Replacements any
    Attributes any
}

// StmtBreak : Stmt
type BreakStmt struct {
    Num any
    Attributes any
}

// StmtCase : Stmt
type CaseStmt struct {
    Cond any
    Stmts any
    Attributes any
}

// StmtCatch : Stmt
type CatchStmt struct {
    Types any
    Var any
    Stmts any
    Attributes any
}

// StmtClass : StmtClassLike
type ClassStmt struct {
    Flags any
    Extends any
    Implements any
    Name any
    Stmts any
    AttrGroups any
    NamespacedName any
    Attributes any
}

// StmtClassConst : Stmt
type ClassConstStmt struct {
    Flags any
    Consts any
    AttrGroups any
    Attributes any
}

// StmtClassMethod : Stmt, FunctionLike
type ClassMethodStmt struct {
    Flags any
    ByRef any
    Name any
    Params any
    ReturnType any
    Stmts any
    AttrGroups any
    MagicNames any
    Attributes any
}

// StmtConst : Stmt
type ConstStmt struct {
    Consts any
    Attributes any
}

// StmtContinue : Stmt
type ContinueStmt struct {
    Num any
    Attributes any
}

// StmtDeclare : Stmt
type DeclareStmt struct {
    Declares any
    Stmts any
    Attributes any
}

// StmtDeclareDeclare : Stmt
type DeclareDeclareStmt struct {
    Key any
    Value any
    Attributes any
}

// StmtDo : Stmt
type DoStmt struct {
    Stmts any
    Cond any
    Attributes any
}

// StmtEcho : Stmt
type EchoStmt struct {
    Exprs any
    Attributes any
}

// StmtElse : Stmt
type ElseStmt struct {
    Stmts any
    Attributes any
}

// StmtElseIf : Stmt
type ElseIfStmt struct {
    Cond any
    Stmts any
    Attributes any
}

// StmtEnum : StmtClassLike
type EnumStmt struct {
    ScalarType any
    Implements any
    Name any
    Stmts any
    AttrGroups any
    NamespacedName any
    Attributes any
}

// StmtEnumCase : Stmt
type EnumCaseStmt struct {
    Name any
    Expr any
    AttrGroups any
    Attributes any
}

// StmtExpression : Stmt
type ExpressionStmt struct {
    Expr any
    Attributes any
}

// StmtFinally : Stmt
type FinallyStmt struct {
    Stmts any
    Attributes any
}

// StmtFor : Stmt
type ForStmt struct {
    Init any
    Cond any
    Loop any
    Stmts any
    Attributes any
}

// StmtForeach : Stmt
type ForeachStmt struct {
    Expr any
    KeyVar any
    ByRef any
    ValueVar any
    Stmts any
    Attributes any
}

// StmtFunction : Stmt, FunctionLike
type FunctionStmt struct {
    ByRef any
    Name any
    Params any
    ReturnType any
    Stmts any
    AttrGroups any
    NamespacedName any
    Attributes any
}

// StmtGlobal : Stmt
type GlobalStmt struct {
    Vars any
    Attributes any
}

// StmtGoto : Stmt
type GotoStmt struct {
    Name any
    Attributes any
}

// StmtGroupUse : Stmt
type GroupUseStmt struct {
    Type any
    Prefix any
    Uses any
    Attributes any
}

// StmtHaltCompiler : Stmt
type HaltCompilerStmt struct {
    Remaining any
    Attributes any
}

// StmtIf : Stmt
type IfStmt struct {
    Cond any
    Stmts any
    Elseifs any
    Else any
    Attributes any
}

// StmtInlineHTML : Stmt
type InlineHTMLStmt struct {
    Value any
    Attributes any
}

// StmtInterface : StmtClassLike
type InterfaceStmt struct {
    Extends any
    Name any
    Stmts any
    AttrGroups any
    NamespacedName any
    Attributes any
}

// StmtLabel : Stmt
type LabelStmt struct {
    Name any
    Attributes any
}

// StmtNamespace : Stmt
type NamespaceStmt struct {
    Name any
    Stmts any
    Attributes any
}

// StmtNop : Stmt
type NopStmt struct {
    Attributes any
}

// StmtProperty : Stmt
type PropertyStmt struct {
    Flags any
    Props any
    Type any
    AttrGroups any
    Attributes any
}

// StmtPropertyProperty : Stmt
type PropertyPropertyStmt struct {
    Name any
    Default any
    Attributes any
}

// StmtReturn : Stmt
type ReturnStmt struct {
    Expr any
    Attributes any
}

// StmtStatic : Stmt
type StaticStmt struct {
    Vars any
    Attributes any
}

// StmtStaticVar : Stmt
type StaticVarStmt struct {
    Var any
    Default any
    Attributes any
}

// StmtSwitch : Stmt
type SwitchStmt struct {
    Cond any
    Cases any
    Attributes any
}

// StmtThrow : Stmt
type ThrowStmt struct {
    Expr any
    Attributes any
}

// StmtTrait : StmtClassLike
type TraitStmt struct {
    Name any
    Stmts any
    AttrGroups any
    NamespacedName any
    Attributes any
}

// StmtTraitUse : Stmt
type TraitUseStmt struct {
    Traits any
    Adaptations any
    Attributes any
}

// StmtTraitUseAdaptationAlias : StmtTraitUseAdaptation
type TraitUseAdaptationAliasStmt struct {
    NewModifier any
    NewName any
    Trait any
    Method any
    Attributes any
}

// StmtTraitUseAdaptationPrecedence : StmtTraitUseAdaptation
type TraitUseAdaptationPrecedenceStmt struct {
    Insteadof any
    Trait any
    Method any
    Attributes any
}

// StmtTryCatch : Stmt
type TryCatchStmt struct {
    Stmts any
    Catches any
    Finally any
    Attributes any
}

// StmtUnset : Stmt
type UnsetStmt struct {
    Vars any
    Attributes any
}

// StmtUse : Stmt
type UseStmt struct {
    Type any
    Uses any
    Attributes any
}

// StmtUseUse : Stmt
type UseUseStmt struct {
    Type any
    Name any
    Alias any
    Attributes any
}

// StmtWhile : Stmt
type WhileStmt struct {
    Cond any
    Stmts any
    Attributes any
}

// UnionType : ComplexType
type UnionType struct {
    Types any
    Attributes any
}

// VarLikeIdentifier : Identifier, Stringable
type VarLikeIdentifier struct {
    Name any
    Attributes any
}

// VariadicPlaceholder : PhpParserNodeAbstract
type VariadicPlaceholder struct {
    Attributes any
}

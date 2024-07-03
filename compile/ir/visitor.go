package ir

import "errors"

type Visitor interface {
	VisitArg(n *Arg)
	VisitName(n *Name)
	VisitIdent(n *Ident)
	VisitParam(n *Param)
	// Type
	VisitSimpleType(n *SimpleType)
	VisitIntersectionType(n *IntersectionType)
	VisitUnionType(n *UnionType)
	VisitNullableType(n *NullableType)
	// Expr
	VisitNullLit(n *NullLit)
	VisitBoolLit(n *BoolLit)
	VisitIntLit(n *IntLit)
	VisitFloatLit(n *FloatLit)
	VisitStringLit(n *StringLit)
	VisitArrayExpr(n *ArrayExpr)
	VisitArrayItemExpr(n *ArrayItemExpr)
	VisitClosureExpr(n *ClosureExpr)
	VisitClosureUseExpr(n *ClosureUseExpr)
	VisitArrowFunctionExpr(n *ArrowFunctionExpr)
	VisitIndexExpr(n *IndexExpr)
	VisitCastExpr(n *CastExpr)
	VisitUnaryExpr(n *UnaryExpr)
	VisitBinaryOpExpr(n *BinaryOpExpr)
	VisitAssignExpr(n *AssignExpr)
	VisitAssignOpExpr(n *AssignOpExpr)
	VisitAssignRefExpr(n *AssignRefExpr)
	VisitIssetExpr(n *IssetExpr)
	VisitEmptyExpr(n *EmptyExpr)
	VisitEvalExpr(n *EvalExpr)
	VisitIncludeExpr(n *IncludeExpr)
	VisitCloneExpr(n *CloneExpr)
	VisitErrorSuppressExpr(n *ErrorSuppressExpr)
	VisitExitExpr(n *ExitExpr)
	VisitConstFetchExpr(n *ConstFetchExpr)
	VisitClassConstFetchExpr(n *ClassConstFetchExpr)
	VisitMagicConstExpr(n *MagicConstExpr)
	VisitInstanceofExpr(n *InstanceofExpr)
	VisitListExpr(n *ListExpr)
	VisitPrintExpr(n *PrintExpr)
	VisitPropertyFetchExpr(n *PropertyFetchExpr)
	VisitStaticPropertyFetchExpr(n *StaticPropertyFetchExpr)
	VisitShellExecExpr(n *ShellExecExpr)
	VisitTernaryExpr(n *TernaryExpr)
	VisitThrowExpr(n *ThrowExpr)
	VisitVariableExpr(n *VariableExpr)
	VisitYieldExpr(n *YieldExpr)
	VisitYieldFromExpr(n *YieldFromExpr)
	VisitFuncCallExpr(n *FuncCallExpr)
	VisitNewExpr(n *NewExpr)
	VisitMethodCallExpr(n *MethodCallExpr)
	VisitStaticCallExpr(n *StaticCallExpr)
	// Stmt
	VisitEmptyStmt(n *EmptyStmt)
	VisitExprStmt(n *ExprStmt)
	VisitReturnStmt(n *ReturnStmt)
	VisitLabelStmt(n *LabelStmt)
	VisitGotoStmt(n *GotoStmt)
	VisitIfStmt(n *IfStmt)
	VisitElseIfStmt(n *ElseIfStmt)
	VisitElseStmt(n *ElseStmt)
	VisitSwitchStmt(n *SwitchStmt)
	VisitCaseStmt(n *CaseStmt)
	VisitForStmt(n *ForStmt)
	VisitForeachStmt(n *ForeachStmt)
	VisitBreakStmt(n *BreakStmt)
	VisitContinueStmt(n *ContinueStmt)
	VisitWhileStmt(n *WhileStmt)
	VisitDoStmt(n *DoStmt)
	VisitTryCatchStmt(n *TryCatchStmt)
	VisitCatchStmt(n *CatchStmt)
	VisitFinallyStmt(n *FinallyStmt)
	VisitConstStmt(n *ConstStmt)
	VisitEchoStmt(n *EchoStmt)
	VisitGlobalStmt(n *GlobalStmt)
	VisitStaticStmt(n *StaticStmt)
	VisitUnsetStmt(n *UnsetStmt)
	VisitUseStmt(n *UseStmt)
	VisitFunctionStmt(n *FunctionStmt)
	VisitInterfaceStmt(n *InterfaceStmt)
	VisitClassStmt(n *ClassStmt)
	VisitClassConstStmt(n *ClassConstStmt)
	VisitPropertyStmt(n *PropertyStmt)
	VisitClassMethodStmt(n *ClassMethodStmt)
	VisitTraitStmt(n *TraitStmt)
	VisitTraitUseStmt(n *TraitUseStmt)
	VisitTraitUseAdaptationAliasStmt(n *TraitUseAdaptationAliasStmt)
	VisitTraitUseAdaptationPrecedenceStmt(n *TraitUseAdaptationPrecedenceStmt)
}

var unexpectedNodeError = errors.New("unexpected node type")

func IsUnexpectedError(err error) bool {
	return errors.Is(err, unexpectedNodeError)
}

func Visit(v Visitor, n Node) error {
	if v == nil {
		return nil
	}

	switch x := n.(type) {
	case *Ident:
		v.VisitIdent(x)
	case *Name:
		v.VisitName(x)
	case *Param:
		v.VisitParam(x)
	case *Arg:
		v.VisitArg(x)
	case Expr:
		return VisitExpr(v, x)
	case Stmt:
		return VisitStmt(v, x)
	case TypeHint:
		return VisitTypeHint(v, x)
	default:
		return unexpectedNodeError
	}

	return nil
}

func VisitTypeHint(v Visitor, n TypeHint) error {
	if v == nil {
		return nil
	}

	switch x := n.(type) {
	case *SimpleType:
		v.VisitSimpleType(x)
	case *IntersectionType:
		v.VisitIntersectionType(x)
	case *UnionType:
		v.VisitUnionType(x)
	case *NullableType:
		v.VisitNullableType(x)
	default:
		return unexpectedNodeError
	}

	return nil
}

func VisitExpr(v Visitor, n Expr) error {
	if v == nil {
		return nil
	}

	switch x := n.(type) {
	case *NullLit:
		v.VisitNullLit(x)
	case *BoolLit:
		v.VisitBoolLit(x)
	case *IntLit:
		v.VisitIntLit(x)
	case *FloatLit:
		v.VisitFloatLit(x)
	case *StringLit:
		v.VisitStringLit(x)
	case *ArrayExpr:
		v.VisitArrayExpr(x)
	case *ArrayItemExpr:
		v.VisitArrayItemExpr(x)
	case *ClosureExpr:
		v.VisitClosureExpr(x)
	case *ClosureUseExpr:
		v.VisitClosureUseExpr(x)
	case *ArrowFunctionExpr:
		v.VisitArrowFunctionExpr(x)
	case *IndexExpr:
		v.VisitIndexExpr(x)
	case *CastExpr:
		v.VisitCastExpr(x)
	case *UnaryExpr:
		v.VisitUnaryExpr(x)
	case *BinaryOpExpr:
		v.VisitBinaryOpExpr(x)
	case *AssignExpr:
		v.VisitAssignExpr(x)
	case *AssignOpExpr:
		v.VisitAssignOpExpr(x)
	case *AssignRefExpr:
		v.VisitAssignRefExpr(x)
	case *IssetExpr:
		v.VisitIssetExpr(x)
	case *EmptyExpr:
		v.VisitEmptyExpr(x)
	case *EvalExpr:
		v.VisitEvalExpr(x)
	case *IncludeExpr:
		v.VisitIncludeExpr(x)
	case *CloneExpr:
		v.VisitCloneExpr(x)
	case *ErrorSuppressExpr:
		v.VisitErrorSuppressExpr(x)
	case *ExitExpr:
		v.VisitExitExpr(x)
	case *ConstFetchExpr:
		v.VisitConstFetchExpr(x)
	case *ClassConstFetchExpr:
		v.VisitClassConstFetchExpr(x)
	case *MagicConstExpr:
		v.VisitMagicConstExpr(x)
	case *InstanceofExpr:
		v.VisitInstanceofExpr(x)
	case *ListExpr:
		v.VisitListExpr(x)
	case *PrintExpr:
		v.VisitPrintExpr(x)
	case *PropertyFetchExpr:
		v.VisitPropertyFetchExpr(x)
	case *StaticPropertyFetchExpr:
		v.VisitStaticPropertyFetchExpr(x)
	case *ShellExecExpr:
		v.VisitShellExecExpr(x)
	case *TernaryExpr:
		v.VisitTernaryExpr(x)
	case *ThrowExpr:
		v.VisitThrowExpr(x)
	case *VariableExpr:
		v.VisitVariableExpr(x)
	case *YieldExpr:
		v.VisitYieldExpr(x)
	case *YieldFromExpr:
		v.VisitYieldFromExpr(x)
	case *FuncCallExpr:
		v.VisitFuncCallExpr(x)
	case *NewExpr:
		v.VisitNewExpr(x)
	case *MethodCallExpr:
		v.VisitMethodCallExpr(x)
	case *StaticCallExpr:
		v.VisitStaticCallExpr(x)
	default:
		return unexpectedNodeError
	}

	return nil
}

func VisitStmt(v Visitor, n Stmt) error {
	if v == nil {
		return nil
	}

	switch x := n.(type) {
	case *EmptyStmt:
		v.VisitEmptyStmt(x)
	case *ExprStmt:
		v.VisitExprStmt(x)
	case *ReturnStmt:
		v.VisitReturnStmt(x)
	case *LabelStmt:
		v.VisitLabelStmt(x)
	case *GotoStmt:
		v.VisitGotoStmt(x)
	case *IfStmt:
		v.VisitIfStmt(x)
	case *ElseIfStmt:
		v.VisitElseIfStmt(x)
	case *ElseStmt:
		v.VisitElseStmt(x)
	case *SwitchStmt:
		v.VisitSwitchStmt(x)
	case *CaseStmt:
		v.VisitCaseStmt(x)
	case *ForStmt:
		v.VisitForStmt(x)
	case *ForeachStmt:
		v.VisitForeachStmt(x)
	case *BreakStmt:
		v.VisitBreakStmt(x)
	case *ContinueStmt:
		v.VisitContinueStmt(x)
	case *WhileStmt:
		v.VisitWhileStmt(x)
	case *DoStmt:
		v.VisitDoStmt(x)
	case *TryCatchStmt:
		v.VisitTryCatchStmt(x)
	case *CatchStmt:
		v.VisitCatchStmt(x)
	case *FinallyStmt:
		v.VisitFinallyStmt(x)
	case *ConstStmt:
		v.VisitConstStmt(x)
	case *EchoStmt:
		v.VisitEchoStmt(x)
	case *GlobalStmt:
		v.VisitGlobalStmt(x)
	case *StaticStmt:
		v.VisitStaticStmt(x)
	case *UnsetStmt:
		v.VisitUnsetStmt(x)
	case *UseStmt:
		v.VisitUseStmt(x)
	case *FunctionStmt:
		v.VisitFunctionStmt(x)
	case *InterfaceStmt:
		v.VisitInterfaceStmt(x)
	case *ClassStmt:
		v.VisitClassStmt(x)
	case *ClassConstStmt:
		v.VisitClassConstStmt(x)
	case *PropertyStmt:
		v.VisitPropertyStmt(x)
	case *ClassMethodStmt:
		v.VisitClassMethodStmt(x)
	case *TraitStmt:
		v.VisitTraitStmt(x)
	case *TraitUseStmt:
		v.VisitTraitUseStmt(x)
	case *TraitUseAdaptationAliasStmt:
		v.VisitTraitUseAdaptationAliasStmt(x)
	case *TraitUseAdaptationPrecedenceStmt:
		v.VisitTraitUseAdaptationPrecedenceStmt(x)
	default:
		return unexpectedNodeError
	}

	return nil
}

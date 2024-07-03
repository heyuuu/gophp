package render

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/php/def"
	"strconv"
)

func (r *render) VisitArg(n *ir.Arg) {
	r.pDefCall("Arg", n.Value, n.Unpack)
}

func (r *render) VisitName(n *ir.Name) {
	r.print(stringLit(n.ToCodeString()))
}

func (r *render) VisitIdent(n *ir.Ident) {
	r.print(stringLit(n.Name))
}

func (r *render) VisitParam(n *ir.Param) {
	r.pDefCall("Param", n.Type, n.Var, n.Default, n.ByRef, n.Variadic)
}

// type

func (r *render) VisitSimpleType(n *ir.SimpleType) {
	r.pDefCall("SimpleType", n.Name.ToCodeString())
}

func (r *render) VisitIntersectionType(n *ir.IntersectionType) {
	r.pDefCall("IntersectionType", n.Types)
}

func (r *render) VisitUnionType(n *ir.UnionType) {
	r.pDefCall("UnionType", n.Types)
}

func (r *render) VisitNullableType(n *ir.NullableType) {
	r.pDefCall("NullableType", n.Type)
}

// Expr
func (r *render) VisitNullLit(n *ir.NullLit) {
	r.pDefCall("Null")
}

func (r *render) VisitBoolLit(n *ir.BoolLit) {
	r.pDefCall("Bool", n.Value)
}

func (r *render) VisitIntLit(n *ir.IntLit) {
	r.pDefCall("Int", strconv.Itoa(n.Value))
}

func (r *render) VisitFloatLit(n *ir.FloatLit) {
	r.pDefCall("Float", fmt.Sprintf("%f", n.Value))
}

func (r *render) VisitStringLit(n *ir.StringLit) {
	r.pDefCall("String", stringLit(n.Value))
}

func (r *render) VisitArrayExpr(n *ir.ArrayExpr) {
	r.pDefCall("Array", func() {
		pList(r, n.Items, ", ")
	})

}

func (r *render) VisitArrayItemExpr(n *ir.ArrayItemExpr) {
	r.pDefCall("ArrayItem", n.Key, n.Value, n.ByRef, n.Unpack)
}

func (r *render) VisitClosureExpr(n *ir.ClosureExpr) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitClosureUseExpr(n *ir.ClosureUseExpr) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitArrowFunctionExpr(n *ir.ArrowFunctionExpr) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitIndexExpr(n *ir.IndexExpr) {
	r.pDefCall("Index", n.Var, n.Dim)
}

func (r *render) VisitCastExpr(n *ir.CastExpr) {
	var castType string
	switch n.Kind {
	case ast.CastArray:
		castType = "ast.CastArray"
	case ast.CastBool:
		castType = "ast.CastBool"
	case ast.CastDouble:
		castType = "ast.CastDouble"
	case ast.CastInt:
		castType = "ast.CastInt"
	case ast.CastObject:
		castType = "ast.CastObject"
	case ast.CastString:
		castType = "ast.CastString"
	case ast.CastUnset:
		castType = "ast.CastUnset"
	}

	r.usePkg(pkgAst)
	r.pDefCall("Cast", castType, n.Expr)
}

func (r *render) VisitUnaryExpr(n *ir.UnaryExpr) {
	var method string

	switch n.Op {
	case ast.UnaryOpPlus:
		method = "PrePlus"
	case ast.UnaryOpMinus:
		method = "PreMinus"
	case ast.UnaryOpBooleanNot:
		method = "BooleanNot"
	case ast.UnaryOpBitwiseNot:
		method = "BitwiseNot"
	case ast.UnaryOpPreInc:
		method = "PrePlus"
	case ast.UnaryOpPreDec:
		method = "PreDec"
	case ast.UnaryOpPostInc:
		method = "PostInc"
	case ast.UnaryOpPostDec:
		method = "PostDec"
	}

	r.pDefCall(method, n.Var)
}

func (r *render) VisitBinaryOpExpr(n *ir.BinaryOpExpr) {
	var method string
	switch n.Op {
	// todo
	}

	r.pDefCall(method, n.Left, n.Right)
}

func (r *render) VisitAssignExpr(n *ir.AssignExpr) {
	r.pDefCall("Assign", n.Var, n.Expr)
}

func (r *render) VisitAssignOpExpr(n *ir.AssignOpExpr) {
	var method string
	switch n.Op {
	// todo
	}

	r.pDefCall(method, n.Var, n.Expr)
}

func (r *render) VisitAssignRefExpr(n *ir.AssignRefExpr) {
	r.pDefCall("AssignRef", n.Var, n.Expr)
}

func (r *render) VisitIssetExpr(n *ir.IssetExpr) {
	r.pDefCall("Isset", n.Vars)
}

func (r *render) VisitEmptyExpr(n *ir.EmptyExpr) {
	r.pDefCall("Empty", n.Expr)
}

func (r *render) VisitEvalExpr(n *ir.EvalExpr) {
	r.pDefCall("Eval", n.Expr)
}

func (r *render) VisitIncludeExpr(n *ir.IncludeExpr) {
	var kind string
	switch n.Kind {
	case ast.KindInclude:
		kind = "KindInclude"
	case ast.KindIncludeOnce:
		kind = "KindIncludeOnce"
	case ast.KindRequire:
		kind = "KindRequire"
	case ast.KindRequireOnce:
		kind = "KindRequireOnce"
	}

	r.pDefCall("Include", func() { r.print(pkgAst, kind) }, n.Expr)
}

func (r *render) VisitCloneExpr(n *ir.CloneExpr) {
	r.pDefCall("Clone", n.Expr)
}

func (r *render) VisitErrorSuppressExpr(n *ir.ErrorSuppressExpr) {
	r.pDefCall("ErrorSuppress", n.Expr)
}

func (r *render) VisitExitExpr(n *ir.ExitExpr) {
	r.pDefCall("ErrorSuppress", n.Expr)
}

func (r *render) VisitConstFetchExpr(n *ir.ConstFetchExpr) {
	r.pDefCall("Const", n.Name.ToString())
}

func (r *render) VisitClassConstFetchExpr(n *ir.ClassConstFetchExpr) {
	r.pDefCall("ClassConst", func() {
		r.pClassName(n)
	}, n.Name)
}

func (r *render) VisitMagicConstExpr(n *ir.MagicConstExpr) {
	var kind string
	var value ir.Expr
	switch n.Kind {
	case ast.MagicConstClass:
		kind = "MagicConstClass"
	case ast.MagicConstDir:
		kind = "MagicConstDir"
	case ast.MagicConstFile:
		kind = "MagicConstFile"
	case ast.MagicConstFunction:
		kind = "MagicConstFunction"
	case ast.MagicConstLine:
		kind = "MagicConstLine"
	case ast.MagicConstMethod:
		kind = "MagicConstMethod"
	case ast.MagicConstNamespace:
		kind = "MagicConstNamespace"
	case ast.MagicConstTrait:
		kind = "MagicConstTrait"
	}

	r.pDefCall("MagicConst", func() {
		r.print(pkgAst, kind)
	}, value)
}

func (r *render) VisitInstanceofExpr(n *ir.InstanceofExpr) {
	r.pDefCall("Instanceof", n.Expr, func() { r.pClassName(n.Class) })
}

func (r *render) VisitListExpr(n *ir.ListExpr) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitPrintExpr(n *ir.PrintExpr) {
	r.pDefCall("Print", n.Expr)
}

func (r *render) VisitPropertyFetchExpr(n *ir.PropertyFetchExpr) {
	if n.Nullsafe {
		r.pDefCall("NullsafeProp", n.Var, n.Name)
	} else {
		r.pDefCall("Prop", n.Var, n.Name)
	}
}

func (r *render) VisitStaticPropertyFetchExpr(n *ir.StaticPropertyFetchExpr) {
	r.pDefCall("StaticProperty", func() { r.pClassName(n.Class) }, n.Name)
}

func (r *render) VisitShellExecExpr(n *ir.ShellExecExpr) {
	r.todo("VisitShellExecExpr")
}

func (r *render) VisitTernaryExpr(n *ir.TernaryExpr) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitThrowExpr(n *ir.ThrowExpr) {
	r.pDefCall("Throw", n.Expr)
}

func (r *render) VisitVariableExpr(n *ir.VariableExpr) {
	r.pDefCall("Var", func() {
		r.pIdentOrExprAsStr(n.Name, def.StageVariableName, "VariableExpr.Name")
	})
}

func (r *render) VisitYieldExpr(n *ir.YieldExpr) {
	r.todo("YieldExpr")
}

func (r *render) VisitYieldFromExpr(n *ir.YieldFromExpr) {
	r.todo("YieldFromExpr")
}

func (r *render) VisitFuncCallExpr(n *ir.FuncCallExpr) {
	r.pDefCall("FuncCall", func() {
		r.pNameOrExprAsStr(n.Name, def.StageFuncName, "FuncCallExpr.Name")

		// args
		if len(n.Args) > 0 {
			r.print(", ")
			pList(r, n.Args, ", ")
		}
	})
}

func (r *render) VisitNewExpr(n *ir.NewExpr) {
	r.pDefCall("New", func() {
		if _, ok := n.Class.(*ir.ClassStmt); ok {
			r.todo("NewExpr.Class is ClassStmt")
		}

		r.pClassName(n.Class)
		// args
		if len(n.Args) > 0 {
			r.print(", ")
			pList(r, n.Args, ", ")
		}
	})
}

func (r *render) VisitMethodCallExpr(n *ir.MethodCallExpr) {
	method := "MethodCall"
	if n.Nullsafe {
		method = "NullsafeMethodCall"
	}

	r.pDefCall(method, n.Var, func() {
		r.pIdentOrExprAsStr(n.Name, def.StageMethodName, "MethodCallExpr.Name")
		// args
		if len(n.Args) > 0 {
			r.print(", ")
			pList(r, n.Args, ", ")
		}
	})
}

func (r *render) VisitStaticCallExpr(n *ir.StaticCallExpr) {
	r.pDefCall("StaticMethodCall", func() {
		// class
		r.pClassName(n.Class)
		// method
		r.print(", ")
		r.pIdentOrExprAsStr(n.Name, def.StageMethodName, "StaticCallExpr.Name")
		// args
		if len(n.Args) > 0 {
			r.print(", ")
			pList(r, n.Args, ", ")
		}
	})
}

func (r *render) VisitEmptyStmt(n *ir.EmptyStmt) {
	// pass
}

func (r *render) VisitExprStmt(n *ir.ExprStmt) {
	r.expr(n.Expr)
}

func (r *render) VisitReturnStmt(n *ir.ReturnStmt) {
	r.print("return ", n.Expr)
}

func (r *render) VisitLabelStmt(n *ir.LabelStmt) {
	r.print(n.Name.Name, ":")
}

func (r *render) VisitGotoStmt(n *ir.GotoStmt) {
	r.print("goto ", n.Name.Name)
}

func (r *render) VisitIfStmt(n *ir.IfStmt) {
	r.print("if ", n.Cond, " {\n", n.Stmts, "}")

	for _, elseif := range n.Elseifs {
		r.VisitElseIfStmt(elseif)
	}

	if n.Else != nil {
		r.VisitElseStmt(n.Else)
	}
}

func (r *render) VisitElseIfStmt(n *ir.ElseIfStmt) {
	r.print(" else if ", n.Cond, " {\n", n.Stmts, "}")
}

func (r *render) VisitElseStmt(n *ir.ElseStmt) {
	r.print(" else {\n", n.Stmts, "}")
}

func (r *render) VisitSwitchStmt(n *ir.SwitchStmt) {
	r.print("switch ", varSwitch, " := ", n.Cond, "; {\n")
	pList(r, n.Cases, "\n")
	r.print("\n}")
}

func (r *render) VisitCaseStmt(n *ir.CaseStmt) {
	if n.Cond != nil {
		r.print("case ")
		r.pAsBool(func() {
			r.pDefCall("Equal", varSwitch, n.Cond)
		})
		r.print(":\n")
	} else {
		r.print("default:\n")
	}
	r.stmtList(n.Stmts, true)
}

func (r *render) VisitForStmt(n *ir.ForStmt) {
	if len(n.Init) > 1 || len(n.Cond) > 1 || len(n.Loop) > 1 {
		r.todo(`ForStmt with multiple init, cond or loop`)
	}

	r.print("for ", n.Init, "; ", n.Cond, "; ", n.Loop, " {\n", n.Stmts, "}")
}

func (r *render) VisitForeachStmt(n *ir.ForeachStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitBreakStmt(n *ir.BreakStmt) {
	if n.Num != nil {
		r.todo("BreakStmt.Num is not nil")
	}
	r.print("break")
}

func (r *render) VisitContinueStmt(n *ir.ContinueStmt) {
	if n.Num != nil {
		r.todo("ContinueStmt.Num is not nil")
	}
	r.print("continue")
}

func (r *render) VisitWhileStmt(n *ir.WhileStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitDoStmt(n *ir.DoStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitTryCatchStmt(n *ir.TryCatchStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitCatchStmt(n *ir.CatchStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitFinallyStmt(n *ir.FinallyStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitConstStmt(n *ir.ConstStmt) {
	r.pDefCall("DeclConst", n.Name, n.Value)
}

func (r *render) VisitEchoStmt(n *ir.EchoStmt) {
	r.pDefCall("EchoVal", n.Exprs)
}

func (r *render) VisitGlobalStmt(n *ir.GlobalStmt) {
	r.pDefCall("DeclGlobal", n.Var) // todo global
}

func (r *render) VisitStaticStmt(n *ir.StaticStmt) {
	r.pDefCall("DeclGlobal", n.Var, n.Default)
}

func (r *render) VisitUnsetStmt(n *ir.UnsetStmt) {
	r.pDefCall("Unset", n.Var)
}

func (r *render) VisitUseStmt(n *ir.UseStmt) {
	var useType string
	switch n.Type {
	case ast.UseNormal:
		useType = "UseNormal"
	case ast.UseFunction:
		useType = "UseFunction"
	case ast.UseConstant:
		useType = "UseConstant"
	}

	var alias string
	if n.Alias != nil {
		alias = n.Alias.Name
	}

	r.pDefCall("Use",
		func() { r.print(pkgAst, ".", useType) },
		n.Name,
		stringLit(alias),
	)
}

func (r *render) VisitFunctionStmt(n *ir.FunctionStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitInterfaceStmt(n *ir.InterfaceStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitClassStmt(n *ir.ClassStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitClassConstStmt(n *ir.ClassConstStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitPropertyStmt(n *ir.PropertyStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitClassMethodStmt(n *ir.ClassMethodStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitTraitStmt(n *ir.TraitStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitTraitUseStmt(n *ir.TraitUseStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitTraitUseAdaptationAliasStmt(n *ir.TraitUseAdaptationAliasStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *render) VisitTraitUseAdaptationPrecedenceStmt(n *ir.TraitUseAdaptationPrecedenceStmt) {
	//TODO implement me
	panic("implement me")
}

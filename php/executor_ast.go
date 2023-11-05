package php

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/php/types"
)

type executeState uint8

const (
	stateNormal executeState = iota
	stateReturn
	stateBreak
	stateContinue
	stateGoto
)

type (
	executeResult interface {
		state() executeState
	}

	returnResult   struct{ retVal Val }
	breakResult    struct{ num int }
	continueResult struct{ num int }
	gotoResult     struct{ label string }
)

func (r returnResult) state() executeState   { return stateReturn }
func (r breakResult) state() executeState    { return stateBreak }
func (r continueResult) state() executeState { return stateContinue }
func (r gotoResult) state() executeState     { return stateGoto }

// private

func (e *executor) executeAstFile(f *ast.File) (Val, error) {
	for _, ns := range f.Namespaces {
		res := e.stmtList(ns.Stmts)
		switch r := res.(type) {
		case *returnResult:
			return r.retVal, nil
		case *continueResult, *breakResult, *gotoResult:
			panic("unreachable")
		}
	}
	return nil, nil
}

func (e *executor) userFunction(fn *types.Function, args []Val) Val {
	// todo 初始化 args
	res := e.stmtList(fn.Stmts())
	switch r := res.(type) {
	case *returnResult:
		return r.retVal
	case *continueResult, *breakResult, *gotoResult:
		panic("unreachable")
	}
	return nil
}

func (e *executor) stmtList(stmts []ast.Stmt) (result executeResult) {
	var labels = map[string]int{}
	for i, stmt := range stmts {
		if label, ok := stmt.(*ast.LabelStmt); ok {
			labels[label.Name.Name] = i
		}
	}

	l := len(stmts)
	for i := 0; i < l; i++ {
		switch x := stmts[i].(type) {
		case *ast.EmptyStmt: // pass
		case *ast.ExprStmt:
			_ = e.expr(x.Expr)
		case *ast.ReturnStmt:
			retVal := e.expr(x.Expr)
			return returnResult{retVal: retVal}
		case *ast.LabelStmt:
			// pass
			// todo goto 能跳到非循环结构内(比如 if)
		case *ast.GotoStmt:
			// todo goto 处理逻辑
			labelName := x.Name.Name
			if v, ok := labels[labelName]; ok {
				i = v
			} else {
				return gotoResult{labelName}
			}
		case *ast.EchoStmt:
			values := e.exprList(x.Exprs)
			for _, value := range values {
				vmEcho(e.ctx, value)
			}
		// todo
		default:
			panic(fmt.Sprintf("todo executor.stmtList(%T)", x))
		}
	}
	return
}

func (e *executor) exprList(exprs []ast.Expr) []Val {
	values := make([]Val, len(exprs))
	for i, expr := range exprs {
		values[i] = e.expr(expr)
	}
	return values
}

func (e *executor) expr(expr ast.Expr) Val {
	switch x := expr.(type) {
	case *ast.IntLit:
		return Long(x.Value)
	case *ast.FloatLit:
		return Double(x.Value)
	case *ast.StringLit:
		return String(x.Value)
	case *ast.ArrayExpr:
		return e.executeArrayExpr(x)
	case *ast.ClosureExpr:
		return e.executeClosureExpr(x)
	case *ast.ClosureUseExpr:
		return e.executeClosureUseExpr(x)
	case *ast.ArrowFunctionExpr:
		return e.executeArrowFunctionExpr(x)
	case *ast.IndexExpr:
		return e.executeIndexExpr(x)
	case *ast.CastExpr:
		return e.executeCastExpr(x)
	case *ast.UnaryExpr:
		return e.executeUnaryExpr(x)
	case *ast.BinaryOpExpr:
		return e.executeBinaryOpExpr(x)
	case *ast.AssignExpr:
		return e.executeAssignExpr(x)
	case *ast.AssignOpExpr:
		return e.executeAssignOpExpr(x)
	case *ast.AssignRefExpr:
		return e.executeAssignRefExpr(x)
	case *ast.IssetExpr:
		return e.executeIssetExpr(x)
	case *ast.EmptyExpr:
		return e.executeEmptyExpr(x)
	case *ast.EvalExpr:
		return e.executeEvalExpr(x)
	case *ast.IncludeExpr:
		return e.executeIncludeExpr(x)
	case *ast.CloneExpr:
		return e.executeCloneExpr(x)
	case *ast.ErrorSuppressExpr:
		return e.executeErrorSuppressExpr(x)
	case *ast.ExitExpr:
		return e.executeExitExpr(x)
	case *ast.ConstFetchExpr:
		return e.executeConstFetchExpr(x)
	case *ast.ClassConstFetchExpr:
		return e.executeClassConstFetchExpr(x)
	case *ast.MagicConstExpr:
		return e.executeMagicConstExpr(x)
	case *ast.InstanceofExpr:
		return e.executeInstanceofExpr(x)
	case *ast.ListExpr:
		return e.executeListExpr(x)
	case *ast.PrintExpr:
		return e.executePrintExpr(x)
	case *ast.PropertyFetchExpr:
		return e.executePropertyFetchExpr(x)
	case *ast.StaticPropertyFetchExpr:
		return e.executeStaticPropertyFetchExpr(x)
	case *ast.ShellExecExpr:
		return e.executeShellExecExpr(x)
	case *ast.TernaryExpr:
		return e.executeTernaryExpr(x)
	case *ast.ThrowExpr:
		return e.executeThrowExpr(x)
	case *ast.VariableExpr:
		return e.executeVariableExpr(x)
	case *ast.YieldExpr:
		return e.executeYieldExpr(x)
	case *ast.YieldFromExpr:
		return e.executeYieldFromExpr(x)
	case *ast.FuncCallExpr:
		return e.executeFuncCallExpr(x)
	case *ast.NewExpr:
		return e.executeNewExpr(x)
	case *ast.MethodCallExpr:
		return e.executeMethodCallExpr(x)
	case *ast.StaticCallExpr:
		return e.executeStaticCallExpr(x)
	case *ast.ArrayItemExpr:
		panic(fmt.Sprintf("unexpected execute type: %T", x))
	default:
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	}
}

func (e *executor) executeBinaryOpExpr(expr *ast.BinaryOpExpr) (val Val) {
	op := e.operator

	// && / || / ?? 操作比较特殊，右表达式节点可能不会执行
	switch expr.Op {
	case ast.BinaryOpBooleanAnd: // &&
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return op.BooleanAnd(left, right)
	case ast.BinaryOpBooleanOr: // ||
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return op.BooleanAnd(left, right)
	case ast.BinaryOpCoalesce: // ??
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return op.Coalesce(left, right)
	}

	// common
	left := e.expr(expr.Left)
	right := e.expr(expr.Right)

	switch expr.Op {
	case ast.BinaryOpPlus: // +
		return op.Add(left, right)
	case ast.BinaryOpMinus: // -
		return op.Sub(left, right)
	case ast.BinaryOpMul: // *
		return op.Mul(left, right)
	case ast.BinaryOpDiv: // /
		return op.Div(left, right)
	case ast.BinaryOpMod: // %
		return op.Mod(left, right)
	case ast.BinaryOpPow: // **
		return op.Pow(left, right)
	case ast.BinaryOpBitwiseAnd: // &
		return op.BitwiseAnd(left, right)
	case ast.BinaryOpBitwiseOr: // n|
		return op.BitwiseOr(left, right)
	case ast.BinaryOpBitwiseXor: // ^
		return op.BitwiseXor(left, right)
	case ast.BinaryOpConcat: // .
		return op.Concat(left, right)
	case ast.BinaryOpEqual: // ==
		return op.Equal(left, right)
	case ast.BinaryOpGreater: // >
		return op.Greater(left, right)
	case ast.BinaryOpGreaterOrEqual: // >=
		return op.GreaterOrEqual(left, right)
	case ast.BinaryOpIdentical: // ===
		return op.Identical(left, right)
	case ast.BinaryOpBooleanXor: // xor
		return op.BooleanXor(left, right)
	case ast.BinaryOpNotEqual: // !=
		return op.NotEqual(left, right)
	case ast.BinaryOpNotIdentical: // !==
		return op.NotIdentical(left, right)
	case ast.BinaryOpShiftLeft: // <<
		return op.SL(left, right)
	case ast.BinaryOpShiftRight: // >>
		return op.SR(left, right)
	case ast.BinaryOpSmaller: // <
		return op.Smaller(left, right)
	case ast.BinaryOpSmallerOrEqual: // <=
		return op.SmallerOrEqual(left, right)
	case ast.BinaryOpSpaceship: // <=>
		return op.Spaceship(left, right)
	default:
		panic("unreachable")
	}
}
func (e *executor) executeArrayExpr(expr *ast.ArrayExpr) Val {
	arr := types.NewArrayCap(len(expr.Items))
	for _, item := range expr.Items {
		if item.ByRef {
			// todo item byref
			panic("todo item byref")
		} else if item.Unpack && item.Key != nil {
			// todo item unpack with key
			panic("todo item unpack with key")
		}

		if item.Key != nil {
			key := e.expr(item.Key)
			value := e.expr(item.Value)
			arrayKey := e.operator.ZvalToArrayKey(key)
			arr.Add(arrayKey, value)
		} else {
			value := e.expr(item.Value)
			arr.Append(value)
		}
	}
	return nil
}
func (e *executor) executeClosureExpr(expr *ast.ClosureExpr) Val {
	panic(fmt.Sprintf("todo executeClosureExpr"))
	return nil
}
func (e *executor) executeClosureUseExpr(expr *ast.ClosureUseExpr) Val {
	panic(fmt.Sprintf("todo executeClosureUseExpr"))
	return nil
}
func (e *executor) executeArrowFunctionExpr(expr *ast.ArrowFunctionExpr) Val {
	panic(fmt.Sprintf("todo executeArrowFunctionExpr"))
	return nil
}
func (e *executor) executeIndexExpr(expr *ast.IndexExpr) Val {
	panic(fmt.Sprintf("todo executeIndexExpr"))
	return nil
}
func (e *executor) executeCastExpr(expr *ast.CastExpr) Val {
	switch expr.Kind {
	case ast.CastArray:
	case ast.CastBool:
	case ast.CastDouble:
	case ast.CastInt:
	case ast.CastObject:
	case ast.CastString:
	case ast.CastUnset:
	}
	return nil
}
func (e *executor) executeUnaryExpr(expr *ast.UnaryExpr) Val {
	panic(fmt.Sprintf("todo executeUnaryExpr"))
	return nil
}
func (e *executor) executeAssignExpr(expr *ast.AssignExpr) Val {
	panic(fmt.Sprintf("todo executeAssignExpr"))
	return nil
}
func (e *executor) executeAssignOpExpr(expr *ast.AssignOpExpr) Val {
	panic(fmt.Sprintf("todo executeAssignOpExpr"))
	return nil
}
func (e *executor) executeAssignRefExpr(expr *ast.AssignRefExpr) Val {
	panic(fmt.Sprintf("todo executeAssignRefExpr"))
	return nil
}
func (e *executor) executeIssetExpr(expr *ast.IssetExpr) Val {
	panic(fmt.Sprintf("todo executeIssetExpr"))
	return nil
}
func (e *executor) executeEmptyExpr(expr *ast.EmptyExpr) Val {
	panic(fmt.Sprintf("todo executeEmptyExpr"))
	return nil
}
func (e *executor) executeEvalExpr(expr *ast.EvalExpr) Val {
	panic(fmt.Sprintf("todo executeEvalExpr"))
	return nil
}
func (e *executor) executeIncludeExpr(expr *ast.IncludeExpr) Val {
	panic(fmt.Sprintf("todo executeIncludeExpr"))
	return nil
}
func (e *executor) executeCloneExpr(expr *ast.CloneExpr) Val {
	panic(fmt.Sprintf("todo executeCloneExpr"))
	return nil
}
func (e *executor) executeErrorSuppressExpr(expr *ast.ErrorSuppressExpr) Val {
	panic(fmt.Sprintf("todo executeErrorSuppressExpr"))
	return nil
}
func (e *executor) executeExitExpr(expr *ast.ExitExpr) Val {
	panic(fmt.Sprintf("todo executeExitExpr"))
	return nil
}
func (e *executor) executeConstFetchExpr(expr *ast.ConstFetchExpr) Val {
	panic(fmt.Sprintf("todo executeConstFetchExpr"))
	return nil
}
func (e *executor) executeClassConstFetchExpr(expr *ast.ClassConstFetchExpr) Val {
	panic(fmt.Sprintf("todo executeClassConstFetchExpr"))
	return nil
}
func (e *executor) executeMagicConstExpr(expr *ast.MagicConstExpr) Val {
	panic(fmt.Sprintf("todo executeMagicConstExpr"))
	return nil
}
func (e *executor) executeInstanceofExpr(expr *ast.InstanceofExpr) Val {
	panic(fmt.Sprintf("todo executeInstanceofExpr"))
	return nil
}
func (e *executor) executeListExpr(expr *ast.ListExpr) Val {
	panic(fmt.Sprintf("todo executeListExpr"))
	return nil
}
func (e *executor) executePrintExpr(expr *ast.PrintExpr) Val {
	panic(fmt.Sprintf("todo executePrintExpr"))
	return nil
}
func (e *executor) executePropertyFetchExpr(expr *ast.PropertyFetchExpr) Val {
	panic(fmt.Sprintf("todo executePropertyFetchExpr"))
	return nil
}
func (e *executor) executeStaticPropertyFetchExpr(expr *ast.StaticPropertyFetchExpr) Val {
	panic(fmt.Sprintf("todo executeStaticPropertyFetchExpr"))
	return nil
}
func (e *executor) executeShellExecExpr(expr *ast.ShellExecExpr) Val {
	panic(fmt.Sprintf("todo executeShellExecExpr"))
	return nil
}
func (e *executor) executeTernaryExpr(expr *ast.TernaryExpr) Val {
	panic(fmt.Sprintf("todo executeTernaryExpr"))
	return nil
}
func (e *executor) executeThrowExpr(expr *ast.ThrowExpr) Val {
	panic(fmt.Sprintf("todo executeThrowExpr"))
	return nil
}
func (e *executor) executeVariableExpr(expr *ast.VariableExpr) Val {
	panic(fmt.Sprintf("todo executeVariableExpr"))
	return nil
}
func (e *executor) executeYieldExpr(expr *ast.YieldExpr) Val {
	panic(fmt.Sprintf("todo executeYieldExpr"))
	return nil
}
func (e *executor) executeYieldFromExpr(expr *ast.YieldFromExpr) Val {
	panic(fmt.Sprintf("todo executeYieldFromExpr"))
	return nil
}
func (e *executor) executeFuncCallExpr(expr *ast.FuncCallExpr) Val {
	var name Val
	switch nameAst := expr.Name.(type) {
	case *ast.Name:
		name = String(nameAst.ToString())
	case ast.Expr:
		name = e.expr(nameAst)
	default:
		panic("unreachable")
	}

	var fn *types.Function
	if name.IsString() {
		fn = e.initStringCall(name.String())
	} else {
		// todo 各种类型的 function name 处理
		panic(fmt.Sprintf("todo executeFuncCallExpr"))
	}

	args := make([]Val, 0, len(expr.Args))
	for _, arg := range expr.Args {
		argVal := e.expr(arg.Value)

		if !arg.Unpack {
			args = append(args, argVal)
		} else {
			// todo unpack args
			panic("todo unpack args")
		}
	}

	return e.function(fn, args)
}
func (e *executor) executeNewExpr(expr *ast.NewExpr) Val {
	panic(fmt.Sprintf("todo executeNewExpr"))
	return nil
}
func (e *executor) executeMethodCallExpr(expr *ast.MethodCallExpr) Val {
	panic(fmt.Sprintf("todo executeMethodCallExpr"))
	return nil
}
func (e *executor) executeStaticCallExpr(expr *ast.StaticCallExpr) Val {
	panic(fmt.Sprintf("todo executeStaticCallExpr"))
	return nil
}

package php

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/operators"
	"github.com/heyuuu/gophp/php/types"
)

// public functions
func ExecuteAstFunction(ctx *Context, executeData *ExecuteData, f *types.Function) (Val, error) {
	Assert(f.IsUserFunction())
	executor := &astExecutor{ctx: ctx, executeData: executeData}
	return executor.executeAstFile(f.AstFile())
}

// errors
type ExecutorError string

func (e ExecutorError) Error() string { return string(e) }

//
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
type astExecutor struct {
	ctx         *Context
	executeData *ExecuteData
	sources     Sources
	currFile    *ast.File
	currNs      *ast.NamespaceStmt
	currRetVal  Val
	operator    *operators.Operator
}

func (e *astExecutor) executeFile(filePath string) (Val, error) {
	code, ok := e.sources.LoadSource(filePath)
	if !ok {
		return nil, ExecutorError("source file is not exist")
	}

	astFile, err := parser.ParseCode(code)
	if err != nil {
		return nil, err
	}

	return e.executeAstFile(astFile)
}

func (e *astExecutor) executeAstFile(f *ast.File) (Val, error) {
	// todo f.Declares

	e.currFile = f
	for _, ns := range f.Namespaces {
		e.currNs = ns
		res, err := e.executeStmts(ns.Stmts)
		if err != nil {
			return nil, err
		}
		switch r := res.(type) {
		case *returnResult:
			return r.retVal, nil
		case *continueResult, *breakResult, *gotoResult:
			panic("unreachable")
		}
	}
	return nil, nil
}

func (e *astExecutor) executeStmts(stmts []ast.Stmt) (result executeResult, err error) {
	var labels = map[string]int{}
	for i, stmt := range stmts {
		if label, ok := stmt.(*ast.LabelStmt); ok {
			labels[label.Name.Name] = i
		}
	}

	l := len(stmts)
	for i := 0; i < l && err == nil; i++ {
		switch x := stmts[i].(type) {
		case *ast.EmptyStmt: // pass
		case *ast.ExprStmt:
			_, err = e.executeExpr(x.Expr)
		case *ast.ReturnStmt:
			retVal, err := e.executeExpr(x.Expr)
			return returnResult{retVal: retVal}, err
		case *ast.LabelStmt:
			// pass
			// todo goto 能跳到非循环结构内(比如 if)
		case *ast.GotoStmt:
			// todo goto 处理逻辑
			labelName := x.Name.Name
			if v, ok := labels[labelName]; ok {
				i = v
			} else {
				return gotoResult{labelName}, err
			}
		case *ast.EchoStmt:
			values, err := e.executeExprs(x.Exprs)
			if err != nil {
				return nil, err
			}
			for _, value := range values {
				vmEcho(e.ctx, value)
			}
		// todo
		default:
			panic(fmt.Sprintf("todo executor.executeStmts(%T)", x))
		}
	}
	return
}

func (e *astExecutor) executeExprs(exprs []ast.Expr) (values []Val, err error) {
	values = make([]Val, len(exprs))
	for i, expr := range exprs {
		values[i], err = e.executeExpr(expr)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (e *astExecutor) executeExpr(expr ast.Expr) (val Val, err error) {
	switch x := expr.(type) {
	case *ast.IntLit:
		return Long(x.Value), nil
	case *ast.FloatLit:
		return Double(x.Value), nil
	case *ast.StringLit:
		return String(x.Value), nil
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

func (e *astExecutor) executeBinaryOpExpr(expr *ast.BinaryOpExpr) (val Val, err error) {
	// && / || / ?? 操作比较特殊，右表达式节点可能不会执行
	switch expr.Op {
	case ast.BinaryOpBooleanAnd: // &&
		if left, err := e.executeExpr(expr.Left); err != nil {
			return nil, err
		} else if !ZvalIsTrue(left) {
			return False(), nil
		}

		if right, err := e.executeExpr(expr.Right); err != nil {
			return nil, err
		} else {
			return Bool(ZvalIsTrue(right)), nil
		}
	case ast.BinaryOpBooleanOr: // ||
		if left, err := e.executeExpr(expr.Left); err != nil {
			return nil, err
		} else if ZvalIsTrue(left) {
			return True(), nil
		}

		if right, err := e.executeExpr(expr.Right); err != nil {
			return nil, err
		} else {
			return Bool(ZvalIsTrue(right)), nil
		}
	case ast.BinaryOpCoalesce: // ??
		if left, err := e.executeExpr(expr.Left); err != nil {
			return nil, err
		} else if !left.IsUndef() && !left.IsNull() {
			return left, nil
		}

		return e.executeExpr(expr.Right)
	}

	// common
	left, err := e.executeExpr(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := e.executeExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Op {
	case ast.BinaryOpPlus: // +
		return vmBinaryOp(e.ctx, left, right, operators.Add)
	case ast.BinaryOpMinus: // -
		return vmBinaryOp(e.ctx, left, right, operators.Sub)
	case ast.BinaryOpMul: // *
		return vmBinaryOp(e.ctx, left, right, operators.Mul)
	case ast.BinaryOpDiv: // /
		return vmBinaryOp(e.ctx, left, right, operators.Div)
	case ast.BinaryOpMod: // %
		return vmBinaryOp(e.ctx, left, right, operators.Mod)
	case ast.BinaryOpPow: // **
		return vmBinaryOp(e.ctx, left, right, operators.Pow)
	case ast.BinaryOpBitwiseAnd: // &
		return vmBinaryOp(e.ctx, left, right, operators.BitwiseAnd)
	case ast.BinaryOpBitwiseOr: // n|
		return vmBinaryOp(e.ctx, left, right, operators.BitwiseOr)
	case ast.BinaryOpBitwiseXor: // ^
		return vmBinaryOp(e.ctx, left, right, operators.BitwiseXor)
	case ast.BinaryOpConcat: // .
		return vmBinaryOp(e.ctx, left, right, operators.Concat)
	case ast.BinaryOpEqual: // ==
		return vmBinaryOp(e.ctx, left, right, operators.Equal)
	case ast.BinaryOpGreater: // >
		return vmBinaryOp(e.ctx, left, right, operators.Greater)
	case ast.BinaryOpGreaterOrEqual: // >=
		return vmBinaryOp(e.ctx, left, right, operators.GreaterOrEqual)
	case ast.BinaryOpIdentical: // ===
		return vmBinaryOp(e.ctx, left, right, operators.Identical)
	case ast.BinaryOpBooleanXor: // xor
		return vmBinaryOp(e.ctx, left, right, operators.BooleanXor)
	case ast.BinaryOpNotEqual: // !=
		return vmBinaryOp(e.ctx, left, right, operators.NotEqual)
	case ast.BinaryOpNotIdentical: // !==
		return vmBinaryOp(e.ctx, left, right, operators.NotIdentical)
	case ast.BinaryOpShiftLeft: // <<
		return vmBinaryOp(e.ctx, left, right, operators.ShiftLeft)
	case ast.BinaryOpShiftRight: // >>
		return vmBinaryOp(e.ctx, left, right, operators.ShiftRight)
	case ast.BinaryOpSmaller: // <
		return vmBinaryOp(e.ctx, left, right, operators.Smaller)
	case ast.BinaryOpSmallerOrEqual: // <=
		return vmBinaryOp(e.ctx, left, right, operators.SmallerOrEqual)
	case ast.BinaryOpSpaceship: // <=>
		return vmBinaryOp(e.ctx, left, right, operators.Spaceship)
	default:
		panic("unreachable")
	}
}
func (e *astExecutor) executeArrayExpr(expr *ast.ArrayExpr) (val Val, err error) {
	arr := types.NewArrayCap(len(expr.Items))
	for _, item := range expr.Items {
		if item.ByRef {
			// todo item byref
			panic("todo item byref")
		} else if item.Unpack && item.Key != nil {
			// todo item unpack with key
			panic("todo item unpack with key")
		}

		var key, value Val
		if item.Key != nil {
			key, err = e.executeExpr(item.Key)
			if err != nil {
				return
			}
		}
		value, err = e.executeExpr(item.Value)
		if err != nil {
			return
		}

		if key == nil {
			// todo array add
		} else {
			// todo array add
		}
	}

	panic(fmt.Sprintf("todo executeArrayExpr"))
	return
}
func (e *astExecutor) executeClosureExpr(expr *ast.ClosureExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeClosureExpr"))
	return
}
func (e *astExecutor) executeClosureUseExpr(expr *ast.ClosureUseExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeClosureUseExpr"))
	return
}
func (e *astExecutor) executeArrowFunctionExpr(expr *ast.ArrowFunctionExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeArrowFunctionExpr"))
	return
}
func (e *astExecutor) executeIndexExpr(expr *ast.IndexExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeIndexExpr"))
	return
}
func (e *astExecutor) executeCastExpr(expr *ast.CastExpr) (val Val, err error) {
	switch expr.Kind {
	case ast.CastArray:
	case ast.CastBool:
	case ast.CastDouble:
	case ast.CastInt:
	case ast.CastObject:
	case ast.CastString:
	case ast.CastUnset:
	}
	return
}
func (e *astExecutor) executeUnaryExpr(expr *ast.UnaryExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeUnaryExpr"))
	return
}
func (e *astExecutor) executeAssignExpr(expr *ast.AssignExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeAssignExpr"))
	return
}
func (e *astExecutor) executeAssignOpExpr(expr *ast.AssignOpExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeAssignOpExpr"))
	return
}
func (e *astExecutor) executeAssignRefExpr(expr *ast.AssignRefExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeAssignRefExpr"))
	return
}
func (e *astExecutor) executeIssetExpr(expr *ast.IssetExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeIssetExpr"))
	return
}
func (e *astExecutor) executeEmptyExpr(expr *ast.EmptyExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeEmptyExpr"))
	return
}
func (e *astExecutor) executeEvalExpr(expr *ast.EvalExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeEvalExpr"))
	return
}
func (e *astExecutor) executeIncludeExpr(expr *ast.IncludeExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeIncludeExpr"))
	return
}
func (e *astExecutor) executeCloneExpr(expr *ast.CloneExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeCloneExpr"))
	return
}
func (e *astExecutor) executeErrorSuppressExpr(expr *ast.ErrorSuppressExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeErrorSuppressExpr"))
	return
}
func (e *astExecutor) executeExitExpr(expr *ast.ExitExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeExitExpr"))
	return
}
func (e *astExecutor) executeConstFetchExpr(expr *ast.ConstFetchExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeConstFetchExpr"))
	return
}
func (e *astExecutor) executeClassConstFetchExpr(expr *ast.ClassConstFetchExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeClassConstFetchExpr"))
	return
}
func (e *astExecutor) executeMagicConstExpr(expr *ast.MagicConstExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeMagicConstExpr"))
	return
}
func (e *astExecutor) executeInstanceofExpr(expr *ast.InstanceofExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeInstanceofExpr"))
	return
}
func (e *astExecutor) executeListExpr(expr *ast.ListExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeListExpr"))
	return
}
func (e *astExecutor) executePrintExpr(expr *ast.PrintExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executePrintExpr"))
	return
}
func (e *astExecutor) executePropertyFetchExpr(expr *ast.PropertyFetchExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executePropertyFetchExpr"))
	return
}
func (e *astExecutor) executeStaticPropertyFetchExpr(expr *ast.StaticPropertyFetchExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeStaticPropertyFetchExpr"))
	return
}
func (e *astExecutor) executeShellExecExpr(expr *ast.ShellExecExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeShellExecExpr"))
	return
}
func (e *astExecutor) executeTernaryExpr(expr *ast.TernaryExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeTernaryExpr"))
	return
}
func (e *astExecutor) executeThrowExpr(expr *ast.ThrowExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeThrowExpr"))
	return
}
func (e *astExecutor) executeVariableExpr(expr *ast.VariableExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeVariableExpr"))
	return
}
func (e *astExecutor) executeYieldExpr(expr *ast.YieldExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeYieldExpr"))
	return
}
func (e *astExecutor) executeYieldFromExpr(expr *ast.YieldFromExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeYieldFromExpr"))
	return
}
func (e *astExecutor) executeFuncCallExpr(expr *ast.FuncCallExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeFuncCallExpr"))
	return
}
func (e *astExecutor) executeNewExpr(expr *ast.NewExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeNewExpr"))
	return
}
func (e *astExecutor) executeMethodCallExpr(expr *ast.MethodCallExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeMethodCallExpr"))
	return
}
func (e *astExecutor) executeStaticCallExpr(expr *ast.StaticCallExpr) (val Val, err error) {
	panic(fmt.Sprintf("todo executeStaticCallExpr"))
	return
}

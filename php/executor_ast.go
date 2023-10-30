package php

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/types"
)

// public functions
func ExecuteAstFunction(ctx *Context, executeData *ExecuteData, f *types.AstFunction) (Val, error) {
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
		val = Long(x.Value)
	case *ast.FloatLit:
		val = Double(x.Value)
	case *ast.StringLit:
		val = String(x.Value)
	case *ast.ArrayExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ArrayItemExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ClosureExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ClosureUseExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ArrowFunctionExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.IndexExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.CastExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.UnaryExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.BinaryOpExpr:
		return e.executeBinaryOpExpr(x)
	case *ast.AssignExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.AssignOpExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.AssignRefExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.IssetExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.EmptyExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.EvalExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.IncludeExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.CloneExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ErrorSuppressExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ExitExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ConstFetchExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ClassConstFetchExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.MagicConstExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.InstanceofExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ListExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.PrintExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.PropertyFetchExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.StaticPropertyFetchExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ShellExecExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.TernaryExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.ThrowExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.VariableExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.YieldExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.YieldFromExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.FuncCallExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.NewExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.MethodCallExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	case *ast.StaticCallExpr:
		// todo
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	default:
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	}

	return
}

func (e *astExecutor) executeBinaryOpExpr(expr *ast.BinaryOpExpr) (val Val, err error) {
	left, err := e.executeExpr(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := e.executeExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Op {
	case ast.BinaryOpBitwiseAnd: // &
		panic("todo BinaryOpBitwiseAnd")
	case ast.BinaryOpBitwiseOr: // |
		panic("todo BinaryOpBitwiseOr")
	case ast.BinaryOpBitwiseXor: // ^
		panic("todo BinaryOpBitwiseXor")
	case ast.BinaryOpBooleanAnd: // &&
		panic("todo BinaryOpBooleanAnd")
	case ast.BinaryOpBooleanOr: // ||
		panic("todo BinaryOpBooleanOr")
	case ast.BinaryOpCoalesce: // ??
		panic("todo BinaryOpCoalesce")
	case ast.BinaryOpConcat: // .
		panic("todo BinaryOpConcat")
	case ast.BinaryOpDiv: // /
		panic("todo BinaryOpDiv")
	case ast.BinaryOpEqual: // ==
		panic("todo BinaryOpEqual")
	case ast.BinaryOpGreater: // >
		panic("todo BinaryOpGreater")
	case ast.BinaryOpGreaterOrEqual: // >=
		panic("todo BinaryOpGreaterOrEqual")
	case ast.BinaryOpIdentical: // ===
		panic("todo BinaryOpIdentical")
	case ast.BinaryOpLogicalAnd: // and
		panic("todo BinaryOpLogicalAnd")
	case ast.BinaryOpLogicalOr: // or
		panic("todo BinaryOpLogicalOr")
	case ast.BinaryOpLogicalXor: // xor
		panic("todo BinaryOpLogicalXor")
	case ast.BinaryOpMinus: // -
		panic("todo BinaryOpMinus")
	case ast.BinaryOpMod: // %
		panic("todo BinaryOpMod")
	case ast.BinaryOpMul: // *
		panic("todo BinaryOpMul")
	case ast.BinaryOpNotEqual: // !=
		panic("todo BinaryOpNotEqual")
	case ast.BinaryOpNotIdentical: // !==
		panic("todo BinaryOpNotIdentical")
	case ast.BinaryOpPlus: // +
		return vmAdd(e.ctx, left, right)
	case ast.BinaryOpPow: // **
		panic("todo BinaryOpPow")
	case ast.BinaryOpShiftLeft: // <<
		panic("todo BinaryOpShiftLeft")
	case ast.BinaryOpShiftRight: // >>
		panic("todo BinaryOpShiftRight")
	case ast.BinaryOpSmaller: // <
		panic("todo BinaryOpSmaller")
	case ast.BinaryOpSmallerOrEqual: // <=
		panic("todo BinaryOpSmallerOrEqual")
	case ast.BinaryOpSpaceship: // <=>
		panic("todo BinaryOpSpaceship")
	default:
		panic("unreachable")
	}
}

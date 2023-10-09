package php

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/types"
	"os"
)

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

//
func Default() *Executor {
	return NewExecutor()
}

type Executor struct {
	sources    Sources
	currFile   *ast.File
	currNs     *ast.NamespaceStmt
	currRetVal Val
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) RunCode(code string) error {
	sources := NewSourcesByCode(code)
	return e.Run(sources, DefaultSourcePath)
}

func (e *Executor) Run(sources Sources, enterFile string) error {
	e.sources = sources
	_, err := e.executeFile(enterFile)
	return err
}

func (e *Executor) executeFile(filePath string) (Val, error) {
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

func (e *Executor) executeAstFile(f *ast.File) (Val, error) {
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

func (e *Executor) executeStmts(stmts []ast.Stmt) (result executeResult, err error) {
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
				// todo change writer
				vmEcho(os.Stdout, value)
			}

		// todo
		default:
			panic(fmt.Sprintf("todo executor.executeStmts(%T)", x))
		}
	}
	return
}

func (e *Executor) executeExprs(exprs []ast.Expr) (values []Val, err error) {
	values = make([]Val, len(exprs))
	for i, expr := range exprs {
		values[i], err = e.executeExpr(expr)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (e *Executor) executeExpr(expr ast.Expr) (val Val, err error) {
	switch x := expr.(type) {
	case *ast.IntLit:
		val = types.NewZvalLong(x.Value)
	case *ast.StringLit:
		val = types.NewZvalString(x.Value)
	// todo
	default:
		panic(fmt.Sprintf("todo executor.executeExpr(%T)", x))
	}

	return
}

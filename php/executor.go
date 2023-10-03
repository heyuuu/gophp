package php

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/types"
)

// errors
type ExecutorError string

func (e ExecutorError) Error() string { return string(e) }

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
	e.currFile, e.currRetVal = f, nil
	for _, ns := range f.Namespaces {
		if err := e.executeNs(ns); err != nil {
			return e.currRetVal, err
		}
	}
	return e.currRetVal, nil
}

func (e *Executor) executeNs(ns *ast.NamespaceStmt) error {
	for _, stmt := range ns.Stmts {
		if err := e.executeStmt(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (e *Executor) executeStmt(stmt ast.Stmt) (err error) {
	switch x := stmt.(type) {
	case *ast.ExprStmt:
		e.executeExpr(x.Expr)
		// todo
	}
	return
}

func (e *Executor) executeExpr(expr ast.Expr) (val Val, err error) {
	switch x := expr.(type) {
	case *ast.IntLit:
		val = types.NewZvalLong(x.Value)
		// todo
	}

	return
}

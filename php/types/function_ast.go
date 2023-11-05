package types

import (
	"github.com/heyuuu/gophp/compile/ast"
)

func NewAstFunction(name string, argInfos []ArgInfo, stmts []ast.Stmt) *Function {
	return &Function{
		typ:          TypeUserFunction,
		functionName: name,
		argInfos:     argInfos,
		stmts:        stmts,
	}
}

func NewAstTopFunction(astFile *ast.File) *Function {
	var stmts []ast.Stmt
	for _, namespace := range astFile.Namespaces {
		stmts = append(stmts, namespace.Stmts...)
	}

	return &Function{
		typ:   TypeUserFunction,
		stmts: stmts,
	}
}

func (f *Function) Stmts() []ast.Stmt  { return f.stmts }
func (f *Function) AstFile() *ast.File { return f.astFile }

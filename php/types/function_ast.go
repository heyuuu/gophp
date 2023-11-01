package types

import (
	"github.com/heyuuu/gophp/compile/ast"
)

type AstFunction struct {
	name     string
	argInfos []ArgInfo
}

func NewAstFunction(name string, argInfos []ArgInfo, stmts []ast.Stmt) *Function {
	return &Function{
		typ:          TypeUserFunction,
		functionName: name,
		argInfos:     argInfos,
		stmts:        stmts,
	}
}

func NewAstTopFunction(ast *ast.File) *Function {
	return &Function{
		typ:     TypeUserFunction,
		astFile: ast,
	}
}

func (f *Function) Stmts() []ast.Stmt  { return f.stmts }
func (f *Function) AstFile() *ast.File { return f.astFile }

package types

import (
	"github.com/heyuuu/gophp/compile/ast"
)

var _ Function = (*AstFunction)(nil)

type AstFunction struct {
	name     string
	argInfos []ArgInfo
	stmts    []ast.Stmt
	astFile  *ast.File
}

func NewAstFunction(name string, argInfos []ArgInfo, stmts []ast.Stmt) *AstFunction {
	return &AstFunction{name: name, argInfos: argInfos, stmts: stmts}
}

func NewAstTopFunction(ast *ast.File) *AstFunction {
	// todo
	return &AstFunction{astFile: ast}
}

func (f *AstFunction) Name() string        { return f.name }
func (f *AstFunction) ArgInfos() []ArgInfo { return f.argInfos }
func (f *AstFunction) Stmts() []ast.Stmt   { return f.stmts }
func (f *AstFunction) AstFile() *ast.File  { return f.astFile }

package types

import "github.com/heyuuu/gophp/compile/ast"

type Function struct {
	data any
}

func (f *Function) GetUserAstFile() *ast.File {
	astFile, ok := f.data.(*ast.File)
	if ok {
		return astFile
	}
	return nil
}

func NewUserFunction(ast *ast.File) *Function {
	// todo
	return &Function{data: ast}
}

package astutil

import (
	"go/ast"
	"go/token"
	"strconv"
)

func Ident(name string) *ast.Ident { return &ast.Ident{Name: name} }

func NilLit() *ast.Ident { return Ident("nil") }
func IntLit(val int) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(val)}
}
func StrLit(val string) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(val)}
}

func KeyValue(key string, value ast.Expr) *ast.KeyValueExpr {
	return &ast.KeyValueExpr{Key: Ident(key), Value: value}
}

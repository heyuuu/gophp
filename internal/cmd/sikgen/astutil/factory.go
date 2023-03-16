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

func Field(ident *ast.Ident, typ ast.Expr) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{ident},
		Type:  typ,
	}
}
func Fields(fields ...*ast.Field) *ast.FieldList {
	return &ast.FieldList{List: fields}
}

func Type(name string) ast.Expr     { return Ident(name) }
func RefType(typ ast.Expr) ast.Expr { return &ast.StarExpr{X: typ} }

func CallExpr(name string, args []ast.Expr) ast.Expr {
	return &ast.CallExpr{
		Fun:  Ident(name),
		Args: args,
	}
}

func MethodCallExpr(instance ast.Expr, method string, args []ast.Expr) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   instance,
			Sel: Ident(method),
		},
		Args: args,
	}
}

func ExprStmt(expr ast.Expr) ast.Stmt {
	return &ast.ExprStmt{X: expr}
}
func AssignStmt(variable ast.Expr, value ast.Expr) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{variable},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{value},
	}
}

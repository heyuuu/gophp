package astutil

import (
	"go/ast"
	"go/token"
	"strconv"
)

var (
	nilIdent   = Ident("nil")
	falseIdent = Ident("false")
	trueIdent  = Ident("true")
)

func Ident(name string) *ast.Ident { return &ast.Ident{Name: name} }
func PkgIdent(pkg string, name string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   &ast.Ident{Name: pkg},
		Sel: &ast.Ident{Name: name},
	}
}

func NilLit() *ast.Ident { return nilIdent }
func IntLit(val int) *ast.BasicLit {
	return &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(val)}
}
func BoolLit(val bool) ast.Expr {
	if val {
		return trueIdent
	}
	return falseIdent
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

func Not(expr ast.Expr) ast.Expr { return &ast.UnaryExpr{Op: token.NOT, X: expr} }

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
func BlockStmt(list ...ast.Stmt) *ast.BlockStmt {
	return &ast.BlockStmt{List: list}
}

func DocComment(comments ...string) *ast.CommentGroup {
	if len(comments) == 0 {
		return nil
	}

	var list []*ast.Comment
	for _, comment := range comments {
		list = append(list, &ast.Comment{
			Text: comment,
		})
	}
	return &ast.CommentGroup{List: list}
}

func ValueSpecDecl(variable *ast.Ident, value ast.Expr) *ast.GenDecl {
	return ValueSpecDeclEx(nil, variable, value)
}

func ValueSpecDeclEx(doc *ast.CommentGroup, variable *ast.Ident, value ast.Expr) *ast.GenDecl {
	return &ast.GenDecl{
		Doc: doc,
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names:  []*ast.Ident{variable},
				Values: []ast.Expr{value},
			},
		},
	}
}

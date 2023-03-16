package zif

import (
	"go/ast"
	"go/token"
	"log"
	f "sik/internal/cmd/sikgen/astutil"
)

func genFileNode(name *ast.Ident, infos []*ZifInfo) *ast.File {
	var decls []ast.Decl
	for _, zifInfo := range infos {
		docComment := &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "\n// generate by " + zifInfo.funcName},
			},
		}

		varExpr := &ast.Ident{Name: zifInfo.defName}
		valueExpr := &ast.CallExpr{
			Fun:  f.Ident("DefFunc"),
			Args: []ast.Expr{genDefFuncOpts(zifInfo)},
		}

		decls = append(decls, &ast.GenDecl{
			Doc: docComment,
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names:  []*ast.Ident{varExpr},
					Values: []ast.Expr{valueExpr},
				},
			},
		})
	}
	return &ast.File{Name: name, Decls: decls}
}

func genDefFuncOpts(zifInfo *ZifInfo) ast.Expr {
	// 构建 DefFuncOpts 字段
	var optElements []ast.Expr
	optElements = append(optElements, f.KeyValue("name", f.StrLit(zifInfo.name)))
	if zifInfo.minNumArgs > 0 {
		optElements = append(optElements, f.KeyValue("minNumArgs", f.IntLit(zifInfo.minNumArgs)))
	}
	optElements = append(optElements, f.KeyValue("handler", genZifHandler(zifInfo)))

	// 构建结构体字面量
	return &ast.CompositeLit{
		Type: f.Ident("DefFuncOpts"),
		Elts: optElements,
	}
}

func genZifHandler(zifInfo *ZifInfo) ast.Expr {
	executeDataIdent := f.Ident("executeData")
	returnValueIdent := f.Ident("returnValue")
	retIdent := f.Ident("ret")

	// type
	funcType := &ast.FuncType{
		Params: f.Fields(
			f.Field(executeDataIdent, f.RefType(f.Type("ZendExecuteData"))),
			f.Field(returnValueIdent, f.RefType(f.Type("Zval"))),
		),
	}

	// body
	var args []ast.Expr
	var stmts []ast.Stmt

	if zifInfo.maxNumArgs >= 0 {
		stmts = append(stmts, &ast.IfStmt{
			Cond: f.Not(
				f.MethodCallExpr(executeDataIdent, "CheckNumArgs", []ast.Expr{
					f.IntLit(zifInfo.minNumArgs),
					f.IntLit(zifInfo.maxNumArgs),
					f.BoolLit(zifInfo.strict),
				}),
			),
			Body: f.BlockStmt(&ast.ReturnStmt{}),
		})
	} else if zifInfo.minNumArgs > 0 {
		stmts = append(stmts, &ast.IfStmt{
			Cond: f.Not(
				f.MethodCallExpr(executeDataIdent, "CheckMinNumArgs", []ast.Expr{
					f.IntLit(zifInfo.minNumArgs),
					f.BoolLit(zifInfo.strict),
				}),
			),
			Body: f.BlockStmt(&ast.ReturnStmt{}),
		})
	}

	if zifInfo.returnArgInfo == nil {
		stmts = append(stmts,
			f.ExprStmt(f.CallExpr(zifInfo.funcName, args)),
		)
	} else {
		var setter string
		switch zifInfo.returnArgInfo.typ {
		case ZvalTypeBool:
			setter = "SetBool"
		case ZvalTypeInt:
			setter = "SetLong"
		case ZvalTypeDouble:
			setter = "SetDouble"
		case ZvalTypeString:
			setter = "SetRawString"
		default:
			log.Fatalln("此 ZvalType 未设置对应 Setter: " + zifInfo.returnArgInfo.typ)
		}

		stmts = append(stmts,
			f.AssignStmt(retIdent, f.CallExpr(zifInfo.funcName, args)),
			f.ExprStmt(f.MethodCallExpr(returnValueIdent, setter, []ast.Expr{retIdent})),
		)
	}

	return &ast.FuncLit{
		Type: funcType,
		Body: &ast.BlockStmt{List: stmts},
	}
}

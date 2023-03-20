package zif

import (
	"go/ast"
	"log"
	f "sik/internal/cmd/sikgen/astutil"
)

var (
	typeEx   = f.RefType(f.Type("ZendExecuteData"))
	typeZval = f.RefType(f.PkgIdent("types", "Zval"))
)

func genFileNode(name string, infos []*ZifInfo) *ast.File {
	fb := f.NewFileBuilder(name)
	fb.AddImport("sik/zend/types")
	fb.AddImport("sik/zend/zpp")

	for _, zifInfo := range infos {
		fb.AddDecl(
			f.ValueSpecDeclEx(
				f.DocComment("\n// generate by "+zifInfo.funcName),
				f.Ident(zifInfo.defName),
				&ast.CallExpr{
					Fun:  f.Ident("DefFunc"),
					Args: []ast.Expr{genDefFuncOpts(zifInfo)},
				},
			),
		)
	}

	return fb.Build()
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
			f.Field(executeDataIdent, typeEx),
			f.Field(returnValueIdent, typeZval),
		),
	}

	// body
	var args []ast.Expr
	var stmts []ast.Stmt

	if zifInfo.minNumArgs == 0 && zifInfo.maxNumArgs <= 0 {
		method := "CheckNumArgsNoneError"
		if zifInfo.strict {
			method = "CheckNumArgsNoneException"
		}
		stmts = append(stmts, &ast.IfStmt{
			Cond: f.Not(
				f.PkgCallExpr("zpp", method, []ast.Expr{
					executeDataIdent,
				}),
			),
			Body: f.BlockStmt(&ast.ReturnStmt{}),
		})
	} else if zifInfo.maxNumArgs >= 0 {
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

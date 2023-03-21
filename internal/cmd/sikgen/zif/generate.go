package zif

import (
	"go/ast"
	"log"
	f "sik/internal/cmd/sikgen/astutil"
)

var (
	// types
	typeEx   = f.RefType(f.Type("ZendExecuteData"))
	typeZval = f.RefType(f.PkgIdent("types", "Zval"))
	// variables
	fpIdent = f.Ident("fp")
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
	log.Printf("ZifInfo: %+v\n", *zifInfo)

	// 构建 ArgInfos
	var realArgInfos []ast.Expr
	for _, argInfo := range zifInfo.argInfos {
		switch argInfo.typ {
		case ZppTypeEx, ZppTypeRet, ZppTypeOpt:
			continue
		default:
			realArgInfos = append(realArgInfos, &ast.CompositeLit{
				Elts: []ast.Expr{
					f.KeyValue("name", f.StrLit(argInfo.name)),
				},
			})
		}
	}

	// 构建 DefFuncOpts 字段
	var optElements []ast.Expr
	optElements = append(optElements, f.KeyValue("name", f.StrLit(zifInfo.name)))
	optElements = append(optElements, f.KeyValue("minNumArgs", f.IntLit(zifInfo.minNumArgs)))
	optElements = append(optElements, f.KeyValue("maxNumArgs", f.IntLit(zifInfo.maxNumArgs)))
	if len(realArgInfos) != 0 {
		optElements = append(optElements, f.KeyValue("argInfos", &ast.CompositeLit{
			Type: f.ArrayType(f.Ident("ArgInfo")),
			Elts: realArgInfos,
		}))
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
	var stmts []ast.Stmt

	// check num args
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
	} else {
		var flags ast.Expr
		if zifInfo.strict {
			flags = &ast.SelectorExpr{X: f.Ident("zpp"), Sel: f.Ident("ZEND_PARSE_PARAMS_THROW")}
		} else {
			flags = f.IntLit(0)
		}

		stmts = append(stmts, f.AssignStmt(
			fpIdent,
			f.PkgCallExpr("zpp", "FastParseStart", []ast.Expr{
				executeDataIdent,
				f.IntLit(zifInfo.minNumArgs),
				f.IntLit(zifInfo.maxNumArgs),
				flags,
			}),
		))
		for _, info := range zifInfo.argInfos {
			argTyp := info.typ
			if argTyp == ZppTypeEx || argTyp == ZppTypeRet {
				continue
			} else if argTyp == ZppTypeOpt {
				stmts = append(stmts, f.ExprStmt(f.PkgCallExpr("zpp", "StartOptional", nil)))
			} else if parseMethod, ok := toZppParseMethod(argTyp); ok {
				stmts = append(stmts, f.AssignStmt(
					f.Ident(info.name),
					f.MethodCallExpr(fpIdent, parseMethod, nil),
				))
			} else {
				log.Fatalf("Zpp类型未定义 Parse 方法: type=%d\n", argTyp)
			}
		}
		stmts = append(stmts, &ast.IfStmt{
			Cond: f.MethodCallExpr(fpIdent, "HasError", nil),
			Body: f.BlockStmt(&ast.ReturnStmt{}),
		})
	}

	var args []ast.Expr
	for _, argInfo := range zifInfo.argInfos {
		switch argInfo.typ {
		case ZppTypeEx:
			args = append(args, executeDataIdent)
		case ZppTypeRet:
			args = append(args, returnValueIdent)
		case ZppTypeOpt:
			args = append(args, f.NilLit())
		default:
			args = append(args, f.Ident(argInfo.name))
		}
	}

	if zifInfo.returnArgInfo == nil {
		stmts = append(stmts,
			f.ExprStmt(f.CallExpr(zifInfo.funcName, args)),
		)
	} else {
		var setter string
		switch zifInfo.returnArgInfo.typ {
		case ZppTypeBool:
			setter = "SetBool"
		case ZppTypeLong:
			setter = "SetLong"
		case ZppTypeDouble:
			setter = "SetDouble"
		case ZppTypeString:
			setter = "SetRawString"
		default:
			log.Fatalln("此 ZppType 未设置对应 Setter: " + zifInfo.returnArgInfo.typ.String())
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

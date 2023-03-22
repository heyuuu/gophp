package zif

import (
	"go/ast"
	"go/token"
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
	okIdent := f.Ident("ok")

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
		// fp := FastParseStart(...)
		var flags []ast.Expr
		if zifInfo.quiet {
			flag := &ast.SelectorExpr{X: f.Ident("zpp"), Sel: f.Ident("FlagQuiet")}
			flags = append(flags, flag)
		}
		if zifInfo.strict {
			flag := &ast.SelectorExpr{X: f.Ident("zpp"), Sel: f.Ident("FlagThrow")}
			flags = append(flags, flag)
		}
		if zifInfo.oldMode {
			flag := &ast.SelectorExpr{X: f.Ident("zpp"), Sel: f.Ident("FlagOldMode")}
			flags = append(flags, flag)
		}
		var flagsExpr ast.Expr
		if len(flags) != 0 {
			flagsExpr = f.BinaryExpr(token.OR, flags[0], flags[1:]...)
		} else {
			flagsExpr = f.IntLit(0)
		}

		// argN := fp.ParseXXX()
		stmts = append(stmts, f.AssignStmt(
			fpIdent,
			f.PkgCallExpr("zpp", "FastParseStart", []ast.Expr{
				executeDataIdent,
				f.IntLit(zifInfo.minNumArgs),
				f.IntLit(zifInfo.maxNumArgs),
				flagsExpr,
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

	// 代用内部方法
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
	realCallExpr := f.CallExpr(zifInfo.funcName, args)

	// 处理返回值
	retInfo := zifInfo.returnInfo
	if retInfo == nil {
		stmts = append(stmts, f.ExprStmt(realCallExpr))
	} else {
		// 返回值setter
		setter, ok := toZppSetMethod(retInfo.typ)
		if !ok {
			log.Fatalf("不支持此类型的返回值: typ=%d\n", retInfo.typ)
		}

		if retInfo.withOk {
			stmts = append(stmts,
				// ret, ok := realCall()
				f.MultiAssignStmt([]ast.Expr{retIdent, okIdent}, realCallExpr),
				// if ok { returnValue.SetXXX(ret); } else { returnValue.SetFalse() }
				&ast.IfStmt{
					Cond: okIdent,
					Body: f.BlockStmt(
						f.ExprStmt(f.MethodCallExpr(returnValueIdent, setter, []ast.Expr{retIdent})),
					),
					Else: f.BlockStmt(
						f.ExprStmt(f.MethodCallExpr(returnValueIdent, "SetFalse", nil)),
					),
				},
			)
		} else {
			stmts = append(stmts,
				// ret := realCall()
				f.MultiAssignStmt([]ast.Expr{retIdent}, realCallExpr),
				// returnValue.SetXXX(ret)
				f.ExprStmt(f.MethodCallExpr(returnValueIdent, setter, []ast.Expr{retIdent})),
			)
		}
	}

	return &ast.FuncLit{
		Type: funcType,
		Body: &ast.BlockStmt{List: stmts},
	}
}

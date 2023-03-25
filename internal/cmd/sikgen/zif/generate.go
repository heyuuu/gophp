package zif

import (
	"go/ast"
	"go/token"
	"log"
	"sik/builtin/strutil"
	f "sik/internal/cmd/sikgen/astutil"
)

var (
	// pkgIdent
	zppPkgIdent = func(name string) *ast.SelectorExpr { return f.PkgIdent("zpp", name) }
	defPkgIdent = func(name string) *ast.SelectorExpr { return f.PkgIdent("def", name) }
	// types
	//typeEx         = f.RefType(f.Type("ZendExecuteData"))
	//typeZval       = f.RefType(f.PkgIdent("types", "Zval"))
	typeEx         = zppPkgIdent("DefEx")
	typeZval       = zppPkgIdent("DefReturn")
	typeArgInfo    = defPkgIdent("ArgInfo")
	typeReturnInfo = defPkgIdent("ReturnInfo")

	// variables
	executeDataIdent = f.Ident("executeData")
	returnValueIdent = f.Ident("returnValue")
	fpIdent          = f.Ident("fp")
	retIdent         = f.Ident("ret")
	okIdent          = f.Ident("ok")
	// flags
	flagQuite   = zppPkgIdent("FlagQuiet")
	flagThrow   = zppPkgIdent("FlagThrow")
	flagOldMode = zppPkgIdent("FlagOldMode")
)

func genFileNode(name string, infos []*ZifInfo) *ast.File {
	fb := f.NewFileBuilder(name)
	fb.AddImport("sik/zend/def")
	fb.AddImport("sik/zend/zpp")

	for _, zifInfo := range infos {
		fb.AddDecl(
			f.ValueSpecDeclEx(
				f.DocComment("\n// generate by "+zifInfo.funcName),
				f.Ident(zifInfo.defName),
				&ast.CallExpr{
					Fun:  defPkgIdent("DefFunc"),
					Args: genDefFuncArgs(zifInfo, zifInfo.name),
				},
			),
		)
		for _, aliasName := range zifInfo.aliasNames {
			fb.AddDecl(
				f.ValueSpecDeclEx(
					f.DocComment("\n// generate by "+zifInfo.funcName),
					f.Ident("Zif"+strutil.UpperCamelCase(aliasName)),
					&ast.CallExpr{
						Fun:  defPkgIdent("DefFunc"),
						Args: genDefFuncArgs(zifInfo, aliasName),
					},
				),
			)
		}
	}

	return fb.Build()
}

func genDefFuncArgs(zifInfo *ZifInfo, phpFuncName string) []ast.Expr {
	//log.Printf("ZifInfo: %+v\n", *zifInfo)

	// 构建 ArgInfos
	var realArgInfos []ast.Expr
	for _, argInfo := range zifInfo.argInfos {
		switch argInfo.typ {
		case ZppTypeEx, ZppTypeRet, ZppTypeOpt:
			continue
		default:
			realArgInfos = append(realArgInfos, &ast.CompositeLit{
				Elts: []ast.Expr{
					f.KeyValue("Name", f.StrLit(argInfo.name)),
				},
			})
		}
	}

	return []ast.Expr{
		f.StrLit(phpFuncName),
		f.IntLit(zifInfo.minNumArgs),
		f.IntLit(zifInfo.maxNumArgs),
		&ast.CompositeLit{Type: f.ArrayType(typeArgInfo), Elts: realArgInfos},
		genZifHandler(zifInfo),
		//genDefFuncOpts(zifInfo),
	}
}

func genDefFuncOpts(zifInfo *ZifInfo) ast.Expr {
	// 构建 DefFuncOpts 字段
	var optElements []ast.Expr

	if len(optElements) == 0 {
		return f.NilLit()
	}

	// 构建结构体字面量
	return &ast.CompositeLit{
		Type: defPkgIdent("DefFuncOpts"),
		Elts: optElements,
	}
}

func genZifHandler(zifInfo *ZifInfo) ast.Expr {
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
			flags = append(flags, flagQuite)
		}
		if zifInfo.strict {
			flags = append(flags, flagThrow)
		}
		if zifInfo.oldMode {
			flags = append(flags, flagOldMode)
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
				stmts = append(stmts, f.ExprStmt(f.MethodCallExpr(fpIdent, "StartOptional", nil)))
			} else if parseMethod, args, ok := toZppParseMethodEx(argTyp); ok {
				stmts = append(stmts, f.AssignStmt(
					f.Ident(info.name),
					f.MethodCallExpr(fpIdent, parseMethod, args),
				))
			} else {
				log.Fatalf("Zpp类型未定义 Parse 方法: type=%s\n", argTyp)
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

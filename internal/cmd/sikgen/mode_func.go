package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	f "sik/internal/cmd/sikgen/astutil"
	"strconv"
	"strings"
)

func runGenFunc(dir string) {
	runClearFunc(dir)
	eachGoFile(dir, func(fileName string) {
		file, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		zifInfos := scanZifInFile(file)
		if len(zifInfos) > 0 {
			zifInfoFileName := fileName[:len(fileName)-3] + ".zif.go" // 将 .go 后缀改为 .zif.go
			zifInfoFileCode := genZifInfoFileContent(file.Name, zifInfos)
			err := ioutil.WriteFile(zifInfoFileName, []byte(zifInfoFileCode), 0644)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Generate File: " + zifInfoFileName)
		}
	})
}

func runClearFunc(dir string) {
	eachGoFile(dir, func(fileName string) {
		if strings.HasSuffix(fileName, ".zif.go") {
			err := os.Remove(fileName)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Remove File: " + fileName)
		}
	})
}

var dumpCfg = printer.Config{
	Mode:     printer.UseSpaces | printer.TabIndent,
	Tabwidth: 8,
}

func genZifInfoFileContent(name *ast.Ident, infos []*ZifInfo) string {
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
	file := &ast.File{Name: name, Decls: decls}

	return printNode(file)
}

func genDefFuncOpts(zifInfo *ZifInfo) ast.Expr {
	// 构建 DefFuncOpts 字段
	var optElements []ast.Expr
	optElements = append(optElements, f.KeyValue("name", f.StrLit(zifInfo.name)))
	if zifInfo.minNumArgs > 0 {
		optElements = append(optElements, f.KeyValue("minNumArgs", f.IntLit(zifInfo.minNumArgs)))
	}
	optElements = append(optElements, f.KeyValue("handler", zifInfo.handler))

	// 构建结构体字面量
	return &ast.CompositeLit{
		Type: f.Ident("DefFuncOpts"),
		Elts: optElements,
	}
}

func printNode(node ast.Node) string {
	var buf strings.Builder
	err := dumpCfg.Fprint(&buf, token.NewFileSet(), node)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func eachGoFile(dir string, handler func(fileName string)) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, "_") || strings.HasPrefix(name, ".") {
			continue
		}

		path := filepath.Join(dir, name)
		if file.IsDir() {
			eachGoFile(path, handler)
		} else if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
			handler(path)
		}
	}
}

type ArgInfo struct {
	name       string
	typ        ZvalType
	isVariadic bool
}

type ZifInfo struct {
	funcName      string
	defName       string
	name          string
	handler       ast.Expr
	minNumArgs    int
	maxNumArgs    int
	useArgNames   bool
	argNames      []string
	argInfos      []ArgInfo
	returnArgInfo *ArgInfo
	strict        bool
}

type ZvalType string

const (
	ZvalTypeInt    ZvalType = "int"
	ZvalTypeDouble ZvalType = "double"
	ZvalTypeString ZvalType = "string"
)

func scanZifInFile(file *ast.File) []*ZifInfo {
	var zifInfos []*ZifInfo
	ast.Inspect(file, func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcName := funcDecl.Name.Name
		if !strings.HasPrefix(funcName, "Zif") {
			return true
		}

		if zifInfo, ok := parseZifInfo(funcDecl); ok {
			zifInfos = append(zifInfos, zifInfo)
		}
		return true
	})
	return zifInfos
}

func parseZifInfo(funcDecl *ast.FuncDecl) (zifInfo *ZifInfo, ok bool) {
	funcName := funcDecl.Name.Name
	params := funcDecl.Type.Params.List
	returns := funcDecl.Type.Results

	// 从参数类型获取信息
	var argNames []string
	var argInfos []ArgInfo
	var returnArgInfo *ArgInfo
	for _, param := range params {
		paramName := param.Names[0].Name
		paramType, ok := asZvalType(param.Type)
		if !ok {
			log.Println("Zif函数未简化或错误，参数类型不合法: " + funcName)
			return nil, false
		}
		argNames = append(argNames, paramName)
		argInfos = append(argInfos, ArgInfo{
			name: paramName,
			typ:  paramType,
		})
	}

	// 从返回类型获取信息
	if returns != nil && len(returns.List) == 1 {
		returnType, ok := asZvalType(returns.List[0].Type)
		if !ok {
			log.Println("Zif函数返回值类型不合法: " + funcName)
			return nil, false
		}
		returnArgInfo = &ArgInfo{
			typ: returnType,
		}
	}

	// 从注解获取信息
	annoArgs := getAnnoArgs(funcDecl.Doc)

	// 构建 zif 信息
	zifName := annoArgs.name
	if zifName == "" {
		zifName = lcName(funcName[len("Zif"):])
	}

	zifInfo = &ZifInfo{
		funcName:      funcName,
		defName:       funcName + "Def",
		name:          zifName,
		minNumArgs:    annoArgs.minNumArgs,
		maxNumArgs:    annoArgs.maxNumArgs,
		useArgNames:   true,
		argNames:      argNames,
		argInfos:      argInfos,
		returnArgInfo: returnArgInfo,
		strict:        annoArgs.strict,
	}

	// 构建 handler
	zifInfo.handler = parseZifHandler(zifInfo, funcDecl)

	return zifInfo, true
}

const zifAnnoName = "@zif"

type zifAnnoFlags struct {
	name       string
	strNumArgs string
	minNumArgs int
	maxNumArgs int
	strict     bool
}

func getAnnoArgs(doc *ast.CommentGroup) zifAnnoFlags {
	annoFlags := zifAnnoFlags{minNumArgs: -1, maxNumArgs: -1}

	// 从注释中找到注解文本
	var args []string
	if doc != nil {
		for _, comment := range doc.List {
			commentText := strings.Trim(comment.Text, "/\t ")
			if strings.HasPrefix(commentText, zifAnnoName) {
				args = strings.Split(commentText, " ")
				break
			}
		}
	}
	if len(args) <= 1 {
		return annoFlags
	}

	flagSet := flag.NewFlagSet(zifAnnoName, flag.ContinueOnError)
	flagSet.StringVar(&annoFlags.name, "n", "", "name")
	flagSet.StringVar(&annoFlags.strNumArgs, "c", "", "num of args")
	flagSet.BoolVar(&annoFlags.strict, "s", false, "open strict mode")
	err := flagSet.Parse(args[1:])
	if err != nil {
		return annoFlags
	}

	/**
	 * 解析 -c 参数，支持多种写法
	 * -c 0    // 最小、最大参数个数相同
	 * -c a,b  // 分别设置最小、最大参数个数
	 * -c a,   // 只设置最小参数个数
	 * -c ,b   // 只设置最大参数个数
	 */
	if len(annoFlags.strNumArgs) != 0 {
		strNumArgs := annoFlags.strNumArgs
		pos := strings.IndexByte(strNumArgs, ',')
		if pos < 0 {
			numArgs := parseFlagInt(strNumArgs)
			annoFlags.minNumArgs = numArgs
			annoFlags.maxNumArgs = numArgs
		} else {
			minNumArgsStr := strings.TrimSpace(strNumArgs[:pos])
			maxNumArgsStr := strings.TrimSpace(strNumArgs[pos+1:])
			if minNumArgsStr != "" {
				annoFlags.minNumArgs = parseFlagInt(minNumArgsStr)
				annoFlags.maxNumArgs = parseFlagInt(maxNumArgsStr)
			}
		}
	}

	return annoFlags
}

func parseFlagInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("解析注解失败: %s 不是合法数字\n", s)
	}
	return val
}

func parseZifHandler(zifInfo *ZifInfo, decl *ast.FuncDecl) ast.Expr {
	//
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
	if zifInfo.returnArgInfo == nil {
		stmts = []ast.Stmt{
			f.ExprStmt(f.CallExpr(zifInfo.funcName, args)),
		}
	} else {
		var setter string
		switch zifInfo.returnArgInfo.typ {
		case ZvalTypeString:
			setter = "SetRawString"
		case ZvalTypeInt:
			setter = "SetLong"
		default:
			log.Fatalln("此 ZvalType 未设置对应 Setter" + zifInfo.returnArgInfo.typ)
		}

		stmts = []ast.Stmt{
			f.AssignStmt(retIdent, f.CallExpr(zifInfo.funcName, args)),
			f.ExprStmt(f.MethodCallExpr(returnValueIdent, setter, []ast.Expr{retIdent})),
		}
	}

	return &ast.FuncLit{
		Type: funcType,
		Body: &ast.BlockStmt{List: stmts},
	}
}

func asZvalType(typ ast.Expr) (ZvalType, bool) {
	typName := printNode(typ)
	switch typName {
	case "int":
		return ZvalTypeInt, true
	case "float64":
		return ZvalTypeDouble, true
	case "string":
		return ZvalTypeString, true
	default:
		return "", false
	}
}

func lcName(name string) string {
	var buf strings.Builder
	for i, c := range name {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf.WriteByte('_')
			}
			buf.WriteRune(c - 'A' + 'a')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

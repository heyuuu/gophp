package zif

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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

		// 跳过没修改的函数
		params := funcDecl.Type.Params.List
		if len(params) > 0 {
			firstParamTypeDesc := printNode(params[0].Type)
			if firstParamTypeDesc == "*ZendExecuteData" || firstParamTypeDesc == "*zend.ZendExecuteData" {
				return true
			}
		}

		// 解析 Zif 信息
		if zifInfo, ok := parseZifInfo(funcDecl); ok {
			zifInfos = append(zifInfos, zifInfo)
		}
		return true
	})
	return zifInfos
}

func parseZifInfo(funcDecl *ast.FuncDecl) (*ZifInfo, bool) {
	funcName := funcDecl.Name.Name

	// 从注解获取信息
	annoArgs, err := getAnnoArgs(funcDecl.Doc)
	if err != nil {
		log.Fatalf("解析 ZifInfo 失败: funcName=%s, error=%s\n", funcName, err.Error())
	}

	// 构建 zif 信息
	zifName := annoArgs.name
	if zifName == "" {
		zifName = lcName(funcName[len("Zif"):])
	}

	zifInfo := &ZifInfo{
		funcName:   funcName,
		defName:    "Def" + funcName,
		name:       zifName,
		aliasNames: annoArgs.aliasNames,
		strict:     annoArgs.strict,
		oldMode:    annoArgs.oldMode,
	}

	// 从参数类型获取信息
	argInfos, err := parseArgInfos(funcDecl)
	if err != nil {
		log.Fatalf("Zif函数 %s 定义错误: %s", funcName, err.Error())
	}
	zifInfo.argInfos = argInfos
	zifInfo.minNumArgs, zifInfo.maxNumArgs = calcNumArgs(argInfos)

	// 从返回类型获取信息
	returnInfo, err := parseReturnInfo(funcDecl)
	if err != nil {
		log.Fatalf("Zif函数 %s 定义错误: %s", funcName, err.Error())
	}
	zifInfo.returnInfo = returnInfo

	return zifInfo, true
}

func getAnnoArgs(doc *ast.CommentGroup) (zifAnnoFlags, error) {
	annoFlags := zifAnnoFlags{maxNumArgs: -1}

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
		return annoFlags, nil
	}

	var strAlias, strNumArgs string
	flagSet := flag.NewFlagSet(zifAnnoName, flag.ContinueOnError)
	flagSet.StringVar(&annoFlags.name, "n", "", "name")
	flagSet.StringVar(&strAlias, "alias", "", "alias name")
	flagSet.StringVar(&strNumArgs, "c", "", "num of args")
	flagSet.BoolVar(&annoFlags.strict, "s", false, "use strict mode")
	flagSet.BoolVar(&annoFlags.quiet, "q", false, "use quite mode")
	flagSet.BoolVar(&annoFlags.oldMode, "old", false, "use old mode")
	err := flagSet.Parse(args[1:])
	if err != nil {
		return annoFlags, err
	}

	// 解析 -alias 参数
	if strAlias != "" {
		annoFlags.aliasNames = strings.Split(strAlias, ",")
		for _, name := range annoFlags.aliasNames {
			if !isValidName(name) {
				return annoFlags, errors.New("@zif 的 --alias 参数错误: 参数值不是合法函数名")
			}
		}
	}

	/**
	 * 解析 -c 参数，支持多种写法
	 * -c m    // 最小、最大参数个数相同
	 * -c m,n  // 分别设置最小、最大参数个数
	 * -c m,   // 只设置最小参数个数
	 */
	if strNumArgs != "" {
		errMsg := "解析注解失败，-c 参数值不合法: " + strNumArgs
		pos := strings.IndexByte(strNumArgs, ',')
		if pos < 0 {
			numArgs := parseFlagInt(strNumArgs, errMsg)
			annoFlags.minNumArgs = numArgs
			annoFlags.maxNumArgs = numArgs
		} else {
			minNumArgsStr := strings.TrimSpace(strNumArgs[:pos])
			maxNumArgsStr := strings.TrimSpace(strNumArgs[pos+1:])
			annoFlags.minNumArgs = parseFlagInt(minNumArgsStr, errMsg)
			if maxNumArgsStr != "" {
				annoFlags.maxNumArgs = parseFlagInt(maxNumArgsStr, errMsg)
			}
		}
	}

	return annoFlags, nil
}

func isValidName(str string) bool {
	for i, c := range str {
		valid := ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_' || (i != 0 && '0' <= c && c <= 'Z')
		if !valid {
			return false
		}
	}
	return true
}

func parseArgInfos(funcDecl *ast.FuncDecl) ([]ArgInfo, error) {
	params := funcDecl.Type.Params.List

	var argInfos []ArgInfo
	hasRealParam := false
	hasVarargs := false
	hasOpt := false
	for i, param := range params {
		paramName := param.Names[0].Name
		paramTypeDesc := printNode(param.Type)
		paramType, ok := toZppType(paramTypeDesc)
		if !ok {
			return nil, errors.New("参数类型不合法:" + paramTypeDesc)
		}
		switch paramType {
		case ZppTypeEx, ZppTypeRet:
			if hasRealParam {
				return nil, errors.New(fmt.Sprintf("参数类型不合法, 特殊类型 %s 必须在所有实际参数前", paramTypeDesc))
			}
		case ZppTypeVariadic:
			if hasVarargs {
				return nil, errors.New("参数类型不合法, 不可有多个变长参数")
			}
			if i != len(params)-1 {
				return nil, errors.New("参数类型不合法, 变长参数必须是最后一个参数")
			}
			hasRealParam = true
			hasVarargs = true
		case ZppTypeOpt:
			if hasOpt {
				return nil, errors.New("参数类型不合法, 不可有多个Opt")
			}
			hasRealParam = true
			hasOpt = true
		default:
			hasRealParam = true
		}
		argInfos = append(argInfos, ArgInfo{
			name: lcName(paramName),
			typ:  paramType,
		})
	}
	return argInfos, nil
}

func parseReturnInfo(funcDecl *ast.FuncDecl) (*ReturnInfo, error) {
	returns := funcDecl.Type.Results
	if returns == nil {
		return nil, nil
	}

	var retTypes []ZppType
	for _, result := range returns.List {
		returnTypeSpec := printNode(result.Type)
		returnType, ok := toZppType(returnTypeSpec)
		if !ok {
			return nil, errors.New("返回值类型不合法: " + returnTypeSpec)
		}
		retTypes = append(retTypes, returnType)
	}

	switch len(retTypes) {
	case 0:
		return nil, nil
	case 1:
		return &ReturnInfo{typ: retTypes[0]}, nil
	case 2:
		if retTypes[1] == ZppTypeBool {
			return &ReturnInfo{typ: retTypes[0], withOk: true}, nil
		}
	}

	return nil, errors.New("不支持此返回值类型组合: " + printNode(funcDecl.Type))
}

func calcNumArgs(argInfos []ArgInfo) (minNumArgs int, maxNumArgs int) {
	minNumArgs, maxNumArgs = -1, 0
outer:
	for _, info := range argInfos {
		switch info.typ {
		case ZppTypeEx, ZppTypeRet:
			// skip
		case ZppTypeOpt:
			minNumArgs = maxNumArgs
		case ZppTypeVariadic:
			maxNumArgs = -1
			break outer
		default:
			maxNumArgs++
		}
	}
	if minNumArgs < 0 {
		minNumArgs = maxNumArgs
	}
	return minNumArgs, maxNumArgs
}

func parseFlagInt(s string, errMsg string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf(errMsg)
	}
	return val
}

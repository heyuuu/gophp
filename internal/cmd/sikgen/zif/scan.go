package zif

import (
	"flag"
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

	// 从注解获取信息
	annoArgs := getAnnoArgs(funcDecl.Doc)

	// 构建 zif 信息
	zifName := annoArgs.name
	if zifName == "" {
		zifName = lcName(funcName[len("Zif"):])
	}

	zifInfo = &ZifInfo{
		funcName:   funcName,
		defName:    "Def" + funcName,
		name:       zifName,
		minNumArgs: annoArgs.minNumArgs,
		maxNumArgs: annoArgs.maxNumArgs,
		strict:     annoArgs.strict,
	}

	// 从参数类型获取信息
	for _, param := range params {
		paramName := param.Names[0].Name
		paramType, ok := toZppType(param.Type)
		if !ok {
			typeDesc := printNode(param.Type)
			if typeDesc == "*ZendExecuteData" || typeDesc == "*zend.ZendExecuteData" {
				//log.Println("Zif函数未简化: " + funcName)
			} else {
				log.Fatalf("Zif函数错误，参数类型不合法: func=%s, type=%s\n", funcName, typeDesc)
			}
			return nil, false
		}
		switch paramType {
		case ZppTypeEx:
			if len(zifInfo.argInfos) > 0 {
				typeDesc := printNode(param.Type)
				log.Fatalf("Zif函数错误，参数类型不合法, DefEx 必须在所有实际参数前: func=%s, type=%s\n", funcName, typeDesc)
			}
			zifInfo.argNeedEx = true
		case ZppTypeRet:
			if len(zifInfo.argInfos) > 0 {
				typeDesc := printNode(param.Type)
				log.Fatalf("Zif函数错误，参数类型不合法, DefRet 必须在所有实际参数前: func=%s, type=%s\n", funcName, typeDesc)
			}
			zifInfo.argNeedRet = true
		default:
			zifInfo.argInfos = append(zifInfo.argInfos, ArgInfo{
				name: paramName,
				typ:  paramType,
			})
		}
	}

	// 从返回类型获取信息
	if returns != nil && len(returns.List) == 1 {
		returnType, ok := toZppType(returns.List[0].Type)
		if !ok {
			typeDesc := returns.List[0].Type
			log.Fatalf("Zif函数错误，返回值类型不合法: func=%s, type=%s\n", funcName, typeDesc)
			return nil, false
		}
		zifInfo.returnArgInfo = &ArgInfo{
			typ: returnType,
		}
	}

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
	 * -c m    // 最小、最大参数个数相同
	 * -c m,n  // 分别设置最小、最大参数个数
	 * -c m,   // 只设置最小参数个数
	 */
	if len(annoFlags.strNumArgs) != 0 {
		strNumArgs := annoFlags.strNumArgs
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

	return annoFlags
}

func parseFlagInt(s string, errMsg string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf(errMsg)
	}
	return val
}

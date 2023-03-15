package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	f "sik/internal/cmd/sikgen/astutil"
	"strings"
)

func runGenFunc(dir string) {
	eachGoFile(dir, func(fileName string) {
		file, err := parser.ParseFile(token.NewFileSet(), fileName, nil, 0)
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
		}
	})
}

func genZifInfoFileContent(name *ast.Ident, infos []ZifInfo) string {
	var decls []ast.Decl
	for _, zifInfo := range infos {
		varExpr := &ast.Ident{Name: zifInfo.defName}
		valueExpr := &ast.CompositeLit{
			Type: f.Ident("FuncDefine"),
			Elts: []ast.Expr{
				f.KeyValue("funcName", f.StrLit(zifInfo.name)),
				f.KeyValue("requiredNumArgs", f.IntLit(zifInfo.requiredNumArgs)),
				f.KeyValue("handler", f.NilLit()),
			},
		}

		decls = append(decls, &ast.GenDecl{
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

	var buf strings.Builder
	dumpConf := printer.Config{}
	err := dumpConf.Fprint(&buf, token.NewFileSet(), file)
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
	typ        string
	isVariadic bool
}

type ZifInfo struct {
	varName         string
	defName         string
	name            string
	requiredNumArgs int
	argInfos        []ArgInfo
	returnArgInfo   *ArgInfo
}

func scanZifInFile(f *ast.File) []ZifInfo {
	var zifInfos []ZifInfo
	ast.Inspect(f, func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcName := funcDecl.Name.Name
		if !strings.HasPrefix(funcName, "Zif") {
			return true
		}

		zifInfo := ZifInfo{
			varName: funcName,
			defName: funcName + "Def",
		}

		zifInfos = append(zifInfos, zifInfo)

		return true
	})
	return zifInfos
}

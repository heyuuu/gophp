package zif

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func RunGenFunc(dir string) {
	RunClearFunc(dir)
	eachGoFile(dir, func(fileName string) {
		file, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		zifInfos := scanZifInFile(file)
		if len(zifInfos) > 0 {
			zifInfoFileName := fileName[:len(fileName)-3] + ".zif.go" // 将 .go 后缀改为 .zif.go
			zifInfoFileCode := printNode(genFileNode(file.Name.Name, zifInfos))
			err := ioutil.WriteFile(zifInfoFileName, []byte(zifInfoFileCode), 0644)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Generate File: " + zifInfoFileName)
		}
	})
}

func RunClearFunc(dir string) {
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

func printNode(node ast.Node) string {
	var buf strings.Builder
	err := dumpCfg.Fprint(&buf, token.NewFileSet(), node)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

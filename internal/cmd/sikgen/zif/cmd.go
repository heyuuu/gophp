package zif

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
)

func RunGenFunc(dir string) {
	handledFiles := make(map[string]bool)
	unchanged, updated, deleted := 0, 0, 0

	eachGoFile(dir, func(fileName string) {
		file, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		zifInfos := scanZifInFile(file)
		if len(zifInfos) > 0 {
			zifInfoFileName := fileName[:len(fileName)-3] + ".zif.go" // 将 .go 后缀改为 .zif.go
			zifInfoFileCode := printNode(genFileNode(file.Name.Name, zifInfos))

			isChanged, err := writeFileIfChanged(zifInfoFileName, zifInfoFileCode)
			if err != nil {
				log.Fatal(err)
			}

			handledFiles[zifInfoFileName] = true
			if isChanged {
				updated++
				log.Println("Update file: " + zifInfoFileName)
			} else {
				unchanged++
				//log.Println("Unchanged file: " + zifInfoFileName)
			}
		}
	})
	eachGoFile(dir, func(fileName string) {
		if !strings.HasSuffix(fileName, ".zif.go") {
			return
		}
		if _, ok := handledFiles[fileName]; ok {
			return
		}

		err := os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}

		deleted++
		log.Println("Remove File: " + fileName)
	})
	log.Printf("处理完成. 共有更新文件 %d, 未变更文件 %d, 移除文件 %d\n", updated, unchanged, deleted)
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

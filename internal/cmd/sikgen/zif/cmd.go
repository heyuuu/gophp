package zif

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	ModeEachFile = "file"
	ModeEachPkg  = "pkg"
)

func RunGenFunc(dir string, mode string) {
	handledFiles := make(map[string]bool)
	unchanged, updated, deleted := 0, 0, 0

	fmt.Println(mode)

	eachGenZifFile(dir, mode == ModeEachPkg, func(zifFile string, code string) {
		isChanged, err := writeFileIfChanged(zifFile, code)
		if err != nil {
			log.Fatal(err)
		}

		handledFiles[zifFile] = true
		if isChanged {
			updated++
			log.Println("Update file: " + zifFile)
		} else {
			unchanged++
			//log.Println("Unchanged file: " + zifInfoFileName)
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

func eachGenZifFile(dir string, isPkgMode bool, handler func(zifFile string, code string)) {
	if !isPkgMode {
		eachGoFile(dir, func(fileName string) {
			pkgName, zifInfos, err := scanZifInFile(fileName)
			if err != nil {
				log.Fatal(err)
			}
			if len(zifInfos) == 0 {
				return
			}

			zifFile := fileName[:len(fileName)-3] + ".zif.go" // 将 .go 后缀改为 .zif.go
			zifCode := printNode(genFileNode(pkgName, zifInfos))
			handler(zifFile, zifCode)
		})
	} else {
		eachGoDir(dir, func(dirPath string, filePaths []string) {
			var pkgName string
			var zifInfos []*ZifInfo

			sort.Strings(filePaths)
			for _, filePath := range filePaths {
				filePkgName, fileZifInfos, err := scanZifInFile(filePath)
				if err != nil {
					log.Fatal(err)
				}
				pkgName = filePkgName
				zifInfos = append(zifInfos, fileZifInfos...)
			}
			if len(zifInfos) == 0 {
				return
			}

			zifFile := filepath.Join(dirPath, pkgName+".zif.go")
			zifCode := printNode(genFileNode(pkgName, zifInfos))
			handler(zifFile, zifCode)
		})
	}
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

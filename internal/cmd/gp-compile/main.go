package main

import (
	"flag"
	"fmt"
	"github.com/heyuuu/gophp/compile/ir"
	irPrinter "github.com/heyuuu/gophp/compile/ir/printer"
	"github.com/heyuuu/gophp/php/parser"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// parse args
	var srcPath string
	var distPath string
	flag.StringVar(&srcPath, "src", "", "php src path")
	flag.StringVar(&distPath, "dist", "", "go dist path")
	flag.Parse()

	// check args
	srcPath = strings.TrimSpace(srcPath)
	if srcPath == "" {
		log.Fatal("src path is not set")
	}
	srcStat, srcErr := os.Stat(srcPath)
	if srcErr != nil {
		log.Fatalf("src path is empty or not exists: src=%s, error=%s", srcPath, srcErr.Error())
	} else if !srcStat.IsDir() {
		log.Fatalf("src path must be a dir")
	}

	distPath = strings.TrimSpace(distPath)
	if distPath == "" {
		log.Fatal("dist path is not set")
	}
	distStat, srcErr := os.Stat(distPath)
	if srcErr != nil && !distStat.IsDir() {
		log.Fatalf("dist path must not exist or be a dir")
	}

	// compile
	err := simpleCompileDir(srcPath, distPath)
	if err != nil {
		log.Fatalf("compile failed: %v", err)
	}
}

func simpleCompileDir(srcPath string, distPath string) error {
	return filepath.Walk(srcPath, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".php") {
			return err
		}

		relativePath := path[len(srcPath):]
		distFile := filepath.Join(distPath, strings.ReplaceAll(relativePath, ".php", ".go"))
		fmt.Printf("relative=%s, distFile=%s\n", relativePath, distPath)
		return simpleCompileFile(path, distFile)
	})
}

func simpleCompileFile(srcFile string, distFile string) error {
	// parse + compile
	astFile, err := parser.ParseFile(srcFile)
	if err != nil {
		return err
	}

	irFile, err := ir.ParseAstFile(astFile)
	if err != nil {
		return err
	}

	// render
	irCode, err := irPrinter.SprintFile(irFile)
	if err != nil {
		return err
	}

	// write file
	return safeWriteFile(distFile, irCode)
}

func safeWriteFile(file string, content string) (err error) {
	dir := filepath.Dir(file)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	if err != nil {
		return err
	}

	return os.WriteFile(file, []byte(content), 0644)
}

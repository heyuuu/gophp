package main

import (
	"flag"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/utils/finder"
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
	distStat, distErr := os.Stat(distPath)
	if distErr == nil && !distStat.IsDir() {
		log.Fatalf("dist path must not exist or be a dir")
	}

	// compile
	err := run(srcPath, distPath)
	if err != nil {
		log.Fatalf("compile failed: %v", err)
	}
}

func run(srcPath string, distPath string) error {
	proj, err := compileDir(srcPath)
	if err != nil {
		return err
	}

	return printIrProject(distPath, proj)
}

func compileDir(srcPath string) (*ir.Project, error) {
	proj := ir.NewProject()

	f := finder.NewFinder(srcPath).Files()
	err := f.Walk(func(f finder.File) error {
		if !strings.HasSuffix(f.Path, ".php") {
			return nil
		}

		irFile, err := compileFile(f.Path)
		if err != nil {
			return err
		}

		return proj.AddFile(f.RelativePath, irFile)
	})
	if err != nil {
		return nil, err
	}

	return proj, nil
}

func compileFile(srcFile string) (*ir.File, error) {
	// parse + compile
	astFile, err := parser.ParseFile(srcFile)
	if err != nil {
		return nil, err
	}

	irFile, err := ir.ParseAstFile(astFile)
	if err != nil {
		return nil, err
	}

	return irFile, nil
}

func printIrProject(distPath string, proj *ir.Project) error {
	codes, err := ir.PrintProject(proj)
	if err != nil {
		return err
	}

	for name, code := range codes {
		// todo 简易文件名规则，后续待优化
		var filename string
		if name == "" {
			filename = "_.go"
		} else {
			filename = strings.ReplaceAll(name, "\\", "__") + ".go"
		}

		distFile := filepath.Join(distPath, filename)
		err := safeWriteFile(distFile, code)
		if err != nil {
			return err
		}
	}

	return nil
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

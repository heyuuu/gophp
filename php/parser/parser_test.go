package parser

import (
	"github.com/heyuuu/gophp/utils/vardumper"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var projRoot = "../../"

func init() {
	wd, _ := os.Getwd()
	SetProjRoot(filepath.Clean(filepath.Join(wd, projRoot)))
}

func TestParseCode(t *testing.T) {
	code := "<?php var_dump(1);"
	nodes, err := ParseCode(code)
	if err != nil {
		log.Fatal(err)
	}

	vardumper.Print(nodes)
}

func TestParseFile(t *testing.T) {
	file := filepath.Join(projRoot, "tools/parser/parser.php")
	nodes, err := ParseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	vardumper.Print(nodes)
}

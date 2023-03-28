package phpparse

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"testing"
)

var projRoot = "../../../../"
var script = filepath.Join(projRoot, "tools/parser/parser.php")
var output = filepath.Join(projRoot, "runtime")
var srcFile = filepath.Join(projRoot, "tools/parser/parser.php")
var astFile = filepath.Join(projRoot, "runtime/parser.php.json")

func Test_parseToJsonFile(t *testing.T) {
	command := exec.Command("pwd")
	result, err := command.CombinedOutput()
	fmt.Println(string(result))

	err = parseToJsonFile(script, srcFile, output)
	if err != nil {
		log.Fatal(err)
	}
}

func Test_loadAstJson(t *testing.T) {
	data, err := loadAstJson(astFile)
	if err != nil {
		log.Fatal(err)
	}
	//_ = data
	fmt.Println(data)
}

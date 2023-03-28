package phpparse

import (
	"errors"
	"fmt"
	"gophp/php/ast"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func ParseCode(code string) ([]ast.Stmt, error) {
	json, err := runParser("-c", code)
	if err != nil {
		return nil, err
	}
	return decodeAstData(json)
}

func ParseFile(file string) ([]ast.Stmt, error) {
	json, err := runParser("-f", file)
	if err != nil {
		return nil, err
	}
	return decodeAstData(json)
}

func parseToJsonFile(script string, filename string, output string) error {
	result, err := runParser("-s", filename, "-o", output)
	fmt.Println(result)
	if err != nil {
		fmt.Println("命令执行失败: " + err.Error())
		return err
	}

	return nil
}

func loadAstJson(jsonFile string) ([]ast.Stmt, error) {
	binData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	return decodeAstData(binData)
}

/* Parser 脚本相关 */
var scriptPath = "./tools/parser/parser.php"

func SetScriptPath(newScriptPath string) {
	scriptPath = newScriptPath
}

func runParser(args ...string) ([]byte, error) {
	if len(scriptPath) == 0 {
		return nil, errors.New("PHP Parser 脚本路径未设置")
	}
	if _, err := os.Stat(scriptPath); err != nil {
		return nil, errors.New("PHP Parser 脚本路径不存在或不可读")
	}

	var commandArgs = make([]string, len(args)+1)
	commandArgs[0] = scriptPath
	copy(commandArgs[1:], args)

	command := exec.Command("php", commandArgs...)
	log.Printf("Run command: %s\n", command.String())
	return command.CombinedOutput()
}

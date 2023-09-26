package phpparse

import (
	"errors"
	"github.com/heyuuu/gophp/compile/ast"
	"log"
	"os"
	"os/exec"
)

func ParseCode(code string) ([]ast.Stmt, error) {
	output, err := runParser("-c", code)
	if err != nil {
		return nil, err
	}
	return decodeOutput(output)
}

func ParseFile(file string) ([]ast.Stmt, error) {
	output, err := runParser("-f", file)
	if err != nil {
		return nil, err
	}
	return decodeOutput(output)
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

	commandArgs := append([]string{scriptPath}, args...)
	command := exec.Command("php", commandArgs...)
	log.Printf("Run command: %s\n", command.String())
	return command.CombinedOutput()
}

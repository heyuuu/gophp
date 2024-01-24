package phpparse

import (
	"github.com/heyuuu/gophp/compile/ast"
	"log"
	"os"
	"os/exec"
)

func ParseCode(code string) (*ast.File, error) {
	output, err := runParser("-c", code)
	if err != nil {
		return nil, err
	}
	return decodeOutput(output)
}

/* Parser 脚本相关 */
var scriptPath = "/Users/heyu/Code/sik/gophp/tools/parser/parser.php"

func runParser(args ...string) ([]byte, error) {
	if _, err := os.Stat(scriptPath); err != nil {
		log.Panicln("PHP Parser 脚本路径不存在或不可读")
	}

	commandArgs := append([]string{scriptPath}, args...)
	command := exec.Command("php", commandArgs...)
	log.Printf("Run command: %s\n", command.String())
	return command.CombinedOutput()
}

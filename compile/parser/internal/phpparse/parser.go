package phpparse

import (
	"github.com/heyuuu/gophp/compile/ast"
	"log"
	"os"
	"os/exec"
)

func ParseCode(code string) (*ast.File, error) {
	output, err := runParser("parse", "-c", code)
	if err != nil {
		return nil, err
	}
	return decodeOutput(output)
}

/* Parser 脚本相关 */
var parser = "/usr/local/bin/gophp-parser"

func runParser(args ...string) ([]byte, error) {
	if _, err := os.Stat(parser); err != nil {
		log.Panicln("PHP Parser 脚本路径不存在或不可读")
	}

	command := exec.Command(parser, args...)
	log.Printf("Run command: %s\n", command.String())
	return command.CombinedOutput()
}

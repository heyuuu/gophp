package phpparse

import (
	"encoding/json"
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"log"
	"os"
	"os/exec"
)

func ParseCode(code string) (*ast.File, error) {
	_, astFile, err := ParseCodeVerbose(code)
	return astFile, err
}

func ParseCodeVerbose(code string) (raw string, astFile *ast.File, err error) {
	raw, err = parseCodeRaw(code)
	if err == nil {
		astFile, err = decodeAstData(raw)
	}
	return
}

func parseCodeRaw(code string) (string, error) {
	output, err := runParser("parse", "-c", code)
	if err != nil {
		return "", fmt.Errorf("php parse run failed: %w", err)
	}

	var res result
	if err = json.Unmarshal(output, &res); err != nil {
		return "", fmt.Errorf("php parse json Unmarshal failed: %s", res.Error)
	}

	if !res.Ok {
		return "", fmt.Errorf("php parse error: %s", res.Error)
	}

	return res.Data, nil
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

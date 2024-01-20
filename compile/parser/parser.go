package parser

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser/internal/phpparse"
	"os"
)

func ParseCode(code string) (*ast.File, error) {
	return ParseCodeEx(code, false)
}

func ParseCodeEx(code string, skipShebang bool) (*ast.File, error) {
	if skipShebang {
		code = "<?php " + code
	}
	return phpparse.ParseCode(code)
}

func ParseFile(file string) (*ast.File, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return phpparse.ParseCode(string(bytes))
}

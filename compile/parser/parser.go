package parser

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser/internal/phpparse"
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

func ParseCodeVerbose(code string) (string, *ast.File, error) {
	return phpparse.ParseCodeVerbose(code)
}

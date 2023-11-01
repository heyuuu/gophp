package php

import (
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/types"
)

func CompileFile(ctx *Context, fileHandle *FileHandle, skipShebang bool) (*types.Function, error) {
	code, err := fileHandle.ReadAll()
	if err != nil {
		return nil, err
	}

	astFile, err := parser.ParseCodeEx(code, skipShebang)
	if err != nil {
		return nil, err
	}

	return types.NewAstTopFunction(astFile), nil
}

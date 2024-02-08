package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/types"
)

func UninitializedZval() types.Zval { return types.Null }

func ExecuteScript(ctx *Context, fileHandle *FileHandle, skipShebang bool) (retval Val, err error) {
	assert.Assert(fileHandle != nil)
	topFunc, err := CompileFile(ctx, fileHandle, skipShebang)
	if err != nil {
		return types.Undef, err
	}
	if topFunc == nil {
		return types.Undef, errors.New("compile code failed")
	}

	executor := ctx.executor
	return executor.Execute(topFunc)
}

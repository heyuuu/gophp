package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func UninitializedZval() types.Zval { return types.Null }

func ExecuteScript(ctx *Context, fileHandle *FileHandle, skipShebang bool) (retval types.Zval, err error) {
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

func Exit(ctx *Context) {
	panic(perr.ErrExit)
}

func ExitWithCode(ctx *Context, code int) {
	// todo handle code
	Exit(ctx)
}

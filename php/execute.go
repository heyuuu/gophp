package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func ExecuteScript(ctx *Context, fileHandle *FileHandle, skipShebang bool) (retval Val, err error) {
	perr.Assert(fileHandle != nil)
	topFunc, err := CompileFile(ctx, fileHandle, skipShebang)
	if err != nil {
		return types.Undef, err
	}
	if topFunc == nil {
		return types.Undef, errors.New("compile code failed")
	}

	executor := NewExecutor(ctx)
	return executor.Execute(topFunc)
}

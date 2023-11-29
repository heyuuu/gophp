package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/lang"
)

func ExecuteScript(ctx *Context, fileHandle *FileHandle, skipShebang bool) (retval Val, err error) {
	lang.Assert(fileHandle != nil)
	topFunc, err := CompileFile(ctx, fileHandle, skipShebang)
	if err != nil {
		return nil, err
	}
	if topFunc == nil {
		return nil, errors.New("compile code failed")
	}

	executor := NewExecutor(ctx)
	return executor.Execute(topFunc)
}

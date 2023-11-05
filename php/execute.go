package php

import (
	"errors"
)

func ExecuteScript(ctx *Context, fileHandle *FileHandle, skipShebang bool) (retval Val, err error) {
	Assert(fileHandle != nil)
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

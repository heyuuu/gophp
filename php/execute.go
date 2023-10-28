package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/types"
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

	return Execute(ctx, topFunc)
}

func Execute(ctx *Context, fun *types.Function) (Val, error) {
	astFile := fun.GetUserAstFile()
	if astFile == nil {
		return nil, errors.New("获取 astFile 失败")
	}

	executor := Default()
	return executor.executeAstFile(astFile)
}

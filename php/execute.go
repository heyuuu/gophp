package php

import (
	"errors"
	"fmt"
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

func Execute(ctx *Context, fun types.Function) (Val, error) {
	// todo push stack
	executeData := NewExecuteData()

	switch f := fun.(type) {
	case *types.AstFunction:
		return ExecuteAstFunction(ctx, executeData, f)
	default:
		panic(fmt.Sprintf("unsupported function type: %T", f))
	}
}

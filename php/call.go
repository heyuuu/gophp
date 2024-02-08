package php

import "github.com/heyuuu/gophp/php/types"

func CallMethod(ctx *Context, object *types.Object, ce *types.Class, fn *types.Function, args []types.Zval) types.Zval {
	// todo
	return ctx.executor.doCall(fn, args, ce)
}

func CallFunction(ctx *Context, fn *types.Function, args []types.Zval, scope any) types.Zval {
	executor := ctx.Executor()
	return executor.doCall(fn, args, scope)
}

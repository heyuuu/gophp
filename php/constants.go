package php

import "github.com/heyuuu/gophp/php/types"

func RegisterConstantEx(ctx *Context, moduleNumber int, name string, value Val, flags uint32) {
	c := types.NewConstant(name, value, flags)
	ctx.EG().ConstantTable().Set(name, c)
}

func RegisterConstant(ctx *Context, moduleNumber int, name string, value Val) {
	RegisterConstantEx(ctx, moduleNumber, name, value, 0)
}

func GetConstant(ctx *Context, name string) *types.Constant {
	name = CleanNsName(name)
	return ctx.EG().ConstantTable().Get(name)
}

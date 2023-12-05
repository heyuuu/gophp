package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func RegisterConstantEx(ctx *Context, moduleNumber int, name string, value Val, flags types.ConstFlag) {
	c := types.NewConstant(name, value, flags, moduleNumber)
	innerRegisterConstant(ctx, c)
}

func RegisterConstant(ctx *Context, moduleNumber int, name string, value Val) {
	RegisterConstantEx(ctx, moduleNumber, name, value, 0)
}

func GetConstant(ctx *Context, name string) *types.Constant {
	name = CleanNsName(name)
	c := ctx.EG().ConstantTable().Get(name)
	if c == nil {
		c = ctx.EG().ConstantTable().Get(ascii.StrToLower(name))
		if c != nil && c.IsCaseSensitive() {
			c = nil
		}
	}

	return c
}

func innerRegisterConstant(ctx *Context, c *types.Constant) {
	name := c.Name()
	if !c.IsCaseSensitive() {
		name = strings.ToLower(name)
	} else {
		// 带命名空间的常量名，命名空间部分大小写无关
		if idx := strings.LastIndexByte(name, NsSeparator); idx >= 0 {
			name = strings.ToLower(name[:idx]) + name[idx:]
		}
	}

	ctx.EG().ConstantTable().Set(name, c)
}

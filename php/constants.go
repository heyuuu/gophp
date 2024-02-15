package php

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

const constHaltOffset = "__COMPILER_HALT_OFFSET__"

func RegisterConstantEx(ctx *Context, moduleNumber int, name string, value types.Zval, flags types.ConstFlag) {
	c := types.NewConstant(name, value, flags, moduleNumber)
	innerRegisterConstant(ctx, c)
}
func RegisterConstant(ctx *Context, moduleNumber int, name string, value types.Zval) {
	RegisterConstantEx(ctx, moduleNumber, name, value, types.ConstPersistent|types.ConstCs)
}
func RegisterLongConstant(ctx *Context, moduleNumber int, name string, value int) {
	RegisterConstant(ctx, moduleNumber, name, types.ZvalLong(value))
}
func RegisterDoubleConstant(ctx *Context, moduleNumber int, name string, value float64) {
	RegisterConstant(ctx, moduleNumber, name, types.ZvalDouble(value))
}
func RegisterStringConstant(ctx *Context, moduleNumber int, name string, str string) {
	RegisterConstant(ctx, moduleNumber, name, types.ZvalString(str))
}

func RegisterUserConstant(ctx *Context, name string, value types.Zval, flags types.ConstFlag) bool {
	c := types.NewConstant(name, value, flags, types.PhpUserConstant)
	return innerRegisterConstant(ctx, c)
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

func GetConstantEx(ctx *Context, name string, scope *types.Class, flags uint32) types.Zval {
	return GetConstant(ctx, name).Value()
}

func innerRegisterConstant(ctx *Context, c *types.Constant) bool {
	name := c.Name()
	if !c.IsCaseSensitive() {
		name = ascii.StrToLower(name)
	} else {
		// 带命名空间的常量名，命名空间部分大小写无关
		if idx := strings.LastIndexByte(name, NsSeparator); idx >= 0 {
			name = ascii.StrToLower(name[:idx]) + name[idx:]
		}
	}

	copyConstant := *c
	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */
	if name == constHaltOffset || !ctx.EG().ConstantTable().Add(name, &copyConstant) {
		Error(ctx, perr.E_NOTICE, fmt.Sprintf("Constant %s already defined", name))
		return false
	}

	return true
}

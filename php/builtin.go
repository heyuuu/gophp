package php

import "github.com/heyuuu/gophp/php/types"

var BuiltinModule = ModuleEntry{
	Name: "Core",
	ModuleStartup: func(ctx *Context, moduleNumber int) bool {
		RegisterCodeClasses(ctx, moduleNumber)
		RegisterStandardConstants(ctx, moduleNumber)
		return true
	},
}

func RegisterCodeClasses(ctx *Context, moduleNumber int) {

}

func RegisterStandardConstants(ctx *Context, moduleNumber int) {
	RegisterConstantEx(ctx, moduleNumber, "TRUE", True(), types.ConstPersistent|types.ConstCtSubst)
	RegisterConstantEx(ctx, moduleNumber, "FALSE", False(), types.ConstPersistent|types.ConstCtSubst)
	RegisterConstantEx(ctx, moduleNumber, "NULL", False(), types.ConstPersistent|types.ConstCtSubst)
}

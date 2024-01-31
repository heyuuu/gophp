package php

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func RegisterInternalClass(ctx *Context, moduleNumber int, name string) *types.Class {
	ce := types.NewInternalClass(name, moduleNumber, 0)
	iRegisterClass(ctx, ce)
	return ce
}
func RegisterUserClass(ctx *Context, name string) *types.Class {
	ce := types.NewUserClass(name)
	iRegisterClass(ctx, ce)
	return ce
}
func iRegisterClass(ctx *Context, ce *types.Class) {
	ctx.EG().ClassTable().Set(ce.LcName(), ce)
}

func RegisterModuleFunctions(ctx *Context, m *Module, functions []types.FunctionEntry) {
	for _, entry := range functions {
		fn := types.NewInternalFunctionByEntry(m.moduleNumber, entry)
		RegisterFunction(ctx, entry.Name(), fn)
	}
}

func RegisterFunction(ctx *Context, name string, fn *types.Function) bool {
	// todo check 冲突
	lcName := strings.ToLower(name)
	return ctx.EG().FunctionTable().Add(lcName, fn)
}

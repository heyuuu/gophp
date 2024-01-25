package php

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

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

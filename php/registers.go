package php

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func RegisterModuleFunctions(ctx *Context, m *Module, functions []FunctionEntry) {
	for _, entry := range functions {
		lcName := strings.ToLower(entry.Name())
		fn := types.NewInternalFunction(entry.Name(), entry.Handler(), m.moduleNumber)
		// todo check 冲突
		ctx.EG().functionTable.Add(lcName, fn)
	}
}

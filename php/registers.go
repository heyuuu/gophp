package php

import (
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

// register class
func RegisterInternalInterface(ctx *Context, moduleNumber int, decl *types.InternalClassDecl) *types.Class {
	decl.Flags |= types.AccInterface
	return RegisterInternalClass(ctx, moduleNumber, decl)
}

func RegisterInternalClass(ctx *Context, moduleNumber int, decl *types.InternalClassDecl) *types.Class {
	assert.Assert(decl.Name != "")
	ce := types.NewInternalClass(decl, moduleNumber)
	ctx.EG().ClassTable().Set(ce.LcName(), ce)
	return ce
}

func RegisterUserClass(ctx *Context, decl *types.UserClassDecl) *types.Class {
	ce := types.NewUserClass(decl)
	ctx.EG().ClassTable().Set(ce.LcName(), ce)
	return ce
}

func RegisterModuleFunctions(ctx *Context, m *Module, functions []types.FunctionDecl) {
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

package builtin

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

func init() {
	php.AddBuiltinModule(BuiltinModuleEntry)
}

var BuiltinModuleEntry = php.ModuleEntry{
	Name:      "Core",
	Functions: zifFunctions,
	ModuleStartup: func(ctx *php.Context, moduleNumber int) bool {
		RegisterCoreClasses(ctx, moduleNumber)
		RegisterStandardConstants(ctx, moduleNumber)
		return true
	},
}

func RegisterCoreClasses(ctx *php.Context, moduleNumber int) {

}

func RegisterStandardConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterConstantEx(ctx, moduleNumber, "TRUE", types.True, types.ConstPersistent|types.ConstCtSubst)
	php.RegisterConstantEx(ctx, moduleNumber, "FALSE", types.False, types.ConstPersistent|types.ConstCtSubst)
	php.RegisterConstantEx(ctx, moduleNumber, "NULL", types.Null, types.ConstPersistent|types.ConstCtSubst)
}

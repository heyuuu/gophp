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
		RegisterStdClass(ctx, moduleNumber)
		RegisterCoreClasses(ctx, moduleNumber)
		php.ZendRegisterStandardConstants(ctx)
		return true
	},
}

func RegisterStdClass(ctx *php.Context, moduleNumber int) {
	php.RegisterInternalClass(ctx, moduleNumber, &types.InternalClassDecl{
		Name: php.StdClassName,
	})
}

func RegisterCoreClasses(ctx *php.Context, moduleNumber int) {

}

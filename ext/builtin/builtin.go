package builtin

import (
	"github.com/heyuuu/gophp/php"
)

func init() {
	php.AddBuiltinModule(BuiltinModuleEntry)
}

var BuiltinModuleEntry = php.ModuleEntry{
	Name:      "Core",
	Functions: zifFunctions,
	ModuleStartup: func(ctx *php.Context, moduleNumber int) bool {
		RegisterCoreClasses(ctx, moduleNumber)
		php.ZendRegisterStandardConstants(ctx)
		return true
	},
}

func RegisterCoreClasses(ctx *php.Context, moduleNumber int) {

}

package standard

import (
	"github.com/heyuuu/gophp/php"
)

func init() {
	php.AddBuiltinModule(BasicModuleEntry)
}

var BasicModuleEntry = php.ModuleEntry{
	Name:          "standard",
	Functions:     zifFunctions,
	ModuleStartup: ZmStartupBasic,
}

func ZmStartupBasic(ctx *php.Context, moduleNumber int) bool {
	RegisterStringConstants(ctx, moduleNumber)
	return true
}

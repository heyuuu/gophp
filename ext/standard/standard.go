package standard

import "github.com/heyuuu/gophp/php"

func init() {
	php.AddBuiltinModule(BasicModuleEntry)
}

var BasicModuleEntry = php.ModuleEntry{
	Name: "standard",
	Functions: []php.FunctionEntry{
		php.DefFunctionEntry("var_dump", ZifVarDump),
	},
}

package php

var builtinModuleEntries []ModuleEntry

func AddBuiltinModule(m ModuleEntry) {
	builtinModuleEntries = append(builtinModuleEntries, m)
}

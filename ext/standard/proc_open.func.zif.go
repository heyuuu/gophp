package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifProcTerminate
var DefZifProcTerminate = def.DefFunc("proc_terminate", 1, 2, []def.ArgInfo{{name: "process"}, {name: "signal"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	process := fp.ParseZval()
	fp.StartOptional()
	signal := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifProcTerminate(executeData, returnValue, process, nil, signal)
})

// generate by ZifProcOpen
var DefZifProcOpen = def.DefFunc("proc_open", 3, 6, []def.ArgInfo{{name: "command"}, {name: "descriptorspec"}, {name: "pipes"}, {name: "cwd"}, {name: "env"}, {name: "other_options"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 6, 0)
	command := fp.ParseZval()
	descriptorspec := fp.ParseZval()
	pipes := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	cwd := fp.ParseZval()
	env := fp.ParseZval()
	other_options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifProcOpen(executeData, returnValue, command, descriptorspec, pipes, nil, cwd, env, other_options)
})

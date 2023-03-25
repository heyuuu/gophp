package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifProcTerminate
var DefZifProcTerminate = def.DefFunc("proc_terminate", 1, 2, []def.ArgInfo{{Name: "process"}, {Name: "signal"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	process := fp.ParseZval()
	fp.StartOptional()
	signal := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifProcTerminate(executeData, returnValue, process, nil, signal)
})

// generate by ZifProcClose
var DefZifProcClose = def.DefFunc("proc_close", 1, 1, []def.ArgInfo{{Name: "process"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	process := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifProcClose(executeData, returnValue, process)
})

// generate by ZifProcGetStatus
var DefZifProcGetStatus = def.DefFunc("proc_get_status", 1, 1, []def.ArgInfo{{Name: "process"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	process := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifProcGetStatus(executeData, returnValue, process)
})

// generate by ZifProcOpen
var DefZifProcOpen = def.DefFunc("proc_open", 3, 6, []def.ArgInfo{{Name: "command"}, {Name: "descriptorspec"}, {Name: "pipes"}, {Name: "cwd"}, {Name: "env"}, {Name: "other_options"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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

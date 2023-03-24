package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifExec
var DefZifExec = def.DefFunc("exec", 1, 3, []def.ArgInfo{{name: "command"}, {name: "output"}, {name: "return_value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	output := fp.ParseZvalEx(false, true)
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifExec(executeData, returnValue, command, nil, output, return_value)
})

// generate by ZifSystem
var DefZifSystem = def.DefFunc("system", 1, 2, []def.ArgInfo{{name: "command"}, {name: "return_value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifSystem(executeData, returnValue, command, nil, return_value)
})

// generate by ZifPassthru
var DefZifPassthru = def.DefFunc("passthru", 1, 2, []def.ArgInfo{{name: "command"}, {name: "return_value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifPassthru(executeData, returnValue, command, nil, return_value)
})

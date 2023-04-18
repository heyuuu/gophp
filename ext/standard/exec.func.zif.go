package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifExec
var DefZifExec = def.DefFunc("exec", 1, 3, []def.ArgInfo{{Name: "command"}, {Name: "output"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifSystem = def.DefFunc("system", 1, 2, []def.ArgInfo{{Name: "command"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifPassthru = def.DefFunc("passthru", 1, 2, []def.ArgInfo{{Name: "command"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifPassthru(executeData, returnValue, command, nil, return_value)
})

// generate by ZifEscapeshellcmd
var DefZifEscapeshellcmd = def.DefFunc("escapeshellcmd", 1, 1, []def.ArgInfo{{Name: "command"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	command := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifEscapeshellcmd(command)
	returnValue.SetStringVal(ret)
})

// generate by ZifEscapeshellarg
var DefZifEscapeshellarg = def.DefFunc("escapeshellarg", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifEscapeshellarg(arg)
	returnValue.SetStringVal(ret)
})

// generate by ZifShellExec
var DefZifShellExec = def.DefFunc("shell_exec", 1, 1, []def.ArgInfo{{Name: "cmd"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	cmd := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifShellExec(executeData, returnValue, cmd)
})

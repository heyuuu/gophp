package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifFsockopen
var DefZifFsockopen = def.DefFunc("fsockopen", 1, 5, []def.ArgInfo{{name: "hostname"}, {name: "port"}, {name: "errno"}, {name: "errstr"}, {name: "timeout"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	hostname := fp.ParseZval()
	fp.StartOptional()
	port := fp.ParseZval()
	errno := fp.ParseZvalEx(false, true)
	errstr := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFsockopen(executeData, returnValue, hostname, nil, port, errno, errstr, timeout)
})

// generate by ZifPfsockopen
var DefZifPfsockopen = def.DefFunc("pfsockopen", 1, 5, []def.ArgInfo{{name: "hostname"}, {name: "port"}, {name: "errno"}, {name: "errstr"}, {name: "timeout"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	hostname := fp.ParseZval()
	fp.StartOptional()
	port := fp.ParseZval()
	errno := fp.ParseZvalEx(false, true)
	errstr := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPfsockopen(executeData, returnValue, hostname, nil, port, errno, errstr, timeout)
})

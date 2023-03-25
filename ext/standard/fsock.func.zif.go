package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifFsockopen
var DefZifFsockopen = def.DefFunc("fsockopen", 1, 5, []def.ArgInfo{{Name: "hostname"}, {Name: "port"}, {Name: "errno"}, {Name: "errstr"}, {Name: "timeout"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifPfsockopen = def.DefFunc("pfsockopen", 1, 5, []def.ArgInfo{{Name: "hostname"}, {Name: "port"}, {Name: "errno"}, {Name: "errstr"}, {Name: "timeout"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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

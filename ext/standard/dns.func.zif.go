package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifGethostname
var DefZifGethostname = def.DefFunc("gethostname", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGethostname(executeData, returnValue)
})

// generate by ZifGethostbyaddr
var DefZifGethostbyaddr = def.DefFunc("gethostbyaddr", 1, 1, []def.ArgInfo{{name: "ip_address"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyaddr(executeData, returnValue, ip_address)
})

// generate by ZifGethostbyname
var DefZifGethostbyname = def.DefFunc("gethostbyname", 1, 1, []def.ArgInfo{{name: "hostname"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyname(executeData, returnValue, hostname)
})

// generate by ZifGethostbynamel
var DefZifGethostbynamel = def.DefFunc("gethostbynamel", 1, 1, []def.ArgInfo{{name: "hostname"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbynamel(executeData, returnValue, hostname)
})

// generate by ZifDnsCheckRecord
var DefZifDnsCheckRecord = def.DefFunc("dns_check_record", 1, 2, []def.ArgInfo{{name: "host"}, {name: "type_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	host := fp.ParseZval()
	fp.StartOptional()
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDnsCheckRecord(executeData, returnValue, host, nil, type_)
})

// generate by ZifDnsGetMx
var DefZifDnsGetMx = def.DefFunc("dns_get_mx", 2, 3, []def.ArgInfo{{name: "hostname"}, {name: "mxhosts"}, {name: "weight"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hostname := fp.ParseZval()
	mxhosts := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	weight := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifDnsGetMx(executeData, returnValue, hostname, mxhosts, nil, weight)
})

package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

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

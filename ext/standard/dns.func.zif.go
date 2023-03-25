package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifGethostname
var DefZifGethostname = def.DefFunc("gethostname", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGethostname(executeData, returnValue)
})

// generate by ZifGethostbyaddr
var DefZifGethostbyaddr = def.DefFunc("gethostbyaddr", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyaddr(executeData, returnValue, ip_address)
})

// generate by ZifGethostbyname
var DefZifGethostbyname = def.DefFunc("gethostbyname", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyname(executeData, returnValue, hostname)
})

// generate by ZifGethostbynamel
var DefZifGethostbynamel = def.DefFunc("gethostbynamel", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbynamel(executeData, returnValue, hostname)
})

// generate by ZifDnsCheckRecord
var DefZifDnsCheckRecord = def.DefFunc("dns_check_record", 1, 2, []def.ArgInfo{{Name: "host"}, {Name: "type_"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	host := fp.ParseZval()
	fp.StartOptional()
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDnsCheckRecord(executeData, returnValue, host, nil, type_)
})

// generate by ZifDnsCheckRecord
var DefZifCheckdnsrr = def.DefFunc("checkdnsrr", 1, 2, []def.ArgInfo{{Name: "host"}, {Name: "type_"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifDnsGetMx = def.DefFunc("dns_get_mx", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "mxhosts"}, {Name: "weight"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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

// generate by ZifDnsGetMx
var DefZifGetmxrr = def.DefFunc("getmxrr", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "mxhosts"}, {Name: "weight"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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

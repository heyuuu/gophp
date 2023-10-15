package cli

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifApacheRequestHeaders
var DefZifApacheRequestHeaders = def.DefFunc("apache_request_headers", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifApacheRequestHeaders(executeData, returnValue)
})

// generate by ZifApacheRequestHeaders
var DefZifGetallheaders = def.DefFunc("getallheaders", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifApacheRequestHeaders(executeData, returnValue)
})

// generate by ZifApacheResponseHeaders
var DefZifApacheResponseHeaders = def.DefFunc("apache_response_headers", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifApacheResponseHeaders(returnValue)
})

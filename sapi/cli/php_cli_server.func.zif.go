package cli

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifApacheRequestHeaders
var DefZifApacheRequestHeaders = def.DefFunc("apache_request_headers", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifApacheRequestHeaders(executeData, returnValue)
})

// generate by ZifApacheResponseHeaders
var DefZifApacheResponseHeaders = def.DefFunc("apache_response_headers", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifApacheResponseHeaders(executeData, returnValue)
})

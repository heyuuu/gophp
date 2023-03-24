package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifHeader
var DefZifHeader = def.DefFunc("header", 1, 3, []def.ArgInfo{{name: "header"}, {name: "replace"}, {name: "http_response_code"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	header := fp.ParseZval()
	fp.StartOptional()
	replace := fp.ParseZval()
	http_response_code := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHeader(executeData, returnValue, header, nil, replace, http_response_code)
})

// generate by ZifHeaderRemove
var DefZifHeaderRemove = def.DefFunc("header_remove", 0, 1, []def.ArgInfo{{name: "name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHeaderRemove(executeData, returnValue, nil, name)
})

// generate by ZifSetcookie
var DefZifSetcookie = def.DefFunc("setcookie", 1, 7, []def.ArgInfo{{name: "name"}, {name: "value"}, {name: "expires_or_options"}, {name: "path"}, {name: "domain"}, {name: "secure"}, {name: "httponly"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 7, 0)
	name := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	expires_or_options := fp.ParseZval()
	path := fp.ParseZval()
	domain := fp.ParseZval()
	secure := fp.ParseZval()
	httponly := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetcookie(executeData, returnValue, name, nil, value, expires_or_options, path, domain, secure, httponly)
})

// generate by ZifSetrawcookie
var DefZifSetrawcookie = def.DefFunc("setrawcookie", 1, 7, []def.ArgInfo{{name: "name"}, {name: "value"}, {name: "expires_or_options"}, {name: "path"}, {name: "domain"}, {name: "secure"}, {name: "httponly"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 7, 0)
	name := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	expires_or_options := fp.ParseZval()
	path := fp.ParseZval()
	domain := fp.ParseZval()
	secure := fp.ParseZval()
	httponly := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetrawcookie(executeData, returnValue, name, nil, value, expires_or_options, path, domain, secure, httponly)
})

// generate by ZifHeadersSent
var DefZifHeadersSent = def.DefFunc("headers_sent", 0, 2, []def.ArgInfo{{name: "file"}, {name: "line"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	file := fp.ParseZvalEx(false, true)
	line := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifHeadersSent(executeData, returnValue, nil, file, line)
})

// generate by ZifHeadersList
var DefZifHeadersList = def.DefFunc("headers_list", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifHeadersList(executeData, returnValue)
})

// generate by ZifHttpResponseCode
var DefZifHttpResponseCode = def.DefFunc("http_response_code", 0, 1, []def.ArgInfo{{name: "response_code"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	response_code := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHttpResponseCode(executeData, returnValue, nil, response_code)
})

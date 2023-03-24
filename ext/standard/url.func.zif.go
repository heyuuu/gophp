package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifParseUrl
var DefZifParseUrl = def.DefFunc("parse_url", 1, 2, []def.ArgInfo{{name: "url"}, {name: "component"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	url := fp.ParseZval()
	fp.StartOptional()
	component := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifParseUrl(executeData, returnValue, url, nil, component)
})

// generate by ZifUrlencode
var DefZifUrlencode = def.DefFunc("urlencode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUrlencode(executeData, returnValue, str)
})

// generate by ZifUrldecode
var DefZifUrldecode = def.DefFunc("urldecode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUrldecode(executeData, returnValue, str)
})

// generate by ZifRawurlencode
var DefZifRawurlencode = def.DefFunc("rawurlencode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRawurlencode(executeData, returnValue, str)
})

// generate by ZifRawurldecode
var DefZifRawurldecode = def.DefFunc("rawurldecode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRawurldecode(executeData, returnValue, str)
})

// generate by ZifGetHeaders
var DefZifGetHeaders = def.DefFunc("get_headers", 1, 3, []def.ArgInfo{{name: "url"}, {name: "format"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	url := fp.ParseZval()
	fp.StartOptional()
	format := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetHeaders(executeData, returnValue, url, nil, format, context)
})

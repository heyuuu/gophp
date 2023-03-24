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

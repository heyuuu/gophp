package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifHttpBuildQuery
var DefZifHttpBuildQuery = def.DefFunc("http_build_query", 1, 4, []def.ArgInfo{{name: "formdata"}, {name: "prefix"}, {name: "arg_separator"}, {name: "enc_type"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	formdata := fp.ParseZval()
	fp.StartOptional()
	prefix := fp.ParseZval()
	arg_separator := fp.ParseZval()
	enc_type := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHttpBuildQuery(executeData, returnValue, formdata, nil, prefix, arg_separator, enc_type)
})

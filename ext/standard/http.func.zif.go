package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifHttpBuildQuery
var DefZifHttpBuildQuery = def.DefFunc("http_build_query", 1, 4, []def.ArgInfo{{Name: "formdata"}, {Name: "prefix"}, {Name: "arg_separator"}, {Name: "enc_type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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

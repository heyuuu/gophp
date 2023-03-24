package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifIptcembed
var DefZifIptcembed = def.DefFunc("iptcembed", 2, 3, []def.ArgInfo{{name: "iptcdata"}, {name: "jpeg_file_name"}, {name: "spool"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	iptcdata := fp.ParseZval()
	jpeg_file_name := fp.ParseZval()
	fp.StartOptional()
	spool := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIptcembed(executeData, returnValue, iptcdata, jpeg_file_name, nil, spool)
})

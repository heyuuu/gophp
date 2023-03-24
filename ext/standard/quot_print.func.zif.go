package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifQuotedPrintableDecode
var DefZifQuotedPrintableDecode = def.DefFunc("quoted_printable_decode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifQuotedPrintableDecode(executeData, returnValue, str)
})

// generate by ZifQuotedPrintableEncode
var DefZifQuotedPrintableEncode = def.DefFunc("quoted_printable_encode", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifQuotedPrintableEncode(executeData, returnValue, str)
})

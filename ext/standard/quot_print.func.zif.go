package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifQuotedPrintableDecode
var DefZifQuotedPrintableDecode = def.DefFunc("quoted_printable_decode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableDecode(str)
	returnValue.SetString(ret)
})

// generate by ZifQuotedPrintableEncode
var DefZifQuotedPrintableEncode = def.DefFunc("quoted_printable_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableEncode(str)
	returnValue.SetString(ret)
})

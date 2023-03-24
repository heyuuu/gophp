package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifRound
var DefZifRound = def.DefFunc("round", 1, 3, []def.ArgInfo{{name: "number"}, {name: "precision"}, {name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	precision := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRound(executeData, returnValue, number, nil, precision, mode)
})

// generate by ZifLog
var DefZifLog = def.DefFunc("log", 1, 2, []def.ArgInfo{{name: "number"}, {name: "base"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	base := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog(executeData, returnValue, number, nil, base)
})

// generate by ZifNumberFormat
var DefZifNumberFormat = def.DefFunc("number_format", 1, 4, []def.ArgInfo{{name: "number"}, {name: "num_decimal_places"}, {name: "dec_separator"}, {name: "thousands_separator"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	num_decimal_places := fp.ParseZval()
	dec_separator := fp.ParseZval()
	thousands_separator := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifNumberFormat(executeData, returnValue, number, nil, num_decimal_places, dec_separator, thousands_separator)
})

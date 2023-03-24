package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifHtmlspecialchars
var DefZifHtmlspecialchars = def.DefFunc("htmlspecialchars", 1, 4, []def.ArgInfo{{name: "string"}, {name: "quote_style"}, {name: "encoding"}, {name: "double_encode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	double_encode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlspecialchars(executeData, returnValue, string, nil, quote_style, encoding, double_encode)
})

// generate by ZifHtmlspecialcharsDecode
var DefZifHtmlspecialcharsDecode = def.DefFunc("htmlspecialchars_decode", 1, 2, []def.ArgInfo{{name: "string"}, {name: "quote_style"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlspecialcharsDecode(executeData, returnValue, string, nil, quote_style)
})

// generate by ZifHtmlEntityDecode
var DefZifHtmlEntityDecode = def.DefFunc("html_entity_decode", 1, 3, []def.ArgInfo{{name: "string"}, {name: "quote_style"}, {name: "encoding"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlEntityDecode(executeData, returnValue, string, nil, quote_style, encoding)
})

// generate by ZifHtmlentities
var DefZifHtmlentities = def.DefFunc("htmlentities", 1, 4, []def.ArgInfo{{name: "string"}, {name: "quote_style"}, {name: "encoding"}, {name: "double_encode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	double_encode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlentities(executeData, returnValue, string, nil, quote_style, encoding, double_encode)
})

// generate by ZifGetHtmlTranslationTable
var DefZifGetHtmlTranslationTable = def.DefFunc("get_html_translation_table", 0, 3, []def.ArgInfo{{name: "table"}, {name: "quote_style"}, {name: "encoding"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 3, 0)
	fp.StartOptional()
	table := fp.ParseZval()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetHtmlTranslationTable(executeData, returnValue, nil, table, quote_style, encoding)
})

package standard

import (
	"sik/zend/zpp"
	"sik/zend/def"
	"sik/zend/types"
)

// generate by ZifAssert
var DefZifAssert = def.DefFunc("assert", 1, 2, []def.ArgInfo{{name: "assertion"}, {name: "description"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	assertion := fp.ParseZval()
	fp.StartOptional()
	description := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAssert(executeData, returnValue, assertion, nil, description)
})

// generate by ZifAssertOptions
var DefZifAssertOptions = def.DefFunc("assert_options", 1, 2, []def.ArgInfo{{name: "what"}, {name: "value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	what := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAssertOptions(executeData, returnValue, what, nil, value)
})

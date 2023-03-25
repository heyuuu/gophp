package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifAssert
var DefZifAssert = def.DefFunc("assert", 1, 2, []def.ArgInfo{{Name: "assertion"}, {Name: "description"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifAssertOptions = def.DefFunc("assert_options", 1, 2, []def.ArgInfo{{Name: "what"}, {Name: "value"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	what := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAssertOptions(executeData, returnValue, what, nil, value)
})

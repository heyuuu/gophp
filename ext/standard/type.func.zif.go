package standard

import (
	"sik/zend/types"
	"sik/zend/zpp"
	"sik/zend/def"
)

// generate by ZifIntval
var DefZifIntval = def.DefFunc("intval", 1, 2, []def.ArgInfo{{name: "var_"}, {name: "base"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	base := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIntval(executeData, returnValue, var_, nil, base)
})

// generate by ZifIsCallable
var DefZifIsCallable = def.DefFunc("is_callable", 1, 3, []def.ArgInfo{{name: "var_"}, {name: "syntax_only"}, {name: "callable_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	syntax_only := fp.ParseZval()
	callable_name := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifIsCallable(executeData, returnValue, var_, nil, syntax_only, callable_name)
})

// generate by ZifIsIterable
var DefZifIsIterable = def.DefFunc("is_iterable", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsIterable(executeData, returnValue, var_)
})

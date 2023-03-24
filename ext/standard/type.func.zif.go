package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifGettype
var DefZifGettype = def.DefFunc("gettype", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGettype(executeData, returnValue, var_)
})

// generate by ZifSettype
var DefZifSettype = def.DefFunc("settype", 2, 2, []def.ArgInfo{{name: "var_"}, {name: "type_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	var_ := fp.ParseZvalEx(false, true)
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSettype(executeData, returnValue, var_, type_)
})

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

// generate by ZifFloatval
var DefZifFloatval = def.DefFunc("floatval", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFloatval(executeData, returnValue, var_)
})

// generate by ZifBoolval
var DefZifBoolval = def.DefFunc("boolval", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBoolval(executeData, returnValue, var_)
})

// generate by ZifStrval
var DefZifStrval = def.DefFunc("strval", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrval(executeData, returnValue, var_)
})

// generate by ZifIsNull
var DefZifIsNull = def.DefFunc("is_null", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNull(executeData, returnValue, var_)
})

// generate by ZifIsResource
var DefZifIsResource = def.DefFunc("is_resource", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsResource(executeData, returnValue, var_)
})

// generate by ZifIsBool
var DefZifIsBool = def.DefFunc("is_bool", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsBool(executeData, returnValue, var_)
})

// generate by ZifIsInt
var DefZifIsInt = def.DefFunc("is_int", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInt(executeData, returnValue, var_)
})

// generate by ZifIsFloat
var DefZifIsFloat = def.DefFunc("is_float", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFloat(executeData, returnValue, var_)
})

// generate by ZifIsString
var DefZifIsString = def.DefFunc("is_string", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsString(executeData, returnValue, var_)
})

// generate by ZifIsArray
var DefZifIsArray = def.DefFunc("is_array", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsArray(executeData, returnValue, var_)
})

// generate by ZifIsObject
var DefZifIsObject = def.DefFunc("is_object", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsObject(executeData, returnValue, var_)
})

// generate by ZifIsNumeric
var DefZifIsNumeric = def.DefFunc("is_numeric", 1, 1, []def.ArgInfo{{name: "value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNumeric(executeData, returnValue, value)
})

// generate by ZifIsScalar
var DefZifIsScalar = def.DefFunc("is_scalar", 1, 1, []def.ArgInfo{{name: "value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsScalar(executeData, returnValue, value)
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

// generate by ZifIsCountable
var DefZifIsCountable = def.DefFunc("is_countable", 1, 1, []def.ArgInfo{{name: "var_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsCountable(executeData, returnValue, var_)
})

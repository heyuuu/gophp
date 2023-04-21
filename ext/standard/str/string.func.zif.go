package str

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifParseStr
var DefZifParseStr = def.DefFunc("parse_str", 1, 2, []def.ArgInfo{{Name: "encoded_string"}, {Name: "result"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	encoded_string := fp.ParseStringVal()
	fp.StartOptional()
	result := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifParseStr(encoded_string, nil, result)
})

// generate by ZifStrGetcsv
var DefZifStrGetcsv = def.DefFunc("str_getcsv", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string_ := fp.ParseStringVal()
	fp.StartOptional()
	delimiter := fp.ParseStringValNullable()
	enclosure := fp.ParseStringValNullable()
	escape := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifStrGetcsv(returnValue, string_, nil, delimiter, enclosure, escape)
	returnValue.SetArray(ret)
})

// generate by ZifSscanf
var DefZifSscanf = def.DefFunc("sscanf", 3, 3, []def.ArgInfo{{Name: "str"}, {Name: "format"}, {Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str := fp.ParseStringVal()
	format := fp.ParseStringVal()
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifSscanf(str, format, vars)
	returnValue.SetBy(ret)
})

// generate by ZifUtf8Encode
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Encode(data)
	returnValue.SetStringVal(ret)
})

// generate by ZifUtf8Decode
var DefZifUtf8Decode = def.DefFunc("utf8_decode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Decode(data)
	returnValue.SetStringVal(ret)
})

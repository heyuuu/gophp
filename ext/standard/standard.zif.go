package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifUtf8Encode, DefZifBin2hex, DefZifHex2bin, DefZifStrcoll, DefZifTrim, DefZifRtrim, DefZifChop, DefZifLtrim, DefZifWordwrap, DefZifExplode, DefZifImplode, DefZifJoin, DefZifStrtok, DefZifStrtoupper, DefZifStrtolower, DefZifBasename, DefZifDirname, DefZifVarDump}

// generate by ZifUtf8Encode
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Encode(data)
	returnValue.SetString(ret)
})

// generate by ZifBin2hex
var DefZifBin2hex = def.DefFunc("bin2hex", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifBin2hex(data)
	returnValue.SetString(ret)
})

// generate by ZifHex2bin
var DefZifHex2bin = def.DefFunc("hex2bin", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHex2bin(executeData.Ctx(), data)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrcoll
var DefZifStrcoll = def.DefFunc("strcoll", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	str1 := fp.ParseString()
	str2 := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrcoll(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifTrim
var DefZifTrim = def.DefFunc("trim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifTrim(executeData.Ctx(), str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifRtrim
var DefZifRtrim = def.DefFunc("rtrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(executeData.Ctx(), str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifRtrim
var DefZifChop = def.DefFunc("chop", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(executeData.Ctx(), str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifLtrim
var DefZifLtrim = def.DefFunc("ltrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifLtrim(executeData.Ctx(), str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifWordwrap
var DefZifWordwrap = def.DefFunc("wordwrap", 1, 4, []def.ArgInfo{{Name: "str"}, {Name: "width"}, {Name: "break"}, {Name: "cut"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 4, 0)
	str := fp.ParseString()
	fp.StartOptional()
	width := fp.ParseLongNullable()
	break_ := fp.ParseStringNullable()
	cut := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifWordwrap(executeData.Ctx(), str, nil, width, break_, cut)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifExplode
var DefZifExplode = def.DefFunc("explode", 2, 3, []def.ArgInfo{{Name: "separator"}, {Name: "str"}, {Name: "limit"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	separator := fp.ParseString()
	str := fp.ParseString()
	fp.StartOptional()
	limit_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifExplode(executeData.Ctx(), separator, str, nil, limit_)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifImplode
var DefZifImplode = def.DefFunc("implode", 1, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	glue_ := fp.ParseZvalPtr()
	fp.StartOptional()
	pieces_ := fp.ParseZvalPtr()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(executeData.Ctx(), glue_, nil, pieces_)
	returnValue.SetString(ret)
})

// generate by ZifImplode
var DefZifJoin = def.DefFunc("join", 1, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	glue_ := fp.ParseZvalPtr()
	fp.StartOptional()
	pieces_ := fp.ParseZvalPtr()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(executeData.Ctx(), glue_, nil, pieces_)
	returnValue.SetString(ret)
})

// generate by ZifStrtok
var DefZifStrtok = def.DefFunc("strtok", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "token"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	token_ := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrtok(executeData.Ctx(), str, nil, token_)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrtoupper
var DefZifStrtoupper = def.DefFunc("strtoupper", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrtoupper(str)
	returnValue.SetString(ret)
})

// generate by ZifStrtolower
var DefZifStrtolower = def.DefFunc("strtolower", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrtolower(str)
	returnValue.SetString(ret)
})

// generate by ZifBasename
var DefZifBasename = def.DefFunc("basename", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "suffix"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	path := fp.ParseString()
	fp.StartOptional()
	suffix := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifBasename(path, nil, suffix)
	returnValue.SetString(ret)
})

// generate by ZifDirname
var DefZifDirname = def.DefFunc("dirname", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "levels"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	path := fp.ParseString()
	fp.StartOptional()
	levels_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifDirname(executeData.Ctx(), path, nil, levels_)
	returnValue.SetString(ret)
})

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 0, -1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 0, -1, 0)
	vars := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifVarDump(executeData.Ctx(), vars)
})

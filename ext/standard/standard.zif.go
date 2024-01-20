package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifUtf8Encode, DefZifBin2hex, DefZifHex2bin, DefZifStrspn, DefZifStrcspn, DefZifStrcoll, DefZifTrim, DefZifRtrim, DefZifChop, DefZifLtrim, DefZifWordwrap, DefZifExplode, DefZifImplode, DefZifJoin, DefZifStrtok, DefZifStrtoupper, DefZifStrtolower, DefZifBasename, DefZifDirname, DefZifPathinfo, DefZifStristr, DefZifStrstr, DefZifStrchr, DefZifStrpos, DefZifStripos, DefZifStrrpos, DefZifStrripos, DefZifStrrchr, DefZifChunkSplit, DefZifSubstr, DefZifSubstrReplace, DefZifQuotemeta, DefZifOrd, DefZifChr, DefZifUcfirst, DefZifLcfirst, DefZifUcwords, DefZifStrtr, DefZifStrrev, DefZifSimilarText, DefZifAddslashes, DefZifAddcslashes, DefZifStripslashes, DefZifStripcslashes, DefZifStrReplace, DefZifStrIreplace, DefZifHebrev, DefZifHebrevc, DefZifNl2br, DefZifStripTags, DefZifStrRepeat, DefZifCountChars, DefZifStrnatcmp, DefZifStrnatcasecmp, DefZifSubstrCount, DefZifStrPad, DefZifStrRot13, DefZifStrShuffle, DefZifStrWordCount, DefZifStrSplit, DefZifStrpbrk, DefZifSubstrCompare, DefZifVarDump}

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

// generate by ZifStrspn
var DefZifStrspn = def.DefFunc("strspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "offset"}, {Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 4, 0)
	str := fp.ParseString()
	mask := fp.ParseString()
	fp.StartOptional()
	offset := fp.ParseLong()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrspn(str, mask, nil, offset, length)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrcspn
var DefZifStrcspn = def.DefFunc("strcspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "offset"}, {Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 4, 0)
	str := fp.ParseString()
	mask := fp.ParseString()
	fp.StartOptional()
	offset := fp.ParseLong()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrcspn(str, mask, nil, offset, length)
	if ok {
		returnValue.SetLong(ret)
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

// generate by ZifPathinfo
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "options"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	path := fp.ParseString()
	fp.StartOptional()
	options := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifPathinfo(path, nil, options)
	returnValue.SetVal(ret)
})

// generate by ZifStristr
var DefZifStristr = def.DefFunc("stristr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStristr(executeData.Ctx(), haystack, needle, nil, part)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrstr
var DefZifStrstr = def.DefFunc("strstr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrstr(executeData.Ctx(), haystack, needle, nil, part)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrstr
var DefZifStrchr = def.DefFunc("strchr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrstr(executeData.Ctx(), haystack, needle, nil, part)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrpos
var DefZifStrpos = def.DefFunc("strpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrpos(executeData.Ctx(), haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStripos
var DefZifStripos = def.DefFunc("stripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStripos(executeData.Ctx(), haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrpos
var DefZifStrrpos = def.DefFunc("strrpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrrpos(executeData.Ctx(), haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrripos
var DefZifStrripos = def.DefFunc("strripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrripos(executeData.Ctx(), haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrchr
var DefZifStrrchr = def.DefFunc("strrchr", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrrchr(executeData.Ctx(), haystack, needle)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifChunkSplit
var DefZifChunkSplit = def.DefFunc("chunk_split", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "chunklen"}, {Name: "ending"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 3, 0)
	str := fp.ParseString()
	fp.StartOptional()
	chunklen_ := fp.ParseLongNullable()
	ending_ := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifChunkSplit(executeData.Ctx(), str, nil, chunklen_, ending_)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstr
var DefZifSubstr = def.DefFunc("substr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "offset"}, {Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	str := fp.ParseString()
	offset := fp.ParseLong()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstr(str, offset, nil, length)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstrReplace
var DefZifSubstrReplace = def.DefFunc("substr_replace", 3, 4, []def.ArgInfo{{Name: "str"}, {Name: "replace"}, {Name: "start"}, {Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 3, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 3, 4, 0)
	str := fp.ParseZval()
	replace := fp.ParseZval()
	start := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZvalPtr()
	if fp.HasError() {
		return
	}
	ZifSubstrReplace(executeData.Ctx(), returnValue, str, replace, start, nil, length)
})

// generate by ZifQuotemeta
var DefZifQuotemeta = def.DefFunc("quotemeta", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifQuotemeta(str)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifOrd
var DefZifOrd = def.DefFunc("ord", 1, 1, []def.ArgInfo{{Name: "character"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	character := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifOrd(character)
	returnValue.SetLong(ret)
})

// generate by ZifChr
var DefZifChr = def.DefFunc("chr", 1, 1, []def.ArgInfo{{Name: "codepoint"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	codepoint := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifChr(codepoint)
	returnValue.SetString(ret)
})

// generate by ZifUcfirst
var DefZifUcfirst = def.DefFunc("ucfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUcfirst(str)
	returnValue.SetString(ret)
})

// generate by ZifLcfirst
var DefZifLcfirst = def.DefFunc("lcfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifLcfirst(str)
	returnValue.SetString(ret)
})

// generate by ZifUcwords
var DefZifUcwords = def.DefFunc("ucwords", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "delimiters"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	delimiters := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifUcwords(executeData.Ctx(), str, nil, delimiters)
	returnValue.SetString(ret)
})

// generate by ZifStrtr
var DefZifStrtr = def.DefFunc("strtr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "from"}, {Name: "to"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	str := fp.ParseString()
	from := fp.ParseZval()
	fp.StartOptional()
	to_ := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrtr(executeData.Ctx(), str, from, nil, to_)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrev
var DefZifStrrev = def.DefFunc("strrev", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrrev(str)
	returnValue.SetString(ret)
})

// generate by ZifSimilarText
var DefZifSimilarText = def.DefFunc("similar_text", 2, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "percent"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 3, 0)
	str1 := fp.ParseString()
	str2 := fp.ParseString()
	fp.StartOptional()
	percent := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ret := ZifSimilarText(executeData.Ctx(), str1, str2, nil, percent)
	returnValue.SetLong(ret)
})

// generate by ZifAddslashes
var DefZifAddslashes = def.DefFunc("addslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifAddslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifAddcslashes
var DefZifAddcslashes = def.DefFunc("addcslashes", 2, 2, []def.ArgInfo{{Name: "str"}, {Name: "charlist"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	str := fp.ParseString()
	charlist := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifAddcslashes(executeData.Ctx(), str, charlist)
	returnValue.SetString(ret)
})

// generate by ZifStripslashes
var DefZifStripslashes = def.DefFunc("stripslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStripslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifStripcslashes
var DefZifStripcslashes = def.DefFunc("stripcslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStripcslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifStrReplace
var DefZifStrReplace = def.DefFunc("str_replace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 3, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 3, 4, 0)
	search := fp.ParseZval()
	replace := fp.ParseZval()
	subject := fp.ParseZval()
	fp.StartOptional()
	replace_count := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ZifStrReplace(executeData.Ctx(), returnValue, search, replace, subject, nil, replace_count)
})

// generate by ZifStrIreplace
var DefZifStrIreplace = def.DefFunc("str_ireplace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 3, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 3, 4, 0)
	search := fp.ParseZval()
	replace := fp.ParseZval()
	subject := fp.ParseZval()
	fp.StartOptional()
	replace_count := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ZifStrIreplace(executeData.Ctx(), returnValue, search, replace, subject, nil, replace_count)
})

// generate by ZifHebrev
var DefZifHebrev = def.DefFunc("hebrev", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	max_chars_per_line := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHebrev(str, nil, max_chars_per_line)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifHebrevc
var DefZifHebrevc = def.DefFunc("hebrevc", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	max_chars_per_line := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHebrevc(str, nil, max_chars_per_line)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifNl2br
var DefZifNl2br = def.DefFunc("nl2br", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "is_xhtml"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	is_xhtml_ := fp.ParseBoolNullable()
	if fp.HasError() {
		return
	}
	ret := ZifNl2br(str, nil, is_xhtml_)
	returnValue.SetString(ret)
})

// generate by ZifStripTags
var DefZifStripTags = def.DefFunc("strip_tags", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "allowable_tags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	allowable_tags := fp.ParseZvalPtr()
	if fp.HasError() {
		return
	}
	ret := ZifStripTags(executeData.Ctx(), str, nil, allowable_tags)
	returnValue.SetString(ret)
})

// generate by ZifStrRepeat
var DefZifStrRepeat = def.DefFunc("str_repeat", 2, 2, []def.ArgInfo{{Name: "input"}, {Name: "mult"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	input := fp.ParseString()
	mult := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrRepeat(executeData.Ctx(), input, mult)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCountChars
var DefZifCountChars = def.DefFunc("count_chars", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	input := fp.ParseString()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCountChars(executeData.Ctx(), input, nil, mode)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrnatcmp
var DefZifStrnatcmp = def.DefFunc("strnatcmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	s1 := fp.ParseString()
	s2 := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrnatcmp(s1, s2)
	returnValue.SetLong(ret)
})

// generate by ZifStrnatcasecmp
var DefZifStrnatcasecmp = def.DefFunc("strnatcasecmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	s1 := fp.ParseString()
	s2 := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrnatcasecmp(s1, s2)
	returnValue.SetLong(ret)
})

// generate by ZifSubstrCount
var DefZifSubstrCount = def.DefFunc("substr_count", 2, 4, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}, {Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 4, 0)
	haystack := fp.ParseString()
	needle := fp.ParseString()
	fp.StartOptional()
	offset := fp.ParseLong()
	length_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstrCount(executeData.Ctx(), haystack, needle, nil, offset, length_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrPad
var DefZifStrPad = def.DefFunc("str_pad", 2, 4, []def.ArgInfo{{Name: "input"}, {Name: "pad_length"}, {Name: "pad_string"}, {Name: "pad_type"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 4, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 4, 0)
	input := fp.ParseString()
	pad_length := fp.ParseLong()
	fp.StartOptional()
	pad_string_ := fp.ParseStringNullable()
	pad_type_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrPad(executeData.Ctx(), input, pad_length, nil, pad_string_, pad_type_)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrRot13
var DefZifStrRot13 = def.DefFunc("str_rot13", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrRot13(str)
	returnValue.SetString(ret)
})

// generate by ZifStrShuffle
var DefZifStrShuffle = def.DefFunc("str_shuffle", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 1, 0)
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrShuffle(str)
	returnValue.SetString(ret)
})

// generate by ZifStrWordCount
var DefZifStrWordCount = def.DefFunc("str_word_count", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "format"}, {Name: "charlist"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 3, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 3, 0)
	str := fp.ParseString()
	fp.StartOptional()
	format := fp.ParseLong()
	charlist := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrWordCount(executeData.Ctx(), str, nil, format, charlist)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrSplit
var DefZifStrSplit = def.DefFunc("str_split", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "split_length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 1, 2, 0)
	str := fp.ParseString()
	fp.StartOptional()
	split_length_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrSplit(executeData.Ctx(), str, nil, split_length_)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrpbrk
var DefZifStrpbrk = def.DefFunc("strpbrk", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "char_list"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 2, 2, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 2, 2, 0)
	haystack := fp.ParseString()
	char_list := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrpbrk(executeData.Ctx(), haystack, char_list)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstrCompare
var DefZifSubstrCompare = def.DefFunc("substr_compare", 3, 5, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}, {Name: "length"}, {Name: "case_insensitivity"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 3, 5, 0) {
		return
	}
	fp := php.NewParamParser(executeData, 3, 5, 0)
	haystack := fp.ParseString()
	needle := fp.ParseString()
	offset := fp.ParseLong()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	case_insensitivity := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstrCompare(executeData.Ctx(), returnValue, haystack, needle, offset, nil, length, case_insensitivity)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
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

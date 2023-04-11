package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifBin2hex
var DefZifBin2hex = def.DefFunc("bin2hex", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifBin2hex(data)
	returnValue.SetStringVal(ret)
})

// generate by ZifHex2bin
var DefZifHex2bin = def.DefFunc("hex2bin", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHex2bin(data)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrspn
var DefZifStrspn = def.DefFunc("strspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "offset"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	str := fp.ParseStringVal()
	mask := fp.ParseStringVal()
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
var DefZifStrcspn = def.DefFunc("strcspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "offset"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	str := fp.ParseStringVal()
	mask := fp.ParseStringVal()
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
var DefZifStrcoll = def.DefFunc("strcoll", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrcoll(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifTrim
var DefZifTrim = def.DefFunc("trim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	character_mask := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifTrim(str, nil, character_mask)
	returnValue.SetStringVal(ret)
})

// generate by ZifRtrim
var DefZifRtrim = def.DefFunc("rtrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	character_mask := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(str, nil, character_mask)
	returnValue.SetStringVal(ret)
})

// generate by ZifRtrim
var DefZifChop = def.DefFunc("chop", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	character_mask := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(str, nil, character_mask)
	returnValue.SetStringVal(ret)
})

// generate by ZifLtrim
var DefZifLtrim = def.DefFunc("ltrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	character_mask := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifLtrim(str, nil, character_mask)
	returnValue.SetStringVal(ret)
})

// generate by ZifWordwrap
var DefZifWordwrap = def.DefFunc("wordwrap", 1, 4, []def.ArgInfo{{Name: "str"}, {Name: "width"}, {Name: "break_"}, {Name: "cut"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	width := fp.ParseLongNullable()
	break_ := fp.ParseStringValNullable()
	cut := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifWordwrap(str, nil, width, break_, cut)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifExplode
var DefZifExplode = def.DefFunc("explode", 2, 3, []def.ArgInfo{{Name: "separator"}, {Name: "str"}, {Name: "limit_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	separator := fp.ParseStringVal()
	str := fp.ParseStringVal()
	fp.StartOptional()
	limit_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifExplode(separator, str, nil, limit_)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifImplode
var DefZifImplode = def.DefFunc("implode", 1, 2, []def.ArgInfo{{Name: "glue_"}, {Name: "pieces_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	glue_ := fp.ParseZval()
	fp.StartOptional()
	pieces_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(glue_, nil, pieces_)
	returnValue.SetStringVal(ret)
})

// generate by ZifImplode
var DefZifJoin = def.DefFunc("join", 1, 2, []def.ArgInfo{{Name: "glue_"}, {Name: "pieces_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	glue_ := fp.ParseZval()
	fp.StartOptional()
	pieces_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(glue_, nil, pieces_)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrtok
var DefZifStrtok = def.DefFunc("strtok", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "token_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	token_ := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrtok(str, nil, token_)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrtoupper
var DefZifStrtoupper = def.DefFunc("strtoupper", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrtoupper(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrtolower
var DefZifStrtolower = def.DefFunc("strtolower", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrtolower(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifBasename
var DefZifBasename = def.DefFunc("basename", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "suffix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseStringVal()
	fp.StartOptional()
	suffix := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifBasename(path, nil, suffix)
	returnValue.SetStringVal(ret)
})

// generate by ZifDirname
var DefZifDirname = def.DefFunc("dirname", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "levels_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseStringVal()
	fp.StartOptional()
	levels_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifDirname(path, nil, levels_)
	returnValue.SetStringVal(ret)
})

// generate by ZifPathinfo
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseStringVal()
	fp.StartOptional()
	options := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifPathinfo(path, nil, options)
	returnValue.SetBy(ret)
})

// generate by ZifStristr
var DefZifStristr = def.DefFunc("stristr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStristr(haystack, needle, nil, part)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrstr
var DefZifStrstr = def.DefFunc("strstr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrstr(haystack, needle, nil, part)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrstr
var DefZifStrchr = def.DefFunc("strchr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrstr(haystack, needle, nil, part)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrpos
var DefZifStrpos = def.DefFunc("strpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrpos(haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStripos
var DefZifStripos = def.DefFunc("stripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStripos(haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrpos
var DefZifStrrpos = def.DefFunc("strrpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrrpos(haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrripos
var DefZifStrripos = def.DefFunc("strripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrripos(haystack, needle, nil, offset)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrchr
var DefZifStrrchr = def.DefFunc("strrchr", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrrchr(haystack, needle)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifChunkSplit
var DefZifChunkSplit = def.DefFunc("chunk_split", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "chunklen_"}, {Name: "ending_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	chunklen_ := fp.ParseLongNullable()
	ending_ := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifChunkSplit(str, nil, chunklen_, ending_)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstr
var DefZifSubstr = def.DefFunc("substr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "offset"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str := fp.ParseStringVal()
	offset := fp.ParseLong()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstr(str, offset, nil, length)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstrReplace
var DefZifSubstrReplace = def.DefFunc("substr_replace", 3, 4, []def.ArgInfo{{Name: "str"}, {Name: "replace"}, {Name: "start"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 4, 0)
	str := fp.ParseZval()
	replace := fp.ParseZval()
	start := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSubstrReplace(returnValue, str, replace, start, nil, length)
})

// generate by ZifQuotemeta
var DefZifQuotemeta = def.DefFunc("quotemeta", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifQuotemeta(str)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifOrd
var DefZifOrd = def.DefFunc("ord", 1, 1, []def.ArgInfo{{Name: "character"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	character := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifOrd(character)
	returnValue.SetLong(ret)
})

// generate by ZifChr
var DefZifChr = def.DefFunc("chr", 1, 1, []def.ArgInfo{{Name: "codepoint"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	codepoint := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifChr(codepoint)
	returnValue.SetStringVal(ret)
})

// generate by ZifUcfirst
var DefZifUcfirst = def.DefFunc("ucfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUcfirst(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifLcfirst
var DefZifLcfirst = def.DefFunc("lcfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifLcfirst(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifUcwords
var DefZifUcwords = def.DefFunc("ucwords", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "delimiters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	delimiters := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifUcwords(str, nil, delimiters)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrtr
var DefZifStrtr = def.DefFunc("strtr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "from"}, {Name: "to_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str := fp.ParseStringVal()
	from := fp.ParseZval()
	fp.StartOptional()
	to_ := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrtr(str, from, nil, to_)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrrev
var DefZifStrrev = def.DefFunc("strrev", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrrev(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifSimilarText
var DefZifSimilarText = def.DefFunc("similar_text", 2, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "percent"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	fp.StartOptional()
	percent := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifSimilarText(str1, str2, nil, percent)
	returnValue.SetLong(ret)
})

// generate by ZifAddslashes
var DefZifAddslashes = def.DefFunc("addslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifAddslashes(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifAddcslashes
var DefZifAddcslashes = def.DefFunc("addcslashes", 2, 2, []def.ArgInfo{{Name: "str"}, {Name: "charlist"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str := fp.ParseStringVal()
	charlist := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifAddcslashes(str, charlist)
	returnValue.SetStringVal(ret)
})

// generate by ZifStripslashes
var DefZifStripslashes = def.DefFunc("stripslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStripslashes(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifStripcslashes
var DefZifStripcslashes = def.DefFunc("stripcslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStripcslashes(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrReplace
var DefZifStrReplace = def.DefFunc("str_replace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 4, 0)
	search := fp.ParseZval()
	replace := fp.ParseZval()
	subject := fp.ParseZval()
	fp.StartOptional()
	replace_count := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifStrReplace(returnValue, search, replace, subject, nil, replace_count)
})

// generate by ZifStrIreplace
var DefZifStrIreplace = def.DefFunc("str_ireplace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 4, 0)
	search := fp.ParseZval()
	replace := fp.ParseZval()
	subject := fp.ParseZval()
	fp.StartOptional()
	replace_count := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifStrIreplace(returnValue, search, replace, subject, nil, replace_count)
})

// generate by ZifHebrev
var DefZifHebrev = def.DefFunc("hebrev", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	max_chars_per_line := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHebrev(str, nil, max_chars_per_line)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifHebrevc
var DefZifHebrevc = def.DefFunc("hebrevc", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	max_chars_per_line := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHebrevc(str, nil, max_chars_per_line)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifNl2br
var DefZifNl2br = def.DefFunc("nl2br", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "is_xhtml_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	is_xhtml_ := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifNl2br(str, nil, is_xhtml_)
	returnValue.SetStringVal(ret)
})

// generate by ZifStripTags
var DefZifStripTags = def.DefFunc("strip_tags", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "allowable_tags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	allowable_tags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifStripTags(str, nil, allowable_tags)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrRepeat
var DefZifStrRepeat = def.DefFunc("str_repeat", 2, 2, []def.ArgInfo{{Name: "input"}, {Name: "mult"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	input := fp.ParseStringVal()
	mult := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrRepeat(input, mult)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCountChars
var DefZifCountChars = def.DefFunc("count_chars", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseStringVal()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCountChars(input, nil, mode)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrnatcmp
var DefZifStrnatcmp = def.DefFunc("strnatcmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	s1 := fp.ParseStringVal()
	s2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrnatcmp(s1, s2)
	returnValue.SetLong(ret)
})

// generate by ZifStrnatcasecmp
var DefZifStrnatcasecmp = def.DefFunc("strnatcasecmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	s1 := fp.ParseStringVal()
	s2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrnatcasecmp(s1, s2)
	returnValue.SetLong(ret)
})

// generate by ZifSubstrCount
var DefZifSubstrCount = def.DefFunc("substr_count", 2, 4, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}, {Name: "length_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseStringVal()
	fp.StartOptional()
	offset := fp.ParseLong()
	length_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstrCount(haystack, needle, nil, offset, length_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrPad
var DefZifStrPad = def.DefFunc("str_pad", 2, 4, []def.ArgInfo{{Name: "input"}, {Name: "pad_length"}, {Name: "pad_string_"}, {Name: "pad_type_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	input := fp.ParseStringVal()
	pad_length := fp.ParseLong()
	fp.StartOptional()
	pad_string_ := fp.ParseStringValNullable()
	pad_type_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrPad(input, pad_length, nil, pad_string_, pad_type_)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrRot13
var DefZifStrRot13 = def.DefFunc("str_rot13", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrRot13(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrShuffle
var DefZifStrShuffle = def.DefFunc("str_shuffle", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrShuffle(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifStrWordCount
var DefZifStrWordCount = def.DefFunc("str_word_count", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "format"}, {Name: "charlist"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	format := fp.ParseLong()
	charlist := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrWordCount(str, nil, format, charlist)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrSplit
var DefZifStrSplit = def.DefFunc("str_split", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "split_length_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	split_length_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrSplit(str, nil, split_length_)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrpbrk
var DefZifStrpbrk = def.DefFunc("strpbrk", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "char_list"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	haystack := fp.ParseStringVal()
	char_list := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrpbrk(haystack, char_list)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSubstrCompare
var DefZifSubstrCompare = def.DefFunc("substr_compare", 3, 5, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}, {Name: "length"}, {Name: "case_insensitivity"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 5, 0)
	haystack := fp.ParseStringVal()
	needle := fp.ParseStringVal()
	offset := fp.ParseLong()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	case_insensitivity := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSubstrCompare(returnValue, haystack, needle, offset, nil, length, case_insensitivity)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifBin2hex
var DefZifBin2hex = def.DefFunc("bin2hex", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBin2hex(executeData, returnValue, data)
})

// generate by ZifHex2bin
var DefZifHex2bin = def.DefFunc("hex2bin", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHex2bin(executeData, returnValue, data)
})

// generate by ZifStrspn
var DefZifStrspn = def.DefFunc("strspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "start"}, {Name: "len_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	str := fp.ParseZval()
	mask := fp.ParseZval()
	fp.StartOptional()
	start := fp.ParseZval()
	len_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrspn(executeData, returnValue, str, mask, nil, start, len_)
})

// generate by ZifStrcspn
var DefZifStrcspn = def.DefFunc("strcspn", 2, 4, []def.ArgInfo{{Name: "str"}, {Name: "mask"}, {Name: "start"}, {Name: "len_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	str := fp.ParseZval()
	mask := fp.ParseZval()
	fp.StartOptional()
	start := fp.ParseZval()
	len_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrcspn(executeData, returnValue, str, mask, nil, start, len_)
})

// generate by ZifStrcoll
var DefZifStrcoll = def.DefFunc("strcoll", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseZval()
	str2 := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrcoll(executeData, returnValue, str1, str2)
})

// generate by ZifTrim
var DefZifTrim = def.DefFunc("trim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	character_mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTrim(executeData, returnValue, str, nil, character_mask)
})

// generate by ZifRtrim
var DefZifRtrim = def.DefFunc("rtrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	character_mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRtrim(executeData, returnValue, str, nil, character_mask)
})

// generate by ZifRtrim
var DefZifChop = def.DefFunc("chop", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	character_mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRtrim(executeData, returnValue, str, nil, character_mask)
})

// generate by ZifLtrim
var DefZifLtrim = def.DefFunc("ltrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	character_mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLtrim(executeData, returnValue, str, nil, character_mask)
})

// generate by ZifWordwrap
var DefZifWordwrap = def.DefFunc("wordwrap", 1, 4, []def.ArgInfo{{Name: "str"}, {Name: "width"}, {Name: "break_"}, {Name: "cut"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	width := fp.ParseZval()
	break_ := fp.ParseZval()
	cut := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifWordwrap(executeData, returnValue, str, nil, width, break_, cut)
})

// generate by ZifExplode
var DefZifExplode = def.DefFunc("explode", 2, 3, []def.ArgInfo{{Name: "separator"}, {Name: "str"}, {Name: "limit"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	separator := fp.ParseZval()
	str := fp.ParseZval()
	fp.StartOptional()
	limit := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifExplode(executeData, returnValue, separator, str, nil, limit)
})

// generate by ZifImplode
var DefZifImplode = def.DefFunc("implode", 2, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	glue := fp.ParseZval()
	pieces := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImplode(executeData, returnValue, glue, pieces)
})

// generate by ZifImplode
var DefZifJoin = def.DefFunc("join", 2, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	glue := fp.ParseZval()
	pieces := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImplode(executeData, returnValue, glue, pieces)
})

// generate by ZifStrtok
var DefZifStrtok = def.DefFunc("strtok", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "token"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	token := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrtok(executeData, returnValue, str, nil, token)
})

// generate by ZifStrtoupper
var DefZifStrtoupper = def.DefFunc("strtoupper", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrtoupper(executeData, returnValue, str)
})

// generate by ZifStrtolower
var DefZifStrtolower = def.DefFunc("strtolower", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrtolower(executeData, returnValue, str)
})

// generate by ZifBasename
var DefZifBasename = def.DefFunc("basename", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "suffix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	suffix := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBasename(executeData, returnValue, path, nil, suffix)
})

// generate by ZifDirname
var DefZifDirname = def.DefFunc("dirname", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "levels"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	levels := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDirname(executeData, returnValue, path, nil, levels)
})

// generate by ZifPathinfo
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPathinfo(executeData, returnValue, path, nil, options)
})

// generate by ZifStristr
var DefZifStristr = def.DefFunc("stristr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStristr(executeData, returnValue, haystack, needle, nil, part)
})

// generate by ZifStrstr
var DefZifStrstr = def.DefFunc("strstr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrstr(executeData, returnValue, haystack, needle, nil, part)
})

// generate by ZifStrstr
var DefZifStrchr = def.DefFunc("strchr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	part := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrstr(executeData, returnValue, haystack, needle, nil, part)
})

// generate by ZifStrpos
var DefZifStrpos = def.DefFunc("strpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrpos(executeData, returnValue, haystack, needle, nil, offset)
})

// generate by ZifStripos
var DefZifStripos = def.DefFunc("stripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStripos(executeData, returnValue, haystack, needle, nil, offset)
})

// generate by ZifStrrpos
var DefZifStrrpos = def.DefFunc("strrpos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrrpos(executeData, returnValue, haystack, needle, nil, offset)
})

// generate by ZifStrripos
var DefZifStrripos = def.DefFunc("strripos", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrripos(executeData, returnValue, haystack, needle, nil, offset)
})

// generate by ZifStrrchr
var DefZifStrrchr = def.DefFunc("strrchr", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrrchr(executeData, returnValue, haystack, needle)
})

// generate by ZifChunkSplit
var DefZifChunkSplit = def.DefFunc("chunk_split", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "chunklen"}, {Name: "ending"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	chunklen := fp.ParseZval()
	ending := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChunkSplit(executeData, returnValue, str, nil, chunklen, ending)
})

// generate by ZifSubstr
var DefZifSubstr = def.DefFunc("substr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "start"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str := fp.ParseZval()
	start := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSubstr(executeData, returnValue, str, start, nil, length)
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
	ZifSubstrReplace(executeData, returnValue, str, replace, start, nil, length)
})

// generate by ZifQuotemeta
var DefZifQuotemeta = def.DefFunc("quotemeta", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifQuotemeta(executeData, returnValue, str)
})

// generate by ZifOrd
var DefZifOrd = def.DefFunc("ord", 1, 1, []def.ArgInfo{{Name: "character"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	character := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOrd(executeData, returnValue, character)
})

// generate by ZifChr
var DefZifChr = def.DefFunc("chr", 1, 1, []def.ArgInfo{{Name: "codepoint"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	codepoint := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChr(executeData, returnValue, codepoint)
})

// generate by ZifUcfirst
var DefZifUcfirst = def.DefFunc("ucfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUcfirst(executeData, returnValue, str)
})

// generate by ZifLcfirst
var DefZifLcfirst = def.DefFunc("lcfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLcfirst(executeData, returnValue, str)
})

// generate by ZifUcwords
var DefZifUcwords = def.DefFunc("ucwords", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "delimiters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	delimiters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUcwords(executeData, returnValue, str, nil, delimiters)
})

// generate by ZifStrtr
var DefZifStrtr = def.DefFunc("strtr", 2, 3, []def.ArgInfo{{Name: "str"}, {Name: "from"}, {Name: "to"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str := fp.ParseZval()
	from := fp.ParseZval()
	fp.StartOptional()
	to := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrtr(executeData, returnValue, str, from, nil, to)
})

// generate by ZifStrrev
var DefZifStrrev = def.DefFunc("strrev", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrrev(executeData, returnValue, str)
})

// generate by ZifSimilarText
var DefZifSimilarText = def.DefFunc("similar_text", 2, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "percent"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	str1 := fp.ParseZval()
	str2 := fp.ParseZval()
	fp.StartOptional()
	percent := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifSimilarText(executeData, returnValue, str1, str2, nil, percent)
})

// generate by ZifAddcslashes
var DefZifAddcslashes = def.DefFunc("addcslashes", 2, 2, []def.ArgInfo{{Name: "str"}, {Name: "charlist"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str := fp.ParseZval()
	charlist := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAddcslashes(executeData, returnValue, str, charlist)
})

// generate by ZifAddslashes
var DefZifAddslashes = def.DefFunc("addslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAddslashes(executeData, returnValue, str)
})

// generate by ZifStripcslashes
var DefZifStripcslashes = def.DefFunc("stripcslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStripcslashes(executeData, returnValue, str)
})

// generate by ZifStripslashes
var DefZifStripslashes = def.DefFunc("stripslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStripslashes(executeData, returnValue, str)
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
	ZifStrReplace(executeData, returnValue, search, replace, subject, nil, replace_count)
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
	ZifStrIreplace(executeData, returnValue, search, replace, subject, nil, replace_count)
})

// generate by ZifHebrev
var DefZifHebrev = def.DefFunc("hebrev", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	max_chars_per_line := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHebrev(executeData, returnValue, str, nil, max_chars_per_line)
})

// generate by ZifHebrevc
var DefZifHebrevc = def.DefFunc("hebrevc", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "max_chars_per_line"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	max_chars_per_line := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHebrevc(executeData, returnValue, str, nil, max_chars_per_line)
})

// generate by ZifNl2br
var DefZifNl2br = def.DefFunc("nl2br", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "is_xhtml"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	is_xhtml := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifNl2br(executeData, returnValue, str, nil, is_xhtml)
})

// generate by ZifStripTags
var DefZifStripTags = def.DefFunc("strip_tags", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "allowable_tags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	allowable_tags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStripTags(executeData, returnValue, str, nil, allowable_tags)
})

// generate by ZifSetlocale
var DefZifSetlocale = def.DefFunc("setlocale", 1, -1, []def.ArgInfo{{Name: "category"}, {Name: "locales"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	category := fp.ParseZval()
	locales := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifSetlocale(executeData, returnValue, category, locales)
})

// generate by ZifParseStr
var DefZifParseStr = def.DefFunc("parse_str", 1, 2, []def.ArgInfo{{Name: "encoded_string"}, {Name: "result"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	encoded_string := fp.ParseZval()
	fp.StartOptional()
	result := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifParseStr(executeData, returnValue, encoded_string, nil, result)
})

// generate by ZifStrGetcsv
var DefZifStrGetcsv = def.DefFunc("str_getcsv", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	delimiter := fp.ParseZval()
	enclosure := fp.ParseZval()
	escape := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrGetcsv(executeData, returnValue, string, nil, delimiter, enclosure, escape)
})

// generate by ZifStrRepeat
var DefZifStrRepeat = def.DefFunc("str_repeat", 2, 2, []def.ArgInfo{{Name: "input"}, {Name: "mult"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	input := fp.ParseZval()
	mult := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrRepeat(executeData, returnValue, input, mult)
})

// generate by ZifCountChars
var DefZifCountChars = def.DefFunc("count_chars", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCountChars(executeData, returnValue, input, nil, mode)
})

// generate by ZifStrnatcmp
var DefZifStrnatcmp = def.DefFunc("strnatcmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	s1 := fp.ParseZval()
	s2 := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrnatcmp(executeData, returnValue, s1, s2)
})

// generate by ZifLocaleconv
var DefZifLocaleconv = def.DefFunc("localeconv", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifLocaleconv(executeData, returnValue)
})

// generate by ZifStrnatcasecmp
var DefZifStrnatcasecmp = def.DefFunc("strnatcasecmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	s1 := fp.ParseZval()
	s2 := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrnatcasecmp(executeData, returnValue, s1, s2)
})

// generate by ZifSubstrCount
var DefZifSubstrCount = def.DefFunc("substr_count", 2, 4, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "offset"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	haystack := fp.ParseZval()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSubstrCount(executeData, returnValue, haystack, needle, nil, offset, length)
})

// generate by ZifStrPad
var DefZifStrPad = def.DefFunc("str_pad", 2, 4, []def.ArgInfo{{Name: "input"}, {Name: "pad_length"}, {Name: "pad_string"}, {Name: "pad_type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	input := fp.ParseZval()
	pad_length := fp.ParseZval()
	fp.StartOptional()
	pad_string := fp.ParseZval()
	pad_type := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrPad(executeData, returnValue, input, pad_length, nil, pad_string, pad_type)
})

// generate by ZifStrRot13
var DefZifStrRot13 = def.DefFunc("str_rot13", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrRot13(executeData, returnValue, str)
})

// generate by ZifStrShuffle
var DefZifStrShuffle = def.DefFunc("str_shuffle", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrShuffle(executeData, returnValue, str)
})

// generate by ZifStrWordCount
var DefZifStrWordCount = def.DefFunc("str_word_count", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "format"}, {Name: "charlist"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	format := fp.ParseZval()
	charlist := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrWordCount(executeData, returnValue, str, nil, format, charlist)
})

// generate by ZifMoneyFormat
var DefZifMoneyFormat = def.DefFunc("money_format", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	format := fp.ParseZval()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMoneyFormat(executeData, returnValue, format, value)
})

// generate by ZifStrSplit
var DefZifStrSplit = def.DefFunc("str_split", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "split_length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	split_length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrSplit(executeData, returnValue, str, nil, split_length)
})

// generate by ZifStrpbrk
var DefZifStrpbrk = def.DefFunc("strpbrk", 2, 2, []def.ArgInfo{{Name: "haystack"}, {Name: "char_list"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	haystack := fp.ParseZval()
	char_list := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrpbrk(executeData, returnValue, haystack, char_list)
})

// generate by ZifSubstrCompare
var DefZifSubstrCompare = def.DefFunc("substr_compare", 3, 5, []def.ArgInfo{{Name: "main_str"}, {Name: "str"}, {Name: "offset"}, {Name: "length"}, {Name: "case_sensitivity"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 5, 0)
	main_str := fp.ParseZval()
	str := fp.ParseZval()
	offset := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	case_sensitivity := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSubstrCompare(executeData, returnValue, main_str, str, offset, nil, length, case_sensitivity)
})

// generate by ZifUtf8Encode
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUtf8Encode(executeData, returnValue, data)
})

// generate by ZifUtf8Decode
var DefZifUtf8Decode = def.DefFunc("utf8_decode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUtf8Decode(executeData, returnValue, data)
})

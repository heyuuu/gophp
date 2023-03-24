package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifStrspn
var DefZifStrspn = def.DefFunc("strspn", 2, 4, []def.ArgInfo{{name: "str"}, {name: "mask"}, {name: "start"}, {name: "len_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrcspn = def.DefFunc("strcspn", 2, 4, []def.ArgInfo{{name: "str"}, {name: "mask"}, {name: "start"}, {name: "len_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifTrim
var DefZifTrim = def.DefFunc("trim", 1, 2, []def.ArgInfo{{name: "str"}, {name: "character_mask"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifRtrim = def.DefFunc("rtrim", 1, 2, []def.ArgInfo{{name: "str"}, {name: "character_mask"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifLtrim = def.DefFunc("ltrim", 1, 2, []def.ArgInfo{{name: "str"}, {name: "character_mask"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifWordwrap = def.DefFunc("wordwrap", 1, 4, []def.ArgInfo{{name: "str"}, {name: "width"}, {name: "break_"}, {name: "cut"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifExplode = def.DefFunc("explode", 2, 3, []def.ArgInfo{{name: "separator"}, {name: "str"}, {name: "limit"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifStrtok
var DefZifStrtok = def.DefFunc("strtok", 1, 2, []def.ArgInfo{{name: "str"}, {name: "token"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	token := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrtok(executeData, returnValue, str, nil, token)
})

// generate by ZifBasename
var DefZifBasename = def.DefFunc("basename", 1, 2, []def.ArgInfo{{name: "path"}, {name: "suffix"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifDirname = def.DefFunc("dirname", 1, 2, []def.ArgInfo{{name: "path"}, {name: "levels"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{name: "path"}, {name: "options"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStristr = def.DefFunc("stristr", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "part"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrstr = def.DefFunc("strstr", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "part"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrpos = def.DefFunc("strpos", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "offset"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStripos = def.DefFunc("stripos", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "offset"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrrpos = def.DefFunc("strrpos", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "offset"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrripos = def.DefFunc("strripos", 2, 3, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "offset"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifChunkSplit
var DefZifChunkSplit = def.DefFunc("chunk_split", 1, 3, []def.ArgInfo{{name: "str"}, {name: "chunklen"}, {name: "ending"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifSubstr = def.DefFunc("substr", 2, 3, []def.ArgInfo{{name: "str"}, {name: "start"}, {name: "length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifSubstrReplace = def.DefFunc("substr_replace", 3, 4, []def.ArgInfo{{name: "str"}, {name: "replace"}, {name: "start"}, {name: "length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifUcwords
var DefZifUcwords = def.DefFunc("ucwords", 1, 2, []def.ArgInfo{{name: "str"}, {name: "delimiters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrtr = def.DefFunc("strtr", 2, 3, []def.ArgInfo{{name: "str"}, {name: "from"}, {name: "to"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifSimilarText
var DefZifSimilarText = def.DefFunc("similar_text", 2, 3, []def.ArgInfo{{name: "str1"}, {name: "str2"}, {name: "percent"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifStrReplace
var DefZifStrReplace = def.DefFunc("str_replace", 3, 4, []def.ArgInfo{{name: "search"}, {name: "replace"}, {name: "subject"}, {name: "replace_count"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrIreplace = def.DefFunc("str_ireplace", 3, 4, []def.ArgInfo{{name: "search"}, {name: "replace"}, {name: "subject"}, {name: "replace_count"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifHebrev = def.DefFunc("hebrev", 1, 2, []def.ArgInfo{{name: "str"}, {name: "max_chars_per_line"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifHebrevc = def.DefFunc("hebrevc", 1, 2, []def.ArgInfo{{name: "str"}, {name: "max_chars_per_line"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifNl2br = def.DefFunc("nl2br", 1, 2, []def.ArgInfo{{name: "str"}, {name: "is_xhtml"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStripTags = def.DefFunc("strip_tags", 1, 2, []def.ArgInfo{{name: "str"}, {name: "allowable_tags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifSetlocale = def.DefFunc("setlocale", -1, -1, []def.ArgInfo{{name: "category"}, {name: "locales"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	category := fp.ParseZval()
	locales := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifSetlocale(executeData, returnValue, category, locales)
})

// generate by ZifParseStr
var DefZifParseStr = def.DefFunc("parse_str", 1, 2, []def.ArgInfo{{name: "encoded_string"}, {name: "result"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrGetcsv = def.DefFunc("str_getcsv", 1, 4, []def.ArgInfo{{name: "string"}, {name: "delimiter"}, {name: "enclosure"}, {name: "escape"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifCountChars
var DefZifCountChars = def.DefFunc("count_chars", 1, 2, []def.ArgInfo{{name: "input"}, {name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCountChars(executeData, returnValue, input, nil, mode)
})

// generate by ZifSubstrCount
var DefZifSubstrCount = def.DefFunc("substr_count", 2, 4, []def.ArgInfo{{name: "haystack"}, {name: "needle"}, {name: "offset"}, {name: "length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrPad = def.DefFunc("str_pad", 2, 4, []def.ArgInfo{{name: "input"}, {name: "pad_length"}, {name: "pad_string"}, {name: "pad_type"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifStrWordCount
var DefZifStrWordCount = def.DefFunc("str_word_count", 1, 3, []def.ArgInfo{{name: "str"}, {name: "format"}, {name: "charlist"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifStrSplit
var DefZifStrSplit = def.DefFunc("str_split", 1, 2, []def.ArgInfo{{name: "str"}, {name: "split_length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifStrpbrk = def.DefFunc("strpbrk", 2, 2, []def.ArgInfo{{name: "haystack"}, {name: "char_list"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	haystack := fp.ParseZval()
	char_list := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrpbrk(executeData, returnValue, haystack, char_list)
})

// generate by ZifSubstrCompare
var DefZifSubstrCompare = def.DefFunc("substr_compare", 3, 5, []def.ArgInfo{{name: "main_str"}, {name: "str"}, {name: "offset"}, {name: "length"}, {name: "case_sensitivity"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{name: "data"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUtf8Encode(executeData, returnValue, data)
})

// generate by ZifUtf8Decode
var DefZifUtf8Decode = def.DefFunc("utf8_decode", 1, 1, []def.ArgInfo{{name: "data"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUtf8Decode(executeData, returnValue, data)
})

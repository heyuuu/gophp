package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifKrsort, DefZifKsort, DefZifCount, DefZifSizeof, DefZifNatsort, DefZifNatcasesort, DefZifAsort, DefZifArsort, DefZifSort, DefZifRsort, DefZifUsort, DefZifUasort, DefZifUksort, DefZifEnd, DefZifPrev, DefZifNext, DefZifReset, DefZifCurrent, DefZifPos, DefZifKey, DefZifMin, DefZifMax, DefZifArrayWalk, DefZifArrayWalkRecursive, DefZifInArray, DefZifArraySearch, DefZifArrayFill, DefZifArrayFillKeys, DefZifRange, DefZifShuffle, DefZifArrayPush, DefZifArrayPop, DefZifArrayShift, DefZifArrayUnshift, DefZifArraySplice, DefZifArraySlice, DefZifArrayMerge, DefZifArrayMergeRecursive, DefZifArrayReplace, DefZifArrayReplaceRecursive, DefZifArrayKeys, DefZifArrayKeyFirst, DefZifArrayKeyLast, DefZifArrayValues, DefZifArrayCountValues, DefZifArrayColumn, DefZifArrayReverse, DefZifArrayPad, DefZifArrayFlip, DefZifArrayChangeKeyCase, DefZifArrayUnique, DefZifArrayIntersectKey, DefZifArrayIntersectUkey, DefZifArrayIntersect, DefZifArrayUintersect, DefZifArrayIntersectAssoc, DefZifArrayIntersectUassoc, DefZifArrayUintersectAssoc, DefZifArrayUintersectUassoc, DefZifArrayDiffKey, DefZifArrayDiffUkey, DefZifArrayDiff, DefZifArrayUdiff, DefZifArrayDiffAssoc, DefZifArrayDiffUassoc, DefZifArrayUdiffAssoc, DefZifArrayUdiffUassoc, DefZifArrayMultisort, DefZifArrayRand, DefZifArraySum, DefZifArrayProduct, DefZifArrayReduce, DefZifArrayFilter, DefZifArrayMap, DefZifArrayKeyExists, DefZifKeyExists, DefZifArrayChunk, DefZifArrayCombine, DefZifBase64Encode, DefZifBase64Decode, DefZifInetNtop, DefZifInetPton, DefZifIp2long, DefZifLong2ip, DefZifIniGet, DefZifIniGetAll, DefZifIniSet, DefZifIniAlter, DefZifIniRestore, DefZifSetIncludePath, DefZifGetIncludePath, DefZifRestoreIncludePath, DefZifPrintR, DefZifParseIniFile, DefZifParseIniString, DefZifCrc32, DefZifStrGetcsv, DefZifGethostname, DefZifGethostbyaddr, DefZifGethostbyname, DefZifGethostbynamel, DefZifDnsGetMx, DefZifGetmxrr, DefZifSprintf, DefZifVsprintf, DefZifPrintf, DefZifVprintf, DefZifHtmlspecialchars, DefZifHtmlspecialcharsDecode, DefZifHtmlEntityDecode, DefZifHtmlentities, DefZifGetHtmlTranslationTable, DefZifHttpBuildQuery, DefZifAbs, DefZifCeil, DefZifFloor, DefZifRound, DefZifSin, DefZifCos, DefZifTan, DefZifAsin, DefZifAcos, DefZifAtan, DefZifAtan2, DefZifSinh, DefZifCosh, DefZifTanh, DefZifAsinh, DefZifAcosh, DefZifAtanh, DefZifPi, DefZifIsFinite, DefZifIsInfinite, DefZifIsNan, DefZifPow, DefZifExp, DefZifExpm1, DefZifLog1p, DefZifLog, DefZifLog10, DefZifSqrt, DefZifHypot, DefZifDeg2rad, DefZifRad2deg, DefZifBindec, DefZifHexdec, DefZifOctdec, DefZifDecbin, DefZifDecoct, DefZifDechex, DefZifBaseConvert, DefZifNumberFormat, DefZifFmod, DefZifIntdiv, DefZifMd5, DefZifMd5File, DefZifMicrotime, DefZifGettimeofday, DefZifGetrusage, DefZifNetGetInterfaces, DefZifPack, DefZifUnpack, DefZifQuotedPrintableDecode, DefZifQuotedPrintableEncode, DefZifMtSrand, DefZifSrand, DefZifMtGetrandmax, DefZifGetrandmax, DefZifMtRand, DefZifRand, DefZifRandomBytes, DefZifRandomInt, DefZifSha1, DefZifSha1File, DefZifSoundex, DefZifUtf8Encode, DefZifUtf8Decode, DefZifBin2hex, DefZifHex2bin, DefZifStrspn, DefZifStrcspn, DefZifStrcoll, DefZifTrim, DefZifRtrim, DefZifChop, DefZifLtrim, DefZifWordwrap, DefZifExplode, DefZifImplode, DefZifJoin, DefZifStrtok, DefZifStrtoupper, DefZifStrtolower, DefZifBasename, DefZifDirname, DefZifPathinfo, DefZifStristr, DefZifStrstr, DefZifStrchr, DefZifStrpos, DefZifStripos, DefZifStrrpos, DefZifStrripos, DefZifStrrchr, DefZifChunkSplit, DefZifSubstr, DefZifSubstrReplace, DefZifQuotemeta, DefZifOrd, DefZifChr, DefZifUcfirst, DefZifLcfirst, DefZifUcwords, DefZifStrtr, DefZifStrrev, DefZifSimilarText, DefZifAddslashes, DefZifAddcslashes, DefZifStripslashes, DefZifStripcslashes, DefZifStrReplace, DefZifStrIreplace, DefZifHebrev, DefZifHebrevc, DefZifNl2br, DefZifStripTags, DefZifStrRepeat, DefZifCountChars, DefZifStrnatcmp, DefZifStrnatcasecmp, DefZifSubstrCount, DefZifStrPad, DefZifStrRot13, DefZifStrShuffle, DefZifStrWordCount, DefZifStrSplit, DefZifStrpbrk, DefZifSubstrCompare, DefZifGettype, DefZifSettype, DefZifIntval, DefZifFloatval, DefZifDoubleval, DefZifBoolval, DefZifStrval, DefZifIsNull, DefZifIsResource, DefZifIsBool, DefZifIsInt, DefZifIsInteger, DefZifIsLong, DefZifIsFloat, DefZifIsDouble, DefZifIsString, DefZifIsArray, DefZifIsObject, DefZifIsNumeric, DefZifIsScalar, DefZifUniqid, DefZifParseUrl, DefZifUrlencode, DefZifUrldecode, DefZifRawurlencode, DefZifRawurldecode, DefZifConvertUuencode, DefZifConvertUudecode, DefZifVarDump, DefZifDebugZvalDump, DefZifVarExport, DefZifSerialize, DefZifVersionCompare}

// generate by ZifKrsort
var DefZifKrsort = def.DefFunc("krsort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifKrsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifKsort
var DefZifKsort = def.DefFunc("ksort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifKsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifCount
var DefZifCount = def.DefFunc("count", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifCount(executeData.Ctx(), var_, nil, mode)
	returnValue.SetLong(ret)
})

// generate by ZifCount
var DefZifSizeof = def.DefFunc("sizeof", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifCount(executeData.Ctx(), var_, nil, mode)
	returnValue.SetLong(ret)
})

// generate by ZifNatsort
var DefZifNatsort = def.DefFunc("natsort", 1, 1, []def.ArgInfo{{Name: "arg", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifNatsort(executeData.Ctx(), arg)
	returnValue.SetBool(ret)
})

// generate by ZifNatcasesort
var DefZifNatcasesort = def.DefFunc("natcasesort", 1, 1, []def.ArgInfo{{Name: "arg", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifNatcasesort(executeData.Ctx(), arg)
	returnValue.SetBool(ret)
})

// generate by ZifAsort
var DefZifAsort = def.DefFunc("asort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifAsort(executeData.Ctx(), arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifArsort
var DefZifArsort = def.DefFunc("arsort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifArsort(executeData.Ctx(), arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifSort
var DefZifSort = def.DefFunc("sort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifSort(executeData.Ctx(), arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifRsort
var DefZifRsort = def.DefFunc("rsort", 1, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "sort_flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifRsort(executeData.Ctx(), arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifUsort
var DefZifUsort = def.DefFunc("usort", 2, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "cmp_function"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUsort(executeData.Ctx(), arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifUasort
var DefZifUasort = def.DefFunc("uasort", 2, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "cmp_function"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUasort(executeData.Ctx(), arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifUksort
var DefZifUksort = def.DefFunc("uksort", 2, 2, []def.ArgInfo{{Name: "arg", ByRef: true}, {Name: "cmp_function"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUksort(executeData.Ctx(), arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifEnd
var DefZifEnd = def.DefFunc("end", 1, 1, []def.ArgInfo{{Name: "array", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifEnd(array)
	returnValue.SetBy(ret)
})

// generate by ZifPrev
var DefZifPrev = def.DefFunc("prev", 1, 1, []def.ArgInfo{{Name: "array", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifPrev(array)
	returnValue.SetBy(ret)
})

// generate by ZifNext
var DefZifNext = def.DefFunc("next", 1, 1, []def.ArgInfo{{Name: "array", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifNext(array)
	returnValue.SetBy(ret)
})

// generate by ZifReset
var DefZifReset = def.DefFunc("reset", 1, 1, []def.ArgInfo{{Name: "array", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifReset(array)
	returnValue.SetBy(ret)
})

// generate by ZifCurrent
var DefZifCurrent = def.DefFunc("current", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCurrent(array)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCurrent
var DefZifPos = def.DefFunc("pos", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCurrent(array)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifKey
var DefZifKey = def.DefFunc("key", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifKey(array)
	returnValue.SetBy(ret)
})

// generate by ZifMin
var DefZifMin = def.DefFunc("min", 1, -1, []def.ArgInfo{{Name: "arg"}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arg := fp.ParseZval()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifMin(executeData.Ctx(), arg, args)
	returnValue.SetBy(ret)
})

// generate by ZifMax
var DefZifMax = def.DefFunc("max", 1, -1, []def.ArgInfo{{Name: "arg"}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arg := fp.ParseZval()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifMax(executeData.Ctx(), arg, args)
	returnValue.SetBy(ret)
})

// generate by ZifArrayWalk
var DefZifArrayWalk = def.DefFunc("array_walk", 2, 3, []def.ArgInfo{{Name: "array", ByRef: true}, {Name: "callable"}, {Name: "arg"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArrayOrObject()
	callable := fp.ParseCallable()
	fp.StartOptional()
	arg := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayWalk(executeData.Ctx(), array, callable, nil, arg)
	returnValue.SetBool(ret)
})

// generate by ZifArrayWalkRecursive
var DefZifArrayWalkRecursive = def.DefFunc("array_walk_recursive", 2, 3, []def.ArgInfo{{Name: "array", ByRef: true}, {Name: "callable"}, {Name: "arg"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArrayOrObject()
	callable := fp.ParseCallable()
	fp.StartOptional()
	arg := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayWalkRecursive(executeData.Ctx(), array, callable, nil, arg)
	returnValue.SetBool(ret)
})

// generate by ZifInArray
var DefZifInArray = def.DefFunc("in_array", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	needle := fp.ParseZval()
	haystack := fp.ParseArray()
	fp.StartOptional()
	strict := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifInArray(executeData.Ctx(), needle, haystack, nil, strict)
	returnValue.SetBool(ret)
})

// generate by ZifArraySearch
var DefZifArraySearch = def.DefFunc("array_search", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	needle := fp.ParseZval()
	haystack := fp.ParseArray()
	fp.StartOptional()
	strict := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifArraySearch(executeData.Ctx(), needle, haystack, nil, strict)
	returnValue.SetBy(ret)
})

// generate by ZifArrayFill
var DefZifArrayFill = def.DefFunc("array_fill", 3, 3, []def.ArgInfo{{Name: "start_key"}, {Name: "num"}, {Name: "val"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 3, 0)
	fp.CheckNumArgs()
	start_key := fp.ParseLong()
	num := fp.ParseLong()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayFill(executeData.Ctx(), start_key, num, val)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayFillKeys
var DefZifArrayFillKeys = def.DefFunc("array_fill_keys", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "val"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	keys := fp.ParseArray()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFillKeys(executeData.Ctx(), keys, val)
	returnValue.SetArray(ret)
})

// generate by ZifRange
var DefZifRange = def.DefFunc("range", 2, 3, []def.ArgInfo{{Name: "low"}, {Name: "high"}, {Name: "step"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	low := fp.ParseZval()
	high := fp.ParseZval()
	fp.StartOptional()
	step := fp.ParseZval()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifRange(executeData.Ctx(), low, high, nil, step)
	if ok {
		returnValue.SetArrayOfZval(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifShuffle
var DefZifShuffle = def.DefFunc("shuffle", 1, 1, []def.ArgInfo{{Name: "arg", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	arg := fp.ParseRefArray()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifShuffle(executeData.Ctx(), arg)
	returnValue.SetBool(ret)
})

// generate by ZifArrayPush
var DefZifArrayPush = def.DefFunc("array_push", 1, -1, []def.ArgInfo{{Name: "stack", ByRef: true}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	stack := fp.ParseRefArray()
	fp.StartOptional()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayPush(executeData.Ctx(), stack, nil, args)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayPop
var DefZifArrayPop = def.DefFunc("array_pop", 1, 1, []def.ArgInfo{{Name: "stack", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	stack := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayPop(stack)
	returnValue.SetBy(ret)
})

// generate by ZifArrayShift
var DefZifArrayShift = def.DefFunc("array_shift", 1, 1, []def.ArgInfo{{Name: "stack", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	stack := fp.ParseRefArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayShift(executeData.Ctx(), stack)
	returnValue.SetBy(ret)
})

// generate by ZifArrayUnshift
var DefZifArrayUnshift = def.DefFunc("array_unshift", 1, -1, []def.ArgInfo{{Name: "stack", ByRef: true}, {Name: "values", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	stack := fp.ParseRefArray()
	values := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayUnshift(stack, values)
	returnValue.SetLong(ret)
})

// generate by ZifArraySplice
var DefZifArraySplice = def.DefFunc("array_splice", 2, 4, []def.ArgInfo{{Name: "array", ByRef: true}, {Name: "offset"}, {Name: "length"}, {Name: "replacement"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
	array := fp.ParseRefArray()
	offset := fp.ParseLong()
	fp.StartOptional()
	length_ := fp.ParseLongNullable()
	replacement := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArraySplice(executeData.Ctx(), array, offset, nil, length_, replacement)
	returnValue.SetArray(ret)
})

// generate by ZifArraySlice
var DefZifArraySlice = def.DefFunc("array_slice", 2, 4, []def.ArgInfo{{Name: "array"}, {Name: "offset"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	offset := fp.ParseLong()
	fp.StartOptional()
	length_ := fp.ParseZval()
	preserve_keys := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifArraySlice(executeData.Ctx(), array, offset, nil, length_, preserve_keys)
	returnValue.SetArray(ret)
})

// generate by ZifArrayMerge
var DefZifArrayMerge = def.DefFunc("array_merge", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifArrayMerge(executeData.Ctx(), returnValue, nil, arrays)
})

// generate by ZifArrayMergeRecursive
var DefZifArrayMergeRecursive = def.DefFunc("array_merge_recursive", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifArrayMergeRecursive(executeData.Ctx(), returnValue, nil, arrays)
})

// generate by ZifArrayReplace
var DefZifArrayReplace = def.DefFunc("array_replace", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayReplace(executeData.Ctx(), arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayReplaceRecursive
var DefZifArrayReplaceRecursive = def.DefFunc("array_replace_recursive", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayReplaceRecursive(executeData.Ctx(), arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeys
var DefZifArrayKeys = def.DefFunc("array_keys", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "search_value"}, {Name: "strict"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	fp.StartOptional()
	search_value := fp.ParseZvalNullable()
	strict := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeys(executeData.Ctx(), array, nil, search_value, strict)
	returnValue.SetArray(ret)
})

// generate by ZifArrayKeyFirst
var DefZifArrayKeyFirst = def.DefFunc("array_key_first", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyFirst(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeyLast
var DefZifArrayKeyLast = def.DefFunc("array_key_last", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyLast(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayValues
var DefZifArrayValues = def.DefFunc("array_values", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayValues(array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayCountValues
var DefZifArrayCountValues = def.DefFunc("array_count_values", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayCountValues(executeData.Ctx(), array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayColumn
var DefZifArrayColumn = def.DefFunc("array_column", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "column_key"}, {Name: "index_key"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	column_key := fp.ParseZvalNullable()
	fp.StartOptional()
	index_key := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayColumn(executeData.Ctx(), array, column_key, nil, index_key)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayReverse
var DefZifArrayReverse = def.DefFunc("array_reverse", 1, 2, []def.ArgInfo{{Name: "array"}, {Name: "preserve_keys"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	fp.StartOptional()
	preserve_keys := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReverse(array, nil, preserve_keys)
	returnValue.SetArray(ret)
})

// generate by ZifArrayPad
var DefZifArrayPad = def.DefFunc("array_pad", 3, 3, []def.ArgInfo{{Name: "array"}, {Name: "pad_size"}, {Name: "pad_value"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	pad_size := fp.ParseLong()
	pad_value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayPad(executeData.Ctx(), array, pad_size, pad_value)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayFlip
var DefZifArrayFlip = def.DefFunc("array_flip", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFlip(executeData.Ctx(), array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayChangeKeyCase
var DefZifArrayChangeKeyCase = def.DefFunc("array_change_key_case", 1, 2, []def.ArgInfo{{Name: "array"}, {Name: "case"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	fp.StartOptional()
	case_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayChangeKeyCase(array, nil, case_)
	returnValue.SetArray(ret)
})

// generate by ZifArrayUnique
var DefZifArrayUnique = def.DefFunc("array_unique", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "flags"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseArray()
	fp.StartOptional()
	flags := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayUnique(executeData.Ctx(), arg, nil, flags)
	returnValue.SetArray(ret)
})

// generate by ZifArrayIntersectKey
var DefZifArrayIntersectKey = def.DefFunc("array_intersect_key", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectKey(executeData.Ctx(), arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectUkey
var DefZifArrayIntersectUkey = def.DefFunc("array_intersect_ukey", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_key_compare_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectUkey(executeData.Ctx(), arrays, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersect
var DefZifArrayIntersect = def.DefFunc("array_intersect", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersect(executeData.Ctx(), arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersect
var DefZifArrayUintersect = def.DefFunc("array_uintersect", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_compare_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_data_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersect(executeData.Ctx(), arrays, callback_data_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectAssoc
var DefZifArrayIntersectAssoc = def.DefFunc("array_intersect_assoc", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectAssoc(executeData.Ctx(), arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectUassoc
var DefZifArrayIntersectUassoc = def.DefFunc("array_intersect_uassoc", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_key_compare_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectUassoc(executeData.Ctx(), arrays, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersectAssoc
var DefZifArrayUintersectAssoc = def.DefFunc("array_uintersect_assoc", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_compare_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_data_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersectAssoc(executeData.Ctx(), arrays, callback_data_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersectUassoc
var DefZifArrayUintersectUassoc = def.DefFunc("array_uintersect_uassoc", 2, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_compare_func"}, {Name: "callback_key_compare_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(2)
	callback_data_compare_func := fp.ParseCallable()
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersectUassoc(executeData.Ctx(), arrays, callback_data_compare_func, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffKey
var DefZifArrayDiffKey = def.DefFunc("array_diff_key", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffKey(executeData.Ctx(), arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffUkey
var DefZifArrayDiffUkey = def.DefFunc("array_diff_ukey", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_key_comp_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffUkey(executeData.Ctx(), arrays, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiff
var DefZifArrayDiff = def.DefFunc("array_diff", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayDiff(executeData.Ctx(), arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayUdiff
var DefZifArrayUdiff = def.DefFunc("array_udiff", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_comp_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_data_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiff(executeData.Ctx(), arrays, callback_data_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffAssoc
var DefZifArrayDiffAssoc = def.DefFunc("array_diff_assoc", 0, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffAssoc(executeData.Ctx(), arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffUassoc
var DefZifArrayDiffUassoc = def.DefFunc("array_diff_uassoc", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_key_comp_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffUassoc(executeData.Ctx(), arrays, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUdiffAssoc
var DefZifArrayUdiffAssoc = def.DefFunc("array_udiff_assoc", 1, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_comp_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(1)
	callback_data_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiffAssoc(executeData.Ctx(), arrays, callback_data_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUdiffUassoc
var DefZifArrayUdiffUassoc = def.DefFunc("array_udiff_uassoc", 2, -1, []def.ArgInfo{{Name: "arrays", Variadic: true}, {Name: "callback_data_comp_func"}, {Name: "callback_key_comp_func"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, -1, 0)
	fp.CheckNumArgs()
	arrays := fp.ParseVariadic(2)
	callback_data_comp_func := fp.ParseCallable()
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiffUassoc(executeData.Ctx(), arrays, callback_data_comp_func, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayMultisort
var DefZifArrayMultisort = def.DefFunc("array_multisort", 0, -1, []def.ArgInfo{{Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayMultisort(executeData.Ctx(), args)
	returnValue.SetBool(ret)
})

// generate by ZifArrayRand
var DefZifArrayRand = def.DefFunc("array_rand", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "num_req"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	arg := fp.ParseArray()
	fp.StartOptional()
	num_req_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayRand(executeData.Ctx(), arg, nil, num_req_)
	returnValue.SetBy(ret)
})

// generate by ZifArraySum
var DefZifArraySum = def.DefFunc("array_sum", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArraySum(executeData.Ctx(), array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayProduct
var DefZifArrayProduct = def.DefFunc("array_product", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret := ZifArrayProduct(executeData.Ctx(), array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayReduce
var DefZifArrayReduce = def.DefFunc("array_reduce", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "callback"}, {Name: "initial"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	callback := fp.ParseCallable()
	fp.StartOptional()
	initial := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReduce(array, callback, nil, initial)
	returnValue.SetBy(ret)
})

// generate by ZifArrayFilter
var DefZifArrayFilter = def.DefFunc("array_filter", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "callback"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	array_ := fp.ParseArray()
	fp.StartOptional()
	callback := fp.ParseCallable()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFilter(executeData.Ctx(), array_, nil, callback, mode)
	returnValue.SetArray(ret)
})

// generate by ZifArrayMap
var DefZifArrayMap = def.DefFunc("array_map", 1, -1, []def.ArgInfo{{Name: "callback"}, {Name: "arrays", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	callback := fp.ParseCallable()
	arrays := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret := ZifArrayMap(executeData.Ctx(), callback, arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeyExists
var DefZifArrayKeyExists = def.DefFunc("array_key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	key := fp.ParseZvalNullable()
	array := fp.ParseArrayOrObjectZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyExists(executeData.Ctx(), key, array)
	returnValue.SetBool(ret)
})

// generate by ZifArrayKeyExists
var DefZifKeyExists = def.DefFunc("key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "array"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	key := fp.ParseZvalNullable()
	array := fp.ParseArrayOrObjectZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyExists(executeData.Ctx(), key, array)
	returnValue.SetBool(ret)
})

// generate by ZifArrayChunk
var DefZifArrayChunk = def.DefFunc("array_chunk", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	array := fp.ParseArray()
	length := fp.ParseLong()
	fp.StartOptional()
	preserve_keys := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifArrayChunk(executeData.Ctx(), array, length, nil, preserve_keys)
	returnValue.SetBy(ret)
})

// generate by ZifArrayCombine
var DefZifArrayCombine = def.DefFunc("array_combine", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "values"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	keys := fp.ParseArray()
	values := fp.ParseArray()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayCombine(executeData.Ctx(), keys, values)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifBase64Encode
var DefZifBase64Encode = def.DefFunc("base64_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifBase64Encode(str)
	returnValue.SetString(ret)
})

// generate by ZifBase64Decode
var DefZifBase64Decode = def.DefFunc("base64_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "strict"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	strict := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBase64Decode(str, nil, strict)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifInetNtop
var DefZifInetNtop = def.DefFunc("inet_ntop", 1, 1, []def.ArgInfo{{Name: "ip"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	ip := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifInetNtop(ip)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifInetPton
var DefZifInetPton = def.DefFunc("inet_pton", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	ip_address := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifInetPton(ip_address)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIp2long
var DefZifIp2long = def.DefFunc("ip2long", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	ip_address := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIp2long(ip_address)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifLong2ip
var DefZifLong2ip = def.DefFunc("long2ip", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	ip_address := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifLong2ip(ip_address)
	returnValue.SetString(ret)
})

// generate by ZifIniGet
var DefZifIniGet = def.DefFunc("ini_get", 1, 1, []def.ArgInfo{{Name: "varname"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	varname := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniGet(executeData.Ctx(), varname)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniGetAll
var DefZifIniGetAll = def.DefFunc("ini_get_all", 0, 2, []def.ArgInfo{{Name: "extension"}, {Name: "details"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	extension := fp.ParseStringNullable()
	details_ := fp.ParseBoolNullable()
	if fp.HasError() {
		return
	}
	ZifIniGetAll(executeData.Ctx(), returnValue, nil, extension, details_)
})

// generate by ZifIniSet
var DefZifIniSet = def.DefFunc("ini_set", 2, 2, []def.ArgInfo{{Name: "varname"}, {Name: "newvalue"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	varname := fp.ParseString()
	newvalue := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniSet(executeData.Ctx(), varname, newvalue)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniSet
var DefZifIniAlter = def.DefFunc("ini_alter", 2, 2, []def.ArgInfo{{Name: "varname"}, {Name: "newvalue"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	varname := fp.ParseString()
	newvalue := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniSet(executeData.Ctx(), varname, newvalue)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniRestore
var DefZifIniRestore = def.DefFunc("ini_restore", 1, 1, []def.ArgInfo{{Name: "var_name"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_name := fp.ParseString()
	if fp.HasError() {
		return
	}
	ZifIniRestore(executeData.Ctx(), var_name)
})

// generate by ZifSetIncludePath
var DefZifSetIncludePath = def.DefFunc("set_include_path", 1, 1, []def.ArgInfo{{Name: "new_include_path"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	new_include_path := fp.ParsePath()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSetIncludePath(executeData.Ctx(), new_include_path)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetIncludePath
var DefZifGetIncludePath = def.DefFunc("get_include_path", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret, ok := ZifGetIncludePath(executeData.Ctx())
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRestoreIncludePath
var DefZifRestoreIncludePath = def.DefFunc("restore_include_path", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ZifRestoreIncludePath(executeData.Ctx())
})

// generate by ZifPrintR
var DefZifPrintR = def.DefFunc("print_r", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "return"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseBool()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifPrintR(executeData.Ctx(), var_, nil, return_)
	returnValue.SetByPtr(ret)
})

// generate by ZifParseIniFile
var DefZifParseIniFile = def.DefFunc("parse_ini_file", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "process_sections"}, {Name: "scanner_mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	filename := fp.ParseString()
	fp.StartOptional()
	process_sections := fp.ParseBool()
	scanner_mode := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifParseIniFile(executeData.Ctx(), filename, nil, process_sections, scanner_mode)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifParseIniString
var DefZifParseIniString = def.DefFunc("parse_ini_string", 1, 3, []def.ArgInfo{{Name: "ini_string"}, {Name: "process_sections"}, {Name: "scanner_mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	ini_string := fp.ParseString()
	fp.StartOptional()
	process_sections := fp.ParseBool()
	scanner_mode := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifParseIniString(executeData.Ctx(), ini_string, nil, process_sections, scanner_mode)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCrc32
var DefZifCrc32 = def.DefFunc("crc32", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifCrc32(str)
	returnValue.SetLong(ret)
})

// generate by ZifStrGetcsv
var DefZifStrGetcsv = def.DefFunc("str_getcsv", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
	string_ := fp.ParseString()
	fp.StartOptional()
	delimiter := fp.ParseString()
	enclosure := fp.ParseString()
	escape := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifStrGetcsv(string_, nil, delimiter, enclosure, escape)
	returnValue.SetArray(ret)
})

// generate by ZifGethostname
var DefZifGethostname = def.DefFunc("gethostname", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret, ok := ZifGethostname()
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGethostbyaddr
var DefZifGethostbyaddr = def.DefFunc("gethostbyaddr", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	ip_address := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifGethostbyaddr(ip_address)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGethostbyname
var DefZifGethostbyname = def.DefFunc("gethostbyname", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	hostname := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifGethostbyname(executeData.Ctx(), hostname)
	returnValue.SetString(ret)
})

// generate by ZifGethostbynamel
var DefZifGethostbynamel = def.DefFunc("gethostbynamel", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	hostname := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifGethostbynamel(executeData.Ctx(), hostname)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifDnsGetMx
var DefZifDnsGetMx = def.DefFunc("dns_get_mx", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "hosts", ByRef: true}, {Name: "weights", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	hostname := fp.ParseString()
	hosts := fp.ParseRefZval()
	fp.StartOptional()
	weights := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ret := ZifDnsGetMx(hostname, hosts, nil, weights)
	returnValue.SetBool(ret)
})

// generate by ZifDnsGetMx
var DefZifGetmxrr = def.DefFunc("getmxrr", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "hosts", ByRef: true}, {Name: "weights", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	hostname := fp.ParseString()
	hosts := fp.ParseRefZval()
	fp.StartOptional()
	weights := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ret := ZifDnsGetMx(hostname, hosts, nil, weights)
	returnValue.SetBool(ret)
})

// generate by ZifSprintf
var DefZifSprintf = def.DefFunc("sprintf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifSprintf(executeData.Ctx(), format, nil, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifVsprintf
var DefZifVsprintf = def.DefFunc("vsprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifVsprintf(executeData.Ctx(), format, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifPrintf
var DefZifPrintf = def.DefFunc("printf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifPrintf(executeData.Ctx(), format, nil, args)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifVprintf
var DefZifVprintf = def.DefFunc("vprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifVprintf(executeData.Ctx(), format, args)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifHtmlspecialchars
var DefZifHtmlspecialchars = def.DefFunc("htmlspecialchars", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "quote_style"}, {Name: "encoding"}, {Name: "double_encode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
	string := fp.ParseString()
	fp.StartOptional()
	quote_style := fp.ParseLongNullable()
	encoding := fp.ParseStringNullable()
	double_encode_ := fp.ParseBoolNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlspecialchars(executeData.Ctx(), string, nil, quote_style, encoding, double_encode_)
	returnValue.SetString(ret)
})

// generate by ZifHtmlspecialcharsDecode
var DefZifHtmlspecialcharsDecode = def.DefFunc("htmlspecialchars_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "quote_style"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	quote_style_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlspecialcharsDecode(executeData.Ctx(), str, nil, quote_style_)
	returnValue.SetString(ret)
})

// generate by ZifHtmlEntityDecode
var DefZifHtmlEntityDecode = def.DefFunc("html_entity_decode", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "quote_style"}, {Name: "encoding"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	quote_style_ := fp.ParseLongNullable()
	encoding := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlEntityDecode(executeData.Ctx(), str, nil, quote_style_, encoding)
	returnValue.SetString(ret)
})

// generate by ZifHtmlentities
var DefZifHtmlentities = def.DefFunc("htmlentities", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "quote_style"}, {Name: "encoding"}, {Name: "double_encode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
	string := fp.ParseString()
	fp.StartOptional()
	quote_style := fp.ParseLongNullable()
	encoding := fp.ParseStringNullable()
	double_encode_ := fp.ParseBoolNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlentities(executeData.Ctx(), string, nil, quote_style, encoding, double_encode_)
	returnValue.SetString(ret)
})

// generate by ZifGetHtmlTranslationTable
var DefZifGetHtmlTranslationTable = def.DefFunc("get_html_translation_table", 0, 3, []def.ArgInfo{{Name: "table"}, {Name: "quote_style"}, {Name: "encoding"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 3, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	table := fp.ParseLongNullable()
	quote_style_ := fp.ParseLongNullable()
	encoding := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifGetHtmlTranslationTable(executeData.Ctx(), returnValue, nil, table, quote_style_, encoding)
	returnValue.SetArray(ret)
})

// generate by ZifHttpBuildQuery
var DefZifHttpBuildQuery = def.DefFunc("http_build_query", 1, 4, []def.ArgInfo{{Name: "formdata"}, {Name: "prefix"}, {Name: "arg_separator"}, {Name: "enc_type"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
	formdata := fp.ParseZval()
	fp.StartOptional()
	prefix := fp.ParseString()
	arg_separator := fp.ParseString()
	enc_type_ := fp.ParseLongNullable()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifHttpBuildQuery(executeData.Ctx(), formdata, nil, prefix, arg_separator, enc_type_)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifAbs
var DefZifAbs = def.DefFunc("abs", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifAbs(executeData.Ctx(), number)
	returnValue.SetBy(ret)
})

// generate by ZifCeil
var DefZifCeil = def.DefFunc("ceil", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCeil(executeData.Ctx(), number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFloor
var DefZifFloor = def.DefFunc("floor", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFloor(executeData.Ctx(), number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRound
var DefZifRound = def.DefFunc("round", 1, 3, []def.ArgInfo{{Name: "number"}, {Name: "precision"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	number := fp.ParseZval()
	fp.StartOptional()
	precision := fp.ParseLong()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifRound(executeData.Ctx(), number, nil, precision, mode_)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSin
var DefZifSin = def.DefFunc("sin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSin(number)
	returnValue.SetDouble(ret)
})

// generate by ZifCos
var DefZifCos = def.DefFunc("cos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifCos(number)
	returnValue.SetDouble(ret)
})

// generate by ZifTan
var DefZifTan = def.DefFunc("tan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifTan(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAsin
var DefZifAsin = def.DefFunc("asin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAsin(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAcos
var DefZifAcos = def.DefFunc("acos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAcos(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtan
var DefZifAtan = def.DefFunc("atan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtan(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtan2
var DefZifAtan2 = def.DefFunc("atan2", 2, 2, []def.ArgInfo{{Name: "y"}, {Name: "x"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	y := fp.ParseDouble()
	x := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtan2(y, x)
	returnValue.SetDouble(ret)
})

// generate by ZifSinh
var DefZifSinh = def.DefFunc("sinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSinh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifCosh
var DefZifCosh = def.DefFunc("cosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifCosh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifTanh
var DefZifTanh = def.DefFunc("tanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifTanh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAsinh
var DefZifAsinh = def.DefFunc("asinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAsinh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAcosh
var DefZifAcosh = def.DefFunc("acosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAcosh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtanh
var DefZifAtanh = def.DefFunc("atanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtanh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifPi
var DefZifPi = def.DefFunc("pi", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifPi()
	returnValue.SetDouble(ret)
})

// generate by ZifIsFinite
var DefZifIsFinite = def.DefFunc("is_finite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsFinite(val)
	returnValue.SetBool(ret)
})

// generate by ZifIsInfinite
var DefZifIsInfinite = def.DefFunc("is_infinite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsInfinite(val)
	returnValue.SetBool(ret)
})

// generate by ZifIsNan
var DefZifIsNan = def.DefFunc("is_nan", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsNan(val)
	returnValue.SetBool(ret)
})

// generate by ZifPow
var DefZifPow = def.DefFunc("pow", 2, 2, []def.ArgInfo{{Name: "base"}, {Name: "exponent"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	base := fp.ParseZval()
	exponent := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifPow(executeData.Ctx(), base, exponent)
	returnValue.SetBy(ret)
})

// generate by ZifExp
var DefZifExp = def.DefFunc("exp", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifExp(number)
	returnValue.SetDouble(ret)
})

// generate by ZifExpm1
var DefZifExpm1 = def.DefFunc("expm1", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifExpm1(number)
	returnValue.SetDouble(ret)
})

// generate by ZifLog1p
var DefZifLog1p = def.DefFunc("log1p", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifLog1p(number)
	returnValue.SetDouble(ret)
})

// generate by ZifLog
var DefZifLog = def.DefFunc("log", 1, 2, []def.ArgInfo{{Name: "number"}, {Name: "base"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	fp.StartOptional()
	base := fp.ParseDoubleNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifLog(executeData.Ctx(), number, nil, base)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifLog10
var DefZifLog10 = def.DefFunc("log10", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifLog10(number)
	returnValue.SetDouble(ret)
})

// generate by ZifSqrt
var DefZifSqrt = def.DefFunc("sqrt", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSqrt(number)
	returnValue.SetDouble(ret)
})

// generate by ZifHypot
var DefZifHypot = def.DefFunc("hypot", 2, 2, []def.ArgInfo{{Name: "num1"}, {Name: "num2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	num1 := fp.ParseDouble()
	num2 := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifHypot(num1, num2)
	returnValue.SetDouble(ret)
})

// generate by ZifDeg2rad
var DefZifDeg2rad = def.DefFunc("deg2rad", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifDeg2rad(number)
	returnValue.SetDouble(ret)
})

// generate by ZifRad2deg
var DefZifRad2deg = def.DefFunc("rad2deg", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifRad2deg(number)
	returnValue.SetDouble(ret)
})

// generate by ZifBindec
var DefZifBindec = def.DefFunc("bindec", 1, 1, []def.ArgInfo{{Name: "binary_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	binary_number := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBindec(executeData.Ctx(), binary_number)
	if ok {
		returnValue.SetByPtr(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifHexdec
var DefZifHexdec = def.DefFunc("hexdec", 1, 1, []def.ArgInfo{{Name: "hexadecimal_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	hexadecimal_number := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHexdec(executeData.Ctx(), hexadecimal_number)
	if ok {
		returnValue.SetByPtr(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifOctdec
var DefZifOctdec = def.DefFunc("octdec", 1, 1, []def.ArgInfo{{Name: "octal_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	octal_number := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifOctdec(executeData.Ctx(), octal_number)
	if ok {
		returnValue.SetByPtr(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifDecbin
var DefZifDecbin = def.DefFunc("decbin", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDecbin(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifDecoct
var DefZifDecoct = def.DefFunc("decoct", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDecoct(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifDechex
var DefZifDechex = def.DefFunc("dechex", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDechex(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifBaseConvert
var DefZifBaseConvert = def.DefFunc("base_convert", 3, 3, []def.ArgInfo{{Name: "number"}, {Name: "frombase"}, {Name: "tobase"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 3, 0)
	fp.CheckNumArgs()
	number := fp.ParseZval()
	frombase := fp.ParseLong()
	tobase := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBaseConvert(executeData.Ctx(), number, frombase, tobase)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifNumberFormat
var DefZifNumberFormat = def.DefFunc("number_format", 1, 4, []def.ArgInfo{{Name: "number"}, {Name: "num_decimal_places"}, {Name: "dec_separator"}, {Name: "thousands_separator"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
	number := fp.ParseDouble()
	fp.StartOptional()
	num_decimal_places := fp.ParseLong()
	dec_separator_ := fp.ParseStringNullable()
	thousands_separator := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifNumberFormat(number, nil, num_decimal_places, dec_separator_, thousands_separator)
	returnValue.SetString(ret)
})

// generate by ZifFmod
var DefZifFmod = def.DefFunc("fmod", 2, 2, []def.ArgInfo{{Name: "x"}, {Name: "y"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	x := fp.ParseDouble()
	y := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifFmod(x, y)
	returnValue.SetDouble(ret)
})

// generate by ZifIntdiv
var DefZifIntdiv = def.DefFunc("intdiv", 2, 2, []def.ArgInfo{{Name: "dividend"}, {Name: "divisor"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	dividend := fp.ParseLong()
	divisor := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifIntdiv(executeData.Ctx(), dividend, divisor)
	returnValue.SetLong(ret)
})

// generate by ZifMd5
var DefZifMd5 = def.DefFunc("md5", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	raw_output := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifMd5(str, nil, raw_output)
	returnValue.SetString(ret)
})

// generate by ZifMd5File
var DefZifMd5File = def.DefFunc("md5_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	filename := fp.ParseString()
	fp.StartOptional()
	raw_output := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifMd5File(filename, nil, raw_output)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifMicrotime
var DefZifMicrotime = def.DefFunc("microtime", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	get_as_float := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifMicrotime(nil, get_as_float)
	returnValue.SetBy(ret)
})

// generate by ZifGettimeofday
var DefZifGettimeofday = def.DefFunc("gettimeofday", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	get_as_float := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifGettimeofday(nil, get_as_float)
	returnValue.SetBy(ret)
})

// generate by ZifGetrusage
var DefZifGetrusage = def.DefFunc("getrusage", 0, 1, []def.ArgInfo{{Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifGetrusage(nil, mode)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifNetGetInterfaces
var DefZifNetGetInterfaces = def.DefFunc("net_get_interfaces", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret, ok := ZifNetGetInterfaces()
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifPack
var DefZifPack = def.DefFunc("pack", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, -1, 0)
	fp.CheckNumArgs()
	format_ := fp.ParseString()
	fp.StartOptional()
	args := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ret, ok := ZifPack(executeData.Ctx(), format_, nil, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifUnpack
var DefZifUnpack = def.DefFunc("unpack", 2, 3, []def.ArgInfo{{Name: "format"}, {Name: "input"}, {Name: "offset"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	format := fp.ParseString()
	input := fp.ParseString()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifUnpack(executeData.Ctx(), format, input, nil, offset)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifQuotedPrintableDecode
var DefZifQuotedPrintableDecode = def.DefFunc("quoted_printable_decode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableDecode(str)
	returnValue.SetString(ret)
})

// generate by ZifQuotedPrintableEncode
var DefZifQuotedPrintableEncode = def.DefFunc("quoted_printable_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableEncode(str)
	returnValue.SetString(ret)
})

// generate by ZifMtSrand
var DefZifMtSrand = def.DefFunc("mt_srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	seed_ := fp.ParseLongNullable()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ZifMtSrand(executeData.Ctx(), nil, seed_, mode_)
})

// generate by ZifMtSrand
var DefZifSrand = def.DefFunc("srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	seed_ := fp.ParseLongNullable()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ZifMtSrand(executeData.Ctx(), nil, seed_, mode_)
})

// generate by ZifMtGetrandmax
var DefZifMtGetrandmax = def.DefFunc("mt_getrandmax", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifMtGetrandmax()
	returnValue.SetLong(ret)
})

// generate by ZifMtGetrandmax
var DefZifGetrandmax = def.DefFunc("getrandmax", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifMtGetrandmax()
	returnValue.SetLong(ret)
})

// generate by ZifMtRand
var DefZifMtRand = def.DefFunc("mt_rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	min_ := fp.ParseLongNullable()
	max_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifMtRand(executeData.Ctx(), nil, min_, max_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRand
var DefZifRand = def.DefFunc("rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	min_ := fp.ParseLongNullable()
	max_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRand(executeData.Ctx(), nil, min_, max_)
	returnValue.SetLong(ret)
})

// generate by ZifRandomBytes
var DefZifRandomBytes = def.DefFunc("random_bytes", 1, 1, []def.ArgInfo{{Name: "length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	length := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomBytes(executeData.Ctx(), length)
	returnValue.SetString(ret)
})

// generate by ZifRandomInt
var DefZifRandomInt = def.DefFunc("random_int", 2, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	min := fp.ParseLong()
	max := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomInt(executeData.Ctx(), min, max)
	returnValue.SetLong(ret)
})

// generate by ZifSha1
var DefZifSha1 = def.DefFunc("sha1", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	raw_output := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifSha1(str, nil, raw_output)
	returnValue.SetString(ret)
})

// generate by ZifSha1File
var DefZifSha1File = def.DefFunc("sha1_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	filename := fp.ParseString()
	fp.StartOptional()
	raw_output := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSha1File(filename, nil, raw_output)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSoundex
var DefZifSoundex = def.DefFunc("soundex", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSoundex(str)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifUtf8Encode
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Encode(data)
	returnValue.SetString(ret)
})

// generate by ZifUtf8Decode
var DefZifUtf8Decode = def.DefFunc("utf8_decode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Decode(data)
	returnValue.SetString(ret)
})

// generate by ZifBin2hex
var DefZifBin2hex = def.DefFunc("bin2hex", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifBin2hex(data)
	returnValue.SetString(ret)
})

// generate by ZifHex2bin
var DefZifHex2bin = def.DefFunc("hex2bin", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifTrim(executeData.Ctx(), executeData, str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifRtrim
var DefZifRtrim = def.DefFunc("rtrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(executeData.Ctx(), executeData, str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifRtrim
var DefZifChop = def.DefFunc("chop", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifRtrim(executeData.Ctx(), executeData, str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifLtrim
var DefZifLtrim = def.DefFunc("ltrim", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "character_mask"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	character_mask := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifLtrim(executeData.Ctx(), executeData, str, nil, character_mask)
	returnValue.SetString(ret)
})

// generate by ZifWordwrap
var DefZifWordwrap = def.DefFunc("wordwrap", 1, 4, []def.ArgInfo{{Name: "str"}, {Name: "width"}, {Name: "break"}, {Name: "cut"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 4, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	separator := fp.ParseString()
	str := fp.ParseString()
	fp.StartOptional()
	limit := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifExplode(executeData.Ctx(), separator, str, nil, limit)
	if ok {
		returnValue.SetArrayOfString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifImplode
var DefZifImplode = def.DefFunc("implode", 1, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	glue_ := fp.ParseZval()
	fp.StartOptional()
	pieces_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(executeData.Ctx(), glue_, nil, pieces_)
	returnValue.SetBy(ret)
})

// generate by ZifImplode
var DefZifJoin = def.DefFunc("join", 1, 2, []def.ArgInfo{{Name: "glue"}, {Name: "pieces"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	glue_ := fp.ParseZval()
	fp.StartOptional()
	pieces_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifImplode(executeData.Ctx(), glue_, nil, pieces_)
	returnValue.SetBy(ret)
})

// generate by ZifStrtok
var DefZifStrtok = def.DefFunc("strtok", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "token"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrtoupper(str)
	returnValue.SetString(ret)
})

// generate by ZifStrtolower
var DefZifStrtolower = def.DefFunc("strtolower", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrtolower(str)
	returnValue.SetString(ret)
})

// generate by ZifBasename
var DefZifBasename = def.DefFunc("basename", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "suffix"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	path := fp.ParseString()
	fp.StartOptional()
	levels_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifDirname(executeData.Ctx(), path, nil, levels_)
	returnValue.SetBy(ret)
})

// generate by ZifPathinfo
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "options"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	path := fp.ParseString()
	fp.StartOptional()
	options := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifPathinfo(path, nil, options)
	returnValue.SetBy(ret)
})

// generate by ZifStristr
var DefZifStristr = def.DefFunc("stristr", 2, 3, []def.ArgInfo{{Name: "haystack"}, {Name: "needle"}, {Name: "part"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	haystack := fp.ParseString()
	needle := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseLong()
	if fp.HasError() {
		returnValue.SetFalse()
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
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 3, 4, 0)
	fp.CheckNumArgs()
	str := fp.ParseZval()
	replace := fp.ParseZval()
	start := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ZifSubstrReplace(executeData.Ctx(), returnValue, str, replace, start, nil, length)
})

// generate by ZifQuotemeta
var DefZifQuotemeta = def.DefFunc("quotemeta", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	character := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifOrd(character)
	returnValue.SetLong(ret)
})

// generate by ZifChr
var DefZifChr = def.DefFunc("chr", 1, 1, []def.ArgInfo{{Name: "codepoint"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	codepoint := fp.ParseLong()
	ret := ZifChr(codepoint)
	returnValue.SetString(ret)
})

// generate by ZifUcfirst
var DefZifUcfirst = def.DefFunc("ucfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUcfirst(str)
	returnValue.SetString(ret)
})

// generate by ZifLcfirst
var DefZifLcfirst = def.DefFunc("lcfirst", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifLcfirst(str)
	returnValue.SetString(ret)
})

// generate by ZifUcwords
var DefZifUcwords = def.DefFunc("ucwords", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "delimiters"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrrev(str)
	returnValue.SetString(ret)
})

// generate by ZifSimilarText
var DefZifSimilarText = def.DefFunc("similar_text", 2, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "percent", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifAddslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifAddcslashes
var DefZifAddcslashes = def.DefFunc("addcslashes", 2, 2, []def.ArgInfo{{Name: "str"}, {Name: "charlist"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStripslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifStripcslashes
var DefZifStripcslashes = def.DefFunc("stripcslashes", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStripcslashes(str)
	returnValue.SetString(ret)
})

// generate by ZifStrReplace
var DefZifStrReplace = def.DefFunc("str_replace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 4, 0)
	fp.CheckNumArgs()
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
var DefZifStrIreplace = def.DefFunc("str_ireplace", 3, 4, []def.ArgInfo{{Name: "search"}, {Name: "replace"}, {Name: "subject"}, {Name: "replace_count", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 4, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	allowable_tags := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifStripTags(executeData.Ctx(), str, nil, allowable_tags)
	returnValue.SetString(ret)
})

// generate by ZifStrRepeat
var DefZifStrRepeat = def.DefFunc("str_repeat", 2, 2, []def.ArgInfo{{Name: "input"}, {Name: "mult"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	input := fp.ParseString()
	mult := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifStrRepeat(executeData.Ctx(), input, mult)
	returnValue.SetBy(ret)
})

// generate by ZifCountChars
var DefZifCountChars = def.DefFunc("count_chars", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "mode"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	input := fp.ParseString()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCountChars(executeData.Ctx(), input, nil, mode)
	if ok {
		returnValue.SetByPtr(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrnatcmp
var DefZifStrnatcmp = def.DefFunc("strnatcmp", 2, 2, []def.ArgInfo{{Name: "s1"}, {Name: "s2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 4, 0)
	fp.CheckNumArgs()
	input := fp.ParseString()
	pad_length := fp.ParseLong()
	fp.StartOptional()
	pad_string_ := fp.ParseStringNullable()
	pad_type_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifStrPad(executeData.Ctx(), executeData, input, pad_length, nil, pad_string_, pad_type_)
	returnValue.SetBy(ret)
})

// generate by ZifStrRot13
var DefZifStrRot13 = def.DefFunc("str_rot13", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrRot13(str)
	returnValue.SetString(ret)
})

// generate by ZifStrShuffle
var DefZifStrShuffle = def.DefFunc("str_shuffle", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrShuffle(str)
	returnValue.SetString(ret)
})

// generate by ZifStrWordCount
var DefZifStrWordCount = def.DefFunc("str_word_count", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "format"}, {Name: "charlist"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 3, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	fp.StartOptional()
	format := fp.ParseLong()
	charlist := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrWordCount(executeData.Ctx(), str, nil, format, charlist)
	if ok {
		returnValue.SetByPtr(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrSplit
var DefZifStrSplit = def.DefFunc("str_split", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "split_length"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
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
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	haystack := fp.ParseString()
	char_list := fp.ParseString()
	if fp.HasError() {
		returnValue.SetFalse()
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
	fp := php.NewParamParser(executeData, 3, 5, 0)
	fp.CheckNumArgs()
	haystack := fp.ParseString()
	needle := fp.ParseString()
	offset := fp.ParseLong()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	case_insensitivity := fp.ParseBool()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifSubstrCompare(executeData.Ctx(), returnValue, haystack, needle, offset, nil, length, case_insensitivity)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGettype
var DefZifGettype = def.DefFunc("gettype", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifGettype(var_)
	returnValue.SetString(ret)
})

// generate by ZifSettype
var DefZifSettype = def.DefFunc("settype", 2, 2, []def.ArgInfo{{Name: "var", ByRef: true}, {Name: "typ"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseRefZval()
	typ := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifSettype(executeData.Ctx(), var_, typ)
	returnValue.SetBool(ret)
})

// generate by ZifIntval
var DefZifIntval = def.DefFunc("intval", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "base"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	fp.StartOptional()
	base_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIntval(executeData.Ctx(), var_, nil, base_)
	returnValue.SetLong(ret)
})

// generate by ZifFloatval
var DefZifFloatval = def.DefFunc("floatval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifFloatval(executeData.Ctx(), var_)
	returnValue.SetDouble(ret)
})

// generate by ZifFloatval
var DefZifDoubleval = def.DefFunc("doubleval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifFloatval(executeData.Ctx(), var_)
	returnValue.SetDouble(ret)
})

// generate by ZifBoolval
var DefZifBoolval = def.DefFunc("boolval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifBoolval(executeData.Ctx(), var_)
	returnValue.SetBool(ret)
})

// generate by ZifStrval
var DefZifStrval = def.DefFunc("strval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifStrval(executeData.Ctx(), var_)
	returnValue.SetString(ret)
})

// generate by ZifIsNull
var DefZifIsNull = def.DefFunc("is_null", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsNull(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsResource
var DefZifIsResource = def.DefFunc("is_resource", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsResource(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsBool
var DefZifIsBool = def.DefFunc("is_bool", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret := ZifIsBool(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsInt
var DefZifIsInt = def.DefFunc("is_int", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsInt(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsInt
var DefZifIsInteger = def.DefFunc("is_integer", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsInt(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsInt
var DefZifIsLong = def.DefFunc("is_long", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsInt(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsFloat
var DefZifIsFloat = def.DefFunc("is_float", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsFloat(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsFloat
var DefZifIsDouble = def.DefFunc("is_double", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsFloat(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsString
var DefZifIsString = def.DefFunc("is_string", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsString(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsArray
var DefZifIsArray = def.DefFunc("is_array", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsArray(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsObject
var DefZifIsObject = def.DefFunc("is_object", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsObject(var_)
	returnValue.SetBool(ret)
})

// generate by ZifIsNumeric
var DefZifIsNumeric = def.DefFunc("is_numeric", 1, 1, []def.ArgInfo{{Name: "value"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	value := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsNumeric(value)
	returnValue.SetBool(ret)
})

// generate by ZifIsScalar
var DefZifIsScalar = def.DefFunc("is_scalar", 1, 1, []def.ArgInfo{{Name: "value"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	value := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIsScalar(value)
	returnValue.SetBool(ret)
})

// generate by ZifUniqid
var DefZifUniqid = def.DefFunc("uniqid", 0, 2, []def.ArgInfo{{Name: "prefix"}, {Name: "more_entropy"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 2, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	prefix := fp.ParseString()
	more_entropy := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifUniqid(nil, prefix, more_entropy)
	returnValue.SetString(ret)
})

// generate by ZifParseUrl
var DefZifParseUrl = def.DefFunc("parse_url", 1, 2, []def.ArgInfo{{Name: "url"}, {Name: "component"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	url := fp.ParseString()
	fp.StartOptional()
	component := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifParseUrl(executeData.Ctx(), url, nil, component)
	returnValue.SetBy(ret)
})

// generate by ZifUrlencode
var DefZifUrlencode = def.DefFunc("urlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUrlencode(str)
	returnValue.SetString(ret)
})

// generate by ZifUrldecode
var DefZifUrldecode = def.DefFunc("urldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUrldecode(str)
	returnValue.SetString(ret)
})

// generate by ZifRawurlencode
var DefZifRawurlencode = def.DefFunc("rawurlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifRawurlencode(str)
	returnValue.SetString(ret)
})

// generate by ZifRawurldecode
var DefZifRawurldecode = def.DefFunc("rawurldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifRawurldecode(str)
	returnValue.SetString(ret)
})

// generate by ZifConvertUuencode
var DefZifConvertUuencode = def.DefFunc("convert_uuencode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	data := fp.ParseString()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifConvertUuencode(data)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifConvertUudecode
var DefZifConvertUudecode = def.DefFunc("convert_uudecode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	data := fp.ParseString()
	if fp.HasError() {
		returnValue.SetFalse()
		return
	}
	ret, ok := ZifConvertUudecode(executeData.Ctx(), data)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "vars", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	vars := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifVarDump(executeData.Ctx(), vars)
})

// generate by ZifDebugZvalDump
var DefZifDebugZvalDump = def.DefFunc("debug_zval_dump", 0, -1, []def.ArgInfo{{Name: "vars", Variadic: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, -1, 0)
	fp.CheckNumArgs()
	vars := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifDebugZvalDump(executeData.Ctx(), vars)
})

// generate by ZifVarExport
var DefZifVarExport = def.DefFunc("var_export", 1, 2, []def.ArgInfo{{Name: "value"}, {Name: "return"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 2, 0)
	fp.CheckNumArgs()
	value := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseBool()
	if fp.HasError() {
		return
	}
	ret := ZifVarExport(executeData.Ctx(), value, nil, return_)
	returnValue.SetBy(ret)
})

// generate by ZifSerialize
var DefZifSerialize = def.DefFunc("serialize", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifSerialize(executeData.Ctx(), var_)
	returnValue.SetBy(ret)
})

// generate by ZifVersionCompare
var DefZifVersionCompare = def.DefFunc("version_compare", 2, 3, []def.ArgInfo{{Name: "ver1"}, {Name: "ver2"}, {Name: "oper"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 3, 0)
	fp.CheckNumArgs()
	ver1 := fp.ParseString()
	ver2 := fp.ParseString()
	fp.StartOptional()
	oper := fp.ParseStringNullable()
	if fp.HasError() {
		return
	}
	ret := ZifVersionCompare(ver1, ver2, nil, oper)
	returnValue.SetByPtr(ret)
})

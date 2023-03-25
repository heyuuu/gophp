package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifKrsort
var DefZifKrsort = def.DefFunc("krsort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifKrsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifKsort
var DefZifKsort = def.DefFunc("ksort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifKsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifCount
var DefZifCount = def.DefFunc("count", 1, 2, []def.ArgInfo{{Name: "var_"}, {Name: "mode"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifCount(returnValue, var_, nil, mode)
})

// generate by ZifCount
var DefZifSizeof = def.DefFunc("sizeof", 1, 2, []def.ArgInfo{{Name: "var_"}, {Name: "mode"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifCount(returnValue, var_, nil, mode)
})

// generate by ZifNatsort
var DefZifNatsort = def.DefFunc("natsort", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifNatsort(executeData, returnValue, arg)
})

// generate by ZifNatcasesort
var DefZifNatcasesort = def.DefFunc("natcasesort", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifNatcasesort(executeData, returnValue, arg)
})

// generate by ZifAsort
var DefZifAsort = def.DefFunc("asort", 3, 3, []def.ArgInfo{{Name: "arg"}, {Name: "_"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arg := fp.ParseZvalEx(false, true)
	_ := fp.ParseArrayEx(false, true)
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifAsort(arg, _, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifArsort
var DefZifArsort = def.DefFunc("arsort", 3, 3, []def.ArgInfo{{Name: "arg"}, {Name: "_"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arg := fp.ParseZvalEx(false, true)
	_ := fp.ParseArrayEx(false, true)
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifArsort(arg, _, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifSort
var DefZifSort = def.DefFunc("sort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifSort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifRsort
var DefZifRsort = def.DefFunc("rsort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifUsort
var DefZifUsort = def.DefFunc("usort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	cmp_function := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUsort(executeData, returnValue, arg, cmp_function)
})

// generate by ZifUasort
var DefZifUasort = def.DefFunc("uasort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	cmp_function := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUasort(executeData, returnValue, arg, cmp_function)
})

// generate by ZifUksort
var DefZifUksort = def.DefFunc("uksort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	cmp_function := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUksort(executeData, returnValue, arg, cmp_function)
})

// generate by ZifEnd
var DefZifEnd = def.DefFunc("end", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifEnd(executeData, returnValue, arg)
})

// generate by ZifPrev
var DefZifPrev = def.DefFunc("prev", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifPrev(executeData, returnValue, arg)
})

// generate by ZifNext
var DefZifNext = def.DefFunc("next", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifNext(executeData, returnValue, arg)
})

// generate by ZifReset
var DefZifReset = def.DefFunc("reset", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifReset(executeData, returnValue, arg)
})

// generate by ZifCurrent
var DefZifCurrent = def.DefFunc("current", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCurrent(executeData, returnValue, arg)
})

// generate by ZifCurrent
var DefZifPos = def.DefFunc("pos", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCurrent(executeData, returnValue, arg)
})

// generate by ZifKey
var DefZifKey = def.DefFunc("key", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifKey(executeData, returnValue, arg)
})

// generate by ZifMin
var DefZifMin = def.DefFunc("min", 0, -1, []def.ArgInfo{{Name: "args"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifMin(executeData, returnValue, args)
})

// generate by ZifMax
var DefZifMax = def.DefFunc("max", 0, -1, []def.ArgInfo{{Name: "args"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifMax(executeData, returnValue, args)
})

// generate by ZifArrayWalk
var DefZifArrayWalk = def.DefFunc("array_walk", 2, 3, []def.ArgInfo{{Name: "input"}, {Name: "funcname"}, {Name: "userdata"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	input := fp.ParseZvalEx(false, true)
	funcname := fp.ParseZval()
	fp.StartOptional()
	userdata := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayWalk(executeData, returnValue, input, funcname, nil, userdata)
})

// generate by ZifArrayWalkRecursive
var DefZifArrayWalkRecursive = def.DefFunc("array_walk_recursive", 2, 3, []def.ArgInfo{{Name: "input"}, {Name: "funcname"}, {Name: "userdata"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	input := fp.ParseZvalEx(false, true)
	funcname := fp.ParseZval()
	fp.StartOptional()
	userdata := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayWalkRecursive(executeData, returnValue, input, funcname, nil, userdata)
})

// generate by ZifInArray
var DefZifInArray = def.DefFunc("in_array", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	needle := fp.ParseZval()
	haystack := fp.ParseZval()
	fp.StartOptional()
	strict := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifInArray(executeData, returnValue, needle, haystack, nil, strict)
})

// generate by ZifArraySearch
var DefZifArraySearch = def.DefFunc("array_search", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	needle := fp.ParseZval()
	haystack := fp.ParseZval()
	fp.StartOptional()
	strict := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySearch(executeData, returnValue, needle, haystack, nil, strict)
})

// generate by ZifCompact
var DefZifCompact = def.DefFunc("compact", 0, -1, []def.ArgInfo{{Name: "var_names"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	var_names := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifCompact(executeData, returnValue, var_names)
})

// generate by ZifArrayFill
var DefZifArrayFill = def.DefFunc("array_fill", 3, 3, []def.ArgInfo{{Name: "start_key"}, {Name: "num"}, {Name: "val"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	start_key := fp.ParseZval()
	num := fp.ParseZval()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayFill(executeData, returnValue, start_key, num, val)
})

// generate by ZifArrayFillKeys
var DefZifArrayFillKeys = def.DefFunc("array_fill_keys", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "val"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	keys := fp.ParseZval()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayFillKeys(executeData, returnValue, keys, val)
})

// generate by ZifRange
var DefZifRange = def.DefFunc("range", 2, 3, []def.ArgInfo{{Name: "low"}, {Name: "high"}, {Name: "step"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	low := fp.ParseZval()
	high := fp.ParseZval()
	fp.StartOptional()
	step := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRange(executeData, returnValue, low, high, nil, step)
})

// generate by ZifShuffle
var DefZifShuffle = def.DefFunc("shuffle", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifShuffle(executeData, returnValue, arg)
})

// generate by ZifArrayPush
var DefZifArrayPush = def.DefFunc("array_push", 1, -1, []def.ArgInfo{{Name: "stack"}, {Name: "vars"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayPush(executeData, returnValue, stack, nil, vars)
})

// generate by ZifArrayPop
var DefZifArrayPop = def.DefFunc("array_pop", 1, 1, []def.ArgInfo{{Name: "stack"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stack := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifArrayPop(executeData, returnValue, stack)
})

// generate by ZifArrayShift
var DefZifArrayShift = def.DefFunc("array_shift", 1, 1, []def.ArgInfo{{Name: "stack"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stack := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifArrayShift(executeData, returnValue, stack)
})

// generate by ZifArrayUnshift
var DefZifArrayUnshift = def.DefFunc("array_unshift", 1, -1, []def.ArgInfo{{Name: "stack"}, {Name: "vars"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayUnshift(executeData, returnValue, stack, nil, vars)
})

// generate by ZifArraySplice
var DefZifArraySplice = def.DefFunc("array_splice", 2, 4, []def.ArgInfo{{Name: "arg"}, {Name: "offset"}, {Name: "length"}, {Name: "replacement"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	arg := fp.ParseZvalEx(false, true)
	offset := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	replacement := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySplice(executeData, returnValue, arg, offset, nil, length, replacement)
})

// generate by ZifArraySlice
var DefZifArraySlice = def.DefFunc("array_slice", 2, 4, []def.ArgInfo{{Name: "arg"}, {Name: "offset"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	arg := fp.ParseZval()
	offset := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	preserve_keys := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySlice(executeData, returnValue, arg, offset, nil, length, preserve_keys)
})

// generate by ZifArrayMerge
var DefZifArrayMerge = def.DefFunc("array_merge", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMerge(executeData, returnValue, nil, arrays)
})

// generate by ZifArrayMergeRecursive
var DefZifArrayMergeRecursive = def.DefFunc("array_merge_recursive", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMergeRecursive(executeData, returnValue, nil, arrays)
})

// generate by ZifArrayReplace
var DefZifArrayReplace = def.DefFunc("array_replace", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayReplace(executeData, returnValue, arr1, nil, arrays)
})

// generate by ZifArrayReplaceRecursive
var DefZifArrayReplaceRecursive = def.DefFunc("array_replace_recursive", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayReplaceRecursive(executeData, returnValue, arr1, nil, arrays)
})

// generate by ZifArrayKeys
var DefZifArrayKeys = def.DefFunc("array_keys", 1, 3, []def.ArgInfo{{Name: "arg"}, {Name: "search_value"}, {Name: "strict"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	arg := fp.ParseZval()
	fp.StartOptional()
	search_value := fp.ParseZval()
	strict := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeys(executeData, returnValue, arg, nil, search_value, strict)
})

// generate by ZifArrayKeyFirst
var DefZifArrayKeyFirst = def.DefFunc("array_key_first", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyFirst(executeData, returnValue, arg)
})

// generate by ZifArrayKeyLast
var DefZifArrayKeyLast = def.DefFunc("array_key_last", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyLast(executeData, returnValue, arg)
})

// generate by ZifArrayValues
var DefZifArrayValues = def.DefFunc("array_values", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayValues(executeData, returnValue, arg)
})

// generate by ZifArrayCountValues
var DefZifArrayCountValues = def.DefFunc("array_count_values", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayCountValues(executeData, returnValue, arg)
})

// generate by ZifArrayColumn
var DefZifArrayColumn = def.DefFunc("array_column", 2, 3, []def.ArgInfo{{Name: "arg"}, {Name: "column_key"}, {Name: "index_key"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	arg := fp.ParseZval()
	column_key := fp.ParseZval()
	fp.StartOptional()
	index_key := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayColumn(executeData, returnValue, arg, column_key, nil, index_key)
})

// generate by ZifArrayReverse
var DefZifArrayReverse = def.DefFunc("array_reverse", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "preserve_keys"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	preserve_keys := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayReverse(executeData, returnValue, input, nil, preserve_keys)
})

// generate by ZifArrayPad
var DefZifArrayPad = def.DefFunc("array_pad", 3, 3, []def.ArgInfo{{Name: "arg"}, {Name: "pad_size"}, {Name: "pad_value"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arg := fp.ParseZval()
	pad_size := fp.ParseZval()
	pad_value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayPad(executeData, returnValue, arg, pad_size, pad_value)
})

// generate by ZifArrayFlip
var DefZifArrayFlip = def.DefFunc("array_flip", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayFlip(executeData, returnValue, arg)
})

// generate by ZifArrayChangeKeyCase
var DefZifArrayChangeKeyCase = def.DefFunc("array_change_key_case", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "case_"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	case_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayChangeKeyCase(executeData, returnValue, input, nil, case_)
})

// generate by ZifArrayUnique
var DefZifArrayUnique = def.DefFunc("array_unique", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUnique(executeData, returnValue, arg, nil, flags)
})

// generate by ZifArrayIntersectKey
var DefZifArrayIntersectKey = def.DefFunc("array_intersect_key", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectKey(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayIntersectUkey
var DefZifArrayIntersectUkey = def.DefFunc("array_intersect_ukey", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_key_compare_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectUkey(executeData, returnValue, arr1, arr2, callback_key_compare_func)
})

// generate by ZifArrayIntersect
var DefZifArrayIntersect = def.DefFunc("array_intersect", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersect(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayUintersect
var DefZifArrayUintersect = def.DefFunc("array_uintersect", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_compare_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_compare_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUintersect(executeData, returnValue, arr1, arr2, callback_data_compare_func)
})

// generate by ZifArrayIntersectAssoc
var DefZifArrayIntersectAssoc = def.DefFunc("array_intersect_assoc", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectAssoc(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayIntersectUassoc
var DefZifArrayIntersectUassoc = def.DefFunc("array_intersect_uassoc", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_key_compare_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectUassoc(executeData, returnValue, arr1, arr2, callback_key_compare_func)
})

// generate by ZifArrayUintersectAssoc
var DefZifArrayUintersectAssoc = def.DefFunc("array_uintersect_assoc", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_compare_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_compare_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUintersectAssoc(executeData, returnValue, arr1, arr2, callback_data_compare_func)
})

// generate by ZifArrayUintersectUassoc
var DefZifArrayUintersectUassoc = def.DefFunc("array_uintersect_uassoc", 4, 4, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_compare_func"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 4, 4, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_compare_func := fp.ParseZval()
	callback_key_compare_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUintersectUassoc(executeData, returnValue, arr1, arr2, callback_data_compare_func, callback_key_compare_func)
})

// generate by ZifArrayDiffKey
var DefZifArrayDiffKey = def.DefFunc("array_diff_key", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiffKey(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayDiffUkey
var DefZifArrayDiffUkey = def.DefFunc("array_diff_ukey", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_key_comp_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayDiffUkey(executeData, returnValue, arr1, arr2, callback_key_comp_func)
})

// generate by ZifArrayDiff
var DefZifArrayDiff = def.DefFunc("array_diff", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiff(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayUdiff
var DefZifArrayUdiff = def.DefFunc("array_udiff", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_comp_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_comp_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUdiff(executeData, returnValue, arr1, arr2, callback_data_comp_func)
})

// generate by ZifArrayDiffAssoc
var DefZifArrayDiffAssoc = def.DefFunc("array_diff_assoc", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiffAssoc(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayDiffUassoc
var DefZifArrayDiffUassoc = def.DefFunc("array_diff_uassoc", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_comp_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_comp_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayDiffUassoc(executeData, returnValue, arr1, arr2, callback_data_comp_func)
})

// generate by ZifArrayUdiffAssoc
var DefZifArrayUdiffAssoc = def.DefFunc("array_udiff_assoc", 3, 3, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_key_comp_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUdiffAssoc(executeData, returnValue, arr1, arr2, callback_key_comp_func)
})

// generate by ZifArrayUdiffUassoc
var DefZifArrayUdiffUassoc = def.DefFunc("array_udiff_uassoc", 4, 4, []def.ArgInfo{{Name: "arr1"}, {Name: "arr2"}, {Name: "callback_data_comp_func"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 4, 4, 0)
	arr1 := fp.ParseZval()
	arr2 := fp.ParseZval()
	callback_data_comp_func := fp.ParseZval()
	callback_key_comp_func := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayUdiffUassoc(executeData, returnValue, arr1, arr2, callback_data_comp_func, callback_key_comp_func)
})

// generate by ZifArrayRand
var DefZifArrayRand = def.DefFunc("array_rand", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "num_req"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZval()
	fp.StartOptional()
	num_req := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayRand(executeData, returnValue, arg, nil, num_req)
})

// generate by ZifArraySum
var DefZifArraySum = def.DefFunc("array_sum", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySum(executeData, returnValue, arg)
})

// generate by ZifArrayProduct
var DefZifArrayProduct = def.DefFunc("array_product", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayProduct(executeData, returnValue, arg)
})

// generate by ZifArrayReduce
var DefZifArrayReduce = def.DefFunc("array_reduce", 2, 3, []def.ArgInfo{{Name: "arg"}, {Name: "callback"}, {Name: "initial"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	arg := fp.ParseZval()
	callback := fp.ParseZval()
	fp.StartOptional()
	initial := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayReduce(executeData, returnValue, arg, callback, nil, initial)
})

// generate by ZifArrayFilter
var DefZifArrayFilter = def.DefFunc("array_filter", 1, 3, []def.ArgInfo{{Name: "arg"}, {Name: "callback"}, {Name: "use_keys"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	arg := fp.ParseZval()
	fp.StartOptional()
	callback := fp.ParseZval()
	use_keys := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayFilter(executeData, returnValue, arg, nil, callback, use_keys)
})

// generate by ZifArrayMap
var DefZifArrayMap = def.DefFunc("array_map", 1, -1, []def.ArgInfo{{Name: "callback"}, {Name: "arrays"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	callback := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMap(executeData, returnValue, callback, arrays)
})

// generate by ZifArrayKeyExists
var DefZifArrayKeyExists = def.DefFunc("array_key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "search"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	search := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyExists(executeData, returnValue, key, search)
})

// generate by ZifArrayKeyExists
var DefZifKeyExists = def.DefFunc("key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "search"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	search := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyExists(executeData, returnValue, key, search)
})

// generate by ZifArrayChunk
var DefZifArrayChunk = def.DefFunc("array_chunk", 2, 3, []def.ArgInfo{{Name: "arg"}, {Name: "size"}, {Name: "preserve_keys"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	arg := fp.ParseZval()
	size := fp.ParseZval()
	fp.StartOptional()
	preserve_keys := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayChunk(executeData, returnValue, arg, size, nil, preserve_keys)
})

// generate by ZifArrayCombine
var DefZifArrayCombine = def.DefFunc("array_combine", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "values"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	keys := fp.ParseZval()
	values := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayCombine(executeData, returnValue, keys, values)
})

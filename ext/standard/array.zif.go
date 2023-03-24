package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifKsort
var DefZifKsort = def.DefFunc("ksort", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "sort_flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifAsort
var DefZifAsort = def.DefFunc("asort", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "sort_flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAsort(executeData, returnValue, arg, nil, sort_flags)
})

// generate by ZifArsort
var DefZifArsort = def.DefFunc("arsort", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "sort_flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArsort(executeData, returnValue, arg, nil, sort_flags)
})

// generate by ZifSort
var DefZifSort = def.DefFunc("sort", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "sort_flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSort(executeData, returnValue, arg, nil, sort_flags)
})

// generate by ZifRsort
var DefZifRsort = def.DefFunc("rsort", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "sort_flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRsort(executeData, returnValue, arg, nil, sort_flags)
})

// generate by ZifMin
var DefZifMin = def.DefFunc("min", -1, -1, []def.ArgInfo{{name: "args"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifMin(executeData, returnValue, args)
})

// generate by ZifMax
var DefZifMax = def.DefFunc("max", -1, -1, []def.ArgInfo{{name: "args"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifMax(executeData, returnValue, args)
})

// generate by ZifArrayWalk
var DefZifArrayWalk = def.DefFunc("array_walk", 2, 3, []def.ArgInfo{{name: "input"}, {name: "funcname"}, {name: "userdata"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayWalkRecursive = def.DefFunc("array_walk_recursive", 2, 3, []def.ArgInfo{{name: "input"}, {name: "funcname"}, {name: "userdata"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifInArray = def.DefFunc("in_array", 2, 3, []def.ArgInfo{{name: "needle"}, {name: "haystack"}, {name: "strict"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArraySearch = def.DefFunc("array_search", 2, 3, []def.ArgInfo{{name: "needle"}, {name: "haystack"}, {name: "strict"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifCompact = def.DefFunc("compact", -1, -1, []def.ArgInfo{{name: "var_names"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	var_names := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifCompact(executeData, returnValue, var_names)
})

// generate by ZifRange
var DefZifRange = def.DefFunc("range", 2, 3, []def.ArgInfo{{name: "low"}, {name: "high"}, {name: "step"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifArrayPush
var DefZifArrayPush = def.DefFunc("array_push", 1, -1, []def.ArgInfo{{name: "stack"}, {name: "vars"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayPush(executeData, returnValue, stack, nil, vars)
})

// generate by ZifArrayUnshift
var DefZifArrayUnshift = def.DefFunc("array_unshift", 1, -1, []def.ArgInfo{{name: "stack"}, {name: "vars"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArraySplice = def.DefFunc("array_splice", 2, 4, []def.ArgInfo{{name: "arg"}, {name: "offset"}, {name: "length"}, {name: "replacement"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArraySlice = def.DefFunc("array_slice", 2, 4, []def.ArgInfo{{name: "arg"}, {name: "offset"}, {name: "length"}, {name: "preserve_keys"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayMerge = def.DefFunc("array_merge", 0, -1, []def.ArgInfo{{name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifArrayMerge(executeData, returnValue, nil, arrays)
})

// generate by ZifArrayMergeRecursive
var DefZifArrayMergeRecursive = def.DefFunc("array_merge_recursive", 0, -1, []def.ArgInfo{{name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifArrayMergeRecursive(executeData, returnValue, nil, arrays)
})

// generate by ZifArrayReplace
var DefZifArrayReplace = def.DefFunc("array_replace", 1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayReplaceRecursive = def.DefFunc("array_replace_recursive", 1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayKeys = def.DefFunc("array_keys", 1, 3, []def.ArgInfo{{name: "arg"}, {name: "search_value"}, {name: "strict"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifArrayColumn
var DefZifArrayColumn = def.DefFunc("array_column", 2, 3, []def.ArgInfo{{name: "arg"}, {name: "column_key"}, {name: "index_key"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayReverse = def.DefFunc("array_reverse", 1, 2, []def.ArgInfo{{name: "input"}, {name: "preserve_keys"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	preserve_keys := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayReverse(executeData, returnValue, input, nil, preserve_keys)
})

// generate by ZifArrayChangeKeyCase
var DefZifArrayChangeKeyCase = def.DefFunc("array_change_key_case", 1, 2, []def.ArgInfo{{name: "input"}, {name: "case_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayUnique = def.DefFunc("array_unique", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayIntersectKey = def.DefFunc("array_intersect_key", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectKey(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayIntersect
var DefZifArrayIntersect = def.DefFunc("array_intersect", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersect(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayIntersectAssoc
var DefZifArrayIntersectAssoc = def.DefFunc("array_intersect_assoc", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayIntersectAssoc(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayDiffKey
var DefZifArrayDiffKey = def.DefFunc("array_diff_key", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiffKey(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayDiff
var DefZifArrayDiff = def.DefFunc("array_diff", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiff(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayDiffAssoc
var DefZifArrayDiffAssoc = def.DefFunc("array_diff_assoc", -1, -1, []def.ArgInfo{{name: "arr1"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiffAssoc(executeData, returnValue, arr1, arrays)
})

// generate by ZifArrayRand
var DefZifArrayRand = def.DefFunc("array_rand", 1, 2, []def.ArgInfo{{name: "arg"}, {name: "num_req"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZval()
	fp.StartOptional()
	num_req := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayRand(executeData, returnValue, arg, nil, num_req)
})

// generate by ZifArrayReduce
var DefZifArrayReduce = def.DefFunc("array_reduce", 2, 3, []def.ArgInfo{{name: "arg"}, {name: "callback"}, {name: "initial"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayFilter = def.DefFunc("array_filter", 1, 3, []def.ArgInfo{{name: "arg"}, {name: "callback"}, {name: "use_keys"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifArrayMap = def.DefFunc("array_map", -1, -1, []def.ArgInfo{{name: "callback"}, {name: "arrays"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, -1, -1, 0)
	callback := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMap(executeData, returnValue, callback, arrays)
})

// generate by ZifArrayChunk
var DefZifArrayChunk = def.DefFunc("array_chunk", 2, 3, []def.ArgInfo{{name: "arg"}, {name: "size"}, {name: "preserve_keys"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

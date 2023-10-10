package array

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifKrsort
var DefZifKrsort = def.DefFunc("krsort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifKsort = def.DefFunc("ksort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifCount = def.DefFunc("count", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifCount(var_, nil, mode)
	returnValue.SetLong(ret)
})

// generate by ZifCount
var DefZifSizeof = def.DefFunc("sizeof", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifCount(var_, nil, mode)
	returnValue.SetLong(ret)
})

// generate by ZifNatsort
var DefZifNatsort = def.DefFunc("natsort", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifNatsort(arg)
	returnValue.SetBool(ret)
})

// generate by ZifNatcasesort
var DefZifNatcasesort = def.DefFunc("natcasesort", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifNatcasesort(arg)
	returnValue.SetBool(ret)
})

// generate by ZifAsort
var DefZifAsort = def.DefFunc("asort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifAsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifArsort
var DefZifArsort = def.DefFunc("arsort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	sort_flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifArsort(arg, nil, sort_flags)
	returnValue.SetBool(ret)
})

// generate by ZifSort
var DefZifSort = def.DefFunc("sort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifRsort = def.DefFunc("rsort", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "sort_flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifUsort = def.DefFunc("usort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseArrayEx2(false, true, false)
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUsort(arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifUasort
var DefZifUasort = def.DefFunc("uasort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseArrayEx2(false, true, false)
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUasort(arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifUksort
var DefZifUksort = def.DefFunc("uksort", 2, 2, []def.ArgInfo{{Name: "arg"}, {Name: "cmp_function"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	arg := fp.ParseArrayEx2(false, true, false)
	cmp_function := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret := ZifUksort(arg, cmp_function)
	returnValue.SetBool(ret)
})

// generate by ZifEnd
var DefZifEnd = def.DefFunc("end", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHtEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifEnd(array)
	returnValue.SetBy(ret)
})

// generate by ZifPrev
var DefZifPrev = def.DefFunc("prev", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHtEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifPrev(array)
	returnValue.SetBy(ret)
})

// generate by ZifNext
var DefZifNext = def.DefFunc("next", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHtEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifNext(array)
	returnValue.SetBy(ret)
})

// generate by ZifReset
var DefZifReset = def.DefFunc("reset", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHtEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifReset(array)
	returnValue.SetBy(ret)
})

// generate by ZifCurrent
var DefZifCurrent = def.DefFunc("current", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
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
var DefZifPos = def.DefFunc("pos", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
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
var DefZifKey = def.DefFunc("key", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifKey(array)
	returnValue.SetBy(ret)
})

// generate by ZifMin
var DefZifMin = def.DefFunc("min", 1, -1, []def.ArgInfo{{Name: "arg"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arg := fp.ParseZval()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifMin(arg, args)
	returnValue.SetBy(ret)
})

// generate by ZifMax
var DefZifMax = def.DefFunc("max", 1, -1, []def.ArgInfo{{Name: "arg"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arg := fp.ParseZval()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifMax(arg, args)
	returnValue.SetBy(ret)
})

// generate by ZifArrayWalk
var DefZifArrayWalk = def.DefFunc("array_walk", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "callable"}, {Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	array := fp.ParseArrayOrObjectEx(false, true)
	callable := fp.ParseCallable()
	fp.StartOptional()
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayWalk(array, callable, nil, arg)
	returnValue.SetBool(ret)
})

// generate by ZifArrayWalkRecursive
var DefZifArrayWalkRecursive = def.DefFunc("array_walk_recursive", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "callable"}, {Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	array := fp.ParseArrayOrObjectEx(false, true)
	callable := fp.ParseCallable()
	fp.StartOptional()
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayWalkRecursive(array, callable, nil, arg)
	returnValue.SetBool(ret)
})

// generate by ZifInArray
var DefZifInArray = def.DefFunc("in_array", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	needle := fp.ParseZval()
	haystack := fp.ParseArrayHt()
	fp.StartOptional()
	strict := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifInArray(needle, haystack, nil, strict)
	returnValue.SetBool(ret)
})

// generate by ZifArraySearch
var DefZifArraySearch = def.DefFunc("array_search", 2, 3, []def.ArgInfo{{Name: "needle"}, {Name: "haystack"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	needle := fp.ParseZval()
	haystack := fp.ParseArrayHt()
	fp.StartOptional()
	strict := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifArraySearch(needle, haystack, nil, strict)
	returnValue.SetBy(ret)
})

// generate by ZifArrayFill
var DefZifArrayFill = def.DefFunc("array_fill", 3, 3, []def.ArgInfo{{Name: "start_key"}, {Name: "num"}, {Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	start_key := fp.ParseLong()
	num := fp.ParseLong()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayFill(start_key, num, val)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayFillKeys
var DefZifArrayFillKeys = def.DefFunc("array_fill_keys", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	keys := fp.ParseArrayHt()
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFillKeys(keys, val)
	returnValue.SetArray(ret)
})

// generate by ZifRange
var DefZifRange = def.DefFunc("range", 2, 3, []def.ArgInfo{{Name: "low"}, {Name: "high"}, {Name: "step"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	low_ := fp.ParseZval()
	high_ := fp.ParseZval()
	fp.StartOptional()
	step_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifRange(low_, high_, nil, step_)
	if ok {
		returnValue.SetArrayOfZval(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifShuffle
var DefZifShuffle = def.DefFunc("shuffle", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifShuffle(arg)
	returnValue.SetBool(ret)
})

// generate by ZifArrayPush
var DefZifArrayPush = def.DefFunc("array_push", 1, -1, []def.ArgInfo{{Name: "stack"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseArrayEx(false, true)
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayPush(stack, nil, args)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayPop
var DefZifArrayPop = def.DefFunc("array_pop", 1, 1, []def.ArgInfo{{Name: "stack"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stack := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifArrayPop(stack)
	returnValue.SetBy(ret)
})

// generate by ZifArrayShift
var DefZifArrayShift = def.DefFunc("array_shift", 1, 1, []def.ArgInfo{{Name: "stack"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stack := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ret := ZifArrayShift(stack)
	returnValue.SetBy(ret)
})

// generate by ZifArrayUnshift
var DefZifArrayUnshift = def.DefFunc("array_unshift", 1, -1, []def.ArgInfo{{Name: "stack"}, {Name: "values"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseZvalEx(false, true)
	values := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayUnshift(stack, values)
	returnValue.SetLong(ret)
})

// generate by ZifArraySplice
var DefZifArraySplice = def.DefFunc("array_splice", 2, 4, []def.ArgInfo{{Name: "array"}, {Name: "offset"}, {Name: "length"}, {Name: "replacement"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	array := fp.ParseArrayEx(false, true)
	offset := fp.ParseLong()
	fp.StartOptional()
	length_ := fp.ParseLongNullable()
	replacement := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArraySplice(array, offset, nil, length_, replacement)
	returnValue.SetArray(ret)
})

// generate by ZifArraySlice
var DefZifArraySlice = def.DefFunc("array_slice", 2, 4, []def.ArgInfo{{Name: "array"}, {Name: "offset"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	array := fp.ParseArrayHt()
	offset := fp.ParseLong()
	fp.StartOptional()
	length_ := fp.ParseZval()
	preserve_keys := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifArraySlice(array, offset, nil, length_, preserve_keys)
	returnValue.SetArray(ret)
})

// generate by ZifArrayMerge
var DefZifArrayMerge = def.DefFunc("array_merge", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMerge(returnValue, nil, arrays)
})

// generate by ZifArrayMergeRecursive
var DefZifArrayMergeRecursive = def.DefFunc("array_merge_recursive", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	fp.StartOptional()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMergeRecursive(returnValue, nil, arrays)
})

// generate by ZifArrayReplace
var DefZifArrayReplace = def.DefFunc("array_replace", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReplace(arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayReplaceRecursive
var DefZifArrayReplaceRecursive = def.DefFunc("array_replace_recursive", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReplaceRecursive(arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeys
var DefZifArrayKeys = def.DefFunc("array_keys", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "search_value"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	array := fp.ParseArrayHt()
	fp.StartOptional()
	search_value := fp.ParseZval()
	strict := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeys(array, nil, search_value, strict)
	returnValue.SetArray(ret)
})

// generate by ZifArrayKeyFirst
var DefZifArrayKeyFirst = def.DefFunc("array_key_first", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyFirst(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeyLast
var DefZifArrayKeyLast = def.DefFunc("array_key_last", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyLast(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayValues
var DefZifArrayValues = def.DefFunc("array_values", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayValues(array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayCountValues
var DefZifArrayCountValues = def.DefFunc("array_count_values", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayCountValues(array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayColumn
var DefZifArrayColumn = def.DefFunc("array_column", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "column_key"}, {Name: "index_key"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	array := fp.ParseArrayHt()
	column_key := fp.ParseZvalEx(true, false)
	fp.StartOptional()
	index_key := fp.ParseZvalEx(true, false)
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayColumn(array, column_key, nil, index_key)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayReverse
var DefZifArrayReverse = def.DefFunc("array_reverse", 1, 2, []def.ArgInfo{{Name: "array"}, {Name: "preserve_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	array := fp.ParseArrayHt()
	fp.StartOptional()
	preserve_keys := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReverse(array, nil, preserve_keys)
	returnValue.SetArray(ret)
})

// generate by ZifArrayPad
var DefZifArrayPad = def.DefFunc("array_pad", 3, 3, []def.ArgInfo{{Name: "array"}, {Name: "pad_size"}, {Name: "pad_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	array := fp.ParseArrayHt()
	pad_size := fp.ParseLong()
	pad_value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayPad(array, pad_size, pad_value)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayFlip
var DefZifArrayFlip = def.DefFunc("array_flip", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFlip(array)
	returnValue.SetArray(ret)
})

// generate by ZifArrayChangeKeyCase
var DefZifArrayChangeKeyCase = def.DefFunc("array_change_key_case", 1, 2, []def.ArgInfo{{Name: "array"}, {Name: "case"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	array := fp.ParseArrayHt()
	fp.StartOptional()
	case_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayChangeKeyCase(array, nil, case_)
	returnValue.SetArray(ret)
})

// generate by ZifArrayUnique
var DefZifArrayUnique = def.DefFunc("array_unique", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayHt()
	fp.StartOptional()
	flags := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayUnique(arg, nil, flags)
	returnValue.SetArray(ret)
})

// generate by ZifArrayIntersectKey
var DefZifArrayIntersectKey = def.DefFunc("array_intersect_key", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectKey(arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectUkey
var DefZifArrayIntersectUkey = def.DefFunc("array_intersect_ukey", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectUkey(arrays, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersect
var DefZifArrayIntersect = def.DefFunc("array_intersect", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersect(arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersect
var DefZifArrayUintersect = def.DefFunc("array_uintersect", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_compare_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_data_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersect(arrays, callback_data_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectAssoc
var DefZifArrayIntersectAssoc = def.DefFunc("array_intersect_assoc", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectAssoc(arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayIntersectUassoc
var DefZifArrayIntersectUassoc = def.DefFunc("array_intersect_uassoc", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayIntersectUassoc(arrays, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersectAssoc
var DefZifArrayUintersectAssoc = def.DefFunc("array_uintersect_assoc", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_compare_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_data_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersectAssoc(arrays, callback_data_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUintersectUassoc
var DefZifArrayUintersectUassoc = def.DefFunc("array_uintersect_uassoc", 2, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_compare_func"}, {Name: "callback_key_compare_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, -1, 0)
	arrays := fp.ParseVariadicEx(2)
	callback_data_compare_func := fp.ParseCallable()
	callback_key_compare_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUintersectUassoc(arrays, callback_data_compare_func, callback_key_compare_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffKey
var DefZifArrayDiffKey = def.DefFunc("array_diff_key", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffKey(arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffUkey
var DefZifArrayDiffUkey = def.DefFunc("array_diff_ukey", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffUkey(arrays, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiff
var DefZifArrayDiff = def.DefFunc("array_diff", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayDiff(arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayUdiff
var DefZifArrayUdiff = def.DefFunc("array_udiff", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_comp_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_data_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiff(arrays, callback_data_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffAssoc
var DefZifArrayDiffAssoc = def.DefFunc("array_diff_assoc", 0, -1, []def.ArgInfo{{Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffAssoc(arrays)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayDiffUassoc
var DefZifArrayDiffUassoc = def.DefFunc("array_diff_uassoc", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayDiffUassoc(arrays, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUdiffAssoc
var DefZifArrayUdiffAssoc = def.DefFunc("array_udiff_assoc", 1, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_comp_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arrays := fp.ParseVariadicEx(1)
	callback_data_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiffAssoc(arrays, callback_data_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayUdiffUassoc
var DefZifArrayUdiffUassoc = def.DefFunc("array_udiff_uassoc", 2, -1, []def.ArgInfo{{Name: "arrays"}, {Name: "callback_data_comp_func"}, {Name: "callback_key_comp_func"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, -1, 0)
	arrays := fp.ParseVariadicEx(2)
	callback_data_comp_func := fp.ParseCallable()
	callback_key_comp_func := fp.ParseCallable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayUdiffUassoc(arrays, callback_data_comp_func, callback_key_comp_func)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifArrayMultisort
var DefZifArrayMultisort = def.DefFunc("array_multisort", 0, -1, []def.ArgInfo{{Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayMultisort(args)
	returnValue.SetBool(ret)
})

// generate by ZifArrayRand
var DefZifArrayRand = def.DefFunc("array_rand", 1, 2, []def.ArgInfo{{Name: "arg"}, {Name: "num_req"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	arg := fp.ParseArrayHt()
	fp.StartOptional()
	num_req_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifArrayRand(arg, nil, num_req_)
	returnValue.SetBy(ret)
})

// generate by ZifArraySum
var DefZifArraySum = def.DefFunc("array_sum", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArraySum(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayProduct
var DefZifArrayProduct = def.DefFunc("array_product", 1, 1, []def.ArgInfo{{Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	array := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifArrayProduct(array)
	returnValue.SetBy(ret)
})

// generate by ZifArrayReduce
var DefZifArrayReduce = def.DefFunc("array_reduce", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "callback"}, {Name: "initial"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	array := fp.ParseArrayHt()
	callback := fp.ParseCallable()
	fp.StartOptional()
	initial := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifArrayReduce(array, callback, nil, initial)
	returnValue.SetBy(ret)
})

// generate by ZifArrayFilter
var DefZifArrayFilter = def.DefFunc("array_filter", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "callback"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	array_ := fp.ParseArrayHt()
	fp.StartOptional()
	callback := fp.ParseCallable()
	mode := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifArrayFilter(array_, nil, callback, mode)
	returnValue.SetArray(ret)
})

// generate by ZifArrayMap
var DefZifArrayMap = def.DefFunc("array_map", 1, -1, []def.ArgInfo{{Name: "callback"}, {Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	callback := fp.ParseCallable()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifArrayMap(callback, arrays)
	returnValue.SetBy(ret)
})

// generate by ZifArrayKeyExists
var DefZifArrayKeyExists = def.DefFunc("array_key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	array := fp.ParseArrayOrObject()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyExists(key, array)
	returnValue.SetBool(ret)
})

// generate by ZifArrayKeyExists
var DefZifKeyExists = def.DefFunc("key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	array := fp.ParseArrayOrObject()
	if fp.HasError() {
		return
	}
	ret := ZifArrayKeyExists(key, array)
	returnValue.SetBool(ret)
})

// generate by ZifArrayChunk
var DefZifArrayChunk = def.DefFunc("array_chunk", 2, 3, []def.ArgInfo{{Name: "array"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	array := fp.ParseArrayHt()
	length := fp.ParseLong()
	fp.StartOptional()
	preserve_keys := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifArrayChunk(array, length, nil, preserve_keys)
	returnValue.SetArray(ret)
})

// generate by ZifArrayCombine
var DefZifArrayCombine = def.DefFunc("array_combine", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "values"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	keys := fp.ParseArrayHt()
	values := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret, ok := ZifArrayCombine(keys, values)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifExtract
var DefZifExtract = def.DefFunc("extract", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "flags"}, {Name: "prefix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	array := fp.ParseArrayEx2(false, true, false)
	fp.StartOptional()
	flags := fp.ParseLong()
	prefix := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifExtract(array, nil, flags, prefix)
	returnValue.SetLong(ret)
})

// generate by ZifCompact
var DefZifCompact = def.DefFunc("compact", 0, -1, []def.ArgInfo{{Name: "var_names"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	var_names := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifCompact(var_names)
	returnValue.SetArray(ret)
})

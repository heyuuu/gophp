package array

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

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
	ret := ZifArrayPop(returnValue, stack)
	returnValue.SetBy(ret)
})

// generate by ZifArrayShift
var DefZifArrayShift = def.DefFunc("array_shift", 1, 1, []def.ArgInfo{{Name: "stack"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stack := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ZifArrayShift(stack)
})

// generate by ZifArrayUnshift
var DefZifArrayUnshift = def.DefFunc("array_unshift", 1, -1, []def.ArgInfo{{Name: "stack"}, {Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	stack := fp.ParseZvalEx(false, true)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayUnshift(executeData, returnValue, stack, vars)
})

// generate by ZifArraySplice
var DefZifArraySplice = def.DefFunc("array_splice", 2, 4, []def.ArgInfo{{Name: "arg"}, {Name: "offset"}, {Name: "length"}, {Name: "replacement"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	arg := fp.ParseArrayEx(false, true)
	offset := fp.ParseLong()
	fp.StartOptional()
	length_ := fp.ParseLongNullable()
	replacement := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySplice(returnValue, arg, offset, nil, length_, replacement)
})

// generate by ZifArraySlice
var DefZifArraySlice = def.DefFunc("array_slice", 2, 4, []def.ArgInfo{{Name: "arg"}, {Name: "offset"}, {Name: "length"}, {Name: "preserve_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifArrayReplace = def.DefFunc("array_replace", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifArrayReplaceRecursive = def.DefFunc("array_replace_recursive", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
	ret := ZifArrayCountValues(executeData, returnValue, array)
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
var DefZifArrayChangeKeyCase = def.DefFunc("array_change_key_case", 1, 2, []def.ArgInfo{{Name: "input"}, {Name: "case"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	input := fp.ParseZval()
	fp.StartOptional()
	case_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayChangeKeyCase(input, nil, case_)
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
var DefZifArrayDiff = def.DefFunc("array_diff", 1, -1, []def.ArgInfo{{Name: "arr1"}, {Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	arr1 := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayDiff(executeData, returnValue, arr1, arrays)
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
	ret := ZifArrayRand(returnValue, arg, nil, num_req_)
	returnValue.SetBy(ret)
})

// generate by ZifArraySum
var DefZifArraySum = def.DefFunc("array_sum", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArraySum(executeData, returnValue, arg)
})

// generate by ZifArrayProduct
var DefZifArrayProduct = def.DefFunc("array_product", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayProduct(executeData, returnValue, arg)
})

// generate by ZifArrayReduce
var DefZifArrayReduce = def.DefFunc("array_reduce", 2, 3, []def.ArgInfo{{Name: "arg"}, {Name: "callback"}, {Name: "initial"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifArrayFilter = def.DefFunc("array_filter", 1, 3, []def.ArgInfo{{Name: "arg"}, {Name: "callback"}, {Name: "use_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifArrayMap = def.DefFunc("array_map", 1, -1, []def.ArgInfo{{Name: "callback"}, {Name: "arrays"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	callback := fp.ParseZval()
	arrays := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifArrayMap(executeData, returnValue, callback, arrays)
})

// generate by ZifArrayKeyExists
var DefZifArrayKeyExists = def.DefFunc("array_key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "search"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	search := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyExists(executeData, returnValue, key, search)
})

// generate by ZifArrayKeyExists
var DefZifKeyExists = def.DefFunc("key_exists", 2, 2, []def.ArgInfo{{Name: "key"}, {Name: "search"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	key := fp.ParseZval()
	search := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayKeyExists(executeData, returnValue, key, search)
})

// generate by ZifArrayChunk
var DefZifArrayChunk = def.DefFunc("array_chunk", 2, 3, []def.ArgInfo{{Name: "arg"}, {Name: "size"}, {Name: "preserve_keys"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifArrayCombine = def.DefFunc("array_combine", 2, 2, []def.ArgInfo{{Name: "keys"}, {Name: "values"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	keys := fp.ParseZval()
	values := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifArrayCombine(executeData, returnValue, keys, values)
})

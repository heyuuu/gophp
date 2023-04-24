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
var DefZifEnd = def.DefFunc("end", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseArrayEx(false, true)
	if fp.HasError() {
		return
	}
	ZifEnd(executeData, returnValue, arg)
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
	ZifKey(executeData, returnValue, array)
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

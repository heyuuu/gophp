package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"sort"
	"strconv"
	"strings"
)

func reserveComparer(comparer types.ArrayComparer) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		return comparer(p2, p1)
	}
}

func valueComparer(comparer func(v1, v2 *types.Zval) int) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		v1 := p1.GetVal().DeIndirect()
		v2 := p2.GetVal().DeIndirect()
		return comparer(v1, v2)
	}
}

func keyComparer(comparer func(v1, v2 types.ArrayKey) int) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		return comparer(p1.GetKey(), p2.GetKey())
	}
}

func twiceComparer(compare1 types.ArrayComparer, compare2 types.ArrayComparer) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		c := compare1(p1, p2)
		if c != 0 {
			return c
		}

		return compare2(p1, p2)
	}
}

func arrayNaturalGeneralCompare(v1, v2 *types.Zval) int {
	str1 := zend.ZvalGetStrVal(v1)
	str2 := zend.ZvalGetStrVal(v2)
	return str.Strnatcmp(str1, str2, false)
}

func arrayNaturalGeneralCaseCompare(v1, v2 *types.Zval) int {
	str1 := zend.ZvalGetStrVal(v1)
	str2 := zend.ZvalGetStrVal(v2)
	return str.Strnatcmp(str1, str2, true)
}
func arrayDataCompare(v1, v2 *types.Zval) int {
	var result types.Zval
	if zend.CompareFunction(&result, v1, v2) == types.FAILURE {
		return 0
	}
	b.Assert(result.IsType(types.IS_LONG))
	return zend.ZEND_NORMALIZE_BOOL(result.Long())
}

func phpGetDataCompareFunc(sortType zend.ZendLong, reverse bool) types.ArrayComparer {
	var comparer types.ArrayComparer
	switch sortType & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		comparer = valueComparer(zend.NumericCompareFunction)
	case PHP_SORT_STRING:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = valueComparer(zend.StringCaseCompareFunction)
		} else {
			comparer = valueComparer(zend.StringCompareFunction)
		}
	case PHP_SORT_NATURAL:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = valueComparer(arrayNaturalGeneralCaseCompare)
		} else {
			comparer = valueComparer(arrayNaturalGeneralCompare)
		}
	case PHP_SORT_LOCALE_STRING:
		comparer = valueComparer(zend.StringLocaleCompareFunction)
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		comparer = valueComparer(arrayDataCompare)
	}
	if reverse {
		comparer = reserveComparer(comparer)
	}
	return comparer
}

func arrayKeyToDouble(k types.ArrayKey) float64 {
	if k.IsStrKey() {
		return zend.StrToDouble(k.StrKey())
	} else {
		return float64(k.IndexKey())
	}
}
func arrayKeyToString(k types.ArrayKey) string {
	if k.IsStrKey() {
		return k.StrKey()
	} else {
		return strconv.Itoa(k.IndexKey())
	}
}
func arrayKeyCompareNumeric(k1 types.ArrayKey, k2 types.ArrayKey) int {
	l1, _, isStr1 := k1.Keys()
	l2, _, isStr2 := k2.Keys()
	if !isStr1 && !isStr2 {
		return b.Cond(l1 < l2, -1, 1)
	} else {
		d1 := arrayKeyToDouble(k1)
		d2 := arrayKeyToDouble(k2)
		return b.Compare(d1, d2)
	}
}
func arrayKeyCompare(k1 types.ArrayKey, k2 types.ArrayKey) int {
	l1, s1, isStr1 := k1.Keys()
	l2, s2, isStr2 := k1.Keys()
	if isStr1 && isStr2 {
		return zend.ZendiSmartStrcmp(s1, s2)
	} else if !isStr1 && !isStr2 {
		return b.Cond(l1 < l2, -1, 1)
	} else {
		d1 := arrayKeyToDouble(k1)
		d2 := arrayKeyToDouble(k2)
		return b.Compare(d1, d2)
	}
}
func arrayKeyCompareStringCase(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return ascii.StrCaseCompare(str1, str2)
}
func arrayKeyCompareString(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return strings.Compare(str1, str2)
}

func arrayKeyCompareStringNaturalCase(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return str.Strnatcmp(str1, str2, true)
}

func arrayKeyCompareStringNatural(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return str.Strnatcmp(str1, str2, false)
}
func arrayKeyCompareStringLocale(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return strcoll(str1, str2)
}
func PhpGetKeyCompareFunc(sortType zend.ZendLong, reverse bool) types.ArrayComparer {
	var comparer types.ArrayComparer
	switch sortType & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		comparer = keyComparer(arrayKeyCompareNumeric)
	case PHP_SORT_STRING:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = keyComparer(arrayKeyCompareStringCase)
		} else {
			comparer = keyComparer(arrayKeyCompareString)
		}
	case PHP_SORT_NATURAL:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = keyComparer(arrayKeyCompareStringNaturalCase)
		} else {
			comparer = keyComparer(arrayKeyCompareStringNatural)
		}
	case PHP_SORT_LOCALE_STRING:
		comparer = keyComparer(arrayKeyCompareStringLocale)
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		comparer = keyComparer(arrayKeyCompare)
	}
	if reverse {
		comparer = reserveComparer(comparer)
	}
	return comparer
}

func arrayUserDataComparer(cmpFunction zpp.Callable) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		arg1 := p1.GetVal()
		arg2 := p2.GetVal()
		if retval, ok := cmpFunction.Call(arg1, arg2); ok && retval.IsNotUndef() {
			var ret = zend.ZvalGetLong(retval)
			return zend.ZEND_NORMALIZE_BOOL(ret)
		} else {
			return 0
		}
	}
}

func arrayUserKeyComparer(cmpFunction zpp.Callable) types.ArrayComparer {
	return func(p1, p2 types.ArrayPair) int {
		arg1 := p1.GetKey().ToZval()
		arg2 := p2.GetKey().ToZval()
		if retval, ok := cmpFunction.Call(arg1, arg2); ok && retval.IsNotUndef() {
			var ret = zend.ZvalGetLong(retval)
			return zend.ZEND_NORMALIZE_BOOL(ret)
		} else {
			return 0
		}
	}
}

func sortPairs(pairs []types.ArrayPair, cmp types.ArrayComparer) {
	sort.Slice(pairs, func(i, j int) bool {
		return cmp(pairs[i], pairs[j]) < 0
	})
}

func arrayDiffWrapper(args []*types.Zval, cmp types.ArrayComparer) (*types.Array, bool) {
	var arrays []*types.Array
	for i, arg := range args {
		if !arg.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
			return nil, false
		}
		arrays[i] = arg.Array()
	}
	return arrayDiff(arrays[0], arrays[1:], cmp), true
}

func arrayDiff(array *types.Array, arrays []*types.Array, cmp types.ArrayComparer) *types.Array {
	retArr := types.ZendArrayDup(array)
	if len(arrays) == 0 {
		return retArr
	}

	// 获取基准键值对，并排序
	basePairs := array.Pairs()
	sortPairs(basePairs, cmp)

	// 获取diff键值对，并排序
	diffPairCount := 0
	for _, array := range arrays {
		diffPairCount += array.Len()
	}
	diffPairs := make([]types.ArrayPair, diffPairCount)
	for _, array := range arrays {
		diffPairs = append(diffPairs, array.Pairs()...)
	}
	sortPairs(diffPairs, cmp)

	// diff
	for len(basePairs) > 0 && len(diffPairs) > 0 {
		c := cmp(basePairs[0], diffPairs[0])
		if c == 0 {
			retArr.Delete(basePairs[0].GetKey())
		}
		if c <= 0 {
			basePairs = basePairs[1:]
		}
		if c >= 0 {
			diffPairs = diffPairs[1:]
		}
	}

	return retArr
}

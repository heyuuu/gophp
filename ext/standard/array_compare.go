package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/cmpkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/fix"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"github.com/heyuuu/gophp/shim/cmp"
	"sort"
	"strconv"
	"strings"
)

func arrayValueComparerWithCtx(ctx *php.Context, comparer func(ctx *php.Context, v1, v2 types.Zval) int) types.ArrayValueComparerFunc {
	return types.ArrayValueComparerFunc(func(v1, v2 types.Zval) int {
		return comparer(ctx, v1, v2)
	})
}

func reserveComparer(comparer types.ArrayComparer) types.ArrayComparer {
	return types.ArrayComparerFunc(func(p1, p2 types.ArrayPair) int {
		return comparer.Compare(p2, p1)
	})
}

func twiceComparer(compare1 types.ArrayComparer, compare2 types.ArrayComparer) types.ArrayComparer {
	return types.ArrayComparerFunc(func(p1, p2 types.ArrayPair) int {
		c := compare1.Compare(p1, p2)
		if c != 0 {
			return c
		}
		return compare2.Compare(p1, p2)
	})
}

func arrayNaturalGeneralCompare(ctx *php.Context, v1, v2 types.Zval) int {
	str1 := php.ZvalGetStrVal(ctx, v1)
	str2 := php.ZvalGetStrVal(ctx, v2)
	return Strnatcmp(str1, str2, false)
}

func arrayNaturalGeneralCaseCompare(ctx *php.Context, v1, v2 types.Zval) int {
	str1 := php.ZvalGetStrVal(ctx, v1)
	str2 := php.ZvalGetStrVal(ctx, v2)
	return Strnatcmp(str1, str2, true)
}
func arrayDataCompare(ctx *php.Context, v1, v2 types.Zval) int {
	if ret, ok := php.ZvalCompareEx(ctx, v1, v2); ok {
		return ret
	} else {
		return 0
	}
}

func getArrayValueComparer(ctx *php.Context, sortType int, reverse bool) types.ArrayComparer {
	var comparerWithCtx func(ctx *php.Context, v1, v2 types.Zval) int
	switch sortType &^ PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		comparerWithCtx = php.NumericCompareFunction
	case PHP_SORT_STRING:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparerWithCtx = php.StringCaseCompareFunction
		} else {
			comparerWithCtx = php.StringCompareFunction
		}
	case PHP_SORT_NATURAL:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparerWithCtx = arrayNaturalGeneralCaseCompare
		} else {
			comparerWithCtx = arrayNaturalGeneralCompare
		}
	case PHP_SORT_LOCALE_STRING:
		comparerWithCtx = php.StringLocaleCompareFunction
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		comparerWithCtx = arrayDataCompare
	}
	comparer := arrayValueComparerWithCtx(ctx, comparerWithCtx)
	if reverse {
		return reserveComparer(comparer)
	}
	return comparer
}

func arrayKeyToDouble(k types.ArrayKey) float64 {
	if k.IsStrKey() {
		return php.ParseDouble(k.StrKey())
	} else {
		return float64(k.IdxKey())
	}
}
func arrayKeyToString(k types.ArrayKey) string {
	if k.IsStrKey() {
		return k.StrKey()
	} else {
		return strconv.Itoa(k.IdxKey())
	}
}
func arrayKeyCompareNumeric(k1 types.ArrayKey, k2 types.ArrayKey) int {
	if k1.IsIdxKey() && k2.IsIdxKey() {
		return cmp.Compare(k1.IdxKey(), k2.IdxKey())
	} else {
		d1 := arrayKeyToDouble(k1)
		d2 := arrayKeyToDouble(k2)
		return cmp.Compare(d1, d2)
	}
}
func arrayKeyCompare(k1 types.ArrayKey, k2 types.ArrayKey) int {
	l1, s1, isStr1 := k1.IdxKey(), k1.StrKey(), k1.IsStrKey()
	l2, s2, isStr2 := k2.IdxKey(), k2.StrKey(), k2.IsStrKey()
	if isStr1 && isStr2 {
		return php.SmartStrCompare(s1, s2)
	} else if !isStr1 && !isStr2 {
		return lang.Cond(l1 < l2, -1, 1)
	} else {
		d1 := arrayKeyToDouble(k1)
		d2 := arrayKeyToDouble(k2)
		return cmp.Compare(d1, d2)
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
	return Strnatcmp(str1, str2, true)
}

func arrayKeyCompareStringNatural(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return Strnatcmp(str1, str2, false)
}
func arrayKeyCompareStringLocale(k1 types.ArrayKey, k2 types.ArrayKey) int {
	str1 := arrayKeyToString(k1)
	str2 := arrayKeyToString(k2)
	return fix.StrColl(str1, str2)
}
func getKeyCompareFunc(sortType int, reverse bool) types.ArrayComparer {
	var comparer types.ArrayKeyComparerFunc
	switch sortType &^ PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		comparer = arrayKeyCompareNumeric
	case PHP_SORT_STRING:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = arrayKeyCompareStringCase
		} else {
			comparer = arrayKeyCompareString
		}
	case PHP_SORT_NATURAL:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparer = arrayKeyCompareStringNaturalCase
		} else {
			comparer = arrayKeyCompareStringNatural
		}
	case PHP_SORT_LOCALE_STRING:
		comparer = arrayKeyCompareStringLocale
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		comparer = arrayKeyCompare
	}
	if reverse {
		return reserveComparer(comparer)
	}
	return comparer
}

func arrayUserValueComparer(ctx *php.Context, cmpFunction zpp.Callable) types.ArrayValueComparerFunc {
	return func(arg1, arg2 types.Zval) int {
		if retval, ok := cmpFunction.Call(arg1, arg2); ok && retval.IsNotUndef() {
			var ret = php.ZvalGetLong(ctx, retval)
			return cmpkit.Normalize(ret)
		} else {
			return 0
		}
	}
}

func arrayUserKeyComparer(ctx *php.Context, cmpFunction zpp.Callable) types.ArrayKeyComparerFunc {
	return func(k1, k2 types.ArrayKey) int {
		if retval, ok := cmpFunction.Call(k1.ToZval(), k2.ToZval()); ok && retval.IsNotUndef() {
			var ret = php.ZvalGetLong(ctx, retval)
			return cmpkit.Normalize(ret)
		} else {
			return 0
		}
	}
}

func sortPairs(pairs []types.ArrayPair, cmp types.ArrayComparer) {
	sort.Slice(pairs, func(i, j int) bool {
		return cmp.Compare(pairs[i], pairs[j]) < 0
	})
}

func checkArrayArgs(ctx *php.Context, args []types.Zval, startIdx int) ([]*types.Array, bool) {
	var arrays = make([]*types.Array, len(args))
	for i, arg := range args {
		if !arg.IsArray() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Expected parameter %d to be an array, %s given", i+startIdx+1, types.ZendZvalTypeName(args[i])))
			return nil, false
		}
		arrays[i] = arg.Array()
	}
	return arrays, true
}

func arrayDiffWrapper(ctx *php.Context, args []types.Zval, cmp types.ArrayComparer) (*types.Array, bool) {
	arrays, ok := checkArrayArgs(ctx, args, 0)
	if !ok {
		return nil, false
	}
	return arrayDiff(arrays[0], arrays[1:], cmp), true
}

func arrayDiff(array *types.Array, arrays []*types.Array, cmp types.ArrayComparer) *types.Array {
	retArr := array.Dup()
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
		c := cmp.Compare(basePairs[0], diffPairs[0])
		if c == 0 {
			retArr.Delete(basePairs[0].Key)
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

func arrayDiffKeyWrapper(ctx *php.Context, args []types.Zval, cmp types.ArrayValueComparerFunc) (*types.Array, bool) {
	arrays, ok := checkArrayArgs(ctx, args, 0)
	if !ok {
		return nil, false
	}
	return arrayDiffKey(arrays[0], arrays[1:], cmp), true
}

func arrayDiffKey(array *types.Array, arrays []*types.Array, dataComparer types.ArrayValueComparerFunc) *types.Array {
	retArr := types.NewArray()
	array.Each(func(key types.ArrayKey, val types.Zval) {
		keep := true
		for _, cmpArray := range arrays {
			if data := cmpArray.Find(key); data.IsUndef() && (dataComparer == nil || dataComparer(val, data) == 0) {
				keep = false
				break
			}
		}
		if keep {
			retArr.Update(key, val)
		}
	})

	return retArr
}

func arrayIntersectWrapper(ctx *php.Context, args []types.Zval, cmp types.ArrayComparer) (*types.Array, bool) {
	arrays, ok := checkArrayArgs(ctx, args, 0)
	if !ok {
		return nil, false
	}
	return arrayIntersect(arrays[0], arrays[1:], cmp), true
}

func arrayIntersect(array *types.Array, arrays []*types.Array, cmp types.ArrayComparer) *types.Array {
	if len(arrays) == 0 {
		return types.NewArray()
	}

	// 获取基准键值对，并排序
	basePairs := array.Pairs()
	sortPairs(basePairs, cmp)

	// 获取对比数组键值对，并各自排序
	cmpPairsList := make([][]types.ArrayPair, len(arrays))
	for i, cmpArray := range arrays {
		cmpPairsList[i] = cmpArray.Pairs()
		sortPairs(cmpPairsList[i], cmp)
	}

	retArr := types.NewArrayCap(array.Len())
	for _, pair := range basePairs {
		keep := true // 标识会否保留此元素
		for cmpIdx, cmpPairs := range cmpPairsList {
			// 去除小于 pair 的元素，找到第一个 >= pair 的元素或到队尾
			for len(cmpPairs) > 0 {
				c := cmp.Compare(pair, cmpPairs[0])
				if c > 0 {
					cmpPairsList[cmpIdx] = cmpPairs[1:]
				} else {
					if c < 0 {
						keep = false
					}
					break
				}
			}
			// 已到队尾说明没有上匹配任何值
			if len(cmpPairs) == 0 {
				keep = false
			}
			if !keep {
				break
			}
		}
		if keep {
			retArr.Update(pair.Key, pair.Val)
		}
	}

	return retArr
}

func arrayIntersectKeyWrapper(ctx *php.Context, args []types.Zval, cmp types.ArrayValueComparerFunc) (*types.Array, bool) {
	arrays, ok := checkArrayArgs(ctx, args, 0)
	if !ok {
		return nil, false
	}
	return arrayIntersectKey(arrays[0], arrays[1:], cmp), true
}

func arrayIntersectKey(array *types.Array, arrays []*types.Array, dataComparer types.ArrayValueComparerFunc) *types.Array {
	retArr := types.NewArrayCap(array.Len())
	array.Each(func(key types.ArrayKey, val types.Zval) {
		keep := true
		for _, cmpArray := range arrays {
			if data := cmpArray.Find(key); data.IsUndef() || dataComparer != nil && dataComparer(val, data) != 0 {
				keep = false
				break
			}
		}
		if keep {
			retArr.Update(key, val)
		}
	})
	return retArr
}

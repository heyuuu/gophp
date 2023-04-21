package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
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

func PhpGetDataCompareFuncEx(sortType zend.ZendLong, reverse bool) types.ArrayComparer {
	var comparser types.ArrayComparer
	switch sortType & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		comparser = valueComparer(zend.NumericCompareFunction)
	case PHP_SORT_STRING:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparser = valueComparer(zend.StringCaseCompareFunction)
		} else {
			comparser = valueComparer(zend.StringCompareFunction)
		}
	case PHP_SORT_NATURAL:
		if (sortType & PHP_SORT_FLAG_CASE) != 0 {
			comparser = valueComparer(arrayNaturalGeneralCaseCompare)
		} else {
			comparser = valueComparer(arrayNaturalGeneralCompare)
		}
	case PHP_SORT_LOCALE_STRING:
		comparser = valueComparer(zend.StringLocaleCompareFunction)
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		comparser = valueComparer(arrayDataCompare)
	}
	if reverse {
		comparser = reserveComparer(comparser)
	}
	return comparser
}

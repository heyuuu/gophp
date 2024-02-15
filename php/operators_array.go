package php

import (
	"github.com/heyuuu/gophp/kits/cmpkit"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

type ZvalComparer func(types.Zval, types.Zval) int

func iArrayKeyCompare(k1, k2 types.ArrayKey) int {
	i1, s1, isStr1 := k1.IdxKey(), k1.StrKey(), k1.IsStrKey()
	i2, s2, isStr2 := k2.IdxKey(), k2.StrKey(), k2.IsStrKey()
	if !isStr1 && !isStr2 {
		return i1 - i2
	} else if isStr1 && isStr2 {
		return strings.Compare(s1, s2)
	} else {
		/* Mixed key types: A string key is considered as larger */
		if isStr1 {
			return 1
		} else {
			return -1
		}
	}
}

func iArrayCompareOrdered(ht1 *types.Array, ht2 *types.Array, comparer ZvalComparer) int {
	// check len
	if ht1.Len() != ht2.Len() {
		return ht1.Len() - ht2.Len()
	}

	pairs1, pairs2 := ht1.Pairs(), ht2.Pairs()
	for idx, pair1 := range pairs1 {
		pair2 := pairs2[idx]

		// compare key
		compareKeyResult := iArrayKeyCompare(pair1.Key, pair2.Key)
		if compareKeyResult != 0 {
			return compareKeyResult
		}

		// compare value
		v1 := pair1.Val.DeRef()
		v2 := pair2.Val.DeRef()
		if v1.IsUndef() {
			if !v2.IsUndef() {
				return -1
			}
		} else if v2.IsUndef() {
			return 1
		} else {
			result := comparer(v1, v2)
			if result != 0 {
				return result
			}
		}
	}
	return 0
}

func iArrayCompareUnordered(ht1 *types.Array, ht2 *types.Array, comparer ZvalComparer) int {
	// check len
	if ht1.Len() != ht2.Len() {
		return ht1.Len() - ht2.Len()
	}

	pairs1 := ht1.Pairs()
	for _, p1 := range pairs1 {
		// find value in ht2
		v2 := ht2.Find(p1.Key)
		if v2.IsUndef() {
			return 1
		}

		// compare value
		v1 := p1.Val.DeRef()
		v2 = v2.DeRef()
		if v1.IsUndef() {
			if !v2.IsUndef() {
				return -1
			}
		} else if v2.IsUndef() {
			return 1
		} else {
			result := comparer(v1, v2)
			if result != 0 {
				return result
			}
		}
	}
	return 0
}

func iArrayCompare(ctx *Context, ht1 *types.Array, ht2 *types.Array, comparer ZvalComparer, ordered bool) int {
	var result int
	if ht1 == ht2 {
		return 0
	}

	/* It's enough to protect only one of the arrays.
	 * The second one may be referenced from the first and this may cause
	 * false recursion detection.
	 */
	if ht1.IsRecursive() {
		ErrorNoreturn(ctx, perr.E_ERROR, "Nesting level too deep - recursive dependency?")
	}

	ht1.ProtectRecursive()

	if ordered {
		result = iArrayCompareOrdered(ht1, ht2, comparer)
	} else {
		result = iArrayCompareUnordered(ht1, ht2, comparer)
	}

	ht1.UnprotectRecursive()

	return cmpkit.Normalize(result)
}

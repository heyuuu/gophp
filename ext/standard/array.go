package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"sort"
)

/**
 * Constants & Global Variables
 */
const EXTR_OVERWRITE = 0
const EXTR_SKIP = 1
const EXTR_PREFIX_SAME = 2
const EXTR_PREFIX_ALL = 3
const EXTR_PREFIX_INVALID = 4
const EXTR_PREFIX_IF_EXISTS = 5
const EXTR_IF_EXISTS = 6
const EXTR_REFS = 0x100
const CASE_LOWER = 0
const CASE_UPPER = 1

const MULTISORT_ORDER = 0
const MULTISORT_TYPE = 1
const MULTISORT_LAST = 2

const PHP_SORT_REGULAR = 0
const PHP_SORT_NUMERIC = 1
const PHP_SORT_STRING = 2
const PHP_SORT_DESC = 3
const PHP_SORT_ASC = 4
const PHP_SORT_LOCALE_STRING = 5
const PHP_SORT_NATURAL = 6
const PHP_SORT_FLAG_CASE = 8
const COUNT_NORMAL = 0
const COUNT_RECURSIVE = 1
const ARRAY_FILTER_USE_BOTH = 1
const ARRAY_FILTER_USE_KEY = 2

/**
 * functions
 */
func RegisterArrayConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterConstant(ctx, moduleNumber, "EXTR_OVERWRITE", php.Long(EXTR_OVERWRITE))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_SKIP", php.Long(EXTR_SKIP))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_PREFIX_SAME", php.Long(EXTR_PREFIX_SAME))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_PREFIX_ALL", php.Long(EXTR_PREFIX_ALL))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_PREFIX_INVALID", php.Long(EXTR_PREFIX_INVALID))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_PREFIX_IF_EXISTS", php.Long(EXTR_PREFIX_IF_EXISTS))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_IF_EXISTS", php.Long(EXTR_IF_EXISTS))
	php.RegisterConstant(ctx, moduleNumber, "EXTR_REFS", php.Long(EXTR_REFS))
	php.RegisterConstant(ctx, moduleNumber, "SORT_ASC", php.Long(PHP_SORT_ASC))
	php.RegisterConstant(ctx, moduleNumber, "SORT_DESC", php.Long(PHP_SORT_DESC))
	php.RegisterConstant(ctx, moduleNumber, "SORT_REGULAR", php.Long(PHP_SORT_REGULAR))
	php.RegisterConstant(ctx, moduleNumber, "SORT_NUMERIC", php.Long(PHP_SORT_NUMERIC))
	php.RegisterConstant(ctx, moduleNumber, "SORT_STRING", php.Long(PHP_SORT_STRING))
	php.RegisterConstant(ctx, moduleNumber, "SORT_LOCALE_STRING", php.Long(PHP_SORT_LOCALE_STRING))
	php.RegisterConstant(ctx, moduleNumber, "SORT_NATURAL", php.Long(PHP_SORT_NATURAL))
	php.RegisterConstant(ctx, moduleNumber, "SORT_FLAG_CASE", php.Long(PHP_SORT_FLAG_CASE))
	php.RegisterConstant(ctx, moduleNumber, "CASE_LOWER", php.Long(CASE_LOWER))
	php.RegisterConstant(ctx, moduleNumber, "CASE_UPPER", php.Long(CASE_UPPER))
	php.RegisterConstant(ctx, moduleNumber, "COUNT_NORMAL", php.Long(COUNT_NORMAL))
	php.RegisterConstant(ctx, moduleNumber, "COUNT_RECURSIVE", php.Long(COUNT_RECURSIVE))
	php.RegisterConstant(ctx, moduleNumber, "ARRAY_FILTER_USE_BOTH", php.Long(ARRAY_FILTER_USE_BOTH))
	php.RegisterConstant(ctx, moduleNumber, "ARRAY_FILTER_USE_KEY", php.Long(ARRAY_FILTER_USE_KEY))
}

// @zif(onError=1)
func ZifKrsort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getKeyCompareFunc(sortFlags, true)
	arg.Sort(cmp, false)
	return true
}

// @zif(onError=1)
func ZifKsort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getKeyCompareFunc(sortFlags, false)
	arg.Sort(cmp, false)
	return true
}

func PhpCountRecursive(ctx *php.Context, ht *types.Array) int {
	if ht.IsRecursive() {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "recursion detected")
		return 0
	}
	ht.ProtectRecursive()

	cnt := ht.Count()
	ht.Each(func(key types.ArrayKey, value types.Zval) {
		value = value.DeRef()
		if value.IsArray() {
			cnt += PhpCountRecursive(ctx, value.Array())
		}
	})

	ht.UnprotectRecursive()
	return cnt
}

// @zif(alias="sizeof")
func ZifCount(ctx *php.Context, var_ *types.Zval, _ zpp.Opt, mode int) int {
	var array = var_
	var cnt int
	switch array.Type() {
	case types.IsNull:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 0
	case types.IsArray:
		if mode != COUNT_RECURSIVE {
			cnt = array.Array().Count()
		} else {
			cnt = PhpCountRecursive(ctx, array.Array())
		}
		return cnt
	case types.IsObject:
		//var retval types.Zval
		//var long int

		/* first, we check if the handler is defined */
		//if array.Object().CanCountElements() {
		//	long = 1
		//	if c, ok := array.Object().CountElements(); ok {
		//		return c
		//	}
		//	php.AssertNoException(ctx)
		//}

		/* if not and the object implements Countable we call its count() method */
		//if php.InstanceofFunction(types.Z_OBJCE_P(array), php.ZendCeCountable) {
		//	php.ZendCallMethodWith0Params(ctx, array, nil, nil, "count", &retval)
		//	if retval.IsNotUndef() {
		//		long = php.ZvalGetLong(ctx, retval)
		//	}
		//	return long
		//}

		/* If There's no handler and it doesn't implement Countable then add a warning */
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 1
	default:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 1
	}
}
func ZifNatsort(ctx *php.Context, arg zpp.RefArray) bool {
	cmp := getArrayValueComparer(ctx, PHP_SORT_NATURAL, false)
	arg.Sort(cmp, false)
	return true
}
func ZifNatcasesort(ctx *php.Context, arg zpp.RefArray) bool {
	cmp := getArrayValueComparer(ctx, PHP_SORT_NATURAL|PHP_SORT_FLAG_CASE, false)
	arg.Sort(cmp, false)
	return true
}

// @zif(onError=1)
func ZifAsort(ctx *php.Context, arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getArrayValueComparer(ctx, sortFlags, false)
	arg.Sort(cmp, false)
	return true
}

// @zif(onError=1)
func ZifArsort(ctx *php.Context, arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getArrayValueComparer(ctx, sortFlags, true)
	arg.Sort(cmp, false)
	return true
}

// @zif(onError=1)
func ZifSort(ctx *php.Context, arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getArrayValueComparer(ctx, sortFlags, false)
	arg.Sort(cmp, true)
	return true
}

// @zif(onError=1)
func ZifRsort(ctx *php.Context, arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := getArrayValueComparer(ctx, sortFlags, true)
	arg.Sort(cmp, true)
	return true
}

func phpUsortEx(arr *types.Array, compareFunc types.ArrayComparer, renumber bool) bool {
	if arr.Len() > 0 {
		/* Copy array, so the in-place modifications will not be visible to the callback function */
		dupArr := arr.Dup()
		dupArr.Sort(compareFunc, renumber)
		arr.SetDataByArray(dupArr)
	}
	return true
}
func ZifUsort(ctx *php.Context, arg zpp.RefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserValueComparer(ctx, cmpFunction)
	phpUsortEx(arg, cmp, true)
	return true
}
func ZifUasort(ctx *php.Context, arg zpp.RefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserValueComparer(ctx, cmpFunction)
	phpUsortEx(arg, cmp, false)
	return true
}
func ZifUksort(ctx *php.Context, arg zpp.RefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserKeyComparer(ctx, cmpFunction)
	phpUsortEx(arg, cmp, false)
	return true
}
func ZifEnd(array zpp.RefArray) types.Zval {
	array.MoveEnd()
	pair := array.Current()
	if !pair.IsValid() {
		return types.False
	}
	return pair.Val.DeRef()
}

func ZifPrev(array zpp.RefArray) types.Zval {
	array.MovePrev()
	pair := array.Current()
	if !pair.IsValid() {
		return types.False
	}
	return pair.Val.DeRef()
}
func ZifNext(array zpp.RefArray) types.Zval {
	array.MoveNext()
	pair := array.Current()
	if !pair.IsValid() {
		return types.False
	}
	return pair.Val.DeRef()
}

func ZifReset(array zpp.RefArray) types.Zval {
	pair := array.Reset()
	if !pair.IsValid() {
		return types.False
	}
	return pair.Val.DeRef()
}

// @zif(alias="pos")
func ZifCurrent(array zpp.ArrayOrObjectHt) (types.Zval, bool) {
	pair := array.Current()
	if !pair.IsValid() {
		return types.Undef, false
	}

	return pair.Val.DeRef(), true
}
func ZifKey(array zpp.ArrayOrObjectHt) types.Zval {
	pair := array.Current()
	if !pair.IsValid() {
		return types.Null
	} else {
		return pair.Key.ToZval()
	}
}

func ZifMin(ctx *php.Context, arg types.Zval, args []types.Zval) types.Zval {
	/* mixed min ( array $values ) */
	if len(args) == 0 {
		if !arg.IsArray() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "When only one parameter is given, it must be an array")
			return types.Null
		} else {
			var result types.Zval
			arg.Array().Each(func(_ types.ArrayKey, value types.Zval) {
				if result.IsUndef() || arrayDataCompare(ctx, result, value) > 0 {
					result = value
				}
			})
			if result.IsNotUndef() {
				return result
			} else {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, "Array must contain at least one element")
				return types.False
			}
		}
	} else {
		result := arg
		for _, value := range args {
			if arrayDataCompare(ctx, result, value) > 0 {
				result = value
			}
		}
		return result
	}
}
func ZifMax(ctx *php.Context, arg types.Zval, args []types.Zval) types.Zval {
	/* mixed max ( array $values ) */
	if len(args) == 0 {
		if !arg.IsArray() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "When only one parameter is given, it must be an array")
			return types.Null
		} else {
			var result types.Zval
			arg.Array().Each(func(_ types.ArrayKey, value types.Zval) {
				if result.IsUndef() || arrayDataCompare(ctx, result, value) < 0 {
					result = value
				}
			})
			if !result.IsUndef() {
				return result
			} else {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, "Array must contain at least one element")
				return types.False
			}
		}
	} else {
		result := arg
		for _, value := range args {
			if arrayDataCompare(ctx, result, value) < 0 {
				result = value
			}
		}
		return result
	}
}

func arrayUserWalkHandler(handler zpp.Callable, userdata types.Zval) func(value types.Zval, key types.ArrayKey) bool {
	return func(value types.Zval, key types.ArrayKey) bool {
		var ok bool
		if userdata.IsNotUndef() {
			_, ok = handler.Call(value, key.ToZval(), userdata)
		} else {
			_, ok = handler.Call(value, key.ToZval())
		}
		return ok
	}
}

func arrayWalk(ctx *php.Context, array types.Zval, recursive bool, handler func(value types.Zval, key types.ArrayKey) bool) bool {
	var targetHash = php.HashOf(array)
	var result = true

	/* Iterate through hash */
	_ = targetHash.EachEx(func(key types.ArrayKey, zv types.Zval) error {
		/* Skip undefined indirect elements */
		//if zv.IsIndirect() {
		//	zv = *zv.Indirect()
		//	if zv.IsUndef() {
		//		return nil
		//	}
		//
		//	/* Add type source for property references. */
		//	if !zv.IsRef() && array.IsType(types.IsObject) {
		//		var prop_info = php.ZendGetTypedPropertyInfoForSlot(array.Object(), &zv)
		//		if prop_info != nil {
		//			zv.SetNewRef(&zv)
		//			php.ZEND_REF_ADD_TYPE_SOURCE(zv.Ref(), prop_info)
		//		}
		//	}
		//}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */
		types.ZVAL_MAKE_REF(&zv)

		/* Back up hash position, as it may change */
		if recursive && zv.RefVal().IsArray() {
			var thash *types.Array
			var ref types.Zval
			types.ZVAL_COPY_VALUE(&ref, &zv)
			zv = zv.DeRef()
			types.SeparateArray(&zv)
			thash = zv.Array()
			if thash.IsRecursive() {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, "recursion detected")
				return lang.BreakErr
			}

			/* backup the fcall info and cache */
			thash.ProtectRecursive()
			result = arrayWalk(ctx, zv, recursive, handler)
			if ref.RefVal().IsArray() && thash == ref.RefVal().Array() {
				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */
				thash.UnprotectRecursive()
			}
		} else {
			/* Call the userland function */
			result = handler(zv, key)
		}
		if !result {
			return lang.BreakErr
		}

		/* Reload array and position -- both may have changed */
		if ctx.EG().HasException() {
			return lang.BreakErr
		}

		return nil
	})

	return result
}

func ZifArrayWalk(ctx *php.Context, array zpp.RefArrayOrObject, callable zpp.Callable, _ zpp.Opt, arg *types.Zval) bool {
	handler := arrayUserWalkHandler(callable, *arg)
	arrayWalk(ctx, array.Val(), false, handler)
	return true
}
func ZifArrayWalkRecursive(ctx *php.Context, array zpp.RefArrayOrObject, callable zpp.Callable, _ zpp.Opt, arg *types.Zval) bool {
	handler := arrayUserWalkHandler(callable, *arg)
	arrayWalk(ctx, array.Val(), true, handler)
	return true
}

func searchArray(ctx *php.Context, value types.Zval, array *types.Array, strict bool) *types.ArrayKey {
	var targetKey *types.ArrayKey = nil
	if strict {
		array.EachEx(func(key types.ArrayKey, entry types.Zval) error {
			entry = entry.DeRef()
			if php.ZvalIsIdentical(ctx, value, entry) {
				targetKey = &key
				return lang.BreakErr
			}
			return nil
		})
	} else {
		array.EachEx(func(key types.ArrayKey, entry types.Zval) error {
			entry = entry.DeRef()
			if php.FastEqualFunction(ctx, value, entry) {
				targetKey = &key
				return lang.BreakErr
			}
			return nil
		})
	}
	return targetKey
}
func ZifInArray(ctx *php.Context, needle types.Zval, haystack *types.Array, _ zpp.Opt, strict bool) bool {
	key := searchArray(ctx, needle, haystack, strict)
	return key != nil
}
func ZifArraySearch(ctx *php.Context, needle types.Zval, haystack *types.Array, _ zpp.Opt, strict bool) types.Zval {
	key := searchArray(ctx, needle, haystack, strict)
	if key == nil {
		return types.False
	} else {
		return key.ToZval()
	}
}

func ZifArrayFill(ctx *php.Context, startKey int, num int, val types.Zval) (*types.Array, bool) {
	if num < 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Number of elements can't be negative")
		return nil, false
	}
	if num == 0 {
		return types.NewArray(), true
	}
	if num > math.MaxInt32 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Too many elements")
		return nil, false
	} else if startKey > math.MaxInt-num+1 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Cannot add element to the array as the next element is already occupied")
		return nil, false
	}

	// todo 尽量创建 packed array

	/* create hash */
	arr := types.NewArrayCap(num)
	for i := 0; i < num; i++ {
		arr.IndexAdd(startKey+i, val)
	}
	return arr, true
}
func ZifArrayFillKeys(ctx *php.Context, keys *types.Array, val types.Zval) *types.Array {
	arr := types.NewArrayCap(keys.Len())
	keys.Each(func(_ types.ArrayKey, entry types.Zval) {
		entry = entry.DeRef()
		if entry.IsLong() {
			arr.IndexUpdate(entry.Long(), val)
		} else {
			key := php.ZvalGetStrVal(ctx, entry)
			arr.SymtableUpdate(key, val)
		}
	})
	return arr
}
func rangeDouble(ctx *php.Context, zLow types.Zval, zHigh types.Zval, step float64) ([]types.Zval, bool) {
	php.Assert(step > 0)
	low := php.ZvalGetDouble(ctx, zLow)
	high := php.ZvalGetDouble(ctx, zHigh)
	if mathkit.IsInf(high) || mathkit.IsInf(low) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid range supplied: start=%0.0f end=%0.0f", low, high))
		return nil, false
	}
	if low > high {
		low, high = high, low
	}
	if high-low < step {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "step exceeds the specified range")
		return nil, false
	}

	size := (high-low)/step + 1
	if size >= float64(types.MaxArraySize) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high))
		return nil, false
	}

	arr := make([]types.Zval, int(size))
	for i := range arr {
		arr[i] = types.ZvalDouble(low + float64(i)*step)
	}
	return arr, true
}

func rangeLong(ctx *php.Context, zLow types.Zval, zHigh types.Zval, step int) ([]types.Zval, bool) {
	php.Assert(step > 0)
	low := php.ZvalGetLong(ctx, zLow)
	high := php.ZvalGetLong(ctx, zHigh)
	if high == low {
		return []types.Zval{types.ZvalLong(high)}, true
	}

	if low > high {
		low, high = high, low
	}
	if high-low < step {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "step exceeds the specified range")
		return nil, false
	}

	size := (high-low)/step + 1
	if size >= types.MaxArraySize {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("The supplied range exceeds the maximum array size: start=%d end=%d", low, high))
		return nil, false
	}

	arr := make([]types.Zval, size)
	for i := range arr {
		arr[i] = types.ZvalLong(low + i*step)
	}
	return arr, true
}

func rangeChar(low byte, high byte, step int) ([]types.Zval, bool) {
	if low == high {
		return []types.Zval{types.ZvalString(string(low))}, true
	}
	if low > high {
		low, high = high, low
	}

	bStep := byte(step)
	size := (high-low)/bStep + 1
	arr := make([]types.Zval, 0, size)
	for i := range arr {
		c := low + byte(i)*bStep
		arr[i] = types.ZvalString(string(c))
	}
	return arr, true
}

// @zif(onError=1)
func ZifRange(ctx *php.Context, low types.Zval, high types.Zval, _ zpp.Opt, step types.Zval) ([]types.Zval, bool) {
	var isStepDouble = false
	var stepVal = 1.0
	if step.IsNotUndef() {
		if step.IsDouble() {
			isStepDouble = true
		} else if step.IsString() {
			r := php.ParseNumber(step.String())
			if r.IsUndef() {
				/* bad number */
				php.ErrorDocRef(ctx, "", perr.E_WARNING, "Invalid range string - must be numeric")
				return nil, false
			}
			if r.IsLong() {
				// pass
			} else if r.IsDouble() {
				isStepDouble = true
			}
		}
		stepVal = php.ZvalGetDouble(ctx, step)

		/* We only want positive step values. */
		if stepVal < 0.0 {
			stepVal *= -1
		}
	}

	/* If the range is given as strings, generate an array of characters. */
	if low.IsString() && high.IsString() && len(low.String()) >= 1 && len(high.String()) >= 1 {
		lowStr := low.String()
		highStr := high.String()

		r1 := php.ParseNumber(lowStr)
		r2 := php.ParseNumber(highStr)
		if r1.IsNotUndef() && r2.IsNotUndef() {
			if r1.IsDouble() || r2.IsDouble() || isStepDouble {
				return rangeDouble(ctx, low, high, stepVal)
			} else if r1.IsLong() || r2.IsLong() {
				return rangeLong(ctx, low, high, int(stepVal))
			}
		}
		return rangeChar(lowStr[0], highStr[0], int(stepVal))
	} else if low.IsDouble() || high.IsDouble() || isStepDouble {
		return rangeDouble(ctx, low, high, stepVal)
	} else {
		return rangeLong(ctx, low, high, int(stepVal))
	}
}
func arrayDataShuffle(ctx *php.Context, array *types.Array) *types.Array {
	values := array.Values()
	for i := len(values) - 1; i >= 0; i-- {
		j := PhpMtRandRange(ctx, 0, i)
		if i != j {
			values[i], values[j] = values[j], values[i]
		}
	}
	return types.NewArrayOfZval(values)
}

// @zif(onError=1)
func ZifShuffle(ctx *php.Context, arg zpp.RefArray) bool {
	if arg.Len() > 1 {
		arg.SetDataByArray(arrayDataShuffle(ctx, arg))
	}
	return true
}

func phpSplice(arr *types.Array, offset int, length int, replace *types.Array) (*types.Array, *types.Array) {
	arrLen := arr.Len()
	// check offset [0,arrLen-1]
	if offset > arrLen {
		offset = arrLen
	} else if offset < 0 {
		offset += arrLen
		if offset < 0 {
			offset = 0
		}
	}
	// check length [0, arrLen - offset]
	if length < 0 {
		length = arrLen - offset + length
	} else if offset+length > arrLen {
		length = arrLen - offset
	}
	if length < 0 {
		length = 0
	}

	outHash := types.NewArray()
	removed := types.NewArray()
	pairs := arr.Pairs()

	// handle range [0, offset)
	for _, pair := range pairs[:offset] {
		key := pair.Key
		val := pair.Val
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	// handle range [offset, offset+length)
	for _, pair := range pairs[offset : offset+length] {
		key := pair.Key
		val := pair.Val
		if !key.IsStrKey() {
			removed.Append(val)
		} else {
			removed.KeyAdd(key.StrKey(), val)
		}
	}

	// handle insert
	if replace != nil {
		replace.Each(func(key types.ArrayKey, value types.Zval) {
			outHash.Append(value)
		})
	}

	// handle range [offset+length, len(array))
	for _, pair := range pairs[offset+length:] {
		key := pair.Key
		val := pair.Val
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	return outHash, removed
}
func ZifArrayPush(ctx *php.Context, stack zpp.RefArray, _ zpp.Opt, args []types.Zval) (int, bool) {
	for _, arg := range args {
		if stack.Append(arg) < 0 {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return 0, false
		}
	}

	return stack.Len(), true
}
func ZifArrayPop(stack zpp.RefArray) types.Zval {
	pair := stack.Pop()
	if !pair.IsValid() {
		return types.Null
	}

	return pair.Val.DeRef()
}
func ZifArrayShift(ctx *php.Context, stack zpp.RefArray) types.Zval {
	if stack.Len() == 0 {
		return types.Null
	}

	/* Get the first value and copy it into the return value */
	pair := stack.First()
	if !pair.IsValid() {
		return types.Null
	}

	key := pair.Key
	val := pair.Val.DeRef()

	/* Delete the first value */
	//if key.IsStrKey() && php.EG__(ctx).SymbolTable().IsHt(stack) {
	//	php.ZendDeleteGlobalVariable(ctx, pair.StrKey())
	//} else {
	stack.Delete(key)
	//}

	/* re-index like it did before */
	var k = 0
	stack.MapWithKey(func(key types.ArrayKey, value types.Zval) (types.ArrayKey, types.Zval) {
		if !key.IsStrKey() {
			key = types.IdxKey(k)
			k++
		}
		return key, value
	})

	return val
}
func ZifArrayUnshift(stack zpp.RefArray, values []types.Zval) int {
	newArr := types.NewArrayCap(stack.Len() + len(values))
	for _, value := range values {
		newArr.Append(value)
	}
	stack.Each(func(key types.ArrayKey, value types.Zval) {
		if key.IsStrKey() {
			newArr.KeyAdd(key.StrKey(), value)
		} else {
			newArr.Append(value)
		}
	})

	stack.SetDataByArray(newArr)
	return stack.Len()
}
func ZifArraySplice(ctx *php.Context, array zpp.RefArray, offset int, _ zpp.Opt, length_ *int, replacement types.Zval) *types.Array {
	var numIn = array.Len()
	var length = lang.Option(length_, numIn)
	var replaceArr *types.Array = nil

	if replacement.IsNotUndef() {
		/* Make sure the last argument, if passed, is an array */
		replaceArr = php.ZvalGetArray(ctx, replacement)
	}

	newArr, removedArr := phpSplice(array, offset, length, replaceArr)
	array.SetDataByArray(newArr)
	return removedArr
}
func ZifArraySlice(ctx *php.Context, array *types.Array, offset int, _ zpp.Opt, length_ types.Zval, preserveKeys bool) *types.Array {
	numIn := array.Len()

	/* We want all entries from offset to the end if length is not passed or is null */
	var length = 0
	if length_.IsUndef() || length_.IsNull() {
		length = numIn
	} else {
		length = php.ZvalGetLong(ctx, length_)
	}

	/* Clamp the offset.. */
	if offset > numIn {
		return types.NewArray()
	} else if offset < 0 && lang.Assign(&offset, numIn+offset) < 0 {
		offset += numIn
		if offset < 0 {
			offset = 0
		}
	}

	/* ..and the length */
	if length < 0 {
		length = numIn - offset + length
	} else if offset+length > numIn {
		length = numIn - offset
	}
	if length <= 0 {
		return types.NewArray()
	}

	/* Initialize returned array */
	retArr := types.NewArrayCap(length)

	/* Start at the beginning and go until we hit offset */
	count := 0
	_ = array.EachEx(func(key types.ArrayKey, value types.Zval) error {
		count++
		if count <= offset {
			return nil
		}
		if count > offset+length {
			return lang.BreakErr
		}
		if key.IsStrKey() || preserveKeys {
			retArr.Add(key, value)
		} else {
			retArr.Append(value)
		}
		return nil
	})
	return retArr
}
func PhpArrayMergeRecursive(ctx *php.Context, dest *types.Array, src *types.Array) bool {
	var ok bool = true
	src.EachEx(func(key types.ArrayKey, value types.Zval) error {
		if !key.IsStrKey() {
			dest.Append(value)
			return nil
		}

		strKey := key.StrKey()
		destEntry := dest.KeyFind(strKey)
		if destEntry.IsUndef() {
			dest.KeyAdd(strKey, value)
		}

		var srcZval = value.DeRef()
		var destZval = destEntry.DeRef()
		var thash *types.Array
		var tmp types.Zval
		if destZval.IsArray() {
			thash = destZval.Array()
		} else {
			thash = nil
		}
		if thash != nil && thash.IsRecursive() || value == destEntry && destEntry.IsRef() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "recursion detected")
			ok = false
			return lang.BreakErr
		}
		//b.Assert(!(destEntry.IsReference()) || destEntry.GetRefcount() > 1)
		types.SeparateZval(&destEntry)
		destZval = destEntry
		if destZval.IsNull() {
			destZval.SetEmptyArray()
			destZval.Array().AddNextIndexNull()
		} else {
			destZval.SetArray(php.ZvalGetArray(ctx, destZval))
		}
		tmp.SetUndef()
		if srcZval.IsObject() {
			tmp = srcZval
			tmp.SetArray(php.ZvalGetArray(ctx, tmp))
			srcZval = tmp
		}
		if srcZval.IsArray() {
			if thash != nil {
				thash.ProtectRecursive()
			}
			ret := PhpArrayMergeRecursive(ctx, destZval.Array(), srcZval.Array())
			if thash != nil {
				thash.UnprotectRecursive()
			}
			if !ret {
				ok = false
				return lang.BreakErr
			}
		} else {
			destZval.Array().Append(srcZval)
		}
		return nil
	})
	return ok
}
func PhpArrayMerge(dest *types.Array, src *types.Array) {
	src.Each(func(key types.ArrayKey, value types.Zval) {
		if key.IsStrKey() {
			dest.KeyUpdate(key.StrKey(), value)
		} else {
			dest.Append(value)
		}
	})
}

func PhpArrayReplaceRecursive(ctx *php.Context, dest *types.Array, src *types.Array) error {
	return src.EachEx(func(key types.ArrayKey, srcEntry types.Zval) error {
		srcZval := srcEntry.DeRef()
		if !srcZval.IsArray() {
			dest.Update(key, srcEntry)
			return nil
		}

		destEntry := dest.Find(key)
		if destEntry.IsNotUndef() || !destEntry.IsArray() && (!(destEntry.IsRef()) || !destEntry.RefVal().IsArray()) {
			dest.Update(key, srcEntry)
			return nil
		}
		destZval := destEntry.DeRef()

		// src/dest 对应值均为 array 的情况下，递归替换
		if destZval.Array().IsRecursive() || srcZval.Array().IsRecursive() || srcEntry.IsRef() && destEntry.IsRef() && srcEntry.Ref() == destEntry.Ref() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "recursion detected")
			return lang.BreakErr
		}
		types.SeparateZval(&destEntry)
		destZval = destEntry

		php.Assert(srcZval.IsArray() && destZval.IsArray())

		srcZval.Array().ProtectRecursive()
		destZval.Array().ProtectRecursive()
		defer srcZval.Array().UnprotectRecursive()
		defer destZval.Array().UnprotectRecursive()

		return PhpArrayReplaceRecursive(ctx, destZval.Array(), srcZval.Array())
	})
}
func arrayMergeWrapper(ctx *php.Context, args []types.Zval, recursive bool) *types.Array {
	var dest *types.Array
	if len(args) == 0 {
		return types.NewArray()
	}

	count := 0
	for i, arg := range args {
		if !arg.IsArray() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(arg)))
			return nil
		}
		count += arg.Array().Len()
	}

	arr := types.NewArrayCap(count)

	args[0].Array().Each(func(key types.ArrayKey, value types.Zval) {
		if key.IsStrKey() {
			arr.KeyUpdate(key.StrKey(), value)
		} else {
			arr.Append(value)
		}
	})

	if recursive {
		for _, arg := range args[1:] {
			PhpArrayMergeRecursive(ctx, dest, arg.Array())
		}
	} else {
		for _, arg := range args[1:] {
			PhpArrayMerge(dest, arg.Array())
		}
	}
	return arr
}
func ZifArrayMerge(ctx *php.Context, returnValue zpp.Ret, _ zpp.Opt, arrays []types.Zval) {
	arr := arrayMergeWrapper(ctx, arrays, false)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}
func ZifArrayMergeRecursive(ctx *php.Context, returnValue zpp.Ret, _ zpp.Opt, arrays []types.Zval) {
	arr := arrayMergeWrapper(ctx, arrays, true)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}

// @zif(numArgs="1,")
func ZifArrayReplace(ctx *php.Context, arrays []types.Zval) types.Zval {
	php.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(ctx, arrays, 0)
	if !ok {
		return types.Null
	}

	dest := arrayHts[0].Clone()
	for _, array := range arrayHts[1:] {
		types.ArrayMerge(dest, array, true)
	}
	return types.ZvalArray(dest)
}

// @zif(numArgs="1,")
func ZifArrayReplaceRecursive(ctx *php.Context, arrays []types.Zval) types.Zval {
	php.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(ctx, arrays, 0)
	if !ok {
		return types.Null
	}

	dest := arrayHts[0].Clone()
	for _, array := range arrayHts[1:] {
		PhpArrayReplaceRecursive(ctx, dest, array)
	}
	return types.ZvalArray(dest)
}
func ZifArrayKeys(ctx *php.Context, array *types.Array, _ zpp.Opt, searchValue *types.Zval, strict bool) *types.Array {
	arrLen := array.Len()

	/* Base case: empty input */
	if arrLen == 0 {
		return types.NewArray()
	}

	/* Initialize return array */
	if searchValue != nil {
		keys := types.NewArray()
		if strict {
			array.Each(func(key types.ArrayKey, entry types.Zval) {
				entry = entry.DeRef()
				if php.ZvalIsIdentical(ctx, *searchValue, entry) {
					keys.Append(key.ToZval())
				}
			})
		} else {
			array.Each(func(key types.ArrayKey, entry types.Zval) {
				if php.FastEqualFunction(ctx, *searchValue, entry) {
					keys.Append(key.ToZval())
				}
			})
		}
		return keys
	} else {
		keys := types.NewArrayCap(arrLen)
		array.Each(func(key types.ArrayKey, _ types.Zval) {
			keys.Append(key.ToZval())
		})
		return keys
	}
}
func ZifArrayKeyFirst(array *types.Array) types.Zval {
	pair := array.First()
	if !pair.IsValid() {
		return types.Null
	}
	return pair.Key.ToZval()
}
func ZifArrayKeyLast(array *types.Array) types.Zval {
	pair := array.Last()
	if !pair.IsValid() {
		return types.Null
	}
	return pair.Key.ToZval()
}
func ZifArrayValues(array *types.Array) *types.Array {
	arrLen := array.Len()
	if arrLen == 0 {
		return types.NewArray()
	}

	/* Initialize return array */
	values := types.NewArrayCap(array.Len())
	array.Each(func(_ types.ArrayKey, entry types.Zval) {
		values.Append(entry)
	})
	return values
}
func ZifArrayCountValues(ctx *php.Context, array *types.Array) *types.Array {
	retArr := types.NewArray()
	array.Each(func(_ types.ArrayKey, entry types.Zval) {
		entry = entry.DeRef()
		var key types.ArrayKey
		if entry.IsLong() {
			key = types.IdxKey(entry.Long())
		} else if entry.IsString() {
			key = types.NumericKey(entry.String())
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Can only count STRING and INTEGER values!")
			return
		}

		if keyCount := retArr.Find(key); keyCount.IsNotUndef() {
			retArr.Update(key, types.ZvalLong(keyCount.Long()+1))
		} else {
			retArr.Update(key, types.ZvalLong(1))
		}
	})

	return retArr
}
func ArrayColumnParamHelper(ctx *php.Context, param *types.Zval, name string) bool {
	switch param.Type() {
	case types.IsDouble:
		param.SetLong(php.ZvalGetLong(ctx, *param))
		return true
	case types.IsLong:
		return true
	case types.IsObject:
		str, ok := php.ZvalTryGetStr(ctx, *param)
		if !ok {
			return false
		}
		param.SetString(str)
		return true
	case types.IsString:
		return true
	default:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("The %s key should be either a string or an integer", name))
		return false
	}
}
func ArrayColumnFetchProp(data *types.Zval, name *types.Zval) types.Zval {
	//var rv types.Zval
	var prop types.Zval
	if data.IsObject() {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		//if data.Object().HasPropertyEx(name, php.ZEND_PROPERTY_EXISTS) || data.Object().HasPropertyEx(name, php.ZEND_PROPERTY_ISSET) {
		//	prop = *data.Object().ReadProperty(name, php.BP_VAR_R, &rv)
		//	if prop.IsNotUndef() {
		//		prop = prop.DeRef()
		//	}
		//}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.IsArray() {
		if name.IsString() {
			prop = data.Array().SymtableFind(name.String())
		} else if name.IsLong() {
			prop = data.Array().IndexFind(name.Long())
		}
		if prop.IsNotUndef() {
			prop = prop.DeRef()
		}
	}
	return prop
}
func ZifArrayColumn(ctx *php.Context, array *types.Array, columnKey *types.Zval, _ zpp.Opt, indexKey zpp.ZvalNullable) (*types.Array, bool) {
	if columnKey != nil && !ArrayColumnParamHelper(ctx, columnKey, "column") || indexKey != nil && !ArrayColumnParamHelper(ctx, indexKey, "index") {
		return nil, false
	}

	retArr := types.NewArrayCap(array.Len())
	if indexKey == nil {
		array.Each(func(_ types.ArrayKey, data types.Zval) {
			var columnVal types.Zval
			data = data.DeRef()
			if columnKey.IsUndef() {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(&data, columnKey); columnVal.IsUndef() {
				return
			}

			retArr.Append(columnVal)
		})
	} else {
		array.Each(func(key types.ArrayKey, data types.Zval) {
			data = data.DeRef()

			// col
			var columnVal types.Zval
			if columnKey == nil {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(&data, columnKey); columnVal.IsUndef() {
				return
			}

			// key
			var keyVal = ArrayColumnFetchProp(&data, indexKey)
			if keyVal.IsNotUndef() {
				switch keyVal.Type() {
				case types.IsString:
					retArr.SymtableUpdate(keyVal.String(), columnVal)
				case types.IsLong:
					retArr.IndexUpdate(keyVal.Long(), columnVal)
				case types.IsObject:
					retArr.SymtableUpdate(php.ZvalGetStrVal(ctx, keyVal), columnVal)
				case types.IsNull:
					retArr.KeyUpdate("", columnVal)
				case types.IsDouble:
					retArr.IndexUpdate(php.DoubleToLong(keyVal.Double()), columnVal)
				case types.IsTrue:
					retArr.IndexUpdate(1, columnVal)
				case types.IsFalse:
					retArr.IndexUpdate(0, columnVal)
				case types.IsResource:
					retArr.IndexUpdate(keyVal.ResourceHandle(), columnVal)
				default:
					retArr.Append(columnVal)
				}
			} else {
				retArr.Append(columnVal)
			}

		})
	}
	return retArr, true
}
func ZifArrayReverse(array *types.Array, _ zpp.Opt, preserveKeys bool) *types.Array {
	retArr := types.NewArrayCap(array.Len())
	array.EachReserve(func(key types.ArrayKey, value types.Zval) {
		if preserveKeys || key.IsStrKey() {
			retArr.Add(key, value)
		} else { // !preserveKeys && !key.IsStrKey()
			retArr.Append(value)
		}
	})
	return retArr
}

const maxPadInOneTimes = 1048576

func ZifArrayPad(ctx *php.Context, array *types.Array, padSize int, padValue types.Zval) (*types.Array, bool) {
	/* Do some initial calculations */
	inputSize := array.Len()
	padSizeAbs := padSize
	if padSizeAbs < 0 {
		padSizeAbs = -padSizeAbs
	}
	if padSizeAbs < 0 || padSizeAbs-inputSize > maxPadInOneTimes {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return nil, false
	}

	/**
	 * notice: 注意此处的差别:
	 * - 当原数组的长度 >= padSize 时，返回数组保留了原数组包括内部指针等信息；
	 * - 当原数组的长度 < padSize 时，产生的新数组内部指针等信息未设置(等于初始值)；
	 */
	if inputSize >= padSizeAbs {
		return array.Clone(), true
	}

	numPads := padSizeAbs - inputSize
	retArr := types.NewArrayCap(padSizeAbs)
	if padSize < 0 {
		for i := 0; i < numPads; i++ {
			retArr.Append(padValue)
		}
	}
	array.Each(func(key types.ArrayKey, value types.Zval) {
		if key.IsStrKey() {
			retArr.KeyAdd(key.StrKey(), value)
		} else {
			retArr.Append(value)
		}
	})
	if padSize > 0 {
		for i := 0; i < numPads; i++ {
			retArr.Append(padValue)
		}
	}
	return retArr, true
}
func ZifArrayFlip(ctx *php.Context, array *types.Array) *types.Array {
	retArr := types.NewArrayCap(array.Len())
	array.Each(func(key types.ArrayKey, value types.Zval) {
		value = value.DeRef()
		if value.IsLong() {
			retArr.IndexUpdate(value.Long(), key.ToZval())
		} else if value.IsString() {
			retArr.SymtableUpdate(value.String(), key.ToZval())
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	})
	return retArr
}
func ZifArrayChangeKeyCase(array *types.Array, _ zpp.Opt, case_ *int) *types.Array {
	var caseFlag = lang.Option(case_, CASE_LOWER)
	retArr := types.NewArrayCap(array.Len())
	array.Each(func(key types.ArrayKey, value types.Zval) {
		if key.IsStrKey() {
			if caseFlag == CASE_LOWER {
				key = types.StrKey(ascii.StrToLower(key.StrKey()))
			} else {
				key = types.StrKey(ascii.StrToUpper(key.StrKey()))
			}
		}

		retArr.Update(key, value)
	})
	return retArr
}
func ZifArrayUnique(ctx *php.Context, arg *types.Array, _ zpp.Opt, flags *int) *types.Array {
	var sortType = lang.Option(flags, PHP_SORT_STRING)

	if sortType == PHP_SORT_STRING {
		existValues := make(map[string]bool, arg.Len())
		retArr := types.NewArrayCap(arg.Len())
		arg.Each(func(key types.ArrayKey, val types.Zval) {
			var strVal = php.ZvalGetStrVal(ctx, val)
			if _, exists := existValues[strVal]; !exists {
				retArr.Add(key, val)
			}

		})
		return retArr
	}

	cmp := getArrayValueComparer(ctx, sortType, false)
	retArr := arg.Dup()

	// 取出数组所有键值对，先排序，后找出重复元素并删除
	pairs := arg.Pairs()
	sort.SliceStable(pairs, func(i, j int) bool {
		return cmp.Compare(pairs[i], pairs[j]) < 0
	})

	last := 0
	for i := 1; i < len(pairs); i++ {
		if cmp.Compare(pairs[last], pairs[i]) != 0 {
			last = i
		} else {
			pair := pairs[i]
			retArr.Delete(pair.Key)
		}
	}

	return retArr
}
func ZvalCompare(ctx *php.Context, first types.Zval, second types.Zval) int {
	return php.StringCompareFunction(ctx, first, second)
}

// @zif(numArgs="2,")
func ZifArrayIntersectKey(ctx *php.Context, arrays []types.Zval) (*types.Array, bool) {
	return arrayIntersectKeyWrapper(ctx, arrays, nil)
}

// @zif(numArgs="3,")
func ZifArrayIntersectUkey(ctx *php.Context, arrays []types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(ctx, callbackKeyCompareFunc)
	return arrayIntersectWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="2,")
func ZifArrayIntersect(ctx *php.Context, arrays []types.Zval) (*types.Array, bool) {
	cmp := arrayValueComparerWithCtx(ctx, php.StringCompareFunction)
	return arrayIntersectWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="3,")
func ZifArrayUintersect(ctx *php.Context, arrays []types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserValueComparer(ctx, callbackDataCompareFunc)
	return arrayIntersectWrapper(ctx, arrays, cmp)
}
func ZifArrayIntersectAssoc(ctx *php.Context, arrays []types.Zval) (*types.Array, bool) {
	cmp := types.ArrayValueComparerFunc(func(v1, v2 types.Zval) int {
		return ZvalCompare(ctx, v1, v2)
	})
	return arrayIntersectKeyWrapper(ctx, arrays, cmp)
}
func ZifArrayIntersectUassoc(ctx *php.Context, arrays []types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(ctx, callbackKeyCompareFunc),
		arrayValueComparerWithCtx(ctx, php.StringCompareFunction),
	)
	return arrayIntersectWrapper(ctx, arrays, cmp)
}
func ZifArrayUintersectAssoc(ctx *php.Context, arrays []types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserValueComparer(ctx, callbackDataCompareFunc)
	return arrayIntersectKeyWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="4,")
func ZifArrayUintersectUassoc(ctx *php.Context, arrays []types.Zval, callbackDataCompareFunc zpp.Callable, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(ctx, callbackKeyCompareFunc),
		arrayUserValueComparer(ctx, callbackDataCompareFunc),
	)
	return arrayIntersectWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="2,")
func ZifArrayDiffKey(ctx *php.Context, arrays []types.Zval) (*types.Array, bool) {
	return arrayDiffKeyWrapper(ctx, arrays, nil)
}

// @zif(numArgs="3,")
func ZifArrayDiffUkey(ctx *php.Context, arrays []types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(ctx, callbackKeyCompFunc)
	return arrayDiffWrapper(ctx, arrays, cmp)
}

func simpleArrayDiff(ctx *php.Context, array *types.Array, arrays []*types.Array) *types.Array {
	if array.Len() == 0 {
		return types.NewArray()
	}

	// array.Len() > 1
	exclude := make(map[string]bool)
	for _, diffArray := range arrays {
		diffArray.Each(func(_ types.ArrayKey, value types.Zval) {
			str := php.ZvalGetStrVal(ctx, value)
			exclude[str] = true
		})
	}
	if len(exclude) == 0 {
		return array
	}

	retArr := types.NewArrayCap(array.Len())
	array.Each(func(key types.ArrayKey, value types.Zval) {
		str := php.ZvalGetStrVal(ctx, value)
		if _, excluded := exclude[str]; !excluded {
			retArr.Add(key, value)
		}
	})
	return retArr
}

// @zif(numArgs="2,")
func ZifArrayDiff(ctx *php.Context, arrays []types.Zval) types.Zval {
	arrayHts, ok := checkArrayArgs(ctx, arrays, 0)
	if !ok {
		return types.Null
	}
	retArr := simpleArrayDiff(ctx, arrayHts[0], arrayHts[1:])
	return types.ZvalArray(retArr)
}

// @zif(numArgs="3,")
func ZifArrayUdiff(ctx *php.Context, arrays []types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserValueComparer(ctx, callbackDataCompFunc)
	return arrayDiffWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="3,")
func ZifArrayDiffAssoc(ctx *php.Context, arrays []types.Zval) (*types.Array, bool) {
	cmp := arrayValueComparerWithCtx(ctx, ZvalCompare)
	return arrayDiffKeyWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="3,")
func ZifArrayDiffUassoc(ctx *php.Context, arrays []types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(ctx, callbackKeyCompFunc),
		arrayValueComparerWithCtx(ctx, php.StringCompareFunction),
	)
	return arrayDiffWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="2,")
func ZifArrayUdiffAssoc(ctx *php.Context, arrays []types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserValueComparer(ctx, callbackDataCompFunc)
	return arrayDiffKeyWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="4,")
func ZifArrayUdiffUassoc(ctx *php.Context, arrays []types.Zval, callbackDataCompFunc zpp.Callable, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(ctx, callbackKeyCompFunc),
		arrayUserValueComparer(ctx, callbackDataCompFunc),
	)
	return arrayDiffWrapper(ctx, arrays, cmp)
}

// @zif(numArgs="1,")
func ZifArrayMultisort(ctx *php.Context, args []types.Zval) bool {
	var parseState = [...]int{0, 0}

	arrays := make([]types.Zval, 0, len(args))
	multisortFunc := make([]types.ArrayComparer, 0, len(args))

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */
	var sortOrder = PHP_SORT_ASC
	var sortType = PHP_SORT_REGULAR
	for i, arg := range args {
		arg = arg.DeRef()
		if arg.IsArray() {
			types.SeparateArray(&arg)

			/* We see the next array, so we update the sort flags of
			 * the previous array and reset the sort flags. */
			if i > 0 {
				multisortFunc = append(multisortFunc, getArrayValueComparer(ctx, sortType, sortOrder != PHP_SORT_ASC))
				sortOrder = PHP_SORT_ASC
				sortType = PHP_SORT_REGULAR
			}
			arrays = append(arrays, arg)

			/* Next one may be an array or a list of sort flags. */
			for k := 0; k < MULTISORT_LAST; k++ {
				parseState[k] = 1
			}

			/* Next one may be an array or a list of sort flags. */

		} else if arg.IsLong() {
			switch arg.Long() &^ PHP_SORT_FLAG_CASE {
			case PHP_SORT_ASC:
				fallthrough
			case PHP_SORT_DESC:
				/* flag allowed here */
				if parseState[MULTISORT_ORDER] == 1 {
					/* Save the flag and make sure then next arg is not the current flag. */
					if arg.Long() == PHP_SORT_DESC {
						sortOrder = PHP_SORT_DESC
					} else {
						sortOrder = PHP_SORT_ASC
					}
					parseState[MULTISORT_ORDER] = 0
				} else {
					php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1))
					return false
				}
			case PHP_SORT_REGULAR:
				fallthrough
			case PHP_SORT_NUMERIC:
				fallthrough
			case PHP_SORT_STRING:
				fallthrough
			case PHP_SORT_NATURAL:
				fallthrough
			case PHP_SORT_LOCALE_STRING:
				/* flag allowed here */
				if parseState[MULTISORT_TYPE] == 1 {
					/* Save the flag and make sure then next arg is not the current flag. */
					sortType = arg.Long()
					parseState[MULTISORT_TYPE] = 0
				} else {
					php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1))
					return false
				}
			default:
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Argument #%d is an unknown sort flag", i+1))
				return false
			}
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Argument #%d is expected to be an array or a sort flag", i+1))
			return false
		}
	}

	/* Take care of the last array sort flags. */
	multisortFunc = append(multisortFunc, getArrayValueComparer(ctx, sortType, sortOrder != PHP_SORT_ASC))

	/* Make sure the arrays are of the same size. */
	arraySize := arrays[0].Array().Len()
	for _, array := range arrays {
		if array.Array().Len() != arraySize {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Array sizes are inconsistent")
			return false
		}
	}

	/* If all arrays are empty we don't need to do anything. */
	if arraySize == 1 {
		return true
	}

	/**
	 * 将所有数组中的键值对取出组成并组成矩阵。(matrix[idx][i] 表示第 i 个数组的第 idx 个键值对)
	 */
	var matrix = make([][]types.ArrayPair, arraySize)
	for _, array := range arrays {
		idx := 0
		array.Array().Each(func(key types.ArrayKey, value types.Zval) {
			matrix[idx] = append(matrix[idx], types.MakeArrayPair(key, value))
			idx++
		})
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */
	sort.SliceStable(matrix, func(idx1, idx2 int) bool {
		col1 := matrix[idx1]
		col2 := matrix[idx2]
		for i, sortFunc := range multisortFunc {
			result := sortFunc.Compare(col1[i], col2[i])
			if result < 0 {
				return true
			} else if result > 0 {
				return false
			}
		}
		return false
	})

	/* Restructure the arrays based on sorted indirect - this is mostly taken from zend_hash_sort() function. */
	for i, array := range arrays {
		newHash := types.NewArrayCap(array.Array().Len())
		for idx := 0; idx < arraySize; idx++ {
			pair := matrix[idx][i]
			if pair.Key.IsStrKey() {
				newHash.KeyAdd(pair.Key.StrKey(), pair.Val)
			} else {
				newHash.Append(pair.Val)
			}
		}
		array.Array().SetDataByArray(newHash)
	}

	return true
}
func ZifArrayRand(ctx *php.Context, arg *types.Array, _ zpp.Opt, numReq_ *int) types.Zval {
	numReq := lang.Option(numReq_, 1)

	numAvail := arg.Len()
	if numAvail == 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Array is empty")
		return types.Undef
	}
	if numReq <= 0 || numReq > numAvail {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Second argument has to be between 1 and the number of elements in the array")
		return types.Undef
	}

	keys := arg.Keys()
	if numReq == 1 {
		randIdx := PhpMtRandRange(ctx, 0, numAvail-1)
		return keys[randIdx].ToZval()
	} else {
		randIdxSet := make(map[int]bool)
		negative := false
		if numReq > arg.Len()/2 {
			negative = true
			numReq = arg.Len() - numReq
		}
		for i := numReq - 1; i >= 0; {
			randIdx := PhpMtRandRange(ctx, 0, numAvail-1)
			if !randIdxSet[randIdx] {
				randIdxSet[randIdx] = true
				i--
			}
		}

		retArr := types.NewArrayCap(numReq)
		for i, key := range keys {
			if !negative && randIdxSet[i] || negative && !randIdxSet[i] {
				retArr.Append(key.ToZval())
			}
		}
		return types.ZvalArray(retArr)
	}
}
func ZifArraySum(ctx *php.Context, array *types.Array) types.Zval {
	ret := types.ZvalLong(0)
	array.Each(func(_ types.ArrayKey, entry types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		num := php.ConvertScalarToNumber(ctx, entry)
		ret = php.OpAdd(ctx, ret, num)
	})
	return ret
}
func ZifArrayProduct(ctx *php.Context, array *types.Array) types.Zval {
	var num types.Zval
	ret := types.ZvalLong(1)
	array.Each(func(_ types.ArrayKey, entry types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		num = php.ConvertScalarToNumber(ctx, entry)
		if num.IsLong() && ret.IsLong() {
			dval := float64(num.Long()) * float64(ret.Long())
			if float64(math.MinInt) <= dval && dval <= float64(math.MaxInt) {
				ret.SetLong(ret.Long() * num.Long())
				return
			}
		}
		ret.SetDouble(php.ZvalGetDouble(ctx, ret) * php.ZvalGetDouble(ctx, num))
	})
	return ret
}
func ZifArrayReduce(array *types.Array, callback zpp.Callable, _ zpp.Opt, initial *types.Zval) types.Zval {
	var result types.Zval
	if initial == nil {
		result = types.Null
	} else {
		result = *initial
	}

	array.Each(func(key types.ArrayKey, value types.Zval) {
		newResult, ok := callback.Call(result, value)
		if ok && newResult.IsNotUndef() {
			result = newResult
		}
	})

	return php.NewZvalZval(result, true, true)
}
func ZifArrayFilter(ctx *php.Context, array_ *types.Array, _ zpp.Opt, callback zpp.Callable, mode int) *types.Array {
	retArr := types.NewArray()
	if array_.Len() == 0 {
		return retArr
	}

	// parse use filter
	var filter func(key types.ArrayKey, value types.Zval) (keep bool, succ bool)
	if callback != nil {
		filter = func(key types.ArrayKey, value types.Zval) (bool, bool) {
			var retVal types.Zval
			var ok bool
			switch mode {
			case ARRAY_FILTER_USE_KEY:
				retVal, ok = callback.Call(key.ToZval())
			case ARRAY_FILTER_USE_BOTH:
				retVal, ok = callback.Call(value, key.ToZval())
			default: // 0
				retVal, ok = callback.Call(value)
			}
			if !ok {
				return false, false
			}
			return php.ZvalIsTrue(ctx, retVal), true
		}
	}

	// filter
	array_.EachEx(func(key types.ArrayKey, value types.Zval) error {
		var keep bool
		if filter != nil {
			var callSucc bool
			keep, callSucc = filter(key, value)
			if !callSucc {
				// 调用 user callback 失败，中断遍历
				return lang.BreakErr
			}
		} else {
			keep = php.ZvalIsTrue(ctx, value)
		}

		if keep {
			retArr.Update(key, value)
		}
		return nil
	})
	return retArr
}

func arrayMapSingle(callback zpp.Callable, array *types.Array) *types.Array {
	retArr := types.NewArrayCap(array.Len())
	err := array.EachEx(func(key types.ArrayKey, value types.Zval) error {
		retVal, ok := callback.Call(value)
		if !ok || retVal.IsUndef() {
			// 调用 callback 失败，中断流程
			return lang.BreakErr
		}

		retArr.Add(key, retVal)
		return nil
	})
	if err != nil {
		return nil
	}
	return retArr
}

func arrayMapMulti(callback zpp.Callable, arrays []*types.Array) *types.Array {
	len_ := arrays[0].Len()
	for _, array := range arrays {
		if array.Len() > len_ {
			len_ = array.Len()
		}
	}

	argMatrix := make([][]types.Zval, len_)
	for _, array := range arrays {
		count := 0
		_ = array.EachEx(func(key types.ArrayKey, value types.Zval) error {
			argMatrix[count] = append(argMatrix[count], value)
			count++
			if count < len_ {
				return nil
			} else {
				return lang.BreakErr
			}
		})
		for i := len_; i < count; i++ {
			argMatrix[i] = append(argMatrix[i], types.Null)
		}
	}

	retArr := types.NewArrayCap(len_)
	for _, argColumns := range argMatrix {
		if callback == nil {
			retArr.Append(types.ZvalArray(types.NewArrayOfZval(argColumns)))
		} else {
			retVal, ok := callback.Call(argColumns...)
			if !ok || retVal.IsUndef() {
				// 调用 callback 失败，中断流程
				return nil
			}
			retArr.Append(retVal)
		}
	}
	return retArr
}

// @zif(numArgs="2,")
func ZifArrayMap(ctx *php.Context, callback zpp.Callable, arrays []types.Zval) types.Zval {
	php.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(ctx, arrays, 1)
	if !ok {
		return types.Null
	}

	var retArr *types.Array
	if len(arrayHts) == 1 {
		retArr = arrayMapSingle(callback, arrayHts[0])
	} else {
		retArr = arrayMapMulti(callback, arrayHts)
	}

	if retArr == nil {
		return types.Null
	}
	return types.ZvalArray(retArr)
}

// @zif(alias="key_exists")
func ZifArrayKeyExists(ctx *php.Context, key *types.Zval, array zpp.ArrayOrObjectZval) bool {
	var ht *types.Array
	if array.IsArray() {
		ht = array.Array()
	} else {
		//ht = php.ZendGetPropertiesFor(&array, types.PropPurposeArrayCast)
		php.ErrorDocRef(ctx, "", perr.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. Use isset() or property_exists() instead")
	}
	switch key.Type() {
	case types.IsString:
		return ht.SymtableExists(key.String())
	case types.IsLong:
		return ht.IndexExists(key.Long())
	case types.IsNull:
		return ht.KeyExists("")
	default:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The first argument should be either a string or an integer")
		return false
	}
}
func ZifArrayChunk(ctx *php.Context, array *types.Array, length int, _ zpp.Opt, preserveKeys bool) types.Zval {
	if length < 1 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Size parameter expected to be greater than 0")
		return types.Null
	}

	chunkCount := (array.Len()-1)/length + 1 // ceil(len/length)
	retArr := types.NewArrayCap(chunkCount)

	var currChunk *types.Array = nil
	itemCount := 0
	array.Each(func(key types.ArrayKey, entry types.Zval) {
		if itemCount%length == 0 {
			currChunk = types.NewArrayCap(length)
			retArr.Append(types.ZvalArray(currChunk))
		}

		if preserveKeys {
			currChunk.Update(key, entry)
		} else {
			currChunk.Append(entry)
		}
		itemCount++
	})
	return php.Array(retArr)
}
func ZifArrayCombine(ctx *php.Context, keys *types.Array, values *types.Array) (*types.Array, bool) {
	if keys.Len() != values.Len() {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Both parameters should have an equal number of elements")
		return nil, false
	}
	if keys.Len() == 0 {
		return types.NewArray(), true
	}

	zvKeys := keys.Values()
	zvValues := values.Values()
	retArr := types.NewArrayCap(keys.Len())
	for i, key := range zvKeys {
		if i > len(zvValues) {
			break
		}
		value := zvValues[i]

		if key.IsLong() {
			retArr.IndexUpdate(key.Long(), value)
		} else {
			strKey := php.ZvalGetStrVal(ctx, key)
			retArr.SymtableUpdate(strKey, value)
		}
	}
	return retArr, true
}

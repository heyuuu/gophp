package array

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
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
func ZmStartupArray(type_ int, module_number int) int {
	zend.RegisterLongConstant("EXTR_OVERWRITE", EXTR_OVERWRITE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_SKIP", EXTR_SKIP, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_PREFIX_SAME", EXTR_PREFIX_SAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_PREFIX_ALL", EXTR_PREFIX_ALL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_PREFIX_INVALID", EXTR_PREFIX_INVALID, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_PREFIX_IF_EXISTS", EXTR_PREFIX_IF_EXISTS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_IF_EXISTS", EXTR_IF_EXISTS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("EXTR_REFS", EXTR_REFS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_ASC", PHP_SORT_ASC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_DESC", PHP_SORT_DESC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_REGULAR", PHP_SORT_REGULAR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_NUMERIC", PHP_SORT_NUMERIC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_STRING", PHP_SORT_STRING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_LOCALE_STRING", PHP_SORT_LOCALE_STRING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_NATURAL", PHP_SORT_NATURAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SORT_FLAG_CASE", PHP_SORT_FLAG_CASE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CASE_LOWER", CASE_LOWER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CASE_UPPER", CASE_UPPER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("COUNT_NORMAL", COUNT_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("COUNT_RECURSIVE", COUNT_RECURSIVE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ARRAY_FILTER_USE_BOTH", ARRAY_FILTER_USE_BOTH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ARRAY_FILTER_USE_KEY", ARRAY_FILTER_USE_KEY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}

func ZifKrsort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := PhpGetKeyCompareFunc(sortFlags, true)
	arg.Array().Sort(cmp, false)
	return true
}
func ZifKsort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := PhpGetKeyCompareFunc(sortFlags, false)
	arg.Array().Sort(cmp, false)
	return true
}

func PhpCountRecursive(ht *types.Array) int {
	if !ht.IsImmutable() {
		if ht.IsRecursive() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		ht.ProtectRecursive()
	}

	cnt := ht.Count()
	ht.Foreach(func(key types.ArrayKey, value *types.Zval) {
		value = types.ZVAL_DEREF(value)
		if value.IsArray() {
			cnt += PhpCountRecursive(value.Array())
		}
	})

	if !ht.IsImmutable() {
		ht.UnprotectRecursive()
	}
	return cnt
}

//@zif -alias sizeof
func ZifCount(var_ *types.Zval, _ zpp.Opt, mode int) int {
	var array = var_
	var cnt zend.ZendLong
	switch array.GetType() {
	case types.IS_NULL:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 0
	case types.IS_ARRAY:
		if mode != COUNT_RECURSIVE {
			cnt = array.Array().Count()
		} else {
			cnt = PhpCountRecursive(array.Array())
		}
		return cnt
	case types.IS_OBJECT:
		var retval types.Zval
		var long int

		/* first, we check if the handler is defined */
		if types.Z_OBJ_HT_P(array).GetCountElements() != nil {
			long = 1
			if types.SUCCESS == types.Z_OBJ_HT(*array).GetCountElements()(array, &long) {
				return long
			}
			if zend.EG__().GetException() != nil {
				return long
			}
		}

		/* if not and the object implements Countable we call its count() method */
		if zend.InstanceofFunction(types.Z_OBJCE_P(array), zend.ZendCeCountable) != 0 {
			zend.ZendCallMethodWith0Params(array, nil, nil, "count", &retval)
			if retval.IsNotUndef() {
				long = zend.ZvalGetLong(&retval)
			}
			return long
		}

		/* If There's no handler and it doesn't implement Countable then add a warning */
		core.PhpErrorDocref(nil, faults.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 1
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return 1
	}
}
func ZifNatsort(arg zpp.RefArray) bool {
	cmp := PhpGetKeyCompareFunc(PHP_SORT_NATURAL, false)
	arg.Array().Sort(cmp, false)
	return true
}
func ZifNatcasesort(arg zpp.RefArray) bool {
	cmp := PhpGetKeyCompareFunc(PHP_SORT_NATURAL|PHP_SORT_FLAG_CASE, false)
	arg.Array().Sort(cmp, false)
	return true
}
func ZifAsort(arg zpp.RefZval, _ zpp.Opt, sortFlags int) bool {
	cmp := phpGetDataCompareFunc(sortFlags, false)
	arg.Array().Sort(cmp, false)
	return true
}
func ZifArsort(arg zpp.RefZval, _ zpp.Opt, sortFlags int) bool {
	cmp := phpGetDataCompareFunc(sortFlags, true)
	arg.Array().Sort(cmp, false)
	return true
}
func ZifSort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := phpGetDataCompareFunc(sortFlags, false)
	arg.Array().Sort(cmp, true)
	return true
}
func ZifRsort(arg zpp.RefArray, _ zpp.Opt, sortFlags int) bool {
	cmp := phpGetDataCompareFunc(sortFlags, true)
	arg.Array().Sort(cmp, true)
	return true
}

func phpUsortEx(array *types.Zval, compareFunc types.ArrayComparer, renumber bool) bool {
	arr := array.Array()
	if arr.Len() > 0 {
		/* Copy array, so the in-place modifications will not be visible to the callback function */
		arr = types.ZendArrayDup(arr)
		arr.Sort(compareFunc, renumber)
		array.SetArray(arr)
	}
	return true
}
func ZifUsort(arg zpp.DerefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserDataComparer(cmpFunction)
	phpUsortEx(arg, cmp, true)
	return true
}
func ZifUasort(arg zpp.DerefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserDataComparer(cmpFunction)
	phpUsortEx(arg, cmp, false)
	return true
}
func ZifUksort(arg zpp.DerefArray, cmpFunction zpp.Callable) bool {
	var cmp = arrayUserKeyComparer(cmpFunction)
	phpUsortEx(arg, cmp, false)
	return true
}
func ZifEnd(executeData zpp.Ex, return_value zpp.Ret, arg zpp.RefArray) {
	var array *types.Array
	var entry *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			array = fp.ParseArrayHtEx(false, true)
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	types.ZendHashInternalPointerEnd(array)
	if zend.USED_RET(executeData) {
		if b.Assign(&entry, types.ZendHashGetCurrentData(array)) == nil {
			return_value.SetFalse()
			return
		}
		if entry.IsIndirect() {
			entry = entry.Indirect()
		}
		types.ZVAL_COPY_DEREF(return_value, entry)
	}
}

func ZifPrev(array zpp.RefArrayHt) *types.Zval {
	types.ZendHashMoveBackwards(array)
	entry := types.ZendHashGetCurrentData(array)
	if entry == nil {
		return types.NewZvalBool(false)
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
	}
	return entry.DeRef()
}
func ZifNext(array zpp.RefArrayHt) *types.Zval {
	types.ZendHashMoveForward(array)
	entry := types.ZendHashGetCurrentData(array)
	if entry == nil {
		return types.NewZvalBool(false)
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
	}
	return entry.DeRef()
}

func ZifReset(array zpp.RefArrayHt) *types.Zval {
	types.ZendHashInternalPointerReset(array)
	entry := types.ZendHashGetCurrentData(array)
	if entry == nil {
		return types.NewZvalBool(false)
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
	}
	return entry.DeRef()
}

//@zif -alias pos
func ZifCurrent(array zpp.ArrayOrObjectHt) (*types.Zval, bool) {
	_, val, ok := array.Current(false)
	if !ok {
		return nil, false
	}

	if val.IsIndirect() {
		val = val.Indirect()
	}
	return val.DeRef(), true
}
func ZifKey(executeData zpp.Ex, return_value zpp.Ret, array zpp.ArrayOrObjectHt) {
	types.ZendHashGetCurrentKeyZval(array, return_value)
}

func ZifMin(arg *types.Zval, args []*types.Zval) *types.Zval {
	/* mixed min ( array $values ) */
	if len(args) == 0 {
		if !arg.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "When only one parameter is given, it must be an array")
			return types.NewZvalNull()
		} else {
			result := arg.Array().Min(arrayDataComparer(arrayDataCompare))
			if result != nil {
				return result.GetVal()
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Array must contain at least one element")
				return types.NewZvalBool(false)
			}
		}
	} else {
		result := arg
		for _, value := range args {
			if arrayDataCompare(result, value) > 0 {
				result = value
			}
		}
		return result
	}
}
func ZifMax(arg *types.Zval, args []*types.Zval) *types.Zval {
	/* mixed max ( array $values ) */
	if len(args) == 0 {
		if !arg.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "When only one parameter is given, it must be an array")
			return types.NewZvalNull()
		} else {
			result := arg.Array().Max(arrayDataComparer(arrayDataCompare))
			if result != nil {
				return result.GetVal()
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Array must contain at least one element")
				return types.NewZvalBool(false)
			}
		}
	} else {
		result := arg
		for _, value := range args {
			if arrayDataCompare(result, value) < 0 {
				result = value
			}
		}
		return result
	}
}

func arrayUserWalkHandler(handler zpp.Callable, userdata *types.Zval) func(value *types.Zval, key types.ArrayKey) bool {
	return func(value *types.Zval, key types.ArrayKey) bool {
		var ok bool
		if userdata != nil {
			_, ok = handler.Call(value, key.ToZval(), userdata)
		} else {
			_, ok = handler.Call(value, key.ToZval())
		}
		return ok
	}
}

func arrayWalk(array *types.Zval, recursive bool, handler func(value *types.Zval, key types.ArrayKey) bool) bool {
	var targetHash = zend.HASH_OF(array)
	var result = true

	/* Iterate through hash */
	targetHash.ForeachEx(func(key types.ArrayKey, zv *types.Zval) bool {
		/* Skip undefined indirect elements */
		if zv.IsIndirect() {
			zv = zv.Indirect()
			if zv.IsUndef() {
				return true
			}

			/* Add type source for property references. */
			if zv.GetType() != types.IS_REFERENCE && array.IsType(types.IS_OBJECT) {
				var prop_info = zend.ZendGetTypedPropertyInfoForSlot(array.Object(), zv)
				if prop_info != nil {
					zv.SetNewRef(zv)
					zend.ZEND_REF_ADD_TYPE_SOURCE(zv.Reference(), prop_info)
				}
			}
		}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */
		types.ZVAL_MAKE_REF(zv)

		/* Back up hash position, as it may change */
		if recursive && types.Z_REFVAL_P(zv).IsType(types.IS_ARRAY) {
			var thash *types.Array
			var ref types.Zval
			types.ZVAL_COPY_VALUE(&ref, zv)
			zv = types.ZVAL_DEREF(zv)
			types.SeparateArray(zv)
			thash = zv.Array()
			if thash.IsRecursive() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
				return false
			}

			/* backup the fcall info and cache */
			thash.ProtectRecursive()
			result = arrayWalk(zv, recursive, handler)
			if types.Z_REFVAL(ref).IsType(types.IS_ARRAY) && thash == types.Z_REFVAL(ref).Array() {
				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */
				thash.UnprotectRecursive()
			}
		} else {
			/* Call the userland function */
			result = handler(zv, key)
		}
		if !result {
			return false
		}

		/* Reload array and position -- both may have changed */
		if zend.EG__().GetException() != nil {
			return false
		}

		return true
	})

	return result
}

func ZifArrayWalk(array zpp.RefArrayOrObject, callable zpp.Callable, _ zpp.Opt, arg *types.Zval) bool {
	handler := arrayUserWalkHandler(callable, arg)
	arrayWalk(array, false, handler)
	return true
}
func ZifArrayWalkRecursive(array zpp.RefArrayOrObject, callable zpp.Callable, _ zpp.Opt, arg *types.Zval) bool {
	handler := arrayUserWalkHandler(callable, arg)
	arrayWalk(array, true, handler)
	return true
}

func searchArray(value *types.Zval, array *types.Array, strict bool) *types.ArrayKey {
	var targetKey *types.ArrayKey = nil
	if strict {
		if value.IsLong() {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if entry.IsLong() && entry.Long() == value.Long() {
					targetKey = &key
					return false
				}
				return true
			})
		} else {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(value, entry) != 0 {
					targetKey = &key
					return false
				}
				return true
			})
		}
	} else {
		if value.IsLong() {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if zend.FastEqualCheckLong(value, entry) != 0 {
					targetKey = &key
					return false
				}
				return true
			})
		} else if value.IsString() {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if zend.FastEqualCheckString(value, entry) != 0 {
					targetKey = &key
					return false
				}
				return true
			})
		} else {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if zend.FastEqualCheckFunction(value, entry) != 0 {
					targetKey = &key
					return false
				}
				return true
			})
		}
	}
	return targetKey
}
func ZifInArray(needle *types.Zval, haystack *types.Array, _ zpp.Opt, strict bool) bool {
	key := searchArray(needle, haystack, strict)
	return key != nil
}
func ZifArraySearch(needle *types.Zval, haystack *types.Array, _ zpp.Opt, strict bool) *types.Zval {
	key := searchArray(needle, haystack, strict)
	if key == nil {
		return types.NewZvalBool(false)
	} else {
		return key.ToZval()
	}
}

func ZifArrayFill(startKey int, num int, val *types.Zval) (*types.Array, bool) {
	if num < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Number of elements can't be negative")
		return nil, false
	}
	if num == 0 {
		return types.NewArray(0), true
	}
	if num > math.MaxInt32 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Too many elements")
		return nil, false
	} else if startKey > math.MaxInt-num+1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
		return nil, false
	}

	// todo 尽量创建 packed array

	/* create hash */
	arr := types.NewArray(num)
	for i := 0; i < num; i++ {
		arr.IndexAdd(startKey+i, val)
	}
	return arr, true
}
func ZifArrayFillKeys(keys *types.Array, val *types.Zval) *types.Array {
	arr := types.NewArray(keys.Len())
	keys.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		entry = types.ZVAL_DEREF(entry)
		if entry.IsLong() {
			arr.IndexUpdate(entry.Long(), val)
		} else {
			key := zend.ZvalGetStrVal(entry)
			arr.SymtableUpdate(key, val)
		}
	})
	return arr
}
func rangeDouble(zLow *types.Zval, zHigh *types.Zval, step float64) ([]*types.Zval, bool) {
	b.Assert(step > 0)
	low := zend.ZvalGetDouble(zLow)
	high := zend.ZvalGetDouble(zHigh)
	if core.ZendIsInf(high) || core.ZendIsInf(low) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid range supplied: start=%0.0f end=%0.0f", low, high)
		return nil, false
	}
	if low > high {
		low, high = high, low
	}
	if high-low < step {
		core.PhpErrorDocref(nil, faults.E_WARNING, "step exceeds the specified range")
		return nil, false
	}

	size := (high-low)/step + 1
	if size >= float64(types.HT_MAX_SIZE) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high)
		return nil, false
	}

	arr := make([]*types.Zval, int(size))
	for i := range arr {
		arr[i] = types.NewZvalDouble(low + float64(i)*step)
	}
	return arr, true
}

func rangeLong(zLow *types.Zval, zHigh *types.Zval, step int) ([]*types.Zval, bool) {
	b.Assert(step > 0)
	low := zend.ZvalGetLong(zLow)
	high := zend.ZvalGetLong(zHigh)
	if high == low {
		return []*types.Zval{types.NewZvalLong(high)}, true
	}

	if low > high {
		low, high = high, low
	}
	if high-low < step {
		core.PhpErrorDocref(nil, faults.E_WARNING, "step exceeds the specified range")
		return nil, false
	}

	size := (high-low)/step + 1
	if size >= types.HT_MAX_SIZE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start=%d end=%d", low, high)
		return nil, false
	}

	arr := make([]*types.Zval, size)
	for i := range arr {
		arr[i] = types.NewZvalLong(low + i*step)
	}
	return arr, true
}

func rangeChar(low byte, high byte, step int) ([]*types.Zval, bool) {
	if low == high {
		return []*types.Zval{types.NewZvalString(string(low))}, true
	}
	if low > high {
		low, high = high, low
	}

	bStep := byte(step)
	size := (high-low)/bStep + 1
	arr := make([]*types.Zval, 0, size)
	for i := range arr {
		c := low + byte(i)*bStep
		arr[i] = types.NewZvalString(string(c))
	}
	return arr, true
}

func ZifRange(low_ *types.Zval, high_ *types.Zval, _ zpp.Opt, step_ *types.Zval) ([]*types.Zval, bool) {
	var zlow *types.Zval = low_
	var zhigh *types.Zval = high_
	var zstep *types.Zval = step_
	var isStepDouble = false

	var step = 1.0
	if zstep != nil {
		if zstep.IsDouble() {
			isStepDouble = true
		} else if zstep.IsString() {
			r := conv.ParseNumber(zstep.StringVal())
			if r.IsInt() {
				// pass
			} else if r.IsFloat() {
				isStepDouble = true
			} else {
				/* bad number */
				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid range string - must be numeric")
				return nil, false
			}
		}
		step = zend.ZvalGetDouble(zstep)

		/* We only want positive step values. */
		if step < 0.0 {
			step *= -1
		}
	}

	/* If the range is given as strings, generate an array of characters. */
	if zlow.IsString() && zhigh.IsString() && zlow.String().GetLen() >= 1 && zhigh.String().GetLen() >= 1 {
		lowStr := zlow.StringVal()
		highStr := zhigh.StringVal()

		r1 := zend.StrToNumber(lowStr)
		r2 := zend.StrToNumber(highStr)
		if r1.IsFloat() || r2.IsFloat() || isStepDouble {
			return rangeDouble(zlow, zhigh, step)
		} else if r1.IsInt() || r2.IsInt() {
			return rangeLong(zlow, zhigh, int(step))
		} else {
			return rangeChar(lowStr[0], highStr[0], int(step))
		}
	} else if zlow.IsDouble() || zhigh.IsDouble() || isStepDouble {
		return rangeDouble(zlow, zhigh, step)
	} else {
		return rangeLong(zlow, zhigh, int(step))
	}
}
func arrayDataShuffle(array *types.Array) *types.Array {
	values := array.Values()
	for i := len(values) - 1; i >= 0; i-- {
		j := standard.PhpMtRandRange(0, i)
		if i != j {
			values[i], values[j] = values[j], values[i]
		}
	}
	return types.NewArrayOfZval(values)
}
func ZifShuffle(arg zpp.RefArray) bool {
	if arg.Array().Len() > 1 {
		arg.SetArray(arrayDataShuffle(arg.Array()))
	}
	return true
}

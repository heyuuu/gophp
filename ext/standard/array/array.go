package array

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
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
	if ht.IsRecursive() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
		return 0
	}
	ht.ProtectRecursive()

	cnt := ht.Count()
	ht.Foreach(func(key types.ArrayKey, value *types.Zval) {
		value = types.ZVAL_DEREF(value)
		if value.IsArray() {
			cnt += PhpCountRecursive(value.Array())
		}
	})

	ht.UnprotectRecursive()
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
		if array.Object().CanCountElements() {
			long = 1
			if types.SUCCESS == array.Object().CountElements(array, &long) {
				return long
			}
			if zend.EG__().GetException() != nil {
				return long
			}
		}

		/* if not and the object implements Countable we call its count() method */
		if operators.InstanceofFunction(types.Z_OBJCE_P(array), zend.ZendCeCountable) != 0 {
			zend.ZendCallMethodWith0Params(array, nil, nil, "count", &retval)
			if retval.IsNotUndef() {
				long = operators.ZvalGetLong(&retval)
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
func ZifEnd(array zpp.RefArrayHt) *types.Zval {
	array.MoveEnd()
	pair := array.Current()
	if pair == nil {
		return types.NewZvalBool(false)
	}
	return pair.GetVal().DeIndirect().DeRef()
}

func ZifPrev(array zpp.RefArrayHt) *types.Zval {
	array.MovePrev()
	pair := array.Current()
	if pair == nil {
		return types.NewZvalBool(false)
	}
	return pair.GetVal().DeIndirect().DeRef()
}
func ZifNext(array zpp.RefArrayHt) *types.Zval {
	array.MoveNext()
	pair := array.Current()
	if pair == nil {
		return types.NewZvalBool(false)
	}
	return pair.GetVal().DeIndirect().DeRef()
}

func ZifReset(array zpp.RefArrayHt) *types.Zval {
	array.ResetInternalPointer()
	pair := array.Current()
	if pair == nil {
		return types.NewZvalBool(false)
	}
	return pair.GetVal().DeIndirect().DeRef()
}

//@zif -alias pos
func ZifCurrent(array zpp.ArrayOrObjectHt) (*types.Zval, bool) {
	pair := array.Current()
	if pair == nil {
		return nil, false
	}

	return pair.GetVal().DeIndirect().DeRef(), true
}
func ZifKey(array zpp.ArrayOrObjectHt) *types.Zval {
	pair := array.Current()
	if pair == nil {
		return types.NewZvalNull()
	} else {
		return pair.GetKey().ToZval()
	}
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
				if operators.FastIsIdenticalFunction(value, entry) != 0 {
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
				if operators.FastEqualCheckLong(value, entry) {
					targetKey = &key
					return false
				}
				return true
			})
		} else if value.IsString() {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if operators.FastEqualCheckString(value, entry) {
					targetKey = &key
					return false
				}
				return true
			})
		} else {
			array.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
				entry = types.ZVAL_DEREF(entry)
				if operators.FastEqualCheckFunction(value, entry) {
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
			key := operators.ZvalGetStrVal(entry)
			arr.SymtableUpdate(key, val)
		}
	})
	return arr
}
func rangeDouble(zLow *types.Zval, zHigh *types.Zval, step float64) ([]*types.Zval, bool) {
	b.Assert(step > 0)
	low := operators.ZvalGetDouble(zLow)
	high := operators.ZvalGetDouble(zHigh)
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
	if size >= float64(types.MaxArraySize) {
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
	low := operators.ZvalGetLong(zLow)
	high := operators.ZvalGetLong(zHigh)
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
	if size >= types.MaxArraySize {
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
		step = operators.ZvalGetDouble(zstep)

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

	outHash := types.NewArray(0)
	removed := types.NewArray(0)
	pairs := arr.Pairs()

	// handle range [0, offset)
	for _, pair := range pairs[:offset] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	// handle range [offset, offset+length)
	for _, pair := range pairs[offset : offset+length] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			removed.Append(val)
		} else {
			removed.KeyAdd(key.StrKey(), val)
		}
	}

	// handle insert
	if replace != nil {
		replace.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			outHash.Append(value)
		})
	}

	// handle range [offset+length, len(array))
	for _, pair := range pairs[offset+length:] {
		key := pair.GetKey()
		val := pair.GetVal()
		if !key.IsStrKey() {
			outHash.Append(val)
		} else {
			outHash.KeyAdd(key.StrKey(), val)
		}
	}

	return outHash, removed
}
func ZifArrayPush(stack zpp.RefArray, _ zpp.Opt, args []*types.Zval) (int, bool) {
	for _, arg := range args {
		if stack.Array().Append(arg) == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return 0, false
		}
	}

	return stack.Array().Len(), true
}
func ZifArrayPop(stack zpp.RefArray) *types.Zval {
	pair := stack.Array().LastIndirect()
	if pair == nil {
		return types.NewZvalNull()
	}

	stack.Array().Delete(pair.GetKey())
	stack.Array().ResetInternalPointer()

	return pair.GetVal().DeRef()
}
func ZifArrayShift(stack zpp.RefArray) *types.Zval {
	if stack.Array().Len() == 0 {
		return nil
	}

	var p *types.Bucket

	/* Get the first value and copy it into the return value */
	pair := stack.Array().FirstIndirect()
	if pair == nil {
		return nil
	}

	key := pair.GetKey()
	val := pair.GetVal().DeRef()

	/* Delete the first value */
	if key.IsStrKey() && stack.Array() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		stack.Array().Delete(key)
	}

	/* re-index like it did before */
	var k = 0
	stack.Array().MapWithKey(func(key types.ArrayKey, value *types.Zval) (types.ArrayKey, *types.Zval) {
		if !key.IsStrKey() {
			key = types.IdxKey(k)
			k++
		}
		return key, value
	})

	return val
}
func ZifArrayUnshift(stack zpp.RefZval, values []*types.Zval) int {
	newArr := types.NewArray(stack.Array().Len() + len(values))
	for _, value := range values {
		newArr.Append(value)
	}
	stack.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			newArr.KeyAdd(key.StrKey(), value)
		} else {
			newArr.Append(value)
		}
	})

	stack.SetArray(newArr)
	return stack.Array().Len()
}
func ZifArraySplice(array zpp.RefArray, offset int, _ zpp.Opt, length_ *int, replacement *types.Zval) *types.Array {
	var numIn = array.Array().Len()
	var length = b.Option(length_, numIn)
	var replaceArr *types.Array = nil

	if replacement != nil {
		/* Make sure the last argument, if passed, is an array */
		operators.ConvertToArrayEx(replacement)
		replaceArr = replacement.Array()
	}

	newArr, removedArr := phpSplice(array.Array(), offset, length, replaceArr)
	array.SetArray(newArr)
	return removedArr
}
func ZifArraySlice(array *types.Array, offset int, _ zpp.Opt, length_ *types.Zval, preserveKeys bool) *types.Array {
	numIn := array.Len()

	/* We want all entries from offset to the end if length is not passed or is null */
	var length = 0
	if length_ == nil || length_.IsNull() {
		length = numIn
	} else {
		length = operators.ZvalGetLong(length_)
	}

	/* Clamp the offset.. */
	if offset > numIn {
		return types.NewArray(0)
	} else if offset < 0 && b.Assign(&offset, numIn+offset) < 0 {
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
		return types.NewArray(0)
	}

	/* Initialize returned array */
	retArr := types.NewArray(length)

	/* Start at the beginning and go until we hit offset */
	count := 0
	array.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
		count++
		if count <= offset {
			return true
		}
		if count > offset+length {
			return false
		}
		if key.IsStrKey() || preserveKeys {
			retArr.Add(key, value)
		} else {
			retArr.Append(value)
		}
		return true
	})
	return retArr
}
func PhpArrayMergeRecursive(dest *types.Array, src *types.Array) int {
	for iter := src.Iterator(); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Current()
		if !key.IsStrKey() {
			dest.Append(value)
			continue
		}

		strKey := key.StrKey()
		destEntry := dest.KeyFind(strKey)
		if destEntry == nil {
			dest.KeyAddNew(strKey, value)
		}

		var srcZval = value.DeRef()
		var destZval = destEntry.DeRef()
		var thash *types.Array
		var tmp types.Zval
		var ret int
		if destZval.IsType(types.IS_ARRAY) {
			thash = destZval.Array()
		} else {
			thash = nil
		}
		if thash != nil && thash.IsRecursive() || value == destEntry && destEntry.IsReference() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		//b.Assert(!(destEntry.IsReference()) || destEntry.GetRefcount() > 1)
		types.SeparateZval(destEntry)
		destZval = destEntry
		if destZval.IsType(types.IS_NULL) {
			operators.ConvertToArrayEx(destZval)
			zend.AddNextIndexNull(destZval)
		} else {
			operators.ConvertToArrayEx(destZval)
		}
		tmp.SetUndef()
		if srcZval.IsType(types.IS_OBJECT) {
			types.ZVAL_COPY(&tmp, srcZval)
			operators.ConvertToArray(&tmp)
			srcZval = &tmp
		}
		if srcZval.IsType(types.IS_ARRAY) {
			if thash != nil {
				thash.ProtectRecursive()
			}
			ret = PhpArrayMergeRecursive(destZval.Array(), srcZval.Array())
			if thash != nil {
				thash.UnprotectRecursive()
			}
			if ret == 0 {
				return 0
			}
		} else {
			destZval.Array().Append(srcZval)
		}
	}
	return 1
}
func PhpArrayMerge(dest *types.Array, src *types.Array) {
	src.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			dest.KeyUpdate(key.StrKey(), value)
		} else {
			dest.Append(value)
		}
	})
}
func PhpArrayReplaceRecursive(dest *types.Array, src *types.Array) bool {
	var destEntry *types.Zval
	var srcZval *types.Zval
	var destZval *types.Zval
	var ret bool
	return src.ForeachEx(func(key types.ArrayKey, srcEntry *types.Zval) bool {
		srcZval = srcEntry.DeRef()
		if !srcZval.IsArray() {
			dest.Update(key, srcEntry)
			return true
		}

		destEntry = dest.Find(key)
		if destEntry == nil || !destEntry.IsArray() && (!(destEntry.IsReference()) || types.Z_REFVAL_P(destEntry).GetType() != types.IS_ARRAY) {
			dest.Update(key, srcEntry)
			return true
		}
		destZval = destEntry.DeRef()

		// src/dest 对应值均为 array 的情况下，递归替换
		if destZval.Array().IsRecursive() || srcZval.Array().IsRecursive() || srcEntry.IsReference() && destEntry.IsReference() && srcEntry.Reference() == destEntry.Reference() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return false
		}
		types.SeparateZval(destEntry)
		destZval = destEntry

		b.Assert(srcZval.IsArray() && destZval.IsArray())
		srcZval.Array().ProtectRecursive()
		destZval.Array().ProtectRecursive()
		ret = PhpArrayReplaceRecursive(destZval.Array(), srcZval.Array())
		srcZval.Array().UnprotectRecursive()
		destZval.Array().UnprotectRecursive()

		if !ret {
			return false
		}
		return true
	})
}
func arrayMergeWrapper(args []*types.Zval, recursive bool) *types.Array {
	var dest *types.Array
	if len(args) == 0 {
		return types.NewArray(0)
	}

	count := 0
	for i, arg := range args {
		if !arg.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(arg))
			return nil
		}
		count += arg.Array().Len()
	}

	arr := types.NewArray(count)

	args[0].Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			arr.KeyUpdate(key.StrKey(), value)
		} else {
			arr.Append(value)
		}
	})

	if recursive {
		for _, arg := range args[1:] {
			PhpArrayMergeRecursive(dest, arg.Array())
		}
	} else {
		for _, arg := range args[1:] {
			PhpArrayMerge(dest, arg.Array())
		}
	}
	return arr
}
func ZifArrayMerge(returnValue zpp.Ret, _ zpp.Opt, arrays []*types.Zval) {
	arr := arrayMergeWrapper(arrays, false)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}
func ZifArrayMergeRecursive(returnValue zpp.Ret, _ zpp.Opt, arrays []*types.Zval) {
	arr := arrayMergeWrapper(arrays, true)
	if arr == nil {
		returnValue.SetNull()
	} else {
		returnValue.SetArray(arr)
	}
}

//@zif -c=1,
func ZifArrayReplace(arrays []*types.Zval) *types.Zval {
	b.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(arrays, 0)
	if !ok {
		return types.NewZvalNull()
	}

	dest := arrayHts[0].Copy()
	for _, array := range arrayHts[1:] {
		types.ZendHashMerge(dest, array, true)
	}
	return types.NewZvalArray(dest)
}

//@zif -c=1,
func ZifArrayReplaceRecursive(arrays []*types.Zval) *types.Zval {
	b.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(arrays, 0)
	if !ok {
		return types.NewZvalNull()
	}

	dest := arrayHts[0].Copy()
	for _, array := range arrayHts[1:] {
		PhpArrayReplaceRecursive(dest, array)
	}
	return types.NewZvalArray(dest)
}
func ZifArrayKeys(array *types.Array, _ zpp.Opt, searchValue *types.Zval, strict bool) *types.Array {
	arrLen := array.Len()

	/* Base case: empty input */
	if arrLen == 0 {
		return types.NewArray(0)
	}

	/* Initialize return array */
	if searchValue != nil {
		keys := types.NewArray(0)
		if strict {
			array.ForeachIndirect(func(key types.ArrayKey, entry *types.Zval) {
				entry = types.ZVAL_DEREF(entry)
				if operators.FastIsIdenticalFunction(searchValue, entry) != 0 {
					keys.Append(key.ToZval())
				}
			})
		} else {
			array.ForeachIndirect(func(key types.ArrayKey, entry *types.Zval) {
				if operators.FastEqualCheckFunction(searchValue, entry) {
					keys.Append(key.ToZval())
				}
			})
		}
		return keys
	} else {
		keys := types.NewArray(arrLen)
		array.ForeachIndirect(func(key types.ArrayKey, _ *types.Zval) {
			keys.Append(key.ToZval())
		})
		return keys
	}
}
func ZifArrayKeyFirst(array *types.Array) *types.Zval {
	pair := array.FirstIndirect()
	if pair == nil {
		return types.NewZvalNull()
	}
	return pair.GetKey().ToZval()
}
func ZifArrayKeyLast(array *types.Array) *types.Zval {
	pair := array.Last()
	if pair == nil {
		return types.NewZvalNull()
	}
	return pair.GetKey().ToZval()
}
func ZifArrayValues(array *types.Array) *types.Array {
	arrLen := array.Len()
	if arrLen == 0 {
		return types.NewArray(0)
	}

	/* Initialize return array */
	values := types.NewArray(array.Len())
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		values.Append(entry)
	})
	return values
}
func ZifArrayCountValues(array *types.Array) *types.Array {
	retArr := types.NewArray(0)
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		entry = types.ZVAL_DEREF(entry)

		var key types.ArrayKey
		if entry.IsLong() {
			key = types.IdxKey(entry.Long())
		} else if entry.IsString() {
			key = types.NumericKey(entry.StringVal())
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only count STRING and INTEGER values!")
			return
		}

		if keyCount := retArr.Find(key); keyCount != nil {
			retArr.Update(key, types.NewZvalLong(keyCount.Long()+1))
		} else {
			retArr.Update(key, types.NewZvalLong(1))
		}
	})

	return retArr
}
func ArrayColumnParamHelper(param *types.Zval, name string) types.ZendBool {
	switch param.GetType() {
	case types.IS_DOUBLE:
		if param.GetType() != types.IS_LONG {
			operators.ConvertToLong(param)
		}
		fallthrough
	case types.IS_LONG:
		return 1
	case types.IS_OBJECT:
		if operators.TryConvertToString(param) == 0 {
			return 0
		}
		fallthrough
	case types.IS_STRING:
		return 1
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The %s key should be either a string or an integer", name)
		return 0
	}
}
func ArrayColumnFetchProp(data *types.Zval, name *types.Zval) *types.Zval {
	var rv types.Zval
	var prop *types.Zval = nil
	if data.IsType(types.IS_OBJECT) {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if data.Object().HasProperty(data, name, zend.ZEND_PROPERTY_EXISTS, nil) != 0 || data.Object().HasProperty(data, name, zend.ZEND_PROPERTY_ISSET, nil) != 0 {
			prop = data.Object().ReadProperty(data, name, zend.BP_VAR_R, nil, &rv)
			if prop != nil {
				prop = types.ZVAL_DEREF(prop)
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.IsType(types.IS_ARRAY) {
		if name.IsString() {
			prop = data.Array().SymtableFind(name.String().GetStr())
		} else if name.IsType(types.IS_LONG) {
			prop = data.Array().IndexFind(name.Long())
		}
		if prop != nil {
			prop = types.ZVAL_DEREF(prop)
		}
	}
	return prop
}
func ZifArrayColumn(array *types.Array, columnKey zpp.ZvalNullable, _ zpp.Opt, indexKey zpp.ZvalNullable) (*types.Array, bool) {
	if columnKey != nil && ArrayColumnParamHelper(columnKey, "column") == 0 || indexKey != nil && ArrayColumnParamHelper(indexKey, "index") == 0 {
		return nil, false
	}

	retArr := types.NewArray(array.Len())
	if indexKey == nil {
		array.Foreach(func(_ types.ArrayKey, data *types.Zval) {
			var columnVal *types.Zval
			data = types.ZVAL_DEREF(data)
			if columnKey == nil {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(data, columnKey); columnVal == nil {
				return
			}

			retArr.Append(columnVal)
		})
	} else {
		array.Foreach(func(key types.ArrayKey, data *types.Zval) {
			data = types.ZVAL_DEREF(data)

			// col
			var columnVal *types.Zval
			if columnKey == nil {
				columnVal = data
			} else if columnVal = ArrayColumnFetchProp(data, columnKey); columnVal == nil {
				return
			}

			// key
			var keyVal = ArrayColumnFetchProp(data, indexKey)
			if keyVal != nil {
				switch keyVal.GetType() {
				case types.IS_STRING:
					retArr.SymtableUpdate(keyVal.String().GetStr(), columnVal)
				case types.IS_LONG:
					retArr.IndexUpdate(keyVal.Long(), columnVal)
				case types.IS_OBJECT:
					retArr.SymtableUpdate(operators.ZvalGetStrVal(keyVal), columnVal)
				case types.IS_NULL:
					retArr.KeyUpdate("", columnVal)
				case types.IS_DOUBLE:
					retArr.IndexUpdate(operators.DvalToLval(keyVal.Double()), columnVal)
				case types.IS_TRUE:
					retArr.IndexUpdate(1, columnVal)
				case types.IS_FALSE:
					retArr.IndexUpdate(0, columnVal)
				case types.IS_RESOURCE:
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
	retArr := types.NewArray(array.Len())
	array.ForeachReserve(func(key types.ArrayKey, value *types.Zval) {
		if preserveKeys || key.IsStrKey() {
			retArr.Add(key, value)
		} else { // !preserveKeys && !key.IsStrKey()
			retArr.Append(value)
		}
	})
	return retArr
}

const maxPadInOneTimes = 1048576

func ZifArrayPad(array *types.Array, padSize int, padValue *types.Zval) (*types.Array, bool) {
	/* Do some initial calculations */
	inputSize := array.Len()
	padSizeAbs := padSize
	if padSizeAbs < 0 {
		padSizeAbs = -padSizeAbs
	}
	if padSizeAbs < 0 || padSizeAbs-inputSize > maxPadInOneTimes {
		core.PhpErrorDocref(nil, faults.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return nil, false
	}

	/**
	 * notice: 注意此处的差别:
	 * - 当原数组的长度 >= padSize 时，返回数组保留了原数组包括内部指针等信息；
	 * - 当原数组的长度 < padSize 时，产生的新数组内部指针等信息未设置(等于初始值)；
	 */
	if inputSize >= padSizeAbs {
		return array.Copy(), true
	}

	numPads := padSizeAbs - inputSize
	retArr := types.NewArray(padSizeAbs)
	if padSize < 0 {
		for i := 0; i < numPads; i++ {
			retArr.Append(padValue)
		}
	}
	array.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
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
func ZifArrayFlip(array *types.Array) *types.Array {
	retArr := types.NewArray(array.Len())
	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
		value = value.DeRef()
		if value.IsLong() {
			retArr.IndexUpdate(value.Long(), key.ToZval())
		} else if value.IsString() {
			retArr.SymtableUpdate(value.StringVal(), key.ToZval())
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	})
	return retArr
}
func ZifArrayChangeKeyCase(array *types.Array, _ zpp.Opt, case_ *int) *types.Array {
	var caseFlag = b.Option(case_, CASE_LOWER)
	retArr := types.NewArray(array.Len())
	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
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
func ZifArrayUnique(arg *types.Array, _ zpp.Opt, flags *int) *types.Array {
	var sortType = b.Option(flags, PHP_SORT_STRING)

	if sortType == PHP_SORT_STRING {
		existValues := make(map[string]bool, arg.Len())
		retArr := types.NewArray(arg.Len())
		arg.ForeachIndirect(func(key types.ArrayKey, val *types.Zval) {
			var strVal = operators.ZvalGetStrVal(val)
			if _, exists := existValues[strVal]; !exists {
				retArr.Add(key, val)
			}

		})
		return retArr
	}

	cmp := phpGetDataCompareFunc(sortType, false)
	retArr := types.ZendArrayDup(arg)

	// 取出数组所有键值对，先排序，后找出重复元素并删除
	pairs := arg.Pairs()
	sort.SliceStable(pairs, func(i, j int) bool {
		return cmp(pairs[i], pairs[j]) < 0
	})

	last := 0
	for i := 1; i < len(pairs); i++ {
		if cmp(pairs[last], pairs[i]) != 0 {
			last = i
		} else {
			pair := pairs[i]
			retArr.Delete(pair.GetKey())
		}
	}

	return retArr
}
func ZvalCompare(first *types.Zval, second *types.Zval) int {
	return operators.StringCompareFunction(first, second)
}

//@zif -c 2,
func ZifArrayIntersectKey(arrays []*types.Zval) (*types.Array, bool) {
	return arrayIntersectKeyWrapper(arrays, nil)
}

//@zif -c=3,
func ZifArrayIntersectUkey(arrays []*types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(callbackKeyCompareFunc)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayIntersect(arrays []*types.Zval) (*types.Array, bool) {
	cmp := arrayDataComparer(operators.StringCompareFunction)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayUintersect(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserDataComparer(callbackDataCompareFunc)
	return arrayIntersectWrapper(arrays, cmp)
}
func ZifArrayIntersectAssoc(arrays []*types.Zval) (*types.Array, bool) {
	cmp := ZvalCompare
	return arrayIntersectKeyWrapper(arrays, cmp)
}
func ZifArrayIntersectUassoc(arrays []*types.Zval, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompareFunc),
		arrayDataComparer(operators.StringCompareFunction),
	)
	return arrayIntersectWrapper(arrays, cmp)
}
func ZifArrayUintersectAssoc(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserZvalComparer(callbackDataCompareFunc)
	return arrayIntersectKeyWrapper(arrays, cmp)
}

//@zif -c=4,
func ZifArrayUintersectUassoc(arrays []*types.Zval, callbackDataCompareFunc zpp.Callable, callbackKeyCompareFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompareFunc),
		arrayUserDataComparer(callbackDataCompareFunc),
	)
	return arrayIntersectWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayDiffKey(arrays []*types.Zval) (*types.Array, bool) {
	return arrayDiffKeyWrapper(arrays, nil)
}

//@zif -c=3,
func ZifArrayDiffUkey(arrays []*types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserKeyComparer(callbackKeyCompFunc)
	return arrayDiffWrapper(arrays, cmp)
}

func simpleArrayDiff(array *types.Array, arrays []*types.Array) *types.Array {
	if array.Len() == 0 {
		return types.NewArray(0)
	}

	// array.Len() > 1
	exclude := make(map[string]bool)
	for _, diffArray := range arrays {
		diffArray.ForeachIndirect(func(_ types.ArrayKey, value *types.Zval) {
			str := operators.ZvalGetStrVal(value)
			exclude[str] = true
		})
	}
	if len(exclude) == 0 {
		return array
	}

	retArr := types.NewArray(array.Len())
	array.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		str := operators.ZvalGetStrVal(value)
		if _, excluded := exclude[str]; !excluded {
			retArr.Add(key, value)
		}
	})
	return retArr
}

//@zif -c=2,
func ZifArrayDiff(arrays []*types.Zval) *types.Zval {
	arrayHts, ok := checkArrayArgs(arrays, 0)
	if !ok {
		return types.NewZvalNull()
	}
	retArr := simpleArrayDiff(arrayHts[0], arrayHts[1:])
	return types.NewZvalArray(retArr)
}

//@zif -c=3,
func ZifArrayUdiff(arrays []*types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserDataComparer(callbackDataCompFunc)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayDiffAssoc(arrays []*types.Zval) (*types.Array, bool) {
	cmp := ZvalCompare
	return arrayDiffKeyWrapper(arrays, cmp)
}

//@zif -c=3,
func ZifArrayDiffUassoc(arrays []*types.Zval, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompFunc),
		arrayDataComparer(operators.StringCompareFunction),
	)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=2,
func ZifArrayUdiffAssoc(arrays []*types.Zval, callbackDataCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := arrayUserZvalComparer(callbackDataCompFunc)
	return arrayDiffKeyWrapper(arrays, cmp)
}

//@zif -c=4,
func ZifArrayUdiffUassoc(arrays []*types.Zval, callbackDataCompFunc zpp.Callable, callbackKeyCompFunc zpp.Callable) (*types.Array, bool) {
	cmp := twiceComparer(
		arrayUserKeyComparer(callbackKeyCompFunc),
		arrayUserDataComparer(callbackDataCompFunc),
	)
	return arrayDiffWrapper(arrays, cmp)
}

//@zif -c=1,
func ZifArrayMultisort(args []*types.Zval) bool {
	var parseState = [...]int{0, 0}

	arrays := make([]*types.Zval, 0, len(args))
	multisortFunc := make([]types.ArrayComparer, 0, len(args))

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */
	var sortOrder = PHP_SORT_ASC
	var sortType = PHP_SORT_REGULAR
	for i, arg := range args {
		arg = types.ZVAL_DEREF(arg)
		if arg.IsArray() {
			types.SeparateArray(arg)

			/* We see the next array, so we update the sort flags of
			 * the previous array and reset the sort flags. */
			if i > 0 {
				multisortFunc = append(multisortFunc, phpGetDataCompareFunc(sortType, sortOrder != PHP_SORT_ASC))
				sortOrder = PHP_SORT_ASC
				sortType = PHP_SORT_REGULAR
			}
			arrays = append(arrays, arg)

			/* Next one may be an array or a list of sort flags. */
			for k := 0; k < MULTISORT_LAST; k++ {
				parseState[k] = 1
			}

			/* Next one may be an array or a list of sort flags. */

		} else if arg.IsType(types.IS_LONG) {
			switch arg.Long() & ^PHP_SORT_FLAG_CASE {
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
					core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
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
					core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					return false
				}
			default:
				core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is an unknown sort flag", i+1)
				return false
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Argument #%d is expected to be an array or a sort flag", i+1)
			return false
		}
	}

	/* Take care of the last array sort flags. */
	multisortFunc = append(multisortFunc, phpGetDataCompareFunc(sortType, sortOrder != PHP_SORT_ASC))

	/* Make sure the arrays are of the same size. */
	arraySize := arrays[0].Array().Len()
	for _, array := range arrays {
		if array.Array().Len() != arraySize {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Array sizes are inconsistent")
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
		array.Array().ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			matrix[idx] = append(matrix[idx], types.MakeArrayPair(key, value))
		})
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */
	sort.SliceStable(matrix, func(idx1, idx2 int) bool {
		col1 := matrix[idx1]
		col2 := matrix[idx2]
		for i, sortFunc := range multisortFunc {
			result := sortFunc(col1[i], col2[i])
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
		newHash := types.NewArray(array.Array().Len())
		for idx := 0; idx < arraySize; idx++ {
			pair := matrix[idx][i]
			if pair.GetKey().IsStrKey() {
				newHash.KeyAdd(pair.GetKey().StrKey(), pair.GetVal())
			} else {
				newHash.Append(pair.GetVal())
			}
		}
	}

	return true
}
func ZifArrayRand(arg *types.Array, _ zpp.Opt, numReq_ *int) *types.Zval {
	numReq := b.Option(numReq_, 1)

	numAvail := arg.Len()
	if numAvail == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Array is empty")
		return nil
	}
	if numReq <= 0 || numReq > numAvail {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Second argument has to be between 1 and the number of elements in the array")
		return nil
	}

	keys := arg.Keys()
	if numReq == 1 {
		randIdx := standard.PhpMtRandRange(0, numAvail-1)
		return keys[randIdx].ToZval()
	} else {
		randIdxSet := make(map[int]bool)
		negative := false
		if numReq > arg.Len()/2 {
			negative = true
			numReq = arg.Len() - numReq
		}
		for i := numReq - 1; i >= 0; {
			randIdx := standard.PhpMtRandRange(0, numAvail-1)
			if !randIdxSet[randIdx] {
				randIdxSet[randIdx] = true
				i--
			}
		}

		retArr := types.NewArray(numReq)
		for i, key := range keys {
			if !negative && randIdxSet[i] || negative && !randIdxSet[i] {
				retArr.Append(key.ToZval())
			}
		}
		return types.NewZvalArray(retArr)
	}
}
func ZifArraySum(array *types.Array) *types.Zval {
	var num types.Zval
	ret := types.NewZvalLong(0)
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		types.ZVAL_COPY(&num, entry)
		operators.ConvertScalarToNumber(&num)
		operators.FastAddFunction(ret, ret, &num)
	})
	return ret
}
func ZifArrayProduct(array *types.Array) *types.Zval {
	var num types.Zval
	ret := types.NewZvalLong(1)
	array.Foreach(func(_ types.ArrayKey, entry *types.Zval) {
		if entry.IsArray() || entry.IsObject() {
			return
		}
		types.ZVAL_COPY(&num, entry)
		operators.ConvertScalarToNumber(&num)
		if num.IsLong() && ret.IsLong() {
			dval := float64(num.Long()) * float64(ret.Long())
			if float64(zend.ZEND_LONG_MIN) <= dval && dval <= float64(zend.ZEND_LONG_MAX) {
				ret.SetLong(ret.Long() * num.Long())
				return
			}
		}
		operators.ConvertToDouble(ret)
		operators.ConvertToDouble(&num)
		ret.SetDouble(ret.Double() * num.Double())
	})
	return ret
}
func ZifArrayReduce(array *types.Array, callback zpp.Callable, _ zpp.Opt, initial *types.Zval) *types.Zval {
	var result types.Zval
	if initial == nil {
		result.SetNull()
	} else {
		result.CopyFrom(initial)
	}

	array.Foreach(func(key types.ArrayKey, value *types.Zval) {
		newResult, ok := callback.Call(&result, value)
		if ok && newResult.IsNotUndef() {
			result.CopyValueFrom(newResult)
		}
	})

	return zend.ZvalZval(&result, true, true)
}
func ZifArrayFilter(array_ *types.Array, _ zpp.Opt, callback zpp.Callable, mode int) *types.Array {
	retArr := types.NewArray(0)
	if array_.Len() == 0 {
		return retArr
	}

	// parse use filter
	var filter func(key types.ArrayKey, value *types.Zval) (keep bool, succ bool)
	if callback != nil {
		filter = func(key types.ArrayKey, value *types.Zval) (bool, bool) {
			var retVal *types.Zval
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
			return operators.ZvalIsTrue(retVal), true
		}
	}

	// filter
	array_.ForeachIndirectEx(func(key types.ArrayKey, value *types.Zval) bool {
		var keep bool
		if filter != nil {
			var callSucc bool
			keep, callSucc = filter(key, value)
			if !callSucc {
				// 调用 user callback 失败，中断遍历
				return false
			}
		} else {
			keep = operators.ZvalIsTrue(value)
		}

		if keep {
			retArr.Update(key, value)
		}
		return true
	})
	return retArr
}

func arrayMapSingle(callback zpp.Callable, array *types.Array) *types.Array {
	retArr := types.NewArray(array.Len())
	ok := array.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
		retVal, ok := callback.Call(value)
		if !ok || retVal.IsUndef() {
			// 调用 callback 失败，中断流程
			return false
		}

		retArr.Add(key, retVal)
		return true
	})
	if !ok {
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

	argMatrix := make([][]*types.Zval, len_)
	for _, array := range arrays {
		count := 0
		array.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
			argMatrix[count] = append(argMatrix[count], value)
			count++
			return count < len_
		})
		for i := len_; i < count; i++ {
			argMatrix[i] = append(argMatrix[i], types.NewZvalNull())
		}
	}

	retArr := types.NewArray(len_)
	for _, argColumns := range argMatrix {
		if callback == nil {
			retArr.Append(types.NewZvalArray(types.NewArrayOfZval(argColumns)))
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

//@zif -c=2,
func ZifArrayMap(callback zpp.Callable, arrays []*types.Zval) *types.Zval {
	b.Assert(len(arrays) >= 1)

	arrayHts, ok := checkArrayArgs(arrays, 1)
	if !ok {
		return types.NewZvalNull()
	}

	var retArr *types.Array
	if len(arrayHts) == 1 {
		retArr = arrayMapSingle(callback, arrayHts[0])
	} else {
		retArr = arrayMapMulti(callback, arrayHts)
	}

	if retArr == nil {
		return types.NewZvalNull()
	}
	return types.NewZvalArray(retArr)
}

//@zif -alias key_exists
func ZifArrayKeyExists(key *types.Zval, array zpp.ArrayOrObject) bool {
	var ht *types.Array
	if array.IsType(types.IS_ARRAY) {
		ht = array.Array()
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. Use isset() or property_exists() instead")
	}
	switch key.GetType() {
	case types.IS_STRING:
		return ht.SymtableExistsInd(key.StringVal())
	case types.IS_LONG:
		return ht.IndexExists(key.Long())
	case types.IS_NULL:
		return ht.KeyExistsIndirect("")
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The first argument should be either a string or an integer")
		return false
	}
}
func ZifArrayChunk(array *types.Array, length int, _ zpp.Opt, preserveKeys bool) *types.Array {
	if length < 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Size parameter expected to be greater than 0")
		return nil
	}

	chunkCount := (array.Len()-1)/length + 1 // ceil(len/length)
	retArr := types.NewArray(chunkCount)

	var currChunk *types.Array = nil
	itemCount := 0
	array.Foreach(func(key types.ArrayKey, entry *types.Zval) {
		if itemCount%length == 0 {
			currChunk = types.NewArray(length)
			retArr.Append(types.NewZvalArray(currChunk))
		}

		if preserveKeys {
			currChunk.Update(key, entry)
		} else {
			currChunk.Append(entry)
		}
		itemCount++
	})
	return retArr
}
func ZifArrayCombine(keys *types.Array, values *types.Array) (*types.Array, bool) {
	if keys.Len() != values.Len() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Both parameters should have an equal number of elements")
		return nil, false
	}
	if keys.Len() == 0 {
		return types.NewArray(0), true
	}

	zvKeys := keys.Values()
	zvValues := values.Values()
	retArr := types.NewArray(keys.Len())
	for i, key := range zvKeys {
		if i > len(zvValues) {
			break
		}
		value := zvValues[i]

		if key.IsLong() {
			retArr.IndexUpdate(key.Long(), value)
		} else {
			strKey := operators.ZvalGetStrVal(key)
			retArr.KeyUpdate(strKey, value)
		}
	}
	return retArr, true
}

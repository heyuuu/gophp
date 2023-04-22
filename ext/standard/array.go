package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
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
 * Bucketindex
 */
type Bucketindex struct {
	b types.Bucket
	i uint
}

func (this *Bucketindex) GetB() types.Bucket      { return this.b }
func (this *Bucketindex) SetB(value types.Bucket) { this.b = value }
func (this *Bucketindex) GetI() uint              { return this.i }
func (this *Bucketindex) SetI(value uint)         { this.i = value }

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
func ZmShutdownArray(type_ int, module_number int) int { return types.SUCCESS }

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
func ZifCount(return_value zpp.Ret, var_ *types.Zval, _ zpp.Opt, mode int) int {
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

		/* first, we check if the handler is defined */
		if types.Z_OBJ_HT_P(array).GetCountElements() != nil {
			return_value.SetLong(1)
			if types.SUCCESS == types.Z_OBJ_HT(*array).GetCountElements()(array, &(return_value.Long())) {
				return
			}
			if zend.EG__().GetException() != nil {
				return
			}
		}

		/* if not and the object implements Countable we call its count() method */
		if zend.InstanceofFunction(types.Z_OBJCE_P(array), zend.ZendCeCountable) != 0 {
			zend.ZendCallMethodWith0Params(array, nil, nil, "count", &retval)
			if retval.IsNotUndef() {
				return_value.SetLong(zend.ZvalGetLong(&retval))
			}
			return
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
func ZifPrev(executeData zpp.Ex, return_value zpp.Ret, arg zpp.RefZval) {
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
	types.ZendHashMoveBackwards(array)
	if zend.USED_RET() {
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
func ZifNext(executeData zpp.Ex, return_value zpp.Ret, arg zpp.RefZval) {
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
	types.ZendHashMoveForward(array)
	if zend.USED_RET() {
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
func ZifReset(executeData zpp.Ex, return_value zpp.Ret, arg zpp.RefZval) {
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
	types.ZendHashInternalPointerReset(array)
	if zend.USED_RET() {
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

//@zif -alias pos
func ZifCurrent(arg zpp.ArrayOrObjectHt) (*types.Zval, bool) {
	_, val, ok := arg.Current(false)
	if !ok {
		return nil, false
	}

	if val.IsIndirect() {
		val = val.Indirect()
	}
	return val.DeRef(), true
}
func ZifKey(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var array *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			array = fp.ParseArrayOrObjectHt()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
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
func PhpArrayWalk(array *types.Zval, userdata *types.Zval, recursive int) int {
	var args []types.Zval
	var retval types.Zval
	var zv *types.Zval
	var target_hash = zend.HASH_OF(array)
	var pos types.ArrayPosition
	var ht_iter uint32
	var result = types.SUCCESS

	/* Set up known arguments */

	args[1].SetUndef()
	if userdata != nil {
		types.ZVAL_COPY(&args[2], userdata)
	}
	BG__().array_walk_fci.retval = &retval
	if userdata != nil {
		BG__().array_walk_fci.param_count = 3
	} else {
		BG__().array_walk_fci.param_count = 2
	}
	BG__().array_walk_fci.params = args
	BG__().array_walk_fci.no_separation = 0
	ht_iter = zend.EG__().AddArrayIterator(target_hash)

	/* Iterate through hash */

	for {

		/* Retrieve value */

		zv = types.ZendHashGetCurrentDataEx(target_hash, &pos)
		if zv == nil {
			break
		}

		/* Skip undefined indirect elements */

		if zv.IsIndirect() {
			zv = zv.Indirect()
			if zv.IsUndef() {
				types.ZendHashMoveForwardEx(target_hash, &pos)
				continue
			}

			/* Add type source for property references. */

			if zv.GetType() != types.IS_REFERENCE && array.IsType(types.IS_OBJECT) {
				var prop_info = zend.ZendGetTypedPropertyInfoForSlot(array.Object(), zv)
				if prop_info != nil {
					zv.SetNewRef(zv)
					zend.ZEND_REF_ADD_TYPE_SOURCE(zv.Reference(), prop_info)
				}
			}

			/* Add type source for property references. */

		}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */

		types.ZVAL_MAKE_REF(zv)

		/* Retrieve key */

		types.ZendHashGetCurrentKeyZvalEx(target_hash, &args[1], &pos)

		/* Move to next element already now -- this mirrors the approach used by foreach
		 * and ensures proper behavior with regard to modifications. */

		types.ZendHashMoveForwardEx(target_hash, &pos)

		/* Back up hash position, as it may change */

		zend.EG__().ArrayIterators()[ht_iter].SetPos(pos)
		if recursive != 0 && types.Z_REFVAL_P(zv).IsType(types.IS_ARRAY) {
			var thash *types.Array
			var orig_array_walk_fci types.ZendFcallInfo
			var orig_array_walk_fci_cache types.ZendFcallInfoCache
			var ref types.Zval
			types.ZVAL_COPY_VALUE(&ref, zv)
			zv = types.ZVAL_DEREF(zv)
			types.SeparateArray(zv)
			thash = zv.Array()
			if thash.IsRecursive() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
				result = types.FAILURE
				break
			}

			/* backup the fcall info and cache */

			orig_array_walk_fci = BG__().array_walk_fci
			orig_array_walk_fci_cache = BG__().array_walk_fci_cache
			// 			ref.AddRefcount()
			thash.ProtectRecursive()
			result = PhpArrayWalk(zv, userdata, recursive)
			if types.Z_REFVAL(ref).IsType(types.IS_ARRAY) && thash == types.Z_REFVAL(ref).Array() {

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

				thash.UnprotectRecursive()

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

			}
			// zend.ZvalPtrDtor(&ref)

			/* restore the fcall info and cache */

			BG__().array_walk_fci = orig_array_walk_fci
			BG__().array_walk_fci_cache = orig_array_walk_fci_cache
		} else {
			types.ZVAL_COPY(&args[0], zv)

			/* Call the userland function */

			result = zend.ZendCallFunction(&(BG__().array_walk_fci), &(BG__().array_walk_fci_cache))
			if result == types.SUCCESS {
				// zend.ZvalPtrDtor(&retval)
			}
			// zend.ZvalPtrDtor(&args[0])
		}
		if args[1].IsNotUndef() {
			// zend.ZvalPtrDtor(&args[1])
			args[1].SetUndef()
		}
		if result == types.FAILURE {
			break
		}

		/* Reload array and position -- both may have changed */

		if array.IsType(types.IS_ARRAY) {
			pos = types.ZendHashIteratorPos(ht_iter, array.Array())
			target_hash = array.Array()
		} else if array.IsType(types.IS_OBJECT) {
			target_hash = types.Z_OBJPROP_P(array)
			pos = types.ZendHashIteratorPos(ht_iter, target_hash)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Iterated value is no longer an array or object")
			result = types.FAILURE
			break
		}

		/* Reload array and position -- both may have changed */

		if zend.EG__().GetException() != nil {
			break
		}
	}
	zend.EG__().DelArrayIterator(ht_iter)
	return result
}
func ZifArrayWalk(executeData zpp.Ex, return_value zpp.Ret, input zpp.RefZval, funcname *types.Zval, _ zpp.Opt, userdata *types.Zval) {
	var array *types.Zval
	var userdata *types.Zval = nil
	var orig_array_walk_fci types.ZendFcallInfo
	var orig_array_walk_fci_cache types.ZendFcallInfoCache
	orig_array_walk_fci = BG__().array_walk_fci
	orig_array_walk_fci_cache = BG__().array_walk_fci_cache
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			array = fp.ParseArrayOrObjectEx(false, true)
			fp.ParseFunc(&(BG__().array_walk_fci), &(BG__().array_walk_fci_cache))
			fp.StartOptional()
			userdata = fp.ParseZval()
			break
		}
		if _error_code != zpp.ZPP_ERROR_OK {
			BG__().array_walk_fci = orig_array_walk_fci
			BG__().array_walk_fci_cache = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 0)
	zend.ZendReleaseFcallInfoCache(&(BG__().array_walk_fci_cache))
	BG__().array_walk_fci = orig_array_walk_fci
	BG__().array_walk_fci_cache = orig_array_walk_fci_cache
	return_value.SetTrue()
	return
}
func ZifArrayWalkRecursive(executeData zpp.Ex, return_value zpp.Ret, input zpp.RefZval, funcname *types.Zval, _ zpp.Opt, userdata *types.Zval) {
	var array *types.Zval
	var userdata *types.Zval = nil
	var orig_array_walk_fci types.ZendFcallInfo
	var orig_array_walk_fci_cache types.ZendFcallInfoCache
	orig_array_walk_fci = BG__().array_walk_fci
	orig_array_walk_fci_cache = BG__().array_walk_fci_cache
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			array = fp.ParseArrayOrObjectEx(false, true)
			fp.ParseFunc(&(BG__().array_walk_fci), &(BG__().array_walk_fci_cache))
			fp.StartOptional()
			userdata = fp.ParseZval()
			break
		}
		if _error_code != zpp.ZPP_ERROR_OK {
			BG__().array_walk_fci = orig_array_walk_fci
			BG__().array_walk_fci_cache = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 1)
	zend.ZendReleaseFcallInfoCache(&(BG__().array_walk_fci_cache))
	BG__().array_walk_fci = orig_array_walk_fci
	BG__().array_walk_fci_cache = orig_array_walk_fci_cache
	return_value.SetTrue()
	return
}
func PhpSearchArray(executeData *zend.ZendExecuteData, return_value *types.Zval, behavior int) {
	var value *types.Zval
	var array *types.Zval
	var entry *types.Zval
	var num_idx zend.ZendUlong
	var str_idx *types.String
	var strict = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			value = fp.ParseZval()
			array = fp.ParseArray()
			fp.StartOptional()
			strict = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if strict != 0 {
		if value.IsType(types.IS_LONG) {
			var __ht = array.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if entry.IsType(types.IS_LONG) && entry.Long() == value.Long() {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else {
			var __ht = array.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		}
	} else {
		if value.IsType(types.IS_LONG) {
			var __ht = array.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckLong(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else if value.IsType(types.IS_STRING) {
			var __ht = array.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckString(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else {
			var __ht = array.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckFunction(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		}
	}
	return_value.SetFalse()
	return
}
func ZifInArray(executeData zpp.Ex, return_value zpp.Ret, needle *types.Zval, haystack *types.Zval, _ zpp.Opt, strict *types.Zval) {
	PhpSearchArray(executeData, return_value, 0)
}
func ZifArraySearch(executeData zpp.Ex, return_value zpp.Ret, needle *types.Zval, haystack *types.Zval, _ zpp.Opt, strict *types.Zval) {
	PhpSearchArray(executeData, return_value, 1)
}
func PhpValidVarName(var_name *byte, var_name_len int) int {
	/* first 256 bits for first character, and second 256 bits for the next */

	var charset = []uint32{0x0, 0x0, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var charset2 = []uint32{0x0, 0x3ff0000, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var i int
	var ch uint32
	if var_name_len == 0 {
		return 0
	}

	/* These are allowed as first char: [a-zA-Z_\x7f-\xff] */

	ch = uint32((*uint8)(var_name))[0]
	if zend.ZEND_BIT_TEST(charset, ch) == 0 {
		return 0
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */

	if var_name_len > 1 {
		i = 1
		for {
			ch = uint32((*uint8)(var_name))[i]
			if zend.ZEND_BIT_TEST(charset2, ch) == 0 {
				return 0
			}
			if b.PreInc(&i) >= var_name_len {
				break
			}
		}
	}
	return 1
}
func PhpPrefixVarname(result *types.Zval, prefix *types.Zval, var_name *byte, var_name_len int, add_underscore types.ZendBool) int {
	result.SetString(types.ZendStringAlloc(prefix.String().GetLen()+b.Cond(add_underscore != 0, 1, 0)+var_name_len, 0))
	memcpy(result.String().GetVal(), prefix.String().GetVal(), prefix.String().GetLen())
	if add_underscore != 0 {
		result.String().GetStr()[prefix.String().GetLen()] = '_'
	}
	memcpy(result.String().GetVal()+prefix.String().GetLen()+b.Cond(add_underscore != 0, 1, 0), var_name, var_name_len+1)
	return types.SUCCESS
}
func PhpExtractRefIfExists(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if var_name.GetStr() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			// zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.Reference())
			count++
		}
	}
	return count
}
func PhpExtractIfExists(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if var_name.GetStr() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
			count++
		}
	}
	return count
}
func PhpExtractRefOverwrite(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			// zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.Reference())
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractOverwrite(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			//entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractRefPrefixIfExists(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						// zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.Reference())
					} else {
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		}
	}
	return count
}
func PhpExtractPrefixIfExists(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							// types.ZendStringReleaseEx(final_name.String(), 0)
							return -1
						}
					} else {
						//entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		}
	}
	return count
}
func PhpExtractRefPrefixSame(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetStr(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.StringVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						// zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.Reference())
					} else {
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "this" {
				goto prefix
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractPrefixSame(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							// types.ZendStringReleaseEx(final_name.String(), 0)
							return -1
						}
					} else {
						//entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "this" {
				goto prefix
			}
			entry = types.ZVAL_DEREF(entry)
			//entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractRefPrefixAll(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
			if final_name.StringVal() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.IsReference() {
					// 					entry.AddRefcount()
				} else {
					types.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
					if orig_var.IsIndirect() {
						orig_var = orig_var.Indirect()
					}
					// zend.ZvalPtrDtor(orig_var)
					orig_var.SetReference(entry.Reference())
				} else {
					symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
				}
				count++
			}
		}

	}
	return count
}
func PhpExtractPrefixAll(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
			if final_name.StringVal() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				entry = types.ZVAL_DEREF(entry)
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
					if orig_var.IsIndirect() {
						orig_var = orig_var.Indirect()
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
					if zend.EG__().GetException() != nil {
						// types.ZendStringReleaseEx(final_name.String(), 0)
						return -1
					}
				} else {
					//entry.TryAddRefcount()
					symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
				}
				count++
			}
		}

	}
	return count
}
func PhpExtractRefPrefixInvalid(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || var_name.GetStr() == "this" {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

				continue
			}
		}
		if final_name.StringVal() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
				if orig_var.IsIndirect() {
					orig_var = orig_var.Indirect()
				}
				// zend.ZvalPtrDtor(orig_var)
				orig_var.SetReference(entry.Reference())
			} else {
				symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
			}
			count++
		}

	}
	return count
}
func PhpExtractPrefixInvalid(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || var_name.GetStr() == "this" {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

				continue
			}
		}
		if final_name.StringVal() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			entry = types.ZVAL_DEREF(entry)
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
				if orig_var.IsIndirect() {
					orig_var = orig_var.Indirect()
				}
				zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
				if zend.EG__().GetException() != nil {
					// types.ZendStringReleaseEx(final_name.String(), 0)
					return -1
				}
			} else {
				//entry.TryAddRefcount()
				symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
			}
			count++
		}

	}
	return count
}
func PhpExtractRefSkip(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
				}
			}
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractSkip(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
				}
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			// entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func ZifExtract(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_array_param *types.Zval
	var prefix *types.Zval = nil
	var extract_refs zend.ZendLong
	var extract_type = EXTR_OVERWRITE
	var count zend.ZendLong
	var symbol_table *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			var_array_param = fp.ParseArrayEx2(false, true, false)
			fp.StartOptional()
			extract_type = fp.ParseLong()
			prefix = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	extract_refs = extract_type & EXTR_REFS
	if extract_refs != 0 {
		types.SeparateArray(var_array_param)
	}
	extract_type &= 0xff
	if extract_type < EXTR_OVERWRITE || extract_type > EXTR_IF_EXISTS {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid extract type")
		return
	}
	if extract_type > EXTR_SKIP && extract_type <= EXTR_PREFIX_IF_EXISTS && executeData.NumArgs() < 3 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "specified extract type requires the prefix parameter")
		return
	}
	if prefix != nil {
		if zend.TryConvertToString(prefix) == 0 {
			return
		}
		if prefix.String().GetLen() != 0 && PhpValidVarName(prefix.String().GetVal(), prefix.String().GetLen()) == 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "prefix is not a valid identifier")
			return
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if extract_refs != 0 {
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractRefIfExists(var_array_param.Array(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractRefOverwrite(var_array_param.Array(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractRefPrefixIfExists(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractRefPrefixSame(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractRefPrefixAll(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractRefPrefixInvalid(var_array_param.Array(), symbol_table, prefix)
		default:
			count = PhpExtractRefSkip(var_array_param.Array(), symbol_table)
		}
	} else {

		/* The array might be stored in a local variable that will be overwritten */

		var array_copy types.Zval
		types.ZVAL_COPY(&array_copy, var_array_param)
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractIfExists(array_copy.Array(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractOverwrite(array_copy.Array(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractPrefixIfExists(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractPrefixSame(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractPrefixAll(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractPrefixInvalid(array_copy.Array(), symbol_table, prefix)
		default:
			count = PhpExtractSkip(array_copy.Array(), symbol_table)
		}
		// zend.ZvalPtrDtor(&array_copy)
	}
	return_value.SetLong(count)
	return
}
func PhpCompactVar(eg_active_symbol_table *types.Array, return_value *types.Zval, entry *types.Zval) {
	var value_ptr *types.Zval
	var data types.Zval
	entry = types.ZVAL_DEREF(entry)
	if entry.IsType(types.IS_STRING) {
		if b.Assign(&value_ptr, types.ZendHashFindInd(eg_active_symbol_table, entry.String().GetStr())) != nil {
			value_ptr = types.ZVAL_DEREF(value_ptr)
			// value_ptr.TryAddRefcount()
			return_value.Array().KeyUpdate(entry.String().GetStr(), value_ptr)
		} else if entry.StringVal() == "this" {
			var object = zend.ZendGetThisObject(zend.CurrEX())
			if object != nil {
				// 				object.AddRefcount()
				data.SetObject(object)
				return_value.Array().KeyUpdate(entry.String().GetStr(), &data)
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "Undefined variable: %s", entry.String().GetVal())
		}
	} else if entry.IsType(types.IS_ARRAY) {
		if entry.IsRefcounted() {
			if entry.IsRecursive() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
				return
			}
			entry.ProtectRecursive()
		}
		var __ht = entry.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value_ptr = _z
			PhpCompactVar(eg_active_symbol_table, return_value, value_ptr)
		}
		if entry.IsRefcounted() {
			entry.UnprotectRecursive()
		}
	}
}
func ZifCompact(executeData zpp.Ex, return_value zpp.Ret, varNames []*types.Zval) {
	var args *types.Zval = nil
	var num_args uint32
	var i uint32
	var symbol_table *types.Array
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			args, num_args = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if zend.ZendForbidDynamicCall("compact()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}

	/* compact() is probably most used with a single array of var_names
	   or multiple string names, rather than a combination of both.
	   So quickly guess a minimum result size based on that */

	if num_args != 0 && args[0].IsType(types.IS_ARRAY) {
		zend.ArrayInitSize(return_value, args[0].Array().Len())
	} else {
		zend.ArrayInitSize(return_value, num_args)
	}
	for i = 0; i < num_args; i++ {
		PhpCompactVar(symbol_table, return_value, &args[i])
	}
}
func ZifArrayFill(executeData zpp.Ex, return_value zpp.Ret, startKey int, num int, val *types.Zval) (*types.Array, bool) {
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
	//if startKey >= 0 && startKey < num {
	//
	//}

	/* create hash */
	arr := types.NewArray(num)
	for i := 0; i < num; i++ {
		arr.IndexAdd(startKey+i, val)
	}
	return arr, true
}
func ZifArrayFillKeys(executeData zpp.Ex, return_value zpp.Ret, keys *types.Zval, val *types.Zval) {
	var keys *types.Zval
	var val *types.Zval
	var entry *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			keys = fp.ParseArray()
			val = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, keys.Array().Len())
	var __ht = keys.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		entry = types.ZVAL_DEREF(entry)
		// val.TryAddRefcount()
		if entry.IsType(types.IS_LONG) {
			return_value.Array().IndexUpdate(entry.Long(), val)
		} else {
			var tmp_key *types.String
			var key = zend.ZvalGetTmpString(entry, &tmp_key)
			return_value.Array().SymtableUpdate(key.GetStr(), val)
		}
	}
}
func ZifRange(executeData zpp.Ex, return_value zpp.Ret, low *types.Zval, high *types.Zval, _ zpp.Opt, step *types.Zval) {
	var zlow *types.Zval
	var zhigh *types.Zval
	var zstep *types.Zval = nil
	var tmp types.Zval
	var err = 0
	var is_step_double = 0
	var step = 1.0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			zlow = fp.ParseZval()
			zhigh = fp.ParseZval()
			fp.StartOptional()
			zstep = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if zstep != nil {
		if zstep.IsType(types.IS_DOUBLE) {
			is_step_double = 1
		} else if zstep.IsType(types.IS_STRING) {
			var type_ int = zend.IsNumericString(zstep.String().GetStr(), nil, nil, 0)
			if type_ == types.IS_DOUBLE {
				is_step_double = 1
			}
			if type_ == 0 {

				/* bad number */

				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid range string - must be numeric")
				return_value.SetFalse()
				return
			}
		}
		step = zend.ZvalGetDouble(zstep)

		/* We only want positive step values. */

		if step < 0.0 {
			step *= -1
		}

		/* We only want positive step values. */

	}

	/* If the range is given as strings, generate an array of characters. */

	if zlow.IsType(types.IS_STRING) && zhigh.IsType(types.IS_STRING) && zlow.String().GetLen() >= 1 && zhigh.String().GetLen() >= 1 {
		var type1 int
		var type2 int
		var low uint8
		var high uint8
		var lstep = zend.ZendLong(step)
		type1 = zend.IsNumericString(zlow.String().GetStr(), nil, nil, 0)
		type2 = zend.IsNumericString(zhigh.String().GetStr(), nil, nil, 0)
		if type1 == types.IS_DOUBLE || type2 == types.IS_DOUBLE || is_step_double != 0 {
			goto double_str
		} else if type1 == types.IS_LONG || type2 == types.IS_LONG {
			goto long_str
		}
		low = uint8(zlow.String().GetStr()[0])
		high = uint8(zhigh.String().GetStr()[0])
		if low > high {
			if lstep <= 0 {
				err = 1
				goto err
			}

			/* Initialize the return_value as an array. */

			zend.ArrayInitSize(return_value, uint32((low-high)/lstep+1))

			for {
				fillScope := types.PackedFillStart(return_value.Array())
				for ; low >= high; low -= uint(lstep) {
					fillScope.FillSetInternedStr(types.NewString(string(low)))
					fillScope.FillNext()
					if signed__int(low-lstep) < 0 {
						break
					}
				}
				fillScope.FillEnd()
				break
			}
		} else if high > low {
			if lstep <= 0 {
				err = 1
				goto err
			}
			zend.ArrayInitSize(return_value, uint32((high-low)/lstep+1))

			for {
				fillScope := types.PackedFillStart(return_value.Array())
				for ; low <= high; low += uint(lstep) {
					fillScope.FillSetInternedStr(types.NewString(string(low)))
					fillScope.FillNext()
					if signed__int(low+lstep) > 255 {
						break
					}
				}
				fillScope.FillEnd()
				break
			}
		} else {
			zend.ArrayInit(return_value)
			tmp.SetStringVal(string(low))
			return_value.Array().NextIndexInsertNew(&tmp)
		}
	} else if zlow.IsType(types.IS_DOUBLE) || zhigh.IsType(types.IS_DOUBLE) || is_step_double != 0 {
		var low float64
		var high float64
		var element float64
		var i uint32
		var size uint32
	double_str:
		low = zend.ZvalGetDouble(zlow)
		high = zend.ZvalGetDouble(zhigh)
		if core.ZendIsInf(high) || core.ZendIsInf(low) {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid range supplied: start=%0.0f end=%0.0f", low, high)
			return_value.SetFalse()
			return
		}
		if low > high {
			if low-high < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size = (low-high)/step + 1
			if __calc_size >= float64(types.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", high, low)
				return_value.SetFalse()
				return
			}
			size = uint32(math.Round(__calc_size))
			zend.ArrayInitSize(return_value, size)

			fillScope := types.PackedFillStart(return_value.Array())
			i = 0
			element = low
			for i < size && element >= high {
				fillScope.FillSetDouble(element)
				fillScope.FillNext()
				i++
				element = low - i*step
			}
			fillScope.FillEnd()
		} else if high > low {
			if high-low < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size = (high-low)/step + 1
			if __calc_size >= float64(types.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high)
				return_value.SetFalse()
				return
			}
			size = uint32(math.Round(__calc_size))
			zend.ArrayInitSize(return_value, size)

			fillScope := types.PackedFillStart(return_value.Array())
			i = 0
			element = low
			for i < size && element <= high {
				fillScope.FillSetDouble(element)
				fillScope.FillNext()
				i++
				element = low + i*step
			}
			fillScope.FillEnd()
		} else {
			zend.ArrayInit(return_value)
			tmp.SetDouble(low)
			return_value.Array().NextIndexInsertNew(&tmp)
		}
	} else {
		var low zend.ZendLong
		var high zend.ZendLong

		/* lstep is a zend_ulong so that comparisons to it don't overflow, i.e. low - high < lstep */

		var lstep zend.ZendUlong
		var i uint32
		var size uint32
	long_str:
		low = zend.ZvalGetLong(zlow)
		high = zend.ZvalGetLong(zhigh)
		if step <= 0 {
			err = 1
			goto err
		}
		lstep = zend.ZendUlong(step)
		if step <= 0 {
			err = 1
			goto err
		}
		if low > high {
			if zend.ZendUlong(low-high < lstep) != 0 {
				err = 1
				goto err
			}
			var __calc_size = zend.ZendUlong(low-high) / lstep
			if __calc_size >= types.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, high, low)
				return_value.SetFalse()
				return
			}
			size = uint32(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)

			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < size; i++ {
				fillScope.FillSetLong(low - i*lstep)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		} else if high > low {
			if zend.ZendUlong(high-low < lstep) != 0 {
				err = 1
				goto err
			}
			var __calc_size = zend.ZendUlong(high-low) / lstep
			if __calc_size >= types.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, low, high)
				return_value.SetFalse()
				return
			}
			size = uint32(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)

			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < size; i++ {
				fillScope.FillSetLong(low + i*lstep)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		} else {
			zend.ArrayInit(return_value)
			tmp.SetLong(low)
			return_value.Array().NextIndexInsertNew(&tmp)
		}
	}
err:
	if err != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "step exceeds the specified range")
		return_value.SetFalse()
		return
	}
}
func arrayDataShuffle(array *types.Array) *types.Array {
	values := array.Values()
	for i := len(values) - 1; i >= 0; i-- {
		j := PhpMtRandRange(0, i)
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
func PhpSplice(in_hash *types.Array, offset zend.ZendLong, length zend.ZendLong, replace *types.Array, removed *types.Array) {
	var out_hash *types.Array
	var num_in zend.ZendLong
	var pos zend.ZendLong
	var idx uint32
	var p *types.Bucket
	var entry *types.Zval
	var iter_pos = types.ZendHashIteratorsLowerPos(in_hash, 0)

	/* Get number of entries in the input hash */

	num_in = in_hash.Len()

	/* Clamp the offset.. */

	if offset > num_in {
		offset = num_in
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if unsigned(offset+unsigned(length)) > unsigned(num_in) {
		length = num_in - offset
	}

	/* Create and initialize output hash */

	out_hash = types.NewArray(b.Cond(length > 0, num_in-length, 0) + b.CondF1(replace != nil, func() int { return replace.Len() }, 0))

	/* Start at the beginning of the input hash and copy entries to output hash until offset is reached */

	pos = 0
	idx = 0
	for ; pos < offset && idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.Bucket(idx)
		if p.GetVal().IsUndef() {
			continue
		}

		/* Get entry and increase reference count */

		entry = p.GetVal()

		/* Update output hash depending on key type */

		if p.GetKey() == nil {
			out_hash.NextIndexInsertNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* If hash for removed entries exists, go until offset+length and copy the entries to it */

	if removed != nil {
		for ; pos < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.Bucket(idx)
			if p.GetVal().IsUndef() {
				continue
			}
			pos++
			entry = p.GetVal()
			// entry.TryAddRefcount()
			if p.GetKey() == nil {
				removed.NextIndexInsertNew(entry)
				types.ZendHashDelBucket(in_hash, p)
			} else {
				removed.KeyAddNew(p.GetKey().GetStr(), entry)
				if in_hash == zend.EG__().GetSymbolTable() {
					zend.ZendDeleteGlobalVariable(p.GetKey())
				} else {
					types.ZendHashDelBucket(in_hash, p)
				}
			}
		}
	} else {
		var pos2 = pos
		for ; pos2 < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.Bucket(idx)
			if p.GetVal().IsUndef() {
				continue
			}
			pos2++
			if p.GetKey() != nil && in_hash == zend.EG__().GetSymbolTable() {
				zend.ZendDeleteGlobalVariable(p.GetKey())
			} else {
				types.ZendHashDelBucket(in_hash, p)
			}
		}
	}
	iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos)

	/* If there are entries to insert.. */

	if replace != nil {
		var __ht = replace
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			entry = _z
			// entry.TryAddRefcount()
			out_hash.NextIndexInsertNew(entry)
			pos++
		}
	}

	/* Copy the remaining input hash entries to the output hash */

	for ; idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.Bucket(idx)
		if p.GetVal().IsUndef() {
			continue
		}
		entry = p.GetVal()
		if p.GetKey() == nil {
			out_hash.NextIndexInsertNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			iter_pos = types.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* replace HashTable data */

	in_hash.SetPDestructor(nil)
	in_hash.Destroy()

	in_hash.SetBy(out_hash)
}
func ZifArrayPush(executeData zpp.Ex, return_value zpp.Ret, stack zpp.RefZval, _ zpp.Opt, vars []*types.Zval) {
	var args *types.Zval
	var stack *types.Zval
	var new_var types.Zval
	var i int
	var argc int
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArrayEx(false, true)
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* For each subsequent argument, make it a reference, increase refcount, and add it to the end of the array */

	for i = 0; i < argc; i++ {
		types.ZVAL_COPY(&new_var, &args[i])
		if stack.Array().NextIndexInsert(&new_var) == nil {
			//new_var.TryDelRefcount()
			core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return_value.SetFalse()
			return
		}
	}

	/* Clean up and return the number of values in the stack */

	return_value.SetLong(stack.Array().Len())

	/* Clean up and return the number of values in the stack */
}
func ZifArrayPop(executeData zpp.Ex, return_value zpp.Ret, stack zpp.RefZval) {
	var stack *types.Zval
	var val *types.Zval
	var idx uint32
	var p *types.Bucket
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			stack = fp.ParseArrayEx(false, true)
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if stack.Array().Len() == 0 {
		return
	}

	/* Get the last value and copy it into the return value */

	idx = stack.Array().GetNNumUsed()
	for true {
		if idx == 0 {
			return
		}
		idx--
		p = stack.Array().Bucket(idx)
		val = p.GetVal()
		if val.IsIndirect() {
			val = val.Indirect()
		}
		if val.IsNotUndef() {
			break
		}
	}
	types.ZVAL_COPY_DEREF(return_value, val)
	if p.GetKey() == nil && stack.Array().GetNNextFreeElement() > 0 && p.GetH() >= zend_ulong(stack.Array().GetNNextFreeElement()-1) {
		stack.Array().SetNNextFreeElement(stack.Array().GetNNextFreeElement() - 1)
	}

	/* Delete the last value */

	if p.GetKey() != nil && stack.Array() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		types.ZendHashDelBucket(stack.Array(), p)
	}
	types.ZendHashInternalPointerReset(stack.Array())
}
func ZifArrayShift(stack zpp.RefArray) {
	if stack.Array().Len() == 0 {
		return
	}

	var p *types.Bucket

	/* Get the first value and copy it into the return value */
	key, val := stack.Array().FirstIndirect()
	if val == nil {
		return
	}

	val = val.DeRef()

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
			key = types.IndexKey(k)
			k++
		}
		return key, value
	})

	// reset internal pointer
	types.ZendHashInternalPointerReset(stack.Array())
}
func ZifArrayUnshift(executeData zpp.Ex, return_value zpp.Ret, stack zpp.RefZval, vars []*types.Zval) {
	var args []*types.Zval = vars
	var new_hash *types.Array
	var argc int
	var i int
	var key *types.String
	var value *types.Zval
	new_hash = types.NewArray(stack.Array().Len() + argc)
	for i = 0; i < argc; i++ {
		args[i].TryAddRefcount()
		new_hash.NextIndexInsertNew(&args[i])
	}
	var __ht = stack.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			new_hash.KeyAddNew(key.GetStr(), value)
		} else {
			new_hash.NextIndexInsertNew(value)
		}
	}

	/* replace HashTable data */

	stack.Array().SetPDestructor(nil)
	stack.Array().Destroy()

	stack.Array().SetBy(new_hash)

	/* Clean up and return the number of elements in the stack */

	return_value.SetLong(stack.Array().Len())

	/* Clean up and return the number of elements in the stack */
}
func ZifArraySplice(executeData zpp.Ex, return_value zpp.Ret, arg zpp.RefZval, offset *types.Zval, _ zpp.Opt, length *types.Zval, replacement *types.Zval) {
	var array *types.Zval
	var repl_array *types.Zval = nil
	var rem_hash *types.Array = nil
	var offset zend.ZendLong
	var length = 0
	var num_in int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			array = fp.ParseArrayEx(false, true)
			offset = fp.ParseLong()
			fp.StartOptional()
			length = fp.ParseLong()
			repl_array = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	num_in = array.Array().Len()
	if executeData.NumArgs() < 3 {
		length = num_in
	}
	if executeData.NumArgs() == 4 {

		/* Make sure the last argument, if passed, is an array */

		zend.ConvertToArrayEx(repl_array)

		/* Make sure the last argument, if passed, is an array */

	}

	/* Don't create the array of removed elements if it's not going
	 * to be used; e.g. only removing and/or replacing elements */

	if zend.USED_RET() {
		var size = length

		/* Clamp the offset.. */

		if offset > num_in {
			offset = num_in
		} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
			offset = 0
		}

		/* ..and the length */

		if length < 0 {
			size = num_in - offset + length
		} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > uint32(num_in) {
			size = num_in - offset
		}

		/* Initialize return value */

		zend.ArrayInitSize(return_value, b.CondF1(size > 0, func() uint32 { return uint32(size) }, 0))
		rem_hash = return_value.Array()
	}

	/* Perform splice */

	PhpSplice(array.Array(), offset, length, b.CondF1(repl_array != nil, func() *types.Array { return repl_array.Array() }, nil), rem_hash)

	/* Perform splice */
}
func ZifArraySlice(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, offset *types.Zval, _ zpp.Opt, length *types.Zval, preserveKeys *types.Zval) {
	var input *types.Zval
	var z_length *types.Zval = nil
	var entry *types.Zval
	var offset zend.ZendLong
	var length = 0
	var preserve_keys = 0
	var num_in int
	var pos int
	var string_key *types.String
	var num_key zend.ZendUlong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			input = fp.ParseArray()
			offset = fp.ParseLong()
			fp.StartOptional()
			z_length = fp.ParseZval()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Get number of entries in the input hash */

	num_in = input.Array().Len()

	/* We want all entries from offset to the end if length is not passed or is null */

	if executeData.NumArgs() < 3 || z_length.IsType(types.IS_NULL) {
		length = num_in
	} else {
		length = zend.ZvalGetLong(z_length)
	}

	/* Clamp the offset.. */

	if offset > num_in {
		return_value.SetEmptyArray()
		return
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > unsigned(num_in) {
		length = num_in - offset
	}
	if length <= 0 {
		return_value.SetEmptyArray()
		return
	}

	/* Initialize returned array */

	zend.ArrayInitSize(return_value, uint32(length))

	/* Start at the beginning and go until we hit offset */

	pos = 0
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		pos++
		if pos <= offset {
			continue
		}
		if pos > offset+length {
			break
		}
		if string_key != nil {
			entry = return_value.Array().KeyAddNew(string_key.GetStr(), entry)
		} else {
			if preserve_keys != 0 {
				entry = return_value.Array().IndexAddNew(num_key, entry)
			} else {
				entry = return_value.Array().NextIndexInsertNew(entry)
			}
		}
		zend.ZvalAddRef(entry)
	}
}
func PhpArrayMergeRecursive(dest *types.Array, src *types.Array) int {
	for iter := src.Iterator(); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Current()
		if !key.IsStrKey() {
			dest.NextIndexInsert(value)
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
		if thash != nil && thash.IsRecursive() || value == destEntry && destEntry.IsReference() && destEntry.GetRefcount()%2 != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		b.Assert(!(destEntry.IsReference()) || destEntry.GetRefcount() > 1)
		types.SeparateZval(destEntry)
		destZval = destEntry
		if destZval.IsType(types.IS_NULL) {
			zend.ConvertToArrayEx(destZval)
			zend.AddNextIndexNull(destZval)
		} else {
			zend.ConvertToArrayEx(destZval)
		}
		tmp.SetUndef()
		if srcZval.IsType(types.IS_OBJECT) {
			types.ZVAL_COPY(&tmp, srcZval)
			zend.ConvertToArray(&tmp)
			srcZval = &tmp
		}
		if srcZval.IsType(types.IS_ARRAY) {
			if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
				thash.ProtectRecursive()
			}
			ret = PhpArrayMergeRecursive(destZval.Array(), srcZval.Array())
			if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
				thash.UnprotectRecursive()
			}
			if ret == 0 {
				return 0
			}
		} else {
			destZval.Array().NextIndexInsert(srcZval)
		}
	}
	return 1
}
func PhpArrayMerge(dest *types.Array, src *types.Array) int {
	src.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if value.IsReference() && value.GetRefcount() == 1 {
			value = types.Z_REFVAL_P(value)
		}
		if key.IsStrKey() {
			dest.KeyUpdate(key.StrKey(), value)
		} else {
			dest.NextIndexInsert(value)
		}
	})
}
func PhpArrayReplaceRecursive(dest *types.Array, src *types.Array) int {
	var src_entry *types.Zval
	var dest_entry *types.Zval
	var src_zval *types.Zval
	var dest_zval *types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var ret int
	var __ht = src
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		src_entry = _z
		src_zval = src_entry
		src_zval = types.ZVAL_DEREF(src_zval)
		if string_key != nil {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.KeyFind(string_key.GetStr())) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv = dest.KeyUpdate(string_key.GetStr(), src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		} else {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.IndexFind(num_key)) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv = dest.IndexUpdate(num_key, src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		}
		dest_zval = dest_entry
		dest_zval = types.ZVAL_DEREF(dest_zval)
		if dest_zval.IsRecursive() || src_zval.IsRecursive() || src_entry.IsReference() && dest_entry.IsReference() && src_entry.Reference() == dest_entry.Reference() && dest_entry.GetRefcount()%2 != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
			return 0
		}
		b.Assert(!(dest_entry.IsReference()) || dest_entry.GetRefcount() > 1)
		types.SeparateZval(dest_entry)
		dest_zval = dest_entry
		if dest_zval.IsRefcounted() {
			dest_zval.ProtectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.ProtectRecursive()
		}
		ret = PhpArrayReplaceRecursive(dest_zval.Array(), src_zval.Array())
		if dest_zval.IsRefcounted() {
			dest_zval.UnprotectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.UnprotectRecursive()
		}
		if ret == 0 {
			return 0
		}
	}
	return 1
}
func PhpArrayReplaceWrapper(executeData *zend.ZendExecuteData, return_value *types.Zval, recursive int) {
	var args *types.Zval = nil
	var arg *types.Zval
	var argc int
	var i int
	var dest *types.Array
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	for i = 0; i < argc; i++ {
		var arg *types.Zval = args + i
		if arg.GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(arg))
			return_value.SetNull()
			return
		}
	}

	/* copy first array */

	arg = args
	dest = types.ZendArrayDup(arg.Array())
	return_value.SetArray(dest)
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayReplaceRecursive(dest, arg.Array())
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			types.ZendHashMerge(dest, arg.Array(), zend.ZvalAddRef, 1)
		}
	}
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
		if value.IsReference() && value.GetRefcount() == 1 {
			value = types.Z_REFVAL_P(value)
		}
		if key.IsStrKey() {
			arr.KeyUpdate(key.StrKey(), value)
		} else {
			arr.NextIndexInsert(value)
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
func ZifArrayReplace(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, _ zpp.Opt, arrays []*types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 0)
}
func ZifArrayReplaceRecursive(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, _ zpp.Opt, arrays []*types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 1)
}
func ZifArrayKeys(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, _ zpp.Opt, searchValue *types.Zval, strict *types.Zval) {
	var input *types.Zval
	var search_value *types.Zval = nil
	var entry *types.Zval
	var new_val types.Zval
	var strict = 0
	var num_idx zend.ZendUlong
	var str_idx *types.String
	var arrval *types.Array
	var elem_count zend.ZendUlong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			input = fp.ParseArray()
			fp.StartOptional()
			search_value = fp.ParseZval()
			strict = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	arrval = input.Array()
	elem_count = arrval.Len()

	/* Base case: empty input */

	if elem_count == 0 {
		zend.ZVAL_ZVAL(return_value, input, 1, 0)
		return
	}

	/* Initialize return array */

	if search_value != nil {
		zend.ArrayInit(return_value)
		if strict != 0 {
			var __ht = arrval
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(search_value, entry) != 0 {
					if str_idx != nil {
						new_val.SetStringCopy(str_idx)
					} else {
						new_val.SetLong(num_idx)
					}
					return_value.Array().NextIndexInsertNew(&new_val)
				}
			}
		} else {
			var __ht = arrval
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckFunction(search_value, entry) != 0 {
					if str_idx != nil {
						new_val.SetStringCopy(str_idx)
					} else {
						new_val.SetLong(num_idx)
					}
					return_value.Array().NextIndexInsertNew(&new_val)
				}
			}
		}
	} else {
		zend.ArrayInitSize(return_value, elem_count)

		fillScope := types.PackedFillStart(return_value.Array())

		/* Go through input array and add keys to the return array */

		var __ht = input.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			num_idx = _p.GetH()
			str_idx = _p.GetKey()
			entry = _z
			if str_idx != nil {
				fillScope.FillSetStringCopy(str_idx)
			} else {
				fillScope.FillSetLong(num_idx)
			}
			fillScope.FillNext()
		}

		/* Go through input array and add keys to the return array */

		fillScope.FillEnd()
	}

	/* Initialize return array */
}
func ZifArrayKeyFirst(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var stack *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			stack = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	var target_hash = stack.Array()
	var pos types.ArrayPosition = 0
	types.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayKeyLast(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var stack *types.Zval
	var pos types.ArrayPosition
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			stack = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	var target_hash = stack.Array()
	types.ZendHashInternalPointerEndEx(target_hash, &pos)
	types.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayValues(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var arrval *types.Array
	var arrlen zend.ZendLong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			input = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	arrval = input.Array()

	/* Return empty input as is */

	arrlen = arrval.Len()
	if arrlen == 0 {
		return_value.SetEmptyArray()
		return
	}

	/* Return vector-like packed arrays as-is */

	/* Initialize return array */

	zend.ArrayInitSize(return_value, arrval.Len())

	/* Go through input array and add values to the return array */

	fillScope := types.PackedFillStart(return_value.Array())
	var __ht = arrval
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		if entry.IsReference() && entry.GetRefcount() == 1 {
			entry = types.Z_REFVAL_P(entry)
		}
		// entry.TryAddRefcount()
		fillScope.FillSet(entry)
		fillScope.FillNext()
	}
	fillScope.FillEnd()

	/* Go through input array and add values to the return array */
}
func ZifArrayCountValues(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var tmp *types.Zval
	var myht *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			input = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)

	/* Go through input array and add values to the return array */

	myht = input.Array()
	var __ht = myht
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if b.Assign(&tmp, return_value.Array().IndexFind(entry.Long())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.Array().IndexUpdate(entry.Long(), &data)
			} else {
				tmp.Long()++
			}
		} else if entry.IsType(types.IS_STRING) {
			if b.Assign(&tmp, return_value.Array().SymtableFind(entry.String().GetStr())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.Array().SymtableUpdate(entry.String().GetStr(), &data)
			} else {
				tmp.Long()++
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only count STRING and INTEGER values!")
		}
	}
}
func ArrayColumnParamHelper(param *types.Zval, name string) types.ZendBool {
	switch param.GetType() {
	case types.IS_DOUBLE:
		if param.GetType() != types.IS_LONG {
			zend.ConvertToLong(param)
		}
		fallthrough
	case types.IS_LONG:
		return 1
	case types.IS_OBJECT:
		if zend.TryConvertToString(param) == 0 {
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
func ArrayColumnFetchProp(data *types.Zval, name *types.Zval, rv *types.Zval) *types.Zval {
	var prop *types.Zval = nil
	if data.IsType(types.IS_OBJECT) {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_EXISTS, nil) != 0 || types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_ISSET, nil) != 0 {
			prop = types.Z_OBJ_HT(*data).GetReadProperty()(data, name, zend.BP_VAR_R, nil, rv)
			if prop != nil {
				prop = types.ZVAL_DEREF(prop)
				if prop != rv {
					// prop.TryAddRefcount()
				}
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.IsType(types.IS_ARRAY) {
		if name.IsType(types.IS_STRING) {
			prop = data.Array().SymtableFind(name.String().GetStr())
		} else if name.IsType(types.IS_LONG) {
			prop = data.Array().IndexFind(name.Long())
		}
		if prop != nil {
			prop = types.ZVAL_DEREF(prop)
			// prop.TryAddRefcount()
		}
	}
	return prop
}
func ZifArrayColumn(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, columnKey *types.Zval, _ zpp.Opt, indexKey *types.Zval) {
	var input *types.Array
	var colval *types.Zval
	var data *types.Zval
	var rv types.Zval
	var column *types.Zval = nil
	var index *types.Zval = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			input = fp.ParseArrayHt()
			column = fp.ParseZvalEx(true, false)
			fp.StartOptional()
			index = fp.ParseZvalEx(true, false)
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if column != nil && ArrayColumnParamHelper(column, "column") == 0 || index != nil && ArrayColumnParamHelper(index, "index") == 0 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInitSize(return_value, input.Len())
	if index == nil {

		fillScope := types.PackedFillStart(return_value.Array())
		var __ht = input
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()

			data = _z
			data = types.ZVAL_DEREF(data)
			if column == nil {
				// data.TryAddRefcount()
				colval = data
			} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
				continue
			}
			fillScope.FillSet(colval)
			fillScope.FillNext()
		}
		fillScope.FillEnd()
	} else {
		var __ht = input
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()

			data = _z
			data = types.ZVAL_DEREF(data)
			if column == nil {
				// data.TryAddRefcount()
				colval = data
			} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
				continue
			}

			/* Failure will leave keyval alone which will land us on the final else block below
			 * which is to append the value as next_index
			 */

			if index != nil {
				var rv types.Zval
				var keyval = ArrayColumnFetchProp(data, index, &rv)
				if keyval != nil {
					switch keyval.GetType() {
					case types.IS_STRING:
						return_value.Array().SymtableUpdate(keyval.String().GetStr(), colval)
					case types.IS_LONG:
						return_value.Array().IndexUpdate(keyval.Long(), colval)
					case types.IS_OBJECT:
						var tmp_key *types.String
						var key = zend.ZvalGetTmpString(keyval, &tmp_key)
						return_value.Array().SymtableUpdate(key.GetStr(), colval)
						// zend.ZendTmpStringRelease(tmp_key)
					case types.IS_NULL:
						return_value.Array().KeyUpdate(types.NewString("").GetStr(), colval)
					case types.IS_DOUBLE:
						return_value.Array().IndexUpdate(zend.DvalToLval(keyval.Double()), colval)
					case types.IS_TRUE:
						return_value.Array().IndexUpdate(1, colval)
					case types.IS_FALSE:
						return_value.Array().IndexUpdate(0, colval)
					case types.IS_RESOURCE:
						return_value.Array().IndexUpdate(types.Z_RES_HANDLE_P(keyval), colval)
					default:
						return_value.Array().NextIndexInsert(colval)
					}
					// zend.ZvalPtrDtor(keyval)
				} else {
					return_value.Array().NextIndexInsert(colval)
				}
			} else {
				return_value.Array().NextIndexInsert(colval)
			}

			/* Failure will leave keyval alone which will land us on the final else block below
			 * which is to append the value as next_index
			 */

		}
	}
}
func ZifArrayReverse(executeData zpp.Ex, return_value zpp.Ret, input *types.Zval, _ zpp.Opt, preserveKeys *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var preserve_keys = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			input = fp.ParseArray()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, input.Array().Len())
	if input.Array().IsPacked() && preserve_keys == 0 {

		fillScope := types.PackedFillStart(return_value.Array())
		var __ht = input.Array()
		for _, _p := range __ht.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			entry = _z
			if entry.IsReference() && entry.GetRefcount() == 1 {
				entry = types.Z_REFVAL_P(entry)
			}
			// entry.TryAddRefcount()
			fillScope.FillSet(entry)
			fillScope.FillNext()
		}
		fillScope.FillEnd()
	} else {
		var __ht = input.Array()
		for _, _p := range __ht.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			num_key = _p.GetH()
			string_key = _p.GetKey()
			entry = _z
			if string_key != nil {
				entry = return_value.Array().KeyAddNew(string_key.GetStr(), entry)
			} else {
				if preserve_keys != 0 {
					entry = return_value.Array().IndexAddNew(num_key, entry)
				} else {
					entry = return_value.Array().NextIndexInsertNew(entry)
				}
			}
			zend.ZvalAddRef(entry)
		}
	}
}
func ZifArrayPad(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, padSize *types.Zval, padValue *types.Zval) {
	var input *types.Zval
	var pad_value *types.Zval
	var pad_size zend.ZendLong
	var pad_size_abs zend.ZendLong
	var input_size zend.ZendLong
	var num_pads zend.ZendLong
	var i zend.ZendLong
	var key *types.String
	var value *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			input = fp.ParseArray()
			pad_size = fp.ParseLong()
			pad_value = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Do some initial calculations */

	input_size = input.Array().Len()
	pad_size_abs = zend.ZEND_ABS(pad_size)
	if pad_size_abs < 0 || pad_size_abs-input_size > int64(1048576) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return_value.SetFalse()
		return
	}
	if input_size >= pad_size_abs {

		/* Copy the original array */

		types.ZVAL_COPY(return_value, input)
		return
	}
	num_pads = pad_size_abs - input_size
	//if pad_value.IsRefcounted() {
	//	pad_value.RefCounted().AddRefcountEx(num_pads)
	//}
	zend.ArrayInitSize(return_value, pad_size_abs)
	if input.Array().IsPacked() {

		if pad_size < 0 {
			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < num_pads; i++ {
				fillScope.FillSet(pad_value)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		}
		fillScope := types.PackedFillStart(return_value.Array())
		var __ht = input.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()

			value = _z
			// value.TryAddRefcount()
			fillScope.FillSet(value)
			fillScope.FillNext()
		}
		fillScope.FillEnd()
		if pad_size > 0 {
			fillScope := types.PackedFillStart(return_value.Array())
			for i = 0; i < num_pads; i++ {
				fillScope.FillSet(pad_value)
				fillScope.FillNext()
			}
			fillScope.FillEnd()
		}
	} else {
		if pad_size < 0 {
			for i = 0; i < num_pads; i++ {
				return_value.Array().NextIndexInsertNew(pad_value)
			}
		}
		var __ht = input.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			key = _p.GetKey()
			value = _z
			// value.TryAddRefcount()
			if key != nil {
				return_value.Array().KeyAddNew(key.GetStr(), value)
			} else {
				return_value.Array().NextIndexInsertNew(value)
			}
		}
		if pad_size > 0 {
			for i = 0; i < num_pads; i++ {
				return_value.Array().NextIndexInsertNew(pad_value)
			}
		}
	}
}
func ZifArrayFlip(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var data types.Zval
	var num_idx zend.ZendUlong
	var str_idx *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			array = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, array.Array().Len())
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_idx = _p.GetH()
		str_idx = _p.GetKey()
		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.Array().IndexUpdate(entry.Long(), &data)
		} else if entry.IsType(types.IS_STRING) {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.Array().SymtableUpdate(entry.String().GetStr(), &data)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	}
}
func ZifArrayChangeKeyCase(executeData zpp.Ex, return_value zpp.Ret, input *types.Zval, _ zpp.Opt, case_ *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var string_key *types.String
	var new_key *types.String
	var num_key zend.ZendUlong
	var change_to_upper = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			array = fp.ParseArray()
			fp.StartOptional()
			change_to_upper = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, array.Array().Len())
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		if string_key == nil {
			entry = return_value.Array().IndexUpdate(num_key, entry)
		} else {
			if change_to_upper != 0 {
				new_key = str.PhpStringToupper(string_key)
			} else {
				new_key = str.PhpStringTolower(string_key)
			}
			entry = return_value.Array().KeyUpdate(new_key.GetStr(), entry)
			// types.ZendStringReleaseEx(new_key, 0)
		}
		zend.ZvalAddRef(entry)
	}
}
func ZifArrayUnique(arg *types.Array, _ zpp.Opt, flags *int) *types.Array {
	var sortType = b.Option(flags, PHP_SORT_STRING)

	if sortType == PHP_SORT_STRING {
		existValues := make(map[string]bool, arg.Len())
		retArr := types.NewArray(arg.Len())
		arg.ForeachIndirect(func(key types.ArrayKey, val *types.Zval) {
			var strVal string = zend.ZvalGetStrVal(val)
			if _, exists := existValues[strVal]; !exists {
				if val.IsReference() && val.GetRefcount() == 1 {
					val = val.DeRef()
				}
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
	return zend.StringCompareFunction(first, second)
}
func ZvalUserCompare(a *types.Zval, b *types.Zval) int {
	var args []types.Zval
	var retval types.Zval
	types.ZVAL_COPY_VALUE(&args[0], a)
	types.ZVAL_COPY_VALUE(&args[1], b)
	BG__().user_compare_fci.param_count = 2
	BG__().user_compare_fci.params = args
	BG__().user_compare_fci.retval = &retval
	BG__().user_compare_fci.no_separation = 0
	if zend.ZendCallFunction(&(BG__().user_compare_fci), &(BG__().user_compare_fci_cache)) == types.SUCCESS && retval.IsNotUndef() {
		var ret = zend.ZvalGetLong(&retval)
		// zend.ZvalPtrDtor(&retval)
		return zend.ZEND_NORMALIZE_BOOL(ret)
	} else {
		return 0
	}
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
	cmp := arrayDataComparer(zend.StringCompareFunction)
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
		arrayDataComparer(zend.StringCompareFunction),
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
func ZifArrayDiff(executeData zpp.Ex, return_value zpp.Ret, arr1 *types.Zval, arrays []*types.Zval) {
	var args []*types.Zval
	var argc int
	var i int
	var num uint32
	var exclude *types.Array
	var value *types.Zval
	var str *types.String
	var tmp_str *types.String
	var key *types.String
	var idx zend.ZendLong
	var dummy types.Zval
	if executeData.NumArgs() < 2 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "at least 2 parameters are required, %d given", executeData.NumArgs())
		return
	}
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if args[0].GetType() != types.IS_ARRAY {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter 1 to be an array, %s given", types.ZendZvalTypeName(&args[0]))
		return_value.SetNull()
		return
	}
	num = args[0].Array().Len()
	if num == 0 {
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
		}
		return_value.SetEmptyArray()
		return
	} else if num == 1 {
		var found = 0
		var search_str *types.String
		var tmp_search_str *types.String
		value = nil
		var __ht *types.Array = args[0].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value = _z
			break
		}
		if value == nil {
			for i = 1; i < argc; i++ {
				if args[i].GetType() != types.IS_ARRAY {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
					return_value.SetNull()
					return
				}
			}
			return_value.SetEmptyArray()
			return
		}
		search_str = zend.ZvalGetTmpString(value, &tmp_search_str)
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
			if found == 0 {
				var __ht *types.Array = args[i].Array()
				for _, _p := range __ht.ForeachData() {
					var _z = _p.GetVal()
					if _z.IsIndirect() {
						_z = _z.Indirect()
						if _z.IsUndef() {
							continue
						}
					}
					value = _z
					str = zend.ZvalGetTmpString(value, &tmp_str)
					if search_str.GetStr() == str.GetStr() {
						// zend.ZendTmpStringRelease(tmp_str)
						found = 1
						break
					}
					// zend.ZendTmpStringRelease(tmp_str)
				}
			}
		}
		// zend.ZendTmpStringRelease(tmp_search_str)
		if found != 0 {
			return_value.SetEmptyArray()
		} else {
			types.ZVAL_COPY(return_value, &args[0])
		}
		return
	}

	/* count number of elements */

	num = 0
	for i = 1; i < argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, types.ZendZvalTypeName(&args[i]))
			return_value.SetNull()
			return
		}
		num += args[i].Array().Len()
	}
	if num == 0 {
		types.ZVAL_COPY(return_value, &args[0])
		return
	}
	dummy.SetNull()

	/* create exclude map */

	exclude = types.NewArray(num)
	for i = 1; i < argc; i++ {
		var __ht *types.Array = args[i].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value = _z
			str = zend.ZvalGetTmpString(value, &tmp_str)
			exclude.KeyAdd(str.GetStr(), &dummy)
			// zend.ZendTmpStringRelease(tmp_str)
		}
	}

	/* copy all elements of first array that are not in exclude set */

	zend.ArrayInitSize(return_value, args[0].Array().Len())
	var __ht *types.Array = args[0].Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		idx = _p.GetH()
		key = _p.GetKey()
		value = _z
		str = zend.ZvalGetTmpString(value, &tmp_str)
		if !exclude.KeyExists(str.GetStr()) {
			if key != nil {
				value = return_value.Array().KeyAddNew(key.GetStr(), value)
			} else {
				value = return_value.Array().IndexAddNew(idx, value)
			}
			zend.ZvalAddRef(value)
		}
		// zend.ZendTmpStringRelease(tmp_str)
	}
	exclude.Destroy()
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
		arrayDataComparer(zend.StringCompareFunction),
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
	var parseState [2]int = [...]int{0, 0}

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
					sortType = int(arg.Long())
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
		//newHash.SetNInternalPointer(0)
		for idx := 0; idx < arraySize; idx++ {
			pair := matrix[idx][i]
			if pair.GetKey().IsStrKey() {
				newHash.KeyAdd(pair.GetKey().StrKey(), pair.GetVal())
			} else {
				newHash.NextIndexInsert(pair.GetVal())
			}
		}
		newHash.SetNNextFreeElement(arraySize)
	}

	return true
}
func ZifArrayRand(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, _ zpp.Opt, numReq *types.Zval) {
	var input *types.Zval
	var num_req = 1
	var string_key *types.String
	var num_key zend.ZendUlong
	var i int
	var num_avail int
	var bitset zend.ZendBitset
	var negative_bitset = 0
	var bitset_len uint32
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			input = fp.ParseArray()
			fp.StartOptional()
			num_req = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	num_avail = input.Array().Len()
	if num_avail == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Array is empty")
		return
	}
	if num_req == 1 {
		var ht = input.Array()
		if uint32(num_avail < ht.GetNNumUsed()-(ht.GetNNumUsed()>>1)) != 0 {

			/* If less than 1/2 of elements are used, don't sample. Instead search for a
			 * specific offset using linear scan. */

			var i = 0
			var randval = PhpMtRandRange(0, num_avail-1)
			var __ht = input.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()

				num_key = _p.GetH()
				string_key = _p.GetKey()
				if i == randval {
					if string_key != nil {
						return_value.SetStringCopy(string_key)
						return
					} else {
						return_value.SetLong(num_key)
						return
					}
				}
				i++
			}
		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

		for {
			var randval = PhpMtRandRange(0, ht.GetNNumUsed()-1)
			var bucket *types.Bucket = ht.GetArData()[randval]
			if !(bucket.GetVal().IsUndef()) {
				if bucket.GetKey() != nil {
					return_value.SetStringCopy(bucket.GetKey())
					return
				} else {
					return_value.SetLong(bucket.GetH())
					return
				}
			}

		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

	}
	if num_req <= 0 || num_req > num_avail {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Second argument has to be between 1 and the number of elements in the array")
		return
	}

	/* Make the return value an array only if we need to pass back more than one result. */

	zend.ArrayInitSize(return_value, uint32(num_req))
	if num_req > num_avail>>1 {
		negative_bitset = 1
		num_req = num_avail - num_req
	}
	bitset_len = zend.ZendBitsetLen(num_avail)
	bitset = zend.ZEND_BITSET_ALLOCA(bitset_len, use_heap)
	zend.ZendBitsetClear(bitset, bitset_len)
	i = num_req
	for i != 0 {
		var randval = PhpMtRandRange(0, num_avail-1)
		if zend.ZendBitsetIn(bitset, randval) == 0 {
			zend.ZendBitsetIncl(bitset, randval)
			i--
		}
	}

	/* i = 0; */

	fillScope := types.PackedFillStart(return_value.Array())

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		if (zend.ZendBitsetIn(bitset, i) ^ negative_bitset) != 0 {
			if string_key != nil {
				fillScope.FillSetStringCopy(string_key)
			} else {
				fillScope.FillSetLong(num_key)
			}
			fillScope.FillNext()
		}
		i++
	}

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	fillScope.FillEnd()
	zend.FreeAlloca(bitset, use_heap)
}
func ZifArraySum(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			input = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetLong(0)
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		zend.FastAddFunction(return_value, return_value, &entry_n)
	}
}
func ZifArrayProduct(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
	var dval float64
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			input = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetLong(1)
	if !(input.Array().Len()) {
		return
	}
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		if entry_n.IsType(types.IS_LONG) && return_value.IsType(types.IS_LONG) {
			dval = float64(return_value.Long() * float64(entry_n.Long()))
			if float64(zend.ZEND_LONG_MIN <= dval && dval <= float64(zend.ZEND_LONG_MAX)) {
				return_value.SetLong(return_value.Long() * entry_n.Long())
				continue
			}
		}
		zend.ConvertToDouble(return_value)
		zend.ConvertToDouble(&entry_n)
		return_value.SetDouble(return_value.Double() * entry_n.Double())
	}
}
func ZifArrayReduce(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, callback *types.Zval, _ zpp.Opt, initial *types.Zval) {
	var input *types.Zval
	var args []types.Zval
	var operand *types.Zval
	var result types.Zval
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var initial *types.Zval = nil
	var htbl *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			input = fp.ParseArray()
			fp.ParseFunc(&fci, &fci_cache)
			fp.StartOptional()
			initial = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() > 2 {
		types.ZVAL_COPY(&result, initial)
	} else {
		result.SetNull()
	}

	/* (zval **)input points to an element of argument stack
	 * the base pointer of which is subject to change.
	 * thus we need to keep the pointer to the hashtable for safety */

	htbl = input.Array()
	if htbl.Len() == 0 {
		types.ZVAL_COPY_VALUE(return_value, &result)
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	fci.SetRetval(&retval)
	fci.SetParamCount(2)
	fci.SetNoSeparation(0)
	var __ht = htbl
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		operand = _z
		types.ZVAL_COPY_VALUE(&args[0], &result)
		types.ZVAL_COPY(&args[1], operand)
		fci.SetParams(args)
		if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
			// zend.ZvalPtrDtor(&args[1])
			// zend.ZvalPtrDtor(&args[0])
			types.ZVAL_COPY_VALUE(&result, &retval)
		} else {
			// zend.ZvalPtrDtor(&args[1])
			// zend.ZvalPtrDtor(&args[0])
			return
		}
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
	zend.ZVAL_ZVAL(return_value, &result, 1, 1)
}
func ZifArrayFilter(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, _ zpp.Opt, callback *types.Zval, useKeys *types.Zval) {
	var array *types.Zval
	var operand *types.Zval
	var key *types.Zval
	var args []types.Zval
	var retval types.Zval
	var have_callback = 0
	var use_type = 0
	var string_key *types.String
	var fci = zend.EmptyFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var num_key zend.ZendUlong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			array = fp.ParseArray()
			fp.StartOptional()
			fp.ParseFunc(&fci, &fci_cache)
			use_type = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInit(return_value)
	if array.Array().Len() == 0 {
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	if executeData.NumArgs() > 1 {
		have_callback = 1
		fci.SetNoSeparation(0)
		fci.SetRetval(&retval)
		if use_type == ARRAY_FILTER_USE_BOTH {
			fci.SetParamCount(2)
			key = &args[1]
		} else {
			fci.SetParamCount(1)
			key = &args[0]
		}
	}
	var __ht = array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		operand = _z
		if have_callback != 0 {
			if use_type != 0 {

				/* Set up the key */

				if string_key == nil {
					key.SetLong(num_key)
				} else {
					key.SetStringCopy(string_key)
				}

				/* Set up the key */

			}
			if use_type != ARRAY_FILTER_USE_KEY {
				types.ZVAL_COPY(&args[0], operand)
			}
			fci.SetParams(args)
			if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS {
				var retval_true int
				// zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					// zend.ZvalPtrDtor(&args[1])
				}
				retval_true = zend.ZendIsTrue(&retval)
				// zend.ZvalPtrDtor(&retval)
				if retval_true == 0 {
					continue
				}
			} else {
				// zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					// zend.ZvalPtrDtor(&args[1])
				}
				return
			}
		} else if zend.ZendIsTrue(operand) == 0 {
			continue
		}
		if string_key != nil {
			operand = return_value.Array().KeyUpdate(string_key.GetStr(), operand)
		} else {
			operand = return_value.Array().IndexUpdate(num_key, operand)
		}
		zend.ZvalAddRef(operand)
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
}
func ZifArrayMap(executeData zpp.Ex, return_value zpp.Ret, callback *types.Zval, arrays []*types.Zval) {
	var arrays *types.Zval = nil
	var n_arrays = 0
	var result types.Zval
	var fci = zend.EmptyFcallInfo
	var fci_cache = zend.EmptyFcallInfoCache
	var i int
	var k uint32
	var maxlen uint32 = 0
	for {
		var _flags = 0
		var _min_num_args = 2
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.ParseFuncEx(&fci, &fci_cache, true, false)
			arrays, n_arrays = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetNull()
	if n_arrays == 1 {
		var num_key zend.ZendUlong
		var str_key *types.String
		var zv *types.Zval
		var arg types.Zval
		var ret int
		if arrays[0].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter 2 to be an array, %s given", types.ZendZvalTypeName(&arrays[0]))
			return
		}
		maxlen = arrays[0].Array().Len()

		/* Short-circuit: if no callback and only one array, just return it. */

		if !(zend.ZEND_FCI_INITIALIZED(fci)) || maxlen == 0 {
			types.ZVAL_COPY(return_value, &arrays[0])
			zend.ZendReleaseFcallInfoCache(&fci_cache)
			return
		}
		zend.ArrayInitSize(return_value, maxlen)
		var __ht = arrays[0].Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			fci.SetRetval(&result)
			fci.SetParamCount(1)
			fci.SetParams(&arg)
			fci.SetNoSeparation(0)
			types.ZVAL_COPY(&arg, zv)
			ret = zend.ZendCallFunction(&fci, &fci_cache)
			// zend.IZvalPtrDtor(&arg)
			if ret != types.SUCCESS || result.IsUndef() {
				return_value.Array().DestroyEx()
				return_value.SetNull()
				return
			}
			if str_key != nil {
				zend._zendHashAppend(return_value.Array(), str_key, &result)
			} else {
				return_value.Array().IndexAddNew(num_key, &result)
			}
		}
		zend.ZendReleaseFcallInfoCache(&fci_cache)
	} else {
		var array_pos = (*types.ArrayPosition)(zend.Ecalloc(n_arrays, b.SizeOf("ArrayPosition")))
		for i = 0; i < n_arrays; i++ {
			if arrays[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Expected parameter %d to be an array, %s given", i+2, types.ZendZvalTypeName(&arrays[i]))
				zend.Efree(array_pos)
				return
			}
			if arrays[i].Array().Len() > maxlen {
				maxlen = arrays[i].Array().Len()
			}
		}
		zend.ArrayInitSize(return_value, maxlen)
		if !(zend.ZEND_FCI_INITIALIZED(fci)) {
			var zv types.Zval

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {

				/* If no callback, the result will be an array, consisting of current
				 * entries from all arrays. */

				zend.ArrayInitSize(&result, n_arrays)
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].Array().GetNNumUsed() {
							zv.SetNull()
							break
						} else if arrays[i].Array().GetArData()[pos].GetVal().IsNotUndef() {
							types.ZVAL_COPY(&zv, arrays[i].Array().GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
					result.Array().NextIndexInsertNew(&zv)
				}
				return_value.Array().NextIndexInsertNew(&result)
			}

			/* We iterate through all the arrays at once. */

		} else {
			var params = (*types.Zval)(zend.SafeEmalloc(n_arrays, b.SizeOf("zval"), 0))

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= arrays[i].Array().GetNNumUsed() {
							params[i].SetNull()
							break
						} else if arrays[i].Array().GetArData()[pos].GetVal().IsNotUndef() {
							types.ZVAL_COPY(&params[i], arrays[i].Array().GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
				}
				fci.SetRetval(&result)
				fci.SetParamCount(n_arrays)
				fci.SetParams(params)
				fci.SetNoSeparation(0)
				if zend.ZendCallFunction(&fci, &fci_cache) != types.SUCCESS || result.IsUndef() {
					zend.Efree(array_pos)
					return_value.Array().DestroyEx()
					for i = 0; i < n_arrays; i++ {
						// zend.ZvalPtrDtor(&params[i])
					}
					zend.Efree(params)
					return_value.SetNull()
					return
				} else {
					for i = 0; i < n_arrays; i++ {
						// zend.ZvalPtrDtor(&params[i])
					}
				}
				return_value.Array().NextIndexInsertNew(&result)
			}
			zend.Efree(params)
			zend.ZendReleaseFcallInfoCache(&fci_cache)
		}
		zend.Efree(array_pos)
	}
}

//@zif -alias key_exists
func ZifArrayKeyExists(executeData zpp.Ex, return_value zpp.Ret, key *types.Zval, search *types.Zval) {
	var key *types.Zval
	var array *types.Zval
	var ht *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			key = fp.ParseZval()
			array = fp.ParseArrayOrObject()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if array.IsType(types.IS_ARRAY) {
		ht = array.Array()
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
	}
	switch key.GetType() {
	case types.IS_STRING:
		return_value.SetBool(ht.SymtableExistsInd(key.String().GetStr()))
	case types.IS_LONG:
		return_value.SetBool(ht.IndexExists(key.Long()))
	case types.IS_NULL:
		return_value.SetBool(ht.KeyExistsIndirect(types.NewString("").GetStr()))
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "The first argument should be either a string or an integer")
		return_value.SetFalse()
	}
	if array.GetType() != types.IS_ARRAY {
		zend.ZendReleaseProperties(ht)
	}
}
func ZifArrayChunk(executeData zpp.Ex, return_value zpp.Ret, arg *types.Zval, size *types.Zval, _ zpp.Opt, preserveKeys *types.Zval) {
	var num_in int
	var size zend.ZendLong
	var current = 0
	var str_key *types.String
	var num_key zend.ZendUlong
	var preserve_keys = 0
	var input *types.Zval = nil
	var chunk types.Zval
	var entry *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			input = fp.ParseArray()
			size = fp.ParseLong()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Do bounds checking for size parameter. */

	if size < 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Size parameter expected to be greater than 0")
		return
	}
	num_in = input.Array().Len()
	if size > num_in {
		if num_in > 0 {
			size = num_in
		} else {
			size = 1
		}
	}
	zend.ArrayInitSize(return_value, uint32((num_in-1)/size+1))
	chunk.SetUndef()
	var __ht = input.Array()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		num_key = _p.GetH()
		str_key = _p.GetKey()
		entry = _z

		/* If new chunk, create and initialize it. */

		if chunk.IsUndef() {
			zend.ArrayInitSize(&chunk, uint32(size))
		}

		/* Add entry to the chunk, preserving keys if necessary. */

		if preserve_keys != 0 {
			if str_key != nil {
				entry = chunk.Array().KeyUpdate(str_key.GetStr(), entry)
			} else {
				entry = chunk.Array().IndexUpdate(num_key, entry)
			}
		} else {
			entry = chunk.Array().NextIndexInsert(entry)
		}
		zend.ZvalAddRef(entry)

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

		if b.PreInc(&current)%size == 0 {
			zend.AddNextIndexZval(return_value, &chunk)
			chunk.SetUndef()
		}

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

	}

	/* Add the final chunk if there is one. */

	if chunk.IsNotUndef() {
		zend.AddNextIndexZval(return_value, &chunk)
	}

	/* Add the final chunk if there is one. */
}
func ZifArrayCombine(executeData zpp.Ex, return_value zpp.Ret, keys *types.Zval, values *types.Zval) {
	var values *types.Array
	var keys *types.Array
	var pos_values uint32 = 0
	var entry_keys *types.Zval
	var entry_values *types.Zval
	var num_keys int
	var num_values int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			keys = fp.ParseArrayHt()
			values = fp.ParseArrayHt()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	num_keys = keys.Len()
	num_values = values.Len()
	if num_keys != num_values {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Both parameters should have an equal number of elements")
		return_value.SetFalse()
		return
	}
	if num_keys == 0 {
		return_value.SetEmptyArray()
		return
	}
	zend.ArrayInitSize(return_value, num_keys)
	var __ht = keys
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		entry_keys = _z
		for true {
			if pos_values >= values.GetNNumUsed() {
				break
			} else if values.GetArData()[pos_values].GetVal().IsNotUndef() {
				entry_values = values.GetArData()[pos_values].GetVal()
				if entry_keys.IsType(types.IS_LONG) {
					entry_values = return_value.Array().IndexUpdate(entry_keys.Long(), entry_values)
				} else {
					var tmp_key *types.String
					var key = zend.ZvalGetTmpString(entry_keys, &tmp_key)
					entry_values = return_value.Array().SymtableUpdate(key.GetStr(), entry_values)
					// zend.ZendTmpStringRelease(tmp_key)
				}
				zend.ZvalAddRef(entry_values)
				pos_values++
				break
			}
			pos_values++
		}
	}
}

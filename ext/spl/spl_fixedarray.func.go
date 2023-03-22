// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func SplFixedArrayFromObj(obj *types.ZendObject) *SplFixedarrayObject {
	return (*SplFixedarrayObject)((*byte)(obj - zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLFIXEDARRAY_P(zv *types.Zval) *SplFixedarrayObject { return SplFixedArrayFromObj(zv.GetObj()) }
func SplFixedarrayInit(array *SplFixedarray, size zend.ZendLong) {
	if size > 0 {
		array.SetSize(0)
		array.SetElements(zend.Ecalloc(size, b.SizeOf("zval")))
		array.SetSize(size)
	} else {
		array.SetElements(nil)
		array.SetSize(0)
	}
}
func SplFixedarrayResize(array *SplFixedarray, size zend.ZendLong) {
	if size == array.GetSize() {

		/* nothing to do */

		return

		/* nothing to do */

	}

	/* first initialization */

	if array.GetSize() == 0 {
		SplFixedarrayInit(array, size)
		return
	}

	/* clearing the array */

	if size == 0 {
		if array.GetElements() != nil {
			var i zend.ZendLong
			var elements *types.Zval = array.GetElements()
			var old_size zend.ZendLong = array.GetSize()
			array.SetElements(nil)
			array.SetSize(0)
			for i = 0; i < old_size; i++ {
				zend.ZvalPtrDtor(&elements[i])
			}
			zend.Efree(elements)
			return
		}
	} else if size > array.GetSize() {
		array.SetElements(zend.SafeErealloc(array.GetElements(), size, b.SizeOf("zval"), 0))
		memset(array.GetElements()+array.GetSize(), '0', b.SizeOf("zval")*(size-array.GetSize()))
	} else {
		var i zend.ZendLong
		for i = size; i < array.GetSize(); i++ {
			zend.ZvalPtrDtor(&array.GetElements()[i])
		}
		array.SetElements(zend.Erealloc(array.GetElements(), b.SizeOf("zval")*size))
	}
	array.SetSize(size)
}
func SplFixedarrayCopy(to *SplFixedarray, from *SplFixedarray) {
	var i int
	for i = 0; i < from.GetSize(); i++ {
		types.ZVAL_COPY(to.GetElements()[i], from.GetElements()[i])
	}
}
func SplFixedarrayObjectGetGc(obj *types.Zval, table **types.Zval, n *int) *types.Array {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *types.Array = zend.ZendStdGetProperties(obj)
	*table = intern.GetArray().GetElements()
	*n = int(intern.GetArray().GetSize())
	return ht
}
func SplFixedarrayObjectGetProperties(obj *types.Zval) *types.Array {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *types.Array = zend.ZendStdGetProperties(obj)
	var i zend.ZendLong = 0
	if intern.GetArray().GetSize() > 0 {
		var j zend.ZendLong = ht.GetNNumOfElements()
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			if !(intern.GetArray().GetElements()[i].IsUndef()) {
				ht.IndexUpdateH(i, intern.GetArray().GetElements()[i])
				intern.GetArray().GetElements()[i].TryAddRefcount()
			} else {
				ht.IndexUpdateH(i, zend.EG__().GetUninitializedZval())
			}
		}
		if j > intern.GetArray().GetSize() {
			for i = intern.GetArray().GetSize(); i < j; i++ {
				types.ZendHashIndexDel(ht, i)
			}
		}
	}
	return ht
}
func SplFixedarrayObjectFreeStorage(object *types.ZendObject) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(object)
	var i zend.ZendLong
	if intern.GetArray().GetSize() > 0 {
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			zend.ZvalPtrDtor(&intern.GetArray().GetElements()[i])
		}
		if intern.GetArray().GetSize() > 0 && intern.GetArray().GetElements() != nil {
			zend.Efree(intern.GetArray().GetElements())
		}
	}
	zend.ZendObjectStdDtor(intern.GetStd())
}
func SplFixedarrayObjectNewEx(class_type *types.ClassEntry, orig *types.Zval, clone_orig int) *types.ZendObject {
	var intern *SplFixedarrayObject
	var parent *types.ClassEntry = class_type
	var inherited int = 0
	var funcs_ptr *zend.ZendClassIteratorFuncs
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_fixedarray_object"), parent)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	intern.SetCurrent(0)
	intern.SetFlags(0)
	if orig != nil && clone_orig != 0 {
		var other *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(orig)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		SplFixedarrayInit(intern.GetArray(), other.GetArray().GetSize())
		SplFixedarrayCopy(intern.GetArray(), other.GetArray())
	}
	for parent != nil {
		if parent == spl_ce_SplFixedArray {
			intern.GetStd().SetHandlers(&spl_handler_SplFixedArray)
			break
		}
		parent = parent.GetParent()
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, faults.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplFixedArray")
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if funcs_ptr.GetZfCurrent() == nil {
		funcs_ptr.SetZfRewind(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "rewind"))
		funcs_ptr.SetZfValid(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "valid"))
		funcs_ptr.SetZfKey(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "key"))
		funcs_ptr.SetZfCurrent(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "current"))
		funcs_ptr.SetZfNext(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "next"))
	}
	if inherited != 0 {
		if funcs_ptr.GetZfRewind().GetScope() != parent {
			intern.SetIsRewind(true)
		}
		if funcs_ptr.GetZfValid().GetScope() != parent {
			intern.SetIsValid(true)
		}
		if funcs_ptr.GetZfKey().GetScope() != parent {
			intern.SetIsKey(true)
		}
		if funcs_ptr.GetZfCurrent().GetScope() != parent {
			intern.SetIsCurrent(true)
		}
		if funcs_ptr.GetZfNext().GetScope() != parent {
			intern.SetIsNext(true)
		}
		intern.SetFptrOffsetGet(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetget"))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetset"))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetexists"))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetunset"))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(types.ZendHashStrFindPtr(class_type.GetFunctionTable(), "count"))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}
	return intern.GetStd()
}
func SplFixedarrayNew(class_type *types.ClassEntry) *types.ZendObject {
	return SplFixedarrayObjectNewEx(class_type, nil, 0)
}
func SplFixedarrayObjectClone(zobject *types.Zval) *types.ZendObject {
	var old_object *types.ZendObject
	var new_object *types.ZendObject
	old_object = zobject.GetObj()
	new_object = SplFixedarrayObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplFixedarrayObjectReadDimensionHelper(intern *SplFixedarrayObject, offset *types.Zval) *types.Zval {
	var index zend.ZendLong

	/* we have to return NULL on error here to avoid memleak because of
	 * ZE duplicating uninitialized_zval_ptr */

	if offset == nil {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	}
	if offset.GetType() != types.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.GetLval()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	} else if intern.GetArray().GetElements()[index].IsUndef() {
		return nil
	} else {
		return intern.GetArray().GetElements()[index]
	}
}
func SplFixedarrayObjectReadDimension(object *types.Zval, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if type_ == zend.BP_VAR_IS && SplFixedarrayObjectHasDimension(object, offset, 0) == 0 {
		return zend.EG__().GetUninitializedZval()
	}
	if intern.GetFptrOffsetGet() != nil {
		var tmp types.Zval
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			types.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetGet(), "offsetGet", rv, offset)
		zend.ZvalPtrDtor(offset)
		if !(rv.IsUndef()) {
			return rv
		}
		return zend.EG__().GetUninitializedZval()
	}
	return SplFixedarrayObjectReadDimensionHelper(intern, offset)
}
func SplFixedarrayObjectWriteDimensionHelper(intern *SplFixedarrayObject, offset *types.Zval, value *types.Zval) {
	var index zend.ZendLong
	if offset == nil {

		/* '$array[] = value' syntax is not supported */

		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	}
	if offset.GetType() != types.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.GetLval()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {

		/* Fix #81429 */

		var ptr *types.Zval = &intern.GetArray().GetElements()[index]
		var tmp types.Zval
		types.ZVAL_COPY_VALUE(&tmp, ptr)
		types.ZVAL_COPY_DEREF(ptr, value)
		zend.ZvalPtrDtor(&tmp)
	}
}
func SplFixedarrayObjectWriteDimension(object *types.Zval, offset *types.Zval, value *types.Zval) {
	var intern *SplFixedarrayObject
	var tmp types.Zval
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetSet() != nil {
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			types.SEPARATE_ARG_IF_REF(offset)
		}
		types.SEPARATE_ARG_IF_REF(value)
		zend.ZendCallMethodWith2Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		zend.ZvalPtrDtor(value)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectWriteDimensionHelper(intern, offset, value)
}
func SplFixedarrayObjectUnsetDimensionHelper(intern *SplFixedarrayObject, offset *types.Zval) {
	var index zend.ZendLong
	if offset.GetType() != types.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.GetLval()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {
		zend.ZvalPtrDtor(&intern.GetArray().GetElements()[index])
		intern.GetArray().GetElements()[index].SetUndef()
	}
}
func SplFixedarrayObjectUnsetDimension(object *types.Zval, offset *types.Zval) {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetDel() != nil {
		types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectUnsetDimensionHelper(intern, offset)
}
func SplFixedarrayObjectHasDimensionHelper(intern *SplFixedarrayObject, offset *types.Zval, check_empty int) int {
	var index zend.ZendLong
	var retval int
	if offset.GetType() != types.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.GetLval()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		retval = 0
	} else {
		if intern.GetArray().GetElements()[index].IsUndef() {
			retval = 0
		} else if check_empty != 0 {
			if zend.ZendIsTrue(intern.GetArray().GetElements()[index]) != 0 {
				retval = 1
			} else {
				retval = 0
			}
		} else {
			retval = 1
		}
	}
	return retval
}
func SplFixedarrayObjectHasDimension(object *types.Zval, offset *types.Zval, check_empty int) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetHas() != nil {
		var rv types.Zval
		var result types.ZendBool
		types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		zend.ZvalPtrDtor(offset)
		result = zend.ZendIsTrue(&rv)
		zend.ZvalPtrDtor(&rv)
		return result
	}
	return SplFixedarrayObjectHasDimensionHelper(intern, offset, check_empty)
}
func SplFixedarrayObjectCountElements(object *types.Zval, count *zend.ZendLong) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if !(rv.IsUndef()) {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
		} else {
			*count = 0
		}
	} else {
		*count = intern.GetArray().GetSize()
	}
	return types.SUCCESS
}
func zim_spl_SplFixedArray___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplFixedarrayObject
	var size zend.ZendLong = 0
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "|l", &size) == types.FAILURE {
		return
	}
	if size < 0 {
		faults.ThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetArray().GetSize() > 0 {

		/* called __construct() twice, bail out */

		return

		/* called __construct() twice, bail out */

	}
	SplFixedarrayInit(intern.GetArray(), size)
}
func zim_spl_SplFixedArray___wakeup(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	var intern_ht *types.Array = zend.ZendStdGetProperties(zend.ZEND_THIS(executeData))
	var data *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if intern.GetArray().GetSize() == 0 {
		var index int = 0
		var size int = intern_ht.GetNNumOfElements()
		SplFixedarrayInit(intern.GetArray(), size)
		var __ht *types.Array = intern_ht
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			data = _z
			types.ZVAL_COPY(intern.GetArray().GetElements()[index], data)
			index++
		}

		/* Remove the unserialised properties, since we now have the elements
		 * within the spl_fixedarray_object structure. */

		intern_ht.Clean()

		/* Remove the unserialised properties, since we now have the elements
		 * within the spl_fixedarray_object structure. */

	}
}
func zim_spl_SplFixedArray_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplFixedarrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	return_value.SetLong(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_toArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if intern.GetArray().GetSize() > 0 {
		var i int = 0
		zend.ArrayInit(return_value)
		for ; i < intern.GetArray().GetSize(); i++ {
			if !(intern.GetArray().GetElements()[i].IsUndef()) {
				return_value.GetArr().IndexUpdateH(i, intern.GetArray().GetElements()[i])
				intern.GetArray().GetElements()[i].TryAddRefcount()
			} else {
				return_value.GetArr().IndexUpdateH(i, zend.EG__().GetUninitializedZval())
			}
		}
	} else {
		types.ZVAL_EMPTY_ARRAY(return_value)
		return
	}
}
func zim_spl_SplFixedArray_fromArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var data *types.Zval
	var array SplFixedarray
	var intern *SplFixedarrayObject
	var num int
	var save_indexes types.ZendBool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "a|b", &data, &save_indexes) == types.FAILURE {
		return
	}
	num = types.Z_ARRVAL_P(data).GetNNumOfElements()
	if num > 0 && save_indexes != 0 {
		var element *types.Zval
		var str_index *types.String
		var num_index zend.ZendUlong
		var max_index zend.ZendUlong = 0
		var tmp zend.ZendLong
		var __ht *types.Array = data.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			num_index = _p.GetH()
			str_index = _p.GetKey()
			if str_index != nil || zend.ZendLong(num_index < 0) != 0 {
				faults.ThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array must contain only positive integer keys")
				return
			}
			if num_index > max_index {
				max_index = num_index
			}
		}
		tmp = max_index + 1
		if tmp <= 0 {
			faults.ThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "integer overflow detected")
			return
		}
		SplFixedarrayInit(&array, tmp)
		var __ht__1 *types.Array = data.GetArr()
		for _, _p := range __ht__1.foreachData() {
			var _z *types.Zval = _p.GetVal()

			num_index = _p.GetH()
			str_index = _p.GetKey()
			element = _z
			types.ZVAL_COPY_DEREF(array.GetElements()[num_index], element)
		}
	} else if num > 0 && save_indexes == 0 {
		var element *types.Zval
		var i zend.ZendLong = 0
		SplFixedarrayInit(&array, num)
		var __ht *types.Array = data.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			element = _z
			types.ZVAL_COPY_DEREF(array.GetElements()[i], element)
			i++
		}
	} else {
		SplFixedarrayInit(&array, 0)
	}
	zend.ObjectInitEx(return_value, spl_ce_SplFixedArray)
	intern = Z_SPLFIXEDARRAY_P(return_value)
	intern.SetArray(array)
}
func zim_spl_SplFixedArray_getSize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplFixedarrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	return_value.SetLong(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_setSize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplFixedarrayObject
	var size zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &size) == types.FAILURE {
		return
	}
	if size < 0 {
		faults.ThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	SplFixedarrayResize(intern.GetArray(), size)
	return_value.SetTrue()
	return
}
func zim_spl_SplFixedArray_offsetExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	types.ZVAL_BOOL(return_value, SplFixedarrayObjectHasDimensionHelper(intern, zindex, 0) != 0)
	return
}
func zim_spl_SplFixedArray_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var value *types.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	value = SplFixedarrayObjectReadDimensionHelper(intern, zindex)
	if value != nil {
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_SplFixedArray_offsetSet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var value *types.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &zindex, &value) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	SplFixedarrayObjectWriteDimensionHelper(intern, zindex, value)
}
func zim_spl_SplFixedArray_offsetUnset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	SplFixedarrayObjectUnsetDimensionHelper(intern, zindex)
}
func SplFixedarrayItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFixedarrayIt = (*SplFixedarrayIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(iterator.GetIntern().GetIt().GetData())
}
func SplFixedarrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsRewind() {
		zend.ZendUserItRewind(iter)
	} else {
		object.SetCurrent(0)
	}
}
func SplFixedarrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsValid() {
		return zend.ZendUserItValid(iter)
	}
	if object.GetCurrent() >= 0 && object.GetCurrent() < object.GetArray().GetSize() {
		return types.SUCCESS
	}
	return types.FAILURE
}
func SplFixedarrayItGetCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var zindex types.Zval
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsCurrent() {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *types.Zval
		zindex.SetLong(object.GetCurrent())
		data = SplFixedarrayObjectReadDimensionHelper(object, &zindex)
		if data == nil {
			data = zend.EG__().GetUninitializedZval()
		}
		return data
	}
}
func SplFixedarrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsKey() {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		key.SetLong(object.GetCurrent())
	}
}
func SplFixedarrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsNext() {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		object.GetCurrent()++
	}
}
func zim_spl_SplFixedArray_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(intern.GetCurrent())
	return
}
func zim_spl_SplFixedArray_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern.GetCurrent()++
}
func zim_spl_SplFixedArray_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZVAL_BOOL(return_value, intern.GetCurrent() >= 0 && intern.GetCurrent() < intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern.SetCurrent(0)
}
func zim_spl_SplFixedArray_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex types.Zval
	var value *types.Zval
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zindex.SetLong(intern.GetCurrent())
	value = SplFixedarrayObjectReadDimensionHelper(intern, &zindex)
	if value != nil {
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func SplFixedarrayGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFixedarrayIt
	if by_ref != 0 {
		faults.ThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_fixedarray_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	object.AddRefcount()
	iterator.GetIntern().GetIt().GetData().SetObject(object.GetObj())
	iterator.GetIntern().GetIt().SetFuncs(&SplFixedarrayItFuncs)
	iterator.GetIntern().SetCe(ce)
	iterator.GetIntern().GetValue().SetUndef()
	return iterator.GetIntern().GetIt()
}
func ZmStartupSplFixedarray(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplFixedArray, "SplFixedArray", SplFixedarrayNew, spl_funcs_SplFixedArray)
	memcpy(&spl_handler_SplFixedArray, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplFixedArray.SetOffset(zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_SplFixedArray.SetCloneObj(SplFixedarrayObjectClone)
	spl_handler_SplFixedArray.SetReadDimension(SplFixedarrayObjectReadDimension)
	spl_handler_SplFixedArray.SetWriteDimension(SplFixedarrayObjectWriteDimension)
	spl_handler_SplFixedArray.SetUnsetDimension(SplFixedarrayObjectUnsetDimension)
	spl_handler_SplFixedArray.SetHasDimension(SplFixedarrayObjectHasDimension)
	spl_handler_SplFixedArray.SetCountElements(SplFixedarrayObjectCountElements)
	spl_handler_SplFixedArray.SetGetProperties(SplFixedarrayObjectGetProperties)
	spl_handler_SplFixedArray.SetGetGc(SplFixedarrayObjectGetGc)
	spl_handler_SplFixedArray.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_SplFixedArray.SetFreeObj(SplFixedarrayObjectFreeStorage)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_Countable)
	spl_ce_SplFixedArray.SetGetIterator(SplFixedarrayGetIterator)
	spl_ce_SplFixedArray.AddCeFlags(zend.ZEND_ACC_REUSE_GET_ITERATOR)
	return types.SUCCESS
}

// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func SplFixedArrayFromObj(obj *zend.ZendObject) *SplFixedarrayObject {
	return (*SplFixedarrayObject)((*byte)(obj - zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLFIXEDARRAY_P(zv *zend.Zval) *SplFixedarrayObject {
	return SplFixedArrayFromObj(zend.Z_OBJ_P(zv))
}
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
			var elements *zend.Zval = array.GetElements()
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
		zend.ZVAL_COPY(to.GetElements()[i], from.GetElements()[i])
	}
}
func SplFixedarrayObjectGetGc(obj *zend.Zval, table **zend.Zval, n *int) *zend.HashTable {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	*table = intern.GetArray().GetElements()
	*n = int(intern.GetArray().GetSize())
	return ht
}
func SplFixedarrayObjectGetProperties(obj *zend.Zval) *zend.HashTable {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	var i zend.ZendLong = 0
	if intern.GetArray().GetSize() > 0 {
		var j zend.ZendLong = ht.GetNNumOfElements()
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			if !(zend.Z_ISUNDEF(intern.GetArray().GetElements()[i])) {
				zend.ZendHashIndexUpdate(ht, i, intern.GetArray().GetElements()[i])
				zend.Z_TRY_ADDREF(intern.GetArray().GetElements()[i])
			} else {
				zend.ZendHashIndexUpdate(ht, i, &(zend.ExecutorGlobals.GetUninitializedZval()))
			}
		}
		if j > intern.GetArray().GetSize() {
			for i = intern.GetArray().GetSize(); i < j; i++ {
				zend.ZendHashIndexDel(ht, i)
			}
		}
	}
	return ht
}
func SplFixedarrayObjectFreeStorage(object *zend.ZendObject) {
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
func SplFixedarrayObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplFixedarrayObject
	var parent *zend.ZendClassEntry = class_type
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
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplFixedArray")
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if funcs_ptr.GetZfCurrent() == nil {
		funcs_ptr.SetZfRewind(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "rewind", b.SizeOf("\"rewind\"")-1))
		funcs_ptr.SetZfValid(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "valid", b.SizeOf("\"valid\"")-1))
		funcs_ptr.SetZfKey(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "key", b.SizeOf("\"key\"")-1))
		funcs_ptr.SetZfCurrent(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "current", b.SizeOf("\"current\"")-1))
		funcs_ptr.SetZfNext(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "next", b.SizeOf("\"next\"")-1))
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
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetget", b.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetset", b.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetexists", b.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetunset", b.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}
	return intern.GetStd()
}
func SplFixedarrayNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplFixedarrayObjectNewEx(class_type, nil, 0)
}
func SplFixedarrayObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zend.Z_OBJ_P(zobject)
	new_object = SplFixedarrayObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplFixedarrayObjectReadDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval) *zend.Zval {
	var index zend.ZendLong

	/* we have to return NULL on error here to avoid memleak because of
	 * ZE duplicating uninitialized_zval_ptr */

	if offset == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	}
	if offset.GetType() != zend.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = zend.Z_LVAL_P(offset)
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	} else if zend.Z_ISUNDEF(intern.GetArray().GetElements()[index]) {
		return nil
	} else {
		return intern.GetArray().GetElements()[index]
	}
}
func SplFixedarrayObjectReadDimension(object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if type_ == zend.BP_VAR_IS && SplFixedarrayObjectHasDimension(object, offset, 0) == 0 {
		return &(zend.ExecutorGlobals.GetUninitializedZval())
	}
	if intern.GetFptrOffsetGet() != nil {
		var tmp zend.Zval
		if offset == nil {
			zend.ZVAL_NULL(&tmp)
			offset = &tmp
		} else {
			zend.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetGet(), "offsetGet", rv, offset)
		zend.ZvalPtrDtor(offset)
		if !(zend.Z_ISUNDEF_P(rv)) {
			return rv
		}
		return &(zend.ExecutorGlobals.GetUninitializedZval())
	}
	return SplFixedarrayObjectReadDimensionHelper(intern, offset)
}
func SplFixedarrayObjectWriteDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval, value *zend.Zval) {
	var index zend.ZendLong
	if offset == nil {

		/* '$array[] = value' syntax is not supported */

		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	}
	if offset.GetType() != zend.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = zend.Z_LVAL_P(offset)
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {

		/* Fix #81429 */

		var ptr *zend.Zval = &intern.GetArray().GetElements()[index]
		var tmp zend.Zval
		zend.ZVAL_COPY_VALUE(&tmp, ptr)
		zend.ZVAL_COPY_DEREF(ptr, value)
		zend.ZvalPtrDtor(&tmp)
	}
}
func SplFixedarrayObjectWriteDimension(object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	var intern *SplFixedarrayObject
	var tmp zend.Zval
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetSet() != nil {
		if offset == nil {
			zend.ZVAL_NULL(&tmp)
			offset = &tmp
		} else {
			zend.SEPARATE_ARG_IF_REF(offset)
		}
		zend.SEPARATE_ARG_IF_REF(value)
		zend.ZendCallMethodWith2Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		zend.ZvalPtrDtor(value)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectWriteDimensionHelper(intern, offset, value)
}
func SplFixedarrayObjectUnsetDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval) {
	var index zend.ZendLong
	if offset.GetType() != zend.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = zend.Z_LVAL_P(offset)
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {
		zend.ZvalPtrDtor(&intern.GetArray().GetElements()[index])
		zend.ZVAL_UNDEF(intern.GetArray().GetElements()[index])
	}
}
func SplFixedarrayObjectUnsetDimension(object *zend.Zval, offset *zend.Zval) {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetDel() != nil {
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectUnsetDimensionHelper(intern, offset)
}
func SplFixedarrayObjectHasDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval, check_empty int) int {
	var index zend.ZendLong
	var retval int
	if offset.GetType() != zend.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = zend.Z_LVAL_P(offset)
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		retval = 0
	} else {
		if zend.Z_ISUNDEF(intern.GetArray().GetElements()[index]) {
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
func SplFixedarrayObjectHasDimension(object *zend.Zval, offset *zend.Zval, check_empty int) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetHas() != nil {
		var rv zend.Zval
		var result zend.ZendBool
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		zend.ZvalPtrDtor(offset)
		result = zend.ZendIsTrue(&rv)
		zend.ZvalPtrDtor(&rv)
		return result
	}
	return SplFixedarrayObjectHasDimensionHelper(intern, offset, check_empty)
}
func SplFixedarrayObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if !(zend.Z_ISUNDEF(rv)) {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
		} else {
			*count = 0
		}
	} else {
		*count = intern.GetArray().GetSize()
	}
	return zend.SUCCESS
}
func zim_spl_SplFixedArray___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplFixedarrayObject
	var size zend.ZendLong = 0
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "|l", &size) == zend.FAILURE {
		return
	}
	if size < 0 {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
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
func zim_spl_SplFixedArray___wakeup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	var intern_ht *zend.HashTable = zend.ZendStdGetProperties(zend.ZEND_THIS)
	var data *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetArray().GetSize() == 0 {
		var index int = 0
		var size int = intern_ht.GetNNumOfElements()
		SplFixedarrayInit(intern.GetArray(), size)
		for {
			var __ht *zend.HashTable = intern_ht
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.GetType() == zend.IS_UNDEF {
					continue
				}
				data = _z
				zend.ZVAL_COPY(intern.GetArray().GetElements()[index], data)
				index++
			}
			break
		}

		/* Remove the unserialised properties, since we now have the elements
		 * within the spl_fixedarray_object structure. */

		zend.ZendHashClean(intern_ht)

		/* Remove the unserialised properties, since we now have the elements
		 * within the spl_fixedarray_object structure. */

	}
}
func zim_spl_SplFixedArray_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplFixedarrayObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	zend.RETVAL_LONG(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_toArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if intern.GetArray().GetSize() > 0 {
		var i int = 0
		zend.ArrayInit(return_value)
		for ; i < intern.GetArray().GetSize(); i++ {
			if !(zend.Z_ISUNDEF(intern.GetArray().GetElements()[i])) {
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), i, intern.GetArray().GetElements()[i])
				zend.Z_TRY_ADDREF(intern.GetArray().GetElements()[i])
			} else {
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), i, &(zend.ExecutorGlobals.GetUninitializedZval()))
			}
		}
	} else {
		zend.RETVAL_EMPTY_ARRAY()
		return
	}
}
func zim_spl_SplFixedArray_fromArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var data *zend.Zval
	var array SplFixedarray
	var intern *SplFixedarrayObject
	var num int
	var save_indexes zend.ZendBool = 1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "a|b", &data, &save_indexes) == zend.FAILURE {
		return
	}
	num = zend.Z_ARRVAL_P(data).GetNNumOfElements()
	if num > 0 && save_indexes != 0 {
		var element *zend.Zval
		var str_index *zend.ZendString
		var num_index zend.ZendUlong
		var max_index zend.ZendUlong = 0
		var tmp zend.ZendLong
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.GetType() == zend.IS_UNDEF {
					continue
				}
				num_index = _p.GetH()
				str_index = _p.GetKey()
				if str_index != nil || zend.ZendLong(num_index < 0) != 0 {
					zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array must contain only positive integer keys")
					return
				}
				if num_index > max_index {
					max_index = num_index
				}
			}
			break
		}
		tmp = max_index + 1
		if tmp <= 0 {
			zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "integer overflow detected")
			return
		}
		SplFixedarrayInit(&array, tmp)
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.GetType() == zend.IS_UNDEF {
					continue
				}
				num_index = _p.GetH()
				str_index = _p.GetKey()
				element = _z
				zend.ZVAL_COPY_DEREF(array.GetElements()[num_index], element)
			}
			break
		}
	} else if num > 0 && save_indexes == 0 {
		var element *zend.Zval
		var i zend.ZendLong = 0
		SplFixedarrayInit(&array, num)
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.GetType() == zend.IS_UNDEF {
					continue
				}
				element = _z
				zend.ZVAL_COPY_DEREF(array.GetElements()[i], element)
				i++
			}
			break
		}
	} else {
		SplFixedarrayInit(&array, 0)
	}
	zend.ObjectInitEx(return_value, spl_ce_SplFixedArray)
	intern = Z_SPLFIXEDARRAY_P(return_value)
	intern.SetArray(array)
}
func zim_spl_SplFixedArray_getSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplFixedarrayObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	zend.RETVAL_LONG(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_setSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplFixedarrayObject
	var size zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &size) == zend.FAILURE {
		return
	}
	if size < 0 {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	SplFixedarrayResize(intern.GetArray(), size)
	zend.RETVAL_TRUE
	return
}
func zim_spl_SplFixedArray_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	zend.RETVAL_BOOL(SplFixedarrayObjectHasDimensionHelper(intern, zindex, 0) != 0)
	return
}
func zim_spl_SplFixedArray_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	value = SplFixedarrayObjectReadDimensionHelper(intern, zindex)
	if value != nil {
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func zim_spl_SplFixedArray_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	SplFixedarrayObjectWriteDimensionHelper(intern, zindex, value)
}
func zim_spl_SplFixedArray_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
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
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func SplFixedarrayItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var zindex zend.Zval
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsCurrent() {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *zend.Zval
		zend.ZVAL_LONG(&zindex, object.GetCurrent())
		data = SplFixedarrayObjectReadDimensionHelper(object, &zindex)
		if data == nil {
			data = &(zend.ExecutorGlobals.GetUninitializedZval())
		}
		return data
	}
}
func SplFixedarrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsKey() {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		zend.ZVAL_LONG(key, object.GetCurrent())
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
func zim_spl_SplFixedArray_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetCurrent())
	return
}
func zim_spl_SplFixedArray_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern.GetCurrent()++
}
func zim_spl_SplFixedArray_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(intern.GetCurrent() >= 0 && intern.GetCurrent() < intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern.SetCurrent(0)
}
func zim_spl_SplFixedArray_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZVAL_LONG(&zindex, intern.GetCurrent())
	value = SplFixedarrayObjectReadDimensionHelper(intern, &zindex)
	if value != nil {
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func SplFixedarrayGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFixedarrayIt
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_fixedarray_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.Z_ADDREF_P(object)
	zend.ZVAL_OBJ(iterator.GetIntern().GetIt().GetData(), zend.Z_OBJ_P(object))
	iterator.GetIntern().GetIt().SetFuncs(&SplFixedarrayItFuncs)
	iterator.GetIntern().SetCe(ce)
	zend.ZVAL_UNDEF(iterator.GetIntern().GetValue())
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
	return zend.SUCCESS
}

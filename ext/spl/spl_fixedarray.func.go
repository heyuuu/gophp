package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func SplFixedArrayFromObj(obj *types.Object) *SplFixedArrayObject {
	return (*SplFixedArrayObject)((*byte)(obj - zend_long((*byte)(&((*SplFixedArrayObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLFIXEDARRAY_P(zv *types.Zval) *SplFixedArrayObject {
	return SplFixedArrayFromObj(zv.Object())
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
			var elements *types.Zval = array.GetElements()
			var old_size zend.ZendLong = array.GetSize()
			array.SetElements(nil)
			array.SetSize(0)
			for i = 0; i < old_size; i++ {
				// zend.ZvalPtrDtor(&elements[i])
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
			// zend.ZvalPtrDtor(&array.GetElements()[i])
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
func SplFixedarrayObjectGetProperties(obj *types.Zval) *types.Array {
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *types.Array = zend.ZendStdGetProperties(obj)
	var i zend.ZendLong = 0
	if intern.GetArray().GetSize() > 0 {
		var j zend.ZendLong = ht.Len()
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			if !(intern.GetArray().GetElements()[i].IsUndef()) {
				ht.IndexUpdate(i, intern.GetArray().GetElements()[i])
				//intern.GetArray().GetElements()[i].TryAddRefcount()
			} else {
				ht.IndexUpdate(i, zend.UninitializedZval())
			}
		}
		if j > intern.GetArray().GetSize() {
			for i = intern.GetArray().GetSize(); i < j; i++ {
				ht.IndexDelete(i)
			}
		}
	}
	return ht
}
func SplFixedarrayObjectFreeStorage(object *types.Object) {
	var intern *SplFixedArrayObject = SplFixedArrayFromObj(object)
	var i zend.ZendLong
	if intern.GetArray().GetSize() > 0 {
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			// zend.ZvalPtrDtor(&intern.GetArray().GetElements()[i])
		}
		if intern.GetArray().GetSize() > 0 && intern.GetArray().GetElements() != nil {
			zend.Efree(intern.GetArray().GetElements())
		}
	}
	zend.ZendObjectStdDtor(intern.GetStd())
}
func SplFixedarrayObjectNewEx(class_type *types.ClassEntry, orig *types.Zval, clone_orig int) *types.Object {
	var intern *SplFixedArrayObject = NewSplFixedArrayObject(class_type)
	var parent *types.ClassEntry = class_type
	var inherited int = 0
	var funcs_ptr *zend.ZendClassIteratorFuncs
	if orig != nil && clone_orig != 0 {
		var other *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(orig)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		SplFixedarrayInit(intern.GetArray(), other.GetArray().GetSize())
		SplFixedarrayCopy(intern.GetArray(), other.GetArray())
	}
	for parent != nil {
		if parent == spl_ce_SplFixedArray {
			break
		}
		parent = parent.GetParent()
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref("", faults.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplFixedArray")
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if funcs_ptr.GetZfCurrent() == nil {
		funcs_ptr.SetZfRewind(class_type.FunctionTable().Get("rewind"))
		funcs_ptr.SetZfValid(class_type.FunctionTable().Get("valid"))
		funcs_ptr.SetZfKey(class_type.FunctionTable().Get("key"))
		funcs_ptr.SetZfCurrent(class_type.FunctionTable().Get("current"))
		funcs_ptr.SetZfNext(class_type.FunctionTable().Get("next"))
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
		intern.SetFptrOffsetGet(class_type.FunctionTable().Get("offsetget"))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(class_type.FunctionTable().Get("offsetset"))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(class_type.FunctionTable().Get("offsetexists"))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(class_type.FunctionTable().Get("offsetunset"))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(class_type.FunctionTable().Get("count"))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}
	return intern.GetStd()
}
func SplFixedarrayNew(class_type *types.ClassEntry) *types.Object {
	return SplFixedarrayObjectNewEx(class_type, nil, 0)
}
func SplFixedarrayObjectClone(zobject *types.Zval) *types.Object {
	var old_object *types.Object
	var new_object *types.Object
	old_object = zobject.Object()
	new_object = SplFixedarrayObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplFixedarrayObjectReadDimensionHelper(intern *SplFixedArrayObject, offset *types.Zval) *types.Zval {
	var index zend.ZendLong

	/* we have to return NULL on error here to avoid memleak because of
	 * ZE duplicating uninitialized_zval_ptr */

	if offset == nil {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	}
	if !offset.IsLong() {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.Long()
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
	var intern *SplFixedArrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if type_ == zend.BP_VAR_IS && SplFixedarrayObjectHasDimension(object, offset, 0) == 0 {
		return zend.UninitializedZval()
	}
	if intern.GetFptrOffsetGet() != nil {
		var tmp types.Zval
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			offset = types.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetGet(), "offsetGet", rv, offset)
		// zend.ZvalPtrDtor(offset)
		if !(rv.IsUndef()) {
			return rv
		}
		return zend.UninitializedZval()
	}
	return SplFixedarrayObjectReadDimensionHelper(intern, offset)
}
func SplFixedarrayObjectWriteDimensionHelper(intern *SplFixedArrayObject, offset *types.Zval, value *types.Zval) {
	var index zend.ZendLong
	if offset == nil {

		/* '$array[] = value' syntax is not supported */

		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	}
	if !offset.IsLong() {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.Long()
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
		// zend.ZvalPtrDtor(&tmp)
	}
}
func SplFixedarrayObjectWriteDimension(object *types.Zval, offset *types.Zval, value *types.Zval) {
	var intern *SplFixedArrayObject
	var tmp types.Zval
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetSet() != nil {
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			offset = types.SEPARATE_ARG_IF_REF(offset)
		}
		value = types.SEPARATE_ARG_IF_REF(value)
		zend.ZendCallMethodWith2Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		// zend.ZvalPtrDtor(value)
		// zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectWriteDimensionHelper(intern, offset, value)
}
func SplFixedarrayObjectUnsetDimensionHelper(intern *SplFixedArrayObject, offset *types.Zval) {
	var index zend.ZendLong
	if !offset.IsLong() {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.Long()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		faults.ThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {
		// zend.ZvalPtrDtor(&intern.GetArray().GetElements()[index])
		intern.GetArray().GetElements()[index].SetUndef()
	}
}
func SplFixedarrayObjectUnsetDimension(object *types.Zval, offset *types.Zval) {
	var intern *SplFixedArrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetDel() != nil {
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		// zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectUnsetDimensionHelper(intern, offset)
}
func SplFixedarrayObjectHasDimensionHelper(intern *SplFixedArrayObject, offset *types.Zval, check_empty int) int {
	var index zend.ZendLong
	var retval int
	if !offset.IsLong() {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.Long()
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		retval = 0
	} else {
		if intern.GetArray().GetElements()[index].IsUndef() {
			retval = 0
		} else if check_empty != 0 {
			if operators.IZendIsTrue(intern.GetArray().GetElements()[index]) != 0 {
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
	var intern *SplFixedArrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetHas() != nil {
		var rv types.Zval
		var result bool
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.GetStd().GetCe(), intern.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		// zend.ZvalPtrDtor(offset)
		result = operators.IZendIsTrue(&rv)
		// zend.ZvalPtrDtor(&rv)
		return result
	}
	return SplFixedarrayObjectHasDimensionHelper(intern, offset, check_empty)
}
func SplFixedarrayObjectCountElements(object *types.Zval, count *zend.ZendLong) int {
	var intern *SplFixedArrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if !(rv.IsUndef()) {
			*count = operators.ZvalGetLong(&rv)
			// zend.ZvalPtrDtor(&rv)
		} else {
			*count = 0
		}
	} else {
		*count = intern.GetArray().GetSize()
	}
	return types.SUCCESS
}
func zim_spl_SplFixedArray___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = executeData.ThisObjectZval()
	var intern *SplFixedArrayObject
	var size zend.ZendLong = 0
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "|l", &size) == types.FAILURE {
		return
	}
	if size < 0 {
		faults.ThrowException(spl_ce_InvalidArgumentException, "array size cannot be less than zero", 0)
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
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	var intern_ht *types.Array = zend.ZendStdGetProperties(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if intern.GetArray().GetSize() == 0 {
		var index int = 0
		var size int = intern_ht.Len()
		SplFixedarrayInit(intern.GetArray(), size)
		intern_ht.Foreach(func(key types.ArrayKey, value *types.Zval) {
			types.ZVAL_COPY(intern.GetArray().GetElements()[index], value)
			index++
		})

		/* Remove the unserialised properties, since we now have the elements
		 * within the spl_fixedarray_object structure. */
		intern_ht.Clean()
	}
}
func zim_spl_SplFixedArray_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = executeData.ThisObjectZval()
	var intern *SplFixedArrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	return_value.SetLong(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_toArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedArrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	if intern.GetArray().GetSize() > 0 {
		var i int = 0
		zend.ArrayInit(return_value)
		for ; i < intern.GetArray().GetSize(); i++ {
			if !(intern.GetArray().GetElements()[i].IsUndef()) {
				return_value.Array().IndexUpdate(i, intern.GetArray().GetElements()[i])
				//intern.GetArray().GetElements()[i].TryAddRefcount()
			} else {
				return_value.Array().IndexUpdate(i, zend.UninitializedZval())
			}
		}
	} else {
		return_value.SetEmptyArray()
		return
	}
}
func zim_spl_SplFixedArray_fromArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var data *types.Zval
	var array SplFixedarray
	var intern *SplFixedArrayObject
	var num int
	var save_indexes bool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "a|b", &data, &save_indexes) == types.FAILURE {
		return
	}
	num = data.Array().Len()
	if num > 0 && save_indexes != 0 {
		var maxIndex uint = 0
		for iter := data.Array().Iterator(); iter.Valid(); iter.Next() {
			key := iter.Key()
			if key.IsStrKey() || key.IdxKey() < 0 {
				faults.ThrowException(spl_ce_InvalidArgumentException, "array must contain only positive integer keys", 0)
				return
			}
			numIndex := uint(key.IdxKey())
			if numIndex > maxIndex {
				maxIndex = numIndex
			}
		}
		tmp := int(maxIndex + 1)
		if tmp <= 0 {
			faults.ThrowException(spl_ce_InvalidArgumentException, "integer overflow detected", 0)
			return
		}
		SplFixedarrayInit(&array, tmp)
		data.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
			types.ZVAL_COPY_DEREF(array.GetElements()[key.IdxKey()], value)
		})
	} else if num > 0 && save_indexes == 0 {
		var i zend.ZendLong = 0
		SplFixedarrayInit(&array, num)
		data.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
			types.ZVAL_COPY_DEREF(array.GetElements()[i], value)
			i++
		})
	} else {
		SplFixedarrayInit(&array, 0)
	}
	zend.ObjectInitEx(return_value, spl_ce_SplFixedArray)
	intern = Z_SPLFIXEDARRAY_P(return_value)
	intern.SetArray(array)
}
func zim_spl_SplFixedArray_getSize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = executeData.ThisObjectZval()
	var intern *SplFixedArrayObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	return_value.SetLong(intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_setSize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = executeData.ThisObjectZval()
	var intern *SplFixedArrayObject
	var size zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &size) == types.FAILURE {
		return
	}
	if size < 0 {
		faults.ThrowException(spl_ce_InvalidArgumentException, "array size cannot be less than zero", 0)
		return
	}
	intern = Z_SPLFIXEDARRAY_P(object)
	SplFixedarrayResize(intern.GetArray(), size)
	return_value.SetTrue()
	return
}
func zim_spl_SplFixedArray_offsetExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var intern *SplFixedArrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	return_value.SetBool(SplFixedarrayObjectHasDimensionHelper(intern, zindex, 0) != 0)
	return
}
func zim_spl_SplFixedArray_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var value *types.Zval
	var intern *SplFixedArrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
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
	var intern *SplFixedArrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &zindex, &value) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	SplFixedarrayObjectWriteDimensionHelper(intern, zindex, value)
}
func zim_spl_SplFixedArray_offsetUnset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex *types.Zval
	var intern *SplFixedArrayObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zindex) == types.FAILURE {
		return
	}
	intern = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	SplFixedarrayObjectUnsetDimensionHelper(intern, zindex)
}
func SplFixedarrayItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFixedarrayIt = (*SplFixedarrayIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	// zend.ZvalPtrDtor(iterator.GetIntern().GetIt().GetData())
}
func SplFixedarrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsRewind() {
		zend.ZendUserItRewind(iter)
	} else {
		object.SetCurrent(0)
	}
}
func SplFixedarrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
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
	var object *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsCurrent() {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *types.Zval
		zindex.SetLong(object.GetCurrent())
		data = SplFixedarrayObjectReadDimensionHelper(object, &zindex)
		if data == nil {
			data = zend.UninitializedZval()
		}
		return data
	}
}
func SplFixedarrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsKey() {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		key.SetLong(object.GetCurrent())
	}
}
func SplFixedarrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(iter.GetData())
	if object.IsNext() {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		object.GetCurrent()++
	}
}
func zim_spl_SplFixedArray_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(intern.GetCurrent())
	return
}
func zim_spl_SplFixedArray_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern.GetCurrent()++
}
func zim_spl_SplFixedArray_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetBool(intern.GetCurrent() >= 0 && intern.GetCurrent() < intern.GetArray().GetSize())
	return
}
func zim_spl_SplFixedArray_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	intern.SetCurrent(0)
}
func zim_spl_SplFixedArray_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zindex types.Zval
	var value *types.Zval
	var intern *SplFixedArrayObject = Z_SPLFIXEDARRAY_P(executeData.ThisObjectZval())
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
	// 	object.AddRefcount()
	iterator.GetIntern().GetIt().GetData().SetObject(object.Object())
	iterator.GetIntern().GetIt().SetFuncs(&SplFixedarrayItFuncs)
	iterator.GetIntern().SetCe(ce)
	iterator.GetIntern().GetValue().SetUndef()
	return iterator.GetIntern().GetIt()
}
func ZmStartupSplFixedarray() int {
	spl_ce_SplFixedArray = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "SplFixedArray",
		CreateObject: SplFixedarrayNew,
		Functions:    spl_funcs_SplFixedArray,
		Interfaces: []*types.ClassEntry{
			spl_ce_Iterator,
			spl_ce_ArrayAccess,
			spl_ce_Countable,
		},
		GetIterator: SplFixedarrayGetIterator,
		CeFlags:     types.AccReuseGetIterator,
	})

	spl_handler_SplFixedArray = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		Offset:         int((*byte)(&((*SplFixedArrayObject)(nil).GetStd())) - (*byte)(nil)),
		CloneObj:       SplFixedarrayObjectClone,
		ReadDimension:  SplFixedarrayObjectReadDimension,
		WriteDimension: SplFixedarrayObjectWriteDimension,
		UnsetDimension: SplFixedarrayObjectUnsetDimension,
		HasDimension:   SplFixedarrayObjectHasDimension,
		CountElements:  SplFixedarrayObjectCountElements,
		GetProperties:  SplFixedarrayObjectGetProperties,
		FreeObj:        SplFixedarrayObjectFreeStorage,
	})

	return types.SUCCESS
}

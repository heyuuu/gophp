// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <ext/spl/spl_fixedarray.h>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) The PHP Group                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author: Antony Dovgal <tony@daylessday.org>                          |
  |         Etienne Kneuss <colder@php.net>                              |
  +----------------------------------------------------------------------+
*/

// #define SPL_FIXEDARRAY_H

var spl_ce_SplFixedArray *zend.ZendClassEntry

// Source: <ext/spl/spl_fixedarray.c>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) The PHP Group                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author: Antony Dovgal <tony@daylessday.org>                          |
  |         Etienne Kneuss <colder@php.net>                              |
  +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "zend_exceptions.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_fixedarray.h"

// # include "spl_exceptions.h"

// # include "spl_iterators.h"

var spl_handler_SplFixedArray zend.ZendObjectHandlers

/* }}} */

/* }}} */

/* }}} */

const SPL_FIXEDARRAY_OVERLOADED_REWIND = 0x1
const SPL_FIXEDARRAY_OVERLOADED_VALID = 0x2
const SPL_FIXEDARRAY_OVERLOADED_KEY = 0x4
const SPL_FIXEDARRAY_OVERLOADED_CURRENT = 0x8
const SPL_FIXEDARRAY_OVERLOADED_NEXT = 0x10

func SplFixedArrayFromObj(obj *zend.ZendObject) *SplFixedarrayObject {
	return (*SplFixedarrayObject)((*byte)(obj - zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

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

/* }}} */

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

/* }}} */

func SplFixedarrayCopy(to *SplFixedarray, from *SplFixedarray) {
	var i int
	for i = 0; i < from.GetSize(); i++ {
		zend.ZVAL_COPY(&to.elements[i], &from.elements[i])
	}
}

/* }}} */

func SplFixedarrayObjectGetGc(obj *zend.Zval, table **zend.Zval, n *int) *zend.HashTable {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	*table = intern.GetArray().GetElements()
	*n = int(intern.GetArray().GetSize())
	return ht
}

/* }}}} */

func SplFixedarrayObjectGetProperties(obj *zend.Zval) *zend.HashTable {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	var i zend.ZendLong = 0
	if intern.GetArray().GetSize() > 0 {
		var j zend.ZendLong = zend.ZendHashNumElements(ht)
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			if !(zend.Z_ISUNDEF(intern.GetArray().GetElements()[i])) {
				zend.ZendHashIndexUpdate(ht, i, &intern.array.GetElements()[i])
				zend.Z_TRY_ADDREF(intern.GetArray().GetElements()[i])
			} else {
				zend.ZendHashIndexUpdate(ht, i, &(zend.ExecutorGlobals.uninitialized_zval))
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

/* }}}} */

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
	zend.ZendObjectStdDtor(&intern.std)
}

/* }}} */

func SplFixedarrayObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplFixedarrayObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	var funcs_ptr *zend.ZendClassIteratorFuncs
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_fixedarray_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.SetCurrent(0)
	intern.SetFlags(0)
	if orig != nil && clone_orig != 0 {
		var other *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(orig)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		SplFixedarrayInit(&intern.array, other.GetArray().GetSize())
		SplFixedarrayCopy(&intern.array, &other.array)
	}
	for parent != nil {
		if parent == spl_ce_SplFixedArray {
			intern.std.handlers = &spl_handler_SplFixedArray
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplFixedArray")
	}
	funcs_ptr = class_type.iterator_funcs_ptr
	if funcs_ptr.zf_current == nil {
		funcs_ptr.zf_rewind = zend.ZendHashStrFindPtr(&class_type.function_table, "rewind", b.SizeOf("\"rewind\"")-1)
		funcs_ptr.zf_valid = zend.ZendHashStrFindPtr(&class_type.function_table, "valid", b.SizeOf("\"valid\"")-1)
		funcs_ptr.zf_key = zend.ZendHashStrFindPtr(&class_type.function_table, "key", b.SizeOf("\"key\"")-1)
		funcs_ptr.zf_current = zend.ZendHashStrFindPtr(&class_type.function_table, "current", b.SizeOf("\"current\"")-1)
		funcs_ptr.zf_next = zend.ZendHashStrFindPtr(&class_type.function_table, "next", b.SizeOf("\"next\"")-1)
	}
	if inherited != 0 {
		if funcs_ptr.zf_rewind.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | SPL_FIXEDARRAY_OVERLOADED_REWIND)
		}
		if funcs_ptr.zf_valid.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | SPL_FIXEDARRAY_OVERLOADED_VALID)
		}
		if funcs_ptr.zf_key.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | SPL_FIXEDARRAY_OVERLOADED_KEY)
		}
		if funcs_ptr.zf_current.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | SPL_FIXEDARRAY_OVERLOADED_CURRENT)
		}
		if funcs_ptr.zf_next.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | SPL_FIXEDARRAY_OVERLOADED_NEXT)
		}
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetget", b.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().common.scope == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetset", b.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().common.scope == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetexists", b.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().common.scope == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetunset", b.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().common.scope == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.function_table, "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().common.scope == parent {
			intern.SetFptrCount(nil)
		}
	}
	return &intern.std
}

/* }}} */

func SplFixedarrayNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplFixedarrayObjectNewEx(class_type, nil, 0)
}

/* }}} */

func SplFixedarrayObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zend.Z_OBJ_P(zobject)
	new_object = SplFixedarrayObjectNewEx(old_object.ce, zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}

/* }}} */

func SplFixedarrayObjectReadDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval) *zend.Zval {
	var index zend.ZendLong

	/* we have to return NULL on error here to avoid memleak because of
	 * ZE duplicating uninitialized_zval_ptr */

	if offset == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	}
	if zend.Z_TYPE_P(offset) != zend.IS_LONG {
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
		return &intern.array.GetElements()[index]
	}
}

/* }}} */

func SplFixedarrayObjectReadDimension(object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if type_ == zend.BP_VAR_IS && SplFixedarrayObjectHasDimension(object, offset, 0) == 0 {
		return &(zend.ExecutorGlobals.uninitialized_zval)
	}
	if intern.GetFptrOffsetGet() != nil {
		var tmp zend.Zval
		if offset == nil {
			zend.ZVAL_NULL(&tmp)
			offset = &tmp
		} else {
			zend.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith1Params(object, intern.std.ce, &intern.fptr_offset_get, "offsetGet", rv, offset)
		zend.ZvalPtrDtor(offset)
		if !(zend.Z_ISUNDEF_P(rv)) {
			return rv
		}
		return &(zend.ExecutorGlobals.uninitialized_zval)
	}
	return SplFixedarrayObjectReadDimensionHelper(intern, offset)
}

/* }}} */

func SplFixedarrayObjectWriteDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval, value *zend.Zval) {
	var index zend.ZendLong
	if offset == nil {

		/* '$array[] = value' syntax is not supported */

		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	}
	if zend.Z_TYPE_P(offset) != zend.IS_LONG {
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

/* }}} */

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
		zend.ZendCallMethodWith2Params(object, intern.std.ce, &intern.fptr_offset_set, "offsetSet", nil, offset, value)
		zend.ZvalPtrDtor(value)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectWriteDimensionHelper(intern, offset, value)
}

/* }}} */

func SplFixedarrayObjectUnsetDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval) {
	var index zend.ZendLong
	if zend.Z_TYPE_P(offset) != zend.IS_LONG {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = zend.Z_LVAL_P(offset)
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {
		zend.ZvalPtrDtor(&intern.GetArray().GetElements()[index])
		zend.ZVAL_UNDEF(&intern.array.GetElements()[index])
	}
}

/* }}} */

func SplFixedarrayObjectUnsetDimension(object *zend.Zval, offset *zend.Zval) {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetDel() != nil {
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.std.ce, &intern.fptr_offset_del, "offsetUnset", nil, offset)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectUnsetDimensionHelper(intern, offset)
}

/* }}} */

func SplFixedarrayObjectHasDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval, check_empty int) int {
	var index zend.ZendLong
	var retval int
	if zend.Z_TYPE_P(offset) != zend.IS_LONG {
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
			if zend.ZendIsTrue(&intern.array.GetElements()[index]) != 0 {
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

/* }}} */

func SplFixedarrayObjectHasDimension(object *zend.Zval, offset *zend.Zval, check_empty int) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrOffsetHas() != nil {
		var rv zend.Zval
		var result zend.ZendBool
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, intern.std.ce, &intern.fptr_offset_has, "offsetExists", &rv, offset)
		zend.ZvalPtrDtor(offset)
		result = zend.ZendIsTrue(&rv)
		zend.ZvalPtrDtor(&rv)
		return result
	}
	return SplFixedarrayObjectHasDimensionHelper(intern, offset, check_empty)
}

/* }}} */

func SplFixedarrayObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplFixedarrayObject
	intern = Z_SPLFIXEDARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith0Params(object, intern.std.ce, &intern.fptr_count, "count", &rv)
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

/* }}} */

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
	SplFixedarrayInit(&intern.array, size)
}

/* }}} */

func zim_spl_SplFixedArray___wakeup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	var intern_ht *zend.HashTable = zend.ZendStdGetProperties(zend.ZEND_THIS)
	var data *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetArray().GetSize() == 0 {
		var index int = 0
		var size int = zend.ZendHashNumElements(intern_ht)
		SplFixedarrayInit(&intern.array, size)
		for {
			var __ht *zend.HashTable = intern_ht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				data = _z
				zend.ZVAL_COPY(&intern.array.GetElements()[index], data)
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

/* }}} */

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

/* }}} */

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
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), i, &intern.array.GetElements()[i])
				zend.Z_TRY_ADDREF(intern.GetArray().GetElements()[i])
			} else {
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), i, &(zend.ExecutorGlobals.uninitialized_zval))
			}
		}
	} else {
		zend.RETVAL_EMPTY_ARRAY()
		return
	}
}

/* }}} */

func zim_spl_SplFixedArray_fromArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var data *zend.Zval
	var array SplFixedarray
	var intern *SplFixedarrayObject
	var num int
	var save_indexes zend.ZendBool = 1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "a|b", &data, &save_indexes) == zend.FAILURE {
		return
	}
	num = zend.ZendHashNumElements(zend.Z_ARRVAL_P(data))
	if num > 0 && save_indexes != 0 {
		var element *zend.Zval
		var str_index *zend.ZendString
		var num_index zend.ZendUlong
		var max_index zend.ZendUlong = 0
		var tmp zend.ZendLong
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_index = _p.h
				str_index = _p.key
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
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_index = _p.h
				str_index = _p.key
				element = _z
				zend.ZVAL_COPY_DEREF(&array.elements[num_index], element)
			}
			break
		}
	} else if num > 0 && save_indexes == 0 {
		var element *zend.Zval
		var i zend.ZendLong = 0
		SplFixedarrayInit(&array, num)
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				element = _z
				zend.ZVAL_COPY_DEREF(&array.elements[i], element)
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

/* }}} */

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

/* }}} */

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
	SplFixedarrayResize(&intern.array, size)
	zend.RETVAL_TRUE
	return
}

/* }}} */

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

/* {{{ proto mixed SplFixedArray::offsetGet(mixed $index)
Returns the value at the specified $index. */

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

/* {{{ proto void SplFixedArray::offsetSet(mixed $index, mixed $newval)
Sets the value at the specified $index to $newval. */

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

/* {{{ proto void SplFixedArray::offsetUnset(mixed $index)
Unsets the value at the specified $index. */

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
	zend.ZvalPtrDtor(&iterator.intern.it.data)
}

/* }}} */

func SplFixedarrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(&iter.data)
	if (object.GetFlags() & SPL_FIXEDARRAY_OVERLOADED_REWIND) != 0 {
		zend.ZendUserItRewind(iter)
	} else {
		object.SetCurrent(0)
	}
}

/* }}} */

func SplFixedarrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(&iter.data)
	if (object.GetFlags() & SPL_FIXEDARRAY_OVERLOADED_VALID) != 0 {
		return zend.ZendUserItValid(iter)
	}
	if object.GetCurrent() >= 0 && object.GetCurrent() < object.GetArray().GetSize() {
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func SplFixedarrayItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var zindex zend.Zval
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(&iter.data)
	if (object.GetFlags() & SPL_FIXEDARRAY_OVERLOADED_CURRENT) != 0 {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *zend.Zval
		zend.ZVAL_LONG(&zindex, object.GetCurrent())
		data = SplFixedarrayObjectReadDimensionHelper(object, &zindex)
		if data == nil {
			data = &(zend.ExecutorGlobals.uninitialized_zval)
		}
		return data
	}
}

/* }}} */

func SplFixedarrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(&iter.data)
	if (object.GetFlags() & SPL_FIXEDARRAY_OVERLOADED_KEY) != 0 {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		zend.ZVAL_LONG(key, object.GetCurrent())
	}
}

/* }}} */

func SplFixedarrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(&iter.data)
	if (object.GetFlags() & SPL_FIXEDARRAY_OVERLOADED_NEXT) != 0 {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		object.GetCurrent()++
	}
}

/* }}} */

func zim_spl_SplFixedArray_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetCurrent())
	return
}

/* }}} */

func zim_spl_SplFixedArray_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern.GetCurrent()++
}

/* }}} */

func zim_spl_SplFixedArray_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(intern.GetCurrent() >= 0 && intern.GetCurrent() < intern.GetArray().GetSize())
	return
}

/* }}} */

func zim_spl_SplFixedArray_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = Z_SPLFIXEDARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern.SetCurrent(0)
}

/* }}} */

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

/* }}} */

var SplFixedarrayItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplFixedarrayItDtor, SplFixedarrayItValid, SplFixedarrayItGetCurrentData, SplFixedarrayItGetCurrentKey, SplFixedarrayItMoveForward, SplFixedarrayItRewind, nil}

func SplFixedarrayGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFixedarrayIt
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_fixedarray_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.Z_ADDREF_P(object)
	zend.ZVAL_OBJ(&iterator.intern.it.data, zend.Z_OBJ_P(object))
	iterator.intern.it.funcs = &SplFixedarrayItFuncs
	iterator.intern.ce = ce
	zend.ZVAL_UNDEF(&iterator.intern.value)
	return &iterator.intern.it
}

/* }}} */

var ArginfoSplfixedarrayConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"size", 0, 0, 0}}
var arginfo_fixedarray_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_fixedarray_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var arginfo_fixedarray_setSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var arginfo_fixedarray_fromArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"array", 0, 0, 0}, {"save_indexes", 0, 0, 0}}
var ArginfoSplfixedarrayVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var spl_funcs_SplFixedArray []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplFixedArray___construct,
		ArginfoSplfixedarrayConstruct,
		uint32_t(b.SizeOf("arginfo_splfixedarray_construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__wakeup",
		zim_spl_SplFixedArray___wakeup,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_SplFixedArray_count,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"toArray",
		zim_spl_SplFixedArray_toArray,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fromArray",
		zim_spl_SplFixedArray_fromArray,
		arginfo_fixedarray_fromArray,
		uint32_t(b.SizeOf("arginfo_fixedarray_fromArray")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_STATIC,
	},
	{
		"getSize",
		zim_spl_SplFixedArray_getSize,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setSize",
		zim_spl_SplFixedArray_setSize,
		arginfo_fixedarray_setSize,
		uint32_t(b.SizeOf("arginfo_fixedarray_setSize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetExists",
		zim_spl_SplFixedArray_offsetExists,
		arginfo_fixedarray_offsetGet,
		uint32_t(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetGet",
		zim_spl_SplFixedArray_offsetGet,
		arginfo_fixedarray_offsetGet,
		uint32_t(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetSet",
		zim_spl_SplFixedArray_offsetSet,
		arginfo_fixedarray_offsetSet,
		uint32_t(b.SizeOf("arginfo_fixedarray_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetUnset",
		zim_spl_SplFixedArray_offsetUnset,
		arginfo_fixedarray_offsetGet,
		uint32_t(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_SplFixedArray_rewind,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_SplFixedArray_current,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_SplFixedArray_key,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_SplFixedArray_next,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_SplFixedArray_valid,
		ArginfoSplfixedarrayVoid,
		uint32_t(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplFixedarray(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplFixedArray, "SplFixedArray", SplFixedarrayNew, spl_funcs_SplFixedArray)
	memcpy(&spl_handler_SplFixedArray, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplFixedArray.offset = zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplFixedArray.clone_obj = SplFixedarrayObjectClone
	spl_handler_SplFixedArray.read_dimension = SplFixedarrayObjectReadDimension
	spl_handler_SplFixedArray.write_dimension = SplFixedarrayObjectWriteDimension
	spl_handler_SplFixedArray.unset_dimension = SplFixedarrayObjectUnsetDimension
	spl_handler_SplFixedArray.has_dimension = SplFixedarrayObjectHasDimension
	spl_handler_SplFixedArray.count_elements = SplFixedarrayObjectCountElements
	spl_handler_SplFixedArray.get_properties = SplFixedarrayObjectGetProperties
	spl_handler_SplFixedArray.get_gc = SplFixedarrayObjectGetGc
	spl_handler_SplFixedArray.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplFixedArray.free_obj = SplFixedarrayObjectFreeStorage
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, spl_ce_Countable)
	spl_ce_SplFixedArray.get_iterator = SplFixedarrayGetIterator
	spl_ce_SplFixedArray.ce_flags |= zend.ZEND_ACC_REUSE_GET_ITERATOR
	return zend.SUCCESS
}

/* }}} */

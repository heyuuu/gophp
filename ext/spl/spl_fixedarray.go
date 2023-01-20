// <<generate>>

package spl

import (
	"sik/core"
	g "sik/runtime/grammar"
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

// @type SplFixedarray struct

/* }}} */

// @type SplFixedarrayObject struct

/* }}} */

// @type SplFixedarrayIt struct

/* }}} */

// #define SPL_FIXEDARRAY_OVERLOADED_REWIND       0x0001

// #define SPL_FIXEDARRAY_OVERLOADED_VALID       0x0002

// #define SPL_FIXEDARRAY_OVERLOADED_KEY       0x0004

// #define SPL_FIXEDARRAY_OVERLOADED_CURRENT       0x0008

// #define SPL_FIXEDARRAY_OVERLOADED_NEXT       0x0010

func SplFixedArrayFromObj(obj *zend.ZendObject) *SplFixedarrayObject {
	return (*SplFixedarrayObject)((*byte)(obj - zend_long((*byte)(&((*SplFixedarrayObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLFIXEDARRAY_P(zv) spl_fixed_array_from_obj ( Z_OBJ_P ( ( zv ) ) )

func SplFixedarrayInit(array *SplFixedarray, size zend.ZendLong) {
	if size > 0 {
		array.SetSize(0)
		array.SetElements(zend._ecalloc(size, g.SizeOf("zval")))
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
			zend._efree(elements)
			return
		}
	} else if size > array.GetSize() {
		array.SetElements(zend._safeErealloc(array.GetElements(), size, g.SizeOf("zval"), 0))
		memset(array.GetElements()+array.GetSize(), '0', g.SizeOf("zval")*(size-array.GetSize()))
	} else {
		var i zend.ZendLong
		for i = size; i < array.GetSize(); i++ {
			zend.ZvalPtrDtor(&array.GetElements()[i])
		}
		array.SetElements(zend._erealloc(array.GetElements(), g.SizeOf("zval")*size))
	}
	array.SetSize(size)
}

/* }}} */

func SplFixedarrayCopy(to *SplFixedarray, from *SplFixedarray) {
	var i int
	for i = 0; i < from.GetSize(); i++ {
		var _z1 *zend.Zval = &to.elements[i]
		var _z2 *zend.Zval = &from.elements[i]
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	}
}

/* }}} */

func SplFixedarrayObjectGetGc(obj *zend.Zval, table **zend.Zval, n *int) *zend.HashTable {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(obj.value.obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	*table = intern.GetArray().GetElements()
	*n = int(intern.GetArray().GetSize())
	return ht
}

/* }}}} */

func SplFixedarrayObjectGetProperties(obj *zend.Zval) *zend.HashTable {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(obj.value.obj)
	var ht *zend.HashTable = zend.ZendStdGetProperties(obj)
	var i zend.ZendLong = 0
	if intern.GetArray().GetSize() > 0 {
		var j zend.ZendLong = ht.nNumOfElements
		for i = 0; i < intern.GetArray().GetSize(); i++ {
			if intern.GetArray().GetElements()[i].u1.v.type_ != 0 {
				zend.ZendHashIndexUpdate(ht, i, &intern.array.GetElements()[i])
				if &intern.GetArray().GetElements()[i].u1.v.type_flags != 0 {
					zend.ZvalAddrefP(&intern.GetArray().GetElements()[i])
				}
			} else {
				zend.ZendHashIndexUpdate(ht, i, &zend.EG.uninitialized_zval)
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
			zend._efree(intern.GetArray().GetElements())
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
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_fixedarray_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.SetCurrent(0)
	intern.SetFlags(0)
	if orig != nil && clone_orig != 0 {
		var other *SplFixedarrayObject = SplFixedArrayFromObj(orig.value.obj)
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
		core.PhpErrorDocref(nil, 1<<6, "Internal compiler error, Class is not child of SplFixedArray")
	}
	funcs_ptr = class_type.iterator_funcs_ptr
	if funcs_ptr.zf_current == nil {
		funcs_ptr.zf_rewind = zend.ZendHashStrFindPtr(&class_type.function_table, "rewind", g.SizeOf("\"rewind\"")-1)
		funcs_ptr.zf_valid = zend.ZendHashStrFindPtr(&class_type.function_table, "valid", g.SizeOf("\"valid\"")-1)
		funcs_ptr.zf_key = zend.ZendHashStrFindPtr(&class_type.function_table, "key", g.SizeOf("\"key\"")-1)
		funcs_ptr.zf_current = zend.ZendHashStrFindPtr(&class_type.function_table, "current", g.SizeOf("\"current\"")-1)
		funcs_ptr.zf_next = zend.ZendHashStrFindPtr(&class_type.function_table, "next", g.SizeOf("\"next\"")-1)
	}
	if inherited != 0 {
		if funcs_ptr.zf_rewind.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | 0x1)
		}
		if funcs_ptr.zf_valid.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | 0x2)
		}
		if funcs_ptr.zf_key.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | 0x4)
		}
		if funcs_ptr.zf_current.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | 0x8)
		}
		if funcs_ptr.zf_next.common.scope != parent {
			intern.SetFlags(intern.GetFlags() | 0x10)
		}
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetget", g.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().common.scope == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetset", g.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().common.scope == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetexists", g.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().common.scope == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetunset", g.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().common.scope == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.function_table, "count", g.SizeOf("\"count\"")-1))
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
	old_object = zobject.value.obj
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
	if offset.u1.v.type_ != 4 {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.value.lval
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return nil
	} else if intern.GetArray().GetElements()[index].u1.v.type_ == 0 {
		return nil
	} else {
		return &intern.array.GetElements()[index]
	}
}

/* }}} */

func SplFixedarrayObjectReadDimension(object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	var intern *SplFixedarrayObject
	intern = SplFixedArrayFromObj(object.value.obj)
	if type_ == 3 && SplFixedarrayObjectHasDimension(object, offset, 0) == 0 {
		return &zend.EG.uninitialized_zval
	}
	if intern.GetFptrOffsetGet() != nil {
		var tmp zend.Zval
		if offset == nil {
			&tmp.u1.type_info = 1
			offset = &tmp
		} else {
			if offset.u1.v.type_ == 10 {
				offset = &(*offset).value.ref.val
			}
			if offset.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(offset)
			}
		}
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_offset_get, "offsetGet", g.SizeOf("\"offsetGet\"")-1, rv, 1, offset, nil)
		zend.ZvalPtrDtor(offset)
		if rv.u1.v.type_ != 0 {
			return rv
		}
		return &zend.EG.uninitialized_zval
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
	if offset.u1.v.type_ != 4 {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.value.lval
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {

		/* Fix #81429 */

		var ptr *zend.Zval = &intern.GetArray().GetElements()[index]
		var tmp zend.Zval
		var _z1 *zend.Zval = &tmp
		var _z2 *zend.Zval = ptr
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = ptr
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		zend.ZvalPtrDtor(&tmp)
	}
}

/* }}} */

func SplFixedarrayObjectWriteDimension(object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	var intern *SplFixedarrayObject
	var tmp zend.Zval
	intern = SplFixedArrayFromObj(object.value.obj)
	if intern.GetFptrOffsetSet() != nil {
		if offset == nil {
			&tmp.u1.type_info = 1
			offset = &tmp
		} else {
			if offset.u1.v.type_ == 10 {
				offset = &(*offset).value.ref.val
			}
			if offset.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(offset)
			}
		}
		if value.u1.v.type_ == 10 {
			value = &(*value).value.ref.val
		}
		if value.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(value)
		}
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_offset_set, "offsetSet", g.SizeOf("\"offsetSet\"")-1, nil, 2, offset, value)
		zend.ZvalPtrDtor(value)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectWriteDimensionHelper(intern, offset, value)
}

/* }}} */

func SplFixedarrayObjectUnsetDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval) {
	var index zend.ZendLong
	if offset.u1.v.type_ != 4 {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.value.lval
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Index invalid or out of range", 0)
		return
	} else {
		zend.ZvalPtrDtor(&intern.GetArray().GetElements()[index])
		&intern.array.GetElements()[index].u1.type_info = 0
	}
}

/* }}} */

func SplFixedarrayObjectUnsetDimension(object *zend.Zval, offset *zend.Zval) {
	var intern *SplFixedarrayObject
	intern = SplFixedArrayFromObj(object.value.obj)
	if intern.GetFptrOffsetDel() != nil {
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		if offset.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(offset)
		}
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_offset_del, "offsetUnset", g.SizeOf("\"offsetUnset\"")-1, nil, 1, offset, nil)
		zend.ZvalPtrDtor(offset)
		return
	}
	SplFixedarrayObjectUnsetDimensionHelper(intern, offset)
}

/* }}} */

func SplFixedarrayObjectHasDimensionHelper(intern *SplFixedarrayObject, offset *zend.Zval, check_empty int) int {
	var index zend.ZendLong
	var retval int
	if offset.u1.v.type_ != 4 {
		index = SplOffsetConvertToLong(offset)
	} else {
		index = offset.value.lval
	}
	if index < 0 || index >= intern.GetArray().GetSize() {
		retval = 0
	} else {
		if intern.GetArray().GetElements()[index].u1.v.type_ == 0 {
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
	intern = SplFixedArrayFromObj(object.value.obj)
	if intern.GetFptrOffsetHas() != nil {
		var rv zend.Zval
		var result zend.ZendBool
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		if offset.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(offset)
		}
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_offset_has, "offsetExists", g.SizeOf("\"offsetExists\"")-1, &rv, 1, offset, nil)
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
	intern = SplFixedArrayFromObj(object.value.obj)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_count, "count", g.SizeOf("\"count\"")-1, &rv, 0, nil, nil)
		if rv.u1.v.type_ != 0 {
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
	var object *zend.Zval = &(execute_data.This)
	var intern *SplFixedarrayObject
	var size zend.ZendLong = 0
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "|l", &size) == zend.FAILURE {
		return
	}
	if size < 0 {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
		return
	}
	intern = SplFixedArrayFromObj(object.value.obj)
	if intern.GetArray().GetSize() > 0 {

		/* called __construct() twice, bail out */

		return

		/* called __construct() twice, bail out */

	}
	SplFixedarrayInit(&intern.array, size)
}

/* }}} */

func zim_spl_SplFixedArray___wakeup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	var intern_ht *zend.HashTable = zend.ZendStdGetProperties(&(execute_data.This))
	var data *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetArray().GetSize() == 0 {
		var index int = 0
		var size int = intern_ht.nNumOfElements
		SplFixedarrayInit(&intern.array, size)
		for {
			var __ht *zend.HashTable = intern_ht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				data = _z
				var _z1 *zend.Zval = &intern.array.GetElements()[index]
				var _z2 *zend.Zval = data
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				if (_t & 0xff00) != 0 {
					zend.ZendGcAddref(&_gc.gc)
				}
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
	var object *zend.Zval = &(execute_data.This)
	var intern *SplFixedarrayObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(object.value.obj)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetArray().GetSize()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplFixedArray_toArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if intern.GetArray().GetSize() > 0 {
		var i int = 0
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for ; i < intern.GetArray().GetSize(); i++ {
			if intern.GetArray().GetElements()[i].u1.v.type_ != 0 {
				zend.ZendHashIndexUpdate(return_value.value.arr, i, &intern.array.GetElements()[i])
				if &intern.GetArray().GetElements()[i].u1.v.type_flags != 0 {
					zend.ZvalAddrefP(&intern.GetArray().GetElements()[i])
				}
			} else {
				zend.ZendHashIndexUpdate(return_value.value.arr, i, &zend.EG.uninitialized_zval)
			}
		}
	} else {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
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
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "a|b", &data, &save_indexes) == zend.FAILURE {
		return
	}
	num = data.value.arr.nNumOfElements
	if num > 0 && save_indexes != 0 {
		var element *zend.Zval
		var str_index *zend.ZendString
		var num_index zend.ZendUlong
		var max_index zend.ZendUlong = 0
		var tmp zend.ZendLong
		for {
			var __ht *zend.HashTable = data.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
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
			var __ht *zend.HashTable = data.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				num_index = _p.h
				str_index = _p.key
				element = _z
				var _z3 *zend.Zval = element
				if (_z3.u1.type_info & 0xff00) != 0 {
					if (_z3.u1.type_info & 0xff) == 10 {
						_z3 = &(*_z3).value.ref.val
						if (_z3.u1.type_info & 0xff00) != 0 {
							zend.ZvalAddrefP(_z3)
						}
					} else {
						zend.ZvalAddrefP(_z3)
					}
				}
				var _z1 *zend.Zval = &array.elements[num_index]
				var _z2 *zend.Zval = _z3
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			}
			break
		}
	} else if num > 0 && save_indexes == 0 {
		var element *zend.Zval
		var i zend.ZendLong = 0
		SplFixedarrayInit(&array, num)
		for {
			var __ht *zend.HashTable = data.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				element = _z
				var _z3 *zend.Zval = element
				if (_z3.u1.type_info & 0xff00) != 0 {
					if (_z3.u1.type_info & 0xff) == 10 {
						_z3 = &(*_z3).value.ref.val
						if (_z3.u1.type_info & 0xff00) != 0 {
							zend.ZvalAddrefP(_z3)
						}
					} else {
						zend.ZvalAddrefP(_z3)
					}
				}
				var _z1 *zend.Zval = &array.elements[i]
				var _z2 *zend.Zval = _z3
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				i++
			}
			break
		}
	} else {
		SplFixedarrayInit(&array, 0)
	}
	zend.ObjectInitEx(return_value, spl_ce_SplFixedArray)
	intern = SplFixedArrayFromObj(return_value.value.obj)
	intern.SetArray(array)
}

/* }}} */

func zim_spl_SplFixedArray_getSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplFixedarrayObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(object.value.obj)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetArray().GetSize()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplFixedArray_setSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplFixedarrayObject
	var size zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &size) == zend.FAILURE {
		return
	}
	if size < 0 {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "array size cannot be less than zero")
		return
	}
	intern = SplFixedArrayFromObj(object.value.obj)
	SplFixedarrayResize(&intern.array, size)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplFixedArray_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if SplFixedarrayObjectHasDimensionHelper(intern, zindex, 0) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed SplFixedArray::offsetGet(mixed $index)
Returns the value at the specified $index. */

func zim_spl_SplFixedArray_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	value = SplFixedarrayObjectReadDimensionHelper(intern, zindex)
	if value != nil {
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto void SplFixedArray::offsetSet(mixed $index, mixed $newval)
Sets the value at the specified $index to $newval. */

func zim_spl_SplFixedArray_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	SplFixedarrayObjectWriteDimensionHelper(intern, zindex, value)
}

/* {{{ proto void SplFixedArray::offsetUnset(mixed $index)
Unsets the value at the specified $index. */

func zim_spl_SplFixedArray_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplFixedarrayObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	SplFixedarrayObjectUnsetDimensionHelper(intern, zindex)
}
func SplFixedarrayItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFixedarrayIt = (*SplFixedarrayIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iterator.intern.it.data)
}

/* }}} */

func SplFixedarrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = SplFixedArrayFromObj(&iter.data.value.obj)
	if (object.GetFlags() & 0x1) != 0 {
		zend.ZendUserItRewind(iter)
	} else {
		object.SetCurrent(0)
	}
}

/* }}} */

func SplFixedarrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFixedarrayObject = SplFixedArrayFromObj(&iter.data.value.obj)
	if (object.GetFlags() & 0x2) != 0 {
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
	var object *SplFixedarrayObject = SplFixedArrayFromObj(&iter.data.value.obj)
	if (object.GetFlags() & 0x8) != 0 {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *zend.Zval
		var __z *zend.Zval = &zindex
		__z.value.lval = object.GetCurrent()
		__z.u1.type_info = 4
		data = SplFixedarrayObjectReadDimensionHelper(object, &zindex)
		if data == nil {
			data = &zend.EG.uninitialized_zval
		}
		return data
	}
}

/* }}} */

func SplFixedarrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFixedarrayObject = SplFixedArrayFromObj(&iter.data.value.obj)
	if (object.GetFlags() & 0x4) != 0 {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		var __z *zend.Zval = key
		__z.value.lval = object.GetCurrent()
		__z.u1.type_info = 4
	}
}

/* }}} */

func SplFixedarrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFixedarrayObject = SplFixedArrayFromObj(&iter.data.value.obj)
	if (object.GetFlags() & 0x10) != 0 {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		object.GetCurrent()++
	}
}

/* }}} */

func zim_spl_SplFixedArray_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetCurrent()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplFixedArray_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern.GetCurrent()++
}

/* }}} */

func zim_spl_SplFixedArray_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetCurrent() >= 0 && intern.GetCurrent() < intern.GetArray().GetSize() {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplFixedArray_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern.SetCurrent(0)
}

/* }}} */

func zim_spl_SplFixedArray_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex zend.Zval
	var value *zend.Zval
	var intern *SplFixedarrayObject = SplFixedArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = &zindex
	__z.value.lval = intern.GetCurrent()
	__z.u1.type_info = 4
	value = SplFixedarrayObjectReadDimensionHelper(intern, &zindex)
	if value != nil {
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
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
	iterator = zend._emalloc(g.SizeOf("spl_fixedarray_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.it.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.it.funcs = &SplFixedarrayItFuncs
	iterator.intern.ce = ce
	&iterator.intern.value.u1.type_info = 0
	return &iterator.intern.it
}

/* }}} */

var ArginfoSplfixedarrayConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"size", 0, 0, 0}}
var arginfo_fixedarray_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_fixedarray_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var arginfo_fixedarray_setSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var arginfo_fixedarray_fromArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"array", 0, 0, 0}, {"save_indexes", 0, 0, 0}}
var ArginfoSplfixedarrayVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var spl_funcs_SplFixedArray []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplFixedArray___construct,
		ArginfoSplfixedarrayConstruct,
		uint32(g.SizeOf("arginfo_splfixedarray_construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__wakeup",
		zim_spl_SplFixedArray___wakeup,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_SplFixedArray_count,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"toArray",
		zim_spl_SplFixedArray_toArray,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fromArray",
		zim_spl_SplFixedArray_fromArray,
		arginfo_fixedarray_fromArray,
		uint32(g.SizeOf("arginfo_fixedarray_fromArray")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<4,
	},
	{
		"getSize",
		zim_spl_SplFixedArray_getSize,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setSize",
		zim_spl_SplFixedArray_setSize,
		arginfo_fixedarray_setSize,
		uint32(g.SizeOf("arginfo_fixedarray_setSize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetExists",
		zim_spl_SplFixedArray_offsetExists,
		arginfo_fixedarray_offsetGet,
		uint32(g.SizeOf("arginfo_fixedarray_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetGet",
		zim_spl_SplFixedArray_offsetGet,
		arginfo_fixedarray_offsetGet,
		uint32(g.SizeOf("arginfo_fixedarray_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetSet",
		zim_spl_SplFixedArray_offsetSet,
		arginfo_fixedarray_offsetSet,
		uint32(g.SizeOf("arginfo_fixedarray_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetUnset",
		zim_spl_SplFixedArray_offsetUnset,
		arginfo_fixedarray_offsetGet,
		uint32(g.SizeOf("arginfo_fixedarray_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_SplFixedArray_rewind,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_SplFixedArray_current,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_SplFixedArray_key,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_SplFixedArray_next,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_SplFixedArray_valid,
		ArginfoSplfixedarrayVoid,
		uint32(g.SizeOf("arginfo_splfixedarray_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplFixedarray(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplFixedArray, "SplFixedArray", SplFixedarrayNew, spl_funcs_SplFixedArray)
	memcpy(&spl_handler_SplFixedArray, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
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
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, zend.ZendCeArrayaccess)
	zend.ZendClassImplements(spl_ce_SplFixedArray, 1, zend.ZendCeCountable)
	spl_ce_SplFixedArray.get_iterator = SplFixedarrayGetIterator
	spl_ce_SplFixedArray.ce_flags |= 1 << 18
	return zend.SUCCESS
}

/* }}} */

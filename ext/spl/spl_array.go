// <<generate>>

package spl

import (
	"sik/core"
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/spl_array.h>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define SPL_ARRAY_H

// # include "php.h"

// # include "php_spl.h"

// # include "spl_iterators.h"

var spl_ce_ArrayObject *zend.ZendClassEntry
var spl_ce_ArrayIterator *zend.ZendClassEntry
var spl_ce_RecursiveArrayIterator *zend.ZendClassEntry

// Source: <ext/spl/spl_array.c>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "ext/standard/php_var.h"

// # include "zend_smart_str.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_array.h"

// # include "spl_exceptions.h"

var spl_handler_ArrayObject zend.ZendObjectHandlers
var spl_handler_ArrayIterator zend.ZendObjectHandlers

// #define SPL_ARRAY_STD_PROP_LIST       0x00000001

// #define SPL_ARRAY_ARRAY_AS_PROPS       0x00000002

// #define SPL_ARRAY_CHILD_ARRAYS_ONLY       0x00000004

// #define SPL_ARRAY_OVERLOADED_REWIND       0x00010000

// #define SPL_ARRAY_OVERLOADED_VALID       0x00020000

// #define SPL_ARRAY_OVERLOADED_KEY       0x00040000

// #define SPL_ARRAY_OVERLOADED_CURRENT       0x00080000

// #define SPL_ARRAY_OVERLOADED_NEXT       0x00100000

// #define SPL_ARRAY_IS_SELF       0x01000000

// #define SPL_ARRAY_USE_OTHER       0x02000000

// #define SPL_ARRAY_INT_MASK       0xFFFF0000

// #define SPL_ARRAY_CLONE_MASK       0x0100FFFF

// #define SPL_ARRAY_METHOD_NO_ARG       0

// #define SPL_ARRAY_METHOD_USE_ARG       1

// #define SPL_ARRAY_METHOD_MAY_USER_ARG       2

// @type SplArrayObject struct

func SplArrayFromObj(obj *zend.ZendObject) *SplArrayObject {
	return (*SplArrayObject)((*byte)(obj - zend_long((*byte)(&((*SplArrayObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLARRAY_P(zv) spl_array_from_obj ( Z_OBJ_P ( ( zv ) ) )

func SplArrayGetHashTablePtr(intern *SplArrayObject) **zend.HashTable {
	//??? TODO: Delay duplication for arrays; only duplicate for write operations

	if (intern.GetArFlags() & 0x1000000) != 0 {
		if intern.std.properties == nil {
			zend.RebuildObjectProperties(&intern.std)
		}
		return &intern.std.properties
	} else if (intern.GetArFlags() & 0x2000000) != 0 {
		var other *SplArrayObject = SplArrayFromObj(&intern.array.value.obj)
		return SplArrayGetHashTablePtr(other)
	} else if intern.array.u1.v.type_ == 7 {
		return &(intern.GetArray()).value.arr
	} else {
		var obj *zend.ZendObject = intern.array.value.obj
		if obj.properties == nil {
			zend.RebuildObjectProperties(obj)
		} else if zend.ZendGcRefcount(&(obj.properties).gc) > 1 {
			if (zend.ZvalGcFlags(obj.properties.gc.u.type_info) & 1 << 6) == 0 {
				zend.ZendGcDelref(&(obj.properties).gc)
			}
			obj.properties = zend.ZendArrayDup(obj.properties)
		}
		return &obj.properties
	}

	//??? TODO: Delay duplication for arrays; only duplicate for write operations
}

/* }}} */

func SplArrayGetHashTable(intern *SplArrayObject) *zend.HashTable {
	return (*SplArrayGetHashTablePtr)(intern)
}

/* }}} */

func SplArrayReplaceHashTable(intern *SplArrayObject, ht *zend.HashTable) {
	var ht_ptr **zend.HashTable = SplArrayGetHashTablePtr(intern)
	zend.ZendArrayDestroy(*ht_ptr)
	*ht_ptr = ht
}

/* }}} */

func SplArrayIsObject(intern *SplArrayObject) zend.ZendBool {
	for (intern.GetArFlags() & 0x2000000) != 0 {
		intern = SplArrayFromObj(&intern.array.value.obj)
	}
	return (intern.GetArFlags()&0x1000000) != 0 || intern.array.u1.v.type_ == 8
}

/* }}} */

func SplArrayCreateHtIter(ht *zend.HashTable, intern *SplArrayObject) {
	intern.SetHtIter(zend.ZendHashIteratorAdd(ht, zend.ZendHashGetCurrentPos(ht)))
	zend.ZendHashInternalPointerResetEx(ht, &zend.EG.ht_iterators[intern.GetHtIter()].pos)
	SplArraySkipProtected(intern, ht)
}

/* }}} */

func SplArrayGetPosPtr(ht *zend.HashTable, intern *SplArrayObject) *uint32 {
	if intern.GetHtIter() == uint32-1 {
		SplArrayCreateHtIter(ht, intern)
	}
	return &zend.EG.ht_iterators[intern.GetHtIter()].pos
}

/* }}} */

func SplArrayObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplArrayObject = SplArrayFromObj(object)
	if intern.GetHtIter() != uint32-1 {
		zend.ZendHashIteratorDel(intern.GetHtIter())
	}
	zend.ZendObjectStdDtor(&intern.std)
	zend.ZvalPtrDtor(&intern.array)
}

/* }}} */

/* {{{ spl_array_object_new_ex */

func SplArrayObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplArrayObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_array_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.SetArFlags(0)
	intern.SetCeGetIterator(spl_ce_ArrayIterator)
	if orig != nil {
		var other *SplArrayObject = SplArrayFromObj(orig.value.obj)
		intern.SetArFlags(intern.GetArFlags() &^ 0x100ffff)
		intern.SetArFlags(intern.GetArFlags() | other.GetArFlags()&0x100ffff)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			if (other.GetArFlags() & 0x1000000) != 0 {
				&intern.array.u1.type_info = 0
			} else if orig.value.obj.handlers == &spl_handler_ArrayObject {
				var __arr *zend.ZendArray = zend.ZendArrayDup(SplArrayGetHashTable(other))
				var __z *zend.Zval = &intern.array
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			} else {
				assert(orig.value.obj.handlers == &spl_handler_ArrayIterator)
				var _z1 *zend.Zval = &intern.array
				var _z2 *zend.Zval = orig
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				if (_t & 0xff00) != 0 {
					zend.ZendGcAddref(&_gc.gc)
				}
				intern.SetArFlags(intern.GetArFlags() | 0x2000000)
			}
		} else {
			var _z1 *zend.Zval = &intern.array
			var _z2 *zend.Zval = orig
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			intern.SetArFlags(intern.GetArFlags() | 0x2000000)
		}
	} else {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &intern.array
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	for parent != nil {
		if parent == spl_ce_ArrayIterator || parent == spl_ce_RecursiveArrayIterator {
			intern.std.handlers = &spl_handler_ArrayIterator
			break
		} else if parent == spl_ce_ArrayObject {
			intern.std.handlers = &spl_handler_ArrayObject
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, 1<<6, "Internal compiler error, Class is not child of ArrayObject or ArrayIterator")
	}
	if inherited != 0 {
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

	/* Cache iterator functions if ArrayIterator or derived. Check current's */

	if intern.std.handlers == &spl_handler_ArrayIterator {
		var funcs_ptr *zend.ZendClassIteratorFuncs = class_type.iterator_funcs_ptr
		if funcs_ptr.zf_current == nil {
			funcs_ptr.zf_rewind = zend.ZendHashStrFindPtr(&class_type.function_table, "rewind", g.SizeOf("\"rewind\"")-1)
			funcs_ptr.zf_valid = zend.ZendHashStrFindPtr(&class_type.function_table, "valid", g.SizeOf("\"valid\"")-1)
			funcs_ptr.zf_key = zend.ZendHashStrFindPtr(&class_type.function_table, "key", g.SizeOf("\"key\"")-1)
			funcs_ptr.zf_current = zend.ZendHashStrFindPtr(&class_type.function_table, "current", g.SizeOf("\"current\"")-1)
			funcs_ptr.zf_next = zend.ZendHashStrFindPtr(&class_type.function_table, "next", g.SizeOf("\"next\"")-1)
		}
		if inherited != 0 {
			if funcs_ptr.zf_rewind.common.scope != parent {
				intern.SetArFlags(intern.GetArFlags() | 0x10000)
			}
			if funcs_ptr.zf_valid.common.scope != parent {
				intern.SetArFlags(intern.GetArFlags() | 0x20000)
			}
			if funcs_ptr.zf_key.common.scope != parent {
				intern.SetArFlags(intern.GetArFlags() | 0x40000)
			}
			if funcs_ptr.zf_current.common.scope != parent {
				intern.SetArFlags(intern.GetArFlags() | 0x80000)
			}
			if funcs_ptr.zf_next.common.scope != parent {
				intern.SetArFlags(intern.GetArFlags() | 0x100000)
			}
		}
	}
	intern.SetHtIter(uint32 - 1)
	return &intern.std
}

/* }}} */

func SplArrayObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplArrayObjectNewEx(class_type, nil, 0)
}

/* }}} */

func SplArrayObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.value.obj
	new_object = SplArrayObjectNewEx(old_object.ce, zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}

/* }}} */

func SplArrayGetDimensionPtr(check_inherited int, intern *SplArrayObject, offset *zend.Zval, type_ int) *zend.Zval {
	var retval *zend.Zval
	var index zend.ZendLong
	var offset_key *zend.ZendString
	var ht *zend.HashTable = SplArrayGetHashTable(intern)
	if offset == nil || offset.u1.v.type_ == 0 || ht == nil {
		return &zend.EG.uninitialized_zval
	}
	if (type_ == 1 || type_ == 2) && intern.GetNApplyCount() > 0 {
		zend.ZendError(1<<1, "Modification of ArrayObject during sorting is prohibited")
		return &zend.EG.error_zval
	}
try_again:
	switch offset.u1.v.type_ {
	case 1:
		offset_key = zend.ZendEmptyString
		goto fetch_dim_string
	case 6:
		offset_key = offset.value.str
	fetch_dim_string:
		retval = zend.ZendSymtableFind(ht, offset_key)
		if retval != nil {
			if retval.u1.v.type_ == 13 {
				retval = retval.value.zv
				if retval.u1.v.type_ == 0 {
					switch type_ {
					case 0:
						zend.ZendError(1<<3, "Undefined index: %s", offset_key.val)
					case 5:

					case 3:
						retval = &zend.EG.uninitialized_zval
						break
					case 2:
						zend.ZendError(1<<3, "Undefined index: %s", offset_key.val)
					case 1:
						retval.u1.type_info = 1
					}
				}
			}
		} else {
			switch type_ {
			case 0:
				zend.ZendError(1<<3, "Undefined index: %s", offset_key.val)
			case 5:

			case 3:
				retval = &zend.EG.uninitialized_zval
				break
			case 2:
				zend.ZendError(1<<3, "Undefined index: %s", offset_key.val)
			case 1:
				var value zend.Zval
				&value.u1.type_info = 1
				retval = zend.ZendSymtableUpdate(ht, offset_key, &value)
			}
		}
		return retval
	case 9:
		zend.ZendError(1<<3, "Resource ID#%d used as offset, casting to integer (%d)", offset.value.res.handle, offset.value.res.handle)
		index = offset.value.res.handle
		goto num_index
	case 5:
		index = zend_long(*offset).value.dval
		goto num_index
	case 2:
		index = 0
		goto num_index
	case 3:
		index = 1
		goto num_index
	case 4:
		index = offset.value.lval
	num_index:
		if g.Assign(&retval, zend.ZendHashIndexFind(ht, index)) == nil {
			switch type_ {
			case 0:
				zend.ZendError(1<<3, "Undefined offset: "+"%"+"lld", index)
			case 5:

			case 3:
				retval = &zend.EG.uninitialized_zval
				break
			case 2:
				zend.ZendError(1<<3, "Undefined offset: "+"%"+"lld", index)
			case 1:
				var value zend.Zval
				&value.u1.type_info = 0
				retval = zend.ZendHashIndexUpdate(ht, index, &value)
			}
		}
		return retval
	case 10:
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		goto try_again
	default:
		zend.ZendError(1<<1, "Illegal offset type")
		if type_ == 1 || type_ == 2 {
			return &zend.EG.error_zval
		} else {
			return &zend.EG.uninitialized_zval
		}
	}
}
func SplArrayReadDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var ret *zend.Zval
	if check_inherited != 0 && (intern.GetFptrOffsetGet() != nil || type_ == 3 && intern.GetFptrOffsetHas() != nil) {
		if type_ == 3 {
			if SplArrayHasDimension(object, offset, 0) == 0 {
				return &zend.EG.uninitialized_zval
			}
		}
		if intern.GetFptrOffsetGet() != nil {
			var tmp zend.Zval
			if offset == nil {
				&tmp.u1.type_info = 0
				offset = &tmp
			} else {
				if offset.u1.v.type_ == 10 {
					offset = &(*offset).value.ref.val
				}
				if offset.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(offset)
				}
			}
			zend.ZendCallMethod(object, object.value.obj.ce, &intern.fptr_offset_get, "offsetGet", g.SizeOf("\"offsetGet\"")-1, rv, 1, offset, nil)
			zend.ZvalPtrDtor(offset)
			if rv.u1.v.type_ != 0 {
				return rv
			}
			return &zend.EG.uninitialized_zval
		}
	}
	ret = SplArrayGetDimensionPtr(check_inherited, intern, offset, type_)

	/* When in a write context,
	 * ZE has to be fooled into thinking this is in a reference set
	 * by separating (if necessary) and returning as IS_REFERENCE (with refcount == 1)
	 */

	if (type_ == 1 || type_ == 2 || type_ == 5) && ret.u1.v.type_ != 10 && ret != &zend.EG.uninitialized_zval {
		var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
		zend.ZendGcSetRefcount(&_ref.gc, 1)
		_ref.gc.u.type_info = 10
		var _z1 *zend.Zval = &_ref.val
		var _z2 *zend.Zval = ret
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		_ref.sources.ptr = nil
		ret.value.ref = _ref
		ret.u1.type_info = 10 | 1<<0<<8
	}
	return ret
}
func SplArrayReadDimension(object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	return SplArrayReadDimensionEx(1, object, offset, type_, rv)
}
func SplArrayWriteDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var index zend.ZendLong
	var ht *zend.HashTable
	if check_inherited != 0 && intern.GetFptrOffsetSet() != nil {
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
		zend.ZendCallMethod(object, object.value.obj.ce, &intern.fptr_offset_set, "offsetSet", g.SizeOf("\"offsetSet\"")-1, nil, 2, offset, value)
		zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(1<<1, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	if value.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(value)
	}
	if offset == nil {
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashNextIndexInsert(ht, value)
		return
	}
try_again:
	switch offset.u1.v.type_ {
	case 6:
		ht = SplArrayGetHashTable(intern)
		zend.ZendSymtableUpdateInd(ht, offset.value.str, value)
		return
	case 5:
		index = zend_long(*offset).value.dval
		goto num_index
	case 9:
		index = offset.value.res.handle
		goto num_index
	case 2:
		index = 0
		goto num_index
	case 3:
		index = 1
		goto num_index
	case 4:
		index = offset.value.lval
	num_index:
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashIndexUpdate(ht, index, value)
		return
	case 1:
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashNextIndexInsert(ht, value)
		return
	case 10:
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		goto try_again
	default:
		zend.ZendError(1<<1, "Illegal offset type")
		zend.ZvalPtrDtor(value)
		return
	}
}
func SplArrayWriteDimension(object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	SplArrayWriteDimensionEx(1, object, offset, value)
}
func SplArrayUnsetDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval) {
	var index zend.ZendLong
	var ht *zend.HashTable
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if check_inherited != 0 && intern.GetFptrOffsetDel() != nil {
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		if offset.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(offset)
		}
		zend.ZendCallMethod(object, object.value.obj.ce, &intern.fptr_offset_del, "offsetUnset", g.SizeOf("\"offsetUnset\"")-1, nil, 1, offset, nil)
		zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(1<<1, "Modification of ArrayObject during sorting is prohibited")
		return
	}
try_again:
	switch offset.u1.v.type_ {
	case 6:
		ht = SplArrayGetHashTable(intern)
		if ht == &zend.EG.symbol_table {
			if zend.ZendDeleteGlobalVariable(offset.value.str) != 0 {
				zend.ZendError(1<<3, "Undefined index: %s", offset.value.str.val)
			}
		} else {
			var data *zend.Zval = zend.ZendSymtableFind(ht, offset.value.str)
			if data != nil {
				if data.u1.v.type_ == 13 {
					data = data.value.zv
					if data.u1.v.type_ == 0 {
						zend.ZendError(1<<3, "Undefined index: %s", offset.value.str.val)
					} else {
						zend.ZvalPtrDtor(data)
						data.u1.type_info = 0
						ht.u.flags |= 1 << 5
						zend.ZendHashMoveForwardEx(ht, SplArrayGetPosPtr(ht, intern))
						if SplArrayIsObject(intern) != 0 {
							SplArraySkipProtected(intern, ht)
						}
					}
				} else if zend.ZendSymtableDel(ht, offset.value.str) == zend.FAILURE {
					zend.ZendError(1<<3, "Undefined index: %s", offset.value.str.val)
				}
			} else {
				zend.ZendError(1<<3, "Undefined index: %s", offset.value.str.val)
			}
		}
		break
	case 5:
		index = zend_long(*offset).value.dval
		goto num_index
	case 9:
		index = offset.value.res.handle
		goto num_index
	case 2:
		index = 0
		goto num_index
	case 3:
		index = 1
		goto num_index
	case 4:
		index = offset.value.lval
	num_index:
		ht = SplArrayGetHashTable(intern)
		if zend.ZendHashIndexDel(ht, index) == zend.FAILURE {
			zend.ZendError(1<<3, "Undefined offset: "+"%"+"lld", index)
		}
		break
	case 10:
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		goto try_again
	default:
		zend.ZendError(1<<1, "Illegal offset type")
		return
	}
}
func SplArrayUnsetDimension(object *zend.Zval, offset *zend.Zval) {
	SplArrayUnsetDimensionEx(1, object, offset)
}
func SplArrayHasDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, check_empty int) int {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var index zend.ZendLong
	var rv zend.Zval
	var value *zend.Zval = nil
	var tmp *zend.Zval
	if check_inherited != 0 && intern.GetFptrOffsetHas() != nil {
		if offset.u1.v.type_ == 10 {
			offset = &(*offset).value.ref.val
		}
		if offset.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(offset)
		}
		zend.ZendCallMethod(object, object.value.obj.ce, &intern.fptr_offset_has, "offsetExists", g.SizeOf("\"offsetExists\"")-1, &rv, 1, offset, nil)
		zend.ZvalPtrDtor(offset)
		if zend.ZendIsTrue(&rv) != 0 {
			zend.ZvalPtrDtor(&rv)
			if check_empty != 1 {
				return 1
			} else if intern.GetFptrOffsetGet() != nil {
				value = SplArrayReadDimensionEx(1, object, offset, 0, &rv)
			}
		} else {
			zend.ZvalPtrDtor(&rv)
			return 0
		}
	}
	if value == nil {
		var ht *zend.HashTable = SplArrayGetHashTable(intern)
	try_again:
		switch offset.u1.v.type_ {
		case 6:
			if g.Assign(&tmp, zend.ZendSymtableFind(ht, offset.value.str)) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
			break
		case 5:
			index = zend_long(*offset).value.dval
			goto num_index
		case 9:
			index = offset.value.res.handle
			goto num_index
		case 2:
			index = 0
			goto num_index
		case 3:
			index = 1
			goto num_index
		case 4:
			index = offset.value.lval
		num_index:
			if g.Assign(&tmp, zend.ZendHashIndexFind(ht, index)) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
			break
		case 10:
			if offset.u1.v.type_ == 10 {
				offset = &(*offset).value.ref.val
			}
			goto try_again
		default:
			zend.ZendError(1<<1, "Illegal offset type")
			return 0
		}
		if check_empty != 0 && check_inherited != 0 && intern.GetFptrOffsetGet() != nil {
			value = SplArrayReadDimensionEx(1, object, offset, 0, &rv)
		} else {
			value = tmp
		}
	}
	var result zend.ZendBool = g.CondF(check_empty != 0, func() int { return zend.ZendIsTrue(value) }, func() bool { return value.u1.v.type_ != 1 })
	if value == &rv {
		zend.ZvalPtrDtor(&rv)
	}
	return result
}
func SplArrayHasDimension(object *zend.Zval, offset *zend.Zval, check_empty int) int {
	return SplArrayHasDimensionEx(1, object, offset, check_empty)
}

/* {{{ proto bool ArrayObject::offsetExists(mixed $index)
    proto bool ArrayIterator::offsetExists(mixed $index)
Returns whether the requested $index exists. */

func zim_spl_Array_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &index) == zend.FAILURE {
		return
	}
	if SplArrayHasDimensionEx(0, &(execute_data.This), index, 2) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed ArrayObject::offsetGet(mixed $index)
    proto mixed ArrayIterator::offsetGet(mixed $index)
Returns the value at the specified $index. */

func zim_spl_Array_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var index *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &index) == zend.FAILURE {
		return
	}
	value = SplArrayReadDimensionEx(0, &(execute_data.This), index, 0, return_value)
	if value != return_value {
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
	}
}

/* {{{ proto void ArrayObject::offsetSet(mixed $index, mixed $newval)
    proto void ArrayIterator::offsetSet(mixed $index, mixed $newval)
Sets the value at the specified $index to $newval. */

func zim_spl_Array_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	var value *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &index, &value) == zend.FAILURE {
		return
	}
	SplArrayWriteDimensionEx(0, &(execute_data.This), index, value)
}
func SplArrayIteratorAppend(object *zend.Zval, append_value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if SplArrayIsObject(intern) != 0 {
		zend.ZendThrowError(nil, "Cannot append properties to objects, use %s::offsetSet() instead", object.value.obj.ce.name.val)
		return
	}
	SplArrayWriteDimension(object, nil, append_value)
}

/* {{{ proto void ArrayObject::append(mixed $newval)
    proto void ArrayIterator::append(mixed $newval)
Appends the value (cannot be called for objects). */

func zim_spl_Array_append(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &value) == zend.FAILURE {
		return
	}
	SplArrayIteratorAppend(&(execute_data.This), value)
}

/* {{{ proto void ArrayObject::offsetUnset(mixed $index)
    proto void ArrayIterator::offsetUnset(mixed $index)
Unsets the value at the specified $index. */

func zim_spl_Array_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &index) == zend.FAILURE {
		return
	}
	SplArrayUnsetDimensionEx(0, &(execute_data.This), index)
}

/* {{{ proto array ArrayObject::getArrayCopy()
   proto array ArrayIterator::getArrayCopy()
Return a copy of the contained array */

func zim_spl_Array_getArrayCopy(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var __arr *zend.ZendArray = zend.ZendArrayDup(SplArrayGetHashTable(intern))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}
func SplArrayGetPropertiesFor(object *zend.Zval, purpose zend.ZendPropPurpose) *zend.HashTable {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var ht *zend.HashTable
	var dup zend.ZendBool
	if (intern.GetArFlags() & 0x1) != 0 {
		return zend.ZendStdGetPropertiesFor(object, purpose)
	}

	/* We are supposed to be the only owner of the internal hashtable.
	 * The "dup" flag decides whether this is a "long-term" use where
	 * we need to duplicate, or a "temporary" one, where we can expect
	 * that no operations on the ArrayObject will be performed in the
	 * meantime. */

	switch purpose {
	case zend.ZEND_PROP_PURPOSE_ARRAY_CAST:
		dup = 1
		break
	case zend.ZEND_PROP_PURPOSE_VAR_EXPORT:

	case zend.ZEND_PROP_PURPOSE_JSON:

	case zend._ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		dup = 0
		break
	default:
		return zend.ZendStdGetPropertiesFor(object, purpose)
	}
	ht = SplArrayGetHashTable(intern)
	if dup != 0 {
		ht = zend.ZendArrayDup(ht)
	} else {
		zend.ZendGcAddref(&ht.gc)
	}
	return ht
}
func SplArrayGetDebugInfo(obj *zend.Zval) *zend.HashTable {
	var storage *zend.Zval
	var zname *zend.ZendString
	var base *zend.ZendClassEntry
	var intern *SplArrayObject = SplArrayFromObj(obj.value.obj)
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	if (intern.GetArFlags() & 0x1000000) != 0 {
		return zend.ZendArrayDup(intern.std.properties)
	} else {
		var debug_info *zend.HashTable
		debug_info = zend._zendNewArray(intern.std.properties.nNumOfElements + 1)
		zend.ZendHashCopy(debug_info, intern.std.properties, zend.CopyCtorFuncT(zend.ZvalAddRef))
		storage = &intern.array
		if storage.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(storage)
		}
		if obj.value.obj.handlers == &spl_handler_ArrayIterator {
			base = spl_ce_ArrayIterator
		} else {
			base = spl_ce_ArrayObject
		}
		zname = SplGenPrivatePropName(base, "storage", g.SizeOf("\"storage\"")-1)
		zend.ZendSymtableUpdate(debug_info, zname, storage)
		zend.ZendStringReleaseEx(zname, 0)
		return debug_info
	}
}

/* }}} */

func SplArrayGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplArrayObject = SplArrayFromObj(obj.value.obj)
	*gc_data = &intern.array
	*gc_data_count = 1
	return zend.ZendStdGetProperties(obj)
}

/* }}} */

func SplArrayReadProperty(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any, rv *zend.Zval) *zend.Zval {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if (intern.GetArFlags()&0x2) != 0 && zend.ZendStdHasProperty(object, member, 0x2, nil) == 0 {
		return SplArrayReadDimension(object, member, type_, rv)
	}
	return zend.ZendStdReadProperty(object, member, type_, cache_slot, rv)
}
func SplArrayWriteProperty(object *zend.Zval, member *zend.Zval, value *zend.Zval, cache_slot *any) *zend.Zval {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if (intern.GetArFlags()&0x2) != 0 && zend.ZendStdHasProperty(object, member, 0x2, nil) == 0 {
		SplArrayWriteDimension(object, member, value)
		return value
	}
	return zend.ZendStdWriteProperty(object, member, value, cache_slot)
}
func SplArrayGetPropertyPtrPtr(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any) *zend.Zval {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if (intern.GetArFlags()&0x2) != 0 && zend.ZendStdHasProperty(object, member, 0x2, nil) == 0 {

		/* If object has offsetGet() overridden, then fallback to read_property,
		 * which will call offsetGet(). */

		if intern.GetFptrOffsetGet() != nil {
			return nil
		}
		return SplArrayGetDimensionPtr(1, intern, member, type_)
	}
	return zend.ZendStdGetPropertyPtrPtr(object, member, type_, cache_slot)
}
func SplArrayHasProperty(object *zend.Zval, member *zend.Zval, has_set_exists int, cache_slot *any) int {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if (intern.GetArFlags()&0x2) != 0 && zend.ZendStdHasProperty(object, member, 0x2, nil) == 0 {
		return SplArrayHasDimension(object, member, has_set_exists)
	}
	return zend.ZendStdHasProperty(object, member, has_set_exists, cache_slot)
}
func SplArrayUnsetProperty(object *zend.Zval, member *zend.Zval, cache_slot *any) {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if (intern.GetArFlags()&0x2) != 0 && zend.ZendStdHasProperty(object, member, 0x2, nil) == 0 {
		SplArrayUnsetDimension(object, member)
		return
	}
	zend.ZendStdUnsetProperty(object, member, cache_slot)
}
func SplArrayCompareObjects(o1 *zend.Zval, o2 *zend.Zval) int {
	var ht1 *zend.HashTable
	var ht2 *zend.HashTable
	var intern1 *SplArrayObject
	var intern2 *SplArrayObject
	var result int = 0
	intern1 = SplArrayFromObj(o1.value.obj)
	intern2 = SplArrayFromObj(o2.value.obj)
	ht1 = SplArrayGetHashTable(intern1)
	ht2 = SplArrayGetHashTable(intern2)
	result = zend.ZendCompareSymbolTables(ht1, ht2)

	/* if we just compared std.properties, don't do it again */

	if result == 0 && !(ht1 == intern1.std.properties && ht2 == intern2.std.properties) {
		result = zend.ZendStdCompareObjects(o1, o2)
	}
	return result
}
func SplArraySkipProtected(intern *SplArrayObject, aht *zend.HashTable) int {
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var data *zend.Zval
	if SplArrayIsObject(intern) != 0 {
		var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
		for {
			if zend.ZendHashGetCurrentKeyEx(aht, &string_key, &num_key, pos_ptr) == 1 {
				data = zend.ZendHashGetCurrentDataEx(aht, pos_ptr)
				if data != nil && data.u1.v.type_ == 13 && g.Assign(&data, data.value.zv).u1.v.type_ == 0 {

				} else if string_key.len_ == 0 || string_key.val[0] {
					return zend.SUCCESS
				}
			} else {
				return zend.SUCCESS
			}
			if g.Cond(zend.ZendHashGetCurrentKeyTypeEx(aht, pos_ptr) == 3, zend.FAILURE, zend.SUCCESS) != zend.SUCCESS {
				return zend.FAILURE
			}
			zend.ZendHashMoveForwardEx(aht, pos_ptr)

		}
	}
	return zend.FAILURE
}
func SplArrayNextEx(intern *SplArrayObject, aht *zend.HashTable) int {
	var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
	zend.ZendHashMoveForwardEx(aht, pos_ptr)
	if SplArrayIsObject(intern) != 0 {
		return SplArraySkipProtected(intern, aht)
	} else {
		if zend.ZendHashGetCurrentKeyTypeEx(aht, pos_ptr) == 3 {
			return zend.FAILURE
		} else {
			return zend.SUCCESS
		}
	}
}
func SplArrayNext(intern *SplArrayObject) int {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	return SplArrayNextEx(intern, aht)
}
func SplArrayItDtor(iter *zend.ZendObjectIterator) {
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iter.data)
}

/* }}} */

func SplArrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplArrayObject = SplArrayFromObj(&iter.data.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if (object.GetArFlags() & 0x20000) != 0 {
		return zend.ZendUserItValid(iter)
	} else {
		if zend.ZendHashGetCurrentKeyTypeEx(aht, SplArrayGetPosPtr(aht, object)) == 3 {
			return zend.FAILURE
		} else {
			return zend.SUCCESS
		}
	}
}

/* }}} */

func SplArrayItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplArrayObject = SplArrayFromObj(&iter.data.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if (object.GetArFlags() & 0x80000) != 0 {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *zend.Zval = zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, object))
		if data != nil && data.u1.v.type_ == 13 {
			data = data.value.zv
		}
		return data
	}
}

/* }}} */

func SplArrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplArrayObject = SplArrayFromObj(&iter.data.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if (object.GetArFlags() & 0x40000) != 0 {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		zend.ZendHashGetCurrentKeyZvalEx(aht, key, SplArrayGetPosPtr(aht, object))
	}
}

/* }}} */

func SplArrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = SplArrayFromObj(&iter.data.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if (object.GetArFlags() & 0x100000) != 0 {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayNextEx(object, aht)
	}
}

/* }}} */

func SplArrayRewind(intern *SplArrayObject) {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if intern.GetHtIter() == uint32-1 {
		SplArrayGetPosPtr(aht, intern)
	} else {
		zend.ZendHashInternalPointerResetEx(aht, SplArrayGetPosPtr(aht, intern))
		SplArraySkipProtected(intern, aht)
	}
}

/* }}} */

func SplArrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = SplArrayFromObj(&iter.data.value.obj)
	if (object.GetArFlags() & 0x10000) != 0 {
		zend.ZendUserItRewind(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayRewind(object)
	}
}

/* }}} */

func SplArraySetArray(object *zend.Zval, intern *SplArrayObject, array *zend.Zval, ar_flags zend.ZendLong, just_array int) {
	if array.u1.v.type_ != 8 && array.u1.v.type_ != 7 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Passed variable is not an array or object", 0)
		return
	}
	if array.u1.v.type_ == 7 {
		zend.ZvalPtrDtor(&intern.array)
		if zend.ZvalRefcountP(array) == 1 {
			var _z1 *zend.Zval = &intern.array
			var _z2 *zend.Zval = array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		} else {

			//??? TODO: try to avoid array duplication

			var __arr *zend.ZendArray = zend.ZendArrayDup(array.value.arr)
			var __z *zend.Zval = &intern.array
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

			//??? TODO: try to avoid array duplication

		}
	} else {
		if array.value.obj.handlers == &spl_handler_ArrayObject || array.value.obj.handlers == &spl_handler_ArrayIterator {
			zend.ZvalPtrDtor(&intern.array)
			if just_array != 0 {
				var other *SplArrayObject = SplArrayFromObj(array.value.obj)
				ar_flags = other.GetArFlags() & ^0xffff0000
			}
			if object.value.obj == array.value.obj {
				ar_flags |= 0x1000000
				&intern.array.u1.type_info = 0
			} else {
				ar_flags |= 0x2000000
				var _z1 *zend.Zval = &intern.array
				var _z2 *zend.Zval = array
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				if (_t & 0xff00) != 0 {
					zend.ZendGcAddref(&_gc.gc)
				}
			}
		} else {
			var handler zend.ZendObjectGetPropertiesT = array.value.obj.handlers.get_properties
			if handler != zend.ZendStdGetProperties {
				zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Overloaded object of type %s is not compatible with %s", array.value.obj.ce.name.val, intern.std.ce.name.val)
				return
			}
			zend.ZvalPtrDtor(&intern.array)
			var _z1 *zend.Zval = &intern.array
			var _z2 *zend.Zval = array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		}
	}
	intern.SetArFlags(intern.GetArFlags() &^ 0x1000000 & ^0x2000000)
	intern.SetArFlags(intern.GetArFlags() | ar_flags)
	intern.SetHtIter(uint32 - 1)
}

/* }}} */

var SplArrayItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplArrayItDtor, SplArrayItValid, SplArrayItGetCurrentData, SplArrayItGetCurrentKey, SplArrayItMoveForward, SplArrayItRewind, nil}

func SplArrayGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *zend.ZendUserIterator
	var array_object *SplArrayObject = SplArrayFromObj(object.value.obj)
	if by_ref != 0 && (array_object.GetArFlags()&0x80000) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend._emalloc(g.SizeOf("zend_user_iterator"))
	zend.ZendIteratorInit(&iterator.it)
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.it.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.it.funcs = &SplArrayItFuncs
	iterator.ce = ce
	&iterator.value.u1.type_info = 0
	return &iterator.it
}

/* }}} */

func zim_spl_Array___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject
	var array *zend.Zval
	var ar_flags zend.ZendLong = 0
	var ce_get_iterator *zend.ZendClassEntry = spl_ce_ArrayIterator
	if execute_data.This.u2.num_args == 0 {
		return
	}
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "z|lC", &array, &ar_flags, &ce_get_iterator) == zend.FAILURE {
		return
	}
	intern = SplArrayFromObj(object.value.obj)
	if execute_data.This.u2.num_args > 2 {
		intern.SetCeGetIterator(ce_get_iterator)
	}
	ar_flags &= ^0xffff0000
	SplArraySetArray(object, intern, array, ar_flags, execute_data.This.u2.num_args == 1)
}

/* }}} */

func zim_spl_ArrayIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject
	var array *zend.Zval
	var ar_flags zend.ZendLong = 0
	if execute_data.This.u2.num_args == 0 {
		return
	}
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "z|l", &array, &ar_flags) == zend.FAILURE {
		return
	}
	intern = SplArrayFromObj(object.value.obj)
	ar_flags &= ^0xffff0000
	SplArraySetArray(object, intern, array, ar_flags, execute_data.This.u2.num_args == 1)
}

/* }}} */

func zim_spl_Array_setIteratorClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var ce_get_iterator *zend.ZendClassEntry = spl_ce_ArrayIterator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgClass(_arg, &ce_get_iterator, _i, 0) == 0 {
				_error_code = 1
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	intern.SetCeGetIterator(ce_get_iterator)
}

/* }}} */

func zim_spl_Array_getIteratorClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendStringAddref(intern.GetCeGetIterator().name)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = intern.GetCeGetIterator().name
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func zim_spl_Array_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetArFlags() & ^0xffff0000
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_Array_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var ar_flags zend.ZendLong = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &ar_flags) == zend.FAILURE {
		return
	}
	intern.SetArFlags(intern.GetArFlags()&0xffff0000 | ar_flags & ^0xffff0000)
}

/* }}} */

func zim_spl_Array_exchangeArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var array *zend.Zval
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &array) == zend.FAILURE {
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(1<<1, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	var __arr *zend.ZendArray = zend.ZendArrayDup(SplArrayGetHashTable(intern))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplArraySetArray(object, intern, array, 0, 1)
}

/* }}} */

func zim_spl_Array_getIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.obj = SplArrayObjectNewEx(intern.GetCeGetIterator(), object, 0)
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
}

/* }}} */

func zim_spl_Array_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplArrayRewind(intern)
}

/* }}} */

func zim_spl_Array_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var opos zend.ZendLong
	var position zend.ZendLong
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	var result int
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &position) == zend.FAILURE {
		return
	}
	opos = position
	if position >= 0 {
		SplArrayRewind(intern)
		result = zend.SUCCESS
		for g.PostDec(&position) > 0 && g.Assign(&result, SplArrayNext(intern)) == zend.SUCCESS {

		}
		if result == zend.SUCCESS && g.Cond(zend.ZendHashGetCurrentKeyTypeEx(aht, SplArrayGetPosPtr(aht, intern)) == 3, zend.FAILURE, zend.SUCCESS) == zend.SUCCESS {
			return
		}
	}
	zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+"%"+"lld"+" is out of range", opos)
}
func SplArrayObjectCountElementsHelper(intern *SplArrayObject) zend.ZendLong {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if SplArrayIsObject(intern) != 0 {
		var count zend.ZendLong = 0
		var key *zend.ZendString
		var val *zend.Zval

		/* Count public/dynamic properties */

		for {
			var __ht *zend.HashTable = aht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				key = _p.key
				val = _z
				if val.u1.v.type_ == 13 {
					if val.value.zv.u1.v.type_ == 0 {
						continue
					}
					if key != nil && key.val[0] == '0' {
						continue
					}
				}
				count++
			}
			break
		}
		return count
	} else {
		return aht.nNumOfElements
	}
}
func SplArrayObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_count, "count", g.SizeOf("\"count\"")-1, &rv, 0, nil, nil)
		if rv.u1.v.type_ != 0 {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return zend.SUCCESS
		}
		*count = 0
		return zend.FAILURE
	}
	*count = SplArrayObjectCountElementsHelper(intern)
	return zend.SUCCESS
}

/* {{{ proto int ArrayObject::count()
    proto int ArrayIterator::count()
Return the number of elements in the Iterator. */

func zim_spl_Array_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = SplArrayObjectCountElementsHelper(intern)
	__z.u1.type_info = 4
	return
}
func SplArrayMethod(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fname string, fname_len int, use_arg int) {
	var intern *SplArrayObject = SplArrayFromObj(&(execute_data.This).value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	var function_name zend.Zval
	var params []zend.Zval
	var arg *zend.Zval = nil
	var __z *zend.Zval = &function_name
	var __s *zend.ZendString = zend.ZendStringInit(fname, fname_len, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
	zend.ZendGcSetRefcount(&_ref.gc, 1)
	_ref.gc.u.type_info = 10
	_ref.sources.ptr = nil
	&params[0].value.ref = _ref
	&params[0].u1.type_info = 10 | 1<<0<<8
	var __arr *zend.ZendArray = aht
	var __z *zend.Zval = &params[0].value.ref.val
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZendGcAddref(&aht.gc)
	if use_arg == 0 {
		intern.GetNApplyCount()++
		zend._callUserFunctionEx(nil, &function_name, return_value, 1, params, 1)
		intern.GetNApplyCount()--
	} else if use_arg == 2 {
		if zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "|z", &arg) == zend.FAILURE {
			zend.ZendThrowException(spl_ce_BadMethodCallException, "Function expects one argument at most", 0)
			goto exit
		}
		if arg != nil {
			var _z1 *zend.Zval = &params[1]
			var _z2 *zend.Zval = arg
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		}
		intern.GetNApplyCount()++
		zend._callUserFunctionEx(nil, &function_name, return_value, g.Cond(arg != nil, 2, 1), params, 1)
		intern.GetNApplyCount()--
	} else {
		if execute_data.This.u2.num_args != 1 || zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "z", &arg) == zend.FAILURE {
			zend.ZendThrowException(spl_ce_BadMethodCallException, "Function expects exactly one argument", 0)
			goto exit
		}
		var _z1 *zend.Zval = &params[1]
		var _z2 *zend.Zval = arg
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		intern.GetNApplyCount()++
		zend._callUserFunctionEx(nil, &function_name, return_value, 2, params, 1)
		intern.GetNApplyCount()--
	}
exit:
	var new_ht *zend.HashTable = &params[0].value.ref.val.value.arr
	if aht != new_ht {
		SplArrayReplaceHashTable(intern, new_ht)
	} else {
		zend.ZendGcDelref(&aht.gc)
	}
	&params[0].value.ref.val.u1.type_info = 1
	zend.ZvalPtrDtor(&params[0])
	zend.ZendStringFree(function_name.value.str)
}

// #define SPL_ARRAY_METHOD(cname,fname,use_arg) SPL_METHOD ( cname , fname ) { spl_array_method ( INTERNAL_FUNCTION_PARAM_PASSTHRU , # fname , sizeof ( # fname ) - 1 , use_arg ) ; }

/* {{{ proto int ArrayObject::asort([int $sort_flags = SORT_REGULAR ])
    proto int ArrayIterator::asort([int $sort_flags = SORT_REGULAR ])
Sort the entries by values. */

func zim_spl_Array_asort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "asort", g.SizeOf("\"asort\"")-1, 2)
}

/* {{{ proto int ArrayObject::ksort([int $sort_flags = SORT_REGULAR ])
    proto int ArrayIterator::ksort([int $sort_flags = SORT_REGULAR ])
Sort the entries by key. */

func zim_spl_Array_ksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "ksort", g.SizeOf("\"ksort\"")-1, 2)
}

/* {{{ proto int ArrayObject::uasort(callback cmp_function)
    proto int ArrayIterator::uasort(callback cmp_function)
Sort the entries by values user defined function. */

func zim_spl_Array_uasort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "uasort", g.SizeOf("\"uasort\"")-1, 1)
}

/* {{{ proto int ArrayObject::uksort(callback cmp_function)
    proto int ArrayIterator::uksort(callback cmp_function)
Sort the entries by key using user defined function. */

func zim_spl_Array_uksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "uksort", g.SizeOf("\"uksort\"")-1, 1)
}

/* {{{ proto int ArrayObject::natsort()
    proto int ArrayIterator::natsort()
Sort the entries by values using "natural order" algorithm. */

func zim_spl_Array_natsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "natsort", g.SizeOf("\"natsort\"")-1, 0)
}

/* {{{ proto int ArrayObject::natcasesort()
    proto int ArrayIterator::natcasesort()
Sort the entries by key using case insensitive "natural order" algorithm. */

func zim_spl_Array_natcasesort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "natcasesort", g.SizeOf("\"natcasesort\"")-1, 0)
}

/* {{{ proto mixed|NULL ArrayIterator::current()
   Return current array entry */

func zim_spl_Array_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var entry *zend.Zval
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if entry.u1.v.type_ == 13 {
		entry = entry.value.zv
		if entry.u1.v.type_ == 0 {
			return
		}
	}
	var _z3 *zend.Zval = entry
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
}

/* }}} */

func zim_spl_Array_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplArrayIteratorKey(&(execute_data.This), return_value)
}
func SplArrayIteratorKey(object *zend.Zval, return_value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	zend.ZendHashGetCurrentKeyZvalEx(aht, return_value, SplArrayGetPosPtr(aht, intern))
}

/* }}} */

func zim_spl_Array_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplArrayNextEx(intern, aht)
}

/* }}} */

func zim_spl_Array_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Cond(zend.ZendHashGetCurrentKeyTypeEx(aht, SplArrayGetPosPtr(aht, intern)) == 3, zend.FAILURE, zend.SUCCESS) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_Array_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var entry *zend.Zval
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if entry.u1.v.type_ == 13 {
		entry = entry.value.zv
	}
	if entry.u1.v.type_ == 10 {
		entry = &(*entry).value.ref.val
	}
	if entry.u1.v.type_ == 7 || entry.u1.v.type_ == 8 && (intern.GetArFlags()&0x4) == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_Array_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var entry *zend.Zval
	var flags zend.Zval
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if entry.u1.v.type_ == 13 {
		entry = entry.value.zv
	}
	if entry.u1.v.type_ == 10 {
		entry = &(*entry).value.ref.val
	}
	if entry.u1.v.type_ == 8 {
		if (intern.GetArFlags() & 0x4) != 0 {
			return
		}
		if zend.InstanceofFunction(entry.value.obj.ce, &(execute_data.This).value.obj.ce) != 0 {
			var __z *zend.Zval = return_value
			__z.value.obj = entry.value.obj
			__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
			zend.ZvalAddrefP(return_value)
			return
		}
	}
	var __z *zend.Zval = &flags
	__z.value.lval = intern.GetArFlags()
	__z.u1.type_info = 4
	SplInstantiateArgEx2(&(execute_data.This).value.obj.ce, return_value, entry, &flags)
}

/* }}} */

func zim_spl_Array_serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var members zend.Zval
	var flags zend.Zval
	var var_hash standard.PhpSerializeDataT
	var buf zend.SmartStr = zend.SmartStr{0}
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var_hash = standard.PhpVarSerializeInit()
	var __z *zend.Zval = &flags
	__z.value.lval = intern.GetArFlags() & 0x100ffff
	__z.u1.type_info = 4

	/* storage */

	zend.SmartStrAppendlEx(&buf, "x:", 2, 0)
	standard.PhpVarSerialize(&buf, &flags, &var_hash)
	if (intern.GetArFlags() & 0x1000000) == 0 {
		standard.PhpVarSerialize(&buf, &intern.array, &var_hash)
		zend.SmartStrAppendcEx(&buf, ';', 0)
	}

	/* members */

	zend.SmartStrAppendlEx(&buf, "m:", 2, 0)
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	var __arr *zend.ZendArray = intern.std.properties
	var __z *zend.Zval = &members
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	standard.PhpVarSerialize(&buf, &members, &var_hash)

	/* done */

	standard.PhpVarSerializeDestroy(var_hash)
	if buf.s != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = buf.s
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	return_value.u1.type_info = 1
	return
}

/* {{{ proto void ArrayObject::unserialize(string serialized)
 * unserialize the object
 */

func zim_spl_Array_unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplArrayObject = SplArrayFromObj(object.value.obj)
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var members *zend.Zval
	var zflags *zend.Zval
	var array *zend.Zval
	var flags zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(1<<1, "Modification of ArrayObject during sorting is prohibited")
		return
	}

	/* storage */

	p = (*uint8)(buf)
	s = p
	var_hash = standard.PhpVarUnserializeInit()
	if (*p) != 'x' || (*(g.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	zflags = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(zflags, &p, s+buf_len, &var_hash) == 0 || zflags.u1.v.type_ != 4 {
		goto outexcept
	}
	p--
	flags = zflags.value.lval

	/* flags needs to be verified and we also need to verify whether the next
	 * thing we get is ';'. After that we require an 'm' or something else
	 * where 'm' stands for members and anything else should be an array. If
	 * neither 'a' or 'm' follows we have an error. */

	if (*p) != ';' {
		goto outexcept
	}
	p++
	if (flags & 0x1000000) != 0 {

		/* If IS_SELF is used, the flags are not followed by an array/object */

		intern.SetArFlags(intern.GetArFlags() &^ 0x100ffff)
		intern.SetArFlags(intern.GetArFlags() | flags&0x100ffff)
		zend.ZvalPtrDtor(&intern.array)
		&intern.array.u1.type_info = 0
	} else {
		if (*p) != 'a' && (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}
		array = standard.VarTmpVar(&var_hash)
		if standard.PhpVarUnserialize(array, &p, s+buf_len, &var_hash) == 0 || array.u1.v.type_ != 7 && array.u1.v.type_ != 8 {
			goto outexcept
		}
		intern.SetArFlags(intern.GetArFlags() &^ 0x100ffff)
		intern.SetArFlags(intern.GetArFlags() | flags&0x100ffff)
		if array.u1.v.type_ == 7 {
			zend.ZvalPtrDtor(&intern.array)
			var _z1 *zend.Zval = &intern.array
			var _z2 *zend.Zval = array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			array.u1.type_info = 1
			var _zv *zend.Zval = &intern.array
			var _arr *zend.ZendArray = _zv.value.arr
			if zend.ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.u1.v.type_flags != 0 {
					zend.ZendGcDelref(&_arr.gc)
				}
				var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
				var __z *zend.Zval = _zv
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			}
		} else {
			SplArraySetArray(object, intern, array, 0, 1)
		}
		if (*p) != ';' {
			goto outexcept
		}
		p++
	}

	/* members */

	if (*p) != 'm' || (*(g.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	members = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(members, &p, s+buf_len, &var_hash) == 0 || members.u1.v.type_ != 7 {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(&intern.std, members.value.arr)

	/* done reading $serialized */

	standard.PhpVarUnserializeDestroy(var_hash)
	return
outexcept:
	standard.PhpVarUnserializeDestroy(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset "+"%"+"lld"+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
	return
}

/* {{{ proto array ArrayObject::__serialize() */

func zim_spl_Array___serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(&(execute_data.This).value.obj)
	var tmp zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneException()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* flags */

	var __z *zend.Zval = &tmp
	__z.value.lval = intern.GetArFlags() & 0x100ffff
	__z.u1.type_info = 4
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)

	/* storage */

	if (intern.GetArFlags() & 0x1000000) != 0 {
		&tmp.u1.type_info = 1
	} else {
		var _z1 *zend.Zval = &tmp
		var _z2 *zend.Zval = &intern.array
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	}
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)

	/* members */

	var __arr *zend.ZendArray = zend.ZendStdGetProperties(&(execute_data.This))
	var __z *zend.Zval = &tmp
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if &tmp.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&tmp)
	}
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)

	/* iterator class */

	if intern.GetCeGetIterator() == spl_ce_ArrayIterator {
		&tmp.u1.type_info = 1
	} else {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = intern.GetCeGetIterator().name
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)
}

/* }}} */

func zim_spl_Array___unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = SplArrayFromObj(&(execute_data.This).value.obj)
	var data *zend.HashTable
	var flags_zv *zend.Zval
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var iterator_class_zv *zend.Zval
	var flags zend.ZendLong
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "h", &data) == zend.FAILURE {
		return
	}
	flags_zv = zend.ZendHashIndexFind(data, 0)
	storage_zv = zend.ZendHashIndexFind(data, 1)
	members_zv = zend.ZendHashIndexFind(data, 2)
	iterator_class_zv = zend.ZendHashIndexFind(data, 3)
	if flags_zv == nil || storage_zv == nil || members_zv == nil || flags_zv.u1.v.type_ != 4 || members_zv.u1.v.type_ != 7 || iterator_class_zv != nil && (iterator_class_zv.u1.v.type_ != 1 && iterator_class_zv.u1.v.type_ != 6) {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	flags = flags_zv.value.lval
	intern.SetArFlags(intern.GetArFlags() &^ 0x100ffff)
	intern.SetArFlags(intern.GetArFlags() | flags&0x100ffff)
	if (flags & 0x1000000) != 0 {
		zend.ZvalPtrDtor(&intern.array)
		&intern.array.u1.type_info = 0
	} else {
		SplArraySetArray(&(execute_data.This), intern, storage_zv, 0, 1)
	}
	zend.ObjectPropertiesLoad(&intern.std, members_zv.value.arr)
	if iterator_class_zv != nil && iterator_class_zv.u1.v.type_ == 6 {
		var ce *zend.ZendClassEntry = zend.ZendLookupClass(iterator_class_zv.value.str)
		if ce == nil {
			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; no such class exists", iterator_class_zv.value.str.val)
			return
		} else if zend.InstanceofFunction(ce, zend.ZendCeIterator) == 0 {
			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; this class does not implement the Iterator interface", iterator_class_zv.value.str.val)
			return
		} else {
			intern.SetCeGetIterator(ce)
		}
	}
}

/* }}} */

func zim_spl_Array___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplArrayGetDebugInfo(g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* {{{ arginfo and function table */

var ArginfoArrayConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"input", 0, 0, 0}, {"flags", 0, 0, 0}, {"iterator_class", 0, 0, 0}}

/* ArrayIterator::__construct and ArrayObject::__construct have different signatures */

var ArginfoArrayIteratorConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"array", 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_array_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_array_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var ArginfoArrayAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoArraySeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"position", 0, 0, 0}}
var arginfo_array_exchangeArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"input", 0, 0, 0}}
var arginfo_array_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_array_setIteratorClass []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"iteratorClass", 0, 0, 0}}
var arginfo_array_uXsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"cmp_function", 0, 0, 0}}
var ArginfoArrayUnserialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"serialized", 0, 0, 0}}
var ArginfoArrayVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var spl_funcs_ArrayObject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_Array___construct,
		ArginfoArrayConstruct,
		uint32(g.SizeOf("arginfo_array___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetExists",
		zim_spl_Array_offsetExists,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetGet",
		zim_spl_Array_offsetGet,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetSet",
		zim_spl_Array_offsetSet,
		arginfo_array_offsetSet,
		uint32(g.SizeOf("arginfo_array_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetUnset",
		zim_spl_Array_offsetUnset,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"append",
		zim_spl_Array_append,
		ArginfoArrayAppend,
		uint32(g.SizeOf("arginfo_array_append")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getArrayCopy",
		zim_spl_Array_getArrayCopy,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_Array_count,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_Array_getFlags,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_Array_setFlags,
		arginfo_array_setFlags,
		uint32(g.SizeOf("arginfo_array_setFlags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"asort",
		zim_spl_Array_asort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"ksort",
		zim_spl_Array_ksort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"uasort",
		zim_spl_Array_uasort,
		arginfo_array_uXsort,
		uint32(g.SizeOf("arginfo_array_uXsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"uksort",
		zim_spl_Array_uksort,
		arginfo_array_uXsort,
		uint32(g.SizeOf("arginfo_array_uXsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"natsort",
		zim_spl_Array_natsort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"natcasesort",
		zim_spl_Array_natcasesort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"unserialize",
		zim_spl_Array_unserialize,
		ArginfoArrayUnserialize,
		uint32(g.SizeOf("arginfo_array_unserialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"serialize",
		zim_spl_Array_serialize,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__unserialize",
		zim_spl_Array___unserialize,
		ArginfoArrayUnserialize,
		uint32(g.SizeOf("arginfo_array_unserialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__serialize",
		zim_spl_Array___serialize,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_Array___debugInfo,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getIterator",
		zim_spl_Array_getIterator,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"exchangeArray",
		zim_spl_Array_exchangeArray,
		arginfo_array_exchangeArray,
		uint32(g.SizeOf("arginfo_array_exchangeArray")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setIteratorClass",
		zim_spl_Array_setIteratorClass,
		arginfo_array_setIteratorClass,
		uint32(g.SizeOf("arginfo_array_setIteratorClass")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getIteratorClass",
		zim_spl_Array_getIteratorClass,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_ArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_ArrayIterator___construct,
		ArginfoArrayIteratorConstruct,
		uint32(g.SizeOf("arginfo_array_iterator___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetExists",
		zim_spl_Array_offsetExists,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetGet",
		zim_spl_Array_offsetGet,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetSet",
		zim_spl_Array_offsetSet,
		arginfo_array_offsetSet,
		uint32(g.SizeOf("arginfo_array_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetUnset",
		zim_spl_Array_offsetUnset,
		arginfo_array_offsetGet,
		uint32(g.SizeOf("arginfo_array_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"append",
		zim_spl_Array_append,
		ArginfoArrayAppend,
		uint32(g.SizeOf("arginfo_array_append")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getArrayCopy",
		zim_spl_Array_getArrayCopy,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_Array_count,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_Array_getFlags,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_Array_setFlags,
		arginfo_array_setFlags,
		uint32(g.SizeOf("arginfo_array_setFlags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"asort",
		zim_spl_Array_asort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"ksort",
		zim_spl_Array_ksort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"uasort",
		zim_spl_Array_uasort,
		arginfo_array_uXsort,
		uint32(g.SizeOf("arginfo_array_uXsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"uksort",
		zim_spl_Array_uksort,
		arginfo_array_uXsort,
		uint32(g.SizeOf("arginfo_array_uXsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"natsort",
		zim_spl_Array_natsort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"natcasesort",
		zim_spl_Array_natcasesort,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"unserialize",
		zim_spl_Array_unserialize,
		ArginfoArrayUnserialize,
		uint32(g.SizeOf("arginfo_array_unserialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"serialize",
		zim_spl_Array_serialize,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__unserialize",
		zim_spl_Array___unserialize,
		ArginfoArrayUnserialize,
		uint32(g.SizeOf("arginfo_array_unserialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__serialize",
		zim_spl_Array___serialize,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_Array___debugInfo,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_Array_rewind,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_Array_current,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_Array_key,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_Array_next,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_Array_valid,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"seek",
		zim_spl_Array_seek,
		ArginfoArraySeek,
		uint32(g.SizeOf("arginfo_array_seek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_RecursiveArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"hasChildren",
		zim_spl_Array_hasChildren,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_Array_getChildren,
		ArginfoArrayVoid,
		uint32(g.SizeOf("arginfo_array_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplArray(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_ArrayObject, "ArrayObject", SplArrayObjectNew, spl_funcs_ArrayObject)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, zend.ZendCeAggregate)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, zend.ZendCeArrayaccess)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, zend.ZendCeSerializable)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, zend.ZendCeCountable)
	memcpy(&spl_handler_ArrayObject, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	spl_handler_ArrayObject.offset = zend_long((*byte)(&((*SplArrayObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_ArrayObject.clone_obj = SplArrayObjectClone
	spl_handler_ArrayObject.read_dimension = SplArrayReadDimension
	spl_handler_ArrayObject.write_dimension = SplArrayWriteDimension
	spl_handler_ArrayObject.unset_dimension = SplArrayUnsetDimension
	spl_handler_ArrayObject.has_dimension = SplArrayHasDimension
	spl_handler_ArrayObject.count_elements = SplArrayObjectCountElements
	spl_handler_ArrayObject.get_properties_for = SplArrayGetPropertiesFor
	spl_handler_ArrayObject.get_gc = SplArrayGetGc
	spl_handler_ArrayObject.read_property = SplArrayReadProperty
	spl_handler_ArrayObject.write_property = SplArrayWriteProperty
	spl_handler_ArrayObject.get_property_ptr_ptr = SplArrayGetPropertyPtrPtr
	spl_handler_ArrayObject.has_property = SplArrayHasProperty
	spl_handler_ArrayObject.unset_property = SplArrayUnsetProperty
	spl_handler_ArrayObject.compare_objects = SplArrayCompareObjects
	spl_handler_ArrayObject.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_ArrayObject.free_obj = SplArrayObjectFreeStorage
	SplRegisterStdClass(&spl_ce_ArrayIterator, "ArrayIterator", SplArrayObjectNew, spl_funcs_ArrayIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, zend.ZendCeArrayaccess)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_SeekableIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, zend.ZendCeSerializable)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, zend.ZendCeCountable)
	memcpy(&spl_handler_ArrayIterator, &spl_handler_ArrayObject, g.SizeOf("zend_object_handlers"))
	spl_ce_ArrayIterator.get_iterator = SplArrayGetIterator
	spl_ce_ArrayIterator.ce_flags |= 1 << 18
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "STD_PROP_LIST", g.SizeOf("\"STD_PROP_LIST\"")-1, zend.ZendLong(0x1))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "ARRAY_AS_PROPS", g.SizeOf("\"ARRAY_AS_PROPS\"")-1, zend.ZendLong(0x2))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "STD_PROP_LIST", g.SizeOf("\"STD_PROP_LIST\"")-1, zend.ZendLong(0x1))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "ARRAY_AS_PROPS", g.SizeOf("\"ARRAY_AS_PROPS\"")-1, zend.ZendLong(0x2))
	SplRegisterSubClass(&spl_ce_RecursiveArrayIterator, spl_ce_ArrayIterator, "RecursiveArrayIterator", SplArrayObjectNew, spl_funcs_RecursiveArrayIterator)
	zend.ZendClassImplements(spl_ce_RecursiveArrayIterator, 1, spl_ce_RecursiveIterator)
	spl_ce_RecursiveArrayIterator.get_iterator = SplArrayGetIterator
	spl_ce_RecursiveArrayIterator.ce_flags |= 1 << 18
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveArrayIterator, "CHILD_ARRAYS_ONLY", g.SizeOf("\"CHILD_ARRAYS_ONLY\"")-1, zend.ZendLong(0x4))
	return zend.SUCCESS
}

/* }}} */

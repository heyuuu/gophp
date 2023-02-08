// <<generate>>

package spl

import (
	b "sik/builtin"
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

var spl_handler_SplFixedArray zend.ZendObjectHandlers

const SPL_FIXEDARRAY_OVERLOADED_REWIND = 0x1
const SPL_FIXEDARRAY_OVERLOADED_VALID = 0x2
const SPL_FIXEDARRAY_OVERLOADED_KEY = 0x4
const SPL_FIXEDARRAY_OVERLOADED_CURRENT = 0x8
const SPL_FIXEDARRAY_OVERLOADED_NEXT = 0x10

/* }}}} */

/* }}}} */

/* {{{ proto mixed SplFixedArray::offsetGet(mixed $index)
Returns the value at the specified $index. */

/* {{{ proto void SplFixedArray::offsetSet(mixed $index, mixed $newval)
Sets the value at the specified $index to $newval. */

/* {{{ proto void SplFixedArray::offsetUnset(mixed $index)
Unsets the value at the specified $index. */

var SplFixedarrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFixedarrayItDtor, SplFixedarrayItValid, SplFixedarrayItGetCurrentData, SplFixedarrayItGetCurrentKey, SplFixedarrayItMoveForward, SplFixedarrayItRewind, nil)
var ArginfoSplfixedarrayConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("size", 0, 0, 0),
}
var arginfo_fixedarray_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("index", 0, 0, 0),
}
var arginfo_fixedarray_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("index", 0, 0, 0),
	zend.MakeZendInternalArgInfo("newval", 0, 0, 0),
}
var arginfo_fixedarray_setSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var arginfo_fixedarray_fromArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("array", 0, 0, 0),
	zend.MakeZendInternalArgInfo("save_indexes", 0, 0, 0),
}
var ArginfoSplfixedarrayVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var spl_funcs_SplFixedArray []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_SplFixedArray___construct, ArginfoSplfixedarrayConstruct, uint32(b.SizeOf("arginfo_splfixedarray_construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__wakeup", zim_spl_SplFixedArray___wakeup, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("count", zim_spl_SplFixedArray_count, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("toArray", zim_spl_SplFixedArray_toArray, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fromArray", zim_spl_SplFixedArray_fromArray, arginfo_fixedarray_fromArray, uint32(b.SizeOf("arginfo_fixedarray_fromArray")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_STATIC),
	zend.MakeZendFunctionEntry("getSize", zim_spl_SplFixedArray_getSize, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setSize", zim_spl_SplFixedArray_setSize, arginfo_fixedarray_setSize, uint32(b.SizeOf("arginfo_fixedarray_setSize")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetExists", zim_spl_SplFixedArray_offsetExists, arginfo_fixedarray_offsetGet, uint32(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetGet", zim_spl_SplFixedArray_offsetGet, arginfo_fixedarray_offsetGet, uint32(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetSet", zim_spl_SplFixedArray_offsetSet, arginfo_fixedarray_offsetSet, uint32(b.SizeOf("arginfo_fixedarray_offsetSet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetUnset", zim_spl_SplFixedArray_offsetUnset, arginfo_fixedarray_offsetGet, uint32(b.SizeOf("arginfo_fixedarray_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("rewind", zim_spl_SplFixedArray_rewind, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("current", zim_spl_SplFixedArray_current, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("key", zim_spl_SplFixedArray_key, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("next", zim_spl_SplFixedArray_next, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("valid", zim_spl_SplFixedArray_valid, ArginfoSplfixedarrayVoid, uint32(b.SizeOf("arginfo_splfixedarray_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

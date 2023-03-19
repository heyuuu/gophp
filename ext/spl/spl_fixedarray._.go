// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
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

var spl_ce_SplFixedArray *types.ClassEntry

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
var ArginfoSplfixedarrayConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("size"),
}
var arginfo_fixedarray_offsetGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("index"),
}
var arginfo_fixedarray_offsetSet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("index"),
	zend.MakeArgInfo("newval"),
}
var arginfo_fixedarray_setSize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var arginfo_fixedarray_fromArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("array"),
	zend.MakeArgInfo("save_indexes"),
}
var ArginfoSplfixedarrayVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var spl_funcs_SplFixedArray []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray___construct, ArginfoSplfixedarrayConstruct),
	types.MakeZendFunctionEntryEx("__wakeup", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray___wakeup, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_count, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("toArray", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_toArray, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("fromArray", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_STATIC, zim_spl_SplFixedArray_fromArray, arginfo_fixedarray_fromArray),
	types.MakeZendFunctionEntryEx("getSize", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_getSize, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("setSize", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_setSize, arginfo_fixedarray_setSize),
	types.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_offsetExists, arginfo_fixedarray_offsetGet),
	types.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_offsetGet, arginfo_fixedarray_offsetGet),
	types.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_offsetSet, arginfo_fixedarray_offsetSet),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_offsetUnset, arginfo_fixedarray_offsetGet),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_rewind, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_current, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_key, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_next, ArginfoSplfixedarrayVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplFixedArray_valid, ArginfoSplfixedarrayVoid),
}

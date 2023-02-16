// <<generate>>

package spl

import (
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

var spl_handler_ArrayObject zend.ZendObjectHandlers
var spl_handler_ArrayIterator zend.ZendObjectHandlers

const SPL_ARRAY_STD_PROP_LIST = 0x1
const SPL_ARRAY_ARRAY_AS_PROPS = 0x2
const SPL_ARRAY_CHILD_ARRAYS_ONLY = 0x4
const SPL_ARRAY_OVERLOADED_REWIND = 0x10000
const SPL_ARRAY_OVERLOADED_VALID = 0x20000
const SPL_ARRAY_OVERLOADED_KEY = 0x40000
const SPL_ARRAY_OVERLOADED_CURRENT = 0x80000
const SPL_ARRAY_OVERLOADED_NEXT = 0x100000
const SPL_ARRAY_IS_SELF = 0x1000000
const SPL_ARRAY_USE_OTHER = 0x2000000
const SPL_ARRAY_INT_MASK = 0xffff0000
const SPL_ARRAY_CLONE_MASK = 0x100ffff
const SPL_ARRAY_METHOD_NO_ARG = 0
const SPL_ARRAY_METHOD_USE_ARG = 1
const SPL_ARRAY_METHOD_MAY_USER_ARG = 2

/* {{{ spl_array_object_new_ex */

/* {{{ proto bool ArrayObject::offsetExists(mixed $index)
    proto bool ArrayIterator::offsetExists(mixed $index)
Returns whether the requested $index exists. */

/* {{{ proto mixed ArrayObject::offsetGet(mixed $index)
    proto mixed ArrayIterator::offsetGet(mixed $index)
Returns the value at the specified $index. */

/* {{{ proto void ArrayObject::offsetSet(mixed $index, mixed $newval)
    proto void ArrayIterator::offsetSet(mixed $index, mixed $newval)
Sets the value at the specified $index to $newval. */

/* {{{ proto void ArrayObject::append(mixed $newval)
    proto void ArrayIterator::append(mixed $newval)
Appends the value (cannot be called for objects). */

/* {{{ proto void ArrayObject::offsetUnset(mixed $index)
    proto void ArrayIterator::offsetUnset(mixed $index)
Unsets the value at the specified $index. */

/* {{{ proto array ArrayObject::getArrayCopy()
   proto array ArrayIterator::getArrayCopy()
Return a copy of the contained array */

var SplArrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplArrayItDtor, SplArrayItValid, SplArrayItGetCurrentData, SplArrayItGetCurrentKey, SplArrayItMoveForward, SplArrayItRewind, nil)

/* {{{ proto int ArrayObject::count()
    proto int ArrayIterator::count()
Return the number of elements in the Iterator. */

/* {{{ proto int ArrayObject::asort([int $sort_flags = SORT_REGULAR ])
    proto int ArrayIterator::asort([int $sort_flags = SORT_REGULAR ])
Sort the entries by values. */

/* {{{ proto int ArrayObject::ksort([int $sort_flags = SORT_REGULAR ])
    proto int ArrayIterator::ksort([int $sort_flags = SORT_REGULAR ])
Sort the entries by key. */

/* {{{ proto int ArrayObject::uasort(callback cmp_function)
    proto int ArrayIterator::uasort(callback cmp_function)
Sort the entries by values user defined function. */

/* {{{ proto int ArrayObject::uksort(callback cmp_function)
    proto int ArrayIterator::uksort(callback cmp_function)
Sort the entries by key using user defined function. */

/* {{{ proto int ArrayObject::natsort()
    proto int ArrayIterator::natsort()
Sort the entries by values using "natural order" algorithm. */

/* {{{ proto int ArrayObject::natcasesort()
    proto int ArrayIterator::natcasesort()
Sort the entries by key using case insensitive "natural order" algorithm. */

/* {{{ proto mixed|NULL ArrayIterator::current()
   Return current array entry */

/* {{{ proto void ArrayObject::unserialize(string serialized)
 * unserialize the object
 */

/* {{{ proto array ArrayObject::__serialize() */

/* {{{ arginfo and function table */

var ArginfoArrayConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("iterator_class"),
}

/* ArrayIterator::__construct and ArrayObject::__construct have different signatures */

var ArginfoArrayIteratorConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("array"),
	zend.MakeArgInfo("flags"),
}
var arginfo_array_offsetGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("index"),
}
var arginfo_array_offsetSet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("index"),
	zend.MakeArgInfo("newval"),
}
var ArginfoArrayAppend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var ArginfoArraySeek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("position"),
}
var arginfo_array_exchangeArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("input"),
}
var arginfo_array_setFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("flags"),
}
var arginfo_array_setIteratorClass []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iteratorClass"),
}
var arginfo_array_uXsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("cmp_function"),
}
var ArginfoArrayUnserialize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("serialized"),
}
var ArginfoArrayVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var spl_funcs_ArrayObject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_Array___construct, ArginfoArrayConstruct),
	zend.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetExists, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetGet, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetSet, arginfo_array_offsetSet),
	zend.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetUnset, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("append", zend.ZEND_ACC_PUBLIC, zim_spl_Array_append, ArginfoArrayAppend),
	zend.MakeZendFunctionEntryEx("getArrayCopy", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getArrayCopy, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_Array_count, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getFlags, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_Array_setFlags, arginfo_array_setFlags),
	zend.MakeZendFunctionEntryEx("asort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_asort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("ksort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_ksort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("uasort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_uasort, arginfo_array_uXsort),
	zend.MakeZendFunctionEntryEx("uksort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_uksort, arginfo_array_uXsort),
	zend.MakeZendFunctionEntryEx("natsort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_natsort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("natcasesort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_natcasesort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array_unserialize, ArginfoArrayUnserialize),
	zend.MakeZendFunctionEntryEx("serialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array_serialize, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("__unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array___unserialize, ArginfoArrayUnserialize),
	zend.MakeZendFunctionEntryEx("__serialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array___serialize, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_Array___debugInfo, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("getIterator", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getIterator, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("exchangeArray", zend.ZEND_ACC_PUBLIC, zim_spl_Array_exchangeArray, arginfo_array_exchangeArray),
	zend.MakeZendFunctionEntryEx("setIteratorClass", zend.ZEND_ACC_PUBLIC, zim_spl_Array_setIteratorClass, arginfo_array_setIteratorClass),
	zend.MakeZendFunctionEntryEx("getIteratorClass", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getIteratorClass, ArginfoArrayVoid),
}
var spl_funcs_ArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_ArrayIterator___construct, ArginfoArrayIteratorConstruct),
	zend.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetExists, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetGet, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetSet, arginfo_array_offsetSet),
	zend.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_Array_offsetUnset, arginfo_array_offsetGet),
	zend.MakeZendFunctionEntryEx("append", zend.ZEND_ACC_PUBLIC, zim_spl_Array_append, ArginfoArrayAppend),
	zend.MakeZendFunctionEntryEx("getArrayCopy", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getArrayCopy, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_Array_count, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getFlags, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_Array_setFlags, arginfo_array_setFlags),
	zend.MakeZendFunctionEntryEx("asort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_asort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("ksort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_ksort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("uasort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_uasort, arginfo_array_uXsort),
	zend.MakeZendFunctionEntryEx("uksort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_uksort, arginfo_array_uXsort),
	zend.MakeZendFunctionEntryEx("natsort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_natsort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("natcasesort", zend.ZEND_ACC_PUBLIC, zim_spl_Array_natcasesort, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array_unserialize, ArginfoArrayUnserialize),
	zend.MakeZendFunctionEntryEx("serialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array_serialize, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("__unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array___unserialize, ArginfoArrayUnserialize),
	zend.MakeZendFunctionEntryEx("__serialize", zend.ZEND_ACC_PUBLIC, zim_spl_Array___serialize, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_Array___debugInfo, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_Array_rewind, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_Array_current, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_Array_key, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_Array_next, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_Array_valid, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("seek", zend.ZEND_ACC_PUBLIC, zim_spl_Array_seek, ArginfoArraySeek),
}
var spl_funcs_RecursiveArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_Array_hasChildren, ArginfoArrayVoid),
	zend.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_Array_getChildren, ArginfoArrayVoid),
}

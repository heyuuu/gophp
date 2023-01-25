// <<generate>>

package spl

import (
	b "sik/builtin"
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

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ spl_array_object_new_ex */

/* }}} */

/* }}} */

/* }}} */

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

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

var SplArrayItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplArrayItDtor, SplArrayItValid, SplArrayItGetCurrentData, SplArrayItGetCurrentKey, SplArrayItMoveForward, SplArrayItRewind, nil}

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

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

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto void ArrayObject::unserialize(string serialized)
 * unserialize the object
 */

/* {{{ proto array ArrayObject::__serialize() */

/* }}} */

/* }}} */

/* {{{ arginfo and function table */

var ArginfoArrayConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"input", 0, 0, 0}, {"flags", 0, 0, 0}, {"iterator_class", 0, 0, 0}}

/* ArrayIterator::__construct and ArrayObject::__construct have different signatures */

var ArginfoArrayIteratorConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"array", 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_array_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_array_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var ArginfoArrayAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var ArginfoArraySeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"position", 0, 0, 0},
}
var arginfo_array_exchangeArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"input", 0, 0, 0},
}
var arginfo_array_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"flags", 0, 0, 0},
}
var arginfo_array_setIteratorClass []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"iteratorClass", 0, 0, 0},
}
var arginfo_array_uXsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"cmp_function", 0, 0, 0},
}
var ArginfoArrayUnserialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"serialized", 0, 0, 0},
}
var ArginfoArrayVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var spl_funcs_ArrayObject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_Array___construct,
		ArginfoArrayConstruct,
		uint32_t(b.SizeOf("arginfo_array___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetExists",
		zim_spl_Array_offsetExists,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetGet",
		zim_spl_Array_offsetGet,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetSet",
		zim_spl_Array_offsetSet,
		arginfo_array_offsetSet,
		uint32_t(b.SizeOf("arginfo_array_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetUnset",
		zim_spl_Array_offsetUnset,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"append",
		zim_spl_Array_append,
		ArginfoArrayAppend,
		uint32_t(b.SizeOf("arginfo_array_append")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getArrayCopy",
		zim_spl_Array_getArrayCopy,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_Array_count,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_Array_getFlags,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_Array_setFlags,
		arginfo_array_setFlags,
		uint32_t(b.SizeOf("arginfo_array_setFlags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"asort",
		zim_spl_Array_asort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"ksort",
		zim_spl_Array_ksort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"uasort",
		zim_spl_Array_uasort,
		arginfo_array_uXsort,
		uint32_t(b.SizeOf("arginfo_array_uXsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"uksort",
		zim_spl_Array_uksort,
		arginfo_array_uXsort,
		uint32_t(b.SizeOf("arginfo_array_uXsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"natsort",
		zim_spl_Array_natsort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"natcasesort",
		zim_spl_Array_natcasesort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"unserialize",
		zim_spl_Array_unserialize,
		ArginfoArrayUnserialize,
		uint32_t(b.SizeOf("arginfo_array_unserialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"serialize",
		zim_spl_Array_serialize,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__unserialize",
		zim_spl_Array___unserialize,
		ArginfoArrayUnserialize,
		uint32_t(b.SizeOf("arginfo_array_unserialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__serialize",
		zim_spl_Array___serialize,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_Array___debugInfo,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getIterator",
		zim_spl_Array_getIterator,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"exchangeArray",
		zim_spl_Array_exchangeArray,
		arginfo_array_exchangeArray,
		uint32_t(b.SizeOf("arginfo_array_exchangeArray")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setIteratorClass",
		zim_spl_Array_setIteratorClass,
		arginfo_array_setIteratorClass,
		uint32_t(b.SizeOf("arginfo_array_setIteratorClass")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getIteratorClass",
		zim_spl_Array_getIteratorClass,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_ArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_ArrayIterator___construct,
		ArginfoArrayIteratorConstruct,
		uint32_t(b.SizeOf("arginfo_array_iterator___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetExists",
		zim_spl_Array_offsetExists,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetGet",
		zim_spl_Array_offsetGet,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetSet",
		zim_spl_Array_offsetSet,
		arginfo_array_offsetSet,
		uint32_t(b.SizeOf("arginfo_array_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetUnset",
		zim_spl_Array_offsetUnset,
		arginfo_array_offsetGet,
		uint32_t(b.SizeOf("arginfo_array_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"append",
		zim_spl_Array_append,
		ArginfoArrayAppend,
		uint32_t(b.SizeOf("arginfo_array_append")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getArrayCopy",
		zim_spl_Array_getArrayCopy,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_Array_count,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_Array_getFlags,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_Array_setFlags,
		arginfo_array_setFlags,
		uint32_t(b.SizeOf("arginfo_array_setFlags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"asort",
		zim_spl_Array_asort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"ksort",
		zim_spl_Array_ksort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"uasort",
		zim_spl_Array_uasort,
		arginfo_array_uXsort,
		uint32_t(b.SizeOf("arginfo_array_uXsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"uksort",
		zim_spl_Array_uksort,
		arginfo_array_uXsort,
		uint32_t(b.SizeOf("arginfo_array_uXsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"natsort",
		zim_spl_Array_natsort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"natcasesort",
		zim_spl_Array_natcasesort,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"unserialize",
		zim_spl_Array_unserialize,
		ArginfoArrayUnserialize,
		uint32_t(b.SizeOf("arginfo_array_unserialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"serialize",
		zim_spl_Array_serialize,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__unserialize",
		zim_spl_Array___unserialize,
		ArginfoArrayUnserialize,
		uint32_t(b.SizeOf("arginfo_array_unserialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__serialize",
		zim_spl_Array___serialize,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_Array___debugInfo,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_Array_rewind,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_Array_current,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_Array_key,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_Array_next,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_Array_valid,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"seek",
		zim_spl_Array_seek,
		ArginfoArraySeek,
		uint32_t(b.SizeOf("arginfo_array_seek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_RecursiveArrayIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"hasChildren",
		zim_spl_Array_hasChildren,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_Array_getChildren,
		ArginfoArrayVoid,
		uint32_t(b.SizeOf("arginfo_array_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

/* }}} */

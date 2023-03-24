package types

import "math"

// Source: <Zend/zend_types.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Xinchen Hui <xinchen.h@zend.com>                            |
   +----------------------------------------------------------------------+
*/

type ZendBool = int
type ZendUchar = uint8
type ZEND_RESULT_CODE = int

func IntBool(value bool) ZendBool {
	if value {
		return 1
	} else {
		return 0
	}
}

func ResultCode(value bool) ZEND_RESULT_CODE {
	if value {
		return SUCCESS
	} else {
		return FAILURE
	}
}

const (
	SUCCESS                  = 0
	FAILURE ZEND_RESULT_CODE = -1
)
const ZEND_SIZE_MAX = math.MaxUint

type ZendIntptrT = uintptr
type ZendUintptrT = uintptr

type CompareFuncT func(any, any) int
type SwapFuncT func(any, any)
type SortFuncT func(any, int, int, CompareFuncT, SwapFuncT)
type DtorFuncT func(pDest *Zval)
type CopyCtorFuncT func(pElement *Zval)

const HT_INVALID_IDX uint32 = math.MaxUint32 // uint32(-1)
const HT_MIN_SIZE = 8
const HT_MAX_SIZE = 0x80000000

/* regular data types */
const IS_UNDEF = 0
const IS_NULL = 1
const IS_FALSE = 2
const IS_TRUE = 3
const IS_LONG = 4
const IS_DOUBLE = 5
const IS_STRING = 6
const IS_ARRAY = 7
const IS_OBJECT = 8
const IS_RESOURCE = 9
const IS_REFERENCE = 10

/* constant expressions */

const IS_CONSTANT_AST = 11

/* internal types */

const IS_INDIRECT = 13
const IS_PTR = 14
const IS_ALIAS_PTR = 15
const IS_ERROR = 15 // _IS_ERROR

/* fake types used only for type hinting (Z_TYPE(zv) can not use them) */

const IS_BOOL = 16 // _IS_BOOL
const IS_CALLABLE = 17
const IS_ITERABLE = 18
const IS_VOID = 19
const IS_NUMBER = 20 // _IS_NUMBER

/* we should never set just Z_TYPE, we should set Z_TYPE_INFO */

const Z_TYPE_MASK = 0xff
const Z_TYPE_FLAGS_MASK = 0xff00
const Z_TYPE_FLAGS_SHIFT = 8

/* zval_gc_flags(zval.value->gc.u.type_info) (common flags) */

const GC_COLLECTABLE = 1 << 4
const GC_PROTECTED = 1 << 5
const GC_IMMUTABLE = 1 << 6
const GC_PERSISTENT = 1 << 7
const GC_PERSISTENT_LOCAL = 1 << 8
const GC_ARRAY = IS_ARRAY | GC_COLLECTABLE<<GC_FLAGS_SHIFT
const GC_OBJECT = IS_OBJECT | GC_COLLECTABLE<<GC_FLAGS_SHIFT

/* zval.u1.v.type_flags */

const IS_TYPE_REFCOUNTED = 1 << 0
const IS_TYPE_COLLECTABLE = 1 << 1

/* This optimized version assumes that we have a single "type_flag" */

/* extended types */

const IS_OBJECT_EX uint32 = IS_OBJECT | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT | IS_TYPE_COLLECTABLE<<Z_TYPE_FLAGS_SHIFT
const IS_REFERENCE_EX uint32 = IS_REFERENCE | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT

/* string flags (zval.value->gc.u.flags) */

const IS_STR_INTERNED = GC_IMMUTABLE
const IS_STR_PERSISTENT = GC_PERSISTENT
const IS_STR_PERMANENT uint32 = 1 << 8
const IS_STR_VALID_UTF8 = 1 << 9

/* array flags */

const IS_ARRAY_IMMUTABLE = GC_IMMUTABLE
const IS_ARRAY_PERSISTENT = GC_PERSISTENT

/* object flags (zval.value->gc.u.flags) */

const IS_OBJ_WEAKLY_REFERENCED = GC_PERSISTENT
const IS_OBJ_DESTRUCTOR_CALLED = 1 << 8
const IS_OBJ_FREE_CALLED = 1 << 9

/* Recursion protection macros must be used only for arrays and objects */

/* All data types < IS_STRING have their constructor/destructors skipped */

/* This optimized version assumes that we have a single "type_flag" */

/* deprecated: (COPYABLE is the same as IS_ARRAY) */

/* deprecated: (IMMUTABLE is the same as IS_ARRAY && !REFCOUNTED) */

/* the following Z_OPT_* macros make better code when Z_TYPE_INFO accessed before */

/* deprecated: (COPYABLE is the same as IS_ARRAY) */

/* ZVAL_COPY_OR_DUP() should be used instead of ZVAL_COPY() and ZVAL_DUP()
 * in all places where the source may be a persistent zval.
 */

/* Properties store a flag distinguishing unset and unintialized properties
 * (both use IS_UNDEF type) in the Z_EXTRA space. As such we also need to copy
 * the Z_EXTRA space when copying property default values etc. We define separate __special__
 * macros for this purpose, so this workaround is easier to remove in the future. */

const IS_PROP_UNINIT = 1

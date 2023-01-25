// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/spl/spl_heap.h>

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
   | Authors: Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

var spl_ce_SplHeap *zend.ZendClassEntry
var spl_ce_SplMinHeap *zend.ZendClassEntry
var spl_ce_SplMaxHeap *zend.ZendClassEntry
var spl_ce_SplPriorityQueue *zend.ZendClassEntry

// Source: <ext/spl/spl_heap.c>

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
   | Authors: Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

const PTR_HEAP_BLOCK_SIZE = 64
const SPL_HEAP_CORRUPTED = 0x1
const SPL_PQUEUE_EXTR_MASK = 0x3
const SPL_PQUEUE_EXTR_BOTH = 0x3
const SPL_PQUEUE_EXTR_DATA = 0x1
const SPL_PQUEUE_EXTR_PRIORITY = 0x2

var spl_handler_SplHeap zend.ZendObjectHandlers
var spl_handler_SplPriorityQueue zend.ZendObjectHandlers

type SplPtrHeapDtorFunc func(any)
type SplPtrHeapCtorFunc func(any)
type SplPtrHeapCmpFunc func(any, any, *zend.Zval) int

/* define an __special__  overloaded iterator structure */

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

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto void SplPriorityQueue::__debugInfo() */

/* iterator handler table */

var SplHeapItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplHeapItDtor, SplHeapItValid, SplHeapItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil}
var SplPqueueItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplHeapItDtor, SplHeapItValid, SplPqueueItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil}

/* }}} */

/* }}} */

var ArginfoHeapInsert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var ArginfoHeapCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value1", 0, 0, 0},
	{"value2", 0, 0, 0},
}
var ArginfoPqueueInsert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
	{"priority", 0, 0, 0},
}
var ArginfoPqueueSetflags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"flags", 0, 0, 0},
}
var ArginfoSplheapVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var spl_funcs_SplMinHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplMinHeap_compare,
		ArginfoHeapCompare,
		uint32_t(b.SizeOf("arginfo_heap_compare")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PROTECTED,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplMaxHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplMaxHeap_compare,
		ArginfoHeapCompare,
		uint32_t(b.SizeOf("arginfo_heap_compare")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PROTECTED,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplPriorityQueue []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplPriorityQueue_compare,
		ArginfoHeapCompare,
		uint32_t(b.SizeOf("arginfo_heap_compare")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"insert",
		zim_spl_SplPriorityQueue_insert,
		ArginfoPqueueInsert,
		uint32_t(b.SizeOf("arginfo_pqueue_insert")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setExtractFlags",
		zim_spl_SplPriorityQueue_setExtractFlags,
		ArginfoPqueueSetflags,
		uint32_t(b.SizeOf("arginfo_pqueue_setflags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getExtractFlags",
		zim_spl_SplPriorityQueue_getExtractFlags,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"top",
		zim_spl_SplPriorityQueue_top,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"extract",
		zim_spl_SplPriorityQueue_extract,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_SplHeap_count,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isEmpty",
		zim_spl_SplHeap_isEmpty,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_SplHeap_rewind,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_SplPriorityQueue_current,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_SplHeap_key,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_SplHeap_next,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_SplHeap_valid,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"recoverFromCorruption",
		zim_spl_SplHeap_recoverFromCorruption,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isCorrupted",
		zim_spl_SplHeap_isCorrupted,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_SplPriorityQueue___debugInfo,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"extract",
		zim_spl_SplHeap_extract,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"insert",
		zim_spl_SplHeap_insert,
		ArginfoHeapInsert,
		uint32_t(b.SizeOf("arginfo_heap_insert")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"top",
		zim_spl_SplHeap_top,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_SplHeap_count,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isEmpty",
		zim_spl_SplHeap_isEmpty,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_SplHeap_rewind,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_SplHeap_current,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_SplHeap_key,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_SplHeap_next,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_SplHeap_valid,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"recoverFromCorruption",
		zim_spl_SplHeap_recoverFromCorruption,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isCorrupted",
		zim_spl_SplHeap_isCorrupted,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_SplHeap___debugInfo,
		ArginfoSplheapVoid,
		uint32_t(b.SizeOf("arginfo_splheap_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"compare",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PROTECTED | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

/* }}} */

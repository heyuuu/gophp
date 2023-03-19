// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
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

var spl_ce_SplHeap *types.ClassEntry
var spl_ce_SplMinHeap *types.ClassEntry
var spl_ce_SplMaxHeap *types.ClassEntry
var spl_ce_SplPriorityQueue *types.ClassEntry

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
type SplPtrHeapCmpFunc func(any, any, *types.Zval) int

/* define an __special__  overloaded iterator structure */

/* {{{ proto void SplPriorityQueue::__debugInfo() */

/* iterator handler table */

var SplHeapItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplHeapItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var SplPqueueItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplPqueueItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var ArginfoHeapInsert []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var ArginfoHeapCompare []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value1"),
	zend.MakeArgInfo("value2"),
}
var ArginfoPqueueInsert []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
	zend.MakeArgInfo("priority"),
}
var ArginfoPqueueSetflags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("flags"),
}
var ArginfoSplheapVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var spl_funcs_SplMinHeap []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("compare", zend.ZEND_ACC_PROTECTED, zim_spl_SplMinHeap_compare, ArginfoHeapCompare),
}
var spl_funcs_SplMaxHeap []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("compare", zend.ZEND_ACC_PROTECTED, zim_spl_SplMaxHeap_compare, ArginfoHeapCompare),
}
var spl_funcs_SplPriorityQueue []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("compare", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_compare, ArginfoHeapCompare),
	types.MakeZendFunctionEntryEx("insert", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_insert, ArginfoPqueueInsert),
	types.MakeZendFunctionEntryEx("setExtractFlags", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_setExtractFlags, ArginfoPqueueSetflags),
	types.MakeZendFunctionEntryEx("getExtractFlags", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_getExtractFlags, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("top", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_top, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("extract", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_extract, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_count, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("isEmpty", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_isEmpty, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_rewind, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue_current, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_key, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_next, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_valid, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("recoverFromCorruption", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_recoverFromCorruption, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("isCorrupted", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_isCorrupted, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplPriorityQueue___debugInfo, ArginfoSplheapVoid),
}
var spl_funcs_SplHeap []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("extract", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_extract, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("insert", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_insert, ArginfoHeapInsert),
	types.MakeZendFunctionEntryEx("top", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_top, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_count, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("isEmpty", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_isEmpty, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_rewind, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_current, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_key, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_next, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_valid, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("recoverFromCorruption", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_recoverFromCorruption, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("isCorrupted", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap_isCorrupted, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplHeap___debugInfo, ArginfoSplheapVoid),
	types.MakeZendFunctionEntryEx("compare", zend.ZEND_ACC_PROTECTED|zend.ZEND_ACC_ABSTRACT, nil, nil),
}

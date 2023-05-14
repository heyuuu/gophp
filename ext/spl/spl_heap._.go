package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplHeap *types.ClassEntry
var spl_ce_SplMinHeap *types.ClassEntry
var spl_ce_SplMaxHeap *types.ClassEntry
var spl_ce_SplPriorityQueue *types.ClassEntry

const PTR_HEAP_BLOCK_SIZE = 64
const SPL_HEAP_CORRUPTED = 0x1
const SPL_PQUEUE_EXTR_MASK = 0x3
const SPL_PQUEUE_EXTR_BOTH = 0x3
const SPL_PQUEUE_EXTR_DATA = 0x1
const SPL_PQUEUE_EXTR_PRIORITY = 0x2

var spl_handler_SplHeap types.ObjectHandlers
var spl_handler_SplPriorityQueue types.ObjectHandlers

type SplPtrHeapDtorFunc func(any)
type SplPtrHeapCtorFunc func(any)
type SplPtrHeapCmpFunc func(any, any, *types.Zval) int

var SplHeapItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplHeapItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var SplPqueueItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplPqueueItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var spl_funcs_SplMinHeap []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("compare", types.AccProtected, zim_spl_SplMinHeap_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
}
var spl_funcs_SplMaxHeap []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("compare", types.AccProtected, zim_spl_SplMaxHeap_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
}
var spl_funcs_SplPriorityQueue []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("compare", types.AccPublic, zim_spl_SplPriorityQueue_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
	types.MakeZendFunctionEntryEx("insert", types.AccPublic, zim_spl_SplPriorityQueue_insert, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
		zend.MakeArgName("priority"),
	}),
	types.MakeZendFunctionEntryEx("setExtractFlags", types.AccPublic, zim_spl_SplPriorityQueue_setExtractFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("getExtractFlags", types.AccPublic, zim_spl_SplPriorityQueue_getExtractFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("top", types.AccPublic, zim_spl_SplPriorityQueue_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("extract", types.AccPublic, zim_spl_SplPriorityQueue_extract, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", types.AccPublic, zim_spl_SplHeap_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isEmpty", types.AccPublic, zim_spl_SplHeap_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_SplHeap_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_SplPriorityQueue_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_SplHeap_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_SplHeap_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic, zim_spl_SplHeap_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("recoverFromCorruption", types.AccPublic, zim_spl_SplHeap_recoverFromCorruption, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isCorrupted", types.AccPublic, zim_spl_SplHeap_isCorrupted, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", types.AccPublic, zim_spl_SplPriorityQueue___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_SplHeap []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("extract", types.AccPublic, zim_spl_SplHeap_extract, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("insert", types.AccPublic, zim_spl_SplHeap_insert, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("top", types.AccPublic, zim_spl_SplHeap_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", types.AccPublic, zim_spl_SplHeap_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isEmpty", types.AccPublic, zim_spl_SplHeap_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_SplHeap_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_SplHeap_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_SplHeap_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_SplHeap_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic, zim_spl_SplHeap_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("recoverFromCorruption", types.AccPublic, zim_spl_SplHeap_recoverFromCorruption, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isCorrupted", types.AccPublic, zim_spl_SplHeap_isCorrupted, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", types.AccPublic, zim_spl_SplHeap___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("compare", types.AccProtected|types.AccAbstract, nil, nil),
}

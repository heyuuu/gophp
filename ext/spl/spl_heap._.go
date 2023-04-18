package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplHeap *types2.ClassEntry
var spl_ce_SplMinHeap *types2.ClassEntry
var spl_ce_SplMaxHeap *types2.ClassEntry
var spl_ce_SplPriorityQueue *types2.ClassEntry

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
type SplPtrHeapCmpFunc func(any, any, *types2.Zval) int

var SplHeapItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplHeapItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var SplPqueueItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplHeapItDtor, SplHeapItValid, SplPqueueItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil)
var spl_funcs_SplMinHeap []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("compare", zend.AccProtected, zim_spl_SplMinHeap_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
}
var spl_funcs_SplMaxHeap []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("compare", zend.AccProtected, zim_spl_SplMaxHeap_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
}
var spl_funcs_SplPriorityQueue []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("compare", zend.AccPublic, zim_spl_SplPriorityQueue_compare, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value1"),
		zend.MakeArgName("value2"),
	}),
	types2.MakeZendFunctionEntryEx("insert", zend.AccPublic, zim_spl_SplPriorityQueue_insert, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
		zend.MakeArgName("priority"),
	}),
	types2.MakeZendFunctionEntryEx("setExtractFlags", zend.AccPublic, zim_spl_SplPriorityQueue_setExtractFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("getExtractFlags", zend.AccPublic, zim_spl_SplPriorityQueue_getExtractFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("top", zend.AccPublic, zim_spl_SplPriorityQueue_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("extract", zend.AccPublic, zim_spl_SplPriorityQueue_extract, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplHeap_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isEmpty", zend.AccPublic, zim_spl_SplHeap_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplHeap_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplPriorityQueue_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplHeap_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplHeap_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplHeap_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("recoverFromCorruption", zend.AccPublic, zim_spl_SplHeap_recoverFromCorruption, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isCorrupted", zend.AccPublic, zim_spl_SplHeap_isCorrupted, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_SplPriorityQueue___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_SplHeap []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("extract", zend.AccPublic, zim_spl_SplHeap_extract, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("insert", zend.AccPublic, zim_spl_SplHeap_insert, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("top", zend.AccPublic, zim_spl_SplHeap_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplHeap_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isEmpty", zend.AccPublic, zim_spl_SplHeap_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplHeap_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplHeap_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplHeap_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplHeap_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplHeap_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("recoverFromCorruption", zend.AccPublic, zim_spl_SplHeap_recoverFromCorruption, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isCorrupted", zend.AccPublic, zim_spl_SplHeap_isCorrupted, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_SplHeap___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("compare", zend.AccProtected|zend.AccAbstract, nil, nil),
}

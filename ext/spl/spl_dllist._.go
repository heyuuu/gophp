package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplDoublyLinkedList *types.ClassEntry
var spl_ce_SplQueue *types.ClassEntry
var spl_ce_SplStack *types.ClassEntry

var spl_handler_SplDoublyLinkedList zend.ZendObjectHandlers

const SPL_DLLIST_IT_DELETE = 0x1
const SPL_DLLIST_IT_LIFO = 0x2
const SPL_DLLIST_IT_MASK = 0x3
const SPL_DLLIST_IT_FIX = 0x4

type SplPtrLlistDtorFunc func(*SplPtrLlistElement)
type SplPtrLlistCtorFunc func(*SplPtrLlistElement)

var SplDllistItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplDllistItDtor, SplDllistItValid, SplDllistItGetCurrentData, SplDllistItGetCurrentKey, SplDllistItMoveForward, SplDllistItRewind, nil)
var spl_funcs_SplQueue []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("enqueue", zend.AccPublic, zim_spl_SplDoublyLinkedList_push, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("dequeue", zend.AccPublic, zim_spl_SplDoublyLinkedList_shift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_SplDoublyLinkedList []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("pop", zend.AccPublic, zim_spl_SplDoublyLinkedList_pop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("shift", zend.AccPublic, zim_spl_SplDoublyLinkedList_shift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("push", zend.AccPublic, zim_spl_SplDoublyLinkedList_push, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("unshift", zend.AccPublic, zim_spl_SplDoublyLinkedList_unshift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("top", zend.AccPublic, zim_spl_SplDoublyLinkedList_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("bottom", zend.AccPublic, zim_spl_SplDoublyLinkedList_bottom, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isEmpty", zend.AccPublic, zim_spl_SplDoublyLinkedList_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setIteratorMode", zend.AccPublic, zim_spl_SplDoublyLinkedList_setIteratorMode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("mode"),
	}),
	types.MakeZendFunctionEntryEx("getIteratorMode", zend.AccPublic, zim_spl_SplDoublyLinkedList_getIteratorMode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_SplDoublyLinkedList___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplDoublyLinkedList_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("add", zend.AccPublic, zim_spl_SplDoublyLinkedList_add, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplDoublyLinkedList_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplDoublyLinkedList_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplDoublyLinkedList_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplDoublyLinkedList_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("prev", zend.AccPublic, zim_spl_SplDoublyLinkedList_prev, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplDoublyLinkedList_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("unserialize", zend.AccPublic, zim_spl_SplDoublyLinkedList_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("serialize", zend.AccPublic, zim_spl_SplDoublyLinkedList_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__unserialize", zend.AccPublic, zim_spl_SplDoublyLinkedList___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("__serialize", zend.AccPublic, zim_spl_SplDoublyLinkedList___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

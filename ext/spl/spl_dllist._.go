package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplDoublyLinkedList *types2.ClassEntry
var spl_ce_SplQueue *types2.ClassEntry
var spl_ce_SplStack *types2.ClassEntry

var spl_handler_SplDoublyLinkedList zend.ZendObjectHandlers

const SPL_DLLIST_IT_DELETE = 0x1
const SPL_DLLIST_IT_LIFO = 0x2
const SPL_DLLIST_IT_MASK = 0x3
const SPL_DLLIST_IT_FIX = 0x4

type SplPtrLlistDtorFunc func(*SplPtrLlistElement)
type SplPtrLlistCtorFunc func(*SplPtrLlistElement)

var SplDllistItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplDllistItDtor, SplDllistItValid, SplDllistItGetCurrentData, SplDllistItGetCurrentKey, SplDllistItMoveForward, SplDllistItRewind, nil)
var spl_funcs_SplQueue []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("enqueue", zend.AccPublic, zim_spl_SplDoublyLinkedList_push, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("dequeue", zend.AccPublic, zim_spl_SplDoublyLinkedList_shift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_SplDoublyLinkedList []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("pop", zend.AccPublic, zim_spl_SplDoublyLinkedList_pop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("shift", zend.AccPublic, zim_spl_SplDoublyLinkedList_shift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("push", zend.AccPublic, zim_spl_SplDoublyLinkedList_push, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("unshift", zend.AccPublic, zim_spl_SplDoublyLinkedList_unshift, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("top", zend.AccPublic, zim_spl_SplDoublyLinkedList_top, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("bottom", zend.AccPublic, zim_spl_SplDoublyLinkedList_bottom, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isEmpty", zend.AccPublic, zim_spl_SplDoublyLinkedList_isEmpty, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setIteratorMode", zend.AccPublic, zim_spl_SplDoublyLinkedList_setIteratorMode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("mode"),
	}),
	types2.MakeZendFunctionEntryEx("getIteratorMode", zend.AccPublic, zim_spl_SplDoublyLinkedList_getIteratorMode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_SplDoublyLinkedList___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplDoublyLinkedList_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types2.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_SplDoublyLinkedList_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("add", zend.AccPublic, zim_spl_SplDoublyLinkedList_add, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplDoublyLinkedList_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplDoublyLinkedList_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplDoublyLinkedList_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplDoublyLinkedList_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("prev", zend.AccPublic, zim_spl_SplDoublyLinkedList_prev, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplDoublyLinkedList_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("unserialize", zend.AccPublic, zim_spl_SplDoublyLinkedList_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types2.MakeZendFunctionEntryEx("serialize", zend.AccPublic, zim_spl_SplDoublyLinkedList_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__unserialize", zend.AccPublic, zim_spl_SplDoublyLinkedList___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types2.MakeZendFunctionEntryEx("__serialize", zend.AccPublic, zim_spl_SplDoublyLinkedList___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

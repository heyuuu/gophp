// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

var spl_ce_SplDoublyLinkedList *zend.ZendClassEntry
var spl_ce_SplQueue *zend.ZendClassEntry
var spl_ce_SplStack *zend.ZendClassEntry
var spl_handler_SplDoublyLinkedList zend.ZendObjectHandlers

const SPL_DLLIST_IT_DELETE = 0x1
const SPL_DLLIST_IT_LIFO = 0x2
const SPL_DLLIST_IT_MASK = 0x3
const SPL_DLLIST_IT_FIX = 0x4

type SplPtrLlistDtorFunc func(*SplPtrLlistElement)
type SplPtrLlistCtorFunc func(*SplPtrLlistElement)

var SplDllistItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplDllistItDtor, SplDllistItValid, SplDllistItGetCurrentData, SplDllistItGetCurrentKey, SplDllistItMoveForward, SplDllistItRewind, nil}
var ArginfoDllistSetiteratormode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"mode", 0, 0, 0},
}
var ArginfoDllistPush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var arginfo_dllist_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_dllist_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var ArginfoDllistVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoDllistSerialized []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"serialized", 0, 0, 0},
}
var spl_funcs_SplQueue []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"enqueue",
		zim_spl_SplDoublyLinkedList_push,
		ArginfoDllistPush,
		uint32_t(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"dequeue",
		zim_spl_SplDoublyLinkedList_shift,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplDoublyLinkedList []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"pop",
		zim_spl_SplDoublyLinkedList_pop,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"shift",
		zim_spl_SplDoublyLinkedList_shift,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"push",
		zim_spl_SplDoublyLinkedList_push,
		ArginfoDllistPush,
		uint32_t(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"unshift",
		zim_spl_SplDoublyLinkedList_unshift,
		ArginfoDllistPush,
		uint32_t(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"top",
		zim_spl_SplDoublyLinkedList_top,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"bottom",
		zim_spl_SplDoublyLinkedList_bottom,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isEmpty",
		zim_spl_SplDoublyLinkedList_isEmpty,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setIteratorMode",
		zim_spl_SplDoublyLinkedList_setIteratorMode,
		ArginfoDllistSetiteratormode,
		uint32_t(b.SizeOf("arginfo_dllist_setiteratormode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getIteratorMode",
		zim_spl_SplDoublyLinkedList_getIteratorMode,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_SplDoublyLinkedList___debugInfo,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_SplDoublyLinkedList_count,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetExists",
		zim_spl_SplDoublyLinkedList_offsetExists,
		arginfo_dllist_offsetGet,
		uint32_t(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetGet",
		zim_spl_SplDoublyLinkedList_offsetGet,
		arginfo_dllist_offsetGet,
		uint32_t(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetSet",
		zim_spl_SplDoublyLinkedList_offsetSet,
		arginfo_dllist_offsetSet,
		uint32_t(b.SizeOf("arginfo_dllist_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetUnset",
		zim_spl_SplDoublyLinkedList_offsetUnset,
		arginfo_dllist_offsetGet,
		uint32_t(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"add",
		zim_spl_SplDoublyLinkedList_add,
		arginfo_dllist_offsetSet,
		uint32_t(b.SizeOf("arginfo_dllist_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_SplDoublyLinkedList_rewind,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_SplDoublyLinkedList_current,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_SplDoublyLinkedList_key,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_SplDoublyLinkedList_next,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"prev",
		zim_spl_SplDoublyLinkedList_prev,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_SplDoublyLinkedList_valid,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"unserialize",
		zim_spl_SplDoublyLinkedList_unserialize,
		ArginfoDllistSerialized,
		uint32_t(b.SizeOf("arginfo_dllist_serialized")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"serialize",
		zim_spl_SplDoublyLinkedList_serialize,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__unserialize",
		zim_spl_SplDoublyLinkedList___unserialize,
		ArginfoDllistSerialized,
		uint32_t(b.SizeOf("arginfo_dllist_serialized")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__serialize",
		zim_spl_SplDoublyLinkedList___serialize,
		ArginfoDllistVoid,
		uint32_t(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

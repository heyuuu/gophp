// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/spl/spl_dllist.h>

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

var spl_ce_SplDoublyLinkedList *zend.ZendClassEntry
var spl_ce_SplQueue *zend.ZendClassEntry
var spl_ce_SplStack *zend.ZendClassEntry

// Source: <ext/spl/spl_dllist.c>

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

var spl_handler_SplDoublyLinkedList zend.ZendObjectHandlers

const SPL_DLLIST_IT_DELETE = 0x1
const SPL_DLLIST_IT_LIFO = 0x2
const SPL_DLLIST_IT_MASK = 0x3
const SPL_DLLIST_IT_FIX = 0x4

type SplPtrLlistDtorFunc func(*SplPtrLlistElement)
type SplPtrLlistCtorFunc func(*SplPtrLlistElement)

/* define an __special__  overloaded iterator structure */

/* {{{  spl_ptr_llist */

/* }}}} */

/* {{{ proto mixed SplDoublyLinkedList::offsetGet(mixed index)
Returns the value at the specified $index. */

/* {{{ proto void SplDoublyLinkedList::offsetSet(mixed index, mixed newval)
Sets the value at the specified $index to $newval. */

/* {{{ proto void SplDoublyLinkedList::offsetUnset(mixed index)
Unsets the value at the specified $index. */

/* {{{ proto void SplDoublyLinkedList::unserialize(string serialized)
Unserializes storage */

/* {{{ proto array SplDoublyLinkedList::__serialize() */

/* {{{ proto void SplDoublyLinkedList::__unserialize(array serialized) */

/* {{{ proto void SplDoublyLinkedList::add(mixed index, mixed newval)
Inserts a new entry before the specified $index consisting of $newval. */

/* {{{ proto void SplDoublyLinkedList::__debugInfo() */

/* {{{ iterator handler table */

var SplDllistItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplDllistItDtor, SplDllistItValid, SplDllistItGetCurrentData, SplDllistItGetCurrentKey, SplDllistItMoveForward, SplDllistItRewind, nil)
var ArginfoDllistSetiteratormode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("mode"),
}
var ArginfoDllistPush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var arginfo_dllist_offsetGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("index"),
}
var arginfo_dllist_offsetSet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("index"),
	zend.MakeArgInfo("newval"),
}
var ArginfoDllistVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoDllistSerialized []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("serialized"),
}
var spl_funcs_SplQueue []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("enqueue", zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush, uint32(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("dequeue", zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var spl_funcs_SplDoublyLinkedList []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("pop", zim_spl_SplDoublyLinkedList_pop, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("shift", zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("push", zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush, uint32(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("unshift", zim_spl_SplDoublyLinkedList_unshift, ArginfoDllistPush, uint32(b.SizeOf("arginfo_dllist_push")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("top", zim_spl_SplDoublyLinkedList_top, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("bottom", zim_spl_SplDoublyLinkedList_bottom, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isEmpty", zim_spl_SplDoublyLinkedList_isEmpty, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setIteratorMode", zim_spl_SplDoublyLinkedList_setIteratorMode, ArginfoDllistSetiteratormode, uint32(b.SizeOf("arginfo_dllist_setiteratormode")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getIteratorMode", zim_spl_SplDoublyLinkedList_getIteratorMode, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__debugInfo", zim_spl_SplDoublyLinkedList___debugInfo, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("count", zim_spl_SplDoublyLinkedList_count, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetExists", zim_spl_SplDoublyLinkedList_offsetExists, arginfo_dllist_offsetGet, uint32(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetGet", zim_spl_SplDoublyLinkedList_offsetGet, arginfo_dllist_offsetGet, uint32(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetSet", zim_spl_SplDoublyLinkedList_offsetSet, arginfo_dllist_offsetSet, uint32(b.SizeOf("arginfo_dllist_offsetSet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("offsetUnset", zim_spl_SplDoublyLinkedList_offsetUnset, arginfo_dllist_offsetGet, uint32(b.SizeOf("arginfo_dllist_offsetGet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("add", zim_spl_SplDoublyLinkedList_add, arginfo_dllist_offsetSet, uint32(b.SizeOf("arginfo_dllist_offsetSet")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("rewind", zim_spl_SplDoublyLinkedList_rewind, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("current", zim_spl_SplDoublyLinkedList_current, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("key", zim_spl_SplDoublyLinkedList_key, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("next", zim_spl_SplDoublyLinkedList_next, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("prev", zim_spl_SplDoublyLinkedList_prev, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("valid", zim_spl_SplDoublyLinkedList_valid, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("unserialize", zim_spl_SplDoublyLinkedList_unserialize, ArginfoDllistSerialized, uint32(b.SizeOf("arginfo_dllist_serialized")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("serialize", zim_spl_SplDoublyLinkedList_serialize, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__unserialize", zim_spl_SplDoublyLinkedList___unserialize, ArginfoDllistSerialized, uint32(b.SizeOf("arginfo_dllist_serialized")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__serialize", zim_spl_SplDoublyLinkedList___serialize, ArginfoDllistVoid, uint32(b.SizeOf("arginfo_dllist_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

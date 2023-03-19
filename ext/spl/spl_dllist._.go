// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
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

var spl_ce_SplDoublyLinkedList *types.ClassEntry
var spl_ce_SplQueue *types.ClassEntry
var spl_ce_SplStack *types.ClassEntry

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
var spl_funcs_SplQueue []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("enqueue", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush),
	types.MakeZendFunctionEntryEx("dequeue", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid),
}
var spl_funcs_SplDoublyLinkedList []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("pop", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_pop, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("shift", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("push", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush),
	types.MakeZendFunctionEntryEx("unshift", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_unshift, ArginfoDllistPush),
	types.MakeZendFunctionEntryEx("top", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_top, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("bottom", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_bottom, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("isEmpty", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_isEmpty, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("setIteratorMode", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_setIteratorMode, ArginfoDllistSetiteratormode),
	types.MakeZendFunctionEntryEx("getIteratorMode", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_getIteratorMode, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___debugInfo, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_count, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetExists, arginfo_dllist_offsetGet),
	types.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetGet, arginfo_dllist_offsetGet),
	types.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetSet, arginfo_dllist_offsetSet),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetUnset, arginfo_dllist_offsetGet),
	types.MakeZendFunctionEntryEx("add", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_add, arginfo_dllist_offsetSet),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_rewind, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_current, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_key, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_next, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("prev", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_prev, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_valid, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_unserialize, ArginfoDllistSerialized),
	types.MakeZendFunctionEntryEx("serialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_serialize, ArginfoDllistVoid),
	types.MakeZendFunctionEntryEx("__unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___unserialize, ArginfoDllistSerialized),
	types.MakeZendFunctionEntryEx("__serialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___serialize, ArginfoDllistVoid),
}

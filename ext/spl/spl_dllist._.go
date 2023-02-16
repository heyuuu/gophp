// <<generate>>

package spl

import (
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
	zend.MakeZendFunctionEntryEx("enqueue", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush),
	zend.MakeZendFunctionEntryEx("dequeue", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid),
}
var spl_funcs_SplDoublyLinkedList []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("pop", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_pop, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("shift", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_shift, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("push", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_push, ArginfoDllistPush),
	zend.MakeZendFunctionEntryEx("unshift", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_unshift, ArginfoDllistPush),
	zend.MakeZendFunctionEntryEx("top", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_top, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("bottom", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_bottom, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("isEmpty", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_isEmpty, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("setIteratorMode", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_setIteratorMode, ArginfoDllistSetiteratormode),
	zend.MakeZendFunctionEntryEx("getIteratorMode", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_getIteratorMode, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___debugInfo, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_count, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetExists, arginfo_dllist_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetGet, arginfo_dllist_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetSet, arginfo_dllist_offsetSet),
	zend.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_offsetUnset, arginfo_dllist_offsetGet),
	zend.MakeZendFunctionEntryEx("add", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_add, arginfo_dllist_offsetSet),
	zend.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_rewind, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_current, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_key, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_next, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("prev", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_prev, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_valid, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_unserialize, ArginfoDllistSerialized),
	zend.MakeZendFunctionEntryEx("serialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList_serialize, ArginfoDllistVoid),
	zend.MakeZendFunctionEntryEx("__unserialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___unserialize, ArginfoDllistSerialized),
	zend.MakeZendFunctionEntryEx("__serialize", zend.ZEND_ACC_PUBLIC, zim_spl_SplDoublyLinkedList___serialize, ArginfoDllistVoid),
}

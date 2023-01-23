// <<generate>>

package spl

import (
	"sik/core"
	"sik/ext/standard"
	g "sik/runtime/grammar"
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

// #define SPL_DLLIST_H

// # include "php.h"

// # include "php_spl.h"

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

// # include "php.h"

// # include "zend_exceptions.h"

// # include "zend_hash.h"

// # include "php_spl.h"

// # include "ext/standard/info.h"

// # include "ext/standard/php_var.h"

// # include "zend_smart_str.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_dllist.h"

// # include "spl_exceptions.h"

var spl_handler_SplDoublyLinkedList zend.ZendObjectHandlers

// #define SPL_LLIST_DELREF(elem) if ( ! -- ( elem ) -> rc ) { efree ( elem ) ; }

// #define SPL_LLIST_CHECK_DELREF(elem) if ( ( elem ) && ! -- ( elem ) -> rc ) { efree ( elem ) ; }

// #define SPL_LLIST_ADDREF(elem) ( elem ) -> rc ++

// #define SPL_LLIST_CHECK_ADDREF(elem) if ( elem ) ( elem ) -> rc ++

// #define SPL_DLLIST_IT_DELETE       0x00000001

// #define SPL_DLLIST_IT_LIFO       0x00000002

// #define SPL_DLLIST_IT_MASK       0x00000003

// #define SPL_DLLIST_IT_FIX       0x00000004

type SplPtrLlistDtorFunc func(*SplPtrLlistElement)
type SplPtrLlistCtorFunc func(*SplPtrLlistElement)

/* define an __special__  overloaded iterator structure */

func SplDllistFromObj(obj *zend.ZendObject) *SplDllistObject {
	return (*SplDllistObject)((*byte)(obj - zend_long((*byte)(&((*SplDllistObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLDLLIST_P(zv) spl_dllist_from_obj ( Z_OBJ_P ( ( zv ) ) )

/* {{{  spl_ptr_llist */

func SplPtrLlistZvalDtor(elem *SplPtrLlistElement) {
	if elem.data.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&elem.data)
		&elem.data.u1.type_info = 0
	}
}

/* }}} */

func SplPtrLlistZvalCtor(elem *SplPtrLlistElement) {
	if elem.data.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&(elem.GetData()))
	}
}

/* }}} */

func SplPtrLlistInit(ctor SplPtrLlistCtorFunc, dtor SplPtrLlistDtorFunc) *SplPtrLlist {
	var llist *SplPtrLlist = zend._emalloc(g.SizeOf("spl_ptr_llist"))
	llist.SetHead(nil)
	llist.SetTail(nil)
	llist.SetCount(0)
	llist.SetDtor(dtor)
	llist.SetCtor(ctor)
	return llist
}

/* }}} */

func SplPtrLlistCount(llist *SplPtrLlist) zend.ZendLong { return zend.ZendLong(llist.GetCount()) }

/* }}} */

func SplPtrLlistDestroy(llist *SplPtrLlist) {
	var current *SplPtrLlistElement = llist.GetHead()
	var next *SplPtrLlistElement
	var dtor SplPtrLlistDtorFunc = llist.GetDtor()
	for current != nil {
		next = current.GetNext()
		if dtor != nil {
			dtor(current)
		}
		if !(g.PreDec(&(current.GetRc()))) {
			zend._efree(current)
		}
		current = next
	}
	zend._efree(llist)
}

/* }}} */

func SplPtrLlistOffset(llist *SplPtrLlist, offset zend.ZendLong, backward int) *SplPtrLlistElement {
	var current *SplPtrLlistElement
	var pos int = 0
	if backward != 0 {
		current = llist.GetTail()
	} else {
		current = llist.GetHead()
	}
	for current != nil && pos < offset {
		pos++
		if backward != 0 {
			current = current.GetPrev()
		} else {
			current = current.GetNext()
		}
	}
	return current
}

/* }}} */

func SplPtrLlistUnshift(llist *SplPtrLlist, data *zend.Zval) {
	var elem *SplPtrLlistElement = zend._emalloc(g.SizeOf("spl_ptr_llist_element"))
	elem.SetRc(1)
	elem.SetPrev(nil)
	elem.SetNext(llist.GetHead())
	var _z1 *zend.Zval = &elem.data
	var _z2 *zend.Zval = data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if llist.GetHead() != nil {
		llist.GetHead().SetPrev(elem)
	} else {
		llist.SetTail(elem)
	}
	llist.SetHead(elem)
	llist.GetCount()++
	if llist.GetCtor() != nil {
		llist.GetCtor()(elem)
	}
}

/* }}} */

func SplPtrLlistPush(llist *SplPtrLlist, data *zend.Zval) {
	var elem *SplPtrLlistElement = zend._emalloc(g.SizeOf("spl_ptr_llist_element"))
	elem.SetRc(1)
	elem.SetPrev(llist.GetTail())
	elem.SetNext(nil)
	var _z1 *zend.Zval = &elem.data
	var _z2 *zend.Zval = data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if llist.GetTail() != nil {
		llist.GetTail().SetNext(elem)
	} else {
		llist.SetHead(elem)
	}
	llist.SetTail(elem)
	llist.GetCount()++
	if llist.GetCtor() != nil {
		llist.GetCtor()(elem)
	}
}

/* }}} */

func SplPtrLlistPop(llist *SplPtrLlist, ret *zend.Zval) {
	var tail *SplPtrLlistElement = llist.GetTail()
	if tail == nil {
		ret.u1.type_info = 0
		return
	}
	if tail.GetPrev() != nil {
		tail.GetPrev().SetNext(nil)
	} else {
		llist.SetHead(nil)
	}
	llist.SetTail(tail.GetPrev())
	llist.GetCount()--
	var _z1 *zend.Zval = ret
	var _z2 *zend.Zval = &tail.data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	tail.SetPrev(nil)
	if llist.GetDtor() != nil {
		llist.GetDtor()(tail)
	}
	&tail.data.u1.type_info = 0
	if !(g.PreDec(&(tail.GetRc()))) {
		zend._efree(tail)
	}
}

/* }}} */

func SplPtrLlistLast(llist *SplPtrLlist) *zend.Zval {
	var tail *SplPtrLlistElement = llist.GetTail()
	if tail == nil {
		return nil
	} else {
		return &tail.data
	}
}

/* }}} */

func SplPtrLlistFirst(llist *SplPtrLlist) *zend.Zval {
	var head *SplPtrLlistElement = llist.GetHead()
	if head == nil {
		return nil
	} else {
		return &head.data
	}
}

/* }}} */

func SplPtrLlistShift(llist *SplPtrLlist, ret *zend.Zval) {
	var head *SplPtrLlistElement = llist.GetHead()
	if head == nil {
		ret.u1.type_info = 0
		return
	}
	if head.GetNext() != nil {
		head.GetNext().SetPrev(nil)
	} else {
		llist.SetTail(nil)
	}
	llist.SetHead(head.GetNext())
	llist.GetCount()--
	var _z1 *zend.Zval = ret
	var _z2 *zend.Zval = &head.data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	head.SetNext(nil)
	if llist.GetDtor() != nil {
		llist.GetDtor()(head)
	}
	&head.data.u1.type_info = 0
	if !(g.PreDec(&(head.GetRc()))) {
		zend._efree(head)
	}
}

/* }}} */

func SplPtrLlistCopy(from *SplPtrLlist, to *SplPtrLlist) {
	var current *SplPtrLlistElement = from.GetHead()
	var next *SplPtrLlistElement

	//???    spl_ptr_llist_ctor_func ctor = from->ctor;

	for current != nil {
		next = current.GetNext()

		/*??? FIXME
		  if (ctor) {
		      ctor(current);
		  }
		*/

		SplPtrLlistPush(to, &current.data)
		current = next
	}

	//???    spl_ptr_llist_ctor_func ctor = from->ctor;
}

/* }}} */

func SplDllistObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplDllistObject = SplDllistFromObj(object)
	var tmp zend.Zval
	zend.ZendObjectStdDtor(&intern.std)
	for intern.GetLlist().GetCount() > 0 {
		SplPtrLlistPop(intern.GetLlist(), &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	if intern.GetGcData() != nil {
		zend._efree(intern.GetGcData())
	}
	SplPtrLlistDestroy(intern.GetLlist())
	if intern.GetTraversePointer() != nil && !(g.PreDec(&(intern.GetTraversePointer().GetRc()))) {
		zend._efree(intern.GetTraversePointer())
	}
}

/* }}} */

func SplDllistObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplDllistObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_dllist_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.SetFlags(0)
	intern.SetTraversePosition(0)
	if orig != nil {
		var other *SplDllistObject = SplDllistFromObj(orig.value.obj)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			intern.SetLlist((*SplPtrLlist)(SplPtrLlistInit(other.GetLlist().GetCtor(), other.GetLlist().GetDtor())))
			SplPtrLlistCopy(other.GetLlist(), intern.GetLlist())
			intern.SetTraversePointer(intern.GetLlist().GetHead())
			if intern.GetTraversePointer() != nil {
				intern.GetTraversePointer().GetRc()++
			}
		} else {
			intern.SetLlist(other.GetLlist())
			intern.SetTraversePointer(intern.GetLlist().GetHead())
			if intern.GetTraversePointer() != nil {
				intern.GetTraversePointer().GetRc()++
			}
		}
		intern.SetFlags(other.GetFlags())
	} else {
		intern.SetLlist((*SplPtrLlist)(SplPtrLlistInit(SplPtrLlistZvalCtor, SplPtrLlistZvalDtor)))
		intern.SetTraversePointer(intern.GetLlist().GetHead())
		if intern.GetTraversePointer() != nil {
			intern.GetTraversePointer().GetRc()++
		}
	}
	for parent != nil {
		if parent == spl_ce_SplStack {
			intern.SetFlags(intern.GetFlags() | 0x4 | 0x2)
			intern.std.handlers = &spl_handler_SplDoublyLinkedList
		} else if parent == spl_ce_SplQueue {
			intern.SetFlags(intern.GetFlags() | 0x4)
			intern.std.handlers = &spl_handler_SplDoublyLinkedList
		}
		if parent == spl_ce_SplDoublyLinkedList {
			intern.std.handlers = &spl_handler_SplDoublyLinkedList
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, 1<<6, "Internal compiler error, Class is not child of SplDoublyLinkedList")
	}
	if inherited != 0 {
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetget", g.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().common.scope == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetset", g.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().common.scope == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetexists", g.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().common.scope == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(&class_type.function_table, "offsetunset", g.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().common.scope == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.function_table, "count", g.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().common.scope == parent {
			intern.SetFptrCount(nil)
		}
	}
	return &intern.std
}

/* }}} */

func SplDllistObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplDllistObjectNewEx(class_type, nil, 0)
}

/* }}} */

func SplDllistObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.value.obj
	new_object = SplDllistObjectNewEx(old_object.ce, zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}

/* }}} */

func SplDllistObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplDllistObject = SplDllistFromObj(object.value.obj)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_count, "count", g.SizeOf("\"count\"")-1, &rv, 0, nil, nil)
		if rv.u1.v.type_ != 0 {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return zend.SUCCESS
		}
		*count = 0
		return zend.FAILURE
	}
	*count = SplPtrLlistCount(intern.GetLlist())
	return zend.SUCCESS
}

/* }}} */

func SplDllistObjectGetDebugInfo(obj *zend.Zval) *zend.HashTable {
	var intern *SplDllistObject = SplDllistFromObj(obj.value.obj)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var next *SplPtrLlistElement
	var tmp zend.Zval
	var dllist_array zend.Zval
	var pnstr *zend.ZendString
	var i int = 0
	var debug_info *zend.HashTable
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	debug_info = zend._zendNewArray(1)
	zend.ZendHashCopy(debug_info, intern.std.properties, zend.CopyCtorFuncT(zend.ZvalAddRef))
	pnstr = SplGenPrivatePropName(spl_ce_SplDoublyLinkedList, "flags", g.SizeOf("\"flags\"")-1)
	var __z *zend.Zval = &tmp
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	zend.ZendHashAdd(debug_info, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &dllist_array
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for current != nil {
		next = current.GetNext()
		zend.AddIndexZval(&dllist_array, i, &current.data)
		if current.data.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(current.GetData()))
		}
		i++
		current = next
	}
	pnstr = SplGenPrivatePropName(spl_ce_SplDoublyLinkedList, "dllist", g.SizeOf("\"dllist\"")-1)
	zend.ZendHashAdd(debug_info, pnstr, &dllist_array)
	zend.ZendStringReleaseEx(pnstr, 0)
	return debug_info
}

/* }}}} */

func SplDllistObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplDllistObject = SplDllistFromObj(obj.value.obj)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var i int = 0
	if intern.GetGcDataCount() < intern.GetLlist().GetCount() {
		intern.SetGcDataCount(intern.GetLlist().GetCount())
		intern.SetGcData(zend._safeErealloc(intern.GetGcData(), intern.GetGcDataCount(), g.SizeOf("zval"), 0))
	}
	for current != nil {
		var _z1 *zend.Zval = &intern.gc_data[g.PostInc(&i)]
		var _z2 *zend.Zval = &current.data
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		current = current.GetNext()
	}
	*gc_data = intern.GetGcData()
	*gc_data_count = i
	return zend.ZendStdGetProperties(obj)
}

/* }}} */

func zim_spl_SplDoublyLinkedList_push(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &value) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	SplPtrLlistPush(intern.GetLlist(), value)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_unshift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &value) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	SplPtrLlistUnshift(intern.GetLlist(), value)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_pop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	SplPtrLlistPop(intern.GetLlist(), return_value)
	if return_value.u1.v.type_ == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't pop from an empty datastructure", 0)
		return_value.u1.type_info = 1
		return
	}
}

/* }}} */

func zim_spl_SplDoublyLinkedList_shift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	SplPtrLlistShift(intern.GetLlist(), return_value)
	if return_value.u1.v.type_ == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't shift from an empty datastructure", 0)
		return_value.u1.type_info = 1
		return
	}
}

/* }}} */

func zim_spl_SplDoublyLinkedList_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	value = SplPtrLlistLast(intern.GetLlist())
	if value == nil || value.u1.v.type_ == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty datastructure", 0)
		return
	}
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func zim_spl_SplDoublyLinkedList_bottom(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	value = SplPtrLlistFirst(intern.GetLlist())
	if value == nil || value.u1.v.type_ == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty datastructure", 0)
		return
	}
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func zim_spl_SplDoublyLinkedList_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	count = SplPtrLlistCount(intern.GetLlist())
	var __z *zend.Zval = return_value
	__z.value.lval = count
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_isEmpty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplDllistObjectCountElements(&(execute_data.This), &count)
	if count == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_setIteratorMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value zend.ZendLong
	var intern *SplDllistObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &value) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	if (intern.GetFlags()&0x4) != 0 && (intern.GetFlags()&0x2) != (value&0x2) {
		zend.ZendThrowException(spl_ce_RuntimeException, "Iterators' LIFO/FIFO modes for SplStack/SplQueue objects are frozen", 0)
		return
	}
	intern.SetFlags(value&0x3 | intern.GetFlags()&0x4)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_getIteratorMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplDllistObject
	var index zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	index = SplOffsetConvertToLong(zindex)
	if index >= 0 && index < intern.GetLlist().GetCount() {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed SplDoublyLinkedList::offsetGet(mixed index)
Returns the value at the specified $index. */

func zim_spl_SplDoublyLinkedList_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var index zend.ZendLong
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	index = SplOffsetConvertToLong(zindex)
	if index < 0 || index >= intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
		return
	}
	element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&0x2)
	if element != nil {
		var value *zend.Zval = &element.data
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
	}
}

/* {{{ proto void SplDoublyLinkedList::offsetSet(mixed index, mixed newval)
Sets the value at the specified $index to $newval. */

func zim_spl_SplDoublyLinkedList_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	if zindex.u1.v.type_ == 1 {

		/* $obj[] = ... */

		SplPtrLlistPush(intern.GetLlist(), value)

		/* $obj[] = ... */

	} else {

		/* $obj[$foo] = ... */

		var index zend.ZendLong
		var element *SplPtrLlistElement
		index = SplOffsetConvertToLong(zindex)
		if index < 0 || index >= intern.GetLlist().GetCount() {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
			return
		}
		element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&0x2)
		if element != nil {

			/* call dtor on the old element as in spl_ptr_llist_pop */

			if intern.GetLlist().GetDtor() != nil {
				intern.GetLlist().GetDtor()(element)
			}

			/* the element is replaced, delref the old one as in
			 * SplDoublyLinkedList::pop() */

			zend.ZvalPtrDtor(&element.data)
			var _z1 *zend.Zval = &element.data
			var _z2 *zend.Zval = value
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t

			/* new element, call ctor as in spl_ptr_llist_push */

			if intern.GetLlist().GetCtor() != nil {
				intern.GetLlist().GetCtor()(element)
			}

			/* new element, call ctor as in spl_ptr_llist_push */

		} else {
			zend.ZvalPtrDtor(value)
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
			return
		}
	}
}

/* {{{ proto void SplDoublyLinkedList::offsetUnset(mixed index)
Unsets the value at the specified $index. */

func zim_spl_SplDoublyLinkedList_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var index zend.ZendLong
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	var llist *SplPtrLlist
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zindex) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	index = SplOffsetConvertToLong(zindex)
	llist = intern.GetLlist()
	if index < 0 || index >= intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset out of range", 0)
		return
	}
	element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&0x2)
	if element != nil {

		/* connect the neightbors */

		if element.GetPrev() != nil {
			element.GetPrev().SetNext(element.GetNext())
		}
		if element.GetNext() != nil {
			element.GetNext().SetPrev(element.GetPrev())
		}

		/* take care of head/tail */

		if element == llist.GetHead() {
			llist.SetHead(element.GetNext())
		}
		if element == llist.GetTail() {
			llist.SetTail(element.GetPrev())
		}

		/* finally, delete the element */

		llist.GetCount()--
		if llist.GetDtor() != nil {
			llist.GetDtor()(element)
		}
		if intern.GetTraversePointer() == element {
			if !(g.PreDec(&(element.GetRc()))) {
				zend._efree(element)
			}
			intern.SetTraversePointer(nil)
		}
		zend.ZvalPtrDtor(&element.data)
		&element.data.u1.type_info = 0
		if !(g.PreDec(&(element.GetRc()))) {
			zend._efree(element)
		}
	} else {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
		return
	}
}
func SplDllistItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	if iterator.GetTraversePointer() != nil && !(g.PreDec(&(iterator.GetTraversePointer().GetRc()))) {
		zend._efree(iterator.GetTraversePointer())
	}
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iterator.intern.it.data)
}

/* }}} */

func SplDllistItHelperRewind(traverse_pointer_ptr **SplPtrLlistElement, traverse_position_ptr *int, llist *SplPtrLlist, flags int) {
	if (*traverse_pointer_ptr) != nil && !(g.PreDec(&((*traverse_pointer_ptr).GetRc()))) {
		zend._efree(*traverse_pointer_ptr)
	}
	if (flags & 0x2) != 0 {
		*traverse_position_ptr = llist.GetCount() - 1
		*traverse_pointer_ptr = llist.GetTail()
	} else {
		*traverse_position_ptr = 0
		*traverse_pointer_ptr = llist.GetHead()
	}
	if (*traverse_pointer_ptr) != nil {
		(*traverse_pointer_ptr).GetRc()++
	}
}

/* }}} */

func SplDllistItHelperMoveForward(traverse_pointer_ptr **SplPtrLlistElement, traverse_position_ptr *int, llist *SplPtrLlist, flags int) {
	if (*traverse_pointer_ptr) != nil {
		var old *SplPtrLlistElement = *traverse_pointer_ptr
		if (flags & 0x2) != 0 {
			*traverse_pointer_ptr = old.GetPrev()
			*traverse_position_ptr--
			if (flags & 0x1) != 0 {
				var prev zend.Zval
				SplPtrLlistPop(llist, &prev)
				zend.ZvalPtrDtor(&prev)
			}
		} else {
			*traverse_pointer_ptr = old.GetNext()
			if (flags & 0x1) != 0 {
				var prev zend.Zval
				SplPtrLlistShift(llist, &prev)
				zend.ZvalPtrDtor(&prev)
			} else {
				*traverse_position_ptr++
			}
		}
		if !(g.PreDec(&(old.GetRc()))) {
			zend._efree(old)
		}
		if (*traverse_pointer_ptr) != nil {
			(*traverse_pointer_ptr).GetRc()++
		}
	}
}

/* }}} */

func SplDllistItRewind(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var object *SplDllistObject = SplDllistFromObj(&iter.data.value.obj)
	var llist *SplPtrLlist = object.GetLlist()
	SplDllistItHelperRewind(&iterator.traverse_pointer, &iterator.traverse_position, llist, object.GetFlags())
}

/* }}} */

func SplDllistItValid(iter *zend.ZendObjectIterator) int {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var element *SplPtrLlistElement = iterator.GetTraversePointer()
	if element != nil {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}

/* }}} */

func SplDllistItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var element *SplPtrLlistElement = iterator.GetTraversePointer()
	if element == nil || element.data.u1.v.type_ == 0 {
		return nil
	}
	return &element.data
}

/* }}} */

func SplDllistItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var __z *zend.Zval = key
	__z.value.lval = iterator.GetTraversePosition()
	__z.u1.type_info = 4
}

/* }}} */

func SplDllistItMoveForward(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var object *SplDllistObject = SplDllistFromObj(&iter.data.value.obj)
	zend.ZendUserItInvalidateCurrent(iter)
	SplDllistItHelperMoveForward(&iterator.traverse_pointer, &iterator.traverse_position, object.GetLlist(), object.GetFlags())
}

/* }}} */

func zim_spl_SplDoublyLinkedList_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetTraversePosition()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_prev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplDllistItHelperMoveForward(&intern.traverse_pointer, &intern.traverse_position, intern.GetLlist(), intern.GetFlags()^0x2)
}

/* }}} */

func zim_spl_SplDoublyLinkedList_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplDllistItHelperMoveForward(&intern.traverse_pointer, &intern.traverse_position, intern.GetLlist(), intern.GetFlags())
}

/* }}} */

func zim_spl_SplDoublyLinkedList_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetTraversePointer() != nil {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplDoublyLinkedList_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplDllistItHelperRewind(&intern.traverse_pointer, &intern.traverse_position, intern.GetLlist(), intern.GetFlags())
}

/* }}} */

func zim_spl_SplDoublyLinkedList_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	var element *SplPtrLlistElement = intern.GetTraversePointer()
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if element == nil || element.data.u1.v.type_ == 0 {
		return_value.u1.type_info = 1
		return
	} else {
		var value *zend.Zval = &element.data
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func zim_spl_SplDoublyLinkedList_serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	var buf zend.SmartStr = zend.SmartStr{0}
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var next *SplPtrLlistElement
	var flags zend.Zval
	var var_hash standard.PhpSerializeDataT
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var_hash = standard.PhpVarSerializeInit()

	/* flags */

	var __z *zend.Zval = &flags
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	standard.PhpVarSerialize(&buf, &flags, &var_hash)

	/* elements */

	for current != nil {
		zend.SmartStrAppendcEx(&buf, ':', 0)
		next = current.GetNext()
		standard.PhpVarSerialize(&buf, &current.data, &var_hash)
		current = next
	}
	zend.SmartStr0(&buf)

	/* done */

	standard.PhpVarSerializeDestroy(var_hash)
	if buf.s != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = buf.s
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto void SplDoublyLinkedList::unserialize(string serialized)
Unserializes storage */

func zim_spl_SplDoublyLinkedList_unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	var flags *zend.Zval
	var elem *zend.Zval
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}
	for intern.GetLlist().GetCount() > 0 {
		var tmp zend.Zval
		SplPtrLlistPop(intern.GetLlist(), &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	p = (*uint8)(buf)
	s = p
	var_hash = standard.PhpVarUnserializeInit()

	/* flags */

	flags = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(flags, &p, s+buf_len, &var_hash) == 0 || flags.u1.v.type_ != 4 {
		goto error
	}
	intern.SetFlags(int(flags.value.lval))

	/* elements */

	for (*p) == ':' {
		p++
		elem = standard.VarTmpVar(&var_hash)
		if standard.PhpVarUnserialize(elem, &p, s+buf_len, &var_hash) == 0 {
			goto error
		}
		standard.VarPushDtor(&var_hash, elem)
		SplPtrLlistPush(intern.GetLlist(), elem)
	}
	if (*p) != '0' {
		goto error
	}
	standard.PhpVarUnserializeDestroy(var_hash)
	return
error:
	standard.PhpVarUnserializeDestroy(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset %zd of %zd bytes", (*byte)(p-buf), buf_len)
	return
}

/* {{{ proto array SplDoublyLinkedList::__serialize() */

func zim_spl_SplDoublyLinkedList___serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var tmp zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneException()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* flags */

	var __z *zend.Zval = &tmp
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)

	/* elements */

	var __arr *zend.ZendArray = zend._zendNewArray(intern.GetLlist().GetCount())
	var __z *zend.Zval = &tmp
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for current != nil {
		zend.ZendHashNextIndexInsert(tmp.value.arr, &current.data)
		if &(current.GetData()).u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(current.GetData()))
		}
		current = current.GetNext()
	}
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)

	/* members */

	var __arr *zend.ZendArray = zend.ZendStdGetProperties(&(execute_data.This))
	var __z *zend.Zval = &tmp
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if &tmp.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&tmp)
	}
	zend.ZendHashNextIndexInsert(return_value.value.arr, &tmp)
}

/* {{{ proto void SplDoublyLinkedList::__unserialize(array serialized) */

func zim_spl_SplDoublyLinkedList___unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = SplDllistFromObj(&(execute_data.This).value.obj)
	var data *zend.HashTable
	var flags_zv *zend.Zval
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var elem *zend.Zval
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "h", &data) == zend.FAILURE {
		return
	}
	flags_zv = zend.ZendHashIndexFind(data, 0)
	storage_zv = zend.ZendHashIndexFind(data, 1)
	members_zv = zend.ZendHashIndexFind(data, 2)
	if flags_zv == nil || storage_zv == nil || members_zv == nil || flags_zv.u1.v.type_ != 4 || storage_zv.u1.v.type_ != 7 || members_zv.u1.v.type_ != 7 {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	intern.SetFlags(int(flags_zv.value.lval))
	for {
		var __ht *zend.HashTable = storage_zv.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			elem = _z
			SplPtrLlistPush(intern.GetLlist(), elem)
		}
		break
	}
	zend.ObjectPropertiesLoad(&intern.std, members_zv.value.arr)
}

/* {{{ proto void SplDoublyLinkedList::add(mixed index, mixed newval)
Inserts a new entry before the specified $index consisting of $newval. */

func zim_spl_SplDoublyLinkedList_add(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	var index zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = SplDllistFromObj(&(execute_data.This).value.obj)
	index = SplOffsetConvertToLong(zindex)
	if index < 0 || index > intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
		return
	}
	if value.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(value)
	}
	if index == intern.GetLlist().GetCount() {

		/* If index is the last entry+1 then we do a push because we're not inserting before any entry */

		SplPtrLlistPush(intern.GetLlist(), value)

		/* If index is the last entry+1 then we do a push because we're not inserting before any entry */

	} else {

		/* Create the new element we want to insert */

		var elem *SplPtrLlistElement = zend._emalloc(g.SizeOf("spl_ptr_llist_element"))

		/* Get the element we want to insert before */

		element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&0x2)
		var _z1 *zend.Zval = &elem.data
		var _z2 *zend.Zval = value
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		elem.SetRc(1)

		/* connect to the neighbours */

		elem.SetNext(element)
		elem.SetPrev(element.GetPrev())

		/* connect the neighbours to this new element */

		if elem.GetPrev() == nil {
			intern.GetLlist().SetHead(elem)
		} else {
			element.GetPrev().SetNext(elem)
		}
		element.SetPrev(elem)
		intern.GetLlist().GetCount()++
		if intern.GetLlist().GetCtor() != nil {
			intern.GetLlist().GetCtor()(elem)
		}
	}
}

/* {{{ proto void SplDoublyLinkedList::__debugInfo() */

func zim_spl_SplDoublyLinkedList___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplDllistObjectGetDebugInfo(g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* {{{ iterator handler table */

var SplDllistItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplDllistItDtor, SplDllistItValid, SplDllistItGetCurrentData, SplDllistItGetCurrentKey, SplDllistItMoveForward, SplDllistItRewind, nil}

func SplDllistGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplDllistIt
	var dllist_object *SplDllistObject = SplDllistFromObj(object.value.obj)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend._emalloc(g.SizeOf("spl_dllist_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.it.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.it.funcs = &SplDllistItFuncs
	iterator.intern.ce = ce
	iterator.SetTraversePosition(dllist_object.GetTraversePosition())
	iterator.SetTraversePointer(dllist_object.GetTraversePointer())
	iterator.SetFlags(dllist_object.GetFlags() & 0x3)
	&iterator.intern.value.u1.type_info = 0
	if iterator.GetTraversePointer() != nil {
		iterator.GetTraversePointer().GetRc()++
	}
	return &iterator.intern.it
}

/* }}} */

var ArginfoDllistSetiteratormode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoDllistPush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var arginfo_dllist_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_dllist_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var ArginfoDllistVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoDllistSerialized []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"serialized", 0, 0, 0}}
var spl_funcs_SplQueue []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"enqueue",
		zim_spl_SplDoublyLinkedList_push,
		ArginfoDllistPush,
		uint32(g.SizeOf("arginfo_dllist_push")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"dequeue",
		zim_spl_SplDoublyLinkedList_shift,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplDoublyLinkedList []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"pop",
		zim_spl_SplDoublyLinkedList_pop,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"shift",
		zim_spl_SplDoublyLinkedList_shift,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"push",
		zim_spl_SplDoublyLinkedList_push,
		ArginfoDllistPush,
		uint32(g.SizeOf("arginfo_dllist_push")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"unshift",
		zim_spl_SplDoublyLinkedList_unshift,
		ArginfoDllistPush,
		uint32(g.SizeOf("arginfo_dllist_push")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"top",
		zim_spl_SplDoublyLinkedList_top,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"bottom",
		zim_spl_SplDoublyLinkedList_bottom,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isEmpty",
		zim_spl_SplDoublyLinkedList_isEmpty,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setIteratorMode",
		zim_spl_SplDoublyLinkedList_setIteratorMode,
		ArginfoDllistSetiteratormode,
		uint32(g.SizeOf("arginfo_dllist_setiteratormode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getIteratorMode",
		zim_spl_SplDoublyLinkedList_getIteratorMode,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_SplDoublyLinkedList___debugInfo,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_SplDoublyLinkedList_count,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetExists",
		zim_spl_SplDoublyLinkedList_offsetExists,
		arginfo_dllist_offsetGet,
		uint32(g.SizeOf("arginfo_dllist_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetGet",
		zim_spl_SplDoublyLinkedList_offsetGet,
		arginfo_dllist_offsetGet,
		uint32(g.SizeOf("arginfo_dllist_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetSet",
		zim_spl_SplDoublyLinkedList_offsetSet,
		arginfo_dllist_offsetSet,
		uint32(g.SizeOf("arginfo_dllist_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetUnset",
		zim_spl_SplDoublyLinkedList_offsetUnset,
		arginfo_dllist_offsetGet,
		uint32(g.SizeOf("arginfo_dllist_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"add",
		zim_spl_SplDoublyLinkedList_add,
		arginfo_dllist_offsetSet,
		uint32(g.SizeOf("arginfo_dllist_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_SplDoublyLinkedList_rewind,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_SplDoublyLinkedList_current,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_SplDoublyLinkedList_key,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_SplDoublyLinkedList_next,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"prev",
		zim_spl_SplDoublyLinkedList_prev,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_SplDoublyLinkedList_valid,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"unserialize",
		zim_spl_SplDoublyLinkedList_unserialize,
		ArginfoDllistSerialized,
		uint32(g.SizeOf("arginfo_dllist_serialized")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"serialize",
		zim_spl_SplDoublyLinkedList_serialize,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__unserialize",
		zim_spl_SplDoublyLinkedList___unserialize,
		ArginfoDllistSerialized,
		uint32(g.SizeOf("arginfo_dllist_serialized")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__serialize",
		zim_spl_SplDoublyLinkedList___serialize,
		ArginfoDllistVoid,
		uint32(g.SizeOf("arginfo_dllist_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplDllist(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplDoublyLinkedList, "SplDoublyLinkedList", SplDllistObjectNew, spl_funcs_SplDoublyLinkedList)
	memcpy(&spl_handler_SplDoublyLinkedList, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	spl_handler_SplDoublyLinkedList.offset = zend_long((*byte)(&((*SplDllistObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplDoublyLinkedList.clone_obj = SplDllistObjectClone
	spl_handler_SplDoublyLinkedList.count_elements = SplDllistObjectCountElements
	spl_handler_SplDoublyLinkedList.get_gc = SplDllistObjectGetGc
	spl_handler_SplDoublyLinkedList.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplDoublyLinkedList.free_obj = SplDllistObjectFreeStorage
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_LIFO", g.SizeOf("\"IT_MODE_LIFO\"")-1, zend.ZendLong(0x2))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_FIFO", g.SizeOf("\"IT_MODE_FIFO\"")-1, zend.ZendLong(0))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_DELETE", g.SizeOf("\"IT_MODE_DELETE\"")-1, zend.ZendLong(0x1))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_KEEP", g.SizeOf("\"IT_MODE_KEEP\"")-1, zend.ZendLong(0))
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, zend.ZendCeCountable)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, zend.ZendCeArrayaccess)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, zend.ZendCeSerializable)
	spl_ce_SplDoublyLinkedList.get_iterator = SplDllistGetIterator
	SplRegisterSubClass(&spl_ce_SplQueue, spl_ce_SplDoublyLinkedList, "SplQueue", SplDllistObjectNew, spl_funcs_SplQueue)
	SplRegisterSubClass(&spl_ce_SplStack, spl_ce_SplDoublyLinkedList, "SplStack", SplDllistObjectNew, nil)
	spl_ce_SplQueue.get_iterator = SplDllistGetIterator
	spl_ce_SplStack.get_iterator = SplDllistGetIterator
	return zend.SUCCESS
}

/* }}} */

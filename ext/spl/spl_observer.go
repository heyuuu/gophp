// <<generate>>

package spl

import (
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/spl_observer.h>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define SPL_OBSERVER_H

// # include "php.h"

// # include "php_spl.h"

var spl_ce_SplObserver *zend.ZendClassEntry
var spl_ce_SplSubject *zend.ZendClassEntry
var spl_ce_SplObjectStorage *zend.ZendClassEntry
var spl_ce_MultipleIterator *zend.ZendClassEntry

// Source: <ext/spl/spl_observer.c>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   |          Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "ext/standard/php_array.h"

// # include "ext/standard/php_var.h"

// # include "zend_smart_str.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_observer.h"

// # include "spl_iterators.h"

// # include "spl_array.h"

// # include "spl_exceptions.h"

var zim_spl_SplObserver_update func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_attach func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_detach func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_notify func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var arginfo_SplObserver_update []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"subject", zend.ZendType("SplSubject"), 0, 0},
}
var spl_funcs_SplObserver []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"update",
		nil,
		arginfo_SplObserver_update,
		uint32(g.SizeOf("arginfo_SplObserver_update")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var arginfo_SplSubject_attach []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"observer", zend.ZendType("SplObserver"), 0, 0},
}
var arginfo_SplSubject_void []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/*ZEND_BEGIN_ARG_INFO_EX(arginfo_SplSubject_notify, 0, 0, 1)
    ZEND_ARG_OBJ_INFO(0, ignore, SplObserver, 1)
ZEND_END_ARG_INFO();*/

var spl_funcs_SplSubject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"attach",
		nil,
		arginfo_SplSubject_attach,
		uint32(g.SizeOf("arginfo_SplSubject_attach")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"detach",
		nil,
		arginfo_SplSubject_attach,
		uint32(g.SizeOf("arginfo_SplSubject_attach")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"notify",
		nil,
		arginfo_SplSubject_void,
		uint32(g.SizeOf("arginfo_SplSubject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var spl_handler_SplObjectStorage zend.ZendObjectHandlers

// @type spl_SplObjectStorage struct

/* {{{ storage is an assoc array of [zend_object*]=>[zval *obj, zval *inf] */

// @type spl_SplObjectStorageElement struct

func SplObjectStorageFromObj(obj *zend.ZendObject) *spl_SplObjectStorage {
	return (*spl_SplObjectStorage)((*byte)(obj - zend_long((*byte)(&((*spl_SplObjectStorage)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLOBJSTORAGE_P(zv) spl_object_storage_from_obj ( Z_OBJ_P ( ( zv ) ) )

func spl_SplObjectStorage_free_storage(object *zend.ZendObject) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(object)
	zend.ZendObjectStdDtor(&intern.std)
	zend.ZendHashDestroy(&intern.storage)
	if intern.GetGcdata() != nil {
		zend._efree(intern.GetGcdata())
	}
}
func SplObjectStorageGetHash(key *zend.ZendHashKey, intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	if intern.GetFptrGetHash() != nil {
		var rv zend.Zval
		zend.ZendCallMethod(this, intern.std.ce, &intern.fptr_get_hash, "getHash", g.SizeOf("\"getHash\"")-1, &rv, 1, obj, nil)
		if rv.u1.v.type_ != 0 {
			if rv.u1.v.type_ == 6 {
				key.key = rv.value.str
				return zend.SUCCESS
			} else {
				zend.ZendThrowException(spl_ce_RuntimeException, "Hash needs to be a string", 0)
				zend.ZvalPtrDtor(&rv)
				return zend.FAILURE
			}
		} else {
			return zend.FAILURE
		}
	} else {
		key.key = nil
		key.h = obj.value.obj.handle
		return zend.SUCCESS
	}
}
func SplObjectStorageFreeHash(intern *spl_SplObjectStorage, key *zend.ZendHashKey) {
	if key.key != nil {
		zend.ZendStringReleaseEx(key.key, 0)
	}
}
func SplObjectStorageDtor(element *zend.Zval) {
	var el *spl_SplObjectStorageElement = element.value.ptr
	zend.ZvalPtrDtor(&el.obj)
	zend.ZvalPtrDtor(&el.inf)
	zend._efree(el)
}
func SplObjectStorageGet(intern *spl_SplObjectStorage, key *zend.ZendHashKey) *spl_SplObjectStorageElement {
	if key.key != nil {
		return zend.ZendHashFindPtr(&intern.storage, key.key)
	} else {
		return zend.ZendHashIndexFindPtr(&intern.storage, key.h)
	}
}
func SplObjectStorageAttach(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval, inf *zend.Zval) *spl_SplObjectStorageElement {
	var pelement *spl_SplObjectStorageElement
	var element spl_SplObjectStorageElement
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return nil
	}
	pelement = SplObjectStorageGet(intern, &key)
	if pelement != nil {
		zend.ZvalPtrDtor(&pelement.inf)
		if inf != nil {
			var _z1 *zend.Zval = &pelement.inf
			var _z2 *zend.Zval = inf
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		} else {
			&pelement.inf.u1.type_info = 1
		}
		SplObjectStorageFreeHash(intern, &key)
		return pelement
	}
	var _z1 *zend.Zval = &element.obj
	var _z2 *zend.Zval = obj
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	if inf != nil {
		var _z1 *zend.Zval = &element.inf
		var _z2 *zend.Zval = inf
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	} else {
		&element.inf.u1.type_info = 1
	}
	if key.key != nil {
		pelement = zend.ZendHashUpdateMem(&intern.storage, key.key, &element, g.SizeOf("spl_SplObjectStorageElement"))
	} else {
		pelement = zend.ZendHashIndexUpdateMem(&intern.storage, key.h, &element, g.SizeOf("spl_SplObjectStorageElement"))
	}
	SplObjectStorageFreeHash(intern, &key)
	return pelement
}
func SplObjectStorageDetach(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	var ret int = zend.FAILURE
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return ret
	}
	if key.key != nil {
		ret = zend.ZendHashDel(&intern.storage, key.key)
	} else {
		ret = zend.ZendHashIndexDel(&intern.storage, key.h)
	}
	SplObjectStorageFreeHash(intern, &key)
	return ret
}
func SplObjectStorageAddall(intern *spl_SplObjectStorage, this *zend.Zval, other *spl_SplObjectStorage) {
	var element *spl_SplObjectStorageElement
	for {
		var __ht *zend.HashTable = &other.storage
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			element = _z.value.ptr
			SplObjectStorageAttach(intern, this, &element.obj, &element.inf)
		}
		break
	}
	intern.SetIndex(0)
}
func SplObjectStorageNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval) *zend.ZendObject {
	var intern *spl_SplObjectStorage
	var parent *zend.ZendClassEntry = class_type
	intern = zend._emalloc(g.SizeOf("spl_SplObjectStorage") + zend.ZendObjectPropertiesSize(parent))
	memset(intern, 0, g.SizeOf("spl_SplObjectStorage")-g.SizeOf("zval"))
	intern.SetPos(0)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	zend._zendHashInit(&intern.storage, 0, SplObjectStorageDtor, 0)
	intern.std.handlers = &spl_handler_SplObjectStorage
	for parent != nil {
		if parent == spl_ce_SplObjectStorage {
			if class_type != spl_ce_SplObjectStorage {
				intern.SetFptrGetHash(zend.ZendHashStrFindPtr(&class_type.function_table, "gethash", g.SizeOf("\"gethash\"")-1))
				if intern.GetFptrGetHash().common.scope == spl_ce_SplObjectStorage {
					intern.SetFptrGetHash(nil)
				}
			}
			break
		}
		parent = parent.parent
	}
	if orig != nil {
		var other *spl_SplObjectStorage = SplObjectStorageFromObj(orig.value.obj)
		SplObjectStorageAddall(intern, orig, other)
	}
	return &intern.std
}

/* }}} */

func SplObjectStorageClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.value.obj
	new_object = SplObjectStorageNewEx(old_object.ce, zobject)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}

/* }}} */

func SplObjectStorageDebugInfo(obj *zend.Zval) *zend.HashTable {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(obj.value.obj)
	var element *spl_SplObjectStorageElement
	var props *zend.HashTable
	var tmp zend.Zval
	var storage zend.Zval
	var md5str *zend.ZendString
	var zname *zend.ZendString
	var debug_info *zend.HashTable
	props = obj.value.obj.handlers.get_properties(&(*obj))
	debug_info = zend._zendNewArray(props.nNumOfElements + 1)
	zend.ZendHashCopy(debug_info, props, zend.CopyCtorFuncT(zend.ZvalAddRef))
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &storage
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = &intern.storage
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			element = _z.value.ptr
			md5str = PhpSplObjectHash(&element.obj)
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = &tmp
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

			/* Incrementing the refcount of obj and inf would confuse the garbage collector.
			 * Prefer to null the destructor */

			&tmp.value.arr.pDestructor = nil
			zend.AddAssocZvalEx(&tmp, "obj", g.SizeOf("\"obj\"")-1, &element.obj)
			zend.AddAssocZvalEx(&tmp, "inf", g.SizeOf("\"inf\"")-1, &element.inf)
			zend.ZendHashUpdate(storage.value.arr, md5str, &tmp)
			zend.ZendStringReleaseEx(md5str, 0)
		}
		break
	}
	zname = SplGenPrivatePropName(spl_ce_SplObjectStorage, "storage", g.SizeOf("\"storage\"")-1)
	zend.ZendSymtableUpdate(debug_info, zname, &storage)
	zend.ZendStringReleaseEx(zname, 0)
	return debug_info
}

/* }}} */

func SplObjectStorageGetGc(obj *zend.Zval, table **zend.Zval, n *int) *zend.HashTable {
	var i int = 0
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(obj.value.obj)
	var element *spl_SplObjectStorageElement
	if intern.storage.nNumOfElements*2 > intern.GetGcdataNum() {
		intern.SetGcdataNum(intern.storage.nNumOfElements * 2)
		intern.SetGcdata((*zend.Zval)(zend._erealloc(intern.GetGcdata(), g.SizeOf("zval")*intern.GetGcdataNum())))
	}
	for {
		var __ht *zend.HashTable = &intern.storage
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			element = _z.value.ptr
			var _z1 *zend.Zval = &intern.gcdata[g.PostInc(&i)]
			var _z2 *zend.Zval = &element.obj
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			var _z1 *zend.Zval = &intern.gcdata[g.PostInc(&i)]
			var _z2 *zend.Zval = &element.inf
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		}
		break
	}
	*table = intern.GetGcdata()
	*n = i
	return zend.ZendStdGetProperties(obj)
}

/* }}} */

func SplObjectStorageCompareInfo(e1 *zend.Zval, e2 *zend.Zval) int {
	var s1 *spl_SplObjectStorageElement = (*spl_SplObjectStorageElement)(e1.value.ptr)
	var s2 *spl_SplObjectStorageElement = (*spl_SplObjectStorageElement)(e2.value.ptr)
	var result zend.Zval
	if zend.CompareFunction(&result, &s1.inf, &s2.inf) == zend.FAILURE {
		return 1
	}
	if result.value.lval != 0 {
		if result.value.lval < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func SplObjectStorageCompareObjects(o1 *zend.Zval, o2 *zend.Zval) int {
	var zo1 *zend.ZendObject = (*zend.ZendObject)(o1.value.obj)
	var zo2 *zend.ZendObject = (*zend.ZendObject)(o2.value.obj)
	if zo1.ce != spl_ce_SplObjectStorage || zo2.ce != spl_ce_SplObjectStorage {
		return 1
	}
	return zend.ZendHashCompare(&(SplObjectStorageFromObj(o1.value.obj)).storage, &(SplObjectStorageFromObj(o2.value.obj)).storage, zend.CompareFuncT(SplObjectStorageCompareInfo), 0)
}

/* }}} */

func spl_SplObjectStorage_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplObjectStorageNewEx(class_type, nil)
}

/* }}} */

func SplObjectStorageContains(intern *spl_SplObjectStorage, this *zend.Zval, obj *zend.Zval) int {
	var found int
	var key zend.ZendHashKey
	if SplObjectStorageGetHash(&key, intern, this, obj) == zend.FAILURE {
		return 0
	}
	if key.key != nil {
		found = zend.ZendHashExists(&intern.storage, key.key)
	} else {
		found = zend.ZendHashIndexExists(&intern.storage, key.h)
	}
	SplObjectStorageFreeHash(intern, &key)
	return found
}

/* {{{ proto void SplObjectStorage::attach(object obj, mixed data = NULL)
Attaches an object to the storage if not yet contained */

func zim_spl_SplObjectStorage_attach(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var inf *zend.Zval = nil
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o|z!", &obj, &inf) == zend.FAILURE {
		return
	}
	SplObjectStorageAttach(intern, &(execute_data.This), obj, inf)
}

/* {{{ proto void SplObjectStorage::detach(object obj)
Detaches an object from the storage */

func zim_spl_SplObjectStorage_detach(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o", &obj) == zend.FAILURE {
		return
	}
	SplObjectStorageDetach(intern, &(execute_data.This), obj)
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	intern.SetIndex(0)
}

/* {{{ proto string SplObjectStorage::getHash(object obj)
Returns the hash of an object */

func zim_spl_SplObjectStorage_getHash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o", &obj) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpSplObjectHash(obj)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* {{{ proto mixed SplObjectStorage::offsetGet(object obj)
Returns associated information for a stored object */

func zim_spl_SplObjectStorage_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var key zend.ZendHashKey
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o", &obj) == zend.FAILURE {
		return
	}
	if SplObjectStorageGetHash(&key, intern, &(execute_data.This), obj) == zend.FAILURE {
		return
	}
	element = SplObjectStorageGet(intern, &key)
	SplObjectStorageFreeHash(intern, &key)
	if element == nil {
		zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Object not found")
	} else {
		var value *zend.Zval = &element.inf
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

/* {{{ proto bool SplObjectStorage::addAll(SplObjectStorage $os)
Add all elements contained in $os */

func zim_spl_SplObjectStorage_addAll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var other *spl_SplObjectStorage
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = SplObjectStorageFromObj(obj.value.obj)
	SplObjectStorageAddall(intern, &(execute_data.This), other)
	var __z *zend.Zval = return_value
	__z.value.lval = &intern.storage.nNumOfElements
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool SplObjectStorage::removeAll(SplObjectStorage $os)
Remove all elements contained in $os */

func zim_spl_SplObjectStorage_removeAll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var other *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = SplObjectStorageFromObj(obj.value.obj)
	zend.ZendHashInternalPointerResetEx(&other.storage, &(&other.storage).nInternalPointer)
	for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&other.storage, &(&other.storage).nInternalPointer)) != nil {
		if SplObjectStorageDetach(intern, &(execute_data.This), &element.obj) == zend.FAILURE {
			zend.ZendHashMoveForwardEx(&other.storage, &(&other.storage).nInternalPointer)
		}
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	intern.SetIndex(0)
	var __z *zend.Zval = return_value
	__z.value.lval = &intern.storage.nNumOfElements
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool SplObjectStorage::removeAllExcept(SplObjectStorage $os)
Remove elements not common to both this SplObjectStorage instance and $os */

func zim_spl_SplObjectStorage_removeAllExcept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var other *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O", &obj, spl_ce_SplObjectStorage) == zend.FAILURE {
		return
	}
	other = SplObjectStorageFromObj(obj.value.obj)
	for {
		var __ht *zend.HashTable = &intern.storage
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			element = _z.value.ptr
			if SplObjectStorageContains(other, &(execute_data.This), &element.obj) == 0 {
				SplObjectStorageDetach(intern, &(execute_data.This), &element.obj)
			}
		}
		break
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	intern.SetIndex(0)
	var __z *zend.Zval = return_value
	__z.value.lval = &intern.storage.nNumOfElements
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplObjectStorage_contains(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o", &obj) == zend.FAILURE {
		return
	}
	if SplObjectStorageContains(intern, &(execute_data.This), obj) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto int SplObjectStorage::count()
Determine number of objects in storage */

func zim_spl_SplObjectStorage_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var mode zend.ZendLong = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|l", &mode) == zend.FAILURE {
		return
	}
	if mode == 1 {
		var ret zend.ZendLong
		if mode != 1 {
			ret = &intern.storage.nNumOfElements
		} else {
			ret = standard.PhpCountRecursive(&intern.storage)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = ret
		__z.u1.type_info = 4
		return
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = &intern.storage.nNumOfElements
	__z.u1.type_info = 4
	return
}

/* {{{ proto void SplObjectStorage::rewind()
Rewind to first position */

func zim_spl_SplObjectStorage_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	intern.SetIndex(0)
}

/* {{{ proto bool SplObjectStorage::valid()
Returns whether current position is valid */

func zim_spl_SplObjectStorage_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Cond(zend.ZendHashGetCurrentKeyTypeEx(&intern.storage, &intern.pos) == 3, zend.FAILURE, zend.SUCCESS) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed SplObjectStorage::key()
Returns current key */

func zim_spl_SplObjectStorage_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetIndex()
	__z.u1.type_info = 4
	return
}

/* {{{ proto mixed SplObjectStorage::current()
Returns current element */

func zim_spl_SplObjectStorage_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) == nil {
		return
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = &element.obj
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
}

/* {{{ proto mixed SplObjectStorage::getInfo()
Returns associated information to current element */

func zim_spl_SplObjectStorage_getInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) == nil {
		return
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = &element.inf
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
}

/* {{{ proto mixed SplObjectStorage::setInfo(mixed $inf)
Sets associated information of current element to $inf */

func zim_spl_SplObjectStorage_setInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var inf *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &inf) == zend.FAILURE {
		return
	}
	if g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) == nil {
		return
	}
	zend.ZvalPtrDtor(&element.inf)
	var _z1 *zend.Zval = &element.inf
	var _z2 *zend.Zval = inf
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
}

/* {{{ proto void SplObjectStorage::next()
Moves position forward */

func zim_spl_SplObjectStorage_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
	intern.GetIndex()++
}

/* {{{ proto string SplObjectStorage::serialize()
Serializes storage */

func zim_spl_SplObjectStorage_serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var element *spl_SplObjectStorageElement
	var members zend.Zval
	var flags zend.Zval
	var pos zend.HashPosition
	var var_hash standard.PhpSerializeDataT
	var buf zend.SmartStr = zend.SmartStr{0}
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var_hash = standard.PhpVarSerializeInit()

	/* storage */

	zend.SmartStrAppendlEx(&buf, "x:", 2, 0)
	var __z *zend.Zval = &flags
	__z.value.lval = &intern.storage.nNumOfElements
	__z.u1.type_info = 4
	standard.PhpVarSerialize(&buf, &flags, &var_hash)
	zend.ZendHashInternalPointerResetEx(&intern.storage, &pos)
	for g.Cond(zend.ZendHashGetCurrentKeyTypeEx(&intern.storage, &pos) == 3, zend.FAILURE, zend.SUCCESS) == zend.SUCCESS {
		if g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &pos)) == nil {
			zend.SmartStrFreeEx(&buf, 0)
			standard.PhpVarSerializeDestroy(var_hash)
			return_value.u1.type_info = 1
			return
		}
		standard.PhpVarSerialize(&buf, &element.obj, &var_hash)
		zend.SmartStrAppendcEx(&buf, ',', 0)
		standard.PhpVarSerialize(&buf, &element.inf, &var_hash)
		zend.SmartStrAppendcEx(&buf, ';', 0)
		zend.ZendHashMoveForwardEx(&intern.storage, &pos)
	}

	/* members */

	zend.SmartStrAppendlEx(&buf, "m:", 2, 0)
	var __arr *zend.ZendArray = zend.ZendArrayDup(zend.ZendStdGetProperties(&(execute_data.This)))
	var __z *zend.Zval = &members
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	standard.PhpVarSerialize(&buf, &members, &var_hash)
	zend.ZvalPtrDtor(&members)

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

/* {{{ proto void SplObjectStorage::unserialize(string serialized)
Unserializes storage */

func zim_spl_SplObjectStorage_unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var entry zend.Zval
	var inf zend.Zval
	var pcount *zend.Zval
	var pmembers *zend.Zval
	var element *spl_SplObjectStorageElement
	var count zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}

	/* storage */

	p = (*uint8)(buf)
	s = p
	var_hash = standard.PhpVarUnserializeInit()
	if (*p) != 'x' || (*(g.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	pcount = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(pcount, &p, s+buf_len, &var_hash) == 0 || pcount.u1.v.type_ != 4 {
		goto outexcept
	}
	p--
	count = pcount.value.lval
	if count < 0 {
		goto outexcept
	}
	&entry.u1.type_info = 0
	&inf.u1.type_info = 0
	for g.PostDec(&count) > 0 {
		var pelement *spl_SplObjectStorageElement
		var key zend.ZendHashKey
		if (*p) != ';' {
			goto outexcept
		}
		p++
		if (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}

		/* store reference to allow cross-references between different elements */

		if standard.PhpVarUnserialize(&entry, &p, s+buf_len, &var_hash) == 0 {
			zend.ZvalPtrDtor(&entry)
			goto outexcept
		}
		if (*p) == ',' {
			p++
			if standard.PhpVarUnserialize(&inf, &p, s+buf_len, &var_hash) == 0 {
				zend.ZvalPtrDtor(&entry)
				zend.ZvalPtrDtor(&inf)
				goto outexcept
			}
		}
		if entry.u1.v.type_ != 8 {
			zend.ZvalPtrDtor(&entry)
			zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		if SplObjectStorageGetHash(&key, intern, &(execute_data.This), &entry) == zend.FAILURE {
			zend.ZvalPtrDtor(&entry)
			zend.ZvalPtrDtor(&inf)
			goto outexcept
		}
		pelement = SplObjectStorageGet(intern, &key)
		SplObjectStorageFreeHash(intern, &key)
		if pelement != nil {
			if pelement.inf.u1.v.type_ != 0 {
				standard.VarPushDtor(&var_hash, &pelement.inf)
			}
			if pelement.obj.u1.v.type_ != 0 {
				standard.VarPushDtor(&var_hash, &pelement.obj)
			}
		}
		element = SplObjectStorageAttach(intern, &(execute_data.This), &entry, g.Cond(inf.u1.v.type_ == 0, nil, &inf))
		standard.VarReplace(&var_hash, &entry, &element.obj)
		standard.VarReplace(&var_hash, &inf, &element.inf)
		zend.ZvalPtrDtor(&entry)
		&entry.u1.type_info = 0
		zend.ZvalPtrDtor(&inf)
		&inf.u1.type_info = 0
	}
	if (*p) != ';' {
		goto outexcept
	}
	p++

	/* members */

	if (*p) != 'm' || (*(g.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	pmembers = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(pmembers, &p, s+buf_len, &var_hash) == 0 || pmembers.u1.v.type_ != 7 {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(&intern.std, pmembers.value.arr)
	standard.PhpVarUnserializeDestroy(var_hash)
	return
outexcept:
	standard.PhpVarUnserializeDestroy(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset %zd of %zd bytes", (*byte)(p-buf), buf_len)
	return
}

/* {{{ proto auto SplObjectStorage::__serialize() */

func zim_spl_SplObjectStorage___serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var elem *spl_SplObjectStorageElement
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

	/* storage */

	var __arr *zend.ZendArray = zend._zendNewArray(2 * &intern.storage.nNumOfElements)
	var __z *zend.Zval = &tmp
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = &intern.storage
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			elem = _z.value.ptr
			if &(elem.GetObj()).u1.v.type_flags != 0 {
				zend.ZvalAddrefP(&(elem.GetObj()))
			}
			zend.ZendHashNextIndexInsert(tmp.value.arr, &elem.obj)
			if &(elem.GetInf()).u1.v.type_flags != 0 {
				zend.ZvalAddrefP(&(elem.GetInf()))
			}
			zend.ZendHashNextIndexInsert(tmp.value.arr, &elem.inf)
		}
		break
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

/* {{{ proto void SplObjectStorage::__unserialize(array serialized) */

func zim_spl_SplObjectStorage___unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	var data *zend.HashTable
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var key *zend.Zval
	var val *zend.Zval
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "h", &data) == zend.FAILURE {
		return
	}
	storage_zv = zend.ZendHashIndexFind(data, 0)
	members_zv = zend.ZendHashIndexFind(data, 1)
	if storage_zv == nil || members_zv == nil || storage_zv.u1.v.type_ != 7 || members_zv.u1.v.type_ != 7 {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	if storage_zv.value.arr.nNumOfElements%2 != 0 {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Odd number of elements", 0)
		return
	}
	key = nil
	for {
		var __ht *zend.HashTable = storage_zv.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			val = _z
			if key != nil {
				if key.u1.v.type_ != 8 {
					zend.ZendThrowException(spl_ce_UnexpectedValueException, "Non-object key", 0)
					return
				}
				SplObjectStorageAttach(intern, &(execute_data.This), key, val)
				key = nil
			} else {
				key = val
			}
		}
		break
	}
	zend.ObjectPropertiesLoad(&intern.std, members_zv.value.arr)
}

/* {{{ proto array SplObjectStorage::__debugInfo() */

func zim_spl_SplObjectStorage___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplObjectStorageDebugInfo(g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* }}} */

var arginfo_Object []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"object", 0, 0, 0}}
var ArginfoAttach []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"object", 0, 0, 0}, {"data", 0, 0, 0}}
var arginfo_Serialized []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"serialized", 0, 0, 0}}
var arginfo_setInfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"info", 0, 0, 0}}
var arginfo_getHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"object", 0, 0, 0}}
var arginfo_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"object", 0, 0, 0}}
var ArginfoSplobjectVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var spl_funcs_SplObjectStorage []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"attach",
		zim_spl_SplObjectStorage_attach,
		ArginfoAttach,
		uint32(g.SizeOf("arginfo_attach")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"detach",
		zim_spl_SplObjectStorage_detach,
		arginfo_Object,
		uint32(g.SizeOf("arginfo_Object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"contains",
		zim_spl_SplObjectStorage_contains,
		arginfo_Object,
		uint32(g.SizeOf("arginfo_Object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addAll",
		zim_spl_SplObjectStorage_addAll,
		arginfo_Object,
		uint32(g.SizeOf("arginfo_Object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"removeAll",
		zim_spl_SplObjectStorage_removeAll,
		arginfo_Object,
		uint32(g.SizeOf("arginfo_Object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"removeAllExcept",
		zim_spl_SplObjectStorage_removeAllExcept,
		arginfo_Object,
		uint32(g.SizeOf("arginfo_Object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getInfo",
		zim_spl_SplObjectStorage_getInfo,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setInfo",
		zim_spl_SplObjectStorage_setInfo,
		arginfo_setInfo,
		uint32(g.SizeOf("arginfo_setInfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getHash",
		zim_spl_SplObjectStorage_getHash,
		arginfo_getHash,
		uint32(g.SizeOf("arginfo_getHash")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__debugInfo",
		zim_spl_SplObjectStorage___debugInfo,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count",
		zim_spl_SplObjectStorage_count,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		zim_spl_SplObjectStorage_rewind,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"valid",
		zim_spl_SplObjectStorage_valid,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		zim_spl_SplObjectStorage_key,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		zim_spl_SplObjectStorage_current,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		zim_spl_SplObjectStorage_next,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unserialize",
		zim_spl_SplObjectStorage_unserialize,
		arginfo_Serialized,
		uint32(g.SizeOf("arginfo_Serialized")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"serialize",
		zim_spl_SplObjectStorage_serialize,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__unserialize",
		zim_spl_SplObjectStorage___unserialize,
		arginfo_Serialized,
		uint32(g.SizeOf("arginfo_Serialized")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__serialize",
		zim_spl_SplObjectStorage___serialize,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetExists",
		zim_spl_SplObjectStorage_contains,
		arginfo_offsetGet,
		uint32(g.SizeOf("arginfo_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetSet",
		zim_spl_SplObjectStorage_attach,
		ArginfoAttach,
		uint32(g.SizeOf("arginfo_attach")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetUnset",
		zim_spl_SplObjectStorage_detach,
		arginfo_offsetGet,
		uint32(g.SizeOf("arginfo_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetGet",
		zim_spl_SplObjectStorage_offsetGet,
		arginfo_offsetGet,
		uint32(g.SizeOf("arginfo_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

type MultipleIteratorFlags = int

const (
	MIT_NEED_ANY     = 0
	MIT_NEED_ALL     = 1
	MIT_KEYS_NUMERIC = 0
	MIT_KEYS_ASSOC   = 2
)

// #define SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT       1

// #define SPL_MULTIPLE_ITERATOR_GET_ALL_KEY       2

/* {{{ proto MultipleIterator::__construct([int flags = MIT_NEED_ALL|MIT_KEYS_NUMERIC])
   Iterator that iterates over several iterators one after the other */

func zim_spl_MultipleIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var flags zend.ZendLong = MIT_NEED_ALL | MIT_KEYS_NUMERIC
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "|l", &flags) == zend.FAILURE {
		return
	}
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	intern.SetFlags(flags)
}

/* }}} */

func zim_spl_MultipleIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_MultipleIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &intern.flags) == zend.FAILURE {
		return
	}
}

/* }}} */

func zim_spl_MultipleIterator_attachIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var iterator *zend.Zval = nil
	var info *zend.Zval = nil
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O|z!", &iterator, zend.ZendCeIterator, &info) == zend.FAILURE {
		return
	}
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if info != nil {
		var element *spl_SplObjectStorageElement
		if info.u1.v.type_ != 4 && info.u1.v.type_ != 6 {
			zend.ZendThrowException(spl_ce_InvalidArgumentException, "Info must be NULL, integer or string", 0)
			return
		}
		zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
		for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) != nil {
			if zend.FastIsIdenticalFunction(info, &element.inf) != 0 {
				zend.ZendThrowException(spl_ce_InvalidArgumentException, "Key duplication error", 0)
				return
			}
			zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
		}
	}
	SplObjectStorageAttach(intern, &(execute_data.This), iterator, info)
}

/* }}} */

func zim_spl_MultipleIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) != nil && zend.EG.exception == nil {
		it = &element.obj
		zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_rewind, "rewind", g.SizeOf("\"rewind\"")-1, nil, 0, nil, nil)
		zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
	}
}

/* }}} */

func zim_spl_MultipleIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) != nil && zend.EG.exception == nil {
		it = &element.obj
		zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_next, "next", g.SizeOf("\"next\"")-1, nil, 0, nil, nil)
		zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
	}
}

/* }}} */

func zim_spl_MultipleIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	var retval zend.Zval
	var expect zend.ZendLong
	var valid zend.ZendLong
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if !(&intern.storage.nNumOfElements) {
		return_value.u1.type_info = 2
		return
	}
	if (intern.GetFlags() & MIT_NEED_ALL) != 0 {
		expect = 1
	} else {
		expect = 0
	}
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) != nil && zend.EG.exception == nil {
		it = &element.obj
		zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_valid, "valid", g.SizeOf("\"valid\"")-1, &retval, 0, nil, nil)
		if retval.u1.v.type_ != 0 {
			valid = retval.u1.v.type_ == 3
			zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if expect != valid {
			if expect == 0 {
				return_value.u1.type_info = 3
			} else {
				return_value.u1.type_info = 2
			}
			return
		}
		zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
	}
	if expect != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func SplMultipleIteratorGetAll(intern *spl_SplObjectStorage, get_type int, return_value *zend.Zval) {
	var element *spl_SplObjectStorageElement
	var it *zend.Zval
	var retval zend.Zval
	var valid int = 1
	var num_elements int
	num_elements = &intern.storage.nNumOfElements
	if num_elements < 1 {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(num_elements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZendHashInternalPointerResetEx(&intern.storage, &intern.pos)
	for g.Assign(&element, zend.ZendHashGetCurrentDataPtrEx(&intern.storage, &intern.pos)) != nil && zend.EG.exception == nil {
		it = &element.obj
		zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_valid, "valid", g.SizeOf("\"valid\"")-1, &retval, 0, nil, nil)
		if retval.u1.v.type_ != 0 {
			valid = retval.u1.v.type_ == 3
			zend.ZvalPtrDtor(&retval)
		} else {
			valid = 0
		}
		if valid != 0 {
			if 1 == get_type {
				zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_current, "current", g.SizeOf("\"current\"")-1, &retval, 0, nil, nil)
			} else {
				zend.ZendCallMethod(it, it.value.obj.ce, &(it.value.obj.ce).iterator_funcs_ptr.zf_key, "key", g.SizeOf("\"key\"")-1, &retval, 0, nil, nil)
			}
			if retval.u1.v.type_ == 0 {
				zend.ZendThrowException(spl_ce_RuntimeException, "Failed to call sub iterator method", 0)
				return
			}
		} else if (intern.GetFlags() & MIT_NEED_ALL) != 0 {
			if 1 == get_type {
				zend.ZendThrowException(spl_ce_RuntimeException, "Called current() with non valid sub iterator", 0)
			} else {
				zend.ZendThrowException(spl_ce_RuntimeException, "Called key() with non valid sub iterator", 0)
			}
			return
		} else {
			&retval.u1.type_info = 1
		}
		if (intern.GetFlags() & MIT_KEYS_ASSOC) != 0 {
			switch element.inf.u1.v.type_ {
			case 4:
				zend.AddIndexZval(return_value, element.inf.value.lval, &retval)
				break
			case 6:
				zend.ZendSymtableUpdate(return_value.value.arr, element.inf.value.str, &retval)
				break
			default:
				zend.ZvalPtrDtor(&retval)
				zend.ZendThrowException(spl_ce_InvalidArgumentException, "Sub-Iterator is associated with NULL", 0)
				return
			}
		} else {
			zend.AddNextIndexZval(return_value, &retval)
		}
		zend.ZendHashMoveForwardEx(&intern.storage, &intern.pos)
	}
}

/* }}} */

func zim_spl_MultipleIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplMultipleIteratorGetAll(intern, 1, return_value)
}

/* }}} */

func zim_spl_MultipleIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *spl_SplObjectStorage
	intern = SplObjectStorageFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplMultipleIteratorGetAll(intern, 2, return_value)
}

/* }}} */

var arginfo_MultipleIterator_attachIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
	{"infos", 0, 0, 0},
}
var arginfo_MultipleIterator_detachIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
}
var arginfo_MultipleIterator_containsIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
}
var arginfo_MultipleIterator_setflags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var spl_funcs_MultipleIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_MultipleIterator___construct,
		arginfo_MultipleIterator_setflags,
		uint32(g.SizeOf("arginfo_MultipleIterator_setflags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getFlags",
		zim_spl_MultipleIterator_getFlags,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setFlags",
		zim_spl_MultipleIterator_setFlags,
		arginfo_MultipleIterator_setflags,
		uint32(g.SizeOf("arginfo_MultipleIterator_setflags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"attachIterator",
		zim_spl_MultipleIterator_attachIterator,
		arginfo_MultipleIterator_attachIterator,
		uint32(g.SizeOf("arginfo_MultipleIterator_attachIterator")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"detachIterator",
		zim_spl_SplObjectStorage_detach,
		arginfo_MultipleIterator_detachIterator,
		uint32(g.SizeOf("arginfo_MultipleIterator_detachIterator")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"containsIterator",
		zim_spl_SplObjectStorage_contains,
		arginfo_MultipleIterator_containsIterator,
		uint32(g.SizeOf("arginfo_MultipleIterator_containsIterator")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"countIterators",
		zim_spl_SplObjectStorage_count,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__debugInfo",
		zim_spl_SplObjectStorage___debugInfo,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		zim_spl_MultipleIterator_rewind,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"valid",
		zim_spl_MultipleIterator_valid,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		zim_spl_MultipleIterator_key,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		zim_spl_MultipleIterator_current,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		zim_spl_MultipleIterator_next,
		ArginfoSplobjectVoid,
		uint32(g.SizeOf("arginfo_splobject_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */

func ZmStartupSplObserver(type_ int, module_number int) int {
	SplRegisterInterface(&spl_ce_SplObserver, "SplObserver", spl_funcs_SplObserver)
	SplRegisterInterface(&spl_ce_SplSubject, "SplSubject", spl_funcs_SplSubject)
	SplRegisterStdClass(&spl_ce_SplObjectStorage, "SplObjectStorage", spl_SplObjectStorage_new, spl_funcs_SplObjectStorage)
	memcpy(&spl_handler_SplObjectStorage, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	spl_handler_SplObjectStorage.offset = zend_long((*byte)(&((*spl_SplObjectStorage)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplObjectStorage.compare_objects = SplObjectStorageCompareObjects
	spl_handler_SplObjectStorage.clone_obj = SplObjectStorageClone
	spl_handler_SplObjectStorage.get_gc = SplObjectStorageGetGc
	spl_handler_SplObjectStorage.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplObjectStorage.free_obj = spl_SplObjectStorage_free_storage
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, zend.ZendCeCountable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, zend.ZendCeSerializable)
	zend.ZendClassImplements(spl_ce_SplObjectStorage, 1, zend.ZendCeArrayaccess)
	SplRegisterStdClass(&spl_ce_MultipleIterator, "MultipleIterator", spl_SplObjectStorage_new, spl_funcs_MultipleIterator)
	zend.ZendClassImplements(spl_ce_MultipleIterator, 1, zend.ZendCeIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ANY", g.SizeOf("\"MIT_NEED_ANY\"")-1, zend.ZendLong(MIT_NEED_ANY))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_NEED_ALL", g.SizeOf("\"MIT_NEED_ALL\"")-1, zend.ZendLong(MIT_NEED_ALL))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_NUMERIC", g.SizeOf("\"MIT_KEYS_NUMERIC\"")-1, zend.ZendLong(MIT_KEYS_NUMERIC))
	zend.ZendDeclareClassConstantLong(spl_ce_MultipleIterator, "MIT_KEYS_ASSOC", g.SizeOf("\"MIT_KEYS_ASSOC\"")-1, zend.ZendLong(MIT_KEYS_ASSOC))
	return zend.SUCCESS
}

/* }}} */

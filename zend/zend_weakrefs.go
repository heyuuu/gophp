// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_weakrefs.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: krakjoe@php.net                                             |
   +----------------------------------------------------------------------+
*/

// #define ZEND_WEAKREFS_H

var ZendCeWeakref *ZendClassEntry

// Source: <Zend/zend_weakrefs.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: krakjoe@php.net                                             |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_interfaces.h"

// # include "zend_objects_API.h"

// # include "zend_weakrefs.h"

var ZendWeakrefHandlers ZendObjectHandlers

func ZendWeakrefFrom(o *ZendObject) *ZendWeakref {
	return (*ZendWeakref)((*byte)(o) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
}
func ZendWeakrefFetch(z *Zval) *ZendWeakref { return ZendWeakrefFrom(Z_OBJ_P(z)) }
func ZendWeakrefUnref(zv *Zval) {
	var wr *ZendWeakref = (*ZendWeakref)(Z_PTR_P(zv))
	GC_DEL_FLAGS(wr.GetReferent(), IS_OBJ_WEAKLY_REFERENCED)
	wr.SetReferent(nil)
}
func ZendWeakrefsInit() {
	ZendHashInit(&(ExecutorGlobals.GetWeakrefs()), 8, nil, ZendWeakrefUnref, 0)
}
func ZendWeakrefsNotify(object *ZendObject) {
	ZendHashIndexDel(&(ExecutorGlobals.GetWeakrefs()), ZendUlong(object))
}
func ZendWeakrefsShutdown() {
	ZendHashDestroy(&(ExecutorGlobals.GetWeakrefs()))
}
func ZendWeakrefNew(ce *ZendClassEntry) *ZendObject {
	var wr *ZendWeakref = ZendObjectAlloc(b.SizeOf("zend_weakref"), ZendCeWeakref)
	ZendObjectStdInit(&wr.std, ZendCeWeakref)
	wr.GetStd().SetHandlers(&ZendWeakrefHandlers)
	return &wr.std
}
func ZendWeakrefFind(referent *Zval, return_value *Zval) ZendBool {
	var wr *ZendWeakref = ZendHashIndexFindPtr(&(ExecutorGlobals.GetWeakrefs()), ZendUlong(Z_OBJ_P(referent)))
	if wr == nil {
		return 0
	}
	GC_ADDREF(&wr.std)
	ZVAL_OBJ(return_value, &wr.std)
	return 1
}
func ZendWeakrefCreate(referent *Zval, return_value *Zval) {
	var wr *ZendWeakref
	ObjectInitEx(return_value, ZendCeWeakref)
	wr = ZendWeakrefFetch(return_value)
	wr.SetReferent(Z_OBJ_P(referent))
	ZendHashIndexAddPtr(&(ExecutorGlobals.GetWeakrefs()), ZendUlong(wr.GetReferent()), wr)
	GC_ADD_FLAGS(wr.GetReferent(), IS_OBJ_WEAKLY_REFERENCED)
}
func ZendWeakrefGet(weakref *Zval, return_value *Zval) {
	var wr *ZendWeakref = ZendWeakrefFetch(weakref)
	if wr.GetReferent() != nil {
		ZVAL_OBJ(return_value, wr.GetReferent())
		Z_ADDREF_P(return_value)
	}
}
func ZendWeakrefFree(zo *ZendObject) {
	var wr *ZendWeakref = ZendWeakrefFrom(zo)
	if wr.GetReferent() != nil {
		ZendHashIndexDel(&(ExecutorGlobals.GetWeakrefs()), ZendUlong(wr.GetReferent()))
	}
	ZendObjectStdDtor(&wr.std)
}
func ZendWeakrefUnsupported(thing string) {
	ZendThrowError(nil, "WeakReference objects do not support "+thing)
}
func ZendWeakrefNoWrite(object *Zval, member *Zval, value *Zval, rtc *any) *Zval {
	ZendWeakrefUnsupported("properties")
	return &(ExecutorGlobals.GetUninitializedZval())
}
func ZendWeakrefNoRead(object *Zval, member *Zval, type_ int, rtc *any, rv *Zval) *Zval {
	if ExecutorGlobals.GetException() == nil {
		ZendWeakrefUnsupported("properties")
	}
	return &(ExecutorGlobals.GetUninitializedZval())
}
func ZendWeakrefNoReadPtr(object *Zval, member *Zval, type_ int, rtc *any) *Zval {
	ZendWeakrefUnsupported("property references")
	return nil
}
func ZendWeakrefNoIsset(object *Zval, member *Zval, hse int, rtc *any) int {
	if hse != 2 {
		ZendWeakrefUnsupported("properties")
	}
	return 0
}
func ZendWeakrefNoUnset(object *Zval, member *Zval, rtc *any) { ZendWeakrefUnsupported("properties") }

var ZendWeakrefCreateArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{
		(*byte)(zend_uintptr_t(1)),
		ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0),
		0,
		0,
	},
	{"referent", ZEND_TYPE_ENCODE(IS_OBJECT, 0), 0, 0},
}
var ZendWeakrefGetArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(0)), ZEND_TYPE_ENCODE(IS_OBJECT, 1), 0, 0},
}

func zim_WeakReference___construct(execute_data *ZendExecuteData, return_value *Zval) {
	ZendThrowError(nil, "Direct instantiation of 'WeakReference' is not allowed, "+"use WeakReference::create instead")
}
func zim_WeakReference_create(execute_data *ZendExecuteData, return_value *Zval) {
	var referent *Zval
	for {
		var _flags int = ZEND_PARSE_PARAMS_THROW
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if UNEXPECTED(_num_args < _min_num_args) || UNEXPECTED(_num_args > _max_num_args) && EXPECTED(_max_num_args >= 0) {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if UNEXPECTED(ZendParseArgObject(_arg, &referent, nil, 0) == 0) {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if UNEXPECTED(_error_code != ZPP_ERROR_OK) {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if ZendWeakrefFind(referent, return_value) != 0 {
		return
	}
	ZendWeakrefCreate(referent, return_value)
}
func zim_WeakReference_get(execute_data *ZendExecuteData, return_value *Zval) {
	for {
		var _flags int = ZEND_PARSE_PARAMS_THROW
		var _min_num_args int = 0
		var _max_num_args int = 0
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if UNEXPECTED(_num_args < _min_num_args) || UNEXPECTED(_num_args > _max_num_args) && EXPECTED(_max_num_args >= 0) {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			break
		}
		if UNEXPECTED(_error_code != ZPP_ERROR_OK) {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	ZendWeakrefGet(getThis(), return_value)
}

var ZendWeakrefMethods []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_WeakReference___construct,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"create",
		zim_WeakReference_create,
		ZendWeakrefCreateArginfo,
		uint32_t(b.SizeOf("zend_weakref_create_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{
		"get",
		zim_WeakReference_get,
		ZendWeakrefGetArginfo,
		uint32_t(b.SizeOf("zend_weakref_get_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

func ZendRegisterWeakrefCe() {
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("WeakReference", b.SizeOf("\"WeakReference\"")-1, 1))
	ce.SetBuiltinFunctions(ZendWeakrefMethods)
	ZendCeWeakref = ZendRegisterInternalClass(&ce)
	ZendCeWeakref.SetCeFlags(ZendCeWeakref.GetCeFlags() | ZEND_ACC_FINAL)
	ZendCeWeakref.create_object = ZendWeakrefNew
	ZendCeWeakref.SetSerialize(ZendClassSerializeDeny)
	ZendCeWeakref.SetUnserialize(ZendClassUnserializeDeny)
	memcpy(&ZendWeakrefHandlers, ZendGetStdObjectHandlers(), b.SizeOf("zend_object_handlers"))
	ZendWeakrefHandlers.SetOffset(zend_long((*byte)(&((*ZendWeakref)(nil).GetStd())) - (*byte)(nil)))
	ZendWeakrefHandlers.SetFreeObj(ZendWeakrefFree)
	ZendWeakrefHandlers.SetReadProperty(ZendWeakrefNoRead)
	ZendWeakrefHandlers.SetWriteProperty(ZendWeakrefNoWrite)
	ZendWeakrefHandlers.SetHasProperty(ZendWeakrefNoIsset)
	ZendWeakrefHandlers.SetUnsetProperty(ZendWeakrefNoUnset)
	ZendWeakrefHandlers.SetGetPropertyPtrPtr(ZendWeakrefNoReadPtr)
	ZendWeakrefHandlers.SetCloneObj(nil)
}

/* }}} */

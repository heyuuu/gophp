// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// @type ZendWeakref struct

var ZendWeakrefHandlers ZendObjectHandlers

// #define zend_weakref_from(o) ( ( zend_weakref * ) ( ( ( char * ) o ) - XtOffsetOf ( zend_weakref , std ) ) )

// #define zend_weakref_fetch(z) zend_weakref_from ( Z_OBJ_P ( z ) )

func ZendWeakrefUnref(zv *Zval) {
	var wr *ZendWeakref = (*ZendWeakref)(zv.GetValue().GetPtr())
	wr.GetReferent().GetGc().SetTypeInfo(wr.GetReferent().GetGc().GetTypeInfo() &^ (1 << 7 << 0))
	wr.SetReferent(nil)
}
func ZendWeakrefsInit() {
	_zendHashInit(&EG.weakrefs, 8, ZendWeakrefUnref, 0)
}
func ZendWeakrefsNotify(object *ZendObject) {
	ZendHashIndexDel(&EG.weakrefs, ZendUlong(object))
}
func ZendWeakrefsShutdown() { ZendHashDestroy(&EG.weakrefs) }
func ZendWeakrefNew(ce *ZendClassEntry) *ZendObject {
	var wr *ZendWeakref = ZendObjectAlloc(g.SizeOf("zend_weakref"), ZendCeWeakref)
	ZendObjectStdInit(&wr.std, ZendCeWeakref)
	wr.GetStd().SetHandlers(&ZendWeakrefHandlers)
	return &wr.std
}
func ZendWeakrefFind(referent *Zval, return_value *Zval) ZendBool {
	var wr *ZendWeakref = ZendHashIndexFindPtr(&EG.weakrefs, zend_ulong(*referent).value.obj)
	if wr == nil {
		return 0
	}
	ZendGcAddref(&(&wr.std).GetGc())
	var __z *Zval = return_value
	__z.GetValue().SetObj(&wr.std)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	return 1
}
func ZendWeakrefCreate(referent *Zval, return_value *Zval) {
	var wr *ZendWeakref
	ObjectInitEx(return_value, ZendCeWeakref)
	wr = (*ZendWeakref)((*byte)(return_value.GetValue().GetObj()) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
	wr.SetReferent(referent.GetValue().GetObj())
	ZendHashIndexAddPtr(&EG.weakrefs, ZendUlong(wr.GetReferent()), wr)
	wr.GetReferent().GetGc().SetTypeInfo(wr.GetReferent().GetGc().GetTypeInfo() | 1<<7<<0)
}
func ZendWeakrefGet(weakref *Zval, return_value *Zval) {
	var wr *ZendWeakref = (*ZendWeakref)((*byte)(weakref.GetValue().GetObj()) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
	if wr.GetReferent() != nil {
		var __z *Zval = return_value
		__z.GetValue().SetObj(wr.GetReferent())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZvalAddrefP(return_value)
	}
}
func ZendWeakrefFree(zo *ZendObject) {
	var wr *ZendWeakref = (*ZendWeakref)((*byte)(zo) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
	if wr.GetReferent() != nil {
		ZendHashIndexDel(&EG.weakrefs, ZendUlong(wr.GetReferent()))
	}
	ZendObjectStdDtor(&wr.std)
}

// #define zend_weakref_unsupported(thing) zend_throw_error ( NULL , "WeakReference objects do not support " thing ) ;

func ZendWeakrefNoWrite(object *Zval, member *Zval, value *Zval, rtc *any) *Zval {
	ZendThrowError(nil, "WeakReference objects do not support "+"properties")
	return &EG.uninitialized_zval
}
func ZendWeakrefNoRead(object *Zval, member *Zval, type_ int, rtc *any, rv *Zval) *Zval {
	if EG.GetException() == nil {
		ZendThrowError(nil, "WeakReference objects do not support "+"properties")
	}
	return &EG.uninitialized_zval
}
func ZendWeakrefNoReadPtr(object *Zval, member *Zval, type_ int, rtc *any) *Zval {
	ZendThrowError(nil, "WeakReference objects do not support "+"property references")
	return nil
}
func ZendWeakrefNoIsset(object *Zval, member *Zval, hse int, rtc *any) int {
	if hse != 2 {
		ZendThrowError(nil, "WeakReference objects do not support "+"properties")
	}
	return 0
}
func ZendWeakrefNoUnset(object *Zval, member *Zval, rtc *any) {
	ZendThrowError(nil, "WeakReference objects do not support "+"properties")
}

var ZendWeakrefCreateArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), ZendType("WeakReference"), 0, 0},
	{"referent", 8<<2 | g.Cond(false, 0x1, 0x0), 0, 0},
}
var ZendWeakrefGetArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(0)), 8<<2 | g.Cond(true, 0x1, 0x0), 0, 0},
}

func zim_WeakReference___construct(execute_data *ZendExecuteData, return_value *Zval) {
	ZendThrowError(nil, "Direct instantiation of 'WeakReference' is not allowed, "+"use WeakReference::create instead")
}
func zim_WeakReference_create(execute_data *ZendExecuteData, return_value *Zval) {
	var referent *Zval
	for {
		var _flags int = 1 << 2
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgObject(_arg, &referent, nil, 0) == 0 {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		var _flags int = 1 << 2
		var _min_num_args int = 0
		var _max_num_args int = 0
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	ZendWeakrefGet(g.CondF1(&(execute_data.GetThis()).GetType() == 8, func() *Zval { return &(execute_data.GetThis()) }, nil), return_value)
}

var ZendWeakrefMethods []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_WeakReference___construct,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"create",
		zim_WeakReference_create,
		ZendWeakrefCreateArginfo,
		uint32(g.SizeOf("zend_weakref_create_arginfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<4,
	},
	{
		"get",
		zim_WeakReference_get,
		ZendWeakrefGetArginfo,
		uint32(g.SizeOf("zend_weakref_get_arginfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func ZendRegisterWeakrefCe() {
	var ce ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("WeakReference", g.SizeOf("\"WeakReference\"")-1, 1))
	ce.SetBuiltinFunctions(ZendWeakrefMethods)
	ZendCeWeakref = ZendRegisterInternalClass(&ce)
	ZendCeWeakref.SetCeFlags(ZendCeWeakref.GetCeFlags() | 1<<5)
	ZendCeWeakref.create_object = ZendWeakrefNew
	ZendCeWeakref.SetSerialize(ZendClassSerializeDeny)
	ZendCeWeakref.SetUnserialize(ZendClassUnserializeDeny)
	memcpy(&ZendWeakrefHandlers, &StdObjectHandlers, g.SizeOf("zend_object_handlers"))
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

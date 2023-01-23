// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_interfaces.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_INTERFACES_H

// # include "zend.h"

// # include "zend_API.h"

var ZendCeTraversable *ZendClassEntry
var ZendCeAggregate *ZendClassEntry
var ZendCeIterator *ZendClassEntry
var ZendCeArrayaccess *ZendClassEntry
var ZendCeSerializable *ZendClassEntry
var ZendCeCountable *ZendClassEntry

// #define zend_call_method_with_0_params(obj,obj_ce,fn_proxy,function_name,retval) zend_call_method ( obj , obj_ce , fn_proxy , function_name , sizeof ( function_name ) - 1 , retval , 0 , NULL , NULL )

// #define zend_call_method_with_1_params(obj,obj_ce,fn_proxy,function_name,retval,arg1) zend_call_method ( obj , obj_ce , fn_proxy , function_name , sizeof ( function_name ) - 1 , retval , 1 , arg1 , NULL )

// #define zend_call_method_with_2_params(obj,obj_ce,fn_proxy,function_name,retval,arg1,arg2) zend_call_method ( obj , obj_ce , fn_proxy , function_name , sizeof ( function_name ) - 1 , retval , 2 , arg1 , arg2 )

// #define REGISTER_MAGIC_INTERFACE(class_name,class_name_str) { zend_class_entry ce ; INIT_CLASS_ENTRY ( ce , # class_name_str , zend_funcs_ ## class_name ) zend_ce_ ## class_name = zend_register_internal_interface ( & ce ) ; zend_ce_ ## class_name -> interface_gets_implemented = zend_implement_ ## class_name ; }

// #define REGISTER_MAGIC_IMPLEMENT(class_name,interface_name) zend_class_implements ( zend_ce_ ## class_name , 1 , zend_ce_ ## interface_name )

// Source: <Zend/zend_interfaces.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

/* {{{ zend_call_method
Only returns the returned zval if retval_ptr != NULL */

func ZendCallMethod(object *Zval, obj_ce *ZendClassEntry, fn_proxy **ZendFunction, function_name string, function_name_len int, retval_ptr *Zval, param_count int, arg1 *Zval, arg2 *Zval) *Zval {
	var result int
	var fci ZendFcallInfo
	var retval Zval
	var params []Zval
	if param_count > 0 {
		var _z1 *Zval = &params[0]
		var _z2 *Zval = arg1
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	if param_count > 1 {
		var _z1 *Zval = &params[1]
		var _z2 *Zval = arg2
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	fci.SetSize(g.SizeOf("fci"))
	if object != nil {
		fci.SetObject(object.GetValue().GetObj())
	} else {
		fci.SetObject(nil)
	}
	if retval_ptr != nil {
		fci.SetRetval(retval_ptr)
	} else {
		fci.SetRetval(&retval)
	}
	fci.SetParamCount(param_count)
	fci.SetParams(params)
	fci.SetNoSeparation(1)
	if fn_proxy == nil && obj_ce == nil {

		/* no interest in caching and no information already present that is
		 * needed later inside zend_call_function. */

		var __z *Zval = &fci.function_name
		var __s *ZendString = ZendStringInit(function_name, function_name_len, 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		result = ZendCallFunction(&fci, nil)
		ZvalPtrDtor(&fci.function_name)
	} else {
		var fcic ZendFcallInfoCache
		&fci.function_name.u1.type_info = 0
		if obj_ce == nil {
			if object != nil {
				obj_ce = object.GetValue().GetObj().GetCe()
			} else {
				obj_ce = nil
			}
		}
		if fn_proxy == nil || (*fn_proxy) == nil {
			if obj_ce != nil {
				fcic.SetFunctionHandler(ZendHashStrFindPtr(&obj_ce.function_table, function_name, function_name_len))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					ZendErrorNoreturn(1<<4, "Couldn't find implementation for method %s::%s", obj_ce.GetName().GetVal(), function_name)

					/* error at c-level */

				}
			} else {
				fcic.SetFunctionHandler(ZendFetchFunctionStr(function_name, function_name_len))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					ZendErrorNoreturn(1<<4, "Couldn't find implementation for function %s", function_name)

					/* error at c-level */

				}
			}
			if fn_proxy != nil {
				*fn_proxy = fcic.GetFunctionHandler()
			}
		} else {
			fcic.SetFunctionHandler(*fn_proxy)
		}
		if object != nil {
			fcic.SetCalledScope(object.GetValue().GetObj().GetCe())
		} else {
			var called_scope *ZendClassEntry = ZendGetCalledScope(EG.GetCurrentExecuteData())
			if obj_ce != nil && (called_scope == nil || InstanceofFunction(called_scope, obj_ce) == 0) {
				fcic.SetCalledScope(obj_ce)
			} else {
				fcic.SetCalledScope(called_scope)
			}
		}
		if object != nil {
			fcic.SetObject(object.GetValue().GetObj())
		} else {
			fcic.SetObject(nil)
		}
		result = ZendCallFunction(&fci, &fcic)
	}
	if result == FAILURE {

		/* error at c-level */

		if obj_ce == nil {
			if object != nil {
				obj_ce = object.GetValue().GetObj().GetCe()
			} else {
				obj_ce = nil
			}
		}
		if EG.GetException() == nil {
			ZendErrorNoreturn(1<<4, "Couldn't execute method %s%s%s", g.CondF1(obj_ce != nil, func() []byte { return obj_ce.GetName().GetVal() }, ""), g.Cond(obj_ce != nil, "::", ""), function_name)
		}
	}
	if retval_ptr == nil {
		ZvalPtrDtor(&retval)
		return nil
	}
	return retval_ptr
}

/* }}} */

func ZendUserItNewIterator(ce *ZendClassEntry, object *Zval, retval *Zval) {
	ZendCallMethod(object, ce, &ce.iterator_funcs_ptr.GetZfNewIterator(), "getiterator", g.SizeOf("\"getiterator\"")-1, retval, 0, nil, nil)
}

/* }}} */

func ZendUserItInvalidateCurrent(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	if iter.GetValue().GetType() != 0 {
		ZvalPtrDtor(&iter.value)
		&iter.value.u1.type_info = 0
	}
}

/* }}} */

func ZendUserItDtor(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = &iter.it.GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZvalPtrDtor(object)
}

/* }}} */

func ZendUserItValid(_iter *ZendObjectIterator) int {
	if _iter != nil {
		var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
		var object *Zval = &iter.it.GetData()
		var more Zval
		var result int
		ZendCallMethod(object, iter.GetCe(), &iter.ce.GetIteratorFuncsPtr().GetZfValid(), "valid", g.SizeOf("\"valid\"")-1, &more, 0, nil, nil)
		result = IZendIsTrue(&more)
		ZvalPtrDtor(&more)
		if result != 0 {
			return SUCCESS
		} else {
			return FAILURE
		}
	}
	return FAILURE
}

/* }}} */

func ZendUserItGetCurrentData(_iter *ZendObjectIterator) *Zval {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = &iter.it.GetData()
	if iter.GetValue().GetType() == 0 {
		ZendCallMethod(object, iter.GetCe(), &iter.ce.GetIteratorFuncsPtr().GetZfCurrent(), "current", g.SizeOf("\"current\"")-1, &iter.value, 0, nil, nil)
	}
	return &iter.value
}

/* }}} */

func ZendUserItGetCurrentKey(_iter *ZendObjectIterator, key *Zval) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = &iter.it.GetData()
	var retval Zval
	ZendCallMethod(object, iter.GetCe(), &iter.ce.GetIteratorFuncsPtr().GetZfKey(), "key", g.SizeOf("\"key\"")-1, &retval, 0, nil, nil)
	if retval.GetType() != 0 {
		var __z *Zval = key
		var __zv *Zval = &retval
		if __zv.GetType() != 10 {
			var _z1 *Zval = __z
			var _z2 *Zval = __zv
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			var _z1 *Zval = __z
			var _z2 *Zval = &(*__zv).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
			ZvalPtrDtor(__zv)
		}
	} else {
		if EG.GetException() == nil {
			ZendError(1<<1, "Nothing returned from %s::key()", iter.GetCe().GetName().GetVal())
		}
		var __z *Zval = key
		__z.GetValue().SetLval(0)
		__z.SetTypeInfo(4)
	}
}

/* }}} */

func ZendUserItMoveForward(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = &iter.it.GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethod(object, iter.GetCe(), &iter.ce.GetIteratorFuncsPtr().GetZfNext(), "next", g.SizeOf("\"next\"")-1, nil, 0, nil, nil)
}

/* }}} */

func ZendUserItRewind(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = &iter.it.GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethod(object, iter.GetCe(), &iter.ce.GetIteratorFuncsPtr().GetZfRewind(), "rewind", g.SizeOf("\"rewind\"")-1, nil, 0, nil, nil)
}

/* }}} */

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = ZendObjectIteratorFuncs{ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent}

/* {{{ zend_user_it_get_iterator */

func ZendUserItGetIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendUserIterator
	if by_ref != 0 {
		ZendThrowError(nil, "An iterator cannot be used with foreach by reference")
		return nil
	}
	iterator = _emalloc(g.SizeOf("zend_user_iterator"))
	ZendIteratorInit((*ZendObjectIterator)(iterator))
	ZvalAddrefP(object)
	var __z *Zval = &iterator.it.GetData()
	__z.GetValue().SetObj(object.GetValue().GetObj())
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	iterator.GetIt().SetFuncs(&ZendInterfaceIteratorFuncsIterator)
	iterator.SetCe(object.GetValue().GetObj().GetCe())
	&iterator.value.u1.type_info = 0
	return (*ZendObjectIterator)(iterator)
}

/* }}} */

func ZendUserItGetNewIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator Zval
	var new_iterator *ZendObjectIterator
	var ce_it *ZendClassEntry
	ZendUserItNewIterator(ce, object, &iterator)
	if iterator.GetType() == 8 {
		ce_it = iterator.GetValue().GetObj().GetCe()
	} else {
		ce_it = nil
	}
	if ce_it == nil || ce_it.GetGetIterator() == nil || ce_it.GetGetIterator() == ZendUserItGetNewIterator && iterator.GetValue().GetObj() == object.GetValue().GetObj() {
		if EG.GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Objects returned by %s::getIterator() must be traversable or implement interface Iterator", g.CondF(ce != nil, func() []byte { return ce.GetName().GetVal() }, func() []byte { return object.GetValue().GetObj().GetCe().GetName().GetVal() }))
		}
		ZvalPtrDtor(&iterator)
		return nil
	}
	new_iterator = ce_it.GetGetIterator()(ce_it, &iterator, by_ref)
	ZvalPtrDtor(&iterator)
	return new_iterator
}

/* }}} */

func ZendImplementTraversable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	/* check that class_type is traversable at c-level or implements at least one of 'aggregate' and 'Iterator' */

	var i uint32
	if class_type.GetGetIterator() != nil || class_type.parent && class_type.parent.get_iterator {
		return SUCCESS
	}
	if class_type.GetNumInterfaces() != 0 {
		r.Assert((class_type.GetCeFlags() & 1 << 20) != 0)
		for i = 0; i < class_type.GetNumInterfaces(); i++ {
			if class_type.interfaces[i] == ZendCeAggregate || class_type.interfaces[i] == ZendCeIterator {
				return SUCCESS
			}
		}
	}
	ZendErrorNoreturn(1<<4, "Class %s must implement interface %s as part of either %s or %s", class_type.GetName().GetVal(), ZendCeTraversable.GetName().GetVal(), ZendCeIterator.GetName().GetVal(), ZendCeAggregate.GetName().GetVal())
	return FAILURE
}

/* }}} */

func ZendImplementAggregate(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	var i uint32
	var t int = -1
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil {
		if class_type.GetType() == 1 {

			/* inheritance ensures the class has necessary userland methods */

			return SUCCESS

			/* inheritance ensures the class has necessary userland methods */

		} else if class_type.GetGetIterator() != ZendUserItGetNewIterator {

			/* c-level get_iterator cannot be changed (exception being only Traversable is implemented) */

			if class_type.GetNumInterfaces() != 0 {
				r.Assert((class_type.GetCeFlags() & 1 << 20) != 0)
				for i = 0; i < class_type.GetNumInterfaces(); i++ {
					if class_type.interfaces[i] == ZendCeIterator {
						ZendErrorNoreturn(1<<0, "Class %s cannot implement both %s and %s at the same time", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeIterator.GetName().GetVal())
						return FAILURE
					}
					if class_type.interfaces[i] == ZendCeTraversable {
						t = i
					}
				}
			}
			if t == -1 {
				return FAILURE
			}
		}
	}
	if class_type.parent && (class_type.parent.ce_flags&1<<18) != 0 {
		class_type.SetGetIterator(class_type.parent.get_iterator)
		class_type.SetCeFlags(class_type.GetCeFlags() | 1<<18)
	} else {
		class_type.SetGetIterator(ZendUserItGetNewIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.GetType() == 1 {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, g.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		}
		funcs_ptr.SetZfNewIterator(ZendHashStrFindPtr(&class_type.function_table, "getiterator", g.SizeOf("\"getiterator\"")-1))
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
			memset(funcs_ptr, 0, g.SizeOf("zend_class_iterator_funcs"))
		} else {
			funcs_ptr.SetZfNewIterator(nil)
		}
	}
	return SUCCESS
}

/* }}} */

func ZendImplementIterator(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil && class_type.GetGetIterator() != ZendUserItGetIterator {
		if class_type.GetType() == 1 {

			/* inheritance ensures the class has the necessary userland methods */

			return SUCCESS

			/* inheritance ensures the class has the necessary userland methods */

		} else {

			/* c-level get_iterator cannot be changed */

			if class_type.GetGetIterator() == ZendUserItGetNewIterator {
				ZendErrorNoreturn(1<<0, "Class %s cannot implement both %s and %s at the same time", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeAggregate.GetName().GetVal())
			}
			return FAILURE
		}
	}
	if class_type.parent && (class_type.parent.ce_flags&1<<18) != 0 {
		class_type.SetGetIterator(class_type.parent.get_iterator)
		class_type.SetCeFlags(class_type.GetCeFlags() | 1<<18)
	} else {
		class_type.SetGetIterator(ZendUserItGetIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.GetType() == 1 {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, g.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		} else {
			funcs_ptr.SetZfRewind(ZendHashStrFindPtr(&class_type.function_table, "rewind", g.SizeOf("\"rewind\"")-1))
			funcs_ptr.SetZfValid(ZendHashStrFindPtr(&class_type.function_table, "valid", g.SizeOf("\"valid\"")-1))
			funcs_ptr.SetZfKey(ZendHashStrFindPtr(&class_type.function_table, "key", g.SizeOf("\"key\"")-1))
			funcs_ptr.SetZfCurrent(ZendHashStrFindPtr(&class_type.function_table, "current", g.SizeOf("\"current\"")-1))
			funcs_ptr.SetZfNext(ZendHashStrFindPtr(&class_type.function_table, "next", g.SizeOf("\"next\"")-1))
		}
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
			memset(funcs_ptr, 0, g.SizeOf("zend_class_iterator_funcs"))
		} else {
			funcs_ptr.SetZfValid(nil)
			funcs_ptr.SetZfCurrent(nil)
			funcs_ptr.SetZfKey(nil)
			funcs_ptr.SetZfNext(nil)
			funcs_ptr.SetZfRewind(nil)
		}
	}
	return SUCCESS
}

/* }}} */

func ZendImplementArrayaccess(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	return SUCCESS
}

/* }}}*/

func ZendUserSerialize(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var retval Zval
	var result int
	ZendCallMethod(object, ce, &ce.serialize_func, "serialize", g.SizeOf("\"serialize\"")-1, &retval, 0, nil, nil)
	if retval.GetType() == 0 || EG.GetException() != nil {
		result = FAILURE
	} else {
		switch retval.GetType() {
		case 1:

			/* we could also make this '*buf_len = 0' but this allows to skip variables */

			ZvalPtrDtor(&retval)
			return FAILURE
		case 6:
			*buffer = (*uint8)(_estrndup(retval.GetValue().GetStr().GetVal(), retval.GetValue().GetStr().GetLen()))
			*buf_len = retval.GetValue().GetStr().GetLen()
			result = SUCCESS
			break
		default:
			result = FAILURE
			break
		}
		ZvalPtrDtor(&retval)
	}
	if result == FAILURE && EG.GetException() == nil {
		ZendThrowExceptionEx(nil, 0, "%s::serialize() must return a string or NULL", ce.GetName().GetVal())
	}
	return result
}

/* }}} */

func ZendUserUnserialize(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	var zdata Zval
	if ObjectInitEx(object, ce) != SUCCESS {
		return FAILURE
	}
	var __z *Zval = &zdata
	var __s *ZendString = ZendStringInit((*byte)(buf), buf_len, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendCallMethod(object, ce, &ce.unserialize_func, "unserialize", g.SizeOf("\"unserialize\"")-1, nil, 1, &zdata, nil)
	ZvalPtrDtor(&zdata)
	if EG.GetException() != nil {
		return FAILURE
	} else {
		return SUCCESS
	}
}

/* }}} */

func ZendClassSerializeDeny(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	ZendThrowExceptionEx(nil, 0, "Serialization of '%s' is not allowed", ce.GetName().GetVal())
	return FAILURE
}

/* }}} */

func ZendClassUnserializeDeny(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	ZendThrowExceptionEx(nil, 0, "Unserialization of '%s' is not allowed", ce.GetName().GetVal())
	return FAILURE
}

/* }}} */

func ZendImplementSerializable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	if class_type.parent && (class_type.parent.serialize || class_type.parent.unserialize) && InstanceofFunctionEx(class_type.parent, ZendCeSerializable, 1) == 0 {
		return FAILURE
	}
	if class_type.GetSerialize() == nil {
		class_type.SetSerialize(ZendUserSerialize)
	}
	if class_type.GetUnserialize() == nil {
		class_type.SetUnserialize(ZendUserUnserialize)
	}
	return SUCCESS
}

/* }}}*/

func ZendImplementCountable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	return SUCCESS
}

/* }}}*/

var ZendFuncsAggregate []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"getIterator",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ZendFuncsIterator []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"current",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"next",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"key",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"valid",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"rewind",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ZendFuncsTraversable *ZendFunctionEntry = nil
var ArginfoArrayaccessOffset []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoArrayaccessOffsetGet []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoArrayaccessOffsetValue []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"offset", 0, 0, 0}, {"value", 0, 0, 0}}
var ZendFuncsArrayaccess []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"offsetExists",
		nil,
		ArginfoArrayaccessOffset,
		uint32(g.SizeOf("arginfo_arrayaccess_offset")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"offsetGet",
		nil,
		ArginfoArrayaccessOffsetGet,
		uint32(g.SizeOf("arginfo_arrayaccess_offset_get")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"offsetSet",
		nil,
		ArginfoArrayaccessOffsetValue,
		uint32(g.SizeOf("arginfo_arrayaccess_offset_value")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"offsetUnset",
		nil,
		ArginfoArrayaccessOffset,
		uint32(g.SizeOf("arginfo_arrayaccess_offset")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoSerializableSerialize []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"serialized", 0, 0, 0}}
var ZendFuncsSerializable []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"serialize",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"unserialize",
		nil,
		ArginfoSerializableSerialize,
		uint32(g.SizeOf("arginfo_serializable_serialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCountableCount []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ZendFuncsCountable []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"count",
		nil,
		ArginfoCountableCount,
		uint32(g.SizeOf("arginfo_countable_count")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZendRegisterInterfaces() {
	var ce zend_class_entry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Traversable", g.SizeOf("\"Traversable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsTraversable)
	ZendCeTraversable = ZendRegisterInternalInterface(&ce)
	ZendCeTraversable.interface_gets_implemented = ZendImplementTraversable
	var ce zend_class_entry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("IteratorAggregate", g.SizeOf("\"IteratorAggregate\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsAggregate)
	ZendCeAggregate = ZendRegisterInternalInterface(&ce)
	ZendCeAggregate.interface_gets_implemented = ZendImplementAggregate
	ZendClassImplements(ZendCeAggregate, 1, ZendCeTraversable)
	var ce zend_class_entry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Iterator", g.SizeOf("\"Iterator\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsIterator)
	ZendCeIterator = ZendRegisterInternalInterface(&ce)
	ZendCeIterator.interface_gets_implemented = ZendImplementIterator
	ZendClassImplements(ZendCeIterator, 1, ZendCeTraversable)
	var ce zend_class_entry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArrayAccess", g.SizeOf("\"ArrayAccess\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsArrayaccess)
	ZendCeArrayaccess = ZendRegisterInternalInterface(&ce)
	ZendCeArrayaccess.interface_gets_implemented = ZendImplementArrayaccess
	var ce zend_class_entry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Serializable", g.SizeOf("\"Serializable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsSerializable)
	ZendCeSerializable = ZendRegisterInternalInterface(&ce)
	ZendCeSerializable.interface_gets_implemented = ZendImplementSerializable
	var ce ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Countable", g.SizeOf("\"Countable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsCountable)
	ZendCeCountable = ZendRegisterInternalInterface(&ce)
	ZendCeCountable.interface_gets_implemented = ZendImplementCountable
}

/* }}} */

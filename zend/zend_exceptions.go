// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_exceptions.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Marcus Boerger <helly@php.net>                              |
   |          Sterling Hughes <sterling@php.net>                          |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_EXCEPTIONS_H

var ZendCeThrowable *ZendClassEntry
var ZendCeException *ZendClassEntry
var ZendCeErrorException *ZendClassEntry
var ZendCeError *ZendClassEntry
var ZendCeCompileError *ZendClassEntry
var ZendCeParseError *ZendClassEntry
var ZendCeTypeError *ZendClassEntry
var ZendCeArgumentCountError *ZendClassEntry
var ZendCeArithmeticError *ZendClassEntry
var ZendCeDivisionByZeroError *ZendClassEntry

/* Deprecated - Use zend_ce_exception directly instead */

/* Deprecated - Use zend_ce_error_exception directly instead */

/* exception_ce   NULL, zend_ce_exception, zend_ce_error, or a derived class
 * message        NULL or the message of the exception */

var ZendThrowExceptionHook func(ex *Zval)

/* show an exception using zend_error(severity,...), severity should be E_ERROR */

// # include "zend_globals.h"

func ZendRethrowException(execute_data *ZendExecuteData) {
	if execute_data.GetOpline().GetOpcode() != 149 {
		EG.SetOplineBeforeException(execute_data.GetOpline())
		execute_data.SetOpline(EG.GetExceptionOp())
	}
}

// Source: <Zend/zend_exceptions.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Marcus Boerger <helly@php.net>                              |
   |          Sterling Hughes <sterling@php.net>                          |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_builtin_functions.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "zend_vm.h"

// # include "zend_dtrace.h"

// # include "zend_smart_str.h"

var DefaultExceptionHandlers ZendObjectHandlers

/* {{{ zend_implement_throwable */

func ZendImplementThrowable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	if InstanceofFunction(class_type, ZendCeException) != 0 || InstanceofFunction(class_type, ZendCeError) != 0 {
		return SUCCESS
	}
	ZendErrorNoreturn(1<<0, "Class %s cannot implement interface %s, extend %s or %s instead", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeException.GetName().GetVal(), ZendCeError.GetName().GetVal())
	return FAILURE
}

/* }}} */

func IGetExceptionBase(object *Zval) *ZendClassEntry {
	if InstanceofFunction(object.GetValue().GetObj().GetCe(), ZendCeException) != 0 {
		return ZendCeException
	} else {
		return ZendCeError
	}
}

/* }}} */

func ZendGetExceptionBase(object *Zval) *ZendClassEntry { return IGetExceptionBase(object) }

/* }}} */

func ZendExceptionSetPrevious(exception *ZendObject, add_previous *ZendObject) {
	var previous *Zval
	var ancestor *Zval
	var ex *Zval
	var pv Zval
	var zv Zval
	var rv Zval
	var base_ce *ZendClassEntry
	if exception == nil || add_previous == nil {
		return
	}
	if exception == add_previous {
		ZendObjectRelease(add_previous)
		return
	}
	var __z *Zval = &pv
	__z.GetValue().SetObj(add_previous)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	if InstanceofFunction(pv.GetValue().GetObj().GetCe(), ZendCeThrowable) == 0 {
		ZendErrorNoreturn(1<<4, "Previous exception must implement Throwable")
		return
	}
	var __z *Zval = &zv
	__z.GetValue().SetObj(exception)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ex = &zv
	for {
		ancestor = ZendReadPropertyEx(IGetExceptionBase(&pv), &pv, ZendKnownStrings[ZEND_STR_PREVIOUS], 1, &rv)
		for ancestor.GetType() == 8 {
			if ancestor.GetValue().GetObj() == ex.GetValue().GetObj() {
				ZendObjectRelease(add_previous)
				return
			}
			ancestor = ZendReadPropertyEx(IGetExceptionBase(ancestor), ancestor, ZendKnownStrings[ZEND_STR_PREVIOUS], 1, &rv)
		}
		base_ce = IGetExceptionBase(ex)
		previous = ZendReadPropertyEx(base_ce, ex, ZendKnownStrings[ZEND_STR_PREVIOUS], 1, &rv)
		if previous.GetType() == 1 {
			ZendUpdatePropertyEx(base_ce, ex, ZendKnownStrings[ZEND_STR_PREVIOUS], &pv)
			ZendGcDelref(&add_previous.gc)
			return
		}
		ex = previous
		if ex.GetValue().GetObj() == add_previous {
			break
		}
	}
}

/* }}} */

func ZendExceptionSave() {
	if EG.GetPrevException() != nil {
		ZendExceptionSetPrevious(EG.GetException(), EG.GetPrevException())
	}
	if EG.GetException() != nil {
		EG.SetPrevException(EG.GetException())
	}
	EG.SetException(nil)
}

/* }}} */

func ZendExceptionRestore() {
	if EG.GetPrevException() != nil {
		if EG.GetException() != nil {
			ZendExceptionSetPrevious(EG.GetException(), EG.GetPrevException())
		} else {
			EG.SetException(EG.GetPrevException())
		}
		EG.SetPrevException(nil)
	}
}

/* }}} */

func ZendThrowExceptionInternal(exception *Zval) {
	if exception != nil {
		var previous *ZendObject = EG.GetException()
		ZendExceptionSetPrevious(exception.GetValue().GetObj(), EG.GetException())
		EG.SetException(exception.GetValue().GetObj())
		if previous != nil {
			return
		}
	}
	if EG.GetCurrentExecuteData() == nil {
		if exception != nil && (exception.GetValue().GetObj().GetCe() == ZendCeParseError || exception.GetValue().GetObj().GetCe() == ZendCeCompileError) {
			return
		}
		if EG.GetException() != nil {
			ZendExceptionError(EG.GetException(), 1<<0)
		}
		ZendErrorNoreturn(1<<4, "Exception thrown without a stack frame")
	}
	if ZendThrowExceptionHook != nil {
		ZendThrowExceptionHook(exception)
	}
	if EG.GetCurrentExecuteData().GetFunc() == nil || (EG.GetCurrentExecuteData().GetFunc().GetCommonType()&1) != 0 || EG.GetCurrentExecuteData().GetOpline().GetOpcode() == 149 {

		/* no need to rethrow the exception */

		return

		/* no need to rethrow the exception */

	}
	EG.SetOplineBeforeException(EG.GetCurrentExecuteData().GetOpline())
	EG.GetCurrentExecuteData().SetOpline(EG.GetExceptionOp())
}

/* }}} */

func ZendClearException() {
	var exception *ZendObject
	if EG.GetPrevException() != nil {
		ZendObjectRelease(EG.GetPrevException())
		EG.SetPrevException(nil)
	}
	if EG.GetException() == nil {
		return
	}

	/* exception may have destructor */

	exception = EG.GetException()
	EG.SetException(nil)
	ZendObjectRelease(exception)
	if EG.GetCurrentExecuteData() != nil {
		EG.GetCurrentExecuteData().SetOpline(EG.GetOplineBeforeException())
	}
}

/* }}} */

func ZendDefaultExceptionNewEx(class_type *ZendClassEntry, skip_top_traces int) *ZendObject {
	var obj Zval
	var tmp Zval
	var object *ZendObject
	var trace Zval
	var base_ce *ZendClassEntry
	var filename *ZendString
	object = ZendObjectsNew(class_type)
	obj.GetValue().SetObj(object)
	obj.GetValue().GetObj().SetHandlers(&DefaultExceptionHandlers)
	ObjectPropertiesInit(object, class_type)
	if EG.GetCurrentExecuteData() != nil {
		ZendFetchDebugBacktrace(&trace, skip_top_traces, g.Cond(EG.GetExceptionIgnoreArgs() != 0, 1<<1, 0), 0)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = &trace
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	}
	ZvalSetRefcountP(&trace, 0)
	base_ce = IGetExceptionBase(&obj)
	if class_type != ZendCeParseError && class_type != ZendCeCompileError || !(g.Assign(&filename, ZendGetCompiledFilename())) {
		var _s *byte = ZendGetExecutedFilename()
		var __z *Zval = &tmp
		var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		ZendUpdatePropertyEx(base_ce, &obj, ZendKnownStrings[ZEND_STR_FILE], &tmp)
		ZvalPtrDtor(&tmp)
		var __z *Zval = &tmp
		__z.GetValue().SetLval(ZendGetExecutedLineno())
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(base_ce, &obj, ZendKnownStrings[ZEND_STR_LINE], &tmp)
	} else {
		var __z *Zval = &tmp
		var __s *ZendString = filename
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		ZendUpdatePropertyEx(base_ce, &obj, ZendKnownStrings[ZEND_STR_FILE], &tmp)
		var __z *Zval = &tmp
		__z.GetValue().SetLval(ZendGetCompiledLineno())
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(base_ce, &obj, ZendKnownStrings[ZEND_STR_LINE], &tmp)
	}
	ZendUpdatePropertyEx(base_ce, &obj, ZendKnownStrings[ZEND_STR_TRACE], &trace)
	return object
}

/* }}} */

func ZendDefaultExceptionNew(class_type *ZendClassEntry) *ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 0)
}

/* }}} */

func ZendErrorExceptionNew(class_type *ZendClassEntry) *ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 2)
}

/* }}} */

func ZimExceptionClone(execute_data *ZendExecuteData, return_value *Zval) {
	/* Should never be executable */

	ZendThrowException(nil, "Cannot clone object using __clone()", 0)

	/* Should never be executable */
}

/* }}} */

func ZimExceptionConstruct(execute_data *ZendExecuteData, return_value *Zval) {
	var message *ZendString = nil
	var code ZendLong = 0
	var tmp Zval
	var object *Zval
	var previous *Zval = nil
	var base_ce *ZendClassEntry
	var argc int = execute_data.GetThis().GetNumArgs()
	object = &(execute_data.GetThis())
	base_ce = IGetExceptionBase(object)
	if ZendParseParametersEx(1<<1, argc, "|SlO!", &message, &code, &previous, ZendCeThrowable) == FAILURE {
		var ce *ZendClassEntry
		if execute_data.GetThis().GetType() == 8 {
			ce = execute_data.GetThis().GetValue().GetObj().GetCe()
		} else if execute_data.GetThis().GetValue().GetCe() != nil {
			ce = execute_data.GetThis().GetValue().GetCe()
		} else {
			ce = base_ce
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code [, Throwable $previous = NULL]]])", ce.GetName().GetVal())
		return
	}
	if message != nil {
		var __z *Zval = &tmp
		var __s *ZendString = message
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		ZendUpdatePropertyEx(base_ce, object, ZendKnownStrings[ZEND_STR_MESSAGE], &tmp)
	}
	if code != 0 {
		var __z *Zval = &tmp
		__z.GetValue().SetLval(code)
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(base_ce, object, ZendKnownStrings[ZEND_STR_CODE], &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(base_ce, object, ZendKnownStrings[ZEND_STR_PREVIOUS], previous)
	}
}

/* }}} */

// #define CHECK_EXC_TYPE(id,type) pvalue = zend_read_property_ex ( i_get_exception_base ( object ) , ( object ) , ZSTR_KNOWN ( id ) , 1 , & value ) ; if ( Z_TYPE_P ( pvalue ) != IS_NULL && Z_TYPE_P ( pvalue ) != type ) { zend_unset_property ( i_get_exception_base ( object ) , object , ZSTR_VAL ( ZSTR_KNOWN ( id ) ) , ZSTR_LEN ( ZSTR_KNOWN ( id ) ) ) ; }

func ZimExceptionWakeup(execute_data *ZendExecuteData, return_value *Zval) {
	var value Zval
	var pvalue *Zval
	var object *Zval = &(execute_data.GetThis())
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_MESSAGE], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 6 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_MESSAGE].GetVal(), ZendKnownStrings[ZEND_STR_MESSAGE].GetLen())
	}
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_STRING], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 6 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_STRING].GetVal(), ZendKnownStrings[ZEND_STR_STRING].GetLen())
	}
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_CODE], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 4 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_CODE].GetVal(), ZendKnownStrings[ZEND_STR_CODE].GetLen())
	}
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_FILE], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 6 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_FILE].GetVal(), ZendKnownStrings[ZEND_STR_FILE].GetLen())
	}
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_LINE], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 4 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_LINE].GetVal(), ZendKnownStrings[ZEND_STR_LINE].GetLen())
	}
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_TRACE], 1, &value)
	if pvalue.GetType() != 1 && pvalue.GetType() != 7 {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZendKnownStrings[ZEND_STR_TRACE].GetVal(), ZendKnownStrings[ZEND_STR_TRACE].GetLen())
	}
	pvalue = ZendReadProperty(IGetExceptionBase(object), object, "previous", g.SizeOf("\"previous\"")-1, 1, &value)
	if pvalue != nil && pvalue.GetType() != 1 && (pvalue.GetType() != 8 || InstanceofFunction(pvalue.GetValue().GetObj().GetCe(), ZendCeThrowable) == 0 || pvalue == object) {
		ZendUnsetProperty(IGetExceptionBase(object), object, "previous", g.SizeOf("\"previous\"")-1)
	}
}

/* }}} */

func ZimErrorExceptionConstruct(execute_data *ZendExecuteData, return_value *Zval) {
	var message *ZendString = nil
	var filename *ZendString = nil
	var code ZendLong = 0
	var severity ZendLong = 1 << 0
	var lineno ZendLong
	var tmp Zval
	var object *Zval
	var previous *Zval = nil
	var argc int = execute_data.GetThis().GetNumArgs()
	if ZendParseParametersEx(1<<1, argc, "|SllSlO!", &message, &code, &severity, &filename, &lineno, &previous, ZendCeThrowable) == FAILURE {
		var ce *ZendClassEntry
		if execute_data.GetThis().GetType() == 8 {
			ce = execute_data.GetThis().GetValue().GetObj().GetCe()
		} else if execute_data.GetThis().GetValue().GetCe() != nil {
			ce = execute_data.GetThis().GetValue().GetCe()
		} else {
			ce = ZendCeErrorException
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code, [ long $severity, [ string $filename, [ long $lineno  [, Throwable $previous = NULL]]]]]])", ce.GetName().GetVal())
		return
	}
	object = &(execute_data.GetThis())
	if message != nil {
		var __z *Zval = &tmp
		var __s *ZendString = message
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_MESSAGE], &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		var __z *Zval = &tmp
		__z.GetValue().SetLval(code)
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_CODE], &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_PREVIOUS], previous)
	}
	var __z *Zval = &tmp
	__z.GetValue().SetLval(severity)
	__z.SetTypeInfo(4)
	ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_SEVERITY], &tmp)
	if argc >= 4 {
		var __z *Zval = &tmp
		var __s *ZendString = filename
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_FILE], &tmp)
		ZvalPtrDtor(&tmp)
		if argc < 5 {
			lineno = 0
		}
		var __z *Zval = &tmp
		__z.GetValue().SetLval(lineno)
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(ZendCeException, object, ZendKnownStrings[ZEND_STR_LINE], &tmp)
	}
}

/* }}} */

// #define DEFAULT_0_PARAMS       if ( zend_parse_parameters_none ( ) == FAILURE ) { return ; }

// #define GET_PROPERTY(object,id) zend_read_property_ex ( i_get_exception_base ( object ) , ( object ) , ZSTR_KNOWN ( id ) , 0 , & rv )

// #define GET_PROPERTY_SILENT(object,id) zend_read_property_ex ( i_get_exception_base ( object ) , ( object ) , ZSTR_KNOWN ( id ) , 1 , & rv )

/* {{{ proto string Exception|Error::getFile()
   Get the file in which the exception occurred */

func zim_exception_getFile(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_FILE], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_exception_getLine(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_LINE], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_exception_getMessage(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_MESSAGE], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_exception_getCode(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_CODE], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_exception_getTrace(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_TRACE], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_error_exception_getSeverity(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	prop = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_SEVERITY], 0, &rv)
	if prop.GetType() == 10 {
		prop = &(*prop).value.GetRef().GetVal()
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = prop
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

// #define TRACE_APPEND_KEY(key) do { tmp = zend_hash_find ( ht , key ) ; if ( tmp ) { if ( Z_TYPE_P ( tmp ) != IS_STRING ) { zend_error ( E_WARNING , "Value for %s is no string" , ZSTR_VAL ( key ) ) ; smart_str_appends ( str , "[unknown]" ) ; } else { smart_str_appends ( str , Z_STRVAL_P ( tmp ) ) ; } } } while ( 0 )

func _buildTraceArgs(arg *Zval, str *SmartStr) {
	/* the trivial way would be to do
	 * convert_to_string_ex(arg);
	 * append it and kill the now tmp arg.
	 * but that could cause some E_NOTICE and also damn long lines.
	 */

	if arg.GetType() == 10 {
		arg = &(*arg).value.GetRef().GetVal()
	}
	switch arg.GetType() {
	case 1:
		SmartStrAppendlEx(str, "NULL, ", strlen("NULL, "), 0)
		break
	case 6:
		SmartStrAppendcEx(str, '\'', 0)
		SmartStrAppendEscaped(str, arg.GetValue().GetStr().GetVal(), g.CondF1(arg.GetValue().GetStr().GetLen() < 15, func() int { return arg.GetValue().GetStr().GetLen() }, 15))
		if arg.GetValue().GetStr().GetLen() > 15 {
			SmartStrAppendlEx(str, "...', ", strlen("...', "), 0)
		} else {
			SmartStrAppendlEx(str, "', ", strlen("', "), 0)
		}
		break
	case 2:
		SmartStrAppendlEx(str, "false, ", strlen("false, "), 0)
		break
	case 3:
		SmartStrAppendlEx(str, "true, ", strlen("true, "), 0)
		break
	case 9:
		SmartStrAppendlEx(str, "Resource id #", strlen("Resource id #"), 0)
		SmartStrAppendLongEx(str, arg.GetValue().GetRes().GetHandle(), 0)
		SmartStrAppendlEx(str, ", ", strlen(", "), 0)
		break
	case 4:
		SmartStrAppendLongEx(str, arg.GetValue().GetLval(), 0)
		SmartStrAppendlEx(str, ", ", strlen(", "), 0)
		break
	case 5:
		SmartStrAppendPrintf(str, "%.*G", int(EG.GetPrecision()), arg.GetValue().GetDval())
		SmartStrAppendlEx(str, ", ", strlen(", "), 0)
		break
	case 7:
		SmartStrAppendlEx(str, "Array, ", strlen("Array, "), 0)
		break
	case 8:
		var class_name *ZendString = arg.GetValue().GetObj().GetHandlers().GetGetClassName()(arg.GetValue().GetObj())
		SmartStrAppendlEx(str, "Object(", strlen("Object("), 0)
		SmartStrAppendlEx(str, class_name.GetVal(), strlen(class_name.GetVal()), 0)
		SmartStrAppendlEx(str, "), ", strlen("), "), 0)
		ZendStringReleaseEx(class_name, 0)
		break
	}
}

/* }}} */

func _buildTraceString(str *SmartStr, ht *HashTable, num uint32) {
	var file *Zval
	var tmp *Zval
	SmartStrAppendcEx(str, '#', 0)
	SmartStrAppendLongEx(str, num, 0)
	SmartStrAppendcEx(str, ' ', 0)
	file = ZendHashFindEx(ht, ZendKnownStrings[ZEND_STR_FILE], 1)
	if file != nil {
		if file.GetType() != 6 {
			ZendError(1<<1, "Function name is no string")
			SmartStrAppendlEx(str, "[unknown function]", strlen("[unknown function]"), 0)
		} else {
			var line ZendLong
			tmp = ZendHashFindEx(ht, ZendKnownStrings[ZEND_STR_LINE], 1)
			if tmp != nil {
				if tmp.GetType() == 4 {
					line = tmp.GetValue().GetLval()
				} else {
					ZendError(1<<1, "Line is no long")
					line = 0
				}
			} else {
				line = 0
			}
			SmartStrAppendEx(str, file.GetValue().GetStr(), 0)
			SmartStrAppendcEx(str, '(', 0)
			SmartStrAppendLongEx(str, line, 0)
			SmartStrAppendlEx(str, "): ", strlen("): "), 0)
		}
	} else {
		SmartStrAppendlEx(str, "[internal function]: ", strlen("[internal function]: "), 0)
	}
	tmp = ZendHashFind(ht, ZendKnownStrings[ZEND_STR_CLASS])
	if tmp != nil {
		if tmp.GetType() != 6 {
			ZendError(1<<1, "Value for %s is no string", ZendKnownStrings[ZEND_STR_CLASS].GetVal())
			SmartStrAppendlEx(str, "[unknown]", strlen("[unknown]"), 0)
		} else {
			SmartStrAppendlEx(str, tmp.GetValue().GetStr().GetVal(), strlen(tmp.GetValue().GetStr().GetVal()), 0)
		}
	}
	tmp = ZendHashFind(ht, ZendKnownStrings[ZEND_STR_TYPE])
	if tmp != nil {
		if tmp.GetType() != 6 {
			ZendError(1<<1, "Value for %s is no string", ZendKnownStrings[ZEND_STR_TYPE].GetVal())
			SmartStrAppendlEx(str, "[unknown]", strlen("[unknown]"), 0)
		} else {
			SmartStrAppendlEx(str, tmp.GetValue().GetStr().GetVal(), strlen(tmp.GetValue().GetStr().GetVal()), 0)
		}
	}
	tmp = ZendHashFind(ht, ZendKnownStrings[ZEND_STR_FUNCTION])
	if tmp != nil {
		if tmp.GetType() != 6 {
			ZendError(1<<1, "Value for %s is no string", ZendKnownStrings[ZEND_STR_FUNCTION].GetVal())
			SmartStrAppendlEx(str, "[unknown]", strlen("[unknown]"), 0)
		} else {
			SmartStrAppendlEx(str, tmp.GetValue().GetStr().GetVal(), strlen(tmp.GetValue().GetStr().GetVal()), 0)
		}
	}
	SmartStrAppendcEx(str, '(', 0)
	tmp = ZendHashFindEx(ht, ZendKnownStrings[ZEND_STR_ARGS], 1)
	if tmp != nil {
		if tmp.GetType() == 7 {
			var last_len int = str.GetS().GetLen()
			var arg *Zval
			for {
				var __ht *HashTable = tmp.GetValue().GetArr()
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					arg = _z
					_buildTraceArgs(arg, str)
				}
				break
			}
			if last_len != str.GetS().GetLen() {
				str.GetS().SetLen(str.GetS().GetLen() - 2)
			}
		} else {
			ZendError(1<<1, "args element is no array")
		}
	}
	SmartStrAppendlEx(str, ")\n", strlen(")\n"), 0)
}

/* }}} */

func zim_exception_getTraceAsString(execute_data *ZendExecuteData, return_value *Zval) {
	var trace *Zval
	var frame *Zval
	var rv Zval
	var index ZendUlong
	var object *Zval
	var base_ce *ZendClassEntry
	var str SmartStr = SmartStr{0}
	var num uint32 = 0
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	object = &(execute_data.GetThis())
	base_ce = IGetExceptionBase(object)
	trace = ZendReadPropertyEx(base_ce, object, ZendKnownStrings[ZEND_STR_TRACE], 1, &rv)
	if trace.GetType() != 7 {
		return_value.SetTypeInfo(2)
		return
	}
	for {
		var __ht *HashTable = trace.GetValue().GetArr()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			index = _p.GetH()
			frame = _z
			if frame.GetType() != 7 {
				ZendError(1<<1, "Expected array for frame "+"%"+"llu", index)
				continue
			}
			_buildTraceString(&str, frame.GetValue().GetArr(), g.PostInc(&num))
		}
		break
	}
	SmartStrAppendcEx(&str, '#', 0)
	SmartStrAppendLongEx(&str, num, 0)
	SmartStrAppendlEx(&str, " {main}", strlen(" {main}"), 0)
	SmartStr0(&str)
	var __z *Zval = return_value
	var __s *ZendString = str.GetS()
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return
}

/* }}} */

func zim_exception_getPrevious(execute_data *ZendExecuteData, return_value *Zval) {
	var rv Zval
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var _z1 *Zval = return_value
	var _z2 *Zval = ZendReadPropertyEx(IGetExceptionBase(&(execute_data.GetThis())), &(execute_data.GetThis()), ZendKnownStrings[ZEND_STR_PREVIOUS], 1, &rv)
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* {{{ proto string Exception|Error::__toString()
   Obtain the string representation of the Exception object */

func zim_exception___toString(execute_data *ZendExecuteData, return_value *Zval) {
	var trace Zval
	var exception *Zval
	var base_ce *ZendClassEntry
	var str *ZendString
	var fci ZendFcallInfo
	var rv Zval
	var tmp Zval
	var fname *ZendString
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	str = ZendEmptyString
	exception = &(execute_data.GetThis())
	fname = ZendStringInit("gettraceasstring", g.SizeOf("\"gettraceasstring\"")-1, 0)
	for exception != nil && exception.GetType() == 8 && InstanceofFunction(exception.GetValue().GetObj().GetCe(), ZendCeThrowable) != 0 {
		var prev_str *ZendString = str
		var message *ZendString = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(exception), exception, ZendKnownStrings[ZEND_STR_MESSAGE], 0, &rv))
		var file *ZendString = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(exception), exception, ZendKnownStrings[ZEND_STR_FILE], 0, &rv))
		var line ZendLong = ZvalGetLong(ZendReadPropertyEx(IGetExceptionBase(exception), exception, ZendKnownStrings[ZEND_STR_LINE], 0, &rv))
		fci.SetSize(g.SizeOf("fci"))
		var __z *Zval = &fci.function_name
		var __s *ZendString = fname
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		fci.SetObject(exception.GetValue().GetObj())
		fci.SetRetval(&trace)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		ZendCallFunction(&fci, nil)
		if trace.GetType() != 6 {
			ZvalPtrDtor(&trace)
			&trace.SetTypeInfo(0)
		}
		if (exception.GetValue().GetObj().GetCe() == ZendCeTypeError || exception.GetValue().GetObj().GetCe() == ZendCeArgumentCountError) && strstr(message.GetVal(), ", called in ") {
			var real_message *ZendString = ZendStrpprintf(0, "%s and defined", message.GetVal())
			ZendStringReleaseEx(message, 0)
			message = real_message
		}
		if message.GetLen() > 0 {
			str = ZendStrpprintf(0, "%s: %s in %s:"+"%"+"lld"+"\nStack trace:\n%s%s%s", exception.GetValue().GetObj().GetCe().GetName().GetVal(), message.GetVal(), file.GetVal(), line, g.CondF1(trace.GetType() == 6 && trace.GetValue().GetStr().GetLen() != 0, func() []byte { return trace.GetValue().GetStr().GetVal() }, "#0 {main}\n"), g.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		} else {
			str = ZendStrpprintf(0, "%s in %s:"+"%"+"lld"+"\nStack trace:\n%s%s%s", exception.GetValue().GetObj().GetCe().GetName().GetVal(), file.GetVal(), line, g.CondF1(trace.GetType() == 6 && trace.GetValue().GetStr().GetLen() != 0, func() []byte { return trace.GetValue().GetStr().GetVal() }, "#0 {main}\n"), g.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		}
		ZendStringReleaseEx(prev_str, 0)
		ZendStringReleaseEx(message, 0)
		ZendStringReleaseEx(file, 0)
		ZvalPtrDtor(&trace)
		exception.GetValue().GetCounted().GetGc().SetTypeInfo(exception.GetValue().GetCounted().GetGc().GetTypeInfo() | 1<<5<<0)
		exception = ZendReadPropertyEx(IGetExceptionBase(exception), exception, ZendKnownStrings[ZEND_STR_PREVIOUS], 0, &rv)
		if exception != nil && exception.GetType() == 8 && (ZvalGcFlags(exception.GetValue().GetCounted().GetGc().GetTypeInfo())&1<<5) != 0 {
			break
		}
	}
	ZendStringReleaseEx(fname, 0)
	exception = &(execute_data.GetThis())

	/* Reset apply counts */

	for exception != nil && exception.GetType() == 8 && g.Assign(&base_ce, IGetExceptionBase(exception)) && InstanceofFunction(exception.GetValue().GetObj().GetCe(), base_ce) != 0 {
		if (ZvalGcFlags(exception.GetValue().GetCounted().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
			exception.GetValue().GetCounted().GetGc().SetTypeInfo(exception.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		} else {
			break
		}
		exception = ZendReadPropertyEx(IGetExceptionBase(exception), exception, ZendKnownStrings[ZEND_STR_PREVIOUS], 0, &rv)
	}
	exception = &(execute_data.GetThis())
	base_ce = IGetExceptionBase(exception)

	/* We store the result in the private property string so we can access
	 * the result in uncaught exception handlers without memleaks. */

	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZendUpdatePropertyEx(base_ce, exception, ZendKnownStrings[ZEND_STR_STRING], &tmp)
	var __z *Zval = return_value
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return
}

/* }}} */

var ZendFuncsThrowable []ZendFunctionEntry = []ZendFunctionEntry{{"getMessage", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getCode", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getFile", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getLine", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getTrace", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getPrevious", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"getTraceAsString", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {"__toString", nil, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<6}, {nil, nil, nil, 0, 0}}

/* }}} */

var ArginfoExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"message", 0, 0, 0}, {"code", 0, 0, 0}, {"previous", 0, 0, 0}}
var DefaultExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{{"__clone", ZimExceptionClone, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<2 | 1<<5}, {"__construct", ZimExceptionConstruct, ArginfoExceptionConstruct, uint32(g.SizeOf("arginfo_exception___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 0}, {"__wakeup", ZimExceptionWakeup, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 0}, {"getMessage", zim_exception_getMessage, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getCode", zim_exception_getCode, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getFile", zim_exception_getFile, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getLine", zim_exception_getLine, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getTrace", zim_exception_getTrace, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getPrevious", zim_exception_getPrevious, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"getTraceAsString", zim_exception_getTraceAsString, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {"__toString", zim_exception___toString, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 0}, {nil, nil, nil, 0, 0}}
var ArginfoErrorExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"message", 0, 0, 0}, {"code", 0, 0, 0}, {"severity", 0, 0, 0}, {"filename", 0, 0, 0}, {"lineno", 0, 0, 0}, {"previous", 0, 0, 0}}
var ErrorExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{{"__construct", ZimErrorExceptionConstruct, ArginfoErrorExceptionConstruct, uint32(g.SizeOf("arginfo_error_exception___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 0}, {"getSeverity", zim_error_exception_getSeverity, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<5}, {nil, nil, nil, 0, 0}}

/* }}} */

func ZendRegisterDefaultException() {
	var ce zend_class_entry
	var ce ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Throwable", g.SizeOf("\"Throwable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsThrowable)
	ZendCeThrowable = ZendRegisterInternalInterface(&ce)
	ZendCeThrowable.interface_gets_implemented = ZendImplementThrowable
	memcpy(&DefaultExceptionHandlers, &StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	DefaultExceptionHandlers.SetCloneObj(nil)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Exception", g.SizeOf("\"Exception\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeException = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeException.create_object = ZendDefaultExceptionNew
	ZendClassImplements(ZendCeException, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeException, "message", g.SizeOf("\"message\"")-1, "", 1<<1)
	ZendDeclarePropertyString(ZendCeException, "string", g.SizeOf("\"string\"")-1, "", 1<<2)
	ZendDeclarePropertyLong(ZendCeException, "code", g.SizeOf("\"code\"")-1, 0, 1<<1)
	ZendDeclarePropertyNull(ZendCeException, "file", g.SizeOf("\"file\"")-1, 1<<1)
	ZendDeclarePropertyNull(ZendCeException, "line", g.SizeOf("\"line\"")-1, 1<<1)
	ZendDeclarePropertyNull(ZendCeException, "trace", g.SizeOf("\"trace\"")-1, 1<<2)
	ZendDeclarePropertyNull(ZendCeException, "previous", g.SizeOf("\"previous\"")-1, 1<<2)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ErrorException", g.SizeOf("\"ErrorException\"")-1, 1))
	ce.SetBuiltinFunctions(ErrorExceptionFunctions)
	ZendCeErrorException = ZendRegisterInternalClassEx(&ce, ZendCeException)
	ZendCeErrorException.create_object = ZendErrorExceptionNew
	ZendDeclarePropertyLong(ZendCeErrorException, "severity", g.SizeOf("\"severity\"")-1, 1<<0, 1<<1)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Error", g.SizeOf("\"Error\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeError = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeError.create_object = ZendDefaultExceptionNew
	ZendClassImplements(ZendCeError, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeError, "message", g.SizeOf("\"message\"")-1, "", 1<<1)
	ZendDeclarePropertyString(ZendCeError, "string", g.SizeOf("\"string\"")-1, "", 1<<2)
	ZendDeclarePropertyLong(ZendCeError, "code", g.SizeOf("\"code\"")-1, 0, 1<<1)
	ZendDeclarePropertyNull(ZendCeError, "file", g.SizeOf("\"file\"")-1, 1<<1)
	ZendDeclarePropertyNull(ZendCeError, "line", g.SizeOf("\"line\"")-1, 1<<1)
	ZendDeclarePropertyNull(ZendCeError, "trace", g.SizeOf("\"trace\"")-1, 1<<2)
	ZendDeclarePropertyNull(ZendCeError, "previous", g.SizeOf("\"previous\"")-1, 1<<2)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("CompileError", g.SizeOf("\"CompileError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeCompileError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeCompileError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ParseError", g.SizeOf("\"ParseError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeParseError = ZendRegisterInternalClassEx(&ce, ZendCeCompileError)
	ZendCeParseError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("TypeError", g.SizeOf("\"TypeError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeTypeError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeTypeError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArgumentCountError", g.SizeOf("\"ArgumentCountError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArgumentCountError = ZendRegisterInternalClassEx(&ce, ZendCeTypeError)
	ZendCeArgumentCountError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArithmeticError", g.SizeOf("\"ArithmeticError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArithmeticError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeArithmeticError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("DivisionByZeroError", g.SizeOf("\"DivisionByZeroError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeDivisionByZeroError = ZendRegisterInternalClassEx(&ce, ZendCeArithmeticError)
	ZendCeDivisionByZeroError.create_object = ZendDefaultExceptionNew
}

/* }}} */

func ZendExceptionGetDefault() *ZendClassEntry { return ZendCeException }

/* }}} */

func ZendGetErrorException() *ZendClassEntry { return ZendCeErrorException }

/* }}} */

func ZendThrowException(exception_ce *ZendClassEntry, message string, code ZendLong) *ZendObject {
	var ex Zval
	var tmp Zval
	if exception_ce != nil {
		if InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
			ZendError(1<<3, "Exceptions must implement Throwable")
			exception_ce = ZendCeException
		}
	} else {
		exception_ce = ZendCeException
	}
	ObjectInitEx(&ex, exception_ce)
	if message {
		var _s *byte = message
		var __z *Zval = &tmp
		var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		ZendUpdatePropertyEx(exception_ce, &ex, ZendKnownStrings[ZEND_STR_MESSAGE], &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		var __z *Zval = &tmp
		__z.GetValue().SetLval(code)
		__z.SetTypeInfo(4)
		ZendUpdatePropertyEx(exception_ce, &ex, ZendKnownStrings[ZEND_STR_CODE], &tmp)
	}
	ZendThrowExceptionInternal(&ex)
	return ex.GetValue().GetObj()
}

/* }}} */

func ZendThrowExceptionEx(exception_ce *ZendClassEntry, code ZendLong, format string, _ ...any) *ZendObject {
	var arg va_list
	var message *byte
	var obj *ZendObject
	va_start(arg, format)
	ZendVspprintf(&message, 0, format, arg)
	va_end(arg)
	obj = ZendThrowException(exception_ce, message, code)
	_efree(message)
	return obj
}

/* }}} */

func ZendThrowErrorException(exception_ce *ZendClassEntry, message *byte, code ZendLong, severity int) *ZendObject {
	var ex Zval
	var tmp Zval
	var obj *ZendObject = ZendThrowException(exception_ce, message, code)
	var __z *Zval = &ex
	__z.GetValue().SetObj(obj)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	var __z *Zval = &tmp
	__z.GetValue().SetLval(severity)
	__z.SetTypeInfo(4)
	ZendUpdatePropertyEx(ZendCeErrorException, &ex, ZendKnownStrings[ZEND_STR_SEVERITY], &tmp)
	return obj
}

/* }}} */

func ZendErrorVa(type_ int, file *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	ZendErrorCb(type_, file, lineno, format, args)
	va_end(args)
}

/* }}} */

func ZendErrorHelper(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var va va_list
	va_start(va, format)
	ZendErrorCb(type_, filename, lineno, format, va)
	va_end(va)
}

/* }}} */

func ZendExceptionError(ex *ZendObject, severity int) {
	var exception Zval
	var rv Zval
	var ce_exception *ZendClassEntry
	var __z *Zval = &exception
	__z.GetValue().SetObj(ex)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ce_exception = ex.GetCe()
	EG.SetException(nil)
	if ce_exception == ZendCeParseError || ce_exception == ZendCeCompileError {
		var message *ZendString = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_MESSAGE], 0, &rv))
		var file *ZendString = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_FILE], 1, &rv))
		var line ZendLong = ZvalGetLong(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_LINE], 1, &rv))
		ZendErrorHelper(g.Cond(ce_exception == ZendCeParseError, 1<<2, 1<<6), file.GetVal(), line, "%s", message.GetVal())
		ZendStringReleaseEx(file, 0)
		ZendStringReleaseEx(message, 0)
	} else if InstanceofFunction(ce_exception, ZendCeThrowable) != 0 {
		var tmp Zval
		var str *ZendString
		var file *ZendString = nil
		var line ZendLong = 0
		ZendCallMethod(&exception, ce_exception, &ex.ce.GetTostring(), "__tostring", g.SizeOf("\"__tostring\"")-1, &tmp, 0, nil, nil)
		if EG.GetException() == nil {
			if tmp.GetType() != 6 {
				ZendError(1<<1, "%s::__toString() must return a string", ce_exception.GetName().GetVal())
			} else {
				ZendUpdatePropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_STRING], &tmp)
			}
		}
		ZvalPtrDtor(&tmp)
		if EG.GetException() != nil {
			var zv Zval
			var __z *Zval = &zv
			__z.GetValue().SetObj(EG.GetException())
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)

			/* do the best we can to inform about the inner exception */

			if InstanceofFunction(ce_exception, ZendCeException) != 0 || InstanceofFunction(ce_exception, ZendCeError) != 0 {
				file = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(&zv), &zv, ZendKnownStrings[ZEND_STR_FILE], 1, &rv))
				line = ZvalGetLong(ZendReadPropertyEx(IGetExceptionBase(&zv), &zv, ZendKnownStrings[ZEND_STR_LINE], 1, &rv))
			}
			ZendErrorVa(1<<1, g.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s in exception handling during call to %s::__tostring()", zv.GetValue().GetObj().GetCe().GetName().GetVal(), ce_exception.GetName().GetVal())
			if file != nil {
				ZendStringReleaseEx(file, 0)
			}
		}
		str = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_STRING], 1, &rv))
		file = ZvalGetString(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_FILE], 1, &rv))
		line = ZvalGetLong(ZendReadPropertyEx(IGetExceptionBase(&exception), &exception, ZendKnownStrings[ZEND_STR_LINE], 1, &rv))
		ZendErrorVa(severity, g.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s\n  thrown", str.GetVal())
		ZendStringReleaseEx(str, 0)
		ZendStringReleaseEx(file, 0)
	} else {
		ZendError(severity, "Uncaught exception '%s'", ce_exception.GetName().GetVal())
	}
	ZendObjectRelease(ex)
}

/* }}} */

func ZendThrowExceptionObject(exception *Zval) {
	var exception_ce *ZendClassEntry
	if exception == nil || exception.GetType() != 8 {
		ZendErrorNoreturn(1<<4, "Need to supply an object when throwing an exception")
	}
	exception_ce = exception.GetValue().GetObj().GetCe()
	if exception_ce == nil || InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
		ZendThrowError(nil, "Cannot throw objects that do not implement Throwable")
		ZvalPtrDtor(exception)
		return
	}
	ZendThrowExceptionInternal(exception)
}

/* }}} */

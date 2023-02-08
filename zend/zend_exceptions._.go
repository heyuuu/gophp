// <<generate>>

package zend

import (
	b "sik/builtin"
)

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

var DefaultExceptionHandlers ZendObjectHandlers

/* {{{ zend_implement_throwable */

/* {{{ proto string Exception|Error::getFile()
   Get the file in which the exception occurred */

/* {{{ proto string Exception|Error::__toString()
   Obtain the string representation of the Exception object */

var ZendFuncsThrowable []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("getMessage", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getCode", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getFile", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getLine", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getTrace", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getPrevious", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getTraceAsString", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("__toString", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	MakeZendInternalArgInfo("message", 0, 0, 0),
	MakeZendInternalArgInfo("code", 0, 0, 0),
	MakeZendInternalArgInfo("previous", 0, 0, 0),
}
var DefaultExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("__clone", ZimExceptionClone, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PRIVATE|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("__construct", ZimExceptionConstruct, ArginfoExceptionConstruct, uint32(b.SizeOf("arginfo_exception___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("__wakeup", ZimExceptionWakeup, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("getMessage", zim_exception_getMessage, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getCode", zim_exception_getCode, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getFile", zim_exception_getFile, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getLine", zim_exception_getLine, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getTrace", zim_exception_getTrace, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getPrevious", zim_exception_getPrevious, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("getTraceAsString", zim_exception_getTraceAsString, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry("__toString", zim_exception___toString, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoErrorExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	MakeZendInternalArgInfo("message", 0, 0, 0),
	MakeZendInternalArgInfo("code", 0, 0, 0),
	MakeZendInternalArgInfo("severity", 0, 0, 0),
	MakeZendInternalArgInfo("filename", 0, 0, 0),
	MakeZendInternalArgInfo("lineno", 0, 0, 0),
	MakeZendInternalArgInfo("previous", 0, 0, 0),
}
var ErrorExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("__construct", ZimErrorExceptionConstruct, ArginfoErrorExceptionConstruct, uint32(b.SizeOf("arginfo_error_exception___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("getSeverity", zim_error_exception_getSeverity, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_FINAL),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

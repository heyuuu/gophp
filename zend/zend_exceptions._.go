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

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto string Exception|Error::getFile()
   Get the file in which the exception occurred */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto string Exception|Error::__toString()
   Obtain the string representation of the Exception object */

/* }}} */

var ZendFuncsThrowable []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"getMessage",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getCode",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getFile",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getLine",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getTrace",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getPrevious",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"getTraceAsString",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"__toString",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

var ArginfoExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"message", 0, 0, 0}, {"code", 0, 0, 0}, {"previous", 0, 0, 0}}
var DefaultExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__clone",
		ZimExceptionClone,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PRIVATE | ZEND_ACC_FINAL,
	},
	{
		"__construct",
		ZimExceptionConstruct,
		ArginfoExceptionConstruct,
		uint32_t(b.SizeOf("arginfo_exception___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"__wakeup",
		ZimExceptionWakeup,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"getMessage",
		zim_exception_getMessage,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getCode",
		zim_exception_getCode,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getFile",
		zim_exception_getFile,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getLine",
		zim_exception_getLine,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getTrace",
		zim_exception_getTrace,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getPrevious",
		zim_exception_getPrevious,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"getTraceAsString",
		zim_exception_getTraceAsString,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{
		"__toString",
		zim_exception___toString,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoErrorExceptionConstruct []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"message", 0, 0, 0}, {"code", 0, 0, 0}, {"severity", 0, 0, 0}, {"filename", 0, 0, 0}, {"lineno", 0, 0, 0}, {"previous", 0, 0, 0}}
var ErrorExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		ZimErrorExceptionConstruct,
		ArginfoErrorExceptionConstruct,
		uint32_t(b.SizeOf("arginfo_error_exception___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"getSeverity",
		zim_error_exception_getSeverity,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_FINAL,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

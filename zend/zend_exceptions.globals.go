// <<generate>>

package zend

import (
	b "sik/builtin"
)

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
var ZendThrowExceptionHook func(ex *Zval)
var DefaultExceptionHandlers ZendObjectHandlers
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

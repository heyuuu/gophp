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
	MakeZendFunctionEntry("getMessage", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getCode", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getFile", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getLine", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getTrace", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getPrevious", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("getTraceAsString", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("__toString", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
}
var ArginfoExceptionConstruct []ArgInfo = []ArgInfo{
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
}
var ArginfoErrorExceptionConstruct []ArgInfo = []ArgInfo{
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

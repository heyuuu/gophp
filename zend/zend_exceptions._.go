// <<generate>>

package zend

import (
	"sik/zend/types"
)

var ZendCeThrowable *types.ClassEntry
var ZendCeException *types.ClassEntry
var ZendCeErrorException *types.ClassEntry
var ZendCeError *types.ClassEntry
var ZendCeCompileError *types.ClassEntry
var ZendCeParseError *types.ClassEntry
var ZendCeTypeError *types.ClassEntry
var ZendCeArgumentCountError *types.ClassEntry
var ZendCeArithmeticError *types.ClassEntry
var ZendCeDivisionByZeroError *types.ClassEntry
var ZendThrowExceptionHook func(ex *types.Zval)

var DefaultExceptionHandlers ZendObjectHandlers

var ZendFuncsThrowable []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntryEx("getMessage", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getCode", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getFile", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getLine", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getTrace", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getPrevious", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("getTraceAsString", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	MakeZendFunctionEntryEx("__toString", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
}
var ArginfoExceptionConstruct []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("message"),
	MakeArgInfo("code"),
	MakeArgInfo("previous"),
}
var DefaultExceptionFunctions = []ZendFunctionEntry{
	MakeZendFunctionEntryEx("__clone", ZEND_ACC_PRIVATE|ZEND_ACC_FINAL, ZimExceptionClone, nil),
	MakeZendFunctionEntryEx("__construct", ZEND_ACC_PUBLIC, ZimExceptionConstruct, ArginfoExceptionConstruct),
	MakeZendFunctionEntryEx("__wakeup", ZEND_ACC_PUBLIC, ZimExceptionWakeup, nil),
	MakeZendFunctionEntryEx("getMessage", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getMessage, nil),
	MakeZendFunctionEntryEx("getCode", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getCode, nil),
	MakeZendFunctionEntryEx("getFile", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getFile, nil),
	MakeZendFunctionEntryEx("getLine", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getLine, nil),
	MakeZendFunctionEntryEx("getTrace", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getTrace, nil),
	MakeZendFunctionEntryEx("getPrevious", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getPrevious, nil),
	MakeZendFunctionEntryEx("getTraceAsString", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_exception_getTraceAsString, nil),
	MakeZendFunctionEntryEx("__toString", 0, zim_exception___toString, nil),
}
var ArginfoErrorExceptionConstruct []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("message"),
	MakeArgInfo("code"),
	MakeArgInfo("severity"),
	MakeArgInfo("filename"),
	MakeArgInfo("lineno"),
	MakeArgInfo("previous"),
}
var ErrorExceptionFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntryEx("__construct", ZEND_ACC_PUBLIC, ZimErrorExceptionConstruct, ArginfoErrorExceptionConstruct),
	MakeZendFunctionEntryEx("getSeverity", ZEND_ACC_PUBLIC|ZEND_ACC_FINAL, zim_error_exception_getSeverity, nil),
}

// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendCeClosure *ZendClassEntry

const ZEND_CLOSURE_PRINT_NAME = "Closure object"

var ClosureHandlers ZendObjectHandlers
var ArginfoClosureBindto []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"newthis", 0, 0, 0}, {"newscope", 0, 0, 0}}
var ArginfoClosureBind []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"closure", 0, 0, 0}, {"newthis", 0, 0, 0}, {"newscope", 0, 0, 0}}
var ArginfoClosureCall []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"newthis", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoClosureFromcallable []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"callable", 0, 0, 0}}
var ClosureFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_Closure___construct,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PRIVATE,
	},
	{
		"bind",
		zim_Closure_bind,
		ArginfoClosureBind,
		uint32_t(b.SizeOf("arginfo_closure_bind")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{
		"bindTo",
		zim_Closure_bind,
		ArginfoClosureBindto,
		uint32_t(b.SizeOf("arginfo_closure_bindto")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"call",
		zim_Closure_call,
		ArginfoClosureCall,
		uint32_t(b.SizeOf("arginfo_closure_call")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"fromCallable",
		zim_Closure_fromCallable,
		ArginfoClosureFromcallable,
		uint32_t(b.SizeOf("arginfo_closure_fromcallable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{nil, nil, nil, 0, 0},
}

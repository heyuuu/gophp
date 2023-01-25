// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendCeGenerator *ZendClassEntry
var zend_ce_ClosedGeneratorException *ZendClassEntry

var ZEND_GENERATOR_CURRENTLY_RUNNING ZendUchar = 0x1
var ZEND_GENERATOR_FORCED_CLOSE ZendUchar = 0x2
var ZEND_GENERATOR_AT_FIRST_YIELD ZendUchar = 0x4
var ZEND_GENERATOR_DO_INIT ZendUchar = 0x8
var ZendGeneratorHandlers ZendObjectHandlers
var ZendGeneratorIteratorFunctions ZendObjectIteratorFuncs = ZendObjectIteratorFuncs{ZendGeneratorIteratorDtor, ZendGeneratorIteratorValid, ZendGeneratorIteratorGetData, ZendGeneratorIteratorGetKey, ZendGeneratorIteratorMoveForward, ZendGeneratorIteratorRewind, nil}
var ArginfoGeneratorVoid []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0},
}
var ArginfoGeneratorSend []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoGeneratorThrow []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"exception", 0, 0, 0}}
var GeneratorFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"rewind",
		zim_Generator_rewind,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_Generator_valid,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_Generator_current,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_Generator_key,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_Generator_next,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"send",
		zim_Generator_send,
		ArginfoGeneratorSend,
		uint32_t(b.SizeOf("arginfo_generator_send")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"throw",
		zim_Generator_throw,
		ArginfoGeneratorThrow,
		uint32_t(b.SizeOf("arginfo_generator_throw")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"getReturn",
		zim_Generator_getReturn,
		ArginfoGeneratorVoid,
		uint32_t(b.SizeOf("arginfo_generator_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

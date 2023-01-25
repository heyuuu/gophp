// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendCeTraversable *ZendClassEntry
var ZendCeAggregate *ZendClassEntry
var ZendCeIterator *ZendClassEntry
var ZendCeArrayaccess *ZendClassEntry
var ZendCeSerializable *ZendClassEntry
var ZendCeCountable *ZendClassEntry

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = ZendObjectIteratorFuncs{ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent}
var ZendFuncsAggregate []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"getIterator",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var ZendFuncsIterator []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"current",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"next",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"key",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"valid",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"rewind",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
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
		uint32_t(b.SizeOf("arginfo_arrayaccess_offset")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"offsetGet",
		nil,
		ArginfoArrayaccessOffsetGet,
		uint32_t(b.SizeOf("arginfo_arrayaccess_offset_get")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"offsetSet",
		nil,
		ArginfoArrayaccessOffsetValue,
		uint32_t(b.SizeOf("arginfo_arrayaccess_offset_value")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"offsetUnset",
		nil,
		ArginfoArrayaccessOffset,
		uint32_t(b.SizeOf("arginfo_arrayaccess_offset")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoSerializableSerialize []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0},
	{"serialized", 0, 0, 0},
}
var ZendFuncsSerializable []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"serialize",
		nil,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{
		"unserialize",
		nil,
		ArginfoSerializableSerialize,
		uint32_t(b.SizeOf("arginfo_serializable_serialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCountableCount []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0},
}
var ZendFuncsCountable []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"count",
		nil,
		ArginfoCountableCount,
		uint32_t(b.SizeOf("arginfo_countable_count")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}

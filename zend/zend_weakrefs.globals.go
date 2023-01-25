// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendCeWeakref *ZendClassEntry

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefCreateArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{
		(*byte)(zend_uintptr_t(1)),
		ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0),
		0,
		0,
	},
	{"referent", ZEND_TYPE_ENCODE(IS_OBJECT, 0), 0, 0},
}
var ZendWeakrefGetArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(0)), ZEND_TYPE_ENCODE(IS_OBJECT, 1), 0, 0},
}
var ZendWeakrefMethods []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_WeakReference___construct,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"create",
		zim_WeakReference_create,
		ZendWeakrefCreateArginfo,
		uint32_t(b.SizeOf("zend_weakref_create_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{
		"get",
		zim_WeakReference_get,
		ZendWeakrefGetArginfo,
		uint32_t(b.SizeOf("zend_weakref_get_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

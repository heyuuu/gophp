// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendCeWeakref *ZendClassEntry

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefCreateArginfo = []ArgInfo{
	MakeReturnArgInfo(1,ArgInfoType(ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0))),
	MakeArgInfo("referent",ArgInfoType(ZEND_TYPE_ENCODE(IS_OBJECT, 0))),
}
var ZendWeakrefGetArginfo = []ArgInfo{
	MakeReturnArgInfo(0,ArgInfoType(ZEND_TYPE_ENCODE(IS_OBJECT, 1))),
}
var ZendWeakrefMethods = []ZendFunctionEntry{
	MakeZendFunctionEntry("__construct", zim_WeakReference___construct, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("create", zim_WeakReference_create, ZendWeakrefCreateArginfo, uint32(b.SizeOf("zend_weakref_create_arginfo")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_STATIC),
	MakeZendFunctionEntry("get", zim_WeakReference_get, ZendWeakrefGetArginfo, uint32(b.SizeOf("zend_weakref_get_arginfo")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

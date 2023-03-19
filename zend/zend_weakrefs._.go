// <<generate>>

package zend

import "sik/zend/types"

var ZendCeWeakref *types.ClassEntry

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefCreateArginfo = []ArgInfo{
	MakeReturnArgInfo(1, ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0))),
	MakeArgInfo("referent", ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_OBJECT, 0))),
}
var ZendWeakrefGetArginfo = []ArgInfo{
	MakeReturnArgInfo(0, ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_OBJECT, 1))),
}
var ZendWeakrefMethods = []ZendFunctionEntry{
	MakeZendFunctionEntryEx("__construct", ZEND_ACC_PUBLIC, zim_WeakReference___construct, nil),
	MakeZendFunctionEntryEx("create", ZEND_ACC_PUBLIC|ZEND_ACC_STATIC, zim_WeakReference_create, ZendWeakrefCreateArginfo),
	MakeZendFunctionEntryEx("get", ZEND_ACC_PUBLIC, zim_WeakReference_get, ZendWeakrefGetArginfo),
}

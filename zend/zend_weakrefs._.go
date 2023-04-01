package zend

import "github.com/heyuuu/gophp/zend/types"

var ZendCeWeakref *types.ClassEntry

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefMethods = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", AccPublic, zim_WeakReference___construct, nil),
	types.MakeZendFunctionEntryEx("create", AccPublic|AccStatic, zim_WeakReference_create, []ArgInfo{MakeReturnArgInfo(1, ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0))),
		MakeArgInfo("referent", ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_OBJECT, 0))),
	}),
	types.MakeZendFunctionEntryEx("get", AccPublic, zim_WeakReference_get, []ArgInfo{MakeReturnArgInfo(0, ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_OBJECT, 1)))}),
}

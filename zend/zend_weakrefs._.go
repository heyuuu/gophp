package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

var ZendCeWeakref *types2.ClassEntry

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefMethods = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", AccPublic, zim_WeakReference___construct, nil),
	types2.MakeZendFunctionEntryEx("create", AccPublic|AccStatic, zim_WeakReference_create, []ArgInfo{MakeReturnArgInfo(1, ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0))),
		MakeArgInfo("referent", ArgInfoType(types2.ZEND_TYPE_ENCODE(types2.IS_OBJECT, 0))),
	}),
	types2.MakeZendFunctionEntryEx("get", AccPublic, zim_WeakReference_get, []ArgInfo{MakeReturnArgInfo(0, ArgInfoType(types2.ZEND_TYPE_ENCODE(types2.IS_OBJECT, 1)))}),
}

package zend

import "sik/zend/types"

var ZendCeTraversable *types.ClassEntry
var ZendCeAggregate *types.ClassEntry
var ZendCeIterator *types.ClassEntry
var ZendCeArrayaccess *types.ClassEntry
var ZendCeSerializable *types.ClassEntry
var ZendCeCountable *types.ClassEntry

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent)

var ZendFuncsAggregate []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("getIterator", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
}
var ZendFuncsIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("current", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("next", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("key", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("valid", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("rewind", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
}
var ZendFuncsTraversable *types.ZendFunctionEntry = nil
var ZendFuncsArrayaccess []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("offsetExists", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("offset"),
		MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
}
var ZendFuncsSerializable []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("serialize", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("unserialize", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(-1),
		MakeArgName("serialized"),
	}),
}
var ZendFuncsCountable []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("count", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, []ArgInfo{MakeReturnArgInfo(-1)}),
}

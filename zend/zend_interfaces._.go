package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

var ZendCeTraversable *types2.ClassEntry
var ZendCeAggregate *types2.ClassEntry
var ZendCeIterator *types2.ClassEntry
var ZendCeArrayaccess *types2.ClassEntry
var ZendCeSerializable *types2.ClassEntry
var ZendCeCountable *types2.ClassEntry

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent)

var ZendFuncsAggregate []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("getIterator", AccPublic|AccAbstract, nil, nil),
}
var ZendFuncsIterator []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("current", AccPublic|AccAbstract, nil, nil),
	types2.MakeZendFunctionEntryEx("next", AccPublic|AccAbstract, nil, nil),
	types2.MakeZendFunctionEntryEx("key", AccPublic|AccAbstract, nil, nil),
	types2.MakeZendFunctionEntryEx("valid", AccPublic|AccAbstract, nil, nil),
	types2.MakeZendFunctionEntryEx("rewind", AccPublic|AccAbstract, nil, nil),
}
var ZendFuncsTraversable *types2.FunctionEntry = nil
var ZendFuncsArrayaccess []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("offsetExists", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types2.MakeZendFunctionEntryEx("offsetGet", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types2.MakeZendFunctionEntryEx("offsetSet", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("offset"),
		MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("offsetUnset", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
}
var ZendFuncsSerializable []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("serialize", AccPublic|AccAbstract, nil, nil),
	types2.MakeZendFunctionEntryEx("unserialize", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1),
		MakeArgName("serialized"),
	}),
}
var ZendFuncsCountable []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("count", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1)}),
}

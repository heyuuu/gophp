package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

var ZendCeTraversable *types.ClassEntry
var ZendCeAggregate *types.ClassEntry
var ZendCeIterator *types.ClassEntry
var ZendCeArrayaccess *types.ClassEntry
var ZendCeSerializable *types.ClassEntry
var ZendCeCountable *types.ClassEntry

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent)

var ZendFuncsAggregate []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("getIterator", AccPublic|AccAbstract, nil, nil),
}
var ZendFuncsIterator []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("current", AccPublic|AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("next", AccPublic|AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("key", AccPublic|AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("valid", AccPublic|AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("rewind", AccPublic|AccAbstract, nil, nil),
}
var ZendFuncsTraversable *types.FunctionEntry = nil
var ZendFuncsArrayaccess []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("offsetExists", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("offset"),
		MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
}
var ZendFuncsSerializable []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("serialize", AccPublic|AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("unserialize", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1),
		MakeArgName("serialized"),
	}),
}
var ZendFuncsCountable []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("count", AccPublic|AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1)}),
}

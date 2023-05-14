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
	types.MakeZendFunctionEntryEx("getIterator", types.AccPublic|types.AccAbstract, nil, nil),
}
var ZendFuncsIterator []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("current", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("next", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("key", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic|types.AccAbstract, nil, nil),
}
var ZendFuncsArrayaccess []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("offsetExists", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("offset"),
		MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("offset"),
	}),
}
var ZendFuncsSerializable []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("serialize", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("unserialize", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1),
		MakeArgName("serialized"),
	}),
}
var ZendFuncsCountable []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("count", types.AccPublic|types.AccAbstract, nil, []ArgInfo{MakeReturnArgInfo(-1)}),
}

package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplFixedArray *types.ClassEntry

var spl_handler_SplFixedArray types.ObjectHandlers

const SPL_FIXEDARRAY_OVERLOADED_REWIND = 0x1
const SPL_FIXEDARRAY_OVERLOADED_VALID = 0x2
const SPL_FIXEDARRAY_OVERLOADED_KEY = 0x4
const SPL_FIXEDARRAY_OVERLOADED_CURRENT = 0x8
const SPL_FIXEDARRAY_OVERLOADED_NEXT = 0x10

var SplFixedarrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFixedarrayItDtor, SplFixedarrayItValid, SplFixedarrayItGetCurrentData, SplFixedarrayItGetCurrentKey, SplFixedarrayItMoveForward, SplFixedarrayItRewind, nil)
var spl_funcs_SplFixedArray []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_SplFixedArray___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("size"),
	}),
	types.MakeZendFunctionEntryEx("__wakeup", types.AccPublic, zim_spl_SplFixedArray___wakeup, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", types.AccPublic, zim_spl_SplFixedArray_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("toArray", types.AccPublic, zim_spl_SplFixedArray_toArray, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fromArray", types.AccPublic|types.AccStatic, zim_spl_SplFixedArray_fromArray, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("array"),
		zend.MakeArgName("save_indexes"),
	}),
	types.MakeZendFunctionEntryEx("getSize", types.AccPublic, zim_spl_SplFixedArray_getSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setSize", types.AccPublic, zim_spl_SplFixedArray_setSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("offsetExists", types.AccPublic, zim_spl_SplFixedArray_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", types.AccPublic, zim_spl_SplFixedArray_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", types.AccPublic, zim_spl_SplFixedArray_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", types.AccPublic, zim_spl_SplFixedArray_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_SplFixedArray_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_SplFixedArray_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_SplFixedArray_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_SplFixedArray_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic, zim_spl_SplFixedArray_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

package spl

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var spl_ce_SplFixedArray *types.ClassEntry

var spl_handler_SplFixedArray zend.ZendObjectHandlers

const SPL_FIXEDARRAY_OVERLOADED_REWIND = 0x1
const SPL_FIXEDARRAY_OVERLOADED_VALID = 0x2
const SPL_FIXEDARRAY_OVERLOADED_KEY = 0x4
const SPL_FIXEDARRAY_OVERLOADED_CURRENT = 0x8
const SPL_FIXEDARRAY_OVERLOADED_NEXT = 0x10

var SplFixedarrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFixedarrayItDtor, SplFixedarrayItValid, SplFixedarrayItGetCurrentData, SplFixedarrayItGetCurrentKey, SplFixedarrayItMoveForward, SplFixedarrayItRewind, nil)
var spl_funcs_SplFixedArray []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_SplFixedArray___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("size"),
	}),
	types.MakeZendFunctionEntryEx("__wakeup", zend.AccPublic, zim_spl_SplFixedArray___wakeup, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplFixedArray_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("toArray", zend.AccPublic, zim_spl_SplFixedArray_toArray, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fromArray", zend.AccPublic|zend.AccStatic, zim_spl_SplFixedArray_fromArray, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("array"),
		zend.MakeArgName("save_indexes"),
	}),
	types.MakeZendFunctionEntryEx("getSize", zend.AccPublic, zim_spl_SplFixedArray_getSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setSize", zend.AccPublic, zim_spl_SplFixedArray_setSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_SplFixedArray_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_SplFixedArray_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_SplFixedArray_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_SplFixedArray_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplFixedArray_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplFixedArray_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplFixedArray_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplFixedArray_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplFixedArray_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

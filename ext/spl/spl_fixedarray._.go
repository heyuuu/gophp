package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplFixedArray *types2.ClassEntry

var spl_handler_SplFixedArray zend.ZendObjectHandlers

const SPL_FIXEDARRAY_OVERLOADED_REWIND = 0x1
const SPL_FIXEDARRAY_OVERLOADED_VALID = 0x2
const SPL_FIXEDARRAY_OVERLOADED_KEY = 0x4
const SPL_FIXEDARRAY_OVERLOADED_CURRENT = 0x8
const SPL_FIXEDARRAY_OVERLOADED_NEXT = 0x10

var SplFixedarrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFixedarrayItDtor, SplFixedarrayItValid, SplFixedarrayItGetCurrentData, SplFixedarrayItGetCurrentKey, SplFixedarrayItMoveForward, SplFixedarrayItRewind, nil)
var spl_funcs_SplFixedArray []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_SplFixedArray___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("size"),
	}),
	types2.MakeZendFunctionEntryEx("__wakeup", zend.AccPublic, zim_spl_SplFixedArray___wakeup, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_SplFixedArray_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("toArray", zend.AccPublic, zim_spl_SplFixedArray_toArray, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fromArray", zend.AccPublic|zend.AccStatic, zim_spl_SplFixedArray_fromArray, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("array"),
		zend.MakeArgName("save_indexes"),
	}),
	types2.MakeZendFunctionEntryEx("getSize", zend.AccPublic, zim_spl_SplFixedArray_getSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setSize", zend.AccPublic, zim_spl_SplFixedArray_setSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_SplFixedArray_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_SplFixedArray_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_SplFixedArray_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types2.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_SplFixedArray_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplFixedArray_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplFixedArray_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplFixedArray_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplFixedArray_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplFixedArray_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

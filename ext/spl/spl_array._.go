package spl

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var spl_ce_ArrayObject *types.ClassEntry
var spl_ce_ArrayIterator *types.ClassEntry
var spl_ce_RecursiveArrayIterator *types.ClassEntry

var spl_handler_ArrayObject zend.ZendObjectHandlers
var spl_handler_ArrayIterator zend.ZendObjectHandlers

const SPL_ARRAY_STD_PROP_LIST = 0x1
const SPL_ARRAY_ARRAY_AS_PROPS = 0x2
const SPL_ARRAY_CHILD_ARRAYS_ONLY = 0x4
const SPL_ARRAY_OVERLOADED_REWIND = 0x10000
const SPL_ARRAY_OVERLOADED_VALID = 0x20000
const SPL_ARRAY_OVERLOADED_KEY = 0x40000
const SPL_ARRAY_OVERLOADED_CURRENT = 0x80000
const SPL_ARRAY_OVERLOADED_NEXT = 0x100000
const SPL_ARRAY_IS_SELF = 0x1000000
const SPL_ARRAY_USE_OTHER = 0x2000000
const SPL_ARRAY_INT_MASK = 0xffff0000
const SPL_ARRAY_CLONE_MASK = 0x100ffff
const SPL_ARRAY_METHOD_NO_ARG = 0
const SPL_ARRAY_METHOD_USE_ARG = 1
const SPL_ARRAY_METHOD_MAY_USER_ARG = 2

var SplArrayItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplArrayItDtor, SplArrayItValid, SplArrayItGetCurrentData, SplArrayItGetCurrentKey, SplArrayItMoveForward, SplArrayItRewind, nil)

/* ArrayIterator::__construct and ArrayObject::__construct have different signatures */
var spl_funcs_ArrayObject []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_Array___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("input"),
		zend.MakeArgName("flags"),
		zend.MakeArgName("iterator_class"),
	}),
	types.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_Array_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_Array_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_Array_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_Array_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("append", zend.AccPublic, zim_spl_Array_append, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("getArrayCopy", zend.AccPublic, zim_spl_Array_getArrayCopy, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_Array_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_Array_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_Array_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("asort", zend.AccPublic, zim_spl_Array_asort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("ksort", zend.AccPublic, zim_spl_Array_ksort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("uasort", zend.AccPublic, zim_spl_Array_uasort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("cmp_function"),
	}),
	types.MakeZendFunctionEntryEx("uksort", zend.AccPublic, zim_spl_Array_uksort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("cmp_function"),
	}),
	types.MakeZendFunctionEntryEx("natsort", zend.AccPublic, zim_spl_Array_natsort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("natcasesort", zend.AccPublic, zim_spl_Array_natcasesort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("unserialize", zend.AccPublic, zim_spl_Array_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("serialize", zend.AccPublic, zim_spl_Array_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__unserialize", zend.AccPublic, zim_spl_Array___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("__serialize", zend.AccPublic, zim_spl_Array___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_Array___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getIterator", zend.AccPublic, zim_spl_Array_getIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("exchangeArray", zend.AccPublic, zim_spl_Array_exchangeArray, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("input"),
	}),
	types.MakeZendFunctionEntryEx("setIteratorClass", zend.AccPublic, zim_spl_Array_setIteratorClass, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("iteratorClass"),
	}),
	types.MakeZendFunctionEntryEx("getIteratorClass", zend.AccPublic, zim_spl_Array_getIteratorClass, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_ArrayIterator []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_ArrayIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("array"),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_Array_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_Array_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_Array_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_Array_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("append", zend.AccPublic, zim_spl_Array_append, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("getArrayCopy", zend.AccPublic, zim_spl_Array_getArrayCopy, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_Array_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_Array_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_Array_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("asort", zend.AccPublic, zim_spl_Array_asort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("ksort", zend.AccPublic, zim_spl_Array_ksort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("uasort", zend.AccPublic, zim_spl_Array_uasort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("cmp_function"),
	}),
	types.MakeZendFunctionEntryEx("uksort", zend.AccPublic, zim_spl_Array_uksort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("cmp_function"),
	}),
	types.MakeZendFunctionEntryEx("natsort", zend.AccPublic, zim_spl_Array_natsort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("natcasesort", zend.AccPublic, zim_spl_Array_natcasesort, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("unserialize", zend.AccPublic, zim_spl_Array_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("serialize", zend.AccPublic, zim_spl_Array_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__unserialize", zend.AccPublic, zim_spl_Array___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("__serialize", zend.AccPublic, zim_spl_Array___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_Array___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_Array_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_Array_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_Array_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_Array_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_Array_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("seek", zend.AccPublic, zim_spl_Array_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("position"),
	}),
}
var spl_funcs_RecursiveArrayIterator []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_Array_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_Array_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

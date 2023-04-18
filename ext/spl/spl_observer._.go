package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplObserver *types2.ClassEntry
var spl_ce_SplSubject *types2.ClassEntry
var spl_ce_SplObjectStorage *types2.ClassEntry
var spl_ce_MultipleIterator *types2.ClassEntry

var spl_funcs_SplObserver = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("update", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("subject", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("SplSubject", 0))),
	}),
}
var spl_funcs_SplSubject = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("attach", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("observer", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0))),
	}),
	types2.MakeZendFunctionEntryEx("detach", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("observer", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0))),
	}),
	types2.MakeZendFunctionEntryEx("notify", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_handler_SplObjectStorage zend.ZendObjectHandlers

var spl_funcs_SplObjectStorage = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("attach", 0, zim_spl_SplObjectStorage_attach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
		zend.MakeArgName("data"),
	}),
	types2.MakeZendFunctionEntryEx("detach", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("contains", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("addAll", 0, zim_spl_SplObjectStorage_addAll, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("removeAll", 0, zim_spl_SplObjectStorage_removeAll, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("removeAllExcept", 0, zim_spl_SplObjectStorage_removeAllExcept, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("getInfo", 0, zim_spl_SplObjectStorage_getInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setInfo", 0, zim_spl_SplObjectStorage_setInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("info"),
	}),
	types2.MakeZendFunctionEntryEx("getHash", 0, zim_spl_SplObjectStorage_getHash, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("count", 0, zim_spl_SplObjectStorage_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("rewind", 0, zim_spl_SplObjectStorage_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", 0, zim_spl_SplObjectStorage_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", 0, zim_spl_SplObjectStorage_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", 0, zim_spl_SplObjectStorage_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", 0, zim_spl_SplObjectStorage_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("unserialize", 0, zim_spl_SplObjectStorage_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types2.MakeZendFunctionEntryEx("serialize", 0, zim_spl_SplObjectStorage_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__unserialize", 0, zim_spl_SplObjectStorage___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types2.MakeZendFunctionEntryEx("__serialize", 0, zim_spl_SplObjectStorage___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("offsetExists", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("offsetSet", 0, zim_spl_SplObjectStorage_attach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
		zend.MakeArgName("data"),
	}),
	types2.MakeZendFunctionEntryEx("offsetUnset", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
	}),
	types2.MakeZendFunctionEntryEx("offsetGet", 0, zim_spl_SplObjectStorage_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
	}),
}

type MultipleIteratorFlags = int

const (
	MIT_NEED_ANY     = 0
	MIT_NEED_ALL     = 1
	MIT_KEYS_NUMERIC = 0
	MIT_KEYS_ASSOC   = 2
)
const SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT = 1
const SPL_MULTIPLE_ITERATOR_GET_ALL_KEY = 2

var spl_funcs_MultipleIterator = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", 0, zim_spl_MultipleIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("getFlags", 0, zim_spl_MultipleIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setFlags", 0, zim_spl_MultipleIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("attachIterator", 0, zim_spl_MultipleIterator_attachIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("infos"),
	}),
	types2.MakeZendFunctionEntryEx("detachIterator", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types2.MakeZendFunctionEntryEx("containsIterator", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types2.MakeZendFunctionEntryEx("countIterators", 0, zim_spl_SplObjectStorage_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("rewind", 0, zim_spl_MultipleIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", 0, zim_spl_MultipleIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", 0, zim_spl_MultipleIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", 0, zim_spl_MultipleIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", 0, zim_spl_MultipleIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */

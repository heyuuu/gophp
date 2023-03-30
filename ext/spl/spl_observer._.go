package spl

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var spl_ce_SplObserver *types.ClassEntry
var spl_ce_SplSubject *types.ClassEntry
var spl_ce_SplObjectStorage *types.ClassEntry
var spl_ce_MultipleIterator *types.ClassEntry

var spl_funcs_SplObserver = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("update", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("subject", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("SplSubject", 0))),
	}),
}
var spl_funcs_SplSubject = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("attach", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("observer", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0))),
	}),
	types.MakeZendFunctionEntryEx("detach", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("observer", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0))),
	}),
	types.MakeZendFunctionEntryEx("notify", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_handler_SplObjectStorage zend.ZendObjectHandlers

var spl_funcs_SplObjectStorage = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("attach", 0, zim_spl_SplObjectStorage_attach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
		zend.MakeArgName("data"),
	}),
	types.MakeZendFunctionEntryEx("detach", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("contains", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("addAll", 0, zim_spl_SplObjectStorage_addAll, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("removeAll", 0, zim_spl_SplObjectStorage_removeAll, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("removeAllExcept", 0, zim_spl_SplObjectStorage_removeAllExcept, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("getInfo", 0, zim_spl_SplObjectStorage_getInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setInfo", 0, zim_spl_SplObjectStorage_setInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("info"),
	}),
	types.MakeZendFunctionEntryEx("getHash", 0, zim_spl_SplObjectStorage_getHash, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", 0, zim_spl_SplObjectStorage_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", 0, zim_spl_SplObjectStorage_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", 0, zim_spl_SplObjectStorage_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", 0, zim_spl_SplObjectStorage_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", 0, zim_spl_SplObjectStorage_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", 0, zim_spl_SplObjectStorage_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("unserialize", 0, zim_spl_SplObjectStorage_unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("serialize", 0, zim_spl_SplObjectStorage_serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__unserialize", 0, zim_spl_SplObjectStorage___unserialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("serialized"),
	}),
	types.MakeZendFunctionEntryEx("__serialize", 0, zim_spl_SplObjectStorage___serialize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("offsetExists", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", 0, zim_spl_SplObjectStorage_attach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
		zend.MakeArgName("data"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("object"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", 0, zim_spl_SplObjectStorage_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
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

var spl_funcs_MultipleIterator = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", 0, zim_spl_MultipleIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("getFlags", 0, zim_spl_MultipleIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", 0, zim_spl_MultipleIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("attachIterator", 0, zim_spl_MultipleIterator_attachIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("infos"),
	}),
	types.MakeZendFunctionEntryEx("detachIterator", 0, zim_spl_SplObjectStorage_detach, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("containsIterator", 0, zim_spl_SplObjectStorage_contains, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("countIterators", 0, zim_spl_SplObjectStorage_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", 0, zim_spl_MultipleIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", 0, zim_spl_MultipleIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", 0, zim_spl_MultipleIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", 0, zim_spl_MultipleIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", 0, zim_spl_MultipleIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */

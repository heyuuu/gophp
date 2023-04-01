package spl

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var spl_ce_Iterator = zend.ZendCeIterator
var spl_ce_Aggregate = zend.ZendCeAggregate
var spl_ce_ArrayAccess = zend.ZendCeArrayaccess
var spl_ce_Serializable = zend.ZendCeSerializable
var spl_ce_Countable = zend.ZendCeCountable

var spl_ce_RecursiveIterator *types.ClassEntry
var spl_ce_RecursiveIteratorIterator *types.ClassEntry
var spl_ce_RecursiveTreeIterator *types.ClassEntry
var spl_ce_FilterIterator *types.ClassEntry
var spl_ce_RecursiveFilterIterator *types.ClassEntry
var spl_ce_ParentIterator *types.ClassEntry
var spl_ce_SeekableIterator *types.ClassEntry
var spl_ce_LimitIterator *types.ClassEntry
var spl_ce_CachingIterator *types.ClassEntry
var spl_ce_RecursiveCachingIterator *types.ClassEntry
var spl_ce_OuterIterator *types.ClassEntry
var spl_ce_IteratorIterator *types.ClassEntry
var spl_ce_NoRewindIterator *types.ClassEntry
var spl_ce_InfiniteIterator *types.ClassEntry
var spl_ce_EmptyIterator *types.ClassEntry
var spl_ce_AppendIterator *types.ClassEntry
var spl_ce_RegexIterator *types.ClassEntry
var spl_ce_RecursiveRegexIterator *types.ClassEntry
var spl_ce_CallbackFilterIterator *types.ClassEntry
var spl_ce_RecursiveCallbackFilterIterator *types.ClassEntry

type DualItType = int

const (
	DIT_Default                            = 0
	DIT_FilterIterator          DualItType = DIT_Default
	DIT_RecursiveFilterIterator DualItType = DIT_Default
	DIT_ParentIterator          DualItType = DIT_Default
	DIT_LimitIterator
	DIT_CachingIterator
	DIT_RecursiveCachingIterator
	DIT_IteratorIterator
	DIT_NoRewindIterator
	DIT_InfiniteIterator
	DIT_AppendIterator
	DIT_RegexIterator
	DIT_RecursiveRegexIterator
	DIT_CallbackFilterIterator
	DIT_RecursiveCallbackFilterIterator
	DIT_Unknown DualItType = ^0
)

type RecursiveItItType = int

const (
	RIT_Default                                     = 0
	RIT_RecursiveIteratorIterator RecursiveItItType = RIT_Default
	RIT_RecursiveTreeIterator
	RIT_Unknow RecursiveItItType = ^0
)
const (
	CIT_CALL_TOSTRING        = 0x1
	CIT_TOSTRING_USE_KEY     = 0x2
	CIT_TOSTRING_USE_CURRENT = 0x4
	CIT_TOSTRING_USE_INNER   = 0x8
	CIT_CATCH_GET_CHILD      = 0x10
	CIT_FULL_CACHE           = 0x100
	CIT_PUBLIC               = 0xffff
	CIT_VALID                = 0x10000
	CIT_HAS_CHILDREN         = 0x20000
)
const (
	REGIT_USE_KEY  = 0x1
	REGIT_INVERTED = 0x2
)

type RegexMode = int

const (
	REGIT_MODE_MATCH = iota
	REGIT_MODE_GET_MATCH
	REGIT_MODE_ALL_MATCHES
	REGIT_MODE_SPLIT
	REGIT_MODE_REPLACE
	REGIT_MODE_MAX
)

type SplIteratorApplyFuncT func(iter *zend.ZendObjectIterator, puser any) int

var spl_funcs_RecursiveIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

type RecursiveIteratorMode = int

const (
	RIT_LEAVES_ONLY = 0
	RIT_SELF_FIRST  = 1
	RIT_CHILD_FIRST = 2
)
const RIT_CATCH_GET_CHILD = CIT_CATCH_GET_CHILD

type RecursiveTreeIteratorFlags = int

const (
	RTIT_BYPASS_CURRENT = 4
	RTIT_BYPASS_KEY     = 8
)

type RecursiveIteratorState = int

const (
	RS_NEXT  = 0
	RS_TEST  = 1
	RS_SELF  = 2
	RS_CHILD = 3
	RS_START = 4
)

var SplHandlersRecItIt zend.ZendObjectHandlers
var SplHandlersDualIt zend.ZendObjectHandlers
var SplRecursiveItIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplRecursiveItDtor, SplRecursiveItValid, SplRecursiveItGetCurrentData, SplRecursiveItGetCurrentKey, SplRecursiveItMoveForward, SplRecursiveItRewind, nil)

var spl_funcs_RecursiveIteratorIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveIteratorIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("mode"),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_RecursiveIteratorIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_RecursiveIteratorIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_RecursiveIteratorIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_RecursiveIteratorIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_RecursiveIteratorIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getDepth", zend.AccPublic, zim_spl_RecursiveIteratorIterator_getDepth, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getSubIterator", zend.AccPublic, zim_spl_RecursiveIteratorIterator_getSubIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("level"),
	}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_RecursiveIteratorIterator_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("beginIteration", zend.AccPublic, zim_spl_RecursiveIteratorIterator_beginIteration, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("endIteration", zend.AccPublic, zim_spl_RecursiveIteratorIterator_endIteration, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("callHasChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_callHasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("callGetChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_callGetChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("beginChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_beginChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("endChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_endChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("nextElement", zend.AccPublic, zim_spl_RecursiveIteratorIterator_nextElement, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setMaxDepth", zend.AccPublic, zim_spl_RecursiveIteratorIterator_setMaxDepth, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("max_depth"),
	}),
	types.MakeZendFunctionEntryEx("getMaxDepth", zend.AccPublic, zim_spl_RecursiveIteratorIterator_getMaxDepth, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_RecursiveTreeIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveTreeIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("flags"),
		zend.MakeArgName("caching_it_flags"),
		zend.MakeArgName("mode"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_RecursiveIteratorIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_RecursiveIteratorIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_RecursiveTreeIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_RecursiveTreeIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_RecursiveIteratorIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("beginIteration", zend.AccPublic, zim_spl_RecursiveIteratorIterator_beginIteration, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("endIteration", zend.AccPublic, zim_spl_RecursiveIteratorIterator_endIteration, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("callHasChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_callHasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("callGetChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_callGetChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("beginChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_beginChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("endChildren", zend.AccPublic, zim_spl_RecursiveIteratorIterator_endChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("nextElement", zend.AccPublic, zim_spl_RecursiveIteratorIterator_nextElement, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getPrefix", zend.AccPublic, zim_spl_RecursiveTreeIterator_getPrefix, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setPrefixPart", zend.AccPublic, zim_spl_RecursiveTreeIterator_setPrefixPart, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgName("part"),
		zend.MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("getEntry", zend.AccPublic, zim_spl_RecursiveTreeIterator_getEntry, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setPostfix", zend.AccPublic, zim_spl_RecursiveTreeIterator_setPostfix, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("postfix"),
	}),
	types.MakeZendFunctionEntryEx("getPostfix", zend.AccPublic, zim_spl_RecursiveTreeIterator_getPostfix, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_FilterIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_FilterIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_FilterIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, ZimSplDualItValid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, ZimSplDualItKey, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, ZimSplDualItCurrent, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_FilterIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("accept", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_CallbackFilterIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_CallbackFilterIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("callback"),
	}),
	types.MakeZendFunctionEntryEx("accept", zend.AccPublic, zim_spl_CallbackFilterIterator_accept, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_RecursiveCallbackFilterIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveCallbackFilterIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
		zend.MakeArgName("callback"),
	}),
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_RecursiveFilterIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_RecursiveCallbackFilterIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_RecursiveFilterIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveFilterIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_RecursiveFilterIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_RecursiveFilterIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_ParentIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_ParentIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("accept", zend.AccPublic, zim_spl_RecursiveFilterIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_RegexIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RegexIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("regex"),
		zend.MakeArgName("mode"),
		zend.MakeArgName("flags"),
		zend.MakeArgName("preg_flags"),
	}),
	types.MakeZendFunctionEntryEx("accept", zend.AccPublic, zim_spl_RegexIterator_accept, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getMode", zend.AccPublic, zim_spl_RegexIterator_getMode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setMode", zend.AccPublic, zim_spl_RegexIterator_setMode, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("mode"),
	}),
	types.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_RegexIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_RegexIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("getPregFlags", zend.AccPublic, zim_spl_RegexIterator_getPregFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setPregFlags", zend.AccPublic, zim_spl_RegexIterator_setPregFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("preg_flags"),
	}),
	types.MakeZendFunctionEntryEx("getRegex", zend.AccPublic, zim_spl_RegexIterator_getRegex, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_RecursiveRegexIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveRegexIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
		zend.MakeArgName("regex"),
		zend.MakeArgName("mode"),
		zend.MakeArgName("flags"),
		zend.MakeArgName("preg_flags"),
	}),
	types.MakeZendFunctionEntryEx("accept", zend.AccPublic, zim_spl_RecursiveRegexIterator_accept, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_RecursiveFilterIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_RecursiveRegexIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_SeekableIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("seek", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("position", zend.ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_LONG, 0))),
	}),
}
var spl_funcs_LimitIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_LimitIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("offset"),
		zend.MakeArgName("count"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_LimitIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_LimitIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, ZimSplDualItKey, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, ZimSplDualItCurrent, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_LimitIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("seek", zend.AccPublic, zim_spl_LimitIterator_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("position"),
	}),
	types.MakeZendFunctionEntryEx("getPosition", zend.AccPublic, zim_spl_LimitIterator_getPosition, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_CachingIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_CachingIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_CachingIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_CachingIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, ZimSplDualItKey, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, ZimSplDualItCurrent, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_CachingIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("hasNext", zend.AccPublic, zim_spl_CachingIterator_hasNext, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__toString", zend.AccPublic, zim_spl_CachingIterator___toString, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_CachingIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_CachingIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("offsetGet", zend.AccPublic, zim_spl_CachingIterator_offsetGet, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetSet", zend.AccPublic, zim_spl_CachingIterator_offsetSet, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("index"),
		zend.MakeArgName("newval"),
	}),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.AccPublic, zim_spl_CachingIterator_offsetUnset, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("offsetExists", zend.AccPublic, zim_spl_CachingIterator_offsetExists, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("index"),
	}),
	types.MakeZendFunctionEntryEx("getCache", zend.AccPublic, zim_spl_CachingIterator_getCache, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_CachingIterator_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_RecursiveCachingIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveCachingIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_RecursiveCachingIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_RecursiveCachingIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_IteratorIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_IteratorIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, ZimSplDualItRewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, ZimSplDualItValid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, ZimSplDualItKey, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, ZimSplDualItCurrent, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, ZimSplDualItNext, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_NoRewindIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_NoRewindIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_NoRewindIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_NoRewindIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_NoRewindIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_NoRewindIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_NoRewindIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_InfiniteIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_InfiniteIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_InfiniteIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_EmptyIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_EmptyIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_EmptyIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_EmptyIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_EmptyIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_EmptyIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_funcs_AppendIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_AppendIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("append", zend.AccPublic, zim_spl_AppendIterator_append, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	}),
	types.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_AppendIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_AppendIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", zend.AccPublic, ZimSplDualItKey, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_AppendIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_AppendIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic, zim_spl_dual_it_getInnerIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getIteratorIndex", zend.AccPublic, zim_spl_AppendIterator_getIteratorIndex, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getArrayIterator", zend.AccPublic, zim_spl_AppendIterator_getArrayIterator, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_funcs_OuterIterator = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.AccPublic|zend.AccAbstract, nil, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

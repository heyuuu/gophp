// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

const spl_ce_Traversable = zend.ZendCeTraversable
const spl_ce_Iterator *zend.ZendClassEntry = zend.ZendCeIterator
const spl_ce_Aggregate = zend.ZendCeAggregate
const spl_ce_ArrayAccess = zend.ZendCeArrayaccess
const spl_ce_Serializable = zend.ZendCeSerializable
const spl_ce_Countable = zend.ZendCeCountable

var spl_ce_RecursiveIterator *zend.ZendClassEntry
var spl_ce_RecursiveIteratorIterator *zend.ZendClassEntry
var spl_ce_RecursiveTreeIterator *zend.ZendClassEntry
var spl_ce_FilterIterator *zend.ZendClassEntry
var spl_ce_RecursiveFilterIterator *zend.ZendClassEntry
var spl_ce_ParentIterator *zend.ZendClassEntry
var spl_ce_SeekableIterator *zend.ZendClassEntry
var spl_ce_LimitIterator *zend.ZendClassEntry
var spl_ce_CachingIterator *zend.ZendClassEntry
var spl_ce_RecursiveCachingIterator *zend.ZendClassEntry
var spl_ce_OuterIterator *zend.ZendClassEntry
var spl_ce_IteratorIterator *zend.ZendClassEntry
var spl_ce_NoRewindIterator *zend.ZendClassEntry
var spl_ce_InfiniteIterator *zend.ZendClassEntry
var spl_ce_EmptyIterator *zend.ZendClassEntry
var spl_ce_AppendIterator *zend.ZendClassEntry
var spl_ce_RegexIterator *zend.ZendClassEntry
var spl_ce_RecursiveRegexIterator *zend.ZendClassEntry
var spl_ce_CallbackFilterIterator *zend.ZendClassEntry
var spl_ce_RecursiveCallbackFilterIterator *zend.ZendClassEntry

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

type __struct___spl_cbfilter_it_intern = _spl_cbfilter_it_intern

type SplIteratorApplyFuncT func(iter *zend.ZendObjectIterator, puser any) int

var ArginfoRecursiveItVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var spl_funcs_RecursiveIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"hasChildren",
		nil,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{
		"getChildren",
		nil,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
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
var SplRecursiveItIteratorFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplRecursiveItDtor, SplRecursiveItValid, SplRecursiveItGetCurrentData, SplRecursiveItGetCurrentKey, SplRecursiveItMoveForward, SplRecursiveItRewind, nil}
var ArginfoRecursiveItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0),
		0,
		0,
	},
	{"mode", 0, 0, 0},
	{"flags", 0, 0, 0},
}
var arginfo_recursive_it_getSubIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"level", 0, 0, 0}}
var arginfo_recursive_it_setMaxDepth []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"max_depth", 0, 0, 0}}
var spl_funcs_RecursiveIteratorIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveIteratorIterator___construct,
		ArginfoRecursiveItConstruct,
		uint32_t(b.SizeOf("arginfo_recursive_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_RecursiveIteratorIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_RecursiveIteratorIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_RecursiveIteratorIterator_key,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_RecursiveIteratorIterator_current,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_RecursiveIteratorIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getDepth",
		zim_spl_RecursiveIteratorIterator_getDepth,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getSubIterator",
		zim_spl_RecursiveIteratorIterator_getSubIterator,
		arginfo_recursive_it_getSubIterator,
		uint32_t(b.SizeOf("arginfo_recursive_it_getSubIterator")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_RecursiveIteratorIterator_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"beginIteration",
		zim_spl_RecursiveIteratorIterator_beginIteration,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"endIteration",
		zim_spl_RecursiveIteratorIterator_endIteration,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"callHasChildren",
		zim_spl_RecursiveIteratorIterator_callHasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"callGetChildren",
		zim_spl_RecursiveIteratorIterator_callGetChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"beginChildren",
		zim_spl_RecursiveIteratorIterator_beginChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"endChildren",
		zim_spl_RecursiveIteratorIterator_endChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"nextElement",
		zim_spl_RecursiveIteratorIterator_nextElement,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setMaxDepth",
		zim_spl_RecursiveIteratorIterator_setMaxDepth,
		arginfo_recursive_it_setMaxDepth,
		uint32_t(b.SizeOf("arginfo_recursive_it_setMaxDepth")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getMaxDepth",
		zim_spl_RecursiveIteratorIterator_getMaxDepth,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRecursiveTreeItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0),
		0,
		0,
	},
	{"flags", 0, 0, 0},
	{"caching_it_flags", 0, 0, 0},
	{"mode", 0, 0, 0},
}
var arginfo_recursive_tree_it_setPrefixPart []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"part", 0, 0, 0}, {"value", 0, 0, 0}}
var arginfo_recursive_tree_it_setPostfix []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"postfix", 0, 0, 0}}
var spl_funcs_RecursiveTreeIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveTreeIterator___construct,
		ArginfoRecursiveTreeItConstruct,
		uint32_t(b.SizeOf("arginfo_recursive_tree_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_RecursiveIteratorIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_RecursiveIteratorIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_RecursiveTreeIterator_key,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_RecursiveTreeIterator_current,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_RecursiveIteratorIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"beginIteration",
		zim_spl_RecursiveIteratorIterator_beginIteration,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"endIteration",
		zim_spl_RecursiveIteratorIterator_endIteration,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"callHasChildren",
		zim_spl_RecursiveIteratorIterator_callHasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"callGetChildren",
		zim_spl_RecursiveIteratorIterator_callGetChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"beginChildren",
		zim_spl_RecursiveIteratorIterator_beginChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"endChildren",
		zim_spl_RecursiveIteratorIterator_endChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"nextElement",
		zim_spl_RecursiveIteratorIterator_nextElement,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPrefix",
		zim_spl_RecursiveTreeIterator_getPrefix,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setPrefixPart",
		zim_spl_RecursiveTreeIterator_setPrefixPart,
		arginfo_recursive_tree_it_setPrefixPart,
		uint32_t(b.SizeOf("arginfo_recursive_tree_it_setPrefixPart")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getEntry",
		zim_spl_RecursiveTreeIterator_getEntry,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setPostfix",
		zim_spl_RecursiveTreeIterator_setPostfix,
		arginfo_recursive_tree_it_setPostfix,
		uint32_t(b.SizeOf("arginfo_recursive_tree_it_setPostfix")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPostfix",
		zim_spl_RecursiveTreeIterator_getPostfix,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
}
var spl_funcs_FilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_FilterIterator___construct,
		ArginfoFilterItConstruct,
		uint32_t(b.SizeOf("arginfo_filter_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_FilterIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		ZimSplDualItValid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_FilterIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"accept",
		nil,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCallbackFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"callback", 0, 0, 0},
}
var spl_funcs_CallbackFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_CallbackFilterIterator___construct,
		ArginfoCallbackFilterItConstruct,
		uint32_t(b.SizeOf("arginfo_callback_filter_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"accept",
		zim_spl_CallbackFilterIterator_accept,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRecursiveCallbackFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0),
		0,
		0,
	},
	{"callback", 0, 0, 0},
}
var spl_funcs_RecursiveCallbackFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveCallbackFilterIterator___construct,
		ArginfoRecursiveCallbackFilterItConstruct,
		uint32_t(b.SizeOf("arginfo_recursive_callback_filter_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_RecursiveCallbackFilterIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoParentItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0),
		0,
		0,
	},
}
var spl_funcs_RecursiveFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveFilterIterator___construct,
		ArginfoParentItConstruct,
		uint32_t(b.SizeOf("arginfo_parent_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_RecursiveFilterIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_ParentIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_ParentIterator___construct,
		ArginfoParentItConstruct,
		uint32_t(b.SizeOf("arginfo_parent_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"accept",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRegexItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(2)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"regex", 0, 0, 0},
	{"mode", 0, 0, 0},
	{"flags", 0, 0, 0},
	{"preg_flags", 0, 0, 0},
}
var ArginfoRegexItSetMode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoRegexItSetFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoRegexItSetPregFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"preg_flags", 0, 0, 0}}
var spl_funcs_RegexIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RegexIterator___construct,
		ArginfoRegexItConstruct,
		uint32_t(b.SizeOf("arginfo_regex_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"accept",
		zim_spl_RegexIterator_accept,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getMode",
		zim_spl_RegexIterator_getMode,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setMode",
		zim_spl_RegexIterator_setMode,
		ArginfoRegexItSetMode,
		uint32_t(b.SizeOf("arginfo_regex_it_set_mode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_RegexIterator_getFlags,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_RegexIterator_setFlags,
		ArginfoRegexItSetFlags,
		uint32_t(b.SizeOf("arginfo_regex_it_set_flags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPregFlags",
		zim_spl_RegexIterator_getPregFlags,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setPregFlags",
		zim_spl_RegexIterator_setPregFlags,
		ArginfoRegexItSetPregFlags,
		uint32_t(b.SizeOf("arginfo_regex_it_set_preg_flags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getRegex",
		zim_spl_RegexIterator_getRegex,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRecRegexItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(2)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0),
		0,
		0,
	},
	{"regex", 0, 0, 0},
	{"mode", 0, 0, 0},
	{"flags", 0, 0, 0},
	{"preg_flags", 0, 0, 0},
}
var spl_funcs_RecursiveRegexIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveRegexIterator___construct,
		ArginfoRecRegexItConstruct,
		uint32_t(b.SizeOf("arginfo_rec_regex_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"accept",
		zim_spl_RecursiveRegexIterator_accept,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_RecursiveRegexIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoSeekableItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"position", zend.ZEND_TYPE_ENCODE(zend.IS_LONG, 0), 0, 0},
}
var spl_funcs_SeekableIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"seek",
		nil,
		ArginfoSeekableItSeek,
		uint32_t(b.SizeOf("arginfo_seekable_it_seek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoLimitItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"offset", 0, 0, 0},
	{"count", 0, 0, 0},
}
var ArginfoLimitItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"position", 0, 0, 0},
}
var spl_funcs_LimitIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_LimitIterator___construct,
		ArginfoLimitItConstruct,
		uint32_t(b.SizeOf("arginfo_limit_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_LimitIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_LimitIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_LimitIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"seek",
		zim_spl_LimitIterator_seek,
		ArginfoLimitItSeek,
		uint32_t(b.SizeOf("arginfo_limit_it_seek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPosition",
		zim_spl_LimitIterator_getPosition,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCachingItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"flags", 0, 0, 0},
}
var arginfo_caching_it_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"flags", 0, 0, 0},
}
var arginfo_caching_it_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"index", 0, 0, 0},
}
var arginfo_caching_it_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"index", 0, 0, 0},
	{"newval", 0, 0, 0},
}
var spl_funcs_CachingIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_CachingIterator___construct,
		ArginfoCachingItConstruct,
		uint32_t(b.SizeOf("arginfo_caching_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_CachingIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_CachingIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_CachingIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasNext",
		zim_spl_CachingIterator_hasNext,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__toString",
		zim_spl_CachingIterator___toString,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_CachingIterator_getFlags,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_CachingIterator_setFlags,
		arginfo_caching_it_setFlags,
		uint32_t(b.SizeOf("arginfo_caching_it_setFlags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetGet",
		zim_spl_CachingIterator_offsetGet,
		arginfo_caching_it_offsetGet,
		uint32_t(b.SizeOf("arginfo_caching_it_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetSet",
		zim_spl_CachingIterator_offsetSet,
		arginfo_caching_it_offsetSet,
		uint32_t(b.SizeOf("arginfo_caching_it_offsetSet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetUnset",
		zim_spl_CachingIterator_offsetUnset,
		arginfo_caching_it_offsetGet,
		uint32_t(b.SizeOf("arginfo_caching_it_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"offsetExists",
		zim_spl_CachingIterator_offsetExists,
		arginfo_caching_it_offsetGet,
		uint32_t(b.SizeOf("arginfo_caching_it_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getCache",
		zim_spl_CachingIterator_getCache,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_CachingIterator_count,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCachingRecItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"flags", 0, 0, 0},
}
var spl_funcs_RecursiveCachingIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveCachingIterator___construct,
		ArginfoCachingRecItConstruct,
		uint32_t(b.SizeOf("arginfo_caching_rec_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_RecursiveCachingIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_RecursiveCachingIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoIteratorItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0),
		0,
		0,
	},
}
var spl_funcs_IteratorIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_IteratorIterator___construct,
		ArginfoIteratorItConstruct,
		uint32_t(b.SizeOf("arginfo_iterator_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		ZimSplDualItRewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		ZimSplDualItValid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		ZimSplDualItNext,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoNorewindItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
}
var spl_funcs_NoRewindIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_NoRewindIterator___construct,
		ArginfoNorewindItConstruct,
		uint32_t(b.SizeOf("arginfo_norewind_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_NoRewindIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_NoRewindIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_NoRewindIterator_key,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_NoRewindIterator_current,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_NoRewindIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_InfiniteIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_InfiniteIterator___construct,
		ArginfoNorewindItConstruct,
		uint32_t(b.SizeOf("arginfo_norewind_it___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_InfiniteIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_EmptyIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"rewind",
		zim_spl_EmptyIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_EmptyIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_EmptyIterator_key,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_EmptyIterator_current,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_EmptyIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoAppendItAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
}
var spl_funcs_AppendIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_AppendIterator___construct,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"append",
		zim_spl_AppendIterator_append,
		ArginfoAppendItAppend,
		uint32_t(b.SizeOf("arginfo_append_it_append")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_AppendIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_AppendIterator_valid,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_AppendIterator_current,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_AppendIterator_next,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getIteratorIndex",
		zim_spl_AppendIterator_getIteratorIndex,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getArrayIterator",
		zim_spl_AppendIterator_getArrayIterator,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_OuterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"getInnerIterator",
		nil,
		ArginfoRecursiveItVoid,
		uint32_t(b.SizeOf("arginfo_recursive_it_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}

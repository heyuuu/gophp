// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
)

// Source: <ext/spl/spl_iterators.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/pcre/php_pcre.h"

const spl_ce_Traversable = zend.ZendCeTraversable
const spl_ce_Iterator *types.ClassEntry = zend.ZendCeIterator
const spl_ce_Aggregate = zend.ZendCeAggregate
const spl_ce_ArrayAccess = zend.ZendCeArrayaccess
const spl_ce_Serializable = zend.ZendCeSerializable
const spl_ce_Countable = zend.ZendCeCountable

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

type __struct___spl_cbfilter_it_intern = _spl_cbfilter_it_intern

type SplIteratorApplyFuncT func(iter *zend.ZendObjectIterator, puser any) int

// Source: <ext/spl/spl_iterators.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

var ArginfoRecursiveItVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var spl_funcs_RecursiveIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, ArginfoRecursiveItVoid),
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
var SplRecursiveItIteratorFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplRecursiveItDtor, SplRecursiveItValid, SplRecursiveItGetCurrentData, SplRecursiveItGetCurrentKey, SplRecursiveItMoveForward, SplRecursiveItRewind, nil)

/* {{{ proto RecursiveIteratorIterator::__construct(RecursiveIterator|IteratorAggregate it [, int mode = RIT_LEAVES_ONLY [, int flags = 0]]) throws InvalidArgumentException
   Creates a RecursiveIteratorIterator from a RecursiveIterator. */

/* {{{ proto void RecursiveIteratorIterator::rewind()
   Rewind the iterator to the first element of the top level inner iterator. */

/* {{{ proto bool RecursiveIteratorIterator::valid()
   Check whether the current position is valid */

/* {{{ proto mixed RecursiveIteratorIterator::key()
   Access the current key */

/* {{{ proto mixed RecursiveIteratorIterator::current()
   Access the current element value */

/* {{{ proto void RecursiveIteratorIterator::next()
   Move forward to the next element */

/* {{{ proto int RecursiveIteratorIterator::getDepth()
   Get the current depth of the recursive iteration */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getSubIterator([int level])
   The current active sub iterator or the iterator at specified level */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getInnerIterator()
   The current active sub iterator */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::beginIteration()
   Called when iteration begins (after first rewind() call) */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::endIteration()
   Called when iteration ends (when valid() first returns false */

/* {{{ proto bool RecursiveIteratorIterator::callHasChildren()
   Called for each element to test whether it has children */

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::callGetChildren()
   Return children of current element */

/* {{{ proto void RecursiveIteratorIterator::beginChildren()
   Called when recursing one level down */

/* {{{ proto void RecursiveIteratorIterator::endChildren()
   Called when end recursing one level */

/* {{{ proto void RecursiveIteratorIterator::nextElement()
   Called when the next element is available */

/* {{{ proto void RecursiveIteratorIterator::setMaxDepth([$max_depth = -1])
   Set the maximum allowed depth (or any depth if pmax_depth = -1] */

/* {{{ proto int|false RecursiveIteratorIterator::getMaxDepth()
   Return the maximum accepted depth or false if any depth is allowed */

/* {{{ spl_RecursiveIteratorIterator_dtor */

var ArginfoRecursiveItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	zend.MakeArgInfo("mode"),
	zend.MakeArgInfo("flags"),
}
var arginfo_recursive_it_getSubIterator []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("level"),
}
var arginfo_recursive_it_setMaxDepth []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("max_depth"),
}
var spl_funcs_RecursiveIteratorIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator___construct, ArginfoRecursiveItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_key, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_current, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getDepth", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_getDepth, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getSubIterator", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_getSubIterator, arginfo_recursive_it_getSubIterator),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_getInnerIterator, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("beginIteration", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_beginIteration, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("endIteration", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_endIteration, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("callHasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_callHasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("callGetChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_callGetChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("beginChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_beginChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("endChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_endChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("nextElement", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_nextElement, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setMaxDepth", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_setMaxDepth, arginfo_recursive_it_setMaxDepth),
	types.MakeZendFunctionEntryEx("getMaxDepth", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_getMaxDepth, ArginfoRecursiveItVoid),
}

/* {{{ proto RecursiveTreeIterator::__construct(RecursiveIterator|IteratorAggregate it [, int flags = RTIT_BYPASS_KEY [, int cit_flags = CIT_CATCH_GET_CHILD [, mode = RIT_SELF_FIRST ]]]) throws InvalidArgumentException
   RecursiveIteratorIterator to generate ASCII graphic trees for the entries in a RecursiveIterator */

/* {{{ proto void RecursiveTreeIterator::setPrefixPart(int part, string prefix) throws OutOfRangeException
   Sets prefix parts as used in getPrefix() */

/* {{{ proto string RecursiveTreeIterator::getPrefix()
   Returns the string to place in front of current element */

/* {{{ proto void RecursiveTreeIterator::setPostfix(string prefix)
   Sets postfix as used in getPostfix() */

/* {{{ proto string RecursiveTreeIterator::getEntry()
   Returns the string presentation built for current element */

/* {{{ proto string RecursiveTreeIterator::getPostfix()
   Returns the string to place after the current element */

/* {{{ proto mixed RecursiveTreeIterator::current()
   Returns the current element prefixed and postfixed */

/* {{{ proto mixed RecursiveTreeIterator::key()
   Returns the current key prefixed and postfixed */

var ArginfoRecursiveTreeItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("caching_it_flags"),
	zend.MakeArgInfo("mode"),
}
var arginfo_recursive_tree_it_setPrefixPart []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("part"),
	zend.MakeArgInfo("value"),
}
var arginfo_recursive_tree_it_setPostfix []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("postfix"),
}
var spl_funcs_RecursiveTreeIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator___construct, ArginfoRecursiveTreeItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_key, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_current, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("beginIteration", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_beginIteration, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("endIteration", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_endIteration, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("callHasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_callHasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("callGetChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_callGetChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("beginChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_beginChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("endChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_endChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("nextElement", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveIteratorIterator_nextElement, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getPrefix", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_getPrefix, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setPrefixPart", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_setPrefixPart, arginfo_recursive_tree_it_setPrefixPart),
	types.MakeZendFunctionEntryEx("getEntry", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_getEntry, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setPostfix", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_setPostfix, arginfo_recursive_tree_it_setPostfix),
	types.MakeZendFunctionEntryEx("getPostfix", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveTreeIterator_getPostfix, ArginfoRecursiveItVoid),
}

/* {{{ proto FilterIterator::__construct(Iterator it)
   Create an Iterator from another iterator */

/* {{{ proto CallbackFilterIterator::__construct(Iterator it, callback func)
   Create an Iterator from another iterator */

/* {{{ proto Iterator FilterIterator::getInnerIterator()
    proto Iterator CachingIterator::getInnerIterator()
    proto Iterator LimitIterator::getInnerIterator()
    proto Iterator ParentIterator::getInnerIterator()
Get the inner iterator */

/* {{{ proto void ParentIterator::rewind()
       proto void IteratorIterator::rewind()
   Rewind the iterator
*/

/* {{{ proto bool FilterIterator::valid()
    proto bool ParentIterator::valid()
    proto bool IteratorIterator::valid()
    proto bool NoRewindIterator::valid()
Check whether the current element is valid */

/* {{{ proto mixed FilterIterator::key()
    proto mixed CachingIterator::key()
    proto mixed LimitIterator::key()
    proto mixed ParentIterator::key()
    proto mixed IteratorIterator::key()
    proto mixed NoRewindIterator::key()
    proto mixed AppendIterator::key()
Get the current key */

/* {{{ proto mixed FilterIterator::current()
    proto mixed CachingIterator::current()
    proto mixed LimitIterator::current()
    proto mixed ParentIterator::current()
    proto mixed IteratorIterator::current()
    proto mixed NoRewindIterator::current()
Get the current element value */

/* {{{ proto void ParentIterator::next()
    proto void IteratorIterator::next()
    proto void NoRewindIterator::next()
Move the iterator forward */

/* {{{ proto void FilterIterator::rewind()
   Rewind the iterator */

/* {{{ proto void FilterIterator::next()
   Move the iterator forward */

/* {{{ proto RecursiveCallbackFilterIterator::__construct(RecursiveIterator it, callback func)
   Create a RecursiveCallbackFilterIterator from a RecursiveIterator */

/* {{{ proto RecursiveFilterIterator::__construct(RecursiveIterator it)
   Create a RecursiveFilterIterator from a RecursiveIterator */

/* {{{ proto bool RecursiveFilterIterator::hasChildren()
   Check whether the inner iterator's current element has children */

/* {{{ proto RecursiveFilterIterator RecursiveFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveFilterIterator */

/* {{{ proto RecursiveCallbackFilterIterator RecursiveCallbackFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveCallbackFilterIterator */

/* {{{ proto ParentIterator::__construct(RecursiveIterator it)
   Create a ParentIterator from a RecursiveIterator */

/* {{{ proto RegexIterator::__construct(Iterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RegexIterator from another iterator and a regular expression */

/* {{{ proto bool CallbackFilterIterator::accept()
   Calls the callback with the current value, the current key and the inner iterator as arguments */

/* {{{ proto string RegexIterator::getRegex()
   Returns current regular expression */

/* {{{ proto bool RegexIterator::getMode()
   Returns current operation mode */

/* {{{ proto bool RegexIterator::setMode(int new_mode)
   Set new operation mode */

/* {{{ proto bool RegexIterator::getFlags()
   Returns current operation flags */

/* {{{ proto bool RegexIterator::setFlags(int new_flags)
   Set operation flags */

/* {{{ proto bool RegexIterator::getFlags()
   Returns current PREG flags (if in use or NULL) */

/* {{{ proto bool RegexIterator::setPregFlags(int new_flags)
   Set PREG flags */

/* {{{ proto RecursiveRegexIterator::__construct(RecursiveIterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RecursiveRegexIterator from another recursive iterator and a regular expression */

/* {{{ proto RecursiveRegexIterator RecursiveRegexIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveRegexIterator */

/* {{{ spl_dual_it_dtor */

var ArginfoFilterItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
}
var spl_funcs_FilterIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_FilterIterator___construct, ArginfoFilterItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_FilterIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, ZimSplDualItValid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, ZimSplDualItKey, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, ZimSplDualItCurrent, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_FilterIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("accept", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, ArginfoRecursiveItVoid),
}
var ArginfoCallbackFilterItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("callback"),
}
var spl_funcs_CallbackFilterIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_CallbackFilterIterator___construct, ArginfoCallbackFilterItConstruct),
	types.MakeZendFunctionEntryEx("accept", zend.ZEND_ACC_PUBLIC, zim_spl_CallbackFilterIterator_accept, ArginfoRecursiveItVoid),
}
var ArginfoRecursiveCallbackFilterItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
	zend.MakeArgInfo("callback"),
}
var spl_funcs_RecursiveCallbackFilterIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveCallbackFilterIterator___construct, ArginfoRecursiveCallbackFilterItConstruct),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator_hasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveCallbackFilterIterator_getChildren, ArginfoRecursiveItVoid),
}
var ArginfoParentItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
}
var spl_funcs_RecursiveFilterIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator___construct, ArginfoParentItConstruct),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator_hasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator_getChildren, ArginfoRecursiveItVoid),
}
var spl_funcs_ParentIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_ParentIterator___construct, ArginfoParentItConstruct),
	types.MakeZendFunctionEntryEx("accept", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator_hasChildren, ArginfoRecursiveItVoid),
}
var ArginfoRegexItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("regex"),
	zend.MakeArgInfo("mode"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("preg_flags"),
}
var ArginfoRegexItSetMode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("mode"),
}
var ArginfoRegexItSetFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("flags"),
}
var ArginfoRegexItSetPregFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("preg_flags"),
}
var spl_funcs_RegexIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator___construct, ArginfoRegexItConstruct),
	types.MakeZendFunctionEntryEx("accept", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_accept, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getMode", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_getMode, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setMode", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_setMode, ArginfoRegexItSetMode),
	types.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_getFlags, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_setFlags, ArginfoRegexItSetFlags),
	types.MakeZendFunctionEntryEx("getPregFlags", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_getPregFlags, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setPregFlags", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_setPregFlags, ArginfoRegexItSetPregFlags),
	types.MakeZendFunctionEntryEx("getRegex", zend.ZEND_ACC_PUBLIC, zim_spl_RegexIterator_getRegex, ArginfoRecursiveItVoid),
}
var ArginfoRecRegexItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("RecursiveIterator", 0))),
	zend.MakeArgInfo("regex"),
	zend.MakeArgInfo("mode"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("preg_flags"),
}
var spl_funcs_RecursiveRegexIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveRegexIterator___construct, ArginfoRecRegexItConstruct),
	types.MakeZendFunctionEntryEx("accept", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveRegexIterator_accept, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveFilterIterator_hasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveRegexIterator_getChildren, ArginfoRecursiveItVoid),
}

/* {{{ proto LimitIterator::__construct(Iterator it [, int offset, int count])
   Construct a LimitIterator from an Iterator with a given starting offset and optionally a maximum count */

/* {{{ proto void LimitIterator::rewind()
   Rewind the iterator to the specified starting offset */

/* {{{ proto bool LimitIterator::valid()
   Check whether the current element is valid */

/* {{{ proto void LimitIterator::next()
   Move the iterator forward */

/* {{{ proto void LimitIterator::seek(int position)
   Seek to the given position */

/* {{{ proto int LimitIterator::getPosition()
   Return the current position */

var ArginfoSeekableItSeek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("position", ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_LONG, 0))),
}
var spl_funcs_SeekableIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("seek", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, ArginfoSeekableItSeek),
}
var ArginfoLimitItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("count"),
}
var ArginfoLimitItSeek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("position"),
}
var spl_funcs_LimitIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator___construct, ArginfoLimitItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, ZimSplDualItKey, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, ZimSplDualItCurrent, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("seek", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator_seek, ArginfoLimitItSeek),
	types.MakeZendFunctionEntryEx("getPosition", zend.ZEND_ACC_PUBLIC, zim_spl_LimitIterator_getPosition, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
}

/* {{{ proto CachingIterator::__construct(Iterator it [, flags = CIT_CALL_TOSTRING])
   Construct a CachingIterator from an Iterator */

/* {{{ proto void CachingIterator::rewind()
   Rewind the iterator */

/* {{{ proto bool CachingIterator::valid()
   Check whether the current element is valid */

/* {{{ proto void CachingIterator::next()
   Move the iterator forward */

/* {{{ proto bool CachingIterator::hasNext()
   Check whether the inner iterator has a valid next element */

/* {{{ proto string CachingIterator::__toString()
   Return the string representation of the current element */

/* {{{ proto void CachingIterator::offsetSet(mixed index, mixed newval)
   Set given index in cache */

var ArginfoCachingItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("flags"),
}
var arginfo_caching_it_setFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("flags"),
}
var arginfo_caching_it_offsetGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("index"),
}
var arginfo_caching_it_offsetSet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("index"),
	zend.MakeArgInfo("newval"),
}
var spl_funcs_CachingIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator___construct, ArginfoCachingItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, ZimSplDualItKey, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, ZimSplDualItCurrent, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("hasNext", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_hasNext, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("__toString", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator___toString, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_getFlags, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_setFlags, arginfo_caching_it_setFlags),
	types.MakeZendFunctionEntryEx("offsetGet", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_offsetGet, arginfo_caching_it_offsetGet),
	types.MakeZendFunctionEntryEx("offsetSet", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_offsetSet, arginfo_caching_it_offsetSet),
	types.MakeZendFunctionEntryEx("offsetUnset", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_offsetUnset, arginfo_caching_it_offsetGet),
	types.MakeZendFunctionEntryEx("offsetExists", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_offsetExists, arginfo_caching_it_offsetGet),
	types.MakeZendFunctionEntryEx("getCache", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_getCache, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_CachingIterator_count, ArginfoRecursiveItVoid),
}

/* {{{ proto RecursiveCachingIterator::__construct(RecursiveIterator it [, flags = CIT_CALL_TOSTRING])
   Create an iterator from a RecursiveIterator */

/* {{{ proto bool RecursiveCachingIterator::hasChildren()
   Check whether the current element of the inner iterator has children */

/* {{{ proto RecursiveCachingIterator RecursiveCachingIterator::getChildren()
Return the inner iterator's children as a RecursiveCachingIterator */

var ArginfoCachingRecItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("flags"),
}
var spl_funcs_RecursiveCachingIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveCachingIterator___construct, ArginfoCachingRecItConstruct),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveCachingIterator_hasChildren, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveCachingIterator_getChildren, ArginfoRecursiveItVoid),
}

/* {{{ proto IteratorIterator::__construct(Traversable it)
   Create an iterator from anything that is traversable */

var ArginfoIteratorItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
}
var spl_funcs_IteratorIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_IteratorIterator___construct, ArginfoIteratorItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, ZimSplDualItRewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, ZimSplDualItValid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, ZimSplDualItKey, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, ZimSplDualItCurrent, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, ZimSplDualItNext, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
}

/* {{{ proto NoRewindIterator::__construct(Iterator it)
   Create an iterator from another iterator */

/* {{{ proto void NoRewindIterator::rewind()
   Prevent a call to inner iterators rewind() */

/* {{{ proto bool NoRewindIterator::valid()
   Return inner iterators valid() */

/* {{{ proto mixed NoRewindIterator::key()
   Return inner iterators key() */

/* {{{ proto mixed NoRewindIterator::current()
   Return inner iterators current() */

/* {{{ proto void NoRewindIterator::next()
   Return inner iterators next() */

var ArginfoNorewindItConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
}
var spl_funcs_NoRewindIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator___construct, ArginfoNorewindItConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator_key, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator_current, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_NoRewindIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
}

/* {{{ proto InfiniteIterator::__construct(Iterator it)
   Create an iterator from another iterator */

/* {{{ proto void InfiniteIterator::next()
   Prevent a call to inner iterators rewind() (internally the current data will be fetched if valid()) */

var spl_funcs_InfiniteIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_InfiniteIterator___construct, ArginfoNorewindItConstruct),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_InfiniteIterator_next, ArginfoRecursiveItVoid),
}

/* {{{ proto void EmptyIterator::rewind()
   Does nothing  */

/* {{{ proto false EmptyIterator::valid()
   Return false */

/* {{{ proto void EmptyIterator::key()
   Throws exception BadMethodCallException */

/* {{{ proto void EmptyIterator::current()
   Throws exception BadMethodCallException */

/* {{{ proto void EmptyIterator::next()
   Does nothing */

var spl_funcs_EmptyIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_EmptyIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_EmptyIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_EmptyIterator_key, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_EmptyIterator_current, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_EmptyIterator_next, ArginfoRecursiveItVoid),
}

/* {{{ proto AppendIterator::__construct()
   Create an AppendIterator */

/* {{{ proto void AppendIterator::append(Iterator it)
   Append an iterator */

/* {{{ proto mixed AppendIterator::current()
   Get the current element value */

/* {{{ proto void AppendIterator::rewind()
   Rewind to the first iterator and rewind the first iterator, too */

/* {{{ proto bool AppendIterator::valid()
   Check if the current state is valid */

/* {{{ proto void AppendIterator::next()
   Forward to next element */

/* {{{ proto int AppendIterator::getIteratorIndex()
   Get index of iterator */

/* {{{ proto ArrayIterator AppendIterator::getArrayIterator()
   Get access to inner ArrayIterator */

var ArginfoAppendItAppend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
}
var spl_funcs_AppendIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator___construct, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("append", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_append, ArginfoAppendItAppend),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_rewind, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_valid, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, ZimSplDualItKey, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_current, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_next, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC, zim_spl_dual_it_getInnerIterator, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getIteratorIndex", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_getIteratorIndex, ArginfoRecursiveItVoid),
	types.MakeZendFunctionEntryEx("getArrayIterator", zend.ZEND_ACC_PUBLIC, zim_spl_AppendIterator_getArrayIterator, ArginfoRecursiveItVoid),
}
var spl_funcs_OuterIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("getInnerIterator", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, ArginfoRecursiveItVoid),
}

/* {{{ PHP_MINIT_FUNCTION(spl_iterators)
 */

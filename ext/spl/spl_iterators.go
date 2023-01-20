// <<generate>>

package spl

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
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

// #define SPL_ITERATORS_H

// # include "php.h"

// # include "php_spl.h"

// failed # include "ext/pcre/php_pcre.h"

// #define spl_ce_Traversable       zend_ce_traversable

// #define spl_ce_Iterator       zend_ce_iterator

// #define spl_ce_Aggregate       zend_ce_aggregate

// #define spl_ce_ArrayAccess       zend_ce_arrayaccess

// #define spl_ce_Serializable       zend_ce_serializable

// #define spl_ce_Countable       zend_ce_countable

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

// @type _spl_cbfilter_it_intern struct
type __struct___spl_cbfilter_it_intern = _spl_cbfilter_it_intern

// @type SplDualItObject struct

func SplDualItFromObj(obj *zend.ZendObject) *SplDualItObject {
	return (*SplDualItObject)((*byte)(obj - zend_long((*byte)(&((*SplDualItObject)(nil).GetStd()))-(*byte)(nil))))
}

// #define Z_SPLDUAL_IT_P(zv) spl_dual_it_from_obj ( Z_OBJ_P ( ( zv ) ) )

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

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_directory.h"

// # include "spl_array.h"

// # include "spl_exceptions.h"

// # include "zend_smart_str.h"

var ArginfoRecursiveItVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var spl_funcs_RecursiveIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"hasChildren",
		nil,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{
		"getChildren",
		nil,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}

type RecursiveIteratorMode = int

const (
	RIT_LEAVES_ONLY = 0
	RIT_SELF_FIRST  = 1
	RIT_CHILD_FIRST = 2
)

// #define RIT_CATCH_GET_CHILD       CIT_CATCH_GET_CHILD

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

// @type SplSubIterator struct

// @type SplRecursiveItObject struct

// @type SplRecursiveItIterator struct

var SplHandlersRecItIt zend.ZendObjectHandlers
var SplHandlersDualIt zend.ZendObjectHandlers

func SplRecursiveItFromObj(obj *zend.ZendObject) *SplRecursiveItObject {
	return (*SplRecursiveItObject)((*byte)(obj - zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLRECURSIVE_IT_P(zv) spl_recursive_it_from_obj ( Z_OBJ_P ( ( zv ) ) )

// #define SPL_FETCH_AND_CHECK_DUAL_IT(var,objzval) do { spl_dual_it_object * it = Z_SPLDUAL_IT_P ( objzval ) ; if ( it -> dit_type == DIT_Unknown ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = it ; } while ( 0 )

// #define SPL_FETCH_SUB_ELEMENT(var,object,element) do { if ( ! ( object ) -> iterators ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = ( object ) -> iterators [ ( object ) -> level ] . element ; } while ( 0 )

// #define SPL_FETCH_SUB_ELEMENT_ADDR(var,object,element) do { if ( ! ( object ) -> iterators ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = & ( object ) -> iterators [ ( object ) -> level ] . element ; } while ( 0 )

// #define SPL_FETCH_SUB_ITERATOR(var,object) SPL_FETCH_SUB_ELEMENT ( var , object , iterator )

func SplRecursiveItDtor(_iter *zend.ZendObjectIterator) {
	var iter *SplRecursiveItIterator = (*SplRecursiveItIterator)(_iter)
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&iter.intern.data.value.obj)
	var sub_iter *zend.ZendObjectIterator
	for object.GetLevel() > 0 {
		if object.GetIterators()[object.GetLevel()].zobject.u1.v.type_ != 0 {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&object.iterators[object.GetLevel()].GetZobject())
		}
		object.GetLevel()--
	}
	object.SetIterators(zend._erealloc(object.GetIterators(), g.SizeOf("spl_sub_iterator")))
	object.SetLevel(0)
	zend.ZvalPtrDtor(&iter.intern.data)
}
func SplRecursiveItValidEx(object *SplRecursiveItObject, zthis *zend.Zval) int {
	var sub_iter *zend.ZendObjectIterator
	var level int = object.GetLevel()
	if object.GetIterators() == nil {
		return zend.FAILURE
	}
	for level >= 0 {
		sub_iter = object.GetIterators()[level].GetIterator()
		if sub_iter.funcs.valid(sub_iter) == zend.SUCCESS {
			return zend.SUCCESS
		}
		level--
	}
	if object.GetEndIteration() != nil && object.GetInIteration() != 0 {
		zend.ZendCallMethod(zthis, object.GetCe(), &object.endIteration, "endIteration", g.SizeOf("\"endIteration\"")-1, nil, 0, nil, nil)
	}
	object.SetInIteration(0)
	return zend.FAILURE
}
func SplRecursiveItValid(iter *zend.ZendObjectIterator) int {
	return SplRecursiveItValidEx(SplRecursiveItFromObj(&iter.data.value.obj), &iter.data)
}
func SplRecursiveItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&iter.data.value.obj)
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	return sub_iter.funcs.get_current_data(sub_iter)
}
func SplRecursiveItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&iter.data.value.obj)
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	if sub_iter.funcs.get_current_key != nil {
		sub_iter.funcs.get_current_key(sub_iter, key)
	} else {
		var __z *zend.Zval = key
		__z.value.lval = iter.index
		__z.u1.type_info = 4
	}
}
func SplRecursiveItMoveForwardEx(object *SplRecursiveItObject, zthis *zend.Zval) {
	var iterator *zend.ZendObjectIterator
	var zobject *zend.Zval
	var ce *zend.ZendClassEntry
	var retval zend.Zval
	var child zend.Zval
	var sub_iter *zend.ZendObjectIterator
	var has_children int
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	iterator = object.GetIterators()[object.GetLevel()].GetIterator()
	for zend.EG.exception == nil {
	next_step:
		iterator = object.GetIterators()[object.GetLevel()].GetIterator()
		switch object.GetIterators()[object.GetLevel()].GetState() {
		case RS_NEXT:
			iterator.funcs.move_forward(iterator)
			if zend.EG.exception != nil {
				if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
				}
			}
		case RS_START:
			if iterator.funcs.valid(iterator) == zend.FAILURE {
				break
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_TEST)
		case RS_TEST:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = &object.iterators[object.GetLevel()].GetZobject()
			if object.GetCallHasChildren() != nil {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.callHasChildren, "callHasChildren", g.SizeOf("\"callHasChildren\"")-1, &retval, 0, nil, nil)
			} else {
				zend.ZendCallMethod(zobject, ce, nil, "haschildren", g.SizeOf("\"haschildren\"")-1, &retval, 0, nil, nil)
			}
			if zend.EG.exception != nil {
				if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					return
				} else {
					zend.ZendClearException()
				}
			}
			if retval.u1.v.type_ != 0 {
				has_children = zend.ZendIsTrue(&retval)
				zend.ZvalPtrDtor(&retval)
				if has_children != 0 {
					if object.GetMaxDepth() == -1 || object.GetMaxDepth() > object.GetLevel() {
						switch object.GetMode() {
						case RIT_LEAVES_ONLY:

						case RIT_CHILD_FIRST:
							object.GetIterators()[object.GetLevel()].SetState(RS_CHILD)
							goto next_step
						case RIT_SELF_FIRST:
							object.GetIterators()[object.GetLevel()].SetState(RS_SELF)
							goto next_step
						}
					} else {

						/* do not recurse into */

						if object.GetMode() == RIT_LEAVES_ONLY {

							/* this is not a leave, so skip it */

							object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
							goto next_step
						}

						/* do not recurse into */

					}
				}
			}
			if object.GetNextElement() != nil {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.nextElement, "nextelement", g.SizeOf("\"nextelement\"")-1, nil, 0, nil, nil)
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			if zend.EG.exception != nil {
				if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
				}
			}
			return
		case RS_SELF:
			if object.GetNextElement() != nil && (object.GetMode() == RIT_SELF_FIRST || object.GetMode() == RIT_CHILD_FIRST) {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.nextElement, "nextelement", g.SizeOf("\"nextelement\"")-1, nil, 0, nil, nil)
			}
			if object.GetMode() == RIT_SELF_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_CHILD)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			return
		case RS_CHILD:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = &object.iterators[object.GetLevel()].GetZobject()
			if object.GetCallGetChildren() != nil {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.callGetChildren, "callGetChildren", g.SizeOf("\"callGetChildren\"")-1, &child, 0, nil, nil)
			} else {
				zend.ZendCallMethod(zobject, ce, nil, "getchildren", g.SizeOf("\"getchildren\"")-1, &child, 0, nil, nil)
			}
			if zend.EG.exception != nil {
				if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
					zend.ZvalPtrDtor(&child)
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					goto next_step
				}
			}
			if child.u1.v.type_ == 0 || child.u1.v.type_ != 8 || !(g.Assign(&ce, child.value.obj.ce) && zend.InstanceofFunction(ce, spl_ce_RecursiveIterator) != 0) {
				zend.ZvalPtrDtor(&child)
				zend.ZendThrowException(spl_ce_UnexpectedValueException, "Objects returned by RecursiveIterator::getChildren() must implement RecursiveIterator", 0)
				return
			}
			if object.GetMode() == RIT_CHILD_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_SELF)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			object.SetIterators(zend._erealloc(object.GetIterators(), g.SizeOf("spl_sub_iterator")*(g.PreInc(&(object.GetLevel()))+1)))
			sub_iter = ce.get_iterator(ce, &child, 0)
			var _z1 *zend.Zval = &object.iterators[object.GetLevel()].GetZobject()
			var _z2 *zend.Zval = &child
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			object.GetIterators()[object.GetLevel()].SetIterator(sub_iter)
			object.GetIterators()[object.GetLevel()].SetCe(ce)
			object.GetIterators()[object.GetLevel()].SetState(RS_START)
			if sub_iter.funcs.rewind != nil {
				sub_iter.funcs.rewind(sub_iter)
			}
			if object.GetBeginChildren() != nil {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.beginChildren, "beginchildren", g.SizeOf("\"beginchildren\"")-1, nil, 0, nil, nil)
				if zend.EG.exception != nil {
					if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
						return
					} else {
						zend.ZendClearException()
					}
				}
			}
			goto next_step
		}

		/* no more elements */

		if object.GetLevel() > 0 {
			if object.GetEndChildren() != nil {
				zend.ZendCallMethod(zthis, object.GetCe(), &object.endChildren, "endchildren", g.SizeOf("\"endchildren\"")-1, nil, 0, nil, nil)
				if zend.EG.exception != nil {
					if (object.GetFlags() & CIT_CATCH_GET_CHILD) == 0 {
						return
					} else {
						zend.ZendClearException()
					}
				}
			}
			if object.GetLevel() > 0 {
				var garbage zend.Zval
				var _z1 *zend.Zval = &garbage
				var _z2 *zend.Zval = &object.iterators[object.GetLevel()].GetZobject()
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				&object.iterators[object.GetLevel()].GetZobject().u1.type_info = 0
				zend.ZvalPtrDtor(&garbage)
				zend.ZendIteratorDtor(iterator)
				object.GetLevel()--
			}
		} else {
			return
		}

		/* no more elements */

	}
}
func SplRecursiveItRewindEx(object *SplRecursiveItObject, zthis *zend.Zval) {
	var sub_iter *zend.ZendObjectIterator
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
	for object.GetLevel() != 0 {
		sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
		zend.ZendIteratorDtor(sub_iter)
		zend.ZvalPtrDtor(&object.iterators[g.PostDec(&(object.GetLevel()))].GetZobject())
		if zend.EG.exception == nil && (object.GetEndChildren() == nil || object.GetEndChildren().common.scope != spl_ce_RecursiveIteratorIterator) {
			zend.ZendCallMethod(zthis, object.GetCe(), &object.endChildren, "endchildren", g.SizeOf("\"endchildren\"")-1, nil, 0, nil, nil)
		}
	}
	object.SetIterators(zend._erealloc(object.GetIterators(), g.SizeOf("spl_sub_iterator")))
	object.GetIterators()[0].SetState(RS_START)
	sub_iter = object.GetIterators()[0].GetIterator()
	if sub_iter.funcs.rewind != nil {
		sub_iter.funcs.rewind(sub_iter)
	}
	if zend.EG.exception == nil && object.GetBeginIteration() != nil && object.GetInIteration() == 0 {
		zend.ZendCallMethod(zthis, object.GetCe(), &object.beginIteration, "beginIteration", g.SizeOf("\"beginIteration\"")-1, nil, 0, nil, nil)
	}
	object.SetInIteration(1)
	SplRecursiveItMoveForwardEx(object, zthis)
}
func SplRecursiveItMoveForward(iter *zend.ZendObjectIterator) {
	SplRecursiveItMoveForwardEx(SplRecursiveItFromObj(&iter.data.value.obj), &iter.data)
}
func SplRecursiveItRewind(iter *zend.ZendObjectIterator) {
	SplRecursiveItRewindEx(SplRecursiveItFromObj(&iter.data.value.obj), &iter.data)
}

var SplRecursiveItIteratorFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplRecursiveItDtor, SplRecursiveItValid, SplRecursiveItGetCurrentData, SplRecursiveItGetCurrentKey, SplRecursiveItMoveForward, SplRecursiveItRewind, nil}

func SplRecursiveItGetIterator(ce *zend.ZendClassEntry, zobject *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplRecursiveItIterator
	var object *SplRecursiveItObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend._emalloc(g.SizeOf("spl_recursive_it_iterator"))
	object = SplRecursiveItFromObj(zobject.value.obj)
	if object.GetIterators() == nil {
		zend.ZendError(1<<0, "The object to be iterated is in an invalid state: "+"the parent constructor has not been called")
	}
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.ZvalAddrefP(zobject)
	var __z *zend.Zval = &iterator.intern.data
	__z.value.obj = zobject.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.funcs = &SplRecursiveItIteratorFuncs
	return (*zend.ZendObjectIterator)(iterator)
}
func SplRecursiveItItConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, rit_type RecursiveItItType) {
	var object *zend.Zval = &(execute_data.This)
	var intern *SplRecursiveItObject
	var iterator *zend.Zval
	var ce_iterator *zend.ZendClassEntry
	var mode zend.ZendLong
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	var caching_it zend.Zval
	var aggregate_retval zend.Zval
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
	switch rit_type {
	case RIT_RecursiveTreeIterator:
		var caching_it_flags zend.Zval
		var user_caching_it_flags *zend.Zval = nil
		mode = RIT_SELF_FIRST
		flags = RTIT_BYPASS_KEY
		if zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "o|lzl", &iterator, &flags, &user_caching_it_flags, &mode) == zend.SUCCESS {
			if zend.InstanceofFunction(iterator.value.obj.ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethod(iterator, iterator.value.obj.ce, &(iterator.value.obj.ce).iterator_funcs_ptr.zf_new_iterator, "getiterator", g.SizeOf("\"getiterator\"")-1, &aggregate_retval, 0, nil, nil)
				iterator = &aggregate_retval
			} else {
				zend.ZvalAddrefP(iterator)
			}
			if user_caching_it_flags != nil {
				var _z1 *zend.Zval = &caching_it_flags
				var _z2 *zend.Zval = user_caching_it_flags
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				if (_t & 0xff00) != 0 {
					zend.ZendGcAddref(&_gc.gc)
				}
			} else {
				var __z *zend.Zval = &caching_it_flags
				__z.value.lval = CIT_CATCH_GET_CHILD
				__z.u1.type_info = 4
			}
			SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &caching_it, iterator, &caching_it_flags)
			zend.ZvalPtrDtor(&caching_it_flags)
			zend.ZvalPtrDtor(iterator)
			iterator = &caching_it
		} else {
			iterator = nil
		}
		break
	case RIT_RecursiveIteratorIterator:

	default:
		mode = RIT_LEAVES_ONLY
		flags = 0
		if zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "o|ll", &iterator, &mode, &flags) == zend.SUCCESS {
			if zend.InstanceofFunction(iterator.value.obj.ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethod(iterator, iterator.value.obj.ce, &(iterator.value.obj.ce).iterator_funcs_ptr.zf_new_iterator, "getiterator", g.SizeOf("\"getiterator\"")-1, &aggregate_retval, 0, nil, nil)
				iterator = &aggregate_retval
			} else {
				zend.ZvalAddrefP(iterator)
			}
		} else {
			iterator = nil
		}
		break
	}
	if iterator == nil || zend.InstanceofFunction(iterator.value.obj.ce, spl_ce_RecursiveIterator) == 0 {
		if iterator != nil {
			zend.ZvalPtrDtor(iterator)
		}
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "An instance of RecursiveIterator or IteratorAggregate creating it is required", 0)
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	intern = SplRecursiveItFromObj(object.value.obj)
	intern.SetIterators(zend._emalloc(g.SizeOf("spl_sub_iterator")))
	intern.SetLevel(0)
	intern.SetMode(mode)
	intern.SetFlags(int(flags))
	intern.SetMaxDepth(-1)
	intern.SetInIteration(0)
	intern.SetCe(object.value.obj.ce)
	intern.SetBeginIteration(zend.ZendHashStrFindPtr(&intern.ce.function_table, "beginiteration", g.SizeOf("\"beginiteration\"")-1))
	if intern.GetBeginIteration().common.scope == ce_base {
		intern.SetBeginIteration(nil)
	}
	intern.SetEndIteration(zend.ZendHashStrFindPtr(&intern.ce.function_table, "enditeration", g.SizeOf("\"enditeration\"")-1))
	if intern.GetEndIteration().common.scope == ce_base {
		intern.SetEndIteration(nil)
	}
	intern.SetCallHasChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "callhaschildren", g.SizeOf("\"callHasChildren\"")-1))
	if intern.GetCallHasChildren().common.scope == ce_base {
		intern.SetCallHasChildren(nil)
	}
	intern.SetCallGetChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "callgetchildren", g.SizeOf("\"callGetChildren\"")-1))
	if intern.GetCallGetChildren().common.scope == ce_base {
		intern.SetCallGetChildren(nil)
	}
	intern.SetBeginChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "beginchildren", g.SizeOf("\"beginchildren\"")-1))
	if intern.GetBeginChildren().common.scope == ce_base {
		intern.SetBeginChildren(nil)
	}
	intern.SetEndChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "endchildren", g.SizeOf("\"endchildren\"")-1))
	if intern.GetEndChildren().common.scope == ce_base {
		intern.SetEndChildren(nil)
	}
	intern.SetNextElement(zend.ZendHashStrFindPtr(&intern.ce.function_table, "nextelement", g.SizeOf("\"nextElement\"")-1))
	if intern.GetNextElement().common.scope == ce_base {
		intern.SetNextElement(nil)
	}
	ce_iterator = iterator.value.obj.ce
	intern.GetIterators()[0].SetIterator(ce_iterator.get_iterator(ce_iterator, iterator, 0))
	var __z *zend.Zval = &intern.iterators[0].GetZobject()
	__z.value.obj = iterator.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	intern.GetIterators()[0].SetCe(ce_iterator)
	intern.GetIterators()[0].SetState(RS_START)
	zend.ZendRestoreErrorHandling(&error_handling)
	if zend.EG.exception != nil {
		var sub_iter *zend.ZendObjectIterator
		for intern.GetLevel() >= 0 {
			sub_iter = intern.GetIterators()[intern.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&intern.iterators[g.PostDec(&(intern.GetLevel()))].GetZobject())
		}
		zend._efree(intern.GetIterators())
		intern.SetIterators(nil)
	}
}

/* {{{ proto RecursiveIteratorIterator::__construct(RecursiveIterator|IteratorAggregate it [, int mode = RIT_LEAVES_ONLY [, int flags = 0]]) throws InvalidArgumentException
   Creates a RecursiveIteratorIterator from a RecursiveIterator. */

func zim_spl_RecursiveIteratorIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplRecursiveItItConstruct(execute_data, return_value, spl_ce_RecursiveIteratorIterator, zend.ZendCeIterator, RIT_RecursiveIteratorIterator)
}

/* {{{ proto void RecursiveIteratorIterator::rewind()
   Rewind the iterator to the first element of the top level inner iterator. */

func zim_spl_RecursiveIteratorIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplRecursiveItRewindEx(object, &(execute_data.This))
}

/* {{{ proto bool RecursiveIteratorIterator::valid()
   Check whether the current position is valid */

func zim_spl_RecursiveIteratorIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if SplRecursiveItValidEx(object, &(execute_data.This)) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed RecursiveIteratorIterator::key()
   Access the current key */

func zim_spl_RecursiveIteratorIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var iterator *zend.ZendObjectIterator
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	iterator = object.GetIterators()[object.GetLevel()].GetIterator()
	if iterator.funcs.get_current_key != nil {
		iterator.funcs.get_current_key(iterator, return_value)
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto mixed RecursiveIteratorIterator::current()
   Access the current element value */

func zim_spl_RecursiveIteratorIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var iterator *zend.ZendObjectIterator
	var data *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	iterator = object.GetIterators()[object.GetLevel()].GetIterator()
	data = iterator.funcs.get_current_data(iterator)
	if data != nil {
		var _z3 *zend.Zval = data
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* {{{ proto void RecursiveIteratorIterator::next()
   Move forward to the next element */

func zim_spl_RecursiveIteratorIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplRecursiveItMoveForwardEx(object, &(execute_data.This))
}

/* {{{ proto int RecursiveIteratorIterator::getDepth()
   Get the current depth of the recursive iteration */

func zim_spl_RecursiveIteratorIterator_getDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = object.GetLevel()
	__z.u1.type_info = 4
	return
}

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getSubIterator([int level])
   The current active sub iterator or the iterator at specified level */

func zim_spl_RecursiveIteratorIterator_getSubIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var level zend.ZendLong = object.GetLevel()
	var value *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|l", &level) == zend.FAILURE {
		return
	}
	if level < 0 || level > object.GetLevel() {
		return_value.u1.type_info = 1
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	value = &object.iterators[level].GetZobject()
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::getInnerIterator()
   The current active sub iterator */

func zim_spl_RecursiveIteratorIterator_getInnerIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var zobject *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	var _z3 *zend.Zval = zobject
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::beginIteration()
   Called when iteration begins (after first rewind() call) */

func zim_spl_RecursiveIteratorIterator_beginIteration(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::endIteration()
   Called when iteration ends (when valid() first returns false */

func zim_spl_RecursiveIteratorIterator_endIteration(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto bool RecursiveIteratorIterator::callHasChildren()
   Called for each element to test whether it has children */

func zim_spl_RecursiveIteratorIterator_callHasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry
	var zobject *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		return_value.u1.type_info = 1
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	if zobject.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
		return
	} else {
		zend.ZendCallMethod(zobject, ce, nil, "haschildren", g.SizeOf("\"haschildren\"")-1, return_value, 0, nil, nil)
		if return_value.u1.v.type_ == 0 {
			return_value.u1.type_info = 2
			return
		}
	}
}

/* {{{ proto RecursiveIterator RecursiveIteratorIterator::callGetChildren()
   Return children of current element */

func zim_spl_RecursiveIteratorIterator_callGetChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry
	var zobject *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	if zobject.u1.v.type_ == 0 {
		return
	} else {
		zend.ZendCallMethod(zobject, ce, nil, "getchildren", g.SizeOf("\"getchildren\"")-1, return_value, 0, nil, nil)
		if return_value.u1.v.type_ == 0 {
			return_value.u1.type_info = 1
			return
		}
	}
}

/* {{{ proto void RecursiveIteratorIterator::beginChildren()
   Called when recursing one level down */

func zim_spl_RecursiveIteratorIterator_beginChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto void RecursiveIteratorIterator::endChildren()
   Called when end recursing one level */

func zim_spl_RecursiveIteratorIterator_endChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto void RecursiveIteratorIterator::nextElement()
   Called when the next element is available */

func zim_spl_RecursiveIteratorIterator_nextElement(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto void RecursiveIteratorIterator::setMaxDepth([$max_depth = -1])
   Set the maximum allowed depth (or any depth if pmax_depth = -1] */

func zim_spl_RecursiveIteratorIterator_setMaxDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var max_depth zend.ZendLong = -1
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|l", &max_depth) == zend.FAILURE {
		return
	}
	if max_depth < -1 {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter max_depth must be >= -1", 0)
		return
	} else if max_depth > 2147483647 {
		max_depth = 2147483647
	}
	object.SetMaxDepth(int(max_depth))
}

/* {{{ proto int|false RecursiveIteratorIterator::getMaxDepth()
   Return the maximum accepted depth or false if any depth is allowed */

func zim_spl_RecursiveIteratorIterator_getMaxDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetMaxDepth() == -1 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = object.GetMaxDepth()
		__z.u1.type_info = 4
		return
	}
}
func SplRecursiveItGetMethod(zobject **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var object *SplRecursiveItObject = SplRecursiveItFromObj(*zobject)
	var level zend.ZendLong = object.GetLevel()
	var zobj *zend.Zval
	if object.GetIterators() == nil {
		core.PhpErrorDocref(nil, 1<<0, "The %s instance wasn't initialized properly", (*zobject).ce.name.val)
	}
	zobj = &object.iterators[level].GetZobject()
	function_handler = zend.ZendStdGetMethod(zobject, method, key)
	if function_handler == nil {
		if g.Assign(&function_handler, zend.ZendHashFindPtr(&(zobj.value.obj.ce).function_table, method)) == nil {
			*zobject = zobj.value.obj
			function_handler = (*zobject).handlers.get_method(zobject, method, key)
		} else {
			*zobject = zobj.value.obj
		}
	}
	return function_handler
}

/* {{{ spl_RecursiveIteratorIterator_dtor */

func spl_RecursiveIteratorIterator_dtor(_object *zend.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	var sub_iter *zend.ZendObjectIterator

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	if object.GetIterators() != nil {
		for object.GetLevel() >= 0 {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&object.iterators[g.PostDec(&(object.GetLevel()))].GetZobject())
		}
		zend._efree(object.GetIterators())
		object.SetIterators(nil)
	}
}

/* }}} */

func spl_RecursiveIteratorIterator_free_storage(_object *zend.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	if object.GetIterators() != nil {
		zend._efree(object.GetIterators())
		object.SetIterators(nil)
		object.SetLevel(0)
	}
	zend.ZendObjectStdDtor(&object.std)
	zend.SmartStrFreeEx(&object.prefix[0], 0)
	zend.SmartStrFreeEx(&object.prefix[1], 0)
	zend.SmartStrFreeEx(&object.prefix[2], 0)
	zend.SmartStrFreeEx(&object.prefix[3], 0)
	zend.SmartStrFreeEx(&object.prefix[4], 0)
	zend.SmartStrFreeEx(&object.prefix[5], 0)
	zend.SmartStrFreeEx(&object.postfix[0], 0)
}

/* }}} */

func spl_RecursiveIteratorIterator_new_ex(class_type *zend.ZendClassEntry, init_prefix int) *zend.ZendObject {
	var intern *SplRecursiveItObject
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_recursive_it_object"), class_type)
	if init_prefix != 0 {
		zend.SmartStrAppendlEx(&intern.prefix[0], "", 0, 0)
		zend.SmartStrAppendlEx(&intern.prefix[1], "| ", 2, 0)
		zend.SmartStrAppendlEx(&intern.prefix[2], "  ", 2, 0)
		zend.SmartStrAppendlEx(&intern.prefix[3], "|-", 2, 0)
		zend.SmartStrAppendlEx(&intern.prefix[4], "\\-", 2, 0)
		zend.SmartStrAppendlEx(&intern.prefix[5], "", 0, 0)
		zend.SmartStrAppendlEx(&intern.postfix[0], "", 0, 0)
	}
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.std.handlers = &SplHandlersRecItIt
	return &intern.std
}

/* }}} */

func spl_RecursiveIteratorIterator_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 0)
}

/* }}} */

func spl_RecursiveTreeIterator_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 1)
}

/* }}} */

var ArginfoRecursiveItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
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
		uint32(g.SizeOf("arginfo_recursive_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_RecursiveIteratorIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_RecursiveIteratorIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_RecursiveIteratorIterator_key,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_RecursiveIteratorIterator_current,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_RecursiveIteratorIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getDepth",
		zim_spl_RecursiveIteratorIterator_getDepth,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getSubIterator",
		zim_spl_RecursiveIteratorIterator_getSubIterator,
		arginfo_recursive_it_getSubIterator,
		uint32(g.SizeOf("arginfo_recursive_it_getSubIterator")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_RecursiveIteratorIterator_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"beginIteration",
		zim_spl_RecursiveIteratorIterator_beginIteration,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"endIteration",
		zim_spl_RecursiveIteratorIterator_endIteration,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"callHasChildren",
		zim_spl_RecursiveIteratorIterator_callHasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"callGetChildren",
		zim_spl_RecursiveIteratorIterator_callGetChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"beginChildren",
		zim_spl_RecursiveIteratorIterator_beginChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"endChildren",
		zim_spl_RecursiveIteratorIterator_endChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"nextElement",
		zim_spl_RecursiveIteratorIterator_nextElement,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setMaxDepth",
		zim_spl_RecursiveIteratorIterator_setMaxDepth,
		arginfo_recursive_it_setMaxDepth,
		uint32(g.SizeOf("arginfo_recursive_it_setMaxDepth")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getMaxDepth",
		zim_spl_RecursiveIteratorIterator_getMaxDepth,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplRecursiveTreeIteratorGetPrefix(object *SplRecursiveItObject, return_value *zend.Zval) {
	var str zend.SmartStr = zend.SmartStr{0}
	var has_next zend.Zval
	var level int
	zend.SmartStrAppendlEx(&str, object.GetPrefix()[0].s.val, object.GetPrefix()[0].s.len_, 0)
	for level = 0; level < object.GetLevel(); level++ {
		zend.ZendCallMethod(&object.iterators[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", g.SizeOf("\"hasnext\"")-1, &has_next, 0, nil, nil)
		if has_next.u1.v.type_ != 0 {
			if has_next.u1.v.type_ == 3 {
				zend.SmartStrAppendlEx(&str, object.GetPrefix()[1].s.val, object.GetPrefix()[1].s.len_, 0)
			} else {
				zend.SmartStrAppendlEx(&str, object.GetPrefix()[2].s.val, object.GetPrefix()[2].s.len_, 0)
			}
			zend.ZvalPtrDtor(&has_next)
		}
	}
	zend.ZendCallMethod(&object.iterators[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", g.SizeOf("\"hasnext\"")-1, &has_next, 0, nil, nil)
	if has_next.u1.v.type_ != 0 {
		if has_next.u1.v.type_ == 3 {
			zend.SmartStrAppendlEx(&str, object.GetPrefix()[3].s.val, object.GetPrefix()[3].s.len_, 0)
		} else {
			zend.SmartStrAppendlEx(&str, object.GetPrefix()[4].s.val, object.GetPrefix()[4].s.len_, 0)
		}
		zend.ZvalPtrDtor(&has_next)
	}
	zend.SmartStrAppendlEx(&str, object.GetPrefix()[5].s.val, object.GetPrefix()[5].s.len_, 0)
	zend.SmartStr0(&str)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str.s
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}
func SplRecursiveTreeIteratorGetEntry(object *SplRecursiveItObject, return_value *zend.Zval) {
	var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	var data *zend.Zval
	data = iterator.funcs.get_current_data(iterator)
	if data != nil {
		if data.u1.v.type_ == 10 {
			data = &(*data).value.ref.val
		}

		/* TODO: Remove this special case? */

		if data.u1.v.type_ == 7 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendKnownStrings[zend.ZEND_STR_ARRAY_CAPITALIZED]
			__z.value.str = __s
			__z.u1.type_info = 6
		} else {
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = data
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			if return_value.u1.v.type_ != 6 {
				zend._convertToString(return_value)
			}
		}

		/* TODO: Remove this special case? */

	}
}
func SplRecursiveTreeIteratorGetPostfix(object *SplRecursiveItObject, return_value *zend.Zval) {
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = object.GetPostfix()[0].s
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	zend.ZvalAddrefP(return_value)
}

/* {{{ proto RecursiveTreeIterator::__construct(RecursiveIterator|IteratorAggregate it [, int flags = RTIT_BYPASS_KEY [, int cit_flags = CIT_CATCH_GET_CHILD [, mode = RIT_SELF_FIRST ]]]) throws InvalidArgumentException
   RecursiveIteratorIterator to generate ASCII graphic trees for the entries in a RecursiveIterator */

func zim_spl_RecursiveTreeIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplRecursiveItItConstruct(execute_data, return_value, spl_ce_RecursiveTreeIterator, zend.ZendCeIterator, RIT_RecursiveTreeIterator)
}

/* {{{ proto void RecursiveTreeIterator::setPrefixPart(int part, string prefix) throws OutOfRangeException
   Sets prefix parts as used in getPrefix() */

func zim_spl_RecursiveTreeIterator_setPrefixPart(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var part zend.ZendLong
	var prefix *byte
	var prefix_len int
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "ls", &part, &prefix, &prefix_len) == zend.FAILURE {
		return
	}
	if 0 > part || part > 5 {
		zend.ZendThrowExceptionEx(spl_ce_OutOfRangeException, 0, "Use RecursiveTreeIterator::PREFIX_* constant")
		return
	}
	zend.SmartStrFreeEx(&object.prefix[part], 0)
	zend.SmartStrAppendlEx(&object.prefix[part], prefix, prefix_len, 0)
}

/* {{{ proto string RecursiveTreeIterator::getPrefix()
   Returns the string to place in front of current element */

func zim_spl_RecursiveTreeIterator_getPrefix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPrefix(object, return_value)
}

/* {{{ proto void RecursiveTreeIterator::setPostfix(string prefix)
   Sets postfix as used in getPostfix() */

func zim_spl_RecursiveTreeIterator_setPostfix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var postfix *byte
	var postfix_len int
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s", &postfix, &postfix_len) == zend.FAILURE {
		return
	}
	zend.SmartStrFreeEx(&object.postfix[0], 0)
	zend.SmartStrAppendlEx(&object.postfix[0], postfix, postfix_len, 0)
}

/* {{{ proto string RecursiveTreeIterator::getEntry()
   Returns the string presentation built for current element */

func zim_spl_RecursiveTreeIterator_getEntry(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetEntry(object, return_value)
}

/* {{{ proto string RecursiveTreeIterator::getPostfix()
   Returns the string to place after the current element */

func zim_spl_RecursiveTreeIterator_getPostfix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, return_value)
}

/* {{{ proto mixed RecursiveTreeIterator::current()
   Returns the current element prefixed and postfixed */

func zim_spl_RecursiveTreeIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var prefix zend.Zval
	var entry zend.Zval
	var postfix zend.Zval
	var ptr *byte
	var str *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	if (object.GetFlags() & RTIT_BYPASS_CURRENT) != 0 {
		var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
		var data *zend.Zval
		if object.GetIterators() == nil {
			zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
			return
		}
		iterator = object.GetIterators()[object.GetLevel()].GetIterator()
		data = iterator.funcs.get_current_data(iterator)
		if data != nil {
			var _z3 *zend.Zval = data
			if (_z3.u1.type_info & 0xff00) != 0 {
				if (_z3.u1.type_info & 0xff) == 10 {
					_z3 = &(*_z3).value.ref.val
					if (_z3.u1.type_info & 0xff00) != 0 {
						zend.ZvalAddrefP(_z3)
					}
				} else {
					zend.ZvalAddrefP(_z3)
				}
			}
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = _z3
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			return
		} else {
			return_value.u1.type_info = 1
			return
		}
	}
	&prefix.u1.type_info = 1
	&entry.u1.type_info = 1
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetEntry(object, &entry)
	if entry.u1.v.type_ != 6 {
		zend.ZvalPtrDtor(&prefix)
		zend.ZvalPtrDtor(&entry)
		return_value.u1.type_info = 1
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = zend.ZendStringAlloc(prefix.value.str.len_+entry.value.str.len_+postfix.value.str.len_, 0)
	ptr = str.val
	memcpy(ptr, prefix.value.str.val, prefix.value.str.len_)
	ptr += prefix.value.str.len_
	memcpy(ptr, entry.value.str.val, entry.value.str.len_)
	ptr += entry.value.str.len_
	memcpy(ptr, postfix.value.str.val, postfix.value.str.len_)
	ptr += postfix.value.str.len_
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&entry)
	zend.ZvalPtrDtor(&postfix)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* {{{ proto mixed RecursiveTreeIterator::key()
   Returns the current key prefixed and postfixed */

func zim_spl_RecursiveTreeIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(&(execute_data.This).value.obj)
	var iterator *zend.ZendObjectIterator
	var prefix zend.Zval
	var key zend.Zval
	var postfix zend.Zval
	var key_copy zend.Zval
	var ptr *byte
	var str *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	iterator = object.GetIterators()[object.GetLevel()].GetIterator()
	if iterator.funcs.get_current_key != nil {
		iterator.funcs.get_current_key(iterator, &key)
	} else {
		&key.u1.type_info = 1
	}
	if (object.GetFlags() & RTIT_BYPASS_KEY) != 0 {
		var __z *zend.Zval = return_value
		var __zv *zend.Zval = &key
		if __zv.u1.v.type_ != 10 {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		} else {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = &(*__zv).value.ref.val
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ZvalPtrDtor(__zv)
		}
		return
	}
	if key.u1.v.type_ != 6 {
		if zend.ZendMakePrintableZval(&key, &key_copy) != 0 {
			key = key_copy
		}
	}
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = zend.ZendStringAlloc(prefix.value.str.len_+key.value.str.len_+postfix.value.str.len_, 0)
	ptr = str.val
	memcpy(ptr, prefix.value.str.val, prefix.value.str.len_)
	ptr += prefix.value.str.len_
	memcpy(ptr, key.value.str.val, key.value.str.len_)
	ptr += key.value.str.len_
	memcpy(ptr, postfix.value.str.val, postfix.value.str.len_)
	ptr += postfix.value.str.len_
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&key)
	zend.ZvalPtrDtor(&postfix)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

var ArginfoRecursiveTreeItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
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
		uint32(g.SizeOf("arginfo_recursive_tree_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_RecursiveIteratorIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_RecursiveIteratorIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_RecursiveTreeIterator_key,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_RecursiveTreeIterator_current,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_RecursiveIteratorIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"beginIteration",
		zim_spl_RecursiveIteratorIterator_beginIteration,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"endIteration",
		zim_spl_RecursiveIteratorIterator_endIteration,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"callHasChildren",
		zim_spl_RecursiveIteratorIterator_callHasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"callGetChildren",
		zim_spl_RecursiveIteratorIterator_callGetChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"beginChildren",
		zim_spl_RecursiveIteratorIterator_beginChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"endChildren",
		zim_spl_RecursiveIteratorIterator_endChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"nextElement",
		zim_spl_RecursiveIteratorIterator_nextElement,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPrefix",
		zim_spl_RecursiveTreeIterator_getPrefix,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setPrefixPart",
		zim_spl_RecursiveTreeIterator_setPrefixPart,
		arginfo_recursive_tree_it_setPrefixPart,
		uint32(g.SizeOf("arginfo_recursive_tree_it_setPrefixPart")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getEntry",
		zim_spl_RecursiveTreeIterator_getEntry,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setPostfix",
		zim_spl_RecursiveTreeIterator_setPostfix,
		arginfo_recursive_tree_it_setPostfix,
		uint32(g.SizeOf("arginfo_recursive_tree_it_setPostfix")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPostfix",
		zim_spl_RecursiveTreeIterator_getPostfix,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplDualItGetMethod(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var intern *SplDualItObject
	intern = SplDualItFromObj(*object)
	function_handler = zend.ZendStdGetMethod(object, method, key)
	if function_handler == nil && intern.GetCe() != nil {
		if g.Assign(&function_handler, zend.ZendHashFindPtr(&intern.inner.ce.function_table, method)) == nil {
			if intern.inner.zobject.value.obj.handlers.get_method != nil {
				*object = intern.inner.zobject.value.obj
				function_handler = (*object).handlers.get_method(object, method, key)
			}
		} else {
			*object = intern.inner.zobject.value.obj
		}
	}
	return function_handler
}

// #define SPL_CHECK_CTOR(intern,classname) if ( intern -> dit_type == DIT_Unknown ) { zend_throw_exception_ex ( spl_ce_BadMethodCallException , 0 , "Classes derived from %s must call %s::__construct()" , ZSTR_VAL ( ( spl_ce_ ## classname ) -> name ) , ZSTR_VAL ( ( spl_ce_ ## classname ) -> name ) ) ; return ; }

// #define APPENDIT_CHECK_CTOR(intern) SPL_CHECK_CTOR ( intern , AppendIterator )

func SplCitCheckFlags(flags zend.ZendLong) int {
	var cnt zend.ZendLong = 0
	if (flags & CIT_CALL_TOSTRING) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_KEY) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_CURRENT) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_INNER) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if cnt <= 1 {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplDualItConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, dit_type DualItType) *SplDualItObject {
	var zobject *zend.Zval
	var retval zend.Zval
	var intern *SplDualItObject
	var ce *zend.ZendClassEntry = nil
	var inc_refcount int = 1
	var error_handling zend.ZendErrorHandling
	intern = SplDualItFromObj(&(execute_data.This).value.obj)
	if intern.GetDitType() != DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s::getIterator() must be called exactly once per instance", ce_base.name.val)
		return nil
	}
	intern.SetDitType(dit_type)
	switch dit_type {
	case DIT_LimitIterator:
		intern.SetOffset(0)
		intern.SetCount(-1)
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "O|ll", &zobject, ce_inner, &intern.u.limit.offset, &intern.u.limit.count) == zend.FAILURE {
			return nil
		}
		if intern.GetOffset() < 0 {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter offset must be >= 0", 0)
			return nil
		}
		if intern.GetCount() < 0 && intern.GetCount() != -1 {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter count must either be -1 or a value greater than or equal 0", 0)
			return nil
		}
		break
	case DIT_CachingIterator:

	case DIT_RecursiveCachingIterator:
		var flags zend.ZendLong = CIT_CALL_TOSTRING
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "O|l", &zobject, ce_inner, &flags) == zend.FAILURE {
			return nil
		}
		if SplCitCheckFlags(flags) != zend.SUCCESS {
			zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
			return nil
		}
		intern.SetUCachingFlags(intern.GetUCachingFlags() | flags&CIT_PUBLIC)
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &intern.u.caching.zcache
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		break
	case DIT_IteratorIterator:
		var ce_cast *zend.ZendClassEntry
		var class_name *zend.ZendString
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "O|S", &zobject, ce_inner, &class_name) == zend.FAILURE {
			return nil
		}
		ce = zobject.value.obj.ce
		if zend.InstanceofFunction(ce, zend.ZendCeIterator) == 0 {
			if execute_data.This.u2.num_args > 1 {
				if !(g.Assign(&ce_cast, zend.ZendLookupClass(class_name))) || zend.InstanceofFunction(ce, ce_cast) == 0 || ce_cast.get_iterator == nil {
					zend.ZendThrowException(spl_ce_LogicException, "Class to downcast to not found or not base class or does not implement Traversable", 0)
					return nil
				}
				ce = ce_cast
			}
			if zend.InstanceofFunction(ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethod(zobject, ce, &ce.iterator_funcs_ptr.zf_new_iterator, "getiterator", g.SizeOf("\"getiterator\"")-1, &retval, 0, nil, nil)
				if zend.EG.exception != nil {
					zend.ZvalPtrDtor(&retval)
					return nil
				}
				if retval.u1.v.type_ != 8 || zend.InstanceofFunction(retval.value.obj.ce, zend.ZendCeTraversable) == 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "%s::getIterator() must return an object that implements Traversable", ce.name.val)
					return nil
				}
				zobject = &retval
				ce = zobject.value.obj.ce
				inc_refcount = 0
			}
		}
		break
	case DIT_AppendIterator:
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		SplInstantiate(spl_ce_ArrayIterator, &intern.u.append.zarrayit)
		zend.ZendCallMethod(&intern.u.append.zarrayit, spl_ce_ArrayIterator, &spl_ce_ArrayIterator.constructor, "__construct", g.SizeOf("\"__construct\"")-1, nil, 0, nil, nil)
		intern.SetUAppendIterator(spl_ce_ArrayIterator.get_iterator(spl_ce_ArrayIterator, &intern.u.append.zarrayit, 0))
		zend.ZendRestoreErrorHandling(&error_handling)
		return intern
	case DIT_RegexIterator:

	case DIT_RecursiveRegexIterator:
		var regex *zend.ZendString
		var mode zend.ZendLong = REGIT_MODE_MATCH
		intern.SetUseFlags(execute_data.This.u2.num_args >= 5)
		intern.SetURegexFlags(0)
		intern.SetPregFlags(0)
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "OS|lll", &zobject, ce_inner, &regex, &mode, &intern.u.regex.flags, &intern.u.regex.preg_flags) == zend.FAILURE {
			return nil
		}
		if mode < 0 || mode >= REGIT_MODE_MAX {
			zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+"%"+"lld", mode)
			return nil
		}
		intern.SetMode(mode)
		intern.SetURegexRegex(zend.ZendStringCopy(regex))
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		intern.SetPce(pcre_get_compiled_regex_cache(regex))
		zend.ZendRestoreErrorHandling(&error_handling)
		if intern.GetPce() == nil {

			/* pcre_get_compiled_regex_cache has already sent error */

			return nil

			/* pcre_get_compiled_regex_cache has already sent error */

		}
		php_pcre_pce_incref(intern.GetPce())
		break
	case DIT_CallbackFilterIterator:

	case DIT_RecursiveCallbackFilterIterator:
		var cfi *_spl_cbfilter_it_intern = zend._emalloc(g.SizeOf("* cfi"))
		cfi.fci.object = nil
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "Of", &zobject, ce_inner, &cfi.fci, &cfi.fcc) == zend.FAILURE {
			zend._efree(cfi)
			return nil
		}
		if &(cfi.fci.function_name).u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(cfi.fci.function_name))
		}
		cfi.SetObject(cfi.fcc.object)
		if cfi.GetObject() != nil {
			zend.ZendGcAddref(&(cfi.GetObject()).gc)
		}
		intern.SetCbfilter(cfi)
		break
	default:
		if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "O", &zobject, ce_inner) == zend.FAILURE {
			return nil
		}
		break
	}
	if inc_refcount != 0 {
		zend.ZvalAddrefP(zobject)
	}
	var __z *zend.Zval = &intern.inner.zobject
	__z.value.obj = zobject.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	if dit_type == DIT_IteratorIterator {
		intern.SetCe(ce)
	} else {
		intern.SetCe(zobject.value.obj.ce)
	}
	intern.SetObject(zobject.value.obj)
	intern.SetInnerIterator(intern.GetCe().get_iterator(intern.GetCe(), zobject, 0))
	return intern
}

/* {{{ proto FilterIterator::__construct(Iterator it)
   Create an Iterator from another iterator */

func zim_spl_FilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_FilterIterator, zend.ZendCeIterator, DIT_FilterIterator)
}

/* {{{ proto CallbackFilterIterator::__construct(Iterator it, callback func)
   Create an Iterator from another iterator */

func zim_spl_CallbackFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_CallbackFilterIterator, zend.ZendCeIterator, DIT_CallbackFilterIterator)
}

/* {{{ proto Iterator FilterIterator::getInnerIterator()
    proto Iterator CachingIterator::getInnerIterator()
    proto Iterator LimitIterator::getInnerIterator()
    proto Iterator ParentIterator::getInnerIterator()
Get the inner iterator */

func zim_spl_dual_it_getInnerIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.inner.zobject.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.inner.zobject
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}
func SplDualItFree(intern *SplDualItObject) {
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().funcs.invalidate_current != nil {
		intern.GetInnerIterator().funcs.invalidate_current(intern.GetInnerIterator())
	}
	if intern.current.data.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&intern.current.data)
		&intern.current.data.u1.type_info = 0
	}
	if intern.current.key.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&intern.current.key)
		&intern.current.key.u1.type_info = 0
	}
	if intern.GetDitType() == DIT_CachingIterator || intern.GetDitType() == DIT_RecursiveCachingIterator {
		if intern.u.caching.zstr.u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&intern.u.caching.zstr)
			&intern.u.caching.zstr.u1.type_info = 0
		}
		if intern.u.caching.zchildren.u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&intern.u.caching.zchildren)
			&intern.u.caching.zchildren.u1.type_info = 0
		}
	}
}
func SplDualItRewind(intern *SplDualItObject) {
	SplDualItFree(intern)
	intern.SetPos(0)
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().funcs.rewind != nil {
		intern.GetInnerIterator().funcs.rewind(intern.GetInnerIterator())
	}
}
func SplDualItValid(intern *SplDualItObject) int {
	if intern.GetInnerIterator() == nil {
		return zend.FAILURE
	}

	/* FAILURE / SUCCESS */

	return intern.GetInnerIterator().funcs.valid(intern.GetInnerIterator())

	/* FAILURE / SUCCESS */
}
func SplDualItFetch(intern *SplDualItObject, check_more int) int {
	var data *zend.Zval
	SplDualItFree(intern)
	if check_more == 0 || SplDualItValid(intern) == zend.SUCCESS {
		data = intern.GetInnerIterator().funcs.get_current_data(intern.GetInnerIterator())
		if data != nil {
			var _z1 *zend.Zval = &intern.current.data
			var _z2 *zend.Zval = data
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		}
		if intern.GetInnerIterator().funcs.get_current_key != nil {
			intern.GetInnerIterator().funcs.get_current_key(intern.GetInnerIterator(), &intern.current.key)
			if zend.EG.exception != nil {
				zend.ZvalPtrDtor(&intern.current.key)
				&intern.current.key.u1.type_info = 0
			}
		} else {
			var __z *zend.Zval = &intern.current.key
			__z.value.lval = intern.GetPos()
			__z.u1.type_info = 4
		}
		if zend.EG.exception != nil {
			return zend.FAILURE
		} else {
			return zend.SUCCESS
		}
	}
	return zend.FAILURE
}
func SplDualItNext(intern *SplDualItObject, do_free int) {
	if do_free != 0 {
		SplDualItFree(intern)
	} else if intern.GetInnerIterator() == nil {
		zend.ZendThrowError(nil, "The inner constructor wasn't initialized with an iterator instance")
		return
	}
	intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
	intern.GetPos()++
}

/* {{{ proto void ParentIterator::rewind()
       proto void IteratorIterator::rewind()
   Rewind the iterator
*/

func ZimSplDualItRewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplDualItFetch(intern, 1)
}

/* {{{ proto bool FilterIterator::valid()
    proto bool ParentIterator::valid()
    proto bool IteratorIterator::valid()
    proto bool NoRewindIterator::valid()
Check whether the current element is valid */

func ZimSplDualItValid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.data.u1.v.type_ != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed FilterIterator::key()
    proto mixed CachingIterator::key()
    proto mixed LimitIterator::key()
    proto mixed ParentIterator::key()
    proto mixed IteratorIterator::key()
    proto mixed NoRewindIterator::key()
    proto mixed AppendIterator::key()
Get the current key */

func ZimSplDualItKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.key.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.current.key
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto mixed FilterIterator::current()
    proto mixed CachingIterator::current()
    proto mixed LimitIterator::current()
    proto mixed ParentIterator::current()
    proto mixed IteratorIterator::current()
    proto mixed NoRewindIterator::current()
Get the current element value */

func ZimSplDualItCurrent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.data.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.current.data
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto void ParentIterator::next()
    proto void IteratorIterator::next()
    proto void NoRewindIterator::next()
Move the iterator forward */

func ZimSplDualItNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	SplDualItFetch(intern, 1)
}
func SplFilterItFetch(zthis *zend.Zval, intern *SplDualItObject) {
	var retval zend.Zval
	for SplDualItFetch(intern, 1) == zend.SUCCESS {
		zend.ZendCallMethod(zthis, intern.std.ce, nil, "accept", g.SizeOf("\"accept\"")-1, &retval, 0, nil, nil)
		if retval.u1.v.type_ != 0 {
			if zend.ZendIsTrue(&retval) != 0 {
				zend.ZvalPtrDtor(&retval)
				return
			}
			zend.ZvalPtrDtor(&retval)
		}
		if zend.EG.exception != nil {
			return
		}
		intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
	}
	SplDualItFree(intern)
}
func SplFilterItRewind(zthis *zend.Zval, intern *SplDualItObject) {
	SplDualItRewind(intern)
	SplFilterItFetch(zthis, intern)
}
func SplFilterItNext(zthis *zend.Zval, intern *SplDualItObject) {
	SplDualItNext(intern, 1)
	SplFilterItFetch(zthis, intern)
}

/* {{{ proto void FilterIterator::rewind()
   Rewind the iterator */

func zim_spl_FilterIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItRewind(&(execute_data.This), intern)
}

/* {{{ proto void FilterIterator::next()
   Move the iterator forward */

func zim_spl_FilterIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItNext(&(execute_data.This), intern)
}

/* {{{ proto RecursiveCallbackFilterIterator::__construct(RecursiveIterator it, callback func)
   Create a RecursiveCallbackFilterIterator from a RecursiveIterator */

func zim_spl_RecursiveCallbackFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveCallbackFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveCallbackFilterIterator)
}

/* {{{ proto RecursiveFilterIterator::__construct(RecursiveIterator it)
   Create a RecursiveFilterIterator from a RecursiveIterator */

func zim_spl_RecursiveFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveFilterIterator)
}

/* {{{ proto bool RecursiveFilterIterator::hasChildren()
   Check whether the inner iterator's current element has children */

func zim_spl_RecursiveFilterIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "haschildren", g.SizeOf("\"haschildren\"")-1, &retval, 0, nil, nil)
	if retval.u1.v.type_ != 0 {
		var __z *zend.Zval = return_value
		var __zv *zend.Zval = &retval
		if __zv.u1.v.type_ != 10 {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		} else {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = &(*__zv).value.ref.val
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ZvalPtrDtor(__zv)
		}
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* {{{ proto RecursiveFilterIterator RecursiveFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveFilterIterator */

func zim_spl_RecursiveFilterIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", g.SizeOf("\"getchildren\"")-1, &retval, 0, nil, nil)
	if zend.EG.exception == nil && retval.u1.v.type_ != 0 {
		SplInstantiateArgEx1(&(execute_data.This).value.obj.ce, return_value, &retval)
	}
	zend.ZvalPtrDtor(&retval)
}

/* {{{ proto RecursiveCallbackFilterIterator RecursiveCallbackFilterIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveCallbackFilterIterator */

func zim_spl_RecursiveCallbackFilterIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", g.SizeOf("\"getchildren\"")-1, &retval, 0, nil, nil)
	if zend.EG.exception == nil && retval.u1.v.type_ != 0 {
		SplInstantiateArgEx2(&(execute_data.This).value.obj.ce, return_value, &retval, &intern.u.cbfilter.fci.function_name)
	}
	zend.ZvalPtrDtor(&retval)
}

/* {{{ proto ParentIterator::__construct(RecursiveIterator it)
   Create a ParentIterator from a RecursiveIterator */

func zim_spl_ParentIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_ParentIterator, spl_ce_RecursiveIterator, DIT_ParentIterator)
}

/* {{{ proto RegexIterator::__construct(Iterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RegexIterator from another iterator and a regular expression */

func zim_spl_RegexIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RegexIterator, zend.ZendCeIterator, DIT_RegexIterator)
}

/* {{{ proto bool CallbackFilterIterator::accept()
   Calls the callback with the current value, the current key and the inner iterator as arguments */

func zim_spl_CallbackFilterIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	var fci *zend.ZendFcallInfo = &intern.u.cbfilter.GetFci()
	var fcc *zend.ZendFcallInfoCache = &intern.u.cbfilter.GetFcc()
	var params []zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.current.data.u1.v.type_ == 0 || intern.current.key.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
		return
	}
	var _z1 *zend.Zval = &params[0]
	var _z2 *zend.Zval = &intern.current.data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	var _z1 *zend.Zval = &params[1]
	var _z2 *zend.Zval = &intern.current.key
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	var _z1 *zend.Zval = &params[2]
	var _z2 *zend.Zval = &intern.inner.zobject
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	fci.retval = return_value
	fci.param_count = 3
	fci.params = params
	fci.no_separation = 0
	if zend.ZendCallFunction(fci, fcc) != zend.SUCCESS || return_value.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
		return
	}
	if zend.EG.exception != nil {
		return_value.u1.type_info = 1
		return
	}

	/* zend_call_function may change args to IS_REF */

	var _z1 *zend.Zval = &intern.current.data
	var _z2 *zend.Zval = &params[0]
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	var _z1 *zend.Zval = &intern.current.key
	var _z2 *zend.Zval = &params[1]
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func zim_spl_RegexIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var result *zend.ZendString
	var subject *zend.ZendString
	var count int = 0
	var zcount zend.Zval
	var rv zend.Zval
	var match_data *pcre2_match_data
	var re *pcre2_code
	var rc int
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.data.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
		return
	}
	if (intern.GetURegexFlags() & REGIT_USE_KEY) != 0 {
		subject = zend.ZvalGetString(&intern.current.key)
	} else {
		if intern.current.data.u1.v.type_ == 7 {
			return_value.u1.type_info = 2
			return
		}
		subject = zend.ZvalGetString(&intern.current.data)
	}

	/* Exception during string conversion. */

	if zend.EG.exception != nil {
		return
	}
	switch intern.GetMode() {
	case REGIT_MODE_MAX:

	case REGIT_MODE_MATCH:
		re = php_pcre_pce_re(intern.GetPce())
		match_data = php_pcre_create_match_data(0, re)
		if match_data == nil {
			return_value.u1.type_info = 2
			return
		}
		rc = pcre2_match(re, PCRE2_SPTR(subject).val, subject.len_, 0, 0, match_data, php_pcre_mctx())
		if rc >= 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		php_pcre_free_match_data(match_data)
		break
	case REGIT_MODE_ALL_MATCHES:

	case REGIT_MODE_GET_MATCH:
		zend.ZvalPtrDtor(&intern.current.data)
		&intern.current.data.u1.type_info = 0
		php_pcre_match_impl(intern.GetPce(), subject, &zcount, &intern.current.data, intern.GetMode() == REGIT_MODE_ALL_MATCHES, intern.GetUseFlags(), intern.GetPregFlags(), 0)
		if zcount.value.lval > 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		break
	case REGIT_MODE_SPLIT:
		zend.ZvalPtrDtor(&intern.current.data)
		&intern.current.data.u1.type_info = 0
		php_pcre_split_impl(intern.GetPce(), subject, &intern.current.data, -1, intern.GetPregFlags())
		count = intern.current.data.value.arr.nNumOfElements
		if count > 1 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		break
	case REGIT_MODE_REPLACE:
		var replacement *zend.Zval = zend.ZendReadProperty(intern.std.ce, &(execute_data.This), "replacement", g.SizeOf("\"replacement\"")-1, 1, &rv)
		var replacement_str *zend.ZendString = zend.ZvalTryGetString(replacement)
		if replacement_str == nil {
			return
		}
		result = php_pcre_replace_impl(intern.GetPce(), subject, subject.val, subject.len_, replacement_str, -1, &count)
		if (intern.GetURegexFlags() & REGIT_USE_KEY) != 0 {
			zend.ZvalPtrDtor(&intern.current.key)
			var __z *zend.Zval = &intern.current.key
			var __s *zend.ZendString = result
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				__z.u1.type_info = 6 | 1<<0<<8
			}
		} else {
			zend.ZvalPtrDtor(&intern.current.data)
			var __z *zend.Zval = &intern.current.data
			var __s *zend.ZendString = result
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				__z.u1.type_info = 6 | 1<<0<<8
			}
		}
		zend.ZendStringRelease(replacement_str)
		if count > 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
	}
	if (intern.GetURegexFlags() & REGIT_INVERTED) != 0 {
		if return_value.u1.v.type_ != 3 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
	}
	zend.ZendStringReleaseEx(subject, 0)
}

/* {{{ proto string RegexIterator::getRegex()
   Returns current regular expression */

func zim_spl_RegexIterator_getRegex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = intern.GetURegexRegex()
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		zend.ZendGcAddref(&__s.gc)
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* {{{ proto bool RegexIterator::getMode()
   Returns current operation mode */

func zim_spl_RegexIterator_getMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetMode()
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool RegexIterator::setMode(int new_mode)
   Set new operation mode */

func zim_spl_RegexIterator_setMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var mode zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &mode) == zend.FAILURE {
		return
	}
	if mode < 0 || mode >= REGIT_MODE_MAX {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+"%"+"lld", mode)
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetMode(mode)
}

/* {{{ proto bool RegexIterator::getFlags()
   Returns current operation flags */

func zim_spl_RegexIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetURegexFlags()
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool RegexIterator::setFlags(int new_flags)
   Set operation flags */

func zim_spl_RegexIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &flags) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetURegexFlags(flags)
}

/* {{{ proto bool RegexIterator::getFlags()
   Returns current PREG flags (if in use or NULL) */

func zim_spl_RegexIterator_getPregFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetUseFlags() != 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = intern.GetPregFlags()
		__z.u1.type_info = 4
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	}
}

/* {{{ proto bool RegexIterator::setPregFlags(int new_flags)
   Set PREG flags */

func zim_spl_RegexIterator_setPregFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var preg_flags zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &preg_flags) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetPregFlags(preg_flags)
	intern.SetUseFlags(1)
}

/* {{{ proto RecursiveRegexIterator::__construct(RecursiveIterator it, string regex [, int mode [, int flags [, int preg_flags]]])
   Create an RecursiveRegexIterator from another recursive iterator and a regular expression */

func zim_spl_RecursiveRegexIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveRegexIterator, spl_ce_RecursiveIterator, DIT_RecursiveRegexIterator)
}

/* {{{ proto RecursiveRegexIterator RecursiveRegexIterator::getChildren()
   Return the inner iterator's children contained in a RecursiveRegexIterator */

func zim_spl_RecursiveRegexIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", g.SizeOf("\"getchildren\"")-1, &retval, 0, nil, nil)
	if zend.EG.exception == nil {
		var args []zend.Zval
		var _z1 *zend.Zval = &args[0]
		var _z2 *zend.Zval = &retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		var __z *zend.Zval = &args[1]
		var __s *zend.ZendString = intern.GetURegexRegex()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		var __z *zval = &args[2]
		__z.value.lval = intern.GetMode()
		__z.u1.type_info = 4
		var __z *zval = &args[3]
		__z.value.lval = intern.GetURegexFlags()
		__z.u1.type_info = 4
		var __z *zend.Zval = &args[4]
		__z.value.lval = intern.GetPregFlags()
		__z.u1.type_info = 4
		SplInstantiateArgN(&(execute_data.This).value.obj.ce, return_value, 5, args)
		zend.ZvalPtrDtor(&args[0])
		zend.ZvalPtrDtor(&args[1])
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveRegexIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.data.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
		return
	} else if intern.current.data.u1.v.type_ == 7 {
		if intern.current.data.value.arr.nNumOfElements > 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	zend.ZendCallMethod(&(execute_data.This), spl_ce_RegexIterator, nil, "accept", g.SizeOf("\"accept\"")-1, return_value, 0, nil, nil)
}

/* {{{ spl_dual_it_dtor */

func SplDualItDtor(_object *zend.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	SplDualItFree(object)
	if object.GetInnerIterator() != nil {
		zend.ZendIteratorDtor(object.GetInnerIterator())
	}
}

/* }}} */

func SplDualItFreeStorage(_object *zend.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)
	if object.inner.zobject.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&object.inner.zobject)
	}
	if object.GetDitType() == DIT_AppendIterator {
		zend.ZendIteratorDtor(object.GetUAppendIterator())
		if object.u.append.zarrayit.u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&object.u.append.zarrayit)
		}
	}
	if object.GetDitType() == DIT_CachingIterator || object.GetDitType() == DIT_RecursiveCachingIterator {
		zend.ZvalPtrDtor(&object.u.caching.zcache)
	}
	if object.GetDitType() == DIT_RegexIterator || object.GetDitType() == DIT_RecursiveRegexIterator {
		if object.GetPce() != nil {
			php_pcre_pce_decref(object.GetPce())
		}
		if object.GetURegexRegex() != nil {
			zend.ZendStringReleaseEx(object.GetURegexRegex(), 0)
		}
	}
	if object.GetDitType() == DIT_CallbackFilterIterator || object.GetDitType() == DIT_RecursiveCallbackFilterIterator {
		if object.GetCbfilter() != nil {
			var cbfilter *_spl_cbfilter_it_intern = object.GetCbfilter()
			object.SetCbfilter(nil)
			zend.ZvalPtrDtor(&cbfilter.fci.function_name)
			if cbfilter.fci.object != nil {
				zend.ZendObjectRelease(cbfilter.fci.object)
			}
			zend._efree(cbfilter)
		}
	}
	zend.ZendObjectStdDtor(&object.std)
}

/* }}} */

func SplDualItNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var intern *SplDualItObject
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_dual_it_object"), class_type)
	intern.SetDitType(DIT_Unknown)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.std.handlers = &SplHandlersDualIt
	return &intern.std
}

/* }}} */

var ArginfoFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
}
var spl_funcs_FilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_FilterIterator___construct,
		ArginfoFilterItConstruct,
		uint32(g.SizeOf("arginfo_filter_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_FilterIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		ZimSplDualItValid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_FilterIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"accept",
		nil,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoCallbackFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
	{"callback", 0, 0, 0},
}
var spl_funcs_CallbackFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_CallbackFilterIterator___construct,
		ArginfoCallbackFilterItConstruct,
		uint32(g.SizeOf("arginfo_callback_filter_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"accept",
		zim_spl_CallbackFilterIterator_accept,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRecursiveCallbackFilterItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("RecursiveIterator"), 0, 0},
	{"callback", 0, 0, 0},
}
var spl_funcs_RecursiveCallbackFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveCallbackFilterIterator___construct,
		ArginfoRecursiveCallbackFilterItConstruct,
		uint32(g.SizeOf("arginfo_recursive_callback_filter_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_RecursiveCallbackFilterIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoParentItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("RecursiveIterator"), 0, 0},
}
var spl_funcs_RecursiveFilterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveFilterIterator___construct,
		ArginfoParentItConstruct,
		uint32(g.SizeOf("arginfo_parent_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_RecursiveFilterIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_ParentIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_ParentIterator___construct,
		ArginfoParentItConstruct,
		uint32(g.SizeOf("arginfo_parent_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"accept",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRegexItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(2)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
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
		uint32(g.SizeOf("arginfo_regex_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"accept",
		zim_spl_RegexIterator_accept,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getMode",
		zim_spl_RegexIterator_getMode,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setMode",
		zim_spl_RegexIterator_setMode,
		ArginfoRegexItSetMode,
		uint32(g.SizeOf("arginfo_regex_it_set_mode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_RegexIterator_getFlags,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_RegexIterator_setFlags,
		ArginfoRegexItSetFlags,
		uint32(g.SizeOf("arginfo_regex_it_set_flags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPregFlags",
		zim_spl_RegexIterator_getPregFlags,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setPregFlags",
		zim_spl_RegexIterator_setPregFlags,
		ArginfoRegexItSetPregFlags,
		uint32(g.SizeOf("arginfo_regex_it_set_preg_flags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getRegex",
		zim_spl_RegexIterator_getRegex,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRecRegexItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(2)), 0, 0, 0},
	{"iterator", zend.ZendType("RecursiveIterator"), 0, 0},
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
		uint32(g.SizeOf("arginfo_rec_regex_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"accept",
		zim_spl_RecursiveRegexIterator_accept,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_RecursiveFilterIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_RecursiveRegexIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplLimitItValid(intern *SplDualItObject) int {
	/* FAILURE / SUCCESS */

	if intern.GetCount() != -1 && intern.GetPos() >= intern.GetOffset()+intern.GetCount() {
		return zend.FAILURE
	} else {
		return SplDualItValid(intern)
	}

	/* FAILURE / SUCCESS */
}
func SplLimitItSeek(intern *SplDualItObject, pos zend.ZendLong) {
	var zpos zend.Zval
	SplDualItFree(intern)
	if pos < intern.GetOffset() {
		zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Cannot seek to "+"%"+"lld"+" which is below the offset "+"%"+"lld", pos, intern.GetOffset())
		return
	}
	if pos >= intern.GetOffset()+intern.GetCount() && intern.GetCount() != -1 {
		zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Cannot seek to "+"%"+"lld"+" which is behind offset "+"%"+"lld"+" plus count "+"%"+"lld", pos, intern.GetOffset(), intern.GetCount())
		return
	}
	if pos != intern.GetPos() && zend.InstanceofFunction(intern.GetCe(), spl_ce_SeekableIterator) != 0 {
		var __z *zend.Zval = &zpos
		__z.value.lval = pos
		__z.u1.type_info = 4
		SplDualItFree(intern)
		zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "seek", g.SizeOf("\"seek\"")-1, nil, 1, &zpos, nil)
		if zend.EG.exception == nil {
			intern.SetPos(pos)
			if SplLimitItValid(intern) == zend.SUCCESS {
				SplDualItFetch(intern, 0)
			}
		}
	} else {

		/* emulate the forward seek, by next() calls */

		if pos < intern.GetPos() {
			SplDualItRewind(intern)
		}
		for pos > intern.GetPos() && SplDualItValid(intern) == zend.SUCCESS {
			SplDualItNext(intern, 1)
		}
		if SplDualItValid(intern) == zend.SUCCESS {
			SplDualItFetch(intern, 1)
		}
	}
}

/* {{{ proto LimitIterator::__construct(Iterator it [, int offset, int count])
   Construct a LimitIterator from an Iterator with a given starting offset and optionally a maximum count */

func zim_spl_LimitIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_LimitIterator, zend.ZendCeIterator, DIT_LimitIterator)
}

/* {{{ proto void LimitIterator::rewind()
   Rewind the iterator to the specified starting offset */

func zim_spl_LimitIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplLimitItSeek(intern, intern.GetOffset())
}

/* {{{ proto bool LimitIterator::valid()
   Check whether the current element is valid */

func zim_spl_LimitIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it

	/*    RETURN_BOOL(spl_limit_it_valid(intern) == SUCCESS);*/

	if (intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount()) && intern.current.data.u1.v.type_ != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto void LimitIterator::next()
   Move the iterator forward */

func zim_spl_LimitIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	if intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount() {
		SplDualItFetch(intern, 1)
	}
}

/* {{{ proto void LimitIterator::seek(int position)
   Seek to the given position */

func zim_spl_LimitIterator_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var pos zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &pos) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplLimitItSeek(intern, pos)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetPos()
	__z.u1.type_info = 4
	return
}

/* {{{ proto int LimitIterator::getPosition()
   Return the current position */

func zim_spl_LimitIterator_getPosition(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetPos()
	__z.u1.type_info = 4
	return
}

var ArginfoSeekableItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"position", 4<<2 | g.Cond(false, 0x1, 0x0), 0, 0},
}
var spl_funcs_SeekableIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"seek",
		nil,
		ArginfoSeekableItSeek,
		uint32(g.SizeOf("arginfo_seekable_it_seek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoLimitItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
	{"offset", 0, 0, 0},
	{"count", 0, 0, 0},
}
var ArginfoLimitItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"position", 0, 0, 0}}
var spl_funcs_LimitIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_LimitIterator___construct,
		ArginfoLimitItConstruct,
		uint32(g.SizeOf("arginfo_limit_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_LimitIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_LimitIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_LimitIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"seek",
		zim_spl_LimitIterator_seek,
		ArginfoLimitItSeek,
		uint32(g.SizeOf("arginfo_limit_it_seek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPosition",
		zim_spl_LimitIterator_getPosition,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplCachingItValid(intern *SplDualItObject) int {
	if (intern.GetUCachingFlags() & CIT_VALID) != 0 {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplCachingItHasNext(intern *SplDualItObject) int { return SplDualItValid(intern) }
func SplCachingItNext(intern *SplDualItObject) {
	if SplDualItFetch(intern, 1) == zend.SUCCESS {
		intern.SetUCachingFlags(intern.GetUCachingFlags() | CIT_VALID)

		/* Full cache ? */

		if (intern.GetUCachingFlags() & CIT_FULL_CACHE) != 0 {
			var key *zend.Zval = &intern.current.key
			var data *zend.Zval = &intern.current.data
			if data.u1.v.type_ == 10 {
				data = &(*data).value.ref.val
			}
			if data.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(data)
			}
			zend.ArraySetZvalKey(intern.u.caching.zcache.value.arr, key, data)
			zend.ZvalPtrDtor(data)
		}

		/* Recursion ? */

		if intern.GetDitType() == DIT_RecursiveCachingIterator {
			var retval zend.Zval
			var zchildren zend.Zval
			var zflags zend.Zval
			zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "haschildren", g.SizeOf("\"haschildren\"")-1, &retval, 0, nil, nil)
			if zend.EG.exception != nil {
				zend.ZvalPtrDtor(&retval)
				if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
					zend.ZendClearException()
				} else {
					return
				}
			} else {
				if zend.ZendIsTrue(&retval) != 0 {
					zend.ZendCallMethod(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", g.SizeOf("\"getchildren\"")-1, &zchildren, 0, nil, nil)
					if zend.EG.exception != nil {
						zend.ZvalPtrDtor(&zchildren)
						if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
							zend.ZendClearException()
						} else {
							zend.ZvalPtrDtor(&retval)
							return
						}
					} else {
						var __z *zend.Zval = &zflags
						__z.value.lval = intern.GetUCachingFlags() & CIT_PUBLIC
						__z.u1.type_info = 4
						SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &intern.u.caching.zchildren, &zchildren, &zflags)
						zend.ZvalPtrDtor(&zchildren)
					}
				}
				zend.ZvalPtrDtor(&retval)
				if zend.EG.exception != nil {
					if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
						zend.ZendClearException()
					} else {
						return
					}
				}
			}
		}
		if (intern.GetUCachingFlags() & (CIT_TOSTRING_USE_INNER | CIT_CALL_TOSTRING)) != 0 {
			var use_copy int
			var expr_copy zend.Zval
			if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_INNER) != 0 {
				var _z1 *zend.Zval = &intern.u.caching.zstr
				var _z2 *zend.Zval = &intern.inner.zobject
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			} else {
				var _z1 *zend.Zval = &intern.u.caching.zstr
				var _z2 *zend.Zval = &intern.current.data
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			}
			use_copy = zend.ZendMakePrintableZval(&intern.u.caching.zstr, &expr_copy)
			if use_copy != 0 {
				var _z1 *zend.Zval = &intern.u.caching.zstr
				var _z2 *zend.Zval = &expr_copy
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			} else {
				if &(intern.GetZstr()).u1.v.type_flags != 0 {
					zend.ZvalAddrefP(&(intern.GetZstr()))
				}
			}
		}
		SplDualItNext(intern, 0)
	} else {
		intern.SetUCachingFlags(intern.GetUCachingFlags() &^ CIT_VALID)
	}
}
func SplCachingItRewind(intern *SplDualItObject) {
	SplDualItRewind(intern)
	zend.ZendHashClean(intern.u.caching.zcache.value.arr)
	SplCachingItNext(intern)
}

/* {{{ proto CachingIterator::__construct(Iterator it [, flags = CIT_CALL_TOSTRING])
   Construct a CachingIterator from an Iterator */

func zim_spl_CachingIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_CachingIterator, zend.ZendCeIterator, DIT_CachingIterator)
}

/* {{{ proto void CachingIterator::rewind()
   Rewind the iterator */

func zim_spl_CachingIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItRewind(intern)
}

/* {{{ proto bool CachingIterator::valid()
   Check whether the current element is valid */

func zim_spl_CachingIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if SplCachingItValid(intern) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto void CachingIterator::next()
   Move the iterator forward */

func zim_spl_CachingIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItNext(intern)
}

/* {{{ proto bool CachingIterator::hasNext()
   Check whether the inner iterator has a valid next element */

func zim_spl_CachingIterator_hasNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if SplCachingItHasNext(intern) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto string CachingIterator::__toString()
   Return the string representation of the current element */

func zim_spl_CachingIterator___toString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & (CIT_CALL_TOSTRING | CIT_TOSTRING_USE_KEY | CIT_TOSTRING_USE_CURRENT | CIT_TOSTRING_USE_INNER)) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not fetch string value (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_KEY) != 0 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &intern.current.key
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		if return_value.u1.v.type_ != 6 {
			zend._convertToString(return_value)
		}
		return
	} else if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_CURRENT) != 0 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &intern.current.data
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		if return_value.u1.v.type_ != 6 {
			zend._convertToString(return_value)
		}
		return
	}
	if intern.u.caching.zstr.u1.v.type_ == 6 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = &intern.u.caching.zstr.value.str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* {{{ proto void CachingIterator::offsetSet(mixed index, mixed newval)
   Set given index in cache */

func zim_spl_CachingIterator_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var value *zend.Zval
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "Sz", &key, &value) == zend.FAILURE {
		return
	}
	if value.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(value)
	}
	zend.ZendSymtableUpdate(intern.u.caching.zcache.value.arr, key, value)
}

/* }}} */

func zim_spl_CachingIterator_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var value *zend.Zval
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S", &key) == zend.FAILURE {
		return
	}
	if g.Assign(&value, zend.ZendSymtableFind(intern.u.caching.zcache.value.arr, key)) == nil {
		zend.ZendError(1<<3, "Undefined index: %s", key.val)
		return
	}
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func zim_spl_CachingIterator_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S", &key) == zend.FAILURE {
		return
	}
	zend.ZendSymtableDel(intern.u.caching.zcache.value.arr, key)
}

/* }}} */

func zim_spl_CachingIterator_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S", &key) == zend.FAILURE {
		return
	}
	if zend.ZendSymtableExists(intern.u.caching.zcache.value.arr, key) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_CachingIterator_getCache(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = &intern.u.caching.zcache
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func zim_spl_CachingIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetUCachingFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_CachingIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &flags) == zend.FAILURE {
		return
	}
	if SplCitCheckFlags(flags) != zend.SUCCESS {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
		return
	}
	if (intern.GetUCachingFlags()&CIT_CALL_TOSTRING) != 0 && (flags&CIT_CALL_TOSTRING) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag CALL_TO_STRING is not possible", 0)
		return
	}
	if (intern.GetUCachingFlags()&CIT_TOSTRING_USE_INNER) != 0 && (flags&CIT_TOSTRING_USE_INNER) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag TOSTRING_USE_INNER is not possible", 0)
		return
	}
	if (flags&CIT_FULL_CACHE) != 0 && (intern.GetUCachingFlags()&CIT_FULL_CACHE) == 0 {

		/* clear on (re)enable */

		zend.ZendHashClean(intern.u.caching.zcache.value.arr)

		/* clear on (re)enable */

	}
	intern.SetUCachingFlags(intern.GetUCachingFlags() & ^CIT_PUBLIC | flags&CIT_PUBLIC)
}

/* }}} */

func zim_spl_CachingIterator_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", &(execute_data.This).value.obj.ce.name.val)
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.u.caching.zcache.value.arr.nNumOfElements
	__z.u1.type_info = 4
	return
}

/* }}} */

var ArginfoCachingItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
	{"flags", 0, 0, 0},
}
var arginfo_caching_it_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_caching_it_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"index", 0, 0, 0}}
var arginfo_caching_it_offsetSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"index", 0, 0, 0}, {"newval", 0, 0, 0}}
var spl_funcs_CachingIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_CachingIterator___construct,
		ArginfoCachingItConstruct,
		uint32(g.SizeOf("arginfo_caching_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_CachingIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_CachingIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_CachingIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasNext",
		zim_spl_CachingIterator_hasNext,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__toString",
		zim_spl_CachingIterator___toString,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_CachingIterator_getFlags,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_CachingIterator_setFlags,
		arginfo_caching_it_setFlags,
		uint32(g.SizeOf("arginfo_caching_it_setFlags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetGet",
		zim_spl_CachingIterator_offsetGet,
		arginfo_caching_it_offsetGet,
		uint32(g.SizeOf("arginfo_caching_it_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetSet",
		zim_spl_CachingIterator_offsetSet,
		arginfo_caching_it_offsetSet,
		uint32(g.SizeOf("arginfo_caching_it_offsetSet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetUnset",
		zim_spl_CachingIterator_offsetUnset,
		arginfo_caching_it_offsetGet,
		uint32(g.SizeOf("arginfo_caching_it_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"offsetExists",
		zim_spl_CachingIterator_offsetExists,
		arginfo_caching_it_offsetGet,
		uint32(g.SizeOf("arginfo_caching_it_offsetGet")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getCache",
		zim_spl_CachingIterator_getCache,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_CachingIterator_count,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto RecursiveCachingIterator::__construct(RecursiveIterator it [, flags = CIT_CALL_TOSTRING])
   Create an iterator from a RecursiveIterator */

func zim_spl_RecursiveCachingIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveCachingIterator, spl_ce_RecursiveIterator, DIT_RecursiveCachingIterator)
}

/* {{{ proto bool RecursiveCachingIterator::hasChildren()
   Check whether the current element of the inner iterator has children */

func zim_spl_RecursiveCachingIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.u.caching.zchildren.u1.v.type_ != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto RecursiveCachingIterator RecursiveCachingIterator::getChildren()
Return the inner iterator's children as a RecursiveCachingIterator */

func zim_spl_RecursiveCachingIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.u.caching.zchildren.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.u.caching.zchildren
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}

var ArginfoCachingRecItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
	{"flags", 0, 0, 0},
}
var spl_funcs_RecursiveCachingIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveCachingIterator___construct,
		ArginfoCachingRecItConstruct,
		uint32(g.SizeOf("arginfo_caching_rec_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_RecursiveCachingIterator_hasChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_RecursiveCachingIterator_getChildren,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto IteratorIterator::__construct(Traversable it)
   Create an iterator from anything that is traversable */

func zim_spl_IteratorIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_IteratorIterator, zend.ZendCeTraversable, DIT_IteratorIterator)
}

var ArginfoIteratorItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
}
var spl_funcs_IteratorIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_IteratorIterator___construct,
		ArginfoIteratorItConstruct,
		uint32(g.SizeOf("arginfo_iterator_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		ZimSplDualItRewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		ZimSplDualItValid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		ZimSplDualItCurrent,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		ZimSplDualItNext,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto NoRewindIterator::__construct(Iterator it)
   Create an iterator from another iterator */

func zim_spl_NoRewindIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_NoRewindIterator, zend.ZendCeIterator, DIT_NoRewindIterator)
}

/* {{{ proto void NoRewindIterator::rewind()
   Prevent a call to inner iterators rewind() */

func zim_spl_NoRewindIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto bool NoRewindIterator::valid()
   Return inner iterators valid() */

func zim_spl_NoRewindIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetInnerIterator().funcs.valid(intern.GetInnerIterator()) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto mixed NoRewindIterator::key()
   Return inner iterators key() */

func zim_spl_NoRewindIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetInnerIterator().funcs.get_current_key != nil {
		intern.GetInnerIterator().funcs.get_current_key(intern.GetInnerIterator(), return_value)
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto mixed NoRewindIterator::current()
   Return inner iterators current() */

func zim_spl_NoRewindIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var data *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	data = intern.GetInnerIterator().funcs.get_current_data(intern.GetInnerIterator())
	if data != nil {
		var _z3 *zend.Zval = data
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* {{{ proto void NoRewindIterator::next()
   Return inner iterators next() */

func zim_spl_NoRewindIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
}

var ArginfoNorewindItConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
}
var spl_funcs_NoRewindIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_NoRewindIterator___construct,
		ArginfoNorewindItConstruct,
		uint32(g.SizeOf("arginfo_norewind_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_NoRewindIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_NoRewindIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_NoRewindIterator_key,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_NoRewindIterator_current,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_NoRewindIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto InfiniteIterator::__construct(Iterator it)
   Create an iterator from another iterator */

func zim_spl_InfiniteIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_InfiniteIterator, zend.ZendCeIterator, DIT_InfiniteIterator)
}

/* {{{ proto void InfiniteIterator::next()
   Prevent a call to inner iterators rewind() (internally the current data will be fetched if valid()) */

func zim_spl_InfiniteIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	if SplDualItValid(intern) == zend.SUCCESS {
		SplDualItFetch(intern, 0)
	} else {
		SplDualItRewind(intern)
		if SplDualItValid(intern) == zend.SUCCESS {
			SplDualItFetch(intern, 0)
		}
	}
}

var spl_funcs_InfiniteIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_InfiniteIterator___construct,
		ArginfoNorewindItConstruct,
		uint32(g.SizeOf("arginfo_norewind_it___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_InfiniteIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto void EmptyIterator::rewind()
   Does nothing  */

func zim_spl_EmptyIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ proto false EmptyIterator::valid()
   Return false */

func zim_spl_EmptyIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	return_value.u1.type_info = 2
	return
}

/* {{{ proto void EmptyIterator::key()
   Throws exception BadMethodCallException */

func zim_spl_EmptyIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the key of an EmptyIterator", 0)
}

/* {{{ proto void EmptyIterator::current()
   Throws exception BadMethodCallException */

func zim_spl_EmptyIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the value of an EmptyIterator", 0)
}

/* {{{ proto void EmptyIterator::next()
   Does nothing */

func zim_spl_EmptyIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

var spl_funcs_EmptyIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"rewind",
		zim_spl_EmptyIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_EmptyIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_EmptyIterator_key,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_EmptyIterator_current,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_EmptyIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplAppendItNextIterator(intern *SplDualItObject) int {
	SplDualItFree(intern)
	if intern.inner.zobject.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&intern.inner.zobject)
		&intern.inner.zobject.u1.type_info = 0
		intern.SetCe(nil)
		if intern.GetInnerIterator() != nil {
			zend.ZendIteratorDtor(intern.GetInnerIterator())
			intern.SetInnerIterator(nil)
		}
	}
	if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) == zend.SUCCESS {
		var it *zend.Zval
		it = intern.GetUAppendIterator().funcs.get_current_data(intern.GetUAppendIterator())
		var _z1 *zend.Zval = &intern.inner.zobject
		var _z2 *zend.Zval = it
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		intern.SetCe(it.value.obj.ce)
		intern.SetInnerIterator(intern.GetCe().get_iterator(intern.GetCe(), it, 0))
		SplDualItRewind(intern)
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplAppendItFetch(intern *SplDualItObject) {
	for SplDualItValid(intern) != zend.SUCCESS {
		intern.GetUAppendIterator().funcs.move_forward(intern.GetUAppendIterator())
		if SplAppendItNextIterator(intern) != zend.SUCCESS {
			return
		}
	}
	SplDualItFetch(intern, 0)
}
func SplAppendItNext(intern *SplDualItObject) {
	if SplDualItValid(intern) == zend.SUCCESS {
		SplDualItNext(intern, 1)
	}
	SplAppendItFetch(intern)
}

/* {{{ proto AppendIterator::__construct()
   Create an AppendIterator */

func zim_spl_AppendIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_AppendIterator, zend.ZendCeIterator, DIT_AppendIterator)
}

/* {{{ proto void AppendIterator::append(Iterator it)
   Append an iterator */

func zim_spl_AppendIterator_append(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *zend.Zval
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "O", &it, zend.ZendCeIterator) == zend.FAILURE {
		return
	}
	if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) == zend.SUCCESS && SplDualItValid(intern) != zend.SUCCESS {
		SplArrayIteratorAppend(&intern.u.append.zarrayit, it)
		intern.GetUAppendIterator().funcs.move_forward(intern.GetUAppendIterator())
	} else {
		SplArrayIteratorAppend(&intern.u.append.zarrayit, it)
	}
	if intern.GetInnerIterator() == nil || SplDualItValid(intern) != zend.SUCCESS {
		if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) != zend.SUCCESS {
			intern.GetUAppendIterator().funcs.rewind(intern.GetUAppendIterator())
		}
		for {
			SplAppendItNextIterator(intern)
			if intern.inner.zobject.value.obj == it.value.obj {
				break
			}
		}
		SplAppendItFetch(intern)
	}
}

/* {{{ proto mixed AppendIterator::current()
   Get the current element value */

func zim_spl_AppendIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItFetch(intern, 1)
	if intern.current.data.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.current.data
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* {{{ proto void AppendIterator::rewind()
   Rewind to the first iterator and rewind the first iterator, too */

func zim_spl_AppendIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetUAppendIterator().funcs.rewind(intern.GetUAppendIterator())
	if SplAppendItNextIterator(intern) == zend.SUCCESS {
		SplAppendItFetch(intern)
	}
}

/* {{{ proto bool AppendIterator::valid()
   Check if the current state is valid */

func zim_spl_AppendIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.current.data.u1.v.type_ != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto void AppendIterator::next()
   Forward to next element */

func zim_spl_AppendIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplAppendItNext(intern)
}

/* {{{ proto int AppendIterator::getIteratorIndex()
   Get index of iterator */

func zim_spl_AppendIterator_getIteratorIndex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "Classes derived from %s must call %s::__construct()", spl_ce_AppendIterator.name.val, spl_ce_AppendIterator.name.val)
		return
	}
	SplArrayIteratorKey(&intern.u.append.zarrayit, return_value)
}

/* {{{ proto ArrayIterator AppendIterator::getArrayIterator()
   Get access to inner ArrayIterator */

func zim_spl_AppendIterator_getArrayIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var value *zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = SplDualItFromObj(&(execute_data.This).value.obj)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	value = &intern.u.append.zarrayit
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

var ArginfoAppendItAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Iterator"), 0, 0},
}
var spl_funcs_AppendIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_AppendIterator___construct,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"append",
		zim_spl_AppendIterator_append,
		ArginfoAppendItAppend,
		uint32(g.SizeOf("arginfo_append_it_append")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_AppendIterator_rewind,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_AppendIterator_valid,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		ZimSplDualItKey,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_AppendIterator_current,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_AppendIterator_next,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInnerIterator",
		zim_spl_dual_it_getInnerIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getIteratorIndex",
		zim_spl_AppendIterator_getIteratorIndex,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getArrayIterator",
		zim_spl_AppendIterator_getArrayIterator,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

func SplIteratorApply(obj *zend.Zval, apply_func SplIteratorApplyFuncT, puser any) int {
	var iter *zend.ZendObjectIterator
	var ce *zend.ZendClassEntry = obj.value.obj.ce
	iter = ce.get_iterator(ce, obj, 0)
	if zend.EG.exception != nil {
		goto done
	}
	iter.index = 0
	if iter.funcs.rewind != nil {
		iter.funcs.rewind(iter)
		if zend.EG.exception != nil {
			goto done
		}
	}
	for iter.funcs.valid(iter) == zend.SUCCESS {
		if zend.EG.exception != nil {
			goto done
		}
		if apply_func(iter, puser) == 1<<1 || zend.EG.exception != nil {
			goto done
		}
		iter.index++
		iter.funcs.move_forward(iter)
		if zend.EG.exception != nil {
			goto done
		}
	}
done:
	if iter != nil {
		zend.ZendIteratorDtor(iter)
	}
	if zend.EG.exception != nil {
		return zend.FAILURE
	} else {
		return zend.SUCCESS
	}
}

/* }}} */

func SplIteratorToArrayApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *zend.Zval
	var return_value *zend.Zval = (*zend.Zval)(puser)
	data = iter.funcs.get_current_data(iter)
	if zend.EG.exception != nil {
		return 1 << 1
	}
	if data == nil {
		return 1 << 1
	}
	if iter.funcs.get_current_key != nil {
		var key zend.Zval
		iter.funcs.get_current_key(iter, &key)
		if zend.EG.exception != nil {
			return 1 << 1
		}
		zend.ArraySetZvalKey(return_value.value.arr, &key, data)
		zend.ZvalPtrDtor(&key)
	} else {
		if data.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(data)
		}
		zend.AddNextIndexZval(return_value, data)
	}
	return 0
}

/* }}} */

func SplIteratorToValuesApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *zend.Zval
	var return_value *zend.Zval = (*zend.Zval)(puser)
	data = iter.funcs.get_current_data(iter)
	if zend.EG.exception != nil {
		return 1 << 1
	}
	if data == nil {
		return 1 << 1
	}
	if data.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(data)
	}
	zend.AddNextIndexZval(return_value, data)
	return 0
}

/* }}} */

func ZifIteratorToArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var use_keys zend.ZendBool = 1
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O|b", &obj, zend.ZendCeTraversable, &use_keys) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplIteratorApply(obj, g.Cond(use_keys != 0, SplIteratorToArrayApply, SplIteratorToValuesApply), any(return_value))
}
func SplIteratorCountApply(iter *zend.ZendObjectIterator, puser any) int {
	*((*zend.ZendLong)(puser))++
	return 0
}

/* }}} */

func ZifIteratorCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var count zend.ZendLong = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "O", &obj, zend.ZendCeTraversable) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if SplIteratorApply(obj, SplIteratorCountApply, any(&count)) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = count
	__z.u1.type_info = 4
	return
}

/* }}} */

// @type SplIteratorApplyInfo struct
func SplIteratorFuncApply(iter *zend.ZendObjectIterator, puser any) int {
	var retval zend.Zval
	var apply_info *SplIteratorApplyInfo = (*SplIteratorApplyInfo)(puser)
	var result int
	apply_info.GetCount()++
	zend.ZendFcallInfoCall(&apply_info.fci, &apply_info.fcc, &retval, nil)
	if zend.ZendIsTrue(&retval) != 0 {
		result = 0
	} else {
		result = 1 << 1
	}
	zend.ZvalPtrDtor(&retval)
	return result
}

/* }}} */

func ZifIteratorApply(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var apply_info SplIteratorApplyInfo
	apply_info.SetArgs(nil)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "Of|a!", &apply_info.obj, zend.ZendCeTraversable, &apply_info.fci, &apply_info.fcc, &apply_info.args) == zend.FAILURE {
		return
	}
	apply_info.SetCount(0)
	zend.ZendFcallInfoArgs(&apply_info.fci, apply_info.GetArgs())
	if SplIteratorApply(apply_info.GetObj(), SplIteratorFuncApply, any(&apply_info)) == zend.FAILURE {
		zend.ZendFcallInfoArgs(&apply_info.fci, nil)
		return
	}
	zend.ZendFcallInfoArgs(&apply_info.fci, nil)
	var __z *zend.Zval = return_value
	__z.value.lval = apply_info.GetCount()
	__z.u1.type_info = 4
	return
}

/* }}} */

var spl_funcs_OuterIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"getInnerIterator",
		nil,
		ArginfoRecursiveItVoid,
		uint32(g.SizeOf("arginfo_recursive_it_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ PHP_MINIT_FUNCTION(spl_iterators)
 */

func ZmStartupSplIterators(type_ int, module_number int) int {
	SplRegisterInterface(&spl_ce_RecursiveIterator, "RecursiveIterator", spl_funcs_RecursiveIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIterator, 1, zend.ZendCeIterator)
	SplRegisterStdClass(&spl_ce_RecursiveIteratorIterator, "RecursiveIteratorIterator", spl_RecursiveIteratorIterator_new, spl_funcs_RecursiveIteratorIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIteratorIterator, 1, zend.ZendCeIterator)
	memcpy(&SplHandlersRecItIt, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	SplHandlersRecItIt.offset = zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd())) - (*byte)(nil))
	SplHandlersRecItIt.get_method = SplRecursiveItGetMethod
	SplHandlersRecItIt.clone_obj = nil
	SplHandlersRecItIt.dtor_obj = spl_RecursiveIteratorIterator_dtor
	SplHandlersRecItIt.free_obj = spl_RecursiveIteratorIterator_free_storage
	memcpy(&SplHandlersDualIt, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	SplHandlersDualIt.offset = zend_long((*byte)(&((*SplDualItObject)(nil).GetStd())) - (*byte)(nil))
	SplHandlersDualIt.get_method = SplDualItGetMethod

	/*spl_handlers_dual_it.call_method = spl_dual_it_call_method;*/

	SplHandlersDualIt.clone_obj = nil
	SplHandlersDualIt.dtor_obj = SplDualItDtor
	SplHandlersDualIt.free_obj = SplDualItFreeStorage
	spl_ce_RecursiveIteratorIterator.get_iterator = SplRecursiveItGetIterator
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "LEAVES_ONLY", g.SizeOf("\"LEAVES_ONLY\"")-1, zend.ZendLong(RIT_LEAVES_ONLY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "SELF_FIRST", g.SizeOf("\"SELF_FIRST\"")-1, zend.ZendLong(RIT_SELF_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CHILD_FIRST", g.SizeOf("\"CHILD_FIRST\"")-1, zend.ZendLong(RIT_CHILD_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CATCH_GET_CHILD", g.SizeOf("\"CATCH_GET_CHILD\"")-1, zend.ZendLong(CIT_CATCH_GET_CHILD))
	SplRegisterInterface(&spl_ce_OuterIterator, "OuterIterator", spl_funcs_OuterIterator)
	zend.ZendClassImplements(spl_ce_OuterIterator, 1, zend.ZendCeIterator)
	SplRegisterStdClass(&spl_ce_IteratorIterator, "IteratorIterator", SplDualItNew, spl_funcs_IteratorIterator)
	zend.ZendClassImplements(spl_ce_IteratorIterator, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_IteratorIterator, 1, spl_ce_OuterIterator)
	SplRegisterSubClass(&spl_ce_FilterIterator, spl_ce_IteratorIterator, "FilterIterator", SplDualItNew, spl_funcs_FilterIterator)
	spl_ce_FilterIterator.ce_flags |= 1 << 6
	SplRegisterSubClass(&spl_ce_RecursiveFilterIterator, spl_ce_FilterIterator, "RecursiveFilterIterator", SplDualItNew, spl_funcs_RecursiveFilterIterator)
	zend.ZendClassImplements(spl_ce_RecursiveFilterIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_CallbackFilterIterator, spl_ce_FilterIterator, "CallbackFilterIterator", SplDualItNew, spl_funcs_CallbackFilterIterator)
	SplRegisterSubClass(&spl_ce_RecursiveCallbackFilterIterator, spl_ce_CallbackFilterIterator, "RecursiveCallbackFilterIterator", SplDualItNew, spl_funcs_RecursiveCallbackFilterIterator)
	zend.ZendClassImplements(spl_ce_RecursiveCallbackFilterIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_ParentIterator, spl_ce_RecursiveFilterIterator, "ParentIterator", SplDualItNew, spl_funcs_ParentIterator)
	SplRegisterInterface(&spl_ce_SeekableIterator, "SeekableIterator", spl_funcs_SeekableIterator)
	zend.ZendClassImplements(spl_ce_SeekableIterator, 1, zend.ZendCeIterator)
	SplRegisterSubClass(&spl_ce_LimitIterator, spl_ce_IteratorIterator, "LimitIterator", SplDualItNew, spl_funcs_LimitIterator)
	SplRegisterSubClass(&spl_ce_CachingIterator, spl_ce_IteratorIterator, "CachingIterator", SplDualItNew, spl_funcs_CachingIterator)
	zend.ZendClassImplements(spl_ce_CachingIterator, 1, zend.ZendCeArrayaccess)
	zend.ZendClassImplements(spl_ce_CachingIterator, 1, zend.ZendCeCountable)
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CALL_TOSTRING", g.SizeOf("\"CALL_TOSTRING\"")-1, zend.ZendLong(CIT_CALL_TOSTRING))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CATCH_GET_CHILD", g.SizeOf("\"CATCH_GET_CHILD\"")-1, zend.ZendLong(CIT_CATCH_GET_CHILD))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_KEY", g.SizeOf("\"TOSTRING_USE_KEY\"")-1, zend.ZendLong(CIT_TOSTRING_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_CURRENT", g.SizeOf("\"TOSTRING_USE_CURRENT\"")-1, zend.ZendLong(CIT_TOSTRING_USE_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_INNER", g.SizeOf("\"TOSTRING_USE_INNER\"")-1, zend.ZendLong(CIT_TOSTRING_USE_INNER))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "FULL_CACHE", g.SizeOf("\"FULL_CACHE\"")-1, zend.ZendLong(CIT_FULL_CACHE))
	SplRegisterSubClass(&spl_ce_RecursiveCachingIterator, spl_ce_CachingIterator, "RecursiveCachingIterator", SplDualItNew, spl_funcs_RecursiveCachingIterator)
	zend.ZendClassImplements(spl_ce_RecursiveCachingIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_NoRewindIterator, spl_ce_IteratorIterator, "NoRewindIterator", SplDualItNew, spl_funcs_NoRewindIterator)
	SplRegisterSubClass(&spl_ce_AppendIterator, spl_ce_IteratorIterator, "AppendIterator", SplDualItNew, spl_funcs_AppendIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIteratorIterator, 1, spl_ce_OuterIterator)
	SplRegisterSubClass(&spl_ce_InfiniteIterator, spl_ce_IteratorIterator, "InfiniteIterator", SplDualItNew, spl_funcs_InfiniteIterator)
	SplRegisterSubClass(&spl_ce_RegexIterator, spl_ce_FilterIterator, "RegexIterator", SplDualItNew, spl_funcs_RegexIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "USE_KEY", g.SizeOf("\"USE_KEY\"")-1, zend.ZendLong(REGIT_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "INVERT_MATCH", g.SizeOf("\"INVERT_MATCH\"")-1, zend.ZendLong(REGIT_INVERTED))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "MATCH", g.SizeOf("\"MATCH\"")-1, zend.ZendLong(REGIT_MODE_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "GET_MATCH", g.SizeOf("\"GET_MATCH\"")-1, zend.ZendLong(REGIT_MODE_GET_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "ALL_MATCHES", g.SizeOf("\"ALL_MATCHES\"")-1, zend.ZendLong(REGIT_MODE_ALL_MATCHES))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "SPLIT", g.SizeOf("\"SPLIT\"")-1, zend.ZendLong(REGIT_MODE_SPLIT))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "REPLACE", g.SizeOf("\"REPLACE\"")-1, zend.ZendLong(REGIT_MODE_REPLACE))
	SplRegisterProperty(spl_ce_RegexIterator, "replacement", g.SizeOf("\"replacement\"")-1, 0)
	SplRegisterSubClass(&spl_ce_RecursiveRegexIterator, spl_ce_RegexIterator, "RecursiveRegexIterator", SplDualItNew, spl_funcs_RecursiveRegexIterator)
	zend.ZendClassImplements(spl_ce_RecursiveRegexIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterStdClass(&spl_ce_EmptyIterator, "EmptyIterator", nil, spl_funcs_EmptyIterator)
	zend.ZendClassImplements(spl_ce_EmptyIterator, 1, zend.ZendCeIterator)
	SplRegisterSubClass(&spl_ce_RecursiveTreeIterator, spl_ce_RecursiveIteratorIterator, "RecursiveTreeIterator", spl_RecursiveTreeIterator_new, spl_funcs_RecursiveTreeIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_CURRENT", g.SizeOf("\"BYPASS_CURRENT\"")-1, zend.ZendLong(RTIT_BYPASS_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_KEY", g.SizeOf("\"BYPASS_KEY\"")-1, zend.ZendLong(RTIT_BYPASS_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_LEFT", g.SizeOf("\"PREFIX_LEFT\"")-1, zend.ZendLong(0))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_HAS_NEXT", g.SizeOf("\"PREFIX_MID_HAS_NEXT\"")-1, zend.ZendLong(1))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_LAST", g.SizeOf("\"PREFIX_MID_LAST\"")-1, zend.ZendLong(2))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_HAS_NEXT", g.SizeOf("\"PREFIX_END_HAS_NEXT\"")-1, zend.ZendLong(3))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_LAST", g.SizeOf("\"PREFIX_END_LAST\"")-1, zend.ZendLong(4))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_RIGHT", g.SizeOf("\"PREFIX_RIGHT\"")-1, zend.ZendLong(5))
	return zend.SUCCESS
}

/* }}} */

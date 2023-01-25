// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

// Source: <Zend/zend_execute.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_EXECUTE_H

// # include "zend_compile.h"

// # include "zend_hash.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

var ZendExecuteEx func(execute_data *ZendExecuteData)
var ZendExecuteInternal func(execute_data *ZendExecuteData, return_value *Zval)

/* export zend_pass_function to allow comparisons against it */

func ZEND_REF_TYPE_SOURCES(ref *ZendReference) ZendPropertyInfoSourceList { return ref.GetSources() }
func ZEND_REF_HAS_TYPE_SOURCES(ref *ZendReference) bool {
	return ZEND_REF_TYPE_SOURCES(ref).GetPtr() != nil
}
func ZEND_REF_FIRST_SOURCE(ref *ZendReference) *ZendPropertyInfo {
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(ref.GetSources().GetList()) != 0 {
		return ZEND_PROPERTY_INFO_SOURCE_TO_LIST(ref.GetSources().GetList()).GetPtr()[0]
	} else {
		return ref.GetSources().GetPtr()
	}
}
func ZendCopyToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, ref *ZendRefcounted) {
	ZVAL_COPY_VALUE(variable_ptr, value)
	if ZEND_CONST_COND(value_type == IS_CONST, 0) {
		if UNEXPECTED(Z_OPT_REFCOUNTED_P(variable_ptr)) {
			Z_ADDREF_P(variable_ptr)
		}
	} else if (value_type & (IS_CONST | IS_CV)) != 0 {
		if Z_OPT_REFCOUNTED_P(variable_ptr) {
			Z_ADDREF_P(variable_ptr)
		}
	} else if ZEND_CONST_COND(value_type == IS_VAR, 1) && UNEXPECTED(ref != nil) {
		if UNEXPECTED(GC_DELREF(ref) == 0) {
			EfreeSize(ref, b.SizeOf("zend_reference"))
		} else if Z_OPT_REFCOUNTED_P(variable_ptr) {
			Z_ADDREF_P(variable_ptr)
		}
	}
}
func ZendAssignToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, strict ZendBool) *Zval {
	var ref *ZendRefcounted = nil
	if ZEND_CONST_COND(value_type&(IS_VAR|IS_CV), 1) && Z_ISREF_P(value) {
		ref = Z_COUNTED_P(value)
		value = Z_REFVAL_P(value)
	}
	for {
		if UNEXPECTED(Z_REFCOUNTED_P(variable_ptr)) {
			var garbage *ZendRefcounted
			if Z_ISREF_P(variable_ptr) {
				if UNEXPECTED(ZEND_REF_HAS_TYPE_SOURCES(Z_REF_P(variable_ptr))) {
					return ZendAssignToTypedRef(variable_ptr, value, value_type, strict, ref)
				}
				variable_ptr = Z_REFVAL_P(variable_ptr)
				if EXPECTED(!(Z_REFCOUNTED_P(variable_ptr))) {
					break
				}
			}
			if Z_TYPE_P(variable_ptr) == IS_OBJECT && UNEXPECTED(Z_OBJ_HANDLER_P(variable_ptr, set) != nil) {
				Z_OBJ_HANDLER_P(variable_ptr, set)(variable_ptr, value)
				return variable_ptr
			}
			garbage = Z_COUNTED_P(variable_ptr)
			ZendCopyToVariable(variable_ptr, value, value_type, ref)
			if GC_DELREF(garbage) == 0 {
				RcDtorFunc(garbage)
			} else {

				/* optimized version of GC_ZVAL_CHECK_POSSIBLE_ROOT(variable_ptr) */

				if UNEXPECTED(GC_MAY_LEAK(garbage)) {
					GcPossibleRoot(garbage)
				}

				/* optimized version of GC_ZVAL_CHECK_POSSIBLE_ROOT(variable_ptr) */

			}
			return variable_ptr
		}
		break
	}
	ZendCopyToVariable(variable_ptr, value, value_type, ref)
	return variable_ptr
}

/* dedicated Zend executor functions - do not use! */

const ZEND_VM_STACK_HEADER_SLOTS *Zval = (ZEND_MM_ALIGNED_SIZE(b.SizeOf("struct _zend_vm_stack")) + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval"))

func ZEND_VM_STACK_ELEMENTS(stack ZendVmStack) __auto__ {
	return (*Zval)(stack) + ZEND_VM_STACK_HEADER_SLOTS
}

/*
 * In general in RELEASE build ZEND_ASSERT() must be zero-cost, but for some
 * reason, GCC generated worse code, performing CSE on assertion code and the
 * following "slow path" and moving memory read operatins from slow path into
 * common header. This made a degradation for the fast path.
 * The following "#if ZEND_DEBUG" eliminates it.
 */

// #define ZEND_ASSERT_VM_STACK(stack)

// #define ZEND_ASSERT_VM_STACK_GLOBAL

func ZendVmInitCallFrame(call *ZendExecuteData, call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) {
	call.SetFunc(func_)
	Z_PTR(call.GetThis()) = object_or_called_scope
	ZEND_CALL_INFO(call) = call_info
	ZEND_CALL_NUM_ARGS(call) = num_args
}
func ZendVmStackPushCallFrameEx(used_stack uint32, call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var call *ZendExecuteData = (*ZendExecuteData)(ExecutorGlobals.GetVmStackTop())
	if UNEXPECTED(used_stack > size_t((*byte)(ExecutorGlobals.GetVmStackEnd())-(*byte)(call))) {
		call = (*ZendExecuteData)(ZendVmStackExtend(used_stack))
		ZendVmInitCallFrame(call, call_info|ZEND_CALL_ALLOCATED, func_, num_args, object_or_called_scope)
		return call
	} else {
		ExecutorGlobals.SetVmStackTop((*Zval)((*byte)(call + used_stack)))
		ZendVmInitCallFrame(call, call_info, func_, num_args, object_or_called_scope)
		return call
	}
}
func ZendVmCalcUsedStack(num_args uint32, func_ *ZendFunction) uint32 {
	var used_stack uint32 = ZEND_CALL_FRAME_SLOT + num_args
	if EXPECTED(ZEND_USER_CODE(func_.GetType())) {
		used_stack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - MIN(func_.GetOpArray().GetNumArgs(), num_args)
	}
	return used_stack * b.SizeOf("zval")
}
func ZendVmStackPushCallFrame(call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var used_stack uint32 = ZendVmCalcUsedStack(num_args, func_)
	return ZendVmStackPushCallFrameEx(used_stack, call_info, func_, num_args, object_or_called_scope)
}
func ZendVmStackFreeExtraArgsEx(call_info uint32, call *ZendExecuteData) {
	if UNEXPECTED((call_info & ZEND_CALL_FREE_EXTRA_ARGS) != 0) {
		var count uint32 = ZEND_CALL_NUM_ARGS(call) - call.GetFunc().GetOpArray().GetNumArgs()
		var p *Zval = ZEND_CALL_VAR_NUM(call, call.GetFunc().GetOpArray().GetLastVar()+call.GetFunc().GetOpArray().GetT())
		for {
			if Z_REFCOUNTED_P(p) {
				var r *ZendRefcounted = Z_COUNTED_P(p)
				if GC_DELREF(r) == 0 {
					ZVAL_NULL(p)
					RcDtorFunc(r)
				} else {
					GcCheckPossibleRoot(r)
				}
			}
			p++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendVmStackFreeExtraArgs(call *ZendExecuteData) {
	ZendVmStackFreeExtraArgsEx(ZEND_CALL_INFO(call), call)
}
func ZendVmStackFreeArgs(call *ZendExecuteData) {
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	if EXPECTED(num_args > 0) {
		var p *Zval = ZEND_CALL_ARG(call, 1)
		for {
			if Z_REFCOUNTED_P(p) {
				var r *ZendRefcounted = Z_COUNTED_P(p)
				if GC_DELREF(r) == 0 {
					ZVAL_NULL(p)
					RcDtorFunc(r)
				}
			}
			p++
			if !(b.PreDec(&num_args)) {
				break
			}
		}
	}
}
func ZendVmStackFreeCallFrameEx(call_info uint32, call *ZendExecuteData) {
	if UNEXPECTED((call_info & ZEND_CALL_ALLOCATED) != 0) {
		var p ZendVmStack = ExecutorGlobals.GetVmStack()
		var prev ZendVmStack = p.GetPrev()
		ZEND_ASSERT(call == (*ZendExecuteData)(ZEND_VM_STACK_ELEMENTS(ExecutorGlobals.GetVmStack())))
		ExecutorGlobals.SetVmStackTop(prev.GetTop())
		ExecutorGlobals.SetVmStackEnd(prev.GetEnd())
		ExecutorGlobals.SetVmStack(prev)
		Efree(p)
	} else {
		ExecutorGlobals.SetVmStackTop((*Zval)(call))
	}
}
func ZendVmStackFreeCallFrame(call *ZendExecuteData) {
	ZendVmStackFreeCallFrameEx(ZEND_CALL_INFO(call), call)
}

/* services */

const ZEND_USER_OPCODE_CONTINUE = 0
const ZEND_USER_OPCODE_RETURN = 1
const ZEND_USER_OPCODE_DISPATCH = 2
const ZEND_USER_OPCODE_ENTER = 3
const ZEND_USER_OPCODE_LEAVE = 4
const ZEND_USER_OPCODE_DISPATCH_TO = 0x100

/* former zend_execute_locks.h */

type ZendFreeOp *Zval

func CACHE_ADDR(num __auto__) *any {
	return (*any)((*byte)(EX(run_time_cache) + num))
}
func CACHED_PTR(num __auto__) any {
	return (*any)((*byte)(EX(run_time_cache) + num))[0]
}
func CACHE_PTR(num __auto__, ptr any) {
	(*any)((*byte)(EX(run_time_cache) + num))[0] = ptr
}
func CACHED_POLYMORPHIC_PTR(num __auto__, ce __auto__) any {
	if EXPECTED((*any)((*byte)(EX(run_time_cache) + num))[0] == any(ce)) {
		return (*any)((*byte)(EX(run_time_cache) + num))[1]
	} else {
		return nil
	}
}
func CACHE_POLYMORPHIC_PTR(num uint32, ce any, ptr any) {
	var slot *any = (*any)((*byte)(EX(run_time_cache) + num))
	slot[0] = ce
	slot[1] = ptr
}
func CACHED_PTR_EX(slot *any) any     { return slot[0] }
func CACHE_PTR_EX(slot *any, ptr any) { slot[0] = ptr }
func CACHED_POLYMORPHIC_PTR_EX(slot __auto__, ce __auto__) __auto__ {
	if EXPECTED(slot[0] == ce) {
		return slot[1]
	} else {
		return nil
	}
}
func CACHE_POLYMORPHIC_PTR_EX(slot *any, ce *ZendClassEntry, ptr any) {
	slot[0] = ce
	slot[1] = ptr
}

const CACHE_SPECIAL = 1 << 0

func IS_SPECIAL_CACHE_VAL(ptr *ZendConstant) int { return uintptr_t(ptr) & CACHE_SPECIAL }
func ENCODE_SPECIAL_CACHE_NUM(num __auto__) any {
	return any(uintptr_t(num)<<1 | CACHE_SPECIAL)
}
func DECODE_SPECIAL_CACHE_NUM(ptr *ZendConstant) int { return uintptr_t(ptr) >> 1 }
func ENCODE_SPECIAL_CACHE_PTR(ptr __auto__) any {
	return any(uintptr_t(ptr) | CACHE_SPECIAL)
}
func DECODE_SPECIAL_CACHE_PTR(ptr __auto__) any {
	return any(uintptr_t(ptr) & ^CACHE_SPECIAL)
}
func SKIP_EXT_OPLINE(opline __auto__) {
	for UNEXPECTED(opline.opcode >= ZEND_EXT_STMT && opline.opcode <= ZEND_TICKS) {
		opline--
	}
}
func ZEND_CLASS_HAS_TYPE_HINTS(ce *ZendClassEntry) bool {
	return (ce.GetCeFlags() & ZEND_ACC_HAS_TYPE_HINTS) == ZEND_ACC_HAS_TYPE_HINTS
}
func ZEND_REF_ADD_TYPE_SOURCE(ref *ZendReference, source *ZendPropertyInfo) {
	ZendRefAddTypeSource(&ZEND_REF_TYPE_SOURCES(ref), source)
}
func ZEND_REF_DEL_TYPE_SOURCE(ref *ZendReference, source *ZendPropertyInfo) {
	ZendRefDelTypeSource(&ZEND_REF_TYPE_SOURCES(ref), source)
}

// #define ZEND_REF_FOREACH_TYPE_SOURCES(ref,prop) do { zend_property_info_source_list * _source_list = & ZEND_REF_TYPE_SOURCES ( ref ) ; zend_property_info * * _prop , * * _end ; zend_property_info_list * _list ; if ( _source_list -> ptr ) { if ( ZEND_PROPERTY_INFO_SOURCE_IS_LIST ( _source_list -> list ) ) { _list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST ( _source_list -> list ) ; _prop = _list -> ptr ; _end = _list -> ptr + _list -> num ; } else { _prop = & _source_list -> ptr ; _end = _prop + 1 ; } for ( ; _prop < _end ; _prop ++ ) { prop = * _prop ;

// #define ZEND_REF_FOREACH_TYPE_SOURCES_END() } } } while ( 0 )

// Source: <Zend/zend_execute.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

const ZEND_INTENSIVE_DEBUGGING = 0

// # include < stdio . h >

// # include < signal . h >

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_API.h"

// # include "zend_ptr_stack.h"

// # include "zend_constants.h"

// # include "zend_extensions.h"

// # include "zend_ini.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "zend_closures.h"

// # include "zend_generators.h"

// # include "zend_vm.h"

// # include "zend_dtrace.h"

// # include "zend_inheritance.h"

// # include "zend_type_info.h"

/* Virtual current working directory support */

// # include "zend_virtual_cwd.h"

const EXECUTE_DATA_D = zend_execute_data * execute_data
const EXECUTE_DATA_C EXECUTE_DATA_D = execute_data

// #define EXECUTE_DATA_DC       , EXECUTE_DATA_D

// #define EXECUTE_DATA_CC       , EXECUTE_DATA_C

// #define NO_EXECUTE_DATA_CC       , NULL

// #define OPLINE_D       const zend_op * opline

const OPLINE_C *ZendOp = opline

// #define OPLINE_DC       , OPLINE_D

// #define OPLINE_CC       , OPLINE_C

const _CONST_CODE = 0
const _TMP_CODE = 1
const _VAR_CODE = 2
const _UNUSED_CODE = 3
const _CV_CODE = 4

type IncdecT func(*Zval) int

func GetZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtr(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrDeref(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrUndef(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *Zval {
	return _getOpDataZvalPtrR(op_type, node, should_free, EXECUTE_DATA_C, OPLINE_C)
}
func GetOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *Zval {
	return _getOpDataZvalPtrDerefR(op_type, node, should_free, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func GetZvalPtrPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func GetObjZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtr(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetObjZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtrUndef(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func RETURN_VALUE_USED(opline *ZendOp) bool {
	return opline.GetResultType() != IS_UNUSED
}
func ZifPass(execute_data *ZendExecuteData, return_value *Zval) {}

var ZendPassFunction ZendInternalFunction = ZendInternalFunction{ZEND_INTERNAL_FUNCTION, {0, 0, 0}, 0, nil, nil, nil, 0, 0, nil, ZifPass, nil, {nil, nil, nil, nil}}

func FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op *Zval, result *Zval) {
	var __container_to_free *Zval = free_op
	if UNEXPECTED(__container_to_free != nil) && EXPECTED(Z_REFCOUNTED_P(__container_to_free)) {
		var __ref *ZendRefcounted = Z_COUNTED_P(__container_to_free)
		if UNEXPECTED(GC_DELREF(__ref) == 0) {
			var __zv *Zval = result
			if EXPECTED(Z_TYPE_P(__zv) == IS_INDIRECT) {
				ZVAL_COPY(__zv, Z_INDIRECT_P(__zv))
			}
			RcDtorFunc(__ref)
		}
	}
}
func FREE_OP(should_free *Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func FREE_UNFETCHED_OP(type_ ZendUchar, var_ uint32) {
	if (type_ & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(EX_VAR(var_))
	}
}
func FREE_OP_VAR_PTR(should_free *Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func CV_DEF_OF(i __auto__) __auto__ { return EX(func_).op_array.vars[i] }

const ZEND_VM_STACK_PAGE_SLOTS = 16 * 1024
const ZEND_VM_STACK_PAGE_SIZE = ZEND_VM_STACK_PAGE_SLOTS * b.SizeOf("zval")

func ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size int, page_size int) int {
	return size + ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval") + (page_size-1) & ^(page_size-1)
}
func ZendVmStackNewPage(size int, prev ZendVmStack) ZendVmStack {
	var page ZendVmStack = ZendVmStack(Emalloc(size))
	page.SetTop(ZEND_VM_STACK_ELEMENTS(page))
	page.SetEnd((*Zval)((*byte)(page + size)))
	page.SetPrev(prev)
	return page
}
func ZendVmStackInit() {
	ExecutorGlobals.SetVmStackPageSize(ZEND_VM_STACK_PAGE_SIZE)
	ExecutorGlobals.SetVmStack(ZendVmStackNewPage(ZEND_VM_STACK_PAGE_SIZE, nil))
	ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStack().GetTop())
	ExecutorGlobals.SetVmStackEnd(ExecutorGlobals.GetVmStack().GetEnd())
}
func ZendVmStackInitEx(page_size int) {
	/* page_size must be a power of 2 */

	ZEND_ASSERT(page_size > 0 && (page_size&page_size-1) == 0)
	ExecutorGlobals.SetVmStackPageSize(page_size)
	ExecutorGlobals.SetVmStack(ZendVmStackNewPage(page_size, nil))
	ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStack().GetTop())
	ExecutorGlobals.SetVmStackEnd(ExecutorGlobals.GetVmStack().GetEnd())
}
func ZendVmStackDestroy() {
	var stack ZendVmStack = ExecutorGlobals.GetVmStack()
	for stack != nil {
		var p ZendVmStack = stack.GetPrev()
		Efree(stack)
		stack = p
	}
}
func ZendVmStackExtend(size int) any {
	var stack ZendVmStack
	var ptr any
	stack = ExecutorGlobals.GetVmStack()
	stack.SetTop(ExecutorGlobals.GetVmStackTop())
	stack = ZendVmStackNewPage(b.CondF(EXPECTED(size < ExecutorGlobals.GetVmStackPageSize()-ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval")), func() int { return ExecutorGlobals.GetVmStackPageSize() }, func() int { return ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size, ExecutorGlobals.GetVmStackPageSize()) }), stack)
	ExecutorGlobals.SetVmStack(stack)
	ptr = stack.GetTop()
	ExecutorGlobals.SetVmStackTop(any((*byte)(ptr) + size))
	ExecutorGlobals.SetVmStackEnd(stack.GetEnd())
	return ptr
}
func ZendGetCompiledVariableValue(execute_data *ZendExecuteData, var_ uint32) *Zval {
	return EX_VAR(var_)
}
func _getZvalPtrTmp(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	ZEND_ASSERT(Z_TYPE_P(ret) != IS_REFERENCE)
	return ret
}
func _getZvalPtrVar(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	return ret
}
func _getZvalPtrVarDeref(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	ZVAL_DEREF(ret)
	return ret
}
func ZvalUndefinedCv(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	if EXPECTED(ExecutorGlobals.GetException() == nil) {
		var cv *ZendString = CV_DEF_OF(EX_VAR_TO_NUM(var_))
		ZendError(E_NOTICE, "Undefined variable: %s", ZSTR_VAL(cv))
	}
	return &(ExecutorGlobals.GetUninitializedZval())
}
func _zvalUndefinedOp1(EXECUTE_DATA_D) *Zval {
	return ZvalUndefinedCv(EX(opline).op1.var_, EXECUTE_DATA_C)
}
func _zvalUndefinedOp2(EXECUTE_DATA_D) *Zval {
	return ZvalUndefinedCv(EX(opline).op2.var_, EXECUTE_DATA_C)
}
func ZVAL_UNDEFINED_OP1() *Zval { return _zvalUndefinedOp1(EXECUTE_DATA_C) }
func ZVAL_UNDEFINED_OP2() *Zval { return _zvalUndefinedOp2(EXECUTE_DATA_C) }
func _getZvalCvLookup(ptr *Zval, var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	switch type_ {
	case BP_VAR_R:

	case BP_VAR_UNSET:
		ptr = ZvalUndefinedCv(var_, EXECUTE_DATA_C)
		break
	case BP_VAR_IS:
		ptr = &(ExecutorGlobals.GetUninitializedZval())
		break
	case BP_VAR_RW:
		ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	case BP_VAR_W:
		ZVAL_NULL(ptr)
		break
	}
	return ptr
}
func _getZvalPtrCv(var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if UNEXPECTED(Z_TYPE_P(ret) == IS_UNDEF) {
		if type_ == BP_VAR_W {
			ZVAL_NULL(ret)
		} else {
			return _getZvalCvLookup(ret, var_, type_, EXECUTE_DATA_C)
		}
	}
	return ret
}
func _getZvalPtrCvDeref(var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if UNEXPECTED(Z_TYPE_P(ret) == IS_UNDEF) {
		if type_ == BP_VAR_W {
			ZVAL_NULL(ret)
			return ret
		} else {
			return _getZvalCvLookup(ret, var_, type_, EXECUTE_DATA_C)
		}
	}
	ZVAL_DEREF(ret)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_R(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if UNEXPECTED(Z_TYPE_P(ret) == IS_UNDEF) {
		return ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	}
	return ret
}
func _get_zval_ptr_cv_deref_BP_VAR_R(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if UNEXPECTED(Z_TYPE_P(ret) == IS_UNDEF) {
		return ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	}
	ZVAL_DEREF(ret)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_IS(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if UNEXPECTED(Z_TYPE_P(ret) == IS_UNDEF) {
		ZVAL_NULL(ret)
		ZvalUndefinedCv(var_, EXECUTE_DATA_C)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if Z_TYPE_P(ret) == IS_UNDEF {
		ZVAL_NULL(ret)
	}
	return ret
}
func _getZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCvDeref(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_deref_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return EX_VAR(node.GetVar())
		} else {
			return nil
		}
	}
}
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if EXPECTED(Z_TYPE_P(ret) == IS_INDIRECT) {
		*should_free = nil
		ret = Z_INDIRECT_P(ret)
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_CV {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
	} else {
		ZEND_ASSERT(op_type == IS_VAR)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	}
}
func _getObjZvalPtr(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &EX(This)
	}
	return GetZvalPtr(op_type, op, should_free, type_)
}
func _getObjZvalPtrUndef(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &EX(This)
	}
	return GetZvalPtrUndef(op_type, op, should_free, type_)
}
func _getObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &EX(This)
	}
	return GetZvalPtrPtr(op_type, node, should_free, type_)
}
func ZendAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval) {
	var ref *ZendReference
	if EXPECTED(!(Z_ISREF_P(value_ptr))) {
		ZVAL_NEW_REF(value_ptr, value_ptr)
	} else if UNEXPECTED(variable_ptr == value_ptr) {
		return
	}
	ref = Z_REF_P(value_ptr)
	GC_ADDREF(ref)
	if Z_REFCOUNTED_P(variable_ptr) {
		var garbage *ZendRefcounted = Z_COUNTED_P(variable_ptr)
		if GC_DELREF(garbage) == 0 {
			ZVAL_REF(variable_ptr, ref)
			RcDtorFunc(garbage)
			return
		} else {
			GcCheckPossibleRoot(garbage)
		}
	}
	ZVAL_REF(variable_ptr, ref)
}
func ZendAssignToTypedPropertyReference(prop_info *ZendPropertyInfo, prop *Zval, value_ptr *Zval, _ EXECUTE_DATA_D) *Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, EX_USES_STRICT_TYPES()) == 0 {
		return &(ExecutorGlobals.GetUninitializedZval())
	}
	if Z_ISREF_P(prop) {
		ZEND_REF_DEL_TYPE_SOURCE(Z_REF_P(prop), prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZEND_REF_ADD_TYPE_SOURCE(Z_REF_P(prop), prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	ZendError(E_NOTICE, "Only variables should be assigned by reference")
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		return &(ExecutorGlobals.GetUninitializedZval())
	}

	/* Use IS_TMP_VAR instead of IS_VAR to avoid ISREF check */

	Z_TRY_ADDREF_P(value_ptr)
	return ZendAssignToVariable(variable_ptr, value_ptr, IS_TMP_VAR, EX_USES_STRICT_TYPES())
}
func ZendFormatType(type_ ZendType, part1 **byte, part2 **byte) {
	if ZEND_TYPE_ALLOW_NULL(type_) {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if ZEND_TYPE_IS_CLASS(type_) {
		if ZEND_TYPE_IS_CE(type_) {
			*part2 = ZSTR_VAL(ZEND_TYPE_CE(type_).GetName())
		} else {
			*part2 = ZSTR_VAL(ZEND_TYPE_NAME(type_))
		}
	} else {
		*part2 = ZendGetTypeByConst(ZEND_TYPE_CODE(type_))
	}
}
func ZendThrowAutoInitInPropError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAutoInitInRefError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAccessUninitPropByRefError(prop *ZendPropertyInfo) {
	ZendThrowError(nil, "Cannot access uninitialized non-nullable property %s::$%s by reference", ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()))
}

/* this should modify object only if it's empty */

func MakeRealObject(object *Zval, property *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	var obj *ZendObject
	var ref *Zval = nil
	if Z_ISREF_P(object) {
		ref = object
		object = Z_REFVAL_P(object)
	}
	if UNEXPECTED(Z_TYPE_P(object) > IS_FALSE && (Z_TYPE_P(object) != IS_STRING || Z_STRLEN_P(object) != 0)) {
		if opline.GetOp1Type() != IS_VAR || EXPECTED(!(Z_ISERROR_P(object))) {
			var tmp_property_name *ZendString
			var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
			if opline.GetOpcode() == ZEND_PRE_INC_OBJ || opline.GetOpcode() == ZEND_PRE_DEC_OBJ || opline.GetOpcode() == ZEND_POST_INC_OBJ || opline.GetOpcode() == ZEND_POST_DEC_OBJ {
				ZendError(E_WARNING, "Attempt to increment/decrement property '%s' of non-object", ZSTR_VAL(property_name))
			} else if opline.GetOpcode() == ZEND_FETCH_OBJ_W || opline.GetOpcode() == ZEND_FETCH_OBJ_RW || opline.GetOpcode() == ZEND_FETCH_OBJ_FUNC_ARG || opline.GetOpcode() == ZEND_ASSIGN_OBJ_REF {
				ZendError(E_WARNING, "Attempt to modify property '%s' of non-object", ZSTR_VAL(property_name))
			} else {
				ZendError(E_WARNING, "Attempt to assign property '%s' of non-object", ZSTR_VAL(property_name))
			}
			ZendTmpStringRelease(tmp_property_name)
		}
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return nil
	}
	if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(Z_REF_P(ref)) {
		if UNEXPECTED(zend_verify_ref_stdClass_assignable(Z_REF_P(ref)) == 0) {
			if RETURN_VALUE_USED(opline) {
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
			return nil
		}
	}
	ZvalPtrDtorNogc(object)
	ObjectInit(object)
	Z_ADDREF_P(object)
	obj = Z_OBJ_P(object)
	ZendError(E_WARNING, "Creating default object from empty value")
	if GC_REFCOUNT(obj) == 1 {

		/* the enclosing container was deleted, obj is unreferenced */

		OBJ_RELEASE(obj)
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return nil
	}
	Z_DELREF_P(object)
	return object
}
func ZendVerifyTypeErrorCommon(zf *ZendFunction, arg_info *ZendArgInfo, ce *ZendClassEntry, value *Zval, fname **byte, fsep **byte, fclass **byte, need_msg **byte, need_kind **byte, need_or_null **byte, given_msg **byte, given_kind **byte) {
	var is_interface ZendBool = 0
	*fname = ZSTR_VAL(zf.GetFunctionName())
	if zf.GetScope() != nil {
		*fsep = "::"
		*fclass = ZSTR_VAL(zf.GetScope().GetName())
	} else {
		*fsep = ""
		*fclass = ""
	}
	if ZEND_TYPE_IS_CLASS(arg_info.GetType()) {
		if ce != nil {
			if (ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0 {
				*need_msg = "implement interface "
				is_interface = 1
			} else {
				*need_msg = "be an instance of "
			}
			*need_kind = ZSTR_VAL(ce.GetName())
		} else {

			/* We don't know whether it's a class or interface, assume it's a class */

			*need_msg = "be an instance of "
			*need_kind = ZSTR_VAL(ZEND_TYPE_NAME(arg_info.GetType()))
		}
	} else {
		switch ZEND_TYPE_CODE(arg_info.GetType()) {
		case IS_OBJECT:
			*need_msg = "be an "
			*need_kind = "object"
			break
		case IS_CALLABLE:
			*need_msg = "be callable"
			*need_kind = ""
			break
		case IS_ITERABLE:
			*need_msg = "be iterable"
			*need_kind = ""
			break
		default:
			*need_msg = "be of the type "
			*need_kind = ZendGetTypeByConst(ZEND_TYPE_CODE(arg_info.GetType()))
			break
		}
	}
	if ZEND_TYPE_ALLOW_NULL(arg_info.GetType()) {
		if is_interface != 0 {
			*need_or_null = " or be null"
		} else {
			*need_or_null = " or null"
		}
	} else {
		*need_or_null = ""
	}
	if value != nil {
		if ZEND_TYPE_IS_CLASS(arg_info.GetType()) && Z_TYPE_P(value) == IS_OBJECT {
			*given_msg = "instance of "
			*given_kind = ZSTR_VAL(Z_OBJCE_P(value).GetName())
		} else {
			*given_msg = ZendZvalTypeName(value)
			*given_kind = ""
		}
	} else {
		*given_msg = "none"
		*given_kind = ""
	}
}
func ZendVerifyArgError(zf *ZendFunction, arg_info *ZendArgInfo, arg_num int, ce *ZendClassEntry, value *Zval) {
	var ptr *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData().GetPrevExecuteData()
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	if ExecutorGlobals.GetException() != nil {

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

		return

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

	}
	if value != nil {
		ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
		if zf.GetCommonType() == ZEND_USER_FUNCTION {
			if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ZSTR_VAL(ptr.GetFunc().GetOpArray().GetFilename()), ptr.GetOpline().GetLineno())
			} else {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
			}
		} else {
			ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
		}
	} else {
		ZendMissingArgError(ptr)
	}
}
func IsNullConstant(scope *ZendClassEntry, default_value *Zval) int {
	if Z_TYPE_P(default_value) == IS_CONSTANT_AST {
		var constant Zval
		ZVAL_COPY(&constant, default_value)
		if UNEXPECTED(ZvalUpdateConstantEx(&constant, scope) != SUCCESS) {
			return 0
		}
		if Z_TYPE(constant) == IS_NULL {
			return 1
		}
		ZvalPtrDtorNogc(&constant)
	}
	return 0
}
func ZendVerifyWeakScalarTypeHint(type_hint ZendUchar, arg *Zval) ZendBool {
	switch type_hint {
	case _IS_BOOL:
		var dest ZendBool
		if ZendParseArgBoolWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_BOOL(arg, dest)
		return 1
	case IS_LONG:
		var dest ZendLong
		if ZendParseArgLongWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_LONG(arg, dest)
		return 1
	case IS_DOUBLE:
		var dest float64
		if ZendParseArgDoubleWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_DOUBLE(arg, dest)
		return 1
	case IS_STRING:
		var dest *ZendString

		/* on success "arg" is converted to IS_STRING */

		return ZendParseArgStrWeak(arg, &dest)

	/* on success "arg" is converted to IS_STRING */

	default:
		return 0
	}
}
func ZendVerifyScalarTypeHint(type_hint ZendUchar, arg *Zval, strict ZendBool) ZendBool {
	if UNEXPECTED(strict != 0) {

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

		if type_hint != IS_DOUBLE || Z_TYPE_P(arg) != IS_LONG {
			return 0
		}

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	} else if UNEXPECTED(Z_TYPE_P(arg) == IS_NULL) {

		/* NULL may be accepted only by nullable hints (this is already checked) */

		return 0

		/* NULL may be accepted only by nullable hints (this is already checked) */

	}
	return ZendVerifyWeakScalarTypeHint(type_hint, arg)
}
func ZendVerifyPropertyTypeError(info *ZendPropertyInfo, property *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte

	/* we _may_ land here in case reading already errored and runtime cache thus has not been updated (i.e. it contains a valid but unrelated info) */

	if ExecutorGlobals.GetException() != nil {
		return
	}

	// TODO Switch to a more standard error message?

	ZendFormatType(info.GetType(), &prop_type1, &prop_type2)
	void(prop_type1)
	if ZEND_TYPE_IS_CLASS(info.GetType()) {
		ZendTypeError("Typed property %s::$%s must be an instance of %s%s, %s used", ZSTR_VAL(info.GetCe().GetName()), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(ZEND_TYPE_ALLOW_NULL(info.GetType()), " or null", ""), b.CondF(Z_TYPE_P(property) == IS_OBJECT, func() []byte { return ZSTR_VAL(Z_OBJCE_P(property).GetName()) }, func() *byte { return ZendGetTypeByConst(Z_TYPE_P(property)) }))
	} else {
		ZendTypeError("Typed property %s::$%s must be %s%s, %s used", ZSTR_VAL(info.GetCe().GetName()), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(ZEND_TYPE_ALLOW_NULL(info.GetType()), " or null", ""), b.CondF(Z_TYPE_P(property) == IS_OBJECT, func() []byte { return ZSTR_VAL(Z_OBJCE_P(property).GetName()) }, func() *byte { return ZendGetTypeByConst(Z_TYPE_P(property)) }))
	}
}
func ZendResolveClassType(type_ *ZendType, self_ce *ZendClassEntry) ZendBool {
	var ce *ZendClassEntry
	var name *ZendString = ZEND_TYPE_NAME(*type_)
	if ZendStringEqualsLiteralCi(name, "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if UNEXPECTED((self_ce.GetCeFlags() & ZEND_ACC_TRAIT) != 0) {
			ZendThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", b.Cond(ZEND_TYPE_ALLOW_NULL(*type_), " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if ZendStringEqualsLiteralCi(name, "parent") {
		if UNEXPECTED(!(self_ce.parent)) {
			ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return 0
		}
		ce = self_ce.parent
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if UNEXPECTED(ce == nil) {
			return 0
		}
	}
	ZendStringRelease(name)
	*type_ = ZEND_TYPE_ENCODE_CE(ce, ZEND_TYPE_ALLOW_NULL(*type_))
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	ZEND_ASSERT(!(Z_ISREF_P(property)))
	if ZEND_TYPE_IS_CLASS(info.GetType()) {
		if UNEXPECTED(Z_TYPE_P(property) != IS_OBJECT) {
			return Z_TYPE_P(property) == IS_NULL && ZEND_TYPE_ALLOW_NULL(info.GetType())
		}
		if UNEXPECTED(!(ZEND_TYPE_IS_CE(info.GetType()))) && UNEXPECTED(ZendResolveClassType(&info.type_, info.GetCe()) == 0) {
			return 0
		}
		return InstanceofFunction(Z_OBJCE_P(property), ZEND_TYPE_CE(info.GetType()))
	}
	ZEND_ASSERT(ZEND_TYPE_CODE(info.GetType()) != IS_CALLABLE)
	if EXPECTED(ZEND_TYPE_CODE(info.GetType()) == Z_TYPE_P(property)) {
		return 1
	} else if EXPECTED(Z_TYPE_P(property) == IS_NULL) {
		return ZEND_TYPE_ALLOW_NULL(info.GetType())
	} else if ZEND_TYPE_CODE(info.GetType()) == _IS_BOOL && EXPECTED(Z_TYPE_P(property) == IS_FALSE || Z_TYPE_P(property) == IS_TRUE) {
		return 1
	} else if ZEND_TYPE_CODE(info.GetType()) == IS_ITERABLE {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(ZEND_TYPE_CODE(info.GetType()), property, strict)
	}
}
func IZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	if IZendCheckPropertyType(info, property, strict) != 0 {
		return 1
	}
	ZendVerifyPropertyTypeError(info, property)
	return 0
}
func ZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *Zval, value *Zval, _ EXECUTE_DATA_D) *Zval {
	var tmp Zval
	ZVAL_DEREF(value)
	ZVAL_COPY(&tmp, value)
	if UNEXPECTED(IZendVerifyPropertyType(info, &tmp, EX_USES_STRICT_TYPES()) == 0) {
		ZvalPtrDtor(&tmp)
		return &(ExecutorGlobals.GetUninitializedZval())
	}
	return ZendAssignToVariable(property_val, &tmp, IS_TMP_VAR, EX_USES_STRICT_TYPES())
}
func ZendCheckType(type_ ZendType, arg *Zval, ce **ZendClassEntry, cache_slot *any, default_value *Zval, scope *ZendClassEntry, is_return_type ZendBool) ZendBool {
	var ref *ZendReference = nil
	if !(ZEND_TYPE_IS_SET(type_)) {
		return 1
	}
	if UNEXPECTED(Z_ISREF_P(arg)) {
		ref = Z_REF_P(arg)
		arg = Z_REFVAL_P(arg)
	}
	if ZEND_TYPE_IS_CLASS(type_) {
		if EXPECTED(*cache_slot) {
			*ce = (*ZendClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass(ZEND_TYPE_NAME(type_), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if UNEXPECTED((*ce) == nil) {
				return Z_TYPE_P(arg) == IS_NULL && (ZEND_TYPE_ALLOW_NULL(type_) || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if EXPECTED(Z_TYPE_P(arg) == IS_OBJECT) {
			return InstanceofFunction(Z_OBJCE_P(arg), *ce)
		}
		return Z_TYPE_P(arg) == IS_NULL && (ZEND_TYPE_ALLOW_NULL(type_) || default_value != nil && IsNullConstant(scope, default_value) != 0)
	} else if EXPECTED(ZEND_TYPE_CODE(type_) == Z_TYPE_P(arg)) {
		return 1
	}
	if Z_TYPE_P(arg) == IS_NULL && (ZEND_TYPE_ALLOW_NULL(type_) || default_value != nil && IsNullConstant(scope, default_value) != 0) {

		/* Null passed to nullable type */

		return 1

		/* Null passed to nullable type */

	}
	if ZEND_TYPE_CODE(type_) == IS_CALLABLE {
		return ZendIsCallable(arg, IS_CALLABLE_CHECK_SILENT, nil)
	} else if ZEND_TYPE_CODE(type_) == IS_ITERABLE {
		return ZendIsIterable(arg)
	} else if ZEND_TYPE_CODE(type_) == _IS_BOOL && EXPECTED(Z_TYPE_P(arg) == IS_FALSE || Z_TYPE_P(arg) == IS_TRUE) {
		return 1
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(ZEND_TYPE_CODE(type_), arg, b.CondF(is_return_type != 0, func() bool { return ZEND_RET_USES_STRICT_TYPES() }, func() bool { return ZEND_ARG_USES_STRICT_TYPES() }))
	}
}
func ZendVerifyArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	if EXPECTED(arg_num <= zf.GetNumArgs()) {
		cur_arg_info = &zf.common.arg_info[arg_num-1]
	} else if UNEXPECTED((zf.GetFnFlags() & ZEND_ACC_VARIADIC) != 0) {
		cur_arg_info = &zf.common.arg_info[zf.GetNumArgs()]
	} else {
		return 1
	}
	ce = nil
	if UNEXPECTED(ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0) {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyRecvArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = &zf.common.arg_info[arg_num-1]
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num <= zf.GetNumArgs())
	cur_arg_info = &zf.common.arg_info[arg_num-1]
	ce = nil
	if UNEXPECTED(ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0) {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num > zf.GetNumArgs())
	ZEND_ASSERT((zf.GetFnFlags() & ZEND_ACC_VARIADIC) != 0)
	cur_arg_info = &zf.common.arg_info[zf.GetNumArgs()]
	ce = nil
	if UNEXPECTED(ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0) {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyInternalArgTypes(fbc *ZendFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	var p *Zval = ZEND_CALL_ARG(call, 1)
	var dummy_cache_slot any
	for i = 0; i < num_args; i++ {
		dummy_cache_slot = nil
		if UNEXPECTED(ZendVerifyArgType(fbc, i+1, p, nil, &dummy_cache_slot) == 0) {
			ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
			return 0
		}
		p++
	}
	return 1
}
func ZendMissingArgError(execute_data *ZendExecuteData) {
	var ptr *ZendExecuteData = EX(prev_execute_data)
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return ZSTR_VAL(EX(func_).common.scope.name) }, ""), b.Cond(EX(func_).common.scope, "::", ""), ZSTR_VAL(EX(func_).common.function_name), EX_NUM_ARGS(), ZSTR_VAL(ptr.GetFunc().GetOpArray().GetFilename()), ptr.GetOpline().GetLineno(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	} else {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return ZSTR_VAL(EX(func_).common.scope.name) }, ""), b.Cond(EX(func_).common.scope, "::", ""), ZSTR_VAL(EX(func_).common.function_name), EX_NUM_ARGS(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	}
}
func ZendVerifyReturnError(zf *ZendFunction, ce *ZendClassEntry, value *Zval) {
	var arg_info *ZendArgInfo = &zf.common.arg_info[-1]
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
	ZendTypeError("Return value of %s%s%s() must %s%s%s, %s%s returned", fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
}
func ZendVerifyReturnType(zf *ZendFunction, ret *Zval, cache_slot *any) {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	var ce *ZendClassEntry = nil
	if UNEXPECTED(ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0) {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf *ZendFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ZEND_TYPE_IS_SET(ret_info.GetType()) && UNEXPECTED(ZEND_TYPE_CODE(ret_info.GetType()) != IS_VOID) {
		var ce *ZendClassEntry = nil
		if ZEND_TYPE_IS_CLASS(ret_info.GetType()) {
			if EXPECTED(*cache_slot) {
				ce = (*ZendClassEntry)(*cache_slot)
			} else {
				ce = ZendFetchClass(ZEND_TYPE_NAME(ret_info.GetType()), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					*cache_slot = any(ce)
				}
			}
		}
		ZendVerifyReturnError(zf, ce, nil)
		return 0
	}
	return 1
}
func ZendUseObjectAsArray() {
	ZendThrowError(nil, "Cannot use object as array")
}
func ZendIllegalOffset() {
	ZendError(E_WARNING, "Illegal offset type")
}
func ZendAssignToObjectDim(object *Zval, dim *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	Z_OBJ_HT_P(object).GetWriteDimension()(object, dim, value)
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
}
func ZendBinaryOp(ret *Zval, op1 *Zval, op2 *Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *Zval, property *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var free_op_data1 ZendFreeOp
	var value *Zval
	var z *Zval
	var rv Zval
	var res Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if b.Assign(&z, Z_OBJ_HT_P(object).GetReadDimension()(object, property, BP_VAR_R, &rv)) != nil {
		if Z_TYPE_P(z) == IS_OBJECT && Z_OBJ_HT_P(z).GetGet() != nil {
			var rv2 Zval
			var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
			if z == &rv {
				ZvalPtrDtor(&rv)
			}
			ZVAL_COPY_VALUE(z, value)
		}
		if ZendBinaryOp(&res, z, value, OPLINE_C) == SUCCESS {
			Z_OBJ_HT_P(object).GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
		}
		ZvalPtrDtor(&res)
	} else {
		ZendUseObjectAsArray()
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
	}
	FREE_OP(free_op_data1)
}
func ZendBinaryAssignOpTypedRef(ref *ZendReference, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && Z_TYPE(ref.GetVal()) == IS_STRING {
		ConcatFunction(&ref.val, &ref.val, value)
		ZEND_ASSERT(Z_TYPE(ref.GetVal()) == IS_STRING && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, &ref.val, value, OPLINE_C)
	if EXPECTED(ZendVerifyRefAssignableZval(ref, &z_copy, EX_USES_STRICT_TYPES()) != 0) {
		ZvalPtrDtor(&ref.val)
		ZVAL_COPY_VALUE(&ref.val, &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && Z_TYPE_P(zptr) == IS_STRING {
		ConcatFunction(zptr, zptr, value)
		ZEND_ASSERT(Z_TYPE_P(zptr) == IS_STRING && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, OPLINE_C)
	if EXPECTED(ZendVerifyPropertyType(prop_info, &z_copy, EX_USES_STRICT_TYPES()) != 0) {
		ZvalPtrDtor(zptr)
		ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *Zval, type_ int, _ EXECUTE_DATA_D) ZendLong {
	var offset ZendLong
try_again:
	if UNEXPECTED(Z_TYPE_P(dim) != IS_LONG) {
		switch Z_TYPE_P(dim) {
		case IS_STRING:
			if IS_LONG == IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				ZendError(E_WARNING, "Illegal string offset '%s'", Z_STRVAL_P(dim))
			}
			break
		case IS_UNDEF:
			ZVAL_UNDEFINED_OP2()
		case IS_DOUBLE:

		case IS_NULL:

		case IS_FALSE:

		case IS_TRUE:
			ZendError(E_NOTICE, "String offset cast occurred")
			break
		case IS_REFERENCE:
			dim = Z_REFVAL_P(dim)
			goto try_again
		default:
			ZendIllegalOffset()
			break
		}
		offset = ZvalGetLongFunc(dim)
	} else {
		offset = Z_LVAL_P(dim)
	}
	return offset
}
func ZendWrongStringOffset(EXECUTE_DATA_D) {
	var msg *byte = nil
	var opline *ZendOp = EX(opline)
	var end *ZendOp
	var var_ uint32
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		return
	}
	switch opline.GetOpcode() {
	case ZEND_ASSIGN_OP:

	case ZEND_ASSIGN_DIM_OP:

	case ZEND_ASSIGN_OBJ_OP:

	case ZEND_ASSIGN_STATIC_PROP_OP:
		msg = "Cannot use assign-op operators with string offsets"
		break
	case ZEND_FETCH_DIM_W:

	case ZEND_FETCH_DIM_RW:

	case ZEND_FETCH_DIM_FUNC_ARG:

	case ZEND_FETCH_DIM_UNSET:

	case ZEND_FETCH_LIST_W:

		/* TODO: Encode the "reason" into opline->extended_value??? */

		var_ = opline.GetResult().GetVar()
		opline++
		end = ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetOpcodes() + ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetLast()
		for opline < end {
			if opline.GetOp1Type() == IS_VAR && opline.GetOp1().GetVar() == var_ {
				switch opline.GetOpcode() {
				case ZEND_FETCH_OBJ_W:

				case ZEND_FETCH_OBJ_RW:

				case ZEND_FETCH_OBJ_FUNC_ARG:

				case ZEND_FETCH_OBJ_UNSET:

				case ZEND_ASSIGN_OBJ:

				case ZEND_ASSIGN_OBJ_OP:

				case ZEND_ASSIGN_OBJ_REF:
					msg = "Cannot use string offset as an object"
					break
				case ZEND_FETCH_DIM_W:

				case ZEND_FETCH_DIM_RW:

				case ZEND_FETCH_DIM_FUNC_ARG:

				case ZEND_FETCH_DIM_UNSET:

				case ZEND_FETCH_LIST_W:

				case ZEND_ASSIGN_DIM:

				case ZEND_ASSIGN_DIM_OP:
					msg = "Cannot use string offset as an array"
					break
				case ZEND_ASSIGN_STATIC_PROP_OP:

				case ZEND_ASSIGN_OP:
					msg = "Cannot use assign-op operators with string offsets"
					break
				case ZEND_PRE_INC_OBJ:

				case ZEND_PRE_DEC_OBJ:

				case ZEND_POST_INC_OBJ:

				case ZEND_POST_DEC_OBJ:

				case ZEND_PRE_INC:

				case ZEND_PRE_DEC:

				case ZEND_POST_INC:

				case ZEND_POST_DEC:
					msg = "Cannot increment/decrement string offsets"
					break
				case ZEND_ASSIGN_REF:

				case ZEND_ADD_ARRAY_ELEMENT:

				case ZEND_INIT_ARRAY:

				case ZEND_MAKE_REF:
					msg = "Cannot create references to/from string offsets"
					break
				case ZEND_RETURN_BY_REF:

				case ZEND_VERIFY_RETURN_TYPE:
					msg = "Cannot return string offsets by reference"
					break
				case ZEND_UNSET_DIM:

				case ZEND_UNSET_OBJ:
					msg = "Cannot unset string offsets"
					break
				case ZEND_YIELD:
					msg = "Cannot yield string offsets by reference"
					break
				case ZEND_SEND_REF:

				case ZEND_SEND_VAR_EX:

				case ZEND_SEND_FUNC_ARG:
					msg = "Only variables can be passed by reference"
					break
				case ZEND_FE_RESET_RW:
					msg = "Cannot iterate on string offsets by reference"
					break
				default:
					break
				}
				break
			}
			if opline.GetOp2Type() == IS_VAR && opline.GetOp2().GetVar() == var_ {
				ZEND_ASSERT(opline.GetOpcode() == ZEND_ASSIGN_REF)
				msg = "Cannot create references to/from string offsets"
				break
			}
			opline++
		}
		break
	default:
		break
	}
	ZEND_ASSERT(msg != nil)
	ZendThrowError(nil, "%s", msg)
}
func ZendWrongPropertyRead(property *Zval) {
	var tmp_property_name *ZendString
	var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
	ZendError(E_NOTICE, "Trying to get property '%s' of non-object", ZSTR_VAL(property_name))
	ZendTmpStringRelease(tmp_property_name)
}
func ZendDeprecatedFunction(fbc *ZendFunction) {
	ZendError(E_DEPRECATED, "Function %s%s%s() is deprecated", b.CondF1(fbc.GetScope() != nil, func() []byte { return ZSTR_VAL(fbc.GetScope().GetName()) }, ""), b.Cond(fbc.GetScope() != nil, "::", ""), ZSTR_VAL(fbc.GetFunctionName()))
}
func ZendAbstractMethod(fbc *ZendFunction) {
	ZendThrowError(nil, "Cannot call abstract method %s::%s()", ZSTR_VAL(fbc.GetScope().GetName()), ZSTR_VAL(fbc.GetFunctionName()))
}
func ZendAssignToStringOffset(str *Zval, dim *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var c ZendUchar
	var string_len int
	var offset ZendLong
	offset = ZendCheckStringOffset(dim, BP_VAR_W, EXECUTE_DATA_C)
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if offset < -ZendLong(Z_STRLEN_P(str)) {

		/* Error on negative offset */

		ZendError(E_WARNING, "Illegal string offset:  "+ZEND_LONG_FMT, offset)
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if Z_TYPE_P(value) != IS_STRING {

		/* Convert to string, just the time to pick the 1st byte */

		var tmp *ZendString = ZvalTryGetStringFunc(value)
		if UNEXPECTED(tmp == nil) {
			if UNEXPECTED(RETURN_VALUE_USED(opline)) {
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
			return
		}
		string_len = ZSTR_LEN(tmp)
		c = ZendUchar(ZSTR_VAL(tmp)[0])
		ZendStringReleaseEx(tmp, 0)
	} else {
		string_len = Z_STRLEN_P(value)
		c = ZendUchar(Z_STRVAL_P(value)[0])
	}
	if string_len == 0 {

		/* Error on empty input string */

		ZendError(E_WARNING, "Cannot assign an empty string to a string offset")
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if offset < 0 {
		offset += ZendLong(Z_STRLEN_P(str))
	}
	if int(offset >= Z_STRLEN_P(str)) != 0 {

		/* Extend string if needed */

		var old_len ZendLong = Z_STRLEN_P(str)
		ZVAL_NEW_STR(str, ZendStringExtend(Z_STR_P(str), offset+1, 0))
		memset(Z_STRVAL_P(str)+old_len, ' ', offset-old_len)
		Z_STRVAL_P(str)[offset+1] = 0
	} else if !(Z_REFCOUNTED_P(str)) {
		ZVAL_NEW_STR(str, ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
	} else if Z_REFCOUNT_P(str) > 1 {
		Z_DELREF_P(str)
		ZVAL_NEW_STR(str, ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
	} else {
		ZendStringForgetHashVal(Z_STR_P(str))
	}
	Z_STRVAL_P(str)[offset] = c
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {

		/* Return the new character */

		ZVAL_INTERNED_STR(EX_VAR(opline.GetResult().GetVar()), ZSTR_CHAR(c))

		/* Return the new character */

	}
}
func ZendGetPropNotAcceptingDouble(ref *ZendReference) *ZendPropertyInfo {
	var prop *ZendPropertyInfo
	var _source_list *ZendPropertyInfoSourceList = &ZEND_REF_TYPE_SOURCES(ref)
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = &_source_list.ptr
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			if ZEND_TYPE_CODE(prop.GetType()) != IS_DOUBLE {
				return prop
			}
		}
	}
	return nil
}
func ZendThrowIncdecRefError(ref *ZendReference, opline *ZendOp) ZendLong {
	var error_prop *ZendPropertyInfo = ZendGetPropNotAcceptingDouble(ref)

	/* Currently there should be no way for a typed reference to accept both int and double.
	 * Generalize this and the related property code once this becomes possible. */

	ZEND_ASSERT(error_prop != nil)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		ZendTypeError("Cannot increment a reference held by property %s::$%s of type %sint past its maximal value", ZSTR_VAL(error_prop.GetCe().GetName()), ZendGetUnmangledPropertyName(error_prop.GetName()), b.Cond(ZEND_TYPE_ALLOW_NULL(error_prop.GetType()), "?", ""))
		return ZEND_LONG_MAX
	} else {
		ZendTypeError("Cannot decrement a reference held by property %s::$%s of type %sint past its minimal value", ZSTR_VAL(error_prop.GetCe().GetName()), ZendGetUnmangledPropertyName(error_prop.GetName()), b.Cond(ZEND_TYPE_ALLOW_NULL(error_prop.GetType()), "?", ""))
		return ZEND_LONG_MIN
	}
}
func ZendThrowIncdecPropError(prop *ZendPropertyInfo, opline *ZendOp) ZendLong {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		ZendTypeError("Cannot increment property %s::$%s of type %s%s past its maximal value", ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MAX
	} else {
		ZendTypeError("Cannot decrement property %s::$%s of type %s%s past its minimal value", ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MIN
	}
}
func ZendIncdecTypedRef(ref *ZendReference, copy *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var tmp Zval
	var var_ptr *Zval = &ref.val
	if copy == nil {
		copy = &tmp
	}
	ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if UNEXPECTED(Z_TYPE_P(var_ptr) == IS_DOUBLE) && Z_TYPE_P(copy) == IS_LONG {
		var val ZendLong = ZendThrowIncdecRefError(ref, OPLINE_C)
		ZVAL_LONG(var_ptr, val)
	} else if UNEXPECTED(ZendVerifyRefAssignableZval(ref, var_ptr, EX_USES_STRICT_TYPES()) == 0) {
		ZvalPtrDtor(var_ptr)
		ZVAL_COPY_VALUE(var_ptr, copy)
		ZVAL_UNDEF(copy)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendIncdecTypedProp(prop_info *ZendPropertyInfo, var_ptr *Zval, copy *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var tmp Zval
	if copy == nil {
		copy = &tmp
	}
	ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if UNEXPECTED(Z_TYPE_P(var_ptr) == IS_DOUBLE) && Z_TYPE_P(copy) == IS_LONG {
		var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
		ZVAL_LONG(var_ptr, val)
	} else if UNEXPECTED(ZendVerifyPropertyType(prop_info, var_ptr, EX_USES_STRICT_TYPES()) == 0) {
		ZvalPtrDtor(var_ptr)
		ZVAL_COPY_VALUE(var_ptr, copy)
		ZVAL_UNDEF(copy)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendPreIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, _ EXECUTE_DATA_D) {
	if EXPECTED(Z_TYPE_P(prop) == IS_LONG) {
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if UNEXPECTED(Z_TYPE_P(prop) != IS_LONG) && UNEXPECTED(prop_info != nil) {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
			ZVAL_LONG(prop, val)
		}
	} else {
		for {
			if Z_ISREF_P(prop) {
				var ref *ZendReference = Z_REF_P(prop)
				prop = Z_REFVAL_P(prop)
				if UNEXPECTED(ZEND_REF_HAS_TYPE_SOURCES(ref)) {
					ZendIncdecTypedRef(ref, nil, OPLINE_C, EXECUTE_DATA_C)
					break
				}
			}
			if UNEXPECTED(prop_info != nil) {
				ZendIncdecTypedProp(prop_info, prop, nil, OPLINE_C, EXECUTE_DATA_C)
			} else if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
			break
		}
	}
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), prop)
	}
}
func ZendPostIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, _ EXECUTE_DATA_D) {
	if EXPECTED(Z_TYPE_P(prop) == IS_LONG) {
		ZVAL_LONG(EX_VAR(opline.GetResult().GetVar()), Z_LVAL_P(prop))
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if UNEXPECTED(Z_TYPE_P(prop) != IS_LONG) && UNEXPECTED(prop_info != nil) {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
			ZVAL_LONG(prop, val)
		}
	} else {
		if Z_ISREF_P(prop) {
			var ref *ZendReference = Z_REF_P(prop)
			prop = Z_REFVAL_P(prop)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, EX_VAR(opline.GetResult().GetVar()), OPLINE_C, EXECUTE_DATA_C)
				return
			}
		}
		if UNEXPECTED(prop_info != nil) {
			ZendIncdecTypedProp(prop_info, prop, EX_VAR(opline.GetResult().GetVar()), OPLINE_C, EXECUTE_DATA_C)
		} else {
			ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), prop)
			if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
		}
	}
}
func ZendPostIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, _ EXECUTE_DATA_D) {
	var rv Zval
	var obj Zval
	var z *Zval
	var z_copy Zval
	ZVAL_OBJ(&obj, Z_OBJ_P(object))
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		OBJ_RELEASE(Z_OBJ(obj))
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return
	}
	if UNEXPECTED(Z_TYPE_P(z) == IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	ZVAL_COPY_DEREF(&z_copy, z)
	ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &z_copy)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	OBJ_RELEASE(Z_OBJ(obj))
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendPreIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, _ EXECUTE_DATA_D) {
	var rv Zval
	var z *Zval
	var obj Zval
	var z_copy Zval
	ZVAL_OBJ(&obj, Z_OBJ_P(object))
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		OBJ_RELEASE(Z_OBJ(obj))
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if UNEXPECTED(Z_TYPE_P(z) == IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	ZVAL_COPY_DEREF(&z_copy, z)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &z_copy)
	}
	Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	OBJ_RELEASE(Z_OBJ(obj))
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendAssignOpOverloadedProperty(object *Zval, property *Zval, cache_slot *any, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z *Zval
	var rv Zval
	var obj Zval
	var res Zval
	ZVAL_OBJ(&obj, Z_OBJ_P(object))
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		OBJ_RELEASE(Z_OBJ(obj))
		if UNEXPECTED(RETURN_VALUE_USED(opline)) {
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if Z_TYPE_P(z) == IS_OBJECT && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	if ZendBinaryOp(&res, z, value, OPLINE_C) == SUCCESS {
		Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &res, cache_slot)
	}
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
	}
	ZvalPtrDtor(z)
	ZvalPtrDtor(&res)
	OBJ_RELEASE(Z_OBJ(obj))
}

/* Utility Functions for Extensions */

func ZendExtensionStatementHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetStatementHandler() != nil {
		extension.GetStatementHandler()(frame)
	}
}
func ZendExtensionFcallBeginHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallBeginHandler() != nil {
		extension.GetFcallBeginHandler()(frame)
	}
}
func ZendExtensionFcallEndHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallEndHandler() != nil {
		extension.GetFcallEndHandler()(frame)
	}
}
func ZendGetTargetSymbolTable(fetch_type int, _ EXECUTE_DATA_D) *HashTable {
	var ht *HashTable
	if EXPECTED((fetch_type & (ZEND_FETCH_GLOBAL_LOCK | ZEND_FETCH_GLOBAL)) != 0) {
		ht = &(ExecutorGlobals.GetSymbolTable())
	} else {
		ZEND_ASSERT((fetch_type & ZEND_FETCH_LOCAL) != 0)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			ZendRebuildSymbolTable()
		}
		ht = EX(symbol_table)
	}
	return ht
}
func ZendUndefinedOffset(lval ZendLong) {
	ZendError(E_NOTICE, "Undefined offset: "+ZEND_LONG_FMT, lval)
}
func ZendUndefinedIndex(offset *ZendString) {
	ZendError(E_NOTICE, "Undefined index: %s", ZSTR_VAL(offset))
}
func ZendUndefinedOffsetWrite(ht *HashTable, lval ZendLong) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (GC_FLAGS(ht) & IS_ARRAY_IMMUTABLE) == 0 {
		GC_ADDREF(ht)
	}
	ZendUndefinedOffset(lval)
	if (GC_FLAGS(ht)&IS_ARRAY_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedIndexWrite(ht *HashTable, offset *ZendString) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (GC_FLAGS(ht) & IS_ARRAY_IMMUTABLE) == 0 {
		GC_ADDREF(ht)
	}
	ZendUndefinedIndex(offset)
	if (GC_FLAGS(ht)&IS_ARRAY_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedMethod(ce *ZendClassEntry, method *ZendString) {
	ZendThrowError(nil, "Call to undefined method %s::%s()", ZSTR_VAL(ce.GetName()), ZSTR_VAL(method))
}
func ZendInvalidMethodCall(object *Zval, function_name *Zval) {
	ZendThrowError(nil, "Call to a member function %s() on %s", Z_STRVAL_P(function_name), ZendGetTypeByConst(Z_TYPE_P(object)))
}
func ZendNonStaticMethodCall(fbc *ZendFunction) {
	if (fbc.GetFnFlags() & ZEND_ACC_ALLOW_STATIC) != 0 {
		ZendError(E_DEPRECATED, "Non-static method %s::%s() should not be called statically", ZSTR_VAL(fbc.GetScope().GetName()), ZSTR_VAL(fbc.GetFunctionName()))
	} else {
		ZendThrowError(ZendCeError, "Non-static method %s::%s() cannot be called statically", ZSTR_VAL(fbc.GetScope().GetName()), ZSTR_VAL(fbc.GetFunctionName()))
	}
}
func ZendParamMustBeRef(func_ *ZendFunction, arg_num uint32) {
	ZendError(E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", arg_num, b.CondF1(func_.GetScope() != nil, func() []byte { return ZSTR_VAL(func_.GetScope().GetName()) }, ""), b.Cond(func_.GetScope() != nil, "::", ""), ZSTR_VAL(func_.GetFunctionName()))
}
func ZendUseScalarAsArray() {
	ZendError(E_WARNING, "Cannot use a scalar value as an array")
}
func ZendCannotAddElement() {
	ZendError(E_WARNING, "Cannot add element to the array as the next element is already occupied")
}
func ZendUseResourceAsOffset(dim *Zval) {
	ZendError(E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", Z_RES_HANDLE_P(dim), Z_RES_HANDLE_P(dim))
}
func ZendUseNewElementForString() {
	ZendThrowError(nil, "[] operator not supported for strings")
}
func ZendBinaryAssignOpDimSlow(container *Zval, dim *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	if UNEXPECTED(Z_TYPE_P(container) == IS_STRING) {
		if opline.GetOp2Type() == IS_UNUSED {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, BP_VAR_RW, EXECUTE_DATA_C)
			ZendWrongStringOffset(EXECUTE_DATA_C)
		}
	} else if EXPECTED(!(Z_ISERROR_P(container))) {
		ZendUseScalarAsArray()
	}
}
func SlowIndexConvert(ht *HashTable, dim *Zval, value *ZendValue, _ EXECUTE_DATA_D) ZendUchar {
	switch Z_TYPE_P(dim) {
	case IS_UNDEF:

		/* The array may be destroyed while throwing the notice.
		 * Temporarily increase the refcount to detect this situation. */

		if (GC_FLAGS(ht) & IS_ARRAY_IMMUTABLE) == 0 {
			GC_ADDREF(ht)
		}
		ZVAL_UNDEFINED_OP2()
		if (GC_FLAGS(ht)&IS_ARRAY_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
			ZendArrayDestroy(ht)
			return IS_NULL
		}
		if ExecutorGlobals.GetException() != nil {
			return IS_NULL
		}
	case IS_NULL:
		value.SetStr(ZSTR_EMPTY_ALLOC())
		return IS_STRING
	case IS_DOUBLE:
		value.SetLval(ZendDvalToLval(Z_DVAL_P(dim)))
		return IS_LONG
	case IS_RESOURCE:
		ZendUseResourceAsOffset(dim)
		value.SetLval(Z_RES_HANDLE_P(dim))
		return IS_LONG
	case IS_FALSE:
		value.SetLval(0)
		return IS_LONG
	case IS_TRUE:
		value.SetLval(1)
		return IS_LONG
	default:
		ZendIllegalOffset()
		return IS_NULL
	}
}
func ZendFetchDimensionAddressInner(ht *HashTable, dim *Zval, dim_type int, type_ int, _ EXECUTE_DATA_D) *Zval {
	var retval *Zval = nil
	var offset_key *ZendString
	var hval ZendUlong
try_again:
	if EXPECTED(Z_TYPE_P(dim) == IS_LONG) {
		hval = Z_LVAL_P(dim)
	num_index:
		ZEND_HASH_INDEX_FIND(ht, hval, retval, num_undef)
		return retval
	num_undef:
		switch type_ {
		case BP_VAR_R:
			ZendUndefinedOffset(hval)
		case BP_VAR_UNSET:

		case BP_VAR_IS:
			retval = &(ExecutorGlobals.GetUninitializedZval())
			break
		case BP_VAR_RW:
			if UNEXPECTED(ZendUndefinedOffsetWrite(ht, hval) == FAILURE) {
				return nil
			}
		case BP_VAR_W:
			retval = ZendHashIndexAddNew(ht, hval, &(ExecutorGlobals.GetUninitializedZval()))
			break
		}
	} else if EXPECTED(Z_TYPE_P(dim) == IS_STRING) {
		offset_key = Z_STR_P(dim)
		if ZEND_CONST_COND(dim_type != IS_CONST, 1) {
			if ZEND_HANDLE_NUMERIC(offset_key, hval) != 0 {
				goto num_index
			}
		}
	str_index:
		retval = ZendHashFindEx(ht, offset_key, ZEND_CONST_COND(dim_type == IS_CONST, 0))
		if retval != nil {

			/* support for $GLOBALS[...] */

			if UNEXPECTED(Z_TYPE_P(retval) == IS_INDIRECT) {
				retval = Z_INDIRECT_P(retval)
				if UNEXPECTED(Z_TYPE_P(retval) == IS_UNDEF) {
					switch type_ {
					case BP_VAR_R:
						ZendUndefinedIndex(offset_key)
					case BP_VAR_UNSET:

					case BP_VAR_IS:
						retval = &(ExecutorGlobals.GetUninitializedZval())
						break
					case BP_VAR_RW:
						if UNEXPECTED(ZendUndefinedIndexWrite(ht, offset_key) != 0) {
							return nil
						}
					case BP_VAR_W:
						ZVAL_NULL(retval)
						break
					}
				}
			}

			/* support for $GLOBALS[...] */

		} else {
			switch type_ {
			case BP_VAR_R:
				ZendUndefinedIndex(offset_key)
			case BP_VAR_UNSET:

			case BP_VAR_IS:
				retval = &(ExecutorGlobals.GetUninitializedZval())
				break
			case BP_VAR_RW:

				/* Key may be released while throwing the undefined index warning. */

				ZendStringAddref(offset_key)
				if UNEXPECTED(ZendUndefinedIndexWrite(ht, offset_key) == FAILURE) {
					ZendStringRelease(offset_key)
					return nil
				}
				retval = ZendHashAddNew(ht, offset_key, &(ExecutorGlobals.GetUninitializedZval()))
				ZendStringRelease(offset_key)
				break
			case BP_VAR_W:
				retval = ZendHashAddNew(ht, offset_key, &(ExecutorGlobals.GetUninitializedZval()))
				break
			}
		}
	} else if EXPECTED(Z_TYPE_P(dim) == IS_REFERENCE) {
		dim = Z_REFVAL_P(dim)
		goto try_again
	} else {
		var val ZendValue
		var t ZendUchar = SlowIndexConvert(ht, dim, &val, EXECUTE_DATA_C)
		if t == IS_STRING {
			offset_key = val.GetStr()
			goto str_index
		} else if t == IS_LONG {
			hval = val.GetLval()
			goto num_index
		} else {
			if type_ == BP_VAR_W || type_ == BP_VAR_RW {
				retval = nil
			} else {
				retval = &(ExecutorGlobals.GetUninitializedZval())
			}
		}
	}
	return retval
}
func zend_fetch_dimension_address_inner_W(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_W_CONST(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_RW(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_RW, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_RW_CONST(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_RW, EXECUTE_DATA_C)
}
func ZendFetchDimensionAddress(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, _ EXECUTE_DATA_D) {
	var retval *Zval
	if EXPECTED(Z_TYPE_P(container) == IS_ARRAY) {
	try_array:
		SEPARATE_ARRAY(container)
	fetch_from_array:
		if dim == nil {
			retval = ZendHashNextIndexInsert(Z_ARRVAL_P(container), &(ExecutorGlobals.GetUninitializedZval()))
			if UNEXPECTED(retval == nil) {
				ZendCannotAddElement()
				ZVAL_ERROR(result)
				return
			}
		} else {
			retval = ZendFetchDimensionAddressInner(Z_ARRVAL_P(container), dim, dim_type, type_, EXECUTE_DATA_C)
			if UNEXPECTED(retval == nil) {
				ZVAL_ERROR(result)
				return
			}
		}
		ZVAL_INDIRECT(result, retval)
		return
	} else if EXPECTED(Z_TYPE_P(container) == IS_REFERENCE) {
		var ref *ZendReference = Z_REF_P(container)
		container = Z_REFVAL_P(container)
		if EXPECTED(Z_TYPE_P(container) == IS_ARRAY) {
			goto try_array
		} else if EXPECTED(Z_TYPE_P(container) <= IS_FALSE) {
			if type_ != BP_VAR_UNSET {
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					if UNEXPECTED(ZendVerifyRefArrayAssignable(ref) == 0) {
						ZVAL_ERROR(result)
						return
					}
				}
				ArrayInit(container)
				goto fetch_from_array
			} else {
				goto return_null
			}
		}
	}
	if UNEXPECTED(Z_TYPE_P(container) == IS_STRING) {
		if dim == nil {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, type_, EXECUTE_DATA_C)
			ZendWrongStringOffset(EXECUTE_DATA_C)
		}
		ZVAL_ERROR(result)
	} else if EXPECTED(Z_TYPE_P(container) == IS_OBJECT) {
		if ZEND_CONST_COND(dim_type == IS_CV, dim != nil) && UNEXPECTED(Z_TYPE_P(dim) == IS_UNDEF) {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && Z_EXTRA_P(dim) == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		if UNEXPECTED(retval == &(ExecutorGlobals.GetUninitializedZval())) {
			var ce *ZendClassEntry = Z_OBJCE_P(container)
			ZVAL_NULL(result)
			ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ZSTR_VAL(ce.GetName()))
		} else if EXPECTED(retval != nil && Z_TYPE_P(retval) != IS_UNDEF) {
			if !(Z_ISREF_P(retval)) {
				if result != retval {
					ZVAL_COPY(result, retval)
					retval = result
				}
				if Z_TYPE_P(retval) != IS_OBJECT {
					var ce *ZendClassEntry = Z_OBJCE_P(container)
					ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ZSTR_VAL(ce.GetName()))
				}
			} else if UNEXPECTED(Z_REFCOUNT_P(retval) == 1) {
				ZVAL_UNREF(retval)
			}
			if result != retval {
				ZVAL_INDIRECT(result, retval)
			}
		} else {
			ZVAL_ERROR(result)
		}
	} else {
		if EXPECTED(Z_TYPE_P(container) <= IS_FALSE) {
			if type_ != BP_VAR_W && UNEXPECTED(Z_TYPE_P(container) == IS_UNDEF) {
				ZVAL_UNDEFINED_OP1()
			}
			if type_ != BP_VAR_UNSET {
				ArrayInit(container)
				goto fetch_from_array
			} else {
			return_null:

				/* for read-mode only */

				if ZEND_CONST_COND(dim_type == IS_CV, dim != nil) && UNEXPECTED(Z_TYPE_P(dim) == IS_UNDEF) {
					ZVAL_UNDEFINED_OP2()
				}
				ZVAL_NULL(result)
			}
		} else if EXPECTED(Z_ISERROR_P(container)) {
			ZVAL_ERROR(result)
		} else {
			if type_ == BP_VAR_UNSET {
				ZendError(E_WARNING, "Cannot unset offset in a non-array variable")
				ZVAL_NULL(result)
			} else {
				ZendUseScalarAsArray()
				ZVAL_ERROR(result)
			}
		}
	}
}
func zend_fetch_dimension_address_W(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_RW(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_RW, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_UNSET(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_UNSET, EXECUTE_DATA_C)
}
func ZendFetchDimensionAddressRead(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, is_list int, slow int, _ EXECUTE_DATA_D) {
	var retval *Zval
	if slow == 0 {
		if EXPECTED(Z_TYPE_P(container) == IS_ARRAY) {
		try_array:
			retval = ZendFetchDimensionAddressInner(Z_ARRVAL_P(container), dim, dim_type, type_, EXECUTE_DATA_C)
			ZVAL_COPY_DEREF(result, retval)
			return
		} else if EXPECTED(Z_TYPE_P(container) == IS_REFERENCE) {
			container = Z_REFVAL_P(container)
			if EXPECTED(Z_TYPE_P(container) == IS_ARRAY) {
				goto try_array
			}
		}
	}
	if is_list == 0 && EXPECTED(Z_TYPE_P(container) == IS_STRING) {
		var offset ZendLong
	try_string_offset:
		if UNEXPECTED(Z_TYPE_P(dim) != IS_LONG) {
			switch Z_TYPE_P(dim) {
			case IS_STRING:
				if IS_LONG == IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), nil, nil, -1) {
					break
				}
				if type_ == BP_VAR_IS {
					ZVAL_NULL(result)
					return
				}
				ZendError(E_WARNING, "Illegal string offset '%s'", Z_STRVAL_P(dim))
				break
			case IS_UNDEF:
				ZVAL_UNDEFINED_OP2()
			case IS_DOUBLE:

			case IS_NULL:

			case IS_FALSE:

			case IS_TRUE:
				if type_ != BP_VAR_IS {
					ZendError(E_NOTICE, "String offset cast occurred")
				}
				break
			case IS_REFERENCE:
				dim = Z_REFVAL_P(dim)
				goto try_string_offset
			default:
				ZendIllegalOffset()
				break
			}
			offset = ZvalGetLongFunc(dim)
		} else {
			offset = Z_LVAL_P(dim)
		}
		if UNEXPECTED(Z_STRLEN_P(container) < b.CondF(offset < 0, func() int { return -int(offset) }, func() int { return int(offset + 1) })) {
			if type_ != BP_VAR_IS {
				ZendError(E_NOTICE, "Uninitialized string offset: "+ZEND_LONG_FMT, offset)
				ZVAL_EMPTY_STRING(result)
			} else {
				ZVAL_NULL(result)
			}
		} else {
			var c ZendUchar
			var real_offset ZendLong
			if UNEXPECTED(offset < 0) {
				real_offset = ZendLong(Z_STRLEN_P(container) + offset)
			} else {
				real_offset = offset
			}
			c = ZendUchar(Z_STRVAL_P(container)[real_offset])
			ZVAL_INTERNED_STR(result, ZSTR_CHAR(c))
		}
	} else if EXPECTED(Z_TYPE_P(container) == IS_OBJECT) {
		if ZEND_CONST_COND(dim_type == IS_CV, 1) && UNEXPECTED(Z_TYPE_P(dim) == IS_UNDEF) {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && Z_EXTRA_P(dim) == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		ZEND_ASSERT(result != nil)
		if retval != nil {
			if result != retval {
				ZVAL_COPY_DEREF(result, retval)
			} else if UNEXPECTED(Z_ISREF_P(retval)) {
				ZendUnwrapReference(result)
			}
		} else {
			ZVAL_NULL(result)
		}
	} else {
		if type_ != BP_VAR_IS && UNEXPECTED(Z_TYPE_P(container) == IS_UNDEF) {
			container = ZVAL_UNDEFINED_OP1()
		}
		if ZEND_CONST_COND(dim_type == IS_CV, 1) && UNEXPECTED(Z_TYPE_P(dim) == IS_UNDEF) {
			ZVAL_UNDEFINED_OP2()
		}
		if is_list == 0 && type_ != BP_VAR_IS {
			ZendError(E_NOTICE, "Trying to access array offset on value of type %s", ZendZvalTypeName(container))
		}
		ZVAL_NULL(result)
	}
}
func zend_fetch_dimension_address_read_R(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 0, 0, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_read_R_slow(container *Zval, dim *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, IS_CV, BP_VAR_R, 0, 1, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_read_IS(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_IS, 0, 0, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_LIST_r(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 1, 0, EXECUTE_DATA_C)
}
func ZendFetchDimensionConst(result *Zval, container *Zval, dim *Zval, type_ int) {
	ZendFetchDimensionAddressRead(result, container, dim, IS_TMP_VAR, type_, 0, 0, nil)
}
func ZendFindArrayDimSlow(ht *HashTable, offset *Zval, _ EXECUTE_DATA_D) *Zval {
	var hval ZendUlong
	if Z_TYPE_P(offset) == IS_DOUBLE {
		hval = ZendDvalToLval(Z_DVAL_P(offset))
	num_idx:
		return ZendHashIndexFind(ht, hval)
	} else if Z_TYPE_P(offset) == IS_NULL {
	str_idx:
		return ZendHashFindExInd(ht, ZSTR_EMPTY_ALLOC(), 1)
	} else if Z_TYPE_P(offset) == IS_FALSE {
		hval = 0
		goto num_idx
	} else if Z_TYPE_P(offset) == IS_TRUE {
		hval = 1
		goto num_idx
	} else if Z_TYPE_P(offset) == IS_RESOURCE {
		hval = Z_RES_HANDLE_P(offset)
		goto num_idx
	} else if Z_TYPE_P(offset) == IS_UNDEF {
		ZVAL_UNDEFINED_OP2()
		goto str_idx
	} else {
		ZendError(E_WARNING, "Illegal offset type in isset or empty")
		return nil
	}
}
func ZendIssetDimSlow(container *Zval, offset *Zval, _ EXECUTE_DATA_D) int {
	if UNEXPECTED(Z_TYPE_P(offset) == IS_UNDEF) {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if EXPECTED(Z_TYPE_P(container) == IS_OBJECT) {
		return Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 0)
	} else if EXPECTED(Z_TYPE_P(container) == IS_STRING) {
		var lval ZendLong
		if EXPECTED(Z_TYPE_P(offset) == IS_LONG) {
			lval = Z_LVAL_P(offset)
		str_offset:
			if UNEXPECTED(lval < 0) {
				lval += ZendLong(Z_STRLEN_P(container))
			}
			if EXPECTED(lval >= 0) && int(lval < Z_STRLEN_P(container)) != 0 {
				return 1
			} else {
				return 0
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			ZVAL_DEREF(offset)

			/*}*/

			if Z_TYPE_P(offset) < IS_STRING || Z_TYPE_P(offset) == IS_STRING && IS_LONG == IsNumericString(Z_STRVAL_P(offset), Z_STRLEN_P(offset), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 0
		}
	} else {
		return 0
	}
}
func ZendIsemptyDimSlow(container *Zval, offset *Zval, _ EXECUTE_DATA_D) int {
	if UNEXPECTED(Z_TYPE_P(offset) == IS_UNDEF) {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if EXPECTED(Z_TYPE_P(container) == IS_OBJECT) {
		return !(Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 1))
	} else if EXPECTED(Z_TYPE_P(container) == IS_STRING) {
		var lval ZendLong
		if EXPECTED(Z_TYPE_P(offset) == IS_LONG) {
			lval = Z_LVAL_P(offset)
		str_offset:
			if UNEXPECTED(lval < 0) {
				lval += ZendLong(Z_STRLEN_P(container))
			}
			if EXPECTED(lval >= 0) && int(lval < Z_STRLEN_P(container)) != 0 {
				return Z_STRVAL_P(container)[lval] == '0'
			} else {
				return 1
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			ZVAL_DEREF(offset)

			/*}*/

			if Z_TYPE_P(offset) < IS_STRING || Z_TYPE_P(offset) == IS_STRING && IS_LONG == IsNumericString(Z_STRVAL_P(offset), Z_STRLEN_P(offset), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 1
		}
	} else {
		return 1
	}
}
func ZendArrayKeyExistsFast(ht *HashTable, key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) uint32 {
	var str *ZendString
	var hval ZendUlong
try_again:
	if EXPECTED(Z_TYPE_P(key) == IS_STRING) {
		str = Z_STR_P(key)
		if ZEND_HANDLE_NUMERIC(str, hval) != 0 {
			goto num_key
		}
	str_key:
		if ZendHashFindInd(ht, str) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if EXPECTED(Z_TYPE_P(key) == IS_LONG) {
		hval = Z_LVAL_P(key)
	num_key:
		if ZendHashIndexFind(ht, hval) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if EXPECTED(Z_ISREF_P(key)) {
		key = Z_REFVAL_P(key)
		goto try_again
	} else if Z_TYPE_P(key) <= IS_NULL {
		if UNEXPECTED(Z_TYPE_P(key) == IS_UNDEF) {
			ZVAL_UNDEFINED_OP1()
		}
		str = ZSTR_EMPTY_ALLOC()
		goto str_key
	} else {
		ZendError(E_WARNING, "array_key_exists(): The first argument should be either a string or an integer")
		return IS_FALSE
	}
}
func ZendArrayKeyExistsSlow(subject *Zval, key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) uint32 {
	if EXPECTED(Z_TYPE_P(subject) == IS_OBJECT) {
		ZendError(E_DEPRECATED, "array_key_exists(): "+"Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
		var ht *HashTable = ZendGetPropertiesFor(subject, ZEND_PROP_PURPOSE_ARRAY_CAST)
		var result uint32 = ZendArrayKeyExistsFast(ht, key, OPLINE_C, EXECUTE_DATA_C)
		ZendReleaseProperties(ht)
		return result
	} else {
		if UNEXPECTED(Z_TYPE_P(key) == IS_UNDEF) {
			ZVAL_UNDEFINED_OP1()
		}
		if UNEXPECTED(Z_TYPE_INFO_P(subject) == IS_UNDEF) {
			ZVAL_UNDEFINED_OP2()
		}
		ZendInternalTypeError(EX_USES_STRICT_TYPES(), "array_key_exists() expects parameter 2 to be array, %s given", ZendGetTypeByConst(Z_TYPE_P(subject)))
		return IS_NULL
	}
}
func PromotesToArray(val *Zval) ZendBool {
	return Z_TYPE_P(val) <= IS_FALSE || Z_ISREF_P(val) && Z_TYPE_P(Z_REFVAL_P(val)) <= IS_FALSE
}
func PromotesToObject(val *Zval) ZendBool {
	ZVAL_DEREF(val)
	return Z_TYPE_P(val) <= IS_FALSE || Z_TYPE_P(val) == IS_STRING && Z_STRLEN_P(val) == 0
}
func CheckTypeArrayAssignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	return ZEND_TYPE_IS_CODE(type_) && (ZEND_TYPE_CODE(type_) == IS_ARRAY || ZEND_TYPE_CODE(type_) == IS_ITERABLE)
}
func check_type_stdClass_assignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	if ZEND_TYPE_IS_CLASS(type_) {
		if ZEND_TYPE_IS_CE(type_) {
			return ZEND_TYPE_CE(type_) == ZendStandardClassDef
		} else {
			return ZendStringEqualsLiteralCi(ZEND_TYPE_NAME(type_), "stdclass")
		}
	} else {
		return ZEND_TYPE_CODE(type_) == IS_OBJECT
	}
}

/* Checks whether an array can be assigned to the reference. Returns conflicting property if
 * assignment is not possible, NULL otherwise. */

func ZendVerifyRefArrayAssignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *ZendPropertyInfoSourceList = &ZEND_REF_TYPE_SOURCES(ref)
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = &_source_list.ptr
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			if CheckTypeArrayAssignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "array")
				return 0
			}
		}
	}
	return 1
}

/* Checks whether an stdClass can be assigned to the reference. Returns conflicting property if
 * assignment is not possible, NULL otherwise. */

func zend_verify_ref_stdClass_assignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *ZendPropertyInfoSourceList = &ZEND_REF_TYPE_SOURCES(ref)
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = &_source_list.ptr
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			if check_type_stdClass_assignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "stdClass")
				return 0
			}
		}
	}
	return 1
}
func ZendObjectFetchPropertyTypeInfo(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	if EXPECTED(!(ZEND_CLASS_HAS_TYPE_HINTS(obj.GetCe()))) {
		return nil
	}

	/* Not a declared property */

	if UNEXPECTED(slot < obj.GetPropertiesTable() || slot >= obj.GetPropertiesTable()+obj.GetCe().GetDefaultPropertiesCount()) {
		return nil
	}
	return ZendGetTypedPropertyInfoForSlot(obj, slot)
}
func ZendHandleFetchObjFlags(result *Zval, ptr *Zval, obj *ZendObject, prop_info *ZendPropertyInfo, flags uint32) ZendBool {
	switch flags {
	case ZEND_FETCH_DIM_WRITE:
		if PromotesToArray(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if CheckTypeArrayAssignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "array")
				if result != nil {
					ZVAL_ERROR(result)
				}
				return 0
			}
		}
		break
	case ZEND_FETCH_OBJ_WRITE:
		if PromotesToObject(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if check_type_stdClass_assignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "stdClass")
				if result != nil {
					ZVAL_ERROR(result)
				}
				return 0
			}
		}
		break
	case ZEND_FETCH_REF:
		if Z_TYPE_P(ptr) != IS_REFERENCE {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if Z_TYPE_P(ptr) == IS_UNDEF {
				if !(ZEND_TYPE_ALLOW_NULL(prop_info.GetType())) {
					ZendThrowAccessUninitPropByRefError(prop_info)
					if result != nil {
						ZVAL_ERROR(result)
					}
					return 0
				}
				ZVAL_NULL(ptr)
			}
			ZVAL_NEW_REF(ptr, ptr)
			ZEND_REF_ADD_TYPE_SOURCE(Z_REF_P(ptr), prop_info)
		}
		break
	default:
		break
	}
	return 1
}
func ZendFetchPropertyAddress(result *Zval, container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, cache_slot *any, type_ int, flags uint32, init_undef ZendBool, opline *ZendOp, _ EXECUTE_DATA_D) {
	var ptr *Zval
	if container_op_type != IS_UNUSED && UNEXPECTED(Z_TYPE_P(container) != IS_OBJECT) {
		for {
			if Z_ISREF_P(container) && Z_TYPE_P(Z_REFVAL_P(container)) == IS_OBJECT {
				container = Z_REFVAL_P(container)
				break
			}
			if container_op_type == IS_CV && type_ != BP_VAR_W && UNEXPECTED(Z_TYPE_P(container) == IS_UNDEF) {
				ZVAL_UNDEFINED_OP1()
			}

			/* this should modify object only if it's empty */

			if type_ == BP_VAR_UNSET {
				ZVAL_NULL(result)
				return
			}
			container = MakeRealObject(container, prop_ptr, OPLINE_C, EXECUTE_DATA_C)
			if UNEXPECTED(container == nil) {
				ZVAL_ERROR(result)
				return
			}
			break
		}
	}
	if prop_op_type == IS_CONST && EXPECTED(Z_OBJCE_P(container) == CACHED_PTR_EX(cache_slot)) {
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *ZendObject = Z_OBJ_P(container)
		if EXPECTED(IS_VALID_PROPERTY_OFFSET(prop_offset)) {
			ptr = OBJ_PROP(zobj, prop_offset)
			if EXPECTED(Z_TYPE_P(ptr) != IS_UNDEF) {
				ZVAL_INDIRECT(result, ptr)
				if flags != 0 {
					var prop_info *ZendPropertyInfo = CACHED_PTR_EX(cache_slot + 2)
					if prop_info != nil {
						ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags)
					}
				}
				return
			}
		} else if EXPECTED(zobj.GetProperties() != nil) {
			if UNEXPECTED(GC_REFCOUNT(zobj.GetProperties()) > 1) {
				if EXPECTED((GC_FLAGS(zobj.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0) {
					GC_DELREF(zobj.GetProperties())
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			ptr = ZendHashFindEx(zobj.GetProperties(), Z_STR_P(prop_ptr), 1)
			if EXPECTED(ptr != nil) {
				ZVAL_INDIRECT(result, ptr)
				return
			}
		}
	}
	ptr = Z_OBJ_HT_P(container).GetGetPropertyPtrPtr()(container, prop_ptr, type_, cache_slot)
	if nil == ptr {
		ptr = Z_OBJ_HT_P(container).GetReadProperty()(container, prop_ptr, type_, cache_slot, result)
		if ptr == result {
			if UNEXPECTED(Z_ISREF_P(ptr) && Z_REFCOUNT_P(ptr) == 1) {
				ZVAL_UNREF(ptr)
			}
			return
		}
		if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			ZVAL_ERROR(result)
			return
		}
	} else if UNEXPECTED(Z_ISERROR_P(ptr)) {
		ZVAL_ERROR(result)
		return
	}
	ZVAL_INDIRECT(result, ptr)
	if flags != 0 {
		var prop_info *ZendPropertyInfo
		if prop_op_type == IS_CONST {
			prop_info = CACHED_PTR_EX(cache_slot + 2)
			if prop_info != nil {
				if UNEXPECTED(ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags) == 0) {
					return
				}
			}
		} else {
			if UNEXPECTED(ZendHandleFetchObjFlags(result, ptr, Z_OBJ_P(container), nil, flags) == 0) {
				return
			}
		}
	}
	if init_undef != 0 && UNEXPECTED(Z_TYPE_P(ptr) == IS_UNDEF) {
		ZVAL_NULL(ptr)
	}
}
func ZendAssignToPropertyReference(container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var variable Zval
	var variable_ptr *Zval = &variable
	var cache_addr *any = b.CondF1(prop_op_type == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION) }, nil)
	ZendFetchPropertyAddress(variable_ptr, container, container_op_type, prop_ptr, prop_op_type, cache_addr, BP_VAR_W, 0, 0, OPLINE_C, EXECUTE_DATA_C)
	if Z_TYPE_P(variable_ptr) == IS_INDIRECT {
		variable_ptr = Z_INDIRECT_P(variable_ptr)
	}
	if UNEXPECTED(Z_ISERROR_P(variable_ptr)) {
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if UNEXPECTED(Z_TYPE(variable) != IS_INDIRECT) {
		ZendThrowError(nil, "Cannot assign by reference to overloaded object")
		ZvalPtrDtor(&variable)
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if UNEXPECTED(Z_ISERROR_P(value_ptr)) {
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && UNEXPECTED(!(Z_ISREF_P(value_ptr))) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, OPLINE_C, EXECUTE_DATA_C)
	} else {
		var prop_info *ZendPropertyInfo = nil
		if prop_op_type == IS_CONST {
			prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_addr + 2))
		} else {
			ZVAL_DEREF(container)
			prop_info = ZendObjectFetchPropertyTypeInfo(Z_OBJ_P(container), variable_ptr)
		}
		if UNEXPECTED(prop_info != nil) {
			variable_ptr = ZendAssignToTypedPropertyReference(prop_info, variable_ptr, value_ptr, EXECUTE_DATA_C)
		} else {
			ZendAssignToVariableReference(variable_ptr, value_ptr)
		}
	}
	if UNEXPECTED(RETURN_VALUE_USED(opline)) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), variable_ptr)
	}
}
func ZendAssignToPropertyReferenceThisConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceThisVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendFetchStaticPropertyAddressEx(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var free_op1 ZendFreeOp
	var name *ZendString
	var tmp_name *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var op1_type ZendUchar = opline.GetOp1Type()
	var op2_type ZendUchar = opline.GetOp2Type()
	if EXPECTED(op2_type == IS_CONST) {
		var class_name *Zval = RT_CONSTANT(opline, opline.GetOp2())
		ZEND_ASSERT(op1_type != IS_CONST || CACHED_PTR(cache_slot) == nil)
		if EXPECTED(b.Assign(&ce, CACHED_PTR(cache_slot)) == nil) {
			ce = ZendFetchClassByName(Z_STR_P(class_name), Z_STR_P(class_name+1), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if UNEXPECTED(ce == nil) {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return FAILURE
			}
			if UNEXPECTED(op1_type != IS_CONST) {
				CACHE_PTR(cache_slot, ce)
			}
		}
	} else {
		if EXPECTED(op2_type == IS_UNUSED) {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if UNEXPECTED(ce == nil) {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return FAILURE
			}
		} else {
			ce = Z_CE_P(EX_VAR(opline.GetOp2().GetVar()))
		}
		if EXPECTED(op1_type == IS_CONST) && EXPECTED(CACHED_PTR(cache_slot) == ce) {
			*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
			*prop_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
			return SUCCESS
		}
	}
	if EXPECTED(op1_type == IS_CONST) {
		name = Z_STR_P(RT_CONSTANT(opline, opline.GetOp1()))
	} else {
		var varname *Zval = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		if EXPECTED(Z_TYPE_P(varname) == IS_STRING) {
			name = Z_STR_P(varname)
			tmp_name = nil
		} else {
			if op1_type == IS_CV && UNEXPECTED(Z_TYPE_P(varname) == IS_UNDEF) {
				ZvalUndefinedCv(opline.GetOp1().GetVar(), EXECUTE_DATA_C)
			}
			name = ZvalGetTmpString(varname, &tmp_name)
		}
	}
	*retval = ZendStdGetStaticPropertyWithInfo(ce, name, fetch_type, &property_info)
	if UNEXPECTED(op1_type != IS_CONST) {
		ZendTmpStringRelease(tmp_name)
		if op1_type != IS_CV {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	if UNEXPECTED((*retval) == nil) {
		return FAILURE
	}
	*prop_info = property_info
	if EXPECTED(op1_type == IS_CONST) {
		CACHE_POLYMORPHIC_PTR(cache_slot, ce, *retval)
		CACHE_PTR(cache_slot+b.SizeOf("void *")*2, property_info)
	}
	return SUCCESS
}
func ZendFetchStaticPropertyAddress(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, flags int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var success int
	var property_info *ZendPropertyInfo
	if opline.GetOp1Type() == IS_CONST && (opline.GetOp2Type() == IS_CONST || opline.GetOp2Type() == IS_UNUSED && (opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_SELF || opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_PARENT)) && EXPECTED(CACHED_PTR(cache_slot) != nil) {
		*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
		property_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
		if (fetch_type == BP_VAR_R || fetch_type == BP_VAR_RW) && UNEXPECTED(Z_TYPE_P(*retval) == IS_UNDEF) && UNEXPECTED(property_info.GetType() != 0) {
			ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", ZSTR_VAL(property_info.GetCe().GetName()), ZendGetUnmangledPropertyName(property_info.GetName()))
			return FAILURE
		}
	} else {
		success = ZendFetchStaticPropertyAddressEx(retval, &property_info, cache_slot, fetch_type, OPLINE_C, EXECUTE_DATA_C)
		if UNEXPECTED(success != SUCCESS) {
			return FAILURE
		}
	}
	if flags != 0 && property_info.GetType() != 0 {
		ZendHandleFetchObjFlags(nil, *retval, nil, property_info, flags)
	}
	if prop_info != nil {
		*prop_info = property_info
	}
	return SUCCESS
}
func ZendThrowRefTypeErrorType(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Reference with value of type %s held by property %s::$%s of type %s%s is not compatible with property %s::$%s of type %s%s", b.CondF(Z_TYPE_P(zv) == IS_OBJECT, func() []byte { return ZSTR_VAL(Z_OBJCE_P(zv).GetName()) }, func() *byte { return ZendGetTypeByConst(Z_TYPE_P(zv)) }), ZSTR_VAL(prop1.GetCe().GetName()), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, ZSTR_VAL(prop2.GetCe().GetName()), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func ZendThrowRefTypeErrorZval(prop *ZendPropertyInfo, zv *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s", b.CondF(Z_TYPE_P(zv) == IS_OBJECT, func() []byte { return ZSTR_VAL(Z_OBJCE_P(zv).GetName()) }, func() *byte { return ZendGetTypeByConst(Z_TYPE_P(zv)) }), ZSTR_VAL(prop.GetCe().GetName()), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowConflictingCoercionError(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s and property %s::$%s of type %s%s, as this would result in an inconsistent type conversion", b.CondF(Z_TYPE_P(zv) == IS_OBJECT, func() []byte { return ZSTR_VAL(Z_OBJCE_P(zv).GetName()) }, func() *byte { return ZendGetTypeByConst(Z_TYPE_P(zv)) }), ZSTR_VAL(prop1.GetCe().GetName()), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, ZSTR_VAL(prop2.GetCe().GetName()), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}

/* 1: valid, 0: invalid, -1: may be valid after type coercion */

func IZendVerifyTypeAssignableZval(type_ptr *ZendType, self_ce *ZendClassEntry, zv *Zval, strict ZendBool) int {
	var type_ ZendType = *type_ptr
	var type_code ZendUchar
	var zv_type ZendUchar = Z_TYPE_P(zv)
	if ZEND_TYPE_ALLOW_NULL(type_) && zv_type == IS_NULL {
		return 1
	}
	if ZEND_TYPE_IS_CLASS(type_) {
		if !(ZEND_TYPE_IS_CE(type_)) {
			if ZendResolveClassType(type_ptr, self_ce) == 0 {
				return 0
			}
			type_ = *type_ptr
		}
		return zv_type == IS_OBJECT && InstanceofFunction(Z_OBJCE_P(zv), ZEND_TYPE_CE(type_)) != 0
	}
	type_code = ZEND_TYPE_CODE(type_)
	if type_code == zv_type || type_code == _IS_BOOL && (zv_type == IS_FALSE || zv_type == IS_TRUE) {
		return 1
	}
	if type_code == IS_ITERABLE {
		return ZendIsIterable(zv)
	}

	/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	if strict != 0 {
		if type_code == IS_DOUBLE && zv_type == IS_LONG {
			return -1
		}
		return 0
	}

	/* No weak conversions for arrays and objects */

	if type_code == IS_ARRAY || type_code == IS_OBJECT {
		return 0
	}

	/* NULL may be accepted only by nullable hints (this is already checked) */

	if zv_type == IS_NULL {
		return 0
	}

	/* Coercion may be necessary, check separately */

	return -1

	/* Coercion may be necessary, check separately */
}
func ZendVerifyRefAssignableZval(ref *ZendReference, zv *Zval, strict ZendBool) ZendBool {
	var prop *ZendPropertyInfo

	/* The value must satisfy each property type, and coerce to the same value for each property
	 * type. Right now, the latter rule means that *if* coercion is necessary, then all types
	 * must be the same (modulo nullability). To handle this, remember the first type we see and
	 * compare against it when coercion becomes necessary. */

	var seen_prop *ZendPropertyInfo = nil
	var seen_type ZendUchar
	var needs_coercion ZendBool = 0
	ZEND_ASSERT(Z_TYPE_P(zv) != IS_REFERENCE)
	var _source_list *ZendPropertyInfoSourceList = &ZEND_REF_TYPE_SOURCES(ref)
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = &_source_list.ptr
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			var result int = IZendVerifyTypeAssignableZval(&prop.type_, prop.GetCe(), zv, strict)
			if result == 0 {
				ZendThrowRefTypeErrorZval(prop, zv)
				return 0
			}
			if result < 0 {
				needs_coercion = 1
			}
			if seen_prop == nil {
				seen_prop = prop
				if ZEND_TYPE_IS_CLASS(prop.GetType()) {
					seen_type = IS_OBJECT
				} else {
					seen_type = ZEND_TYPE_CODE(prop.GetType())
				}
			} else if needs_coercion != 0 && seen_type != ZEND_TYPE_CODE(prop.GetType()) {
				ZendThrowConflictingCoercionError(seen_prop, prop, zv)
				return 0
			}
		}
	}
	if UNEXPECTED(needs_coercion != 0 && ZendVerifyWeakScalarTypeHint(seen_type, zv) == 0) {
		ZendThrowRefTypeErrorZval(seen_prop, zv)
		return 0
	}
	return 1
}
func IZvalPtrDtorNoref(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = Z_COUNTED_P(zval_ptr)
		ZEND_ASSERT(Z_TYPE_P(zval_ptr) != IS_REFERENCE)
		if GC_DELREF(ref) == 0 {
			RcDtorFunc(ref)
		} else if UNEXPECTED(GC_MAY_LEAK(ref)) {
			GcPossibleRoot(ref)
		}
	}
}
func ZendAssignToTypedRef(variable_ptr *Zval, orig_value *Zval, value_type ZendUchar, strict ZendBool, ref *ZendRefcounted) *Zval {
	var ret ZendBool
	var value Zval
	ZVAL_COPY(&value, orig_value)
	ret = ZendVerifyRefAssignableZval(Z_REF_P(variable_ptr), &value, strict)
	variable_ptr = Z_REFVAL_P(variable_ptr)
	if EXPECTED(ret != 0) {
		IZvalPtrDtorNoref(variable_ptr)
		ZVAL_COPY_VALUE(variable_ptr, &value)
	} else {
		ZvalPtrDtorNogc(&value)
	}
	if (value_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		if UNEXPECTED(ref != nil) {
			if UNEXPECTED(GC_DELREF(ref) == 0) {
				ZvalPtrDtor(orig_value)
				EfreeSize(ref, b.SizeOf("zend_reference"))
			}
		} else {
			IZvalPtrDtorNoref(orig_value)
		}
	}
	return variable_ptr
}
func ZendVerifyPropAssignableByRef(prop_info *ZendPropertyInfo, orig_val *Zval, strict ZendBool) ZendBool {
	var val *Zval = orig_val
	if Z_ISREF_P(val) && ZEND_REF_HAS_TYPE_SOURCES(Z_REF_P(val)) {
		var result int
		val = Z_REFVAL_P(val)
		result = IZendVerifyTypeAssignableZval(&prop_info.type_, prop_info.GetCe(), val, strict)
		if result > 0 {
			return 1
		}
		if result < 0 {
			var ref_prop *ZendPropertyInfo = ZEND_REF_FIRST_SOURCE(Z_REF_P(orig_val))
			if ZEND_TYPE_CODE(prop_info.GetType()) != ZEND_TYPE_CODE(ref_prop.GetType()) {

				/* Invalid due to conflicting coercion */

				ZendThrowRefTypeErrorType(ref_prop, prop_info, val)
				return 0
			}
			if ZendVerifyWeakScalarTypeHint(ZEND_TYPE_CODE(prop_info.GetType()), val) != 0 {
				return 1
			}
		}
	} else {
		ZVAL_DEREF(val)
		if IZendCheckPropertyType(prop_info, val, strict) != 0 {
			return 1
		}
	}
	ZendVerifyPropertyTypeError(prop_info, val)
	return 0
}
func ZendRefAddTypeSource(source_list *ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *ZendPropertyInfoList
	if source_list.GetPtr() == nil {
		source_list.SetPtr(prop)
		return
	}
	list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		list = Emalloc(b.SizeOf("zend_property_info_list") + (4-1)*b.SizeOf("zend_property_info *"))
		list.GetPtr()[0] = source_list.GetPtr()
		list.SetNumAllocated(4)
		list.SetNum(1)
	} else if list.GetNumAllocated() == list.GetNum() {
		list.SetNumAllocated(list.GetNum() * 2)
		list = Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))
	}
	list.GetPtr()[b.PostInc(&(list.GetNum()))] = prop
	source_list.SetList(ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list))
}
func ZendRefDelTypeSource(source_list *ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *ZendPropertyInfoList = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	var ptr **ZendPropertyInfo
	var end ***ZendPropertyInfo
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		ZEND_ASSERT(source_list.GetPtr() == prop)
		source_list.SetPtr(nil)
		return
	}
	if list.GetNum() == 1 {
		ZEND_ASSERT((*list).ptr == prop)
		Efree(list)
		source_list.SetPtr(nil)
		return
	}

	/* Checking against end here to get a more graceful failure mode if we missed adding a type
	 * source at some point. */

	ptr = list.GetPtr()
	end = ptr + list.GetNum()
	for ptr < end && (*ptr) != prop {
		ptr++
	}
	ZEND_ASSERT((*ptr) == prop)

	/* Copy the last list element into the deleted slot. */

	*ptr = list.GetPtr()[b.PreDec(&(list.GetNum()))]
	if list.GetNum() >= 4 && list.GetNum()*4 == list.GetNumAllocated() {
		list.SetNumAllocated(list.GetNum() * 2)
		source_list.SetList(ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))))
	}
}
func ZendFetchThisVar(type_ int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	switch type_ {
	case BP_VAR_R:
		if EXPECTED(Z_TYPE(EX(This)) == IS_OBJECT) {
			ZVAL_OBJ(result, Z_OBJ(EX(This)))
			Z_ADDREF_P(result)
		} else {
			ZVAL_NULL(result)
			ZendError(E_NOTICE, "Undefined variable: this")
		}
		break
	case BP_VAR_IS:
		if EXPECTED(Z_TYPE(EX(This)) == IS_OBJECT) {
			ZVAL_OBJ(result, Z_OBJ(EX(This)))
			Z_ADDREF_P(result)
		} else {
			ZVAL_NULL(result)
		}
		break
	case BP_VAR_RW:

	case BP_VAR_W:
		ZVAL_UNDEF(result)
		ZendThrowError(nil, "Cannot re-assign $this")
		break
	case BP_VAR_UNSET:
		ZVAL_UNDEF(result)
		ZendThrowError(nil, "Cannot unset $this")
		break
	default:
		break
	}
}
func ZendWrongCloneCall(clone *ZendFunction, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s %s::__clone() from context '%s'", ZendVisibilityString(clone.GetFnFlags()), ZSTR_VAL(clone.GetScope().GetName()), b.CondF1(scope != nil, func() []byte { return ZSTR_VAL(scope.GetName()) }, ""))
}

// #define CHECK_SYMBOL_TABLES()

func ExecuteInternal(execute_data *ZendExecuteData, return_value *Zval) {
	execute_data.GetFunc().GetInternalFunction().GetHandler()(execute_data, return_value)
}
func ZendCleanAndCacheSymbolTable(symbol_table *ZendArray) {
	/* Clean before putting into the cache, since clean could call dtors,
	 * which could use the cached hash. Also do this before the check for
	 * available cache slots, as those may be used by a dtor as well. */

	ZendSymtableClean(symbol_table)
	if ExecutorGlobals.GetSymtableCachePtr() >= ExecutorGlobals.GetSymtableCacheLimit() {
		ZendArrayDestroy(symbol_table)
	} else {
		*(b.PostInc(&(ExecutorGlobals.GetSymtableCachePtr()))) = symbol_table
	}
}

/* }}} */

func IFreeCompiledVariables(execute_data *ZendExecuteData) {
	var cv *Zval = EX_VAR_NUM(0)
	var count int = EX(func_).op_array.last_var
	for EXPECTED(count != 0) {
		if Z_REFCOUNTED_P(cv) {
			var r *ZendRefcounted = Z_COUNTED_P(cv)
			if GC_DELREF(r) == 0 {
				ZVAL_NULL(cv)
				RcDtorFunc(r)
			} else {
				GcCheckPossibleRoot(r)
			}
		}
		cv++
		count--
	}
}

/* }}} */

func ZendFreeCompiledVariables(execute_data *ZendExecuteData) { IFreeCompiledVariables(execute_data) }

/* }}} */

func ZEND_VM_INTERRUPT_CHECK() {
	if UNEXPECTED(ExecutorGlobals.GetVmInterrupt() != 0) {
		ZEND_VM_INTERRUPT()
	}
}
func ZEND_VM_LOOP_INTERRUPT_CHECK() {
	if UNEXPECTED(ExecutorGlobals.GetVmInterrupt() != 0) {
		ZEND_VM_LOOP_INTERRUPT()
	}
}

/*
 * Stack Frame Layout (the whole stack frame is allocated at once)
 * ==================
 *
 *                             +========================================+
 * EG(current_execute_data) -> | zend_execute_data                      |
 *                             +----------------------------------------+
 *     EX_VAR_NUM(0) --------> | VAR[0] = ARG[1]                        |
 *                             | ...                                    |
 *                             | VAR[op_array->num_args-1] = ARG[N]     |
 *                             | ...                                    |
 *                             | VAR[op_array->last_var-1]              |
 *                             | VAR[op_array->last_var] = TMP[0]       |
 *                             | ...                                    |
 *                             | VAR[op_array->last_var+op_array->T-1]  |
 *                             | ARG[N+1] (extra_args)                  |
 *                             | ...                                    |
 *                             +----------------------------------------+
 */

func ZendCopyExtraArgs(EXECUTE_DATA_D) {
	var op_array *ZendOpArray = &EX(func_).op_array
	var first_extra_arg uint32 = op_array.GetNumArgs()
	var num_args uint32 = EX_NUM_ARGS()
	var src *Zval
	var delta int
	var count uint32
	var type_flags uint32 = 0
	if EXPECTED((op_array.GetFnFlags() & ZEND_ACC_HAS_TYPE_HINTS) == 0) {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		EX(opline) += first_extra_arg

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* move extra args into separate array after all CV and TMP vars */

	src = EX_VAR_NUM(num_args - 1)
	delta = op_array.GetLastVar() + op_array.GetT() - first_extra_arg
	count = num_args - first_extra_arg
	if EXPECTED(delta != 0) {
		delta *= b.SizeOf("zval")
		for {
			type_flags |= Z_TYPE_INFO_P(src)
			ZVAL_COPY_VALUE((*Zval)((*byte)(src)+delta), src)
			ZVAL_UNDEF(src)
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
		if Z_TYPE_INFO_REFCOUNTED(type_flags) {
			ZEND_ADD_CALL_FLAG(execute_data, ZEND_CALL_FREE_EXTRA_ARGS)
		}
	} else {
		for {
			if Z_REFCOUNTED_P(src) {
				ZEND_ADD_CALL_FLAG(execute_data, ZEND_CALL_FREE_EXTRA_ARGS)
				break
			}
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendInitCvs(first uint32, last uint32, _ EXECUTE_DATA_D) {
	if EXPECTED(first < last) {
		var count uint32 = last - first
		var var_ *Zval = EX_VAR_NUM(first)
		for {
			ZVAL_UNDEF(var_)
			var_++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func IInitFuncExecuteData(op_array *ZendOpArray, return_value *Zval, may_be_trampoline ZendBool, _ EXECUTE_DATA_D) {
	var first_extra_arg uint32
	var num_args uint32
	ZEND_ASSERT(EX(func_) == (*ZendFunction)(op_array))
	EX(opline) = op_array.GetOpcodes()
	EX(call) = nil
	EX(return_value) = return_value

	/* Handle arguments */

	first_extra_arg = op_array.GetNumArgs()
	num_args = EX_NUM_ARGS()
	if UNEXPECTED(num_args > first_extra_arg) {
		if may_be_trampoline == 0 || EXPECTED((op_array.GetFnFlags()&ZEND_ACC_CALL_VIA_TRAMPOLINE) == 0) {
			ZendCopyExtraArgs(EXECUTE_DATA_C)
		}
	} else if EXPECTED((op_array.GetFnFlags() & ZEND_ACC_HAS_TYPE_HINTS) == 0) {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		EX(opline) += num_args

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* Initialize CV variables (skip arguments) */

	ZendInitCvs(num_args, op_array.GetLastVar(), EXECUTE_DATA_C)
	EX(run_time_cache) = RUN_TIME_CACHE(op_array)
	ExecutorGlobals.SetCurrentExecuteData(execute_data)
}

/* }}} */

func InitFuncRunTimeCacheI(op_array *ZendOpArray) {
	var run_time_cache *any
	ZEND_ASSERT(RUN_TIME_CACHE(op_array) == nil)
	run_time_cache = ZendArenaAlloc(&(CompilerGlobals.GetArena()), op_array.GetCacheSize())
	memset(run_time_cache, 0, op_array.GetCacheSize())
	ZEND_MAP_PTR_SET(op_array.run_time_cache, run_time_cache)
}

/* }}} */

func InitFuncRunTimeCache(op_array *ZendOpArray) { InitFuncRunTimeCacheI(op_array) }

/* }}} */

func ZendFetchFunction(name *ZendString) *ZendFunction {
	var zv *Zval = ZendHashFind(ExecutorGlobals.GetFunctionTable(), name)
	if EXPECTED(zv != nil) {
		var fbc *ZendFunction = Z_FUNC_P(zv)
		if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
			InitFuncRunTimeCacheI(&fbc.op_array)
		}
		return fbc
	}
	return nil
}
func ZendFetchFunctionStr(name string, len_ int) *ZendFunction {
	var zv *Zval = ZendHashStrFind(ExecutorGlobals.GetFunctionTable(), name, len_)
	if EXPECTED(zv != nil) {
		var fbc *ZendFunction = Z_FUNC_P(zv)
		if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
			InitFuncRunTimeCacheI(&fbc.op_array)
		}
		return fbc
	}
	return nil
}
func ZendInitFuncRunTimeCache(op_array *ZendOpArray) {
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCacheI(op_array)
	}
}
func IInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	ZEND_ASSERT(EX(func_) == (*ZendFunction)(op_array))
	EX(opline) = op_array.GetOpcodes()
	EX(call) = nil
	EX(return_value) = return_value
	ZendAttachSymbolTable(execute_data)
	if op_array.GetRunTimeCachePtr() == nil {
		var ptr any
		ZEND_ASSERT((op_array.GetFnFlags() & ZEND_ACC_HEAP_RT_CACHE) != 0)
		ptr = Emalloc(op_array.GetCacheSize() + b.SizeOf("void *"))
		ZEND_MAP_PTR_INIT(op_array.run_time_cache, ptr)
		ptr = (*byte)(ptr + b.SizeOf("void *"))
		ZEND_MAP_PTR_SET(op_array.run_time_cache, ptr)
		memset(ptr, 0, op_array.GetCacheSize())
	}
	EX(run_time_cache) = RUN_TIME_CACHE(op_array)
	ExecutorGlobals.SetCurrentExecuteData(execute_data)
}

/* }}} */

func ZendInitFuncExecuteData(ex *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	var execute_data *ZendExecuteData = ex
	EX(prev_execute_data) = ExecutorGlobals.GetCurrentExecuteData()
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCache(op_array)
	}
	IInitFuncExecuteData(op_array, return_value, 1, EXECUTE_DATA_C)
}

/* }}} */

func ZendInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	EX(prev_execute_data) = ExecutorGlobals.GetCurrentExecuteData()
	IInitCodeExecuteData(execute_data, op_array, return_value)
}

/* }}} */

func ZendInitExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		ZendInitCodeExecuteData(execute_data, op_array, return_value)
	} else {
		ZendInitFuncExecuteData(execute_data, op_array, return_value)
	}
}

/* }}} */

func ZendVmStackCopyCallFrame(call *ZendExecuteData, passed_args uint32, additional_args uint32) *ZendExecuteData {
	var new_call *ZendExecuteData
	var used_stack int = ExecutorGlobals.GetVmStackTop() - (*Zval)(call) + additional_args

	/* copy call frame into new stack segment */

	new_call = ZendVmStackExtend(used_stack * b.SizeOf("zval"))
	*new_call = *call
	ZEND_ADD_CALL_FLAG(new_call, ZEND_CALL_ALLOCATED)
	if passed_args != 0 {
		var src *Zval = ZEND_CALL_ARG(call, 1)
		var dst *Zval = ZEND_CALL_ARG(new_call, 1)
		for {
			ZVAL_COPY_VALUE(dst, src)
			passed_args--
			src++
			dst++
			if passed_args == 0 {
				break
			}
		}
	}

	/* delete old call_frame from previous stack segment */

	ExecutorGlobals.GetVmStack().GetPrev().SetTop((*Zval)(call))

	/* delete previous stack segment if it became empty */

	if UNEXPECTED(ExecutorGlobals.GetVmStack().GetPrev().GetTop() == ZEND_VM_STACK_ELEMENTS(ExecutorGlobals.GetVmStack().GetPrev())) {
		var r ZendVmStack = ExecutorGlobals.GetVmStack().GetPrev()
		ExecutorGlobals.GetVmStack().SetPrev(r.GetPrev())
		Efree(r)
	}
	return new_call
}

/* }}} */

func ZendVmStackExtendCallFrame(call **ZendExecuteData, passed_args uint32, additional_args uint32) {
	if EXPECTED(uint32_t(ExecutorGlobals.GetVmStackEnd()-ExecutorGlobals.GetVmStackTop()) > additional_args) {
		ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStackTop() + additional_args)
	} else {
		*call = ZendVmStackCopyCallFrame(*call, passed_args, additional_args)
	}
}

/* }}} */

func ZendGetRunningGenerator(EXECUTE_DATA_D) *ZendGenerator {
	/* The generator object is stored in EX(return_value) */

	var generator *ZendGenerator = (*ZendGenerator)(EX(return_value))

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */

	return generator

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */
}

/* }}} */

func CleanupUnfinishedCalls(execute_data *ZendExecuteData, op_num uint32) {
	if UNEXPECTED(EX(call)) {
		var call *ZendExecuteData = EX(call)
		var opline *ZendOp = EX(func_).op_array.opcodes + op_num
		var level int
		var do_exit int
		if UNEXPECTED(opline.GetOpcode() == ZEND_INIT_FCALL || opline.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_DYNAMIC_CALL || opline.GetOpcode() == ZEND_INIT_USER_CALL || opline.GetOpcode() == ZEND_INIT_METHOD_CALL || opline.GetOpcode() == ZEND_INIT_STATIC_METHOD_CALL || opline.GetOpcode() == ZEND_NEW) {
			ZEND_ASSERT(op_num != 0)
			opline--
		}
		for {

			/* If the exception was thrown during a function call there might be
			 * arguments pushed to the stack that have to be dtor'ed. */

			level = 0
			do_exit = 0
			for {
				switch opline.GetOpcode() {
				case ZEND_DO_FCALL:

				case ZEND_DO_ICALL:

				case ZEND_DO_UCALL:

				case ZEND_DO_FCALL_BY_NAME:
					level++
					break
				case ZEND_INIT_FCALL:

				case ZEND_INIT_FCALL_BY_NAME:

				case ZEND_INIT_NS_FCALL_BY_NAME:

				case ZEND_INIT_DYNAMIC_CALL:

				case ZEND_INIT_USER_CALL:

				case ZEND_INIT_METHOD_CALL:

				case ZEND_INIT_STATIC_METHOD_CALL:

				case ZEND_NEW:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = 0
						do_exit = 1
					}
					level--
					break
				case ZEND_SEND_VAL:

				case ZEND_SEND_VAL_EX:

				case ZEND_SEND_VAR:

				case ZEND_SEND_VAR_EX:

				case ZEND_SEND_FUNC_ARG:

				case ZEND_SEND_REF:

				case ZEND_SEND_VAR_NO_REF:

				case ZEND_SEND_VAR_NO_REF_EX:

				case ZEND_SEND_USER:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = opline.GetOp2().GetNum()
						do_exit = 1
					}
					break
				case ZEND_SEND_ARRAY:

				case ZEND_SEND_UNPACK:
					if level == 0 {
						do_exit = 1
					}
					break
				}
				if do_exit == 0 {
					opline--
				}
				if do_exit != 0 {
					break
				}
			}
			if call.GetPrevExecuteData() != nil {

				/* skip current call region */

				level = 0
				do_exit = 0
				for {
					switch opline.GetOpcode() {
					case ZEND_DO_FCALL:

					case ZEND_DO_ICALL:

					case ZEND_DO_UCALL:

					case ZEND_DO_FCALL_BY_NAME:
						level++
						break
					case ZEND_INIT_FCALL:

					case ZEND_INIT_FCALL_BY_NAME:

					case ZEND_INIT_NS_FCALL_BY_NAME:

					case ZEND_INIT_DYNAMIC_CALL:

					case ZEND_INIT_USER_CALL:

					case ZEND_INIT_METHOD_CALL:

					case ZEND_INIT_STATIC_METHOD_CALL:

					case ZEND_NEW:
						if level == 0 {
							do_exit = 1
						}
						level--
						break
					}
					opline--
					if do_exit != 0 {
						break
					}
				}
			}
			ZendVmStackFreeArgs(EX(call))
			if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
				OBJ_RELEASE(Z_OBJ(call.GetThis()))
			}
			if (call.GetFunc().GetFnFlags() & ZEND_ACC_CLOSURE) != 0 {
				ZendObjectRelease(ZEND_CLOSURE_OBJECT(call.GetFunc()))
			} else if (call.GetFunc().GetFnFlags() & ZEND_ACC_CALL_VIA_TRAMPOLINE) != 0 {
				ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
				ZendFreeTrampoline(call.GetFunc())
			}
			EX(call) = call.GetPrevExecuteData()
			ZendVmStackFreeCallFrame(call)
			call = EX(call)
			if call == nil {
				break
			}
		}
	}
}

/* }}} */

func FindLiveRange(op_array *ZendOpArray, op_num uint32, var_num uint32) *ZendLiveRange {
	var i int
	for i = 0; i < op_array.GetLastLiveRange(); i++ {
		var range_ *ZendLiveRange = &op_array.live_range[i]
		if op_num >= range_.GetStart() && op_num < range_.GetEnd() && var_num == (range_.GetVar() & ^ZEND_LIVE_MASK) {
			return range_
		}
	}
	return nil
}

/* }}} */

func CleanupLiveVars(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	var i int
	for i = 0; i < EX(func_).op_array.last_live_range; i++ {
		var range_ *ZendLiveRange = &EX(func_).op_array.live_range[i]
		if range_.GetStart() > op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		} else if op_num < range_.GetEnd() {
			if catch_op_num == 0 || catch_op_num >= range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *Zval = EX_VAR(var_num)
				if kind == ZEND_LIVE_TMPVAR {
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_NEW {
					var obj *ZendObject
					ZEND_ASSERT(Z_TYPE_P(var_) == IS_OBJECT)
					obj = Z_OBJ_P(var_)
					ZendObjectStoreCtorFailed(obj)
					OBJ_RELEASE(obj)
				} else if kind == ZEND_LIVE_LOOP {
					if Z_TYPE_P(var_) != IS_ARRAY && Z_FE_ITER_P(var_) != uint32_t-1 {
						ZendHashIteratorDel(Z_FE_ITER_P(var_))
					}
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_ROPE {
					var rope **ZendString = (**ZendString)(var_)
					var last *ZendOp = EX(func_).op_array.opcodes + op_num
					for last.GetOpcode() != ZEND_ROPE_ADD && last.GetOpcode() != ZEND_ROPE_INIT || last.GetResult().GetVar() != var_num {
						ZEND_ASSERT(last >= EX(func_).op_array.opcodes)
						last--
					}
					if last.GetOpcode() == ZEND_ROPE_INIT {
						ZendStringReleaseEx(*rope, 0)
					} else {
						var j int = last.GetExtendedValue()
						for {
							ZendStringReleaseEx(rope[j], 0)
							if !(b.PostDec(&j)) {
								break
							}
						}
					}
				} else if kind == ZEND_LIVE_SILENCE {

					/* restore previous error_reporting value */

					if ExecutorGlobals.GetErrorReporting() == 0 && Z_LVAL_P(var_) != 0 {
						ExecutorGlobals.SetErrorReporting(Z_LVAL_P(var_))
					}

					/* restore previous error_reporting value */

				}
			}
		}
	}
}

/* }}} */

func ZendCleanupUnfinishedExecution(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	CleanupUnfinishedCalls(execute_data, op_num)
	CleanupLiveVars(execute_data, op_num, catch_op_num)
}
func ZendSwapOperands(op *ZendOp) {
	var tmp ZnodeOp
	var tmp_type ZendUchar
	tmp = op.GetOp1()
	tmp_type = op.GetOp1Type()
	op.SetOp1(op.GetOp2())
	op.SetOp1Type(op.GetOp2Type())
	op.SetOp2(tmp)
	op.SetOp2Type(tmp_type)
}

/* }}} */

func ZendInitDynamicCallString(function *ZendString, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var func_ *Zval
	var called_scope *ZendClassEntry
	var lcname *ZendString
	var colon *byte
	if b.Assign(&colon, ZendMemrchr(ZSTR_VAL(function), ':', ZSTR_LEN(function))) != nil && colon > ZSTR_VAL(function) && (*(colon - 1)) == ':' {
		var mname *ZendString
		var cname_length int = colon - ZSTR_VAL(function) - 1
		var mname_length int = ZSTR_LEN(function) - cname_length - (b.SizeOf("\"::\"") - 1)
		lcname = ZendStringInit(ZSTR_VAL(function), cname_length, 0)
		called_scope = ZendFetchClassByName(lcname, nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
		if UNEXPECTED(called_scope == nil) {
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		mname = ZendStringInit(ZSTR_VAL(function)+(cname_length+b.SizeOf("\"::\"")-1), mname_length, 0)
		if called_scope.GetGetStaticMethod() != nil {
			fbc = called_scope.GetGetStaticMethod()(called_scope, mname)
		} else {
			fbc = ZendStdGetStaticMethod(called_scope, mname, nil)
		}
		if UNEXPECTED(fbc == nil) {
			if EXPECTED(ExecutorGlobals.GetException() == nil) {
				ZendUndefinedMethod(called_scope, mname)
			}
			ZendStringReleaseEx(lcname, 0)
			ZendStringReleaseEx(mname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		ZendStringReleaseEx(mname, 0)
		if UNEXPECTED((fbc.GetFnFlags() & ZEND_ACC_STATIC) == 0) {
			ZendNonStaticMethodCall(fbc)
			if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
				return nil
			}
		}
		if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	} else {
		if ZSTR_VAL(function)[0] == '\\' {
			lcname = ZendStringAlloc(ZSTR_LEN(function)-1, 0)
			ZendStrTolowerCopy(ZSTR_VAL(lcname), ZSTR_VAL(function)+1, ZSTR_LEN(function)-1)
		} else {
			lcname = ZendStringTolower(function)
		}
		if UNEXPECTED(b.Assign(&func_, ZendHashFind(ExecutorGlobals.GetFunctionTable(), lcname)) == nil) {
			ZendThrowError(nil, "Call to undefined function %s()", ZSTR_VAL(function))
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		fbc = Z_FUNC_P(func_)
		if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		called_scope = nil
	}
	return ZendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION|ZEND_CALL_DYNAMIC, fbc, num_args, called_scope)
}

/* }}} */

func ZendInitDynamicCallObject(function *Zval, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var called_scope *ZendClassEntry
	var object *ZendObject
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if EXPECTED(Z_OBJ_HANDLER_P(function, get_closure)) && EXPECTED(Z_OBJ_HANDLER_P(function, get_closure)(function, &called_scope, &fbc, &object) == SUCCESS) {
		object_or_called_scope = called_scope
		if (fbc.GetFnFlags() & ZEND_ACC_CLOSURE) != 0 {

			/* Delay closure destruction until its invocation */

			GC_ADDREF(ZEND_CLOSURE_OBJECT(fbc))
			call_info |= ZEND_CALL_CLOSURE
			if (fbc.GetFnFlags() & ZEND_ACC_FAKE_CLOSURE) != 0 {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if object != nil {
				call_info |= ZEND_CALL_HAS_THIS
				object_or_called_scope = object
			}
		} else if object != nil {
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
			GC_ADDREF(object)
			object_or_called_scope = object
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
		InitFuncRunTimeCache(&fbc.op_array)
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}

/* }}} */

func ZendInitDynamicCallArray(function *ZendArray, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if ZendHashNumElements(function) == 2 {
		var obj *Zval
		var method *Zval
		obj = ZendHashIndexFind(function, 0)
		method = ZendHashIndexFind(function, 1)
		if UNEXPECTED(obj == nil) || UNEXPECTED(method == nil) {
			ZendThrowError(nil, "Array callback has to contain indices 0 and 1")
			return nil
		}
		ZVAL_DEREF(obj)
		if UNEXPECTED(Z_TYPE_P(obj) != IS_STRING) && UNEXPECTED(Z_TYPE_P(obj) != IS_OBJECT) {
			ZendThrowError(nil, "First array member is not a valid class name or object")
			return nil
		}
		ZVAL_DEREF(method)
		if UNEXPECTED(Z_TYPE_P(method) != IS_STRING) {
			ZendThrowError(nil, "Second array member is not a valid method")
			return nil
		}
		if Z_TYPE_P(obj) == IS_STRING {
			var called_scope *ZendClassEntry = ZendFetchClassByName(Z_STR_P(obj), nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if UNEXPECTED(called_scope == nil) {
				return nil
			}
			if called_scope.GetGetStaticMethod() != nil {
				fbc = called_scope.GetGetStaticMethod()(called_scope, Z_STR_P(method))
			} else {
				fbc = ZendStdGetStaticMethod(called_scope, Z_STR_P(method), nil)
			}
			if UNEXPECTED(fbc == nil) {
				if EXPECTED(ExecutorGlobals.GetException() == nil) {
					ZendUndefinedMethod(called_scope, Z_STR_P(method))
				}
				return nil
			}
			if (fbc.GetFnFlags() & ZEND_ACC_STATIC) == 0 {
				ZendNonStaticMethodCall(fbc)
				if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
					return nil
				}
			}
			object_or_called_scope = called_scope
		} else {
			var object *ZendObject = Z_OBJ_P(obj)
			fbc = Z_OBJ_HT_P(obj).GetGetMethod()(&object, Z_STR_P(method), nil)
			if UNEXPECTED(fbc == nil) {
				if EXPECTED(ExecutorGlobals.GetException() == nil) {
					ZendUndefinedMethod(object.GetCe(), Z_STR_P(method))
				}
				return nil
			}
			if (fbc.GetFnFlags() & ZEND_ACC_STATIC) != 0 {
				object_or_called_scope = object.GetCe()
			} else {
				call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
				GC_ADDREF(object)
				object_or_called_scope = object
			}
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if EXPECTED(fbc.GetType() == ZEND_USER_FUNCTION) && UNEXPECTED(!(RUN_TIME_CACHE(&fbc.op_array))) {
		InitFuncRunTimeCache(&fbc.op_array)
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}

/* }}} */

const ZEND_FAKE_OP_ARRAY *ZendOpArray = (*ZendOpArray)(zend_intptr_t - 1)

func ZendIncludeOrEval(inc_filename *Zval, type_ int) *ZendOpArray {
	var new_op_array *ZendOpArray = nil
	var tmp_inc_filename Zval
	ZVAL_UNDEF(&tmp_inc_filename)
	if Z_TYPE_P(inc_filename) != IS_STRING {
		var tmp *ZendString = ZvalTryGetStringFunc(inc_filename)
		if UNEXPECTED(tmp == nil) {
			return nil
		}
		ZVAL_STR(&tmp_inc_filename, tmp)
		inc_filename = &tmp_inc_filename
	}
	switch type_ {
	case ZEND_INCLUDE_ONCE:

	case ZEND_REQUIRE_ONCE:
		var file_handle ZendFileHandle
		var resolved_path *ZendString
		resolved_path = ZendResolvePath(Z_STRVAL_P(inc_filename), Z_STRLEN_P(inc_filename))
		if EXPECTED(resolved_path != nil) {
			if ZendHashExists(&(ExecutorGlobals.GetIncludedFiles()), resolved_path) != 0 {
				goto already_compiled
			}
		} else if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			break
		} else if UNEXPECTED(strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename)) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		} else {
			resolved_path = ZendStringCopy(Z_STR_P(inc_filename))
		}
		if SUCCESS == ZendStreamOpen(ZSTR_VAL(resolved_path), &file_handle) {
			if file_handle.GetOpenedPath() == nil {
				file_handle.SetOpenedPath(ZendStringCopy(resolved_path))
			}
			if ZendHashAddEmptyElement(&(ExecutorGlobals.GetIncludedFiles()), file_handle.GetOpenedPath()) != nil {
				var op_array *ZendOpArray = ZendCompileFile(&file_handle, b.Cond(type_ == ZEND_INCLUDE_ONCE, ZEND_INCLUDE, ZEND_REQUIRE))
				ZendDestroyFileHandle(&file_handle)
				ZendStringReleaseEx(resolved_path, 0)
				if Z_TYPE(tmp_inc_filename) != IS_UNDEF {
					ZvalPtrDtorStr(&tmp_inc_filename)
				}
				return op_array
			} else {
				ZendFileHandleDtor(&file_handle)
			already_compiled:
				new_op_array = ZEND_FAKE_OP_ARRAY
			}
		} else {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
		}
		ZendStringReleaseEx(resolved_path, 0)
		break
	case ZEND_INCLUDE:

	case ZEND_REQUIRE:
		if UNEXPECTED(strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename)) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		}
		new_op_array = CompileFilename(type_, inc_filename)
		break
	case ZEND_EVAL:
		var eval_desc *byte = ZendMakeCompiledStringDescription("eval()'d code")
		new_op_array = ZendCompileString(inc_filename, eval_desc)
		Efree(eval_desc)
		break
	default:
		break
	}
	if Z_TYPE(tmp_inc_filename) != IS_UNDEF {
		ZvalPtrDtorStr(&tmp_inc_filename)
	}
	return new_op_array
}

/* }}} */

func ZendDoFcallOverloaded(call *ZendExecuteData, ret *Zval) int {
	var fbc *ZendFunction = call.GetFunc()
	var object *ZendObject

	/* Not sure what should be done here if it's a static method */

	if UNEXPECTED(Z_TYPE(call.GetThis()) != IS_OBJECT) {
		ZendVmStackFreeArgs(call)
		if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			ZendStringReleaseEx(fbc.GetFunctionName(), 0)
		}
		Efree(fbc)
		ZendVmStackFreeCallFrame(call)
		ZendThrowError(nil, "Cannot call overloaded function for non-object")
		return 0
	}
	object = Z_OBJ(call.GetThis())
	ZVAL_NULL(ret)
	ExecutorGlobals.SetCurrentExecuteData(call)
	object.GetHandlers().GetCallMethod()(fbc.GetFunctionName(), object, call, ret)
	ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
	ZendVmStackFreeArgs(call)
	if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
		ZendStringReleaseEx(fbc.GetFunctionName(), 0)
	}
	Efree(fbc)
	return 1
}

/* }}} */

func ZendFeResetIterator(array_ptr *Zval, by_ref int, opline *ZendOp, _ EXECUTE_DATA_D) ZendBool {
	var ce *ZendClassEntry = Z_OBJCE_P(array_ptr)
	var iter *ZendObjectIterator = ce.GetGetIterator()(ce, array_ptr, by_ref)
	var is_empty ZendBool
	if UNEXPECTED(iter == nil) || UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		if iter != nil {
			OBJ_RELEASE(&iter.std)
		}
		if ExecutorGlobals.GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ZSTR_VAL(ce.GetName()))
		}
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return 1
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			OBJ_RELEASE(&iter.std)
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			return 1
		}
	}
	is_empty = iter.GetFuncs().GetValid()(iter) != SUCCESS
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		OBJ_RELEASE(&iter.std)
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return 1
	}
	iter.SetIndex(-1)
	ZVAL_OBJ(EX_VAR(opline.GetResult().GetVar()), &iter.std)
	Z_FE_ITER_P(EX_VAR(opline.GetResult().GetVar())) = uint32_t - 1
	return is_empty
}

/* }}} */

func _zendQuickGetConstant(key *Zval, flags uint32, check_defined_only int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var zv *Zval
	var orig_key *Zval = key
	var c *ZendConstant = nil
	zv = ZendHashFindEx(ExecutorGlobals.GetZendConstants(), Z_STR_P(key), 1)
	if zv != nil {
		c = (*ZendConstant)(Z_PTR_P(zv))
	} else {
		key++
		zv = ZendHashFindEx(ExecutorGlobals.GetZendConstants(), Z_STR_P(key), 1)
		if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(Z_PTR_P(zv)))&CONST_CS) == 0 {
			c = (*ZendConstant)(Z_PTR_P(zv))
		} else {
			if (flags & (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED)) == (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED) {
				key++
				zv = ZendHashFindEx(ExecutorGlobals.GetZendConstants(), Z_STR_P(key), 1)
				if zv != nil {
					c = (*ZendConstant)(Z_PTR_P(zv))
				} else {
					key++
					zv = ZendHashFindEx(ExecutorGlobals.GetZendConstants(), Z_STR_P(key), 1)
					if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(Z_PTR_P(zv)))&CONST_CS) == 0 {
						c = (*ZendConstant)(Z_PTR_P(zv))
					}
				}
			}
		}
	}
	if c == nil {
		if check_defined_only == 0 {
			if (opline.GetOp1().GetNum() & IS_CONSTANT_UNQUALIFIED) != 0 {
				var actual *byte = (*byte)(ZendMemrchr(Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())), '\\', Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))))
				if actual == nil {
					ZVAL_STR_COPY(EX_VAR(opline.GetResult().GetVar()), Z_STR_P(RT_CONSTANT(opline, opline.GetOp2())))
				} else {
					actual++
					ZVAL_STRINGL(EX_VAR(opline.GetResult().GetVar()), actual, Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))-(actual-Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2()))))
				}

				/* non-qualified constant - allow text substitution */

				ZendError(E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())), Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())))

				/* non-qualified constant - allow text substitution */

			} else {
				ZendThrowError(nil, "Undefined constant '%s'", Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())))
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
		}
		return FAILURE
	}
	if check_defined_only == 0 {
		ZVAL_COPY_OR_DUP(EX_VAR(opline.GetResult().GetVar()), &c.value)
		if (ZEND_CONSTANT_FLAGS(c) & (CONST_CS | CONST_CT_SUBST)) == 0 {
			var ns_sep *byte
			var shortname_offset int
			var shortname_len int
			var is_deprecated ZendBool
			if (flags & IS_CONSTANT_UNQUALIFIED) != 0 {
				var access_key *Zval
				if (flags & IS_CONSTANT_IN_NAMESPACE) == 0 {
					access_key = orig_key - 1
				} else {
					if key < orig_key+2 {
						goto check_short_name
					} else {
						access_key = orig_key + 2
					}
				}
				is_deprecated = !(ZendStringEquals(c.GetName(), Z_STR_P(access_key)))
			} else {
			check_short_name:

				/* Namespaces are always case-insensitive. Only compare shortname. */

				ns_sep = ZendMemrchr(ZSTR_VAL(c.GetName()), '\\', ZSTR_LEN(c.GetName()))
				if ns_sep != nil {
					shortname_offset = ns_sep - ZSTR_VAL(c.GetName()) + 1
					shortname_len = ZSTR_LEN(c.GetName()) - shortname_offset
				} else {
					shortname_offset = 0
					shortname_len = ZSTR_LEN(c.GetName())
				}
				is_deprecated = memcmp(ZSTR_VAL(c.GetName())+shortname_offset, Z_STRVAL_P(orig_key-1)+shortname_offset, shortname_len) != 0
			}
			if is_deprecated != 0 {
				ZendError(E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", ZSTR_VAL(c.GetName()))
				return SUCCESS
			}
		}
	}
	CACHE_PTR(opline.GetExtendedValue(), c)
	return SUCCESS
}

/* }}} */

func ZendQuickGetConstant(key *Zval, flags uint32, opline *ZendOp, _ EXECUTE_DATA_D) {
	_zendQuickGetConstant(key, flags, 0, OPLINE_C, EXECUTE_DATA_C)
}
func ZendQuickCheckConstant(key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) int {
	return _zendQuickGetConstant(key, 0, 1, OPLINE_C, EXECUTE_DATA_C)
}

const _zendVmStackPushCallFrameEx = ZendVmStackPushCallFrameEx
const _zendVmStackPushCallFrame = ZendVmStackPushCallFrame

// #define ZEND_VM_NEXT_OPCODE_EX(check_exception,skip) CHECK_SYMBOL_TABLES ( ) if ( check_exception ) { OPLINE = EX ( opline ) + ( skip ) ; } else { ZEND_ASSERT ( ! EG ( exception ) ) ; OPLINE = opline + ( skip ) ; } ZEND_VM_CONTINUE ( )

func ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION() {
	OPLINE = EX(opline) + 1
	ZEND_VM_CONTINUE()
}
func ZEND_VM_NEXT_OPCODE() {
	ZEND_ASSERT(ExecutorGlobals.GetException() == nil)
	OPLINE = opline + 1
	ZEND_VM_CONTINUE()
}

// #define ZEND_VM_SET_NEXT_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op

// #define ZEND_VM_SET_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op ; ZEND_VM_INTERRUPT_CHECK ( )

func ZEND_VM_SET_RELATIVE_OPCODE(opline *ZendOp, offset uint32) {
	OPLINE = ZEND_OFFSET_TO_OPLINE(opline, offset)
	ZEND_VM_INTERRUPT_CHECK()
}
func ZEND_VM_JMP_EX(new_op *ZendOp, check_exception int) {
	if check_exception != 0 && UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		HANDLE_EXCEPTION()
	}
	OPLINE = new_op
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_JMP(new_op *ZendOp) { ZEND_VM_JMP_EX(new_op, 1) }
func ZEND_VM_INC_OPCODE() int {
	OPLINE++
	return OPLINE - 1
}

const VM_SMART_OPCODES = 1

// #define ZEND_VM_REPEATABLE_OPCODE       do {

// #define ZEND_VM_REPEAT_OPCODE(_opcode) } while ( UNEXPECTED ( ( ++ opline ) -> opcode == _opcode ) ) ; OPLINE = opline ; ZEND_VM_CONTINUE ( )

func ZEND_VM_SMART_BRANCH(_result __auto__, _check int) {
	for {
		if _check != 0 && UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			break
		}
		if EXPECTED((opline + 1).opcode == ZEND_JMPZ) {
			if _result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else if EXPECTED((opline + 1).opcode == ZEND_JMPNZ) {
			if !_result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else {
			break
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPZ(_result int, _check int) {
	for {
		if _check != 0 && UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			break
		}
		if _result != 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPNZ(_result int, _check int) {
	for {
		if _check != 0 && UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			break
		}
		if _result == 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_TRUE() {
	if EXPECTED((opline + 1).opcode == ZEND_JMPNZ) {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	} else if EXPECTED((opline + 1).opcode == ZEND_JMPZ) {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPNZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE() {
	if EXPECTED((opline + 1).opcode == ZEND_JMPNZ) {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	} else if EXPECTED((opline + 1).opcode == ZEND_JMPZ) {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPNZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}

// #define ZEND_VM_GUARD(name)

func UNDEF_RESULT() {
	if (opline.result_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		ZVAL_UNDEF(EX_VAR(opline.result.var_))
	}
}

// # include "zend_vm_execute.h"

func ZendSetUserOpcodeHandler(opcode ZendUchar, handler UserOpcodeHandlerT) int {
	if opcode != ZEND_USER_OPCODE {
		if handler == nil {

			/* restore the original handler */

			ZendUserOpcodes[opcode] = opcode

			/* restore the original handler */

		} else {
			ZendUserOpcodes[opcode] = ZEND_USER_OPCODE
		}
		ZendUserOpcodeHandlers[opcode] = handler
		return SUCCESS
	}
	return FAILURE
}
func ZendGetUserOpcodeHandler(opcode ZendUchar) UserOpcodeHandlerT {
	return ZendUserOpcodeHandlers[opcode]
}
func ZendGetZvalPtr(opline *ZendOp, op_type int, node *ZnodeOp, execute_data *ZendExecuteData, should_free *ZendFreeOp, type_ int) *Zval {
	var ret *Zval
	switch op_type {
	case IS_CONST:
		ret = RT_CONSTANT(opline, *node)
		*should_free = nil
		break
	case IS_TMP_VAR:

	case IS_VAR:
		ret = EX_VAR(node.GetVar())
		*should_free = ret
		break
	case IS_CV:
		ret = EX_VAR(node.GetVar())
		*should_free = nil
		break
	default:
		ret = nil
		*should_free = ret
		break
	}
	return ret
}

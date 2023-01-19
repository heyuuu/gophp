// <<generate>>

package zend

import g "sik/runtime/grammar"

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

// #define ZEND_REF_TYPE_SOURCES(ref) ( ref ) -> sources

// #define ZEND_REF_HAS_TYPE_SOURCES(ref) ( ZEND_REF_TYPE_SOURCES ( ref ) . ptr != NULL )

// #define ZEND_REF_FIRST_SOURCE(ref) ( ZEND_PROPERTY_INFO_SOURCE_IS_LIST ( ( ref ) -> sources . list ) ? ZEND_PROPERTY_INFO_SOURCE_TO_LIST ( ( ref ) -> sources . list ) -> ptr [ 0 ] : ( ref ) -> sources . ptr )

func ZendCopyToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, ref *ZendRefcounted) {
	var _z1 *Zval = variable_ptr
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (value_type & (1<<0 | 1<<3)) != 0 {
		if (variable_ptr.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(variable_ptr)
		}
	} else if ref != nil {
		if ZendGcDelref(&ref.gc) == 0 {
			_efree(ref)
		} else if (variable_ptr.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(variable_ptr)
		}
	}
}
func ZendAssignToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, strict ZendBool) *Zval {
	var ref *ZendRefcounted = nil
	if value.GetType() == 10 {
		ref = value.GetValue().GetCounted()
		value = &(*value).value.GetRef().GetVal()
	}
	for {
		if variable_ptr.GetTypeFlags() != 0 {
			var garbage *ZendRefcounted
			if variable_ptr.GetType() == 10 {
				if variable_ptr.GetValue().GetRef().GetSources().GetPtr() != nil {
					return ZendAssignToTypedRef(variable_ptr, value, value_type, strict, ref)
				}
				variable_ptr = &(*variable_ptr).value.GetRef().GetVal()
				if variable_ptr.GetTypeFlags() == 0 {
					break
				}
			}
			if variable_ptr.GetType() == 8 && variable_ptr.GetValue().GetObj().GetHandlers().GetSet() != nil {
				variable_ptr.GetValue().GetObj().GetHandlers().GetSet()(variable_ptr, value)
				return variable_ptr
			}
			garbage = variable_ptr.GetValue().GetCounted()
			ZendCopyToVariable(variable_ptr, value, value_type, ref)
			if ZendGcDelref(&garbage.gc) == 0 {
				RcDtorFunc(garbage)
			} else {

				/* optimized version of GC_ZVAL_CHECK_POSSIBLE_ROOT(variable_ptr) */

				if (garbage.GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
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

// @type _zendVmStack struct

// #define ZEND_VM_STACK_HEADER_SLOTS       ( ( ZEND_MM_ALIGNED_SIZE ( sizeof ( struct _zend_vm_stack ) ) + ZEND_MM_ALIGNED_SIZE ( sizeof ( zval ) ) - 1 ) / ZEND_MM_ALIGNED_SIZE ( sizeof ( zval ) ) )

// #define ZEND_VM_STACK_ELEMENTS(stack) ( ( ( zval * ) ( stack ) ) + ZEND_VM_STACK_HEADER_SLOTS )

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
	call.GetThis().GetValue().SetPtr(object_or_called_scope)
	call.GetThis().SetTypeInfo(call_info)
	call.GetThis().SetNumArgs(num_args)
}
func ZendVmStackPushCallFrameEx(used_stack uint32, call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var call *ZendExecuteData = (*ZendExecuteData)(EG.GetVmStackTop())
	if used_stack > size_t((*byte)(EG.GetVmStackEnd())-(*byte)(call)) {
		call = (*ZendExecuteData)(ZendVmStackExtend(used_stack))
		ZendVmInitCallFrame(call, call_info|1<<18, func_, num_args, object_or_called_scope)
		return call
	} else {
		EG.SetVmStackTop((*Zval)((*byte)(call + used_stack)))
		ZendVmInitCallFrame(call, call_info, func_, num_args, object_or_called_scope)
		return call
	}
}
func ZendVmCalcUsedStack(num_args uint32, func_ *ZendFunction) uint32 {
	var used_stack uint32 = int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + num_args
	if (func_.GetType() & 1) == 0 {
		used_stack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - g.CondF1(func_.GetOpArray().GetNumArgs() < num_args, func() uint32 { return func_.GetOpArray().GetNumArgs() }, num_args)
	}
	return used_stack * g.SizeOf("zval")
}
func ZendVmStackPushCallFrame(call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var used_stack uint32 = ZendVmCalcUsedStack(num_args, func_)
	return ZendVmStackPushCallFrameEx(used_stack, call_info, func_, num_args, object_or_called_scope)
}
func ZendVmStackFreeExtraArgsEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & 1 << 19) != 0 {
		var count uint32 = call.GetThis().GetNumArgs() - call.GetFunc().GetOpArray().GetNumArgs()
		var p *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(call.GetFunc().GetOpArray().GetLastVar()+call.GetFunc().GetOpArray().GetT()))
		for {
			if p.GetTypeFlags() != 0 {
				var r *ZendRefcounted = p.GetValue().GetCounted()
				if ZendGcDelref(&r.gc) == 0 {
					p.SetTypeInfo(1)
					RcDtorFunc(r)
				} else {
					GcCheckPossibleRoot(r)
				}
			}
			p++
			if !(g.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendVmStackFreeExtraArgs(call *ZendExecuteData) {
	ZendVmStackFreeExtraArgsEx(call.GetThis().GetTypeInfo(), call)
}
func ZendVmStackFreeArgs(call *ZendExecuteData) {
	var num_args uint32 = call.GetThis().GetNumArgs()
	if num_args > 0 {
		var p *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		for {
			if p.GetTypeFlags() != 0 {
				var r *ZendRefcounted = p.GetValue().GetCounted()
				if ZendGcDelref(&r.gc) == 0 {
					p.SetTypeInfo(1)
					RcDtorFunc(r)
				}
			}
			p++
			if !(g.PreDec(&num_args)) {
				break
			}
		}
	}
}
func ZendVmStackFreeCallFrameEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & 1 << 18) != 0 {
		var p ZendVmStack = EG.GetVmStack()
		var prev ZendVmStack = p.GetPrev()
		assert(call == (*ZendExecuteData)((*Zval)(EG.GetVmStack())+((g.SizeOf("struct _zend_vm_stack")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))))
		EG.SetVmStackTop(prev.GetTop())
		EG.SetVmStackEnd(prev.GetEnd())
		EG.SetVmStack(prev)
		_efree(p)
	} else {
		EG.SetVmStackTop((*Zval)(call))
	}
}
func ZendVmStackFreeCallFrame(call *ZendExecuteData) {
	ZendVmStackFreeCallFrameEx(call.GetThis().GetTypeInfo(), call)
}

/* services */

// #define ZEND_USER_OPCODE_CONTINUE       0

// #define ZEND_USER_OPCODE_RETURN       1

// #define ZEND_USER_OPCODE_DISPATCH       2

// #define ZEND_USER_OPCODE_ENTER       3

// #define ZEND_USER_OPCODE_LEAVE       4

// #define ZEND_USER_OPCODE_DISPATCH_TO       0x100

/* former zend_execute_locks.h */

type ZendFreeOp *Zval

// #define CACHE_ADDR(num) ( ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) )

// #define CACHED_PTR(num) ( ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) ) [ 0 ]

// #define CACHE_PTR(num,ptr) do { ( ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) ) [ 0 ] = ( ptr ) ; } while ( 0 )

// #define CACHED_POLYMORPHIC_PTR(num,ce) ( EXPECTED ( ( ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) ) [ 0 ] == ( void * ) ( ce ) ) ? ( ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) ) [ 1 ] : NULL )

// #define CACHE_POLYMORPHIC_PTR(num,ce,ptr) do { void * * slot = ( void * * ) ( ( char * ) EX ( run_time_cache ) + ( num ) ) ; slot [ 0 ] = ( ce ) ; slot [ 1 ] = ( ptr ) ; } while ( 0 )

// #define CACHED_PTR_EX(slot) ( slot ) [ 0 ]

// #define CACHE_PTR_EX(slot,ptr) do { ( slot ) [ 0 ] = ( ptr ) ; } while ( 0 )

// #define CACHED_POLYMORPHIC_PTR_EX(slot,ce) ( EXPECTED ( ( slot ) [ 0 ] == ( ce ) ) ? ( slot ) [ 1 ] : NULL )

// #define CACHE_POLYMORPHIC_PTR_EX(slot,ce,ptr) do { ( slot ) [ 0 ] = ( ce ) ; ( slot ) [ 1 ] = ( ptr ) ; } while ( 0 )

// #define CACHE_SPECIAL       ( 1 << 0 )

// #define IS_SPECIAL_CACHE_VAL(ptr) ( ( ( uintptr_t ) ( ptr ) ) & CACHE_SPECIAL )

// #define ENCODE_SPECIAL_CACHE_NUM(num) ( ( void * ) ( ( ( ( uintptr_t ) ( num ) ) << 1 ) | CACHE_SPECIAL ) )

// #define DECODE_SPECIAL_CACHE_NUM(ptr) ( ( ( uintptr_t ) ( ptr ) ) >> 1 )

// #define ENCODE_SPECIAL_CACHE_PTR(ptr) ( ( void * ) ( ( ( uintptr_t ) ( ptr ) ) | CACHE_SPECIAL ) )

// #define DECODE_SPECIAL_CACHE_PTR(ptr) ( ( void * ) ( ( ( uintptr_t ) ( ptr ) ) & ~ CACHE_SPECIAL ) )

// #define SKIP_EXT_OPLINE(opline) do { while ( UNEXPECTED ( ( opline ) -> opcode >= ZEND_EXT_STMT && ( opline ) -> opcode <= ZEND_TICKS ) ) { ( opline ) -- ; } } while ( 0 )

// #define ZEND_CLASS_HAS_TYPE_HINTS(ce) ( ( ce -> ce_flags & ZEND_ACC_HAS_TYPE_HINTS ) == ZEND_ACC_HAS_TYPE_HINTS )

// #define ZEND_REF_ADD_TYPE_SOURCE(ref,source) zend_ref_add_type_source ( & ZEND_REF_TYPE_SOURCES ( ref ) , source )

// #define ZEND_REF_DEL_TYPE_SOURCE(ref,source) zend_ref_del_type_source ( & ZEND_REF_TYPE_SOURCES ( ref ) , source )

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

// #define ZEND_INTENSIVE_DEBUGGING       0

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

// #define EXECUTE_DATA_D       zend_execute_data * execute_data

// #define EXECUTE_DATA_C       execute_data

// #define EXECUTE_DATA_DC       , EXECUTE_DATA_D

// #define EXECUTE_DATA_CC       , EXECUTE_DATA_C

// #define NO_EXECUTE_DATA_CC       , NULL

// #define OPLINE_D       const zend_op * opline

// #define OPLINE_C       opline

// #define OPLINE_DC       , OPLINE_D

// #define OPLINE_CC       , OPLINE_C

// #define _CONST_CODE       0

// #define _TMP_CODE       1

// #define _VAR_CODE       2

// #define _UNUSED_CODE       3

// #define _CV_CODE       4

type IncdecT func(*Zval) int

// #define get_zval_ptr(op_type,node,should_free,type) _get_zval_ptr ( op_type , node , should_free , type EXECUTE_DATA_CC OPLINE_CC )

// #define get_zval_ptr_deref(op_type,node,should_free,type) _get_zval_ptr_deref ( op_type , node , should_free , type EXECUTE_DATA_CC OPLINE_CC )

// #define get_zval_ptr_undef(op_type,node,should_free,type) _get_zval_ptr_undef ( op_type , node , should_free , type EXECUTE_DATA_CC OPLINE_CC )

// #define get_op_data_zval_ptr_r(op_type,node,should_free) _get_op_data_zval_ptr_r ( op_type , node , should_free EXECUTE_DATA_CC OPLINE_CC )

// #define get_op_data_zval_ptr_deref_r(op_type,node,should_free) _get_op_data_zval_ptr_deref_r ( op_type , node , should_free EXECUTE_DATA_CC OPLINE_CC )

// #define get_zval_ptr_ptr(op_type,node,should_free,type) _get_zval_ptr_ptr ( op_type , node , should_free , type EXECUTE_DATA_CC )

// #define get_zval_ptr_ptr_undef(op_type,node,should_free,type) _get_zval_ptr_ptr ( op_type , node , should_free , type EXECUTE_DATA_CC )

// #define get_obj_zval_ptr(op_type,node,should_free,type) _get_obj_zval_ptr ( op_type , node , should_free , type EXECUTE_DATA_CC OPLINE_CC )

// #define get_obj_zval_ptr_undef(op_type,node,should_free,type) _get_obj_zval_ptr_undef ( op_type , node , should_free , type EXECUTE_DATA_CC OPLINE_CC )

// #define get_obj_zval_ptr_ptr(op_type,node,should_free,type) _get_obj_zval_ptr_ptr ( op_type , node , should_free , type EXECUTE_DATA_CC )

// #define RETURN_VALUE_USED(opline) ( ( opline ) -> result_type != IS_UNUSED )

func ZifPass(execute_data *ZendExecuteData, return_value *Zval) {}

var ZendPassFunction ZendInternalFunction = ZendInternalFunction{1, {0, 0, 0}, 0, nil, nil, nil, 0, 0, nil, ZifPass, nil, {nil, nil, nil, nil}}

// #define FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op,result) do { zval * __container_to_free = ( free_op ) ; if ( UNEXPECTED ( __container_to_free ) && EXPECTED ( Z_REFCOUNTED_P ( __container_to_free ) ) ) { zend_refcounted * __ref = Z_COUNTED_P ( __container_to_free ) ; if ( UNEXPECTED ( ! GC_DELREF ( __ref ) ) ) { zval * __zv = ( result ) ; if ( EXPECTED ( Z_TYPE_P ( __zv ) == IS_INDIRECT ) ) { ZVAL_COPY ( __zv , Z_INDIRECT_P ( __zv ) ) ; } rc_dtor_func ( __ref ) ; } } } while ( 0 )

// #define FREE_OP(should_free) if ( should_free ) { zval_ptr_dtor_nogc ( should_free ) ; }

// #define FREE_UNFETCHED_OP(type,var) if ( ( type ) & ( IS_TMP_VAR | IS_VAR ) ) { zval_ptr_dtor_nogc ( EX_VAR ( var ) ) ; }

// #define FREE_OP_VAR_PTR(should_free) if ( should_free ) { zval_ptr_dtor_nogc ( should_free ) ; }

// #define CV_DEF_OF(i) ( EX ( func ) -> op_array . vars [ i ] )

// #define ZEND_VM_STACK_PAGE_SLOTS       ( 16 * 1024 )

// #define ZEND_VM_STACK_PAGE_SIZE       ( ZEND_VM_STACK_PAGE_SLOTS * sizeof ( zval ) )

// #define ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size,page_size) ( ( ( size ) + ZEND_VM_STACK_HEADER_SLOTS * sizeof ( zval ) + ( ( page_size ) - 1 ) ) & ~ ( ( page_size ) - 1 ) )

func ZendVmStackNewPage(size int, prev ZendVmStack) ZendVmStack {
	var page ZendVmStack = ZendVmStack(_emalloc(size))
	page.SetTop((*Zval)(page) + ((g.SizeOf("struct _zend_vm_stack")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))
	page.SetEnd((*Zval)((*byte)(page + size)))
	page.SetPrev(prev)
	return page
}
func ZendVmStackInit() {
	EG.SetVmStackPageSize(16 * 1024 * g.SizeOf("zval"))
	EG.SetVmStack(ZendVmStackNewPage(16*1024*g.SizeOf("zval"), nil))
	EG.SetVmStackTop(EG.GetVmStack().GetTop())
	EG.SetVmStackEnd(EG.GetVmStack().GetEnd())
}
func ZendVmStackInitEx(page_size int) {
	/* page_size must be a power of 2 */

	assert(page_size > 0 && (page_size&page_size-1) == 0)
	EG.SetVmStackPageSize(page_size)
	EG.SetVmStack(ZendVmStackNewPage(page_size, nil))
	EG.SetVmStackTop(EG.GetVmStack().GetTop())
	EG.SetVmStackEnd(EG.GetVmStack().GetEnd())
}
func ZendVmStackDestroy() {
	var stack ZendVmStack = EG.GetVmStack()
	for stack != nil {
		var p ZendVmStack = stack.GetPrev()
		_efree(stack)
		stack = p
	}
}
func ZendVmStackExtend(size int) any {
	var stack ZendVmStack
	var ptr any
	stack = EG.GetVmStack()
	stack.SetTop(EG.GetVmStackTop())
	stack = ZendVmStackNewPage(g.CondF(size < EG.GetVmStackPageSize()-((g.SizeOf("struct _zend_vm_stack")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))*g.SizeOf("zval"), func() int { return EG.GetVmStackPageSize() }, func() int {
		return size + ((g.SizeOf("struct _zend_vm_stack")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))*g.SizeOf("zval") + (EG.GetVmStackPageSize()-1) & ^(EG.GetVmStackPageSize()-1)
	}), stack)
	EG.SetVmStack(stack)
	ptr = stack.GetTop()
	EG.SetVmStackTop(any((*byte)(ptr) + size))
	EG.SetVmStackEnd(stack.GetEnd())
	return ptr
}
func ZendGetCompiledVariableValue(execute_data *ZendExecuteData, var_ uint32) *Zval {
	return (*Zval)((*byte)(execute_data) + int(var_))
}
func _getZvalPtrTmp(var_ uint32, should_free *ZendFreeOp, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	*should_free = ret
	assert(ret.GetType() != 10)
	return ret
}
func _getZvalPtrVar(var_ uint32, should_free *ZendFreeOp, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	*should_free = ret
	return ret
}
func _getZvalPtrVarDeref(var_ uint32, should_free *ZendFreeOp, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	*should_free = ret
	if ret.GetType() == 10 {
		ret = &(*ret).value.GetRef().GetVal()
	}
	return ret
}
func ZvalUndefinedCv(var_ uint32, execute_data *ZendExecuteData) *Zval {
	if EG.GetException() == nil {
		var cv *ZendString = execute_data.GetFunc().GetOpArray().GetVars()[uint32((*Zval)((*byte)(nil)+int(var_))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0))))]
		ZendError(1<<3, "Undefined variable: %s", cv.GetVal())
	}
	return &EG.uninitialized_zval
}
func _zvalUndefinedOp1(execute_data *ZendExecuteData) *Zval {
	return ZvalUndefinedCv(execute_data.GetOpline().GetOp1().GetVar(), execute_data)
}
func _zvalUndefinedOp2(execute_data *ZendExecuteData) *Zval {
	return ZvalUndefinedCv(execute_data.GetOpline().GetOp2().GetVar(), execute_data)
}

// #define ZVAL_UNDEFINED_OP1() _zval_undefined_op1 ( EXECUTE_DATA_C )

// #define ZVAL_UNDEFINED_OP2() _zval_undefined_op2 ( EXECUTE_DATA_C )

func _getZvalCvLookup(ptr *Zval, var_ uint32, type_ int, execute_data *ZendExecuteData) *Zval {
	switch type_ {
	case 0:

	case 5:
		ptr = ZvalUndefinedCv(var_, execute_data)
		break
	case 3:
		ptr = &EG.uninitialized_zval
		break
	case 2:
		ZvalUndefinedCv(var_, execute_data)
	case 1:
		ptr.SetTypeInfo(1)
		break
	}
	return ptr
}
func _getZvalPtrCv(var_ uint32, type_ int, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		if type_ == 1 {
			ret.SetTypeInfo(1)
		} else {
			return _getZvalCvLookup(ret, var_, type_, execute_data)
		}
	}
	return ret
}
func _getZvalPtrCvDeref(var_ uint32, type_ int, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		if type_ == 1 {
			ret.SetTypeInfo(1)
			return ret
		} else {
			return _getZvalCvLookup(ret, var_, type_, execute_data)
		}
	}
	if ret.GetType() == 10 {
		ret = &(*ret).value.GetRef().GetVal()
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_R(var_ uint32, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		return ZvalUndefinedCv(var_, execute_data)
	}
	return ret
}
func _get_zval_ptr_cv_deref_BP_VAR_R(var_ uint32, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		return ZvalUndefinedCv(var_, execute_data)
	}
	if ret.GetType() == 10 {
		ret = &(*ret).value.GetRef().GetVal()
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_IS(var_ uint32, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	return ret
}
func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		ret.SetTypeInfo(1)
		ZvalUndefinedCv(var_, execute_data)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 0 {
		ret.SetTypeInfo(1)
	}
	return ret
}
func _getZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if (op_type & (1<<1 | 1<<2)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, execute_data)
	} else {
		*should_free = nil
		if op_type == 1<<0 {
			return (*Zval)((*byte)(opline) + int32(node).constant)
		} else if op_type == 1<<3 {
			return _getZvalPtrCv(node.GetVar(), type_, execute_data)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if (op_type & (1<<1 | 1<<2)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, execute_data)
	} else {
		*should_free = nil
		if op_type == 1<<0 {
			return (*Zval)((*byte)(opline+1) + int32(node).constant)
		} else if op_type == 1<<3 {
			return _get_zval_ptr_cv_BP_VAR_R(node.GetVar(), execute_data)
		} else {
			return nil
		}
	}
}
func _getZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if (op_type & (1<<1 | 1<<2)) != 0 {
		if op_type == 1<<1 {
			return _getZvalPtrTmp(node.GetVar(), should_free, execute_data)
		} else {
			assert(op_type == 1<<2)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, execute_data)
		}
	} else {
		*should_free = nil
		if op_type == 1<<0 {
			return (*Zval)((*byte)(opline) + int32(node).constant)
		} else if op_type == 1<<3 {
			return _getZvalPtrCvDeref(node.GetVar(), type_, execute_data)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if (op_type & (1<<1 | 1<<2)) != 0 {
		if op_type == 1<<1 {
			return _getZvalPtrTmp(node.GetVar(), should_free, execute_data)
		} else {
			assert(op_type == 1<<2)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, execute_data)
		}
	} else {
		*should_free = nil
		if op_type == 1<<0 {
			return (*Zval)((*byte)(opline+1) + int32(node).constant)
		} else if op_type == 1<<3 {
			return _get_zval_ptr_cv_deref_BP_VAR_R(node.GetVar(), execute_data)
		} else {
			return nil
		}
	}
}
func _getZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if (op_type & (1<<1 | 1<<2)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, execute_data)
	} else {
		*should_free = nil
		if op_type == 1<<0 {
			return (*Zval)((*byte)(opline) + int32(node).constant)
		} else if op_type == 1<<3 {
			return (*Zval)((*byte)(execute_data) + int(node.GetVar()))
		} else {
			return nil
		}
	}
}
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, execute_data *ZendExecuteData) *Zval {
	var ret *Zval = (*Zval)((*byte)(execute_data) + int(var_))
	if ret.GetType() == 13 {
		*should_free = nil
		ret = ret.GetValue().GetZv()
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData) *Zval {
	if op_type == 1<<3 {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, execute_data)
	} else {
		assert(op_type == 1<<2)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, execute_data)
	}
}
func _getObjZvalPtr(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if op_type == 0 {
		*should_free = nil
		return &(execute_data.GetThis())
	}
	return _getZvalPtr(op_type, op, should_free, type_, execute_data, opline)
}
func _getObjZvalPtrUndef(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData, opline *ZendOp) *Zval {
	if op_type == 0 {
		*should_free = nil
		return &(execute_data.GetThis())
	}
	return _getZvalPtrUndef(op_type, op, should_free, type_, execute_data, opline)
}
func _getObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, execute_data *ZendExecuteData) *Zval {
	if op_type == 0 {
		*should_free = nil
		return &(execute_data.GetThis())
	}
	return _getZvalPtrPtr(op_type, node, should_free, type_, execute_data)
}
func ZendAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval) {
	var ref *ZendReference
	if value_ptr.GetType() != 10 {
		var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
		ZendGcSetRefcount(&_ref.gc, 1)
		_ref.GetGc().SetTypeInfo(10)
		var _z1 *Zval = &_ref.val
		var _z2 *Zval = value_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		_ref.GetSources().SetPtr(nil)
		value_ptr.GetValue().SetRef(_ref)
		value_ptr.SetTypeInfo(10 | 1<<0<<8)
	} else if variable_ptr == value_ptr {
		return
	}
	ref = value_ptr.GetValue().GetRef()
	ZendGcAddref(&ref.gc)
	if variable_ptr.GetTypeFlags() != 0 {
		var garbage *ZendRefcounted = variable_ptr.GetValue().GetCounted()
		if ZendGcDelref(&garbage.gc) == 0 {
			var __z *Zval = variable_ptr
			__z.GetValue().SetRef(ref)
			__z.SetTypeInfo(10 | 1<<0<<8)
			RcDtorFunc(garbage)
			return
		} else {
			GcCheckPossibleRoot(garbage)
		}
	}
	var __z *Zval = variable_ptr
	__z.GetValue().SetRef(ref)
	__z.SetTypeInfo(10 | 1<<0<<8)
}
func ZendAssignToTypedPropertyReference(prop_info *ZendPropertyInfo, prop *Zval, value_ptr *Zval, execute_data *ZendExecuteData) *Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) == 0 {
		return &EG.uninitialized_zval
	}
	if prop.GetType() == 10 {
		ZendRefDelTypeSource(&(prop.GetValue().GetRef()).sources, prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZendRefAddTypeSource(&(prop.GetValue().GetRef()).sources, prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) *Zval {
	ZendError(1<<3, "Only variables should be assigned by reference")
	if EG.GetException() != nil {
		return &EG.uninitialized_zval
	}

	/* Use IS_TMP_VAR instead of IS_VAR to avoid ISREF check */

	if value_ptr.GetTypeFlags() != 0 {
		ZvalAddrefP(value_ptr)
	}
	return ZendAssignToVariable(variable_ptr, value_ptr, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
}
func ZendFormatType(type_ ZendType, part1 **byte, part2 **byte) {
	if (type_ & 0x1) != 0 {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if type_ > 0x3ff {
		if (type_ & 0x2) != 0 {
			*part2 = (*ZendClassEntry)(type_ & ^0x3).GetName().GetVal()
		} else {
			*part2 = (*ZendString)(type_ & ^0x3).GetVal()
		}
	} else {
		*part2 = ZendGetTypeByConst(type_ >> 2)
	}
}
func ZendThrowAutoInitInPropError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAutoInitInRefError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAccessUninitPropByRefError(prop *ZendPropertyInfo) {
	ZendThrowError(nil, "Cannot access uninitialized non-nullable property %s::$%s by reference", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()))
}

/* this should modify object only if it's empty */

func MakeRealObject(object *Zval, property *Zval, opline *ZendOp, execute_data *ZendExecuteData) *Zval {
	var obj *ZendObject
	var ref *Zval = nil
	if object.GetType() == 10 {
		ref = object
		object = &(*object).value.GetRef().GetVal()
	}
	if object.GetType() > 2 && (object.GetType() != 6 || object.GetValue().GetStr().GetLen() != 0) {
		if opline.GetOp1Type() != 1<<2 || object.GetType() != 15 {
			var tmp_property_name *ZendString
			var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
			if opline.GetOpcode() == 132 || opline.GetOpcode() == 133 || opline.GetOpcode() == 134 || opline.GetOpcode() == 135 {
				ZendError(1<<1, "Attempt to increment/decrement property '%s' of non-object", property_name.GetVal())
			} else if opline.GetOpcode() == 85 || opline.GetOpcode() == 88 || opline.GetOpcode() == 94 || opline.GetOpcode() == 32 {
				ZendError(1<<1, "Attempt to modify property '%s' of non-object", property_name.GetVal())
			} else {
				ZendError(1<<1, "Attempt to assign property '%s' of non-object", property_name.GetVal())
			}
			ZendTmpStringRelease(tmp_property_name)
		}
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		return nil
	}
	if ref != nil && ref.GetValue().GetRef().GetSources().GetPtr() != nil {
		if zend_verify_ref_stdClass_assignable(ref.GetValue().GetRef()) == 0 {
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return nil
		}
	}
	ZvalPtrDtorNogc(object)
	ObjectInit(object)
	ZvalAddrefP(object)
	obj = object.GetValue().GetObj()
	ZendError(1<<1, "Creating default object from empty value")
	if ZendGcRefcount(&obj.gc) == 1 {

		/* the enclosing container was deleted, obj is unreferenced */

		ZendObjectRelease(obj)
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		return nil
	}
	ZvalDelrefP(object)
	return object
}
func ZendVerifyTypeErrorCommon(zf *ZendFunction, arg_info *ZendArgInfo, ce *ZendClassEntry, value *Zval, fname **byte, fsep **byte, fclass **byte, need_msg **byte, need_kind **byte, need_or_null **byte, given_msg **byte, given_kind **byte) {
	var is_interface ZendBool = 0
	*fname = zf.GetFunctionName().GetVal()
	if zf.GetScope() != nil {
		*fsep = "::"
		*fclass = zf.GetScope().GetName().GetVal()
	} else {
		*fsep = ""
		*fclass = ""
	}
	if arg_info.GetType() > 0x3ff {
		if ce != nil {
			if (ce.GetCeFlags() & 1 << 0) != 0 {
				*need_msg = "implement interface "
				is_interface = 1
			} else {
				*need_msg = "be an instance of "
			}
			*need_kind = ce.GetName().GetVal()
		} else {

			/* We don't know whether it's a class or interface, assume it's a class */

			*need_msg = "be an instance of "
			*need_kind = (*ZendString)(arg_info.GetType() & ^0x3).GetVal()
		}
	} else {
		switch arg_info.GetType() >> 2 {
		case 8:
			*need_msg = "be an "
			*need_kind = "object"
			break
		case 17:
			*need_msg = "be callable"
			*need_kind = ""
			break
		case 18:
			*need_msg = "be iterable"
			*need_kind = ""
			break
		default:
			*need_msg = "be of the type "
			*need_kind = ZendGetTypeByConst(arg_info.GetType() >> 2)
			break
		}
	}
	if (arg_info.GetType() & 0x1) != 0 {
		if is_interface != 0 {
			*need_or_null = " or be null"
		} else {
			*need_or_null = " or null"
		}
	} else {
		*need_or_null = ""
	}
	if value != nil {
		if arg_info.GetType() > 0x3ff && value.GetType() == 8 {
			*given_msg = "instance of "
			*given_kind = value.GetValue().GetObj().GetCe().GetName().GetVal()
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
	var ptr *ZendExecuteData = EG.GetCurrentExecuteData().GetPrevExecuteData()
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	if EG.GetException() != nil {

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

		return

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

	}
	if value != nil {
		ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
		if zf.GetCommonType() == 2 {
			if ptr != nil && ptr.GetFunc() != nil && (ptr.GetFunc().GetCommonType()&1) == 0 {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno())
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
	if default_value.GetType() == 11 {
		var constant Zval
		var _z1 *Zval = &constant
		var _z2 *Zval = default_value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		if ZvalUpdateConstantEx(&constant, scope) != SUCCESS {
			return 0
		}
		if constant.GetType() == 1 {
			return 1
		}
		ZvalPtrDtorNogc(&constant)
	}
	return 0
}
func ZendVerifyWeakScalarTypeHint(type_hint ZendUchar, arg *Zval) ZendBool {
	switch type_hint {
	case 16:
		var dest ZendBool
		if ZendParseArgBoolWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		if dest != 0 {
			arg.SetTypeInfo(3)
		} else {
			arg.SetTypeInfo(2)
		}
		return 1
	case 4:
		var dest ZendLong
		if ZendParseArgLongWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		var __z *Zval = arg
		__z.GetValue().SetLval(dest)
		__z.SetTypeInfo(4)
		return 1
	case 5:
		var dest float64
		if ZendParseArgDoubleWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		var __z *Zval = arg
		__z.GetValue().SetDval(dest)
		__z.SetTypeInfo(5)
		return 1
	case 6:
		var dest *ZendString

		/* on success "arg" is converted to IS_STRING */

		return ZendParseArgStrWeak(arg, &dest)

	/* on success "arg" is converted to IS_STRING */

	default:
		return 0
	}
}
func ZendVerifyScalarTypeHint(type_hint ZendUchar, arg *Zval, strict ZendBool) ZendBool {
	if strict != 0 {

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

		if type_hint != 5 || arg.GetType() != 4 {
			return 0
		}

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	} else if arg.GetType() == 1 {

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

	if EG.GetException() != nil {
		return
	}

	// TODO Switch to a more standard error message?

	ZendFormatType(info.GetType(), &prop_type1, &prop_type2)
	void(prop_type1)
	if info.GetType() > 0x3ff {
		ZendTypeError("Typed property %s::$%s must be an instance of %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, g.Cond((info.GetType()&0x1) != 0, " or null", ""), g.CondF(property.GetType() == 8, func() []byte { return property.GetValue().GetObj().GetCe().GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	} else {
		ZendTypeError("Typed property %s::$%s must be %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, g.Cond((info.GetType()&0x1) != 0, " or null", ""), g.CondF(property.GetType() == 8, func() []byte { return property.GetValue().GetObj().GetCe().GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	}
}
func ZendResolveClassType(type_ *ZendType, self_ce *ZendClassEntry) ZendBool {
	var ce *ZendClassEntry
	var name *ZendString = (*ZendString)((*type_) & ^0x3)
	if name.GetLen() == g.SizeOf("\"self\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "self", g.SizeOf("\"self\"")-1) == 0 {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if (self_ce.GetCeFlags() & 1 << 1) != 0 {
			ZendThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", g.Cond(((*type_)&0x1) != 0, " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if name.GetLen() == g.SizeOf("\"parent\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "parent", g.SizeOf("\"parent\"")-1) == 0 {
		if !(self_ce.parent) {
			ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return 0
		}
		ce = self_ce.parent
	} else {
		ce = ZendLookupClassEx(name, nil, 0x80)
		if ce == nil {
			return 0
		}
	}
	ZendStringRelease(name)
	*type_ = uintptr_t(ce) | g.Cond(((*type_)&0x1) != 0, 0x3, 0x2)
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	assert(property.GetType() != 10)
	if info.GetType() > 0x3ff {
		if property.GetType() != 8 {
			return property.GetType() == 1 && (info.GetType()&0x1) != 0
		}
		if (info.GetType()&0x2) == 0 && ZendResolveClassType(&info.type_, info.GetCe()) == 0 {
			return 0
		}
		return InstanceofFunction(property.GetValue().GetObj().GetCe(), (*ZendClassEntry)(info.GetType() & ^0x3))
	}
	assert(info.GetType()>>2 != 17)
	if info.GetType()>>2 == property.GetType() {
		return 1
	} else if property.GetType() == 1 {
		return (info.GetType() & 0x1) != 0
	} else if info.GetType()>>2 == 16 && property.GetType() == 2 || property.GetType() == 3 {
		return 1
	} else if info.GetType()>>2 == 18 {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType()>>2, property, strict)
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
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *Zval, value *Zval, execute_data *ZendExecuteData) *Zval {
	var tmp Zval
	if value.GetType() == 10 {
		value = &(*value).value.GetRef().GetVal()
	}
	var _z1 *Zval = &tmp
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	if IZendVerifyPropertyType(info, &tmp, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) == 0 {
		ZvalPtrDtor(&tmp)
		return &EG.uninitialized_zval
	}
	return ZendAssignToVariable(property_val, &tmp, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
}
func ZendCheckType(type_ ZendType, arg *Zval, ce **ZendClassEntry, cache_slot *any, default_value *Zval, scope *ZendClassEntry, is_return_type ZendBool) ZendBool {
	var ref *ZendReference = nil
	if type_ <= 0x3 {
		return 1
	}
	if arg.GetType() == 10 {
		ref = arg.GetValue().GetRef()
		arg = &(*arg).value.GetRef().GetVal()
	}
	if type_ > 0x3ff {
		if *cache_slot {
			*ce = (*ZendClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass((*ZendString)(type_ & ^0x3), 4|0x80)
			if (*ce) == nil {
				return arg.GetType() == 1 && ((type_&0x1) != 0 || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if arg.GetType() == 8 {
			return InstanceofFunction(arg.GetValue().GetObj().GetCe(), *ce)
		}
		return arg.GetType() == 1 && ((type_&0x1) != 0 || default_value != nil && IsNullConstant(scope, default_value) != 0)
	} else if type_>>2 == arg.GetType() {
		return 1
	}
	if arg.GetType() == 1 && ((type_&0x1) != 0 || default_value != nil && IsNullConstant(scope, default_value) != 0) {

		/* Null passed to nullable type */

		return 1

		/* Null passed to nullable type */

	}
	if type_>>2 == 17 {
		return ZendIsCallable(arg, 1<<3, nil)
	} else if type_>>2 == 18 {
		return ZendIsIterable(arg)
	} else if type_>>2 == 16 && arg.GetType() == 2 || arg.GetType() == 3 {
		return 1
	} else if ref != nil && ref.GetSources().GetPtr() != nil {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(type_>>2, arg, g.CondF(is_return_type != 0, func() bool { return (EG.GetCurrentExecuteData().GetFunc().GetFnFlags() & 1 << 31) != 0 }, func() bool {
			return EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0
		}))
	}
}
func ZendVerifyArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	if arg_num <= zf.GetNumArgs() {
		cur_arg_info = &zf.common.arg_info[arg_num-1]
	} else if (zf.GetFnFlags() & 1 << 14) != 0 {
		cur_arg_info = &zf.common.arg_info[zf.GetNumArgs()]
	} else {
		return 1
	}
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyRecvArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = &zf.common.arg_info[arg_num-1]
	var ce *ZendClassEntry
	assert(arg_num <= zf.GetNumArgs())
	cur_arg_info = &zf.common.arg_info[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	assert(arg_num > zf.GetNumArgs())
	assert((zf.GetFnFlags() & 1 << 14) != 0)
	cur_arg_info = &zf.common.arg_info[zf.GetNumArgs()]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyInternalArgTypes(fbc *ZendFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = call.GetThis().GetNumArgs()
	var p *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
	var dummy_cache_slot any
	for i = 0; i < num_args; i++ {
		dummy_cache_slot = nil
		if ZendVerifyArgType(fbc, i+1, p, nil, &dummy_cache_slot) == 0 {
			EG.SetCurrentExecuteData(call.GetPrevExecuteData())
			return 0
		}
		p++
	}
	return 1
}
func ZendMissingArgError(execute_data *ZendExecuteData) {
	var ptr *ZendExecuteData = execute_data.GetPrevExecuteData()
	if ptr != nil && ptr.GetFunc() != nil && (ptr.GetFunc().GetCommonType()&1) == 0 {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", g.CondF1(execute_data.GetFunc().GetScope() != nil, func() []byte { return execute_data.GetFunc().GetScope().GetName().GetVal() }, ""), g.Cond(execute_data.GetFunc().GetScope() != nil, "::", ""), execute_data.GetFunc().GetFunctionName().GetVal(), execute_data.GetThis().GetNumArgs(), ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno(), g.Cond(execute_data.GetFunc().GetRequiredNumArgs() == execute_data.GetFunc().GetNumArgs(), "exactly", "at least"), execute_data.GetFunc().GetRequiredNumArgs())
	} else {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", g.CondF1(execute_data.GetFunc().GetScope() != nil, func() []byte { return execute_data.GetFunc().GetScope().GetName().GetVal() }, ""), g.Cond(execute_data.GetFunc().GetScope() != nil, "::", ""), execute_data.GetFunc().GetFunctionName().GetVal(), execute_data.GetThis().GetNumArgs(), g.Cond(execute_data.GetFunc().GetRequiredNumArgs() == execute_data.GetFunc().GetNumArgs(), "exactly", "at least"), execute_data.GetFunc().GetRequiredNumArgs())
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
	if ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0 {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf *ZendFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType() > 0x3 && ret_info.GetType()>>2 != 19 {
		var ce *ZendClassEntry = nil
		if ret_info.GetType() > 0x3ff {
			if *cache_slot {
				ce = (*ZendClassEntry)(*cache_slot)
			} else {
				ce = ZendFetchClass((*ZendString)(ret_info.GetType() & ^0x3), 4|0x80)
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
func ZendIllegalOffset() { ZendError(1<<1, "Illegal offset type") }
func ZendAssignToObjectDim(object *Zval, dim *Zval, value *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	object.GetValue().GetObj().GetHandlers().GetWriteDimension()(object, dim, value)
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
}
func ZendBinaryOp(ret *Zval, op1 *Zval, op2 *Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-1](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *Zval, property *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var free_op_data1 ZendFreeOp
	var value *Zval
	var z *Zval
	var rv Zval
	var res Zval
	value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1, execute_data, opline)
	if g.Assign(&z, object.GetValue().GetObj().GetHandlers().GetReadDimension()(object, property, 0, &rv)) != nil {
		if z.GetType() == 8 && z.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var rv2 Zval
			var value *Zval = z.GetValue().GetObj().GetHandlers().GetGet()(z, &rv2)
			if z == &rv {
				ZvalPtrDtor(&rv)
			}
			var _z1 *Zval = z
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
		if ZendBinaryOp(&res, z, value, opline) == SUCCESS {
			object.GetValue().GetObj().GetHandlers().GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = &res
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
		ZvalPtrDtor(&res)
	} else {
		ZendUseObjectAsArray()
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
	}
	if free_op_data1 != nil {
		ZvalPtrDtorNogc(free_op_data1)
	}
}
func ZendBinaryAssignOpTypedRef(ref *ZendReference, value *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == 8 && ref.GetVal().GetType() == 6 {
		ConcatFunction(&ref.val, &ref.val, value)
		assert(ref.GetVal().GetType() == 6 && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, &ref.val, value, opline)
	if ZendVerifyRefAssignableZval(ref, &z_copy, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) != 0 {
		ZvalPtrDtor(&ref.val)
		var _z1 *Zval = &ref.val
		var _z2 *Zval = &z_copy
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *Zval, value *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == 8 && zptr.GetType() == 6 {
		ConcatFunction(zptr, zptr, value)
		assert(zptr.GetType() == 6 && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, opline)
	if ZendVerifyPropertyType(prop_info, &z_copy, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) != 0 {
		ZvalPtrDtor(zptr)
		var _z1 *Zval = zptr
		var _z2 *Zval = &z_copy
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *Zval, type_ int, execute_data *ZendExecuteData) ZendLong {
	var offset ZendLong
try_again:
	if dim.GetType() != 4 {
		switch dim.GetType() {
		case 6:
			if 4 == IsNumericString(dim.GetValue().GetStr().GetVal(), dim.GetValue().GetStr().GetLen(), nil, nil, -1) {
				break
			}
			if type_ != 5 {
				ZendError(1<<1, "Illegal string offset '%s'", dim.GetValue().GetStr().GetVal())
			}
			break
		case 0:
			_zvalUndefinedOp2(execute_data)
		case 5:

		case 1:

		case 2:

		case 3:
			ZendError(1<<3, "String offset cast occurred")
			break
		case 10:
			dim = &(*dim).value.GetRef().GetVal()
			goto try_again
		default:
			ZendIllegalOffset()
			break
		}
		offset = ZvalGetLongFunc(dim)
	} else {
		offset = dim.GetValue().GetLval()
	}
	return offset
}
func ZendWrongStringOffset(execute_data *ZendExecuteData) {
	var msg *byte = nil
	var opline *ZendOp = execute_data.GetOpline()
	var end *ZendOp
	var var_ uint32
	if EG.GetException() != nil {
		return
	}
	switch opline.GetOpcode() {
	case 26:

	case 27:

	case 28:

	case 29:
		msg = "Cannot use assign-op operators with string offsets"
		break
	case 84:

	case 87:

	case 93:

	case 96:

	case 155:

		/* TODO: Encode the "reason" into opline->extended_value??? */

		var_ = opline.GetResult().GetVar()
		opline++
		end = EG.GetCurrentExecuteData().GetFunc().GetOpArray().GetOpcodes() + EG.GetCurrentExecuteData().GetFunc().GetOpArray().GetLast()
		for opline < end {
			if opline.GetOp1Type() == 1<<2 && opline.GetOp1().GetVar() == var_ {
				switch opline.GetOpcode() {
				case 85:

				case 88:

				case 94:

				case 97:

				case 24:

				case 28:

				case 32:
					msg = "Cannot use string offset as an object"
					break
				case 84:

				case 87:

				case 93:

				case 96:

				case 155:

				case 23:

				case 27:
					msg = "Cannot use string offset as an array"
					break
				case 29:

				case 26:
					msg = "Cannot use assign-op operators with string offsets"
					break
				case 132:

				case 133:

				case 134:

				case 135:

				case 34:

				case 35:

				case 36:

				case 37:
					msg = "Cannot increment/decrement string offsets"
					break
				case 30:

				case 72:

				case 71:

				case 140:
					msg = "Cannot create references to/from string offsets"
					break
				case 111:

				case 124:
					msg = "Cannot return string offsets by reference"
					break
				case 75:

				case 76:
					msg = "Cannot unset string offsets"
					break
				case 160:
					msg = "Cannot yield string offsets by reference"
					break
				case 67:

				case 66:

				case 185:
					msg = "Only variables can be passed by reference"
					break
				case 125:
					msg = "Cannot iterate on string offsets by reference"
					break
				default:
					break
				}
				break
			}
			if opline.GetOp2Type() == 1<<2 && opline.GetOp2().GetVar() == var_ {
				assert(opline.GetOpcode() == 30)
				msg = "Cannot create references to/from string offsets"
				break
			}
			opline++
		}
		break
	default:
		break
	}
	assert(msg != nil)
	ZendThrowError(nil, "%s", msg)
}
func ZendWrongPropertyRead(property *Zval) {
	var tmp_property_name *ZendString
	var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
	ZendError(1<<3, "Trying to get property '%s' of non-object", property_name.GetVal())
	ZendTmpStringRelease(tmp_property_name)
}
func ZendDeprecatedFunction(fbc *ZendFunction) {
	ZendError(1<<13, "Function %s%s%s() is deprecated", g.CondF1(fbc.GetScope() != nil, func() []byte { return fbc.GetScope().GetName().GetVal() }, ""), g.Cond(fbc.GetScope() != nil, "::", ""), fbc.GetFunctionName().GetVal())
}
func ZendAbstractMethod(fbc *ZendFunction) {
	ZendThrowError(nil, "Cannot call abstract method %s::%s()", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
}
func ZendAssignToStringOffset(str *Zval, dim *Zval, value *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var c ZendUchar
	var string_len int
	var offset ZendLong
	offset = ZendCheckStringOffset(dim, 1, execute_data)
	if EG.GetException() != nil {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return
	}
	if offset < -(zend_long(str.GetValue().GetStr()).len_) {

		/* Error on negative offset */

		ZendError(1<<1, "Illegal string offset:  "+"%"+"lld", offset)
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		return
	}
	if value.GetType() != 6 {

		/* Convert to string, just the time to pick the 1st byte */

		var tmp *ZendString = ZvalTryGetStringFunc(value)
		if tmp == nil {
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return
		}
		string_len = tmp.GetLen()
		c = zend_uchar(tmp).val[0]
		ZendStringReleaseEx(tmp, 0)
	} else {
		string_len = value.GetValue().GetStr().GetLen()
		c = zend_uchar(value.GetValue().GetStr()).val[0]
	}
	if string_len == 0 {

		/* Error on empty input string */

		ZendError(1<<1, "Cannot assign an empty string to a string offset")
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		return
	}
	if offset < 0 {
		offset += zend_long(str.GetValue().GetStr()).len_
	}
	if int(offset >= str.GetValue().GetStr().GetLen()) != 0 {

		/* Extend string if needed */

		var old_len ZendLong = str.GetValue().GetStr().GetLen()
		var __z *Zval = str
		var __s *ZendString = ZendStringExtend(str.GetValue().GetStr(), offset+1, 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		memset(str.GetValue().GetStr().GetVal()+old_len, ' ', offset-old_len)
		str.GetValue().GetStr().GetVal()[offset+1] = 0
	} else if str.GetTypeFlags() == 0 {
		var __z *Zval = str
		var __s *ZendString = ZendStringInit(str.GetValue().GetStr().GetVal(), str.GetValue().GetStr().GetLen(), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else if ZvalRefcountP(str) > 1 {
		ZvalDelrefP(str)
		var __z *Zval = str
		var __s *ZendString = ZendStringInit(str.GetValue().GetStr().GetVal(), str.GetValue().GetStr().GetLen(), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else {
		ZendStringForgetHashVal(str.GetValue().GetStr())
	}
	str.GetValue().GetStr().GetVal()[offset] = c
	if opline.GetResultType() != 0 {

		/* Return the new character */

		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = ZendOneCharString[c]
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)

		/* Return the new character */

	}
}
func ZendGetPropNotAcceptingDouble(ref *ZendReference) *ZendPropertyInfo {
	var prop *ZendPropertyInfo
	var _source_list *ZendPropertyInfoSourceList = &ref.sources
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if (_source_list.GetList() & 0x1) != 0 {
			_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = &_source_list.ptr
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			if prop.GetType()>>2 != 5 {
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

	assert(error_prop != nil)
	if (opline.GetOpcode() & 1) == 0 {
		ZendTypeError("Cannot increment a reference held by property %s::$%s of type %sint past its maximal value", error_prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(error_prop.GetName()), g.Cond((error_prop.GetType()&0x1) != 0, "?", ""))
		return INT64_MAX
	} else {
		ZendTypeError("Cannot decrement a reference held by property %s::$%s of type %sint past its minimal value", error_prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(error_prop.GetName()), g.Cond((error_prop.GetType()&0x1) != 0, "?", ""))
		return INT64_MIN
	}
}
func ZendThrowIncdecPropError(prop *ZendPropertyInfo, opline *ZendOp) ZendLong {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	if (opline.GetOpcode() & 1) == 0 {
		ZendTypeError("Cannot increment property %s::$%s of type %s%s past its maximal value", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return INT64_MAX
	} else {
		ZendTypeError("Cannot decrement property %s::$%s of type %s%s past its minimal value", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return INT64_MIN
	}
}
func ZendIncdecTypedRef(ref *ZendReference, copy *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var tmp Zval
	var var_ptr *Zval = &ref.val
	if copy == nil {
		copy = &tmp
	}
	var _z1 *Zval = copy
	var _z2 *Zval = var_ptr
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	if (opline.GetOpcode() & 1) == 0 {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if var_ptr.GetType() == 5 && copy.GetType() == 4 {
		var val ZendLong = ZendThrowIncdecRefError(ref, opline)
		var __z *Zval = var_ptr
		__z.GetValue().SetLval(val)
		__z.SetTypeInfo(4)
	} else if ZendVerifyRefAssignableZval(ref, var_ptr, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) == 0 {
		ZvalPtrDtor(var_ptr)
		var _z1 *Zval = var_ptr
		var _z2 *Zval = copy
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		copy.SetTypeInfo(0)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendIncdecTypedProp(prop_info *ZendPropertyInfo, var_ptr *Zval, copy *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var tmp Zval
	if copy == nil {
		copy = &tmp
	}
	var _z1 *Zval = copy
	var _z2 *Zval = var_ptr
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	if (opline.GetOpcode() & 1) == 0 {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if var_ptr.GetType() == 5 && copy.GetType() == 4 {
		var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
		var __z *Zval = var_ptr
		__z.GetValue().SetLval(val)
		__z.SetTypeInfo(4)
	} else if ZendVerifyPropertyType(prop_info, var_ptr, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0) == 0 {
		ZvalPtrDtor(var_ptr)
		var _z1 *Zval = var_ptr
		var _z2 *Zval = copy
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		copy.SetTypeInfo(0)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendPreIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, execute_data *ZendExecuteData) {
	if prop.GetType() == 4 {
		if (opline.GetOpcode() & 1) == 0 {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if prop.GetType() != 4 && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
			var __z *Zval = prop
			__z.GetValue().SetLval(val)
			__z.SetTypeInfo(4)
		}
	} else {
		for {
			if prop.GetType() == 10 {
				var ref *ZendReference = prop.GetValue().GetRef()
				prop = &(*prop).value.GetRef().GetVal()
				if ref.GetSources().GetPtr() != nil {
					ZendIncdecTypedRef(ref, nil, opline, execute_data)
					break
				}
			}
			if prop_info != nil {
				ZendIncdecTypedProp(prop_info, prop, nil, opline, execute_data)
			} else if (opline.GetOpcode() & 1) == 0 {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
			break
		}
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = prop
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
}
func ZendPostIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, execute_data *ZendExecuteData) {
	if prop.GetType() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(prop.GetValue().GetLval())
		__z.SetTypeInfo(4)
		if (opline.GetOpcode() & 1) == 0 {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if prop.GetType() != 4 && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
			var __z *Zval = prop
			__z.GetValue().SetLval(val)
			__z.SetTypeInfo(4)
		}
	} else {
		if prop.GetType() == 10 {
			var ref *ZendReference = prop.GetValue().GetRef()
			prop = &(*prop).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendIncdecTypedRef(ref, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), opline, execute_data)
				return
			}
		}
		if prop_info != nil {
			ZendIncdecTypedProp(prop_info, prop, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), opline, execute_data)
		} else {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = prop
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
			if (opline.GetOpcode() & 1) == 0 {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
		}
	}
}
func ZendPostIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, execute_data *ZendExecuteData) {
	var rv Zval
	var obj Zval
	var z *Zval
	var z_copy Zval
	var __z *Zval = &obj
	__z.GetValue().SetObj(object.GetValue().GetObj())
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ZvalAddrefP(&obj)
	z = obj.GetValue().GetObj().GetHandlers().GetReadProperty()(&obj, property, 0, cache_slot, &rv)
	if EG.GetException() != nil {
		ZendObjectRelease(obj.GetValue().GetObj())
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return
	}
	if z.GetType() == 8 && z.GetValue().GetObj().GetHandlers().GetGet() != nil {
		var rv2 Zval
		var value *Zval = z.GetValue().GetObj().GetHandlers().GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		var _z1 *Zval = z
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	var _z3 *Zval = z
	if (_z3.GetTypeInfo() & 0xff00) != 0 {
		if (_z3.GetTypeInfo() & 0xff) == 10 {
			_z3 = &(*_z3).value.GetRef().GetVal()
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(_z3)
			}
		} else {
			ZvalAddrefP(_z3)
		}
	}
	var _z1 *Zval = &z_copy
	var _z2 *Zval = _z3
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = &z_copy
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	if (opline.GetOpcode() & 1) == 0 {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	obj.GetValue().GetObj().GetHandlers().GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	ZendObjectRelease(obj.GetValue().GetObj())
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendPreIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, execute_data *ZendExecuteData) {
	var rv Zval
	var z *Zval
	var obj Zval
	var z_copy Zval
	var __z *Zval = &obj
	__z.GetValue().SetObj(object.GetValue().GetObj())
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ZvalAddrefP(&obj)
	z = obj.GetValue().GetObj().GetHandlers().GetReadProperty()(&obj, property, 0, cache_slot, &rv)
	if EG.GetException() != nil {
		ZendObjectRelease(obj.GetValue().GetObj())
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		return
	}
	if z.GetType() == 8 && z.GetValue().GetObj().GetHandlers().GetGet() != nil {
		var rv2 Zval
		var value *Zval = z.GetValue().GetObj().GetHandlers().GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		var _z1 *Zval = z
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	var _z3 *Zval = z
	if (_z3.GetTypeInfo() & 0xff00) != 0 {
		if (_z3.GetTypeInfo() & 0xff) == 10 {
			_z3 = &(*_z3).value.GetRef().GetVal()
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(_z3)
			}
		} else {
			ZvalAddrefP(_z3)
		}
	}
	var _z1 *Zval = &z_copy
	var _z2 *Zval = _z3
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (opline.GetOpcode() & 1) == 0 {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = &z_copy
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	obj.GetValue().GetObj().GetHandlers().GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	ZendObjectRelease(obj.GetValue().GetObj())
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendAssignOpOverloadedProperty(object *Zval, property *Zval, cache_slot *any, value *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var z *Zval
	var rv Zval
	var obj Zval
	var res Zval
	var __z *Zval = &obj
	__z.GetValue().SetObj(object.GetValue().GetObj())
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ZvalAddrefP(&obj)
	z = obj.GetValue().GetObj().GetHandlers().GetReadProperty()(&obj, property, 0, cache_slot, &rv)
	if EG.GetException() != nil {
		ZendObjectRelease(obj.GetValue().GetObj())
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return
	}
	if z.GetType() == 8 && z.GetValue().GetObj().GetHandlers().GetGet() != nil {
		var rv2 Zval
		var value *Zval = z.GetValue().GetObj().GetHandlers().GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		var _z1 *Zval = z
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	if ZendBinaryOp(&res, z, value, opline) == SUCCESS {
		obj.GetValue().GetObj().GetHandlers().GetWriteProperty()(&obj, property, &res, cache_slot)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = &res
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtor(z)
	ZvalPtrDtor(&res)
	ZendObjectRelease(obj.GetValue().GetObj())
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
func ZendGetTargetSymbolTable(fetch_type int, execute_data *ZendExecuteData) *HashTable {
	var ht *HashTable
	if (fetch_type & (1<<3 | 1<<1)) != 0 {
		ht = &EG.symbol_table
	} else {
		assert((fetch_type & 1 << 2) != 0)
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) == 0 {
			ZendRebuildSymbolTable()
		}
		ht = execute_data.GetSymbolTable()
	}
	return ht
}
func ZendUndefinedOffset(lval ZendLong) {
	ZendError(1<<3, "Undefined offset: "+"%"+"lld", lval)
}
func ZendUndefinedIndex(offset *ZendString) {
	ZendError(1<<3, "Undefined index: %s", offset.GetVal())
}
func ZendUndefinedOffsetWrite(ht *HashTable, lval ZendLong) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&ht.gc)
	}
	ZendUndefinedOffset(lval)
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if EG.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedIndexWrite(ht *HashTable, offset *ZendString) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&ht.gc)
	}
	ZendUndefinedIndex(offset)
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if EG.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedMethod(ce *ZendClassEntry, method *ZendString) {
	ZendThrowError(nil, "Call to undefined method %s::%s()", ce.GetName().GetVal(), method.GetVal())
}
func ZendInvalidMethodCall(object *Zval, function_name *Zval) {
	ZendThrowError(nil, "Call to a member function %s() on %s", function_name.GetValue().GetStr().GetVal(), ZendGetTypeByConst(object.GetType()))
}
func ZendNonStaticMethodCall(fbc *ZendFunction) {
	if (fbc.GetFnFlags() & 1 << 17) != 0 {
		ZendError(1<<13, "Non-static method %s::%s() should not be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	} else {
		ZendThrowError(ZendCeError, "Non-static method %s::%s() cannot be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	}
}
func ZendParamMustBeRef(func_ *ZendFunction, arg_num uint32) {
	ZendError(1<<1, "Parameter %d to %s%s%s() expected to be a reference, value given", arg_num, g.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), g.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
}
func ZendUseScalarAsArray() {
	ZendError(1<<1, "Cannot use a scalar value as an array")
}
func ZendCannotAddElement() {
	ZendError(1<<1, "Cannot add element to the array as the next element is already occupied")
}
func ZendUseResourceAsOffset(dim *Zval) {
	ZendError(1<<3, "Resource ID#%d used as offset, casting to integer (%d)", dim.GetValue().GetRes().GetHandle(), dim.GetValue().GetRes().GetHandle())
}
func ZendUseNewElementForString() {
	ZendThrowError(nil, "[] operator not supported for strings")
}
func ZendBinaryAssignOpDimSlow(container *Zval, dim *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	if container.GetType() == 6 {
		if opline.GetOp2Type() == 0 {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, 2, execute_data)
			ZendWrongStringOffset(execute_data)
		}
	} else if container.GetType() != 15 {
		ZendUseScalarAsArray()
	}
}
func SlowIndexConvert(ht *HashTable, dim *Zval, value *ZendValue, execute_data *ZendExecuteData) ZendUchar {
	switch dim.GetType() {
	case 0:

		/* The array may be destroyed while throwing the notice.
		 * Temporarily increase the refcount to detect this situation. */

		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			ZendGcAddref(&ht.gc)
		}
		_zvalUndefinedOp2(execute_data)
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
			ZendArrayDestroy(ht)
			return 1
		}
		if EG.GetException() != nil {
			return 1
		}
	case 1:
		value.SetStr(ZendEmptyString)
		return 6
	case 5:
		value.SetLval(ZendDvalToLval(dim.GetValue().GetDval()))
		return 4
	case 9:
		ZendUseResourceAsOffset(dim)
		value.SetLval(dim.GetValue().GetRes().GetHandle())
		return 4
	case 2:
		value.SetLval(0)
		return 4
	case 3:
		value.SetLval(1)
		return 4
	default:
		ZendIllegalOffset()
		return 1
	}
}
func ZendFetchDimensionAddressInner(ht *HashTable, dim *Zval, dim_type int, type_ int, execute_data *ZendExecuteData) *Zval {
	var retval *Zval = nil
	var offset_key *ZendString
	var hval ZendUlong
try_again:
	if dim.GetType() == 4 {
		hval = dim.GetValue().GetLval()
	num_index:
		if (ht.GetUFlags() & 1 << 2) != 0 {
			if zend_ulong(hval) < zend_ulong(ht).nNumUsed {
				retval = &ht.arData[hval].GetVal()
				if retval.GetType() == 0 {
					goto num_undef
				}
			} else {
				goto num_undef
			}
		} else {
			retval = _zendHashIndexFind(ht, hval)
			if retval == nil {
				goto num_undef
			}
		}
		return retval
	num_undef:
		switch type_ {
		case 0:
			ZendUndefinedOffset(hval)
		case 5:

		case 3:
			retval = &EG.uninitialized_zval
			break
		case 2:
			if ZendUndefinedOffsetWrite(ht, hval) == FAILURE {
				return nil
			}
		case 1:
			retval = ZendHashIndexAddNew(ht, hval, &EG.uninitialized_zval)
			break
		}
	} else if dim.GetType() == 6 {
		offset_key = dim.GetValue().GetStr()
		if _zendHandleNumericStr(offset_key.GetVal(), offset_key.GetLen(), &hval) != 0 {
			goto num_index
		}
	str_index:
		retval = ZendHashFindEx(ht, offset_key, 0)
		if retval != nil {

			/* support for $GLOBALS[...] */

			if retval.GetType() == 13 {
				retval = retval.GetValue().GetZv()
				if retval.GetType() == 0 {
					switch type_ {
					case 0:
						ZendUndefinedIndex(offset_key)
					case 5:

					case 3:
						retval = &EG.uninitialized_zval
						break
					case 2:
						if ZendUndefinedIndexWrite(ht, offset_key) != 0 {
							return nil
						}
					case 1:
						retval.SetTypeInfo(1)
						break
					}
				}
			}

			/* support for $GLOBALS[...] */

		} else {
			switch type_ {
			case 0:
				ZendUndefinedIndex(offset_key)
			case 5:

			case 3:
				retval = &EG.uninitialized_zval
				break
			case 2:

				/* Key may be released while throwing the undefined index warning. */

				ZendStringAddref(offset_key)
				if ZendUndefinedIndexWrite(ht, offset_key) == FAILURE {
					ZendStringRelease(offset_key)
					return nil
				}
				retval = ZendHashAddNew(ht, offset_key, &EG.uninitialized_zval)
				ZendStringRelease(offset_key)
				break
			case 1:
				retval = ZendHashAddNew(ht, offset_key, &EG.uninitialized_zval)
				break
			}
		}
	} else if dim.GetType() == 10 {
		dim = &(*dim).value.GetRef().GetVal()
		goto try_again
	} else {
		var val ZendValue
		var t ZendUchar = SlowIndexConvert(ht, dim, &val, execute_data)
		if t == 6 {
			offset_key = val.GetStr()
			goto str_index
		} else if t == 4 {
			hval = val.GetLval()
			goto num_index
		} else {
			if type_ == 1 || type_ == 2 {
				retval = nil
			} else {
				retval = &EG.uninitialized_zval
			}
		}
	}
	return retval
}
func zend_fetch_dimension_address_inner_W(ht *HashTable, dim *Zval, execute_data *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, 1<<1, 1, execute_data)
}
func zend_fetch_dimension_address_inner_W_CONST(ht *HashTable, dim *Zval, execute_data *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, 1<<0, 1, execute_data)
}
func zend_fetch_dimension_address_inner_RW(ht *HashTable, dim *Zval, execute_data *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, 1<<1, 2, execute_data)
}
func zend_fetch_dimension_address_inner_RW_CONST(ht *HashTable, dim *Zval, execute_data *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, 1<<0, 2, execute_data)
}
func ZendFetchDimensionAddress(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, execute_data *ZendExecuteData) {
	var retval *Zval
	if container.GetType() == 7 {
	try_array:
		var _zv *Zval = container
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
	fetch_from_array:
		if dim == nil {
			retval = ZendHashNextIndexInsert(container.GetValue().GetArr(), &EG.uninitialized_zval)
			if retval == nil {
				ZendCannotAddElement()
				result.SetTypeInfo(15)
				return
			}
		} else {
			retval = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, dim_type, type_, execute_data)
			if retval == nil {
				result.SetTypeInfo(15)
				return
			}
		}
		result.GetValue().SetZv(retval)
		result.SetTypeInfo(13)
		return
	} else if container.GetType() == 10 {
		var ref *ZendReference = container.GetValue().GetRef()
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto try_array
		} else if container.GetType() <= 2 {
			if type_ != 5 {
				if ref.GetSources().GetPtr() != nil {
					if ZendVerifyRefArrayAssignable(ref) == 0 {
						result.SetTypeInfo(15)
						return
					}
				}
				var __arr *ZendArray = _zendNewArray(0)
				var __z *Zval = container
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto fetch_from_array
			} else {
				goto return_null
			}
		}
	}
	if container.GetType() == 6 {
		if dim == nil {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, type_, execute_data)
			ZendWrongStringOffset(execute_data)
		}
		result.SetTypeInfo(15)
	} else if container.GetType() == 8 {
		if dim != nil && dim.GetType() == 0 {
			dim = _zvalUndefinedOp2(execute_data)
		}
		if dim_type == 1<<0 && dim.GetU2Extra() == 1 {
			dim++
		}
		retval = container.GetValue().GetObj().GetHandlers().GetReadDimension()(container, dim, type_, result)
		if retval == &EG.uninitialized_zval {
			var ce *ZendClassEntry = container.GetValue().GetObj().GetCe()
			result.SetTypeInfo(1)
			ZendError(1<<3, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
		} else if retval != nil && retval.GetType() != 0 {
			if retval.GetType() != 10 {
				if result != retval {
					var _z1 *Zval = result
					var _z2 *Zval = retval
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
					retval = result
				}
				if retval.GetType() != 8 {
					var ce *ZendClassEntry = container.GetValue().GetObj().GetCe()
					ZendError(1<<3, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
				}
			} else if ZvalRefcountP(retval) == 1 {
				var _z *Zval = retval
				var ref *ZendReference
				assert(_z.GetType() == 10)
				ref = _z.GetValue().GetRef()
				var _z1 *Zval = _z
				var _z2 *Zval = &ref.val
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_efree(ref)
			}
			if result != retval {
				result.GetValue().SetZv(retval)
				result.SetTypeInfo(13)
			}
		} else {
			result.SetTypeInfo(15)
		}
	} else {
		if container.GetType() <= 2 {
			if type_ != 1 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if type_ != 5 {
				var __arr *ZendArray = _zendNewArray(0)
				var __z *Zval = container
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto fetch_from_array
			} else {
			return_null:

				/* for read-mode only */

				if dim != nil && dim.GetType() == 0 {
					_zvalUndefinedOp2(execute_data)
				}
				result.SetTypeInfo(1)
			}
		} else if container.GetType() == 15 {
			result.SetTypeInfo(15)
		} else {
			if type_ == 5 {
				ZendError(1<<1, "Cannot unset offset in a non-array variable")
				result.SetTypeInfo(1)
			} else {
				ZendUseScalarAsArray()
				result.SetTypeInfo(15)
			}
		}
	}
}
func zend_fetch_dimension_address_W(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, 1, execute_data)
}
func zend_fetch_dimension_address_RW(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, 2, execute_data)
}
func zend_fetch_dimension_address_UNSET(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, 5, execute_data)
}
func ZendFetchDimensionAddressRead(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, is_list int, slow int, execute_data *ZendExecuteData) {
	var retval *Zval
	if slow == 0 {
		if container.GetType() == 7 {
		try_array:
			retval = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, dim_type, type_, execute_data)
			var _z3 *Zval = retval
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = result
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			return
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto try_array
			}
		}
	}
	if is_list == 0 && container.GetType() == 6 {
		var offset ZendLong
	try_string_offset:
		if dim.GetType() != 4 {
			switch dim.GetType() {
			case 6:
				if 4 == IsNumericString(dim.GetValue().GetStr().GetVal(), dim.GetValue().GetStr().GetLen(), nil, nil, -1) {
					break
				}
				if type_ == 3 {
					result.SetTypeInfo(1)
					return
				}
				ZendError(1<<1, "Illegal string offset '%s'", dim.GetValue().GetStr().GetVal())
				break
			case 0:
				_zvalUndefinedOp2(execute_data)
			case 5:

			case 1:

			case 2:

			case 3:
				if type_ != 3 {
					ZendError(1<<3, "String offset cast occurred")
				}
				break
			case 10:
				dim = &(*dim).value.GetRef().GetVal()
				goto try_string_offset
			default:
				ZendIllegalOffset()
				break
			}
			offset = ZvalGetLongFunc(dim)
		} else {
			offset = dim.GetValue().GetLval()
		}
		if container.GetValue().GetStr().GetLen() < g.CondF(offset < 0, func() int { return -int(offset) }, func() int { return int(offset + 1) }) {
			if type_ != 3 {
				ZendError(1<<3, "Uninitialized string offset: "+"%"+"lld", offset)
				var __z *Zval = result
				var __s *ZendString = ZendEmptyString
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
			} else {
				result.SetTypeInfo(1)
			}
		} else {
			var c ZendUchar
			var real_offset ZendLong
			if offset < 0 {
				real_offset = zend_long(container.GetValue().GetStr()).len_ + offset
			} else {
				real_offset = offset
			}
			c = zend_uchar(container.GetValue().GetStr()).val[real_offset]
			var __z *Zval = result
			var __s *ZendString = ZendOneCharString[c]
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
	} else if container.GetType() == 8 {
		if dim.GetType() == 0 {
			dim = _zvalUndefinedOp2(execute_data)
		}
		if dim_type == 1<<0 && dim.GetU2Extra() == 1 {
			dim++
		}
		retval = container.GetValue().GetObj().GetHandlers().GetReadDimension()(container, dim, type_, result)
		assert(result != nil)
		if retval != nil {
			if result != retval {
				var _z3 *Zval = retval
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					if (_z3.GetTypeInfo() & 0xff) == 10 {
						_z3 = &(*_z3).value.GetRef().GetVal()
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(_z3)
						}
					} else {
						ZvalAddrefP(_z3)
					}
				}
				var _z1 *Zval = result
				var _z2 *Zval = _z3
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if retval.GetType() == 10 {
				ZendUnwrapReference(result)
			}
		} else {
			result.SetTypeInfo(1)
		}
	} else {
		if type_ != 3 && container.GetType() == 0 {
			container = _zvalUndefinedOp1(execute_data)
		}
		if dim.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		if is_list == 0 && type_ != 3 {
			ZendError(1<<3, "Trying to access array offset on value of type %s", ZendZvalTypeName(container))
		}
		result.SetTypeInfo(1)
	}
}
func zend_fetch_dimension_address_read_R(container *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, 0, 0, 0, execute_data)
}
func zend_fetch_dimension_address_read_R_slow(container *Zval, dim *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddressRead(result, container, dim, 1<<3, 0, 0, 1, execute_data)
}
func zend_fetch_dimension_address_read_IS(container *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, 3, 0, 0, execute_data)
}
func zend_fetch_dimension_address_LIST_r(container *Zval, dim *Zval, dim_type int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, 0, 1, 0, execute_data)
}
func ZendFetchDimensionConst(result *Zval, container *Zval, dim *Zval, type_ int) {
	ZendFetchDimensionAddressRead(result, container, dim, 1<<1, type_, 0, 0, nil)
}
func ZendFindArrayDimSlow(ht *HashTable, offset *Zval, execute_data *ZendExecuteData) *Zval {
	var hval ZendUlong
	if offset.GetType() == 5 {
		hval = ZendDvalToLval(offset.GetValue().GetDval())
	num_idx:
		return ZendHashIndexFind(ht, hval)
	} else if offset.GetType() == 1 {
	str_idx:
		return ZendHashFindExInd(ht, ZendEmptyString, 1)
	} else if offset.GetType() == 2 {
		hval = 0
		goto num_idx
	} else if offset.GetType() == 3 {
		hval = 1
		goto num_idx
	} else if offset.GetType() == 9 {
		hval = offset.GetValue().GetRes().GetHandle()
		goto num_idx
	} else if offset.GetType() == 0 {
		_zvalUndefinedOp2(execute_data)
		goto str_idx
	} else {
		ZendError(1<<1, "Illegal offset type in isset or empty")
		return nil
	}
}
func ZendIssetDimSlow(container *Zval, offset *Zval, execute_data *ZendExecuteData) int {
	if offset.GetType() == 0 {
		offset = _zvalUndefinedOp2(execute_data)
	}
	if container.GetType() == 8 {
		return container.GetValue().GetObj().GetHandlers().GetHasDimension()(container, offset, 0)
	} else if container.GetType() == 6 {
		var lval ZendLong
		if offset.GetType() == 4 {
			lval = offset.GetValue().GetLval()
		str_offset:
			if lval < 0 {
				lval += zend_long(container.GetValue().GetStr()).len_
			}
			if lval >= 0 && int(lval < container.GetValue().GetStr().GetLen()) != 0 {
				return 1
			} else {
				return 0
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			if offset.GetType() == 10 {
				offset = &(*offset).value.GetRef().GetVal()
			}

			/*}*/

			if offset.GetType() < 6 || offset.GetType() == 6 && 4 == IsNumericString(offset.GetValue().GetStr().GetVal(), offset.GetValue().GetStr().GetLen(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 0
		}
	} else {
		return 0
	}
}
func ZendIsemptyDimSlow(container *Zval, offset *Zval, execute_data *ZendExecuteData) int {
	if offset.GetType() == 0 {
		offset = _zvalUndefinedOp2(execute_data)
	}
	if container.GetType() == 8 {
		return !(container.GetValue().GetObj().GetHandlers().GetHasDimension()(container, offset, 1))
	} else if container.GetType() == 6 {
		var lval ZendLong
		if offset.GetType() == 4 {
			lval = offset.GetValue().GetLval()
		str_offset:
			if lval < 0 {
				lval += zend_long(container.GetValue().GetStr()).len_
			}
			if lval >= 0 && int(lval < container.GetValue().GetStr().GetLen()) != 0 {
				return container.GetValue().GetStr().GetVal()[lval] == '0'
			} else {
				return 1
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			if offset.GetType() == 10 {
				offset = &(*offset).value.GetRef().GetVal()
			}

			/*}*/

			if offset.GetType() < 6 || offset.GetType() == 6 && 4 == IsNumericString(offset.GetValue().GetStr().GetVal(), offset.GetValue().GetStr().GetLen(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 1
		}
	} else {
		return 1
	}
}
func ZendArrayKeyExistsFast(ht *HashTable, key *Zval, opline *ZendOp, execute_data *ZendExecuteData) uint32 {
	var str *ZendString
	var hval ZendUlong
try_again:
	if key.GetType() == 6 {
		str = key.GetValue().GetStr()
		if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
			goto num_key
		}
	str_key:
		if ZendHashFindInd(ht, str) != nil {
			return 3
		} else {
			return 2
		}
	} else if key.GetType() == 4 {
		hval = key.GetValue().GetLval()
	num_key:
		if ZendHashIndexFind(ht, hval) != nil {
			return 3
		} else {
			return 2
		}
	} else if key.GetType() == 10 {
		key = &(*key).value.GetRef().GetVal()
		goto try_again
	} else if key.GetType() <= 1 {
		if key.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		str = ZendEmptyString
		goto str_key
	} else {
		ZendError(1<<1, "array_key_exists(): The first argument should be either a string or an integer")
		return 2
	}
}
func ZendArrayKeyExistsSlow(subject *Zval, key *Zval, opline *ZendOp, execute_data *ZendExecuteData) uint32 {
	if subject.GetType() == 8 {
		ZendError(1<<13, "array_key_exists(): "+"Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
		var ht *HashTable = ZendGetPropertiesFor(subject, ZEND_PROP_PURPOSE_ARRAY_CAST)
		var result uint32 = ZendArrayKeyExistsFast(ht, key, opline, execute_data)
		if ht != nil && (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
			ZendArrayDestroy(ht)
		}
		return result
	} else {
		if key.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		if subject.GetTypeInfo() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		ZendInternalTypeError((execute_data.GetFunc().GetFnFlags()&1<<31) != 0, "array_key_exists() expects parameter 2 to be array, %s given", ZendGetTypeByConst(subject.GetType()))
		return 1
	}
}
func PromotesToArray(val *Zval) ZendBool {
	return val.GetType() <= 2 || val.GetType() == 10 && &(*val).value.GetRef().GetVal().u1.v.type_ <= 2
}
func PromotesToObject(val *Zval) ZendBool {
	if val.GetType() == 10 {
		val = &(*val).value.GetRef().GetVal()
	}
	return val.GetType() <= 2 || val.GetType() == 6 && val.GetValue().GetStr().GetLen() == 0
}
func CheckTypeArrayAssignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	return type_ > 0x3 && type_ <= 0x3ff && (type_>>2 == 7 || type_>>2 == 18)
}
func check_type_stdClass_assignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	if type_ > 0x3ff {
		if (type_ & 0x2) != 0 {
			return (*ZendClassEntry)(type_ & ^0x3) == ZendStandardClassDef
		} else {
			return (*ZendString)(type_ & ^0x3).GetLen() == g.SizeOf("\"stdclass\"")-1 && ZendBinaryStrcasecmp((*ZendString)(type_ & ^0x3).GetVal(), (*ZendString)(type_ & ^0x3).GetLen(), "stdclass", g.SizeOf("\"stdclass\"")-1) == 0
		}
	} else {
		return type_>>2 == 8
	}
}

/* Checks whether an array can be assigned to the reference. Returns conflicting property if
 * assignment is not possible, NULL otherwise. */

func ZendVerifyRefArrayAssignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	assert(ref.GetSources().GetPtr() != nil)
	var _source_list *ZendPropertyInfoSourceList = &ref.sources
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if (_source_list.GetList() & 0x1) != 0 {
			_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
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
	assert(ref.GetSources().GetPtr() != nil)
	var _source_list *ZendPropertyInfoSourceList = &ref.sources
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if (_source_list.GetList() & 0x1) != 0 {
			_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
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
	if (obj.GetCe().GetCeFlags() & 1 << 8) != 1<<8 {
		return nil
	}

	/* Not a declared property */

	if slot < obj.GetPropertiesTable() || slot >= obj.GetPropertiesTable()+obj.GetCe().GetDefaultPropertiesCount() {
		return nil
	}
	return ZendGetTypedPropertyInfoForSlot(obj, slot)
}
func ZendHandleFetchObjFlags(result *Zval, ptr *Zval, obj *ZendObject, prop_info *ZendPropertyInfo, flags uint32) ZendBool {
	switch flags {
	case 2:
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
					result.SetTypeInfo(15)
				}
				return 0
			}
		}
		break
	case 3:
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
					result.SetTypeInfo(15)
				}
				return 0
			}
		}
		break
	case 1:
		if ptr.GetType() != 10 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if ptr.GetType() == 0 {
				if (prop_info.GetType() & 0x1) == 0 {
					ZendThrowAccessUninitPropByRefError(prop_info)
					if result != nil {
						result.SetTypeInfo(15)
					}
					return 0
				}
				ptr.SetTypeInfo(1)
			}
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			ptr.GetValue().SetRef(_ref)
			ptr.SetTypeInfo(10 | 1<<0<<8)
			ZendRefAddTypeSource(&(ptr.GetValue().GetRef()).sources, prop_info)
		}
		break
	default:
		break
	}
	return 1
}
func ZendFetchPropertyAddress(result *Zval, container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, cache_slot *any, type_ int, flags uint32, init_undef ZendBool, opline *ZendOp, execute_data *ZendExecuteData) {
	var ptr *Zval
	if container_op_type != 0 && container.GetType() != 8 {
		for {
			if container.GetType() == 10 && &(*container).value.GetRef().GetVal().u1.v.type_ == 8 {
				container = &(*container).value.GetRef().GetVal()
				break
			}
			if container_op_type == 1<<3 && type_ != 1 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}

			/* this should modify object only if it's empty */

			if type_ == 5 {
				result.SetTypeInfo(1)
				return
			}
			container = MakeRealObject(container, prop_ptr, opline, execute_data)
			if container == nil {
				result.SetTypeInfo(15)
				return
			}
			break
		}
	}
	if prop_op_type == 1<<0 && container.GetValue().GetObj().GetCe() == cache_slot[0] {
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = container.GetValue().GetObj()
		if intptr_t(prop_offset) > 0 {
			ptr = (*Zval)((*byte)(zobj + prop_offset))
			if ptr.GetType() != 0 {
				result.GetValue().SetZv(ptr)
				result.SetTypeInfo(13)
				if flags != 0 {
					var prop_info *ZendPropertyInfo = (cache_slot + 2)[0]
					if prop_info != nil {
						ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags)
					}
				}
				return
			}
		} else if zobj.GetProperties() != nil {
			if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
				if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(zobj.GetProperties()).gc)
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			ptr = ZendHashFindEx(zobj.GetProperties(), prop_ptr.GetValue().GetStr(), 1)
			if ptr != nil {
				result.GetValue().SetZv(ptr)
				result.SetTypeInfo(13)
				return
			}
		}
	}
	ptr = container.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(container, prop_ptr, type_, cache_slot)
	if nil == ptr {
		ptr = container.GetValue().GetObj().GetHandlers().GetReadProperty()(container, prop_ptr, type_, cache_slot, result)
		if ptr == result {
			if ptr.GetType() == 10 && ZvalRefcountP(ptr) == 1 {
				var _z *Zval = ptr
				var ref *ZendReference
				assert(_z.GetType() == 10)
				ref = _z.GetValue().GetRef()
				var _z1 *Zval = _z
				var _z2 *Zval = &ref.val
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_efree(ref)
			}
			return
		}
		if EG.GetException() != nil {
			result.SetTypeInfo(15)
			return
		}
	} else if ptr.GetType() == 15 {
		result.SetTypeInfo(15)
		return
	}
	result.GetValue().SetZv(ptr)
	result.SetTypeInfo(13)
	if flags != 0 {
		var prop_info *ZendPropertyInfo
		if prop_op_type == 1<<0 {
			prop_info = (cache_slot + 2)[0]
			if prop_info != nil {
				if ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags) == 0 {
					return
				}
			}
		} else {
			if ZendHandleFetchObjFlags(result, ptr, container.GetValue().GetObj(), nil, flags) == 0 {
				return
			}
		}
	}
	if init_undef != 0 && ptr.GetType() == 0 {
		ptr.SetTypeInfo(1)
	}
}
func ZendAssignToPropertyReference(container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	var variable Zval
	var variable_ptr *Zval = &variable
	var cache_addr *any = g.CondF1(prop_op_type == 1<<0, func() *any {
		return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))
	}, nil)
	ZendFetchPropertyAddress(variable_ptr, container, container_op_type, prop_ptr, prop_op_type, cache_addr, 1, 0, 0, opline, execute_data)
	if variable_ptr.GetType() == 13 {
		variable_ptr = variable_ptr.GetValue().GetZv()
	}
	if variable_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if variable.GetType() != 13 {
		ZendThrowError(nil, "Cannot assign by reference to overloaded object")
		ZvalPtrDtor(&variable)
		variable_ptr = &EG.uninitialized_zval
	} else if value_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if (opline.GetExtendedValue()&1<<0) != 0 && value_ptr.GetType() != 10 {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, execute_data)
	} else {
		var prop_info *ZendPropertyInfo = nil
		if prop_op_type == 1<<0 {
			prop_info = (*ZendPropertyInfo)((cache_addr + 2)[0])
		} else {
			if container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
			}
			prop_info = ZendObjectFetchPropertyTypeInfo(container.GetValue().GetObj(), variable_ptr)
		}
		if prop_info != nil {
			variable_ptr = ZendAssignToTypedPropertyReference(prop_info, variable_ptr, value_ptr, execute_data)
		} else {
			ZendAssignToVariableReference(variable_ptr, value_ptr)
		}
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = variable_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
}
func ZendAssignToPropertyReferenceThisConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	ZendAssignToPropertyReference(container, 0, prop_ptr, 1<<0, value_ptr, opline, execute_data)
}
func ZendAssignToPropertyReferenceVarConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	ZendAssignToPropertyReference(container, 1<<2, prop_ptr, 1<<0, value_ptr, opline, execute_data)
}
func ZendAssignToPropertyReferenceThisVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	ZendAssignToPropertyReference(container, 0, prop_ptr, 1<<2, value_ptr, opline, execute_data)
}
func ZendAssignToPropertyReferenceVarVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, execute_data *ZendExecuteData) {
	ZendAssignToPropertyReference(container, 1<<2, prop_ptr, 1<<2, value_ptr, opline, execute_data)
}
func ZendFetchStaticPropertyAddressEx(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, opline *ZendOp, execute_data *ZendExecuteData) int {
	var free_op1 ZendFreeOp
	var name *ZendString
	var tmp_name *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var op1_type ZendUchar = opline.GetOp1Type()
	var op2_type ZendUchar = opline.GetOp2Type()
	if op2_type == 1<<0 {
		var class_name *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		assert(op1_type != 1<<0 || (*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))[0] == nil)
		if g.Assign(&ce, (*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))[0]) == nil {
			ce = ZendFetchClassByName(class_name.GetValue().GetStr(), (class_name + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				if (op1_type & (1<<1 | 1<<2)) != 0 {
					ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
				}
				return FAILURE
			}
			if op1_type != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))[0] = ce
			}
		}
	} else {
		if op2_type == 0 {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				if (op1_type & (1<<1 | 1<<2)) != 0 {
					ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
				}
				return FAILURE
			}
		} else {
			ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())).GetValue().GetCe()
		}
		if op1_type == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))[0] == ce {
			*retval = (*any)((*byte)(execute_data.GetRunTimeCache() + (cache_slot + g.SizeOf("void *"))))[0]
			*prop_info = (*any)((*byte)(execute_data.GetRunTimeCache() + (cache_slot + g.SizeOf("void *")*2)))[0]
			return SUCCESS
		}
	}
	if op1_type == 1<<0 {
		name = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant).GetValue().GetStr()
	} else {
		var varname *Zval = _getZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
		if varname.GetType() == 6 {
			name = varname.GetValue().GetStr()
			tmp_name = nil
		} else {
			if op1_type == 1<<3 && varname.GetType() == 0 {
				ZvalUndefinedCv(opline.GetOp1().GetVar(), execute_data)
			}
			name = ZvalGetTmpString(varname, &tmp_name)
		}
	}
	*retval = ZendStdGetStaticPropertyWithInfo(ce, name, fetch_type, &property_info)
	if op1_type != 1<<0 {
		ZendTmpStringRelease(tmp_name)
		if op1_type != 1<<3 {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	if (*retval) == nil {
		return FAILURE
	}
	*prop_info = property_info
	if op1_type == 1<<0 {
		var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))
		slot[0] = ce
		slot[1] = *retval
		(*any)((*byte)(execute_data.GetRunTimeCache() + (cache_slot + g.SizeOf("void *")*2)))[0] = property_info
	}
	return SUCCESS
}
func ZendFetchStaticPropertyAddress(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, flags int, opline *ZendOp, execute_data *ZendExecuteData) int {
	var success int
	var property_info *ZendPropertyInfo
	if opline.GetOp1Type() == 1<<0 && (opline.GetOp2Type() == 1<<0 || opline.GetOp2Type() == 0 && (opline.GetOp2().GetNum() == 1 || opline.GetOp2().GetNum() == 2)) && (*any)((*byte)(execute_data.GetRunTimeCache() + cache_slot))[0] != nil {
		*retval = (*any)((*byte)(execute_data.GetRunTimeCache() + (cache_slot + g.SizeOf("void *"))))[0]
		property_info = (*any)((*byte)(execute_data.GetRunTimeCache() + (cache_slot + g.SizeOf("void *")*2)))[0]
		if (fetch_type == 0 || fetch_type == 2) && (*retval).GetType() == 0 && property_info.GetType() != 0 {
			ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(property_info.GetName()))
			return FAILURE
		}
	} else {
		success = ZendFetchStaticPropertyAddressEx(retval, &property_info, cache_slot, fetch_type, opline, execute_data)
		if success != SUCCESS {
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
	ZendTypeError("Reference with value of type %s held by property %s::$%s of type %s%s is not compatible with property %s::$%s of type %s%s", g.CondF(zv.GetType() == 8, func() []byte { return zv.GetValue().GetObj().GetCe().GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func ZendThrowRefTypeErrorZval(prop *ZendPropertyInfo, zv *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s", g.CondF(zv.GetType() == 8, func() []byte { return zv.GetValue().GetObj().GetCe().GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowConflictingCoercionError(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s and property %s::$%s of type %s%s, as this would result in an inconsistent type conversion", g.CondF(zv.GetType() == 8, func() []byte { return zv.GetValue().GetObj().GetCe().GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}

/* 1: valid, 0: invalid, -1: may be valid after type coercion */

func IZendVerifyTypeAssignableZval(type_ptr *ZendType, self_ce *ZendClassEntry, zv *Zval, strict ZendBool) int {
	var type_ ZendType = *type_ptr
	var type_code ZendUchar
	var zv_type ZendUchar = zv.GetType()
	if (type_&0x1) != 0 && zv_type == 1 {
		return 1
	}
	if type_ > 0x3ff {
		if (type_ & 0x2) == 0 {
			if ZendResolveClassType(type_ptr, self_ce) == 0 {
				return 0
			}
			type_ = *type_ptr
		}
		return zv_type == 8 && InstanceofFunction(zv.GetValue().GetObj().GetCe(), (*ZendClassEntry)(type_ & ^0x3)) != 0
	}
	type_code = type_ >> 2
	if type_code == zv_type || type_code == 16 && (zv_type == 2 || zv_type == 3) {
		return 1
	}
	if type_code == 18 {
		return ZendIsIterable(zv)
	}

	/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	if strict != 0 {
		if type_code == 5 && zv_type == 4 {
			return -1
		}
		return 0
	}

	/* No weak conversions for arrays and objects */

	if type_code == 7 || type_code == 8 {
		return 0
	}

	/* NULL may be accepted only by nullable hints (this is already checked) */

	if zv_type == 1 {
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
	assert(zv.GetType() != 10)
	var _source_list *ZendPropertyInfoSourceList = &ref.sources
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if (_source_list.GetList() & 0x1) != 0 {
			_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
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
				if prop.GetType() > 0x3ff {
					seen_type = 8
				} else {
					seen_type = prop.GetType() >> 2
				}
			} else if needs_coercion != 0 && seen_type != prop.GetType()>>2 {
				ZendThrowConflictingCoercionError(seen_prop, prop, zv)
				return 0
			}
		}
	}
	if needs_coercion != 0 && ZendVerifyWeakScalarTypeHint(seen_type, zv) == 0 {
		ZendThrowRefTypeErrorZval(seen_prop, zv)
		return 0
	}
	return 1
}
func IZvalPtrDtorNoref(zval_ptr *Zval) {
	if zval_ptr.GetTypeFlags() != 0 {
		var ref *ZendRefcounted = zval_ptr.GetValue().GetCounted()
		assert(zval_ptr.GetType() != 10)
		if ZendGcDelref(&ref.gc) == 0 {
			RcDtorFunc(ref)
		} else if (ref.GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
			GcPossibleRoot(ref)
		}
	}
}
func ZendAssignToTypedRef(variable_ptr *Zval, orig_value *Zval, value_type ZendUchar, strict ZendBool, ref *ZendRefcounted) *Zval {
	var ret ZendBool
	var value Zval
	var _z1 *Zval = &value
	var _z2 *Zval = orig_value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	ret = ZendVerifyRefAssignableZval(variable_ptr.GetValue().GetRef(), &value, strict)
	variable_ptr = &(*variable_ptr).value.GetRef().GetVal()
	if ret != 0 {
		IZvalPtrDtorNoref(variable_ptr)
		var _z1 *Zval = variable_ptr
		var _z2 *Zval = &value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		ZvalPtrDtorNogc(&value)
	}
	if (value_type & (1<<2 | 1<<1)) != 0 {
		if ref != nil {
			if ZendGcDelref(&ref.gc) == 0 {
				ZvalPtrDtor(orig_value)
				_efree(ref)
			}
		} else {
			IZvalPtrDtorNoref(orig_value)
		}
	}
	return variable_ptr
}
func ZendVerifyPropAssignableByRef(prop_info *ZendPropertyInfo, orig_val *Zval, strict ZendBool) ZendBool {
	var val *Zval = orig_val
	if val.GetType() == 10 && val.GetValue().GetRef().GetSources().GetPtr() != nil {
		var result int
		val = &(*val).value.GetRef().GetVal()
		result = IZendVerifyTypeAssignableZval(&prop_info.type_, prop_info.GetCe(), val, strict)
		if result > 0 {
			return 1
		}
		if result < 0 {
			var ref_prop *ZendPropertyInfo = g.CondF((orig_val.GetValue().GetRef().GetSources().GetList()&0x1) != 0, func() *ZendPropertyInfo {
				return (*ZendPropertyInfoList)(orig_val.GetValue().GetRef().GetSources().GetList() & ^0x1).GetPtr()[0]
			}, func() *ZendPropertyInfo { return orig_val.GetValue().GetRef().GetSources().GetPtr() })
			if prop_info.GetType()>>2 != ref_prop.GetType()>>2 {

				/* Invalid due to conflicting coercion */

				ZendThrowRefTypeErrorType(ref_prop, prop_info, val)
				return 0
			}
			if ZendVerifyWeakScalarTypeHint(prop_info.GetType()>>2, val) != 0 {
				return 1
			}
		}
	} else {
		if val.GetType() == 10 {
			val = &(*val).value.GetRef().GetVal()
		}
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
	list = (*ZendPropertyInfoList)(source_list.GetList() & ^0x1)
	if (source_list.GetList() & 0x1) == 0 {
		list = _emalloc(g.SizeOf("zend_property_info_list") + (4-1)*g.SizeOf("zend_property_info *"))
		list.GetPtr()[0] = source_list.GetPtr()
		list.SetNumAllocated(4)
		list.SetNum(1)
	} else if list.GetNumAllocated() == list.GetNum() {
		list.SetNumAllocated(list.GetNum() * 2)
		list = _erealloc(list, g.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*g.SizeOf("zend_property_info *"))
	}
	list.GetPtr()[g.PostInc(&(list.GetNum()))] = prop
	source_list.SetList(0x1 | uintptr_t(list))
}
func ZendRefDelTypeSource(source_list *ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *ZendPropertyInfoList = (*ZendPropertyInfoList)(source_list.GetList() & ^0x1)
	var ptr **ZendPropertyInfo
	var end ***ZendPropertyInfo
	if (source_list.GetList() & 0x1) == 0 {
		assert(source_list.GetPtr() == prop)
		source_list.SetPtr(nil)
		return
	}
	if list.GetNum() == 1 {
		assert((*list).ptr == prop)
		_efree(list)
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
	assert((*ptr) == prop)

	/* Copy the last list element into the deleted slot. */

	*ptr = list.GetPtr()[g.PreDec(&(list.GetNum()))]
	if list.GetNum() >= 4 && list.GetNum()*4 == list.GetNumAllocated() {
		list.SetNumAllocated(list.GetNum() * 2)
		source_list.SetList(0x1 | uintptr_t(_erealloc(list, g.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*g.SizeOf("zend_property_info *"))))
	}
}
func ZendFetchThisVar(type_ int, opline *ZendOp, execute_data *ZendExecuteData) {
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	switch type_ {
	case 0:
		if execute_data.GetThis().GetType() == 8 {
			var __z *Zval = result
			__z.GetValue().SetObj(execute_data.GetThis().GetValue().GetObj())
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			ZvalAddrefP(result)
		} else {
			result.SetTypeInfo(1)
			ZendError(1<<3, "Undefined variable: this")
		}
		break
	case 3:
		if execute_data.GetThis().GetType() == 8 {
			var __z *Zval = result
			__z.GetValue().SetObj(execute_data.GetThis().GetValue().GetObj())
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			ZvalAddrefP(result)
		} else {
			result.SetTypeInfo(1)
		}
		break
	case 2:

	case 1:
		result.SetTypeInfo(0)
		ZendThrowError(nil, "Cannot re-assign $this")
		break
	case 5:
		result.SetTypeInfo(0)
		ZendThrowError(nil, "Cannot unset $this")
		break
	default:
		break
	}
}
func ZendWrongCloneCall(clone *ZendFunction, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s %s::__clone() from context '%s'", ZendVisibilityString(clone.GetFnFlags()), clone.GetScope().GetName().GetVal(), g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
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
	if EG.GetSymtableCachePtr() >= EG.GetSymtableCacheLimit() {
		ZendArrayDestroy(symbol_table)
	} else {
		*(g.PostInc(&(EG.GetSymtableCachePtr()))) = symbol_table
	}
}

/* }}} */

func IFreeCompiledVariables(execute_data *ZendExecuteData) {
	var cv *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(0))
	var count int = execute_data.GetFunc().GetOpArray().GetLastVar()
	for count != 0 {
		if cv.GetTypeFlags() != 0 {
			var r *ZendRefcounted = cv.GetValue().GetCounted()
			if ZendGcDelref(&r.gc) == 0 {
				cv.SetTypeInfo(1)
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

// #define ZEND_VM_INTERRUPT_CHECK() do { if ( UNEXPECTED ( EG ( vm_interrupt ) ) ) { ZEND_VM_INTERRUPT ( ) ; } } while ( 0 )

// #define ZEND_VM_LOOP_INTERRUPT_CHECK() do { if ( UNEXPECTED ( EG ( vm_interrupt ) ) ) { ZEND_VM_LOOP_INTERRUPT ( ) ; } } while ( 0 )

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

func ZendCopyExtraArgs(execute_data *ZendExecuteData) {
	var op_array *ZendOpArray = &(execute_data.GetFunc()).op_array
	var first_extra_arg uint32 = op_array.GetNumArgs()
	var num_args uint32 = execute_data.GetThis().GetNumArgs()
	var src *Zval
	var delta int
	var count uint32
	var type_flags uint32 = 0
	if (op_array.GetFnFlags() & 1 << 8) == 0 {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		execute_data.SetOpline(execute_data.GetOpline() + first_extra_arg)

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* move extra args into separate array after all CV and TMP vars */

	src = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(num_args-1))
	delta = op_array.GetLastVar() + op_array.GetT() - first_extra_arg
	count = num_args - first_extra_arg
	if delta != 0 {
		delta *= g.SizeOf("zval")
		for {
			type_flags |= src.GetTypeInfo()
			var _z1 *Zval = (*Zval)((*byte)(src) + delta)
			var _z2 *Zval = src
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			src.SetTypeInfo(0)
			src--
			if !(g.PreDec(&count)) {
				break
			}
		}
		if (type_flags & 0xff00) != 0 {
			execute_data.GetThis().SetTypeInfo(execute_data.GetThis().GetTypeInfo() | 1<<19)
		}
	} else {
		for {
			if src.GetTypeFlags() != 0 {
				execute_data.GetThis().SetTypeInfo(execute_data.GetThis().GetTypeInfo() | 1<<19)
				break
			}
			src--
			if !(g.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendInitCvs(first uint32, last uint32, execute_data *ZendExecuteData) {
	if first < last {
		var count uint32 = last - first
		var var_ *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(first))
		for {
			var_.SetTypeInfo(0)
			var_++
			if !(g.PreDec(&count)) {
				break
			}
		}
	}
}
func IInitFuncExecuteData(op_array *ZendOpArray, return_value *Zval, may_be_trampoline ZendBool, execute_data *ZendExecuteData) {
	var first_extra_arg uint32
	var num_args uint32
	assert(execute_data.GetFunc() == (*ZendFunction)(op_array))
	execute_data.SetOpline(op_array.GetOpcodes())
	execute_data.SetCall(nil)
	execute_data.SetReturnValue(return_value)

	/* Handle arguments */

	first_extra_arg = op_array.GetNumArgs()
	num_args = execute_data.GetThis().GetNumArgs()
	if num_args > first_extra_arg {
		if may_be_trampoline == 0 || (op_array.GetFnFlags()&1<<18) == 0 {
			ZendCopyExtraArgs(execute_data)
		}
	} else if (op_array.GetFnFlags() & 1 << 8) == 0 {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		execute_data.SetOpline(execute_data.GetOpline() + num_args)

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* Initialize CV variables (skip arguments) */

	ZendInitCvs(num_args, op_array.GetLastVar(), execute_data)
	if (uintptr_t(op_array).run_time_cache__ptr & 1) != 0 {
		execute_data.SetRunTimeCache(*((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1))))
	} else {
		execute_data.SetRunTimeCache(any(*(op_array.GetRunTimeCachePtr())))
	}
	EG.SetCurrentExecuteData(execute_data)
}

/* }}} */

func InitFuncRunTimeCacheI(op_array *ZendOpArray) {
	var run_time_cache *any
	assert(g.CondF((uintptr_t(op_array).run_time_cache__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1)))
	}, func() any { return any(*(op_array.GetRunTimeCachePtr())) }) == nil)
	run_time_cache = ZendArenaAlloc(&CG.arena, op_array.GetCacheSize())
	memset(run_time_cache, 0, op_array.GetCacheSize())
	if (uintPtr(op_array.GetRunTimeCachePtr()) & 1) != 0 {
		*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetRunTimeCachePtr()-1)))) = run_time_cache
	} else {
		*(op_array.GetRunTimeCachePtr()) = run_time_cache
	}
}

/* }}} */

func InitFuncRunTimeCache(op_array *ZendOpArray) { InitFuncRunTimeCacheI(op_array) }

/* }}} */

func ZendFetchFunction(name *ZendString) *ZendFunction {
	var zv *Zval = ZendHashFind(EG.GetFunctionTable(), name)
	if zv != nil {
		var fbc *ZendFunction = zv.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCacheI(&fbc.op_array)
		}
		return fbc
	}
	return nil
}
func ZendFetchFunctionStr(name string, len_ int) *ZendFunction {
	var zv *Zval = ZendHashStrFind(EG.GetFunctionTable(), name, len_)
	if zv != nil {
		var fbc *ZendFunction = zv.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCacheI(&fbc.op_array)
		}
		return fbc
	}
	return nil
}
func ZendInitFuncRunTimeCache(op_array *ZendOpArray) {
	if !(g.CondF((uintptr_t(op_array).run_time_cache__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1)))
	}, func() any { return any(*(op_array.GetRunTimeCachePtr())) })) {
		InitFuncRunTimeCacheI(op_array)
	}
}
func IInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	assert(execute_data.GetFunc() == (*ZendFunction)(op_array))
	execute_data.SetOpline(op_array.GetOpcodes())
	execute_data.SetCall(nil)
	execute_data.SetReturnValue(return_value)
	ZendAttachSymbolTable(execute_data)
	if op_array.GetRunTimeCachePtr() == nil {
		var ptr any
		assert((op_array.GetFnFlags() & 1 << 22) != 0)
		ptr = _emalloc(op_array.GetCacheSize() + g.SizeOf("void *"))
		op_array.SetRunTimeCachePtr(ptr)
		ptr = (*byte)(ptr + g.SizeOf("void *"))
		if (uintPtr(op_array.GetRunTimeCachePtr()) & 1) != 0 {
			*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetRunTimeCachePtr()-1)))) = ptr
		} else {
			*(op_array.GetRunTimeCachePtr()) = ptr
		}
		memset(ptr, 0, op_array.GetCacheSize())
	}
	if (uintptr_t(op_array).run_time_cache__ptr & 1) != 0 {
		execute_data.SetRunTimeCache(*((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1))))
	} else {
		execute_data.SetRunTimeCache(any(*(op_array.GetRunTimeCachePtr())))
	}
	EG.SetCurrentExecuteData(execute_data)
}

/* }}} */

func ZendInitFuncExecuteData(ex *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	var execute_data *ZendExecuteData = ex
	execute_data.SetPrevExecuteData(EG.GetCurrentExecuteData())
	if !(g.CondF((uintptr_t(op_array).run_time_cache__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1)))
	}, func() any { return any(*(op_array.GetRunTimeCachePtr())) })) {
		InitFuncRunTimeCache(op_array)
	}
	IInitFuncExecuteData(op_array, return_value, 1, execute_data)
}

/* }}} */

func ZendInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	execute_data.SetPrevExecuteData(EG.GetCurrentExecuteData())
	IInitCodeExecuteData(execute_data, op_array, return_value)
}

/* }}} */

func ZendInitExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	if (execute_data.GetThis().GetTypeInfo() & 1 << 20) != 0 {
		ZendInitCodeExecuteData(execute_data, op_array, return_value)
	} else {
		ZendInitFuncExecuteData(execute_data, op_array, return_value)
	}
}

/* }}} */

func ZendVmStackCopyCallFrame(call *ZendExecuteData, passed_args uint32, additional_args uint32) *ZendExecuteData {
	var new_call *ZendExecuteData
	var used_stack int = EG.GetVmStackTop() - (*Zval)(call) + additional_args

	/* copy call frame into new stack segment */

	new_call = ZendVmStackExtend(used_stack * g.SizeOf("zval"))
	*new_call = *call
	new_call.GetThis().SetTypeInfo(new_call.GetThis().GetTypeInfo() | 1<<18)
	if passed_args != 0 {
		var src *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		var dst *Zval = (*Zval)(new_call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		for {
			var _z1 *Zval = dst
			var _z2 *Zval = src
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			passed_args--
			src++
			dst++
			if passed_args == 0 {
				break
			}
		}
	}

	/* delete old call_frame from previous stack segment */

	EG.GetVmStack().GetPrev().SetTop((*Zval)(call))

	/* delete previous stack segment if it became empty */

	if EG.GetVmStack().GetPrev().GetTop() == (*Zval)(EG.GetVmStack().GetPrev())+((g.SizeOf("struct _zend_vm_stack")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)) {
		var r ZendVmStack = EG.GetVmStack().GetPrev()
		EG.GetVmStack().SetPrev(r.GetPrev())
		_efree(r)
	}
	return new_call
}

/* }}} */

func ZendVmStackExtendCallFrame(call **ZendExecuteData, passed_args uint32, additional_args uint32) {
	if uint32(EG.GetVmStackEnd()-EG.GetVmStackTop()) > additional_args {
		EG.SetVmStackTop(EG.GetVmStackTop() + additional_args)
	} else {
		*call = ZendVmStackCopyCallFrame(*call, passed_args, additional_args)
	}
}

/* }}} */

func ZendGetRunningGenerator(execute_data *ZendExecuteData) *ZendGenerator {
	/* The generator object is stored in EX(return_value) */

	var generator *ZendGenerator = (*ZendGenerator)(execute_data.GetReturnValue())

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */

	return generator

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */
}

/* }}} */

func CleanupUnfinishedCalls(execute_data *ZendExecuteData, op_num uint32) {
	if execute_data.GetCall() != nil {
		var call *ZendExecuteData = execute_data.GetCall()
		var opline *ZendOp = execute_data.GetFunc().GetOpArray().GetOpcodes() + op_num
		var level int
		var do_exit int
		if opline.GetOpcode() == 61 || opline.GetOpcode() == 59 || opline.GetOpcode() == 69 || opline.GetOpcode() == 128 || opline.GetOpcode() == 118 || opline.GetOpcode() == 112 || opline.GetOpcode() == 113 || opline.GetOpcode() == 68 {
			assert(op_num != 0)
			opline--
		}
		for {

			/* If the exception was thrown during a function call there might be
			 * arguments pushed to the stack that have to be dtor'ed. */

			level = 0
			do_exit = 0
			for {
				switch opline.GetOpcode() {
				case 60:

				case 129:

				case 130:

				case 131:
					level++
					break
				case 61:

				case 59:

				case 69:

				case 128:

				case 118:

				case 112:

				case 113:

				case 68:
					if level == 0 {
						call.GetThis().SetNumArgs(0)
						do_exit = 1
					}
					level--
					break
				case 65:

				case 116:

				case 117:

				case 66:

				case 185:

				case 67:

				case 106:

				case 50:

				case 120:
					if level == 0 {
						call.GetThis().SetNumArgs(opline.GetOp2().GetNum())
						do_exit = 1
					}
					break
				case 119:

				case 165:
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
					case 60:

					case 129:

					case 130:

					case 131:
						level++
						break
					case 61:

					case 59:

					case 69:

					case 128:

					case 118:

					case 112:

					case 113:

					case 68:
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
			ZendVmStackFreeArgs(execute_data.GetCall())
			if (call.GetThis().GetTypeInfo() & 1 << 21) != 0 {
				ZendObjectRelease(call.GetThis().GetValue().GetObj())
			}
			if (call.GetFunc().GetFnFlags() & 1 << 20) != 0 {
				ZendObjectRelease((*ZendObject)((*byte)(call.GetFunc() - g.SizeOf("zend_object"))))
			} else if (call.GetFunc().GetFnFlags() & 1 << 18) != 0 {
				ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
				if call.GetFunc() == &EG.trampoline {
					EG.GetTrampoline().SetFunctionName(nil)
				} else {
					_efree(call.GetFunc())
				}
			}
			execute_data.SetCall(call.GetPrevExecuteData())
			ZendVmStackFreeCallFrame(call)
			call = execute_data.GetCall()
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
		if op_num >= range_.GetStart() && op_num < range_.GetEnd() && var_num == (range_.GetVar() & ^7) {
			return range_
		}
	}
	return nil
}

/* }}} */

func CleanupLiveVars(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	var i int
	for i = 0; i < execute_data.GetFunc().GetOpArray().GetLastLiveRange(); i++ {
		var range_ *ZendLiveRange = &(execute_data.GetFunc()).op_array.GetLiveRange()[i]
		if range_.GetStart() > op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		} else if op_num < range_.GetEnd() {
			if catch_op_num == 0 || catch_op_num >= range_.GetEnd() {
				var kind uint32 = range_.GetVar() & 7
				var var_num uint32 = range_.GetVar() & ^7
				var var_ *Zval = (*Zval)((*byte)(execute_data) + int(var_num))
				if kind == 0 {
					ZvalPtrDtorNogc(var_)
				} else if kind == 4 {
					var obj *ZendObject
					assert(var_.GetType() == 8)
					obj = var_.GetValue().GetObj()
					ZendObjectStoreCtorFailed(obj)
					ZendObjectRelease(obj)
				} else if kind == 1 {
					if var_.GetType() != 7 && var_.GetFeIterIdx() != uint32-1 {
						ZendHashIteratorDel(var_.GetFeIterIdx())
					}
					ZvalPtrDtorNogc(var_)
				} else if kind == 3 {
					var rope **ZendString = (**ZendString)(var_)
					var last *ZendOp = execute_data.GetFunc().GetOpArray().GetOpcodes() + op_num
					for last.GetOpcode() != 55 && last.GetOpcode() != 54 || last.GetResult().GetVar() != var_num {
						assert(last >= execute_data.GetFunc().GetOpArray().GetOpcodes())
						last--
					}
					if last.GetOpcode() == 54 {
						ZendStringReleaseEx(*rope, 0)
					} else {
						var j int = last.GetExtendedValue()
						for {
							ZendStringReleaseEx(rope[j], 0)
							if !(g.PostDec(&j)) {
								break
							}
						}
					}
				} else if kind == 2 {

					/* restore previous error_reporting value */

					if EG.GetErrorReporting() == 0 && var_.GetValue().GetLval() != 0 {
						EG.SetErrorReporting(var_.GetValue().GetLval())
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
	if g.Assign(&colon, ZendMemrchr(function.GetVal(), ':', function.GetLen())) != nil && colon > function.GetVal() && (*(colon - 1)) == ':' {
		var mname *ZendString
		var cname_length int = colon - function.GetVal() - 1
		var mname_length int = function.GetLen() - cname_length - (g.SizeOf("\"::\"") - 1)
		lcname = ZendStringInit(function.GetVal(), cname_length, 0)
		called_scope = ZendFetchClassByName(lcname, nil, 0|0x200)
		if called_scope == nil {
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		mname = ZendStringInit(function.GetVal()+(cname_length+g.SizeOf("\"::\"")-1), mname_length, 0)
		if called_scope.GetGetStaticMethod() != nil {
			fbc = called_scope.GetGetStaticMethod()(called_scope, mname)
		} else {
			fbc = ZendStdGetStaticMethod(called_scope, mname, nil)
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(called_scope, mname)
			}
			ZendStringReleaseEx(lcname, 0)
			ZendStringReleaseEx(mname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		ZendStringReleaseEx(mname, 0)
		if (fbc.GetFnFlags() & 1 << 4) == 0 {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return nil
			}
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	} else {
		if function.GetVal()[0] == '\\' {
			lcname = ZendStringAlloc(function.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), function.GetVal()+1, function.GetLen()-1)
		} else {
			lcname = ZendStringTolowerEx(function, 0)
		}
		if g.Assign(&func_, ZendHashFind(EG.GetFunctionTable(), lcname)) == nil {
			ZendThrowError(nil, "Call to undefined function %s()", function.GetVal())
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		fbc = func_.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		called_scope = nil
	}
	return ZendVmStackPushCallFrame(0<<16|0<<17|1<<25, fbc, num_args, called_scope)
}

/* }}} */

func ZendInitDynamicCallObject(function *Zval, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var called_scope *ZendClassEntry
	var object *ZendObject
	var call_info uint32 = 0<<16 | 0<<17 | 1<<25
	if function.GetValue().GetObj().GetHandlers().GetGetClosure() != nil && function.GetValue().GetObj().GetHandlers().GetGetClosure()(function, &called_scope, &fbc, &object) == SUCCESS {
		object_or_called_scope = called_scope
		if (fbc.GetFnFlags() & 1 << 20) != 0 {

			/* Delay closure destruction until its invocation */

			ZendGcAddref(&((*ZendObject)((*byte)(fbc - g.SizeOf("zend_object")))).gc)
			call_info |= 1 << 22
			if (fbc.GetFnFlags() & 1 << 21) != 0 {
				call_info |= 1 << 23
			}
			if object != nil {
				call_info |= 8 | 1<<0<<8 | 1<<1<<8
				object_or_called_scope = object
			}
		} else if object != nil {
			call_info |= 1<<21 | (8 | 1<<0<<8 | 1<<1<<8)
			ZendGcAddref(&object.gc)
			object_or_called_scope = object
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
	}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
		InitFuncRunTimeCache(&fbc.op_array)
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}

/* }}} */

func ZendInitDynamicCallArray(function *ZendArray, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var call_info uint32 = 0<<16 | 0<<17 | 1<<25
	if function.GetNNumOfElements() == 2 {
		var obj *Zval
		var method *Zval
		obj = ZendHashIndexFind(function, 0)
		method = ZendHashIndexFind(function, 1)
		if obj == nil || method == nil {
			ZendThrowError(nil, "Array callback has to contain indices 0 and 1")
			return nil
		}
		if obj.GetType() == 10 {
			obj = &(*obj).value.GetRef().GetVal()
		}
		if obj.GetType() != 6 && obj.GetType() != 8 {
			ZendThrowError(nil, "First array member is not a valid class name or object")
			return nil
		}
		if method.GetType() == 10 {
			method = &(*method).value.GetRef().GetVal()
		}
		if method.GetType() != 6 {
			ZendThrowError(nil, "Second array member is not a valid method")
			return nil
		}
		if obj.GetType() == 6 {
			var called_scope *ZendClassEntry = ZendFetchClassByName(obj.GetValue().GetStr(), nil, 0|0x200)
			if called_scope == nil {
				return nil
			}
			if called_scope.GetGetStaticMethod() != nil {
				fbc = called_scope.GetGetStaticMethod()(called_scope, method.GetValue().GetStr())
			} else {
				fbc = ZendStdGetStaticMethod(called_scope, method.GetValue().GetStr(), nil)
			}
			if fbc == nil {
				if EG.GetException() == nil {
					ZendUndefinedMethod(called_scope, method.GetValue().GetStr())
				}
				return nil
			}
			if (fbc.GetFnFlags() & 1 << 4) == 0 {
				ZendNonStaticMethodCall(fbc)
				if EG.GetException() != nil {
					return nil
				}
			}
			object_or_called_scope = called_scope
		} else {
			var object *ZendObject = obj.GetValue().GetObj()
			fbc = obj.GetValue().GetObj().GetHandlers().GetGetMethod()(&object, method.GetValue().GetStr(), nil)
			if fbc == nil {
				if EG.GetException() == nil {
					ZendUndefinedMethod(object.GetCe(), method.GetValue().GetStr())
				}
				return nil
			}
			if (fbc.GetFnFlags() & 1 << 4) != 0 {
				object_or_called_scope = object.GetCe()
			} else {
				call_info |= 1<<21 | (8 | 1<<0<<8 | 1<<1<<8)
				ZendGcAddref(&object.gc)
				object_or_called_scope = object
			}
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
	}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
		InitFuncRunTimeCache(&fbc.op_array)
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}

/* }}} */

// #define ZEND_FAKE_OP_ARRAY       ( ( zend_op_array * ) ( zend_intptr_t ) - 1 )

func ZendIncludeOrEval(inc_filename *Zval, type_ int) *ZendOpArray {
	var new_op_array *ZendOpArray = nil
	var tmp_inc_filename Zval
	&tmp_inc_filename.SetTypeInfo(0)
	if inc_filename.GetType() != 6 {
		var tmp *ZendString = ZvalTryGetStringFunc(inc_filename)
		if tmp == nil {
			return nil
		}
		var __z *Zval = &tmp_inc_filename
		var __s *ZendString = tmp
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		inc_filename = &tmp_inc_filename
	}
	switch type_ {
	case 1 << 2:

	case 1 << 4:
		var file_handle ZendFileHandle
		var resolved_path *ZendString
		resolved_path = ZendResolvePath(inc_filename.GetValue().GetStr().GetVal(), inc_filename.GetValue().GetStr().GetLen())
		if resolved_path != nil {
			if ZendHashExists(&EG.included_files, resolved_path) != 0 {
				goto already_compiled
			}
		} else if EG.GetException() != nil {
			break
		} else if strlen(inc_filename.GetValue().GetStr().GetVal()) != inc_filename.GetValue().GetStr().GetLen() {
			ZendMessageDispatcher(g.Cond(type_ == 1<<2, 1, 2), inc_filename.GetValue().GetStr().GetVal())
			break
		} else {
			resolved_path = ZendStringCopy(inc_filename.GetValue().GetStr())
		}
		if SUCCESS == ZendStreamOpen(resolved_path.GetVal(), &file_handle) {
			if file_handle.GetOpenedPath() == nil {
				file_handle.SetOpenedPath(ZendStringCopy(resolved_path))
			}
			if ZendHashAddEmptyElement(&EG.included_files, file_handle.GetOpenedPath()) != nil {
				var op_array *ZendOpArray = ZendCompileFile(&file_handle, g.Cond(type_ == 1<<2, 1<<1, 1<<3))
				ZendDestroyFileHandle(&file_handle)
				ZendStringReleaseEx(resolved_path, 0)
				if tmp_inc_filename.GetType() != 0 {
					ZvalPtrDtorStr(&tmp_inc_filename)
				}
				return op_array
			} else {
				ZendFileHandleDtor(&file_handle)
			already_compiled:
				new_op_array = (*ZendOpArray)(zend_intptr_t - 1)
			}
		} else {
			ZendMessageDispatcher(g.Cond(type_ == 1<<2, 1, 2), inc_filename.GetValue().GetStr().GetVal())
		}
		ZendStringReleaseEx(resolved_path, 0)
		break
	case 1 << 1:

	case 1 << 3:
		if strlen(inc_filename.GetValue().GetStr().GetVal()) != inc_filename.GetValue().GetStr().GetLen() {
			ZendMessageDispatcher(g.Cond(type_ == 1<<1, 1, 2), inc_filename.GetValue().GetStr().GetVal())
			break
		}
		new_op_array = CompileFilename(type_, inc_filename)
		break
	case 1 << 0:
		var eval_desc *byte = ZendMakeCompiledStringDescription("eval()'d code")
		new_op_array = ZendCompileString(inc_filename, eval_desc)
		_efree(eval_desc)
		break
	default:
		break
	}
	if tmp_inc_filename.GetType() != 0 {
		ZvalPtrDtorStr(&tmp_inc_filename)
	}
	return new_op_array
}

/* }}} */

func ZendDoFcallOverloaded(call *ZendExecuteData, ret *Zval) int {
	var fbc *ZendFunction = call.GetFunc()
	var object *ZendObject

	/* Not sure what should be done here if it's a static method */

	if call.GetThis().GetType() != 8 {
		ZendVmStackFreeArgs(call)
		if fbc.GetType() == 5 {
			ZendStringReleaseEx(fbc.GetFunctionName(), 0)
		}
		_efree(fbc)
		ZendVmStackFreeCallFrame(call)
		ZendThrowError(nil, "Cannot call overloaded function for non-object")
		return 0
	}
	object = call.GetThis().GetValue().GetObj()
	ret.SetTypeInfo(1)
	EG.SetCurrentExecuteData(call)
	object.GetHandlers().GetCallMethod()(fbc.GetFunctionName(), object, call, ret)
	EG.SetCurrentExecuteData(call.GetPrevExecuteData())
	ZendVmStackFreeArgs(call)
	if fbc.GetType() == 5 {
		ZendStringReleaseEx(fbc.GetFunctionName(), 0)
	}
	_efree(fbc)
	return 1
}

/* }}} */

func ZendFeResetIterator(array_ptr *Zval, by_ref int, opline *ZendOp, execute_data *ZendExecuteData) ZendBool {
	var ce *ZendClassEntry = array_ptr.GetValue().GetObj().GetCe()
	var iter *ZendObjectIterator = ce.GetGetIterator()(ce, array_ptr, by_ref)
	var is_empty ZendBool
	if iter == nil || EG.GetException() != nil {
		if iter != nil {
			ZendObjectRelease(&iter.std)
		}
		if EG.GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 1
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if EG.GetException() != nil {
			ZendObjectRelease(&iter.std)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 1
		}
	}
	is_empty = iter.GetFuncs().GetValid()(iter) != SUCCESS
	if EG.GetException() != nil {
		ZendObjectRelease(&iter.std)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 1
	}
	iter.SetIndex(-1)
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetObj(&iter.std)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
	return is_empty
}

/* }}} */

func _zendQuickGetConstant(key *Zval, flags uint32, check_defined_only int, opline *ZendOp, execute_data *ZendExecuteData) int {
	var zv *Zval
	var orig_key *Zval = key
	var c *ZendConstant = nil
	zv = ZendHashFindEx(EG.GetZendConstants(), key.GetValue().GetStr(), 1)
	if zv != nil {
		c = (*ZendConstant)(zv.GetValue().GetPtr())
	} else {
		key++
		zv = ZendHashFindEx(EG.GetZendConstants(), key.GetValue().GetStr(), 1)
		if zv != nil && ((*ZendConstant)(zv.GetValue().GetPtr()).GetValue().GetConstantFlags()&0xff&1<<0) == 0 {
			c = (*ZendConstant)(zv.GetValue().GetPtr())
		} else {
			if (flags & (0x100 | 0x10)) == (0x100 | 0x10) {
				key++
				zv = ZendHashFindEx(EG.GetZendConstants(), key.GetValue().GetStr(), 1)
				if zv != nil {
					c = (*ZendConstant)(zv.GetValue().GetPtr())
				} else {
					key++
					zv = ZendHashFindEx(EG.GetZendConstants(), key.GetValue().GetStr(), 1)
					if zv != nil && ((*ZendConstant)(zv.GetValue().GetPtr()).GetValue().GetConstantFlags()&0xff&1<<0) == 0 {
						c = (*ZendConstant)(zv.GetValue().GetPtr())
					}
				}
			}
		}
	}
	if c == nil {
		if check_defined_only == 0 {
			if (opline.GetOp1().GetNum() & 0x10) != 0 {
				var actual *byte = (*byte)(ZendMemrchr((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal(), '\\', (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetLen()))
				if actual == nil {
					var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var __s *ZendString = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetStr()
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
				} else {
					actual++
					var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var __s *ZendString = ZendStringInit(actual, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetLen()-(actual-(*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal()), 0)
					__z.GetValue().SetStr(__s)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}

				/* non-qualified constant - allow text substitution */

				ZendError(1<<1, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetStr().GetVal(), (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetStr().GetVal())

				/* non-qualified constant - allow text substitution */

			} else {
				ZendThrowError(nil, "Undefined constant '%s'", (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal())
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
		}
		return FAILURE
	}
	if check_defined_only == 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = &c.value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
				ZendGcAddref(&_gc.gc)
			} else {
				ZvalCopyCtorFunc(_z1)
			}
		}
		if (c.GetValue().GetConstantFlags() & 0xff & (1<<0 | 1<<2)) == 0 {
			var ns_sep *byte
			var shortname_offset int
			var shortname_len int
			var is_deprecated ZendBool
			if (flags & 0x10) != 0 {
				var access_key *Zval
				if (flags & 0x100) == 0 {
					access_key = orig_key - 1
				} else {
					if key < orig_key+2 {
						goto check_short_name
					} else {
						access_key = orig_key + 2
					}
				}
				is_deprecated = !(ZendStringEquals(c.GetName(), access_key.GetValue().GetStr()))
			} else {
			check_short_name:

				/* Namespaces are always case-insensitive. Only compare shortname. */

				ns_sep = ZendMemrchr(c.GetName().GetVal(), '\\', c.GetName().GetLen())
				if ns_sep != nil {
					shortname_offset = ns_sep - c.GetName().GetVal() + 1
					shortname_len = c.GetName().GetLen() - shortname_offset
				} else {
					shortname_offset = 0
					shortname_len = c.GetName().GetLen()
				}
				is_deprecated = memcmp(c.GetName().GetVal()+shortname_offset, (orig_key-1).value.str.val+shortname_offset, shortname_len) != 0
			}
			if is_deprecated != 0 {
				ZendError(1<<13, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
				return SUCCESS
			}
		}
	}
	(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = c
	return SUCCESS
}

/* }}} */

func ZendQuickGetConstant(key *Zval, flags uint32, opline *ZendOp, execute_data *ZendExecuteData) {
	_zendQuickGetConstant(key, flags, 0, opline, execute_data)
}
func ZendQuickCheckConstant(key *Zval, opline *ZendOp, execute_data *ZendExecuteData) int {
	return _zendQuickGetConstant(key, 0, 1, opline, execute_data)
}

// #define _zend_vm_stack_push_call_frame_ex       zend_vm_stack_push_call_frame_ex

// #define _zend_vm_stack_push_call_frame       zend_vm_stack_push_call_frame

// #define ZEND_VM_NEXT_OPCODE_EX(check_exception,skip) CHECK_SYMBOL_TABLES ( ) if ( check_exception ) { OPLINE = EX ( opline ) + ( skip ) ; } else { ZEND_ASSERT ( ! EG ( exception ) ) ; OPLINE = opline + ( skip ) ; } ZEND_VM_CONTINUE ( )

// #define ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION() ZEND_VM_NEXT_OPCODE_EX ( 1 , 1 )

// #define ZEND_VM_NEXT_OPCODE() ZEND_VM_NEXT_OPCODE_EX ( 0 , 1 )

// #define ZEND_VM_SET_NEXT_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op

// #define ZEND_VM_SET_OPCODE(new_op) CHECK_SYMBOL_TABLES ( ) OPLINE = new_op ; ZEND_VM_INTERRUPT_CHECK ( )

// #define ZEND_VM_SET_RELATIVE_OPCODE(opline,offset) ZEND_VM_SET_OPCODE ( ZEND_OFFSET_TO_OPLINE ( opline , offset ) )

// #define ZEND_VM_JMP_EX(new_op,check_exception) do { if ( check_exception && UNEXPECTED ( EG ( exception ) ) ) { HANDLE_EXCEPTION ( ) ; } ZEND_VM_SET_OPCODE ( new_op ) ; ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_JMP(new_op) ZEND_VM_JMP_EX ( new_op , 1 )

// #define ZEND_VM_INC_OPCODE() OPLINE ++

// #define VM_SMART_OPCODES       1

// #define ZEND_VM_REPEATABLE_OPCODE       do {

// #define ZEND_VM_REPEAT_OPCODE(_opcode) } while ( UNEXPECTED ( ( ++ opline ) -> opcode == _opcode ) ) ; OPLINE = opline ; ZEND_VM_CONTINUE ( )

// #define ZEND_VM_SMART_BRANCH(_result,_check) do { if ( ( _check ) && UNEXPECTED ( EG ( exception ) ) ) { break ; } if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPZ ) ) { if ( _result ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; } else { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; } } else if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPNZ ) ) { if ( ! ( _result ) ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; } else { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; } } else { break ; } ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_JMPZ(_result,_check) do { if ( ( _check ) && UNEXPECTED ( EG ( exception ) ) ) { break ; } if ( _result ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; } else { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; } ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_JMPNZ(_result,_check) do { if ( ( _check ) && UNEXPECTED ( EG ( exception ) ) ) { break ; } if ( ! ( _result ) ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; } else { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; } ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_TRUE() do { if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPNZ ) ) { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; ZEND_VM_CONTINUE ( ) ; } else if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPZ ) ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; ZEND_VM_CONTINUE ( ) ; } } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_TRUE_JMPZ() do { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_TRUE_JMPNZ() do { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_FALSE() do { if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPNZ ) ) { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; ZEND_VM_CONTINUE ( ) ; } else if ( EXPECTED ( ( opline + 1 ) -> opcode == ZEND_JMPZ ) ) { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; ZEND_VM_CONTINUE ( ) ; } } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_FALSE_JMPZ() do { ZEND_VM_SET_OPCODE ( OP_JMP_ADDR ( opline + 1 , ( opline + 1 ) -> op2 ) ) ; ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_SMART_BRANCH_FALSE_JMPNZ() do { ZEND_VM_SET_NEXT_OPCODE ( opline + 2 ) ; ZEND_VM_CONTINUE ( ) ; } while ( 0 )

// #define ZEND_VM_GUARD(name)

// #define UNDEF_RESULT() do { if ( opline -> result_type & ( IS_VAR | IS_TMP_VAR ) ) { ZVAL_UNDEF ( EX_VAR ( opline -> result . var ) ) ; } } while ( 0 )

// # include "zend_vm_execute.h"

func ZendSetUserOpcodeHandler(opcode ZendUchar, handler UserOpcodeHandlerT) int {
	if opcode != 150 {
		if handler == nil {

			/* restore the original handler */

			ZendUserOpcodes[opcode] = opcode

			/* restore the original handler */

		} else {
			ZendUserOpcodes[opcode] = 150
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
	case 1 << 0:
		ret = (*Zval)((*byte)(opline) + int32(*node).constant)
		*should_free = nil
		break
	case 1 << 1:

	case 1 << 2:
		ret = (*Zval)((*byte)(execute_data) + int(node.GetVar()))
		*should_free = ret
		break
	case 1 << 3:
		ret = (*Zval)((*byte)(execute_data) + int(node.GetVar()))
		*should_free = nil
		break
	default:
		ret = nil
		*should_free = ret
		break
	}
	return ret
}

// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_compile.h>

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
   +----------------------------------------------------------------------+
*/

// #define ZEND_COMPILE_H

// # include "zend.h"

// # include "zend_ast.h"

// # include < stdarg . h >

// # include "zend_llist.h"

// #define SET_UNUSED(op) op ## _type = IS_UNUSED

// #define MAKE_NOP(opline) do { ( opline ) -> op1 . num = 0 ; ( opline ) -> op2 . num = 0 ; ( opline ) -> result . num = 0 ; ( opline ) -> opcode = ZEND_NOP ; ( opline ) -> op1_type = IS_UNUSED ; ( opline ) -> op2_type = IS_UNUSED ; ( opline ) -> result_type = IS_UNUSED ; } while ( 0 )

// #define RESET_DOC_COMMENT() do { if ( CG ( doc_comment ) ) { zend_string_release_ex ( CG ( doc_comment ) , 0 ) ; CG ( doc_comment ) = NULL ; } } while ( 0 )

/* On 64-bit systems less optimal, but more compact VM code leads to better
 * performance. So on 32-bit systems we use absolute addresses for jump
 * targets and constants, but on 64-bit systems realtive 32-bit offsets */

// #define ZEND_USE_ABS_JMP_ADDR       0

// #define ZEND_USE_ABS_CONST_ADDR       0

// @type ZnodeOp struct

// @type Znode struct

/* Temporarily defined here, to avoid header ordering issues */

// @type ZendAstZnode struct

func ZendAstGetZnode(ast *ZendAst) *Znode { return &((*ZendAstZnode)(ast)).node }

// @type ZendDeclarables struct

/* Compilation context that is different for each file, but shared between op arrays. */

// @type ZendFileContext struct

// @type ZendParserStackElem struct

type UserOpcodeHandlerT func(execute_data *ZendExecuteData) int

// @type ZendOp struct
// @type ZendBrkContElement struct

// @type ZendLabel struct

// @type ZendTryCatchElement struct

// #define ZEND_LIVE_TMPVAR       0

// #define ZEND_LIVE_LOOP       1

// #define ZEND_LIVE_SILENCE       2

// #define ZEND_LIVE_ROPE       3

// #define ZEND_LIVE_NEW       4

// #define ZEND_LIVE_MASK       7

// @type ZendLiveRange struct

/* Compilation context that is different for each op array. */

// @type ZendOparrayContext struct

/* Class, property and method flags                  class|meth.|prop.|const*/

// #define ZEND_ACC_PUBLIC       ( 1 << 0 )

// #define ZEND_ACC_PROTECTED       ( 1 << 1 )

// #define ZEND_ACC_PRIVATE       ( 1 << 2 )

/*                                                        |     |     |     */

// #define ZEND_ACC_CHANGED       ( 1 << 3 )

/*                                                        |     |     |     */

// #define ZEND_ACC_STATIC       ( 1 << 4 )

/*                                                        |     |     |     */

// #define ZEND_ACC_FINAL       ( 1 << 5 )

/*                                                        |     |     |     */

// #define ZEND_ACC_ABSTRACT       ( 1 << 6 )

// #define ZEND_ACC_EXPLICIT_ABSTRACT_CLASS       ( 1 << 6 )

/*                                                        |     |     |     */

// #define ZEND_ACC_IMMUTABLE       ( 1 << 7 )

/*                                                        |     |     |     */

// #define ZEND_ACC_HAS_TYPE_HINTS       ( 1 << 8 )

/*                                                        |     |     |     */

// #define ZEND_ACC_TOP_LEVEL       ( 1 << 9 )

/*                                                        |     |     |     */

// #define ZEND_ACC_PRELOADED       ( 1 << 10 )

/*                                                        |     |     |     */

// #define ZEND_ACC_INTERFACE       ( 1 << 0 )

// #define ZEND_ACC_TRAIT       ( 1 << 1 )

// #define ZEND_ACC_ANON_CLASS       ( 1 << 2 )

/*                                                        |     |     |     */

// #define ZEND_ACC_LINKED       ( 1 << 3 )

/*                                                        |     |     |     */

// #define ZEND_ACC_IMPLICIT_ABSTRACT_CLASS       ( 1 << 4 )

/*                                                        |     |     |     */

// #define ZEND_ACC_USE_GUARDS       ( 1 << 11 )

/*                                                        |     |     |     */

// #define ZEND_ACC_CONSTANTS_UPDATED       ( 1 << 12 )

/*                                                        |     |     |     */

// #define ZEND_ACC_INHERITED       ( 1 << 13 )

/*                                                        |     |     |     */

// #define ZEND_ACC_IMPLEMENT_INTERFACES       ( 1 << 14 )

/*                                                        |     |     |     */

// #define ZEND_ACC_IMPLEMENT_TRAITS       ( 1 << 15 )

/*                                                        |     |     |     */

// #define ZEND_HAS_STATIC_IN_METHODS       ( 1 << 16 )

/*                                                        |     |     |     */

// #define ZEND_ACC_PROPERTY_TYPES_RESOLVED       ( 1 << 17 )

/*                                                        |     |     |     */

// #define ZEND_ACC_REUSE_GET_ITERATOR       ( 1 << 18 )

/*                                                        |     |     |     */

// #define ZEND_ACC_RESOLVED_PARENT       ( 1 << 19 )

/*                                                        |     |     |     */

// #define ZEND_ACC_RESOLVED_INTERFACES       ( 1 << 20 )

/*                                                        |     |     |     */

// #define ZEND_ACC_UNRESOLVED_VARIANCE       ( 1 << 21 )

/*                                                        |     |     |     */

// #define ZEND_ACC_NEARLY_LINKED       ( 1 << 22 )

/*                                                        |     |     |     */

// #define ZEND_ACC_HAS_UNLINKED_USES       ( 1 << 23 )

/*                                                        |     |     |     */

// #define ZEND_ACC_DEPRECATED       ( 1 << 11 )

/*                                                        |     |     |     */

// #define ZEND_ACC_RETURN_REFERENCE       ( 1 << 12 )

/*                                                        |     |     |     */

// #define ZEND_ACC_HAS_RETURN_TYPE       ( 1 << 13 )

/*                                                        |     |     |     */

// #define ZEND_ACC_VARIADIC       ( 1 << 14 )

/*                                                        |     |     |     */

// #define ZEND_ACC_HAS_FINALLY_BLOCK       ( 1 << 15 )

/*                                                        |     |     |     */

// #define ZEND_ACC_EARLY_BINDING       ( 1 << 16 )

/*                                                        |     |     |     */

// #define ZEND_ACC_ALLOW_STATIC       ( 1 << 17 )

/*                                                        |     |     |     */

// #define ZEND_ACC_CALL_VIA_TRAMPOLINE       ( 1 << 18 )

/*                                                        |     |     |     */

// #define ZEND_ACC_NEVER_CACHE       ( 1 << 19 )

/*                                                        |     |     |     */

// #define ZEND_ACC_CLOSURE       ( 1 << 20 )

// #define ZEND_ACC_FAKE_CLOSURE       ( 1 << 21 )

/*                                                        |     |     |     */

// #define ZEND_ACC_HEAP_RT_CACHE       ( 1 << 22 )

/*                                                        |     |     |     */

// #define ZEND_ACC_USER_ARG_INFO       ( 1 << 22 )

/*                                                        |     |     |     */

// #define ZEND_ACC_GENERATOR       ( 1 << 24 )

/*                                                        |     |     |     */

// #define ZEND_ACC_DONE_PASS_TWO       ( 1 << 25 )

/*                                                        |     |     |     */

// #define ZEND_ACC_ARENA_ALLOCATED       ( 1 << 25 )

/*                                                        |     |     |     */

// #define ZEND_ACC_TRAIT_CLONE       ( 1 << 27 )

/*                                                        |     |     |     */

// #define ZEND_ACC_CTOR       ( 1 << 28 )

/*                                                        |     |     |     */

// #define ZEND_ACC_DTOR       ( 1 << 29 )

/*                                                        |     |     |     */

// #define ZEND_ACC_USES_THIS       ( 1 << 30 )

/*                                                        |     |     |     */

// #define ZEND_ACC_STRICT_TYPES       ( 1U << 31 )

// #define ZEND_ACC_PPP_MASK       ( ZEND_ACC_PUBLIC | ZEND_ACC_PROTECTED | ZEND_ACC_PRIVATE )

/* call through internal function handler. e.g. Closure::invoke() */

// #define ZEND_ACC_CALL_VIA_HANDLER       ZEND_ACC_CALL_VIA_TRAMPOLINE

// @type ZendPropertyInfo struct

// #define OBJ_PROP(obj,offset) ( ( zval * ) ( ( char * ) ( obj ) + offset ) )

// #define OBJ_PROP_NUM(obj,num) ( & ( obj ) -> properties_table [ ( num ) ] )

// #define OBJ_PROP_TO_OFFSET(num) ( ( uint32_t ) ( XtOffsetOf ( zend_object , properties_table ) + sizeof ( zval ) * ( num ) ) )

// #define OBJ_PROP_TO_NUM(offset) ( ( offset - OBJ_PROP_TO_OFFSET ( 0 ) ) / sizeof ( zval ) )

// @type ZendClassConstant struct

/* arg_info for internal functions */

// @type ZendInternalArgInfo struct

/* arg_info for user functions */

// @type ZendArgInfo struct

/* the following structure repeats the layout of zend_internal_arg_info,
 * but its fields have different meaning. It's used as the first element of
 * arg_info array to define properties __special__  of internal functions.
 * It's also used for the return type.
 */

// @type ZendInternalFunctionInfo struct

// @type ZendOpArray struct

// #define ZEND_RETURN_VALUE       0

// #define ZEND_RETURN_REFERENCE       1

/* zend_internal_function_handler */

type ZifHandler func(execute_data *ZendExecuteData, return_value *Zval)

// @type ZendInternalFunction struct

// #define ZEND_FN_SCOPE_NAME(function) ( ( function ) && ( function ) -> common . scope ? ZSTR_VAL ( ( function ) -> common . scope -> name ) : "" )

// @type ZendFunction struct
// @type ZendExecuteData struct

// #define ZEND_CALL_HAS_THIS       IS_OBJECT_EX

/* Top 16 bits of Z_TYPE_INFO(EX(This)) are used as call_info flags */

// #define ZEND_CALL_FUNCTION       ( 0 << 16 )

// #define ZEND_CALL_CODE       ( 1 << 16 )

// #define ZEND_CALL_NESTED       ( 0 << 17 )

// #define ZEND_CALL_TOP       ( 1 << 17 )

// #define ZEND_CALL_ALLOCATED       ( 1 << 18 )

// #define ZEND_CALL_FREE_EXTRA_ARGS       ( 1 << 19 )

// #define ZEND_CALL_HAS_SYMBOL_TABLE       ( 1 << 20 )

// #define ZEND_CALL_RELEASE_THIS       ( 1 << 21 )

// #define ZEND_CALL_CLOSURE       ( 1 << 22 )

// #define ZEND_CALL_FAKE_CLOSURE       ( 1 << 23 )

// #define ZEND_CALL_GENERATOR       ( 1 << 24 )

// #define ZEND_CALL_DYNAMIC       ( 1 << 25 )

// #define ZEND_CALL_SEND_ARG_BY_REF       ( 1u << 31 )

// #define ZEND_CALL_NESTED_FUNCTION       ( ZEND_CALL_FUNCTION | ZEND_CALL_NESTED )

// #define ZEND_CALL_NESTED_CODE       ( ZEND_CALL_CODE | ZEND_CALL_NESTED )

// #define ZEND_CALL_TOP_FUNCTION       ( ZEND_CALL_TOP | ZEND_CALL_FUNCTION )

// #define ZEND_CALL_TOP_CODE       ( ZEND_CALL_CODE | ZEND_CALL_TOP )

// #define ZEND_CALL_INFO(call) Z_TYPE_INFO ( ( call ) -> This )

// #define ZEND_CALL_KIND_EX(call_info) ( call_info & ( ZEND_CALL_CODE | ZEND_CALL_TOP ) )

// #define ZEND_CALL_KIND(call) ZEND_CALL_KIND_EX ( ZEND_CALL_INFO ( call ) )

// #define ZEND_ADD_CALL_FLAG_EX(call_info,flag) do { call_info |= ( flag ) ; } while ( 0 )

// #define ZEND_DEL_CALL_FLAG_EX(call_info,flag) do { call_info &= ~ ( flag ) ; } while ( 0 )

// #define ZEND_ADD_CALL_FLAG(call,flag) do { ZEND_ADD_CALL_FLAG_EX ( Z_TYPE_INFO ( ( call ) -> This ) , flag ) ; } while ( 0 )

// #define ZEND_DEL_CALL_FLAG(call,flag) do { ZEND_DEL_CALL_FLAG_EX ( Z_TYPE_INFO ( ( call ) -> This ) , flag ) ; } while ( 0 )

// #define ZEND_CALL_NUM_ARGS(call) ( call ) -> This . u2 . num_args

// #define ZEND_CALL_FRAME_SLOT       ( ( int ) ( ( ZEND_MM_ALIGNED_SIZE ( sizeof ( zend_execute_data ) ) + ZEND_MM_ALIGNED_SIZE ( sizeof ( zval ) ) - 1 ) / ZEND_MM_ALIGNED_SIZE ( sizeof ( zval ) ) ) )

// #define ZEND_CALL_VAR(call,n) ( ( zval * ) ( ( ( char * ) ( call ) ) + ( ( int ) ( n ) ) ) )

// #define ZEND_CALL_VAR_NUM(call,n) ( ( ( zval * ) ( call ) ) + ( ZEND_CALL_FRAME_SLOT + ( ( int ) ( n ) ) ) )

// #define ZEND_CALL_ARG(call,n) ZEND_CALL_VAR_NUM ( call , ( ( int ) ( n ) ) - 1 )

// #define EX(element) ( ( execute_data ) -> element )

// #define EX_CALL_INFO() ZEND_CALL_INFO ( execute_data )

// #define EX_CALL_KIND() ZEND_CALL_KIND ( execute_data )

// #define EX_NUM_ARGS() ZEND_CALL_NUM_ARGS ( execute_data )

// #define ZEND_CALL_USES_STRICT_TYPES(call) ( ( ( call ) -> func -> common . fn_flags & ZEND_ACC_STRICT_TYPES ) != 0 )

// #define EX_USES_STRICT_TYPES() ZEND_CALL_USES_STRICT_TYPES ( execute_data )

// #define ZEND_ARG_USES_STRICT_TYPES() ( EG ( current_execute_data ) -> prev_execute_data && EG ( current_execute_data ) -> prev_execute_data -> func && ZEND_CALL_USES_STRICT_TYPES ( EG ( current_execute_data ) -> prev_execute_data ) )

// #define ZEND_RET_USES_STRICT_TYPES() ZEND_CALL_USES_STRICT_TYPES ( EG ( current_execute_data ) )

// #define EX_VAR(n) ZEND_CALL_VAR ( execute_data , n )

// #define EX_VAR_NUM(n) ZEND_CALL_VAR_NUM ( execute_data , n )

// #define EX_VAR_TO_NUM(n) ( ( uint32_t ) ( ZEND_CALL_VAR ( NULL , n ) - ZEND_CALL_VAR_NUM ( NULL , 0 ) ) )

// #define ZEND_OPLINE_TO_OFFSET(opline,target) ( ( char * ) ( target ) - ( char * ) ( opline ) )

// #define ZEND_OPLINE_NUM_TO_OFFSET(op_array,opline,opline_num) ( ( char * ) & ( op_array ) -> opcodes [ opline_num ] - ( char * ) ( opline ) )

// #define ZEND_OFFSET_TO_OPLINE(base,offset) ( ( zend_op * ) ( ( ( char * ) ( base ) ) + ( int ) offset ) )

// #define ZEND_OFFSET_TO_OPLINE_NUM(op_array,base,offset) ( ZEND_OFFSET_TO_OPLINE ( base , offset ) - op_array -> opcodes )

/* run-time jump target */

// #define OP_JMP_ADDR(opline,node) ZEND_OFFSET_TO_OPLINE ( opline , ( node ) . jmp_offset )

// #define ZEND_SET_OP_JMP_ADDR(opline,node,val) do { ( node ) . jmp_offset = ZEND_OPLINE_TO_OFFSET ( opline , val ) ; } while ( 0 )

/* convert jump target from compile-time to run-time */

// #define ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array,opline,node) do { ( node ) . jmp_offset = ZEND_OPLINE_NUM_TO_OFFSET ( op_array , opline , ( node ) . opline_num ) ; } while ( 0 )

/* convert jump target back from run-time to compile-time */

// #define ZEND_PASS_TWO_UNDO_JMP_TARGET(op_array,opline,node) do { ( node ) . opline_num = ZEND_OFFSET_TO_OPLINE_NUM ( op_array , opline , ( node ) . jmp_offset ) ; } while ( 0 )

/* constant-time constant */

// #define CT_CONSTANT_EX(op_array,num) ( ( op_array ) -> literals + ( num ) )

// #define CT_CONSTANT(node) CT_CONSTANT_EX ( CG ( active_op_array ) , ( node ) . constant )

/* At run-time, constants are allocated together with op_array->opcodes
 * and addressed relatively to current opline.
 */

// #define RT_CONSTANT(opline,node) ( ( zval * ) ( ( ( char * ) ( opline ) ) + ( int32_t ) ( node ) . constant ) )

/* convert constant from compile-time to run-time */

// #define ZEND_PASS_TWO_UPDATE_CONSTANT(op_array,opline,node) do { ( node ) . constant = ( ( ( char * ) CT_CONSTANT_EX ( op_array , ( node ) . constant ) ) - ( ( char * ) opline ) ) ; } while ( 0 )

/* convert constant back from run-time to compile-time */

// #define ZEND_PASS_TWO_UNDO_CONSTANT(op_array,opline,node) do { ( node ) . constant = RT_CONSTANT ( opline , node ) - ( op_array ) -> literals ; } while ( 0 )

// #define RUN_TIME_CACHE(op_array) ZEND_MAP_PTR_GET ( ( op_array ) -> run_time_cache )

// #define ZEND_OP_ARRAY_EXTENSION(op_array,handle) ( ( void * * ) RUN_TIME_CACHE ( op_array ) ) [ handle ]

// #define IS_UNUSED       0

// #define IS_CONST       ( 1 << 0 )

// #define IS_TMP_VAR       ( 1 << 1 )

// #define IS_VAR       ( 1 << 2 )

// #define IS_CV       ( 1 << 3 )

// #define ZEND_EXTRA_VALUE       1

// # include "zend_globals.h"

var ZendCompileFile func(file_handle *ZendFileHandle, type_ int) *ZendOpArray
var ZendCompileString func(source_string *Zval, filename *byte) *ZendOpArray

type UnaryOpType func(*Zval, *Zval) int
type BinaryOpType func(*Zval, *Zval, *Zval) int

/* Used during AST construction */

/* parser-driven code generators */

var ZendDoExtendedInfo func()

// #define INITIAL_OP_ARRAY_SIZE       64

/* helper functions in zend_language_scanner.l */

// #define zend_try_exception_handler() do { if ( UNEXPECTED ( EG ( exception ) ) ) { if ( Z_TYPE ( EG ( user_exception_handler ) ) != IS_UNDEF ) { zend_user_exception_handler ( ) ; } } } while ( 0 )

// #define zend_unmangle_property_name(mangled_property,class_name,prop_name) zend_unmangle_property_name_ex ( mangled_property , class_name , prop_name , NULL )

func ZendGetUnmangledPropertyName(mangled_prop *ZendString) *byte {
	var class_name *byte
	var prop_name *byte
	ZendUnmanglePropertyNameEx(mangled_prop, &class_name, &prop_name, nil)
	return prop_name
}

// #define ZEND_FUNCTION_DTOR       zend_function_dtor

// #define ZEND_CLASS_DTOR       destroy_zend_class

type ZendNeedsLiveRangeCb func(op_array *ZendOpArray, opline *ZendOp) ZendBool
type ZendAutoGlobalCallback func(name *ZendString) ZendBool

// @type ZendAutoGlobal struct

/* BEGIN: OPCODES */

// # include "zend_vm_opcodes.h"

/* END: OPCODES */

// #define ZEND_FETCH_CLASS_DEFAULT       0

// #define ZEND_FETCH_CLASS_SELF       1

// #define ZEND_FETCH_CLASS_PARENT       2

// #define ZEND_FETCH_CLASS_STATIC       3

// #define ZEND_FETCH_CLASS_AUTO       4

// #define ZEND_FETCH_CLASS_INTERFACE       5

// #define ZEND_FETCH_CLASS_TRAIT       6

// #define ZEND_FETCH_CLASS_MASK       0x0f

// #define ZEND_FETCH_CLASS_NO_AUTOLOAD       0x80

// #define ZEND_FETCH_CLASS_SILENT       0x0100

// #define ZEND_FETCH_CLASS_EXCEPTION       0x0200

// #define ZEND_FETCH_CLASS_ALLOW_UNLINKED       0x0400

// #define ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED       0x0800

// #define ZEND_PARAM_REF       ( 1 << 0 )

// #define ZEND_PARAM_VARIADIC       ( 1 << 1 )

// #define ZEND_NAME_FQ       0

// #define ZEND_NAME_NOT_FQ       1

// #define ZEND_NAME_RELATIVE       2

// #define ZEND_TYPE_NULLABLE       ( 1 << 8 )

// #define ZEND_ARRAY_SYNTAX_LIST       1

// #define ZEND_ARRAY_SYNTAX_LONG       2

// #define ZEND_ARRAY_SYNTAX_SHORT       3

/* var status for backpatching */

// #define BP_VAR_R       0

// #define BP_VAR_W       1

// #define BP_VAR_RW       2

// #define BP_VAR_IS       3

// #define BP_VAR_FUNC_ARG       4

// #define BP_VAR_UNSET       5

// #define ZEND_INTERNAL_FUNCTION       1

// #define ZEND_USER_FUNCTION       2

// #define ZEND_OVERLOADED_FUNCTION       3

// #define ZEND_EVAL_CODE       4

// #define ZEND_OVERLOADED_FUNCTION_TEMPORARY       5

/* A quick check (type == ZEND_USER_FUNCTION || type == ZEND_EVAL_CODE) */

// #define ZEND_USER_CODE(type) ( ( type & 1 ) == 0 )

// #define ZEND_INTERNAL_CLASS       1

// #define ZEND_USER_CLASS       2

// #define ZEND_EVAL       ( 1 << 0 )

// #define ZEND_INCLUDE       ( 1 << 1 )

// #define ZEND_INCLUDE_ONCE       ( 1 << 2 )

// #define ZEND_REQUIRE       ( 1 << 3 )

// #define ZEND_REQUIRE_ONCE       ( 1 << 4 )

/* global/local fetches */

// #define ZEND_FETCH_GLOBAL       ( 1 << 1 )

// #define ZEND_FETCH_LOCAL       ( 1 << 2 )

// #define ZEND_FETCH_GLOBAL_LOCK       ( 1 << 3 )

// #define ZEND_FETCH_TYPE_MASK       0xe

/* Only one of these can ever be in use */

// #define ZEND_FETCH_REF       1

// #define ZEND_FETCH_DIM_WRITE       2

// #define ZEND_FETCH_OBJ_WRITE       3

// #define ZEND_FETCH_OBJ_FLAGS       3

// #define ZEND_ISEMPTY       ( 1 << 0 )

// #define ZEND_LAST_CATCH       ( 1 << 0 )

// #define ZEND_FREE_ON_RETURN       ( 1 << 0 )

// #define ZEND_FREE_SWITCH       ( 1 << 1 )

// #define ZEND_SEND_BY_VAL       0u

// #define ZEND_SEND_BY_REF       1u

// #define ZEND_SEND_PREFER_REF       2u

// #define ZEND_DIM_IS       ( 1 << 0 )

// #define ZEND_DIM_ALTERNATIVE_SYNTAX       ( 1 << 1 )

// #define IS_CONSTANT_UNQUALIFIED       0x010

// #define IS_CONSTANT_CLASS       0x080

// #define IS_CONSTANT_IN_NAMESPACE       0x100

func ZendCheckArgSendType(zf *ZendFunction, arg_num uint32, mask uint32) int {
	arg_num--
	if arg_num >= zf.GetNumArgs() {
		if (zf.GetFnFlags() & 1 << 14) == 0 {
			return 0
		}
		arg_num = zf.GetNumArgs()
	}
	return (zf.GetArgInfo()[arg_num].GetPassByReference() & mask) != 0
}

// #define ARG_MUST_BE_SENT_BY_REF(zf,arg_num) zend_check_arg_send_type ( zf , arg_num , ZEND_SEND_BY_REF )

// #define ARG_SHOULD_BE_SENT_BY_REF(zf,arg_num) zend_check_arg_send_type ( zf , arg_num , ZEND_SEND_BY_REF | ZEND_SEND_PREFER_REF )

// #define ARG_MAY_BE_SENT_BY_REF(zf,arg_num) zend_check_arg_send_type ( zf , arg_num , ZEND_SEND_PREFER_REF )

/* Quick API to check first 12 arguments */

// #define MAX_ARG_FLAG_NUM       12

// #define ZEND_SET_ARG_FLAG(zf,arg_num,mask) do { ( zf ) -> quick_arg_flags |= ( ( ( mask ) << 6 ) << ( arg_num ) * 2 ) ; } while ( 0 )

// #define ZEND_CHECK_ARG_FLAG(zf,arg_num,mask) ( ( ( zf ) -> quick_arg_flags >> ( ( ( arg_num ) + 3 ) * 2 ) ) & ( mask ) )

// #define QUICK_ARG_MUST_BE_SENT_BY_REF(zf,arg_num) ZEND_CHECK_ARG_FLAG ( zf , arg_num , ZEND_SEND_BY_REF )

// #define QUICK_ARG_SHOULD_BE_SENT_BY_REF(zf,arg_num) ZEND_CHECK_ARG_FLAG ( zf , arg_num , ZEND_SEND_BY_REF | ZEND_SEND_PREFER_REF )

// #define QUICK_ARG_MAY_BE_SENT_BY_REF(zf,arg_num) ZEND_CHECK_ARG_FLAG ( zf , arg_num , ZEND_SEND_PREFER_REF )

// #define ZEND_RETURN_VAL       0

// #define ZEND_RETURN_REF       1

// #define ZEND_BIND_VAL       0

// #define ZEND_BIND_REF       1

// #define ZEND_BIND_IMPLICIT       2

// #define ZEND_RETURNS_FUNCTION       ( 1 << 0 )

// #define ZEND_RETURNS_VALUE       ( 1 << 1 )

// #define ZEND_ARRAY_ELEMENT_REF       ( 1 << 0 )

// #define ZEND_ARRAY_NOT_PACKED       ( 1 << 1 )

// #define ZEND_ARRAY_SIZE_SHIFT       2

/* Attribute for ternary inside parentheses */

// #define ZEND_PARENTHESIZED_CONDITIONAL       1

/* For "use" AST nodes and the seen symbol table */

// #define ZEND_SYMBOL_CLASS       ( 1 << 0 )

// #define ZEND_SYMBOL_FUNCTION       ( 1 << 1 )

// #define ZEND_SYMBOL_CONST       ( 1 << 2 )

/* All increment opcodes are even (decrement are odd) */

// #define ZEND_IS_INCREMENT(opcode) ( ( ( opcode ) & 1 ) == 0 )

// #define ZEND_IS_BINARY_ASSIGN_OP_OPCODE(opcode) ( ( ( opcode ) >= ZEND_ADD ) && ( ( opcode ) <= ZEND_POW ) )

/* Pseudo-opcodes that are used only temporarily during compilation */

// #define ZEND_PARENTHESIZED_CONCAT       252

// #define ZEND_GOTO       253

// #define ZEND_BRK       254

// #define ZEND_CONT       255

// #define ZEND_CLONE_FUNC_NAME       "__clone"

// #define ZEND_CONSTRUCTOR_FUNC_NAME       "__construct"

// #define ZEND_DESTRUCTOR_FUNC_NAME       "__destruct"

// #define ZEND_GET_FUNC_NAME       "__get"

// #define ZEND_SET_FUNC_NAME       "__set"

// #define ZEND_UNSET_FUNC_NAME       "__unset"

// #define ZEND_ISSET_FUNC_NAME       "__isset"

// #define ZEND_CALL_FUNC_NAME       "__call"

// #define ZEND_CALLSTATIC_FUNC_NAME       "__callstatic"

// #define ZEND_TOSTRING_FUNC_NAME       "__tostring"

// #define ZEND_AUTOLOAD_FUNC_NAME       "__autoload"

// #define ZEND_INVOKE_FUNC_NAME       "__invoke"

// #define ZEND_DEBUGINFO_FUNC_NAME       "__debuginfo"

/* The following constants may be combined in CG(compiler_options)
 * to change the default compiler behavior */

// #define ZEND_COMPILE_EXTENDED_STMT       ( 1 << 0 )

// #define ZEND_COMPILE_EXTENDED_FCALL       ( 1 << 1 )

// #define ZEND_COMPILE_EXTENDED_INFO       ( ZEND_COMPILE_EXTENDED_STMT | ZEND_COMPILE_EXTENDED_FCALL )

/* call op_array handler of extendions */

// #define ZEND_COMPILE_HANDLE_OP_ARRAY       ( 1 << 2 )

/* generate ZEND_INIT_FCALL_BY_NAME for internal functions instead of ZEND_INIT_FCALL */

// #define ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS       ( 1 << 3 )

/* don't perform early binding for classes inherited form internal ones;
 * in namespaces assume that internal class that doesn't exist at compile-time
 * may apper in run-time */

// #define ZEND_COMPILE_IGNORE_INTERNAL_CLASSES       ( 1 << 4 )

/* generate ZEND_DECLARE_CLASS_DELAYED opcode to delay early binding */

// #define ZEND_COMPILE_DELAYED_BINDING       ( 1 << 5 )

/* disable constant substitution at compile-time */

// #define ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION       ( 1 << 6 )

/* disable usage of builtin instruction for strlen() */

// #define ZEND_COMPILE_NO_BUILTIN_STRLEN       ( 1 << 7 )

/* disable substitution of persistent constants at compile-time */

// #define ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION       ( 1 << 8 )

/* generate ZEND_INIT_FCALL_BY_NAME for userland functions instead of ZEND_INIT_FCALL */

// #define ZEND_COMPILE_IGNORE_USER_FUNCTIONS       ( 1 << 9 )

/* force ZEND_ACC_USE_GUARDS for all classes */

// #define ZEND_COMPILE_GUARDS       ( 1 << 10 )

/* disable builtin special case function calls */

// #define ZEND_COMPILE_NO_BUILTINS       ( 1 << 11 )

/* result of compilation may be stored in file cache */

// #define ZEND_COMPILE_WITH_FILE_CACHE       ( 1 << 12 )

/* ignore functions and classes declared in other files */

// #define ZEND_COMPILE_IGNORE_OTHER_FILES       ( 1 << 13 )

/* this flag is set when compiler invoked by opcache_compile_file() */

// #define ZEND_COMPILE_WITHOUT_EXECUTION       ( 1 << 14 )

/* this flag is set when compiler invoked during preloading */

// #define ZEND_COMPILE_PRELOAD       ( 1 << 15 )

/* disable jumptable optimization for switch statements */

// #define ZEND_COMPILE_NO_JUMPTABLES       ( 1 << 16 )

/* this flag is set when compiler invoked during preloading in separate process */

// #define ZEND_COMPILE_PRELOAD_IN_CHILD       ( 1 << 17 )

/* The default value for CG(compiler_options) */

// #define ZEND_COMPILE_DEFAULT       ZEND_COMPILE_HANDLE_OP_ARRAY

/* The default value for CG(compiler_options) during eval() */

// #define ZEND_COMPILE_DEFAULT_FOR_EVAL       0

// Source: <Zend/zend_compile.c>

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
   |          Nikita Popov <nikic@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include < zend_language_parser . h >

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_constants.h"

// # include "zend_llist.h"

// # include "zend_API.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "zend_virtual_cwd.h"

// # include "zend_multibyte.h"

// # include "zend_language_scanner.h"

// # include "zend_inheritance.h"

// # include "zend_vm.h"

// #define SET_NODE(target,src) do { target ## _type = ( src ) -> op_type ; if ( ( src ) -> op_type == IS_CONST ) { target . constant = zend_add_literal ( & ( src ) -> u . constant ) ; } else { target = ( src ) -> u . op ; } } while ( 0 )

// #define GET_NODE(target,src) do { ( target ) -> op_type = src ## _type ; if ( ( target ) -> op_type == IS_CONST ) { ZVAL_COPY_VALUE ( & ( target ) -> u . constant , CT_CONSTANT ( src ) ) ; } else { ( target ) -> u . op = src ; } } while ( 0 )

// #define FC(member) ( CG ( file_context ) . member )

// @type ZendLoopVar struct

func ZendAllocCacheSlots(count unsigned) uint32 {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var ret uint32 = op_array.GetCacheSize()
	op_array.SetCacheSize(op_array.GetCacheSize() + count*g.SizeOf("void *"))
	return ret
}
func ZendAllocCacheSlot() uint32 { return ZendAllocCacheSlots(1) }

var CompilerGlobals ZendCompilerGlobals
var ExecutorGlobals ZendExecutorGlobals

func InitOp(op *ZendOp) {
	op.GetOp1().SetNum(0)
	op.GetOp2().SetNum(0)
	op.GetResult().SetNum(0)
	op.SetOpcode(0)
	op.SetOp1Type(0)
	op.SetOp2Type(0)
	op.SetResultType(0)
	op.SetExtendedValue(0)
	op.SetLineno(CG.GetZendLineno())
}
func GetNextOpNumber() uint32 { return CG.GetActiveOpArray().GetLast() }
func GetNextOp() *ZendOp {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var next_op_num uint32 = g.PostInc(&(op_array.GetLast()))
	var next_op *ZendOp
	if next_op_num >= CG.GetContext().GetOpcodesSize() {
		CG.GetContext().SetOpcodesSize(CG.GetContext().GetOpcodesSize() * 4)
		op_array.SetOpcodes(_erealloc(op_array.GetOpcodes(), CG.GetContext().GetOpcodesSize()*g.SizeOf("zend_op")))
	}
	next_op = &op_array.GetOpcodes()[next_op_num]
	InitOp(next_op)
	return next_op
}
func GetNextBrkContElement() *ZendBrkContElement {
	CG.GetContext().GetLastBrkCont()++
	CG.GetContext().SetBrkContArray(_erealloc(CG.GetContext().GetBrkContArray(), g.SizeOf("zend_brk_cont_element")*CG.GetContext().GetLastBrkCont()))
	return &CG.context.GetBrkContArray()[CG.GetContext().GetLastBrkCont()-1]
}
func ZendDestroyPropertyInfoInternal(zv *Zval) {
	var property_info *ZendPropertyInfo = zv.GetValue().GetPtr()
	ZendStringRelease(property_info.GetName())
	Free(property_info)
}

/* }}} */

func ZendBuildRuntimeDefinitionKey(name *ZendString, start_lineno uint32) *ZendString {
	var filename *ZendString = CG.GetActiveOpArray().GetFilename()
	var result *ZendString = ZendStrpprintf(0, "%c%s%s:%"+"u"+"$%"+PRIx32, '0', name.GetVal(), filename.GetVal(), start_lineno, g.PostInc(&(CG.GetRtdKeyCounter())))
	return ZendNewInternedString(result)
}

/* }}} */

func ZendGetUnqualifiedName(name *ZendString, result **byte, result_len *int) ZendBool {
	var ns_separator *byte = ZendMemrchr(name.GetVal(), '\\', name.GetLen())
	if ns_separator != nil {
		*result = ns_separator + 1
		*result_len = name.GetVal() + name.GetLen() - (*result)
		return 1
	}
	return 0
}

/* }}} */

// @type ReservedClassName struct
var ReservedClassNames []ReservedClassName = []ReservedClassName{{"bool", g.SizeOf("\"bool\"") - 1}, {"false", g.SizeOf("\"false\"") - 1}, {"float", g.SizeOf("\"float\"") - 1}, {"int", g.SizeOf("\"int\"") - 1}, {"null", g.SizeOf("\"null\"") - 1}, {"parent", g.SizeOf("\"parent\"") - 1}, {"self", g.SizeOf("\"self\"") - 1}, {"static", g.SizeOf("\"static\"") - 1}, {"string", g.SizeOf("\"string\"") - 1}, {"true", g.SizeOf("\"true\"") - 1}, {"void", g.SizeOf("\"void\"") - 1}, {"iterable", g.SizeOf("\"iterable\"") - 1}, {"object", g.SizeOf("\"object\"") - 1}, {nil, 0}}

func ZendIsReservedClassName(name *ZendString) ZendBool {
	var reserved *ReservedClassName = ReservedClassNames
	var uqname *byte = name.GetVal()
	var uqname_len int = name.GetLen()
	ZendGetUnqualifiedName(name, &uqname, &uqname_len)
	for ; reserved.GetName() != nil; reserved++ {
		if uqname_len == reserved.GetLen() && ZendBinaryStrcasecmp(uqname, uqname_len, reserved.GetName(), reserved.GetLen()) == 0 {
			return 1
		}
	}
	return 0
}

/* }}} */

func ZendAssertValidClassName(name *ZendString) {
	if ZendIsReservedClassName(name) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot use '%s' as class name as it is reserved", name.GetVal())
	}
}

/* }}} */

// @type BuiltinTypeInfo struct

var BuiltinTypes []BuiltinTypeInfo = []BuiltinTypeInfo{{"int", g.SizeOf("\"int\"") - 1, 4}, {"float", g.SizeOf("\"float\"") - 1, 5}, {"string", g.SizeOf("\"string\"") - 1, 6}, {"bool", g.SizeOf("\"bool\"") - 1, 16}, {"void", g.SizeOf("\"void\"") - 1, 19}, {"iterable", g.SizeOf("\"iterable\"") - 1, 18}, {"object", g.SizeOf("\"object\"") - 1, 8}, {nil, 0, 0}}

func ZendLookupBuiltinTypeByName(name *ZendString) ZendUchar {
	var info *BuiltinTypeInfo = &BuiltinTypes[0]
	for ; info.GetName() != nil; info++ {
		if name.GetLen() == info.GetNameLen() && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), info.GetName(), info.GetNameLen()) == 0 {
			return info.GetType()
		}
	}
	return 0
}

/* }}} */

func ZendOparrayContextBegin(prev_context *ZendOparrayContext) {
	*prev_context = CG.GetContext()
	CG.GetContext().SetOpcodesSize(64)
	CG.GetContext().SetVarsSize(0)
	CG.GetContext().SetLiteralsSize(0)
	CG.GetContext().SetFastCallVar(-1)
	CG.GetContext().SetTryCatchOffset(-1)
	CG.GetContext().SetCurrentBrkCont(-1)
	CG.GetContext().SetLastBrkCont(0)
	CG.GetContext().SetBrkContArray(nil)
	CG.GetContext().SetLabels(nil)
}

/* }}} */

func ZendOparrayContextEnd(prev_context *ZendOparrayContext) {
	if CG.GetContext().GetBrkContArray() != nil {
		_efree(CG.GetContext().GetBrkContArray())
		CG.GetContext().SetBrkContArray(nil)
	}
	if CG.GetContext().GetLabels() != nil {
		ZendHashDestroy(CG.GetContext().GetLabels())
		_efree(CG.GetContext().GetLabels())
		CG.GetContext().SetLabels(nil)
	}
	CG.SetContext(*prev_context)
}

/* }}} */

func ZendResetImportTables() {
	if CG.GetFileContext().GetImports() != nil {
		ZendHashDestroy(CG.GetFileContext().GetImports())
		_efree(CG.GetFileContext().GetImports())
		CG.GetFileContext().SetImports(nil)
	}
	if CG.GetFileContext().GetImportsFunction() != nil {
		ZendHashDestroy(CG.GetFileContext().GetImportsFunction())
		_efree(CG.GetFileContext().GetImportsFunction())
		CG.GetFileContext().SetImportsFunction(nil)
	}
	if CG.GetFileContext().GetImportsConst() != nil {
		ZendHashDestroy(CG.GetFileContext().GetImportsConst())
		_efree(CG.GetFileContext().GetImportsConst())
		CG.GetFileContext().SetImportsConst(nil)
	}
}

/* }}} */

func ZendEndNamespace() {
	CG.GetFileContext().SetInNamespace(0)
	ZendResetImportTables()
	if CG.GetFileContext().GetCurrentNamespace() != nil {
		ZendStringReleaseEx(CG.GetFileContext().GetCurrentNamespace(), 0)
		CG.GetFileContext().SetCurrentNamespace(nil)
	}
}

/* }}} */

func ZendFileContextBegin(prev_context *ZendFileContext) {
	*prev_context = CG.GetFileContext()
	CG.GetFileContext().SetImports(nil)
	CG.GetFileContext().SetImportsFunction(nil)
	CG.GetFileContext().SetImportsConst(nil)
	CG.GetFileContext().SetCurrentNamespace(nil)
	CG.GetFileContext().SetInNamespace(0)
	CG.GetFileContext().SetHasBracketedNamespaces(0)
	CG.GetFileContext().GetDeclarables().SetTicks(0)
	_zendHashInit(&(CG.GetFileContext().GetSeenSymbols()), 8, nil, 0)
}

/* }}} */

func ZendFileContextEnd(prev_context *ZendFileContext) {
	ZendEndNamespace()
	ZendHashDestroy(&(CG.GetFileContext().GetSeenSymbols()))
	CG.SetFileContext(*prev_context)
}

/* }}} */

func ZendInitCompilerDataStructures() {
	ZendStackInit(&CG.loop_var_stack, g.SizeOf("zend_loop_var"))
	ZendStackInit(&CG.delayed_oplines_stack, g.SizeOf("zend_op"))
	CG.SetActiveClassEntry(nil)
	CG.SetInCompilation(0)
	CG.SetSkipShebang(0)
	CG.SetEncodingDeclared(0)
	CG.SetMemoizedExprs(nil)
	CG.SetMemoizeMode(0)
}

/* }}} */

func ZendRegisterSeenSymbol(name *ZendString, kind uint32) {
	var zv *Zval = ZendHashFind(&(CG.GetFileContext().GetSeenSymbols()), name)
	if zv != nil {
		zv.GetValue().SetLval(zv.GetValue().GetLval() | kind)
	} else {
		var tmp Zval
		var __z *Zval = &tmp
		__z.GetValue().SetLval(kind)
		__z.SetTypeInfo(4)
		ZendHashAddNew(&(CG.GetFileContext().GetSeenSymbols()), name, &tmp)
	}
}
func ZendHaveSeenSymbol(name *ZendString, kind uint32) ZendBool {
	var zv *Zval = ZendHashFind(&(CG.GetFileContext().GetSeenSymbols()), name)
	return zv != nil && (zv.GetValue().GetLval()&kind) != 0
}
func FileHandleDtor(fh *ZendFileHandle) { ZendFileHandleDtor(fh) }

/* }}} */

func InitCompiler() {
	CG.SetArena(ZendArenaCreate(64 * 1024))
	CG.SetActiveOpArray(nil)
	memset(&CG.context, 0, g.SizeOf("CG ( context )"))
	ZendInitCompilerDataStructures()
	ZendInitRsrcList()
	_zendHashInit(&CG.filenames_table, 8, ZvalPtrDtor, 0)
	ZendLlistInit(&CG.open_files, g.SizeOf("zend_file_handle"), (func(any))(FileHandleDtor), 0)
	CG.SetUncleanShutdown(0)
	CG.SetDelayedVarianceObligations(nil)
	CG.SetDelayedAutoloads(nil)
}

/* }}} */

func ShutdownCompiler() {
	ZendStackDestroy(&CG.loop_var_stack)
	ZendStackDestroy(&CG.delayed_oplines_stack)
	ZendHashDestroy(&CG.filenames_table)
	ZendArenaDestroy(CG.GetArena())
	if CG.GetDelayedVarianceObligations() != nil {
		ZendHashDestroy(CG.GetDelayedVarianceObligations())
		_efree(CG.GetDelayedVarianceObligations())
		CG.SetDelayedVarianceObligations(nil)
	}
	if CG.GetDelayedAutoloads() != nil {
		ZendHashDestroy(CG.GetDelayedAutoloads())
		_efree(CG.GetDelayedAutoloads())
		CG.SetDelayedAutoloads(nil)
	}
}

/* }}} */

func ZendSetCompiledFilename(new_compiled_filename *ZendString) *ZendString {
	var p *Zval
	var rv Zval
	if g.Assign(&p, ZendHashFind(&CG.filenames_table, new_compiled_filename)) {
		assert(p.GetType() == 6)
		CG.SetCompiledFilename(p.GetValue().GetStr())
		return p.GetValue().GetStr()
	}
	new_compiled_filename = ZendNewInternedString(ZendStringCopy(new_compiled_filename))
	var __z *Zval = &rv
	var __s *ZendString = new_compiled_filename
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZendHashAddNew(&CG.filenames_table, new_compiled_filename, &rv)
	CG.SetCompiledFilename(new_compiled_filename)
	return new_compiled_filename
}

/* }}} */

func ZendRestoreCompiledFilename(original_compiled_filename *ZendString) {
	CG.SetCompiledFilename(original_compiled_filename)
}

/* }}} */

func ZendGetCompiledFilename() *ZendString { return CG.GetCompiledFilename() }

/* }}} */

func ZendGetCompiledLineno() int { return CG.GetZendLineno() }

/* }}} */

func ZendIsCompiling() ZendBool { return CG.GetInCompilation() }

/* }}} */

func GetTemporaryVariable() uint32 {
	return uint32(g.PostInc(&(CG.GetActiveOpArray().GetT())))
}

/* }}} */

func LookupCv(name *ZendString) int {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var i int = 0
	var hash_value ZendUlong = ZendStringHashVal(name)
	for i < op_array.GetLastVar() {
		if op_array.GetVars()[i].GetH() == hash_value && ZendStringEquals(op_array.GetVars()[i], name) != 0 {
			return int(zend_intptr_t((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(i))))
		}
		i++
	}
	i = op_array.GetLastVar()
	op_array.GetLastVar()++
	if op_array.GetLastVar() > CG.GetContext().GetVarsSize() {
		CG.GetContext().SetVarsSize(CG.GetContext().GetVarsSize() + 16)
		op_array.SetVars(_erealloc(op_array.GetVars(), CG.GetContext().GetVarsSize()*g.SizeOf("zend_string *")))
	}
	op_array.GetVars()[i] = ZendStringCopy(name)
	return int(zend_intptr_t((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(i))))
}

/* }}} */

func ZendDelLiteral(op_array *ZendOpArray, n int) {
	ZvalPtrDtorNogc(op_array.GetLiterals() + n)
	if n+1 == op_array.GetLastLiteral() {
		op_array.GetLastLiteral()--
	} else {
		(op_array.GetLiterals() + n).SetTypeInfo(0)
	}
}

/* }}} */

/* Common part of zend_add_literal and zend_append_individual_literal */

func ZendInsertLiteral(op_array *ZendOpArray, zv *Zval, literal_position int) {
	var lit *Zval = op_array.GetLiterals() + literal_position
	if zv.GetType() == 6 {
		ZvalMakeInternedString(zv)
	}
	var _z1 *Zval = lit
	var _z2 *Zval = zv
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	lit.SetU2Extra(0)
}

/* }}} */

func ZendAddLiteral(zv *Zval) int {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var i int = op_array.GetLastLiteral()
	op_array.GetLastLiteral()++
	if i >= CG.GetContext().GetLiteralsSize() {
		for i >= CG.GetContext().GetLiteralsSize() {
			CG.GetContext().SetLiteralsSize(CG.GetContext().GetLiteralsSize() + 16)
		}
		op_array.SetLiterals((*Zval)(_erealloc(op_array.GetLiterals(), CG.GetContext().GetLiteralsSize()*g.SizeOf("zval"))))
	}
	ZendInsertLiteral(op_array, zv, i)
	return i
}

/* }}} */

func ZendAddLiteralString(str **ZendString) int {
	var ret int
	var zv Zval
	var __z *Zval = &zv
	var __s *ZendString = *str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ret = ZendAddLiteral(&zv)
	*str = zv.GetValue().GetStr()
	return ret
}

/* }}} */

func ZendAddFuncNameLiteral(name *ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *ZendString = ZendStringTolowerEx(name, 0)
	ZendAddLiteralString(&lc_name)
	return ret
}

/* }}} */

func ZendAddNsFuncNameLiteral(name *ZendString) int {
	var unqualified_name *byte
	var unqualified_name_len int

	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *ZendString = ZendStringTolowerEx(name, 0)
	ZendAddLiteralString(&lc_name)

	/* Lowercased unqualfied name */

	if ZendGetUnqualifiedName(name, &unqualified_name, &unqualified_name_len) != 0 {
		lc_name = ZendStringAlloc(unqualified_name_len, 0)
		ZendStrTolowerCopy(lc_name.GetVal(), unqualified_name, unqualified_name_len)
		ZendAddLiteralString(&lc_name)
	}
	return ret
}

/* }}} */

func ZendAddClassNameLiteral(name *ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *ZendString = ZendStringTolowerEx(name, 0)
	ZendAddLiteralString(&lc_name)
	return ret
}

/* }}} */

func ZendAddConstNameLiteral(name *ZendString, unqualified ZendBool) int {
	var tmp_name *ZendString
	var ret int = ZendAddLiteralString(&name)
	var ns_len int = 0
	var after_ns_len int = name.GetLen()
	var after_ns *byte = ZendMemrchr(name.GetVal(), '\\', name.GetLen())
	if after_ns != nil {
		after_ns += 1
		ns_len = after_ns - name.GetVal() - 1
		after_ns_len = name.GetLen() - ns_len - 1

		/* lowercased namespace name & original constant name */

		tmp_name = ZendStringInit(name.GetVal(), name.GetLen(), 0)
		ZendStrTolower(tmp_name.GetVal(), ns_len)
		ZendAddLiteralString(&tmp_name)

		/* lowercased namespace name & lowercased constant name */

		tmp_name = ZendStringTolowerEx(name, 0)
		ZendAddLiteralString(&tmp_name)
		if unqualified == 0 {
			return ret
		}
	} else {
		after_ns = name.GetVal()
	}

	/* original unqualified constant name */

	tmp_name = ZendStringInit(after_ns, after_ns_len, 0)
	ZendAddLiteralString(&tmp_name)

	/* lowercased unqualified constant name */

	tmp_name = ZendStringAlloc(after_ns_len, 0)
	ZendStrTolowerCopy(tmp_name.GetVal(), after_ns, after_ns_len)
	ZendAddLiteralString(&tmp_name)
	return ret
}

/* }}} */

// #define LITERAL_STR(op,str) do { zval _c ; ZVAL_STR ( & _c , str ) ; op . constant = zend_add_literal ( & _c ) ; } while ( 0 )

func ZendStopLexing() {
	if LANG_SCNG.GetOnEvent() != nil {
		LANG_SCNG.GetOnEvent()(ON_STOP, END, 0, LANG_SCNG.GetOnEventContext())
	}
	LANG_SCNG.SetYyCursor(LANG_SCNG.GetYyLimit())
}
func ZendBeginLoop(free_opcode ZendUchar, loop_var *Znode, is_switch ZendBool) {
	var brk_cont_element *ZendBrkContElement
	var parent int = CG.GetContext().GetCurrentBrkCont()
	var info ZendLoopVar = ZendLoopVar{0}
	CG.GetContext().SetCurrentBrkCont(CG.GetContext().GetLastBrkCont())
	brk_cont_element = GetNextBrkContElement()
	brk_cont_element.SetParent(parent)
	brk_cont_element.SetIsSwitch(is_switch)
	if loop_var != nil && (loop_var.GetOpType()&(1<<2|1<<1)) != 0 {
		var start uint32 = GetNextOpNumber()
		info.SetOpcode(free_opcode)
		info.SetVarType(loop_var.GetOpType())
		info.SetVarNum(loop_var.GetOp().GetVar())
		brk_cont_element.SetStart(start)
	} else {
		info.SetOpcode(0)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

		brk_cont_element.SetStart(-1)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

	}
	ZendStackPush(&CG.loop_var_stack, &info)
}

/* }}} */

func ZendEndLoop(cont_addr int, var_node *Znode) {
	var end uint32 = GetNextOpNumber()
	var brk_cont_element *ZendBrkContElement = &CG.context.GetBrkContArray()[CG.GetContext().GetCurrentBrkCont()]
	brk_cont_element.SetCont(cont_addr)
	brk_cont_element.SetBrk(end)
	CG.GetContext().SetCurrentBrkCont(brk_cont_element.GetParent())
	ZendStackDelTop(&CG.loop_var_stack)
}

/* }}} */

func ZendDoFree(op1 *Znode) {
	if op1.GetOpType() == 1<<1 {
		var opline *ZendOp = &CG.active_op_array.GetOpcodes()[CG.GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == 58 {
			opline--
		}
		if opline.GetResultType() == 1<<1 && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == 52 || opline.GetOpcode() == 14 {
				return
			}
		}
		ZendEmitOp(nil, 70, op1, nil)
	} else if op1.GetOpType() == 1<<2 {
		var opline *ZendOp = &CG.active_op_array.GetOpcodes()[CG.GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == 58 || opline.GetOpcode() == 103 || opline.GetOpcode() == 137 {
			opline--
		}
		if opline.GetResultType() == 1<<2 && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == 184 {
				opline.SetOpcode(0)
				opline.SetResultType(0)
			} else {
				opline.SetResultType(0)
			}
		} else {
			for opline >= CG.GetActiveOpArray().GetOpcodes() {
				if (opline.GetOpcode() == 98 || opline.GetOpcode() == 155) && opline.GetOp1Type() == 1<<2 && opline.GetOp1().GetVar() == op1.GetOp().GetVar() {
					ZendEmitOp(nil, 70, op1, nil)
					return
				}
				if opline.GetResultType() == 1<<2 && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
					if opline.GetOpcode() == 68 {
						ZendEmitOp(nil, 70, op1, nil)
					}
					break
				}
				opline--
			}
		}
	} else if op1.GetOpType() == 1<<0 {

		/* Destroy value without using GC: When opcache moves arrays into SHM it will
		 * free the zend_array structure, so references to it from outside the op array
		 * become invalid. GC would cause such a reference in the root buffer. */

		ZvalPtrDtorNogc(&op1.u.constant)

		/* Destroy value without using GC: When opcache moves arrays into SHM it will
		 * free the zend_array structure, so references to it from outside the op array
		 * become invalid. GC would cause such a reference in the root buffer. */

	}
}

/* }}} */

func ZendAddClassModifier(flags uint32, new_flag uint32) uint32 {
	var new_flags uint32 = flags | new_flag
	if (flags&1<<6) != 0 && (new_flag&1<<6) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&1<<5) != 0 && (new_flag&1<<5) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&1<<6) != 0 && (new_flags&1<<5) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class", 0)
		return 0
	}
	return new_flags
}

/* }}} */

func ZendAddMemberModifier(flags uint32, new_flag uint32) uint32 {
	var new_flags uint32 = flags | new_flag
	if (flags&(1<<0|1<<1|1<<2)) != 0 && (new_flag&(1<<0|1<<1|1<<2)) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple access type modifiers are not allowed", 0)
		return 0
	}
	if (flags&1<<6) != 0 && (new_flag&1<<6) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&1<<4) != 0 && (new_flag&1<<4) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple static modifiers are not allowed", 0)
		return 0
	}
	if (flags&1<<5) != 0 && (new_flag&1<<5) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&1<<6) != 0 && (new_flags&1<<5) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class member", 0)
		return 0
	}
	return new_flags
}

/* }}} */

func ZendConcat3(str1 *byte, str1_len int, str2 string, str2_len int, str3 *byte, str3_len int) *ZendString {
	var len_ int = str1_len + str2_len + str3_len
	var res *ZendString = ZendStringAlloc(len_, 0)
	memcpy(res.GetVal(), str1, str1_len)
	memcpy(res.GetVal()+str1_len, str2, str2_len)
	memcpy(res.GetVal()+str1_len+str2_len, str3, str3_len)
	res.GetVal()[len_] = '0'
	return res
}
func ZendConcatNames(name1 *byte, name1_len int, name2 *byte, name2_len int) *ZendString {
	return ZendConcat3(name1, name1_len, "\\", 1, name2, name2_len)
}
func ZendPrefixWithNs(name *ZendString) *ZendString {
	if CG.GetFileContext().GetCurrentNamespace() != nil {
		var ns *ZendString = CG.GetFileContext().GetCurrentNamespace()
		return ZendConcatNames(ns.GetVal(), ns.GetLen(), name.GetVal(), name.GetLen())
	} else {
		return ZendStringCopy(name)
	}
}
func ZendHashFindPtrLc(ht *HashTable, str *byte, len_ int) any {
	var result any
	var lcname *ZendString
	lcname = (*ZendString)(_emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + len_ + 1 + (8-1) & ^(8-1)))
	ZendGcSetRefcount(&lcname.gc, 1)
	lcname.GetGc().SetTypeInfo(6)
	lcname.SetH(0)
	lcname.SetLen(len_)
	ZendStrTolowerCopy(lcname.GetVal(), str, len_)
	result = ZendHashFindPtr(ht, lcname)
	_efree(lcname)
	return result
}
func ZendResolveNonClassName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool, case_sensitive ZendBool, current_import_sub *HashTable) *ZendString {
	var compound *byte
	*is_fully_qualified = 0
	if name.GetVal()[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		*is_fully_qualified = 1
		return ZendStringInit(name.GetVal()+1, name.GetLen()-1, 0)
	}
	if type_ == 0 {
		*is_fully_qualified = 1
		return ZendStringCopy(name)
	}
	if type_ == 2 {
		*is_fully_qualified = 1
		return ZendPrefixWithNs(name)
	}
	if current_import_sub != nil {

		/* If an unqualified name is a function/const alias, replace it. */

		var import_name *ZendString
		if case_sensitive != 0 {
			import_name = ZendHashFindPtr(current_import_sub, name)
		} else {
			import_name = ZendHashFindPtrLc(current_import_sub, name.GetVal(), name.GetLen())
		}
		if import_name != nil {
			*is_fully_qualified = 1
			return ZendStringCopy(import_name)
		}
	}
	compound = memchr(name.GetVal(), '\\', name.GetLen())
	if compound != nil {
		*is_fully_qualified = 1
	}
	if compound != nil && CG.GetFileContext().GetImports() != nil {

		/* If the first part of a qualified name is an alias, substitute it. */

		var len_ int = compound - name.GetVal()
		var import_name *ZendString = ZendHashFindPtrLc(CG.GetFileContext().GetImports(), name.GetVal(), len_)
		if import_name != nil {
			return ZendConcatNames(import_name.GetVal(), import_name.GetLen(), name.GetVal()+len_+1, name.GetLen()-len_-1)
		}
	}
	return ZendPrefixWithNs(name)
}

/* }}} */

func ZendResolveFunctionName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool) *ZendString {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 0, CG.GetFileContext().GetImportsFunction())
}

/* }}} */

func ZendResolveConstName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool) *ZendString {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 1, CG.GetFileContext().GetImportsConst())
}

/* }}} */

func ZendResolveClassName(name *ZendString, type_ uint32) *ZendString {
	var compound *byte
	if type_ == 2 {
		return ZendPrefixWithNs(name)
	}
	if type_ == 0 || name.GetVal()[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		if name.GetVal()[0] == '\\' {
			name = ZendStringInit(name.GetVal()+1, name.GetLen()-1, 0)
		} else {
			ZendStringAddref(name)
		}

		/* Ensure that \self, \parent and \static are not used */

		if 0 != ZendGetClassFetchType(name) {
			ZendErrorNoreturn(1<<6, "'\\%s' is an invalid class name", name.GetVal())
		}
		return name
	}
	if CG.GetFileContext().GetImports() != nil {
		compound = memchr(name.GetVal(), '\\', name.GetLen())
		if compound != nil {

			/* If the first part of a qualified name is an alias, substitute it. */

			var len_ int = compound - name.GetVal()
			var import_name *ZendString = ZendHashFindPtrLc(CG.GetFileContext().GetImports(), name.GetVal(), len_)
			if import_name != nil {
				return ZendConcatNames(import_name.GetVal(), import_name.GetLen(), name.GetVal()+len_+1, name.GetLen()-len_-1)
			}
		} else {

			/* If an unqualified name is an alias, replace it. */

			var import_name *ZendString = ZendHashFindPtrLc(CG.GetFileContext().GetImports(), name.GetVal(), name.GetLen())
			if import_name != nil {
				return ZendStringCopy(import_name)
			}
		}
	}

	/* If not fully qualified and not an alias, prepend the current namespace */

	return ZendPrefixWithNs(name)

	/* If not fully qualified and not an alias, prepend the current namespace */
}

/* }}} */

func ZendResolveClassNameAst(ast *ZendAst) *ZendString {
	var class_name *Zval = ZendAstGetZval(ast)
	if class_name.GetType() != 6 {
		ZendErrorNoreturn(1<<6, "Illegal class name")
	}
	return ZendResolveClassName(class_name.GetValue().GetStr(), ast.GetAttr())
}

/* }}} */

func LabelPtrDtor(zv *Zval) { _efree(zv.GetValue().GetPtr()) }

/* }}} */

func StrDtor(zv *Zval) {
	ZendStringReleaseEx(zv.GetValue().GetStr(), 0)
}

/* }}} */

func ZendAddTryElement(try_op uint32) uint32 {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var try_catch_offset uint32 = g.PostInc(&(op_array.GetLastTryCatch()))
	var elem *ZendTryCatchElement
	op_array.SetTryCatchArray(_safeErealloc(op_array.GetTryCatchArray(), g.SizeOf("zend_try_catch_element"), op_array.GetLastTryCatch(), 0))
	elem = &op_array.try_catch_array[try_catch_offset]
	elem.SetTryOp(try_op)
	elem.SetCatchOp(0)
	elem.SetFinallyOp(0)
	elem.SetFinallyEnd(0)
	return try_catch_offset
}

/* }}} */

func FunctionAddRef(function *ZendFunction) {
	if function.GetType() == 2 {
		var op_array *ZendOpArray = &function.op_array
		if op_array.GetRefcount() != nil {
			(*op_array).refcount++
		}
		if op_array.GetStaticVariables() != nil {
			if (ZvalGcFlags(op_array.GetStaticVariables().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendGcAddref(&(op_array.GetStaticVariables()).gc)
			}
		}
		if (CG.GetCompilerOptions() & 1 << 15) != 0 {
			assert((op_array.GetFnFlags() & 1 << 10) != 0)
			op_array.SetRunTimeCachePtr(ZendMapPtrNew())
			op_array.SetStaticVariablesPtrPtr(ZendMapPtrNew())
		} else {
			op_array.SetStaticVariablesPtrPtr(&op_array.static_variables)
			op_array.SetRunTimeCachePtr(ZendArenaAlloc(&CG.arena, g.SizeOf("void *")))
			if (uintPtr(op_array.GetRunTimeCachePtr()) & 1) != 0 {
				*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetRunTimeCachePtr()-1)))) = nil
			} else {
				*(op_array.GetRunTimeCachePtr()) = nil
			}
		}
	} else if function.GetType() == 1 {
		if function.GetFunctionName() != nil {
			ZendStringAddref(function.GetFunctionName())
		}
	}
}

/* }}} */

func DoBindFunctionError(lcname *ZendString, op_array *ZendOpArray, compile_time ZendBool) {
	var zv *Zval = ZendHashFindEx(g.CondF(compile_time != 0, func() *HashTable { return CG.GetFunctionTable() }, func() *HashTable { return EG.GetFunctionTable() }), lcname, 1)
	var error_level int = g.Cond(compile_time != 0, 1<<6, 1<<0)
	var old_function *ZendFunction
	assert(zv != nil)
	old_function = (*ZendFunction)(zv.GetValue().GetPtr())
	if old_function.GetType() == 2 && old_function.GetOpArray().GetLast() > 0 {
		ZendErrorNoreturn(error_level, "Cannot redeclare %s() (previously declared in %s:%d)", g.CondF(op_array != nil, func() []byte { return op_array.GetFunctionName().GetVal() }, func() []byte { return old_function.GetFunctionName().GetVal() }), old_function.GetOpArray().GetFilename().GetVal(), old_function.GetOpArray().GetOpcodes()[0].GetLineno())
	} else {
		ZendErrorNoreturn(error_level, "Cannot redeclare %s()", g.CondF(op_array != nil, func() []byte { return op_array.GetFunctionName().GetVal() }, func() []byte { return old_function.GetFunctionName().GetVal() }))
	}
}
func DoBindFunction(lcname *Zval) int {
	var function *ZendFunction
	var rtd_key *Zval
	var zv *Zval
	rtd_key = lcname + 1
	zv = ZendHashFindEx(EG.GetFunctionTable(), rtd_key.GetValue().GetStr(), 1)
	if zv == nil {
		DoBindFunctionError(lcname.GetValue().GetStr(), nil, 0)
		return FAILURE
	}
	function = (*ZendFunction)(zv.GetValue().GetPtr())
	if (function.GetFnFlags()&1<<10) != 0 && (CG.GetCompilerOptions()&1<<15) == 0 {
		zv = ZendHashAdd(EG.GetFunctionTable(), lcname.GetValue().GetStr(), zv)
	} else {
		zv = ZendHashSetBucketKey(EG.GetFunctionTable(), (*Bucket)(zv), lcname.GetValue().GetStr())
	}
	if zv == nil {
		DoBindFunctionError(lcname.GetValue().GetStr(), &function.op_array, 0)
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func DoBindClass(lcname *Zval, lc_parent_name *ZendString) int {
	var ce *ZendClassEntry
	var rtd_key *Zval
	var zv *Zval
	rtd_key = lcname + 1
	zv = ZendHashFindEx(EG.GetClassTable(), rtd_key.GetValue().GetStr(), 1)
	if zv == nil {
		ce = ZendHashFindPtr(EG.GetClassTable(), lcname.GetValue().GetStr())
		if ce != nil {
			ZendErrorNoreturn(1<<6, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
			return FAILURE
		} else {
			for {
				assert((EG.GetCurrentExecuteData().GetFunc().GetOpArray().GetFnFlags() & 1 << 10) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(EG.GetCurrentExecuteData().GetFunc().GetOpArray().GetFilename()) == SUCCESS {
					zv = ZendHashFindEx(EG.GetClassTable(), rtd_key.GetValue().GetStr(), 1)
					if zv != nil {
						break
					}
				}
				ZendErrorNoreturn(1<<0, "Class %s wasn't preloaded", lcname.GetValue().GetStr().GetVal())
				return FAILURE
				break
			}
		}
	}

	/* Register the derived class */

	ce = (*ZendClassEntry)(zv.GetValue().GetPtr())
	zv = ZendHashSetBucketKey(EG.GetClassTable(), (*Bucket)(zv), lcname.GetValue().GetStr())
	if zv == nil {
		ZendErrorNoreturn(1<<6, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
		return FAILURE
	}
	if ZendDoLinkClass(ce, lc_parent_name) == FAILURE {

		/* Reload bucket pointer, the hash table may have been reallocated */

		zv = ZendHashFind(EG.GetClassTable(), lcname.GetValue().GetStr())
		ZendHashSetBucketKey(EG.GetClassTable(), (*Bucket)(zv), rtd_key.GetValue().GetStr())
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func ZendMarkFunctionAsGenerator() {
	if CG.GetActiveOpArray().GetFunctionName() == nil {
		ZendErrorNoreturn(1<<6, "The \"yield\" expression can only be used inside a function")
	}
	if (CG.GetActiveOpArray().GetFnFlags() & 1 << 13) != 0 {
		var return_info ZendArgInfo = CG.GetActiveOpArray().GetArgInfo()[-1]
		if return_info.GetType()>>2 != 18 {
			var msg *byte = "Generators may only declare a return type of Generator, Iterator, Traversable, or iterable, %s is not permitted"
			if return_info.GetType() <= 0x3ff {
				ZendErrorNoreturn(1<<6, msg, ZendGetTypeByConst(return_info.GetType()>>2))
			}
			if !((*ZendString)(return_info.GetType() & ^0x3).GetLen() == g.SizeOf("\"Traversable\"")-1 && ZendBinaryStrcasecmp((*ZendString)(return_info.GetType() & ^0x3).GetVal(), (*ZendString)(return_info.GetType() & ^0x3).GetLen(), "Traversable", g.SizeOf("\"Traversable\"")-1) == 0) && !((*ZendString)(return_info.GetType() & ^0x3).GetLen() == g.SizeOf("\"Iterator\"")-1 && ZendBinaryStrcasecmp((*ZendString)(return_info.GetType() & ^0x3).GetVal(), (*ZendString)(return_info.GetType() & ^0x3).GetLen(), "Iterator", g.SizeOf("\"Iterator\"")-1) == 0) && !((*ZendString)(return_info.GetType() & ^0x3).GetLen() == g.SizeOf("\"Generator\"")-1 && ZendBinaryStrcasecmp((*ZendString)(return_info.GetType() & ^0x3).GetVal(), (*ZendString)(return_info.GetType() & ^0x3).GetLen(), "Generator", g.SizeOf("\"Generator\"")-1) == 0) {
				ZendErrorNoreturn(1<<6, msg, (*ZendString)(return_info.GetType() & ^0x3).GetVal())
			}
		}
	}
	CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<24)
}

/* }}} */

func ZendBuildDelayedEarlyBindingList(op_array *ZendOpArray) uint32 {
	if (op_array.GetFnFlags() & 1 << 16) != 0 {
		var first_early_binding_opline uint32 = uint32 - 1
		var prev_opline_num *uint32 = &first_early_binding_opline
		var opline *ZendOp = op_array.GetOpcodes()
		var end *ZendOp = opline + op_array.GetLast()
		for opline < end {
			if opline.GetOpcode() == 145 {
				*prev_opline_num = opline - op_array.GetOpcodes()
				prev_opline_num = &opline.result.GetOplineNum()
			}
			opline++
		}
		*prev_opline_num = -1
		return first_early_binding_opline
	}
	return uint32 - 1
}

/* }}} */

func ZendDoDelayedEarlyBinding(op_array *ZendOpArray, first_early_binding_opline uint32) {
	if first_early_binding_opline != uint32-1 {
		var orig_in_compilation ZendBool = CG.GetInCompilation()
		var opline_num uint32 = first_early_binding_opline
		var run_time_cache *any
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
			run_time_cache = *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(op_array).run_time_cache__ptr - 1)))
		} else {
			run_time_cache = any(*(op_array.GetRunTimeCachePtr()))
		}
		CG.SetInCompilation(1)
		for opline_num != uint32-1 {
			var opline *ZendOp = &op_array.opcodes[opline_num]
			var lcname *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
			var zv *Zval = ZendHashFindEx(EG.GetClassTable(), (lcname + 1).GetValue().GetStr(), 1)
			if zv != nil {
				var ce *ZendClassEntry = zv.GetValue().GetCe()
				var lc_parent_name *ZendString = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetStr()
				var parent_ce *ZendClassEntry = ZendHashFindExPtr(EG.GetClassTable(), lc_parent_name, 1)
				if parent_ce != nil {
					if ZendTryEarlyBind(ce, parent_ce, lcname.GetValue().GetStr(), zv) != 0 {

						/* Store in run-time cache */

						(*any)((*byte)(run_time_cache + opline.GetExtendedValue()))[0] = ce

						/* Store in run-time cache */

					}
				}
			}
			opline_num = op_array.GetOpcodes()[opline_num].GetResult().GetOplineNum()
		}
		CG.SetInCompilation(orig_in_compilation)
	}
}

/* }}} */

func ZendManglePropertyName(src1 *byte, src1_length int, src2 string, src2_length int, internal int) *ZendString {
	var prop_name_length int = 1 + src1_length + 1 + src2_length
	var prop_name *ZendString = ZendStringAlloc(prop_name_length, internal)
	prop_name.GetVal()[0] = '0'
	memcpy(prop_name.GetVal()+1, src1, src1_length+1)
	memcpy(prop_name.GetVal()+1+src1_length+1, src2, src2_length+1)
	return prop_name
}

/* }}} */

func ZendStrnlen(s *byte, maxlen int) int {
	var len_ int = 0
	for g.PostInc(&(*s)) && g.PostDec(&maxlen) {
		len_++
	}
	return len_
}

/* }}} */

func ZendUnmanglePropertyNameEx(name *ZendString, class_name **byte, prop_name **byte, prop_len *int) int {
	var class_name_len int
	var anonclass_src_len int
	*class_name = nil
	if name.GetLen() == 0 || name.GetVal()[0] != '0' {
		*prop_name = name.GetVal()
		if prop_len != nil {
			*prop_len = name.GetLen()
		}
		return SUCCESS
	}
	if name.GetLen() < 3 || name.GetVal()[1] == '0' {
		ZendError(1<<3, "Illegal member variable name")
		*prop_name = name.GetVal()
		if prop_len != nil {
			*prop_len = name.GetLen()
		}
		return FAILURE
	}
	class_name_len = ZendStrnlen(name.GetVal()+1, name.GetLen()-2)
	if class_name_len >= name.GetLen()-2 || name.GetVal()[class_name_len+1] != '0' {
		ZendError(1<<3, "Corrupt member variable name")
		*prop_name = name.GetVal()
		if prop_len != nil {
			*prop_len = name.GetLen()
		}
		return FAILURE
	}
	*class_name = name.GetVal() + 1
	anonclass_src_len = ZendStrnlen((*class_name)+class_name_len+1, name.GetLen()-class_name_len-2)
	if class_name_len+anonclass_src_len+2 != name.GetLen() {
		class_name_len += anonclass_src_len + 1
	}
	*prop_name = name.GetVal() + class_name_len + 2
	if prop_len != nil {
		*prop_len = name.GetLen() - class_name_len - 2
	}
	return SUCCESS
}

/* }}} */

func ZendLookupReservedConst(name *byte, len_ int) *ZendConstant {
	var c *ZendConstant = ZendHashFindPtrLc(EG.GetZendConstants(), name, len_)
	if c != nil && (c.GetValue().GetConstantFlags()&0xff&1<<0) == 0 && (c.GetValue().GetConstantFlags()&0xff&1<<2) != 0 {
		return c
	}
	return nil
}

/* }}} */

func ZendTryCtEvalConst(zv *Zval, name *ZendString, is_fully_qualified ZendBool) ZendBool {
	var c *ZendConstant

	/* Substitute case-sensitive (or lowercase) constants */

	c = ZendHashFindPtr(EG.GetZendConstants(), name)
	if c != nil && ((c.GetValue().GetConstantFlags()&0xff&1<<1) != 0 && (CG.GetCompilerOptions()&1<<8) == 0 && ((c.GetValue().GetConstantFlags()&0xff&1<<3) == 0 || (CG.GetCompilerOptions()&1<<12) == 0) || c.GetValue().GetType() < 8 && (CG.GetCompilerOptions()&1<<6) == 0) {
		var _z1 *Zval = zv
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
		return 1
	}

	/* Substitute true, false and null (including unqualified usage in namespaces) */

	var lookup_name *byte = name.GetVal()
	var lookup_len int = name.GetLen()
	if is_fully_qualified == 0 {
		ZendGetUnqualifiedName(name, &lookup_name, &lookup_len)
	}
	c = ZendLookupReservedConst(lookup_name, lookup_len)
	if c != nil {
		var _z1 *Zval = zv
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
		return 1
	}
	return 0
}

/* }}} */

func ZendIsScopeKnown() ZendBool {
	if (CG.GetActiveOpArray().GetFnFlags() & 1 << 20) != 0 {

		/* Closures can be rebound to a different scope */

		return 0

		/* Closures can be rebound to a different scope */

	}
	if CG.GetActiveClassEntry() == nil {

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

		return CG.GetActiveOpArray().GetFunctionName() != nil

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

	}

	/* For traits self etc refers to the using class, not the trait itself */

	return (CG.GetActiveClassEntry().GetCeFlags() & 1 << 1) == 0

	/* For traits self etc refers to the using class, not the trait itself */
}

/* }}} */

func ClassNameRefersToActiveCe(class_name *ZendString, fetch_type uint32) ZendBool {
	if CG.GetActiveClassEntry() == nil {
		return 0
	}
	if fetch_type == 1 && ZendIsScopeKnown() != 0 {
		return 1
	}
	return fetch_type == 0 && (class_name.GetLen() == CG.GetActiveClassEntry().GetName().GetLen() && ZendBinaryStrcasecmp(class_name.GetVal(), class_name.GetLen(), CG.GetActiveClassEntry().GetName().GetVal(), CG.GetActiveClassEntry().GetName().GetLen()) == 0)
}

/* }}} */

func ZendGetClassFetchType(name *ZendString) uint32 {
	if name.GetLen() == g.SizeOf("\"self\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "self", g.SizeOf("\"self\"")-1) == 0 {
		return 1
	} else if name.GetLen() == g.SizeOf("\"parent\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "parent", g.SizeOf("\"parent\"")-1) == 0 {
		return 2
	} else if name.GetLen() == g.SizeOf("\"static\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "static", g.SizeOf("\"static\"")-1) == 0 {
		return 3
	} else {
		return 0
	}
}

/* }}} */

func ZendGetClassFetchTypeAst(name_ast *ZendAst) uint32 {
	/* Fully qualified names are always default refs */

	if name_ast.GetAttr() == 0 {
		return 0
	}
	return ZendGetClassFetchType(ZendAstGetStr(name_ast))
}

/* }}} */

func ZendEnsureValidClassFetchType(fetch_type uint32) {
	if fetch_type != 0 && ZendIsScopeKnown() != 0 {
		var ce *ZendClassEntry = CG.GetActiveClassEntry()
		if ce == nil {
			ZendErrorNoreturn(1<<6, "Cannot use \"%s\" when no class scope is active", g.Cond(g.Cond(fetch_type == 1, "self", fetch_type == 2), "parent", "static"))
		} else if fetch_type == 2 && !(ce.parent_name) {
			ZendError(1<<13, "Cannot use \"parent\" when current class scope has no parent")
		}
	}
}

/* }}} */

func ZendTryCompileConstExprResolveClassName(zv *Zval, class_ast *ZendAst) ZendBool {
	var fetch_type uint32
	var class_name *Zval
	if class_ast.GetKind() != ZEND_AST_ZVAL {
		ZendErrorNoreturn(1<<6, "Cannot use ::class with dynamic class name")
	}
	class_name = ZendAstGetZval(class_ast)
	if class_name.GetType() != 6 {
		ZendErrorNoreturn(1<<6, "Illegal class name")
	}
	fetch_type = ZendGetClassFetchType(class_name.GetValue().GetStr())
	ZendEnsureValidClassFetchType(fetch_type)
	switch fetch_type {
	case 1:
		if CG.GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
			var __z *Zval = zv
			var __s *ZendString = CG.GetActiveClassEntry().GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			return 1
		}
		return 0
	case 2:
		if CG.GetActiveClassEntry() != nil && CG.GetActiveClassEntry().parent_name && ZendIsScopeKnown() != 0 {
			var __z *Zval = zv
			var __s *ZendString = CG.GetActiveClassEntry().parent_name
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			return 1
		}
		return 0
	case 3:
		return 0
	case 0:
		var __z *Zval = zv
		var __s *ZendString = ZendResolveClassNameAst(class_ast)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		return 1
	default:
		break
	}
}

/* }}} */

func ZendVerifyCtConstAccess(c *ZendClassConstant, scope *ZendClassEntry) ZendBool {
	if (c.GetValue().GetAccessFlags() & 1 << 0) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & 1 << 2) != 0 {
		return c.GetCe() == scope
	} else {
		var ce *ZendClassEntry = c.GetCe()
		for true {
			if ce == scope {
				return 1
			}
			if !(ce.parent) {
				break
			}
			if (ce.GetCeFlags() & 1 << 19) != 0 {
				ce = ce.parent
			} else {
				ce = ZendHashFindPtrLc(CG.GetClassTable(), ce.parent_name.val, ce.parent_name.len_)
				if ce == nil {
					break
				}
			}
		}

		/* Reverse case cannot be true during compilation */

		return 0

		/* Reverse case cannot be true during compilation */

	}
}
func ZendTryCtEvalClassConst(zv *Zval, class_name *ZendString, name *ZendString) ZendBool {
	var fetch_type uint32 = ZendGetClassFetchType(class_name)
	var cc *ZendClassConstant
	var c *Zval
	if ClassNameRefersToActiveCe(class_name, fetch_type) != 0 {
		cc = ZendHashFindPtr(&CG.active_class_entry.GetConstantsTable(), name)
	} else if fetch_type == 0 && (CG.GetCompilerOptions()&1<<6) == 0 {
		var ce *ZendClassEntry = ZendHashFindPtrLc(CG.GetClassTable(), class_name.GetVal(), class_name.GetLen())
		if ce != nil {
			cc = ZendHashFindPtr(&ce.constants_table, name)
		} else {
			return 0
		}
	} else {
		return 0
	}
	if (CG.GetCompilerOptions() & 1 << 8) != 0 {
		return 0
	}
	if cc == nil || ZendVerifyCtConstAccess(cc, CG.GetActiveClassEntry()) == 0 {
		return 0
	}
	c = &cc.value

	/* Substitute case-sensitive (or lowercase) persistent class constants */

	if c.GetType() < 8 {
		var _z1 *Zval = zv
		var _z2 *Zval = c
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
		return 1
	}
	return 0
}

/* }}} */

func ZendAddToList(result any, item any) {
	var list *any = *((*any)(result))
	var n int = 0
	if list != nil {
		for list[n] {
			n++
		}
	}
	list = _erealloc(list, g.SizeOf("void *")*(n+2))
	list[n] = item
	list[n+1] = nil
	*((*any)(result)) = list
}

/* }}} */

func ZendDoExtendedStmt() {
	var opline *ZendOp
	if (CG.GetCompilerOptions() & 1 << 0) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(101)
}

/* }}} */

func ZendDoExtendedFcallBegin() {
	var opline *ZendOp
	if (CG.GetCompilerOptions() & 1 << 1) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(102)
}

/* }}} */

func ZendDoExtendedFcallEnd() {
	var opline *ZendOp
	if (CG.GetCompilerOptions() & 1 << 1) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(103)
}

/* }}} */

func ZendIsAutoGlobalStr(name string, len_ int) ZendBool {
	var auto_global *ZendAutoGlobal
	if g.Assign(&auto_global, ZendHashStrFindPtr(CG.GetAutoGlobals(), name, len_)) != nil {
		if auto_global.GetArmed() != 0 {
			auto_global.SetArmed(auto_global.GetAutoGlobalCallback()(auto_global.GetName()))
		}
		return 1
	}
	return 0
}

/* }}} */

func ZendIsAutoGlobal(name *ZendString) ZendBool {
	var auto_global *ZendAutoGlobal
	if g.Assign(&auto_global, ZendHashFindPtr(CG.GetAutoGlobals(), name)) != nil {
		if auto_global.GetArmed() != 0 {
			auto_global.SetArmed(auto_global.GetAutoGlobalCallback()(auto_global.GetName()))
		}
		return 1
	}
	return 0
}

/* }}} */

func ZendRegisterAutoGlobal(name *ZendString, jit ZendBool, auto_global_callback ZendAutoGlobalCallback) int {
	var auto_global ZendAutoGlobal
	var retval int
	auto_global.SetName(name)
	auto_global.SetAutoGlobalCallback(auto_global_callback)
	auto_global.SetJit(jit)
	if ZendHashAddMem(CG.GetAutoGlobals(), auto_global.GetName(), &auto_global, g.SizeOf("zend_auto_global")) != nil {
		retval = SUCCESS
	} else {
		retval = FAILURE
	}
	return retval
}

/* }}} */

func ZendActivateAutoGlobals() {
	var auto_global *ZendAutoGlobal
	for {
		var __ht *HashTable = CG.GetAutoGlobals()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			auto_global = _z.GetValue().GetPtr()
			if auto_global.GetJit() != 0 {
				auto_global.SetArmed(1)
			} else if auto_global.GetAutoGlobalCallback() != nil {
				auto_global.SetArmed(auto_global.GetAutoGlobalCallback()(auto_global.GetName()))
			} else {
				auto_global.SetArmed(0)
			}
		}
		break
	}
}

/* }}} */

func Zendlex(elem *ZendParserStackElem) int {
	var zv Zval
	var ret int
	if CG.GetIncrementLineno() != 0 {
		CG.GetZendLineno()++
		CG.SetIncrementLineno(0)
	}
	ret = LexScan(&zv, elem)
	assert(EG.GetException() == nil || ret == T_ERROR)
	return ret
}

/* }}} */

func ZendInitializeClassData(ce *ZendClassEntry, nullify_handlers ZendBool) {
	var persistent_hashes ZendBool = ce.GetType() == 1
	ce.SetRefcount(1)
	ce.SetCeFlags(1 << 12)
	if (CG.GetCompilerOptions() & 1 << 10) != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
	}
	ce.SetDefaultPropertiesTable(nil)
	ce.SetDefaultStaticMembersTable(nil)
	_zendHashInit(&ce.properties_info, 8, g.Cond(persistent_hashes != 0, ZendDestroyPropertyInfoInternal, nil), persistent_hashes)
	_zendHashInit(&ce.constants_table, 8, nil, persistent_hashes)
	_zendHashInit(&ce.function_table, 8, ZendFunctionDtor, persistent_hashes)
	if ce.GetType() == 1 {
		ce.SetStaticMembersTablePtr(nil)
	} else {
		ce.SetStaticMembersTablePtr(&ce.default_static_members_table)
		ce.SetDocComment(nil)
	}
	ce.SetDefaultPropertiesCount(0)
	ce.SetDefaultStaticMembersCount(0)
	ce.SetPropertiesInfoTable(nil)
	if nullify_handlers != 0 {
		ce.SetConstructor(nil)
		ce.SetDestructor(nil)
		ce.SetClone(nil)
		ce.SetGet(nil)
		ce.SetSet(nil)
		ce.SetUnset(nil)
		ce.SetIsset(nil)
		ce.SetCall(nil)
		ce.SetCallstatic(nil)
		ce.SetTostring(nil)
		ce.create_object = nil
		ce.SetGetIterator(nil)
		ce.SetIteratorFuncsPtr(nil)
		ce.SetGetStaticMethod(nil)
		ce.parent = nil
		ce.parent_name = nil
		ce.SetNumInterfaces(0)
		ce.interfaces = nil
		ce.SetNumTraits(0)
		ce.SetTraitNames(nil)
		ce.SetTraitAliases(nil)
		ce.SetTraitPrecedences(nil)
		ce.SetSerialize(nil)
		ce.SetUnserialize(nil)
		ce.SetSerializeFunc(nil)
		ce.SetUnserializeFunc(nil)
		ce.SetDebugInfo(nil)
		if ce.GetType() == 1 {
			ce.SetModule(nil)
			ce.SetBuiltinFunctions(nil)
		}
	}
}

/* }}} */

func ZendGetCompiledVariableName(op_array *ZendOpArray, var_ uint32) *ZendString {
	return op_array.GetVars()[uint32((*Zval)((*byte)(nil)+int(var_))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0))))]
}

/* }}} */

func ZendAstAppendStr(left_ast *ZendAst, right_ast *ZendAst) *ZendAst {
	var left_zv *Zval = ZendAstGetZval(left_ast)
	var left *ZendString = left_zv.GetValue().GetStr()
	var right *ZendString = ZendAstGetStr(right_ast)
	var result *ZendString
	var left_len int = left.GetLen()
	var len_ int = left_len + right.GetLen() + 1
	result = ZendStringExtend(left, len_, 0)
	result.GetVal()[left_len] = '\\'
	memcpy(&result.val[left_len+1], right.GetVal(), right.GetLen())
	result.GetVal()[len_] = '0'
	ZendStringReleaseEx(right, 0)
	var __z *Zval = left_zv
	var __s *ZendString = result
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return left_ast
}

/* }}} */

func ZendNegateNumString(ast *ZendAst) *ZendAst {
	var zv *Zval = ZendAstGetZval(ast)
	if zv.GetType() == 4 {
		if zv.GetValue().GetLval() == 0 {
			var __z *Zval = zv
			var __s *ZendString = ZendStringInit("-0", g.SizeOf("\"-0\"")-1, 0)
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
		} else {
			assert(zv.GetValue().GetLval() > 0)
			zv.GetValue().SetLval(zv.GetValue().GetLval() * -1)
		}
	} else if zv.GetType() == 6 {
		var orig_len int = zv.GetValue().GetStr().GetLen()
		zv.GetValue().SetStr(ZendStringExtend(zv.GetValue().GetStr(), orig_len+1, 0))
		memmove(zv.GetValue().GetStr().GetVal()+1, zv.GetValue().GetStr().GetVal(), orig_len+1)
		zv.GetValue().GetStr().GetVal()[0] = '-'
	} else {
		assert(false)
	}
	return ast
}

/* }}} */

func ZendVerifyNamespace() {
	if CG.GetFileContext().GetHasBracketedNamespaces() != 0 && CG.GetFileContext().GetInNamespace() == 0 {
		ZendErrorNoreturn(1<<6, "No code may exist outside of namespace {}")
	}
}

/* }}} */

func ZendDirname(path *byte, len_ int) int {
	var end *byte = path + len_ - 1
	var len_adjust uint = 0
	if len_ == 0 {

		/* Illegal use of this function */

		return 0

		/* Illegal use of this function */

	}

	/* Strip trailing slashes */

	for end >= path && (*end) == '/' {
		end--
	}
	if end < path {

		/* The path only contained slashes */

		path[0] = '/'
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip filename */

	for end >= path && (*end) != '/' {
		end--
	}
	if end < path {

		/* No slash found, therefore return '.' */

		path[0] = '.'
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip slashes which came before the file name */

	for end >= path && (*end) == '/' {
		end--
	}
	if end < path {
		path[0] = '/'
		path[1] = '0'
		return 1 + len_adjust
	}
	*(end + 1) = '0'
	return size_t(end+1-path) + len_adjust
}

/* }}} */

func ZendAdjustForFetchType(opline *ZendOp, result *Znode, type_ uint32) {
	var factor ZendUchar = g.Cond(opline.GetOpcode() == 173, 1, 3)
	switch type_ {
	case 0:
		opline.SetResultType(1 << 1)
		result.SetOpType(1 << 1)
		return
	case 1:
		opline.SetOpcode(opline.GetOpcode() + 1*factor)
		return
	case 2:
		opline.SetOpcode(opline.GetOpcode() + 2*factor)
		return
	case 3:
		opline.SetResultType(1 << 1)
		result.SetOpType(1 << 1)
		opline.SetOpcode(opline.GetOpcode() + 3*factor)
		return
	case 4:
		opline.SetOpcode(opline.GetOpcode() + 4*factor)
		return
	case 5:
		opline.SetOpcode(opline.GetOpcode() + 5*factor)
		return
	default:
		break
	}
}

/* }}} */

func ZendMakeVarResult(result *Znode, opline *ZendOp) {
	opline.SetResultType(1 << 2)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == 1<<0 {
		var _z1 *Zval = &result.u.constant
		var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetResult().GetConstant()
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		result.SetOp(opline.GetResult())
	}
}

/* }}} */

func ZendMakeTmpResult(result *Znode, opline *ZendOp) {
	opline.SetResultType(1 << 1)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == 1<<0 {
		var _z1 *Zval = &result.u.constant
		var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetResult().GetConstant()
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		result.SetOp(opline.GetResult())
	}
}

/* }}} */

func ZendEmitOp(result *Znode, opcode ZendUchar, op1 *Znode, op2 *Znode) *ZendOp {
	var opline *ZendOp = GetNextOp()
	opline.SetOpcode(opcode)
	if op1 != nil {
		opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&op2.u.constant))
		} else {
			opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeVarResult(result, opline)
	}
	return opline
}

/* }}} */

func ZendEmitOpTmp(result *Znode, opcode ZendUchar, op1 *Znode, op2 *Znode) *ZendOp {
	var opline *ZendOp = GetNextOp()
	opline.SetOpcode(opcode)
	if op1 != nil {
		opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&op2.u.constant))
		} else {
			opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeTmpResult(result, opline)
	}
	return opline
}

/* }}} */

func ZendEmitTick() {
	var opline *ZendOp

	/* This prevents a double TICK generated by the parser statement of "declare()" */

	if CG.GetActiveOpArray().GetLast() != 0 && CG.GetActiveOpArray().GetOpcodes()[CG.GetActiveOpArray().GetLast()-1].GetOpcode() == 105 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(105)
	opline.SetExtendedValue(CG.GetFileContext().GetDeclarables().GetTicks())
}

/* }}} */

func ZendEmitOpData(value *Znode) *ZendOp { return ZendEmitOp(nil, 137, value, nil) }

/* }}} */

func ZendEmitJump(opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *ZendOp = ZendEmitOp(nil, 42, nil, nil)
	opline.GetOp1().SetOplineNum(opnum_target)
	return opnum
}

/* }}} */

func ZendIsSmartBranch(opline *ZendOp) int {
	switch opline.GetOpcode() {
	case 16:

	case 17:

	case 18:

	case 19:

	case 20:

	case 21:

	case 48:

	case 154:

	case 114:

	case 115:

	case 148:

	case 180:

	case 138:

	case 123:

	case 122:

	case 189:

	case 194:
		return 1
	default:
		return 0
	}
}

/* }}} */

func ZendEmitCondJump(opcode ZendUchar, cond *Znode, opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *ZendOp
	if (cond.GetOpType()&(1<<3|1<<0)) != 0 && opnum > 0 && ZendIsSmartBranch(CG.GetActiveOpArray().GetOpcodes()+opnum-1) != 0 {

		/* emit extra NOP to avoid incorrect SMART_BRANCH in very rare cases */

		ZendEmitOp(nil, 0, nil, nil)
		opnum = GetNextOpNumber()
	}
	opline = ZendEmitOp(nil, opcode, cond, nil)
	opline.GetOp2().SetOplineNum(opnum_target)
	return opnum
}

/* }}} */

func ZendUpdateJumpTarget(opnum_jump uint32, opnum_target uint32) {
	var opline *ZendOp = &CG.active_op_array.GetOpcodes()[opnum_jump]
	switch opline.GetOpcode() {
	case 42:
		opline.GetOp1().SetOplineNum(opnum_target)
		break
	case 43:

	case 44:

	case 46:

	case 47:

	case 152:

	case 169:
		opline.GetOp2().SetOplineNum(opnum_target)
		break
	default:
		break
	}
}

/* }}} */

func ZendUpdateJumpTargetToNext(opnum_jump uint32) {
	ZendUpdateJumpTarget(opnum_jump, GetNextOpNumber())
}

/* }}} */

func ZendDelayedEmitOp(result *Znode, opcode ZendUchar, op1 *Znode, op2 *Znode) *ZendOp {
	var tmp_opline ZendOp
	InitOp(&tmp_opline)
	tmp_opline.SetOpcode(opcode)
	if op1 != nil {
		tmp_opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == 1<<0 {
			tmp_opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			tmp_opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		tmp_opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == 1<<0 {
			tmp_opline.GetOp2().SetConstant(ZendAddLiteral(&op2.u.constant))
		} else {
			tmp_opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeVarResult(result, &tmp_opline)
	}
	ZendStackPush(&CG.delayed_oplines_stack, &tmp_opline)
	return ZendStackTop(&CG.delayed_oplines_stack)
}

/* }}} */

func ZendDelayedCompileBegin() uint32 {
	return ZendStackCount(&CG.delayed_oplines_stack)
}

/* }}} */

func ZendDelayedCompileEnd(offset uint32) *ZendOp {
	var opline *ZendOp = nil
	var oplines *ZendOp = ZendStackBase(&CG.delayed_oplines_stack)
	var i uint32
	var count uint32 = ZendStackCount(&CG.delayed_oplines_stack)
	assert(count >= offset)
	for i = offset; i < count; i++ {
		opline = GetNextOp()
		memcpy(opline, &oplines[i], g.SizeOf("zend_op"))
	}
	CG.GetDelayedOplinesStack().SetTop(offset)
	return opline
}

/* }}} */

// #define ZEND_MEMOIZE_NONE       0

// #define ZEND_MEMOIZE_COMPILE       1

// #define ZEND_MEMOIZE_FETCH       2

func ZendCompileMemoizedExpr(result *Znode, expr *ZendAst) {
	var memoize_mode int = CG.GetMemoizeMode()
	if memoize_mode == 1 {
		var memoized_result Znode

		/* Go through normal compilation */

		CG.SetMemoizeMode(0)
		ZendCompileExpr(result, expr)
		CG.SetMemoizeMode(1)
		if result.GetOpType() == 1<<2 {
			ZendEmitOp(&memoized_result, 167, result, nil)
		} else if result.GetOpType() == 1<<1 {
			ZendEmitOpTmp(&memoized_result, 167, result, nil)
		} else {
			if result.GetOpType() == 1<<0 {
				if &(result.GetConstant()).GetTypeFlags() != 0 {
					ZvalAddrefP(&(result.GetConstant()))
				}
			}
			memoized_result = *result
		}
		ZendHashIndexUpdateMem(CG.GetMemoizedExprs(), uintPtr(expr), &memoized_result, g.SizeOf("znode"))
	} else if memoize_mode == 2 {
		var memoized_result *Znode = ZendHashIndexFindPtr(CG.GetMemoizedExprs(), uintPtr(expr))
		*result = *memoized_result
		if result.GetOpType() == 1<<0 {
			if &(result.GetConstant()).GetTypeFlags() != 0 {
				ZvalAddrefP(&(result.GetConstant()))
			}
		}
	} else {
		assert(false)
	}
}

/* }}} */

func ZendEmitReturnTypeCheck(expr *Znode, return_info *ZendArgInfo, implicit ZendBool) {
	if return_info.GetType() > 0x3 {
		var opline *ZendOp

		/* `return ...;` is illegal in a void function (but `return;` isn't) */

		if return_info.GetType()>>2 == 19 {
			if expr != nil {
				if expr.GetOpType() == 1<<0 && expr.GetConstant().GetType() == 1 {
					ZendErrorNoreturn(1<<6, "A void function must not return a value "+"(did you mean \"return;\" instead of \"return null;\"?)")
				} else {
					ZendErrorNoreturn(1<<6, "A void function must not return a value")
				}
			}

			/* we don't need run-time check */

			return

			/* we don't need run-time check */

		}
		if expr == nil && implicit == 0 {
			if (return_info.GetType() & 0x1) != 0 {
				ZendErrorNoreturn(1<<6, "A function with return type must return a value "+"(did you mean \"return null;\" instead of \"return;\"?)")
			} else {
				ZendErrorNoreturn(1<<6, "A function with return type must return a value")
			}
		}
		if expr != nil && expr.GetOpType() == 1<<0 {
			if return_info.GetType()>>2 == expr.GetConstant().GetType() || return_info.GetType()>>2 == 16 && (expr.GetConstant().GetType() == 2 || expr.GetConstant().GetType() == 3) || (return_info.GetType()&0x1) != 0 && expr.GetConstant().GetType() == 1 {

				/* we don't need run-time check */

				return

				/* we don't need run-time check */

			}
		}
		opline = ZendEmitOp(nil, 124, expr, nil)
		if expr != nil && expr.GetOpType() == 1<<0 {
			expr.SetOpType(1 << 1)
			opline.SetResultType(expr.GetOpType())
			expr.GetOp().SetVar(GetTemporaryVariable())
			opline.GetResult().SetVar(expr.GetOp().GetVar())
		}
		if return_info.GetType() > 0x3ff {
			opline.GetOp2().SetNum(CG.GetActiveOpArray().GetCacheSize())
			CG.GetActiveOpArray().SetCacheSize(CG.GetActiveOpArray().GetCacheSize() + g.SizeOf("void *"))
		} else {
			opline.GetOp2().SetNum(-1)
		}
	}
}

/* }}} */

func ZendEmitFinalReturn(return_one int) {
	var zn Znode
	var ret *ZendOp
	var returns_reference ZendBool = (CG.GetActiveOpArray().GetFnFlags() & 1 << 12) != 0
	if (CG.GetActiveOpArray().GetFnFlags()&1<<13) != 0 && (CG.GetActiveOpArray().GetFnFlags()&1<<24) == 0 {
		ZendEmitReturnTypeCheck(nil, CG.GetActiveOpArray().GetArgInfo()-1, 1)
	}
	zn.SetOpType(1 << 0)
	if return_one != 0 {
		var __z *Zval = &zn.u.constant
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
	} else {
		&zn.u.constant.u1.type_info = 1
	}
	ret = ZendEmitOp(nil, g.Cond(returns_reference != 0, 111, 62), &zn, nil)
	ret.SetExtendedValue(-1)
}

/* }}} */

func ZendIsVariable(ast *ZendAst) ZendBool {
	return ast.GetKind() == ZEND_AST_VAR || ast.GetKind() == ZEND_AST_DIM || ast.GetKind() == ZEND_AST_PROP || ast.GetKind() == ZEND_AST_STATIC_PROP
}

/* }}} */

func ZendIsCall(ast *ZendAst) ZendBool {
	return ast.GetKind() == ZEND_AST_CALL || ast.GetKind() == ZEND_AST_METHOD_CALL || ast.GetKind() == ZEND_AST_STATIC_CALL
}

/* }}} */

func ZendIsVariableOrCall(ast *ZendAst) ZendBool {
	return ZendIsVariable(ast) != 0 || ZendIsCall(ast) != 0
}

/* }}} */

func ZendIsUntickedStmt(ast *ZendAst) ZendBool {
	return ast.GetKind() == ZEND_AST_STMT_LIST || ast.GetKind() == ZEND_AST_LABEL || ast.GetKind() == ZEND_AST_PROP_DECL || ast.GetKind() == ZEND_AST_CLASS_CONST_DECL || ast.GetKind() == ZEND_AST_USE_TRAIT || ast.GetKind() == ZEND_AST_METHOD
}

/* }}} */

func ZendCanWriteToVariable(ast *ZendAst) ZendBool {
	for ast.GetKind() == ZEND_AST_DIM || ast.GetKind() == ZEND_AST_PROP {
		ast = ast.GetChild()[0]
	}
	return ZendIsVariableOrCall(ast)
}

/* }}} */

func ZendIsConstDefaultClassRef(name_ast *ZendAst) ZendBool {
	if name_ast.GetKind() != ZEND_AST_ZVAL {
		return 0
	}
	return 0 == ZendGetClassFetchTypeAst(name_ast)
}

/* }}} */

func ZendHandleNumericOp(node *Znode) {
	if node.GetOpType() == 1<<0 && node.GetConstant().GetType() == 6 {
		var index ZendUlong
		if _zendHandleNumericStr(node.GetConstant().GetValue().GetStr().GetVal(), node.GetConstant().GetValue().GetStr().GetLen(), &index) != 0 {
			ZvalPtrDtor(&node.u.constant)
			var __z *Zval = &node.u.constant
			__z.GetValue().SetLval(index)
			__z.SetTypeInfo(4)
		}
	}
}

/* }}} */

func ZendHandleNumericDim(opline *ZendOp, dim_node *Znode) {
	if dim_node.GetConstant().GetType() == 6 {
		var index ZendUlong
		if _zendHandleNumericStr(dim_node.GetConstant().GetValue().GetStr().GetVal(), dim_node.GetConstant().GetValue().GetStr().GetLen(), &index) != 0 {

			/* For numeric indexes we also keep the original value to use by ArrayAccess
			 * See bug #63217
			 */

			var c int = ZendAddLiteral(&dim_node.u.constant)
			assert(opline.GetOp2().GetConstant()+1 == c)
			var __z *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()
			__z.GetValue().SetLval(index)
			__z.SetTypeInfo(4)
			(CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()).u2.extra = 1
			return
		}
	}
}

/* }}} */

func ZendSetClassNameOp1(opline *ZendOp, class_node *Znode) {
	if class_node.GetOpType() == 1<<0 {
		opline.SetOp1Type(1 << 0)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().GetValue().GetStr()))
	} else {
		opline.SetOp1Type(class_node.GetOpType())
		if class_node.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&class_node.u.constant))
		} else {
			opline.SetOp1(class_node.GetOp())
		}
	}
}

/* }}} */

func ZendCompileClassRef(result *Znode, name_ast *ZendAst, fetch_flags uint32) {
	var fetch_type uint32
	if name_ast.GetKind() != ZEND_AST_ZVAL {
		var name_node Znode
		ZendCompileExpr(&name_node, name_ast)
		if name_node.GetOpType() == 1<<0 {
			var name *ZendString
			if name_node.GetConstant().GetType() != 6 {
				ZendErrorNoreturn(1<<6, "Illegal class name")
			}
			name = name_node.GetConstant().GetValue().GetStr()
			fetch_type = ZendGetClassFetchType(name)
			if fetch_type == 0 {
				result.SetOpType(1 << 0)
				var __z *Zval = &result.u.constant
				var __s *ZendString = ZendResolveClassName(name, 0)
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				result.SetOpType(0)
				result.GetOp().SetNum(fetch_type | fetch_flags)
			}
			ZendStringReleaseEx(name, 0)
		} else {
			var opline *ZendOp = ZendEmitOp(result, 109, nil, &name_node)
			opline.GetOp1().SetNum(0 | fetch_flags)
		}
		return
	}

	/* Fully qualified names are always default refs */

	if name_ast.GetAttr() == 0 {
		result.SetOpType(1 << 0)
		var __z *Zval = &result.u.constant
		var __s *ZendString = ZendResolveClassNameAst(name_ast)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		return
	}
	fetch_type = ZendGetClassFetchType(ZendAstGetStr(name_ast))
	if 0 == fetch_type {
		result.SetOpType(1 << 0)
		var __z *Zval = &result.u.constant
		var __s *ZendString = ZendResolveClassNameAst(name_ast)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	} else {
		ZendEnsureValidClassFetchType(fetch_type)
		result.SetOpType(0)
		result.GetOp().SetNum(fetch_type | fetch_flags)
	}
}

/* }}} */

func ZendTryCompileCv(result *Znode, ast *ZendAst) int {
	var name_ast *ZendAst = ast.GetChild()[0]
	if name_ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(name_ast)
		var name *ZendString
		if zv.GetType() == 6 {
			name = ZvalMakeInternedString(zv)
		} else {
			name = ZendNewInternedString(ZvalGetStringFunc(zv))
		}
		if ZendIsAutoGlobal(name) != 0 {
			return FAILURE
		}
		result.SetOpType(1 << 3)
		result.GetOp().SetVar(LookupCv(name))
		if zv.GetType() != 6 {
			ZendStringReleaseEx(name, 0)
		}
		return SUCCESS
	}
	return FAILURE
}

/* }}} */

func ZendCompileSimpleVarNoCv(result *Znode, ast *ZendAst, type_ uint32, delayed int) *ZendOp {
	var name_ast *ZendAst = ast.GetChild()[0]
	var name_node Znode
	var opline *ZendOp
	ZendCompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == 1<<0 {
		if &name_node.u.constant.u1.v.type_ != 6 {
			_convertToString(&name_node.u.constant)
		}
	}
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, 80, &name_node, nil)
	} else {
		opline = ZendEmitOp(result, 80, &name_node, nil)
	}
	if name_node.GetOpType() == 1<<0 && ZendIsAutoGlobal(name_node.GetConstant().GetValue().GetStr()) != 0 {
		opline.SetExtendedValue(1 << 1)
	} else {
		opline.SetExtendedValue(1 << 2)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}

/* }}} */

func IsThisFetch(ast *ZendAst) ZendBool {
	if ast.GetKind() == ZEND_AST_VAR && ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL {
		var name *Zval = ZendAstGetZval(ast.GetChild()[0])
		return name.GetType() == 6 && (name.GetValue().GetStr().GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(name.GetValue().GetStr().GetVal(), "this", g.SizeOf("\"this\"")-1)))
	}
	return 0
}

/* }}} */

func ZendCompileSimpleVar(result *Znode, ast *ZendAst, type_ uint32, delayed int) *ZendOp {
	if IsThisFetch(ast) != 0 {
		var opline *ZendOp = ZendEmitOp(result, 184, nil, nil)
		if type_ == 0 || type_ == 3 {
			opline.SetResultType(1 << 1)
			result.SetOpType(1 << 1)
		}
		CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<30)
		return opline
	} else if ZendTryCompileCv(result, ast) == FAILURE {
		return ZendCompileSimpleVarNoCv(result, ast, type_, delayed)
	}
	return nil
}

/* }}} */

func ZendSeparateIfCallAndWrite(node *Znode, ast *ZendAst, type_ uint32) {
	if type_ != 0 && type_ != 3 && ZendIsCall(ast) != 0 {
		if node.GetOpType() == 1<<2 {
			var opline *ZendOp = ZendEmitOp(nil, 156, node, nil)
			opline.SetResultType(1 << 2)
			opline.GetResult().SetVar(opline.GetOp1().GetVar())
		} else {
			ZendErrorNoreturn(1<<6, "Cannot use result of built-in function in write context")
		}
	}
}

/* }}} */

func ZendEmitAssignZnode(var_ast *ZendAst, value_node *Znode) {
	var dummy_node Znode
	var assign_ast *ZendAst = ZendAstCreate2(ZEND_AST_ASSIGN, var_ast, ZendAstCreateZnode(value_node))
	ZendCompileAssign(&dummy_node, assign_ast)
	ZendDoFree(&dummy_node)
}

/* }}} */

func ZendDelayedCompileDim(result *Znode, ast *ZendAst, type_ uint32) *ZendOp {
	if ast.GetAttr() == 1<<1 {
		ZendError(1<<13, "Array and string offset access syntax with curly braces is deprecated")
	}
	var var_ast *ZendAst = ast.GetChild()[0]
	var dim_ast *ZendAst = ast.GetChild()[1]
	var opline *ZendOp
	var var_node Znode
	var dim_node Znode
	opline = ZendDelayedCompileVar(&var_node, var_ast, type_, 0)
	if opline != nil && type_ == 1 && (opline.GetOpcode() == 174 || opline.GetOpcode() == 85) {
		opline.SetExtendedValue(opline.GetExtendedValue() | 2)
	}
	ZendSeparateIfCallAndWrite(&var_node, var_ast, type_)
	if dim_ast == nil {
		if type_ == 0 || type_ == 3 {
			ZendErrorNoreturn(1<<6, "Cannot use [] for reading")
		}
		if type_ == 5 {
			ZendErrorNoreturn(1<<6, "Cannot use [] for unsetting")
		}
		dim_node.SetOpType(0)
	} else {
		ZendCompileExpr(&dim_node, dim_ast)
	}
	opline = ZendDelayedEmitOp(result, 81, &var_node, &dim_node)
	ZendAdjustForFetchType(opline, result, type_)
	if dim_node.GetOpType() == 1<<0 {
		ZendHandleNumericDim(opline, &dim_node)
	}
	return opline
}

/* }}} */

func ZendCompileDim(result *Znode, ast *ZendAst, type_ uint32) *ZendOp {
	var offset uint32 = ZendDelayedCompileBegin()
	ZendDelayedCompileDim(result, ast, type_)
	return ZendDelayedCompileEnd(offset)
}

/* }}} */

func ZendDelayedCompileProp(result *Znode, ast *ZendAst, type_ uint32) *ZendOp {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var prop_ast *ZendAst = ast.GetChild()[1]
	var obj_node Znode
	var prop_node Znode
	var opline *ZendOp
	if IsThisFetch(obj_ast) != 0 {
		obj_node.SetOpType(0)
		CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<30)
	} else {
		opline = ZendDelayedCompileVar(&obj_node, obj_ast, type_, 0)
		if opline != nil && type_ == 1 && (opline.GetOpcode() == 174 || opline.GetOpcode() == 85) {
			opline.SetExtendedValue(opline.GetExtendedValue() | 3)
		}
		ZendSeparateIfCallAndWrite(&obj_node, obj_ast, type_)
	}
	ZendCompileExpr(&prop_node, prop_ast)
	opline = ZendDelayedEmitOp(result, 82, &obj_node, &prop_node)
	if opline.GetOp2Type() == 1<<0 {
		if (CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()).u1.v.type_ != 6 {
			_convertToString(CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant())
		}
		opline.SetExtendedValue(ZendAllocCacheSlots(3))
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}

/* }}} */

func ZendCompileProp(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *ZendOp {
	var offset uint32 = ZendDelayedCompileBegin()
	var opline *ZendOp = ZendDelayedCompileProp(result, ast, type_)
	if by_ref != 0 {
		opline.SetExtendedValue(opline.GetExtendedValue() | 1)
	}
	return ZendDelayedCompileEnd(offset)
}

/* }}} */

func ZendCompileStaticProp(result *Znode, ast *ZendAst, type_ uint32, by_ref int, delayed int) *ZendOp {
	var class_ast *ZendAst = ast.GetChild()[0]
	var prop_ast *ZendAst = ast.GetChild()[1]
	var class_node Znode
	var prop_node Znode
	var opline *ZendOp
	ZendCompileClassRef(&class_node, class_ast, 0x200)
	ZendCompileExpr(&prop_node, prop_ast)
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, 173, &prop_node, nil)
	} else {
		opline = ZendEmitOp(result, 173, &prop_node, nil)
	}
	if opline.GetOp1Type() == 1<<0 {
		if (CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant()).u1.v.type_ != 6 {
			_convertToString(CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant())
		}
		opline.SetExtendedValue(ZendAllocCacheSlots(3))
	}
	if class_node.GetOpType() == 1<<0 {
		opline.SetOp2Type(1 << 0)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().GetValue().GetStr()))
		if opline.GetOp1Type() != 1<<0 {
			opline.SetExtendedValue(ZendAllocCacheSlot())
		}
	} else {
		opline.SetOp2Type(&class_node.GetOpType())
		if &class_node.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&class_node).u.constant))
		} else {
			opline.SetOp2(&class_node.GetOp())
		}
	}
	if by_ref != 0 && (type_ == 1 || type_ == 4) {
		opline.SetExtendedValue(opline.GetExtendedValue() | 1)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}

/* }}} */

func ZendVerifyListAssignTarget(var_ast *ZendAst, old_style ZendBool) {
	if var_ast.GetKind() == ZEND_AST_ARRAY {
		if var_ast.GetAttr() == 2 {
			ZendErrorNoreturn(1<<6, "Cannot assign to array(), use [] instead")
		}
		if old_style != var_ast.GetAttr() {
			ZendErrorNoreturn(1<<6, "Cannot mix [] and list()")
		}
	} else if ZendCanWriteToVariable(var_ast) == 0 {
		ZendErrorNoreturn(1<<6, "Assignments can only happen to writable values")
	}
}

/* }}} */

/* Propagate refs used on leaf elements to the surrounding list() structures. */

func ZendPropagateListRefs(ast *ZendAst) ZendBool {
	var list *ZendAstList = ZendAstGetList(ast)
	var has_refs ZendBool = 0
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		if elem_ast != nil {
			var var_ast *ZendAst = elem_ast.GetChild()[0]
			if var_ast.GetKind() == ZEND_AST_ARRAY {
				elem_ast.SetAttr(ZendPropagateListRefs(var_ast))
			}
			has_refs |= elem_ast.GetAttr()
		}
	}
	return has_refs
}

/* }}} */

func ZendCompileListAssign(result *Znode, ast *ZendAst, expr_node *Znode, old_style ZendBool) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var has_elems ZendBool = 0
	var is_keyed ZendBool = list.GetChildren() > 0 && list.GetChild()[0] != nil && list.GetChild()[0].GetChild()[1] != nil
	if list.GetChildren() != 0 && expr_node.GetOpType() == 1<<0 && expr_node.GetConstant().GetType() == 6 {
		ZvalMakeInternedString(&expr_node.u.constant)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var var_ast *ZendAst
		var key_ast *ZendAst
		var fetch_result Znode
		var dim_node Znode
		var opline *ZendOp
		if elem_ast == nil {
			if is_keyed != 0 {
				ZendError(1<<6, "Cannot use empty array entries in keyed array assignment")
			} else {
				continue
			}
		}
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			ZendError(1<<6, "Spread operator is not supported in assignments")
		}
		var_ast = elem_ast.GetChild()[0]
		key_ast = elem_ast.GetChild()[1]
		has_elems = 1
		if is_keyed != 0 {
			if key_ast == nil {
				ZendError(1<<6, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			ZendCompileExpr(&dim_node, key_ast)
		} else {
			if key_ast != nil {
				ZendError(1<<6, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			dim_node.SetOpType(1 << 0)
			var __z *Zval = &dim_node.u.constant
			__z.GetValue().SetLval(i)
			__z.SetTypeInfo(4)
		}
		if expr_node.GetOpType() == 1<<0 {
			if &(expr_node.GetConstant()).GetTypeFlags() != 0 {
				ZvalAddrefP(&(expr_node.GetConstant()))
			}
		}
		ZendVerifyListAssignTarget(var_ast, old_style)
		opline = ZendEmitOp(&fetch_result, g.CondF1(elem_ast.GetAttr() != 0, func() int {
			if expr_node.GetOpType() == 1<<3 {
				return 84
			} else {
				return 155
			}
		}, 98), expr_node, &dim_node)
		if dim_node.GetOpType() == 1<<0 {
			ZendHandleNumericDim(opline, &dim_node)
		}
		if var_ast.GetKind() == ZEND_AST_ARRAY {
			if elem_ast.GetAttr() != 0 {
				ZendEmitOp(&fetch_result, 140, &fetch_result, nil)
			}
			ZendCompileListAssign(nil, var_ast, &fetch_result, var_ast.GetAttr())
		} else if elem_ast.GetAttr() != 0 {
			ZendEmitAssignRefZnode(var_ast, &fetch_result)
		} else {
			ZendEmitAssignZnode(var_ast, &fetch_result)
		}
	}
	if has_elems == 0 {
		ZendErrorNoreturn(1<<6, "Cannot use empty list")
	}
	if result != nil {
		*result = *expr_node
	} else {
		ZendDoFree(expr_node)
	}
}

/* }}} */

func ZendEnsureWritableVariable(ast *ZendAst) {
	if ast.GetKind() == ZEND_AST_CALL {
		ZendErrorNoreturn(1<<6, "Can't use function return value in write context")
	}
	if ast.GetKind() == ZEND_AST_METHOD_CALL || ast.GetKind() == ZEND_AST_STATIC_CALL {
		ZendErrorNoreturn(1<<6, "Can't use method return value in write context")
	}
}

/* }}} */

func ZendIsAssignToSelf(var_ast *ZendAst, expr_ast *ZendAst) ZendBool {
	if expr_ast.GetKind() != ZEND_AST_VAR || expr_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
		return 0
	}
	for ZendIsVariable(var_ast) != 0 && var_ast.GetKind() != ZEND_AST_VAR {
		var_ast = var_ast.GetChild()[0]
	}
	if var_ast.GetKind() != ZEND_AST_VAR || var_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
		return 0
	}
	var name1 *ZendString = ZvalGetString(ZendAstGetZval(var_ast.GetChild()[0]))
	var name2 *ZendString = ZvalGetString(ZendAstGetZval(expr_ast.GetChild()[0]))
	var result ZendBool = ZendStringEquals(name1, name2)
	ZendStringReleaseEx(name1, 0)
	ZendStringReleaseEx(name2, 0)
	return result
}

/* }}} */

func ZendCompileAssign(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var expr_ast *ZendAst = ast.GetChild()[1]
	var var_node Znode
	var expr_node Znode
	var opline *ZendOp
	var offset uint32
	if IsThisFetch(var_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(&var_node, var_ast, 1, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		ZendEmitOp(result, 22, &var_node, &expr_node)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(result, var_ast, 1, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(25)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileDim(result, var_ast, 1)
		if ZendIsAssignToSelf(var_ast, expr_ast) != 0 && IsThisFetch(expr_ast) == 0 {

			/* $a[0] = $a should evaluate the right $a first */

			var cv_node Znode
			if ZendTryCompileCv(&cv_node, expr_ast) == FAILURE {
				ZendCompileSimpleVarNoCv(&expr_node, expr_ast, 0, 0)
			} else {
				ZendEmitOpTmp(&expr_node, 31, &cv_node, nil)
			}
		} else {
			ZendCompileExpr(&expr_node, expr_ast)
		}
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(23)
		opline = ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileProp(result, var_ast, 1)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(24)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_ARRAY:
		if ZendPropagateListRefs(var_ast) != 0 {
			if ZendIsVariableOrCall(expr_ast) == 0 {
				ZendErrorNoreturn(1<<6, "Cannot assign reference to non referencable value")
			}
			ZendCompileVar(&expr_node, expr_ast, 1, 1)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

			ZendEmitOp(&expr_node, 140, &expr_node, nil)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

		} else {
			if expr_ast.GetKind() == ZEND_AST_VAR {

				/* list($a, $b) = $a should evaluate the right $a first */

				var cv_node Znode
				if ZendTryCompileCv(&cv_node, expr_ast) == FAILURE {
					ZendCompileSimpleVarNoCv(&expr_node, expr_ast, 0, 0)
				} else {
					ZendEmitOpTmp(&expr_node, 31, &cv_node, nil)
				}
			} else {
				ZendCompileExpr(&expr_node, expr_ast)
			}
		}
		ZendCompileListAssign(result, var_ast, &expr_node, var_ast.GetAttr())
		return
	default:
		break
	}
}

/* }}} */

func ZendCompileAssignRef(result *Znode, ast *ZendAst) {
	var target_ast *ZendAst = ast.GetChild()[0]
	var source_ast *ZendAst = ast.GetChild()[1]
	var target_node Znode
	var source_node Znode
	var opline *ZendOp
	var offset uint32
	var flags uint32
	if IsThisFetch(target_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(target_ast)
	offset = ZendDelayedCompileBegin()
	ZendDelayedCompileVar(&target_node, target_ast, 1, 1)
	ZendCompileVar(&source_node, source_ast, 1, 1)
	if (target_ast.GetKind() != ZEND_AST_VAR || target_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL) && source_node.GetOpType() != 1<<3 {

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

		ZendEmitOp(&source_node, 140, &source_node, nil)

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

	}
	opline = ZendDelayedCompileEnd(offset)
	if source_node.GetOpType() != 1<<2 && ZendIsCall(source_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot use result of built-in function in write context")
	}
	if ZendIsCall(source_ast) != 0 {
		flags = 1 << 0
	} else {
		flags = 0
	}
	if opline != nil && opline.GetOpcode() == 85 {
		opline.SetOpcode(32)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ 1)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else if opline != nil && opline.GetOpcode() == 174 {
		opline.SetOpcode(33)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ 1)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else {
		opline = ZendEmitOp(result, 30, &target_node, &source_node)
		opline.SetExtendedValue(flags)
	}
}

/* }}} */

func ZendEmitAssignRefZnode(var_ast *ZendAst, value_node *Znode) {
	var dummy_node Znode
	var assign_ast *ZendAst = ZendAstCreate2(ZEND_AST_ASSIGN_REF, var_ast, ZendAstCreateZnode(value_node))
	ZendCompileAssignRef(&dummy_node, assign_ast)
	ZendDoFree(&dummy_node)
}

/* }}} */

func ZendCompileCompoundAssign(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var expr_ast *ZendAst = ast.GetChild()[1]
	var opcode uint32 = ast.GetAttr()
	var var_node Znode
	var expr_node Znode
	var opline *ZendOp
	var offset uint32
	var cache_slot uint32
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(&var_node, var_ast, 2, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		opline = ZendEmitOp(result, 26, &var_node, &expr_node)
		opline.SetExtendedValue(opcode)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(result, var_ast, 2, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(29)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileDim(result, var_ast, 2)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(27)
		opline.SetExtendedValue(opcode)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileProp(result, var_ast, 2)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(28)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	default:
		break
	}
}

/* }}} */

func ZendCompileArgs(ast *ZendAst, fbc *ZendFunction) uint32 {
	var args *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var uses_arg_unpack ZendBool = 0
	var arg_count uint32 = 0
	for i = 0; i < args.GetChildren(); i++ {
		var arg *ZendAst = args.GetChild()[i]
		var arg_num uint32 = i + 1
		var arg_node Znode
		var opline *ZendOp
		var opcode ZendUchar
		if arg.GetKind() == ZEND_AST_UNPACK {
			uses_arg_unpack = 1
			fbc = nil
			ZendCompileExpr(&arg_node, arg.GetChild()[0])
			opline = ZendEmitOp(nil, 165, &arg_node, nil)
			opline.GetOp2().SetNum(arg_count)
			opline.GetResult().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(arg_count)-1))))
			continue
		}
		if uses_arg_unpack != 0 {
			ZendErrorNoreturn(1<<6, "Cannot use positional argument after argument unpacking")
		}
		arg_count++
		if ZendIsVariableOrCall(arg) != 0 {
			if ZendIsCall(arg) != 0 {
				ZendCompileVar(&arg_node, arg, 0, 0)
				if (arg_node.GetOpType() & (1<<0 | 1<<1)) != 0 {

					/* Function call was converted into builtin instruction */

					if fbc == nil || ZendCheckArgSendType(fbc, arg_num, 1) != 0 {
						opcode = 116
					} else {
						opcode = 65
					}

					/* Function call was converted into builtin instruction */

				} else {
					if fbc != nil {
						if ZendCheckArgSendType(fbc, arg_num, 1) != 0 {
							opcode = 106
						} else if ZendCheckArgSendType(fbc, arg_num, 2) != 0 {
							opcode = 65
						} else {
							opcode = 117
						}
					} else {
						opcode = 50
					}
				}
			} else if fbc != nil {
				if ZendCheckArgSendType(fbc, arg_num, 1|2) != 0 {
					ZendCompileVar(&arg_node, arg, 1, 1)
					opcode = 67
				} else {
					ZendCompileVar(&arg_node, arg, 0, 0)
					if arg_node.GetOpType() == 1<<1 {
						opcode = 65
					} else {
						opcode = 117
					}
				}
			} else {
				for {
					if arg.GetKind() == ZEND_AST_VAR {
						CG.SetZendLineno(ZendAstGetLineno(ast))
						if IsThisFetch(arg) != 0 {
							ZendEmitOp(&arg_node, 184, nil, nil)
							opcode = 66
							CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<30)
							break
						} else if ZendTryCompileCv(&arg_node, arg) == SUCCESS {
							opcode = 66
							break
						}
					}
					opline = ZendEmitOp(nil, 100, nil, nil)
					opline.GetOp2().SetNum(arg_num)
					ZendCompileVar(&arg_node, arg, 4, 1)
					opcode = 185
					break
				}
			}
		} else {
			ZendCompileExpr(&arg_node, arg)
			if arg_node.GetOpType() == 1<<2 {

				/* pass ++$a or something similar */

				if fbc != nil {
					if ZendCheckArgSendType(fbc, arg_num, 1) != 0 {
						opcode = 106
					} else if ZendCheckArgSendType(fbc, arg_num, 2) != 0 {
						opcode = 65
					} else {
						opcode = 117
					}
				} else {
					opcode = 50
				}

				/* pass ++$a or something similar */

			} else if arg_node.GetOpType() == 1<<3 {
				if fbc != nil {
					if ZendCheckArgSendType(fbc, arg_num, 1|2) != 0 {
						opcode = 67
					} else {
						opcode = 117
					}
				} else {
					opcode = 66
				}
			} else {
				if fbc != nil {
					opcode = 65
					if ZendCheckArgSendType(fbc, arg_num, 1) != 0 {
						ZendErrorNoreturn(1<<6, "Only variables can be passed by reference")
					}
				} else {
					opcode = 116
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, &arg_node, nil)
		opline.GetOp2().SetOplineNum(arg_num)
		opline.GetResult().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(arg_num)-1))))
	}
	return arg_count
}

/* }}} */

func ZendGetCallOp(init_op *ZendOp, fbc *ZendFunction) ZendUchar {
	if fbc != nil {
		if fbc.GetType() == 1 && (CG.GetCompilerOptions()&1<<3) == 0 {
			if init_op.GetOpcode() == 61 && ZendExecuteInternal == nil {
				if (fbc.GetFnFlags() & (1<<6 | 1<<11 | 1<<8 | 1<<12)) == 0 {
					return 129
				} else {
					return 131
				}
			}
		} else if (CG.GetCompilerOptions() & 1 << 9) == 0 {
			if ZendExecuteEx == ExecuteEx && (fbc.GetFnFlags()&1<<6) == 0 {
				return 130
			}
		}
	} else if ZendExecuteEx == ExecuteEx && ZendExecuteInternal == nil && (init_op.GetOpcode() == 59 || init_op.GetOpcode() == 69) {
		return 131
	}
	return 60
}

/* }}} */

func ZendCompileCallCommon(result *Znode, args_ast *ZendAst, fbc *ZendFunction) {
	var opline *ZendOp
	var opnum_init uint32 = GetNextOpNumber() - 1
	var arg_count uint32
	arg_count = ZendCompileArgs(args_ast, fbc)
	ZendDoExtendedFcallBegin()
	opline = &CG.active_op_array.GetOpcodes()[opnum_init]
	opline.SetExtendedValue(arg_count)
	if opline.GetOpcode() == 61 {
		opline.GetOp1().SetNum(ZendVmCalcUsedStack(arg_count, fbc))
	}
	opline = ZendEmitOp(result, ZendGetCallOp(opline, fbc), nil, nil)
	ZendDoExtendedFcallEnd()
}

/* }}} */

func ZendCompileFunctionName(name_node *Znode, name_ast *ZendAst) ZendBool {
	var orig_name *ZendString = ZendAstGetStr(name_ast)
	var is_fully_qualified ZendBool
	name_node.SetOpType(1 << 0)
	var __z *Zval = &name_node.u.constant
	var __s *ZendString = ZendResolveFunctionName(orig_name, name_ast.GetAttr(), &is_fully_qualified)
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return is_fully_qualified == 0 && CG.GetFileContext().GetCurrentNamespace() != nil
}

/* }}} */

func ZendCompileNsCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	var opline *ZendOp = GetNextOp()
	opline.SetOpcode(69)
	opline.SetOp2Type(1 << 0)
	opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name_node.GetConstant().GetValue().GetStr()))
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	ZendCompileCallCommon(result, args_ast, nil)
}

/* }}} */

func ZendCompileDynamicCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	if name_node.GetOpType() == 1<<0 && name_node.GetConstant().GetType() == 6 {
		var colon *byte
		var str *ZendString = name_node.GetConstant().GetValue().GetStr()
		if g.Assign(&colon, ZendMemrchr(str.GetVal(), ':', str.GetLen())) != nil && colon > str.GetVal() && (*(colon - 1)) == ':' {
			var class *ZendString = ZendStringInit(str.GetVal(), colon-str.GetVal()-1, 0)
			var method *ZendString = ZendStringInit(colon+1, str.GetLen()-(colon-str.GetVal())-1, 0)
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(113)
			opline.SetOp1Type(1 << 0)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class))
			opline.SetOp2Type(1 << 0)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method))

			/* 2 slots, for class and method */

			opline.GetResult().SetNum(ZendAllocCacheSlots(2))
			ZvalPtrDtor(&name_node.u.constant)
		} else {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(59)
			opline.SetOp2Type(1 << 0)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(str))
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
	} else {
		ZendEmitOp(nil, 128, nil, name_node)
	}
	ZendCompileCallCommon(result, args_ast, nil)
}

/* }}} */

func ZendArgsContainUnpack(args *ZendAstList) ZendBool {
	var i uint32
	for i = 0; i < args.GetChildren(); i++ {
		if args.GetChild()[i].GetKind() == ZEND_AST_UNPACK {
			return 1
		}
	}
	return 0
}

/* }}} */

func ZendCompileFuncStrlen(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if (CG.GetCompilerOptions()&1<<7) != 0 || args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	if arg_node.GetOpType() == 1<<0 && arg_node.GetConstant().GetType() == 6 {
		result.SetOpType(1 << 0)
		var __z *Zval = &result.u.constant
		__z.GetValue().SetLval(arg_node.GetConstant().GetValue().GetStr().GetLen())
		__z.SetTypeInfo(4)
		ZvalPtrDtorStr(&arg_node.u.constant)
	} else {
		ZendEmitOpTmp(result, 121, &arg_node, nil)
	}
	return SUCCESS
}

/* }}} */

func ZendCompileFuncTypecheck(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, 123, &arg_node, nil)
	if type_ != 16 {
		opline.SetExtendedValue(1 << type_)
	} else {
		opline.SetExtendedValue(1<<2 | 1<<3)
	}
	return SUCCESS
}

/* }}} */

func ZendCompileFuncCast(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, 51, &arg_node, nil)
	opline.SetExtendedValue(type_)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncDefined(result *Znode, args *ZendAstList) int {
	var name *ZendString
	var opline *ZendOp
	if args.GetChildren() != 1 || args.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
		return FAILURE
	}
	name = ZvalGetString(ZendAstGetZval(args.GetChild()[0]))
	if ZendMemrchr(name.GetVal(), '\\', name.GetLen()) || ZendMemrchr(name.GetVal(), ':', name.GetLen()) {
		ZendStringReleaseEx(name, 0)
		return FAILURE
	}
	if ZendTryCtEvalConst(&result.u.constant, name, 0) != 0 {
		ZendStringReleaseEx(name, 0)
		ZvalPtrDtor(&result.u.constant)
		&result.u.constant.u1.type_info = 3
		result.SetOpType(1 << 0)
		return SUCCESS
	}
	opline = ZendEmitOpTmp(result, 122, nil, nil)
	opline.SetOp1Type(1 << 0)
	var _c Zval
	var __z *Zval = &_c
	var __s *ZendString = name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	opline.GetOp1().SetConstant(ZendAddLiteral(&_c))
	opline.SetExtendedValue(ZendAllocCacheSlot())

	/* Lowercase constant name in a separate literal */

	var c Zval
	var lcname *ZendString = ZendStringTolowerEx(name, 0)
	var __z *Zval = &c
	var __s *ZendString = lcname
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendAddLiteral(&c)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncChr(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0]).GetType() == 4 {
		var c ZendLong = ZendAstGetZval(args.GetChild()[0]).GetValue().GetLval() & 0xff
		result.SetOpType(1 << 0)
		var __z *Zval = &result.u.constant
		var __s *ZendString = ZendOneCharString[c]
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileFuncOrd(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0]).GetType() == 6 {
		result.SetOpType(1 << 0)
		var __z *Zval = &result.u.constant
		__z.GetValue().SetLval(uint8(ZendAstGetZval(args.GetChild()[0]).GetValue().GetStr().GetVal()[0]))
		__z.SetTypeInfo(4)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func FbcIsFinalized(fbc *ZendFunction) ZendBool {
	return (fbc.GetType()&1) != 0 || (fbc.GetFnFlags()&1<<25) != 0
}
func ZendTryCompileCtBoundInitUserFunc(name_ast *ZendAst, num_args uint32) int {
	var name *ZendString
	var lcname *ZendString
	var fbc *ZendFunction
	var opline *ZendOp
	if name_ast.GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(name_ast).GetType() != 6 {
		return FAILURE
	}
	name = ZendAstGetStr(name_ast)
	lcname = ZendStringTolowerEx(name, 0)
	fbc = ZendHashFindPtr(CG.GetFunctionTable(), lcname)
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == 1 && (CG.GetCompilerOptions()&1<<3) != 0 || fbc.GetType() == 2 && (CG.GetCompilerOptions()&1<<9) != 0 || fbc.GetType() == 2 && (CG.GetCompilerOptions()&1<<13) != 0 && fbc.GetOpArray().GetFilename() != CG.GetActiveOpArray().GetFilename() {
		ZendStringReleaseEx(lcname, 0)
		return FAILURE
	}
	opline = ZendEmitOp(nil, 61, nil, nil)
	opline.SetExtendedValue(num_args)
	opline.GetOp1().SetNum(ZendVmCalcUsedStack(num_args, fbc))
	opline.SetOp2Type(1 << 0)
	var _c Zval
	var __z *Zval = &_c
	var __s *ZendString = lcname
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	opline.GetOp2().SetConstant(ZendAddLiteral(&_c))
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	return SUCCESS
}

/* }}} */

func ZendCompileInitUserFunc(name_ast *ZendAst, num_args uint32, orig_func_name *ZendString) {
	var opline *ZendOp
	var name_node Znode
	if ZendTryCompileCtBoundInitUserFunc(name_ast, num_args) == SUCCESS {
		return
	}
	ZendCompileExpr(&name_node, name_ast)
	opline = ZendEmitOp(nil, 118, nil, &name_node)
	opline.SetOp1Type(1 << 0)
	var _c Zval
	var __z *Zval = &_c
	var __s *ZendString = ZendStringCopy(orig_func_name)
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	opline.GetOp1().SetConstant(ZendAddLiteral(&_c))
	opline.SetExtendedValue(num_args)
}

/* }}} */

func ZendCompileFuncCufa(result *Znode, args *ZendAstList, lcname *ZendString) int {
	var arg_node Znode
	if args.GetChildren() != 2 {
		return FAILURE
	}
	ZendCompileInitUserFunc(args.GetChild()[0], 0, lcname)
	if args.GetChild()[1].GetKind() == ZEND_AST_CALL && args.GetChild()[1].GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[1].GetChild()[0]).GetType() == 6 && args.GetChild()[1].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST {
		var orig_name *ZendString = ZendAstGetStr(args.GetChild()[1].GetChild()[0])
		var list *ZendAstList = ZendAstGetList(args.GetChild()[1].GetChild()[1])
		var is_fully_qualified ZendBool
		var name *ZendString = ZendResolveFunctionName(orig_name, args.GetChild()[1].GetChild()[0].GetAttr(), &is_fully_qualified)
		if name.GetLen() == g.SizeOf("\"array_slice\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "array_slice", g.SizeOf("\"array_slice\"")-1) == 0 && list.GetChildren() == 3 && list.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
			var zv *Zval = ZendAstGetZval(list.GetChild()[1])
			if zv.GetType() == 4 && zv.GetValue().GetLval() >= 0 && zv.GetValue().GetLval() <= 0x7fffffff {
				var opline *ZendOp
				var len_node Znode
				ZendCompileExpr(&arg_node, list.GetChild()[0])
				ZendCompileExpr(&len_node, list.GetChild()[2])
				opline = ZendEmitOp(nil, 119, &arg_node, &len_node)
				opline.SetExtendedValue(zv.GetValue().GetLval())
				ZendEmitOp(result, 60, nil, nil)
				ZendStringReleaseEx(name, 0)
				return SUCCESS
			}
		}
		ZendStringReleaseEx(name, 0)
	}
	ZendCompileExpr(&arg_node, args.GetChild()[1])
	ZendEmitOp(nil, 119, &arg_node, nil)
	ZendEmitOp(result, 60, nil, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncCuf(result *Znode, args *ZendAstList, lcname *ZendString) int {
	var i uint32
	if args.GetChildren() < 1 {
		return FAILURE
	}
	ZendCompileInitUserFunc(args.GetChild()[0], args.GetChildren()-1, lcname)
	for i = 1; i < args.GetChildren(); i++ {
		var arg_ast *ZendAst = args.GetChild()[i]
		var arg_node Znode
		var opline *ZendOp
		ZendCompileExpr(&arg_node, arg_ast)
		opline = ZendEmitOp(nil, 120, &arg_node, nil)
		opline.GetOp2().SetNum(i)
		opline.GetResult().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(i)-1))))
	}
	ZendEmitOp(result, 60, nil, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileAssert(result *Znode, args *ZendAstList, name *ZendString, fbc *ZendFunction) {
	if EG.GetAssertions() >= 0 {
		var name_node Znode
		var opline *ZendOp
		var check_op_number uint32 = GetNextOpNumber()
		ZendEmitOp(nil, 151, nil, nil)
		if fbc != nil && FbcIsFinalized(fbc) != 0 {
			name_node.SetOpType(1 << 0)
			var __z *Zval = &name_node.u.constant
			var __s *ZendString = name
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			opline = ZendEmitOp(nil, 61, nil, &name_node)
		} else {
			opline = ZendEmitOp(nil, 69, nil, nil)
			opline.SetOp2Type(1 << 0)
			opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name))
		}
		opline.GetResult().SetNum(ZendAllocCacheSlot())
		if args.GetChildren() == 1 && (args.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(args.GetChild()[0]).GetType() != 6) {

			/* add "assert(condition) as assertion message */

			ZendAstListAdd((*ZendAst)(args), ZendAstCreateZvalFromStr(ZendAstExport("assert(", args.GetChild()[0], ")")))

			/* add "assert(condition) as assertion message */

		}
		ZendCompileCallCommon(result, (*ZendAst)(args), fbc)
		opline = &CG.active_op_array.GetOpcodes()[check_op_number]
		opline.GetOp2().SetOplineNum(GetNextOpNumber())
		opline.SetResultType(result.GetOpType())
		if result.GetOpType() == 1<<0 {
			opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetResult(result.GetOp())
		}
	} else {
		if fbc == nil {
			ZendStringReleaseEx(name, 0)
		}
		result.SetOpType(1 << 0)
		&result.u.constant.u1.type_info = 3
	}
}

/* }}} */

func ZendCompileFuncInArray(result *Znode, args *ZendAstList) int {
	var strict ZendBool = 0
	var array Znode
	var needly Znode
	var opline *ZendOp
	if args.GetChildren() == 3 {
		if args.GetChild()[2].GetKind() == ZEND_AST_ZVAL {
			strict = ZendIsTrue(ZendAstGetZval(args.GetChild()[2]))
		} else if args.GetChild()[2].GetKind() == ZEND_AST_CONST {
			var value Zval
			var name_ast *ZendAst = args.GetChild()[2].GetChild()[0]
			var is_fully_qualified ZendBool
			var resolved_name *ZendString = ZendResolveConstName(ZendAstGetStr(name_ast), name_ast.GetAttr(), &is_fully_qualified)
			if ZendTryCtEvalConst(&value, resolved_name, is_fully_qualified) == 0 {
				ZendStringReleaseEx(resolved_name, 0)
				return FAILURE
			}
			ZendStringReleaseEx(resolved_name, 0)
			strict = ZendIsTrue(&value)
			ZvalPtrDtor(&value)
		} else {
			return FAILURE
		}
	} else if args.GetChildren() != 2 {
		return FAILURE
	}
	if args.GetChild()[1].GetKind() != ZEND_AST_ARRAY || ZendTryCtEvalArray(&array.u.constant, args.GetChild()[1]) == 0 {
		return FAILURE
	}
	if array.GetConstant().GetValue().GetArr().GetNNumOfElements() > 0 {
		var ok ZendBool = 1
		var val *Zval
		var tmp Zval
		var src *HashTable = array.GetConstant().GetValue().GetArr()
		var dst *HashTable = _zendNewArray(src.GetNNumOfElements())
		&tmp.SetTypeInfo(3)
		if strict != 0 {
			for {
				var __ht *HashTable = src
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					val = _z
					if val.GetType() == 6 {
						ZendHashAdd(dst, val.GetValue().GetStr(), &tmp)
					} else if val.GetType() == 4 {
						ZendHashIndexAdd(dst, val.GetValue().GetLval(), &tmp)
					} else {
						ZendArrayDestroy(dst)
						ok = 0
						break
					}
				}
				break
			}
		} else {
			for {
				var __ht *HashTable = src
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					val = _z
					if val.GetType() != 6 || IsNumericString(val.GetValue().GetStr().GetVal(), val.GetValue().GetStr().GetLen(), nil, nil, 0) != 0 {
						ZendArrayDestroy(dst)
						ok = 0
						break
					}
					ZendHashAdd(dst, val.GetValue().GetStr(), &tmp)
				}
				break
			}
		}
		ZendArrayDestroy(src)
		if ok == 0 {
			return FAILURE
		}
		array.GetConstant().GetValue().SetArr(dst)
	}
	array.SetOpType(1 << 0)
	ZendCompileExpr(&needly, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, 189, &needly, &array)
	opline.SetExtendedValue(strict)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncCount(result *Znode, args *ZendAstList, lcname *ZendString) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, 190, &arg_node, nil)
	opline.SetExtendedValue(lcname.GetLen() == g.SizeOf("\"sizeof\"")-1 && !(memcmp(lcname.GetVal(), "sizeof", g.SizeOf("\"sizeof\"")-1)))
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGetClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 0 {
		ZendEmitOpTmp(result, 191, nil, nil)
	} else {
		var arg_node Znode
		if args.GetChildren() != 1 {
			return FAILURE
		}
		ZendCompileExpr(&arg_node, args.GetChild()[0])
		ZendEmitOpTmp(result, 191, &arg_node, nil)
	}
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGetCalledClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() != 0 {
		return FAILURE
	}
	ZendEmitOpTmp(result, 192, nil, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGettype(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	ZendEmitOpTmp(result, 193, &arg_node, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncNumArgs(result *Znode, args *ZendAstList) int {
	if CG.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, 171, nil, nil)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileFuncGetArgs(result *Znode, args *ZendAstList) int {
	if CG.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, 172, nil, nil)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileFuncArrayKeyExists(result *Znode, args *ZendAstList) int {
	var subject Znode
	var needle Znode
	if args.GetChildren() != 2 {
		return FAILURE
	}
	ZendCompileExpr(&needle, args.GetChild()[0])
	ZendCompileExpr(&subject, args.GetChild()[1])
	ZendEmitOpTmp(result, 194, &needle, &subject)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncArraySlice(result *Znode, args *ZendAstList) int {
	if CG.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 2 && args.GetChild()[0].GetKind() == ZEND_AST_CALL && args.GetChild()[0].GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0].GetChild()[0]).GetType() == 6 && args.GetChild()[0].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST && args.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
		var orig_name *ZendString = ZendAstGetStr(args.GetChild()[0].GetChild()[0])
		var is_fully_qualified ZendBool
		var name *ZendString = ZendResolveFunctionName(orig_name, args.GetChild()[0].GetChild()[0].GetAttr(), &is_fully_qualified)
		var list *ZendAstList = ZendAstGetList(args.GetChild()[0].GetChild()[1])
		var zv *Zval = ZendAstGetZval(args.GetChild()[1])
		var first Znode
		if name.GetLen() == g.SizeOf("\"func_get_args\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "func_get_args", g.SizeOf("\"func_get_args\"")-1) == 0 && list.GetChildren() == 0 && zv.GetType() == 4 && zv.GetValue().GetLval() >= 0 {
			first.SetOpType(1 << 0)
			var __z *Zval = &first.u.constant
			__z.GetValue().SetLval(zv.GetValue().GetLval())
			__z.SetTypeInfo(4)
			ZendEmitOpTmp(result, 172, &first, nil)
			ZendStringReleaseEx(name, 0)
			return SUCCESS
		}
		ZendStringReleaseEx(name, 0)
	}
	return FAILURE
}

/* }}} */

func ZendTryCompileSpecialFunc(result *Znode, lcname *ZendString, args *ZendAstList, fbc *ZendFunction, type_ uint32) int {
	if fbc.GetInternalFunction().GetHandler() == ZifDisplayDisabledFunction {
		return FAILURE
	}
	if (CG.GetCompilerOptions() & 1 << 11) != 0 {
		return FAILURE
	}
	if ZendArgsContainUnpack(args) != 0 {
		return FAILURE
	}
	if lcname.GetLen() == g.SizeOf("\"strlen\"")-1 && !(memcmp(lcname.GetVal(), "strlen", g.SizeOf("\"strlen\"")-1)) {
		return ZendCompileFuncStrlen(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"is_null\"")-1 && !(memcmp(lcname.GetVal(), "is_null", g.SizeOf("\"is_null\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 1)
	} else if lcname.GetLen() == g.SizeOf("\"is_bool\"")-1 && !(memcmp(lcname.GetVal(), "is_bool", g.SizeOf("\"is_bool\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 16)
	} else if lcname.GetLen() == g.SizeOf("\"is_long\"")-1 && !(memcmp(lcname.GetVal(), "is_long", g.SizeOf("\"is_long\"")-1)) || lcname.GetLen() == g.SizeOf("\"is_int\"")-1 && !(memcmp(lcname.GetVal(), "is_int", g.SizeOf("\"is_int\"")-1)) || lcname.GetLen() == g.SizeOf("\"is_integer\"")-1 && !(memcmp(lcname.GetVal(), "is_integer", g.SizeOf("\"is_integer\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 4)
	} else if lcname.GetLen() == g.SizeOf("\"is_float\"")-1 && !(memcmp(lcname.GetVal(), "is_float", g.SizeOf("\"is_float\"")-1)) || lcname.GetLen() == g.SizeOf("\"is_double\"")-1 && !(memcmp(lcname.GetVal(), "is_double", g.SizeOf("\"is_double\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 5)
	} else if lcname.GetLen() == g.SizeOf("\"is_string\"")-1 && !(memcmp(lcname.GetVal(), "is_string", g.SizeOf("\"is_string\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 6)
	} else if lcname.GetLen() == g.SizeOf("\"is_array\"")-1 && !(memcmp(lcname.GetVal(), "is_array", g.SizeOf("\"is_array\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 7)
	} else if lcname.GetLen() == g.SizeOf("\"is_object\"")-1 && !(memcmp(lcname.GetVal(), "is_object", g.SizeOf("\"is_object\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 8)
	} else if lcname.GetLen() == g.SizeOf("\"is_resource\"")-1 && !(memcmp(lcname.GetVal(), "is_resource", g.SizeOf("\"is_resource\"")-1)) {
		return ZendCompileFuncTypecheck(result, args, 9)
	} else if lcname.GetLen() == g.SizeOf("\"boolval\"")-1 && !(memcmp(lcname.GetVal(), "boolval", g.SizeOf("\"boolval\"")-1)) {
		return ZendCompileFuncCast(result, args, 16)
	} else if lcname.GetLen() == g.SizeOf("\"intval\"")-1 && !(memcmp(lcname.GetVal(), "intval", g.SizeOf("\"intval\"")-1)) {
		return ZendCompileFuncCast(result, args, 4)
	} else if lcname.GetLen() == g.SizeOf("\"floatval\"")-1 && !(memcmp(lcname.GetVal(), "floatval", g.SizeOf("\"floatval\"")-1)) || lcname.GetLen() == g.SizeOf("\"doubleval\"")-1 && !(memcmp(lcname.GetVal(), "doubleval", g.SizeOf("\"doubleval\"")-1)) {
		return ZendCompileFuncCast(result, args, 5)
	} else if lcname.GetLen() == g.SizeOf("\"strval\"")-1 && !(memcmp(lcname.GetVal(), "strval", g.SizeOf("\"strval\"")-1)) {
		return ZendCompileFuncCast(result, args, 6)
	} else if lcname.GetLen() == g.SizeOf("\"defined\"")-1 && !(memcmp(lcname.GetVal(), "defined", g.SizeOf("\"defined\"")-1)) {
		return ZendCompileFuncDefined(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"chr\"")-1 && !(memcmp(lcname.GetVal(), "chr", g.SizeOf("\"chr\"")-1)) && type_ == 0 {
		return ZendCompileFuncChr(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"ord\"")-1 && !(memcmp(lcname.GetVal(), "ord", g.SizeOf("\"ord\"")-1)) && type_ == 0 {
		return ZendCompileFuncOrd(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"call_user_func_array\"")-1 && !(memcmp(lcname.GetVal(), "call_user_func_array", g.SizeOf("\"call_user_func_array\"")-1)) {
		return ZendCompileFuncCufa(result, args, lcname)
	} else if lcname.GetLen() == g.SizeOf("\"call_user_func\"")-1 && !(memcmp(lcname.GetVal(), "call_user_func", g.SizeOf("\"call_user_func\"")-1)) {
		return ZendCompileFuncCuf(result, args, lcname)
	} else if lcname.GetLen() == g.SizeOf("\"in_array\"")-1 && !(memcmp(lcname.GetVal(), "in_array", g.SizeOf("\"in_array\"")-1)) {
		return ZendCompileFuncInArray(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"count\"")-1 && !(memcmp(lcname.GetVal(), "count", g.SizeOf("\"count\"")-1)) || lcname.GetLen() == g.SizeOf("\"sizeof\"")-1 && !(memcmp(lcname.GetVal(), "sizeof", g.SizeOf("\"sizeof\"")-1)) {
		return ZendCompileFuncCount(result, args, lcname)
	} else if lcname.GetLen() == g.SizeOf("\"get_class\"")-1 && !(memcmp(lcname.GetVal(), "get_class", g.SizeOf("\"get_class\"")-1)) {
		return ZendCompileFuncGetClass(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"get_called_class\"")-1 && !(memcmp(lcname.GetVal(), "get_called_class", g.SizeOf("\"get_called_class\"")-1)) {
		return ZendCompileFuncGetCalledClass(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"gettype\"")-1 && !(memcmp(lcname.GetVal(), "gettype", g.SizeOf("\"gettype\"")-1)) {
		return ZendCompileFuncGettype(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"func_num_args\"")-1 && !(memcmp(lcname.GetVal(), "func_num_args", g.SizeOf("\"func_num_args\"")-1)) {
		return ZendCompileFuncNumArgs(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"func_get_args\"")-1 && !(memcmp(lcname.GetVal(), "func_get_args", g.SizeOf("\"func_get_args\"")-1)) {
		return ZendCompileFuncGetArgs(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"array_slice\"")-1 && !(memcmp(lcname.GetVal(), "array_slice", g.SizeOf("\"array_slice\"")-1)) {
		return ZendCompileFuncArraySlice(result, args)
	} else if lcname.GetLen() == g.SizeOf("\"array_key_exists\"")-1 && !(memcmp(lcname.GetVal(), "array_key_exists", g.SizeOf("\"array_key_exists\"")-1)) {
		return ZendCompileFuncArrayKeyExists(result, args)
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileCall(result *Znode, ast *ZendAst, type_ uint32) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var args_ast *ZendAst = ast.GetChild()[1]
	var name_node Znode
	if name_ast.GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(name_ast).GetType() != 6 {
		ZendCompileExpr(&name_node, name_ast)
		ZendCompileDynamicCall(result, &name_node, args_ast)
		return
	}
	var runtime_resolution ZendBool = ZendCompileFunctionName(&name_node, name_ast)
	if runtime_resolution != 0 {
		if ZendAstGetStr(name_ast).GetLen() == g.SizeOf("\"assert\"")-1 && ZendBinaryStrcasecmp(ZendAstGetStr(name_ast).GetVal(), ZendAstGetStr(name_ast).GetLen(), "assert", g.SizeOf("\"assert\"")-1) == 0 {
			ZendCompileAssert(result, ZendAstGetList(args_ast), name_node.GetConstant().GetValue().GetStr(), nil)
		} else {
			ZendCompileNsCall(result, &name_node, args_ast)
		}
		return
	}
	var name *Zval = &name_node.u.constant
	var lcname *ZendString
	var fbc *ZendFunction
	var opline *ZendOp
	lcname = ZendStringTolowerEx(name.GetValue().GetStr(), 0)
	fbc = ZendHashFindPtr(CG.GetFunctionTable(), lcname)

	/* Special assert() handling should apply independently of compiler flags. */

	if fbc != nil && (lcname.GetLen() == g.SizeOf("\"assert\"")-1 && !(memcmp(lcname.GetVal(), "assert", g.SizeOf("\"assert\"")-1))) {
		ZendCompileAssert(result, ZendAstGetList(args_ast), lcname, fbc)
		ZendStringRelease(lcname)
		ZvalPtrDtor(&name_node.u.constant)
		return
	}
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == 1 && (CG.GetCompilerOptions()&1<<3) != 0 || fbc.GetType() == 2 && (CG.GetCompilerOptions()&1<<9) != 0 || fbc.GetType() == 2 && (CG.GetCompilerOptions()&1<<13) != 0 && fbc.GetOpArray().GetFilename() != CG.GetActiveOpArray().GetFilename() {
		ZendStringReleaseEx(lcname, 0)
		ZendCompileDynamicCall(result, &name_node, args_ast)
		return
	}
	if ZendTryCompileSpecialFunc(result, lcname, ZendAstGetList(args_ast), fbc, type_) == SUCCESS {
		ZendStringReleaseEx(lcname, 0)
		ZvalPtrDtor(&name_node.u.constant)
		return
	}
	ZvalPtrDtor(&name_node.u.constant)
	var __z *Zval = &name_node.u.constant
	var __s *ZendString = lcname
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	opline = ZendEmitOp(nil, 61, nil, &name_node)
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	ZendCompileCallCommon(result, args_ast, fbc)
}

/* }}} */

func ZendCompileMethodCall(result *Znode, ast *ZendAst, type_ uint32) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	var args_ast *ZendAst = ast.GetChild()[2]
	var obj_node Znode
	var method_node Znode
	var opline *ZendOp
	var fbc *ZendFunction = nil
	if IsThisFetch(obj_ast) != 0 {
		obj_node.SetOpType(0)
		CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<30)
	} else {
		ZendCompileExpr(&obj_node, obj_ast)
	}
	ZendCompileExpr(&method_node, method_ast)
	opline = ZendEmitOp(nil, 112, &obj_node, nil)
	if method_node.GetOpType() == 1<<0 {
		if method_node.GetConstant().GetType() != 6 {
			ZendErrorNoreturn(1<<6, "Method name must be a string")
		}
		opline.SetOp2Type(1 << 0)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().GetValue().GetStr()))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		opline.SetOp2Type(&method_node.GetOpType())
		if &method_node.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&method_node).u.constant))
		} else {
			opline.SetOp2(&method_node.GetOp())
		}
	}

	/* Check if this calls a known method on $this */

	if opline.GetOp1Type() == 0 && opline.GetOp2Type() == 1<<0 && CG.GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
		var lcname *ZendString = (CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant() + 1).value.str
		fbc = ZendHashFindPtr(&CG.active_class_entry.GetFunctionTable(), lcname)

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

		if fbc != nil && (fbc.GetFnFlags()&(1<<2|1<<5)) == 0 {
			fbc = nil
		}

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

	}
	ZendCompileCallCommon(result, args_ast, fbc)
}

/* }}} */

func ZendIsConstructor(name *ZendString) ZendBool {
	return name.GetLen() == g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "__construct", g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1) == 0
}

/* }}} */

func ZendCompileStaticCall(result *Znode, ast *ZendAst, type_ uint32) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	var args_ast *ZendAst = ast.GetChild()[2]
	var class_node Znode
	var method_node Znode
	var opline *ZendOp
	var fbc *ZendFunction = nil
	ZendCompileClassRef(&class_node, class_ast, 0x200)
	ZendCompileExpr(&method_node, method_ast)
	if method_node.GetOpType() == 1<<0 {
		var name *Zval = &method_node.u.constant
		if name.GetType() != 6 {
			ZendErrorNoreturn(1<<6, "Method name must be a string")
		}
		if ZendIsConstructor(name.GetValue().GetStr()) != 0 {
			ZvalPtrDtor(name)
			method_node.SetOpType(0)
		}
	}
	opline = GetNextOp()
	opline.SetOpcode(113)
	ZendSetClassNameOp1(opline, &class_node)
	if method_node.GetOpType() == 1<<0 {
		opline.SetOp2Type(1 << 0)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().GetValue().GetStr()))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		if opline.GetOp1Type() == 1<<0 {
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
		opline.SetOp2Type(&method_node.GetOpType())
		if &method_node.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&method_node).u.constant))
		} else {
			opline.SetOp2(&method_node.GetOp())
		}
	}

	/* Check if we already know which method we're calling */

	if opline.GetOp2Type() == 1<<0 {
		var ce *ZendClassEntry = nil
		if opline.GetOp1Type() == 1<<0 {
			var lcname *ZendString = (CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant() + 1).value.str
			ce = ZendHashFindPtr(CG.GetClassTable(), lcname)
			if ce == nil && CG.GetActiveClassEntry() != nil && (CG.GetActiveClassEntry().GetName().GetLen() == lcname.GetLen() && ZendBinaryStrcasecmp(CG.GetActiveClassEntry().GetName().GetVal(), CG.GetActiveClassEntry().GetName().GetLen(), lcname.GetVal(), lcname.GetLen()) == 0) {
				ce = CG.GetActiveClassEntry()
			}
		} else if opline.GetOp1Type() == 0 && (opline.GetOp1().GetNum()&0xf) == 1 && ZendIsScopeKnown() != 0 {
			ce = CG.GetActiveClassEntry()
		}
		if ce != nil {
			var lcname *ZendString = (CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant() + 1).value.str
			fbc = ZendHashFindPtr(&ce.function_table, lcname)
			if fbc != nil && (fbc.GetFnFlags()&1<<0) == 0 {
				if ce != CG.GetActiveClassEntry() && ((fbc.GetFnFlags()&1<<2) != 0 || (fbc.GetScope().GetCeFlags()&1<<3) == 0 || CG.GetActiveClassEntry() != nil && (CG.GetActiveClassEntry().GetCeFlags()&1<<3) == 0 || ZendCheckProtected(g.CondF(fbc.GetPrototype() != nil, func() *ZendClassEntry { return fbc.GetPrototype().GetScope() }, func() *ZendClassEntry { return fbc.GetScope() }), CG.GetActiveClassEntry()) == 0) {

					/* incompatibe function */

					fbc = nil

					/* incompatibe function */

				}
			}
		}
	}
	ZendCompileCallCommon(result, args_ast, fbc)
}

/* }}} */

func ZendCompileNew(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var args_ast *ZendAst = ast.GetChild()[1]
	var class_node Znode
	var ctor_result Znode
	var opline *ZendOp
	if class_ast.GetKind() == ZEND_AST_CLASS {

		/* anon class declaration */

		opline = ZendCompileClassDecl(class_ast, 0)
		class_node.SetOpType(opline.GetResultType())
		class_node.GetOp().SetVar(opline.GetResult().GetVar())
	} else {
		ZendCompileClassRef(&class_node, class_ast, 0x200)
	}
	opline = ZendEmitOp(result, 68, nil, nil)
	if class_node.GetOpType() == 1<<0 {
		opline.SetOp1Type(1 << 0)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().GetValue().GetStr()))
		opline.GetOp2().SetNum(ZendAllocCacheSlot())
	} else {
		opline.SetOp1Type(&class_node.GetOpType())
		if &class_node.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&(&class_node).u.constant))
		} else {
			opline.SetOp1(&class_node.GetOp())
		}
	}
	ZendCompileCallCommon(&ctor_result, args_ast, nil)
	ZendDoFree(&ctor_result)
}

/* }}} */

func ZendCompileClone(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var obj_node Znode
	ZendCompileExpr(&obj_node, obj_ast)
	ZendEmitOpTmp(result, 110, &obj_node, nil)
}

/* }}} */

func ZendCompileGlobalVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var name_ast *ZendAst = var_ast.GetChild()[0]
	var name_node Znode
	var result Znode
	ZendCompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == 1<<0 {
		if &name_node.u.constant.u1.v.type_ != 6 {
			_convertToString(&name_node.u.constant)
		}
	}
	if IsThisFetch(var_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot use $this as global variable")
	} else if ZendTryCompileCv(&result, var_ast) == SUCCESS {
		var opline *ZendOp = ZendEmitOp(nil, 168, &result, &name_node)
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {

		/* name_ast should be evaluated only. FETCH_GLOBAL_LOCK instructs FETCH_W
		 * to not free the name_node operand, so it can be reused in the following
		 * ASSIGN_REF, which then frees it. */

		var opline *ZendOp = ZendEmitOp(&result, 83, &name_node, nil)
		opline.SetExtendedValue(1 << 3)
		if name_node.GetOpType() == 1<<0 {
			ZendStringAddref(name_node.GetConstant().GetValue().GetStr())
		}
		ZendEmitAssignRefZnode(ZendAstCreate1(ZEND_AST_VAR, ZendAstCreateZnode(&name_node)), &result)
	}
}

/* }}} */

func ZendCompileStaticVarCommon(var_name *ZendString, value *Zval, mode uint32) {
	var opline *ZendOp
	if CG.GetActiveOpArray().GetStaticVariables() == nil {
		if CG.GetActiveOpArray().GetScope() != nil {
			CG.GetActiveOpArray().GetScope().SetCeFlags(CG.GetActiveOpArray().GetScope().GetCeFlags() | 1<<16)
		}
		CG.GetActiveOpArray().SetStaticVariables(_zendNewArray(8))
	}
	value = ZendHashUpdate(CG.GetActiveOpArray().GetStaticVariables(), var_name, value)
	if var_name.GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.GetVal(), "this", g.SizeOf("\"this\"")-1)) {
		ZendErrorNoreturn(1<<6, "Cannot use $this as static variable")
	}
	opline = ZendEmitOp(nil, 183, nil, nil)
	opline.SetOp1Type(1 << 3)
	opline.GetOp1().SetVar(LookupCv(var_name))
	opline.SetExtendedValue(uint32((*byte)(value-(*byte)(CG.GetActiveOpArray().GetStaticVariables().GetArData()))) | mode)
}

/* }}} */

func ZendCompileStaticVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var value_ast *ZendAst = ast.GetChild()[1]
	var value_zv Zval
	if value_ast != nil {
		ZendConstExprToZval(&value_zv, value_ast)
	} else {
		&value_zv.SetTypeInfo(1)
	}
	ZendCompileStaticVarCommon(ZendAstGetStr(var_ast), &value_zv, 1)
}

/* }}} */

func ZendCompileUnset(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var var_node Znode
	var opline *ZendOp
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		if IsThisFetch(var_ast) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot unset $this")
		} else if ZendTryCompileCv(&var_node, var_ast) == SUCCESS {
			opline = ZendEmitOp(nil, 153, &var_node, nil)
		} else {
			opline = ZendCompileSimpleVarNoCv(nil, var_ast, 5, 0)
			opline.SetOpcode(74)
		}
		return
	case ZEND_AST_DIM:
		opline = ZendCompileDim(nil, var_ast, 5)
		opline.SetOpcode(75)
		return
	case ZEND_AST_PROP:
		opline = ZendCompileProp(nil, var_ast, 5, 0)
		opline.SetOpcode(76)
		return
	case ZEND_AST_STATIC_PROP:
		opline = ZendCompileStaticProp(nil, var_ast, 5, 0, 0)
		opline.SetOpcode(179)
		return
	default:
		break
	}
}

/* }}} */

func ZendHandleLoopsAndFinallyEx(depth ZendLong, return_value *Znode) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = ZendStackTop(&CG.loop_var_stack)
	if loop_var == nil {
		return 1
	}
	base = ZendStackBase(&CG.loop_var_stack)
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == 162 {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(162)
			opline.SetResultType(1 << 1)
			opline.GetResult().SetVar(loop_var.GetVarNum())
			if return_value != nil {
				opline.SetOp2Type(return_value.GetOpType())
				if return_value.GetOpType() == 1<<0 {
					opline.GetOp2().SetConstant(ZendAddLiteral(&return_value.u.constant))
				} else {
					opline.SetOp2(return_value.GetOp())
				}
			}
			opline.GetOp1().SetNum(loop_var.GetTryCatchOffset())
		} else if loop_var.GetOpcode() == 159 {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(159)
			opline.SetOp1Type(1 << 1)
			opline.GetOp1().SetVar(loop_var.GetVarNum())
		} else if loop_var.GetOpcode() == 62 {

			/* Stack separator */

			break

			/* Stack separator */

		} else if depth <= 1 {
			return 1
		} else if loop_var.GetOpcode() == 0 {

			/* Loop doesn't have freeable variable */

			depth--

			/* Loop doesn't have freeable variable */

		} else {
			var opline *ZendOp
			assert((loop_var.GetVarType() & (1<<2 | 1<<1)) != 0)
			opline = GetNextOp()
			opline.SetOpcode(loop_var.GetOpcode())
			opline.SetOp1Type(loop_var.GetVarType())
			opline.GetOp1().SetVar(loop_var.GetVarNum())
			opline.SetExtendedValue(1 << 0)
			depth--
		}
	}
	return depth == 0
}

/* }}} */

func ZendHandleLoopsAndFinally(return_value *Znode) int {
	return ZendHandleLoopsAndFinallyEx(ZendStackCount(&CG.loop_var_stack)+1, return_value)
}

/* }}} */

func ZendHasFinallyEx(depth ZendLong) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = ZendStackTop(&CG.loop_var_stack)
	if loop_var == nil {
		return 0
	}
	base = ZendStackBase(&CG.loop_var_stack)
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == 162 {
			return 1
		} else if loop_var.GetOpcode() == 159 {

		} else if loop_var.GetOpcode() == 62 {

			/* Stack separator */

			return 0

			/* Stack separator */

		} else if depth <= 1 {
			return 0
		} else {
			depth--
		}
	}
	return 0
}

/* }}} */

func ZendHasFinally() int {
	return ZendHasFinallyEx(ZendStackCount(&CG.loop_var_stack) + 1)
}

/* }}} */

func ZendCompileReturn(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var is_generator ZendBool = (CG.GetActiveOpArray().GetFnFlags() & 1 << 24) != 0
	var by_ref ZendBool = (CG.GetActiveOpArray().GetFnFlags() & 1 << 12) != 0
	var expr_node Znode
	var opline *ZendOp
	if is_generator != 0 {

		/* For generators the by-ref flag refers to yields, not returns */

		by_ref = 0

		/* For generators the by-ref flag refers to yields, not returns */

	}
	if expr_ast == nil {
		expr_node.SetOpType(1 << 0)
		&expr_node.u.constant.u1.type_info = 1
	} else if by_ref != 0 && ZendIsVariable(expr_ast) != 0 {
		ZendCompileVar(&expr_node, expr_ast, 1, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
	}
	if (CG.GetActiveOpArray().GetFnFlags()&1<<15) != 0 && (expr_node.GetOpType() == 1<<3 || by_ref != 0 && expr_node.GetOpType() == 1<<2) && ZendHasFinally() != 0 {

		/* Copy return value into temporary VAR to avoid modification in finally code */

		if by_ref != 0 {
			ZendEmitOp(&expr_node, 140, &expr_node, nil)
		} else {
			ZendEmitOpTmp(&expr_node, 31, &expr_node, nil)
		}

		/* Copy return value into temporary VAR to avoid modification in finally code */

	}

	/* Generator return types are handled separately */

	if is_generator == 0 && (CG.GetActiveOpArray().GetFnFlags()&1<<13) != 0 {
		ZendEmitReturnTypeCheck(g.Cond(expr_ast != nil, &expr_node, nil), CG.GetActiveOpArray().GetArgInfo()-1, 0)
	}
	ZendHandleLoopsAndFinally(g.Cond((expr_node.GetOpType()&(1<<1|1<<2)) != 0, &expr_node, nil))
	opline = ZendEmitOp(nil, g.Cond(by_ref != 0, 111, 62), &expr_node, nil)
	if by_ref != 0 && expr_ast != nil {
		if ZendIsCall(expr_ast) != 0 {
			opline.SetExtendedValue(1 << 0)
		} else if ZendIsVariable(expr_ast) == 0 {
			opline.SetExtendedValue(1 << 1)
		}
	}
}

/* }}} */

func ZendCompileEcho(ast *ZendAst) {
	var opline *ZendOp
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, 136, &expr_node, nil)
	opline.SetExtendedValue(0)
}

/* }}} */

func ZendCompileThrow(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	ZendEmitOp(nil, 108, &expr_node, nil)
}

/* }}} */

func ZendCompileBreakContinue(ast *ZendAst) {
	var depth_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	var depth ZendLong
	assert(ast.GetKind() == ZEND_AST_BREAK || ast.GetKind() == ZEND_AST_CONTINUE)
	if depth_ast != nil {
		var depth_zv *Zval
		if depth_ast.GetKind() != ZEND_AST_ZVAL {
			ZendErrorNoreturn(1<<6, "'%s' operator with non-integer operand "+"is no longer supported", g.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth_zv = ZendAstGetZval(depth_ast)
		if depth_zv.GetType() != 4 || depth_zv.GetValue().GetLval() < 1 {
			ZendErrorNoreturn(1<<6, "'%s' operator accepts only positive integers", g.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth = depth_zv.GetValue().GetLval()
	} else {
		depth = 1
	}
	if CG.GetContext().GetCurrentBrkCont() == -1 {
		ZendErrorNoreturn(1<<6, "'%s' not in the 'loop' or 'switch' context", g.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
	} else {
		if ZendHandleLoopsAndFinallyEx(depth, nil) == 0 {
			ZendErrorNoreturn(1<<6, "Cannot '%s' "+"%"+"lld"+" level%s", g.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"), depth, g.Cond(depth == 1, "", "s"))
		}
	}
	if ast.GetKind() == ZEND_AST_CONTINUE {
		var d int
		var cur int = CG.GetContext().GetCurrentBrkCont()
		for d = depth - 1; d > 0; d-- {
			cur = CG.GetContext().GetBrkContArray()[cur].GetParent()
			assert(cur != -1)
		}
		if CG.GetContext().GetBrkContArray()[cur].GetIsSwitch() != 0 {
			if depth == 1 {
				ZendError(1<<1, "\"continue\" targeting switch is equivalent to \"break\". "+"Did you mean to use \"continue "+"%"+"lld"+"\"?", depth+1)
			} else {
				ZendError(1<<1, "\"continue "+"%"+"lld"+"\" targeting switch is equivalent to \"break "+"%"+"lld"+"\". "+"Did you mean to use \"continue "+"%"+"lld"+"\"?", depth, depth, depth+1)
			}
		}
	}
	opline = ZendEmitOp(nil, g.Cond(ast.GetKind() == ZEND_AST_BREAK, 254, 255), nil, nil)
	opline.GetOp1().SetNum(CG.GetContext().GetCurrentBrkCont())
	opline.GetOp2().SetNum(depth)
}

/* }}} */

func ZendResolveGotoLabel(op_array *ZendOpArray, opline *ZendOp) {
	var dest *ZendLabel
	var current int
	var remove_oplines int = opline.GetOp1().GetNum()
	var label *Zval
	var opnum uint32 = opline - op_array.GetOpcodes()
	label = op_array.GetLiterals() + opline.GetOp2().GetConstant()
	if CG.GetContext().GetLabels() == nil || g.Assign(&dest, ZendHashFindPtr(CG.GetContext().GetLabels(), label.GetValue().GetStr())) == nil {
		CG.SetInCompilation(1)
		CG.SetActiveOpArray(op_array)
		CG.SetZendLineno(opline.GetLineno())
		ZendErrorNoreturn(1<<6, "'goto' to undefined label '%s'", label.GetValue().GetStr().GetVal())
	}
	ZvalPtrDtorStr(label)
	label.SetTypeInfo(1)
	current = opline.GetExtendedValue()
	for ; current != dest.GetBrkCont(); current = CG.GetContext().GetBrkContArray()[current].GetParent() {
		if current == -1 {
			CG.SetInCompilation(1)
			CG.SetActiveOpArray(op_array)
			CG.SetZendLineno(opline.GetLineno())
			ZendErrorNoreturn(1<<6, "'goto' into loop or switch statement is disallowed")
		}
		if CG.GetContext().GetBrkContArray()[current].GetStart() >= 0 {
			remove_oplines--
		}
	}
	for current = 0; current < op_array.GetLastTryCatch(); current++ {
		var elem *ZendTryCatchElement = &op_array.try_catch_array[current]
		if elem.GetTryOp() > opnum {
			break
		}
		if elem.GetFinallyOp() != 0 && opnum < elem.GetFinallyOp()-1 && (dest.GetOplineNum() > elem.GetFinallyEnd() || dest.GetOplineNum() < elem.GetTryOp()) {
			remove_oplines--
		}
	}
	opline.SetOpcode(42)
	opline.GetOp1().SetOplineNum(dest.GetOplineNum())
	opline.SetExtendedValue(0)
	opline.SetOp1Type(0)
	opline.SetOp2Type(0)
	opline.SetResultType(0)
	assert(remove_oplines >= 0)
	for g.PostDec(&remove_oplines) {
		opline--
		opline.GetOp1().SetNum(0)
		opline.GetOp2().SetNum(0)
		opline.GetResult().SetNum(0)
		opline.SetOpcode(0)
		opline.SetOp1Type(0)
		opline.SetOp2Type(0)
		opline.SetResultType(0)
		ZendVmSetOpcodeHandler(opline)
	}
}

/* }}} */

func ZendCompileGoto(ast *ZendAst) {
	var label_ast *ZendAst = ast.GetChild()[0]
	var label_node Znode
	var opline *ZendOp
	var opnum_start uint32 = GetNextOpNumber()
	ZendCompileExpr(&label_node, label_ast)

	/* Label resolution and unwinding adjustments happen in pass two. */

	ZendHandleLoopsAndFinally(nil)
	opline = ZendEmitOp(nil, 253, nil, &label_node)
	opline.GetOp1().SetNum(GetNextOpNumber() - opnum_start - 1)
	opline.SetExtendedValue(CG.GetContext().GetCurrentBrkCont())
}

/* }}} */

func ZendCompileLabel(ast *ZendAst) {
	var label *ZendString = ZendAstGetStr(ast.GetChild()[0])
	var dest ZendLabel
	if CG.GetContext().GetLabels() == nil {
		CG.GetContext().SetLabels((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
		_zendHashInit(CG.GetContext().GetLabels(), 8, LabelPtrDtor, 0)
	}
	dest.SetBrkCont(CG.GetContext().GetCurrentBrkCont())
	dest.SetOplineNum(GetNextOpNumber())
	if !(ZendHashAddMem(CG.GetContext().GetLabels(), label, &dest, g.SizeOf("zend_label"))) {
		ZendErrorNoreturn(1<<6, "Label '%s' already defined", label.GetVal())
	}
}

/* }}} */

func ZendCompileWhile(ast *ZendAst) {
	var cond_ast *ZendAst = ast.GetChild()[0]
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var cond_node Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_cond uint32
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(0, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendUpdateJumpTarget(opnum_jmp, opnum_cond)
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(44, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}

/* }}} */

func ZendCompileDoWhile(ast *ZendAst) {
	var stmt_ast *ZendAst = ast.GetChild()[0]
	var cond_ast *ZendAst = ast.GetChild()[1]
	var cond_node Znode
	var opnum_start uint32
	var opnum_cond uint32
	ZendBeginLoop(0, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(44, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}

/* }}} */

func ZendCompileExprList(result *Znode, ast *ZendAst) {
	var list *ZendAstList
	var i uint32
	result.SetOpType(1 << 0)
	&result.u.constant.u1.type_info = 3
	if ast == nil {
		return
	}
	list = ZendAstGetList(ast)
	for i = 0; i < list.GetChildren(); i++ {
		var expr_ast *ZendAst = list.GetChild()[i]
		ZendDoFree(result)
		ZendCompileExpr(result, expr_ast)
	}
}

/* }}} */

func ZendCompileFor(ast *ZendAst) {
	var init_ast *ZendAst = ast.GetChild()[0]
	var cond_ast *ZendAst = ast.GetChild()[1]
	var loop_ast *ZendAst = ast.GetChild()[2]
	var stmt_ast *ZendAst = ast.GetChild()[3]
	var result Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_loop uint32
	ZendCompileExprList(&result, init_ast)
	ZendDoFree(&result)
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(0, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_loop = GetNextOpNumber()
	ZendCompileExprList(&result, loop_ast)
	ZendDoFree(&result)
	ZendUpdateJumpTargetToNext(opnum_jmp)
	ZendCompileExprList(&result, cond_ast)
	ZendDoExtendedStmt()
	ZendEmitCondJump(44, &result, opnum_start)
	ZendEndLoop(opnum_loop, nil)
}

/* }}} */

func ZendCompileForeach(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var value_ast *ZendAst = ast.GetChild()[1]
	var key_ast *ZendAst = ast.GetChild()[2]
	var stmt_ast *ZendAst = ast.GetChild()[3]
	var by_ref ZendBool = value_ast.GetKind() == ZEND_AST_REF
	var is_variable ZendBool = ZendIsVariable(expr_ast) != 0 && ZendCanWriteToVariable(expr_ast) != 0
	var expr_node Znode
	var reset_node Znode
	var value_node Znode
	var key_node Znode
	var opline *ZendOp
	var opnum_reset uint32
	var opnum_fetch uint32
	if key_ast != nil {
		if key_ast.GetKind() == ZEND_AST_REF {
			ZendErrorNoreturn(1<<6, "Key element cannot be a reference")
		}
		if key_ast.GetKind() == ZEND_AST_ARRAY {
			ZendErrorNoreturn(1<<6, "Cannot use list as key element")
		}
	}
	if by_ref != 0 {
		value_ast = value_ast.GetChild()[0]
	}
	if value_ast.GetKind() == ZEND_AST_ARRAY && ZendPropagateListRefs(value_ast) != 0 {
		by_ref = 1
	}
	if by_ref != 0 && is_variable != 0 {
		ZendCompileVar(&expr_node, expr_ast, 1, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
	}
	if by_ref != 0 {
		ZendSeparateIfCallAndWrite(&expr_node, expr_ast, 1)
	}
	opnum_reset = GetNextOpNumber()
	opline = ZendEmitOp(&reset_node, g.Cond(by_ref != 0, 125, 77), &expr_node, nil)
	ZendBeginLoop(127, &reset_node, 0)
	opnum_fetch = GetNextOpNumber()
	opline = ZendEmitOp(nil, g.Cond(by_ref != 0, 126, 78), &reset_node, nil)
	if IsThisFetch(value_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot re-assign $this")
	} else if value_ast.GetKind() == ZEND_AST_VAR && ZendTryCompileCv(&value_node, value_ast) == SUCCESS {
		opline.SetOp2Type(&value_node.GetOpType())
		if &value_node.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&value_node).u.constant))
		} else {
			opline.SetOp2(&value_node.GetOp())
		}
	} else {
		opline.SetOp2Type(1 << 2)
		opline.GetOp2().SetVar(GetTemporaryVariable())
		&value_node.SetOpType(opline.GetOp2Type())
		if &value_node.GetOpType() == 1<<0 {
			var _z1 *Zval = &(&value_node).u.constant
			var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			&value_node.SetOp(opline.GetOp2())
		}
		if value_ast.GetKind() == ZEND_AST_ARRAY {
			ZendCompileListAssign(nil, value_ast, &value_node, value_ast.GetAttr())
		} else if by_ref != 0 {
			ZendEmitAssignRefZnode(value_ast, &value_node)
		} else {
			ZendEmitAssignZnode(value_ast, &value_node)
		}
	}
	if key_ast != nil {
		opline = &CG.active_op_array.GetOpcodes()[opnum_fetch]
		ZendMakeTmpResult(&key_node, opline)
		ZendEmitAssignZnode(key_ast, &key_node)
	}
	ZendCompileStmt(stmt_ast)

	/* Place JMP and FE_FREE on the line where foreach starts. It would be
	 * better to use the end line, but this information is not available
	 * currently. */

	CG.SetZendLineno(ast.GetLineno())
	ZendEmitJump(opnum_fetch)
	opline = &CG.active_op_array.GetOpcodes()[opnum_reset]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
	opline = &CG.active_op_array.GetOpcodes()[opnum_fetch]
	opline.SetExtendedValue(GetNextOpNumber())
	ZendEndLoop(opnum_fetch, &reset_node)
	opline = ZendEmitOp(nil, 127, &reset_node, nil)
}

/* }}} */

func ZendCompileIf(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var jmp_opnums *uint32 = nil
	if list.GetChildren() > 1 {
		jmp_opnums = _safeEmalloc(g.SizeOf("uint32_t"), list.GetChildren()-1, 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var cond_ast *ZendAst = elem_ast.GetChild()[0]
		var stmt_ast *ZendAst = elem_ast.GetChild()[1]
		if cond_ast != nil {
			var cond_node Znode
			var opnum_jmpz uint32
			ZendCompileExpr(&cond_node, cond_ast)
			opnum_jmpz = ZendEmitCondJump(43, &cond_node, 0)
			ZendCompileStmt(stmt_ast)
			if i != list.GetChildren()-1 {
				jmp_opnums[i] = ZendEmitJump(0)
			}
			ZendUpdateJumpTargetToNext(opnum_jmpz)
		} else {

			/* "else" can only occur as last element. */

			assert(i == list.GetChildren()-1)
			ZendCompileStmt(stmt_ast)
		}
	}
	if list.GetChildren() > 1 {
		for i = 0; i < list.GetChildren()-1; i++ {
			ZendUpdateJumpTargetToNext(jmp_opnums[i])
		}
		_efree(jmp_opnums)
	}
}

/* }}} */

func DetermineSwitchJumptableType(cases *ZendAstList) ZendUchar {
	var i uint32
	var common_type ZendUchar = 0
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast **ZendAst = &case_ast.child[0]
		var cond_zv *Zval
		if case_ast.GetChild()[0] == nil {

			/* Skip default clause */

			continue

			/* Skip default clause */

		}
		ZendEvalConstExpr(cond_ast)
		if (*cond_ast).GetKind() != ZEND_AST_ZVAL {

			/* Non-constant case */

			return 0

			/* Non-constant case */

		}
		cond_zv = ZendAstGetZval(case_ast.GetChild()[0])
		if cond_zv.GetType() != 4 && cond_zv.GetType() != 6 {

			/* We only optimize switched on integers and strings */

			return 0

			/* We only optimize switched on integers and strings */

		}
		if common_type == 0 {
			common_type = cond_zv.GetType()
		} else if common_type != cond_zv.GetType() {

			/* Non-uniform case types */

			return 0

			/* Non-uniform case types */

		}
		if cond_zv.GetType() == 6 && IsNumericString(cond_zv.GetValue().GetStr().GetVal(), cond_zv.GetValue().GetStr().GetLen(), nil, nil, 0) != 0 {

			/* Numeric strings cannot be compared with a simple hash lookup */

			return 0

			/* Numeric strings cannot be compared with a simple hash lookup */

		}
	}
	return common_type
}
func ShouldUseJumptable(cases *ZendAstList, jumptable_type ZendUchar) ZendBool {
	if (CG.GetCompilerOptions() & 1 << 16) != 0 {
		return 0
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */

	if jumptable_type == 4 {
		return cases.GetChildren() >= 5
	} else {
		assert(jumptable_type == 6)
		return cases.GetChildren() >= 2
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */
}
func ZendCompileSwitch(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var cases *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	var i uint32
	var has_default_case ZendBool = 0
	var expr_node Znode
	var case_node Znode
	var opline *ZendOp
	var jmpnz_opnums *uint32
	var opnum_default_jmp uint32
	var opnum_switch uint32 = uint32 - 1
	var jumptable_type ZendUchar
	var jumptable *HashTable = nil
	ZendCompileExpr(&expr_node, expr_ast)
	ZendBeginLoop(70, &expr_node, 1)
	case_node.SetOpType(1 << 1)
	case_node.GetOp().SetVar(GetTemporaryVariable())
	jumptable_type = DetermineSwitchJumptableType(cases)
	if jumptable_type != 0 && ShouldUseJumptable(cases, jumptable_type) != 0 {
		var jumptable_op Znode
		jumptable = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
		_zendHashInit(jumptable, cases.GetChildren(), nil, 0)
		jumptable_op.SetOpType(1 << 0)
		var __arr *ZendArray = jumptable
		var __z *Zval = &jumptable_op.u.constant
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		opline = ZendEmitOp(nil, g.Cond(jumptable_type == 4, 187, 188), &expr_node, &jumptable_op)
		if opline.GetOp1Type() == 1<<0 {
			if (CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant()).u1.v.type_flags != 0 {
				ZvalAddrefP(CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant())
			}
		}
		opnum_switch = opline - CG.GetActiveOpArray().GetOpcodes()
	}
	jmpnz_opnums = _safeEmalloc(g.SizeOf("uint32_t"), cases.GetChildren(), 0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast *ZendAst = case_ast.GetChild()[0]
		var cond_node Znode
		if cond_ast == nil {
			if has_default_case != 0 {
				CG.SetZendLineno(case_ast.GetLineno())
				ZendErrorNoreturn(1<<6, "Switch statements may only contain one default clause")
			}
			has_default_case = 1
			continue
		}
		ZendCompileExpr(&cond_node, cond_ast)
		if expr_node.GetOpType() == 1<<0 && expr_node.GetConstant().GetType() == 2 {
			jmpnz_opnums[i] = ZendEmitCondJump(43, &cond_node, 0)
		} else if expr_node.GetOpType() == 1<<0 && expr_node.GetConstant().GetType() == 3 {
			jmpnz_opnums[i] = ZendEmitCondJump(44, &cond_node, 0)
		} else {
			opline = ZendEmitOp(nil, g.Cond((expr_node.GetOpType()&(1<<2|1<<1)) != 0, 48, 18), &expr_node, &cond_node)
			opline.SetResultType(&case_node.GetOpType())
			if &case_node.GetOpType() == 1<<0 {
				opline.GetResult().SetConstant(ZendAddLiteral(&(&case_node).u.constant))
			} else {
				opline.SetResult(&case_node.GetOp())
			}
			if opline.GetOp1Type() == 1<<0 {
				if (CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant()).u1.v.type_flags != 0 {
					ZvalAddrefP(CG.GetActiveOpArray().GetLiterals() + opline.GetOp1().GetConstant())
				}
			}
			jmpnz_opnums[i] = ZendEmitCondJump(44, &case_node, 0)
		}
	}
	opnum_default_jmp = ZendEmitJump(0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast *ZendAst = case_ast.GetChild()[0]
		var stmt_ast *ZendAst = case_ast.GetChild()[1]
		if cond_ast != nil {
			ZendUpdateJumpTargetToNext(jmpnz_opnums[i])
			if jumptable != nil {
				var cond_zv *Zval = ZendAstGetZval(cond_ast)
				var jmp_target Zval
				var __z *Zval = &jmp_target
				__z.GetValue().SetLval(GetNextOpNumber())
				__z.SetTypeInfo(4)
				assert(cond_zv.GetType() == jumptable_type)
				if cond_zv.GetType() == 4 {
					ZendHashIndexAdd(jumptable, cond_zv.GetValue().GetLval(), &jmp_target)
				} else {
					assert(cond_zv.GetType() == 6)
					ZendHashAdd(jumptable, cond_zv.GetValue().GetStr(), &jmp_target)
				}
			}
		} else {
			ZendUpdateJumpTargetToNext(opnum_default_jmp)
			if jumptable != nil {
				assert(opnum_switch != uint32-1)
				opline = &CG.active_op_array.GetOpcodes()[opnum_switch]
				opline.SetExtendedValue(GetNextOpNumber())
			}
		}
		ZendCompileStmt(stmt_ast)
	}
	if has_default_case == 0 {
		ZendUpdateJumpTargetToNext(opnum_default_jmp)
		if jumptable != nil {
			opline = &CG.active_op_array.GetOpcodes()[opnum_switch]
			opline.SetExtendedValue(GetNextOpNumber())
		}
	}
	ZendEndLoop(GetNextOpNumber(), &expr_node)
	if (expr_node.GetOpType() & (1<<2 | 1<<1)) != 0 {
		opline = ZendEmitOp(nil, 70, &expr_node, nil)
		opline.SetExtendedValue(1 << 1)
	} else if expr_node.GetOpType() == 1<<0 {
		ZvalPtrDtorNogc(&expr_node.u.constant)
	}
	_efree(jmpnz_opnums)
}

/* }}} */

func ZendCompileTry(ast *ZendAst) {
	var try_ast *ZendAst = ast.GetChild()[0]
	var catches *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	var finally_ast *ZendAst = ast.GetChild()[2]
	var i uint32
	var j uint32
	var opline *ZendOp
	var try_catch_offset uint32
	var jmp_opnums *uint32 = _safeEmalloc(g.SizeOf("uint32_t"), catches.GetChildren(), 0)
	var orig_fast_call_var uint32 = CG.GetContext().GetFastCallVar()
	var orig_try_catch_offset uint32 = CG.GetContext().GetTryCatchOffset()
	if catches.GetChildren() == 0 && finally_ast == nil {
		ZendErrorNoreturn(1<<6, "Cannot use try without catch or finally")
	}

	/* label: try { } must not be equal to try { label: } */

	if CG.GetContext().GetLabels() != nil {
		var label *ZendLabel
		for {
			var __ht *HashTable = CG.GetContext().GetLabels()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				label = _z.GetValue().GetPtr()
				if label.GetOplineNum() == GetNextOpNumber() {
					ZendEmitOp(nil, 0, nil, nil)
				}
				break
			}
			break
		}
	}
	try_catch_offset = ZendAddTryElement(GetNextOpNumber())
	if finally_ast != nil {
		var fast_call ZendLoopVar
		if (CG.GetActiveOpArray().GetFnFlags() & 1 << 15) == 0 {
			CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<15)
		}
		CG.GetContext().SetFastCallVar(GetTemporaryVariable())

		/* Push FAST_CALL on unwind stack */

		fast_call.SetOpcode(162)
		fast_call.SetVarType(1 << 1)
		fast_call.SetVarNum(CG.GetContext().GetFastCallVar())
		fast_call.SetTryCatchOffset(try_catch_offset)
		ZendStackPush(&CG.loop_var_stack, &fast_call)
	}
	CG.GetContext().SetTryCatchOffset(try_catch_offset)
	ZendCompileStmt(try_ast)
	if catches.GetChildren() != 0 {
		jmp_opnums[0] = ZendEmitJump(0)
	}
	for i = 0; i < catches.GetChildren(); i++ {
		var catch_ast *ZendAst = catches.GetChild()[i]
		var classes *ZendAstList = ZendAstGetList(catch_ast.GetChild()[0])
		var var_ast *ZendAst = catch_ast.GetChild()[1]
		var stmt_ast *ZendAst = catch_ast.GetChild()[2]
		var var_name *ZendString = ZvalMakeInternedString(ZendAstGetZval(var_ast))
		var is_last_catch ZendBool = i+1 == catches.GetChildren()
		var jmp_multicatch *uint32 = _safeEmalloc(g.SizeOf("uint32_t"), classes.GetChildren()-1, 0)
		var opnum_catch uint32 = uint32 - 1
		CG.SetZendLineno(catch_ast.GetLineno())
		for j = 0; j < classes.GetChildren(); j++ {
			var class_ast *ZendAst = classes.GetChild()[j]
			var is_last_class ZendBool = j+1 == classes.GetChildren()
			if ZendIsConstDefaultClassRef(class_ast) == 0 {
				ZendErrorNoreturn(1<<6, "Bad class name in the catch statement")
			}
			opnum_catch = GetNextOpNumber()
			if i == 0 && j == 0 {
				CG.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetCatchOp(opnum_catch)
			}
			opline = GetNextOp()
			opline.SetOpcode(107)
			opline.SetOp1Type(1 << 0)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(ZendResolveClassNameAst(class_ast)))
			opline.SetExtendedValue(ZendAllocCacheSlot())
			if var_name.GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.GetVal(), "this", g.SizeOf("\"this\"")-1)) {
				ZendErrorNoreturn(1<<6, "Cannot re-assign $this")
			}
			opline.SetResultType(1 << 3)
			opline.GetResult().SetVar(LookupCv(var_name))
			if is_last_catch != 0 && is_last_class != 0 {
				opline.SetExtendedValue(opline.GetExtendedValue() | 1<<0)
			}
			if is_last_class == 0 {
				jmp_multicatch[j] = ZendEmitJump(0)
				opline = &CG.active_op_array.GetOpcodes()[opnum_catch]
				opline.GetOp2().SetOplineNum(GetNextOpNumber())
			}
		}
		for j = 0; j < classes.GetChildren()-1; j++ {
			ZendUpdateJumpTargetToNext(jmp_multicatch[j])
		}
		_efree(jmp_multicatch)
		ZendCompileStmt(stmt_ast)
		if is_last_catch == 0 {
			jmp_opnums[i+1] = ZendEmitJump(0)
		}
		assert(opnum_catch != uint32-1 && "Should have at least one class")
		opline = &CG.active_op_array.GetOpcodes()[opnum_catch]
		if is_last_catch == 0 {
			opline.GetOp2().SetOplineNum(GetNextOpNumber())
		}
	}
	for i = 0; i < catches.GetChildren(); i++ {
		ZendUpdateJumpTargetToNext(jmp_opnums[i])
	}
	if finally_ast != nil {
		var discard_exception ZendLoopVar
		var opnum_jmp uint32 = GetNextOpNumber() + 1

		/* Pop FAST_CALL from unwind stack */

		ZendStackDelTop(&CG.loop_var_stack)

		/* Push DISCARD_EXCEPTION on unwind stack */

		discard_exception.SetOpcode(159)
		discard_exception.SetVarType(1 << 1)
		discard_exception.SetVarNum(CG.GetContext().GetFastCallVar())
		ZendStackPush(&CG.loop_var_stack, &discard_exception)
		CG.SetZendLineno(finally_ast.GetLineno())
		opline = ZendEmitOp(nil, 162, nil, nil)
		opline.GetOp1().SetNum(try_catch_offset)
		opline.SetResultType(1 << 1)
		opline.GetResult().SetVar(CG.GetContext().GetFastCallVar())
		ZendEmitOp(nil, 42, nil, nil)
		ZendCompileStmt(finally_ast)
		CG.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyOp(opnum_jmp + 1)
		CG.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyEnd(GetNextOpNumber())
		opline = ZendEmitOp(nil, 163, nil, nil)
		opline.SetOp1Type(1 << 1)
		opline.GetOp1().SetVar(CG.GetContext().GetFastCallVar())
		opline.GetOp2().SetNum(orig_try_catch_offset)
		ZendUpdateJumpTargetToNext(opnum_jmp)
		CG.GetContext().SetFastCallVar(orig_fast_call_var)

		/* Pop DISCARD_EXCEPTION from unwind stack */

		ZendStackDelTop(&CG.loop_var_stack)

		/* Pop DISCARD_EXCEPTION from unwind stack */

	}
	CG.GetContext().SetTryCatchOffset(orig_try_catch_offset)
	_efree(jmp_opnums)
}

/* }}} */

func ZendHandleEncodingDeclaration(ast *ZendAst) ZendBool {
	var declares *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < declares.GetChildren(); i++ {
		var declare_ast *ZendAst = declares.GetChild()[i]
		var name_ast *ZendAst = declare_ast.GetChild()[0]
		var value_ast *ZendAst = declare_ast.GetChild()[1]
		var name *ZendString = ZendAstGetStr(name_ast)
		if name.GetLen() == g.SizeOf("\"encoding\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "encoding", g.SizeOf("\"encoding\"")-1) == 0 {
			if value_ast.GetKind() != ZEND_AST_ZVAL {
				ZendThrowException(ZendCeCompileError, "Encoding must be a literal", 0)
				return 0
			}
			if CG.GetMultibyte() != 0 {
				var encoding_name *ZendString = ZvalGetString(ZendAstGetZval(value_ast))
				var new_encoding *ZendEncoding
				var old_encoding *ZendEncoding
				var old_input_filter ZendEncodingFilter
				CG.SetEncodingDeclared(1)
				new_encoding = ZendMultibyteFetchEncoding(encoding_name.GetVal())
				if new_encoding == nil {
					ZendError(1<<7, "Unsupported encoding [%s]", encoding_name.GetVal())
				} else {
					old_input_filter = LANG_SCNG.GetInputFilter()
					old_encoding = LANG_SCNG.GetScriptEncoding()
					ZendMultibyteSetFilter(new_encoding)

					/* need to re-scan if input filter changed */

					if old_input_filter != LANG_SCNG.GetInputFilter() || old_input_filter != nil && new_encoding != old_encoding {
						ZendMultibyteYyinputAgain(old_input_filter, old_encoding)
					}

					/* need to re-scan if input filter changed */

				}
				ZendStringReleaseEx(encoding_name, 0)
			} else {
				ZendError(1<<7, "declare(encoding=...) ignored because "+"Zend multibyte feature is turned off by settings")
			}
		}
	}
	return 1
}

/* }}} */

func ZendDeclareIsFirstStatement(ast *ZendAst) int {
	var i uint32 = 0
	var file_ast *ZendAstList = ZendAstGetList(CG.GetAst())

	/* Check to see if this declare is preceded only by declare statements */

	for i < file_ast.GetChildren() {
		if file_ast.GetChild()[i] == ast {
			return SUCCESS
		} else if file_ast.GetChild()[i] == nil {

			/* Empty statements are not allowed prior to a declare */

			return FAILURE

			/* Empty statements are not allowed prior to a declare */

		} else if file_ast.GetChild()[i].GetKind() != ZEND_AST_DECLARE {

			/* declares can only be preceded by other declares */

			return FAILURE

			/* declares can only be preceded by other declares */

		}
		i++
	}
	return FAILURE
}

/* }}} */

func ZendCompileDeclare(ast *ZendAst) {
	var declares *ZendAstList = ZendAstGetList(ast.GetChild()[0])
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var orig_declarables ZendDeclarables = CG.GetFileContext().GetDeclarables()
	var i uint32
	for i = 0; i < declares.GetChildren(); i++ {
		var declare_ast *ZendAst = declares.GetChild()[i]
		var name_ast *ZendAst = declare_ast.GetChild()[0]
		var value_ast *ZendAst = declare_ast.GetChild()[1]
		var name *ZendString = ZendAstGetStr(name_ast)
		if value_ast.GetKind() != ZEND_AST_ZVAL {
			ZendErrorNoreturn(1<<6, "declare(%s) value must be a literal", name.GetVal())
		}
		if name.GetLen() == g.SizeOf("\"ticks\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "ticks", g.SizeOf("\"ticks\"")-1) == 0 {
			var value_zv Zval
			ZendConstExprToZval(&value_zv, value_ast)
			CG.GetFileContext().GetDeclarables().SetTicks(ZvalGetLong(&value_zv))
			ZvalPtrDtorNogc(&value_zv)
		} else if name.GetLen() == g.SizeOf("\"encoding\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "encoding", g.SizeOf("\"encoding\"")-1) == 0 {
			if FAILURE == ZendDeclareIsFirstStatement(ast) {
				ZendErrorNoreturn(1<<6, "Encoding declaration pragma must be "+"the very first statement in the script")
			}
		} else if name.GetLen() == g.SizeOf("\"strict_types\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "strict_types", g.SizeOf("\"strict_types\"")-1) == 0 {
			var value_zv Zval
			if FAILURE == ZendDeclareIsFirstStatement(ast) {
				ZendErrorNoreturn(1<<6, "strict_types declaration must be "+"the very first statement in the script")
			}
			if ast.GetChild()[1] != nil {
				ZendErrorNoreturn(1<<6, "strict_types declaration must not "+"use block mode")
			}
			ZendConstExprToZval(&value_zv, value_ast)
			if value_zv.GetType() != 4 || value_zv.GetValue().GetLval() != 0 && value_zv.GetValue().GetLval() != 1 {
				ZendErrorNoreturn(1<<6, "strict_types declaration must have 0 or 1 as its value")
			}
			if value_zv.GetValue().GetLval() == 1 {
				CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<31)
			}
		} else {
			ZendError(1<<7, "Unsupported declare '%s'", name.GetVal())
		}
	}
	if stmt_ast != nil {
		ZendCompileStmt(stmt_ast)
		CG.GetFileContext().SetDeclarables(orig_declarables)
	}
}

/* }}} */

func ZendCompileStmtList(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		ZendCompileStmt(list.GetChild()[i])
	}
}

/* }}} */

func ZendSetFunctionArgFlags(func_ *ZendFunction) {
	var i uint32
	var n uint32
	func_.GetArgFlags()[0] = 0
	func_.GetArgFlags()[1] = 0
	func_.GetArgFlags()[2] = 0
	if func_.GetArgInfo() != nil {
		if func_.GetNumArgs() < 12 {
			n = func_.GetNumArgs()
		} else {
			n = 12
		}
		i = 0
		for i < n {
			func_.SetQuickArgFlags(func_.GetQuickArgFlags() | func_.GetArgInfo()[i].GetPassByReference()<<6<<(i+1)*2)
			i++
		}
		if (func_.GetFnFlags()&1<<14) != 0 && func_.GetArgInfo()[i].GetPassByReference() != 0 {
			var pass_by_reference uint32 = func_.GetArgInfo()[i].GetPassByReference()
			for i < 12 {
				func_.SetQuickArgFlags(func_.GetQuickArgFlags() | pass_by_reference<<6<<(i+1)*2)
				i++
			}
		}
	}
}

/* }}} */

func ZendCompileTypename(ast *ZendAst, force_allow_null ZendBool) ZendType {
	var allow_null ZendBool = force_allow_null
	var orig_ast_attr ZendAstAttr = ast.GetAttr()
	var type_ ZendType
	if (ast.GetAttr() & 1 << 8) != 0 {
		allow_null = 1
		ast.SetAttr(ast.GetAttr() &^ (1 << 8))
	}
	if ast.GetKind() == ZEND_AST_TYPE {
		return ast.GetAttr()<<2 | g.Cond(allow_null != 0, 0x1, 0x0)
	} else {
		var class_name *ZendString = ZendAstGetStr(ast)
		var type_code ZendUchar = ZendLookupBuiltinTypeByName(class_name)
		if type_code != 0 {
			if (ast.GetAttr() & 1) != 1 {
				ZendErrorNoreturn(1<<6, "Type declaration '%s' must be unqualified", ZendStringTolowerEx(class_name, 0).GetVal())
			}
			type_ = type_code<<2 | g.Cond(allow_null != 0, 0x1, 0x0)
		} else {
			var fetch_type uint32 = ZendGetClassFetchTypeAst(ast)
			if fetch_type == 0 {
				class_name = ZendResolveClassNameAst(ast)
				ZendAssertValidClassName(class_name)
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				ZendStringAddref(class_name)
			}
			type_ = uintptr_t(class_name) | g.Cond(allow_null != 0, 0x1, 0x0)
		}
	}
	ast.SetAttr(orig_ast_attr)
	return type_
}

/* }}} */

func ZendCompileParams(ast *ZendAst, return_type_ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var arg_infos *ZendArgInfo
	if return_type_ast != nil {

		/* Use op_array->arg_info[-1] for return type */

		arg_infos = _safeEmalloc(g.SizeOf("zend_arg_info"), list.GetChildren()+1, 0)
		arg_infos.SetName(nil)
		arg_infos.SetPassByReference((op_array.GetFnFlags() & 1 << 12) != 0)
		arg_infos.SetIsVariadic(0)
		arg_infos.SetType(ZendCompileTypename(return_type_ast, 0))
		if arg_infos.GetType()>>2 == 19 && (arg_infos.GetType()&0x1) != 0 {
			ZendErrorNoreturn(1<<6, "Void type cannot be nullable")
		}
		arg_infos++
		op_array.SetFnFlags(op_array.GetFnFlags() | 1<<13)
	} else {
		if list.GetChildren() == 0 {
			return
		}
		arg_infos = _safeEmalloc(g.SizeOf("zend_arg_info"), list.GetChildren(), 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var param_ast *ZendAst = list.GetChild()[i]
		var type_ast *ZendAst = param_ast.GetChild()[0]
		var var_ast *ZendAst = param_ast.GetChild()[1]
		var default_ast *ZendAst = param_ast.GetChild()[2]
		var name *ZendString = ZvalMakeInternedString(ZendAstGetZval(var_ast))
		var is_ref ZendBool = (param_ast.GetAttr() & 1 << 0) != 0
		var is_variadic ZendBool = (param_ast.GetAttr() & 1 << 1) != 0
		var var_node Znode
		var default_node Znode
		var opcode ZendUchar
		var opline *ZendOp
		var arg_info *ZendArgInfo
		if ZendIsAutoGlobal(name) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot re-assign auto-global variable %s", name.GetVal())
		}
		var_node.SetOpType(1 << 3)
		var_node.GetOp().SetVar(LookupCv(name))
		if uint32((*Zval)((*byte)(nil)+int(var_node.GetOp().GetVar()))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0)))) != i {
			ZendErrorNoreturn(1<<6, "Redefinition of parameter $%s", name.GetVal())
		} else if name.GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(name.GetVal(), "this", g.SizeOf("\"this\"")-1)) {
			ZendErrorNoreturn(1<<6, "Cannot use $this as parameter")
		}
		if (op_array.GetFnFlags() & 1 << 14) != 0 {
			ZendErrorNoreturn(1<<6, "Only the last parameter can be variadic")
		}
		if is_variadic != 0 {
			opcode = 164
			default_node.SetOpType(0)
			op_array.SetFnFlags(op_array.GetFnFlags() | 1<<14)
			if default_ast != nil {
				ZendErrorNoreturn(1<<6, "Variadic parameter cannot have a default value")
			}
		} else if default_ast != nil {

			/* we cannot substitute constants here or it will break ReflectionParameter::getDefaultValueConstantName() and ReflectionParameter::isDefaultValueConstant() */

			var cops uint32 = CG.GetCompilerOptions()
			CG.SetCompilerOptions(CG.GetCompilerOptions() | 1<<6 | 1<<8)
			opcode = 64
			default_node.SetOpType(1 << 0)
			ZendConstExprToZval(&default_node.u.constant, default_ast)
			CG.SetCompilerOptions(cops)
		} else {
			opcode = 63
			default_node.SetOpType(0)
			op_array.SetRequiredNumArgs(i + 1)
		}
		arg_info = &arg_infos[i]
		arg_info.SetName(ZendStringCopy(name))
		arg_info.SetPassByReference(is_ref)
		arg_info.SetIsVariadic(is_variadic)

		/* TODO: Keep compatibility, but may be better reset "allow_null" ??? */

		arg_info.SetType(0<<2 | g.Cond(true, 0x1, 0x0))
		if type_ast != nil {
			var has_null_default ZendBool = default_ast != nil && (default_node.GetConstant().GetType() == 1 || default_node.GetConstant().GetType() == 11 && (*ZendAst)((*byte)(default_node.GetConstant().GetValue().GetAst())+g.SizeOf("zend_ast_ref")).GetKind() == ZEND_AST_CONSTANT && strcasecmp(ZendAstGetConstantName((*ZendAst)((*byte)(default_node.GetConstant().GetValue().GetAst())+g.SizeOf("zend_ast_ref"))).GetVal(), "NULL") == 0)
			op_array.SetFnFlags(op_array.GetFnFlags() | 1<<8)
			arg_info.SetType(ZendCompileTypename(type_ast, has_null_default))
			if arg_info.GetType()>>2 == 19 {
				ZendErrorNoreturn(1<<6, "void cannot be used as a parameter type")
			}
			if type_ast.GetKind() == ZEND_AST_TYPE {
				if arg_info.GetType()>>2 == 7 {
					if default_ast != nil && has_null_default == 0 && default_node.GetConstant().GetType() != 7 && default_node.GetConstant().GetType() != 11 {
						ZendErrorNoreturn(1<<6, "Default value for parameters "+"with array type can only be an array or NULL")
					}
				} else if arg_info.GetType()>>2 == 17 && default_ast != nil {
					if has_null_default == 0 && default_node.GetConstant().GetType() != 11 {
						ZendErrorNoreturn(1<<6, "Default value for parameters "+"with callable type can only be NULL")
					}
				}
			} else {
				if default_ast != nil && has_null_default == 0 && default_node.GetConstant().GetType() != 11 {
					if arg_info.GetType() > 0x3ff {
						ZendErrorNoreturn(1<<6, "Default value for parameters "+"with a class type can only be NULL")
					} else {
						switch arg_info.GetType() >> 2 {
						case 5:
							if default_node.GetConstant().GetType() != 5 && default_node.GetConstant().GetType() != 4 {
								ZendErrorNoreturn(1<<6, "Default value for parameters "+"with a float type can only be float, integer, or NULL")
							}
							ConvertToDouble(&default_node.u.constant)
							break
						case 18:
							if default_node.GetConstant().GetType() != 7 {
								ZendErrorNoreturn(1<<6, "Default value for parameters "+"with iterable type can only be an array or NULL")
							}
							break
						case 8:
							ZendErrorNoreturn(1<<6, "Default value for parameters "+"with an object type can only be NULL")
							break
						default:
							if !(arg_info.GetType()>>2 == default_node.GetConstant().GetType() || arg_info.GetType()>>2 == 16 && (default_node.GetConstant().GetType() == 3 || default_node.GetConstant().GetType() == 2)) {
								ZendErrorNoreturn(1<<6, "Default value for parameters "+"with a %s type can only be %s or NULL", ZendGetTypeByConst(arg_info.GetType()>>2), ZendGetTypeByConst(arg_info.GetType()>>2))
							}
							break
						}
					}
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, nil, &default_node)
		opline.SetResultType(&var_node.GetOpType())
		if &var_node.GetOpType() == 1<<0 {
			opline.GetResult().SetConstant(ZendAddLiteral(&(&var_node).u.constant))
		} else {
			opline.SetResult(&var_node.GetOp())
		}
		opline.GetOp1().SetNum(i + 1)
		if type_ast != nil {

			/* Allocate cache slot to speed-up run-time class resolution */

			if opline.GetOpcode() == 64 {
				if arg_info.GetType() > 0x3ff {
					opline.SetExtendedValue(ZendAllocCacheSlot())
				}
			} else {
				if arg_info.GetType() > 0x3ff {
					opline.GetOp2().SetNum(op_array.GetCacheSize())
					op_array.SetCacheSize(op_array.GetCacheSize() + g.SizeOf("void *"))
				} else {
					opline.GetOp2().SetNum(-1)
				}
			}

			/* Allocate cache slot to speed-up run-time class resolution */

		} else {
			if opline.GetOpcode() != 64 {
				opline.GetOp2().SetNum(-1)
			}
		}
	}

	/* These are assigned at the end to avoid uninitialized memory in case of an error */

	op_array.SetNumArgs(list.GetChildren())
	op_array.SetArgInfo(arg_infos)

	/* Don't count the variadic argument */

	if (op_array.GetFnFlags() & 1 << 14) != 0 {
		op_array.GetNumArgs()--
	}
	ZendSetFunctionArgFlags((*ZendFunction)(op_array))
}

/* }}} */

func ZendCompileClosureBinding(closure *Znode, op_array *ZendOpArray, uses_ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(uses_ast)
	var i uint32
	if list.GetChildren() == 0 {
		return
	}
	if op_array.GetStaticVariables() == nil {
		op_array.SetStaticVariables(_zendNewArray(8))
	}
	for i = 0; i < list.GetChildren(); i++ {
		var var_name_ast *ZendAst = list.GetChild()[i]
		var var_name *ZendString = ZvalMakeInternedString(ZendAstGetZval(var_name_ast))
		var mode uint32 = var_name_ast.GetAttr()
		var opline *ZendOp
		var value *Zval
		if var_name.GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.GetVal(), "this", g.SizeOf("\"this\"")-1)) {
			ZendErrorNoreturn(1<<6, "Cannot use $this as lexical variable")
		}
		if ZendIsAutoGlobal(var_name) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot use auto-global as lexical variable")
		}
		value = ZendHashAdd(op_array.GetStaticVariables(), var_name, &EG.uninitialized_zval)
		if value == nil {
			ZendErrorNoreturn(1<<6, "Cannot use variable $%s twice", var_name.GetVal())
		}
		CG.SetZendLineno(ZendAstGetLineno(var_name_ast))
		opline = ZendEmitOp(nil, 182, closure, nil)
		opline.SetOp2Type(1 << 3)
		opline.GetOp2().SetVar(LookupCv(var_name))
		opline.SetExtendedValue(uint32((*byte)(value-(*byte)(op_array.GetStaticVariables().GetArData()))) | mode)
	}
}

/* }}} */

// @type ClosureInfo struct
func FindImplicitBindsRecursively(info *ClosureInfo, ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_VAR {
		var name_ast *ZendAst = ast.GetChild()[0]
		if name_ast.GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(name_ast).GetType() == 6 {
			var name *ZendString = ZendAstGetStr(name_ast)
			if ZendIsAutoGlobal(name) != 0 {

				/* These is no need to explicitly import auto-globals. */

				return

				/* These is no need to explicitly import auto-globals. */

			}
			if name.GetLen() == g.SizeOf("\"this\"")-1 && !(memcmp(name.GetVal(), "this", g.SizeOf("\"this\"")-1)) {

				/* $this does not need to be explicitly imported. */

				return

				/* $this does not need to be explicitly imported. */

			}
			ZendHashAddEmptyElement(&info.uses, name)
		} else {
			info.SetVarvarsUsed(1)
			FindImplicitBindsRecursively(info, name_ast)
		}
	} else if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			FindImplicitBindsRecursively(info, list.GetChild()[i])
		}
	} else if ast.GetKind() == ZEND_AST_CLOSURE {

		/* For normal closures add the use() list. */

		var closure_ast *ZendAstDecl = (*ZendAstDecl)(ast)
		var uses_ast *ZendAst = closure_ast.GetChild()[1]
		if uses_ast != nil {
			var uses_list *ZendAstList = ZendAstGetList(uses_ast)
			var i uint32
			for i = 0; i < uses_list.GetChildren(); i++ {
				ZendHashAddEmptyElement(&info.uses, ZendAstGetStr(uses_list.GetChild()[i]))
			}
		}
	} else if ast.GetKind() == ZEND_AST_ARROW_FUNC {

		/* For arrow functions recursively check the expression. */

		var closure_ast *ZendAstDecl = (*ZendAstDecl)(ast)
		FindImplicitBindsRecursively(info, closure_ast.GetChild()[2])
	} else if ZendAstIsSpecial(ast) == 0 {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		for i = 0; i < children; i++ {
			FindImplicitBindsRecursively(info, ast.GetChild()[i])
		}
	}
}
func FindImplicitBinds(info *ClosureInfo, params_ast *ZendAst, stmt_ast *ZendAst) {
	var param_list *ZendAstList = ZendAstGetList(params_ast)
	var i uint32
	_zendHashInit(&info.uses, param_list.GetChildren(), nil, 0)
	FindImplicitBindsRecursively(info, stmt_ast)

	/* Remove variables that are parameters */

	for i = 0; i < param_list.GetChildren(); i++ {
		var param_ast *ZendAst = param_list.GetChild()[i]
		ZendHashDel(&info.uses, ZendAstGetStr(param_ast.GetChild()[1]))
	}

	/* Remove variables that are parameters */
}
func CompileImplicitLexicalBinds(info *ClosureInfo, closure *Znode, op_array *ZendOpArray) {
	var var_name *ZendString
	var opline *ZendOp

	/* TODO We might want to use a special binding mode if varvars_used is set. */

	if &info.uses.nNumOfElements == 0 {
		return
	}
	if op_array.GetStaticVariables() == nil {
		op_array.SetStaticVariables(_zendNewArray(8))
	}
	for {
		var __ht *HashTable = &info.uses
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			var_name = _p.GetKey()
			var value *Zval = ZendHashAdd(op_array.GetStaticVariables(), var_name, &EG.uninitialized_zval)
			var offset uint32 = uint32((*byte)(value - (*byte)(op_array.GetStaticVariables().GetArData())))
			opline = ZendEmitOp(nil, 182, closure, nil)
			opline.SetOp2Type(1 << 3)
			opline.GetOp2().SetVar(LookupCv(var_name))
			opline.SetExtendedValue(offset | 2)
		}
		break
	}
}
func ZendCompileClosureUses(ast *ZendAst) {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var var_ast *ZendAst = list.GetChild()[i]
		var var_name *ZendString = ZendAstGetStr(var_ast)
		var zv Zval
		&zv.SetTypeInfo(1)
		var i int
		for i = 0; i < op_array.GetLastVar(); i++ {
			if ZendStringEquals(op_array.GetVars()[i], var_name) != 0 {
				ZendErrorNoreturn(1<<6, "Cannot use lexical variable $%s as a parameter name", var_name.GetVal())
			}
		}
		CG.SetZendLineno(ZendAstGetLineno(var_ast))
		ZendCompileStaticVarCommon(var_name, &zv, g.Cond(var_ast.GetAttr() != 0, 1, 0))
	}
}

/* }}} */

func ZendCompileImplicitClosureUses(info *ClosureInfo) {
	var var_name *ZendString
	for {
		var __ht *HashTable = &info.uses
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			var_name = _p.GetKey()
			var zv Zval
			&zv.SetTypeInfo(1)
			ZendCompileStaticVarCommon(var_name, &zv, 2)
		}
		break
	}
}
func ZendBeginMethodDecl(op_array *ZendOpArray, name *ZendString, has_body ZendBool) {
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	var in_interface ZendBool = (ce.GetCeFlags() & 1 << 0) != 0
	var in_trait ZendBool = (ce.GetCeFlags() & 1 << 1) != 0
	var is_public ZendBool = (op_array.GetFnFlags() & 1 << 0) != 0
	var is_static ZendBool = (op_array.GetFnFlags() & 1 << 4) != 0
	var lcname *ZendString
	if in_interface != 0 {
		if is_public == 0 || (op_array.GetFnFlags()&(1<<5|1<<6)) != 0 {
			ZendErrorNoreturn(1<<6, "Access type for interface method "+"%s::%s() must be omitted", ce.GetName().GetVal(), name.GetVal())
		}
		op_array.SetFnFlags(op_array.GetFnFlags() | 1<<6)
	}
	if (op_array.GetFnFlags() & 1 << 6) != 0 {
		if (op_array.GetFnFlags() & 1 << 2) != 0 {
			ZendErrorNoreturn(1<<6, "%s function %s::%s() cannot be declared private", g.Cond(in_interface != 0, "Interface", "Abstract"), ce.GetName().GetVal(), name.GetVal())
		}
		if has_body != 0 {
			ZendErrorNoreturn(1<<6, "%s function %s::%s() cannot contain body", g.Cond(in_interface != 0, "Interface", "Abstract"), ce.GetName().GetVal(), name.GetVal())
		}
		ce.SetCeFlags(ce.GetCeFlags() | 1<<4)
	} else if has_body == 0 {
		ZendErrorNoreturn(1<<6, "Non-abstract method %s::%s() must contain body", ce.GetName().GetVal(), name.GetVal())
	}
	op_array.SetScope(ce)
	op_array.SetFunctionName(ZendStringCopy(name))
	lcname = ZendStringTolowerEx(name, 0)
	lcname = ZendNewInternedString(lcname)
	if ZendHashAddPtr(&ce.function_table, lcname, op_array) == nil {
		ZendErrorNoreturn(1<<6, "Cannot redeclare %s::%s()", ce.GetName().GetVal(), name.GetVal())
	}
	if in_interface != 0 {
		if lcname.GetVal()[0] != '_' || lcname.GetVal()[1] != '_' {

		} else if lcname.GetLen() == g.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__call", g.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __call() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__callstatic", g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
			if is_public == 0 || is_static == 0 {
				ZendError(1<<1, "The magic method __callStatic() must have "+"public visibility and be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__get", g.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __get() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__set", g.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __set() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__unset", g.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__isset", g.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__tostring", g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__invoke", g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
		}
	} else {
		if in_trait == 0 && (lcname.GetLen() == ce.GetName().GetLen() && ZendBinaryStrcasecmp(lcname.GetVal(), lcname.GetLen(), ce.GetName().GetVal(), ce.GetName().GetLen()) == 0) {
			if ce.GetConstructor() == nil {
				ce.SetConstructor((*ZendFunction)(op_array))
			}
		} else if lcname.GetLen() == g.SizeOf("\"serialize\"")-1 && !(memcmp(lcname.GetVal(), "serialize", g.SizeOf("\"serialize\"")-1)) {
			ce.SetSerializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | 1<<17)
			}
		} else if lcname.GetLen() == g.SizeOf("\"unserialize\"")-1 && !(memcmp(lcname.GetVal(), "unserialize", g.SizeOf("\"unserialize\"")-1)) {
			ce.SetUnserializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | 1<<17)
			}
		} else if lcname.GetVal()[0] != '_' || lcname.GetVal()[1] != '_' {
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | 1<<17)
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__construct", g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1)) {
			ce.SetConstructor((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__destruct", g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) {
			ce.SetDestructor((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__clone", g.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) {
			ce.SetClone((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__call", g.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __call() must have "+"public visibility and cannot be static")
			}
			ce.SetCall((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__callstatic", g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
			if is_public == 0 || is_static == 0 {
				ZendError(1<<1, "The magic method __callStatic() must have "+"public visibility and be static")
			}
			ce.SetCallstatic((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__get", g.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __get() must have "+"public visibility and cannot be static")
			}
			ce.SetGet((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
		} else if lcname.GetLen() == g.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__set", g.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __set() must have "+"public visibility and cannot be static")
			}
			ce.SetSet((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
		} else if lcname.GetLen() == g.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__unset", g.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
			ce.SetUnset((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
		} else if lcname.GetLen() == g.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__isset", g.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
			ce.SetIsset((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
		} else if lcname.GetLen() == g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__tostring", g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
			ce.SetTostring((*ZendFunction)(op_array))
		} else if lcname.GetLen() == g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__invoke", g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if lcname.GetLen() == g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) {
			if is_public == 0 || is_static != 0 {
				ZendError(1<<1, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
			ce.SetDebugInfo((*ZendFunction)(op_array))
		} else if is_static == 0 {
			op_array.SetFnFlags(op_array.GetFnFlags() | 1<<17)
		}
	}
	ZendStringReleaseEx(lcname, 0)
}

/* }}} */

func ZendBeginFuncDecl(result *Znode, op_array *ZendOpArray, decl *ZendAstDecl, toplevel ZendBool) {
	var params_ast *ZendAst = decl.GetChild()[0]
	var unqualified_name *ZendString
	var name *ZendString
	var lcname *ZendString
	var key *ZendString
	var opline *ZendOp
	unqualified_name = decl.GetName()
	name = ZendPrefixWithNs(unqualified_name)
	op_array.SetFunctionName(name)
	lcname = ZendStringTolowerEx(name, 0)
	if CG.GetFileContext().GetImportsFunction() != nil {
		var import_name *ZendString = ZendHashFindPtrLc(CG.GetFileContext().GetImportsFunction(), unqualified_name.GetVal(), unqualified_name.GetLen())
		if import_name != nil && !(lcname.GetLen() == import_name.GetLen() && ZendBinaryStrcasecmp(lcname.GetVal(), lcname.GetLen(), import_name.GetVal(), import_name.GetLen()) == 0) {
			ZendErrorNoreturn(1<<6, "Cannot declare function %s "+"because the name is already in use", name.GetVal())
		}
	}
	if lcname.GetLen() == g.SizeOf("ZEND_AUTOLOAD_FUNC_NAME")-1 && !(memcmp(lcname.GetVal(), "__autoload", g.SizeOf("ZEND_AUTOLOAD_FUNC_NAME")-1)) {
		if ZendAstGetList(params_ast).GetChildren() != 1 {
			ZendErrorNoreturn(1<<6, "%s() must take exactly 1 argument", "__autoload")
		}
		ZendError(1<<13, "__autoload() is deprecated, use spl_autoload_register() instead")
	}
	if unqualified_name.GetLen() == g.SizeOf("\"assert\"")-1 && ZendBinaryStrcasecmp(unqualified_name.GetVal(), unqualified_name.GetLen(), "assert", g.SizeOf("\"assert\"")-1) == 0 {
		ZendError(1<<13, "Defining a custom assert() function is deprecated, "+"as the function has special semantics")
	}
	ZendRegisterSeenSymbol(lcname, 1<<1)
	if toplevel != 0 {
		if ZendHashAddPtr(CG.GetFunctionTable(), lcname, op_array) == nil {
			DoBindFunctionError(lcname, op_array, 1)
		}
		ZendStringReleaseEx(lcname, 0)
		return
	}

	/* Generate RTD keys until we find one that isn't in use yet. */

	key = nil
	for {
		ZendTmpStringRelease(key)
		key = ZendBuildRuntimeDefinitionKey(lcname, decl.GetStartLineno())
		if ZendHashAddPtr(CG.GetFunctionTable(), key, op_array) {
			break
		}
	}
	if (op_array.GetFnFlags() & 1 << 20) != 0 {
		opline = ZendEmitOpTmp(result, 142, nil, nil)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetOp1Type(1 << 0)
		var _c Zval
		var __z *Zval = &_c
		var __s *ZendString = key
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		opline.GetOp1().SetConstant(ZendAddLiteral(&_c))
	} else {
		opline = GetNextOp()
		opline.SetOpcode(141)
		opline.SetOp1Type(1 << 0)
		var _c Zval
		var __z *Zval = &_c
		var __s *ZendString = ZendStringCopy(lcname)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		opline.GetOp1().SetConstant(ZendAddLiteral(&_c))

		/* RTD key is placed after lcname literal in op1 */

		ZendAddLiteralString(&key)

		/* RTD key is placed after lcname literal in op1 */

	}
	ZendStringReleaseEx(lcname, 0)
}

/* }}} */

func ZendCompileFuncDecl(result *Znode, ast *ZendAst, toplevel ZendBool) {
	var decl *ZendAstDecl = (*ZendAstDecl)(ast)
	var params_ast *ZendAst = decl.GetChild()[0]
	var uses_ast *ZendAst = decl.GetChild()[1]
	var stmt_ast *ZendAst = decl.GetChild()[2]
	var return_type_ast *ZendAst = decl.GetChild()[3]
	var is_method ZendBool = decl.GetKind() == ZEND_AST_METHOD
	var orig_class_entry *ZendClassEntry = CG.GetActiveClassEntry()
	var orig_op_array *ZendOpArray = CG.GetActiveOpArray()
	var op_array *ZendOpArray = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_op_array"))
	var orig_oparray_context ZendOparrayContext
	var info ClosureInfo
	memset(&info, 0, g.SizeOf("closure_info"))
	InitOpArray(op_array, 2, 64)
	if (CG.GetCompilerOptions() & 1 << 15) != 0 {
		op_array.SetFnFlags(op_array.GetFnFlags() | 1<<10)
		op_array.SetRunTimeCachePtr(ZendMapPtrNew())
		op_array.SetStaticVariablesPtrPtr(ZendMapPtrNew())
	} else {
		op_array.SetRunTimeCachePtr(ZendArenaAlloc(&CG.arena, g.SizeOf("void *")))
		if (uintPtr(op_array.GetRunTimeCachePtr()) & 1) != 0 {
			*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetRunTimeCachePtr()-1)))) = nil
		} else {
			*(op_array.GetRunTimeCachePtr()) = nil
		}
	}
	op_array.SetFnFlags(op_array.GetFnFlags() | orig_op_array.GetFnFlags()&1<<31)
	op_array.SetFnFlags(op_array.GetFnFlags() | decl.GetFlags())
	op_array.SetLineStart(decl.GetStartLineno())
	op_array.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		op_array.SetDocComment(ZendStringCopy(decl.GetDocComment()))
	}
	if decl.GetKind() == ZEND_AST_CLOSURE || decl.GetKind() == ZEND_AST_ARROW_FUNC {
		op_array.SetFnFlags(op_array.GetFnFlags() | 1<<20)
	}
	if is_method != 0 {
		var has_body ZendBool = stmt_ast != nil
		ZendBeginMethodDecl(op_array, decl.GetName(), has_body)
	} else {
		ZendBeginFuncDecl(result, op_array, decl, toplevel)
		if decl.GetKind() == ZEND_AST_ARROW_FUNC {
			FindImplicitBinds(&info, params_ast, stmt_ast)
			CompileImplicitLexicalBinds(&info, result, op_array)
		} else if uses_ast != nil {
			ZendCompileClosureBinding(result, op_array, uses_ast)
		}
	}
	CG.SetActiveOpArray(op_array)

	/* Do not leak the class scope into free standing functions, even if they are dynamically
	 * defined inside a class method. This is necessary for correct handling of magic constants.
	 * For example __CLASS__ should always be "" inside a free standing function. */

	if decl.GetKind() == ZEND_AST_FUNC_DECL {
		CG.SetActiveClassEntry(nil)
	}
	if toplevel != 0 {
		op_array.SetFnFlags(op_array.GetFnFlags() | 1<<9)
	}
	ZendOparrayContextBegin(&orig_oparray_context)
	if (CG.GetCompilerOptions() & 1 << 0) != 0 {
		var opline_ext *ZendOp = ZendEmitOp(nil, 104, nil, nil)
		opline_ext.SetLineno(decl.GetStartLineno())
	}

	/* Push a separator to the loop variable stack */

	var dummy_var ZendLoopVar
	dummy_var.SetOpcode(62)
	ZendStackPush(&CG.loop_var_stack, any(&dummy_var))
	ZendCompileParams(params_ast, return_type_ast)
	if (CG.GetActiveOpArray().GetFnFlags() & 1 << 24) != 0 {
		ZendMarkFunctionAsGenerator()
		ZendEmitOp(nil, 139, nil, nil)
	}
	if decl.GetKind() == ZEND_AST_ARROW_FUNC {
		ZendCompileImplicitClosureUses(&info)
		ZendHashDestroy(&info.uses)
	} else if uses_ast != nil {
		ZendCompileClosureUses(uses_ast)
	}
	ZendCompileStmt(stmt_ast)
	if is_method != 0 {
		ZendCheckMagicMethodImplementation(CG.GetActiveClassEntry(), (*ZendFunction)(op_array), 1<<6)
	}

	/* put the implicit return on the really last line */

	CG.SetZendLineno(decl.GetEndLineno())
	ZendDoExtendedStmt()
	ZendEmitFinalReturn(0)
	PassTwo(CG.GetActiveOpArray())
	ZendOparrayContextEnd(&orig_oparray_context)

	/* Pop the loop variable stack separator */

	ZendStackDelTop(&CG.loop_var_stack)
	CG.SetActiveOpArray(orig_op_array)
	CG.SetActiveClassEntry(orig_class_entry)
}

/* }}} */

func ZendCompilePropDecl(ast *ZendAst, type_ast *ZendAst, flags uint32) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	var i uint32
	var children uint32 = list.GetChildren()
	if (ce.GetCeFlags() & 1 << 0) != 0 {
		ZendErrorNoreturn(1<<6, "Interfaces may not include member variables")
	}
	if (flags & 1 << 6) != 0 {
		ZendErrorNoreturn(1<<6, "Properties cannot be declared abstract")
	}
	for i = 0; i < children; i++ {
		var prop_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = prop_ast.GetChild()[0]
		var value_ast *ZendAst = prop_ast.GetChild()[1]
		var doc_comment_ast *ZendAst = prop_ast.GetChild()[2]
		var name *ZendString = ZvalMakeInternedString(ZendAstGetZval(name_ast))
		var doc_comment *ZendString = nil
		var value_zv Zval
		var type_ ZendType = 0
		if type_ast != nil {
			type_ = ZendCompileTypename(type_ast, 0)
			if type_>>2 == 19 || type_>>2 == 17 {
				ZendErrorNoreturn(1<<6, "Property %s::$%s cannot have type %s", ce.GetName().GetVal(), name.GetVal(), ZendGetTypeByConst(type_>>2))
			}
		}

		/* Doc comment has been appended as last element in ZEND_AST_PROP_ELEM ast */

		if doc_comment_ast != nil {
			doc_comment = ZendStringCopy(ZendAstGetStr(doc_comment_ast))
		}
		if (flags & 1 << 5) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot declare property %s::$%s final, "+"the final modifier is allowed only for methods and classes", ce.GetName().GetVal(), name.GetVal())
		}
		if ZendHashExists(&ce.properties_info, name) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot redeclare %s::$%s", ce.GetName().GetVal(), name.GetVal())
		}
		if value_ast != nil {
			ZendConstExprToZval(&value_zv, value_ast)
			if type_ > 0x3 && value_zv.GetType() != 11 {
				if value_zv.GetType() == 1 {
					if (type_ & 0x1) == 0 {
						var name *byte = g.CondF(type_ > 0x3ff, func() []byte { return (*ZendString)(type_ & ^0x3).GetVal() }, func() *byte { return ZendGetTypeByConst(type_ >> 2) })
						ZendErrorNoreturn(1<<6, "Default value for property of type %s may not be null. "+"Use the nullable type ?%s to allow null default value", name, name)
					}
				} else if type_ > 0x3ff {
					ZendErrorNoreturn(1<<6, "Property of type %s may not have default value", (*ZendString)(type_ & ^0x3).GetVal())
				} else if type_>>2 == 7 || type_>>2 == 18 {
					if value_zv.GetType() != 7 {
						ZendErrorNoreturn(1<<6, "Default value for property of type %s can only be an array", ZendGetTypeByConst(type_>>2))
					}
				} else if type_>>2 == 5 {
					if value_zv.GetType() != 5 && value_zv.GetType() != 4 {
						ZendErrorNoreturn(1<<6, "Default value for property of type float can only be float or int")
					}
					ConvertToDouble(&value_zv)
				} else if !(type_>>2 == value_zv.GetType() || type_>>2 == 16 && (value_zv.GetType() == 3 || value_zv.GetType() == 2)) {
					ZendErrorNoreturn(1<<6, "Default value for property of type %s can only be %s", ZendGetTypeByConst(type_>>2), ZendGetTypeByConst(type_>>2))
				}
			}
		} else if type_ <= 0x3 {
			&value_zv.SetTypeInfo(1)
		} else {
			&value_zv.SetTypeInfo(0)
		}
		ZendDeclareTypedProperty(ce, name, &value_zv, flags, doc_comment, type_)
	}
}

/* }}} */

func ZendCompilePropGroup(list *ZendAst) {
	var type_ast *ZendAst = list.GetChild()[0]
	var prop_ast *ZendAst = list.GetChild()[1]
	ZendCompilePropDecl(prop_ast, type_ast, list.GetAttr())
}

/* }}} */

func ZendCompileClassConstDecl(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	var i uint32
	if (ce.GetCeFlags() & 1 << 1) != 0 {
		ZendErrorNoreturn(1<<6, "Traits cannot have constants")
		return
	}
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = const_ast.GetChild()[0]
		var value_ast *ZendAst = const_ast.GetChild()[1]
		var doc_comment_ast *ZendAst = const_ast.GetChild()[2]
		var name *ZendString = ZvalMakeInternedString(ZendAstGetZval(name_ast))
		var doc_comment *ZendString = g.CondF1(doc_comment_ast != nil, func() *ZendString { return ZendStringCopy(ZendAstGetStr(doc_comment_ast)) }, nil)
		var value_zv Zval
		if (ast.GetAttr() & (1<<4 | 1<<6 | 1<<5)) != 0 {
			if (ast.GetAttr() & 1 << 4) != 0 {
				ZendErrorNoreturn(1<<6, "Cannot use 'static' as constant modifier")
			} else if (ast.GetAttr() & 1 << 6) != 0 {
				ZendErrorNoreturn(1<<6, "Cannot use 'abstract' as constant modifier")
			} else if (ast.GetAttr() & 1 << 5) != 0 {
				ZendErrorNoreturn(1<<6, "Cannot use 'final' as constant modifier")
			}
		}
		ZendConstExprToZval(&value_zv, value_ast)
		ZendDeclareClassConstantEx(ce, name, &value_zv, ast.GetAttr(), doc_comment)
	}
}

/* }}} */

func ZendCompileMethodRef(ast *ZendAst, method_ref *ZendTraitMethodReference) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	method_ref.SetMethodName(ZendStringCopy(ZendAstGetStr(method_ast)))
	if class_ast != nil {
		method_ref.SetClassName(ZendResolveClassNameAst(class_ast))
	} else {
		method_ref.SetClassName(nil)
	}
}

/* }}} */

func ZendCompileTraitPrecedence(ast *ZendAst) {
	var method_ref_ast *ZendAst = ast.GetChild()[0]
	var insteadof_ast *ZendAst = ast.GetChild()[1]
	var insteadof_list *ZendAstList = ZendAstGetList(insteadof_ast)
	var i uint32
	var precedence *ZendTraitPrecedence = _emalloc(g.SizeOf("zend_trait_precedence") + (insteadof_list.GetChildren()-1)*g.SizeOf("zend_string *"))
	ZendCompileMethodRef(method_ref_ast, &precedence.trait_method)
	precedence.SetNumExcludes(insteadof_list.GetChildren())
	for i = 0; i < insteadof_list.GetChildren(); i++ {
		var name_ast *ZendAst = insteadof_list.GetChild()[i]
		precedence.GetExcludeClassNames()[i] = ZendResolveClassNameAst(name_ast)
	}
	ZendAddToList(&CG.active_class_entry.GetTraitPrecedences(), precedence)
}

/* }}} */

func ZendCompileTraitAlias(ast *ZendAst) {
	var method_ref_ast *ZendAst = ast.GetChild()[0]
	var alias_ast *ZendAst = ast.GetChild()[1]
	var modifiers uint32 = ast.GetAttr()
	var alias *ZendTraitAlias
	if modifiers == 1<<4 {
		ZendErrorNoreturn(1<<6, "Cannot use 'static' as method modifier")
	} else if modifiers == 1<<6 {
		ZendErrorNoreturn(1<<6, "Cannot use 'abstract' as method modifier")
	} else if modifiers == 1<<5 {
		ZendErrorNoreturn(1<<6, "Cannot use 'final' as method modifier")
	}
	alias = _emalloc(g.SizeOf("zend_trait_alias"))
	ZendCompileMethodRef(method_ref_ast, &alias.trait_method)
	alias.SetModifiers(modifiers)
	if alias_ast != nil {
		alias.SetAlias(ZendStringCopy(ZendAstGetStr(alias_ast)))
	} else {
		alias.SetAlias(nil)
	}
	ZendAddToList(&CG.active_class_entry.GetTraitAliases(), alias)
}

/* }}} */

func ZendCompileUseTrait(ast *ZendAst) {
	var traits *ZendAstList = ZendAstGetList(ast.GetChild()[0])
	var adaptations *ZendAstList = g.CondF1(ast.GetChild()[1] != nil, func() *ZendAstList { return ZendAstGetList(ast.GetChild()[1]) }, nil)
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	var i uint32
	ce.SetCeFlags(ce.GetCeFlags() | 1<<15)
	ce.SetTraitNames(_erealloc(ce.GetTraitNames(), g.SizeOf("zend_class_name")*(ce.GetNumTraits()+traits.GetChildren())))
	for i = 0; i < traits.GetChildren(); i++ {
		var trait_ast *ZendAst = traits.GetChild()[i]
		var name *ZendString = ZendAstGetStr(trait_ast)
		if (ce.GetCeFlags() & 1 << 0) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot use traits inside of interfaces. "+"%s is used in %s", name.GetVal(), ce.GetName().GetVal())
		}
		switch ZendGetClassFetchType(name) {
		case 1:

		case 2:

		case 3:
			ZendErrorNoreturn(1<<6, "Cannot use '%s' as trait name "+"as it is reserved", name.GetVal())
			break
		}
		ce.GetTraitNames()[ce.GetNumTraits()].SetName(ZendResolveClassNameAst(trait_ast))
		ce.GetTraitNames()[ce.GetNumTraits()].SetLcName(ZendStringTolowerEx(ce.GetTraitNames()[ce.GetNumTraits()].GetName(), 0))
		ce.GetNumTraits()++
	}
	if adaptations == nil {
		return
	}
	for i = 0; i < adaptations.GetChildren(); i++ {
		var adaptation_ast *ZendAst = adaptations.GetChild()[i]
		switch adaptation_ast.GetKind() {
		case ZEND_AST_TRAIT_PRECEDENCE:
			ZendCompileTraitPrecedence(adaptation_ast)
			break
		case ZEND_AST_TRAIT_ALIAS:
			ZendCompileTraitAlias(adaptation_ast)
			break
		default:
			break
		}
	}
}

/* }}} */

func ZendCompileImplements(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	var interface_names *ZendClassName
	var i uint32
	interface_names = _emalloc(g.SizeOf("zend_class_name") * list.GetChildren())
	for i = 0; i < list.GetChildren(); i++ {
		var class_ast *ZendAst = list.GetChild()[i]
		var name *ZendString = ZendAstGetStr(class_ast)
		if ZendIsConstDefaultClassRef(class_ast) == 0 {
			_efree(interface_names)
			ZendErrorNoreturn(1<<6, "Cannot use '%s' as interface name as it is reserved", name.GetVal())
		}
		interface_names[i].SetName(ZendResolveClassNameAst(class_ast))
		interface_names[i].SetLcName(ZendStringTolowerEx(interface_names[i].GetName(), 0))
	}
	ce.SetCeFlags(ce.GetCeFlags() | 1<<14)
	ce.SetNumInterfaces(list.GetChildren())
	ce.interface_names = interface_names
}

/* }}} */

func ZendGenerateAnonClassName(start_lineno uint32) *ZendString {
	var filename *ZendString = CG.GetActiveOpArray().GetFilename()
	var result *ZendString = ZendStrpprintf(0, "class@anonymous%c%s:%"+"u"+"$%"+PRIx32, '0', filename.GetVal(), start_lineno, g.PostInc(&(CG.GetRtdKeyCounter())))
	return ZendNewInternedString(result)
}

/* }}} */

func ZendCompileClassDecl(ast *ZendAst, toplevel ZendBool) *ZendOp {
	var decl *ZendAstDecl = (*ZendAstDecl)(ast)
	var extends_ast *ZendAst = decl.GetChild()[0]
	var implements_ast *ZendAst = decl.GetChild()[1]
	var stmt_ast *ZendAst = decl.GetChild()[2]
	var name *ZendString
	var lcname *ZendString
	var ce *ZendClassEntry = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_class_entry"))
	var opline *ZendOp
	var original_ce *ZendClassEntry = CG.GetActiveClassEntry()
	if (decl.GetFlags() & 1 << 2) == 0 {
		var unqualified_name *ZendString = decl.GetName()
		if CG.GetActiveClassEntry() != nil {
			ZendErrorNoreturn(1<<6, "Class declarations may not be nested")
		}
		ZendAssertValidClassName(unqualified_name)
		name = ZendPrefixWithNs(unqualified_name)
		name = ZendNewInternedString(name)
		lcname = ZendStringTolowerEx(name, 0)
		if CG.GetFileContext().GetImports() != nil {
			var import_name *ZendString = ZendHashFindPtrLc(CG.GetFileContext().GetImports(), unqualified_name.GetVal(), unqualified_name.GetLen())
			if import_name != nil && !(lcname.GetLen() == import_name.GetLen() && ZendBinaryStrcasecmp(lcname.GetVal(), lcname.GetLen(), import_name.GetVal(), import_name.GetLen()) == 0) {
				ZendErrorNoreturn(1<<6, "Cannot declare class %s "+"because the name is already in use", name.GetVal())
			}
		}
		ZendRegisterSeenSymbol(lcname, 1<<0)
	} else {

		/* Find an anon class name that is not in use yet. */

		name = nil
		lcname = nil
		for {
			ZendTmpStringRelease(name)
			ZendTmpStringRelease(lcname)
			name = ZendGenerateAnonClassName(decl.GetStartLineno())
			lcname = ZendStringTolowerEx(name, 0)
			if ZendHashExists(CG.GetClassTable(), lcname) == 0 {
				break
			}
		}
	}
	lcname = ZendNewInternedString(lcname)
	ce.SetType(2)
	ce.SetName(name)
	ZendInitializeClassData(ce, 1)
	if (CG.GetCompilerOptions() & 1 << 15) != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | 1<<10)
		ce.SetStaticMembersTablePtr(ZendMapPtrNew())
	}
	ce.SetCeFlags(ce.GetCeFlags() | decl.GetFlags())
	ce.SetFilename(ZendGetCompiledFilename())
	ce.SetLineStart(decl.GetStartLineno())
	ce.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		ce.SetDocComment(ZendStringCopy(decl.GetDocComment()))
	}
	if (decl.GetFlags() & 1 << 2) != 0 {

		/* Serialization is not supported for anonymous classes */

		ce.SetSerialize(ZendClassSerializeDeny)
		ce.SetUnserialize(ZendClassUnserializeDeny)
	}
	if extends_ast != nil {
		var extends_node Znode
		var extends_name *ZendString
		if ZendIsConstDefaultClassRef(extends_ast) == 0 {
			extends_name = ZendAstGetStr(extends_ast)
			ZendErrorNoreturn(1<<6, "Cannot use '%s' as class name as it is reserved", extends_name.GetVal())
		}
		ZendCompileExpr(&extends_node, extends_ast)
		if extends_node.GetOpType() != 1<<0 || extends_node.GetConstant().GetType() != 6 {
			ZendErrorNoreturn(1<<6, "Illegal class name")
		}
		extends_name = extends_node.GetConstant().GetValue().GetStr()
		ce.parent_name = ZendResolveClassName(extends_name, g.CondF1(extends_ast.GetKind() == ZEND_AST_ZVAL, func() ZendAstAttr { return extends_ast.GetAttr() }, 0))
		ZendStringReleaseEx(extends_name, 0)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<13)
	}
	CG.SetActiveClassEntry(ce)
	ZendCompileStmt(stmt_ast)

	/* Reset lineno for final opcodes and errors */

	CG.SetZendLineno(ast.GetLineno())
	if (ce.GetCeFlags() & 1 << 15) == 0 {

		/* For traits this check is delayed until after trait binding */

		ZendCheckDeprecatedConstructor(ce)

		/* For traits this check is delayed until after trait binding */

	}
	if ce.GetConstructor() != nil {
		ce.GetConstructor().SetFnFlags(ce.GetConstructor().GetFnFlags() | 1<<28)
		if (ce.GetConstructor().GetFnFlags() & 1 << 4) != 0 {
			ZendErrorNoreturn(1<<6, "Constructor %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
		if (ce.GetConstructor().GetFnFlags() & 1 << 13) != 0 {
			ZendErrorNoreturn(1<<6, "Constructor %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
	}
	if ce.GetDestructor() != nil {
		ce.GetDestructor().SetFnFlags(ce.GetDestructor().GetFnFlags() | 1<<29)
		if (ce.GetDestructor().GetFnFlags() & 1 << 4) != 0 {
			ZendErrorNoreturn(1<<6, "Destructor %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetDestructor().GetFunctionName().GetVal())
		} else if (ce.GetDestructor().GetFnFlags() & 1 << 13) != 0 {
			ZendErrorNoreturn(1<<6, "Destructor %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetDestructor().GetFunctionName().GetVal())
		}
	}
	if ce.GetClone() != nil {
		if (ce.GetClone().GetFnFlags() & 1 << 4) != 0 {
			ZendErrorNoreturn(1<<6, "Clone method %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetClone().GetFunctionName().GetVal())
		} else if (ce.GetClone().GetFnFlags() & 1 << 13) != 0 {
			ZendErrorNoreturn(1<<6, "Clone method %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetClone().GetFunctionName().GetVal())
		}
	}
	if implements_ast != nil {
		ZendCompileImplements(implements_ast)
	}
	if (ce.GetCeFlags() & (1<<4 | 1<<0 | 1<<1 | 1<<6)) == 1<<4 {
		ZendVerifyAbstractClass(ce)
	}
	CG.SetActiveClassEntry(original_ce)
	if toplevel != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | 1<<9)
	}
	if toplevel != 0 && (ce.GetCeFlags()&(1<<14|1<<15)) == 0 && (CG.GetCompilerOptions()&1<<15) == 0 {
		if extends_ast != nil {
			var parent_ce *ZendClassEntry = ZendLookupClassEx(ce.parent_name, nil, 0x80)
			if parent_ce != nil && (parent_ce.GetType() != 1 || (CG.GetCompilerOptions()&1<<4) == 0) && (parent_ce.GetType() != 2 || (CG.GetCompilerOptions()&1<<13) == 0 || parent_ce.GetFilename() == ce.GetFilename()) {
				CG.SetZendLineno(decl.GetEndLineno())
				if ZendTryEarlyBind(ce, parent_ce, lcname, nil) != 0 {
					CG.SetZendLineno(ast.GetLineno())
					ZendStringRelease(lcname)
					return nil
				}
				CG.SetZendLineno(ast.GetLineno())
			}
		} else {
			if ZendHashAddPtr(CG.GetClassTable(), lcname, ce) != nil {
				ZendStringRelease(lcname)
				ZendBuildPropertiesInfoTable(ce)
				ce.SetCeFlags(ce.GetCeFlags() | 1<<3)
				return nil
			}
		}
	}
	opline = GetNextOp()
	if ce.parent_name {

		/* Lowercased parent name */

		var lc_parent_name *ZendString = ZendStringTolowerEx(ce.parent_name, 0)
		opline.SetOp2Type(1 << 0)
		var _c Zval
		var __z *Zval = &_c
		var __s *ZendString = lc_parent_name
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		opline.GetOp2().SetConstant(ZendAddLiteral(&_c))
	}
	opline.SetOp1Type(1 << 0)
	var _c Zval
	var __z *Zval = &_c
	var __s *ZendString = lcname
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	opline.GetOp1().SetConstant(ZendAddLiteral(&_c))
	if (decl.GetFlags() & 1 << 2) != 0 {
		opline.SetOpcode(146)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetResultType(1 << 2)
		opline.GetResult().SetVar(GetTemporaryVariable())
		if !(ZendHashAddPtr(CG.GetClassTable(), lcname, ce)) {

			/* We checked above that the class name is not used. This really shouldn't happen. */

			ZendErrorNoreturn(1<<0, "Runtime definition key collision for %s. This is a bug", name.GetVal())

			/* We checked above that the class name is not used. This really shouldn't happen. */

		}
	} else {

		/* Generate RTD keys until we find one that isn't in use yet. */

		var key *ZendString = nil
		for {
			ZendTmpStringRelease(key)
			key = ZendBuildRuntimeDefinitionKey(lcname, decl.GetStartLineno())
			if ZendHashAddPtr(CG.GetClassTable(), key, ce) {
				break
			}
		}

		/* RTD key is placed after lcname literal in op1 */

		ZendAddLiteralString(&key)
		opline.SetOpcode(144)
		if extends_ast != nil && toplevel != 0 && (CG.GetCompilerOptions()&1<<5) != 0 && (ce.GetCeFlags()&(1<<14|1<<15)) == 0 {
			CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<16)
			opline.SetOpcode(145)
			opline.SetExtendedValue(ZendAllocCacheSlot())
			opline.SetResultType(0)
			opline.GetResult().SetOplineNum(-1)
		}
	}
	return opline
}

/* }}} */

func ZendGetImportHt(type_ uint32) *HashTable {
	switch type_ {
	case 1 << 0:
		if CG.GetFileContext().GetImports() == nil {
			CG.GetFileContext().SetImports(_emalloc(g.SizeOf("HashTable")))
			_zendHashInit(CG.GetFileContext().GetImports(), 8, StrDtor, 0)
		}
		return CG.GetFileContext().GetImports()
	case 1 << 1:
		if CG.GetFileContext().GetImportsFunction() == nil {
			CG.GetFileContext().SetImportsFunction(_emalloc(g.SizeOf("HashTable")))
			_zendHashInit(CG.GetFileContext().GetImportsFunction(), 8, StrDtor, 0)
		}
		return CG.GetFileContext().GetImportsFunction()
	case 1 << 2:
		if CG.GetFileContext().GetImportsConst() == nil {
			CG.GetFileContext().SetImportsConst(_emalloc(g.SizeOf("HashTable")))
			_zendHashInit(CG.GetFileContext().GetImportsConst(), 8, StrDtor, 0)
		}
		return CG.GetFileContext().GetImportsConst()
	default:
		break
	}
	return nil
}

/* }}} */

func ZendGetUseTypeStr(type_ uint32) *byte {
	switch type_ {
	case 1 << 0:
		return ""
	case 1 << 1:
		return " function"
	case 1 << 2:
		return " const"
	default:
		break
	}
	return " unknown"
}

/* }}} */

func ZendCheckAlreadyInUse(type_ uint32, old_name *ZendString, new_name *ZendString, check_name *ZendString) {
	if old_name.GetLen() == check_name.GetLen() && ZendBinaryStrcasecmp(old_name.GetVal(), old_name.GetLen(), check_name.GetVal(), check_name.GetLen()) == 0 {
		return
	}
	ZendErrorNoreturn(1<<6, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), old_name.GetVal(), new_name.GetVal())
}

/* }}} */

func ZendCompileUse(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var current_ns *ZendString = CG.GetFileContext().GetCurrentNamespace()
	var type_ uint32 = ast.GetAttr()
	var current_import *HashTable = ZendGetImportHt(type_)
	var case_sensitive ZendBool = type_ == 1<<2
	for i = 0; i < list.GetChildren(); i++ {
		var use_ast *ZendAst = list.GetChild()[i]
		var old_name_ast *ZendAst = use_ast.GetChild()[0]
		var new_name_ast *ZendAst = use_ast.GetChild()[1]
		var old_name *ZendString = ZendAstGetStr(old_name_ast)
		var new_name *ZendString
		var lookup_name *ZendString
		if new_name_ast != nil {
			new_name = ZendStringCopy(ZendAstGetStr(new_name_ast))
		} else {
			var unqualified_name *byte
			var unqualified_name_len int
			if ZendGetUnqualifiedName(old_name, &unqualified_name, &unqualified_name_len) != 0 {

				/* The form "use A\B" is equivalent to "use A\B as B" */

				new_name = ZendStringInit(unqualified_name, unqualified_name_len, 0)

				/* The form "use A\B" is equivalent to "use A\B as B" */

			} else {
				new_name = ZendStringCopy(old_name)
				if current_ns == nil {
					if type_ == T_CLASS && (new_name.GetLen() == g.SizeOf("\"strict\"")-1 && !(memcmp(new_name.GetVal(), "strict", g.SizeOf("\"strict\"")-1))) {
						ZendErrorNoreturn(1<<6, "You seem to be trying to use a different language...")
					}
					ZendError(1<<1, "The use statement with non-compound name '%s' "+"has no effect", new_name.GetVal())
				}
			}
		}
		if case_sensitive != 0 {
			lookup_name = ZendStringCopy(new_name)
		} else {
			lookup_name = ZendStringTolowerEx(new_name, 0)
		}
		if type_ == 1<<0 && ZendIsReservedClassName(new_name) != 0 {
			ZendErrorNoreturn(1<<6, "Cannot use %s as %s because '%s' "+"is a special class name", old_name.GetVal(), new_name.GetVal(), new_name.GetVal())
		}
		if current_ns != nil {
			var ns_name *ZendString = ZendStringAlloc(current_ns.GetLen()+1+new_name.GetLen(), 0)
			ZendStrTolowerCopy(ns_name.GetVal(), current_ns.GetVal(), current_ns.GetLen())
			ns_name.GetVal()[current_ns.GetLen()] = '\\'
			memcpy(ns_name.GetVal()+current_ns.GetLen()+1, lookup_name.GetVal(), lookup_name.GetLen()+1)
			if ZendHaveSeenSymbol(ns_name, type_) != 0 {
				ZendCheckAlreadyInUse(type_, old_name, new_name, ns_name)
			}
			ZendStringEfree(ns_name)
		} else {
			if ZendHaveSeenSymbol(lookup_name, type_) != 0 {
				ZendCheckAlreadyInUse(type_, old_name, new_name, lookup_name)
			}
		}
		ZendStringAddref(old_name)
		old_name = ZendNewInternedString(old_name)
		if !(ZendHashAddPtr(current_import, lookup_name, old_name)) {
			ZendErrorNoreturn(1<<6, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), old_name.GetVal(), new_name.GetVal())
		}
		ZendStringReleaseEx(lookup_name, 0)
		ZendStringReleaseEx(new_name, 0)
	}
}

/* }}} */

func ZendCompileGroupUse(ast *ZendAst) {
	var i uint32
	var ns *ZendString = ZendAstGetStr(ast.GetChild()[0])
	var list *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	for i = 0; i < list.GetChildren(); i++ {
		var inline_use *ZendAst
		var use *ZendAst = list.GetChild()[i]
		var name_zval *Zval = ZendAstGetZval(use.GetChild()[0])
		var name *ZendString = name_zval.GetValue().GetStr()
		var compound_ns *ZendString = ZendConcatNames(ns.GetVal(), ns.GetLen(), name.GetVal(), name.GetLen())
		ZendStringReleaseEx(name, 0)
		var __z *Zval = name_zval
		var __s *ZendString = compound_ns
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		inline_use = ZendAstCreateList1(ZEND_AST_USE, use)
		if ast.GetAttr() != 0 {
			inline_use.SetAttr(ast.GetAttr())
		} else {
			inline_use.SetAttr(use.GetAttr())
		}
		ZendCompileUse(inline_use)
	}
}

/* }}} */

func ZendCompileConstDecl(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = const_ast.GetChild()[0]
		var value_ast *ZendAst = const_ast.GetChild()[1]
		var unqualified_name *ZendString = ZendAstGetStr(name_ast)
		var name *ZendString
		var name_node Znode
		var value_node Znode
		var value_zv *Zval = &value_node.u.constant
		value_node.SetOpType(1 << 0)
		ZendConstExprToZval(value_zv, value_ast)
		if ZendLookupReservedConst(unqualified_name.GetVal(), unqualified_name.GetLen()) != nil {
			ZendErrorNoreturn(1<<6, "Cannot redeclare constant '%s'", unqualified_name.GetVal())
		}
		name = ZendPrefixWithNs(unqualified_name)
		name = ZendNewInternedString(name)
		if CG.GetFileContext().GetImportsConst() != nil {
			var import_name *ZendString = ZendHashFindPtr(CG.GetFileContext().GetImportsConst(), unqualified_name)
			if import_name != nil && ZendStringEquals(import_name, name) == 0 {
				ZendErrorNoreturn(1<<6, "Cannot declare const %s because "+"the name is already in use", name.GetVal())
			}
		}
		name_node.SetOpType(1 << 0)
		var __z *Zval = &name_node.u.constant
		var __s *ZendString = name
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		ZendEmitOp(nil, 143, &name_node, &value_node)
		ZendRegisterSeenSymbol(name, 1<<2)
	}
}

/* }}}*/

func ZendCompileNamespace(ast *ZendAst) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var name *ZendString
	var with_bracket ZendBool = stmt_ast != nil

	/* handle mixed syntax declaration or nested namespaces */

	if CG.GetFileContext().GetHasBracketedNamespaces() == 0 {
		if CG.GetFileContext().GetCurrentNamespace() != nil {

			/* previous namespace declarations were unbracketed */

			if with_bracket != 0 {
				ZendErrorNoreturn(1<<6, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
			}

			/* previous namespace declarations were unbracketed */

		}
	} else {

		/* previous namespace declarations were bracketed */

		if with_bracket == 0 {
			ZendErrorNoreturn(1<<6, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
		} else if CG.GetFileContext().GetCurrentNamespace() != nil || CG.GetFileContext().GetInNamespace() != 0 {
			ZendErrorNoreturn(1<<6, "Namespace declarations cannot be nested")
		}

		/* previous namespace declarations were bracketed */

	}
	if (with_bracket == 0 && CG.GetFileContext().GetCurrentNamespace() == nil || with_bracket != 0 && CG.GetFileContext().GetHasBracketedNamespaces() == 0) && CG.GetActiveOpArray().GetLast() > 0 {

		/* ignore ZEND_EXT_STMT and ZEND_TICKS */

		var num uint32 = CG.GetActiveOpArray().GetLast()
		for num > 0 && (CG.GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == 101 || CG.GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == 105) {
			num--
		}
		if num > 0 {
			ZendErrorNoreturn(1<<6, "Namespace declaration statement has to be "+"the very first statement or after any declare call in the script")
		}
	}
	if CG.GetFileContext().GetCurrentNamespace() != nil {
		ZendStringReleaseEx(CG.GetFileContext().GetCurrentNamespace(), 0)
	}
	if name_ast != nil {
		name = ZendAstGetStr(name_ast)
		if 0 != ZendGetClassFetchType(name) {
			ZendErrorNoreturn(1<<6, "Cannot use '%s' as namespace name", name.GetVal())
		}
		CG.GetFileContext().SetCurrentNamespace(ZendStringCopy(name))
	} else {
		CG.GetFileContext().SetCurrentNamespace(nil)
	}
	ZendResetImportTables()
	CG.GetFileContext().SetInNamespace(1)
	if with_bracket != 0 {
		CG.GetFileContext().SetHasBracketedNamespaces(1)
	}
	if stmt_ast != nil {
		ZendCompileTopStmt(stmt_ast)
		ZendEndNamespace()
	}
}

/* }}} */

func ZendCompileHaltCompiler(ast *ZendAst) {
	var offset_ast *ZendAst = ast.GetChild()[0]
	var offset ZendLong = ZendAstGetZval(offset_ast).GetValue().GetLval()
	var filename *ZendString
	var name *ZendString
	var const_name []byte = "__COMPILER_HALT_OFFSET__"
	if CG.GetFileContext().GetHasBracketedNamespaces() != 0 && CG.GetFileContext().GetInNamespace() != 0 {
		ZendErrorNoreturn(1<<6, "__HALT_COMPILER() can only be used from the outermost scope")
	}
	filename = ZendGetCompiledFilename()
	name = ZendManglePropertyName(const_name, g.SizeOf("const_name")-1, filename.GetVal(), filename.GetLen(), 0)
	ZendRegisterLongConstant(name.GetVal(), name.GetLen(), offset, 1<<0, 0)
	ZendStringReleaseEx(name, 0)
}

/* }}} */

func ZendTryCtEvalMagicConst(zv *Zval, ast *ZendAst) ZendBool {
	var op_array *ZendOpArray = CG.GetActiveOpArray()
	var ce *ZendClassEntry = CG.GetActiveClassEntry()
	switch ast.GetAttr() {
	case T_LINE:
		var __z *Zval = zv
		__z.GetValue().SetLval(ast.GetLineno())
		__z.SetTypeInfo(4)
		break
	case T_FILE:
		var __z *Zval = zv
		var __s *ZendString = CG.GetCompiledFilename()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	case T_DIR:
		var filename *ZendString = CG.GetCompiledFilename()
		var dirname *ZendString = ZendStringInit(filename.GetVal(), filename.GetLen(), 0)
		dirname.SetLen(ZendDirname(dirname.GetVal(), dirname.GetLen()))
		if strcmp(dirname.GetVal(), ".") == 0 {
			dirname = ZendStringExtend(dirname, 256, 0)
			void(getcwd(dirname.GetVal(), 256))
			dirname.SetLen(strlen(dirname.GetVal()))
		}
		var __z *Zval = zv
		var __s *ZendString = dirname
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	case T_FUNC_C:
		if op_array != nil && op_array.GetFunctionName() != nil {
			var __z *Zval = zv
			var __s *ZendString = op_array.GetFunctionName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = zv
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	case T_METHOD_C:

		/* Detect whether we are directly inside a class (e.g. a class constant) and treat
		 * this as not being inside a function. */

		if op_array != nil && ce != nil && op_array.GetScope() == nil && (op_array.GetFnFlags()&1<<20) == 0 {
			op_array = nil
		}
		if op_array != nil && op_array.GetFunctionName() != nil {
			if op_array.GetScope() != nil {
				var __z *Zval = zv
				var __s *ZendString = ZendConcat3(op_array.GetScope().GetName().GetVal(), op_array.GetScope().GetName().GetLen(), "::", 2, op_array.GetFunctionName().GetVal(), op_array.GetFunctionName().GetLen())
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6 | 1<<0<<8)
			} else {
				var __z *Zval = zv
				var __s *ZendString = op_array.GetFunctionName()
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
		} else {
			var __z *Zval = zv
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	case T_CLASS_C:
		if ce != nil {
			if (ce.GetCeFlags() & 1 << 1) != 0 {
				return 0
			} else {
				var __z *Zval = zv
				var __s *ZendString = ce.GetName()
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
		} else {
			var __z *Zval = zv
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	case T_TRAIT_C:
		if ce != nil && (ce.GetCeFlags()&1<<1) != 0 {
			var __z *Zval = zv
			var __s *ZendString = ce.GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = zv
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	case T_NS_C:
		if CG.GetFileContext().GetCurrentNamespace() != nil {
			var __z *Zval = zv
			var __s *ZendString = CG.GetFileContext().GetCurrentNamespace()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = zv
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	default:
		break
	}
	return 1
}

/* }}} */

func ZendBinaryOpProducesNumericStringError(opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	if !(opcode == 1 || opcode == 2 || opcode == 3 || opcode == 4 || opcode == 12 || opcode == 5 || opcode == 6 || opcode == 7 || opcode == 9 || opcode == 10 || opcode == 11) {
		return 0
	}

	/* While basic arithmetic operators always produce numeric string errors,
	 * bitwise operators don't produce errors if both operands are strings */

	if (opcode == 9 || opcode == 10 || opcode == 11) && op1.GetType() == 6 && op2.GetType() == 6 {
		return 0
	}
	if op1.GetType() == 6 && IsNumericString(op1.GetValue().GetStr().GetVal(), op1.GetValue().GetStr().GetLen(), nil, nil, 0) == 0 {
		return 1
	}
	if op2.GetType() == 6 && IsNumericString(op2.GetValue().GetStr().GetVal(), op2.GetValue().GetStr().GetLen(), nil, nil, 0) == 0 {
		return 1
	}
	return 0
}

/* }}} */

func ZendBinaryOpProducesArrayConversionError(opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	if opcode == 8 && (op1.GetType() == 7 || op2.GetType() == 7) {
		return 1
	}
	return 0
}

/* }}} */

func ZendTryCtEvalBinaryOp(result *Zval, opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	var fn BinaryOpType = GetBinaryOp(opcode)

	/* don't evaluate division by zero at compile-time */

	if (opcode == 4 || opcode == 5) && ZvalGetLong(op2) == 0 {
		return 0
	} else if (opcode == 6 || opcode == 7) && ZvalGetLong(op2) < 0 {
		return 0
	}

	/* don't evaluate numeric string error-producing operations at compile-time */

	if ZendBinaryOpProducesNumericStringError(opcode, op1, op2) != 0 {
		return 0
	}

	/* don't evaluate array to string conversions at compile-time */

	if ZendBinaryOpProducesArrayConversionError(opcode, op1, op2) != 0 {
		return 0
	}
	fn(result, op1, op2)
	return 1
}

/* }}} */

func ZendCtEvalUnaryOp(result *Zval, opcode uint32, op *Zval) {
	var fn UnaryOpType = GetUnaryOp(opcode)
	fn(result, op)
}

/* }}} */

func ZendTryCtEvalUnaryPm(result *Zval, kind ZendAstKind, op *Zval) ZendBool {
	var left Zval
	var __z *Zval = &left
	if kind == ZEND_AST_UNARY_PLUS {
		__z.GetValue().SetLval(1)
	} else {
		__z.GetValue().SetLval(-1)
	}
	__z.SetTypeInfo(4)
	return ZendTryCtEvalBinaryOp(result, 3, &left, op)
}

/* }}} */

func ZendCtEvalGreater(result *Zval, kind ZendAstKind, op1 *Zval, op2 *Zval) {
	var fn BinaryOpType = g.Cond(kind == ZEND_AST_GREATER, IsSmallerFunction, IsSmallerOrEqualFunction)
	fn(result, op2, op1)
}

/* }}} */

func ZendTryCtEvalArray(result *Zval, ast *ZendAst) ZendBool {
	var list *ZendAstList = ZendAstGetList(ast)
	var last_elem_ast *ZendAst = nil
	var i uint32
	var is_constant ZendBool = 1
	if ast.GetAttr() == 1 {
		ZendError(1<<6, "Cannot use list() as standalone expression")
	}

	/* First ensure that *all* child nodes are constant and by-val */

	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		if elem_ast == nil {

			/* Report error at line of last non-empty element */

			if last_elem_ast != nil {
				CG.SetZendLineno(ZendAstGetLineno(last_elem_ast))
			}
			ZendError(1<<6, "Cannot use empty array elements in arrays")
		}
		if elem_ast.GetKind() != ZEND_AST_UNPACK {
			ZendEvalConstExpr(&elem_ast.child[0])
			ZendEvalConstExpr(&elem_ast.child[1])
			if elem_ast.GetAttr() != 0 || elem_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || elem_ast.GetChild()[1] != nil && elem_ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
				is_constant = 0
			}
		} else {
			ZendEvalConstExpr(&elem_ast.child[0])
			if elem_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
				is_constant = 0
			}
		}
		last_elem_ast = elem_ast
	}
	if is_constant == 0 {
		return 0
	}
	if list.GetChildren() == 0 {
		var __z *Zval = result
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
		return 1
	}
	var __arr *ZendArray = _zendNewArray(list.GetChildren())
	var __z *Zval = result
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var value_ast *ZendAst = elem_ast.GetChild()[0]
		var key_ast *ZendAst
		var value *Zval = ZendAstGetZval(value_ast)
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			if value.GetType() == 7 {
				var ht *HashTable = value.GetValue().GetArr()
				var val *Zval
				var key *ZendString
				for {
					var __ht *HashTable = ht
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						key = _p.GetKey()
						val = _z
						if key != nil {
							ZendErrorNoreturn(1<<6, "Cannot unpack array with string keys")
						}
						if ZendHashNextIndexInsert(result.GetValue().GetArr(), val) == nil {
							ZvalPtrDtor(result)
							return 0
						}
						if val.GetTypeFlags() != 0 {
							ZvalAddrefP(val)
						}
					}
					break
				}
				continue
			} else {
				ZendErrorNoreturn(1<<6, "Only arrays and Traversables can be unpacked")
			}
		}
		if value.GetTypeFlags() != 0 {
			ZvalAddrefP(value)
		}
		key_ast = elem_ast.GetChild()[1]
		if key_ast != nil {
			var key *Zval = ZendAstGetZval(key_ast)
			switch key.GetType() {
			case 4:
				ZendHashIndexUpdate(result.GetValue().GetArr(), key.GetValue().GetLval(), value)
				break
			case 6:
				ZendSymtableUpdate(result.GetValue().GetArr(), key.GetValue().GetStr(), value)
				break
			case 5:
				ZendHashIndexUpdate(result.GetValue().GetArr(), ZendDvalToLval(key.GetValue().GetDval()), value)
				break
			case 2:
				ZendHashIndexUpdate(result.GetValue().GetArr(), 0, value)
				break
			case 3:
				ZendHashIndexUpdate(result.GetValue().GetArr(), 1, value)
				break
			case 1:
				ZendHashUpdate(result.GetValue().GetArr(), ZendEmptyString, value)
				break
			default:
				ZendErrorNoreturn(1<<6, "Illegal offset type")
				break
			}
		} else {
			if ZendHashNextIndexInsert(result.GetValue().GetArr(), value) == nil {
				ZvalPtrDtorNogc(value)
				ZvalPtrDtor(result)
				return 0
			}
		}
	}
	return 1
}

/* }}} */

func ZendCompileBinaryOp(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.GetChild()[0]
	var right_ast *ZendAst = ast.GetChild()[1]
	var opcode uint32 = ast.GetAttr()
	if (opcode == 1 || opcode == 2) && left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == 8 {
		ZendError(1<<13, "The behavior of unparenthesized expressions containing both '.' and '+'/'-' will change in PHP 8: '+'/'-' will take a higher precedence")
	}
	if (opcode == 6 || opcode == 7) && (left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == 8 || right_ast.GetKind() == ZEND_AST_BINARY_OP && right_ast.GetAttr() == 8) {
		ZendError(1<<13, "The behavior of unparenthesized expressions containing both '.' and '>>'/'<<' will change in PHP 8: '<<'/'>>' will take a higher precedence")
	}
	if opcode == 252 {
		opcode = 8
	}
	var left_node Znode
	var right_node Znode
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == 1<<0 && right_node.GetOpType() == 1<<0 {
		if ZendTryCtEvalBinaryOp(&result.u.constant, opcode, &left_node.u.constant, &right_node.u.constant) != 0 {
			result.SetOpType(1 << 0)
			ZvalPtrDtor(&left_node.u.constant)
			ZvalPtrDtor(&right_node.u.constant)
			return
		}
	}
	for {
		if opcode == 18 || opcode == 19 {
			if left_node.GetOpType() == 1<<0 {
				if left_node.GetConstant().GetType() == 2 {
					if opcode == 19 {
						opcode = 52
					} else {
						opcode = 14
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				} else if left_node.GetConstant().GetType() == 3 {
					if opcode == 18 {
						opcode = 52
					} else {
						opcode = 14
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				}
			} else if right_node.GetOpType() == 1<<0 {
				if right_node.GetConstant().GetType() == 2 {
					if opcode == 19 {
						opcode = 52
					} else {
						opcode = 14
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				} else if right_node.GetConstant().GetType() == 3 {
					if opcode == 18 {
						opcode = 52
					} else {
						opcode = 14
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				}
			}
		}
		if opcode == 8 {

			/* convert constant operands to strings at compile-time */

			if left_node.GetOpType() == 1<<0 {
				if left_node.GetConstant().GetType() == 7 {
					ZendEmitOpTmp(&left_node, 51, &left_node, nil).SetExtendedValue(6)
				} else {
					if &left_node.u.constant.u1.v.type_ != 6 {
						_convertToString(&left_node.u.constant)
					}
				}
			}
			if right_node.GetOpType() == 1<<0 {
				if right_node.GetConstant().GetType() == 7 {
					ZendEmitOpTmp(&right_node, 51, &right_node, nil).SetExtendedValue(6)
				} else {
					if &right_node.u.constant.u1.v.type_ != 6 {
						_convertToString(&right_node.u.constant)
					}
				}
			}
			if left_node.GetOpType() == 1<<0 && right_node.GetOpType() == 1<<0 {
				opcode = 53
			}
		}
		ZendEmitOpTmp(result, opcode, &left_node, &right_node)
		break
	}
}

/* }}} */

func ZendCompileGreater(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.GetChild()[0]
	var right_ast *ZendAst = ast.GetChild()[1]
	var left_node Znode
	var right_node Znode
	assert(ast.GetKind() == ZEND_AST_GREATER || ast.GetKind() == ZEND_AST_GREATER_EQUAL)
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == 1<<0 && right_node.GetOpType() == 1<<0 {
		result.SetOpType(1 << 0)
		ZendCtEvalGreater(&result.u.constant, ast.GetKind(), &left_node.u.constant, &right_node.u.constant)
		ZvalPtrDtor(&left_node.u.constant)
		ZvalPtrDtor(&right_node.u.constant)
		return
	}
	ZendEmitOpTmp(result, g.Cond(ast.GetKind() == ZEND_AST_GREATER, 20, 21), &right_node, &left_node)
}

/* }}} */

func ZendCompileUnaryOp(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var opcode uint32 = ast.GetAttr()
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == 1<<0 {
		result.SetOpType(1 << 0)
		ZendCtEvalUnaryOp(&result.u.constant, opcode, &expr_node.u.constant)
		ZvalPtrDtor(&expr_node.u.constant)
		return
	}
	ZendEmitOpTmp(result, opcode, &expr_node, nil)
}

/* }}} */

func ZendCompileUnaryPm(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	var lefthand_node Znode
	assert(ast.GetKind() == ZEND_AST_UNARY_PLUS || ast.GetKind() == ZEND_AST_UNARY_MINUS)
	ZendCompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == 1<<0 {
		if ZendTryCtEvalUnaryPm(&result.u.constant, ast.GetKind(), &expr_node.u.constant) != 0 {
			result.SetOpType(1 << 0)
			ZvalPtrDtor(&expr_node.u.constant)
			return
		}
	}
	lefthand_node.SetOpType(1 << 0)
	var __z *Zval = &lefthand_node.u.constant
	if ast.GetKind() == ZEND_AST_UNARY_PLUS {
		__z.GetValue().SetLval(1)
	} else {
		__z.GetValue().SetLval(-1)
	}
	__z.SetTypeInfo(4)
	ZendEmitOpTmp(result, 3, &lefthand_node, &expr_node)
}

/* }}} */

func ZendCompileShortCircuiting(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.GetChild()[0]
	var right_ast *ZendAst = ast.GetChild()[1]
	var left_node Znode
	var right_node Znode
	var opline_jmpz *ZendOp
	var opline_bool *ZendOp
	var opnum_jmpz uint32
	assert(ast.GetKind() == ZEND_AST_AND || ast.GetKind() == ZEND_AST_OR)
	ZendCompileExpr(&left_node, left_ast)
	if left_node.GetOpType() == 1<<0 {
		if ast.GetKind() == ZEND_AST_AND && ZendIsTrue(&left_node.u.constant) == 0 || ast.GetKind() == ZEND_AST_OR && ZendIsTrue(&left_node.u.constant) != 0 {
			result.SetOpType(1 << 0)
			if ZendIsTrue(&left_node.u.constant) != 0 {
				&result.u.constant.u1.type_info = 3
			} else {
				&result.u.constant.u1.type_info = 2
			}
		} else {
			ZendCompileExpr(&right_node, right_ast)
			if right_node.GetOpType() == 1<<0 {
				result.SetOpType(1 << 0)
				if ZendIsTrue(&right_node.u.constant) != 0 {
					&result.u.constant.u1.type_info = 3
				} else {
					&result.u.constant.u1.type_info = 2
				}
				ZvalPtrDtor(&right_node.u.constant)
			} else {
				ZendEmitOpTmp(result, 52, &right_node, nil)
			}
		}
		ZvalPtrDtor(&left_node.u.constant)
		return
	}
	opnum_jmpz = GetNextOpNumber()
	opline_jmpz = ZendEmitOp(nil, g.Cond(ast.GetKind() == ZEND_AST_AND, 46, 47), &left_node, nil)
	if left_node.GetOpType() == 1<<1 {
		opline_jmpz.SetResultType(&left_node.GetOpType())
		if &left_node.GetOpType() == 1<<0 {
			opline_jmpz.GetResult().SetConstant(ZendAddLiteral(&(&left_node).u.constant))
		} else {
			opline_jmpz.SetResult(&left_node.GetOp())
		}
	} else {
		opline_jmpz.GetResult().SetVar(GetTemporaryVariable())
		opline_jmpz.SetResultType(1 << 1)
	}
	result.SetOpType(opline_jmpz.GetResultType())
	if result.GetOpType() == 1<<0 {
		var _z1 *Zval = &result.u.constant
		var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline_jmpz.GetResult().GetConstant()
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		result.SetOp(opline_jmpz.GetResult())
	}
	ZendCompileExpr(&right_node, right_ast)
	opline_bool = ZendEmitOp(nil, 52, &right_node, nil)
	opline_bool.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline_bool.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline_bool.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmpz)
}

/* }}} */

func ZendCompilePostIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	assert(ast.GetKind() == ZEND_AST_POST_INC || ast.GetKind() == ZEND_AST_POST_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.GetKind() == ZEND_AST_PROP {
		var opline *ZendOp = ZendCompileProp(nil, var_ast, 2, 0)
		if ast.GetKind() == ZEND_AST_POST_INC {
			opline.SetOpcode(134)
		} else {
			opline.SetOpcode(135)
		}
		ZendMakeTmpResult(result, opline)
	} else if var_ast.GetKind() == ZEND_AST_STATIC_PROP {
		var opline *ZendOp = ZendCompileStaticProp(nil, var_ast, 2, 0, 0)
		if ast.GetKind() == ZEND_AST_POST_INC {
			opline.SetOpcode(40)
		} else {
			opline.SetOpcode(41)
		}
		ZendMakeTmpResult(result, opline)
	} else {
		var var_node Znode
		ZendCompileVar(&var_node, var_ast, 2, 0)
		ZendEmitOpTmp(result, g.Cond(ast.GetKind() == ZEND_AST_POST_INC, 36, 37), &var_node, nil)
	}
}

/* }}} */

func ZendCompilePreIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	assert(ast.GetKind() == ZEND_AST_PRE_INC || ast.GetKind() == ZEND_AST_PRE_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.GetKind() == ZEND_AST_PROP {
		var opline *ZendOp = ZendCompileProp(result, var_ast, 2, 0)
		if ast.GetKind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(132)
		} else {
			opline.SetOpcode(133)
		}
	} else if var_ast.GetKind() == ZEND_AST_STATIC_PROP {
		var opline *ZendOp = ZendCompileStaticProp(result, var_ast, 2, 0, 0)
		if ast.GetKind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(38)
		} else {
			opline.SetOpcode(39)
		}
	} else {
		var var_node Znode
		ZendCompileVar(&var_node, var_ast, 2, 0)
		ZendEmitOp(result, g.Cond(ast.GetKind() == ZEND_AST_PRE_INC, 34, 35), &var_node, nil)
	}
}

/* }}} */

func ZendCompileCast(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	var opline *ZendOp
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOpTmp(result, 51, &expr_node, nil)
	opline.SetExtendedValue(ast.GetAttr())
	if ast.GetAttr() == 1 {
		ZendError(1<<13, "The (unset) cast is deprecated")
	}
}

/* }}} */

func ZendCompileShorthandConditional(result *Znode, ast *ZendAst) {
	var cond_ast *ZendAst = ast.GetChild()[0]
	var false_ast *ZendAst = ast.GetChild()[2]
	var cond_node Znode
	var false_node Znode
	var opline_qm_assign *ZendOp
	var opnum_jmp_set uint32
	assert(ast.GetChild()[1] == nil)
	ZendCompileExpr(&cond_node, cond_ast)
	opnum_jmp_set = GetNextOpNumber()
	ZendEmitOpTmp(result, 152, &cond_node, nil)
	ZendCompileExpr(&false_node, false_ast)
	opline_qm_assign = ZendEmitOpTmp(nil, 31, &false_node, nil)
	opline_qm_assign.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline_qm_assign.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline_qm_assign.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmp_set)
}

/* }}} */

func ZendCompileConditional(result *Znode, ast *ZendAst) {
	var cond_ast *ZendAst = ast.GetChild()[0]
	var true_ast *ZendAst = ast.GetChild()[1]
	var false_ast *ZendAst = ast.GetChild()[2]
	var cond_node Znode
	var true_node Znode
	var false_node Znode
	var opline_qm_assign2 *ZendOp
	var opnum_jmpz uint32
	var opnum_jmp uint32
	if cond_ast.GetKind() == ZEND_AST_CONDITIONAL && cond_ast.GetAttr() != 1 {
		if cond_ast.GetChild()[1] != nil {
			if true_ast != nil {
				ZendError(1<<13, "Unparenthesized `a ? b : c ? d : e` is deprecated. "+"Use either `(a ? b : c) ? d : e` or `a ? b : (c ? d : e)`")
			} else {
				ZendError(1<<13, "Unparenthesized `a ? b : c ?: d` is deprecated. "+"Use either `(a ? b : c) ?: d` or `a ? b : (c ?: d)`")
			}
		} else {
			if true_ast != nil {
				ZendError(1<<13, "Unparenthesized `a ?: b ? c : d` is deprecated. "+"Use either `(a ?: b) ? c : d` or `a ?: (b ? c : d)`")
			}
		}
	}
	if true_ast == nil {
		ZendCompileShorthandConditional(result, ast)
		return
	}
	ZendCompileExpr(&cond_node, cond_ast)
	opnum_jmpz = ZendEmitCondJump(43, &cond_node, 0)
	ZendCompileExpr(&true_node, true_ast)
	ZendEmitOpTmp(result, 31, &true_node, nil)
	opnum_jmp = ZendEmitJump(0)
	ZendUpdateJumpTargetToNext(opnum_jmpz)
	ZendCompileExpr(&false_node, false_ast)
	opline_qm_assign2 = ZendEmitOp(nil, 31, &false_node, nil)
	opline_qm_assign2.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline_qm_assign2.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline_qm_assign2.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmp)
}

/* }}} */

func ZendCompileCoalesce(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var default_ast *ZendAst = ast.GetChild()[1]
	var expr_node Znode
	var default_node Znode
	var opline *ZendOp
	var opnum uint32
	ZendCompileVar(&expr_node, expr_ast, 3, 0)
	opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, 169, &expr_node, nil)
	ZendCompileExpr(&default_node, default_ast)
	opline = ZendEmitOpTmp(nil, 31, &default_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline = &CG.active_op_array.GetOpcodes()[opnum]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
}

/* }}} */

func ZnodeDtor(zv *Zval) {
	var node *Znode = zv.GetValue().GetPtr()
	if node.GetOpType() == 1<<0 {
		ZvalPtrDtorNogc(&node.u.constant)
	}
	_efree(node)
}
func ZendCompileAssignCoalesce(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var default_ast *ZendAst = ast.GetChild()[1]
	var var_node_is Znode
	var var_node_w Znode
	var default_node Znode
	var assign_node Znode
	var node *Znode
	var opline *ZendOp
	var coalesce_opnum uint32
	var need_frees ZendBool = 0

	/* Remember expressions compiled during the initial BP_VAR_IS lookup,
	 * to avoid double-evaluation when we compile again with BP_VAR_W. */

	var orig_memoized_exprs *HashTable = CG.GetMemoizedExprs()
	var orig_memoize_mode int = CG.GetMemoizeMode()
	ZendEnsureWritableVariable(var_ast)
	if IsThisFetch(var_ast) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot re-assign $this")
	}
	CG.SetMemoizedExprs((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
	_zendHashInit(CG.GetMemoizedExprs(), 0, ZnodeDtor, 0)
	CG.SetMemoizeMode(1)
	ZendCompileVar(&var_node_is, var_ast, 3, 0)
	coalesce_opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, 169, &var_node_is, nil)
	CG.SetMemoizeMode(0)
	ZendCompileExpr(&default_node, default_ast)
	CG.SetMemoizeMode(2)
	ZendCompileVar(&var_node_w, var_ast, 1, 0)

	/* Reproduce some of the zend_compile_assign() opcode fixup logic here. */

	opline = &CG.active_op_array.GetOpcodes()[CG.GetActiveOpArray().GetLast()-1]
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		ZendEmitOp(&assign_node, 22, &var_node_w, &default_node)
		break
	case ZEND_AST_STATIC_PROP:
		opline.SetOpcode(25)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	case ZEND_AST_DIM:
		opline.SetOpcode(23)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	case ZEND_AST_PROP:
		opline.SetOpcode(24)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	default:
		break
	}
	opline = ZendEmitOpTmp(nil, 31, &assign_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	for {
		var __ht *HashTable = CG.GetMemoizedExprs()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			node = _z.GetValue().GetPtr()
			if node.GetOpType() == 1<<1 || node.GetOpType() == 1<<2 {
				need_frees = 1
				break
			}
		}
		break
	}

	/* Free DUPed expressions if there are any */

	if need_frees != 0 {
		var jump_opnum uint32 = ZendEmitJump(0)
		ZendUpdateJumpTargetToNext(coalesce_opnum)
		for {
			var __ht *HashTable = CG.GetMemoizedExprs()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				node = _z.GetValue().GetPtr()
				if node.GetOpType() == 1<<1 || node.GetOpType() == 1<<2 {
					ZendEmitOp(nil, 70, node, nil)
				}
			}
			break
		}
		ZendUpdateJumpTargetToNext(jump_opnum)
	} else {
		ZendUpdateJumpTargetToNext(coalesce_opnum)
	}
	ZendHashDestroy(CG.GetMemoizedExprs())
	_efree(CG.GetMemoizedExprs())
	CG.SetMemoizedExprs(orig_memoized_exprs)
	CG.SetMemoizeMode(orig_memoize_mode)
}

/* }}} */

func ZendCompilePrint(result *Znode, ast *ZendAst) {
	var opline *ZendOp
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, 136, &expr_node, nil)
	opline.SetExtendedValue(1)
	result.SetOpType(1 << 0)
	var __z *Zval = &result.u.constant
	__z.GetValue().SetLval(1)
	__z.SetTypeInfo(4)
}

/* }}} */

func ZendCompileExit(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	if expr_ast != nil {
		var expr_node Znode
		ZendCompileExpr(&expr_node, expr_ast)
		ZendEmitOp(nil, 79, &expr_node, nil)
	} else {
		ZendEmitOp(nil, 79, nil, nil)
	}
	result.SetOpType(1 << 0)
	&result.u.constant.u1.type_info = 3
}

/* }}} */

func ZendCompileYield(result *Znode, ast *ZendAst) {
	var value_ast *ZendAst = ast.GetChild()[0]
	var key_ast *ZendAst = ast.GetChild()[1]
	var value_node Znode
	var key_node Znode
	var value_node_ptr *Znode = nil
	var key_node_ptr *Znode = nil
	var opline *ZendOp
	var returns_by_ref ZendBool = (CG.GetActiveOpArray().GetFnFlags() & 1 << 12) != 0
	ZendMarkFunctionAsGenerator()
	if key_ast != nil {
		ZendCompileExpr(&key_node, key_ast)
		key_node_ptr = &key_node
	}
	if value_ast != nil {
		if returns_by_ref != 0 && ZendIsVariable(value_ast) != 0 {
			ZendCompileVar(&value_node, value_ast, 1, 1)
		} else {
			ZendCompileExpr(&value_node, value_ast)
		}
		value_node_ptr = &value_node
	}
	opline = ZendEmitOp(result, 160, value_node_ptr, key_node_ptr)
	if value_ast != nil && returns_by_ref != 0 && ZendIsCall(value_ast) != 0 {
		opline.SetExtendedValue(1 << 0)
	}
}

/* }}} */

func ZendCompileYieldFrom(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendMarkFunctionAsGenerator()
	if (CG.GetActiveOpArray().GetFnFlags() & 1 << 12) != 0 {
		ZendErrorNoreturn(1<<6, "Cannot use \"yield from\" inside a by-reference generator")
	}
	ZendCompileExpr(&expr_node, expr_ast)
	ZendEmitOpTmp(result, 166, &expr_node, nil)
}

/* }}} */

func ZendCompileInstanceof(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var class_ast *ZendAst = ast.GetChild()[1]
	var obj_node Znode
	var class_node Znode
	var opline *ZendOp
	ZendCompileExpr(&obj_node, obj_ast)
	if obj_node.GetOpType() == 1<<0 {
		ZendDoFree(&obj_node)
		result.SetOpType(1 << 0)
		&result.u.constant.u1.type_info = 2
		return
	}
	ZendCompileClassRef(&class_node, class_ast, 0x80|0x200)
	opline = ZendEmitOpTmp(result, 138, &obj_node, nil)
	if class_node.GetOpType() == 1<<0 {
		opline.SetOp2Type(1 << 0)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().GetValue().GetStr()))
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {
		opline.SetOp2Type(&class_node.GetOpType())
		if &class_node.GetOpType() == 1<<0 {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&class_node).u.constant))
		} else {
			opline.SetOp2(&class_node.GetOp())
		}
	}
}

/* }}} */

func ZendCompileIncludeOrEval(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	var opline *ZendOp
	ZendDoExtendedFcallBegin()
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(result, 73, &expr_node, nil)
	opline.SetExtendedValue(ast.GetAttr())
	ZendDoExtendedFcallEnd()
}

/* }}} */

func ZendCompileIssetOrEmpty(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var var_node Znode
	var opline *ZendOp = nil
	assert(ast.GetKind() == ZEND_AST_ISSET || ast.GetKind() == ZEND_AST_EMPTY)
	if ZendIsVariable(var_ast) == 0 {
		if ast.GetKind() == ZEND_AST_EMPTY {

			/* empty(expr) can be transformed to !expr */

			var not_ast *ZendAst = ZendAstCreateEx1(ZEND_AST_UNARY_OP, 14, var_ast)
			ZendCompileExpr(result, not_ast)
			return
		} else {
			ZendErrorNoreturn(1<<6, "Cannot use isset() on the result of an expression "+"(you can use \"null !== expression\" instead)")
		}
	}
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		if IsThisFetch(var_ast) != 0 {
			opline = ZendEmitOp(result, 186, nil, nil)
			CG.GetActiveOpArray().SetFnFlags(CG.GetActiveOpArray().GetFnFlags() | 1<<30)
		} else if ZendTryCompileCv(&var_node, var_ast) == SUCCESS {
			opline = ZendEmitOp(result, 154, &var_node, nil)
		} else {
			opline = ZendCompileSimpleVarNoCv(result, var_ast, 3, 0)
			opline.SetOpcode(114)
		}
		break
	case ZEND_AST_DIM:
		opline = ZendCompileDim(result, var_ast, 3)
		opline.SetOpcode(115)
		break
	case ZEND_AST_PROP:
		opline = ZendCompileProp(result, var_ast, 3, 0)
		opline.SetOpcode(148)
		break
	case ZEND_AST_STATIC_PROP:
		opline = ZendCompileStaticProp(result, var_ast, 3, 0, 0)
		opline.SetOpcode(180)
		break
	default:
		break
	}
	opline.SetResultType(1 << 1)
	result.SetOpType(opline.GetResultType())
	if ast.GetKind() != ZEND_AST_ISSET {
		opline.SetExtendedValue(opline.GetExtendedValue() | 1<<0)
	}
}

/* }}} */

func ZendCompileSilence(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var silence_node Znode
	ZendEmitOpTmp(&silence_node, 57, nil, nil)
	if expr_ast.GetKind() == ZEND_AST_VAR {

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

		ZendCompileSimpleVarNoCv(result, expr_ast, 0, 0)

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

	} else {
		ZendCompileExpr(result, expr_ast)
	}
	ZendEmitOp(nil, 58, &silence_node, nil)
}

/* }}} */

func ZendCompileShellExec(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var fn_name Zval
	var name_ast *ZendAst
	var args_ast *ZendAst
	var call_ast *ZendAst
	var _s *byte = "shell_exec"
	var __z *Zval = &fn_name
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	name_ast = ZendAstCreateZval(&fn_name)
	args_ast = ZendAstCreateList1(ZEND_AST_ARG_LIST, expr_ast)
	call_ast = ZendAstCreate2(ZEND_AST_CALL, name_ast, args_ast)
	ZendCompileExpr(result, call_ast)
	ZvalPtrDtor(&fn_name)
}

/* }}} */

func ZendCompileArray(result *Znode, ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var opline *ZendOp
	var i uint32
	var opnum_init uint32 = -1
	var packed ZendBool = 1
	if ZendTryCtEvalArray(&result.u.constant, ast) != 0 {
		result.SetOpType(1 << 0)
		return
	}

	/* Empty arrays are handled at compile-time */

	assert(list.GetChildren() > 0)
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var value_ast *ZendAst
		var key_ast *ZendAst
		var by_ref ZendBool
		var value_node Znode
		var key_node Znode
		var key_node_ptr *Znode = nil
		if elem_ast == nil {
			ZendError(1<<6, "Cannot use empty array elements in arrays")
		}
		value_ast = elem_ast.GetChild()[0]
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			ZendCompileExpr(&value_node, value_ast)
			if i == 0 {
				opnum_init = GetNextOpNumber()
				opline = ZendEmitOpTmp(result, 71, nil, nil)
			}
			opline = ZendEmitOp(nil, 147, &value_node, nil)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == 1<<0 {
				opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
			} else {
				opline.SetResult(result.GetOp())
			}
			continue
		}
		key_ast = elem_ast.GetChild()[1]
		by_ref = elem_ast.GetAttr()
		if key_ast != nil {
			ZendCompileExpr(&key_node, key_ast)
			ZendHandleNumericOp(&key_node)
			key_node_ptr = &key_node
		}
		if by_ref != 0 {
			ZendEnsureWritableVariable(value_ast)
			ZendCompileVar(&value_node, value_ast, 1, 1)
		} else {
			ZendCompileExpr(&value_node, value_ast)
		}
		if i == 0 {
			opnum_init = GetNextOpNumber()
			opline = ZendEmitOpTmp(result, 71, &value_node, key_node_ptr)
			opline.SetExtendedValue(list.GetChildren() << 2)
		} else {
			opline = ZendEmitOp(nil, 72, &value_node, key_node_ptr)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == 1<<0 {
				opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
			} else {
				opline.SetResult(result.GetOp())
			}
		}
		opline.SetExtendedValue(opline.GetExtendedValue() | by_ref)
		if key_ast != nil && key_node.GetOpType() == 1<<0 && key_node.GetConstant().GetType() == 6 {
			packed = 0
		}
	}

	/* Add a flag to INIT_ARRAY if we know this array cannot be packed */

	if packed == 0 {
		assert(opnum_init != uint32-1)
		opline = &CG.active_op_array.GetOpcodes()[opnum_init]
		opline.SetExtendedValue(opline.GetExtendedValue() | 1<<1)
	}

	/* Add a flag to INIT_ARRAY if we know this array cannot be packed */
}

/* }}} */

func ZendCompileConst(result *Znode, ast *ZendAst) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	var is_fully_qualified ZendBool
	var orig_name *ZendString = ZendAstGetStr(name_ast)
	var resolved_name *ZendString = ZendResolveConstName(orig_name, name_ast.GetAttr(), &is_fully_qualified)
	if resolved_name.GetLen() == g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(resolved_name.GetVal(), "__COMPILER_HALT_OFFSET__", g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) || name_ast.GetAttr() != 2 && (orig_name.GetLen() == g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(orig_name.GetVal(), "__COMPILER_HALT_OFFSET__", g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1))) {
		var last *ZendAst = CG.GetAst()
		for last != nil && last.GetKind() == ZEND_AST_STMT_LIST {
			var list *ZendAstList = ZendAstGetList(last)
			if list.GetChildren() == 0 {
				break
			}
			last = list.GetChild()[list.GetChildren()-1]
		}
		if last != nil && last.GetKind() == ZEND_AST_HALT_COMPILER {
			result.SetOpType(1 << 0)
			var __z *Zval = &result.u.constant
			__z.GetValue().SetLval(ZendAstGetZval(last.GetChild()[0]).GetValue().GetLval())
			__z.SetTypeInfo(4)
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
	}
	if ZendTryCtEvalConst(&result.u.constant, resolved_name, is_fully_qualified) != 0 {
		result.SetOpType(1 << 0)
		ZendStringReleaseEx(resolved_name, 0)
		return
	}
	opline = ZendEmitOpTmp(result, 99, nil, nil)
	opline.SetOp2Type(1 << 0)
	if is_fully_qualified != 0 {
		opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name, 0))
	} else {
		opline.GetOp1().SetNum(0x10)
		if CG.GetFileContext().GetCurrentNamespace() != nil {
			opline.GetOp1().SetNum(opline.GetOp1().GetNum() | 0x100)
			opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name, 1))
		} else {
			opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name, 0))
		}
	}
	opline.SetExtendedValue(ZendAllocCacheSlot())
}

/* }}} */

func ZendCompileClassConst(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var const_ast *ZendAst = ast.GetChild()[1]
	var class_node Znode
	var const_node Znode
	var opline *ZendOp
	ZendEvalConstExpr(&ast.child[0])
	ZendEvalConstExpr(&ast.child[1])
	class_ast = ast.GetChild()[0]
	const_ast = ast.GetChild()[1]
	if class_ast.GetKind() == ZEND_AST_ZVAL {
		var resolved_name *ZendString
		resolved_name = ZendResolveClassNameAst(class_ast)
		if const_ast.GetKind() == ZEND_AST_ZVAL && ZendTryCtEvalClassConst(&result.u.constant, resolved_name, ZendAstGetStr(const_ast)) != 0 {
			result.SetOpType(1 << 0)
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
		ZendStringReleaseEx(resolved_name, 0)
	}
	ZendCompileClassRef(&class_node, class_ast, 0x200)
	ZendCompileExpr(&const_node, const_ast)
	opline = ZendEmitOpTmp(result, 181, nil, &const_node)
	ZendSetClassNameOp1(opline, &class_node)
	opline.SetExtendedValue(ZendAllocCacheSlots(2))
}

/* }}} */

func ZendCompileClassName(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	if ZendTryCompileConstExprResolveClassName(&result.u.constant, class_ast) != 0 {
		result.SetOpType(1 << 0)
		return
	}
	opline = ZendEmitOpTmp(result, 157, nil, nil)
	opline.GetOp1().SetNum(ZendGetClassFetchType(ZendAstGetStr(class_ast)))
}

/* }}} */

func ZendCompileRopeAddEx(opline *ZendOp, result *Znode, num uint32, elem_node *Znode) *ZendOp {
	if num == 0 {
		result.SetOpType(1 << 1)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(54)
	} else {
		opline.SetOpcode(55)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == 1<<0 {
		opline.GetOp2().SetConstant(ZendAddLiteral(&elem_node.u.constant))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline.SetExtendedValue(num)
	return opline
}

/* }}} */

func ZendCompileRopeAdd(result *Znode, num uint32, elem_node *Znode) *ZendOp {
	var opline *ZendOp = GetNextOp()
	if num == 0 {
		result.SetOpType(1 << 1)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(54)
	} else {
		opline.SetOpcode(55)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == 1<<0 {
			opline.GetOp1().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == 1<<0 {
		opline.GetOp2().SetConstant(ZendAddLiteral(&elem_node.u.constant))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == 1<<0 {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline.SetExtendedValue(num)
	return opline
}

/* }}} */

func ZendCompileEncapsList(result *Znode, ast *ZendAst) {
	var i uint32
	var j uint32
	var rope_init_lineno uint32 = -1
	var opline *ZendOp = nil
	var init_opline *ZendOp
	var elem_node Znode
	var last_const_node Znode
	var list *ZendAstList = ZendAstGetList(ast)
	var reserved_op_number uint32 = -1
	assert(list.GetChildren() > 0)
	j = 0
	last_const_node.SetOpType(0)
	for i = 0; i < list.GetChildren(); i++ {
		ZendCompileExpr(&elem_node, list.GetChild()[i])
		if elem_node.GetOpType() == 1<<0 {
			if &elem_node.u.constant.u1.v.type_ != 6 {
				_convertToString(&elem_node.u.constant)
			}
			if elem_node.GetConstant().GetValue().GetStr().GetLen() == 0 {
				ZvalPtrDtor(&elem_node.u.constant)
			} else if last_const_node.GetOpType() == 1<<0 {
				ConcatFunction(&last_const_node.u.constant, &last_const_node.u.constant, &elem_node.u.constant)
				ZvalPtrDtor(&elem_node.u.constant)
			} else {
				last_const_node.SetOpType(1 << 0)
				var _z1 *Zval = &last_const_node.u.constant
				var _z2 *Zval = &elem_node.u.constant
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)

				/* Reserve place for ZEND_ROPE_ADD instruction */

				reserved_op_number = GetNextOpNumber()
				opline = GetNextOp()
				opline.SetOpcode(0)
			}
			continue
		} else {
			if j == 0 {
				if last_const_node.GetOpType() == 1<<0 {
					rope_init_lineno = reserved_op_number
				} else {
					rope_init_lineno = GetNextOpNumber()
				}
			}
			if last_const_node.GetOpType() == 1<<0 {
				opline = &CG.active_op_array.GetOpcodes()[reserved_op_number]
				ZendCompileRopeAddEx(opline, result, g.PostInc(&j), &last_const_node)
				last_const_node.SetOpType(0)
			}
			opline = ZendCompileRopeAdd(result, g.PostInc(&j), &elem_node)
		}
	}
	if j == 0 {
		result.SetOpType(1 << 0)
		if last_const_node.GetOpType() == 1<<0 {
			var _z1 *Zval = &result.u.constant
			var _z2 *Zval = &last_const_node.u.constant
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			var __z *Zval = &result.u.constant
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		CG.GetActiveOpArray().SetLast(reserved_op_number - 1)
		return
	} else if last_const_node.GetOpType() == 1<<0 {
		opline = &CG.active_op_array.GetOpcodes()[reserved_op_number]
		opline = ZendCompileRopeAddEx(opline, result, g.PostInc(&j), &last_const_node)
	}
	init_opline = CG.GetActiveOpArray().GetOpcodes() + rope_init_lineno
	if j == 1 {
		if opline.GetOp2Type() == 1<<0 {
			result.SetOpType(opline.GetOp2Type())
			if result.GetOpType() == 1<<0 {
				var _z1 *Zval = &result.u.constant
				var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				result.SetOp(opline.GetOp2())
			}
			opline.GetOp1().SetNum(0)
			opline.GetOp2().SetNum(0)
			opline.GetResult().SetNum(0)
			opline.SetOpcode(0)
			opline.SetOp1Type(0)
			opline.SetOp2Type(0)
			opline.SetResultType(0)
		} else {
			opline.SetOpcode(51)
			opline.SetExtendedValue(6)
			opline.SetOp1Type(opline.GetOp2Type())
			opline.SetOp1(opline.GetOp2())
			opline.SetResultType(1 << 1)
			opline.GetResult().SetVar(GetTemporaryVariable())
			opline.SetOp2Type(0)
			result.SetOpType(opline.GetResultType())
			if result.GetOpType() == 1<<0 {
				var _z1 *Zval = &result.u.constant
				var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetResult().GetConstant()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				result.SetOp(opline.GetResult())
			}
		}
	} else if j == 2 {
		opline.SetOpcode(53)
		opline.SetExtendedValue(0)
		opline.SetOp1Type(init_opline.GetOp2Type())
		opline.SetOp1(init_opline.GetOp2())
		opline.SetResultType(1 << 1)
		opline.GetResult().SetVar(GetTemporaryVariable())
		init_opline.GetOp1().SetNum(0)
		init_opline.GetOp2().SetNum(0)
		init_opline.GetResult().SetNum(0)
		init_opline.SetOpcode(0)
		init_opline.SetOp1Type(0)
		init_opline.SetOp2Type(0)
		init_opline.SetResultType(0)
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == 1<<0 {
			var _z1 *Zval = &result.u.constant
			var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetResult().GetConstant()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			result.SetOp(opline.GetResult())
		}
	} else {
		var var_ uint32
		init_opline.SetExtendedValue(j)
		opline.SetOpcode(56)
		opline.GetResult().SetVar(GetTemporaryVariable())
		opline.GetOp1().SetVar(GetTemporaryVariable())
		var_ = opline.GetOp1().GetVar()
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == 1<<0 {
			var _z1 *Zval = &result.u.constant
			var _z2 *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetResult().GetConstant()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			result.SetOp(opline.GetResult())
		}

		/* Allocates the necessary number of zval slots to keep the rope */

		i = (j*g.SizeOf("zend_string *") + (g.SizeOf("zval") - 1)) / g.SizeOf("zval")
		for i > 1 {
			GetTemporaryVariable()
			i--
		}

		/* Update all the previous opcodes to use the same variable */

		for opline != init_opline {
			opline--
			if opline.GetOpcode() == 55 && opline.GetResult().GetVar() == uint32-1 {
				opline.GetOp1().SetVar(var_)
				opline.GetResult().SetVar(var_)
			} else if opline.GetOpcode() == 54 && opline.GetResult().GetVar() == uint32-1 {
				opline.GetResult().SetVar(var_)
			}
		}

		/* Update all the previous opcodes to use the same variable */

	}
}

/* }}} */

func ZendCompileMagicConst(result *Znode, ast *ZendAst) {
	var opline *ZendOp
	if ZendTryCtEvalMagicConst(&result.u.constant, ast) != 0 {
		result.SetOpType(1 << 0)
		return
	}
	assert(ast.GetAttr() == T_CLASS_C && CG.GetActiveClassEntry() != nil && (CG.GetActiveClassEntry().GetCeFlags()&1<<1) != 0)
	opline = ZendEmitOpTmp(result, 157, nil, nil)
	opline.GetOp1().SetNum(1)
}

/* }}} */

func ZendIsAllowedInConstExpr(kind ZendAstKind) ZendBool {
	return kind == ZEND_AST_ZVAL || kind == ZEND_AST_BINARY_OP || kind == ZEND_AST_GREATER || kind == ZEND_AST_GREATER_EQUAL || kind == ZEND_AST_AND || kind == ZEND_AST_OR || kind == ZEND_AST_UNARY_OP || kind == ZEND_AST_UNARY_PLUS || kind == ZEND_AST_UNARY_MINUS || kind == ZEND_AST_CONDITIONAL || kind == ZEND_AST_DIM || kind == ZEND_AST_ARRAY || kind == ZEND_AST_ARRAY_ELEM || kind == ZEND_AST_UNPACK || kind == ZEND_AST_CONST || kind == ZEND_AST_CLASS_CONST || kind == ZEND_AST_CLASS_NAME || kind == ZEND_AST_MAGIC_CONST || kind == ZEND_AST_COALESCE
}

/* }}} */

func ZendCompileConstExprClassConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var class_ast *ZendAst = ast.GetChild()[0]
	var const_ast *ZendAst = ast.GetChild()[1]
	var class_name *ZendString
	var const_name *ZendString = ZendAstGetStr(const_ast)
	var name *ZendString
	var fetch_type int
	if class_ast.GetKind() != ZEND_AST_ZVAL {
		ZendErrorNoreturn(1<<6, "Dynamic class names are not allowed in compile-time class constant references")
	}
	class_name = ZendAstGetStr(class_ast)
	fetch_type = ZendGetClassFetchType(class_name)
	if 3 == fetch_type {
		ZendErrorNoreturn(1<<6, "\"static::\" is not allowed in compile-time constants")
	}
	if 0 == fetch_type {
		class_name = ZendResolveClassNameAst(class_ast)
	} else {
		ZendStringAddref(class_name)
	}
	name = ZendConcat3(class_name.GetVal(), class_name.GetLen(), "::", 2, const_name.GetVal(), const_name.GetLen())
	ZendAstDestroy(ast)
	ZendStringReleaseEx(class_name, 0)
	*ast_ptr = ZendAstCreateConstant(name, fetch_type|0x200)
}

/* }}} */

func ZendCompileConstExprClassName(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var class_ast *ZendAst = ast.GetChild()[0]
	var class_name *ZendString = ZendAstGetStr(class_ast)
	var fetch_type uint32 = ZendGetClassFetchType(class_name)
	switch fetch_type {
	case 1:

	case 2:

		/* For the const-eval representation store the fetch type instead of the name. */

		ZendStringRelease(class_name)
		ast.GetChild()[0] = nil
		ast.SetAttr(fetch_type)
		return
	case 3:
		ZendErrorNoreturn(1<<6, "static::class cannot be used for compile-time class name resolution")
		return
	default:
		break
	}
}
func ZendCompileConstExprConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var name_ast *ZendAst = ast.GetChild()[0]
	var orig_name *ZendString = ZendAstGetStr(name_ast)
	var is_fully_qualified ZendBool
	var result Zval
	var resolved_name *ZendString
	resolved_name = ZendResolveConstName(orig_name, name_ast.GetAttr(), &is_fully_qualified)
	if ZendTryCtEvalConst(&result, resolved_name, is_fully_qualified) != 0 {
		ZendStringReleaseEx(resolved_name, 0)
		ZendAstDestroy(ast)
		*ast_ptr = ZendAstCreateZval(&result)
		return
	}
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateConstant(resolved_name, g.Cond(is_fully_qualified == 0, 0x10, 0))
}

/* }}} */

func ZendCompileConstExprMagicConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr

	/* Other cases already resolved by constant folding */

	assert(ast.GetAttr() == T_CLASS_C)
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreate0(ZEND_AST_CONSTANT_CLASS)
}

/* }}} */

func ZendCompileConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	if ast == nil || ast.GetKind() == ZEND_AST_ZVAL {
		return
	}
	if ZendIsAllowedInConstExpr(ast.GetKind()) == 0 {
		ZendErrorNoreturn(1<<6, "Constant expression contains invalid operations")
	}
	switch ast.GetKind() {
	case ZEND_AST_CLASS_CONST:
		ZendCompileConstExprClassConst(ast_ptr)
		break
	case ZEND_AST_CLASS_NAME:
		ZendCompileConstExprClassName(ast_ptr)
		break
	case ZEND_AST_CONST:
		ZendCompileConstExprConst(ast_ptr)
		break
	case ZEND_AST_MAGIC_CONST:
		ZendCompileConstExprMagicConst(ast_ptr)
		break
	default:
		ZendAstApply(ast, ZendCompileConstExpr)
		break
	}
}

/* }}} */

func ZendConstExprToZval(result *Zval, ast *ZendAst) {
	var orig_ast *ZendAst = ast
	ZendEvalConstExpr(&ast)
	ZendCompileConstExpr(&ast)
	if ast.GetKind() == ZEND_AST_ZVAL {
		var _z1 *Zval = result
		var _z2 *Zval = ZendAstGetZval(ast)
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		var __z *Zval = result
		__z.GetValue().SetAst(ZendAstCopy(ast))
		__z.SetTypeInfo(11 | 1<<0<<8)

		/* destroy the ast here, it might have been replaced */

		ZendAstDestroy(ast)

		/* destroy the ast here, it might have been replaced */

	}

	/* Kill this branch of the original AST, as it was already destroyed.
	 * It would be nice to find a better solution to this problem in the
	 * future. */

	orig_ast.SetKind(0)

	/* Kill this branch of the original AST, as it was already destroyed.
	 * It would be nice to find a better solution to this problem in the
	 * future. */
}

/* }}} */

func ZendCompileTopStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_STMT_LIST {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			ZendCompileTopStmt(list.GetChild()[i])
		}
		return
	}
	if ast.GetKind() == ZEND_AST_FUNC_DECL {
		CG.SetZendLineno(ast.GetLineno())
		ZendCompileFuncDecl(nil, ast, 1)
		CG.SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
	} else if ast.GetKind() == ZEND_AST_CLASS {
		CG.SetZendLineno(ast.GetLineno())
		ZendCompileClassDecl(ast, 1)
		CG.SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
	} else {
		ZendCompileStmt(ast)
	}
	if ast.GetKind() != ZEND_AST_NAMESPACE && ast.GetKind() != ZEND_AST_HALT_COMPILER {
		ZendVerifyNamespace()
	}
}

/* }}} */

func ZendCompileStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	CG.SetZendLineno(ast.GetLineno())
	if (CG.GetCompilerOptions()&1<<0) != 0 && ZendIsUntickedStmt(ast) == 0 {
		ZendDoExtendedStmt()
	}
	switch ast.GetKind() {
	case ZEND_AST_STMT_LIST:
		ZendCompileStmtList(ast)
		break
	case ZEND_AST_GLOBAL:
		ZendCompileGlobalVar(ast)
		break
	case ZEND_AST_STATIC:
		ZendCompileStaticVar(ast)
		break
	case ZEND_AST_UNSET:
		ZendCompileUnset(ast)
		break
	case ZEND_AST_RETURN:
		ZendCompileReturn(ast)
		break
	case ZEND_AST_ECHO:
		ZendCompileEcho(ast)
		break
	case ZEND_AST_THROW:
		ZendCompileThrow(ast)
		break
	case ZEND_AST_BREAK:

	case ZEND_AST_CONTINUE:
		ZendCompileBreakContinue(ast)
		break
	case ZEND_AST_GOTO:
		ZendCompileGoto(ast)
		break
	case ZEND_AST_LABEL:
		ZendCompileLabel(ast)
		break
	case ZEND_AST_WHILE:
		ZendCompileWhile(ast)
		break
	case ZEND_AST_DO_WHILE:
		ZendCompileDoWhile(ast)
		break
	case ZEND_AST_FOR:
		ZendCompileFor(ast)
		break
	case ZEND_AST_FOREACH:
		ZendCompileForeach(ast)
		break
	case ZEND_AST_IF:
		ZendCompileIf(ast)
		break
	case ZEND_AST_SWITCH:
		ZendCompileSwitch(ast)
		break
	case ZEND_AST_TRY:
		ZendCompileTry(ast)
		break
	case ZEND_AST_DECLARE:
		ZendCompileDeclare(ast)
		break
	case ZEND_AST_FUNC_DECL:

	case ZEND_AST_METHOD:
		ZendCompileFuncDecl(nil, ast, 0)
		break
	case ZEND_AST_PROP_GROUP:
		ZendCompilePropGroup(ast)
		break
	case ZEND_AST_CLASS_CONST_DECL:
		ZendCompileClassConstDecl(ast)
		break
	case ZEND_AST_USE_TRAIT:
		ZendCompileUseTrait(ast)
		break
	case ZEND_AST_CLASS:
		ZendCompileClassDecl(ast, 0)
		break
	case ZEND_AST_GROUP_USE:
		ZendCompileGroupUse(ast)
		break
	case ZEND_AST_USE:
		ZendCompileUse(ast)
		break
	case ZEND_AST_CONST_DECL:
		ZendCompileConstDecl(ast)
		break
	case ZEND_AST_NAMESPACE:
		ZendCompileNamespace(ast)
		break
	case ZEND_AST_HALT_COMPILER:
		ZendCompileHaltCompiler(ast)
		break
	default:
		var result Znode
		ZendCompileExpr(&result, ast)
		ZendDoFree(&result)
	}
	if CG.GetFileContext().GetDeclarables().GetTicks() != 0 && ZendIsUntickedStmt(ast) == 0 {
		ZendEmitTick()
	}
}

/* }}} */

func ZendCompileExpr(result *Znode, ast *ZendAst) {
	/* CG(zend_lineno) = ast->lineno; */

	CG.SetZendLineno(ZendAstGetLineno(ast))
	if CG.GetMemoizeMode() != 0 {
		ZendCompileMemoizedExpr(result, ast)
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_ZVAL:
		var _z1 *Zval = &result.u.constant
		var _z2 *Zval = ZendAstGetZval(ast)
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		result.SetOpType(1 << 0)
		return
	case ZEND_AST_ZNODE:
		*result = (*ZendAstGetZnode)(ast)
		return
	case ZEND_AST_VAR:

	case ZEND_AST_DIM:

	case ZEND_AST_PROP:

	case ZEND_AST_STATIC_PROP:

	case ZEND_AST_CALL:

	case ZEND_AST_METHOD_CALL:

	case ZEND_AST_STATIC_CALL:
		ZendCompileVar(result, ast, 0, 0)
		return
	case ZEND_AST_ASSIGN:
		ZendCompileAssign(result, ast)
		return
	case ZEND_AST_ASSIGN_REF:
		ZendCompileAssignRef(result, ast)
		return
	case ZEND_AST_NEW:
		ZendCompileNew(result, ast)
		return
	case ZEND_AST_CLONE:
		ZendCompileClone(result, ast)
		return
	case ZEND_AST_ASSIGN_OP:
		ZendCompileCompoundAssign(result, ast)
		return
	case ZEND_AST_BINARY_OP:
		ZendCompileBinaryOp(result, ast)
		return
	case ZEND_AST_GREATER:

	case ZEND_AST_GREATER_EQUAL:
		ZendCompileGreater(result, ast)
		return
	case ZEND_AST_UNARY_OP:
		ZendCompileUnaryOp(result, ast)
		return
	case ZEND_AST_UNARY_PLUS:

	case ZEND_AST_UNARY_MINUS:
		ZendCompileUnaryPm(result, ast)
		return
	case ZEND_AST_AND:

	case ZEND_AST_OR:
		ZendCompileShortCircuiting(result, ast)
		return
	case ZEND_AST_POST_INC:

	case ZEND_AST_POST_DEC:
		ZendCompilePostIncdec(result, ast)
		return
	case ZEND_AST_PRE_INC:

	case ZEND_AST_PRE_DEC:
		ZendCompilePreIncdec(result, ast)
		return
	case ZEND_AST_CAST:
		ZendCompileCast(result, ast)
		return
	case ZEND_AST_CONDITIONAL:
		ZendCompileConditional(result, ast)
		return
	case ZEND_AST_COALESCE:
		ZendCompileCoalesce(result, ast)
		return
	case ZEND_AST_ASSIGN_COALESCE:
		ZendCompileAssignCoalesce(result, ast)
		return
	case ZEND_AST_PRINT:
		ZendCompilePrint(result, ast)
		return
	case ZEND_AST_EXIT:
		ZendCompileExit(result, ast)
		return
	case ZEND_AST_YIELD:
		ZendCompileYield(result, ast)
		return
	case ZEND_AST_YIELD_FROM:
		ZendCompileYieldFrom(result, ast)
		return
	case ZEND_AST_INSTANCEOF:
		ZendCompileInstanceof(result, ast)
		return
	case ZEND_AST_INCLUDE_OR_EVAL:
		ZendCompileIncludeOrEval(result, ast)
		return
	case ZEND_AST_ISSET:

	case ZEND_AST_EMPTY:
		ZendCompileIssetOrEmpty(result, ast)
		return
	case ZEND_AST_SILENCE:
		ZendCompileSilence(result, ast)
		return
	case ZEND_AST_SHELL_EXEC:
		ZendCompileShellExec(result, ast)
		return
	case ZEND_AST_ARRAY:
		ZendCompileArray(result, ast)
		return
	case ZEND_AST_CONST:
		ZendCompileConst(result, ast)
		return
	case ZEND_AST_CLASS_CONST:
		ZendCompileClassConst(result, ast)
		return
	case ZEND_AST_CLASS_NAME:
		ZendCompileClassName(result, ast)
		return
	case ZEND_AST_ENCAPS_LIST:
		ZendCompileEncapsList(result, ast)
		return
	case ZEND_AST_MAGIC_CONST:
		ZendCompileMagicConst(result, ast)
		return
	case ZEND_AST_CLOSURE:

	case ZEND_AST_ARROW_FUNC:
		ZendCompileFuncDecl(result, ast, 0)
		return
	default:
		assert(false)
	}
}

/* }}} */

func ZendCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *ZendOp {
	CG.SetZendLineno(ZendAstGetLineno(ast))
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return ZendCompileSimpleVar(result, ast, type_, 0)
	case ZEND_AST_DIM:
		return ZendCompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		return ZendCompileProp(result, ast, type_, by_ref)
	case ZEND_AST_STATIC_PROP:
		return ZendCompileStaticProp(result, ast, type_, by_ref, 0)
	case ZEND_AST_CALL:
		ZendCompileCall(result, ast, type_)
		return nil
	case ZEND_AST_METHOD_CALL:
		ZendCompileMethodCall(result, ast, type_)
		return nil
	case ZEND_AST_STATIC_CALL:
		ZendCompileStaticCall(result, ast, type_)
		return nil
	case ZEND_AST_ZNODE:
		*result = (*ZendAstGetZnode)(ast)
		return nil
	default:
		if type_ == 1 || type_ == 2 || type_ == 5 {
			ZendErrorNoreturn(1<<6, "Cannot use temporary expression in write context")
		}
		ZendCompileExpr(result, ast)
		return nil
	}
}

/* }}} */

func ZendDelayedCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref ZendBool) *ZendOp {
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return ZendCompileSimpleVar(result, ast, type_, 1)
	case ZEND_AST_DIM:
		return ZendDelayedCompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		var opline *ZendOp = ZendDelayedCompileProp(result, ast, type_)
		if by_ref != 0 {
			opline.SetExtendedValue(opline.GetExtendedValue() | 1)
		}
		return opline
	case ZEND_AST_STATIC_PROP:
		return ZendCompileStaticProp(result, ast, type_, by_ref, 1)
	default:
		return ZendCompileVar(result, ast, type_, 0)
	}
}

/* }}} */

func ZendEvalConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var result Zval
	if ast == nil {
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		if ZendTryCtEvalBinaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1])) == 0 {
			return
		}
		break
	case ZEND_AST_GREATER:

	case ZEND_AST_GREATER_EQUAL:
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalGreater(&result, ast.GetKind(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1]))
		break
	case ZEND_AST_AND:

	case ZEND_AST_OR:
		var child0_is_true ZendBool
		var child1_is_true ZendBool
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		child0_is_true = ZendIsTrue(ZendAstGetZval(ast.GetChild()[0]))
		if child0_is_true == (ast.GetKind() == ZEND_AST_OR) {
			if ast.GetKind() == ZEND_AST_OR {
				&result.SetTypeInfo(3)
			} else {
				&result.SetTypeInfo(2)
			}
			break
		}
		if ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		child1_is_true = ZendIsTrue(ZendAstGetZval(ast.GetChild()[1]))
		if ast.GetKind() == ZEND_AST_OR {
			if child0_is_true != 0 || child1_is_true != 0 {
				&result.SetTypeInfo(3)
			} else {
				&result.SetTypeInfo(2)
			}
		} else {
			if child0_is_true != 0 && child1_is_true != 0 {
				&result.SetTypeInfo(3)
			} else {
				&result.SetTypeInfo(2)
			}
		}
		break
	case ZEND_AST_UNARY_OP:
		ZendEvalConstExpr(&ast.child[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalUnaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]))
		break
	case ZEND_AST_UNARY_PLUS:

	case ZEND_AST_UNARY_MINUS:
		ZendEvalConstExpr(&ast.child[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		if ZendTryCtEvalUnaryPm(&result, ast.GetKind(), ZendAstGetZval(ast.GetChild()[0])) == 0 {
			return
		}
		break
	case ZEND_AST_COALESCE:

		/* Set isset fetch indicator here, opcache disallows runtime altering of the AST */

		if ast.GetChild()[0].GetKind() == ZEND_AST_DIM {
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | 1<<0)
		}
		ZendEvalConstExpr(&ast.child[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			ZendEvalConstExpr(&ast.child[1])
			return
		}
		if ZendAstGetZval(ast.GetChild()[0]).GetType() == 1 {
			ZendEvalConstExpr(&ast.child[1])
			*ast_ptr = ast.GetChild()[1]
			ast.GetChild()[1] = nil
			ZendAstDestroy(ast)
		} else {
			*ast_ptr = ast.GetChild()[0]
			ast.GetChild()[0] = nil
			ZendAstDestroy(ast)
		}
		return
	case ZEND_AST_CONDITIONAL:
		var child **ZendAst
		var child_ast **ZendAst
		ZendEvalConstExpr(&ast.child[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			if ast.GetChild()[1] != nil {
				ZendEvalConstExpr(&ast.child[1])
			}
			ZendEvalConstExpr(&ast.child[2])
			return
		}
		child = &ast.child[2-ZendIsTrue(ZendAstGetZval(ast.GetChild()[0]))]
		if (*child) == nil {
			child--
		}
		child_ast = *child
		*child = nil
		ZendAstDestroy(ast)
		*ast_ptr = child_ast
		ZendEvalConstExpr(ast_ptr)
		return
	case ZEND_AST_DIM:

		/* constant expression should be always read context ... */

		var container *Zval
		var dim *Zval
		if ast.GetChild()[1] == nil {
			ZendErrorNoreturn(1<<6, "Cannot use [] for reading")
		}
		if (ast.GetAttr() & 1 << 1) != 0 {
			ast.SetAttr(ast.GetAttr() &^ (1 << 1))
			ZendError(1<<13, "Array and string offset access syntax with curly braces is deprecated")
		}

		/* Set isset fetch indicator here, opcache disallows runtime altering of the AST */

		if (ast.GetAttr()&1<<0) != 0 && ast.GetChild()[0].GetKind() == ZEND_AST_DIM {
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | 1<<0)
		}
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		container = ZendAstGetZval(ast.GetChild()[0])
		dim = ZendAstGetZval(ast.GetChild()[1])
		if container.GetType() == 7 {
			var el *Zval
			if dim.GetType() == 4 {
				el = ZendHashIndexFind(container.GetValue().GetArr(), dim.GetValue().GetLval())
				if el != nil {
					var _z1 *Zval = &result
					var _z2 *Zval = el
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				} else {
					return
				}
			} else if dim.GetType() == 6 {
				el = ZendSymtableFind(container.GetValue().GetArr(), dim.GetValue().GetStr())
				if el != nil {
					var _z1 *Zval = &result
					var _z2 *Zval = el
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				} else {
					return
				}
			} else {
				return
			}
		} else if container.GetType() == 6 {
			var offset ZendLong
			var c ZendUchar
			if dim.GetType() == 4 {
				offset = dim.GetValue().GetLval()
			} else if dim.GetType() != 6 || IsNumericString(dim.GetValue().GetStr().GetVal(), dim.GetValue().GetStr().GetLen(), &offset, nil, 1) != 4 {
				return
			}
			if offset < 0 || int(offset >= container.GetValue().GetStr().GetLen()) != 0 {
				return
			}
			c = zend_uchar(container.GetValue().GetStr()).val[offset]
			var __z *Zval = &result
			var __s *ZendString = ZendOneCharString[c]
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		} else if container.GetType() <= 2 {
			&result.SetTypeInfo(1)
		} else {
			return
		}
		break
	case ZEND_AST_ARRAY:
		if ZendTryCtEvalArray(&result, ast) == 0 {
			return
		}
		break
	case ZEND_AST_MAGIC_CONST:
		if ZendTryCtEvalMagicConst(&result, ast) == 0 {
			return
		}
		break
	case ZEND_AST_CONST:
		var name_ast *ZendAst = ast.GetChild()[0]
		var is_fully_qualified ZendBool
		var resolved_name *ZendString = ZendResolveConstName(ZendAstGetStr(name_ast), name_ast.GetAttr(), &is_fully_qualified)
		if ZendTryCtEvalConst(&result, resolved_name, is_fully_qualified) == 0 {
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
		ZendStringReleaseEx(resolved_name, 0)
		break
	case ZEND_AST_CLASS_CONST:
		var class_ast *ZendAst
		var name_ast *ZendAst
		var resolved_name *ZendString
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		class_ast = ast.GetChild()[0]
		name_ast = ast.GetChild()[1]
		if class_ast.GetKind() != ZEND_AST_ZVAL || name_ast.GetKind() != ZEND_AST_ZVAL {
			return
		}
		resolved_name = ZendResolveClassNameAst(class_ast)
		if ZendTryCtEvalClassConst(&result, resolved_name, ZendAstGetStr(name_ast)) == 0 {
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
		ZendStringReleaseEx(resolved_name, 0)
		break
	case ZEND_AST_CLASS_NAME:
		var class_ast *ZendAst = ast.GetChild()[0]
		if ZendTryCompileConstExprResolveClassName(&result, class_ast) == 0 {
			return
		}
		break
	default:
		return
	}
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateZval(&result)
}

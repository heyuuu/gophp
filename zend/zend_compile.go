// <<generate>>

package zend

import (
	b "sik/builtin"
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

func MAKE_NOP(opline *ZendOp) {
	opline.GetOp1().SetNum(0)
	opline.GetOp2().SetNum(0)
	opline.GetResult().SetNum(0)
	opline.SetOpcode(ZEND_NOP)
	opline.SetOp1Type(IS_UNUSED)
	opline.SetOp2Type(IS_UNUSED)
	opline.SetResultType(IS_UNUSED)
}
func RESET_DOC_COMMENT() {
	if CG(doc_comment) {
		ZendStringReleaseEx(CG(doc_comment), 0)
		CG(doc_comment) = nil
	}
}

/* On 64-bit systems less optimal, but more compact VM code leads to better
 * performance. So on 32-bit systems we use absolute addresses for jump
 * targets and constants, but on 64-bit systems realtive 32-bit offsets */

const ZEND_USE_ABS_JMP_ADDR = 0
const ZEND_USE_ABS_CONST_ADDR = 0

/* Temporarily defined here, to avoid header ordering issues */

func ZendAstGetZnode(ast *ZendAst) *Znode { return &((*ZendAstZnode)(ast)).node }

/* Compilation context that is different for each file, but shared between op arrays. */

type UserOpcodeHandlerT func(execute_data *ZendExecuteData) int

const ZEND_LIVE_TMPVAR = 0
const ZEND_LIVE_LOOP = 1
const ZEND_LIVE_SILENCE = 2
const ZEND_LIVE_ROPE = 3
const ZEND_LIVE_NEW = 4
const ZEND_LIVE_MASK = 7

/* Compilation context that is different for each op array. */

/* Class, property and method flags                  class|meth.|prop.|const*/

const ZEND_ACC_PUBLIC uint32 = 1 << 0
const ZEND_ACC_PROTECTED = 1 << 1
const ZEND_ACC_PRIVATE = 1 << 2

/*                                                        |     |     |     */

const ZEND_ACC_CHANGED = 1 << 3

/*                                                        |     |     |     */

const ZEND_ACC_STATIC = 1 << 4

/*                                                        |     |     |     */

const ZEND_ACC_FINAL = 1 << 5

/*                                                        |     |     |     */

const ZEND_ACC_ABSTRACT = 1 << 6
const ZEND_ACC_EXPLICIT_ABSTRACT_CLASS = 1 << 6

/*                                                        |     |     |     */

const ZEND_ACC_IMMUTABLE = 1 << 7

/*                                                        |     |     |     */

const ZEND_ACC_HAS_TYPE_HINTS = 1 << 8

/*                                                        |     |     |     */

const ZEND_ACC_TOP_LEVEL = 1 << 9

/*                                                        |     |     |     */

const ZEND_ACC_PRELOADED = 1 << 10

/*                                                        |     |     |     */

const ZEND_ACC_INTERFACE uint32 = 1 << 0
const ZEND_ACC_TRAIT = 1 << 1
const ZEND_ACC_ANON_CLASS = 1 << 2

/*                                                        |     |     |     */

const ZEND_ACC_LINKED = 1 << 3

/*                                                        |     |     |     */

const ZEND_ACC_IMPLICIT_ABSTRACT_CLASS = 1 << 4

/*                                                        |     |     |     */

const ZEND_ACC_USE_GUARDS = 1 << 11

/*                                                        |     |     |     */

const ZEND_ACC_CONSTANTS_UPDATED uint32 = 1 << 12

/*                                                        |     |     |     */

const ZEND_ACC_INHERITED = 1 << 13

/*                                                        |     |     |     */

const ZEND_ACC_IMPLEMENT_INTERFACES = 1 << 14

/*                                                        |     |     |     */

const ZEND_ACC_IMPLEMENT_TRAITS = 1 << 15

/*                                                        |     |     |     */

const ZEND_HAS_STATIC_IN_METHODS = 1 << 16

/*                                                        |     |     |     */

const ZEND_ACC_PROPERTY_TYPES_RESOLVED = 1 << 17

/*                                                        |     |     |     */

const ZEND_ACC_REUSE_GET_ITERATOR = 1 << 18

/*                                                        |     |     |     */

const ZEND_ACC_RESOLVED_PARENT = 1 << 19

/*                                                        |     |     |     */

const ZEND_ACC_RESOLVED_INTERFACES = 1 << 20

/*                                                        |     |     |     */

const ZEND_ACC_UNRESOLVED_VARIANCE = 1 << 21

/*                                                        |     |     |     */

const ZEND_ACC_NEARLY_LINKED = 1 << 22

/*                                                        |     |     |     */

const ZEND_ACC_HAS_UNLINKED_USES = 1 << 23

/*                                                        |     |     |     */

const ZEND_ACC_DEPRECATED = 1 << 11

/*                                                        |     |     |     */

const ZEND_ACC_RETURN_REFERENCE = 1 << 12

/*                                                        |     |     |     */

const ZEND_ACC_HAS_RETURN_TYPE = 1 << 13

/*                                                        |     |     |     */

const ZEND_ACC_VARIADIC = 1 << 14

/*                                                        |     |     |     */

const ZEND_ACC_HAS_FINALLY_BLOCK = 1 << 15

/*                                                        |     |     |     */

const ZEND_ACC_EARLY_BINDING = 1 << 16

/*                                                        |     |     |     */

const ZEND_ACC_ALLOW_STATIC = 1 << 17

/*                                                        |     |     |     */

const ZEND_ACC_CALL_VIA_TRAMPOLINE = 1 << 18

/*                                                        |     |     |     */

const ZEND_ACC_NEVER_CACHE = 1 << 19

/*                                                        |     |     |     */

const ZEND_ACC_CLOSURE = 1 << 20
const ZEND_ACC_FAKE_CLOSURE = 1 << 21

/*                                                        |     |     |     */

const ZEND_ACC_HEAP_RT_CACHE = 1 << 22

/*                                                        |     |     |     */

const ZEND_ACC_USER_ARG_INFO = 1 << 22

/*                                                        |     |     |     */

const ZEND_ACC_GENERATOR = 1 << 24

/*                                                        |     |     |     */

const ZEND_ACC_DONE_PASS_TWO = 1 << 25

/*                                                        |     |     |     */

const ZEND_ACC_ARENA_ALLOCATED = 1 << 25

/*                                                        |     |     |     */

const ZEND_ACC_TRAIT_CLONE = 1 << 27

/*                                                        |     |     |     */

const ZEND_ACC_CTOR = 1 << 28

/*                                                        |     |     |     */

const ZEND_ACC_DTOR = 1 << 29

/*                                                        |     |     |     */

const ZEND_ACC_USES_THIS = 1 << 30

/*                                                        |     |     |     */

const ZEND_ACC_STRICT_TYPES = 1 << 31
const ZEND_ACC_PPP_MASK = ZEND_ACC_PUBLIC | ZEND_ACC_PROTECTED | ZEND_ACC_PRIVATE

/* call through internal function handler. e.g. Closure::invoke() */

const ZEND_ACC_CALL_VIA_HANDLER = ZEND_ACC_CALL_VIA_TRAMPOLINE

func OBJ_PROP(obj *ZendObject, offset *ZendObject) *Zval { return (*Zval)((*byte)(obj + offset)) }
func OBJ_PROP_NUM(obj __auto__, num __auto__) __auto__   { return &obj.properties_table[num] }
func OBJ_PROP_TO_OFFSET(num int) __auto__ {
	return uint32_t(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil)) + b.SizeOf("zval")*num)
}
func OBJ_PROP_TO_NUM(offset uint32) int {
	return (offset - OBJ_PROP_TO_OFFSET(0)) / b.SizeOf("zval")
}

/* arg_info for internal functions */

/* arg_info for user functions */

/* the following structure repeats the layout of zend_internal_arg_info,
 * but its fields have different meaning. It's used as the first element of
 * arg_info array to define properties __special__  of internal functions.
 * It's also used for the return type.
 */

const ZEND_RETURN_VALUE = 0
const ZEND_RETURN_REFERENCE = 1

/* zend_internal_function_handler */

type ZifHandler func(execute_data *ZendExecuteData, return_value *Zval)

func ZEND_FN_SCOPE_NAME(function *ZendFunction) string {
	if function != nil && function.GetScope() != nil {
		return ZSTR_VAL(function.GetScope().GetName())
	} else {
		return ""
	}
}

const ZEND_CALL_HAS_THIS = IS_OBJECT_EX

/* Top 16 bits of Z_TYPE_INFO(EX(This)) are used as call_info flags */

const ZEND_CALL_FUNCTION uint32 = 0 << 16
const ZEND_CALL_CODE = 1 << 16
const ZEND_CALL_NESTED = 0 << 17
const ZEND_CALL_TOP uint32 = 1 << 17
const ZEND_CALL_ALLOCATED uint32 = 1 << 18
const ZEND_CALL_FREE_EXTRA_ARGS uint32 = 1 << 19
const ZEND_CALL_HAS_SYMBOL_TABLE uint32 = 1 << 20
const ZEND_CALL_RELEASE_THIS = 1 << 21
const ZEND_CALL_CLOSURE uint32 = 1 << 22
const ZEND_CALL_FAKE_CLOSURE = 1 << 23
const ZEND_CALL_GENERATOR = 1 << 24
const ZEND_CALL_DYNAMIC = 1 << 25
const ZEND_CALL_SEND_ARG_BY_REF uint32 = 1 << 31
const ZEND_CALL_NESTED_FUNCTION uint32 = ZEND_CALL_FUNCTION | ZEND_CALL_NESTED
const ZEND_CALL_NESTED_CODE = ZEND_CALL_CODE | ZEND_CALL_NESTED
const ZEND_CALL_TOP_FUNCTION = ZEND_CALL_TOP | ZEND_CALL_FUNCTION
const ZEND_CALL_TOP_CODE = ZEND_CALL_CODE | ZEND_CALL_TOP

func ZEND_CALL_INFO(call *ZendExecuteData) uint32 { return Z_TYPE_INFO(call.GetThis()) }
func ZEND_CALL_KIND_EX(call_info uint32) int {
	return call_info & (ZEND_CALL_CODE | ZEND_CALL_TOP)
}
func ZEND_CALL_KIND(call *ZendExecuteData) int {
	return ZEND_CALL_KIND_EX(ZEND_CALL_INFO(call))
}
func ZEND_ADD_CALL_FLAG_EX(call_info uint32, flag int)    { call_info |= flag }
func ZEND_DEL_CALL_FLAG_EX(call_info uint32, flag uint32) { call_info &= ^flag }
func ZEND_ADD_CALL_FLAG(call *ZendExecuteData, flag uint32) {
	ZEND_ADD_CALL_FLAG_EX(Z_TYPE_INFO(call.GetThis()), flag)
}
func ZEND_DEL_CALL_FLAG(call __auto__, flag uint32) {
	ZEND_DEL_CALL_FLAG_EX(Z_TYPE_INFO(call.This), flag)
}
func ZEND_CALL_NUM_ARGS(call *ZendExecuteData) uint32 { return call.GetThis().GetNumArgs() }

const ZEND_CALL_FRAME_SLOT = int((ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_execute_data")) + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")))

func ZEND_CALL_VAR(call *ZendExecuteData, n uint32) *Zval { return (*Zval)((*byte)(call) + int(n)) }
func ZEND_CALL_VAR_NUM(call *ZendExecuteData, n int) *Zval {
	return (*Zval)(call) + (ZEND_CALL_FRAME_SLOT + int(n))
}
func ZEND_CALL_ARG(call *ZendExecuteData, n int) *Zval {
	return ZEND_CALL_VAR_NUM(call, int(n)-1)
}
func EX(element __auto__) __auto__ { return execute_data.element }
func EX_CALL_INFO() uint32         { return ZEND_CALL_INFO(execute_data) }
func EX_CALL_KIND() int            { return ZEND_CALL_KIND(execute_data) }
func EX_NUM_ARGS() uint32          { return ZEND_CALL_NUM_ARGS(execute_data) }
func ZEND_CALL_USES_STRICT_TYPES(call *ZendExecuteData) bool {
	return (call.GetFunc().GetFnFlags() & ZEND_ACC_STRICT_TYPES) != 0
}
func EX_USES_STRICT_TYPES() bool {
	return ZEND_CALL_USES_STRICT_TYPES(execute_data)
}
func ZEND_ARG_USES_STRICT_TYPES() bool {
	return EG(current_execute_data).prev_execute_data && EG(current_execute_data).prev_execute_data.func_ && ZEND_CALL_USES_STRICT_TYPES(EG(current_execute_data).prev_execute_data)
}
func ZEND_RET_USES_STRICT_TYPES() bool {
	return ZEND_CALL_USES_STRICT_TYPES(EG(current_execute_data))
}
func EX_VAR(n uint32) *Zval { return ZEND_CALL_VAR(execute_data, n) }
func EX_VAR_NUM(n int) *Zval {
	return ZEND_CALL_VAR_NUM(execute_data, n)
}
func EX_VAR_TO_NUM(n uint32) __auto__ {
	return uint32_t(ZEND_CALL_VAR(nil, n) - ZEND_CALL_VAR_NUM(nil, 0))
}
func ZEND_OPLINE_TO_OFFSET(opline __auto__, target *byte) *byte {
	return (*byte)(target - (*byte)(opline))
}
func ZEND_OPLINE_NUM_TO_OFFSET(op_array *ZendOpArray, opline *ZendOp, opline_num uint32) *byte {
	return (*byte)(&op_array.opcodes[opline_num] - (*byte)(opline))
}
func ZEND_OFFSET_TO_OPLINE(base *ZendOp, offset uint32) *ZendOp {
	return (*ZendOp)((*byte)(base) + int(offset))
}
func ZEND_OFFSET_TO_OPLINE_NUM(op_array __auto__, base *ZendOp, offset uint32) int {
	return ZEND_OFFSET_TO_OPLINE(base, offset) - op_array.opcodes
}

/* run-time jump target */

func OP_JMP_ADDR(opline *ZendOp, node ZnodeOp) *ZendOp {
	return ZEND_OFFSET_TO_OPLINE(opline, node.GetJmpOffset())
}
func ZEND_SET_OP_JMP_ADDR(opline __auto__, node __auto__, val *byte) {
	node.jmp_offset = ZEND_OPLINE_TO_OFFSET(opline, val)
}

/* convert jump target from compile-time to run-time */

func ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array *ZendOpArray, opline *ZendOp, node ZnodeOp) {
	node.SetJmpOffset(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, node.GetOplineNum()))
}

/* convert jump target back from run-time to compile-time */

func ZEND_PASS_TWO_UNDO_JMP_TARGET(op_array __auto__, opline *ZendOp, node __auto__) {
	node.opline_num = ZEND_OFFSET_TO_OPLINE_NUM(op_array, opline, node.jmp_offset)
}

/* constant-time constant */

func CT_CONSTANT_EX(op_array *ZendOpArray, num *Zval) __auto__ { return op_array.GetLiterals() + num }
func CT_CONSTANT(node ZnodeOp) __auto__ {
	return CT_CONSTANT_EX(CG(active_op_array), node.GetConstant())
}

/* At run-time, constants are allocated together with op_array->opcodes
 * and addressed relatively to current opline.
 */

func RT_CONSTANT(opline *ZendOp, node ZnodeOp) *Zval {
	return (*Zval)((*byte)(opline) + int32_t(node).constant)
}

/* convert constant from compile-time to run-time */

func ZEND_PASS_TWO_UPDATE_CONSTANT(op_array *ZendOpArray, opline *ZendOp, node ZnodeOp) {
	node.SetConstant((*byte)(CT_CONSTANT_EX(op_array, node.GetConstant())) - (*byte)(opline))
}

/* convert constant back from run-time to compile-time */

func ZEND_PASS_TWO_UNDO_CONSTANT(op_array __auto__, opline *ZendOp, node ZnodeOp) {
	node.SetConstant(RT_CONSTANT(opline, node) - op_array.literals)
}
func RUN_TIME_CACHE(op_array __auto__) any {
	return ZEND_MAP_PTR_GET(op_array.run_time_cache)
}
func ZEND_OP_ARRAY_EXTENSION(op_array __auto__, handle __auto__) any {
	return (*any)(RUN_TIME_CACHE(op_array))[handle]
}

const IS_UNUSED = 0
const IS_CONST ZendUchar = 1 << 0
const IS_TMP_VAR ZendUchar = 1 << 1
const IS_VAR ZendUchar = 1 << 2
const IS_CV ZendUchar = 1 << 3
const ZEND_EXTRA_VALUE = 1

// # include "zend_globals.h"

var ZendCompileFile func(file_handle *ZendFileHandle, type_ int) *ZendOpArray
var ZendCompileString func(source_string *Zval, filename *byte) *ZendOpArray

type UnaryOpType func(*Zval, *Zval) int
type BinaryOpType func(*Zval, *Zval, *Zval) int

/* Used during AST construction */

/* parser-driven code generators */

var ZendDoExtendedInfo func()

const INITIAL_OP_ARRAY_SIZE = 64

/* helper functions in zend_language_scanner.l */

func ZendTryExceptionHandler() {
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		if Z_TYPE(ExecutorGlobals.GetUserExceptionHandler()) != IS_UNDEF {
			ZendUserExceptionHandler()
		}
	}
}
func ZendUnmanglePropertyName(mangled_property *ZendString, class_name **byte, prop_name **byte) int {
	return ZendUnmanglePropertyNameEx(mangled_property, class_name, prop_name, nil)
}
func ZendGetUnmangledPropertyName(mangled_prop *ZendString) *byte {
	var class_name *byte
	var prop_name *byte
	ZendUnmanglePropertyName(mangled_prop, &class_name, &prop_name)
	return prop_name
}

const ZEND_FUNCTION_DTOR DtorFuncT = ZendFunctionDtor
const ZEND_CLASS_DTOR DtorFuncT = DestroyZendClass

type ZendNeedsLiveRangeCb func(op_array *ZendOpArray, opline *ZendOp) ZendBool
type ZendAutoGlobalCallback func(name *ZendString) ZendBool

/* BEGIN: OPCODES */

// # include "zend_vm_opcodes.h"

/* END: OPCODES */

const ZEND_FETCH_CLASS_DEFAULT = 0
const ZEND_FETCH_CLASS_SELF = 1
const ZEND_FETCH_CLASS_PARENT = 2
const ZEND_FETCH_CLASS_STATIC = 3
const ZEND_FETCH_CLASS_AUTO = 4
const ZEND_FETCH_CLASS_INTERFACE = 5
const ZEND_FETCH_CLASS_TRAIT = 6
const ZEND_FETCH_CLASS_MASK = 0xf
const ZEND_FETCH_CLASS_NO_AUTOLOAD = 0x80
const ZEND_FETCH_CLASS_SILENT = 0x100
const ZEND_FETCH_CLASS_EXCEPTION = 0x200
const ZEND_FETCH_CLASS_ALLOW_UNLINKED = 0x400
const ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED = 0x800
const ZEND_PARAM_REF = 1 << 0
const ZEND_PARAM_VARIADIC = 1 << 1
const ZEND_NAME_FQ = 0
const ZEND_NAME_NOT_FQ = 1
const ZEND_NAME_RELATIVE = 2
const ZEND_TYPE_NULLABLE = 1 << 8
const ZEND_ARRAY_SYNTAX_LIST = 1
const ZEND_ARRAY_SYNTAX_LONG = 2
const ZEND_ARRAY_SYNTAX_SHORT = 3

/* var status for backpatching */

const BP_VAR_R = 0
const BP_VAR_W = 1
const BP_VAR_RW = 2
const BP_VAR_IS = 3
const BP_VAR_FUNC_ARG = 4
const BP_VAR_UNSET = 5
const ZEND_INTERNAL_FUNCTION = 1
const ZEND_USER_FUNCTION = 2
const ZEND_OVERLOADED_FUNCTION = 3
const ZEND_EVAL_CODE = 4
const ZEND_OVERLOADED_FUNCTION_TEMPORARY = 5

/* A quick check (type == ZEND_USER_FUNCTION || type == ZEND_EVAL_CODE) */

func ZEND_USER_CODE(type_ ZendUchar) bool { return (type_ & 1) == 0 }

const ZEND_INTERNAL_CLASS = 1
const ZEND_USER_CLASS = 2
const ZEND_EVAL = 1 << 0
const ZEND_INCLUDE = 1 << 1
const ZEND_INCLUDE_ONCE = 1 << 2
const ZEND_REQUIRE = 1 << 3
const ZEND_REQUIRE_ONCE = 1 << 4

/* global/local fetches */

const ZEND_FETCH_GLOBAL uint32 = 1 << 1
const ZEND_FETCH_LOCAL uint32 = 1 << 2
const ZEND_FETCH_GLOBAL_LOCK uint32 = 1 << 3
const ZEND_FETCH_TYPE_MASK = 0xe

/* Only one of these can ever be in use */

const ZEND_FETCH_REF = 1
const ZEND_FETCH_DIM_WRITE = 2
const ZEND_FETCH_OBJ_WRITE = 3
const ZEND_FETCH_OBJ_FLAGS = 3
const ZEND_ISEMPTY = 1 << 0
const ZEND_LAST_CATCH = 1 << 0
const ZEND_FREE_ON_RETURN uint32 = 1 << 0
const ZEND_FREE_SWITCH uint32 = 1 << 1
const ZEND_SEND_BY_VAL = 0
const ZEND_SEND_BY_REF = 1
const ZEND_SEND_PREFER_REF = 2
const ZEND_DIM_IS = 1 << 0
const ZEND_DIM_ALTERNATIVE_SYNTAX = 1 << 1
const IS_CONSTANT_UNQUALIFIED = 0x10
const IS_CONSTANT_CLASS = 0x80
const IS_CONSTANT_IN_NAMESPACE = 0x100

func ZendCheckArgSendType(zf *ZendFunction, arg_num uint32, mask uint32) int {
	arg_num--
	if UNEXPECTED(arg_num >= zf.GetNumArgs()) {
		if EXPECTED((zf.GetFnFlags() & ZEND_ACC_VARIADIC) == 0) {
			return 0
		}
		arg_num = zf.GetNumArgs()
	}
	return UNEXPECTED((zf.GetArgInfo()[arg_num].GetPassByReference() & mask) != 0)
}
func ARG_MUST_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF)
}
func ARG_SHOULD_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func ARG_MAY_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_PREFER_REF)
}

/* Quick API to check first 12 arguments */

const MAX_ARG_FLAG_NUM = 12

func ZEND_SET_ARG_FLAG(zf *ZendFunction, arg_num uint32, mask __auto__) {
	zf.SetQuickArgFlags(zf.GetQuickArgFlags() | mask<<6<<arg_num*2)
}
func ZEND_CHECK_ARG_FLAG(zf *ZendFunction, arg_num int, mask __auto__) int {
	return zf.GetQuickArgFlags() >> (arg_num + 3) * 2 & mask
}
func QUICK_ARG_MUST_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF)
}
func QUICK_ARG_SHOULD_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func QUICK_ARG_MAY_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_PREFER_REF)
}

const ZEND_RETURN_VAL = 0
const ZEND_RETURN_REF = 1
const ZEND_BIND_VAL = 0
const ZEND_BIND_REF = 1
const ZEND_BIND_IMPLICIT = 2
const ZEND_RETURNS_FUNCTION uint32 = 1 << 0
const ZEND_RETURNS_VALUE uint32 = 1 << 1
const ZEND_ARRAY_ELEMENT_REF = 1 << 0
const ZEND_ARRAY_NOT_PACKED = 1 << 1
const ZEND_ARRAY_SIZE_SHIFT = 2

/* Attribute for ternary inside parentheses */

const ZEND_PARENTHESIZED_CONDITIONAL = 1

/* For "use" AST nodes and the seen symbol table */

const ZEND_SYMBOL_CLASS uint32 = 1 << 0
const ZEND_SYMBOL_FUNCTION uint32 = 1 << 1
const ZEND_SYMBOL_CONST uint32 = 1 << 2

/* All increment opcodes are even (decrement are odd) */

func ZEND_IS_INCREMENT(opcode ZendUchar) bool { return (opcode & 1) == 0 }
func ZEND_IS_BINARY_ASSIGN_OP_OPCODE(opcode __auto__) bool {
	return opcode >= ZEND_ADD && opcode <= ZEND_POW
}

/* Pseudo-opcodes that are used only temporarily during compilation */

const ZEND_PARENTHESIZED_CONCAT = 252
const ZEND_GOTO = 253
const ZEND_BRK = 254
const ZEND_CONT = 255
const ZEND_CLONE_FUNC_NAME = "__clone"
const ZEND_CONSTRUCTOR_FUNC_NAME = "__construct"
const ZEND_DESTRUCTOR_FUNC_NAME = "__destruct"
const ZEND_GET_FUNC_NAME = "__get"
const ZEND_SET_FUNC_NAME = "__set"
const ZEND_UNSET_FUNC_NAME = "__unset"
const ZEND_ISSET_FUNC_NAME = "__isset"
const ZEND_CALL_FUNC_NAME = "__call"
const ZEND_CALLSTATIC_FUNC_NAME = "__callstatic"
const ZEND_TOSTRING_FUNC_NAME = "__tostring"
const ZEND_AUTOLOAD_FUNC_NAME = "__autoload"
const ZEND_INVOKE_FUNC_NAME = "__invoke"
const ZEND_DEBUGINFO_FUNC_NAME = "__debuginfo"

/* The following constants may be combined in CG(compiler_options)
 * to change the default compiler behavior */

const ZEND_COMPILE_EXTENDED_STMT = 1 << 0
const ZEND_COMPILE_EXTENDED_FCALL = 1 << 1
const ZEND_COMPILE_EXTENDED_INFO = ZEND_COMPILE_EXTENDED_STMT | ZEND_COMPILE_EXTENDED_FCALL

/* call op_array handler of extendions */

const ZEND_COMPILE_HANDLE_OP_ARRAY = 1 << 2

/* generate ZEND_INIT_FCALL_BY_NAME for internal functions instead of ZEND_INIT_FCALL */

const ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS = 1 << 3

/* don't perform early binding for classes inherited form internal ones;
 * in namespaces assume that internal class that doesn't exist at compile-time
 * may apper in run-time */

const ZEND_COMPILE_IGNORE_INTERNAL_CLASSES = 1 << 4

/* generate ZEND_DECLARE_CLASS_DELAYED opcode to delay early binding */

const ZEND_COMPILE_DELAYED_BINDING = 1 << 5

/* disable constant substitution at compile-time */

const ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION = 1 << 6

/* disable usage of builtin instruction for strlen() */

const ZEND_COMPILE_NO_BUILTIN_STRLEN = 1 << 7

/* disable substitution of persistent constants at compile-time */

const ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION = 1 << 8

/* generate ZEND_INIT_FCALL_BY_NAME for userland functions instead of ZEND_INIT_FCALL */

const ZEND_COMPILE_IGNORE_USER_FUNCTIONS = 1 << 9

/* force ZEND_ACC_USE_GUARDS for all classes */

const ZEND_COMPILE_GUARDS = 1 << 10

/* disable builtin special case function calls */

const ZEND_COMPILE_NO_BUILTINS = 1 << 11

/* result of compilation may be stored in file cache */

const ZEND_COMPILE_WITH_FILE_CACHE = 1 << 12

/* ignore functions and classes declared in other files */

const ZEND_COMPILE_IGNORE_OTHER_FILES = 1 << 13

/* this flag is set when compiler invoked by opcache_compile_file() */

const ZEND_COMPILE_WITHOUT_EXECUTION = 1 << 14

/* this flag is set when compiler invoked during preloading */

const ZEND_COMPILE_PRELOAD = 1 << 15

/* disable jumptable optimization for switch statements */

const ZEND_COMPILE_NO_JUMPTABLES = 1 << 16

/* this flag is set when compiler invoked during preloading in separate process */

const ZEND_COMPILE_PRELOAD_IN_CHILD = 1 << 17

/* The default value for CG(compiler_options) */

const ZEND_COMPILE_DEFAULT = ZEND_COMPILE_HANDLE_OP_ARRAY

/* The default value for CG(compiler_options) during eval() */

const ZEND_COMPILE_DEFAULT_FOR_EVAL = 0

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

func FC(member __auto__) __auto__ {
	return CompilerGlobals.file_context.member
}

func ZendAllocCacheSlots(count unsigned) uint32 {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var ret uint32 = op_array.GetCacheSize()
	op_array.SetCacheSize(op_array.GetCacheSize() + count*b.SizeOf("void *"))
	return ret
}
func ZendAllocCacheSlot() uint32 { return ZendAllocCacheSlots(1) }

var CompilerGlobals ZendCompilerGlobals
var ExecutorGlobals ZendExecutorGlobals

func InitOp(op *ZendOp) {
	MAKE_NOP(op)
	op.SetExtendedValue(0)
	op.SetLineno(CompilerGlobals.GetZendLineno())
}
func GetNextOpNumber() uint32 {
	return CompilerGlobals.GetActiveOpArray().GetLast()
}
func GetNextOp() *ZendOp {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var next_op_num uint32 = b.PostInc(&(op_array.GetLast()))
	var next_op *ZendOp
	if UNEXPECTED(next_op_num >= CompilerGlobals.GetContext().GetOpcodesSize()) {
		CompilerGlobals.GetContext().SetOpcodesSize(CompilerGlobals.GetContext().GetOpcodesSize() * 4)
		op_array.SetOpcodes(Erealloc(op_array.GetOpcodes(), CompilerGlobals.GetContext().GetOpcodesSize()*b.SizeOf("zend_op")))
	}
	next_op = &op_array.GetOpcodes()[next_op_num]
	InitOp(next_op)
	return next_op
}
func GetNextBrkContElement() *ZendBrkContElement {
	CompilerGlobals.GetContext().GetLastBrkCont()++
	CompilerGlobals.GetContext().SetBrkContArray(Erealloc(CompilerGlobals.GetContext().GetBrkContArray(), b.SizeOf("zend_brk_cont_element")*CompilerGlobals.GetContext().GetLastBrkCont()))
	return &(CompilerGlobals.GetContext()).brk_cont_array[CompilerGlobals.GetContext().GetLastBrkCont()-1]
}
func ZendDestroyPropertyInfoInternal(zv *Zval) {
	var property_info *ZendPropertyInfo = Z_PTR_P(zv)
	ZendStringRelease(property_info.GetName())
	Free(property_info)
}

/* }}} */

func ZendBuildRuntimeDefinitionKey(name *ZendString, start_lineno uint32) *ZendString {
	var filename *ZendString = CompilerGlobals.GetActiveOpArray().GetFilename()
	var result *ZendString = ZendStrpprintf(0, "%c%s%s:%"+"u"+"$%"+PRIx32, '0', ZSTR_VAL(name), ZSTR_VAL(filename), start_lineno, b.PostInc(&(CompilerGlobals.GetRtdKeyCounter())))
	return ZendNewInternedString(result)
}

/* }}} */

func ZendGetUnqualifiedName(name *ZendString, result **byte, result_len *int) ZendBool {
	var ns_separator *byte = ZendMemrchr(ZSTR_VAL(name), '\\', ZSTR_LEN(name))
	if ns_separator != nil {
		*result = ns_separator + 1
		*result_len = ZSTR_VAL(name) + ZSTR_LEN(name) - (*result)
		return 1
	}
	return 0
}

/* }}} */

var ReservedClassNames []ReservedClassName = []ReservedClassName{{ZEND_STRL("bool")}, {ZEND_STRL("false")}, {ZEND_STRL("float")}, {ZEND_STRL("int")}, {ZEND_STRL("null")}, {ZEND_STRL("parent")}, {ZEND_STRL("self")}, {ZEND_STRL("static")}, {ZEND_STRL("string")}, {ZEND_STRL("true")}, {ZEND_STRL("void")}, {ZEND_STRL("iterable")}, {ZEND_STRL("object")}, {nil, 0}}

func ZendIsReservedClassName(name *ZendString) ZendBool {
	var reserved *ReservedClassName = ReservedClassNames
	var uqname *byte = ZSTR_VAL(name)
	var uqname_len int = ZSTR_LEN(name)
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
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as class name as it is reserved", ZSTR_VAL(name))
	}
}

/* }}} */

var BuiltinTypes []BuiltinTypeInfo = []BuiltinTypeInfo{{ZEND_STRL("int"), IS_LONG}, {ZEND_STRL("float"), IS_DOUBLE}, {ZEND_STRL("string"), IS_STRING}, {ZEND_STRL("bool"), _IS_BOOL}, {ZEND_STRL("void"), IS_VOID}, {ZEND_STRL("iterable"), IS_ITERABLE}, {ZEND_STRL("object"), IS_OBJECT}, {nil, 0, IS_UNDEF}}

func ZendLookupBuiltinTypeByName(name *ZendString) ZendUchar {
	var info *BuiltinTypeInfo = &BuiltinTypes[0]
	for ; info.GetName() != nil; info++ {
		if ZSTR_LEN(name) == info.GetNameLen() && ZendBinaryStrcasecmp(ZSTR_VAL(name), ZSTR_LEN(name), info.GetName(), info.GetNameLen()) == 0 {
			return info.GetType()
		}
	}
	return 0
}

/* }}} */

func ZendOparrayContextBegin(prev_context *ZendOparrayContext) {
	*prev_context = CompilerGlobals.GetContext()
	CompilerGlobals.GetContext().SetOpcodesSize(INITIAL_OP_ARRAY_SIZE)
	CompilerGlobals.GetContext().SetVarsSize(0)
	CompilerGlobals.GetContext().SetLiteralsSize(0)
	CompilerGlobals.GetContext().SetFastCallVar(-1)
	CompilerGlobals.GetContext().SetTryCatchOffset(-1)
	CompilerGlobals.GetContext().SetCurrentBrkCont(-1)
	CompilerGlobals.GetContext().SetLastBrkCont(0)
	CompilerGlobals.GetContext().SetBrkContArray(nil)
	CompilerGlobals.GetContext().SetLabels(nil)
}

/* }}} */

func ZendOparrayContextEnd(prev_context *ZendOparrayContext) {
	if CompilerGlobals.GetContext().GetBrkContArray() != nil {
		Efree(CompilerGlobals.GetContext().GetBrkContArray())
		CompilerGlobals.GetContext().SetBrkContArray(nil)
	}
	if CompilerGlobals.GetContext().GetLabels() != nil {
		ZendHashDestroy(CompilerGlobals.GetContext().GetLabels())
		FREE_HASHTABLE(CompilerGlobals.GetContext().GetLabels())
		CompilerGlobals.GetContext().SetLabels(nil)
	}
	CompilerGlobals.SetContext(*prev_context)
}

/* }}} */

func ZendResetImportTables() {
	if FC(imports) {
		ZendHashDestroy(FC(imports))
		Efree(FC(imports))
		FC(imports) = nil
	}
	if FC(imports_function) {
		ZendHashDestroy(FC(imports_function))
		Efree(FC(imports_function))
		FC(imports_function) = nil
	}
	if FC(imports_const) {
		ZendHashDestroy(FC(imports_const))
		Efree(FC(imports_const))
		FC(imports_const) = nil
	}
}

/* }}} */

func ZendEndNamespace() {
	FC(in_namespace) = 0
	ZendResetImportTables()
	if FC(current_namespace) {
		ZendStringReleaseEx(FC(current_namespace), 0)
		FC(current_namespace) = nil
	}
}

/* }}} */

func ZendFileContextBegin(prev_context *ZendFileContext) {
	*prev_context = CompilerGlobals.GetFileContext()
	FC(imports) = nil
	FC(imports_function) = nil
	FC(imports_const) = nil
	FC(current_namespace) = nil
	FC(in_namespace) = 0
	FC(has_bracketed_namespaces) = 0
	FC(declarables).ticks = 0
	ZendHashInit(&FC(seen_symbols), 8, nil, nil, 0)
}

/* }}} */

func ZendFileContextEnd(prev_context *ZendFileContext) {
	ZendEndNamespace()
	ZendHashDestroy(&FC(seen_symbols))
	CompilerGlobals.SetFileContext(*prev_context)
}

/* }}} */

func ZendInitCompilerDataStructures() {
	ZendStackInit(&(CompilerGlobals.GetLoopVarStack()), b.SizeOf("zend_loop_var"))
	ZendStackInit(&(CompilerGlobals.GetDelayedOplinesStack()), b.SizeOf("zend_op"))
	CompilerGlobals.SetActiveClassEntry(nil)
	CompilerGlobals.SetInCompilation(0)
	CompilerGlobals.SetSkipShebang(0)
	CompilerGlobals.SetEncodingDeclared(0)
	CompilerGlobals.SetMemoizedExprs(nil)
	CompilerGlobals.SetMemoizeMode(0)
}

/* }}} */

func ZendRegisterSeenSymbol(name *ZendString, kind uint32) {
	var zv *Zval = ZendHashFind(&FC(seen_symbols), name)
	if zv != nil {
		Z_LVAL_P(zv) |= kind
	} else {
		var tmp Zval
		ZVAL_LONG(&tmp, kind)
		ZendHashAddNew(&FC(seen_symbols), name, &tmp)
	}
}
func ZendHaveSeenSymbol(name *ZendString, kind uint32) ZendBool {
	var zv *Zval = ZendHashFind(&FC(seen_symbols), name)
	return zv != nil && (Z_LVAL_P(zv)&kind) != 0
}
func FileHandleDtor(fh *ZendFileHandle) { ZendFileHandleDtor(fh) }

/* }}} */

func InitCompiler() {
	CompilerGlobals.SetArena(ZendArenaCreate(64 * 1024))
	CompilerGlobals.SetActiveOpArray(nil)
	memset(&(CompilerGlobals.GetContext()), 0, b.SizeOf("CG ( context )"))
	ZendInitCompilerDataStructures()
	ZendInitRsrcList()
	ZendHashInit(&(CompilerGlobals.GetFilenamesTable()), 8, nil, ZVAL_PTR_DTOR, 0)
	ZendLlistInit(&(CompilerGlobals.GetOpenFiles()), b.SizeOf("zend_file_handle"), (func(any))(FileHandleDtor), 0)
	CompilerGlobals.SetUncleanShutdown(0)
	CompilerGlobals.SetDelayedVarianceObligations(nil)
	CompilerGlobals.SetDelayedAutoloads(nil)
}

/* }}} */

func ShutdownCompiler() {
	ZendStackDestroy(&(CompilerGlobals.GetLoopVarStack()))
	ZendStackDestroy(&(CompilerGlobals.GetDelayedOplinesStack()))
	ZendHashDestroy(&(CompilerGlobals.GetFilenamesTable()))
	ZendArenaDestroy(CompilerGlobals.GetArena())
	if CompilerGlobals.GetDelayedVarianceObligations() != nil {
		ZendHashDestroy(CompilerGlobals.GetDelayedVarianceObligations())
		FREE_HASHTABLE(CompilerGlobals.GetDelayedVarianceObligations())
		CompilerGlobals.SetDelayedVarianceObligations(nil)
	}
	if CompilerGlobals.GetDelayedAutoloads() != nil {
		ZendHashDestroy(CompilerGlobals.GetDelayedAutoloads())
		FREE_HASHTABLE(CompilerGlobals.GetDelayedAutoloads())
		CompilerGlobals.SetDelayedAutoloads(nil)
	}
}

/* }}} */

func ZendSetCompiledFilename(new_compiled_filename *ZendString) *ZendString {
	var p *Zval
	var rv Zval
	if b.Assign(&p, ZendHashFind(&(CompilerGlobals.GetFilenamesTable()), new_compiled_filename)) {
		ZEND_ASSERT(Z_TYPE_P(p) == IS_STRING)
		CompilerGlobals.SetCompiledFilename(Z_STR_P(p))
		return Z_STR_P(p)
	}
	new_compiled_filename = ZendNewInternedString(ZendStringCopy(new_compiled_filename))
	ZVAL_STR(&rv, new_compiled_filename)
	ZendHashAddNew(&(CompilerGlobals.GetFilenamesTable()), new_compiled_filename, &rv)
	CompilerGlobals.SetCompiledFilename(new_compiled_filename)
	return new_compiled_filename
}

/* }}} */

func ZendRestoreCompiledFilename(original_compiled_filename *ZendString) {
	CompilerGlobals.SetCompiledFilename(original_compiled_filename)
}

/* }}} */

func ZendGetCompiledFilename() *ZendString {
	return CompilerGlobals.GetCompiledFilename()
}

/* }}} */

func ZendGetCompiledLineno() int { return CompilerGlobals.GetZendLineno() }

/* }}} */

func ZendIsCompiling() ZendBool {
	return CompilerGlobals.GetInCompilation()
}

/* }}} */

func GetTemporaryVariable() uint32 {
	uint32_t(CompilerGlobals.GetActiveOpArray()).T++
	return uint32_t(CompilerGlobals.GetActiveOpArray()).T - 1
}

/* }}} */

func LookupCv(name *ZendString) int {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var i int = 0
	var hash_value ZendUlong = ZendStringHashVal(name)
	for i < op_array.GetLastVar() {
		if ZSTR_H(op_array.GetVars()[i]) == hash_value && ZendStringEquals(op_array.GetVars()[i], name) != 0 {
			return int(ZendIntptrT(ZEND_CALL_VAR_NUM(nil, i)))
		}
		i++
	}
	i = op_array.GetLastVar()
	op_array.GetLastVar()++
	if op_array.GetLastVar() > CompilerGlobals.GetContext().GetVarsSize() {
		CompilerGlobals.GetContext().SetVarsSize(CompilerGlobals.GetContext().GetVarsSize() + 16)
		op_array.SetVars(Erealloc(op_array.GetVars(), CompilerGlobals.GetContext().GetVarsSize()*b.SizeOf("zend_string *")))
	}
	op_array.GetVars()[i] = ZendStringCopy(name)
	return int(ZendIntptrT(ZEND_CALL_VAR_NUM(nil, i)))
}

/* }}} */

func ZendDelLiteral(op_array *ZendOpArray, n int) {
	ZvalPtrDtorNogc(CT_CONSTANT_EX(op_array, n))
	if n+1 == op_array.GetLastLiteral() {
		op_array.GetLastLiteral()--
	} else {
		ZVAL_UNDEF(CT_CONSTANT_EX(op_array, n))
	}
}

/* }}} */

/* Common part of zend_add_literal and zend_append_individual_literal */

func ZendInsertLiteral(op_array *ZendOpArray, zv *Zval, literal_position int) {
	var lit *Zval = CT_CONSTANT_EX(op_array, literal_position)
	if Z_TYPE_P(zv) == IS_STRING {
		ZvalMakeInternedString(zv)
	}
	ZVAL_COPY_VALUE(lit, zv)
	Z_EXTRA_P(lit) = 0
}

/* }}} */

func ZendAddLiteral(zv *Zval) int {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var i int = op_array.GetLastLiteral()
	op_array.GetLastLiteral()++
	if i >= CompilerGlobals.GetContext().GetLiteralsSize() {
		for i >= CompilerGlobals.GetContext().GetLiteralsSize() {
			CompilerGlobals.GetContext().SetLiteralsSize(CompilerGlobals.GetContext().GetLiteralsSize() + 16)
		}
		op_array.SetLiterals((*Zval)(Erealloc(op_array.GetLiterals(), CompilerGlobals.GetContext().GetLiteralsSize()*b.SizeOf("zval"))))
	}
	ZendInsertLiteral(op_array, zv, i)
	return i
}

/* }}} */

func ZendAddLiteralString(str **ZendString) int {
	var ret int
	var zv Zval
	ZVAL_STR(&zv, *str)
	ret = ZendAddLiteral(&zv)
	*str = Z_STR(zv)
	return ret
}

/* }}} */

func ZendAddFuncNameLiteral(name *ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *ZendString = ZendStringTolower(name)
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

	var lc_name *ZendString = ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)

	/* Lowercased unqualfied name */

	if ZendGetUnqualifiedName(name, &unqualified_name, &unqualified_name_len) != 0 {
		lc_name = ZendStringAlloc(unqualified_name_len, 0)
		ZendStrTolowerCopy(ZSTR_VAL(lc_name), unqualified_name, unqualified_name_len)
		ZendAddLiteralString(&lc_name)
	}
	return ret
}

/* }}} */

func ZendAddClassNameLiteral(name *ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *ZendString = ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)
	return ret
}

/* }}} */

func ZendAddConstNameLiteral(name *ZendString, unqualified ZendBool) int {
	var tmp_name *ZendString
	var ret int = ZendAddLiteralString(&name)
	var ns_len int = 0
	var after_ns_len int = ZSTR_LEN(name)
	var after_ns *byte = ZendMemrchr(ZSTR_VAL(name), '\\', ZSTR_LEN(name))
	if after_ns != nil {
		after_ns += 1
		ns_len = after_ns - ZSTR_VAL(name) - 1
		after_ns_len = ZSTR_LEN(name) - ns_len - 1

		/* lowercased namespace name & original constant name */

		tmp_name = ZendStringInit(ZSTR_VAL(name), ZSTR_LEN(name), 0)
		ZendStrTolower(ZSTR_VAL(tmp_name), ns_len)
		ZendAddLiteralString(&tmp_name)

		/* lowercased namespace name & lowercased constant name */

		tmp_name = ZendStringTolower(name)
		ZendAddLiteralString(&tmp_name)
		if unqualified == 0 {
			return ret
		}
	} else {
		after_ns = ZSTR_VAL(name)
	}

	/* original unqualified constant name */

	tmp_name = ZendStringInit(after_ns, after_ns_len, 0)
	ZendAddLiteralString(&tmp_name)

	/* lowercased unqualified constant name */

	tmp_name = ZendStringAlloc(after_ns_len, 0)
	ZendStrTolowerCopy(ZSTR_VAL(tmp_name), after_ns, after_ns_len)
	ZendAddLiteralString(&tmp_name)
	return ret
}

/* }}} */

func LITERAL_STR(op ZnodeOp, str *ZendString) {
	var _c Zval
	ZVAL_STR(&_c, str)
	op.SetConstant(ZendAddLiteral(&_c))
}
func ZendStopLexing() {
	if LanguageScannerGlobals.GetOnEvent() != nil {
		LanguageScannerGlobals.GetOnEvent()(ON_STOP, END, 0, LanguageScannerGlobals.GetOnEventContext())
	}
	LanguageScannerGlobals.SetYyCursor(LanguageScannerGlobals.GetYyLimit())
}
func ZendBeginLoop(free_opcode ZendUchar, loop_var *Znode, is_switch ZendBool) {
	var brk_cont_element *ZendBrkContElement
	var parent int = CompilerGlobals.GetContext().GetCurrentBrkCont()
	var info ZendLoopVar = ZendLoopVar{0}
	CompilerGlobals.GetContext().SetCurrentBrkCont(CompilerGlobals.GetContext().GetLastBrkCont())
	brk_cont_element = GetNextBrkContElement()
	brk_cont_element.SetParent(parent)
	brk_cont_element.SetIsSwitch(is_switch)
	if loop_var != nil && (loop_var.GetOpType()&(IS_VAR|IS_TMP_VAR)) != 0 {
		var start uint32 = GetNextOpNumber()
		info.SetOpcode(free_opcode)
		info.SetVarType(loop_var.GetOpType())
		info.SetVarNum(loop_var.GetOp().GetVar())
		brk_cont_element.SetStart(start)
	} else {
		info.SetOpcode(ZEND_NOP)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

		brk_cont_element.SetStart(-1)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

	}
	ZendStackPush(&(CompilerGlobals.GetLoopVarStack()), &info)
}

/* }}} */

func ZendEndLoop(cont_addr int, var_node *Znode) {
	var end uint32 = GetNextOpNumber()
	var brk_cont_element *ZendBrkContElement = &(CompilerGlobals.GetContext()).brk_cont_array[CompilerGlobals.GetContext().GetCurrentBrkCont()]
	brk_cont_element.SetCont(cont_addr)
	brk_cont_element.SetBrk(end)
	CompilerGlobals.GetContext().SetCurrentBrkCont(brk_cont_element.GetParent())
	ZendStackDelTop(&(CompilerGlobals.GetLoopVarStack()))
}

/* }}} */

func ZendDoFree(op1 *Znode) {
	if op1.GetOpType() == IS_TMP_VAR {
		var opline *ZendOp = &(CompilerGlobals.GetActiveOpArray()).opcodes[CompilerGlobals.GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == ZEND_END_SILENCE {
			opline--
		}
		if opline.GetResultType() == IS_TMP_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == ZEND_BOOL || opline.GetOpcode() == ZEND_BOOL_NOT {
				return
			}
		}
		ZendEmitOp(nil, ZEND_FREE, op1, nil)
	} else if op1.GetOpType() == IS_VAR {
		var opline *ZendOp = &(CompilerGlobals.GetActiveOpArray()).opcodes[CompilerGlobals.GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == ZEND_END_SILENCE || opline.GetOpcode() == ZEND_EXT_FCALL_END || opline.GetOpcode() == ZEND_OP_DATA {
			opline--
		}
		if opline.GetResultType() == IS_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == ZEND_FETCH_THIS {
				opline.SetOpcode(ZEND_NOP)
				opline.SetResultType(IS_UNUSED)
			} else {
				opline.SetResultType(IS_UNUSED)
			}
		} else {
			for opline >= CompilerGlobals.GetActiveOpArray().GetOpcodes() {
				if (opline.GetOpcode() == ZEND_FETCH_LIST_R || opline.GetOpcode() == ZEND_FETCH_LIST_W) && opline.GetOp1Type() == IS_VAR && opline.GetOp1().GetVar() == op1.GetOp().GetVar() {
					ZendEmitOp(nil, ZEND_FREE, op1, nil)
					return
				}
				if opline.GetResultType() == IS_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
					if opline.GetOpcode() == ZEND_NEW {
						ZendEmitOp(nil, ZEND_FREE, op1, nil)
					}
					break
				}
				opline--
			}
		}
	} else if op1.GetOpType() == IS_CONST {

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
	if (flags&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 && (new_flag&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_FINAL) != 0 && (new_flag&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 && (new_flags&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class", 0)
		return 0
	}
	return new_flags
}

/* }}} */

func ZendAddMemberModifier(flags uint32, new_flag uint32) uint32 {
	var new_flags uint32 = flags | new_flag
	if (flags&ZEND_ACC_PPP_MASK) != 0 && (new_flag&ZEND_ACC_PPP_MASK) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple access type modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_ABSTRACT) != 0 && (new_flag&ZEND_ACC_ABSTRACT) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_STATIC) != 0 && (new_flag&ZEND_ACC_STATIC) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple static modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_FINAL) != 0 && (new_flag&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&ZEND_ACC_ABSTRACT) != 0 && (new_flags&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class member", 0)
		return 0
	}
	return new_flags
}

/* }}} */

func ZendConcat3(str1 *byte, str1_len int, str2 string, str2_len int, str3 *byte, str3_len int) *ZendString {
	var len_ int = str1_len + str2_len + str3_len
	var res *ZendString = ZendStringAlloc(len_, 0)
	memcpy(ZSTR_VAL(res), str1, str1_len)
	memcpy(ZSTR_VAL(res)+str1_len, str2, str2_len)
	memcpy(ZSTR_VAL(res)+str1_len+str2_len, str3, str3_len)
	ZSTR_VAL(res)[len_] = '0'
	return res
}
func ZendConcatNames(name1 *byte, name1_len int, name2 *byte, name2_len int) *ZendString {
	return ZendConcat3(name1, name1_len, "\\", 1, name2, name2_len)
}
func ZendPrefixWithNs(name *ZendString) *ZendString {
	if FC(current_namespace) {
		var ns *ZendString = FC(current_namespace)
		return ZendConcatNames(ZSTR_VAL(ns), ZSTR_LEN(ns), ZSTR_VAL(name), ZSTR_LEN(name))
	} else {
		return ZendStringCopy(name)
	}
}
func ZendHashFindPtrLc(ht *HashTable, str *byte, len_ int) any {
	var result any
	var lcname *ZendString
	ZSTR_ALLOCA_ALLOC(lcname, len_, use_heap)
	ZendStrTolowerCopy(ZSTR_VAL(lcname), str, len_)
	result = ZendHashFindPtr(ht, lcname)
	ZSTR_ALLOCA_FREE(lcname, use_heap)
	return result
}
func ZendResolveNonClassName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool, case_sensitive ZendBool, current_import_sub *HashTable) *ZendString {
	var compound *byte
	*is_fully_qualified = 0
	if ZSTR_VAL(name)[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		*is_fully_qualified = 1
		return ZendStringInit(ZSTR_VAL(name)+1, ZSTR_LEN(name)-1, 0)
	}
	if type_ == ZEND_NAME_FQ {
		*is_fully_qualified = 1
		return ZendStringCopy(name)
	}
	if type_ == ZEND_NAME_RELATIVE {
		*is_fully_qualified = 1
		return ZendPrefixWithNs(name)
	}
	if current_import_sub != nil {

		/* If an unqualified name is a function/const alias, replace it. */

		var import_name *ZendString
		if case_sensitive != 0 {
			import_name = ZendHashFindPtr(current_import_sub, name)
		} else {
			import_name = ZendHashFindPtrLc(current_import_sub, ZSTR_VAL(name), ZSTR_LEN(name))
		}
		if import_name != nil {
			*is_fully_qualified = 1
			return ZendStringCopy(import_name)
		}
	}
	compound = memchr(ZSTR_VAL(name), '\\', ZSTR_LEN(name))
	if compound != nil {
		*is_fully_qualified = 1
	}
	if compound != nil && FC(imports) {

		/* If the first part of a qualified name is an alias, substitute it. */

		var len_ int = compound - ZSTR_VAL(name)
		var import_name *ZendString = ZendHashFindPtrLc(FC(imports), ZSTR_VAL(name), len_)
		if import_name != nil {
			return ZendConcatNames(ZSTR_VAL(import_name), ZSTR_LEN(import_name), ZSTR_VAL(name)+len_+1, ZSTR_LEN(name)-len_-1)
		}
	}
	return ZendPrefixWithNs(name)
}

/* }}} */

func ZendResolveFunctionName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool) *ZendString {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 0, FC(imports_function))
}

/* }}} */

func ZendResolveConstName(name *ZendString, type_ uint32, is_fully_qualified *ZendBool) *ZendString {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 1, FC(imports_const))
}

/* }}} */

func ZendResolveClassName(name *ZendString, type_ uint32) *ZendString {
	var compound *byte
	if type_ == ZEND_NAME_RELATIVE {
		return ZendPrefixWithNs(name)
	}
	if type_ == ZEND_NAME_FQ || ZSTR_VAL(name)[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		if ZSTR_VAL(name)[0] == '\\' {
			name = ZendStringInit(ZSTR_VAL(name)+1, ZSTR_LEN(name)-1, 0)
		} else {
			ZendStringAddref(name)
		}

		/* Ensure that \self, \parent and \static are not used */

		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "'\\%s' is an invalid class name", ZSTR_VAL(name))
		}
		return name
	}
	if FC(imports) {
		compound = memchr(ZSTR_VAL(name), '\\', ZSTR_LEN(name))
		if compound != nil {

			/* If the first part of a qualified name is an alias, substitute it. */

			var len_ int = compound - ZSTR_VAL(name)
			var import_name *ZendString = ZendHashFindPtrLc(FC(imports), ZSTR_VAL(name), len_)
			if import_name != nil {
				return ZendConcatNames(ZSTR_VAL(import_name), ZSTR_LEN(import_name), ZSTR_VAL(name)+len_+1, ZSTR_LEN(name)-len_-1)
			}
		} else {

			/* If an unqualified name is an alias, replace it. */

			var import_name *ZendString = ZendHashFindPtrLc(FC(imports), ZSTR_VAL(name), ZSTR_LEN(name))
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
	if Z_TYPE_P(class_name) != IS_STRING {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Illegal class name")
	}
	return ZendResolveClassName(Z_STR_P(class_name), ast.GetAttr())
}

/* }}} */

func LabelPtrDtor(zv *Zval) {
	EfreeSize(Z_PTR_P(zv), b.SizeOf("zend_label"))
}

/* }}} */

func StrDtor(zv *Zval) { ZendStringReleaseEx(Z_STR_P(zv), 0) }

/* }}} */

func ZendAddTryElement(try_op uint32) uint32 {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var try_catch_offset uint32 = b.PostInc(&(op_array.GetLastTryCatch()))
	var elem *ZendTryCatchElement
	op_array.SetTryCatchArray(SafeErealloc(op_array.GetTryCatchArray(), b.SizeOf("zend_try_catch_element"), op_array.GetLastTryCatch(), 0))
	elem = &op_array.try_catch_array[try_catch_offset]
	elem.SetTryOp(try_op)
	elem.SetCatchOp(0)
	elem.SetFinallyOp(0)
	elem.SetFinallyEnd(0)
	return try_catch_offset
}

/* }}} */

func FunctionAddRef(function *ZendFunction) {
	if function.GetType() == ZEND_USER_FUNCTION {
		var op_array *ZendOpArray = &function.op_array
		if op_array.GetRefcount() != nil {
			(*op_array).refcount++
		}
		if op_array.GetStaticVariables() != nil {
			if (GC_FLAGS(op_array.GetStaticVariables()) & IS_ARRAY_IMMUTABLE) == 0 {
				GC_ADDREF(op_array.GetStaticVariables())
			}
		}
		if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
			ZEND_ASSERT((op_array.GetFnFlags() & ZEND_ACC_PRELOADED) != 0)
			ZEND_MAP_PTR_NEW(op_array.run_time_cache)
			ZEND_MAP_PTR_NEW(op_array.static_variables_ptr)
		} else {
			ZEND_MAP_PTR_INIT(op_array.static_variables_ptr, &op_array.static_variables)
			ZEND_MAP_PTR_INIT(op_array.run_time_cache, ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("void *")))
			ZEND_MAP_PTR_SET(op_array.run_time_cache, nil)
		}
	} else if function.GetType() == ZEND_INTERNAL_FUNCTION {
		if function.GetFunctionName() != nil {
			ZendStringAddref(function.GetFunctionName())
		}
	}
}

/* }}} */

func DoBindFunctionError(lcname *ZendString, op_array *ZendOpArray, compile_time ZendBool) {
	var zv *Zval = ZendHashFindEx(b.CondF(compile_time != 0, func() *HashTable { return CompilerGlobals.GetFunctionTable() }, func() *HashTable { return ExecutorGlobals.GetFunctionTable() }), lcname, 1)
	var error_level int = b.Cond(compile_time != 0, E_COMPILE_ERROR, E_ERROR)
	var old_function *ZendFunction
	ZEND_ASSERT(zv != nil)
	old_function = (*ZendFunction)(Z_PTR_P(zv))
	if old_function.GetType() == ZEND_USER_FUNCTION && old_function.GetOpArray().GetLast() > 0 {
		ZendErrorNoreturn(error_level, "Cannot redeclare %s() (previously declared in %s:%d)", b.CondF(op_array != nil, func() []byte { return ZSTR_VAL(op_array.GetFunctionName()) }, func() []byte { return ZSTR_VAL(old_function.GetFunctionName()) }), ZSTR_VAL(old_function.GetOpArray().GetFilename()), old_function.GetOpArray().GetOpcodes()[0].GetLineno())
	} else {
		ZendErrorNoreturn(error_level, "Cannot redeclare %s()", b.CondF(op_array != nil, func() []byte { return ZSTR_VAL(op_array.GetFunctionName()) }, func() []byte { return ZSTR_VAL(old_function.GetFunctionName()) }))
	}
}
func DoBindFunction(lcname *Zval) int {
	var function *ZendFunction
	var rtd_key *Zval
	var zv *Zval
	rtd_key = lcname + 1
	zv = ZendHashFindEx(ExecutorGlobals.GetFunctionTable(), Z_STR_P(rtd_key), 1)
	if UNEXPECTED(zv == nil) {
		DoBindFunctionError(Z_STR_P(lcname), nil, 0)
		return FAILURE
	}
	function = (*ZendFunction)(Z_PTR_P(zv))
	if UNEXPECTED((function.GetFnFlags()&ZEND_ACC_PRELOADED) != 0) && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_PRELOAD) == 0 {
		zv = ZendHashAdd(ExecutorGlobals.GetFunctionTable(), Z_STR_P(lcname), zv)
	} else {
		zv = ZendHashSetBucketKey(ExecutorGlobals.GetFunctionTable(), (*Bucket)(zv), Z_STR_P(lcname))
	}
	if UNEXPECTED(zv == nil) {
		DoBindFunctionError(Z_STR_P(lcname), &function.op_array, 0)
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
	zv = ZendHashFindEx(ExecutorGlobals.GetClassTable(), Z_STR_P(rtd_key), 1)
	if UNEXPECTED(zv == nil) {
		ce = ZendHashFindPtr(ExecutorGlobals.GetClassTable(), Z_STR_P(lcname))
		if ce != nil {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ZSTR_VAL(ce.GetName()))
			return FAILURE
		} else {
			for {
				ZEND_ASSERT((ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetFnFlags() & ZEND_ACC_PRELOADED) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetFilename()) == SUCCESS {
					zv = ZendHashFindEx(ExecutorGlobals.GetClassTable(), Z_STR_P(rtd_key), 1)
					if EXPECTED(zv != nil) {
						break
					}
				}
				ZendErrorNoreturn(E_ERROR, "Class %s wasn't preloaded", Z_STRVAL_P(lcname))
				return FAILURE
				break
			}
		}
	}

	/* Register the derived class */

	ce = (*ZendClassEntry)(Z_PTR_P(zv))
	zv = ZendHashSetBucketKey(ExecutorGlobals.GetClassTable(), (*Bucket)(zv), Z_STR_P(lcname))
	if UNEXPECTED(zv == nil) {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ZSTR_VAL(ce.GetName()))
		return FAILURE
	}
	if ZendDoLinkClass(ce, lc_parent_name) == FAILURE {

		/* Reload bucket pointer, the hash table may have been reallocated */

		zv = ZendHashFind(ExecutorGlobals.GetClassTable(), Z_STR_P(lcname))
		ZendHashSetBucketKey(ExecutorGlobals.GetClassTable(), (*Bucket)(zv), Z_STR_P(rtd_key))
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func ZendMarkFunctionAsGenerator() {
	if CompilerGlobals.GetActiveOpArray().GetFunctionName() == nil {
		ZendErrorNoreturn(E_COMPILE_ERROR, "The \"yield\" expression can only be used inside a function")
	}
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {
		var return_info ZendArgInfo = CompilerGlobals.GetActiveOpArray().GetArgInfo()[-1]
		if ZEND_TYPE_CODE(return_info.GetType()) != IS_ITERABLE {
			var msg *byte = "Generators may only declare a return type of Generator, Iterator, Traversable, or iterable, %s is not permitted"
			if !(ZEND_TYPE_IS_CLASS(return_info.GetType())) {
				ZendErrorNoreturn(E_COMPILE_ERROR, msg, ZendGetTypeByConst(ZEND_TYPE_CODE(return_info.GetType())))
			}
			if !(ZendStringEqualsLiteralCi(ZEND_TYPE_NAME(return_info.GetType()), "Traversable")) && !(ZendStringEqualsLiteralCi(ZEND_TYPE_NAME(return_info.GetType()), "Iterator")) && !(ZendStringEqualsLiteralCi(ZEND_TYPE_NAME(return_info.GetType()), "Generator")) {
				ZendErrorNoreturn(E_COMPILE_ERROR, msg, ZSTR_VAL(ZEND_TYPE_NAME(return_info.GetType())))
			}
		}
	}
	CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_GENERATOR)
}

/* }}} */

func ZendBuildDelayedEarlyBindingList(op_array *ZendOpArray) uint32 {
	if (op_array.GetFnFlags() & ZEND_ACC_EARLY_BINDING) != 0 {
		var first_early_binding_opline uint32 = uint32_t - 1
		var prev_opline_num *uint32 = &first_early_binding_opline
		var opline *ZendOp = op_array.GetOpcodes()
		var end *ZendOp = opline + op_array.GetLast()
		for opline < end {
			if opline.GetOpcode() == ZEND_DECLARE_CLASS_DELAYED {
				*prev_opline_num = opline - op_array.GetOpcodes()
				prev_opline_num = &opline.result.GetOplineNum()
			}
			opline++
		}
		*prev_opline_num = -1
		return first_early_binding_opline
	}
	return uint32_t - 1
}

/* }}} */

func ZendDoDelayedEarlyBinding(op_array *ZendOpArray, first_early_binding_opline uint32) {
	if first_early_binding_opline != uint32_t-1 {
		var orig_in_compilation ZendBool = CompilerGlobals.GetInCompilation()
		var opline_num uint32 = first_early_binding_opline
		var run_time_cache *any
		if op_array.GetRunTimeCachePtr() == nil {
			var ptr any
			ZEND_ASSERT((op_array.GetFnFlags() & ZEND_ACC_HEAP_RT_CACHE) != 0)
			ptr = Emalloc(op_array.GetCacheSize() + b.SizeOf("void *"))
			ZEND_MAP_PTR_INIT(op_array.run_time_cache, ptr)
			ptr = (*byte)(ptr + b.SizeOf("void *"))
			ZEND_MAP_PTR_SET(op_array.run_time_cache, ptr)
			memset(ptr, 0, op_array.GetCacheSize())
		}
		run_time_cache = RUN_TIME_CACHE(op_array)
		CompilerGlobals.SetInCompilation(1)
		for opline_num != uint32_t-1 {
			var opline *ZendOp = &op_array.opcodes[opline_num]
			var lcname *Zval = RT_CONSTANT(opline, opline.GetOp1())
			var zv *Zval = ZendHashFindEx(ExecutorGlobals.GetClassTable(), Z_STR_P(lcname+1), 1)
			if zv != nil {
				var ce *ZendClassEntry = Z_CE_P(zv)
				var lc_parent_name *ZendString = Z_STR_P(RT_CONSTANT(opline, opline.GetOp2()))
				var parent_ce *ZendClassEntry = ZendHashFindExPtr(ExecutorGlobals.GetClassTable(), lc_parent_name, 1)
				if parent_ce != nil {
					if ZendTryEarlyBind(ce, parent_ce, Z_STR_P(lcname), zv) != 0 {

						/* Store in run-time cache */

						(*any)((*byte)(run_time_cache + opline.GetExtendedValue()))[0] = ce

						/* Store in run-time cache */

					}
				}
			}
			opline_num = op_array.GetOpcodes()[opline_num].GetResult().GetOplineNum()
		}
		CompilerGlobals.SetInCompilation(orig_in_compilation)
	}
}

/* }}} */

func ZendManglePropertyName(src1 *byte, src1_length int, src2 string, src2_length int, internal int) *ZendString {
	var prop_name_length int = 1 + src1_length + 1 + src2_length
	var prop_name *ZendString = ZendStringAlloc(prop_name_length, internal)
	ZSTR_VAL(prop_name)[0] = '0'
	memcpy(ZSTR_VAL(prop_name)+1, src1, src1_length+1)
	memcpy(ZSTR_VAL(prop_name)+1+src1_length+1, src2, src2_length+1)
	return prop_name
}

/* }}} */

func ZendStrnlen(s *byte, maxlen int) int {
	var len_ int = 0
	for b.PostInc(&(*s)) && b.PostDec(&maxlen) {
		len_++
	}
	return len_
}

/* }}} */

func ZendUnmanglePropertyNameEx(name *ZendString, class_name **byte, prop_name **byte, prop_len *int) int {
	var class_name_len int
	var anonclass_src_len int
	*class_name = nil
	if ZSTR_LEN(name) == 0 || ZSTR_VAL(name)[0] != '0' {
		*prop_name = ZSTR_VAL(name)
		if prop_len != nil {
			*prop_len = ZSTR_LEN(name)
		}
		return SUCCESS
	}
	if ZSTR_LEN(name) < 3 || ZSTR_VAL(name)[1] == '0' {
		ZendError(E_NOTICE, "Illegal member variable name")
		*prop_name = ZSTR_VAL(name)
		if prop_len != nil {
			*prop_len = ZSTR_LEN(name)
		}
		return FAILURE
	}
	class_name_len = ZendStrnlen(ZSTR_VAL(name)+1, ZSTR_LEN(name)-2)
	if class_name_len >= ZSTR_LEN(name)-2 || ZSTR_VAL(name)[class_name_len+1] != '0' {
		ZendError(E_NOTICE, "Corrupt member variable name")
		*prop_name = ZSTR_VAL(name)
		if prop_len != nil {
			*prop_len = ZSTR_LEN(name)
		}
		return FAILURE
	}
	*class_name = ZSTR_VAL(name) + 1
	anonclass_src_len = ZendStrnlen((*class_name)+class_name_len+1, ZSTR_LEN(name)-class_name_len-2)
	if class_name_len+anonclass_src_len+2 != ZSTR_LEN(name) {
		class_name_len += anonclass_src_len + 1
	}
	*prop_name = ZSTR_VAL(name) + class_name_len + 2
	if prop_len != nil {
		*prop_len = ZSTR_LEN(name) - class_name_len - 2
	}
	return SUCCESS
}

/* }}} */

func ZendLookupReservedConst(name *byte, len_ int) *ZendConstant {
	var c *ZendConstant = ZendHashFindPtrLc(ExecutorGlobals.GetZendConstants(), name, len_)
	if c != nil && (ZEND_CONSTANT_FLAGS(c)&CONST_CS) == 0 && (ZEND_CONSTANT_FLAGS(c)&CONST_CT_SUBST) != 0 {
		return c
	}
	return nil
}

/* }}} */

func ZendTryCtEvalConst(zv *Zval, name *ZendString, is_fully_qualified ZendBool) ZendBool {
	var c *ZendConstant

	/* Substitute case-sensitive (or lowercase) constants */

	c = ZendHashFindPtr(ExecutorGlobals.GetZendConstants(), name)
	if c != nil && ((ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT) != 0 && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION) == 0 && ((ZEND_CONSTANT_FLAGS(c)&CONST_NO_FILE_CACHE) == 0 || (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_WITH_FILE_CACHE) == 0) || Z_TYPE(c.GetValue()) < IS_OBJECT && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0) {
		ZVAL_COPY_OR_DUP(zv, &c.value)
		return 1
	}

	/* Substitute true, false and null (including unqualified usage in namespaces) */

	var lookup_name *byte = ZSTR_VAL(name)
	var lookup_len int = ZSTR_LEN(name)
	if is_fully_qualified == 0 {
		ZendGetUnqualifiedName(name, &lookup_name, &lookup_len)
	}
	c = ZendLookupReservedConst(lookup_name, lookup_len)
	if c != nil {
		ZVAL_COPY_OR_DUP(zv, &c.value)
		return 1
	}
	return 0
}

/* }}} */

func ZendIsScopeKnown() ZendBool {
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_CLOSURE) != 0 {

		/* Closures can be rebound to a different scope */

		return 0

		/* Closures can be rebound to a different scope */

	}
	if CompilerGlobals.GetActiveClassEntry() == nil {

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

		return CompilerGlobals.GetActiveOpArray().GetFunctionName() != nil

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

	}

	/* For traits self etc refers to the using class, not the trait itself */

	return (CompilerGlobals.GetActiveClassEntry().GetCeFlags() & ZEND_ACC_TRAIT) == 0

	/* For traits self etc refers to the using class, not the trait itself */
}

/* }}} */

func ClassNameRefersToActiveCe(class_name *ZendString, fetch_type uint32) ZendBool {
	if CompilerGlobals.GetActiveClassEntry() == nil {
		return 0
	}
	if fetch_type == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() != 0 {
		return 1
	}
	return fetch_type == ZEND_FETCH_CLASS_DEFAULT && ZendStringEqualsCi(class_name, CompilerGlobals.GetActiveClassEntry().GetName())
}

/* }}} */

func ZendGetClassFetchType(name *ZendString) uint32 {
	if ZendStringEqualsLiteralCi(name, "self") {
		return ZEND_FETCH_CLASS_SELF
	} else if ZendStringEqualsLiteralCi(name, "parent") {
		return ZEND_FETCH_CLASS_PARENT
	} else if ZendStringEqualsLiteralCi(name, "static") {
		return ZEND_FETCH_CLASS_STATIC
	} else {
		return ZEND_FETCH_CLASS_DEFAULT
	}
}

/* }}} */

func ZendGetClassFetchTypeAst(name_ast *ZendAst) uint32 {
	/* Fully qualified names are always default refs */

	if name_ast.GetAttr() == ZEND_NAME_FQ {
		return ZEND_FETCH_CLASS_DEFAULT
	}
	return ZendGetClassFetchType(ZendAstGetStr(name_ast))
}

/* }}} */

func ZendEnsureValidClassFetchType(fetch_type uint32) {
	if fetch_type != ZEND_FETCH_CLASS_DEFAULT && ZendIsScopeKnown() != 0 {
		var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
		if ce == nil {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		} else if fetch_type == ZEND_FETCH_CLASS_PARENT && !(ce.parent_name) {
			ZendError(E_DEPRECATED, "Cannot use \"parent\" when current class scope has no parent")
		}
	}
}

/* }}} */

func ZendTryCompileConstExprResolveClassName(zv *Zval, class_ast *ZendAst) ZendBool {
	var fetch_type uint32
	var class_name *Zval
	if class_ast.GetKind() != ZEND_AST_ZVAL {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use ::class with dynamic class name")
	}
	class_name = ZendAstGetZval(class_ast)
	if Z_TYPE_P(class_name) != IS_STRING {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Illegal class name")
	}
	fetch_type = ZendGetClassFetchType(Z_STR_P(class_name))
	ZendEnsureValidClassFetchType(fetch_type)
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		if CompilerGlobals.GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
			ZVAL_STR_COPY(zv, CompilerGlobals.GetActiveClassEntry().GetName())
			return 1
		}
		return 0
	case ZEND_FETCH_CLASS_PARENT:
		if CompilerGlobals.GetActiveClassEntry() != nil && CompilerGlobals.GetActiveClassEntry().parent_name && ZendIsScopeKnown() != 0 {
			ZVAL_STR_COPY(zv, CompilerGlobals.GetActiveClassEntry().parent_name)
			return 1
		}
		return 0
	case ZEND_FETCH_CLASS_STATIC:
		return 0
	case ZEND_FETCH_CLASS_DEFAULT:
		ZVAL_STR(zv, ZendResolveClassNameAst(class_ast))
		return 1
	default:
		break
	}
}

/* }}} */

func ZendVerifyCtConstAccess(c *ZendClassConstant, scope *ZendClassEntry) ZendBool {
	if (Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PUBLIC) != 0 {
		return 1
	} else if (Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PRIVATE) != 0 {
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
			if (ce.GetCeFlags() & ZEND_ACC_RESOLVED_PARENT) != 0 {
				ce = ce.parent
			} else {
				ce = ZendHashFindPtrLc(CompilerGlobals.GetClassTable(), ZSTR_VAL(ce.parent_name), ZSTR_LEN(ce.parent_name))
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
		cc = ZendHashFindPtr(&(CompilerGlobals.GetActiveClassEntry()).constants_table, name)
	} else if fetch_type == ZEND_FETCH_CLASS_DEFAULT && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0 {
		var ce *ZendClassEntry = ZendHashFindPtrLc(CompilerGlobals.GetClassTable(), ZSTR_VAL(class_name), ZSTR_LEN(class_name))
		if ce != nil {
			cc = ZendHashFindPtr(&ce.constants_table, name)
		} else {
			return 0
		}
	} else {
		return 0
	}
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION) != 0 {
		return 0
	}
	if cc == nil || ZendVerifyCtConstAccess(cc, CompilerGlobals.GetActiveClassEntry()) == 0 {
		return 0
	}
	c = &cc.value

	/* Substitute case-sensitive (or lowercase) persistent class constants */

	if Z_TYPE_P(c) < IS_OBJECT {
		ZVAL_COPY_OR_DUP(zv, c)
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
	list = Erealloc(list, b.SizeOf("void *")*(n+2))
	list[n] = item
	list[n+1] = nil
	*((*any)(result)) = list
}

/* }}} */

func ZendDoExtendedStmt() {
	var opline *ZendOp
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_STMT)
}

/* }}} */

func ZendDoExtendedFcallBegin() {
	var opline *ZendOp
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_EXTENDED_FCALL) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_FCALL_BEGIN)
}

/* }}} */

func ZendDoExtendedFcallEnd() {
	var opline *ZendOp
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_EXTENDED_FCALL) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_FCALL_END)
}

/* }}} */

func ZendIsAutoGlobalStr(name string, len_ int) ZendBool {
	var auto_global *ZendAutoGlobal
	if b.Assign(&auto_global, ZendHashStrFindPtr(CompilerGlobals.GetAutoGlobals(), name, len_)) != nil {
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
	if b.Assign(&auto_global, ZendHashFindPtr(CompilerGlobals.GetAutoGlobals(), name)) != nil {
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
	if ZendHashAddMem(CompilerGlobals.GetAutoGlobals(), auto_global.GetName(), &auto_global, b.SizeOf("zend_auto_global")) != nil {
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
		var __ht *HashTable = CompilerGlobals.GetAutoGlobals()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			auto_global = Z_PTR_P(_z)
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
	if CompilerGlobals.GetIncrementLineno() != 0 {
		CompilerGlobals.GetZendLineno()++
		CompilerGlobals.SetIncrementLineno(0)
	}
	ret = LexScan(&zv, elem)
	ZEND_ASSERT(ExecutorGlobals.GetException() == nil || ret == T_ERROR)
	return ret
}

/* }}} */

func ZendInitializeClassData(ce *ZendClassEntry, nullify_handlers ZendBool) {
	var persistent_hashes ZendBool = ce.GetType() == ZEND_INTERNAL_CLASS
	ce.SetRefcount(1)
	ce.SetCeFlags(ZEND_ACC_CONSTANTS_UPDATED)
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_GUARDS) != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
	}
	ce.SetDefaultPropertiesTable(nil)
	ce.SetDefaultStaticMembersTable(nil)
	ZendHashInitEx(&ce.properties_info, 8, nil, b.Cond(persistent_hashes != 0, ZendDestroyPropertyInfoInternal, nil), persistent_hashes, 0)
	ZendHashInitEx(&ce.constants_table, 8, nil, nil, persistent_hashes, 0)
	ZendHashInitEx(&ce.function_table, 8, nil, ZEND_FUNCTION_DTOR, persistent_hashes, 0)
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		ZEND_MAP_PTR_INIT(ce.static_members_table, nil)
	} else {
		ZEND_MAP_PTR_INIT(ce.static_members_table, &ce.default_static_members_table)
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
		if ce.GetType() == ZEND_INTERNAL_CLASS {
			ce.SetModule(nil)
			ce.SetBuiltinFunctions(nil)
		}
	}
}

/* }}} */

func ZendGetCompiledVariableName(op_array *ZendOpArray, var_ uint32) *ZendString {
	return op_array.GetVars()[EX_VAR_TO_NUM(var_)]
}

/* }}} */

func ZendAstAppendStr(left_ast *ZendAst, right_ast *ZendAst) *ZendAst {
	var left_zv *Zval = ZendAstGetZval(left_ast)
	var left *ZendString = Z_STR_P(left_zv)
	var right *ZendString = ZendAstGetStr(right_ast)
	var result *ZendString
	var left_len int = ZSTR_LEN(left)
	var len_ int = left_len + ZSTR_LEN(right) + 1
	result = ZendStringExtend(left, len_, 0)
	ZSTR_VAL(result)[left_len] = '\\'
	memcpy(&ZSTR_VAL(result)[left_len+1], ZSTR_VAL(right), ZSTR_LEN(right))
	ZSTR_VAL(result)[len_] = '0'
	ZendStringReleaseEx(right, 0)
	ZVAL_STR(left_zv, result)
	return left_ast
}

/* }}} */

func ZendNegateNumString(ast *ZendAst) *ZendAst {
	var zv *Zval = ZendAstGetZval(ast)
	if Z_TYPE_P(zv) == IS_LONG {
		if Z_LVAL_P(zv) == 0 {
			ZVAL_NEW_STR(zv, ZendStringInit("-0", b.SizeOf("\"-0\"")-1, 0))
		} else {
			ZEND_ASSERT(Z_LVAL_P(zv) > 0)
			Z_LVAL_P(zv) *= -1
		}
	} else if Z_TYPE_P(zv) == IS_STRING {
		var orig_len int = Z_STRLEN_P(zv)
		Z_STR_P(zv) = ZendStringExtend(Z_STR_P(zv), orig_len+1, 0)
		memmove(Z_STRVAL_P(zv)+1, Z_STRVAL_P(zv), orig_len+1)
		Z_STRVAL_P(zv)[0] = '-'
	} else {
		ZEND_ASSERT(false)
	}
	return ast
}

/* }}} */

func ZendVerifyNamespace() {
	if FC(has_bracketed_namespaces) && !(FC(in_namespace)) {
		ZendErrorNoreturn(E_COMPILE_ERROR, "No code may exist outside of namespace {}")
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

	for end >= path && IS_SLASH_P(end) {
		end--
	}
	if end < path {

		/* The path only contained slashes */

		path[0] = DEFAULT_SLASH
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip filename */

	for end >= path && !(IS_SLASH_P(end)) {
		end--
	}
	if end < path {

		/* No slash found, therefore return '.' */

		path[0] = '.'
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip slashes which came before the file name */

	for end >= path && IS_SLASH_P(end) {
		end--
	}
	if end < path {
		path[0] = DEFAULT_SLASH
		path[1] = '0'
		return 1 + len_adjust
	}
	*(end + 1) = '0'
	return size_t(end+1-path) + len_adjust
}

/* }}} */

func ZendAdjustForFetchType(opline *ZendOp, result *Znode, type_ uint32) {
	var factor ZendUchar = b.Cond(opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_R, 1, 3)
	switch type_ {
	case BP_VAR_R:
		opline.SetResultType(IS_TMP_VAR)
		result.SetOpType(IS_TMP_VAR)
		return
	case BP_VAR_W:
		opline.SetOpcode(opline.GetOpcode() + 1*factor)
		return
	case BP_VAR_RW:
		opline.SetOpcode(opline.GetOpcode() + 2*factor)
		return
	case BP_VAR_IS:
		opline.SetResultType(IS_TMP_VAR)
		result.SetOpType(IS_TMP_VAR)
		opline.SetOpcode(opline.GetOpcode() + 3*factor)
		return
	case BP_VAR_FUNC_ARG:
		opline.SetOpcode(opline.GetOpcode() + 4*factor)
		return
	case BP_VAR_UNSET:
		opline.SetOpcode(opline.GetOpcode() + 5*factor)
		return
	default:
		break
	}
}

/* }}} */

func ZendMakeVarResult(result *Znode, opline *ZendOp) {
	opline.SetResultType(IS_VAR)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == IS_CONST {
		ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetResult()))
	} else {
		result.SetOp(opline.GetResult())
	}
}

/* }}} */

func ZendMakeTmpResult(result *Znode, opline *ZendOp) {
	opline.SetResultType(IS_TMP_VAR)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == IS_CONST {
		ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetResult()))
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
		if op1.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
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
		if op1.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
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

	if CompilerGlobals.GetActiveOpArray().GetLast() != 0 && CompilerGlobals.GetActiveOpArray().GetOpcodes()[CompilerGlobals.GetActiveOpArray().GetLast()-1].GetOpcode() == ZEND_TICKS {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_TICKS)
	opline.SetExtendedValue(FC(declarables).ticks)
}

/* }}} */

func ZendEmitOpData(value *Znode) *ZendOp {
	return ZendEmitOp(nil, ZEND_OP_DATA, value, nil)
}

/* }}} */

func ZendEmitJump(opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *ZendOp = ZendEmitOp(nil, ZEND_JMP, nil, nil)
	opline.GetOp1().SetOplineNum(opnum_target)
	return opnum
}

/* }}} */

func ZendIsSmartBranch(opline *ZendOp) int {
	switch opline.GetOpcode() {
	case ZEND_IS_IDENTICAL:

	case ZEND_IS_NOT_IDENTICAL:

	case ZEND_IS_EQUAL:

	case ZEND_IS_NOT_EQUAL:

	case ZEND_IS_SMALLER:

	case ZEND_IS_SMALLER_OR_EQUAL:

	case ZEND_CASE:

	case ZEND_ISSET_ISEMPTY_CV:

	case ZEND_ISSET_ISEMPTY_VAR:

	case ZEND_ISSET_ISEMPTY_DIM_OBJ:

	case ZEND_ISSET_ISEMPTY_PROP_OBJ:

	case ZEND_ISSET_ISEMPTY_STATIC_PROP:

	case ZEND_INSTANCEOF:

	case ZEND_TYPE_CHECK:

	case ZEND_DEFINED:

	case ZEND_IN_ARRAY:

	case ZEND_ARRAY_KEY_EXISTS:
		return 1
	default:
		return 0
	}
}

/* }}} */

func ZendEmitCondJump(opcode ZendUchar, cond *Znode, opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *ZendOp
	if (cond.GetOpType()&(IS_CV|IS_CONST)) != 0 && opnum > 0 && ZendIsSmartBranch(CompilerGlobals.GetActiveOpArray().GetOpcodes()+opnum-1) != 0 {

		/* emit extra NOP to avoid incorrect SMART_BRANCH in very rare cases */

		ZendEmitOp(nil, ZEND_NOP, nil, nil)
		opnum = GetNextOpNumber()
	}
	opline = ZendEmitOp(nil, opcode, cond, nil)
	opline.GetOp2().SetOplineNum(opnum_target)
	return opnum
}

/* }}} */

func ZendUpdateJumpTarget(opnum_jump uint32, opnum_target uint32) {
	var opline *ZendOp = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_jump]
	switch opline.GetOpcode() {
	case ZEND_JMP:
		opline.GetOp1().SetOplineNum(opnum_target)
		break
	case ZEND_JMPZ:

	case ZEND_JMPNZ:

	case ZEND_JMPZ_EX:

	case ZEND_JMPNZ_EX:

	case ZEND_JMP_SET:

	case ZEND_COALESCE:
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
		if op1.GetOpType() == IS_CONST {
			tmp_opline.GetOp1().SetConstant(ZendAddLiteral(&op1.u.constant))
		} else {
			tmp_opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		tmp_opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
			tmp_opline.GetOp2().SetConstant(ZendAddLiteral(&op2.u.constant))
		} else {
			tmp_opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeVarResult(result, &tmp_opline)
	}
	ZendStackPush(&(CompilerGlobals.GetDelayedOplinesStack()), &tmp_opline)
	return ZendStackTop(&(CompilerGlobals.GetDelayedOplinesStack()))
}

/* }}} */

func ZendDelayedCompileBegin() uint32 {
	return ZendStackCount(&(CompilerGlobals.GetDelayedOplinesStack()))
}

/* }}} */

func ZendDelayedCompileEnd(offset uint32) *ZendOp {
	var opline *ZendOp = nil
	var oplines *ZendOp = ZendStackBase(&(CompilerGlobals.GetDelayedOplinesStack()))
	var i uint32
	var count uint32 = ZendStackCount(&(CompilerGlobals.GetDelayedOplinesStack()))
	ZEND_ASSERT(count >= offset)
	for i = offset; i < count; i++ {
		opline = GetNextOp()
		memcpy(opline, &oplines[i], b.SizeOf("zend_op"))
	}
	CompilerGlobals.GetDelayedOplinesStack().SetTop(offset)
	return opline
}

/* }}} */

const ZEND_MEMOIZE_NONE = 0
const ZEND_MEMOIZE_COMPILE = 1
const ZEND_MEMOIZE_FETCH = 2

func ZendCompileMemoizedExpr(result *Znode, expr *ZendAst) {
	var memoize_mode int = CompilerGlobals.GetMemoizeMode()
	if memoize_mode == ZEND_MEMOIZE_COMPILE {
		var memoized_result Znode

		/* Go through normal compilation */

		CompilerGlobals.SetMemoizeMode(ZEND_MEMOIZE_NONE)
		ZendCompileExpr(result, expr)
		CompilerGlobals.SetMemoizeMode(ZEND_MEMOIZE_COMPILE)
		if result.GetOpType() == IS_VAR {
			ZendEmitOp(&memoized_result, ZEND_COPY_TMP, result, nil)
		} else if result.GetOpType() == IS_TMP_VAR {
			ZendEmitOpTmp(&memoized_result, ZEND_COPY_TMP, result, nil)
		} else {
			if result.GetOpType() == IS_CONST {
				Z_TRY_ADDREF(result.GetConstant())
			}
			memoized_result = *result
		}
		ZendHashIndexUpdateMem(CompilerGlobals.GetMemoizedExprs(), uintPtr(expr), &memoized_result, b.SizeOf("znode"))
	} else if memoize_mode == ZEND_MEMOIZE_FETCH {
		var memoized_result *Znode = ZendHashIndexFindPtr(CompilerGlobals.GetMemoizedExprs(), uintPtr(expr))
		*result = *memoized_result
		if result.GetOpType() == IS_CONST {
			Z_TRY_ADDREF(result.GetConstant())
		}
	} else {
		ZEND_ASSERT(false)
	}
}

/* }}} */

func ZendEmitReturnTypeCheck(expr *Znode, return_info *ZendArgInfo, implicit ZendBool) {
	if ZEND_TYPE_IS_SET(return_info.GetType()) {
		var opline *ZendOp

		/* `return ...;` is illegal in a void function (but `return;` isn't) */

		if ZEND_TYPE_CODE(return_info.GetType()) == IS_VOID {
			if expr != nil {
				if expr.GetOpType() == IS_CONST && Z_TYPE(expr.GetConstant()) == IS_NULL {
					ZendErrorNoreturn(E_COMPILE_ERROR, "A void function must not return a value "+"(did you mean \"return;\" instead of \"return null;\"?)")
				} else {
					ZendErrorNoreturn(E_COMPILE_ERROR, "A void function must not return a value")
				}
			}

			/* we don't need run-time check */

			return

			/* we don't need run-time check */

		}
		if expr == nil && implicit == 0 {
			if ZEND_TYPE_ALLOW_NULL(return_info.GetType()) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "A function with return type must return a value "+"(did you mean \"return null;\" instead of \"return;\"?)")
			} else {
				ZendErrorNoreturn(E_COMPILE_ERROR, "A function with return type must return a value")
			}
		}
		if expr != nil && expr.GetOpType() == IS_CONST {
			if ZEND_TYPE_CODE(return_info.GetType()) == Z_TYPE(expr.GetConstant()) || ZEND_TYPE_CODE(return_info.GetType()) == _IS_BOOL && (Z_TYPE(expr.GetConstant()) == IS_FALSE || Z_TYPE(expr.GetConstant()) == IS_TRUE) || ZEND_TYPE_ALLOW_NULL(return_info.GetType()) && Z_TYPE(expr.GetConstant()) == IS_NULL {

				/* we don't need run-time check */

				return

				/* we don't need run-time check */

			}
		}
		opline = ZendEmitOp(nil, ZEND_VERIFY_RETURN_TYPE, expr, nil)
		if expr != nil && expr.GetOpType() == IS_CONST {
			expr.SetOpType(IS_TMP_VAR)
			opline.SetResultType(expr.GetOpType())
			expr.GetOp().SetVar(GetTemporaryVariable())
			opline.GetResult().SetVar(expr.GetOp().GetVar())
		}
		if ZEND_TYPE_IS_CLASS(return_info.GetType()) {
			opline.GetOp2().SetNum(CompilerGlobals.GetActiveOpArray().GetCacheSize())
			CompilerGlobals.GetActiveOpArray().SetCacheSize(CompilerGlobals.GetActiveOpArray().GetCacheSize() + b.SizeOf("void *"))
		} else {
			opline.GetOp2().SetNum(-1)
		}
	}
}

/* }}} */

func ZendEmitFinalReturn(return_one int) {
	var zn Znode
	var ret *ZendOp
	var returns_reference ZendBool = (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags()&ZEND_ACC_HAS_RETURN_TYPE) != 0 && (CompilerGlobals.GetActiveOpArray().GetFnFlags()&ZEND_ACC_GENERATOR) == 0 {
		ZendEmitReturnTypeCheck(nil, CompilerGlobals.GetActiveOpArray().GetArgInfo()-1, 1)
	}
	zn.SetOpType(IS_CONST)
	if return_one != 0 {
		ZVAL_LONG(&zn.u.constant, 1)
	} else {
		ZVAL_NULL(&zn.u.constant)
	}
	ret = ZendEmitOp(nil, b.Cond(returns_reference != 0, ZEND_RETURN_BY_REF, ZEND_RETURN), &zn, nil)
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
	return ZEND_FETCH_CLASS_DEFAULT == ZendGetClassFetchTypeAst(name_ast)
}

/* }}} */

func ZendHandleNumericOp(node *Znode) {
	if node.GetOpType() == IS_CONST && Z_TYPE(node.GetConstant()) == IS_STRING {
		var index ZendUlong
		if ZEND_HANDLE_NUMERIC(Z_STR(node.GetConstant()), index) != 0 {
			ZvalPtrDtor(&node.u.constant)
			ZVAL_LONG(&node.u.constant, index)
		}
	}
}

/* }}} */

func ZendHandleNumericDim(opline *ZendOp, dim_node *Znode) {
	if Z_TYPE(dim_node.GetConstant()) == IS_STRING {
		var index ZendUlong
		if ZEND_HANDLE_NUMERIC(Z_STR(dim_node.GetConstant()), index) != 0 {

			/* For numeric indexes we also keep the original value to use by ArrayAccess
			 * See bug #63217
			 */

			var c int = ZendAddLiteral(&dim_node.u.constant)
			ZEND_ASSERT(opline.GetOp2().GetConstant()+1 == c)
			ZVAL_LONG(CT_CONSTANT(opline.GetOp2()), index)
			Z_EXTRA_P(CT_CONSTANT(opline.GetOp2())) = ZEND_EXTRA_VALUE
			return
		}
	}
}

/* }}} */

func ZendSetClassNameOp1(opline *ZendOp, class_node *Znode) {
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp1Type(IS_CONST)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(Z_STR(class_node.GetConstant())))
	} else {
		opline.SetOp1Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
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
		if name_node.GetOpType() == IS_CONST {
			var name *ZendString
			if Z_TYPE(name_node.GetConstant()) != IS_STRING {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Illegal class name")
			}
			name = Z_STR(name_node.GetConstant())
			fetch_type = ZendGetClassFetchType(name)
			if fetch_type == ZEND_FETCH_CLASS_DEFAULT {
				result.SetOpType(IS_CONST)
				ZVAL_STR(&result.u.constant, ZendResolveClassName(name, ZEND_NAME_FQ))
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				result.SetOpType(IS_UNUSED)
				result.GetOp().SetNum(fetch_type | fetch_flags)
			}
			ZendStringReleaseEx(name, 0)
		} else {
			var opline *ZendOp = ZendEmitOp(result, ZEND_FETCH_CLASS, nil, &name_node)
			opline.GetOp1().SetNum(ZEND_FETCH_CLASS_DEFAULT | fetch_flags)
		}
		return
	}

	/* Fully qualified names are always default refs */

	if name_ast.GetAttr() == ZEND_NAME_FQ {
		result.SetOpType(IS_CONST)
		ZVAL_STR(&result.u.constant, ZendResolveClassNameAst(name_ast))
		return
	}
	fetch_type = ZendGetClassFetchType(ZendAstGetStr(name_ast))
	if ZEND_FETCH_CLASS_DEFAULT == fetch_type {
		result.SetOpType(IS_CONST)
		ZVAL_STR(&result.u.constant, ZendResolveClassNameAst(name_ast))
	} else {
		ZendEnsureValidClassFetchType(fetch_type)
		result.SetOpType(IS_UNUSED)
		result.GetOp().SetNum(fetch_type | fetch_flags)
	}
}

/* }}} */

func ZendTryCompileCv(result *Znode, ast *ZendAst) int {
	var name_ast *ZendAst = ast.GetChild()[0]
	if name_ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(name_ast)
		var name *ZendString
		if EXPECTED(Z_TYPE_P(zv) == IS_STRING) {
			name = ZvalMakeInternedString(zv)
		} else {
			name = ZendNewInternedString(ZvalGetStringFunc(zv))
		}
		if ZendIsAutoGlobal(name) != 0 {
			return FAILURE
		}
		result.SetOpType(IS_CV)
		result.GetOp().SetVar(LookupCv(name))
		if UNEXPECTED(Z_TYPE_P(zv) != IS_STRING) {
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
	if name_node.GetOpType() == IS_CONST {
		ConvertToString(&name_node.u.constant)
	}
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, ZEND_FETCH_R, &name_node, nil)
	} else {
		opline = ZendEmitOp(result, ZEND_FETCH_R, &name_node, nil)
	}
	if name_node.GetOpType() == IS_CONST && ZendIsAutoGlobal(Z_STR(name_node.GetConstant())) != 0 {
		opline.SetExtendedValue(ZEND_FETCH_GLOBAL)
	} else {
		opline.SetExtendedValue(ZEND_FETCH_LOCAL)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}

/* }}} */

func IsThisFetch(ast *ZendAst) ZendBool {
	if ast.GetKind() == ZEND_AST_VAR && ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL {
		var name *Zval = ZendAstGetZval(ast.GetChild()[0])
		return Z_TYPE_P(name) == IS_STRING && ZendStringEqualsLiteral(Z_STR_P(name), "this")
	}
	return 0
}

/* }}} */

func ZendCompileSimpleVar(result *Znode, ast *ZendAst, type_ uint32, delayed int) *ZendOp {
	if IsThisFetch(ast) != 0 {
		var opline *ZendOp = ZendEmitOp(result, ZEND_FETCH_THIS, nil, nil)
		if type_ == BP_VAR_R || type_ == BP_VAR_IS {
			opline.SetResultType(IS_TMP_VAR)
			result.SetOpType(IS_TMP_VAR)
		}
		CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_USES_THIS)
		return opline
	} else if ZendTryCompileCv(result, ast) == FAILURE {
		return ZendCompileSimpleVarNoCv(result, ast, type_, delayed)
	}
	return nil
}

/* }}} */

func ZendSeparateIfCallAndWrite(node *Znode, ast *ZendAst, type_ uint32) {
	if type_ != BP_VAR_R && type_ != BP_VAR_IS && ZendIsCall(ast) != 0 {
		if node.GetOpType() == IS_VAR {
			var opline *ZendOp = ZendEmitOp(nil, ZEND_SEPARATE, node, nil)
			opline.SetResultType(IS_VAR)
			opline.GetResult().SetVar(opline.GetOp1().GetVar())
		} else {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use result of built-in function in write context")
		}
	}
}

/* }}} */

func ZendEmitAssignZnode(var_ast *ZendAst, value_node *Znode) {
	var dummy_node Znode
	var assign_ast *ZendAst = ZendAstCreate(ZEND_AST_ASSIGN, var_ast, ZendAstCreateZnode(value_node))
	ZendCompileAssign(&dummy_node, assign_ast)
	ZendDoFree(&dummy_node)
}

/* }}} */

func ZendDelayedCompileDim(result *Znode, ast *ZendAst, type_ uint32) *ZendOp {
	if ast.GetAttr() == ZEND_DIM_ALTERNATIVE_SYNTAX {
		ZendError(E_DEPRECATED, "Array and string offset access syntax with curly braces is deprecated")
	}
	var var_ast *ZendAst = ast.GetChild()[0]
	var dim_ast *ZendAst = ast.GetChild()[1]
	var opline *ZendOp
	var var_node Znode
	var dim_node Znode
	opline = ZendDelayedCompileVar(&var_node, var_ast, type_, 0)
	if opline != nil && type_ == BP_VAR_W && (opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W || opline.GetOpcode() == ZEND_FETCH_OBJ_W) {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_DIM_WRITE)
	}
	ZendSeparateIfCallAndWrite(&var_node, var_ast, type_)
	if dim_ast == nil {
		if type_ == BP_VAR_R || type_ == BP_VAR_IS {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if type_ == BP_VAR_UNSET {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use [] for unsetting")
		}
		dim_node.SetOpType(IS_UNUSED)
	} else {
		ZendCompileExpr(&dim_node, dim_ast)
	}
	opline = ZendDelayedEmitOp(result, ZEND_FETCH_DIM_R, &var_node, &dim_node)
	ZendAdjustForFetchType(opline, result, type_)
	if dim_node.GetOpType() == IS_CONST {
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
		obj_node.SetOpType(IS_UNUSED)
		CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_USES_THIS)
	} else {
		opline = ZendDelayedCompileVar(&obj_node, obj_ast, type_, 0)
		if opline != nil && type_ == BP_VAR_W && (opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W || opline.GetOpcode() == ZEND_FETCH_OBJ_W) {
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_OBJ_WRITE)
		}
		ZendSeparateIfCallAndWrite(&obj_node, obj_ast, type_)
	}
	ZendCompileExpr(&prop_node, prop_ast)
	opline = ZendDelayedEmitOp(result, ZEND_FETCH_OBJ_R, &obj_node, &prop_node)
	if opline.GetOp2Type() == IS_CONST {
		ConvertToString(CT_CONSTANT(opline.GetOp2()))
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
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
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
	ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	ZendCompileExpr(&prop_node, prop_ast)
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, ZEND_FETCH_STATIC_PROP_R, &prop_node, nil)
	} else {
		opline = ZendEmitOp(result, ZEND_FETCH_STATIC_PROP_R, &prop_node, nil)
	}
	if opline.GetOp1Type() == IS_CONST {
		ConvertToString(CT_CONSTANT(opline.GetOp1()))
		opline.SetExtendedValue(ZendAllocCacheSlots(3))
	}
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(Z_STR(class_node.GetConstant())))
		if opline.GetOp1Type() != IS_CONST {
			opline.SetExtendedValue(ZendAllocCacheSlot())
		}
	} else {
		opline.SetOp2Type(&class_node.GetOpType())
		if &class_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&class_node).u.constant))
		} else {
			opline.SetOp2(&class_node.GetOp())
		}
	}
	if by_ref != 0 && (type_ == BP_VAR_W || type_ == BP_VAR_FUNC_ARG) {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}

/* }}} */

func ZendVerifyListAssignTarget(var_ast *ZendAst, old_style ZendBool) {
	if var_ast.GetKind() == ZEND_AST_ARRAY {
		if var_ast.GetAttr() == ZEND_ARRAY_SYNTAX_LONG {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot assign to array(), use [] instead")
		}
		if old_style != var_ast.GetAttr() {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot mix [] and list()")
		}
	} else if ZendCanWriteToVariable(var_ast) == 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Assignments can only happen to writable values")
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
	if list.GetChildren() != 0 && expr_node.GetOpType() == IS_CONST && Z_TYPE(expr_node.GetConstant()) == IS_STRING {
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
				ZendError(E_COMPILE_ERROR, "Cannot use empty array entries in keyed array assignment")
			} else {
				continue
			}
		}
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			ZendError(E_COMPILE_ERROR, "Spread operator is not supported in assignments")
		}
		var_ast = elem_ast.GetChild()[0]
		key_ast = elem_ast.GetChild()[1]
		has_elems = 1
		if is_keyed != 0 {
			if key_ast == nil {
				ZendError(E_COMPILE_ERROR, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			ZendCompileExpr(&dim_node, key_ast)
		} else {
			if key_ast != nil {
				ZendError(E_COMPILE_ERROR, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			dim_node.SetOpType(IS_CONST)
			ZVAL_LONG(&dim_node.u.constant, i)
		}
		if expr_node.GetOpType() == IS_CONST {
			Z_TRY_ADDREF(expr_node.GetConstant())
		}
		ZendVerifyListAssignTarget(var_ast, old_style)
		opline = ZendEmitOp(&fetch_result, b.CondF1(elem_ast.GetAttr() != 0, func() __auto__ {
			if expr_node.GetOpType() == IS_CV {
				return ZEND_FETCH_DIM_W
			} else {
				return ZEND_FETCH_LIST_W
			}
		}, ZEND_FETCH_LIST_R), expr_node, &dim_node)
		if dim_node.GetOpType() == IS_CONST {
			ZendHandleNumericDim(opline, &dim_node)
		}
		if var_ast.GetKind() == ZEND_AST_ARRAY {
			if elem_ast.GetAttr() != 0 {
				ZendEmitOp(&fetch_result, ZEND_MAKE_REF, &fetch_result, nil)
			}
			ZendCompileListAssign(nil, var_ast, &fetch_result, var_ast.GetAttr())
		} else if elem_ast.GetAttr() != 0 {
			ZendEmitAssignRefZnode(var_ast, &fetch_result)
		} else {
			ZendEmitAssignZnode(var_ast, &fetch_result)
		}
	}
	if has_elems == 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use empty list")
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
		ZendErrorNoreturn(E_COMPILE_ERROR, "Can't use function return value in write context")
	}
	if ast.GetKind() == ZEND_AST_METHOD_CALL || ast.GetKind() == ZEND_AST_STATIC_CALL {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Can't use method return value in write context")
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
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(&var_node, var_ast, BP_VAR_W, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		ZendEmitOp(result, ZEND_ASSIGN, &var_node, &expr_node)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(result, var_ast, BP_VAR_W, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileDim(result, var_ast, BP_VAR_W)
		if ZendIsAssignToSelf(var_ast, expr_ast) != 0 && IsThisFetch(expr_ast) == 0 {

			/* $a[0] = $a should evaluate the right $a first */

			var cv_node Znode
			if ZendTryCompileCv(&cv_node, expr_ast) == FAILURE {
				ZendCompileSimpleVarNoCv(&expr_node, expr_ast, BP_VAR_R, 0)
			} else {
				ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &cv_node, nil)
			}
		} else {
			ZendCompileExpr(&expr_node, expr_ast)
		}
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_DIM)
		opline = ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileProp(result, var_ast, BP_VAR_W)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_OBJ)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_ARRAY:
		if ZendPropagateListRefs(var_ast) != 0 {
			if ZendIsVariableOrCall(expr_ast) == 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot assign reference to non referencable value")
			}
			ZendCompileVar(&expr_node, expr_ast, BP_VAR_W, 1)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

			ZendEmitOp(&expr_node, ZEND_MAKE_REF, &expr_node, nil)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

		} else {
			if expr_ast.GetKind() == ZEND_AST_VAR {

				/* list($a, $b) = $a should evaluate the right $a first */

				var cv_node Znode
				if ZendTryCompileCv(&cv_node, expr_ast) == FAILURE {
					ZendCompileSimpleVarNoCv(&expr_node, expr_ast, BP_VAR_R, 0)
				} else {
					ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &cv_node, nil)
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
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(target_ast)
	offset = ZendDelayedCompileBegin()
	ZendDelayedCompileVar(&target_node, target_ast, BP_VAR_W, 1)
	ZendCompileVar(&source_node, source_ast, BP_VAR_W, 1)
	if (target_ast.GetKind() != ZEND_AST_VAR || target_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL) && source_node.GetOpType() != IS_CV {

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

		ZendEmitOp(&source_node, ZEND_MAKE_REF, &source_node, nil)

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

	}
	opline = ZendDelayedCompileEnd(offset)
	if source_node.GetOpType() != IS_VAR && ZendIsCall(source_ast) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use result of built-in function in write context")
	}
	if ZendIsCall(source_ast) != 0 {
		flags = ZEND_RETURNS_FUNCTION
	} else {
		flags = 0
	}
	if opline != nil && opline.GetOpcode() == ZEND_FETCH_OBJ_W {
		opline.SetOpcode(ZEND_ASSIGN_OBJ_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ ZEND_FETCH_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else if opline != nil && opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W {
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ ZEND_FETCH_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else {
		opline = ZendEmitOp(result, ZEND_ASSIGN_REF, &target_node, &source_node)
		opline.SetExtendedValue(flags)
	}
}

/* }}} */

func ZendEmitAssignRefZnode(var_ast *ZendAst, value_node *Znode) {
	var dummy_node Znode
	var assign_ast *ZendAst = ZendAstCreate(ZEND_AST_ASSIGN_REF, var_ast, ZendAstCreateZnode(value_node))
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
		ZendDelayedCompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		opline = ZendEmitOp(result, ZEND_ASSIGN_OP, &var_node, &expr_node)
		opline.SetExtendedValue(opcode)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(result, var_ast, BP_VAR_RW, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP_OP)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileDim(result, var_ast, BP_VAR_RW)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_DIM_OP)
		opline.SetExtendedValue(opcode)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileProp(result, var_ast, BP_VAR_RW)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(ZEND_ASSIGN_OBJ_OP)
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
			opline = ZendEmitOp(nil, ZEND_SEND_UNPACK, &arg_node, nil)
			opline.GetOp2().SetNum(arg_count)
			opline.GetResult().SetVar(uint32(ZendIntptrT(ZEND_CALL_ARG(nil, arg_count))))
			continue
		}
		if uses_arg_unpack != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use positional argument after argument unpacking")
		}
		arg_count++
		if ZendIsVariableOrCall(arg) != 0 {
			if ZendIsCall(arg) != 0 {
				ZendCompileVar(&arg_node, arg, BP_VAR_R, 0)
				if (arg_node.GetOpType() & (IS_CONST | IS_TMP_VAR)) != 0 {

					/* Function call was converted into builtin instruction */

					if fbc == nil || ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAL_EX
					} else {
						opcode = ZEND_SEND_VAL
					}

					/* Function call was converted into builtin instruction */

				} else {
					if fbc != nil {
						if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
							opcode = ZEND_SEND_VAR_NO_REF
						} else if ARG_MAY_BE_SENT_BY_REF(fbc, arg_num) != 0 {
							opcode = ZEND_SEND_VAL
						} else {
							opcode = ZEND_SEND_VAR
						}
					} else {
						opcode = ZEND_SEND_VAR_NO_REF_EX
					}
				}
			} else if fbc != nil {
				if ARG_SHOULD_BE_SENT_BY_REF(fbc, arg_num) != 0 {
					ZendCompileVar(&arg_node, arg, BP_VAR_W, 1)
					opcode = ZEND_SEND_REF
				} else {
					ZendCompileVar(&arg_node, arg, BP_VAR_R, 0)
					if arg_node.GetOpType() == IS_TMP_VAR {
						opcode = ZEND_SEND_VAL
					} else {
						opcode = ZEND_SEND_VAR
					}
				}
			} else {
				for {
					if arg.GetKind() == ZEND_AST_VAR {
						CompilerGlobals.SetZendLineno(ZendAstGetLineno(ast))
						if IsThisFetch(arg) != 0 {
							ZendEmitOp(&arg_node, ZEND_FETCH_THIS, nil, nil)
							opcode = ZEND_SEND_VAR_EX
							CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_USES_THIS)
							break
						} else if ZendTryCompileCv(&arg_node, arg) == SUCCESS {
							opcode = ZEND_SEND_VAR_EX
							break
						}
					}
					opline = ZendEmitOp(nil, ZEND_CHECK_FUNC_ARG, nil, nil)
					opline.GetOp2().SetNum(arg_num)
					ZendCompileVar(&arg_node, arg, BP_VAR_FUNC_ARG, 1)
					opcode = ZEND_SEND_FUNC_ARG
					break
				}
			}
		} else {
			ZendCompileExpr(&arg_node, arg)
			if arg_node.GetOpType() == IS_VAR {

				/* pass ++$a or something similar */

				if fbc != nil {
					if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAR_NO_REF
					} else if ARG_MAY_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAL
					} else {
						opcode = ZEND_SEND_VAR
					}
				} else {
					opcode = ZEND_SEND_VAR_NO_REF_EX
				}

				/* pass ++$a or something similar */

			} else if arg_node.GetOpType() == IS_CV {
				if fbc != nil {
					if ARG_SHOULD_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_REF
					} else {
						opcode = ZEND_SEND_VAR
					}
				} else {
					opcode = ZEND_SEND_VAR_EX
				}
			} else {
				if fbc != nil {
					opcode = ZEND_SEND_VAL
					if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Only variables can be passed by reference")
					}
				} else {
					opcode = ZEND_SEND_VAL_EX
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, &arg_node, nil)
		opline.GetOp2().SetOplineNum(arg_num)
		opline.GetResult().SetVar(uint32(ZendIntptrT(ZEND_CALL_ARG(nil, arg_num))))
	}
	return arg_count
}

/* }}} */

func ZendGetCallOp(init_op *ZendOp, fbc *ZendFunction) ZendUchar {
	if fbc != nil {
		if fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) == 0 {
			if init_op.GetOpcode() == ZEND_INIT_FCALL && ZendExecuteInternal == nil {
				if (fbc.GetFnFlags() & (ZEND_ACC_ABSTRACT | ZEND_ACC_DEPRECATED | ZEND_ACC_HAS_TYPE_HINTS | ZEND_ACC_RETURN_REFERENCE)) == 0 {
					return ZEND_DO_ICALL
				} else {
					return ZEND_DO_FCALL_BY_NAME
				}
			}
		} else if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_IGNORE_USER_FUNCTIONS) == 0 {
			if ZendExecuteEx == ExecuteEx && (fbc.GetFnFlags()&ZEND_ACC_ABSTRACT) == 0 {
				return ZEND_DO_UCALL
			}
		}
	} else if ZendExecuteEx == ExecuteEx && ZendExecuteInternal == nil && (init_op.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || init_op.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME) {
		return ZEND_DO_FCALL_BY_NAME
	}
	return ZEND_DO_FCALL
}

/* }}} */

func ZendCompileCallCommon(result *Znode, args_ast *ZendAst, fbc *ZendFunction) {
	var opline *ZendOp
	var opnum_init uint32 = GetNextOpNumber() - 1
	var arg_count uint32
	arg_count = ZendCompileArgs(args_ast, fbc)
	ZendDoExtendedFcallBegin()
	opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_init]
	opline.SetExtendedValue(arg_count)
	if opline.GetOpcode() == ZEND_INIT_FCALL {
		opline.GetOp1().SetNum(ZendVmCalcUsedStack(arg_count, fbc))
	}
	opline = ZendEmitOp(result, ZendGetCallOp(opline, fbc), nil, nil)
	ZendDoExtendedFcallEnd()
}

/* }}} */

func ZendCompileFunctionName(name_node *Znode, name_ast *ZendAst) ZendBool {
	var orig_name *ZendString = ZendAstGetStr(name_ast)
	var is_fully_qualified ZendBool
	name_node.SetOpType(IS_CONST)
	ZVAL_STR(&name_node.u.constant, ZendResolveFunctionName(orig_name, name_ast.GetAttr(), &is_fully_qualified))
	return is_fully_qualified == 0 && FC(current_namespace)
}

/* }}} */

func ZendCompileNsCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	var opline *ZendOp = GetNextOp()
	opline.SetOpcode(ZEND_INIT_NS_FCALL_BY_NAME)
	opline.SetOp2Type(IS_CONST)
	opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(Z_STR(name_node.GetConstant())))
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	ZendCompileCallCommon(result, args_ast, nil)
}

/* }}} */

func ZendCompileDynamicCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	if name_node.GetOpType() == IS_CONST && Z_TYPE(name_node.GetConstant()) == IS_STRING {
		var colon *byte
		var str *ZendString = Z_STR(name_node.GetConstant())
		if b.Assign(&colon, ZendMemrchr(ZSTR_VAL(str), ':', ZSTR_LEN(str))) != nil && colon > ZSTR_VAL(str) && (*(colon - 1)) == ':' {
			var class *ZendString = ZendStringInit(ZSTR_VAL(str), colon-ZSTR_VAL(str)-1, 0)
			var method *ZendString = ZendStringInit(colon+1, ZSTR_LEN(str)-(colon-ZSTR_VAL(str))-1, 0)
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
			opline.SetOp1Type(IS_CONST)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class))
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method))

			/* 2 slots, for class and method */

			opline.GetResult().SetNum(ZendAllocCacheSlots(2))
			ZvalPtrDtor(&name_node.u.constant)
		} else {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_FCALL_BY_NAME)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(str))
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
	} else {
		ZendEmitOp(nil, ZEND_INIT_DYNAMIC_CALL, nil, name_node)
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
	if (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_NO_BUILTIN_STRLEN) != 0 || args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	if arg_node.GetOpType() == IS_CONST && Z_TYPE(arg_node.GetConstant()) == IS_STRING {
		result.SetOpType(IS_CONST)
		ZVAL_LONG(&result.u.constant, Z_STRLEN(arg_node.GetConstant()))
		ZvalPtrDtorStr(&arg_node.u.constant)
	} else {
		ZendEmitOpTmp(result, ZEND_STRLEN, &arg_node, nil)
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
	opline = ZendEmitOpTmp(result, ZEND_TYPE_CHECK, &arg_node, nil)
	if type_ != _IS_BOOL {
		opline.SetExtendedValue(1 << type_)
	} else {
		opline.SetExtendedValue(1<<IS_FALSE | 1<<IS_TRUE)
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
	opline = ZendEmitOpTmp(result, ZEND_CAST, &arg_node, nil)
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
	if ZendMemrchr(ZSTR_VAL(name), '\\', ZSTR_LEN(name)) || ZendMemrchr(ZSTR_VAL(name), ':', ZSTR_LEN(name)) {
		ZendStringReleaseEx(name, 0)
		return FAILURE
	}
	if ZendTryCtEvalConst(&result.u.constant, name, 0) != 0 {
		ZendStringReleaseEx(name, 0)
		ZvalPtrDtor(&result.u.constant)
		ZVAL_TRUE(&result.u.constant)
		result.SetOpType(IS_CONST)
		return SUCCESS
	}
	opline = ZendEmitOpTmp(result, ZEND_DEFINED, nil, nil)
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), name)
	opline.SetExtendedValue(ZendAllocCacheSlot())

	/* Lowercase constant name in a separate literal */

	var c Zval
	var lcname *ZendString = ZendStringTolower(name)
	ZVAL_NEW_STR(&c, lcname)
	ZendAddLiteral(&c)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncChr(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && Z_TYPE_P(ZendAstGetZval(args.GetChild()[0])) == IS_LONG {
		var c ZendLong = Z_LVAL_P(ZendAstGetZval(args.GetChild()[0])) & 0xff
		result.SetOpType(IS_CONST)
		ZVAL_INTERNED_STR(&result.u.constant, ZSTR_CHAR(c))
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileFuncOrd(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && Z_TYPE_P(ZendAstGetZval(args.GetChild()[0])) == IS_STRING {
		result.SetOpType(IS_CONST)
		ZVAL_LONG(&result.u.constant, uint8(Z_STRVAL_P(ZendAstGetZval(args.GetChild()[0]))[0]))
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func FbcIsFinalized(fbc *ZendFunction) ZendBool {
	return !(ZEND_USER_CODE(fbc.GetType())) || (fbc.GetFnFlags()&ZEND_ACC_DONE_PASS_TWO) != 0
}
func ZendTryCompileCtBoundInitUserFunc(name_ast *ZendAst, num_args uint32) int {
	var name *ZendString
	var lcname *ZendString
	var fbc *ZendFunction
	var opline *ZendOp
	if name_ast.GetKind() != ZEND_AST_ZVAL || Z_TYPE_P(ZendAstGetZval(name_ast)) != IS_STRING {
		return FAILURE
	}
	name = ZendAstGetStr(name_ast)
	lcname = ZendStringTolower(name)
	fbc = ZendHashFindPtr(CompilerGlobals.GetFunctionTable(), lcname)
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CompilerGlobals.GetActiveOpArray().GetFilename() {
		ZendStringReleaseEx(lcname, 0)
		return FAILURE
	}
	opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, nil)
	opline.SetExtendedValue(num_args)
	opline.GetOp1().SetNum(ZendVmCalcUsedStack(num_args, fbc))
	opline.SetOp2Type(IS_CONST)
	LITERAL_STR(opline.GetOp2(), lcname)
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
	opline = ZendEmitOp(nil, ZEND_INIT_USER_CALL, nil, &name_node)
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), ZendStringCopy(orig_func_name))
	opline.SetExtendedValue(num_args)
}

/* }}} */

func ZendCompileFuncCufa(result *Znode, args *ZendAstList, lcname *ZendString) int {
	var arg_node Znode
	if args.GetChildren() != 2 {
		return FAILURE
	}
	ZendCompileInitUserFunc(args.GetChild()[0], 0, lcname)
	if args.GetChild()[1].GetKind() == ZEND_AST_CALL && args.GetChild()[1].GetChild()[0].GetKind() == ZEND_AST_ZVAL && Z_TYPE_P(ZendAstGetZval(args.GetChild()[1].GetChild()[0])) == IS_STRING && args.GetChild()[1].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST {
		var orig_name *ZendString = ZendAstGetStr(args.GetChild()[1].GetChild()[0])
		var list *ZendAstList = ZendAstGetList(args.GetChild()[1].GetChild()[1])
		var is_fully_qualified ZendBool
		var name *ZendString = ZendResolveFunctionName(orig_name, args.GetChild()[1].GetChild()[0].GetAttr(), &is_fully_qualified)
		if ZendStringEqualsLiteralCi(name, "array_slice") && list.GetChildren() == 3 && list.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
			var zv *Zval = ZendAstGetZval(list.GetChild()[1])
			if Z_TYPE_P(zv) == IS_LONG && Z_LVAL_P(zv) >= 0 && Z_LVAL_P(zv) <= 0x7fffffff {
				var opline *ZendOp
				var len_node Znode
				ZendCompileExpr(&arg_node, list.GetChild()[0])
				ZendCompileExpr(&len_node, list.GetChild()[2])
				opline = ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, &len_node)
				opline.SetExtendedValue(Z_LVAL_P(zv))
				ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
				ZendStringReleaseEx(name, 0)
				return SUCCESS
			}
		}
		ZendStringReleaseEx(name, 0)
	}
	ZendCompileExpr(&arg_node, args.GetChild()[1])
	ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, nil)
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
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
		opline = ZendEmitOp(nil, ZEND_SEND_USER, &arg_node, nil)
		opline.GetOp2().SetNum(i)
		opline.GetResult().SetVar(uint32(ZendIntptrT(ZEND_CALL_ARG(nil, i))))
	}
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileAssert(result *Znode, args *ZendAstList, name *ZendString, fbc *ZendFunction) {
	if ExecutorGlobals.GetAssertions() >= 0 {
		var name_node Znode
		var opline *ZendOp
		var check_op_number uint32 = GetNextOpNumber()
		ZendEmitOp(nil, ZEND_ASSERT_CHECK, nil, nil)
		if fbc != nil && FbcIsFinalized(fbc) != 0 {
			name_node.SetOpType(IS_CONST)
			ZVAL_STR_COPY(&name_node.u.constant, name)
			opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
		} else {
			opline = ZendEmitOp(nil, ZEND_INIT_NS_FCALL_BY_NAME, nil, nil)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name))
		}
		opline.GetResult().SetNum(ZendAllocCacheSlot())
		if args.GetChildren() == 1 && (args.GetChild()[0].GetKind() != ZEND_AST_ZVAL || Z_TYPE_P(ZendAstGetZval(args.GetChild()[0])) != IS_STRING) {

			/* add "assert(condition) as assertion message */

			ZendAstListAdd((*ZendAst)(args), ZendAstCreateZvalFromStr(ZendAstExport("assert(", args.GetChild()[0], ")")))

			/* add "assert(condition) as assertion message */

		}
		ZendCompileCallCommon(result, (*ZendAst)(args), fbc)
		opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[check_op_number]
		opline.GetOp2().SetOplineNum(GetNextOpNumber())
		opline.SetResultType(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetResult(result.GetOp())
		}
	} else {
		if fbc == nil {
			ZendStringReleaseEx(name, 0)
		}
		result.SetOpType(IS_CONST)
		ZVAL_TRUE(&result.u.constant)
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
	if ZendHashNumElements(Z_ARRVAL(array.GetConstant())) > 0 {
		var ok ZendBool = 1
		var val *Zval
		var tmp Zval
		var src *HashTable = Z_ARRVAL(array.GetConstant())
		var dst *HashTable = ZendNewArray(ZendHashNumElements(src))
		ZVAL_TRUE(&tmp)
		if strict != 0 {
			for {
				var __ht *HashTable = src
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					val = _z
					if Z_TYPE_P(val) == IS_STRING {
						ZendHashAdd(dst, Z_STR_P(val), &tmp)
					} else if Z_TYPE_P(val) == IS_LONG {
						ZendHashIndexAdd(dst, Z_LVAL_P(val), &tmp)
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

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					val = _z
					if Z_TYPE_P(val) != IS_STRING || IsNumericString(Z_STRVAL_P(val), Z_STRLEN_P(val), nil, nil, 0) != 0 {
						ZendArrayDestroy(dst)
						ok = 0
						break
					}
					ZendHashAdd(dst, Z_STR_P(val), &tmp)
				}
				break
			}
		}
		ZendArrayDestroy(src)
		if ok == 0 {
			return FAILURE
		}
		Z_ARRVAL(array.GetConstant()) = dst
	}
	array.SetOpType(IS_CONST)
	ZendCompileExpr(&needly, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, ZEND_IN_ARRAY, &needly, &array)
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
	opline = ZendEmitOpTmp(result, ZEND_COUNT, &arg_node, nil)
	opline.SetExtendedValue(ZendStringEqualsLiteral(lcname, "sizeof"))
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGetClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_GET_CLASS, nil, nil)
	} else {
		var arg_node Znode
		if args.GetChildren() != 1 {
			return FAILURE
		}
		ZendCompileExpr(&arg_node, args.GetChild()[0])
		ZendEmitOpTmp(result, ZEND_GET_CLASS, &arg_node, nil)
	}
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGetCalledClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() != 0 {
		return FAILURE
	}
	ZendEmitOpTmp(result, ZEND_GET_CALLED_CLASS, nil, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncGettype(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if args.GetChildren() != 1 {
		return FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	ZendEmitOpTmp(result, ZEND_GET_TYPE, &arg_node, nil)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncNumArgs(result *Znode, args *ZendAstList) int {
	if CompilerGlobals.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_NUM_ARGS, nil, nil)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ZendCompileFuncGetArgs(result *Znode, args *ZendAstList) int {
	if CompilerGlobals.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, nil, nil)
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
	ZendEmitOpTmp(result, ZEND_ARRAY_KEY_EXISTS, &needle, &subject)
	return SUCCESS
}

/* }}} */

func ZendCompileFuncArraySlice(result *Znode, args *ZendAstList) int {
	if CompilerGlobals.GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 2 && args.GetChild()[0].GetKind() == ZEND_AST_CALL && args.GetChild()[0].GetChild()[0].GetKind() == ZEND_AST_ZVAL && Z_TYPE_P(ZendAstGetZval(args.GetChild()[0].GetChild()[0])) == IS_STRING && args.GetChild()[0].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST && args.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
		var orig_name *ZendString = ZendAstGetStr(args.GetChild()[0].GetChild()[0])
		var is_fully_qualified ZendBool
		var name *ZendString = ZendResolveFunctionName(orig_name, args.GetChild()[0].GetChild()[0].GetAttr(), &is_fully_qualified)
		var list *ZendAstList = ZendAstGetList(args.GetChild()[0].GetChild()[1])
		var zv *Zval = ZendAstGetZval(args.GetChild()[1])
		var first Znode
		if ZendStringEqualsLiteralCi(name, "func_get_args") && list.GetChildren() == 0 && Z_TYPE_P(zv) == IS_LONG && Z_LVAL_P(zv) >= 0 {
			first.SetOpType(IS_CONST)
			ZVAL_LONG(&first.u.constant, Z_LVAL_P(zv))
			ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, &first, nil)
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
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_NO_BUILTINS) != 0 {
		return FAILURE
	}
	if ZendArgsContainUnpack(args) != 0 {
		return FAILURE
	}
	if ZendStringEqualsLiteral(lcname, "strlen") {
		return ZendCompileFuncStrlen(result, args)
	} else if ZendStringEqualsLiteral(lcname, "is_null") {
		return ZendCompileFuncTypecheck(result, args, IS_NULL)
	} else if ZendStringEqualsLiteral(lcname, "is_bool") {
		return ZendCompileFuncTypecheck(result, args, _IS_BOOL)
	} else if ZendStringEqualsLiteral(lcname, "is_long") || ZendStringEqualsLiteral(lcname, "is_int") || ZendStringEqualsLiteral(lcname, "is_integer") {
		return ZendCompileFuncTypecheck(result, args, IS_LONG)
	} else if ZendStringEqualsLiteral(lcname, "is_float") || ZendStringEqualsLiteral(lcname, "is_double") {
		return ZendCompileFuncTypecheck(result, args, IS_DOUBLE)
	} else if ZendStringEqualsLiteral(lcname, "is_string") {
		return ZendCompileFuncTypecheck(result, args, IS_STRING)
	} else if ZendStringEqualsLiteral(lcname, "is_array") {
		return ZendCompileFuncTypecheck(result, args, IS_ARRAY)
	} else if ZendStringEqualsLiteral(lcname, "is_object") {
		return ZendCompileFuncTypecheck(result, args, IS_OBJECT)
	} else if ZendStringEqualsLiteral(lcname, "is_resource") {
		return ZendCompileFuncTypecheck(result, args, IS_RESOURCE)
	} else if ZendStringEqualsLiteral(lcname, "boolval") {
		return ZendCompileFuncCast(result, args, _IS_BOOL)
	} else if ZendStringEqualsLiteral(lcname, "intval") {
		return ZendCompileFuncCast(result, args, IS_LONG)
	} else if ZendStringEqualsLiteral(lcname, "floatval") || ZendStringEqualsLiteral(lcname, "doubleval") {
		return ZendCompileFuncCast(result, args, IS_DOUBLE)
	} else if ZendStringEqualsLiteral(lcname, "strval") {
		return ZendCompileFuncCast(result, args, IS_STRING)
	} else if ZendStringEqualsLiteral(lcname, "defined") {
		return ZendCompileFuncDefined(result, args)
	} else if ZendStringEqualsLiteral(lcname, "chr") && type_ == BP_VAR_R {
		return ZendCompileFuncChr(result, args)
	} else if ZendStringEqualsLiteral(lcname, "ord") && type_ == BP_VAR_R {
		return ZendCompileFuncOrd(result, args)
	} else if ZendStringEqualsLiteral(lcname, "call_user_func_array") {
		return ZendCompileFuncCufa(result, args, lcname)
	} else if ZendStringEqualsLiteral(lcname, "call_user_func") {
		return ZendCompileFuncCuf(result, args, lcname)
	} else if ZendStringEqualsLiteral(lcname, "in_array") {
		return ZendCompileFuncInArray(result, args)
	} else if ZendStringEqualsLiteral(lcname, "count") || ZendStringEqualsLiteral(lcname, "sizeof") {
		return ZendCompileFuncCount(result, args, lcname)
	} else if ZendStringEqualsLiteral(lcname, "get_class") {
		return ZendCompileFuncGetClass(result, args)
	} else if ZendStringEqualsLiteral(lcname, "get_called_class") {
		return ZendCompileFuncGetCalledClass(result, args)
	} else if ZendStringEqualsLiteral(lcname, "gettype") {
		return ZendCompileFuncGettype(result, args)
	} else if ZendStringEqualsLiteral(lcname, "func_num_args") {
		return ZendCompileFuncNumArgs(result, args)
	} else if ZendStringEqualsLiteral(lcname, "func_get_args") {
		return ZendCompileFuncGetArgs(result, args)
	} else if ZendStringEqualsLiteral(lcname, "array_slice") {
		return ZendCompileFuncArraySlice(result, args)
	} else if ZendStringEqualsLiteral(lcname, "array_key_exists") {
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
	if name_ast.GetKind() != ZEND_AST_ZVAL || Z_TYPE_P(ZendAstGetZval(name_ast)) != IS_STRING {
		ZendCompileExpr(&name_node, name_ast)
		ZendCompileDynamicCall(result, &name_node, args_ast)
		return
	}
	var runtime_resolution ZendBool = ZendCompileFunctionName(&name_node, name_ast)
	if runtime_resolution != 0 {
		if ZendStringEqualsLiteralCi(ZendAstGetStr(name_ast), "assert") {
			ZendCompileAssert(result, ZendAstGetList(args_ast), Z_STR(name_node.GetConstant()), nil)
		} else {
			ZendCompileNsCall(result, &name_node, args_ast)
		}
		return
	}
	var name *Zval = &name_node.u.constant
	var lcname *ZendString
	var fbc *ZendFunction
	var opline *ZendOp
	lcname = ZendStringTolower(Z_STR_P(name))
	fbc = ZendHashFindPtr(CompilerGlobals.GetFunctionTable(), lcname)

	/* Special assert() handling should apply independently of compiler flags. */

	if fbc != nil && ZendStringEqualsLiteral(lcname, "assert") {
		ZendCompileAssert(result, ZendAstGetList(args_ast), lcname, fbc)
		ZendStringRelease(lcname)
		ZvalPtrDtor(&name_node.u.constant)
		return
	}
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CompilerGlobals.GetActiveOpArray().GetFilename() {
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
	ZVAL_NEW_STR(&name_node.u.constant, lcname)
	opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
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
		obj_node.SetOpType(IS_UNUSED)
		CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_USES_THIS)
	} else {
		ZendCompileExpr(&obj_node, obj_ast)
	}
	ZendCompileExpr(&method_node, method_ast)
	opline = ZendEmitOp(nil, ZEND_INIT_METHOD_CALL, &obj_node, nil)
	if method_node.GetOpType() == IS_CONST {
		if Z_TYPE(method_node.GetConstant()) != IS_STRING {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Method name must be a string")
		}
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(Z_STR(method_node.GetConstant())))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		opline.SetOp2Type(&method_node.GetOpType())
		if &method_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&method_node).u.constant))
		} else {
			opline.SetOp2(&method_node.GetOp())
		}
	}

	/* Check if this calls a known method on $this */

	if opline.GetOp1Type() == IS_UNUSED && opline.GetOp2Type() == IS_CONST && CompilerGlobals.GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
		var lcname *ZendString = Z_STR_P(CT_CONSTANT(opline.GetOp2()) + 1)
		fbc = ZendHashFindPtr(&(CompilerGlobals.GetActiveClassEntry()).function_table, lcname)

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

		if fbc != nil && (fbc.GetFnFlags()&(ZEND_ACC_PRIVATE|ZEND_ACC_FINAL)) == 0 {
			fbc = nil
		}

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

	}
	ZendCompileCallCommon(result, args_ast, fbc)
}

/* }}} */

func ZendIsConstructor(name *ZendString) ZendBool {
	return ZendStringEqualsLiteralCi(name, ZEND_CONSTRUCTOR_FUNC_NAME)
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
	ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	ZendCompileExpr(&method_node, method_ast)
	if method_node.GetOpType() == IS_CONST {
		var name *Zval = &method_node.u.constant
		if Z_TYPE_P(name) != IS_STRING {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Method name must be a string")
		}
		if ZendIsConstructor(Z_STR_P(name)) != 0 {
			ZvalPtrDtor(name)
			method_node.SetOpType(IS_UNUSED)
		}
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
	ZendSetClassNameOp1(opline, &class_node)
	if method_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(Z_STR(method_node.GetConstant())))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		if opline.GetOp1Type() == IS_CONST {
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
		opline.SetOp2Type(&method_node.GetOpType())
		if &method_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&method_node).u.constant))
		} else {
			opline.SetOp2(&method_node.GetOp())
		}
	}

	/* Check if we already know which method we're calling */

	if opline.GetOp2Type() == IS_CONST {
		var ce *ZendClassEntry = nil
		if opline.GetOp1Type() == IS_CONST {
			var lcname *ZendString = Z_STR_P(CT_CONSTANT(opline.GetOp1()) + 1)
			ce = ZendHashFindPtr(CompilerGlobals.GetClassTable(), lcname)
			if ce == nil && CompilerGlobals.GetActiveClassEntry() != nil && ZendStringEqualsCi(CompilerGlobals.GetActiveClassEntry().GetName(), lcname) {
				ce = CompilerGlobals.GetActiveClassEntry()
			}
		} else if opline.GetOp1Type() == IS_UNUSED && (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() != 0 {
			ce = CompilerGlobals.GetActiveClassEntry()
		}
		if ce != nil {
			var lcname *ZendString = Z_STR_P(CT_CONSTANT(opline.GetOp2()) + 1)
			fbc = ZendHashFindPtr(&ce.function_table, lcname)
			if fbc != nil && (fbc.GetFnFlags()&ZEND_ACC_PUBLIC) == 0 {
				if ce != CompilerGlobals.GetActiveClassEntry() && ((fbc.GetFnFlags()&ZEND_ACC_PRIVATE) != 0 || (fbc.GetScope().GetCeFlags()&ZEND_ACC_LINKED) == 0 || CompilerGlobals.GetActiveClassEntry() != nil && (CompilerGlobals.GetActiveClassEntry().GetCeFlags()&ZEND_ACC_LINKED) == 0 || ZendCheckProtected(ZendGetFunctionRootClass(fbc), CompilerGlobals.GetActiveClassEntry()) == 0) {

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
		ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	}
	opline = ZendEmitOp(result, ZEND_NEW, nil, nil)
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp1Type(IS_CONST)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(Z_STR(class_node.GetConstant())))
		opline.GetOp2().SetNum(ZendAllocCacheSlot())
	} else {
		opline.SetOp1Type(&class_node.GetOpType())
		if &class_node.GetOpType() == IS_CONST {
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
	ZendEmitOpTmp(result, ZEND_CLONE, &obj_node, nil)
}

/* }}} */

func ZendCompileGlobalVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var name_ast *ZendAst = var_ast.GetChild()[0]
	var name_node Znode
	var result Znode
	ZendCompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == IS_CONST {
		ConvertToString(&name_node.u.constant)
	}
	if IsThisFetch(var_ast) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use $this as global variable")
	} else if ZendTryCompileCv(&result, var_ast) == SUCCESS {
		var opline *ZendOp = ZendEmitOp(nil, ZEND_BIND_GLOBAL, &result, &name_node)
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {

		/* name_ast should be evaluated only. FETCH_GLOBAL_LOCK instructs FETCH_W
		 * to not free the name_node operand, so it can be reused in the following
		 * ASSIGN_REF, which then frees it. */

		var opline *ZendOp = ZendEmitOp(&result, ZEND_FETCH_W, &name_node, nil)
		opline.SetExtendedValue(ZEND_FETCH_GLOBAL_LOCK)
		if name_node.GetOpType() == IS_CONST {
			ZendStringAddref(Z_STR(name_node.GetConstant()))
		}
		ZendEmitAssignRefZnode(ZendAstCreate(ZEND_AST_VAR, ZendAstCreateZnode(&name_node)), &result)
	}
}

/* }}} */

func ZendCompileStaticVarCommon(var_name *ZendString, value *Zval, mode uint32) {
	var opline *ZendOp
	if CompilerGlobals.GetActiveOpArray().GetStaticVariables() == nil {
		if CompilerGlobals.GetActiveOpArray().GetScope() != nil {
			CompilerGlobals.GetActiveOpArray().GetScope().SetCeFlags(CompilerGlobals.GetActiveOpArray().GetScope().GetCeFlags() | ZEND_HAS_STATIC_IN_METHODS)
		}
		CompilerGlobals.GetActiveOpArray().SetStaticVariables(ZendNewArray(8))
	}
	value = ZendHashUpdate(CompilerGlobals.GetActiveOpArray().GetStaticVariables(), var_name, value)
	if ZendStringEqualsLiteral(var_name, "this") {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use $this as static variable")
	}
	opline = ZendEmitOp(nil, ZEND_BIND_STATIC, nil, nil)
	opline.SetOp1Type(IS_CV)
	opline.GetOp1().SetVar(LookupCv(var_name))
	opline.SetExtendedValue(uint32_t((*byte)(value-(*byte)(CompilerGlobals.GetActiveOpArray().GetStaticVariables().GetArData()))) | mode)
}

/* }}} */

func ZendCompileStaticVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var value_ast *ZendAst = ast.GetChild()[1]
	var value_zv Zval
	if value_ast != nil {
		ZendConstExprToZval(&value_zv, value_ast)
	} else {
		ZVAL_NULL(&value_zv)
	}
	ZendCompileStaticVarCommon(ZendAstGetStr(var_ast), &value_zv, ZEND_BIND_REF)
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
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot unset $this")
		} else if ZendTryCompileCv(&var_node, var_ast) == SUCCESS {
			opline = ZendEmitOp(nil, ZEND_UNSET_CV, &var_node, nil)
		} else {
			opline = ZendCompileSimpleVarNoCv(nil, var_ast, BP_VAR_UNSET, 0)
			opline.SetOpcode(ZEND_UNSET_VAR)
		}
		return
	case ZEND_AST_DIM:
		opline = ZendCompileDim(nil, var_ast, BP_VAR_UNSET)
		opline.SetOpcode(ZEND_UNSET_DIM)
		return
	case ZEND_AST_PROP:
		opline = ZendCompileProp(nil, var_ast, BP_VAR_UNSET, 0)
		opline.SetOpcode(ZEND_UNSET_OBJ)
		return
	case ZEND_AST_STATIC_PROP:
		opline = ZendCompileStaticProp(nil, var_ast, BP_VAR_UNSET, 0, 0)
		opline.SetOpcode(ZEND_UNSET_STATIC_PROP)
		return
	default:
		break
	}
}

/* }}} */

func ZendHandleLoopsAndFinallyEx(depth ZendLong, return_value *Znode) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = ZendStackTop(&(CompilerGlobals.GetLoopVarStack()))
	if loop_var == nil {
		return 1
	}
	base = ZendStackBase(&(CompilerGlobals.GetLoopVarStack()))
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == ZEND_FAST_CALL {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_FAST_CALL)
			opline.SetResultType(IS_TMP_VAR)
			opline.GetResult().SetVar(loop_var.GetVarNum())
			if return_value != nil {
				opline.SetOp2Type(return_value.GetOpType())
				if return_value.GetOpType() == IS_CONST {
					opline.GetOp2().SetConstant(ZendAddLiteral(&return_value.u.constant))
				} else {
					opline.SetOp2(return_value.GetOp())
				}
			}
			opline.GetOp1().SetNum(loop_var.GetTryCatchOffset())
		} else if loop_var.GetOpcode() == ZEND_DISCARD_EXCEPTION {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_DISCARD_EXCEPTION)
			opline.SetOp1Type(IS_TMP_VAR)
			opline.GetOp1().SetVar(loop_var.GetVarNum())
		} else if loop_var.GetOpcode() == ZEND_RETURN {

			/* Stack separator */

			break

			/* Stack separator */

		} else if depth <= 1 {
			return 1
		} else if loop_var.GetOpcode() == ZEND_NOP {

			/* Loop doesn't have freeable variable */

			depth--

			/* Loop doesn't have freeable variable */

		} else {
			var opline *ZendOp
			ZEND_ASSERT((loop_var.GetVarType() & (IS_VAR | IS_TMP_VAR)) != 0)
			opline = GetNextOp()
			opline.SetOpcode(loop_var.GetOpcode())
			opline.SetOp1Type(loop_var.GetVarType())
			opline.GetOp1().SetVar(loop_var.GetVarNum())
			opline.SetExtendedValue(ZEND_FREE_ON_RETURN)
			depth--
		}
	}
	return depth == 0
}

/* }}} */

func ZendHandleLoopsAndFinally(return_value *Znode) int {
	return ZendHandleLoopsAndFinallyEx(ZendStackCount(&(CompilerGlobals.GetLoopVarStack()))+1, return_value)
}

/* }}} */

func ZendHasFinallyEx(depth ZendLong) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = ZendStackTop(&(CompilerGlobals.GetLoopVarStack()))
	if loop_var == nil {
		return 0
	}
	base = ZendStackBase(&(CompilerGlobals.GetLoopVarStack()))
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == ZEND_FAST_CALL {
			return 1
		} else if loop_var.GetOpcode() == ZEND_DISCARD_EXCEPTION {

		} else if loop_var.GetOpcode() == ZEND_RETURN {

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
	return ZendHasFinallyEx(ZendStackCount(&(CompilerGlobals.GetLoopVarStack())) + 1)
}

/* }}} */

func ZendCompileReturn(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var is_generator ZendBool = (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_GENERATOR) != 0
	var by_ref ZendBool = (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0
	var expr_node Znode
	var opline *ZendOp
	if is_generator != 0 {

		/* For generators the by-ref flag refers to yields, not returns */

		by_ref = 0

		/* For generators the by-ref flag refers to yields, not returns */

	}
	if expr_ast == nil {
		expr_node.SetOpType(IS_CONST)
		ZVAL_NULL(&expr_node.u.constant)
	} else if by_ref != 0 && ZendIsVariable(expr_ast) != 0 {
		ZendCompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
	}
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags()&ZEND_ACC_HAS_FINALLY_BLOCK) != 0 && (expr_node.GetOpType() == IS_CV || by_ref != 0 && expr_node.GetOpType() == IS_VAR) && ZendHasFinally() != 0 {

		/* Copy return value into temporary VAR to avoid modification in finally code */

		if by_ref != 0 {
			ZendEmitOp(&expr_node, ZEND_MAKE_REF, &expr_node, nil)
		} else {
			ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &expr_node, nil)
		}

		/* Copy return value into temporary VAR to avoid modification in finally code */

	}

	/* Generator return types are handled separately */

	if is_generator == 0 && (CompilerGlobals.GetActiveOpArray().GetFnFlags()&ZEND_ACC_HAS_RETURN_TYPE) != 0 {
		ZendEmitReturnTypeCheck(b.Cond(expr_ast != nil, &expr_node, nil), CompilerGlobals.GetActiveOpArray().GetArgInfo()-1, 0)
	}
	ZendHandleLoopsAndFinally(b.Cond((expr_node.GetOpType()&(IS_TMP_VAR|IS_VAR)) != 0, &expr_node, nil))
	opline = ZendEmitOp(nil, b.Cond(by_ref != 0, ZEND_RETURN_BY_REF, ZEND_RETURN), &expr_node, nil)
	if by_ref != 0 && expr_ast != nil {
		if ZendIsCall(expr_ast) != 0 {
			opline.SetExtendedValue(ZEND_RETURNS_FUNCTION)
		} else if ZendIsVariable(expr_ast) == 0 {
			opline.SetExtendedValue(ZEND_RETURNS_VALUE)
		}
	}
}

/* }}} */

func ZendCompileEcho(ast *ZendAst) {
	var opline *ZendOp
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, ZEND_ECHO, &expr_node, nil)
	opline.SetExtendedValue(0)
}

/* }}} */

func ZendCompileThrow(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	ZendEmitOp(nil, ZEND_THROW, &expr_node, nil)
}

/* }}} */

func ZendCompileBreakContinue(ast *ZendAst) {
	var depth_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	var depth ZendLong
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_BREAK || ast.GetKind() == ZEND_AST_CONTINUE)
	if depth_ast != nil {
		var depth_zv *Zval
		if depth_ast.GetKind() != ZEND_AST_ZVAL {
			ZendErrorNoreturn(E_COMPILE_ERROR, "'%s' operator with non-integer operand "+"is no longer supported", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth_zv = ZendAstGetZval(depth_ast)
		if Z_TYPE_P(depth_zv) != IS_LONG || Z_LVAL_P(depth_zv) < 1 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "'%s' operator accepts only positive integers", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth = Z_LVAL_P(depth_zv)
	} else {
		depth = 1
	}
	if CompilerGlobals.GetContext().GetCurrentBrkCont() == -1 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "'%s' not in the 'loop' or 'switch' context", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
	} else {
		if ZendHandleLoopsAndFinallyEx(depth, nil) == 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot '%s' "+ZEND_LONG_FMT+" level%s", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"), depth, b.Cond(depth == 1, "", "s"))
		}
	}
	if ast.GetKind() == ZEND_AST_CONTINUE {
		var d int
		var cur int = CompilerGlobals.GetContext().GetCurrentBrkCont()
		for d = depth - 1; d > 0; d-- {
			cur = CompilerGlobals.GetContext().GetBrkContArray()[cur].GetParent()
			ZEND_ASSERT(cur != -1)
		}
		if CompilerGlobals.GetContext().GetBrkContArray()[cur].GetIsSwitch() != 0 {
			if depth == 1 {
				ZendError(E_WARNING, "\"continue\" targeting switch is equivalent to \"break\". "+"Did you mean to use \"continue "+ZEND_LONG_FMT+"\"?", depth+1)
			} else {
				ZendError(E_WARNING, "\"continue "+ZEND_LONG_FMT+"\" targeting switch is equivalent to \"break "+ZEND_LONG_FMT+"\". "+"Did you mean to use \"continue "+ZEND_LONG_FMT+"\"?", depth, depth, depth+1)
			}
		}
	}
	opline = ZendEmitOp(nil, b.Cond(ast.GetKind() == ZEND_AST_BREAK, ZEND_BRK, ZEND_CONT), nil, nil)
	opline.GetOp1().SetNum(CompilerGlobals.GetContext().GetCurrentBrkCont())
	opline.GetOp2().SetNum(depth)
}

/* }}} */

func ZendResolveGotoLabel(op_array *ZendOpArray, opline *ZendOp) {
	var dest *ZendLabel
	var current int
	var remove_oplines int = opline.GetOp1().GetNum()
	var label *Zval
	var opnum uint32 = opline - op_array.GetOpcodes()
	label = CT_CONSTANT_EX(op_array, opline.GetOp2().GetConstant())
	if CompilerGlobals.GetContext().GetLabels() == nil || b.Assign(&dest, ZendHashFindPtr(CompilerGlobals.GetContext().GetLabels(), Z_STR_P(label))) == nil {
		CompilerGlobals.SetInCompilation(1)
		CompilerGlobals.SetActiveOpArray(op_array)
		CompilerGlobals.SetZendLineno(opline.GetLineno())
		ZendErrorNoreturn(E_COMPILE_ERROR, "'goto' to undefined label '%s'", Z_STRVAL_P(label))
	}
	ZvalPtrDtorStr(label)
	ZVAL_NULL(label)
	current = opline.GetExtendedValue()
	for ; current != dest.GetBrkCont(); current = CompilerGlobals.GetContext().GetBrkContArray()[current].GetParent() {
		if current == -1 {
			CompilerGlobals.SetInCompilation(1)
			CompilerGlobals.SetActiveOpArray(op_array)
			CompilerGlobals.SetZendLineno(opline.GetLineno())
			ZendErrorNoreturn(E_COMPILE_ERROR, "'goto' into loop or switch statement is disallowed")
		}
		if CompilerGlobals.GetContext().GetBrkContArray()[current].GetStart() >= 0 {
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
	opline.SetOpcode(ZEND_JMP)
	opline.GetOp1().SetOplineNum(dest.GetOplineNum())
	opline.SetExtendedValue(0)
	opline.SetOp1Type(IS_UNUSED)
	opline.SetOp2Type(IS_UNUSED)
	opline.SetResultType(IS_UNUSED)
	ZEND_ASSERT(remove_oplines >= 0)
	for b.PostDec(&remove_oplines) {
		opline--
		MAKE_NOP(opline)
		ZEND_VM_SET_OPCODE_HANDLER(opline)
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
	opline = ZendEmitOp(nil, ZEND_GOTO, nil, &label_node)
	opline.GetOp1().SetNum(GetNextOpNumber() - opnum_start - 1)
	opline.SetExtendedValue(CompilerGlobals.GetContext().GetCurrentBrkCont())
}

/* }}} */

func ZendCompileLabel(ast *ZendAst) {
	var label *ZendString = ZendAstGetStr(ast.GetChild()[0])
	var dest ZendLabel
	if CompilerGlobals.GetContext().GetLabels() == nil {
		ALLOC_HASHTABLE(CompilerGlobals.GetContext().GetLabels())
		ZendHashInit(CompilerGlobals.GetContext().GetLabels(), 8, nil, LabelPtrDtor, 0)
	}
	dest.SetBrkCont(CompilerGlobals.GetContext().GetCurrentBrkCont())
	dest.SetOplineNum(GetNextOpNumber())
	if !(ZendHashAddMem(CompilerGlobals.GetContext().GetLabels(), label, &dest, b.SizeOf("zend_label"))) {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Label '%s' already defined", ZSTR_VAL(label))
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
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendUpdateJumpTarget(opnum_jmp, opnum_cond)
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}

/* }}} */

func ZendCompileDoWhile(ast *ZendAst) {
	var stmt_ast *ZendAst = ast.GetChild()[0]
	var cond_ast *ZendAst = ast.GetChild()[1]
	var cond_node Znode
	var opnum_start uint32
	var opnum_cond uint32
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}

/* }}} */

func ZendCompileExprList(result *Znode, ast *ZendAst) {
	var list *ZendAstList
	var i uint32
	result.SetOpType(IS_CONST)
	ZVAL_TRUE(&result.u.constant)
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
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_loop = GetNextOpNumber()
	ZendCompileExprList(&result, loop_ast)
	ZendDoFree(&result)
	ZendUpdateJumpTargetToNext(opnum_jmp)
	ZendCompileExprList(&result, cond_ast)
	ZendDoExtendedStmt()
	ZendEmitCondJump(ZEND_JMPNZ, &result, opnum_start)
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
			ZendErrorNoreturn(E_COMPILE_ERROR, "Key element cannot be a reference")
		}
		if key_ast.GetKind() == ZEND_AST_ARRAY {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use list as key element")
		}
	}
	if by_ref != 0 {
		value_ast = value_ast.GetChild()[0]
	}
	if value_ast.GetKind() == ZEND_AST_ARRAY && ZendPropagateListRefs(value_ast) != 0 {
		by_ref = 1
	}
	if by_ref != 0 && is_variable != 0 {
		ZendCompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
	}
	if by_ref != 0 {
		ZendSeparateIfCallAndWrite(&expr_node, expr_ast, BP_VAR_W)
	}
	opnum_reset = GetNextOpNumber()
	opline = ZendEmitOp(&reset_node, b.Cond(by_ref != 0, ZEND_FE_RESET_RW, ZEND_FE_RESET_R), &expr_node, nil)
	ZendBeginLoop(ZEND_FE_FREE, &reset_node, 0)
	opnum_fetch = GetNextOpNumber()
	opline = ZendEmitOp(nil, b.Cond(by_ref != 0, ZEND_FE_FETCH_RW, ZEND_FE_FETCH_R), &reset_node, nil)
	if IsThisFetch(value_ast) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign $this")
	} else if value_ast.GetKind() == ZEND_AST_VAR && ZendTryCompileCv(&value_node, value_ast) == SUCCESS {
		opline.SetOp2Type(&value_node.GetOpType())
		if &value_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(&(&value_node).u.constant))
		} else {
			opline.SetOp2(&value_node.GetOp())
		}
	} else {
		opline.SetOp2Type(IS_VAR)
		opline.GetOp2().SetVar(GetTemporaryVariable())
		&value_node.SetOpType(opline.GetOp2Type())
		if &value_node.GetOpType() == IS_CONST {
			ZVAL_COPY_VALUE(&(&value_node).u.constant, CT_CONSTANT(opline.GetOp2()))
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
		opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_fetch]
		ZendMakeTmpResult(&key_node, opline)
		ZendEmitAssignZnode(key_ast, &key_node)
	}
	ZendCompileStmt(stmt_ast)

	/* Place JMP and FE_FREE on the line where foreach starts. It would be
	 * better to use the end line, but this information is not available
	 * currently. */

	CompilerGlobals.SetZendLineno(ast.GetLineno())
	ZendEmitJump(opnum_fetch)
	opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_reset]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
	opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_fetch]
	opline.SetExtendedValue(GetNextOpNumber())
	ZendEndLoop(opnum_fetch, &reset_node)
	opline = ZendEmitOp(nil, ZEND_FE_FREE, &reset_node, nil)
}

/* }}} */

func ZendCompileIf(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var jmp_opnums *uint32 = nil
	if list.GetChildren() > 1 {
		jmp_opnums = SafeEmalloc(b.SizeOf("uint32_t"), list.GetChildren()-1, 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var cond_ast *ZendAst = elem_ast.GetChild()[0]
		var stmt_ast *ZendAst = elem_ast.GetChild()[1]
		if cond_ast != nil {
			var cond_node Znode
			var opnum_jmpz uint32
			ZendCompileExpr(&cond_node, cond_ast)
			opnum_jmpz = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
			ZendCompileStmt(stmt_ast)
			if i != list.GetChildren()-1 {
				jmp_opnums[i] = ZendEmitJump(0)
			}
			ZendUpdateJumpTargetToNext(opnum_jmpz)
		} else {

			/* "else" can only occur as last element. */

			ZEND_ASSERT(i == list.GetChildren()-1)
			ZendCompileStmt(stmt_ast)
		}
	}
	if list.GetChildren() > 1 {
		for i = 0; i < list.GetChildren()-1; i++ {
			ZendUpdateJumpTargetToNext(jmp_opnums[i])
		}
		Efree(jmp_opnums)
	}
}

/* }}} */

func DetermineSwitchJumptableType(cases *ZendAstList) ZendUchar {
	var i uint32
	var common_type ZendUchar = IS_UNDEF
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

			return IS_UNDEF

			/* Non-constant case */

		}
		cond_zv = ZendAstGetZval(case_ast.GetChild()[0])
		if Z_TYPE_P(cond_zv) != IS_LONG && Z_TYPE_P(cond_zv) != IS_STRING {

			/* We only optimize switched on integers and strings */

			return IS_UNDEF

			/* We only optimize switched on integers and strings */

		}
		if common_type == IS_UNDEF {
			common_type = Z_TYPE_P(cond_zv)
		} else if common_type != Z_TYPE_P(cond_zv) {

			/* Non-uniform case types */

			return IS_UNDEF

			/* Non-uniform case types */

		}
		if Z_TYPE_P(cond_zv) == IS_STRING && IsNumericString(Z_STRVAL_P(cond_zv), Z_STRLEN_P(cond_zv), nil, nil, 0) != 0 {

			/* Numeric strings cannot be compared with a simple hash lookup */

			return IS_UNDEF

			/* Numeric strings cannot be compared with a simple hash lookup */

		}
	}
	return common_type
}
func ShouldUseJumptable(cases *ZendAstList, jumptable_type ZendUchar) ZendBool {
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_NO_JUMPTABLES) != 0 {
		return 0
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */

	if jumptable_type == IS_LONG {
		return cases.GetChildren() >= 5
	} else {
		ZEND_ASSERT(jumptable_type == IS_STRING)
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
	var opnum_switch uint32 = uint32_t - 1
	var jumptable_type ZendUchar
	var jumptable *HashTable = nil
	ZendCompileExpr(&expr_node, expr_ast)
	ZendBeginLoop(ZEND_FREE, &expr_node, 1)
	case_node.SetOpType(IS_TMP_VAR)
	case_node.GetOp().SetVar(GetTemporaryVariable())
	jumptable_type = DetermineSwitchJumptableType(cases)
	if jumptable_type != IS_UNDEF && ShouldUseJumptable(cases, jumptable_type) != 0 {
		var jumptable_op Znode
		ALLOC_HASHTABLE(jumptable)
		ZendHashInit(jumptable, cases.GetChildren(), nil, nil, 0)
		jumptable_op.SetOpType(IS_CONST)
		ZVAL_ARR(&jumptable_op.u.constant, jumptable)
		opline = ZendEmitOp(nil, b.Cond(jumptable_type == IS_LONG, ZEND_SWITCH_LONG, ZEND_SWITCH_STRING), &expr_node, &jumptable_op)
		if opline.GetOp1Type() == IS_CONST {
			Z_TRY_ADDREF_P(CT_CONSTANT(opline.GetOp1()))
		}
		opnum_switch = opline - CompilerGlobals.GetActiveOpArray().GetOpcodes()
	}
	jmpnz_opnums = SafeEmalloc(b.SizeOf("uint32_t"), cases.GetChildren(), 0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast *ZendAst = case_ast.GetChild()[0]
		var cond_node Znode
		if cond_ast == nil {
			if has_default_case != 0 {
				CompilerGlobals.SetZendLineno(case_ast.GetLineno())
				ZendErrorNoreturn(E_COMPILE_ERROR, "Switch statements may only contain one default clause")
			}
			has_default_case = 1
			continue
		}
		ZendCompileExpr(&cond_node, cond_ast)
		if expr_node.GetOpType() == IS_CONST && Z_TYPE(expr_node.GetConstant()) == IS_FALSE {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
		} else if expr_node.GetOpType() == IS_CONST && Z_TYPE(expr_node.GetConstant()) == IS_TRUE {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &cond_node, 0)
		} else {
			opline = ZendEmitOp(nil, b.Cond((expr_node.GetOpType()&(IS_VAR|IS_TMP_VAR)) != 0, ZEND_CASE, ZEND_IS_EQUAL), &expr_node, &cond_node)
			opline.SetResultType(&case_node.GetOpType())
			if &case_node.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(&(&case_node).u.constant))
			} else {
				opline.SetResult(&case_node.GetOp())
			}
			if opline.GetOp1Type() == IS_CONST {
				Z_TRY_ADDREF_P(CT_CONSTANT(opline.GetOp1()))
			}
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &case_node, 0)
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
				ZVAL_LONG(&jmp_target, GetNextOpNumber())
				ZEND_ASSERT(Z_TYPE_P(cond_zv) == jumptable_type)
				if Z_TYPE_P(cond_zv) == IS_LONG {
					ZendHashIndexAdd(jumptable, Z_LVAL_P(cond_zv), &jmp_target)
				} else {
					ZEND_ASSERT(Z_TYPE_P(cond_zv) == IS_STRING)
					ZendHashAdd(jumptable, Z_STR_P(cond_zv), &jmp_target)
				}
			}
		} else {
			ZendUpdateJumpTargetToNext(opnum_default_jmp)
			if jumptable != nil {
				ZEND_ASSERT(opnum_switch != uint32_t-1)
				opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_switch]
				opline.SetExtendedValue(GetNextOpNumber())
			}
		}
		ZendCompileStmt(stmt_ast)
	}
	if has_default_case == 0 {
		ZendUpdateJumpTargetToNext(opnum_default_jmp)
		if jumptable != nil {
			opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_switch]
			opline.SetExtendedValue(GetNextOpNumber())
		}
	}
	ZendEndLoop(GetNextOpNumber(), &expr_node)
	if (expr_node.GetOpType() & (IS_VAR | IS_TMP_VAR)) != 0 {
		opline = ZendEmitOp(nil, ZEND_FREE, &expr_node, nil)
		opline.SetExtendedValue(ZEND_FREE_SWITCH)
	} else if expr_node.GetOpType() == IS_CONST {
		ZvalPtrDtorNogc(&expr_node.u.constant)
	}
	Efree(jmpnz_opnums)
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
	var jmp_opnums *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), catches.GetChildren(), 0)
	var orig_fast_call_var uint32 = CompilerGlobals.GetContext().GetFastCallVar()
	var orig_try_catch_offset uint32 = CompilerGlobals.GetContext().GetTryCatchOffset()
	if catches.GetChildren() == 0 && finally_ast == nil {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use try without catch or finally")
	}

	/* label: try { } must not be equal to try { label: } */

	if CompilerGlobals.GetContext().GetLabels() != nil {
		var label *ZendLabel
		for {
			var __ht *HashTable = CompilerGlobals.GetContext().GetLabels()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				label = Z_PTR_P(_z)
				if label.GetOplineNum() == GetNextOpNumber() {
					ZendEmitOp(nil, ZEND_NOP, nil, nil)
				}
				break
			}
			break
		}
	}
	try_catch_offset = ZendAddTryElement(GetNextOpNumber())
	if finally_ast != nil {
		var fast_call ZendLoopVar
		if (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_HAS_FINALLY_BLOCK) == 0 {
			CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_HAS_FINALLY_BLOCK)
		}
		CompilerGlobals.GetContext().SetFastCallVar(GetTemporaryVariable())

		/* Push FAST_CALL on unwind stack */

		fast_call.SetOpcode(ZEND_FAST_CALL)
		fast_call.SetVarType(IS_TMP_VAR)
		fast_call.SetVarNum(CompilerGlobals.GetContext().GetFastCallVar())
		fast_call.SetTryCatchOffset(try_catch_offset)
		ZendStackPush(&(CompilerGlobals.GetLoopVarStack()), &fast_call)
	}
	CompilerGlobals.GetContext().SetTryCatchOffset(try_catch_offset)
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
		var jmp_multicatch *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), classes.GetChildren()-1, 0)
		var opnum_catch uint32 = uint32_t - 1
		CompilerGlobals.SetZendLineno(catch_ast.GetLineno())
		for j = 0; j < classes.GetChildren(); j++ {
			var class_ast *ZendAst = classes.GetChild()[j]
			var is_last_class ZendBool = j+1 == classes.GetChildren()
			if ZendIsConstDefaultClassRef(class_ast) == 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Bad class name in the catch statement")
			}
			opnum_catch = GetNextOpNumber()
			if i == 0 && j == 0 {
				CompilerGlobals.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetCatchOp(opnum_catch)
			}
			opline = GetNextOp()
			opline.SetOpcode(ZEND_CATCH)
			opline.SetOp1Type(IS_CONST)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(ZendResolveClassNameAst(class_ast)))
			opline.SetExtendedValue(ZendAllocCacheSlot())
			if ZendStringEqualsLiteral(var_name, "this") {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign $this")
			}
			opline.SetResultType(IS_CV)
			opline.GetResult().SetVar(LookupCv(var_name))
			if is_last_catch != 0 && is_last_class != 0 {
				opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_LAST_CATCH)
			}
			if is_last_class == 0 {
				jmp_multicatch[j] = ZendEmitJump(0)
				opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_catch]
				opline.GetOp2().SetOplineNum(GetNextOpNumber())
			}
		}
		for j = 0; j < classes.GetChildren()-1; j++ {
			ZendUpdateJumpTargetToNext(jmp_multicatch[j])
		}
		Efree(jmp_multicatch)
		ZendCompileStmt(stmt_ast)
		if is_last_catch == 0 {
			jmp_opnums[i+1] = ZendEmitJump(0)
		}
		ZEND_ASSERT(opnum_catch != uint32_t-1 && "Should have at least one class")
		opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_catch]
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

		ZendStackDelTop(&(CompilerGlobals.GetLoopVarStack()))

		/* Push DISCARD_EXCEPTION on unwind stack */

		discard_exception.SetOpcode(ZEND_DISCARD_EXCEPTION)
		discard_exception.SetVarType(IS_TMP_VAR)
		discard_exception.SetVarNum(CompilerGlobals.GetContext().GetFastCallVar())
		ZendStackPush(&(CompilerGlobals.GetLoopVarStack()), &discard_exception)
		CompilerGlobals.SetZendLineno(finally_ast.GetLineno())
		opline = ZendEmitOp(nil, ZEND_FAST_CALL, nil, nil)
		opline.GetOp1().SetNum(try_catch_offset)
		opline.SetResultType(IS_TMP_VAR)
		opline.GetResult().SetVar(CompilerGlobals.GetContext().GetFastCallVar())
		ZendEmitOp(nil, ZEND_JMP, nil, nil)
		ZendCompileStmt(finally_ast)
		CompilerGlobals.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyOp(opnum_jmp + 1)
		CompilerGlobals.GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyEnd(GetNextOpNumber())
		opline = ZendEmitOp(nil, ZEND_FAST_RET, nil, nil)
		opline.SetOp1Type(IS_TMP_VAR)
		opline.GetOp1().SetVar(CompilerGlobals.GetContext().GetFastCallVar())
		opline.GetOp2().SetNum(orig_try_catch_offset)
		ZendUpdateJumpTargetToNext(opnum_jmp)
		CompilerGlobals.GetContext().SetFastCallVar(orig_fast_call_var)

		/* Pop DISCARD_EXCEPTION from unwind stack */

		ZendStackDelTop(&(CompilerGlobals.GetLoopVarStack()))

		/* Pop DISCARD_EXCEPTION from unwind stack */

	}
	CompilerGlobals.GetContext().SetTryCatchOffset(orig_try_catch_offset)
	Efree(jmp_opnums)
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
		if ZendStringEqualsLiteralCi(name, "encoding") {
			if value_ast.GetKind() != ZEND_AST_ZVAL {
				ZendThrowException(ZendCeCompileError, "Encoding must be a literal", 0)
				return 0
			}
			if CompilerGlobals.GetMultibyte() != 0 {
				var encoding_name *ZendString = ZvalGetString(ZendAstGetZval(value_ast))
				var new_encoding *ZendEncoding
				var old_encoding *ZendEncoding
				var old_input_filter ZendEncodingFilter
				CompilerGlobals.SetEncodingDeclared(1)
				new_encoding = ZendMultibyteFetchEncoding(ZSTR_VAL(encoding_name))
				if new_encoding == nil {
					ZendError(E_COMPILE_WARNING, "Unsupported encoding [%s]", ZSTR_VAL(encoding_name))
				} else {
					old_input_filter = LanguageScannerGlobals.GetInputFilter()
					old_encoding = LanguageScannerGlobals.GetScriptEncoding()
					ZendMultibyteSetFilter(new_encoding)

					/* need to re-scan if input filter changed */

					if old_input_filter != LanguageScannerGlobals.GetInputFilter() || old_input_filter != nil && new_encoding != old_encoding {
						ZendMultibyteYyinputAgain(old_input_filter, old_encoding)
					}

					/* need to re-scan if input filter changed */

				}
				ZendStringReleaseEx(encoding_name, 0)
			} else {
				ZendError(E_COMPILE_WARNING, "declare(encoding=...) ignored because "+"Zend multibyte feature is turned off by settings")
			}
		}
	}
	return 1
}

/* }}} */

func ZendDeclareIsFirstStatement(ast *ZendAst) int {
	var i uint32 = 0
	var file_ast *ZendAstList = ZendAstGetList(CompilerGlobals.GetAst())

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
	var orig_declarables ZendDeclarables = FC(declarables)
	var i uint32
	for i = 0; i < declares.GetChildren(); i++ {
		var declare_ast *ZendAst = declares.GetChild()[i]
		var name_ast *ZendAst = declare_ast.GetChild()[0]
		var value_ast *ZendAst = declare_ast.GetChild()[1]
		var name *ZendString = ZendAstGetStr(name_ast)
		if value_ast.GetKind() != ZEND_AST_ZVAL {
			ZendErrorNoreturn(E_COMPILE_ERROR, "declare(%s) value must be a literal", ZSTR_VAL(name))
		}
		if ZendStringEqualsLiteralCi(name, "ticks") {
			var value_zv Zval
			ZendConstExprToZval(&value_zv, value_ast)
			FC(declarables).ticks = ZvalGetLong(&value_zv)
			ZvalPtrDtorNogc(&value_zv)
		} else if ZendStringEqualsLiteralCi(name, "encoding") {
			if FAILURE == ZendDeclareIsFirstStatement(ast) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Encoding declaration pragma must be "+"the very first statement in the script")
			}
		} else if ZendStringEqualsLiteralCi(name, "strict_types") {
			var value_zv Zval
			if FAILURE == ZendDeclareIsFirstStatement(ast) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "strict_types declaration must be "+"the very first statement in the script")
			}
			if ast.GetChild()[1] != nil {
				ZendErrorNoreturn(E_COMPILE_ERROR, "strict_types declaration must not "+"use block mode")
			}
			ZendConstExprToZval(&value_zv, value_ast)
			if Z_TYPE(value_zv) != IS_LONG || Z_LVAL(value_zv) != 0 && Z_LVAL(value_zv) != 1 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "strict_types declaration must have 0 or 1 as its value")
			}
			if Z_LVAL(value_zv) == 1 {
				CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_STRICT_TYPES)
			}
		} else {
			ZendError(E_COMPILE_WARNING, "Unsupported declare '%s'", ZSTR_VAL(name))
		}
	}
	if stmt_ast != nil {
		ZendCompileStmt(stmt_ast)
		FC(declarables) = orig_declarables
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
		n = MIN(func_.GetNumArgs(), MAX_ARG_FLAG_NUM)
		i = 0
		for i < n {
			ZEND_SET_ARG_FLAG(func_, i+1, func_.GetArgInfo()[i].GetPassByReference())
			i++
		}
		if UNEXPECTED((func_.GetFnFlags()&ZEND_ACC_VARIADIC) != 0 && func_.GetArgInfo()[i].GetPassByReference() != 0) {
			var pass_by_reference uint32 = func_.GetArgInfo()[i].GetPassByReference()
			for i < MAX_ARG_FLAG_NUM {
				ZEND_SET_ARG_FLAG(func_, i+1, pass_by_reference)
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
	if (ast.GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
		allow_null = 1
		ast.SetAttr(ast.GetAttr() &^ ZEND_TYPE_NULLABLE)
	}
	if ast.GetKind() == ZEND_AST_TYPE {
		return ZEND_TYPE_ENCODE(ast.GetAttr(), allow_null)
	} else {
		var class_name *ZendString = ZendAstGetStr(ast)
		var type_code ZendUchar = ZendLookupBuiltinTypeByName(class_name)
		if type_code != 0 {
			if (ast.GetAttr() & ZEND_NAME_NOT_FQ) != ZEND_NAME_NOT_FQ {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Type declaration '%s' must be unqualified", ZSTR_VAL(ZendStringTolower(class_name)))
			}
			type_ = ZEND_TYPE_ENCODE(type_code, allow_null)
		} else {
			var fetch_type uint32 = ZendGetClassFetchTypeAst(ast)
			if fetch_type == ZEND_FETCH_CLASS_DEFAULT {
				class_name = ZendResolveClassNameAst(ast)
				ZendAssertValidClassName(class_name)
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				ZendStringAddref(class_name)
			}
			type_ = ZEND_TYPE_ENCODE_CLASS(class_name, allow_null)
		}
	}
	ast.SetAttr(orig_ast_attr)
	return type_
}

/* }}} */

func ZendCompileParams(ast *ZendAst, return_type_ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var arg_infos *ZendArgInfo
	if return_type_ast != nil {

		/* Use op_array->arg_info[-1] for return type */

		arg_infos = SafeEmalloc(b.SizeOf("zend_arg_info"), list.GetChildren()+1, 0)
		arg_infos.SetName(nil)
		arg_infos.SetPassByReference((op_array.GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0)
		arg_infos.SetIsVariadic(0)
		arg_infos.SetType(ZendCompileTypename(return_type_ast, 0))
		if ZEND_TYPE_CODE(arg_infos.GetType()) == IS_VOID && ZEND_TYPE_ALLOW_NULL(arg_infos.GetType()) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Void type cannot be nullable")
		}
		arg_infos++
		op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_HAS_RETURN_TYPE)
	} else {
		if list.GetChildren() == 0 {
			return
		}
		arg_infos = SafeEmalloc(b.SizeOf("zend_arg_info"), list.GetChildren(), 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var param_ast *ZendAst = list.GetChild()[i]
		var type_ast *ZendAst = param_ast.GetChild()[0]
		var var_ast *ZendAst = param_ast.GetChild()[1]
		var default_ast *ZendAst = param_ast.GetChild()[2]
		var name *ZendString = ZvalMakeInternedString(ZendAstGetZval(var_ast))
		var is_ref ZendBool = (param_ast.GetAttr() & ZEND_PARAM_REF) != 0
		var is_variadic ZendBool = (param_ast.GetAttr() & ZEND_PARAM_VARIADIC) != 0
		var var_node Znode
		var default_node Znode
		var opcode ZendUchar
		var opline *ZendOp
		var arg_info *ZendArgInfo
		if ZendIsAutoGlobal(name) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign auto-global variable %s", ZSTR_VAL(name))
		}
		var_node.SetOpType(IS_CV)
		var_node.GetOp().SetVar(LookupCv(name))
		if EX_VAR_TO_NUM(var_node.GetOp().GetVar()) != i {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Redefinition of parameter $%s", ZSTR_VAL(name))
		} else if ZendStringEqualsLiteral(name, "this") {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use $this as parameter")
		}
		if (op_array.GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Only the last parameter can be variadic")
		}
		if is_variadic != 0 {
			opcode = ZEND_RECV_VARIADIC
			default_node.SetOpType(IS_UNUSED)
			op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_VARIADIC)
			if default_ast != nil {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Variadic parameter cannot have a default value")
			}
		} else if default_ast != nil {

			/* we cannot substitute constants here or it will break ReflectionParameter::getDefaultValueConstantName() and ReflectionParameter::isDefaultValueConstant() */

			var cops uint32 = CompilerGlobals.GetCompilerOptions()
			CompilerGlobals.SetCompilerOptions(CompilerGlobals.GetCompilerOptions() | ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION | ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION)
			opcode = ZEND_RECV_INIT
			default_node.SetOpType(IS_CONST)
			ZendConstExprToZval(&default_node.u.constant, default_ast)
			CompilerGlobals.SetCompilerOptions(cops)
		} else {
			opcode = ZEND_RECV
			default_node.SetOpType(IS_UNUSED)
			op_array.SetRequiredNumArgs(i + 1)
		}
		arg_info = &arg_infos[i]
		arg_info.SetName(ZendStringCopy(name))
		arg_info.SetPassByReference(is_ref)
		arg_info.SetIsVariadic(is_variadic)

		/* TODO: Keep compatibility, but may be better reset "allow_null" ??? */

		arg_info.SetType(ZEND_TYPE_ENCODE(0, 1))
		if type_ast != nil {
			var has_null_default ZendBool = default_ast != nil && (Z_TYPE(default_node.GetConstant()) == IS_NULL || Z_TYPE(default_node.GetConstant()) == IS_CONSTANT_AST && Z_ASTVAL(default_node.GetConstant()).GetKind() == ZEND_AST_CONSTANT && strcasecmp(ZSTR_VAL(ZendAstGetConstantName(Z_ASTVAL(default_node.GetConstant()))), "NULL") == 0)
			op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_HAS_TYPE_HINTS)
			arg_info.SetType(ZendCompileTypename(type_ast, has_null_default))
			if ZEND_TYPE_CODE(arg_info.GetType()) == IS_VOID {
				ZendErrorNoreturn(E_COMPILE_ERROR, "void cannot be used as a parameter type")
			}
			if type_ast.GetKind() == ZEND_AST_TYPE {
				if ZEND_TYPE_CODE(arg_info.GetType()) == IS_ARRAY {
					if default_ast != nil && has_null_default == 0 && Z_TYPE(default_node.GetConstant()) != IS_ARRAY && Z_TYPE(default_node.GetConstant()) != IS_CONSTANT_AST {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with array type can only be an array or NULL")
					}
				} else if ZEND_TYPE_CODE(arg_info.GetType()) == IS_CALLABLE && default_ast != nil {
					if has_null_default == 0 && Z_TYPE(default_node.GetConstant()) != IS_CONSTANT_AST {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with callable type can only be NULL")
					}
				}
			} else {
				if default_ast != nil && has_null_default == 0 && Z_TYPE(default_node.GetConstant()) != IS_CONSTANT_AST {
					if ZEND_TYPE_IS_CLASS(arg_info.GetType()) {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with a class type can only be NULL")
					} else {
						switch ZEND_TYPE_CODE(arg_info.GetType()) {
						case IS_DOUBLE:
							if Z_TYPE(default_node.GetConstant()) != IS_DOUBLE && Z_TYPE(default_node.GetConstant()) != IS_LONG {
								ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with a float type can only be float, integer, or NULL")
							}
							ConvertToDouble(&default_node.u.constant)
							break
						case IS_ITERABLE:
							if Z_TYPE(default_node.GetConstant()) != IS_ARRAY {
								ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with iterable type can only be an array or NULL")
							}
							break
						case IS_OBJECT:
							ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with an object type can only be NULL")
							break
						default:
							if !(ZEND_SAME_FAKE_TYPE(ZEND_TYPE_CODE(arg_info.GetType()), Z_TYPE(default_node.GetConstant()))) {
								ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for parameters "+"with a %s type can only be %s or NULL", ZendGetTypeByConst(ZEND_TYPE_CODE(arg_info.GetType())), ZendGetTypeByConst(ZEND_TYPE_CODE(arg_info.GetType())))
							}
							break
						}
					}
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, nil, &default_node)
		opline.SetResultType(&var_node.GetOpType())
		if &var_node.GetOpType() == IS_CONST {
			opline.GetResult().SetConstant(ZendAddLiteral(&(&var_node).u.constant))
		} else {
			opline.SetResult(&var_node.GetOp())
		}
		opline.GetOp1().SetNum(i + 1)
		if type_ast != nil {

			/* Allocate cache slot to speed-up run-time class resolution */

			if opline.GetOpcode() == ZEND_RECV_INIT {
				if ZEND_TYPE_IS_CLASS(arg_info.GetType()) {
					opline.SetExtendedValue(ZendAllocCacheSlot())
				}
			} else {
				if ZEND_TYPE_IS_CLASS(arg_info.GetType()) {
					opline.GetOp2().SetNum(op_array.GetCacheSize())
					op_array.SetCacheSize(op_array.GetCacheSize() + b.SizeOf("void *"))
				} else {
					opline.GetOp2().SetNum(-1)
				}
			}

			/* Allocate cache slot to speed-up run-time class resolution */

		} else {
			if opline.GetOpcode() != ZEND_RECV_INIT {
				opline.GetOp2().SetNum(-1)
			}
		}
	}

	/* These are assigned at the end to avoid uninitialized memory in case of an error */

	op_array.SetNumArgs(list.GetChildren())
	op_array.SetArgInfo(arg_infos)

	/* Don't count the variadic argument */

	if (op_array.GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
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
		op_array.SetStaticVariables(ZendNewArray(8))
	}
	for i = 0; i < list.GetChildren(); i++ {
		var var_name_ast *ZendAst = list.GetChild()[i]
		var var_name *ZendString = ZvalMakeInternedString(ZendAstGetZval(var_name_ast))
		var mode uint32 = var_name_ast.GetAttr()
		var opline *ZendOp
		var value *Zval
		if ZendStringEqualsLiteral(var_name, "this") {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use $this as lexical variable")
		}
		if ZendIsAutoGlobal(var_name) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use auto-global as lexical variable")
		}
		value = ZendHashAdd(op_array.GetStaticVariables(), var_name, &(ExecutorGlobals.GetUninitializedZval()))
		if value == nil {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use variable $%s twice", ZSTR_VAL(var_name))
		}
		CompilerGlobals.SetZendLineno(ZendAstGetLineno(var_name_ast))
		opline = ZendEmitOp(nil, ZEND_BIND_LEXICAL, closure, nil)
		opline.SetOp2Type(IS_CV)
		opline.GetOp2().SetVar(LookupCv(var_name))
		opline.SetExtendedValue(uint32_t((*byte)(value-(*byte)(op_array.GetStaticVariables().GetArData()))) | mode)
	}
}

/* }}} */

func FindImplicitBindsRecursively(info *ClosureInfo, ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_VAR {
		var name_ast *ZendAst = ast.GetChild()[0]
		if name_ast.GetKind() == ZEND_AST_ZVAL && Z_TYPE_P(ZendAstGetZval(name_ast)) == IS_STRING {
			var name *ZendString = ZendAstGetStr(name_ast)
			if ZendIsAutoGlobal(name) != 0 {

				/* These is no need to explicitly import auto-globals. */

				return

				/* These is no need to explicitly import auto-globals. */

			}
			if ZendStringEqualsLiteral(name, "this") {

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
	ZendHashInit(&info.uses, param_list.GetChildren(), nil, nil, 0)
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

	if ZendHashNumElements(&info.uses) == 0 {
		return
	}
	if op_array.GetStaticVariables() == nil {
		op_array.SetStaticVariables(ZendNewArray(8))
	}
	for {
		var __ht *HashTable = &info.uses
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			var_name = _p.GetKey()
			var value *Zval = ZendHashAdd(op_array.GetStaticVariables(), var_name, &(ExecutorGlobals.GetUninitializedZval()))
			var offset uint32 = uint32_t((*byte)(value - (*byte)(op_array.GetStaticVariables().GetArData())))
			opline = ZendEmitOp(nil, ZEND_BIND_LEXICAL, closure, nil)
			opline.SetOp2Type(IS_CV)
			opline.GetOp2().SetVar(LookupCv(var_name))
			opline.SetExtendedValue(offset | ZEND_BIND_IMPLICIT)
		}
		break
	}
}
func ZendCompileClosureUses(ast *ZendAst) {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var var_ast *ZendAst = list.GetChild()[i]
		var var_name *ZendString = ZendAstGetStr(var_ast)
		var zv Zval
		ZVAL_NULL(&zv)
		var i int
		for i = 0; i < op_array.GetLastVar(); i++ {
			if ZendStringEquals(op_array.GetVars()[i], var_name) != 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use lexical variable $%s as a parameter name", ZSTR_VAL(var_name))
			}
		}
		CompilerGlobals.SetZendLineno(ZendAstGetLineno(var_ast))
		ZendCompileStaticVarCommon(var_name, &zv, b.Cond(var_ast.GetAttr() != 0, ZEND_BIND_REF, 0))
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			var_name = _p.GetKey()
			var zv Zval
			ZVAL_NULL(&zv)
			ZendCompileStaticVarCommon(var_name, &zv, ZEND_BIND_IMPLICIT)
		}
		break
	}
}
func ZendBeginMethodDecl(op_array *ZendOpArray, name *ZendString, has_body ZendBool) {
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var in_interface ZendBool = (ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0
	var in_trait ZendBool = (ce.GetCeFlags() & ZEND_ACC_TRAIT) != 0
	var is_public ZendBool = (op_array.GetFnFlags() & ZEND_ACC_PUBLIC) != 0
	var is_static ZendBool = (op_array.GetFnFlags() & ZEND_ACC_STATIC) != 0
	var lcname *ZendString
	if in_interface != 0 {
		if is_public == 0 || (op_array.GetFnFlags()&(ZEND_ACC_FINAL|ZEND_ACC_ABSTRACT)) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Access type for interface method "+"%s::%s() must be omitted", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
		}
		op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_ABSTRACT)
	}
	if (op_array.GetFnFlags() & ZEND_ACC_ABSTRACT) != 0 {
		if (op_array.GetFnFlags() & ZEND_ACC_PRIVATE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "%s function %s::%s() cannot be declared private", b.Cond(in_interface != 0, "Interface", "Abstract"), ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
		}
		if has_body != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "%s function %s::%s() cannot contain body", b.Cond(in_interface != 0, "Interface", "Abstract"), ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
		}
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
	} else if has_body == 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Non-abstract method %s::%s() must contain body", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
	}
	op_array.SetScope(ce)
	op_array.SetFunctionName(ZendStringCopy(name))
	lcname = ZendStringTolower(name)
	lcname = ZendNewInternedString(lcname)
	if ZendHashAddPtr(&ce.function_table, lcname, op_array) == nil {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot redeclare %s::%s()", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
	}
	if in_interface != 0 {
		if ZSTR_VAL(lcname)[0] != '_' || ZSTR_VAL(lcname)[1] != '_' {

		} else if ZendStringEqualsLiteral(lcname, ZEND_CALL_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __call() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_CALLSTATIC_FUNC_NAME) {
			if is_public == 0 || is_static == 0 {
				ZendError(E_WARNING, "The magic method __callStatic() must have "+"public visibility and be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_GET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __get() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_SET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __set() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_UNSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_ISSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_TOSTRING_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_INVOKE_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_DEBUGINFO_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
		}
	} else {
		if in_trait == 0 && ZendStringEqualsCi(lcname, ce.GetName()) {
			if ce.GetConstructor() == nil {
				ce.SetConstructor((*ZendFunction)(op_array))
			}
		} else if ZendStringEqualsLiteral(lcname, "serialize") {
			ce.SetSerializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_ALLOW_STATIC)
			}
		} else if ZendStringEqualsLiteral(lcname, "unserialize") {
			ce.SetUnserializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_ALLOW_STATIC)
			}
		} else if ZSTR_VAL(lcname)[0] != '_' || ZSTR_VAL(lcname)[1] != '_' {
			if is_static == 0 {
				op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_ALLOW_STATIC)
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_CONSTRUCTOR_FUNC_NAME) {
			ce.SetConstructor((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_DESTRUCTOR_FUNC_NAME) {
			ce.SetDestructor((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_CLONE_FUNC_NAME) {
			ce.SetClone((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_CALL_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __call() must have "+"public visibility and cannot be static")
			}
			ce.SetCall((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_CALLSTATIC_FUNC_NAME) {
			if is_public == 0 || is_static == 0 {
				ZendError(E_WARNING, "The magic method __callStatic() must have "+"public visibility and be static")
			}
			ce.SetCallstatic((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_GET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __get() must have "+"public visibility and cannot be static")
			}
			ce.SetGet((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
		} else if ZendStringEqualsLiteral(lcname, ZEND_SET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __set() must have "+"public visibility and cannot be static")
			}
			ce.SetSet((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
		} else if ZendStringEqualsLiteral(lcname, ZEND_UNSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
			ce.SetUnset((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
		} else if ZendStringEqualsLiteral(lcname, ZEND_ISSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
			ce.SetIsset((*ZendFunction)(op_array))
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
		} else if ZendStringEqualsLiteral(lcname, ZEND_TOSTRING_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
			ce.SetTostring((*ZendFunction)(op_array))
		} else if ZendStringEqualsLiteral(lcname, ZEND_INVOKE_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if ZendStringEqualsLiteral(lcname, ZEND_DEBUGINFO_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				ZendError(E_WARNING, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
			ce.SetDebugInfo((*ZendFunction)(op_array))
		} else if is_static == 0 {
			op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_ALLOW_STATIC)
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
	lcname = ZendStringTolower(name)
	if FC(imports_function) {
		var import_name *ZendString = ZendHashFindPtrLc(FC(imports_function), ZSTR_VAL(unqualified_name), ZSTR_LEN(unqualified_name))
		if import_name != nil && !(ZendStringEqualsCi(lcname, import_name)) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare function %s "+"because the name is already in use", ZSTR_VAL(name))
		}
	}
	if ZendStringEqualsLiteral(lcname, ZEND_AUTOLOAD_FUNC_NAME) {
		if ZendAstGetList(params_ast).GetChildren() != 1 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "%s() must take exactly 1 argument", ZEND_AUTOLOAD_FUNC_NAME)
		}
		ZendError(E_DEPRECATED, "__autoload() is deprecated, use spl_autoload_register() instead")
	}
	if ZendStringEqualsLiteralCi(unqualified_name, "assert") {
		ZendError(E_DEPRECATED, "Defining a custom assert() function is deprecated, "+"as the function has special semantics")
	}
	ZendRegisterSeenSymbol(lcname, ZEND_SYMBOL_FUNCTION)
	if toplevel != 0 {
		if UNEXPECTED(ZendHashAddPtr(CompilerGlobals.GetFunctionTable(), lcname, op_array) == nil) {
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
		if ZendHashAddPtr(CompilerGlobals.GetFunctionTable(), key, op_array) {
			break
		}
	}
	if (op_array.GetFnFlags() & ZEND_ACC_CLOSURE) != 0 {
		opline = ZendEmitOpTmp(result, ZEND_DECLARE_LAMBDA_FUNCTION, nil, nil)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetOp1Type(IS_CONST)
		LITERAL_STR(opline.GetOp1(), key)
	} else {
		opline = GetNextOp()
		opline.SetOpcode(ZEND_DECLARE_FUNCTION)
		opline.SetOp1Type(IS_CONST)
		LITERAL_STR(opline.GetOp1(), ZendStringCopy(lcname))

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
	var orig_class_entry *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var orig_op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var op_array *ZendOpArray = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_op_array"))
	var orig_oparray_context ZendOparrayContext
	var info ClosureInfo
	memset(&info, 0, b.SizeOf("closure_info"))
	InitOpArray(op_array, ZEND_USER_FUNCTION, INITIAL_OP_ARRAY_SIZE)
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_PRELOADED)
		ZEND_MAP_PTR_NEW(op_array.run_time_cache)
		ZEND_MAP_PTR_NEW(op_array.static_variables_ptr)
	} else {
		ZEND_MAP_PTR_INIT(op_array.run_time_cache, ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("void *")))
		ZEND_MAP_PTR_SET(op_array.run_time_cache, nil)
	}
	op_array.SetFnFlags(op_array.GetFnFlags() | orig_op_array.GetFnFlags()&ZEND_ACC_STRICT_TYPES)
	op_array.SetFnFlags(op_array.GetFnFlags() | decl.GetFlags())
	op_array.SetLineStart(decl.GetStartLineno())
	op_array.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		op_array.SetDocComment(ZendStringCopy(decl.GetDocComment()))
	}
	if decl.GetKind() == ZEND_AST_CLOSURE || decl.GetKind() == ZEND_AST_ARROW_FUNC {
		op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_CLOSURE)
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
	CompilerGlobals.SetActiveOpArray(op_array)

	/* Do not leak the class scope into free standing functions, even if they are dynamically
	 * defined inside a class method. This is necessary for correct handling of magic constants.
	 * For example __CLASS__ should always be "" inside a free standing function. */

	if decl.GetKind() == ZEND_AST_FUNC_DECL {
		CompilerGlobals.SetActiveClassEntry(nil)
	}
	if toplevel != 0 {
		op_array.SetFnFlags(op_array.GetFnFlags() | ZEND_ACC_TOP_LEVEL)
	}
	ZendOparrayContextBegin(&orig_oparray_context)
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) != 0 {
		var opline_ext *ZendOp = ZendEmitOp(nil, ZEND_EXT_NOP, nil, nil)
		opline_ext.SetLineno(decl.GetStartLineno())
	}

	/* Push a separator to the loop variable stack */

	var dummy_var ZendLoopVar
	dummy_var.SetOpcode(ZEND_RETURN)
	ZendStackPush(&(CompilerGlobals.GetLoopVarStack()), any(&dummy_var))
	ZendCompileParams(params_ast, return_type_ast)
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_GENERATOR) != 0 {
		ZendMarkFunctionAsGenerator()
		ZendEmitOp(nil, ZEND_GENERATOR_CREATE, nil, nil)
	}
	if decl.GetKind() == ZEND_AST_ARROW_FUNC {
		ZendCompileImplicitClosureUses(&info)
		ZendHashDestroy(&info.uses)
	} else if uses_ast != nil {
		ZendCompileClosureUses(uses_ast)
	}
	ZendCompileStmt(stmt_ast)
	if is_method != 0 {
		ZendCheckMagicMethodImplementation(CompilerGlobals.GetActiveClassEntry(), (*ZendFunction)(op_array), E_COMPILE_ERROR)
	}

	/* put the implicit return on the really last line */

	CompilerGlobals.SetZendLineno(decl.GetEndLineno())
	ZendDoExtendedStmt()
	ZendEmitFinalReturn(0)
	PassTwo(CompilerGlobals.GetActiveOpArray())
	ZendOparrayContextEnd(&orig_oparray_context)

	/* Pop the loop variable stack separator */

	ZendStackDelTop(&(CompilerGlobals.GetLoopVarStack()))
	CompilerGlobals.SetActiveOpArray(orig_op_array)
	CompilerGlobals.SetActiveClassEntry(orig_class_entry)
}

/* }}} */

func ZendCompilePropDecl(ast *ZendAst, type_ast *ZendAst, flags uint32) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var i uint32
	var children uint32 = list.GetChildren()
	if (ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Interfaces may not include member variables")
	}
	if (flags & ZEND_ACC_ABSTRACT) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Properties cannot be declared abstract")
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
			if ZEND_TYPE_CODE(type_) == IS_VOID || ZEND_TYPE_CODE(type_) == IS_CALLABLE {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Property %s::$%s cannot have type %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name), ZendGetTypeByConst(ZEND_TYPE_CODE(type_)))
			}
		}

		/* Doc comment has been appended as last element in ZEND_AST_PROP_ELEM ast */

		if doc_comment_ast != nil {
			doc_comment = ZendStringCopy(ZendAstGetStr(doc_comment_ast))
		}
		if (flags & ZEND_ACC_FINAL) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare property %s::$%s final, "+"the final modifier is allowed only for methods and classes", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
		}
		if ZendHashExists(&ce.properties_info, name) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot redeclare %s::$%s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name))
		}
		if value_ast != nil {
			ZendConstExprToZval(&value_zv, value_ast)
			if ZEND_TYPE_IS_SET(type_) && !(Z_CONSTANT(value_zv)) {
				if Z_TYPE(value_zv) == IS_NULL {
					if !(ZEND_TYPE_ALLOW_NULL(type_)) {
						var name *byte = b.CondF(ZEND_TYPE_IS_CLASS(type_), func() []byte { return ZSTR_VAL(ZEND_TYPE_NAME(type_)) }, func() *byte { return ZendGetTypeByConst(ZEND_TYPE_CODE(type_)) })
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for property of type %s may not be null. "+"Use the nullable type ?%s to allow null default value", name, name)
					}
				} else if ZEND_TYPE_IS_CLASS(type_) {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Property of type %s may not have default value", ZSTR_VAL(ZEND_TYPE_NAME(type_)))
				} else if ZEND_TYPE_CODE(type_) == IS_ARRAY || ZEND_TYPE_CODE(type_) == IS_ITERABLE {
					if Z_TYPE(value_zv) != IS_ARRAY {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for property of type %s can only be an array", ZendGetTypeByConst(ZEND_TYPE_CODE(type_)))
					}
				} else if ZEND_TYPE_CODE(type_) == IS_DOUBLE {
					if Z_TYPE(value_zv) != IS_DOUBLE && Z_TYPE(value_zv) != IS_LONG {
						ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for property of type float can only be float or int")
					}
					ConvertToDouble(&value_zv)
				} else if !(ZEND_SAME_FAKE_TYPE(ZEND_TYPE_CODE(type_), Z_TYPE(value_zv))) {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Default value for property of type %s can only be %s", ZendGetTypeByConst(ZEND_TYPE_CODE(type_)), ZendGetTypeByConst(ZEND_TYPE_CODE(type_)))
				}
			}
		} else if !(ZEND_TYPE_IS_SET(type_)) {
			ZVAL_NULL(&value_zv)
		} else {
			ZVAL_UNDEF(&value_zv)
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
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var i uint32
	if (ce.GetCeFlags() & ZEND_ACC_TRAIT) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Traits cannot have constants")
		return
	}
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = const_ast.GetChild()[0]
		var value_ast *ZendAst = const_ast.GetChild()[1]
		var doc_comment_ast *ZendAst = const_ast.GetChild()[2]
		var name *ZendString = ZvalMakeInternedString(ZendAstGetZval(name_ast))
		var doc_comment *ZendString = b.CondF1(doc_comment_ast != nil, func() *ZendString { return ZendStringCopy(ZendAstGetStr(doc_comment_ast)) }, nil)
		var value_zv Zval
		if UNEXPECTED((ast.GetAttr() & (ZEND_ACC_STATIC | ZEND_ACC_ABSTRACT | ZEND_ACC_FINAL)) != 0) {
			if (ast.GetAttr() & ZEND_ACC_STATIC) != 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'static' as constant modifier")
			} else if (ast.GetAttr() & ZEND_ACC_ABSTRACT) != 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'abstract' as constant modifier")
			} else if (ast.GetAttr() & ZEND_ACC_FINAL) != 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'final' as constant modifier")
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
	var precedence *ZendTraitPrecedence = Emalloc(b.SizeOf("zend_trait_precedence") + (insteadof_list.GetChildren()-1)*b.SizeOf("zend_string *"))
	ZendCompileMethodRef(method_ref_ast, &precedence.trait_method)
	precedence.SetNumExcludes(insteadof_list.GetChildren())
	for i = 0; i < insteadof_list.GetChildren(); i++ {
		var name_ast *ZendAst = insteadof_list.GetChild()[i]
		precedence.GetExcludeClassNames()[i] = ZendResolveClassNameAst(name_ast)
	}
	ZendAddToList(&(CompilerGlobals.GetActiveClassEntry()).trait_precedences, precedence)
}

/* }}} */

func ZendCompileTraitAlias(ast *ZendAst) {
	var method_ref_ast *ZendAst = ast.GetChild()[0]
	var alias_ast *ZendAst = ast.GetChild()[1]
	var modifiers uint32 = ast.GetAttr()
	var alias *ZendTraitAlias
	if modifiers == ZEND_ACC_STATIC {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'static' as method modifier")
	} else if modifiers == ZEND_ACC_ABSTRACT {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'abstract' as method modifier")
	} else if modifiers == ZEND_ACC_FINAL {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use 'final' as method modifier")
	}
	alias = Emalloc(b.SizeOf("zend_trait_alias"))
	ZendCompileMethodRef(method_ref_ast, &alias.trait_method)
	alias.SetModifiers(modifiers)
	if alias_ast != nil {
		alias.SetAlias(ZendStringCopy(ZendAstGetStr(alias_ast)))
	} else {
		alias.SetAlias(nil)
	}
	ZendAddToList(&(CompilerGlobals.GetActiveClassEntry()).trait_aliases, alias)
}

/* }}} */

func ZendCompileUseTrait(ast *ZendAst) {
	var traits *ZendAstList = ZendAstGetList(ast.GetChild()[0])
	var adaptations *ZendAstList = b.CondF1(ast.GetChild()[1] != nil, func() *ZendAstList { return ZendAstGetList(ast.GetChild()[1]) }, nil)
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var i uint32
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_IMPLEMENT_TRAITS)
	ce.SetTraitNames(Erealloc(ce.GetTraitNames(), b.SizeOf("zend_class_name")*(ce.GetNumTraits()+traits.GetChildren())))
	for i = 0; i < traits.GetChildren(); i++ {
		var trait_ast *ZendAst = traits.GetChild()[i]
		var name *ZendString = ZendAstGetStr(trait_ast)
		if (ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use traits inside of interfaces. "+"%s is used in %s", ZSTR_VAL(name), ZSTR_VAL(ce.GetName()))
		}
		switch ZendGetClassFetchType(name) {
		case ZEND_FETCH_CLASS_SELF:

		case ZEND_FETCH_CLASS_PARENT:

		case ZEND_FETCH_CLASS_STATIC:
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as trait name "+"as it is reserved", ZSTR_VAL(name))
			break
		}
		ce.GetTraitNames()[ce.GetNumTraits()].SetName(ZendResolveClassNameAst(trait_ast))
		ce.GetTraitNames()[ce.GetNumTraits()].SetLcName(ZendStringTolower(ce.GetTraitNames()[ce.GetNumTraits()].GetName()))
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
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	var interface_names *ZendClassName
	var i uint32
	interface_names = Emalloc(b.SizeOf("zend_class_name") * list.GetChildren())
	for i = 0; i < list.GetChildren(); i++ {
		var class_ast *ZendAst = list.GetChild()[i]
		var name *ZendString = ZendAstGetStr(class_ast)
		if ZendIsConstDefaultClassRef(class_ast) == 0 {
			Efree(interface_names)
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as interface name as it is reserved", ZSTR_VAL(name))
		}
		interface_names[i].SetName(ZendResolveClassNameAst(class_ast))
		interface_names[i].SetLcName(ZendStringTolower(interface_names[i].GetName()))
	}
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_IMPLEMENT_INTERFACES)
	ce.SetNumInterfaces(list.GetChildren())
	ce.interface_names = interface_names
}

/* }}} */

func ZendGenerateAnonClassName(start_lineno uint32) *ZendString {
	var filename *ZendString = CompilerGlobals.GetActiveOpArray().GetFilename()
	var result *ZendString = ZendStrpprintf(0, "class@anonymous%c%s:%"+"u"+"$%"+PRIx32, '0', ZSTR_VAL(filename), start_lineno, b.PostInc(&(CompilerGlobals.GetRtdKeyCounter())))
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
	var ce *ZendClassEntry = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_class_entry"))
	var opline *ZendOp
	var original_ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	if EXPECTED((decl.GetFlags() & ZEND_ACC_ANON_CLASS) == 0) {
		var unqualified_name *ZendString = decl.GetName()
		if CompilerGlobals.GetActiveClassEntry() != nil {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Class declarations may not be nested")
		}
		ZendAssertValidClassName(unqualified_name)
		name = ZendPrefixWithNs(unqualified_name)
		name = ZendNewInternedString(name)
		lcname = ZendStringTolower(name)
		if FC(imports) {
			var import_name *ZendString = ZendHashFindPtrLc(FC(imports), ZSTR_VAL(unqualified_name), ZSTR_LEN(unqualified_name))
			if import_name != nil && !(ZendStringEqualsCi(lcname, import_name)) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare class %s "+"because the name is already in use", ZSTR_VAL(name))
			}
		}
		ZendRegisterSeenSymbol(lcname, ZEND_SYMBOL_CLASS)
	} else {

		/* Find an anon class name that is not in use yet. */

		name = nil
		lcname = nil
		for {
			ZendTmpStringRelease(name)
			ZendTmpStringRelease(lcname)
			name = ZendGenerateAnonClassName(decl.GetStartLineno())
			lcname = ZendStringTolower(name)
			if ZendHashExists(CompilerGlobals.GetClassTable(), lcname) == 0 {
				break
			}
		}
	}
	lcname = ZendNewInternedString(lcname)
	ce.SetType(ZEND_USER_CLASS)
	ce.SetName(name)
	ZendInitializeClassData(ce, 1)
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_PRELOADED)
		ZEND_MAP_PTR_NEW(ce.static_members_table)
	}
	ce.SetCeFlags(ce.GetCeFlags() | decl.GetFlags())
	ce.SetFilename(ZendGetCompiledFilename())
	ce.SetLineStart(decl.GetStartLineno())
	ce.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		ce.SetDocComment(ZendStringCopy(decl.GetDocComment()))
	}
	if UNEXPECTED((decl.GetFlags() & ZEND_ACC_ANON_CLASS) != 0) {

		/* Serialization is not supported for anonymous classes */

		ce.SetSerialize(ZendClassSerializeDeny)
		ce.SetUnserialize(ZendClassUnserializeDeny)
	}
	if extends_ast != nil {
		var extends_node Znode
		var extends_name *ZendString
		if ZendIsConstDefaultClassRef(extends_ast) == 0 {
			extends_name = ZendAstGetStr(extends_ast)
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as class name as it is reserved", ZSTR_VAL(extends_name))
		}
		ZendCompileExpr(&extends_node, extends_ast)
		if extends_node.GetOpType() != IS_CONST || Z_TYPE(extends_node.GetConstant()) != IS_STRING {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Illegal class name")
		}
		extends_name = Z_STR(extends_node.GetConstant())
		ce.parent_name = ZendResolveClassName(extends_name, b.CondF1(extends_ast.GetKind() == ZEND_AST_ZVAL, func() ZendAstAttr { return extends_ast.GetAttr() }, ZEND_NAME_FQ))
		ZendStringReleaseEx(extends_name, 0)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_INHERITED)
	}
	CompilerGlobals.SetActiveClassEntry(ce)
	ZendCompileStmt(stmt_ast)

	/* Reset lineno for final opcodes and errors */

	CompilerGlobals.SetZendLineno(ast.GetLineno())
	if (ce.GetCeFlags() & ZEND_ACC_IMPLEMENT_TRAITS) == 0 {

		/* For traits this check is delayed until after trait binding */

		ZendCheckDeprecatedConstructor(ce)

		/* For traits this check is delayed until after trait binding */

	}
	if ce.GetConstructor() != nil {
		ce.GetConstructor().SetFnFlags(ce.GetConstructor().GetFnFlags() | ZEND_ACC_CTOR)
		if (ce.GetConstructor().GetFnFlags() & ZEND_ACC_STATIC) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Constructor %s::%s() cannot be static", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetConstructor().GetFunctionName()))
		}
		if (ce.GetConstructor().GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Constructor %s::%s() cannot declare a return type", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetConstructor().GetFunctionName()))
		}
	}
	if ce.GetDestructor() != nil {
		ce.GetDestructor().SetFnFlags(ce.GetDestructor().GetFnFlags() | ZEND_ACC_DTOR)
		if (ce.GetDestructor().GetFnFlags() & ZEND_ACC_STATIC) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Destructor %s::%s() cannot be static", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetDestructor().GetFunctionName()))
		} else if (ce.GetDestructor().GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Destructor %s::%s() cannot declare a return type", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetDestructor().GetFunctionName()))
		}
	}
	if ce.GetClone() != nil {
		if (ce.GetClone().GetFnFlags() & ZEND_ACC_STATIC) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Clone method %s::%s() cannot be static", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetClone().GetFunctionName()))
		} else if (ce.GetClone().GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Clone method %s::%s() cannot declare a return type", ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetClone().GetFunctionName()))
		}
	}
	if implements_ast != nil {
		ZendCompileImplements(implements_ast)
	}
	if (ce.GetCeFlags() & (ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) == ZEND_ACC_IMPLICIT_ABSTRACT_CLASS {
		ZendVerifyAbstractClass(ce)
	}
	CompilerGlobals.SetActiveClassEntry(original_ce)
	if toplevel != 0 {
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_TOP_LEVEL)
	}
	if toplevel != 0 && (ce.GetCeFlags()&(ZEND_ACC_IMPLEMENT_INTERFACES|ZEND_ACC_IMPLEMENT_TRAITS)) == 0 && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_PRELOAD) == 0 {
		if extends_ast != nil {
			var parent_ce *ZendClassEntry = ZendLookupClassEx(ce.parent_name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if parent_ce != nil && (parent_ce.GetType() != ZEND_INTERNAL_CLASS || (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_CLASSES) == 0) && (parent_ce.GetType() != ZEND_USER_CLASS || (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) == 0 || parent_ce.GetFilename() == ce.GetFilename()) {
				CompilerGlobals.SetZendLineno(decl.GetEndLineno())
				if ZendTryEarlyBind(ce, parent_ce, lcname, nil) != 0 {
					CompilerGlobals.SetZendLineno(ast.GetLineno())
					ZendStringRelease(lcname)
					return nil
				}
				CompilerGlobals.SetZendLineno(ast.GetLineno())
			}
		} else {
			if EXPECTED(ZendHashAddPtr(CompilerGlobals.GetClassTable(), lcname, ce) != nil) {
				ZendStringRelease(lcname)
				ZendBuildPropertiesInfoTable(ce)
				ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_LINKED)
				return nil
			}
		}
	}
	opline = GetNextOp()
	if ce.parent_name {

		/* Lowercased parent name */

		var lc_parent_name *ZendString = ZendStringTolower(ce.parent_name)
		opline.SetOp2Type(IS_CONST)
		LITERAL_STR(opline.GetOp2(), lc_parent_name)
	}
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), lcname)
	if (decl.GetFlags() & ZEND_ACC_ANON_CLASS) != 0 {
		opline.SetOpcode(ZEND_DECLARE_ANON_CLASS)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetResultType(IS_VAR)
		opline.GetResult().SetVar(GetTemporaryVariable())
		if !(ZendHashAddPtr(CompilerGlobals.GetClassTable(), lcname, ce)) {

			/* We checked above that the class name is not used. This really shouldn't happen. */

			ZendErrorNoreturn(E_ERROR, "Runtime definition key collision for %s. This is a bug", ZSTR_VAL(name))

			/* We checked above that the class name is not used. This really shouldn't happen. */

		}
	} else {

		/* Generate RTD keys until we find one that isn't in use yet. */

		var key *ZendString = nil
		for {
			ZendTmpStringRelease(key)
			key = ZendBuildRuntimeDefinitionKey(lcname, decl.GetStartLineno())
			if ZendHashAddPtr(CompilerGlobals.GetClassTable(), key, ce) {
				break
			}
		}

		/* RTD key is placed after lcname literal in op1 */

		ZendAddLiteralString(&key)
		opline.SetOpcode(ZEND_DECLARE_CLASS)
		if extends_ast != nil && toplevel != 0 && (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_DELAYED_BINDING) != 0 && (ce.GetCeFlags()&(ZEND_ACC_IMPLEMENT_INTERFACES|ZEND_ACC_IMPLEMENT_TRAITS)) == 0 {
			CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_EARLY_BINDING)
			opline.SetOpcode(ZEND_DECLARE_CLASS_DELAYED)
			opline.SetExtendedValue(ZendAllocCacheSlot())
			opline.SetResultType(IS_UNUSED)
			opline.GetResult().SetOplineNum(-1)
		}
	}
	return opline
}

/* }}} */

func ZendGetImportHt(type_ uint32) *HashTable {
	switch type_ {
	case ZEND_SYMBOL_CLASS:
		if !(FC(imports)) {
			FC(imports) = Emalloc(b.SizeOf("HashTable"))
			ZendHashInit(FC(imports), 8, nil, StrDtor, 0)
		}
		return FC(imports)
	case ZEND_SYMBOL_FUNCTION:
		if !(FC(imports_function)) {
			FC(imports_function) = Emalloc(b.SizeOf("HashTable"))
			ZendHashInit(FC(imports_function), 8, nil, StrDtor, 0)
		}
		return FC(imports_function)
	case ZEND_SYMBOL_CONST:
		if !(FC(imports_const)) {
			FC(imports_const) = Emalloc(b.SizeOf("HashTable"))
			ZendHashInit(FC(imports_const), 8, nil, StrDtor, 0)
		}
		return FC(imports_const)
	default:
		break
	}
	return nil
}

/* }}} */

func ZendGetUseTypeStr(type_ uint32) *byte {
	switch type_ {
	case ZEND_SYMBOL_CLASS:
		return ""
	case ZEND_SYMBOL_FUNCTION:
		return " function"
	case ZEND_SYMBOL_CONST:
		return " const"
	default:
		break
	}
	return " unknown"
}

/* }}} */

func ZendCheckAlreadyInUse(type_ uint32, old_name *ZendString, new_name *ZendString, check_name *ZendString) {
	if ZendStringEqualsCi(old_name, check_name) {
		return
	}
	ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), ZSTR_VAL(old_name), ZSTR_VAL(new_name))
}

/* }}} */

func ZendCompileUse(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var current_ns *ZendString = FC(current_namespace)
	var type_ uint32 = ast.GetAttr()
	var current_import *HashTable = ZendGetImportHt(type_)
	var case_sensitive ZendBool = type_ == ZEND_SYMBOL_CONST
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
					if type_ == T_CLASS && ZendStringEqualsLiteral(new_name, "strict") {
						ZendErrorNoreturn(E_COMPILE_ERROR, "You seem to be trying to use a different language...")
					}
					ZendError(E_WARNING, "The use statement with non-compound name '%s' "+"has no effect", ZSTR_VAL(new_name))
				}
			}
		}
		if case_sensitive != 0 {
			lookup_name = ZendStringCopy(new_name)
		} else {
			lookup_name = ZendStringTolower(new_name)
		}
		if type_ == ZEND_SYMBOL_CLASS && ZendIsReservedClassName(new_name) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use %s as %s because '%s' "+"is a special class name", ZSTR_VAL(old_name), ZSTR_VAL(new_name), ZSTR_VAL(new_name))
		}
		if current_ns != nil {
			var ns_name *ZendString = ZendStringAlloc(ZSTR_LEN(current_ns)+1+ZSTR_LEN(new_name), 0)
			ZendStrTolowerCopy(ZSTR_VAL(ns_name), ZSTR_VAL(current_ns), ZSTR_LEN(current_ns))
			ZSTR_VAL(ns_name)[ZSTR_LEN(current_ns)] = '\\'
			memcpy(ZSTR_VAL(ns_name)+ZSTR_LEN(current_ns)+1, ZSTR_VAL(lookup_name), ZSTR_LEN(lookup_name)+1)
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
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), ZSTR_VAL(old_name), ZSTR_VAL(new_name))
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
		var name *ZendString = Z_STR_P(name_zval)
		var compound_ns *ZendString = ZendConcatNames(ZSTR_VAL(ns), ZSTR_LEN(ns), ZSTR_VAL(name), ZSTR_LEN(name))
		ZendStringReleaseEx(name, 0)
		ZVAL_STR(name_zval, compound_ns)
		inline_use = ZendAstCreateList(1, ZEND_AST_USE, use)
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
		value_node.SetOpType(IS_CONST)
		ZendConstExprToZval(value_zv, value_ast)
		if ZendLookupReservedConst(ZSTR_VAL(unqualified_name), ZSTR_LEN(unqualified_name)) != nil {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot redeclare constant '%s'", ZSTR_VAL(unqualified_name))
		}
		name = ZendPrefixWithNs(unqualified_name)
		name = ZendNewInternedString(name)
		if FC(imports_const) {
			var import_name *ZendString = ZendHashFindPtr(FC(imports_const), unqualified_name)
			if import_name != nil && ZendStringEquals(import_name, name) == 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare const %s because "+"the name is already in use", ZSTR_VAL(name))
			}
		}
		name_node.SetOpType(IS_CONST)
		ZVAL_STR(&name_node.u.constant, name)
		ZendEmitOp(nil, ZEND_DECLARE_CONST, &name_node, &value_node)
		ZendRegisterSeenSymbol(name, ZEND_SYMBOL_CONST)
	}
}

/* }}}*/

func ZendCompileNamespace(ast *ZendAst) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var name *ZendString
	var with_bracket ZendBool = stmt_ast != nil

	/* handle mixed syntax declaration or nested namespaces */

	if !(FC(has_bracketed_namespaces)) {
		if FC(current_namespace) {

			/* previous namespace declarations were unbracketed */

			if with_bracket != 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
			}

			/* previous namespace declarations were unbracketed */

		}
	} else {

		/* previous namespace declarations were bracketed */

		if with_bracket == 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
		} else if FC(current_namespace) || FC(in_namespace) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Namespace declarations cannot be nested")
		}

		/* previous namespace declarations were bracketed */

	}
	if (with_bracket == 0 && !(FC(current_namespace)) || with_bracket != 0 && !(FC(has_bracketed_namespaces))) && CompilerGlobals.GetActiveOpArray().GetLast() > 0 {

		/* ignore ZEND_EXT_STMT and ZEND_TICKS */

		var num uint32 = CompilerGlobals.GetActiveOpArray().GetLast()
		for num > 0 && (CompilerGlobals.GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == ZEND_EXT_STMT || CompilerGlobals.GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == ZEND_TICKS) {
			num--
		}
		if num > 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Namespace declaration statement has to be "+"the very first statement or after any declare call in the script")
		}
	}
	if FC(current_namespace) {
		ZendStringReleaseEx(FC(current_namespace), 0)
	}
	if name_ast != nil {
		name = ZendAstGetStr(name_ast)
		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as namespace name", ZSTR_VAL(name))
		}
		FC(current_namespace) = ZendStringCopy(name)
	} else {
		FC(current_namespace) = nil
	}
	ZendResetImportTables()
	FC(in_namespace) = 1
	if with_bracket != 0 {
		FC(has_bracketed_namespaces) = 1
	}
	if stmt_ast != nil {
		ZendCompileTopStmt(stmt_ast)
		ZendEndNamespace()
	}
}

/* }}} */

func ZendCompileHaltCompiler(ast *ZendAst) {
	var offset_ast *ZendAst = ast.GetChild()[0]
	var offset ZendLong = Z_LVAL_P(ZendAstGetZval(offset_ast))
	var filename *ZendString
	var name *ZendString
	var const_name []byte = "__COMPILER_HALT_OFFSET__"
	if FC(has_bracketed_namespaces) && FC(in_namespace) {
		ZendErrorNoreturn(E_COMPILE_ERROR, "__HALT_COMPILER() can only be used from the outermost scope")
	}
	filename = ZendGetCompiledFilename()
	name = ZendManglePropertyName(const_name, b.SizeOf("const_name")-1, ZSTR_VAL(filename), ZSTR_LEN(filename), 0)
	ZendRegisterLongConstant(ZSTR_VAL(name), ZSTR_LEN(name), offset, CONST_CS, 0)
	ZendStringReleaseEx(name, 0)
}

/* }}} */

func ZendTryCtEvalMagicConst(zv *Zval, ast *ZendAst) ZendBool {
	var op_array *ZendOpArray = CompilerGlobals.GetActiveOpArray()
	var ce *ZendClassEntry = CompilerGlobals.GetActiveClassEntry()
	switch ast.GetAttr() {
	case T_LINE:
		ZVAL_LONG(zv, ast.GetLineno())
		break
	case T_FILE:
		ZVAL_STR_COPY(zv, CompilerGlobals.GetCompiledFilename())
		break
	case T_DIR:
		var filename *ZendString = CompilerGlobals.GetCompiledFilename()
		var dirname *ZendString = ZendStringInit(ZSTR_VAL(filename), ZSTR_LEN(filename), 0)
		ZSTR_LEN(dirname) = ZendDirname(ZSTR_VAL(dirname), ZSTR_LEN(dirname))
		if strcmp(ZSTR_VAL(dirname), ".") == 0 {
			dirname = ZendStringExtend(dirname, MAXPATHLEN, 0)
			ZEND_IGNORE_VALUE(VCWD_GETCWD(ZSTR_VAL(dirname), MAXPATHLEN))
			ZSTR_LEN(dirname) = strlen(ZSTR_VAL(dirname))
		}
		ZVAL_STR(zv, dirname)
		break
	case T_FUNC_C:
		if op_array != nil && op_array.GetFunctionName() != nil {
			ZVAL_STR_COPY(zv, op_array.GetFunctionName())
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
		break
	case T_METHOD_C:

		/* Detect whether we are directly inside a class (e.g. a class constant) and treat
		 * this as not being inside a function. */

		if op_array != nil && ce != nil && op_array.GetScope() == nil && (op_array.GetFnFlags()&ZEND_ACC_CLOSURE) == 0 {
			op_array = nil
		}
		if op_array != nil && op_array.GetFunctionName() != nil {
			if op_array.GetScope() != nil {
				ZVAL_NEW_STR(zv, ZendConcat3(ZSTR_VAL(op_array.GetScope().GetName()), ZSTR_LEN(op_array.GetScope().GetName()), "::", 2, ZSTR_VAL(op_array.GetFunctionName()), ZSTR_LEN(op_array.GetFunctionName())))
			} else {
				ZVAL_STR_COPY(zv, op_array.GetFunctionName())
			}
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
		break
	case T_CLASS_C:
		if ce != nil {
			if (ce.GetCeFlags() & ZEND_ACC_TRAIT) != 0 {
				return 0
			} else {
				ZVAL_STR_COPY(zv, ce.GetName())
			}
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
		break
	case T_TRAIT_C:
		if ce != nil && (ce.GetCeFlags()&ZEND_ACC_TRAIT) != 0 {
			ZVAL_STR_COPY(zv, ce.GetName())
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
		break
	case T_NS_C:
		if FC(current_namespace) {
			ZVAL_STR_COPY(zv, FC(current_namespace))
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
		break
	default:
		break
	}
	return 1
}

/* }}} */

func ZendBinaryOpProducesNumericStringError(opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	if !(opcode == ZEND_ADD || opcode == ZEND_SUB || opcode == ZEND_MUL || opcode == ZEND_DIV || opcode == ZEND_POW || opcode == ZEND_MOD || opcode == ZEND_SL || opcode == ZEND_SR || opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) {
		return 0
	}

	/* While basic arithmetic operators always produce numeric string errors,
	 * bitwise operators don't produce errors if both operands are strings */

	if (opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) && Z_TYPE_P(op1) == IS_STRING && Z_TYPE_P(op2) == IS_STRING {
		return 0
	}
	if Z_TYPE_P(op1) == IS_STRING && IsNumericString(Z_STRVAL_P(op1), Z_STRLEN_P(op1), nil, nil, 0) == 0 {
		return 1
	}
	if Z_TYPE_P(op2) == IS_STRING && IsNumericString(Z_STRVAL_P(op2), Z_STRLEN_P(op2), nil, nil, 0) == 0 {
		return 1
	}
	return 0
}

/* }}} */

func ZendBinaryOpProducesArrayConversionError(opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	if opcode == ZEND_CONCAT && (Z_TYPE_P(op1) == IS_ARRAY || Z_TYPE_P(op2) == IS_ARRAY) {
		return 1
	}
	return 0
}

/* }}} */

func ZendTryCtEvalBinaryOp(result *Zval, opcode uint32, op1 *Zval, op2 *Zval) ZendBool {
	var fn BinaryOpType = GetBinaryOp(opcode)

	/* don't evaluate division by zero at compile-time */

	if (opcode == ZEND_DIV || opcode == ZEND_MOD) && ZvalGetLong(op2) == 0 {
		return 0
	} else if (opcode == ZEND_SL || opcode == ZEND_SR) && ZvalGetLong(op2) < 0 {
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
	ZVAL_LONG(&left, b.Cond(kind == ZEND_AST_UNARY_PLUS, 1, -1))
	return ZendTryCtEvalBinaryOp(result, ZEND_MUL, &left, op)
}

/* }}} */

func ZendCtEvalGreater(result *Zval, kind ZendAstKind, op1 *Zval, op2 *Zval) {
	var fn BinaryOpType = b.Cond(kind == ZEND_AST_GREATER, IsSmallerFunction, IsSmallerOrEqualFunction)
	fn(result, op2, op1)
}

/* }}} */

func ZendTryCtEvalArray(result *Zval, ast *ZendAst) ZendBool {
	var list *ZendAstList = ZendAstGetList(ast)
	var last_elem_ast *ZendAst = nil
	var i uint32
	var is_constant ZendBool = 1
	if ast.GetAttr() == ZEND_ARRAY_SYNTAX_LIST {
		ZendError(E_COMPILE_ERROR, "Cannot use list() as standalone expression")
	}

	/* First ensure that *all* child nodes are constant and by-val */

	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		if elem_ast == nil {

			/* Report error at line of last non-empty element */

			if last_elem_ast != nil {
				CompilerGlobals.SetZendLineno(ZendAstGetLineno(last_elem_ast))
			}
			ZendError(E_COMPILE_ERROR, "Cannot use empty array elements in arrays")
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
		ZVAL_EMPTY_ARRAY(result)
		return 1
	}
	ArrayInitSize(result, list.GetChildren())
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var value_ast *ZendAst = elem_ast.GetChild()[0]
		var key_ast *ZendAst
		var value *Zval = ZendAstGetZval(value_ast)
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			if Z_TYPE_P(value) == IS_ARRAY {
				var ht *HashTable = Z_ARRVAL_P(value)
				var val *Zval
				var key *ZendString
				for {
					var __ht *HashTable = ht
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
							continue
						}
						key = _p.GetKey()
						val = _z
						if key != nil {
							ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot unpack array with string keys")
						}
						if ZendHashNextIndexInsert(Z_ARRVAL_P(result), val) == nil {
							ZvalPtrDtor(result)
							return 0
						}
						Z_TRY_ADDREF_P(val)
					}
					break
				}
				continue
			} else {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Only arrays and Traversables can be unpacked")
			}
		}
		Z_TRY_ADDREF_P(value)
		key_ast = elem_ast.GetChild()[1]
		if key_ast != nil {
			var key *Zval = ZendAstGetZval(key_ast)
			switch Z_TYPE_P(key) {
			case IS_LONG:
				ZendHashIndexUpdate(Z_ARRVAL_P(result), Z_LVAL_P(key), value)
				break
			case IS_STRING:
				ZendSymtableUpdate(Z_ARRVAL_P(result), Z_STR_P(key), value)
				break
			case IS_DOUBLE:
				ZendHashIndexUpdate(Z_ARRVAL_P(result), ZendDvalToLval(Z_DVAL_P(key)), value)
				break
			case IS_FALSE:
				ZendHashIndexUpdate(Z_ARRVAL_P(result), 0, value)
				break
			case IS_TRUE:
				ZendHashIndexUpdate(Z_ARRVAL_P(result), 1, value)
				break
			case IS_NULL:
				ZendHashUpdate(Z_ARRVAL_P(result), ZSTR_EMPTY_ALLOC(), value)
				break
			default:
				ZendErrorNoreturn(E_COMPILE_ERROR, "Illegal offset type")
				break
			}
		} else {
			if ZendHashNextIndexInsert(Z_ARRVAL_P(result), value) == nil {
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
	if (opcode == ZEND_ADD || opcode == ZEND_SUB) && left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == ZEND_CONCAT {
		ZendError(E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '+'/'-' will change in PHP 8: '+'/'-' will take a higher precedence")
	}
	if (opcode == ZEND_SL || opcode == ZEND_SR) && (left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == ZEND_CONCAT || right_ast.GetKind() == ZEND_AST_BINARY_OP && right_ast.GetAttr() == ZEND_CONCAT) {
		ZendError(E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '>>'/'<<' will change in PHP 8: '<<'/'>>' will take a higher precedence")
	}
	if opcode == ZEND_PARENTHESIZED_CONCAT {
		opcode = ZEND_CONCAT
	}
	var left_node Znode
	var right_node Znode
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		if ZendTryCtEvalBinaryOp(&result.u.constant, opcode, &left_node.u.constant, &right_node.u.constant) != 0 {
			result.SetOpType(IS_CONST)
			ZvalPtrDtor(&left_node.u.constant)
			ZvalPtrDtor(&right_node.u.constant)
			return
		}
	}
	for {
		if opcode == ZEND_IS_EQUAL || opcode == ZEND_IS_NOT_EQUAL {
			if left_node.GetOpType() == IS_CONST {
				if Z_TYPE(left_node.GetConstant()) == IS_FALSE {
					if opcode == ZEND_IS_NOT_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				} else if Z_TYPE(left_node.GetConstant()) == IS_TRUE {
					if opcode == ZEND_IS_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				}
			} else if right_node.GetOpType() == IS_CONST {
				if Z_TYPE(right_node.GetConstant()) == IS_FALSE {
					if opcode == ZEND_IS_NOT_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				} else if Z_TYPE(right_node.GetConstant()) == IS_TRUE {
					if opcode == ZEND_IS_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				}
			}
		}
		if opcode == ZEND_CONCAT {

			/* convert constant operands to strings at compile-time */

			if left_node.GetOpType() == IS_CONST {
				if Z_TYPE(left_node.GetConstant()) == IS_ARRAY {
					ZendEmitOpTmp(&left_node, ZEND_CAST, &left_node, nil).SetExtendedValue(IS_STRING)
				} else {
					ConvertToString(&left_node.u.constant)
				}
			}
			if right_node.GetOpType() == IS_CONST {
				if Z_TYPE(right_node.GetConstant()) == IS_ARRAY {
					ZendEmitOpTmp(&right_node, ZEND_CAST, &right_node, nil).SetExtendedValue(IS_STRING)
				} else {
					ConvertToString(&right_node.u.constant)
				}
			}
			if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
				opcode = ZEND_FAST_CONCAT
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
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_GREATER || ast.GetKind() == ZEND_AST_GREATER_EQUAL)
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
		ZendCtEvalGreater(&result.u.constant, ast.GetKind(), &left_node.u.constant, &right_node.u.constant)
		ZvalPtrDtor(&left_node.u.constant)
		ZvalPtrDtor(&right_node.u.constant)
		return
	}
	ZendEmitOpTmp(result, b.Cond(ast.GetKind() == ZEND_AST_GREATER, ZEND_IS_SMALLER, ZEND_IS_SMALLER_OR_EQUAL), &right_node, &left_node)
}

/* }}} */

func ZendCompileUnaryOp(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var opcode uint32 = ast.GetAttr()
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
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
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_UNARY_PLUS || ast.GetKind() == ZEND_AST_UNARY_MINUS)
	ZendCompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == IS_CONST {
		if ZendTryCtEvalUnaryPm(&result.u.constant, ast.GetKind(), &expr_node.u.constant) != 0 {
			result.SetOpType(IS_CONST)
			ZvalPtrDtor(&expr_node.u.constant)
			return
		}
	}
	lefthand_node.SetOpType(IS_CONST)
	ZVAL_LONG(&lefthand_node.u.constant, b.Cond(ast.GetKind() == ZEND_AST_UNARY_PLUS, 1, -1))
	ZendEmitOpTmp(result, ZEND_MUL, &lefthand_node, &expr_node)
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
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_AND || ast.GetKind() == ZEND_AST_OR)
	ZendCompileExpr(&left_node, left_ast)
	if left_node.GetOpType() == IS_CONST {
		if ast.GetKind() == ZEND_AST_AND && ZendIsTrue(&left_node.u.constant) == 0 || ast.GetKind() == ZEND_AST_OR && ZendIsTrue(&left_node.u.constant) != 0 {
			result.SetOpType(IS_CONST)
			ZVAL_BOOL(&result.u.constant, ZendIsTrue(&left_node.u.constant))
		} else {
			ZendCompileExpr(&right_node, right_ast)
			if right_node.GetOpType() == IS_CONST {
				result.SetOpType(IS_CONST)
				ZVAL_BOOL(&result.u.constant, ZendIsTrue(&right_node.u.constant))
				ZvalPtrDtor(&right_node.u.constant)
			} else {
				ZendEmitOpTmp(result, ZEND_BOOL, &right_node, nil)
			}
		}
		ZvalPtrDtor(&left_node.u.constant)
		return
	}
	opnum_jmpz = GetNextOpNumber()
	opline_jmpz = ZendEmitOp(nil, b.Cond(ast.GetKind() == ZEND_AST_AND, ZEND_JMPZ_EX, ZEND_JMPNZ_EX), &left_node, nil)
	if left_node.GetOpType() == IS_TMP_VAR {
		opline_jmpz.SetResultType(&left_node.GetOpType())
		if &left_node.GetOpType() == IS_CONST {
			opline_jmpz.GetResult().SetConstant(ZendAddLiteral(&(&left_node).u.constant))
		} else {
			opline_jmpz.SetResult(&left_node.GetOp())
		}
	} else {
		opline_jmpz.GetResult().SetVar(GetTemporaryVariable())
		opline_jmpz.SetResultType(IS_TMP_VAR)
	}
	result.SetOpType(opline_jmpz.GetResultType())
	if result.GetOpType() == IS_CONST {
		ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline_jmpz.GetResult()))
	} else {
		result.SetOp(opline_jmpz.GetResult())
	}
	ZendCompileExpr(&right_node, right_ast)
	opline_bool = ZendEmitOp(nil, ZEND_BOOL, &right_node, nil)
	opline_bool.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline_bool.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline_bool.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmpz)
}

/* }}} */

func ZendCompilePostIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_POST_INC || ast.GetKind() == ZEND_AST_POST_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.GetKind() == ZEND_AST_PROP {
		var opline *ZendOp = ZendCompileProp(nil, var_ast, BP_VAR_RW, 0)
		if ast.GetKind() == ZEND_AST_POST_INC {
			opline.SetOpcode(ZEND_POST_INC_OBJ)
		} else {
			opline.SetOpcode(ZEND_POST_DEC_OBJ)
		}
		ZendMakeTmpResult(result, opline)
	} else if var_ast.GetKind() == ZEND_AST_STATIC_PROP {
		var opline *ZendOp = ZendCompileStaticProp(nil, var_ast, BP_VAR_RW, 0, 0)
		if ast.GetKind() == ZEND_AST_POST_INC {
			opline.SetOpcode(ZEND_POST_INC_STATIC_PROP)
		} else {
			opline.SetOpcode(ZEND_POST_DEC_STATIC_PROP)
		}
		ZendMakeTmpResult(result, opline)
	} else {
		var var_node Znode
		ZendCompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendEmitOpTmp(result, b.Cond(ast.GetKind() == ZEND_AST_POST_INC, ZEND_POST_INC, ZEND_POST_DEC), &var_node, nil)
	}
}

/* }}} */

func ZendCompilePreIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_PRE_INC || ast.GetKind() == ZEND_AST_PRE_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.GetKind() == ZEND_AST_PROP {
		var opline *ZendOp = ZendCompileProp(result, var_ast, BP_VAR_RW, 0)
		if ast.GetKind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(ZEND_PRE_INC_OBJ)
		} else {
			opline.SetOpcode(ZEND_PRE_DEC_OBJ)
		}
	} else if var_ast.GetKind() == ZEND_AST_STATIC_PROP {
		var opline *ZendOp = ZendCompileStaticProp(result, var_ast, BP_VAR_RW, 0, 0)
		if ast.GetKind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(ZEND_PRE_INC_STATIC_PROP)
		} else {
			opline.SetOpcode(ZEND_PRE_DEC_STATIC_PROP)
		}
	} else {
		var var_node Znode
		ZendCompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendEmitOp(result, b.Cond(ast.GetKind() == ZEND_AST_PRE_INC, ZEND_PRE_INC, ZEND_PRE_DEC), &var_node, nil)
	}
}

/* }}} */

func ZendCompileCast(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	var opline *ZendOp
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOpTmp(result, ZEND_CAST, &expr_node, nil)
	opline.SetExtendedValue(ast.GetAttr())
	if ast.GetAttr() == IS_NULL {
		ZendError(E_DEPRECATED, "The (unset) cast is deprecated")
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
	ZEND_ASSERT(ast.GetChild()[1] == nil)
	ZendCompileExpr(&cond_node, cond_ast)
	opnum_jmp_set = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_JMP_SET, &cond_node, nil)
	ZendCompileExpr(&false_node, false_ast)
	opline_qm_assign = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &false_node, nil)
	opline_qm_assign.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
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
	if cond_ast.GetKind() == ZEND_AST_CONDITIONAL && cond_ast.GetAttr() != ZEND_PARENTHESIZED_CONDITIONAL {
		if cond_ast.GetChild()[1] != nil {
			if true_ast != nil {
				ZendError(E_DEPRECATED, "Unparenthesized `a ? b : c ? d : e` is deprecated. "+"Use either `(a ? b : c) ? d : e` or `a ? b : (c ? d : e)`")
			} else {
				ZendError(E_DEPRECATED, "Unparenthesized `a ? b : c ?: d` is deprecated. "+"Use either `(a ? b : c) ?: d` or `a ? b : (c ?: d)`")
			}
		} else {
			if true_ast != nil {
				ZendError(E_DEPRECATED, "Unparenthesized `a ?: b ? c : d` is deprecated. "+"Use either `(a ?: b) ? c : d` or `a ?: (b ? c : d)`")
			}
		}
	}
	if true_ast == nil {
		ZendCompileShorthandConditional(result, ast)
		return
	}
	ZendCompileExpr(&cond_node, cond_ast)
	opnum_jmpz = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
	ZendCompileExpr(&true_node, true_ast)
	ZendEmitOpTmp(result, ZEND_QM_ASSIGN, &true_node, nil)
	opnum_jmp = ZendEmitJump(0)
	ZendUpdateJumpTargetToNext(opnum_jmpz)
	ZendCompileExpr(&false_node, false_ast)
	opline_qm_assign2 = ZendEmitOp(nil, ZEND_QM_ASSIGN, &false_node, nil)
	opline_qm_assign2.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
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
	ZendCompileVar(&expr_node, expr_ast, BP_VAR_IS, 0)
	opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_COALESCE, &expr_node, nil)
	ZendCompileExpr(&default_node, default_ast)
	opline = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &default_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
}

/* }}} */

func ZnodeDtor(zv *Zval) {
	var node *Znode = Z_PTR_P(zv)
	if node.GetOpType() == IS_CONST {
		ZvalPtrDtorNogc(&node.u.constant)
	}
	Efree(node)
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

	var orig_memoized_exprs *HashTable = CompilerGlobals.GetMemoizedExprs()
	var orig_memoize_mode int = CompilerGlobals.GetMemoizeMode()
	ZendEnsureWritableVariable(var_ast)
	if IsThisFetch(var_ast) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ALLOC_HASHTABLE(CompilerGlobals.GetMemoizedExprs())
	ZendHashInit(CompilerGlobals.GetMemoizedExprs(), 0, nil, ZnodeDtor, 0)
	CompilerGlobals.SetMemoizeMode(ZEND_MEMOIZE_COMPILE)
	ZendCompileVar(&var_node_is, var_ast, BP_VAR_IS, 0)
	coalesce_opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_COALESCE, &var_node_is, nil)
	CompilerGlobals.SetMemoizeMode(ZEND_MEMOIZE_NONE)
	ZendCompileExpr(&default_node, default_ast)
	CompilerGlobals.SetMemoizeMode(ZEND_MEMOIZE_FETCH)
	ZendCompileVar(&var_node_w, var_ast, BP_VAR_W, 0)

	/* Reproduce some of the zend_compile_assign() opcode fixup logic here. */

	opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[CompilerGlobals.GetActiveOpArray().GetLast()-1]
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		ZendEmitOp(&assign_node, ZEND_ASSIGN, &var_node_w, &default_node)
		break
	case ZEND_AST_STATIC_PROP:
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	case ZEND_AST_DIM:
		opline.SetOpcode(ZEND_ASSIGN_DIM)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	case ZEND_AST_PROP:
		opline.SetOpcode(ZEND_ASSIGN_OBJ)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
		break
	default:
		break
	}
	opline = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &assign_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
	} else {
		opline.SetResult(result.GetOp())
	}
	for {
		var __ht *HashTable = CompilerGlobals.GetMemoizedExprs()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			node = Z_PTR_P(_z)
			if node.GetOpType() == IS_TMP_VAR || node.GetOpType() == IS_VAR {
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
			var __ht *HashTable = CompilerGlobals.GetMemoizedExprs()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				node = Z_PTR_P(_z)
				if node.GetOpType() == IS_TMP_VAR || node.GetOpType() == IS_VAR {
					ZendEmitOp(nil, ZEND_FREE, node, nil)
				}
			}
			break
		}
		ZendUpdateJumpTargetToNext(jump_opnum)
	} else {
		ZendUpdateJumpTargetToNext(coalesce_opnum)
	}
	ZendHashDestroy(CompilerGlobals.GetMemoizedExprs())
	FREE_HASHTABLE(CompilerGlobals.GetMemoizedExprs())
	CompilerGlobals.SetMemoizedExprs(orig_memoized_exprs)
	CompilerGlobals.SetMemoizeMode(orig_memoize_mode)
}

/* }}} */

func ZendCompilePrint(result *Znode, ast *ZendAst) {
	var opline *ZendOp
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, ZEND_ECHO, &expr_node, nil)
	opline.SetExtendedValue(1)
	result.SetOpType(IS_CONST)
	ZVAL_LONG(&result.u.constant, 1)
}

/* }}} */

func ZendCompileExit(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	if expr_ast != nil {
		var expr_node Znode
		ZendCompileExpr(&expr_node, expr_ast)
		ZendEmitOp(nil, ZEND_EXIT, &expr_node, nil)
	} else {
		ZendEmitOp(nil, ZEND_EXIT, nil, nil)
	}
	result.SetOpType(IS_CONST)
	ZVAL_BOOL(&result.u.constant, 1)
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
	var returns_by_ref ZendBool = (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0
	ZendMarkFunctionAsGenerator()
	if key_ast != nil {
		ZendCompileExpr(&key_node, key_ast)
		key_node_ptr = &key_node
	}
	if value_ast != nil {
		if returns_by_ref != 0 && ZendIsVariable(value_ast) != 0 {
			ZendCompileVar(&value_node, value_ast, BP_VAR_W, 1)
		} else {
			ZendCompileExpr(&value_node, value_ast)
		}
		value_node_ptr = &value_node
	}
	opline = ZendEmitOp(result, ZEND_YIELD, value_node_ptr, key_node_ptr)
	if value_ast != nil && returns_by_ref != 0 && ZendIsCall(value_ast) != 0 {
		opline.SetExtendedValue(ZEND_RETURNS_FUNCTION)
	}
}

/* }}} */

func ZendCompileYieldFrom(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendMarkFunctionAsGenerator()
	if (CompilerGlobals.GetActiveOpArray().GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use \"yield from\" inside a by-reference generator")
	}
	ZendCompileExpr(&expr_node, expr_ast)
	ZendEmitOpTmp(result, ZEND_YIELD_FROM, &expr_node, nil)
}

/* }}} */

func ZendCompileInstanceof(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var class_ast *ZendAst = ast.GetChild()[1]
	var obj_node Znode
	var class_node Znode
	var opline *ZendOp
	ZendCompileExpr(&obj_node, obj_ast)
	if obj_node.GetOpType() == IS_CONST {
		ZendDoFree(&obj_node)
		result.SetOpType(IS_CONST)
		ZVAL_FALSE(&result.u.constant)
		return
	}
	ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_NO_AUTOLOAD|ZEND_FETCH_CLASS_EXCEPTION)
	opline = ZendEmitOpTmp(result, ZEND_INSTANCEOF, &obj_node, nil)
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(Z_STR(class_node.GetConstant())))
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {
		opline.SetOp2Type(&class_node.GetOpType())
		if &class_node.GetOpType() == IS_CONST {
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
	opline = ZendEmitOp(result, ZEND_INCLUDE_OR_EVAL, &expr_node, nil)
	opline.SetExtendedValue(ast.GetAttr())
	ZendDoExtendedFcallEnd()
}

/* }}} */

func ZendCompileIssetOrEmpty(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var var_node Znode
	var opline *ZendOp = nil
	ZEND_ASSERT(ast.GetKind() == ZEND_AST_ISSET || ast.GetKind() == ZEND_AST_EMPTY)
	if ZendIsVariable(var_ast) == 0 {
		if ast.GetKind() == ZEND_AST_EMPTY {

			/* empty(expr) can be transformed to !expr */

			var not_ast *ZendAst = ZendAstCreateEx(ZEND_AST_UNARY_OP, ZEND_BOOL_NOT, var_ast)
			ZendCompileExpr(result, not_ast)
			return
		} else {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use isset() on the result of an expression "+"(you can use \"null !== expression\" instead)")
		}
	}
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		if IsThisFetch(var_ast) != 0 {
			opline = ZendEmitOp(result, ZEND_ISSET_ISEMPTY_THIS, nil, nil)
			CompilerGlobals.GetActiveOpArray().SetFnFlags(CompilerGlobals.GetActiveOpArray().GetFnFlags() | ZEND_ACC_USES_THIS)
		} else if ZendTryCompileCv(&var_node, var_ast) == SUCCESS {
			opline = ZendEmitOp(result, ZEND_ISSET_ISEMPTY_CV, &var_node, nil)
		} else {
			opline = ZendCompileSimpleVarNoCv(result, var_ast, BP_VAR_IS, 0)
			opline.SetOpcode(ZEND_ISSET_ISEMPTY_VAR)
		}
		break
	case ZEND_AST_DIM:
		opline = ZendCompileDim(result, var_ast, BP_VAR_IS)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_DIM_OBJ)
		break
	case ZEND_AST_PROP:
		opline = ZendCompileProp(result, var_ast, BP_VAR_IS, 0)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_PROP_OBJ)
		break
	case ZEND_AST_STATIC_PROP:
		opline = ZendCompileStaticProp(result, var_ast, BP_VAR_IS, 0, 0)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_STATIC_PROP)
		break
	default:
		break
	}
	opline.SetResultType(IS_TMP_VAR)
	result.SetOpType(opline.GetResultType())
	if ast.GetKind() != ZEND_AST_ISSET {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_ISEMPTY)
	}
}

/* }}} */

func ZendCompileSilence(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var silence_node Znode
	ZendEmitOpTmp(&silence_node, ZEND_BEGIN_SILENCE, nil, nil)
	if expr_ast.GetKind() == ZEND_AST_VAR {

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

		ZendCompileSimpleVarNoCv(result, expr_ast, BP_VAR_R, 0)

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

	} else {
		ZendCompileExpr(result, expr_ast)
	}
	ZendEmitOp(nil, ZEND_END_SILENCE, &silence_node, nil)
}

/* }}} */

func ZendCompileShellExec(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var fn_name Zval
	var name_ast *ZendAst
	var args_ast *ZendAst
	var call_ast *ZendAst
	ZVAL_STRING(&fn_name, "shell_exec")
	name_ast = ZendAstCreateZval(&fn_name)
	args_ast = ZendAstCreateList(1, ZEND_AST_ARG_LIST, expr_ast)
	call_ast = ZendAstCreate(ZEND_AST_CALL, name_ast, args_ast)
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
		result.SetOpType(IS_CONST)
		return
	}

	/* Empty arrays are handled at compile-time */

	ZEND_ASSERT(list.GetChildren() > 0)
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var value_ast *ZendAst
		var key_ast *ZendAst
		var by_ref ZendBool
		var value_node Znode
		var key_node Znode
		var key_node_ptr *Znode = nil
		if elem_ast == nil {
			ZendError(E_COMPILE_ERROR, "Cannot use empty array elements in arrays")
		}
		value_ast = elem_ast.GetChild()[0]
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			ZendCompileExpr(&value_node, value_ast)
			if i == 0 {
				opnum_init = GetNextOpNumber()
				opline = ZendEmitOpTmp(result, ZEND_INIT_ARRAY, nil, nil)
			}
			opline = ZendEmitOp(nil, ZEND_ADD_ARRAY_UNPACK, &value_node, nil)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == IS_CONST {
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
			ZendCompileVar(&value_node, value_ast, BP_VAR_W, 1)
		} else {
			ZendCompileExpr(&value_node, value_ast)
		}
		if i == 0 {
			opnum_init = GetNextOpNumber()
			opline = ZendEmitOpTmp(result, ZEND_INIT_ARRAY, &value_node, key_node_ptr)
			opline.SetExtendedValue(list.GetChildren() << ZEND_ARRAY_SIZE_SHIFT)
		} else {
			opline = ZendEmitOp(nil, ZEND_ADD_ARRAY_ELEMENT, &value_node, key_node_ptr)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(&result.u.constant))
			} else {
				opline.SetResult(result.GetOp())
			}
		}
		opline.SetExtendedValue(opline.GetExtendedValue() | by_ref)
		if key_ast != nil && key_node.GetOpType() == IS_CONST && Z_TYPE(key_node.GetConstant()) == IS_STRING {
			packed = 0
		}
	}

	/* Add a flag to INIT_ARRAY if we know this array cannot be packed */

	if packed == 0 {
		ZEND_ASSERT(opnum_init != uint32_t-1)
		opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[opnum_init]
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_ARRAY_NOT_PACKED)
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
	if ZendStringEqualsLiteral(resolved_name, "__COMPILER_HALT_OFFSET__") || name_ast.GetAttr() != ZEND_NAME_RELATIVE && ZendStringEqualsLiteral(orig_name, "__COMPILER_HALT_OFFSET__") {
		var last *ZendAst = CompilerGlobals.GetAst()
		for last != nil && last.GetKind() == ZEND_AST_STMT_LIST {
			var list *ZendAstList = ZendAstGetList(last)
			if list.GetChildren() == 0 {
				break
			}
			last = list.GetChild()[list.GetChildren()-1]
		}
		if last != nil && last.GetKind() == ZEND_AST_HALT_COMPILER {
			result.SetOpType(IS_CONST)
			ZVAL_LONG(&result.u.constant, Z_LVAL_P(ZendAstGetZval(last.GetChild()[0])))
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
	}
	if ZendTryCtEvalConst(&result.u.constant, resolved_name, is_fully_qualified) != 0 {
		result.SetOpType(IS_CONST)
		ZendStringReleaseEx(resolved_name, 0)
		return
	}
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CONSTANT, nil, nil)
	opline.SetOp2Type(IS_CONST)
	if is_fully_qualified != 0 {
		opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name, 0))
	} else {
		opline.GetOp1().SetNum(IS_CONSTANT_UNQUALIFIED)
		if FC(current_namespace) {
			opline.GetOp1().SetNum(opline.GetOp1().GetNum() | IS_CONSTANT_IN_NAMESPACE)
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
			result.SetOpType(IS_CONST)
			ZendStringReleaseEx(resolved_name, 0)
			return
		}
		ZendStringReleaseEx(resolved_name, 0)
	}
	ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	ZendCompileExpr(&const_node, const_ast)
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_CONSTANT, nil, &const_node)
	ZendSetClassNameOp1(opline, &class_node)
	opline.SetExtendedValue(ZendAllocCacheSlots(2))
}

/* }}} */

func ZendCompileClassName(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	if ZendTryCompileConstExprResolveClassName(&result.u.constant, class_ast) != 0 {
		result.SetOpType(IS_CONST)
		return
	}
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_NAME, nil, nil)
	opline.GetOp1().SetNum(ZendGetClassFetchType(ZendAstGetStr(class_ast)))
}

/* }}} */

func ZendCompileRopeAddEx(opline *ZendOp, result *Znode, num uint32, elem_node *Znode) *ZendOp {
	if num == 0 {
		result.SetOpType(IS_TMP_VAR)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(ZEND_ROPE_INIT)
	} else {
		opline.SetOpcode(ZEND_ROPE_ADD)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == IS_CONST {
		opline.GetOp2().SetConstant(ZendAddLiteral(&elem_node.u.constant))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
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
		result.SetOpType(IS_TMP_VAR)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(ZEND_ROPE_INIT)
	} else {
		opline.SetOpcode(ZEND_ROPE_ADD)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(&result.u.constant))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == IS_CONST {
		opline.GetOp2().SetConstant(ZendAddLiteral(&elem_node.u.constant))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
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
	ZEND_ASSERT(list.GetChildren() > 0)
	j = 0
	last_const_node.SetOpType(IS_UNUSED)
	for i = 0; i < list.GetChildren(); i++ {
		ZendCompileExpr(&elem_node, list.GetChild()[i])
		if elem_node.GetOpType() == IS_CONST {
			ConvertToString(&elem_node.u.constant)
			if Z_STRLEN(elem_node.GetConstant()) == 0 {
				ZvalPtrDtor(&elem_node.u.constant)
			} else if last_const_node.GetOpType() == IS_CONST {
				ConcatFunction(&last_const_node.u.constant, &last_const_node.u.constant, &elem_node.u.constant)
				ZvalPtrDtor(&elem_node.u.constant)
			} else {
				last_const_node.SetOpType(IS_CONST)
				ZVAL_COPY_VALUE(&last_const_node.u.constant, &elem_node.u.constant)

				/* Reserve place for ZEND_ROPE_ADD instruction */

				reserved_op_number = GetNextOpNumber()
				opline = GetNextOp()
				opline.SetOpcode(ZEND_NOP)
			}
			continue
		} else {
			if j == 0 {
				if last_const_node.GetOpType() == IS_CONST {
					rope_init_lineno = reserved_op_number
				} else {
					rope_init_lineno = GetNextOpNumber()
				}
			}
			if last_const_node.GetOpType() == IS_CONST {
				opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[reserved_op_number]
				ZendCompileRopeAddEx(opline, result, b.PostInc(&j), &last_const_node)
				last_const_node.SetOpType(IS_UNUSED)
			}
			opline = ZendCompileRopeAdd(result, b.PostInc(&j), &elem_node)
		}
	}
	if j == 0 {
		result.SetOpType(IS_CONST)
		if last_const_node.GetOpType() == IS_CONST {
			ZVAL_COPY_VALUE(&result.u.constant, &last_const_node.u.constant)
		} else {
			ZVAL_EMPTY_STRING(&result.u.constant)
		}
		CompilerGlobals.GetActiveOpArray().SetLast(reserved_op_number - 1)
		return
	} else if last_const_node.GetOpType() == IS_CONST {
		opline = &(CompilerGlobals.GetActiveOpArray()).opcodes[reserved_op_number]
		opline = ZendCompileRopeAddEx(opline, result, b.PostInc(&j), &last_const_node)
	}
	init_opline = CompilerGlobals.GetActiveOpArray().GetOpcodes() + rope_init_lineno
	if j == 1 {
		if opline.GetOp2Type() == IS_CONST {
			result.SetOpType(opline.GetOp2Type())
			if result.GetOpType() == IS_CONST {
				ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetOp2()))
			} else {
				result.SetOp(opline.GetOp2())
			}
			MAKE_NOP(opline)
		} else {
			opline.SetOpcode(ZEND_CAST)
			opline.SetExtendedValue(IS_STRING)
			opline.SetOp1Type(opline.GetOp2Type())
			opline.SetOp1(opline.GetOp2())
			opline.SetResultType(IS_TMP_VAR)
			opline.GetResult().SetVar(GetTemporaryVariable())
			opline.SetOp2Type(IS_UNUSED)
			result.SetOpType(opline.GetResultType())
			if result.GetOpType() == IS_CONST {
				ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetResult()))
			} else {
				result.SetOp(opline.GetResult())
			}
		}
	} else if j == 2 {
		opline.SetOpcode(ZEND_FAST_CONCAT)
		opline.SetExtendedValue(0)
		opline.SetOp1Type(init_opline.GetOp2Type())
		opline.SetOp1(init_opline.GetOp2())
		opline.SetResultType(IS_TMP_VAR)
		opline.GetResult().SetVar(GetTemporaryVariable())
		MAKE_NOP(init_opline)
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == IS_CONST {
			ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetResult()))
		} else {
			result.SetOp(opline.GetResult())
		}
	} else {
		var var_ uint32
		init_opline.SetExtendedValue(j)
		opline.SetOpcode(ZEND_ROPE_END)
		opline.GetResult().SetVar(GetTemporaryVariable())
		opline.GetOp1().SetVar(GetTemporaryVariable())
		var_ = opline.GetOp1().GetVar()
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == IS_CONST {
			ZVAL_COPY_VALUE(&result.u.constant, CT_CONSTANT(opline.GetResult()))
		} else {
			result.SetOp(opline.GetResult())
		}

		/* Allocates the necessary number of zval slots to keep the rope */

		i = (j*b.SizeOf("zend_string *") + (b.SizeOf("zval") - 1)) / b.SizeOf("zval")
		for i > 1 {
			GetTemporaryVariable()
			i--
		}

		/* Update all the previous opcodes to use the same variable */

		for opline != init_opline {
			opline--
			if opline.GetOpcode() == ZEND_ROPE_ADD && opline.GetResult().GetVar() == uint32_t-1 {
				opline.GetOp1().SetVar(var_)
				opline.GetResult().SetVar(var_)
			} else if opline.GetOpcode() == ZEND_ROPE_INIT && opline.GetResult().GetVar() == uint32_t-1 {
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
		result.SetOpType(IS_CONST)
		return
	}
	ZEND_ASSERT(ast.GetAttr() == T_CLASS_C && CompilerGlobals.GetActiveClassEntry() != nil && (CompilerGlobals.GetActiveClassEntry().GetCeFlags()&ZEND_ACC_TRAIT) != 0)
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_NAME, nil, nil)
	opline.GetOp1().SetNum(ZEND_FETCH_CLASS_SELF)
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
		ZendErrorNoreturn(E_COMPILE_ERROR, "Dynamic class names are not allowed in compile-time class constant references")
	}
	class_name = ZendAstGetStr(class_ast)
	fetch_type = ZendGetClassFetchType(class_name)
	if ZEND_FETCH_CLASS_STATIC == fetch_type {
		ZendErrorNoreturn(E_COMPILE_ERROR, "\"static::\" is not allowed in compile-time constants")
	}
	if ZEND_FETCH_CLASS_DEFAULT == fetch_type {
		class_name = ZendResolveClassNameAst(class_ast)
	} else {
		ZendStringAddref(class_name)
	}
	name = ZendConcat3(ZSTR_VAL(class_name), ZSTR_LEN(class_name), "::", 2, ZSTR_VAL(const_name), ZSTR_LEN(const_name))
	ZendAstDestroy(ast)
	ZendStringReleaseEx(class_name, 0)
	*ast_ptr = ZendAstCreateConstant(name, fetch_type|ZEND_FETCH_CLASS_EXCEPTION)
}

/* }}} */

func ZendCompileConstExprClassName(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var class_ast *ZendAst = ast.GetChild()[0]
	var class_name *ZendString = ZendAstGetStr(class_ast)
	var fetch_type uint32 = ZendGetClassFetchType(class_name)
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:

	case ZEND_FETCH_CLASS_PARENT:

		/* For the const-eval representation store the fetch type instead of the name. */

		ZendStringRelease(class_name)
		ast.GetChild()[0] = nil
		ast.SetAttr(fetch_type)
		return
	case ZEND_FETCH_CLASS_STATIC:
		ZendErrorNoreturn(E_COMPILE_ERROR, "static::class cannot be used for compile-time class name resolution")
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
	*ast_ptr = ZendAstCreateConstant(resolved_name, b.Cond(is_fully_qualified == 0, IS_CONSTANT_UNQUALIFIED, 0))
}

/* }}} */

func ZendCompileConstExprMagicConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr

	/* Other cases already resolved by constant folding */

	ZEND_ASSERT(ast.GetAttr() == T_CLASS_C)
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreate(ZEND_AST_CONSTANT_CLASS)
}

/* }}} */

func ZendCompileConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	if ast == nil || ast.GetKind() == ZEND_AST_ZVAL {
		return
	}
	if ZendIsAllowedInConstExpr(ast.GetKind()) == 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Constant expression contains invalid operations")
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
		ZVAL_COPY_VALUE(result, ZendAstGetZval(ast))
	} else {
		ZVAL_AST(result, ZendAstCopy(ast))

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
		CompilerGlobals.SetZendLineno(ast.GetLineno())
		ZendCompileFuncDecl(nil, ast, 1)
		CompilerGlobals.SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
	} else if ast.GetKind() == ZEND_AST_CLASS {
		CompilerGlobals.SetZendLineno(ast.GetLineno())
		ZendCompileClassDecl(ast, 1)
		CompilerGlobals.SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
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
	CompilerGlobals.SetZendLineno(ast.GetLineno())
	if (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_EXTENDED_STMT) != 0 && ZendIsUntickedStmt(ast) == 0 {
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
	if FC(declarables).ticks && ZendIsUntickedStmt(ast) == 0 {
		ZendEmitTick()
	}
}

/* }}} */

func ZendCompileExpr(result *Znode, ast *ZendAst) {
	/* CG(zend_lineno) = ast->lineno; */

	CompilerGlobals.SetZendLineno(ZendAstGetLineno(ast))
	if CompilerGlobals.GetMemoizeMode() != ZEND_MEMOIZE_NONE {
		ZendCompileMemoizedExpr(result, ast)
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_ZVAL:
		ZVAL_COPY(&result.u.constant, ZendAstGetZval(ast))
		result.SetOpType(IS_CONST)
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
		ZendCompileVar(result, ast, BP_VAR_R, 0)
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
		ZEND_ASSERT(false)
	}
}

/* }}} */

func ZendCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *ZendOp {
	CompilerGlobals.SetZendLineno(ZendAstGetLineno(ast))
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
		if type_ == BP_VAR_W || type_ == BP_VAR_RW || type_ == BP_VAR_UNSET {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use temporary expression in write context")
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
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
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
			ZVAL_BOOL(&result, ast.GetKind() == ZEND_AST_OR)
			break
		}
		if ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		child1_is_true = ZendIsTrue(ZendAstGetZval(ast.GetChild()[1]))
		if ast.GetKind() == ZEND_AST_OR {
			ZVAL_BOOL(&result, child0_is_true != 0 || child1_is_true != 0)
		} else {
			ZVAL_BOOL(&result, child0_is_true != 0 && child1_is_true != 0)
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
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | ZEND_DIM_IS)
		}
		ZendEvalConstExpr(&ast.child[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			ZendEvalConstExpr(&ast.child[1])
			return
		}
		if Z_TYPE_P(ZendAstGetZval(ast.GetChild()[0])) == IS_NULL {
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
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if (ast.GetAttr() & ZEND_DIM_ALTERNATIVE_SYNTAX) != 0 {
			ast.SetAttr(ast.GetAttr() &^ ZEND_DIM_ALTERNATIVE_SYNTAX)
			ZendError(E_DEPRECATED, "Array and string offset access syntax with curly braces is deprecated")
		}

		/* Set isset fetch indicator here, opcache disallows runtime altering of the AST */

		if (ast.GetAttr()&ZEND_DIM_IS) != 0 && ast.GetChild()[0].GetKind() == ZEND_AST_DIM {
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | ZEND_DIM_IS)
		}
		ZendEvalConstExpr(&ast.child[0])
		ZendEvalConstExpr(&ast.child[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		container = ZendAstGetZval(ast.GetChild()[0])
		dim = ZendAstGetZval(ast.GetChild()[1])
		if Z_TYPE_P(container) == IS_ARRAY {
			var el *Zval
			if Z_TYPE_P(dim) == IS_LONG {
				el = ZendHashIndexFind(Z_ARR_P(container), Z_LVAL_P(dim))
				if el != nil {
					ZVAL_COPY(&result, el)
				} else {
					return
				}
			} else if Z_TYPE_P(dim) == IS_STRING {
				el = ZendSymtableFind(Z_ARR_P(container), Z_STR_P(dim))
				if el != nil {
					ZVAL_COPY(&result, el)
				} else {
					return
				}
			} else {
				return
			}
		} else if Z_TYPE_P(container) == IS_STRING {
			var offset ZendLong
			var c ZendUchar
			if Z_TYPE_P(dim) == IS_LONG {
				offset = Z_LVAL_P(dim)
			} else if Z_TYPE_P(dim) != IS_STRING || IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), &offset, nil, 1) != IS_LONG {
				return
			}
			if offset < 0 || int(offset >= Z_STRLEN_P(container)) != 0 {
				return
			}
			c = ZendUchar(Z_STRVAL_P(container)[offset])
			ZVAL_INTERNED_STR(&result, ZSTR_CHAR(c))
		} else if Z_TYPE_P(container) <= IS_FALSE {
			ZVAL_NULL(&result)
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

/* }}} */

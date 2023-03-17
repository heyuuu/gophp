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

/* On 64-bit systems less optimal, but more compact VM code leads to better
 * performance. So on 32-bit systems we use absolute addresses for jump
 * targets and constants, but on 64-bit systems realtive 32-bit offsets */

const ZEND_USE_ABS_JMP_ADDR = 0
const ZEND_USE_ABS_CONST_ADDR = 0

/* Temporarily defined here, to avoid header ordering issues */

/* Compilation context that is different for each file, but shared between op arrays. */

type UserOpcodeHandlerT func(executeData *ZendExecuteData) int

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

type ZifHandler func(executeData *ZendExecuteData, return_value *Zval)

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

// math.ceil(sizeof(*ZendExecuteData)/sizeof(Zval))
const ZEND_CALL_FRAME_SLOT = int((ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_execute_data")) + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")))

/* run-time jump target */

/* convert jump target from compile-time to run-time */

/* convert jump target back from run-time to compile-time */

/* constant-time constant */

/* At run-time, constants are allocated together with op_array->opcodes
 * and addressed relatively to current opline.
 */

/* convert constant from compile-time to run-time */

/* convert constant back from run-time to compile-time */

const IS_UNUSED = 0
const IS_CONST ZendUchar = 1 << 0
const IS_TMP_VAR ZendUchar = 1 << 1
const IS_VAR ZendUchar = 1 << 2
const IS_CV ZendUchar = 1 << 3
const ZEND_EXTRA_VALUE = 1

var ZendCompileFile func(file_handle *ZendFileHandle, type_ int) *ZendOpArray
var ZendCompileString func(source_string *Zval, filename *byte) *ZendOpArray

type UnaryOpType func(*Zval, *Zval) int
type BinaryOpType func(*Zval, *Zval, *Zval) int

/* Used during AST construction */

/* parser-driven code generators */

var ZendDoExtendedInfo func()

const INITIAL_OP_ARRAY_SIZE = 64

/* helper functions in zend_language_scanner.l */

const ZEND_FUNCTION_DTOR DtorFuncT = ZendFunctionDtor
const ZEND_CLASS_DTOR DtorFuncT = DestroyZendClass

type ZendNeedsLiveRangeCb func(op_array *ZendOpArray, opline *ZendOp) ZendBool
type ZendAutoGlobalCallback func(name *ZendString) ZendBool

/* BEGIN: OPCODES */

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

/* Quick API to check first 12 arguments */

const MAX_ARG_FLAG_NUM = 12
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

var CompilerGlobals ZendCompilerGlobals
var ExecutorGlobals ZendExecutorGlobals
var ReservedClassNames []ReservedClassName = []ReservedClassName{
	MakeReservedClassName(ZEND_STRL("bool")),
	MakeReservedClassName(ZEND_STRL("false")),
	MakeReservedClassName(ZEND_STRL("float")),
	MakeReservedClassName(ZEND_STRL("int")),
	MakeReservedClassName(ZEND_STRL("null")),
	MakeReservedClassName(ZEND_STRL("parent")),
	MakeReservedClassName(ZEND_STRL("self")),
	MakeReservedClassName(ZEND_STRL("static")),
	MakeReservedClassName(ZEND_STRL("string")),
	MakeReservedClassName(ZEND_STRL("true")),
	MakeReservedClassName(ZEND_STRL("void")),
	MakeReservedClassName(ZEND_STRL("iterable")),
	MakeReservedClassName(ZEND_STRL("object")),
	MakeReservedClassName(nil, 0),
}

var BuiltinTypes []BuiltinTypeInfo = []BuiltinTypeInfo{
	MakeBuiltinTypeInfo(ZEND_STRL("int"), IS_LONG),
	MakeBuiltinTypeInfo(ZEND_STRL("float"), IS_DOUBLE),
	MakeBuiltinTypeInfo(ZEND_STRL("string"), IS_STRING),
	MakeBuiltinTypeInfo(ZEND_STRL("bool"), _IS_BOOL),
	MakeBuiltinTypeInfo(ZEND_STRL("void"), IS_VOID),
	MakeBuiltinTypeInfo(ZEND_STRL("iterable"), IS_ITERABLE),
	MakeBuiltinTypeInfo(ZEND_STRL("object"), IS_OBJECT),
	MakeBuiltinTypeInfo(nil, 0, IS_UNDEF),
}

/* Common part of zend_add_literal and zend_append_individual_literal */

const ZEND_MEMOIZE_NONE = 0
const ZEND_MEMOIZE_COMPILE = 1
const ZEND_MEMOIZE_FETCH = 2

/* Propagate refs used on leaf elements to the surrounding list() structures. */

/* }}}*/

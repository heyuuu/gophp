package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

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
const (

	/* Class, property and method flags                  class|meth.|prop.|const*/
	// Common flags
	// ============
	/* Visibility flags (public < protected < private)        |     |     |     */
	AccPublic    = 1 << 0 /*                                  |  X  |  X  |  X  */
	AccProtected = 1 << 1 /*                                  |  X  |  X  |  X  */
	AccPrivate   = 1 << 2 /*                                  |  X  |  X  |  X  */
	/* Property or method overrides private one */
	AccChanged = 1 << 3 /*                                    |  X  |  X  |     */
	/* Staic method or property */
	AccStatic = 1 << 4 /*                                     |  X  |  X  |     */
	/* Final class or method */
	AccFinal = 1 << 5 /*                                   X  |  X  |     |     */
	// Abstract method
	AccAbstract              = 1 << 6 /*                   X  |  X  |     |     */
	AccExplicitAbstractClass = 1 << 6 /*                   X  |     |     |     */
	// Immutable op_array and class_entries
	// (implemented only for lazy loading of op_arrays)
	AccImmutable = 1 << 7 /*                               X  |  X  |     |     */
	// Function has typed arguments / class has typed props
	AccHasTypeHints = 1 << 8 /*                            X  |  X  |     |     */
	// Top-level class or function declaration
	AccTopLevel = 1 << 9 /*                                X  |  X  |     |     */
	// op_array or class is preloaded
	AccPreloaded = 1 << 10 /*                              X  |  X  |     |     */
	// Class Flags (unused: 24...)
	// ===========
	// Special class types
	AccInterface = 1 << 0 /*                               X  |     |     |     */
	AccTrait     = 1 << 1 /*                               X  |     |     |     */
	AccAnonClass = 1 << 2 /*                               X  |     |     |     */
	// Class linked with parent, interfacs and traits
	AccLinked = 1 << 3 /*                                  X  |     |     |     */
	// class is abstarct, since it is set by any
	// abstract method
	AccImplicitAbstractClass = 1 << 4 /*                   X  |     |     |     */
	// Class has magic methods __get/__set/__unset/
	// __isset that use guards
	AccUseGuards = 1 << 11 /*                              X  |     |     |     */
	// Class constants updated
	AccConstantsUpdated = 1 << 12 /*                       X  |     |     |     */
	// Class extends another class
	AccInherited = 1 << 13 /*                              X  |     |     |     */
	// Class implements interface(s)
	AccImplementInterfaces = 1 << 14 /*                    X  |     |     |     */
	// Class uses trait(s)
	AccImplementTraits = 1 << 15 /*                        X  |     |     |     */
	// User class has methods with static variables
	AccHasStaticInMethods = 1 << 16 /*                     X  |     |     |     */
	// Whether all property types are resolved to CEs
	AccPropertyTypesResolved = 1 << 17 /*                  X  |     |     |     */
	// Children must reuse parent get_iterator()
	AccReuseGetIterator = 1 << 18 /*                       X  |     |     |     */
	// Parent class is resolved (CE).
	AccResolvedParent = 1 << 19 /*                         X  |     |     |     */
	// Interfaces are resolved (CEs).
	AccResolvedInterfaces = 1 << 20 /*                     X  |     |     |     */
	// Class has unresolved variance obligations.
	AccUnresolvedVariance = 1 << 21 /*                     X  |     |     |     */
	// Class is linked apart from variance obligations.
	AccNearlyLinked = 1 << 22 /*                           X  |     |     |     */
	// Whether this class was used in its unlinked state.
	AccHasUnlinkedUses = 1 << 23 /*                        X  |     |     |     */
	// Function Flags (unused: 23, 26)
	// ==============
	// deprecation flag
	AccDeprecated = 1 << 11 /*                                |  X  |     |     */
	// Function returning by reference
	AccReturnReference = 1 << 12 /*                           |  X  |     |     */
	// Function has a return type
	AccHasReturnType = 1 << 13 /*                             |  X  |     |     */
	// Function with variable number of arguments
	AccVariadic = 1 << 14 /*                                  |  X  |     |     */
	// op_array has finally blocks (user only)
	AccHasFinallyBlock = 1 << 15 /*                           |  X  |     |     */
	// "main" op_array with
	// ZEND_DECLARE_CLASS_DELAYED opcodes
	AccEarlyBinding = 1 << 16 /*                              |  X  |     |     */
	// method flag (bc only), any method that has this
	// flag can be used statically and non statically.
	AccAllowStatic = 1 << 17 /*                               |  X  |     |     */
	// call through user function trampoline. e.g.
	// __call, __callstatic
	AccCallViaTrampoline = 1 << 18 /*                         |  X  |     |     */
	// disable inline caching
	AccNeverCache = 1 << 19 /*                                |  X  |     |     */
	// Closure related
	AccClosure     = 1 << 20 /*                               |  X  |     |     */
	AccFakeClosure = 1 << 21 /*                               |  X  |     |     */
	// run_time_cache allocated on heap (user only)
	AccHeapRtCache = 1 << 22 /*                               |  X  |     |     */
	// method flag used by Closure::__invoke() (int only)     |     |     |     */
	AccUserArgInfo = 1 << 22 /*                               |  X  |     |     */
	AccGenerator   = 1 << 24 /*                               |  X  |     |     */
	// function was processed by pass two (user only)
	AccDonePassTwo = 1 << 25 /*                               |  X  |     |     */
	// internal function is allocated at arena (int only)
	AccArenaAllocated = 1 << 25 /*                            |  X  |     |     */
	// op_array is a clone of trait method
	AccTraitClone = 1 << 27 /*                                |  X  |     |     */
	// functions is a constructor
	AccCtor = 1 << 28 /*                                      |  X  |     |     */
	// function is a destructor
	AccDtor = 1 << 29 /*                                      |  X  |     |     */
	// closure uses $this
	AccUsesThis = 1 << 30 /*                                  |  X  |     |     */
	// op_array uses strict mode types
	AccStrictTypes = 1 << 31 /*                               |  X  |     |     */

	//
	AccPppMask = AccPublic | AccProtected | AccPrivate

	// call through internal function handler. e.g. Closure::invoke()
	AccCallViaHandler = AccCallViaTrampoline
)

/* zend_internal_function_handler */

type ZifHandler func(executeData *ZendExecuteData, return_value *types.Zval)

const ZEND_CALL_HAS_THIS = types.IS_OBJECT_EX

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
const ZEND_CALL_NESTED_FUNCTION = ZEND_CALL_FUNCTION | ZEND_CALL_NESTED
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
const IS_CONST = 1 << 0
const IS_TMP_VAR = 1 << 1
const IS_VAR = 1 << 2
const IS_CV = 1 << 3
const ZEND_EXTRA_VALUE = 1

type UnaryOpType func(*types.Zval, *types.Zval) int
type BinaryOpType func(*types.Zval, *types.Zval, *types.Zval) int

/* Used during AST construction */

/* parser-driven code generators */

const INITIAL_OP_ARRAY_SIZE = 64

/* helper functions in zend_language_scanner.l */

const ZEND_FUNCTION_DTOR types.DtorFuncT = ZendFunctionDtor

type ZendNeedsLiveRangeCb func(op_array *types.ZendOpArray, opline *ZendOp) types.ZendBool
type ZendAutoGlobalCallback func(name *types.String) types.ZendBool

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

var CompilerGlobals ZendCompilerGlobals
var ExecutorGlobals ZendExecutorGlobals
var ReservedClassNames = []ReservedClassName{
	MakeReservedClassName("bool"),
	MakeReservedClassName("false"),
	MakeReservedClassName("float"),
	MakeReservedClassName("int"),
	MakeReservedClassName("null"),
	MakeReservedClassName("parent"),
	MakeReservedClassName("self"),
	MakeReservedClassName("static"),
	MakeReservedClassName("string"),
	MakeReservedClassName("true"),
	MakeReservedClassName("void"),
	MakeReservedClassName("iterable"),
	MakeReservedClassName("object"),
}

var BuiltinTypes = []BuiltinTypeInfo{
	MakeBuiltinTypeInfo("int", types.IS_LONG),
	MakeBuiltinTypeInfo("float", types.IS_DOUBLE),
	MakeBuiltinTypeInfo("string", types.IS_STRING),
	MakeBuiltinTypeInfo("bool", types.IS_BOOL),
	MakeBuiltinTypeInfo("void", types.IS_VOID),
	MakeBuiltinTypeInfo("iterable", types.IS_ITERABLE),
	MakeBuiltinTypeInfo("object", types.IS_OBJECT),
}

/* Common part of zend_add_literal and zend_append_individual_literal */

const ZEND_MEMOIZE_NONE = 0
const ZEND_MEMOIZE_COMPILE = 1
const ZEND_MEMOIZE_FETCH = 2

/* Propagate refs used on leaf elements to the surrounding list() structures. */

/* }}}*/

package types

import "math"

const MaxLong = math.MaxInt
const MinLong = math.MinInt
const SizeofLong = 8
const MaxLengthOfLong = 20

const MaxArraySize = math.MaxInt32 + 1

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

/* Used to get hash of the properties of the object, as hash of zval's */
type PropPurposeType int

const (
	PropPurposeDebug PropPurposeType = iota
	PropPurposeArrayCast
	PropPurposeSerialize
	PropPurposeVarExport
	PropPurposeJson
	PropPurposeArrayKeyExists
)

// string
const MaxStrLen = math.MaxInt32

// ZendKnownString
var (
	STR_FILE                 = "file"
	STR_LINE                 = "line"
	STR_FUNCTION             = "function"
	STR_CLASS                = "class"
	STR_OBJECT               = "object"
	STR_TYPE                 = "type"
	STR_OBJECT_OPERATOR      = "->"
	STR_PAAMAYIM_NEKUDOTAYIM = "::"
	STR_ARGS                 = "args"
	STR_UNKNOWN              = "unknown"
	STR_EVAL                 = "eval"
	STR_INCLUDE              = "include"
	STR_REQUIRE              = "require"
	STR_INCLUDE_ONCE         = "include_once"
	STR_REQUIRE_ONCE         = "require_once"
	STR_SCALAR               = "scalar"
	STR_ERROR_REPORTING      = "error_reporting"
	STR_THIS                 = "this"
	STR_VALUE                = "value"
	STR_KEY                  = "key"
	STR_MAGIC_AUTOLOAD       = "__autoload"
	STR_MAGIC_INVOKE         = "__invoke"
	STR_PREVIOUS             = "previous"
	STR_CODE                 = "code"
	STR_MESSAGE              = "message"
	STR_SEVERITY             = "severity"
	STR_STRING               = "string"
	STR_TRACE                = "trace"
	STR_ARGV                 = "argv"
	STR_ARGC                 = "argc"
)

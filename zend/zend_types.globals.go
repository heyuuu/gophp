// <<generate>>

package zend

type ZendBool = uint8
type ZendUchar = uint8
type ZEND_RESULT_CODE = int

const (
	SUCCESS                  = 0
	FAILURE ZEND_RESULT_CODE = -1
)
const ZEND_SIZE_MAX = SIZE_MAX

type ZendIntptrT = intPtr
type ZendUintptrT = uintPtr

type CompareFuncT func(any, any) int
type SwapFuncT func(any, any)
type SortFuncT func(any, int, int, CompareFuncT, SwapFuncT)
type DtorFuncT func(pDest *Zval)
type CopyCtorFuncT func(pElement *Zval)
type ZendType = uintPtr
type HashTable = ZendArray

const HT_INVALID_IDX uint32 = uint32_t - 1
const HT_MIN_MASK uint32 = uint32_t - 2
const HT_MIN_SIZE = 8
const HT_MAX_SIZE = 0x80000000

type HashPosition = uint32

const IS_UNDEF = 0
const IS_NULL = 1
const IS_FALSE = 2
const IS_TRUE = 3
const IS_LONG = 4
const IS_DOUBLE = 5
const IS_STRING = 6
const IS_ARRAY = 7
const IS_OBJECT = 8
const IS_RESOURCE = 9
const IS_REFERENCE = 10
const IS_CONSTANT_AST = 11
const IS_INDIRECT = 13
const IS_PTR = 14
const IS_ALIAS_PTR = 15
const _IS_ERROR = 15
const _IS_BOOL = 16
const IS_CALLABLE = 17
const IS_ITERABLE = 18
const IS_VOID = 19
const _IS_NUMBER = 20
const Z_TYPE_MASK = 0xff
const Z_TYPE_FLAGS_MASK = 0xff00
const Z_TYPE_FLAGS_SHIFT = 8
const GC_TYPE_MASK = 0xf
const GC_FLAGS_MASK = 0x3f0
const GC_INFO_MASK = 0xfffffc00
const GC_FLAGS_SHIFT = 0
const GC_INFO_SHIFT = 10
const GC_COLLECTABLE = 1 << 4
const GC_PROTECTED = 1 << 5
const GC_IMMUTABLE = 1 << 6
const GC_PERSISTENT = 1 << 7
const GC_PERSISTENT_LOCAL = 1 << 8
const GC_ARRAY = IS_ARRAY | GC_COLLECTABLE<<GC_FLAGS_SHIFT
const GC_OBJECT = IS_OBJECT | GC_COLLECTABLE<<GC_FLAGS_SHIFT
const IS_TYPE_REFCOUNTED = 1 << 0
const IS_TYPE_COLLECTABLE = 1 << 1
const IS_INTERNED_STRING_EX uint32 = IS_STRING
const IS_STRING_EX uint32 = IS_STRING | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT
const IS_ARRAY_EX uint32 = IS_ARRAY | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT | IS_TYPE_COLLECTABLE<<Z_TYPE_FLAGS_SHIFT
const IS_OBJECT_EX uint32 = IS_OBJECT | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT | IS_TYPE_COLLECTABLE<<Z_TYPE_FLAGS_SHIFT
const IS_RESOURCE_EX uint32 = IS_RESOURCE | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT
const IS_REFERENCE_EX uint32 = IS_REFERENCE | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT
const IS_CONSTANT_AST_EX uint32 = IS_CONSTANT_AST | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT
const IS_STR_INTERNED = GC_IMMUTABLE
const IS_STR_PERSISTENT = GC_PERSISTENT
const IS_STR_PERMANENT uint32 = 1 << 8
const IS_STR_VALID_UTF8 = 1 << 9
const IS_ARRAY_IMMUTABLE = GC_IMMUTABLE
const IS_ARRAY_PERSISTENT = GC_PERSISTENT
const IS_OBJ_WEAKLY_REFERENCED = GC_PERSISTENT
const IS_OBJ_DESTRUCTOR_CALLED = 1 << 8
const IS_OBJ_FREE_CALLED = 1 << 9
const ZEND_RC_DEBUG = 0
const IS_PROP_UNINIT = 1

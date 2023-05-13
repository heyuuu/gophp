package types

import "math"

type ZendBool = int
type ZEND_RESULT_CODE = int

func IntBool(value bool) ZendBool {
	if value {
		return 1
	} else {
		return 0
	}
}

func ResultCode(value bool) ZEND_RESULT_CODE {
	if value {
		return SUCCESS
	} else {
		return FAILURE
	}
}

const (
	SUCCESS                  = 0
	FAILURE ZEND_RESULT_CODE = -1
)
const ZEND_SIZE_MAX = math.MaxUint

type ZendIntptrT = uintptr
type ZendUintptrT = uintptr

type ZvalComparer func(*Zval, *Zval) int

type ZvalType = uint8

const (
	/* regular data types */
	IS_UNDEF     ZvalType = 0
	IS_NULL      ZvalType = 1
	IS_FALSE     ZvalType = 2
	IS_TRUE      ZvalType = 3
	IS_LONG      ZvalType = 4
	IS_DOUBLE    ZvalType = 5
	IS_STRING    ZvalType = 6
	IS_ARRAY     ZvalType = 7
	IS_OBJECT    ZvalType = 8
	IS_RESOURCE  ZvalType = 9
	IS_REFERENCE ZvalType = 10

	/* constant expressions */

	IS_CONSTANT_AST ZvalType = 11

	/* internal types */

	IS_INDIRECT  ZvalType = 13
	IS_PTR       ZvalType = 14
	IS_ALIAS_PTR ZvalType = 15
	IS_ERROR     ZvalType = 15 // _IS_ERROR

	/* fake types used only for type hinting (Z_TYPE(zv) can not use them) */

	IS_BOOL     ZvalType = 16 // _IS_BOOL
	IS_CALLABLE ZvalType = 17
	IS_ITERABLE ZvalType = 18
	IS_VOID     ZvalType = 19
	IS_NUMBER   ZvalType = 20 // _IS_NUMBER

	_IS_IMMUTABLE_ARRAY ZvalType = 30
)

/* we should never set just Z_TYPE, we should set Z_TYPE_INFO */

const Z_TYPE_MASK = 0xff
const Z_TYPE_FLAGS_MASK = 0xff00
const Z_TYPE_FLAGS_SHIFT = 8

/* zval_gc_flags(zval.value->gc.u.type_info) (common flags) */

const GC_COLLECTABLE = 1 << 4
const GC_PROTECTED = 1 << 5
const GC_IMMUTABLE = 1 << 6

/* zval.u1.v.type_flags */

const IS_TYPE_REFCOUNTED = 1 << 0
const IS_TYPE_COLLECTABLE = 1 << 1

/* This optimized version assumes that we have a single "type_flag" */

/* extended types */

const IS_OBJECT_EX uint32 = IS_OBJECT | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT | IS_TYPE_COLLECTABLE<<Z_TYPE_FLAGS_SHIFT
const IS_REFERENCE_EX uint32 = IS_REFERENCE | IS_TYPE_REFCOUNTED<<Z_TYPE_FLAGS_SHIFT

/* object flags (zval.value->gc.u.flags) */

//const IS_OBJ_WEAKLY_REFERENCED = GC_PERSISTENT
const IS_OBJ_DESTRUCTOR_CALLED = 1 << 8
const IS_OBJ_FREE_CALLED = 1 << 9

const IS_PROP_UNINIT = 1

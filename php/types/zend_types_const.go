package types

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

/* object flags (zval.value->gc.u.flags) */

const IS_PROP_UNINIT = 1

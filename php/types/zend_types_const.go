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

/* we should never set just Z_TYPE, we should set Z_TYPE_INFO */

const Z_TYPE_MASK = 0xff
const Z_TYPE_FLAGS_MASK = 0xff00
const Z_TYPE_FLAGS_SHIFT = 8

/* object flags (zval.value->gc.u.flags) */

const IS_PROP_UNINIT = 1

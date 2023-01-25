// <<generate>>

package zend

import (
	"sik/core"
)

func ZEND_LONG_INT_OVFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong > ZendLong(core.INT_MAX))
}
func ZEND_LONG_INT_UDFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong < ZendLong(core.INT_MIN))
}
func ZEND_LONG_EXCEEDS_INT(zlong __auto__) __auto__ {
	return UNEXPECTED(ZEND_LONG_INT_OVFL(zlong) || ZEND_LONG_INT_UDFL(zlong))
}
func ZEND_LONG_UINT_OVFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong < 0 || zlong > ZendLong(UINT_MAX))
}
func ZEND_SIZE_T_INT_OVFL(size int) __auto__ {
	return UNEXPECTED(size > int(core.INT_MAX))
}
func ZEND_SIZE_T_UINT_OVFL(size __auto__) __auto__ { return UNEXPECTED(size > int(UINT_MAX)) }
func ZEND_SIZE_T_GT_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong < 0 || size > size_t(zlong)
}
func ZEND_SIZE_T_GTE_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong < 0 || size >= size_t(zlong)
}
func ZEND_SIZE_T_LT_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong >= 0 && size < size_t(zlong)
}
func ZEND_SIZE_T_LTE_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong >= 0 && size <= size_t(zlong)
}

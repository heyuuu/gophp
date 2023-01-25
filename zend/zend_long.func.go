// <<generate>>

package zend

import (
	"sik/core"
)

func Z_L(i int) __auto__  { return int64(i) }
func Z_UL(i int) __auto__ { return uint64(i) }
func ZEND_LTOA(i __auto__, s []char, len_ __auto__) {
	var st int = core.Snprintf(s, len_, ZEND_LONG_FMT, i)
	s[st] = '0'
}
func ZEND_ATOL(i __auto__, s __auto__) __auto__ {
	i = atoll(s)
	return i
}
func ZEND_STRTOL(s0 __auto__, s1 **byte, base int) __auto__  { return strtoll(s0, s1, base) }
func ZEND_STRTOUL(s0 __auto__, s1 **byte, base int) __auto__ { return strtoull(s0, s1, base) }

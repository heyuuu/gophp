// <<generate>>

package zend

func ZEND_ATOL(i __auto__, s __auto__) __auto__ {
	i = atoll(s)
	return i
}
func ZEND_STRTOL(s0 __auto__, s1 **byte, base int) __auto__  { return strtoll(s0, s1, base) }
func ZEND_STRTOUL(s0 __auto__, s1 **byte, base int) __auto__ { return strtoull(s0, s1, base) }

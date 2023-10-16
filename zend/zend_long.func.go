package zend

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"strconv"
)

func ZEND_ATOL(i __auto__, s __auto__) __auto__ {
	i = atoll(s)
	return i
}
func ZEND_STRTOL(s0 __auto__, s1 **byte, base int) __auto__  { return strtoll(s0, s1, base) }
func ZEND_STRTOUL(s0 __auto__, s1 **byte, base int) __auto__ { return strtoull(s0, s1, base) }

func TryStrToLong(str string) (n int, useLen int, ok bool) {
	if len(str) == 0 || !ascii.IsDigit(str[0]) {
		return 0, 0, false
	}

	useLen = 1
	for useLen < len(str) && ascii.IsDigit(str[useLen]) {
		useLen++
	}

	if val, err := strconv.Atoi(str[:useLen]); err == nil {
		return val, useLen, true
	} else {
		return 0, 0, false
	}
}

func StrToLong(s string, base int) (int, error) {
	i := 0

	// skip spaces
	for i < len(s) && s[i] == ' ' {
		i++
	}
	start := i

	// scan digits
	for i < len(s) && '0' <= s[i] && s[i] <= '9' {
		i++
	}

	// parse
	val, err := strconv.ParseInt(s[start:i], base, 64)
	if err != nil {
		return 0, err
	}
	return int(val), err
}

func StrToLongWithUnit(str string) ZendLong {
	if len(str) == 0 {
		return 0
	}
	retval, _ := StrToLong(str, 0)
	switch str[len(str)-1] {
	case 'g', 'G':
		retval *= 1024
		fallthrough
	case 'm', 'M':
		retval *= 1024
		fallthrough
	case 'k', 'K':
		retval *= 1024
	}

	return retval
}

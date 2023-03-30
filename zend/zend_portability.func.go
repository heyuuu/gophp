package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
)

func ZEND_IGNORE_VALUE(x __auto__) { void(x) }
func ZendQuietWrite()              { ZEND_IGNORE_VALUE(write(__VA_ARGS__)) }
func DL_LOAD(libname *byte) __auto__ {
	return dlopen(libname, PHP_RTLD_MODE|RTLD_GLOBAL)
}
func DoAlloca(p int, use_heap __auto__) any { return Emalloc(p) }
func FreeAlloca(p any, use_heap __auto__)   { Efree(p) }
func ZEND_BIT_TEST(bits []uint32, bit uint32) int {
	return bits[bit/(b.SizeOf("( bits ) [ 0 ]")*8)] >> (bit&b.SizeOf("( bits ) [ 0 ]")*8 - 1) & 1
}
func ZEND_STRL(str string) string { return str }
func ZEND_NORMALIZE_BOOL(n ZendLong) int {
	if n != 0 {
		if n < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}
func ZEND_SECURE_ZERO(var_ __auto__, size __auto__) __auto__ { return core.ExplicitBzero(var_, size) }
func ZEND_VALID_SOCKET(sock core.PhpSocketT) bool            { return sock >= 0 }
func VaCopy(dest ...any, src ...any) __auto__ {
	return memcpy(&dest, &src, b.SizeOf("va_list"))
}
func ZEND_EXPAND_VA(code __auto__) __auto__ { return code }

// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func __hasAttribute(x __auto__) int { return 0 }
func __hasBuiltin(x __auto__) int   { return 0 }
func __hasFeature(x __auto__) int   { return 0 }
func ZEND_ASSERT(c bool)            {}
func ZEND_IGNORE_VALUE(x __auto__)  { void(x) }
func ZendQuietWrite()               { ZEND_IGNORE_VALUE(write(__VA_ARGS__)) }
func DL_LOAD(libname *byte) __auto__ {
	return dlopen(libname, PHP_RTLD_MODE|RTLD_GLOBAL)
}
func ZEND_CONST_COND(_condition __auto__, _default __auto__) __auto__ { return _default }
func EXPECTED(condition bool) __auto__                                { return __builtin_expect(!!condition, 1) }
func UNEXPECTED(condition bool) __auto__                              { return __builtin_expect(!!condition, 0) }
func DoAlloca(p int, use_heap __auto__) any                           { return Emalloc(p) }
func FreeAlloca(p any, use_heap __auto__)                             { Efree(p) }
func SETJMP(a JMP_BUF) __auto__                                       { return sigsetjmp(a, 0) }
func LONGJMP(a JMP_BUF, b ZEND_RESULT_CODE) __auto__                  { return siglongjmp(a, b) }
func MAX(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func MIN(a __auto__, b __auto__) __auto__ {
	if a < b {
		return a
	} else {
		return b
	}
}
func ZEND_BIT_TEST(bits []uint32, bit uint32) int {
	return bits[bit/(b.SizeOf("( bits ) [ 0 ]")*8)] >> (bit&b.SizeOf("( bits ) [ 0 ]")*8 - 1) & 1
}
func _zendGetInf() float64 { return HUGE_VAL }
func _zendGetNan() float64 { return 0.0 / 0.0 }
func ZEND_STRL(str string) {
	str
	b.SizeOf("str") - 1
}
func ZEND_STRS(str __auto__) {
	str
	b.SizeOf("str")
}
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
func ZEND_TRUTH(x __auto__) int {
	if x {
		return 1
	} else {
		return 0
	}
}
func ZEND_LOG_XOR(a __auto__, b __auto__) int                { return ZEND_TRUTH(a) ^ ZEND_TRUTH(b) }
func ZEND_SECURE_ZERO(var_ __auto__, size __auto__) __auto__ { return core.ExplicitBzero(var_, size) }
func ZEND_VALID_SOCKET(sock core.PhpSocketT) bool            { return sock >= 0 }
func VaCopy(dest ...any, src ...any) __auto__ {
	return memcpy(&dest, &src, b.SizeOf("va_list"))
}
func ZEND_SLIDE_TO_ALIGNED(alignment int, ptr __auto__) int {
	return zend_uintptr_t(ptr) + (alignment-1) & ^(alignment-1)
}
func ZEND_SLIDE_TO_ALIGNED16(ptr __auto__) int {
	return ZEND_SLIDE_TO_ALIGNED(uint64(16), ptr)
}
func ZEND_EXPAND_VA(code __auto__) __auto__ { return code }

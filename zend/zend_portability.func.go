package zend

import (
	"github.com/heyuuu/gophp/core"
)

func ZEND_IGNORE_VALUE(x __auto__)      { void(x) }
func ZendQuietWrite()                   { ZEND_IGNORE_VALUE(write(__VA_ARGS__)) }
func DoAlloca(p int, use_heap bool) any { return Emalloc(p) }
func FreeAlloca(p any, use_heap bool)   { Efree(p) }
func ZEND_NORMALIZE_BOOL[T int | float64](n T) int {
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
func ZEND_SECURE_ZERO(dst *byte, siz int) {
	var i int = 0
	var buf *uint8 = (*uint8)(dst)
	for ; i < siz; i++ {
		buf[i] = 0
	}
}
func ZEND_VALID_SOCKET(sock core.PhpSocketT) bool { return sock >= 0 }

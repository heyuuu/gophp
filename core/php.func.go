package core

import (
	"github.com/heyuuu/gophp/zend"
)

func PHP_STRLCPY(dst []char, src __auto__, size int, src_size int) {
	var php_str_len int
	if src_size >= size {
		php_str_len = size - 1
	} else {
		php_str_len = src_size
	}
	memcpy(dst, src, php_str_len)
	dst[php_str_len] = '0'
}
func STR_PRINT(str *byte) string {
	if str != nil {
		return str
	} else {
		return ""
	}
}
func PhpIgnoreValue(x __auto__) { zend.ZEND_IGNORE_VALUE(x) }

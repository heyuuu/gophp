package core

import (
	"github.com/heyuuu/gophp/zend"
)

func Spprintf(message *string, max_len int, format string, args ...any) int {
	result := zend.ZendSprintfEx(max_len, format, args...)
	*message = result
	return len(result)
}

func Vspprintf(pbuf *string, max_len int, format string, args ...any) int {
	/* since there are places where (v)spprintf called without checking for null,
	   a bit of defensive coding here */
	if pbuf == nil {
		return 0
	}
	return Spprintf(pbuf, max_len, format, args...)
}

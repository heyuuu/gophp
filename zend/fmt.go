package zend

import (
	"github.com/heyuuu/gophp/core/pfmt"
)

// 替代各种 sprintf 方法(限制长度)
func ZendSprintfEx(maxLen int, format string, args ...any) string {
	if maxLen != 0 {
		return pfmt.Snprintf(maxLen, format, args...)
	} else {
		return pfmt.Sprintf(format, args...)
	}
}

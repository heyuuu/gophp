package zend

import (
	"github.com/heyuuu/gophp/core/pfmt"
	"github.com/heyuuu/gophp/zend/types"
)

// 替代各种 sprintf 方法(限制长度)
func ZendSprintfEx(maxLen int, format string, args ...any) string {
	if maxLen != 0 {
		return pfmt.Snprintf(maxLen, format, args)
	} else {
		return pfmt.Sprintf(format, args)
	}
}

// 替代各种 sprintf 方法
func ZendSprintf(format string, args ...any) string {
	return pfmt.Sprintf(format, args)
}

// 替代各种 sprintf 方法
func ZendSprintfZStr(format string, args ...any) *types.String {
	result := pfmt.Sprintf(format, args)
	return types.NewString(result)
}

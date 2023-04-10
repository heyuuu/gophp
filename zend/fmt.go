package zend

import "github.com/heyuuu/gophp/zend/types"

// 替代各种 sprintf 方法(限制长度)
func ZendSprintfEx(maxLen int, format string, args ...any) string {
	result := ZendSprintf(format, args...)
	if maxLen != 0 && len(result) > maxLen {
		return result[:maxLen]
	}
	return result
}

// 替代各种 sprintf 方法
func ZendSprintf(format string, args ...any) string {
	var buf = SmartStr{}
	ZendPrintfToSmartStr(&buf, format, args...)
	return buf.GetStr()
}

// 替代各种 sprintf 方法
func ZendSprintfZStr(format string, args ...any) *types.String {
	return types.NewString(ZendSprintf(format, args...))
}

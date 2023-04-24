package zend

import (
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"strconv"
)

func StrToDouble(str string) float64 {
	// todo 注意，此处异常有差异未处理
	// - 越界时，返回 +INF 或 -INF；此时逻辑一致
	// - 语法错误，但前序字符串合法时，PHP 返回前序字符串对应值，Golang 返回 0。此处不一致待修改
	d, _ := strconv.ParseFloat(str, 64)
	return d
}

/**
 * constants
 */
type ConvertNumericMode int

const (
	ConvertRefuseErrors     ConvertNumericMode = 0  // 不允许错误
	ConvertContinueOnErrors ConvertNumericMode = 1  // 允许不完全匹配
	ConvertNoticeOnErrors   ConvertNumericMode = -1 // 允许不完全匹配，不完全匹配时触发 Zend Notice (可能产生 ZendException)
)

func StrToNumber(str string) conv.ParseNumberResult {
	return StrToNumberEx(str, ConvertRefuseErrors)
}
func StrToNumberAllowErrors(str string) conv.ParseNumberResult {
	return StrToNumberEx(str, ConvertContinueOnErrors)
}
func StrToNumberNoticeErrors(str string) conv.ParseNumberResult {
	return StrToNumberEx(str, ConvertNoticeOnErrors)
}

/**
 * StrToNumberEx 	尝试转换字符串为数字
 * @param	str		待转换的字符串
 * @param	mode 	是否允许错误，具体参看上方常量
 * @return 	conv.ParseNumberResult
 */
func StrToNumberEx(str string, mode ConvertNumericMode) conv.ParseNumberResult {
	switch mode {
	case ConvertRefuseErrors:
		return conv.ParseNumber(str)
	case ConvertNoticeOnErrors:
		result, matchLen := conv.ParseNumberPrefix(str, false)
		if matchLen != len(str) {
			// notice: 此处可能会触发 Exception
			faults.Error(faults.E_NOTICE, "A non well formed numeric value encountered")
		}
		return result
	default:
		fallthrough
	case ConvertContinueOnErrors:
		result, _ := conv.ParseNumberPrefix(str, false)
		return result
	}
}

func StrToNumberZvalEx(str string, mode ConvertNumericMode) *types.Zval {
	r := StrToNumberEx(str, mode)
	if r.IsInt() {
		return types.NewZvalLong(r.Int())
	} else if r.IsFloat() {
		return types.NewZvalDouble(r.Float())
	} else {
		return nil
	}
}

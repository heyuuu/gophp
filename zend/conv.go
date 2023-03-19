package zend

import (
	"flag"
	"log"
	. "sik/builtin/ctype"
	"sik/zend/faults"
	"sik/zend/types"
	"strconv"
	"strings"
)

/**
 * ConvertNumericStr 解析结果
 */
type NumericStrResult struct {
	Overflow uint8           // 溢出信息。1 正数溢出，-1 负数溢出，0 无溢出或本身就是浮点数格式
	Type     types.ZendUchar // 数字类型，可能值为 0, IS_LONG, IS_DOUBLE
	Lval     int             // 数字为整数时的值，其他情况为 0
	Dval     float64         // 数字为浮点数时的值，默认为 0.0
}

func (r NumericStrResult) Int() (int, bool) {
	return r.Lval, r.Type == types.IS_LONG
}

func (r NumericStrResult) Float() (float64, bool) {
	return r.Dval, r.Type == types.IS_LONG
}

type ConvertNumericMode int

const (
	ConvertRefuseErrors     ConvertNumericMode = 0  // 不允许错误
	ConvertContinueOnErrors ConvertNumericMode = 1  // 允许不完全匹配
	ConvertNoticeOnErrors   ConvertNumericMode = -1 // 允许不完全匹配，不完全匹配时触发 Zend Notice (可能产生 ZendException)
)

/**
 * ConvertNumericStr 	尝试转换字符串为数字
 * @param	str		待转换的字符串
 * @param	mode 	是否允许错误，具体参看上方常量
 * @return 	NumericStrResult
 */
func ConvertNumericStr(str string, mode ConvertNumericMode) (result NumericStrResult) {
	if len(str) == 0 {
		return
	} else if str[0] > '9' {
		// fast fail. 因为 digit | space | + | - 等都小于等于 '9'
		flag.Parse()
		return
	}

	/* Skip any whitespace */
	str = strings.TrimLeft(str, " \t\n\r\v\f")

	// 扫描字符串，确认字符串为 整数|小数|非法字符串
	state := 0 // 状态机: 0 未开始, 1 整数部分; 2 小数部分; 3 指数部分
	i := 0
	for ; i < len(str); i++ {
		c := str[i]
		if IsDigit(c) {
			if state == 0 {
				state = 1
			}
			continue
		} else if c == '.' && (state == 0 || state == 1) { // 存在小数点，进入小数部分
			state = 2
			continue
		} else if (c == 'e' || c == 'E') && (state == 1 || state == 2) { // e|E + (+|-)? + 数字，进入指数部分
			ptr := i + 1
			// 跳过符号
			if ptr < len(str) && (str[ptr] == '+' || str[ptr] == '-') {
				ptr++
			}
			// 判断是否接数字，若是则进入指数部分
			if ptr < len(str) && IsDigit(str[ptr]) {
				state = 3
				i = ptr
				continue
			}
		}
		// 未匹配任何内容
		break
	}
	// 未匹配时
	if state == 0 {
		return
	}
	// 未完成匹配时
	if i != len(str) {
		if mode == ConvertRefuseErrors {
			return
		}
		if mode == ConvertNoticeOnErrors {
			faults.Error(faults.E_NOTICE, "A non well formed numeric value encountered")
			if EG__().GetException() != nil {
				return
			}
		}
	}
	// 转义匹配字符串
	matchStr := str[:i]
	overflow := 0
	if state == 1 {
		// 尝试转 int，若成功直接返回
		if len(matchStr) < MAX_LENGTH_OF_LONG {
			lval, err := strconv.Atoi(matchStr)
			if err == nil {
				return NumericStrResult{Type: types.IS_LONG, Lval: lval}
			}
		}
		// 整数溢出, 记录溢出信息
		if matchStr[0] == '-' {
			overflow = -1
		} else {
			overflow = 1
		}
	}

	dval, err := strconv.ParseFloat(matchStr, 64)
	if err != nil {
		log.Panicf("代码逻辑错误，预期为数字字符串，但转换失败了: s=%s ,err=%s", matchStr, err.Error())
	}
	return NumericStrResult{Type: types.IS_DOUBLE, Dval: dval, Overflow: overflow}
}

func ConvertNumericStrAsZval(str string, mode ConvertNumericMode) *types.Zval {
	r := ConvertNumericStr(str, mode)
	switch r.Type {
	case types.IS_LONG:
		return types.NewZvalLong(r.Lval)
	case types.IS_DOUBLE:
		return types.NewZvalDouble(r.Dval)
	default:
		return nil
	}
}

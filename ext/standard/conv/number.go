package conv

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/zend/faults"
	"log"
	"strconv"
	"strings"
)

/**
 * constants && types
 */

const maxLengthOfLong = 20

type ParseNumberResult struct {
	typ      int8    // 数字类型，可能值为 0: 解析失败, 1: 整数, 2: 浮点数
	overflow int8    // 溢出信息。1 正数溢出，-1 负数溢出，0 无溢出或本身就是浮点数格式
	iVal     int     // 数字为整数时的值，其他情况为 0
	fVal     float64 // 数字为浮点数时的值，默认为 0.0
}

func parseNumberResultInt(i int) ParseNumberResult {
	return ParseNumberResult{typ: 1, iVal: i}
}
func parseNumberResultFloat(f float64, overflow int8) ParseNumberResult {
	return ParseNumberResult{typ: 1, fVal: f, overflow: overflow}
}
func (r ParseNumberResult) IsSucc() bool     { return r.typ != 0 }
func (r ParseNumberResult) IsInt() bool      { return r.typ == 1 }
func (r ParseNumberResult) IsFloat() bool    { return r.typ == 2 }
func (r ParseNumberResult) IsOverflow() bool { return r.overflow != 0 }
func (r ParseNumberResult) Int() int         { return r.iVal }
func (r ParseNumberResult) Float() float64   { return r.fVal }
func (r ParseNumberResult) Overflow() int    { return int(r.overflow) }
func (r ParseNumberResult) ToFloat() float64 {
	if r.IsFloat() {
		return r.Float()
	} else {
		return float64(r.Int())
	}
}

/**
 * functions
 */
// 尝试转换字符串整体为数字
func ParseNumber(str string) ParseNumberResult {
	r, _ := ParseNumberPrefix(str, true)
	return r
}

// 尽量尝试转换字符串前缀为数字，返回转换结果+匹配长度 (类似 strconv.parseFloatPrefix())
func ParseNumberPrefix(str string, strict bool) (result ParseNumberResult, matchLen int) {
	if len(str) == 0 {
		return
	} else if str[0] > '9' {
		// fast fail. 因为 digit | space | + | - 等都小于等于 '9'
		return
	}

	/* Skip any whitespace */
	str = strings.TrimLeft(str, " \t\n\r\v\f")

	// 扫描字符串，确认字符串为 整数|小数|非法字符串
	state := 0 // 状态机: 0 未开始, 1 整数部分; 2 小数部分; 3 指数部分
	i := 0
	for ; i < len(str); i++ {
		c := str[i]
		if ascii.IsDigit(c) {
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
			if ptr < len(str) && ascii.IsDigit(str[ptr]) {
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
	matchLen = i

	// 未完成匹配时
	if matchLen != len(str) {
		if strict {
			return
		}
	}

	if i != len(str) {
		if mode == ModeRefuseErrors {
			return
		}
		if mode == ModeNoticeOnErrors {
			faults.Error(faults.E_NOTICE, "A non well formed numeric value encountered")
			if EG__().HasException() {
				return
			}
		}
	}
	// 转义匹配字符串
	matchStr := str[:i]
	overflow := int8(0)
	if state == 1 {
		// 尝试转 int，若成功直接返回
		if len(matchStr) < maxLengthOfLong {
			lval, err := strconv.Atoi(matchStr)
			if err == nil {
				return parseNumberResultInt(lval), matchLen
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
	return parseNumberResultFloat(dval, overflow), matchLen
}

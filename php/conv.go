package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php/types"
	"log"
	"math"
	"strconv"
)

func ParseDouble(str string) float64 {
	matchStr, _, _ := matchNumberPrefix(str)
	if matchStr == "" { // not match
		return 0 // todo 异常处理
	}
	d, _ := strconv.ParseFloat(matchStr, 64)
	return d
}

func ParseNumber(str string) Val {
	zv, _ := ParseNumberEx(str)
	return zv
}
func ParseNumberEx(str string) (Val, int) {
	zv, overflow, matchLen := parseNumberPrefix(str)
	if matchLen != len(str) {
		return types.Undef, 0
	}
	return zv, overflow
}
func ParseNumberPrefix(str string) (Val, int) {
	zv, _, matchLen := parseNumberPrefix(str)
	return zv, matchLen
}

func DoubleToLong(d float64) int {
	if !mathkit.IsFinite(d) {
		return 0
	} else if math.MinInt <= d && d < math.MaxInt {
		return int(d)
	}

	// zend_dval_to_lval_slow
	// 越界 mod 处理 (和 go 强制转换机制不同，无法直接转换)
	dmod := math.Mod(d, 1<<64)
	if dmod > math.MaxInt {
		dmod -= 1 << 64
	} else if dmod < math.MinInt {
		dmod += 1 << 64
	}
	return int(dmod)
}
func DoubleToLongCap(d float64) int {
	if !mathkit.IsFinite(d) {
		return 0
	} else if d < math.MinInt {
		return math.MinInt
	} else if d >= math.MaxInt {
		return math.MaxInt
	} else {
		return int(d)
	}
}

/**
 * private functions
 */
const maxLengthOfLong = 20

// _is_numeric_string_ex
// 尽量尝试转换字符串前缀为数字，返回转换结果+匹配长度 (类似 strconv.parseFloatPrefix())
func parseNumberPrefix(str string) (zv Val, overflow int, matchLen int) {
	matchStr, matchLen, maybeLong := matchNumberPrefix(str)
	if matchStr == "" { // not match
		return
	}

	// 转义匹配字符串
	overflow = 0
	if maybeLong {
		// 尝试转 int，若成功直接返回
		if len(matchStr) < maxLengthOfLong {
			lval, err := strconv.Atoi(matchStr)
			if err == nil {
				return Long(lval), overflow, matchLen
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
	return Double(dval), overflow, matchLen
}

// 尽量尝试匹配字符串中可能为数字的前缀，返回匹配位置及是否可能为整形
func matchNumberPrefix(str string) (matchStr string, matchLen int, maybeLong bool) {
	if len(str) == 0 {
		return
	} else if str[0] > '9' {
		// fast fail. 因为 digit | space | + | - 等都小于等于 '9'
		return
	}

	/* Skip any whitespace */
	start := 0
	for start < len(str) && ascii.IsSpace(str[start]) {
		start++
	}

	// 扫描字符串，确认字符串为 整数|小数|非法字符串
	state := 0 // 状态机: 0 未开始, 1 整数部分; 2 小数部分; 3 指数部分
	idx := start
	if idx < len(str) && (str[idx] == '+' || str[idx] == '-') {
		idx++
	}
	for ; idx < len(str); idx++ {
		c := str[idx]
		if ascii.IsDigit(c) {
			if state == 0 {
				state = 1
			}
			continue
		} else if c == '.' && (state == 0 || state == 1) { // 存在小数点，进入小数部分
			state = 2
			continue
		} else if (c == 'e' || c == 'E') && (state == 1 || state == 2) { // e|E + (+|-)? + 数字，进入指数部分
			ptr := idx + 1
			// 跳过符号
			if ptr < len(str) && (str[ptr] == '+' || str[ptr] == '-') {
				ptr++
			}
			// 判断是否接数字，若是则进入指数部分
			if ptr < len(str) && ascii.IsDigit(str[ptr]) {
				state = 3
				idx = ptr
				continue
			}
		}
		// 未匹配任何内容
		break
	}

	if state == 0 {
		return "", 0, false
	}
	return str[start:idx], idx, state == 1
}

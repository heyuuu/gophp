package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"math"
	"strconv"
	"strings"
)

func StringCompare(s1, s2 string) int {
	// notice: 在 PHP < 8.2.0 的版本，很多字符串比较返回值并非限定为 -1/0/1，故不能直接使用 strings.Compare
	for i := 0; i < len(s1) && i < len(s2); i++ {
		diff := int(s1[i]) - int(s2[i])
		if diff != 0 {
			return diff
		}
	}
	return len(s1) - len(s2)
}

func StringCaseCompare(s1, s2 string) int {
	// notice: 在 PHP < 8.2.0 的版本，很多字符串比较返回值并非限定为 -1/0/1，故不能直接使用 ascii.StrCaseCompare
	for i := 0; i < len(s1) && i < len(s2); i++ {
		diff := int(ascii.ToLower(s1[i])) - int(ascii.ToLower(s2[i]))
		if diff != 0 {
			return diff
		}
	}
	return len(s1) - len(s2)
}

func StringNCaseCompare(s1, s2 string, n int) int {
	if len(s1) > n {
		s1 = s1[:n]
	}
	if len(s2) > n {
		s2 = s2[:n]
	}
	return StringCaseCompare(s1, s2)
}

// @see: FORMAT_CONV_MAX_PRECISION
const formatMaxPrec = 500

// php 对浮点数的格式化方式，兼容一些细微差异
func FormatDouble(f float64, fmtTyp byte, prec int) string {
	if math.IsNaN(f) {
		return "NAN"
	}

	if fmtTyp == 'F' {
		fmtTyp = 'f'
	}
	if prec > formatMaxPrec {
		prec = formatMaxPrec
	}
	str := strconv.FormatFloat(f, fmtTyp, prec, 64)

	// fix: 兼容细节差异
	switch fmtTyp {
	case 'G', 'g':
		// fix: g|G 模式下，科学计数法基数只有正数位时，php 会补充 ".0"（go 会省略小数点和小数部分）
		if idx := strings.LastIndexAny(str, "eE"); idx >= 0 {
			if pointIdx := strings.IndexByte(str[:idx], '.'); pointIdx < 0 {
				str = str[:idx] + ".0" + str[idx:]
			}
		}
		fallthrough
	case 'e', 'E':
		// fix: e|E 模式下，指数不会有前导'0' (go 默认会补齐前导0至最少2位)
		if idx := strings.LastIndexAny(str, "eE"); idx >= 0 {
			if idx+3 < len(str) && str[idx+2] == '0' {
				str = str[:idx+2] + str[idx+3:]
			}
		}
	}

	return str
}

// php 对浮点数的格式化方式，兼容一些细微差异
// todo
func SerializeDouble(f float64, prec int) string {
	str := strconv.FormatFloat(f, 'G', prec, 64)

	if idx := strings.LastIndexAny(str, "eE"); idx >= 0 {
		var buf strings.Builder

		buf.WriteString(str[:idx])
		if pointIdx := strings.IndexByte(str[:idx], '.'); pointIdx < 0 {
			buf.WriteString(".0")
		}

		if idx+3 < len(str) && str[idx+2] == '0' {
			buf.WriteString(str[idx : idx+2])
			buf.WriteString(str[idx+3:])
		} else {
			buf.WriteString(str[idx:])
		}
		str = buf.String()
	}
	return str
}

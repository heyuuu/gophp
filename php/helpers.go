package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
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

// php 对浮点数的格式化方式，兼容一些细微差异
func FormatDouble(f float64, fmtTyp byte, prec int) string {
	if fmtTyp == 'F' {
		fmtTyp = 'f'
	}
	str := strconv.FormatFloat(f, fmtTyp, prec, 64)

	// fix: 科学计数法且位数为指数个位数时，go 默认会补齐到2位，但 php 保持位数不变。此处修复此差别
	if idx := strings.LastIndexAny(str, "eE"); idx >= 0 {
		if idx+3 < len(str) && str[idx+2] == '0' {
			str = str[:idx+2] + str[idx+3:]
		}
	}

	return str
}

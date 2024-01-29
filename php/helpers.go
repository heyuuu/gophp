package php

import "github.com/heyuuu/gophp/kits/ascii"

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

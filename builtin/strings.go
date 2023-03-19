package builtin

import "strings"

func StrCaseCompare(s1 string, s2 string) bool {
	var l1 = strings.ToLower(s1)
	var l2 = strings.ToLower(s2)
	return l1 < l2
}

// 对应 c 函数 strlen
func Strlen(s string) int {
	pos := strings.IndexByte(s, '\000')
	if pos < 0 {
		return len(s)
	} else {
		return pos
	}
}

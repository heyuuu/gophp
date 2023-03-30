package builtin

import (
	"strings"
)

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

func AsciiCaseEquals(s1 string, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func ByteToLowerAscii(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func StringToLowerAscii(s string) string {
	i := 0
	for ; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			break
		}
	}
	if i == len(s) {
		return s
	}

	var buf strings.Builder
	buf.Grow(len(s))
	buf.WriteString(s[:i])
	for j := i; j < len(s); j++ {
		c := s[j]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		buf.WriteByte(c)
	}
	return buf.String()
}

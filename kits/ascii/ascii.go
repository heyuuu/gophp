package ascii

import "strings"

func IsAscii(c byte) bool {
	return c <= 0x7f
}
func IsAsciiRune(r rune) bool {
	return 0 <= r && r <= 0x7f
}

func IsControl(c byte) bool {
	return c <= 0x1f || c == 0x7f
}

func IsAlphaNum(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func IsAlpha(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

func IsUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsLower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func IsXDigit(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func ParseXDigit(c byte) (byte, bool) {
	if c >= '0' && c <= '9' {
		return c - '0', true
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10, true
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10, true
	} else {
		return 0, false
	}
}

func IsSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r'
}

func IsSpaceRune(r rune) bool { return r <= 0xff && IsSpace(byte(r)) }

func ToUpper(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func ToLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

// 字符串转小写。与 strings.ToUpper() 的区别是，它不支持除英文字母外的其他unicode字母
func StrToUpper(s string) string {
	var idx int
	for idx = 0; idx < len(s); idx++ {
		if 'a' <= s[idx] && s[idx] <= 'z' {
			break
		}
	}
	if idx == len(s) {
		return s
	}

	var buf strings.Builder
	buf.WriteString(s[:idx])
	for i := idx; i < len(s); i++ {
		buf.WriteByte(ToUpper(s[i]))
	}
	return buf.String()
}

// 字符串转小写。与 strings.ToLower() 的区别是，它不处理除英文字母外的其他unicode字母
func StrToLower(s string) string {
	var idx int
	for idx = 0; idx < len(s); idx++ {
		if 'A' <= s[idx] && s[idx] <= 'Z' {
			break
		}
	}
	if idx == len(s) {
		return s
	}

	var buf strings.Builder
	buf.WriteString(s[:idx])
	for i := idx; i < len(s); i++ {
		buf.WriteByte(ToLower(s[i]))
	}
	return buf.String()
}

func StrCaseEquals(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if ToLower(a[i]) != ToLower(b[i]) {
			return false
		}
	}
	return true
}

func StrCaseCompare(a string, b string) int {
	a = StrToLower(a)
	b = StrToLower(b)
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func StrCaseHasPrefix(a string, b string) bool {
	return len(a) >= len(b) && StrCaseEquals(a[:len(b)], b)
}

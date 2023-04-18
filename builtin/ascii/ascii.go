package ascii

import "strings"

func IsAscii(c byte) bool {
	return c <= 0x7f
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
	return strings.Map(func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	}, s)
}

// 字符串转小写。与 strings.ToLower() 的区别是，它不支持除英文字母外的其他unicode字母
func StrToLower(s string) string {
	return strings.Map(func(r rune) rune {
		if 'A' <= r && r <= 'Z' {
			return r - 'A' + 'a'
		}
		return r
	}, s)
}

func StrCaseEquals(a string, b string) bool {
	return StrToLower(a) == StrToLower(b)
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

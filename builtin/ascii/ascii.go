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

func StrToLower(s string) string {
	i := 0
	for ; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
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
		buf.WriteByte(ToLower(s[j]))
	}
	return buf.String()
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

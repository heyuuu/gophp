package strutil

import "strings"

func AfterByte(str string, c byte) (string, bool) {
	if pos := strings.IndexByte(str, c); pos >= 0 && pos+1 < len(str) {
		return str[pos+1:], true
	}
	return "", false
}

func LastAfterByte(str string, c byte) (string, bool) {
	if pos := strings.LastIndexByte(str, c); pos >= 0 {
		return str[pos+1:], true
	}
	return "", false
}

func BeforeByte(str string, c byte) (string, bool) {
	if pos := strings.IndexByte(str, c); pos >= 0 {
		return str[:pos], true
	}
	return "", false
}

func LastBeforeByte(str string, c byte) (string, bool) {
	if pos := strings.LastIndexByte(str, c); pos >= 0 {
		return str[:pos], true
	}
	return "", false
}

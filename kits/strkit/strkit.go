package strkit

import "strings"

func Cut(s string, sep string) (before, after string, found bool) {
	return strings.Cut(s, sep)
}

func LastCut(s string, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return "", s, false
}

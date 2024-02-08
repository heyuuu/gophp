package standard

import (
	"strings"
)

const PHP_URL_SCHEME = 0
const PHP_URL_HOST = 1
const PHP_URL_PORT = 2
const PHP_URL_USER = 3
const PHP_URL_PASS = 4
const PHP_URL_PATH = 5
const PHP_URL_QUERY = 6
const PHP_URL_FRAGMENT = 7
const PHP_QUERY_RFC1738 = 1
const PHP_QUERY_RFC3986 = 2

const lcHexChars = "0123456789abcdef"

// --- functions

func PhpUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c == ' ' {
			buf.WriteByte('+')
		} else if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' {
			buf.WriteByte('%')
			buf.WriteByte(lcHexChars[c>>4])
			buf.WriteByte(lcHexChars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func PhpRawUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' && c != '~' {
			buf.WriteByte('%')
			buf.WriteByte(lcHexChars[c>>4])
			buf.WriteByte(lcHexChars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}

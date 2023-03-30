package builtin

import (
	"strings"
)

// 对应 c 函数 strlen
func Strlen(s string) int {
	pos := strings.IndexByte(s, '\000')
	if pos < 0 {
		return len(s)
	} else {
		return pos
	}
}

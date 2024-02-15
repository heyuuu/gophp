package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"strconv"
)

// strtoll(s0, s1, base)
func ParseLong(s string, base int) int {
	value, _ := parseLongPrefix(s, base)
	return value
}

func parseLongPrefix(s string, base int) (value int, n int) {
	pos := 0
	if pos < len(s) && (s[pos] == '-' || s[pos] == '+') {
		pos++
	}
	if base == 0 {
		if pos+1 < len(s) && s[pos] == '0' {
			c := s[pos+1]
			if c == 'b' || c == 'B' {
				base = 2
				pos += 2
			} else if c == 'x' || c == 'X' {
				base = 2
				pos += 2
			} else if ascii.IsDigit(c) {
				base = 8
				pos++
			} else { // `0` | `+0` | `-0`
				return 0, pos + 1
			}
		} else {
			base = 10
		}
	}

	for ; pos < len(s); pos++ {
		digit, ok := ascii.ParseXDigit(s[pos])
		if !ok || int(digit) >= base {
			break
		}
	}
	num, _ := strconv.ParseInt(s[:pos], base, 64)
	return int(num), pos
}

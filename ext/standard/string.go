package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/types"
	"strings"
)

func PhpStringToupper(s *types.String) *types.String {
	return types.NewString(ascii.StrToUpper(s.GetStr()))
}

func PhpStringTolower(s *types.String) *types.String {
	return types.NewString(ascii.StrToLower(s.GetStr()))
}

func substr(str string, offset int, length *int) (string, bool) {
	negativeOffset := offset < 0
	rawLen := len(str)
	// handle offset
	if offset > len(str) {
		return "", false
	} else if offset < 0 {
		/* if "offset" position is negative, count start position from the end
		 * of the string
		 */
		offset += len(str)
		if offset < 0 {
			offset = 0
		}
	}
	if offset > 0 {
		str = str[offset:]
	}

	// handle length
	if length != nil {
		l := *length
		if l < 0 {
			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */
			if -l > len(str) {
				if negativeOffset && -l <= rawLen {
					l = 0
				} else {
					return "", false
				}
			} else {
				l = len(str) + l
			}
		}

		if l < len(str) {
			str = str[:l]
		}
	}

	return str, true
}

func PhpAddslashes(str string) string {
	if str == "" {
		return ""
	}
	if pos := strings.IndexByte(str, '\\'); pos < 0 {
		return str
	}
	replacer := strings.NewReplacer(
		"\\000", "\\0",
		`'`, `\'`,
		`"`, `\"`,
		`\`, `\\`,
	)
	return replacer.Replace(str)
}

func PhpAddcslashes(str string, what string) string {
	mask, _ := PhpCharmaskEx(what)

	strings.NewReplacer()

	var buf strings.Builder
	for _, c := range []byte(str) {
		if strings.ContainsRune(mask, rune(c)) {
			if c < 32 || c > 126 {
				buf.WriteByte('\\')
				switch c {
				case '\n':
					buf.WriteByte('n')
				case '\t':
					buf.WriteByte('t')
				case '\r':
					buf.WriteByte('r')
				case '\a':
					buf.WriteByte('a')
				case '\v':
					buf.WriteByte('v')
				case '\b':
					buf.WriteByte('b')
				case '\f':
					buf.WriteByte('f')
				default:
					buf.WriteString(fmt.Sprintf("%03o", c))
				}
				continue
			}
			buf.WriteByte('\\')
		}
		buf.WriteByte(c)
	}

	return buf.String()
}

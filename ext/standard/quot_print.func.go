package standard

import (
	"github.com/heyuuu/gophp/kits/ascii"
	b "github.com/heyuuu/gophp/php/lang"
	"strings"
)

func PhpHex2int(c byte) byte {
	if ascii.IsDigit(c) {
		return c - '0'
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else {
		panic("unreachable")
	}
}

func ZifQuotedPrintableDecode(str string) string {
	if str == "" {
		return ""
	}

	var buf strings.Builder
	l := len(str)
	for i := 0; i < len(str); {
		switch str[i] {
		case '=':
			if i+2 < l && ascii.IsXDigit(str[i+1]) && ascii.IsXDigit(str[i+2]) {
				c := (PhpHex2int(str[i+1]) << 4) + PhpHex2int(str[i+2])
				buf.WriteByte(c)
				i += 3
			} else {
				k := 1
				for i+k < l && (str[i+k] == 32 || str[i+k] == 9) {
					/* Possibly, skip spaces/tabs at the end of line */
					k++
				}
				if i+k == l {
					/* End of line reached */
					i += k
				} else if i+k+1 < l && str[i+k] == 13 && str[i+k+1] == 10 {
					/* CRLF */
					i += k + 2
				} else if str[i+k] == 13 || str[i+k] == 10 {
					/* CR or LF */
					i += k + 1
				} else {
					buf.WriteByte(str[i])
					i++
				}
			}
		default:
			buf.WriteByte(str[i])
			i++
		}
	}
	return buf.String()
}
func ZifQuotedPrintableEncode(str string) string {
	if str == "" {
		return ""
	}

	var lp = 0
	var hex = "0123456789ABCDEF"
	var buf strings.Builder
	idx := 0
	for length := len(str) - 1; length >= 0; length-- {
		c := str[idx]
		idx++
		if c == '\015' && str[idx] == '\012' && length > 0 {
			// CR & LF
			buf.WriteByte('\015')
			buf.WriteByte('\012')
			idx++
			length--
			lp = 0
		} else {
			if ascii.IsControl(c) || (c&0x80) != 0 || c == '=' || c == ' ' && str[idx] == '0' {
				if b.Assign(&lp, lp+3) > PHP_QPRINT_MAXL && c <= 0x7f || c > 0x7f && c <= 0xdf && lp+3 > PHP_QPRINT_MAXL || c > 0xdf && c <= 0xef && lp+6 > PHP_QPRINT_MAXL || c > 0xef && c <= 0xf4 && lp+9 > PHP_QPRINT_MAXL {
					buf.WriteString("=00")
					lp = 3
				}
				buf.Write([]byte{'=', hex[c>>4], hex[c&0xf]})
			} else {
				if b.PreInc(&lp) > PHP_QPRINT_MAXL {
					buf.WriteString("=00")
					lp = 1
				}
				buf.WriteByte(c)
			}
		}
	}
	return buf.String()
}

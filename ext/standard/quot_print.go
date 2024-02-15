package standard

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"strings"
)

const qprintMaxL = 75

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
				c1, _ := ascii.ParseXDigit(str[i+1])
				c2, _ := ascii.ParseXDigit(str[i+2])
				buf.WriteByte(c1<<4 + c2)
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
	var hex = ucHexChars
	var buf strings.Builder

	for i := 0; i < len(str); i++ {
		var c, nc byte = str[i], 0
		if i+1 < len(str) {
			nc = str[i+1]
		}
		if c == '\015' && nc == '\012' {
			// CR & LF
			buf.WriteByte('\015')
			i++
			lp = 0
		} else {
			if ascii.IsControl(c) || (c&0x80) != 0 || c == '=' || c == ' ' && nc == '\015' {
				lp += 3
				if lp > qprintMaxL && c <= 0x7f || c > 0x7f && c <= 0xdf && lp+3 > qprintMaxL || c > 0xdf && c <= 0xef && lp+6 > qprintMaxL || c > 0xef && c <= 0xf4 && lp+9 > qprintMaxL {
					buf.WriteString("=\015\012")
					lp = 3
				}
				buf.Write([]byte{'=', hex[c>>4], hex[c&0xf]})
			} else {
				lp++
				if lp > qprintMaxL {
					buf.WriteString("=\015\012")
					lp = 1
				}
				buf.WriteByte(c)
			}
		}
	}
	return buf.String()
}

// <<generate>>

package zend

import (
	b "sik/builtin"
)

/* Windows uses VK_ESCAPE instead of \e */

const VK_ESCAPE = 'e'

func ZendComputeEscapedStringLen(s *byte, l int) int {
	var i int
	var len_ int = l
	for i = 0; i < l; i++ {
		var c byte = s[i]
		if c == '\n' || c == '\r' || c == '\t' || c == 'f' || c == 'v' || c == '\\' || c == VK_ESCAPE {
			len_ += 1
		} else if c < 32 || c > 126 {
			len_ += 3
		}
	}
	return len_
}
func SmartStrAppendEscaped(str *SmartStr, s *byte, l int) {
	var res *byte
	var i int
	var len_ int = ZendComputeEscapedStringLen(s, l)
	str.Alloc(len_)
	res = &str.GetS().GetVal()[str.GetS().GetLen()]
	str.GetS().GetLen() += len_
	for i = 0; i < l; i++ {
		var c uint8 = s[i]
		if c < 32 || c == '\\' || c > 126 {
			b.PostInc(&(*res)) = '\\'
			switch c {
			case '\n':
				b.PostInc(&(*res)) = 'n'
				break
			case '\r':
				b.PostInc(&(*res)) = 'r'
				break
			case '\t':
				b.PostInc(&(*res)) = 't'
				break
			case 'f':
				b.PostInc(&(*res)) = 'f'
				break
			case 'v':
				b.PostInc(&(*res)) = 'v'
				break
			case '\\':
				b.PostInc(&(*res)) = '\\'
				break
			case VK_ESCAPE:
				b.PostInc(&(*res)) = 'e'
				break
			default:
				b.PostInc(&(*res)) = 'x'
				if c>>4 < 10 {
					b.PostInc(&(*res)) = (c >> 4) + '0'
				} else {
					b.PostInc(&(*res)) = (c >> 4) + 'A' - 10
				}
				if (c & 0xf) < 10 {
					b.PostInc(&(*res)) = (c & 0xf) + '0'
				} else {
					b.PostInc(&(*res)) = (c & 0xf) + 'A' - 10
				}
			}
		} else {
			b.PostInc(&(*res)) = c
		}
	}
}
func SmartStrAppendPrintf(dest *SmartStr, format string, _ ...any) {
	var arg va_list
	va_start(arg, format)
	ZendPrintfToSmartStr(dest, format, arg)
	va_end(arg)
}

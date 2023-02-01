// <<generate>>

package zend

import (
	b "sik/builtin"
)

func SmartStrAppendsEx(dest *SmartStr, src *byte, what ZendBool) {
	SmartStrAppendlEx(dest, src, strlen(src), what)
}
func SmartStrAppends(dest *SmartStr, src *byte)            { SmartStrAppendl(dest, src, strlen(src)) }
func SmartStrAppendc(dest *SmartStr, c byte)               { SmartStrAppendcEx(dest, c, 0) }
func SmartStrAppendl(dest *SmartStr, src *byte, len_ int)  { SmartStrAppendlEx(dest, src, len_, 0) }
func SmartStrAppend(dest *SmartStr, src *ZendString)       { SmartStrAppendEx(dest, src, 0) }
func SmartStrAppendSmartStr(dest *SmartStr, src *SmartStr) { SmartStrAppendSmartStrEx(dest, src, 0) }
func SmartStrAppendLong(dest *SmartStr, val ZendLong)      { SmartStrAppendLongEx(dest, val, 0) }
func SmartStrAppendUnsigned(dest *SmartStr, val ZendUlong) { SmartStrAppendUnsignedEx(dest, val, 0) }
func SmartStrFree(dest *SmartStr)                          { SmartStrFreeEx(dest, 0) }
func SmartStrAlloc(str *SmartStr, len_ int, persistent ZendBool) int {
	if str.GetS() == nil {
		goto do_smart_str_realloc
	} else {
		len_ += str.GetS().GetLen()
		if len_ >= str.GetA() {
		do_smart_str_realloc:
			if persistent != 0 {
				SmartStrRealloc(str, len_)
			} else {
				SmartStrErealloc(str, len_)
			}
		}
	}
	return len_
}
func SmartStrFreeEx(str *SmartStr, persistent ZendBool) {
	if str.GetS() != nil {
		ZendStringReleaseEx(str.GetS(), persistent)
		str.SetS(nil)
	}
	str.SetA(0)
}
func SmartStr0(str *SmartStr) {
	if str.GetS() != nil {
		str.GetS().GetVal()[str.GetS().GetLen()] = '0'
	}
}
func SmartStrAppendcEx(dest *SmartStr, ch byte, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, 1, persistent)
	dest.GetS().GetVal()[new_len-1] = ch
	dest.GetS().GetLen() = new_len
}
func SmartStrAppendlEx(dest *SmartStr, str string, len_ int, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, len_, persistent)
	memcpy(dest.GetS().GetVal()+dest.GetS().GetLen(), str, len_)
	dest.GetS().GetLen() = new_len
}
func SmartStrAppendEx(dest *SmartStr, src *ZendString, persistent ZendBool) {
	SmartStrAppendlEx(dest, src.GetVal(), src.GetLen(), persistent)
}
func SmartStrAppendSmartStrEx(dest *SmartStr, src *SmartStr, persistent ZendBool) {
	if src.GetS() != nil && src.GetS().GetLen() != 0 {
		SmartStrAppendEx(dest, src.GetS(), persistent)
	}
}
func SmartStrAppendLongEx(dest *SmartStr, num ZendLong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintLongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStrAppendUnsignedEx(dest *SmartStr, num ZendUlong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintUlongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStrSetl(dest *SmartStr, src *byte, len_ int) {
	SmartStrFree(dest)
	SmartStrAppendl(dest, src, len_)
}
func SMART_STR_NEW_LEN(len_ int) int {
	return ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STR_OVERHEAD, SMART_STR_PAGE) - SMART_STR_OVERHEAD
}
func SmartStrErealloc(str *SmartStr, len_ int) {
	if str.GetS() == nil {
		if len_ <= SMART_STR_START_LEN {
			str.SetA(SMART_STR_START_LEN)
		} else {
			str.SetA(SMART_STR_NEW_LEN(len_))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 0))
		str.GetS().GetLen() = 0
	} else {
		str.SetA(SMART_STR_NEW_LEN(len_))
		str.SetS((*ZendString)(Erealloc2(str.GetS(), str.GetA()+_ZSTR_HEADER_SIZE+1, _ZSTR_HEADER_SIZE+str.GetS().GetLen())))
	}
}
func SmartStrRealloc(str *SmartStr, len_ int) {
	if str.GetS() == nil {
		if len_ <= SMART_STR_START_LEN {
			str.SetA(SMART_STR_START_LEN)
		} else {
			str.SetA(SMART_STR_NEW_LEN(len_))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 1))
		str.GetS().GetLen() = 0
	} else {
		str.SetA(SMART_STR_NEW_LEN(len_))
		str.SetS((*ZendString)(Perealloc(str.GetS(), str.GetA()+_ZSTR_HEADER_SIZE+1, 1)))
	}
}
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
	SmartStrAlloc(str, len_, 0)
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

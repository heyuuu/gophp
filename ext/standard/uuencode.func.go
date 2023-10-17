package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/zend/faults"
	"math"
	"strings"
)

func PHP_UU_ENC(c byte) byte {
	if c != 0 {
		return (c & 077) + ' '
	} else {
		return '`'
	}
}
func PHP_UU_ENC_C2(c int) __auto__ {
	return PHP_UU_ENC((*c)<<4&060 | (*(c + 1))>>4&017)
}
func PHP_UU_ENC_C3(c int) __auto__ {
	return PHP_UU_ENC((*(c + 1))<<2&074 | (*(c + 2))>>6&3)
}
func PHP_UU_DEC(c byte) byte { return (c - ' ') & 077 }
func PhpUuencode(src string) string {
	var len_ int = 45
	var s *uint8
	var e *uint8
	var ee *uint8
	var buf strings.Builder
	var srcLen = len(src)

	/* encoded length is ~ 38% greater than the original
	   Use 1.5 for easier calculation.
	*/
	buf.Grow(srcLen/2*3 + 46)
	s = (*uint8)(src)
	e = s + srcLen
	for s+3 < e {
		ee = s + len_
		if ee > e {
			ee = e
			len_ = ee - s
			if len_%3 != 0 {
				ee = s + int(floor(float64(len_/3))*3)
			}
		}
		buf.WriteByte()
		buf.WriteByte(PHP_UU_ENC(len_))
		for s < ee {
			buf.WriteByte(PHP_UU_ENC((*s) >> 2))
			buf.WriteByte(PHP_UU_ENC_C2(s))
			buf.WriteByte(PHP_UU_ENC_C3(s))
			buf.WriteByte(PHP_UU_ENC((*(s + 2)) & 077))
			s += 3
		}
		if len_ == 45 {
			buf.WriteByte('\n')
		}
	}
	if s < e {
		if len_ == 45 {
			buf.WriteByte(PHP_UU_ENC(e - s))
			len_ = 0
		}
		buf.WriteByte(PHP_UU_ENC((*s) >> 2))
		buf.WriteByte(PHP_UU_ENC_C2(s))
		if e-s > 1 {
			buf.WriteByte(PHP_UU_ENC_C3(s))
		} else {
			buf.WriteByte(PHP_UU_ENC('0'))
		}
		if e-s > 2 {
			buf.WriteByte(PHP_UU_ENC((*(s + 2)) & 077))
		} else {
			buf.WriteByte(PHP_UU_ENC('0'))
		}
	}
	if len_ < 45 {
		buf.WriteByte('\n')
	}
	buf.WriteByte(PHP_UU_ENC('0'))
	buf.WriteByte('\n')
	return buf.String()
}
func PhpUudecode(src string) (string, bool) {
	var len_ int
	var totalLen int = 0
	var s *byte
	var buf strings.Builder

	srcLen := len(src)
	buf.Grow(int(math.Ceil(float64(srcLen) * 0.75)))

	idx := 0
	for idx < srcLen {
		len_ := int(PHP_UU_DEC(src[idx]))
		idx++
		if len_ == 0 {
			break
		}

		/* sanity check */
		if len_ > srcLen {
			return "", false
		}
		totalLen += len_
		idx2 := idx + lang.Cond(len_ == 45, 60, int(math.Floor(float64(len_)*1.33)))

		/* sanity check */
		if idx2 >= srcLen {
			return "", false
		}
		for ; idx < idx2; idx += 4 {
			if idx+4 >= srcLen {
				return "", false
			}
			buf.WriteByte(PHP_UU_DEC(src[idx])<<2 | PHP_UU_DEC(src[idx+1])>>4)
			buf.WriteByte(PHP_UU_DEC(src[idx+1])<<4 | PHP_UU_DEC(src[idx+2])>>2)
			buf.WriteByte(PHP_UU_DEC(src[idx+2])<<6 | PHP_UU_DEC(src[idx+3]))
		}
		if len_ < 45 {
			break
		}

		/* skip \n */
		idx++
	}
	if len_ = totalLen; len_ > buf.Len() {
		buf.WriteByte(PHP_UU_DEC(*s)<<2 | PHP_UU_DEC(src[idx+1])>>4)
		if len_ > 1 {
			buf.WriteByte(PHP_UU_DEC(src[idx+1])<<4 | PHP_UU_DEC(src[idx+2])>>2)
			if len_ > 2 {
				buf.WriteByte(PHP_UU_DEC(src[idx+2])<<6 | PHP_UU_DEC(src[idx+3]))
			}
		}
	}
	return buf.String()[:totalLen], true
}
func ZifConvertUuencode(data string) (string, bool) {
	// notice: PHP < 8.0 时，对空字符串输入会返回 false, >= 8.0 后取消了这个逻辑
	if len(data) == 0 {
		return "", false
	}
	return PhpUuencode(data), true
}
func ZifConvertUudecode(data string) (string, bool) {
	if len(data) == 0 {
		return "", false
	}

	if result, ok := PhpUudecode(data); ok {
		return result, true
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "The given parameter is not a valid uuencoded string")
		return "", false
	}
}

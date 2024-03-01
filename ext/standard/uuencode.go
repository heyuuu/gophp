package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"strings"
)

func uuencodeByte(c byte) byte {
	if c != 0 {
		return (c & 077) + ' '
	} else {
		return '`'
	}
}
func uuencode(src string) string {
	var srcLen = len(src)
	var buf strings.Builder
	buf.Grow(srcLen/2*3 + 46)

	/* encoded length is ~ 38% greater than the original
	   Use 1.5 for easier calculation.
	*/
	for i := 0; i < srcLen; i += 3 {
		// line start
		if i%45 == 0 {
			// new line
			if i != 0 {
				buf.WriteByte('\n')
			}
			// line length
			length := min(45, srcLen-i)
			buf.WriteByte(uuencodeByte(byte(length)))
		}
		// chunk
		// 使用临时数组截取最多3位字节，不足部分用0填充
		var c [3]byte
		copy(c[:], src[i:])
		buf.Write([]byte{
			uuencodeByte(c[0] >> 2),
			uuencodeByte(c[0]<<4&0b110000 | c[1]>>4&0b001111),
			uuencodeByte(c[1]<<2&0b111100 | c[2]>>6&0b000011),
			uuencodeByte(c[2] & 0b111111),
		})
	}
	buf.WriteByte('\n')
	buf.WriteByte(uuencodeByte(0))
	buf.WriteByte('\n')
	return buf.String()
}

func uudecodeByte(c byte) byte { return (c - ' ') & 077 }
func uudecode(src string) (string, bool) {
	var srcLen = len(src)
	var buf strings.Builder
	buf.Grow(srcLen / 4 * 3)

	var totalLen int = 0
	idx := 0
	for idx < srcLen {
		// line length
		length := int(uudecodeByte(src[idx]))
		idx++
		if length == 0 {
			break
		}
		totalLen += length

		// read each chunk
		for i := 0; i < length; i += 3 {
			/* sanity check */
			if idx+4 > srcLen {
				return "", false
			}

			var c [4]byte
			copy(c[:], src[idx:])
			idx += 4

			buf.Write([]byte{
				uudecodeByte(c[0])<<2 | uudecodeByte(c[1])>>4,
				uudecodeByte(c[1])<<4 | uudecodeByte(c[2])>>2,
				uudecodeByte(c[2])<<6 | uudecodeByte(c[3]),
			})
		}

		if length < 45 {
			break
		}
		/* skip \n */
		idx++
	}
	return buf.String()[:totalLen], true
}

// @zif(onError=1)
func ZifConvertUuencode(data string) (string, bool) {
	// notice: PHP < 8.0 时，对空字符串输入会返回 false, >= 8.0 后取消了这个逻辑
	if len(data) == 0 {
		return "", false
	}
	return uuencode(data), true
}

// @zif(onError=1)
func ZifConvertUudecode(ctx *php.Context, data string) (string, bool) {
	if len(data) == 0 {
		return "", false
	}

	if result, ok := uudecode(data); ok {
		return result, true
	} else {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The given parameter is not a valid uuencoded string")
		return "", false
	}
}

package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func LIMIT_ALL(all __auto__, doctype int, charset EntityCharset) {
	all = all && !(CHARSET_PARTIAL_SUPPORT(charset)) && doctype != ENT_HTML_DOC_XML1
}
func CHECK_LEN(pos int, chars_need int) bool { return str_len-pos >= chars_need }
func Utf8Lead(c uint8) bool {
	return c < 0x80 || c >= 0xc2 && c <= 0xf4
}
func Utf8Trail(c uint8) bool { return c >= 0x80 && c <= 0xbf }
func Gb2312Lead(c uint8) bool {
	return c != 0x8e && c != 0x8f && c != 0xa0 && c != 0xff
}
func Gb2312Trail(c uint8) bool { return c >= 0xa1 && c <= 0xfe }
func SjisLead(c uint8) bool {
	return c != 0x80 && c != 0xa0 && c < 0xfd
}
func SjisTrail(c uint8) bool {
	return c >= 0x40 && c != 0x7f && c < 0xfd
}
func GetDefaultCharset() string {
	if core.PG__().internal_encoding && core.PG__().internal_encoding[0] {
		return core.PG__().internal_encoding
	} else if core.SG__().default_charset && core.SG__().default_charset[0] {
		return core.SG__().default_charset
	}
	return ""
}
func GetNextChar(charset EntityCharset, str *uint8, str_len int, cursor *int, status *int) uint {
	var pos int = *cursor
	var this_char uint = 0
	*status = types.SUCCESS
	b.Assert(pos <= str_len)
	if !(CHECK_LEN(pos, 1)) {
		*cursor = pos + 1
		*status = types.FAILURE
		return 0
	}
	switch charset {
	case CsUtf8:

		/* We'll follow strategy 2. from section 3.6.1 of UTR #36:
		 * "In a reported illegal byte sequence, do not include any
		 *  non-initial byte that encodes a valid character or is a leading
		 *  byte for a valid sequence." */

		var c uint8
		c = str[pos]
		if c < 0x80 {
			this_char = c
			pos++
		} else if c < 0xc2 {
			*cursor = pos + 1
			*status = types.FAILURE
			return 0
		} else if c < 0xe0 {
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			if !(Utf8Trail(str[pos+1])) {
				*cursor = pos + lang.Cond(Utf8Lead(str[pos+1]), 1, 2)
				*status = types.FAILURE
				return 0
			}
			this_char = (c&0x1f)<<6 | str[pos+1]&0x3f
			if this_char < 0x80 {
				*cursor = pos + 2
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else if c < 0xf0 {
			var avail int = str_len - pos
			if avail < 3 || !(Utf8Trail(str[pos+1])) || !(Utf8Trail(str[pos+2])) {
				if avail < 2 || Utf8Lead(str[pos+1]) {
					*cursor = pos + 1
					*status = types.FAILURE
					return 0
				} else if avail < 3 || Utf8Lead(str[pos+2]) {
					*cursor = pos + 2
					*status = types.FAILURE
					return 0
				} else {
					*cursor = pos + 3
					*status = types.FAILURE
					return 0
				}
			}
			this_char = (c&0xf)<<12 | (str[pos+1]&0x3f)<<6 | str[pos+2]&0x3f
			if this_char < 0x800 {
				*cursor = pos + 3
				*status = types.FAILURE
				return 0
			} else if this_char >= 0xd800 && this_char <= 0xdfff {
				*cursor = pos + 3
				*status = types.FAILURE
				return 0
			}
			pos += 3
		} else if c < 0xf5 {
			var avail int = str_len - pos
			if avail < 4 || !(Utf8Trail(str[pos+1])) || !(Utf8Trail(str[pos+2])) || !(Utf8Trail(str[pos+3])) {
				if avail < 2 || Utf8Lead(str[pos+1]) {
					*cursor = pos + 1
					*status = types.FAILURE
					return 0
				} else if avail < 3 || Utf8Lead(str[pos+2]) {
					*cursor = pos + 2
					*status = types.FAILURE
					return 0
				} else if avail < 4 || Utf8Lead(str[pos+3]) {
					*cursor = pos + 3
					*status = types.FAILURE
					return 0
				} else {
					*cursor = pos + 4
					*status = types.FAILURE
					return 0
				}
			}
			this_char = (c&0x7)<<18 | (str[pos+1]&0x3f)<<12 | (str[pos+2]&0x3f)<<6 | str[pos+3]&0x3f
			if this_char < 0x10000 || this_char > 0x10ffff {
				*cursor = pos + 4
				*status = types.FAILURE
				return 0
			}
			pos += 4
		} else {
			*cursor = pos + 1
			*status = types.FAILURE
			return 0
		}
	case CsBig5:

		/* reference http://demo.icu-project.org/icu-bin/convexp?conv=big5 */

		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0xfe {
			var next uint8
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
				this_char = c<<8 | next
			} else {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else {
			this_char = c
			pos += 1
		}
	case CsBig5hkscs:
		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0xfe {
			var next uint8
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
				this_char = c<<8 | next
			} else if next != 0x80 && next != 0xff {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else {
			this_char = c
			pos += 1
		}
	case CsGb2312:
		var c uint8 = str[pos]
		if c >= 0xa1 && c <= 0xfe {
			var next uint8
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if Gb2312Trail(next) {
				this_char = c<<8 | next
			} else if Gb2312Lead(next) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else if Gb2312Lead(c) {
			this_char = c
			pos += 1
		} else {
			*cursor = pos + 1
			*status = types.FAILURE
			return 0
		}
	case CsSjis:
		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0x9f || c >= 0xe0 && c <= 0xfc {
			var next uint8
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if SjisTrail(next) {
				this_char = c<<8 | next
			} else if SjisLead(next) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else if c < 0x80 || c >= 0xa1 && c <= 0xdf {
			this_char = c
			pos += 1
		} else {
			*cursor = pos + 1
			*status = types.FAILURE
			return 0
		}
	case CsEucjp:
		var c uint8 = str[pos]
		if c >= 0xa1 && c <= 0xfe {
			var next uint
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0xa1 && next <= 0xfe {

				/* this a jis kanji char */

				this_char = c<<8 | next

				/* this a jis kanji char */

			} else {
				*cursor = pos + lang.Cond(next != 0xa0 && next != 0xff, 1, 2)
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else if c == 0x8e {
			var next uint
			if !(CHECK_LEN(pos, 2)) {
				*cursor = pos + 1
				*status = types.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0xa1 && next <= 0xdf {

				/* JIS X 0201 kana */

				this_char = c<<8 | next

				/* JIS X 0201 kana */

			} else {
				*cursor = pos + lang.Cond(next != 0xa0 && next != 0xff, 1, 2)
				*status = types.FAILURE
				return 0
			}
			pos += 2
		} else if c == 0x8f {
			var avail int = str_len - pos
			if avail < 3 || !(str[pos+1] >= 0xa1 && str[pos+1] <= 0xfe) || !(str[pos+2] >= 0xa1 && str[pos+2] <= 0xfe) {
				if avail < 2 || str[pos+1] != 0xa0 && str[pos+1] != 0xff {
					*cursor = pos + 1
					*status = types.FAILURE
					return 0
				} else if avail < 3 || str[pos+2] != 0xa0 && str[pos+2] != 0xff {
					*cursor = pos + 2
					*status = types.FAILURE
					return 0
				} else {
					*cursor = pos + 3
					*status = types.FAILURE
					return 0
				}
			} else {

				/* JIS X 0212 hojo-kanji */

				this_char = c<<16 | str[pos+1]<<8 | str[pos+2]

				/* JIS X 0212 hojo-kanji */

			}
			pos += 3
		} else if c != 0xa0 && c != 0xff {

			/* character encoded in 1 code unit */

			this_char = c
			pos += 1
		} else {
			*cursor = pos + 1
			*status = types.FAILURE
			return 0
		}
	default:

		/* single-byte charsets */

		this_char = str[lang.PostInc(&pos)]
	}
	*cursor = pos
	return this_char
}
func PhpNextUtf8Char(str *uint8, str_len int, cursor *int, status *int) uint {
	return GetNextChar(CsUtf8, str, str_len, cursor, status)
}
func DetermineCharset(charset_hint *byte) EntityCharset {
	var i int
	var charset EntityCharset = CsUtf8
	var len_ int = 0

	/* Default is now UTF-8 */

	if charset_hint == nil {
		return CsUtf8
	}
	if lang.Assign(&len_, strlen(charset_hint)) != 0 {
		goto det_charset
	}
	charset_hint = core.SG__().default_charset
	if charset_hint != nil && lang.Assign(&len_, strlen(charset_hint)) != 0 {
		goto det_charset
	}

	/* try to detect the charset for the locale */

	/* try to figure out the charset from the locale */

	var localename *byte
	var dot *byte
	var at *byte

	/* lang[_territory][.codeset][@modifier] */

	localename = setlocale(LC_CTYPE, nil)
	dot = strchr(localename, '.')
	if dot != nil {
		dot++

		/* locale specifies a codeset */

		at = strchr(dot, '@')
		if at != nil {
			len_ = at - dot
		} else {
			len_ = strlen(dot)
		}
		charset_hint = dot
	} else {

		/* no explicit name; see if the name itself
		 * is the charset */

		charset_hint = localename
		len_ = strlen(charset_hint)
	}
det_charset:
	if charset_hint != nil {
		var found bool
		/* look for the codeset */
		charset, found = GetCharset(b.CastStr(charset_hint, len_))
		if !found {
			core.PhpErrorDocref("", faults.E_WARNING, "charset `%s' not supported, assuming utf-8", charset_hint)
		}
	}
	return charset
}
func PhpUtf32Utf8Ex(k uint) []byte {
	/* assert(0x0 <= k <= 0x10FFFF); */
	/* UTF-8 has been restricted to max 4 bytes since RFC 3629 */
	if k < 0x80 {
		return []byte{
			byte(k),
		}
	} else if k < 0x800 {
		return []byte{
			byte(0xc0 | k>>6),
			byte(0x80 | k&0x3f),
		}
	} else if k < 0x10000 {
		return []byte{
			byte(0xe0 | k>>12),
			byte(0x80 | k>>6&0x3f),
			byte(0x80 | k&0x3f),
		}
	} else {
		return []byte{
			byte(0xf0 | k>>18),
			byte(0x80 | k>>12&0x3f),
			byte(0x80 | k>>6&0x3f),
			byte(0x80 | k&0x3f),
		}
	}
}
func PhpUtf32Utf8(buf *uint8, k uint) int {
	var retval int = 0

	/* assert(0x0 <= k <= 0x10FFFF); */

	if k < 0x80 {
		buf[0] = k
		retval = 1
	} else if k < 0x800 {
		buf[0] = 0xc0 | k>>6
		buf[1] = 0x80 | k&0x3f
		retval = 2
	} else if k < 0x10000 {
		buf[0] = 0xe0 | k>>12
		buf[1] = 0x80 | k>>6&0x3f
		buf[2] = 0x80 | k&0x3f
		retval = 3
	} else {
		buf[0] = 0xf0 | k>>18
		buf[1] = 0x80 | k>>12&0x3f
		buf[2] = 0x80 | k>>6&0x3f
		buf[3] = 0x80 | k&0x3f
		retval = 4
	}

	/* UTF-8 has been restricted to max 4 bytes since RFC 3629 */

	return retval
}
func UnimapBsearch(table *UniToEnc, code_key_a uint, num int) uint8 {
	var l *UniToEnc = table
	var h *UniToEnc = &table[num-1]
	var m *UniToEnc
	var code_key uint16

	/* we have no mappings outside the BMP */

	if code_key_a > 0xffff {
		return 0
	}
	code_key = uint16(code_key_a)
	for l <= h {
		m = l + (h-l)/2
		if code_key < m.GetUnCodePoint() {
			h = m - 1
		} else if code_key > m.GetUnCodePoint() {
			l = m + 1
		} else {
			return m.GetCsCode()
		}
	}
	return 0
}
func MapFromUnicode(code uint, charset EntityCharset, res *uint) int {
	var found uint8
	var table *UniToEnc
	var table_size int
	switch charset {
	case Cs88591:

		/* identity mapping of code points to unicode */

		if code > 0xff {
			return types.FAILURE
		}
		*res = code
	case Cs88595:
		if code <= 0xa0 || code == 0xad {
			*res = code
		} else if code == 0x2116 {
			*res = 0xf0
		} else if code == 0xa7 {
			*res = 0xfd
		} else if code >= 0x401 && code <= 0x44f {
			if code == 0x40d || code == 0x450 || code == 0x45d {
				return types.FAILURE
			}
			*res = code - 0x360
		} else {
			return types.FAILURE
		}
	case Cs885915:
		if code < 0xa4 || code > 0xbe && code <= 0xff {
			*res = code
		} else {
			found = UnimapBsearch(UnimapIso885915, code, b.SizeOf("unimap_iso885915")/b.SizeOf("* unimap_iso885915"))
			if found != 0 {
				*res = found
			} else {
				return types.FAILURE
			}
		}
	case CsCp1252:
		if code <= 0x7f || code >= 0xa0 && code <= 0xff {
			*res = code
		} else {
			found = UnimapBsearch(UnimapWin1252, code, b.SizeOf("unimap_win1252")/b.SizeOf("* unimap_win1252"))
			if found != 0 {
				*res = found
			} else {
				return types.FAILURE
			}
		}
	case CsMacroman:
		if code == 0x7f {
			return types.FAILURE
		}
		table = UnimapMacroman
		table_size = b.SizeOf("unimap_macroman") / b.SizeOf("* unimap_macroman")
		goto table_over_7F
	case CsCp1251:
		table = UnimapWin1251
		table_size = b.SizeOf("unimap_win1251") / b.SizeOf("* unimap_win1251")
		goto table_over_7F
	case CsKoi8r:
		table = UnimapKoi8r
		table_size = b.SizeOf("unimap_koi8r") / b.SizeOf("* unimap_koi8r")
		goto table_over_7F
	case CsCp866:
		table = UnimapCp866
		table_size = b.SizeOf("unimap_cp866") / b.SizeOf("* unimap_cp866")
	table_over_7F:
		if code <= 0x7f {
			*res = code
		} else {
			found = UnimapBsearch(table, code, table_size)
			if found != 0 {
				*res = found
			} else {
				return types.FAILURE
			}
		}
	case CsSjis:
		fallthrough
	case CsEucjp:

		/* we interpret 0x5C as the Yen symbol. This is not universal.
		 * See <http://www.w3.org/Submission/japanese-xml/#ambiguity_of_yen> */

		if code >= 0x20 && code <= 0x7d {
			if code == 0x5c {
				return types.FAILURE
			}
			*res = code
		} else {
			return types.FAILURE
		}
	case CsBig5:
		fallthrough
	case CsBig5hkscs:
		fallthrough
	case CsGb2312:
		if code >= 0x20 && code <= 0x7d {
			*res = code
		} else {
			return types.FAILURE
		}
	default:
		return types.FAILURE
	}
	return types.SUCCESS
}
func MapToUnicode(code uint, table *EncToUni, res *uint) {
	/* only single byte encodings are currently supported; assumed code <= 0xFF */

	*res = table.GetInner()[ENT_ENC_TO_UNI_STAGE1(code)].GetUniCp()[ENT_ENC_TO_UNI_STAGE2(code)]

	/* only single byte encodings are currently supported; assumed code <= 0xFF */
}
func UnicodeCpIsAllowed(uni_cp uint, document_type int) int {
	/* XML 1.0                HTML 4.01            HTML 5
	 * 0x09..0x0A            0x09..0x0A            0x09..0x0A
	 * 0x0D                    0x0D                0x0C..0x0D
	 * 0x0020..0xD7FF        0x20..0x7E            0x20..0x7E
	 *                        0x00A0..0xD7FF        0x00A0..0xD7FF
	 * 0xE000..0xFFFD        0xE000..0x10FFFF    0xE000..0xFDCF
	 * 0x010000..0x10FFFF                        0xFDF0..0x10FFFF (*)
	 *
	 * (*) exclude code points where ((code & 0xFFFF) >= 0xFFFE)
	 *
	 * References:
	 * XML 1.0:   <http://www.w3.org/TR/REC-xml/#charsets>
	 * HTML 4.01: <http://www.w3.org/TR/1999/PR-html40-19990824/sgml/sgmldecl.html>
	 * HTML 5:    <http://dev.w3.org/html5/spec/Overview.html#preprocessing-the-input-stream>
	 *
	 * Not sure this is the relevant part for HTML 5, though. I opted to
	 * disallow the characters that would result in a parse error when
	 * preprocessing of the input stream. See also section 8.1.3.
	 *
	 * It's unclear if XHTML 1.0 allows C1 characters. I'll opt to apply to
	 * XHTML 1.0 the same rules as for XML 1.0.
	 * See <http://cmsmcq.com/2007/C1.xml>.
	 */

	switch document_type {
	case ENT_HTML_DOC_HTML401:
		return uni_cp >= 0x20 && uni_cp <= 0x7e || (uni_cp == 0xa || uni_cp == 0x9 || uni_cp == 0xd) || uni_cp >= 0xa0 && uni_cp <= 0xd7ff || uni_cp >= 0xe000 && uni_cp <= 0x10ffff
	case ENT_HTML_DOC_HTML5:
		return uni_cp >= 0x20 && uni_cp <= 0x7e || uni_cp >= 0x9 && uni_cp <= 0xd && uni_cp != 0xb || uni_cp >= 0xa0 && uni_cp <= 0xd7ff || uni_cp >= 0xe000 && uni_cp <= 0x10ffff && (uni_cp&0xffff) < 0xfffe && (uni_cp < 0xfdd0 || uni_cp > 0xfdef)
	case ENT_HTML_DOC_XHTML:
		fallthrough
	case ENT_HTML_DOC_XML1:
		return uni_cp >= 0x20 && uni_cp <= 0xd7ff || (uni_cp == 0xa || uni_cp == 0x9 || uni_cp == 0xd) || uni_cp >= 0xe000 && uni_cp <= 0x10ffff && uni_cp != 0xfffe && uni_cp != 0xffff
	default:
		return 1
	}

	/* XML 1.0                HTML 4.01            HTML 5
	 * 0x09..0x0A            0x09..0x0A            0x09..0x0A
	 * 0x0D                    0x0D                0x0C..0x0D
	 * 0x0020..0xD7FF        0x20..0x7E            0x20..0x7E
	 *                        0x00A0..0xD7FF        0x00A0..0xD7FF
	 * 0xE000..0xFFFD        0xE000..0x10FFFF    0xE000..0xFDCF
	 * 0x010000..0x10FFFF                        0xFDF0..0x10FFFF (*)
	 *
	 * (*) exclude code points where ((code & 0xFFFF) >= 0xFFFE)
	 *
	 * References:
	 * XML 1.0:   <http://www.w3.org/TR/REC-xml/#charsets>
	 * HTML 4.01: <http://www.w3.org/TR/1999/PR-html40-19990824/sgml/sgmldecl.html>
	 * HTML 5:    <http://dev.w3.org/html5/spec/Overview.html#preprocessing-the-input-stream>
	 *
	 * Not sure this is the relevant part for HTML 5, though. I opted to
	 * disallow the characters that would result in a parse error when
	 * preprocessing of the input stream. See also section 8.1.3.
	 *
	 * It's unclear if XHTML 1.0 allows C1 characters. I'll opt to apply to
	 * XHTML 1.0 the same rules as for XML 1.0.
	 * See <http://cmsmcq.com/2007/C1.xml>.
	 */
}
func NumericEntityIsAllowed(uni_cp uint, document_type int) int {
	/* less restrictive than unicode_cp_is_allowed */

	switch document_type {
	case ENT_HTML_DOC_HTML401:

		/* all non-SGML characters (those marked with UNUSED in DESCSET) should be
		 * representable with numeric entities */

		return uni_cp <= 0x10ffff
	case ENT_HTML_DOC_HTML5:

		/* 8.1.4. The numeric character reference forms described above are allowed to
		 * reference any Unicode code point other than U+0000, U+000D, permanently
		 * undefined Unicode characters (noncharacters), and control characters other
		 * than space characters (U+0009, U+000A, U+000C and U+000D) */

		return uni_cp >= 0x20 && uni_cp <= 0x7e || uni_cp >= 0x9 && uni_cp <= 0xc && uni_cp != 0xb || uni_cp >= 0xa0 && uni_cp <= 0x10ffff && (uni_cp&0xffff) < 0xfffe && (uni_cp < 0xfdd0 || uni_cp > 0xfdef)
	case ENT_HTML_DOC_XHTML:
		fallthrough
	case ENT_HTML_DOC_XML1:

		/* OTOH, XML 1.0 requires "character references to match the production for Char
		 * See <http://www.w3.org/TR/REC-xml/#NT-CharRef> */

		return UnicodeCpIsAllowed(uni_cp, document_type)
	default:
		return 1
	}

	/* less restrictive than unicode_cp_is_allowed */
}
func ProcessNumericEntity(buf **byte, code_point *uint) int {
	var code_l zend.ZendLong
	var hexadecimal int = (*(*buf)) == 'x' || (*(*buf)) == 'X'
	var endptr *byte
	if hexadecimal != 0 && (*(*buf)) != '0' {
		*buf++
	}

	/* strtol allows whitespace and other stuff in the beginning
	 * we're not interested */

	if hexadecimal != 0 && !(isxdigit(*(*buf))) || hexadecimal == 0 && !(isdigit(*(*buf))) {
		return types.FAILURE
	}
	code_l = zend.ZEND_STRTOL(*buf, &endptr, lang.Cond(hexadecimal != 0, 16, 10))

	/* we're guaranteed there were valid digits, so *endptr > buf */

	*buf = endptr
	if (*(*buf)) != ';' {
		return types.FAILURE
	}

	/* many more are invalid, but that depends on whether it's HTML
	 * (and which version) or XML. */

	if code_l > int64(0x10ffff) {
		return types.FAILURE
	}
	if code_point != nil {
		*code_point = uint(code_l)
	}
	return types.SUCCESS
}
func ProcessNamedEntityHtml(buf **byte, start **byte, length *int) int {
	*start = *buf

	/* "&" is represented by a 0x26 in all supported encodings. That means
	 * the byte after represents a character or is the leading byte of an
	 * sequence of 8-bit code units. If in the ranges below, it represents
	 * necessarily a alpha character because none of the supported encodings
	 * has an overlap with ASCII in the leading byte (only on the second one) */

	for (*(*buf)) >= 'a' && (*(*buf)) <= 'z' || (*(*buf)) >= 'A' && (*(*buf)) <= 'Z' || (*(*buf)) >= '0' && (*(*buf)) <= '9' {
		*buf++
	}
	if (*(*buf)) != ';' {
		return types.FAILURE
	}

	/* cast to size_t OK as the quantity is always non-negative */

	*length = (*buf) - (*start)
	if (*length) == 0 {
		return types.FAILURE
	}
	return types.SUCCESS
}
func ResolveNamedEntityHtml(start *byte, length int, ht *EntityHt, uni_cp1 *uint, uni_cp2 *uint) int {
	var s *EntityCpMap
	var hash zend.ZendUlong = b.HashStr(b.CastStr(start, length))
	s = ht.GetBuckets()[hash%ht.GetNumElems()]
	for s.GetEntity() != nil {
		if s.GetEntityLen() == length {
			if memcmp(start, s.GetEntity(), length) == 0 {
				*uni_cp1 = s.GetCodepoint1()
				*uni_cp2 = s.GetCodepoint2()
				return types.SUCCESS
			}
		}
		s++
	}
	return types.FAILURE
}
func WriteOctetSequenceAsBytes(charset EntityCharset, code uint) []byte {
	/* code is not necessarily a unicode code point */
	switch charset {
	case CsUtf8:
		return PhpUtf32Utf8Ex(code)
	case Cs88591, CsCp1252, Cs885915, CsKoi8r, CsCp1251, Cs88595, CsCp866, CsMacroman:
		/* single byte stuff */
		return []byte{byte(code)}
	case CsBig5, CsBig5hkscs, CsSjis, CsGb2312:
		/* we don't have complete unicode mappings for these yet in entity_decode,
		 * and we opt to pass through the octet sequences for these in htmlentities
		 * instead of converting to an int and then converting back. */
		return []byte{byte(code)}
	case CsEucjp:
		return []byte{byte(code)}
	default:
		b.Assert(false)
		return nil
	}
	/* code is not necessarily a unicode code point */
}
func WriteOctetSequence(buf *uint8, charset EntityCharset, code uint) int {
	/* code is not necessarily a unicode code point */

	switch charset {
	case CsUtf8:
		return PhpUtf32Utf8(buf, code)
	case Cs88591, CsCp1252, Cs885915, CsKoi8r, CsCp1251, Cs88595, CsCp866, CsMacroman:

		/* single byte stuff */

		*buf = code
		return 1
	case CsBig5, CsBig5hkscs, CsSjis, CsGb2312:

		/* we don't have complete unicode mappings for these yet in entity_decode,
		 * and we opt to pass through the octet sequences for these in htmlentities
		 * instead of converting to an int and then converting back. */

		*buf = code
		return 1
	case CsEucjp:
		*buf = code
		return 1
	default:
		b.Assert(false)
		return 0
	}

	/* code is not necessarily a unicode code point */
}
func TRAVERSE_FOR_ENTITIES_EXPAND_SIZE(oldlen int) int { return oldlen + oldlen/5 + 2 }
func TraverseForEntities(old string, all bool, flags int, inv_map *EntityHt, charset EntityCharset) string {
	var doctype int = flags & ENT_HTML_DOC_TYPE_MASK
	var p *byte
	p = old
	oldLen := len(old)
	var buf strings.Builder
	for idx := 0; idx < oldLen; {
		var code uint
		var code2 uint = 0
		var next *byte = nil

		/* Shift JIS, Big5 and HKSCS use multi-byte encodings where an
		 * ASCII range byte can be part of a multi-byte sequence.
		 * However, they start at 0x40, therefore if we find a 0x26 byte,
		 * we're sure it represents the '&' character. */

		if old[idx] != '&' || idx+3 >= oldLen {
			buf.WriteByte(old[idx])
			idx++
			continue
		}

		/* now p[3] is surely valid and is no terminator */

		if old[idx+1] == '#' {
			next = &p[2]
			if ProcessNumericEntity(&next, &code) == types.FAILURE {
				goto invalid_code
			}

			/* If we're in htmlspecialchars_decode, we're only decoding entities
			 * that represent &, <, >, " and '. Is this one of them? */
			if !all && (code > 63 || Stage3TableBeApos00000[code].GetEntity() == nil) {
				goto invalid_code
			}

			/* are we allowed to decode this entity in this document type?
			 * HTML 5 is the only that has a character that cannot be used in
			 * a numeric entity but is allowed literally (U+000D). The
			 * unoptimized version would be ... || !numeric_entity_is_allowed(code) */

			if UnicodeCpIsAllowed(code, doctype) == 0 || doctype == ENT_HTML_DOC_HTML5 && code == 0xd {
				goto invalid_code
			}

			/* are we allowed to decode this entity in this document type?
			 * HTML 5 is the only that has a character that cannot be used in
			 * a numeric entity but is allowed literally (U+000D). The
			 * unoptimized version would be ... || !numeric_entity_is_allowed(code) */

		} else {
			var start *byte
			var ent_len int
			next = &p[1]
			start = next
			if ProcessNamedEntityHtml(&next, &start, &ent_len) == types.FAILURE {
				goto invalid_code
			}
			if ResolveNamedEntityHtml(start, ent_len, inv_map, &code, &code2) == types.FAILURE {
				if doctype == ENT_HTML_DOC_XHTML && ent_len == 4 && start[0] == 'a' && start[1] == 'p' && start[2] == 'o' && start[3] == 's' {

					/* uses html4 inv_map, which doesn't include apos;. This is a
					 * hack to support it */

					code = uint('\'')

					/* uses html4 inv_map, which doesn't include apos;. This is a
					 * hack to support it */

				} else {
					goto invalid_code
				}
			}
		}
		b.Assert((*next) == ';')
		if code == '\'' && (flags&ENT_HTML_QUOTE_SINGLE) == 0 || code == '"' && (flags&ENT_HTML_QUOTE_DOUBLE) == 0 {
			goto invalid_code
		}

		/* UTF-8 doesn't need mapping (ISO-8859-1 doesn't either, but
		 * the call is needed to ensure the codepoint <= U+00FF)  */
		if charset != CsUtf8 {
			/* replace unicode code point */
			if MapFromUnicode(code, charset, &code) == types.FAILURE || code2 != 0 {
				goto invalid_code
			}
		}
		buf.Write(WriteOctetSequenceAsBytes(charset, code))
		if code2 != 0 {
			buf.Write(WriteOctetSequenceAsBytes(charset, code2))
		}

		/* jump over the valid entity; may go beyond size of buffer; np */
		p = next + 1
		continue
	invalid_code:
		for ; p < next; p++ {
			buf.WriteByte(*p)
		}
	}
	return buf.String()
}
func UnescapeInverseMap(all bool, flags int) *EntityHt {
	var document_type int = flags & ENT_HTML_DOC_TYPE_MASK
	if all {
		switch document_type {
		case ENT_HTML_DOC_HTML401:
			fallthrough
		case ENT_HTML_DOC_XHTML:
			return &EntHtHtml4
		case ENT_HTML_DOC_HTML5:
			return &EntHtHtml5
		default:
			return &EntHtBeApos
		}
	} else {
		switch document_type {
		case ENT_HTML_DOC_HTML401:
			return &EntHtBeNoapos
		default:
			return &EntHtBeApos
		}
	}
}
func DetermineEntityTable(all bool, doctype int) EntityTableOpt {
	var retval EntityTableOpt = MakeEntityTableOpt(nil, nil)
	b.Assert(!(doctype == ENT_HTML_DOC_XML1 && all))
	if all {
		if doctype == ENT_HTML_DOC_HTML5 {
			retval.SetMsTable(EntityMsTableHtml5)
		} else {
			retval.SetMsTable(EntityMsTableHtml4)
		}
	} else {
		if doctype == ENT_HTML_DOC_HTML401 {
			retval.SetTable(Stage3TableBeNoapos00000)
		} else {
			retval.SetTable(Stage3TableBeApos00000)
		}
	}
	return retval
}
func PhpUnescapeHtmlEntities(str string, all bool, flags int, hintCharset string) string {
	if strings.IndexByte(str, '&') < 0 {
		return str
	}

	charset := Cs88591
	if all {
		charset = DetermineCharset(hintCharset)
	}

	var inverse_map *EntityHt
	var new_size int

	/* don't use LIMIT_ALL! */
	new_size = TRAVERSE_FOR_ENTITIES_EXPAND_SIZE(len(str))
	if len(str) > new_size {
		/* overflow, refuse to do anything */
		return str
	}
	inverse_map = UnescapeInverseMap(all, flags)

	/* replace numeric entities */
	return TraverseForEntities(str, all, flags, inverse_map, charset)
}

func PhpEscapeHtmlEntities(old string, all bool, flags int, hintCharset string) string {
	return PhpEscapeHtmlEntitiesEx(old, all, flags, hintCharset, true)
}

func FindEntityForChar(
	k uint,
	charset EntityCharset,
	table *EntityStage1Row,
	entity **uint8,
	entity_len *int,
	old *uint8,
	oldlen int,
	cursor *int,
) {
	var stage1_idx uint = ENT_STAGE1_INDEX(k)
	var c *EntityStage3Row
	if stage1_idx > 0x1d {
		*entity = nil
		*entity_len = 0
		return
	}
	c = &table[stage1_idx][ENT_STAGE2_INDEX(k)][ENT_STAGE3_INDEX(k)]
	if !(c.GetAmbiguous()) {
		*entity = (*uint8)(c.GetEntity())
		*entity_len = c.GetEntityLen()
	} else {

		/* peek at next char */

		var cursor_before int = *cursor
		var status int = types.SUCCESS
		var next_char uint
		if (*cursor) >= oldlen {
			goto no_suitable_2nd
		}
		next_char = GetNextChar(charset, old, oldlen, cursor, &status)
		if status == types.FAILURE {
			goto no_suitable_2nd
		}
		var s *EntityMulticodepointRow
		var e *EntityMulticodepointRow
		s = c.GetMulticodepointTable()[1]
		e = s - 1 + c.GetMulticodepointTable()[0].GetSize()

		/* we could do a binary search but it's not worth it since we have
		 * at most two entries... */

		for ; s <= e; s++ {
			if s.GetSecondCp() == next_char {
				*entity = (*uint8)(s.GetEntity())
				*entity_len = s.GetEntityLen()
				return
			}
		}

		/* we could do a binary search but it's not worth it since we have
		 * at most two entries... */

	no_suitable_2nd:
		*cursor = cursor_before
		*entity = (*uint8)(c.GetMulticodepointTable()[0].GetDefaultEntity())
		*entity_len = c.GetMulticodepointTable()[0].GetDefaultEntityLen()
	}
}
func FindEntityForCharBasic(k uint, table *EntityStage3Row, entity **uint8, entity_len *int) {
	if k >= 64 {
		*entity = nil
		*entity_len = 0
		return
	}
	*entity = (*uint8)(table[k].GetEntity())
	*entity_len = table[k].GetEntityLen()
}
func PhpEscapeHtmlEntitiesEx(old string, all bool, flags int, hintCharset string, doubleEncode bool) string {
	var cursor int
	var buf strings.Builder
	var charset EntityCharset = DetermineCharset(hintCharset)
	var doctype int = flags & ENT_HTML_DOC_TYPE_MASK
	var entity_table EntityTableOpt
	var to_uni_table *EncToUni = nil
	var inv_map *EntityHt = nil
	oldlen := len(old)

	/* only used if flags includes ENT_HTML_IGNORE_ERRORS or ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS */
	var replacement string
	var replacement_len int = 0
	if all {
		if CHARSET_PARTIAL_SUPPORT(charset) {
			core.PhpErrorDocref("", faults.E_NOTICE, "Only basic entities substitution is supported for multi-byte encodings other than UTF-8; functionality is equivalent to htmlspecialchars")
		}
		LIMIT_ALL(all, doctype, charset)
	}
	entity_table = DetermineEntityTable(all, doctype)
	if all && !(CHARSET_UNICODE_COMPAT(charset)) {
		to_uni_table = EncToUniIndex[charset]
	}
	if !doubleEncode {

		/* first arg is 1 because we want to identify valid named entities
		 * even if we are only encoding the basic ones */

		inv_map = UnescapeInverseMap(true, flags)

		/* first arg is 1 because we want to identify valid named entities
		 * even if we are only encoding the basic ones */

	}
	if (flags & (ENT_HTML_SUBSTITUTE_ERRORS | ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS)) != 0 {
		if charset == CsUtf8 {
			replacement = "\xEF\xBF\xBD"
		} else {
			replacement = "&#xFFFD;"
		}
		replacement_len = len(replacement)
	}

	/* initial estimate */
	cursor = 0
	for cursor < oldlen {
		var mbsequence *uint8 = nil
		var mbseqlen int = 0
		var cursor_before int = cursor
		var status int = types.SUCCESS
		var this_char uint = GetNextChar(charset, old, oldlen, &cursor, &status)

		/* guarantee we have at least 40 bytes to write.
		 * In HTML5, entities may take up to 33 bytes */
		if status == types.FAILURE {
			/* invalid MB sequence */
			if (flags & ENT_HTML_IGNORE_ERRORS) != 0 {
				continue
			} else if (flags & ENT_HTML_SUBSTITUTE_ERRORS) != 0 {
				buf.WriteString(replacement)
				continue
			} else {
				return ""
			}
		} else {
			mbsequence = &old[cursor_before]
			mbseqlen = cursor - cursor_before
		}
		if this_char != '&' {
			var rep *uint8 = nil
			var rep_len int = 0
			if this_char == '\'' && (flags&ENT_HTML_QUOTE_SINGLE) == 0 || this_char == '"' && (flags&ENT_HTML_QUOTE_DOUBLE) == 0 {
				goto pass_char_through
			}
			if all {
				if to_uni_table != nil {
					/* !CHARSET_UNICODE_COMPAT therefore not UTF-8; since UTF-8
					 * is the only multibyte encoding with !CHARSET_PARTIAL_SUPPORT,
					 * we're using a single byte encoding */
					MapToUnicode(this_char, to_uni_table, &this_char)
					if this_char == 0xffff {
						goto pass_char_through
					}
				}

				/* the cursor may advance */
				FindEntityForChar(this_char, charset, entity_table.GetMsTable(), &rep, &rep_len, old, oldlen, &cursor)
			} else {
				FindEntityForCharBasic(this_char, entity_table.GetTable(), &rep, &rep_len)
			}
			if rep != nil {
				buf.WriteByte('&')
				buf.WriteString(rep)
				buf.WriteByte(';')
			} else {

				/* we did not find an entity for this char.
				 * check for its validity, if its valid pass it unchanged */
				if (flags & ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS) != 0 {
					if CHARSET_UNICODE_COMPAT(charset) {
						if UnicodeCpIsAllowed(this_char, doctype) == 0 {
							mbsequence = replacement
							mbseqlen = replacement_len
						}
					} else if to_uni_table != nil {
						if all == 0 {
							MapToUnicode(this_char, to_uni_table, &this_char)
						}
						if UnicodeCpIsAllowed(this_char, doctype) == 0 {
							mbsequence = replacement
							mbseqlen = replacement_len
						}
					} else {

						/* not a unicode code point, unless, coincidentally, it's in
						 * the 0x20..0x7D range (except 0x5C in sjis). We know nothing
						 * about other code points, because we have no tables. Since
						 * Unicode code points in that range are not disallowed in any
						 * document type, we could do nothing. However, conversion
						 * tables frequently map 0x00-0x1F to the respective C0 code
						 * points. Let's play it safe and admit that's the case */

						if this_char <= 0x7d && UnicodeCpIsAllowed(this_char, doctype) == 0 {
							mbsequence = replacement
							mbseqlen = replacement_len
						}
					}
				}
			pass_char_through:
				if mbseqlen > 1 {
					buf.WriteString(b.CastStr(mbsequence, mbseqlen))
				} else {
					buf.WriteByte(mbsequence[0])
				}
			}
		} else {
			if doubleEncode != 0 {
			encode_amp:
				buf.WriteString("&amp;")
			} else {

				/* check if entity is valid */

				var ent_len int

				/* peek at next char */

				if old[cursor] == '#' {
					var code_point uint
					var valid int
					var pos *byte = (*byte)(&old[cursor+1])
					valid = ProcessNumericEntity((**byte)(&pos), &code_point)
					if valid == types.FAILURE {
						goto encode_amp
					}
					if (flags & ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS) != 0 {
						if NumericEntityIsAllowed(code_point, doctype) == 0 {
							goto encode_amp
						}
					}
					ent_len = pos - (*byte)(&old[cursor])
				} else {

					/* check for vality of named entity */

					var start *byte = (*byte)(&old[cursor])
					var next *byte = start
					var dummy1 uint
					var dummy2 uint
					if ProcessNamedEntityHtml(&next, &start, &ent_len) == types.FAILURE {
						goto encode_amp
					}
					if ResolveNamedEntityHtml(start, ent_len, inv_map, &dummy1, &dummy2) == types.FAILURE {
						if !(doctype == ENT_HTML_DOC_XHTML && ent_len == 4 && start[0] == 'a' && start[1] == 'p' && start[2] == 'o' && start[3] == 's') {
							/* uses html4 inv_map, which doesn't include apos;. This is a hack to support it */
							goto encode_amp
						}
					}
				}

				/* checks passed; copy entity to result */
				buf.WriteByte('&')
				buf.WriteString(old[cursor : cursor+ent_len])
				buf.WriteByte(';')

				cursor += ent_len + 1
			}
		}
	}
	return buf.String()
}
func PhpHtmlEntities(executeData *zend.ZendExecuteData, return_value *types.Zval, all int) {
	var str *types.String
	var hint_charset *types.String
	var flags zend.ZendLong = ENT_COMPAT
	var double_encode bool = 1
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 4, 0)
			str = fp.ParseStr()
			fp.StartOptional()
			flags = fp.ParseLong()
			hint_charset = fp.ParseStrEx(true, false)
			double_encode = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	var hintCharset string
	if hint_charset == nil {
		hintCharset = GetDefaultCharset()
	} else {
		hintCharset = hint_charset.GetStr()
	}
	replaced := PhpEscapeHtmlEntitiesEx(str.GetStr(), all != 0, flags, hintCharset, double_encode)
	return_value.SetString(replaced)
}
func RegisterHtmlConstants(moduleNumber int) {
	zend.RegisterLongConstant("HTML_SPECIALCHARS", HTML_SPECIALCHARS, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("HTML_ENTITIES", HTML_ENTITIES, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_COMPAT", ENT_COMPAT, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_QUOTES", ENT_QUOTES, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_NOQUOTES", ENT_NOQUOTES, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_IGNORE", ENT_IGNORE, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_SUBSTITUTE", ENT_SUBSTITUTE, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_DISALLOWED", ENT_DISALLOWED, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_HTML401", ENT_HTML401, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_XML1", ENT_XML1, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_XHTML", ENT_XHTML, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
	zend.RegisterLongConstant("ENT_HTML5", ENT_HTML5, zend.CONST_PERSISTENT|zend.CONST_CS, moduleNumber)
}
func ZifHtmlspecialchars(executeData zpp.Ex, return_value zpp.Ret, string *types.Zval, _ zpp.Opt, quoteStyle *types.Zval, encoding *types.Zval, doubleEncode *types.Zval) {
	PhpHtmlEntities(executeData, return_value, 0)
}
func ZifHtmlspecialcharsDecode(str string, _ zpp.Opt, quoteStyle_ *int) string {
	var quoteStyle = b.Option(quoteStyle_, ENT_COMPAT)
	return PhpUnescapeHtmlEntities(str, false, quoteStyle, "")
}
func ZifHtmlEntityDecode(str string, _ zpp.Opt, quoteStyle_ *int, encoding *string) string {
	quoteStyle := b.Option(quoteStyle_, ENT_COMPAT)
	var charset string
	if encoding != nil {
		charset = *encoding
	} else {
		charset = GetDefaultCharset()
	}
	return PhpUnescapeHtmlEntities(str, true, quoteStyle, charset)
}
func ZifHtmlentities(executeData zpp.Ex, return_value zpp.Ret, string *types.Zval, _ zpp.Opt, quoteStyle *types.Zval, encoding *types.Zval, doubleEncode *types.Zval) {
	PhpHtmlEntities(executeData, return_value, 1)
}
func WriteS3rowData(r *EntityStage3Row, orig_cp uint, charset EntityCharset, arr *types.Zval) {
	var key []byte = ""
	var entity []byte = []byte{'&'}
	var written_k1 int
	written_k1 = WriteOctetSequence((*uint8)(key), charset, orig_cp)
	if !(r.GetAmbiguous()) {
		var l int = r.GetEntityLen()
		memcpy(&entity[1], r.GetEntity(), l)
		entity[l+1] = ';'
		zend.AddAssocStringlEx(arr, b.CastStr(key, written_k1), b.CastStr(entity, l+2))
	} else {
		var i uint
		var num_entries uint
		var mcpr *EntityMulticodepointRow = r.GetMulticodepointTable()
		if mcpr[0].GetDefaultEntity() != nil {
			var l int = mcpr[0].GetDefaultEntityLen()
			memcpy(&entity[1], mcpr[0].GetDefaultEntity(), l)
			entity[l+1] = ';'
			zend.AddAssocStringlEx(arr, b.CastStr(key, written_k1), b.CastStr(entity, l+2))
		}
		num_entries = mcpr[0].GetSize()
		for i = 1; i <= num_entries; i++ {
			var l int
			var written_k2 int
			var uni_cp uint
			var spe_cp uint
			uni_cp = mcpr[i].GetSecondCp()
			l = mcpr[i].GetEntityLen()
			if !(CHARSET_UNICODE_COMPAT(charset)) {
				if MapFromUnicode(uni_cp, charset, &spe_cp) == types.FAILURE {
					continue
				}
			} else {
				spe_cp = uni_cp
			}
			written_k2 = WriteOctetSequence((*uint8)(&key[written_k1]), charset, spe_cp)
			memcpy(&entity[1], mcpr[i].GetEntity(), l)
			entity[l+1] = ';'
			zend.AddAssocStringlEx(arr, b.CastStr(key, written_k1+written_k2), b.CastStr(entity, l+2))
		}
	}
}
func ZifGetHtmlTranslationTable(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, table *types.Zval, quoteStyle *types.Zval, encoding *types.Zval) {
	var all zend.ZendLong = HTML_SPECIALCHARS
	var flags zend.ZendLong = ENT_COMPAT
	var doctype int
	var entity_table EntityTableOpt
	var to_uni_table *EncToUni = nil
	var charset_hint *byte = nil
	var charset_hint_len int
	var charset EntityCharset

	/* in this function we have to jump through some loops because we're
	 * getting the translated table from data structures that are optimized for
	 * random access, not traversal */

	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 3, 0)
			fp.StartOptional()
			all = fp.ParseLong()
			flags = fp.ParseLong()
			charset_hint, charset_hint_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	charset = DetermineCharset(charset_hint)
	doctype = flags & ENT_HTML_DOC_TYPE_MASK
	LIMIT_ALL(all, doctype, charset)
	zend.ArrayInit(return_value)
	entity_table = DetermineEntityTable(all != 0, doctype)
	if all != 0 && !(CHARSET_UNICODE_COMPAT(charset)) {
		to_uni_table = EncToUniIndex[charset]
	}
	if all != 0 {
		var ms_table *EntityStage1Row = entity_table.GetMsTable()
		if CHARSET_UNICODE_COMPAT(charset) {
			var i uint
			var j uint
			var k uint
			var max_i uint
			var max_j uint
			var max_k uint

			/* no mapping to unicode required */

			if CHARSET_SINGLE_BYTE(charset) {
				max_i = 1
				max_j = 4
				max_k = 64
			} else {
				max_i = 0x1e
				max_j = 64
				max_k = 64
			}
			for i = 0; i < max_i; i++ {
				if ms_table[i] == EmptyStage2Table {
					continue
				}
				for j = 0; j < max_j; j++ {
					if ms_table[i][j] == EmptyStage3Table {
						continue
					}
					for k = 0; k < max_k; k++ {
						var r *EntityStage3Row = &ms_table[i][j][k]
						var code uint
						if r.GetEntity() == nil {
							continue
						}
						code = ENT_CODE_POINT_FROM_STAGES(i, j, k)
						if code == '\'' && (flags&ENT_HTML_QUOTE_SINGLE) == 0 || code == '"' && (flags&ENT_HTML_QUOTE_DOUBLE) == 0 {
							continue
						}
						WriteS3rowData(r, code, charset, return_value)
					}
				}
			}
		} else {

			/* we have to iterate through the set of code points for this
			 * encoding and map them to unicode code points */

			var i uint
			for i = 0; i <= 0xff; i++ {
				var r *EntityStage3Row
				var uni_cp uint

				/* can be done before mapping, they're invariant */

				if i == '\'' && (flags&ENT_HTML_QUOTE_SINGLE) == 0 || i == '"' && (flags&ENT_HTML_QUOTE_DOUBLE) == 0 {
					continue
				}
				MapToUnicode(i, to_uni_table, &uni_cp)
				r = &ms_table[ENT_STAGE1_INDEX(uni_cp)][ENT_STAGE2_INDEX(uni_cp)][ENT_STAGE3_INDEX(uni_cp)]
				if r.GetEntity() == nil {
					continue
				}
				WriteS3rowData(r, i, charset, return_value)
			}
		}
	} else {

		/* we could use sizeof(stage3_table_be_apos_00000) as well */

		var j uint
		var numelems uint = b.SizeOf("stage3_table_be_noapos_00000") / b.SizeOf("* stage3_table_be_noapos_00000")
		for j = 0; j < numelems; j++ {
			var r *EntityStage3Row = entity_table.GetTable()[j]
			if r.GetEntity() == nil {
				continue
			}
			if j == '\'' && (flags&ENT_HTML_QUOTE_SINGLE) == 0 || j == '"' && (flags&ENT_HTML_QUOTE_DOUBLE) == 0 {
				continue
			}

			/* charset is indifferent, used cs_8859_1 for efficiency */

			WriteS3rowData(r, j, Cs88591, return_value)

			/* charset is indifferent, used cs_8859_1 for efficiency */

		}
	}
}

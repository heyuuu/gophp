// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/html.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// #define HTML_H

// #define ENT_HTML_QUOTE_NONE       0

// #define ENT_HTML_QUOTE_SINGLE       1

// #define ENT_HTML_QUOTE_DOUBLE       2

// #define ENT_HTML_IGNORE_ERRORS       4

// #define ENT_HTML_SUBSTITUTE_ERRORS       8

// #define ENT_HTML_DOC_TYPE_MASK       ( 16 | 32 )

// #define ENT_HTML_DOC_HTML401       0

// #define ENT_HTML_DOC_XML1       16

// #define ENT_HTML_DOC_XHTML       32

// #define ENT_HTML_DOC_HTML5       ( 16 | 32 )

/* reserve bit 6 */

// #define ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS       128

// #define ENT_COMPAT       ENT_HTML_QUOTE_DOUBLE

// #define ENT_QUOTES       ( ENT_HTML_QUOTE_DOUBLE | ENT_HTML_QUOTE_SINGLE )

// #define ENT_NOQUOTES       ENT_HTML_QUOTE_NONE

// #define ENT_IGNORE       ENT_HTML_IGNORE_ERRORS

// #define ENT_SUBSTITUTE       ENT_HTML_SUBSTITUTE_ERRORS

// #define ENT_HTML401       0

// #define ENT_XML1       16

// #define ENT_XHTML       32

// #define ENT_HTML5       ( 16 | 32 )

// #define ENT_DISALLOWED       128

// Source: <ext/standard/html.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Jaakko Hyvätti <jaakko.hyvatti@iki.fi>                      |
   |          Wez Furlong    <wez@thebrainroom.com>                       |
   |          Gustavo Lopes  <cataphract@php.net>                         |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < php_config . h >

// # include "php_standard.h"

// # include "php_string.h"

// # include "SAPI.h"

// # include < locale . h >

// # include < langinfo . h >

// # include < zend_hash . h >

// # include "html_tables.h"

/* Macro for disabling flag of translation of non-basic entities where this isn't supported.
 * Not appropriate for html_entity_decode/htmlspecialchars_decode */

// #define LIMIT_ALL(all,doctype,charset) do { ( all ) = ( all ) && ! CHARSET_PARTIAL_SUPPORT ( ( charset ) ) && ( ( doctype ) != ENT_HTML_DOC_XML1 ) ; } while ( 0 )

// #define MB_FAILURE(pos,advance) do { * cursor = pos + ( advance ) ; * status = FAILURE ; return 0 ; } while ( 0 )

// #define CHECK_LEN(pos,chars_need) ( ( str_len - ( pos ) ) >= ( chars_need ) )

/* valid as single byte character or leading byte */

// #define utf8_lead(c) ( ( c ) < 0x80 || ( ( c ) >= 0xC2 && ( c ) <= 0xF4 ) )

/* whether it's actually valid depends on other stuff;
 * this macro cannot check for non-shortest forms, surrogates or
 * code points above 0x10FFFF */

// #define utf8_trail(c) ( ( c ) >= 0x80 && ( c ) <= 0xBF )

// #define gb2312_lead(c) ( ( c ) != 0x8E && ( c ) != 0x8F && ( c ) != 0xA0 && ( c ) != 0xFF )

// #define gb2312_trail(c) ( ( c ) >= 0xA1 && ( c ) <= 0xFE )

// #define sjis_lead(c) ( ( c ) != 0x80 && ( c ) != 0xA0 && ( c ) < 0xFD )

// #define sjis_trail(c) ( ( c ) >= 0x40 && ( c ) != 0x7F && ( c ) < 0xFD )

/* {{{ get_default_charset
 */

func GetDefaultCharset() *byte {
	if core.CoreGlobals.internal_encoding != nil && core.CoreGlobals.internal_encoding[0] {
		return core.CoreGlobals.internal_encoding
	} else if core.sapi_globals.default_charset != nil && core.sapi_globals.default_charset[0] {
		return core.sapi_globals.default_charset
	}
	return nil
}

/* }}} */

func GetNextChar(charset EntityCharset, str *uint8, str_len int, cursor *int, status *int) uint {
	var pos int = *cursor
	var this_char uint = 0
	*status = zend.SUCCESS
	assert(pos <= str_len)
	if str_len-pos < 1 {
		*cursor = pos + 1
		*status = zend.FAILURE
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
			*status = zend.FAILURE
			return 0
		} else if c < 0xe0 {
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			if !(str[pos+1] >= 0x80 && str[pos+1] <= 0xbf) {
				*cursor = pos + g.Cond(str[pos+1] < 0x80 || str[pos+1] >= 0xc2 && str[pos+1] <= 0xf4, 1, 2)
				*status = zend.FAILURE
				return 0
			}
			this_char = (c&0x1f)<<6 | str[pos+1]&0x3f
			if this_char < 0x80 {
				*cursor = pos + 2
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else if c < 0xf0 {
			var avail int = str_len - pos
			if avail < 3 || !(str[pos+1] >= 0x80 && str[pos+1] <= 0xbf) || !(str[pos+2] >= 0x80 && str[pos+2] <= 0xbf) {
				if avail < 2 || (str[pos+1] < 0x80 || str[pos+1] >= 0xc2 && str[pos+1] <= 0xf4) {
					*cursor = pos + 1
					*status = zend.FAILURE
					return 0
				} else if avail < 3 || (str[pos+2] < 0x80 || str[pos+2] >= 0xc2 && str[pos+2] <= 0xf4) {
					*cursor = pos + 2
					*status = zend.FAILURE
					return 0
				} else {
					*cursor = pos + 3
					*status = zend.FAILURE
					return 0
				}
			}
			this_char = (c&0xf)<<12 | (str[pos+1]&0x3f)<<6 | str[pos+2]&0x3f
			if this_char < 0x800 {
				*cursor = pos + 3
				*status = zend.FAILURE
				return 0
			} else if this_char >= 0xd800 && this_char <= 0xdfff {
				*cursor = pos + 3
				*status = zend.FAILURE
				return 0
			}
			pos += 3
		} else if c < 0xf5 {
			var avail int = str_len - pos
			if avail < 4 || !(str[pos+1] >= 0x80 && str[pos+1] <= 0xbf) || !(str[pos+2] >= 0x80 && str[pos+2] <= 0xbf) || !(str[pos+3] >= 0x80 && str[pos+3] <= 0xbf) {
				if avail < 2 || (str[pos+1] < 0x80 || str[pos+1] >= 0xc2 && str[pos+1] <= 0xf4) {
					*cursor = pos + 1
					*status = zend.FAILURE
					return 0
				} else if avail < 3 || (str[pos+2] < 0x80 || str[pos+2] >= 0xc2 && str[pos+2] <= 0xf4) {
					*cursor = pos + 2
					*status = zend.FAILURE
					return 0
				} else if avail < 4 || (str[pos+3] < 0x80 || str[pos+3] >= 0xc2 && str[pos+3] <= 0xf4) {
					*cursor = pos + 3
					*status = zend.FAILURE
					return 0
				} else {
					*cursor = pos + 4
					*status = zend.FAILURE
					return 0
				}
			}
			this_char = (c&0x7)<<18 | (str[pos+1]&0x3f)<<12 | (str[pos+2]&0x3f)<<6 | str[pos+3]&0x3f
			if this_char < 0x10000 || this_char > 0x10ffff {
				*cursor = pos + 4
				*status = zend.FAILURE
				return 0
			}
			pos += 4
		} else {
			*cursor = pos + 1
			*status = zend.FAILURE
			return 0
		}
		break
	case CsBig5:

		/* reference http://demo.icu-project.org/icu-bin/convexp?conv=big5 */

		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0xfe {
			var next uint8
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
				this_char = c<<8 | next
			} else {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else {
			this_char = c
			pos += 1
		}
		break
	case CsBig5hkscs:
		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0xfe {
			var next uint8
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
				this_char = c<<8 | next
			} else if next != 0x80 && next != 0xff {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else {
			this_char = c
			pos += 1
		}
		break
	case CsGb2312:
		var c uint8 = str[pos]
		if c >= 0xa1 && c <= 0xfe {
			var next uint8
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0xa1 && next <= 0xfe {
				this_char = c<<8 | next
			} else if next != 0x8e && next != 0x8f && next != 0xa0 && next != 0xff {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else if c != 0x8e && c != 0x8f && c != 0xa0 && c != 0xff {
			this_char = c
			pos += 1
		} else {
			*cursor = pos + 1
			*status = zend.FAILURE
			return 0
		}
		break
	case CsSjis:
		var c uint8 = str[pos]
		if c >= 0x81 && c <= 0x9f || c >= 0xe0 && c <= 0xfc {
			var next uint8
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0x40 && next != 0x7f && next < 0xfd {
				this_char = c<<8 | next
			} else if next != 0x80 && next != 0xa0 && next < 0xfd {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			} else {
				*cursor = pos + 2
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else if c < 0x80 || c >= 0xa1 && c <= 0xdf {
			this_char = c
			pos += 1
		} else {
			*cursor = pos + 1
			*status = zend.FAILURE
			return 0
		}
		break
	case CsEucjp:
		var c uint8 = str[pos]
		if c >= 0xa1 && c <= 0xfe {
			var next unsigned
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0xa1 && next <= 0xfe {

				/* this a jis kanji char */

				this_char = c<<8 | next

				/* this a jis kanji char */

			} else {
				*cursor = pos + g.Cond(next != 0xa0 && next != 0xff, 1, 2)
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else if c == 0x8e {
			var next unsigned
			if str_len-pos < 2 {
				*cursor = pos + 1
				*status = zend.FAILURE
				return 0
			}
			next = str[pos+1]
			if next >= 0xa1 && next <= 0xdf {

				/* JIS X 0201 kana */

				this_char = c<<8 | next

				/* JIS X 0201 kana */

			} else {
				*cursor = pos + g.Cond(next != 0xa0 && next != 0xff, 1, 2)
				*status = zend.FAILURE
				return 0
			}
			pos += 2
		} else if c == 0x8f {
			var avail int = str_len - pos
			if avail < 3 || !(str[pos+1] >= 0xa1 && str[pos+1] <= 0xfe) || !(str[pos+2] >= 0xa1 && str[pos+2] <= 0xfe) {
				if avail < 2 || str[pos+1] != 0xa0 && str[pos+1] != 0xff {
					*cursor = pos + 1
					*status = zend.FAILURE
					return 0
				} else if avail < 3 || str[pos+2] != 0xa0 && str[pos+2] != 0xff {
					*cursor = pos + 2
					*status = zend.FAILURE
					return 0
				} else {
					*cursor = pos + 3
					*status = zend.FAILURE
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
			*status = zend.FAILURE
			return 0
		}
		break
	default:

		/* single-byte charsets */

		this_char = str[g.PostInc(&pos)]
		break
	}
	*cursor = pos
	return this_char
}

/* }}} */

func PhpNextUtf8Char(str *uint8, str_len int, cursor *int, status *int) uint {
	return GetNextChar(CsUtf8, str, str_len, cursor, status)
}

/* }}} */

func DetermineCharset(charset_hint *byte) EntityCharset {
	var i int
	var charset EntityCharset = CsUtf8
	var len_ int = 0
	var zenc *zend.ZendEncoding

	/* Default is now UTF-8 */

	if charset_hint == nil {
		return CsUtf8
	}
	if g.Assign(&len_, strlen(charset_hint)) != 0 {
		goto det_charset
	}
	zenc = zend.ZendMultibyteGetInternalEncoding()
	if zenc != nil {
		charset_hint = (*byte)(zend.ZendMultibyteGetEncodingName(zenc))
		if charset_hint != nil && g.Assign(&len_, strlen(charset_hint)) != 0 {
			if len_ == 4 && (!(memcmp("pass", charset_hint, 4)) || !(memcmp("auto", charset_hint, 4))) {
				charset_hint = nil
				len_ = 0
			} else {
				goto det_charset
			}
		}
	}
	charset_hint = core.sapi_globals.default_charset
	if charset_hint != nil && g.Assign(&len_, strlen(charset_hint)) != 0 {
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
		var found int = 0

		/* now walk the charset map and look for the codeset */

		for i = 0; i < g.SizeOf("charset_map")/g.SizeOf("charset_map [ 0 ]"); i++ {
			if len_ == CharsetMap[i].codeset_len && zend.ZendBinaryStrcasecmp(charset_hint, len_, CharsetMap[i].codeset, len_) == 0 {
				charset = CharsetMap[i].charset
				found = 1
				break
			}
		}
		if found == 0 {
			core.PhpErrorDocref(nil, 1<<1, "charset `%s' not supported, assuming utf-8", charset_hint)
		}
	}
	return charset
}

/* }}} */

func PhpUtf32Utf8(buf *uint8, k unsigned) int {
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

	/* UTF-8 has been restricted to max 4 bytes since RFC 3629 */
}

/* }}} */

func UnimapBsearch(table *UniToEnc, code_key_a unsigned, num int) uint8 {
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

/* }}} */

func MapFromUnicode(code unsigned, charset EntityCharset, res *unsigned) int {
	var found uint8
	var table *UniToEnc
	var table_size int
	switch charset {
	case Cs88591:

		/* identity mapping of code points to unicode */

		if code > 0xff {
			return zend.FAILURE
		}
		*res = code
		break
	case Cs88595:
		if code <= 0xa0 || code == 0xad {
			*res = code
		} else if code == 0x2116 {
			*res = 0xf0
		} else if code == 0xa7 {
			*res = 0xfd
		} else if code >= 0x401 && code <= 0x44f {
			if code == 0x40d || code == 0x450 || code == 0x45d {
				return zend.FAILURE
			}
			*res = code - 0x360
		} else {
			return zend.FAILURE
		}
		break
	case Cs885915:
		if code < 0xa4 || code > 0xbe && code <= 0xff {
			*res = code
		} else {
			found = UnimapBsearch(UnimapIso885915, code, g.SizeOf("unimap_iso885915")/g.SizeOf("* unimap_iso885915"))
			if found != 0 {
				*res = found
			} else {
				return zend.FAILURE
			}
		}
		break
	case CsCp1252:
		if code <= 0x7f || code >= 0xa0 && code <= 0xff {
			*res = code
		} else {
			found = UnimapBsearch(UnimapWin1252, code, g.SizeOf("unimap_win1252")/g.SizeOf("* unimap_win1252"))
			if found != 0 {
				*res = found
			} else {
				return zend.FAILURE
			}
		}
		break
	case CsMacroman:
		if code == 0x7f {
			return zend.FAILURE
		}
		table = UnimapMacroman
		table_size = g.SizeOf("unimap_macroman") / g.SizeOf("* unimap_macroman")
		goto table_over_7F
	case CsCp1251:
		table = UnimapWin1251
		table_size = g.SizeOf("unimap_win1251") / g.SizeOf("* unimap_win1251")
		goto table_over_7F
	case CsKoi8r:
		table = UnimapKoi8r
		table_size = g.SizeOf("unimap_koi8r") / g.SizeOf("* unimap_koi8r")
		goto table_over_7F
	case CsCp866:
		table = UnimapCp866
		table_size = g.SizeOf("unimap_cp866") / g.SizeOf("* unimap_cp866")
	table_over_7F:
		if code <= 0x7f {
			*res = code
		} else {
			found = UnimapBsearch(table, code, table_size)
			if found != 0 {
				*res = found
			} else {
				return zend.FAILURE
			}
		}
		break
	case CsSjis:

	case CsEucjp:

		/* we interpret 0x5C as the Yen symbol. This is not universal.
		 * See <http://www.w3.org/Submission/japanese-xml/#ambiguity_of_yen> */

		if code >= 0x20 && code <= 0x7d {
			if code == 0x5c {
				return zend.FAILURE
			}
			*res = code
		} else {
			return zend.FAILURE
		}
		break
	case CsBig5:

	case CsBig5hkscs:

	case CsGb2312:
		if code >= 0x20 && code <= 0x7d {
			*res = code
		} else {
			return zend.FAILURE
		}
		break
	default:
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func MapToUnicode(code unsigned, table *EncToUni, res *unsigned) {
	/* only single byte encodings are currently supported; assumed code <= 0xFF */

	*res = table.GetInner()[(code&0xc0)>>6].GetUniCp()[code&0x3f]

	/* only single byte encodings are currently supported; assumed code <= 0xFF */
}

/* }}} */

func UnicodeCpIsAllowed(uni_cp unsigned, document_type int) int {
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
	case 0:
		return uni_cp >= 0x20 && uni_cp <= 0x7e || (uni_cp == 0xa || uni_cp == 0x9 || uni_cp == 0xd) || uni_cp >= 0xa0 && uni_cp <= 0xd7ff || uni_cp >= 0xe000 && uni_cp <= 0x10ffff
	case 16 | 32:
		return uni_cp >= 0x20 && uni_cp <= 0x7e || uni_cp >= 0x9 && uni_cp <= 0xd && uni_cp != 0xb || uni_cp >= 0xa0 && uni_cp <= 0xd7ff || uni_cp >= 0xe000 && uni_cp <= 0x10ffff && (uni_cp&0xffff) < 0xfffe && (uni_cp < 0xfdd0 || uni_cp > 0xfdef)
	case 32:

	case 16:
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

/* }}} */

func NumericEntityIsAllowed(uni_cp unsigned, document_type int) int {
	/* less restrictive than unicode_cp_is_allowed */

	switch document_type {
	case 0:

		/* all non-SGML characters (those marked with UNUSED in DESCSET) should be
		 * representable with numeric entities */

		return uni_cp <= 0x10ffff
	case 16 | 32:

		/* 8.1.4. The numeric character reference forms described above are allowed to
		 * reference any Unicode code point other than U+0000, U+000D, permanently
		 * undefined Unicode characters (noncharacters), and control characters other
		 * than space characters (U+0009, U+000A, U+000C and U+000D) */

		return uni_cp >= 0x20 && uni_cp <= 0x7e || uni_cp >= 0x9 && uni_cp <= 0xc && uni_cp != 0xb || uni_cp >= 0xa0 && uni_cp <= 0x10ffff && (uni_cp&0xffff) < 0xfffe && (uni_cp < 0xfdd0 || uni_cp > 0xfdef)
	case 32:

	case 16:

		/* OTOH, XML 1.0 requires "character references to match the production for Char
		 * See <http://www.w3.org/TR/REC-xml/#NT-CharRef> */

		return UnicodeCpIsAllowed(uni_cp, document_type)
	default:
		return 1
	}

	/* less restrictive than unicode_cp_is_allowed */
}

/* }}} */

func ProcessNumericEntity(buf **byte, code_point *unsigned) int {
	var code_l zend.ZendLong
	var hexadecimal int = (*(*buf)) == 'x' || (*(*buf)) == 'X'
	var endptr *byte
	if hexadecimal != 0 && (*(*buf)) != '0' {
		*buf++
	}

	/* strtol allows whitespace and other stuff in the beginning
	 * we're not interested */

	if hexadecimal != 0 && !(isxdigit(*(*buf))) || hexadecimal == 0 && !(isdigit(*(*buf))) {
		return zend.FAILURE
	}
	code_l = strtoll(*buf, &endptr, g.Cond(hexadecimal != 0, 16, 10))

	/* we're guaranteed there were valid digits, so *endptr > buf */

	*buf = endptr
	if (*(*buf)) != ';' {
		return zend.FAILURE
	}

	/* many more are invalid, but that depends on whether it's HTML
	 * (and which version) or XML. */

	if code_l > 0x10ffff {
		return zend.FAILURE
	}
	if code_point != nil {
		*code_point = unsigned(code_l)
	}
	return zend.SUCCESS
}

/* }}} */

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
		return zend.FAILURE
	}

	/* cast to size_t OK as the quantity is always non-negative */

	*length = (*buf) - (*start)
	if (*length) == 0 {
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func ResolveNamedEntityHtml(start *byte, length int, ht *EntityHt, uni_cp1 *unsigned, uni_cp2 *unsigned) int {
	var s *EntityCpMap
	var hash zend.ZendUlong = zend.ZendInlineHashFunc(start, length)
	s = ht.GetBuckets()[hash%ht.GetNumElems()]
	for s.GetEntity() != nil {
		if s.GetEntityLen() == length {
			if memcmp(start, s.GetEntity(), length) == 0 {
				*uni_cp1 = s.GetCodepoint1()
				*uni_cp2 = s.GetCodepoint2()
				return zend.SUCCESS
			}
		}
		s++
	}
	return zend.FAILURE
}

/* }}} */

func WriteOctetSequence(buf *uint8, charset EntityCharset, code unsigned) int {
	/* code is not necessarily a unicode code point */

	switch charset {
	case CsUtf8:
		return PhpUtf32Utf8(buf, code)
	case Cs88591:

	case CsCp1252:

	case Cs885915:

	case CsKoi8r:

	case CsCp1251:

	case Cs88595:

	case CsCp866:

	case CsMacroman:

		/* single byte stuff */

		*buf = code
		return 1
	case CsBig5:

	case CsBig5hkscs:

	case CsSjis:

	case CsGb2312:

		/* we don't have complete unicode mappings for these yet in entity_decode,
		 * and we opt to pass through the octet sequences for these in htmlentities
		 * instead of converting to an int and then converting back. */

		*buf = code
		return 1
	case CsEucjp:
		*buf = code
		return 1
	default:
		assert(false)
		return 0
	}

	/* code is not necessarily a unicode code point */
}

/* {{{ traverse_for_entities
 * Auxiliary function to php_unescape_html_entities().
 * - The argument "all" determines if all numeric entities are decode or only those
 *   that correspond to quotes (depending on quote_style).
 */

// #define TRAVERSE_FOR_ENTITIES_EXPAND_SIZE(oldlen) ( ( oldlen ) + ( oldlen ) / 5 + 2 )

func TraverseForEntities(old *byte, oldlen int, ret *zend.ZendString, all int, flags int, inv_map *EntityHt, charset EntityCharset) {
	var p *byte
	var lim *byte
	var q *byte
	var doctype int = flags & (16 | 32)
	lim = old + oldlen
	assert((*lim) == '0')
	p = old
	q = ret.val
	for p < lim {
		var code unsigned
		var code2 unsigned = 0
		var next *byte = nil

		/* Shift JIS, Big5 and HKSCS use multi-byte encodings where an
		 * ASCII range byte can be part of a multi-byte sequence.
		 * However, they start at 0x40, therefore if we find a 0x26 byte,
		 * we're sure it represents the '&' character. */

		if p[0] != '&' || p+3 >= lim {
			*(g.PostInc(&q)) = *(g.PostInc(&p))
			continue
		}

		/* now p[3] is surely valid and is no terminator */

		if p[1] == '#' {
			next = &p[2]
			if ProcessNumericEntity(&next, &code) == zend.FAILURE {
				goto invalid_code
			}

			/* If we're in htmlspecialchars_decode, we're only decoding entities
			 * that represent &, <, >, " and '. Is this one of them? */

			if all == 0 && (code > 63 || Stage3TableBeApos00000[code].GetEntity() == nil) {
				goto invalid_code
			}

			/* are we allowed to decode this entity in this document type?
			 * HTML 5 is the only that has a character that cannot be used in
			 * a numeric entity but is allowed literally (U+000D). The
			 * unoptimized version would be ... || !numeric_entity_is_allowed(code) */

			if UnicodeCpIsAllowed(code, doctype) == 0 || doctype == (16|32) && code == 0xd {
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
			if ProcessNamedEntityHtml(&next, &start, &ent_len) == zend.FAILURE {
				goto invalid_code
			}
			if ResolveNamedEntityHtml(start, ent_len, inv_map, &code, &code2) == zend.FAILURE {
				if doctype == 32 && ent_len == 4 && start[0] == 'a' && start[1] == 'p' && start[2] == 'o' && start[3] == 's' {

					/* uses html4 inv_map, which doesn't include apos;. This is a
					 * hack to support it */

					code = unsigned('\'')

					/* uses html4 inv_map, which doesn't include apos;. This is a
					 * hack to support it */

				} else {
					goto invalid_code
				}
			}
		}
		assert((*next) == ';')
		if code == '\'' && (flags&1) == 0 || code == '"' && (flags&2) == 0 {
			goto invalid_code
		}

		/* UTF-8 doesn't need mapping (ISO-8859-1 doesn't either, but
		 * the call is needed to ensure the codepoint <= U+00FF)  */

		if charset != CsUtf8 {

			/* replace unicode code point */

			if MapFromUnicode(code, charset, &code) == zend.FAILURE || code2 != 0 {
				goto invalid_code
			}

			/* replace unicode code point */

		}
		q += WriteOctetSequence((*uint8)(q), charset, code)
		if code2 {
			q += WriteOctetSequence((*uint8)(q), charset, code2)
		}

		/* jump over the valid entity; may go beyond size of buffer; np */

		p = next + 1
		continue
	invalid_code:
		for ; p < next; p++ {
			*(g.PostInc(&q)) = *p
		}
	}
	*q = '0'
	ret.len_ = size_t(q - ret.val)
}

/* }}} */

func UnescapeInverseMap(all int, flags int) *EntityHt {
	var document_type int = flags & (16 | 32)
	if all != 0 {
		switch document_type {
		case 0:

		case 32:
			return &EntHtHtml4
		case 16 | 32:
			return &EntHtHtml5
		default:
			return &EntHtBeApos
		}
	} else {
		switch document_type {
		case 0:
			return &EntHtBeNoapos
		default:
			return &EntHtBeApos
		}
	}
}

/* }}} */

func DetermineEntityTable(all int, doctype int) EntityTableOpt {
	var retval EntityTableOpt = EntityTableOpt{nil}
	assert(!(doctype == 16 && all != 0))
	if all != 0 {
		if doctype == (16 | 32) {
			retval.SetMsTable(EntityMsTableHtml5)
		} else {
			retval.SetMsTable(EntityMsTableHtml4)
		}
	} else {
		if doctype == 0 {
			retval.SetTable(Stage3TableBeNoapos00000)
		} else {
			retval.SetTable(Stage3TableBeApos00000)
		}
	}
	return retval
}

/* }}} */

func PhpUnescapeHtmlEntities(str *zend.ZendString, all int, flags int, hint_charset *byte) *zend.ZendString {
	var ret *zend.ZendString
	var charset EntityCharset
	var inverse_map *EntityHt
	var new_size int
	if !(memchr(str.val, '&', str.len_)) {
		return zend.ZendStringCopy(str)
	}
	if all != 0 {
		charset = DetermineCharset(hint_charset)
	} else {
		charset = Cs88591
	}

	/* don't use LIMIT_ALL! */

	new_size = str.len_ + str.len_/5 + 2
	if str.len_ > new_size {

		/* overflow, refuse to do anything */

		return zend.ZendStringCopy(str)

		/* overflow, refuse to do anything */

	}
	ret = zend.ZendStringAlloc(new_size, 0)
	inverse_map = UnescapeInverseMap(all, flags)

	/* replace numeric entities */

	TraverseForEntities(str.val, str.len_, ret, all, flags, inverse_map, charset)
	return ret
}

/* }}} */

func PhpEscapeHtmlEntities(old *uint8, oldlen int, all int, flags int, hint_charset string) *zend.ZendString {
	return PhpEscapeHtmlEntitiesEx(old, oldlen, all, flags, hint_charset, 1)
}

/* {{{ find_entity_for_char */

func FindEntityForChar(k uint, charset EntityCharset, table *EntityStage1Row, entity **uint8, entity_len *int, old *uint8, oldlen int, cursor *int) {
	var stage1_idx unsigned = (k & 0xfff000) >> 12
	var c *EntityStage3Row
	if stage1_idx > 0x1d {
		*entity = nil
		*entity_len = 0
		return
	}
	c = &table[stage1_idx][(k&0xfc0)>>6][k&0x3f]
	if !(c.GetAmbiguous()) {
		*entity = (*uint8)(c.GetEntity())
		*entity_len = c.GetEntityLen()
	} else {

		/* peek at next char */

		var cursor_before int = *cursor
		var status int = zend.SUCCESS
		var next_char unsigned
		if (*cursor) >= oldlen {
			goto no_suitable_2nd
		}
		next_char = GetNextChar(charset, old, oldlen, cursor, &status)
		if status == zend.FAILURE {
			goto no_suitable_2nd
		}
		var s *EntityMulticodepointRow
		var e *EntityMulticodepointRow
		s = &c.data.multicodepoint_table[1]
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

/* }}} */

func FindEntityForCharBasic(k uint, table *EntityStage3Row, entity **uint8, entity_len *int) {
	if k >= 64 {
		*entity = nil
		*entity_len = 0
		return
	}
	*entity = (*uint8)(table[k].GetEntity())
	*entity_len = table[k].GetEntityLen()
}

/* }}} */

func PhpEscapeHtmlEntitiesEx(old *uint8, oldlen int, all int, flags int, hint_charset *byte, double_encode zend.ZendBool) *zend.ZendString {
	var cursor int
	var maxlen int
	var len_ int
	var replaced *zend.ZendString
	var charset EntityCharset = DetermineCharset(hint_charset)
	var doctype int = flags & (16 | 32)
	var entity_table EntityTableOpt
	var to_uni_table *EncToUni = nil
	var inv_map *EntityHt = nil

	/* only used if flags includes ENT_HTML_IGNORE_ERRORS or ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS */

	var replacement *uint8 = nil
	var replacement_len int = 0
	if all != 0 {
		if charset >= CsBig5 {
			core.PhpErrorDocref(nil, 1<<3, "Only basic entities "+"substitution is supported for multi-byte encodings other than UTF-8; "+"functionality is equivalent to htmlspecialchars")
		}
		all = all != 0 && charset < CsBig5 && doctype != 16
	}
	entity_table = DetermineEntityTable(all, doctype)
	if all != 0 && charset > Cs88591 {
		to_uni_table = EncToUniIndex[charset]
	}
	if double_encode == 0 {

		/* first arg is 1 because we want to identify valid named entities
		 * even if we are only encoding the basic ones */

		inv_map = UnescapeInverseMap(1, flags)

		/* first arg is 1 because we want to identify valid named entities
		 * even if we are only encoding the basic ones */

	}
	if (flags & (8 | 128)) != 0 {
		if charset == CsUtf8 {
			replacement = (*uint8)("xEFxBFxBD")
			replacement_len = g.SizeOf("\"\\xEF\\xBF\\xBD\"") - 1
		} else {
			replacement = (*uint8)("&#xFFFD;")
			replacement_len = g.SizeOf("\"&#xFFFD;\"") - 1
		}
	}

	/* initial estimate */

	if oldlen < 64 {
		maxlen = 128
	} else {
		maxlen = zend.ZendSafeAddmult(oldlen, 2, 0, "html_entities")
	}
	replaced = zend.ZendStringAlloc(maxlen, 0)
	len_ = 0
	cursor = 0
	for cursor < oldlen {
		var mbsequence *uint8 = nil
		var mbseqlen int = 0
		var cursor_before int = cursor
		var status int = zend.SUCCESS
		var this_char uint = GetNextChar(charset, old, oldlen, &cursor, &status)

		/* guarantee we have at least 40 bytes to write.
		 * In HTML5, entities may take up to 33 bytes */

		if len_ > maxlen-40 {
			replaced = zend.ZendStringSafeRealloc(replaced, maxlen, 1, 128, 0)
			maxlen += 128
		}
		if status == zend.FAILURE {

			/* invalid MB sequence */

			if (flags & 4) != 0 {
				continue
			} else if (flags & 8) != 0 {
				memcpy(&replaced.val[len_], replacement, replacement_len)
				len_ += replacement_len
				continue
			} else {
				zend.ZendStringEfree(replaced)
				return zend.ZendEmptyString
			}

			/* invalid MB sequence */

		} else {
			mbsequence = &old[cursor_before]
			mbseqlen = cursor - cursor_before
		}
		if this_char != '&' {
			var rep *uint8 = nil
			var rep_len int = 0
			if this_char == '\'' && (flags&1) == 0 || this_char == '"' && (flags&2) == 0 {
				goto pass_char_through
			}
			if all != 0 {
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

				/* the cursor may advance */

			} else {
				FindEntityForCharBasic(this_char, entity_table.GetTable(), &rep, &rep_len)
			}
			if rep != nil {
				replaced.val[g.PostInc(&len_)] = '&'
				memcpy(&replaced.val[len_], rep, rep_len)
				len_ += rep_len
				replaced.val[g.PostInc(&len_)] = ';'
			} else {

				/* we did not find an entity for this char.
				 * check for its validity, if its valid pass it unchanged */

				if (flags & 128) != 0 {
					if charset <= Cs88591 {
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

						/* not a unicode code point, unless, coincidentally, it's in
						 * the 0x20..0x7D range (except 0x5C in sjis). We know nothing
						 * about other code points, because we have no tables. Since
						 * Unicode code points in that range are not disallowed in any
						 * document type, we could do nothing. However, conversion
						 * tables frequently map 0x00-0x1F to the respective C0 code
						 * points. Let's play it safe and admit that's the case */

					}
				}
			pass_char_through:
				if mbseqlen > 1 {
					memcpy(replaced.val+len_, mbsequence, mbseqlen)
					len_ += mbseqlen
				} else {
					replaced.val[g.PostInc(&len_)] = mbsequence[0]
				}
			}
		} else {
			if double_encode != 0 {
			encode_amp:
				memcpy(&replaced.val[len_], "&amp;", g.SizeOf("\"&amp;\"")-1)
				len_ += g.SizeOf("\"&amp;\"") - 1
			} else {

				/* check if entity is valid */

				var ent_len int

				/* peek at next char */

				if old[cursor] == '#' {
					var code_point unsigned
					var valid int
					var pos *byte = (*byte)(&old[cursor+1])
					valid = ProcessNumericEntity((**byte)(&pos), &code_point)
					if valid == zend.FAILURE {
						goto encode_amp
					}
					if (flags & 128) != 0 {
						if NumericEntityIsAllowed(code_point, doctype) == 0 {
							goto encode_amp
						}
					}
					ent_len = pos - (*byte)(&old[cursor])
				} else {

					/* check for vality of named entity */

					var start *byte = (*byte)(&old[cursor])
					var next *byte = start
					var dummy1 unsigned
					var dummy2 unsigned
					if ProcessNamedEntityHtml(&next, &start, &ent_len) == zend.FAILURE {
						goto encode_amp
					}
					if ResolveNamedEntityHtml(start, ent_len, inv_map, &dummy1, &dummy2) == zend.FAILURE {
						if !(doctype == 32 && ent_len == 4 && start[0] == 'a' && start[1] == 'p' && start[2] == 'o' && start[3] == 's') {

							/* uses html4 inv_map, which doesn't include apos;. This is a
							 * hack to support it */

							goto encode_amp

							/* uses html4 inv_map, which doesn't include apos;. This is a
							 * hack to support it */

						}
					}
				}

				/* checks passed; copy entity to result */

				if maxlen-len_ < ent_len+2 {

					/* ent_len < oldlen, which is certainly <= SIZE_MAX/2 */

					replaced = zend.ZendStringSafeRealloc(replaced, maxlen, 1, ent_len+128, 0)
					maxlen += ent_len + 128
				}
				replaced.val[g.PostInc(&len_)] = '&'
				memcpy(&replaced.val[len_], &old[cursor], ent_len)
				len_ += ent_len
				replaced.val[g.PostInc(&len_)] = ';'
				cursor += ent_len + 1
			}
		}
	}
	replaced.val[len_] = '0'
	replaced.len_ = len_
	return replaced
}

/* }}} */

func PhpHtmlEntities(execute_data *zend.ZendExecuteData, return_value *zend.Zval, all int) {
	var str *zend.ZendString
	var hint_charset *zend.ZendString = nil
	var default_charset *byte
	var flags zend.ZendLong = 2
	var replaced *zend.ZendString
	var double_encode zend.ZendBool = 1
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &hint_charset, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &double_encode, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if hint_charset == nil {
		default_charset = GetDefaultCharset()
	}
	replaced = PhpEscapeHtmlEntitiesEx((*uint8)(str.val), str.len_, all, int(flags), g.CondF1(hint_charset != nil, func() []byte { return hint_charset.val }, default_charset), double_encode)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = replaced
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

// #define HTML_SPECIALCHARS       0

// #define HTML_ENTITIES       1

/* {{{ register_html_constants
 */

func RegisterHtmlConstants(type_ int, module_number int) {
	zend.ZendRegisterLongConstant("HTML_SPECIALCHARS", g.SizeOf("\"HTML_SPECIALCHARS\"")-1, 0, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("HTML_ENTITIES", g.SizeOf("\"HTML_ENTITIES\"")-1, 1, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_COMPAT", g.SizeOf("\"ENT_COMPAT\"")-1, 2, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_QUOTES", g.SizeOf("\"ENT_QUOTES\"")-1, 2|1, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_NOQUOTES", g.SizeOf("\"ENT_NOQUOTES\"")-1, 0, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_IGNORE", g.SizeOf("\"ENT_IGNORE\"")-1, 4, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_SUBSTITUTE", g.SizeOf("\"ENT_SUBSTITUTE\"")-1, 8, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_DISALLOWED", g.SizeOf("\"ENT_DISALLOWED\"")-1, 128, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_HTML401", g.SizeOf("\"ENT_HTML401\"")-1, 0, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_XML1", g.SizeOf("\"ENT_XML1\"")-1, 16, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_XHTML", g.SizeOf("\"ENT_XHTML\"")-1, 32, 1<<1|1<<0, module_number)
	zend.ZendRegisterLongConstant("ENT_HTML5", g.SizeOf("\"ENT_HTML5\"")-1, 16|32, 1<<1|1<<0, module_number)
}

/* }}} */

func ZifHtmlspecialchars(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHtmlEntities(execute_data, return_value, 0)
}

/* }}} */

func ZifHtmlspecialcharsDecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var quote_style zend.ZendLong = 2
	var replaced *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &quote_style, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	replaced = PhpUnescapeHtmlEntities(str, 0, int(quote_style), nil)
	if replaced != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = replaced
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifHtmlEntityDecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var hint_charset *zend.ZendString = nil
	var default_charset *byte
	var quote_style zend.ZendLong = 2
	var replaced *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &quote_style, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &hint_charset, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if hint_charset == nil {
		default_charset = GetDefaultCharset()
	}
	replaced = PhpUnescapeHtmlEntities(str, 1, int(quote_style), g.CondF1(hint_charset != nil, func() []byte { return hint_charset.val }, default_charset))
	if replaced != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = replaced
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifHtmlentities(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHtmlEntities(execute_data, return_value, 1)
}

/* }}} */

func WriteS3rowData(r *EntityStage3Row, orig_cp unsigned, charset EntityCharset, arr *zend.Zval) {
	var key []byte = ""
	var entity []byte = []byte{'&'}
	var written_k1 int
	written_k1 = WriteOctetSequence((*uint8)(key), charset, orig_cp)
	if !(r.GetAmbiguous()) {
		var l int = r.GetEntityLen()
		memcpy(&entity[1], r.GetEntity(), l)
		entity[l+1] = ';'
		zend.AddAssocStringlEx(arr, key, written_k1, entity, l+2)
	} else {
		var i unsigned
		var num_entries unsigned
		var mcpr *EntityMulticodepointRow = r.GetMulticodepointTable()
		if mcpr[0].GetDefaultEntity() != nil {
			var l int = mcpr[0].GetDefaultEntityLen()
			memcpy(&entity[1], mcpr[0].GetDefaultEntity(), l)
			entity[l+1] = ';'
			zend.AddAssocStringlEx(arr, key, written_k1, entity, l+2)
		}
		num_entries = mcpr[0].GetSize()
		for i = 1; i <= num_entries; i++ {
			var l int
			var written_k2 int
			var uni_cp unsigned
			var spe_cp unsigned
			uni_cp = mcpr[i].GetSecondCp()
			l = mcpr[i].GetEntityLen()
			if charset > Cs88591 {
				if MapFromUnicode(uni_cp, charset, &spe_cp) == zend.FAILURE {
					continue
				}
			} else {
				spe_cp = uni_cp
			}
			written_k2 = WriteOctetSequence((*uint8)(&key[written_k1]), charset, spe_cp)
			memcpy(&entity[1], mcpr[i].GetEntity(), l)
			entity[l+1] = ';'
			zend.AddAssocStringlEx(arr, key, written_k1+written_k2, entity, l+2)
		}
	}
}

/* }}} */

func ZifGetHtmlTranslationTable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var all zend.ZendLong = 0
	var flags zend.ZendLong = 2
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
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &all, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &charset_hint, &charset_hint_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	charset = DetermineCharset(charset_hint)
	doctype = flags & (16 | 32)
	all = all != 0 && charset < CsBig5 && doctype != 16
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	entity_table = DetermineEntityTable(int(all), doctype)
	if all != 0 && charset > Cs88591 {
		to_uni_table = EncToUniIndex[charset]
	}
	if all != 0 {
		var ms_table *EntityStage1Row = entity_table.GetMsTable()
		if charset <= Cs88591 {
			var i unsigned
			var j unsigned
			var k unsigned
			var max_i unsigned
			var max_j unsigned
			var max_k unsigned

			/* no mapping to unicode required */

			if charset > CsUtf8 && charset < CsBig5 {
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
						var code unsigned
						if r.GetEntity() == nil {
							continue
						}
						code = i<<12 | j<<6 | k
						if code == '\'' && (flags&1) == 0 || code == '"' && (flags&2) == 0 {
							continue
						}
						WriteS3rowData(r, code, charset, return_value)
					}
				}
			}
		} else {

			/* we have to iterate through the set of code points for this
			 * encoding and map them to unicode code points */

			var i unsigned
			for i = 0; i <= 0xff; i++ {
				var r *EntityStage3Row
				var uni_cp unsigned

				/* can be done before mapping, they're invariant */

				if i == '\'' && (flags&1) == 0 || i == '"' && (flags&2) == 0 {
					continue
				}
				MapToUnicode(i, to_uni_table, &uni_cp)
				r = &ms_table[(uni_cp&0xfff000)>>12][(uni_cp&0xfc0)>>6][uni_cp&0x3f]
				if r.GetEntity() == nil {
					continue
				}
				WriteS3rowData(r, i, charset, return_value)
			}
		}
	} else {

		/* we could use sizeof(stage3_table_be_apos_00000) as well */

		var j unsigned
		var numelems unsigned = g.SizeOf("stage3_table_be_noapos_00000") / g.SizeOf("* stage3_table_be_noapos_00000")
		for j = 0; j < numelems; j++ {
			var r *EntityStage3Row = &entity_table.table[j]
			if r.GetEntity() == nil {
				continue
			}
			if j == '\'' && (flags&1) == 0 || j == '"' && (flags&2) == 0 {
				continue
			}

			/* charset is indifferent, used cs_8859_1 for efficiency */

			WriteS3rowData(r, j, Cs88591, return_value)

			/* charset is indifferent, used cs_8859_1 for efficiency */

		}
	}
}

/* }}} */

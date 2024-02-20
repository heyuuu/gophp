package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/ext/standard/charsets"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	EntHtmlQuoteNone        = 0
	EntHtmlQuoteSingle      = 1
	EntHtmlQuoteDouble      = 2
	EntHtmlIgnoreErrors     = 4
	EntHtmlSubstituteErrors = 8
	EntHtmlDocTypeMask      = 16 | 32
	EntHtmlDocHtml401       = 0
	EntHtmlDocXml1          = 16
	EntHtmlDocXhtml         = 32
	EntHtmlDocHtml5         = 16 | 32
	/* reserve bit 6 */
	EntHtmlSubstituteDisallowedChars = 128
)

const (
	EntCompat     = EntHtmlQuoteDouble
	EntQuotes     = EntHtmlQuoteDouble | EntHtmlQuoteSingle
	EntNoquotes   = EntHtmlQuoteNone
	EntIgnore     = EntHtmlIgnoreErrors
	EntSubstitute = EntHtmlSubstituteErrors
	EntHtml401    = 0
	EntXml1       = 16
	EntXhtml      = 32
	EntHtml5      = 16 | 32
	EntDisallowed = 128
)

const (
	HtmlSpecialchars = 0
	HtmlEntities     = 1
)

// --- functions

func mapToUnicode(uniTable EncToUni, code uint) uint {
	idx1, idx2 := (code&0xc0)>>6, code&0x3f

	/* only single byte encodings are currently supported; assumed code <= 0xFF */
	return uint(uniTable[idx1][idx2])
}
func mapEntityRow(table []EntityStage1Row, k uint) *EntityStage3Row {
	stage1Idx, stage2Idx, stage3Idx := (k&0xfff000)>>12, (k&0xfc0)>>6, k&0x3f
	if stage1Idx > 0x1d {
		return nil
	}
	stage1Row := table[stage1Idx]
	if stage1Row == nil {
		return nil
	}
	stage2Row := stage1Row[stage2Idx]
	if stage2Row == nil {
		return nil
	}
	return stage2Row[stage3Idx]
}
func codePointFromStages(i, j, k uint) uint {
	return i<<12 | j<<6 | k
}
func eachUnicodeEntity(charset charsets.Charset, msTable []EntityStage1Row, handler func(code uint, r *EntityStage3Row)) {
	if charset.UnicodeCompat() {
		eachUnicodeCompatEntity(charset, msTable, handler)
	} else {
		toUniTable := EncToUniIndex[charset]
		/* we have to iterate through the set of code points for this
		 * encoding and map them to unicode code points */

		var i uint
		for i = 0; i <= 0xff; i++ {
			uniCp := mapToUnicode(*toUniTable, i)
			r := mapEntityRow(msTable, uniCp)
			if r != nil {
				handler(i, r)
			}
		}
	}
}

func eachUnicodeCompatEntity(charset charsets.Charset, msTable []EntityStage1Row, handler func(code uint, r *EntityStage3Row)) {
	assert.Assert(charset.UnicodeCompat())

	/* no mapping to unicode required */
	var maxI, maxJ, maxK uint
	if charset.SingleByte() {
		maxI, maxJ, maxK = 1, 4, 64
	} else {
		maxI, maxJ, maxK = 0x1e, 64, 64
	}
	for i, stage1Row := range msTable[:maxI] {
		if stage1Row == nil {
			continue
		}
		for j, stage2Row := range stage1Row[:maxJ] {
			if stage2Row == nil {
				continue
			}
			for k, stage3Row := range stage2Row[:maxK] {
				if stage3Row == nil {
					continue
				}
				code := codePointFromStages(uint(i), uint(j), uint(k))
				handler(code, stage3Row)
			}
		}
	}
}

func charsetLimitAll(all bool, doctype int, charset charsets.Charset) bool {
	return all && !charset.PartialSupport() && doctype != EntHtmlDocXml1
}
func GetDefaultCharset(ctx *php.Context) string {
	if internalEncoding := ctx.PG().InternalEncoding(); internalEncoding != "" {
		return internalEncoding
	}
	if defaultCharset := ctx.PG().DefaultCharset(); defaultCharset != "" {
		return defaultCharset
	}
	return ""
}

func DetermineCharset(ctx *php.Context, charsetHint string) charsets.Charset {
	return DetermineCharsetEx(ctx, &charsetHint)
}
func DetermineCharsetEx(ctx *php.Context, charsetHint *string) charsets.Charset {
	var charset = charsets.CsUtf8

	/* Default is now UTF-8 */
	if charsetHint == nil {
		return charsets.CsUtf8
	}
	charsetStr := *charsetHint
	if charsetStr == "" {
		charsetStr = ctx.PG().DefaultCharset()
	}
	if charsetStr != "" {
		/* look for the codeset */
		if matchCharset, found := charsets.GetCharset(charsetStr); found {
			charset = matchCharset
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("charset `%s' not supported, assuming utf-8", charsetStr))
		}
	}
	return charset
}

func UnimapBsearch(table []UniToEnc, codeKeyA uint) (uint8, bool) {
	/* we have no mappings outside the BMP */
	if codeKeyA > 0xffff {
		return 0, false
	}

	codeKey := uint16(codeKeyA)
	idx := sort.Search(len(table), func(i int) bool {
		return codeKey >= table[i].UnCodePoint()
	})
	if idx < len(table) && table[idx].UnCodePoint() == codeKey {
		return table[idx].CsCode(), true
	}
	return 0, false
}
func MapFromUnicode(code uint, charset charsets.Charset) (uint8, bool) {
	switch charset {
	case charsets.Cs88591:
		/* identity mapping of code points to unicode */
		if code > 0xff {
			return 0, false
		}
		return uint8(code), true
	case charsets.Cs88595:
		var res uint8
		if code <= 0xa0 || code == 0xad {
			res = uint8(code)
		} else if code == 0x2116 {
			res = 0xf0
		} else if code == 0xa7 {
			res = 0xfd
		} else if code >= 0x401 && code <= 0x44f {
			if code == 0x40d || code == 0x450 || code == 0x45d {
				return 0, false
			}
			res = uint8(code - 0x360)
		} else {
			return 0, false
		}
		return res, true
	case charsets.Cs885915:
		if code < 0xa4 || code > 0xbe && code <= 0xff {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapIso885915, code)
		}
	case charsets.CsCp1252:
		if code <= 0x7f || code >= 0xa0 && code <= 0xff {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapWin1252, code)
		}
	case charsets.CsMacroman:
		if code == 0x7f {
			return 0, false
		} else if code < 0x7f {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapMacroman, code)
		}
	case charsets.CsCp1251:
		if code <= 0x7f {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapWin1251, code)
		}
	case charsets.CsKoi8r:
		if code <= 0x7f {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapKoi8r, code)
		}
	case charsets.CsCp866:
		if code <= 0x7f {
			return uint8(code), true
		} else {
			return UnimapBsearch(UnimapCp866, code)
		}
	case charsets.CsSjis,
		charsets.CsEucjp:

		/* we interpret 0x5C as the Yen symbol. This is not universal.
		 * See <http://www.w3.org/Submission/japanese-xml/#ambiguity_of_yen> */
		if code >= 0x20 && code <= 0x7d && code != 0x5c {
			return uint8(code), true
		} else {
			return 0, false
		}
	case charsets.CsBig5,
		charsets.CsBig5hkscs,
		charsets.CsGb2312:
		if code >= 0x20 && code <= 0x7d {
			return uint8(code), true
		} else {
			return 0, false
		}
	default:
		return 0, false
	}
}
func UnicodeCpIsAllowed(uniCp uint, documentType int) bool {
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
	switch documentType {
	case EntHtmlDocHtml401:
		return uniCp >= 0x20 && uniCp <= 0x7e || (uniCp == 0xa || uniCp == 0x9 || uniCp == 0xd) || uniCp >= 0xa0 && uniCp <= 0xd7ff || uniCp >= 0xe000 && uniCp <= 0x10ffff
	case EntHtmlDocHtml5:
		return uniCp >= 0x20 && uniCp <= 0x7e || uniCp >= 0x9 && uniCp <= 0xd && uniCp != 0xb || uniCp >= 0xa0 && uniCp <= 0xd7ff || uniCp >= 0xe000 && uniCp <= 0x10ffff && (uniCp&0xffff) < 0xfffe && (uniCp < 0xfdd0 || uniCp > 0xfdef)
	case EntHtmlDocXhtml:
		fallthrough
	case EntHtmlDocXml1:
		return uniCp >= 0x20 && uniCp <= 0xd7ff || (uniCp == 0xa || uniCp == 0x9 || uniCp == 0xd) || uniCp >= 0xe000 && uniCp <= 0x10ffff && uniCp != 0xfffe && uniCp != 0xffff
	default:
		return true
	}
}
func NumericEntityIsAllowed(uniCp uint, documentType int) bool {
	/* less restrictive than unicode_cp_is_allowed */

	switch documentType {
	case EntHtmlDocHtml401:
		/* all non-SGML characters (those marked with UNUSED in DESCSET) should be
		 * representable with numeric entities */
		return uniCp <= 0x10ffff
	case EntHtmlDocHtml5:
		/* 8.1.4. The numeric character reference forms described above are allowed to
		 * reference any Unicode code point other than U+0000, U+000D, permanently
		 * undefined Unicode characters (noncharacters), and control characters other
		 * than space characters (U+0009, U+000A, U+000C and U+000D) */
		return uniCp >= 0x20 && uniCp <= 0x7e || uniCp >= 0x9 && uniCp <= 0xc && uniCp != 0xb || uniCp >= 0xa0 && uniCp <= 0x10ffff && (uniCp&0xffff) < 0xfffe && (uniCp < 0xfdd0 || uniCp > 0xfdef)
	case EntHtmlDocXhtml, EntHtmlDocXml1:
		/* OTOH, XML 1.0 requires "character references to match the production for Char
		 * See <http://www.w3.org/TR/REC-xml/#NT-CharRef> */
		return UnicodeCpIsAllowed(uniCp, documentType)
	default:
		return true
	}
}

func ProcessNumericEntity(s string) (cp uint, n int, ok bool) {
	if s == "" {
		return 0, 0, false
	}

	var hexadecimal = s[0] == 'x' || s[0] == 'X'
	if hexadecimal {
		s = s[1:]
		n++
	}

	/* strtol allows whitespace and other stuff in the beginning we're not interested */
	if s == "" || hexadecimal && !ascii.IsXDigit(s[0]) || !hexadecimal && !(ascii.IsDigit(s[0])) {
		return 0, n, false
	}

	code, used := php.ParseLongPrefix(s, lang.Cond(hexadecimal, 16, 10))
	s = s[used:]
	n += used

	/* we're guaranteed there were valid digits, so *endptr > buf */
	if s == "" || s[0] != ';' {
		return 0, n, false
	}

	/* many more are invalid, but that depends on whether it's HTML
	 * (and which version) or XML. */
	if code > 0x10ffff {
		return 0, n, false
	}
	return uint(code), n, true
}

func ProcessNamedEntityHtml(s string) (name string, n int, ok bool) {
	/* "&" is represented by a 0x26 in all supported encodings. That means
	 * the byte after represents a character or is the leading byte of an
	 * sequence of 8-bit code units. If in the ranges below, it represents
	 * necessarily a alpha character because none of the supported encodings
	 * has an overlap with ASCII in the leading byte (only on the second one) */

	idx := 0
	for idx < len(s) && ascii.IsAlphaNum(s[idx]) {
		idx++
	}
	if idx == 0 || idx >= len(s) || s[idx] != ';' {
		return "", idx, false
	}

	/* cast to size_t OK as the quantity is always non-negative */
	return s[:idx], idx, true
}

func ResolveNamedEntityHtml(ht EntityMap, key string) (cp1 uint, cp2 uint, ok bool) {
	if cp, ok := ht[key]; ok {
		return cp[0], cp[1], true
	}
	return 0, 0, false
}

func WriteOctetSequenceToString(charset charsets.Charset, code uint) string {
	return string(WriteOctetSequenceToBytes(charset, code))
}

func WriteOctetSequenceToBytes(charset charsets.Charset, code uint) []byte {
	/* code is not necessarily a unicode code point */
	switch charset {
	case charsets.CsUtf8:
		/* assert(0x0 <= k <= 0x10FFFF); */
		/* UTF-8 has been restricted to max 4 bytes since RFC 3629 */
		var buf [4]byte
		n := utf8.EncodeRune(buf[:], rune(code))
		return buf[:n]
	case charsets.Cs88591, charsets.CsCp1252, charsets.Cs885915, charsets.CsKoi8r, charsets.CsCp1251, charsets.Cs88595, charsets.CsCp866, charsets.CsMacroman:
		/* single byte stuff */
		return []byte{byte(code)}
	case charsets.CsBig5, charsets.CsBig5hkscs, charsets.CsSjis, charsets.CsGb2312:
		/* we don't have complete unicode mappings for these yet in entity_decode,
		 * and we opt to pass through the octet sequences for these in htmlentities
		 * instead of converting to an int and then converting back. */
		return []byte{byte(code)}
	case charsets.CsEucjp:
		return []byte{byte(code)}
	default:
		php.Assert(false)
		return nil
	}
}

func TraverseForEntities(s string, all bool, flags int, inv_map EntityMap, charset charsets.Charset) string {
	var doctype = flags & EntHtmlDocTypeMask
	var buf strings.Builder

	p := s
	for len(p) > 0 {
		/* Shift JIS, Big5 and HKSCS use multi-byte encodings where an
		 * ASCII range byte can be part of a multi-byte sequence.
		 * However, they start at 0x40, therefore if we find a 0x26 byte,
		 * we're sure it represents the '&' character. */
		if len(p) <= 3 {
			buf.WriteString(p)
			break
		} else if p[0] != '&' {
			buf.WriteByte(p[0])
			p = p[1:]
			continue
		}

		/* now p[3] is surely valid and is no terminator */
		var read int
		var code, code2 uint
		if p[1] == '#' {
			var n int
			var ok bool
			code, n, ok = ProcessNumericEntity(p[2:])
			read = 2 + n
			if !ok {
				goto invalidCode
			}

			/* If we're in htmlspecialchars_decode, we're only decoding entities
			 * that represent &, <, >, " and '. Is this one of them? */
			if !all && (code > 63 || Stage3TableBeApos00000[code] == nil) {
				goto invalidCode
			}

			/* are we allowed to decode this entity in this document type?
			 * HTML 5 is the only that has a character that cannot be used in
			 * a numeric entity but is allowed literally (U+000D). The
			 * unoptimized version would be ... || !numeric_entity_is_allowed(code) */

			if !UnicodeCpIsAllowed(code, doctype) || doctype == EntHtmlDocHtml5 && code == 0xd {
				goto invalidCode
			}
		} else {
			name, n, ok := ProcessNamedEntityHtml(p[1:])
			read = 1 + n
			if !ok {
				goto invalidCode
			}
			if code, code2, ok = ResolveNamedEntityHtml(inv_map, name); !ok {
				if doctype == EntHtmlDocXhtml && name == "apos" {
					/* uses html4 inv_map, which doesn't include apos;. This is a
					 * hack to support it */
					code = uint('\'')
				} else {
					goto invalidCode
				}
			}
		}
		assert.Assert(p[read] == ';')
		if code == '\'' && (flags&EntHtmlQuoteSingle) == 0 || code == '"' && (flags&EntHtmlQuoteDouble) == 0 {
			goto invalidCode
		}
		p = p[read+1:]

		/* UTF-8 doesn't need mapping (ISO-8859-1 doesn't either, but
		 * the call is needed to ensure the codepoint <= U+00FF)  */
		if charset != charsets.CsUtf8 {
			/* replace unicode code point */
			ret, ok := MapFromUnicode(code, charset)
			if ok {
				code = uint(ret)
			}
			if !ok || code2 != 0 {
				goto invalidCode
			}
		}
		buf.Write(WriteOctetSequenceToBytes(charset, code))
		if code2 != 0 {
			buf.Write(WriteOctetSequenceToBytes(charset, code2))
		}

		continue
	invalidCode:
		buf.WriteString(p[:read])
		p = p[read:]
	}
	return buf.String()
}

func UnescapeInverseMap(all bool, flags int) EntityMap {
	var documentType = flags & EntHtmlDocTypeMask
	if all {
		switch documentType {
		case EntHtmlDocHtml401:
			fallthrough
		case EntHtmlDocXhtml:
			return EntHtHtml4
		case EntHtmlDocHtml5:
			return EntHtHtml5
		default:
			return EntHtBeApos
		}
	} else {
		switch documentType {
		case EntHtmlDocHtml401:
			return EntHtBeNoapos
		default:
			return EntHtBeApos
		}
	}
}
func DetermineEntityTable(all bool, doctype int) EntityTableOpt {
	php.Assert(!(doctype == EntHtmlDocXml1 && all))

	if all {
		if doctype == EntHtmlDocHtml5 {
			return MakeEntityTableOpt(EntityMsTableHtml5, nil)
		} else {
			return MakeEntityTableOpt(EntityMsTableHtml4, nil)
		}
	} else {
		if doctype == EntHtmlDocHtml401 {
			return MakeEntityTableOpt(nil, Stage3TableBeNoapos00000)
		} else {
			return MakeEntityTableOpt(nil, Stage3TableBeApos00000)
		}
	}
}
func PhpUnescapeHtmlEntities(ctx *php.Context, str string, all bool, flags int, hintCharset string) string {
	if strings.IndexByte(str, '&') < 0 {
		return str
	}

	var charset charsets.Charset
	if all {
		charset = DetermineCharset(ctx, hintCharset)
	} else {
		charset = charsets.Cs88591 /* charset shouldn't matter, use ISO-8859-1 for performance */
	}

	/* don't use LIMIT_ALL! */
	inverseMap := UnescapeInverseMap(all, flags)

	/* replace numeric entities */
	return TraverseForEntities(str, all, flags, inverseMap, charset)
}

func PhpEscapeHtmlEntities(ctx *php.Context, old string, all bool, flags int, hintCharset string) string {
	return PhpEscapeHtmlEntitiesEx(ctx, old, all, flags, hintCharset, true)
}

func FindEntityForChar(k uint, table []EntityStage1Row, charReader *charsets.CharReader) string {
	var c = mapEntityRow(table, k)
	if c == nil {
		return ""
	}

	if !(c.Ambiguous()) {
		return c.Entity()
	} else {
		var mcpr = c.MultiCodepointTable()

		/* peek at next char */
		nextChar, n, ok := charReader.PeekChar()
		if !ok {
			return mcpr[0].DefaultEntity()
		}

		/* we could do a binary search but it's not worth it since we have at most two entries... */
		for _, row := range mcpr[1:] {
			if row.SecondCp() == nextChar {
				charReader.Skip(n)
				return row.Entity()
			}
		}

		return c.MultiCodepointTable()[0].DefaultEntity()
	}
}
func FindEntityForCharBasic(k uint, table []*EntityStage3Row) string {
	if k >= 64 || table[k] == nil {
		return ""
	}
	return table[k].Entity()
}
func PhpEscapeHtmlEntitiesEx(ctx *php.Context, old string, all bool, flags int, hintCharset string, doubleEncode bool) string {
	var buf strings.Builder
	var charset = DetermineCharset(ctx, hintCharset)
	var doctype = flags & EntHtmlDocTypeMask
	var entityTable EntityTableOpt
	var toUniTable *EncToUni = nil
	var invMap EntityMap = nil

	/* only used if flags includes ENT_HTML_IGNORE_ERRORS or ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS */
	var replacement string
	if all {
		if charset.PartialSupport() {
			php.ErrorDocRef(ctx, "", perr.E_NOTICE, "Only basic entities substitution is supported for multi-byte encodings other than UTF-8; functionality is equivalent to htmlspecialchars")
		}
		all = charsetLimitAll(all, doctype, charset)
	}
	entityTable = DetermineEntityTable(all, doctype)
	if all && !charset.UnicodeCompat() {
		toUniTable = EncToUniIndex[charset]
	}
	if !doubleEncode {
		/* first arg is 1 because we want to identify valid named entities
		 * even if we are only encoding the basic ones */
		invMap = UnescapeInverseMap(true, flags)
	}
	if (flags & (EntHtmlSubstituteErrors | EntHtmlSubstituteDisallowedChars)) != 0 {
		if charset == charsets.CsUtf8 {
			replacement = "\xEF\xBF\xBD"
		} else {
			replacement = "&#xFFFD;"
		}
	}

	charReader := charsets.NewCharReader(charset, old)

	/* initial estimate */
	for charReader.Valid() {
		thisChar, mbsequence, ok := charReader.ReadChar()

		/* guarantee we have at least 40 bytes to write.
		 * In HTML5, entities may take up to 33 bytes */
		if !ok {
			/* invalid MB sequence */
			if (flags & EntHtmlIgnoreErrors) != 0 {
				continue
			} else if (flags & EntHtmlSubstituteErrors) != 0 {
				buf.WriteString(replacement)
				continue
			} else {
				return ""
			}
		}
		if thisChar != '&' {
			var repStr string
			if thisChar == '\'' && (flags&EntHtmlQuoteSingle) == 0 || thisChar == '"' && (flags&EntHtmlQuoteDouble) == 0 {
				buf.WriteString(mbsequence)
				continue
			}
			if all {
				if toUniTable != nil {
					/* !CHARSET_UNICODE_COMPAT therefore not UTF-8; since UTF-8
					 * is the only multibyte encoding with !CHARSET_PARTIAL_SUPPORT,
					 * we're using a single byte encoding */
					thisChar = mapToUnicode(*toUniTable, thisChar)
					if thisChar == 0xffff {
						buf.WriteString(mbsequence)
						continue
					}
				}

				/* the cursor may advance */
				repStr = FindEntityForChar(thisChar, entityTable.MsTable(), charReader)
			} else {
				repStr = FindEntityForCharBasic(thisChar, entityTable.Table())
			}
			if repStr != "" {
				buf.WriteByte('&')
				buf.WriteString(repStr)
				buf.WriteByte(';')
			} else {

				/* we did not find an entity for this char.
				 * check for its validity, if its valid pass it unchanged */
				if (flags & EntHtmlSubstituteDisallowedChars) != 0 {
					if charset.UnicodeCompat() {
						if !UnicodeCpIsAllowed(thisChar, doctype) {
							mbsequence = replacement
						}
					} else if toUniTable != nil {
						if !all {
							thisChar = mapToUnicode(*toUniTable, thisChar)
						}
						if !UnicodeCpIsAllowed(thisChar, doctype) {
							mbsequence = replacement
						}
					} else {

						/* not a unicode code point, unless, coincidentally, it's in
						 * the 0x20..0x7D range (except 0x5C in sjis). We know nothing
						 * about other code points, because we have no tables. Since
						 * Unicode code points in that range are not disallowed in any
						 * document type, we could do nothing. However, conversion
						 * tables frequently map 0x00-0x1F to the respective C0 code
						 * points. Let's play it safe and admit that's the case */

						if thisChar <= 0x7d && !UnicodeCpIsAllowed(thisChar, doctype) {
							mbsequence = replacement
						}
					}
				}
				buf.WriteString(mbsequence)
			}
		} else {
			if doubleEncode {
				buf.WriteString("&amp;")
			} else {
				/* check if entity is valid */
				var ent_len int

				/* peek at next char */
				left := charReader.Left()
				if left != "" && left[0] == '#' {
					code, n, ok := ProcessNumericEntity(left[1:])
					if !ok {
						buf.WriteString("&amp;")
						continue
					}
					if (flags & EntHtmlSubstituteDisallowedChars) != 0 {
						if !NumericEntityIsAllowed(code, doctype) {
							buf.WriteString("&amp;")
							continue
						}
					}
					ent_len = n + 1
				} else {
					/* check for vality of named entity */
					name, n, ok := ProcessNamedEntityHtml(left)
					if !ok {
						buf.WriteString("&amp;")
						continue
					}
					if _, _, ok = ResolveNamedEntityHtml(invMap, name); !ok {
						if !(doctype == EntHtmlDocXhtml && name == "apos") {
							/* uses html4 inv_map, which doesn't include apos;. This is a hack to support it */
							buf.WriteString("&amp;")
							continue
						}
					}
					ent_len = n
				}

				/* checks passed; copy entity to result */
				buf.WriteByte('&')
				buf.WriteString(left[:ent_len])
				buf.WriteByte(';')

				charReader.Skip(ent_len + 1)
			}
		}
	}
	return buf.String()
}
func PhpHtmlEntities(ctx *php.Context, str string, flags_ *int, encoding *string, doubleEncode bool, all bool) string {
	var flags = lang.Option(flags_, EntCompat)
	var hintCharset string
	if encoding == nil {
		hintCharset = GetDefaultCharset(ctx)
	} else {
		hintCharset = *encoding
	}
	return PhpEscapeHtmlEntitiesEx(ctx, str, all, flags, hintCharset, doubleEncode)
}
func RegisterHtmlConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterLongConstant(ctx, moduleNumber, "HTML_SPECIALCHARS", HtmlSpecialchars)
	php.RegisterLongConstant(ctx, moduleNumber, "HTML_ENTITIES", HtmlEntities)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_COMPAT", EntCompat)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_QUOTES", EntQuotes)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_NOQUOTES", EntNoquotes)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_IGNORE", EntIgnore)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_SUBSTITUTE", EntSubstitute)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_DISALLOWED", EntDisallowed)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_HTML401", EntHtml401)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_XML1", EntXml1)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_XHTML", EntXhtml)
	php.RegisterLongConstant(ctx, moduleNumber, "ENT_HTML5", EntHtml5)
}
func ZifHtmlspecialchars(ctx *php.Context, string string, _ zpp.Opt, quoteStyle *int, encoding *string, doubleEncode_ *bool) string {
	var doubleEncode = lang.Option(doubleEncode_, true)
	return PhpHtmlEntities(ctx, string, quoteStyle, encoding, doubleEncode, false)
}
func ZifHtmlspecialcharsDecode(ctx *php.Context, str string, _ zpp.Opt, quoteStyle_ *int) string {
	var quoteStyle = lang.Option(quoteStyle_, EntCompat)
	return PhpUnescapeHtmlEntities(ctx, str, false, quoteStyle, "")
}
func ZifHtmlEntityDecode(ctx *php.Context, str string, _ zpp.Opt, quoteStyle_ *int, encoding *string) string {
	quoteStyle := lang.Option(quoteStyle_, EntCompat)
	var charset string
	if encoding != nil {
		charset = *encoding
	} else {
		charset = GetDefaultCharset(ctx)
	}
	return PhpUnescapeHtmlEntities(ctx, str, true, quoteStyle, charset)
}
func ZifHtmlentities(ctx *php.Context, string string, _ zpp.Opt, quoteStyle *int, encoding *string, doubleEncode_ *bool) string {
	var doubleEncode = lang.Option(doubleEncode_, true)
	return PhpHtmlEntities(ctx, string, quoteStyle, encoding, doubleEncode, true)
}
func writeS3rowData(r *EntityStage3Row, origCp uint, charset charsets.Charset, arr_ *types.Array) {
	key1 := WriteOctetSequenceToString(charset, origCp)
	if !(r.Ambiguous()) {
		arr_.AddAssocStr(key1, "&"+r.Entity()+";")
	} else {
		var mcpr = r.MultiCodepointTable()
		if mcpr[0].DefaultEntity() != "" {
			arr_.AddAssocStr(key1, "&"+mcpr[0].DefaultEntity()+";")
		}

		numEntries := mcpr[0].Size()
		for i := uint(1); i <= numEntries; i++ {
			var speCp uint
			uniCp := mcpr[i].SecondCp()
			if !charset.UnicodeCompat() {
				if ret, ok := MapFromUnicode(uniCp, charset); ok {
					speCp = uint(ret)
				} else {
					continue
				}
			} else {
				speCp = uniCp
			}
			key2 := WriteOctetSequenceToString(charset, speCp)
			arr_.AddAssocStr(key1+key2, "&"+mcpr[i].Entity()+";")
		}
	}
}

func ZifGetHtmlTranslationTable(ctx *php.Context, returnValue zpp.Ret, _ zpp.Opt, table *int, quoteStyle_ *int, encoding *string) *types.Array {
	var all = lang.Option(table, HtmlSpecialchars) != 0
	var flags = lang.Option(quoteStyle_, EntCompat)
	var doctype int
	var entityTable EntityTableOpt
	var charset charsets.Charset

	/* in this function we have to jump through some loops because we're
	 * getting the translated table from data structures that are optimized for
	 * random access, not traversal */
	charset = DetermineCharsetEx(ctx, encoding)
	doctype = flags & EntHtmlDocTypeMask
	all = charsetLimitAll(all, doctype, charset)
	returnValue.SetEmptyArray()
	entityTable = DetermineEntityTable(all, doctype)

	arr := types.NewArray()
	if all {
		var msTable = entityTable.MsTable()
		eachUnicodeEntity(charset, msTable, func(code uint, r *EntityStage3Row) {
			if code == '\'' && (flags&EntHtmlQuoteSingle) == 0 || code == '"' && (flags&EntHtmlQuoteDouble) == 0 {
				return
			}
			writeS3rowData(r, code, charset, arr)
		})
	} else {
		for code, r := range entityTable.Table() {
			if r == nil {
				continue
			}
			if code == '\'' && (flags&EntHtmlQuoteSingle) == 0 || code == '"' && (flags&EntHtmlQuoteDouble) == 0 {
				continue
			}

			/* charset is indifferent, used cs_8859_1 for efficiency */
			writeS3rowData(r, uint(code), charsets.Cs88591, arr)
		}
	}
	return arr
}

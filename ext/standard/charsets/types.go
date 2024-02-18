package charsets

import "github.com/heyuuu/gophp/kits/ascii"

type Charset uint8

const (
	CsUtf8 Charset = iota
	Cs88591
	CsCp1252
	Cs885915
	CsCp1251
	Cs88595
	CsCp866
	CsMacroman
	CsKoi8r
	CsBig5
	CsGb2312
	CsBig5hkscs
	CsSjis
	CsEucjp
)

// @see: CHARSET_UNICODE_COMPAT
func (cs Charset) UnicodeCompat() bool { return cs <= Cs88591 }

// @see: CHARSET_SINGLE_BYTE
func (cs Charset) SingleByte() bool { return cs > CsUtf8 && cs < CsBig5 }

// @see: CHARSET_SINGLE_BYTE
func (cs Charset) PartialSupport() bool { return cs >= CsBig5 }

//
var charsetMap = map[string]Charset{
	"iso-8859-1":   Cs88591,     // ISO-8859-1
	"iso8859-1":    Cs88591,     // ISO8859-1
	"iso-8859-15":  Cs885915,    // ISO-8859-15
	"iso8859-15":   Cs885915,    // ISO8859-15
	"utf-8":        CsUtf8,      // utf-8
	"cp1252":       CsCp1252,    // cp1252
	"windows-1252": CsCp1252,    // Windows-1252
	"1252":         CsCp1252,    // 1252
	"big5":         CsBig5,      // BIG5
	"950":          CsBig5,      // 950
	"gb2312":       CsGb2312,    // GB2312
	"936":          CsGb2312,    // 936
	"big5-hkscs":   CsBig5hkscs, // BIG5-HKSCS
	"shift_jis":    CsSjis,      // Shift_JIS
	"sjis":         CsSjis,      // SJIS
	"932":          CsSjis,      // 932
	"sjis-win":     CsSjis,      // SJIS-win
	"cp932":        CsSjis,      // CP932
	"eucjp":        CsEucjp,     // EUCJP
	"euc-jp":       CsEucjp,     // EUC-JP
	"eucjp-win":    CsEucjp,     // eucJP-win
	"koi8-r":       CsKoi8r,     // KOI8-R
	"koi8-ru":      CsKoi8r,     // koi8-ru
	"koi8r":        CsKoi8r,     // koi8r
	"cp1251":       CsCp1251,    // cp1251
	"windows-1251": CsCp1251,    // Windows-1251
	"win-1251":     CsCp1251,    // win-1251
	"iso8859-5":    Cs88595,     // iso8859-5
	"iso-8859-5":   Cs88595,     // iso-8859-5
	"cp866":        CsCp866,     // cp866
	"866":          CsCp866,     // 866
	"ibm866":       CsCp866,     // ibm866
	"macroman":     CsMacroman,  // MacRoman
}

func GetCharset(hint string) (Charset, bool) {
	lcHint := ascii.StrToLower(hint)
	if charset, ok := charsetMap[lcHint]; ok {
		return charset, ok
	}
	return 0, false
}

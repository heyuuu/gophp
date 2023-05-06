package standard

import (
	"github.com/heyuuu/gophp/builtin/ascii"
)

type charsetEntry struct {
	codeset string
	charset EntityCharset
}

var charsets = []charsetEntry{
	{"ISO-8859-1", Cs88591},
	{"ISO8859-1", Cs88591},
	{"ISO-8859-15", Cs885915},
	{"ISO8859-15", Cs885915},
	{"utf-8", CsUtf8},
	{"cp1252", CsCp1252},
	{"Windows-1252", CsCp1252},
	{"1252", CsCp1252},
	{"BIG5", CsBig5},
	{"950", CsBig5},
	{"GB2312", CsGb2312},
	{"936", CsGb2312},
	{"BIG5-HKSCS", CsBig5hkscs},
	{"Shift_JIS", CsSjis},
	{"SJIS", CsSjis},
	{"932", CsSjis},
	{"SJIS-win", CsSjis},
	{"CP932", CsSjis},
	{"EUCJP", CsEucjp},
	{"EUC-JP", CsEucjp},
	{"eucJP-win", CsEucjp},
	{"KOI8-R", CsKoi8r},
	{"koi8-ru", CsKoi8r},
	{"koi8r", CsKoi8r},
	{"cp1251", CsCp1251},
	{"Windows-1251", CsCp1251},
	{"win-1251", CsCp1251},
	{"iso8859-5", Cs88595},
	{"iso-8859-5", Cs88595},
	{"cp866", CsCp866},
	{"866", CsCp866},
	{"ibm866", CsCp866},
	{"MacRoman", CsMacroman},
}
var lcCharsets = buildLcCharsets()

func buildLcCharsets() map[string]charsetEntry {
	m := make(map[string]charsetEntry, len(charsets))
	for _, charset := range charsets {
		m[ascii.StrToLower(charset.codeset)] = charset
	}
	return m
}

func GetCharset(hint string) (EntityCharset, bool) {
	lcHint := ascii.StrToLower(hint)
	if charset, ok := lcCharsets[lcHint]; ok {
		return charset.charset, ok
	}
	return 0, false
}

func CheckCodeSet(hint string) (string, bool) {
	lcHint := ascii.StrToLower(hint)
	if charset, ok := lcCharsets[lcHint]; ok {
		return charset.codeset, ok
	}
	return "", false
}

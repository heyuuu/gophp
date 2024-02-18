package charsets

import (
	"github.com/heyuuu/gophp/php/lang"
	"unicode/utf8"
)

type CharDecoder = func(s string) (char uint, n int, ok bool)

func GetCharDecoder(charset Charset) CharDecoder {
	switch charset {
	case CsUtf8:
		return DecodeCharUtf8
	case CsBig5:
		return DecodeCharBig5
	case CsBig5hkscs:
		return DecodeCharBig5hkscs
	case CsGb2312:
		return DecodeCharGb2312
	case CsSjis:
		return DecodeCharSjis
	case CsEucjp:
		return DecodeCharEucjp
	default:
		return DecodeCharSingleByte
	}
}

func DecodeCharUtf8(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return 0, size, false
	} else {
		return uint(r), size, true
	}
}

func DecodeCharBig5(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	/* reference http://demo.icu-project.org/icu-bin/convexp?conv=big5 */
	var c = s[0]
	if c >= 0x81 && c <= 0xfe {
		if len(s) < 2 {
			return 0, 1, false
		}
		next := s[1]
		if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
			char = uint(c)<<8 | uint(next)
		} else {
			return 0, 1, false
		}
		return char, 2, true
	} else {
		return uint(c), 1, true
	}
}

func DecodeCharBig5hkscs(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	var c = s[0]
	if c >= 0x81 && c <= 0xfe {
		if len(s) < 2 {
			return 0, 1, false
		}
		next := s[1]
		if next >= 0x40 && next <= 0x7e || next >= 0xa1 && next <= 0xfe {
			char = uint(c)<<8 | uint(next)
		} else if next != 0x80 && next != 0xff {
			return 0, 1, false
		} else {
			return 0, 2, false
		}
		return char, 2, true
	} else {
		return uint(c), 1, true
	}
}

func Gb2312Lead(c uint8) bool {
	return c != 0x8e && c != 0x8f && c != 0xa0 && c != 0xff
}
func Gb2312Trail(c uint8) bool { return c >= 0xa1 && c <= 0xfe }

func DecodeCharGb2312(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	var c = s[0]
	if c >= 0xa1 && c <= 0xfe {
		var next uint8
		if len(s) < 2 {
			return 0, 1, false
		}
		next = s[1]
		if Gb2312Trail(next) {
			char = uint(c)<<8 | uint(next)
		} else if Gb2312Lead(next) {
			return 0, 1, false
		} else {
			return 0, 2, false
		}
		return char, 2, true
	} else if Gb2312Lead(c) {
		return uint(c), 1, true
	} else {
		return 0, 1, false
	}
}

func SjisLead(c uint8) bool {
	return c != 0x80 && c != 0xa0 && c < 0xfd
}
func SjisTrail(c uint8) bool {
	return c >= 0x40 && c != 0x7f && c < 0xfd
}
func DecodeCharSjis(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	var c = s[0]
	if c >= 0x81 && c <= 0x9f || c >= 0xe0 && c <= 0xfc {
		var next uint8
		if len(s) < 2 {
			return 0, 1, false
		}
		next = s[1]
		if SjisTrail(next) {
			char = uint(c)<<8 | uint(next)
		} else if SjisLead(next) {
			return 0, 1, false
		} else {
			return 0, 2, false
		}
		return char, 2, true
	} else if c < 0x80 || c >= 0xa1 && c <= 0xdf {
		return uint(c), 1, true
	} else {
		return 0, 1, false
	}
}

func DecodeCharEucjp(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}

	var c = s[0]
	if c >= 0xa1 && c <= 0xfe {
		if len(s) < 2 {
			return 0, 1, false
		}
		next := s[1]
		if next >= 0xa1 && next <= 0xfe {
			/* this a jis kanji char */
			char = uint(c)<<8 | uint(next)
		} else {
			return 0, lang.Cond(next != 0xa0 && next != 0xff, 1, 2), false
		}
		return char, 2, true
	} else if c == 0x8e {
		if len(s) < 2 {
			return 0, 1, false
		}
		next := s[1]
		if next >= 0xa1 && next <= 0xdf {
			/* JIS X 0201 kana */
			char = uint(c)<<8 | uint(next)
		} else {
			return 0, lang.Cond(next != 0xa0 && next != 0xff, 1, 2), false
		}
		return char, 2, true
	} else if c == 0x8f {
		var avail = len(s)
		if avail < 3 || !(s[1] >= 0xa1 && s[1] <= 0xfe) || !(s[2] >= 0xa1 && s[2] <= 0xfe) {
			if avail < 2 || s[1] != 0xa0 && s[1] != 0xff {
				return 0, 1, false
			} else if avail < 3 || s[2] != 0xa0 && s[2] != 0xff {
				return 0, 2, false
			} else {
				return 0, 3, false
			}
		} else {
			/* JIS X 0212 hojo-kanji */
			char = uint(c)<<16 | uint(s[1])<<8 | uint(s[2])
		}
		return char, 3, true
	} else if c != 0xa0 && c != 0xff {
		/* character encoded in 1 code unit */
		return uint(c), 1, true
	} else {
		return 0, 1, false
	}
}

func DecodeCharSingleByte(s string) (char uint, n int, ok bool) {
	if s == "" {
		return 0, 1, false
	}
	/* single-byte charsets */
	return uint(s[0]), 1, true
}

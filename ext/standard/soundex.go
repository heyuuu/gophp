package standard

import (
	"github.com/heyuuu/gophp/kits/ascii"
)

func ZifSoundex(str string) (string, bool) {
	// notice: PHP >= 8.0, 取消了空字符串输入返回 false 的特例
	if str == "" {
		return "", false
	}

	var soundex = [4]byte{'0', '0', '0', '0'}
	var soundexTable = []byte{
		0,   /* A */
		'1', /* B */
		'2', /* C */
		'3', /* D */
		0,   /* E */
		'1', /* F */
		'2', /* G */
		0,   /* H */
		0,   /* I */
		'2', /* J */
		'2', /* K */
		'4', /* L */
		'5', /* M */
		'5', /* N */
		0,   /* O */
		'1', /* P */
		'2', /* Q */
		'6', /* R */
		'2', /* S */
		'3', /* T */
		0,   /* U */
		'1', /* V */
		0,   /* W */
		'2', /* X */
		0,   /* Y */
		'2', /* Z */
	}

	/* build soundex string */
	var last byte = 0
	small := 0
	for _, code := range []byte(str) {
		code = ascii.ToUpper(code)
		if 'A' <= code && code <= 'Z' {
			if small == 0 {
				/* remember first valid char */
				soundex[small] = code
				small++
				last = soundexTable[code-'A']
			} else {
				/* ignore sequences of consonants with same soundex */
				code = soundexTable[code-'A']
				if code != last && code != 0 {
					soundex[small] = code
					small++
				}
				last = code
			}
			if small >= 4 {
				break
			}
		}
	}

	return string(soundex[:]), true
}

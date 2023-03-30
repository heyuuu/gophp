package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifSoundex(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var str *byte
	var i int
	var _small int
	var str_len int
	var code int
	var last int
	var soundex []byte
	var soundex_table []byte = []byte{0, '1', '2', '3', 0, '1', '2', 0, 0, '2', '2', '4', '5', '5', 0, '1', '2', '6', '2', '3', 0, '1', 0, '2', 0, '2'}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str, str_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if str_len == 0 {
		return_value.SetFalse()
		return
	}

	/* build soundex string */

	last = -1
	i = 0
	_small = 0
	for ; i < str_len && _small < 4; i++ {

		/* convert chars to upper case and strip non-letter chars */

		code = toupper(int(uint8(str[i])))
		if code >= 'A' && code <= 'Z' {
			if _small == 0 {

				/* remember first valid char */

				soundex[b.PostInc(&_small)] = byte(code)
				last = soundex_table[code-'A']
			} else {

				/* ignore sequences of consonants with same soundex */

				code = soundex_table[code-'A']
				if code != last {
					if code != 0 {
						soundex[b.PostInc(&_small)] = byte(code)
					}
					last = code
				}
			}
		}
	}

	/* pad with '0' and terminate with 0 ;-) */

	for _small < 4 {
		soundex[b.PostInc(&_small)] = '0'
	}
	soundex[_small] = '0'
	return_value.SetStringVal(b.CastStr(soundex, _small))
	return
}

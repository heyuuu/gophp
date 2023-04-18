package standard

import (
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpConvertCyrString(str *uint8, length int, from byte, to byte) *byte {
	var from_table *uint8
	var to_table *uint8
	var tmp uint8
	var i int
	from_table = nil
	to_table = nil
	switch toupper(int(uint8(from))) {
	case 'W':
		from_table = _cyrWin1251
	case 'A':
		fallthrough
	case 'D':
		from_table = _cyrCp866
	case 'I':
		from_table = _cyrIso88595
	case 'M':
		from_table = _cyrMac
	case 'K':

	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown source charset: %c", from)
	}
	switch toupper(int(uint8(to))) {
	case 'W':
		to_table = _cyrWin1251
	case 'A':
		fallthrough
	case 'D':
		to_table = _cyrCp866
	case 'I':
		to_table = _cyrIso88595
	case 'M':
		to_table = _cyrMac
	case 'K':

	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown destination charset: %c", to)
	}
	if str == nil {
		return (*byte)(str)
	}
	for i = 0; i < length; i++ {
		if from_table == nil {
			tmp = str[i]
		} else {
			tmp = from_table[str[i]]
		}
		if to_table == nil {
			str[i] = tmp
		} else {
			str[i] = to_table[tmp+256]
		}
	}
	return (*byte)(str)
}
func ZifConvertCyrString(executeData zpp.Ex, return_value zpp.Ret, str *types2.Zval, from *types2.Zval, to *types2.Zval) {
	var input *byte
	var fr_cs *byte
	var to_cs *byte
	var input_len int
	var fr_cs_len int
	var to_cs_len int
	var str *types2.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			input, input_len = fp.ParseString()
			fr_cs, fr_cs_len = fp.ParseString()
			to_cs, to_cs_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	str = types2.NewString(b.CastStr(input, input_len))
	PhpConvertCyrString((*uint8)(str.GetVal()), str.GetLen(), fr_cs[0], to_cs[0])
	return_value.SetString(str)
}

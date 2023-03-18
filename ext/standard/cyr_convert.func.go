// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
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
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown source charset: %c", from)
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
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown destination charset: %c", to)
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
func ZifConvertCyrString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *byte
	var fr_cs *byte
	var to_cs *byte
	var input_len int
	var fr_cs_len int
	var to_cs_len int
	var str *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input, input_len = fp.ParseString()
			fr_cs, fr_cs_len = fp.ParseString()
			to_cs, to_cs_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	str = types.ZendStringInit(input, input_len, 0)
	PhpConvertCyrString((*uint8)(str.GetVal()), str.GetLen(), fr_cs[0], to_cs[0])
	return_value.SetString(str)
}

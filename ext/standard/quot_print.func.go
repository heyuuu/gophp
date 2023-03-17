// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func PhpHex2int(c int) byte {
	if isdigit(c) {
		return c - '0'
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else {
		return -1
	}
}
func PhpQuotPrintDecode(str *uint8, length int, replace_us_by_ws int) *zend.ZendString {
	var i int
	var p1 *uint8
	var p2 *uint8
	var h_nbl uint
	var l_nbl uint
	var decoded_len int
	var buf_size int
	var retval *zend.ZendString
	var hexval_tbl []uint = []uint{64, 64, 64, 64, 64, 64, 64, 64, 64, 32, 16, 64, 64, 16, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 32, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 64, 64, 64, 64, 64, 64, 64, 10, 11, 12, 13, 14, 15, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 10, 11, 12, 13, 14, 15, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}
	if replace_us_by_ws != 0 {
		replace_us_by_ws = '_'
	}
	i = length
	p1 = str
	buf_size = length
	for i > 1 && (*p1) != '0' {
		if (*p1) == '=' {
			buf_size -= 2
			p1++
			i--
		}
		p1++
		i--
	}
	retval = zend.ZendStringAlloc(buf_size, 0)
	i = length
	p1 = str
	p2 = (*uint8)(retval.GetVal())
	decoded_len = 0
	for i > 0 && (*p1) != '0' {
		if (*p1) == '=' {
			i--
			p1++
			if i == 0 || (*p1) == '0' {
				break
			}
			h_nbl = hexval_tbl[*p1]
			if h_nbl < 16 {

				/* next char should be a hexadecimal digit */

				if b.PreDec(&i) == 0 || b.Assign(&l_nbl, hexval_tbl[*(b.PreInc(&p1))]) >= 16 {
					zend.Efree(retval)
					return nil
				}
				*(b.PostInc(&p2)) = h_nbl<<4 | l_nbl
				decoded_len++
				i--
				p1++
			} else if h_nbl < 64 {

				/* soft line break */

				for h_nbl == 32 {
					if b.PreDec(&i) == 0 || b.Assign(&h_nbl, hexval_tbl[*(b.PreInc(&p1))]) == 64 {
						zend.Efree(retval)
						return nil
					}
				}
				if p1[0] == '\r' && i >= 2 && p1[1] == '\n' {
					i--
					p1++
				}
				i--
				p1++
			} else {
				zend.Efree(retval)
				return nil
			}
		} else {
			if replace_us_by_ws == (*p1) {
				*(b.PostInc(&p2)) = 'x'
			} else {
				*(b.PostInc(&p2)) = *p1
			}
			i--
			p1++
			decoded_len++
		}
	}
	*p2 = '0'
	retval.SetLen(decoded_len)
	return retval
}
func PhpQuotPrintEncode(str *uint8, length int) *zend.ZendString {
	var lp zend.ZendUlong = 0
	var c uint8
	var d *uint8
	var hex *byte = "0123456789ABCDEF"
	var ret *zend.ZendString
	ret = zend.ZendStringSafeAlloc(3, length+(3*length/(PHP_QPRINT_MAXL-9)+1), 0, 0)
	d = (*uint8)(ret.GetVal())
	for b.PostDec(&length) {
		if b.Assign(&c, b.PostInc(&(*str))) == '0' && (*str) == '0' && length > 0 {
			b.PostInc(&(*d)) = '0'
			*str++
			b.PostInc(&(*d)) = (*str) - 1
			length--
			lp = 0
		} else {
			if iscntrl(c) || c == 0x7f || (c&0x80) != 0 || c == '=' || c == ' ' && (*str) == '0' {
				if b.AssignOp(&lp, "+=", 3) > PHP_QPRINT_MAXL && c <= 0x7f || c > 0x7f && c <= 0xdf && lp+3 > PHP_QPRINT_MAXL || c > 0xdf && c <= 0xef && lp+6 > PHP_QPRINT_MAXL || c > 0xef && c <= 0xf4 && lp+9 > PHP_QPRINT_MAXL {
					b.PostInc(&(*d)) = '='
					b.PostInc(&(*d)) = '0'
					b.PostInc(&(*d)) = '0'
					lp = 3
				}
				b.PostInc(&(*d)) = '='
				b.PostInc(&(*d)) = hex[c>>4]
				b.PostInc(&(*d)) = hex[c&0xf]
			} else {
				if b.PreInc(&lp) > PHP_QPRINT_MAXL {
					b.PostInc(&(*d)) = '='
					b.PostInc(&(*d)) = '0'
					b.PostInc(&(*d)) = '0'
					lp = 1
				}
				b.PostInc(&(*d)) = c
			}
		}
	}
	*d = '0'
	ret = zend.ZendStringTruncate(ret, d-(*uint8)(ret.GetVal()), 0)
	return ret
}
func ZifQuotedPrintableDecode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.ZendString
	var str_in *byte
	var str_out *zend.ZendString
	var i int = 0
	var j int = 0
	var k int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if arg1.GetLen() == 0 {

		/* shortcut */

		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
	str_in = arg1.GetVal()
	str_out = zend.ZendStringAlloc(arg1.GetLen(), 0)
	for str_in[i] {
		switch str_in[i] {
		case '=':
			if str_in[i+1] && str_in[i+2] && isxdigit(int(str_in[i+1])) && isxdigit(int(str_in[i+2])) {
				str_out.GetVal()[b.PostInc(&j)] = (PhpHex2int(int(str_in[i+1])) << 4) + PhpHex2int(int(str_in[i+2]))
				i += 3
			} else {
				k = 1
				for str_in[i+k] && (str_in[i+k] == 32 || str_in[i+k] == 9) {

					/* Possibly, skip spaces/tabs at the end of line */

					k++

					/* Possibly, skip spaces/tabs at the end of line */

				}
				if !(str_in[i+k]) {

					/* End of line reached */

					i += k

					/* End of line reached */

				} else if str_in[i+k] == 13 && str_in[i+k+1] == 10 {

					/* CRLF */

					i += k + 2

					/* CRLF */

				} else if str_in[i+k] == 13 || str_in[i+k] == 10 {

					/* CR or LF */

					i += k + 1

					/* CR or LF */

				} else {
					str_out.GetVal()[b.PostInc(&j)] = str_in[b.PostInc(&i)]
				}
			}
		default:
			str_out.GetVal()[b.PostInc(&j)] = str_in[b.PostInc(&i)]
		}
	}
	str_out.GetVal()[j] = '0'
	str_out.SetLen(j)
	return_value.SetString(str_out)
}
func ZifQuotedPrintableEncode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var new_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if str.GetLen() == 0 {
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
	new_str = PhpQuotPrintEncode((*uint8)(str.GetVal()), str.GetLen())
	return_value.SetString(new_str)
	return
}

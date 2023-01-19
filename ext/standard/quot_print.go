// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/quot_print.h>

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
   | Author: Kirill Maximov (kir@rus.net)                                 |
   +----------------------------------------------------------------------+
*/

// #define QUOT_PRINT_H

// Source: <ext/standard/quot_print.c>

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
   | Author: Kirill Maximov <kir@actimind.com>                            |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include < string . h >

// # include < errno . h >

// # include "php.h"

// # include "quot_print.h"

// # include < stdio . h >

/*
*  Converting HEX char to INT value
 */

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

/* }}} */

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
	p2 = (*uint8)(retval.val)
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

				if g.PreDec(&i) == 0 || g.Assign(&l_nbl, hexval_tbl[*(g.PreInc(&p1))]) >= 16 {
					zend._efree(retval)
					return nil
				}
				*(g.PostInc(&p2)) = h_nbl<<4 | l_nbl
				decoded_len++
				i--
				p1++
			} else if h_nbl < 64 {

				/* soft line break */

				for h_nbl == 32 {
					if g.PreDec(&i) == 0 || g.Assign(&h_nbl, hexval_tbl[*(g.PreInc(&p1))]) == 64 {
						zend._efree(retval)
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
				zend._efree(retval)
				return nil
			}
		} else {
			if replace_us_by_ws == (*p1) {
				*(g.PostInc(&p2)) = 'x'
			} else {
				*(g.PostInc(&p2)) = *p1
			}
			i--
			p1++
			decoded_len++
		}
	}
	*p2 = '0'
	retval.len_ = decoded_len
	return retval
}

/* }}} */

// #define PHP_QPRINT_MAXL       75

func PhpQuotPrintEncode(str *uint8, length int) *zend.ZendString {
	var lp zend.ZendUlong = 0
	var c uint8
	var d *uint8
	var hex *byte = "0123456789ABCDEF"
	var ret *zend.ZendString
	ret = zend.ZendStringSafeAlloc(3, length+(3*length/(75-9)+1), 0, 0)
	d = (*uint8)(ret.val)
	for g.PostDec(&length) {
		if g.Assign(&c, g.PostInc(&(*str))) == '0' && (*str) == '0' && length > 0 {
			g.PostInc(&(*d)) = '0'
			*str++
			g.PostInc(&(*d)) = (*str) - 1
			length--
			lp = 0
		} else {
			if iscntrl(c) || c == 0x7f || (c&0x80) != 0 || c == '=' || c == ' ' && (*str) == '0' {
				if g.AssignOp(&lp, "+=", 3) > 75 && c <= 0x7f || c > 0x7f && c <= 0xdf && lp+3 > 75 || c > 0xdf && c <= 0xef && lp+6 > 75 || c > 0xef && c <= 0xf4 && lp+9 > 75 {
					g.PostInc(&(*d)) = '='
					g.PostInc(&(*d)) = '0'
					g.PostInc(&(*d)) = '0'
					lp = 3
				}
				g.PostInc(&(*d)) = '='
				g.PostInc(&(*d)) = hex[c>>4]
				g.PostInc(&(*d)) = hex[c&0xf]
			} else {
				if g.PreInc(&lp) > 75 {
					g.PostInc(&(*d)) = '='
					g.PostInc(&(*d)) = '0'
					g.PostInc(&(*d)) = '0'
					lp = 1
				}
				g.PostInc(&(*d)) = c
			}
		}
	}
	*d = '0'
	ret = zend.ZendStringTruncate(ret, d-(*uint8)(ret.val), 0)
	return ret
}

/* }}} */

func ZifQuotedPrintableDecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
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

			if zend.ZendParseArgStr(_arg, &arg1, 0) == 0 {
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
	if arg1.len_ == 0 {

		/* shortcut */

		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	str_in = arg1.val
	str_out = zend.ZendStringAlloc(arg1.len_, 0)
	for str_in[i] {
		switch str_in[i] {
		case '=':
			if str_in[i+1] && str_in[i+2] && isxdigit(int(str_in[i+1])) && isxdigit(int(str_in[i+2])) {
				str_out.val[g.PostInc(&j)] = (PhpHex2int(int(str_in[i+1])) << 4) + PhpHex2int(int(str_in[i+2]))
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
					str_out.val[g.PostInc(&j)] = str_in[g.PostInc(&i)]
				}
			}
			break
		default:
			str_out.val[g.PostInc(&j)] = str_in[g.PostInc(&i)]
		}
	}
	str_out.val[j] = '0'
	str_out.len_ = j
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str_out
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
}

/* }}} */

func ZifQuotedPrintableEncode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var new_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	new_str = PhpQuotPrintEncode((*uint8)(str.val), str.len_)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = new_str
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

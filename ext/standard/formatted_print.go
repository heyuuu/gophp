// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/formatted_print.c>

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
   | Author: Stig S�ther Bakken <ssb@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include < math . h >

// # include "php.h"

// # include "ext/standard/head.h"

// # include "php_string.h"

// # include "zend_execute.h"

// # include < stdio . h >

// # include < locale . h >

// #define LCONV_DECIMAL_POINT       ( * lconv -> decimal_point )

// #define ALIGN_LEFT       0

// #define ALIGN_RIGHT       1

// #define ADJ_WIDTH       1

// #define ADJ_PRECISION       2

// #define NUM_BUF_SIZE       500

// #define FLOAT_PRECISION       6

// #define MAX_FLOAT_PRECISION       53

// #define PRINTF_DEBUG(arg)

var Hexchars []byte = "0123456789abcdef"
var HEXCHARS []byte = "0123456789ABCDEF"

/* php_spintf_appendchar() {{{ */

func PhpSprintfAppendchar(buffer **zend.ZendString, pos *int, add byte) {
	if (*pos)+1 >= (*buffer).len_ {
		*buffer = zend.ZendStringExtend(*buffer, (*buffer).len_<<1, 0)
	}
	(*buffer).val[g.PostInc(&(*pos))] = add
}

/* }}} */

func PhpSprintfAppendchars(buffer **zend.ZendString, pos *int, add *byte, len_ int) {
	if (*pos)+len_ >= (*buffer).len_ {
		var nlen int = (*buffer).len_
		for {
			nlen = nlen << 1
			if (*pos)+len_ < nlen {
				break
			}
		}
		*buffer = zend.ZendStringExtend(*buffer, nlen, 0)
	}
	memcpy((*buffer).val+(*pos), add, len_)
	*pos += len_
}

/* }}} */

func PhpSprintfAppendstring(buffer **zend.ZendString, pos *int, add *byte, min_width int, max_width int, padding byte, alignment int, len_ int, neg int, expprec int, always_sign int) {
	var npad int
	var req_size int
	var copy_len int
	var m_width int
	if expprec != 0 {
		if max_width < len_ {
			copy_len = max_width
		} else {
			copy_len = len_
		}
	} else {
		copy_len = len_
	}
	if min_width < copy_len {
		npad = 0
	} else {
		npad = min_width - copy_len
	}
	if min_width > copy_len {
		m_width = min_width
	} else {
		m_width = copy_len
	}
	if m_width > 2147483647-(*pos)-1 {
		zend.ZendErrorNoreturn(1<<0, "Field width %zd is too long", m_width)
	}
	req_size = (*pos) + m_width + 1
	if req_size > (*buffer).len_ {
		var size int = (*buffer).len_
		for req_size > size {
			if size > SIZE_MAX/2 {
				zend.ZendErrorNoreturn(1<<0, "Field width %zd is too long", req_size)
			}
			size <<= 1
		}
		*buffer = zend.ZendStringExtend(*buffer, size, 0)
	}
	if alignment == 1 {
		if (neg != 0 || always_sign != 0) && padding == '0' {
			if neg != 0 {
				(*buffer).val[g.PostInc(&(*pos))] = '-'
			} else {
				(*buffer).val[g.PostInc(&(*pos))] = '+'
			}
			add++
			len_--
			copy_len--
		}
		for g.PostDec(&npad) > 0 {
			(*buffer).val[g.PostInc(&(*pos))] = padding
		}
	}
	memcpy(&(*buffer).val[*pos], add, copy_len+1)
	*pos += copy_len
	if alignment == 0 {
		for g.PostDec(&npad) {
			(*buffer).val[g.PostInc(&(*pos))] = padding
		}
	}
}

/* }}} */

func PhpSprintfAppendint(buffer **zend.ZendString, pos *int, number zend.ZendLong, width int, padding byte, alignment int, always_sign int) {
	var numbuf []byte
	var magn zend.ZendUlong
	var nmagn zend.ZendUlong
	var i uint = 500 - 1
	var neg uint = 0
	if number < 0 {
		neg = 1
		magn = zend_ulong - (number + 1) + 1
	} else {
		magn = zend.ZendUlong(number)
	}

	/* Can't right-pad 0's on integers */

	if alignment == 0 && padding == '0' {
		padding = ' '
	}
	numbuf[i] = '0'
	for {
		nmagn = magn / 10
		numbuf[g.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 1) {
			break
		}
	}
	if neg != 0 {
		numbuf[g.PreDec(&i)] = '-'
	} else if always_sign != 0 {
		numbuf[g.PreDec(&i)] = '+'
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, 500-1-i, neg, 0, always_sign)
}

/* }}} */

func PhpSprintfAppenduint(buffer **zend.ZendString, pos *int, number zend.ZendUlong, width int, padding byte, alignment int) {
	var numbuf []byte
	var magn zend.ZendUlong
	var nmagn zend.ZendUlong
	var i uint = 500 - 1
	magn = zend.ZendUlong(number)

	/* Can't right-pad 0's on integers */

	if alignment == 0 && padding == '0' {
		padding = ' '
	}
	numbuf[i] = '0'
	for {
		nmagn = magn / 10
		numbuf[g.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 0) {
			break
		}
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, 500-1-i, 0, 0, 0)
}

/* }}} */

func PhpSprintfAppenddouble(buffer **zend.ZendString, pos *int, number float64, width int, padding byte, alignment int, precision int, adjust int, fmt byte, always_sign int) {
	var num_buf []byte
	var s *byte = nil
	var s_len int = 0
	var is_negative int = 0
	var lconv *__struct__lconv
	if (adjust & 2) == 0 {
		precision = 6
	} else if precision > 53 {
		core.PhpErrorDocref(nil, 1<<3, "Requested precision of %d digits was truncated to PHP maximum of %d digits", precision, 53)
		precision = 53
	}
	if isnan(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buffer, pos, "NaN", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	if isinf(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buffer, pos, "INF", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	switch fmt {
	case 'e':

	case 'E':

	case 'f':

	case 'F':
		lconv = localeconv()
		s = core.PhpConvFp(g.Cond(fmt == 'f', 'F', fmt), number, 0, precision, g.CondF1(fmt == 'f', func() __auto__ { return (*lconv).decimal_point }, '.'), &is_negative, &num_buf[1], &s_len)
		if is_negative != 0 {
			num_buf[0] = '-'
			s = num_buf
			s_len++
		} else if always_sign != 0 {
			num_buf[0] = '+'
			s = num_buf
			s_len++
		}
		break
	case 'g':

	case 'G':
		if precision == 0 {
			precision = 1
		}

		/*
		 * * We use &num_buf[ 1 ], so that we have room for the sign
		 */

		lconv = localeconv()
		s = core.PhpGcvt(number, precision, (*lconv).decimal_point, g.Cond(fmt == 'G', 'E', 'e'), &num_buf[1])
		is_negative = 0
		if (*s) == '-' {
			is_negative = 1
			s = &num_buf[1]
		} else if always_sign != 0 {
			num_buf[0] = '+'
			s = num_buf
		}
		s_len = strlen(s)
		break
	}
	PhpSprintfAppendstring(buffer, pos, s, width, 0, padding, alignment, s_len, is_negative, 0, always_sign)
}

/* }}} */

func PhpSprintfAppend2n(buffer **zend.ZendString, pos *int, number zend.ZendLong, width int, padding byte, alignment int, n int, chartable *byte, expprec int) {
	var numbuf []byte
	var num zend.ZendUlong
	var i zend.ZendUlong = 500 - 1
	var andbits int = (1 << n) - 1
	num = zend.ZendUlong(number)
	numbuf[i] = '0'
	for {
		numbuf[g.PreDec(&i)] = chartable[num&andbits]
		num >>= n
		if num <= 0 {
			break
		}
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, 500-1-i, 0, expprec, 0)
}

/* }}} */

func PhpSprintfGetnumber(buffer **byte, len_ *int) int {
	var endptr *byte
	var num zend.ZendLong = strtoll(*buffer, &endptr, 10)
	var i int
	if endptr != nil {
		i = endptr - (*buffer)
		*len_ -= i
		*buffer = endptr
	}
	if num >= 2147483647 || num < 0 {
		return -1
	} else {
		return int(num)
	}
}

/* }}} */

func PhpFormattedPrint(z_format *zend.Zval, args *zend.Zval, argc int) *zend.ZendString {
	var size int = 240
	var outpos int = 0
	var alignment int
	var currarg int
	var adjusting int
	var argnum int
	var width int
	var precision int
	var format *byte
	var temppos *byte
	var padding byte
	var result *zend.ZendString
	var always_sign int
	var format_len int
	if zend.TryConvertToString(z_format) == 0 {
		return nil
	}
	format = z_format.value.str.val
	format_len = z_format.value.str.len_
	result = zend.ZendStringAlloc(size, 0)
	currarg = 0
	for format_len != 0 {
		var expprec int
		var tmp *zend.Zval
		temppos = memchr(format, '%', format_len)
		if temppos == nil {
			PhpSprintfAppendchars(&result, &outpos, format, format_len)
			break
		} else if temppos != format {
			PhpSprintfAppendchars(&result, &outpos, format, temppos-format)
			format_len -= temppos - format
			format = temppos
		}
		format++
		format_len--
		if (*format) == '%' {
			PhpSprintfAppendchar(&result, &outpos, '%')
			format++
			format_len--
		} else {

			/* starting a new format specifier, reset variables */

			alignment = 1
			adjusting = 0
			padding = ' '
			always_sign = 0
			expprec = 0
			if isalpha(int(*format)) {
				precision = 0
				width = precision
				currarg++
				argnum = currarg - 1
			} else {

				/* first look for argnum */

				temppos = format
				for isdigit(int(*temppos)) {
					temppos++
				}
				if (*temppos) == '$' {
					argnum = PhpSprintfGetnumber(&format, &format_len)
					if argnum <= 0 {
						zend.ZendStringEfree(result)
						core.PhpErrorDocref(nil, 1<<1, "Argument number must be greater than zero")
						return nil
					}
					argnum--
					format++
					format_len--
				} else {
					currarg++
					argnum = currarg - 1
				}

				/* after argnum comes modifiers */

				for {
					if (*format) == ' ' || (*format) == '0' {
						padding = *format
					} else if (*format) == '-' {
						alignment = 0
					} else if (*format) == '+' {
						always_sign = 1
					} else if (*format) == '\'' && format_len > 1 {
						format++
						format_len--
						padding = *format
					} else {
						break
					}
					format++
					format_len--
				}

				/* after modifiers comes width */

				if isdigit(int(*format)) {
					if g.Assign(&width, PhpSprintfGetnumber(&format, &format_len)) < 0 {
						zend._efree(result)
						core.PhpErrorDocref(nil, 1<<1, "Width must be greater than zero and less than %d", 2147483647)
						return nil
					}
					adjusting |= 1
				} else {
					width = 0
				}

				/* after width and argnum comes precision */

				if (*format) == '.' {
					format++
					format_len--
					if isdigit(int(*format)) {
						if g.Assign(&precision, PhpSprintfGetnumber(&format, &format_len)) < 0 {
							zend._efree(result)
							core.PhpErrorDocref(nil, 1<<1, "Precision must be greater than zero and less than %d", 2147483647)
							return nil
						}
						adjusting |= 2
						expprec = 1
					} else {
						precision = 0
					}
				} else {
					precision = 0
				}
			}
			if argnum >= argc {
				zend._efree(result)
				core.PhpErrorDocref(nil, 1<<1, "Too few arguments")
				return nil
			}
			if (*format) == 'l' {
				format++
				format_len--
			}

			/* now we expect to find a type specifier */

			tmp = &args[argnum]
			switch *format {
			case 's':
				var t *zend.ZendString
				var str *zend.ZendString = zend.ZvalGetTmpString(tmp, &t)
				PhpSprintfAppendstring(&result, &outpos, str.val, width, precision, padding, alignment, str.len_, 0, expprec, 0)
				zend.ZendTmpStringRelease(t)
				break
			case 'd':
				PhpSprintfAppendint(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, always_sign)
				break
			case 'u':
				PhpSprintfAppenduint(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment)
				break
			case 'g':

			case 'G':

			case 'e':

			case 'E':

			case 'f':

			case 'F':
				PhpSprintfAppenddouble(&result, &outpos, zend.ZvalGetDouble(tmp), width, padding, alignment, precision, adjusting, *format, always_sign)
				break
			case 'c':
				PhpSprintfAppendchar(&result, &outpos, byte(zend.ZvalGetLong(tmp)))
				break
			case 'o':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 3, Hexchars, expprec)
				break
			case 'x':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 4, Hexchars, expprec)
				break
			case 'X':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 4, HEXCHARS, expprec)
				break
			case 'b':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 1, Hexchars, expprec)
				break
			case '%':
				PhpSprintfAppendchar(&result, &outpos, '%')
				break
			case '0':
				if format_len == 0 {
					goto exit
				}
				break
			default:
				break
			}
			format++
			format_len--
		}
	}
exit:

	/* possibly, we have to make sure we have room for the terminating null? */

	result.val[outpos] = 0
	result.len_ = outpos
	return result
}

/* }}} */

func PhpFormattedPrintGetArray(array *zend.Zval, argc *int) *zend.Zval {
	var args *zend.Zval
	var zv *zend.Zval
	var n int
	if array.u1.v.type_ != 7 {
		zend.ConvertToArray(array)
	}
	n = array.value.arr.nNumOfElements
	args = (*zend.Zval)(zend._safeEmalloc(n, g.SizeOf("zval"), 0))
	n = 0
	for {
		var __ht *zend.HashTable = array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			zv = _z
			var _z1 *zend.Zval = &args[n]
			var _z2 *zend.Zval = zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			n++
		}
		break
	}
	*argc = n
	return args
}

/* }}} */

func ZifUserSprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var format *zend.Zval
	var args *zend.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func ZifVsprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var format *zend.Zval
	var array *zend.Zval
	var args *zend.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend._efree(args)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func ZifUserPrintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var rlen int
	var format *zend.Zval
	var args *zend.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	rlen = core.PhpOutputWrite(result.val, result.len_)
	zend.ZendStringEfree(result)
	var __z *zend.Zval = return_value
	__z.value.lval = rlen
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifVprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var rlen int
	var format *zend.Zval
	var array *zend.Zval
	var args *zend.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend._efree(args)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	rlen = core.PhpOutputWrite(result.val, result.len_)
	zend.ZendStringEfree(result)
	var __z *zend.Zval = return_value
	__z.value.lval = rlen
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifFprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var arg1 *zend.Zval
	var format *zend.Zval
	var args *zend.Zval
	var argc int
	var result *zend.ZendString
	if execute_data.This.u2.num_args < 2 {
		zend.ZendWrongParamCount()
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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

			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(arg1, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	streams._phpStreamWrite(stream, result.val, result.len_)
	var __z *zend.Zval = return_value
	__z.value.lval = result.len_
	__z.u1.type_info = 4
	zend.ZendStringEfree(result)
}

/* }}} */

func ZifVfprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var arg1 *zend.Zval
	var format *zend.Zval
	var array *zend.Zval
	var args *zend.Zval
	var argc int
	var result *zend.ZendString
	if execute_data.This.u2.num_args != 3 {
		zend.ZendWrongParamCount()
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
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

			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
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

			zend.ZendParseArgZvalDeref(_arg, &format, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(arg1, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend._efree(args)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	streams._phpStreamWrite(stream, result.val, result.len_)
	var __z *zend.Zval = return_value
	__z.value.lval = result.len_
	__z.u1.type_info = 4
	zend.ZendStringEfree(result)
}

/* }}} */

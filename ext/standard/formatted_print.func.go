// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
)

func PhpSprintfAppendchar(buffer **zend.ZendString, pos *int, add byte) {
	if (*pos)+1 >= zend.ZSTR_LEN(*buffer) {
		*buffer = zend.ZendStringExtend(*buffer, zend.ZSTR_LEN(*buffer)<<1, 0)
	}
	zend.ZSTR_VAL(*buffer)[b.PostInc(&(*pos))] = add
}
func PhpSprintfAppendchars(buffer **zend.ZendString, pos *int, add *byte, len_ int) {
	if (*pos)+len_ >= zend.ZSTR_LEN(*buffer) {
		var nlen int = zend.ZSTR_LEN(*buffer)
		for {
			nlen = nlen << 1
			if (*pos)+len_ < nlen {
				break
			}
		}
		*buffer = zend.ZendStringExtend(*buffer, nlen, 0)
	}
	memcpy(zend.ZSTR_VAL(*buffer)+(*pos), add, len_)
	*pos += len_
}
func PhpSprintfAppendstring(buffer **zend.ZendString, pos *int, add *byte, min_width int, max_width int, padding byte, alignment int, len_ int, neg int, expprec int, always_sign int) {
	var npad int
	var req_size int
	var copy_len int
	var m_width int
	if expprec != 0 {
		copy_len = cli.MIN(max_width, len_)
	} else {
		copy_len = len_
	}
	if min_width < copy_len {
		npad = 0
	} else {
		npad = min_width - copy_len
	}
	m_width = zend.MAX(min_width, copy_len)
	if m_width > core.INT_MAX-(*pos)-1 {
		zend.ZendErrorNoreturn(zend.E_ERROR, "Field width %zd is too long", m_width)
	}
	req_size = (*pos) + m_width + 1
	if req_size > zend.ZSTR_LEN(*buffer) {
		var size int = zend.ZSTR_LEN(*buffer)
		for req_size > size {
			if size > zend.ZEND_SIZE_MAX/2 {
				zend.ZendErrorNoreturn(zend.E_ERROR, "Field width %zd is too long", req_size)
			}
			size <<= 1
		}
		*buffer = zend.ZendStringExtend(*buffer, size, 0)
	}
	if alignment == ALIGN_RIGHT {
		if (neg != 0 || always_sign != 0) && padding == '0' {
			if neg != 0 {
				zend.ZSTR_VAL(*buffer)[b.PostInc(&(*pos))] = '-'
			} else {
				zend.ZSTR_VAL(*buffer)[b.PostInc(&(*pos))] = '+'
			}
			add++
			len_--
			copy_len--
		}
		for b.PostDec(&npad) > 0 {
			zend.ZSTR_VAL(*buffer)[b.PostInc(&(*pos))] = padding
		}
	}
	memcpy(&zend.ZSTR_VAL(*buffer)[*pos], add, copy_len+1)
	*pos += copy_len
	if alignment == ALIGN_LEFT {
		for b.PostDec(&npad) {
			zend.ZSTR_VAL(*buffer)[b.PostInc(&(*pos))] = padding
		}
	}
}
func PhpSprintfAppendint(buffer **zend.ZendString, pos *int, number zend.ZendLong, width int, padding byte, alignment int, always_sign int) {
	var numbuf []byte
	var magn zend.ZendUlong
	var nmagn zend.ZendUlong
	var i uint = NUM_BUF_SIZE - 1
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
		numbuf[b.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 1) {
			break
		}
	}
	if neg != 0 {
		numbuf[b.PreDec(&i)] = '-'
	} else if always_sign != 0 {
		numbuf[b.PreDec(&i)] = '+'
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, neg, 0, always_sign)
}
func PhpSprintfAppenduint(buffer **zend.ZendString, pos *int, number zend.ZendUlong, width int, padding byte, alignment int) {
	var numbuf []byte
	var magn zend.ZendUlong
	var nmagn zend.ZendUlong
	var i uint = NUM_BUF_SIZE - 1
	magn = zend.ZendUlong(number)

	/* Can't right-pad 0's on integers */

	if alignment == 0 && padding == '0' {
		padding = ' '
	}
	numbuf[i] = '0'
	for {
		nmagn = magn / 10
		numbuf[b.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 0) {
			break
		}
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, 0, 0, 0)
}
func PhpSprintfAppenddouble(buffer **zend.ZendString, pos *int, number float64, width int, padding byte, alignment int, precision int, adjust int, fmt byte, always_sign int) {
	var num_buf []byte
	var s *byte = nil
	var s_len int = 0
	var is_negative int = 0
	var lconv *__struct__lconv
	if (adjust & ADJ_PRECISION) == 0 {
		precision = FLOAT_PRECISION
	} else if precision > MAX_FLOAT_PRECISION {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "Requested precision of %d digits was truncated to PHP maximum of %d digits", precision, MAX_FLOAT_PRECISION)
		precision = MAX_FLOAT_PRECISION
	}
	if core.ZendIsnan(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buffer, pos, "NaN", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	if core.ZendIsinf(number) {
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
		s = core.PhpConvFp(b.Cond(fmt == 'f', 'F', fmt), number, 0, precision, b.Cond(fmt == 'f', LCONV_DECIMAL_POINT, '.'), &is_negative, &num_buf[1], &s_len)
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
		s = core.PhpGcvt(number, precision, LCONV_DECIMAL_POINT, b.Cond(fmt == 'G', 'E', 'e'), &num_buf[1])
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
func PhpSprintfAppend2n(buffer **zend.ZendString, pos *int, number zend.ZendLong, width int, padding byte, alignment int, n int, chartable *byte, expprec int) {
	var numbuf []byte
	var num zend.ZendUlong
	var i zend.ZendUlong = NUM_BUF_SIZE - 1
	var andbits int = (1 << n) - 1
	num = zend.ZendUlong(number)
	numbuf[i] = '0'
	for {
		numbuf[b.PreDec(&i)] = chartable[num&andbits]
		num >>= n
		if num <= 0 {
			break
		}
	}
	PhpSprintfAppendstring(buffer, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, 0, expprec, 0)
}
func PhpSprintfGetnumber(buffer **byte, len_ *int) int {
	var endptr *byte
	var num zend.ZendLong = zend.ZEND_STRTOL(*buffer, &endptr, 10)
	var i int
	if endptr != nil {
		i = endptr - (*buffer)
		*len_ -= i
		*buffer = endptr
	}
	if num >= core.INT_MAX || num < 0 {
		return -1
	} else {
		return int(num)
	}
}
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
	format = zend.Z_STRVAL_P(z_format)
	format_len = zend.Z_STRLEN_P(z_format)
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

			alignment = ALIGN_RIGHT
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
						core.PhpErrorDocref(nil, zend.E_WARNING, "Argument number must be greater than zero")
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
						alignment = ALIGN_LEFT
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
					if b.Assign(&width, PhpSprintfGetnumber(&format, &format_len)) < 0 {
						zend.Efree(result)
						core.PhpErrorDocref(nil, zend.E_WARNING, "Width must be greater than zero and less than %d", core.INT_MAX)
						return nil
					}
					adjusting |= ADJ_WIDTH
				} else {
					width = 0
				}

				/* after width and argnum comes precision */

				if (*format) == '.' {
					format++
					format_len--
					if isdigit(int(*format)) {
						if b.Assign(&precision, PhpSprintfGetnumber(&format, &format_len)) < 0 {
							zend.Efree(result)
							core.PhpErrorDocref(nil, zend.E_WARNING, "Precision must be greater than zero and less than %d", core.INT_MAX)
							return nil
						}
						adjusting |= ADJ_PRECISION
						expprec = 1
					} else {
						precision = 0
					}
				} else {
					precision = 0
				}
			}
			if argnum >= argc {
				zend.Efree(result)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Too few arguments")
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
				PhpSprintfAppendstring(&result, &outpos, zend.ZSTR_VAL(str), width, precision, padding, alignment, zend.ZSTR_LEN(str), 0, expprec, 0)
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

	zend.ZSTR_VAL(result)[outpos] = 0
	zend.ZSTR_LEN(result) = outpos
	return result
}
func PhpFormattedPrintGetArray(array *zend.Zval, argc *int) *zend.Zval {
	var args *zend.Zval
	var zv *zend.Zval
	var n int
	if zend.Z_TYPE_P(array) != zend.IS_ARRAY {
		zend.ConvertToArray(array)
	}
	n = zend.Z_ARRVAL_P(array).NumElements()
	args = (*zend.Zval)(zend.SafeEmalloc(n, b.SizeOf("zval"), 0))
	n = 0
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.GetVal()

			if zend.Z_TYPE_P(_z) == zend.IS_UNDEF {
				continue
			}
			zv = _z
			zend.ZVAL_COPY_VALUE(&args[n], zv)
			n++
		}
		break
	}
	*argc = n
	return args
}
func ZifUserSprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var format *zend.Zval
	var args *zend.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(result)
}
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
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(result)
}
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
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	rlen = core.PHPWRITE(zend.ZSTR_VAL(result), zend.ZSTR_LEN(result))
	zend.ZendStringEfree(result)
	zend.RETVAL_LONG(rlen)
	return
}
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
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	rlen = core.PHPWRITE(zend.ZSTR_VAL(result), zend.ZSTR_LEN(result))
	zend.ZendStringEfree(result)
	zend.RETVAL_LONG(rlen)
	return
}
func ZifFprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var arg1 *zend.Zval
	var format *zend.Zval
	var args *zend.Zval
	var argc int
	var result *zend.ZendString
	if zend.ZEND_NUM_ARGS() < 2 {
		zend.WRONG_PARAM_COUNT
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamWrite(stream, zend.ZSTR_VAL(result), zend.ZSTR_LEN(result))
	zend.RETVAL_LONG(zend.ZSTR_LEN(result))
	zend.ZendStringEfree(result)
}
func ZifVfprintf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var arg1 *zend.Zval
	var format *zend.Zval
	var array *zend.Zval
	var args *zend.Zval
	var argc int
	var result *zend.ZendString
	if zend.ZEND_NUM_ARGS() != 3 {
		zend.WRONG_PARAM_COUNT
	}
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &format, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &array, 0)
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamWrite(stream, zend.ZSTR_VAL(result), zend.ZSTR_LEN(result))
	zend.RETVAL_LONG(zend.ZSTR_LEN(result))
	zend.ZendStringEfree(result)
}

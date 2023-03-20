// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func PhpSprintfAppendchar(buffer **types.String, pos *int, add byte) {
	if (*pos)+1 >= buffer.GetLen() {
		*buffer = types.ZendStringExtend(*buffer, buffer.GetLen()<<1, 0)
	}
	buffer.GetVal()[b.PostInc(&(*pos))] = add
}
func PhpSprintfAppendchars(buffer **types.String, pos *int, add *byte, len_ int) {
	if (*pos)+len_ >= buffer.GetLen() {
		var nlen int = buffer.GetLen()
		for {
			nlen = nlen << 1
			if (*pos)+len_ < nlen {
				break
			}
		}
		*buffer = types.ZendStringExtend(*buffer, nlen, 0)
	}
	memcpy(buffer.GetVal()+(*pos), add, len_)
	*pos += len_
}
func PhpSprintfAppendstring(
	buffer **types.String,
	pos *int,
	add *byte,
	min_width int,
	max_width int,
	padding byte,
	alignment int,
	len_ int,
	neg int,
	expprec int,
	always_sign int,
) {
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
	m_width = b.Max(min_width, copy_len)
	if m_width > core.INT_MAX-(*pos)-1 {
		faults.ErrorNoreturn(faults.E_ERROR, "Field width %zd is too long", m_width)
	}
	req_size = (*pos) + m_width + 1
	if req_size > buffer.GetLen() {
		var size int = buffer.GetLen()
		for req_size > size {
			if size > types.ZEND_SIZE_MAX/2 {
				faults.ErrorNoreturn(faults.E_ERROR, "Field width %zd is too long", req_size)
			}
			size <<= 1
		}
		*buffer = types.ZendStringExtend(*buffer, size, 0)
	}
	if alignment == ALIGN_RIGHT {
		if (neg != 0 || always_sign != 0) && padding == '0' {
			if neg != 0 {
				buffer.GetVal()[b.PostInc(&(*pos))] = '-'
			} else {
				buffer.GetVal()[b.PostInc(&(*pos))] = '+'
			}
			add++
			len_--
			copy_len--
		}
		for b.PostDec(&npad) > 0 {
			buffer.GetVal()[b.PostInc(&(*pos))] = padding
		}
	}
	memcpy(&buffer.GetVal()[*pos], add, copy_len+1)
	*pos += copy_len
	if alignment == ALIGN_LEFT {
		for b.PostDec(&npad) {
			buffer.GetVal()[b.PostInc(&(*pos))] = padding
		}
	}
}
func PhpSprintfAppendint(
	buffer **types.String,
	pos *int,
	number zend.ZendLong,
	width int,
	padding byte,
	alignment int,
	always_sign int,
) {
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
func PhpSprintfAppenduint(
	buffer **types.String,
	pos *int,
	number zend.ZendUlong,
	width int,
	padding byte,
	alignment int,
) {
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
func PhpSprintfAppenddouble(
	buffer **types.String,
	pos *int,
	number float64,
	width int,
	padding byte,
	alignment int,
	precision int,
	adjust int,
	fmt byte,
	always_sign int,
) {
	var num_buf []byte
	var s *byte = nil
	var s_len int = 0
	var is_negative int = 0
	var lconv *__struct__lconv
	if (adjust & ADJ_PRECISION) == 0 {
		precision = FLOAT_PRECISION
	} else if precision > MAX_FLOAT_PRECISION {
		core.PhpErrorDocref(nil, faults.E_NOTICE, "Requested precision of %d digits was truncated to PHP maximum of %d digits", precision, MAX_FLOAT_PRECISION)
		precision = MAX_FLOAT_PRECISION
	}
	if core.ZendIsNaN(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buffer, pos, "NaN", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	if core.ZendIsInf(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buffer, pos, "INF", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	switch fmt {
	case 'e':
		fallthrough
	case 'E':
		fallthrough
	case 'f':
		fallthrough
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
	case 'g':
		fallthrough
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
	}
	PhpSprintfAppendstring(buffer, pos, s, width, 0, padding, alignment, s_len, is_negative, 0, always_sign)
}
func PhpSprintfAppend2n(
	buffer **types.String,
	pos *int,
	number zend.ZendLong,
	width int,
	padding byte,
	alignment int,
	n int,
	chartable *byte,
	expprec int,
) {
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
func PhpFormattedPrint(z_format *types.Zval, args *types.Zval, argc int) *types.String {
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
	var result *types.String
	var always_sign int
	var format_len int
	if zend.TryConvertToString(z_format) == 0 {
		return nil
	}
	format = z_format.GetStr().GetVal()
	format_len = z_format.GetStr().GetLen()
	result = types.ZendStringAlloc(size, 0)
	currarg = 0
	for format_len != 0 {
		var expprec int
		var tmp *types.Zval
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
						types.ZendStringEfree(result)
						core.PhpErrorDocref(nil, faults.E_WARNING, "Argument number must be greater than zero")
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
						core.PhpErrorDocref(nil, faults.E_WARNING, "Width must be greater than zero and less than %d", core.INT_MAX)
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
							core.PhpErrorDocref(nil, faults.E_WARNING, "Precision must be greater than zero and less than %d", core.INT_MAX)
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
				core.PhpErrorDocref(nil, faults.E_WARNING, "Too few arguments")
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
				var t *types.String
				var str *types.String = zend.ZvalGetTmpString(tmp, &t)
				PhpSprintfAppendstring(&result, &outpos, str.GetVal(), width, precision, padding, alignment, str.GetLen(), 0, expprec, 0)
				zend.ZendTmpStringRelease(t)
			case 'd':
				PhpSprintfAppendint(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, always_sign)
			case 'u':
				PhpSprintfAppenduint(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment)
			case 'g':
				fallthrough
			case 'G':
				fallthrough
			case 'e':
				fallthrough
			case 'E':
				fallthrough
			case 'f':
				fallthrough
			case 'F':
				PhpSprintfAppenddouble(&result, &outpos, zend.ZvalGetDouble(tmp), width, padding, alignment, precision, adjusting, *format, always_sign)
			case 'c':
				PhpSprintfAppendchar(&result, &outpos, byte(zend.ZvalGetLong(tmp)))
			case 'o':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 3, Hexchars, expprec)
			case 'x':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 4, Hexchars, expprec)
			case 'X':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 4, HEXCHARS, expprec)
			case 'b':
				PhpSprintfAppend2n(&result, &outpos, zend.ZvalGetLong(tmp), width, padding, alignment, 1, Hexchars, expprec)
			case '%':
				PhpSprintfAppendchar(&result, &outpos, '%')
			case '0':
				if format_len == 0 {
					goto exit
				}
			default:

			}
			format++
			format_len--
		}
	}
exit:

	/* possibly, we have to make sure we have room for the terminating null? */

	result.GetVal()[outpos] = 0
	result.SetLen(outpos)
	return result
}
func PhpFormattedPrintGetArray(array *types.Zval, argc *int) *types.Zval {
	var args *types.Zval
	var zv *types.Zval
	var n int
	if array.GetType() != types.IS_ARRAY {
		zend.ConvertToArray(array)
	}
	n = types.Z_ARRVAL_P(array).GetNNumOfElements()
	args = (*types.Zval)(zend.SafeEmalloc(n, b.SizeOf("zval"), 0))
	n = 0
	var __ht *types.Array = array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		zv = _z
		types.ZVAL_COPY_VALUE(&args[n], zv)
		n++
	}
	*argc = n
	return args
}
func ZifUserSprintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var result *types.String
	var format *types.Zval
	var args *types.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			format = fp.ParseZval()
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetString(result)
}
func ZifVsprintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var result *types.String
	var format *types.Zval
	var array *types.Zval
	var args *types.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetString(result)
}
func ZifUserPrintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var result *types.String
	var rlen int
	var format *types.Zval
	var args *types.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			format = fp.ParseZval()
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.SetFalse()
		return
	}
	rlen = core.PHPWRITE(result.GetVal(), result.GetLen())
	types.ZendStringEfree(result)
	return_value.SetLong(rlen)
	return
}
func ZifVprintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var result *types.String
	var rlen int
	var format *types.Zval
	var array *types.Zval
	var args *types.Zval
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		return_value.SetFalse()
		return
	}
	rlen = core.PHPWRITE(result.GetVal(), result.GetLen())
	types.ZendStringEfree(result)
	return_value.SetLong(rlen)
	return
}
func ZifFprintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stream *core.PhpStream
	var arg1 *types.Zval
	var format *types.Zval
	var args *types.Zval
	var argc int
	var result *types.String
	if executeData.NumArgs() < 2 {
		zend.ZendWrongParamCount()
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg1 = fp.ParseResource()
			format = fp.ParseZval()
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return_value.SetFalse()
		return
	}
	core.PhpStreamWrite(stream, result.GetVal(), result.GetLen())
	return_value.SetLong(result.GetLen())
	types.ZendStringEfree(result)
}
func ZifVfprintf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stream *core.PhpStream
	var arg1 *types.Zval
	var format *types.Zval
	var array *types.Zval
	var args *types.Zval
	var argc int
	var result *types.String
	if executeData.NumArgs() != 3 {
		zend.ZendWrongParamCount()
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg1 = fp.ParseResource()
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	args = PhpFormattedPrintGetArray(array, &argc)
	result = PhpFormattedPrint(format, args, argc)
	zend.Efree(args)
	if result == nil {
		return_value.SetFalse()
		return
	}
	core.PhpStreamWrite(stream, result.GetVal(), result.GetLen())
	return_value.SetLong(result.GetLen())
	types.ZendStringEfree(result)
}

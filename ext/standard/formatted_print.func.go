package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func PhpSprintfAppendstring(
	buf *strings.Builder,
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
	if alignment == ALIGN_RIGHT {
		if (neg != 0 || always_sign != 0) && padding == '0' {
			if neg != 0 {
				*pos++
				buf.WriteByte('-')
			} else {
				*pos++
				buf.WriteByte('+')
			}
			add++
			len_--
			copy_len--
		}
		for lang.PostDec(&npad) > 0 {
			*pos++
			buf.WriteByte(padding)
		}
	}
	buf.WriteString(add[:copy_len])
	*pos += copy_len
	if alignment == ALIGN_LEFT {
		for lang.PostDec(&npad) != 0 {
			*pos++
			buf.WriteByte(padding)
		}
	}
}
func PhpSprintfAppendint(
	buf *strings.Builder,
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
		numbuf[lang.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 1) {
			break
		}
	}
	if neg != 0 {
		numbuf[lang.PreDec(&i)] = '-'
	} else if always_sign != 0 {
		numbuf[lang.PreDec(&i)] = '+'
	}
	PhpSprintfAppendstring(buf, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, neg, 0, always_sign)
}
func PhpSprintfAppenduint(
	buf *strings.Builder,
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
		numbuf[lang.PreDec(&i)] = uint8(magn - nmagn*10 + '0')
		magn = nmagn
		if !(magn > 0 && i > 0) {
			break
		}
	}
	PhpSprintfAppendstring(buf, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, 0, 0, 0)
}
func PhpSprintfAppenddouble(
	buf *strings.Builder,
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
		core.PhpErrorDocref("", faults.E_NOTICE, "Requested precision of %d digits was truncated to PHP maximum of %d digits", precision, MAX_FLOAT_PRECISION)
		precision = MAX_FLOAT_PRECISION
	}
	if core.ZendIsNaN(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buf, pos, "NaN", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
		return
	}
	if core.ZendIsInf(number) {
		is_negative = number < 0
		PhpSprintfAppendstring(buf, pos, "INF", 3, 0, padding, alignment, 3, is_negative, 0, always_sign)
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
		s = core.PhpConvFp(lang.Cond(fmt == 'f', 'F', fmt), number, 0, precision, lang.Cond(fmt == 'f', LCONV_DECIMAL_POINT, '.'), &is_negative, &num_buf[1], &s_len)
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
		s = core.PhpGcvt(number, precision, LCONV_DECIMAL_POINT, lang.Cond(fmt == 'G', 'E', 'e'), &num_buf[1])
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
	PhpSprintfAppendstring(buf, pos, s, width, 0, padding, alignment, s_len, is_negative, 0, always_sign)
}
func PhpSprintfAppend2n(
	buf *strings.Builder,
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
		numbuf[lang.PreDec(&i)] = chartable[num&andbits]
		num >>= n
		if num <= 0 {
			break
		}
	}
	PhpSprintfAppendstring(buf, pos, &numbuf[i], width, 0, padding, alignment, NUM_BUF_SIZE-1-i, 0, expprec, 0)
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
	if operators.TryConvertToString(z_format) == 0 {
		return nil
	}
	var buf strings.Builder
	format = z_format.StringEx().GetVal()
	format_len = z_format.StringEx().GetLen()
	currarg = 0
	for format_len != 0 {
		var expprec int
		var tmp *types.Zval
		temppos = memchr(format, '%', format_len)
		if temppos == nil {
			buf.WriteString(b.CastStr(format, format_len))
			break
		} else if temppos != format {
			buf.WriteString(b.CastStr(format, temppos-format))
			format_len -= temppos - format
			format = temppos
		}
		format++
		format_len--
		if (*format) == '%' {
			buf.WriteByte('%')
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
						core.PhpErrorDocref("", faults.E_WARNING, "Argument number must be greater than zero")
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
					if lang.Assign(&width, PhpSprintfGetnumber(&format, &format_len)) < 0 {
						core.PhpErrorDocref("", faults.E_WARNING, "Width must be greater than zero and less than %d", core.INT_MAX)
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
						if lang.Assign(&precision, PhpSprintfGetnumber(&format, &format_len)) < 0 {
							core.PhpErrorDocref("", faults.E_WARNING, "Precision must be greater than zero and less than %d", core.INT_MAX)
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
				core.PhpErrorDocref("", faults.E_WARNING, "Too few arguments")
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
				var str *types.String = operators.ZvalGetString(tmp)
				PhpSprintfAppendstring(&buf, &outpos, str.GetVal(), width, precision, padding, alignment, str.GetLen(), 0, expprec, 0)
			case 'd':
				PhpSprintfAppendint(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment, always_sign)
			case 'u':
				PhpSprintfAppenduint(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment)
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
				PhpSprintfAppenddouble(&buf, &outpos, operators.ZvalGetDouble(tmp), width, padding, alignment, precision, adjusting, *format, always_sign)
			case 'c':
				buf.WriteByte(byte(operators.ZvalGetLong(tmp)))
			case 'o':
				PhpSprintfAppend2n(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment, 3, Hexchars, expprec)
			case 'x':
				PhpSprintfAppend2n(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment, 4, Hexchars, expprec)
			case 'X':
				PhpSprintfAppend2n(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment, 4, HEXCHARS, expprec)
			case 'b':
				PhpSprintfAppend2n(&buf, &outpos, operators.ZvalGetLong(tmp), width, padding, alignment, 1, Hexchars, expprec)
			case '%':
				buf.WriteByte('%')
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

	return types.NewString(buf.String())
}
func PhpFormattedPrintGetArray(array *types.Zval, argc *int) *types.Zval {
	var args *types.Zval
	var zv *types.Zval
	var n int
	if !array.IsArray() {
		operators.ConvertToArray(array)
	}
	n = array.Array().Len()
	args = (*types.Zval)(zend.SafeEmalloc(n, b.SizeOf("zval"), 0))
	n = 0
	array.Array().Foreach(func(_ types.ArrayKey, zv *types.Zval) {
		types.ZVAL_COPY_VALUE(&args[n], zv)
		n++
	})
	*argc = n
	return args
}

//@zif -name sprintf
func ZifUserSprintf(format *types.Zval, _ zpp.Opt, args []*types.Zval) (string, bool) {
	var result *types.String
	var argc int
	result = PhpFormattedPrint(format, args, argc)
	if result == nil {
		return "", false
	}
	return result.GetStr(), true
}
func ZifVsprintf(executeData zpp.Ex, return_value zpp.Ret, format *types.Zval, args *types.Zval) {
	var result *types.String
	var format *types.Zval
	var array *types.Zval
	var args *types.Zval
	var argc int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
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
	return_value.SetStringEx(result)
}

//@zif -name printf
func ZifUserPrintf(executeData zpp.Ex, return_value zpp.Ret, format *types.Zval, _ zpp.Opt, args []*types.Zval) {
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
	rlen = core.PUTS(result.GetStr())
	// types.ZendStringEfree(result)
	return_value.SetLong(rlen)
	return
}
func ZifVprintf(executeData zpp.Ex, return_value zpp.Ret, format *types.Zval, args *types.Zval) {
	var result *types.String
	var rlen int
	var format *types.Zval
	var array *types.Zval
	var args *types.Zval
	var argc int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
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
	rlen = core.PUTS(result.GetStr())
	// types.ZendStringEfree(result)
	return_value.SetLong(rlen)
	return
}
func ZifFprintf(executeData zpp.Ex, return_value zpp.Ret, stream *types.Zval, format *types.Zval, _ zpp.Opt, args []*types.Zval) {
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
	// types.ZendStringEfree(result)
}
func ZifVfprintf(executeData zpp.Ex, return_value zpp.Ret, stream *types.Zval, format *types.Zval, args *types.Zval) {
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
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			arg1 = fp.ParseResource()
			format = fp.ParseZval()
			array = fp.ParseZval()
			if fp.HasError() {
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
	// types.ZendStringEfree(result)
}

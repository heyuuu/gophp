package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func __cvt(
	value float64,
	ndigit int,
	decpt *int,
	sign *int,
	fmode int,
	pad int,
) *byte {
	var s *byte = nil
	var p *byte
	var rve *byte
	var c byte
	var siz int
	if ndigit < 0 {
		siz = -ndigit + 1
	} else {
		siz = ndigit + 1
	}

	/* __dtoa() doesn't allocate space for 0 so we do it by hand */

	if value == 0.0 {
		*decpt = 1 - fmode
		*sign = 0
		if b.Assign(&rve, b.Assign(&s, (*byte)(zend.Malloc(b.Cond(ndigit != 0, siz, 2))))) == nil {
			return nil
		}
		b.PostInc(&(*rve)) = '0'
		*rve = '0'
		if ndigit == 0 {
			return s
		}
	} else {
		p = zend.ZendDtoa(value, fmode+2, ndigit, decpt, sign, &rve)
		if (*decpt) == 9999 {

			/* Infinity or Nan, convert to inf or nan like printf */

			*decpt = 0
			c = *p
			zend.ZendFreedtoa(p)
			return strdup(b.Cond(c == 'I', "INF", "NAN"))
		}

		/* Make a local copy and adjust rve to be in terms of s */

		if pad != 0 && fmode != 0 {
			siz += *decpt
		}
		if b.Assign(&s, (*byte)(zend.Malloc(siz+1))) == nil {
			zend.ZendFreedtoa(p)
			return nil
		}
		void(strlcpy(s, p, siz))
		rve = s + (rve - p)
		zend.ZendFreedtoa(p)
	}

	/* Add trailing zeros */

	if pad != 0 {
		siz -= rve - s
		for b.PreDec(&siz) {
			b.PostInc(&(*rve)) = '0'
		}
		*rve = '0'
	}
	return s
}
func PhpEcvt(value float64, ndigit int, decpt *int, sign *int) *byte {
	return __cvt(value, ndigit, decpt, sign, 0, 1)
}
func PhpFcvt(value float64, ndigit int, decpt *int, sign *int) *byte {
	return __cvt(value, ndigit, decpt, sign, 1, 1)
}
func PhpGcvt(value float64, ndigit int, dec_point byte, exponent byte, buf *byte) *byte {
	var digits *byte
	var dst *byte
	var src *byte
	var i int
	var decpt int
	var sign int
	var mode int = b.Cond(ndigit >= 0, 2, 0)
	if mode == 0 {
		ndigit = 17
	}
	digits = zend.ZendDtoa(value, mode, ndigit, &decpt, &sign, nil)
	if decpt == 9999 {

		/*
		 * Infinity or NaN, convert to inf or nan with sign.
		 * We assume the buffer is at least ndigit long.
		 */

		Snprintf(buf, ndigit+1, "%s%s", b.Cond(sign != 0 && (*digits) == 'I', "-", ""), b.Cond((*digits) == 'I', "INF", "NAN"))
		zend.ZendFreedtoa(digits)
		return buf
	}
	dst = buf
	if sign != 0 {
		b.PostInc(&(*dst)) = '-'
	}
	if decpt >= 0 && decpt > ndigit || decpt < -3 {

		/* exponential format (e.g. 1.2345e+13) */

		if b.PreDec(&decpt) < 0 {
			sign = 1
			decpt = -decpt
		} else {
			sign = 0
		}
		src = digits
		*src++
		b.PostInc(&(*dst)) = (*src) - 1
		b.PostInc(&(*dst)) = dec_point
		if (*src) == '0' {
			b.PostInc(&(*dst)) = '0'
		} else {
			for {
				*src++
				b.PostInc(&(*dst)) = (*src) - 1
				if (*src) == '0' {
					break
				}
			}
		}
		b.PostInc(&(*dst)) = exponent
		if sign != 0 {
			b.PostInc(&(*dst)) = '-'
		} else {
			b.PostInc(&(*dst)) = '+'
		}
		if decpt < 10 {
			b.PostInc(&(*dst)) = '0' + decpt
			*dst = '0'
		} else {

			/* XXX - optimize */

			sign = decpt
			i = 0
			for ; b.AssignOp(&sign, "/=", 10) != 0; i++ {

			}
			dst[i+1] = '0'
			for decpt != 0 {
				dst[b.PostDec(&i)] = '0' + decpt%10
				decpt /= 10
			}
		}
	} else if decpt < 0 {

		/* standard format 0. */

		b.PostInc(&(*dst)) = '0'
		b.PostInc(&(*dst)) = dec_point
		for {
			b.PostInc(&(*dst)) = '0'
			if b.PreInc(&decpt) >= 0 {
				break
			}
		}
		src = digits
		for (*src) != '0' {
			*src++
			b.PostInc(&(*dst)) = (*src) - 1
		}
		*dst = '0'
	} else {

		/* standard format */

		i = 0
		src = digits
		for ; i < decpt; i++ {
			if (*src) != '0' {
				*src++
				b.PostInc(&(*dst)) = (*src) - 1
			} else {
				b.PostInc(&(*dst)) = '0'
			}
		}
		if (*src) != '0' {
			if src == digits {
				b.PostInc(&(*dst)) = '0'
			}
			b.PostInc(&(*dst)) = dec_point
			for i = decpt; digits[i] != '0'; i++ {
				b.PostInc(&(*dst)) = digits[i]
			}
		}
		*dst = '0'
	}
	zend.ZendFreedtoa(digits)
	return buf
}
func ApPhpConv10(num WideInt, is_unsigned BoolInt, is_negative *BoolInt, buf_end *byte, len_ *int) *byte {
	var p *byte = buf_end
	var magnitude UWideInt
	if is_unsigned != 0 {
		magnitude = UWideInt(num)
		*is_negative = FALSE
	} else {
		*is_negative = num < 0

		/*
		 * On a 2's complement machine, negating the most negative integer
		 * results in a number that cannot be represented as a signed integer.
		 * Here is what we do to obtain the number's magnitude:
		 *      a. add 1 to the number
		 *      b. negate it (becomes positive)
		 *      c. convert it to unsigned
		 *      d. add 1
		 */

		if (*is_negative) != 0 {
			var t WideInt = num + 1
			magnitude = u_wide_int - t + 1
		} else {
			magnitude = UWideInt(num)
		}

		/*
		 * On a 2's complement machine, negating the most negative integer
		 * results in a number that cannot be represented as a signed integer.
		 * Here is what we do to obtain the number's magnitude:
		 *      a. add 1 to the number
		 *      b. negate it (becomes positive)
		 *      c. convert it to unsigned
		 *      d. add 1
		 */

	}

	/*
	 * We use a do-while loop so that we write at least 1 digit
	 */

	for {
		var new_magnitude UWideInt = magnitude / 10
		*(b.PreDec(&p)) = byte(magnitude - new_magnitude*10 + '0')
		magnitude = new_magnitude
		if !magnitude {
			break
		}
	}
	*len_ = buf_end - p
	return p
}
func PhpConvFp(
	format byte,
	num float64,
	add_dp BooleanE,
	precision int,
	dec_point byte,
	is_negative *BoolInt,
	buf *byte,
	len_ *int,
) *byte {
	var s *byte = buf
	var p *byte
	var p_orig *byte
	var decimal_point int
	if precision >= NDIG-1 {
		precision = NDIG - 2
	}
	if format == 'F' {
		p = PhpFcvt(num, precision, &decimal_point, is_negative)
		p_orig = p
	} else {
		p = PhpEcvt(num, precision+1, &decimal_point, is_negative)
		p_orig = p
	}

	/*
	 * Check for Infinity and NaN
	 */

	if isalpha(int(*p)) {
		*len_ = strlen(p)
		memcpy(buf, p, (*len_)+1)
		*is_negative = FALSE
		zend.Free(p_orig)
		return buf
	}
	if format == 'F' {
		if decimal_point <= 0 {
			if num != 0 || precision > 0 {
				b.PostInc(&(*s)) = '0'
				if precision > 0 {
					b.PostInc(&(*s)) = dec_point
					for b.PostInc(&decimal_point) < 0 {
						b.PostInc(&(*s)) = '0'
					}
				} else if add_dp != 0 {
					b.PostInc(&(*s)) = dec_point
				}
			}
		} else {
			var addz int = b.Cond(decimal_point >= NDIG, decimal_point-NDIG+1, 0)
			decimal_point -= addz
			for b.PostDec(&decimal_point) > 0 {
				*p++
				b.PostInc(&(*s)) = (*p) - 1
			}
			for b.PostDec(&addz) > 0 {
				b.PostInc(&(*s)) = '0'
			}
			if precision > 0 || add_dp != 0 {
				b.PostInc(&(*s)) = dec_point
			}
		}
	} else {
		*p++
		b.PostInc(&(*s)) = (*p) - 1
		if precision > 0 || add_dp != 0 {
			b.PostInc(&(*s)) = '.'
		}
	}

	/*
	 * copy the rest of p, the NUL is NOT copied
	 */

	for *p {
		*p++
		b.PostInc(&(*s)) = (*p) - 1
	}
	if format != 'F' {
		var temp []byte
		var t_len int
		var exponent_is_negative BoolInt
		b.PostInc(&(*s)) = format
		decimal_point--
		if decimal_point != 0 {
			p = ApPhpConv10(WideInt(decimal_point), FALSE, &exponent_is_negative, &temp[EXPONENT_LENGTH], &t_len)
			if exponent_is_negative != 0 {
				b.PostInc(&(*s)) = '-'
			} else {
				b.PostInc(&(*s)) = '+'
			}

			/*
			 * Make sure the exponent has at least 2 digits
			 */

			for b.PostDec(&t_len) {
				*p++
				b.PostInc(&(*s)) = (*p) - 1
			}

			/*
			 * Make sure the exponent has at least 2 digits
			 */

		} else {
			b.PostInc(&(*s)) = '+'
			b.PostInc(&(*s)) = '0'
		}
	}
	*len_ = s - buf
	zend.Free(p_orig)
	return buf
}
func ApPhpConvP2(num UWideInt, nbits int, format byte, buf_end *byte, len_ *int) *byte {
	var mask int = (1 << nbits) - 1
	var p *byte = buf_end
	var low_digits []byte = "0123456789abcdef"
	var upper_digits []byte = "0123456789ABCDEF"
	var digits *byte = b.Cond(format == 'X', upper_digits, low_digits)
	for {
		*(b.PreDec(&p)) = digits[num&mask]
		num >>= nbits
		if !num {
			break
		}
	}
	*len_ = buf_end - p
	return p
}
func NUM(c char) int { return c - '0' }
func STR_TO_DEC(str *byte, num int) {
	num = NUM(b.PostInc(&(*str)))
	for isdigit(int(*str)) {
		num *= 10
		num += NUM(b.PostInc(&(*str)))
	}
}
func FIX_PRECISION(adjust BooleanE, precision int, s *byte, s_len int) {
	if adjust != 0 {
		for s_len < int(precision) {
			*(b.PreDec(&s)) = '0'
			s_len++
		}
	}
}
func PAD(width int, len_ int, ch byte) {
	for {
		if sp < bep {
			b.PostInc(&(*sp)) = ch
		}
		cc++
		width--
		if int(width > len_) == 0 {
			break
		}
	}
}
func FormatConverter(odp *Buffy, fmt *byte, ap ...any) int {
	var sp *byte
	var bep *byte
	var cc int = 0
	var i int
	var s *byte = nil
	var s_len int
	var free_zcopy int
	var zvp *types.Zval
	var zcopy types.Zval
	var min_width int = 0
	var precision int = 0
	var adjust int
	var pad_char byte
	var prefix_char byte
	var fp_num float64
	var i_num WideInt = WideInt(0)
	var ui_num UWideInt
	var num_buf []byte
	var char_buf []byte
	var lconv *__struct__lconv = nil

	/*
	 * Flag variables
	 */

	var modifier LengthModifierE
	var alternate_form BooleanE
	var print_sign BooleanE
	var print_blank BooleanE
	var adjust_precision BooleanE
	var adjust_width BooleanE
	var is_negative BoolInt
	sp = odp.GetNextb()
	bep = odp.GetBufEnd()
	for *fmt {
		if (*fmt) != '%' {
			if sp < bep {
				b.PostInc(&(*sp)) = *fmt
			}
			cc++
		} else {

			/*
			 * Default variable settings
			 */

			adjust = RIGHT
			print_blank = NO
			print_sign = print_blank
			alternate_form = print_sign
			pad_char = ' '
			prefix_char = NUL
			free_zcopy = 0
			fmt++

			/*
			 * Try to avoid checking for flags, width or precision
			 */

			if isascii(int(*fmt)) && !(islower(int(*fmt))) {

				/*
				 * Recognize flags: -, #, BLANK, +
				 */

				for ; ; fmt++ {
					if (*fmt) == '-' {
						adjust = LEFT
					} else if (*fmt) == '+' {
						print_sign = YES
					} else if (*fmt) == '#' {
						alternate_form = YES
					} else if (*fmt) == ' ' {
						print_blank = YES
					} else if (*fmt) == '0' {
						pad_char = '0'
					} else {
						break
					}
				}

				/*
				 * Check if a width was specified
				 */

				if isdigit(int(*fmt)) {
					STR_TO_DEC(fmt, min_width)
					adjust_width = YES
				} else if (*fmt) == '*' {
					min_width = __va_arg(ap, int(_))
					fmt++
					adjust_width = YES
					if min_width < 0 {
						adjust = LEFT
						min_width = -min_width
					}
				} else {
					adjust_width = NO
				}

				/*
				 * Check if a precision was specified
				 */

				if (*fmt) == '.' {
					adjust_precision = YES
					fmt++
					if isdigit(int(*fmt)) {
						STR_TO_DEC(fmt, precision)
					} else if (*fmt) == '*' {
						precision = __va_arg(ap, int(_))
						fmt++
						if precision < 0 {
							precision = 0
						}
					} else {
						precision = 0
					}
					if precision > FORMAT_CONV_MAX_PRECISION {
						precision = FORMAT_CONV_MAX_PRECISION
					}
				} else {
					adjust_precision = NO
				}

				/*
				 * Check if a precision was specified
				 */

			} else {
				adjust_width = NO
				adjust_precision = adjust_width
			}

			/*
			 * Modifier check
			 */

			switch *fmt {
			case 'L':
				fmt++
				modifier = LM_LONG_DOUBLE
			case 'I':
				fmt++
				if (*fmt) == '6' && (*(fmt + 1)) == '4' {
					fmt += 2
					modifier = LM_LONG_LONG
				} else if (*fmt) == '3' && (*(fmt + 1)) == '2' {
					fmt += 2
					modifier = LM_LONG
				} else {
					modifier = LM_LONG
				}
			case 'l':
				fmt++
				if (*fmt) == 'l' {
					fmt++
					modifier = LM_LONG_LONG
				} else {
					modifier = LM_LONG
				}
			case 'z':
				fmt++
				modifier = LM_SIZE_T
			case 'j':
				fmt++
				modifier = LM_INTMAX_T
			case 't':
				fmt++
				modifier = LM_PTRDIFF_T
			case 'p':
				fmt++
				modifier = LM_PHP_INT_T
			case 'h':
				fmt++
				if (*fmt) == 'h' {
					fmt++
				}
				fallthrough
			default:
				modifier = LM_STD
			}

			/*
			 * Argument extraction and printing.
			 * First we determine the argument type.
			 * Then, we convert the argument to a string.
			 * On exit from the switch, s points to the string that
			 * must be printed, s_len has the length of the string
			 * The precision requirements, if any, are reflected in s_len.
			 *
			 * NOTE: pad_char may be set to '0' because of the 0 flag.
			 *   It is reset to ' ' by non-numeric formats
			 */

			switch *fmt {
			case 'Z':
				zvp = (*types.Zval)(__va_arg(ap, (*types.Zval)(_)))
				free_zcopy = zend.ZendMakePrintableZval(zvp, &zcopy)
				if free_zcopy != 0 {
					zvp = &zcopy
				}
				s_len = zvp.GetStr().GetLen()
				s = zvp.GetStr().GetVal()
				if adjust_precision != 0 && int(precision < s_len) != 0 {
					s_len = precision
				}
			case 'u':
				switch modifier {
				default:
					i_num = WideInt(__va_arg(ap, uint(_)))
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					i_num = WideInt(__va_arg(ap, unsigned__long__int(_)))
				case LM_SIZE_T:
					i_num = WideInt(__va_arg(ap, int(_)))
				case LM_LONG_LONG:
					i_num = WideInt(__va_arg(ap, UWideInt(_)))
				case LM_INTMAX_T:
					i_num = WideInt(__va_arg(ap, uintmax_t(_)))
				case LM_PTRDIFF_T:
					i_num = WideInt(__va_arg(ap, ptrdiff_t(_)))
				case LM_PHP_INT_T:
					i_num = WideInt(__va_arg(ap, zend.ZendUlong(_)))
				}
				fallthrough
			case 'd':
				fallthrough
			case 'i':

				/*
				 * Get the arg if we haven't already.
				 */

				if (*fmt) != 'u' {
					switch modifier {
					default:
						i_num = WideInt(__va_arg(ap, int(_)))
					case LM_LONG_DOUBLE:
						goto fmt_error
					case LM_LONG:
						i_num = WideInt(__va_arg(ap, long__int(_)))
					case LM_SIZE_T:
						i_num = WideInt(__va_arg(ap, ssize_t(_)))
					case LM_LONG_LONG:
						i_num = WideInt(__va_arg(ap, WideInt(_)))
					case LM_INTMAX_T:
						i_num = WideInt(__va_arg(ap, intmax_t(_)))
					case LM_PTRDIFF_T:
						i_num = WideInt(__va_arg(ap, ptrdiff_t(_)))
					case LM_PHP_INT_T:
						i_num = WideInt(__va_arg(ap, zend.ZendLong(_)))
					}
				}
				s = ApPhpConv10(i_num, (*fmt) == 'u', &is_negative, &num_buf[NUM_BUF_SIZE], &s_len)
				FIX_PRECISION(adjust_precision, precision, s, s_len)
				if (*fmt) != 'u' {
					if is_negative != 0 {
						prefix_char = '-'
					} else if print_sign != 0 {
						prefix_char = '+'
					} else if print_blank != 0 {
						prefix_char = ' '
					}
				}
			case 'o':
				switch modifier {
				default:
					ui_num = UWideInt(__va_arg(ap, uint(_)))
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					ui_num = UWideInt(__va_arg(ap, unsigned__long__int(_)))
				case LM_SIZE_T:
					ui_num = UWideInt(__va_arg(ap, int(_)))
				case LM_LONG_LONG:
					ui_num = UWideInt(__va_arg(ap, UWideInt(_)))
				case LM_INTMAX_T:
					ui_num = UWideInt(__va_arg(ap, uintmax_t(_)))
				case LM_PTRDIFF_T:
					ui_num = UWideInt(__va_arg(ap, ptrdiff_t(_)))
				case LM_PHP_INT_T:
					ui_num = UWideInt(__va_arg(ap, zend.ZendUlong(_)))
				}
				s = ApPhpConvP2(ui_num, 3, *fmt, &num_buf[NUM_BUF_SIZE], &s_len)
				FIX_PRECISION(adjust_precision, precision, s, s_len)
				if alternate_form != 0 && (*s) != '0' {
					*(b.PreDec(&s)) = '0'
					s_len++
				}
			case 'x':
				fallthrough
			case 'X':
				switch modifier {
				default:
					ui_num = UWideInt(__va_arg(ap, uint(_)))
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					ui_num = UWideInt(__va_arg(ap, unsigned__long__int(_)))
				case LM_SIZE_T:
					ui_num = UWideInt(__va_arg(ap, int(_)))
				case LM_LONG_LONG:
					ui_num = UWideInt(__va_arg(ap, UWideInt(_)))
				case LM_INTMAX_T:
					ui_num = UWideInt(__va_arg(ap, uintmax_t(_)))
				case LM_PTRDIFF_T:
					ui_num = UWideInt(__va_arg(ap, ptrdiff_t(_)))
				case LM_PHP_INT_T:
					ui_num = UWideInt(__va_arg(ap, zend.ZendUlong(_)))
				}
				s = ApPhpConvP2(ui_num, 4, *fmt, &num_buf[NUM_BUF_SIZE], &s_len)
				FIX_PRECISION(adjust_precision, precision, s, s_len)
				if alternate_form != 0 && i_num != 0 {
					*(b.PreDec(&s)) = *fmt
					*(b.PreDec(&s)) = '0'
					s_len += 2
				}
			case 's':
				fallthrough
			case 'v':
				s = __va_arg(ap, (*byte)(_))
				if s != nil {
					s_len = strlen(s)
					if adjust_precision != 0 && int(precision < s_len) != 0 {
						s_len = precision
					}
				} else {
					s = S_NULL
					s_len = S_NULL_LEN
				}
				pad_char = ' '
			case 'f':
				fallthrough
			case 'F':
				fallthrough
			case 'e':
				fallthrough
			case 'E':
				switch modifier {
				case LM_LONG_DOUBLE:
					fp_num = float64(__va_arg(ap, long__double(_)))
				case LM_STD:
					fp_num = __va_arg(ap, float64(_))
				default:
					goto fmt_error
				}
				if ZendIsNaN(fp_num) {
					s = "NAN"
					s_len = 3
				} else if ZendIsInf(fp_num) {
					s = "INF"
					s_len = 3
				} else {
					if lconv == nil {
						lconv = localeconv()
					}
					s = PhpConvFp(b.Cond((*fmt) == 'f', 'F', *fmt), fp_num, alternate_form, b.Cond(adjust_precision == NO, FLOAT_DIGITS, precision), b.Cond((*fmt) == 'f', LCONV_DECIMAL_POINT, '.'), &is_negative, &num_buf[1], &s_len)
					if is_negative != 0 {
						prefix_char = '-'
					} else if print_sign != 0 {
						prefix_char = '+'
					} else if print_blank != 0 {
						prefix_char = ' '
					}
				}
			case 'g':
				fallthrough
			case 'k':
				fallthrough
			case 'G':
				fallthrough
			case 'H':
				switch modifier {
				case LM_LONG_DOUBLE:
					fp_num = float64(__va_arg(ap, long__double(_)))
				case LM_STD:
					fp_num = __va_arg(ap, float64(_))
				default:
					goto fmt_error
				}
				if ZendIsNaN(fp_num) {
					s = "NAN"
					s_len = 3
					break
				} else if ZendIsInf(fp_num) {
					if fp_num > 0 {
						s = "INF"
						s_len = 3
					} else {
						s = "-INF"
						s_len = 4
					}
					break
				}
				if adjust_precision == NO {
					precision = FLOAT_DIGITS
				} else if precision == 0 {
					precision = 1
				}

				/*
				 * * We use &num_buf[ 1 ], so that we have room for the sign
				 */

				if lconv == nil {
					lconv = localeconv()
				}
				s = PhpGcvt(fp_num, precision, b.Cond((*fmt) == 'H' || (*fmt) == 'k', '.', LCONV_DECIMAL_POINT), b.Cond((*fmt) == 'G' || (*fmt) == 'H', 'E', 'e'), &num_buf[1])
				if (*s) == '-' {
					*s++
					prefix_char = (*s) - 1
				} else if print_sign != 0 {
					prefix_char = '+'
				} else if print_blank != 0 {
					prefix_char = ' '
				}
				s_len = strlen(s)
				if alternate_form != 0 && strchr(s, '.') == nil {
					s[b.PostInc(&s_len)] = '.'
				}
			case 'c':
				char_buf[0] = byte(__va_arg(ap, int(_)))
				s = &char_buf[0]
				s_len = 1
				pad_char = ' '
			case '%':
				char_buf[0] = '%'
				s = &char_buf[0]
				s_len = 1
				pad_char = ' '
			case 'n':
				*(__va_arg(ap, (*int)(_))) = cc
				goto skip_output
			case 'p':
				if b.SizeOf("char *") <= b.SizeOf("u_wide_int") {
					ui_num = u_wide_int(int(__va_arg(ap, (*byte)(_))))
					s = ApPhpConvP2(ui_num, 4, 'x', &num_buf[NUM_BUF_SIZE], &s_len)
					if ui_num != 0 {
						*(b.PreDec(&s)) = 'x'
						*(b.PreDec(&s)) = '0'
						s_len += 2
					}
				} else {
					s = "%p"
					s_len = 2
				}
				pad_char = ' '
			case NUL:

				/*
				 * The last character of the format string was %.
				 * We ignore it.
				 */

				continue
			fmt_error:
				PhpError(faults.E_ERROR, "Illegal length modifier specified '%c' in s[np]printf call", *fmt)
			default:
				char_buf[0] = '%'
				char_buf[1] = *fmt
				s = char_buf
				s_len = 2
				pad_char = ' '
			}
			if prefix_char != NUL {
				*(b.PreDec(&s)) = prefix_char
				s_len++
			}
			if adjust_width != 0 && adjust == RIGHT && int(min_width > s_len) != 0 {
				if pad_char == '0' && prefix_char != NUL {
					if sp < bep {
						b.PostInc(&(*sp)) = *s
					}
					cc++
					s++
					s_len--
					min_width--
				}
				PAD(min_width, s_len, pad_char)
			}

			/*
			 * Print the string s.
			 */

			for i = s_len; i != 0; i-- {
				if sp < bep {
					b.PostInc(&(*sp)) = *s
				}
				cc++
				s++
			}
			if adjust_width != 0 && adjust == LEFT && int(min_width > s_len) != 0 {
				PAD(min_width, s_len, pad_char)
			}
			if free_zcopy != 0 {

			}
		}
	skip_output:
		fmt++
	}
	odp.SetNextb(sp)
	return cc
}
func StrxPrintv(ccp *int, buf *byte, len_ int, format *byte, ap ...any) {
	var od Buffy
	var cc int

	/*
	 * First initialize the descriptor
	 * Notice that if no length is given, we initialize buf_end to the
	 * highest possible address.
	 */

	if len_ == 0 {
		od.SetBufEnd((*byte)(^0))
		od.SetNextb((*byte)(^0))
	} else {
		od.SetBufEnd(&buf[len_-1])
		od.SetNextb(buf)
	}

	/*
	 * Do the conversion
	 */

	cc = FormatConverter(&od, format, ap)
	if len_ != 0 && od.GetNextb() <= od.GetBufEnd() {
		*(od.GetNextb()) = '0'
	}
	if ccp != nil {
		*ccp = cc
	}
}
func ApPhpSlprintf(buf *byte, len_ int, format *byte, _ ...any) int {
	var cc int
	var ap va_list
	va_start(ap, format)
	StrxPrintv(&cc, buf, len_, format, ap)
	va_end(ap)
	if int(cc >= len_) != 0 {
		cc = int(len_ - 1)
		buf[cc] = '0'
	}
	return cc
}
func ApPhpSnprintf(buf *byte, len_ int, format *byte, _ ...any) int {
	var cc int
	var ap va_list
	va_start(ap, format)
	StrxPrintv(&cc, buf, len_, format, ap)
	va_end(ap)
	return cc
}

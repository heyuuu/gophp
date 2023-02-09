// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/zend"
)

func INS_STRING(xbuf any, str *byte, len_ int, is_char zend.ZendBool) {
	if is_char != 0 {
		(*zend.SmartString)(xbuf).AppendString(b.CastStr(str, len_))
	} else {
		(*zend.SmartStr)(xbuf).AppendString(b.CastStr(str, len_))
	}
}
func PAD_CHAR(xbuf any, ch byte, count int, is_char zend.ZendBool) {
	if is_char != 0 {
		(*zend.SmartString)(xbuf).Alloc(count)
		memset((*zend.SmartString)(xbuf).GetC()+(*zend.SmartString)(xbuf).GetLen(), ch, count)
		(*zend.SmartString)(xbuf).SetLen((*zend.SmartString)(xbuf).GetLen() + count)
	} else {
		(*zend.SmartStr)(xbuf).Alloc(count)
		memset((*zend.SmartStr)(xbuf).GetS().GetVal()+(*zend.SmartStr)(xbuf).GetS().GetLen(), ch, count)
		(*zend.SmartStr)(xbuf).GetS().GetLen() += count
	}
}
func XbufFormatConverter(xbuf any, is_char zend.ZendBool, fmt *byte, ap ...any) {
	var s *byte = nil
	var s_len int
	var free_zcopy int
	var zvp *zend.Zval
	var zcopy zend.Zval
	var min_width int = 0
	var precision int = 0
	var adjust int
	var pad_char byte
	var prefix_char byte
	var fp_num float64
	var i_num WideInt = WideInt(0)
	var ui_num UWideInt = UWideInt(0)
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
	for *fmt {
		if (*fmt) != '%' {
			if is_char != 0 {
				(*zend.SmartString)(xbuf).AppendByte(*fmt)
			} else {
				(*zend.SmartStr)(xbuf).AppendByte(*fmt)
			}
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
						if precision < -1 {
							precision = -1
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
				var __next byte = *(fmt + 1)
				if 'd' == __next || 'u' == __next || 'x' == __next || 'o' == __next {
					fmt++
					modifier = LM_PHP_INT_T
				} else {
					modifier = LM_STD
				}
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
				zvp = (*zend.Zval)(__va_arg(ap, (*zend.Zval)(_)))
				free_zcopy = zend.ZendMakePrintableZval(zvp, &zcopy)
				if free_zcopy != 0 {
					zvp = &zcopy
				}
				s_len = zend.Z_STRLEN_P(zvp)
				s = zend.Z_STRVAL_P(zvp)
				if adjust_precision != 0 && int(precision < s_len) != 0 {
					s_len = precision
				}
			case 'u':
				switch modifier {
				default:
					i_num = WideInt(__va_arg(ap, uint(_)))
				case LM_LONG_DOUBLE:
					goto fmt_error
					fallthrough
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
						fallthrough
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
					fallthrough
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
					fallthrough
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
				if alternate_form != 0 && ui_num != 0 {
					*(b.PreDec(&s)) = *fmt
					*(b.PreDec(&s)) = '0'
					s_len += 2
				}
			case 's':
				fallthrough
			case 'v':
				s = __va_arg(ap, (*byte)(_))
				if s != nil {
					if adjust_precision == 0 {
						s_len = strlen(s)
					} else {
						s_len = strnlen(s, precision)
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
				if ZendIsnan(fp_num) {
					s = "nan"
					s_len = 3
				} else if ZendIsinf(fp_num) {
					s = "inf"
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
				if ZendIsnan(fp_num) {
					s = "NAN"
					s_len = 3
					break
				} else if ZendIsinf(fp_num) {
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
				if is_char != 0 {
					*(__va_arg(ap, (*int)(_))) = int((*zend.SmartString)(xbuf).GetLen())
				} else {
					*(__va_arg(ap, (*int)(_))) = int((*zend.SmartStr)(xbuf).GetS().GetLen())
				}
				goto skip_output
				fallthrough
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
				PhpError(zend.E_ERROR, "Illegal length modifier specified '%c' in s[np]printf call", *fmt)
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
					if is_char != 0 {
						(*zend.SmartString)(xbuf).AppendByte(*s)
					} else {
						(*zend.SmartStr)(xbuf).AppendByte(*s)
					}
					s++
					s_len--
					min_width--
				}
				PAD_CHAR(xbuf, pad_char, min_width-s_len, is_char)
			}

			/*
			 * Print the string s.
			 */

			INS_STRING(xbuf, s, s_len, is_char)
			if adjust_width != 0 && adjust == LEFT && int(min_width > s_len) != 0 {
				PAD_CHAR(xbuf, pad_char, min_width-s_len, is_char)
			}
			if free_zcopy != 0 {
				zend.ZvalPtrDtorStr(&zcopy)
			}
		}
	skip_output:
		fmt++
	}
	return
}
func PhpPrintfToSmartString(buf *zend.SmartString, format *byte, ap ...any) {
	XbufFormatConverter(buf, 1, format, ap)
}
func PhpPrintfToSmartStr(buf *zend.SmartStr, format *byte, ap ...any) {
	XbufFormatConverter(buf, 0, format, ap)
}

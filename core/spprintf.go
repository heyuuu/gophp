// <<generate>>

package core

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/spprintf.h>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define SPPRINTF_H

// # include "snprintf.h"

// # include "zend_smart_str_public.h"

// # include "zend_smart_string_public.h"

// #define spprintf       zend_spprintf

// #define strpprintf       zend_strpprintf

// #define vspprintf       zend_vspprintf

// #define vstrpprintf       zend_vstrpprintf

// Source: <main/spprintf.c>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define _GNU_SOURCE

// # include "php.h"

// # include < stddef . h >

// # include < stdio . h >

// # include < ctype . h >

// # include < sys / types . h >

// # include < stdarg . h >

// # include < string . h >

// # include < stdlib . h >

// # include < math . h >

// # include < inttypes . h >

// # include < locale . h >

// #define LCONV_DECIMAL_POINT       ( * lconv -> decimal_point )

// # include "snprintf.h"

// #define FALSE       0

// #define TRUE       1

// #define NUL       '\0'

// #define INT_NULL       ( ( int * ) 0 )

// #define S_NULL       "(null)"

// #define S_NULL_LEN       6

// #define FLOAT_DIGITS       6

// #define EXPONENT_LENGTH       10

// # include "zend_smart_str.h"

// # include "zend_smart_string.h"

/* {{{ macros */

// #define INS_CHAR(xbuf,ch,is_char) do { if ( ( is_char ) ) { smart_string_appendc ( ( smart_string * ) ( xbuf ) , ( ch ) ) ; } else { smart_str_appendc ( ( smart_str * ) ( xbuf ) , ( ch ) ) ; } } while ( 0 ) ;

// #define INS_STRING(xbuf,str,len,is_char) do { if ( ( is_char ) ) { smart_string_appendl ( ( smart_string * ) ( xbuf ) , ( str ) , ( len ) ) ; } else { smart_str_appendl ( ( smart_str * ) ( xbuf ) , ( str ) , ( len ) ) ; } } while ( 0 ) ;

// #define PAD_CHAR(xbuf,ch,count,is_char) do { if ( ( is_char ) ) { smart_string_alloc ( ( ( smart_string * ) ( xbuf ) ) , ( count ) , 0 ) ; memset ( ( ( smart_string * ) ( xbuf ) ) -> c + ( ( smart_string * ) ( xbuf ) ) -> len , ( ch ) , ( count ) ) ; ( ( smart_string * ) ( xbuf ) ) -> len += ( count ) ; } else { smart_str_alloc ( ( ( smart_str * ) ( xbuf ) ) , ( count ) , 0 ) ; memset ( ZSTR_VAL ( ( ( smart_str * ) ( xbuf ) ) -> s ) + ZSTR_LEN ( ( ( smart_str * ) ( xbuf ) ) -> s ) , ( ch ) , ( count ) ) ; ZSTR_LEN ( ( ( smart_str * ) ( xbuf ) ) -> s ) += ( count ) ; } } while ( 0 ) ;

/*
 * NUM_BUF_SIZE is the size of the buffer used for arithmetic conversions
 * which can be at most max length of double
 */

// #define NUM_BUF_SIZE       PHP_DOUBLE_MAX_LENGTH

// #define NUM(c) ( c - '0' )

// #define STR_TO_DEC(str,num) do { num = NUM ( * str ++ ) ; while ( isdigit ( ( int ) * str ) ) { num *= 10 ; num += NUM ( * str ++ ) ; if ( num >= INT_MAX / 10 ) { while ( isdigit ( ( int ) * str ++ ) ) ; break ; } } } while ( 0 )

/*
 * This macro does zero padding so that the precision
 * requirement is satisfied. The padding is done by
 * adding '0's to the left of the string that is going
 * to be printed.
 */

// #define FIX_PRECISION(adjust,precision,s,s_len) do { if ( adjust ) while ( s_len < ( size_t ) precision ) { * -- s = '0' ; s_len ++ ; } } while ( 0 )

/* }}} */

/*
 * Do format conversion placing the output in buffer
 */

func XbufFormatConverter(xbuf any, is_char zend.ZendBool, fmt *byte, ap va_list) {
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
				zend.SmartStringAppendcEx((*zend.SmartString)(xbuf), *fmt, 0)
			} else {
				zend.SmartStrAppendcEx((*zend.SmartStr)(xbuf), *fmt, 0)
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
			prefix_char = '0'
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
					for {
						min_width = g.PostInc(&(*fmt)) - '0'
						for isdigit(int(*fmt)) {
							min_width *= 10
							min_width += g.PostInc(&(*fmt)) - '0'
							if min_width >= 2147483647/10 {
								for isdigit(int(g.PostInc(&(*fmt)))) {

								}
								break
							}
						}
						break
					}
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
						for {
							precision = g.PostInc(&(*fmt)) - '0'
							for isdigit(int(*fmt)) {
								precision *= 10
								precision += g.PostInc(&(*fmt)) - '0'
								if precision >= 2147483647/10 {
									for isdigit(int(g.PostInc(&(*fmt)))) {

									}
									break
								}
							}
							break
						}
					} else if (*fmt) == '*' {
						precision = __va_arg(ap, int(_))
						fmt++
						if precision < -1 {
							precision = -1
						}
					} else {
						precision = 0
					}
					if precision > 500 {
						precision = 500
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
				break
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
				break
			case 'l':
				fmt++
				if (*fmt) == 'l' {
					fmt++
					modifier = LM_LONG_LONG
				} else {
					modifier = LM_LONG
				}
				break
			case 'z':
				fmt++
				modifier = LM_SIZE_T
				break
			case 'j':
				fmt++
				modifier = LM_INTMAX_T
				break
			case 't':
				fmt++
				modifier = LM_PTRDIFF_T
				break
			case 'p':
				var __next byte = *(fmt + 1)
				if 'd' == __next || 'u' == __next || 'x' == __next || 'o' == __next {
					fmt++
					modifier = LM_PHP_INT_T
				} else {
					modifier = LM_STD
				}
				break
			case 'h':
				fmt++
				if (*fmt) == 'h' {
					fmt++
				}
			default:
				modifier = LM_STD
				break
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
				s_len = zvp.value.str.len_
				s = zvp.value.str.val
				if adjust_precision != 0 && int(precision < s_len) != 0 {
					s_len = precision
				}
				break
			case 'u':
				switch modifier {
				default:
					i_num = WideInt(__va_arg(ap, uint(_)))
					break
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					i_num = WideInt(__va_arg(ap, unsigned__long__int(_)))
					break
				case LM_SIZE_T:
					i_num = WideInt(__va_arg(ap, int(_)))
					break
				case LM_LONG_LONG:
					i_num = WideInt(__va_arg(ap, UWideInt(_)))
					break
				case LM_INTMAX_T:
					i_num = WideInt(__va_arg(ap, uintmax_t(_)))
					break
				case LM_PTRDIFF_T:
					i_num = WideInt(__va_arg(ap, ptrdiff_t(_)))
					break
				case LM_PHP_INT_T:
					i_num = WideInt(__va_arg(ap, zend.ZendUlong(_)))
					break
				}
			case 'd':

			case 'i':

				/*
				 * Get the arg if we haven't already.
				 */

				if (*fmt) != 'u' {
					switch modifier {
					default:
						i_num = WideInt(__va_arg(ap, int(_)))
						break
					case LM_LONG_DOUBLE:
						goto fmt_error
					case LM_LONG:
						i_num = WideInt(__va_arg(ap, long__int(_)))
						break
					case LM_SIZE_T:
						i_num = WideInt(__va_arg(ap, ssize_t(_)))
						break
					case LM_LONG_LONG:
						i_num = WideInt(__va_arg(ap, WideInt(_)))
						break
					case LM_INTMAX_T:
						i_num = WideInt(__va_arg(ap, intmax_t(_)))
						break
					case LM_PTRDIFF_T:
						i_num = WideInt(__va_arg(ap, ptrdiff_t(_)))
						break
					case LM_PHP_INT_T:
						i_num = WideInt(__va_arg(ap, zend.ZendLong(_)))
						break
					}
				}
				s = ApPhpConv10(i_num, (*fmt) == 'u', &is_negative, &num_buf[1080], &s_len)
				if adjust_precision != 0 {
					for s_len < int(precision) {
						*(g.PreDec(&s)) = '0'
						s_len++
					}
				}
				if (*fmt) != 'u' {
					if is_negative != 0 {
						prefix_char = '-'
					} else if print_sign != 0 {
						prefix_char = '+'
					} else if print_blank != 0 {
						prefix_char = ' '
					}
				}
				break
			case 'o':
				switch modifier {
				default:
					ui_num = UWideInt(__va_arg(ap, uint(_)))
					break
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					ui_num = UWideInt(__va_arg(ap, unsigned__long__int(_)))
					break
				case LM_SIZE_T:
					ui_num = UWideInt(__va_arg(ap, int(_)))
					break
				case LM_LONG_LONG:
					ui_num = UWideInt(__va_arg(ap, UWideInt(_)))
					break
				case LM_INTMAX_T:
					ui_num = UWideInt(__va_arg(ap, uintmax_t(_)))
					break
				case LM_PTRDIFF_T:
					ui_num = UWideInt(__va_arg(ap, ptrdiff_t(_)))
					break
				case LM_PHP_INT_T:
					ui_num = UWideInt(__va_arg(ap, zend.ZendUlong(_)))
					break
				}
				s = ApPhpConvP2(ui_num, 3, *fmt, &num_buf[1080], &s_len)
				if adjust_precision != 0 {
					for s_len < int(precision) {
						*(g.PreDec(&s)) = '0'
						s_len++
					}
				}
				if alternate_form != 0 && (*s) != '0' {
					*(g.PreDec(&s)) = '0'
					s_len++
				}
				break
			case 'x':

			case 'X':
				switch modifier {
				default:
					ui_num = UWideInt(__va_arg(ap, uint(_)))
					break
				case LM_LONG_DOUBLE:
					goto fmt_error
				case LM_LONG:
					ui_num = UWideInt(__va_arg(ap, unsigned__long__int(_)))
					break
				case LM_SIZE_T:
					ui_num = UWideInt(__va_arg(ap, int(_)))
					break
				case LM_LONG_LONG:
					ui_num = UWideInt(__va_arg(ap, UWideInt(_)))
					break
				case LM_INTMAX_T:
					ui_num = UWideInt(__va_arg(ap, uintmax_t(_)))
					break
				case LM_PTRDIFF_T:
					ui_num = UWideInt(__va_arg(ap, ptrdiff_t(_)))
					break
				case LM_PHP_INT_T:
					ui_num = UWideInt(__va_arg(ap, zend.ZendUlong(_)))
					break
				}
				s = ApPhpConvP2(ui_num, 4, *fmt, &num_buf[1080], &s_len)
				if adjust_precision != 0 {
					for s_len < int(precision) {
						*(g.PreDec(&s)) = '0'
						s_len++
					}
				}
				if alternate_form != 0 && ui_num != 0 {
					*(g.PreDec(&s)) = *fmt
					*(g.PreDec(&s)) = '0'
					s_len += 2
				}
				break
			case 's':

			case 'v':
				s = __va_arg(ap, (*byte)(_))
				if s != nil {
					if adjust_precision == 0 {
						s_len = strlen(s)
					} else {
						s_len = strnlen(s, precision)
					}
				} else {
					s = "(null)"
					s_len = 6
				}
				pad_char = ' '
				break
			case 'f':

			case 'F':

			case 'e':

			case 'E':
				switch modifier {
				case LM_LONG_DOUBLE:
					fp_num = float64(__va_arg(ap, long__double(_)))
					break
				case LM_STD:
					fp_num = __va_arg(ap, float64(_))
					break
				default:
					goto fmt_error
				}
				if isnan(fp_num) {
					s = "nan"
					s_len = 3
				} else if isinf(fp_num) {
					s = "inf"
					s_len = 3
				} else {
					if lconv == nil {
						lconv = localeconv()
					}
					s = PhpConvFp(g.Cond((*fmt) == 'f', 'F', *fmt), fp_num, alternate_form, g.Cond(adjust_precision == NO, 6, precision), g.CondF1((*fmt) == 'f', func() __auto__ { return (*lconv).decimal_point }, '.'), &is_negative, &num_buf[1], &s_len)
					if is_negative != 0 {
						prefix_char = '-'
					} else if print_sign != 0 {
						prefix_char = '+'
					} else if print_blank != 0 {
						prefix_char = ' '
					}
				}
				break
			case 'g':

			case 'k':

			case 'G':

			case 'H':
				switch modifier {
				case LM_LONG_DOUBLE:
					fp_num = float64(__va_arg(ap, long__double(_)))
					break
				case LM_STD:
					fp_num = __va_arg(ap, float64(_))
					break
				default:
					goto fmt_error
				}
				if isnan(fp_num) {
					s = "NAN"
					s_len = 3
					break
				} else if isinf(fp_num) {
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
					precision = 6
				} else if precision == 0 {
					precision = 1
				}

				/*
				 * * We use &num_buf[ 1 ], so that we have room for the sign
				 */

				if lconv == nil {
					lconv = localeconv()
				}
				s = PhpGcvt(fp_num, precision, g.CondF2((*fmt) == 'H' || (*fmt) == 'k', '.', func() __auto__ { return (*lconv).decimal_point }), g.Cond((*fmt) == 'G' || (*fmt) == 'H', 'E', 'e'), &num_buf[1])
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
					s[g.PostInc(&s_len)] = '.'
				}
				break
			case 'c':
				char_buf[0] = byte(__va_arg(ap, int(_)))
				s = &char_buf[0]
				s_len = 1
				pad_char = ' '
				break
			case '%':
				char_buf[0] = '%'
				s = &char_buf[0]
				s_len = 1
				pad_char = ' '
				break
			case 'n':
				if is_char != 0 {
					*(__va_arg(ap, (*int)(_))) = int((*zend.SmartString)(xbuf).len_)
				} else {
					*(__va_arg(ap, (*int)(_))) = int((*zend.SmartStr)(xbuf).s.len_)
				}
				goto skip_output
			case 'p':
				if g.SizeOf("char *") <= g.SizeOf("u_wide_int") {
					ui_num = u_wide_int(int(__va_arg(ap, (*byte)(_))))
					s = ApPhpConvP2(ui_num, 4, 'x', &num_buf[1080], &s_len)
					if ui_num != 0 {
						*(g.PreDec(&s)) = 'x'
						*(g.PreDec(&s)) = '0'
						s_len += 2
					}
				} else {
					s = "%p"
					s_len = 2
				}
				pad_char = ' '
				break
			case '0':

				/*
				 * The last character of the format string was %.
				 * We ignore it.
				 */

				continue
			fmt_error:
				zend.ZendError(1<<0, "Illegal length modifier specified '%c' in s[np]printf call", *fmt)
			default:
				char_buf[0] = '%'
				char_buf[1] = *fmt
				s = char_buf
				s_len = 2
				pad_char = ' '
				break
			}
			if prefix_char != '0' {
				*(g.PreDec(&s)) = prefix_char
				s_len++
			}
			if adjust_width != 0 && adjust == RIGHT && int(min_width > s_len) != 0 {
				if pad_char == '0' && prefix_char != '0' {
					if is_char != 0 {
						zend.SmartStringAppendcEx((*zend.SmartString)(xbuf), *s, 0)
					} else {
						zend.SmartStrAppendcEx((*zend.SmartStr)(xbuf), *s, 0)
					}
					s++
					s_len--
					min_width--
				}
				if is_char != 0 {
					zend.SmartStringAlloc((*zend.SmartString)(xbuf), min_width-s_len, 0)
					memset((*zend.SmartString)(xbuf).c+(*zend.SmartString)(xbuf).len_, pad_char, min_width-s_len)
					(*zend.SmartString)(xbuf).len_ += min_width - s_len
				} else {
					zend.SmartStrAlloc((*zend.SmartStr)(xbuf), min_width-s_len, 0)
					memset((*zend.SmartStr)(xbuf).s.val+(*zend.SmartStr)(xbuf).s.len_, pad_char, min_width-s_len)
					(*zend.SmartStr)(xbuf).s.len_ += min_width - s_len
				}
			}

			/*
			 * Print the string s.
			 */

			if is_char != 0 {
				zend.SmartStringAppendlEx((*zend.SmartString)(xbuf), s, s_len, 0)
			} else {
				zend.SmartStrAppendlEx((*zend.SmartStr)(xbuf), s, s_len, 0)
			}
			if adjust_width != 0 && adjust == LEFT && int(min_width > s_len) != 0 {
				if is_char != 0 {
					zend.SmartStringAlloc((*zend.SmartString)(xbuf), min_width-s_len, 0)
					memset((*zend.SmartString)(xbuf).c+(*zend.SmartString)(xbuf).len_, pad_char, min_width-s_len)
					(*zend.SmartString)(xbuf).len_ += min_width - s_len
				} else {
					zend.SmartStrAlloc((*zend.SmartStr)(xbuf), min_width-s_len, 0)
					memset((*zend.SmartStr)(xbuf).s.val+(*zend.SmartStr)(xbuf).s.len_, pad_char, min_width-s_len)
					(*zend.SmartStr)(xbuf).s.len_ += min_width - s_len
				}
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

/* }}} */

func PhpPrintfToSmartString(buf *zend.SmartString, format *byte, ap va_list) {
	XbufFormatConverter(buf, 1, format, ap)
}

/* }}} */

func PhpPrintfToSmartStr(buf *zend.SmartStr, format *byte, ap va_list) {
	XbufFormatConverter(buf, 0, format, ap)
}

/* }}} */

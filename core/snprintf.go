// <<generate>>

package core

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/snprintf.h>

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
   | Author: Stig Sæther Bakken <ssb@php.net>                             |
   |         Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define SNPRINTF_H

type BoolInt = int
type BooleanE = int

const (
	NO  = 0
	YES = 1
)

var Php0cvt func(value float64, ndigit int, dec_point byte, exponent byte, buf *byte) *byte

// #define slprintf       ap_php_slprintf

// #define vslprintf       ap_php_vslprintf

// #define snprintf       ap_php_snprintf

// #define vsnprintf       ap_php_vsnprintf

type LengthModifierE = int

const (
	LM_STD = 0
	LM_INTMAX_T
	LM_PTRDIFF_T
	LM_LONG_LONG
	LM_SIZE_T
	LM_LONG
	LM_LONG_DOUBLE
	LM_PHP_INT_T
)

// #define WIDE_INT       long long

type WideInt = long__long
type UWideInt = unsigned__long__long

/* The maximum precision that's allowed for float conversion. Does not include
 * decimal separator, exponent, sign, terminator. Currently does not affect
 * the modes e/f, only g/k/H, as those have a different limit enforced at
 * another level (see NDIG in php_conv_fp()).
 * Applies to the formatting functions of both spprintf.c and snprintf.c, which
 * use equally sized buffers of MAX_BUF_SIZE = 512 to hold the result of the
 * call to php_gcvt().
 * This should be reasonably smaller than MAX_BUF_SIZE (I think MAX_BUF_SIZE - 9
 * should be enough, but let's give some more space) */

// #define FORMAT_CONV_MAX_PRECISION       500

// Source: <main/snprintf.c>

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
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

// #define _GNU_SOURCE

// # include "php.h"

// # include < zend_strtod . h >

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

/*
 * Copyright (c) 2002, 2006 Todd C. Miller <Todd.Miller@courtesan.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 *
 * Sponsored in part by the Defense Advanced Research Projects
 * Agency (DARPA) and Air Force Research Laboratory, Air Force
 * Materiel Command, USAF, under agreement number F39502-99-1-0512.
 */

func __cvt(value float64, ndigit int, decpt *int, sign *int, fmode int, pad int) *byte {
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
		if g.Assign(&rve, g.Assign(&s, (*byte)(zend.Malloc(g.Cond(ndigit != 0, siz, 2))))) == nil {
			return nil
		}
		g.PostInc(&(*rve)) = '0'
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
			return strdup(g.Cond(c == 'I', "INF", "NAN"))
		}

		/* Make a local copy and adjust rve to be in terms of s */

		if pad != 0 && fmode != 0 {
			siz += *decpt
		}
		if g.Assign(&s, (*byte)(zend.Malloc(siz+1))) == nil {
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
		for g.PreDec(&siz) {
			g.PostInc(&(*rve)) = '0'
		}
		*rve = '0'
	}
	return s
}

/* }}} */

func PhpEcvt(value float64, ndigit int, decpt *int, sign *int) *byte {
	return __cvt(value, ndigit, decpt, sign, 0, 1)
}

/* }}} */

func PhpFcvt(value float64, ndigit int, decpt *int, sign *int) *byte {
	return __cvt(value, ndigit, decpt, sign, 1, 1)
}

/* }}} */

func PhpGcvt(value float64, ndigit int, dec_point byte, exponent byte, buf *byte) *byte {
	var digits *byte
	var dst *byte
	var src *byte
	var i int
	var decpt int
	var sign int
	var mode int = g.Cond(ndigit >= 0, 2, 0)
	if mode == 0 {
		ndigit = 17
	}
	digits = zend.ZendDtoa(value, mode, ndigit, &decpt, &sign, nil)
	if decpt == 9999 {

		/*
		 * Infinity or NaN, convert to inf or nan with sign.
		 * We assume the buffer is at least ndigit long.
		 */

		ApPhpSnprintf(buf, ndigit+1, "%s%s", g.Cond(sign != 0 && (*digits) == 'I', "-", ""), g.Cond((*digits) == 'I', "INF", "NAN"))
		zend.ZendFreedtoa(digits)
		return buf
	}
	dst = buf
	if sign != 0 {
		g.PostInc(&(*dst)) = '-'
	}
	if decpt >= 0 && decpt > ndigit || decpt < -3 {

		/* exponential format (e.g. 1.2345e+13) */

		if g.PreDec(&decpt) < 0 {
			sign = 1
			decpt = -decpt
		} else {
			sign = 0
		}
		src = digits
		*src++
		g.PostInc(&(*dst)) = (*src) - 1
		g.PostInc(&(*dst)) = dec_point
		if (*src) == '0' {
			g.PostInc(&(*dst)) = '0'
		} else {
			for {
				*src++
				g.PostInc(&(*dst)) = (*src) - 1
				if (*src) == '0' {
					break
				}
			}
		}
		g.PostInc(&(*dst)) = exponent
		if sign != 0 {
			g.PostInc(&(*dst)) = '-'
		} else {
			g.PostInc(&(*dst)) = '+'
		}
		if decpt < 10 {
			g.PostInc(&(*dst)) = '0' + decpt
			*dst = '0'
		} else {

			/* XXX - optimize */

			sign = decpt
			i = 0
			for ; g.AssignOp(&sign, "/=", 10) != 0; i++ {

			}
			dst[i+1] = '0'
			for decpt != 0 {
				dst[g.PostDec(&i)] = '0' + decpt%10
				decpt /= 10
			}
		}
	} else if decpt < 0 {

		/* standard format 0. */

		g.PostInc(&(*dst)) = '0'
		g.PostInc(&(*dst)) = dec_point
		for {
			g.PostInc(&(*dst)) = '0'
			if g.PreInc(&decpt) >= 0 {
				break
			}
		}
		src = digits
		for (*src) != '0' {
			*src++
			g.PostInc(&(*dst)) = (*src) - 1
		}
		*dst = '0'
	} else {

		/* standard format */

		i = 0
		src = digits
		for ; i < decpt; i++ {
			if (*src) != '0' {
				*src++
				g.PostInc(&(*dst)) = (*src) - 1
			} else {
				g.PostInc(&(*dst)) = '0'
			}
		}
		if (*src) != '0' {
			if src == digits {
				g.PostInc(&(*dst)) = '0'
			}
			g.PostInc(&(*dst)) = dec_point
			for i = decpt; digits[i] != '0'; i++ {
				g.PostInc(&(*dst)) = digits[i]
			}
		}
		*dst = '0'
	}
	zend.ZendFreedtoa(digits)
	return buf
}

/* }}} */

// #define FALSE       0

// #define TRUE       1

// #define NUL       '\0'

// #define INT_NULL       ( ( int * ) 0 )

// #define S_NULL       "(null)"

// #define S_NULL_LEN       6

// #define FLOAT_DIGITS       6

// #define EXPONENT_LENGTH       10

/*
 * Convert num to its decimal format.
 * Return value:
 *   - a pointer to a string containing the number (no sign)
 *   - len contains the length of the string
 *   - is_negative is set to TRUE or FALSE depending on the sign
 *     of the number (always set to FALSE if is_unsigned is TRUE)
 *
 * The caller provides a buffer for the string: that is the buf_end argument
 * which is a pointer to the END of the buffer + 1 (i.e. if the buffer
 * is declared as buf[ 100 ], buf_end should be &buf[ 100 ])
 */

func ApPhpConv10(num WideInt, is_unsigned BoolInt, is_negative *BoolInt, buf_end *byte, len_ *int) *byte {
	var p *byte = buf_end
	var magnitude UWideInt
	if is_unsigned != 0 {
		magnitude = UWideInt(num)
		*is_negative = 0
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
		*(g.PreDec(&p)) = byte(magnitude - new_magnitude*10 + '0')
		magnitude = new_magnitude
		if !magnitude {
			break
		}
	}
	*len_ = buf_end - p
	return p
}

/* }}} */

// #define NDIG       320

/*
 * Convert a floating point number to a string formats 'f', 'e' or 'E'.
 * The result is placed in buf, and len denotes the length of the string
 * The sign is returned in the is_negative argument (and is not placed
 * in buf).
 */

func PhpConvFp(format byte, num float64, add_dp BooleanE, precision int, dec_point byte, is_negative *BoolInt, buf *byte, len_ *int) *byte {
	var s *byte = buf
	var p *byte
	var p_orig *byte
	var decimal_point int
	if precision >= 320-1 {
		precision = 320 - 2
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
		*is_negative = 0
		zend.Free(p_orig)
		return buf
	}
	if format == 'F' {
		if decimal_point <= 0 {
			if num != 0 || precision > 0 {
				g.PostInc(&(*s)) = '0'
				if precision > 0 {
					g.PostInc(&(*s)) = dec_point
					for g.PostInc(&decimal_point) < 0 {
						g.PostInc(&(*s)) = '0'
					}
				} else if add_dp != 0 {
					g.PostInc(&(*s)) = dec_point
				}
			}
		} else {
			var addz int = g.Cond(decimal_point >= 320, decimal_point-320+1, 0)
			decimal_point -= addz
			for g.PostDec(&decimal_point) > 0 {
				*p++
				g.PostInc(&(*s)) = (*p) - 1
			}
			for g.PostDec(&addz) > 0 {
				g.PostInc(&(*s)) = '0'
			}
			if precision > 0 || add_dp != 0 {
				g.PostInc(&(*s)) = dec_point
			}
		}
	} else {
		*p++
		g.PostInc(&(*s)) = (*p) - 1
		if precision > 0 || add_dp != 0 {
			g.PostInc(&(*s)) = '.'
		}
	}

	/*
	 * copy the rest of p, the NUL is NOT copied
	 */

	for *p {
		*p++
		g.PostInc(&(*s)) = (*p) - 1
	}
	if format != 'F' {
		var temp []byte
		var t_len int
		var exponent_is_negative BoolInt
		g.PostInc(&(*s)) = format
		decimal_point--
		if decimal_point != 0 {
			p = ApPhpConv10(WideInt(decimal_point), 0, &exponent_is_negative, &temp[10], &t_len)
			if exponent_is_negative != 0 {
				g.PostInc(&(*s)) = '-'
			} else {
				g.PostInc(&(*s)) = '+'
			}

			/*
			 * Make sure the exponent has at least 2 digits
			 */

			for g.PostDec(&t_len) {
				*p++
				g.PostInc(&(*s)) = (*p) - 1
			}

			/*
			 * Make sure the exponent has at least 2 digits
			 */

		} else {
			g.PostInc(&(*s)) = '+'
			g.PostInc(&(*s)) = '0'
		}
	}
	*len_ = s - buf
	zend.Free(p_orig)
	return buf
}

/* }}} */

func ApPhpConvP2(num UWideInt, nbits int, format byte, buf_end *byte, len_ *int) *byte {
	var mask int = (1 << nbits) - 1
	var p *byte = buf_end
	var low_digits []byte = "0123456789abcdef"
	var upper_digits []byte = "0123456789ABCDEF"
	var digits *byte = g.Cond(format == 'X', upper_digits, low_digits)
	for {
		*(g.PreDec(&p)) = digits[num&mask]
		num >>= nbits
		if !num {
			break
		}
	}
	*len_ = buf_end - p
	return p
}

/* }}} */

// #define NUM_BUF_SIZE       2048

/*
 * Descriptor for buffer area
 */

// @type BufArea struct
type Buffy = BufArea

/*
 * The INS_CHAR macro inserts a character in the buffer and writes
 * the buffer back to disk if necessary
 * It uses the char pointers sp and bep:
 *      sp points to the next available character in the buffer
 *      bep points to the end-of-buffer+1
 * While using this macro, note that the nextb pointer is NOT updated.
 *
 * NOTE: Evaluation of the c argument should not have any side-effects
 */

// #define INS_CHAR(c,sp,bep,cc) { if ( sp < bep ) { * sp ++ = c ; } cc ++ ; }

// #define NUM(c) ( c - '0' )

// #define STR_TO_DEC(str,num) num = NUM ( * str ++ ) ; while ( isdigit ( ( int ) * str ) ) { num *= 10 ; num += NUM ( * str ++ ) ; }

/*
 * This macro does zero padding so that the precision
 * requirement is satisfied. The padding is done by
 * adding '0's to the left of the string that is going
 * to be printed.
 */

// #define FIX_PRECISION(adjust,precision,s,s_len) if ( adjust ) while ( s_len < ( size_t ) precision ) { * -- s = '0' ; s_len ++ ; }

/*
 * Macro that does padding. The padding is done by printing
 * the character ch.
 */

// #define PAD(width,len,ch) do { INS_CHAR ( ch , sp , bep , cc ) ; width -- ; } while ( ( size_t ) width > len )

/*
 * Prefix the character ch to the string str
 * Increase length
 * Set the has_prefix flag
 */

// #define PREFIX(str,length,ch) * -- str = ch ; length ++ ; has_prefix = YES

/*
 * Do format conversion placing the output in buffer
 */

func FormatConverter(odp *Buffy, fmt *byte, ap va_list) int {
	var sp *byte
	var bep *byte
	var cc int = 0
	var i int
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
				g.PostInc(&(*sp)) = *fmt
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
					min_width = g.PostInc(&(*fmt)) - '0'
					for isdigit(int(*fmt)) {
						min_width *= 10
						min_width += g.PostInc(&(*fmt)) - '0'
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
						precision = g.PostInc(&(*fmt)) - '0'
						for isdigit(int(*fmt)) {
							precision *= 10
							precision += g.PostInc(&(*fmt)) - '0'
						}
					} else if (*fmt) == '*' {
						precision = __va_arg(ap, int(_))
						fmt++
						if precision < 0 {
							precision = 0
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
				fmt++
				modifier = LM_PHP_INT_T
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
				s = ApPhpConv10(i_num, (*fmt) == 'u', &is_negative, &num_buf[2048], &s_len)
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
				s = ApPhpConvP2(ui_num, 3, *fmt, &num_buf[2048], &s_len)
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
				s = ApPhpConvP2(ui_num, 4, *fmt, &num_buf[2048], &s_len)
				if adjust_precision != 0 {
					for s_len < int(precision) {
						*(g.PreDec(&s)) = '0'
						s_len++
					}
				}
				if alternate_form != 0 && i_num != 0 {
					*(g.PreDec(&s)) = *fmt
					*(g.PreDec(&s)) = '0'
					s_len += 2
				}
				break
			case 's':

			case 'v':
				s = __va_arg(ap, (*byte)(_))
				if s != nil {
					s_len = strlen(s)
					if adjust_precision != 0 && int(precision < s_len) != 0 {
						s_len = precision
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
					s = "NAN"
					s_len = 3
				} else if isinf(fp_num) {
					s = "INF"
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
				*(__va_arg(ap, (*int)(_))) = cc
				goto skip_output
			case 'p':
				if g.SizeOf("char *") <= g.SizeOf("u_wide_int") {
					ui_num = u_wide_int(int(__va_arg(ap, (*byte)(_))))
					s = ApPhpConvP2(ui_num, 4, 'x', &num_buf[2048], &s_len)
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
					if sp < bep {
						g.PostInc(&(*sp)) = *s
					}
					cc++
					s++
					s_len--
					min_width--
				}
				for {
					if sp < bep {
						g.PostInc(&(*sp)) = pad_char
					}
					cc++
					min_width--
					if int(min_width > s_len) == 0 {
						break
					}
				}
			}

			/*
			 * Print the string s.
			 */

			for i = s_len; i != 0; i-- {
				if sp < bep {
					g.PostInc(&(*sp)) = *s
				}
				cc++
				s++
			}
			if adjust_width != 0 && adjust == LEFT && int(min_width > s_len) != 0 {
				for {
					if sp < bep {
						g.PostInc(&(*sp)) = pad_char
					}
					cc++
					min_width--
					if int(min_width > s_len) == 0 {
						break
					}
				}
			}
			if free_zcopy != 0 {
				zend.ZvalPtrDtorStr(&zcopy)
			}
		}
	skip_output:
		fmt++
	}
	odp.SetNextb(sp)
	return cc
}

/* }}} */

func StrxPrintv(ccp *int, buf *byte, len_ int, format *byte, ap va_list) {
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

/* }}} */

func ApPhpSlprintf(buf *byte, len_ int, format string, _ ...any) int {
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

/* }}} */

func ApPhpVslprintf(buf *byte, len_ int, format *byte, ap va_list) int {
	var cc int
	StrxPrintv(&cc, buf, len_, format, ap)
	if int(cc >= len_) != 0 {
		cc = int(len_ - 1)
		buf[cc] = '0'
	}
	return cc
}

/* }}} */

func ApPhpSnprintf(buf *byte, len_ int, format string, _ ...any) int {
	var cc int
	var ap va_list
	va_start(ap, format)
	StrxPrintv(&cc, buf, len_, format, ap)
	va_end(ap)
	return cc
}

/* }}} */

func ApPhpVsnprintf(buf *byte, len_ int, format *byte, ap va_list) int {
	var cc int
	StrxPrintv(&cc, buf, len_, format, ap)
	return cc
}

/* }}} */

func ApPhpVasprintf(buf **byte, format *byte, ap va_list) int {
	var ap2 va_list
	var cc int
	memcpy(&ap2, &ap, g.SizeOf("va_list"))
	cc = ApPhpVsnprintf(nil, 0, format, ap2)
	va_end(ap2)
	*buf = nil
	if cc >= 0 {
		if g.Assign(&(*buf), zend.Malloc(g.PreInc(&cc))) != nil {
			if g.Assign(&cc, ApPhpVsnprintf(*buf, cc, format, ap)) < 0 {
				zend.Free(*buf)
				*buf = nil
			}
		}
	}
	return cc
}

/* }}} */

func ApPhpAsprintf(buf **byte, format *byte, _ ...any) int {
	var cc int
	var ap va_list
	va_start(ap, format)
	cc = vasprintf(buf, format, ap)
	va_end(ap)
	return cc
}

/* }}} */

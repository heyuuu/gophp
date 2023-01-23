// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/scanf.h>

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
   | Author: Clayton Collie <clcollie@mindspring.com>                     |
   +----------------------------------------------------------------------+
*/

// #define SCANF_H

// #define SCAN_MAX_ARGS       0xFF

/* passed to (f|s)scanf. This is an artificial   */

// #define SCAN_SUCCESS       SUCCESS

// #define SCAN_ERROR_EOF       - 1

/* can be caused by bad parameters or format*/

// #define SCAN_ERROR_INVALID_FORMAT       ( SCAN_ERROR_EOF - 1 )

// #define SCAN_ERROR_VAR_PASSED_BYVAL       ( SCAN_ERROR_INVALID_FORMAT - 1 )

// #define SCAN_ERROR_WRONG_PARAM_COUNT       ( SCAN_ERROR_VAR_PASSED_BYVAL - 1 )

// #define SCAN_ERROR_INTERNAL       ( SCAN_ERROR_WRONG_PARAM_COUNT - 1 )

/*
 * The following are here solely for the benefit of the scanf type functions
 * e.g. fscanf
 */

// Source: <ext/standard/scanf.c>

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
   | Author: Clayton Collie <clcollie@mindspring.com>                     |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < limits . h >

// # include < ctype . h >

// # include "php.h"

// # include "php_variables.h"

// # include < locale . h >

// # include "zend_execute.h"

// # include "zend_operators.h"

// # include "zend_strtod.h"

// # include "php_globals.h"

// # include "basic_functions.h"

// # include "scanf.h"

/*
 * Flag values used internally by [f|s]canf.
 */

// #define SCAN_NOSKIP       0x1

// #define SCAN_SUPPRESS       0x2

// #define SCAN_UNSIGNED       0x4

// #define SCAN_WIDTH       0x8

// #define SCAN_SIGNOK       0x10

// #define SCAN_NODIGITS       0x20

// #define SCAN_NOZERO       0x40

// #define SCAN_XOK       0x80

// #define SCAN_PTOK       0x100

// #define SCAN_EXPOK       0x200

// #define UCHAR(x) ( zend_uchar ) ( x )

/*
 * The following structure contains the information associated with
 * a character set.
 */

/*
 * Declarations for functions used only in this file.
 */

/* {{{ BuildCharSet
 *----------------------------------------------------------------------
 *
 * BuildCharSet --
 *
 *    This function examines a character set format specification
 *    and builds a CharSet containing the individual characters and
 *    character ranges specified.
 *
 * Results:
 *    Returns the next format position.
 *
 * Side effects:
 *    Initializes the charset.
 *
 *----------------------------------------------------------------------
 */

func BuildCharSet(cset *CharSet, format *byte) *byte {
	var ch *byte
	var start byte
	var nranges int
	var end *byte
	memset(cset, 0, g.SizeOf("CharSet"))
	ch = format
	if (*ch) == '^' {
		cset.SetExclude(1)
		format++
		ch = format
	}
	end = format + 1

	/*
	 * Find the close bracket so we can overallocate the set.
	 */

	if (*ch) == ']' {
		end++
		ch = end - 1
	}
	nranges = 0
	for (*ch) != ']' {
		if (*ch) == '-' {
			nranges++
		}
		end++
		ch = end - 1
	}
	cset.SetChars((*byte)(zend._safeEmalloc(g.SizeOf("char"), end-format-1, 0)))
	if nranges > 0 {
		cset.SetRanges((*__struct__Range)(zend._safeEmalloc(g.SizeOf("struct Range"), nranges, 0)))
	} else {
		cset.SetRanges(nil)
	}

	/*
	 * Now build the character set.
	 */

	cset.SetNranges(0)
	cset.SetNchars(cset.GetNranges())
	format++
	ch = format - 1
	start = *ch
	if (*ch) == ']' || (*ch) == '-' {
		cset.GetChars()[g.PostInc(&(cset.GetNchars()))] = *ch
		format++
		ch = format - 1
	}
	for (*ch) != ']' {
		if (*format) == '-' {

			/*
			 * This may be the first character of a range, so don't add
			 * it yet.
			 */

			start = *ch

			/*
			 * This may be the first character of a range, so don't add
			 * it yet.
			 */

		} else if (*ch) == '-' {

			/*
			 * Check to see if this is the last character in the set, in which
			 * case it is not a range and we should add the previous character
			 * as well as the dash.
			 */

			if (*format) == ']' {
				cset.GetChars()[g.PostInc(&(cset.GetNchars()))] = start
				cset.GetChars()[g.PostInc(&(cset.GetNchars()))] = *ch
			} else {
				format++
				ch = format - 1

				/*
				 * Check to see if the range is in reverse order.
				 */

				if start < (*ch) {
					cset.GetRanges()[cset.GetNranges()].start = start
					cset.GetRanges()[cset.GetNranges()].end = *ch
				} else {
					cset.GetRanges()[cset.GetNranges()].start = *ch
					cset.GetRanges()[cset.GetNranges()].end = start
				}
				cset.GetNranges()++
			}

			/*
			 * Check to see if this is the last character in the set, in which
			 * case it is not a range and we should add the previous character
			 * as well as the dash.
			 */

		} else {
			cset.GetChars()[g.PostInc(&(cset.GetNchars()))] = *ch
		}
		format++
		ch = format - 1
	}
	return format
}

/* }}} */

func CharInSet(cset *CharSet, c int) int {
	var ch byte = byte(c)
	var i int
	var match int = 0
	for i = 0; i < cset.GetNchars(); i++ {
		if cset.GetChars()[i] == ch {
			match = 1
			break
		}
	}
	if match == 0 {
		for i = 0; i < cset.GetNranges(); i++ {
			if cset.GetRanges()[i].start <= ch && ch <= cset.GetRanges()[i].end {
				match = 1
				break
			}
		}
	}
	if cset.GetExclude() != 0 {
		return !match
	} else {
		return match
	}
}

/* }}} */

func ReleaseCharSet(cset *CharSet) {
	zend._efree((*byte)(cset.GetChars()))
	if cset.GetRanges() != nil {
		zend._efree((*byte)(cset.GetRanges()))
	}
}

/* }}} */

func ValidateFormat(format *byte, numVars int, totalSubs *int) int {
	// #define STATIC_LIST_SIZE       16

	var gotXpg int
	var gotSequential int
	var value int
	var i int
	var flags int
	var end *byte
	var ch *byte = nil
	var staticAssign []int
	var nassign *int = staticAssign
	var objIndex int
	var xpgSize int
	var nspace int = 16

	/*
	 * Initialize an array that records the number of times a variable
	 * is assigned to by the format string.  We use this to detect if
	 * a variable is multiply assigned or left unassigned.
	 */

	if numVars > nspace {
		nassign = (*int)(zend._safeEmalloc(g.SizeOf("int"), numVars, 0))
		nspace = numVars
	}
	for i = 0; i < nspace; i++ {
		nassign[i] = 0
	}
	gotSequential = 0
	gotXpg = gotSequential
	objIndex = gotXpg
	xpgSize = objIndex
	for (*format) != '0' {
		format++
		ch = format - 1
		flags = 0
		if (*ch) != '%' {
			continue
		}
		format++
		ch = format - 1
		if (*ch) == '%' {
			continue
		}
		if (*ch) == '*' {
			flags |= 0x2
			format++
			ch = format - 1
			goto xpgCheckDone
		}
		if isdigit(int(*ch)) {

			/*
			 * Check for an XPG3-style %n$ specification.  Note: there
			 * must not be a mixture of XPG3 specs and non-XPG3 specs
			 * in the same format string.
			 */

			value = strtoull(format-1, &end, 10)
			if (*end) != '$' {
				goto notXpg
			}
			format = end + 1
			format++
			ch = format - 1
			gotXpg = 1
			if gotSequential != 0 {
				goto mixedXPG
			}
			objIndex = value - 1
			if objIndex < 0 || numVars != 0 && objIndex >= numVars {
				goto badIndex
			} else if numVars == 0 {

				/*
				 * In the case where no vars are specified, the user can
				 * specify %9999$ legally, so we have to consider special
				 * rules for growing the assign array.  'value' is
				 * guaranteed to be > 0.
				 */

				if value > 0xff {
					goto badIndex
				}
				if xpgSize > value {
					xpgSize = xpgSize
				} else {
					xpgSize = value
				}
			}
			goto xpgCheckDone
		}
	notXpg:
		gotSequential = 1
		if gotXpg != 0 {
		mixedXPG:
			core.PhpErrorDocref(nil, 1<<1, "%s", "cannot mix \"%\" and \"%n$\" conversion specifiers")
			goto error
		}
	xpgCheckDone:

		/*
		 * Parse any width specifier.
		 */

		if isdigit(zend_uchar(*ch)) {
			value = strtoull(format-1, &format, 10)
			flags |= 0x8
			format++
			ch = format - 1
		}

		/*
		 * Ignore size specifier.
		 */

		if (*ch) == 'l' || (*ch) == 'L' || (*ch) == 'h' {
			format++
			ch = format - 1
		}
		if (flags&0x2) == 0 && numVars != 0 && objIndex >= numVars {
			goto badIndex
		}

		/*
		 * Handle the various field types.
		 */

		switch *ch {
		case 'n':

		case 'd':

		case 'D':

		case 'i':

		case 'o':

		case 'x':

		case 'X':

		case 'u':

		case 'f':

		case 'e':

		case 'E':

		case 'g':

		case 's':
			break
		case 'c':

			/* we differ here with the TCL implementation in allowing for */

			break
		case '[':
			if (*format) == '0' {
				goto badSet
			}
			format++
			ch = format - 1
			if (*ch) == '^' {
				if (*format) == '0' {
					goto badSet
				}
				format++
				ch = format - 1
			}
			if (*ch) == ']' {
				if (*format) == '0' {
					goto badSet
				}
				format++
				ch = format - 1
			}
			for (*ch) != ']' {
				if (*format) == '0' {
					goto badSet
				}
				format++
				ch = format - 1
			}
			break
		badSet:
			core.PhpErrorDocref(nil, 1<<1, "Unmatched [ in format string")
			goto error
		default:
			core.PhpErrorDocref(nil, 1<<1, "Bad scan conversion character \"%c\"", *ch)
			goto error
		}
		if (flags & 0x2) == 0 {
			if objIndex >= nspace {

				/*
				 * Expand the nassign buffer.  If we are using XPG specifiers,
				 * make sure that we grow to a large enough size.  xpgSize is
				 * guaranteed to be at least one larger than objIndex.
				 */

				value = nspace
				if xpgSize != 0 {
					nspace = xpgSize
				} else {
					nspace += 16
				}
				if nassign == staticAssign {
					nassign = any(zend._safeEmalloc(nspace, g.SizeOf("int"), 0))
					for i = 0; i < 16; i++ {
						nassign[i] = staticAssign[i]
					}
				} else {
					nassign = any(zend._erealloc(any(nassign), nspace*g.SizeOf("int")))
				}
				for i = value; i < nspace; i++ {
					nassign[i] = 0
				}
			}
			nassign[objIndex]++
			objIndex++
		}
	}

	/*
	 * Verify that all of the variable were assigned exactly once.
	 */

	if numVars == 0 {
		if xpgSize != 0 {
			numVars = xpgSize
		} else {
			numVars = objIndex
		}
	}
	if totalSubs != nil {
		*totalSubs = numVars
	}
	for i = 0; i < numVars; i++ {
		if nassign[i] > 1 {
			core.PhpErrorDocref(nil, 1<<1, "%s", "Variable is assigned by multiple \"%n$\" conversion specifiers")
			goto error
		} else if xpgSize == 0 && nassign[i] == 0 {

			/*
			 * If the space is empty, and xpgSize is 0 (means XPG wasn't
			 * used, and/or numVars != 0), then too many vars were given
			 */

			core.PhpErrorDocref(nil, 1<<1, "Variable is not assigned by any conversion specifiers")
			goto error
		}
	}
	if nassign != staticAssign {
		zend._efree((*byte)(nassign))
	}
	return zend.SUCCESS
badIndex:
	if gotXpg != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%s", "\"%n$\" argument index out of range")
	} else {
		core.PhpErrorDocref(nil, 1<<1, "Different numbers of variable names and field specifiers")
	}
error:
	if nassign != staticAssign {
		zend._efree((*byte)(nassign))
	}
	return -1 - 1
}

/* }}} */

func PhpSscanfInternal(string *byte, format *byte, argCount int, args *zend.Zval, varStart int, return_value *zend.Zval) int {
	var numVars int
	var nconversions int
	var totalVars int = -1
	var i int
	var result int
	var value zend.ZendLong
	var objIndex int
	var end *byte
	var baseString *byte
	var current *zend.Zval
	var op byte = 0
	var base int = 0
	var underflow int = 0
	var width int
	var fn func() zend.ZendLong = nil
	var ch *byte
	var sch byte
	var flags int
	var buf []byte

	/* do some sanity checking */

	if varStart > argCount || varStart < 0 {
		varStart = 0xff + 1
	}
	numVars = argCount - varStart
	if numVars < 0 {
		numVars = 0
	}

	/*
	 * Check for errors in the format string.
	 */

	if ValidateFormat(format, numVars, &totalVars) != zend.SUCCESS {
		ScanSetErrorReturn(numVars, return_value)
		return -1 - 1
	}
	if numVars != 0 {
		objIndex = varStart
	} else {
		objIndex = 0
	}

	/*
	 * If any variables are passed, make sure they are all passed by reference
	 */

	if numVars != 0 {
		for i = varStart; i < argCount; i++ {
			if args[i].u1.v.type_ != 10 {
				core.PhpErrorDocref(nil, 1<<1, "Parameter %d must be passed by reference", i)
				ScanSetErrorReturn(numVars, return_value)
				return -1 - 1 - 1
			}
		}
	}

	/*
	 * Allocate space for the result objects. Only happens when no variables
	 * are specified
	 */

	if numVars == 0 {
		var tmp zend.Zval

		/* allocate an array for return */

		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for i = 0; i < totalVars; i++ {
			&tmp.u1.type_info = 1
			if zend.AddNextIndexZval(return_value, &tmp) == zend.FAILURE {
				ScanSetErrorReturn(0, return_value)
				return zend.FAILURE
			}
		}
		varStart = 0
	}
	baseString = string

	/*
	 * Iterate over the format string filling in the result objects until
	 * we reach the end of input, the end of the format string, or there
	 * is a mismatch.
	 */

	nconversions = 0

	/* note ! - we need to limit the loop for objIndex to keep it in bounds */

	for (*format) != '0' {
		format++
		ch = format - 1
		flags = 0

		/*
		 * If we see whitespace in the format, skip whitespace in the string.
		 */

		if isspace(int(*ch)) {
			sch = *string
			for isspace(int(sch)) {
				if (*string) == '0' {
					goto done
				}
				string++
				sch = *string
			}
			continue
		}
		if (*ch) != '%' {
		literal:
			if (*string) == '0' {
				underflow = 1
				goto done
			}
			sch = *string
			string++
			if (*ch) != sch {
				goto done
			}
			continue
		}
		format++
		ch = format - 1
		if (*ch) == '%' {
			goto literal
		}

		/*
		 * Check for assignment suppression ('*') or an XPG3-style
		 * assignment ('%n$').
		 */

		if (*ch) == '*' {
			flags |= 0x2
			format++
			ch = format - 1
		} else if isdigit(zend_uchar(*ch)) {
			value = strtoull(format-1, &end, 10)
			if (*end) == '$' {
				format = end + 1
				format++
				ch = format - 1
				objIndex = varStart + value - 1
			}
		}

		/*
		 * Parse any width specifier.
		 */

		if isdigit(zend_uchar(*ch)) {
			width = strtoull(format-1, &format, 10)
			format++
			ch = format - 1
		} else {
			width = 0
		}

		/*
		 * Ignore size specifier.
		 */

		if (*ch) == 'l' || (*ch) == 'L' || (*ch) == 'h' {
			format++
			ch = format - 1
		}

		/*
		 * Handle the various field types.
		 */

		switch *ch {
		case 'n':
			if (flags & 0x2) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + g.PostInc(&objIndex)
					for {
						r.Assert(current.u1.v.type_ == 10)
						for {
							var _zv *zend.Zval = current
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefLong(ref, zend_long(string-baseString))
								break
							}
							_zv = &ref.val
							zend.ZvalPtrDtor(_zv)
							var __z *zend.Zval = _zv
							__z.value.lval = zend_long(string - baseString)
							__z.u1.type_info = 4
							break
						}
						break
					}
				} else {
					zend.AddIndexLong(return_value, g.PostInc(&objIndex), string-baseString)
				}
			}
			nconversions++
			continue
		case 'd':

		case 'D':
			op = 'i'
			base = 10
			fn = (func() zend.ZendLong)(strtoll)
			break
		case 'i':
			op = 'i'
			base = 0
			fn = (func() zend.ZendLong)(strtoll)
			break
		case 'o':
			op = 'i'
			base = 8
			fn = (func() zend.ZendLong)(strtoll)
			break
		case 'x':

		case 'X':
			op = 'i'
			base = 16
			fn = (func() zend.ZendLong)(strtoll)
			break
		case 'u':
			op = 'i'
			base = 10
			flags |= 0x4
			fn = (func() zend.ZendLong)(strtoull)
			break
		case 'f':

		case 'e':

		case 'E':

		case 'g':
			op = 'f'
			break
		case 's':
			op = 's'
			break
		case 'c':
			op = 's'
			flags |= 0x1

			/*-cc-*/

			if 0 == width {
				width = 1
			}

			/*-cc-*/

			break
		case '[':
			op = '['
			flags |= 0x1
			break
		}

		/*
		 * At this point, we will need additional characters from the
		 * string to proceed.
		 */

		if (*string) == '0' {
			underflow = 1
			goto done
		}

		/*
		 * Skip any leading whitespace at the beginning of a field unless
		 * the format suppresses this behavior.
		 */

		if (flags & 0x1) == 0 {
			for (*string) != '0' {
				sch = *string
				if !(isspace(int(sch))) {
					break
				}
				string++
			}
			if (*string) == '0' {
				underflow = 1
				goto done
			}
		}

		/*
		 * Perform the requested scanning operation.
		 */

		switch op {
		case 'c':

		case 's':

			/*
			 * Scan a string up to width characters or whitespace.
			 */

			if width == 0 {
				width = int(^0)
			}
			end = string
			for (*end) != '0' {
				sch = *end
				if isspace(int(sch)) {
					break
				}
				end++
				if g.PreDec(&width) == 0 {
					break
				}
			}
			if (flags & 0x2) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + g.PostInc(&objIndex)
					for {
						r.Assert(current.u1.v.type_ == 10)
						for {
							var _zv *zend.Zval = current
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefStringl(ref, string, end-string)
								break
							}
							_zv = &ref.val
							zend.ZvalPtrDtor(_zv)
							var __z *zend.Zval = _zv
							var __s *zend.ZendString = zend.ZendStringInit(string, end-string, 0)
							__z.value.str = __s
							__z.u1.type_info = 6 | 1<<0<<8
							break
						}
						break
					}
				} else {
					zend.AddIndexStringl(return_value, g.PostInc(&objIndex), string, end-string)
				}
			}
			string = end
			break
		case '[':
			var cset CharSet
			if width == 0 {
				width = int(^0)
			}
			end = string
			format = BuildCharSet(&cset, format)
			for (*end) != '0' {
				sch = *end
				if CharInSet(&cset, int(sch)) == 0 {
					break
				}
				end++
				if g.PreDec(&width) == 0 {
					break
				}
			}
			ReleaseCharSet(&cset)
			if string == end {

				/*
				 * Nothing matched the range, stop processing
				 */

				goto done

				/*
				 * Nothing matched the range, stop processing
				 */

			}
			if (flags & 0x2) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + g.PostInc(&objIndex)
					for {
						r.Assert(current.u1.v.type_ == 10)
						for {
							var _zv *zend.Zval = current
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefStringl(ref, string, end-string)
								break
							}
							_zv = &ref.val
							zend.ZvalPtrDtor(_zv)
							var __z *zend.Zval = _zv
							var __s *zend.ZendString = zend.ZendStringInit(string, end-string, 0)
							__z.value.str = __s
							__z.u1.type_info = 6 | 1<<0<<8
							break
						}
						break
					}
				} else {
					zend.AddIndexStringl(return_value, g.PostInc(&objIndex), string, end-string)
				}
			}
			string = end
			break
		case 'i':

			/*
			 * Scan an unsigned or signed integer.
			 */

			buf[0] = '0'

			/*-cc-*/

			if width == 0 || width > g.SizeOf("buf")-1 {
				width = g.SizeOf("buf") - 1
			}
			flags |= 0x10 | 0x20 | 0x40
			for end = buf; width > 0; width-- {
				switch *string {
				case '0':

					/*-cc-*/

					if base == 16 {
						flags |= 0x80
					}

					/*-cc-*/

					if base == 0 {
						base = 8
						flags |= 0x80
					}
					if (flags & 0x40) != 0 {
						flags &= ^(0x10 | 0x20 | 0x40)
					} else {
						flags &= ^(0x10 | 0x80 | 0x20)
					}
					goto addToInt
				case '1':

				case '2':

				case '3':

				case '4':

				case '5':

				case '6':

				case '7':
					if base == 0 {
						base = 10
					}
					flags &= ^(0x10 | 0x80 | 0x20)
					goto addToInt
				case '8':

				case '9':
					if base == 0 {
						base = 10
					}
					if base <= 8 {
						break
					}
					flags &= ^(0x10 | 0x80 | 0x20)
					goto addToInt
				case 'A':

				case 'B':

				case 'C':

				case 'D':

				case 'E':

				case 'F':

				case 'a':

				case 'b':

				case 'c':

				case 'd':

				case 'e':

				case 'f':
					if base <= 10 {
						break
					}
					flags &= ^(0x10 | 0x80 | 0x20)
					goto addToInt
				case '+':

				case '-':
					if (flags & 0x10) != 0 {
						flags &= ^0x10
						goto addToInt
					}
					break
				case 'x':

				case 'X':
					if (flags&0x80) != 0 && end == buf+1 {
						base = 16
						flags &= ^0x80
						goto addToInt
					}
					break
				}

				/*
				 * We got an illegal character so we are done accumulating.
				 */

				break
			addToInt:

				/*
				 * Add the character to the temporary buffer.
				 */

				*string++
				g.PostInc(&(*end)) = (*string) - 1
				if (*string) == '0' {
					break
				}
			}

			/*
			 * Check to see if we need to back up because we only got a
			 * sign or a trailing x after a 0.
			 */

			if (flags & 0x20) != 0 {
				if (*string) == '0' {
					underflow = 1
				}
				goto done
			} else if end[-1] == 'x' || end[-1] == 'X' {
				end--
				string--
			}

			/*
			 * Scan the value from the temporary buffer.  If we are
			 * returning a large unsigned value, we have to convert it back
			 * to a string since PHP only supports signed values.
			 */

			if (flags & 0x2) == 0 {
				*end = '0'
				value = zend_long(*fn)(buf, nil, base)
				if (flags&0x4) != 0 && value < 0 {
					core.ApPhpSnprintf(buf, g.SizeOf("buf"), "%"+"llu", value)
					if numVars != 0 && objIndex >= argCount {
						break
					} else if numVars != 0 {

						/* change passed value type to string */

						current = args + g.PostInc(&objIndex)
						for {
							r.Assert(current.u1.v.type_ == 10)
							for {
								var _zv *zend.Zval = current
								var ref *zend.ZendReference = _zv.value.ref
								if ref.sources.ptr != nil {
									zend.ZendTryAssignTypedRefString(ref, buf)
									break
								}
								_zv = &ref.val
								zend.ZvalPtrDtor(_zv)
								var _s *byte = buf
								var __z *zend.Zval = _zv
								var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
								__z.value.str = __s
								__z.u1.type_info = 6 | 1<<0<<8
								break
							}
							break
						}
					} else {
						zend.AddIndexString(return_value, g.PostInc(&objIndex), buf)
					}
				} else {
					if numVars != 0 && objIndex >= argCount {
						break
					} else if numVars != 0 {
						current = args + g.PostInc(&objIndex)
						for {
							r.Assert(current.u1.v.type_ == 10)
							for {
								var _zv *zend.Zval = current
								var ref *zend.ZendReference = _zv.value.ref
								if ref.sources.ptr != nil {
									zend.ZendTryAssignTypedRefLong(ref, value)
									break
								}
								_zv = &ref.val
								zend.ZvalPtrDtor(_zv)
								var __z *zend.Zval = _zv
								__z.value.lval = value
								__z.u1.type_info = 4
								break
							}
							break
						}
					} else {
						zend.AddIndexLong(return_value, g.PostInc(&objIndex), value)
					}
				}
			}
			break
		case 'f':

			/*
			 * Scan a floating point number
			 */

			buf[0] = '0'
			if width == 0 || width > g.SizeOf("buf")-1 {
				width = g.SizeOf("buf") - 1
			}
			flags |= 0x10 | 0x20 | 0x100 | 0x200
			for end = buf; width > 0; width-- {
				switch *string {
				case '0':

				case '1':

				case '2':

				case '3':

				case '4':

				case '5':

				case '6':

				case '7':

				case '8':

				case '9':
					flags &= ^(0x10 | 0x20)
					goto addToFloat
				case '+':

				case '-':
					if (flags & 0x10) != 0 {
						flags &= ^0x10
						goto addToFloat
					}
					break
				case '.':
					if (flags & 0x100) != 0 {
						flags &= ^(0x10 | 0x100)
						goto addToFloat
					}
					break
				case 'e':

				case 'E':

					/*
					 * An exponent is not allowed until there has
					 * been at least one digit.
					 */

					if (flags & (0x20 | 0x200)) == 0x200 {
						flags = flags & ^(0x200|0x100) | 0x10 | 0x20
						goto addToFloat
					}
					break
				}

				/*
				 * We got an illegal character so we are done accumulating.
				 */

				break
			addToFloat:

				/*
				 * Add the character to the temporary buffer.
				 */

				*string++
				g.PostInc(&(*end)) = (*string) - 1
				if (*string) == '0' {
					break
				}
			}

			/*
			 * Check to see if we need to back up because we saw a
			 * trailing 'e' or sign.
			 */

			if (flags & 0x20) != 0 {
				if (flags & 0x200) != 0 {

					/*
					 * There were no digits at all so scanning has
					 * failed and we are done.
					 */

					if (*string) == '0' {
						underflow = 1
					}
					goto done
				}

				/*
				 * We got a bad exponent ('e' and maybe a sign).
				 */

				end--
				string--
				if (*end) != 'e' && (*end) != 'E' {
					end--
					string--
				}
			}

			/*
			 * Scan the value from the temporary buffer.
			 */

			if (flags & 0x2) == 0 {
				var dvalue float64
				*end = '0'
				dvalue = zend.ZendStrtod(buf, nil)
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + g.PostInc(&objIndex)
					for {
						r.Assert(current.u1.v.type_ == 10)
						for {
							var _zv *zend.Zval = current
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefDouble(ref, dvalue)
								break
							}
							_zv = &ref.val
							zend.ZvalPtrDtor(_zv)
							var __z *zend.Zval = _zv
							__z.value.dval = dvalue
							__z.u1.type_info = 5
							break
						}
						break
					}
				} else {
					zend.AddIndexDouble(return_value, g.PostInc(&objIndex), dvalue)
				}
			}
			break
		}
		nconversions++
	}
done:
	result = zend.SUCCESS
	if underflow != 0 && 0 == nconversions {
		ScanSetErrorReturn(numVars, return_value)
		result = -1
	} else if numVars != 0 {
		zend.ZvalPtrDtor(return_value)
		var __z *zend.Zval = return_value
		__z.value.lval = nconversions
		__z.u1.type_info = 4
	} else if nconversions < totalVars {

	}
	return result
}

/* }}} */

func ScanSetErrorReturn(numVars int, return_value *zend.Zval) {
	if numVars != 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = -1
		__z.u1.type_info = 4
	} else {

		/* convert_to_null calls destructor */

		zend.ConvertToNull(return_value)

		/* convert_to_null calls destructor */

	}
}

/* }}} */

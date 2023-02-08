// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func UCHAR(x byte) __auto__ { return zend_uchar(x) }
func BuildCharSet(cset *CharSet, format *byte) *byte {
	var ch *byte
	var start byte
	var nranges int
	var end *byte
	memset(cset, 0, b.SizeOf("CharSet"))
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
	cset.SetChars((*byte)(zend.SafeEmalloc(b.SizeOf("char"), end-format-1, 0)))
	if nranges > 0 {
		cset.SetRanges((*__struct__Range)(zend.SafeEmalloc(b.SizeOf("struct Range"), nranges, 0)))
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
		cset.GetChars()[b.PostInc(&(cset.GetNchars()))] = *ch
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
				cset.GetChars()[b.PostInc(&(cset.GetNchars()))] = start
				cset.GetChars()[b.PostInc(&(cset.GetNchars()))] = *ch
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
			cset.GetChars()[b.PostInc(&(cset.GetNchars()))] = *ch
		}
		format++
		ch = format - 1
	}
	return format
}
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
func ReleaseCharSet(cset *CharSet) {
	zend.Efree((*byte)(cset.GetChars()))
	if cset.GetRanges() != nil {
		zend.Efree((*byte)(cset.GetRanges()))
	}
}
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
		nassign = (*int)(zend.SafeEmalloc(b.SizeOf("int"), numVars, 0))
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
			flags |= SCAN_SUPPRESS
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

			value = zend.ZEND_STRTOUL(format-1, &end, 10)
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

				if value > SCAN_MAX_ARGS {
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
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s", "cannot mix \"%\" and \"%n$\" conversion specifiers")
			goto error
		}
	xpgCheckDone:

		/*
		 * Parse any width specifier.
		 */

		if isdigit(UCHAR(*ch)) {
			value = zend.ZEND_STRTOUL(format-1, &format, 10)
			flags |= SCAN_WIDTH
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
		if (flags&SCAN_SUPPRESS) == 0 && numVars != 0 && objIndex >= numVars {
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
			core.PhpErrorDocref(nil, zend.E_WARNING, "Unmatched [ in format string")
			goto error
		default:
			core.PhpErrorDocref(nil, zend.E_WARNING, "Bad scan conversion character \"%c\"", *ch)
			goto error
		}
		if (flags & SCAN_SUPPRESS) == 0 {
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
					nassign = any(zend.SafeEmalloc(nspace, b.SizeOf("int"), 0))
					for i = 0; i < 16; i++ {
						nassign[i] = staticAssign[i]
					}
				} else {
					nassign = any(zend.Erealloc(any(nassign), nspace*b.SizeOf("int")))
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
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s", "Variable is assigned by multiple \"%n$\" conversion specifiers")
			goto error
		} else if xpgSize == 0 && nassign[i] == 0 {

			/*
			 * If the space is empty, and xpgSize is 0 (means XPG wasn't
			 * used, and/or numVars != 0), then too many vars were given
			 */

			core.PhpErrorDocref(nil, zend.E_WARNING, "Variable is not assigned by any conversion specifiers")
			goto error
		}
	}
	if nassign != staticAssign {
		zend.Efree((*byte)(nassign))
	}
	return SCAN_SUCCESS
badIndex:
	if gotXpg != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", "\"%n$\" argument index out of range")
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Different numbers of variable names and field specifiers")
	}
error:
	if nassign != staticAssign {
		zend.Efree((*byte)(nassign))
	}
	return SCAN_ERROR_INVALID_FORMAT
}
func PhpSscanfInternal(
	string *byte,
	format *byte,
	argCount int,
	args *zend.Zval,
	varStart int,
	return_value *zend.Zval,
) int {
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
		varStart = SCAN_MAX_ARGS + 1
	}
	numVars = argCount - varStart
	if numVars < 0 {
		numVars = 0
	}

	/*
	 * Check for errors in the format string.
	 */

	if ValidateFormat(format, numVars, &totalVars) != SCAN_SUCCESS {
		ScanSetErrorReturn(numVars, return_value)
		return SCAN_ERROR_INVALID_FORMAT
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
			if !(args[i].IsReference()) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter %d must be passed by reference", i)
				ScanSetErrorReturn(numVars, return_value)
				return SCAN_ERROR_VAR_PASSED_BYVAL
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

		zend.ArrayInit(return_value)
		for i = 0; i < totalVars; i++ {
			tmp.SetNull()
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
			flags |= SCAN_SUPPRESS
			format++
			ch = format - 1
		} else if isdigit(UCHAR(*ch)) {
			value = zend.ZEND_STRTOUL(format-1, &end, 10)
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

		if isdigit(UCHAR(*ch)) {
			width = zend.ZEND_STRTOUL(format-1, &format, 10)
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
			if (flags & SCAN_SUPPRESS) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + b.PostInc(&objIndex)
					zend.ZEND_TRY_ASSIGN_REF_LONG(current, zend_long(string-baseString))
				} else {
					zend.AddIndexLong(return_value, b.PostInc(&objIndex), string-baseString)
				}
			}
			nconversions++
			continue
		case 'd':

		case 'D':
			op = 'i'
			base = 10
			fn = (func() zend.ZendLong)(zend.ZEND_STRTOL_PTR)
			break
		case 'i':
			op = 'i'
			base = 0
			fn = (func() zend.ZendLong)(zend.ZEND_STRTOL_PTR)
			break
		case 'o':
			op = 'i'
			base = 8
			fn = (func() zend.ZendLong)(zend.ZEND_STRTOL_PTR)
			break
		case 'x':

		case 'X':
			op = 'i'
			base = 16
			fn = (func() zend.ZendLong)(zend.ZEND_STRTOL_PTR)
			break
		case 'u':
			op = 'i'
			base = 10
			flags |= SCAN_UNSIGNED
			fn = (func() zend.ZendLong)(zend.ZEND_STRTOUL_PTR)
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
			flags |= SCAN_NOSKIP

			/*-cc-*/

			if 0 == width {
				width = 1
			}

			/*-cc-*/

			break
		case '[':
			op = '['
			flags |= SCAN_NOSKIP
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

		if (flags & SCAN_NOSKIP) == 0 {
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
				if b.PreDec(&width) == 0 {
					break
				}
			}
			if (flags & SCAN_SUPPRESS) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + b.PostInc(&objIndex)
					zend.ZEND_TRY_ASSIGN_REF_STRINGL(current, string, end-string)
				} else {
					zend.AddIndexStringl(return_value, b.PostInc(&objIndex), string, end-string)
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
				if b.PreDec(&width) == 0 {
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
			if (flags & SCAN_SUPPRESS) == 0 {
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + b.PostInc(&objIndex)
					zend.ZEND_TRY_ASSIGN_REF_STRINGL(current, string, end-string)
				} else {
					zend.AddIndexStringl(return_value, b.PostInc(&objIndex), string, end-string)
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

			if width == 0 || width > b.SizeOf("buf")-1 {
				width = b.SizeOf("buf") - 1
			}
			flags |= SCAN_SIGNOK | SCAN_NODIGITS | SCAN_NOZERO
			for end = buf; width > 0; width-- {
				switch *string {
				case '0':

					/*-cc-*/

					if base == 16 {
						flags |= SCAN_XOK
					}

					/*-cc-*/

					if base == 0 {
						base = 8
						flags |= SCAN_XOK
					}
					if (flags & SCAN_NOZERO) != 0 {
						flags &= ^(SCAN_SIGNOK | SCAN_NODIGITS | SCAN_NOZERO)
					} else {
						flags &= ^(SCAN_SIGNOK | SCAN_XOK | SCAN_NODIGITS)
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
					flags &= ^(SCAN_SIGNOK | SCAN_XOK | SCAN_NODIGITS)
					goto addToInt
				case '8':

				case '9':
					if base == 0 {
						base = 10
					}
					if base <= 8 {
						break
					}
					flags &= ^(SCAN_SIGNOK | SCAN_XOK | SCAN_NODIGITS)
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
					flags &= ^(SCAN_SIGNOK | SCAN_XOK | SCAN_NODIGITS)
					goto addToInt
				case '+':

				case '-':
					if (flags & SCAN_SIGNOK) != 0 {
						flags &= ^SCAN_SIGNOK
						goto addToInt
					}
					break
				case 'x':

				case 'X':
					if (flags&SCAN_XOK) != 0 && end == buf+1 {
						base = 16
						flags &= ^SCAN_XOK
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
				b.PostInc(&(*end)) = (*string) - 1
				if (*string) == '0' {
					break
				}
			}

			/*
			 * Check to see if we need to back up because we only got a
			 * sign or a trailing x after a 0.
			 */

			if (flags & SCAN_NODIGITS) != 0 {
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

			if (flags & SCAN_SUPPRESS) == 0 {
				*end = '0'
				value = zend_long(*fn)(buf, nil, base)
				if (flags&SCAN_UNSIGNED) != 0 && value < 0 {
					core.Snprintf(buf, b.SizeOf("buf"), zend.ZEND_ULONG_FMT, value)
					if numVars != 0 && objIndex >= argCount {
						break
					} else if numVars != 0 {

						/* change passed value type to string */

						current = args + b.PostInc(&objIndex)
						zend.ZEND_TRY_ASSIGN_REF_STRING(current, buf)
					} else {
						zend.AddIndexString(return_value, b.PostInc(&objIndex), buf)
					}
				} else {
					if numVars != 0 && objIndex >= argCount {
						break
					} else if numVars != 0 {
						current = args + b.PostInc(&objIndex)
						zend.ZEND_TRY_ASSIGN_REF_LONG(current, value)
					} else {
						zend.AddIndexLong(return_value, b.PostInc(&objIndex), value)
					}
				}
			}
			break
		case 'f':

			/*
			 * Scan a floating point number
			 */

			buf[0] = '0'
			if width == 0 || width > b.SizeOf("buf")-1 {
				width = b.SizeOf("buf") - 1
			}
			flags |= SCAN_SIGNOK | SCAN_NODIGITS | SCAN_PTOK | SCAN_EXPOK
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
					flags &= ^(SCAN_SIGNOK | SCAN_NODIGITS)
					goto addToFloat
				case '+':

				case '-':
					if (flags & SCAN_SIGNOK) != 0 {
						flags &= ^SCAN_SIGNOK
						goto addToFloat
					}
					break
				case '.':
					if (flags & SCAN_PTOK) != 0 {
						flags &= ^(SCAN_SIGNOK | SCAN_PTOK)
						goto addToFloat
					}
					break
				case 'e':

				case 'E':

					/*
					 * An exponent is not allowed until there has
					 * been at least one digit.
					 */

					if (flags & (SCAN_NODIGITS | SCAN_EXPOK)) == SCAN_EXPOK {
						flags = flags & ^(SCAN_EXPOK|SCAN_PTOK) | SCAN_SIGNOK | SCAN_NODIGITS
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
				b.PostInc(&(*end)) = (*string) - 1
				if (*string) == '0' {
					break
				}
			}

			/*
			 * Check to see if we need to back up because we saw a
			 * trailing 'e' or sign.
			 */

			if (flags & SCAN_NODIGITS) != 0 {
				if (flags & SCAN_EXPOK) != 0 {

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

			if (flags & SCAN_SUPPRESS) == 0 {
				var dvalue float64
				*end = '0'
				dvalue = zend.ZendStrtod(buf, nil)
				if numVars != 0 && objIndex >= argCount {
					break
				} else if numVars != 0 {
					current = args + b.PostInc(&objIndex)
					zend.ZEND_TRY_ASSIGN_REF_DOUBLE(current, dvalue)
				} else {
					zend.AddIndexDouble(return_value, b.PostInc(&objIndex), dvalue)
				}
			}
			break
		}
		nconversions++
	}
done:
	result = SCAN_SUCCESS
	if underflow != 0 && 0 == nconversions {
		ScanSetErrorReturn(numVars, return_value)
		result = SCAN_ERROR_EOF
	} else if numVars != 0 {
		zend.ZvalPtrDtor(return_value)
		return_value.SetLong(nconversions)
	} else if nconversions < totalVars {

	}
	return result
}
func ScanSetErrorReturn(numVars int, return_value *zend.Zval) {
	if numVars != 0 {
		return_value.SetLong(SCAN_ERROR_EOF)
	} else {

		/* convert_to_null calls destructor */

		zend.ConvertToNull(return_value)

		/* convert_to_null calls destructor */

	}
}

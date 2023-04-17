package pfmt

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
	"math"
	"strconv"
	"strings"
)

const FORMAT_CONV_MAX_PRECISION = 500
const (
	LM_STD = iota
	LM_INTMAX_T
	LM_PTRDIFF_T
	LM_LONG_LONG
	LM_SIZE_T
	LM_LONG
	LM_LONG_DOUBLE
	LM_PHP_INT_T
)

type printer struct {
	buf []byte
	err error
}

func (p *printer) writeByte(c byte) {
	p.buf = append(p.buf, c)
}
func (p *printer) writeString(s string) {
	p.buf = append(p.buf, []byte(s)...)
}
func (p *printer) len() int {
	return len(p.buf)
}
func (p *printer) doPrint(fmt string, args ...any) {
	argIdx := 0
	for r := newByteReader(fmt); r.valid(); {
		c := r.read()
		if c != '%' {
			p.writeByte(c)
			continue
		}

		/*
		 * Default variable settings
		 */
		adjustRight := true
		printBlank := false
		printSign := false
		alternateForm := false
		padChar := byte(' ')
		prefixChar := byte(0)
		adjustWidth := false
		adjustPrecision := false
		minWidth := 0
		precision := 0

		/*
		 * Try to avoid checking for flags, width or precision
		 */
		if r.isAscii() && !r.isLower() {
			/*
			 * Recognize flags: -, #, BLANK, +
			 */
		flagsLoop:
			for ; ; r.read() {
				switch r.curr() {
				case '-':
					adjustRight = false
				case '+':
					printSign = true
				case '#':
					alternateForm = true
				case ' ':
					printBlank = true
				case '0':
					padChar = '0'
				default:
					break flagsLoop
				}
			}

			/*
			 * Check if a width was specified
			 */
			if r.isDigit() {
				minWidth = r.readDec()
				adjustWidth = true
			} else if r.curr() == '*' {
				minWidth = args[argIdx].(int)
				argIdx++
				r.read()
				adjustWidth = true
				if minWidth < 0 {
					adjustRight = false
					minWidth = -minWidth
				}
			}

			/*
			 * Check if a precision was specified
			 */
			if r.curr() == '.' {
				adjustPrecision = true
				r.read()
				if r.isDigit() {
					precision = r.readDec()
				} else if r.curr() == '*' {
					precision = args[argIdx].(int)
					argIdx++
					if precision < -1 {
						precision = -1
					}
				} else {
					precision = 0
				}
				if precision > FORMAT_CONV_MAX_PRECISION {
					precision = FORMAT_CONV_MAX_PRECISION
				}
			}
		}

		/*
		 * Modifier check
		 */
		var modifier int
		_ = modifier
		switch r.curr() {
		case 'L':
			r.read()
			modifier = LM_LONG_DOUBLE
		case 'I':
			r.read()
			if r.tryRead("64") {
				modifier = LM_LONG_LONG
			} else if r.tryRead("32") {
				modifier = LM_LONG
			} else {
				modifier = LM_LONG
			}
		case 'l':
			r.read()
			if r.tryRead("l") {
				modifier = LM_LONG_LONG
			} else {
				modifier = LM_LONG
			}
		case 'z':
			r.read()
			modifier = LM_SIZE_T
		case 'j':
			r.read()
			modifier = LM_INTMAX_T
		case 't':
			r.read()
			modifier = LM_PTRDIFF_T
		case 'p':
			next := r.offset(1)
			if next == 'd' || next == 'u' || next == 'x' || next == 'o' {
				r.read()
				modifier = LM_PHP_INT_T
			} else {
				modifier = LM_STD
			}
		case 'h':
			r.read()
			if r.curr() == 'h' {
				r.read()
			}
			modifier = LM_STD
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
		var s string
		switch r.curr() {
		case 'Z':
			zvp := args[argIdx].(*types.Zval)
			s = zend.PrintZval(zvp)
			if adjustPrecision && precision < len(s) {
				s = s[:precision]
			}
		case 'u':
			num := args[argIdx].(uint)
			argIdx++
			s = strconv.FormatUint(uint64(num), 0)
			s = fixPrecision(adjustPrecision, precision, s)
		case 'd', 'i':
			num := args[argIdx].(int)
			argIdx++
			if num >= 0 {
				s = strconv.Itoa(num)
			} else {
				s = strconv.Itoa(-num)
			}
			s = fixPrecision(adjustPrecision, precision, s)
			if num < 0 {
				prefixChar = '-'
			} else if printSign {
				prefixChar = '+'
			} else if printBlank {
				prefixChar = ' '
			}
		case 'o':
			num := args[argIdx].(uint)
			argIdx++
			s = strconv.FormatUint(uint64(num), 8)
			s = fixPrecision(adjustPrecision, precision, s)
			if alternateForm && s[0] != '0' {
				s = "0" + s
			}
		case 'x', 'X':
			num := args[argIdx].(uint)
			argIdx++
			s = strconv.FormatUint(uint64(num), 16)
			s = fixPrecision(adjustPrecision, precision, s)
			if alternateForm && s[0] != '0' {
				s = "0" + string(r.curr()) + s
			}
		case 's', 'v':
			sp := args[argIdx].(*byte)
			argIdx++
			if sp != nil {
				s = b.CastStrAuto(sp)
				if adjustPrecision && len(s) > precision {
					s = s[:precision]
				}
			} else {
				s = ""
			}
			padChar = ' '
		case 'f', 'F', 'e', 'E':
			num := args[argIdx].(float64)
			argIdx++
			if math.IsNaN(num) {
				s = "nan"
			} else if math.IsInf(num, 1) || math.IsInf(num, -1) {
				s = "inf"
			} else {
				sign := r.curr()

				if num < 0 {
					s = strconv.FormatFloat(-num, sign, 0, 64)
				} else {
					s = strconv.FormatFloat(num, sign, 0, 64)
				}
				if num < 0 {
					prefixChar = '-'
				} else if printSign {
					prefixChar = '+'
				} else {
					prefixChar = ' '
				}
			}
		case 'g', 'k', 'G', 'H':
			num := args[argIdx].(float64)
			argIdx++
			if math.IsNaN(num) {
				s = "NAN"
			} else if math.IsInf(num, 1) {
				s = "INF"
			} else if math.IsInf(num, -1) {
				s = "-INF"
			} else {
				// todo 此处有部分细节未处理
				sign := r.curr()

				if num < 0 {
					s = strconv.FormatFloat(-num, sign, 0, 64)
				} else {
					s = strconv.FormatFloat(num, sign, 0, 64)
				}
				if num < 0 {
					prefixChar = '-'
				} else if printSign {
					prefixChar = '+'
				} else {
					prefixChar = ' '
				}
			}
		case 'c':
			byte_ := byte(args[argIdx].(int))
			argIdx++
			s = string(byte_)
			padChar = ' '
		case '%':
			s = "%"
			padChar = ' '
		case 'n':
			args[argIdx] = p.len()
			argIdx++
			continue
		case 'p':
			// todo
		case 0:
			continue
		default:
			s = "%" + string(r.curr())
			padChar = ' '
		}
		if prefixChar != 0 {
			s = string(prefixChar) + s
		}
		if adjustWidth && adjustRight && minWidth > len(s) {
			if padChar == '0' && prefixChar != 0 {
				p.writeByte(s[0])
				s = s[1:]
				minWidth--
			}
			for i := 0; i < minWidth-len(s); i++ {
				p.writeByte(padChar)
			}
		}

		/*
		 * Print the string s.
		 */
		p.writeString(s)
		if adjustWidth && !adjustRight && minWidth > len(s) {
			for i := 0; i < minWidth-len(s); i++ {
				p.writeByte(padChar)
			}
		}
	}
}

func fixPrecision(adjust bool, precision int, s string) string {
	if adjust && len(s) < precision {
		return strings.Repeat("0", precision-len(s)) + s
	}
	return s
}

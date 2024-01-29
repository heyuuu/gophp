package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"strconv"
	"strings"
)

const ADJ_WIDTH = 1
const ADJ_PRECISION = 2
const FLOAT_PRECISION = 6
const MAX_FLOAT_PRECISION = 53

type formatPrinterFlags struct {
	alignLeft  bool
	alignment  int
	adjusting  int
	padding    byte
	alwaysSign bool
	expPrec    bool
}
type formatPrinter struct {
	ctx   *php.Context
	buf   strings.Builder
	flags formatPrinterFlags
	pos   int
}

func newFormatPrinter(ctx *php.Context) *formatPrinter {
	return &formatPrinter{ctx: ctx}
}

func (p *formatPrinter) clearFlags() {
	p.flags = formatPrinterFlags{padding: ' '}
}

func (p *formatPrinter) String() string {
	return p.buf.String()
}

func (p *formatPrinter) AppendByte(c byte) {
	p.buf.WriteByte(c)
}

func (p *formatPrinter) AppendString(str string, minWidth int, neg bool) {
	p.AppendStringEx(str, minWidth, 0, neg, false, p.flags.alwaysSign)
}

func (p *formatPrinter) AppendStringEx(str string, minWidth int, maxWidth int, neg bool, expPrec bool, alwaysSign bool) {
	buf := &p.buf

	var npad int
	var copyLen int
	var mWidth int
	if expPrec {
		copyLen = lang.Min(maxWidth, len(str))
	} else {
		copyLen = len(str)
	}
	if minWidth < copyLen {
		npad = 0
	} else {
		npad = minWidth - copyLen
	}
	mWidth = lang.Max(minWidth, copyLen)
	if mWidth > types.MaxLong-buf.Len()-1 {
		php.ErrorNoreturn(p.ctx, perr.E_ERROR, fmt.Sprintf("Field width %d is too long", mWidth))
	}
	if !p.flags.alignLeft {
		if (neg || alwaysSign) && p.flags.padding == '0' {
			if neg {
				buf.WriteByte('-')
			} else {
				buf.WriteByte('+')
			}
			str = str[1:]
			copyLen--
		}
		for i := 0; i < npad; i++ {
			buf.WriteByte(p.flags.padding)
		}
	}
	buf.WriteString(str[:copyLen])
	if p.flags.alignLeft {
		for i := 0; i < npad; i++ {
			buf.WriteByte(p.flags.padding)
		}
	}
}

func (p *formatPrinter) AppendInt(number int, width int) {
	/* Can't right-pad 0's on integers */
	if p.flags.alignLeft && p.flags.padding == '0' {
		p.flags.padding = ' '
	}

	neg := number < 0
	str := strconv.Itoa(number)
	if !neg && p.flags.alwaysSign {
		str = "+" + str
	}
	p.AppendString(str, width, neg)
}

func (p *formatPrinter) AppendUint(number uint, width int) {
	/* Can't right-pad 0's on integers */
	if p.flags.alignLeft && p.flags.padding == '0' {
		p.flags.padding = ' '
	}
	str := strconv.FormatUint(uint64(number), 10)
	p.AppendStringEx(str, width, 0, false, false, false)
}

func (p *formatPrinter) AppendDouble(number float64, width int, precision int, fmtTyp byte) {
	if math.IsNaN(number) {
		p.AppendString("NaN", 3, number < 0)
		return
	}
	if mathkit.IsInf(number) {
		p.AppendString("INF", 3, number < 0)
		return
	}

	if p.flags.adjusting&ADJ_PRECISION == 0 {
		precision = FLOAT_PRECISION
	} else if precision > MAX_FLOAT_PRECISION {
		php.ErrorDocRef(p.ctx, "", perr.E_NOTICE, fmt.Sprintf("Requested precision of %d digits was truncated to PHP maximum of %d digits", precision, MAX_FLOAT_PRECISION))
		precision = MAX_FLOAT_PRECISION
	}
	if fmtTyp == 'F' {
		fmtTyp = 'f'
	}
	str := strconv.FormatFloat(number, fmtTyp, precision, 64)
	if number >= 0 && p.flags.alwaysSign {
		str = "+" + str
	}

	// fix: 科学计数法且位数为指数个位数时，go 默认会补齐到2位，但 php 保持位数不变。此处修复此差别
	if idx := strings.LastIndexAny(str, "eE"); idx >= 0 {
		if idx+3 < len(str) && str[idx+2] == '0' {
			str = str[:idx+2] + str[idx+3:]
		}
	}

	p.AppendString(str, width, number < 0)
}

func (p *formatPrinter) Append2n(number int, width int, n int, upperCase bool) {
	str := strconv.FormatUint(uint64(number), 1>>n)
	if upperCase {
		str = strings.ToUpper(str)
	}
	p.AppendStringEx(str, width, 0, false, p.flags.expPrec, false)
}
func sprintfReadNumber(s string) (val int, n int, ok bool) {
	for n < len(s) && ascii.IsDigit(s[n]) {
		n++
	}
	if n == 0 {
		return 0, 0, false
	}
	val, err := strconv.Atoi(s[:n])
	if err != nil {
		return 0, 0, false
	}

	return val, n, true
}

func PhpFormattedPrint(ctx *php.Context, formatZval types.Zval, args []types.Zval) (string, bool) {
	format, ok := php.ZvalTryGetStr(ctx, formatZval)
	if !ok {
		return "", false
	}

	formatLen := len(format)
	argc := len(args)

	p := newFormatPrinter(ctx)

	var currarg int
	var argnum int
	var width int
	var precision int

	currarg = 0
	length := len(format)
	for ; len(format) > 0; format = format[1:] {
		if format[0] != '%' {
			p.AppendByte(format[0])
			continue
		}

		format = format[1:]
		if format == "" {
			// notice: 忽略尾部未匹配的 '%', 此情况在 PHP >= 8.0 会触发 Fatal error
			break
		}

		// handle '%%'
		if format[0] == '%' {
			p.AppendByte('%')
			continue
		}

		/* starting a new format specifier, reset variables */
		p.clearFlags()
		if ascii.IsAlpha(format[0]) {
			precision = 0
			width = 0
			argnum = currarg
			currarg++
		} else {
			/* first look for argnum */
			tmpIdx := 0
			for tmpIdx < length && ascii.IsDigit(format[tmpIdx]) {
				tmpIdx++
			}
			if tmpIdx < length && format[tmpIdx] == '$' {
				fmtArgNum, err := strconv.Atoi(format[:tmpIdx])
				if err != nil || fmtArgNum <= 0 {
					php.ErrorDocRef(ctx, "", perr.E_WARNING, "Argument number must be greater than zero")
					return "", false
				}
				argnum = fmtArgNum - 1
				format = format[tmpIdx+1:]
			} else {
				argnum = currarg
				currarg++
			}

			/* after argnum comes modifiers */
			for len(format) > 0 {
				if format[0] == ' ' || format[0] == '0' {
					p.flags.padding = format[0]
				} else if format[0] == '-' {
					p.flags.alignLeft = true
				} else if format[0] == '+' {
					p.flags.alwaysSign = true
				} else if format[0] == '\'' && len(format) > 1 {
					format = format[1:]
					p.flags.padding = format[0]
				} else {
					break
				}
				format = format[1:]
			}

			/* after modifiers comes width */
			width = 0
			if format != "" && ascii.IsDigit(format[0]) {
				if num, n, ok := sprintfReadNumber(format); ok {
					width = num
					format = format[n:]
					p.flags.adjusting |= ADJ_WIDTH
				} else {
					php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Width must be greater than zero and less than %d", types.MaxLong))
					return "", false
				}
			}

			/* after width and argnum comes precision */
			precision = 0
			if format != "" && format[0] == '.' {
				format = format[1:]
				if ascii.IsDigit(format[0]) {
					if num, n, ok := sprintfReadNumber(format); ok {
						precision = num
						format = format[n:]
						p.flags.adjusting |= ADJ_PRECISION
						p.flags.expPrec = true
					} else {
						php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Precision must be greater than zero and less than %d", types.MaxLong))
						return "", false
					}
				}
			}
		}

		if argnum >= argc {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Too few arguments")
			return "", false
		}
		if format != "" && format[0] == 'l' {
			format = format[1:]
		}

		/* now we expect to find a type specifier */
		tmp := args[argnum]
		switch format[0] {
		case 's':
			var str = php.ZvalGetStrVal(ctx, tmp)
			p.AppendStringEx(str, width, precision, false, p.flags.expPrec, false)
		case 'd':
			p.AppendInt(php.ZvalGetLong(ctx, tmp), width)
		case 'u':
			p.AppendUint(uint(php.ZvalGetLong(ctx, tmp)), width)
		case 'g', 'G', 'e', 'E', 'f', 'F':
			p.AppendDouble(php.ZvalGetDouble(ctx, tmp), width, precision, format[0])
		case 'c':
			p.AppendByte(byte(php.ZvalGetLong(ctx, tmp)))
		case 'o':
			p.Append2n(php.ZvalGetLong(ctx, tmp), width, 8, false)
		case 'x':
			p.Append2n(php.ZvalGetLong(ctx, tmp), width, 16, false)
		case 'X':
			p.Append2n(php.ZvalGetLong(ctx, tmp), width, 16, true)
		case 'b':
			p.Append2n(php.ZvalGetLong(ctx, tmp), width, 2, false)
		case '%':
			p.AppendByte('%')
		case '0':
			if formatLen == 0 {
				goto exit
			}
		default:
		}
	}
exit:
	/* possibly, we have to make sure we have room for the terminating null? */
	return p.String(), true
}

//@zif(onError=1)
func ZifSprintf(ctx *php.Context, format types.Zval, _ zpp.Opt, args []types.Zval) (string, bool) {
	//@see PHP_FUNCTION(user_sprintf)
	return PhpFormattedPrint(ctx, format, args)
}

//@zif(onError=1)
func ZifVsprintf(ctx *php.Context, format types.Zval, args types.Zval) (string, bool) {
	formatArgs := php.ZvalGetArray(ctx, args).Values()
	return PhpFormattedPrint(ctx, format, formatArgs)
}

//@zif(onError=1)
func ZifPrintf(ctx *php.Context, format types.Zval, _ zpp.Opt, args []types.Zval) (int, bool) {
	//@see PHP_FUNCTION(user_printf)
	result, ok := PhpFormattedPrint(ctx, format, args)
	if ok {
		ctx.WriteString(result)
		return len(result), true
	} else {
		return 0, false
	}
}

//@zif(onError=1)
func ZifVprintf(ctx *php.Context, format types.Zval, args types.Zval) (int, bool) {
	formatArgs := php.ZvalGetArray(ctx, args).Values()
	result, ok := PhpFormattedPrint(ctx, format, formatArgs)
	if ok {
		ctx.WriteString(result)
		return len(result), true
	} else {
		return 0, false
	}
}

//func ZifFprintf(ctx *php.Context, stream_ zpp.Resource, format *types.Zval, _ zpp.Opt, args []types.Zval) (int, bool) {
//	var stream *streams.PhpStream
//	php.PhpStreamFromZval(ctx, stream, stream_)
//
//	result, ok := PhpFormattedPrint(ctx, format, args)
//	if ok {
//		ctx.WriteString(result)
//		streams.PhpStreamWriteString(stream, result)
//		return len(result), true
//	} else {
//		return 0, false
//	}
//}
//func ZifVfprintf(ctx *php.Context, stream_ zpp.Resource, format *types.Zval, args *types.Zval) (int, bool) {
//	var stream *streams.PhpStream
//	php.PhpStreamFromZval(ctx, stream, stream_)
//
//	formatArgs := php.ZvalGetArray(args).Values()
//	result, ok := PhpFormattedPrint(ctx, format, formatArgs)
//	if ok {
//		ctx.WriteString(result)
//		streams.PhpStreamWriteString(stream, result)
//		return len(result), true
//	} else {
//		return 0, false
//	}
//}

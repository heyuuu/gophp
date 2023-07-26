package standard

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
	"strconv"
	"strings"
)

func zvalGetFloat(zv *types.Zval) (float64, bool) {
	if zv.IsLong() {
		return float64(zv.Long()), true
	} else if zv.IsDouble() {
		return zv.Double(), true
	} else {
		return 0, false
	}
}

func phpRound(f float64, dec int, mode int) float64 {
	v := f * math.Pow(10, float64(dec))
	if math.IsInf(v, 1) || math.IsInf(v, -1) {
		return f
	}
	switch mode {
	case PHP_ROUND_HALF_UP:
		v = math.Round(v)
	case PHP_ROUND_HALF_DOWN:
		v = math.Trunc(v)
	case PHP_ROUND_HALF_EVEN:
		v = math.RoundToEven(v)
	default: // PHP_ROUND_HALF_ODD
		v = math.RoundToEven(v+1) - 1
	}
	v /= math.Pow(10, float64(dec))
	return v
}

func ZifAbs(number *types.Zval) (*types.Zval, bool) {
	switch number.GetType() {
	case types.IsDouble:
		num := number.Double()
		result := math.Abs(num)
		return types.NewZvalDouble(result), true
	case types.IsLong:
		num := number.Long()
		if num == math.MinInt { // overflow
			result := -float64(num)
			return types.NewZvalDouble(result), true
		} else {
			result := -num
			return types.NewZvalLong(result), true
		}
	default:
		return nil, false
	}
}
func ZifCeil(number *types.Zval) (float64, bool) {
	operators.ConvertScalarToNumberEx(number)
	switch number.GetType() {
	case types.IsDouble:
		result := math.Ceil(number.Double())
		return result, true
	case types.IsLong:
		result := float64(number.Long())
		return result, true
	default:
		return 0, false
	}
}
func ZifFloor(number *types.Zval) (float64, bool) {
	operators.ConvertScalarToNumberEx(number)
	switch number.GetType() {
	case types.IsDouble:
		result := math.Floor(number.Double())
		return result, true
	case types.IsLong:
		result := float64(number.Long())
		return result, true
	default:
		return 0, false
	}
}
func ZifRound(number *types.Zval, _ zpp.Opt, precision int, mode_ *int) (float64, bool) {
	var value *types.Zval = number
	var mode = b.Option(mode_, PHP_ROUND_HALF_UP)

	precision = b.FixRange(precision, core.INT_MIN+1, core.INT_MAX)
	operators.ConvertScalarToNumberEx(value)
	if value.IsLong() && precision >= 0 {
		/* Simple case - long that doesn't need to be rounded. */
		return float64(value.Long()), true
	} else if value.IsLong() || value.IsDouble() {
		val, _ := zvalGetFloat(value)
		return phpRound(val, precision, mode), true
	}
	return 0, false
}
func ZifSin(number float64) float64 {
	return math.Sin(number)
}
func ZifCos(number float64) float64 {
	return math.Cos(number)
}
func ZifTan(number float64) float64 {
	return math.Tan(number)
}
func ZifAsin(number float64) float64 {
	return math.Asin(number)
}
func ZifAcos(number float64) float64 {
	return math.Acos(number)
}
func ZifAtan(number float64) float64 {
	return math.Atan(number)
}
func ZifAtan2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}
func ZifSinh(number float64) float64 {
	return math.Sinh(number)
}
func ZifCosh(number float64) float64 {
	return math.Cosh(number)
}
func ZifTanh(number float64) float64 {
	return math.Tanh(number)
}
func ZifAsinh(number float64) float64 {
	return math.Asinh(number)
}
func ZifAcosh(number float64) float64 {
	return math.Acosh(number)
}
func ZifAtanh(number float64) float64 {
	return math.Atanh(number)
}
func ZifPi() float64 {
	return M_PI
}
func ZifIsFinite(val float64) bool {
	return core.ZendFinite(val)
}
func ZifIsInfinite(val float64) bool {
	return math.IsInf(val, 1) || math.IsInf(val, -1)
}
func ZifIsNan(val float64) bool {
	return math.IsNaN(val)
}
func ZifPow(returnValue zpp.Ret, base *types.Zval, exponent *types.Zval) {
	operators.PowFunction(returnValue, base, exponent)
}
func ZifExp(number float64) float64 {
	return math.Exp(number)
}
func ZifExpm1(number float64) float64 {
	return math.Expm1(number)
}
func ZifLog1p(number float64) float64 {
	return math.Log1p(number)
}
func ZifLog(number float64, _ zpp.Opt, base *float64) (float64, bool) {
	if base == nil {
		return math.Log(number), true
	}
	if *base == 10.0 {
		return math.Log10(number), true
	}
	if *base == 1.0 {
		return math.NaN(), true
	}
	if *base <= 0.0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "base must be greater than 0")
		return 0, false
	}
	return math.Log(number) / math.Log(*base), true
}
func ZifLog10(number float64) float64 {
	return math.Log10(number)
}
func ZifSqrt(number float64) float64 {
	return math.Sqrt(number)
}
func ZifHypot(num1 float64, num2 float64) float64 {
	return math.Hypot(num1, num2)
}
func ZifDeg2rad(number float64) float64 {
	return number / 180.0 * M_PI
}
func ZifRad2deg(number float64) float64 {
	return number / M_PI * 180
}

func _parseStringToNumberZval(s string, base int) (*types.Zval, bool) {
	num, ok := _parseStringToNumber(s, base)
	if !ok {
		return nil, false
	} else if num.IsInt() {
		return types.NewZvalLong(num.Int()), true
	} else {
		return types.NewZvalDouble(num.Float()), true
	}
}
func _parseStringToNumber(s string, base int) (types.Number, bool) {
	s = strings.TrimFunc(s, ascii.IsSpaceRune)
	if len(s) >= 2 {
		if base == 16 && (s[:2] == "0x" || s[:2] == "0X") {
			s = s[2:]
		}
		if base == 8 && (s[:2] == "0o" || s[:2] == "0O") {
			s = s[2:]
		}
		if base == 2 && (s[:2] == "0b" || s[:2] == "0B") {
			s = s[2:]
		}
	}

	cutoff := zend.ZEND_LONG_MAX / base
	cutlim := zend.ZEND_LONG_MAX % base
	inum := 0
	fnum := float64(0)
	isFloat := false

	invalidchars := 0
	for _, c := range s {
		var digit int
		if c >= '0' && c <= '9' {
			digit = int(c - '0')
		} else if c >= 'A' && c <= 'Z' {
			digit = int(c - 'A' + 10)
		} else if c >= 'a' && c <= 'z' {
			digit = int(c - 'a' + 10)
		} else {
			invalidchars++
			continue
		}
		if digit >= base {
			invalidchars++
			continue
		}
		if !isFloat && (inum < cutoff || inum == cutoff && digit <= cutlim) {
			inum = inum*base + digit
		} else {
			isFloat = true
			fnum = fnum*float64(base) + float64(digit)
		}
	}
	if invalidchars > 0 {
		faults.Error(faults.E_DEPRECATED, "Invalid characters passed for attempted conversion, these have been ignored")
	}
	if isFloat {
		return types.FloatNumber(fnum), true
	} else {
		return types.IntNumber(inum), true
	}
}

func ZifBindec(binaryNumber *types.Zval) (*types.Zval, bool) {
	operators.ConvertToStringEx(binaryNumber)
	return _parseStringToNumberZval(binaryNumber.StringVal(), 2)
}
func ZifHexdec(hexadecimalNumber *types.Zval) (*types.Zval, bool) {
	operators.ConvertToStringEx(hexadecimalNumber)
	return _parseStringToNumberZval(hexadecimalNumber.StringVal(), 16)
}
func ZifOctdec(octalNumber *types.Zval) (*types.Zval, bool) {
	operators.ConvertToStringEx(octalNumber)
	return _parseStringToNumberZval(octalNumber.StringVal(), 8)
}
func ZifDecbin(decimalNumber int) string {
	return strconv.FormatInt(int64(decimalNumber), 2)
}
func ZifDecoct(decimalNumber int) string {
	return strconv.FormatInt(int64(decimalNumber), 8)
}
func ZifDechex(decimalNumber int) string {
	return strconv.FormatInt(int64(decimalNumber), 16)
}
func ZifBaseConvert(number *types.Zval, frombase int, tobase int) (string, bool) {
	if operators.TryConvertToString(number) == 0 { // fail
		return "", false
	}
	if frombase < 2 || frombase > 36 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid `from base' ("+zend.ZEND_LONG_FMT+")", frombase)
		return "", false
	}
	if tobase < 2 || tobase > 36 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid `to base' ("+zend.ZEND_LONG_FMT+")", tobase)
		return "", false
	}

	num, ok := _parseStringToNumber(number.StringVal(), frombase)
	if !ok {
		return "", false
	}

	/* Don't try to convert +/- infinity */
	if num.IsInf() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Number too large")
		return "", true
	}

	val := num.Floor()
	return strconv.FormatInt(int64(val), tobase), true
}
func phpMathNumberFormat(d float64, dec int, decPoint string, thousandSep string) string {
	// special cases
	if math.IsNaN(d) {
		return "nan"
	} else if math.IsInf(d, 1) {
		return "inf"
	} else if math.IsInf(d, -1) {
		return "-inf"
	}

	//
	isNegative := false
	if d < 0 {
		isNegative = true
		d = -d
	}
	if dec < 0 {
		dec = 0
	}
	d = phpRound(d, dec, PHP_ROUND_HALF_UP)

	/* Check if the number is no longer negative after rounding */
	if isNegative && d == 0 {
		isNegative = false
	}

	//
	tmp := fmt.Sprintf("%f", dec)

	var buf strings.Builder
	if isNegative {
		buf.WriteByte('-')
	}
	var pointPos int
	if pos := strings.Index(tmp, "."); pos >= 0 {
		pointPos = pos
	} else {
		pointPos = len(tmp)
	}

	first := true
	for i := pointPos % 3; i < pointPos; i += 3 {
		if first {
			first = false
		} else {
			buf.WriteString(thousandSep)
		}
		buf.WriteString(tmp[i : i+3])
	}

	if pointPos < len(tmp) {
		buf.WriteString(decPoint)
		if pointPos+1+dec < len(tmp) {
			buf.WriteString(tmp[pointPos+1:])
		} else {
			buf.WriteString(tmp[pointPos+1 : pointPos+1+dec])
		}
	}

	return buf.String()
}
func ZifNumberFormat(number float64, _ zpp.Opt, numDecimalPlaces int, decSeparator_ *string, thousandsSeparator *string) string {
	var decPoint = b.Option(decSeparator_, ".")
	var thousandSep = b.Option(thousandsSeparator, ",")
	return phpMathNumberFormat(number, numDecimalPlaces, decPoint, thousandSep)
}
func ZifFmod(x float64, y float64) float64 {
	return math.Mod(x, y)
}
func ZifIntdiv(dividend int, divisor int) int {
	if divisor == 0 {
		faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Division by zero")
		return 0
	} else if divisor == -1 && dividend == zend.ZEND_LONG_MIN {
		/* Prevent overflow error/crash ... really should not happen:
		   We don't return a float here as that violates function contract */
		faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Division of PHP_INT_MIN by -1 is not an integer")
		return 0
	}
	return dividend / divisor
}

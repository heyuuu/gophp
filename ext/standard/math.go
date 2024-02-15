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

const M_E = 2.7182817
const M_LOG2E = 1.442695
const M_LOG10E = 0.4342945
const M_LN2 = 0.6931472
const M_LN10 = 2.3025851
const M_PI = 3.1415927
const M_PI_2 = 1.5707964
const M_PI_4 = 0.7853982
const M_1_PI = 0.31830987
const M_2_PI = 0.63661975
const M_SQRTPI = 1.7724539
const M_2_SQRTPI = 1.1283792
const M_LNPI = 1.1447299
const M_EULER = 0.5772157
const M_SQRT2 = 1.4142135
const M_SQRT1_2 = 0.70710677
const M_SQRT3 = 1.7320508

/* Define rounding modes (all are round-to-nearest) */

const PHP_ROUND_HALF_UP = 0x1
const PHP_ROUND_HALF_DOWN = 0x2
const PHP_ROUND_HALF_EVEN = 0x3
const PHP_ROUND_HALF_ODD = 0x4

// --- functions

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

func ZifAbs(ctx *php.Context, number types.Zval) types.Zval {
	number = php.ConvertScalarToNumber(ctx, number)

	switch number.Type() {
	case types.IsDouble:
		num := number.Double()
		result := math.Abs(num)
		return types.ZvalDouble(result)
	case types.IsLong:
		num := number.Long()
		if num == math.MinInt { // overflow
			result := -float64(num)
			return types.ZvalDouble(result)
		} else {
			result := num
			if num < 0 {
				result = -num
			}
			return types.ZvalLong(result)
		}
	default:
		return types.False
	}
}
func ZifCeil(ctx *php.Context, number types.Zval) (float64, bool) {
	number = php.ConvertScalarToNumber(ctx, number)
	switch number.Type() {
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
func ZifFloor(ctx *php.Context, number types.Zval) (float64, bool) {
	number = php.ConvertScalarToNumber(ctx, number)
	switch number.Type() {
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
func ZifRound(ctx *php.Context, number types.Zval, _ zpp.Opt, precision int, mode_ *int) (float64, bool) {
	mode := lang.Option(mode_, PHP_ROUND_HALF_UP)
	precision = lang.FixRange(precision, types.MinLong+1, types.MaxLong)
	value := php.ConvertScalarToNumber(ctx, number)
	if value.IsLong() {
		if precision >= 0 {
			/* Simple case - long that doesn't need to be rounded. */
			return float64(value.Long()), true
		}
		return phpRound(float64(value.Long()), precision, mode), true
	} else if value.IsDouble() {
		return phpRound(value.Double(), precision, mode), true
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
	return mathkit.IsFinite(val)
}
func ZifIsInfinite(val float64) bool {
	return math.IsInf(val, 1) || math.IsInf(val, -1)
}
func ZifIsNan(val float64) bool {
	return math.IsNaN(val)
}
func ZifPow(ctx *php.Context, base types.Zval, exponent types.Zval) types.Zval {
	return php.OpPow(ctx, base, exponent)
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
func ZifLog(ctx *php.Context, number float64, _ zpp.Opt, base *float64) (float64, bool) {
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
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "base must be greater than 0")
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

func _parseStringToNumberZval(ctx *php.Context, s string, base int) (*types.Zval, bool) {
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

	cutoff := types.MaxLong / base
	cutlim := types.MaxLong % base
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
		php.Error(ctx, perr.E_DEPRECATED, "Invalid characters passed for attempted conversion, these have been ignored")
	}
	if isFloat {
		return types.NewZvalDouble(fnum), true
	} else {
		return types.NewZvalLong(inum), true
	}
}

func ZifBindec(ctx *php.Context, binaryNumber *types.Zval) (*types.Zval, bool) {
	php.ConvertToString(ctx, binaryNumber)
	return _parseStringToNumberZval(ctx, binaryNumber.String(), 2)
}
func ZifHexdec(ctx *php.Context, hexadecimalNumber *types.Zval) (*types.Zval, bool) {
	php.ConvertToString(ctx, hexadecimalNumber)
	return _parseStringToNumberZval(ctx, hexadecimalNumber.String(), 16)
}
func ZifOctdec(ctx *php.Context, octalNumber *types.Zval) (*types.Zval, bool) {
	php.ConvertToString(ctx, octalNumber)
	return _parseStringToNumberZval(ctx, octalNumber.String(), 8)
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
func ZifBaseConvert(ctx *php.Context, number types.Zval, frombase int, tobase int) (string, bool) {
	numberStr, ok := php.ZvalTryGetStr(ctx, number)
	if !ok { // fail
		return "", false
	}
	if frombase < 2 || frombase > 36 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid `from base' (%d)", frombase))
		return "", false
	}
	if tobase < 2 || tobase > 36 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid `to base' (%d)", tobase))
		return "", false
	}

	num, ok := _parseStringToNumberZval(ctx, numberStr, frombase)
	if !ok {
		return "", false
	}

	/* Don't try to convert +/- infinity */
	if num.IsDouble() && (math.IsInf(num.Double(), 1) || math.IsInf(num.Double(), -1)) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Number too large")
		return "", true
	}

	var base int
	if num.IsLong() {
		base = num.Long()
	} else { // num.IsDouble
		base = int(math.Floor(num.Double()))
	}
	return strconv.FormatInt(int64(base), tobase), true
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
	tmp := strconv.FormatFloat(d, 'f', dec, 64)

	var buf strings.Builder
	if isNegative {
		buf.WriteByte('-')
	}
	var pointPos int
	if pos := strings.Index(tmp, "."); pos >= 0 {
		if pos == 0 {
			buf.WriteByte('0')
		}
		pointPos = pos
	} else {
		pointPos = len(tmp)
	}

	for i := 0; i < pointPos; i++ {
		buf.WriteByte(tmp[i])
		if i+1 < pointPos && (pointPos-i-1)%3 == 0 {
			buf.WriteString(thousandSep)
		}
	}

	if pointPos < len(tmp) {
		buf.WriteString(decPoint)
		if pointPos+1+dec < len(tmp) {
			buf.WriteString(tmp[pointPos+1 : pointPos+1+dec])
		} else {
			buf.WriteString(tmp[pointPos+1:])
		}
	}

	return buf.String()
}
func ZifNumberFormat(number float64, _ zpp.Opt, numDecimalPlaces int, decSeparator_ *string, thousandsSeparator *string) string {
	var decPoint = lang.Option(decSeparator_, ".")
	var thousandSep = lang.Option(thousandsSeparator, ",")
	return phpMathNumberFormat(number, numDecimalPlaces, decPoint, thousandSep)
}
func ZifFmod(x float64, y float64) float64 {
	return math.Mod(x, y)
}
func ZifIntdiv(ctx *php.Context, dividend int, divisor int) int {
	if divisor == 0 {
		php.ThrowException(ctx, nil, "Division by zero", 0)
		return 0
	} else if divisor == -1 && dividend == types.MinLong {
		/* Prevent overflow error/crash ... really should not happen:
		   We don't return a float here as that violates function contract */
		php.ThrowException(ctx, nil, "Division of PHP_INT_MIN by -1 is not an integer", 0)
		return 0
	}
	return dividend / divisor
}

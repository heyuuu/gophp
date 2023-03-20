// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func PhpIntlog10abs(value float64) int {
	var result int
	value = fabs(value)
	if value < 1.0e-8 || value > 1.0e22 {
		result = int(floor(log10(value)))
	} else {
		var values []float64 = []float64{1.0e-8, 1.0e-7, 1.0e-6, 1.0e-5, 1.0e-4, 0.001, 0.01, 0.1, 1.0, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0e7, 1.0e8, 1.0e9, 1.0e10, 9.9999998e10, 1.0e12, 9.9999998e12, 1.0e14, 9.9999999e14, 1.00000003e16, 9.9999998e16, 9.9999998e17, 1.0e19, 1.0e20, 1.0e21, 1.0e22}

		/* Do a binary search with 5 steps */

		result = 15
		if value < values[result] {
			result -= 8
		} else {
			result += 8
		}
		if value < values[result] {
			result -= 4
		} else {
			result += 4
		}
		if value < values[result] {
			result -= 2
		} else {
			result += 2
		}
		if value < values[result] {
			result -= 1
		} else {
			result += 1
		}
		if value < values[result] {
			result -= 1
		}
		result -= 8
	}
	return result
}
func PhpIntpow10(power int) float64 {
	var powers []float64 = []float64{1.0, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0e7, 1.0e8, 1.0e9, 1.0e10, 9.9999998e10, 1.0e12, 9.9999998e12, 1.0e14, 9.9999999e14, 1.00000003e16, 9.9999998e16, 9.9999998e17, 1.0e19, 1.0e20, 1.0e21, 1.0e22}

	/* Not in lookup table */

	if power < 0 || power > 22 {
		return pow(10.0, float64(power))
	}
	return powers[power]
}
func PhpRoundHelper(value float64, mode int) float64 {
	var tmp_value float64
	if value >= 0.0 {
		tmp_value = floor(value + 0.5)
		if mode == PHP_ROUND_HALF_DOWN && value == -0.5+tmp_value || mode == PHP_ROUND_HALF_EVEN && value == 0.5+2*floor(tmp_value/2.0) || mode == PHP_ROUND_HALF_ODD && value == 0.5+2*floor(tmp_value/2.0)-1.0 {
			tmp_value = tmp_value - 1.0
		}
	} else {
		tmp_value = ceil(value - 0.5)
		if mode == PHP_ROUND_HALF_DOWN && value == 0.5+tmp_value || mode == PHP_ROUND_HALF_EVEN && value == -0.5+2*ceil(tmp_value/2.0) || mode == PHP_ROUND_HALF_ODD && value == -0.5+2*ceil(tmp_value/2.0)+1.0 {
			tmp_value = tmp_value + 1.0
		}
	}
	return tmp_value
}
func _phpMathRound(value float64, places int, mode int) float64 {
	var f1 float64
	var f2 float64
	var tmp_value float64
	var precision_places int
	if !(core.ZendFinite(value)) || value == 0.0 {
		return value
	}
	if places < core.INT_MIN+1 {
		places = core.INT_MIN + 1
	} else {
		places = places
	}
	precision_places = 14 - PhpIntlog10abs(value)
	f1 = PhpIntpow10(abs(places))

	/* If the decimal precision guaranteed by FP arithmetic is higher than
	   the requested places BUT is small enough to make sure a non-zero value
	   is returned, pre-round the result to the precision */

	if precision_places > places && precision_places-15 < places {
		var use_precision int64 = b.Cond(precision_places < core.INT_MIN+1, core.INT_MIN+1, precision_places)
		f2 = PhpIntpow10(abs(int(use_precision)))
		if use_precision >= 0 {
			tmp_value = value * f2
		} else {
			tmp_value = value / f2
		}

		/* preround the result (tmp_value will always be something * 1e14,
		   thus never larger than 1e15 here) */

		tmp_value = PhpRoundHelper(tmp_value, mode)
		use_precision = places - precision_places
		if use_precision < core.INT_MIN+1 {
			use_precision = core.INT_MIN + 1
		} else {
			use_precision = use_precision
		}

		/* now correctly move the decimal point */

		f2 = PhpIntpow10(abs(int(use_precision)))

		/* because places < precision_places */

		tmp_value = tmp_value / f2

		/* because places < precision_places */

	} else {

		/* adjust the value */

		if places >= 0 {
			tmp_value = value * f1
		} else {
			tmp_value = value / f1
		}

		/* This value is beyond our precision, so rounding it is pointless */

		if fabs(tmp_value) >= 9.9999999e14 {
			return value
		}

		/* This value is beyond our precision, so rounding it is pointless */

	}

	/* round the temp value */

	tmp_value = PhpRoundHelper(tmp_value, mode)

	/* see if it makes sense to use simple division to round the value */

	if abs(places) < 23 {
		if places > 0 {
			tmp_value = tmp_value / f1
		} else {
			tmp_value = tmp_value * f1
		}
	} else {

		/* Simple division can't be used since that will cause wrong results.
		   Instead, the number is converted to a string and back again using
		   strtod(). strtod() will return the nearest possible FP value for
		   that string. */

		var buf []byte
		core.Snprintf(buf, 39, "%15fe%d", tmp_value, -places)
		buf[39] = '0'
		tmp_value = zend.ZendStrtod(buf, nil)

		/* couldn't convert to string and back */

		if !(core.ZendFinite(tmp_value)) || core.ZendIsNaN(tmp_value) {
			tmp_value = value
		}

		/* couldn't convert to string and back */

	}
	return tmp_value
}
func PhpAsinh(z float64) float64 { return asinh(z) }
func PhpAcosh(x float64) float64 { return acosh(x) }
func PhpAtanh(z float64) float64 { return atanh(z) }
func PhpLog1p(x float64) float64 { return log1p(x) }
func PhpExpm1(x float64) float64 { return expm1(x) }
func ZifAbs(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			value = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertScalarToNumberEx(value)
	if value.IsType(types.IS_DOUBLE) {
		return_value.SetDouble(fabs(value.GetDval()))
		return
	} else if value.IsType(types.IS_LONG) {
		if value.GetLval() == zend.ZEND_LONG_MIN {
			return_value.SetDouble(-float64(zend.ZEND_LONG_MIN))
			return
		} else {
			return_value.SetLong(b.CondF(value.GetLval() < 0, func() int { return -(value.GetLval()) }, func() zend.ZendLong { return value.GetLval() }))
			return
		}
	}
	return_value.SetFalse()
	return
}
func ZifCeil(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			value = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertScalarToNumberEx(value)
	if value.IsType(types.IS_DOUBLE) {
		return_value.SetDouble(ceil(value.GetDval()))
		return
	} else if value.IsType(types.IS_LONG) {
		return_value.SetDouble(zend.ZvalGetDouble(value))
		return
	}
	return_value.SetFalse()
	return
}
func ZifFloor(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			value = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertScalarToNumberEx(value)
	if value.IsType(types.IS_DOUBLE) {
		return_value.SetDouble(floor(value.GetDval()))
		return
	} else if value.IsType(types.IS_LONG) {
		return_value.SetDouble(zend.ZvalGetDouble(value))
		return
	}
	return_value.SetFalse()
	return
}
func ZifRound(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	var places int = 0
	var precision zend.ZendLong = 0
	var mode zend.ZendLong = PHP_ROUND_HALF_UP
	var return_val float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			value = fp.ParseZval()
			fp.StartOptional()
			precision = fp.ParseLong()
			mode = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() >= 2 {
		if precision >= 0 {
			if precision > core.INT_MAX {
				places = core.INT_MAX
			} else {
				places = int(precision)
			}
		} else {
			if precision <= core.INT_MIN {
				places = core.INT_MIN + 1
			} else {
				places = int(precision)
			}
		}
	}
	zend.ConvertScalarToNumberEx(value)
	switch value.GetType() {
	case types.IS_LONG:

		/* Simple case - long that doesn't need to be rounded. */

		if places >= 0 {
			return_value.SetDouble(float64(value.GetLval()))
			return
		}
		fallthrough
	case types.IS_DOUBLE:
		if value.IsType(types.IS_LONG) {
			return_val = float64(value.GetLval())
		} else {
			return_val = value.GetDval()
		}
		return_val = _phpMathRound(return_val, int(places), int(mode))
		return_value.SetDouble(return_val)
		return
	default:
		return_value.SetFalse()
		return
	}
}
func ZifSin(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(sin(num))
	return
}
func ZifCos(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(cos(num))
	return
}
func ZifTan(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(tan(num))
	return
}
func ZifAsin(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(asin(num))
	return
}
func ZifAcos(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(acos(num))
	return
}
func ZifAtan(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(atan(num))
	return
}
func ZifAtan2(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num1 float64
	var num2 float64
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num1 = fp.ParseDouble()
			num2 = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(atan2(num1, num2))
	return
}
func ZifSinh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(sinh(num))
	return
}
func ZifCosh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(cosh(num))
	return
}
func ZifTanh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(tanh(num))
	return
}
func ZifAsinh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(PhpAsinh(num))
	return
}
func ZifAcosh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(PhpAcosh(num))
	return
}
func ZifAtanh(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(PhpAtanh(num))
	return
}
func ZifPi(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	return_value.SetDouble(M_PI)
	return
}
func ZifIsFinite(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dval = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, core.ZendFinite(dval))
	return
}
func ZifIsInfinite(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dval = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, core.ZendIsInf(dval))
	return
}
func ZifIsNan(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dval = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, core.ZendIsNaN(dval))
	return
}
func ZifPow(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zbase *types.Zval
	var zexp *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zbase = fp.ParseZval()
			zexp = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.PowFunction(return_value, zbase, zexp)
}
func ZifExp(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(exp(num))
	return
}
func ZifExpm1(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(PhpExpm1(num))
	return
}
func ZifLog1p(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(PhpLog1p(num))
	return
}
func ZifLog(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	var base float64 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			fp.StartOptional()
			base = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 1 {
		return_value.SetDouble(log(num))
		return
	}
	if base == 10.0 {
		return_value.SetDouble(log10(num))
		return
	}
	if base == 1.0 {
		return_value.SetDouble(zend.ZEND_NAN)
		return
	}
	if base <= 0.0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "base must be greater than 0")
		return_value.SetFalse()
		return
	}
	return_value.SetDouble(log(num) / log(base))
	return
}
func ZifLog10(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(log10(num))
	return
}
func ZifSqrt(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(sqrt(num))
	return
}
func ZifHypot(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num1 float64
	var num2 float64
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num1 = fp.ParseDouble()
			num2 = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(hypot(num1, num2))
	return
}
func ZifDeg2rad(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var deg float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			deg = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(deg / 180.0 * M_PI)
	return
}
func ZifRad2deg(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var rad float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			rad = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(rad / M_PI * 180)
	return
}
func _phpMathBasetolong(arg *types.Zval, base int) zend.ZendLong {
	var num zend.ZendLong = 0
	var digit zend.ZendLong
	var onum zend.ZendLong
	var i zend.ZendLong
	var c byte
	var s *byte
	if arg.GetType() != types.IS_STRING || base < 2 || base > 36 {
		return 0
	}
	s = arg.GetStr().GetVal()
	for i = arg.GetStr().GetLen(); i > 0; i-- {
		*s++
		c = (*s) - 1
		if b.Cond(b.Cond(c >= '0' && c <= '9', c-'0', c >= 'A' && c <= 'Z'), c-'A'+10, c >= 'a' && c <= 'z') {
			digit = c - 'a' + 10
		} else {
			digit = base
		}
		if digit >= base {
			continue
		}
		onum = num
		num = num*base + digit
		if num > onum {
			continue
		}
		core.PhpErrorDocref(nil, faults.E_WARNING, "Number '%s' is too big to fit in long", s)
		return zend.ZEND_LONG_MAX
	}
	return num
}
func _phpMathBasetozval(arg *types.Zval, base int, ret *types.Zval) int {
	var num zend.ZendLong = 0
	var fnum float64 = 0
	var mode int = 0
	var c byte
	var s *byte
	var e *byte
	var cutoff zend.ZendLong
	var cutlim int
	var invalidchars int = 0
	if arg.GetType() != types.IS_STRING || base < 2 || base > 36 {
		return types.FAILURE
	}
	s = arg.GetStr().GetVal()
	e = s + arg.GetStr().GetLen()

	for s < e && isspace(*s) {
		s++
	}

	/* Skip trailing whitespace */

	for s < e && isspace(*(e - 1)) {
		e--
	}
	if e-s >= 2 {
		if base == 16 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			s += 2
		}
		if base == 8 && s[0] == '0' && (s[1] == 'o' || s[1] == 'O') {
			s += 2
		}
		if base == 2 && s[0] == '0' && (s[1] == 'b' || s[1] == 'B') {
			s += 2
		}
	}
	cutoff = zend.ZEND_LONG_MAX / base
	cutlim = zend.ZEND_LONG_MAX % base
	for s < e {
		*s++
		c = (*s) - 1

		/* might not work for EBCDIC */

		if c >= '0' && c <= '9' {
			c -= '0'
		} else if c >= 'A' && c <= 'Z' {
			c -= 'A' - 10
		} else if c >= 'a' && c <= 'z' {
			c -= 'a' - 10
		} else {
			invalidchars++
			continue
		}
		if c >= base {
			invalidchars++
			continue
		}
		switch mode {
		case 0:
			if num < cutoff || num == cutoff && c <= cutlim {
				num = num*base + c
				break
			} else {
				fnum = float64(num)
				mode = 1
			}
			fallthrough
		case 1:
			fnum = fnum*base + c
		}
	}
	if invalidchars > 0 {
		faults.Error(faults.E_DEPRECATED, "Invalid characters passed for attempted conversion, these have been ignored")
	}
	if mode == 1 {
		ret.SetDouble(fnum)
	} else {
		ret.SetLong(num)
	}
	return types.SUCCESS
}
func _phpMathLongtobase(arg *types.Zval, base int) *types.String {
	var digits []byte = "0123456789abcdefghijklmnopqrstuvwxyz"
	var buf []byte
	var ptr *byte
	var end *byte
	var value zend.ZendUlong
	if arg.GetType() != types.IS_LONG || base < 2 || base > 36 {
		return types.ZSTR_EMPTY_ALLOC()
	}
	value = arg.GetLval()
	ptr = buf + b.SizeOf("buf") - 1
	end = ptr
	*ptr = '0'
	for {
		b.Assert(ptr > buf)
		*(b.PreDec(&ptr)) = digits[value%base]
		value /= base
		if value == 0 {
			break
		}
	}
	return types.NewString(b.CastStr(ptr, end-ptr))
}
func _phpMathZvaltobase(arg *types.Zval, base int) *types.String {
	var digits []byte = "0123456789abcdefghijklmnopqrstuvwxyz"
	if arg.GetType() != types.IS_LONG && arg.GetType() != types.IS_DOUBLE || base < 2 || base > 36 {
		return types.ZSTR_EMPTY_ALLOC()
	}
	if arg.IsType(types.IS_DOUBLE) {
		var fvalue float64 = floor(arg.GetDval())
		var ptr *byte
		var end *byte
		var buf []byte

		/* Don't try to convert +/- infinity */

		if fvalue == zend.ZEND_INFINITY || fvalue == -zend.ZEND_INFINITY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Number too large")
			return types.ZSTR_EMPTY_ALLOC()
		}
		ptr = buf + b.SizeOf("buf") - 1
		end = ptr
		*ptr = '0'
		for {
			*(b.PreDec(&ptr)) = digits[int(fmod(fvalue, base))]
			fvalue /= base
			if !(ptr > buf && fabs(fvalue) >= 1) {
				break
			}
		}
		return types.NewString(b.CastStr(ptr, end-ptr))
	}
	return _phpMathLongtobase(arg, base)
}
func ZifBindec(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertToStringEx(arg)
	if _phpMathBasetozval(arg, 2, return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifHexdec(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertToStringEx(arg)
	if _phpMathBasetozval(arg, 16, return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifOctdec(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ConvertToStringEx(arg)
	if _phpMathBasetozval(arg, 8, return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifDecbin(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if arg.GetType() != types.IS_LONG {
		zend.ConvertToLong(arg)
	}
	result = _phpMathLongtobase(arg, 2)
	return_value.SetString(result)
	return
}
func ZifDecoct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if arg.GetType() != types.IS_LONG {
		zend.ConvertToLong(arg)
	}
	result = _phpMathLongtobase(arg, 8)
	return_value.SetString(result)
	return
}
func ZifDechex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if arg.GetType() != types.IS_LONG {
		zend.ConvertToLong(arg)
	}
	result = _phpMathLongtobase(arg, 16)
	return_value.SetString(result)
	return
}
func ZifBaseConvert(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var number *types.Zval
	var temp types.Zval
	var frombase zend.ZendLong
	var tobase zend.ZendLong
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			number = fp.ParseZval()
			frombase = fp.ParseLong()
			tobase = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if zend.TryConvertToString(number) == 0 {
		return
	}
	if frombase < 2 || frombase > 36 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid `from base' ("+zend.ZEND_LONG_FMT+")", frombase)
		return_value.SetFalse()
		return
	}
	if tobase < 2 || tobase > 36 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid `to base' ("+zend.ZEND_LONG_FMT+")", tobase)
		return_value.SetFalse()
		return
	}
	if _phpMathBasetozval(number, int(frombase), &temp) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	result = _phpMathZvaltobase(&temp, int(tobase))
	return_value.SetString(result)
}
func _phpMathNumberFormat(d float64, dec int, dec_point byte, thousand_sep byte) *types.String {
	return _phpMathNumberFormatEx(d, dec, &dec_point, 1, &thousand_sep, 1)
}
func _phpMathNumberFormatEx(
	d float64,
	dec int,
	dec_point *byte,
	dec_point_len int,
	thousand_sep *byte,
	thousand_sep_len int,
) *types.String {
	var res *types.String
	var tmpbuf *types.String
	var s *byte
	var t *byte
	var dp *byte
	var integral int
	var reslen int = 0
	var count int = 0
	var is_negative int = 0
	if d < 0 {
		is_negative = 1
		d = -d
	}
	dec = b.Max(0, dec)
	d = _phpMathRound(d, dec, PHP_ROUND_HALF_UP)
	tmpbuf = core.Strpprintf(0, "%.*F", dec, d)
	if tmpbuf == nil {
		return nil
	} else if !(isdigit(int(tmpbuf.GetVal()[0]))) {
		return tmpbuf
	}

	/* Check if the number is no longer negative after rounding */

	if is_negative != 0 && d == 0 {
		is_negative = 0
	}

	/* find decimal point, if expected */

	if dec != 0 {
		dp = strpbrk(tmpbuf.GetVal(), ".,")
	} else {
		dp = nil
	}

	/* calculate the length of the return buffer */

	if dp != nil {
		integral = dp - tmpbuf.GetVal()
	} else {

		/* no decimal point was found */

		integral = tmpbuf.GetLen()

		/* no decimal point was found */

	}

	/* allow for thousand separators */

	if thousand_sep != nil {
		integral = zend.ZendSafeAddmult((integral-1)/3, thousand_sep_len, integral, "number formatting")
	}
	reslen = integral
	if dec != 0 {
		reslen += dec
		if dec_point != nil {
			reslen = zend.ZendSafeAddmult(reslen, 1, dec_point_len, "number formatting")
		}
	}

	/* add a byte for minus sign */

	if is_negative != 0 {
		reslen++
	}
	res = types.ZendStringAlloc(reslen, 0)
	s = tmpbuf.GetVal() + tmpbuf.GetLen() - 1
	t = res.GetVal() + reslen
	b.PostDec(&(*t)) = '0'

	/* copy the decimal places.
	 * Take care, as the sprintf implementation may return less places than
	 * we requested due to internal buffer limitations */

	if dec != 0 {
		var declen int = b.Cond(dp != nil, s-dp, 0)
		var topad int = int(b.Cond(dec > declen, dec-declen, 0))

		/* pad with '0's */

		for b.PostDec(&topad) {
			b.PostDec(&(*t)) = '0'
		}
		if dp != nil {
			s -= declen + 1
			t -= declen

			/* now copy the chars after the point */

			memcpy(t+1, dp+1, declen)

			/* now copy the chars after the point */

		}

		/* add decimal point */

		if dec_point != nil {
			t -= dec_point_len
			memcpy(t+1, dec_point, dec_point_len)
		}

		/* add decimal point */

	}

	/* copy the numbers before the decimal point, adding thousand
	 * separator every three digits */

	for s >= tmpbuf.GetVal() {
		*s--
		b.PostDec(&(*t)) = (*s) + 1
		if thousand_sep != nil && b.PreInc(&count)%3 == 0 && s >= tmpbuf.GetVal() {
			t -= thousand_sep_len
			memcpy(t+1, thousand_sep, thousand_sep_len)
		}
	}

	/* and a minus sign, if needed */

	if is_negative != 0 {
		b.PostDec(&(*t)) = '-'
	}
	res.SetLen(reslen)
	types.ZendStringReleaseEx(tmpbuf, 0)
	return res
}
func ZifNumberFormat(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num float64
	var dec zend.ZendLong = 0
	var thousand_sep *byte = nil
	var dec_point *byte = nil
	var thousand_sep_chr byte = ','
	var dec_point_chr byte = '.'
	var thousand_sep_len int = 0
	var dec_point_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseDouble()
			fp.StartOptional()
			dec = fp.ParseLong()
			dec_point, dec_point_len = fp.ParseStringEx(true)
			thousand_sep, thousand_sep_len = fp.ParseStringEx(true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	switch executeData.NumArgs() {
	case 1:
		return_value.SetString(_phpMathNumberFormat(num, 0, dec_point_chr, thousand_sep_chr))
		return
	case 2:
		return_value.SetString(_phpMathNumberFormat(num, int(dec), dec_point_chr, thousand_sep_chr))
		return
	case 4:
		if dec_point == nil {
			dec_point = &dec_point_chr
			dec_point_len = 1
		}
		if thousand_sep == nil {
			thousand_sep = &thousand_sep_chr
			thousand_sep_len = 1
		}
		return_value.SetString(_phpMathNumberFormatEx(num, int(dec), dec_point, dec_point_len, thousand_sep, thousand_sep_len))
	default:
		zend.ZendWrongParamCount()
		return
	}
}
func ZifFmod(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num1 float64
	var num2 float64
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num1 = fp.ParseDouble()
			num2 = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(fmod(num1, num2))
	return
}
func ZifIntdiv(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dividend zend.ZendLong
	var divisor zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dividend = fp.ParseLong()
			divisor = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if divisor == 0 {
		faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Division by zero")
		return
	} else if divisor == -1 && dividend == zend.ZEND_LONG_MIN {

		/* Prevent overflow error/crash ... really should not happen:
		   We don't return a float here as that violates function contract */

		faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Division of PHP_INT_MIN by -1 is not an integer")
		return
	}
	return_value.SetLong(dividend / divisor)
	return
}

package zend

import (
	"math"
)

func isArgUseStrictTypes() bool { return CurrEX().IsArgUseStrictTypes() }

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }
func parseArgNull[T any]() (T, bool, bool)      { var temp T; return temp, true, true }
func parseArgFail[T any]() (T, bool, bool)      { var temp T; return temp, false, false }

func ParseArgBool(arg *Zval, checkNull bool, strict bool) (dest bool, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsTrue() {
		return parseArgSucc(true)
	} else if arg.IsFalse() {
		return parseArgSucc(false)
	}

	// weak parse
	if !strict {
		dest, ok = ParseArgBoolWeak(arg)
	}

	return
}

func ParseArgBoolWeak(arg *Zval) (dest bool, ok bool) {
	if arg.GetType() <= IS_STRING {
		return ZendIsTrueEx(arg), true
	}
	return false, false
}

func ParseArgLong(arg *Zval, checkNull bool, cap bool, strict bool) (dest ZendLong, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetLval())
	}

	// weak parse
	if !strict {
		dest, ok = ParseArgLongWeak(arg, cap)
	}

	return
}

func ParseArgLongWeak(arg *Zval, cap bool) (dest ZendLong, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = ConvertNumericStrAsZval(arg.GetStr().GetStr(), ConvertNoticeOnErrors)
		if arg == nil {
			return // fail
		}
		if EG__().GetException() != nil {
			return // fail
		}
	}

	switch arg.GetType() {
	case IS_UNDEF, IS_NULL, IS_FALSE:
		dest = 0
	case IS_TRUE:
		dest = 1
	case IS_LONG:
		dest = arg.GetLval()
	case IS_DOUBLE:
		return parseArgLongWeak_DvalToLval(arg.GetDval(), cap)
	default:
		return // fail
	}
	// success
	return dest, true
}

func parseArgLongWeak_DvalToLval(dval float64, cap bool) (ZendLong, bool) {
	if math.IsNaN(dval) {
		return 0, false
	}
	if cap {
		return ZendDvalToLvalCap(dval), true
	} else {
		if !ZEND_DOUBLE_FITS_LONG(dval) {
			return 0, false
		}
		return ZendDvalToLval(dval), true
	}
}

func ParseArgDouble(arg *Zval, checkNull bool, strict bool) (dest float64, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetDval())
	} else if arg.IsLong() {
		return parseArgSucc(float64(arg.GetLval()))
	}

	// weak parse
	if !strict {
		dest, ok = ParseArgDoubleWeak(arg)
	}

	return
}

func ParseArgDoubleWeak(arg *Zval) (dest float64, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = ConvertNumericStrAsZval(arg.GetStr().GetStr(), ConvertNoticeOnErrors)
		if arg == nil {
			return // fail
		}
		if EG__().GetException() != nil {
			return // fail
		}
	}

	switch arg.GetType() {
	case IS_UNDEF, IS_NULL, IS_FALSE:
		dest = 0
	case IS_TRUE:
		dest = 1
	case IS_LONG:
		dest = float64(arg.GetLval())
	case IS_DOUBLE:
		dest = arg.GetDval()
	default:
		return // fail
	}
	// success
	return dest, true
}

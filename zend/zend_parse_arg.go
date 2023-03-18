package zend

import (
	"math"
	"sik/zend/types"
)

func isArgUseStrictTypes() bool { return CurrEX().IsArgUseStrictTypes() }
func isArgUseWeakTypes() bool   { return !CurrEX().IsArgUseStrictTypes() }

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }
func parseArgNull[T any]() (T, bool, bool)      { var temp T; return temp, true, true }
func parseArgFail[T any]() (T, bool, bool)      { var temp T; return temp, false, false }

func ParseArgBool(arg *types.Zval, checkNull bool) (dest bool, isNull bool, ok bool) {
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
	if isArgUseWeakTypes() {
		dest, ok = ParseArgBoolWeak(arg)
	}

	return
}

func ParseArgBoolWeak(arg *types.Zval) (dest bool, ok bool) {
	if arg.GetType() <= types.IS_STRING {
		return ZendIsTrueEx(arg), true
	}
	return false, false
}

func ParseArgLong(arg *types.Zval, checkNull bool, cap bool) (dest ZendLong, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetLval())
	}

	// weak parse
	if isArgUseWeakTypes() {
		dest, ok = ParseArgLongWeak(arg, cap)
	}

	return
}

func ParseArgLongWeak(arg *types.Zval, cap bool) (dest ZendLong, ok bool) {
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
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		dest = 0
	case types.IS_TRUE:
		dest = 1
	case types.IS_LONG:
		dest = arg.GetLval()
	case types.IS_DOUBLE:
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

func ParseArgDouble(arg *types.Zval, checkNull bool) (dest float64, isNull bool, ok bool) {
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
	if isArgUseWeakTypes() {
		dest, ok = ParseArgDoubleWeak(arg)
	}

	return
}

func ParseArgDoubleWeak(arg *types.Zval) (dest float64, ok bool) {
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
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		dest = 0
	case types.IS_TRUE:
		dest = 1
	case types.IS_LONG:
		dest = float64(arg.GetLval())
	case types.IS_DOUBLE:
		dest = arg.GetDval()
	default:
		return // fail
	}
	// success
	return dest, true
}

func ParseArgStr(arg *types.Zval, checkNull bool, strict bool) (dest *types.ZendString, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsString() {
		return parseArgSucc(arg.GetStr())
	}

	// weak parse
	if !strict {
		dest, ok = ParseArgStrWeak(arg)
	}

	return
}

func ParseArgStrWeak(arg *types.Zval) (*types.ZendString, bool) {
	if arg.GetType() < types.IS_STRING {
		ConvertToString(arg)
		return arg.GetStr(), true
	} else if arg.IsString() {
		return arg.GetStr(), true
	} else if arg.IsObject() {
		handlers := arg.GetObj().GetHandlers()
		if castFunc := handlers.GetCastObject(); castFunc != nil {
			var obj types.Zval
			if castFunc(arg, &obj, types.IS_STRING) == types.SUCCESS {
				ZvalPtrDtor(arg)
				types.ZVAL_COPY_VALUE(arg, &obj)
				return arg.GetStr(), true
			}
		} else if getFunc := handlers.GetGet(); getFunc != nil {
			var rv types.Zval
			var z *types.Zval = getFunc(arg, &rv)
			if z.GetType() != types.IS_OBJECT {
				ZvalPtrDtor(arg)
				if z.IsString() {
					types.ZVAL_COPY_VALUE(arg, z)
				} else {
					arg.SetString(ZvalGetStringFunc(z))
					ZvalPtrDtor(z)
				}
				return arg.GetStr(), true
			}
			ZvalPtrDtor(z)
		}
		return nil, false
	} else {
		return nil, false
	}
}

package argparse

import (
	"math"
	"sik/zend"
	"sik/zend/types"
)

func isArgUseWeakTypes() bool { return !zend.CurrEX().IsArgUseStrictTypes() }

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }
func parseArgNull[T any]() (T, bool, bool)      { var temp T; return temp, true, true }
func parseArgFail[T any]() (T, bool, bool)      { var temp T; return temp, false, false }

func ParseBool(arg *types.Zval, checkNull bool) (dest bool, isNull bool, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return dest, true, true
	}

	// base parse
	if arg.IsTrue() {
		return parseArgSucc(true)
	} else if arg.IsFalse() {
		return parseArgSucc(false)
	}

	// weak parse
	if isArgUseWeakTypes() {
		dest, ok = ParseBoolWeak(arg)
	}

	return
}

func ParseBoolWeak(arg *types.Zval) (dest bool, ok bool) {
	if arg.GetType() <= types.IS_STRING {
		return zend.ZendIsTrueEx(arg), true
	}
	return false, false
}

func ParseLong(arg *types.Zval, checkNull bool, cap bool) (dest int, isNull bool, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return dest, true, true
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetLval())
	}

	// weak parse
	if isArgUseWeakTypes() {
		dest, ok = ParseLongWeak(arg, cap)
	}

	return
}

func ParseLongWeak(arg *types.Zval, cap bool) (dest int, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = zend.ConvertNumericStrAsZval(arg.GetStr().GetStr(), zend.ConvertNoticeOnErrors)
		if arg == nil {
			return // fail
		}
		if zend.EG__().GetException() != nil {
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
		return parseArgWeak_DvalToLval(arg.GetDval(), cap)
	default:
		return // fail
	}
	// success
	return dest, true
}

func parseArgWeak_DvalToLval(dval float64, cap bool) (zend.ZendLong, bool) {
	if math.IsNaN(dval) {
		return 0, false
	}
	if cap {
		return zend.ZendDvalToLvalCap(dval), true
	} else {
		if !zend.ZEND_DOUBLE_FITS_LONG(dval) {
			return 0, false
		}
		return zend.ZendDvalToLval(dval), true
	}
}

func ParseDouble(arg *types.Zval, checkNull bool) (dest float64, isNull bool, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return dest, true, true
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetDval())
	} else if arg.IsLong() {
		return parseArgSucc(float64(arg.GetLval()))
	}

	// weak parse
	if isArgUseWeakTypes() {
		dest, ok = ParseDoubleWeak(arg)
	}

	return
}

func ParseDoubleWeak(arg *types.Zval) (dest float64, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = zend.ConvertNumericStrAsZval(arg.GetStr().GetStr(), zend.ConvertNoticeOnErrors)
		if arg == nil {
			return // fail
		}
		if zend.EG__().GetException() != nil {
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

/**
 * 与 int/float 等类型不同，为空时 *dest 直接为 nil，不需单独的 is_null 字符安
 */
func ParseZStr(arg *types.Zval, checkNull bool) (dest *types.ZendString, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return nil, true
	}

	// base parse
	if arg.IsString() {
		return arg.GetStr(), true
	}

	// weak parse
	if isArgUseWeakTypes() {
		return ParseZStrWeak(arg)
	}

	// fail
	return
}

func ParseZStrWeak(arg *types.Zval) (*types.ZendString, bool) {
	if arg.GetType() < types.IS_STRING {
		zend.ConvertToString(arg)
		return arg.GetStr(), true
	} else if arg.IsString() {
		return arg.GetStr(), true
	} else if arg.IsObject() {
		handlers := arg.GetObj().GetHandlers()
		if castFunc := handlers.GetCastObject(); castFunc != nil {
			var obj types.Zval
			if castFunc(arg, &obj, types.IS_STRING) == types.SUCCESS {
				zend.ZvalPtrDtor(arg)
				types.ZVAL_COPY_VALUE(arg, &obj)
				return arg.GetStr(), true
			}
		} else if getFunc := handlers.GetGet(); getFunc != nil {
			var rv types.Zval
			var z *types.Zval = getFunc(arg, &rv)
			if z.GetType() != types.IS_OBJECT {
				zend.ZvalPtrDtor(arg)
				if z.IsString() {
					types.ZVAL_COPY_VALUE(arg, z)
				} else {
					arg.SetString(zend.ZvalGetStringFunc(z))
					zend.ZvalPtrDtor(z)
				}
				return arg.GetStr(), true
			}
			zend.ZvalPtrDtor(z)
		}
		return nil, false
	} else {
		return nil, false
	}
}

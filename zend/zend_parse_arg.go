package zend

import "math"

func isArgUseWeakTypes() bool { return !CurrEX().IsArgUseStrictTypes() }

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }
func parseArgNull[T any]() (T, bool, bool)      { var temp T; return temp, true, true }
func parseArgFail[T any]() (T, bool, bool)      { var temp T; return temp, false, false }

func ParseArgBool(arg *Zval, checkNull bool, weak bool) (dest bool, isNull bool, ok bool) {
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
	if weak {
		if arg.GetType() <= IS_STRING {
			return parseArgSucc(ZendIsTrueEx(arg))
		}
	}

	return
}

func ParseArgLong(arg *Zval, checkNull bool, cap bool, weak bool) (dest ZendLong, isNull bool, ok bool) {
	// check null
	if isNull = checkNull && arg.IsNull(); isNull {
		return
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.GetLval())
	}

	// weak parse
	if weak {
		if cap {
			dest, ok = ParseArgLongCapWeak(arg)
		} else {
			dest, ok = ParseArgLongWeak(arg)
		}
	}

	return
}

func ParseArgLongCapWeak(arg *Zval) (dest ZendLong, ok bool) {
	switch arg.GetType() {
	case IS_UNDEF, IS_NULL, IS_FALSE:
		return 0, true
	case IS_TRUE:
		return 1, true
	case IS_LONG:
		return arg.GetLval(), true
	case IS_DOUBLE:
		if math.IsNaN(arg.GetDval()) {
			return 0, false
		}
		return ZendDvalToLvalCap(arg.GetDval()), true
	case IS_STRING:
		// todo
		return
	default:
		return 0, false
	}
}

func ParseArgLongWeak(arg *Zval) (dest ZendLong, ok bool) {
	// todo
	return
}

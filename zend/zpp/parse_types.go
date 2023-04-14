package zpp

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"math"
)

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }

func ParseBool(arg *types.Zval, checkNull bool, weak bool) (dest bool, isNull bool, ok bool) {
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
	if weak {
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

func ParseLong(arg *types.Zval, checkNull bool, cap bool, weak bool) (dest int, isNull bool, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return dest, true, true
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.Long())
	}

	// weak parse
	if weak {
		dest, ok = ParseLongWeak(arg, cap)
	}

	return
}

func ParseLongWeak(arg *types.Zval, cap bool) (dest int, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = zend.ConvertNumericStrAsZval(arg.String().GetStr(), zend.ConvertNoticeOnErrors)
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
		dest = arg.Long()
	case types.IS_DOUBLE:
		return parseArgWeak_DvalToLval(arg.Double(), cap)
	default:
		return // fail
	}
	// success
	return dest, true
}

func parseArgWeak_DvalToLval(dval float64, cap bool) (int, bool) {
	if math.IsNaN(dval) {
		return 0, false
	}
	if cap {
		return zend.DvalToLvalCap(dval), true
	} else {
		if !zend.DoubleFitsLong(dval) {
			return 0, false
		}
		return zend.DvalToLval(dval), true
	}
}

func ParseDouble(arg *types.Zval, checkNull bool, weak bool) (dest float64, isNull bool, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return dest, true, true
	}

	// base parse
	if arg.IsLong() {
		return parseArgSucc(arg.Double())
	} else if arg.IsLong() {
		return parseArgSucc(float64(arg.Long()))
	}

	// weak parse
	if weak {
		dest, ok = ParseDoubleWeak(arg)
	}

	return
}

func ParseDoubleWeak(arg *types.Zval) (dest float64, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = zend.ConvertNumericStrAsZval(arg.String().GetStr(), zend.ConvertNoticeOnErrors)
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
		dest = float64(arg.Long())
	case types.IS_DOUBLE:
		dest = arg.Double()
	default:
		return // fail
	}
	// success
	return dest, true
}

/**
 * 与 int/float 等类型不同，为空时 *dest 直接为 nil，不需单独的 is_null 字符安
 */
func ParseZStr(arg *types.Zval, checkNull bool, weak bool) (dest *types.String, ok bool) {
	// check null
	if checkNull && arg.IsNull() {
		return nil, true
	}

	// base parse
	if arg.IsString() {
		return arg.String(), true
	}

	// weak parse
	if weak {
		return ParseZStrWeak(arg)
	}

	// fail
	return
}

func ParseZStrWeak(arg *types.Zval) (*types.String, bool) {
	if arg.GetType() < types.IS_STRING {
		zend.ConvertToString(arg)
		return arg.String(), true
	} else if arg.IsString() {
		return arg.String(), true
	} else if arg.IsObject() {
		handlers := arg.Object().GetHandlers()
		if castFunc := handlers.GetCastObject(); castFunc != nil {
			var obj types.Zval
			if castFunc(arg, &obj, types.IS_STRING) == types.SUCCESS {
				zend.ZvalPtrDtor(arg)
				types.ZVAL_COPY_VALUE(arg, &obj)
				return arg.String(), true
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
				return arg.String(), true
			}
			zend.ZvalPtrDtor(z)
		}
		return nil, false
	} else {
		return nil, false
	}
}

func ParseStrPtr(arg *types.Zval, checkNull bool, weak bool) (str *byte, len_ int, ok bool) {
	val, ok := ParseZStr(arg, checkNull, weak)
	if !ok {
		return nil, 0, false
	}

	if checkNull && val == nil {
		return nil, 0, true
	} else {
		return val.GetValPtr(), val.GetLen(), true
	}
}

// @see Micro CHECK_NULL_PATH
func checkNullPath(s string) bool {
	// todo 待确认此逻辑的生效方式 (当前代码一直为false)
	// 可能生效方式: 确认字符串是二进制安全的(即不包含 \0 字符)
	return len(s) != b.Strlen(s)
}

func ParsePathStr(arg *types.Zval, checkNull bool, weak bool) (dest *types.String, ok bool) {
	dest, ok = ParseZStr(arg, checkNull, weak)
	if !ok {
		return
	}

	if dest != nil && checkNullPath(dest.GetStr()) {
		return nil, false
	}

	return
}

func ParsePathStrPtr(arg *types.Zval, checkNull bool, weak bool) (str *byte, len_ int, ok bool) {
	val, ok := ParsePathStr(arg, checkNull, weak)
	if !ok {
		return nil, 0, false
	}

	if checkNull && val == nil {
		return nil, 0, true
	} else {
		return val.GetValPtr(), val.GetLen(), true
	}
}

func ParseArray(arg *types.Zval, checkNull bool, orObject bool) (dest *types.Zval, ok bool) {
	if arg.IsArray() || (orObject && arg.IsObject()) {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseArrayHt(arg *types.Zval, checkNull bool, orObject bool, separate bool) (dest *types.Array, ok bool) {
	if arg.IsArray() {
		return arg.Array(), true
	} else if orObject && arg.IsObject() {
		if separate && types.Z_OBJ_P(arg).GetProperties() != nil && types.Z_OBJ_P(arg).GetProperties().GetRefcount() > 1 {
			if (types.Z_OBJ_P(arg).GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
				types.Z_OBJ_P(arg).GetProperties().DelRefcount()
			}
			types.Z_OBJ_P(arg).SetProperties(types.ZendArrayDup(types.Z_OBJ_P(arg).GetProperties()))
		}
		return types.Z_OBJ_HT_P(arg).GetGetProperties()(arg), true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseObject(arg *types.Zval, ce *types.ClassEntry, checkNull bool) (dest *types.Zval, ok bool) {
	if arg.IsObject() && (ce == nil || zend.InstanceofFunction(types.Z_OBJCE_P(arg), ce) != 0) {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseResource(arg *types.Zval, checkNull bool) (dest *types.Zval, ok bool) {
	if arg.IsResource() {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseFunc(arg *types.Zval, dest_fci *types.ZendFcallInfo, dest_fcc *types.ZendFcallInfoCache, checkNull bool) (error *string, ok bool) {
	if checkNull && arg.IsNull() {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		return nil, true
	}

	// notice: 此处在成功时 error 也有可能不为 nil (例如在产生 Deprecated 信息时)
	state := zend.ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, &error)
	return error, state == types.SUCCESS
}

func ParseZvalDeref(arg *types.Zval, checkNull bool) (dest *types.Zval) {
	if checkNull && arg.IsNull() {
		return nil
	} else {
		return arg
	}
}

func ParseClass(arg *types.Zval, baseCe *types.ClassEntry, num int, checkNull bool) (ce *types.ClassEntry, ok bool) {
	if checkNull && arg.IsNull() {
		return nil, true
	}

	if zend.TryConvertToString(arg) == 0 {
		return nil, false
	}
	ce = zend.ZendLookupClass(arg.String())
	if baseCe != nil {
		if ce == nil || zend.InstanceofFunction(ce, baseCe) == 0 {
			faults.InternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a class name derived from %s, '%s' given", zend.GetActiveCalleeName(), num, baseCe.Name(), arg.GetStrVal())
			return nil, false
		}
	}
	if ce == nil {
		faults.InternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a valid class name, '%s' given", zend.GetActiveCalleeName(), num, arg.String().GetVal())
		return nil, false
	}
	return ce, true
}

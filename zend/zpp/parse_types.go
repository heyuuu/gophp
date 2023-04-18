package zpp

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"math"
)

func parseArgSucc[T any](val T) (T, bool, bool) { return val, false, true }

func ParseBool(arg *types2.Zval, checkNull bool, weak bool) (dest bool, isNull bool, ok bool) {
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

func ParseBoolWeak(arg *types2.Zval) (dest bool, ok bool) {
	if arg.GetType() <= types2.IS_STRING {
		return zend.ZendIsTrueEx(arg), true
	}
	return false, false
}

func ParseLong(arg *types2.Zval, checkNull bool, cap bool, weak bool) (dest int, isNull bool, ok bool) {
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

func ParseLongWeak(arg *types2.Zval, cap bool) (dest int, ok bool) {
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
	case types2.IS_UNDEF, types2.IS_NULL, types2.IS_FALSE:
		dest = 0
	case types2.IS_TRUE:
		dest = 1
	case types2.IS_LONG:
		dest = arg.Long()
	case types2.IS_DOUBLE:
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

func ParseDouble(arg *types2.Zval, checkNull bool, weak bool) (dest float64, isNull bool, ok bool) {
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

func ParseDoubleWeak(arg *types2.Zval) (dest float64, ok bool) {
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
	case types2.IS_UNDEF, types2.IS_NULL, types2.IS_FALSE:
		dest = 0
	case types2.IS_TRUE:
		dest = 1
	case types2.IS_LONG:
		dest = float64(arg.Long())
	case types2.IS_DOUBLE:
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
func ParseZStr(arg *types2.Zval, checkNull bool, weak bool) (dest *types2.String, ok bool) {
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

func ParseZStrWeak(arg *types2.Zval) (*types2.String, bool) {
	if arg.GetType() < types2.IS_STRING {
		zend.ConvertToString(arg)
		return arg.String(), true
	} else if arg.IsString() {
		return arg.String(), true
	} else if arg.IsObject() {
		handlers := arg.Object().GetHandlers()
		if castFunc := handlers.GetCastObject(); castFunc != nil {
			var obj types2.Zval
			if castFunc(arg, &obj, types2.IS_STRING) == types2.SUCCESS {
				// zend.ZvalPtrDtor(arg)
				types2.ZVAL_COPY_VALUE(arg, &obj)
				return arg.String(), true
			}
		} else if getFunc := handlers.GetGet(); getFunc != nil {
			var rv types2.Zval
			var z *types2.Zval = getFunc(arg, &rv)
			if z.GetType() != types2.IS_OBJECT {
				// zend.ZvalPtrDtor(arg)
				if z.IsString() {
					arg.CopyValueFrom(z)
				} else {
					arg.SetString(zend.ZvalGetStringFunc(z))
					// zend.ZvalPtrDtor(z)
				}
				return arg.String(), true
			}
			// zend.ZvalPtrDtor(z)
		}
		return nil, false
	} else {
		return nil, false
	}
}

// @see Micro CHECK_NULL_PATH
func checkNullPath(s string) bool {
	// todo 待确认此逻辑的生效方式 (当前代码一直为false)
	// 可能生效方式: 确认字符串是二进制安全的(即不包含 \0 字符)
	return len(s) != b.Strlen(s)
}

func ParsePathStr(arg *types2.Zval, checkNull bool, weak bool) (dest *types2.String, ok bool) {
	dest, ok = ParseZStr(arg, checkNull, weak)
	if !ok {
		return
	}

	if dest != nil && checkNullPath(dest.GetStr()) {
		return nil, false
	}

	return
}

func ParseArray(arg *types2.Zval, checkNull bool, orObject bool) (dest *types2.Zval, ok bool) {
	if arg.IsArray() || (orObject && arg.IsObject()) {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseArrayHt(arg *types2.Zval, checkNull bool, orObject bool, separate bool) (dest *types2.Array, ok bool) {
	if arg.IsArray() {
		return arg.Array(), true
	} else if orObject && arg.IsObject() {
		if separate && types2.Z_OBJ_P(arg).GetProperties() != nil && types2.Z_OBJ_P(arg).GetProperties().GetRefcount() > 1 {
			if (types2.Z_OBJ_P(arg).GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
				types2.Z_OBJ_P(arg).GetProperties().DelRefcount()
			}
			types2.Z_OBJ_P(arg).SetProperties(types2.ZendArrayDup(types2.Z_OBJ_P(arg).GetProperties()))
		}
		return types2.Z_OBJ_HT_P(arg).GetGetProperties()(arg), true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseObject(arg *types2.Zval, ce *types2.ClassEntry, checkNull bool) (dest *types2.Zval, ok bool) {
	if arg.IsObject() && (ce == nil || zend.InstanceofFunction(types2.Z_OBJCE_P(arg), ce) != 0) {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseResource(arg *types2.Zval, checkNull bool) (dest *types2.Zval, ok bool) {
	if arg.IsResource() {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func ParseFunc(arg *types2.Zval, dest_fci *types2.ZendFcallInfo, dest_fcc *types2.ZendFcallInfoCache, checkNull bool) (error *string, ok bool) {
	if checkNull && arg.IsNull() {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		return nil, true
	}

	// notice: 此处在成功时 error 也有可能不为 nil (例如在产生 Deprecated 信息时)
	state := zend.ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, &error)
	return error, state == types2.SUCCESS
}

func ParseZvalDeref(arg *types2.Zval, checkNull bool) (dest *types2.Zval) {
	if checkNull && arg.IsNull() {
		return nil
	} else {
		return arg
	}
}

func ParseClass(arg *types2.Zval, baseCe *types2.ClassEntry, num int, checkNull bool) (ce *types2.ClassEntry, ok bool) {
	if checkNull && arg.IsNull() {
		return nil, true
	}

	if zend.TryConvertToString(arg) == 0 {
		return nil, false
	}
	ce = zend.ZendLookupClass(arg.String())
	if baseCe != nil {
		if ce == nil || zend.InstanceofFunction(ce, baseCe) == 0 {
			faults.InternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a class name derived from %s, '%s' given", zend.GetActiveCalleeName(), num, baseCe.Name(), arg.StringVal())
			return nil, false
		}
	}
	if ce == nil {
		faults.InternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a valid class name, '%s' given", zend.GetActiveCalleeName(), num, arg.String().GetVal())
		return nil, false
	}
	return ce, true
}

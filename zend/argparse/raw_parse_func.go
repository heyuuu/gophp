package argparse

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

func Z_PARAM_PROLOGUE(deref int, separate int) {
	_i++
	b.Assert(_i <= _min_num_args || _optional == 1)
	b.Assert(_i > _min_num_args || _optional == 0)
	if _optional {
		if _i > _num_args {
			break
		}
	}
	_real_arg++
	_arg = _real_arg
	if deref != 0 {
		if _arg.IsReference() {
			_arg = types.Z_REFVAL_P(_arg)
		}
	}
	if separate != 0 {
		types.SEPARATE_ZVAL_NOREF(_arg)
	}
}
func ZendParseArgBool(arg *types.Zval, dest *types.ZendBool, is_null *types.ZendBool, check_null int) int {
	val, isNull, ok := ParseBool(arg, check_null != 0)
	*dest = types.IntBool(val)
	if check_null != 0 {
		*is_null = types.IntBool(isNull)
	}
	return types.IntBool(ok)
}

func ZendParseArgLong00(arg *types.Zval, dest *zend.ZendLong) bool {
	val, _, ok := ParseLong(arg, false, false)
	*dest = val
	return ok
}

func ZendParseArgLong(arg *types.Zval, dest *zend.ZendLong, is_null *types.ZendBool, check_null int, cap int) bool {
	val, isNull, ok := ParseLong(arg, check_null != 0, cap != 0)
	*dest = val
	if is_null != nil {
		*is_null = types.IntBool(isNull)
	}
	return ok
}
func ZendParseArgDouble(arg *types.Zval, dest *float64, is_null *types.ZendBool, check_null int) int {
	val, isNull, ok := ParseDouble(arg, check_null != 0)
	*dest = val
	if is_null != nil {
		*is_null = types.IntBool(isNull)
	}
	return types.IntBool(ok)
}
func ZendParseArgStr(arg *types.Zval, dest **types.ZendString, check_null int) int {
	// 为空时 *dest 直接为 nil，不需单独的 is_null 字符安
	val, ok := ParseZStr(arg, check_null != 0)
	*dest = val
	return types.IntBool(ok)
}
func ZendParseArgString(arg *types.Zval, dest **byte, dest_len *int, check_null int) int {
	var str *types.ZendString
	if ZendParseArgStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgPathStr(arg *types.Zval, dest **types.ZendString, check_null int) int {
	if ZendParseArgStr(arg, dest, check_null) == 0 || (*dest) != nil && zend.CHECK_NULL_PATH(dest.GetVal(), dest.GetLen()) {
		return 0
	}
	return 1
}
func ZendParseArgPath(arg *types.Zval, dest **byte, dest_len *int, check_null int) int {
	var str *types.ZendString
	if ZendParseArgPathStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgArray(arg *types.Zval, dest **types.Zval, check_null int, or_object int) int {
	if arg.IsArray() || or_object != 0 && arg.IsObject() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgArrayHt(arg *types.Zval, dest **types.HashTable, check_null int, or_object int, separate int) int {
	if arg.IsArray() {
		*dest = arg.GetArr()
	} else if or_object != 0 && arg.IsObject() {
		if separate != 0 && types.Z_OBJ_P(arg).GetProperties() != nil && types.Z_OBJ_P(arg).GetProperties().GetRefcount() > 1 {
			if (types.Z_OBJ_P(arg).GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
				types.Z_OBJ_P(arg).GetProperties().DelRefcount()
			}
			types.Z_OBJ_P(arg).SetProperties(zend.ZendArrayDup(types.Z_OBJ_P(arg).GetProperties()))
		}
		*dest = types.Z_OBJ_HT_P(arg).GetGetProperties()(arg)
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgObject(arg *types.Zval, dest **types.Zval, ce *zend.ZendClassEntry, check_null int) int {
	if arg.IsObject() && (ce == nil || zend.InstanceofFunction(types.Z_OBJCE_P(arg), ce) != 0) {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgResource(arg *types.Zval, dest **types.Zval, check_null int) int {
	if arg.IsResource() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgFunc(arg *types.Zval, dest_fci *zend.ZendFcallInfo, dest_fcc *zend.ZendFcallInfoCache, check_null int, error **byte) int {
	if check_null != 0 && arg.IsNull() {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		*error = nil
	} else if zend.ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, error) != types.SUCCESS {
		return 0
	}
	return 1
}
func ZendParseArgZvalDeref(arg *types.Zval, dest **types.Zval, check_null int) {
	if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		*dest = arg
	}
}

func ZendParseArgClass(arg *types.Zval, pce **zend.ZendClassEntry, num int, check_null int) int {
	var ce_base *zend.ZendClassEntry = *pce
	if check_null != 0 && arg.IsNull() {
		*pce = nil
		return 1
	}
	if zend.TryConvertToString(arg) == 0 {
		*pce = nil
		return 0
	}
	*pce = zend.ZendLookupClass(arg.GetStr())
	if ce_base != nil {
		if (*pce) == nil || zend.InstanceofFunction(*pce, ce_base) == 0 {
			zend.ZendInternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a class name derived from %s, '%s' given", zend.GetActiveCalleeName(), num, ce_base.GetName().GetVal(), arg.GetStr().GetVal())
			*pce = nil
			return 0
		}
	}
	if (*pce) == nil {
		zend.ZendInternalTypeError(zend.CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a valid class name, '%s' given", zend.GetActiveCalleeName(), num, arg.GetStr().GetVal())
		return 0
	}
	return 1
}

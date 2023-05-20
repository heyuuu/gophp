package operators

import (
	"fmt"
	"github.com/heyuuu/gophp/core/pfmt"
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"strconv"
)

func IZendIsTrue(op *types.Zval) int { return types.IntBool(ZvalIsTrue(op)) }

func ZvalIsTrue(op *types.Zval) bool {
again:
	switch op.GetType() {
	case types.IS_TRUE:
		return true
	case types.IS_LONG:
		return op.Long() != 0
	case types.IS_DOUBLE:
		return op.Double() != 0
	case types.IS_STRING:
		str := op.StringVal()
		return str != "" && str != "0"
	case types.IS_ARRAY:
		return op.Array().Len() != 0
	case types.IS_OBJECT:
		return ZendObjectIsTrue(op)
	case types.IS_RESOURCE:
		return op.ResourceHandle() != 0
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto again
	}
	return false
}

func ZvalGetLong(op *types.Zval) int {
	if op.IsLong() {
		return op.Long()
	} else {
		return _zvalGetLongFuncEx(op, true)
	}
}

func _zvalGetLongFuncNoisy(op *types.Zval) int { return _zvalGetLongFuncEx(op, false) }
func _zvalGetLongFuncEx(op *types.Zval, silent bool) int {
	op = op.DeRef()
	switch op.GetType() {
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		return 0
	case types.IS_TRUE:
		return 1
	case types.IS_RESOURCE:
		return op.ResourceHandle()
	case types.IS_LONG:
		return op.Long()
	case types.IS_DOUBLE:
		return DvalToLval(op.Double())
	case types.IS_STRING:
		var r conv.ParseNumberResult
		if silent {
			r = zend.StrToNumberAllowErrors(op.StringVal())
		} else {
			r = zend.StrToNumberNoticeErrors(op.StringVal())
		}
		if r.IsInt() {
			return r.Int()
		} else if r.IsFloat() {
			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */
			return DvalToLvalCap(r.Float())
		} else {
			if !silent {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
			return 0
		}
	case types.IS_ARRAY:
		if op.Array().Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_LONG, ConvertToLong)
		if dst.IsLong() {
			return dst.Long()
		} else {
			return 1
		}
	default:
		return 0
	}
}

func ZvalGetDouble(op *types.Zval) float64 {
	if op.IsDouble() {
		return op.Double()
	} else {
		return ZvalGetDoubleFunc(op)
	}
}
func ZvalGetDoubleFunc(op *types.Zval) float64 {
	op = op.DeRef()
	switch op.GetType() {
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		return 0.0
	case types.IS_TRUE:
		return 1.0
	case types.IS_RESOURCE:
		return float64(op.ResourceHandle())
	case types.IS_LONG:
		return float64(op.Long())
	case types.IS_DOUBLE:
		return op.Double()
	case types.IS_STRING:
		return zend.ZendStrtod(op.StringVal(), nil)
	case types.IS_ARRAY:
		if op.Array().Len() != 0 {
			return 1.0
		} else {
			return 0.0
		}
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_DOUBLE, ConvertToDouble)
		if dst.IsDouble() {
			return dst.Double()
		} else {
			return 1.0
		}
	default:
		return 0.0
	}
}

func ZvalGetStrVal(op *types.Zval) string {
	str, _ := ZvalGetStr(op)
	return str
}

func ZvalGetStr(op *types.Zval) (string, bool) {
	return __zvalGetStrFunc(op, false)
}
func ZvalGetString(op *types.Zval) *types.String {
	if str, ok := ZvalGetStr(op); ok {
		return types.NewString(str)
	}
	return nil
}

func ZvalTryGetStr(op *types.Zval) (string, bool) {
	return __zvalGetStrFunc(op, true)
}
func ZvalTryGetString(op *types.Zval) *types.String {
	if str, ok := ZvalTryGetStr(op); ok {
		return types.NewString(str)
	}
	return nil
}

/**
 * 从 Zval 转字符串
 * @return string 返回的字符串值。
 * @return bool   是否成功。
 */
func __zvalGetStrFunc(op *types.Zval, try bool) (string, bool) {
	op = op.DeRef()
	switch op.GetType() {
	case types.IS_STRING:
		return op.StringVal(), true
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		return "", true
	case types.IS_TRUE:
		return "1", true
	case types.IS_RESOURCE:
		return fmt.Sprintf("Resource id #%d", op.ResourceHandle()), true
	case types.IS_LONG:
		return strconv.Itoa(op.Long()), true
	case types.IS_DOUBLE:
		return pfmt.Sprintf("%.*G", zend.EG__().GetPrecision(), op.Double()), true
	case types.IS_ARRAY:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try && zend.EG__().GetException() != nil {
			return "", false
		}
		return types.STR_ARRAY_CAPITALIZED, true
	case types.IS_OBJECT:
		var tmp types.Zval
		if op.Object().CanCast() {
			if op.Object().Cast(&tmp, types.IS_STRING) == types.SUCCESS {
				return tmp.StringVal(), true
			}
		} else if op.Object().CanGet() {
			var z *types.Zval = op.Object().Get(&tmp)
			if z.GetType() != types.IS_OBJECT {
				return __zvalGetStrFunc(z, try)
			}
		}
		if zend.EG__().GetException() == nil {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).GetName().GetVal())
		}
		if try {
			return "", false
		} else {
			return "", true
		}
	default:
		return "", false
	}
}

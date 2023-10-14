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
	switch op.Type() {
	case types.IsTrue:
		return true
	case types.IsLong:
		return op.Long() != 0
	case types.IsDouble:
		return op.Double() != 0
	case types.IsString:
		str := op.String()
		return str != "" && str != "0"
	case types.IsArray:
		return op.Array().Len() != 0
	case types.IsObject:
		return ZendObjectIsTrue(op)
	case types.IsResource:
		return op.ResourceHandle() != 0
	case types.IsRef:
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
	switch op.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return 0
	case types.IsTrue:
		return 1
	case types.IsResource:
		return op.ResourceHandle()
	case types.IsLong:
		return op.Long()
	case types.IsDouble:
		return DvalToLval(op.Double())
	case types.IsString:
		var r conv.ParseNumberResult
		if silent {
			r = zend.StrToNumberAllowErrors(op.String())
		} else {
			r = zend.StrToNumberNoticeErrors(op.String())
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
	case types.IsArray:
		if op.Array().Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsLong, ConvertToLong)
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
	switch op.Type() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		return 0.0
	case types.IsTrue:
		return 1.0
	case types.IsResource:
		return float64(op.ResourceHandle())
	case types.IsLong:
		return float64(op.Long())
	case types.IsDouble:
		return op.Double()
	case types.IsString:
		return zend.ZendStrtod(op.String(), nil)
	case types.IsArray:
		if op.Array().Len() != 0 {
			return 1.0
		} else {
			return 0.0
		}
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsDouble, ConvertToDouble)
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
	switch op.Type() {
	case types.IsString:
		return op.String(), true
	case types.IsUndef, types.IsNull, types.IsFalse:
		return "", true
	case types.IsTrue:
		return "1", true
	case types.IsResource:
		return fmt.Sprintf("Resource id #%d", op.ResourceHandle()), true
	case types.IsLong:
		return strconv.Itoa(op.Long()), true
	case types.IsDouble:
		return pfmt.Sprintf("%.*G", zend.EG__().GetPrecision(), op.Double()), true
	case types.IsArray:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try && zend.EG__().HasException() {
			return "", false
		}
		return "Array", true
	case types.IsObject:
		var tmp types.Zval
		if op.Object().CanCast() {
			if op.Object().Cast(&tmp, types.IsString) == types.SUCCESS {
				return tmp.String(), true
			}
		} else if op.Object().CanGet() {
			var z *types.Zval = op.Object().Get(&tmp)
			if !z.IsObject() {
				return __zvalGetStrFunc(z, try)
			}
		}
		if zend.EG__().NoException() {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).Name())
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

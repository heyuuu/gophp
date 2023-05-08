package zend

import (
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
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
		if types.Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			return true
		} else {
			return ZendObjectIsTrue(op)
		}
	case types.IS_RESOURCE:
		return types.Z_RES_HANDLE_P(op) != 0
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto again
	}
	return false
}

func ZvalGetLong(op *types.Zval) ZendLong {
	if op.IsLong() {
		return op.Long()
	} else {
		return _zvalGetLongFuncEx(op, true)
	}
}

func _zvalGetLongFuncNoisy(op *types.Zval) ZendLong { return _zvalGetLongFuncEx(op, false) }
func _zvalGetLongFuncEx(op *types.Zval, silent bool) ZendLong {
	op = op.DeRef()
	switch op.GetType() {
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		return 0
	case types.IS_TRUE:
		return 1
	case types.IS_RESOURCE:
		return types.Z_RES_HANDLE_P(op)
	case types.IS_LONG:
		return op.Long()
	case types.IS_DOUBLE:
		return DvalToLval(op.Double())
	case types.IS_STRING:
		var r conv.ParseNumberResult
		if silent {
			r = StrToNumberAllowErrors(op.StringVal())
		} else {
			r = StrToNumberNoticeErrors(op.StringVal())
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
		return float64(types.Z_RES_HANDLE_P(op))
	case types.IS_LONG:
		return float64(op.Long())
	case types.IS_DOUBLE:
		return op.Double()
	case types.IS_STRING:
		return ZendStrtod(op.StringVal(), nil)
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
	if op.IsString() {
		return op.StringVal()
	} else {
		zstr := ZvalGetStringFunc(op)
		if zstr == nil {
			return ""
		}
		return zstr.GetStr()
	}
}

func ZvalGetString(op *types.Zval) *types.String {
	if op.IsString() {
		return op.String().Copy()
	} else {
		return ZvalGetStringFunc(op)
	}
}

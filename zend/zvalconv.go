package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
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
		return op.ResourceHandle() != 0
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
		return op.ResourceHandle()
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
		return float64(op.ResourceHandle())
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
	str, _ := ZvalGetStrValEx(op)
	return str
}
func ZvalGetStrValEx(op *types.Zval) (string, bool) {
	if op.IsString() {
		return op.StringVal(), true
	} else {
		zstr := ZvalGetStringFunc(op)
		if zstr == nil {
			return "", false
		}
		return zstr.GetStr(), true
	}
}
func ZvalGetString(op *types.Zval) *types.String {
	if str, ok := ZvalGetStrValEx(op); ok {
		return types.NewString(str)
	}
	return nil
}

func ZvalGetTmpString(op *types.Zval, tmp **types.String) *types.String {
	if op.IsString() {
		*tmp = nil
		return op.String()
	} else {
		*tmp = ZvalGetStringFunc(op)
		return *tmp
	}
}
func ZvalTryGetString(op *types.Zval) *types.String {
	if op.IsString() {
		var ret *types.String = op.String().Copy()
		return ret
	} else {
		return ZvalTryGetStringFunc(op)
	}
}
func ZvalTryGetTmpString(op *types.Zval, tmp **types.String) *types.String {
	if op.IsString() {
		*tmp = nil
		return op.String()
	} else {
		*tmp = ZvalTryGetStringFunc(op)
		return *tmp
	}
}

func __zvalGetStrFunc(op *types.Zval, try types.ZendBool) (string, bool) {
try_again:
	switch op.GetType() {
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		return "", true
	case types.IS_TRUE:
		return "1", true
	case types.IS_RESOURCE:
		str := ZendSprintf("Resource id #"+ZEND_LONG_FMT, op.ResourceHandle())
		return fmt.Sprintf("Resource id #%d", op.ResourceHandle()), true
	case types.IS_LONG:
		return ZendLongToStr(op.Long())
	case types.IS_DOUBLE:
		str := ZendSprintf("%.*G", EG__().GetPrecision(), op.Double())
		return types.NewString(str)
	case types.IS_ARRAY:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try != 0 && EG__().GetException() != nil {
			return nil
		} else {
			return types.NewString(types.STR_ARRAY_CAPITALIZED)
		}
	case types.IS_OBJECT:
		var tmp types.Zval
		if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
			if types.Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, types.IS_STRING) == types.SUCCESS {
				return tmp.String()
			}
		} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
			var z *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != types.IS_OBJECT {
				var str *types.String = b.CondF(try != 0, func() *types.String { return ZvalTryGetString(z) }, func() *types.String { return ZvalGetString(z) })
				return str
			}
		}
		if EG__().GetException() == nil {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return types.NewString("")
		}
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto try_again
	case types.IS_STRING:
		return op.String().Copy()
	default:
	}
	return nil
}

func __zvalGetStringFunc(op *types.Zval, try types.ZendBool) *types.String {
try_again:
	switch op.GetType() {
	case types.IS_UNDEF, types.IS_NULL, types.IS_FALSE:
		return types.NewString("")
	case types.IS_TRUE:
		return types.NewString("1")
	case types.IS_RESOURCE:
		str := ZendSprintf("Resource id #"+ZEND_LONG_FMT, op.ResourceHandle())
		return types.NewString(str)
	case types.IS_LONG:
		return ZendLongToStr(op.Long())
	case types.IS_DOUBLE:
		str := ZendSprintf("%.*G", EG__().GetPrecision(), op.Double())
		return types.NewString(str)
	case types.IS_ARRAY:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try != 0 && EG__().GetException() != nil {
			return nil
		} else {
			return types.NewString(types.STR_ARRAY_CAPITALIZED)
		}
	case types.IS_OBJECT:
		var tmp types.Zval
		if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
			if types.Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, types.IS_STRING) == types.SUCCESS {
				return tmp.String()
			}
		} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
			var z *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != types.IS_OBJECT {
				var str *types.String = b.CondF(try != 0, func() *types.String { return ZvalTryGetString(z) }, func() *types.String { return ZvalGetString(z) })
				return str
			}
		}
		if EG__().GetException() == nil {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return types.NewString("")
		}
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto try_again
	case types.IS_STRING:
		return op.String().Copy()
	default:
	}
	return nil
}
func ZvalGetStringFunc(op *types.Zval) *types.String    { return __zvalGetStringFunc(op, 0) }
func ZvalTryGetStringFunc(op *types.Zval) *types.String { return __zvalGetStringFunc(op, 1) }

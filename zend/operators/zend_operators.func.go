package operators

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/conv"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"math"
	"strings"
)

func IsFinite(f float64) bool {
	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return false
	}
	return true
}

func DoubleFitsLong(d float64) bool {
	return !(d >= zend.ZEND_LONG_MAX || d < zend.ZEND_LONG_MIN)
}
func DvalToLval(d float64) zend.ZendLong {
	if !IsFinite(d) {
		return 0
	} else if !(DoubleFitsLong(d)) {
		return ZendDvalToLvalSlow(d)
	}
	return zend.ZendLong(d)
}
func DvalToLvalCap(d float64) zend.ZendLong {
	if !IsFinite(d) {
		return 0
	} else if !(DoubleFitsLong(d)) {
		if d > 0 {
			return zend.ZEND_LONG_MAX
		} else {
			return zend.ZEND_LONG_MIN
		}
	}
	return zend.ZendLong(d)
}
func IsNumericString(str string, lval *zend.ZendLong, dval *float64, allow_errors zend.ConvertNumericMode) uint8 {
	r := zend.StrToNumberEx(str, allow_errors)

	*lval = r.Int()
	if dval != nil {
		*dval = r.Float()
	}
	if r.IsInt() {
		return types.IsLong
	} else if r.IsFloat() {
		return types.IsDouble
	} else {
		return 0
	}
}
func ZendMemnstr(haystack *byte, needle string, needle_len int, end *byte) *byte {
	// todo 替换 - 查找haystack中needle首次出现的位置，没出现则返回nil
	pos := strings.Index(b.CastStr(haystack, end-haystack), b.CastStr(needle))
	if pos < 0 {
		return nil
	}
	return haystack + pos
}
func ZendMemrchr(s *byte, c byte, n int) *byte {
	str := b.CastStr(s, n)

	if pos := strings.LastIndexByte(str, c); pos >= 0 {
		return s + pos
	} else {
		return nil
	}
}
func TryConvertToString(op *types.Zval) bool {
	if op.IsString() {
		return true
	} else if str, ok := ZvalTryGetStr(op); ok {
		op.SetStringVal(str)
		return true
	} else {
		return false
	}
}
func ConvertToString(op *types.Zval) {
	TryConvertToString(op)
}
func ZendStringTolower(str *types.String) *types.String {
	return types.NewString(ascii.StrToLower(str.GetStr()))
}
func ConvertToStringEx(pzv *types.Zval) {
	if !pzv.IsString() {
		ConvertToString(pzv)
	}
}
func ConvertToArrayEx(pzv *types.Zval) {
	if !pzv.IsArray() {
		ConvertToArray(pzv)
	}
}
func ConvertScalarToNumberEx(pzv *types.Zval) {
	if !pzv.IsLong() && !pzv.IsDouble() {
		ConvertScalarToNumber(pzv)
	}
}
func FastLongIncrementFunction(op1 *types.Zval) {
	if op1.Long() == zend.ZEND_LONG_MAX {
		/* switch to double */
		op1.SetDouble(zend.ZEND_LONG_MAX + 1.0)
	} else {
		op1.SetLong(op1.Long() + 1)
	}
}
func FastLongDecrementFunction(op1 *types.Zval) {
	if op1.Long() == zend.ZEND_LONG_MIN {
		/* switch to double */
		op1.SetDouble(zend.ZEND_LONG_MIN - 1.0)
	} else {
		op1.SetLong(op1.Long() - 1)
	}
}
func FastLongAddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */
	l1, l2 := op1.Long(), op2.Long()
	if l1&LONG_SIGN_MASK == l2&LONG_SIGN_MASK && l1&LONG_SIGN_MASK != l1+l2&LONG_SIGN_MASK {
		result.SetDouble(float64(l1) + float64(l2))
	} else {
		result.SetLong(l1 + l2)
	}
}
func FastAddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if op1.IsLong() {
		if op2.IsLong() {
			FastLongAddFunction(result, op1, op2)
			return types.SUCCESS
		} else if op2.IsDouble() {
			result.SetDouble(float64(op1.Long()) + op2.Double())
			return types.SUCCESS
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			result.SetDouble(op1.Double() + op2.Double())
			return types.SUCCESS
		} else if op2.IsLong() {
			result.SetDouble(op1.Double() + float64(op2.Long()))
			return types.SUCCESS
		}
	}
	return AddFunction(result, op1, op2)
}
func FastLongSubFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	result.SetLong(op1.Long() - op2.Long())
	if (op1.Long()&LONG_SIGN_MASK) != (op2.Long()&LONG_SIGN_MASK) && (op1.Long()&LONG_SIGN_MASK) != (result.Long()&LONG_SIGN_MASK) {
		result.SetDouble(float64(op1.Long()) - float64(op2.Long()))
	}
}
func FastDivFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	return DivFunction(result, op1, op2)
}
func ZendFastEqualStringsEx(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) > 0 && s1[0] <= '9' && len(s2) > 0 && s2[0] <= '9' {
		return ZendiSmartStreq(s1, s2)
	} else {
		return false
	}
}
func FastEqualCheckFunction(op1 *types.Zval, op2 *types.Zval) bool {
	var result types.Zval
	if op1.IsLong() {
		if op2.IsLong() {
			return op1.Long() == op2.Long()
		} else if op2.IsDouble() {
			return float64(op1.Long()) == op2.Double()
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			return op1.Double() == op2.Double()
		} else if op2.IsLong() {
			return op1.Double() == float64(op2.Long())
		}
	} else if op1.IsString() {
		if op2.IsString() {
			return ZendFastEqualStringsEx(op1.StringVal(), op2.StringVal())
		}
	}
	CompareFunction(&result, op1, op2)
	return result.Long() == 0
}
func FastEqualCheckLong(op1 *types.Zval, op2 *types.Zval) bool {
	b.Assert(op1.IsLong()) // @add
	var result types.Zval
	if op2.IsLong() {
		return op1.Long() == op2.Long()
	}
	CompareFunction(&result, op1, op2)
	return result.Long() == 0
}
func FastEqualCheckString(op1 *types.Zval, op2 *types.Zval) bool {
	var result types.Zval
	if op2.IsString() {
		return ZendFastEqualStringsEx(op1.StringVal(), op2.StringVal())
	}
	CompareFunction(&result, op1, op2)
	return result.Long() == 0
}
func FastIsIdenticalFunction(op1 *types.Zval, op2 *types.Zval) bool {
	if op1.GetType() != op2.GetType() {
		return false
	} else if op1.IsSignType() {
		return true
	}
	return ZendIsIdentical(op1, op2)
}
func FastIsNotIdenticalFunction(op1 *types.Zval, op2 *types.Zval) bool {
	if op1.GetType() != op2.GetType() {
		return true
	} else if op1.IsSignType() {
		return false
	}
	return !ZendIsIdentical(op1, op2)
}
func ZendUnwrapReference(op *types.Zval) {
	types.ZVAL_COPY(op, types.Z_REFVAL_P(op))
}
func ConvertObjectToType(op *types.Zval, dst *types.Zval, ctype types.ZvalType, convFunc func(op *types.Zval)) {
	dst.SetUndef()
	if op.Object().CanCast() {
		if op.Object().Cast(dst, ctype) == types.FAILURE {
			faults.Error(faults.E_RECOVERABLE_ERROR, "Object of class %s could not be converted to %s", types.Z_OBJCE_P(op).Name(), types.ZendGetTypeByConst(ctype))
		}
	} else if op.Object().CanGet() {
		var newop *types.Zval = op.Object().Get(dst)
		if !newop.IsObject() {
			dst.CopyValueFrom(newop)
			convFunc(dst)
		}
	}
}
func _convertScalarToNumber(op *types.Zval, silent bool, check bool) {
	op = op.DeRef()
	switch op.GetType() {
	case types.IsString:
		var r conv.ParseNumberResult
		if silent {
			r = zend.StrToNumberAllowErrors(op.StringVal())
		} else {
			r = zend.StrToNumberNoticeErrors(op.StringVal())
		}
		if r.IsInt() {
			op.SetLong(r.Int())
		} else if r.IsFloat() {
			op.SetDouble(r.Float())
		} else { // fail
			op.SetLong(0)
			if !silent {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
		}
	case types.IsNull,
		types.IsFalse:
		op.SetLong(0)
	case types.IsTrue:
		op.SetLong(1)
	case types.IsResource:
		var l zend.ZendLong = op.ResourceHandle()
		op.SetLong(l)
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsNumber, ConvertScalarToNumber)
		if check && zend.EG__().GetException() != nil {
			return
		}
		if dst.IsLong() || dst.IsDouble() {
			types.ZVAL_COPY_VALUE(op, &dst)
		} else {
			op.SetLong(1)
		}
	}
}
func ConvertScalarToNumber(op *types.Zval) { _convertScalarToNumber(op, true, false) }
func _zendiConvertScalarToNumberEx(op *types.Zval, holder *types.Zval, silent bool) *types.Zval {
	switch op.GetType() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		holder.SetLong(0)
		return holder
	case types.IsTrue:
		holder.SetLong(1)
		return holder
	case types.IsString:
		var mode zend.ConvertNumericMode
		if silent {
			mode = zend.ConvertContinueOnErrors
		} else {
			mode = zend.ConvertNoticeOnErrors
		}
		r := zend.StrToNumberEx(op.StringVal(), mode)
		if r.IsInt() {
			holder.SetLong(r.Int())
		} else if r.IsFloat() {
			holder.SetDouble(r.Float())
		} else {
			holder.SetLong(0)
			if !silent {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
		}
		return holder
	case types.IsResource:
		holder.SetLong(op.ResourceHandle())
		return holder
	case types.IsObject:
		ConvertObjectToType(op, holder, types.IsNumber, ConvertScalarToNumber)
		if zend.EG__().GetException() != nil || !holder.IsLong() && !holder.IsDouble() {
			holder.SetLong(1)
		}
		return holder
	case types.IsLong:
		fallthrough
	case types.IsDouble:
		fallthrough
	default:
		return op
	}
}
func _zendiConvertScalarToNumber(op *types.Zval, holder *types.Zval) *types.Zval {
	return _zendiConvertScalarToNumberEx(op, holder, true)
}
func _zendiConvertScalarToNumberNoisy(op *types.Zval, holder *types.Zval) *types.Zval {
	return _zendiConvertScalarToNumberEx(op, holder, false)
}
func ZendiConvertScalarToNumber(op *types.Zval, holder *types.Zval, result *types.Zval, silent bool) *types.Zval {
	if op.IsLong() || op.IsDouble() {
		return op
	} else {
		if op == result {
			_convertScalarToNumber(op, silent, true)
			return op
		} else {
			if silent {
				return _zendiConvertScalarToNumber(op, holder)
			} else {
				return _zendiConvertScalarToNumberNoisy(op, holder)
			}
		}
	}
}
func ConvertToLong(op *types.Zval) {
	if !op.IsLong() {
		ConvertToLongBase(op, 10)
	}
}
func ConvertToLongBase(op *types.Zval, base int) {
	var tmp zend.ZendLong
try_again:
	switch op.GetType() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		op.SetLong(0)
	case types.IsTrue:
		op.SetLong(1)
	case types.IsResource:
		tmp = op.ResourceHandle()
		// ZvalPtrDtor(op)
		op.SetLong(tmp)
	case types.IsLong:

	case types.IsDouble:
		op.SetLong(DvalToLval(op.Double()))
	case types.IsString:
		var str *types.String = op.String()
		if base == 10 {
			op.SetLong(ZvalGetLong(op))
		} else {
			op.SetLong(zend.ZEND_STRTOL(str.GetVal(), nil, base))
		}
	case types.IsArray:
		if op.Array().Len() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		op.SetLong(tmp)
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsLong, ConvertToLong)
		if dst.IsLong() {
			op.SetLong(dst.Long())
		} else {
			op.SetLong(1)
		}
		return
	case types.IsRef:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToDouble(op *types.Zval) {
	var tmp float64
try_again:
	switch op.GetType() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		op.SetDouble(0.0)
	case types.IsTrue:
		op.SetDouble(1.0)
	case types.IsResource:
		var d float64 = float64(op.ResourceHandle())
		// ZvalPtrDtor(op)
		op.SetDouble(d)
	case types.IsLong:
		op.SetDouble(float64(op.Long()))
	case types.IsDouble:

	case types.IsString:
		var str *types.String = op.String()
		op.SetDouble(zend.ZendStrtod(str.GetStr(), nil))
	case types.IsArray:
		if op.Array().Len() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		// ZvalPtrDtor(op)
		op.SetDouble(tmp)
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsDouble, ConvertToDouble)
		// ZvalPtrDtor(op)
		if dst.IsDouble() {
			op.SetDouble(dst.Double())
		} else {
			op.SetDouble(1.0)
		}
	case types.IsRef:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToNull(op *types.Zval) {
	op.SetNull()
}
func ConvertToBoolean(op *types.Zval) {
	var tmp int
try_again:
	switch op.GetType() {
	case types.IsFalse:
		fallthrough
	case types.IsTrue:

	case types.IsNull:
		op.SetFalse()
	case types.IsResource:
		var l zend.ZendLong = lang.Cond(op.ResourceHandle() != 0, 1, 0)
		// ZvalPtrDtor(op)
		op.SetBool(l != 0)
	case types.IsLong:
		op.SetBool(op.Long() != 0)
	case types.IsDouble:
		op.SetBool(op.Double() != 0)
	case types.IsString:
		var str *types.String = op.String()
		if str.GetLen() == 0 || str.GetLen() == 1 && str.GetStr()[0] == '0' {
			op.SetFalse()
		} else {
			op.SetTrue()
		}
		// types.ZendStringReleaseEx(str, 0)
	case types.IsArray:
		if op.Array().Len() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		// ZvalPtrDtor(op)
		op.SetBool(tmp != 0)
	case types.IsObject:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IsBool, ConvertToBoolean)
		if dst.IsFalse() {
			op.SetFalse()
		} else { // dst.IsTrue() or others
			op.SetTrue()
		}
	case types.IsRef:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}

func ConvertScalarToArray(op *types.Zval) {
	var ht *types.Array = types.NewArray(1)
	ht.IndexAddNew(0, op)
	op.SetArray(ht)
}
func ConvertToArray(op *types.Zval) {
try_again:
	switch op.GetType() {
	case types.IsRef:
		ZendUnwrapReference(op)
		goto try_again
	case types.IsArray:
		// pass
	case types.IsObject:
		if types.Z_OBJCE_P(op) == zend.ZendCeClosure {
			ConvertScalarToArray(op)
		} else {
			var objHt *types.Array = zend.ZendGetPropertiesFor(op, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
			if objHt != nil {
				var newObjHt *types.Array = types.ZendProptableToSymtable(objHt, types.Z_OBJCE_P(op).GetDefaultPropertiesCount() != 0 || op.Object().GetHandlers() != zend.StdObjectHandlersPtr || objHt.IsRecursive())
				op.SetArray(newObjHt)
			} else {
				zend.ArrayInit(op)
			}
		}
	case types.IsNull:
		/*ZVAL_EMPTY_ARRAY(op);*/
		zend.ArrayInit(op)
	default:
		ConvertScalarToArray(op)
	}
}
func ConvertToObject(op *types.Zval) {
try_again:
	switch op.GetType() {
	case types.IsRef:
		ZendUnwrapReference(op)
		goto try_again
	case types.IsArray:
		var ht = types.ZendSymtableToProptable(op.Array())
		//if ht.IsImmutable() {
		//	/* TODO: try not to duplicate immutable arrays as well ??? */
		//	ht = types.ZendArrayDup(ht)
		//} else if ht != op.Array() {
		//	// ZvalPtrDtor(op)
		//} else {
		//	//ht.DelRefcount()
		//}
		op.SetObject(zend.NewStdClassObject(ht))
	case types.IsObject:

	case types.IsNull:
		zend.ObjectInit(op)
	default:
		var tmp types.Zval
		types.ZVAL_COPY_VALUE(&tmp, op)
		zend.ObjectInit(op)
		types.Z_OBJPROP_P(op).KeyAddNew(types.STR_SCALAR, &tmp)
	}
}
func AddFunctionArray(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	if result == op1 && op1.Array() == op2.Array() {
		/* $a += $a */
		return
	}
	if result != op1 {
		result.SetArray(types.ZendArrayDup(op1.Array()))
	} else {
		types.SeparateArray(result)
	}
	types.ZendHashMerge(result.Array(), op2.Array(), false)
}
func _zvalFastGetDouble(op *types.Zval) float64 {
	if op.IsLong() {
		return float64(op.Long())
	} else if op.IsDouble() {
		return op.Double()
	} else {
		return 0
	}
}

func MulFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair = TypePair(op1.GetType(), op2.GetType())
		if type_pair == TypePair(types.IsLong, types.IsLong) {
			if iVal, dVal, overflow := zend.SignedMultiplyLong(op1.Long(), op2.Long()); overflow {
				result.SetDouble(dVal)
			} else {
				result.SetLong(iVal)
			}
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsDouble) {
			result.SetDouble(op1.Double() * op2.Double())
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsLong, types.IsDouble) {
			result.SetDouble(float64(op1.Long()) * op2.Double())
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsLong) {
			result.SetDouble(op1.Double() * float64(op2.Long()))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
					var ret int
					var rv types.Zval
					var objval *types.Zval = op1.Object().Get(&rv)
					// objval.TryAddRefcount()
					ret = MulFunction(objval, objval, op2)
					op1.Object().Set(objval)
					// ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && op1.Object().CanDoOperation() {
					if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_MUL, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_MUL, result, op1, op2) {
					return types.SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, false)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, false)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, false)
					op2 = op1
				}
				if zend.EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func PowFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair = TypePair(op1.GetType(), op2.GetType())
		if type_pair == TypePair(types.IsLong, types.IsLong) {
			if op2.Long() >= 0 {
				var l1 zend.ZendLong = 1
				var l2 zend.ZendLong = op1.Long()
				var i zend.ZendLong = op2.Long()
				if i == 0 {
					result.SetLong(1)
					return types.SUCCESS
				} else if l2 == 0 {
					result.SetLong(0)
					return types.SUCCESS
				}
				for i >= 1 {
					if i%2 != 0 {
						i--
						if iVal, dVal, overflow := zend.SignedMultiplyLong(l1, l2); overflow {
							result.SetDouble(dVal * math.Pow(float64(l2), float64(i)))
							return types.SUCCESS
						} else {
							l1 = iVal
						}
					} else {
						i /= 2
						if iVal, dVal, overflow := zend.SignedMultiplyLong(l1, l2); overflow {
							result.SetDouble(float64(l1) * math.Pow(dVal, float64(i)))
							return types.SUCCESS
						} else {
							l2 = iVal
						}
					}
				}

				/* i == 0 */

				result.SetLong(l1)

				/* i == 0 */

			} else {
				result.SetDouble(math.Pow(float64(op1.Long()), float64(op2.Long())))
			}
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsDouble) {
			result.SetDouble(math.Pow(op1.Double(), op2.Double()))
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsLong, types.IsDouble) {
			result.SetDouble(math.Pow(float64(op1.Long()), op2.Double()))
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsLong) {
			result.SetDouble(math.Pow(op1.Double(), float64(op2.Long())))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
					var ret int
					var rv types.Zval
					var objval *types.Zval = op1.Object().Get(&rv)
					// objval.TryAddRefcount()
					ret = PowFunction(objval, objval, op2)
					op1.Object().Set(objval)
					// ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && op1.Object().CanDoOperation() {
					if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_POW, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_POW, result, op1, op2) {
					return types.SUCCESS
				}
				if op1 != op2 {
					if op1.IsArray() {
						if op1 == result {
							// ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return types.SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, false)
					}
					if op2.IsArray() {
						if op1 == result {
							// ZvalPtrDtor(result)
						}
						result.SetLong(1)
						return types.SUCCESS
					} else {
						op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, false)
					}
				} else {
					if op1.IsArray() {
						if op1 == result {
							// ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return types.SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, false)
					}
					op2 = op1
				}
				if zend.EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func DivFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair uint = TypePair(op1.GetType(), op2.GetType())
		if type_pair == TypePair(types.IsLong, types.IsLong) {
			if op2.Long() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
				result.SetDouble(float64(op1.Long() / float64(op2.Long())))
				return types.SUCCESS
			} else if op2.Long() == -1 && op1.Long() == zend.ZEND_LONG_MIN {

				/* Prevent overflow error/crash */

				result.SetDouble(float64(zend.ZEND_LONG_MIN / -1))
				return types.SUCCESS
			}
			if op1.Long()%op2.Long() == 0 {
				result.SetLong(op1.Long() / op2.Long())
			} else {
				result.SetDouble(float64(op1.Long()) / op2.Long())
			}
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsDouble) {
			if op2.Double() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.Double() / op2.Double())
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsDouble, types.IsLong) {
			if op2.Long() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.Double() / float64(op2.Long()))
			return types.SUCCESS
		} else if type_pair == TypePair(types.IsLong, types.IsDouble) {
			if op2.Double() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(float64(op1.Long() / op2.Double()))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
					var ret int
					var rv types.Zval
					var objval *types.Zval = op1.Object().Get(&rv)
					// objval.TryAddRefcount()
					ret = DivFunction(objval, objval, op2)
					op1.Object().Set(objval)
					// ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && op1.Object().CanDoOperation() {
					if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_DIV, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_DIV, result, op1, op2) {
					return types.SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = op1
				}
				if zend.EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func ModFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	for {
		if !op1.IsLong() {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.Long()
					break
				}
			}
			if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
				var ret int
				var rv types.Zval
				var objval *types.Zval = op1.Object().Get(&rv)
				// objval.TryAddRefcount()
				ret = ModFunction(objval, objval, op2)
				op1.Object().Set(objval)
				// ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && op1.Object().CanDoOperation() {
				if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_MOD, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.Long()
		}
		break
	}
	for {
		if !op2.IsLong() {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.Long()
					break
				}
			}
			if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_MOD, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op2_lval = op2.Long()
		}
		break
	}
	if op2_lval == 0 {

		/* modulus by zero */

		if zend.CurrEX() != nil && zend.CG__().GetInCompilation() == 0 {
			faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Modulo by zero")
		} else {
			faults.ErrorNoreturn(faults.E_ERROR, "Modulo by zero")
		}
		if op1 != result {
			result.SetUndef()
		}
		return types.FAILURE
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}
	if op2_lval == -1 {

		/* Prevent overflow error/crash if op1==LONG_MIN */

		result.SetLong(0)
		return types.SUCCESS
	}
	result.SetLong(op1_lval % op2_lval)
	return types.SUCCESS
}
func BooleanXorFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_val int
	var op2_val int
	for {
		if op1.IsFalse() {
			op1_val = 0
		} else if op1.IsTrue() {
			op1_val = 1
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsFalse() {
					op1_val = 0
					break
				} else if op1.IsTrue() {
					op1_val = 1
					break
				}
			}
			if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
				var ret int
				var rv types.Zval
				var objval *types.Zval = op1.Object().Get(&rv)
				// objval.TryAddRefcount()
				ret = BooleanXorFunction(objval, objval, op2)
				op1.Object().Set(objval)
				// ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && op1.Object().CanDoOperation() {
				if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BOOL_XOR, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_val = IZendIsTrue(op1)
		}
		break
	}
	for {
		if op2.IsFalse() {
			op2_val = 0
		} else if op2.IsTrue() {
			op2_val = 1
		} else {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsFalse() {
					op2_val = 0
					break
				} else if op2.IsTrue() {
					op2_val = 1
					break
				}
			}
			if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_BOOL_XOR, result, op1, op2) {
				return types.SUCCESS
			}
			op2_val = IZendIsTrue(op2)
		}
		break
	}
	result.SetBool((op1_val ^ op2_val) != 0)
	return types.SUCCESS
}
func BooleanNotFunction(result *types.Zval, op1 *types.Zval) int {
	if op1.GetType() < types.IsTrue {
		result.SetTrue()
	} else if op1.IsTrue() {
		result.SetFalse()
	} else {
		if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
			if op1.GetType() < types.IsTrue {
				result.SetTrue()
				return types.SUCCESS
			} else if op1.IsTrue() {
				result.SetFalse()
				return types.SUCCESS
			}
		}
		if op1.IsObject() && op1.Object().CanDoOperation() && types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BOOL_NOT, result, op1, nil) {
			return types.SUCCESS
		}
		result.SetBool(!ZvalIsTrue(op1))
	}
	return types.SUCCESS
}
func BitwiseNotFunction(result *types.Zval, op1 *types.Zval) int {
try_again:
	switch op1.GetType() {
	case types.IsLong:
		result.SetLong(^(op1.Long()))
		return types.SUCCESS
	case types.IsDouble:
		result.SetLong(^(DvalToLval(op1.Double())))
		return types.SUCCESS
	case types.IsString:
		str := []byte(op1.StringVal())
		for i, c := range str {
			str[i] = ^c
		}
		result.SetStringVal(string(str))
		return types.SUCCESS
	case types.IsRef:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		if op1.IsObject() && op1.Object().CanDoOperation() && types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BW_NOT, result, op1, nil) {
			return types.SUCCESS
		}
		if result != op1 {
			result.SetUndef()
		}
		faults.ThrowError(nil, "Unsupported operand types")
		return types.FAILURE
	}
}
func BitwiseOrFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.Long() | op2.Long())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.StringVal(), op2.StringVal()
		str := make([]byte, b.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] | s2[i]
		}
		result.SetStringVal(string(str))
		return types.SUCCESS
	}
	if !op1.IsLong() {
		if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
			var ret int
			var rv types.Zval
			var objval *types.Zval = op1.Object().Get(&rv)
			// objval.TryAddRefcount()
			ret = BitwiseOrFunction(objval, objval, op2)
			op1.Object().Set(objval)
			// ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && op1.Object().CanDoOperation() {
			if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BW_OR, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.Long()
	}
	if !op2.IsLong() {
		if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_BW_OR, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.Long()
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval | op2_lval)
	return types.SUCCESS
}
func BitwiseAndFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.Long() & op2.Long())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.StringVal(), op2.StringVal()
		str := make([]byte, b.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		result.SetStringVal(string(str))
		return types.SUCCESS
	}
	if !op1.IsLong() {
		if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
			var ret int
			var rv types.Zval
			var objval *types.Zval = op1.Object().Get(&rv)
			// objval.TryAddRefcount()
			ret = BitwiseAndFunction(objval, objval, op2)
			op1.Object().Set(objval)
			// ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && op1.Object().CanDoOperation() {
			if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BW_AND, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.Long()
	}
	if !op2.IsLong() {
		if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_BW_AND, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.Long()
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval & op2_lval)
	return types.SUCCESS
}
func BitwiseXorFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.Long() ^ op2.Long())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.StringVal(), op2.StringVal()
		str := make([]byte, b.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] ^ s2[i]
		}
		result.SetStringVal(string(str))
		return types.SUCCESS
	}
	if !op1.IsLong() {
		if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
			var ret int
			var rv types.Zval
			var objval *types.Zval = op1.Object().Get(&rv)
			// objval.TryAddRefcount()
			ret = BitwiseXorFunction(objval, objval, op2)
			op1.Object().Set(objval)
			// ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && op1.Object().CanDoOperation() {
			if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_BW_XOR, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.Long()
	}
	if !op2.IsLong() {
		if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_BW_XOR, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if zend.EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.Long()
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval ^ op2_lval)
	return types.SUCCESS
}
func ShiftLeftFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	for {
		if !op1.IsLong() {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.Long()
					break
				}
			}
			if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
				var ret int
				var rv types.Zval
				var objval *types.Zval = op1.Object().Get(&rv)
				// objval.TryAddRefcount()
				ret = ShiftLeftFunction(objval, objval, op2)
				op1.Object().Set(objval)
				// ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && op1.Object().CanDoOperation() {
				if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_SL, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.Long()
		}
		break
	}
	for {
		if !op2.IsLong() {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.Long()
					break
				}
			}
			if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_SL, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op2_lval = op2.Long()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where << 64 + x == << x */

	if zend.ZendUlong(op2_lval >= zend.SIZEOF_ZEND_LONG*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				// ZvalPtrDtor(result)
			}
			result.SetLong(0)
			return types.SUCCESS
		} else {
			if zend.CurrEX() != nil && zend.CG__().GetInCompilation() == 0 {
				faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				faults.ErrorNoreturn(faults.E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}

	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	result.SetLong(zend_long(zend.ZendUlong(op1_lval << op2_lval)))
	return types.SUCCESS
}
func ShiftRightFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval zend.ZendLong
	var op2_lval zend.ZendLong
	for {
		if !op1.IsLong() {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.Long()
					break
				}
			}
			if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
				var ret int
				var rv types.Zval
				var objval *types.Zval = op1.Object().Get(&rv)
				// objval.TryAddRefcount()
				ret = ShiftRightFunction(objval, objval, op2)
				op1.Object().Set(objval)
				// ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && op1.Object().CanDoOperation() {
				if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_SR, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.Long()
		}
		break
	}
	for {
		if !op2.IsLong() {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.Long()
					break
				}
			}
			if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_SR, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if zend.EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op2_lval = op2.Long()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where >> 64 + x == >> x */

	if zend.ZendUlong(op2_lval >= zend.SIZEOF_ZEND_LONG*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				// ZvalPtrDtor(result)
			}
			result.SetLong(lang.Cond(op1_lval < 0, -1, 0))
			return types.SUCCESS
		} else {
			if zend.CurrEX() != nil && zend.CG__().GetInCompilation() == 0 {
				faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				faults.ErrorNoreturn(faults.E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}
	}
	if op1 == result {
		// ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval >> op2_lval)
	return types.SUCCESS
}
func ConcatFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var orig_op1 *types.Zval = op1
	var op1_copy types.Zval
	var op2_copy types.Zval
	op1_copy.SetUndef()
	op2_copy.SetUndef()
	for {
		if !op1.IsString() {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsString() {
					break
				}
			}
			if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
				var ret int
				var rv types.Zval
				var objval *types.Zval = op1.Object().Get(&rv)
				// objval.TryAddRefcount()
				ret = ConcatFunction(objval, objval, op2)
				op1.Object().Set(objval)
				// ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && op1.Object().CanDoOperation() {
				if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_CONCAT, result, op1, op2) {
					return types.SUCCESS
				}
			} else if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_CONCAT, result, op1, op2) {
				return types.SUCCESS
			}
			op1_copy.SetString(ZvalGetString(op1))
			if zend.EG__().GetException() != nil {

				if orig_op1 != result {
					result.SetUndef()
				}
				return types.FAILURE
			}
			if result == op1 {
				if op1 == op2 {
					op2 = &op1_copy
				}
			}
			op1 = &op1_copy
		}
		break
	}
	for {
		if !op2.IsString() {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsString() {
					break
				}
			}
			if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_CONCAT, result, op1, op2) {
				return types.SUCCESS
			}
			op2_copy.SetString(ZvalGetString(op2))
			if zend.EG__().GetException() != nil {

				if orig_op1 != result {
					result.SetUndef()
				}
				return types.FAILURE
			}
			op2 = &op2_copy
		}
		break
	}
	if op1.String().GetLen() == 0 {
		if result != op2 {
			types.ZVAL_COPY(result, op2)
		}
	} else if op2.String().GetLen() == 0 {
		if result != op1 {
			types.ZVAL_COPY(result, op1)
		}
	} else {
		var op1_len int = op1.String().GetLen()
		var op2_len int = op2.String().GetLen()
		if op1_len > types.STR_MAX_LEN-op2_len {
			faults.ThrowError(nil, "String size overflow")
			if orig_op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}

		/* This has to happen first to account for the cases where result == op1 == op2 and
		 * the realloc is done. In this case this line will also update Z_STRVAL_P(op2) to
		 * point to the new string. The first op2_len bytes of result will still be the same. */
		result.SetStringVal(op1.StringVal() + op2.StringVal())
	}

	return types.SUCCESS
}
func StringCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var str1 = ZvalGetStrVal(op1)
	var str2 = ZvalGetStrVal(op2)
	return strings.Compare(str1, str2)
}
func StringCaseCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var str1 = ZvalGetStrVal(op1)
	var str2 = ZvalGetStrVal(op2)
	return ascii.StrCaseCompare(str1, str2)
}
func StringLocaleCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var str1 = ZvalGetStrVal(op1)
	var str2 = ZvalGetStrVal(op2)
	return b.StrColl(str1, str2)
}
func NumericCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(op1)
	d2 = ZvalGetDouble(op2)
	return zend.ZEND_NORMALIZE_BOOL(d1 - d2)
}
func ConvertCompareResultToLong(result *types.Zval) {
	if result.IsDouble() {
		result.SetLong(zend.ZEND_NORMALIZE_BOOL(result.Double()))
	} else {
		ConvertToLong(result)
	}
}

func CompareFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var ret int
	var converted int = 0
	var op1_copy types.Zval
	var op2_copy types.Zval
	var op_free *types.Zval
	var tmp_free types.Zval
	for true {
		switch TypePair(op1.GetType(), op2.GetType()) {
		case TypePair(types.IsLong, types.IsLong):
			result.SetLong(lang.CondF2(op1.Long() > op2.Long(), 1, func() int {
				if op1.Long() < op2.Long() {
					return -1
				} else {
					return 0
				}
			}))
			return types.SUCCESS
		case TypePair(types.IsDouble, types.IsLong):
			result.SetDouble(op1.Double() - float64(op2.Long()))
			result.SetLong(zend.ZEND_NORMALIZE_BOOL(result.Double()))
			return types.SUCCESS
		case TypePair(types.IsLong, types.IsDouble):
			result.SetDouble(float64(op1.Long() - op2.Double()))
			result.SetLong(zend.ZEND_NORMALIZE_BOOL(result.Double()))
			return types.SUCCESS
		case TypePair(types.IsDouble, types.IsDouble):
			if op1.Double() == op2.Double() {
				result.SetLong(0)
			} else {
				result.SetDouble(op1.Double() - op2.Double())
				result.SetLong(zend.ZEND_NORMALIZE_BOOL(result.Double()))
			}
			return types.SUCCESS
		case TypePair(types.IsArray, types.IsArray):
			result.SetLong(ZendCompareArrays(op1, op2))
			return types.SUCCESS
		case TypePair(types.IsNull, types.IsNull):
			fallthrough
		case TypePair(types.IsNull, types.IsFalse):
			fallthrough
		case TypePair(types.IsFalse, types.IsNull):
			fallthrough
		case TypePair(types.IsFalse, types.IsFalse):
			fallthrough
		case TypePair(types.IsTrue, types.IsTrue):
			result.SetLong(0)
			return types.SUCCESS
		case TypePair(types.IsNull, types.IsTrue):
			result.SetLong(-1)
			return types.SUCCESS
		case TypePair(types.IsTrue, types.IsNull):
			result.SetLong(1)
			return types.SUCCESS
		case TypePair(types.IsString, types.IsString):
			if op1.String() == op2.String() {
				result.SetLong(0)
				return types.SUCCESS
			}
			result.SetLong(ZendiSmartStrcmp(op1.StringVal(), op2.StringVal()))
			return types.SUCCESS
		case TypePair(types.IsNull, types.IsString):
			result.SetLong(lang.Cond(op2.String().GetLen() == 0, 0, -1))
			return types.SUCCESS
		case TypePair(types.IsString, types.IsNull):
			result.SetLong(lang.Cond(op1.String().GetLen() == 0, 0, 1))
			return types.SUCCESS
		case TypePair(types.IsObject, types.IsNull):
			result.SetLong(1)
			return types.SUCCESS
		case TypePair(types.IsNull, types.IsObject):
			result.SetLong(-1)
			return types.SUCCESS
		default:
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				continue
			}
			if op1.IsObject() && op1.Object().CanCompare() {
				ret = op1.Object().Compare(result, op1, op2)
				if !result.IsLong() {
					ConvertCompareResultToLong(result)
				}
				return ret
			} else if op2.IsObject() && op2.Object().CanCompare() {
				ret = op2.Object().Compare(result, op1, op2)
				if !result.IsLong() {
					ConvertCompareResultToLong(result)
				}
				return ret
			}
			if op1.IsObject() && op2.IsObject() {
				if op1.Object() == op2.Object() {
					/* object handles are identical, apparently this is the same object */
					result.SetLong(0)
					return types.SUCCESS
				}
				if op1.Object().CanCompareObjectsTo(op2.Object()) {
					result.SetLong(op1.Object().CompareObjectsTo(op2.Object()))
					return types.SUCCESS
				}
			}
			if op1.IsObject() {
				if op1.Object().CanGet() {
					var rv types.Zval
					op_free = op1.Object().Get(&rv)
					ret = CompareFunction(result, op_free, op2)
					//ZendFreeObjGetResult(op_free)
					return ret
				} else if !op2.IsObject() && op1.Object().CanCast() {
					tmp_free.SetUndef()
					if op1.Object().Cast(&tmp_free, lang.CondF2(op2.IsFalse() || op2.IsTrue(), types.IsBool, func() __auto__ { return op2.GetType() })) == types.FAILURE {
						result.SetLong(1)
						//ZendFreeObjGetResult(&tmp_free)
						return types.SUCCESS
					}
					ret = CompareFunction(result, &tmp_free, op2)
					//ZendFreeObjGetResult(&tmp_free)
					return ret
				}
			}
			if op2.IsObject() {
				if op2.Object().CanGet() {
					var rv types.Zval
					op_free = op2.Object().Get(&rv)
					ret = CompareFunction(result, op1, op_free)
					//ZendFreeObjGetResult(op_free)
					return ret
				} else if !op1.IsObject() && op2.Object().CanCast() {
					tmp_free.SetUndef()

					var castType = op1.GetType()
					if op1.IsBool() {
						castType = types.IsBool
					}

					if op2.Object().Cast(&tmp_free, castType) == types.FAILURE {
						result.SetLong(-1)
						return types.SUCCESS
					}
					ret = CompareFunction(result, op1, &tmp_free)
					return ret
				} else if op1.IsObject() {
					result.SetLong(1)
					return types.SUCCESS
				}
			}
			if converted == 0 {
				if op1.GetType() < types.IsTrue {
					result.SetLong(lang.Cond(ZvalIsTrue(op2), -1, 0))
					return types.SUCCESS
				} else if op1.IsTrue() {
					result.SetLong(lang.Cond(ZvalIsTrue(op2), 0, 1))
					return types.SUCCESS
				} else if op2.GetType() < types.IsTrue {
					result.SetLong(lang.Cond(ZvalIsTrue(op1), 1, 0))
					return types.SUCCESS
				} else if op2.IsTrue() {
					result.SetLong(lang.Cond(ZvalIsTrue(op1), 0, -1))
					return types.SUCCESS
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 1)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 1)
					if zend.EG__().GetException() != nil {
						if result != op1 {
							result.SetUndef()
						}
						return types.FAILURE
					}
					converted = 1
				}
			} else if op1.IsArray() {
				result.SetLong(1)
				return types.SUCCESS
			} else if op2.IsArray() {
				result.SetLong(-1)
				return types.SUCCESS
			} else {
				b.Assert(false)
				faults.ThrowError(nil, "Unsupported operand types")
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		}
	}
}
func HashZvalIdenticalFunction(z1 *types.Zval, z2 *types.Zval) int {
	/* is_identical_function() returns 1 in case of identity and 0 in case
	 * of a difference;
	 * whereas this comparison function is expected to return 0 on identity,
	 * and non zero otherwise.
	 */
	z1 = types.ZVAL_DEREF(z1)
	z2 = types.ZVAL_DEREF(z2)
	return FastIsNotIdenticalFunction(z1, z2)
}
func ZendIsIdentical(op1 *types.Zval, op2 *types.Zval) bool {
	if op1.GetType() != op2.GetType() {
		return false
	}
	switch op1.GetType() {
	case types.IsNull, types.IsFalse, types.IsTrue:
		return true
	case types.IsLong:
		return op1.Long() == op2.Long()
	case types.IsResource:
		return op1.Resource() == op2.Resource()
	case types.IsDouble:
		return op1.Double() == op2.Double()
	case types.IsString:
		return op1.StringVal() == op2.StringVal()
	case types.IsArray:
		return op1.Array() == op2.Array() || types.ZendHashCompare(op1.Array(), op2.Array(), HashZvalIdenticalFunction, 1) == 0
	case types.IsObject:
		return op1.Object() == op2.Object()
	default:
		return false
	}
}
func IsIdenticalFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	result.SetBool(ZendIsIdentical(op1, op2))
	return types.SUCCESS
}
func IsNotIdenticalFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	result.SetBool(!ZendIsIdentical(op1, op2))
	return types.SUCCESS
}
func IsEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	result.SetBool(result.Long() == 0)
	return types.SUCCESS
}
func IsNotEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	result.SetBool(result.Long() != 0)
	return types.SUCCESS
}
func IsSmallerFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	result.SetBool(result.Long() < 0)
	return types.SUCCESS
}
func IsSmallerOrEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	result.SetBool(result.Long() <= 0)
	return types.SUCCESS
}
func InstanceofClass(instance_ce *types.ClassEntry, ce *types.ClassEntry) bool {
	for {
		if instance_ce == ce {
			return true
		}
		instance_ce = instance_ce.GetParent()
		if instance_ce == nil {
			break
		}
	}
	return false
}
func InstanceofInterface(instance_ce *types.ClassEntry, ce *types.ClassEntry) bool {
	if instance_ce.GetNumInterfaces() != 0 {
		b.Assert(instance_ce.IsResolvedInterfaces())
		for i := 0; i < instance_ce.GetNumInterfaces(); i++ {
			if instance_ce.GetInterfaces()[i] == ce {
				return true
			}
		}
	}
	return instance_ce == ce
}
func InstanceofFunctionEx(instance_ce *types.ClassEntry, ce *types.ClassEntry, isInterface bool) bool {
	if isInterface {
		b.Assert(ce.IsInterface())
		return InstanceofInterface(instance_ce, ce)
	} else {
		b.Assert(!ce.IsInterface())
		return InstanceofClass(instance_ce, ce)
	}
}
func InstanceofFunction(instance_ce *types.ClassEntry, ce *types.ClassEntry) bool {
	if ce.IsInterface() {
		return InstanceofInterface(instance_ce, ce)
	} else {
		return InstanceofClass(instance_ce, ce)
	}
}
func IncrementString(str *types.Zval) {
	// notice: 前置要求 str 必须是 IS_STRING 类型，且其值不为数字字符串
	str.SetStringVal(IncrementStringEx(str.StringVal()))
}

func IncrementStringEx(str string) string {
	// notice: 前置要求 str 其值不为数字字符串
	if str == "" {
		return "1"
	}

	s := []byte(str)
	last := 0
	carry := false
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if ascii.IsLower(c) {
			last = LOWER_CASE
			carry = c == 'z'
			s[i] = lang.Cond(carry, 'a', c+1)
		} else if ascii.IsUpper(c) {
			last = UPPER_CASE
			carry = c == 'Z'
			s[i] = lang.Cond(carry, 'A', c+1)
		} else if ascii.IsDigit(c) {
			last = NUMERIC
			carry = c == '9'
			s[i] = lang.Cond(carry, '0', c+1)
		} else {
			carry = false
			break
		}
		if !carry {
			break
		}
	}

	if carry {
		switch last {
		case NUMERIC:
			return "1" + string(s)
		case UPPER_CASE:
			return "A" + string(s)
		case LOWER_CASE:
			return "a" + string(s)
		}
	}

	return string(s)
}

func IncrementFunction(op1 *types.Zval) int {
try_again:
	switch op1.GetType() {
	case types.IsLong:
		FastLongIncrementFunction(op1)
	case types.IsDouble:
		op1.SetDouble(op1.Double() + 1)
	case types.IsNull:
		op1.SetLong(1)
	case types.IsString:
		var lval zend.ZendLong
		var dval float64
		switch IsNumericString(op1.String().GetStr(), &lval, &dval, 0) {
		case types.IsLong:

			if lval == zend.ZEND_LONG_MAX {

				/* switch to double */

				var d float64 = float64(lval)
				op1.SetDouble(d + 1)
			} else {
				op1.SetLong(lval + 1)
			}
		case types.IsDouble:

			op1.SetDouble(dval + 1)
		default:
			/* Perl style string increment */
			IncrementString(op1)
		}
	case types.IsObject:
		if op1.Object().CanGet() && op1.Object().CanSet() {

			/* proxy object */

			var rv types.Zval
			var val *types.Zval
			val = op1.Object().Get(&rv)
			// val.TryAddRefcount()
			IncrementFunction(val)
			op1.Object().Set(val)
			// ZvalPtrDtor(val)
		} else if op1.Object().CanDoOperation() {
			var op2 types.Zval
			var res int
			op2.SetLong(1)
			res = op1.Object().DoOperation(zend.ZEND_ADD, op1, op1, &op2)
			return res
		}
		return types.FAILURE
	case types.IsRef:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		return types.FAILURE
	}
	return types.SUCCESS
}
func DecrementFunction(op1 *types.Zval) int {
	var lval zend.ZendLong
	var dval float64
try_again:
	switch op1.GetType() {
	case types.IsLong:
		FastLongDecrementFunction(op1)
	case types.IsDouble:
		op1.SetDouble(op1.Double() - 1)
	case types.IsString:
		if op1.String().GetLen() == 0 {

			op1.SetLong(-1)
			break
		}
		switch IsNumericString(op1.String().GetStr(), &lval, &dval, 0) {
		case types.IsLong:

			if lval == zend.ZEND_LONG_MIN {
				var d float64 = float64(lval)
				op1.SetDouble(d - 1)
			} else {
				op1.SetLong(lval - 1)
			}
		case types.IsDouble:

			op1.SetDouble(dval - 1)
		}
	case types.IsObject:
		if op1.Object().CanGet() && op1.Object().CanSet() {

			/* proxy object */

			var rv types.Zval
			var val *types.Zval
			val = op1.Object().Get(&rv)
			// val.TryAddRefcount()
			DecrementFunction(val)
			op1.Object().Set(val)
			// ZvalPtrDtor(val)
		} else if op1.Object().CanDoOperation() {
			var op2 types.Zval
			var res int
			op2.SetLong(1)
			res = op1.Object().DoOperation(zend.ZEND_SUB, op1, op1, &op2)
			return res
		}
		return types.FAILURE
	case types.IsRef:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendObjectIsTrue(op *types.Zval) bool {
	if op.Object().CanCast() {
		var tmp types.Zval
		if op.Object().Cast(&tmp, types.IsBool) == types.SUCCESS {
			return tmp.IsTrue()
		}
		faults.Error(faults.E_RECOVERABLE_ERROR, "Object of class %s could not be converted to bool", op.Object().GetCe().Name())
	} else if op.Object().CanGet() {
		var result bool
		var rv types.Zval
		var tmp *types.Zval = op.Object().Get(&rv)
		if !tmp.IsObject() {
			/* for safety - avoid loop */
			result = ZvalIsTrue(tmp)
			return result
		}
	}
	return true
}
func ZendStrTolowerCopy(dest *byte, source *byte, length int) *byte {
	var str *uint8 = (*uint8)(source)
	var result *uint8 = (*uint8)(dest)
	var end *uint8 = str + length
	for str < end {
		lang.PostInc(&(*result)) = ascii.ToLower(lang.PostInc(&(*str)))
	}
	*result = '0'
	return dest
}
func ZendBinaryStrncmp(s1 string, s2 string, length int) int {
	if len(s1) > length {
		s1 = s1[:length]
	}
	if len(s2) > length {
		s2 = s2[:length]
	}
	return strings.Compare(s1, s2)
}
func ZendBinaryStrcasecmp(s1 string, s2 string) int { return ascii.StrCaseCompare(s1, s2) }
func ZendBinaryStrncasecmp(s1 string, s2 string, length int) int {
	if len(s1) > length {
		s1 = s1[:length]
	}
	if len(s2) > length {
		s2 = s2[:length]
	}
	return ascii.StrCaseCompare(s1, s2)
}
func ZendiSmartStreq(s1 string, s2 string) bool {
	r1 := zend.StrToNumberEx(s1, 0)
	r2 := zend.StrToNumberEx(s2, 0)

	if r1.IsSucc() && r2.IsSucc() {
		if r1.Overflow() != 0 && r1.Overflow() == r2.Overflow() && r1.Float()-r2.Float() == 0.0 {
			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */
			return s1 == s2
		}

		// int vs. int
		if r1.IsInt() && r2.IsInt() {
			return r1.Int() == r2.Int()
		}

		//
		if !r1.IsFloat() {
			if r2.Overflow() != 0 {
				/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */
				return false
			}
		} else if !r2.IsFloat() {
			if r1.Overflow() != 0 {
				return false
			}
		}

		d1 := r1.ToFloat()
		d2 := r2.ToFloat()
		if d1 == d2 && !(core.ZendFinite(d1)) {
			/* Both values overflowed and have the same sign,
			 * so a numeric comparison would be inaccurate */
			return s1 == s2
		}

		return d1 == d2
	}

	return s1 == s2
}
func ZendiSmartStrcmp(s1 string, s2 string) int {
	var r1, r2 conv.ParseNumberResult
	r1 = zend.StrToNumberEx(s1, 0)
	if !r1.IsSucc() {
		goto string_cmp
	}
	r2 = zend.StrToNumberEx(s2, 0)
	if !r2.IsSucc() {
		goto string_cmp
	}

	if r1.Overflow() != 0 && r1.Overflow() == r2.Overflow() && r1.Float()-r2.Float() == 0.0 {
		/* both values are integers overflown to the same side, and the
		 * double comparison may have resulted in crucial accuracy lost */
		goto string_cmp
	}
	if r1.IsFloat() || r2.IsFloat() {
		dval1, dval2 := r1.Float(), r2.Float()
		if r1.IsInt() {
			if r2.Overflow() != 0 {
				/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */
				return -1 * r2.Overflow()
			}
			dval1 = float64(r1.Int())
		} else if r2.IsInt() {
			if r1.Overflow() != 0 {
				return r1.Overflow()
			}
			dval2 = float64(r2.Int())
		} else if r2.Float() == r2.Float() && !(core.ZendFinite(r1.Float())) {
			/* Both values overflowed and have the same sign,
			 * so a numeric comparison would be inaccurate */
			goto string_cmp
		}
		return b.Compare(dval1, dval2)
	} else {
		return b.Compare(r1.Int(), r2.Int())
	}

string_cmp:
	return strings.Compare(s1, s2)
}
func HashZvalCompareFunction(z1 *types.Zval, z2 *types.Zval) int {
	var result types.Zval
	if CompareFunction(&result, z1, z2) == types.FAILURE {
		return 1
	}
	return result.Long()
}
func ZendCompareSymbolTables(ht1 *types.Array, ht2 *types.Array) int {
	if ht1 == ht2 {
		return 0
	} else {
		return types.ZendHashCompare(ht1, ht2, HashZvalCompareFunction, 0)
	}
}
func ZendCompareArrays(a1 *types.Zval, a2 *types.Zval) int {
	return ZendCompareSymbolTables(a1.Array(), a2.Array())
}

func ZendDvalToLvalSlow(d float64) zend.ZendLong {
	dmod := math.Mod(d, 1<<64)
	if dmod > math.MaxInt {
		dmod -= 1 << 64
	} else if dmod < math.MinInt {
		dmod += 1 << 64
	}
	return zend.ZendLong(dmod)
}

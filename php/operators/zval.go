package operators

import (
	"fmt"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/shim/cmp"
	"strconv"
)

// bool
func (op *Operator) IsTrue(v Val) bool { return op.ToBool(v) }
func (op *Operator) ToBool(v Val) bool {
again:
	switch v.Type() {
	case types.IsTrue:
		return true
	case types.IsLong:
		return v.Long() != 0
	case types.IsDouble:
		return v.Double() != 0
	case types.IsString:
		str := v.String()
		return str != "" && str != "0"
	case types.IsArray:
		return v.Array().Len() != 0
	case types.IsObject:
		dst := op.convertObjectToType(v.Object(), types.IsBool)
		if dst != nil {
			return dst.IsTrue()
		}
		return true
	case types.IsResource:
		return v.ResourceHandle() != 0
	case types.IsRef:
		v = v.RefVal()
		goto again
	}
	return false
}

// long

func (op *Operator) ToLong(v Val) int      { return op.ToLongEx(v, true) }
func (op *Operator) ToLongNoisy(v Val) int { return op.ToLongEx(v, false) }
func (op *Operator) ToLongEx(v Val, silent bool) int {
	// fast
	if v.IsLong() {
		return v.Long()
	}

	// common
again:
	switch v.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return 0
	case types.IsTrue:
		return 1
	case types.IsResource:
		return v.ResourceHandle()
	case types.IsLong:
		return v.Long()
	case types.IsDouble:
		return DoubleToLong(v.Double())
	case types.IsString:
		var r Val = op.parseNumberPrefix(v.String(), silent)
		if r == nil {
			if !silent {
				op.Error(perr.E_WARNING, "A non-numeric value encountered")
			}
			return 0
		}
		if r.IsLong() {
			return r.Long()
		} else { //  r.IsDouble()
			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */
			return DoubleToLongCap(r.Double())
		}
	case types.IsArray:
		if v.Array().Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IsObject:
		dst := op.convertObjectToType(v.Object(), types.IsLong)
		if dst.IsLong() {
			return dst.Long()
		} else {
			return 1
		}
	case types.IsRef:
		v = v.DeRef()
		goto again
	default:
		return 0
	}
}

// TryToLong。 相比 ToLong，不考虑 Array/Object/Resource 等复杂类型。
func (op *Operator) TryToLong(v Val) (int, bool) {
	v = v.DeRef()
	if v.Type() < types.IsString {
		return op.ToLong(v), true
	} else if v.IsString() {
		v, err := strconv.Atoi(v.String())
		if err == nil {
			return v, true
		}
	}
	return 0, false
}

// double
func (op *Operator) ToDouble(v Val) float64 {
	if v.IsDouble() {
		return v.Double()
	}

	v = v.DeRef()
	switch v.Type() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		return 0.0
	case types.IsTrue:
		return 1.0
	case types.IsResource:
		return float64(v.ResourceHandle())
	case types.IsLong:
		return float64(v.Long())
	case types.IsDouble:
		return v.Double()
	case types.IsString:
		return ParseDouble(v.String())
	case types.IsArray:
		if v.Array().Len() != 0 {
			return 1.0
		} else {
			return 0.0
		}
	case types.IsObject:
		dst := op.convertObjectToType(v.Object(), types.IsDouble)
		if dst.IsDouble() {
			return dst.Double()
		} else {
			return 1.0
		}
	default:
		return 0.0
	}
}

// scalar to number
func (op *Operator) ToNumber(v Val) Val { return op.ToNumberEx(v, true) }
func (op *Operator) ToNumberEx(v Val, silent bool) Val {
	switch v.Type() {
	case types.IsNull, types.IsFalse:
		return Long(0)
	case types.IsTrue:
		return Long(1)
	case types.IsLong:
		return Long(v.Long())
	case types.IsDouble:
		return Double(v.Double())
	case types.IsString:
		r := op.parseNumberPrefix(v.String(), silent)
		if r == nil {
			if !silent {
				op.Error(perr.E_WARNING, "A non-numeric value encountered")
			}
			return Long(0)
		}
		return r
	case types.IsResource:
		var l = v.ResourceHandle()
		return Long(l)
	case types.IsObject:
		dst := op.convertObjectToType(v.Object(), types.IsNumber)
		if op.HasException() {
			return Long(1)
		}
		if dst.IsLong() || dst.IsDouble() {
			return dst
		} else {
			return Long(1)
		}
	default:
		// todo fail log
		return nil
	}
}

// string
func (op *Operator) ToStrVal(v Val) string {
	str, _ := op.toStrEx(v, false)
	return str
}
func (op *Operator) ToStr(v Val) (string, bool) {
	return op.toStrEx(v, false)
}
func (op *Operator) TryToStrVal(v Val) string {
	str, _ := op.toStrEx(v, true)
	return str
}
func (op *Operator) TryToStr(v Val) (string, bool) {
	return op.toStrEx(v, true)
}

/**
 * 从 Zval 转字符串
 * @return string 返回的字符串值。
 * @return bool   是否成功。
 */
func (op *Operator) toStrEx(v Val, try bool) (string, bool) {
	v = v.DeRef()
	switch v.Type() {
	case types.IsString:
		return v.String(), true
	case types.IsUndef, types.IsNull, types.IsFalse:
		return "", true
	case types.IsTrue:
		return "1", true
	case types.IsResource:
		return fmt.Sprintf("Resource id #%d", v.ResourceHandle()), true
	case types.IsLong:
		return strconv.Itoa(v.Long()), true
	case types.IsDouble:
		return fmt.Sprintf("%.*G", op.Precision(), v.Double()), true
	case types.IsArray:
		op.Error(perr.E_NOTICE, "Array to string conversion")
		if try && op.HasException() {
			return "", false
		}
		return "Array", true
	case types.IsObject:
		if tmp, ok := v.Object().Cast(types.IsString); ok {
			return tmp.String(), true
		}
		if !op.HasException() {
			op.ThrowError(nil, fmt.Sprintf("Object of class %s could not be converted to string", v.Object().CeName()))
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

// array
func (op *Operator) ToArray(v Val) *types.Array {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		return v.Array()
	case types.IsObject:
		return op.ObjectGetArray(v.Object())
	case types.IsNull:
		return types.NewArray()
	default:
		return types.NewArrayOf(v)
	}
}

// object
func (op *Operator) ToObject(v Val) *types.Object {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		var ht = v.Array()
		// todo
		return op.NewObject(ht)
	case types.IsObject:
		return v.Object()
	case types.IsNull:
		return op.NewObject(nil)
	default:
		obj := op.NewObject(nil)
		//obj.GetPropertiesArray().KeyAdd(types.STR_SCALAR, v.CloneValue())
		return obj
	}
}

func (op *Operator) convertObjectToType(obj *types.Object, ctype types.ZvalType) Val {
	if result, ok := obj.Cast(ctype); ok {
		return result
	} else if obj.CanCast() {
		op.Error(perr.E_RECOVERABLE_ERROR, fmt.Sprintf("Object of class %s could not be converted to %s", obj.CeName(), types.ZendGetTypeByConst(ctype)))
	}
	return nil
}

// compare
func (op *Operator) Compare(v1 Val, v2 Val) int {
	result, ok := op.CompareEx(v1, v2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}
func (op *Operator) CompareEx(v1 Val, v2 Val) (int, bool) {
	var converted int = 0

	v1 = v1.DeRef()
	v2 = v2.DeRef()
	for {
		switch typePair(v1, v2) {
		case IsLongLong:
			return cmp.Compare(v1.Long(), v2.Long()), true
		case IsLongDouble, IsDoubleLong, IsDoubleDouble:
			d1 := fastGetDouble(v1)
			d2 := fastGetDouble(v2)
			return cmp.Compare(d1, d2), true
		case IsArrayArray:
			return op.CompareArray(v1.Array(), v2.Array()), true
		case IsNullNull, IsNullFalse, IsFalseNull, IsFalseFalse, IsTrueTrue:
			return 0, true
		case IsNullTrue:
			return -1, true
		case IsTrueNull:
			return 1, true
		case IsStringString:
			if v1.String() == v2.String() {
				return 0, true
			}
			return SmartStrCompare(v1.String(), v2.String()), true
		case IsNullString:
			return lang.Cond(len(v2.String()) == 0, 0, -1), true
		case IsStringNull:
			return lang.Cond(len(v1.String()) == 0, 0, 1), true
		case IsObjectNull:
			return 1, true
		case IsNullObject:
			return -1, true
		default:
			if v1.IsObject() && v1.Object().CanCompare() {
				return op.opObjectCompare(v1.Object(), v1, v2)
			} else if v2.IsObject() && v2.Object().CanCompare() {
				return op.opObjectCompare(v2.Object(), v1, v2)
			}
			if v1.IsObject() && v2.IsObject() {
				if v1.Object() == v2.Object() {
					/* object handles are identical, apparently this is the same object */
					return 0, true
				}
				if retval, ok := v1.Object().CompareObjectsTo(v2.Object()); ok {
					return retval, true
				}
				return 1, true
			}
			if v1.IsObject() && !v2.IsObject() && v1.Object().CanCast() {
				if tmp, ok := v1.Object().Cast(v2.Type()); ok {
					return op.CompareEx(tmp, v2)
				} else {
					return 1, true
				}
			}
			if v2.IsObject() && !v1.IsObject() && v2.Object().CanCast() {
				if tmp, ok := v2.Object().Cast(v1.Type()); ok {
					return op.CompareEx(v1, tmp)
				} else {
					return -1, true
				}
			}

			if converted == 0 {
				if v1.Type() < types.IsTrue {
					return lang.Cond(op.IsTrue(v2), -1, 0), true
				} else if v1.IsTrue() {
					return lang.Cond(op.IsTrue(v2), 0, 1), true
				} else if v2.Type() < types.IsTrue {
					return lang.Cond(op.IsTrue(v1), 1, 0), true
				} else if v2.IsTrue() {
					return lang.Cond(op.IsTrue(v1), 0, -1), true
				} else {
					v1, v2 = op.opScalarGetNumberEx(v1, v2, true)
					if op.HasException() {
						return 0, false
					}
					converted = 1
				}
			} else if v1.IsArray() {
				return 1, true
			} else if v2.IsArray() {
				return -1, true
			} else {
				lang.Assert(false)
				op.ThrowError(nil, "Unsupported operand types")
				return 0, false
			}
		}
	}
}

func (op *Operator) CompareArray(ht1, ht2 *types.Array) int {
	// todo
	panic("todo")
}

// equals
func (op *Operator) Equals(op1, op2 Val) bool {
	result, ok := op.IsEquals(op1, op2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}

func (op *Operator) IsEquals(v1, v2 Val) (result bool, ok bool) {
	switch typePair(v1, v2) {
	case IsLongLong:
		return v1.Long() == v2.Long(), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(v1)
		d2 := fastGetDouble(v2)
		return d1 == d2, true
	case IsStringString:
		return SmartStrEquals(v1.String(), v2.String()), true
	default:
		ret, ok := op.CompareEx(v1, v2)
		if !ok {
			return false, false
		}
		return ret == 0, true
	}
}

// identical
func (op *Operator) IsIdentical(v1 Val, v2 Val) bool {
	if v1.Type() != v2.Type() {
		return false
	}
	switch v1.Type() {
	case types.IsNull, types.IsFalse, types.IsTrue:
		return true
	case types.IsLong:
		return v1.Long() == v2.Long()
	case types.IsResource:
		return v1.Resource() == v2.Resource()
	case types.IsDouble:
		return v1.Double() == v2.Double()
	case types.IsString:
		return v1.String() == v2.String()
	case types.IsArray:
		return v1.Array() == v2.Array() || op.IsIdenticalArray(v1.Array(), v2.Array())
	case types.IsObject:
		return v1.Object() == v2.Object()
	default:
		return false
	}
}

func (op *Operator) IsIdenticalArray(ht1, ht2 *types.Array) bool {
	// todo
	return false
}

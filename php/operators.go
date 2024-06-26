package php

import (
	"cmp"
	"fmt"
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"math"
	"strconv"
	"strings"
)

// -- zval functions

// bool
func ZvalIsTrue(ctx *Context, v types.Zval) bool {
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
		dst := convertObjectToType(ctx, v.Object(), types.IsBool)
		if dst.IsNotUndef() {
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
func ZvalGetLong(ctx *Context, v types.Zval) int      { return ZvalGetLongEx(ctx, v, true) }
func ZvalGetLongNoisy(ctx *Context, v types.Zval) int { return ZvalGetLongEx(ctx, v, false) }
func ZvalGetLongEx(ctx *Context, v types.Zval, silent bool) int {
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
		var r = opParseNumberPrefix(ctx, v.String(), silent)
		if r.IsUndef() {
			if !silent {
				Error(ctx, perr.E_WARNING, "A non-numeric value encountered")
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
		dst := convertObjectToType(ctx, v.Object(), types.IsLong)
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

// ZvalTryGetLong。 相比 ZvalGetLong，不考虑 Array/Object/Resource 等复杂类型。
func ZvalTryGetLong(ctx *Context, v types.Zval) (int, bool) {
	v = v.DeRef()
	if v.Type() < types.IsString {
		return ZvalGetLong(ctx, v), true
	} else if v.IsString() {
		v, err := strconv.Atoi(v.String())
		if err == nil {
			return v, true
		}
	}
	return 0, false
}

// double
func ZvalGetDouble(ctx *Context, v types.Zval) float64 {
	v = v.DeRef()
	switch v.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return 0.0
	case types.IsTrue:
		return 1.0
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
		dst := convertObjectToType(ctx, v.Object(), types.IsDouble)
		if dst.IsDouble() {
			return dst.Double()
		} else {
			return 1.0
		}
	case types.IsResource:
		return float64(v.ResourceHandle())
	default:
		return 0.0
	}
}

// scalar to number
func ConvertScalarToNumber(ctx *Context, v types.Zval) types.Zval {
	return ConvertScalarToNumberEx(ctx, v, true)
}
func ConvertScalarToNumberEx(ctx *Context, v types.Zval, silent bool) types.Zval {
	switch v.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return Long(0)
	case types.IsTrue:
		return Long(1)
	case types.IsLong:
		return Long(v.Long())
	case types.IsDouble:
		return Double(v.Double())
	case types.IsString:
		r := opParseNumberPrefix(ctx, v.String(), silent)
		if r.IsUndef() {
			if !silent {
				Error(ctx, perr.E_WARNING, "A non-numeric value encountered")
			}
			return Long(0)
		}
		return r
	case types.IsResource:
		var l = v.ResourceHandle()
		return Long(l)
	case types.IsObject:
		dst := convertObjectToType(ctx, v.Object(), types.IsNumber)
		if ctx.EG().HasException() {
			return Long(1)
		}
		if dst.IsLong() || dst.IsDouble() {
			return dst
		} else {
			return Long(1)
		}
	default:
		return v
	}
}

// string
func ZvalGetStrVal(ctx *Context, v types.Zval) string {
	str, _ := zvalGetStrEx(ctx, v, false)
	return str
}
func ZvalGetStr(ctx *Context, v types.Zval) (string, bool) {
	return zvalGetStrEx(ctx, v, false)
}
func ZvalTryGetStrVal(ctx *Context, v types.Zval) string {
	str, _ := zvalGetStrEx(ctx, v, true)
	return str
}
func ZvalTryGetStr(ctx *Context, v types.Zval) (string, bool) {
	return zvalGetStrEx(ctx, v, true)
}

/**
 * 从 Zval 转字符串
 * @return string 返回的字符串值。
 * @return bool   是否成功。
 */
func zvalGetStrEx(ctx *Context, v types.Zval, try bool) (string, bool) {
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
		return FormatDouble(v.Double(), 'G', ctx.EG().Precision()), true
	case types.IsArray:
		Error(ctx, perr.E_NOTICE, "Array to string conversion")
		if try && ctx.EG().HasException() {
			return "", false
		}
		return "Array", true
	case types.IsObject:
		if tmp, ok := v.Object().Cast(types.IsString); ok {
			return tmp.String(), true
		}
		if !ctx.EG().HasException() {
			ThrowError(ctx, nil, fmt.Sprintf("Object of class %s could not be converted to string", v.Object().ClassName()))
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
func ZvalGetArray(ctx *Context, v types.Zval) *types.Array {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		return v.Array()
	case types.IsObject:
		return opObjectGetArray(ctx, v.Object())
	case types.IsNull:
		return types.NewArray()
	default:
		return types.NewArrayOf(v)
	}
}

// object
func ZvalGetObject(ctx *Context, v types.Zval) *types.Object {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		var ht = v.Array()
		// todo
		return opNewObject(ht)
	case types.IsObject:
		return v.Object()
	case types.IsNull:
		return opNewObject(nil)
	default:
		obj := opNewObject(nil)
		//obj.GetPropertiesArray().KeyAdd(types.STR_SCALAR, v.CloneValue())
		return obj
	}
}

func convertObjectToType(ctx *Context, obj *types.Object, ctype types.ZvalType) types.Zval {
	if result, ok := obj.Cast(ctype); ok {
		return result
	} else {
		Error(ctx, perr.E_RECOVERABLE_ERROR, fmt.Sprintf("Object of class %s could not be converted to %s", obj.ClassName(), types.ZendGetTypeByConst(ctype)))
		return types.Undef
	}
}

// compare
func ZvalCompare(ctx *Context, v1 types.Zval, v2 types.Zval) int {
	result, ok := ZvalCompareEx(ctx, v1, v2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}
func ZvalCompareEx(ctx *Context, v1 types.Zval, v2 types.Zval) (int, bool) {
	var converted = 0

	v1 = v1.DeRef()
	v2 = v2.DeRef()
	for {
		if v1 == v2 {
			return 0, true
		}
		switch TypePair(v1, v2) {
		case IsLongLong:
			return cmp.Compare(v1.Long(), v2.Long()), true
		case IsLongDouble, IsDoubleLong, IsDoubleDouble:
			d1 := fastGetDouble(v1)
			d2 := fastGetDouble(v2)
			return cmp.Compare(d1, d2), true
		case IsArrayArray:
			return OpCompareArray(ctx, v1.Array(), v2.Array()), true
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
			if v1.IsObject() {
				// CompareTo
				if v, ok := v1.Object().CompareTo(v2); ok {
					return v, true
				}

				// CompareObjectTo
				if v2.IsObject() {
					if v1.Object() == v2.Object() {
						/* object handles are identical, apparently this is the same object */
						return 0, true
					}
					if retval, ok := v1.Object().CompareObjectTo(v2.Object()); ok {
						return retval, true
					}
					return 1, true
				}

				// !v2.IsObject()
				if tmp, ok := v1.Object().Cast(v2.Type()); ok {
					return ZvalCompareEx(ctx, tmp, v2)
				} else {
					return 1, true
				}
			} else if v2.IsObject() {
				// CompareTo
				if v, ok := v2.Object().CompareTo(v1); ok {
					return -v, true
				}

				// !v1.IsObject()
				if tmp, ok := v2.Object().Cast(v1.Type()); ok {
					return ZvalCompareEx(ctx, v1, tmp)
				} else {
					return -1, true
				}
			}

			if converted == 0 {
				if v1.Type() < types.IsTrue {
					return lang.Cond(ZvalIsTrue(ctx, v2), -1, 0), true
				} else if v1.IsTrue() {
					return lang.Cond(ZvalIsTrue(ctx, v2), 0, 1), true
				} else if v2.Type() < types.IsTrue {
					return lang.Cond(ZvalIsTrue(ctx, v1), 1, 0), true
				} else if v2.IsTrue() {
					return lang.Cond(ZvalIsTrue(ctx, v1), 0, -1), true
				} else {
					v1 = ConvertScalarToNumber(ctx, v1)
					v2 = ConvertScalarToNumber(ctx, v2)
					if ctx.EG().HasException() {
						return 0, false
					}
					converted = 1
				}
			} else if v1.IsArray() {
				return 1, true
			} else if v2.IsArray() {
				return -1, true
			} else {
				assert.Assert(false)
				ThrowError(ctx, nil, "Unsupported operand types")
				return 0, false
			}
		}
	}
}

func OpCompareArray(ctx *Context, ht1, ht2 *types.Array) int {
	var comparer ZvalComparer = func(v1 types.Zval, v2 types.Zval) int {
		if v, ok := ZvalCompareEx(ctx, v1, v2); ok {
			return v
		} else {
			return 1
		}
	}
	return iArrayCompare(ctx, ht1, ht2, comparer, false)
}

// SmartStrCompare: zendi_smart_strcmp
func SmartStrCompare(s1 string, s2 string) int {
	v1, overflow1 := ParseNumberEx(s1)
	v2, overflow2 := ParseNumberEx(s2)
	if v1.IsUndef() || v2.IsUndef() {
		goto stringCmp
	}

	if overflow1 != 0 && overflow1 == overflow2 && v1.Double()-v2.Double() == 0.0 {
		/* both values are integers overflown to the same side, and the
		 * double comparison may have resulted in crucial accuracy lost */
		goto stringCmp
	}
	if v1.IsDouble() || v2.IsDouble() {
		dval1, dval2 := v1.Double(), v2.Double()
		if v1.IsLong() {
			if overflow2 != 0 {
				/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */
				return -1 * overflow2
			}
			dval1 = float64(v1.Long())
		} else if v2.IsLong() {
			if overflow1 != 0 {
				return overflow1
			}
			dval2 = float64(v2.Long())
		} else if v1.Double() == v2.Double() && !(mathkit.IsFinite(v1.Double())) {
			/* Both values overflowed and have the same sign,
			 * so a numeric comparison would be inaccurate */
			goto stringCmp
		}
		return cmp.Compare(dval1, dval2)
	} else {
		return cmp.Compare(v1.Long(), v2.Long())
	}

stringCmp:
	return strings.Compare(s1, s2)
}

// SmartStrEquals: zend_fast_equal_strings
func SmartStrEquals(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) > 0 && s1[0] <= '9' && len(s2) > 0 && s2[0] <= '9' {
		return SmartStrCompare(s1, s2) == 0
	} else {
		return false
	}
}

// equals
func ZvalEquals(ctx *Context, op1, op2 types.Zval) bool {
	result, ok := ZvalEqualsEx(ctx, op1, op2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}

func ZvalEqualsEx(ctx *Context, v1, v2 types.Zval) (result bool, ok bool) {
	v1 = v1.DeRef()
	v2 = v2.DeRef()
	if v1 == v2 {
		return true, true
	}
	switch TypePair(v1, v2) {
	case IsLongLong:
		return v1.Long() == v2.Long(), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(v1)
		d2 := fastGetDouble(v2)
		return d1 == d2, true
	case IsStringString:
		return SmartStrEquals(v1.String(), v2.String()), true
	default:
		ret, ok := ZvalCompareEx(ctx, v1, v2)
		if !ok {
			return false, false
		}
		return ret == 0, true
	}
}

// identical
func ZvalIsIdentical(ctx *Context, v1 types.Zval, v2 types.Zval) bool {
	if v1 == v2 {
		return true
	}
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
		return v1.Array() == v2.Array() || zvalIsIdenticalArray(ctx, v1.Array(), v2.Array())
	case types.IsObject:
		return v1.Object() == v2.Object()
	default:
		return false
	}
}

func zvalIsIdenticalArray(ctx *Context, ht1, ht2 *types.Array) bool {
	var comparer ZvalComparer = func(v1 types.Zval, v2 types.Zval) int {
		v1 = v1.DeRef()
		v2 = v2.DeRef()
		if ZvalIsIdentical(ctx, v1, v2) {
			return 0
		} else {
			return 1
		}
	}
	return iArrayCompare(ctx, ht1, ht2, comparer, true) == 0
}

// Add
func OpAdd(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch TypePair(op1, op2) {
	case IsLongLong:
		return OpAddLong(ctx, op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) + fastGetDouble(op2))
	case IsArrayArray:
		retArr := OpAddArray(ctx, op1.Array(), op2.Array())
		return Array(retArr)
	default:
		if converted {
			// fail
			opThrowError(ctx, nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = opScalarGetNumber(ctx, op1, op2)
		goto again
	}
}

func OpAddLong(ctx *Context, i1, i2 int) types.Zval {
	if sign(i1) == sign(i2) && sign(i1) != sign(i1+i2) { // 判断相加是否越界
		return Double(float64(i1) + float64(i2))
	} else {
		return Long(i1 + i2)
	}
}

func OpAddArray(ctx *Context, a1, a2 *types.Array) *types.Array {
	// todo AddArray
	panic(perr.Unreachable())
}

// Sub
func OpSub(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch TypePair(op1, op2) {
	case IsLongLong:
		return OpSubLong(ctx, op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) - fastGetDouble(op2))
	default:
		if converted {
			// fail
			opThrowError(ctx, nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = opScalarGetNumber(ctx, op1, op2)
		goto again
	}
}

func OpSubLong(ctx *Context, i1, i2 int) types.Zval {
	if sign(i1) != sign(i2) && sign(i1) != sign(i1-i2) { // 判断是否越界
		return Double(float64(i1) - float64(i2))
	} else {
		return Long(i1 - i2)
	}
}

// Mul
func OpMul(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch TypePair(op1, op2) {
	case IsLongLong:
		return OpMulLong(ctx, op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) * fastGetDouble(op2))
	default:
		if converted {
			// fail
			opThrowError(ctx, nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = opScalarGetNumber(ctx, op1, op2)
		goto again
	}
}

func OpMulLong(ctx *Context, i1, i2 int) types.Zval {
	if iVal, dVal, overflow := OpMulLongVal(ctx, i1, i2); !overflow {
		return Long(iVal)
	} else {
		return Double(dVal)
	}
}

func OpMulLongVal(ctx *Context, i1, i2 int) (iVal int, dVal float64, overflow bool) {
	// ZEND_SIGNED_MULTIPLY_LONG
	iVal = i1 * i2
	dVal = float64(i1) * float64(i2)
	delta := dVal - float64(iVal)
	if dVal+delta == dVal { // 判断是否越界
		return iVal, 0, false
	} else {
		return 0, dVal, true
	}
}

// Div
func OpDiv(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch TypePair(op1, op2) {
	case IsLongLong:
		return OpDivLong(ctx, op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		if d2 == 0 {
			Error(ctx, perr.E_WARNING, "Division by zero")
			return Double(math.Inf(int(d1)))
		}
		return Double(d1 / d2)
	default:
		if converted {
			// fail
			opThrowError(ctx, nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = opScalarGetNumber(ctx, op1, op2)
		goto again
	}
}

func OpDivLong(ctx *Context, i1, i2 int) types.Zval {
	if i2 == 0 {
		Error(ctx, perr.E_WARNING, "Division by zero")
		return Double(math.Inf(i1))
	} else if i2 == -1 && i1 == math.MinInt {
		/* Prevent overflow error/crash */
		return Double(float64(i1) / float64(i2))
	}
	if i1%i2 == 0 {
		return Long(i1 / i2)
	} else {
		return Double(float64(i1) / float64(i2))
	}
}

// Mod
func OpMod(ctx *Context, op1, op2 types.Zval) types.Zval {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = ZvalGetLongNoisy(ctx, op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = ZvalGetLongNoisy(ctx, op2)
	}

	if op2Lval == 0 {
		/* modulus by zero */
		opThrowException(ctx, nil, "Modulo by zero")
	}

	return Long(op1Lval % op2Lval)
}

// ShiftLeft (SL)
func OpSL(ctx *Context, op1, op2 types.Zval) types.Zval {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = ZvalGetLongNoisy(ctx, op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = ZvalGetLongNoisy(ctx, op2)
	}

	if op2Lval < 0 {
		opThrowException(ctx, nil, "Bit shift by negative number")
	}

	return Long(op1Lval << op2Lval)
}

// ShiftRight (SR)
func OpSR(ctx *Context, op1, op2 types.Zval) types.Zval {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = ZvalGetLongNoisy(ctx, op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = ZvalGetLongNoisy(ctx, op2)
	}

	if op2Lval < 0 {
		opThrowException(ctx, nil, "Bit shift by negative number")
	}

	return Long(op1Lval >> op2Lval)
}

// Concat
func OpConcat(ctx *Context, op1, op2 types.Zval) types.Zval {
	var s1, s2 string

	op1 = op1.DeRef()
	op2 = op2.DeRef()

	if op1.IsString() {
		s1 = op1.String()
	} else {
		s1 = ZvalGetStrVal(ctx, op1)
	}

	if op2.IsString() {
		s2 = op2.String()
	} else {
		s2 = ZvalGetStrVal(ctx, op2)
	}

	if len(s1)+len(s2) > math.MaxInt {
		opThrowError(ctx, nil, "String size overflow")
	}

	return String(s1 + s2)
}

// Pow
func OpPow(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch TypePair(op1, op2) {
	case IsLongLong:
		return OpPowLong(ctx, op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		return Double(math.Pow(d1, d2))
	default:
		if converted {
			// fail
			opThrowError(ctx, nil, "Unsupported operand types")
		}

		// array type
		if op1.IsArray() {
			return Long(0)
		} else if op2.IsArray() {
			return Long(1)
		}

		// convert
		op1, op2 = opScalarGetNumber(ctx, op1, op2)

		converted = true
		goto again
	}
}

func OpPowLong(ctx *Context, i1, i2 int) types.Zval {
	if i2 >= 0 {
		if i2 == 0 || i1 == 1 {
			return Long(1)
		} else if i1 == 0 {
			return Long(0)
		} else if i2 == 1 {
			return Long(i1)
		}
		if i2 == 0 {
			return Long(1)
		} else if i1 == 0 {
			return Long(0)
		}

		// result = l1 * l2 ^ pow
		// 		  = l1 * (l2 * l2) ^ (pow/2)  (when pow % 2 == 0 )
		// 		  = (l1 * l2) * l2 ^ (pow-1)  (when pow % 2 == 1 )
		l1, l2, pow := 1, i1, i2
		overflow := false
		for pow >= 1 {
			if pow%2 != 0 {
				pow--
				l1, _, overflow = OpMulLongVal(ctx, l1, l2)
			} else {
				i2 /= 2
				l2, _, overflow = OpMulLongVal(ctx, l2, l2)
			}
			if overflow {
				goto doubleVal
			}
		}

		/* pow == 0 */
		return Long(l1)
	}
doubleVal:
	return Double(math.Pow(float64(i1), float64(i2)))
}

// BitwiseAnd
func OpBitwiseAnd(ctx *Context, op1, op2 types.Zval) types.Zval {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() & op2.Long())
	}

	// common
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		return String(string(str))
	}

	// common
	var op1Lval = ZvalGetLongNoisy(ctx, op1)
	var op2Lval = ZvalGetLongNoisy(ctx, op2)
	return Long(op1Lval & op2Lval)
}

// BitwiseOr
func OpBitwiseOr(ctx *Context, op1, op2 types.Zval) types.Zval {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() | op2.Long())
	}

	// common
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] | s2[i]
		}
		return String(string(str))
	}

	var op1Lval = ZvalGetLongNoisy(ctx, op1)
	var op2Lval = ZvalGetLongNoisy(ctx, op2)
	return Long(op1Lval | op2Lval)
}

// BitwiseXor
func OpBitwiseXor(ctx *Context, op1, op2 types.Zval) types.Zval {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() ^ op2.Long())
	}

	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		return String(string(str))
	}

	var op1Lval = ZvalGetLongNoisy(ctx, op1)
	var op2Lval = ZvalGetLongNoisy(ctx, op2)
	return Long(op1Lval ^ op2Lval)
}

// BitwiseNot
func OpBitwiseNot(ctx *Context, op1 types.Zval) types.Zval {
again:
	switch op1.Type() {
	case types.IsLong:
		v := ^op1.Long()
		return Long(v)
	case types.IsDouble:
		v := ^DoubleToLong(op1.Double())
		return Long(v)
	case types.IsString:
		str := []byte(op1.String())
		for i, c := range str {
			str[i] = ^c
		}
		return String(string(str))
	case types.IsRef:
		op1 = op1.DeRef()
		goto again
	default:
		opThrowError(ctx, nil, "Unsupported operand types")
		panic("unreachable")
	}
}

// Coalesce(??)
func OpCoalesce(ctx *Context, op1 types.Zval, op2 func() types.Zval) types.Zval {
	if !op1.IsUndef() && !op1.IsNull() {
		return op1
	}
	return op2()
}

// BooleanAnd
func OpBooleanAnd(ctx *Context, op1 types.Zval, op2 func() types.Zval) types.Zval {
	op1Val := ZvalIsTrue(ctx, op1)
	if !op1Val {
		return types.False
	}

	op2Val := ZvalIsTrue(ctx, op2())
	return Bool(op2Val)
}

// BooleanOr
func OpBooleanOr(ctx *Context, op1 types.Zval, op2 func() types.Zval) types.Zval {
	op1Val := ZvalIsTrue(ctx, op1)
	if op1Val {
		return types.True
	}

	op2Val := ZvalIsTrue(ctx, op2())
	return Bool(op2Val)
}

// BooleanNot
func OpBooleanNot(ctx *Context, op1 types.Zval) types.Zval {
	op1Val := ZvalIsTrue(ctx, op1)
	return Bool(!op1Val)
}

// BooleanXor
func OpBooleanXor(ctx *Context, op1, op2 types.Zval) types.Zval {
	op1Val := ZvalIsTrue(ctx, op1)
	op2Val := ZvalIsTrue(ctx, op2)
	return Bool(lang.Xor(op1Val, op2Val))
}

// Identical
func OpIdentical(ctx *Context, op1, op2 types.Zval) types.Zval {
	return Bool(ZvalIsIdentical(ctx, op1, op2))
}

// NotIdentical
func OpNotIdentical(ctx *Context, op1, op2 types.Zval) types.Zval {
	return Bool(!ZvalIsIdentical(ctx, op1, op2))
}

// Equal
func OpEqual(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalEquals(ctx, op1, op2)
	return Bool(result)
}

// NotEqual
func OpNotEqual(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalEquals(ctx, op1, op2)
	return Bool(!result)
}

// Greater
func OpGreater(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalCompare(ctx, op1, op2)
	return Bool(result > 0)
}

// GreaterOrEqual
func OpGreaterOrEqual(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalCompare(ctx, op1, op2)
	return Bool(result >= 0)
}

// Smaller
func OpSmaller(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalCompare(ctx, op1, op2)
	return Bool(result < 0)
}

// SmallerOrEqual
func OpSmallerOrEqual(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalCompare(ctx, op1, op2)
	return Bool(result <= 0)
}

// Spaceship
func OpSpaceship(ctx *Context, op1, op2 types.Zval) types.Zval {
	result := ZvalCompare(ctx, op1, op2)
	return Long(result)
}

// -- internal methods

func opScalarGetNumber(ctx *Context, v1, v2 types.Zval) (types.Zval, types.Zval) {
	if v1 != v2 {
		v1 = ConvertScalarToNumberEx(ctx, v1, false)
		v2 = ConvertScalarToNumberEx(ctx, v2, false)
	} else {
		v1 = ConvertScalarToNumberEx(ctx, v1, false)
		v2 = v1
	}
	return v1, v2
}

func opParseNumberPrefix(ctx *Context, str string, silent bool) types.Zval {
	zv, matchLen := ParseNumberPrefix(str)
	if matchLen != len(str) && !silent {
		// notice: 此处可能会触发 Exception
		if matchLen > 0 {
			Error(ctx, perr.E_NOTICE, "A non well formed numeric value encountered")
		}
		if ctx.EG().HasException() {
			return types.Undef
		}
	}
	return zv
}

func ZvalToArrayKey(ctx *Context, offset types.Zval) (key types.ArrayKey, ok bool) {
	offset = offset.DeRef()
	switch offset.Type() {
	case types.IsUndef, types.IsNull:
		return types.StrKey(""), true
	case types.IsFalse:
		return types.IdxKey(0), true
	case types.IsTrue:
		return types.IdxKey(1), true
	case types.IsLong:
		return types.IdxKey(offset.Long()), true
	case types.IsDouble:
		return types.IdxKey(DoubleToLong(offset.Double())), true
	case types.IsString:
		return types.NumericKey(offset.String()), true
	default:
		ZendIllegalOffset(ctx)
		return types.IdxKey(0), false
	}
}

// todo
func opThrowError(ctx *Context, exceptionCe *types.Class, message string) {
	ThrowError(ctx, exceptionCe, message)
	panic(perr.Unreachable())
}
func opThrowException(ctx *Context, exceptionCe *types.Class, message string) {
	panic(perr.Todof("opThrowException"))
}
func opObjectGetArray(ctx *Context, obj *types.Object) *types.Array {
	panic(perr.Todof("opObjectGetArray"))
}

func opNewObject(properties *types.Array) *types.Object {
	panic(perr.Todof("opNewObject"))
}

package operators

import (
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"math"
)

// Add
func Add(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return AddLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) + fastGetDouble(op2)), true
	case IsArrayArray:
		retArr := AddArray(op1.Array(), op2.Array())
		return Array(retArr), true
	default:
		if converted {
			// fail
			faults.ThrowError(nil, "Unsupported operand types")
			return nil, false
		}

		// convert
		op1, op2 = opScalarGetNumber(op1, op2)
		if hasException() {
			return nil, false
		}

		converted = true
		goto again
	}
}

func AddLong(i1, i2 int) Val {
	if sign(i1) == sign(i2) && sign(i1) != sign(i1+i2) { // 判断相加是否越界
		return Double(float64(i1) + float64(i2))
	} else {
		return Long(i1 + i2)
	}
}

func AddArray(a1, a2 *types.Array) *types.Array {
	// todo AddArray
	panic("unreachable")
}

// Sub
func Sub(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return SubLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) - fastGetDouble(op2)), true
	default:
		if converted {
			// fail
			faults.ThrowError(nil, "Unsupported operand types")
			return nil, false
		}

		// convert
		op1, op2 = opScalarGetNumber(op1, op2)
		if hasException() {
			return nil, false
		}

		converted = true
		goto again
	}
}

func SubLong(i1, i2 int) Val {
	if sign(i1) != sign(i2) && sign(i1) != sign(i1-i2) { // 判断是否越界
		return Double(float64(i1) - float64(i2))
	} else {
		return Long(i1 - i2)
	}
}

// Mul
func Mul(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return MulLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) * fastGetDouble(op2)), true
	default:
		if converted {
			// fail
			faults.ThrowError(nil, "Unsupported operand types")
			return nil, false
		}

		// convert
		op1, op2 = opScalarGetNumber(op1, op2)
		if hasException() {
			return nil, false
		}

		converted = true
		goto again
	}
}

func MulLong(i1, i2 int) Val {
	if iVal, dVal, overflow := mulLong(i1, i2); !overflow {
		return Long(iVal)
	} else {
		return Double(dVal)
	}
}

func mulLong(i1, i2 int) (iVal int, dVal float64, overflow bool) {
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
func Div(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return DivLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		if d2 == 0 {
			faults.Error(faults.E_WARNING, "Division by zero")
			return Double(math.Inf(int(d1))), true
		}
		return Double(d1 / d2), true
	default:
		if converted {
			// fail
			faults.ThrowError(nil, "Unsupported operand types")
			return nil, false
		}

		// convert
		op1, op2 = opScalarGetNumber(op1, op2)
		if hasException() {
			return nil, false
		}

		converted = true
		goto again
	}
}

func DivLong(i1, i2 int) Val {
	if i2 == 0 {
		faults.Error(faults.E_WARNING, "Division by zero")
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
func Mod(op1, op2 Val) (Val, bool) {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = zvalGetLongNoisy(op1)
		if hasException() {
			return nil, false
		}
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = zvalGetLongNoisy(op2)
		if hasException() {
			return nil, false
		}
	}

	if op2Lval == 0 {
		/* modulus by zero */
		throwIfExecuting(nil, "Modulo by zero")
		return nil, false
	}

	return Long(op1Lval % op2Lval), true
}

// ShiftLeft (SL)
func ShiftLeft(op1, op2 Val) (Val, bool) {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = zvalGetLongNoisy(op1)
		if hasException() {
			return nil, false
		}
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = zvalGetLongNoisy(op2)
		if hasException() {
			return nil, false
		}
	}

	if op2Lval < 0 {
		throwIfExecuting(nil, "Bit shift by negative number")
		return nil, false
	}

	return Long(op1Lval << op2Lval), true
}

// ShiftRight (SR)
func ShiftRight(op1, op2 Val) (Val, bool) {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = zvalGetLongNoisy(op1)
		if hasException() {
			return nil, false
		}
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = zvalGetLongNoisy(op2)
		if hasException() {
			return nil, false
		}
	}

	if op2Lval < 0 {
		throwIfExecuting(nil, "Bit shift by negative number")
		return nil, false
	}

	return Long(op1Lval >> op2Lval), true
}

// Concat
func Concat(op1, op2 Val) (Val, bool) {
	var s1, s2 string

	op1 = op1.DeRef()
	op2 = op2.DeRef()

	if op1.IsString() {
		s1 = op1.String()
	} else {
		s1 = ZvalGetStrVal(op1)
		if hasException() {
			return nil, false
		}
	}

	if op2.IsString() {
		s2 = op2.String()
	} else {
		s2 = ZvalGetStrVal(op2)
		if hasException() {
			return nil, false
		}
	}

	if len(s1)+len(s2) > math.MaxInt {
		faults.ThrowError(nil, "String size overflow")
		return nil, false
	}

	return String(s1 + s2), true
}

// Pow
func Pow(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return PowLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		return Double(math.Pow(d1, d2)), true
	default:
		if converted {
			// fail
			faults.ThrowError(nil, "Unsupported operand types")
			return nil, false
		}

		// array type
		if op1.IsArray() {
			return Long(0), true
		} else if op2.IsArray() {
			return Long(1), true
		}

		// convert
		op1, op2 = opScalarGetNumber(op1, op2)
		if hasException() {
			return nil, false
		}

		converted = true
		goto again
	}
}

func PowLong(i1, i2 int) Val {
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
				l1, _, overflow = mulLong(l1, l2)
			} else {
				i2 /= 2
				l2, _, overflow = mulLong(l2, l2)
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
func BitwiseAnd(op1, op2 Val) (Val, bool) {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() & op2.Long()), true
	}

	// common
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, lang.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		return String(string(str)), true
	}

	// common
	var op1Lval = zvalGetLongNoisy(op1)
	if hasException() {
		return nil, false
	}

	var op2Lval = zvalGetLongNoisy(op2)
	if hasException() {
		return nil, false
	}

	return Long(op1Lval & op2Lval), true
}

// BitwiseOr
func BitwiseOr(op1, op2 Val) (Val, bool) {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() | op2.Long()), true
	}

	// common
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, lang.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] | s2[i]
		}
		return String(string(str)), true
	}

	var op1Lval = zvalGetLongNoisy(op1)
	if hasException() {
		return nil, false
	}

	var op2Lval = zvalGetLongNoisy(op2)
	if hasException() {
		return nil, false
	}

	return Long(op1Lval | op2Lval), true
}

// BitwiseXor
func BitwiseXor(op1, op2 Val) (Val, bool) {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() ^ op2.Long()), true
	}

	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, lang.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		return String(string(str)), true
	}

	var op1Lval = zvalGetLongNoisy(op1)
	if hasException() {
		return nil, false
	}

	var op2Lval = zvalGetLongNoisy(op2)
	if hasException() {
		return nil, false
	}

	return Long(op1Lval ^ op2Lval), true
}

// BooleanXor
func BooleanXor(op1, op2 Val) (Val, bool) {
	op1Val := ZvalIsTrue(op1)
	op2Val := ZvalIsTrue(op2)
	return Bool(lang.Xor(op1Val, op2Val)), true
}

// Identical
func Identical(op1, op2 Val) (Val, bool) {
	return Bool(ZvalIsIdentical(op1, op2)), true
}

// NotIdentical
func NotIdentical(op1, op2 Val) (Val, bool) {
	return Bool(!ZvalIsIdentical(op1, op2)), true
}

// Equal
func Equal(op1, op2 Val) (Val, bool) {
	result, ok := ZvalEquals(op1, op2)
	return Bool(result), ok
}

// NotEqual
func NotEqual(op1, op2 Val) (Val, bool) {
	result, ok := ZvalEquals(op1, op2)
	return Bool(!result), ok
}

// Greater
func Greater(op1, op2 Val) (Val, bool) {
	result, ok := ZvalCompare(op1, op2)
	return Bool(result > 0), ok
}

// GreaterOrEqual
func GreaterOrEqual(op1, op2 Val) (Val, bool) {
	result, ok := ZvalCompare(op1, op2)
	return Bool(result >= 0), ok
}

// Smaller
func Smaller(op1, op2 Val) (Val, bool) {
	result, ok := ZvalCompare(op1, op2)
	return Bool(result < 0), ok
}

// SmallerOrEqual
func SmallerOrEqual(op1, op2 Val) (Val, bool) {
	result, ok := ZvalCompare(op1, op2)
	return Bool(result <= 0), ok
}

// Spaceship
func Spaceship(op1, op2 Val) (Val, bool) {
	result, ok := ZvalCompare(op1, op2)
	return Long(result), ok
}

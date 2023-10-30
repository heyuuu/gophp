package operators

import (
	"github.com/heyuuu/gophp/php/types"
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
			return nil, false
		}

		converted = true
		goto again
	}
}
func SubLong(i1, i2 int) Val {
	if sign(i1) != sign(i2) && sign(i1) != sign(i1-i2) { // 判断相加是否越界
		return Double(float64(i1) - float64(i2))
	} else {
		return Long(i1 + i2)
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
func Div(op1, op2 Val) (Val, bool) {}

// Mod
func Mod(op1, op2 Val) (Val, bool) {}

// Pow
func Pow(op1, op2 Val) (Val, bool) {}

// BitwiseAnd
func BitwiseAnd(op1, op2 Val) (Val, bool) {}

// BitwiseOr
func BitwiseOr(op1, op2 Val) (Val, bool) {}

// BitwiseXor
func BitwiseXor(op1, op2 Val) (Val, bool) {}

// BooleanAnd
func BooleanAnd(op1, op2 Val) (Val, bool) {}

// BooleanOr
func BooleanOr(op1, op2 Val) (Val, bool) {}

// Coalesce
func Coalesce(op1, op2 Val) (Val, bool) {}

// Concat
func Concat(op1, op2 Val) (Val, bool) {}

// Equal
func Equal(op1, op2 Val) (Val, bool) {}

// Greater
func Greater(op1, op2 Val) (Val, bool) {}

// GreaterOrEqual
func GreaterOrEqual(op1, op2 Val) (Val, bool) {}

// Identical
func Identical(op1, op2 Val) (Val, bool) {}

// LogicalAnd
func LogicalAnd(op1, op2 Val) (Val, bool) {}

// LogicalOr
func LogicalOr(op1, op2 Val) (Val, bool) {}

// LogicalXor
func LogicalXor(op1, op2 Val) (Val, bool) {}

// NotEqual
func NotEqual(op1, op2 Val) (Val, bool) {}

// NotIdentical
func NotIdentical(op1, op2 Val) (Val, bool) {}

// ShiftLeft
func ShiftLeft(op1, op2 Val) (Val, bool) {}

// ShiftRight
func ShiftRight(op1, op2 Val) (Val, bool) {}

// Smaller
func Smaller(op1, op2 Val) (Val, bool) {}

// SmallerOrEqual
func SmallerOrEqual(op1, op2 Val) (Val, bool) {}

// Spaceship
func Spaceship(op1, op2 Val) (Val, bool) {}

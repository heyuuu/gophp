package operators

import (
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"math"
)

var _ iOperator = (*Operator)(nil)

type Operator struct {
}

// Add
func (op *Operator) Add(op1, op2 Val) Val {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return op.AddLong(op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) + fastGetDouble(op2))
	case IsArrayArray:
		retArr := op.AddArray(op1.Array(), op2.Array())
		return Array(retArr)
	default:
		if !converted {
			// fail
			op.throwErrorNoReturn(nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = op.opScalarGetNumber(op1, op2)
		goto again
	}
}

func (op *Operator) AddLong(i1, i2 int) Val {
	if sign(i1) == sign(i2) && sign(i1) != sign(i1+i2) { // 判断相加是否越界
		return Double(float64(i1) + float64(i2))
	} else {
		return Long(i1 + i2)
	}
}

func (op *Operator) AddArray(a1, a2 *types.Array) *types.Array {
	// todo AddArray
	panic(perr.Unreachable())
}

// Sub
func (op *Operator) Sub(op1, op2 Val) Val {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return op.SubLong(op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) - fastGetDouble(op2))
	default:
		if converted {
			// fail
			op.throwErrorNoReturn(nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = op.opScalarGetNumber(op1, op2)
		goto again
	}
}

func (op *Operator) SubLong(i1, i2 int) Val {
	if sign(i1) != sign(i2) && sign(i1) != sign(i1-i2) { // 判断是否越界
		return Double(float64(i1) - float64(i2))
	} else {
		return Long(i1 - i2)
	}
}

// Mul
func (op *Operator) Mul(op1, op2 Val) Val {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return op.MulLong(op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) * fastGetDouble(op2))
	default:
		if converted {
			// fail
			op.throwErrorNoReturn(nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = op.opScalarGetNumber(op1, op2)
		goto again
	}
}

func (op *Operator) MulLong(i1, i2 int) Val {
	if iVal, dVal, overflow := op.mulLong(i1, i2); !overflow {
		return Long(iVal)
	} else {
		return Double(dVal)
	}
}

func (op *Operator) mulLong(i1, i2 int) (iVal int, dVal float64, overflow bool) {
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
func (op *Operator) Div(op1, op2 Val) Val {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return op.DivLong(op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		if d2 == 0 {
			op.error(faults.E_WARNING, "Division by zero")
			return Double(math.Inf(int(d1)))
		}
		return Double(d1 / d2)
	default:
		if converted {
			// fail
			op.throwErrorNoReturn(nil, "Unsupported operand types")
		}

		// convert
		converted = true
		op1, op2 = op.opScalarGetNumber(op1, op2)
		goto again
	}
}

func (op *Operator) DivLong(i1, i2 int) Val {
	if i2 == 0 {
		op.error(faults.E_WARNING, "Division by zero")
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
func (op *Operator) Mod(op1, op2 Val) Val {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = op.zvalGetLongNoisy(op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = op.zvalGetLongNoisy(op2)
	}

	if op2Lval == 0 {
		/* modulus by zero */
		op.throwIfExecutingNoReturn(nil, "Modulo by zero")
	}

	return Long(op1Lval % op2Lval)
}

// ShiftLeft (SL)
func (op *Operator) SL(op1, op2 Val) Val {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = op.zvalGetLongNoisy(op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = op.zvalGetLongNoisy(op2)
	}

	if op2Lval < 0 {
		op.throwIfExecutingNoReturn(nil, "Bit shift by negative number")
	}

	return Long(op1Lval << op2Lval)
}

// ShiftRight (SR)
func (op *Operator) SR(op1, op2 Val) Val {
	var op1Lval int
	var op2Lval int

	op1 = op1.DeRef()
	if op1.IsLong() {
		op1Lval = op1.Long()
	} else {
		op1Lval = op.zvalGetLongNoisy(op1)
	}

	op2 = op2.DeRef()
	if op2.IsLong() {
		op2Lval = op2.Long()
	} else {
		op2Lval = op.zvalGetLongNoisy(op2)
	}

	if op2Lval < 0 {
		op.throwIfExecutingNoReturn(nil, "Bit shift by negative number")
	}

	return Long(op1Lval >> op2Lval)
}

// Concat
func (op *Operator) Concat(op1, op2 Val) Val {
	var s1, s2 string

	op1 = op1.DeRef()
	op2 = op2.DeRef()

	if op1.IsString() {
		s1 = op1.String()
	} else {
		s1 = op.ZvalGetStrVal(op1)
	}

	if op2.IsString() {
		s2 = op2.String()
	} else {
		s2 = op.ZvalGetStrVal(op2)
	}

	if len(s1)+len(s2) > math.MaxInt {
		op.throwErrorNoReturn(nil, "String size overflow")
	}

	return String(s1 + s2)
}

// Pow
func (op *Operator) Pow(op1, op2 Val) Val {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	converted := false
again:
	switch typePair(op1, op2) {
	case IsLongLong:
		return op.PowLong(op1.Long(), op2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		d1 := fastGetDouble(op1)
		d2 := fastGetDouble(op2)
		return Double(math.Pow(d1, d2))
	default:
		if converted {
			// fail
			op.throwErrorNoReturn(nil, "Unsupported operand types")
		}

		// array type
		if op1.IsArray() {
			return Long(0)
		} else if op2.IsArray() {
			return Long(1)
		}

		// convert
		op1, op2 = op.opScalarGetNumber(op1, op2)

		converted = true
		goto again
	}
}

func (op *Operator) PowLong(i1, i2 int) Val {
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
				l1, _, overflow = op.mulLong(l1, l2)
			} else {
				i2 /= 2
				l2, _, overflow = op.mulLong(l2, l2)
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
func (op *Operator) BitwiseAnd(op1, op2 Val) Val {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() & op2.Long())
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
		return String(string(str))
	}

	// common
	var op1Lval = op.zvalGetLongNoisy(op1)
	var op2Lval = op.zvalGetLongNoisy(op2)
	return Long(op1Lval & op2Lval)
}

// BitwiseOr
func (op *Operator) BitwiseOr(op1, op2 Val) Val {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() | op2.Long())
	}

	// common
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, lang.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] | s2[i]
		}
		return String(string(str))
	}

	var op1Lval = op.zvalGetLongNoisy(op1)
	var op2Lval = op.zvalGetLongNoisy(op2)
	return Long(op1Lval | op2Lval)
}

// BitwiseXor
func (op *Operator) BitwiseXor(op1, op2 Val) Val {
	// fast
	if op1.IsLong() && op2.IsLong() {
		return Long(op1.Long() ^ op2.Long())
	}

	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if op1.IsString() && op2.IsString() {
		s1, s2 := op1.String(), op2.String()
		str := make([]byte, lang.Min(len(s1), len(s2)))
		for i := range str {
			str[i] = s1[i] & s2[i]
		}
		return String(string(str))
	}

	var op1Lval = op.zvalGetLongNoisy(op1)
	var op2Lval = op.zvalGetLongNoisy(op2)
	return Long(op1Lval ^ op2Lval)
}

// Coalesce(??)
func (op *Operator) Coalesce(op1 Val, op2 LazyVal) Val {
	if !op1.IsUndef() && !op1.IsNull() {
		return op1
	}
	return op2()
}

// BooleanAnd
func (op *Operator) BooleanAnd(op1 Val, op2 LazyVal) Val {
	op1Val := ZvalIsTrue(op1)
	if !op1Val {
		return False()
	}

	op2Val := ZvalIsTrue(op2())
	return Bool(op2Val)
}

// BooleanOr
func (op *Operator) BooleanOr(op1 Val, op2 LazyVal) Val {
	op1Val := ZvalIsTrue(op1)
	if op1Val {
		return True()
	}

	op2Val := ZvalIsTrue(op2())
	return Bool(op2Val)
}

// BooleanXor
func (op *Operator) BooleanXor(op1, op2 Val) Val {
	op1Val := ZvalIsTrue(op1)
	op2Val := ZvalIsTrue(op2)
	return Bool(lang.Xor(op1Val, op2Val))
}

// Identical
func (op *Operator) Identical(op1, op2 Val) Val {
	return Bool(ZvalIsIdentical(op1, op2))
}

// NotIdentical
func (op *Operator) NotIdentical(op1, op2 Val) Val {
	return Bool(!ZvalIsIdentical(op1, op2))
}

// Equal
func (op *Operator) Equal(op1, op2 Val) Val {
	result := op.Equals(op1, op2)
	return Bool(result)
}

// NotEqual
func (op *Operator) NotEqual(op1, op2 Val) Val {
	result := op.Equals(op1, op2)
	return Bool(!result)
}

// Greater
func (op *Operator) Greater(op1, op2 Val) Val {
	result := op.Compare(op1, op2)
	return Bool(result > 0)
}

// GreaterOrEqual
func (op *Operator) GreaterOrEqual(op1, op2 Val) Val {
	result := op.Compare(op1, op2)
	return Bool(result >= 0)
}

// Smaller
func (op *Operator) Smaller(op1, op2 Val) Val {
	result := op.Compare(op1, op2)
	return Bool(result < 0)
}

// SmallerOrEqual
func (op *Operator) SmallerOrEqual(op1, op2 Val) Val {
	result := op.Compare(op1, op2)
	return Bool(result <= 0)
}

// Spaceship
func (op *Operator) Spaceship(op1, op2 Val) Val {
	result := op.Compare(op1, op2)
	return Long(result)
}

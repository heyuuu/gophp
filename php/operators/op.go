package operators

import "github.com/heyuuu/gophp/php/types"

func Add(op1, op2 Val) (Val, bool) {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if ret, ok := AddFast(op1, op2); ok {
		return ret, true
	} else {
		return addSlow(op1, op2)
	}
}

func AddFast(op1, op2 Val) (Val, bool) {
	switch typePair(op1, op2) {
	case IsLongLong:
		return AddLong(op1.Long(), op2.Long()), true
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(op1) + fastGetDouble(op2)), true
	case IsArrayArray:
		retArr := AddArray(op1.Array(), op2.Array())
		return Array(retArr), true
	default:
		return nil, false
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

func addSlow(op1, op2 Val) (Val, bool) {
	// todo addSlow
	return nil, false
}

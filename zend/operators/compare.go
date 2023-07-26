package operators

import "github.com/heyuuu/gophp/php/types"

func compareResult[T int | float64](r T) int {
	if r < 0 {
		return -1
	} else if r > 0 {
		return 1
	} else {
		return 0
	}
}

func Compare(op1, op2 *types.Zval) int {
	switch TypePair(op1.Type(), op2.Type()) {
	case TypeLongLong:
		return compareResult(op1.Long() - op2.Long())
	case TypeDoubleDouble:
		if op1.Double() == op2.Double() {
			return 0
		}
		return compareResult(op1.Double() - op2.Double())
	case TypeLongDouble, TypeDoubleLong:
		return compareResult(_zvalFastGetDouble(op1) - _zvalFastGetDouble(op2))
	case TypeArrayArray:
		return CompareArrays(op1.Array(), op2.Array())
		// todo
	}
}

func CompareArrays(arr1, arr2 *types.Array) int {
	if arr1 == arr2 {
		return 0
	}

	// todo
}

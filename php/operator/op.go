package operator

import "github.com/heyuuu/gophp/php/types"

func fastGetDouble(v Val) float64 {
	if v.IsLong() {
		return float64(v.Long())
	} else if v.IsDouble() {
		return v.Double()
	} else {
		return 0
	}
}

func Add(v1, v2 Val) Val {
	// todo deref
	switch typePair(v1, v2) {
	case IsLongLong:
		return AddLong(v1.Long(), v2.Long())
	case IsLongDouble, IsDoubleLong, IsDoubleDouble:
		return Double(fastGetDouble(v1) + fastGetDouble(v2))
	default:
		// todo
		panic("unreachable")
	}
}

func AddLong(i1, i2 int) types.Zval {
	if sign(i1) == sign(i2) && sign(i1) != sign(i1+i2) { // 判断相加是否越界
		return Double(float64(i1) + float64(i2))
	} else {
		return Long()
	}
}

func AddArray(a1, a2 *types.Array) *types.Array {
	panic("unreachable")
}

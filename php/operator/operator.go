package operator

import "github.com/heyuuu/gophp/php/types"

type Operator interface {
	Add(v1, v2 *types.Zval) *types.Zval
	Sub(v1, v2 *types.Zval) *types.Zval
	Mul(v1, v2 *types.Zval) *types.Zval
	Div(v1, v2 *types.Zval) *types.Zval
	Mod(v1, v2 *types.Zval) *types.Zval
	ShiftLeft(v1, v2 *types.Zval) *types.Zval
	ShiftRight(v1, v2 *types.Zval) *types.Zval
	Concat(v1, v2 *types.Zval) *types.Zval
}

func DefaultOperator() Operator {
	return &baseOperator{}
}

// baseOperator extends Operator 默认实现
type baseOperator struct{}

var _ Operator = (*baseOperator)(nil)

func (op *baseOperator) Add(v1, v2 *types.Zval) *types.Zval {
	switch typePair(v1, v2) {
	case IsLongLong:
		l1, l2 := v1.Long(), v2.Long()
		if sign(l1) == sign(l2) && sign(l1) != sign(l1+l2) { // 判断相加是否越界
			return Double(float64(l1) + float64(l2))
		} else {
			return Long(l1 + l2)
		}
	case IsLongDouble:
		return Double(float64(v1.Long()) + v2.Double())
	case IsDoubleLong:
		return Double(v1.Double() + float64(v2.Long()))
	case IsDoubleDouble:
		return Double(v1.Double() + v2.Double())
	case IsArrayArray:
		// todo
		panic("implement me")
	default:
		// todo
		panic("implement me")
	}
}

func (op *baseOperator) Sub(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) Mul(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) Div(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) Mod(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) ShiftLeft(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) ShiftRight(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

func (op *baseOperator) Concat(v1, v2 *types.Zval) *types.Zval {
	//TODO implement me
	panic("implement me")
}

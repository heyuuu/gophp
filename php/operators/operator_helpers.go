package operators

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func (op *Operator) opScalarGetNumber(v1, v2 Val) (Val, Val) {
	v1, v2 = op.opScalarGetNumberEx(v1, v2, false)
	if op.HasException() {
		//return nil, false
		panic(perr.Unreachable())
	}
	return v1, v2
}

func (op *Operator) opScalarGetNumberEx(v1, v2 Val, silent bool) (Val, Val) {
	if v1 != v2 {
		v1 = op.ToNumberEx(v1, silent)
		v2 = op.ToNumberEx(v2, silent)
	} else {
		v1 = op.ToNumberEx(v1, silent)
		v2 = v1
	}
	return v1, v2
}

func (op *Operator) opObjectCompare(obj *types.Object, v1, v2 Val) (result int, ok bool) {
	if v1.Object() == v2.Object() {
		return 0, true
	}
	// todo 暂未支持 object 比较
	return 0, false
}

func (op *Operator) parseNumberPrefix(str string, silent bool) Val {
	zv, matchLen := ParseNumberPrefix(str)
	if matchLen != len(str) && !silent {
		// notice: 此处可能会触发 Exception
		op.Error(perr.E_NOTICE, "A non well formed numeric value encountered")
		if op.HasException() {
			return nil
		}
	}
	return zv
}

package operators

import (
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func (op *Operator) ZvalToArrayKey(offset Val) types.ArrayKey {
	// todo
	return types.IdxKey(0)
}

func (op *Operator) error(typ int, message string) {
	faults.Error(typ, message)
}

func (op *Operator) throwErrorNoReturn(exceptionCe *types.Class, message string) {
	// fail
	faults.ThrowError(exceptionCe, message)
	panic(perr.Unreachable())
}

func (op *Operator) throwIfExecutingNoReturn(exceptionCe *types.Class, message string) {
	// todo
	throwIfExecuting(exceptionCe, message)
	panic(perr.Unreachable())
}

func (op *Operator) opScalarGetNumber(op1, op2 Val) (Val, Val) {
	op1, op2 = opScalarGetNumber(op1, op2)
	if hasException() {
		//return nil, false
		panic(perr.Unreachable())
	}
	return op1, op2
}

func (op *Operator) zvalGetLongNoisy(v Val) int {
	ret := zvalGetLongNoisy(v)
	if hasException() {
		//return nil, false
		panic(perr.Unreachable())
	}
	return ret
}

func (op *Operator) ZvalGetStrVal(v Val) string {
	ret := ZvalGetStrVal(v)
	if hasException() {
		//return nil, false
		panic(perr.Unreachable())
	}
	return ret
}

func (op *Operator) Equals(op1, op2 Val) bool {
	result, ok := ZvalEquals(op1, op2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}

func (op *Operator) Compare(op1, op2 Val) int {
	result, ok := ZvalCompare(op1, op2)
	if !ok {
		// todo
		panic(perr.Unreachable())
	}
	return result
}

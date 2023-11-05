package operators

import (
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/types"
)

func (op *Operator) throwErrorNoReturn(exceptionCe *types.Class, message string) {
	// fail
	faults.ThrowError(exceptionCe, message)
	//return nil, false
	panic("unreachable")
}

func (op *Operator) throwIfExecuting(exceptionCe *types.Class, message string) {
	// todo
	throwIfExecuting(exceptionCe, message)
	panic("unreachable")
}

func (op *Operator) opScalarGetNumber(op1, op2 Val) (Val, Val) {
	op1, op2 = opScalarGetNumber(op1, op2)
	if hasException() {
		//return nil, false
		panic("unreachable")
	}
	return op1, op2
}

func (op *Operator) zvalGetLongNoisy(v Val) int {
	ret := zvalGetLongNoisy(v)
	if hasException() {
		//return nil, false
		panic("unreachable")
	}
	return ret
}

func (op *Operator) ZvalGetStrVal(v Val) string {
	ret := ZvalGetStrVal(v)
	if hasException() {
		//return nil, false
		panic("unreachable")
	}
	return ret
}

func (op *Operator) Equals(op1, op2 Val) bool {
	result, ok := ZvalEquals(op1, op2)
	if !ok {
		// todo
		panic("unreachable")
	}
	return result
}

func (op *Operator) Compare(op1, op2 Val) int {
	result, ok := ZvalCompare(op1, op2)
	if !ok {
		// todo
		panic("unreachable")
	}
	return result
}

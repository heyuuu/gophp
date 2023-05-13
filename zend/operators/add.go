package operators

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func AddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	return types.IntBool(AddFunctionEx(result, op1, op2))
}

func AddFunctionEx(result *types.Zval, op1 *types.Zval, op2 *types.Zval) bool {
	op1 = op1.DeRef()
	op2 = op2.DeRef()
	if _addFunctionFast(result, op1, op2) {
		return true
	} else if _addFunctionSlow(result, op1, op2) {
		return true
	} else {
		return false
	}
}

func _addFunctionFast(result *types.Zval, op1 *types.Zval, op2 *types.Zval) bool {
	switch TypePair(op1.GetType(), op1.GetType()) {
	case TypeLongLong:
		FastLongAddFunction(result, op1, op2)
		return true
	case TypeLongDouble, TypeDoubleLong, TypeDoubleDouble:
		result.SetDouble(_zvalFastGetDouble(op1) + _zvalFastGetDouble(op2))
		return true
	case TypeArrayArray:
		AddFunctionArray(result, op1, op2)
		return true
	default:
		return false
	}
}
func _addFunctionSlow(result *types.Zval, op1 *types.Zval, op2 *types.Zval) bool {
	var op1Copy types.Zval
	var op2Copy types.Zval

	// convert
	if op1.IsObject() && op1 == result && op1.Object().Handlers().GetGet() != nil && op1.Object().Handlers().GetSet() != nil {
		var rv types.Zval
		var objval *types.Zval = op1.Object().Handlers().GetGet()(op1, &rv)
		ret := AddFunctionEx(objval, objval, op2)
		op1.Object().Handlers().GetSet()(op1, objval)
		return ret
	} else if op1.IsObject() && op1.Object().Handlers().GetDoOperation() != nil {
		if types.SUCCESS == op1.Object().Handlers().GetDoOperation()(zend.ZEND_ADD, result, op1, op2) {
			return true
		}
	} else if op2.IsObject() && op2.Object().Handlers().GetDoOperation() != nil && types.SUCCESS == op2.Object().Handlers().GetDoOperation()(zend.ZEND_ADD, result, op1, op2) {
		return true
	}
	if op1 != op2 {
		op1 = ZendiConvertScalarToNumber(op1, &op1Copy, result, 0)
		op2 = ZendiConvertScalarToNumber(op2, &op2Copy, result, 0)
	} else {
		op1 = ZendiConvertScalarToNumber(op1, &op1Copy, result, 0)
		op2 = op1
	}
	if zend.EG__().GetException() != nil {
		if result != op1 {
			result.SetUndef()
		}
		return false
	}

	// try again
	if _addFunctionFast(result, op1, op2) {
		return true
	}

	// fail
	if result != op1 {
		result.SetUndef()
	}
	faults.ThrowError(nil, "Unsupported operand types")
	return false
}

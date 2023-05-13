package operators

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func SubFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if _subFunctionFast(result, op1, op2) == types.SUCCESS {
		return types.SUCCESS
	} else {
		return _subFunctionSlow(result, op1, op2)
	}
}

func _subFunctionFast(result *types.Zval, op1 *types.Zval, op2 *types.Zval) bool {
	switch TypePair(op1.GetType(), op2.GetType()) {
	case TypeLongLong:
		FastLongSubFunction(result, op1, op2)
		return true
	case TypeLongDouble, TypeDoubleLong, TypeDoubleDouble:
		result.SetDouble(_zvalFastGetDouble(op1) - _zvalFastGetDouble(op2))
		return true
	default:
		return false
	}
}
func _subFunctionSlow(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1Copy types.Zval
	var op2Copy types.Zval

	// convert
	if op1.IsObject() && op1 == result && op1.Object().Handlers().GetGet() != nil && op1.Object().Handlers().GetSet() != nil {
		var ret int
		var rv types.Zval
		var objval *types.Zval = op1.Object().Handlers().GetGet()(op1, &rv)
		ret = SubFunction(objval, objval, op2)
		op1.Object().Handlers().GetSet()(op1, objval)
		return ret
	} else if op1.IsObject() && op1.Object().Handlers().GetDoOperation() != nil {
		if types.SUCCESS == op1.Object().Handlers().GetDoOperation()(zend.ZEND_SUB, result, op1, op2) {
			return types.SUCCESS
		}
	} else if op2.IsObject() && op2.Object().Handlers().GetDoOperation() != nil && types.SUCCESS == op2.Object().Handlers().GetDoOperation()(zend.ZEND_SUB, result, op1, op2) {
		return types.SUCCESS
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
		return types.FAILURE
	}

	// try again
	if _subFunctionFast(result, op1, op2) {
		return types.SUCCESS
	}

	// fail
	if result != op1 {
		result.SetUndef()
	}
	faults.ThrowError(nil, "Unsupported operand types")
	return types.FAILURE
}

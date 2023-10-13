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
	switch TypePair(op1.Type(), op2.Type()) {
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
	if op1.IsObject() && op1 == result && op1.Object().CanGet() && op1.Object().CanSet() {
		var ret int
		var rv types.Zval
		var objval *types.Zval = op1.Object().Get(&rv)
		ret = SubFunction(objval, objval, op2)
		op1.Object().Set(objval)
		return ret
	} else if op1.IsObject() && op1.Object().CanDoOperation() {
		if types.SUCCESS == op1.Object().DoOperation(zend.ZEND_SUB, result, op1, op2) {
			return types.SUCCESS
		}
	} else if op2.IsObject() && op2.Object().CanDoOperation() && types.SUCCESS == op2.Object().DoOperation(zend.ZEND_SUB, result, op1, op2) {
		return types.SUCCESS
	}
	if op1 != op2 {
		op1 = ZendiConvertScalarToNumber(op1, &op1Copy, result, 0)
		op2 = ZendiConvertScalarToNumber(op2, &op2Copy, result, 0)
	} else {
		op1 = ZendiConvertScalarToNumber(op1, &op1Copy, result, 0)
		op2 = op1
	}
	if zend.EG__().HasException() {
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

package zend

import (
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_COUNT_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var count ZendLong
	op1 = executeData.GetOp1(opline)
	for true {
		if op1.IsArray() {
			count = op1.Array().Count()
			break
		} else if op1.IsObject() {

			/* first, we check if the handler is defined */

			if op1.Object().CanCountElements() {
				if c, ok := op1.Object().CountElements(); ok {
					count = c
					break
				}
				if EG__().GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if operators.InstanceofFunction(types.Z_OBJCE_P(op1), ZendCeCountable) != 0 {
				var retval types.Zval
				ZendCallMethodWith0Params(op1, nil, nil, "count", &retval)
				count = operators.ZvalGetLong(&retval)
				// ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if op1.Type() <= types.IsNull {
			count = 0
		} else {
			count = 1
		}
		faults.Error(faults.E_WARNING, "%s(): Parameter must be an array or an object that implements Countable", b.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	opline.Result().SetLong(count)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var count ZendLong
	op1 = opline.Op1()
	for true {
		if op1.IsArray() {
			count = op1.Array().Count()
			break
		} else if op1.IsObject() {

			/* first, we check if the handler is defined */

			if op1.Object().CanCountElements() {
				if c, ok := op1.Object().CountElements(); ok {
					count = c
					break
				}
				if EG__().GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if operators.InstanceofFunction(types.Z_OBJCE_P(op1), ZendCeCountable) != 0 {
				var retval types.Zval
				ZendCallMethodWith0Params(op1, nil, nil, "count", &retval)
				count = operators.ZvalGetLong(&retval)
				// ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if op1.IsRef() {
			op1 = types.Z_REFVAL_P(op1)
			continue
		} else if op1.Type() <= types.IsNull {
			if op1.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			count = 0
		} else {
			count = 1
		}
		faults.Error(faults.E_WARNING, "%s(): Parameter must be an array or an object that implements Countable", b.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	opline.Result().SetLong(count)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_COUNT_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var count ZendLong
	op1 = opline.Op1()
	for true {
		if op1.IsArray() {
			count = op1.Array().Count()
			break
		} else if op1.IsObject() {

			/* first, we check if the handler is defined */

			if op1.Object().CanCountElements() {
				if c, ok := op1.Object().CountElements(); ok {
					count = c
					break
				}
				if EG__().GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if operators.InstanceofFunction(types.Z_OBJCE_P(op1), ZendCeCountable) != 0 {
				var retval types.Zval
				ZendCallMethodWith0Params(op1, nil, nil, "count", &retval)
				count = operators.ZvalGetLong(&retval)
				// ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if op1.IsRef() {
			op1 = types.Z_REFVAL_P(op1)
			continue
		} else if op1.Type() <= types.IsNull {
			if op1.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			count = 0
		} else {
			count = 1
		}
		faults.Error(faults.E_WARNING, "%s(): Parameter must be an array or an object that implements Countable", b.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	opline.Result().SetLong(count)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

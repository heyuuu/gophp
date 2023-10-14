package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_JMP_SET_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = opline.Const1()
	ret = operators.IZendIsTrue(value)
	if EG__().HasException() {
		opline.Result().SetUndef()
		return 0
	}
	if ret != 0 {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)
		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_JMP_SET_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsRef() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	ret = operators.IZendIsTrue(value)
	if EG__().HasException() {
		// ZvalPtrDtorNogc(free_op1)
		opline.Result().SetUndef()
		return 0
	}
	if ret != 0 {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_JMP_SET_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = opline.Op1()
	if value.IsRef() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	ret = operators.IZendIsTrue(value)
	if EG__().HasException() {
		// ZvalPtrDtorNogc(free_op1)
		opline.Result().SetUndef()
		return 0
	}
	if ret != 0 {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_JMP_SET_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = opline.Cv1OrUndef()
	if value.IsRef() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	ret = operators.IZendIsTrue(value)
	if EG__().HasException() {
		opline.Result().SetUndef()
		return 0
	}
	if ret != 0 {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

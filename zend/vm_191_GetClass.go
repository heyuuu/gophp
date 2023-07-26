package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = executeData.GetOp1(opline)
		for true {
			if op1.IsObject() {
				opline.Result().SetString(types.Z_OBJCE_P(op1).Name())
			} else {
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.Type()))
				opline.Result().SetFalse()
			}
			break
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()

	{
		var free_op1 ZendFreeOp
		var op1 *types.Zval
		op1 = opline.Op1()
		for true {
			if op1.IsObject() {
				opline.Result().SetString(types.Z_OBJCE_P(op1).Name())
			} else if op1.IsRef() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1(executeData)
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.Type()))
				opline.Result().SetFalse()
			}
			break
		}
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	{
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_class() called without object from outside a class")
			opline.Result().SetFalse()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			opline.Result().SetString(executeData.GetFunc().common.scope.name.GetStr())
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		}
	}

}
func ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = opline.Op1()
		for true {
			if op1.IsObject() {
				opline.Result().SetString(types.Z_OBJCE_P(op1).Name())
			} else if op1.IsRef() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1(executeData)
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.Type()))
				opline.Result().SetFalse()
			}
			break
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}

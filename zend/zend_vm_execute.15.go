package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func zend_pre_inc_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1(executeData)
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.Reference()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, nil, opline, executeData)
				break
			}
		}
		operators.IncrementFunction(var_ptr)
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), var_ptr)
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_pre_dec_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1(executeData)
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.Reference()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, nil, opline, executeData)
				break
			}
		}
		operators.DecrementFunction(var_ptr)
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), var_ptr)
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_post_inc_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsError() {
		opline.Result().SetNull()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1(executeData)
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.Reference()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, opline.Result(), opline, executeData)
				break
			}
		}
		types.ZVAL_COPY(opline.Result(), var_ptr)
		operators.IncrementFunction(var_ptr)
		break
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_post_dec_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsError() {
		opline.Result().SetNull()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1(executeData)
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.Reference()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, opline.Result(), opline, executeData)
				break
			}
		}
		types.ZVAL_COPY(opline.Result(), var_ptr)
		operators.DecrementFunction(var_ptr)
		break
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	varptr = opline.Op1()
	if varptr.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

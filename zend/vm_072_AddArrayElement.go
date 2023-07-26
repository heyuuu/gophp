package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	expr_ptr = opline.Const1()
	// expr_ptr.TryAddRefcount()
	{
		var offset *types.Zval = opline.Const2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	{
		expr_ptr = opline.Const1()

		{
			// expr_ptr.TryAddRefcount()
		}

	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
		// ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	{
		expr_ptr = opline.Const1()

		{
			// expr_ptr.TryAddRefcount()
		}

	}

	{
		if opline.Result().Array().Append(expr_ptr) == nil {
			ZendCannotAddElement()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	{
		expr_ptr = opline.Const1()

		{
			// expr_ptr.TryAddRefcount()
		}

	}
	{
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = nil
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	}
	{
		var offset *types.Zval = opline.Const2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = nil
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
		// ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = nil
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	}

	{
		if opline.Result().Array().Append(expr_ptr) == nil {
			ZendCannotAddElement()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = nil
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	}
	{
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
		if free_op1 != nil {
			// ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = opline.Op1()
	}
	{
		var offset *types.Zval = opline.Const2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
		if free_op1 != nil {
			// ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = opline.Op1()
	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
		// ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
		if free_op1 != nil {
			// ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = opline.Op1()
	}

	{
		if opline.Result().Array().Append(expr_ptr) == nil {
			ZendCannotAddElement()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
		if free_op1 != nil {
			// ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = opline.Op1()
	}
	{
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = opline.Cv1OrUndef()
	}
	{
		var offset *types.Zval = opline.Const2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = opline.Cv1OrUndef()
	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
		// ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = opline.Cv1OrUndef()
	}

	{
		if opline.Result().Array().Append(expr_ptr) == nil {
			ZendCannotAddElement()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsRef() {
			// 			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = opline.Cv1OrUndef()
	}
	{
		var offset *types.Zval = opline.Op2()
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.Result().Array().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.Long()
		num_index:
			opline.Result().Array().IndexUpdate(hval, expr_ptr)
		} else if offset.IsRef() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.NewString("")
			goto str_index
		} else if offset.IsDouble() {
			hval = operators.DvalToLval(offset.Double())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = offset.ResourceHandle()
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			str = types.NewString("")
			goto str_index
		} else {
			ZendIllegalOffset()
			// ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

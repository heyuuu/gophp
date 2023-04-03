package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_VM_DISPATCH(executeData *ZendExecuteData, opcode OpCode, opline *ZendOp) int {
	return ZendVmGetOpcodeHandler(opcode, opline)(executeData)
}
func zend_mul_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	MulFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_mod_by_zero_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Modulo by zero")
	opline.GetResultZval().SetUndef()
	return 0
}
func zend_mod_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	ModFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_shift_left_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	ShiftLeftFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_shift_right_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	ShiftRightFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_is_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	CompareFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG__().GetException() != nil {
		return 0
	}
	if opline.GetResultZval().GetLval() == 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.GetResultZval().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_not_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	CompareFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG__().GetException() != nil {
		return 0
	}
	if opline.GetResultZval().GetLval() != 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.GetResultZval().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_smaller_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	CompareFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG__().GetException() != nil {
		return 0
	}
	if opline.GetResultZval().GetLval() < 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.GetResultZval().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_smaller_or_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	CompareFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG__().GetException() != nil {
		return 0
	}
	if opline.GetResultZval().GetLval() <= 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.GetResultZval().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_bw_or_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	BitwiseOrFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_bw_and_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	BitwiseAndFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_bw_xor_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	BitwiseXorFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_this_not_in_object_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Using $this when not in object context")
	if (opline + 1).GetOpcode() == ZEND_OP_DATA {
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
	}
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	UNDEF_RESULT()
	return 0
}
func zend_undefined_function_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	function_name = RT_CONSTANT(opline, opline.GetOp2())
	faults.ThrowError(nil, "Call to undefined function %s()", function_name.GetStr().GetVal())
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_OP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	/* This helper actually never will receive IS_VAR as second op, and has the same handling for VAR and TMP in the first op, but for interoperability with the other binary_assign_op helpers, it is necessary to "include" it */

	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	var ref *types.ZendReference
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, (opline+1).GetExtendedValue(), BP_VAR_RW, 0, opline, executeData) != types.SUCCESS {
		b.Assert(EG__().GetException() != nil)
		UNDEF_RESULT()
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
		return 0
	}
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
	for {
		if prop.IsReference() {
			ref = prop.GetRef()
			prop = types.Z_REFVAL_P(prop)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
				break
			}
		}
		if prop_info.GetType() != 0 {

			/* special case for typed properties */

			ZendBinaryAssignOpTypedProp(prop_info, prop, value, opline, executeData)

			/* special case for typed properties */

		} else {
			ZendBinaryOp(prop, prop, value, opline)
		}
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), prop)
	}
	FREE_OP(free_op_data)

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_PRE_INC_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var prop *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_RW, 0, opline, executeData) != types.SUCCESS {
		UNDEF_RESULT()
		return 0
	}
	ZendPreIncdecPropertyZval(prop, b.Cond(prop_info.GetType() != 0, prop_info, nil), opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POST_INC_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var prop *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_RW, 0, opline, executeData) != types.SUCCESS {
		UNDEF_RESULT()
		return 0
	}
	ZendPostIncdecPropertyZval(prop, b.Cond(prop_info.GetType() != 0, prop_info, nil), opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_fetch_static_prop_helper_SPEC(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var prop *types.Zval
	if ZendFetchStaticPropertyAddress(&prop, nil, opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS, type_, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, opline, executeData) != types.SUCCESS {
		b.Assert(EG__().GetException() != nil || type_ == BP_VAR_IS)
		prop = EG__().GetUninitializedZval()
	}
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), prop)
	} else {
		opline.GetResultZval().SetIndirect(prop)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_STATIC_PROP_R_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_R, executeData)
}
func ZEND_FETCH_STATIC_PROP_W_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_W, executeData)
}
func ZEND_FETCH_STATIC_PROP_RW_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_RW, executeData)
}
func ZEND_FETCH_STATIC_PROP_FUNC_ARG_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_static_prop_helper_SPEC(fetch_type, executeData)
}
func ZEND_FETCH_STATIC_PROP_UNSET_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_UNSET, executeData)
}
func ZEND_FETCH_STATIC_PROP_IS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_IS, executeData)
}
func zend_use_tmp_in_write_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot use temporary expression in write context")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	opline.GetResultZval().SetUndef()
	return 0
}
func zend_use_undef_in_read_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot use [] for reading")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	opline.GetResultZval().SetUndef()
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		UNDEF_RESULT()
		return 0
	}
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
	} else {
		value = ZendAssignToVariable(prop, value, IS_CONST, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
		UNDEF_RESULT()
		return 0
	}
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
		ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
		UNDEF_RESULT()
		return 0
	}
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
		ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, IS_VAR, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		UNDEF_RESULT()
		return 0
	}
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
	} else {
		value = ZendAssignToVariable(prop, value, IS_CV, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_REF_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value_ptr *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION, BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
		UNDEF_RESULT()
		return 0
	}
	value_ptr = GetZvalPtrPtr((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, BP_VAR_W)
	if (opline+1).GetOp1Type() == IS_VAR && value_ptr.IsError() {
		prop = EG__().GetUninitializedZval()
	} else if (opline+1).GetOp1Type() == IS_VAR && (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && !(value_ptr.IsReference()) {
		if ZendWrongAssignToVariableReference(prop, value_ptr, opline, executeData) == nil {
			prop = EG__().GetUninitializedZval()
		}
	} else if prop_info.GetType() != 0 {
		prop = ZendAssignToTypedPropertyReference(prop_info, prop, value_ptr, executeData)
	} else {
		ZendAssignToVariableReference(prop, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), prop)
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func zend_leave_helper_SPEC(executeData *ZendExecuteData) int {
	var old_execute_data *ZendExecuteData
	var call_info uint32 = EX_CALL_INFO()
	if (call_info & (ZEND_CALL_CODE | ZEND_CALL_TOP | ZEND_CALL_HAS_SYMBOL_TABLE | ZEND_CALL_FREE_EXTRA_ARGS | ZEND_CALL_ALLOCATED)) == 0 {
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		IFreeCompiledVariables(executeData)
		if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
			OBJ_RELEASE(executeData.GetThis().GetObj())
		} else if (call_info & ZEND_CALL_CLOSURE) != 0 {
			OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
		}
		EG__().SetVmStackTop((*types.Zval)(executeData))
		executeData = executeData.GetPrevExecuteData()
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else if (call_info & (ZEND_CALL_CODE | ZEND_CALL_TOP)) == 0 {
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		IFreeCompiledVariables(executeData)
		if (call_info & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			ZendCleanAndCacheSymbolTable(executeData.GetSymbolTable(

				/* Free extra args before releasing the closure,
				 * as that may free the op_array. */))
		}

		ZendVmStackFreeExtraArgsEx(call_info, executeData)
		if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
			OBJ_RELEASE(executeData.GetThis().GetObj())
		} else if (call_info & ZEND_CALL_CLOSURE) != 0 {
			OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
		}
		old_execute_data = executeData
		executeData = executeData.GetPrevExecuteData()
		ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else if (call_info & ZEND_CALL_TOP) == 0 {
		ZendDetachSymbolTable(executeData)
		DestroyOpArray(executeData.GetFunc().op_array)
		EfreeSize(executeData.GetFunc(), b.SizeOf("zend_op_array"))
		old_execute_data = executeData
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		executeData = CurrEX()
		ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
		ZendAttachSymbolTable(executeData)
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else {
		if (call_info & ZEND_CALL_CODE) == 0 {
			EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
			IFreeCompiledVariables(executeData)
			if (call_info & (ZEND_CALL_HAS_SYMBOL_TABLE | ZEND_CALL_FREE_EXTRA_ARGS)) != 0 {
				if (call_info & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
					ZendCleanAndCacheSymbolTable(executeData.GetSymbolTable())
				}
				ZendVmStackFreeExtraArgsEx(call_info, executeData)
			}
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
			}
			return -1
		} else {
			var symbol_table *types.Array = executeData.GetSymbolTable()
			ZendDetachSymbolTable(executeData)
			old_execute_data = executeData.GetPrevExecuteData()
			for old_execute_data != nil {
				if old_execute_data.GetFunc() != nil && (ZEND_CALL_INFO(old_execute_data)&ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
					if old_execute_data.GetSymbolTable() == symbol_table {
						ZendAttachSymbolTable(old_execute_data)
					}
					break
				}
				old_execute_data = old_execute_data.GetPrevExecuteData()
			}
			EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
			return -1
		}
	}
}
func ZEND_JMP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
func ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	call.SetPrevExecuteData(executeData)
	EG__().SetCurrentExecuteData(call)
	ret = &retval
	ret.SetNull()
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG__().SetCurrentExecuteData(executeData)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)
	IZvalPtrDtor(ret)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	call.SetPrevExecuteData(executeData)
	EG__().SetCurrentExecuteData(call)
	ret = opline.GetResultZval()
	ret.SetNull()
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG__().SetCurrentExecuteData(executeData)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)

	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_UCALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	ret = nil

	call.SetPrevExecuteData(executeData)
	executeData = call
	IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
	return 1
}
func ZEND_DO_UCALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	ret = nil
	ret = opline.GetResultZval()
	call.SetPrevExecuteData(executeData)
	executeData = call
	IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
	return 1
}
func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil

		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		return 1
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		if fbc.IsDeprecated() {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				UNDEF_RESULT()
				ret = &retval
				ret.SetUndef()
				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			UNDEF_RESULT()
			ret = &retval
			ret.SetUndef()
			goto fcall_by_name_end
		}
		ret = &retval
		ret.SetNull()
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG__().SetCurrentExecuteData(executeData)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)
		IZvalPtrDtor(ret)
	}
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil
		ret = opline.GetResultZval()
		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		return 1
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		if fbc.IsDeprecated() {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				UNDEF_RESULT()

				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			UNDEF_RESULT()

			goto fcall_by_name_end
		}
		ret = opline.GetResultZval()
		ret.SetNull()
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG__().SetCurrentExecuteData(executeData)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)

	}
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_FCALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.HasFnFlags(AccAbstract | AccDeprecated) {
		if fbc.IsAbstract() {
			ZendAbstractMethod(fbc)
		fcall_except:
			UNDEF_RESULT()
			ret = &retval
			ret.SetUndef()
			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil

		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 1, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < ZEND_USER_FUNCTION {
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = &retval
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(executeData)
	fcall_end:
		ZendVmStackFreeArgs(call)
		IZvalPtrDtor(ret)
	} else {
		ret = &retval
		call.SetPrevExecuteData(executeData)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			UNDEF_RESULT()
			return 0
		}
		ZvalPtrDtor(ret)
	}
	if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
		OBJ_RELEASE(call.GetThis().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_FCALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.HasFnFlags(AccAbstract | AccDeprecated) {
		if fbc.IsAbstract() {
			ZendAbstractMethod(fbc)
		fcall_except:
			UNDEF_RESULT()

			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil
		ret = opline.GetResultZval()
		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 1, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < ZEND_USER_FUNCTION {
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = opline.GetResultZval()
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(executeData)
	fcall_end:
		ZendVmStackFreeArgs(call)

	} else {
		ret = opline.GetResultZval()
		call.SetPrevExecuteData(executeData)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			UNDEF_RESULT()
			return 0
		}

	}
	if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
		OBJ_RELEASE(call.GetThis().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_GENERATOR_CREATE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var return_value *types.Zval = executeData.GetReturnValue()
	if return_value != nil {
		var opline *ZendOp = executeData.GetOpline()
		var generator *ZendGenerator
		var gen_execute_data *ZendExecuteData
		var num_args uint32
		var used_stack uint32
		var call_info uint32
		ObjectInitEx(return_value, ZendCeGenerator)

		/*
		 * Normally the executeData is allocated on the VM stack (because it does
		 * not actually do any allocation and thus is faster). For generators
		 * though this behavior would be suboptimal, because the (rather large)
		 * structure would have to be copied back and forth every time execution is
		 * suspended or resumed. That's why for generators the execution context
		 * is allocated on heap.
		 */

		num_args = executeData.NumArgs()
		if num_args <= executeData.GetFunc().op_array.num_args {
			used_stack = (ZEND_CALL_FRAME_SLOT + executeData.GetFunc().op_array.last_var + executeData.GetFunc().op_array.T) * b.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(Emalloc(used_stack))
			used_stack = (ZEND_CALL_FRAME_SLOT + executeData.GetFunc().op_array.last_var) * b.SizeOf("zval")
		} else {
			used_stack = (ZEND_CALL_FRAME_SLOT + num_args + executeData.GetFunc().op_array.last_var + executeData.GetFunc().op_array.T - executeData.GetFunc().op_array.num_args) * b.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(Emalloc(used_stack))
		}
		memcpy(gen_execute_data, executeData, used_stack)

		/* Save execution context in generator object. */

		generator = (*ZendGenerator)(executeData.GetReturnValue().GetObj())
		generator.SetExecuteData(gen_execute_data)
		generator.SetFrozenCallStack(nil)
		generator.GetExecuteFake().SetOpline(nil)
		generator.GetExecuteFake().SetFunc(nil)
		generator.GetExecuteFake().SetPrevExecuteData(nil)
		generator.GetExecuteFake().GetThis().SetObject((*types.ZendObject)(generator))
		gen_execute_data.SetOpline(opline + 1)

		/* EX(return_value) keeps pointer to zend_object (not a real zval) */

		gen_execute_data.SetReturnValue((*types.Zval)(generator))
		call_info = executeData.GetThis().GetTypeInfo()
		if (call_info&types.Z_TYPE_MASK) == types.IS_OBJECT && ((call_info&(ZEND_CALL_CLOSURE|ZEND_CALL_RELEASE_THIS)) == 0 || ZendExecuteEx != ExecuteEx) {
			ZEND_ADD_CALL_FLAG_EX(call_info, ZEND_CALL_RELEASE_THIS)
			gen_execute_data.GetThis().AddRefcount()
		}
		ZEND_ADD_CALL_FLAG_EX(call_info, ZEND_CALL_TOP_FUNCTION|ZEND_CALL_ALLOCATED|ZEND_CALL_GENERATOR)
		gen_execute_data.GetThis().GetTypeInfo() = call_info
		gen_execute_data.SetPrevExecuteData(nil)
		call_info = EX_CALL_INFO()
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		if (call_info & (ZEND_CALL_TOP | ZEND_CALL_ALLOCATED)) == 0 {
			EG__().SetVmStackTop((*types.Zval)(executeData))
			executeData = executeData.GetPrevExecuteData()
			ZEND_VM_INC_OPCODE(executeData)
			return 2
		} else if (call_info & ZEND_CALL_TOP) == 0 {
			var old_execute_data *ZendExecuteData = executeData
			executeData = executeData.GetPrevExecuteData()
			ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
			ZEND_VM_INC_OPCODE(executeData)
			return 2
		} else {
			return -1
		}
	} else {
		return zend_leave_helper_SPEC(executeData)
	}
}

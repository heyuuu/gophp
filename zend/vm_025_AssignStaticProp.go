package zend

import "github.com/heyuuu/gophp/php/types"

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
		types.ZVAL_COPY(opline.Result(), value)
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
		// ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
		UNDEF_RESULT()
		return 0
	}
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
		// ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
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
		// ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
		UNDEF_RESULT()
		return 0
	}
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
		// ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, IS_VAR, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
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
	value = opline.Offset(1).Cv1OrUndef()
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, executeData)
	} else {
		value = ZendAssignToVariable(prop, value, IS_CV, executeData.IsCallUseStrictTypes())
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}

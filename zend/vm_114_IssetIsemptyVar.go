package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	var result int
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Const1()
	{
		name = varname.String()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	if value == nil {
		result = opline.GetExtendedValue() & ZEND_ISEMPTY
	} else {
		if value.IsIndirect() {
			value = value.Indirect()
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
			if value.IsReference() {
				value = types2.Z_REFVAL_P(value)
			}
			result = value.GetType() > types2.IS_NULL
		} else {
			result = !(IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	var result int
	var free_op1 ZendFreeOp
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Op1()

	{
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	// ZvalPtrDtorNogc(free_op1)
	if value == nil {
		result = opline.GetExtendedValue() & ZEND_ISEMPTY
	} else {
		if value.IsIndirect() {
			value = value.Indirect()
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
			if value.IsReference() {
				value = types2.Z_REFVAL_P(value)
			}
			result = value.GetType() > types2.IS_NULL
		} else {
			result = !(IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	var result int
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Op1()
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	if value == nil {
		result = opline.GetExtendedValue() & ZEND_ISEMPTY
	} else {
		if value.IsIndirect() {
			value = value.Indirect()
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
			if value.IsReference() {
				value = types2.Z_REFVAL_P(value)
			}
			result = value.GetType() > types2.IS_NULL
		} else {
			result = !(IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

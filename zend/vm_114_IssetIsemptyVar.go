package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
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
				value = types.Z_REFVAL_P(value)
			}
			result = value.GetType() > types.IsNull
		} else {
			result = !(operators.IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var varname *types.Zval
	var name *types.String
	var target_symbol_table *types.Array
	varname = opline.Op1()
	name = operators.ZvalGetString(varname)
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
				value = types.Z_REFVAL_P(value)
			}
			result = value.GetType() > types.IsNull
		} else {
			result = !(operators.IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
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
				value = types.Z_REFVAL_P(value)
			}
			result = value.GetType() > types.IsNull
		} else {
			result = !(operators.IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Const1()
	{
		name = varname.String()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types2.ZendHashDelInd(target_symbol_table, name.GetStr())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	var free_op1 ZendFreeOp
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.String()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1(executeData)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types2.ZendHashDelInd(target_symbol_table, name.GetStr())
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.String()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1(executeData)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types2.ZendHashDelInd(target_symbol_table, name.GetStr())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

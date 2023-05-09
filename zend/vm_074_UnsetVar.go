package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var target_symbol_table *types.Array
	varname = opline.Const1()
	name = varname.String()

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name string
	var target_symbol_table *types.Array
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.StringVal()
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1(executeData)
		}
		var ok bool
		if name, ok = ZvalGetStr(varname); !ok {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, *name)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name string
	var target_symbol_table *types.Array
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.StringVal()
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1(executeData)
		}
		var ok bool
		if name, ok = ZvalGetStr(varname); !ok {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

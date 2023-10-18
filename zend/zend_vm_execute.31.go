package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func zend_fetch_var_address_helper_SPEC_CV_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var retval *types.Zval
	var name string
	var target_symbol_table *types.Array
	var ok bool
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.String()
	} else {
		if varname.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		name, ok = operators.ZvalTryGetStr(varname)
		if !ok {
			opline.Result().SetUndef()
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name)
	if retval == nil {
		if name == types.STR_THIS {
			goto fetchThis
		}
		if type_ == BP_VAR_W {
			retval = target_symbol_table.KeyAddNew(name, UninitializedZval())
		} else if type_ == BP_VAR_IS {
			retval = UninitializedZval()
		} else {
			faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined variable: %s", name))
			if type_ == BP_VAR_RW {
				retval = target_symbol_table.KeyUpdate(name, UninitializedZval())
			} else {
				retval = UninitializedZval()
			}
		}
	} else if retval.IsIndirect() {
		retval = retval.Indirect()
		if retval.IsUndef() {
			if name == types.STR_THIS {
				goto fetchThis
			}
			if type_ == BP_VAR_W {
				retval.SetNull()
			} else if type_ == BP_VAR_IS {
				retval = UninitializedZval()
			} else {
				faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined variable: %s", name))
				if type_ == BP_VAR_RW {
					retval.SetNull()
				} else {
					retval = UninitializedZval()
				}
			}
		}
	}

	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else {
		opline.Result().SetIndirect(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
fetchThis:
	ZendFetchThisVar(type_, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

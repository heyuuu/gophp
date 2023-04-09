package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func zend_fetch_var_address_helper_SPEC_CONST_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var retval *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = opline.Const1()
	{
		name = varname.GetStr()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if name.GetStr() == types.STR_THIS {
		fetch_this:
			ZendFetchThisVar(type_, opline, executeData)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if type_ == BP_VAR_W {
			retval = target_symbol_table.KeyAddNew(name.GetStr(), EG__().GetUninitializedZval())
		} else if type_ == BP_VAR_IS {
			retval = EG__().GetUninitializedZval()
		} else {
			faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
			if type_ == BP_VAR_RW {
				retval = target_symbol_table.KeyUpdate(name.GetStr(), EG__().GetUninitializedZval())
			} else {
				retval = EG__().GetUninitializedZval()
			}
		}
	} else if retval.IsIndirect() {
		retval = retval.GetZv()
		if retval.IsUndef() {
			if name.GetStr() == types.STR_THIS {
				goto fetch_this
			}
			if type_ == BP_VAR_W {
				retval.SetNull()
			} else if type_ == BP_VAR_IS {
				retval = EG__().GetUninitializedZval()
			} else {
				faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
				if type_ == BP_VAR_RW {
					retval.SetNull()
				} else {
					retval = EG__().GetUninitializedZval()
				}
			}
		}
	}
	if (opline.GetExtendedValue() & ZEND_FETCH_GLOBAL_LOCK) == 0 {
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else {
		opline.Result().SetIndirect(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

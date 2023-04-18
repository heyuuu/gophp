package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varname *types2.Zval
	var retval *types2.Zval
	var name *types2.String
	var tmp_name *types2.String
	var target_symbol_table *types2.Array
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.String()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			// ZvalPtrDtorNogc(free_op1)
			opline.Result().SetUndef()
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if name.GetStr() == types2.STR_THIS {
		fetch_this:
			ZendFetchThisVar(type_, opline, executeData)
			{
				ZendTmpStringRelease(tmp_name)
			}
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
		retval = retval.Indirect()
		if retval.IsUndef() {
			if name.GetStr() == types2.STR_THIS {
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
		// ZvalPtrDtorNogc(free_op1)
	}
	{
		ZendTmpStringRelease(tmp_name)
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types2.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else {
		opline.Result().SetIndirect(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

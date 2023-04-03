package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varname *types.Zval
	var retval *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = opline.getZvalPtrVar1(&free_op1)

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			ZvalPtrDtorNogc(free_op1)
			opline.GetResultZval().SetUndef()
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
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
		retval = retval.GetZv()
		if retval.IsUndef() {
			if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
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
		ZvalPtrDtorNogc(free_op1)
	}
	{
		ZendTmpStringRelease(tmp_name)
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else {
		opline.GetResultZval().SetIndirect(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

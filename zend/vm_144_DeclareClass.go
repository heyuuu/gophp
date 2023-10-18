package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()

	var lcname *types.Zval = opline.Const1()
	var rtdKey *types.Zval = lcname + 1

	if opline.GetOp2Type() == IS_CONST {
		DoBindClass(lcname, rtdKey, opline.Const2().String())
	} else {
		DoBindClass(lcname, rtdKey, nil)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

func DoBindClass(lcname *types.Zval, rtdKey *types.Zval, lcParentName string) int {
	ce := EG__().ClassTable().Get(rtdKey.String())
	if ce == nil {
		if EG__().ClassTable().Exists(lcname.String()) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.Name()))
			return types.FAILURE
		} else {
			b.Assert(CurrEX().GetFunc().GetOpArray().IsPreloaded())
			faults.ErrorNoreturn(faults.E_ERROR, fmt.Sprintf("Class %s wasn't preloaded", lcname.String()))
			return types.FAILURE
		}
	}

	if EG__().ClassTable().Exists(lcname.String()) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.Name()))
		return types.FAILURE
	}

	if ZendDoLinkClass(ce, lcParentName) == types.FAILURE {
		return types.FAILURE
	}

	EG__().ClassTable().Del(rtdKey.String())
	EG__().ClassTable().Add(lcname.String(), ce)

	return types.SUCCESS
}

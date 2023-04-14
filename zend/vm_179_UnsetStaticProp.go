package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_UNSET_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String = nil
	var ce *types.ClassEntry
	var free_op1 ZendFreeOp
	if opline.GetOp2Type() == IS_CONST {
		ce = CACHED_PTR(opline.GetExtendedValue())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const2().GetStr(), (opline.Const2() + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
				return 0
			}
		}
	} else if opline.GetOp2Type() == IS_UNUSED {
		ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
			return 0
		}
	} else {
		ce = opline.Op2().GetCe()
	}
	varname = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
	if opline.GetOp1Type() == IS_CONST {
		name = varname.GetStr()
	} else if varname.IsString() {
		name = varname.GetStr()
	} else {
		if opline.GetOp1Type() == IS_CV && varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1(executeData)
		}
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	ZendStdUnsetStaticProperty(ce, name)
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

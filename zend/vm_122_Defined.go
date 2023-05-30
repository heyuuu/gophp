package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_DEFINED_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var c *ZendConstant
	c = CACHED_PTR(opline.GetExtendedValue())
	if c != nil {
		if IS_SPECIAL_CACHE_VAL(c) == 0 {
			goto defined_true
		} else if EG__().ConstantTable().Len() == DECODE_SPECIAL_CACHE_NUM(c) {
			goto defined_false
		}
	}
	if ZendQuickCheckConstant(opline.Const1(), opline, executeData) != types.SUCCESS {
		CACHE_PTR(opline.GetExtendedValue(), ENCODE_SPECIAL_CACHE_NUM(EG__().ConstantTable().Len()))
		goto defined_false
	} else {
		goto defined_true
	}

	panic("unreachable")
defined_true:
	ZEND_VM_SMART_BRANCH_TRUE()
	opline.Result().SetTrue()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)

defined_false:
	ZEND_VM_SMART_BRANCH_FALSE()
	opline.Result().SetFalse()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

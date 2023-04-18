package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var name *types2.Zval
	var val *types2.Zval
	var c ZendConstant
	name = opline.Const1()
	val = opline.Const2()
	types2.ZVAL_COPY(c.Value(), val)
	if c.Value().IsConstantAst() {
		if ZvalUpdateConstantEx(c.Value(), executeData.GetFunc().GetOpArray().scope) != types2.SUCCESS {
			// ZvalPtrDtorNogc(c.Value())
			return 0
		}
	}

	/* non persistent, case sensitive */

	c.SetFlags(CONST_CS, PHP_USER_CONSTANT)
	c.SetName(name.StringVal())
	ZendRegisterConstant(&c)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

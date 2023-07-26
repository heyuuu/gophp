package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.Object = container.Object()
	var retval *types.Zval

	retval = zobj.ReadPropertyEx(offset, BP_VAR_R, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

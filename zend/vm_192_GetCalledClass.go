package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		opline.Result().SetStringVal(executeData.GetThis().Object().GetCe().Name())
	} else if executeData.GetThis().Class() != nil {
		opline.Result().SetStringVal(executeData.GetThis().Object().GetCe().Name())
	} else {
		opline.Result().SetFalse()
		if executeData.GetFunc().GetScope() == nil {
			faults.Error(faults.E_WARNING, "get_called_class() called from outside a class")
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

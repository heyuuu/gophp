package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		opline.Result().SetStringCopy(types2.Z_OBJCE(executeData.GetThis()).GetName())
	} else if executeData.GetThis().Class() != nil {
		opline.Result().SetStringCopy(types.Z_CE(executeData.GetThis()).GetName())
	} else {
		opline.Result().SetFalse()
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_called_class() called from outside a class")
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

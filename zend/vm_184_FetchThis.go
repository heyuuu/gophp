package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_FETCH_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		var result *types.Zval = opline.Result()
		result.SetObject(executeData.GetThis().Object())
		// 		result.AddRefcount()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
}

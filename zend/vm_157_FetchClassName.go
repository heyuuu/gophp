package zend

import (
	"fmt"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type uint32
	var called_scope *types.ClassEntry
	var scope *types.ClassEntry
	var opline *types.ZendOp = executeData.GetOpline()
	fetch_type = opline.GetOp1().GetNum()
	scope = executeData.GetFunc().GetOpArray().scope
	if scope == nil {
		faults.ThrowError(nil, fmt.Sprintf(`Cannot use "%s" when no class scope is active`, fetchTypeName(fetch_type)))
		opline.Result().SetUndef()
		return 0
	}
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		opline.Result().SetString(scope.Name())
	case ZEND_FETCH_CLASS_PARENT:
		if scope.GetParent() == nil {
			faults.ThrowError(nil, `Cannot use "parent" when current class scope has no parent`)
			opline.Result().SetUndef()
			return 0
		}
		opline.Result().SetString(scope.GetParent().name.GetStr())
	case ZEND_FETCH_CLASS_STATIC:
		called_scope = executeData.ThisClass()
		opline.Result().SetString(called_scope.Name())
	default:

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

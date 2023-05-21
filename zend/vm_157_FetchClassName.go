package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type uint32
	var called_scope *types.ClassEntry
	var scope *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	fetch_type = opline.GetOp1().GetNum()
	scope = executeData.GetFunc().GetOpArray().scope
	if scope == nil {
		faults.ThrowError(nil, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		opline.Result().SetUndef()
		return 0
	}
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		opline.Result().SetStringVal(scope.Name())
	case ZEND_FETCH_CLASS_PARENT:
		if scope.GetParent() == nil {
			faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
			opline.Result().SetUndef()
			return 0
		}
		opline.Result().SetStringVal(scope.GetParent().name.GetStr())
	case ZEND_FETCH_CLASS_STATIC:
		if executeData.GetThis().IsObject() {
			called_scope = types.Z_OBJCE(executeData.GetThis())
		} else {
			called_scope = executeData.GetThis().Class()
		}
		opline.Result().SetStringVal(called_scope.Name())
	default:

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

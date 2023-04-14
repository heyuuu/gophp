package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_ types.IFunction
	var object *types.Zval
	var called_scope *types.ClassEntry
	func_ = CACHED_PTR(opline.GetExtendedValue())
	if func_ == nil {
		func_ = EG__().FunctionTable().Get(opline.Const1().String().GetStr())
		b.Assert(func_ != nil)
		b.Assert(func_.GetType() == ZEND_USER_FUNCTION)
		CACHE_PTR(opline.GetExtendedValue(), func_)
	}
	if executeData.GetThis().IsObject() {
		called_scope = types.Z_OBJCE(executeData.GetThis())
		if func_.IsStatic() || (executeData.GetFunc().GetFnFlags()&AccStatic) != 0 {
			object = nil
		} else {
			object = &(executeData.GetThis())
		}
	} else {
		called_scope = executeData.GetThis().Class()
		object = nil
	}
	ZendCreateClosure(opline.Result(), func_, executeData.GetFunc().GetOpArray().GetScope(), called_scope, object)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

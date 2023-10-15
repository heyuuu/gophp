package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var func_ types.IFunction
	var object *types.Object
	var called_scope *types.ClassEntry
	func_ = CACHED_PTR(opline.GetExtendedValue())
	if func_ == nil {
		func_ = EG__().FunctionTable().Get(opline.Const1().String())
		b.Assert(func_ != nil)
		b.Assert(func_.GetType() == ZEND_USER_FUNCTION)
		CACHE_PTR(opline.GetExtendedValue(), func_)
	}
	if executeData.InScope() {
		called_scope = executeData.ThisClass()
		if func_.IsStatic() || (executeData.GetFunc().GetFnFlags()&types.AccStatic) != 0 {
			object = nil
		} else {
			object = executeData.ThisObject()
		}
	} else {
		called_scope = executeData.ThisClass()
		object = nil
	}
	ZendCreateClosureEx(opline.Result(), func_, executeData.GetFunc().GetOpArray().GetScope(), called_scope, object)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

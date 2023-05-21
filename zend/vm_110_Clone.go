package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func cloneHandlerHelper(executeData *ZendExecuteData, obj *types.Zval) int {
	var opline *ZendOp = executeData.GetOpline()

	ce := types.Z_OBJCE_P(obj)
	clone := ce.GetClone()
	if !obj.Object().CanClone() {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.Name())
		opline.Result().SetUndef()
		return 0
	}
	if clone != nil && !clone.IsPublic() {
		scope := executeData.GetFunc().GetOpArray().GetScope()
		if clone.GetScope() != scope {
			if clone.IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) {
				ZendWrongCloneCall(clone, scope)
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	opline.Result().SetObject(obj.Object().Clone())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

func ZEND_CLONE_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	opline.Result().SetUndef()
	faults.ThrowError(nil, "__clone method called on non-object")
	return 0
}
func ZEND_CLONE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	obj = opline.Op1()
	for {
		if obj.GetType() != types.IS_OBJECT {
			if obj.IsReference() {
				obj = types.Z_REFVAL_P(obj)
				if obj.IsObject() {
					break
				}
			}
			opline.Result().SetUndef()
			if obj.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "__clone method called on non-object")
			return 0
		}
		break
	}

	return cloneHandlerHelper(executeData, obj)
}
func ZEND_CLONE_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var obj *types.Zval
	obj = executeData.GetThis()
	if obj.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	return cloneHandlerHelper(executeData, obj)
}
func ZEND_CLONE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	obj = opline.Op1()
	for {
		if obj.GetType() != types.IS_OBJECT {
			if obj.IsReference() {
				obj = types.Z_REFVAL_P(obj)
				if obj.IsObject() {
					break
				}
			}
			opline.Result().SetUndef()
			if obj.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "__clone method called on non-object")
			return 0
		}
		break
	}

	return cloneHandlerHelper(executeData, obj)
}

package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Const1()
	offset = opline.Const2()
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Const1()
	offset = opline.Op2()
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Const1()
	offset = opline.Cv2OrUndef()
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Op2()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op2)
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Cv2OrUndef()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Const2()
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Op2()
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Cv2OrUndef()
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Op2()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Cv2OrUndef()
	if !container.IsObject() {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if !container.IsObject() {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ container.Object().HasProperty(offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

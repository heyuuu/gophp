package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_UNSET_OBJ_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = opline.Const2()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = opline.Op2()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	// ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = opline.Cv2OrUndef()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Const2()
	container.Object().UnsetPropertyEx(offset)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Op2()
	container.Object().UnsetPropertyEx(offset)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Cv2OrUndef()
	container.Object().UnsetPropertyEx(offset)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Op2()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Cv2OrUndef()
	for {
		if !container.IsObject() {
			if container.IsRef() {
				container = types.Z_REFVAL_P(container)
				if !container.IsObject() {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1(executeData)
					}
					break
				}
			} else {
				break
			}
		}
		container.Object().UnsetPropertyEx(offset)
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

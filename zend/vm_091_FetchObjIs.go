package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func fetchObjIsHandler0(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	opline.Result().SetNull()
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Op2()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Cv2OrUndef()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Const2()
	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Op2()
	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = executeData.ThisObjectZval()
	if container == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Cv2OrUndef()
	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Const2()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Op2()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = opline.Op1()
	offset = opline.Cv2OrUndef()
	if !container.IsObject() {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.Result().SetNull()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj = container.Object()
	var retval *types.Zval
	retval = zobj.ReadPropertyEx(offset, BP_VAR_IS, opline.Result())
	if retval != opline.Result() {
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		operators.ZendUnwrapReference(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

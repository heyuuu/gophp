package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var offset *types.Zval
	offset = opline.Const2()
	ZendWrongPropertyRead(offset)
	opline.Result().SetNull()
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Const1()
	offset = opline.Op2()
	{
		for {
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Const1()
	offset = opline.Op2()
	{
		for {
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Op1()
	offset = opline.Const2()
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_REF)
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if !retval.IsUndef() {
					{
						goto fetch_obj_r_copy
					}

				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().Bucket(idx)))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.GetStr() || (p.IsStrKey() && p.StrKey() == offset.StringVal())) {
							retval = p.GetVal()
							goto fetch_obj_r_copy
						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.StringVal())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
					{
						goto fetch_obj_r_copy
					}

				}
			}
		}
	}

	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Op1()
	offset = opline.Op2()
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op2)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Op1()
	offset = opline.Op2()
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Op2()
	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Op2()
	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Op1()
	offset = opline.Op2()
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.Op1()
	offset = opline.Op2()
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			ZendWrongPropertyRead(offset)
			opline.Result().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

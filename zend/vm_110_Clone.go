package zend

func ZEND_CLONE_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone types.IFunction
	var clone_call ZendObjectCloneObjT
	obj = opline.Const1()
	for {
		{
			opline.Result().SetUndef()
			faults.ThrowError(nil, "__clone method called on non-object")
			return 0
		}
		break
	}
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		opline.Result().SetUndef()
		return 0
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().GetOpArray().scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	opline.Result().SetObject(clone_call(obj))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CLONE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone types.IFunction
	var clone_call ZendObjectCloneObjT
	obj = opline.Op1Ptr(&free_op1)
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
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		break
	}
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		// ZvalPtrDtorNogc(free_op1)
		opline.Result().SetUndef()
		return 0
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().GetOpArray().scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				// ZvalPtrDtorNogc(free_op1)
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	opline.Result().SetObject(clone_call(obj))
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CLONE_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone types.IFunction
	var clone_call ZendObjectCloneObjT
	obj = &(executeData.GetThis())
	if obj.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	for {
		break
	}
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		opline.Result().SetUndef()
		return 0
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().GetOpArray().scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	opline.Result().SetObject(clone_call(obj))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CLONE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone types.IFunction
	var clone_call ZendObjectCloneObjT
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
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		opline.Result().SetUndef()
		return 0
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().GetOpArray().scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	opline.Result().SetObject(clone_call(obj))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

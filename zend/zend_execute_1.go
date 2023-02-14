// <<generate>>

package zend

func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsUndef() {
		ret.SetNull()
		ZvalUndefinedCv(var_, EXECUTE_DATA_C)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsUndef() {
		ret.SetNull()
	}
	return ret
}
func _getZvalPtr(
	op_type int,
	node ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	_ EXECUTE_DATA_D,
	opline *ZendOp,
) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrDeref(
	op_type int,
	node ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	_ EXECUTE_DATA_D,
	opline *ZendOp,
) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCvDeref(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_deref_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrUndef(
	op_type int,
	node ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	_ EXECUTE_DATA_D,
	opline *ZendOp,
) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return EX_VAR(node.GetVar())
		} else {
			return nil
		}
	}
}
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsIndirect() {
		*should_free = nil
		ret = ret.GetZv()
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_CV {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
	} else {
		ZEND_ASSERT(op_type == IS_VAR)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	}
}
func _getObjZvalPtr(
	op_type int,
	op ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	_ EXECUTE_DATA_D,
	opline *ZendOp,
) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtr(op_type, op, should_free, type_)
}
func _getObjZvalPtrUndef(
	op_type int,
	op ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	_ EXECUTE_DATA_D,
	opline *ZendOp,
) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtrUndef(op_type, op, should_free, type_)
}
func _getObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtrPtr(op_type, node, should_free, type_)
}
func ZendAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval) {
	var ref *ZendReference
	if !(value_ptr.IsReference()) {
		value_ptr.SetNewRef(value_ptr)
	} else if variable_ptr == value_ptr {
		return
	}
	ref = value_ptr.GetRef()
	ref.AddRefcount()
	if variable_ptr.IsRefcounted() {
		var garbage *ZendRefcounted = variable_ptr.GetCounted()
		if garbage.DelRefcount() == 0 {
			variable_ptr.SetReference(ref)
			RcDtorFunc(garbage)
			return
		} else {
			GcCheckPossibleRoot(garbage)
		}
	}
	variable_ptr.SetReference(ref)
}
func ZendAssignToTypedPropertyReference(prop_info *ZendPropertyInfo, prop *Zval, value_ptr *Zval, _ EXECUTE_DATA_D) *Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, EX_USES_STRICT_TYPES()) == 0 {
		return EG__().GetUninitializedZval()
	}
	if prop.IsReference() {
		ZEND_REF_DEL_TYPE_SOURCE(prop.GetRef(), prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZEND_REF_ADD_TYPE_SOURCE(prop.GetRef(), prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	ZendError(E_NOTICE, "Only variables should be assigned by reference")
	if EG__().GetException() != nil {
		return EG__().GetUninitializedZval()
	}

	/* Use IS_TMP_VAR instead of IS_VAR to avoid ISREF check */

	value_ptr.TryAddRefcount()
	return ZendAssignToVariable(variable_ptr, value_ptr, IS_TMP_VAR, EX_USES_STRICT_TYPES())
}
func ZendFormatType(type_ ZendType, part1 **byte, part2 **byte) {
	if type_.AllowNull() {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if type_.IsClass() {
		if type_.IsCe() {
			*part2 = ZEND_TYPE_CE(type_).GetName().GetVal()
		} else {
			*part2 = ZEND_TYPE_NAME(type_).GetVal()
		}
	} else {
		*part2 = ZendGetTypeByConst(type_.Code())
	}
}
func ZendThrowAutoInitInPropError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAutoInitInRefError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAccessUninitPropByRefError(prop *ZendPropertyInfo) {
	ZendThrowError(nil, "Cannot access uninitialized non-nullable property %s::$%s by reference", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()))
}
func MakeRealObject(object *Zval, property *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	var obj *ZendObject
	var ref *Zval = nil
	if object.IsReference() {
		ref = object
		object = Z_REFVAL_P(object)
	}
	if object.GetType() > IS_FALSE && (object.GetType() != IS_STRING || Z_STRLEN_P(object) != 0) {
		if opline.GetOp1Type() != IS_VAR || !(object.IsError()) {
			var tmp_property_name *ZendString
			var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
			if opline.GetOpcode() == ZEND_PRE_INC_OBJ || opline.GetOpcode() == ZEND_PRE_DEC_OBJ || opline.GetOpcode() == ZEND_POST_INC_OBJ || opline.GetOpcode() == ZEND_POST_DEC_OBJ {
				ZendError(E_WARNING, "Attempt to increment/decrement property '%s' of non-object", property_name.GetVal())
			} else if opline.GetOpcode() == ZEND_FETCH_OBJ_W || opline.GetOpcode() == ZEND_FETCH_OBJ_RW || opline.GetOpcode() == ZEND_FETCH_OBJ_FUNC_ARG || opline.GetOpcode() == ZEND_ASSIGN_OBJ_REF {
				ZendError(E_WARNING, "Attempt to modify property '%s' of non-object", property_name.GetVal())
			} else {
				ZendError(E_WARNING, "Attempt to assign property '%s' of non-object", property_name.GetVal())
			}
			ZendTmpStringRelease(tmp_property_name)
		}
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
		return nil
	}
	if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref.GetRef()) {
		if zend_verify_ref_stdClass_assignable(ref.GetRef()) == 0 {
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
			}
			return nil
		}
	}
	ZvalPtrDtorNogc(object)
	ObjectInit(object)
	object.AddRefcount()
	obj = object.GetObj()
	ZendError(E_WARNING, "Creating default object from empty value")
	if obj.GetRefcount() == 1 {

		/* the enclosing container was deleted, obj is unreferenced */

		OBJ_RELEASE(obj)
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
		return nil
	}
	object.DelRefcount()
	return object
}
func ZendVerifyTypeErrorCommon(
	zf *ZendFunction,
	arg_info *ZendArgInfo,
	ce *ZendClassEntry,
	value *Zval,
	fname **byte,
	fsep **byte,
	fclass **byte,
	need_msg **byte,
	need_kind **byte,
	need_or_null **byte,
	given_msg **byte,
	given_kind **byte,
) {
	var is_interface ZendBool = 0
	*fname = zf.GetFunctionName().GetVal()
	if zf.GetScope() != nil {
		*fsep = "::"
		*fclass = zf.GetScope().GetName().GetVal()
	} else {
		*fsep = ""
		*fclass = ""
	}
	if arg_info.GetType().IsClass() {
		if ce != nil {
			if ce.IsInterface() {
				*need_msg = "implement interface "
				is_interface = 1
			} else {
				*need_msg = "be an instance of "
			}
			*need_kind = ce.GetName().GetVal()
		} else {

			/* We don't know whether it's a class or interface, assume it's a class */

			*need_msg = "be an instance of "
			*need_kind = ZEND_TYPE_NAME(arg_info.GetType()).GetVal()
		}
	} else {
		switch arg_info.GetType().Code() {
		case IS_OBJECT:
			*need_msg = "be an "
			*need_kind = "object"
		case IS_CALLABLE:
			*need_msg = "be callable"
			*need_kind = ""
		case IS_ITERABLE:
			*need_msg = "be iterable"
			*need_kind = ""
		default:
			*need_msg = "be of the type "
			*need_kind = ZendGetTypeByConst(arg_info.GetType().Code())
		}
	}
	if arg_info.GetType().AllowNull() {
		if is_interface != 0 {
			*need_or_null = " or be null"
		} else {
			*need_or_null = " or null"
		}
	} else {
		*need_or_null = ""
	}
	if value != nil {
		if arg_info.GetType().IsClass() && value.IsObject() {
			*given_msg = "instance of "
			*given_kind = Z_OBJCE_P(value).GetName().GetVal()
		} else {
			*given_msg = ZendZvalTypeName(value)
			*given_kind = ""
		}
	} else {
		*given_msg = "none"
		*given_kind = ""
	}
}
func ZendVerifyArgError(zf *ZendFunction, arg_info *ZendArgInfo, arg_num int, ce *ZendClassEntry, value *Zval) {
	var ptr *ZendExecuteData = EG__().GetCurrentExecuteData().GetPrevExecuteData()
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	if EG__().GetException() != nil {

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

		return

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

	}
	if value != nil {
		ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
		if zf.GetCommonType() == ZEND_USER_FUNCTION {
			if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno())
			} else {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
			}
		} else {
			ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
		}
	} else {
		ZendMissingArgError(ptr)
	}
}
func IsNullConstant(scope *ZendClassEntry, default_value *Zval) int {
	if default_value.IsConstant() {
		var constant Zval
		ZVAL_COPY(&constant, default_value)
		if ZvalUpdateConstantEx(&constant, scope) != SUCCESS {
			return 0
		}
		if constant.IsNull() {
			return 1
		}
		ZvalPtrDtorNogc(&constant)
	}
	return 0
}

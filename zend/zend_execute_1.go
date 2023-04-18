package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, executeData *ZendExecuteData) *types2.Zval {
	var ret *types2.Zval = EX_VAR(executeData, var_)
	if ret.IsUndef() {
		ret.SetNull()
		ZvalUndefinedCv(var_, executeData)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, executeData *ZendExecuteData) *types2.Zval {
	var ret *types2.Zval = EX_VAR(executeData, var_)
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
	executeData *ZendExecuteData,
	opline *ZendOp,
) *types2.Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, executeData)
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCv(node.GetVar(), type_, executeData)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp, executeData *ZendExecuteData, opline *ZendOp) *types2.Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, executeData)
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_BP_VAR_R(node.GetVar(), executeData)
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
	executeData *ZendExecuteData,
	opline *ZendOp,
) *types2.Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, executeData)
		} else {
			b.Assert(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, executeData)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCvDeref(node.GetVar(), type_, executeData)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp, executeData *ZendExecuteData, opline *ZendOp) *types2.Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, executeData)
		} else {
			b.Assert(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, executeData)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_deref_BP_VAR_R(node.GetVar(), executeData)
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
	executeData *ZendExecuteData,
	opline *ZendOp,
) *types2.Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		return _getZvalPtrVar(node.GetVar(), should_free, executeData)
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
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types2.Zval {
	var ret *types2.Zval = EX_VAR(executeData, var_)
	if ret.IsIndirect() {
		*should_free = nil
		ret = ret.Indirect()
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, executeData *ZendExecuteData) *types2.Zval {
	if op_type == IS_CV {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, executeData)
	} else {
		b.Assert(op_type == IS_VAR)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, executeData)
	}
}
func _getObjZvalPtr(
	op_type int,
	op ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	executeData *ZendExecuteData,
	opline *ZendOp,
) *types2.Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(executeData.GetThis())
	}
	return GetZvalPtr(op_type, op, should_free, type_)
}
func _getObjZvalPtrUndef(
	op_type int,
	op ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	executeData *ZendExecuteData,
	opline *ZendOp,
) *types2.Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(executeData.GetThis())
	}
	return GetZvalPtrUndef(op_type, op, should_free, type_)
}
func _getObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, executeData *ZendExecuteData) *types2.Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(executeData.GetThis())
	}
	return GetZvalPtrPtr(op_type, node, should_free, type_)
}
func ZendAssignToVariableReference(variable_ptr *types2.Zval, value_ptr *types2.Zval) {
	var ref *types2.ZendReference
	if !(value_ptr.IsReference()) {
		value_ptr.SetNewRef(value_ptr)
	} else if variable_ptr == value_ptr {
		return
	}
	ref = value_ptr.Reference()
	// 	ref.AddRefcount()
	if variable_ptr.IsRefcounted() {
		var garbage *types2.ZendRefcounted = variable_ptr.RefCounted()
		if garbage.DelRefcount() == 0 {
			variable_ptr.SetReference(ref)
			//RcDtorFunc(garbage)
			return
		} else {
			//GcCheckPossibleRoot(garbage)
		}
	}
	variable_ptr.SetReference(ref)
}
func ZendAssignToTypedPropertyReference(prop_info *ZendPropertyInfo, prop *types2.Zval, value_ptr *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, executeData.IsCallUseStrictTypes()) == 0 {
		return EG__().GetUninitializedZval()
	}
	if prop.IsReference() {
		ZEND_REF_DEL_TYPE_SOURCE(prop.Reference(), prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZEND_REF_ADD_TYPE_SOURCE(prop.Reference(), prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *types2.Zval, value_ptr *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) *types2.Zval {
	faults.Error(faults.E_NOTICE, "Only variables should be assigned by reference")
	if EG__().GetException() != nil {
		return EG__().GetUninitializedZval()
	}

	/* Use IS_TMP_VAR instead of IS_VAR to avoid ISREF check */

	// value_ptr.TryAddRefcount()
	return ZendAssignToVariable(variable_ptr, value_ptr, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
}
func ZendFormatType(type_ types2.ZendType, part1 **byte, part2 **byte) {
	if type_.AllowNull() {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if type_.IsClass() {
		if type_.IsCe() {
			*part2 = types2.ZEND_TYPE_CE(type_).GetName().GetVal()
		} else {
			*part2 = types2.ZEND_TYPE_NAME(type_).GetVal()
		}
	} else {
		*part2 = types2.ZendGetTypeByConst(type_.Code())
	}
}
func ZendThrowAutoInitInPropError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	faults.TypeError("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAutoInitInRefError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	faults.TypeError("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAccessUninitPropByRefError(prop *ZendPropertyInfo) {
	faults.ThrowError(nil, "Cannot access uninitialized non-nullable property %s::$%s by reference", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()))
}
func MakeRealObject(object *types2.Zval, property *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) *types2.Zval {
	var obj *types2.ZendObject
	var ref *types2.Zval = nil
	if object.IsReference() {
		ref = object
		object = types2.Z_REFVAL_P(object)
	}
	if object.GetType() > types2.IS_FALSE && (object.GetType() != types2.IS_STRING || object.String().GetLen() != 0) {
		if opline.GetOp1Type() != IS_VAR || !(object.IsError()) {
			var tmp_property_name *types2.String
			var property_name *types2.String = ZvalGetTmpString(property, &tmp_property_name)
			if opline.GetOpcode() == ZEND_PRE_INC_OBJ || opline.GetOpcode() == ZEND_PRE_DEC_OBJ || opline.GetOpcode() == ZEND_POST_INC_OBJ || opline.GetOpcode() == ZEND_POST_DEC_OBJ {
				faults.Error(faults.E_WARNING, "Attempt to increment/decrement property '%s' of non-object", property_name.GetVal())
			} else if opline.GetOpcode() == ZEND_FETCH_OBJ_W || opline.GetOpcode() == ZEND_FETCH_OBJ_RW || opline.GetOpcode() == ZEND_FETCH_OBJ_FUNC_ARG || opline.GetOpcode() == ZEND_ASSIGN_OBJ_REF {
				faults.Error(faults.E_WARNING, "Attempt to modify property '%s' of non-object", property_name.GetVal())
			} else {
				faults.Error(faults.E_WARNING, "Attempt to assign property '%s' of non-object", property_name.GetVal())
			}
			ZendTmpStringRelease(tmp_property_name)
		}
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return nil
	}
	if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref.Reference()) {
		if zend_verify_ref_stdClass_assignable(ref.Reference()) == 0 {
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetUndef()
			}
			return nil
		}
	}
	// ZvalPtrDtorNogc(object)
	ObjectInit(object)
	// 	object.AddRefcount()
	obj = object.Object()
	faults.Error(faults.E_WARNING, "Creating default object from empty value")
	if obj.GetRefcount() == 1 {

		/* the enclosing container was deleted, obj is unreferenced */

		// OBJ_RELEASE(obj)
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return nil
	}
	object.DelRefcount()
	return object
}
func ZendVerifyTypeErrorCommon(
	zf types2.IFunction,
	arg_info *ZendArgInfo,
	ce *types2.ClassEntry,
	value *types2.Zval,
	fname **byte,
	fsep **byte,
	fclass **byte,
	need_msg **byte,
	need_kind **byte,
	need_or_null **byte,
	given_msg **byte,
	given_kind **byte,
) {
	var is_interface types2.ZendBool = 0
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
			*need_kind = types2.ZEND_TYPE_NAME(arg_info.GetType()).GetVal()
		}
	} else {
		switch arg_info.GetType().Code() {
		case types2.IS_OBJECT:
			*need_msg = "be an "
			*need_kind = "object"
		case types2.IS_CALLABLE:
			*need_msg = "be callable"
			*need_kind = ""
		case types2.IS_ITERABLE:
			*need_msg = "be iterable"
			*need_kind = ""
		default:
			*need_msg = "be of the type "
			*need_kind = types2.ZendGetTypeByConst(arg_info.GetType().Code())
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
			*given_kind = types2.Z_OBJCE_P(value).GetName().GetVal()
		} else {
			*given_msg = types2.ZendZvalTypeName(value)
			*given_kind = ""
		}
	} else {
		*given_msg = "none"
		*given_kind = ""
	}
}
func ZendVerifyArgError(zf types2.IFunction, arg_info *ZendArgInfo, arg_num int, ce *types2.ClassEntry, value *types2.Zval) {
	var ptr *ZendExecuteData = CurrEX().GetPrevExecuteData()
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
		if zf.GetType() == ZEND_USER_FUNCTION {
			if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetType()) {
				faults.TypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno())
			} else {
				faults.TypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
			}
		} else {
			faults.TypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
		}
	} else {
		ZendMissingArgError(ptr)
	}
}
func IsNullConstant(scope *types2.ClassEntry, default_value *types2.Zval) int {
	if default_value.IsConstantAst() {
		var constant types2.Zval
		types2.ZVAL_COPY(&constant, default_value)
		if ZvalUpdateConstantEx(&constant, scope) != types2.SUCCESS {
			return 0
		}
		if constant.IsNull() {
			return 1
		}
		// ZvalPtrDtorNogc(&constant)
	}
	return 0
}

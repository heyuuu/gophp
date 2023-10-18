package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	if ret.IsUndef() {
		ret.SetNull()
		ZvalUndefinedCv(var_, executeData)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	if ret.IsUndef() {
		ret.SetNull()
	}
	return ret
}
func _getZvalPtr(
	op_type int,
	node types.ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	executeData *ZendExecuteData,
	opline *types.ZendOp,
) *types.Zval {
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
func _getOpDataZvalPtrR(op_type int, node types.ZnodeOp, should_free *ZendFreeOp, executeData *ZendExecuteData, opline *types.ZendOp) *types.Zval {
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
func _getZvalPtrUndef(
	op_type int,
	node types.ZnodeOp,
	should_free *ZendFreeOp,
	type_ int,
	executeData *ZendExecuteData,
	opline *types.ZendOp,
) *types.Zval {
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
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	if ret.IsIndirect() {
		*should_free = nil
		ret = ret.Indirect()
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node types.ZnodeOp, should_free *ZendFreeOp, type_ int, executeData *ZendExecuteData) *types.Zval {
	if op_type == IS_CV {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, executeData)
	} else {
		b.Assert(op_type == IS_VAR)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, executeData)
	}
}
func ZendAssignToVariableReference(variable_ptr *types.Zval, value_ptr *types.Zval) {
	var ref *types.Reference
	if !(value_ptr.IsRef()) {
		value_ptr.SetNewRef(value_ptr)
	} else if variable_ptr == value_ptr {
		return
	}
	ref = value_ptr.Ref()
	variable_ptr.SetReference(ref)
}
func ZendAssignToTypedPropertyReference(prop_info *types.PropertyInfo, prop *types.Zval, value_ptr *types.Zval, executeData *ZendExecuteData) *types.Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, executeData.IsCallUseStrictTypes()) == 0 {
		return UninitializedZval()
	}
	if prop.IsRef() {
		ZEND_REF_DEL_TYPE_SOURCE(prop.Ref(), prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZEND_REF_ADD_TYPE_SOURCE(prop.Ref(), prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *types.Zval, value_ptr *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) *types.Zval {
	faults.Error(faults.E_NOTICE, "Only variables should be assigned by reference")
	if EG__().HasException() {
		return UninitializedZval()
	}

	return ZendAssignToVariable(variable_ptr, value_ptr, executeData.IsCallUseStrictTypes())
}
func ZendFormatTypeEx(typ types.TypeHint) (part1 string, part2 string) {
	if typ.AllowNull() {
		part1 = "?"
	}
	if typ.IsClass() {
		if typ.IsCe() {
			part2 = typ.Ce().Name()
		} else {
			part2 = typ.Name()
		}
	} else {
		part2 = types.ZendGetTypeByConst(typ.Code())
	}
	return part1, part2
}

func ZendFormatType(type_ types.TypeHint, part1 **byte, part2 **byte) {
	if type_.AllowNull() {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if type_.IsClass() {
		if type_.IsCe() {
			*part2 = type_.Ce().Name()
		} else {
			*part2 = type_.Name()
		}
	} else {
		*part2 = types.ZendGetTypeByConst(type_.Code())
	}
}
func ZendThrowAutoInitInPropError(prop *types.PropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	faults.TypeError(fmt.Sprintf("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName()), prop_type1, prop_type2))
}
func ZendThrowAutoInitInRefError(prop *types.PropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	faults.TypeError(fmt.Sprintf("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName()), prop_type1, prop_type2))
}
func ZendThrowAccessUninitPropByRefError(prop *types.PropertyInfo) {
	faults.ThrowError(nil, fmt.Sprintf("Cannot access uninitialized non-nullable property %s::$%s by reference", prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName())))
}
func MakeRealObject(object *types.Zval, property *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) *types.Zval {
	var obj *types.Object
	var ref *types.Zval = nil
	if object.IsRef() {
		ref = object
		object = types.Z_REFVAL_P(object)
	}
	if object.Type() > types.IsFalse && (!object.IsString() || len(object.String()) != 0) {
		if opline.GetOp1Type() != IS_VAR || !(object.IsError()) {
			var property_name *types.String = operators.ZvalGetString(property)
			if opline.GetOpcode() == ZEND_PRE_INC_OBJ || opline.GetOpcode() == ZEND_PRE_DEC_OBJ || opline.GetOpcode() == ZEND_POST_INC_OBJ || opline.GetOpcode() == ZEND_POST_DEC_OBJ {
				faults.Error(faults.E_WARNING, fmt.Sprintf("Attempt to increment/decrement property '%s' of non-object", property_name.GetStr()))
			} else if opline.GetOpcode() == ZEND_FETCH_OBJ_W || opline.GetOpcode() == ZEND_FETCH_OBJ_RW || opline.GetOpcode() == ZEND_FETCH_OBJ_FUNC_ARG || opline.GetOpcode() == ZEND_ASSIGN_OBJ_REF {
				faults.Error(faults.E_WARNING, fmt.Sprintf("Attempt to modify property '%s' of non-object", property_name.GetStr()))
			} else {
				faults.Error(faults.E_WARNING, fmt.Sprintf("Attempt to assign property '%s' of non-object", property_name.GetStr()))
			}
			//ZendTmpStringRelease(tmp_property_name)
		}
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return nil
	}
	if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref.Ref()) {
		if zend_verify_ref_stdClass_assignable(ref.Ref()) == 0 {
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetUndef()
			}
			return nil
		}
	}
	ObjectInit(object)
	obj = object.Object()
	faults.Error(faults.E_WARNING, "Creating default object from empty value")
	return object
}
func ZendVerifyTypeErrorCommon(
	zf types.IFunction,
	arg_info *ZendArgInfo,
	ce *types.ClassEntry,
	value *types.Zval,
	fname **byte,
	fsep **byte,
	fclass **byte,
	need_msg **byte,
	need_kind **byte,
	need_or_null **byte,
	given_msg **byte,
	given_kind **byte,
) {
	var is_interface bool = 0
	*fname = zf.FunctionName()
	if zf.GetScope() != nil {
		*fsep = "::"
		*fclass = zf.GetScope().Name()
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
			*need_kind = ce.Name()
		} else {

			/* We don't know whether it's a class or interface, assume it's a class */

			*need_msg = "be an instance of "
			*need_kind = arg_info.GetType().Name()
		}
	} else {
		switch arg_info.GetType().Code() {
		case types.IsObject:
			*need_msg = "be an "
			*need_kind = "object"
		case types.IsCallable:
			*need_msg = "be callable"
			*need_kind = ""
		case types.IsIterable:
			*need_msg = "be iterable"
			*need_kind = ""
		default:
			*need_msg = "be of the type "
			*need_kind = types.ZendGetTypeByConst(arg_info.GetType().Code())
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
			*given_kind = types.Z_OBJCE_P(value).Name()
		} else {
			*given_msg = types.ZendZvalTypeName(value)
			*given_kind = ""
		}
	} else {
		*given_msg = "none"
		*given_kind = ""
	}
}
func ZendVerifyArgError(zf types.IFunction, arg_info *ZendArgInfo, arg_num int, ce *types.ClassEntry, value *types.Zval) {
	var ptr *ZendExecuteData = CurrEX().GetPrevExecuteData()
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	if EG__().HasException() {

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
				faults.TypeError(fmt.Sprintf("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ptr.GetFunc().GetOpArray().GetFilename(), ptr.GetOpline().GetLineno()))
			} else {
				faults.TypeError(fmt.Sprintf("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind))
			}
		} else {
			faults.TypeError(fmt.Sprintf("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind))
		}
	} else {
		ZendMissingArgError(ptr)
	}
}
func IsNullConstant(scope *types.ClassEntry, default_value *types.Zval) int {
	if default_value.IsConstantAst() {
		var constant types.Zval
		types.ZVAL_COPY(&constant, default_value)
		if ZvalUpdateConstantEx(&constant, scope) != types.SUCCESS {
			return 0
		}
		if constant.IsNull() {
			return 1
		}
		// ZvalPtrDtorNogc(&constant)
	}
	return 0
}

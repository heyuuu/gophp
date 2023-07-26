package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendObjectFetchPropertyTypeInfo(obj *types.Object, slot *types.Zval) *types.PropertyInfo {
	if !(ZEND_CLASS_HAS_TYPE_HINTS(obj.GetCe())) {
		return nil
	}

	/* Not a declared property */

	if slot < obj.GetPropertiesTable() || slot >= obj.GetPropertiesTable()+obj.GetCe().GetDefaultPropertiesCount() {
		return nil
	}
	return ZendGetTypedPropertyInfoForSlot(obj, slot)
}
func ZendHandleFetchObjFlags(result *types.Zval, ptr *types.Zval, obj *types.Object, prop_info *types.PropertyInfo, flags uint32) bool {
	switch flags {
	case ZEND_FETCH_DIM_WRITE:
		if PromotesToArray(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if CheckTypeArrayAssignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "array")
				if result != nil {
					result.IsError()
				}
				return 0
			}
		}
	case ZEND_FETCH_OBJ_WRITE:
		if PromotesToObject(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if check_type_stdClass_assignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "stdClass")
				if result != nil {
					result.IsError()
				}
				return 0
			}
		}
	case ZEND_FETCH_REF:
		if !ptr.IsRef() {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if ptr.IsUndef() {
				if !(prop_info.GetType().AllowNull()) {
					ZendThrowAccessUninitPropByRefError(prop_info)
					if result != nil {
						result.IsError()
					}
					return 0
				}
				ptr.SetNull()
			}
			ptr.SetNewRef(ptr)
			ZEND_REF_ADD_TYPE_SOURCE(ptr.Reference(), prop_info)
		}
	default:

	}
	return 1
}
func ZendFetchPropertyAddress(
	result *types.Zval,
	container *types.Zval,
	container_op_type uint32,
	prop_ptr *types.Zval,
	prop_op_type uint32,
	cache_slot *any,
	type_ int,
	flags uint32,
	init_undef bool,
	opline *types.ZendOp,
	executeData *ZendExecuteData,
) {
	var ptr *types.Zval
	if container_op_type != IS_UNUSED && !container.IsObject() {
		for {
			if container.IsRef() && types.Z_REFVAL_P(container).IsObject() {
				container = types.Z_REFVAL_P(container)
				break
			}
			if container_op_type == IS_CV && type_ != BP_VAR_W && container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}

			/* this should modify object only if it's empty */
			if type_ == BP_VAR_UNSET {
				result.SetNull()
				return
			}
			container = MakeRealObject(container, prop_ptr, opline, executeData)
			if container == nil {
				result.IsError()
				return
			}
			break
		}
	}
	if prop_op_type == IS_CONST && types.Z_OBJCE_P(container) == CACHED_PTR_EX(cache_slot) {
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.Object = container.Object()
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			ptr = OBJ_PROP(zobj, prop_offset)
			if ptr.IsNotUndef() {
				result.SetIndirect(ptr)
				if flags != 0 {
					var prop_info *types.PropertyInfo = CACHED_PTR_EX(cache_slot + 2)
					if prop_info != nil {
						ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags)
					}
				}
				return
			}
		} else if zobj.GetProperties() != nil {
			zobj.DupProperties()
			ptr = zobj.GetProperties().KeyFind(prop_ptr.String().GetStr())
			if ptr != nil {
				result.SetIndirect(ptr)
				return
			}
		}
	}
	ptr = container.Object().GetPropertyPtrEx(prop_ptr, type_)
	if nil == ptr {
		ptr = container.Object().ReadPropertyEx(prop_ptr, type_, result)
		if ptr == result {
			return
		}
		if EG__().GetException() != nil {
			result.IsError()
			return
		}
	} else if ptr.IsError() {
		result.IsError()
		return
	}
	result.SetIndirect(ptr)
	if flags != 0 {
		// todo check
		//var prop_info *types.PropertyInfo
		//if prop_op_type == IS_CONST {
		//	prop_info = CACHED_PTR_EX(cache_slot + 2)
		//	if prop_info != nil {
		//		if !ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags) {
		//			return
		//		}
		//	}
		//} else {
		//	if !ZendHandleFetchObjFlags(result, ptr, container.Object(), nil, flags) {
		//		return
		//	}
		//}
		if !ZendHandleFetchObjFlags(result, ptr, container.Object(), nil, flags) {
			return
		}
	}
	if init_undef && ptr.IsUndef() {
		ptr.SetNull()
	}
}
func ZendAssignToPropertyReference(
	container *types.Zval,
	container_op_type uint32,
	prop_ptr *types.Zval,
	prop_op_type uint32,
	value_ptr *types.Zval,
	opline *types.ZendOp,
	executeData *ZendExecuteData,
) {
	var variable types.Zval
	var variable_ptr *types.Zval = &variable
	var cache_addr *any = lang.CondF1(prop_op_type == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION) }, nil)
	ZendFetchPropertyAddress(variable_ptr, container, container_op_type, prop_ptr, prop_op_type, cache_addr, BP_VAR_W, 0, 0, opline, executeData)
	if variable_ptr.IsIndirect() {
		variable_ptr = variable_ptr.Indirect()
	}
	if variable_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if !variable.IsIndirect() {
		faults.ThrowError(nil, "Cannot assign by reference to overloaded object")
		variable_ptr = UninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && !(value_ptr.IsRef()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		var prop_info *types.PropertyInfo = nil
		if prop_op_type == IS_CONST {
			prop_info = (*types.PropertyInfo)(CACHED_PTR_EX(cache_addr + 2))
		} else {
			container = types.ZVAL_DEREF(container)
			prop_info = ZendObjectFetchPropertyTypeInfo(container.Object(), variable_ptr)
		}
		if prop_info != nil {
			variable_ptr = ZendAssignToTypedPropertyReference(prop_info, variable_ptr, value_ptr, executeData)
		} else {
			ZendAssignToVariableReference(variable_ptr, value_ptr)
		}
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), variable_ptr)
	}
}
func ZendAssignToPropertyReferenceThisConst(container *types.Zval, prop_ptr *types.Zval, value_ptr *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_CONST, value_ptr, opline, executeData)
}
func ZendAssignToPropertyReferenceVarConst(container *types.Zval, prop_ptr *types.Zval, value_ptr *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_CONST, value_ptr, opline, executeData)
}
func ZendAssignToPropertyReferenceThisVar(container *types.Zval, prop_ptr *types.Zval, value_ptr *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_VAR, value_ptr, opline, executeData)
}
func ZendAssignToPropertyReferenceVarVar(container *types.Zval, prop_ptr *types.Zval, value_ptr *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_VAR, value_ptr, opline, executeData)
}
func ZendFetchStaticPropertyAddressEx(
	retval **types.Zval,
	prop_info **types.PropertyInfo,
	cache_slot uint32,
	fetch_type int,
	opline *types.ZendOp,
	executeData *ZendExecuteData,
) int {
	var free_op1 ZendFreeOp
	var name *types.String
	var tmp_name *types.String
	var ce *types.ClassEntry
	var property_info *types.PropertyInfo
	var op1_type uint8 = opline.GetOp1Type()
	var op2_type uint8 = opline.GetOp2Type()
	if op2_type == IS_CONST {
		var class_name *types.Zval = opline.Const2()
		b.Assert(op1_type != IS_CONST || CACHED_PTR(cache_slot) == nil)
		if lang.Assign(&ce, CACHED_PTR(cache_slot)) == nil {
			ce = ZendFetchClassByName(class_name.String(), (class_name + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return types.FAILURE
			}
			if op1_type != IS_CONST {
				CACHE_PTR(cache_slot, ce)
			}
		}
	} else {
		if op2_type == IS_UNUSED {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return types.FAILURE
			}
		} else {
			ce = opline.Op2().Class()
		}
		if op1_type == IS_CONST && CACHED_PTR(cache_slot) == ce {
			*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
			*prop_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
			return types.SUCCESS
		}
	}
	if op1_type == IS_CONST {
		name = opline.Const1().String()
	} else {
		var varname *types.Zval = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		if varname.IsString() {
			name = varname.String()
			tmp_name = nil
		} else {
			if op1_type == IS_CV && varname.IsUndef() {
				ZvalUndefinedCv(opline.GetOp1().GetVar(), executeData)
			}
			name = operators.ZvalGetString(varname)
		}
	}
	*retval = ZendStdGetStaticPropertyWithInfo(ce, name, fetch_type, &property_info)
	if op1_type != IS_CONST {
		//ZendTmpStringRelease(tmp_name)
		if op1_type != IS_CV {
			// ZvalPtrDtorNogc(free_op1)
		}
	}
	if (*retval) == nil {
		return types.FAILURE
	}
	*prop_info = property_info
	if op1_type == IS_CONST {
		CACHE_POLYMORPHIC_PTR(cache_slot, ce, *retval)
		CACHE_PTR(cache_slot+b.SizeOf("void *")*2, property_info)
	}
	return types.SUCCESS
}
func ZendFetchStaticPropertyAddress(
	retval **types.Zval,
	prop_info **types.PropertyInfo,
	cache_slot uint32,
	fetch_type int,
	flags int,
	opline *types.ZendOp,
	executeData *ZendExecuteData,
) int {
	var success int
	var property_info *types.PropertyInfo
	if opline.GetOp1Type() == IS_CONST && (opline.GetOp2Type() == IS_CONST || opline.GetOp2Type() == IS_UNUSED && (opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_SELF || opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_PARENT)) && CACHED_PTR(cache_slot) != nil {
		*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
		property_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
		if (fetch_type == BP_VAR_R || fetch_type == BP_VAR_RW) && retval.IsUndef() && property_info.GetType() != 0 {
			faults.ThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().Name(), ZendGetUnmangledPropertyNameEx(property_info.GetName()))
			return types.FAILURE
		}
	} else {
		success = ZendFetchStaticPropertyAddressEx(retval, &property_info, cache_slot, fetch_type, opline, executeData)
		if success != types.SUCCESS {
			return types.FAILURE
		}
	}
	if flags != 0 && property_info.GetType() != 0 {
		ZendHandleFetchObjFlags(nil, *retval, nil, property_info, flags)
	}
	if prop_info != nil {
		*prop_info = property_info
	}
	return types.SUCCESS
}
func ZendThrowRefTypeErrorType(prop1 *types.PropertyInfo, prop2 *types.PropertyInfo, zv *types.Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	faults.TypeError("Reference with value of type %s held by property %s::$%s of type %s%s is not compatible with property %s::$%s of type %s%s", lang.CondF(zv.IsObject(), func() []byte { return types.Z_OBJCE_P(zv).Name() }, func() *byte { return types.ZendGetTypeByConst(zv.Type()) }), prop1.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop2.GetName()), prop2_type1, prop2_type2)
}
func ZendThrowRefTypeErrorZval(prop *types.PropertyInfo, zv *types.Zval) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	faults.TypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s", lang.CondF(zv.IsObject(), func() []byte { return types.Z_OBJCE_P(zv).Name() }, func() *byte { return types.ZendGetTypeByConst(zv.Type()) }), prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowConflictingCoercionError(prop1 *types.PropertyInfo, prop2 *types.PropertyInfo, zv *types.Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	faults.TypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s and property %s::$%s of type %s%s, as this would result in an inconsistent type conversion", lang.CondF(zv.IsObject(), func() []byte { return types.Z_OBJCE_P(zv).Name() }, func() *byte { return types.ZendGetTypeByConst(zv.Type()) }), prop1.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop2.GetName()), prop2_type1, prop2_type2)
}
func IZendVerifyTypeAssignableZval(type_ptr *types.TypeHint, self_ce *types.ClassEntry, zv *types.Zval, strict bool) int {
	var type_ types.TypeHint = *type_ptr
	var type_code uint8
	var zv_type uint8 = zv.Type()
	if type_.AllowNull() && zv_type == types.IsNull {
		return 1
	}
	if type_.IsClass() {
		if !(type_.IsCe()) {
			if !ZendResolveClassType(type_ptr, self_ce) {
				return 0
			}
			type_ = *type_ptr
		}
		return zv_type == types.IsObject && operators.InstanceofFunction(types.Z_OBJCE_P(zv), type_.Ce()) != 0
	}
	type_code = type_.Code()
	if type_code == zv_type || type_code == types.IsBool && (zv_type == types.IsFalse || zv_type == types.IsTrue) {
		return 1
	}
	if type_code == types.IsIterable {
		return ZendIsIterable(zv)
	}

	/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	if strict != 0 {
		if type_code == types.IsDouble && zv_type == types.IsLong {
			return -1
		}
		return 0
	}

	/* No weak conversions for arrays and objects */

	if type_code == types.IsArray || type_code == types.IsObject {
		return 0
	}

	/* NULL may be accepted only by nullable hints (this is already checked) */

	if zv_type == types.IsNull {
		return 0
	}

	/* Coercion may be necessary, check separately */

	return -1

	/* Coercion may be necessary, check separately */
}
func ZendVerifyRefAssignableZval(ref *types.Reference, zv *types.Zval, strict bool) bool {
	var prop *types.PropertyInfo

	/* The value must satisfy each property type, and coerce to the same value for each property
	 * type. Right now, the latter rule means that *if* coercion is necessary, then all types
	 * must be the same (modulo nullability). To handle this, remember the first type we see and
	 * compare against it when coercion becomes necessary. */

	var seen_prop *types.PropertyInfo = nil
	var seen_type uint8
	var needs_coercion bool = 0
	b.Assert(!zv.IsRef())
	var _source_list *types.ZendPropertyInfoSourceList = &(ref.GetSources())
	var _prop **types.PropertyInfo
	var _end ***types.PropertyInfo
	var _list *types.ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if types.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = types.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
			_prop = _list.GetPtr()
			_end = _list.GetPtr() + _list.GetNum()
		} else {
			_prop = _source_list.GetPtr()
			_end = _prop + 1
		}
		for ; _prop < _end; _prop++ {
			prop = *_prop
			var result int = IZendVerifyTypeAssignableZval(prop.GetType(), prop.GetCe(), zv, strict)
			if result == 0 {
				ZendThrowRefTypeErrorZval(prop, zv)
				return 0
			}
			if result < 0 {
				needs_coercion = 1
			}
			if seen_prop == nil {
				seen_prop = prop
				if prop.GetType().IsClass() {
					seen_type = types.IsObject
				} else {
					seen_type = prop.GetType().Code()
				}
			} else if needs_coercion != 0 && seen_type != prop.GetType().Code() {
				ZendThrowConflictingCoercionError(seen_prop, prop, zv)
				return 0
			}
		}
	}
	if needs_coercion != 0 && ZendVerifyWeakScalarTypeHint(seen_type, zv) == 0 {
		ZendThrowRefTypeErrorZval(seen_prop, zv)
		return 0
	}
	return 1
}

// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendObjectFetchPropertyTypeInfo(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	if !(ZEND_CLASS_HAS_TYPE_HINTS(obj.GetCe())) {
		return nil
	}

	/* Not a declared property */

	if slot < obj.GetPropertiesTable() || slot >= obj.GetPropertiesTable()+obj.GetCe().GetDefaultPropertiesCount() {
		return nil
	}
	return ZendGetTypedPropertyInfoForSlot(obj, slot)
}
func ZendHandleFetchObjFlags(result *Zval, ptr *Zval, obj *ZendObject, prop_info *ZendPropertyInfo, flags uint32) ZendBool {
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
		if ptr.GetType() != IS_REFERENCE {
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
			ZEND_REF_ADD_TYPE_SOURCE(ptr.GetRef(), prop_info)
		}
	default:

	}
	return 1
}
func ZendFetchPropertyAddress(
	result *Zval,
	container *Zval,
	container_op_type uint32,
	prop_ptr *Zval,
	prop_op_type uint32,
	cache_slot *any,
	type_ int,
	flags uint32,
	init_undef ZendBool,
	opline *ZendOp,
	_ EXECUTE_DATA_D,
) {
	var ptr *Zval
	if container_op_type != IS_UNUSED && container.GetType() != IS_OBJECT {
		for {
			if container.IsReference() && Z_REFVAL_P(container).IsObject() {
				container = Z_REFVAL_P(container)
				break
			}
			if container_op_type == IS_CV && type_ != BP_VAR_W && container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}

			/* this should modify object only if it's empty */

			if type_ == BP_VAR_UNSET {
				result.SetNull()
				return
			}
			container = MakeRealObject(container, prop_ptr, OPLINE_C, EXECUTE_DATA_C)
			if container == nil {
				result.IsError()
				return
			}
			break
		}
	}
	if prop_op_type == IS_CONST && Z_OBJCE_P(container) == CACHED_PTR_EX(cache_slot) {
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *ZendObject = container.GetObj()
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			ptr = OBJ_PROP(zobj, prop_offset)
			if ptr.GetType() != IS_UNDEF {
				result.SetIndirect(ptr)
				if flags != 0 {
					var prop_info *ZendPropertyInfo = CACHED_PTR_EX(cache_slot + 2)
					if prop_info != nil {
						ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags)
					}
				}
				return
			}
		} else if zobj.GetProperties() != nil {
			if zobj.GetProperties().GetRefcount() > 1 {
				if (zobj.GetProperties().GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
					zobj.GetProperties().DelRefcount()
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			ptr = zobj.GetProperties().KeyFind(prop_ptr.GetStr().GetStr())
			if ptr != nil {
				result.SetIndirect(ptr)
				return
			}
		}
	}
	ptr = Z_OBJ_HT_P(container).GetGetPropertyPtrPtr()(container, prop_ptr, type_, cache_slot)
	if nil == ptr {
		ptr = Z_OBJ_HT_P(container).GetReadProperty()(container, prop_ptr, type_, cache_slot, result)
		if ptr == result {
			if ptr.IsReference() && ptr.GetRefcount() == 1 {
				ZVAL_UNREF(ptr)
			}
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
		var prop_info *ZendPropertyInfo
		if prop_op_type == IS_CONST {
			prop_info = CACHED_PTR_EX(cache_slot + 2)
			if prop_info != nil {
				if ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags) == 0 {
					return
				}
			}
		} else {
			if ZendHandleFetchObjFlags(result, ptr, container.GetObj(), nil, flags) == 0 {
				return
			}
		}
	}
	if init_undef != 0 && ptr.IsUndef() {
		ptr.SetNull()
	}
}
func ZendAssignToPropertyReference(
	container *Zval,
	container_op_type uint32,
	prop_ptr *Zval,
	prop_op_type uint32,
	value_ptr *Zval,
	opline *ZendOp,
	_ EXECUTE_DATA_D,
) {
	var variable Zval
	var variable_ptr *Zval = &variable
	var cache_addr *any = b.CondF1(prop_op_type == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION) }, nil)
	ZendFetchPropertyAddress(variable_ptr, container, container_op_type, prop_ptr, prop_op_type, cache_addr, BP_VAR_W, 0, 0, OPLINE_C, EXECUTE_DATA_C)
	if variable_ptr.IsIndirect() {
		variable_ptr = variable_ptr.GetZv()
	}
	if variable_ptr.IsError() {
		variable_ptr = EG__().GetUninitializedZval()
	} else if variable.GetType() != IS_INDIRECT {
		ZendThrowError(nil, "Cannot assign by reference to overloaded object")
		ZvalPtrDtor(&variable)
		variable_ptr = EG__().GetUninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = EG__().GetUninitializedZval()
	} else if (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && !(value_ptr.IsReference()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, OPLINE_C, EXECUTE_DATA_C)
	} else {
		var prop_info *ZendPropertyInfo = nil
		if prop_op_type == IS_CONST {
			prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_addr + 2))
		} else {
			ZVAL_DEREF(container)
			prop_info = ZendObjectFetchPropertyTypeInfo(container.GetObj(), variable_ptr)
		}
		if prop_info != nil {
			variable_ptr = ZendAssignToTypedPropertyReference(prop_info, variable_ptr, value_ptr, EXECUTE_DATA_C)
		} else {
			ZendAssignToVariableReference(variable_ptr, value_ptr)
		}
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), variable_ptr)
	}
}
func ZendAssignToPropertyReferenceThisConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceThisVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendFetchStaticPropertyAddressEx(
	retval **Zval,
	prop_info **ZendPropertyInfo,
	cache_slot uint32,
	fetch_type int,
	opline *ZendOp,
	_ EXECUTE_DATA_D,
) int {
	var free_op1 ZendFreeOp
	var name *ZendString
	var tmp_name *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var op1_type ZendUchar = opline.GetOp1Type()
	var op2_type ZendUchar = opline.GetOp2Type()
	if op2_type == IS_CONST {
		var class_name *Zval = RT_CONSTANT(opline, opline.GetOp2())
		ZEND_ASSERT(op1_type != IS_CONST || CACHED_PTR(cache_slot) == nil)
		if b.Assign(&ce, CACHED_PTR(cache_slot)) == nil {
			ce = ZendFetchClassByName(class_name.GetStr(), (class_name + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return FAILURE
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
				return FAILURE
			}
		} else {
			ce = EX_VAR(opline.GetOp2().GetVar()).GetCe()
		}
		if op1_type == IS_CONST && CACHED_PTR(cache_slot) == ce {
			*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
			*prop_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
			return SUCCESS
		}
	}
	if op1_type == IS_CONST {
		name = RT_CONSTANT(opline, opline.GetOp1()).GetStr()
	} else {
		var varname *Zval = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		if varname.IsString() {
			name = varname.GetStr()
			tmp_name = nil
		} else {
			if op1_type == IS_CV && varname.IsUndef() {
				ZvalUndefinedCv(opline.GetOp1().GetVar(), EXECUTE_DATA_C)
			}
			name = ZvalGetTmpString(varname, &tmp_name)
		}
	}
	*retval = ZendStdGetStaticPropertyWithInfo(ce, name, fetch_type, &property_info)
	if op1_type != IS_CONST {
		ZendTmpStringRelease(tmp_name)
		if op1_type != IS_CV {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	if (*retval) == nil {
		return FAILURE
	}
	*prop_info = property_info
	if op1_type == IS_CONST {
		CACHE_POLYMORPHIC_PTR(cache_slot, ce, *retval)
		CACHE_PTR(cache_slot+b.SizeOf("void *")*2, property_info)
	}
	return SUCCESS
}
func ZendFetchStaticPropertyAddress(
	retval **Zval,
	prop_info **ZendPropertyInfo,
	cache_slot uint32,
	fetch_type int,
	flags int,
	opline *ZendOp,
	_ EXECUTE_DATA_D,
) int {
	var success int
	var property_info *ZendPropertyInfo
	if opline.GetOp1Type() == IS_CONST && (opline.GetOp2Type() == IS_CONST || opline.GetOp2Type() == IS_UNUSED && (opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_SELF || opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_PARENT)) && CACHED_PTR(cache_slot) != nil {
		*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
		property_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
		if (fetch_type == BP_VAR_R || fetch_type == BP_VAR_RW) && retval.IsUndef() && property_info.GetType() != 0 {
			ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(property_info.GetName()))
			return FAILURE
		}
	} else {
		success = ZendFetchStaticPropertyAddressEx(retval, &property_info, cache_slot, fetch_type, OPLINE_C, EXECUTE_DATA_C)
		if success != SUCCESS {
			return FAILURE
		}
	}
	if flags != 0 && property_info.GetType() != 0 {
		ZendHandleFetchObjFlags(nil, *retval, nil, property_info, flags)
	}
	if prop_info != nil {
		*prop_info = property_info
	}
	return SUCCESS
}
func ZendThrowRefTypeErrorType(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Reference with value of type %s held by property %s::$%s of type %s%s is not compatible with property %s::$%s of type %s%s", b.CondF(zv.IsObject(), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func ZendThrowRefTypeErrorZval(prop *ZendPropertyInfo, zv *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s", b.CondF(zv.IsObject(), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowConflictingCoercionError(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s and property %s::$%s of type %s%s, as this would result in an inconsistent type conversion", b.CondF(zv.IsObject(), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func IZendVerifyTypeAssignableZval(type_ptr *ZendType, self_ce *ZendClassEntry, zv *Zval, strict ZendBool) int {
	var type_ ZendType = *type_ptr
	var type_code ZendUchar
	var zv_type ZendUchar = zv.GetType()
	if type_.AllowNull() && zv_type == IS_NULL {
		return 1
	}
	if type_.IsClass() {
		if !(type_.IsCe()) {
			if ZendResolveClassType(type_ptr, self_ce) == 0 {
				return 0
			}
			type_ = *type_ptr
		}
		return zv_type == IS_OBJECT && InstanceofFunction(Z_OBJCE_P(zv), type_.Ce()) != 0
	}
	type_code = type_.Code()
	if type_code == zv_type || type_code == _IS_BOOL && (zv_type == IS_FALSE || zv_type == IS_TRUE) {
		return 1
	}
	if type_code == IS_ITERABLE {
		return ZendIsIterable(zv)
	}

	/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	if strict != 0 {
		if type_code == IS_DOUBLE && zv_type == IS_LONG {
			return -1
		}
		return 0
	}

	/* No weak conversions for arrays and objects */

	if type_code == IS_ARRAY || type_code == IS_OBJECT {
		return 0
	}

	/* NULL may be accepted only by nullable hints (this is already checked) */

	if zv_type == IS_NULL {
		return 0
	}

	/* Coercion may be necessary, check separately */

	return -1

	/* Coercion may be necessary, check separately */
}
func ZendVerifyRefAssignableZval(ref *ZendReference, zv *Zval, strict ZendBool) ZendBool {
	var prop *ZendPropertyInfo

	/* The value must satisfy each property type, and coerce to the same value for each property
	 * type. Right now, the latter rule means that *if* coercion is necessary, then all types
	 * must be the same (modulo nullability). To handle this, remember the first type we see and
	 * compare against it when coercion becomes necessary. */

	var seen_prop *ZendPropertyInfo = nil
	var seen_type ZendUchar
	var needs_coercion ZendBool = 0
	ZEND_ASSERT(zv.GetType() != IS_REFERENCE)
	var _source_list *ZendPropertyInfoSourceList = &(ref.GetSources())
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
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
					seen_type = IS_OBJECT
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

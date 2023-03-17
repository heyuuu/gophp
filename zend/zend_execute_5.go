// <<generate>>

package zend

import (
	b "sik/builtin"
)

func zend_fetch_dimension_address_W(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_RW(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_RW, executeData)
}
func zend_fetch_dimension_address_UNSET(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_UNSET, executeData)
}
func ZendFetchDimensionAddressRead(
	result *Zval,
	container *Zval,
	dim *Zval,
	dim_type int,
	type_ int,
	is_list int,
	slow int,
	executeData *ZendExecuteData,
) {
	var retval *Zval
	if slow == 0 {
		if container.IsArray() {
		try_array:
			retval = ZendFetchDimensionAddressInner(container.GetArr(), dim, dim_type, type_, executeData)
			ZVAL_COPY_DEREF(result, retval)
			return
		} else if container.IsReference() {
			container = Z_REFVAL_P(container)
			if container.IsArray() {
				goto try_array
			}
		}
	}
	if is_list == 0 && container.IsString() {
		var offset ZendLong
	try_string_offset:
		if dim.GetType() != IS_LONG {
			switch dim.GetType() {
			case IS_STRING:
				if IS_LONG == IsNumericString(dim.GetStr().GetVal(), dim.GetStr().GetLen(), nil, nil, -1) {
					break
				}
				if type_ == BP_VAR_IS {
					result.SetNull()
					return
				}
				ZendError(E_WARNING, "Illegal string offset '%s'", dim.GetStr().GetVal())
			case IS_UNDEF:
				ZVAL_UNDEFINED_OP2()
				fallthrough
			case IS_DOUBLE:
				fallthrough
			case IS_NULL:
				fallthrough
			case IS_FALSE:
				fallthrough
			case IS_TRUE:
				if type_ != BP_VAR_IS {
					ZendError(E_NOTICE, "String offset cast occurred")
				}
			case IS_REFERENCE:
				dim = Z_REFVAL_P(dim)
				goto try_string_offset
			default:
				ZendIllegalOffset()
			}
			offset = ZvalGetLongFunc(dim)
		} else {
			offset = dim.GetLval()
		}
		if container.GetStr().GetLen() < b.CondF(offset < 0, func() int { return -int(offset) }, func() int { return int(offset + 1) }) {
			if type_ != BP_VAR_IS {
				ZendError(E_NOTICE, "Uninitialized string offset: "+ZEND_LONG_FMT, offset)
				ZVAL_EMPTY_STRING(result)
			} else {
				result.SetNull()
			}
		} else {
			var c ZendUchar
			var real_offset ZendLong
			if offset < 0 {
				real_offset = ZendLong(container.GetStr().GetLen() + offset)
			} else {
				real_offset = offset
			}
			c = ZendUchar(container.GetStr().GetVal()[real_offset])
			result.SetInternedString(ZSTR_CHAR(c))
		}
	} else if container.IsObject() {
		if dim.IsUndef() {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		ZEND_ASSERT(result != nil)
		if retval != nil {
			if result != retval {
				ZVAL_COPY_DEREF(result, retval)
			} else if retval.IsReference() {
				ZendUnwrapReference(result)
			}
		} else {
			result.SetNull()
		}
	} else {
		if type_ != BP_VAR_IS && container.IsUndef() {
			container = ZVAL_UNDEFINED_OP1()
		}
		if dim.IsUndef() {
			ZVAL_UNDEFINED_OP2()
		}
		if is_list == 0 && type_ != BP_VAR_IS {
			ZendError(E_NOTICE, "Trying to access array offset on value of type %s", ZendZvalTypeName(container))
		}
		result.SetNull()
	}
}
func zend_fetch_dimension_address_read_R(container *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 0, 0, executeData)
}
func zend_fetch_dimension_address_read_R_slow(container *Zval, dim *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, IS_CV, BP_VAR_R, 0, 1, executeData)
}
func zend_fetch_dimension_address_read_IS(container *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_IS, 0, 0, executeData)
}
func zend_fetch_dimension_address_LIST_r(container *Zval, dim *Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 1, 0, executeData)
}
func ZendFetchDimensionConst(result *Zval, container *Zval, dim *Zval, type_ int) {
	ZendFetchDimensionAddressRead(result, container, dim, IS_TMP_VAR, type_, 0, 0, nil)
}
func ZendFindArrayDimSlow(ht *HashTable, offset *Zval, executeData *ZendExecuteData) *Zval {
	var hval ZendUlong
	if offset.IsDouble() {
		hval = ZendDvalToLval(offset.GetDval())
	num_idx:
		return ht.IndexFindH(hval)
	} else if offset.IsNull() {
	str_idx:
		return ZendHashFindExInd(ht, ZSTR_EMPTY_ALLOC(), 1)
	} else if offset.IsFalse() {
		hval = 0
		goto num_idx
	} else if offset.IsTrue() {
		hval = 1
		goto num_idx
	} else if offset.IsResource() {
		hval = Z_RES_HANDLE_P(offset)
		goto num_idx
	} else if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2()
		goto str_idx
	} else {
		ZendError(E_WARNING, "Illegal offset type in isset or empty")
		return nil
	}
}
func ZendIssetDimSlow(container *Zval, offset *Zval, executeData *ZendExecuteData) int {
	if offset.IsUndef() {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if container.IsObject() {
		return Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 0)
	} else if container.IsString() {
		var lval ZendLong
		if offset.IsLong() {
			lval = offset.GetLval()
		str_offset:
			if lval < 0 {
				lval += ZendLong(container.GetStr().GetLen())
			}
			if lval >= 0 && int(lval < container.GetStr().GetLen()) != 0 {
				return 1
			} else {
				return 0
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			offset = ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < IS_STRING || offset.IsString() && IS_LONG == IsNumericString(offset.GetStr().GetVal(), offset.GetStr().GetLen(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 0
		}
	} else {
		return 0
	}
}
func ZendIsemptyDimSlow(container *Zval, offset *Zval, executeData *ZendExecuteData) int {
	if offset.IsUndef() {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if container.IsObject() {
		return !(Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 1))
	} else if container.IsString() {
		var lval ZendLong
		if offset.IsLong() {
			lval = offset.GetLval()
		str_offset:
			if lval < 0 {
				lval += ZendLong(container.GetStr().GetLen())
			}
			if lval >= 0 && int(lval < container.GetStr().GetLen()) != 0 {
				return container.GetStr().GetVal()[lval] == '0'
			} else {
				return 1
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			offset = ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < IS_STRING || offset.IsString() && IS_LONG == IsNumericString(offset.GetStr().GetVal(), offset.GetStr().GetLen(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 1
		}
	} else {
		return 1
	}
}
func ZendArrayKeyExistsFast(ht *HashTable, key *Zval, opline *ZendOp, executeData *ZendExecuteData) uint32 {
	var str *ZendString
	var hval ZendUlong
try_again:
	if key.IsString() {
		str = key.GetStr()
		if ZEND_HANDLE_NUMERIC(str, &hval) {
			goto num_key
		}
	str_key:
		if ZendHashFindInd(ht, str) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if key.IsLong() {
		hval = key.GetLval()
	num_key:
		if ht.IndexFindH(hval) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if key.IsReference() {
		key = Z_REFVAL_P(key)
		goto try_again
	} else if key.GetType() <= IS_NULL {
		if key.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		str = ZSTR_EMPTY_ALLOC()
		goto str_key
	} else {
		ZendError(E_WARNING, "array_key_exists(): The first argument should be either a string or an integer")
		return IS_FALSE
	}
}
func ZendArrayKeyExistsSlow(subject *Zval, key *Zval, opline *ZendOp, executeData *ZendExecuteData) uint32 {
	if subject.IsObject() {
		ZendError(E_DEPRECATED, "array_key_exists(): "+"Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
		var ht *HashTable = ZendGetPropertiesFor(subject, ZEND_PROP_PURPOSE_ARRAY_CAST)
		var result uint32 = ZendArrayKeyExistsFast(ht, key, opline, executeData)
		ZendReleaseProperties(ht)
		return result
	} else {
		if key.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		if subject.GetTypeInfo() == IS_UNDEF {
			ZVAL_UNDEFINED_OP2()
		}
		ZendInternalTypeError(executeData.IsCallUseStrictTypes(), "array_key_exists() expects parameter 2 to be array, %s given", ZendGetTypeByConst(subject.GetType()))
		return IS_NULL
	}
}
func PromotesToArray(val *Zval) ZendBool {
	return val.GetType() <= IS_FALSE || val.IsReference() && Z_REFVAL_P(val).GetType() <= IS_FALSE
}
func PromotesToObject(val *Zval) ZendBool {
	val = ZVAL_DEREF(val)
	return val.GetType() <= IS_FALSE || val.IsString() && val.GetStr().GetLen() == 0
}
func CheckTypeArrayAssignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	return type_.IsCode() && (type_.Code() == IS_ARRAY || type_.Code() == IS_ITERABLE)
}
func check_type_stdClass_assignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	if type_.IsClass() {
		if type_.IsCe() {
			return type_.Ce() == ZendStandardClassDef
		} else {
			return ZendStringEqualsLiteralCi(type_.Name(), "stdclass")
		}
	} else {
		return type_.Code() == IS_OBJECT
	}
}
func ZendVerifyRefArrayAssignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
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
			if CheckTypeArrayAssignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "array")
				return 0
			}
		}
	}
	return 1
}
func zend_verify_ref_stdClass_assignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
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
			if check_type_stdClass_assignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "stdClass")
				return 0
			}
		}
	}
	return 1
}

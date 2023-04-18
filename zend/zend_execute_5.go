package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func zend_fetch_dimension_address_W(container_ptr *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_RW(container_ptr *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_RW, executeData)
}
func zend_fetch_dimension_address_UNSET(container_ptr *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_UNSET, executeData)
}
func ZendFetchDimensionAddressRead(
	result *types2.Zval,
	container *types2.Zval,
	dim *types2.Zval,
	dim_type int,
	type_ int,
	is_list int,
	slow int,
	executeData *ZendExecuteData,
) {
	var retval *types2.Zval
	if slow == 0 {
		if container.IsArray() {
		try_array:
			retval = ZendFetchDimensionAddressInner(container.Array(), dim, dim_type, type_, executeData)
			types2.ZVAL_COPY_DEREF(result, retval)
			return
		} else if container.IsReference() {
			container = types2.Z_REFVAL_P(container)
			if container.IsArray() {
				goto try_array
			}
		}
	}
	if is_list == 0 && container.IsString() {
		var offset ZendLong
	try_string_offset:
		if dim.GetType() != types2.IS_LONG {
			switch dim.GetType() {
			case types2.IS_STRING:
				if types2.IS_LONG == IsNumericString(dim.String().GetStr(), nil, nil, -1) {
					break
				}
				if type_ == BP_VAR_IS {
					result.SetNull()
					return
				}
				faults.Error(faults.E_WARNING, "Illegal string offset '%s'", dim.String().GetVal())
			case types2.IS_UNDEF:
				ZVAL_UNDEFINED_OP2(executeData)
				fallthrough
			case types2.IS_DOUBLE:
				fallthrough
			case types2.IS_NULL:
				fallthrough
			case types2.IS_FALSE:
				fallthrough
			case types2.IS_TRUE:
				if type_ != BP_VAR_IS {
					faults.Error(faults.E_NOTICE, "String offset cast occurred")
				}
			case types2.IS_REFERENCE:
				dim = types2.Z_REFVAL_P(dim)
				goto try_string_offset
			default:
				ZendIllegalOffset()
			}
			offset = ZvalGetLongFunc(dim)
		} else {
			offset = dim.Long()
		}
		if container.String().GetLen() < b.CondF(offset < 0, func() int { return -int(offset) }, func() int { return int(offset + 1) }) {
			if type_ != BP_VAR_IS {
				faults.Error(faults.E_NOTICE, "Uninitialized string offset: "+ZEND_LONG_FMT, offset)
				result.SetStringVal("")
			} else {
				result.SetNull()
			}
		} else {
			var c types2.ZendUchar
			var real_offset ZendLong
			if offset < 0 {
				real_offset = ZendLong(container.String().GetLen() + offset)
			} else {
				real_offset = offset
			}
			c = types2.ZendUchar(container.String().GetStr()[real_offset])
			result.SetStringVal(string(c))
		}
	} else if container.IsObject() {
		if dim.IsUndef() {
			dim = ZVAL_UNDEFINED_OP2(executeData)
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = types2.Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		b.Assert(result != nil)
		if retval != nil {
			if result != retval {
				types2.ZVAL_COPY_DEREF(result, retval)
			} else if retval.IsReference() {
				ZendUnwrapReference(result)
			}
		} else {
			result.SetNull()
		}
	} else {
		if type_ != BP_VAR_IS && container.IsUndef() {
			container = ZVAL_UNDEFINED_OP1(executeData)
		}
		if dim.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		if is_list == 0 && type_ != BP_VAR_IS {
			faults.Error(faults.E_NOTICE, "Trying to access array offset on value of type %s", types2.ZendZvalTypeName(container))
		}
		result.SetNull()
	}
}
func zend_fetch_dimension_address_read_R(container *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 0, 0, executeData)
}
func zend_fetch_dimension_address_read_R_slow(container *types2.Zval, dim *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddressRead(result, container, dim, IS_CV, BP_VAR_R, 0, 1, executeData)
}
func zend_fetch_dimension_address_read_IS(container *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_IS, 0, 0, executeData)
}
func zend_fetch_dimension_address_LIST_r(container *types2.Zval, dim *types2.Zval, dim_type int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types2.Zval = opline.Result()
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 1, 0, executeData)
}
func ZendFetchDimensionConst(result *types2.Zval, container *types2.Zval, dim *types2.Zval, type_ int) {
	ZendFetchDimensionAddressRead(result, container, dim, IS_TMP_VAR, type_, 0, 0, nil)
}
func ZendFindArrayDimSlow(ht *types2.Array, offset *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	var hval ZendUlong
	if offset.IsDouble() {
		hval = DvalToLval(offset.Double())
	num_idx:
		return ht.IndexFind(hval)
	} else if offset.IsNull() {
	str_idx:
		return types2.ZendHashFindInd(ht, types2.NewString("").GetStr())
	} else if offset.IsFalse() {
		hval = 0
		goto num_idx
	} else if offset.IsTrue() {
		hval = 1
		goto num_idx
	} else if offset.IsResource() {
		hval = types2.Z_RES_HANDLE_P(offset)
		goto num_idx
	} else if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2(executeData)
		goto str_idx
	} else {
		faults.Error(faults.E_WARNING, "Illegal offset type in isset or empty")
		return nil
	}
}
func ZendIssetDimSlow(container *types2.Zval, offset *types2.Zval, executeData *ZendExecuteData) int {
	if offset.IsUndef() {
		offset = ZVAL_UNDEFINED_OP2(executeData)
	}
	if container.IsObject() {
		return types2.Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 0)
	} else if container.IsString() {
		var lval ZendLong
		if offset.IsLong() {
			lval = offset.Long()
		str_offset:
			if lval < 0 {
				lval += ZendLong(container.String().GetLen())
			}
			if lval >= 0 && int(lval < container.String().GetLen()) != 0 {
				return 1
			} else {
				return 0
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			offset = types2.ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < types2.IS_STRING || offset.IsString() && types2.IS_LONG == IsNumericString(offset.String().GetStr(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 0
		}
	} else {
		return 0
	}
}
func ZendIsemptyDimSlow(container *types2.Zval, offset *types2.Zval, executeData *ZendExecuteData) int {
	if offset.IsUndef() {
		offset = ZVAL_UNDEFINED_OP2(executeData)
	}
	if container.IsObject() {
		return !(types2.Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 1))
	} else if container.IsString() {
		var lval ZendLong
		if offset.IsLong() {
			lval = offset.Long()
		str_offset:
			if lval < 0 {
				lval += ZendLong(container.String().GetLen())
			}
			if lval >= 0 && int(lval < container.String().GetLen()) != 0 {
				return container.String().GetStr()[lval] == '0'
			} else {
				return 1
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			offset = types2.ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < types2.IS_STRING || offset.IsString() && types2.IS_LONG == IsNumericString(offset.String().GetStr(), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 1
		}
	} else {
		return 1
	}
}
func ZendArrayKeyExistsFast(ht *types2.Array, key *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) uint32 {
	var str *types2.String
	var hval ZendUlong
try_again:
	if key.IsString() {
		str = key.String()
		if types2.HandleNumericStr(str.GetStr(), &hval) {
			goto num_key
		}
	str_key:
		if types2.ZendHashFindInd(ht, str.GetStr()) != nil {
			return types2.IS_TRUE
		} else {
			return types2.IS_FALSE
		}
	} else if key.IsLong() {
		hval = key.Long()
	num_key:
		if ht.IndexFind(hval) != nil {
			return types2.IS_TRUE
		} else {
			return types2.IS_FALSE
		}
	} else if key.IsReference() {
		key = types2.Z_REFVAL_P(key)
		goto try_again
	} else if key.GetType() <= types2.IS_NULL {
		if key.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		str = types2.NewString("")
		goto str_key
	} else {
		faults.Error(faults.E_WARNING, "array_key_exists(): The first argument should be either a string or an integer")
		return types2.IS_FALSE
	}
}
func ZendArrayKeyExistsSlow(subject *types2.Zval, key *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) uint32 {
	if subject.IsObject() {
		faults.Error(faults.E_DEPRECATED, "array_key_exists(): "+"Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
		var ht *types2.Array = ZendGetPropertiesFor(subject, ZEND_PROP_PURPOSE_ARRAY_CAST)
		var result uint32 = ZendArrayKeyExistsFast(ht, key, opline, executeData)
		ZendReleaseProperties(ht)
		return result
	} else {
		if key.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		if subject.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "array_key_exists() expects parameter 2 to be array, %s given", types2.ZendGetTypeByConst(subject.GetType()))
		return types2.IS_NULL
	}
}
func PromotesToArray(val *types2.Zval) types2.ZendBool {
	return val.GetType() <= types2.IS_FALSE || val.IsReference() && types2.Z_REFVAL_P(val).GetType() <= types2.IS_FALSE
}
func PromotesToObject(val *types2.Zval) types2.ZendBool {
	val = types2.ZVAL_DEREF(val)
	return val.GetType() <= types2.IS_FALSE || val.IsString() && val.String().GetLen() == 0
}
func CheckTypeArrayAssignable(type_ types2.ZendType) types2.ZendBool {
	if type_ == 0 {
		return 1
	}
	return type_.IsCode() && (type_.Code() == types2.IS_ARRAY || type_.Code() == types2.IS_ITERABLE)
}
func check_type_stdClass_assignable(type_ types2.ZendType) types2.ZendBool {
	if type_ == 0 {
		return 1
	}
	if type_.IsClass() {
		if type_.IsCe() {
			return type_.Ce() == ZendStandardClassDef
		} else {
			return ascii.StrCaseEquals(type_.Name().GetStr(), "stdclass")
		}
	} else {
		return type_.Code() == types2.IS_OBJECT
	}
}
func ZendVerifyRefArrayAssignable(ref *types2.ZendReference) types2.ZendBool {
	var prop *ZendPropertyInfo
	b.Assert(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *types2.ZendPropertyInfoSourceList = &(ref.GetSources())
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *types2.ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if types2.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = types2.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
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
func zend_verify_ref_stdClass_assignable(ref *types2.ZendReference) types2.ZendBool {
	var prop *ZendPropertyInfo
	b.Assert(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *types2.ZendPropertyInfoSourceList = &(ref.GetSources())
	var _prop **ZendPropertyInfo
	var _end ***ZendPropertyInfo
	var _list *types2.ZendPropertyInfoList
	if _source_list.GetPtr() != nil {
		if types2.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
			_list = types2.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
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

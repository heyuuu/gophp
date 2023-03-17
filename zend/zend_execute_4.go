// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendPreIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, executeData *ZendExecuteData) {
	var rv Zval
	var z *Zval
	var obj Zval
	var z_copy Zval
	obj.SetObject(object.GetObj())
	obj.AddRefcount()
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if EG__().GetException() != nil {
		OBJ_RELEASE(obj.GetObj())
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
		return
	}
	if z.IsObject() && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	ZVAL_COPY_DEREF(&z_copy, z)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &z_copy)
	}
	Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	OBJ_RELEASE(obj.GetObj())
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendAssignOpOverloadedProperty(
	object *Zval,
	property *Zval,
	cache_slot *any,
	value *Zval,
	opline *ZendOp,
	executeData *ZendExecuteData,
) {
	var z *Zval
	var rv Zval
	var obj Zval
	var res Zval
	obj.SetObject(object.GetObj())
	obj.AddRefcount()
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if EG__().GetException() != nil {
		OBJ_RELEASE(obj.GetObj())
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
		}
		return
	}
	if z.IsObject() && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	if ZendBinaryOp(&res, z, value, opline) == SUCCESS {
		Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &res, cache_slot)
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
	}
	ZvalPtrDtor(z)
	ZvalPtrDtor(&res)
	OBJ_RELEASE(obj.GetObj())
}
func ZendExtensionStatementHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetStatementHandler() != nil {
		extension.GetStatementHandler()(frame)
	}
}
func ZendExtensionFcallBeginHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallBeginHandler() != nil {
		extension.GetFcallBeginHandler()(frame)
	}
}
func ZendExtensionFcallEndHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallEndHandler() != nil {
		extension.GetFcallEndHandler()(frame)
	}
}
func ZendGetTargetSymbolTable(fetch_type int, executeData *ZendExecuteData) *HashTable {
	var ht *HashTable
	if (fetch_type & (ZEND_FETCH_GLOBAL_LOCK | ZEND_FETCH_GLOBAL)) != 0 {
		ht = EG__().GetSymbolTable()
	} else {
		ZEND_ASSERT((fetch_type & ZEND_FETCH_LOCAL) != 0)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			ZendRebuildSymbolTable()
		}
		ht = executeData.GetSymbolTable()
	}
	return ht
}
func ZendUndefinedOffset(lval ZendLong) {
	ZendError(E_NOTICE, "Undefined offset: "+ZEND_LONG_FMT, lval)
}
func ZendUndefinedIndex(offset *ZendString) {
	ZendError(E_NOTICE, "Undefined index: %s", offset.GetVal())
}
func ZendUndefinedOffsetWrite(ht *HashTable, lval ZendLong) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
	}
	ZendUndefinedOffset(lval)
	if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
		ht.DestroyEx()
		return FAILURE
	}
	if EG__().GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedIndexWrite(ht *HashTable, offset *ZendString) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
	}
	ZendUndefinedIndex(offset)
	if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
		ht.DestroyEx()
		return FAILURE
	}
	if EG__().GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedMethod(ce *ZendClassEntry, method *ZendString) {
	ZendThrowError(nil, "Call to undefined method %s::%s()", ce.GetName().GetVal(), method.GetVal())
}
func ZendInvalidMethodCall(object *Zval, function_name *Zval) {
	ZendThrowError(nil, "Call to a member function %s() on %s", Z_STRVAL_P(function_name), ZendGetTypeByConst(object.GetType()))
}
func ZendNonStaticMethodCall(fbc *ZendFunction) {
	if fbc.IsAllowStatic() {
		ZendError(E_DEPRECATED, "Non-static method %s::%s() should not be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	} else {
		ZendThrowError(ZendCeError, "Non-static method %s::%s() cannot be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	}
}
func ZendParamMustBeRef(func_ *ZendFunction, arg_num uint32) {
	ZendError(E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", arg_num, b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
}
func ZendUseScalarAsArray() {
	ZendError(E_WARNING, "Cannot use a scalar value as an array")
}
func ZendCannotAddElement() {
	ZendError(E_WARNING, "Cannot add element to the array as the next element is already occupied")
}
func ZendUseResourceAsOffset(dim *Zval) {
	ZendError(E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", Z_RES_HANDLE_P(dim), Z_RES_HANDLE_P(dim))
}
func ZendUseNewElementForString() {
	ZendThrowError(nil, "[] operator not supported for strings")
}
func ZendBinaryAssignOpDimSlow(container *Zval, dim *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	if container.IsString() {
		if opline.GetOp2Type() == IS_UNUSED {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, BP_VAR_RW, executeData)
			ZendWrongStringOffset(executeData)
		}
	} else if !(container.IsError()) {
		ZendUseScalarAsArray()
	}
}
func SlowIndexConvert(ht *HashTable, dim *Zval, value *ZendValue, executeData *ZendExecuteData) ZendUchar {
	switch dim.GetType() {
	case IS_UNDEF:

		/* The array may be destroyed while throwing the notice.
		 * Temporarily increase the refcount to detect this situation. */

		if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
			ht.AddRefcount()
		}
		ZVAL_UNDEFINED_OP2()
		if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
			ht.DestroyEx()
			return IS_NULL
		}
		if EG__().GetException() != nil {
			return IS_NULL
		}
		fallthrough
	case IS_NULL:
		value.SetStr(ZSTR_EMPTY_ALLOC())
		return IS_STRING
	case IS_DOUBLE:
		value.SetLval(ZendDvalToLval(dim.GetDval()))
		return IS_LONG
	case IS_RESOURCE:
		ZendUseResourceAsOffset(dim)
		value.SetLval(Z_RES_HANDLE_P(dim))
		return IS_LONG
	case IS_FALSE:
		value.SetLval(0)
		return IS_LONG
	case IS_TRUE:
		value.SetLval(1)
		return IS_LONG
	default:
		ZendIllegalOffset()
		return IS_NULL
	}
}
func ZendFetchDimensionAddressInner(ht *HashTable, dim *Zval, dim_type int, type_ int, executeData *ZendExecuteData) *Zval {
	var retval *Zval = nil
	var offset_key *ZendString
	var hval ZendUlong
try_again:
	if dim.IsLong() {
		hval = dim.GetLval()
	num_index:
		ZEND_HASH_INDEX_FIND(ht, hval, retval, num_undef)
		return retval
	num_undef:
		switch type_ {
		case BP_VAR_R:
			ZendUndefinedOffset(hval)
			fallthrough
		case BP_VAR_UNSET:
			fallthrough
		case BP_VAR_IS:
			retval = EG__().GetUninitializedZval()
		case BP_VAR_RW:
			if ZendUndefinedOffsetWrite(ht, hval) == FAILURE {
				return nil
			}
			fallthrough
		case BP_VAR_W:
			retval = ht.IndexAddNewH(hval, EG__().GetUninitializedZval())
		}
	} else if dim.IsString() {
		offset_key = dim.GetStr()
		{
			if ZEND_HANDLE_NUMERIC(offset_key, &hval) {
				goto num_index
			}
		}
	str_index:
		retval = ht.KeyFind(offset_key.GetStr())
		if retval != nil {

			/* support for $GLOBALS[...] */

			if retval.IsIndirect() {
				retval = retval.GetZv()
				if retval.IsUndef() {
					switch type_ {
					case BP_VAR_R:
						ZendUndefinedIndex(offset_key)
						fallthrough
					case BP_VAR_UNSET:
						fallthrough
					case BP_VAR_IS:
						retval = EG__().GetUninitializedZval()
					case BP_VAR_RW:
						if ZendUndefinedIndexWrite(ht, offset_key) != 0 {
							return nil
						}
						fallthrough
					case BP_VAR_W:
						retval.SetNull()
					}
				}
			}

			/* support for $GLOBALS[...] */

		} else {
			switch type_ {
			case BP_VAR_R:
				ZendUndefinedIndex(offset_key)
				fallthrough
			case BP_VAR_UNSET:
				fallthrough
			case BP_VAR_IS:
				retval = EG__().GetUninitializedZval()
			case BP_VAR_RW:

				/* Key may be released while throwing the undefined index warning. */

				offset_key.AddRefcount()
				if ZendUndefinedIndexWrite(ht, offset_key) == FAILURE {
					ZendStringRelease(offset_key)
					return nil
				}
				retval = ht.KeyAddNew(offset_key.GetStr(), EG__().GetUninitializedZval())
				ZendStringRelease(offset_key)
			case BP_VAR_W:
				retval = ht.KeyAddNew(offset_key.GetStr(), EG__().GetUninitializedZval())
			}
		}
	} else if dim.IsReference() {
		dim = Z_REFVAL_P(dim)
		goto try_again
	} else {
		var val ZendValue
		var t ZendUchar = SlowIndexConvert(ht, dim, &val, executeData)
		if t == IS_STRING {
			offset_key = val.GetStr()
			goto str_index
		} else if t == IS_LONG {
			hval = val.GetLval()
			goto num_index
		} else {
			if type_ == BP_VAR_W || type_ == BP_VAR_RW {
				retval = nil
			} else {
				retval = EG__().GetUninitializedZval()
			}
		}
	}
	return retval
}
func zend_fetch_dimension_address_inner_W(ht *HashTable, dim *Zval, executeData *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_inner_W_CONST(ht *HashTable, dim *Zval, executeData *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_inner_RW(ht *HashTable, dim *Zval, executeData *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_RW, executeData)
}
func zend_fetch_dimension_address_inner_RW_CONST(ht *HashTable, dim *Zval, executeData *ZendExecuteData) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_RW, executeData)
}
func ZendFetchDimensionAddress(
	result *Zval,
	container *Zval,
	dim *Zval,
	dim_type int,
	type_ int,
	executeData *ZendExecuteData,
) {
	var retval *Zval
	if container.IsArray() {
	try_array:
		SEPARATE_ARRAY(container)
	fetch_from_array:
		if dim == nil {
			retval = container.GetArr().NextIndexInsert(EG__().GetUninitializedZval())
			if retval == nil {
				ZendCannotAddElement()
				result.IsError()
				return
			}
		} else {
			retval = ZendFetchDimensionAddressInner(container.GetArr(), dim, dim_type, type_, executeData)
			if retval == nil {
				result.IsError()
				return
			}
		}
		result.SetIndirect(retval)
		return
	} else if container.IsReference() {
		var ref *ZendReference = container.GetRef()
		container = Z_REFVAL_P(container)
		if container.IsArray() {
			goto try_array
		} else if container.GetType() <= IS_FALSE {
			if type_ != BP_VAR_UNSET {
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					if ZendVerifyRefArrayAssignable(ref) == 0 {
						result.IsError()
						return
					}
				}
				ArrayInit(container)
				goto fetch_from_array
			} else {
				goto return_null
			}
		}
	}
	if container.IsString() {
		if dim == nil {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, type_, executeData)
			ZendWrongStringOffset(executeData)
		}
		result.IsError()
	} else if container.IsObject() {
		if dim != nil && dim.IsUndef() {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		if retval == EG__().GetUninitializedZval() {
			var ce *ZendClassEntry = Z_OBJCE_P(container)
			result.SetNull()
			ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
		} else if retval != nil && retval.GetType() != IS_UNDEF {
			if !(retval.IsReference()) {
				if result != retval {
					ZVAL_COPY(result, retval)
					retval = result
				}
				if retval.GetType() != IS_OBJECT {
					var ce *ZendClassEntry = Z_OBJCE_P(container)
					ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
				}
			} else if retval.GetRefcount() == 1 {
				ZVAL_UNREF(retval)
			}
			if result != retval {
				result.SetIndirect(retval)
			}
		} else {
			result.IsError()
		}
	} else {
		if container.GetType() <= IS_FALSE {
			if type_ != BP_VAR_W && container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			if type_ != BP_VAR_UNSET {
				ArrayInit(container)
				goto fetch_from_array
			} else {
			return_null:

				/* for read-mode only */

				if dim != nil && dim.IsUndef() {
					ZVAL_UNDEFINED_OP2()
				}
				result.SetNull()
			}
		} else if container.IsError() {
			result.IsError()
		} else {
			if type_ == BP_VAR_UNSET {
				ZendError(E_WARNING, "Cannot unset offset in a non-array variable")
				result.SetNull()
			} else {
				ZendUseScalarAsArray()
				result.IsError()
			}
		}
	}
}

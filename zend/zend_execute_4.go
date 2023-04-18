package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendPreIncdecOverloadedProperty(object *types2.Zval, property *types2.Zval, cache_slot *any, opline *ZendOp, executeData *ZendExecuteData) {
	var rv types2.Zval
	var z *types2.Zval
	var obj types2.Zval
	var z_copy types2.Zval
	obj.SetObject(object.Object())
	// 	obj.AddRefcount()
	z = types2.Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if EG__().GetException() != nil {
		// OBJ_RELEASE(obj.Object())
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return
	}
	if z.IsObject() && types2.Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 types2.Zval
		var value *types2.Zval = types2.Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			// ZvalPtrDtor(&rv)
		}
		z.CopyValueFrom(value)
	}
	types2.ZVAL_COPY_DEREF(&z_copy, z)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	if RETURN_VALUE_USED(opline) {
		types2.ZVAL_COPY(opline.Result(), &z_copy)
	}
	types2.Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	// OBJ_RELEASE(obj.Object())
	// ZvalPtrDtor(&z_copy)
	// ZvalPtrDtor(z)
}
func ZendAssignOpOverloadedProperty(
	object *types2.Zval,
	property *types2.Zval,
	cache_slot *any,
	value *types2.Zval,
	opline *ZendOp,
	executeData *ZendExecuteData,
) {
	var z *types2.Zval
	var rv types2.Zval
	var obj types2.Zval
	var res types2.Zval
	obj.SetObject(object.Object())
	// 	obj.AddRefcount()
	z = types2.Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if EG__().GetException() != nil {
		// OBJ_RELEASE(obj.Object())
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetUndef()
		}
		return
	}
	if z.IsObject() && types2.Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 types2.Zval
		var value *types2.Zval = types2.Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			// ZvalPtrDtor(&rv)
		}
		z.CopyValueFrom(value)
	}
	if ZendBinaryOp(&res, z, value, opline) == types2.SUCCESS {
		types2.Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &res, cache_slot)
	}
	if RETURN_VALUE_USED(opline) {
		types2.ZVAL_COPY(opline.Result(), &res)
	}
	// ZvalPtrDtor(z)
	// ZvalPtrDtor(&res)
	// OBJ_RELEASE(obj.Object())
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
func ZendGetTargetSymbolTable(fetch_type int, executeData *ZendExecuteData) *types2.Array {
	var ht *types2.Array
	if (fetch_type & (ZEND_FETCH_GLOBAL_LOCK | ZEND_FETCH_GLOBAL)) != 0 {
		ht = EG__().GetSymbolTable()
	} else {
		b.Assert((fetch_type & ZEND_FETCH_LOCAL) != 0)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			ZendRebuildSymbolTable()
		}
		ht = executeData.GetSymbolTable()
	}
	return ht
}
func ZendUndefinedOffset(lval ZendLong) {
	faults.Error(faults.E_NOTICE, "Undefined offset: "+ZEND_LONG_FMT, lval)
}
func ZendUndefinedIndex(offset *types2.String) {
	faults.Error(faults.E_NOTICE, "Undefined index: %s", offset.GetVal())
}
func ZendUndefinedOffsetWrite(ht *types2.Array, lval ZendLong) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
		// 		ht.AddRefcount()
	}
	ZendUndefinedOffset(lval)
	if (ht.GetGcFlags()&types2.IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
		ht.DestroyEx()
		return types2.FAILURE
	}
	if EG__().GetException() != nil {
		return types2.FAILURE
	}
	return types2.SUCCESS
}
func ZendUndefinedIndexWrite(ht *types2.Array, offset *types2.String) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
		// 		ht.AddRefcount()
	}
	ZendUndefinedIndex(offset)
	if (ht.GetGcFlags()&types2.IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
		ht.DestroyEx()
		return types2.FAILURE
	}
	if EG__().GetException() != nil {
		return types2.FAILURE
	}
	return types2.SUCCESS
}
func ZendUndefinedMethod(ce *types2.ClassEntry, method *types2.String) {
	faults.ThrowError(nil, "Call to undefined method %s::%s()", ce.GetName().GetVal(), method.GetVal())
}
func ZendInvalidMethodCall(object *types2.Zval, function_name *types2.Zval) {
	faults.ThrowError(nil, "Call to a member function %s() on %s", function_name.String().GetVal(), types2.ZendGetTypeByConst(object.GetType()))
}
func ZendNonStaticMethodCall(fbc types2.IFunction) {
	if fbc.IsAllowStatic() {
		faults.Error(faults.E_DEPRECATED, "Non-static method %s::%s() should not be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	} else {
		faults.ThrowError(faults.ZendCeError, "Non-static method %s::%s() cannot be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	}
}
func ZendParamMustBeRef(func_ types2.IFunction, arg_num uint32) {
	faults.Error(faults.E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", arg_num, b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
}
func ZendUseScalarAsArray() {
	faults.Error(faults.E_WARNING, "Cannot use a scalar value as an array")
}
func ZendCannotAddElement() {
	faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
}
func ZendUseResourceAsOffset(dim *types2.Zval) {
	faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", types2.Z_RES_HANDLE_P(dim), types2.Z_RES_HANDLE_P(dim))
}
func ZendUseNewElementForString() {
	faults.ThrowError(nil, "[] operator not supported for strings")
}
func ZendBinaryAssignOpDimSlow(container *types2.Zval, dim *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
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
func SlowIndexConvertEx(ht *types2.Array, dim *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	switch dim.GetType() {
	case types2.IS_UNDEF:
		/* The array may be destroyed while throwing the notice.
		 * Temporarily increase the refcount to detect this situation. */
		if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
			// 			ht.AddRefcount()
		}
		ZVAL_UNDEFINED_OP2(executeData)
		if (ht.GetGcFlags()&types2.IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
			ht.DestroyEx()
			return types2.NewZvalNull()
		}
		if EG__().GetException() != nil {
			return types2.NewZvalNull()
		}
		fallthrough
	case types2.IS_NULL:
		return types2.NewZvalString("")
	case types2.IS_DOUBLE:
		return types2.NewZvalLong(DvalToLval(dim.Double()))
	case types2.IS_RESOURCE:
		ZendUseResourceAsOffset(dim)
		return types2.NewZvalLong(types2.Z_RES_HANDLE_P(dim))
	case types2.IS_FALSE:
		return types2.NewZvalLong(0)
	case types2.IS_TRUE:
		return types2.NewZvalLong(1)
	default:
		ZendIllegalOffset()
		return types2.NewZvalNull()
	}
}
func ZendFetchDimensionAddressInner(ht *types2.Array, dim *types2.Zval, dim_type int, type_ int, executeData *ZendExecuteData) *types2.Zval {
	var retval *types2.Zval = nil
	var offset_key *types2.String
	var hval ZendUlong
try_again:
	if dim.IsLong() {
		hval = dim.Long()
	num_index:
		retval = ht.IndexFind(hval)
		if retval == nil {
			goto num_undef
		}
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
			if ZendUndefinedOffsetWrite(ht, hval) == types2.FAILURE {
				return nil
			}
			fallthrough
		case BP_VAR_W:
			retval = ht.IndexAddNew(hval, EG__().GetUninitializedZval())
		}
	} else if dim.IsString() {
		offset_key = dim.String()
		{
			if types2.HandleNumericStr(offset_key.GetStr(), &hval) {
				goto num_index
			}
		}
	str_index:
		retval = ht.KeyFind(offset_key.GetStr())
		if retval != nil {

			/* support for $GLOBALS[...] */

			if retval.IsIndirect() {
				retval = retval.Indirect()
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

				//offset_key.AddRefcount()
				if ZendUndefinedIndexWrite(ht, offset_key) == types2.FAILURE {
					// types.ZendStringRelease(offset_key)
					return nil
				}
				retval = ht.KeyAddNew(offset_key.GetStr(), EG__().GetUninitializedZval())
				// types.ZendStringRelease(offset_key)
			case BP_VAR_W:
				retval = ht.KeyAddNew(offset_key.GetStr(), EG__().GetUninitializedZval())
			}
		}
	} else if dim.IsReference() {
		dim = types2.Z_REFVAL_P(dim)
		goto try_again
	} else {
		var zv = SlowIndexConvertEx(ht, dim, executeData)
		if zv.IsString() {
			offset_key = zv.String()
			goto str_index
		} else if zv.IsLong() {
			hval = zv.Long()
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
func zend_fetch_dimension_address_inner_W(ht *types2.Array, dim *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_inner_W_CONST(ht *types2.Array, dim *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_W, executeData)
}
func zend_fetch_dimension_address_inner_RW(ht *types2.Array, dim *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_RW, executeData)
}
func zend_fetch_dimension_address_inner_RW_CONST(ht *types2.Array, dim *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_RW, executeData)
}
func ZendFetchDimensionAddress(
	result *types2.Zval,
	container *types2.Zval,
	dim *types2.Zval,
	dim_type int,
	type_ int,
	executeData *ZendExecuteData,
) {
	var retval *types2.Zval
	if container.IsArray() {
	try_array:
		types2.SeparateArray(container)
	fetch_from_array:
		if dim == nil {
			retval = container.Array().NextIndexInsert(EG__().GetUninitializedZval())
			if retval == nil {
				ZendCannotAddElement()
				result.IsError()
				return
			}
		} else {
			retval = ZendFetchDimensionAddressInner(container.Array(), dim, dim_type, type_, executeData)
			if retval == nil {
				result.IsError()
				return
			}
		}
		result.SetIndirect(retval)
		return
	} else if container.IsReference() {
		var ref *types2.ZendReference = container.Reference()
		container = types2.Z_REFVAL_P(container)
		if container.IsArray() {
			goto try_array
		} else if container.GetType() <= types2.IS_FALSE {
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
			dim = ZVAL_UNDEFINED_OP2(executeData)
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = types2.Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		if retval == EG__().GetUninitializedZval() {
			var ce *types2.ClassEntry = types2.Z_OBJCE_P(container)
			result.SetNull()
			faults.Error(faults.E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
		} else if retval != nil && retval.IsNotUndef() {
			if !(retval.IsReference()) {
				if result != retval {
					types2.ZVAL_COPY(result, retval)
					retval = result
				}
				if retval.GetType() != types2.IS_OBJECT {
					var ce *types2.ClassEntry = types2.Z_OBJCE_P(container)
					faults.Error(faults.E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
				}
			} else if retval.GetRefcount() == 1 {
				types2.ZVAL_UNREF(retval)
			}
			if result != retval {
				result.SetIndirect(retval)
			}
		} else {
			result.IsError()
		}
	} else {
		if container.GetType() <= types2.IS_FALSE {
			if type_ != BP_VAR_W && container.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
			}
			if type_ != BP_VAR_UNSET {
				ArrayInit(container)
				goto fetch_from_array
			} else {
			return_null:

				/* for read-mode only */

				if dim != nil && dim.IsUndef() {
					ZVAL_UNDEFINED_OP2(executeData)
				}
				result.SetNull()
			}
		} else if container.IsError() {
			result.IsError()
		} else {
			if type_ == BP_VAR_UNSET {
				faults.Error(faults.E_WARNING, "Cannot unset offset in a non-array variable")
				result.SetNull()
			} else {
				ZendUseScalarAsArray()
				result.IsError()
			}
		}
	}
}

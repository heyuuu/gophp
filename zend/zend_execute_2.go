// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendVerifyWeakScalarTypeHint(type_hint ZendUchar, arg *Zval) ZendBool {
	switch type_hint {
	case _IS_BOOL:
		var dest ZendBool
		if ZendParseArgBoolWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_BOOL(arg, dest != 0)
		return 1
	case IS_LONG:
		var dest ZendLong
		if ZendParseArgLongWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		arg.SetLong(dest)
		return 1
	case IS_DOUBLE:
		var dest float64
		if ZendParseArgDoubleWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		arg.SetDouble(dest)
		return 1
	case IS_STRING:
		var dest *ZendString

		/* on success "arg" is converted to IS_STRING */

		return ZendParseArgStrWeak(arg, &dest)

	/* on success "arg" is converted to IS_STRING */

	default:
		return 0
	}
}
func ZendVerifyScalarTypeHint(type_hint ZendUchar, arg *Zval, strict ZendBool) ZendBool {
	if strict != 0 {

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

		if type_hint != IS_DOUBLE || arg.GetType() != IS_LONG {
			return 0
		}

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	} else if arg.IsNull() {

		/* NULL may be accepted only by nullable hints (this is already checked) */

		return 0

		/* NULL may be accepted only by nullable hints (this is already checked) */

	}
	return ZendVerifyWeakScalarTypeHint(type_hint, arg)
}
func ZendVerifyPropertyTypeError(info *ZendPropertyInfo, property *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte

	/* we _may_ land here in case reading already errored and runtime cache thus has not been updated (i.e. it contains a valid but unrelated info) */

	if EG__().GetException() != nil {
		return
	}

	// TODO Switch to a more standard error message?

	ZendFormatType(info.GetType(), &prop_type1, &prop_type2)
	void(prop_type1)
	if info.GetType().IsClass() {
		ZendTypeError("Typed property %s::$%s must be an instance of %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	} else {
		ZendTypeError("Typed property %s::$%s must be %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	}
}
func ZendResolveClassType(type_ *ZendType, self_ce *ZendClassEntry) ZendBool {
	var ce *ZendClassEntry
	var name *ZendString = type_.Name()
	if ZendStringEqualsLiteralCi(name, "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if self_ce.IsTrait() {
			ZendThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", b.Cond(type_.AllowNull(), " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if ZendStringEqualsLiteralCi(name, "parent") {
		if !(self_ce.GetParent()) {
			ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return 0
		}
		ce = self_ce.GetParent()
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			return 0
		}
	}
	ZendStringRelease(name)
	*type_ = ZEND_TYPE_ENCODE_CE(ce, type_.AllowNull())
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	ZEND_ASSERT(!(property.IsReference()))
	if info.GetType().IsClass() {
		if property.GetType() != IS_OBJECT {
			return property.IsNull() && info.GetType().AllowNull()
		}
		if !(info.GetType().IsCe()) && ZendResolveClassType(info.GetType(), info.GetCe()) == 0 {
			return 0
		}
		return InstanceofFunction(Z_OBJCE_P(property), info.GetType().Ce())
	}
	ZEND_ASSERT(info.GetType().Code() != IS_CALLABLE)
	if info.GetType().Code() == property.GetType() {
		return 1
	} else if property.IsNull() {
		return info.GetType().AllowNull()
	} else if info.GetType().Code() == _IS_BOOL && property.IsFalse() || property.IsTrue() {
		return 1
	} else if info.GetType().Code() == IS_ITERABLE {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType().Code(), property, strict)
	}
}
func IZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	if IZendCheckPropertyType(info, property, strict) != 0 {
		return 1
	}
	ZendVerifyPropertyTypeError(info, property)
	return 0
}
func ZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *Zval, value *Zval, executeData *ZendExecuteData) *Zval {
	var tmp Zval
	value = ZVAL_DEREF(value)
	ZVAL_COPY(&tmp, value)
	if IZendVerifyPropertyType(info, &tmp, executeData.IsCallUseStrictTypes()) == 0 {
		ZvalPtrDtor(&tmp)
		return EG__().GetUninitializedZval()
	}
	return ZendAssignToVariable(property_val, &tmp, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
}
func ZendCheckType(
	type_ ZendType,
	arg *Zval,
	ce **ZendClassEntry,
	cache_slot *any,
	default_value *Zval,
	scope *ZendClassEntry,
	is_return_type ZendBool,
) ZendBool {
	var ref *ZendReference = nil
	if !(type_.IsSet()) {
		return 1
	}
	if arg.IsReference() {
		ref = arg.GetRef()
		arg = Z_REFVAL_P(arg)
	}
	if type_.IsClass() {
		if *cache_slot {
			*ce = (*ZendClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass(type_.Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if (*ce) == nil {
				return arg.IsNull() && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if arg.IsObject() {
			return InstanceofFunction(Z_OBJCE_P(arg), *ce)
		}
		return arg.IsNull() && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
	} else if type_.Code() == arg.GetType() {
		return 1
	}
	if arg.IsNull() && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0) {

		/* Null passed to nullable type */

		return 1

		/* Null passed to nullable type */

	}
	if type_.Code() == IS_CALLABLE {
		return ZendIsCallable(arg, IS_CALLABLE_CHECK_SILENT, nil)
	} else if type_.Code() == IS_ITERABLE {
		return ZendIsIterable(arg)
	} else if type_.Code() == _IS_BOOL && arg.IsFalse() || arg.IsTrue() {
		return 1
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(type_.Code(), arg, b.CondF(is_return_type != 0, func() bool { return CurrEX().IsCallUseStrictTypes() }, func() bool { return CurrEX().IsArgUseStrictTypes() }))
	}
}
func ZendVerifyArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	if arg_num <= zf.GetNumArgs() {
		cur_arg_info = zf.GetArgInfo()[arg_num-1]
	} else if zf.IsVariadic() {
		cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	} else {
		return 1
	}
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyRecvArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = zf.GetArgInfo()[arg_num-1]
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num <= zf.GetNumArgs())
	cur_arg_info = zf.GetArgInfo()[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num > zf.GetNumArgs())
	ZEND_ASSERT(zf.IsVariadic())
	cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyInternalArgTypes(fbc *ZendFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	var p *Zval = call.Arg(1)
	var dummy_cache_slot any
	for i = 0; i < num_args; i++ {
		dummy_cache_slot = nil
		if ZendVerifyArgType(fbc, i+1, p, nil, &dummy_cache_slot) == 0 {
			EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
			return 0
		}
		p++
	}
	return 1
}
func ZendMissingArgError(executeData *ZendExecuteData) {
	var ptr *ZendExecuteData = EX(prev_execute_data)
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return EX(func_).common.scope.name.GetVal() }, ""), b.Cond(EX(func_).common.scope, "::", ""), EX(func_).common.function_name.GetVal(), EX_NUM_ARGS(), ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	} else {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return EX(func_).common.scope.name.GetVal() }, ""), b.Cond(EX(func_).common.scope, "::", ""), EX(func_).common.function_name.GetVal(), EX_NUM_ARGS(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	}
}
func ZendVerifyReturnError(zf *ZendFunction, ce *ZendClassEntry, value *Zval) {
	var arg_info *ZendArgInfo = zf.GetArgInfo()[-1]
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
	ZendTypeError("Return value of %s%s%s() must %s%s%s, %s%s returned", fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
}
func ZendVerifyReturnType(zf *ZendFunction, ret *Zval, cache_slot *any) {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	var ce *ZendClassEntry = nil
	if ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0 {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf *ZendFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType().IsSet() && ret_info.GetType().Code() != IS_VOID {
		var ce *ZendClassEntry = nil
		if ret_info.GetType().IsClass() {
			if *cache_slot {
				ce = (*ZendClassEntry)(*cache_slot)
			} else {
				ce = ZendFetchClass(ret_info.GetType().Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					*cache_slot = any(ce)
				}
			}
		}
		ZendVerifyReturnError(zf, ce, nil)
		return 0
	}
	return 1
}
func ZendUseObjectAsArray() {
	ZendThrowError(nil, "Cannot use object as array")
}
func ZendIllegalOffset() {
	ZendError(E_WARNING, "Illegal offset type")
}
func ZendAssignToObjectDim(object *Zval, dim *Zval, value *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	Z_OBJ_HT_P(object).GetWriteDimension()(object, dim, value)
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
}
func ZendBinaryOp(ret *Zval, op1 *Zval, op2 *Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *Zval, property *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var free_op_data1 ZendFreeOp
	var value *Zval
	var z *Zval
	var rv Zval
	var res Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if b.Assign(&z, Z_OBJ_HT_P(object).GetReadDimension()(object, property, BP_VAR_R, &rv)) != nil {
		if z.IsObject() && Z_OBJ_HT_P(z).GetGet() != nil {
			var rv2 Zval
			var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
			if z == &rv {
				ZvalPtrDtor(&rv)
			}
			ZVAL_COPY_VALUE(z, value)
		}
		if ZendBinaryOp(&res, z, value, opline) == SUCCESS {
			Z_OBJ_HT_P(object).GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		if RETURN_VALUE_USED(opline) {
			ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
		}
		ZvalPtrDtor(&res)
	} else {
		ZendUseObjectAsArray()
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
	}
	FREE_OP(free_op_data1)
}
func ZendBinaryAssignOpTypedRef(ref *ZendReference, value *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && ref.GetVal().IsString() {
		ConcatFunction(ref.GetVal(), ref.GetVal(), value)
		ZEND_ASSERT(ref.GetVal().IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, ref.GetVal(), value, opline)
	if ZendVerifyRefAssignableZval(ref, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		ZvalPtrDtor(ref.GetVal())
		ZVAL_COPY_VALUE(ref.GetVal(), &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *Zval, value *Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && zptr.IsString() {
		ConcatFunction(zptr, zptr, value)
		ZEND_ASSERT(zptr.IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, opline)
	if ZendVerifyPropertyType(prop_info, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		ZvalPtrDtor(zptr)
		ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *Zval, type_ int, executeData *ZendExecuteData) ZendLong {
	var offset ZendLong
try_again:
	if dim.GetType() != IS_LONG {
		switch dim.GetType() {
		case IS_STRING:
			if IS_LONG == IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				ZendError(E_WARNING, "Illegal string offset '%s'", Z_STRVAL_P(dim))
			}
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
			ZendError(E_NOTICE, "String offset cast occurred")
		case IS_REFERENCE:
			dim = Z_REFVAL_P(dim)
			goto try_again
		default:
			ZendIllegalOffset()
		}
		offset = ZvalGetLongFunc(dim)
	} else {
		offset = dim.GetLval()
	}
	return offset
}

// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendVerifyWeakScalarTypeHint(type_hint types.ZendUchar, arg *types.Zval) types.ZendBool {
	switch type_hint {
	case types.IS_BOOL:
		var dest types.ZendBool
		if argparse.ZendParseArgBoolWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		types.ZVAL_BOOL(arg, dest != 0)
		return 1
	case types.IS_LONG:
		var dest ZendLong
		if argparse.ZendParseArgLongWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		arg.SetLong(dest)
		return 1
	case types.IS_DOUBLE:
		var dest float64
		if argparse.ZendParseArgDoubleWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		arg.SetDouble(dest)
		return 1
	case types.IS_STRING:
		var dest *types.String

		/* on success "arg" is converted to IS_STRING */

		return argparse.ZendParseArgStrWeak(arg, &dest)

	/* on success "arg" is converted to IS_STRING */

	default:
		return 0
	}
}
func ZendVerifyScalarTypeHint(type_hint types.ZendUchar, arg *types.Zval, strict types.ZendBool) types.ZendBool {
	if strict != 0 {

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

		if type_hint != types.IS_DOUBLE || arg.GetType() != types.IS_LONG {
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
func ZendVerifyPropertyTypeError(info *ZendPropertyInfo, property *types.Zval) {
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
		faults.TypeError("Typed property %s::$%s must be an instance of %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return types.Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return types.ZendGetTypeByConst(property.GetType()) }))
	} else {
		faults.TypeError("Typed property %s::$%s must be %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return types.Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return types.ZendGetTypeByConst(property.GetType()) }))
	}
}
func ZendResolveClassType(type_ *types.ZendType, self_ce *types.ClassEntry) types.ZendBool {
	var ce *types.ClassEntry
	var name *types.String = type_.Name()
	if types.ZendStringEqualsLiteralCi(name, "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if self_ce.IsTrait() {
			faults.ThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", b.Cond(type_.AllowNull(), " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if types.ZendStringEqualsLiteralCi(name, "parent") {
		if !(self_ce.GetParent()) {
			faults.ThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return 0
		}
		ce = self_ce.GetParent()
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			return 0
		}
	}
	types.ZendStringRelease(name)
	*type_ = types.ZEND_TYPE_ENCODE_CE(ce, type_.AllowNull())
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *types.Zval, strict types.ZendBool) types.ZendBool {
	b.Assert(!(property.IsReference()))
	if info.GetType().IsClass() {
		if property.GetType() != types.IS_OBJECT {
			return property.IsNull() && info.GetType().AllowNull()
		}
		if !(info.GetType().IsCe()) && ZendResolveClassType(info.GetType(), info.GetCe()) == 0 {
			return 0
		}
		return InstanceofFunction(types.Z_OBJCE_P(property), info.GetType().Ce())
	}
	b.Assert(info.GetType().Code() != types.IS_CALLABLE)
	if info.GetType().Code() == property.GetType() {
		return 1
	} else if property.IsNull() {
		return info.GetType().AllowNull()
	} else if info.GetType().Code() == types.IS_BOOL && property.IsFalse() || property.IsTrue() {
		return 1
	} else if info.GetType().Code() == types.IS_ITERABLE {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType().Code(), property, strict)
	}
}
func IZendVerifyPropertyType(info *ZendPropertyInfo, property *types.Zval, strict types.ZendBool) types.ZendBool {
	if IZendCheckPropertyType(info, property, strict) != 0 {
		return 1
	}
	ZendVerifyPropertyTypeError(info, property)
	return 0
}
func ZendVerifyPropertyType(info *ZendPropertyInfo, property *types.Zval, strict types.ZendBool) types.ZendBool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *types.Zval, value *types.Zval, executeData *ZendExecuteData) *types.Zval {
	var tmp types.Zval
	value = types.ZVAL_DEREF(value)
	types.ZVAL_COPY(&tmp, value)
	if IZendVerifyPropertyType(info, &tmp, executeData.IsCallUseStrictTypes()) == 0 {
		ZvalPtrDtor(&tmp)
		return EG__().GetUninitializedZval()
	}
	return ZendAssignToVariable(property_val, &tmp, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
}
func ZendCheckType(
	type_ types.ZendType,
	arg *types.Zval,
	ce **types.ClassEntry,
	cache_slot *any,
	default_value *types.Zval,
	scope *types.ClassEntry,
	is_return_type types.ZendBool,
) types.ZendBool {
	var ref *types.ZendReference = nil
	if !(type_.IsSet()) {
		return 1
	}
	if arg.IsReference() {
		ref = arg.GetRef()
		arg = types.Z_REFVAL_P(arg)
	}
	if type_.IsClass() {
		if *cache_slot {
			*ce = (*types.ClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass(type_.Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if (*ce) == nil {
				return arg.IsNull() && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if arg.IsObject() {
			return InstanceofFunction(types.Z_OBJCE_P(arg), *ce)
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
	if type_.Code() == types.IS_CALLABLE {
		return ZendIsCallable(arg, IS_CALLABLE_CHECK_SILENT, nil)
	} else if type_.Code() == types.IS_ITERABLE {
		return ZendIsIterable(arg)
	} else if type_.Code() == types.IS_BOOL && arg.IsFalse() || arg.IsTrue() {
		return 1
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(type_.Code(), arg, b.CondF(is_return_type != 0, func() bool { return CurrEX().IsCallUseStrictTypes() }, func() bool { return CurrEX().IsArgUseStrictTypes() }))
	}
}
func ZendVerifyArgType(zf *ZendFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *types.ClassEntry
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
func ZendVerifyRecvArgType(zf *ZendFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = zf.GetArgInfo()[arg_num-1]
	var ce *types.ClassEntry
	b.Assert(arg_num <= zf.GetNumArgs())
	cur_arg_info = zf.GetArgInfo()[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf *ZendFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *types.ClassEntry
	b.Assert(arg_num > zf.GetNumArgs())
	b.Assert(zf.IsVariadic())
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
	var num_args uint32 = call.NumArgs()
	var p *types.Zval = call.Arg(1)
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
	var ptr *ZendExecuteData = executeData.GetPrevExecuteData()
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
		faults.ThrowError(faults.ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", b.CondF1(executeData.GetFunc().common.scope, func() []byte { return executeData.GetFunc().common.scope.name.GetVal() }, ""), b.Cond(executeData.GetFunc().common.scope, "::", ""), executeData.GetFunc().common.function_name.GetVal(), executeData.NumArgs(), ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno(), b.Cond(executeData.GetFunc().common.required_num_args == executeData.GetFunc().common.num_args, "exactly", "at least"), executeData.GetFunc().common.required_num_args)
	} else {
		faults.ThrowError(faults.ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", b.CondF1(executeData.GetFunc().common.scope, func() []byte { return executeData.GetFunc().common.scope.name.GetVal() }, ""), b.Cond(executeData.GetFunc().common.scope, "::", ""), executeData.GetFunc().common.function_name.GetVal(), executeData.NumArgs(), b.Cond(executeData.GetFunc().common.required_num_args == executeData.GetFunc().common.num_args, "exactly", "at least"), executeData.GetFunc().common.required_num_args)
	}
}
func ZendVerifyReturnError(zf *ZendFunction, ce *types.ClassEntry, value *types.Zval) {
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
	faults.TypeError("Return value of %s%s%s() must %s%s%s, %s%s returned", fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
}
func ZendVerifyReturnType(zf *ZendFunction, ret *types.Zval, cache_slot *any) {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	var ce *types.ClassEntry = nil
	if ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0 {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf *ZendFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType().IsSet() && ret_info.GetType().Code() != types.IS_VOID {
		var ce *types.ClassEntry = nil
		if ret_info.GetType().IsClass() {
			if *cache_slot {
				ce = (*types.ClassEntry)(*cache_slot)
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
	faults.ThrowError(nil, "Cannot use object as array")
}
func ZendIllegalOffset() {
	faults.Error(faults.E_WARNING, "Illegal offset type")
}
func ZendAssignToObjectDim(object *types.Zval, dim *types.Zval, value *types.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	types.Z_OBJ_HT_P(object).GetWriteDimension()(object, dim, value)
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
}
func ZendBinaryOp(ret *types.Zval, op1 *types.Zval, op2 *types.Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *types.Zval, property *types.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var free_op_data1 ZendFreeOp
	var value *types.Zval
	var z *types.Zval
	var rv types.Zval
	var res types.Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if b.Assign(&z, types.Z_OBJ_HT_P(object).GetReadDimension()(object, property, BP_VAR_R, &rv)) != nil {
		if z.IsObject() && types.Z_OBJ_HT_P(z).GetGet() != nil {
			var rv2 types.Zval
			var value *types.Zval = types.Z_OBJ_HT_P(z).GetGet()(z, &rv2)
			if z == &rv {
				ZvalPtrDtor(&rv)
			}
			types.ZVAL_COPY_VALUE(z, value)
		}
		if ZendBinaryOp(&res, z, value, opline) == types.SUCCESS {
			types.Z_OBJ_HT_P(object).GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
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
func ZendBinaryAssignOpTypedRef(ref *types.ZendReference, value *types.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy types.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && ref.GetVal().IsString() {
		ConcatFunction(ref.GetVal(), ref.GetVal(), value)
		b.Assert(ref.GetVal().IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, ref.GetVal(), value, opline)
	if ZendVerifyRefAssignableZval(ref, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		ZvalPtrDtor(ref.GetVal())
		types.ZVAL_COPY_VALUE(ref.GetVal(), &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *types.Zval, value *types.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy types.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && zptr.IsString() {
		ConcatFunction(zptr, zptr, value)
		b.Assert(zptr.IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, opline)
	if ZendVerifyPropertyType(prop_info, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		ZvalPtrDtor(zptr)
		types.ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *types.Zval, type_ int, executeData *ZendExecuteData) ZendLong {
	var offset ZendLong
try_again:
	if dim.GetType() != types.IS_LONG {
		switch dim.GetType() {
		case types.IS_STRING:
			if types.IS_LONG == IsNumericString(dim.GetStr().GetStr(), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				faults.Error(faults.E_WARNING, "Illegal string offset '%s'", dim.GetStr().GetVal())
			}
		case types.IS_UNDEF:
			ZVAL_UNDEFINED_OP2()
			fallthrough
		case types.IS_DOUBLE:
			fallthrough
		case types.IS_NULL:
			fallthrough
		case types.IS_FALSE:
			fallthrough
		case types.IS_TRUE:
			faults.Error(faults.E_NOTICE, "String offset cast occurred")
		case types.IS_REFERENCE:
			dim = types.Z_REFVAL_P(dim)
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

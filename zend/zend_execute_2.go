package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZendVerifyWeakScalarTypeHint(type_hint types2.ZendUchar, arg *types2.Zval) types2.ZendBool {
	switch type_hint {
	case types2.IS_BOOL:
		if val, ok := zpp.ParseBoolWeak(arg); ok {
			// ZvalPtrDtor(arg)
			arg.SetBool(val)
			return 1
		}
	case types2.IS_LONG:
		if val, ok := zpp.ParseLongWeak(arg, false); ok {
			// ZvalPtrDtor(arg)
			arg.SetLong(val)
			return 1
		}
	case types2.IS_DOUBLE:
		if val, ok := zpp.ParseDoubleWeak(arg); ok {
			// ZvalPtrDtor(arg)
			arg.SetDouble(val)
			return 1
		}
	case types2.IS_STRING:
		if val, ok := zpp.ParseZStrWeak(arg); ok {
			arg.SetString(val)
			return 1
		}
	}
	return 0
}
func ZendVerifyScalarTypeHint(type_hint types2.ZendUchar, arg *types2.Zval, strict types2.ZendBool) types2.ZendBool {
	if strict != 0 {
		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */
		if type_hint != types2.IS_DOUBLE || arg.GetType() != types2.IS_LONG {
			return 0
		}
	} else if arg.IsNull() {
		/* NULL may be accepted only by nullable hints (this is already checked) */
		return 0
	}
	return ZendVerifyWeakScalarTypeHint(type_hint, arg)
}
func ZendVerifyPropertyTypeError(info *ZendPropertyInfo, property *types2.Zval) {
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
		faults.TypeError("Typed property %s::$%s must be an instance of %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return types2.Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return types2.ZendGetTypeByConst(property.GetType()) }))
	} else {
		faults.TypeError("Typed property %s::$%s must be %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsObject(), func() []byte { return types2.Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return types2.ZendGetTypeByConst(property.GetType()) }))
	}
}
func ZendResolveClassType(type_ *types2.ZendType, self_ce *types2.ClassEntry) types2.ZendBool {
	var ce *types2.ClassEntry
	var name *types2.String = type_.Name()
	if ascii.StrCaseEquals(name.GetStr(), "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if self_ce.IsTrait() {
			faults.ThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", b.Cond(type_.AllowNull(), " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if ascii.StrCaseEquals(name.GetStr(), "parent") {
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
	// types.ZendStringRelease(name)
	*type_ = types2.ZEND_TYPE_ENCODE_CE(ce, type_.AllowNull())
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *types2.Zval, strict types2.ZendBool) types2.ZendBool {
	b.Assert(!(property.IsReference()))
	if info.GetType().IsClass() {
		if property.GetType() != types2.IS_OBJECT {
			return property.IsNull() && info.GetType().AllowNull()
		}
		if !(info.GetType().IsCe()) && ZendResolveClassType(info.GetType(), info.GetCe()) == 0 {
			return 0
		}
		return InstanceofFunction(types2.Z_OBJCE_P(property), info.GetType().Ce())
	}
	b.Assert(info.GetType().Code() != types2.IS_CALLABLE)
	if info.GetType().Code() == property.GetType() {
		return 1
	} else if property.IsNull() {
		return info.GetType().AllowNull()
	} else if info.GetType().Code() == types2.IS_BOOL && property.IsFalse() || property.IsTrue() {
		return 1
	} else if info.GetType().Code() == types2.IS_ITERABLE {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType().Code(), property, strict)
	}
}
func IZendVerifyPropertyType(info *ZendPropertyInfo, property *types2.Zval, strict types2.ZendBool) types2.ZendBool {
	if IZendCheckPropertyType(info, property, strict) != 0 {
		return 1
	}
	ZendVerifyPropertyTypeError(info, property)
	return 0
}
func ZendVerifyPropertyType(info *ZendPropertyInfo, property *types2.Zval, strict types2.ZendBool) types2.ZendBool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *types2.Zval, value *types2.Zval, executeData *ZendExecuteData) *types2.Zval {
	var tmp types2.Zval
	value = types2.ZVAL_DEREF(value)
	types2.ZVAL_COPY(&tmp, value)
	if IZendVerifyPropertyType(info, &tmp, executeData.IsCallUseStrictTypes()) == 0 {
		// ZvalPtrDtor(&tmp)
		return EG__().GetUninitializedZval()
	}
	return ZendAssignToVariable(property_val, &tmp, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
}
func ZendCheckType(
	type_ types2.ZendType,
	arg *types2.Zval,
	ce **types2.ClassEntry,
	cache_slot *any,
	default_value *types2.Zval,
	scope *types2.ClassEntry,
	is_return_type types2.ZendBool,
) types2.ZendBool {
	var ref *types2.ZendReference = nil
	if !(type_.IsSet()) {
		return 1
	}
	if arg.IsReference() {
		ref = arg.Reference()
		arg = types2.Z_REFVAL_P(arg)
	}
	if type_.IsClass() {
		if *cache_slot {
			*ce = (*types2.ClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass(type_.Name().GetStr(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if (*ce) == nil {
				return arg.IsNull() && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if arg.IsObject() {
			return InstanceofFunction(types2.Z_OBJCE_P(arg), *ce)
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
	if type_.Code() == types2.IS_CALLABLE {
		return ZendIsCallable(arg, IS_CALLABLE_CHECK_SILENT, nil)
	} else if type_.Code() == types2.IS_ITERABLE {
		return ZendIsIterable(arg)
	} else if type_.Code() == types2.IS_BOOL && arg.IsFalse() || arg.IsTrue() {
		return 1
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(type_.Code(), arg, b.CondF(is_return_type != 0, func() bool { return CurrEX().IsCallUseStrictTypes() }, func() bool { return CurrEX().IsArgUseStrictTypes() }))
	}
}
func ZendVerifyArgType(zf types2.IFunction, arg_num uint32, arg *types2.Zval, default_value *types2.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *types2.ClassEntry
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
func ZendVerifyRecvArgType(zf types2.IFunction, arg_num uint32, arg *types2.Zval, default_value *types2.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = zf.GetArgInfo()[arg_num-1]
	var ce *types2.ClassEntry
	b.Assert(arg_num <= zf.GetNumArgs())
	cur_arg_info = zf.GetArgInfo()[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf types2.IFunction, arg_num uint32, arg *types2.Zval, default_value *types2.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *types2.ClassEntry
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
func ZendVerifyInternalArgTypes(fbc types2.IFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = call.NumArgs()
	var p *types2.Zval = call.Arg(1)
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
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetType()) {
		faults.ThrowError(faults.ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", b.CondF1(executeData.GetFunc().common.scope, func() []byte { return executeData.GetFunc().common.scope.name.GetVal() }, ""), b.Cond(executeData.GetFunc().common.scope, "::", ""), executeData.GetFunc().common.function_name.GetVal(), executeData.NumArgs(), ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno(), b.Cond(executeData.GetFunc().common.required_num_args == executeData.GetFunc().common.num_args, "exactly", "at least"), executeData.GetFunc().common.required_num_args)
	} else {
		faults.ThrowError(faults.ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", b.CondF1(executeData.GetFunc().common.scope, func() []byte { return executeData.GetFunc().common.scope.name.GetVal() }, ""), b.Cond(executeData.GetFunc().common.scope, "::", ""), executeData.GetFunc().common.function_name.GetVal(), executeData.NumArgs(), b.Cond(executeData.GetFunc().common.required_num_args == executeData.GetFunc().common.num_args, "exactly", "at least"), executeData.GetFunc().common.required_num_args)
	}
}
func ZendVerifyReturnError(zf types2.IFunction, ce *types2.ClassEntry, value *types2.Zval) {
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
func ZendVerifyReturnType(zf types2.IFunction, ret *types2.Zval, cache_slot *any) {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	var ce *types2.ClassEntry = nil
	if ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0 {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf types2.IFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType().IsSet() && ret_info.GetType().Code() != types2.IS_VOID {
		var ce *types2.ClassEntry = nil
		if ret_info.GetType().IsClass() {
			if *cache_slot {
				ce = (*types2.ClassEntry)(*cache_slot)
			} else {
				ce = ZendFetchClass(ret_info.GetType().Name().GetStr(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
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
func ZendAssignToObjectDim(object *types2.Zval, dim *types2.Zval, value *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	types2.Z_OBJ_HT_P(object).GetWriteDimension()(object, dim, value)
	if RETURN_VALUE_USED(opline) {
		types2.ZVAL_COPY(opline.Result(), value)
	}
}
func ZendBinaryOp(ret *types2.Zval, op1 *types2.Zval, op2 *types2.Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *types2.Zval, property *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var free_op_data1 ZendFreeOp
	var value *types2.Zval
	var z *types2.Zval
	var rv types2.Zval
	var res types2.Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if b.Assign(&z, types2.Z_OBJ_HT_P(object).GetReadDimension()(object, property, BP_VAR_R, &rv)) != nil {
		if z.IsObject() && types2.Z_OBJ_HT_P(z).GetGet() != nil {
			var rv2 types2.Zval
			var value *types2.Zval = types2.Z_OBJ_HT_P(z).GetGet()(z, &rv2)
			if z == &rv {
				// ZvalPtrDtor(&rv)
			}
			z.CopyValueFrom(value)
		}
		if ZendBinaryOp(&res, z, value, opline) == types2.SUCCESS {
			types2.Z_OBJ_HT_P(object).GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			// ZvalPtrDtor(&rv)
		}
		if RETURN_VALUE_USED(opline) {
			types2.ZVAL_COPY(opline.Result(), &res)
		}
		// ZvalPtrDtor(&res)
	} else {
		ZendUseObjectAsArray()
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
	}
	// 	FREE_OP(free_op_data1)
}
func ZendBinaryAssignOpTypedRef(ref *types2.ZendReference, value *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy types2.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && ref.GetVal().IsString() {
		ConcatFunction(ref.GetVal(), ref.GetVal(), value)
		b.Assert(ref.GetVal().IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, ref.GetVal(), value, opline)
	if ZendVerifyRefAssignableZval(ref, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		// ZvalPtrDtor(ref.GetVal())
		types2.ZVAL_COPY_VALUE(ref.GetVal(), &z_copy)
	} else {
		// ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *types2.Zval, value *types2.Zval, opline *ZendOp, executeData *ZendExecuteData) {
	var z_copy types2.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && zptr.IsString() {
		ConcatFunction(zptr, zptr, value)
		b.Assert(zptr.IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, opline)
	if ZendVerifyPropertyType(prop_info, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		// ZvalPtrDtor(zptr)
		types2.ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		// ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *types2.Zval, type_ int, executeData *ZendExecuteData) ZendLong {
	var offset ZendLong
try_again:
	if dim.GetType() != types2.IS_LONG {
		switch dim.GetType() {
		case types2.IS_STRING:
			if types2.IS_LONG == IsNumericString(dim.String().GetStr(), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				faults.Error(faults.E_WARNING, "Illegal string offset '%s'", dim.String().GetVal())
			}
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
			faults.Error(faults.E_NOTICE, "String offset cast occurred")
		case types2.IS_REFERENCE:
			dim = types2.Z_REFVAL_P(dim)
			goto try_again
		default:
			ZendIllegalOffset()
		}
		offset = ZvalGetLongFunc(dim)
	} else {
		offset = dim.Long()
	}
	return offset
}

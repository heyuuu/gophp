package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZendVerifyWeakScalarTypeHint(type_hint uint8, arg *types.Zval) bool {
	switch type_hint {
	case types.IsBool:
		if val, ok := zpp.ParseBoolWeak(arg); ok {
			arg.SetBool(val)
			return true
		}
	case types.IsLong:
		if val, ok := zpp.ParseLongWeak(arg, false); ok {
			arg.SetLong(val)
			return true
		}
	case types.IsDouble:
		if val, ok := zpp.ParseDoubleWeak(arg); ok {
			arg.SetDouble(val)
			return true
		}
	case types.IsString:
		if val, ok := zpp.ParseZStrWeak(arg); ok {
			arg.SetStringEx(val)
			return true
		}
	}
	return false
}
func ZendVerifyScalarTypeHint(type_hint uint8, arg *types.Zval, strict bool) bool {
	if strict {
		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */
		if type_hint != types.IsDouble || !arg.IsLong() {
			return false
		}
	} else if arg.IsNull() {
		/* NULL may be accepted only by nullable hints (this is already checked) */
		return false
	}
	return ZendVerifyWeakScalarTypeHint(type_hint, arg)
}
func ZendVerifyPropertyTypeError(info *types.PropertyInfo, property *types.Zval) {
	/* we _may_ land here in case reading already errored and runtime cache thus has not been updated (i.e. it contains a valid but unrelated info) */
	if EG__().HasException() {
		return
	}

	ceName := info.GetCe().Name()
	propName := ZendGetUnmangledPropertyNameEx(info.GetName())
	_, propType2 := ZendFormatTypeEx(*info.GetType())
	orNullTips := ""
	if info.GetType().AllowNull() {
		orNullTips = " or null"
	}

	var propTypeName string
	if property.IsObject() {
		propTypeName = property.Object().GetCe().Name()
	} else {
		propTypeName = types.ZendGetTypeByConst(property.Type())
	}

	if info.GetType().IsClass() {
		faults.TypeError(fmt.Sprintf("Typed property %s::$%s must be an instance of %s%s, %s used", ceName, propName, propType2, orNullTips, propTypeName))
	} else {
		faults.TypeError(fmt.Sprintf("Typed property %s::$%s must be %s%s, %s used", ceName, propName, propType2, orNullTips, propTypeName))
	}
}
func ZendResolveClassType(type_ *types.TypeHint, selfCe *types.ClassEntry) bool {
	var ce *types.ClassEntry
	var name = type_.Name()
	if ascii.StrCaseEquals(name, "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if selfCe.IsTrait() {
			faults.ThrowError(nil, fmt.Sprintf("Cannot write a%s value to a 'self' typed static property of a trait", lang.Cond(type_.AllowNull(), " non-null", "")))
			return false
		}
		ce = selfCe
	} else if ascii.StrCaseEquals(name, "parent") {
		if selfCe.GetParent() == nil {
			faults.ThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return false
		}
		ce = selfCe.GetParent()
	} else {
		ce = ZendLookupClassEx(name, "", ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			return false
		}
	}
	*type_ = types.TypeHintCe(ce, type_.AllowNull())
	return true
}
func IZendCheckPropertyType(info *types.PropertyInfo, property *types.Zval, strict bool) bool {
	b.Assert(!(property.IsRef()))
	if info.GetType().IsClass() {
		if !property.IsObject() {
			return property.IsNull() && info.GetType().AllowNull()
		}
		if !(info.GetType().IsCe()) && !ZendResolveClassType(info.GetType(), info.GetCe()) {
			return false
		}
		return operators.InstanceofFunction(types.Z_OBJCE_P(property), info.GetType().Ce())
	}
	b.Assert(info.GetType().Code() != types.IsCallable)
	if info.GetType().Code() == property.Type() {
		return true
	} else if property.IsNull() {
		return info.GetType().AllowNull()
	} else if info.GetType().Code() == types.IsBool && property.IsFalse() || property.IsTrue() {
		return true
	} else if info.GetType().Code() == types.IsIterable {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType().Code(), property, strict)
	}
}
func IZendVerifyPropertyType(info *types.PropertyInfo, property *types.Zval, strict bool) bool {
	if IZendCheckPropertyType(info, property, strict) {
		return true
	}
	ZendVerifyPropertyTypeError(info, property)
	return false
}
func ZendVerifyPropertyType(info *types.PropertyInfo, property *types.Zval, strict bool) bool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *types.PropertyInfo, property_val *types.Zval, value *types.Zval, executeData *ZendExecuteData) *types.Zval {
	var tmp types.Zval
	value = types.ZVAL_DEREF(value)
	types.ZVAL_COPY(&tmp, value)
	if !IZendVerifyPropertyType(info, &tmp, executeData.IsCallUseStrictTypes()) {
		return UninitializedZval()
	}
	return ZendAssignToVariable(property_val, &tmp, executeData.IsCallUseStrictTypes())
}
func ZendCheckType(typ types.TypeHint, arg *types.Zval, ce **types.ClassEntry, cacheSlot *any, defaultValue *types.Zval, scope *types.ClassEntry) bool {
	var ref *types.Reference = nil
	if !(typ.IsSet()) {
		return true
	}
	if arg.IsRef() {
		ref = arg.Ref()
		arg = types.Z_REFVAL_P(arg)
	}
	if typ.IsClass() {
		if *cacheSlot != nil {
			*ce = (*types.ClassEntry)(*cacheSlot)
		} else {
			*ce = ZendFetchClass(typ.Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if (*ce) == nil {
				return arg.IsNull() && (typ.AllowNull() || defaultValue != nil && IsNullConstant(scope, defaultValue) != 0)
			}
			*cacheSlot = any(*ce)
		}
		if arg.IsObject() {
			return operators.InstanceofFunction(types.Z_OBJCE_P(arg), *ce)
		}
		return arg.IsNull() && (typ.AllowNull() || defaultValue != nil && IsNullConstant(scope, defaultValue) != 0)
	} else if typ.Code() == arg.Type() {
		return true
	}
	if arg.IsNull() && (typ.AllowNull() || defaultValue != nil && IsNullConstant(scope, defaultValue) != 0) {
		/* Null passed to nullable type */
		return true
	}
	if typ.Code() == types.IsCallable {
		return IsCallable(arg, nil, IS_CALLABLE_CHECK_SILENT)
	} else if typ.Code() == types.IsIterable {
		return ZendIsIterable(arg)
	} else if typ.Code() == types.IsBool && arg.IsFalse() || arg.IsTrue() {
		return true
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return false
	} else {
		return ZendVerifyScalarTypeHint(typ.Code(), arg, CurrEX().IsArgUseStrictTypes())
	}
}
func ZendVerifyArgType(zf types.IFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) bool {
	var cur_arg_info *ZendArgInfo
	var ce *types.ClassEntry
	if arg_num <= zf.GetNumArgs() {
		cur_arg_info = zf.GetArgInfo()[arg_num-1]
	} else if zf.IsVariadic() {
		cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	} else {
		return true
	}
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope()) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return false
	}
	return true
}
func ZendVerifyRecvArgType(zf types.IFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = zf.GetArgInfo()[arg_num-1]
	var ce *types.ClassEntry
	b.Assert(arg_num <= zf.GetNumArgs())
	cur_arg_info = zf.GetArgInfo()[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope()) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf types.IFunction, arg_num uint32, arg *types.Zval, default_value *types.Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *types.ClassEntry
	b.Assert(arg_num > zf.GetNumArgs())
	b.Assert(zf.IsVariadic())
	cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope()) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyInternalArgTypes(fbc types.IFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = call.NumArgs()
	var p *types.Zval = call.Arg(1)
	var dummy_cache_slot any
	for i = 0; i < num_args; i++ {
		dummy_cache_slot = nil
		if !ZendVerifyArgType(fbc, i+1, p, nil, &dummy_cache_slot) {
			EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
			return 0
		}
		p++
	}
	return 1
}
func ZendMissingArgError(executeData *ZendExecuteData) {
	scope := executeData.GetFunc().GetScope()
	functionName := executeData.GetFunc().FunctionName()
	calleeName := functionName
	if scope != nil {
		calleeName = scope.Name() + "::" + functionName
	}

	requiredNumArgs := executeData.GetFunc().GetRequiredNumArgs()
	numArgs := executeData.GetFunc().GetNumArgs()

	var ptr = executeData.GetPrevExecuteData()
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetType()) {
		faults.ThrowError(faults.ZendCeArgumentCountError, fmt.Sprintf("Too few arguments to function %s(), %d passed in %s on line %d and %s %d expected", calleeName, executeData.NumArgs(), ptr.GetFunc().GetOpArray().GetFilename(), ptr.GetOpline().GetLineno(), lang.Cond(requiredNumArgs == numArgs, "exactly", "at least"), requiredNumArgs))
	} else {
		faults.ThrowError(faults.ZendCeArgumentCountError, fmt.Sprintf("Too few arguments to function %s(), %d passed and %s %d expected", calleeName, executeData.NumArgs(), lang.Cond(requiredNumArgs == numArgs, "exactly", "at least"), requiredNumArgs))
	}
}
func ZendVerifyReturnError(zf types.IFunction, ce *types.ClassEntry, value *types.Zval) {
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
	faults.TypeError(fmt.Sprintf("Return value of %s%s%s() must %s%s%s, %s%s returned", fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind))
}
func ZendVerifyMissingReturnType(zf types.IFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType().IsSet() && ret_info.GetType().Code() != types.IsVoid {
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
func ZendAssignToObjectDim(object *types.Zval, dim *types.Zval, value *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	object.Object().WriteDimension(dim, value)
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
}
func ZendBinaryOp(ret *types.Zval, op1 *types.Zval, op2 *types.Zval, opline *types.ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{operators.AddFunction, operators.SubFunction, operators.MulFunction, operators.DivFunction, operators.ModFunction, operators.ShiftLeftFunction, operators.ShiftRightFunction, operators.ConcatFunction, operators.BitwiseOrFunction, operators.BitwiseAndFunction, operators.BitwiseXorFunction, operators.PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *types.Zval, property *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var free_op_data1 ZendFreeOp
	var value *types.Zval
	var z *types.Zval
	var rv types.Zval
	var res types.Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if lang.Assign(&z, object.Object().ReadDimension(property, BP_VAR_R, &rv)) != nil {
		if z.IsObject() && z.Object().CanGet() {
			var rv2 types.Zval
			var value *types.Zval = z.Object().Get(&rv2)
			if z == &rv {
				// ZvalPtrDtor(&rv)
			}
			z.CopyValueFrom(value)
		}
		if ZendBinaryOp(&res, z, value, opline) == types.SUCCESS {
			object.Object().WriteDimension(property, &res)
		}
		if z == &rv {
			// ZvalPtrDtor(&rv)
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), &res)
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
func ZendBinaryAssignOpTypedRef(ref *types.Reference, value *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var z_copy types.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && ref.GetVal().IsString() {
		operators.ConcatFunction(ref.GetVal(), ref.GetVal(), value)
		b.Assert(ref.GetVal().IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, ref.GetVal(), value, opline)
	if ZendVerifyRefAssignableZval(ref, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		// ZvalPtrDtor(ref.GetVal())
		types.ZVAL_COPY_VALUE(ref.GetVal(), &z_copy)
	} else {
		// ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *types.PropertyInfo, zptr *types.Zval, value *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var z_copy types.Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && zptr.IsString() {
		operators.ConcatFunction(zptr, zptr, value)
		b.Assert(zptr.IsString() && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, opline)
	if ZendVerifyPropertyType(prop_info, &z_copy, executeData.IsCallUseStrictTypes()) != 0 {
		// ZvalPtrDtor(zptr)
		types.ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		// ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *types.Zval, type_ int, executeData *ZendExecuteData) ZendLong {
	dim = dim.DeRef()

	var offset ZendLong
	if !dim.IsLong() {
		switch dim.Type() {
		case types.IsString:
			if types.IsLong == operators.IsNumericString(dim.String(), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				faults.Error(faults.E_WARNING, fmt.Sprintf("Illegal string offset '%s'", dim.StringEx().GetVal()))
			}
		case types.IsUndef:
			ZVAL_UNDEFINED_OP2(executeData)
			fallthrough
		case types.IsDouble:
			fallthrough
		case types.IsNull:
			fallthrough
		case types.IsFalse:
			fallthrough
		case types.IsTrue:
			faults.Error(faults.E_NOTICE, "String offset cast occurred")
		default:
			ZendIllegalOffset()
		}
		offset = operators.ZvalGetLong(dim)
	} else {
		offset = dim.Long()
	}
	return offset
}

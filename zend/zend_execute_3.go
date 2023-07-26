package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendWrongStringOffset(executeData *ZendExecuteData) {
	var msg string
	var opline *types.ZendOp = executeData.GetOpline()
	var end *types.ZendOp
	var var_ uint32
	if EG__().GetException() != nil {
		return
	}
	switch opline.GetOpcode() {
	case ZEND_ASSIGN_OP:
		fallthrough
	case ZEND_ASSIGN_DIM_OP:
		fallthrough
	case ZEND_ASSIGN_OBJ_OP:
		fallthrough
	case ZEND_ASSIGN_STATIC_PROP_OP:
		msg = "Cannot use assign-op operators with string offsets"
	case ZEND_FETCH_DIM_W:
		fallthrough
	case ZEND_FETCH_DIM_RW:
		fallthrough
	case ZEND_FETCH_DIM_FUNC_ARG:
		fallthrough
	case ZEND_FETCH_DIM_UNSET:
		fallthrough
	case ZEND_FETCH_LIST_W:

		/* TODO: Encode the "reason" into opline->extended_value??? */

		var_ = opline.GetResult().GetVar()
		opline++
		end = CurrEX().GetFunc().GetOpArray().GetOpcodes() + CurrEX().GetFunc().GetOpArray().GetLast()
		for opline < end {
			if opline.GetOp1Type() == IS_VAR && opline.GetOp1().GetVar() == var_ {
				switch opline.GetOpcode() {
				case ZEND_FETCH_OBJ_W:
					fallthrough
				case ZEND_FETCH_OBJ_RW:
					fallthrough
				case ZEND_FETCH_OBJ_FUNC_ARG:
					fallthrough
				case ZEND_FETCH_OBJ_UNSET:
					fallthrough
				case ZEND_ASSIGN_OBJ:
					fallthrough
				case ZEND_ASSIGN_OBJ_OP:
					fallthrough
				case ZEND_ASSIGN_OBJ_REF:
					msg = "Cannot use string offset as an object"
				case ZEND_FETCH_DIM_W:
					fallthrough
				case ZEND_FETCH_DIM_RW:
					fallthrough
				case ZEND_FETCH_DIM_FUNC_ARG:
					fallthrough
				case ZEND_FETCH_DIM_UNSET:
					fallthrough
				case ZEND_FETCH_LIST_W:
					fallthrough
				case ZEND_ASSIGN_DIM:
					fallthrough
				case ZEND_ASSIGN_DIM_OP:
					msg = "Cannot use string offset as an array"
				case ZEND_ASSIGN_STATIC_PROP_OP:
					fallthrough
				case ZEND_ASSIGN_OP:
					msg = "Cannot use assign-op operators with string offsets"
				case ZEND_PRE_INC_OBJ:
					fallthrough
				case ZEND_PRE_DEC_OBJ:
					fallthrough
				case ZEND_POST_INC_OBJ:
					fallthrough
				case ZEND_POST_DEC_OBJ:
					fallthrough
				case ZEND_PRE_INC:
					fallthrough
				case ZEND_PRE_DEC:
					fallthrough
				case ZEND_POST_INC:
					fallthrough
				case ZEND_POST_DEC:
					msg = "Cannot increment/decrement string offsets"
				case ZEND_ASSIGN_REF:
					fallthrough
				case ZEND_ADD_ARRAY_ELEMENT:
					fallthrough
				case ZEND_INIT_ARRAY:
					fallthrough
				case ZEND_MAKE_REF:
					msg = "Cannot create references to/from string offsets"
				case ZEND_RETURN_BY_REF:
					fallthrough
				case ZEND_VERIFY_RETURN_TYPE:
					msg = "Cannot return string offsets by reference"
				case ZEND_UNSET_DIM:
					fallthrough
				case ZEND_UNSET_OBJ:
					msg = "Cannot unset string offsets"
				case ZEND_YIELD:
					msg = "Cannot yield string offsets by reference"
				case ZEND_SEND_REF:
					fallthrough
				case ZEND_SEND_VAR_EX:
					fallthrough
				case ZEND_SEND_FUNC_ARG:
					msg = "Only variables can be passed by reference"
				case ZEND_FE_RESET_RW:
					msg = "Cannot iterate on string offsets by reference"
				default:

				}
				break
			}
			if opline.GetOp2Type() == IS_VAR && opline.GetOp2().GetVar() == var_ {
				b.Assert(opline.GetOpcode() == ZEND_ASSIGN_REF)
				msg = "Cannot create references to/from string offsets"
				break
			}
			opline++
		}
	default:

	}
	faults.ThrowErrorEx(nil, msg)
}
func ZendWrongPropertyRead(property *types.Zval) {
	var property_name *types.String = operators.ZvalGetString(property)
	faults.Error(faults.E_NOTICE, "Trying to get property '%s' of non-object", property_name.GetVal())
}
func ZendDeprecatedFunction(fbc types.IFunction) {
	faults.Error(faults.E_DEPRECATED, "Function %s%s%s() is deprecated", lang.CondF1(fbc.GetScope() != nil, func() []byte { return fbc.GetScope().Name() }, ""), lang.Cond(fbc.GetScope() != nil, "::", ""), fbc.FunctionName())
}
func ZendAbstractMethod(fbc types.IFunction) {
	faults.ThrowError(nil, "Cannot call abstract method %s::%s()", fbc.GetScope().Name(), fbc.FunctionName())
}
func ZendAssignToStringOffset(str *types.Zval, dim *types.Zval, value *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var c uint8
	var string_len int
	var offset ZendLong
	offset = ZendCheckStringOffset(dim, BP_VAR_W, executeData)
	if EG__().GetException() != nil {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetUndef()
		}
		return
	}
	if offset < -ZendLong(str.String().GetLen()) {

		faults.Error(faults.E_WARNING, "Illegal string offset:  "+ZEND_LONG_FMT, offset)
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return
	}
	if !value.IsString() {

		/* Convert to string, just the time to pick the 1st byte */

		var tmp *types.String = operators.ZvalTryGetString(value)
		if tmp == nil {
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetUndef()
			}
			return
		}
		string_len = tmp.GetLen()
		c = uint8(tmp.GetStr()[0])
		// types.ZendStringReleaseEx(tmp, 0)
	} else {
		string_len = value.String().GetLen()
		c = uint8(value.String().GetStr()[0])
	}
	if string_len == 0 {

		/* Error on empty input string */

		faults.Error(faults.E_WARNING, "Cannot assign an empty string to a string offset")
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetNull()
		}
		return
	}
	if offset < 0 {
		offset += ZendLong(str.String().GetLen())
	}

	s := str.StringVal()
	if offset >= len(s) {
		newBytes := make([]byte, offset+1)
		copy(newBytes, []byte(s))
		for i := len(s); i < offset; i++ {
			newBytes[i] = ' '
		}
		newBytes[offset] = c

		str.SetString(string(newBytes))
	} else {
		newBytes := []byte(s)
		newBytes[offset] = c

		str.SetString(string(newBytes))
	}

	opline.Result().SetString(string(c))
}
func ZendGetPropNotAcceptingDouble(ref *types.Reference) *types.PropertyInfo {
	var prop *types.PropertyInfo
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
			if prop.GetType().Code() != types.IsDouble {
				return prop
			}
		}
	}
	return nil
}
func ZendThrowIncdecRefError(ref *types.Reference, opline *types.ZendOp) ZendLong {
	var error_prop *types.PropertyInfo = ZendGetPropNotAcceptingDouble(ref)

	/* Currently there should be no way for a typed reference to accept both int and double.
	 * Generalize this and the related property code once this becomes possible. */

	b.Assert(error_prop != nil)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		faults.TypeError("Cannot increment a reference held by property %s::$%s of type %sint past its maximal value", error_prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(error_prop.GetName()), lang.Cond(error_prop.GetType().AllowNull(), "?", ""))
		return ZEND_LONG_MAX
	} else {
		faults.TypeError("Cannot decrement a reference held by property %s::$%s of type %sint past its minimal value", error_prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(error_prop.GetName()), lang.Cond(error_prop.GetType().AllowNull(), "?", ""))
		return ZEND_LONG_MIN
	}
}
func ZendThrowIncdecPropError(prop *types.PropertyInfo, opline *types.ZendOp) ZendLong {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		faults.TypeError("Cannot increment property %s::$%s of type %s%s past its maximal value", prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MAX
	} else {
		faults.TypeError("Cannot decrement property %s::$%s of type %s%s past its minimal value", prop.GetCe().Name(), ZendGetUnmangledPropertyNameEx(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MIN
	}
}
func ZendIncdecTypedRef(ref *types.Reference, copy *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var tmp types.Zval
	var var_ptr *types.Zval = ref.GetVal()
	if copy == nil {
		copy = &tmp
	}
	types.ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		operators.IncrementFunction(var_ptr)
	} else {
		operators.DecrementFunction(var_ptr)
	}
	if var_ptr.IsDouble() && copy.IsLong() {
		var val ZendLong = ZendThrowIncdecRefError(ref, opline)
		var_ptr.SetLong(val)
	} else if ZendVerifyRefAssignableZval(ref, var_ptr, executeData.IsCallUseStrictTypes()) == 0 {
		// ZvalPtrDtor(var_ptr)
		var_ptr.CopyValueFrom(copy)
		copy.SetUndef()
	} else if copy == &tmp {
		// ZvalPtrDtor(&tmp)
	}
}
func ZendIncdecTypedProp(prop_info *types.PropertyInfo, var_ptr *types.Zval, copy *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var tmp types.Zval
	if copy == nil {
		copy = &tmp
	}
	types.ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		operators.IncrementFunction(var_ptr)
	} else {
		operators.DecrementFunction(var_ptr)
	}
	if var_ptr.IsDouble() && copy.IsLong() {
		var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
		var_ptr.SetLong(val)
	} else if ZendVerifyPropertyType(prop_info, var_ptr, executeData.IsCallUseStrictTypes()) == 0 {
		// ZvalPtrDtor(var_ptr)
		var_ptr.CopyValueFrom(copy)
		copy.SetUndef()
	} else if copy == &tmp {
		// ZvalPtrDtor(&tmp)
	}
}
func ZendPreIncdecPropertyZval(prop *types.Zval, prop_info *types.PropertyInfo, opline *types.ZendOp, executeData *ZendExecuteData) {
	if prop.IsLong() {
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			operators.FastLongIncrementFunction(prop)
		} else {
			operators.FastLongDecrementFunction(prop)
		}
		if !prop.IsLong() && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
			prop.SetLong(val)
		}
	} else {
		for {
			if prop.IsRef() {
				var ref *types.Reference = prop.Reference()
				prop = types.Z_REFVAL_P(prop)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendIncdecTypedRef(ref, nil, opline, executeData)
					break
				}
			}
			if prop_info != nil {
				ZendIncdecTypedProp(prop_info, prop, nil, opline, executeData)
			} else if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				operators.IncrementFunction(prop)
			} else {
				operators.DecrementFunction(prop)
			}
			break
		}
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), prop)
	}
}
func ZendPostIncdecPropertyZval(prop *types.Zval, prop_info *types.PropertyInfo, opline *types.ZendOp, executeData *ZendExecuteData) {
	if prop.IsLong() {
		opline.Result().SetLong(prop.Long())
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			operators.FastLongIncrementFunction(prop)
		} else {
			operators.FastLongDecrementFunction(prop)
		}
		if !prop.IsLong() && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, opline)
			prop.SetLong(val)
		}
	} else {
		if prop.IsRef() {
			var ref *types.Reference = prop.Reference()
			prop = types.Z_REFVAL_P(prop)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, opline.Result(), opline, executeData)
				return
			}
		}
		if prop_info != nil {
			ZendIncdecTypedProp(prop_info, prop, opline.Result(), opline, executeData)
		} else {
			types.ZVAL_COPY(opline.Result(), prop)
			if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				operators.IncrementFunction(prop)
			} else {
				operators.DecrementFunction(prop)
			}
		}
	}
}
func ZendPostIncdecOverloadedProperty(object *types.Zval, property *types.Zval, opline *types.ZendOp, executeData *ZendExecuteData) {
	var rv types.Zval
	var obj types.Zval
	var z *types.Zval
	var z_copy types.Zval
	obj.SetObject(object.Object())
	// 	obj.AddRefcount()
	z = obj.Object().ReadPropertyEx(property, BP_VAR_R, &rv)
	if EG__().GetException() != nil {
		// OBJ_RELEASE(obj.Object())
		opline.Result().SetUndef()
		return
	}
	if z.IsObject() && z.Object().CanGet() {
		var rv2 types.Zval
		var value *types.Zval = z.Object().Get(&rv2)
		z.CopyValueFrom(value)
	}
	types.ZVAL_COPY_DEREF(&z_copy, z)
	types.ZVAL_COPY(opline.Result(), &z_copy)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		operators.IncrementFunction(&z_copy)
	} else {
		operators.DecrementFunction(&z_copy)
	}
	obj.Object().WritePropertyEx(property, &z_copy)
}

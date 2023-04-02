package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if retval.IsNotUndef() {
					{
						goto fetch_obj_is_copy
					}

				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().Bucket(idx)))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.GetStr() || (p.IsStrKey() && p.StrKey() == offset.GetStrVal())) {
							retval = p.GetVal()
							{
								goto fetch_obj_is_copy
							}

						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.GetStr().GetStr())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
					{
						goto fetch_obj_is_copy
					}

				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	if op1.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op2_str)
			}

			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			EX_VAR(opline.GetResult().GetVar()).SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	{
		op2_str = op2.GetStr()
	}

	for {
		{
			if op1_str.GetLen() == 0 {
				{
					if op2.IsRefcounted() {
						op2_str.AddRefcount()
					}
				}
				EX_VAR(opline.GetResult().GetVar()).SetString(op2_str)
				types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		EX_VAR(opline.GetResult().GetVar()).SetString(str)
		{
			types.ZendStringReleaseEx(op1_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1()
					if EG__().GetException() != nil {
						return 0
					}
				}
				{
					function_name = RT_CONSTANT(opline, opline.GetOp2())
				}
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op1)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = RT_CONSTANT(opline, opline.GetOp2())
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) && obj == orig_obj {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), called_scope, fbc)
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		ZvalPtrDtorNogc(free_op1)
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			case_true:
				ZEND_VM_SMART_BRANCH_TRUE()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			case_false:
				ZEND_VM_SMART_BRANCH_FALSE()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto case_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		case_double:
			if d1 == d2 {
				goto case_true
			} else {
				goto case_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto case_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.IsArray() {
		var ht *types.Array
		var value *types.Zval
		var str *types.String
	isset_dim_obj_array:
		ht = container.GetArr()
	isset_again:
		if offset.IsString() {
			str = offset.GetStr()
			value = types.ZendHashFindInd(ht, str.GetStr())
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index_prop:
			value = ht.IndexFind(hval)
		} else {
			value = ZendFindArrayDimSlow(ht, offset, executeData)
			if EG__().GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL)
			{

				/* avoid exception check */

				ZEND_VM_SMART_BRANCH(result, 0)
				types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto isset_dim_obj_array
		}
	}
	if offset.GetU2Extra() == ZEND_EXTRA_VALUE {
		offset++
	}
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.GetType() != types.IS_OBJECT {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.GetType() != types.IS_OBJECT {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	subject = RT_CONSTANT(opline, opline.GetOp2())
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	EX_VAR(opline.GetResult().GetVar()).SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry
		{
			ce = CACHED_PTR(opline.GetExtendedValue())
			if ce == nil {
				ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp2()).GetStr(), (RT_CONSTANT(opline, opline.GetOp2()) + 1).GetStr(), ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					CACHE_PTR(opline.GetExtendedValue(), ce)
				}
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op2_str)
			}

			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op1_str)
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			EX_VAR(opline.GetResult().GetVar()).SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZvalPtrDtorNogc(free_op1)
		ZvalPtrDtorNogc(free_op2)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_not_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}

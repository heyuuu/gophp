// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	BooleanXorFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_TMP_VAR|IS_VAR, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			{
			}

			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			ZendWrongPropertyRead(offset)
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.GetTypeInfo() == types.IS_UNDEF {
		ZVAL_UNDEFINED_OP2()
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	}

	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
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
		ZEND_VM_NEXT_OPCODE()
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2()
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				{
				}

				EX_VAR(opline.GetResult().GetVar()).SetString(op2_str)
				types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				{
				}

				EX_VAR(opline.GetResult().GetVar()).SetString(op1_str)
				types.ZendStringReleaseEx(op2_str, 0)
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
		{
			types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc *ZendFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	{
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				if EG__().GetException() != nil {
					ZvalPtrDtorNogc(free_op1)
					HANDLE_EXCEPTION()
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			ZvalPtrDtorNogc(free_op1)
			HANDLE_EXCEPTION()
			break
		}
	}
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
						{
							ZvalPtrDtorNogc(free_op2)
						}
						HANDLE_EXCEPTION()
					}
				}
				{
				}

				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op2)
				ZvalPtrDtorNogc(free_op1)
				HANDLE_EXCEPTION()
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		{
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			ZvalPtrDtorNogc(free_op1)
			HANDLE_EXCEPTION()
		}
		{
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
	{
		ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		ZvalPtrDtorNogc(free_op1)
		if EG__().GetException() != nil {
			HANDLE_EXCEPTION()
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
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
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
			case_true:
				ZEND_VM_SMART_BRANCH_TRUE()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			case_false:
				ZEND_VM_SMART_BRANCH_FALSE()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
			ZvalPtrDtorNogc(free_op2)
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.IsArray() {
		var ht *types.Array
		var value *types.Zval
		var str *types.String
	isset_dim_obj_array:
		ht = container.GetArr()
	isset_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index_prop
				}
			}
			value = types.ZendHashFindInd(ht, str.GetStr())
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index_prop:
			value = ht.IndexFindH(hval)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto isset_again
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

				ZvalPtrDtorNogc(free_op2)
				ZEND_VM_SMART_BRANCH(result, 0)
				types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
				ZEND_VM_NEXT_OPCODE()
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
	{
	}

	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	subject = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	EX_VAR(opline.GetResult().GetVar()).SetTypeInfo(result)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = EX_VAR(opline.GetOp2().GetVar()).GetCe()
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varname *types.Zval
	var retval *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			ZvalPtrDtorNogc(free_op1)
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
		fetch_this:
			ZendFetchThisVar(type_, opline, executeData)
			{
				ZendTmpStringRelease(tmp_name)
			}
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
		if type_ == BP_VAR_W {
			retval = target_symbol_table.KeyAddNew(name.GetStr(), EG__().GetUninitializedZval())
		} else if type_ == BP_VAR_IS {
			retval = EG__().GetUninitializedZval()
		} else {
			faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
			if type_ == BP_VAR_RW {
				retval = target_symbol_table.KeyUpdate(name.GetStr(), EG__().GetUninitializedZval())
			} else {
				retval = EG__().GetUninitializedZval()
			}
		}
	} else if retval.IsIndirect() {
		retval = retval.GetZv()
		if retval.IsUndef() {
			if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
				goto fetch_this
			}
			if type_ == BP_VAR_W {
				retval.SetNull()
			} else if type_ == BP_VAR_IS {
				retval = EG__().GetUninitializedZval()
			} else {
				faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
				if type_ == BP_VAR_RW {
					retval.SetNull()
				} else {
					retval = EG__().GetUninitializedZval()
				}
			}
		}
	}
	if (opline.GetExtendedValue() & ZEND_FETCH_GLOBAL_LOCK) == 0 {
		ZvalPtrDtorNogc(free_op1)
	}
	{
		ZendTmpStringRelease(tmp_name)
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetIndirect(retval)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_R_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_R, executeData)
}
func ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_W, executeData)
}
func ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_RW, executeData)
}
func ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(fetch_type, executeData)
}
func ZEND_FETCH_UNSET_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_UNSET, executeData)
}
func ZEND_FETCH_IS_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_IS, executeData)
}
func ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	var free_op1 ZendFreeOp
	varname = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			ZvalPtrDtorNogc(free_op1)
			HANDLE_EXCEPTION()
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var free_op1 ZendFreeOp
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)

	{
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	ZvalPtrDtorNogc(free_op1)
	if value == nil {
		result = opline.GetExtendedValue() & ZEND_ISEMPTY
	} else {
		if value.IsIndirect() {
			value = value.GetZv()
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
			if value.IsReference() {
				value = types.Z_REFVAL_P(value)
			}
			result = value.GetType() > types.IS_NULL
		} else {
			result = !(IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				ZvalPtrDtorNogc(free_op1)
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var count ZendLong
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	for true {
		if op1.IsArray() {
			count = op1.GetArr().Count()
			break
		} else if op1.IsObject() {

			/* first, we check if the handler is defined */

			if types.Z_OBJ_HT_P(op1).GetCountElements() != nil {
				if types.SUCCESS == types.Z_OBJ_HT_P(op1).GetCountElements()(op1, &count) {
					break
				}
				if EG__().GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if InstanceofFunction(types.Z_OBJCE_P(op1), ZendCeCountable) != 0 {
				var retval types.Zval
				ZendCallMethodWith0Params(op1, nil, nil, "count", &retval)
				count = ZvalGetLong(&retval)
				ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
			continue
		} else if op1.GetType() <= types.IS_NULL {
			if op1.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			count = 0
		} else {
			count = 1
		}
		faults.Error(faults.E_WARNING, "%s(): Parameter must be an array or an object that implements Countable", b.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	EX_VAR(opline.GetResult().GetVar()).SetLong(count)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var free_op1 ZendFreeOp
		var op1 *types.Zval
		op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		for true {
			if op1.IsObject() {
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1()
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
			}
			break
		}
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_COPY_TMP_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	types.ZVAL_COPY(result, value)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_DIV_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = EX_VAR(opline.GetOp2().GetVar())
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
		ZEND_VM_NEXT_OPCODE()
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_SPACESHIP_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}

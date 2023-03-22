// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = EX_VAR(opline.GetOp2().GetVar())
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CV, BP_VAR_R, executeData)
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

	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_read_IS(container, EX_VAR(opline.GetOp2().GetVar()), IS_CV, opline, executeData)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = EX_VAR(opline.GetOp2().GetVar())
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
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
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
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
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
		function_name = EX_VAR(opline.GetOp2().GetVar())
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
						}
						HANDLE_EXCEPTION()
					}
				}
				{
				}

				ZendInvalidMethodCall(object, function_name)
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

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
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
func ZEND_CASE_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = EX_VAR(opline.GetOp2().GetVar())
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
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = EX_VAR(opline.GetOp2().GetVar())
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
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
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
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	subject = EX_VAR(opline.GetOp2().GetVar())
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
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	EX_VAR(opline.GetResult().GetVar()).SetTypeInfo(result)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_RETURN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	return_value = executeData.GetReturnValue()
	if retval_ptr.GetTypeInfo() == types.IS_UNDEF {
		retval_ptr = ZVAL_UNDEFINED_OP1()
		if return_value != nil {
			return_value.SetNull()
		}
	} else if return_value == nil {
		{
			if free_op1.IsRefcounted() && free_op1.DelRefcount() == 0 {
				RcDtorFunc(free_op1.GetCounted())
			}
		}
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
			{
			}

		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var free_op1 ZendFreeOp
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
			if !(executeData.GetReturnValue()) {
				ZvalPtrDtorNogc(free_op1)
			} else {
				if retval_ptr.IsReference() {
					types.ZVAL_COPY_VALUE(executeData.GetReturnValue(), retval_ptr)
					break
				}
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
				{
				}

			}
			break
		}
		retval_ptr = nil
		{
			b.Assert(retval_ptr != EG__().GetUninitializedZval())
			if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(retval_ptr.IsReference()) {
				faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
				if executeData.GetReturnValue() {
					executeData.GetReturnValue().
						SetNewRef(retval_ptr)
				}
				break
			}
		}
		if executeData.GetReturnValue() {
			if retval_ptr.IsReference() {
				retval_ptr.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(retval_ptr, 2)
			}
			executeData.GetReturnValue().
				SetReference(retval_ptr.GetRef())
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_GENERATOR_RETURN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var free_op1 ZendFreeOp
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
		{
		}

	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_THROW_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	for {
		if value.GetType() != types.IS_OBJECT {
			if value.IsReference() {
				value = types.Z_REFVAL_P(value)
				if value.IsObject() {
					break
				}
			}
			if value.IsUndef() {
				ZVAL_UNDEFINED_OP1()
				if EG__().GetException() != nil {
					HANDLE_EXCEPTION()
				}
			}
			faults.ThrowError(nil, "Can only throw objects")
			ZvalPtrDtorNogc(free_op1)
			HANDLE_EXCEPTION()
		}
		break
	}
	faults.ExceptionSave()
	{
	}

	faults.ThrowExceptionObject(value)
	faults.ExceptionRestore()
	HANDLE_EXCEPTION()
}
func ZEND_SEND_VAL_EX_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_val_by_ref:
		return zend_cannot_pass_by_ref_helper_SPEC(executeData)
	}
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAL_EX_SPEC_TMP_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_val_by_ref
	}
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_USER_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	var free_op1 ZendFreeOp
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CAST_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	var ht *types.Array
	expr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	switch opline.GetExtendedValue() {
	case types.IS_NULL:
		result.SetNull()
	case types.IS_BOOL:
		types.ZVAL_BOOL(result, ZendIsTrue(expr) != 0)
	case types.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			types.ZVAL_COPY_VALUE(result, expr)
			{
			}

			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			if expr.GetType() != types.IS_OBJECT || types.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewZendArray(1))
					expr = result.GetArr().IndexAddNewH(0, expr)

					{
						if expr.IsRefcounted() {
							expr.AddRefcount()
						}
					}
				} else {
					types.ZVAL_EMPTY_ARRAY(result)
				}
			} else {
				var obj_ht *types.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(expr).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					types.ZVAL_EMPTY_ARRAY(result)
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types.ZendSymtableToProptable(expr.GetArr())
				if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types.IS_NULL {
				ht = types.NewZendArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.ZSTR_SCALAR.GetStr(), expr)

				{
					if expr.IsRefcounted() {
						expr.AddRefcount()
					}
				}
			}
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FE_RESET_R_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, array_ptr)
		{
		}

		result.SetFePos(0)
		ZEND_VM_NEXT_OPCODE()
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			if types.Z_OBJ_P(array_ptr).GetProperties() != nil && types.Z_OBJ_P(array_ptr).GetProperties().GetRefcount() > 1 {
				if (types.Z_OBJ_P(array_ptr).GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
					types.Z_OBJ_P(array_ptr).GetProperties().DelRefcount()
				}
				types.Z_OBJ_P(array_ptr).SetProperties(types.ZendArrayDup(types.Z_OBJ_P(array_ptr).GetProperties()))
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			result = EX_VAR(opline.GetResult().GetVar())
			types.ZVAL_COPY_VALUE(result, array_ptr)
			{
			}

			if properties.GetNNumOfElements() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(types.ZendHashIteratorAdd(properties, 0))
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
			var is_empty types.ZendBool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
			ZvalPtrDtorNogc(free_op1)
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			} else if is_empty != 0 {
				ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
			} else {
				ZEND_VM_NEXT_OPCODE()
			}
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_RW_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = nil
		array_ref = array_ptr
		if array_ref.IsReference() {
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
	}

	if array_ptr.IsArray() {
		{
			if array_ptr == array_ref {
				array_ref.SetNewRef(array_ref)
				array_ptr = types.Z_REFVAL_P(array_ref)
			}
			array_ref.AddRefcount()
			types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), array_ref)
		}

		{
			types.SEPARATE_ARRAY(array_ptr)
		}
		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(types.ZendHashIteratorAdd(array_ptr.GetArr(), 0))
		{
		}
		ZEND_VM_NEXT_OPCODE()
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			{
				if array_ptr == array_ref {
					array_ref.SetNewRef(array_ref)
					array_ptr = types.Z_REFVAL_P(array_ref)
				}
				array_ref.AddRefcount()
				types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), array_ref)
			}

			if types.Z_OBJ_P(array_ptr).GetProperties() != nil && types.Z_OBJ_P(array_ptr).GetProperties().GetRefcount() > 1 {
				if (types.Z_OBJ_P(array_ptr).GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
					types.Z_OBJ_P(array_ptr).GetProperties().DelRefcount()
				}
				types.Z_OBJ_P(array_ptr).SetProperties(types.ZendArrayDup(types.Z_OBJ_P(array_ptr).GetProperties()))
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			if properties.GetNNumOfElements() == 0 {
				EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
				ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(types.ZendHashIteratorAdd(properties, 0))
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
			var is_empty types.ZendBool = ZendFeResetIterator(array_ptr, 1, opline, executeData)
			{
			}

			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			} else if is_empty != 0 {
				ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
			} else {
				ZEND_VM_NEXT_OPCODE()
			}
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
		{
		}

		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_END_SILENCE_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetErrorReporting() == 0 && EX_VAR(opline.GetOp1().GetVar()).GetLval() != 0 {
		EG__().SetErrorReporting(EX_VAR(opline.GetOp1().GetVar()).GetLval())
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_JMP_SET_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	ret = IZendIsTrue(value)
	if EG__().GetException() != nil {
		ZvalPtrDtorNogc(free_op1)
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	if ret != 0 {
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, value)

		{
			if result.IsRefcounted() {
				result.AddRefcount()
			}
		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_COALESCE_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, value)

		{
			if result.IsRefcounted() {
				result.AddRefcount()
			}
		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE()
}

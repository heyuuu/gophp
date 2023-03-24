package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_INIT_FCALL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fname *types.Zval
	var func_ *types.Zval
	var fbc *ZendFunction
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		fname = (*types.Zval)(RT_CONSTANT(opline, opline.GetOp2()))
		func_ = EG__().GetFunctionTable().KeyFind(fname.GetStr().GetStr())
		if func_ == nil {
			return zend_undefined_function_helper_SPEC(executeData)
		}
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = _zendVmStackPushCallFrameEx(opline.GetOp1().GetNum(), ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_RECV_INIT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32
	var param *types.Zval
	for {
		arg_num = opline.GetOp1().GetNum()
		param = EX_VAR(opline.GetResult().GetVar())
		if arg_num > executeData.NumArgs() {
			var default_value *types.Zval = RT_CONSTANT(opline, opline.GetOp2())
			if default_value.IsConstant() {
				var cache_val *types.Zval = (*types.Zval)(CACHE_ADDR(default_value.GetCacheSlot()))

				/* we keep in cache only not refcounted values */

				if cache_val.IsNotUndef() {
					types.ZVAL_COPY_VALUE(param, cache_val)
				} else {
					types.ZVAL_COPY(param, default_value)
					if ZvalUpdateConstantEx(param, executeData.GetFunc().op_array.scope) != types.SUCCESS {
						ZvalPtrDtorNogc(param)
						param.SetUndef()
						HANDLE_EXCEPTION()
					}
					if !(param.IsRefcounted()) {
						types.ZVAL_COPY_VALUE(cache_val, param)
					}
				}
				goto recv_init_check_type
			} else {
				types.ZVAL_COPY(param, default_value)
			}
		} else {
		recv_init_check_type:
			if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_HAS_TYPE_HINTS) != 0 {
				var default_value *types.Zval = RT_CONSTANT(opline, opline.GetOp2())
				if ZendVerifyRecvArgType(executeData.GetFunc(), arg_num, param, default_value, CACHE_ADDR(opline.GetExtendedValue())) == 0 {
					HANDLE_EXCEPTION()
				}
			}
		}
		if b.PreInc(&opline).opcode != ZEND_RECV_INIT {
			break
		}
	}
	OPLINE = opline
	return 0
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
try_function_name:
	if function_name.IsString() {
		call = ZendInitDynamicCallString(function_name.GetStr(), opline.GetExtendedValue())
	} else if function_name.IsObject() {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else if function_name.IsReference() {
		function_name = types.Z_REFVAL_P(function_name)
		goto try_function_name
	} else {
		if function_name.IsUndef() {
			ZVAL_UNDEFINED_OP2()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	ZvalPtrDtorNogc(free_op2)
	if call == nil {
		HANDLE_EXCEPTION()
	}
	{
		if EG__().GetException() != nil {
			if call != nil {
				if call.GetFunc().IsCallViaTrampoline() {
					types.ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					ZendFreeTrampoline(call.GetFunc())
				}
				ZendVmStackFreeCallFrame(call)
			}
			HANDLE_EXCEPTION()
		}
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_RECV_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	if arg_num > executeData.NumArgs() {
		ZendMissingArgError(executeData)
		HANDLE_EXCEPTION()
	} else {
		var param *types.Zval = EX_VAR(opline.GetResult().GetVar())
		if ZendVerifyRecvArgType(executeData.GetFunc(), arg_num, param, nil, CACHE_ADDR(opline.GetOp2().GetNum())) == 0 {
			HANDLE_EXCEPTION()
		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	var arg_count uint32 = executeData.NumArgs()
	var params *types.Zval
	params = EX_VAR(opline.GetResult().GetVar())
	if arg_num <= arg_count {
		var param *types.Zval
		ArrayInitSize(params, arg_count-arg_num+1)
		types.ZendHashRealInitPacked(params.GetArr())
		for {
			fillScope := types.PackedFillStart(params.GetArr())
			param = executeData.VarNum(executeData.GetFunc().op_array.last_var + executeData.GetFunc().op_array.T)
			if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_HAS_TYPE_HINTS) != 0 {
				ZEND_ADD_CALL_FLAG(executeData, ZEND_CALL_FREE_EXTRA_ARGS)
				for {
					ZendVerifyVariadicArgType(executeData.GetFunc(), arg_num, param, nil, CACHE_ADDR(opline.GetOp2().GetNum()))

					param.TryAddRefcount()

					fillScope.FillSet(param)
					fillScope.FillNext()
					param++
					if b.PreInc(&arg_num) > arg_count {
						break
					}
				}
			} else {
				for {

					param.TryAddRefcount()

					fillScope.FillSet(param)
					fillScope.FillNext()
					param++
					if b.PreInc(&arg_num) > arg_count {
						break
					}
				}
			}
			fillScope.FillEnd()
			break
		}
	} else {
		params.SetEmptyArray()
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = EX_VAR(opline.GetOp2().GetVar())
try_function_name:
	if function_name.IsString() {
		call = ZendInitDynamicCallString(function_name.GetStr(), opline.GetExtendedValue())
	} else if function_name.IsObject() {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else if function_name.IsReference() {
		function_name = types.Z_REFVAL_P(function_name)
		goto try_function_name
	} else {
		if function_name.IsUndef() {
			ZVAL_UNDEFINED_OP2()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		HANDLE_EXCEPTION()
	}
	{
		if EG__().GetException() != nil {
			if call != nil {
				if call.GetFunc().IsCallViaTrampoline() {
					types.ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					ZendFreeTrampoline(call.GetFunc())
				}
				ZendVmStackFreeCallFrame(call)
			}
			HANDLE_EXCEPTION()
		}
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BW_NOT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	if op1.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(^(op1.GetLval()))
		ZEND_VM_NEXT_OPCODE()
	}
	BitwiseNotFunction(EX_VAR(opline.GetResult().GetVar()), op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_NOT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) == 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ECHO_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var z *types.Zval
	z = RT_CONSTANT(opline, opline.GetOp1())
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else {
		}

		types.ZendStringReleaseEx(str, 0)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_JMPZ_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPNZ_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		ZEND_VM_NEXT_OPCODE()
	}
	if IZendIsTrue(val) != 0 {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	} else {
		opline++
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZNZ_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZ_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		opline++
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPNZ_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()

		{
			ZEND_VM_NEXT_OPCODE()
		}
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		opline++
	}
	ZEND_VM_JMP(opline)
}
func ZEND_RETURN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = RT_CONSTANT(opline, opline.GetOp1())
	return_value = executeData.GetReturnValue()

	if return_value == nil {
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
			{

				return_value.TryAddRefcount()

			}
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = RT_CONSTANT(opline, opline.GetOp1())
			if !(executeData.GetReturnValue()) {
			} else {
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
				{
					retval_ptr.TryAddRefcount()
				}
			}
			break
		}
		retval_ptr = nil
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
func ZEND_GENERATOR_RETURN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = RT_CONSTANT(opline, opline.GetOp1())

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
		{

			generator.GetRetval().TryAddRefcount()

		}
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_THROW_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp1())
	for {
		{
			faults.ThrowError(nil, "Can only throw objects")
			HANDLE_EXCEPTION()
		}
		break
	}
	faults.ExceptionSave()
	{
		value.TryAddRefcount()
	}
	faults.ThrowExceptionObject(value)
	faults.ExceptionRestore()
	HANDLE_EXCEPTION()
}
func ZEND_CATCH_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ce *types.ClassEntry
	var catch_ce *types.ClassEntry
	var exception *types.ZendObject
	var ex *types.Zval

	/* Check whether an exception has been thrown, if not, jump over code */

	faults.ExceptionRestore()
	if EG__().GetException() == nil {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	catch_ce = CACHED_PTR(opline.GetExtendedValue() & ^ZEND_LAST_CATCH)
	if catch_ce == nil {
		catch_ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_NO_AUTOLOAD)
		CACHE_PTR(opline.GetExtendedValue() & ^ZEND_LAST_CATCH, catch_ce)
	}
	ce = EG__().GetException().GetCe()
	if ce != catch_ce {
		if catch_ce == nil || InstanceofFunction(ce, catch_ce) == 0 {
			if (opline.GetExtendedValue() & ZEND_LAST_CATCH) != 0 {
				faults.RethrowException(executeData)
				HANDLE_EXCEPTION()
			}
			ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
		}
	}
	exception = EG__().GetException()
	ex = EX_VAR(opline.GetResult().GetVar())

	/* Always perform a strict assignment. There is a reasonable expectation that if you
	 * write "catch (Exception $e)" then $e will actually be instanceof Exception. As such,
	 * we should not permit coercion to string here. */

	var tmp types.Zval
	tmp.SetObject(exception)
	EG__().SetException(nil)
	ZendAssignToVariable(ex, &tmp, IS_TMP_VAR, 1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SEND_VAL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp1())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{

		arg.TryAddRefcount()

	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAL_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_val_by_ref:
		return zend_cannot_pass_by_ref_helper_SPEC(executeData)
	}
	value = RT_CONSTANT(opline, opline.GetOp1())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{

		arg.TryAddRefcount()

	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAL_EX_SPEC_CONST_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_val_by_ref
	}
	value = RT_CONSTANT(opline, opline.GetOp1())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{

		arg.TryAddRefcount()

	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_USER_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = RT_CONSTANT(opline, opline.GetOp1())
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) != 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CLONE_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = RT_CONSTANT(opline, opline.GetOp1())
	for {
		{
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			faults.ThrowError(nil, "__clone method called on non-object")
			HANDLE_EXCEPTION()
		}
		break
	}
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().op_array.scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
		}
	}
	EX_VAR(opline.GetResult().GetVar()).SetObject(clone_call(obj))
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CAST_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	var ht *types.Array
	expr = RT_CONSTANT(opline, opline.GetOp1())
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
		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			types.ZVAL_COPY_VALUE(result, expr)
			{

				result.TryAddRefcount()

			}

			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			{
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
					expr = result.GetArr().IndexAddNew(0, expr)
					{

						expr.TryAddRefcount()

					}

				} else {
					result.SetEmptyArray()
				}
			}

			/* fast copy */

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
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.ZSTR_SCALAR.GetStr(), expr)
				{

					expr.TryAddRefcount()

				}

			}
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INCLUDE_OR_EVAL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *ZendOpArray
	var inc_filename *types.Zval
	inc_filename = RT_CONSTANT(opline, opline.GetOp1())
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	if EG__().GetException() != nil {
		if new_op_array != ZEND_FAKE_OP_ARRAY && new_op_array != nil {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		}
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	} else if new_op_array == ZEND_FAKE_OP_ARRAY {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetTrue()
		}
	} else if new_op_array != nil {
		var return_value *types.Zval = nil
		var call *ZendExecuteData
		if RETURN_VALUE_USED(opline) {
			return_value = EX_VAR(opline.GetResult().GetVar())
		}
		new_op_array.SetScope(executeData.GetFunc().op_array.scope)
		call = ZendVmStackPushCallFrame(executeData.GetThis().GetTypeInfo()&ZEND_CALL_HAS_THIS|ZEND_CALL_NESTED_CODE|ZEND_CALL_HAS_SYMBOL_TABLE, (*ZendFunction)(new_op_array), 0, executeData.GetThis().GetPtr())
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			call.SetSymbolTable(executeData.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(executeData)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			UNDEF_RESULT()
			HANDLE_EXCEPTION()
		}
	} else if RETURN_VALUE_USED(opline) {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_FE_RESET_R_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = RT_CONSTANT(opline, opline.GetOp1())
	if array_ptr.IsArray() {
		result = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, array_ptr)
		if result.IsRefcounted() {
			array_ptr.AddRefcount()
		}
		result.SetFePos(0)
		ZEND_VM_NEXT_OPCODE()
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_RW_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var array_ref *types.Zval

	{
		array_ptr = RT_CONSTANT(opline, opline.GetOp1())
		array_ref = array_ptr
	}
	if array_ptr.IsArray() {

		{
			array_ref = EX_VAR(opline.GetResult().GetVar())
			array_ref.SetNewRef(array_ptr)
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
		{
			array_ptr.SetArray(types.ZendArrayDup(array_ptr.GetArr()))
		}

		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(types.ZendHashIteratorAdd(array_ptr.GetArr(), 0))
		ZEND_VM_NEXT_OPCODE()
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_JMP_SET_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = RT_CONSTANT(opline, opline.GetOp1())
	ret = IZendIsTrue(value)
	if EG__().GetException() != nil {
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	if ret != 0 {
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, value)
		{

			result.TryAddRefcount()

		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_COALESCE_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	value = RT_CONSTANT(opline, opline.GetOp1())
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, value)
		{

			result.TryAddRefcount()

		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_QM_ASSIGN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	value = RT_CONSTANT(opline, opline.GetOp1())
	{
		types.ZVAL_COPY_VALUE(result, value)
		{

			result.TryAddRefcount()

		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	DoBindClass(RT_CONSTANT(opline, opline.GetOp1()), b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types.String { return RT_CONSTANT(opline, opline.GetOp2()).GetStr() }, nil))
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_FROM_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)

		val.TryAddRefcount()

		generator.GetValues().GetFePos() = 0
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if RETURN_VALUE_USED(opline) {
		EX_VAR(opline.GetResult().GetVar()).SetNull()
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE()

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_STRLEN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp1())
	if value.IsString() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetStr().GetLen())
		ZEND_VM_NEXT_OPCODE()
	} else {
		var strict types.ZendBool
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					EX_VAR(opline.GetResult().GetVar()).SetLong(str.GetLen())
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.GetType()))
			}
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			break
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}

package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = opline.getZvalPtrVar2(&free_op2)
	{
		for {
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			ZendWrongPropertyRead(offset)
			opline.GetResultZval().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2()
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.GetResultZval())
	if retval != opline.GetResultZval() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = opline.getZvalPtrVar2(&free_op2)
	{
		for {
			opline.GetResultZval().SetNull()
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, opline.GetResultZval())
	if retval != opline.GetResultZval() {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, opline.getZvalPtrVar2(&free_op2), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.getZvalPtrVar2(&free_op2)
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
			{
				opline.GetResultZval().SetStringCopy(op1_str)
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.GetResultZval().SetString(str)
			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	{
		op1_str = op1.GetStr()
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
			if op2_str.GetLen() == 0 {
				{
					if op1.IsRefcounted() {
						op1_str.AddRefcount()
					}
				}
				opline.GetResultZval().SetString(op1_str)
				types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.GetResultZval().SetString(str)
		{
			types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = RT_CONSTANT(opline, opline.GetOp1())
	{
		function_name = opline.getZvalPtrVar2(&free_op2)
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
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	{
		for {
			{
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op2)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				ZvalPtrDtorNogc(opline.GetOp2Zval())
				return 0
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
		}
	}

	{
		var free_op2 ZendFreeOp
		function_name = opline.getZvalPtrVar2(&free_op2)
		{
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
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
			ZvalPtrDtorNogc(free_op2)
		}
	}

	if !fbc.IsStatic() {
		if executeData.GetThis().IsObject() && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:
		/* previous opcode is ZEND_FETCH_CLASS */

		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ types.IFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = opline.getZvalPtrVar2(&free_op2)
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				ZvalPtrDtorNogc(free_op2)
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if func_.IsClosure() {

			/* Delay closure destruction until its invocation */

			ZEND_CLOSURE_OBJECT(func_).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if func_.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= ZEND_CALL_HAS_THIS
			}
		} else if fcc.GetObject() != nil {
			fcc.GetObject().AddRefcount()
			object_or_called_scope = fcc.GetObject()
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
		}
		ZvalPtrDtorNogc(free_op2)
		if EG__().GetException() != nil {
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				ZendObjectRelease(ZEND_CLOSURE_OBJECT(func_))
			} else if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
				ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", RT_CONSTANT(opline, opline.GetOp1()).GetStr().GetVal(), error)
		Efree(error)
		ZvalPtrDtorNogc(free_op2)
		if EG__().GetException() != nil {
			return 0
		}
		func_ = (types.IFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	{
		expr_ptr = RT_CONSTANT(opline, opline.GetOp1())

		{
			expr_ptr.TryAddRefcount()
		}

	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.getZvalPtrVar2(&free_op2)
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			opline.GetResultZval().GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			opline.GetResultZval().GetArr().IndexUpdate(hval, expr_ptr)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else if offset.IsDouble() {
			hval = DvalToLval(offset.GetDval())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = types.Z_RES_HANDLE_P(offset)
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2()
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
		ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.GetResultZval()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = opline.getZvalPtrVar2(&free_op2)
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
			value = ht.IndexFind(hval)
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
				types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else {
	}
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = opline.getZvalPtrVar2(&free_op2)
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = RT_CONSTANT(opline, opline.GetOp1())
	subject = opline.getZvalPtrVar2(&free_op2)
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
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.GetResultZval().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_CONST_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = RT_CONSTANT(opline, opline.GetOp1())
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = RT_CONSTANT(opline, opline.GetOp1())

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetLval())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.GetResultZval())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CONST_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = RT_CONSTANT(opline, opline.GetOp1())
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = RT_CONSTANT(opline, opline.GetOp1())

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.getZvalPtrVar2(&free_op2)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetLval())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.GetResultZval())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func zend_fetch_var_address_helper_SPEC_CONST_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var retval *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = RT_CONSTANT(opline, opline.GetOp1())
	{
		name = varname.GetStr()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
		fetch_this:
			ZendFetchThisVar(type_, opline, executeData)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
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
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else {
		opline.GetResultZval().SetIndirect(retval)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_R_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_R, executeData)
}
func ZEND_FETCH_W_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_W, executeData)
}
func ZEND_FETCH_RW_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_RW, executeData)
}
func ZEND_FETCH_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(fetch_type, executeData)
}
func ZEND_FETCH_UNSET_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_UNSET, executeData)
}
func ZEND_FETCH_IS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_IS, executeData)
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				return 0
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
		}
	}

	{
		if ce.GetConstructor() == nil {
			faults.ThrowError(nil, "Cannot call constructor")
			return 0
		}
		if executeData.GetThis().IsObject() && types.Z_OBJ(executeData.GetThis()).GetCe() != ce.GetConstructor().GetScope() && ce.GetConstructor().IsPrivate() {
			faults.ThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	if !fbc.IsStatic() {
		if executeData.GetThis().IsObject() && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:
		/* previous opcode is ZEND_FETCH_CLASS */

		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

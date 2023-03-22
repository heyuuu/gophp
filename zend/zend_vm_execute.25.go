// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	{
	}

assign_object:
	{
	}

	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
	{
	}

assign_object:
	{
	}

	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	{

		{
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, executeData)
		}
	}

	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)
	{

		{
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, executeData)
		}
	}

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ROPE_INIT_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(EX_VAR(opline.GetResult().GetVar()))

	{
		var_ = EX_VAR(opline.GetOp2().GetVar())
		if var_.IsString() {
			{
				rope[0] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[0] = ZvalGetStringFunc(var_)
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		class_name = EX_VAR(opline.GetOp2().GetVar())
	try_class_name:
		if class_name.IsObject() {
			EX_VAR(opline.GetResult().GetVar()).SetCe(types.Z_OBJCE_P(class_name))
		} else if class_name.IsString() {
			EX_VAR(opline.GetResult().GetVar()).SetCe(ZendFetchClass(class_name.GetStr(), opline.GetOp1().GetNum()))
		} else if class_name.IsReference() {
			class_name = types.Z_REFVAL_P(class_name)
			goto try_class_name
		} else {
			if class_name.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				if EG__().GetException() != nil {
					HANDLE_EXCEPTION()
				}
			}
			faults.ThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc *ZendFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
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
					HANDLE_EXCEPTION()
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			HANDLE_EXCEPTION()
			break
		}
	}
	{
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
			HANDLE_EXCEPTION()
		}
		{
		}
		{
		}

		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		{
		}

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
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			HANDLE_EXCEPTION()
		}
	}

	{
		function_name = EX_VAR(opline.GetOp2().GetVar())
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
							HANDLE_EXCEPTION()
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					HANDLE_EXCEPTION()
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			HANDLE_EXCEPTION()
		}
		{
		}

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
		}
	}

	if !fbc.IsStatic() {
		if executeData.GetThis().u1.v.type_ == types.IS_OBJECT && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			if executeData.GetThis().u1.v.type_ == types.IS_OBJECT {
				ce = types.Z_OBJCE(executeData.GetThis())
			} else {
				ce = executeData.GetThis().GetCe()
			}
		}
		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewZendArray(0))
		ZEND_VM_NEXT_OPCODE()
	}
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	for {
		{
		}

		types.Z_OBJ_HT_P(container).GetUnsetProperty()(container, offset, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	{
	}

	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
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

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	{
		var key *types.Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)

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

		generator.SetSendTarget(EX_VAR(opline.GetResult().GetVar()))
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE()

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_BW_NOT_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	if op1.GetTypeInfo() == types.IS_LONG {
		EX_VAR(opline.GetResult().GetVar()).SetLong(^(op1.GetLval()))
		ZEND_VM_NEXT_OPCODE()
	}
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	BitwiseNotFunction(EX_VAR(opline.GetResult().GetVar()), op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_NOT_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) == 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func zend_pre_inc_helper_SPEC_CV(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
		ZEND_VM_NEXT_OPCODE()
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1()
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.GetRef()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, nil, opline, executeData)
				break
			}
		}
		IncrementFunction(var_ptr)
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)

		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_inc_helper_SPEC_CV(executeData)
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_inc_helper_SPEC_CV(executeData)
}
func zend_pre_dec_helper_SPEC_CV(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
		ZEND_VM_NEXT_OPCODE()
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1()
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.GetRef()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, nil, opline, executeData)
				break
			}
		}
		DecrementFunction(var_ptr)
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_DEC_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		FastLongDecrementFunction(var_ptr)

		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_dec_helper_SPEC_CV(executeData)
}
func ZEND_PRE_DEC_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		FastLongDecrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_dec_helper_SPEC_CV(executeData)
}
func zend_post_inc_helper_SPEC_CV(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsError() {
		EX_VAR(opline.GetResult().GetVar()).SetNull()
		ZEND_VM_NEXT_OPCODE()
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1()
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.GetRef()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, EX_VAR(opline.GetResult().GetVar()), opline, executeData)
				break
			}
		}
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		IncrementFunction(var_ptr)
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
		FastLongIncrementFunction(var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_post_inc_helper_SPEC_CV(executeData)
}
func zend_post_dec_helper_SPEC_CV(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsError() {
		EX_VAR(opline.GetResult().GetVar()).SetNull()
		ZEND_VM_NEXT_OPCODE()
	}
	if var_ptr.IsUndef() {
		var_ptr.SetNull()
		ZVAL_UNDEFINED_OP1()
	}
	for {
		if var_ptr.IsReference() {
			var ref *types.ZendReference = var_ptr.GetRef()
			var_ptr = types.Z_REFVAL_P(var_ptr)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, EX_VAR(opline.GetResult().GetVar()), opline, executeData)
				break
			}
		}
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		DecrementFunction(var_ptr)
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_DEC_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	if var_ptr.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
		FastLongDecrementFunction(var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_post_dec_helper_SPEC_CV(executeData)
}
func ZEND_ECHO_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var z *types.Zval
	z = EX_VAR(opline.GetOp1().GetVar())
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else if z.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		types.ZendStringReleaseEx(str, 0)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_JMPZ_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.GetTypeInfo() == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPNZ_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.GetTypeInfo() == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_NEXT_OPCODE()
	}
	if IZendIsTrue(val) != 0 {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	} else {
		opline++
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZNZ_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.GetTypeInfo() == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZ_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if val.GetTypeInfo() == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
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
func ZEND_JMPNZ_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.GetTypeInfo() == types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if val.GetTypeInfo() == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
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
func ZEND_RETURN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = EX_VAR(opline.GetOp1().GetVar())
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
func ZEND_RETURN_BY_REF_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
			if !(executeData.GetReturnValue()) {
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
		retval_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
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
func ZEND_GENERATOR_RETURN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)

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
func ZEND_THROW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
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
func ZEND_SEND_VAR_SPEC_CV_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	varptr = EX_VAR(opline.GetOp1().GetVar())
	if varptr.GetTypeInfo() == types.IS_UNDEF {
		ZVAL_UNDEFINED_OP1()
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAR_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_SEND_VAR_SPEC_CV_INLINE_HANDLER(executeData)
}
func ZEND_SEND_REF_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	varptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	if varptr.IsError() {
		arg.SetNewEmptyRef()
		types.Z_REFVAL_P(arg).SetNull()
		ZEND_VM_NEXT_OPCODE()
	}
	if varptr.IsReference() {
		varptr.AddRefcount()
	} else {
		types.ZVAL_MAKE_REF_EX(varptr, 2)
	}
	arg.SetReference(varptr.GetRef())
	ZEND_VM_NEXT_OPCODE()
}

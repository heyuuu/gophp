package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc *types.ZendFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			ZvalPtrDtorNogc(EX_VAR(opline.GetOp2().GetVar()))
			HANDLE_EXCEPTION()
		}
	}

	{
		var free_op2 ZendFreeOp
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
					ZvalPtrDtorNogc(free_op2)
					HANDLE_EXCEPTION()
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
			HANDLE_EXCEPTION()
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
				HANDLE_EXCEPTION()
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			if executeData.GetThis().IsObject() {
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
func ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		ZEND_VM_NEXT_OPCODE()
	}
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	for {
		types.Z_OBJ_HT_P(container).GetUnsetProperty()(container, offset, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
		break
	}
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_UNUSED_TMP_HANDLER(executeData *ZendExecuteData) int {
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
func ZEND_YIELD_SPEC_UNUSED_VAR_HANDLER(executeData *ZendExecuteData) int {
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
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_FETCH_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	{
		EX_VAR(opline.GetResult().GetVar()).SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc *types.ZendFunction
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
		if ce.GetConstructor() == nil {
			faults.ThrowError(nil, "Cannot call constructor")
			HANDLE_EXCEPTION()
		}
		if executeData.GetThis().IsObject() && types.Z_OBJ(executeData.GetThis()).GetCe() != ce.GetConstructor().GetScope() && ce.GetConstructor().IsPrivate() {
			faults.ThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			HANDLE_EXCEPTION()
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
				HANDLE_EXCEPTION()
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			if executeData.GetThis().IsObject() {
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
func ZEND_VERIFY_RETURN_TYPE_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	{
		ZendVerifyMissingReturnType(executeData.GetFunc(), CACHE_ADDR(opline.GetOp2().GetNum()))
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_NEW_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var result *types.Zval
	var constructor *types.ZendFunction
	var ce *types.ClassEntry
	var call *ZendExecuteData

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
	}

	result = EX_VAR(opline.GetResult().GetVar())
	if ObjectInitEx(result, ce) != types.SUCCESS {
		result.SetUndef()
		HANDLE_EXCEPTION()
	}
	constructor = types.Z_OBJ_HT_P(result).GetGetConstructor()(result.GetObj())
	if constructor == nil {
		if EG__().GetException() != nil {
			HANDLE_EXCEPTION()
		}

		/* If there are no arguments, skip over the DO_FCALL opcode. We check if the next
		 * opcode is DO_FCALL in case EXT instructions are used. */

		if opline.GetExtendedValue() == 0 && (opline+1).GetOpcode() == ZEND_DO_FCALL {
			OPLINE = executeData.GetOpline() + 2
			return 0
		}

		/* Perform a dummy function call */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION, (*types.ZendFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

		/* Perform a dummy function call */

	} else {
		if constructor.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(constructor.GetOpArray())) {
			InitFuncRunTimeCache(constructor.GetOpArray())
		}

		/* We are not handling overloaded classes right now */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION|ZEND_CALL_RELEASE_THIS|ZEND_CALL_HAS_THIS, constructor, opline.GetExtendedValue(), result.GetObj())
		result.AddRefcount()
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		ZEND_VM_NEXT_OPCODE()
	}
}
func ZEND_YIELD_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
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
func ZEND_FETCH_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		result.SetObject(executeData.GetThis().GetObj())
		result.AddRefcount()
		ZEND_VM_NEXT_OPCODE()
	} else {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
}
func ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), (opline.GetExtendedValue()&ZEND_ISEMPTY^executeData.GetThis().IsObject()) != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	{
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_class() called without object from outside a class")
			EX_VAR(opline.GetResult().GetVar()).SetFalse()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
			EX_VAR(opline.GetResult().GetVar()).SetStringCopy(executeData.GetFunc().common.scope.name)
			ZEND_VM_NEXT_OPCODE()
		}
	}

}
func ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(types.Z_OBJCE(executeData.GetThis()).GetName())
	} else if executeData.GetThis().GetCe() != nil {
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(types.Z_CE(executeData.GetThis()).GetName())
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_called_class() called from outside a class")
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	EX_VAR(opline.GetResult().GetVar()).SetLong(executeData.NumArgs())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_FUNC_GET_ARGS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = executeData.NumArgs()

	{
		skip = 0
		result_size = arg_count
	}
	if result_size != 0 {
		var first_extra_arg uint32 = executeData.GetFunc().op_array.num_args
		ht = types.NewArray(result_size)
		EX_VAR(opline.GetResult().GetVar()).SetArray(ht)
		types.ZendHashRealInitPacked(ht)
		fillScope := types.PackedFillStart(ht)
		var p *types.Zval
		var q *types.Zval
		var i uint32 = skip
		p = executeData.VarNum(i)
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if !q.IsUndef() {
					q = types.ZVAL_DEREF(q)

					q.TryAddRefcount()

					fillScope.FillSet(q)
				} else {
					fillScope.FillSetNull()
				}
				fillScope.FillNext()
				p++
				i++
			}
			if skip < first_extra_arg {
				skip = 0
			} else {
				skip -= first_extra_arg
			}
			p = executeData.VarNum(executeData.GetFunc().op_array.last_var + executeData.GetFunc().op_array.T + skip)
		}
		for i < arg_count {
			q = p
			if !q.IsUndef() {
				q = types.ZVAL_DEREF(q)

				q.TryAddRefcount()

				fillScope.FillSet(q)
			} else {
				fillScope.FillSetNull()
			}
			fillScope.FillNext()
			p++
			i++
		}
		fillScope.FillEnd()
		ht.SetNNumOfElements(result_size)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetEmptyArray()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	for {
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
	assign_op_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {
				var orig_zptr *types.Zval = zptr
				var ref *types.ZendReference
				for {
					if zptr.IsReference() {
						ref = zptr.GetRef()
						zptr = types.Z_REFVAL_P(zptr)
						if ZEND_REF_HAS_TYPE_SOURCES(ref) {
							ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
							break
						}
					}

					{
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, executeData)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), zptr)
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, executeData)
		}
		break
	}
	FREE_OP(free_op_data)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	for {
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {

				{
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), zptr)
				}
				ZendPreIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_OBJ_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	for {
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			} else {

				{
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), zptr)
				}
				ZendPostIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = EX_VAR(opline.GetOp2().GetVar())
	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
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
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
assign_object:
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
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
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
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
assign_object:
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

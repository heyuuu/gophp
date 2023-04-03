package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_ROPE_INIT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(opline.GetResultZval())
	{
		var_ = RT_CONSTANT(opline, opline.GetOp2())
		rope[0] = var_.GetStr()

		var_.TryAddRefcount()

	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		var ce *types.ClassEntry = CACHED_PTR(opline.GetExtendedValue())
		if ce == nil {
			class_name = RT_CONSTANT(opline, opline.GetOp2())
			ce = ZendFetchClassByName(class_name.GetStr(), (class_name + 1).GetStr(), opline.GetOp1().GetNum())
			CACHE_PTR(opline.GetExtendedValue(), ce)
		}
		opline.GetResultZval().SetCe(ce)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
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
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) && obj == orig_obj {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), called_scope, fbc)
		}
		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			return 0
		}
	}

	if CACHED_PTR(opline.GetResult().GetNum()) == ce {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		function_name = RT_CONSTANT(opline, opline.GetOp2())
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
		}
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
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CONSTANT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var c *ZendConstant
	c = CACHED_PTR(opline.GetExtendedValue())
	if c != nil && IS_SPECIAL_CACHE_VAL(c) == 0 {
		types.ZVAL_COPY_OR_DUP(opline.GetResultZval(), c.Value())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	ZendQuickGetConstant(RT_CONSTANT(opline, opline.GetOp2())+1, opline.GetOp1().GetNum(), opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *ZendClassConstant
	var value *types.Zval
	var zv *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {

		{
			{
				ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
				if ce == nil {
					b.Assert(EG__().GetException() != nil)
					opline.GetResultZval().SetUndef()
					return 0
				}
			}

			if CACHED_PTR(opline.GetExtendedValue()) == ce {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			}
		}
		zv = ce.GetConstantsTable().KeyFind(RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetStr())
		if zv != nil {
			c = zv.GetPtr()
			scope = executeData.GetFunc().op_array.scope
			if ZendVerifyConstAccess(c, scope) == 0 {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
				opline.GetResultZval().SetUndef()
				return 0
			}
			value = c.GetValue()
			if value.IsConstant() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					opline.GetResultZval().SetUndef()
					return 0
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
			opline.GetResultZval().SetUndef()
			return 0
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(opline.GetResultZval(), value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.GetResultZval()

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = RT_CONSTANT(opline, opline.GetOp2())
	for {
		types.Z_OBJ_HT_P(container).GetUnsetProperty()(container, offset, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = RT_CONSTANT(opline, opline.GetOp2())
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
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
		var key *types.Zval = RT_CONSTANT(opline, opline.GetOp2())

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			generator.GetKey().TryAddRefcount()

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
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
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
	property = opline.getZvalPtrVar2(&free_op2)
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
					opline.GetResultZval().SetNull()
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
					types.ZVAL_COPY(opline.GetResultZval(), zptr)
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, executeData)
		}
		break
	}
	FREE_OP(free_op_data)
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	for {
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					opline.GetResultZval().SetNull()
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
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	for {
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				opline.GetResultZval().SetNull()
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
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.getZvalPtrVar2(&free_op2)
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
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.getZvalPtrVar2(&free_op2)
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
assign_object:
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	{

		{
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, executeData)
		}
	}

	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.getZvalPtrVar2(&free_op2)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)
	{

		{
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, executeData)
		}
	}

	ZvalPtrDtorNogc(free_op2)
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(opline.GetResultZval())

	{
		var_ = opline.getZvalPtrVar2(&free_op2)
		if var_.IsString() {
			{
				rope[0] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[0] = ZvalGetStringFunc(var_)
			ZvalPtrDtorNogc(free_op2)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var free_op2 ZendFreeOp
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		class_name = opline.getZvalPtrVar2(&free_op2)
	try_class_name:
		if class_name.IsObject() {
			opline.GetResultZval().SetCe(types.Z_OBJCE_P(class_name))
		} else if class_name.IsString() {
			opline.GetResultZval().SetCe(ZendFetchClass(class_name.GetStr(), opline.GetOp1().GetNum()))
		} else if class_name.IsReference() {
			class_name = types.Z_REFVAL_P(class_name)
			goto try_class_name
		} else {
			if class_name.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
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
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
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

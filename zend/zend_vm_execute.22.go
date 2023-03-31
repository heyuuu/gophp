package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_YIELD_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
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
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_RETURN_REFERENCE) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

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
func ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var variable_ptr *types.Zval
	var value_type uint32
	var fe_ht *types.Array
	var pos types.ArrayPosition
	var p *types.Bucket
	array = EX_VAR(opline.GetOp1().GetVar())
	fe_ht = array.GetArr()
	pos = array.GetFePos()
	p = fe_ht.Bucket(pos)
	for true {
		if pos >= fe_ht.GetNNumUsed() {

			/* reached end of iteration */

			ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
			return 0
		}
		value = p.GetVal()
		value_type = value.GetTypeInfo()
		if value_type != types.IS_UNDEF {
			if value_type == types.IS_INDIRECT {
				value = value.GetZv()
				value_type = value.GetTypeInfo()
				if value_type != types.IS_UNDEF {
					break
				}
			} else {
				break
			}
		}
		pos++
		p++
	}
	array.SetFePos(pos + 1)

	variable_ptr = EX_VAR(opline.GetOp2().GetVar())
	ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var variable_ptr *types.Zval
	var value_type uint32
	var fe_ht *types.Array
	var pos types.ArrayPosition
	var p *types.Bucket
	array = EX_VAR(opline.GetOp1().GetVar())
	fe_ht = array.GetArr()
	pos = array.GetFePos()
	p = fe_ht.Bucket(pos)
	for true {
		if pos >= fe_ht.GetNNumUsed() {

			/* reached end of iteration */

			ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
			return 0
		}
		value = p.GetVal()
		value_type = value.GetTypeInfo()
		if value_type != types.IS_UNDEF {
			if value_type == types.IS_INDIRECT {
				value = value.GetZv()
				value_type = value.GetTypeInfo()
				if value_type != types.IS_UNDEF {
					break
				}
			} else {
				break
			}
		}
		pos++
		p++
	}
	array.SetFePos(pos + 1)
	if p.GetKey() == nil {
		EX_VAR(opline.GetResult().GetVar()).SetLong(p.GetH())
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(p.GetKey())
	}
	variable_ptr = EX_VAR(opline.GetOp2().GetVar())
	ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_ADD_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	} else {
		ZEND_DEL_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_ADD_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	} else {
		ZEND_DEL_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CLONE_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone *types.ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = &(executeData.GetThis())
	if obj.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	for {
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
func ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type uint32
	var called_scope *types.ClassEntry
	var scope *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	fetch_type = opline.GetOp1().GetNum()
	scope = executeData.GetFunc().op_array.scope
	if scope == nil {
		faults.ThrowError(nil, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(scope.GetName())
	case ZEND_FETCH_CLASS_PARENT:
		if scope.GetParent() == nil {
			faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(scope.GetParent().name)
	case ZEND_FETCH_CLASS_STATIC:
		if executeData.GetThis().IsObject() {
			called_scope = types.Z_OBJCE(executeData.GetThis())
		} else {
			called_scope = executeData.GetThis().GetCe()
		}
		EX_VAR(opline.GetResult().GetVar()).SetStringCopy(called_scope.GetName())
	default:

	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
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
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
	assign_op_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR((opline + 1).GetExtendedValue())
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
						prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
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
func ZEND_PRE_INC_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
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
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		}

		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {
				{
					prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
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
func ZEND_POST_INC_OBJ_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
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
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		}

		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			} else {
				{
					prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
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
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = RT_CONSTANT(opline, opline.GetOp2())
	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_REF)
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if !retval.IsUndef() {

					{
					fetch_obj_r_fast_copy:
						types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
						ZEND_VM_NEXT_OPCODE()
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
								goto fetch_obj_r_fast_copy
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
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
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
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(executeData)
}
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = RT_CONSTANT(opline, opline.GetOp2())
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
					fetch_obj_is_fast_copy:
						types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
						ZEND_VM_NEXT_OPCODE()
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
								goto fetch_obj_is_fast_copy
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
						goto fetch_obj_is_fast_copy
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_UNUSED, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
assign_object:
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.IsNotUndef() {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_CONST, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				{

					value.TryAddRefcount()

				}

				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
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
	property = RT_CONSTANT(opline, opline.GetOp2())
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
assign_object:
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.IsNotUndef() {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
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
	property = RT_CONSTANT(opline, opline.GetOp2())
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
assign_object:
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.IsNotUndef() {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_VAR, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
assign_object:
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.IsNotUndef() {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_CV, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	{
		{
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, executeData)
		}

	}

	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = RT_CONSTANT(opline, opline.GetOp2())
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)
	{
		{
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, executeData)
		}

	}

	OPLINE = executeData.GetOpline() + 2
	return 0
}

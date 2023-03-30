package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_GET_TYPE_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var type_ *types.String
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		EX_VAR(opline.GetResult().GetVar()).SetInternedString(type_)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetStringVal("unknown type")
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_ROPE_ADD_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))

	{
		var_ = EX_VAR(opline.GetOp2().GetVar())
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ROPE_END_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))

	{
		var_ = EX_VAR(opline.GetOp2().GetVar())
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			if EG__().GetException() != nil {
				for i = 0; i <= opline.GetExtendedValue(); i++ {
					types.ZendStringReleaseEx(rope[i], 0)
				}
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
		}
	}
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		len_ += rope[i].GetLen()
	}
	ret = EX_VAR(opline.GetResult().GetVar())
	ret.SetString(types.ZendStringAlloc(len_, 0))
	target = ret.GetStr().GetVal()
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		memcpy(target, rope[i].GetVal(), rope[i].GetLen())
		target += rope[i].GetLen()
		types.ZendStringReleaseEx(rope[i], 0)
	}
	*target = '0'
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = nil
		if expr_ptr.IsReference() {
			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	}
	{
		var offset *types.Zval = EX_VAR(opline.GetOp2().GetVar())
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
			EX_VAR(opline.GetResult().GetVar()).GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdate(hval, expr_ptr)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else if offset.IsDouble() {
			hval = ZendDvalToLval(offset.GetDval())
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
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER(executeData)
	}

}
func ZEND_YIELD_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
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
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

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
func ZEND_BIND_LEXICAL_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var closure *types.Zval
	var var_ *types.Zval
	closure = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {

		/* By-ref binding */

		var_ = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), executeData)
		if var_.IsReference() {
			var_.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(var_, 2)
		}
	} else {
		var_ = EX_VAR(opline.GetOp2().GetVar())
		if var_.IsUndef() && (opline.GetExtendedValue()&ZEND_BIND_IMPLICIT) == 0 {
			var_ = ZVAL_UNDEFINED_OP2()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		var_ = types.ZVAL_DEREF(var_)
		var_.TryAddRefcount()
	}
	ZendClosureBindVarEx(closure, opline.GetExtendedValue() & ^(ZEND_BIND_REF|ZEND_BIND_IMPLICIT), var_)
	ZEND_VM_NEXT_OPCODE()
}
func zend_pre_inc_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_INC_SPEC_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)

		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_inc_helper_SPEC_VAR(executeData)
}
func ZEND_PRE_INC_SPEC_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_inc_helper_SPEC_VAR(executeData)
}
func zend_pre_dec_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_DEC_SPEC_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongDecrementFunction(var_ptr)

		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_dec_helper_SPEC_VAR(executeData)
}
func ZEND_PRE_DEC_SPEC_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongDecrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_pre_dec_helper_SPEC_VAR(executeData)
}
func zend_post_inc_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
		FastLongIncrementFunction(var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_post_inc_helper_SPEC_VAR(executeData)
}
func zend_post_dec_helper_SPEC_VAR(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_DEC_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
		FastLongDecrementFunction(var_ptr)
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_post_dec_helper_SPEC_VAR(executeData)
}
func ZEND_RETURN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	return_value = executeData.GetReturnValue()
	if retval_ptr.IsUndef() {
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
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var free_op1 ZendFreeOp
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
			if !(executeData.GetReturnValue()) {
				ZvalPtrDtorNogc(free_op1)
			} else {
				if retval_ptr.IsReference() {
					types.ZVAL_COPY_VALUE(executeData.GetReturnValue(), retval_ptr)
					break
				}
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
			}
			break
		}
		retval_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		{
			b.Assert(retval_ptr != EG__().GetUninitializedZval())
			if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(retval_ptr.IsReference()) {
				faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
				if executeData.GetReturnValue() {
					executeData.GetReturnValue().
						SetNewRef(retval_ptr)
				} else {
					if free_op1 != nil {
						ZvalPtrDtorNogc(free_op1)
					}
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
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_GENERATOR_RETURN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var free_op1 ZendFreeOp
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_THROW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	faults.ThrowExceptionObject(value)
	faults.ExceptionRestore()
	ZvalPtrDtorNogc(free_op1)
	HANDLE_EXCEPTION()
}
func ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if varptr.IsUndef() {
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
func ZEND_SEND_VAR_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData)
}
func ZEND_SEND_VAR_NO_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, varptr)
	if varptr.IsReference() {
		ZEND_VM_NEXT_OPCODE()
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {
		return ZEND_SEND_VAR_SPEC_VAR_HANDLER(executeData)
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, varptr)
	if varptr.IsReference() || ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_VM_NEXT_OPCODE()
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {
		return ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData)
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, varptr)
	if varptr.IsReference() || QUICK_ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_VM_NEXT_OPCODE()
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	varptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAR_EX_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_var_by_ref:
		return ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData)
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if varptr.IsUndef() {
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
func ZEND_SEND_VAR_EX_SPEC_VAR_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_var_by_ref
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if varptr.IsUndef() {
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
func ZEND_SEND_FUNC_ARG_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		return ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData)
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	if varptr.IsReference() {
		var ref *types.ZendRefcounted = varptr.GetCounted()
		varptr = types.Z_REFVAL_P(varptr)
		types.ZVAL_COPY_VALUE(arg, varptr)
		if ref.DelRefcount() == 0 {
			EfreeSize(ref, b.SizeOf("zend_reference"))
		} else {
			arg.TryAddRefcount()
		}

	} else {
		types.ZVAL_COPY_VALUE(arg, varptr)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_USER_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	var free_op1 ZendFreeOp
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CAST_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	var ht *types.Array
	expr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
			ZvalPtrDtorNogc(free_op1)
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			if expr.GetType() != types.IS_OBJECT || types.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
					expr = result.GetArr().IndexAddNew(0, expr)

					{

						expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(expr).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
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
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.ZSTR_SCALAR.GetStr(), expr)

				{

					expr.TryAddRefcount()

				}
			}
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FE_RESET_R_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = EX_VAR(opline.GetResult().GetVar())
		types.ZVAL_COPY_VALUE(result, array_ptr)
		result.SetFePos(0)
		ZvalPtrDtorNogc(free_op1)
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
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				ZvalPtrDtorNogc(free_op1)
				ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(types.ZendHashIteratorAdd(properties, 0))
			ZvalPtrDtorNogc(free_op1)
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

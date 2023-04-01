package zend

import (
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_QM_ASSIGN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		result.SetNull()
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	{
		types.ZVAL_COPY_DEREF(result, value)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_YIELD_FROM_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		ZvalPtrDtorNogc(free_op1)
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		generator.GetValues().GetFePos() = 0
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			if new_gen.GetRetval().IsUndef() {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					faults.ThrowError(nil, "Impossible to yield from the Generator being currently run")
					ZvalPtrDtor(val)
					UNDEF_RESULT()
					HANDLE_EXCEPTION()
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				faults.ThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				ZvalPtrDtor(val)
				UNDEF_RESULT()
				HANDLE_EXCEPTION()
			} else {
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), new_gen.GetRetval())
				}
				ZEND_VM_NEXT_OPCODE()
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			ZvalPtrDtorNogc(free_op1)
			if iter == nil || EG__().GetException() != nil {
				if EG__().GetException() == nil {
					faults.ThrowError(nil, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				UNDEF_RESULT()
				HANDLE_EXCEPTION()
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG__().GetException() != nil {
					OBJ_RELEASE(iter.GetStd())
					UNDEF_RESULT()
					HANDLE_EXCEPTION()
				}
			}
			generator.GetValues().SetObject(iter.GetStd())
		}
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		ZvalPtrDtorNogc(free_op1)
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
func ZEND_IS_IDENTICAL_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_ROPE_ADD_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))
	{
		var_ = RT_CONSTANT(opline, opline.GetOp2())
		rope[opline.GetExtendedValue()] = var_.GetStr()

		var_.TryAddRefcount()

	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ROPE_END_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))
	{
		var_ = RT_CONSTANT(opline, opline.GetOp2())
		rope[opline.GetExtendedValue()] = var_.GetStr()

		var_.TryAddRefcount()

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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
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
		var offset *types.Zval = RT_CONSTANT(opline, opline.GetOp2())
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdate(hval, expr_ptr)
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
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER(executeData)
	}

}
func ZEND_YIELD_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

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
func ZEND_IN_ARRAY_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var ht *types.Array = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	var result *types.Zval
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStr().GetStr())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.GetLval())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
		result = ht.KeyFind(types.ZSTR_EMPTY_ALLOC().GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.GetLval() == 0 {
				result = val
				break
			}
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != nil)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))

	{
		var_ = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			ZvalPtrDtorNogc(free_op2)
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(EX_VAR(opline.GetOp1().GetVar()))

	{
		var_ = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			ZvalPtrDtorNogc(free_op2)
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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
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
		var free_op2 ZendFreeOp
		var offset *types.Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_IS_IDENTICAL_SPEC_TMP_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_TMP_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_TMP_TMP_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

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
func ZEND_YIELD_SPEC_TMP_VAR_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

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
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
func ZEND_VERIFY_RETURN_TYPE_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		if EX_VAR(opline.GetResult().GetVar()).GetArr().NextIndexInsert(expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER(executeData)
	}

}
func ZEND_YIELD_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

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

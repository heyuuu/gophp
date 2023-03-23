// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_FE_RESET_RW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
			if free_op1 != nil {
				ZvalPtrDtorNogc(free_op1)
			}
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
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
				ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(types.ZendHashIteratorAdd(properties, 0))
			if free_op1 != nil {
				ZvalPtrDtorNogc(free_op1)
			}
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
			var is_empty types.ZendBool = ZendFeResetIterator(array_ptr, 1, opline, executeData)
			{
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
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
			if free_op1 != nil {
				ZvalPtrDtorNogc(free_op1)
			}
		}

		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_FETCH_R_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var value_type uint32
	var fe_ht *types.Array
	var pos types.HashPosition
	var p *types.Bucket
	array = EX_VAR(opline.GetOp1().GetVar())
	if array.IsArray() {
		fe_ht = array.GetArr()
		pos = array.GetFePos()
		p = fe_ht.GetArData() + pos
		for true {
			if pos >= fe_ht.GetNNumUsed() {

				/* reached end of iteration */

			fe_fetch_r_exit:
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
		if RETURN_VALUE_USED(opline) {
			if p.GetKey() == nil {
				EX_VAR(opline.GetResult().GetVar()).SetLong(p.GetH())
			} else {
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(p.GetKey())
			}
		}
	} else {
		var iter *ZendObjectIterator
		b.Assert(array.IsObject())
		if b.Assign(&iter, ZendIteratorUnwrap(array)) == nil {

			/* plain object */

			fe_ht = types.Z_OBJPROP_P(array)
			pos = types.ZendHashIteratorPos(array.GetFeIterIdx(), fe_ht)
			p = fe_ht.GetArData() + pos
			for true {
				if pos >= fe_ht.GetNNumUsed() {

					/* reached end of iteration */

					goto fe_fetch_r_exit

					/* reached end of iteration */

				}
				value = p.GetVal()
				value_type = value.GetTypeInfo()
				if value_type != types.IS_UNDEF {
					if value_type == types.IS_INDIRECT {
						value = value.GetZv()
						value_type = value.GetTypeInfo()
						if value_type != types.IS_UNDEF && ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 0) == types.SUCCESS {
							break
						}
					} else if types.Z_OBJCE_P(array).GetDefaultPropertiesCount() == 0 || p.GetKey() == nil || ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 1) == types.SUCCESS {
						break
					}
				}
				pos++
				p++
			}
			if RETURN_VALUE_USED(opline) {
				if p.GetKey() == nil {
					EX_VAR(opline.GetResult().GetVar()).SetLong(p.GetH())
				} else if p.GetKey().GetVal()[0] {
					EX_VAR(opline.GetResult().GetVar()).SetStringCopy(p.GetKey())
				} else {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					ZendUnmanglePropertyNameEx(p.GetKey(), &class_name, &prop_name, &prop_name_len)
					EX_VAR(opline.GetResult().GetVar()).SetRawString(b.CastStr(prop_name, prop_name_len))
				}
			}
			EG__().GetHtIterators()[types.Z_FE_ITER_P(array)].SetPos(pos + 1)
		} else {
			if b.PreInc(&(iter.GetIndex())) > 0 {

				/* This could cause an endless loop if index becomes zero again.
				 * In case that ever happens we need an additional flag. */

				iter.GetFuncs().GetMoveForward()(iter)
				if EG__().GetException() != nil {
					UNDEF_RESULT()
					HANDLE_EXCEPTION()
				}
				if iter.GetFuncs().GetValid()(iter) == types.FAILURE {

					/* reached end of iteration */

					if EG__().GetException() != nil {
						UNDEF_RESULT()
						HANDLE_EXCEPTION()
					}
					goto fe_fetch_r_exit
				}
			}
			value = iter.GetFuncs().GetGetCurrentData()(iter)
			if EG__().GetException() != nil {
				UNDEF_RESULT()
				HANDLE_EXCEPTION()
			}
			if value == nil {

				/* failure in get_current_data */

				goto fe_fetch_r_exit

				/* failure in get_current_data */

			}
			if RETURN_VALUE_USED(opline) {
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					iter.GetFuncs().GetGetCurrentKey()(iter, EX_VAR(opline.GetResult().GetVar()))
					if EG__().GetException() != nil {
						UNDEF_RESULT()
						HANDLE_EXCEPTION()
					}
				} else {
					EX_VAR(opline.GetResult().GetVar()).SetLong(iter.GetIndex())
				}
			}
			value_type = value.GetTypeInfo()
		}
	}
	if opline.GetOp2Type() == IS_CV {
		var variable_ptr *types.Zval = EX_VAR(opline.GetOp2().GetVar())
		ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
	} else {
		var res *types.Zval = EX_VAR(opline.GetOp2().GetVar())
		var gc *types.ZendRefcounted = value.GetCounted()
		types.ZVAL_COPY_VALUE_EX(res, value, gc, value_type)
		if types.Z_TYPE_INFO_REFCOUNTED(value_type) {
			gc.AddRefcount()
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FE_FETCH_RW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var value_type uint32
	var fe_ht *types.Array
	var pos types.HashPosition
	var p *types.Bucket
	array = EX_VAR(opline.GetOp1().GetVar())
	array = types.ZVAL_DEREF(array)
	if array.IsArray() {
		pos = types.ZendHashIteratorPosEx(EX_VAR(opline.GetOp1().GetVar()).GetFeIterIdx(), array)
		fe_ht = array.GetArr()
		p = fe_ht.GetArData() + pos
		for true {
			if pos >= fe_ht.GetNNumUsed() {

				/* reached end of iteration */

				goto fe_fetch_w_exit

				/* reached end of iteration */

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
		if RETURN_VALUE_USED(opline) {
			if p.GetKey() == nil {
				EX_VAR(opline.GetResult().GetVar()).SetLong(p.GetH())
			} else {
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(p.GetKey())
			}
		}
		EG__().GetHtIterators()[types.Z_FE_ITER_P(EX_VAR(opline.GetOp1().GetVar()))].SetPos(pos + 1)
	} else if array.IsObject() {
		var iter *ZendObjectIterator
		if b.Assign(&iter, ZendIteratorUnwrap(array)) == nil {

			/* plain object */

			fe_ht = types.Z_OBJPROP_P(array)
			pos = types.ZendHashIteratorPos(EX_VAR(opline.GetOp1().GetVar()).GetFeIterIdx(), fe_ht)
			p = fe_ht.GetArData() + pos
			for true {
				if pos >= fe_ht.GetNNumUsed() {

					/* reached end of iteration */

					goto fe_fetch_w_exit

					/* reached end of iteration */

				}
				value = p.GetVal()
				value_type = value.GetTypeInfo()
				if value_type != types.IS_UNDEF {
					if value_type == types.IS_INDIRECT {
						value = value.GetZv()
						value_type = value.GetTypeInfo()
						if value_type != types.IS_UNDEF && ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 0) == types.SUCCESS {
							if (value_type & types.Z_TYPE_MASK) != types.IS_REFERENCE {
								var prop_info *ZendPropertyInfo = ZendGetTypedPropertyInfoForSlot(array.GetObj(), value)
								if prop_info != nil {
									value.SetNewRef(value)
									ZEND_REF_ADD_TYPE_SOURCE(value.GetRef(), prop_info)
									value_type = types.IS_REFERENCE_EX
								}
							}
							break
						}
					} else if types.Z_OBJCE_P(array).GetDefaultPropertiesCount() == 0 || p.GetKey() == nil || ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 1) == types.SUCCESS {
						break
					}
				}
				pos++
				p++
			}
			if RETURN_VALUE_USED(opline) {
				if p.GetKey() == nil {
					EX_VAR(opline.GetResult().GetVar()).SetLong(p.GetH())
				} else if p.GetKey().GetVal()[0] {
					EX_VAR(opline.GetResult().GetVar()).SetStringCopy(p.GetKey())
				} else {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					ZendUnmanglePropertyNameEx(p.GetKey(), &class_name, &prop_name, &prop_name_len)
					EX_VAR(opline.GetResult().GetVar()).SetRawString(b.CastStr(prop_name, prop_name_len))
				}
			}
			EG__().GetHtIterators()[types.Z_FE_ITER_P(EX_VAR(opline.GetOp1().GetVar()))].SetPos(pos + 1)
		} else {
			if b.PreInc(&(iter.GetIndex())) > 0 {

				/* This could cause an endless loop if index becomes zero again.
				 * In case that ever happens we need an additional flag. */

				iter.GetFuncs().GetMoveForward()(iter)
				if EG__().GetException() != nil {
					UNDEF_RESULT()
					HANDLE_EXCEPTION()
				}
				if iter.GetFuncs().GetValid()(iter) == types.FAILURE {

					/* reached end of iteration */

					if EG__().GetException() != nil {
						UNDEF_RESULT()
						HANDLE_EXCEPTION()
					}
					goto fe_fetch_w_exit
				}
			}
			value = iter.GetFuncs().GetGetCurrentData()(iter)
			if EG__().GetException() != nil {
				UNDEF_RESULT()
				HANDLE_EXCEPTION()
			}
			if value == nil {

				/* failure in get_current_data */

				goto fe_fetch_w_exit

				/* failure in get_current_data */

			}
			if RETURN_VALUE_USED(opline) {
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					iter.GetFuncs().GetGetCurrentKey()(iter, EX_VAR(opline.GetResult().GetVar()))
					if EG__().GetException() != nil {
						UNDEF_RESULT()
						HANDLE_EXCEPTION()
					}
				} else {
					EX_VAR(opline.GetResult().GetVar()).SetLong(iter.GetIndex())
				}
			}
			value_type = value.GetTypeInfo()
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		if EG__().GetException() != nil {
			UNDEF_RESULT()
			HANDLE_EXCEPTION()
		}
	fe_fetch_w_exit:
		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	}
	if (value_type & types.Z_TYPE_MASK) != types.IS_REFERENCE {
		var gc *types.ZendRefcounted = value.GetCounted()
		var ref *types.Zval
		value.SetNewEmptyRef()
		ref = types.Z_REFVAL_P(value)
		types.ZVAL_COPY_VALUE_EX(ref, value, gc, value_type)
	}
	if opline.GetOp2Type() == IS_CV {
		var variable_ptr *types.Zval = EX_VAR(opline.GetOp2().GetVar())
		if variable_ptr != value {
			var ref *types.ZendReference
			ref = value.GetRef()
			ref.AddRefcount()
			IZvalPtrDtor(variable_ptr)
			variable_ptr.SetReference(ref)
		}
	} else {
		value.AddRefcount()
		EX_VAR(opline.GetOp2().GetVar()).SetReference(value.GetRef())
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_JMP_SET_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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

			result.TryAddRefcount()

		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_COALESCE_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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

			result.TryAddRefcount()

		}

		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_QM_ASSIGN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
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
func ZEND_YIELD_FROM_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		ZvalPtrDtorNogc(free_op1)
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		{
		}

		generator.GetValues().GetFePos() = 0
		ZvalPtrDtorNogc(free_op1)
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			{
			}

			ZvalPtrDtorNogc(free_op1)
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
func ZEND_SEND_VAR_SIMPLE_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY(arg, varptr)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAR_EX_SIMPLE_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		return ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData)
	}
	varptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY(arg, varptr)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_IDENTICAL_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto assign_op_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = RT_CONSTANT(opline, opline.GetOp2())

		{
			{
				var_ptr = zend_fetch_dimension_address_inner_RW_CONST(container.GetArr(), dim, executeData)
			}

			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = RT_CONSTANT(opline, opline.GetOp2())
		if container.IsObject() {
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewZendArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OP_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
	} else {
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_INC_OBJ_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto pre_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_OBJ_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	{
	}

	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto post_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
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
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_W_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_W(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_RW(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}

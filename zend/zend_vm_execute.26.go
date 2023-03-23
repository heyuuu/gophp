// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_SEND_VAR_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_var_by_ref:
		return ZEND_SEND_REF_SPEC_CV_HANDLER(executeData)
	}
	varptr = EX_VAR(opline.GetOp1().GetVar())
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
func ZEND_SEND_VAR_EX_SPEC_CV_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_var_by_ref
	}
	varptr = EX_VAR(opline.GetOp1().GetVar())
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
func ZEND_SEND_USER_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = EX_VAR(opline.GetOp1().GetVar())
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) != 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CLONE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = EX_VAR(opline.GetOp1().GetVar())
	{
	}

	for {
		if obj.GetType() != types.IS_OBJECT {
			if obj.IsReference() {
				obj = types.Z_REFVAL_P(obj)
				if obj.IsObject() {
					break
				}
			}
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			if obj.IsUndef() {
				ZVAL_UNDEFINED_OP1()
				if EG__().GetException() != nil {
					HANDLE_EXCEPTION()
				}
			}
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
func ZEND_CAST_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	var ht *types.Array
	expr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
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

						expr.TryAddRefcount()

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

					expr.TryAddRefcount()

				}
			}
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INCLUDE_OR_EVAL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *ZendOpArray
	var inc_filename *types.Zval
	inc_filename = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
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
func ZEND_FE_RESET_R_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
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
		ZEND_VM_JMP(OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_RW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
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
func ZEND_JMP_SET_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	var ret int
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
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
func ZEND_COALESCE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	value = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), executeData)
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
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_QM_ASSIGN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	value = EX_VAR(opline.GetOp1().GetVar())
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
func ZEND_YIELD_FROM_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	val = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		{
		}

		generator.GetValues().GetFePos() = 0
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			{
			}

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
func ZEND_STRLEN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	if value.IsString() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetStr().GetLen())
		ZEND_VM_NEXT_OPCODE()
	} else {
		var strict types.ZendBool
		if value.IsReference() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetStr().GetLen())
				ZEND_VM_NEXT_OPCODE()
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1()
		}
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
func ZEND_TYPE_CHECK_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	value = EX_VAR(opline.GetOp1().GetVar())
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else if value.IsReference() {
		value = types.Z_REFVAL_P(value)
		if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
			goto type_check_resource
		}
	} else if value.IsUndef() {
		result = (1 << types.IS_NULL & opline.GetExtendedValue()) != 0
		ZVAL_UNDEFINED_OP1()
		if EG__().GetException() != nil {
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
	}
	{
		ZEND_VM_SMART_BRANCH(result, 1)
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}

}
func ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	var_ptr.GetLval()++

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	var_ptr.GetLval()++
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	FastLongIncrementFunction(var_ptr)

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	FastLongIncrementFunction(var_ptr)
	types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	var_ptr.GetLval()--

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	var_ptr.GetLval()--
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	FastLongDecrementFunction(var_ptr)

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	FastLongDecrementFunction(var_ptr)
	types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), var_ptr)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_POST_INC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	var_ptr.GetLval()++
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_POST_INC_LONG_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	FastLongIncrementFunction(var_ptr)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_POST_DEC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	var_ptr.GetLval()--
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_POST_DEC_LONG_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetLong(var_ptr.GetLval())
	FastLongDecrementFunction(var_ptr)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAR_SIMPLE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	varptr = EX_VAR(opline.GetOp1().GetVar())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY(arg, varptr)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SEND_VAR_EX_SIMPLE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		return ZEND_SEND_REF_SPEC_CV_HANDLER(executeData)
	}
	varptr = EX_VAR(opline.GetOp1().GetVar())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY(arg, varptr)
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_DIV_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	if op1.IsString() {
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
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			EX_VAR(opline.GetResult().GetVar()).SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
			{
			}

		}
		ZEND_VM_NEXT_OPCODE()
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		{
		}

		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_IS_IDENTICAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_EQUAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			{
			}

			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			{
			}

			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			{
			}

			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}

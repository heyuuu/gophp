package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = EX_VAR(opline.GetOp1().GetVar())
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = EX_VAR(opline.GetOp1().GetVar())
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_VERIFY_RETURN_TYPE_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsReference() {
			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	}

	{
		if EX_VAR(opline.GetResult().GetVar()).GetArr().NextIndexInsert(expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(executeData)
	}

}
func ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ *types.Zval = EX_VAR(opline.GetOp1().GetVar())
	if var_.IsRefcounted() {
		var garbage *types.ZendRefcounted = var_.GetCounted()
		var_.SetUndef()
		if garbage.DelRefcount() == 0 {
			RcDtorFunc(garbage)
		} else {
			GcCheckPossibleRoot(garbage)
		}
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	} else {
		var_.SetUndef()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = EX_VAR(opline.GetOp1().GetVar())

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			HANDLE_EXCEPTION()
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_SET_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	if value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL) {
		ZEND_VM_SMART_BRANCH_TRUE()
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_NEXT_OPCODE()
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		ZEND_VM_NEXT_OPCODE()
	}
}
func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_EMPTY_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	var result int
	result = !(IZendIsTrue(value))
	if EG__().GetException() != nil {
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	ZEND_VM_SMART_BRANCH(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), executeData)

	{
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	if value == nil {
		result = opline.GetExtendedValue() & ZEND_ISEMPTY
	} else {
		if value.IsIndirect() {
			value = value.GetZv()
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
			if value.IsReference() {
				value = types.Z_REFVAL_P(value)
			}
			result = value.GetType() > types.IS_NULL
		} else {
			result = !(IZendIsTrue(value))
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INSTANCEOF_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result types.ZendBool
	expr = EX_VAR(opline.GetOp1().GetVar())
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_RETURN_REFERENCE) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)

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
func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var value *types.Zval
	var variable_ptr *types.Zval
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
	IZvalPtrDtor(variable_ptr)
	ht = ZEND_MAP_PTR_GET(executeData.GetFunc().op_array.static_variables_ptr)
	if ht == nil {
		b.Assert((executeData.GetFunc().op_array.fn_flags & (ZEND_ACC_IMMUTABLE | ZEND_ACC_PRELOADED)) != 0)
		ht = types.ZendArrayDup(executeData.GetFunc().op_array.static_variables)
		ZEND_MAP_PTR_SET(executeData.GetFunc().op_array.static_variables_ptr, ht)
	} else if ht.GetRefcount() > 1 {
		if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
			ht.DelRefcount()
		}
		ht = types.ZendArrayDup(ht)
		ZEND_MAP_PTR_SET(executeData.GetFunc().op_array.static_variables_ptr, ht)
	}
	value = (*types.Zval)((*byte)(ht.GetArData() + (opline.GetExtendedValue() & ^(ZEND_BIND_REF | ZEND_BIND_IMPLICIT))))
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {
		if value.IsConstant() {
			if ZvalUpdateConstantEx(value, executeData.GetFunc().op_array.scope) != types.SUCCESS {
				variable_ptr.SetNull()
				HANDLE_EXCEPTION()
			}
		}
		if !(value.IsReference()) {
			var ref *types.ZendReference = (*types.ZendReference)(Emalloc(b.SizeOf("zend_reference")))
			ref.SetRefcount(2)
			ref.GetGcTypeInfo() = types.IS_REFERENCE
			types.ZVAL_COPY_VALUE(ref.GetVal(), value)
			ref.GetSources().SetPtr(nil)
			value.SetTypeReference()
			variable_ptr.SetReference(ref)
		} else {
			value.AddRefcount()
			variable_ptr.SetReference(value.GetRef())
		}
	} else {
		types.ZVAL_COPY(variable_ptr, value)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CHECK_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = EX_VAR(opline.GetOp1().GetVar())
	if op1.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = EX_VAR(opline.GetOp1().GetVar())
	{
		if op1.IsUndef() {
			op1.SetNewEmptyRef()
			op1.SetRefcount(2)
			types.Z_REFVAL_P(op1).SetNull()
			EX_VAR(opline.GetResult().GetVar()).SetReference(op1.GetRef())
		} else {
			if op1.IsReference() {
				op1.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(op1, 2)
			}
			EX_VAR(opline.GetResult().GetVar()).SetReference(op1.GetRef())
		}
	}

	ZEND_VM_NEXT_OPCODE()
}
func ZEND_COUNT_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var count ZendLong
	op1 = EX_VAR(opline.GetOp1().GetVar())
	for true {
		if op1.IsArray() {
			count = op1.GetArr().Count()
			break
		} else if op1.IsObject() {

			/* first, we check if the handler is defined */

			if types.Z_OBJ_HT_P(op1).GetCountElements() != nil {
				if types.SUCCESS == types.Z_OBJ_HT_P(op1).GetCountElements()(op1, &count) {
					break
				}
				if EG__().GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if InstanceofFunction(types.Z_OBJCE_P(op1), ZendCeCountable) != 0 {
				var retval types.Zval
				ZendCallMethodWith0Params(op1, nil, nil, "count", &retval)
				count = ZvalGetLong(&retval)
				ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
			continue
		} else if op1.GetType() <= types.IS_NULL {
			if op1.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			count = 0
		} else {
			count = 1
		}
		faults.Error(faults.E_WARNING, "%s(): Parameter must be an array or an object that implements Countable", b.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	EX_VAR(opline.GetResult().GetVar()).SetLong(count)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = EX_VAR(opline.GetOp1().GetVar())
		for true {
			if op1.IsObject() {
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1()
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
			}
			break
		}
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var type_ *types.String
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		EX_VAR(opline.GetResult().GetVar()).SetInternedString(type_)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetStringVal("unknown type")
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_DIV_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	if (op1.IsString()) && (op2.IsString()) {
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
		} else if op2_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op1_str)
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
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
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		ZEND_VM_NEXT_OPCODE()
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_IS_IDENTICAL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_EQUAL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_EQUAL_SPEC_CV_CV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}

// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_VERIFY_RETURN_TYPE_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	{
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_NEW_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var result *types.Zval
	var constructor *ZendFunction
	var ce *types.ClassEntry
	var call *ZendExecuteData
	{
		ce = CACHED_PTR(opline.GetOp2().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
			CACHE_PTR(opline.GetOp2().GetNum(), ce)
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

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION, (*ZendFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		if EX_VAR(opline.GetResult().GetVar()).GetArr().NextIndexInsert(expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewZendArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(executeData)
	}

}
func ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = RT_CONSTANT(opline, opline.GetOp1())
	{
		name = varname.GetStr()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	{
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = RT_CONSTANT(opline, opline.GetOp1())
	{
		name = varname.GetStr()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	value = target_symbol_table.KeyFind(name.GetStr())
	{
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
func ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_ *ZendFunction
	var zfunc *types.Zval
	var object *types.Zval
	var called_scope *types.ClassEntry
	func_ = CACHED_PTR(opline.GetExtendedValue())
	if func_ == nil {
		zfunc = EG__().GetFunctionTable().KeyFind(RT_CONSTANT(opline, opline.GetOp1()).GetStr().GetStr())
		b.Assert(zfunc != nil)
		func_ = zfunc.GetFunc()
		b.Assert(func_.GetType() == ZEND_USER_FUNCTION)
		CACHE_PTR(opline.GetExtendedValue(), func_)
	}
	if executeData.GetThis().IsObject() {
		called_scope = types.Z_OBJCE(executeData.GetThis())
		if func_.IsStatic() || (executeData.GetFunc().common.fn_flags&ZEND_ACC_STATIC) != 0 {
			object = nil
		} else {
			object = &(executeData.GetThis())
		}
	} else {
		called_scope = executeData.GetThis().GetCe()
		object = nil
	}
	ZendCreateClosure(EX_VAR(opline.GetResult().GetVar()), func_, executeData.GetFunc().op_array.scope, called_scope, object)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_YIELD_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
				value = RT_CONSTANT(opline, opline.GetOp1())
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{
					if generator.GetValue().IsRefcounted() {
						generator.GetValue().AddRefcount()
					}
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
				if generator.GetValue().IsRefcounted() {
					generator.GetValue().AddRefcount()
				}
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
func ZEND_COUNT_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var count ZendLong
	op1 = RT_CONSTANT(opline, opline.GetOp1())
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

		} else if op1.GetType() <= types.IS_NULL {
			{
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
func ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = RT_CONSTANT(opline, opline.GetOp1())
		for true {
			if op1.IsObject() {
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else {
				{
				}

				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
			}
			break
		}
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var type_ *types.String
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		EX_VAR(opline.GetResult().GetVar()).SetInternedString(type_)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetRawString("unknown type")
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FUNC_GET_ARGS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = executeData.NumArgs()
	{
		skip = RT_CONSTANT(opline, opline.GetOp1()).GetLval()
		if arg_count < skip {
			result_size = 0
		} else {
			result_size = arg_count - skip
		}
	}

	if result_size != 0 {
		var first_extra_arg uint32 = executeData.GetFunc().op_array.num_args
		ht = types.NewZendArray(result_size)
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
					if q.IsRefcounted() {
						q.AddRefcount()
					}
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
				if q.IsRefcounted() {
					q.AddRefcount()
				}
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
		types.ZVAL_EMPTY_ARRAY(EX_VAR(opline.GetResult().GetVar()))
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_DIV_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
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
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		ZEND_VM_NEXT_OPCODE()
	} else {
		{
		}

		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_SPACESHIP_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	dim = EX_VAR(opline.GetOp2().GetVar())

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_CV, opline, executeData)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_read_IS(container, EX_VAR(opline.GetOp2().GetVar()), IS_CV, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		{
		}

		return ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = RT_CONSTANT(opline, opline.GetOp1())
	{
	}

	offset = EX_VAR(opline.GetOp2().GetVar())
	{
		for {
			{
			}
			{
			}

			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			ZendWrongPropertyRead(offset)
			EX_VAR(opline.GetResult().GetVar()).SetNull()
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
func ZEND_FETCH_OBJ_IS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = RT_CONSTANT(opline, opline.GetOp1())
	{
	}

	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	{
		for {
			{
			}

			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, EX_VAR(opline.GetOp2().GetVar()), IS_CV, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
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
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		ZEND_VM_NEXT_OPCODE()
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
		}

		{
			if op2_str.GetLen() == 0 {
				{
					if op1.IsRefcounted() {
						op1_str.AddRefcount()
					}
				}
				EX_VAR(opline.GetResult().GetVar()).SetString(op1_str)
				types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		EX_VAR(opline.GetResult().GetVar()).SetString(str)
		{
		}

		{
			types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc *ZendFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = RT_CONSTANT(opline, opline.GetOp1())
	{
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
		for {
			{
				{
				}
				{
				}
				{
				}

				ZendInvalidMethodCall(object, function_name)
				HANDLE_EXCEPTION()
			}
			break
		}
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				HANDLE_EXCEPTION()
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
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
		{
		}

		/* previous opcode is ZEND_FETCH_CLASS */

		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_INIT_USER_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ *ZendFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
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
		if EG__().GetException() != nil {
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				ZendObjectRelease(ZEND_CLOSURE_OBJECT(func_))
			} else if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
				ZendObjectRelease(fcc.GetObject())
			}
			HANDLE_EXCEPTION()
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", RT_CONSTANT(opline, opline.GetOp1()).GetStr().GetVal(), error)
		Efree(error)
		if EG__().GetException() != nil {
			HANDLE_EXCEPTION()
		}
		func_ = (*ZendFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
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
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdateH(hval, expr_ptr)
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
